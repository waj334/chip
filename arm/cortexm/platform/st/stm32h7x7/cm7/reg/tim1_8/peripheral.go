//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package tim1_8

import (
	"unsafe"
	"volatile"
)

var (
	Tim1 = (*_tim1_8)(unsafe.Pointer(uintptr(0x40010000)))
	Tim8 = (*_tim1_8)(unsafe.Pointer(uintptr(0x40010400)))

	Instances = [2]*_tim1_8{
		Tim1,
		Tim8,
	}
)

type _tim1_8 struct {
	Cr1         RegisterCr1Type
	Cr2         RegisterCr2Type
	Smcr        RegisterSmcrType
	Dier        RegisterDierType
	Sr          RegisterSrType
	Egr         RegisterEgrType
	Ccmr1output RegisterCcmr1outputType
	Ccmr1input  RegisterCcmr1inputType
	Ccmr2output RegisterCcmr2outputType
	Ccmr2input  RegisterCcmr2inputType
	Ccer        RegisterCcerType
	Cnt         RegisterCntType
	Psc         RegisterPscType
	Arr         RegisterArrType
	Rcr         RegisterRcrType
	Ccr1        RegisterCcr1Type
	Ccr2        RegisterCcr2Type
	Ccr3        RegisterCcr3Type
	Ccr4        RegisterCcr4Type
	Bdtr        RegisterBdtrType
	Dcr         RegisterDcrType
	Dmar        RegisterDmarType
	_           [4]byte
	Ccmr3output RegisterCcmr3outputType
	Ccr5        RegisterCcr5Type
	Ccr6        RegisterCcr6Type
	Af1         RegisterAf1Type
	Af2         RegisterAf2Type
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

const (
	RegisterCr2FieldOis2nShift = 11
	RegisterCr2FieldOis2nMask  = 0x800
)

// GetOis2n Output Idle state 2
func (r *RegisterCr2Type) GetOis2n() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis2nMask) != 0
}

// SetOis2n Output Idle state 2
func (r *RegisterCr2Type) SetOis2n(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldOis2nMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldOis2nMask)
	}
}

const (
	RegisterCr2FieldOis3Shift = 12
	RegisterCr2FieldOis3Mask  = 0x1000
)

// GetOis3 Output Idle state 3
func (r *RegisterCr2Type) GetOis3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis3Mask) != 0
}

// SetOis3 Output Idle state 3
func (r *RegisterCr2Type) SetOis3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldOis3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldOis3Mask)
	}
}

const (
	RegisterCr2FieldOis3nShift = 13
	RegisterCr2FieldOis3nMask  = 0x2000
)

// GetOis3n Output Idle state 3
func (r *RegisterCr2Type) GetOis3n() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis3nMask) != 0
}

// SetOis3n Output Idle state 3
func (r *RegisterCr2Type) SetOis3n(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldOis3nMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldOis3nMask)
	}
}

const (
	RegisterCr2FieldOis4Shift = 14
	RegisterCr2FieldOis4Mask  = 0x4000
)

// GetOis4 Output Idle state 4
func (r *RegisterCr2Type) GetOis4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis4Mask) != 0
}

// SetOis4 Output Idle state 4
func (r *RegisterCr2Type) SetOis4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldOis4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldOis4Mask)
	}
}

const (
	RegisterCr2FieldOis5Shift = 16
	RegisterCr2FieldOis5Mask  = 0x10000
)

// GetOis5 Output Idle state 5
func (r *RegisterCr2Type) GetOis5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis5Mask) != 0
}

// SetOis5 Output Idle state 5
func (r *RegisterCr2Type) SetOis5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldOis5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldOis5Mask)
	}
}

const (
	RegisterCr2FieldOis6Shift = 18
	RegisterCr2FieldOis6Mask  = 0x40000
)

// GetOis6 Output Idle state 6
func (r *RegisterCr2Type) GetOis6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis6Mask) != 0
}

// SetOis6 Output Idle state 6
func (r *RegisterCr2Type) SetOis6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldOis6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldOis6Mask)
	}
}

const (
	RegisterCr2FieldMms2Shift = 20
	RegisterCr2FieldMms2Mask  = 0xf00000
)

// GetMms2 Master mode selection 2
func (r *RegisterCr2Type) GetMms2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldMms2Mask) >> RegisterCr2FieldMms2Shift)
}

// SetMms2 Master mode selection 2
func (r *RegisterCr2Type) SetMms2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldMms2Mask)|(uint32(value)<<RegisterCr2FieldMms2Shift))
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
	RegisterSrFieldB2ifShift = 8
	RegisterSrFieldB2ifMask  = 0x100
)

// GetB2if Break 2 interrupt flag
func (r *RegisterSrType) GetB2if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldB2ifMask) != 0
}

// SetB2if Break 2 interrupt flag
func (r *RegisterSrType) SetB2if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldB2ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldB2ifMask)
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

const (
	RegisterSrFieldSbifShift = 13
	RegisterSrFieldSbifMask  = 0x2000
)

// GetSbif System Break interrupt flag
func (r *RegisterSrType) GetSbif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSbifMask) != 0
}

