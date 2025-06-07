//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package tim15

import (
	"unsafe"
	"volatile"
)

var (
	Tim15 = (*_tim15)(unsafe.Pointer(uintptr(0x40014000)))
)

type _tim15 struct {
	Cr1         RegisterCr1Type
	Cr2         RegisterCr2Type
	Smcr        RegisterSmcrType
	Dier        RegisterDierType
	Sr          RegisterSrType
	Egr         RegisterEgrType
	Ccmr1output RegisterCcmr1outputType
	Ccmr1input  RegisterCcmr1inputType
	_           [4]byte
	Ccer        RegisterCcerType
	Cnt         RegisterCntType
	Psc         RegisterPscType
	Arr         RegisterArrType
	Rcr         RegisterRcrType
	Ccr1        RegisterCcr1Type
	Ccr2        RegisterCcr2Type
	_           [8]byte
	Bdtr        RegisterBdtrType
	Dcr         RegisterDcrType
	Dmar        RegisterDmarType
	_           [16]byte
	Af          RegisterAfType
	_           [4]byte
	Tisel       RegisterTiselType
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
	RegisterCr2FieldCcpcShift = 0
	RegisterCr2FieldCcpcMask  = 0x1
)

// GetCcpc Capture/compare preloaded control
func (r *RegisterCr2Type) GetCcpc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCcpcMask) != 0
}

// SetCcpc Capture/compare preloaded control
func (r *RegisterCr2Type) SetCcpc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldCcpcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldCcpcMask)
	}
}

const (
	RegisterCr2FieldCcusShift = 2
	RegisterCr2FieldCcusMask  = 0x4
)

// GetCcus Capture/compare control update selection
func (r *RegisterCr2Type) GetCcus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCcusMask) != 0
}

// SetCcus Capture/compare control update selection
func (r *RegisterCr2Type) SetCcus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldCcusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldCcusMask)
	}
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

const (
	RegisterCr2FieldOis1Shift = 8
	RegisterCr2FieldOis1Mask  = 0x100
)

// GetOis1 Output Idle state 1
func (r *RegisterCr2Type) GetOis1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis1Mask) != 0
}

// SetOis1 Output Idle state 1
func (r *RegisterCr2Type) SetOis1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldOis1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldOis1Mask)
	}
}

const (
	RegisterCr2FieldOis1nShift = 9
	RegisterCr2FieldOis1nMask  = 0x200
)

// GetOis1n Output Idle state 1
func (r *RegisterCr2Type) GetOis1n() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis1nMask) != 0
}

// SetOis1n Output Idle state 1
func (r *RegisterCr2Type) SetOis1n(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldOis1nMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldOis1nMask)
	}
}

const (
	RegisterCr2FieldOis2Shift = 10
	RegisterCr2FieldOis2Mask  = 0x400
)

// GetOis2 Output Idle state 2
func (r *RegisterCr2Type) GetOis2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis2Mask) != 0
}

// SetOis2 Output Idle state 2
func (r *RegisterCr2Type) SetOis2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldOis2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldOis2Mask)
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
	RegisterSmcrFieldTs20Shift = 4
	RegisterSmcrFieldTs20Mask  = 0x70
)

// GetTs20 Trigger selection
func (r *RegisterSmcrType) GetTs20() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldTs20Mask) >> RegisterSmcrFieldTs20Shift)
}

// SetTs20 Trigger selection
func (r *RegisterSmcrType) SetTs20(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldTs20Mask)|(uint32(value)<<RegisterSmcrFieldTs20Shift))
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
	RegisterSmcrFieldSms3Shift = 16
	RegisterSmcrFieldSms3Mask  = 0x10000
)

// GetSms3 Slave mode selection bit 3
func (r *RegisterSmcrType) GetSms3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldSms3Mask) != 0
}

// SetSms3 Slave mode selection bit 3
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

// GetTs43 Trigger selection - bit 4:3
func (r *RegisterSmcrType) GetTs43() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldTs43Mask) >> RegisterSmcrFieldTs43Shift)
}

// SetTs43 Trigger selection - bit 4:3
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
	RegisterDierFieldComieShift = 5
	RegisterDierFieldComieMask  = 0x20
)

// GetComie COM interrupt enable
func (r *RegisterDierType) GetComie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldComieMask) != 0
}

// SetComie COM interrupt enable
func (r *RegisterDierType) SetComie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldComieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldComieMask)
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
	RegisterDierFieldBieShift = 7
	RegisterDierFieldBieMask  = 0x80
)

