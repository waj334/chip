//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package tim1

import (
	"unsafe"
	"volatile"
)

var (
	Tim1 = (*_tim1)(unsafe.Pointer(uintptr(0x40010000)))
	Tim8 = (*_tim1)(unsafe.Pointer(uintptr(0x40010400)))
)

type _tim1 struct {
	Cr1          registerCr1Type
	Cr2          registerCr2Type
	Smcr         registerSmcrType
	Dier         registerDierType
	Sr           registerSrType
	Egr          registerEgrType
	Ccmr1_output registerCcmr1_outputType
	Ccmr1_input  registerCcmr1_inputType
	Ccmr2_output registerCcmr2_outputType
	Ccmr2_input  registerCcmr2_inputType
	Ccer         registerCcerType
	Cnt          registerCntType
	Psc          registerPscType
	Arr          registerArrType
	Rcr          registerRcrType
	Ccr1         registerCcr1Type
	Ccr2         registerCcr2Type
	Ccr3         registerCcr3Type
	Ccr4         registerCcr4Type
	Bdtr         registerBdtrType
	Dcr          registerDcrType
	Dmar         registerDmarType
	_            [4]byte
	Ccmr3_output registerCcmr3_outputType
	Ccr5         registerCcr5Type
	Ccr6         registerCcr6Type
	Af1          registerAf1Type
	Af2          registerAf2Type
	Tisel        registerTiselType
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
	RegisterCr2FieldCcpcShift = 0
	RegisterCr2FieldCcpcMask  = 0x1
)

// GetCcpc Capture/compare preloaded control
func (r *registerCr2Type) GetCcpc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCcpcMask) != 0
}

// SetCcpc Capture/compare preloaded control
func (r *registerCr2Type) SetCcpc(value bool) {
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
func (r *registerCr2Type) GetCcus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCcusMask) != 0
}

// SetCcus Capture/compare control update selection
func (r *registerCr2Type) SetCcus(value bool) {
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

const (
	RegisterCr2FieldOis1Shift = 8
	RegisterCr2FieldOis1Mask  = 0x100
)

// GetOis1 Output Idle state 1
func (r *registerCr2Type) GetOis1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis1Mask) != 0
}

// SetOis1 Output Idle state 1
func (r *registerCr2Type) SetOis1(value bool) {
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
func (r *registerCr2Type) GetOis1n() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis1nMask) != 0
}

// SetOis1n Output Idle state 1
func (r *registerCr2Type) SetOis1n(value bool) {
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
func (r *registerCr2Type) GetOis2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis2Mask) != 0
}

// SetOis2 Output Idle state 2
func (r *registerCr2Type) SetOis2(value bool) {
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
func (r *registerCr2Type) GetOis2n() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis2nMask) != 0
}

// SetOis2n Output Idle state 2
func (r *registerCr2Type) SetOis2n(value bool) {
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
func (r *registerCr2Type) GetOis3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis3Mask) != 0
}

// SetOis3 Output Idle state 3
func (r *registerCr2Type) SetOis3(value bool) {
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
func (r *registerCr2Type) GetOis3n() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis3nMask) != 0
}

// SetOis3n Output Idle state 3
func (r *registerCr2Type) SetOis3n(value bool) {
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
func (r *registerCr2Type) GetOis4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis4Mask) != 0
}

// SetOis4 Output Idle state 4
func (r *registerCr2Type) SetOis4(value bool) {
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
func (r *registerCr2Type) GetOis5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis5Mask) != 0
}

// SetOis5 Output Idle state 5
func (r *registerCr2Type) SetOis5(value bool) {
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
func (r *registerCr2Type) GetOis6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOis6Mask) != 0
}

// SetOis6 Output Idle state 6
func (r *registerCr2Type) SetOis6(value bool) {
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
func (r *registerCr2Type) GetMms2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldMms2Mask) >> RegisterCr2FieldMms2Shift)
}

// SetMms2 Master mode selection 2
func (r *registerCr2Type) SetMms2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldMms2Mask)|(uint32(value)<<RegisterCr2FieldMms2Shift))
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
	RegisterSmcrFieldSms_3Shift = 16
	RegisterSmcrFieldSms_3Mask  = 0x10000
)

// GetSms_3 Slave mode selection - bit 3
func (r *registerSmcrType) GetSms_3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldSms_3Mask) != 0
}

// SetSms_3 Slave mode selection - bit 3
func (r *registerSmcrType) SetSms_3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSmcrFieldSms_3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldSms_3Mask)
	}
}

const (
	RegisterSmcrFieldTs_4_3Shift = 20
	RegisterSmcrFieldTs_4_3Mask  = 0x300000
)

// GetTs_4_3 Trigger selection - bit 4:3
func (r *registerSmcrType) GetTs_4_3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldTs_4_3Mask) >> RegisterSmcrFieldTs_4_3Shift)
}

// SetTs_4_3 Trigger selection - bit 4:3
func (r *registerSmcrType) SetTs_4_3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldTs_4_3Mask)|(uint32(value)<<RegisterSmcrFieldTs_4_3Shift))
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
	RegisterDierFieldComieShift = 5
	RegisterDierFieldComieMask  = 0x20
)

// GetComie COM interrupt enable
func (r *registerDierType) GetComie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldComieMask) != 0
}

// SetComie COM interrupt enable
func (r *registerDierType) SetComie(value bool) {
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
	RegisterDierFieldBieShift = 7
	RegisterDierFieldBieMask  = 0x80
)

// GetBie Break interrupt enable
func (r *registerDierType) GetBie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldBieMask) != 0
}

// SetBie Break interrupt enable
func (r *registerDierType) SetBie(value bool) {
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
	RegisterDierFieldComdeShift = 13
	RegisterDierFieldComdeMask  = 0x2000
)

// GetComde COM DMA request enable
func (r *registerDierType) GetComde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldComdeMask) != 0
}

// SetComde COM DMA request enable
func (r *registerDierType) SetComde(value bool) {
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
	RegisterSrFieldComifShift = 5
	RegisterSrFieldComifMask  = 0x20
)

// GetComif COM interrupt flag
func (r *registerSrType) GetComif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldComifMask) != 0
}

// SetComif COM interrupt flag
func (r *registerSrType) SetComif(value bool) {
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
	RegisterSrFieldBifShift = 7
	RegisterSrFieldBifMask  = 0x80
)

// GetBif Break interrupt flag
func (r *registerSrType) GetBif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldBifMask) != 0
}