// SetSbif System Break interrupt flag
func (r *RegisterSrType) SetSbif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldSbifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldSbifMask)
	}
}

const (
	RegisterSrFieldCc5ifShift = 16
	RegisterSrFieldCc5ifMask  = 0x10000
)

// GetCc5if Compare 5 interrupt flag
func (r *RegisterSrType) GetCc5if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc5ifMask) != 0
}

// SetCc5if Compare 5 interrupt flag
func (r *RegisterSrType) SetCc5if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCc5ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCc5ifMask)
	}
}

const (
	RegisterSrFieldCc6ifShift = 17
	RegisterSrFieldCc6ifMask  = 0x20000
)

// GetCc6if Compare 6 interrupt flag
func (r *RegisterSrType) GetCc6if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc6ifMask) != 0
}

// SetCc6if Compare 6 interrupt flag
func (r *RegisterSrType) SetCc6if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCc6ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCc6ifMask)
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

const (
	RegisterEgrFieldB2gShift = 8
	RegisterEgrFieldB2gMask  = 0x100
)

// GetB2g Break 2 generation
func (r *RegisterEgrType) GetB2g() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldB2gMask) != 0
}

// SetB2g Break 2 generation
func (r *RegisterEgrType) SetB2g(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldB2gMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldB2gMask)
	}
}

// RegisterCcmr1outputType capture/compare mode register 1 (output mode)
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
	RegisterCcmr1outputFieldOc1ceShift = 7
	RegisterCcmr1outputFieldOc1ceMask  = 0x80
)

// GetOc1ce Output Compare 1 clear enable
func (r *RegisterCcmr1outputType) GetOc1ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldOc1ceMask) != 0
}

// SetOc1ce Output Compare 1 clear enable
func (r *RegisterCcmr1outputType) SetOc1ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1outputFieldOc1ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1outputFieldOc1ceMask)
	}
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
	RegisterCcmr1outputFieldOc2ceShift = 15
	RegisterCcmr1outputFieldOc2ceMask  = 0x8000
)

// GetOc2ce Output Compare 2 clear enable
func (r *RegisterCcmr1outputType) GetOc2ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldOc2ceMask) != 0
}

// SetOc2ce Output Compare 2 clear enable
func (r *RegisterCcmr1outputType) SetOc2ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1outputFieldOc2ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1outputFieldOc2ceMask)
	}
}

const (
	RegisterCcmr1outputFieldOc1m3Shift = 16
	RegisterCcmr1outputFieldOc1m3Mask  = 0x10000
)

// GetOc1m3 Output Compare 1 mode - bit 3
func (r *RegisterCcmr1outputType) GetOc1m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldOc1m3Mask) != 0
}

// SetOc1m3 Output Compare 1 mode - bit 3
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

// GetOc2m3 Output Compare 2 mode - bit 3
func (r *RegisterCcmr1outputType) GetOc2m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1outputFieldOc2m3Mask) != 0
}

// SetOc2m3 Output Compare 2 mode - bit 3
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
	RegisterCcmr1inputFieldIcpcsShift = 2
	RegisterCcmr1inputFieldIcpcsMask  = 0xc
)

// GetIcpcs Input capture 1 prescaler
func (r *RegisterCcmr1inputType) GetIcpcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1inputFieldIcpcsMask) >> RegisterCcmr1inputFieldIcpcsShift)
}

// SetIcpcs Input capture 1 prescaler
func (r *RegisterCcmr1inputType) SetIcpcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1inputFieldIcpcsMask)|(uint32(value)<<RegisterCcmr1inputFieldIcpcsShift))
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
	RegisterCcmr1inputFieldIc2pcsShift = 10
	RegisterCcmr1inputFieldIc2pcsMask  = 0xc00
)

// GetIc2pcs Input capture 2 prescaler
func (r *RegisterCcmr1inputType) GetIc2pcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1inputFieldIc2pcsMask) >> RegisterCcmr1inputFieldIc2pcsShift)
}

// SetIc2pcs Input capture 2 prescaler
func (r *RegisterCcmr1inputType) SetIc2pcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1inputFieldIc2pcsMask)|(uint32(value)<<RegisterCcmr1inputFieldIc2pcsShift))
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

// RegisterCcmr2outputType capture/compare mode register 2 (output mode)
type RegisterCcmr2outputType uint32

func (r *RegisterCcmr2outputType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcmr2outputType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcmr2outputType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcmr2outputType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcmr2outputType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcmr2outputFieldCc3sShift = 0
	RegisterCcmr2outputFieldCc3sMask  = 0x3
)

// GetCc3s Capture/Compare 3 selection
func (r *RegisterCcmr2outputType) GetCc3s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2outputFieldCc3sMask) >> RegisterCcmr2outputFieldCc3sShift)
}

// SetCc3s Capture/Compare 3 selection
func (r *RegisterCcmr2outputType) SetCc3s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2outputFieldCc3sMask)|(uint32(value)<<RegisterCcmr2outputFieldCc3sShift))
}

const (
	RegisterCcmr2outputFieldOc3feShift = 2
	RegisterCcmr2outputFieldOc3feMask  = 0x4
)

