package dsihost

import (
	"pkg.si-go.dev/chip/arm/cortexm"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	regdsi "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/dsihost"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	corehal "pkg.si-go.dev/chip/core/hal"
)

func validateConfig(cfg Config) error {
	if cfg.Lanes != OneDataLane && cfg.Lanes != TwoDataLanes {
		return corehal.ErrInvalidConfig
	}

	if cfg.VirtualChannelID > 3 {
		return corehal.ErrInvalidConfig
	}

	switch cfg.ColorCoding {
	case ColorRGB565, ColorRGB666, ColorRGB888:
	default:
		return corehal.ErrInvalidConfig
	}

	switch cfg.VideoMode {
	case VideoModeNonBurstSyncPulse, VideoModeNonBurstSyncEvent, VideoModeBurst:
	default:
		return corehal.ErrInvalidConfig
	}

	if cfg.TXEscapeCkdiv < 2 {
		return corehal.ErrInvalidConfig
	}

	if cfg.PacketSize == 0 ||
		cfg.HorizontalLine == 0 ||
		cfg.VerticalActive == 0 {
		return corehal.ErrInvalidConfig
	}

	if cfg.HorizontalSyncActive > 0x0fff ||
		cfg.HorizontalBackPorch > 0x0fff ||
		cfg.HorizontalLine > 0x7fff ||
		cfg.VerticalSyncActive > 0x3fff ||
		cfg.VerticalBackPorch > 0x3fff ||
		cfg.VerticalFrontPorch > 0x3fff ||
		cfg.VerticalActive > 0x3fff ||
		cfg.PacketSize > 0x3fff ||
		cfg.NumberOfChunks > 0x1fff ||
		cfg.NullPacketSize > 0x1fff {
		return corehal.ErrInvalidConfig
	}

	if !cfg.PLL.Enable {
		return corehal.ErrInvalidConfig
	}

	if cfg.PLL.NDIV < 10 || cfg.PLL.NDIV > 125 {
		return corehal.ErrInvalidConfig
	}
	if cfg.PLL.IDF < PLLIDF1 || cfg.PLL.IDF > PLLIDF7 {
		return corehal.ErrInvalidConfig
	}
	if cfg.PLL.ODF > PLLODF8 {
		return corehal.ErrInvalidConfig
	}

	return nil
}

func enableClock() {
	rcc.Rcc.Apb3enr.SetDsien(true)
	for !rcc.Rcc.Apb3enr.GetDsien() {
	}

	// On dual-core H7, also enable the CPU1 access gate.
	rcc.Rcc.C1apb3enr.SetDsien(true)
	for !rcc.Rcc.C1apb3enr.GetDsien() {
	}

	cortexm.DSB()
}

func resetPeripheral() {
	rcc.Rcc.Apb3rstr.SetDsirst(true)
	for !rcc.Rcc.Apb3rstr.GetDsirst() {
	}

	rcc.Rcc.Apb3rstr.SetDsirst(false)
	for rcc.Rcc.Apb3rstr.GetDsirst() {
	}

	cortexm.DSB()
}

// configurePLL turns on the wrapper regulator, waits for it, then programs and
// enables the wrapper PLL, finally waiting for lock. Matches the PLL section
// of ST HAL_DSI_Init.
func configurePLL(pll PLLConfig) error {
	if !pll.Enable {
		regdsi.Dsihost.Dsiwrpcr.SetPllen(false)
		return corehal.ErrInvalidConfig
	}

	// Ensure PLL is off before reprogramming.
	regdsi.Dsihost.Dsiwrpcr.SetPllen(false)

	// Regulator enable + wait for RRS.
	regdsi.Dsihost.Dsiwrpcr.SetRegen(pll.RegulatorEnable)
	if pll.RegulatorEnable {
		if err := waitRegulatorReady(); err != nil {
			return err
		}
	}

	// Program dividers.
	regdsi.Dsihost.Dsiwrpcr.SetIdf(uint8(pll.IDF))
	regdsi.Dsihost.Dsiwrpcr.SetNdiv(pll.NDIV)
	regdsi.Dsihost.Dsiwrpcr.SetOdf(uint8(pll.ODF))

	// Enable PLL.
	regdsi.Dsihost.Dsiwrpcr.SetPllen(true)
	cortexm.DSB()

	// Wait for lock. ST HAL inserts a 1 ms delay before checking PLLLS; our
	// polling loop's own overhead far exceeds the minimum 400 us settling time,
	// so checking immediately is safe - the bit will just read 0 until lock.
	return waitWrapperPLLLock()
}

