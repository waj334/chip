//go:build stm32h7x7

// Package dma2d is a HAL for the STM32H7x7 Chrom-ART Accelerator (DMA2D).
//
// DMA2D is a dedicated 2D graphics DMA that can fill, copy, pixel-format
// convert, and alpha-blend rectangular memory regions without CPU
// involvement. It supports four transfer modes:
//
//	Fill	register-to-memory	solid color fill of a rectangle
//	Copy	memory-to-memory	raw copy, source and destination formats equal
//	Convert	memory-to-memory	copy with pixel-format conversion (PFC)
//	Blend	memory-to-memory	alpha blend foreground over background
//
// All addresses are caller-supplied physical addresses (framebuffers in
// SRAM, AXI SRAM, or external SDRAM). The HAL never allocates: geometry and
// layer descriptors are passed by value and live on the caller's stack.
//
// Cache coherency is the caller's responsibility. DMA2D is an AHB bus master
// and does not participate in the Cortex-M7 data cache. When buffers live in
// cacheable memory, the caller must clean source data out of the D-cache
// before starting a transfer and invalidate the destination region after the
// transfer completes — or place the buffers in a non-cacheable MPU region.
//
// This HAL targets the CM7 core. A single transfer runs at a time; concurrent
// callers are serialized, and a request issued while a transfer is in flight
// returns ErrInvalidState.
package dma2d

import (
	"sync"
	"volatile"

	stm32h7x7 "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/dma2d"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	corehal "pkg.si-go.dev/chip/core/hal"
)

// DMA2D is the single Chrom-ART Accelerator instance.
var DMA2D _dma2d

// Format identifies a pixel colour mode.
//
// All formats are valid as a transfer input (foreground/background). Only the
// first five (ARGB8888 through ARGB4444) are valid as a transfer output; the
// indexed and alpha-only formats exist solely as inputs to the PFC unit.
type Format uint8

const (
	FormatARGB8888 Format = iota // 32 bpp
	FormatRGB888                 // 24 bpp
	FormatRGB565                 // 16 bpp
	FormatARGB1555               // 16 bpp
	FormatARGB4444               // 16 bpp
	FormatL8                     // 8 bpp  indexed (input only)
	FormatAL44                   // 8 bpp  indexed + alpha (input only)
	FormatAL88                   // 16 bpp indexed + alpha (input only)
	FormatL4                     // 4 bpp  indexed (input only)
	FormatA8                     // 8 bpp  alpha only (input only)
	FormatA4                     // 4 bpp  alpha only (input only)
)

func (f Format) validInput() bool  { return f <= FormatA4 }
func (f Format) validOutput() bool { return f <= FormatARGB4444 }

// AlphaMode selects how a blend layer's alpha channel is derived.
type AlphaMode uint8

const (
	// AlphaNoModify uses the layer's own per-pixel alpha unchanged.
	AlphaNoModify AlphaMode = iota
	// AlphaReplace replaces the layer alpha with the constant Alpha value.
	AlphaReplace
	// AlphaCombine multiplies the layer's per-pixel alpha by Alpha/255.
	AlphaCombine
)

// Geometry limits imposed by the NLR and offset register field widths.
const (
	maxPixelsPerLine = 0x3FFF // NLR.PL: 14 bits
	maxLineOffset    = 0x3FFF // xOR.LO: 14 bits
	// NLR.NL is 16 bits, so any uint16 line count is valid.
)

// Internal CR.MODE encodings.
const (
	modeM2M      uint8 = 0 // memory-to-memory
	modeM2MPFC   uint8 = 1 // memory-to-memory with pixel-format conversion
	modeM2MBlend uint8 = 2 // memory-to-memory with blending
	modeR2M      uint8 = 3 // register-to-memory (fill)
)

// Source describes an input region for a Copy or Convert transfer.
type Source struct {
	Addr       uintptr // start address of the first pixel
	Format     Format  // input pixel format
	LineOffset uint16  // pixels to skip between lines (0 for a contiguous buffer)
}

// BlendSource describes one of the two input layers for a Blend transfer.
type BlendSource struct {
	Addr        uintptr   // start address of the first pixel
	Format      Format    // input pixel format
	LineOffset  uint16    // pixels to skip between lines
	AlphaMode   AlphaMode // how to derive this layer's alpha
	Alpha       uint8     // constant alpha for AlphaReplace / AlphaCombine
	RedBlueSwap bool      // swap the red and blue channels of this layer
}

