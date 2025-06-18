//go:build stm32h7x7

package sdio

import (
	"runtime"
	"sync"
	"time"
	"unsafe"
	"volatile"

	stm32h7x7 "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/sdmmc"

	corehal "pkg.si-go.dev/chip/core/hal"
	corepin "pkg.si-go.dev/chip/core/hal/pin"
	coresdio "pkg.si-go.dev/chip/core/hal/sdio"
)

const (
	errNone _error = iota
	errCommandFail
	errCommandCrcFail
	errTimeout
	errNotReady
)

func (e _error) Error() string {
	switch e {
	case errNone:
		return "no error"
	case errCommandFail:
		return "command fail"
	case errCommandCrcFail:
		return "command crc fail"
	case errTimeout:
		return "command timeout"
	case errNotReady:
		return "the card is not ready"
	default:
		return "unknown error"
	}
}

const (
	SDIO1 SDIO = 0
	SDIO2 SDIO = 1
)

const (
	_ALT_CKIN = iota + 1
	_ALT_CDIR
	_ALT_D0DIR
	_ALT_D123DIR
	_ALT_D0
	_ALT_D1
	_ALT_D2
	_ALT_D3
	_ALT_D4
	_ALT_D5
	_ALT_D6
	_ALT_D7
	_ALT_CK
	_ALT_CMD
)

const (
	Width1Bit = coresdio.Width1Bit
	Width4Bit = coresdio.Width4Bit
	Width8Bit = coresdio.Width8Bit

	Sdr12  = coresdio.Sdr12
	Sdr25  = coresdio.Sdr25
	Sdr50  = coresdio.Sdr50
	Sdr104 = coresdio.Sdr104
	Ddr50  = coresdio.Ddr50
	Ds     = coresdio.Ds
	Hs     = coresdio.Hs

	//DDR200 = coresdio.DDR200
)

const (
	NoResponse uint8 = iota
	ShortResponse
	ShortResponseNoCrc
	LongResponse
)

const (
	/* Masks for errors in an R5 response */
	r5ComCrcError    = 0x80
	r5IllegalCommand = 0x40
	r5IoCurrentState = 0x30
	r5Error          = 0x08
	r5FunctionNumber = 0x02
	r5OutOfRange     = 0x01
	r5ErrorBits      = r5ComCrcError | r5IllegalCommand | r5Error | r5FunctionNumber | r5OutOfRange
)

var instances [2]_state
var defaultTimeout = time.Millisecond * 500

type (
	SDIO uint8

	_state struct {
		mutex  sync.Mutex
		config Config

		data     []uint32
		index    int
		cmdDone  bool
		dataDone bool
		done     bool
		hasData  bool
		write    bool

		lastResponse [4]uint32
		lastError    error
	}

	Config struct {
		Enable  bool
		CKIN    pin.Pin
		CDIR    pin.Pin
		D0DIR   pin.Pin
		D123DIR pin.Pin
		Dn      [8]pin.Pin
		CK      pin.Pin
		CMD     pin.Pin
		DMA     bool

		IrqCallback func(status sdmmc.RegisterStarType)
	}

	SecondaryConfig struct {
		NegEdge             bool
		BusWidth            BusWidth
		BusSpeed            BusSpeed
		PowerSave           bool
		HardwareFlowControl bool
	}

	_error uint8

	BusWidth     = coresdio.BusWidth
	BusSpeed     = coresdio.BusSpeed
	Command      = coresdio.Command
	CommandIndex = coresdio.CommandType
	Response     = coresdio.Response
	Transfer     = coresdio.Transfer
)

