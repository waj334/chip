//go:build stm32h7x7

package sdio

import (
	"runtime"
	"sync"
	"time"
	"unsafe"
	"volatile"

	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/sdmmc"

	stm32h7x7 "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
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

// SDMMC_MASKR / SDMMC_STAR bit positions per RM0399 §58.10.13 / §58.10.11.
// These constants define the bit index for each interrupt enable / status flag.
const (
	bitCCRCFAIL = 0
	bitDCRCFAIL = 1
	bitCTIMEOUT = 2
	bitDTIMEOUT = 3
	bitTXUNDERR = 4
	bitRXOVERR  = 5
	bitCMDREND  = 6
	bitCMDSENT  = 7
	bitDATAEND  = 8
	bitTXFIFOHE = 14
	bitRXFIFOHF = 15
	bitRXFIFOF  = 17
	bitIDMABTC  = 28
)

var instances [2]_state
var defaultTimeout = time.Millisecond * 500

type (
	SDIO uint8

	_state struct {
		mutex  sync.Mutex
		config Config

		data []byte

		index       int32 // bytes processed
		done        int32 // 0 = not done, 1 = done
		cmdDone     int32 // 0 = command phase not complete, 1 = complete
		dataDone    int32 // 0 = data phase not complete, 1 = complete
		useDMA      int32 // 0 = PIO/IRQ FIFO path, 1 = IDMA path
		wordAligned int32 // 0 = byte path, 1 = aligned fast path
		write       int32 // 0 = read, 1 = write

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
		state.mutex.Unlock()
		return err
	}
	if err := pinCfg(config.CDIR, _ALT_CDIR); err != nil {
		state.mutex.Unlock()
		return err
	}
	if err := pinCfg(config.D0DIR, _ALT_D0DIR); err != nil {
		state.mutex.Unlock()
		return err
	}
	if err := pinCfg(config.D123DIR, _ALT_D123DIR); err != nil {
		state.mutex.Unlock()
		return err
	}
	if err := pinCfg(config.CK, _ALT_CK); err != nil {
		state.mutex.Unlock()
		return err
	}
	if err := pinCfg(config.CMD, _ALT_CMD); err != nil {
		state.mutex.Unlock()
		return err
	}

	for i, Dn := range config.Dn {
		if err := pinCfg(Dn, _ALT_D0+i); err != nil {
			state.mutex.Unlock()
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
	s.setClockFrequency(400_000)
	regs.Power.SetPwrctrl(sdmmc.RegisterPowerFieldPwrctrlEnumPoweron)

	// 74 SD clock cycles at 400 KHz is ~185 us. 1 ms is already generous.
	time.Sleep(time.Millisecond)

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

	if config.BusWidth != 0 {
		err := s.setBusWidth(config.BusWidth)
		if err != nil {
			state.mutex.Unlock()
			return err
		}
	}

	if config.BusSpeed != 0 {
		err := s.setBusSpeed(config.BusSpeed)
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

	resp, err := s.sendCommand(Command{
		Class:    CMD52,
		Argument: arg.Value(),
	})
	if err != nil {
		return err
	}

	r5 := coresdio.R5(resp[0])
	if r5.Flags()&r5ErrorBits != 0 {
		return coresdio.ErrCommandFail
	}

	val := r5.Data()
	if speed == Ds {
		val &= ^uint8(1 << 1)
	} else if speed == Hs {
		val |= 1 << 1
	} else {
		val &= ^uint8(0x7)
		val |= speedBits
		val |= 1 << 1
	}

	arg.Data = uint32(val)
	arg.ReadWrite = coresdio.Write

	resp, err = s.sendCommand(Command{
		Class:    CMD52,
		Argument: arg.Value(),
	})
	if err != nil {
		return err
	}

	r5Write := coresdio.R5(resp[0])
	if r5Write.Flags()&r5ErrorBits != 0 {
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

	cmdIndex := uint8(cmd.Class & 0xFF)

	var responseType uint8
	switch cmdIndex {
	case 0:
		responseType = NoResponse
	case 2, 9, 10:
		responseType = LongResponse
	case 3, 5, 7:
		responseType = ShortResponseNoCrc
	default:
		responseType = ShortResponse
	}

	transfer := cmdIndex == uint8(CMD53&0xFF)
	write := false
	blockMode := false
	var xferBytes uint32

	if transfer {
		write = (cmd.Argument & (1 << 31)) != 0
		blockMode = (cmd.Argument & (1 << 27)) != 0
		count := cmd.Argument & 0x1FF

		if blockMode {
			if cmd.BlockSize == 0 || count == 0 {
				return Response{}, corehal.ErrInvalidConfig
			}
			if !isPowerOfTwo(int(cmd.BlockSize)) {
				return Response{}, corehal.ErrInvalidConfig
			}
			xferBytes = cmd.BlockSize * count
		} else {
			if count == 0 {
				count = 512
			}
			if count > 512 {
				return Response{}, corehal.ErrInvalidConfig
			}
			xferBytes = count
		}

		if uint32(len(data)) != xferBytes {
			return Response{}, corehal.ErrInvalidConfig
		}
	}

	wordAligned := isWordAligned(data)
	useDMA := transfer &&
		state.config.DMA &&
		IsSDMMCDMABuffer(data) &&
		wordAligned

	// If the caller configured DMA but the buffer isn't DMA-accessible,
	// return an error rather than silently falling back to PIO.
	if transfer && state.config.DMA && !useDMA {
		return Response{}, corehal.ErrInvalidBuffer
	}

	// Make sure prior command/data activity is gone before reprogramming.
	regs.Maskr.Store(0)
	regs.Idmactrlr.Store(0)
	regs.Icr.StoreBits(0xFFFF_FFFF)

	for regs.Star.GetCpsmact() {
	}
	for regs.Star.GetDpsmact() {
	}

	// Publish shared state before enabling interrupts / issuing the command.
	state.data = data
	volatile.StoreInt32(&state.index, 0)
	volatile.StoreInt32(&state.done, 0)
	volatile.StoreInt32(&state.cmdDone, 0)
	volatile.StoreInt32(&state.dataDone, boolToInt32(!transfer))
	volatile.StoreInt32(&state.useDMA, boolToInt32(useDMA))
	volatile.StoreInt32(&state.wordAligned, boolToInt32(wordAligned))
	volatile.StoreInt32(&state.write, boolToInt32(write))
	state.lastError = nil
	state.lastResponse = [4]uint32{}

	if transfer {
		var dctrl sdmmc.RegisterDctrlType

		regs.Dtimer.Store(0xFFFF_FFFF)

		if blockMode {
			regs.Dlenr.Store(xferBytes)
			dctrl.SetDtmode(0)
			dctrl.SetDblocksize(blockSize(int(cmd.BlockSize)))
		} else {
			regs.Dlenr.Store(xferBytes)
			dctrl.SetDtmode(1) // SDIO multibyte mode
			dctrl.SetDblocksize(0)
		}

		dctrl.SetDtdir(!write)

		// Do NOT set DTEN for SD/SDIO/eMMC card transfers.
		// CMDTRANS is what should start the DPSM for CMD53.
		regs.Dctrl.Store(uint32(dctrl))

		if useDMA {
			const baseUnit = 32 // bytes
			regs.Idmabsizer.SetIdmabndt(uint8((len(data) + baseUnit - 1) / baseUnit))
			regs.Idmabase0r.SetIdmabase0(uint32(uintptr(dataPointer(data))))

			var idmaCtrl sdmmc.RegisterIdmactrlrType
			idmaCtrl.SetIdmabact(false)
			idmaCtrl.SetIdmabmode(false)
			idmaCtrl.SetIdmaen(true)
			regs.Idmactrlr.Store(uint32(idmaCtrl))
		}
	}

	// Clear flags again after programming transfer state, before enabling IRQs.
	regs.Icr.StoreBits(0xFFFF_FFFF)

	// Enable only the interrupts you actually use.
	// Bit positions per RM0399 SDMMC_MASKR (§58.10.13):
	//   0: CCRCFAILIE    1: DCRCFAILIE    2: CTIMEOUTIE   3: DTIMEOUTIE
	//   4: TXUNDERRIE    5: RXOVERRIE     6: CMDRENDIE    7: CMDSENTIE
	//   8: DATAENDIE    14: TXFIFOHEIE   15: RXFIFOHFIE  17: RXFIFOFIE
	//  28: IDMABTCIE
	var mask uint32
	mask |= 1 << bitCCRCFAIL // CCRCFAILIE
	mask |= 1 << bitCTIMEOUT // CTIMEOUTIE

	if responseType == NoResponse {
		mask |= 1 << bitCMDSENT // CMDSENTIE
	} else {
		mask |= 1 << bitCMDREND // CMDRENDIE
	}

	if transfer {
		mask |= 1 << bitTXUNDERR // TXUNDERRIE
		mask |= 1 << bitRXOVERR  // RXOVERRIE
		mask |= 1 << bitDATAEND  // DATAENDIE
		mask |= 1 << bitDCRCFAIL // DCRCFAILIE
		mask |= 1 << bitDTIMEOUT // DTIMEOUTIE

		if useDMA {
			mask |= 1 << bitIDMABTC // IDMABTCIE
		} else if write {
			mask |= 1 << bitTXFIFOHE // TXFIFOHEIE
		} else {
			mask |= 1 << bitRXFIFOHF // RXFIFOHFIE
			mask |= 1 << bitRXFIFOF  // RXFIFOFIE
		}
	}

	regs.Maskr.Store(mask)

	// Issue command.
	regs.Argr.SetCmdarg(cmd.Argument)

	var cmdr sdmmc.RegisterCmdrType
	cmdr.SetCmdindex(cmdIndex)
	cmdr.SetWaitresp(responseType)
	cmdr.SetWaitint(false)
	cmdr.SetCpsmen(true)
	cmdr.SetCmdtrans(transfer) // This is the correct way to start CMD53 data path.
	regs.Cmdr.Store(uint32(cmdr))

	// Wait for ISR completion.
	deadline := time.Now().Add(defaultTimeout)

	for volatile.LoadInt32(&state.done) == 0 {
		// For non-DMA PIO reads, poll-drain the FIFO from the goroutine
		// context. The ISR drains the FIFO when RXFIFOHF fires, but
		// RXFIFOHF requires ≥8 words (32 bytes) in the FIFO. Transfers
		// smaller than 32 bytes — or the final chunk of any transfer —
		// never trigger RXFIFOHF. Without this drain the DPSM stays in
		// Wait_R (it won't move to Idle until the FIFO is empty), DATAEND
		// never fires, and sendCommand times out.
		//
		// The SDMMC IRQ is briefly disabled so we don't race with the ISR
		// over the FIFO and the shared index.
		if transfer && !useDMA && !write {
			s.disableIRQ()
			idx := int(volatile.LoadInt32(&state.index))
			if idx < len(data) && !regs.Star.GetRxfifoe() {
				fifop := (*uint32)(&regs.Fifor)
				for idx < len(data) && !regs.Star.GetRxfifoe() {
					w := volatile.LoadUint32(fifop)
					if wordAligned && idx+4 <= len(data) {
						*(*uint32)(unsafe.Add(dataPointer(data), idx)) = w
						idx += 4
					} else {
						unpackWordLE(data, idx, w)
						rem := len(data) - idx
						if rem > 4 {
							rem = 4
						}
						idx += rem
					}
				}
				volatile.StoreInt32(&state.index, int32(idx))
			}

			// The FIFO may now be empty — check whether DATAEND fired.
			if regs.Star.GetDataend() && volatile.LoadInt32(&state.dataDone) == 0 {
				regs.Icr.SetDataendc(true)
				volatile.StoreInt32(&state.dataDone, 1)
				if volatile.LoadInt32(&state.cmdDone) != 0 {
					regs.Maskr.Store(0)
					regs.Idmactrlr.Store(0)
					volatile.StoreInt32(&state.done, 1)
				}
			}
			s.enableIRQ()
		}

		if time.Now().After(deadline) {
			regs.Maskr.Store(0)
			regs.Idmactrlr.Store(0)

			for regs.Star.GetCpsmact() {
				runtime.Gosched()
			}

			for regs.Star.GetDpsmact() {
				regs.Cmdr.Store(0x1080)
				for regs.Star.GetCpsmact() {
					runtime.Gosched()
				}
			}

			regs.Icr.StoreBits(0xFFFF_FFFF)
			volatile.StoreInt32(&state.done, 1)
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

func (s SDIO) disableIRQ() {
	switch s {
	case SDIO1:
		stm32h7x7.IrqSdmmc1.Disable()
	case SDIO2:
		stm32h7x7.IrqSdmmc2.Disable()
	}
}

func (s SDIO) enableIRQ() {
	switch s {
	case SDIO1:
		stm32h7x7.IrqSdmmc1.Enable()
	case SDIO2:
		stm32h7x7.IrqSdmmc2.Enable()
	}
}

func irqHandler(instance SDIO) {
	state := &instances[instance]
	regs := sdmmc.Instances[instance]
	status := &regs.Star
	icr := &regs.Icr
	fifof := &regs.Fifor

	if volatile.LoadInt32(&state.done) != 0 {
		icr.StoreBits(0xFFFF_FFFF)
		regs.Maskr.Store(0)
		return
	}

	data := state.data
	length := len(data)
	index := int(volatile.LoadInt32(&state.index))
	write := volatile.LoadInt32(&state.write) != 0
	useDMA := volatile.LoadInt32(&state.useDMA) != 0
	wordAligned := volatile.LoadInt32(&state.wordAligned) != 0
	dataPending := volatile.LoadInt32(&state.dataDone) == 0

	// PIO FIFO servicing.
	if dataPending && !useDMA && regs.Star.GetDpsmact() && index < length {
		if write {
			if wordAligned {
				for index+4 <= length {
					if status.GetTxfifof() {
						break
					}

					remainingWords := (length - index) >> 2
					if remainingWords > 8 {
						remainingWords = 8
					}

					base := unsafe.Add(dataPointer(data), index)
					for i := 0; i < remainingWords && !status.GetTxfifof(); i++ {
						volatile.StoreUint32((*uint32)(fifof), *(*uint32)(unsafe.Add(base, i*4)))
						index += 4
					}
				}
			}

			if index < length && (!wordAligned || (length-index) < 4) {
				for index < length && !status.GetTxfifof() {
					volatile.StoreUint32((*uint32)(fifof), packWordLE(data, index))

					remaining := length - index
					if remaining > 4 {
						remaining = 4
					}
					index += remaining
				}
			}
		} else {
			if wordAligned {
				for index+4 <= length {
					if status.GetRxfifoe() {
						break
					}

					remainingWords := (length - index) >> 2
					if remainingWords > 8 {
						remainingWords = 8
					}

					base := unsafe.Add(dataPointer(data), index)
					for i := 0; i < remainingWords && !status.GetRxfifoe(); i++ {
						*(*uint32)(unsafe.Add(base, i*4)) = volatile.LoadUint32((*uint32)(fifof))
						index += 4
					}
				}
			}

			if index < length && (!wordAligned || (length-index) < 4) {
				for index < length && !status.GetRxfifoe() {
					v := volatile.LoadUint32((*uint32)(fifof))
					unpackWordLE(data, index, v)

					remaining := length - index
					if remaining > 4 {
						remaining = 4
					}
					index += remaining
				}
			}
		}

		volatile.StoreInt32(&state.index, int32(index))
	}

	if status.GetSdioit() {
		icr.SetSdioitc(true)
	}

	if status.GetIdmabtc() {
		icr.SetIdmabtcc(true)
	}

	if status.GetDataend() {
		icr.SetDataendc(true)
		volatile.StoreInt32(&state.dataDone, 1)
	}

	// Errors.
	if status.GetCcrcfail() ||
		status.GetDcrcfail() ||
		status.GetCtimeout() ||
		status.GetDtimeout() ||
		status.GetRxoverr() ||
		status.GetTxunderr() {

		if status.GetCcrcfail() {
			icr.SetCcrcfailc(true)
			state.lastError = coresdio.ErrCommandCrcFail
		}

		if status.GetDcrcfail() {
			icr.SetDcrcfailc(true)
			if state.lastError == nil {
				state.lastError = coresdio.ErrDataError
			}
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

		regs.Maskr.Store(0)

		if !useDMA && regs.Star.GetDpsmact() {
			var dctrl sdmmc.RegisterDctrlType
			dctrl.SetFiforst(true)
			regs.Dctrl.Store(uint32(dctrl))
		}

		volatile.StoreInt32(&state.done, 1)

		if state.config.IrqCallback != nil {
			state.config.IrqCallback(*status)
		}
		return
	}

	if status.GetCmdrend() {
		icr.SetCmdrendc(true)
		state.lastResponse[0] = regs.Resp1r.Load()
		state.lastResponse[1] = regs.Resp2r.Load()
		state.lastResponse[2] = regs.Resp3r.Load()
		state.lastResponse[3] = regs.Resp4r.Load()
		volatile.StoreInt32(&state.cmdDone, 1)
	}

	if status.GetCmdsent() {
		icr.SetCmdsentc(true)
		volatile.StoreInt32(&state.cmdDone, 1)
	}

	if volatile.LoadInt32(&state.cmdDone) != 0 && volatile.LoadInt32(&state.dataDone) != 0 {
		regs.Maskr.Store(0)
		regs.Idmactrlr.Store(0)
		volatile.StoreInt32(&state.done, 1)
	}

	if state.config.IrqCallback != nil {
		state.config.IrqCallback(*status)
	}
}
