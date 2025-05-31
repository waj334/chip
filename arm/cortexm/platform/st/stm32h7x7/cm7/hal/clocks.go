//go:build stm32h7x7

package hal

import (
	cortexm "pkg.si-go.dev/chip/arm/cortexm/runtime"

	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/flash"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/pwr"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/syscfg"
)

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

var (
	Divn1Prescaler = 60
	Divn2Prescaler = 16
	Divn3Prescaler = 16

	Divm1FrequencyHz uint64 = 16_000_000
	Divm2FrequencyHz uint64 = 16_000_000
	Divm3FrequencyHz uint64 = 16_000_000

	Divp1FrequencyHz uint64 = 480_000_000
	Divq1FrequencyHz uint64 = 480_000_000
	Divr1FrequencyHz uint64 = 480_000_000

	Divp2FrequencyHz uint64 = 128_000_000
	Divq2FrequencyHz uint64 = 128_000_000
	Divr2FrequencyHz uint64 = 128_000_000

	Divp3FrequencyHz uint64 = 128_000_000
	Divq3FrequencyHz uint64 = 128_000_000
	Divr3FrequencyHz uint64 = 128_000_000

	D1cpre    = Div1
	Hpre      = Div2
	D1ppre    = Div2
	D2ppre1   = Div2
	D2ppre2   = Div2
	D3ppre    = Div2
	DivRtcHse = Div2

	cpu1FrequencyHz = Divp1FrequencyHz / uint64(D1cpre)
	HpreFrequencyHz = Divp1FrequencyHz / uint64(Hpre)

	Hclk1FrequencyHz = HpreFrequencyHz
	Pclk1FrequencyHz = HpreFrequencyHz / uint64(D2ppre1)

	Hclk2FrequencyHz = HpreFrequencyHz
	Pclk2FrequencyHz = HpreFrequencyHz / uint64(D2ppre2)

	Hclk3FrequencyHz = HpreFrequencyHz
	Pclk3FrequencyHz = HpreFrequencyHz / uint64(D1ppre)

	Hclk4FrequencyHz = HpreFrequencyHz
	Pclk4FrequencyHz = HpreFrequencyHz / uint64(D3ppre)

	// HseFrequencyHz High-speed external oscillator. 4 MHz - 48 MHz.
	HseFrequencyHz               uint64 = 16_000_000 // Default: 16 MHz
	LseFrequencyHz               uint64 = 32_768
	PllSourceFrequencyHz         uint64 = HsiFrequencyHz
	I2c13SourceFrequencyHz       uint64 = Pclk1FrequencyHz
	I2c4SourceFrequencyHz        uint64 = Pclk4FrequencyHz
	Usart16SourceFrequencyHz     uint64 = Pclk2FrequencyHz
	Usart234578SourceFrequencyHz uint64 = Pclk1FrequencyHz
	RtcSourceFrequencyHz         uint64 = LsiFrequencyHz

	PllSource         = ClockSourceHsi
	I2c13Source       = ClockSourcePclk1
	I2c4Source        = ClockSourcePclk4
	Usart234578Source = ClockSourcePclk1
	Usart16Source     = ClockSourcePclk2
	RtcClockSource    = ClockSourceLsi

	EnableHSE = false
	EnableLSE = false
)

const (
	// HsiFrequencyHz High-speed internal oscillator.
	HsiFrequencyHz = 64_000_000 // 64 MHz
	CsiFrequencyHz = 4_000_000  // 8 MHz
	LsiFrequencyHz = 32_000
)