func altFunction(p pin.Pin, alt int, instance SDIO) (corepin.Mode, error) {
	type altKey struct {
		pin      pin.Pin
		function int
		instance SDIO
		mode     corepin.Mode
	}

	var altFunctionTable = [...]altKey{
		// SDIO1
		{pin.PB8, _ALT_CKIN, SDIO1, pin.Alt7},
		{pin.PB8, _ALT_D4, SDIO1, pin.Alt12},
		{pin.PB9, _ALT_CDIR, SDIO1, pin.Alt7},
		{pin.PB9, _ALT_D5, SDIO1, pin.Alt12},
		{pin.PC6, _ALT_D0DIR, SDIO1, pin.Alt8},
		{pin.PC6, _ALT_D6, SDIO1, pin.Alt12},
		{pin.PC7, _ALT_D123DIR, SDIO1, pin.Alt8},
		{pin.PC7, _ALT_D7, SDIO1, pin.Alt12},
		{pin.PC8, _ALT_D0, SDIO1, pin.Alt12},
		{pin.PC9, _ALT_D1, SDIO1, pin.Alt12},
		{pin.PC10, _ALT_D2, SDIO1, pin.Alt12},
		{pin.PC11, _ALT_D3, SDIO1, pin.Alt12},
		{pin.PC12, _ALT_CK, SDIO1, pin.Alt12},
		{pin.PD2, _ALT_CMD, SDIO1, pin.Alt12},

		// SDIO2
		{pin.PA0, _ALT_CMD, SDIO2, pin.Alt9},
		{pin.PB3, _ALT_D2, SDIO2, pin.Alt9},
		{pin.PB4, _ALT_D3, SDIO2, pin.Alt9},
		{pin.PB8, _ALT_D4, SDIO2, pin.Alt10},
		{pin.PB9, _ALT_D5, SDIO2, pin.Alt10},
		{pin.PC14, _ALT_D0, SDIO2, pin.Alt9},
		{pin.PC15, _ALT_D1, SDIO2, pin.Alt9},
		{pin.PC1, _ALT_CK, SDIO2, pin.Alt9},
		{pin.PC6, _ALT_D6, SDIO2, pin.Alt10},
		{pin.PC7, _ALT_D7, SDIO2, pin.Alt10},
		{pin.PD6, _ALT_CK, SDIO2, pin.Alt11},
		{pin.PD7, _ALT_CMD, SDIO2, pin.Alt11},
		{pin.PG11, _ALT_D2, SDIO2, pin.Alt10},
	}

	for _, entry := range altFunctionTable {
		if entry.pin == p && entry.function == alt && entry.instance == instance {
			return entry.mode, nil
		}
	}
	return 0, corehal.ErrInvalidPinout
}

