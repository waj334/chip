//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package tim2_3_4_5

import (
	"unsafe"
	"volatile"
)

var (
	Tim2 = (*_tim2_3_4_5)(unsafe.Pointer(uintptr(0x40000000)))
	Tim3 = (*_tim2_3_4_5)(unsafe.Pointer(uintptr(0x40000400)))
	Tim4 = (*_tim2_3_4_5)(unsafe.Pointer(uintptr(0x40000800)))
	Tim5 = (*_tim2_3_4_5)(unsafe.Pointer(uintptr(0x40000c00)))

	Instances = [4]*_tim2_3_4_5{
		Tim2,
		Tim3,
		Tim4,
		Tim5,
	}
)

type _tim2_3_4_5 struct {
	Cr1   registerCr1Type
	Cr2   registerCr2Type
	Smcr  registerSmcrType
	Dier  registerDierType
	Sr    registerSrType
	Egr   registerEgrType
	Ccmr1 registerCcmr1Type
	Ccmr2 registerCcmr2Type
	Ccer  registerCcerType
	Cnt   registerCntType
	Psc   registerPscType
	Arr   registerArrType
	_     [4]byte
	Ccr1  registerCcr1Type
	Ccr2  registerCcr2Type
	Ccr3  registerCcr3Type
	Ccr4  registerCcr4Type
	_     [4]byte
	Dcr   registerDcrType
	Dmar  registerDmarType
	_     [16]byte
	Af1   registerAf1Type
	_     [4]byte
	Tisel registerTiselType
}

// registerCr1Type control register 1
type registerCr1Type uint32

const (
	RegisterCr1FieldCenShift = 0
	RegisterCr1FieldCenMask  = 0x1
)

// GetCen Counter enable
func (r *registerCr1Type) GetCen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCenMask) != 0
}

// SetCen Counter enable
func (r *registerCr1Type) SetCen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldCenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldCenMask)
	}
}

const (
	RegisterCr1FieldUdisShift = 1
	RegisterCr1FieldUdisMask  = 0x2
)

// GetUdis Update disable
func (r *registerCr1Type) GetUdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldUdisMask) != 0
}

// SetUdis Update disable
func (r *registerCr1Type) SetUdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldUdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldUdisMask)
	}
}

const (
	RegisterCr1FieldUrsShift = 2
	RegisterCr1FieldUrsMask  = 0x4
)

// GetUrs Update request source
func (r *registerCr1Type) GetUrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldUrsMask) != 0
}

// SetUrs Update request source
func (r *registerCr1Type) SetUrs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldUrsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldUrsMask)
	}
}

const (
	RegisterCr1FieldOpmShift = 3
	RegisterCr1FieldOpmMask  = 0x8
)

// GetOpm One-pulse mode
func (r *registerCr1Type) GetOpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldOpmMask) != 0
}

// SetOpm One-pulse mode
func (r *registerCr1Type) SetOpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldOpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldOpmMask)
	}
}

const (
	RegisterCr1FieldDirShift = 4
	RegisterCr1FieldDirMask  = 0x10
)

// GetDir Direction
func (r *registerCr1Type) GetDir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDirMask) != 0
}

// SetDir Direction
func (r *registerCr1Type) SetDir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDirMask)
	}
}

const (
	RegisterCr1FieldCmsShift = 5
	RegisterCr1FieldCmsMask  = 0x60
)

// GetCms Center-aligned mode selection
func (r *registerCr1Type) GetCms() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCmsMask) >> RegisterCr1FieldCmsShift)
}

// SetCms Center-aligned mode selection
func (r *registerCr1Type) SetCms(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldCmsMask)|(uint32(value)<<RegisterCr1FieldCmsShift))
}

const (
	RegisterCr1FieldArpeShift = 7
	RegisterCr1FieldArpeMask  = 0x80
)

// GetArpe Auto-reload preload enable
func (r *registerCr1Type) GetArpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldArpeMask) != 0
}

// SetArpe Auto-reload preload enable
func (r *registerCr1Type) SetArpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldArpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldArpeMask)
	}
}

const (
	RegisterCr1FieldCkdShift = 8
	RegisterCr1FieldCkdMask  = 0x300
)

// GetCkd Clock division
func (r *registerCr1Type) GetCkd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCkdMask) >> RegisterCr1FieldCkdShift)
}

// SetCkd Clock division
func (r *registerCr1Type) SetCkd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldCkdMask)|(uint32(value)<<RegisterCr1FieldCkdShift))
}

const (
	RegisterCr1FieldUifremapShift = 11
	RegisterCr1FieldUifremapMask  = 0x800
)

// GetUifremap UIF status bit remapping
func (r *registerCr1Type) GetUifremap() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldUifremapMask) != 0
}

// SetUifremap UIF status bit remapping
func (r *registerCr1Type) SetUifremap(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldUifremapMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldUifremapMask)
	}
}

// registerCr2Type control register 2
type registerCr2Type uint32

const (
	RegisterCr2FieldCcdsShift = 3
	RegisterCr2FieldCcdsMask  = 0x8
)

// GetCcds Capture/compare DMA selection
func (r *registerCr2Type) GetCcds() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCcdsMask) != 0
}

// SetCcds Capture/compare DMA selection
func (r *registerCr2Type) SetCcds(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldCcdsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldCcdsMask)
	}
}

const (
	RegisterCr2FieldMmsShift = 4
	RegisterCr2FieldMmsMask  = 0x70
)

// GetMms Master mode selection
func (r *registerCr2Type) GetMms() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldMmsMask) >> RegisterCr2FieldMmsShift)
}

// SetMms Master mode selection
func (r *registerCr2Type) SetMms(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldMmsMask)|(uint32(value)<<RegisterCr2FieldMmsShift))
}

const (
	RegisterCr2FieldTi1sShift = 7
	RegisterCr2FieldTi1sMask  = 0x80
)

// GetTi1s TI1 selection
func (r *registerCr2Type) GetTi1s() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTi1sMask) != 0
}

// SetTi1s TI1 selection
func (r *registerCr2Type) SetTi1s(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTi1sMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTi1sMask)
	}
}

// registerSmcrType slave mode control register
type registerSmcrType uint32

const (
	RegisterSmcrFieldSmsShift = 0
	RegisterSmcrFieldSmsMask  = 0x7
)

// GetSms Slave mode selection
func (r *registerSmcrType) GetSms() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldSmsMask) >> RegisterSmcrFieldSmsShift)
}

