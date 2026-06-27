//go:build stm32h7x7

package hal

import (
	"pkg.si-go.dev/chip/arm/cortexm"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/flash"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/pwr"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/syscfg"
)

// ClockSource identifies a hardware clock source for selectable peripheral
// clock muxes. Used as the value type for the Default*ClockSource constants.
type ClockSource uint8

const (
	ClockSourceNone ClockSource = iota
	ClockSourcePclk1
	ClockSourcePclk2
	ClockSourcePclk3
	ClockSourcePclk4
	ClockSourcePll1p
	ClockSourcePll1q
	ClockSourcePll1r
	ClockSourcePll2p
	ClockSourcePll2q
	ClockSourcePll2r
	ClockSourcePll3p
	ClockSourcePll3q
	ClockSourcePll3r
	ClockSourceHse
	ClockSourceLse
	ClockSourceHsi
	ClockSourceCsi
	ClockSourceLsi
	ClockSourceRc48
	ClockSourcePerCk
	ClockSourceI2sCkin
	ClockSourceTim2
	ClockSourceCpu1
	ClockSourceSysClk
	ClockSourceHclk3
	ClockSourceHseRtc
)

// Divider expresses an integer clock divider. Values are the actual division
// factor (Div2 = divide by 2). Used both at compile time for frequency math
// and at runtime to select the matching register field encoding.
type Divider uint16

const (
	Div1   Divider = 1
	Div2   Divider = 2
	Div4   Divider = 4
	Div8   Divider = 8
	Div16  Divider = 16
	Div32  Divider = 32
	Div64  Divider = 64
	Div128 Divider = 128
	Div256 Divider = 256
	Div512 Divider = 512
)

// Internal oscillator frequencies — fixed by hardware.
const (
	HsiFrequencyHz  uint64 = 64_000_000
	CsiFrequencyHz  uint64 = 4_000_000
	LsiFrequencyHz  uint64 = 32_000
	Rc48FrequencyHz uint64 = 48_000_000
)

// Default board oscillator frequencies. Boards may use different crystals;
// these are only used by the default ConfigureClocks.
const (
	DefaultHseFrequencyHz uint64 = 16_000_000
	DefaultLseFrequencyHz uint64 = 32_768
)

// Default oscillator enables.
const (
	DefaultEnableRc48 = true
	DefaultEnableHse  = false
	DefaultEnableLse  = false
)

// Default PLL input dividers. The PLL input frequency must be in the range
// 1–16 MHz; check the reference manual for valid combinations with the
// configured PLL source.
const (
	DefaultDivm1 uint8 = 4  // PLL1 input = 64M / 4 = 16 MHz
	DefaultDivm2 uint8 = 32 // PLL2 input = 64M / 32 = 2 MHz
	DefaultDivm3 uint8 = 16 // PLL3 input = 64M / 16 = 4 MHz
)

// Default PLL multipliers (Divn) and output dividers.
//
// PLL1 VCO = (input × Divn) where input = HSI / Divm1 = 16 MHz, so
// VCO = 16 × 60 = 960 MHz. Output frequencies follow from the dividers.
const (
	DefaultDivn1 uint16 = 60
	DefaultDivn2 uint16 = 300
	DefaultDivn3 uint16 = 16

	DefaultDivp1 uint8 = 2 // 960 / 2 = 480 MHz (CPU)
	DefaultDivq1 uint8 = 8 // 960 / 8 = 120 MHz (SPI1/2/3, RNG, SDMMC alt)
	DefaultDivr1 uint8 = 2 // 960 / 2 = 480 MHz

	DefaultDivp2 uint8 = 4 // 600 / 4 = 150 MHz
	DefaultDivq2 uint8 = 4 // 600 / 4 = 150 MHz
	DefaultDivr2 uint8 = 6 // 600 / 6 = 100 MHz (SDMMC clock source)

	DefaultDivp3 uint8 = 2 // 64 / 2 = 32 MHz (placeholder; reconfigure as needed)
	DefaultDivq3 uint8 = 2 // 64 / 2 = 32 MHz
	DefaultDivr3 uint8 = 2 // 64 / 2 = 32 MHz
)

