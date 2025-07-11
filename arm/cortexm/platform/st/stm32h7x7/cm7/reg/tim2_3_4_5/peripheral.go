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
	Cr1   RegisterCr1Type
	Cr2   RegisterCr2Type
	Smcr  RegisterSmcrType
	Dier  RegisterDierType
	Sr    RegisterSrType
	Egr   RegisterEgrType
	Ccmr1 RegisterCcmr1Type
	Ccmr2 RegisterCcmr2Type
	Ccer  RegisterCcerType
	Cnt   RegisterCntType
	Psc   RegisterPscType
	Arr   RegisterArrType
	_     [4]byte
	Ccr1  RegisterCcr1Type
	Ccr2  RegisterCcr2Type
	Ccr3  RegisterCcr3Type
	Ccr4  RegisterCcr4Type
	_     [4]byte
	Dcr   RegisterDcrType
	Dmar  RegisterDmarType
	_     [16]byte
	Af1   RegisterAf1Type
	_     [4]byte
	Tisel RegisterTiselType
}

// RegisterCr1Type control register 1
type RegisterCr1Type uint32

func (r *RegisterCr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCr1FieldCenShift = 0
	RegisterCr1FieldCenMask  = 0x1
)

// GetCen Counter enable
func (r *RegisterCr1Type) GetCen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCenMask) != 0
}

// SetCen Counter enable
func (r *RegisterCr1Type) SetCen(value bool) {
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
func (r *RegisterCr1Type) GetUdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldUdisMask) != 0
}

// SetUdis Update disable
func (r *RegisterCr1Type) SetUdis(value bool) {
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
func (r *RegisterCr1Type) GetUrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldUrsMask) != 0
}

// SetUrs Update request source
func (r *RegisterCr1Type) SetUrs(value bool) {
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
func (r *RegisterCr1Type) GetOpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldOpmMask) != 0
}

// SetOpm One-pulse mode
func (r *RegisterCr1Type) SetOpm(value bool) {
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
func (r *RegisterCr1Type) GetDir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDirMask) != 0
}

// SetDir Direction
func (r *RegisterCr1Type) SetDir(value bool) {
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
func (r *RegisterCr1Type) GetCms() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCmsMask) >> RegisterCr1FieldCmsShift)
}

// SetCms Center-aligned mode selection
func (r *RegisterCr1Type) SetCms(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldCmsMask)|(uint32(value)<<RegisterCr1FieldCmsShift))
}

const (
	RegisterCr1FieldArpeShift = 7
	RegisterCr1FieldArpeMask  = 0x80
)

// GetArpe Auto-reload preload enable
func (r *RegisterCr1Type) GetArpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldArpeMask) != 0
}

// SetArpe Auto-reload preload enable
func (r *RegisterCr1Type) SetArpe(value bool) {
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
func (r *RegisterCr1Type) GetCkd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCkdMask) >> RegisterCr1FieldCkdShift)
}

// SetCkd Clock division
func (r *RegisterCr1Type) SetCkd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldCkdMask)|(uint32(value)<<RegisterCr1FieldCkdShift))
}

const (
	RegisterCr1FieldUifremapShift = 11
	RegisterCr1FieldUifremapMask  = 0x800
)

// GetUifremap UIF status bit remapping
func (r *RegisterCr1Type) GetUifremap() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldUifremapMask) != 0
}

// SetUifremap UIF status bit remapping
func (r *RegisterCr1Type) SetUifremap(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldUifremapMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldUifremapMask)
	}
}

// RegisterCr2Type control register 2
type RegisterCr2Type uint32

func (r *RegisterCr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCr2FieldCcdsShift = 3
	RegisterCr2FieldCcdsMask  = 0x8
)

// GetCcds Capture/compare DMA selection
func (r *RegisterCr2Type) GetCcds() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCcdsMask) != 0
}

// SetCcds Capture/compare DMA selection
func (r *RegisterCr2Type) SetCcds(value bool) {
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
func (r *RegisterCr2Type) GetMms() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldMmsMask) >> RegisterCr2FieldMmsShift)
}

// SetMms Master mode selection
func (r *RegisterCr2Type) SetMms(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldMmsMask)|(uint32(value)<<RegisterCr2FieldMmsShift))
}

const (
	RegisterCr2FieldTi1sShift = 7
	RegisterCr2FieldTi1sMask  = 0x80
)

// GetTi1s TI1 selection
func (r *RegisterCr2Type) GetTi1s() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTi1sMask) != 0
}

// SetTi1s TI1 selection
func (r *RegisterCr2Type) SetTi1s(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTi1sMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTi1sMask)
	}
}

// RegisterSmcrType slave mode control register
type RegisterSmcrType uint32

func (r *RegisterSmcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSmcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSmcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSmcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSmcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSmcrFieldSmsShift = 0
	RegisterSmcrFieldSmsMask  = 0x7
)

// GetSms Slave mode selection
func (r *RegisterSmcrType) GetSms() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldSmsMask) >> RegisterSmcrFieldSmsShift)
}

