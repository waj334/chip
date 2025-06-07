//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package dcmi

import (
	"unsafe"
	"volatile"
)

var (
	Dcmi = (*_dcmi)(unsafe.Pointer(uintptr(0x48020000)))
)

type _dcmi struct {
	Cr     RegisterCrType
	Sr     RegisterSrType
	Ris    RegisterRisType
	Ier    RegisterIerType
	Mis    RegisterMisType
	Icr    RegisterIcrType
	Escr   RegisterEscrType
	Esur   RegisterEsurType
	Cwstrt RegisterCwstrtType
	Cwsize RegisterCwsizeType
	Dr     RegisterDrType
}

// RegisterCrType control register 1
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
	RegisterCrFieldCaptureShift = 0
	RegisterCrFieldCaptureMask  = 0x1
)

// GetCapture Capture enable
func (r *RegisterCrType) GetCapture() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCaptureMask) != 0
}

// SetCapture Capture enable
func (r *RegisterCrType) SetCapture(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCaptureMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCaptureMask)
	}
}

const (
	RegisterCrFieldCmShift = 1
	RegisterCrFieldCmMask  = 0x2
)

// GetCm Capture mode
func (r *RegisterCrType) GetCm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCmMask) != 0
}

// SetCm Capture mode
func (r *RegisterCrType) SetCm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCmMask)
	}
}

const (
	RegisterCrFieldCropShift = 2
	RegisterCrFieldCropMask  = 0x4
)

// GetCrop Crop feature
func (r *RegisterCrType) GetCrop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCropMask) != 0
}

// SetCrop Crop feature
func (r *RegisterCrType) SetCrop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCropMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCropMask)
	}
}

const (
	RegisterCrFieldJpegShift = 3
	RegisterCrFieldJpegMask  = 0x8
)

// GetJpeg JPEG format
func (r *RegisterCrType) GetJpeg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldJpegMask) != 0
}

// SetJpeg JPEG format
func (r *RegisterCrType) SetJpeg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldJpegMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldJpegMask)
	}
}

const (
	RegisterCrFieldEssShift = 4
	RegisterCrFieldEssMask  = 0x10
)

// GetEss Embedded synchronization select
func (r *RegisterCrType) GetEss() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEssMask) != 0
}

// SetEss Embedded synchronization select
func (r *RegisterCrType) SetEss(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEssMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEssMask)
	}
}

const (
	RegisterCrFieldPckpolShift = 5
	RegisterCrFieldPckpolMask  = 0x20
)

// GetPckpol Pixel clock polarity
func (r *RegisterCrType) GetPckpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPckpolMask) != 0
}

// SetPckpol Pixel clock polarity
func (r *RegisterCrType) SetPckpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldPckpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPckpolMask)
	}
}

const (
	RegisterCrFieldHspolShift = 6
	RegisterCrFieldHspolMask  = 0x40
)

// GetHspol Horizontal synchronization polarity
func (r *RegisterCrType) GetHspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldHspolMask) != 0
}

// SetHspol Horizontal synchronization polarity
func (r *RegisterCrType) SetHspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldHspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldHspolMask)
	}
}

const (
	RegisterCrFieldVspolShift = 7
	RegisterCrFieldVspolMask  = 0x80
)

// GetVspol Vertical synchronization polarity
func (r *RegisterCrType) GetVspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldVspolMask) != 0
}

// SetVspol Vertical synchronization polarity
func (r *RegisterCrType) SetVspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldVspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldVspolMask)
	}
}

const (
	RegisterCrFieldFcrcShift = 8
	RegisterCrFieldFcrcMask  = 0x300
)

// GetFcrc Frame capture rate control
func (r *RegisterCrType) GetFcrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldFcrcMask) >> RegisterCrFieldFcrcShift)
}

// SetFcrc Frame capture rate control
func (r *RegisterCrType) SetFcrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldFcrcMask)|(uint32(value)<<RegisterCrFieldFcrcShift))
}

const (
	RegisterCrFieldEdmShift = 10
	RegisterCrFieldEdmMask  = 0xc00
)