// GetOc3fe Output compare 3 fast enable
func (r *RegisterCcmr2outputType) GetOc3fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2outputFieldOc3feMask) != 0
}

// SetOc3fe Output compare 3 fast enable
func (r *RegisterCcmr2outputType) SetOc3fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2outputFieldOc3feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2outputFieldOc3feMask)
	}
}

const (
	RegisterCcmr2outputFieldOc3peShift = 3
	RegisterCcmr2outputFieldOc3peMask  = 0x8
)

// GetOc3pe Output compare 3 preload enable
func (r *RegisterCcmr2outputType) GetOc3pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2outputFieldOc3peMask) != 0
}

// SetOc3pe Output compare 3 preload enable
func (r *RegisterCcmr2outputType) SetOc3pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2outputFieldOc3peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2outputFieldOc3peMask)
	}
}

const (
	RegisterCcmr2outputFieldOc3mShift = 4
	RegisterCcmr2outputFieldOc3mMask  = 0x70
)

// GetOc3m Output compare 3 mode
func (r *RegisterCcmr2outputType) GetOc3m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2outputFieldOc3mMask) >> RegisterCcmr2outputFieldOc3mShift)
}

// SetOc3m Output compare 3 mode
func (r *RegisterCcmr2outputType) SetOc3m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2outputFieldOc3mMask)|(uint32(value)<<RegisterCcmr2outputFieldOc3mShift))
}

const (
	RegisterCcmr2outputFieldOc3ceShift = 7
	RegisterCcmr2outputFieldOc3ceMask  = 0x80
)

// GetOc3ce Output compare 3 clear enable
func (r *RegisterCcmr2outputType) GetOc3ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2outputFieldOc3ceMask) != 0
}

// SetOc3ce Output compare 3 clear enable
func (r *RegisterCcmr2outputType) SetOc3ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2outputFieldOc3ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2outputFieldOc3ceMask)
	}
}

const (
	RegisterCcmr2outputFieldCc4sShift = 8
	RegisterCcmr2outputFieldCc4sMask  = 0x300
)

// GetCc4s Capture/Compare 4 selection
func (r *RegisterCcmr2outputType) GetCc4s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2outputFieldCc4sMask) >> RegisterCcmr2outputFieldCc4sShift)
}

// SetCc4s Capture/Compare 4 selection
func (r *RegisterCcmr2outputType) SetCc4s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2outputFieldCc4sMask)|(uint32(value)<<RegisterCcmr2outputFieldCc4sShift))
}

const (
	RegisterCcmr2outputFieldOc4feShift = 10
	RegisterCcmr2outputFieldOc4feMask  = 0x400
)

// GetOc4fe Output compare 4 fast enable
func (r *RegisterCcmr2outputType) GetOc4fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2outputFieldOc4feMask) != 0
}

// SetOc4fe Output compare 4 fast enable
func (r *RegisterCcmr2outputType) SetOc4fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2outputFieldOc4feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2outputFieldOc4feMask)
	}
}

const (
	RegisterCcmr2outputFieldOc4peShift = 11
	RegisterCcmr2outputFieldOc4peMask  = 0x800
)

// GetOc4pe Output compare 4 preload enable
func (r *RegisterCcmr2outputType) GetOc4pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2outputFieldOc4peMask) != 0
}

// SetOc4pe Output compare 4 preload enable
func (r *RegisterCcmr2outputType) SetOc4pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2outputFieldOc4peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2outputFieldOc4peMask)
	}
}

const (
	RegisterCcmr2outputFieldOc4mShift = 12
	RegisterCcmr2outputFieldOc4mMask  = 0x7000
)

// GetOc4m Output compare 4 mode
func (r *RegisterCcmr2outputType) GetOc4m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2outputFieldOc4mMask) >> RegisterCcmr2outputFieldOc4mShift)
}

// SetOc4m Output compare 4 mode
func (r *RegisterCcmr2outputType) SetOc4m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2outputFieldOc4mMask)|(uint32(value)<<RegisterCcmr2outputFieldOc4mShift))
}

const (
	RegisterCcmr2outputFieldOc4ceShift = 15
	RegisterCcmr2outputFieldOc4ceMask  = 0x8000
)

// GetOc4ce Output compare 4 clear enable
func (r *RegisterCcmr2outputType) GetOc4ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2outputFieldOc4ceMask) != 0
}

// SetOc4ce Output compare 4 clear enable
func (r *RegisterCcmr2outputType) SetOc4ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2outputFieldOc4ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2outputFieldOc4ceMask)
	}
}

const (
	RegisterCcmr2outputFieldOc3m3Shift = 16
	RegisterCcmr2outputFieldOc3m3Mask  = 0x10000
)

// GetOc3m3 Output Compare 3 mode - bit 3
func (r *RegisterCcmr2outputType) GetOc3m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2outputFieldOc3m3Mask) != 0
}

// SetOc3m3 Output Compare 3 mode - bit 3
func (r *RegisterCcmr2outputType) SetOc3m3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2outputFieldOc3m3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2outputFieldOc3m3Mask)
	}
}