// SetSms Slave mode selection
func (r *RegisterSmcrType) SetSms(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldSmsMask)|(uint32(value)<<RegisterSmcrFieldSmsShift))
}

const (
	RegisterSmcrFieldTsShift = 4
	RegisterSmcrFieldTsMask  = 0x70
)

// GetTs Trigger selection
func (r *RegisterSmcrType) GetTs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldTsMask) >> RegisterSmcrFieldTsShift)
}

// SetTs Trigger selection
func (r *RegisterSmcrType) SetTs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldTsMask)|(uint32(value)<<RegisterSmcrFieldTsShift))
}

const (
	RegisterSmcrFieldMsmShift = 7
	RegisterSmcrFieldMsmMask  = 0x80
)

// GetMsm Master/Slave mode
func (r *RegisterSmcrType) GetMsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldMsmMask) != 0
}

// SetMsm Master/Slave mode
func (r *RegisterSmcrType) SetMsm(value bool) {
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
func (r *RegisterSmcrType) GetEtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldEtfMask) >> RegisterSmcrFieldEtfShift)
}

// SetEtf External trigger filter
func (r *RegisterSmcrType) SetEtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldEtfMask)|(uint32(value)<<RegisterSmcrFieldEtfShift))
}

const (
	RegisterSmcrFieldEtpsShift = 12
	RegisterSmcrFieldEtpsMask  = 0x3000
)

// GetEtps External trigger prescaler
func (r *RegisterSmcrType) GetEtps() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldEtpsMask) >> RegisterSmcrFieldEtpsShift)
}

// SetEtps External trigger prescaler
func (r *RegisterSmcrType) SetEtps(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldEtpsMask)|(uint32(value)<<RegisterSmcrFieldEtpsShift))
}

const (
	RegisterSmcrFieldEceShift = 14
	RegisterSmcrFieldEceMask  = 0x4000
)

// GetEce External clock enable
func (r *RegisterSmcrType) GetEce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldEceMask) != 0
}

// SetEce External clock enable
func (r *RegisterSmcrType) SetEce(value bool) {
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
func (r *RegisterSmcrType) GetEtp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldEtpMask) != 0
}

// SetEtp External trigger polarity
func (r *RegisterSmcrType) SetEtp(value bool) {
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
func (r *RegisterSmcrType) GetSms3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldSms3Mask) != 0
}

// SetSms3 Slave mode selection - bit 3
func (r *RegisterSmcrType) SetSms3(value bool) {
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
func (r *RegisterSmcrType) GetTs43() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldTs43Mask) >> RegisterSmcrFieldTs43Shift)
}

// SetTs43 Trigger selection
func (r *RegisterSmcrType) SetTs43(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldTs43Mask)|(uint32(value)<<RegisterSmcrFieldTs43Shift))
}

// RegisterDierType DMA/Interrupt enable register
type RegisterDierType uint32

func (r *RegisterDierType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDierType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDierType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDierType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDierType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDierFieldUieShift = 0
	RegisterDierFieldUieMask  = 0x1
)

// GetUie Update interrupt enable
func (r *RegisterDierType) GetUie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldUieMask) != 0
}

// SetUie Update interrupt enable
func (r *RegisterDierType) SetUie(value bool) {
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
func (r *RegisterDierType) GetCc1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc1ieMask) != 0
}

// SetCc1ie Capture/Compare 1 interrupt enable
func (r *RegisterDierType) SetCc1ie(value bool) {
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
func (r *RegisterDierType) GetCc2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc2ieMask) != 0
}

// SetCc2ie Capture/Compare 2 interrupt enable
func (r *RegisterDierType) SetCc2ie(value bool) {
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
func (r *RegisterDierType) GetCc3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc3ieMask) != 0
}

// SetCc3ie Capture/Compare 3 interrupt enable
func (r *RegisterDierType) SetCc3ie(value bool) {
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
func (r *RegisterDierType) GetCc4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc4ieMask) != 0
}

// SetCc4ie Capture/Compare 4 interrupt enable
func (r *RegisterDierType) SetCc4ie(value bool) {
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
func (r *RegisterDierType) GetTie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldTieMask) != 0
}

// SetTie Trigger interrupt enable
func (r *RegisterDierType) SetTie(value bool) {
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
func (r *RegisterDierType) GetUde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldUdeMask) != 0
}

// SetUde Update DMA request enable
func (r *RegisterDierType) SetUde(value bool) {
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
func (r *RegisterDierType) GetCc1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc1deMask) != 0
}

// SetCc1de Capture/Compare 1 DMA request enable
func (r *RegisterDierType) SetCc1de(value bool) {
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
func (r *RegisterDierType) GetCc2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc2deMask) != 0
}

// SetCc2de Capture/Compare 2 DMA request enable
func (r *RegisterDierType) SetCc2de(value bool) {
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
func (r *RegisterDierType) GetCc3de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc3deMask) != 0
}

// SetCc3de Capture/Compare 3 DMA request enable
func (r *RegisterDierType) SetCc3de(value bool) {
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
func (r *RegisterDierType) GetCc4de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCc4deMask) != 0
}