// SetBif Break interrupt flag
func (r *registerSrType) SetBif(value bool) {
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
func (r *registerSrType) GetB2if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldB2ifMask) != 0
}

// SetB2if Break 2 interrupt flag
func (r *registerSrType) SetB2if(value bool) {
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

const (
	RegisterSrFieldSbifShift = 13
	RegisterSrFieldSbifMask  = 0x2000
)

// GetSbif System Break interrupt flag
func (r *registerSrType) GetSbif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSbifMask) != 0
}

// SetSbif System Break interrupt flag
func (r *registerSrType) SetSbif(value bool) {
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
func (r *registerSrType) GetCc5if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc5ifMask) != 0
}

// SetCc5if Compare 5 interrupt flag
func (r *registerSrType) SetCc5if(value bool) {
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
func (r *registerSrType) GetCc6if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCc6ifMask) != 0
}

// SetCc6if Compare 6 interrupt flag
func (r *registerSrType) SetCc6if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCc6ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCc6ifMask)
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
	RegisterEgrFieldComgShift = 5
	RegisterEgrFieldComgMask  = 0x20
)

// GetComg Capture/Compare control update generation
func (r *registerEgrType) GetComg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldComgMask) != 0
}

// SetComg Capture/Compare control update generation
func (r *registerEgrType) SetComg(value bool) {
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

const (
	RegisterEgrFieldBgShift = 7
	RegisterEgrFieldBgMask  = 0x80
)

// GetBg Break generation
func (r *registerEgrType) GetBg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldBgMask) != 0
}

// SetBg Break generation
func (r *registerEgrType) SetBg(value bool) {
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
func (r *registerEgrType) GetB2g() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldB2gMask) != 0
}

// SetB2g Break 2 generation
func (r *registerEgrType) SetB2g(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldB2gMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldB2gMask)
	}
}

// registerCcmr1_outputType capture/compare mode register 1 (output mode)
type registerCcmr1_outputType uint32

const (
	RegisterCcmr1_outputFieldCc1sShift = 0
	RegisterCcmr1_outputFieldCc1sMask  = 0x3
)

// GetCc1s Capture/Compare 1 selection
func (r *registerCcmr1_outputType) GetCc1s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldCc1sMask) >> RegisterCcmr1_outputFieldCc1sShift)
}

// SetCc1s Capture/Compare 1 selection
func (r *registerCcmr1_outputType) SetCc1s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_outputFieldCc1sMask)|(uint32(value)<<RegisterCcmr1_outputFieldCc1sShift))
}

const (
	RegisterCcmr1_outputFieldOc1feShift = 2
	RegisterCcmr1_outputFieldOc1feMask  = 0x4
)

// GetOc1fe Output Compare 1 fast enable
func (r *registerCcmr1_outputType) GetOc1fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldOc1feMask) != 0
}

// SetOc1fe Output Compare 1 fast enable
func (r *registerCcmr1_outputType) SetOc1fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1_outputFieldOc1feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_outputFieldOc1feMask)
	}
}

const (
	RegisterCcmr1_outputFieldOc1peShift = 3
	RegisterCcmr1_outputFieldOc1peMask  = 0x8
)

// GetOc1pe Output Compare 1 preload enable
func (r *registerCcmr1_outputType) GetOc1pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldOc1peMask) != 0
}

// SetOc1pe Output Compare 1 preload enable
func (r *registerCcmr1_outputType) SetOc1pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1_outputFieldOc1peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_outputFieldOc1peMask)
	}
}

const (
	RegisterCcmr1_outputFieldOc1mShift = 4
	RegisterCcmr1_outputFieldOc1mMask  = 0x70
)

// GetOc1m Output Compare 1 mode
func (r *registerCcmr1_outputType) GetOc1m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldOc1mMask) >> RegisterCcmr1_outputFieldOc1mShift)
}

// SetOc1m Output Compare 1 mode
func (r *registerCcmr1_outputType) SetOc1m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_outputFieldOc1mMask)|(uint32(value)<<RegisterCcmr1_outputFieldOc1mShift))
}

const (
	RegisterCcmr1_outputFieldOc1ceShift = 7
	RegisterCcmr1_outputFieldOc1ceMask  = 0x80
)

// GetOc1ce Output Compare 1 clear enable
func (r *registerCcmr1_outputType) GetOc1ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldOc1ceMask) != 0
}

// SetOc1ce Output Compare 1 clear enable
func (r *registerCcmr1_outputType) SetOc1ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1_outputFieldOc1ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_outputFieldOc1ceMask)
	}
}

const (
	RegisterCcmr1_outputFieldCc2sShift = 8
	RegisterCcmr1_outputFieldCc2sMask  = 0x300
)

// GetCc2s Capture/Compare 2 selection
func (r *registerCcmr1_outputType) GetCc2s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldCc2sMask) >> RegisterCcmr1_outputFieldCc2sShift)
}

// SetCc2s Capture/Compare 2 selection
func (r *registerCcmr1_outputType) SetCc2s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_outputFieldCc2sMask)|(uint32(value)<<RegisterCcmr1_outputFieldCc2sShift))
}

const (
	RegisterCcmr1_outputFieldOc2feShift = 10
	RegisterCcmr1_outputFieldOc2feMask  = 0x400
)

// GetOc2fe Output Compare 2 fast enable
func (r *registerCcmr1_outputType) GetOc2fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldOc2feMask) != 0
}

// SetOc2fe Output Compare 2 fast enable
func (r *registerCcmr1_outputType) SetOc2fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1_outputFieldOc2feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_outputFieldOc2feMask)
	}
}

const (
	RegisterCcmr1_outputFieldOc2peShift = 11
	RegisterCcmr1_outputFieldOc2peMask  = 0x800
)

// GetOc2pe Output Compare 2 preload enable
func (r *registerCcmr1_outputType) GetOc2pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldOc2peMask) != 0
}

// SetOc2pe Output Compare 2 preload enable
func (r *registerCcmr1_outputType) SetOc2pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1_outputFieldOc2peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_outputFieldOc2peMask)
	}
}

const (
	RegisterCcmr1_outputFieldOc2mShift = 12
	RegisterCcmr1_outputFieldOc2mMask  = 0x7000
)

// GetOc2m Output Compare 2 mode
func (r *registerCcmr1_outputType) GetOc2m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldOc2mMask) >> RegisterCcmr1_outputFieldOc2mShift)
}