func (s SDIO) Configure(config Config) error {
	regs := sdmmc.Instances[s]
	state := &instances[s]

	state.mutex.Lock()

	switch s {
	case SDIO1:
		stm32h7x7.IrqSdmmc1.Disable()
		rcc.Rcc.Ahb3rstr.SetSdmmc1rst(true)
		for !rcc.Rcc.Ahb3rstr.GetSdmmc1rst() {
		}

		rcc.Rcc.Ahb3rstr.SetSdmmc1rst(false)
		for rcc.Rcc.Ahb3rstr.GetSdmmc1rst() {
		}
	case SDIO2:
		stm32h7x7.IrqSdmmc2.Disable()
		rcc.Rcc.Ahb2rstr.SetSdmmc2rst(true)
		for !rcc.Rcc.Ahb2rstr.GetSdmmc2rst() {
		}

		rcc.Rcc.Ahb2rstr.SetSdmmc2rst(false)
		for rcc.Rcc.Ahb2rstr.GetSdmmc2rst() {
		}
	}

	if !config.Enable {
		state.mutex.Unlock()
		return nil
	}

	// Enable SDIO in RCC.
	switch s {
	case SDIO1:
		rcc.Rcc.Ahb3enr.SetSdmmc1en(true)
		for !rcc.Rcc.Ahb3enr.GetSdmmc1en() {
		}
	case SDIO2:
		rcc.Rcc.Ahb2enr.SetSdmmc2en(true)
		for !rcc.Rcc.Ahb2enr.GetSdmmc2en() {
		}
	}

	// Set up pins.
	pinCfg := func(p pin.Pin, alt int) error {
		if p == 0 {
			return nil
		}
		mode, err := altFunction(p, alt, s)
		if err != nil {
			return err
		}
		p.SetMode(pin.AltFunction | mode)
		return nil
	}

	if err := pinCfg(config.CKIN, _ALT_CKIN); err != nil {
		return err
	}

	if err := pinCfg(config.CDIR, _ALT_CDIR); err != nil {
		return err
	}

	if err := pinCfg(config.D0DIR, _ALT_D0DIR); err != nil {
		return err
	}

	if err := pinCfg(config.D123DIR, _ALT_D123DIR); err != nil {
		return err
	}

	if err := pinCfg(config.CK, _ALT_CK); err != nil {
		return err
	}

	if err := pinCfg(config.CMD, _ALT_CMD); err != nil {
		return err
	}

	for i, Dn := range config.Dn {
		if Dn != 0 {
			Dn.SetMode(pin.Output)
			Dn.High()
		}

		if err := pinCfg(Dn, _ALT_D0+i); err != nil {
			return err
		}

	}

	// Power cycle the card and then turn it off before initializing to SDIO peripheral.
	regs.Power.SetPwrctrl(sdmmc.RegisterPowerFieldPwrctrlEnumPowercycle)
	time.Sleep(time.Millisecond * 50)
	regs.Power.SetPwrctrl(sdmmc.RegisterPowerFieldPwrctrlEnumPoweroff)
	time.Sleep(time.Millisecond * 50)

	// Initialize with default settings.
	regs.Clkcr.SetNegedge(false)
	regs.Clkcr.SetPwrsav(false)
	regs.Clkcr.SetWidbus(0)
	regs.Clkcr.SetHwfcen(false)
	s.setClockFrequency(400_000) // 400KHz init clock.
	regs.Power.SetPwrctrl(sdmmc.RegisterPowerFieldPwrctrlEnumPoweron)

	// Wait 74 cycles.
	time.Sleep(time.Millisecond * 200)

	switch s {
	case SDIO1:
		stm32h7x7.IrqSdmmc1.Enable()
	case SDIO2:
		stm32h7x7.IrqSdmmc2.Enable()
	}

	state.config = config
	state.mutex.Unlock()
	return nil
}

func (s SDIO) Reconfigure(config SecondaryConfig) error {
	state := &instances[s]
	state.mutex.Lock()
	regs := sdmmc.Instances[s]

	var err error
	if config.BusWidth != 0 {
		err := s.setBusWidth(config.BusWidth)
		if err != nil {
			state.mutex.Unlock()
			return err
		}
	}

	if config.BusSpeed != 0 {
		err = s.setBusSpeed(config.BusSpeed)
		if err != nil {
			state.mutex.Unlock()
			return err
		}
	}

	regs.Clkcr.SetNegedge(config.NegEdge)
	regs.Clkcr.SetPwrsav(config.PowerSave)
	regs.Clkcr.SetHwfcen(config.HardwareFlowControl)
	state.mutex.Unlock()
	return nil
}

func (s SDIO) SetBusWidth(width BusWidth) error {
	state := &instances[s]
	state.mutex.Lock()
	err := s.setBusWidth(width)
	state.mutex.Unlock()
	return err
}

func (s SDIO) setBusWidth(width BusWidth) error {
	var cccr uint32
	var buswid uint8

	switch width {
	case Width1Bit:
		buswid = 0
		cccr = 0
	case Width4Bit:
		buswid = 1
		cccr = 2
	case Width8Bit:
		buswid = 2
		cccr = 3
	default:
		return corehal.ErrInvalidConfig
	}

	arg := coresdio.CMD52Args{
		Data:      cccr,
		Address:   0x07,
		Raw:       coresdio.Normal,
		Function:  0,
		ReadWrite: coresdio.Write,
	}

	resp, err := s.sendCommand(Command{
		Class:    CMD52,
		Argument: arg.Value(),
	})

	if err != nil {
		return err
	}

	// Check the response for an error.
	r5 := coresdio.R5(resp[0])
	if r5.Flags()&r5ErrorBits != 0 {
		return coresdio.ErrCommandFail
	}

	// Set the host bus width register.
	regs := sdmmc.Instances[s]
	regs.Clkcr.SetWidbus(buswid)
	return nil
}

func (s SDIO) SetBusSpeed(speed BusSpeed) error {
	state := &instances[s]
	state.mutex.Lock()
	err := s.setBusSpeed(speed)
	state.mutex.Unlock()
	return err
}

