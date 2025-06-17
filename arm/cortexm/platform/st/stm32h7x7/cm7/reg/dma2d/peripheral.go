//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package dma2d

import (
	"unsafe"
	"volatile"
)

var (
	Dma2d = (*_dma2d)(unsafe.Pointer(uintptr(0x52001000)))
)

type _dma2d struct {
	Cr      RegisterCrType
	Isr     RegisterIsrType
	Ifcr    RegisterIfcrType
	Fgmar   RegisterFgmarType
	Fgor    RegisterFgorType
	Bgmar   RegisterBgmarType
	Bgor    RegisterBgorType
	Fgpfccr RegisterFgpfccrType
	Fgcolr  RegisterFgcolrType
	Bgpfccr RegisterBgpfccrType
	Bgcolr  RegisterBgcolrType
	Fgcmar  RegisterFgcmarType
	Bgcmar  RegisterBgcmarType
	Opfccr  RegisterOpfccrType
	Ocolr   RegisterOcolrType
	Omar    RegisterOmarType
	Oor     RegisterOorType
	Nlr     RegisterNlrType
	Lwr     RegisterLwrType
	Amtcr   RegisterAmtcrType
}

// RegisterCrType DMA2D control register
type RegisterCrType uint32

func (r *RegisterCrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrFieldStartShift = 0
	RegisterCrFieldStartMask  = 0x1
)

// GetStart Start This bit can be used to launch the DMA2D according to the parameters loaded in the various configuration registers
func (r *RegisterCrType) GetStart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldStartMask) != 0
}

// SetStart Start This bit can be used to launch the DMA2D according to the parameters loaded in the various configuration registers
func (r *RegisterCrType) SetStart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldStartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldStartMask)
	}
}

const (
	RegisterCrFieldSuspShift = 1
	RegisterCrFieldSuspMask  = 0x2
)

// GetSusp Suspend This bit can be used to suspend the current transfer. This bit is set and reset by software. It is automatically reset by hardware when the START bit is reset.
func (r *RegisterCrType) GetSusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSuspMask) != 0
}

// SetSusp Suspend This bit can be used to suspend the current transfer. This bit is set and reset by software. It is automatically reset by hardware when the START bit is reset.
func (r *RegisterCrType) SetSusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldSuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSuspMask)
	}
}

const (
	RegisterCrFieldAbortShift = 2
	RegisterCrFieldAbortMask  = 0x4
)

// GetAbort Abort This bit can be used to abort the current transfer. This bit is set by software and is automatically reset by hardware when the START bit is reset.
func (r *RegisterCrType) GetAbort() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAbortMask) != 0
}

// SetAbort Abort This bit can be used to abort the current transfer. This bit is set by software and is automatically reset by hardware when the START bit is reset.
func (r *RegisterCrType) SetAbort(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAbortMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAbortMask)
	}
}

const (
	RegisterCrFieldTeieShift = 8
	RegisterCrFieldTeieMask  = 0x100
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTeieMask)
	}
}

const (
	RegisterCrFieldTcieShift = 9
	RegisterCrFieldTcieMask  = 0x200
)

// GetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTcieMask)
	}
}

const (
	RegisterCrFieldTwieShift = 10
	RegisterCrFieldTwieMask  = 0x400
)

// GetTwie Transfer watermark interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) GetTwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTwieMask) != 0
}

// SetTwie Transfer watermark interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) SetTwie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTwieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTwieMask)
	}
}

const (
	RegisterCrFieldCaeieShift = 11
	RegisterCrFieldCaeieMask  = 0x800
)

// GetCaeie CLUT access error interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) GetCaeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCaeieMask) != 0
}

// SetCaeie CLUT access error interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) SetCaeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCaeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCaeieMask)
	}
}

const (
	RegisterCrFieldCtcieShift = 12
	RegisterCrFieldCtcieMask  = 0x1000
)

// GetCtcie CLUT transfer complete interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCtcieMask) != 0
}

// SetCtcie CLUT transfer complete interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCtcieMask)
	}
}

const (
	RegisterCrFieldCeieShift = 13
	RegisterCrFieldCeieMask  = 0x2000
)

// GetCeie Configuration Error Interrupt Enable This bit is set and cleared by software.
func (r *RegisterCrType) GetCeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCeieMask) != 0
}

// SetCeie Configuration Error Interrupt Enable This bit is set and cleared by software.
func (r *RegisterCrType) SetCeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCeieMask)
	}
}

const (
	RegisterCrFieldModeShift = 16
	RegisterCrFieldModeMask  = 0x30000
)

// GetMode DMA2D mode This bit is set and cleared by software. It cannot be modified while a transfer is ongoing.
func (r *RegisterCrType) GetMode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldModeMask) >> RegisterCrFieldModeShift)
}

// SetMode DMA2D mode This bit is set and cleared by software. It cannot be modified while a transfer is ongoing.
func (r *RegisterCrType) SetMode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldModeMask)|(uint32(value)<<RegisterCrFieldModeShift))
}

// RegisterIsrType DMA2D Interrupt Status Register
type RegisterIsrType uint32

func (r *RegisterIsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIsrFieldTeifShift = 0
	RegisterIsrFieldTeifMask  = 0x1
)

// GetTeif Transfer error interrupt flag This bit is set when an error occurs during a DMA transfer (data transfer or automatic CLUT loading).
func (r *RegisterIsrType) GetTeif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTeifMask) != 0
}

// SetTeif Transfer error interrupt flag This bit is set when an error occurs during a DMA transfer (data transfer or automatic CLUT loading).
func (r *RegisterIsrType) SetTeif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTeifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTeifMask)
	}
}

const (
	RegisterIsrFieldTcifShift = 1
	RegisterIsrFieldTcifMask  = 0x2
)