// GetEdm Extended data mode
func (r *RegisterCrType) GetEdm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEdmMask) >> RegisterCrFieldEdmShift)
}

// SetEdm Extended data mode
func (r *RegisterCrType) SetEdm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEdmMask)|(uint32(value)<<RegisterCrFieldEdmShift))
}

const (
	RegisterCrFieldEnableShift = 14
	RegisterCrFieldEnableMask  = 0x4000
)

// GetEnable DCMI enable
func (r *RegisterCrType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEnableMask) != 0
}

// SetEnable DCMI enable
func (r *RegisterCrType) SetEnable(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEnableMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEnableMask)
	}
}

const (
	RegisterCrFieldBsmShift = 16
	RegisterCrFieldBsmMask  = 0x30000
)

// GetBsm Byte Select mode
func (r *RegisterCrType) GetBsm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldBsmMask) >> RegisterCrFieldBsmShift)
}

// SetBsm Byte Select mode
func (r *RegisterCrType) SetBsm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldBsmMask)|(uint32(value)<<RegisterCrFieldBsmShift))
}

const (
	RegisterCrFieldOebsShift = 18
	RegisterCrFieldOebsMask  = 0x40000
)

// GetOebs Odd/Even Byte Select (Byte Select Start)
func (r *RegisterCrType) GetOebs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldOebsMask) != 0
}

// SetOebs Odd/Even Byte Select (Byte Select Start)
func (r *RegisterCrType) SetOebs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldOebsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldOebsMask)
	}
}

const (
	RegisterCrFieldLsmShift = 19
	RegisterCrFieldLsmMask  = 0x80000
)

// GetLsm Line Select mode
func (r *RegisterCrType) GetLsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLsmMask) != 0
}

// SetLsm Line Select mode
func (r *RegisterCrType) SetLsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldLsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldLsmMask)
	}
}

const (
	RegisterCrFieldOelsShift = 20
	RegisterCrFieldOelsMask  = 0x100000
)

// GetOels Odd/Even Line Select (Line Select Start)
func (r *RegisterCrType) GetOels() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldOelsMask) != 0
}

// SetOels Odd/Even Line Select (Line Select Start)
func (r *RegisterCrType) SetOels(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldOelsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldOelsMask)
	}
}

// RegisterSrType status register
type RegisterSrType uint32

func (r *RegisterSrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSrFieldHsyncShift = 0
	RegisterSrFieldHsyncMask  = 0x1
)

// GetHsync HSYNC
func (r *RegisterSrType) GetHsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldHsyncMask) != 0
}

// SetHsync HSYNC
func (r *RegisterSrType) SetHsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldHsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldHsyncMask)
	}
}

const (
	RegisterSrFieldVsyncShift = 1
	RegisterSrFieldVsyncMask  = 0x2
)

// GetVsync VSYNC
func (r *RegisterSrType) GetVsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldVsyncMask) != 0
}

// SetVsync VSYNC
func (r *RegisterSrType) SetVsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldVsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldVsyncMask)
	}
}

const (
	RegisterSrFieldFneShift = 2
	RegisterSrFieldFneMask  = 0x4
)

// GetFne FIFO not empty
func (r *RegisterSrType) GetFne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldFneMask) != 0
}

// SetFne FIFO not empty
func (r *RegisterSrType) SetFne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldFneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldFneMask)
	}
}

// RegisterRisType raw interrupt status register
type RegisterRisType uint32

func (r *RegisterRisType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRisType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRisType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRisType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRisType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRisFieldFramerisShift = 0
	RegisterRisFieldFramerisMask  = 0x1
)

// GetFrameris Capture complete raw interrupt status
func (r *RegisterRisType) GetFrameris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldFramerisMask) != 0
}

// SetFrameris Capture complete raw interrupt status
func (r *RegisterRisType) SetFrameris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisFieldFramerisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisFieldFramerisMask)
	}
}

const (
	RegisterRisFieldOvrrisShift = 1
	RegisterRisFieldOvrrisMask  = 0x2
)

// GetOvrris Overrun raw interrupt status
func (r *RegisterRisType) GetOvrris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldOvrrisMask) != 0
}