// SetSms Slave mode selection
func (r *registerSmcrType) SetSms(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldSmsMask)|(uint32(value)<<RegisterSmcrFieldSmsShift))
}

const (
	RegisterSmcrFieldTsShift = 4
	RegisterSmcrFieldTsMask  = 0x70
)

// GetTs Trigger selection
func (r *registerSmcrType) GetTs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldTsMask) >> RegisterSmcrFieldTsShift)
}

// SetTs Trigger selection
func (r *registerSmcrType) SetTs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldTsMask)|(uint32(value)<<RegisterSmcrFieldTsShift))
}

const (
	RegisterSmcrFieldMsmShift = 7
	RegisterSmcrFieldMsmMask  = 0x80
)

// GetMsm Master/Slave mode
func (r *registerSmcrType) GetMsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldMsmMask) != 0
}

// SetMsm Master/Slave mode
func (r *registerSmcrType) SetMsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSmcrFieldMsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldMsmMask)
	}
}

const (
	RegisterSmcrFieldEtfShift = 8
	RegisterSmcrFieldEtfMask  = 0xf00
)

// GetEtf External trigger filter
func (r *registerSmcrType) GetEtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldEtfMask) >> RegisterSmcrFieldEtfShift)
}

// SetEtf External trigger filter
func (r *registerSmcrType) SetEtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldEtfMask)|(uint32(value)<<RegisterSmcrFieldEtfShift))
}

const (
	RegisterSmcrFieldEtpsShift = 12
	RegisterSmcrFieldEtpsMask  = 0x3000
)

// GetEtps External trigger prescaler
func (r *registerSmcrType) GetEtps() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldEtpsMask) >> RegisterSmcrFieldEtpsShift)
}

// SetEtps External trigger prescaler
func (r *registerSmcrType) SetEtps(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldEtpsMask)|(uint32(value)<<RegisterSmcrFieldEtpsShift))
}

const (
	RegisterSmcrFieldEceShift = 14
	RegisterSmcrFieldEceMask  = 0x4000
)

// GetEce External clock enable
func (r *registerSmcrType) GetEce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldEceMask) != 0
}

// SetEce External clock enable
func (r *registerSmcrType) SetEce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSmcrFieldEceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldEceMask)
	}
}

const (
	RegisterSmcrFieldEtpShift = 15
	RegisterSmcrFieldEtpMask  = 0x8000
)

// GetEtp External trigger polarity
func (r *registerSmcrType) GetEtp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldEtpMask) != 0
}

// SetEtp External trigger polarity
func (r *registerSmcrType) SetEtp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSmcrFieldEtpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldEtpMask)
	}
}

const (
	RegisterSmcrFieldSms3Shift = 16
	RegisterSmcrFieldSms3Mask  = 0x10000
)

// GetSms3 Slave mode selection - bit 3
func (r *registerSmcrType) GetSms3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldSms3Mask) != 0
}

// SetSms3 Slave mode selection - bit 3
func (r *registerSmcrType) SetSms3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSmcrFieldSms3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldSms3Mask)
	}
}

const (
	RegisterSmcrFieldTs43Shift = 20
	RegisterSmcrFieldTs43Mask  = 0x300000
)

// GetTs43 Trigger selection
func (r *registerSmcrType) GetTs43() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldTs43Mask) >> RegisterSmcrFieldTs43Shift)
}

// SetTs43 Trigger selection
func (r *registerSmcrType) SetTs43(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldTs43Mask)|(uint32(value)<<RegisterSmcrFieldTs43Shift))
}

// registerDierType DMA/Interrupt enable register
type registerDierType uint32

const (
	RegisterDierFieldUieShift = 0
	RegisterDierFieldUieMask  = 0x1
)

// GetUie Update interrupt enable
func (r *registerDierType) GetUie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldUieMask) != 0
}

// SetUie Update interrupt enable
func (r *registerDierType) SetUie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldUieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldUieMask)
	}
}

const (
	RegisterDierFieldCc1ieShift = 1
	RegisterDierFieldCc1ieMask  = 0x2
)

// GetCc1ie Capture/Compare 1 interrupt enable
func (r *registerDierType) GetCc1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc1ieMask) != 0
}

// SetCc1ie Capture/Compare 1 interrupt enable
func (r *registerDierType) SetCc1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldCc1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldCc1ieMask)
	}
}

const (
	RegisterDierFieldCc2ieShift = 2
	RegisterDierFieldCc2ieMask  = 0x4
)

// GetCc2ie Capture/Compare 2 interrupt enable
func (r *registerDierType) GetCc2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc2ieMask) != 0
}

// SetCc2ie Capture/Compare 2 interrupt enable
func (r *registerDierType) SetCc2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldCc2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldCc2ieMask)
	}
}

const (
	RegisterDierFieldCc3ieShift = 3
	RegisterDierFieldCc3ieMask  = 0x8
)

// GetCc3ie Capture/Compare 3 interrupt enable
func (r *registerDierType) GetCc3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc3ieMask) != 0
}

// SetCc3ie Capture/Compare 3 interrupt enable
func (r *registerDierType) SetCc3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldCc3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldCc3ieMask)
	}
}

const (
	RegisterDierFieldCc4ieShift = 4
	RegisterDierFieldCc4ieMask  = 0x10
)

// GetCc4ie Capture/Compare 4 interrupt enable
func (r *registerDierType) GetCc4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc4ieMask) != 0
}

// SetCc4ie Capture/Compare 4 interrupt enable
func (r *registerDierType) SetCc4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldCc4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldCc4ieMask)
	}
}

const (
	RegisterDierFieldTieShift = 6
	RegisterDierFieldTieMask  = 0x40
)

// GetTie Trigger interrupt enable
func (r *registerDierType) GetTie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldTieMask) != 0
}

// SetTie Trigger interrupt enable
func (r *registerDierType) SetTie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldTieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldTieMask)
	}
}

const (
	RegisterDierFieldUdeShift = 8
	RegisterDierFieldUdeMask  = 0x100
)

// GetUde Update DMA request enable
func (r *registerDierType) GetUde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldUdeMask) != 0
}

// SetUde Update DMA request enable
func (r *registerDierType) SetUde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldUdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldUdeMask)
	}
}

const (
	RegisterDierFieldCc1deShift = 9
	RegisterDierFieldCc1deMask  = 0x200
)

// GetCc1de Capture/Compare 1 DMA request enable
func (r *registerDierType) GetCc1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc1deMask) != 0
}