// Default bus prescalers from the CPU clock down through the AHB and APB
// hierarchy. With CPU = 480 MHz, HPRE = /2 yields HCLK = 240 MHz which
// satisfies the 240 MHz HCLK ceiling for VOS0.
const (
	DefaultD1cpre    Divider = Div1
	DefaultHpre      Divider = Div2
	DefaultD1ppre    Divider = Div2
	DefaultD2ppre1   Divider = Div2
	DefaultD2ppre2   Divider = Div2
	DefaultD3ppre    Divider = Div2
	DefaultDivRtcHse Divider = Div32
)

// Default peripheral clock source selections.
const (
	DefaultPllSource         = ClockSourceHsi
	DefaultRtcClockSource    = ClockSourceLse
	DefaultI2c13Source       = ClockSourcePclk1
	DefaultI2c4Source        = ClockSourcePclk4
	DefaultUsart16Source     = ClockSourcePclk2
	DefaultUsart234578Source = ClockSourcePclk1
	DefaultSpi123ClockSource = ClockSourcePll1q
	DefaultSpi45ClockSource  = ClockSourceHsi
	DefaultSpi6ClockSource   = ClockSourceHsi
	DefaultSdmmcClockSource  = ClockSourcePll2r
	DefaultRngClockSource    = ClockSourceRc48
	DefaultFmcClockSource    = ClockSourceHclk3
)

// Default frequencies derived from the configuration above. These are the
// values ConfigureClocks will write into the corresponding globals.
//
// Boards that override ConfigureClocks compute their own equivalents and
// write to the globals directly.
const (
	defaultPll1VcoHz uint64 = (HsiFrequencyHz / uint64(DefaultDivm1)) * uint64(DefaultDivn1)
	defaultPll2VcoHz uint64 = (HsiFrequencyHz / uint64(DefaultDivm2)) * uint64(DefaultDivn2)
	defaultPll3VcoHz uint64 = (HsiFrequencyHz / uint64(DefaultDivm3)) * uint64(DefaultDivn3)

	DefaultDivp1FrequencyHz uint64 = defaultPll1VcoHz / uint64(DefaultDivp1)
	DefaultDivq1FrequencyHz uint64 = defaultPll1VcoHz / uint64(DefaultDivq1)
	DefaultDivr1FrequencyHz uint64 = defaultPll1VcoHz / uint64(DefaultDivr1)
	DefaultDivp2FrequencyHz uint64 = defaultPll2VcoHz / uint64(DefaultDivp2)
	DefaultDivq2FrequencyHz uint64 = defaultPll2VcoHz / uint64(DefaultDivq2)
	DefaultDivr2FrequencyHz uint64 = defaultPll2VcoHz / uint64(DefaultDivr2)
	DefaultDivp3FrequencyHz uint64 = defaultPll3VcoHz / uint64(DefaultDivp3)
	DefaultDivq3FrequencyHz uint64 = defaultPll3VcoHz / uint64(DefaultDivq3)
	DefaultDivr3FrequencyHz uint64 = defaultPll3VcoHz / uint64(DefaultDivr3)

	DefaultCpu1FrequencyHz uint64 = DefaultDivp1FrequencyHz / uint64(DefaultD1cpre)
	DefaultHpreFrequencyHz uint64 = DefaultDivp1FrequencyHz / uint64(DefaultHpre)

	DefaultPclk1FrequencyHz uint64 = DefaultHpreFrequencyHz / uint64(DefaultD2ppre1)
	DefaultPclk2FrequencyHz uint64 = DefaultHpreFrequencyHz / uint64(DefaultD2ppre2)
	DefaultPclk3FrequencyHz uint64 = DefaultHpreFrequencyHz / uint64(DefaultD1ppre)
	DefaultPclk4FrequencyHz uint64 = DefaultHpreFrequencyHz / uint64(DefaultD3ppre)
)