// Target describes the output region for any transfer.
type Target struct {
	Addr        uintptr // start address of the first pixel
	Format      Format  // output pixel format (must be ARGB8888..ARGB4444)
	LineOffset  uint16  // pixels to skip between lines
	RedBlueSwap bool    // swap the red and blue channels on output
	AlphaInvert bool    // invert the output alpha channel
}

// Config holds one-time peripheral configuration.
type Config struct {
	// Interrupt enables completion via the DMA2D global interrupt. When set,
	// transfer methods return immediately, and completion is reported through
	// the registered CompletionHandler and observable via Wait. When clear,
	// transfer methods block by polling and return the result directly.
	Interrupt bool

	// Priority is the NVIC priority for the DMA2D interrupt (only used when
	// Interrupt is set).
	Priority uint8

	// DeadTimeEnable inserts a minimum idle period between two consecutive
	// AHB accesses, throttling DMA2D bus usage so the CPU and other masters
	// retain bandwidth.
	DeadTimeEnable bool

	// DeadTime is the dead-time duration in AHB clock cycles (only used when
	// DeadTimeEnable is set).
	DeadTime uint8
}

// Event reports the outcome of a completed transfer.
type Event struct {
	Complete    bool // transfer completed
	TransferErr bool // AHB bus transfer error
	ClutErr     bool // CLUT access error
	ConfigErr   bool // configuration error
}

// OK reports whether the transfer completed without any error.
func (e Event) OK() bool {
	return e.Complete && !e.TransferErr && !e.ClutErr && !e.ConfigErr
}

// Err returns a non-nil error if the transfer reported any error condition.
func (e Event) Err() corehal.Error {
	if e.TransferErr || e.ClutErr || e.ConfigErr {
		return corehal.ErrInvalidState
	}
	return corehal.NoError
}

// CompletionHandler is invoked from the interrupt context when an interrupt-mode
// transfer completes. It must be inexpensive and must not block.
type CompletionHandler func(Event)

// _dma2d holds the driver state for the single DMA2D peripheral.
type _dma2d struct {
	mutex        sync.Mutex
	busy         uint32 // 1 while a transfer is in flight (volatile access)
	useInterrupt bool
	handler      CompletionHandler
	lastEvent    Event
}

// Configure enables the DMA2D clock, applies the supplied configuration, and
// arms or disarms the interrupt. It must be called once before any transfer.
func (d *_dma2d) Configure(config Config) error {
	d.mutex.Lock()

	// Enable the peripheral clock (AHB3) and pulse the reset line so the
	// peripheral starts from a known state.
	rcc.Rcc.Ahb3enr.SetDma2den(true)
	rcc.Rcc.Ahb3rstr.SetDma2drst(true)
	rcc.Rcc.Ahb3rstr.SetDma2drst(false)

	// Mask all interrupt sources until a transfer arms them.
	dma2d.Dma2d.Cr.SetTcie(false)
	dma2d.Dma2d.Cr.SetTeie(false)
	dma2d.Dma2d.Cr.SetTwie(false)
	dma2d.Dma2d.Cr.SetCaeie(false)
	dma2d.Dma2d.Cr.SetCeie(false)
	dma2d.Dma2d.Cr.SetCtcie(false)

	// AHB dead-time throttling.
	dma2d.Dma2d.Amtcr.SetDt(config.DeadTime)
	dma2d.Dma2d.Amtcr.SetEn(config.DeadTimeEnable)

	d.useInterrupt = config.Interrupt
	if config.Interrupt {
		stm32h7x7.IrqDma2d.SetPriority(config.Priority)
		stm32h7x7.IrqDma2d.Enable()
	} else {
		stm32h7x7.IrqDma2d.Disable()
	}

	volatile.StoreUint32(&d.busy, 0)
	d.mutex.Unlock()
	return nil
}

// SetCompletionHandler registers the callback invoked when an interrupt-mode
// transfer completes. Pass nil to clear it. Set it before issuing transfers.
func (d *_dma2d) SetCompletionHandler(fn CompletionHandler) {
	d.mutex.Lock()
	d.handler = fn
	d.mutex.Unlock()
}