// SetCc1de Capture/Compare 1 DMA request enable
func (r *registerDierType) SetCc1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldCc1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldCc1deMask)
	}
}

const (
	RegisterDierFieldCc2deShift = 10
	RegisterDierFieldCc2deMask  = 0x400
)

// GetCc2de Capture/Compare 2 DMA request enable
func (r *registerDierType) GetCc2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc2deMask) != 0
}

// SetCc2de Capture/Compare 2 DMA request enable
func (r *registerDierType) SetCc2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldCc2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldCc2deMask)
	}
}

const (
	RegisterDierFieldCc3deShift = 11
	RegisterDierFieldCc3deMask  = 0x800
)

// GetCc3de Capture/Compare 3 DMA request enable
func (r *registerDierType) GetCc3de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc3deMask) != 0
}

// SetCc3de Capture/Compare 3 DMA request enable
func (r *registerDierType) SetCc3de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldCc3deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldCc3deMask)
	}
}

const (
	RegisterDierFieldCc4deShift = 12
	RegisterDierFieldCc4deMask  = 0x1000
)

// GetCc4de Capture/Compare 4 DMA request enable
func (r *registerDierType) GetCc4de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc4deMask) != 0
}

// SetCc4de Capture/Compare 4 DMA request enable
func (r *registerDierType) SetCc4de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldCc4deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldCc4deMask)
	}
}

const (
	RegisterDierFieldTdeShift = 14
	RegisterDierFieldTdeMask  = 0x4000
)

// GetTde Trigger DMA request enable
func (r *registerDierType) GetTde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldTdeMask) != 0
}

// SetTde Trigger DMA request enable
func (r *registerDierType) SetTde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldTdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldTdeMask)
	}
}

// registerSrType status register
type registerSrType uint32

const (
	RegisterSrFieldUifShift = 0
	RegisterSrFieldUifMask  = 0x1
)

// GetUif Update interrupt flag
func (r *registerSrType) GetUif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldUifMask) != 0
}

// SetUif Update interrupt flag
func (r *registerSrType) SetUif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldUifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldUifMask)
	}
}

const (
	RegisterSrFieldCc1ifShift = 1
	RegisterSrFieldCc1ifMask  = 0x2
)

// GetCc1if Capture/compare 1 interrupt flag
func (r *registerSrType) GetCc1if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc1ifMask) != 0
}

// SetCc1if Capture/compare 1 interrupt flag
func (r *registerSrType) SetCc1if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCc1ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCc1ifMask)
	}
}

const (
	RegisterSrFieldCc2ifShift = 2
	RegisterSrFieldCc2ifMask  = 0x4
)

// GetCc2if Capture/Compare 2 interrupt flag
func (r *registerSrType) GetCc2if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc2ifMask) != 0
}

// SetCc2if Capture/Compare 2 interrupt flag
func (r *registerSrType) SetCc2if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCc2ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCc2ifMask)
	}
}

const (
	RegisterSrFieldCc3ifShift = 3
	RegisterSrFieldCc3ifMask  = 0x8
)

// GetCc3if Capture/Compare 3 interrupt flag
func (r *registerSrType) GetCc3if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc3ifMask) != 0
}

// SetCc3if Capture/Compare 3 interrupt flag
func (r *registerSrType) SetCc3if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCc3ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCc3ifMask)
	}
}

const (
	RegisterSrFieldCc4ifShift = 4
	RegisterSrFieldCc4ifMask  = 0x10
)

// GetCc4if Capture/Compare 4 interrupt flag
func (r *registerSrType) GetCc4if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc4ifMask) != 0
}

// SetCc4if Capture/Compare 4 interrupt flag
func (r *registerSrType) SetCc4if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCc4ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCc4ifMask)
	}
}

const (
	RegisterSrFieldTifShift = 6
	RegisterSrFieldTifMask  = 0x40
)

// GetTif Trigger interrupt flag
func (r *registerSrType) GetTif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTifMask) != 0
}

// SetTif Trigger interrupt flag
func (r *registerSrType) SetTif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldTifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldTifMask)
	}
}

const (
	RegisterSrFieldCc1ofShift = 9
	RegisterSrFieldCc1ofMask  = 0x200
)

// GetCc1of Capture/Compare 1 overcapture flag
func (r *registerSrType) GetCc1of() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc1ofMask) != 0
}

// SetCc1of Capture/Compare 1 overcapture flag
func (r *registerSrType) SetCc1of(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCc1ofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCc1ofMask)
	}
}

const (
	RegisterSrFieldCc2ofShift = 10
	RegisterSrFieldCc2ofMask  = 0x400
)

// GetCc2of Capture/compare 2 overcapture flag
func (r *registerSrType) GetCc2of() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc2ofMask) != 0
}

// SetCc2of Capture/compare 2 overcapture flag
func (r *registerSrType) SetCc2of(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCc2ofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCc2ofMask)
	}
}

const (
	RegisterSrFieldCc3ofShift = 11
	RegisterSrFieldCc3ofMask  = 0x800
)

// GetCc3of Capture/Compare 3 overcapture flag
func (r *registerSrType) GetCc3of() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc3ofMask) != 0
}

// SetCc3of Capture/Compare 3 overcapture flag
func (r *registerSrType) SetCc3of(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCc3ofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCc3ofMask)
	}
}

const (
	RegisterSrFieldCc4ofShift = 12
	RegisterSrFieldCc4ofMask  = 0x1000
)

// GetCc4of Capture/Compare 4 overcapture flag
func (r *registerSrType) GetCc4of() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc4ofMask) != 0
}

// SetCc4of Capture/Compare 4 overcapture flag
func (r *registerSrType) SetCc4of(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCc4ofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCc4ofMask)
	}
}

// registerEgrType event generation register
type registerEgrType uint32

const (
	RegisterEgrFieldUgShift = 0
	RegisterEgrFieldUgMask  = 0x1
)

// GetUg Update generation
func (r *registerEgrType) GetUg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldUgMask) != 0
}

// SetUg Update generation
func (r *registerEgrType) SetUg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldUgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldUgMask)
	}
}

const (
	RegisterEgrFieldCc1gShift = 1
	RegisterEgrFieldCc1gMask  = 0x2
)

// GetCc1g Capture/compare 1 generation
func (r *registerEgrType) GetCc1g() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldCc1gMask) != 0
}

// SetCc1g Capture/compare 1 generation
func (r *registerEgrType) SetCc1g(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldCc1gMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldCc1gMask)
	}
}

