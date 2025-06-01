//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package tim12

import (
	"unsafe"
	"volatile"
)

var (
	Tim12 = (*_tim12)(unsafe.Pointer(uintptr(0x40001800)))
)

type _tim12 struct {
	Cr1   registerCr1Type
	Cr2   registerCr2Type
	Smcr  registerSmcrType
	Dier  registerDierType
	Sr    registerSrType
	Egr   registerEgrType
	Ccmr1 registerCcmr1Type
	_     [4]byte
	Ccer  registerCcerType
	Cnt   registerCntType
	Psc   registerPscType
	Arr   registerArrType
	_     [4]byte
	Ccr1  registerCcr1Type
	Ccr2  registerCcr2Type
	_     [44]byte
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

// registerCcmr1Type capture/compare mode register 1 (output mode)
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

// registerCntType counter
type registerCntType uint32

const (
	RegisterCntFieldCntShift = 0
	RegisterCntFieldCntMask  = 0xffff
)

// GetCnt low counter value
func (r *registerCntType) GetCnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldCntMask) >> RegisterCntFieldCntShift)
}

// SetCnt low counter value
func (r *registerCntType) SetCnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldCntMask)|(uint32(value)<<RegisterCntFieldCntShift))
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
	RegisterArrFieldArrShift = 0
	RegisterArrFieldArrMask  = 0xffff
)

// GetArr Low Auto-reload value
func (r *registerArrType) GetArr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterArrFieldArrMask) >> RegisterArrFieldArrShift)
}

// SetArr Low Auto-reload value
func (r *registerArrType) SetArr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArrFieldArrMask)|(uint32(value)<<RegisterArrFieldArrShift))
}

// registerCcr1Type capture/compare register 1
type registerCcr1Type uint32

const (
	RegisterCcr1FieldCcr1Shift = 0
	RegisterCcr1FieldCcr1Mask  = 0xffff
)

// GetCcr1 Low Capture/Compare 1 value
func (r *registerCcr1Type) GetCcr1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldCcr1Mask) >> RegisterCcr1FieldCcr1Shift)
}

// SetCcr1 Low Capture/Compare 1 value
func (r *registerCcr1Type) SetCcr1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldCcr1Mask)|(uint32(value)<<RegisterCcr1FieldCcr1Shift))
}

// registerCcr2Type capture/compare register 2
type registerCcr2Type uint32

const (
	RegisterCcr2FieldCcr2Shift = 0
	RegisterCcr2FieldCcr2Mask  = 0xffff
)

// GetCcr2 Low Capture/Compare 2 value
func (r *registerCcr2Type) GetCcr2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldCcr2Mask) >> RegisterCcr2FieldCcr2Shift)
}

// SetCcr2 Low Capture/Compare 2 value
func (r *registerCcr2Type) SetCcr2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldCcr2Mask)|(uint32(value)<<RegisterCcr2FieldCcr2Shift))
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