// Configured clock frequencies. Populated by ConfigureClocks (or a board's
// equivalent) during postinit. Peripheral drivers read these to derive
// their own timings.
//
// Reading these before postinit returns zero — drivers must not be used
// before clocks are configured.
var (
	D1cpre    Divider
	Hpre      Divider
	D1ppre    Divider
	D2ppre1   Divider
	D2ppre2   Divider
	D3ppre    Divider
	DivRtcHse Divider

	HseFrequencyHz uint64
	LseFrequencyHz uint64

	Cpu1FrequencyHz uint64
	HpreFrequencyHz uint64

	Hclk1FrequencyHz uint64
	Hclk2FrequencyHz uint64
	Hclk3FrequencyHz uint64
	Hclk4FrequencyHz uint64

	Pclk1FrequencyHz uint64
	Pclk2FrequencyHz uint64
	Pclk3FrequencyHz uint64
	Pclk4FrequencyHz uint64

	Divp1FrequencyHz uint64
	Divq1FrequencyHz uint64
	Divr1FrequencyHz uint64

	Divp2FrequencyHz uint64
	Divq2FrequencyHz uint64
	Divr2FrequencyHz uint64

	Divp3FrequencyHz uint64
	Divq3FrequencyHz uint64
	Divr3FrequencyHz uint64

	PllSourceFrequencyHz         uint64
	RtcSourceFrequencyHz         uint64
	I2c13SourceFrequencyHz       uint64
	I2c4SourceFrequencyHz        uint64
	Usart16SourceFrequencyHz     uint64
	Usart234578SourceFrequencyHz uint64
	Spi123SourceFrequencyHz      uint64
	Spi45SourceFrequencyHz       uint64
	Spi6SourceFrequencyHz        uint64
	SdmmcSourceFrequencyHz       uint64
	RngSourceFrequencyHz         uint64
	FmcSourceFrequencyHz         uint64
	QspiSourceFrequencyHz        uint64
	LtdcSourceFrequencyHz        uint64
	DsiHostReferenceFrequencyHz  uint64
	DsiHostPhyFrequencyHz        uint64
	DsiHostLaneByteFrequencyHz   uint64
	DsiHostTxEscapeFrequencyHz   uint64
)