// SetOc2m Output Compare 2 mode
func (r *registerCcmr1_outputType) SetOc2m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_outputFieldOc2mMask)|(uint32(value)<<RegisterCcmr1_outputFieldOc2mShift))
}

const (
	RegisterCcmr1_outputFieldOc2ceShift = 15
	RegisterCcmr1_outputFieldOc2ceMask  = 0x8000
)

// GetOc2ce Output Compare 2 clear enable
func (r *registerCcmr1_outputType) GetOc2ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldOc2ceMask) != 0
}

// SetOc2ce Output Compare 2 clear enable
func (r *registerCcmr1_outputType) SetOc2ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1_outputFieldOc2ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_outputFieldOc2ceMask)
	}
}

const (
	RegisterCcmr1_outputFieldOc1m_3Shift = 16
	RegisterCcmr1_outputFieldOc1m_3Mask  = 0x10000
)

// GetOc1m_3 Output Compare 1 mode - bit 3
func (r *registerCcmr1_outputType) GetOc1m_3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldOc1m_3Mask) != 0
}

// SetOc1m_3 Output Compare 1 mode - bit 3
func (r *registerCcmr1_outputType) SetOc1m_3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1_outputFieldOc1m_3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_outputFieldOc1m_3Mask)
	}
}

const (
	RegisterCcmr1_outputFieldOc2m_3Shift = 24
	RegisterCcmr1_outputFieldOc2m_3Mask  = 0x1000000
)

// GetOc2m_3 Output Compare 2 mode - bit 3
func (r *registerCcmr1_outputType) GetOc2m_3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldOc2m_3Mask) != 0
}

// SetOc2m_3 Output Compare 2 mode - bit 3
func (r *registerCcmr1_outputType) SetOc2m_3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr1_outputFieldOc2m_3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_outputFieldOc2m_3Mask)
	}
}

// registerCcmr1_inputType capture/compare mode register 1 (input mode)
type registerCcmr1_inputType uint32

const (
	RegisterCcmr1_inputFieldCc1sShift = 0
	RegisterCcmr1_inputFieldCc1sMask  = 0x3
)

// GetCc1s Capture/Compare 1 selection
func (r *registerCcmr1_inputType) GetCc1s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_inputFieldCc1sMask) >> RegisterCcmr1_inputFieldCc1sShift)
}

// SetCc1s Capture/Compare 1 selection
func (r *registerCcmr1_inputType) SetCc1s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_inputFieldCc1sMask)|(uint32(value)<<RegisterCcmr1_inputFieldCc1sShift))
}

const (
	RegisterCcmr1_inputFieldIcpcsShift = 2
	RegisterCcmr1_inputFieldIcpcsMask  = 0xc
)

// GetIcpcs Input capture 1 prescaler
func (r *registerCcmr1_inputType) GetIcpcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_inputFieldIcpcsMask) >> RegisterCcmr1_inputFieldIcpcsShift)
}

// SetIcpcs Input capture 1 prescaler
func (r *registerCcmr1_inputType) SetIcpcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_inputFieldIcpcsMask)|(uint32(value)<<RegisterCcmr1_inputFieldIcpcsShift))
}

const (
	RegisterCcmr1_inputFieldIc1fShift = 4
	RegisterCcmr1_inputFieldIc1fMask  = 0xf0
)

// GetIc1f Input capture 1 filter
func (r *registerCcmr1_inputType) GetIc1f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_inputFieldIc1fMask) >> RegisterCcmr1_inputFieldIc1fShift)
}

// SetIc1f Input capture 1 filter
func (r *registerCcmr1_inputType) SetIc1f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_inputFieldIc1fMask)|(uint32(value)<<RegisterCcmr1_inputFieldIc1fShift))
}

const (
	RegisterCcmr1_inputFieldCc2sShift = 8
	RegisterCcmr1_inputFieldCc2sMask  = 0x300
)

// GetCc2s Capture/Compare 2 selection
func (r *registerCcmr1_inputType) GetCc2s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_inputFieldCc2sMask) >> RegisterCcmr1_inputFieldCc2sShift)
}

// SetCc2s Capture/Compare 2 selection
func (r *registerCcmr1_inputType) SetCc2s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_inputFieldCc2sMask)|(uint32(value)<<RegisterCcmr1_inputFieldCc2sShift))
}

const (
	RegisterCcmr1_inputFieldIc2pcsShift = 10
	RegisterCcmr1_inputFieldIc2pcsMask  = 0xc00
)

// GetIc2pcs Input capture 2 prescaler
func (r *registerCcmr1_inputType) GetIc2pcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_inputFieldIc2pcsMask) >> RegisterCcmr1_inputFieldIc2pcsShift)
}

// SetIc2pcs Input capture 2 prescaler
func (r *registerCcmr1_inputType) SetIc2pcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_inputFieldIc2pcsMask)|(uint32(value)<<RegisterCcmr1_inputFieldIc2pcsShift))
}

const (
	RegisterCcmr1_inputFieldIc2fShift = 12
	RegisterCcmr1_inputFieldIc2fMask  = 0xf000
)

// GetIc2f Input capture 2 filter
func (r *registerCcmr1_inputType) GetIc2f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_inputFieldIc2fMask) >> RegisterCcmr1_inputFieldIc2fShift)
}

// SetIc2f Input capture 2 filter
func (r *registerCcmr1_inputType) SetIc2f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_inputFieldIc2fMask)|(uint32(value)<<RegisterCcmr1_inputFieldIc2fShift))
}

// registerCcmr2_outputType capture/compare mode register 2 (output mode)
type registerCcmr2_outputType uint32

const (
	RegisterCcmr2_outputFieldCc3sShift = 0
	RegisterCcmr2_outputFieldCc3sMask  = 0x3
)

// GetCc3s Capture/Compare 3 selection
func (r *registerCcmr2_outputType) GetCc3s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_outputFieldCc3sMask) >> RegisterCcmr2_outputFieldCc3sShift)
}

// SetCc3s Capture/Compare 3 selection
func (r *registerCcmr2_outputType) SetCc3s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_outputFieldCc3sMask)|(uint32(value)<<RegisterCcmr2_outputFieldCc3sShift))
}

const (
	RegisterCcmr2_outputFieldOc3feShift = 2
	RegisterCcmr2_outputFieldOc3feMask  = 0x4
)

// GetOc3fe Output compare 3 fast enable
func (r *registerCcmr2_outputType) GetOc3fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_outputFieldOc3feMask) != 0
}