// GetTcif Transfer complete interrupt flag This bit is set when a DMA2D transfer operation is complete (data transfer only).
func (r *RegisterIsrType) GetTcif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcifMask) != 0
}

// SetTcif Transfer complete interrupt flag This bit is set when a DMA2D transfer operation is complete (data transfer only).
func (r *RegisterIsrType) SetTcif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTcifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTcifMask)
	}
}

const (
	RegisterIsrFieldTwifShift = 2
	RegisterIsrFieldTwifMask  = 0x4
)

// GetTwif Transfer watermark interrupt flag This bit is set when the last pixel of the watermarked line has been transferred.
func (r *RegisterIsrType) GetTwif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTwifMask) != 0
}

// SetTwif Transfer watermark interrupt flag This bit is set when the last pixel of the watermarked line has been transferred.
func (r *RegisterIsrType) SetTwif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTwifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTwifMask)
	}
}

const (
	RegisterIsrFieldCaeifShift = 3
	RegisterIsrFieldCaeifMask  = 0x8
)

// GetCaeif CLUT access error interrupt flag This bit is set when the CPU accesses the CLUT while the CLUT is being automatically copied from a system memory to the internal DMA2D.
func (r *RegisterIsrType) GetCaeif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCaeifMask) != 0
}

// SetCaeif CLUT access error interrupt flag This bit is set when the CPU accesses the CLUT while the CLUT is being automatically copied from a system memory to the internal DMA2D.
func (r *RegisterIsrType) SetCaeif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldCaeifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldCaeifMask)
	}
}

const (
	RegisterIsrFieldCtcifShift = 4
	RegisterIsrFieldCtcifMask  = 0x10
)

// GetCtcif CLUT transfer complete interrupt flag This bit is set when the CLUT copy from a system memory area to the internal DMA2D memory is complete.
func (r *RegisterIsrType) GetCtcif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCtcifMask) != 0
}

// SetCtcif CLUT transfer complete interrupt flag This bit is set when the CLUT copy from a system memory area to the internal DMA2D memory is complete.
func (r *RegisterIsrType) SetCtcif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldCtcifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldCtcifMask)
	}
}

const (
	RegisterIsrFieldCeifShift = 5
	RegisterIsrFieldCeifMask  = 0x20
)

// GetCeif Configuration error interrupt flag This bit is set when the START bit of DMA2D_CR, DMA2DFGPFCCR or DMA2D_BGPFCCR is set and a wrong configuration has been programmed.
func (r *RegisterIsrType) GetCeif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCeifMask) != 0
}

// SetCeif Configuration error interrupt flag This bit is set when the START bit of DMA2D_CR, DMA2DFGPFCCR or DMA2D_BGPFCCR is set and a wrong configuration has been programmed.
func (r *RegisterIsrType) SetCeif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldCeifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldCeifMask)
	}
}

// RegisterIfcrType DMA2D interrupt flag clear register
type RegisterIfcrType uint32

func (r *RegisterIfcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIfcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIfcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIfcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIfcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIfcrFieldCteifShift = 0
	RegisterIfcrFieldCteifMask  = 0x1
)

// GetCteif Clear Transfer error interrupt flag Programming this bit to 1 clears the TEIF flag in the DMA2D_ISR register
func (r *RegisterIfcrType) GetCteif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCteifMask) != 0
}

// SetCteif Clear Transfer error interrupt flag Programming this bit to 1 clears the TEIF flag in the DMA2D_ISR register
func (r *RegisterIfcrType) SetCteif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCteifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCteifMask)
	}
}

const (
	RegisterIfcrFieldCtcifShift = 1
	RegisterIfcrFieldCtcifMask  = 0x2
)

// GetCtcif Clear transfer complete interrupt flag Programming this bit to 1 clears the TCIF flag in the DMA2D_ISR register
func (r *RegisterIfcrType) GetCtcif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCtcifMask) != 0
}

// SetCtcif Clear transfer complete interrupt flag Programming this bit to 1 clears the TCIF flag in the DMA2D_ISR register
func (r *RegisterIfcrType) SetCtcif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCtcifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCtcifMask)
	}
}

const (
	RegisterIfcrFieldCtwifShift = 2
	RegisterIfcrFieldCtwifMask  = 0x4
)

// GetCtwif Clear transfer watermark interrupt flag Programming this bit to 1 clears the TWIF flag in the DMA2D_ISR register
func (r *RegisterIfcrType) GetCtwif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCtwifMask) != 0
}

// SetCtwif Clear transfer watermark interrupt flag Programming this bit to 1 clears the TWIF flag in the DMA2D_ISR register
func (r *RegisterIfcrType) SetCtwif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCtwifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCtwifMask)
	}
}

const (
	RegisterIfcrFieldCaecifShift = 3
	RegisterIfcrFieldCaecifMask  = 0x8
)

// GetCaecif Clear CLUT access error interrupt flag Programming this bit to 1 clears the CAEIF flag in the DMA2D_ISR register
func (r *RegisterIfcrType) GetCaecif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCaecifMask) != 0
}

// SetCaecif Clear CLUT access error interrupt flag Programming this bit to 1 clears the CAEIF flag in the DMA2D_ISR register
func (r *RegisterIfcrType) SetCaecif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCaecifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCaecifMask)
	}
}

const (
	RegisterIfcrFieldCctcifShift = 4
	RegisterIfcrFieldCctcifMask  = 0x10
)

// GetCctcif Clear CLUT transfer complete interrupt flag Programming this bit to 1 clears the CTCIF flag in the DMA2D_ISR register
func (r *RegisterIfcrType) GetCctcif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCctcifMask) != 0
}