// ConfigureClocks programs the chip's clock tree to the default configuration.
// It reads only constants and writes only to hardware registers. Safe to call
// from preinit because no global state is touched.
//
// Boards that need different settings replace this entirely with their own
// version (typically by providing their own preinit that does not call
// hal.ConfigureClocks).
func ConfigureClocks() {
	state := cortexm.DisableInterrupts()

	// VOS0 (480 MHz) is reachable only from VOS1.
	if pwr.Pwr.D3cr.GetVos() != pwr.RegisterD3crFieldVosEnumScale1 {
		pwr.Pwr.D3cr.SetVos(pwr.RegisterD3crFieldVosEnumScale1)
	}
	pwr.Pwr.Cr3.SetSden(false)

	rcc.Rcc.Apb4enr.SetSyscfgen(true)
	syscfg.Syscfg.Pwrcr.SetOden(true)
	for !pwr.Pwr.D3cr.GetVosrdy() {
	}

	if DefaultEnableRc48 {
		rcc.Rcc.Cr.SetRc48on(true)
		for !rcc.Rcc.Cr.GetRc48rdy() {
		}
	}
	if DefaultEnableHse {
		rcc.Rcc.Cr.SetHseon(true)
		for !rcc.Rcc.Cr.GetHserdy() {
		}
	}
	if DefaultEnableLse {
		rcc.Rcc.Bdcr.SetLseon(true)
		for !rcc.Rcc.Bdcr.GetLseon() {
		}
	}

	// RTC source mux.
	rcc.Rcc.Cfgr.SetRtcpre(uint8(DefaultDivRtcHse))
	switch DefaultRtcClockSource {
	case ClockSourceLse:
		rcc.Rcc.Bdcr.SetRtcsrc(1)
	case ClockSourceLsi:
		rcc.Rcc.Bdcr.SetRtcsrc(2)
	case ClockSourceHseRtc:
		rcc.Rcc.Bdcr.SetRtcsrc(3)
	}

	// PLL source mux.
	switch DefaultPllSource {
	case ClockSourceHse:
		rcc.Rcc.Pllckselr.SetPllsrc(rcc.RegisterPllckselrFieldPllsrcEnumHse)
	case ClockSourceHsi:
		rcc.Rcc.Pllckselr.SetPllsrc(rcc.RegisterPllckselrFieldPllsrcEnumHsi)
	case ClockSourceCsi:
		rcc.Rcc.Pllckselr.SetPllsrc(rcc.RegisterPllckselrFieldPllsrcEnumCsi)
	}

	// PLL input dividers.
	rcc.Rcc.Pllckselr.SetDivm1(DefaultDivm1)
	rcc.Rcc.Pllckselr.SetDivm2(DefaultDivm2)
	rcc.Rcc.Pllckselr.SetDivm3(DefaultDivm3)

	// PLL input range. Hardcoded for the default config (HSI/Divm1=16MHz,
	// HSI/Divm2=2MHz, HSI/Divm3=4MHz). Boards changing Divm values must
	// adjust these.
	rcc.Rcc.Pllcfgr.SetPll1rge(rcc.RegisterPllcfgrFieldPll1rgeEnum8to16mhz)
	rcc.Rcc.Pllcfgr.SetPll2rge(rcc.RegisterPllcfgrFieldPll2rgeEnum1to2mhz)
	rcc.Rcc.Pllcfgr.SetPll3rge(rcc.RegisterPllcfgrFieldPll3rgeEnum2to4mhz)

	// PLL multipliers and output dividers. Register fields are 0-based
	// (value = N-1).
	rcc.Rcc.Pll1divr.SetDivn1(DefaultDivn1 - 1)
	rcc.Rcc.Pll1divr.SetDivp1(DefaultDivp1 - 1)
	rcc.Rcc.Pll1divr.SetDivq1(DefaultDivq1 - 1)
	rcc.Rcc.Pll1divr.SetDivr1(DefaultDivr1 - 1)

	rcc.Rcc.Pll2divr.SetDivn2(DefaultDivn2 - 1)
	rcc.Rcc.Pll2divr.SetDivp2(DefaultDivp2 - 1)
	rcc.Rcc.Pll2divr.SetDivq2(DefaultDivq2 - 1)
	rcc.Rcc.Pll2divr.SetDivr2(DefaultDivr2 - 1)

	rcc.Rcc.Pll3divr.SetDivn3(DefaultDivn3 - 1)
	rcc.Rcc.Pll3divr.SetDivp3(DefaultDivp3 - 1)
	rcc.Rcc.Pll3divr.SetDivq3(DefaultDivq3 - 1)
	rcc.Rcc.Pll3divr.SetDivr3(DefaultDivr3 - 1)

	// Enable PLL outputs.
	rcc.Rcc.Pllcfgr.SetDivp1en(true)
	rcc.Rcc.Pllcfgr.SetDivq1en(true)
	rcc.Rcc.Pllcfgr.SetDivr1en(true)
	rcc.Rcc.Pllcfgr.SetDivp2en(true)
	rcc.Rcc.Pllcfgr.SetDivq2en(true)
	rcc.Rcc.Pllcfgr.SetDivr2en(true)
	rcc.Rcc.Pllcfgr.SetDivp3en(true)
	rcc.Rcc.Pllcfgr.SetDivq3en(true)
	rcc.Rcc.Pllcfgr.SetDivr3en(true)

	// Start the PLLs.
	rcc.Rcc.Cr.SetPll1on(true)
	rcc.Rcc.Cr.SetPll2on(true)
	rcc.Rcc.Cr.SetPll3on(true)

	// Bus prescalers — apply before switching SYSCLK so HCLK and the APBs
	// don't briefly overshoot when SYSCLK jumps from HSI to PLL1.
	setD1cpre(DefaultD1cpre)
	setHpre(DefaultHpre)
	setD1ppre(DefaultD1ppre)
	setD2ppre1(DefaultD2ppre1)
	setD2ppre2(DefaultD2ppre2)
	setD3ppre(DefaultD3ppre)

	// Wait for PLL1 lock.
	for !rcc.Rcc.Cr.GetPll1rdy() {
	}

	// Switch SYSCLK to PLL1.
	rcc.Rcc.Cfgr.SetSw(rcc.RegisterCfgrFieldSwEnumPll1)
	for rcc.Rcc.Cfgr.GetSws() != rcc.RegisterCfgrFieldSwsEnumPll1 {
	}

	// Flash wait states for VOS0 + 240 MHz HCLK + AXI 240 MHz.
	flash.Flash.Bank[0].Acr.SetLatency(4)
	for flash.Flash.Bank[0].Acr.GetLatency() != 4 {
	}
	flash.Flash.Bank[0].Acr.SetWrhighfreq(3)
	for flash.Flash.Bank[0].Acr.GetWrhighfreq() != 3 {
	}

	// GPIO clocks. All banks; boards can disable individual ones later.
	rcc.Rcc.Ahb4enr.SetGpioaen(true)
	rcc.Rcc.Ahb4enr.SetGpioben(true)
	rcc.Rcc.Ahb4enr.SetGpiocen(true)
	rcc.Rcc.Ahb4enr.SetGpioden(true)
	rcc.Rcc.Ahb4enr.SetGpioeen(true)
	rcc.Rcc.Ahb4enr.SetGpiofen(true)
	rcc.Rcc.Ahb4enr.SetGpiogen(true)
	rcc.Rcc.Ahb4enr.SetGpiohen(true)
	rcc.Rcc.Ahb4enr.SetGpioien(true)
	rcc.Rcc.Ahb4enr.SetGpiojen(true)
	rcc.Rcc.Ahb4enr.SetGpioken(true)

	// DAC, ADC, RNG infrastructure clocks.
	rcc.Rcc.Apb1lenr.SetDac12en(true)
	rcc.Rcc.D3ccipr.SetAdcsrc(0b00) // PLL2.
	rcc.Rcc.Ahb1enr.SetAdc12en(true)
	rcc.Rcc.Ahb4enr.SetAdc3en(true)
	rcc.Rcc.Ahb2enr.SetRngen(true)

	// Peripheral source muxes — hardcoded for the default config since
	// every selection is a known constant.

	// I2C1-3.
	switch DefaultI2c13Source {
	case ClockSourcePclk1:
		rcc.Rcc.D2ccip2r.SetI2c123src(0)
	case ClockSourcePll3r:
		rcc.Rcc.D2ccip2r.SetI2c123src(1)
	case ClockSourceHsi:
		rcc.Rcc.D2ccip2r.SetI2c123src(2)
	case ClockSourceCsi:
		rcc.Rcc.D2ccip2r.SetI2c123src(3)
	}

	// I2C4.
	switch DefaultI2c4Source {
	case ClockSourcePclk4:
		rcc.Rcc.D3ccipr.SetI2c4src(0)
	case ClockSourcePll3r:
		rcc.Rcc.D3ccipr.SetI2c4src(1)
	case ClockSourceHsi:
		rcc.Rcc.D3ccipr.SetI2c4src(2)
	case ClockSourceCsi:
		rcc.Rcc.D3ccipr.SetI2c4src(3)
	}

	// USART1/6.
	switch DefaultUsart16Source {
	case ClockSourcePclk2:
		rcc.Rcc.D2ccip2r.SetUsart16src(0)
	case ClockSourcePll2q:
		rcc.Rcc.D2ccip2r.SetUsart16src(1)
	case ClockSourcePll3q:
		rcc.Rcc.D2ccip2r.SetUsart16src(2)
	case ClockSourceHsi:
		rcc.Rcc.D2ccip2r.SetUsart16src(3)
	case ClockSourceCsi:
		rcc.Rcc.D2ccip2r.SetUsart16src(4)
	case ClockSourceLse:
		rcc.Rcc.D2ccip2r.SetUsart16src(5)
	}

	// USART2/3/4/5/7/8.
	switch DefaultUsart234578Source {
	case ClockSourcePclk1:
		rcc.Rcc.D2ccip2r.SetUsart234578src(0)
	case ClockSourcePll2q:
		rcc.Rcc.D2ccip2r.SetUsart234578src(1)
	case ClockSourcePll3q:
		rcc.Rcc.D2ccip2r.SetUsart234578src(2)
	case ClockSourceHsi:
		rcc.Rcc.D2ccip2r.SetUsart234578src(3)
	case ClockSourceCsi:
		rcc.Rcc.D2ccip2r.SetUsart234578src(4)
	case ClockSourceLse:
		rcc.Rcc.D2ccip2r.SetUsart234578src(5)
	}

	// SPI1/2/3.
	switch DefaultSpi123ClockSource {
	case ClockSourcePll1q:
		rcc.Rcc.D2ccip1r.SetSpi123src(0)
	case ClockSourcePll2p:
		rcc.Rcc.D2ccip1r.SetSpi123src(1)
	case ClockSourcePll3p:
		rcc.Rcc.D2ccip1r.SetSpi123src(2)
	case ClockSourceI2sCkin:
		rcc.Rcc.D2ccip1r.SetSpi123src(3)
	case ClockSourcePerCk:
		rcc.Rcc.D2ccip1r.SetSpi123src(4)
	}

	// SPI4/5.
	switch DefaultSpi45ClockSource {
	case ClockSourcePclk2:
		rcc.Rcc.D2ccip1r.SetSpi45src(0)
	case ClockSourcePll2q:
		rcc.Rcc.D2ccip1r.SetSpi45src(1)
	case ClockSourcePll3q:
		rcc.Rcc.D2ccip1r.SetSpi45src(2)
	case ClockSourceHsi:
		rcc.Rcc.D2ccip1r.SetSpi45src(3)
	case ClockSourceCsi:
		rcc.Rcc.D2ccip1r.SetSpi45src(4)
	case ClockSourceHse:
		rcc.Rcc.D2ccip1r.SetSpi45src(5)
	}

	// SPI6.
	switch DefaultSpi6ClockSource {
	case ClockSourcePclk4:
		rcc.Rcc.D3ccipr.SetSpi6src(0)
	case ClockSourcePll2q:
		rcc.Rcc.D3ccipr.SetSpi6src(1)
	case ClockSourcePll3q:
		rcc.Rcc.D3ccipr.SetSpi6src(2)
	case ClockSourceHsi:
		rcc.Rcc.D3ccipr.SetSpi6src(3)
	case ClockSourceCsi:
		rcc.Rcc.D3ccipr.SetSpi6src(4)
	case ClockSourceHse:
		rcc.Rcc.D3ccipr.SetSpi6src(5)
	}

	// SDMMC.
	switch DefaultSdmmcClockSource {
	case ClockSourcePll1q:
		rcc.Rcc.D1ccipr.SetSdmmcsrc(false)
	case ClockSourcePll2r:
		rcc.Rcc.D1ccipr.SetSdmmcsrc(true)
	}

	// RNG.
	switch DefaultRngClockSource {
	case ClockSourceRc48:
		rcc.Rcc.D2ccip2r.SetRngsrc(0)
	case ClockSourcePll1q:
		rcc.Rcc.D2ccip2r.SetRngsrc(1)
	case ClockSourceLse:
		rcc.Rcc.D2ccip2r.SetRngsrc(2)
	case ClockSourceLsi:
		rcc.Rcc.D2ccip2r.SetRngsrc(3)
	}

	// FMC.
	switch DefaultFmcClockSource {
	case ClockSourceHclk3:
		rcc.Rcc.D1ccipr.SetFmcsrc(0)
	case ClockSourcePll1q:
		rcc.Rcc.D1ccipr.SetFmcsrc(1)
	case ClockSourcePll2r:
		rcc.Rcc.D1ccipr.SetFmcsrc(2)
	case ClockSourcePerCk:
		rcc.Rcc.D1ccipr.SetFmcsrc(3)
	}

	cortexm.EnableInterrupts(state)
}