// SetOvrris Overrun raw interrupt status
func (r *RegisterRisType) SetOvrris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisFieldOvrrisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisFieldOvrrisMask)
	}
}

const (
	RegisterRisFieldErrrisShift = 2
	RegisterRisFieldErrrisMask  = 0x4
)

// GetErrris Synchronization error raw interrupt status
func (r *RegisterRisType) GetErrris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldErrrisMask) != 0
}

// SetErrris Synchronization error raw interrupt status
func (r *RegisterRisType) SetErrris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisFieldErrrisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisFieldErrrisMask)
	}
}

const (
	RegisterRisFieldVsyncrisShift = 3
	RegisterRisFieldVsyncrisMask  = 0x8
)

// GetVsyncris VSYNC raw interrupt status
func (r *RegisterRisType) GetVsyncris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldVsyncrisMask) != 0
}

// SetVsyncris VSYNC raw interrupt status
func (r *RegisterRisType) SetVsyncris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisFieldVsyncrisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisFieldVsyncrisMask)
	}
}

const (
	RegisterRisFieldLinerisShift = 4
	RegisterRisFieldLinerisMask  = 0x10
)

// GetLineris Line raw interrupt status
func (r *RegisterRisType) GetLineris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldLinerisMask) != 0
}

// SetLineris Line raw interrupt status
func (r *RegisterRisType) SetLineris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisFieldLinerisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisFieldLinerisMask)
	}
}

// RegisterIerType interrupt enable register
type RegisterIerType uint32

func (r *RegisterIerType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIerType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIerType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIerType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIerType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIerFieldFrameieShift = 0
	RegisterIerFieldFrameieMask  = 0x1
)

// GetFrameie Capture complete interrupt enable
func (r *RegisterIerType) GetFrameie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldFrameieMask) != 0
}

// SetFrameie Capture complete interrupt enable
func (r *RegisterIerType) SetFrameie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldFrameieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldFrameieMask)
	}
}

const (
	RegisterIerFieldOvrieShift = 1
	RegisterIerFieldOvrieMask  = 0x2
)

// GetOvrie Overrun interrupt enable
func (r *RegisterIerType) GetOvrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldOvrieMask) != 0
}

// SetOvrie Overrun interrupt enable
func (r *RegisterIerType) SetOvrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldOvrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldOvrieMask)
	}
}

const (
	RegisterIerFieldErrieShift = 2
	RegisterIerFieldErrieMask  = 0x4
)

// GetErrie Synchronization error interrupt enable
func (r *RegisterIerType) GetErrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldErrieMask) != 0
}

// SetErrie Synchronization error interrupt enable
func (r *RegisterIerType) SetErrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldErrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldErrieMask)
	}
}

const (
	RegisterIerFieldVsyncieShift = 3
	RegisterIerFieldVsyncieMask  = 0x8
)

// GetVsyncie VSYNC interrupt enable
func (r *RegisterIerType) GetVsyncie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldVsyncieMask) != 0
}

// SetVsyncie VSYNC interrupt enable
func (r *RegisterIerType) SetVsyncie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldVsyncieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldVsyncieMask)
	}
}

const (
	RegisterIerFieldLineieShift = 4
	RegisterIerFieldLineieMask  = 0x10
)

// GetLineie Line interrupt enable
func (r *RegisterIerType) GetLineie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldLineieMask) != 0
}

// SetLineie Line interrupt enable
func (r *RegisterIerType) SetLineie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldLineieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldLineieMask)
	}
}

// RegisterMisType masked interrupt status register
type RegisterMisType uint32

func (r *RegisterMisType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMisType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMisType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMisType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMisType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMisFieldFramemisShift = 0
	RegisterMisFieldFramemisMask  = 0x1
)

// GetFramemis Capture complete masked interrupt status
func (r *RegisterMisType) GetFramemis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldFramemisMask) != 0
}

// SetFramemis Capture complete masked interrupt status
func (r *RegisterMisType) SetFramemis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisFieldFramemisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisFieldFramemisMask)
	}
}