// GetBie Break interrupt enable
func (r *RegisterDierType) GetBie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldBieMask) != 0
}

// SetBie Break interrupt enable
func (r *RegisterDierType) SetBie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldBieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldBieMask)
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
	RegisterDierFieldComdeShift = 13
	RegisterDierFieldComdeMask  = 0x2000
)

// GetComde COM DMA request enable
func (r *RegisterDierType) GetComde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldComdeMask) != 0
}

// SetComde COM DMA request enable
func (r *RegisterDierType) SetComde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldComdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldComdeMask)
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
	RegisterSrFieldComifShift = 5
	RegisterSrFieldComifMask  = 0x20
)

// GetComif COM interrupt flag
func (r *RegisterSrType) GetComif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldComifMask) != 0
}

// SetComif COM interrupt flag
func (r *RegisterSrType) SetComif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldComifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldComifMask)
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
	RegisterSrFieldBifShift = 7
	RegisterSrFieldBifMask  = 0x80
)

// GetBif Break interrupt flag
func (r *RegisterSrType) GetBif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldBifMask) != 0
}

// SetBif Break interrupt flag
func (r *RegisterSrType) SetBif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldBifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldBifMask)
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
	RegisterEgrFieldComgShift = 5
	RegisterEgrFieldComgMask  = 0x20
)

// GetComg Capture/Compare control update generation
func (r *RegisterEgrType) GetComg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldComgMask) != 0
}

// SetComg Capture/Compare control update generation
func (r *RegisterEgrType) SetComg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldComgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldComgMask)
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

const (
	RegisterEgrFieldBgShift = 7
	RegisterEgrFieldBgMask  = 0x80
)

// GetBg Break generation
func (r *RegisterEgrType) GetBg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldBgMask) != 0
}

// SetBg Break generation
func (r *RegisterEgrType) SetBg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldBgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldBgMask)
	}
}

// RegisterCcmr1outputType capture/compare mode register (output mode)
type RegisterCcmr1outputType uint32

func (r *RegisterCcmr1outputType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcmr1outputType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcmr1outputType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcmr1outputType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcmr1outputType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcmr1outputFieldCc1sShift = 0
	RegisterCcmr1outputFieldCc1sMask  = 0x3
)

// GetCc1s Capture/Compare 1 selection
func (r *RegisterCcmr1outputType) GetCc1s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldCc1sMask) >> RegisterCcmr1outputFieldCc1sShift)
}

// SetCc1s Capture/Compare 1 selection
func (r *RegisterCcmr1outputType) SetCc1s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1outputFieldCc1sMask)|(uint32(value)<<RegisterCcmr1outputFieldCc1sShift))
}

const (
	RegisterCcmr1outputFieldOc1feShift = 2
	RegisterCcmr1outputFieldOc1feMask  = 0x4
)

// GetOc1fe Output Compare 1 fast enable
func (r *RegisterCcmr1outputType) GetOc1fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldOc1feMask) != 0
}

// SetOc1fe Output Compare 1 fast enable
func (r *RegisterCcmr1outputType) SetOc1fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1outputFieldOc1feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1outputFieldOc1feMask)
	}
}

const (
	RegisterCcmr1outputFieldOc1peShift = 3
	RegisterCcmr1outputFieldOc1peMask  = 0x8
)

// GetOc1pe Output Compare 1 preload enable
func (r *RegisterCcmr1outputType) GetOc1pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldOc1peMask) != 0
}

// SetOc1pe Output Compare 1 preload enable
func (r *RegisterCcmr1outputType) SetOc1pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1outputFieldOc1peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1outputFieldOc1peMask)
	}
}

const (
	RegisterCcmr1outputFieldOc1mShift = 4
	RegisterCcmr1outputFieldOc1mMask  = 0x70
)

// GetOc1m Output Compare 1 mode
func (r *RegisterCcmr1outputType) GetOc1m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldOc1mMask) >> RegisterCcmr1outputFieldOc1mShift)
}

// SetOc1m Output Compare 1 mode
func (r *RegisterCcmr1outputType) SetOc1m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1outputFieldOc1mMask)|(uint32(value)<<RegisterCcmr1outputFieldOc1mShift))
}

const (
	RegisterCcmr1outputFieldCc2sShift = 8
	RegisterCcmr1outputFieldCc2sMask  = 0x300
)