// SetCc4de Capture/Compare 4 DMA request enable
func (r *RegisterDierType) SetCc4de(value bool) {
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
func (r *RegisterDierType) GetTde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldTdeMask) != 0
}

// SetTde Trigger DMA request enable
func (r *RegisterDierType) SetTde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldTdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldTdeMask)
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
	RegisterSrFieldUifShift = 0
	RegisterSrFieldUifMask  = 0x1
)

// GetUif Update interrupt flag
func (r *RegisterSrType) GetUif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldUifMask) != 0
}

// SetUif Update interrupt flag
func (r *RegisterSrType) SetUif(value bool) {
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
func (r *RegisterSrType) GetCc1if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc1ifMask) != 0
}

// SetCc1if Capture/compare 1 interrupt flag
func (r *RegisterSrType) SetCc1if(value bool) {
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
func (r *RegisterSrType) GetCc2if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc2ifMask) != 0
}

// SetCc2if Capture/Compare 2 interrupt flag
func (r *RegisterSrType) SetCc2if(value bool) {
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
func (r *RegisterSrType) GetCc3if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc3ifMask) != 0
}

// SetCc3if Capture/Compare 3 interrupt flag
func (r *RegisterSrType) SetCc3if(value bool) {
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
func (r *RegisterSrType) GetCc4if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc4ifMask) != 0
}

// SetCc4if Capture/Compare 4 interrupt flag
func (r *RegisterSrType) SetCc4if(value bool) {
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
func (r *RegisterSrType) GetTif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTifMask) != 0
}

// SetTif Trigger interrupt flag
func (r *RegisterSrType) SetTif(value bool) {
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
func (r *RegisterSrType) GetCc1of() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc1ofMask) != 0
}

// SetCc1of Capture/Compare 1 overcapture flag
func (r *RegisterSrType) SetCc1of(value bool) {
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
func (r *RegisterSrType) GetCc2of() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc2ofMask) != 0
}

// SetCc2of Capture/compare 2 overcapture flag
func (r *RegisterSrType) SetCc2of(value bool) {
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
func (r *RegisterSrType) GetCc3of() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc3ofMask) != 0
}

// SetCc3of Capture/Compare 3 overcapture flag
func (r *RegisterSrType) SetCc3of(value bool) {
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
func (r *RegisterSrType) GetCc4of() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc4ofMask) != 0
}

// SetCc4of Capture/Compare 4 overcapture flag
func (r *RegisterSrType) SetCc4of(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCc4ofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCc4ofMask)
	}
}

// RegisterEgrType event generation register
type RegisterEgrType uint32

func (r *RegisterEgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterEgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterEgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterEgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterEgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterEgrFieldUgShift = 0
	RegisterEgrFieldUgMask  = 0x1
)

// GetUg Update generation
func (r *RegisterEgrType) GetUg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldUgMask) != 0
}

// SetUg Update generation
func (r *RegisterEgrType) SetUg(value bool) {
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
func (r *RegisterEgrType) GetCc1g() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldCc1gMask) != 0
}

// SetCc1g Capture/compare 1 generation
func (r *RegisterEgrType) SetCc1g(value bool) {
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
func (r *RegisterEgrType) GetCc2g() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldCc2gMask) != 0
}

// SetCc2g Capture/compare 2 generation
func (r *RegisterEgrType) SetCc2g(value bool) {
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
func (r *RegisterEgrType) GetCc3g() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldCc3gMask) != 0
}

// SetCc3g Capture/compare 3 generation
func (r *RegisterEgrType) SetCc3g(value bool) {
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
func (r *RegisterEgrType) GetCc4g() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldCc4gMask) != 0
}

// SetCc4g Capture/compare 4 generation
func (r *RegisterEgrType) SetCc4g(value bool) {
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
func (r *RegisterEgrType) GetTg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldTgMask) != 0
}

// SetTg Trigger generation
func (r *RegisterEgrType) SetTg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldTgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldTgMask)
	}
}

// RegisterCcmr1Type capture/compare mode register 1
type RegisterCcmr1Type uint32

func (r *RegisterCcmr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcmr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcmr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcmr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcmr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcmr1FieldCc1sShift = 0
	RegisterCcmr1FieldCc1sMask  = 0x3
)

// GetCc1s CC1S
func (r *RegisterCcmr1Type) GetCc1s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldCc1sMask) >> RegisterCcmr1FieldCc1sShift)
}

// SetCc1s CC1S
func (r *RegisterCcmr1Type) SetCc1s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldCc1sMask)|(uint32(value)<<RegisterCcmr1FieldCc1sShift))
}

const (
	RegisterCcmr1FieldOc1feShift = 2
	RegisterCcmr1FieldOc1feMask  = 0x4
)

// GetOc1fe OC1FE
func (r *RegisterCcmr1Type) GetOc1fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc1feMask) != 0
}

// SetOc1fe OC1FE
func (r *RegisterCcmr1Type) SetOc1fe(value bool) {
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
func (r *RegisterCcmr1Type) GetOc1pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc1peMask) != 0
}