// SetDefaultFrequencies writes the chip-default frequency values into
// the package globals. Called by the default postinit after ConfigureClocks
// has applied the matching configuration.
//
// Boards that program different clock values must publish their actual
// frequencies (either by writing to the globals directly or by providing
// their own publish function) so peripheral drivers see consistent values.
func SetDefaultFrequencies() {
	D1cpre = DefaultD1cpre
	Hpre = DefaultHpre
	D1ppre = DefaultD1ppre
	D2ppre1 = DefaultD2ppre1
	D2ppre2 = DefaultD2ppre2
	D3ppre = DefaultD3ppre
	DivRtcHse = DefaultDivRtcHse

	HseFrequencyHz = DefaultHseFrequencyHz
	LseFrequencyHz = DefaultLseFrequencyHz

	Cpu1FrequencyHz = DefaultCpu1FrequencyHz
	HpreFrequencyHz = DefaultHpreFrequencyHz

	Hclk1FrequencyHz = DefaultHpreFrequencyHz
	Hclk2FrequencyHz = DefaultHpreFrequencyHz
	Hclk3FrequencyHz = DefaultHpreFrequencyHz
	Hclk4FrequencyHz = DefaultHpreFrequencyHz

	Pclk1FrequencyHz = DefaultPclk1FrequencyHz
	Pclk2FrequencyHz = DefaultPclk2FrequencyHz
	Pclk3FrequencyHz = DefaultPclk3FrequencyHz
	Pclk4FrequencyHz = DefaultPclk4FrequencyHz

	Divp1FrequencyHz = DefaultDivp1FrequencyHz
	Divq1FrequencyHz = DefaultDivq1FrequencyHz
	Divr1FrequencyHz = DefaultDivr1FrequencyHz
	Divp2FrequencyHz = DefaultDivp2FrequencyHz
	Divq2FrequencyHz = DefaultDivq2FrequencyHz
	Divr2FrequencyHz = DefaultDivr2FrequencyHz
	Divp3FrequencyHz = DefaultDivp3FrequencyHz
	Divq3FrequencyHz = DefaultDivq3FrequencyHz
	Divr3FrequencyHz = DefaultDivr3FrequencyHz

	PllSourceFrequencyHz = HsiFrequencyHz

	switch DefaultRtcClockSource {
	case ClockSourceLse:
		RtcSourceFrequencyHz = DefaultLseFrequencyHz
	case ClockSourceLsi:
		RtcSourceFrequencyHz = LsiFrequencyHz
	case ClockSourceHseRtc:
		RtcSourceFrequencyHz = DefaultHseFrequencyHz / uint64(DefaultDivRtcHse)
	}

	I2c13SourceFrequencyHz = DefaultPclk1FrequencyHz
	I2c4SourceFrequencyHz = DefaultPclk4FrequencyHz
	Usart16SourceFrequencyHz = DefaultPclk2FrequencyHz
	Usart234578SourceFrequencyHz = DefaultPclk1FrequencyHz
	Spi123SourceFrequencyHz = DefaultDivq1FrequencyHz
	Spi45SourceFrequencyHz = HsiFrequencyHz
	Spi6SourceFrequencyHz = HsiFrequencyHz
	SdmmcSourceFrequencyHz = DefaultDivr2FrequencyHz
	RngSourceFrequencyHz = Rc48FrequencyHz

	// SysTick on the new CPU clock.
	cortexm.UpdateSysTickFrequency(uint32(Cpu1FrequencyHz))
}

