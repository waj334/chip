package ltdc

import (
	"unsafe"

	"pkg.si-go.dev/chip/arm/cortexm"
	regltdc "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/ltdc"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	corehal "pkg.si-go.dev/chip/core/hal"
)

// Package ltdc provides a thin HAL for STM32H7 LTDC framebuffer scanout.
//
// Framebuffer memory must be visible to LTDC over the bus matrix. Do not place
// LTDC framebuffers in DTCM. Use AXI SRAM, SRAM1/2/3, or external SDRAM/FMC.
//
// If the framebuffer region is cacheable, callers must clean D-cache for dirty
// regions before reload / scanout. For the first bring-up, prefer a
// non-cacheable MPU region.
//
// LTDC ownership is assumed to be single-core. If both M7 and M4 can touch
// LTDC registers, guard access externally, for example, with HSEM.

type PixelFormat uint8

const (
	// These values map directly to LTDC_LxPFCR.PF.
	PixelFormatARGB8888 PixelFormat = 0b000
	PixelFormatRGB888   PixelFormat = 0b001
	PixelFormatRGB565   PixelFormat = 0b010
	PixelFormatARGB1555 PixelFormat = 0b011
	PixelFormatARGB4444 PixelFormat = 0b100
	PixelFormatL8       PixelFormat = 0b101
	PixelFormatAL44     PixelFormat = 0b110
	PixelFormatAL88     PixelFormat = 0b111
)

type Layer uint8

const (
	Layer1 Layer = 0
	Layer2 Layer = 1
)

type Timing struct {
	Width  uint16
	Height uint16

	HSYNC uint16
	HBP   uint16
	HFP   uint16

	VSYNC uint16
	VBP   uint16
	VFP   uint16
}

type LayerConfig struct {
	Enabled bool

	Framebuffer unsafe.Pointer

	Width  uint16
	Height uint16

	// Bytes per line in memory.
	// For RGB888 480 wide, this is 480 * 3 = 1440.
	Pitch uint16

	Format PixelFormat

	// Window position inside the active display area.
	// X1/Y1 are exclusive. If X1/Y1 are zero, they default to X0+Width/Y0+Height.
	X0 uint16
	Y0 uint16
	X1 uint16
	Y1 uint16

	// 0 = transparent, 255 = opaque.
	Alpha uint8
}

type Config struct {
	Enable bool

	Timing Timing

	BackgroundRed   uint8
	BackgroundGreen uint8
	BackgroundBlue  uint8

	// Polarity bits are written directly into LTDC_GCR.
	// For internal LTDC<->DSI bridges, note that DE polarity is INVERTED
	// between LTDC and DSI sides (see ST stm32h7xx_hal_ltdc_ex.c).
	HSPolHigh          bool
	VSPolHigh          bool
	DEPolHigh          bool
	PixelClockInverted bool

	Layers [2]LayerConfig
}

type LTDC struct{}

const (
	maxHSW    = 0x03ff
	maxHAccum = 0x0fff
	maxVAccum = 0x07ff

	// LTDC blending factor field values.
	// BF1 = PAxCA, BF2 = 1 - PAxCA.
	blendingFactor1PAxCA = 0x6
	blendingFactor2PAxCA = 0x7
)

type timingRegs struct {
	hsw    uint32
	vsh    uint32
	ahbp   uint32
	avbp   uint32
	aaw    uint32
	aah    uint32
	totalw uint32
	totalh uint32
}

type preparedConfig struct {
	timing timingRegs

	layerEnabled [2]bool
	layers       [2]LayerConfig
}

func (l LTDC) Configure(cfg Config) error {
	if !cfg.Enable {
		l.Disable()
		return nil
	}

	prepared, err := prepareConfig(cfg)
	if err != nil {
		return err
	}

	enableClock()
	resetPeripheral()

	disableLayer(Layer1)
	disableLayer(Layer2)

	configurePolarity(cfg)
	configureGlobalTiming(prepared.timing)
	configureBackground(cfg)

	if prepared.layerEnabled[0] {
		configureLayer(Layer1, cfg.Timing, prepared.layers[0])
	}
	if prepared.layerEnabled[1] {
		configureLayer(Layer2, cfg.Timing, prepared.layers[1])
	}

	clearInterruptFlags()
	reloadImmediate()

	l.Enable()
	return nil
}