// SetOc1pe OC1PE
func (r *RegisterCcmr1Type) SetOc1pe(value bool) {
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
func (r *RegisterCcmr1Type) GetOc1m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc1mMask) >> RegisterCcmr1FieldOc1mShift)
}

// SetOc1m OC1M
func (r *RegisterCcmr1Type) SetOc1m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldOc1mMask)|(uint32(value)<<RegisterCcmr1FieldOc1mShift))
}

const (
	RegisterCcmr1FieldOc1ceShift = 7
	RegisterCcmr1FieldOc1ceMask  = 0x80
)

// GetOc1ce OC1CE
func (r *RegisterCcmr1Type) GetOc1ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc1ceMask) != 0
}

// SetOc1ce OC1CE
func (r *RegisterCcmr1Type) SetOc1ce(value bool) {
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
func (r *RegisterCcmr1Type) GetCc2s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldCc2sMask) >> RegisterCcmr1FieldCc2sShift)
}

// SetCc2s CC2S
func (r *RegisterCcmr1Type) SetCc2s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldCc2sMask)|(uint32(value)<<RegisterCcmr1FieldCc2sShift))
}

const (
	RegisterCcmr1FieldOc2feShift = 10
	RegisterCcmr1FieldOc2feMask  = 0x400
)

// GetOc2fe OC2FE
func (r *RegisterCcmr1Type) GetOc2fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc2feMask) != 0
}

// SetOc2fe OC2FE
func (r *RegisterCcmr1Type) SetOc2fe(value bool) {
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
func (r *RegisterCcmr1Type) GetOc2pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc2peMask) != 0
}

// SetOc2pe OC2PE
func (r *RegisterCcmr1Type) SetOc2pe(value bool) {
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
func (r *RegisterCcmr1Type) GetOc2m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc2mMask) >> RegisterCcmr1FieldOc2mShift)
}

// SetOc2m OC2M
func (r *RegisterCcmr1Type) SetOc2m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldOc2mMask)|(uint32(value)<<RegisterCcmr1FieldOc2mShift))
}

const (
	RegisterCcmr1FieldOc2ceShift = 15
	RegisterCcmr1FieldOc2ceMask  = 0x8000
)

// GetOc2ce OC2CE
func (r *RegisterCcmr1Type) GetOc2ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc2ceMask) != 0
}

// SetOc2ce OC2CE
func (r *RegisterCcmr1Type) SetOc2ce(value bool) {
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
func (r *RegisterCcmr1Type) GetOc1m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc1m3Mask) != 0
}

// SetOc1m3 Output Compare 1 mode - bit 3
func (r *RegisterCcmr1Type) SetOc1m3(value bool) {
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
func (r *RegisterCcmr1Type) GetOc2m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldOc2m3Mask) != 0
}

// SetOc2m3 Output Compare 2 mode - bit 3
func (r *RegisterCcmr1Type) SetOc2m3(value bool) {
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
func (r *RegisterCcmr1Type) GetIcpcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldIcpcsMask) >> RegisterCcmr1FieldIcpcsShift)
}

// SetIcpcs Input capture 1 prescaler
func (r *RegisterCcmr1Type) SetIcpcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldIcpcsMask)|(uint32(value)<<RegisterCcmr1FieldIcpcsShift))
}

const (
	RegisterCcmr1FieldIc1fShift = 4
	RegisterCcmr1FieldIc1fMask  = 0xf0
)

// GetIc1f Input capture 1 filter
func (r *RegisterCcmr1Type) GetIc1f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldIc1fMask) >> RegisterCcmr1FieldIc1fShift)
}

// SetIc1f Input capture 1 filter
func (r *RegisterCcmr1Type) SetIc1f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldIc1fMask)|(uint32(value)<<RegisterCcmr1FieldIc1fShift))
}

const (
	RegisterCcmr1FieldIc2pcsShift = 10
	RegisterCcmr1FieldIc2pcsMask  = 0xc00
)

// GetIc2pcs Input capture 2 prescaler
func (r *RegisterCcmr1Type) GetIc2pcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldIc2pcsMask) >> RegisterCcmr1FieldIc2pcsShift)
}

// SetIc2pcs Input capture 2 prescaler
func (r *RegisterCcmr1Type) SetIc2pcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldIc2pcsMask)|(uint32(value)<<RegisterCcmr1FieldIc2pcsShift))
}

const (
	RegisterCcmr1FieldIc2fShift = 12
	RegisterCcmr1FieldIc2fMask  = 0xf000
)

// GetIc2f Input capture 2 filter
func (r *RegisterCcmr1Type) GetIc2f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1FieldIc2fMask) >> RegisterCcmr1FieldIc2fShift)
}

// SetIc2f Input capture 2 filter
func (r *RegisterCcmr1Type) SetIc2f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1FieldIc2fMask)|(uint32(value)<<RegisterCcmr1FieldIc2fShift))
}

// RegisterCcmr2Type capture/compare mode register 2
type RegisterCcmr2Type uint32