const (
	RegisterCcmr2outputFieldOc4m4Shift = 24
	RegisterCcmr2outputFieldOc4m4Mask  = 0x1000000
)

// GetOc4m4 Output Compare 4 mode - bit 3
func (r *RegisterCcmr2outputType) GetOc4m4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2outputFieldOc4m4Mask) != 0
}

// SetOc4m4 Output Compare 4 mode - bit 3
func (r *RegisterCcmr2outputType) SetOc4m4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2outputFieldOc4m4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2outputFieldOc4m4Mask)
	}
}

// RegisterCcmr2inputType capture/compare mode register 2 (input mode)
type RegisterCcmr2inputType uint32

func (r *RegisterCcmr2inputType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcmr2inputType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcmr2inputType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcmr2inputType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcmr2inputType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcmr2inputFieldCc3sShift = 0
	RegisterCcmr2inputFieldCc3sMask  = 0x3
)

// GetCc3s Capture/compare 3 selection
func (r *RegisterCcmr2inputType) GetCc3s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2inputFieldCc3sMask) >> RegisterCcmr2inputFieldCc3sShift)
}

// SetCc3s Capture/compare 3 selection
func (r *RegisterCcmr2inputType) SetCc3s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2inputFieldCc3sMask)|(uint32(value)<<RegisterCcmr2inputFieldCc3sShift))
}

const (
	RegisterCcmr2inputFieldIc3pscShift = 2
	RegisterCcmr2inputFieldIc3pscMask  = 0xc
)

// GetIc3psc Input capture 3 prescaler
func (r *RegisterCcmr2inputType) GetIc3psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2inputFieldIc3pscMask) >> RegisterCcmr2inputFieldIc3pscShift)
}

// SetIc3psc Input capture 3 prescaler
func (r *RegisterCcmr2inputType) SetIc3psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2inputFieldIc3pscMask)|(uint32(value)<<RegisterCcmr2inputFieldIc3pscShift))
}

const (
	RegisterCcmr2inputFieldIc3fShift = 4
	RegisterCcmr2inputFieldIc3fMask  = 0xf0
)

// GetIc3f Input capture 3 filter
func (r *RegisterCcmr2inputType) GetIc3f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2inputFieldIc3fMask) >> RegisterCcmr2inputFieldIc3fShift)
}

// SetIc3f Input capture 3 filter
func (r *RegisterCcmr2inputType) SetIc3f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2inputFieldIc3fMask)|(uint32(value)<<RegisterCcmr2inputFieldIc3fShift))
}

const (
	RegisterCcmr2inputFieldCc4sShift = 8
	RegisterCcmr2inputFieldCc4sMask  = 0x300
)

// GetCc4s Capture/Compare 4 selection
func (r *RegisterCcmr2inputType) GetCc4s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2inputFieldCc4sMask) >> RegisterCcmr2inputFieldCc4sShift)
}

// SetCc4s Capture/Compare 4 selection
func (r *RegisterCcmr2inputType) SetCc4s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2inputFieldCc4sMask)|(uint32(value)<<RegisterCcmr2inputFieldCc4sShift))
}

const (
	RegisterCcmr2inputFieldIc4pscShift = 10
	RegisterCcmr2inputFieldIc4pscMask  = 0xc00
)

// GetIc4psc Input capture 4 prescaler
func (r *RegisterCcmr2inputType) GetIc4psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2inputFieldIc4pscMask) >> RegisterCcmr2inputFieldIc4pscShift)
}

// SetIc4psc Input capture 4 prescaler
func (r *RegisterCcmr2inputType) SetIc4psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2inputFieldIc4pscMask)|(uint32(value)<<RegisterCcmr2inputFieldIc4pscShift))
}

const (
	RegisterCcmr2inputFieldIc4fShift = 12
	RegisterCcmr2inputFieldIc4fMask  = 0xf000
)

// GetIc4f Input capture 4 filter
func (r *RegisterCcmr2inputType) GetIc4f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2inputFieldIc4fMask) >> RegisterCcmr2inputFieldIc4fShift)
}

// SetIc4f Input capture 4 filter
func (r *RegisterCcmr2inputType) SetIc4f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2inputFieldIc4fMask)|(uint32(value)<<RegisterCcmr2inputFieldIc4fShift))
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
	RegisterCcerFieldCc2neShift = 6
	RegisterCcerFieldCc2neMask  = 0x40
)

// GetCc2ne Capture/Compare 2 complementary output enable
func (r *RegisterCcerType) GetCc2ne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc2neMask) != 0
}

// SetCc2ne Capture/Compare 2 complementary output enable
func (r *RegisterCcerType) SetCc2ne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc2neMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc2neMask)
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
	RegisterCcerFieldCc3neShift = 10
	RegisterCcerFieldCc3neMask  = 0x400
)

// GetCc3ne Capture/Compare 3 complementary output enable
func (r *RegisterCcerType) GetCc3ne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc3neMask) != 0
}

// SetCc3ne Capture/Compare 3 complementary output enable
func (r *RegisterCcerType) SetCc3ne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc3neMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc3neMask)
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