const (
	RegisterMisFieldOvrmisShift = 1
	RegisterMisFieldOvrmisMask  = 0x2
)

// GetOvrmis Overrun masked interrupt status
func (r *RegisterMisType) GetOvrmis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldOvrmisMask) != 0
}

// SetOvrmis Overrun masked interrupt status
func (r *RegisterMisType) SetOvrmis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisFieldOvrmisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisFieldOvrmisMask)
	}
}

const (
	RegisterMisFieldErrmisShift = 2
	RegisterMisFieldErrmisMask  = 0x4
)

// GetErrmis Synchronization error masked interrupt status
func (r *RegisterMisType) GetErrmis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldErrmisMask) != 0
}

// SetErrmis Synchronization error masked interrupt status
func (r *RegisterMisType) SetErrmis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisFieldErrmisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisFieldErrmisMask)
	}
}

const (
	RegisterMisFieldVsyncmisShift = 3
	RegisterMisFieldVsyncmisMask  = 0x8
)

// GetVsyncmis VSYNC masked interrupt status
func (r *RegisterMisType) GetVsyncmis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldVsyncmisMask) != 0
}

// SetVsyncmis VSYNC masked interrupt status
func (r *RegisterMisType) SetVsyncmis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisFieldVsyncmisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisFieldVsyncmisMask)
	}
}

const (
	RegisterMisFieldLinemisShift = 4
	RegisterMisFieldLinemisMask  = 0x10
)

// GetLinemis Line masked interrupt status
func (r *RegisterMisType) GetLinemis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldLinemisMask) != 0
}

// SetLinemis Line masked interrupt status
func (r *RegisterMisType) SetLinemis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisFieldLinemisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisFieldLinemisMask)
	}
}

// RegisterIcrType interrupt clear register
type RegisterIcrType uint32

func (r *RegisterIcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIcrFieldFrameiscShift = 0
	RegisterIcrFieldFrameiscMask  = 0x1
)

// GetFrameisc Capture complete interrupt status clear
func (r *RegisterIcrType) GetFrameisc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldFrameiscMask) != 0
}

// SetFrameisc Capture complete interrupt status clear
func (r *RegisterIcrType) SetFrameisc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldFrameiscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldFrameiscMask)
	}
}

const (
	RegisterIcrFieldOvriscShift = 1
	RegisterIcrFieldOvriscMask  = 0x2
)

// GetOvrisc Overrun interrupt status clear
func (r *RegisterIcrType) GetOvrisc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldOvriscMask) != 0
}

// SetOvrisc Overrun interrupt status clear
func (r *RegisterIcrType) SetOvrisc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldOvriscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldOvriscMask)
	}
}

const (
	RegisterIcrFieldErriscShift = 2
	RegisterIcrFieldErriscMask  = 0x4
)

// GetErrisc Synchronization error interrupt status clear
func (r *RegisterIcrType) GetErrisc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldErriscMask) != 0
}

// SetErrisc Synchronization error interrupt status clear
func (r *RegisterIcrType) SetErrisc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldErriscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldErriscMask)
	}
}

const (
	RegisterIcrFieldVsynciscShift = 3
	RegisterIcrFieldVsynciscMask  = 0x8
)

// GetVsyncisc Vertical synch interrupt status clear
func (r *RegisterIcrType) GetVsyncisc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldVsynciscMask) != 0
}

// SetVsyncisc Vertical synch interrupt status clear
func (r *RegisterIcrType) SetVsyncisc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldVsynciscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldVsynciscMask)
	}
}

const (
	RegisterIcrFieldLineiscShift = 4
	RegisterIcrFieldLineiscMask  = 0x10
)

// GetLineisc line interrupt status clear
func (r *RegisterIcrType) GetLineisc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldLineiscMask) != 0
}

// SetLineisc line interrupt status clear
func (r *RegisterIcrType) SetLineisc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldLineiscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldLineiscMask)
	}
}

// RegisterEscrType embedded synchronization code register
type RegisterEscrType uint32

func (r *RegisterEscrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterEscrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterEscrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterEscrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterEscrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterEscrFieldFscShift = 0
	RegisterEscrFieldFscMask  = 0xff
)