func (s SDIO) setBusSpeed(speed BusSpeed) error {
	var speedBits uint8
	var ddr bool
	var busspeed bool

	switch speed {
	case Sdr12:
		speedBits = 0
	case Sdr25:
		speedBits = 1
	case Sdr50:
		speedBits = 2
		busspeed = true
	case Sdr104:
		speedBits = 3
		busspeed = true
	case Ddr50:
		speedBits = 4
		ddr = true
		busspeed = true
	case Hs, Ds:
		speedBits = 0
	default:
		return corehal.ErrInvalidConfig
	}

	arg := coresdio.CMD52Args{
		Address:   0x13,
		Raw:       coresdio.Normal,
		Function:  0,
		ReadWrite: coresdio.Read,
	}

	// CMD52 read to CCCR 0x13.
	resp, err := s.sendCommand(Command{
		Class:    CMD52,
		Argument: arg.Value(),
	})
	if err != nil {
		return err
	}

	r5 := coresdio.R5(resp[0])
	val := r5.Data()
	if speed == Ds {
		// Unset EHS bit.
		val |= 0 << 1
	} else if speed == Hs {
		// Set EHS bit.
		val |= 1 << 1
	} else {
		val &= ^uint8(0x7)    // Clear speed bits.
		val |= speedBits << 0 // Set speed.
		val |= 1 << 1         // EHS.
	}

	arg.Data = uint32(val)
	arg.ReadWrite = coresdio.Write

	// CMD52 write to CCCR 0x13.
	resp, err = s.sendCommand(Command{
		Class:    CMD52,
		Argument: arg.Value(),
	})
	if err != nil {
		return err
	}

	// Check the response for an error.
	if r5.Flags()&r5ErrorBits != 0 {
		return coresdio.ErrCommandFail
	}

	regs := sdmmc.Instances[s]
	regs.Clkcr.SetBusspeed(busspeed)
	regs.Clkcr.SetDdr(ddr)
	return nil
}

func (s SDIO) SetClockFrequency(hz uint32) error {
	state := &instances[s]
	state.mutex.Lock()
	s.setClockFrequency(hz)
	state.mutex.Unlock()
	return nil
}

func (s SDIO) setClockFrequency(hz uint32) {
	if hz == 0 {
		panic("SDMMC clock frequency cannot be zero")
	}
	regs := sdmmc.Instances[s]
	div := uint16(hal.SdmmcSourceFrequencyHz / uint64(2*hz))
	regs.Clkcr.SetClkdiv(div)
}

func (s SDIO) SendCommand(cmd Command) (Response, error) {
	state := &instances[s]
	state.mutex.Lock()
	resp, err := s.sendCommand(cmd)
	state.mutex.Unlock()
	return resp, err
}