func (r *RegisterCcmr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcmr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcmr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcmr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcmr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcmr2FieldCc3sShift = 0
	RegisterCcmr2FieldCc3sMask  = 0x3
)

// GetCc3s CC3S
func (r *RegisterCcmr2Type) GetCc3s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldCc3sMask) >> RegisterCcmr2FieldCc3sShift)
}

// SetCc3s CC3S
func (r *RegisterCcmr2Type) SetCc3s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldCc3sMask)|(uint32(value)<<RegisterCcmr2FieldCc3sShift))
}

const (
	RegisterCcmr2FieldOc3feShift = 2
	RegisterCcmr2FieldOc3feMask  = 0x4
)

// GetOc3fe OC3FE
func (r *RegisterCcmr2Type) GetOc3fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc3feMask) != 0
}

// SetOc3fe OC3FE
func (r *RegisterCcmr2Type) SetOc3fe(value bool) {
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
func (r *RegisterCcmr2Type) GetOc3pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc3peMask) != 0
}

// SetOc3pe OC3PE
func (r *RegisterCcmr2Type) SetOc3pe(value bool) {
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
func (r *RegisterCcmr2Type) GetOc3m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc3mMask) >> RegisterCcmr2FieldOc3mShift)
}

// SetOc3m OC3M
func (r *RegisterCcmr2Type) SetOc3m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldOc3mMask)|(uint32(value)<<RegisterCcmr2FieldOc3mShift))
}

const (
	RegisterCcmr2FieldOc3ceShift = 7
	RegisterCcmr2FieldOc3ceMask  = 0x80
)

// GetOc3ce OC3CE
func (r *RegisterCcmr2Type) GetOc3ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc3ceMask) != 0
}

// SetOc3ce OC3CE
func (r *RegisterCcmr2Type) SetOc3ce(value bool) {
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
func (r *RegisterCcmr2Type) GetCc4s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldCc4sMask) >> RegisterCcmr2FieldCc4sShift)
}

// SetCc4s CC4S
func (r *RegisterCcmr2Type) SetCc4s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldCc4sMask)|(uint32(value)<<RegisterCcmr2FieldCc4sShift))
}

const (
	RegisterCcmr2FieldOc4feShift = 10
	RegisterCcmr2FieldOc4feMask  = 0x400
)

// GetOc4fe OC4FE
func (r *RegisterCcmr2Type) GetOc4fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc4feMask) != 0
}

// SetOc4fe OC4FE
func (r *RegisterCcmr2Type) SetOc4fe(value bool) {
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
func (r *RegisterCcmr2Type) GetOc4pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc4peMask) != 0
}

// SetOc4pe OC4PE
func (r *RegisterCcmr2Type) SetOc4pe(value bool) {
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
func (r *RegisterCcmr2Type) GetOc4m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc4mMask) >> RegisterCcmr2FieldOc4mShift)
}

// SetOc4m OC4M
func (r *RegisterCcmr2Type) SetOc4m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldOc4mMask)|(uint32(value)<<RegisterCcmr2FieldOc4mShift))
}

const (
	RegisterCcmr2FieldOc4ceShift = 15
	RegisterCcmr2FieldOc4ceMask  = 0x8000
)

// GetOc4ce OC4CE
func (r *RegisterCcmr2Type) GetOc4ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc4ceMask) != 0
}

// SetOc4ce OC4CE
func (r *RegisterCcmr2Type) SetOc4ce(value bool) {
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
func (r *RegisterCcmr2Type) GetOc3m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc3m3Mask) != 0
}

// SetOc3m3 Output Compare 1 mode - bit 3
func (r *RegisterCcmr2Type) SetOc3m3(value bool) {
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
func (r *RegisterCcmr2Type) GetOc4m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldOc4m3Mask) != 0
}

// SetOc4m3 Output Compare 2 mode - bit 3
func (r *RegisterCcmr2Type) SetOc4m3(value bool) {
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
func (r *RegisterCcmr2Type) GetIc3psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldIc3pscMask) >> RegisterCcmr2FieldIc3pscShift)
}

// SetIc3psc Input capture 3 prescaler
func (r *RegisterCcmr2Type) SetIc3psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldIc3pscMask)|(uint32(value)<<RegisterCcmr2FieldIc3pscShift))
}

const (
	RegisterCcmr2FieldIc3fShift = 4
	RegisterCcmr2FieldIc3fMask  = 0xf0
)

// GetIc3f Input capture 3 filter
func (r *RegisterCcmr2Type) GetIc3f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldIc3fMask) >> RegisterCcmr2FieldIc3fShift)
}

// SetIc3f Input capture 3 filter
func (r *RegisterCcmr2Type) SetIc3f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldIc3fMask)|(uint32(value)<<RegisterCcmr2FieldIc3fShift))
}

const (
	RegisterCcmr2FieldIc4pscShift = 10
	RegisterCcmr2FieldIc4pscMask  = 0xc00
)

// GetIc4psc Input capture 4 prescaler
func (r *RegisterCcmr2Type) GetIc4psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldIc4pscMask) >> RegisterCcmr2FieldIc4pscShift)
}