// GetCc4np Capture/Compare 4 complementary output polarity
func (r *RegisterCcerType) GetCc4np() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc4npMask) != 0
}

// SetCc4np Capture/Compare 4 complementary output polarity
func (r *RegisterCcerType) SetCc4np(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc4npMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc4npMask)
	}
}

const (
	RegisterCcerFieldCc5eShift = 16
	RegisterCcerFieldCc5eMask  = 0x10000
)

// GetCc5e Capture/Compare 5 output enable
func (r *RegisterCcerType) GetCc5e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc5eMask) != 0
}

// SetCc5e Capture/Compare 5 output enable
func (r *RegisterCcerType) SetCc5e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc5eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc5eMask)
	}
}

const (
	RegisterCcerFieldCc5pShift = 17
	RegisterCcerFieldCc5pMask  = 0x20000
)

// GetCc5p Capture/Compare 5 output polarity
func (r *RegisterCcerType) GetCc5p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc5pMask) != 0
}

// SetCc5p Capture/Compare 5 output polarity
func (r *RegisterCcerType) SetCc5p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc5pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc5pMask)
	}
}

const (
	RegisterCcerFieldCc6eShift = 20
	RegisterCcerFieldCc6eMask  = 0x100000
)

// GetCc6e Capture/Compare 6 output enable
func (r *RegisterCcerType) GetCc6e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc6eMask) != 0
}

// SetCc6e Capture/Compare 6 output enable
func (r *RegisterCcerType) SetCc6e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc6eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc6eMask)
	}
}

const (
	RegisterCcerFieldCc6pShift = 21
	RegisterCcerFieldCc6pMask  = 0x200000
)

// GetCc6p Capture/Compare 6 output polarity
func (r *RegisterCcerType) GetCc6p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc6pMask) != 0
}

// SetCc6p Capture/Compare 6 output polarity
func (r *RegisterCcerType) SetCc6p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc6pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc6pMask)
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
	RegisterCcr3FieldCcr3Shift = 0
	RegisterCcr3FieldCcr3Mask  = 0xffff
)

// GetCcr3 Capture/Compare value
func (r *RegisterCcr3Type) GetCcr3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldCcr3Mask) >> RegisterCcr3FieldCcr3Shift)
}

// SetCcr3 Capture/Compare value
func (r *RegisterCcr3Type) SetCcr3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldCcr3Mask)|(uint32(value)<<RegisterCcr3FieldCcr3Shift))
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
	RegisterCcr4FieldCcr4Shift = 0
	RegisterCcr4FieldCcr4Mask  = 0xffff
)

// GetCcr4 Capture/Compare value
func (r *RegisterCcr4Type) GetCcr4() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldCcr4Mask) >> RegisterCcr4FieldCcr4Shift)
}

// SetCcr4 Capture/Compare value
func (r *RegisterCcr4Type) SetCcr4(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldCcr4Mask)|(uint32(value)<<RegisterCcr4FieldCcr4Shift))
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

const (
	RegisterBdtrFieldBk2fShift = 20
	RegisterBdtrFieldBk2fMask  = 0xf00000
)

// GetBk2f Break 2 filter
func (r *RegisterBdtrType) GetBk2f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldBk2fMask) >> RegisterBdtrFieldBk2fShift)
}

// SetBk2f Break 2 filter
func (r *RegisterBdtrType) SetBk2f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldBk2fMask)|(uint32(value)<<RegisterBdtrFieldBk2fShift))
}

const (
	RegisterBdtrFieldBk2eShift = 24
	RegisterBdtrFieldBk2eMask  = 0x1000000
)

// GetBk2e Break 2 enable
func (r *RegisterBdtrType) GetBk2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldBk2eMask) != 0
}

// SetBk2e Break 2 enable
func (r *RegisterBdtrType) SetBk2e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtrFieldBk2eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldBk2eMask)
	}
}

const (
	RegisterBdtrFieldBk2pShift = 25
	RegisterBdtrFieldBk2pMask  = 0x2000000
)

// GetBk2p Break 2 polarity
func (r *RegisterBdtrType) GetBk2p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldBk2pMask) != 0
}

// SetBk2p Break 2 polarity
func (r *RegisterBdtrType) SetBk2p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtrFieldBk2pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldBk2pMask)
	}
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

// RegisterCcmr3outputType capture/compare mode register 3 (output mode)
type RegisterCcmr3outputType uint32

func (r *RegisterCcmr3outputType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcmr3outputType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcmr3outputType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcmr3outputType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcmr3outputType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcmr3outputFieldOc5feShift = 2
	RegisterCcmr3outputFieldOc5feMask  = 0x4
)

// GetOc5fe Output compare 5 fast enable
func (r *RegisterCcmr3outputType) GetOc5fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3outputFieldOc5feMask) != 0
}

// SetOc5fe Output compare 5 fast enable
func (r *RegisterCcmr3outputType) SetOc5fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3outputFieldOc5feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3outputFieldOc5feMask)
	}
}

const (
	RegisterCcmr3outputFieldOc5peShift = 3
	RegisterCcmr3outputFieldOc5peMask  = 0x8
)