// GetFsc Frame start delimiter code
func (r *RegisterEscrType) GetFsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEscrFieldFscMask) >> RegisterEscrFieldFscShift)
}

// SetFsc Frame start delimiter code
func (r *RegisterEscrType) SetFsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEscrFieldFscMask)|(uint32(value)<<RegisterEscrFieldFscShift))
}

const (
	RegisterEscrFieldLscShift = 8
	RegisterEscrFieldLscMask  = 0xff00
)

// GetLsc Line start delimiter code
func (r *RegisterEscrType) GetLsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEscrFieldLscMask) >> RegisterEscrFieldLscShift)
}

// SetLsc Line start delimiter code
func (r *RegisterEscrType) SetLsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEscrFieldLscMask)|(uint32(value)<<RegisterEscrFieldLscShift))
}

const (
	RegisterEscrFieldLecShift = 16
	RegisterEscrFieldLecMask  = 0xff0000
)

// GetLec Line end delimiter code
func (r *RegisterEscrType) GetLec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEscrFieldLecMask) >> RegisterEscrFieldLecShift)
}

// SetLec Line end delimiter code
func (r *RegisterEscrType) SetLec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEscrFieldLecMask)|(uint32(value)<<RegisterEscrFieldLecShift))
}

const (
	RegisterEscrFieldFecShift = 24
	RegisterEscrFieldFecMask  = 0xff000000
)

// GetFec Frame end delimiter code
func (r *RegisterEscrType) GetFec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEscrFieldFecMask) >> RegisterEscrFieldFecShift)
}

// SetFec Frame end delimiter code
func (r *RegisterEscrType) SetFec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEscrFieldFecMask)|(uint32(value)<<RegisterEscrFieldFecShift))
}

// RegisterEsurType embedded synchronization unmask register
type RegisterEsurType uint32

func (r *RegisterEsurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterEsurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterEsurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterEsurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterEsurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterEsurFieldFsuShift = 0
	RegisterEsurFieldFsuMask  = 0xff
)

// GetFsu Frame start delimiter unmask
func (r *RegisterEsurType) GetFsu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEsurFieldFsuMask) >> RegisterEsurFieldFsuShift)
}

// SetFsu Frame start delimiter unmask
func (r *RegisterEsurType) SetFsu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEsurFieldFsuMask)|(uint32(value)<<RegisterEsurFieldFsuShift))
}

const (
	RegisterEsurFieldLsuShift = 8
	RegisterEsurFieldLsuMask  = 0xff00
)

// GetLsu Line start delimiter unmask
func (r *RegisterEsurType) GetLsu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEsurFieldLsuMask) >> RegisterEsurFieldLsuShift)
}

// SetLsu Line start delimiter unmask
func (r *RegisterEsurType) SetLsu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEsurFieldLsuMask)|(uint32(value)<<RegisterEsurFieldLsuShift))
}

const (
	RegisterEsurFieldLeuShift = 16
	RegisterEsurFieldLeuMask  = 0xff0000
)

// GetLeu Line end delimiter unmask
func (r *RegisterEsurType) GetLeu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEsurFieldLeuMask) >> RegisterEsurFieldLeuShift)
}

// SetLeu Line end delimiter unmask
func (r *RegisterEsurType) SetLeu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEsurFieldLeuMask)|(uint32(value)<<RegisterEsurFieldLeuShift))
}

const (
	RegisterEsurFieldFeuShift = 24
	RegisterEsurFieldFeuMask  = 0xff000000
)

// GetFeu Frame end delimiter unmask
func (r *RegisterEsurType) GetFeu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEsurFieldFeuMask) >> RegisterEsurFieldFeuShift)
}

// SetFeu Frame end delimiter unmask
func (r *RegisterEsurType) SetFeu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEsurFieldFeuMask)|(uint32(value)<<RegisterEsurFieldFeuShift))
}

// RegisterCwstrtType crop window start
type RegisterCwstrtType uint32

func (r *RegisterCwstrtType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCwstrtType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCwstrtType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCwstrtType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCwstrtType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCwstrtFieldHoffcntShift = 0
	RegisterCwstrtFieldHoffcntMask  = 0x3fff
)

