package sdram

import (
	"pkg.si-go.dev/chip/arm/cortexm"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/fmc"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
)

const (
	SDRAM0 = 0
	SDRAM1 = 1
)

const (
	// FMC SDCMR command modes.
	cmdNormal       = 0b000
	cmdClockEnable  = 0b001
	cmdPrechargeAll = 0b010
	cmdAutoRefresh  = 0b011
	cmdLoadModeReg  = 0b100

	// FMC_SDCMR: CTB1 = bit 4 (Bank 1, 0xC0000000), CTB2 = bit 3 (Bank 2, 0xD0000000).
	sdcmdTargetBank1 = uint32(1 << 4) // CTB1
	sdcmdTargetBank2 = uint32(1 << 3) // CTB2
)

type BusWidth uint8

const (
	BusWidth8  BusWidth = 0
	BusWidth16 BusWidth = 1
	BusWidth32 BusWidth = 2
)

type RowBits uint8

const (
	RowBits11Bit RowBits = 0
	RowBits12Bit RowBits = 1
	RowBits13Bit RowBits = 2
)

type ColumnBits uint8

const (
	ColumnBits8Bit  ColumnBits = 0
	ColumnBits9Bit  ColumnBits = 1
	ColumnBits10Bit ColumnBits = 2
	ColumnBits11Bit ColumnBits = 3
)

type BankCount bool

const (
	BankCount2 = false
	BankCount4 = true
)

type ReadPipe uint8

const (
	NoDelay     ReadPipe = 0
	Delay1Cycle ReadPipe = 1
	Delay2Cycle ReadPipe = 2
)

type CASLatency uint8

const (
	CAS1 CASLatency = 1
	CAS2 CASLatency = 2
	CAS3 CASLatency = 3
)

type ClockPeriod uint8

const (
	Period0Cycles ClockPeriod = 0
	Period2Cycles ClockPeriod = 2
	Period3Cycles ClockPeriod = 3
)

type SDRAM uint8

type TimingConfig struct {
	TRCD uint8 // Row to column delay
	TRP  uint8 // Row precharge delay
	TWR  uint8 // Recovery Delay
	TRC  uint8 // Row cycle delay
	TRAS uint8 // Self-refresh time
	TXSR uint8 // Exit self-refresh delay
	TMRD uint8 // Load Mode Register to Active
	_    uint8
}

type BankConfig struct {
	Enable bool

	BusWidth      BusWidth
	RowBits       RowBits
	ColumnBits    ColumnBits
	InternalBanks BankCount

	BurstRead       bool
	WriteProtection bool

	ReadPipeDelay ReadPipe
	CASLatency    CASLatency
	ClockPeriod   ClockPeriod

	LoadMode uint32

	Timing TimingConfig
}

type Config struct {
	Bank [2]BankConfig

	CLK  pin.Pin     // SDRAM clock
	CKE  pin.Pin     // SDRAM Bank Clock Enable
	CS   pin.Pin     // SDRAM Bank Chip Enable
	RAS  pin.Pin     // SDRAM Row Address Strobe
	CAS  pin.Pin     // SDRAM Column Address Strobe
	WE   pin.Pin     // SDRAM Write Enable
	ADDR [13]pin.Pin // SDRAM Address
	DATA [32]pin.Pin // SDRAM Bidirectional data bus
	BA   [2]pin.Pin  // SDRAM Bank Address
	NBL  [4]pin.Pin  // SDRAM Output Byte Mask for write accesses (memory signal name: DQM[3:0]

	RT                 uint32 // Refresh rate in the number of SDCLK clock cycles between the refresh cycles
	PowerUpDelayCycles uint32
}