// GetOc5pe Output compare 5 preload enable
func (r *RegisterCcmr3outputType) GetOc5pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3outputFieldOc5peMask) != 0
}

// SetOc5pe Output compare 5 preload enable
func (r *RegisterCcmr3outputType) SetOc5pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3outputFieldOc5peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3outputFieldOc5peMask)
	}
}

const (
	RegisterCcmr3outputFieldOc5mShift = 4
	RegisterCcmr3outputFieldOc5mMask  = 0x70
)

// GetOc5m Output compare 5 mode
func (r *RegisterCcmr3outputType) GetOc5m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3outputFieldOc5mMask) >> RegisterCcmr3outputFieldOc5mShift)
}

// SetOc5m Output compare 5 mode
func (r *RegisterCcmr3outputType) SetOc5m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3outputFieldOc5mMask)|(uint32(value)<<RegisterCcmr3outputFieldOc5mShift))
}

const (
	RegisterCcmr3outputFieldOc5ceShift = 7
	RegisterCcmr3outputFieldOc5ceMask  = 0x80
)

// GetOc5ce Output compare 5 clear enable
func (r *RegisterCcmr3outputType) GetOc5ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3outputFieldOc5ceMask) != 0
}

// SetOc5ce Output compare 5 clear enable
func (r *RegisterCcmr3outputType) SetOc5ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3outputFieldOc5ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3outputFieldOc5ceMask)
	}
}

const (
	RegisterCcmr3outputFieldOc6feShift = 10
	RegisterCcmr3outputFieldOc6feMask  = 0x400
)

// GetOc6fe Output compare 6 fast enable
func (r *RegisterCcmr3outputType) GetOc6fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3outputFieldOc6feMask) != 0
}

// SetOc6fe Output compare 6 fast enable
func (r *RegisterCcmr3outputType) SetOc6fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3outputFieldOc6feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3outputFieldOc6feMask)
	}
}

const (
	RegisterCcmr3outputFieldOc6peShift = 11
	RegisterCcmr3outputFieldOc6peMask  = 0x800
)

// GetOc6pe Output compare 6 preload enable
func (r *RegisterCcmr3outputType) GetOc6pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3outputFieldOc6peMask) != 0
}

// SetOc6pe Output compare 6 preload enable
func (r *RegisterCcmr3outputType) SetOc6pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3outputFieldOc6peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3outputFieldOc6peMask)
	}
}

const (
	RegisterCcmr3outputFieldOc6mShift = 12
	RegisterCcmr3outputFieldOc6mMask  = 0x7000
)

// GetOc6m Output compare 6 mode
func (r *RegisterCcmr3outputType) GetOc6m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3outputFieldOc6mMask) >> RegisterCcmr3outputFieldOc6mShift)
}

// SetOc6m Output compare 6 mode
func (r *RegisterCcmr3outputType) SetOc6m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3outputFieldOc6mMask)|(uint32(value)<<RegisterCcmr3outputFieldOc6mShift))
}

const (
	RegisterCcmr3outputFieldOc6ceShift = 15
	RegisterCcmr3outputFieldOc6ceMask  = 0x8000
)

// GetOc6ce Output compare 6 clear enable
func (r *RegisterCcmr3outputType) GetOc6ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3outputFieldOc6ceMask) != 0
}

// SetOc6ce Output compare 6 clear enable
func (r *RegisterCcmr3outputType) SetOc6ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3outputFieldOc6ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3outputFieldOc6ceMask)
	}
}

const (
	RegisterCcmr3outputFieldOc5m3Shift = 16
	RegisterCcmr3outputFieldOc5m3Mask  = 0x10000
)

// GetOc5m3 Output Compare 5 mode
func (r *RegisterCcmr3outputType) GetOc5m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3outputFieldOc5m3Mask) != 0
}

// SetOc5m3 Output Compare 5 mode
func (r *RegisterCcmr3outputType) SetOc5m3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3outputFieldOc5m3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3outputFieldOc5m3Mask)
	}
}

const (
	RegisterCcmr3outputFieldOc6m3Shift = 24
	RegisterCcmr3outputFieldOc6m3Mask  = 0x1000000
)

// GetOc6m3 Output Compare 6 mode
func (r *RegisterCcmr3outputType) GetOc6m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3outputFieldOc6m3Mask) != 0
}

// SetOc6m3 Output Compare 6 mode
func (r *RegisterCcmr3outputType) SetOc6m3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3outputFieldOc6m3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3outputFieldOc6m3Mask)
	}
}

// RegisterCcr5Type capture/compare register 5
type RegisterCcr5Type uint32

func (r *RegisterCcr5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcr5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcr5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcr5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcr5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcr5FieldCcr5Shift = 0
	RegisterCcr5FieldCcr5Mask  = 0xffff
)

// GetCcr5 Capture/Compare 5 value
func (r *RegisterCcr5Type) GetCcr5() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldCcr5Mask) >> RegisterCcr5FieldCcr5Shift)
}

// SetCcr5 Capture/Compare 5 value
func (r *RegisterCcr5Type) SetCcr5(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldCcr5Mask)|(uint32(value)<<RegisterCcr5FieldCcr5Shift))
}

