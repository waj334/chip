//go:build stm32h7x7

package sdio

import (
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/sdmmc"
	"runtime"
	"sync"
	"time"
	"unsafe"

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

	CMD0  = coresdio.CMD0
	CMD2  = coresdio.CMD2
	CMD3  = coresdio.CMD3
	CMD5  = coresdio.CMD5
	CMD7  = coresdio.CMD7
	CMD8  = coresdio.CMD8
	CMD9  = coresdio.CMD9
	CMD12 = coresdio.CMD12
	CMD17 = coresdio.CMD17
	CMD24 = coresdio.CMD24
	CMD41 = coresdio.CMD41
	CMD52 = coresdio.CMD52
	CMD53 = coresdio.CMD53
	CMD55 = coresdio.CMD55
)

const (
	NoResponse uint8 = iota
	ShortResponse
	ShortResponseNoCrc
	LongResponse
)

var instances = [2]_state{}
var defaultTimeout = time.Millisecond * 500

type (
	SDIO uint8

	_state struct {
		mutex sync.Mutex
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
	CommandIndex = coresdio.CommandIndex
	Response     = coresdio.Response
	Transfer     = coresdio.Transfer
)

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

func altFunction(p pin.Pin, alt int, instance SDIO) (corepin.Mode, error) {
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
		//stm32h7x7.IrqSdmmc1.Disable()
		rcc.Rcc.Ahb3rstr.SetSdmmc1rst(true)
		for !rcc.Rcc.Ahb3rstr.GetSdmmc1rst() {
		}

		rcc.Rcc.Ahb3rstr.SetSdmmc1rst(false)
		for rcc.Rcc.Ahb3rstr.GetSdmmc1rst() {
		}
	case SDIO2:
		//stm32h7x7.IrqSdmmc2.Disable()
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

	/*
		switch s {
		case SDIO1:
			stm32h7x7.IrqSdmmc1.Enable()
		case SDIO2:
			stm32h7x7.IrqSdmmc2.Enable()
		}
	*/

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

	resp, err := s.sendCommand(Command{
		Index: CMD52,
		Argument: (1 << 31) | // Write flag.
			(0 << 28) | // Function 0
			(0 << 27) | // No RAW
			(0x07 << 9) | // Address = 0x07
			cccr, // Data.
	})
	if err != nil {
		return err
	}

	// Check the response for an error.
	flags := uint8(resp[0] >> 8)
	if flags&0b1100_1011 != 0 {
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

	// CMD52 read to CCCR 0x13.
	resp, err := s.sendCommand(Command{
		Index: CMD52,
		Argument: (0 << 31) | // Read.
			(0 << 28) | // Func 0.
			(0 << 27) | // No RAW.
			(0x13 << 9), // Addr.
	})
	if err != nil {
		return err
	}

	// Data is in bits 7:0.
	val := uint8(resp[0])
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

	// CMD52 write to CCCR 0x13.
	resp, err = s.sendCommand(Command{
		Index: CMD52,
		Argument: (1 << 31) | // Write.
			(0 << 28) | // Func 0.
			(0 << 27) | // No RAW.
			(0x13 << 9) | // Addr.
			uint32(val),
	})
	if err != nil {
		return err
	}

	// Check the response for an error.
	flags := uint8(resp[0] >> 8)
	if flags&0b1100_1011 != 0 {
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
	regs := sdmmc.Instances[s]

	interrupt := false
	transfer := false
	var responseType uint8

	switch cmd.Index {
	case CMD0:
		responseType = NoResponse
	case CMD3, CMD5, CMD7:
		responseType = ShortResponseNoCrc
	case CMD53:
		transfer = true
		responseType = ShortResponseNoCrc
	default:
		responseType = ShortResponse
	}

	// Clear pending flags.
	regs.Icr.StoreBits(0xFFFF_FFFF)

	// Wait for CPSM idle.
	for regs.Star.GetCpsmact() {
	}

	// Issue command.
	regs.Argr.SetCmdarg(cmd.Argument)

	var cmdr sdmmc.RegisterCmdrType
	cmdr.SetCmdindex(uint8(cmd.Index))
	cmdr.SetWaitresp(responseType)
	cmdr.SetWaitint(interrupt)
	cmdr.SetCpsmen(true)
	cmdr.SetCmdtrans(transfer)
	regs.Cmdr.Store(uint32(cmdr))

	// Busy-wait for ISR to complete command.
	deadline := time.Now().Add(defaultTimeout)
	status := &regs.Star
	for {
		switch {
		case status.GetCmdrend():
			regs.Icr.SetCmdrendc(true)
			goto response
		case status.GetCmdsent():
			regs.Icr.SetCmdsentc(true)
			return Response{}, nil
		case status.GetCtimeout():
			regs.Icr.SetCtimeoutc(true)
			return Response{}, coresdio.ErrTimeout
		case status.GetCcrcfail():
			regs.Icr.SetCcrcfailc(true)
			return Response{}, coresdio.ErrCommandCrcFail
		case time.Now().After(deadline):
			return Response{}, coresdio.ErrTimeout
		}

		// Do something else while waiting for the response.
		runtime.Gosched()
	}

response:
	var resp Response
	resp[0] = regs.Resp1r.GetCardstatus1()
	resp[1] = regs.Resp2r.GetCardstatus2()
	resp[2] = regs.Resp3r.GetCardstatus3()
	resp[3] = regs.Resp4r.GetCardstatus4()

	return resp, nil
}

func (s SDIO) ReadBytes(buf []byte, transfer Transfer) error {
	state := &instances[s]
	state.mutex.Lock()
	err := s.transfer(buf, transfer, false, false)
	state.mutex.Unlock()
	return err
}

func (s SDIO) WriteBytes(buf []byte, transfer Transfer) error {
	state := &instances[s]
	state.mutex.Lock()
	err := s.transfer(buf, transfer, true, false)
	state.mutex.Unlock()
	return err
}

func (s SDIO) ReadBlocks(buf []byte, transfer Transfer) error {
	state := &instances[s]
	state.mutex.Lock()
	err := s.transferBlocks(buf, false, transfer)
	state.mutex.Unlock()
	return err
}

func (s SDIO) WriteBlocks(buf []byte, transfer Transfer) error {
	state := &instances[s]
	state.mutex.Lock()
	err := s.transferBlocks(buf, true, transfer)
	state.mutex.Unlock()
	return err
}

func (s SDIO) transferBlocks(buf []byte, write bool, transfer Transfer) error {
	var err error

	if transfer.BlockSize == 0 {
		return corehal.ErrInvalidParameter
	}

	if transfer.BlockCount == 0 {
		transfer.BlockCount = uint16(len(buf)) / transfer.BlockSize
	}

	remaining := len(buf) % int(transfer.BlockSize)
	fullBlockBytes := len(buf) - remaining
	if fullBlockBytes > 0 {
		err = s.transfer(buf[:fullBlockBytes], transfer, write, true)
		if err != nil {
			return err
		}
	}

	if remaining > 0 {
		// Allocate a block to store the remaining bytes.
		b := make([]byte, transfer.BlockSize)
		copy(b, buf[fullBlockBytes:])

		// Send the remaining block.
		t := transfer
		t.BlockCount = 1
		if t.Increment {
			t.Address += uint32(fullBlockBytes)
		}
		err = s.transfer(b, t, write, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s SDIO) transfer(buf []byte, transfer Transfer, write bool, blockMode bool) error {
	if len(buf) == 0 {
		return corehal.ErrInvalidConfig
	}

	regs := sdmmc.Instances[s]
	arg := uint32(0)

	// Enable interrupts required from IDMA.
	regs.Maskr.SetDcrcfailie(true)
	regs.Maskr.SetDtimeoutie(true)
	regs.Maskr.SetRxoverrie(true)
	regs.Maskr.SetTxunderrie(true)
	regs.Maskr.SetDataendie(true)

	// Set up the data transfer before sending CMD53.
	var dctrl sdmmc.RegisterDctrlType
	regs.Dtimer.Store(0xFFFF_FFFF)
	if blockMode {
		if transfer.BlockSize == 0 || transfer.BlockCount == 0 {
			return corehal.ErrInvalidConfig
		}
		dctrl.SetDtmode(0) // Block data transfer ending on block count.
		regs.Dlenr.Store(uint32(transfer.BlockSize * transfer.BlockCount))
		dctrl.SetDblocksize(blockSize(int(transfer.BlockSize)))
		arg |= uint32(transfer.BlockCount & 0x1FF)
		arg |= 1 << 27
	} else {
		if len(buf) > 511 {
			return corehal.ErrInvalidConfig
		}
		dctrl.SetDtmode(1) // SDIO multibyte data transfer.
		regs.Dlenr.Store(uint32(len(buf)))
		arg |= uint32(len(buf) & 0x1FF)
	}

	dctrl.SetDtdir(!write)
	//dctrl.SetDten(true)
	regs.Dctrl.Store(uint32(dctrl))

	// Set up IDMA transfer.
	const baseUnit = 32 // bytes.
	regs.Idmabsizer.SetIdmabndt(uint8((len(buf) + baseUnit - 1) / baseUnit))
	regs.Idmabase0r.SetIdmabase0(uint32(uintptr(unsafe.Pointer(unsafe.SliceData(buf)))))

	var idmaCtrl sdmmc.RegisterIdmactrlrType
	idmaCtrl.SetIdmabact(false)
	idmaCtrl.SetIdmabmode(false)
	idmaCtrl.SetIdmaen(true)
	regs.Idmactrlr.Store(uint32(idmaCtrl))

	if write {
		arg |= 1 << 31
	}

	if transfer.Increment {
		arg |= 1 << 26
	}

	arg |= uint32(transfer.Function&0x3) << 28
	arg |= (transfer.Address & 0x1FFFF) << 9

	// Send CMD53 to begin the byte transfer.
	resp, err := s.sendCommand(Command{
		Index:    CMD53,
		Argument: arg, // Packed CMD53 argument: function, RW, mode, addr, count
	})
	if err != nil {
		return err
	}

	// Check the response for an error.
	flags := uint8(resp[0] >> 8)
	const cmd53ErrorMask = 0b11001011
	if flags&cmd53ErrorMask != 0 {
		return coresdio.ErrCommandFail
	}

	//deadline := time.Now().Add(defaultTimeout)
	status := &regs.Star
	for {
		if status.GetDataend() {
			break
		}
		if status.GetDtimeout() { // || time.Now().After(deadline) {
			return coresdio.ErrTimeout
		}
		if status.GetDcrcfail() {
			return coresdio.ErrCommandCrcFail
		}
		if !write && status.GetRxoverr() {
			return coresdio.ErrDataError
		}
		if write && status.GetTxunderr() {
			return coresdio.ErrDataError
		}

		// Do something else while waiting for the transfer to complete.
		runtime.Gosched()
	}
	return nil
}

/*
//sigo:interrupt sdmmc1Handler Sdmmc1Handler
func sdmmc1Handler() {
	irqHandler(SDIO1)
}

//sigo:interrupt sdmmc2Handler Sdmmc2Handler
func sdmmc2Handler() {
	irqHandler(SDIO2)
}

func irqHandler(instance SDIO) {
	// Not really needed right now. Perhaps later for wake up logic.
}
*/

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