// GetCc2s Capture/Compare 2 selection
func (r *RegisterCcmr1outputType) GetCc2s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldCc2sMask) >> RegisterCcmr1outputFieldCc2sShift)
}

// SetCc2s Capture/Compare 2 selection
func (r *RegisterCcmr1outputType) SetCc2s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1outputFieldCc2sMask)|(uint32(value)<<RegisterCcmr1outputFieldCc2sShift))
}

const (
	RegisterCcmr1outputFieldOc2feShift = 10
	RegisterCcmr1outputFieldOc2feMask  = 0x400
)

// GetOc2fe Output Compare 2 fast enable
func (r *RegisterCcmr1outputType) GetOc2fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldOc2feMask) != 0
}

// SetOc2fe Output Compare 2 fast enable
func (r *RegisterCcmr1outputType) SetOc2fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1outputFieldOc2feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1outputFieldOc2feMask)
	}
}

const (
	RegisterCcmr1outputFieldOc2peShift = 11
	RegisterCcmr1outputFieldOc2peMask  = 0x800
)

// GetOc2pe Output Compare 2 preload enable
func (r *RegisterCcmr1outputType) GetOc2pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldOc2peMask) != 0
}

// SetOc2pe Output Compare 2 preload enable
func (r *RegisterCcmr1outputType) SetOc2pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1outputFieldOc2peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1outputFieldOc2peMask)
	}
}

const (
	RegisterCcmr1outputFieldOc2mShift = 12
	RegisterCcmr1outputFieldOc2mMask  = 0x7000
)

// GetOc2m Output Compare 2 mode
func (r *RegisterCcmr1outputType) GetOc2m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldOc2mMask) >> RegisterCcmr1outputFieldOc2mShift)
}

// SetOc2m Output Compare 2 mode
func (r *RegisterCcmr1outputType) SetOc2m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1outputFieldOc2mMask)|(uint32(value)<<RegisterCcmr1outputFieldOc2mShift))
}

const (
	RegisterCcmr1outputFieldOc1m3Shift = 16
	RegisterCcmr1outputFieldOc1m3Mask  = 0x10000
)

// GetOc1m3 Output Compare 1 mode bit 3
func (r *RegisterCcmr1outputType) GetOc1m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldOc1m3Mask) != 0
}

// SetOc1m3 Output Compare 1 mode bit 3
func (r *RegisterCcmr1outputType) SetOc1m3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1outputFieldOc1m3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1outputFieldOc1m3Mask)
	}
}

const (
	RegisterCcmr1outputFieldOc2m3Shift = 24
	RegisterCcmr1outputFieldOc2m3Mask  = 0x1000000
)

// GetOc2m3 Output Compare 2 mode bit 3
func (r *RegisterCcmr1outputType) GetOc2m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldOc2m3Mask) != 0
}

// SetOc2m3 Output Compare 2 mode bit 3
func (r *RegisterCcmr1outputType) SetOc2m3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1outputFieldOc2m3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1outputFieldOc2m3Mask)
	}
}

// RegisterCcmr1inputType capture/compare mode register 1 (input mode)
type RegisterCcmr1inputType uint32

func (r *RegisterCcmr1inputType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcmr1inputType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcmr1inputType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcmr1inputType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcmr1inputType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcmr1inputFieldCc1sShift = 0
	RegisterCcmr1inputFieldCc1sMask  = 0x3
)

// GetCc1s Capture/Compare 1 selection
func (r *RegisterCcmr1inputType) GetCc1s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1inputFieldCc1sMask) >> RegisterCcmr1inputFieldCc1sShift)
}

// SetCc1s Capture/Compare 1 selection
func (r *RegisterCcmr1inputType) SetCc1s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1inputFieldCc1sMask)|(uint32(value)<<RegisterCcmr1inputFieldCc1sShift))
}

const (
	RegisterCcmr1inputFieldIc1pscShift = 2
	RegisterCcmr1inputFieldIc1pscMask  = 0xc
)

// GetIc1psc Input capture 1 prescaler
func (r *RegisterCcmr1inputType) GetIc1psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1inputFieldIc1pscMask) >> RegisterCcmr1inputFieldIc1pscShift)
}

// SetIc1psc Input capture 1 prescaler
func (r *RegisterCcmr1inputType) SetIc1psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1inputFieldIc1pscMask)|(uint32(value)<<RegisterCcmr1inputFieldIc1pscShift))
}

const (
	RegisterCcmr1inputFieldIc1fShift = 4
	RegisterCcmr1inputFieldIc1fMask  = 0xf0
)

