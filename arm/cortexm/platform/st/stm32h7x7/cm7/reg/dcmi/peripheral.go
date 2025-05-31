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
	Cr     registerCrType
	Sr     registerSrType
	Ris    registerRisType
	Ier    registerIerType
	Mis    registerMisType
	Icr    registerIcrType
	Escr   registerEscrType
	Esur   registerEsurType
	Cwstrt registerCwstrtType
	Cwsize registerCwsizeType
	Dr     registerDrType
}

// registerCrType control register 1
type registerCrType uint32

const (
	RegisterCrFieldCaptureShift = 0
	RegisterCrFieldCaptureMask  = 0x1
)

// GetCapture Capture enable
func (r *registerCrType) GetCapture() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCaptureMask) != 0
}

// SetCapture Capture enable
func (r *registerCrType) SetCapture(value bool) {
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
func (r *registerCrType) GetCm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCmMask) != 0
}

// SetCm Capture mode
func (r *registerCrType) SetCm(value bool) {
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
func (r *registerCrType) GetCrop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCropMask) != 0
}

// SetCrop Crop feature
func (r *registerCrType) SetCrop(value bool) {
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
func (r *registerCrType) GetJpeg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldJpegMask) != 0
}

// SetJpeg JPEG format
func (r *registerCrType) SetJpeg(value bool) {
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
func (r *registerCrType) GetEss() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEssMask) != 0
}

// SetEss Embedded synchronization select
func (r *registerCrType) SetEss(value bool) {
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
func (r *registerCrType) GetPckpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPckpolMask) != 0
}

// SetPckpol Pixel clock polarity
func (r *registerCrType) SetPckpol(value bool) {
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
func (r *registerCrType) GetHspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldHspolMask) != 0
}

// SetHspol Horizontal synchronization polarity
func (r *registerCrType) SetHspol(value bool) {
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
func (r *registerCrType) GetVspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldVspolMask) != 0
}

// SetVspol Vertical synchronization polarity
func (r *registerCrType) SetVspol(value bool) {
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
func (r *registerCrType) GetFcrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldFcrcMask) >> RegisterCrFieldFcrcShift)
}

// SetFcrc Frame capture rate control
func (r *registerCrType) SetFcrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldFcrcMask)|(uint32(value)<<RegisterCrFieldFcrcShift))
}

const (
	RegisterCrFieldEdmShift = 10
	RegisterCrFieldEdmMask  = 0xc00
)

// GetEdm Extended data mode
func (r *registerCrType) GetEdm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEdmMask) >> RegisterCrFieldEdmShift)
}

// SetEdm Extended data mode
func (r *registerCrType) SetEdm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEdmMask)|(uint32(value)<<RegisterCrFieldEdmShift))
}

const (
	RegisterCrFieldEnableShift = 14
	RegisterCrFieldEnableMask  = 0x4000
)

// GetEnable DCMI enable
func (r *registerCrType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEnableMask) != 0
}

// SetEnable DCMI enable
func (r *registerCrType) SetEnable(value bool) {
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
func (r *registerCrType) GetBsm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldBsmMask) >> RegisterCrFieldBsmShift)
}

// SetBsm Byte Select mode
func (r *registerCrType) SetBsm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldBsmMask)|(uint32(value)<<RegisterCrFieldBsmShift))
}

const (
	RegisterCrFieldOebsShift = 18
	RegisterCrFieldOebsMask  = 0x40000
)

// GetOebs Odd/Even Byte Select (Byte Select Start)
func (r *registerCrType) GetOebs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldOebsMask) != 0
}

// SetOebs Odd/Even Byte Select (Byte Select Start)
func (r *registerCrType) SetOebs(value bool) {
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
func (r *registerCrType) GetLsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLsmMask) != 0
}

// SetLsm Line Select mode
func (r *registerCrType) SetLsm(value bool) {
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
func (r *registerCrType) GetOels() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldOelsMask) != 0
}

// SetOels Odd/Even Line Select (Line Select Start)
func (r *registerCrType) SetOels(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldOelsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldOelsMask)
	}
}

// registerSrType status register
type registerSrType uint32

const (
	RegisterSrFieldHsyncShift = 0
	RegisterSrFieldHsyncMask  = 0x1
)

// GetHsync HSYNC
func (r *registerSrType) GetHsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldHsyncMask) != 0
}