// GetHoffcnt Horizontal offset count
func (r *RegisterCwstrtType) GetHoffcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCwstrtFieldHoffcntMask) >> RegisterCwstrtFieldHoffcntShift)
}

// SetHoffcnt Horizontal offset count
func (r *RegisterCwstrtType) SetHoffcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwstrtFieldHoffcntMask)|(uint32(value)<<RegisterCwstrtFieldHoffcntShift))
}

const (
	RegisterCwstrtFieldVstShift = 16
	RegisterCwstrtFieldVstMask  = 0x1fff0000
)

// GetVst Vertical start line count
func (r *RegisterCwstrtType) GetVst() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCwstrtFieldVstMask) >> RegisterCwstrtFieldVstShift)
}

// SetVst Vertical start line count
func (r *RegisterCwstrtType) SetVst(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwstrtFieldVstMask)|(uint32(value)<<RegisterCwstrtFieldVstShift))
}

// RegisterCwsizeType crop window size
type RegisterCwsizeType uint32

func (r *RegisterCwsizeType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCwsizeType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCwsizeType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCwsizeType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCwsizeType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCwsizeFieldCapcntShift = 0
	RegisterCwsizeFieldCapcntMask  = 0x3fff
)

// GetCapcnt Capture count
func (r *RegisterCwsizeType) GetCapcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCwsizeFieldCapcntMask) >> RegisterCwsizeFieldCapcntShift)
}

// SetCapcnt Capture count
func (r *RegisterCwsizeType) SetCapcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwsizeFieldCapcntMask)|(uint32(value)<<RegisterCwsizeFieldCapcntShift))
}

const (
	RegisterCwsizeFieldVlineShift = 16
	RegisterCwsizeFieldVlineMask  = 0x3fff0000
)

// GetVline Vertical line count
func (r *RegisterCwsizeType) GetVline() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCwsizeFieldVlineMask) >> RegisterCwsizeFieldVlineShift)
}

// SetVline Vertical line count
func (r *RegisterCwsizeType) SetVline(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwsizeFieldVlineMask)|(uint32(value)<<RegisterCwsizeFieldVlineShift))
}

// RegisterDrType data register
type RegisterDrType uint32

func (r *RegisterDrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDrFieldByte0Shift = 0
	RegisterDrFieldByte0Mask  = 0xff
)

// GetByte0 Data byte 0
func (r *RegisterDrType) GetByte0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldByte0Mask) >> RegisterDrFieldByte0Shift)
}

// SetByte0 Data byte 0
func (r *RegisterDrType) SetByte0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldByte0Mask)|(uint32(value)<<RegisterDrFieldByte0Shift))
}

const (
	RegisterDrFieldByte1Shift = 8
	RegisterDrFieldByte1Mask  = 0xff00
)

// GetByte1 Data byte 1
func (r *RegisterDrType) GetByte1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldByte1Mask) >> RegisterDrFieldByte1Shift)
}

// SetByte1 Data byte 1
func (r *RegisterDrType) SetByte1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldByte1Mask)|(uint32(value)<<RegisterDrFieldByte1Shift))
}

const (
	RegisterDrFieldByte2Shift = 16
	RegisterDrFieldByte2Mask  = 0xff0000
)

// GetByte2 Data byte 2
func (r *RegisterDrType) GetByte2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldByte2Mask) >> RegisterDrFieldByte2Shift)
}

// SetByte2 Data byte 2
func (r *RegisterDrType) SetByte2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldByte2Mask)|(uint32(value)<<RegisterDrFieldByte2Shift))
}

const (
	RegisterDrFieldByte3Shift = 24
	RegisterDrFieldByte3Mask  = 0xff000000
)

// GetByte3 Data byte 3
func (r *RegisterDrType) GetByte3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldByte3Mask) >> RegisterDrFieldByte3Shift)
}

// SetByte3 Data byte 3
func (r *RegisterDrType) SetByte3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldByte3Mask)|(uint32(value)<<RegisterDrFieldByte3Shift))
}