// configureCommandMode programs DSI_CMCR. Each *TX bit controls whether
// packets of that type are transmitted in LP (bit=1) or HS (bit=0).
// For panel init we want LP for everything.
func configureCommandMode(cfg Config) {
	lp := cfg.LPCommandMode
	var cmcr regdsi.RegisterDsicmcrType

	// Generic short writes.
	cmcr.SetGsw0tx(lp)
	cmcr.SetGsw1tx(lp)
	cmcr.SetGsw2tx(lp)

	// Generic long write.
	cmcr.SetGlwtx(lp)

	// DCS short writes.
	cmcr.SetDsw0tx(lp)
	cmcr.SetDsw1tx(lp)

	// DCS long write.
	cmcr.SetDlwtx(lp)

	// Reads in LP too during bring-up.
	cmcr.SetGsr0tx(lp)
	cmcr.SetGsr1tx(lp)
	cmcr.SetGsr2tx(lp)
	cmcr.SetDsr0tx(lp)

	cmcr.SetMrdps(lp)

	regdsi.Dsihost.Dsicmcr.Store(uint32(cmcr))
}

// configureVideoMode mirrors ST HAL_DSI_ConfigVideoMode in order and content.
func configureVideoMode(cfg Config) {
	// Video mode (not adapted-command) on both host and wrapper.
	regdsi.Dsihost.Dsimcr.SetCmdm(false)
	regdsi.Dsihost.Dsiwcfgr.SetDsim(false)

	// Video transmission type (burst, sync-pulse, sync-event).
	regdsi.Dsihost.Dsivmcr.SetVmt(uint8(cfg.VideoMode))

	// Video packet size, chunks, null packet size.
	regdsi.Dsihost.Dsivpcr.SetVpsize(cfg.PacketSize)
	regdsi.Dsihost.Dsivccr.SetNumc(cfg.NumberOfChunks)
	regdsi.Dsihost.Dsivnpcr.SetNpsize(cfg.NullPacketSize)

	// Virtual channel for LTDC interface traffic.
	regdsi.Dsihost.Dsilvcidr.SetVcid(cfg.VirtualChannelID)

	// Control signal polarities at the DSI<->panel interface.
	regdsi.Dsihost.Dsilpcr.SetDep(cfg.DEPolarityHigh)
	regdsi.Dsihost.Dsilpcr.SetVsp(cfg.VSPolarityHigh)
	regdsi.Dsihost.Dsilpcr.SetHsp(cfg.HSPolarityHigh)

	// Color coding for host (LCOLCR) and wrapper (WCFGR.COLMUX).
	regdsi.Dsihost.Dsilcolcr.SetColc(uint8(cfg.ColorCoding))
	regdsi.Dsihost.Dsiwcfgr.SetColmux(uint8(cfg.ColorMux))

	// LooselyPacked only meaningful for RGB666.
	if cfg.ColorCoding == ColorRGB666 {
		regdsi.Dsihost.Dsilcolcr.SetLpe(cfg.LooselyPacked)
	}

	// Horizontal timing (lane-byte clock cycles).
	regdsi.Dsihost.Dsivhsacr.SetHsa(cfg.HorizontalSyncActive)
	regdsi.Dsihost.Dsivhbpcr.SetHbp(cfg.HorizontalBackPorch)
	regdsi.Dsihost.Dsivlcr.SetHline(cfg.HorizontalLine)

	// Vertical timing (lines).
	regdsi.Dsihost.Dsivvsacr.SetVsa(cfg.VerticalSyncActive)
	regdsi.Dsihost.Dsivvbpcr.SetVbp(cfg.VerticalBackPorch)
	regdsi.Dsihost.Dsivvfpcr.SetVfp(cfg.VerticalFrontPorch)
	regdsi.Dsihost.Dsivvacr.SetVa(cfg.VerticalActive)

	// LP command enable while video streaming.
	regdsi.Dsihost.Dsivmcr.SetLpce(cfg.LPCommandEnable)

	// LP largest packet sizes.
	regdsi.Dsihost.Dsilpmcr.SetLpsize(cfg.LPLargestPacketSize)
	regdsi.Dsihost.Dsilpmcr.SetVlpsize(cfg.LPVACTLargestPacketSize)

	// LP transition enable per region.
	regdsi.Dsihost.Dsivmcr.SetLphfpe(cfg.LPHFPEnable)
	regdsi.Dsihost.Dsivmcr.SetLphbpe(cfg.LPHBPEnable)
	regdsi.Dsihost.Dsivmcr.SetLpvae(cfg.LPVACTEnable)
	regdsi.Dsihost.Dsivmcr.SetLpvfpe(cfg.LPVFPEnable)
	regdsi.Dsihost.Dsivmcr.SetLpvbpe(cfg.LPVBPEnable)
	regdsi.Dsihost.Dsivmcr.SetLpvsae(cfg.LPVSAEnable)

	// End-of-frame BTA acknowledge.
	regdsi.Dsihost.Dsivmcr.SetFbtaae(cfg.FrameBTAAcknowledge)

	// Pattern generator OFF for normal use.
	regdsi.Dsihost.Dsivmcr.SetPge(false)
	regdsi.Dsihost.Dsivmcr.SetPgm(false)
	regdsi.Dsihost.Dsivmcr.SetPgo(false)
}