func Configure(config Config) error {
	config0 := config.Bank[SDRAM0]
	config1 := config.Bank[SDRAM1]

	// Enable the peripheral in RCC.
	rcc.Rcc.Ahb3enr.SetFmcen(true)
	_ = rcc.Rcc.Ahb3enr.GetFmcen()

	// GPIO banks used by the FMC SDRAM signals.
	rcc.Rcc.Ahb4enr.SetGpiocen(true)
	rcc.Rcc.Ahb4enr.SetGpioden(true)
	rcc.Rcc.Ahb4enr.SetGpioeen(true)
	rcc.Rcc.Ahb4enr.SetGpiofen(true)
	rcc.Rcc.Ahb4enr.SetGpiogen(true)
	rcc.Rcc.Ahb4enr.SetGpiohen(true)

	// Configure the pins.
	configurePin(config.CLK)
	configurePin(config.CKE)
	configurePin(config.CS)
	configurePin(config.RAS)
	configurePin(config.CAS)
	configurePin(config.WE)

	for _, p := range config.ADDR {
		configurePin(p)
	}

	for _, p := range config.DATA {
		configurePin(p)
	}

	for _, p := range config.BA {
		configurePin(p)
	}

	for _, p := range config.NBL {
		configurePin(p)
	}

	if config0.Enable {
		// Program the memory device features into the FMC_SDCR1 register.
		var cr fmc.RegisterSdcr1Type
		cr.SetMwid(uint8(config0.BusWidth))
		cr.SetNr(uint8(config0.RowBits))
		cr.SetNc(uint8(config0.ColumnBits))
		cr.SetNb(bool(config0.InternalBanks))
		cr.SetWp(config0.WriteProtection)
		cr.SetCas(uint8(config0.CASLatency))

		// These can only be set in SDCR1.
		cr.SetSdclk(uint8(config0.ClockPeriod))
		cr.SetRburst(config0.BurstRead)
		cr.SetRpipe(uint8(config0.ReadPipeDelay))

		// Write all to the SDCR1 register.
		fmc.Fmc.Sdcr1.Store(uint32(cr))

		// Program the memory device timing into the FMC_SDTR1 register.
		var tr fmc.RegisterSdtr1Type
		tr.SetTrcd(config0.Timing.TRCD - 1)
		tr.SetTwr(config0.Timing.TWR - 1)
		tr.SetTxsr(config0.Timing.TXSR - 1)
		tr.SetTras(config0.Timing.TRAS - 1)
		tr.SetTmrd(config0.Timing.TMRD - 1)

		// The TRP and TRC timings must be programmed in the FMC_SDTR1 register.
		tr.SetTrp(config0.Timing.TRP - 1)
		tr.SetTrc(config0.Timing.TRC - 1)

		// Write all to the SDTR1 register.
		fmc.Fmc.Sdtr1.Store(uint32(tr))
	}

	if config1.Enable {
		// Program the memory device features into the FMC_SDCR2 register.
		var cr fmc.RegisterSdcr2Type
		cr.SetMwid(uint8(config1.BusWidth))
		cr.SetNr(uint8(config1.RowBits))
		cr.SetNc(uint8(config1.ColumnBits))
		cr.SetNb(bool(config1.InternalBanks))
		cr.SetWp(config1.WriteProtection)

		// Write all to the SDCR2 register.
		fmc.Fmc.Sdcr2.Store(uint32(cr))

		// Program the memory device timing into the FMC_SDTR2 register.
		var tr fmc.RegisterSdtr2Type
		tr.SetTrcd(config1.Timing.TRCD - 1)
		tr.SetTwr(config1.Timing.TWR - 1)
		tr.SetTxsr(config1.Timing.TXSR - 1)
		tr.SetTras(config1.Timing.TRAS - 1)
		tr.SetTmrd(config1.Timing.TMRD - 1)

		// Write all to the SDTR2 register.
		fmc.Fmc.Sdtr2.Store(uint32(tr))
	}

	// Enable the FMC before issuing commands.
	fmc.Fmc.Bcr1.SetFmcen(true)

	var target uint32
	if config0.Enable {
		target |= sdcmdTargetBank1
	}

	if config1.Enable {
		target |= sdcmdTargetBank2
	}

	// Perform the initialization sequence.
	sendSDRAMCommand(target, cmdClockEnable, 1, 0)

	// Wait during the prescribed delay period.
	delayCycles(config.PowerUpDelayCycles)

	sendSDRAMCommand(target, cmdPrechargeAll, 1, 0)
	sendSDRAMCommand(target, cmdAutoRefresh, 8, 0)

	// Apply the Load Mode to each SDRAM chip.
	if config0.Enable {
		sendSDRAMCommand(sdcmdTargetBank1, cmdLoadModeReg, 1, config0.LoadMode)
	}
	if config1.Enable {
		sendSDRAMCommand(sdcmdTargetBank2, cmdLoadModeReg, 1, config1.LoadMode)
	}

	// Program the refresh rate.
	fmc.Fmc.Sdrtr.Store(config.RT)

	return nil
}

func sendSDRAMCommand(target, mode, refreshCount, mrd uint32) {
	cmd := (mode & 0x7) |
		target |
		((refreshCount - 1) << 5) |
		(mrd << 9)

	fmc.Fmc.Sdcmr.Store(cmd)
	cortexm.DSB()

	// Wait for the command to be accepted (MODE clears to 0).
	for fmc.Fmc.Sdcmr.Load()&0x7 != 0 {
	}

	cortexm.DSB()
}

func delayCycles(cycles uint32) {
	for i := uint32(0); i < cycles; i++ {
		// TODO: Consider using a more efficient delay mechanism.
		_ = fmc.Fmc.Sdsr.Load()
	}
}

func configurePin(p pin.Pin) {
	if p == 0 {
		return
	}

	p.SetMode(pin.AltFunction | pin.Alt12)
	p.SetSpeedMode(pin.VeryHighSpeed)
	p.SetPullMode(pin.PullUp)
	p.SetOutputMode(pin.PushPull)
}