const (
	RegisterEgrFieldCc2gShift = 2
	RegisterEgrFieldCc2gMask  = 0x4
)

// GetCc2g Capture/compare 2 generation
func (r *registerEgrType) GetCc2g() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldCc2gMask) != 0
}

// SetCc2g Capture/compare 2 generation
func (r *registerEgrType) SetCc2g(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldCc2gMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldCc2gMask)
	}
}

const (
	RegisterEgrFieldCc3gShift = 3
	RegisterEgrFieldCc3gMask  = 0x8
)

// GetCc3g Capture/compare 3 generation
func (r *registerEgrType) GetCc3g() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldCc3gMask) != 0
}

// SetCc3g Capture/compare 3 generation
func (r *registerEgrType) SetCc3g(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldCc3gMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldCc3gMask)
	}
}

const (
	RegisterEgrFieldCc4gShift = 4
	RegisterEgrFieldCc4gMask  = 0x10
)

// GetCc4g Capture/compare 4 generation
func (r *registerEgrType) GetCc4g() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldCc4gMask) != 0
}

// SetCc4g Capture/compare 4 generation
func (r *registerEgrType) SetCc4g(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldCc4gMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldCc4gMask)
	}
}

const (
	RegisterEgrFieldTgShift = 6
	RegisterEgrFieldTgMask  = 0x40
)

// GetTg Trigger generation
func (r *registerEgrType) GetTg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldTgMask) != 0
}

// SetTg Trigger generation
func (r *registerEgrType) SetTg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldTgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldTgMask)
	}
}

// registerCcmr1Type capture/compare mode register 1
type registerCcmr1Type uint32

const (
	RegisterCcmr1FieldCc1sShift = 0
	RegisterCcmr1FieldCc1sMask  = 0x3
)

// GetCc1s CC1S
func (r *registerCcmr1Type) GetCc1s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldCc1sMask) >> RegisterCcmr1FieldCc1sShift)
}

// SetCc1s CC1S
func (r *registerCcmr1Type) SetCc1s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldCc1sMask)|(uint32(value)<<RegisterCcmr1FieldCc1sShift))
}

const (
	RegisterCcmr1FieldOc1feShift = 2
	RegisterCcmr1FieldOc1feMask  = 0x4
)

// GetOc1fe OC1FE
func (r *registerCcmr1Type) GetOc1fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc1feMask) != 0
}

// SetOc1fe OC1FE
func (r *registerCcmr1Type) SetOc1fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1FieldOc1feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldOc1feMask)
	}
}

const (
	RegisterCcmr1FieldOc1peShift = 3
	RegisterCcmr1FieldOc1peMask  = 0x8
)

// GetOc1pe OC1PE
func (r *registerCcmr1Type) GetOc1pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc1peMask) != 0
}

// SetOc1pe OC1PE
func (r *registerCcmr1Type) SetOc1pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1FieldOc1peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldOc1peMask)
	}
}

const (
	RegisterCcmr1FieldOc1mShift = 4
	RegisterCcmr1FieldOc1mMask  = 0x70
)

// GetOc1m OC1M
func (r *registerCcmr1Type) GetOc1m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc1mMask) >> RegisterCcmr1FieldOc1mShift)
}

// SetOc1m OC1M
func (r *registerCcmr1Type) SetOc1m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldOc1mMask)|(uint32(value)<<RegisterCcmr1FieldOc1mShift))
}

const (
	RegisterCcmr1FieldOc1ceShift = 7
	RegisterCcmr1FieldOc1ceMask  = 0x80
)

// GetOc1ce OC1CE
func (r *registerCcmr1Type) GetOc1ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc1ceMask) != 0
}

// SetOc1ce OC1CE
func (r *registerCcmr1Type) SetOc1ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1FieldOc1ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldOc1ceMask)
	}
}

const (
	RegisterCcmr1FieldCc2sShift = 8
	RegisterCcmr1FieldCc2sMask  = 0x300
)

// GetCc2s CC2S
func (r *registerCcmr1Type) GetCc2s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldCc2sMask) >> RegisterCcmr1FieldCc2sShift)
}

// SetCc2s CC2S
func (r *registerCcmr1Type) SetCc2s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldCc2sMask)|(uint32(value)<<RegisterCcmr1FieldCc2sShift))
}

const (
	RegisterCcmr1FieldOc2feShift = 10
	RegisterCcmr1FieldOc2feMask  = 0x400
)

// GetOc2fe OC2FE
func (r *registerCcmr1Type) GetOc2fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc2feMask) != 0
}

// SetOc2fe OC2FE
func (r *registerCcmr1Type) SetOc2fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1FieldOc2feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldOc2feMask)
	}
}

const (
	RegisterCcmr1FieldOc2peShift = 11
	RegisterCcmr1FieldOc2peMask  = 0x800
)

// GetOc2pe OC2PE
func (r *registerCcmr1Type) GetOc2pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc2peMask) != 0
}

// SetOc2pe OC2PE
func (r *registerCcmr1Type) SetOc2pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1FieldOc2peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldOc2peMask)
	}
}

const (
	RegisterCcmr1FieldOc2mShift = 12
	RegisterCcmr1FieldOc2mMask  = 0x7000
)

// GetOc2m OC2M
func (r *registerCcmr1Type) GetOc2m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc2mMask) >> RegisterCcmr1FieldOc2mShift)
}

// SetOc2m OC2M
func (r *registerCcmr1Type) SetOc2m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldOc2mMask)|(uint32(value)<<RegisterCcmr1FieldOc2mShift))
}

const (
	RegisterCcmr1FieldOc2ceShift = 15
	RegisterCcmr1FieldOc2ceMask  = 0x8000
)

// GetOc2ce OC2CE
func (r *registerCcmr1Type) GetOc2ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc2ceMask) != 0
}

// SetOc2ce OC2CE
func (r *registerCcmr1Type) SetOc2ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1FieldOc2ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldOc2ceMask)
	}
}

const (
	RegisterCcmr1FieldOc1m3Shift = 16
	RegisterCcmr1FieldOc1m3Mask  = 0x10000
)

// GetOc1m3 Output Compare 1 mode - bit 3
func (r *registerCcmr1Type) GetOc1m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc1m3Mask) != 0
}

// SetOc1m3 Output Compare 1 mode - bit 3
func (r *registerCcmr1Type) SetOc1m3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1FieldOc1m3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldOc1m3Mask)
	}
}