// SetCctcif Clear CLUT transfer complete interrupt flag Programming this bit to 1 clears the CTCIF flag in the DMA2D_ISR register
func (r *RegisterIfcrType) SetCctcif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCctcifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCctcifMask)
	}
}

const (
	RegisterIfcrFieldCceifShift = 5
	RegisterIfcrFieldCceifMask  = 0x20
)

// GetCceif Clear configuration error interrupt flag Programming this bit to 1 clears the CEIF flag in the DMA2D_ISR register
func (r *RegisterIfcrType) GetCceif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCceifMask) != 0
}

// SetCceif Clear configuration error interrupt flag Programming this bit to 1 clears the CEIF flag in the DMA2D_ISR register
func (r *RegisterIfcrType) SetCceif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCceifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCceifMask)
	}
}

// RegisterFgmarType DMA2D foreground memory address register
type RegisterFgmarType uint32

func (r *RegisterFgmarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterFgmarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterFgmarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterFgmarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterFgmarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterFgmarFieldMaShift = 0
	RegisterFgmarFieldMaMask  = 0xffffffff
)

// GetMa Memory address Address of the data used for the foreground image. This register can only be written when data transfers are disabled. Once the data transfer has started, this register is read-only. The address alignment must match the image format selected e.g. a 32-bit per pixel format must be 32-bit aligned, a 16-bit per pixel format must be 16-bit aligned and a 4-bit per pixel format must be 8-bit aligned.
func (r *RegisterFgmarType) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFgmarFieldMaMask) >> RegisterFgmarFieldMaShift)
}

// SetMa Memory address Address of the data used for the foreground image. This register can only be written when data transfers are disabled. Once the data transfer has started, this register is read-only. The address alignment must match the image format selected e.g. a 32-bit per pixel format must be 32-bit aligned, a 16-bit per pixel format must be 16-bit aligned and a 4-bit per pixel format must be 8-bit aligned.
func (r *RegisterFgmarType) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFgmarFieldMaMask)|(uint32(value)<<RegisterFgmarFieldMaShift))
}

// RegisterFgorType DMA2D foreground offset register
type RegisterFgorType uint32

func (r *RegisterFgorType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterFgorType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterFgorType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterFgorType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterFgorType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterFgorFieldLoShift = 0
	RegisterFgorFieldLoMask  = 0x3fff
)

// GetLo Line offset Line offset used for the foreground expressed in pixel. This value is used to generate the address. It is added at the end of each line to determine the starting address of the next line. These bits can only be written when data transfers are disabled. Once a data transfer has started, they become read-only. If the image format is 4-bit per pixel, the line offset must be even.
func (r *RegisterFgorType) GetLo() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFgorFieldLoMask) >> RegisterFgorFieldLoShift)
}

// SetLo Line offset Line offset used for the foreground expressed in pixel. This value is used to generate the address. It is added at the end of each line to determine the starting address of the next line. These bits can only be written when data transfers are disabled. Once a data transfer has started, they become read-only. If the image format is 4-bit per pixel, the line offset must be even.
func (r *RegisterFgorType) SetLo(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFgorFieldLoMask)|(uint32(value)<<RegisterFgorFieldLoShift))
}

// RegisterBgmarType DMA2D background memory address register
type RegisterBgmarType uint32

func (r *RegisterBgmarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBgmarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBgmarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBgmarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBgmarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBgmarFieldMaShift = 0
	RegisterBgmarFieldMaMask  = 0xffffffff
)

// GetMa Memory address Address of the data used for the background image. This register can only be written when data transfers are disabled. Once a data transfer has started, this register is read-only. The address alignment must match the image format selected e.g. a 32-bit per pixel format must be 32-bit aligned, a 16-bit per pixel format must be 16-bit aligned and a 4-bit per pixel format must be 8-bit aligned.
func (r *RegisterBgmarType) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterBgmarFieldMaMask) >> RegisterBgmarFieldMaShift)
}

// SetMa Memory address Address of the data used for the background image. This register can only be written when data transfers are disabled. Once a data transfer has started, this register is read-only. The address alignment must match the image format selected e.g. a 32-bit per pixel format must be 32-bit aligned, a 16-bit per pixel format must be 16-bit aligned and a 4-bit per pixel format must be 8-bit aligned.
func (r *RegisterBgmarType) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBgmarFieldMaMask)|(uint32(value)<<RegisterBgmarFieldMaShift))
}

// RegisterBgorType DMA2D background offset register
type RegisterBgorType uint32

func (r *RegisterBgorType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBgorType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBgorType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBgorType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBgorType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBgorFieldLoShift = 0
	RegisterBgorFieldLoMask  = 0x3fff
)

// GetLo Line offset Line offset used for the background image (expressed in pixel). This value is used for the address generation. It is added at the end of each line to determine the starting address of the next line. These bits can only be written when data transfers are disabled. Once data transfer has started, they become read-only. If the image format is 4-bit per pixel, the line offset must be even.
func (r *RegisterBgorType) GetLo() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBgorFieldLoMask) >> RegisterBgorFieldLoShift)
}

// SetLo Line offset Line offset used for the background image (expressed in pixel). This value is used for the address generation. It is added at the end of each line to determine the starting address of the next line. These bits can only be written when data transfers are disabled. Once data transfer has started, they become read-only. If the image format is 4-bit per pixel, the line offset must be even.
func (r *RegisterBgorType) SetLo(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBgorFieldLoMask)|(uint32(value)<<RegisterBgorFieldLoShift))
}

// RegisterFgpfccrType DMA2D foreground PFC control register
type RegisterFgpfccrType uint32

func (r *RegisterFgpfccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterFgpfccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterFgpfccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterFgpfccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterFgpfccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterFgpfccrFieldCmShift = 0
	RegisterFgpfccrFieldCmMask  = 0xf
)