func (l LTDC) Enable() {
	enableClock()
	regltdc.Ltdc.Gcr.SetLtdcen(true)
}

func (l LTDC) Disable() {
	enableClock()

	disableLayer(Layer1)
	disableLayer(Layer2)
	clearInterruptFlags()
	reloadImmediate()

	regltdc.Ltdc.Gcr.SetLtdcen(false)
}

func (l LTDC) Enabled() bool {
	return regltdc.Ltdc.Gcr.GetLtdcen()
}

func (l LTDC) Reload() {
	reloadImmediate()
}

func (l LTDC) ReloadVerticalBlank() {
	reloadVerticalBlank()
}

// ConfigureLayer programs one LTDC layer without resetting the LTDC peripheral.
// This mirrors HAL_LTDC_ConfigLayer and is useful for the DSI bring-up order
// where LTDC global timing is enabled before layer state is installed.
func (l LTDC) ConfigureLayer(layer Layer, t Timing, cfg LayerConfig, verticalBlank bool) error {
	switch layer {
	case Layer1, Layer2:
	default:
		return corehal.ErrInvalidConfig
	}

	normalized, err := normalizeLayer(cfg)
	if err != nil {
		return err
	}
	if err := validateLayer(t, normalized); err != nil {
		return err
	}

	configureLayer(layer, t, normalized)
	if verticalBlank {
		reloadVerticalBlank()
	} else {
		reloadImmediate()
	}
	return nil
}

func (l LTDC) EnableLayer(layer Layer, verticalBlank bool) error {
	switch layer {
	case Layer1:
		regltdc.Ltdc.L1cr.SetLen(true)
	case Layer2:
		regltdc.Ltdc.L2cr.SetLen(true)
	default:
		return corehal.ErrInvalidConfig
	}

	if verticalBlank {
		reloadVerticalBlank()
	} else {
		reloadImmediate()
	}
	return nil
}

func (l LTDC) SetLayerFramebuffer(layer Layer, framebuffer unsafe.Pointer, verticalBlank bool) error {
	if framebuffer == nil {
		return corehal.ErrInvalidConfig
	}

	switch layer {
	case Layer1:
		regltdc.Ltdc.L1cfbar.Store(uint32(uintptr(framebuffer)))
	case Layer2:
		regltdc.Ltdc.L2cfbar.Store(uint32(uintptr(framebuffer)))
	default:
		return corehal.ErrInvalidConfig
	}

	if verticalBlank {
		reloadVerticalBlank()
	} else {
		reloadImmediate()
	}
	return nil
}

func (l LTDC) DisableLayer(layer Layer, verticalBlank bool) error {
	switch layer {
	case Layer1, Layer2:
		disableLayer(layer)
	default:
		return corehal.ErrInvalidConfig
	}

	if verticalBlank {
		reloadVerticalBlank()
	} else {
		reloadImmediate()
	}
	return nil
}

func (l LTDC) InterruptStatus() uint32 {
	return regltdc.Ltdc.Isr.Load()
}

func (l LTDC) ClearInterrupts() {
	clearInterruptFlags()
}

func (l LTDC) CurrentX() uint16 {
	return regltdc.Ltdc.Cpsr.GetCxpos()
}

func (l LTDC) CurrentY() uint16 {
	return regltdc.Ltdc.Cpsr.GetCypos()
}

func (l LTDC) DisplayStatus() uint32 {
	return regltdc.Ltdc.Cdsr.Load()
}

func prepareConfig(cfg Config) (preparedConfig, error) {
	timing, err := computeTimingRegs(cfg.Timing)
	if err != nil {
		return preparedConfig{}, err
	}

	prepared := preparedConfig{timing: timing}

	for i := 0; i < len(cfg.Layers); i++ {
		layer := cfg.Layers[i]
		if !layer.Enabled {
			continue
		}

		normalized, err := normalizeLayer(layer)
		if err != nil {
			return preparedConfig{}, err
		}

		if err := validateLayer(cfg.Timing, normalized); err != nil {
			return preparedConfig{}, err
		}

		prepared.layerEnabled[i] = true
		prepared.layers[i] = normalized
	}

	return prepared, nil
}