// SetIc4psc Input capture 4 prescaler
func (r *RegisterCcmr2Type) SetIc4psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldIc4pscMask)|(uint32(value)<<RegisterCcmr2FieldIc4pscShift))
}

const (
	RegisterCcmr2FieldIc4fShift = 12
	RegisterCcmr2FieldIc4fMask  = 0xf000
)

// GetIc4f Input capture 4 filter
func (r *RegisterCcmr2Type) GetIc4f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2FieldIc4fMask) >> RegisterCcmr2FieldIc4fShift)
}

// SetIc4f Input capture 4 filter
func (r *RegisterCcmr2Type) SetIc4f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2FieldIc4fMask)|(uint32(value)<<RegisterCcmr2FieldIc4fShift))
}

// RegisterCcerType capture/compare enable register
type RegisterCcerType uint32

func (r *RegisterCcerType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcerType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcerType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcerType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcerType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcerFieldCc1eShift = 0
	RegisterCcerFieldCc1eMask  = 0x1
)

// GetCc1e Capture/Compare 1 output enable
func (r *RegisterCcerType) GetCc1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc1eMask) != 0
}

// SetCc1e Capture/Compare 1 output enable
func (r *RegisterCcerType) SetCc1e(value bool) {
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
func (r *RegisterCcerType) GetCc1p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc1pMask) != 0
}

// SetCc1p Capture/Compare 1 output Polarity
func (r *RegisterCcerType) SetCc1p(value bool) {
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
func (r *RegisterCcerType) GetCc1np() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc1npMask) != 0
}

// SetCc1np Capture/Compare 1 output Polarity
func (r *RegisterCcerType) SetCc1np(value bool) {
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
func (r *RegisterCcerType) GetCc2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc2eMask) != 0
}

// SetCc2e Capture/Compare 2 output enable
func (r *RegisterCcerType) SetCc2e(value bool) {
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
func (r *RegisterCcerType) GetCc2p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc2pMask) != 0
}

// SetCc2p Capture/Compare 2 output Polarity
func (r *RegisterCcerType) SetCc2p(value bool) {
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
func (r *RegisterCcerType) GetCc2np() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc2npMask) != 0
}

// SetCc2np Capture/Compare 2 output Polarity
func (r *RegisterCcerType) SetCc2np(value bool) {
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
func (r *RegisterCcerType) GetCc3e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc3eMask) != 0
}

// SetCc3e Capture/Compare 3 output enable
func (r *RegisterCcerType) SetCc3e(value bool) {
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
func (r *RegisterCcerType) GetCc3p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc3pMask) != 0
}

// SetCc3p Capture/Compare 3 output Polarity
func (r *RegisterCcerType) SetCc3p(value bool) {
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
func (r *RegisterCcerType) GetCc3np() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc3npMask) != 0
}

// SetCc3np Capture/Compare 3 output Polarity
func (r *RegisterCcerType) SetCc3np(value bool) {
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
func (r *RegisterCcerType) GetCc4e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc4eMask) != 0
}

// SetCc4e Capture/Compare 4 output enable
func (r *RegisterCcerType) SetCc4e(value bool) {
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
func (r *RegisterCcerType) GetCc4p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc4pMask) != 0
}

// SetCc4p Capture/Compare 3 output Polarity
func (r *RegisterCcerType) SetCc4p(value bool) {
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
func (r *RegisterCcerType) GetCc4np() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc4npMask) != 0
}

// SetCc4np Capture/Compare 4 output Polarity
func (r *RegisterCcerType) SetCc4np(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc4npMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc4npMask)
	}
}

// RegisterCntType counter
type RegisterCntType uint32

func (r *RegisterCntType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCntType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCntType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCntType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCntType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCntFieldCntlShift = 0
	RegisterCntFieldCntlMask  = 0xffff
)

// GetCntl low counter value
func (r *RegisterCntType) GetCntl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldCntlMask) >> RegisterCntFieldCntlShift)
}

// SetCntl low counter value
func (r *RegisterCntType) SetCntl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldCntlMask)|(uint32(value)<<RegisterCntFieldCntlShift))
}

const (
	RegisterCntFieldCnthShift = 16
	RegisterCntFieldCnthMask  = 0xffff0000
)

// GetCnth High counter value
func (r *RegisterCntType) GetCnth() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldCnthMask) >> RegisterCntFieldCnthShift)
}

// SetCnth High counter value
func (r *RegisterCntType) SetCnth(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldCnthMask)|(uint32(value)<<RegisterCntFieldCnthShift))
}

// RegisterPscType prescaler
type RegisterPscType uint32

func (r *RegisterPscType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPscType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPscType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPscType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPscType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPscFieldPscShift = 0
	RegisterPscFieldPscMask  = 0xffff
)

// GetPsc Prescaler value
func (r *RegisterPscType) GetPsc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPscFieldPscMask) >> RegisterPscFieldPscShift)
}

// SetPsc Prescaler value
func (r *RegisterPscType) SetPsc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPscFieldPscMask)|(uint32(value)<<RegisterPscFieldPscShift))
}

// RegisterArrType auto-reload register
type RegisterArrType uint32