// GetCm Color mode These bits defines the color format of the foreground image. They can only be written when data transfers are disabled. Once the transfer has started, they are read-only. others: meaningless
func (r *RegisterFgpfccrType) GetCm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFgpfccrFieldCmMask) >> RegisterFgpfccrFieldCmShift)
}

// SetCm Color mode These bits defines the color format of the foreground image. They can only be written when data transfers are disabled. Once the transfer has started, they are read-only. others: meaningless
func (r *RegisterFgpfccrType) SetCm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFgpfccrFieldCmMask)|(uint32(value)<<RegisterFgpfccrFieldCmShift))
}

const (
	RegisterFgpfccrFieldCcmShift = 4
	RegisterFgpfccrFieldCcmMask  = 0x10
)

// GetCcm CLUT color mode This bit defines the color format of the CLUT. It can only be written when the transfer is disabled. Once the CLUT transfer has started, this bit is read-only.
func (r *RegisterFgpfccrType) GetCcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFgpfccrFieldCcmMask) != 0
}

// SetCcm CLUT color mode This bit defines the color format of the CLUT. It can only be written when the transfer is disabled. Once the CLUT transfer has started, this bit is read-only.
func (r *RegisterFgpfccrType) SetCcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFgpfccrFieldCcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFgpfccrFieldCcmMask)
	}
}

const (
	RegisterFgpfccrFieldStartShift = 5
	RegisterFgpfccrFieldStartMask  = 0x20
)

// GetStart Start This bit can be set to start the automatic loading of the CLUT. It is automatically reset: ** at the end of the transfer ** when the transfer is aborted by the user application by setting the ABORT bit in DMA2D_CR ** when a transfer error occurs ** when the transfer has not started due to a configuration error or another transfer operation already ongoing (data transfer or automatic background CLUT transfer).
func (r *RegisterFgpfccrType) GetStart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFgpfccrFieldStartMask) != 0
}

// SetStart Start This bit can be set to start the automatic loading of the CLUT. It is automatically reset: ** at the end of the transfer ** when the transfer is aborted by the user application by setting the ABORT bit in DMA2D_CR ** when a transfer error occurs ** when the transfer has not started due to a configuration error or another transfer operation already ongoing (data transfer or automatic background CLUT transfer).
func (r *RegisterFgpfccrType) SetStart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFgpfccrFieldStartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFgpfccrFieldStartMask)
	}
}

const (
	RegisterFgpfccrFieldCsShift = 8
	RegisterFgpfccrFieldCsMask  = 0xff00
)

// GetCs CLUT size These bits define the size of the CLUT used for the foreground image. Once the CLUT transfer has started, this field is read-only. The number of CLUT entries is equal to CS[7:0] + 1.
func (r *RegisterFgpfccrType) GetCs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFgpfccrFieldCsMask) >> RegisterFgpfccrFieldCsShift)
}

// SetCs CLUT size These bits define the size of the CLUT used for the foreground image. Once the CLUT transfer has started, this field is read-only. The number of CLUT entries is equal to CS[7:0] + 1.
func (r *RegisterFgpfccrType) SetCs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFgpfccrFieldCsMask)|(uint32(value)<<RegisterFgpfccrFieldCsShift))
}

const (
	RegisterFgpfccrFieldAmShift = 16
	RegisterFgpfccrFieldAmMask  = 0x30000
)

// GetAm Alpha mode These bits select the alpha channel value to be used for the foreground image. They can only be written data the transfer are disabled. Once the transfer has started, they become read-only. other configurations are meaningless
func (r *RegisterFgpfccrType) GetAm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFgpfccrFieldAmMask) >> RegisterFgpfccrFieldAmShift)
}

// SetAm Alpha mode These bits select the alpha channel value to be used for the foreground image. They can only be written data the transfer are disabled. Once the transfer has started, they become read-only. other configurations are meaningless
func (r *RegisterFgpfccrType) SetAm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFgpfccrFieldAmMask)|(uint32(value)<<RegisterFgpfccrFieldAmShift))
}

const (
	RegisterFgpfccrFieldCssShift = 18
	RegisterFgpfccrFieldCssMask  = 0xc0000
)

// GetCss Chroma Sub-Sampling These bits define the chroma sub-sampling mode for YCbCr color mode. Once the transfer has started, these bits are read-only. others: meaningless
func (r *RegisterFgpfccrType) GetCss() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFgpfccrFieldCssMask) >> RegisterFgpfccrFieldCssShift)
}

// SetCss Chroma Sub-Sampling These bits define the chroma sub-sampling mode for YCbCr color mode. Once the transfer has started, these bits are read-only. others: meaningless
func (r *RegisterFgpfccrType) SetCss(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFgpfccrFieldCssMask)|(uint32(value)<<RegisterFgpfccrFieldCssShift))
}

const (
	RegisterFgpfccrFieldAiShift = 20
	RegisterFgpfccrFieldAiMask  = 0x100000
)

// GetAi Alpha Inverted This bit inverts the alpha value. Once the transfer has started, this bit is read-only.
func (r *RegisterFgpfccrType) GetAi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFgpfccrFieldAiMask) != 0
}

// SetAi Alpha Inverted This bit inverts the alpha value. Once the transfer has started, this bit is read-only.
func (r *RegisterFgpfccrType) SetAi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFgpfccrFieldAiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFgpfccrFieldAiMask)
	}
}

const (
	RegisterFgpfccrFieldRbsShift = 21
	RegisterFgpfccrFieldRbsMask  = 0x200000
)

// GetRbs Red Blue Swap This bit allows to swap the R &amp; B to support BGR or ABGR color formats. Once the transfer has started, this bit is read-only.
func (r *RegisterFgpfccrType) GetRbs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFgpfccrFieldRbsMask) != 0
}