// Fill writes a solid colour into a width×height rectangle of the destination
// (register-to-memory mode). The colour must already be encoded in the
// destination's pixel format, right-aligned within the 32-bit value.
func (d *_dma2d) Fill(dst Target, width, height uint16, color uint32) corehal.Error {
	d.mutex.Lock()

	if volatile.LoadUint32(&d.busy) != 0 {
		d.mutex.Unlock()
		return corehal.ErrInvalidState
	}
	if err := validateSize(width, height); err != corehal.NoError {
		d.mutex.Unlock()
		return err
	}
	if !dst.Format.validOutput() || dst.LineOffset > maxLineOffset {
		d.mutex.Unlock()
		return corehal.ErrInvalidConfig
	}

	dma2d.Dma2d.Opfccr.SetCm(uint8(dst.Format))
	dma2d.Dma2d.Opfccr.SetRbs(dst.RedBlueSwap)
	dma2d.Dma2d.Opfccr.SetAi(dst.AlphaInvert)
	dma2d.Dma2d.Ocolr.Store(color)
	dma2d.Dma2d.Omar.SetMa(uint32(dst.Addr))
	dma2d.Dma2d.Oor.SetLo(dst.LineOffset)
	dma2d.Dma2d.Nlr.SetPl(width)
	dma2d.Dma2d.Nlr.SetNl(height)

	d.launch(modeR2M)
	err := d.finish()
	d.mutex.Unlock()
	return err
}

// Copy performs a raw memory-to-memory copy of a width×height rectangle. The
// source and destination formats must be identical; use Convert when they
// differ.
func (d *_dma2d) Copy(src Source, dst Target, width, height uint16) corehal.Error {
	d.mutex.Lock()

	if volatile.LoadUint32(&d.busy) != 0 {
		d.mutex.Unlock()
		return corehal.ErrInvalidState
	}
	if err := validateSize(width, height); err != corehal.NoError {
		d.mutex.Unlock()
		return err
	}
	if !src.Format.validInput() || !dst.Format.validOutput() {
		d.mutex.Unlock()
		return corehal.ErrInvalidConfig
	}
	if src.Format != dst.Format {
		d.mutex.Unlock()
		return corehal.ErrInvalidConfig
	}
	if src.LineOffset > maxLineOffset || dst.LineOffset > maxLineOffset {
		d.mutex.Unlock()
		return corehal.ErrInvalidConfig
	}

	dma2d.Dma2d.Fgpfccr.SetCm(uint8(src.Format))
	dma2d.Dma2d.Fgpfccr.SetAm(uint8(AlphaNoModify))
	dma2d.Dma2d.Fgpfccr.SetAlpha(0)
	dma2d.Dma2d.Fgpfccr.SetRbs(false)
	dma2d.Dma2d.Fgmar.SetMa(uint32(src.Addr))
	dma2d.Dma2d.Fgor.SetLo(src.LineOffset)

	dma2d.Dma2d.Opfccr.SetCm(uint8(dst.Format))
	dma2d.Dma2d.Opfccr.SetRbs(dst.RedBlueSwap)
	dma2d.Dma2d.Opfccr.SetAi(dst.AlphaInvert)
	dma2d.Dma2d.Omar.SetMa(uint32(dst.Addr))
	dma2d.Dma2d.Oor.SetLo(dst.LineOffset)
	dma2d.Dma2d.Nlr.SetPl(width)
	dma2d.Dma2d.Nlr.SetNl(height)

	d.launch(modeM2M)
	err := d.finish()
	d.mutex.Unlock()
	return err
}