const (
	RegisterCcmr1FieldOc2m3Shift = 24
	RegisterCcmr1FieldOc2m3Mask  = 0x1000000
)

// GetOc2m3 Output Compare 2 mode - bit 3
func (r *registerCcmr1Type) GetOc2m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc2m3Mask) != 0
}

// SetOc2m3 Output Compare 2 mode - bit 3
func (r *registerCcmr1Type) SetOc2m3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1FieldOc2m3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldOc2m3Mask)
	}
}

const (
	RegisterCcmr1FieldIcpcsShift = 2
	RegisterCcmr1FieldIcpcsMask  = 0xc
)

// GetIcpcs Input capture 1 prescaler
func (r *registerCcmr1Type) GetIcpcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldIcpcsMask) >> RegisterCcmr1FieldIcpcsShift)
}

// SetIcpcs Input capture 1 prescaler
func (r *registerCcmr1Type) SetIcpcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldIcpcsMask)|(uint32(value)<<RegisterCcmr1FieldIcpcsShift))
}

const (
	RegisterCcmr1FieldIc1fShift = 4
	RegisterCcmr1FieldIc1fMask  = 0xf0
)

// GetIc1f Input capture 1 filter
func (r *registerCcmr1Type) GetIc1f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldIc1fMask) >> RegisterCcmr1FieldIc1fShift)
}

// SetIc1f Input capture 1 filter
func (r *registerCcmr1Type) SetIc1f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldIc1fMask)|(uint32(value)<<RegisterCcmr1FieldIc1fShift))
}

const (
	RegisterCcmr1FieldIc2pcsShift = 10
	RegisterCcmr1FieldIc2pcsMask  = 0xc00
)

// GetIc2pcs Input capture 2 prescaler
func (r *registerCcmr1Type) GetIc2pcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldIc2pcsMask) >> RegisterCcmr1FieldIc2pcsShift)
}

// SetIc2pcs Input capture 2 prescaler
func (r *registerCcmr1Type) SetIc2pcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldIc2pcsMask)|(uint32(value)<<RegisterCcmr1FieldIc2pcsShift))
}

const (
	RegisterCcmr1FieldIc2fShift = 12
	RegisterCcmr1FieldIc2fMask  = 0xf000
)

// GetIc2f Input capture 2 filter
func (r *registerCcmr1Type) GetIc2f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldIc2fMask) >> RegisterCcmr1FieldIc2fShift)
}

// SetIc2f Input capture 2 filter
func (r *registerCcmr1Type) SetIc2f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldIc2fMask)|(uint32(value)<<RegisterCcmr1FieldIc2fShift))
}

// registerCcmr2Type capture/compare mode register 2
type registerCcmr2Type uint32

const (
	RegisterCcmr2FieldCc3sShift = 0
	RegisterCcmr2FieldCc3sMask  = 0x3
)

// GetCc3s CC3S
func (r *registerCcmr2Type) GetCc3s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldCc3sMask) >> RegisterCcmr2FieldCc3sShift)
}

// SetCc3s CC3S
func (r *registerCcmr2Type) SetCc3s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldCc3sMask)|(uint32(value)<<RegisterCcmr2FieldCc3sShift))
}

const (
	RegisterCcmr2FieldOc3feShift = 2
	RegisterCcmr2FieldOc3feMask  = 0x4
)

// GetOc3fe OC3FE
func (r *registerCcmr2Type) GetOc3fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc3feMask) != 0
}

// SetOc3fe OC3FE
func (r *registerCcmr2Type) SetOc3fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2FieldOc3feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldOc3feMask)
	}
}

const (
	RegisterCcmr2FieldOc3peShift = 3
	RegisterCcmr2FieldOc3peMask  = 0x8
)

// GetOc3pe OC3PE
func (r *registerCcmr2Type) GetOc3pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc3peMask) != 0
}

// SetOc3pe OC3PE
func (r *registerCcmr2Type) SetOc3pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2FieldOc3peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldOc3peMask)
	}
}

const (
	RegisterCcmr2FieldOc3mShift = 4
	RegisterCcmr2FieldOc3mMask  = 0x70
)

// GetOc3m OC3M
func (r *registerCcmr2Type) GetOc3m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc3mMask) >> RegisterCcmr2FieldOc3mShift)
}

// SetOc3m OC3M
func (r *registerCcmr2Type) SetOc3m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldOc3mMask)|(uint32(value)<<RegisterCcmr2FieldOc3mShift))
}

const (
	RegisterCcmr2FieldOc3ceShift = 7
	RegisterCcmr2FieldOc3ceMask  = 0x80
)

// GetOc3ce OC3CE
func (r *registerCcmr2Type) GetOc3ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc3ceMask) != 0
}

// SetOc3ce OC3CE
func (r *registerCcmr2Type) SetOc3ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2FieldOc3ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldOc3ceMask)
	}
}

const (
	RegisterCcmr2FieldCc4sShift = 8
	RegisterCcmr2FieldCc4sMask  = 0x300
)

// GetCc4s CC4S
func (r *registerCcmr2Type) GetCc4s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldCc4sMask) >> RegisterCcmr2FieldCc4sShift)
}

// SetCc4s CC4S
func (r *registerCcmr2Type) SetCc4s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldCc4sMask)|(uint32(value)<<RegisterCcmr2FieldCc4sShift))
}

const (
	RegisterCcmr2FieldOc4feShift = 10
	RegisterCcmr2FieldOc4feMask  = 0x400
)

// GetOc4fe OC4FE
func (r *registerCcmr2Type) GetOc4fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc4feMask) != 0
}

// SetOc4fe OC4FE
func (r *registerCcmr2Type) SetOc4fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2FieldOc4feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldOc4feMask)
	}
}

const (
	RegisterCcmr2FieldOc4peShift = 11
	RegisterCcmr2FieldOc4peMask  = 0x800
)

// GetOc4pe OC4PE
func (r *registerCcmr2Type) GetOc4pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc4peMask) != 0
}

// SetOc4pe OC4PE
func (r *registerCcmr2Type) SetOc4pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2FieldOc4peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldOc4peMask)
	}
}

const (
	RegisterCcmr2FieldOc4mShift = 12
	RegisterCcmr2FieldOc4mMask  = 0x7000
)

// GetOc4m OC4M
func (r *registerCcmr2Type) GetOc4m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc4mMask) >> RegisterCcmr2FieldOc4mShift)
}