func computeTimingRegs(t Timing) (timingRegs, error) {
	if t.Width == 0 || t.Height == 0 ||
		t.HSYNC == 0 || t.HBP == 0 || t.HFP == 0 ||
		t.VSYNC == 0 || t.VBP == 0 || t.VFP == 0 {
		return timingRegs{}, corehal.ErrInvalidConfig
	}

	hsync := uint32(t.HSYNC)
	hbp := uint32(t.HBP)
	width := uint32(t.Width)
	hfp := uint32(t.HFP)

	vsync := uint32(t.VSYNC)
	vbp := uint32(t.VBP)
	height := uint32(t.Height)
	vfp := uint32(t.VFP)

	regs := timingRegs{
		hsw:    hsync - 1,
		vsh:    vsync - 1,
		ahbp:   hsync + hbp - 1,
		avbp:   vsync + vbp - 1,
		aaw:    hsync + hbp + width - 1,
		aah:    vsync + vbp + height - 1,
		totalw: hsync + hbp + width + hfp - 1,
		totalh: vsync + vbp + height + vfp - 1,
	}

	if regs.hsw > maxHSW || regs.vsh > maxVAccum ||
		regs.ahbp > maxHAccum || regs.avbp > maxVAccum ||
		regs.aaw > maxHAccum || regs.aah > maxVAccum ||
		regs.totalw > maxHAccum || regs.totalh > maxVAccum {
		return timingRegs{}, corehal.ErrInvalidConfig
	}

	return regs, nil
}

func normalizeLayer(layer LayerConfig) (LayerConfig, error) {
	if layer.X1 == 0 {
		x1 := uint32(layer.X0) + uint32(layer.Width)
		if x1 > 0xffff {
			return LayerConfig{}, corehal.ErrInvalidConfig
		}
		layer.X1 = uint16(x1)
	}

	if layer.Y1 == 0 {
		y1 := uint32(layer.Y0) + uint32(layer.Height)
		if y1 > 0xffff {
			return LayerConfig{}, corehal.ErrInvalidConfig
		}
		layer.Y1 = uint16(y1)
	}

	return layer, nil
}

func validateLayer(t Timing, layer LayerConfig) error {
	if layer.Framebuffer == nil {
		return corehal.ErrInvalidConfig
	}
	if layer.Width == 0 || layer.Height == 0 {
		return corehal.ErrInvalidConfig
	}
	if layer.X1 <= layer.X0 || layer.Y1 <= layer.Y0 {
		return corehal.ErrInvalidConfig
	}
	if layer.X1 > t.Width || layer.Y1 > t.Height {
		return corehal.ErrInvalidConfig
	}
	if uint32(layer.X1)-uint32(layer.X0) != uint32(layer.Width) {
		return corehal.ErrInvalidConfig
	}
	if uint32(layer.Y1)-uint32(layer.Y0) != uint32(layer.Height) {
		return corehal.ErrInvalidConfig
	}

	bpp := bytesPerPixel(layer.Format)
	if bpp == 0 {
		return corehal.ErrInvalidConfig
	}

	activeLineBytes := uint32(layer.Width) * uint32(bpp)
	if activeLineBytes == 0 || activeLineBytes > 0x1ffc {
		return corehal.ErrInvalidConfig
	}

	if uint32(layer.Pitch) < activeLineBytes || uint32(layer.Pitch) > 0x1fff {
		return corehal.ErrInvalidConfig
	}

	hsync := uint32(t.HSYNC)
	hbp := uint32(t.HBP)
	vsync := uint32(t.VSYNC)
	vbp := uint32(t.VBP)

	ahbp := hsync + hbp - 1
	avbp := vsync + vbp - 1

	whstpos := uint32(layer.X0) + ahbp + 1
	whsppos := uint32(layer.X1) + ahbp

	wvstpos := uint32(layer.Y0) + avbp + 1
	wvsppos := uint32(layer.Y1) + avbp

	if whstpos > maxHAccum || whsppos > maxHAccum ||
		wvstpos > maxVAccum || wvsppos > maxVAccum {
		return corehal.ErrInvalidConfig
	}

	return nil
}