func (s SDIO) sendCommand(cmd Command) (Response, error) {
	state := &instances[s]
	regs := sdmmc.Instances[s]
	data := cmd.Data

	// Precompute words length.
	nwords := (len(data) + 3) / 4
	if nwords == 0 {
		nwords = 1
	}
	words := unsafe.Slice((*uint32)(unsafe.Pointer(unsafe.SliceData(data))), nwords)

	transfer := false
	write := false
	blockMode := true
	count := uint32(len(cmd.Data))
	var responseType uint8

	switch cmd.Class {
	case CMD0:
		responseType = NoResponse
	case CMD3, CMD5, CMD7:
		responseType = ShortResponseNoCrc
	default:
		responseType = ShortResponse
	}

	// Determine transfer parameters.
	if cmd.Class&coresdio.CmdFlagBlockRead != 0 {
		transfer = true
	} else if cmd.Class&coresdio.CmdFlagBlockWrite != 0 {
		transfer = true
		write = true
	} else if cmd.Class == CMD53 {
		args := coresdio.DecodeCMD53(cmd.Argument)
		transfer = true
		write = args.ReadWrite == coresdio.Write
		blockMode = args.BlockMode == coresdio.Blocks
		count = args.Count
	}

	var irq sdmmc.RegisterMaskrType
	irq.SetCmdsentie(true)
	irq.SetCmdrendie(true)
	irq.SetCtimeoutie(true)
	irq.SetCcrcfailie(true)

	if transfer {
		irq.SetDataendie(true)
		irq.SetDtimeoutie(true)
		irq.SetDcrcfailie(true)

		if write {
			irq.SetTxunderrie(true)
		} else {
			irq.SetRxoverrie(true)
		}

		// Set up transfer.
		if !state.config.DMA {
			if write {
				irq.SetTxfifoheie(true) // TX FIFO half-empty IRQ
				irq.SetTxfifoeie(true)
			} else {
				irq.SetRxfifohfie(true) // RX FIFO half-full IRQ
				irq.SetRxfifofie(true)
			}
		}

		// Set up the data transfer before sending CMD53.
		var dctrl sdmmc.RegisterDctrlType
		regs.Dtimer.Store(0xFFFF_FFFF)
		if blockMode {
			if cmd.BlockSize == 0 || count == 0 {
				return Response{}, corehal.ErrInvalidConfig
			}
			dctrl.SetDtmode(0) // Block data transfer ending on block count.
			regs.Dlenr.Store(cmd.BlockSize * count)
			dctrl.SetDblocksize(blockSize(int(cmd.BlockSize)))
		} else {
			if count == 512 {
				// NOTE: BlockCount == 0 -> 512 bytes
				count = 0
			} else if count > 512 {
				return Response{}, corehal.ErrInvalidConfig
			}

			regs.Dlenr.Store(count)
			dctrl.SetDblocksize(0)
			dctrl.SetDtmode(1) // SDIO multibyte data transfer.
		}

		dctrl.SetDtdir(!write)
		dctrl.SetSdioen(true)
		dctrl.SetFiforst(true)
		regs.Dctrl.Store(uint32(dctrl))

		if state.config.DMA {
			// Set up IDMA transfer.
			const baseUnit = 32 // bytes.
			regs.Idmabsizer.SetIdmabndt(uint8((len(data) + baseUnit - 1) / baseUnit))
			regs.Idmabase0r.SetIdmabase0(uint32(uintptr(unsafe.Pointer(unsafe.SliceData(data)))))

			var idmaCtrl sdmmc.RegisterIdmactrlrType
			idmaCtrl.SetIdmabact(false)
			idmaCtrl.SetIdmabmode(false)
			idmaCtrl.SetIdmaen(true)
			regs.Idmactrlr.Store(uint32(idmaCtrl))
		}
	}

	// Clear pending flags.
	regs.Icr.StoreBits(0xFFFF_FFFF)

	// Wait for CPSM idle.
	for regs.Star.GetCpsmact() {
	}

	// Reset state.
	state.done = false
	state.cmdDone = false
	state.dataDone = !transfer
	state.data = words
	state.hasData = transfer
	state.index = 0
	state.lastError = nil
	state.lastResponse = [4]uint32{}
	state.write = write

	// Enable interrupts.
	regs.Maskr.Store(uint32(irq))

	// Issue command.
	regs.Argr.SetCmdarg(cmd.Argument)

	var cmdr sdmmc.RegisterCmdrType
	cmdr.SetCmdindex(uint8(cmd.Class & 0xFF))
	cmdr.SetWaitresp(responseType)
	cmdr.SetWaitint(false)
	cmdr.SetCpsmen(true)
	cmdr.SetCmdtrans(transfer)
	regs.Cmdr.Store(uint32(cmdr))

	// Busy-wait for ISR to complete command.
	deadline := time.Now().Add(defaultTimeout)
	for !state.done {
		if time.Now().After(deadline) {
			// Wait for the command path to become inactive.
			for regs.Star.GetCpsmact() {
				runtime.Gosched()
			}

			// Check if the data path is active.
			for regs.Star.GetDpsmact() {
				// Send stop command.
				regs.Cmdr.Store(0x1080)

				// Wait for the command path to become inactive.
				for regs.Star.GetCpsmact() {
					runtime.Gosched()
				}
			}

			return Response{}, coresdio.ErrTimeout
		}
		runtime.Gosched()
	}

	return state.lastResponse, state.lastError
}