// SetOc3fe Output compare 3 fast enable
func (r *registerCcmr2_outputType) SetOc3fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2_outputFieldOc3feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_outputFieldOc3feMask)
	}
}

const (
	RegisterCcmr2_outputFieldOc3peShift = 3
	RegisterCcmr2_outputFieldOc3peMask  = 0x8
)

// GetOc3pe Output compare 3 preload enable
func (r *registerCcmr2_outputType) GetOc3pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_outputFieldOc3peMask) != 0
}

// SetOc3pe Output compare 3 preload enable
func (r *registerCcmr2_outputType) SetOc3pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2_outputFieldOc3peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_outputFieldOc3peMask)
	}
}

const (
	RegisterCcmr2_outputFieldOc3mShift = 4
	RegisterCcmr2_outputFieldOc3mMask  = 0x70
)

// GetOc3m Output compare 3 mode
func (r *registerCcmr2_outputType) GetOc3m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_outputFieldOc3mMask) >> RegisterCcmr2_outputFieldOc3mShift)
}

// SetOc3m Output compare 3 mode
func (r *registerCcmr2_outputType) SetOc3m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_outputFieldOc3mMask)|(uint32(value)<<RegisterCcmr2_outputFieldOc3mShift))
}

const (
	RegisterCcmr2_outputFieldOc3ceShift = 7
	RegisterCcmr2_outputFieldOc3ceMask  = 0x80
)

// GetOc3ce Output compare 3 clear enable
func (r *registerCcmr2_outputType) GetOc3ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_outputFieldOc3ceMask) != 0
}

// SetOc3ce Output compare 3 clear enable
func (r *registerCcmr2_outputType) SetOc3ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2_outputFieldOc3ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_outputFieldOc3ceMask)
	}
}

const (
	RegisterCcmr2_outputFieldCc4sShift = 8
	RegisterCcmr2_outputFieldCc4sMask  = 0x300
)

// GetCc4s Capture/Compare 4 selection
func (r *registerCcmr2_outputType) GetCc4s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_outputFieldCc4sMask) >> RegisterCcmr2_outputFieldCc4sShift)
}

// SetCc4s Capture/Compare 4 selection
func (r *registerCcmr2_outputType) SetCc4s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_outputFieldCc4sMask)|(uint32(value)<<RegisterCcmr2_outputFieldCc4sShift))
}

const (
	RegisterCcmr2_outputFieldOc4feShift = 10
	RegisterCcmr2_outputFieldOc4feMask  = 0x400
)

// GetOc4fe Output compare 4 fast enable
func (r *registerCcmr2_outputType) GetOc4fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_outputFieldOc4feMask) != 0
}

// SetOc4fe Output compare 4 fast enable
func (r *registerCcmr2_outputType) SetOc4fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2_outputFieldOc4feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_outputFieldOc4feMask)
	}
}

const (
	RegisterCcmr2_outputFieldOc4peShift = 11
	RegisterCcmr2_outputFieldOc4peMask  = 0x800
)

// GetOc4pe Output compare 4 preload enable
func (r *registerCcmr2_outputType) GetOc4pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_outputFieldOc4peMask) != 0
}

// SetOc4pe Output compare 4 preload enable
func (r *registerCcmr2_outputType) SetOc4pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2_outputFieldOc4peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_outputFieldOc4peMask)
	}
}

const (
	RegisterCcmr2_outputFieldOc4mShift = 12
	RegisterCcmr2_outputFieldOc4mMask  = 0x7000
)

// GetOc4m Output compare 4 mode
func (r *registerCcmr2_outputType) GetOc4m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_outputFieldOc4mMask) >> RegisterCcmr2_outputFieldOc4mShift)
}

// SetOc4m Output compare 4 mode
func (r *registerCcmr2_outputType) SetOc4m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_outputFieldOc4mMask)|(uint32(value)<<RegisterCcmr2_outputFieldOc4mShift))
}

const (
	RegisterCcmr2_outputFieldOc4ceShift = 15
	RegisterCcmr2_outputFieldOc4ceMask  = 0x8000
)

// GetOc4ce Output compare 4 clear enable
func (r *registerCcmr2_outputType) GetOc4ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_outputFieldOc4ceMask) != 0
}

// SetOc4ce Output compare 4 clear enable
func (r *registerCcmr2_outputType) SetOc4ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2_outputFieldOc4ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_outputFieldOc4ceMask)
	}
}

const (
	RegisterCcmr2_outputFieldOc3m_3Shift = 16
	RegisterCcmr2_outputFieldOc3m_3Mask  = 0x10000
)

// GetOc3m_3 Output Compare 3 mode - bit 3
func (r *registerCcmr2_outputType) GetOc3m_3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_outputFieldOc3m_3Mask) != 0
}

// SetOc3m_3 Output Compare 3 mode - bit 3
func (r *registerCcmr2_outputType) SetOc3m_3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2_outputFieldOc3m_3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_outputFieldOc3m_3Mask)
	}
}

const (
	RegisterCcmr2_outputFieldOc4m_4Shift = 24
	RegisterCcmr2_outputFieldOc4m_4Mask  = 0x1000000
)

// GetOc4m_4 Output Compare 4 mode - bit 3
func (r *registerCcmr2_outputType) GetOc4m_4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_outputFieldOc4m_4Mask) != 0
}

// SetOc4m_4 Output Compare 4 mode - bit 3
func (r *registerCcmr2_outputType) SetOc4m_4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr2_outputFieldOc4m_4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_outputFieldOc4m_4Mask)
	}
}

// registerCcmr2_inputType capture/compare mode register 2 (input mode)
type registerCcmr2_inputType uint32

const (
	RegisterCcmr2_inputFieldCc3sShift = 0
	RegisterCcmr2_inputFieldCc3sMask  = 0x3
)

// GetCc3s Capture/compare 3 selection
func (r *registerCcmr2_inputType) GetCc3s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_inputFieldCc3sMask) >> RegisterCcmr2_inputFieldCc3sShift)
}

// SetCc3s Capture/compare 3 selection
func (r *registerCcmr2_inputType) SetCc3s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_inputFieldCc3sMask)|(uint32(value)<<RegisterCcmr2_inputFieldCc3sShift))
}

const (
	RegisterCcmr2_inputFieldIc3pscShift = 2
	RegisterCcmr2_inputFieldIc3pscMask  = 0xc
)

// GetIc3psc Input capture 3 prescaler
func (r *registerCcmr2_inputType) GetIc3psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_inputFieldIc3pscMask) >> RegisterCcmr2_inputFieldIc3pscShift)
}