// Convert copies a width×height rectangle while converting from the source
// pixel format to the destination pixel format (memory-to-memory with PFC).
// When the source has no alpha channel and the destination does, the output
// alpha is forced to fully opaque by the hardware.
func (d *_dma2d) Convert(src Source, dst Target, width, height uint16) error {
	d.mutex.Lock()

	if volatile.LoadUint32(&d.busy) != 0 {
		d.mutex.Unlock()
		return corehal.ErrInvalidState
	}
	if err := validateSize(width, height); err != corehal.NoError {
		d.mutex.Unlock()
		return err
	}
	if !src.Format.validInput() || !dst.Format.validOutput() {
		d.mutex.Unlock()
		return corehal.ErrInvalidConfig
	}
	if src.LineOffset > maxLineOffset || dst.LineOffset > maxLineOffset {
		d.mutex.Unlock()
		return corehal.ErrInvalidConfig
	}

	dma2d.Dma2d.Fgpfccr.SetCm(uint8(src.Format))
	dma2d.Dma2d.Fgpfccr.SetAm(uint8(AlphaNoModify))
	dma2d.Dma2d.Fgpfccr.SetAlpha(0)
	dma2d.Dma2d.Fgpfccr.SetRbs(false)
	dma2d.Dma2d.Fgmar.SetMa(uint32(src.Addr))
	dma2d.Dma2d.Fgor.SetLo(src.LineOffset)

	dma2d.Dma2d.Opfccr.SetCm(uint8(dst.Format))
	dma2d.Dma2d.Opfccr.SetRbs(dst.RedBlueSwap)
	dma2d.Dma2d.Opfccr.SetAi(dst.AlphaInvert)
	dma2d.Dma2d.Omar.SetMa(uint32(dst.Addr))
	dma2d.Dma2d.Oor.SetLo(dst.LineOffset)
	dma2d.Dma2d.Nlr.SetPl(width)
	dma2d.Dma2d.Nlr.SetNl(height)

	d.launch(modeM2MPFC)
	err := d.finish()
	d.mutex.Unlock()
	return err
}

// Blend alpha-blends the foreground over the background into the destination
// (memory-to-memory with blending). All three regions share the same
// width×height geometry.
func (d *_dma2d) Blend(fg, bg BlendSource, dst Target, width, height uint16) corehal.Error {
	d.mutex.Lock()

	if volatile.LoadUint32(&d.busy) != 0 {
		d.mutex.Unlock()
		return corehal.ErrInvalidState
	}
	if err := validateSize(width, height); err != corehal.NoError {
		d.mutex.Unlock()
		return err
	}
	if !fg.Format.validInput() || !bg.Format.validInput() || !dst.Format.validOutput() {
		d.mutex.Unlock()
		return corehal.ErrInvalidConfig
	}
	if fg.AlphaMode > AlphaCombine || bg.AlphaMode > AlphaCombine {
		d.mutex.Unlock()
		return corehal.ErrInvalidConfig
	}
	if fg.LineOffset > maxLineOffset || bg.LineOffset > maxLineOffset || dst.LineOffset > maxLineOffset {
		d.mutex.Unlock()
		return corehal.ErrInvalidConfig
	}

	dma2d.Dma2d.Fgpfccr.SetCm(uint8(fg.Format))
	dma2d.Dma2d.Fgpfccr.SetAm(uint8(fg.AlphaMode))
	dma2d.Dma2d.Fgpfccr.SetAlpha(fg.Alpha)
	dma2d.Dma2d.Fgpfccr.SetRbs(fg.RedBlueSwap)
	dma2d.Dma2d.Fgmar.SetMa(uint32(fg.Addr))
	dma2d.Dma2d.Fgor.SetLo(fg.LineOffset)

	dma2d.Dma2d.Bgpfccr.SetCm(uint8(bg.Format))
	dma2d.Dma2d.Bgpfccr.SetAm(uint8(bg.AlphaMode))
	dma2d.Dma2d.Bgpfccr.SetAlpha(bg.Alpha)
	dma2d.Dma2d.Bgpfccr.SetRbs(bg.RedBlueSwap)
	dma2d.Dma2d.Bgmar.SetMa(uint32(bg.Addr))
	dma2d.Dma2d.Bgor.SetLo(bg.LineOffset)

	dma2d.Dma2d.Opfccr.SetCm(uint8(dst.Format))
	dma2d.Dma2d.Opfccr.SetRbs(dst.RedBlueSwap)
	dma2d.Dma2d.Opfccr.SetAi(dst.AlphaInvert)
	dma2d.Dma2d.Omar.SetMa(uint32(dst.Addr))
	dma2d.Dma2d.Oor.SetLo(dst.LineOffset)
	dma2d.Dma2d.Nlr.SetPl(width)
	dma2d.Dma2d.Nlr.SetNl(height)

	d.launch(modeM2MBlend)
	err := d.finish()
	d.mutex.Unlock()
	return err
}