func bytesPerPixel(format PixelFormat) uint8 {
	switch format {
	case PixelFormatARGB8888:
		return 4
	case PixelFormatRGB888:
		return 3
	case PixelFormatRGB565, PixelFormatARGB1555, PixelFormatARGB4444, PixelFormatAL88:
		return 2
	case PixelFormatL8, PixelFormatAL44:
		return 1
	default:
		return 0
	}
}

func enableClock() {
	rcc.Rcc.Apb3enr.SetLtdcen(true)
	for !rcc.Rcc.Apb3enr.GetLtdcen() {
	}
}

func resetPeripheral() {
	rcc.Rcc.Apb3rstr.SetLtdcrst(true)
	for !rcc.Rcc.Apb3rstr.GetLtdcrst() {
	}

	rcc.Rcc.Apb3rstr.SetLtdcrst(false)
	for rcc.Rcc.Apb3rstr.GetLtdcrst() {
	}
}

func configurePolarity(cfg Config) {
	const polarityMask = regltdc.RegisterGcrFieldPcpolMask |
		regltdc.RegisterGcrFieldDepolMask |
		regltdc.RegisterGcrFieldVspolMask |
		regltdc.RegisterGcrFieldHspolMask

	value := regltdc.Ltdc.Gcr.Load() &^ polarityMask

	if cfg.PixelClockInverted {
		value |= regltdc.RegisterGcrFieldPcpolMask
	}
	if cfg.DEPolHigh {
		value |= regltdc.RegisterGcrFieldDepolMask
	}
	if cfg.VSPolHigh {
		value |= regltdc.RegisterGcrFieldVspolMask
	}
	if cfg.HSPolHigh {
		value |= regltdc.RegisterGcrFieldHspolMask
	}

	regltdc.Ltdc.Gcr.Store(value)
}

func configureGlobalTiming(regs timingRegs) {
	regltdc.Ltdc.Sscr.Store((regs.hsw << 16) | regs.vsh)
	regltdc.Ltdc.Bpcr.Store((regs.ahbp << 16) | regs.avbp)
	regltdc.Ltdc.Awcr.Store((regs.aaw << 16) | regs.aah)
	regltdc.Ltdc.Twcr.Store((regs.totalw << 16) | regs.totalh)
}

func configureBackground(cfg Config) {
	value := uint32(cfg.BackgroundBlue) |
		(uint32(cfg.BackgroundGreen) << 8) |
		(uint32(cfg.BackgroundRed) << 16)

	regltdc.Ltdc.Bccr.Store(value)
}