// SetIc3psc Input capture 3 prescaler
func (r *registerCcmr2_inputType) SetIc3psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_inputFieldIc3pscMask)|(uint32(value)<<RegisterCcmr2_inputFieldIc3pscShift))
}

const (
	RegisterCcmr2_inputFieldIc3fShift = 4
	RegisterCcmr2_inputFieldIc3fMask  = 0xf0
)

// GetIc3f Input capture 3 filter
func (r *registerCcmr2_inputType) GetIc3f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_inputFieldIc3fMask) >> RegisterCcmr2_inputFieldIc3fShift)
}

// SetIc3f Input capture 3 filter
func (r *registerCcmr2_inputType) SetIc3f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_inputFieldIc3fMask)|(uint32(value)<<RegisterCcmr2_inputFieldIc3fShift))
}

const (
	RegisterCcmr2_inputFieldCc4sShift = 8
	RegisterCcmr2_inputFieldCc4sMask  = 0x300
)

// GetCc4s Capture/Compare 4 selection
func (r *registerCcmr2_inputType) GetCc4s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_inputFieldCc4sMask) >> RegisterCcmr2_inputFieldCc4sShift)
}

// SetCc4s Capture/Compare 4 selection
func (r *registerCcmr2_inputType) SetCc4s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_inputFieldCc4sMask)|(uint32(value)<<RegisterCcmr2_inputFieldCc4sShift))
}

const (
	RegisterCcmr2_inputFieldIc4pscShift = 10
	RegisterCcmr2_inputFieldIc4pscMask  = 0xc00
)

// GetIc4psc Input capture 4 prescaler
func (r *registerCcmr2_inputType) GetIc4psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_inputFieldIc4pscMask) >> RegisterCcmr2_inputFieldIc4pscShift)
}

// SetIc4psc Input capture 4 prescaler
func (r *registerCcmr2_inputType) SetIc4psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_inputFieldIc4pscMask)|(uint32(value)<<RegisterCcmr2_inputFieldIc4pscShift))
}

const (
	RegisterCcmr2_inputFieldIc4fShift = 12
	RegisterCcmr2_inputFieldIc4fMask  = 0xf000
)

// GetIc4f Input capture 4 filter
func (r *registerCcmr2_inputType) GetIc4f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr2_inputFieldIc4fMask) >> RegisterCcmr2_inputFieldIc4fShift)
}

// SetIc4f Input capture 4 filter
func (r *registerCcmr2_inputType) SetIc4f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr2_inputFieldIc4fMask)|(uint32(value)<<RegisterCcmr2_inputFieldIc4fShift))
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
	RegisterCcerFieldCc1neShift = 2
	RegisterCcerFieldCc1neMask  = 0x4
)

// GetCc1ne Capture/Compare 1 complementary output enable
func (r *registerCcerType) GetCc1ne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc1neMask) != 0
}

// SetCc1ne Capture/Compare 1 complementary output enable
func (r *registerCcerType) SetCc1ne(value bool) {
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
	RegisterCcerFieldCc2neShift = 6
	RegisterCcerFieldCc2neMask  = 0x40
)

// GetCc2ne Capture/Compare 2 complementary output enable
func (r *registerCcerType) GetCc2ne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc2neMask) != 0
}

// SetCc2ne Capture/Compare 2 complementary output enable
func (r *registerCcerType) SetCc2ne(value bool) {
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
	RegisterCcerFieldCc3neShift = 10
	RegisterCcerFieldCc3neMask  = 0x400
)

// GetCc3ne Capture/Compare 3 complementary output enable
func (r *registerCcerType) GetCc3ne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc3neMask) != 0
}

// SetCc3ne Capture/Compare 3 complementary output enable
func (r *registerCcerType) SetCc3ne(value bool) {
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

// GetCc4np Capture/Compare 4 complementary output polarity
func (r *registerCcerType) GetCc4np() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc4npMask) != 0
}

// SetCc4np Capture/Compare 4 complementary output polarity
func (r *registerCcerType) SetCc4np(value bool) {
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
func (r *registerCcerType) GetCc5e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc5eMask) != 0
}

// SetCc5e Capture/Compare 5 output enable
func (r *registerCcerType) SetCc5e(value bool) {
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
func (r *registerCcerType) GetCc5p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc5pMask) != 0
}

// SetCc5p Capture/Compare 5 output polarity
func (r *registerCcerType) SetCc5p(value bool) {
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
func (r *registerCcerType) GetCc6e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc6eMask) != 0
}

// SetCc6e Capture/Compare 6 output enable
func (r *registerCcerType) SetCc6e(value bool) {
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
func (r *registerCcerType) GetCc6p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCc6pMask) != 0
}

// SetCc6p Capture/Compare 6 output polarity
func (r *registerCcerType) SetCc6p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCc6pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCc6pMask)
	}
}

// registerCntType counter
type registerCntType uint32

const (
	RegisterCntFieldCntShift = 0
	RegisterCntFieldCntMask  = 0xffff
)

// GetCnt counter value
func (r *registerCntType) GetCnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldCntMask) >> RegisterCntFieldCntShift)
}

// SetCnt counter value
func (r *registerCntType) SetCnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldCntMask)|(uint32(value)<<RegisterCntFieldCntShift))
}

const (
	RegisterCntFieldUifcpyShift = 31
	RegisterCntFieldUifcpyMask  = 0x80000000
)

// GetUifcpy UIF copy
func (r *registerCntType) GetUifcpy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldUifcpyMask) != 0
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

// GetArr Auto-reload value
func (r *registerArrType) GetArr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterArrFieldArrMask) >> RegisterArrFieldArrShift)
}

// SetArr Auto-reload value
func (r *registerArrType) SetArr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArrFieldArrMask)|(uint32(value)<<RegisterArrFieldArrShift))
}

// registerRcrType repetition counter register
type registerRcrType uint32

const (
	RegisterRcrFieldRepShift = 0
	RegisterRcrFieldRepMask  = 0xff
)

// GetRep Repetition counter value
func (r *registerRcrType) GetRep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRcrFieldRepMask) >> RegisterRcrFieldRepShift)
}

// SetRep Repetition counter value
func (r *registerRcrType) SetRep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRcrFieldRepMask)|(uint32(value)<<RegisterRcrFieldRepShift))
}

// registerCcr1Type capture/compare register 1
type registerCcr1Type uint32

const (
	RegisterCcr1FieldCcr1Shift = 0
	RegisterCcr1FieldCcr1Mask  = 0xffff
)