func ConfigureClocks() {
	state := cortexm.DisableInterrupts()

	// Enable the LDO.
	// pwr.Pwr.Cr3.SetLdoen(false)
	// pwr.Pwr.Cr3.SetLdoen(true)

	// NOTE: VOS0 can only be activated in VOS1 mode.
	if pwr.Pwr.D3cr.GetVos() != pwr.RegisterD3crFieldVosEnumScale1 {
		// Set the LDO voltage scaler to 1.
		pwr.Pwr.D3cr.SetVos(pwr.RegisterD3crFieldVosEnumScale1)
	}

	// Disable the step-down converter.
	// TODO: Determine if this is always safe!
	pwr.Pwr.Cr3.SetSden(false)

	// Enable the SYSCFG clock.
	rcc.Rcc.Apb4enr.SetSyscfgen(true)

	// Enable overdrive. This will implicitly enable VOS0.
	syscfg.Syscfg.Pwrcr.SetOden(true)

	// Wait for the ready bit.
	for !pwr.Pwr.D3cr.GetVosrdy() {
	}

	if EnableHSE {
		// Enable the HSE clock source.
		rcc.Rcc.Cr.SetHseon(true)
		for !rcc.Rcc.Cr.GetHserdy() {
		}
	}

	if EnableLSE {
		// Enable the LSE clock source.
		rcc.Rcc.Bdcr.SetLseon(true)
		for !rcc.Rcc.Bdcr.GetLseon() {
		}
	}

	// Configure RTC clock settings.
	rcc.Rcc.Cfgr.SetRtcpre(uint8(DivRtcHse))

	switch RtcClockSource {
	case ClockSourceNone:
		rcc.Rcc.Bdcr.SetRtcsrc(0)
		RtcSourceFrequencyHz = 0
	case ClockSourceLse:
		rcc.Rcc.Bdcr.SetRtcsrc(1)
		RtcSourceFrequencyHz = LseFrequencyHz
	case ClockSourceLsi:
		rcc.Rcc.Bdcr.SetRtcsrc(2)
		RtcSourceFrequencyHz = LsiFrequencyHz
	case ClockSourceHseRtc:
		rcc.Rcc.Bdcr.SetRtcsrc(3)
		RtcSourceFrequencyHz = HseFrequencyHz / uint64(DivRtcHse)
	}

	var fDivm1 uint64
	var fDivm2 uint64
	var fDivm3 uint64

	switch PllSource {
	case ClockSourceHse:
		// Select the external clock source.
		rcc.Rcc.Pllckselr.SetPllsrc(rcc.RegisterPllckselrFieldPllsrcEnumHse)
		PllSourceFrequencyHz = HseFrequencyHz

		fDivm1 = HseFrequencyHz / Divm1FrequencyHz
		fDivm2 = HseFrequencyHz / Divm2FrequencyHz
		fDivm3 = HseFrequencyHz / Divm3FrequencyHz

		rcc.Rcc.Pllckselr.SetDivm1(uint8(fDivm1))
		rcc.Rcc.Pllckselr.SetDivm2(uint8(fDivm2))
		rcc.Rcc.Pllckselr.SetDivm3(uint8(fDivm3))
	case ClockSourceHsi:
		// Select the internal clock source.
		rcc.Rcc.Pllckselr.SetPllsrc(rcc.RegisterPllckselrFieldPllsrcEnumHsi)
		PllSourceFrequencyHz = HsiFrequencyHz

		fDivm1 = HsiFrequencyHz / Divm1FrequencyHz
		fDivm2 = HsiFrequencyHz / Divm2FrequencyHz
		fDivm3 = HsiFrequencyHz / Divm3FrequencyHz

		rcc.Rcc.Pllckselr.SetDivm1(uint8(fDivm1))
		rcc.Rcc.Pllckselr.SetDivm2(uint8(fDivm2))
		rcc.Rcc.Pllckselr.SetDivm3(uint8(fDivm3))
	case ClockSourceCsi:
		// Select the CSI clock source.
		rcc.Rcc.Pllckselr.SetPllsrc(rcc.RegisterPllckselrFieldPllsrcEnumCsi)
		PllSourceFrequencyHz = CsiFrequencyHz

		fDivm1 = CsiFrequencyHz / Divm1FrequencyHz
		fDivm2 = CsiFrequencyHz / Divm2FrequencyHz
		fDivm3 = CsiFrequencyHz / Divm3FrequencyHz

		rcc.Rcc.Pllckselr.SetDivm1(uint8(fDivm1))
		rcc.Rcc.Pllckselr.SetDivm2(uint8(fDivm2))
		rcc.Rcc.Pllckselr.SetDivm3(uint8(fDivm3))
	}

	fDivn1 := uint64(Divn1Prescaler) * (PllSourceFrequencyHz / fDivm1)
	fDivn2 := uint64(Divn2Prescaler) * (PllSourceFrequencyHz / fDivm2)
	fDivn3 := uint64(Divn3Prescaler) * (PllSourceFrequencyHz / fDivm3)

	divp1 := max(1, (fDivn1/Divp1FrequencyHz)-1)
	divq1 := max(1, (fDivn1/Divq1FrequencyHz)-1)
	divr1 := max(1, (fDivn1/Divr1FrequencyHz)-1)

	divp2 := max(1, (fDivn2/Divp2FrequencyHz)-1)
	divq2 := max(1, (fDivn2/Divq2FrequencyHz)-1)
	divr2 := max(1, (fDivn2/Divr2FrequencyHz)-1)

	divp3 := max(1, (fDivn3/Divp3FrequencyHz)-1)
	divq3 := max(1, (fDivn3/Divq3FrequencyHz)-1)
	divr3 := max(1, (fDivn3/Divr3FrequencyHz)-1)

	switch {
	case PllSourceFrequencyHz < 2_000_000:
		rcc.Rcc.Pllcfgr.SetPll1rge(rcc.RegisterPllcfgrFieldPll1rgeEnum1to2mhz)
		rcc.Rcc.Pllcfgr.SetPll2rge(rcc.RegisterPllcfgrFieldPll2rgeEnum1to2mhz)
		rcc.Rcc.Pllcfgr.SetPll3rge(rcc.RegisterPllcfgrFieldPll3rgeEnum1to2mhz)
	case PllSourceFrequencyHz < 4_000_000:
		rcc.Rcc.Pllcfgr.SetPll1rge(rcc.RegisterPllcfgrFieldPll1rgeEnum2to4mhz)
		rcc.Rcc.Pllcfgr.SetPll2rge(rcc.RegisterPllcfgrFieldPll2rgeEnum2to4mhz)
		rcc.Rcc.Pllcfgr.SetPll3rge(rcc.RegisterPllcfgrFieldPll3rgeEnum2to4mhz)
	case PllSourceFrequencyHz < 8_000_000:
		rcc.Rcc.Pllcfgr.SetPll1rge(rcc.RegisterPllcfgrFieldPll1rgeEnum4to8mhz)
		rcc.Rcc.Pllcfgr.SetPll2rge(rcc.RegisterPllcfgrFieldPll2rgeEnum4to8mhz)
		rcc.Rcc.Pllcfgr.SetPll3rge(rcc.RegisterPllcfgrFieldPll3rgeEnum4to8mhz)
	default:
		rcc.Rcc.Pllcfgr.SetPll1rge(rcc.RegisterPllcfgrFieldPll1rgeEnum8to16mhz)
		rcc.Rcc.Pllcfgr.SetPll2rge(rcc.RegisterPllcfgrFieldPll2rgeEnum8to16mhz)
		rcc.Rcc.Pllcfgr.SetPll3rge(rcc.RegisterPllcfgrFieldPll3rgeEnum8to16mhz)
	}

	// Configure PLL1 clock source.
	rcc.Rcc.Pll1divr.SetDivn1(uint16(Divn1Prescaler - 1))
	rcc.Rcc.Pll1divr.SetDivp1(uint8(divp1))
	rcc.Rcc.Pll1divr.SetDivq1(uint8(divq1))
	rcc.Rcc.Pll1divr.SetDivr1(uint8(divr1))

	// Disable fractional latch.
	rcc.Rcc.Pllcfgr.SetPll1fracen(false)
	rcc.Rcc.Pll1fracr.SetFracn1(0)

	// Select wide VCO range.
	rcc.Rcc.Pllcfgr.SetPll1vcosel(false)

	// Enable divider outputs.
	rcc.Rcc.Pllcfgr.SetDivp1en(true)
	rcc.Rcc.Pllcfgr.SetDivq1en(true)
	rcc.Rcc.Pllcfgr.SetDivr1en(true)

	// Enable PLL1.
	rcc.Rcc.Pllcfgr.SetPll1fracen(true)
	rcc.Rcc.Cr.SetPll1on(true)
	for !rcc.Rcc.Cr.GetPll1rdy() {
	}

	// Configure PLL2 clock source.
	rcc.Rcc.Pll2divr.SetDivn2(uint16(Divn2Prescaler - 1))
	rcc.Rcc.Pll2divr.SetDivp2(uint8(divp2))
	rcc.Rcc.Pll2divr.SetDivq2(uint8(divq2))
	rcc.Rcc.Pll2divr.SetDivr2(uint8(divr2))
	rcc.Rcc.Pllcfgr.SetPll2fracen(false)
	rcc.Rcc.Pll2fracr.SetFracn2(0)
	rcc.Rcc.Pllcfgr.SetPll2vcosel(true)
	rcc.Rcc.Pllcfgr.SetDivp2en(true)
	rcc.Rcc.Pllcfgr.SetDivq2en(true)
	rcc.Rcc.Pllcfgr.SetDivr2en(true)

	// Enable PLL2.
	rcc.Rcc.Pllcfgr.SetPll2fracen(true)
	rcc.Rcc.Cr.SetPll2on(true)
	for !rcc.Rcc.Cr.GetPll2rdy() {
	}

	// Configure PLL3 clock source.
	rcc.Rcc.Pll3divr.SetDivn3(uint16(Divn3Prescaler - 1))
	rcc.Rcc.Pll3divr.SetDivp3(uint8(divp3))
	rcc.Rcc.Pll3divr.SetDivq3(uint8(divq3))
	rcc.Rcc.Pll3divr.SetDivr3(uint8(divr3))
	rcc.Rcc.Pllcfgr.SetPll3fracen(false)
	rcc.Rcc.Pll3fracr.SetFracn3(0)
	rcc.Rcc.Pllcfgr.SetPll3vcosel(true)
	rcc.Rcc.Pllcfgr.SetDivp3en(true)
	rcc.Rcc.Pllcfgr.SetDivq3en(true)
	rcc.Rcc.Pllcfgr.SetDivr3en(true)

	// Enable PLL3.
	rcc.Rcc.Pllcfgr.SetPll3fracen(true)
	rcc.Rcc.Cr.SetPll3on(true)
	for !rcc.Rcc.Cr.GetPll3rdy() {
	}

	// Set D1 domain Core prescaler.
	switch D1cpre {
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
		panic("invalid divider value")
	}

	// Set D1 domain AHB prescaler.
	switch Hpre {
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
		panic("invalid divider value")
	}

	// Set D1 domain APB3 prescaler.
	switch D1ppre {
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
		panic("invalid divider value")
	}

	// Set D2 domain APB1 prescaler.
	switch D2ppre1 {
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
		panic("invalid divider value")
	}

	// Set D2 domain APB2 prescaler.
	switch D2ppre2 {
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
		panic("invalid divider value")
	}

	// Set D3 domain APB4 prescaler.
	switch D3ppre {
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
		panic("invalid divider value")
	}

	// Wait for PLL1 to be ready
	for !rcc.Rcc.Cr.GetPll1rdy() {
	}

	// Set the System Clock source to PLL1.
	rcc.Rcc.Cfgr.SetSw(rcc.RegisterCfgrFieldSwEnumPll1)
	for rcc.Rcc.Cfgr.GetSws() != rcc.RegisterCfgrFieldSwsEnumPll1 {
	}

	// Set flash wait state for 480Mhz.
	flash.Flash.Bank[0].Acr.SetLatency(4) // Use 4 wait states.
	for flash.Flash.Bank[0].Acr.GetLatency() != 4 {
	}

	flash.Flash.Bank[0].Acr.SetWrhighfreq(3) // Delay.
	for flash.Flash.Bank[0].Acr.GetWrhighfreq() != 3 {
	}

	// Update clock frequency values.
	cpu1FrequencyHz = Divp1FrequencyHz / uint64(D1cpre)
	HpreFrequencyHz = Divp1FrequencyHz / uint64(Hpre)

	Hclk1FrequencyHz = HpreFrequencyHz
	Pclk1FrequencyHz = HpreFrequencyHz / uint64(D2ppre1)

	Hclk2FrequencyHz = HpreFrequencyHz
	Pclk2FrequencyHz = HpreFrequencyHz / uint64(D2ppre2)

	Hclk3FrequencyHz = HpreFrequencyHz
	Pclk3FrequencyHz = HpreFrequencyHz / uint64(D1ppre)

	Hclk4FrequencyHz = HpreFrequencyHz
	Pclk4FrequencyHz = HpreFrequencyHz / uint64(D3ppre)

	// Enable GPIO clocks.
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

	// Enable DAC clock.
	rcc.Rcc.Apb1lenr.SetDac12en(true)

	// Enable ADC clocks.
	rcc.Rcc.D3ccipr.SetAdcsrc(0b00) // Select PLL2.
	rcc.Rcc.Ahb1enr.SetAdc12en(true)
	rcc.Rcc.Ahb4enr.SetAdc3en(true)

	// Handle I2C1 - I2C3
	if I2c13Source != ClockSourceNone {
		// Set the clock source.
		switch I2c13Source {
		case ClockSourcePclk1:
			rcc.Rcc.D2ccip2r.SetI2c123src(0)
			I2c13SourceFrequencyHz = Pclk1FrequencyHz
		case ClockSourcePll3r:
			rcc.Rcc.D2ccip2r.SetI2c123src(1)
			I2c13SourceFrequencyHz = Divr3FrequencyHz
		case ClockSourceHsi:
			rcc.Rcc.D2ccip2r.SetI2c123src(2)
			I2c13SourceFrequencyHz = HsiFrequencyHz
		case ClockSourceCsi:
			rcc.Rcc.D2ccip2r.SetI2c123src(3)
			I2c13SourceFrequencyHz = CsiFrequencyHz
		default:
			panic("invalid I2C1_3 clock source")
		}
	}

	// Handle I2C4
	if I2c4Source != ClockSourceNone {
		// Set the clock source.
		switch I2c4Source {
		case ClockSourcePclk4:
			rcc.Rcc.D3ccipr.SetI2c4src(0)
			I2c4SourceFrequencyHz = Pclk4FrequencyHz
		case ClockSourcePll3r:
			rcc.Rcc.D3ccipr.SetI2c4src(1)
			I2c4SourceFrequencyHz = Divr3FrequencyHz
		case ClockSourceHsi:
			rcc.Rcc.D3ccipr.SetI2c4src(2)
			I2c4SourceFrequencyHz = HsiFrequencyHz
		case ClockSourceCsi:
			rcc.Rcc.D3ccipr.SetI2c4src(3)
			I2c4SourceFrequencyHz = CsiFrequencyHz
		default:
			panic("invalid I2C4 clock source")
		}
	}

	if Usart16Source != ClockSourceNone {
		switch Usart16Source {
		case ClockSourcePclk2:
			rcc.Rcc.D2ccip2r.SetUsart16src(0)
			Usart16SourceFrequencyHz = Pclk1FrequencyHz
		case ClockSourcePll2q:
			rcc.Rcc.D2ccip2r.SetUsart16src(1)
			Usart16SourceFrequencyHz = Divq2FrequencyHz
		case ClockSourcePll3q:
			rcc.Rcc.D2ccip2r.SetUsart16src(2)
			Usart16SourceFrequencyHz = Divq3FrequencyHz
		case ClockSourceHsi:
			rcc.Rcc.D2ccip2r.SetUsart16src(3)
			Usart16SourceFrequencyHz = HsiFrequencyHz
		case ClockSourceCsi:
			rcc.Rcc.D2ccip2r.SetUsart16src(4)
			Usart16SourceFrequencyHz = CsiFrequencyHz
		case ClockSourceLse:
			rcc.Rcc.D2ccip2r.SetUsart16src(4)
			Usart16SourceFrequencyHz = HseFrequencyHz
		}
	}

	if Usart234578Source != ClockSourceNone {
		switch Usart16Source {
		case ClockSourcePclk1:
			rcc.Rcc.D2ccip2r.SetUsart234578src(0)
			Usart234578SourceFrequencyHz = Pclk1FrequencyHz
		case ClockSourcePll2q:
			rcc.Rcc.D2ccip2r.SetUsart234578src(1)
			Usart234578SourceFrequencyHz = Divq2FrequencyHz
		case ClockSourcePll3q:
			rcc.Rcc.D2ccip2r.SetUsart234578src(2)
			Usart234578SourceFrequencyHz = Divq3FrequencyHz
		case ClockSourceHsi:
			rcc.Rcc.D2ccip2r.SetUsart234578src(3)
			Usart234578SourceFrequencyHz = HsiFrequencyHz
		case ClockSourceCsi:
			rcc.Rcc.D2ccip2r.SetUsart234578src(4)
			Usart234578SourceFrequencyHz = CsiFrequencyHz
		case ClockSourceLse:
			rcc.Rcc.D2ccip2r.SetUsart234578src(4)
			Usart234578SourceFrequencyHz = HseFrequencyHz
		}
	}

	// Update the SysTick frequency to match the CPU clock frequency.
	cortexm.UpdateSysTickFrequency(uint32(Divp1FrequencyHz))
	cortexm.EnableInterrupts(state)
}