// SetOc4m OC4M
func (r *registerCcmr2Type) SetOc4m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldOc4mMask)|(uint32(value)<<RegisterCcmr2FieldOc4mShift))
}

const (
	RegisterCcmr2FieldOc4ceShift = 15
	RegisterCcmr2FieldOc4ceMask  = 0x8000
)

// GetOc4ce OC4CE
func (r *registerCcmr2Type) GetOc4ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc4ceMask) != 0
}

// SetOc4ce OC4CE
func (r *registerCcmr2Type) SetOc4ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2FieldOc4ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldOc4ceMask)
	}
}

const (
	RegisterCcmr2FieldOc3m3Shift = 16
	RegisterCcmr2FieldOc3m3Mask  = 0x10000
)

// GetOc3m3 Output Compare 1 mode - bit 3
func (r *registerCcmr2Type) GetOc3m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc3m3Mask) != 0
}

// SetOc3m3 Output Compare 1 mode - bit 3
func (r *registerCcmr2Type) SetOc3m3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2FieldOc3m3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldOc3m3Mask)
	}
}

const (
	RegisterCcmr2FieldOc4m3Shift = 24
	RegisterCcmr2FieldOc4m3Mask  = 0x1000000
)

// GetOc4m3 Output Compare 2 mode - bit 3
func (r *registerCcmr2Type) GetOc4m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc4m3Mask) != 0
}

// SetOc4m3 Output Compare 2 mode - bit 3
func (r *registerCcmr2Type) SetOc4m3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2FieldOc4m3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldOc4m3Mask)
	}
}

const (
	RegisterCcmr2FieldIc3pscShift = 2
	RegisterCcmr2FieldIc3pscMask  = 0xc
)

// GetIc3psc Input capture 3 prescaler
func (r *registerCcmr2Type) GetIc3psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldIc3pscMask) >> RegisterCcmr2FieldIc3pscShift)
}

// SetIc3psc Input capture 3 prescaler
func (r *registerCcmr2Type) SetIc3psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldIc3pscMask)|(uint32(value)<<RegisterCcmr2FieldIc3pscShift))
}

const (
	RegisterCcmr2FieldIc3fShift = 4
	RegisterCcmr2FieldIc3fMask  = 0xf0
)

// GetIc3f Input capture 3 filter
func (r *registerCcmr2Type) GetIc3f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldIc3fMask) >> RegisterCcmr2FieldIc3fShift)
}

// SetIc3f Input capture 3 filter
func (r *registerCcmr2Type) SetIc3f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldIc3fMask)|(uint32(value)<<RegisterCcmr2FieldIc3fShift))
}

const (
	RegisterCcmr2FieldIc4pscShift = 10
	RegisterCcmr2FieldIc4pscMask  = 0xc00
)

// GetIc4psc Input capture 4 prescaler
func (r *registerCcmr2Type) GetIc4psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldIc4pscMask) >> RegisterCcmr2FieldIc4pscShift)
}

// SetIc4psc Input capture 4 prescaler
func (r *registerCcmr2Type) SetIc4psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldIc4pscMask)|(uint32(value)<<RegisterCcmr2FieldIc4pscShift))
}

const (
	RegisterCcmr2FieldIc4fShift = 12
	RegisterCcmr2FieldIc4fMask  = 0xf000
)

// GetIc4f Input capture 4 filter
func (r *registerCcmr2Type) GetIc4f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldIc4fMask) >> RegisterCcmr2FieldIc4fShift)
}

// SetIc4f Input capture 4 filter
func (r *registerCcmr2Type) SetIc4f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldIc4fMask)|(uint32(value)<<RegisterCcmr2FieldIc4fShift))
}

// registerCcerType capture/compare enable register
type registerCcerType uint32

const (
	RegisterCcerFieldCc1eShift = 0
	RegisterCcerFieldCc1eMask  = 0x1
)

// GetCc1e Capture/Compare 1 output enable
func (r *registerCcerType) GetCc1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc1eMask) != 0
}

// SetCc1e Capture/Compare 1 output enable
func (r *registerCcerType) SetCc1e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc1eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc1eMask)
	}
}

const (
	RegisterCcerFieldCc1pShift = 1
	RegisterCcerFieldCc1pMask  = 0x2
)

// GetCc1p Capture/Compare 1 output Polarity
func (r *registerCcerType) GetCc1p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc1pMask) != 0
}

// SetCc1p Capture/Compare 1 output Polarity
func (r *registerCcerType) SetCc1p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc1pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc1pMask)
	}
}

const (
	RegisterCcerFieldCc1npShift = 3
	RegisterCcerFieldCc1npMask  = 0x8
)

// GetCc1np Capture/Compare 1 output Polarity
func (r *registerCcerType) GetCc1np() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc1npMask) != 0
}

// SetCc1np Capture/Compare 1 output Polarity
func (r *registerCcerType) SetCc1np(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc1npMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc1npMask)
	}
}

const (
	RegisterCcerFieldCc2eShift = 4
	RegisterCcerFieldCc2eMask  = 0x10
)

// GetCc2e Capture/Compare 2 output enable
func (r *registerCcerType) GetCc2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc2eMask) != 0
}

// SetCc2e Capture/Compare 2 output enable
func (r *registerCcerType) SetCc2e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc2eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc2eMask)
	}
}

const (
	RegisterCcerFieldCc2pShift = 5
	RegisterCcerFieldCc2pMask  = 0x20
)

// GetCc2p Capture/Compare 2 output Polarity
func (r *registerCcerType) GetCc2p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc2pMask) != 0
}

// SetCc2p Capture/Compare 2 output Polarity
func (r *registerCcerType) SetCc2p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc2pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc2pMask)
	}
}

const (
	RegisterCcerFieldCc2npShift = 7
	RegisterCcerFieldCc2npMask  = 0x80
)

// GetCc2np Capture/Compare 2 output Polarity
func (r *registerCcerType) GetCc2np() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc2npMask) != 0
}

// SetCc2np Capture/Compare 2 output Polarity
func (r *registerCcerType) SetCc2np(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc2npMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc2npMask)
	}
}

const (
	RegisterCcerFieldCc3eShift = 8
	RegisterCcerFieldCc3eMask  = 0x100
)

// GetCc3e Capture/Compare 3 output enable
func (r *registerCcerType) GetCc3e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc3eMask) != 0
}

// SetCc3e Capture/Compare 3 output enable
func (r *registerCcerType) SetCc3e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc3eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc3eMask)
	}
}