// GetCcr1 Capture/Compare 1 value
func (r *registerCcr1Type) GetCcr1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldCcr1Mask) >> RegisterCcr1FieldCcr1Shift)
}

// SetCcr1 Capture/Compare 1 value
func (r *registerCcr1Type) SetCcr1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldCcr1Mask)|(uint32(value)<<RegisterCcr1FieldCcr1Shift))
}

// registerCcr2Type capture/compare register 2
type registerCcr2Type uint32

const (
	RegisterCcr2FieldCcr2Shift = 0
	RegisterCcr2FieldCcr2Mask  = 0xffff
)

// GetCcr2 Capture/Compare 2 value
func (r *registerCcr2Type) GetCcr2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldCcr2Mask) >> RegisterCcr2FieldCcr2Shift)
}

// SetCcr2 Capture/Compare 2 value
func (r *registerCcr2Type) SetCcr2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldCcr2Mask)|(uint32(value)<<RegisterCcr2FieldCcr2Shift))
}

// registerCcr3Type capture/compare register 3
type registerCcr3Type uint32

const (
	RegisterCcr3FieldCcr3Shift = 0
	RegisterCcr3FieldCcr3Mask  = 0xffff
)

// GetCcr3 Capture/Compare value
func (r *registerCcr3Type) GetCcr3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldCcr3Mask) >> RegisterCcr3FieldCcr3Shift)
}

// SetCcr3 Capture/Compare value
func (r *registerCcr3Type) SetCcr3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldCcr3Mask)|(uint32(value)<<RegisterCcr3FieldCcr3Shift))
}

// registerCcr4Type capture/compare register 4
type registerCcr4Type uint32

const (
	RegisterCcr4FieldCcr4Shift = 0
	RegisterCcr4FieldCcr4Mask  = 0xffff
)

// GetCcr4 Capture/Compare value
func (r *registerCcr4Type) GetCcr4() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldCcr4Mask) >> RegisterCcr4FieldCcr4Shift)
}

// SetCcr4 Capture/Compare value
func (r *registerCcr4Type) SetCcr4(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldCcr4Mask)|(uint32(value)<<RegisterCcr4FieldCcr4Shift))
}

// registerBdtrType break and dead-time register
type registerBdtrType uint32

const (
	RegisterBdtrFieldDtgShift = 0
	RegisterBdtrFieldDtgMask  = 0xff
)

// GetDtg Dead-time generator setup
func (r *registerBdtrType) GetDtg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldDtgMask) >> RegisterBdtrFieldDtgShift)
}

// SetDtg Dead-time generator setup
func (r *registerBdtrType) SetDtg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldDtgMask)|(uint32(value)<<RegisterBdtrFieldDtgShift))
}

const (
	RegisterBdtrFieldLockShift = 8
	RegisterBdtrFieldLockMask  = 0x300
)

// GetLock Lock configuration
func (r *registerBdtrType) GetLock() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldLockMask) >> RegisterBdtrFieldLockShift)
}

// SetLock Lock configuration
func (r *registerBdtrType) SetLock(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldLockMask)|(uint32(value)<<RegisterBdtrFieldLockShift))
}

const (
	RegisterBdtrFieldOssiShift = 10
	RegisterBdtrFieldOssiMask  = 0x400
)

// GetOssi Off-state selection for Idle mode
func (r *registerBdtrType) GetOssi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldOssiMask) != 0
}

// SetOssi Off-state selection for Idle mode
func (r *registerBdtrType) SetOssi(value bool) {
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
func (r *registerBdtrType) GetOssr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldOssrMask) != 0
}

// SetOssr Off-state selection for Run mode
func (r *registerBdtrType) SetOssr(value bool) {
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
func (r *registerBdtrType) GetBke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldBkeMask) != 0
}

// SetBke Break enable
func (r *registerBdtrType) SetBke(value bool) {
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
func (r *registerBdtrType) GetBkp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldBkpMask) != 0
}

// SetBkp Break polarity
func (r *registerBdtrType) SetBkp(value bool) {
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
func (r *registerBdtrType) GetAoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldAoeMask) != 0
}

// SetAoe Automatic output enable
func (r *registerBdtrType) SetAoe(value bool) {
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
func (r *registerBdtrType) GetMoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldMoeMask) != 0
}

// SetMoe Main output enable
func (r *registerBdtrType) SetMoe(value bool) {
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
func (r *registerBdtrType) GetBkf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldBkfMask) >> RegisterBdtrFieldBkfShift)
}

// SetBkf Break filter
func (r *registerBdtrType) SetBkf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldBkfMask)|(uint32(value)<<RegisterBdtrFieldBkfShift))
}

const (
	RegisterBdtrFieldBk2fShift = 20
	RegisterBdtrFieldBk2fMask  = 0xf00000
)

// GetBk2f Break 2 filter
func (r *registerBdtrType) GetBk2f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldBk2fMask) >> RegisterBdtrFieldBk2fShift)
}

// SetBk2f Break 2 filter
func (r *registerBdtrType) SetBk2f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldBk2fMask)|(uint32(value)<<RegisterBdtrFieldBk2fShift))
}

const (
	RegisterBdtrFieldBk2eShift = 24
	RegisterBdtrFieldBk2eMask  = 0x1000000
)

// GetBk2e Break 2 enable
func (r *registerBdtrType) GetBk2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldBk2eMask) != 0
}

// SetBk2e Break 2 enable
func (r *registerBdtrType) SetBk2e(value bool) {
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
func (r *registerBdtrType) GetBk2p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtrFieldBk2pMask) != 0
}

// SetBk2p Break 2 polarity
func (r *registerBdtrType) SetBk2p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtrFieldBk2pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtrFieldBk2pMask)
	}
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

// registerCcmr3_outputType capture/compare mode register 3 (output mode)
type registerCcmr3_outputType uint32

const (
	RegisterCcmr3_outputFieldOc5feShift = 2
	RegisterCcmr3_outputFieldOc5feMask  = 0x4
)

// GetOc5fe Output compare 5 fast enable
func (r *registerCcmr3_outputType) GetOc5fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3_outputFieldOc5feMask) != 0
}

// SetOc5fe Output compare 5 fast enable
func (r *registerCcmr3_outputType) SetOc5fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3_outputFieldOc5feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3_outputFieldOc5feMask)
	}
}