// GetIc1f Input capture 1 filter
func (r *RegisterCcmr1inputType) GetIc1f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1inputFieldIc1fMask) >> RegisterCcmr1inputFieldIc1fShift)
}

// SetIc1f Input capture 1 filter
func (r *RegisterCcmr1inputType) SetIc1f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1inputFieldIc1fMask)|(uint32(value)<<RegisterCcmr1inputFieldIc1fShift))
}

const (
	RegisterCcmr1inputFieldCc2sShift = 8
	RegisterCcmr1inputFieldCc2sMask  = 0x300
)

// GetCc2s Capture/Compare 2 selection
func (r *RegisterCcmr1inputType) GetCc2s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1inputFieldCc2sMask) >> RegisterCcmr1inputFieldCc2sShift)
}

// SetCc2s Capture/Compare 2 selection
func (r *RegisterCcmr1inputType) SetCc2s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1inputFieldCc2sMask)|(uint32(value)<<RegisterCcmr1inputFieldCc2sShift))
}

const (
	RegisterCcmr1inputFieldIc2pscShift = 10
	RegisterCcmr1inputFieldIc2pscMask  = 0xc00
)

// GetIc2psc Input capture 2 prescaler
func (r *RegisterCcmr1inputType) GetIc2psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1inputFieldIc2pscMask) >> RegisterCcmr1inputFieldIc2pscShift)
}

// SetIc2psc Input capture 2 prescaler
func (r *RegisterCcmr1inputType) SetIc2psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1inputFieldIc2pscMask)|(uint32(value)<<RegisterCcmr1inputFieldIc2pscShift))
}

const (
	RegisterCcmr1inputFieldIc2fShift = 12
	RegisterCcmr1inputFieldIc2fMask  = 0xf000
)

// GetIc2f Input capture 2 filter
func (r *RegisterCcmr1inputType) GetIc2f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1inputFieldIc2fMask) >> RegisterCcmr1inputFieldIc2fShift)
}

// SetIc2f Input capture 2 filter
func (r *RegisterCcmr1inputType) SetIc2f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1inputFieldIc2fMask)|(uint32(value)<<RegisterCcmr1inputFieldIc2fShift))
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
	RegisterCcerFieldCc1neShift = 2
	RegisterCcerFieldCc1neMask  = 0x4
)

// GetCc1ne Capture/Compare 1 complementary output enable
func (r *RegisterCcerType) GetCc1ne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc1neMask) != 0
}

// SetCc1ne Capture/Compare 1 complementary output enable
func (r *RegisterCcerType) SetCc1ne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc1neMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc1neMask)
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
	RegisterCntFieldCntShift = 0
	RegisterCntFieldCntMask  = 0xffff
)

// GetCnt counter value
func (r *RegisterCntType) GetCnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldCntMask) >> RegisterCntFieldCntShift)
}

// SetCnt counter value
func (r *RegisterCntType) SetCnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldCntMask)|(uint32(value)<<RegisterCntFieldCntShift))
}

const (
	RegisterCntFieldUifcpyShift = 31
	RegisterCntFieldUifcpyMask  = 0x80000000
)

// GetUifcpy UIF copy
func (r *RegisterCntType) GetUifcpy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldUifcpyMask) != 0
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
	RegisterArrFieldArrShift = 0
	RegisterArrFieldArrMask  = 0xffff
)

// GetArr Auto-reload value
func (r *RegisterArrType) GetArr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterArrFieldArrMask) >> RegisterArrFieldArrShift)
}

// SetArr Auto-reload value
func (r *RegisterArrType) SetArr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArrFieldArrMask)|(uint32(value)<<RegisterArrFieldArrShift))
}

// RegisterRcrType repetition counter register
type RegisterRcrType uint32

func (r *RegisterRcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRcrFieldRepShift = 0
	RegisterRcrFieldRepMask  = 0xff
)

// GetRep Repetition counter value
func (r *RegisterRcrType) GetRep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRcrFieldRepMask) >> RegisterRcrFieldRepShift)
}

// SetRep Repetition counter value
func (r *RegisterRcrType) SetRep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRcrFieldRepMask)|(uint32(value)<<RegisterRcrFieldRepShift))
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
	RegisterCcr1FieldCcr1Shift = 0
	RegisterCcr1FieldCcr1Mask  = 0xffff
)

// GetCcr1 Capture/Compare 1 value
func (r *RegisterCcr1Type) GetCcr1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldCcr1Mask) >> RegisterCcr1FieldCcr1Shift)
}