// SetRbs Red Blue Swap This bit allows to swap the R &amp; B to support BGR or ABGR color formats. Once the transfer has started, this bit is read-only.
func (r *RegisterFgpfccrType) SetRbs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFgpfccrFieldRbsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFgpfccrFieldRbsMask)
	}
}

const (
	RegisterFgpfccrFieldAlphaShift = 24
	RegisterFgpfccrFieldAlphaMask  = 0xff000000
)

// GetAlpha Alpha value These bits define a fixed alpha channel value which can replace the original alpha value or be multiplied by the original alpha value according to the alpha mode selected through the AM[1:0] bits. These bits can only be written when data transfers are disabled. Once a transfer has started, they become read-only.
func (r *RegisterFgpfccrType) GetAlpha() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFgpfccrFieldAlphaMask) >> RegisterFgpfccrFieldAlphaShift)
}

// SetAlpha Alpha value These bits define a fixed alpha channel value which can replace the original alpha value or be multiplied by the original alpha value according to the alpha mode selected through the AM[1:0] bits. These bits can only be written when data transfers are disabled. Once a transfer has started, they become read-only.
func (r *RegisterFgpfccrType) SetAlpha(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFgpfccrFieldAlphaMask)|(uint32(value)<<RegisterFgpfccrFieldAlphaShift))
}

// RegisterFgcolrType DMA2D foreground color register
type RegisterFgcolrType uint32

func (r *RegisterFgcolrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterFgcolrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterFgcolrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterFgcolrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterFgcolrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterFgcolrFieldBlueShift = 0
	RegisterFgcolrFieldBlueMask  = 0xff
)

// GetBlue Blue Value These bits defines the blue value for the A4 or A8 mode of the foreground image. They can only be written when data transfers are disabled. Once the transfer has started, They are read-only.
func (r *RegisterFgcolrType) GetBlue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFgcolrFieldBlueMask) >> RegisterFgcolrFieldBlueShift)
}

// SetBlue Blue Value These bits defines the blue value for the A4 or A8 mode of the foreground image. They can only be written when data transfers are disabled. Once the transfer has started, They are read-only.
func (r *RegisterFgcolrType) SetBlue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFgcolrFieldBlueMask)|(uint32(value)<<RegisterFgcolrFieldBlueShift))
}

const (
	RegisterFgcolrFieldGreenShift = 8
	RegisterFgcolrFieldGreenMask  = 0xff00
)

// GetGreen Green Value These bits defines the green value for the A4 or A8 mode of the foreground image. They can only be written when data transfers are disabled. Once the transfer has started, They are read-only.
func (r *RegisterFgcolrType) GetGreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFgcolrFieldGreenMask) >> RegisterFgcolrFieldGreenShift)
}

// SetGreen Green Value These bits defines the green value for the A4 or A8 mode of the foreground image. They can only be written when data transfers are disabled. Once the transfer has started, They are read-only.
func (r *RegisterFgcolrType) SetGreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFgcolrFieldGreenMask)|(uint32(value)<<RegisterFgcolrFieldGreenShift))
}

const (
	RegisterFgcolrFieldRedShift = 16
	RegisterFgcolrFieldRedMask  = 0xff0000
)

// GetRed Red Value These bits defines the red value for the A4 or A8 mode of the foreground image. They can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterFgcolrType) GetRed() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFgcolrFieldRedMask) >> RegisterFgcolrFieldRedShift)
}

// SetRed Red Value These bits defines the red value for the A4 or A8 mode of the foreground image. They can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterFgcolrType) SetRed(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFgcolrFieldRedMask)|(uint32(value)<<RegisterFgcolrFieldRedShift))
}

// RegisterBgpfccrType DMA2D background PFC control register
type RegisterBgpfccrType uint32

func (r *RegisterBgpfccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBgpfccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBgpfccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBgpfccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBgpfccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBgpfccrFieldCmShift = 0
	RegisterBgpfccrFieldCmMask  = 0xf
)

// GetCm Color mode These bits define the color format of the foreground image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. others: meaningless
func (r *RegisterBgpfccrType) GetCm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBgpfccrFieldCmMask) >> RegisterBgpfccrFieldCmShift)
}

// SetCm Color mode These bits define the color format of the foreground image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. others: meaningless
func (r *RegisterBgpfccrType) SetCm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBgpfccrFieldCmMask)|(uint32(value)<<RegisterBgpfccrFieldCmShift))
}

const (
	RegisterBgpfccrFieldCcmShift = 4
	RegisterBgpfccrFieldCcmMask  = 0x10
)

// GetCcm CLUT Color mode These bits define the color format of the CLUT. This register can only be written when the transfer is disabled. Once the CLUT transfer has started, this bit is read-only.
func (r *RegisterBgpfccrType) GetCcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBgpfccrFieldCcmMask) != 0
}

// SetCcm CLUT Color mode These bits define the color format of the CLUT. This register can only be written when the transfer is disabled. Once the CLUT transfer has started, this bit is read-only.
func (r *RegisterBgpfccrType) SetCcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBgpfccrFieldCcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBgpfccrFieldCcmMask)
	}
}

const (
	RegisterBgpfccrFieldStartShift = 5
	RegisterBgpfccrFieldStartMask  = 0x20
)

// GetStart Start This bit is set to start the automatic loading of the CLUT. This bit is automatically reset: ** at the end of the transfer ** when the transfer is aborted by the user application by setting the ABORT bit in the DMA2D_CR ** when a transfer error occurs ** when the transfer has not started due to a configuration error or another transfer operation already on going (data transfer or automatic BackGround CLUT transfer).
func (r *RegisterBgpfccrType) GetStart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBgpfccrFieldStartMask) != 0
}