const (
	RegisterCcmr3_outputFieldOc5peShift = 3
	RegisterCcmr3_outputFieldOc5peMask  = 0x8
)

// GetOc5pe Output compare 5 preload enable
func (r *registerCcmr3_outputType) GetOc5pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3_outputFieldOc5peMask) != 0
}

// SetOc5pe Output compare 5 preload enable
func (r *registerCcmr3_outputType) SetOc5pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3_outputFieldOc5peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3_outputFieldOc5peMask)
	}
}

const (
	RegisterCcmr3_outputFieldOc5mShift = 4
	RegisterCcmr3_outputFieldOc5mMask  = 0x70
)

// GetOc5m Output compare 5 mode
func (r *registerCcmr3_outputType) GetOc5m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3_outputFieldOc5mMask) >> RegisterCcmr3_outputFieldOc5mShift)
}

// SetOc5m Output compare 5 mode
func (r *registerCcmr3_outputType) SetOc5m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3_outputFieldOc5mMask)|(uint32(value)<<RegisterCcmr3_outputFieldOc5mShift))
}

const (
	RegisterCcmr3_outputFieldOc5ceShift = 7
	RegisterCcmr3_outputFieldOc5ceMask  = 0x80
)

// GetOc5ce Output compare 5 clear enable
func (r *registerCcmr3_outputType) GetOc5ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3_outputFieldOc5ceMask) != 0
}

// SetOc5ce Output compare 5 clear enable
func (r *registerCcmr3_outputType) SetOc5ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3_outputFieldOc5ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3_outputFieldOc5ceMask)
	}
}

const (
	RegisterCcmr3_outputFieldOc6feShift = 10
	RegisterCcmr3_outputFieldOc6feMask  = 0x400
)

// GetOc6fe Output compare 6 fast enable
func (r *registerCcmr3_outputType) GetOc6fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3_outputFieldOc6feMask) != 0
}

// SetOc6fe Output compare 6 fast enable
func (r *registerCcmr3_outputType) SetOc6fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3_outputFieldOc6feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3_outputFieldOc6feMask)
	}
}

const (
	RegisterCcmr3_outputFieldOc6peShift = 11
	RegisterCcmr3_outputFieldOc6peMask  = 0x800
)

// GetOc6pe Output compare 6 preload enable
func (r *registerCcmr3_outputType) GetOc6pe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3_outputFieldOc6peMask) != 0
}

// SetOc6pe Output compare 6 preload enable
func (r *registerCcmr3_outputType) SetOc6pe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3_outputFieldOc6peMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3_outputFieldOc6peMask)
	}
}

const (
	RegisterCcmr3_outputFieldOc6mShift = 12
	RegisterCcmr3_outputFieldOc6mMask  = 0x7000
)

// GetOc6m Output compare 6 mode
func (r *registerCcmr3_outputType) GetOc6m() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3_outputFieldOc6mMask) >> RegisterCcmr3_outputFieldOc6mShift)
}

// SetOc6m Output compare 6 mode
func (r *registerCcmr3_outputType) SetOc6m(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3_outputFieldOc6mMask)|(uint32(value)<<RegisterCcmr3_outputFieldOc6mShift))
}

const (
	RegisterCcmr3_outputFieldOc6ceShift = 15
	RegisterCcmr3_outputFieldOc6ceMask  = 0x8000
)

// GetOc6ce Output compare 6 clear enable
func (r *registerCcmr3_outputType) GetOc6ce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3_outputFieldOc6ceMask) != 0
}

// SetOc6ce Output compare 6 clear enable
func (r *registerCcmr3_outputType) SetOc6ce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3_outputFieldOc6ceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3_outputFieldOc6ceMask)
	}
}

const (
	RegisterCcmr3_outputFieldOc5m3Shift = 16
	RegisterCcmr3_outputFieldOc5m3Mask  = 0x10000
)

// GetOc5m3 Output Compare 5 mode
func (r *registerCcmr3_outputType) GetOc5m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3_outputFieldOc5m3Mask) != 0
}

// SetOc5m3 Output Compare 5 mode
func (r *registerCcmr3_outputType) SetOc5m3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3_outputFieldOc5m3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3_outputFieldOc5m3Mask)
	}
}

const (
	RegisterCcmr3_outputFieldOc6m3Shift = 24
	RegisterCcmr3_outputFieldOc6m3Mask  = 0x1000000
)

// GetOc6m3 Output Compare 6 mode
func (r *registerCcmr3_outputType) GetOc6m3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr3_outputFieldOc6m3Mask) != 0
}

// SetOc6m3 Output Compare 6 mode
func (r *registerCcmr3_outputType) SetOc6m3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmr3_outputFieldOc6m3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmr3_outputFieldOc6m3Mask)
	}
}

// registerCcr5Type capture/compare register 5
type registerCcr5Type uint32

const (
	RegisterCcr5FieldCcr5Shift = 0
	RegisterCcr5FieldCcr5Mask  = 0xffff
)

// GetCcr5 Capture/Compare 5 value
func (r *registerCcr5Type) GetCcr5() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldCcr5Mask) >> RegisterCcr5FieldCcr5Shift)
}

// SetCcr5 Capture/Compare 5 value
func (r *registerCcr5Type) SetCcr5(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldCcr5Mask)|(uint32(value)<<RegisterCcr5FieldCcr5Shift))
}

const (
	RegisterCcr5FieldGc5c1Shift = 29
	RegisterCcr5FieldGc5c1Mask  = 0x20000000
)

// GetGc5c1 Group Channel 5 and Channel 1
func (r *registerCcr5Type) GetGc5c1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldGc5c1Mask) != 0
}

// SetGc5c1 Group Channel 5 and Channel 1
func (r *registerCcr5Type) SetGc5c1(value bool) {
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
func (r *registerCcr5Type) GetGc5c2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldGc5c2Mask) != 0
}

// SetGc5c2 Group Channel 5 and Channel 2
func (r *registerCcr5Type) SetGc5c2(value bool) {
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
func (r *registerCcr5Type) GetGc5c3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldGc5c3Mask) != 0
}

// SetGc5c3 Group Channel 5 and Channel 3
func (r *registerCcr5Type) SetGc5c3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldGc5c3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldGc5c3Mask)
	}
}

// registerCcr6Type capture/compare register 6
type registerCcr6Type uint32

const (
	RegisterCcr6FieldCcr6Shift = 0
	RegisterCcr6FieldCcr6Mask  = 0xffff
)