const (
	RegisterCcr5FieldGc5c1Shift = 29
	RegisterCcr5FieldGc5c1Mask  = 0x20000000
)

// GetGc5c1 Group Channel 5 and Channel 1
func (r *RegisterCcr5Type) GetGc5c1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldGc5c1Mask) != 0
}

// SetGc5c1 Group Channel 5 and Channel 1
func (r *RegisterCcr5Type) SetGc5c1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldGc5c1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldGc5c1Mask)
	}
}

const (
	RegisterCcr5FieldGc5c2Shift = 30
	RegisterCcr5FieldGc5c2Mask  = 0x40000000
)

// GetGc5c2 Group Channel 5 and Channel 2
func (r *RegisterCcr5Type) GetGc5c2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldGc5c2Mask) != 0
}

// SetGc5c2 Group Channel 5 and Channel 2
func (r *RegisterCcr5Type) SetGc5c2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldGc5c2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldGc5c2Mask)
	}
}

const (
	RegisterCcr5FieldGc5c3Shift = 31
	RegisterCcr5FieldGc5c3Mask  = 0x80000000
)

// GetGc5c3 Group Channel 5 and Channel 3
func (r *RegisterCcr5Type) GetGc5c3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldGc5c3Mask) != 0
}

// SetGc5c3 Group Channel 5 and Channel 3
func (r *RegisterCcr5Type) SetGc5c3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldGc5c3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldGc5c3Mask)
	}
}

// RegisterCcr6Type capture/compare register 6
type RegisterCcr6Type uint32

func (r *RegisterCcr6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcr6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcr6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcr6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcr6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcr6FieldCcr6Shift = 0
	RegisterCcr6FieldCcr6Mask  = 0xffff
)

// GetCcr6 Capture/Compare 6 value
func (r *RegisterCcr6Type) GetCcr6() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldCcr6Mask) >> RegisterCcr6FieldCcr6Shift)
}

// SetCcr6 Capture/Compare 6 value
func (r *RegisterCcr6Type) SetCcr6(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldCcr6Mask)|(uint32(value)<<RegisterCcr6FieldCcr6Shift))
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
	RegisterAf1FieldBkineShift = 0
	RegisterAf1FieldBkineMask  = 0x1
)

// GetBkine BRK BKIN input enable
func (r *RegisterAf1Type) GetBkine() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkineMask) != 0
}

// SetBkine BRK BKIN input enable
func (r *RegisterAf1Type) SetBkine(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf1FieldBkineMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf1FieldBkineMask)
	}
}

const (
	RegisterAf1FieldBkcmp1eShift = 1
	RegisterAf1FieldBkcmp1eMask  = 0x2
)

// GetBkcmp1e BRK COMP1 enable
func (r *RegisterAf1Type) GetBkcmp1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkcmp1eMask) != 0
}

// SetBkcmp1e BRK COMP1 enable
func (r *RegisterAf1Type) SetBkcmp1e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf1FieldBkcmp1eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf1FieldBkcmp1eMask)
	}
}

const (
	RegisterAf1FieldBkcmp2eShift = 2
	RegisterAf1FieldBkcmp2eMask  = 0x4
)

// GetBkcmp2e BRK COMP2 enable
func (r *RegisterAf1Type) GetBkcmp2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkcmp2eMask) != 0
}

// SetBkcmp2e BRK COMP2 enable
func (r *RegisterAf1Type) SetBkcmp2e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf1FieldBkcmp2eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf1FieldBkcmp2eMask)
	}
}

const (
	RegisterAf1FieldBkdf1bk0eShift = 8
	RegisterAf1FieldBkdf1bk0eMask  = 0x100
)

// GetBkdf1bk0e BRK dfsdm1_break[0] enable
func (r *RegisterAf1Type) GetBkdf1bk0e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkdf1bk0eMask) != 0
}

// SetBkdf1bk0e BRK dfsdm1_break[0] enable
func (r *RegisterAf1Type) SetBkdf1bk0e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf1FieldBkdf1bk0eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf1FieldBkdf1bk0eMask)
	}
}

const (
	RegisterAf1FieldBkinpShift = 9
	RegisterAf1FieldBkinpMask  = 0x200
)

// GetBkinp BRK BKIN input polarity
func (r *RegisterAf1Type) GetBkinp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkinpMask) != 0
}

// SetBkinp BRK BKIN input polarity
func (r *RegisterAf1Type) SetBkinp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf1FieldBkinpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf1FieldBkinpMask)
	}
}

const (
	RegisterAf1FieldBkcmp1pShift = 10
	RegisterAf1FieldBkcmp1pMask  = 0x400
)

// GetBkcmp1p BRK COMP1 input polarity
func (r *RegisterAf1Type) GetBkcmp1p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkcmp1pMask) != 0
}

// SetBkcmp1p BRK COMP1 input polarity
func (r *RegisterAf1Type) SetBkcmp1p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf1FieldBkcmp1pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf1FieldBkcmp1pMask)
	}
}

const (
	RegisterAf1FieldBkcmp2pShift = 11
	RegisterAf1FieldBkcmp2pMask  = 0x800
)