func configureLayer(layerIndex Layer, t Timing, layer LayerConfig) {
	hsync := uint32(t.HSYNC)
	hbp := uint32(t.HBP)
	vsync := uint32(t.VSYNC)
	vbp := uint32(t.VBP)

	ahbp := hsync + hbp - 1
	avbp := vsync + vbp - 1

	whstpos := uint32(layer.X0) + ahbp + 1
	whsppos := uint32(layer.X1) + ahbp

	wvstpos := uint32(layer.Y0) + avbp + 1
	wvsppos := uint32(layer.Y1) + avbp

	bpp := uint32(bytesPerPixel(layer.Format))
	activeLineBytes := uint32(layer.Width) * bpp

	// CFBLL = active line length in bytes + 7 (matches ST HAL LTDC_SetConfig).
	// RM allows +3 minimum, but +7 is more conservative and matches ST's
	// validated drivers.
	lineLength := activeLineBytes + 3
	pitch := uint32(layer.Pitch)

	whpcr := (whsppos << 16) | whstpos
	wvpcr := (wvsppos << 16) | wvstpos
	pfcr := uint32(layer.Format)
	cacr := uint32(layer.Alpha)
	dccr := uint32(0)
	bfcr := uint32(blendingFactor1PAxCA<<8) | uint32(blendingFactor2PAxCA)
	cfbar := uint32(uintptr(layer.Framebuffer))
	cfblr := (pitch << 16) | lineLength
	cfblnr := uint32(layer.Height)

	switch layerIndex {
	case Layer1:
		regltdc.Ltdc.L1cr.SetLen(false)

		regltdc.Ltdc.L1whpcr.Store(whpcr)
		regltdc.Ltdc.L1wvpcr.Store(wvpcr)
		regltdc.Ltdc.L1ckcr.Store(0)
		regltdc.Ltdc.L1pfcr.Store(pfcr)
		regltdc.Ltdc.L1cacr.Store(cacr)
		regltdc.Ltdc.L1dccr.Store(dccr)
		regltdc.Ltdc.L1bfcr.Store(bfcr)
		regltdc.Ltdc.L1cfbar.Store(cfbar)
		regltdc.Ltdc.L1cfblr.Store(cfblr)
		regltdc.Ltdc.L1cfblnr.Store(cfblnr)

		regltdc.Ltdc.L1cr.SetColken(false)
		regltdc.Ltdc.L1cr.SetCluten(false)
		if layer.Enabled {
			regltdc.Ltdc.L1cr.SetLen(true)
		}

	case Layer2:
		regltdc.Ltdc.L2cr.SetLen(false)

		regltdc.Ltdc.L2whpcr.Store(whpcr)
		regltdc.Ltdc.L2wvpcr.Store(wvpcr)
		regltdc.Ltdc.L2ckcr.Store(0)
		regltdc.Ltdc.L2pfcr.Store(pfcr)
		regltdc.Ltdc.L2cacr.Store(cacr)
		regltdc.Ltdc.L2dccr.Store(dccr)
		regltdc.Ltdc.L2bfcr.Store(bfcr)
		regltdc.Ltdc.L2cfbar.Store(cfbar)
		regltdc.Ltdc.L2cfblr.Store(cfblr)
		regltdc.Ltdc.L2cfblnr.Store(cfblnr)

		regltdc.Ltdc.L2cr.SetColken(false)
		regltdc.Ltdc.L2cr.SetCluten(false)
		if layer.Enabled {
			regltdc.Ltdc.L2cr.SetLen(true)
		}
	}
}

func disableLayer(layer Layer) {
	switch layer {
	case Layer1:
		regltdc.Ltdc.L1cr.SetLen(false)
	case Layer2:
		regltdc.Ltdc.L2cr.SetLen(false)
	}
}

func reloadImmediate() {
	regltdc.Ltdc.Srcr.SetImr(true)
}

func reloadVerticalBlank() {
	regltdc.Ltdc.Srcr.SetVbr(true)
}

func clearInterruptFlags() {
	regltdc.Ltdc.Icr.Store(
		regltdc.RegisterIcrFieldClifMask |
			regltdc.RegisterIcrFieldCfuifMask |
			regltdc.RegisterIcrFieldCterrifMask |
			regltdc.RegisterIcrFieldCrrifMask,
	)
}

const reloadPollLimit = 20_000_000

func (l LTDC) SetLayerFramebufferAndWait(
	layer Layer,
	framebuffer unsafe.Pointer,
) error {
	if framebuffer == nil {
		return corehal.ErrInvalidConfig
	}
	if !l.Enabled() {
		return corehal.ErrInvalidState
	}

	switch layer {
	case Layer1:
		regltdc.Ltdc.L1cfbar.Store(uint32(uintptr(framebuffer)))
	case Layer2:
		regltdc.Ltdc.L2cfbar.Store(uint32(uintptr(framebuffer)))
	default:
		return corehal.ErrInvalidConfig
	}

	cortexm.DSB()

	// Schedule shadow-register reload at the next vertical blank.
	regltdc.Ltdc.Srcr.SetVbr(true)
	cortexm.DSB()

	// VBR is cleared by hardware when the requested reload is applied.
	for i := uint32(0); regltdc.Ltdc.Srcr.GetVbr(); i++ {
		if i >= reloadPollLimit {
			return corehal.ErrTimeout
		}
	}

	return nil
}