// Busy reports whether a transfer is currently in flight.
func (d *_dma2d) Busy() bool {
	return volatile.LoadUint32(&d.busy) != 0
}

// Wait blocks until the in-flight transfer completes and returns its result.
//
// In interrupt mode it spins until the interrupt handler clears the busy flag
// and then returns the recorded outcome. In polling mode it drains the
// transfer directly. Calling Wait when no transfer is in flight returns the
// outcome of the most recently completed transfer (nil if none ran).
func (d *_dma2d) Wait() corehal.Error {
	if d.useInterrupt {
		for volatile.LoadUint32(&d.busy) != 0 {
		}
		return d.lastEvent.Err()
	}

	for dma2d.Dma2d.Cr.GetStart() {
	}

	e := readStatus()
	clearFlags()
	d.lastEvent = e
	volatile.StoreUint32(&d.busy, 0)
	return e.Err()
}

// Abort cancels an in-flight transfer. It blocks until the abort takes effect.
func (d *_dma2d) Abort() {
	d.mutex.Lock()
	dma2d.Dma2d.Cr.SetAbort(true)
	for dma2d.Dma2d.Cr.GetAbort() {
	}
	clearFlags()
	volatile.StoreUint32(&d.busy, 0)
	d.mutex.Unlock()
}

// launch clears stale status, programs the mode, arms interrupts when
// configured, marks the peripheral busy, and starts the transfer. The caller
// must hold the mutex and must have verified the peripheral is not busy.
func (d *_dma2d) launch(mode uint8) {
	clearFlags()

	dma2d.Dma2d.Cr.SetMode(mode)

	if d.useInterrupt {
		dma2d.Dma2d.Cr.SetTcie(true)
		dma2d.Dma2d.Cr.SetTeie(true)
		dma2d.Dma2d.Cr.SetCaeie(true)
		dma2d.Dma2d.Cr.SetCeie(true)
	}

	volatile.StoreUint32(&d.busy, 1)
	dma2d.Dma2d.Cr.SetStart(true)
}

// finish returns immediately in interrupt mode (completion arrives via the
// handler / Wait) and otherwise blocks until the transfer drains.
func (d *_dma2d) finish() corehal.Error {
	if d.useInterrupt {
		return corehal.NoError
	}
	return d.Wait()
}

func validateSize(width, height uint16) corehal.Error {
	if width == 0 || height == 0 {
		return corehal.ErrInvalidParameter
	}
	if width > maxPixelsPerLine {
		return corehal.ErrInvalidParameter
	}
	return corehal.NoError
}

func readStatus() Event {
	return Event{
		Complete:    dma2d.Dma2d.Isr.GetTcif(),
		TransferErr: dma2d.Dma2d.Isr.GetTeif(),
		ClutErr:     dma2d.Dma2d.Isr.GetCaeif(),
		ConfigErr:   dma2d.Dma2d.Isr.GetCeif(),
	}
}

func clearFlags() {
	dma2d.Dma2d.Ifcr.SetCtcif(true)  // transfer complete
	dma2d.Dma2d.Ifcr.SetCteif(true)  // transfer error
	dma2d.Dma2d.Ifcr.SetCaecif(true) // CLUT access error
	dma2d.Dma2d.Ifcr.SetCceif(true)  // configuration error
	dma2d.Dma2d.Ifcr.SetCctcif(true) // CLUT transfer complete
	dma2d.Dma2d.Ifcr.SetCtwif(true)  // transfer watermark
}

//sigo:interrupt dma2dHandler Dma2dHandler
func dma2dHandler() {
	e := readStatus()

	// Disarm interrupt sources before clearing the flags so the line drops.
	dma2d.Dma2d.Cr.SetTcie(false)
	dma2d.Dma2d.Cr.SetTeie(false)
	dma2d.Dma2d.Cr.SetCaeie(false)
	dma2d.Dma2d.Cr.SetCeie(false)
	clearFlags()

	DMA2D.lastEvent = e
	volatile.StoreUint32(&DMA2D.busy, 0)

	if h := DMA2D.handler; h != nil {
		h(e)
	}
}
