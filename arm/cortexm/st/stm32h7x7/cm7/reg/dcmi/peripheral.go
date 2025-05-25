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
	RegisterRisFieldFrame_risShift = 0
	RegisterRisFieldFrame_risMask  = 0x1
)

// GetFrame_ris Capture complete raw interrupt status
func (r *registerRisType) GetFrame_ris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldFrame_risMask) != 0
}

// SetFrame_ris Capture complete raw interrupt status
func (r *registerRisType) SetFrame_ris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisFieldFrame_risMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisFieldFrame_risMask)
	}
}

const (
	RegisterRisFieldOvr_risShift = 1
	RegisterRisFieldOvr_risMask  = 0x2
)

// GetOvr_ris Overrun raw interrupt status
func (r *registerRisType) GetOvr_ris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldOvr_risMask) != 0
}

// SetOvr_ris Overrun raw interrupt status
func (r *registerRisType) SetOvr_ris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisFieldOvr_risMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisFieldOvr_risMask)
	}
}

const (
	RegisterRisFieldErr_risShift = 2
	RegisterRisFieldErr_risMask  = 0x4
)

// GetErr_ris Synchronization error raw interrupt status
func (r *registerRisType) GetErr_ris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldErr_risMask) != 0
}

// SetErr_ris Synchronization error raw interrupt status
func (r *registerRisType) SetErr_ris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisFieldErr_risMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisFieldErr_risMask)
	}
}

const (
	RegisterRisFieldVsync_risShift = 3
	RegisterRisFieldVsync_risMask  = 0x8
)

// GetVsync_ris VSYNC raw interrupt status
func (r *registerRisType) GetVsync_ris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldVsync_risMask) != 0
}

// SetVsync_ris VSYNC raw interrupt status
func (r *registerRisType) SetVsync_ris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisFieldVsync_risMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisFieldVsync_risMask)
	}
}

const (
	RegisterRisFieldLine_risShift = 4
	RegisterRisFieldLine_risMask  = 0x10
)

// GetLine_ris Line raw interrupt status
func (r *registerRisType) GetLine_ris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisFieldLine_risMask) != 0
}

// SetLine_ris Line raw interrupt status
func (r *registerRisType) SetLine_ris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisFieldLine_risMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisFieldLine_risMask)
	}
}

// registerIerType interrupt enable register
type registerIerType uint32

const (
	RegisterIerFieldFrame_ieShift = 0
	RegisterIerFieldFrame_ieMask  = 0x1
)

// GetFrame_ie Capture complete interrupt enable
func (r *registerIerType) GetFrame_ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldFrame_ieMask) != 0
}

// SetFrame_ie Capture complete interrupt enable
func (r *registerIerType) SetFrame_ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldFrame_ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldFrame_ieMask)
	}
}

const (
	RegisterIerFieldOvr_ieShift = 1
	RegisterIerFieldOvr_ieMask  = 0x2
)

// GetOvr_ie Overrun interrupt enable
func (r *registerIerType) GetOvr_ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldOvr_ieMask) != 0
}

// SetOvr_ie Overrun interrupt enable
func (r *registerIerType) SetOvr_ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldOvr_ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldOvr_ieMask)
	}
}

const (
	RegisterIerFieldErr_ieShift = 2
	RegisterIerFieldErr_ieMask  = 0x4
)

// GetErr_ie Synchronization error interrupt enable
func (r *registerIerType) GetErr_ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldErr_ieMask) != 0
}

// SetErr_ie Synchronization error interrupt enable
func (r *registerIerType) SetErr_ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldErr_ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldErr_ieMask)
	}
}

const (
	RegisterIerFieldVsync_ieShift = 3
	RegisterIerFieldVsync_ieMask  = 0x8
)

// GetVsync_ie VSYNC interrupt enable
func (r *registerIerType) GetVsync_ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldVsync_ieMask) != 0
}

// SetVsync_ie VSYNC interrupt enable
func (r *registerIerType) SetVsync_ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldVsync_ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldVsync_ieMask)
	}
}

const (
	RegisterIerFieldLine_ieShift = 4
	RegisterIerFieldLine_ieMask  = 0x10
)

// GetLine_ie Line interrupt enable
func (r *registerIerType) GetLine_ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldLine_ieMask) != 0
}

// SetLine_ie Line interrupt enable
func (r *registerIerType) SetLine_ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldLine_ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldLine_ieMask)
	}
}

// registerMisType masked interrupt status register
type registerMisType uint32