func (r *RegisterArrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterArrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterArrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterArrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterArrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterArrFieldArrlShift = 0
	RegisterArrFieldArrlMask  = 0xffff
)

// GetArrl Low Auto-reload value
func (r *RegisterArrType) GetArrl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterArrFieldArrlMask) >> RegisterArrFieldArrlShift)
}

// SetArrl Low Auto-reload value
func (r *RegisterArrType) SetArrl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArrFieldArrlMask)|(uint32(value)<<RegisterArrFieldArrlShift))
}

const (
	RegisterArrFieldArrhShift = 16
	RegisterArrFieldArrhMask  = 0xffff0000
)

// GetArrh High Auto-reload value
func (r *RegisterArrType) GetArrh() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterArrFieldArrhMask) >> RegisterArrFieldArrhShift)
}

// SetArrh High Auto-reload value
func (r *RegisterArrType) SetArrh(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArrFieldArrhMask)|(uint32(value)<<RegisterArrFieldArrhShift))
}

// RegisterCcr1Type capture/compare register 1
type RegisterCcr1Type uint32

func (r *RegisterCcr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcr1FieldCcr1lShift = 0
	RegisterCcr1FieldCcr1lMask  = 0xffff
)

// GetCcr1l Low Capture/Compare 1 value
func (r *RegisterCcr1Type) GetCcr1l() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldCcr1lMask) >> RegisterCcr1FieldCcr1lShift)
}

// SetCcr1l Low Capture/Compare 1 value
func (r *RegisterCcr1Type) SetCcr1l(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldCcr1lMask)|(uint32(value)<<RegisterCcr1FieldCcr1lShift))
}

const (
	RegisterCcr1FieldCcr1hShift = 16
	RegisterCcr1FieldCcr1hMask  = 0xffff0000
)

// GetCcr1h High Capture/Compare 1 value
func (r *RegisterCcr1Type) GetCcr1h() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldCcr1hMask) >> RegisterCcr1FieldCcr1hShift)
}

// SetCcr1h High Capture/Compare 1 value
func (r *RegisterCcr1Type) SetCcr1h(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldCcr1hMask)|(uint32(value)<<RegisterCcr1FieldCcr1hShift))
}

// RegisterCcr2Type capture/compare register 2
type RegisterCcr2Type uint32

func (r *RegisterCcr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcr2FieldCcr2lShift = 0
	RegisterCcr2FieldCcr2lMask  = 0xffff
)

// GetCcr2l Low Capture/Compare 2 value
func (r *RegisterCcr2Type) GetCcr2l() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldCcr2lMask) >> RegisterCcr2FieldCcr2lShift)
}

// SetCcr2l Low Capture/Compare 2 value
func (r *RegisterCcr2Type) SetCcr2l(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldCcr2lMask)|(uint32(value)<<RegisterCcr2FieldCcr2lShift))
}

const (
	RegisterCcr2FieldCcr2hShift = 16
	RegisterCcr2FieldCcr2hMask  = 0xffff0000
)

// GetCcr2h High Capture/Compare 2 value
func (r *RegisterCcr2Type) GetCcr2h() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldCcr2hMask) >> RegisterCcr2FieldCcr2hShift)
}

// SetCcr2h High Capture/Compare 2 value
func (r *RegisterCcr2Type) SetCcr2h(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldCcr2hMask)|(uint32(value)<<RegisterCcr2FieldCcr2hShift))
}

// RegisterCcr3Type capture/compare register 3
type RegisterCcr3Type uint32

func (r *RegisterCcr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcr3FieldCcr3lShift = 0
	RegisterCcr3FieldCcr3lMask  = 0xffff
)

// GetCcr3l Low Capture/Compare value
func (r *RegisterCcr3Type) GetCcr3l() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldCcr3lMask) >> RegisterCcr3FieldCcr3lShift)
}

// SetCcr3l Low Capture/Compare value
func (r *RegisterCcr3Type) SetCcr3l(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldCcr3lMask)|(uint32(value)<<RegisterCcr3FieldCcr3lShift))
}

const (
	RegisterCcr3FieldCcr3hShift = 16
	RegisterCcr3FieldCcr3hMask  = 0xffff0000
)

// GetCcr3h High Capture/Compare value
func (r *RegisterCcr3Type) GetCcr3h() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldCcr3hMask) >> RegisterCcr3FieldCcr3hShift)
}

// SetCcr3h High Capture/Compare value
func (r *RegisterCcr3Type) SetCcr3h(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldCcr3hMask)|(uint32(value)<<RegisterCcr3FieldCcr3hShift))
}

// RegisterCcr4Type capture/compare register 4
type RegisterCcr4Type uint32

func (r *RegisterCcr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcr4FieldCcr4lShift = 0
	RegisterCcr4FieldCcr4lMask  = 0xffff
)

// GetCcr4l Low Capture/Compare value
func (r *RegisterCcr4Type) GetCcr4l() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldCcr4lMask) >> RegisterCcr4FieldCcr4lShift)
}