// SetHsync HSYNC
func (r *registerSrType) SetHsync(value bool) {
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
func (r *registerSrType) GetVsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldVsyncMask) != 0
}

// SetVsync VSYNC
func (r *registerSrType) SetVsync(value bool) {
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
func (r *registerSrType) GetFne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldFneMask) != 0
}

// SetFne FIFO not empty
func (r *registerSrType) SetFne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldFneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldFneMask)
	}
}

// registerRisType raw interrupt status register
type registerRisType uint32

const (
	RegisterRisFieldFramerisShift = 0
	RegisterRisFieldFramerisMask  = 0x1
)

// GetFrameris Capture complete raw interrupt status
func (r *registerRisType) GetFrameris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldFramerisMask) != 0
}

// SetFrameris Capture complete raw interrupt status
func (r *registerRisType) SetFrameris(value bool) {
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
func (r *registerRisType) GetOvrris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldOvrrisMask) != 0
}

// SetOvrris Overrun raw interrupt status
func (r *registerRisType) SetOvrris(value bool) {
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
func (r *registerRisType) GetErrris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldErrrisMask) != 0
}

// SetErrris Synchronization error raw interrupt status
func (r *registerRisType) SetErrris(value bool) {
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
func (r *registerRisType) GetVsyncris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldVsyncrisMask) != 0
}

// SetVsyncris VSYNC raw interrupt status
func (r *registerRisType) SetVsyncris(value bool) {
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
func (r *registerRisType) GetLineris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldLinerisMask) != 0
}

// SetLineris Line raw interrupt status
func (r *registerRisType) SetLineris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisFieldLinerisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisFieldLinerisMask)
	}
}

// registerIerType interrupt enable register
type registerIerType uint32

const (
	RegisterIerFieldFrameieShift = 0
	RegisterIerFieldFrameieMask  = 0x1
)

// GetFrameie Capture complete interrupt enable
func (r *registerIerType) GetFrameie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldFrameieMask) != 0
}

// SetFrameie Capture complete interrupt enable
func (r *registerIerType) SetFrameie(value bool) {
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
func (r *registerIerType) GetOvrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldOvrieMask) != 0
}

// SetOvrie Overrun interrupt enable
func (r *registerIerType) SetOvrie(value bool) {
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
func (r *registerIerType) GetErrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldErrieMask) != 0
}

// SetErrie Synchronization error interrupt enable
func (r *registerIerType) SetErrie(value bool) {
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
func (r *registerIerType) GetVsyncie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldVsyncieMask) != 0
}

// SetVsyncie VSYNC interrupt enable
func (r *registerIerType) SetVsyncie(value bool) {
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
func (r *registerIerType) GetLineie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldLineieMask) != 0
}

// SetLineie Line interrupt enable
func (r *registerIerType) SetLineie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldLineieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldLineieMask)
	}
}

// registerMisType masked interrupt status register
type registerMisType uint32

const (
	RegisterMisFieldFramemisShift = 0
	RegisterMisFieldFramemisMask  = 0x1
)

// GetFramemis Capture complete masked interrupt status
func (r *registerMisType) GetFramemis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldFramemisMask) != 0
}

// SetFramemis Capture complete masked interrupt status
func (r *registerMisType) SetFramemis(value bool) {
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
func (r *registerMisType) GetOvrmis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldOvrmisMask) != 0
}

// SetOvrmis Overrun masked interrupt status
func (r *registerMisType) SetOvrmis(value bool) {
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
func (r *registerMisType) GetErrmis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldErrmisMask) != 0
}

// SetErrmis Synchronization error masked interrupt status
func (r *registerMisType) SetErrmis(value bool) {
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
func (r *registerMisType) GetVsyncmis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldVsyncmisMask) != 0
}

// SetVsyncmis VSYNC masked interrupt status
func (r *registerMisType) SetVsyncmis(value bool) {
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
func (r *registerMisType) GetLinemis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldLinemisMask) != 0
}

// SetLinemis Line masked interrupt status
func (r *registerMisType) SetLinemis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisFieldLinemisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisFieldLinemisMask)
	}
}

// registerIcrType interrupt clear register
type registerIcrType uint32

const (
	RegisterIcrFieldFrameiscShift = 0
	RegisterIcrFieldFrameiscMask  = 0x1
)

// GetFrameisc Capture complete interrupt status clear
func (r *registerIcrType) GetFrameisc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldFrameiscMask) != 0
}