// GetCcr6 Capture/Compare 6 value
func (r *registerCcr6Type) GetCcr6() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldCcr6Mask) >> RegisterCcr6FieldCcr6Shift)
}

// SetCcr6 Capture/Compare 6 value
func (r *registerCcr6Type) SetCcr6(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldCcr6Mask)|(uint32(value)<<RegisterCcr6FieldCcr6Shift))
}

// registerAf1Type TIM1 alternate function option register 1
type registerAf1Type uint32

const (
	RegisterAf1FieldBkineShift = 0
	RegisterAf1FieldBkineMask  = 0x1
)

// GetBkine BRK BKIN input enable
func (r *registerAf1Type) GetBkine() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkineMask) != 0
}

// SetBkine BRK BKIN input enable
func (r *registerAf1Type) SetBkine(value bool) {
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
func (r *registerAf1Type) GetBkcmp1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkcmp1eMask) != 0
}

// SetBkcmp1e BRK COMP1 enable
func (r *registerAf1Type) SetBkcmp1e(value bool) {
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
func (r *registerAf1Type) GetBkcmp2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkcmp2eMask) != 0
}

// SetBkcmp2e BRK COMP2 enable
func (r *registerAf1Type) SetBkcmp2e(value bool) {
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
func (r *registerAf1Type) GetBkdf1bk0e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkdf1bk0eMask) != 0
}

// SetBkdf1bk0e BRK dfsdm1_break[0] enable
func (r *registerAf1Type) SetBkdf1bk0e(value bool) {
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
func (r *registerAf1Type) GetBkinp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkinpMask) != 0
}

// SetBkinp BRK BKIN input polarity
func (r *registerAf1Type) SetBkinp(value bool) {
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
func (r *registerAf1Type) GetBkcmp1p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkcmp1pMask) != 0
}

// SetBkcmp1p BRK COMP1 input polarity
func (r *registerAf1Type) SetBkcmp1p(value bool) {
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
func (r *registerAf1Type) GetBkcmp2p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldBkcmp2pMask) != 0
}

// SetBkcmp2p BRK COMP2 input polarity
func (r *registerAf1Type) SetBkcmp2p(value bool) {
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
func (r *registerAf1Type) GetEtrsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAf1FieldEtrselMask) >> RegisterAf1FieldEtrselShift)
}

// SetEtrsel ETR source selection
func (r *registerAf1Type) SetEtrsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAf1FieldEtrselMask)|(uint32(value)<<RegisterAf1FieldEtrselShift))
}

// registerAf2Type TIM1 Alternate function odfsdm1_breakster 2
type registerAf2Type uint32

const (
	RegisterAf2FieldBk2ineShift = 0
	RegisterAf2FieldBk2ineMask  = 0x1
)

// GetBk2ine BRK2 BKIN input enable
func (r *registerAf2Type) GetBk2ine() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2ineMask) != 0
}

// SetBk2ine BRK2 BKIN input enable
func (r *registerAf2Type) SetBk2ine(value bool) {
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
func (r *registerAf2Type) GetBk2cmp1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2cmp1eMask) != 0
}

// SetBk2cmp1e BRK2 COMP1 enable
func (r *registerAf2Type) SetBk2cmp1e(value bool) {
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
func (r *registerAf2Type) GetBk2cmp2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2cmp2eMask) != 0
}

// SetBk2cmp2e BRK2 COMP2 enable
func (r *registerAf2Type) SetBk2cmp2e(value bool) {
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
func (r *registerAf2Type) GetBk2df1bk1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2df1bk1eMask) != 0
}

// SetBk2df1bk1e BRK2 dfsdm1_break[1] enable
func (r *registerAf2Type) SetBk2df1bk1e(value bool) {
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
func (r *registerAf2Type) GetBk2inp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2inpMask) != 0
}

// SetBk2inp BRK2 BKIN2 input polarity
func (r *registerAf2Type) SetBk2inp(value bool) {
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
func (r *registerAf2Type) GetBk2cmp1p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2cmp1pMask) != 0
}

// SetBk2cmp1p BRK2 COMP1 input polarit
func (r *registerAf2Type) SetBk2cmp1p(value bool) {
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
func (r *registerAf2Type) GetBk2cmp2p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAf2FieldBk2cmp2pMask) != 0
}

// SetBk2cmp2p BRK2 COMP2 input polarity
func (r *registerAf2Type) SetBk2cmp2p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAf2FieldBk2cmp2pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAf2FieldBk2cmp2pMask)
	}
}

// registerTiselType TIM1 timer input selection register
type registerTiselType uint32

const (
	RegisterTiselFieldTi1selShift = 0
	RegisterTiselFieldTi1selMask  = 0xf
)

// GetTi1sel selects TI1[0] to TI1[15] input
func (r *registerTiselType) GetTi1sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi1selMask) >> RegisterTiselFieldTi1selShift)
}

// SetTi1sel selects TI1[0] to TI1[15] input
func (r *registerTiselType) SetTi1sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi1selMask)|(uint32(value)<<RegisterTiselFieldTi1selShift))
}

const (
	RegisterTiselFieldTi2selShift = 8
	RegisterTiselFieldTi2selMask  = 0xf00
)

// GetTi2sel selects TI2[0] to TI2[15] input
func (r *registerTiselType) GetTi2sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi2selMask) >> RegisterTiselFieldTi2selShift)
}

// SetTi2sel selects TI2[0] to TI2[15] input
func (r *registerTiselType) SetTi2sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi2selMask)|(uint32(value)<<RegisterTiselFieldTi2selShift))
}

const (
	RegisterTiselFieldTi3selShift = 16
	RegisterTiselFieldTi3selMask  = 0xf0000
)

// GetTi3sel selects TI3[0] to TI3[15] input
func (r *registerTiselType) GetTi3sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi3selMask) >> RegisterTiselFieldTi3selShift)
}

// SetTi3sel selects TI3[0] to TI3[15] input
func (r *registerTiselType) SetTi3sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi3selMask)|(uint32(value)<<RegisterTiselFieldTi3selShift))
}

const (
	RegisterTiselFieldTi4selShift = 24
	RegisterTiselFieldTi4selMask  = 0xf000000
)

// GetTi4sel selects TI4[0] to TI4[15] input
func (r *registerTiselType) GetTi4sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi4selMask) >> RegisterTiselFieldTi4selShift)
}

// SetTi4sel selects TI4[0] to TI4[15] input
func (r *registerTiselType) SetTi4sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi4selMask)|(uint32(value)<<RegisterTiselFieldTi4selShift))
}