// configureClocking sets TX escape clock divisor.
func configureClocking(cfg Config) {
	regdsi.Dsihost.Dsiccr.SetTxeckdiv(cfg.TXEscapeCkdiv)
	regdsi.Dsihost.Dsiccr.SetTockdiv(0)
}

// configurePHYTimer mirrors ST HAL_DSI_ConfigPhyTimer. The clock-lane LP2HS
// and HS2LP times are forced to max(LP2HS, HS2LP) on both fields because the
// hardware computes 2*HS2LP instead of LP2HS+HS2LP (ST workaround).
func configurePHYTimer(cfg Config) {
	if !cfg.PhyTiming.Enable {
		return
	}

	pt := cfg.PhyTiming

	// Clock lane timer: workaround for 2*HS2LP hardware quirk.
	maxTime := pt.ClockLP2HS
	if pt.ClockHS2LP > maxTime {
		maxTime = pt.ClockHS2LP
	}
	regdsi.Dsihost.Dsicltcr.SetLp2hstime(maxTime)
	regdsi.Dsihost.Dsicltcr.SetHs2lptime(maxTime)

	// Data lane timer.
	regdsi.Dsihost.Dsidltcr.SetLp2hstime(pt.DataLP2HS)
	regdsi.Dsihost.Dsidltcr.SetHs2lptime(pt.DataHS2LP)
	regdsi.Dsihost.Dsidltcr.SetMrdtime(pt.MaxReadTime)
}

// configureProtocol turns off flow-control / error-response features for
// first-light. Matches what Arduino doesn't touch (defaults remain 0).
func configureProtocol(cfg Config) {
	regdsi.Dsihost.Dsipcr.SetEttxe(false)
	regdsi.Dsihost.Dsipcr.SetEtrxe(false)
	regdsi.Dsihost.Dsipcr.SetBtae(false)
	regdsi.Dsihost.Dsipcr.SetEccrxe(false)
	regdsi.Dsihost.Dsipcr.SetCrcrxe(false)

	regdsi.Dsihost.Dsigvcidr.SetVcid(cfg.VirtualChannelID)
}

// setUIX4 computes and writes the HS bit period to DSI_WPCR0.UIX4. This is
// mandatory for the D-PHY to time HS transmissions correctly.
//
//	UIX4 = (4_000_000 * IDF * 2^ODF) / ((HSE_MHz * 1000) * NDIV)
//
// For HSE=16 MHz, NDIV=125, IDF=4, ODF=0: UIX4 = 8 (i.e. 2 ns bit period).
//
// NOTE: WPCR0 register naming in the generated regdsi package may vary. If
// the field setter name differs, the read-modify-write below works as long
// as the register exists as a uint32 mapped register with Load/Store.
func setUIX4(pll PLLConfig) error {
	if pll.NDIV == 0 || pll.IDF == 0 {
		return corehal.ErrInvalidConfig
	}

	// 2^ODF: ODF encodes 0=/1, 1=/2, 2=/4, 3=/8.
	odfMul := uint64(1) << uint64(pll.ODF&0x3)

	uix4 := (uint64(4_000_000) * uint64(pll.IDF) * odfMul) /
		((uint64(hal.HseFrequencyHz) / 1000) * uint64(pll.NDIV))

	// UIX4 occupies bits[5:0] of WPCR0.
	const uix4Mask = uint32(0x3F)
	v := regdsi.Dsihost.Dsiwpcr0.Load() &^ uix4Mask
	v |= uint32(uix4) & uix4Mask
	regdsi.Dsihost.Dsiwpcr0.Store(v)
	cortexm.DSB()

	return nil
}

// clearInterruptEnables disables all DSI host error interrupts (IER0, IER1).
// Matches ST HAL_DSI_Init.
func clearInterruptEnables() {
	regdsi.Dsihost.Dsiier0.Store(0)
	regdsi.Dsihost.Dsiier1.Store(0)
}

func dsiCycles(pixelClocks uint32) uint16 {
	return uint16(pixelClocks * 62_500_000 / uint32(hal.LtdcSourceFrequencyHz))
}