// SetCcr4l Low Capture/Compare value
func (r *RegisterCcr4Type) SetCcr4l(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldCcr4lMask)|(uint32(value)<<RegisterCcr4FieldCcr4lShift))
}

const (
	RegisterCcr4FieldCcr4hShift = 16
	RegisterCcr4FieldCcr4hMask  = 0xffff0000
)

// GetCcr4h High Capture/Compare value
func (r *RegisterCcr4Type) GetCcr4h() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldCcr4hMask) >> RegisterCcr4FieldCcr4hShift)
}

// SetCcr4h High Capture/Compare value
func (r *RegisterCcr4Type) SetCcr4h(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldCcr4hMask)|(uint32(value)<<RegisterCcr4FieldCcr4hShift))
}

// RegisterDcrType DMA control register
type RegisterDcrType uint32

func (r *RegisterDcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDcrFieldDbaShift = 0
	RegisterDcrFieldDbaMask  = 0x1f
)

// GetDba DMA base address
func (r *RegisterDcrType) GetDba() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDcrFieldDbaMask) >> RegisterDcrFieldDbaShift)
}

// SetDba DMA base address
func (r *RegisterDcrType) SetDba(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcrFieldDbaMask)|(uint32(value)<<RegisterDcrFieldDbaShift))
}

const (
	RegisterDcrFieldDblShift = 8
	RegisterDcrFieldDblMask  = 0x1f00
)

// GetDbl DMA burst length
func (r *RegisterDcrType) GetDbl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDcrFieldDblMask) >> RegisterDcrFieldDblShift)
}

// SetDbl DMA burst length
func (r *RegisterDcrType) SetDbl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcrFieldDblMask)|(uint32(value)<<RegisterDcrFieldDblShift))
}

// RegisterDmarType DMA address for full transfer
type RegisterDmarType uint32

func (r *RegisterDmarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmarFieldDmabShift = 0
	RegisterDmarFieldDmabMask  = 0xffff
)

// GetDmab DMA register for burst accesses
func (r *RegisterDmarType) GetDmab() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDmarFieldDmabMask) >> RegisterDmarFieldDmabShift)
}

// SetDmab DMA register for burst accesses
func (r *RegisterDmarType) SetDmab(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmarFieldDmabMask)|(uint32(value)<<RegisterDmarFieldDmabShift))
}

// RegisterAf1Type alternate function option register 1
type RegisterAf1Type uint32

func (r *RegisterAf1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAf1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAf1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAf1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAf1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAf1FieldEtrselShift = 14
	RegisterAf1FieldEtrselMask  = 0x3c000
)

// GetEtrsel ETR source selection
func (r *RegisterAf1Type) GetEtrsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldEtrselMask) >> RegisterAf1FieldEtrselShift)
}

// SetEtrsel ETR source selection
func (r *RegisterAf1Type) SetEtrsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAf1FieldEtrselMask)|(uint32(value)<<RegisterAf1FieldEtrselShift))
}

// RegisterTiselType timer input selection register
type RegisterTiselType uint32

func (r *RegisterTiselType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTiselType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTiselType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTiselType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTiselType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTiselFieldTi1selShift = 0
	RegisterTiselFieldTi1selMask  = 0xf
)

// GetTi1sel TI1[0] to TI1[15] input selection
func (r *RegisterTiselType) GetTi1sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi1selMask) >> RegisterTiselFieldTi1selShift)
}

// SetTi1sel TI1[0] to TI1[15] input selection
func (r *RegisterTiselType) SetTi1sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi1selMask)|(uint32(value)<<RegisterTiselFieldTi1selShift))
}

const (
	RegisterTiselFieldTi2selShift = 8
	RegisterTiselFieldTi2selMask  = 0xf00
)

// GetTi2sel TI2[0] to TI2[15] input selection
func (r *RegisterTiselType) GetTi2sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi2selMask) >> RegisterTiselFieldTi2selShift)
}

// SetTi2sel TI2[0] to TI2[15] input selection
func (r *RegisterTiselType) SetTi2sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi2selMask)|(uint32(value)<<RegisterTiselFieldTi2selShift))
}

const (
	RegisterTiselFieldTi3selShift = 16
	RegisterTiselFieldTi3selMask  = 0xf0000
)

// GetTi3sel TI3[0] to TI3[15] input selection
func (r *RegisterTiselType) GetTi3sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi3selMask) >> RegisterTiselFieldTi3selShift)
}

// SetTi3sel TI3[0] to TI3[15] input selection
func (r *RegisterTiselType) SetTi3sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi3selMask)|(uint32(value)<<RegisterTiselFieldTi3selShift))
}

const (
	RegisterTiselFieldTi4selShift = 24
	RegisterTiselFieldTi4selMask  = 0xf000000
)

// GetTi4sel TI4[0] to TI4[15] input selection
func (r *RegisterTiselType) GetTi4sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi4selMask) >> RegisterTiselFieldTi4selShift)
}

// SetTi4sel TI4[0] to TI4[15] input selection
func (r *RegisterTiselType) SetTi4sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi4selMask)|(uint32(value)<<RegisterTiselFieldTi4selShift))
}