// SetStart Start This bit is set to start the automatic loading of the CLUT. This bit is automatically reset: ** at the end of the transfer ** when the transfer is aborted by the user application by setting the ABORT bit in the DMA2D_CR ** when a transfer error occurs ** when the transfer has not started due to a configuration error or another transfer operation already on going (data transfer or automatic BackGround CLUT transfer).
func (r *RegisterBgpfccrType) SetStart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBgpfccrFieldStartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBgpfccrFieldStartMask)
	}
}

const (
	RegisterBgpfccrFieldCsShift = 8
	RegisterBgpfccrFieldCsMask  = 0xff00
)

// GetCs CLUT size These bits define the size of the CLUT used for the BG. Once the CLUT transfer has started, this field is read-only. The number of CLUT entries is equal to CS[7:0] + 1.
func (r *RegisterBgpfccrType) GetCs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBgpfccrFieldCsMask) >> RegisterBgpfccrFieldCsShift)
}

// SetCs CLUT size These bits define the size of the CLUT used for the BG. Once the CLUT transfer has started, this field is read-only. The number of CLUT entries is equal to CS[7:0] + 1.
func (r *RegisterBgpfccrType) SetCs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBgpfccrFieldCsMask)|(uint32(value)<<RegisterBgpfccrFieldCsShift))
}

const (
	RegisterBgpfccrFieldAmShift = 16
	RegisterBgpfccrFieldAmMask  = 0x30000
)

// GetAm Alpha mode These bits define which alpha channel value to be used for the background image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. others: meaningless
func (r *RegisterBgpfccrType) GetAm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBgpfccrFieldAmMask) >> RegisterBgpfccrFieldAmShift)
}

// SetAm Alpha mode These bits define which alpha channel value to be used for the background image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. others: meaningless
func (r *RegisterBgpfccrType) SetAm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBgpfccrFieldAmMask)|(uint32(value)<<RegisterBgpfccrFieldAmShift))
}

const (
	RegisterBgpfccrFieldAiShift = 20
	RegisterBgpfccrFieldAiMask  = 0x100000
)

// GetAi Alpha Inverted This bit inverts the alpha value. Once the transfer has started, this bit is read-only.
func (r *RegisterBgpfccrType) GetAi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBgpfccrFieldAiMask) != 0
}

// SetAi Alpha Inverted This bit inverts the alpha value. Once the transfer has started, this bit is read-only.
func (r *RegisterBgpfccrType) SetAi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBgpfccrFieldAiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBgpfccrFieldAiMask)
	}
}

const (
	RegisterBgpfccrFieldRbsShift = 21
	RegisterBgpfccrFieldRbsMask  = 0x200000
)

// GetRbs Red Blue Swap This bit allows to swap the R &amp; B to support BGR or ABGR color formats. Once the transfer has started, this bit is read-only.
func (r *RegisterBgpfccrType) GetRbs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBgpfccrFieldRbsMask) != 0
}

// SetRbs Red Blue Swap This bit allows to swap the R &amp; B to support BGR or ABGR color formats. Once the transfer has started, this bit is read-only.
func (r *RegisterBgpfccrType) SetRbs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBgpfccrFieldRbsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBgpfccrFieldRbsMask)
	}
}

const (
	RegisterBgpfccrFieldAlphaShift = 24
	RegisterBgpfccrFieldAlphaMask  = 0xff000000
)

// GetAlpha Alpha value These bits define a fixed alpha channel value which can replace the original alpha value or be multiplied with the original alpha value according to the alpha mode selected with bits AM[1: 0]. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterBgpfccrType) GetAlpha() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBgpfccrFieldAlphaMask) >> RegisterBgpfccrFieldAlphaShift)
}

// SetAlpha Alpha value These bits define a fixed alpha channel value which can replace the original alpha value or be multiplied with the original alpha value according to the alpha mode selected with bits AM[1: 0]. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterBgpfccrType) SetAlpha(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBgpfccrFieldAlphaMask)|(uint32(value)<<RegisterBgpfccrFieldAlphaShift))
}

// RegisterBgcolrType DMA2D background color register
type RegisterBgcolrType uint32

func (r *RegisterBgcolrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBgcolrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBgcolrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBgcolrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBgcolrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBgcolrFieldBlueShift = 0
	RegisterBgcolrFieldBlueMask  = 0xff
)

// GetBlue Blue Value These bits define the blue value for the A4 or A8 mode of the background. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterBgcolrType) GetBlue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBgcolrFieldBlueMask) >> RegisterBgcolrFieldBlueShift)
}

// SetBlue Blue Value These bits define the blue value for the A4 or A8 mode of the background. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterBgcolrType) SetBlue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBgcolrFieldBlueMask)|(uint32(value)<<RegisterBgcolrFieldBlueShift))
}

const (
	RegisterBgcolrFieldGreenShift = 8
	RegisterBgcolrFieldGreenMask  = 0xff00
)

// GetGreen Green Value These bits define the green value for the A4 or A8 mode of the background. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterBgcolrType) GetGreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBgcolrFieldGreenMask) >> RegisterBgcolrFieldGreenShift)
}

// SetGreen Green Value These bits define the green value for the A4 or A8 mode of the background. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterBgcolrType) SetGreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBgcolrFieldGreenMask)|(uint32(value)<<RegisterBgcolrFieldGreenShift))
}

const (
	RegisterBgcolrFieldRedShift = 16
	RegisterBgcolrFieldRedMask  = 0xff0000
)

// GetRed Red Value These bits define the red value for the A4 or A8 mode of the background. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterBgcolrType) GetRed() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBgcolrFieldRedMask) >> RegisterBgcolrFieldRedShift)
}

// SetRed Red Value These bits define the red value for the A4 or A8 mode of the background. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterBgcolrType) SetRed(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBgcolrFieldRedMask)|(uint32(value)<<RegisterBgcolrFieldRedShift))
}