const (
	RegisterMisFieldFrame_misShift = 0
	RegisterMisFieldFrame_misMask  = 0x1
)

// GetFrame_mis Capture complete masked interrupt status
func (r *registerMisType) GetFrame_mis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldFrame_misMask) != 0
}

// SetFrame_mis Capture complete masked interrupt status
func (r *registerMisType) SetFrame_mis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisFieldFrame_misMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisFieldFrame_misMask)
	}
}

const (
	RegisterMisFieldOvr_misShift = 1
	RegisterMisFieldOvr_misMask  = 0x2
)

// GetOvr_mis Overrun masked interrupt status
func (r *registerMisType) GetOvr_mis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldOvr_misMask) != 0
}

// SetOvr_mis Overrun masked interrupt status
func (r *registerMisType) SetOvr_mis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisFieldOvr_misMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisFieldOvr_misMask)
	}
}

const (
	RegisterMisFieldErr_misShift = 2
	RegisterMisFieldErr_misMask  = 0x4
)

// GetErr_mis Synchronization error masked interrupt status
func (r *registerMisType) GetErr_mis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldErr_misMask) != 0
}

// SetErr_mis Synchronization error masked interrupt status
func (r *registerMisType) SetErr_mis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisFieldErr_misMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisFieldErr_misMask)
	}
}

const (
	RegisterMisFieldVsync_misShift = 3
	RegisterMisFieldVsync_misMask  = 0x8
)

// GetVsync_mis VSYNC masked interrupt status
func (r *registerMisType) GetVsync_mis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldVsync_misMask) != 0
}

// SetVsync_mis VSYNC masked interrupt status
func (r *registerMisType) SetVsync_mis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisFieldVsync_misMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisFieldVsync_misMask)
	}
}

const (
	RegisterMisFieldLine_misShift = 4
	RegisterMisFieldLine_misMask  = 0x10
)

// GetLine_mis Line masked interrupt status
func (r *registerMisType) GetLine_mis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisFieldLine_misMask) != 0
}

// SetLine_mis Line masked interrupt status
func (r *registerMisType) SetLine_mis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisFieldLine_misMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisFieldLine_misMask)
	}
}

// registerIcrType interrupt clear register
type registerIcrType uint32

const (
	RegisterIcrFieldFrame_iscShift = 0
	RegisterIcrFieldFrame_iscMask  = 0x1
)

// GetFrame_isc Capture complete interrupt status clear
func (r *registerIcrType) GetFrame_isc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldFrame_iscMask) != 0
}

// SetFrame_isc Capture complete interrupt status clear
func (r *registerIcrType) SetFrame_isc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldFrame_iscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldFrame_iscMask)
	}
}

const (
	RegisterIcrFieldOvr_iscShift = 1
	RegisterIcrFieldOvr_iscMask  = 0x2
)

// GetOvr_isc Overrun interrupt status clear
func (r *registerIcrType) GetOvr_isc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldOvr_iscMask) != 0
}

// SetOvr_isc Overrun interrupt status clear
func (r *registerIcrType) SetOvr_isc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldOvr_iscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldOvr_iscMask)
	}
}

const (
	RegisterIcrFieldErr_iscShift = 2
	RegisterIcrFieldErr_iscMask  = 0x4
)

// GetErr_isc Synchronization error interrupt status clear
func (r *registerIcrType) GetErr_isc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldErr_iscMask) != 0
}

// SetErr_isc Synchronization error interrupt status clear
func (r *registerIcrType) SetErr_isc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldErr_iscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldErr_iscMask)
	}
}

const (
	RegisterIcrFieldVsync_iscShift = 3
	RegisterIcrFieldVsync_iscMask  = 0x8
)

// GetVsync_isc Vertical synch interrupt status clear
func (r *registerIcrType) GetVsync_isc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldVsync_iscMask) != 0
}

// SetVsync_isc Vertical synch interrupt status clear
func (r *registerIcrType) SetVsync_isc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldVsync_iscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldVsync_iscMask)
	}
}

const (
	RegisterIcrFieldLine_iscShift = 4
	RegisterIcrFieldLine_iscMask  = 0x10
)

// GetLine_isc line interrupt status clear
func (r *registerIcrType) GetLine_isc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldLine_iscMask) != 0
}

// SetLine_isc line interrupt status clear
func (r *registerIcrType) SetLine_isc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldLine_iscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldLine_iscMask)
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