// TimerMultiplier returns the multiplier between the APB peripheral clock
// and the timer clock for a given APB prescaler. Used by timer drivers to
// derive timer-clock frequencies from PCLK frequencies.
func TimerMultiplier(d2ppre Divider) uint64 {
	timpre := rcc.Rcc.Cfgr.GetTimpre()
	switch d2ppre {
	case Div1:
		return 1
	case Div2:
		return 2
	case Div4, Div8, Div16:
		if timpre {
			return 4
		}
		return 2
	default:
		panic("invalid divider value")
	}
}

// Helper functions for divider→register-field translation. Each takes a
// constant Divider at every call site, so the compiler reduces these to a
// single register write at -O2.

func setD1cpre(d Divider) {
	switch d {
	case Div1:
		rcc.Rcc.D1cfgr.SetD1cpre(rcc.RegisterD1cfgrFieldD1cpreEnumDiv1)
	case Div2:
		rcc.Rcc.D1cfgr.SetD1cpre(rcc.RegisterD1cfgrFieldD1cpreEnumDiv2)
	case Div4:
		rcc.Rcc.D1cfgr.SetD1cpre(rcc.RegisterD1cfgrFieldD1cpreEnumDiv4)
	case Div8:
		rcc.Rcc.D1cfgr.SetD1cpre(rcc.RegisterD1cfgrFieldD1cpreEnumDiv8)
	case Div16:
		rcc.Rcc.D1cfgr.SetD1cpre(rcc.RegisterD1cfgrFieldD1cpreEnumDiv16)
	case Div64:
		rcc.Rcc.D1cfgr.SetD1cpre(rcc.RegisterD1cfgrFieldD1cpreEnumDiv64)
	case Div128:
		rcc.Rcc.D1cfgr.SetD1cpre(rcc.RegisterD1cfgrFieldD1cpreEnumDiv128)
	case Div256:
		rcc.Rcc.D1cfgr.SetD1cpre(rcc.RegisterD1cfgrFieldD1cpreEnumDiv256)
	case Div512:
		rcc.Rcc.D1cfgr.SetD1cpre(rcc.RegisterD1cfgrFieldD1cpreEnumDiv512)
	default:
		panic("invalid D1CPRE divider")
	}
}