// SetCcr1 Capture/Compare 1 value
func (r *RegisterCcr1Type) SetCcr1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldCcr1Mask)|(uint32(value)<<RegisterCcr1FieldCcr1Shift))
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
	RegisterCcr2FieldCcr2Shift = 0
	RegisterCcr2FieldCcr2Mask  = 0xffff
)

// GetCcr2 Capture/Compare 2 value
func (r *RegisterCcr2Type) GetCcr2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldCcr2Mask) >> RegisterCcr2FieldCcr2Shift)
}

// SetCcr2 Capture/Compare 2 value
func (r *RegisterCcr2Type) SetCcr2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldCcr2Mask)|(uint32(value)<<RegisterCcr2FieldCcr2Shift))
}

// RegisterBdtrType break and dead-time register
type RegisterBdtrType uint32

func (r *RegisterBdtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBdtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBdtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBdtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBdtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBdtrFieldDtgShift = 0
	RegisterBdtrFieldDtgMask  = 0xff
)

// GetDtg Dead-time generator setup
func (r *RegisterBdtrType) GetDtg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldDtgMask) >> RegisterBdtrFieldDtgShift)
}

// SetDtg Dead-time generator setup
func (r *RegisterBdtrType) SetDtg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldDtgMask)|(uint32(value)<<RegisterBdtrFieldDtgShift))
}

const (
	RegisterBdtrFieldLockShift = 8
	RegisterBdtrFieldLockMask  = 0x300
)

// GetLock Lock configuration
func (r *RegisterBdtrType) GetLock() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldLockMask) >> RegisterBdtrFieldLockShift)
}

// SetLock Lock configuration
func (r *RegisterBdtrType) SetLock(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldLockMask)|(uint32(value)<<RegisterBdtrFieldLockShift))
}

const (
	RegisterBdtrFieldOssiShift = 10
	RegisterBdtrFieldOssiMask  = 0x400
)

// GetOssi Off-state selection for Idle mode
func (r *RegisterBdtrType) GetOssi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldOssiMask) != 0
}

// SetOssi Off-state selection for Idle mode
func (r *RegisterBdtrType) SetOssi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtrFieldOssiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldOssiMask)
	}
}

const (
	RegisterBdtrFieldOssrShift = 11
	RegisterBdtrFieldOssrMask  = 0x800
)

// GetOssr Off-state selection for Run mode
func (r *RegisterBdtrType) GetOssr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldOssrMask) != 0
}

// SetOssr Off-state selection for Run mode
func (r *RegisterBdtrType) SetOssr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtrFieldOssrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldOssrMask)
	}
}

const (
	RegisterBdtrFieldBkeShift = 12
	RegisterBdtrFieldBkeMask  = 0x1000
)

// GetBke Break enable
func (r *RegisterBdtrType) GetBke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldBkeMask) != 0
}

// SetBke Break enable
func (r *RegisterBdtrType) SetBke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtrFieldBkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldBkeMask)
	}
}

const (
	RegisterBdtrFieldBkpShift = 13
	RegisterBdtrFieldBkpMask  = 0x2000
)

// GetBkp Break polarity
func (r *RegisterBdtrType) GetBkp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldBkpMask) != 0
}

// SetBkp Break polarity
func (r *RegisterBdtrType) SetBkp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtrFieldBkpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldBkpMask)
	}
}

const (
	RegisterBdtrFieldAoeShift = 14
	RegisterBdtrFieldAoeMask  = 0x4000
)

// GetAoe Automatic output enable
func (r *RegisterBdtrType) GetAoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldAoeMask) != 0
}

// SetAoe Automatic output enable
func (r *RegisterBdtrType) SetAoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtrFieldAoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldAoeMask)
	}
}

const (
	RegisterBdtrFieldMoeShift = 15
	RegisterBdtrFieldMoeMask  = 0x8000
)

// GetMoe Main output enable
func (r *RegisterBdtrType) GetMoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldMoeMask) != 0
}

// SetMoe Main output enable
func (r *RegisterBdtrType) SetMoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtrFieldMoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldMoeMask)
	}
}

const (
	RegisterBdtrFieldBkfShift = 16
	RegisterBdtrFieldBkfMask  = 0xf0000
)

// GetBkf Break filter
func (r *RegisterBdtrType) GetBkf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldBkfMask) >> RegisterBdtrFieldBkfShift)
}