const (
	RegisterCcerFieldCc3pShift = 9
	RegisterCcerFieldCc3pMask  = 0x200
)

// GetCc3p Capture/Compare 3 output Polarity
func (r *registerCcerType) GetCc3p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc3pMask) != 0
}

// SetCc3p Capture/Compare 3 output Polarity
func (r *registerCcerType) SetCc3p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc3pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc3pMask)
	}
}

const (
	RegisterCcerFieldCc3npShift = 11
	RegisterCcerFieldCc3npMask  = 0x800
)

// GetCc3np Capture/Compare 3 output Polarity
func (r *registerCcerType) GetCc3np() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc3npMask) != 0
}

// SetCc3np Capture/Compare 3 output Polarity
func (r *registerCcerType) SetCc3np(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc3npMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc3npMask)
	}
}

const (
	RegisterCcerFieldCc4eShift = 12
	RegisterCcerFieldCc4eMask  = 0x1000
)

// GetCc4e Capture/Compare 4 output enable
func (r *registerCcerType) GetCc4e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc4eMask) != 0
}

// SetCc4e Capture/Compare 4 output enable
func (r *registerCcerType) SetCc4e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc4eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc4eMask)
	}
}

const (
	RegisterCcerFieldCc4pShift = 13
	RegisterCcerFieldCc4pMask  = 0x2000
)

// GetCc4p Capture/Compare 3 output Polarity
func (r *registerCcerType) GetCc4p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc4pMask) != 0
}

// SetCc4p Capture/Compare 3 output Polarity
func (r *registerCcerType) SetCc4p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc4pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc4pMask)
	}
}

const (
	RegisterCcerFieldCc4npShift = 15
	RegisterCcerFieldCc4npMask  = 0x8000
)

// GetCc4np Capture/Compare 4 output Polarity
func (r *registerCcerType) GetCc4np() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc4npMask) != 0
}

// SetCc4np Capture/Compare 4 output Polarity
func (r *registerCcerType) SetCc4np(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc4npMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc4npMask)
	}
}

// registerCntType counter
type registerCntType uint32

const (
	RegisterCntFieldCntlShift = 0
	RegisterCntFieldCntlMask  = 0xffff
)

// GetCntl low counter value
func (r *registerCntType) GetCntl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldCntlMask) >> RegisterCntFieldCntlShift)
}

// SetCntl low counter value
func (r *registerCntType) SetCntl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldCntlMask)|(uint32(value)<<RegisterCntFieldCntlShift))
}

const (
	RegisterCntFieldCnthShift = 16
	RegisterCntFieldCnthMask  = 0xffff0000
)

// GetCnth High counter value
func (r *registerCntType) GetCnth() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldCnthMask) >> RegisterCntFieldCnthShift)
}

// SetCnth High counter value
func (r *registerCntType) SetCnth(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldCnthMask)|(uint32(value)<<RegisterCntFieldCnthShift))
}

// registerPscType prescaler
type registerPscType uint32

const (
	RegisterPscFieldPscShift = 0
	RegisterPscFieldPscMask  = 0xffff
)

// GetPsc Prescaler value
func (r *registerPscType) GetPsc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPscFieldPscMask) >> RegisterPscFieldPscShift)
}

// SetPsc Prescaler value
func (r *registerPscType) SetPsc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPscFieldPscMask)|(uint32(value)<<RegisterPscFieldPscShift))
}

// registerArrType auto-reload register
type registerArrType uint32

const (
	RegisterArrFieldArrlShift = 0
	RegisterArrFieldArrlMask  = 0xffff
)

// GetArrl Low Auto-reload value
func (r *registerArrType) GetArrl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterArrFieldArrlMask) >> RegisterArrFieldArrlShift)
}

// SetArrl Low Auto-reload value
func (r *registerArrType) SetArrl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArrFieldArrlMask)|(uint32(value)<<RegisterArrFieldArrlShift))
}

const (
	RegisterArrFieldArrhShift = 16
	RegisterArrFieldArrhMask  = 0xffff0000
)

// GetArrh High Auto-reload value
func (r *registerArrType) GetArrh() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterArrFieldArrhMask) >> RegisterArrFieldArrhShift)
}

// SetArrh High Auto-reload value
func (r *registerArrType) SetArrh(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArrFieldArrhMask)|(uint32(value)<<RegisterArrFieldArrhShift))
}

// registerCcr1Type capture/compare register 1
type registerCcr1Type uint32

const (
	RegisterCcr1FieldCcr1lShift = 0
	RegisterCcr1FieldCcr1lMask  = 0xffff
)

// GetCcr1l Low Capture/Compare 1 value
func (r *registerCcr1Type) GetCcr1l() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldCcr1lMask) >> RegisterCcr1FieldCcr1lShift)
}

// SetCcr1l Low Capture/Compare 1 value
func (r *registerCcr1Type) SetCcr1l(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldCcr1lMask)|(uint32(value)<<RegisterCcr1FieldCcr1lShift))
}

const (
	RegisterCcr1FieldCcr1hShift = 16
	RegisterCcr1FieldCcr1hMask  = 0xffff0000
)

// GetCcr1h High Capture/Compare 1 value
func (r *registerCcr1Type) GetCcr1h() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldCcr1hMask) >> RegisterCcr1FieldCcr1hShift)
}

// SetCcr1h High Capture/Compare 1 value
func (r *registerCcr1Type) SetCcr1h(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldCcr1hMask)|(uint32(value)<<RegisterCcr1FieldCcr1hShift))
}

// registerCcr2Type capture/compare register 2
type registerCcr2Type uint32

const (
	RegisterCcr2FieldCcr2lShift = 0
	RegisterCcr2FieldCcr2lMask  = 0xffff
)

// GetCcr2l Low Capture/Compare 2 value
func (r *registerCcr2Type) GetCcr2l() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldCcr2lMask) >> RegisterCcr2FieldCcr2lShift)
}

// SetCcr2l Low Capture/Compare 2 value
func (r *registerCcr2Type) SetCcr2l(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldCcr2lMask)|(uint32(value)<<RegisterCcr2FieldCcr2lShift))
}

const (
	RegisterCcr2FieldCcr2hShift = 16
	RegisterCcr2FieldCcr2hMask  = 0xffff0000
)

// GetCcr2h High Capture/Compare 2 value
func (r *registerCcr2Type) GetCcr2h() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldCcr2hMask) >> RegisterCcr2FieldCcr2hShift)
}

