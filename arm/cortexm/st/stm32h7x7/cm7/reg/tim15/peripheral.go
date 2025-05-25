package tim15

import (
	"unsafe"
	"volatile"
)

var (
	Tim15 = (*_tim15)(unsafe.Pointer(uintptr(0x40014000)))
)

type _tim15 struct {
	Cr1          registerCr1Type
	Cr2          registerCr2Type
	Smcr         registerSmcrType
	Dier         registerDierType
	Sr           registerSrType
	Egr          registerEgrType
	Ccmr1_output registerCcmr1_outputType
	Ccmr1_input  registerCcmr1_inputType
	_            [4]byte
	Ccer         registerCcerType
	Cnt          registerCntType
	Psc          registerPscType
	Arr          registerArrType
	Rcr          registerRcrType
	Ccr1         registerCcr1Type
	Ccr2         registerCcr2Type
	_            [8]byte
	Bdtr         registerBdtrType
	Dcr          registerDcrType
	Dmar         registerDmarType
	_            [16]byte
	Af1          registerAf1Type
	_            [4]byte
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
	RegisterSmcrFieldTs_2_0Shift = 4
	RegisterSmcrFieldTs_2_0Mask  = 0x70
)

// GetTs_2_0 Trigger selection
func (r *registerSmcrType) GetTs_2_0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldTs_2_0Mask) >> RegisterSmcrFieldTs_2_0Shift)
}

// SetTs_2_0 Trigger selection
func (r *registerSmcrType) SetTs_2_0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmcrFieldTs_2_0Mask)|(uint32(value)<<RegisterSmcrFieldTs_2_0Shift))
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
	RegisterSmcrFieldSms_3Shift = 16
	RegisterSmcrFieldSms_3Mask  = 0x10000
)

// GetSms_3 Slave mode selection bit 3
func (r *registerSmcrType) GetSms_3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSmcrFieldSms_3Mask) != 0
}

// SetSms_3 Slave mode selection bit 3
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

// registerCcmr1_outputType capture/compare mode register (output mode)
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
	RegisterCcmr1_outputFieldOc1m_3Shift = 16
	RegisterCcmr1_outputFieldOc1m_3Mask  = 0x10000
)

// GetOc1m_3 Output Compare 1 mode bit 3
func (r *registerCcmr1_outputType) GetOc1m_3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldOc1m_3Mask) != 0
}

// SetOc1m_3 Output Compare 1 mode bit 3
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

// GetOc2m_3 Output Compare 2 mode bit 3
func (r *registerCcmr1_outputType) GetOc2m_3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_outputFieldOc2m_3Mask) != 0
}

// SetOc2m_3 Output Compare 2 mode bit 3
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
	RegisterCcmr1_inputFieldIc1pscShift = 2
	RegisterCcmr1_inputFieldIc1pscMask  = 0xc
)

// GetIc1psc Input capture 1 prescaler
func (r *registerCcmr1_inputType) GetIc1psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_inputFieldIc1pscMask) >> RegisterCcmr1_inputFieldIc1pscShift)
}

// SetIc1psc Input capture 1 prescaler
func (r *registerCcmr1_inputType) SetIc1psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_inputFieldIc1pscMask)|(uint32(value)<<RegisterCcmr1_inputFieldIc1pscShift))
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
	RegisterCcmr1_inputFieldIc2pscShift = 10
	RegisterCcmr1_inputFieldIc2pscMask  = 0xc00
)

// GetIc2psc Input capture 2 prescaler
func (r *registerCcmr1_inputType) GetIc2psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmr1_inputFieldIc2pscMask) >> RegisterCcmr1_inputFieldIc2pscShift)
}

// SetIc2psc Input capture 2 prescaler
func (r *registerCcmr1_inputType) SetIc2psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmr1_inputFieldIc2pscMask)|(uint32(value)<<RegisterCcmr1_inputFieldIc2pscShift))
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

// registerAf1Type TIM15 alternate fdfsdm1_breakon register 1
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

// registerTiselType TIM15 input selection register
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