// SetFrameisc Capture complete interrupt status clear
func (r *registerIcrType) SetFrameisc(value bool) {
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
func (r *registerIcrType) GetOvrisc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldOvriscMask) != 0
}

// SetOvrisc Overrun interrupt status clear
func (r *registerIcrType) SetOvrisc(value bool) {
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
func (r *registerIcrType) GetErrisc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldErriscMask) != 0
}

// SetErrisc Synchronization error interrupt status clear
func (r *registerIcrType) SetErrisc(value bool) {
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
func (r *registerIcrType) GetVsyncisc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldVsynciscMask) != 0
}

// SetVsyncisc Vertical synch interrupt status clear
func (r *registerIcrType) SetVsyncisc(value bool) {
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
func (r *registerIcrType) GetLineisc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldLineiscMask) != 0
}

// SetLineisc line interrupt status clear
func (r *registerIcrType) SetLineisc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldLineiscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldLineiscMask)
	}
}

// registerEscrType embedded synchronization code register
type registerEscrType uint32

const (
	RegisterEscrFieldFscShift = 0
	RegisterEscrFieldFscMask  = 0xff
)

// GetFsc Frame start delimiter code
func (r *registerEscrType) GetFsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEscrFieldFscMask) >> RegisterEscrFieldFscShift)
}

// SetFsc Frame start delimiter code
func (r *registerEscrType) SetFsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEscrFieldFscMask)|(uint32(value)<<RegisterEscrFieldFscShift))
}

const (
	RegisterEscrFieldLscShift = 8
	RegisterEscrFieldLscMask  = 0xff00
)

// GetLsc Line start delimiter code
func (r *registerEscrType) GetLsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEscrFieldLscMask) >> RegisterEscrFieldLscShift)
}

// SetLsc Line start delimiter code
func (r *registerEscrType) SetLsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEscrFieldLscMask)|(uint32(value)<<RegisterEscrFieldLscShift))
}

const (
	RegisterEscrFieldLecShift = 16
	RegisterEscrFieldLecMask  = 0xff0000
)

// GetLec Line end delimiter code
func (r *registerEscrType) GetLec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEscrFieldLecMask) >> RegisterEscrFieldLecShift)
}

// SetLec Line end delimiter code
func (r *registerEscrType) SetLec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEscrFieldLecMask)|(uint32(value)<<RegisterEscrFieldLecShift))
}

const (
	RegisterEscrFieldFecShift = 24
	RegisterEscrFieldFecMask  = 0xff000000
)

// GetFec Frame end delimiter code
func (r *registerEscrType) GetFec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEscrFieldFecMask) >> RegisterEscrFieldFecShift)
}

// SetFec Frame end delimiter code
func (r *registerEscrType) SetFec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEscrFieldFecMask)|(uint32(value)<<RegisterEscrFieldFecShift))
}

// registerEsurType embedded synchronization unmask register
type registerEsurType uint32

const (
	RegisterEsurFieldFsuShift = 0
	RegisterEsurFieldFsuMask  = 0xff
)

// GetFsu Frame start delimiter unmask
func (r *registerEsurType) GetFsu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEsurFieldFsuMask) >> RegisterEsurFieldFsuShift)
}

// SetFsu Frame start delimiter unmask
func (r *registerEsurType) SetFsu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEsurFieldFsuMask)|(uint32(value)<<RegisterEsurFieldFsuShift))
}

const (
	RegisterEsurFieldLsuShift = 8
	RegisterEsurFieldLsuMask  = 0xff00
)

// GetLsu Line start delimiter unmask
func (r *registerEsurType) GetLsu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEsurFieldLsuMask) >> RegisterEsurFieldLsuShift)
}

// SetLsu Line start delimiter unmask
func (r *registerEsurType) SetLsu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEsurFieldLsuMask)|(uint32(value)<<RegisterEsurFieldLsuShift))
}

const (
	RegisterEsurFieldLeuShift = 16
	RegisterEsurFieldLeuMask  = 0xff0000
)

// GetLeu Line end delimiter unmask
func (r *registerEsurType) GetLeu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEsurFieldLeuMask) >> RegisterEsurFieldLeuShift)
}

// SetLeu Line end delimiter unmask
func (r *registerEsurType) SetLeu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEsurFieldLeuMask)|(uint32(value)<<RegisterEsurFieldLeuShift))
}

const (
	RegisterEsurFieldFeuShift = 24
	RegisterEsurFieldFeuMask  = 0xff000000
)