// SetBkf Break filter
func (r *RegisterBdtrType) SetBkf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldBkfMask)|(uint32(value)<<RegisterBdtrFieldBkfShift))
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

// RegisterAfType alternate fdfsdm1_breakon register 1
type RegisterAfType uint32

func (r *RegisterAfType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAfType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAfType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAfType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAfType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAfFieldBkineShift = 0
	RegisterAfFieldBkineMask  = 0x1
)

// GetBkine BRK BKIN input enable
func (r *RegisterAfType) GetBkine() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfFieldBkineMask) != 0
}

// SetBkine BRK BKIN input enable
func (r *RegisterAfType) SetBkine(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfFieldBkineMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfFieldBkineMask)
	}
}

const (
	RegisterAfFieldBkcmp1eShift = 1
	RegisterAfFieldBkcmp1eMask  = 0x2
)

// GetBkcmp1e BRK COMP1 enable
func (r *RegisterAfType) GetBkcmp1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfFieldBkcmp1eMask) != 0
}

// SetBkcmp1e BRK COMP1 enable
func (r *RegisterAfType) SetBkcmp1e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfFieldBkcmp1eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfFieldBkcmp1eMask)
	}
}

const (
	RegisterAfFieldBkcmp2eShift = 2
	RegisterAfFieldBkcmp2eMask  = 0x4
)

// GetBkcmp2e BRK COMP2 enable
func (r *RegisterAfType) GetBkcmp2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfFieldBkcmp2eMask) != 0
}

// SetBkcmp2e BRK COMP2 enable
func (r *RegisterAfType) SetBkcmp2e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfFieldBkcmp2eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfFieldBkcmp2eMask)
	}
}

const (
	RegisterAfFieldBkdf1bk0eShift = 8
	RegisterAfFieldBkdf1bk0eMask  = 0x100
)

// GetBkdf1bk0e BRK dfsdm1_break[0] enable
func (r *RegisterAfType) GetBkdf1bk0e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfFieldBkdf1bk0eMask) != 0
}

// SetBkdf1bk0e BRK dfsdm1_break[0] enable
func (r *RegisterAfType) SetBkdf1bk0e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfFieldBkdf1bk0eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfFieldBkdf1bk0eMask)
	}
}

const (
	RegisterAfFieldBkinpShift = 9
	RegisterAfFieldBkinpMask  = 0x200
)

// GetBkinp BRK BKIN input polarity
func (r *RegisterAfType) GetBkinp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfFieldBkinpMask) != 0
}

// SetBkinp BRK BKIN input polarity
func (r *RegisterAfType) SetBkinp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfFieldBkinpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfFieldBkinpMask)
	}
}

const (
	RegisterAfFieldBkcmp1pShift = 10
	RegisterAfFieldBkcmp1pMask  = 0x400
)

// GetBkcmp1p BRK COMP1 input polarity
func (r *RegisterAfType) GetBkcmp1p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfFieldBkcmp1pMask) != 0
}

// SetBkcmp1p BRK COMP1 input polarity
func (r *RegisterAfType) SetBkcmp1p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfFieldBkcmp1pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfFieldBkcmp1pMask)
	}
}

const (
	RegisterAfFieldBkcmp2pShift = 11
	RegisterAfFieldBkcmp2pMask  = 0x800
)

// GetBkcmp2p BRK COMP2 input polarity
func (r *RegisterAfType) GetBkcmp2p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfFieldBkcmp2pMask) != 0
}

// SetBkcmp2p BRK COMP2 input polarity
func (r *RegisterAfType) SetBkcmp2p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfFieldBkcmp2pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfFieldBkcmp2pMask)
	}
}

// RegisterTiselType input selection register
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

// GetTi1sel selects TI1[0] to TI1[15] input
func (r *RegisterTiselType) GetTi1sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi1selMask) >> RegisterTiselFieldTi1selShift)
}

// SetTi1sel selects TI1[0] to TI1[15] input
func (r *RegisterTiselType) SetTi1sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi1selMask)|(uint32(value)<<RegisterTiselFieldTi1selShift))
}

const (
	RegisterTiselFieldTi2selShift = 8
	RegisterTiselFieldTi2selMask  = 0xf00
)

// GetTi2sel selects TI2[0] to TI2[15] input
func (r *RegisterTiselType) GetTi2sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi2selMask) >> RegisterTiselFieldTi2selShift)
}

// SetTi2sel selects TI2[0] to TI2[15] input
func (r *RegisterTiselType) SetTi2sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi2selMask)|(uint32(value)<<RegisterTiselFieldTi2selShift))
}