func setHpre(d Divider) {
	switch d {
	case Div1:
		rcc.Rcc.D1cfgr.SetHpre(rcc.RegisterD1cfgrFieldHpreEnumDiv1)
	case Div2:
		rcc.Rcc.D1cfgr.SetHpre(rcc.RegisterD1cfgrFieldHpreEnumDiv2)
	case Div4:
		rcc.Rcc.D1cfgr.SetHpre(rcc.RegisterD1cfgrFieldHpreEnumDiv4)
	case Div8:
		rcc.Rcc.D1cfgr.SetHpre(rcc.RegisterD1cfgrFieldHpreEnumDiv8)
	case Div16:
		rcc.Rcc.D1cfgr.SetHpre(rcc.RegisterD1cfgrFieldHpreEnumDiv16)
	case Div64:
		rcc.Rcc.D1cfgr.SetHpre(rcc.RegisterD1cfgrFieldHpreEnumDiv64)
	case Div128:
		rcc.Rcc.D1cfgr.SetHpre(rcc.RegisterD1cfgrFieldHpreEnumDiv128)
	case Div256:
		rcc.Rcc.D1cfgr.SetHpre(rcc.RegisterD1cfgrFieldHpreEnumDiv256)
	case Div512:
		rcc.Rcc.D1cfgr.SetHpre(rcc.RegisterD1cfgrFieldHpreEnumDiv512)
	default:
		panic("invalid HPRE divider")
	}
}