// GetFeu Frame end delimiter unmask
func (r *registerEsurType) GetFeu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEsurFieldFeuMask) >> RegisterEsurFieldFeuShift)
}

// SetFeu Frame end delimiter unmask
func (r *registerEsurType) SetFeu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEsurFieldFeuMask)|(uint32(value)<<RegisterEsurFieldFeuShift))
}

// registerCwstrtType crop window start
type registerCwstrtType uint32

const (
	RegisterCwstrtFieldHoffcntShift = 0
	RegisterCwstrtFieldHoffcntMask  = 0x3fff
)

// GetHoffcnt Horizontal offset count
func (r *registerCwstrtType) GetHoffcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCwstrtFieldHoffcntMask) >> RegisterCwstrtFieldHoffcntShift)
}

// SetHoffcnt Horizontal offset count
func (r *registerCwstrtType) SetHoffcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwstrtFieldHoffcntMask)|(uint32(value)<<RegisterCwstrtFieldHoffcntShift))
}

const (
	RegisterCwstrtFieldVstShift = 16
	RegisterCwstrtFieldVstMask  = 0x1fff0000
)

// GetVst Vertical start line count
func (r *registerCwstrtType) GetVst() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCwstrtFieldVstMask) >> RegisterCwstrtFieldVstShift)
}

// SetVst Vertical start line count
func (r *registerCwstrtType) SetVst(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwstrtFieldVstMask)|(uint32(value)<<RegisterCwstrtFieldVstShift))
}

// registerCwsizeType crop window size
type registerCwsizeType uint32

const (
	RegisterCwsizeFieldCapcntShift = 0
	RegisterCwsizeFieldCapcntMask  = 0x3fff
)

// GetCapcnt Capture count
func (r *registerCwsizeType) GetCapcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCwsizeFieldCapcntMask) >> RegisterCwsizeFieldCapcntShift)
}

// SetCapcnt Capture count
func (r *registerCwsizeType) SetCapcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwsizeFieldCapcntMask)|(uint32(value)<<RegisterCwsizeFieldCapcntShift))
}

const (
	RegisterCwsizeFieldVlineShift = 16
	RegisterCwsizeFieldVlineMask  = 0x3fff0000
)

// GetVline Vertical line count
func (r *registerCwsizeType) GetVline() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCwsizeFieldVlineMask) >> RegisterCwsizeFieldVlineShift)
}

// SetVline Vertical line count
func (r *registerCwsizeType) SetVline(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwsizeFieldVlineMask)|(uint32(value)<<RegisterCwsizeFieldVlineShift))
}

// registerDrType data register
type registerDrType uint32

const (
	RegisterDrFieldByte0Shift = 0
	RegisterDrFieldByte0Mask  = 0xff
)

// GetByte0 Data byte 0
func (r *registerDrType) GetByte0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldByte0Mask) >> RegisterDrFieldByte0Shift)
}

// SetByte0 Data byte 0
func (r *registerDrType) SetByte0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldByte0Mask)|(uint32(value)<<RegisterDrFieldByte0Shift))
}

const (
	RegisterDrFieldByte1Shift = 8
	RegisterDrFieldByte1Mask  = 0xff00
)

// GetByte1 Data byte 1
func (r *registerDrType) GetByte1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldByte1Mask) >> RegisterDrFieldByte1Shift)
}

// SetByte1 Data byte 1
func (r *registerDrType) SetByte1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldByte1Mask)|(uint32(value)<<RegisterDrFieldByte1Shift))
}

const (
	RegisterDrFieldByte2Shift = 16
	RegisterDrFieldByte2Mask  = 0xff0000
)

// GetByte2 Data byte 2
func (r *registerDrType) GetByte2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldByte2Mask) >> RegisterDrFieldByte2Shift)
}

// SetByte2 Data byte 2
func (r *registerDrType) SetByte2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldByte2Mask)|(uint32(value)<<RegisterDrFieldByte2Shift))
}

const (
	RegisterDrFieldByte3Shift = 24
	RegisterDrFieldByte3Mask  = 0xff000000
)

// GetByte3 Data byte 3
func (r *registerDrType) GetByte3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldByte3Mask) >> RegisterDrFieldByte3Shift)
}

// SetByte3 Data byte 3
func (r *registerDrType) SetByte3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldByte3Mask)|(uint32(value)<<RegisterDrFieldByte3Shift))
}