// RegisterFgcmarType DMA2D foreground CLUT memory address register
type RegisterFgcmarType uint32

func (r *RegisterFgcmarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterFgcmarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterFgcmarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterFgcmarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterFgcmarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterFgcmarFieldMaShift = 0
	RegisterFgcmarFieldMaMask  = 0xffffffff
)

// GetMa Memory Address Address of the data used for the CLUT address dedicated to the foreground image. This register can only be written when no transfer is ongoing. Once the CLUT transfer has started, this register is read-only. If the foreground CLUT format is 32-bit, the address must be 32-bit aligned.
func (r *RegisterFgcmarType) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFgcmarFieldMaMask) >> RegisterFgcmarFieldMaShift)
}

// SetMa Memory Address Address of the data used for the CLUT address dedicated to the foreground image. This register can only be written when no transfer is ongoing. Once the CLUT transfer has started, this register is read-only. If the foreground CLUT format is 32-bit, the address must be 32-bit aligned.
func (r *RegisterFgcmarType) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFgcmarFieldMaMask)|(uint32(value)<<RegisterFgcmarFieldMaShift))
}

// RegisterBgcmarType DMA2D background CLUT memory address register
type RegisterBgcmarType uint32

func (r *RegisterBgcmarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBgcmarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBgcmarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBgcmarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBgcmarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBgcmarFieldMaShift = 0
	RegisterBgcmarFieldMaMask  = 0xffffffff
)

// GetMa Memory address Address of the data used for the CLUT address dedicated to the background image. This register can only be written when no transfer is on going. Once the CLUT transfer has started, this register is read-only. If the background CLUT format is 32-bit, the address must be 32-bit aligned.
func (r *RegisterBgcmarType) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterBgcmarFieldMaMask) >> RegisterBgcmarFieldMaShift)
}

// SetMa Memory address Address of the data used for the CLUT address dedicated to the background image. This register can only be written when no transfer is on going. Once the CLUT transfer has started, this register is read-only. If the background CLUT format is 32-bit, the address must be 32-bit aligned.
func (r *RegisterBgcmarType) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBgcmarFieldMaMask)|(uint32(value)<<RegisterBgcmarFieldMaShift))
}

// RegisterOpfccrType DMA2D output PFC control register
type RegisterOpfccrType uint32

func (r *RegisterOpfccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOpfccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOpfccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOpfccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOpfccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOpfccrFieldCmShift = 0
	RegisterOpfccrFieldCmMask  = 0x7
)

// GetCm Color mode These bits define the color format of the output image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. others: meaningless
func (r *RegisterOpfccrType) GetCm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpfccrFieldCmMask) >> RegisterOpfccrFieldCmShift)
}

// SetCm Color mode These bits define the color format of the output image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. others: meaningless
func (r *RegisterOpfccrType) SetCm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpfccrFieldCmMask)|(uint32(value)<<RegisterOpfccrFieldCmShift))
}

const (
	RegisterOpfccrFieldAiShift = 20
	RegisterOpfccrFieldAiMask  = 0x100000
)

// GetAi Alpha Inverted This bit inverts the alpha value. Once the transfer has started, this bit is read-only.
func (r *RegisterOpfccrType) GetAi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpfccrFieldAiMask) != 0
}

// SetAi Alpha Inverted This bit inverts the alpha value. Once the transfer has started, this bit is read-only.
func (r *RegisterOpfccrType) SetAi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpfccrFieldAiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpfccrFieldAiMask)
	}
}

const (
	RegisterOpfccrFieldRbsShift = 21
	RegisterOpfccrFieldRbsMask  = 0x200000
)

// GetRbs Red Blue Swap This bit allows to swap the R &amp; B to support BGR or ABGR color formats. Once the transfer has started, this bit is read-only.
func (r *RegisterOpfccrType) GetRbs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpfccrFieldRbsMask) != 0
}

// SetRbs Red Blue Swap This bit allows to swap the R &amp; B to support BGR or ABGR color formats. Once the transfer has started, this bit is read-only.
func (r *RegisterOpfccrType) SetRbs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpfccrFieldRbsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpfccrFieldRbsMask)
	}
}

// RegisterOcolrType DMA2D output color register
type RegisterOcolrType uint32

func (r *RegisterOcolrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOcolrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOcolrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOcolrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOcolrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOcolrFieldBlueShift = 0
	RegisterOcolrFieldBlueMask  = 0xff
)

// GetBlue Blue Value These bits define the blue value of the output image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterOcolrType) GetBlue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOcolrFieldBlueMask) >> RegisterOcolrFieldBlueShift)
}

// SetBlue Blue Value These bits define the blue value of the output image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterOcolrType) SetBlue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOcolrFieldBlueMask)|(uint32(value)<<RegisterOcolrFieldBlueShift))
}

const (
	RegisterOcolrFieldGreenShift = 8
	RegisterOcolrFieldGreenMask  = 0xff00
)

// GetGreen Green Value These bits define the green value of the output image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterOcolrType) GetGreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOcolrFieldGreenMask) >> RegisterOcolrFieldGreenShift)
}

// SetGreen Green Value These bits define the green value of the output image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterOcolrType) SetGreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOcolrFieldGreenMask)|(uint32(value)<<RegisterOcolrFieldGreenShift))
}

const (
	RegisterOcolrFieldRedShift = 16
	RegisterOcolrFieldRedMask  = 0xff0000
)

// GetRed Red Value These bits define the red value of the output image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterOcolrType) GetRed() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOcolrFieldRedMask) >> RegisterOcolrFieldRedShift)
}

// SetRed Red Value These bits define the red value of the output image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterOcolrType) SetRed(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOcolrFieldRedMask)|(uint32(value)<<RegisterOcolrFieldRedShift))
}