func setD1ppre(d Divider) {
	switch d {
	case Div1:
		rcc.Rcc.D1cfgr.SetD1ppre(rcc.RegisterD1cfgrFieldD1ppreEnumDiv1)
	case Div2:
		rcc.Rcc.D1cfgr.SetD1ppre(rcc.RegisterD1cfgrFieldD1ppreEnumDiv2)
	case Div4:
		rcc.Rcc.D1cfgr.SetD1ppre(rcc.RegisterD1cfgrFieldD1ppreEnumDiv4)
	case Div8:
		rcc.Rcc.D1cfgr.SetD1ppre(rcc.RegisterD1cfgrFieldD1ppreEnumDiv8)
	case Div16:
		rcc.Rcc.D1cfgr.SetD1ppre(rcc.RegisterD1cfgrFieldD1ppreEnumDiv16)
	default:
		panic("invalid D1PPRE divider")
	}
}

func setD2ppre1(d Divider) {
	switch d {
	case Div1:
		rcc.Rcc.D2cfgr.SetD2ppre1(rcc.RegisterD2cfgrFieldD2ppre1EnumDiv1)
	case Div2:
		rcc.Rcc.D2cfgr.SetD2ppre1(rcc.RegisterD2cfgrFieldD2ppre1EnumDiv2)
	case Div4:
		rcc.Rcc.D2cfgr.SetD2ppre1(rcc.RegisterD2cfgrFieldD2ppre1EnumDiv4)
	case Div8:
		rcc.Rcc.D2cfgr.SetD2ppre1(rcc.RegisterD2cfgrFieldD2ppre1EnumDiv8)
	case Div16:
		rcc.Rcc.D2cfgr.SetD2ppre1(rcc.RegisterD2cfgrFieldD2ppre1EnumDiv16)
	default:
		panic("invalid D2PPRE1 divider")
	}
}

func setD2ppre2(d Divider) {
	switch d {
	case Div1:
		rcc.Rcc.D2cfgr.SetD2ppre2(rcc.RegisterD2cfgrFieldD2ppre2EnumDiv1)
	case Div2:
		rcc.Rcc.D2cfgr.SetD2ppre2(rcc.RegisterD2cfgrFieldD2ppre2EnumDiv2)
	case Div4:
		rcc.Rcc.D2cfgr.SetD2ppre2(rcc.RegisterD2cfgrFieldD2ppre2EnumDiv4)
	case Div8:
		rcc.Rcc.D2cfgr.SetD2ppre2(rcc.RegisterD2cfgrFieldD2ppre2EnumDiv8)
	case Div16:
		rcc.Rcc.D2cfgr.SetD2ppre2(rcc.RegisterD2cfgrFieldD2ppre2EnumDiv16)
	default:
		panic("invalid D2PPRE2 divider")
	}
}

func setD3ppre(d Divider) {
	switch d {
	case Div1:
		rcc.Rcc.D3cfgr.SetD3ppre(rcc.RegisterD3cfgrFieldD3ppreEnumDiv1)
	case Div2:
		rcc.Rcc.D3cfgr.SetD3ppre(rcc.RegisterD3cfgrFieldD3ppreEnumDiv2)
	case Div4:
		rcc.Rcc.D3cfgr.SetD3ppre(rcc.RegisterD3cfgrFieldD3ppreEnumDiv4)
	case Div8:
		rcc.Rcc.D3cfgr.SetD3ppre(rcc.RegisterD3cfgrFieldD3ppreEnumDiv8)
	case Div16:
		rcc.Rcc.D3cfgr.SetD3ppre(rcc.RegisterD3cfgrFieldD3ppreEnumDiv16)
	default:
		panic("invalid D3PPRE divider")
	}
}