// SetCcr2h High Capture/Compare 2 value
func (r *registerCcr2Type) SetCcr2h(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldCcr2hMask)|(uint32(value)<<RegisterCcr2FieldCcr2hShift))
}

// registerCcr3Type capture/compare register 3
type registerCcr3Type uint32

const (
	RegisterCcr3FieldCcr3lShift = 0
	RegisterCcr3FieldCcr3lMask  = 0xffff
)

// GetCcr3l Low Capture/Compare value
func (r *registerCcr3Type) GetCcr3l() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldCcr3lMask) >> RegisterCcr3FieldCcr3lShift)
}

// SetCcr3l Low Capture/Compare value
func (r *registerCcr3Type) SetCcr3l(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldCcr3lMask)|(uint32(value)<<RegisterCcr3FieldCcr3lShift))
}

const (
	RegisterCcr3FieldCcr3hShift = 16
	RegisterCcr3FieldCcr3hMask  = 0xffff0000
)

// GetCcr3h High Capture/Compare value
func (r *registerCcr3Type) GetCcr3h() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldCcr3hMask) >> RegisterCcr3FieldCcr3hShift)
}

// SetCcr3h High Capture/Compare value
func (r *registerCcr3Type) SetCcr3h(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldCcr3hMask)|(uint32(value)<<RegisterCcr3FieldCcr3hShift))
}

// registerCcr4Type capture/compare register 4
type registerCcr4Type uint32

const (
	RegisterCcr4FieldCcr4lShift = 0
	RegisterCcr4FieldCcr4lMask  = 0xffff
)

// GetCcr4l Low Capture/Compare value
func (r *registerCcr4Type) GetCcr4l() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldCcr4lMask) >> RegisterCcr4FieldCcr4lShift)
}

// SetCcr4l Low Capture/Compare value
func (r *registerCcr4Type) SetCcr4l(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldCcr4lMask)|(uint32(value)<<RegisterCcr4FieldCcr4lShift))
}

const (
	RegisterCcr4FieldCcr4hShift = 16
	RegisterCcr4FieldCcr4hMask  = 0xffff0000
)

// GetCcr4h High Capture/Compare value
func (r *registerCcr4Type) GetCcr4h() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldCcr4hMask) >> RegisterCcr4FieldCcr4hShift)
}

// SetCcr4h High Capture/Compare value
func (r *registerCcr4Type) SetCcr4h(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldCcr4hMask)|(uint32(value)<<RegisterCcr4FieldCcr4hShift))
}

// registerDcrType DMA control register
type registerDcrType uint32

const (
	RegisterDcrFieldDbaShift = 0
	RegisterDcrFieldDbaMask  = 0x1f
)

// GetDba DMA base address
func (r *registerDcrType) GetDba() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDcrFieldDbaMask) >> RegisterDcrFieldDbaShift)
}

// SetDba DMA base address
func (r *registerDcrType) SetDba(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcrFieldDbaMask)|(uint32(value)<<RegisterDcrFieldDbaShift))
}

const (
	RegisterDcrFieldDblShift = 8
	RegisterDcrFieldDblMask  = 0x1f00
)

// GetDbl DMA burst length
func (r *registerDcrType) GetDbl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDcrFieldDblMask) >> RegisterDcrFieldDblShift)
}

// SetDbl DMA burst length
func (r *registerDcrType) SetDbl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcrFieldDblMask)|(uint32(value)<<RegisterDcrFieldDblShift))
}

// registerDmarType DMA address for full transfer
type registerDmarType uint32

const (
	RegisterDmarFieldDmabShift = 0
	RegisterDmarFieldDmabMask  = 0xffff
)

// GetDmab DMA register for burst accesses
func (r *registerDmarType) GetDmab() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDmarFieldDmabMask) >> RegisterDmarFieldDmabShift)
}

// SetDmab DMA register for burst accesses
func (r *registerDmarType) SetDmab(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmarFieldDmabMask)|(uint32(value)<<RegisterDmarFieldDmabShift))
}

// registerAf1Type alternate function option register 1
type registerAf1Type uint32

const (
	RegisterAf1FieldEtrselShift = 14
	RegisterAf1FieldEtrselMask  = 0x3c000
)

// GetEtrsel ETR source selection
func (r *registerAf1Type) GetEtrsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldEtrselMask) >> RegisterAf1FieldEtrselShift)
}

// SetEtrsel ETR source selection
func (r *registerAf1Type) SetEtrsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAf1FieldEtrselMask)|(uint32(value)<<RegisterAf1FieldEtrselShift))
}

// registerTiselType timer input selection register
type registerTiselType uint32

const (
	RegisterTiselFieldTi1selShift = 0
	RegisterTiselFieldTi1selMask  = 0xf
)

// GetTi1sel TI1[0] to TI1[15] input selection
func (r *registerTiselType) GetTi1sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi1selMask) >> RegisterTiselFieldTi1selShift)
}

// SetTi1sel TI1[0] to TI1[15] input selection
func (r *registerTiselType) SetTi1sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi1selMask)|(uint32(value)<<RegisterTiselFieldTi1selShift))
}

const (
	RegisterTiselFieldTi2selShift = 8
	RegisterTiselFieldTi2selMask  = 0xf00
)

// GetTi2sel TI2[0] to TI2[15] input selection
func (r *registerTiselType) GetTi2sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi2selMask) >> RegisterTiselFieldTi2selShift)
}

// SetTi2sel TI2[0] to TI2[15] input selection
func (r *registerTiselType) SetTi2sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi2selMask)|(uint32(value)<<RegisterTiselFieldTi2selShift))
}

const (
	RegisterTiselFieldTi3selShift = 16
	RegisterTiselFieldTi3selMask  = 0xf0000
)

// GetTi3sel TI3[0] to TI3[15] input selection
func (r *registerTiselType) GetTi3sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi3selMask) >> RegisterTiselFieldTi3selShift)
}

// SetTi3sel TI3[0] to TI3[15] input selection
func (r *registerTiselType) SetTi3sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi3selMask)|(uint32(value)<<RegisterTiselFieldTi3selShift))
}

const (
	RegisterTiselFieldTi4selShift = 24
	RegisterTiselFieldTi4selMask  = 0xf000000
)

// GetTi4sel TI4[0] to TI4[15] input selection
func (r *registerTiselType) GetTi4sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi4selMask) >> RegisterTiselFieldTi4selShift)
}

// SetTi4sel TI4[0] to TI4[15] input selection
func (r *registerTiselType) SetTi4sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi4selMask)|(uint32(value)<<RegisterTiselFieldTi4selShift))
}