const (
	RegisterOcolrFieldAlphaShift = 24
	RegisterOcolrFieldAlphaMask  = 0xff000000
)

// GetAlpha Alpha Channel Value These bits define the alpha channel of the output color. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterOcolrType) GetAlpha() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOcolrFieldAlphaMask) >> RegisterOcolrFieldAlphaShift)
}

// SetAlpha Alpha Channel Value These bits define the alpha channel of the output color. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterOcolrType) SetAlpha(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOcolrFieldAlphaMask)|(uint32(value)<<RegisterOcolrFieldAlphaShift))
}

// RegisterOmarType DMA2D output memory address register
type RegisterOmarType uint32

func (r *RegisterOmarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOmarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOmarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOmarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOmarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOmarFieldMaShift = 0
	RegisterOmarFieldMaMask  = 0xffffffff
)

// GetMa Memory Address Address of the data used for the output FIFO. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. The address alignment must match the image format selected e.g. a 32-bit per pixel format must be 32-bit aligned and a 16-bit per pixel format must be 16-bit aligned.
func (r *RegisterOmarType) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOmarFieldMaMask) >> RegisterOmarFieldMaShift)
}

// SetMa Memory Address Address of the data used for the output FIFO. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. The address alignment must match the image format selected e.g. a 32-bit per pixel format must be 32-bit aligned and a 16-bit per pixel format must be 16-bit aligned.
func (r *RegisterOmarType) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOmarFieldMaMask)|(uint32(value)<<RegisterOmarFieldMaShift))
}

// RegisterOorType DMA2D output offset register
type RegisterOorType uint32

func (r *RegisterOorType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOorType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOorType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOorType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOorType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOorFieldLoShift = 0
	RegisterOorFieldLoMask  = 0x3fff
)

// GetLo Line Offset Line offset used for the output (expressed in pixels). This value is used for the address generation. It is added at the end of each line to determine the starting address of the next line. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterOorType) GetLo() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOorFieldLoMask) >> RegisterOorFieldLoShift)
}

// SetLo Line Offset Line offset used for the output (expressed in pixels). This value is used for the address generation. It is added at the end of each line to determine the starting address of the next line. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterOorType) SetLo(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOorFieldLoMask)|(uint32(value)<<RegisterOorFieldLoShift))
}

// RegisterNlrType DMA2D number of line register
type RegisterNlrType uint32

func (r *RegisterNlrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterNlrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterNlrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterNlrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterNlrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterNlrFieldNlShift = 0
	RegisterNlrFieldNlMask  = 0xffff
)

// GetNl Number of lines Number of lines of the area to be transferred. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterNlrType) GetNl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterNlrFieldNlMask) >> RegisterNlrFieldNlShift)
}

// SetNl Number of lines Number of lines of the area to be transferred. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterNlrType) SetNl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterNlrFieldNlMask)|(uint32(value)<<RegisterNlrFieldNlShift))
}

const (
	RegisterNlrFieldPlShift = 16
	RegisterNlrFieldPlMask  = 0x3fff0000
)

// GetPl Pixel per lines Number of pixels per lines of the area to be transferred. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. If any of the input image format is 4-bit per pixel, pixel per lines must be even.
func (r *RegisterNlrType) GetPl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterNlrFieldPlMask) >> RegisterNlrFieldPlShift)
}

// SetPl Pixel per lines Number of pixels per lines of the area to be transferred. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. If any of the input image format is 4-bit per pixel, pixel per lines must be even.
func (r *RegisterNlrType) SetPl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterNlrFieldPlMask)|(uint32(value)<<RegisterNlrFieldPlShift))
}

// RegisterLwrType DMA2D line watermark register
type RegisterLwrType uint32

func (r *RegisterLwrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterLwrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterLwrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterLwrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterLwrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterLwrFieldLwShift = 0
	RegisterLwrFieldLwMask  = 0xffff
)

// GetLw Line watermark These bits allow to configure the line watermark for interrupt generation. An interrupt is raised when the last pixel of the watermarked line has been transferred. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterLwrType) GetLw() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterLwrFieldLwMask) >> RegisterLwrFieldLwShift)
}

// SetLw Line watermark These bits allow to configure the line watermark for interrupt generation. An interrupt is raised when the last pixel of the watermarked line has been transferred. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.
func (r *RegisterLwrType) SetLw(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterLwrFieldLwMask)|(uint32(value)<<RegisterLwrFieldLwShift))
}

// RegisterAmtcrType DMA2D AXI master timer configuration register
type RegisterAmtcrType uint32

func (r *RegisterAmtcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAmtcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAmtcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAmtcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAmtcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAmtcrFieldEnShift = 0
	RegisterAmtcrFieldEnMask  = 0x1
)

// GetEn Enable Enables the dead time functionality.
func (r *RegisterAmtcrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAmtcrFieldEnMask) != 0
}

// SetEn Enable Enables the dead time functionality.
func (r *RegisterAmtcrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAmtcrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAmtcrFieldEnMask)
	}
}

const (
	RegisterAmtcrFieldDtShift = 8
	RegisterAmtcrFieldDtMask  = 0xff00
)

// GetDt Dead Time Dead time value in the AXI clock cycle inserted between two consecutive accesses on the AXI master port. These bits represent the minimum guaranteed number of cycles between two consecutive AXI accesses.
func (r *RegisterAmtcrType) GetDt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAmtcrFieldDtMask) >> RegisterAmtcrFieldDtShift)
}

// SetDt Dead Time Dead time value in the AXI clock cycle inserted between two consecutive accesses on the AXI master port. These bits represent the minimum guaranteed number of cycles between two consecutive AXI accesses.
func (r *RegisterAmtcrType) SetDt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAmtcrFieldDtMask)|(uint32(value)<<RegisterAmtcrFieldDtShift))
}