// GetBkcmp2p BRK COMP2 input polarity
func (r *RegisterAf1Type) GetBkcmp2p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkcmp2pMask) != 0
}

// SetBkcmp2p BRK COMP2 input polarity
func (r *RegisterAf1Type) SetBkcmp2p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf1FieldBkcmp2pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf1FieldBkcmp2pMask)
	}
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

// RegisterAf2Type Alternate function odfsdm1_breakster 2
type RegisterAf2Type uint32

func (r *RegisterAf2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAf2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAf2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAf2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAf2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAf2FieldBk2ineShift = 0
	RegisterAf2FieldBk2ineMask  = 0x1
)

// GetBk2ine BRK2 BKIN input enable
func (r *RegisterAf2Type) GetBk2ine() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2ineMask) != 0
}

// SetBk2ine BRK2 BKIN input enable
func (r *RegisterAf2Type) SetBk2ine(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf2FieldBk2ineMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf2FieldBk2ineMask)
	}
}

const (
	RegisterAf2FieldBk2cmp1eShift = 1
	RegisterAf2FieldBk2cmp1eMask  = 0x2
)

// GetBk2cmp1e BRK2 COMP1 enable
func (r *RegisterAf2Type) GetBk2cmp1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2cmp1eMask) != 0
}

// SetBk2cmp1e BRK2 COMP1 enable
func (r *RegisterAf2Type) SetBk2cmp1e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf2FieldBk2cmp1eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf2FieldBk2cmp1eMask)
	}
}

const (
	RegisterAf2FieldBk2cmp2eShift = 2
	RegisterAf2FieldBk2cmp2eMask  = 0x4
)

// GetBk2cmp2e BRK2 COMP2 enable
func (r *RegisterAf2Type) GetBk2cmp2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2cmp2eMask) != 0
}

// SetBk2cmp2e BRK2 COMP2 enable
func (r *RegisterAf2Type) SetBk2cmp2e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf2FieldBk2cmp2eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf2FieldBk2cmp2eMask)
	}
}

const (
	RegisterAf2FieldBk2df1bk1eShift = 8
	RegisterAf2FieldBk2df1bk1eMask  = 0x100
)

// GetBk2df1bk1e BRK2 dfsdm1_break[1] enable
func (r *RegisterAf2Type) GetBk2df1bk1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2df1bk1eMask) != 0
}

// SetBk2df1bk1e BRK2 dfsdm1_break[1] enable
func (r *RegisterAf2Type) SetBk2df1bk1e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf2FieldBk2df1bk1eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf2FieldBk2df1bk1eMask)
	}
}

const (
	RegisterAf2FieldBk2inpShift = 9
	RegisterAf2FieldBk2inpMask  = 0x200
)

// GetBk2inp BRK2 BKIN2 input polarity
func (r *RegisterAf2Type) GetBk2inp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2inpMask) != 0
}

// SetBk2inp BRK2 BKIN2 input polarity
func (r *RegisterAf2Type) SetBk2inp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf2FieldBk2inpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf2FieldBk2inpMask)
	}
}

const (
	RegisterAf2FieldBk2cmp1pShift = 10
	RegisterAf2FieldBk2cmp1pMask  = 0x400
)

// GetBk2cmp1p BRK2 COMP1 input polarit
func (r *RegisterAf2Type) GetBk2cmp1p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2cmp1pMask) != 0
}

// SetBk2cmp1p BRK2 COMP1 input polarit
func (r *RegisterAf2Type) SetBk2cmp1p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf2FieldBk2cmp1pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf2FieldBk2cmp1pMask)
	}
}

const (
	RegisterAf2FieldBk2cmp2pShift = 11
	RegisterAf2FieldBk2cmp2pMask  = 0x800
)

// GetBk2cmp2p BRK2 COMP2 input polarity
func (r *RegisterAf2Type) GetBk2cmp2p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2cmp2pMask) != 0
}

// SetBk2cmp2p BRK2 COMP2 input polarity
func (r *RegisterAf2Type) SetBk2cmp2p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf2FieldBk2cmp2pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf2FieldBk2cmp2pMask)
	}
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

const (
	RegisterTiselFieldTi3selShift = 16
	RegisterTiselFieldTi3selMask  = 0xf0000
)

// GetTi3sel selects TI3[0] to TI3[15] input
func (r *RegisterTiselType) GetTi3sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi3selMask) >> RegisterTiselFieldTi3selShift)
}

// SetTi3sel selects TI3[0] to TI3[15] input
func (r *RegisterTiselType) SetTi3sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi3selMask)|(uint32(value)<<RegisterTiselFieldTi3selShift))
}

const (
	RegisterTiselFieldTi4selShift = 24
	RegisterTiselFieldTi4selMask  = 0xf000000
)

// GetTi4sel selects TI4[0] to TI4[15] input
func (r *RegisterTiselType) GetTi4sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi4selMask) >> RegisterTiselFieldTi4selShift)
}

// SetTi4sel selects TI4[0] to TI4[15] input
func (r *RegisterTiselType) SetTi4sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi4selMask)|(uint32(value)<<RegisterTiselFieldTi4selShift))
}