//sigo:interrupt sdmmc1Handler Sdmmc1Handler
func sdmmc1Handler() {
	irqHandler(SDIO1)
}

//sigo:interrupt sdmmc2Handler Sdmmc2Handler
func sdmmc2Handler() {
	irqHandler(SDIO2)
}

func irqHandler(instance SDIO) {
	state := &instances[instance]
	regs := sdmmc.Instances[instance]
	status := &regs.Star
	icr := &regs.Icr
	fifof := &regs.Fifor

	length := len(state.data)
	data := state.data
	index := state.index
	write := state.write
	hasData := state.hasData

	if state.config.IrqCallback != nil {
		defer state.config.IrqCallback(*status)
	}

	if hasData && index < length {
		if write {
			for !status.GetTxfifof() && index < length {
				volatile.StoreUint32((*uint32)(fifof), data[index])
				index++
			}
		} else {
			for !status.GetRxfifoe() && index < length {
				state.data[index] = volatile.LoadUint32((*uint32)(fifof))
				index++
			}
		}
		state.index = index
	}

	if status.GetSdioit() {
		// Just clear the SDIO interrupt status.
		icr.SetSdioitc(true)
	}

	if status.GetDataend() {
		icr.SetDataendc(true)

		// Disable data-related interrupts.
		regs.Maskr.ClearBits(0x0000_003F)

		// All data has finished transmission.
		state.hasData = false
		state.dataDone = true
	}

	// If any error
	if status.GetDcrcfail() || status.GetCtimeout() || status.GetDtimeout() || status.GetRxoverr() || status.GetTxunderr() {
		if status.GetDcrcfail() {
			icr.SetDcrcfailc(true)
			state.lastError = coresdio.ErrCommandCrcFail
		}

		if status.GetCtimeout() {
			icr.SetCtimeoutc(true)
			state.lastError = coresdio.ErrTimeout
		}

		if status.GetDtimeout() {
			icr.SetDtimeoutc(true)
			state.lastError = coresdio.ErrTimeout
		}

		if status.GetRxoverr() {
			icr.SetRxoverrc(true)
			state.lastError = coresdio.ErrDataError
		}

		if status.GetTxunderr() {
			icr.SetTxunderrc(true)
			state.lastError = coresdio.ErrDataError
		}

		// Clear all flags.
		icr.StoreBits(0xFFFF_FFFF)

		// Disable interrupts
		regs.Maskr.Store(0)

		state.done = true
		return
	}

	// If command complete
	if status.GetCmdrend() {
		icr.SetCmdrendc(true)
		state.lastResponse[0] = regs.Resp1r.Load()
		state.lastResponse[1] = regs.Resp2r.Load()
		state.lastResponse[2] = regs.Resp3r.Load()
		state.lastResponse[3] = regs.Resp4r.Load()
		state.cmdDone = true
	}

	// If command sent without response
	if status.GetCmdsent() {
		icr.SetCmdsentc(true)
		state.cmdDone = true
	}

	// If IDMA boundary (double buffer complete)
	if status.GetIdmabtc() {
		icr.SetIdmabtcc(true)
		// TODO: Track buffer switching here.
	}

	if state.cmdDone && state.dataDone {
		state.done = true

		// Disable interrupts.
		regs.Maskr.Store(0)
	}
}

func blockSize(size int) uint8 {
	switch {
	case size >= 16384:
		return 0b1110
	case size >= 8192:
		return 0b1101
	case size >= 4096:
		return 0b1100
	case size >= 2048:
		return 0b1011
	case size >= 1024:
		return 0b1010
	case size >= 512:
		return 0b1001
	case size >= 256:
		return 0b1000
	case size >= 128:
		return 0b0111
	case size >= 64:
		return 0b0110
	case size >= 32:
		return 0b0101
	case size >= 16:
		return 0b0100
	case size >= 8:
		return 0b0011
	case size >= 4:
		return 0b0010
	case size >= 2:
		return 0b0001
	case size >= 1:
		return 0b0000
	default:
		return 0b0010
	}
}
