//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package adc

import (
	"unsafe"
	"volatile"
)

var (
	Adc1 = (*_adc)(unsafe.Pointer(uintptr(0x40022000)))
	Adc2 = (*_adc)(unsafe.Pointer(uintptr(0x40022100)))
	Adc3 = (*_adc)(unsafe.Pointer(uintptr(0x58026000)))

	Instances = [3]*_adc{
		Adc1,
		Adc2,
		Adc3,
	}
)

type _adc struct {
	Isr      RegisterIsrType
	Ier      RegisterIerType
	Cr       RegisterCrType
	Cfgr     RegisterCfgrType
	Cfgr2    RegisterCfgr2Type
	Smpr1    RegisterSmpr1Type
	Smpr2    RegisterSmpr2Type
	Pcsel    RegisterPcselType
	Ltr1     RegisterLtr1Type
	Lhtr1    RegisterLhtr1Type
	_        [8]byte
	Sqr1     RegisterSqr1Type
	Sqr2     RegisterSqr2Type
	Sqr3     RegisterSqr3Type
	Sqr4     RegisterSqr4Type
	Dr       RegisterDrType
	_        [8]byte
	Jsqr     RegisterJsqrType
	_        [16]byte
	Ofr1     RegisterOfr1Type
	Ofr2     RegisterOfr2Type
	Ofr3     RegisterOfr3Type
	Ofr4     RegisterOfr4Type
	_        [16]byte
	Jdr1     RegisterJdr1Type
	Jdr2     RegisterJdr2Type
	Jdr3     RegisterJdr3Type
	Jdr4     RegisterJdr4Type
	_        [16]byte
	Awd2cr   RegisterAwd2crType
	Awd3cr   RegisterAwd3crType
	_        [8]byte
	Ltr2     RegisterLtr2Type
	Htr2     RegisterHtr2Type
	Ltr3     RegisterLtr3Type
	Htr3     RegisterHtr3Type
	Difsel   RegisterDifselType
	Calfact  RegisterCalfactType
	Calfact2 RegisterCalfact2Type
}

// RegisterIsrType ADC interrupt and status register
type RegisterIsrType uint32

func (r *RegisterIsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIsrFieldAdrdyShift = 0
	RegisterIsrFieldAdrdyMask  = 0x1
)

// GetAdrdy ADC ready flag
func (r *RegisterIsrType) GetAdrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldAdrdyMask) != 0
}

// SetAdrdy ADC ready flag
func (r *RegisterIsrType) SetAdrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldAdrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldAdrdyMask)
	}
}

const (
	RegisterIsrFieldEosmpShift = 1
	RegisterIsrFieldEosmpMask  = 0x2
)

// GetEosmp ADC group regular end of sampling flag
func (r *RegisterIsrType) GetEosmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldEosmpMask) != 0
}

// SetEosmp ADC group regular end of sampling flag
func (r *RegisterIsrType) SetEosmp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldEosmpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldEosmpMask)
	}
}

const (
	RegisterIsrFieldEocShift = 2
	RegisterIsrFieldEocMask  = 0x4
)

// GetEoc ADC group regular end of unitary conversion flag
func (r *RegisterIsrType) GetEoc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldEocMask) != 0
}

// SetEoc ADC group regular end of unitary conversion flag
func (r *RegisterIsrType) SetEoc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldEocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldEocMask)
	}
}

const (
	RegisterIsrFieldEosShift = 3
	RegisterIsrFieldEosMask  = 0x8
)

// GetEos ADC group regular end of sequence conversions flag
func (r *RegisterIsrType) GetEos() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldEosMask) != 0
}

// SetEos ADC group regular end of sequence conversions flag
func (r *RegisterIsrType) SetEos(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldEosMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldEosMask)
	}
}

const (
	RegisterIsrFieldOvrShift = 4
	RegisterIsrFieldOvrMask  = 0x10
)

// GetOvr ADC group regular overrun flag
func (r *RegisterIsrType) GetOvr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldOvrMask) != 0
}

// SetOvr ADC group regular overrun flag
func (r *RegisterIsrType) SetOvr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldOvrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldOvrMask)
	}
}

const (
	RegisterIsrFieldJeocShift = 5
	RegisterIsrFieldJeocMask  = 0x20
)

// GetJeoc ADC group injected end of unitary conversion flag
func (r *RegisterIsrType) GetJeoc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldJeocMask) != 0
}

// SetJeoc ADC group injected end of unitary conversion flag
func (r *RegisterIsrType) SetJeoc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldJeocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldJeocMask)
	}
}

const (
	RegisterIsrFieldJeosShift = 6
	RegisterIsrFieldJeosMask  = 0x40
)

// GetJeos ADC group injected end of sequence conversions flag
func (r *RegisterIsrType) GetJeos() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldJeosMask) != 0
}

// SetJeos ADC group injected end of sequence conversions flag
func (r *RegisterIsrType) SetJeos(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldJeosMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldJeosMask)
	}
}

const (
	RegisterIsrFieldAwd1Shift = 7
	RegisterIsrFieldAwd1Mask  = 0x80
)

// GetAwd1 ADC analog watchdog 1 flag
func (r *RegisterIsrType) GetAwd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldAwd1Mask) != 0
}

// SetAwd1 ADC analog watchdog 1 flag
func (r *RegisterIsrType) SetAwd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldAwd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldAwd1Mask)
	}
}

const (
	RegisterIsrFieldAwd2Shift = 8
	RegisterIsrFieldAwd2Mask  = 0x100
)

// GetAwd2 ADC analog watchdog 2 flag
func (r *RegisterIsrType) GetAwd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldAwd2Mask) != 0
}

// SetAwd2 ADC analog watchdog 2 flag
func (r *RegisterIsrType) SetAwd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldAwd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldAwd2Mask)
	}
}

const (
	RegisterIsrFieldAwd3Shift = 9
	RegisterIsrFieldAwd3Mask  = 0x200
)

// GetAwd3 ADC analog watchdog 3 flag
func (r *RegisterIsrType) GetAwd3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldAwd3Mask) != 0
}

// SetAwd3 ADC analog watchdog 3 flag
func (r *RegisterIsrType) SetAwd3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldAwd3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldAwd3Mask)
	}
}

const (
	RegisterIsrFieldJqovfShift = 10
	RegisterIsrFieldJqovfMask  = 0x400
)

// GetJqovf ADC group injected contexts queue overflow flag
func (r *RegisterIsrType) GetJqovf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldJqovfMask) != 0
}

// SetJqovf ADC group injected contexts queue overflow flag
func (r *RegisterIsrType) SetJqovf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldJqovfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldJqovfMask)
	}
}

// RegisterIerType ADC interrupt enable register
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
	RegisterIerFieldAdrdyieShift = 0
	RegisterIerFieldAdrdyieMask  = 0x1
)

// GetAdrdyie ADC ready interrupt
func (r *RegisterIerType) GetAdrdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldAdrdyieMask) != 0
}

// SetAdrdyie ADC ready interrupt
func (r *RegisterIerType) SetAdrdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldAdrdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldAdrdyieMask)
	}
}

const (
	RegisterIerFieldEosmpieShift = 1
	RegisterIerFieldEosmpieMask  = 0x2
)

// GetEosmpie ADC group regular end of sampling interrupt
func (r *RegisterIerType) GetEosmpie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldEosmpieMask) != 0
}

// SetEosmpie ADC group regular end of sampling interrupt
func (r *RegisterIerType) SetEosmpie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldEosmpieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldEosmpieMask)
	}
}

const (
	RegisterIerFieldEocieShift = 2
	RegisterIerFieldEocieMask  = 0x4
)

// GetEocie ADC group regular end of unitary conversion interrupt
func (r *RegisterIerType) GetEocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldEocieMask) != 0
}

// SetEocie ADC group regular end of unitary conversion interrupt
func (r *RegisterIerType) SetEocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldEocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldEocieMask)
	}
}

const (
	RegisterIerFieldEosieShift = 3
	RegisterIerFieldEosieMask  = 0x8
)

// GetEosie ADC group regular end of sequence conversions interrupt
func (r *RegisterIerType) GetEosie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldEosieMask) != 0
}

// SetEosie ADC group regular end of sequence conversions interrupt
func (r *RegisterIerType) SetEosie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldEosieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldEosieMask)
	}
}

const (
	RegisterIerFieldOvrieShift = 4
	RegisterIerFieldOvrieMask  = 0x10
)

// GetOvrie ADC group regular overrun interrupt
func (r *RegisterIerType) GetOvrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldOvrieMask) != 0
}

// SetOvrie ADC group regular overrun interrupt
func (r *RegisterIerType) SetOvrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldOvrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldOvrieMask)
	}
}

const (
	RegisterIerFieldJeocieShift = 5
	RegisterIerFieldJeocieMask  = 0x20
)

// GetJeocie ADC group injected end of unitary conversion interrupt
func (r *RegisterIerType) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldJeocieMask) != 0
}

// SetJeocie ADC group injected end of unitary conversion interrupt
func (r *RegisterIerType) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldJeocieMask)
	}
}

const (
	RegisterIerFieldJeosieShift = 6
	RegisterIerFieldJeosieMask  = 0x40
)

// GetJeosie ADC group injected end of sequence conversions interrupt
func (r *RegisterIerType) GetJeosie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldJeosieMask) != 0
}

// SetJeosie ADC group injected end of sequence conversions interrupt
func (r *RegisterIerType) SetJeosie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldJeosieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldJeosieMask)
	}
}

const (
	RegisterIerFieldAwd1ieShift = 7
	RegisterIerFieldAwd1ieMask  = 0x80
)

// GetAwd1ie ADC analog watchdog 1 interrupt
func (r *RegisterIerType) GetAwd1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldAwd1ieMask) != 0
}

// SetAwd1ie ADC analog watchdog 1 interrupt
func (r *RegisterIerType) SetAwd1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldAwd1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldAwd1ieMask)
	}
}

const (
	RegisterIerFieldAwd2ieShift = 8
	RegisterIerFieldAwd2ieMask  = 0x100
)

// GetAwd2ie ADC analog watchdog 2 interrupt
func (r *RegisterIerType) GetAwd2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldAwd2ieMask) != 0
}

// SetAwd2ie ADC analog watchdog 2 interrupt
func (r *RegisterIerType) SetAwd2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldAwd2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldAwd2ieMask)
	}
}

const (
	RegisterIerFieldAwd3ieShift = 9
	RegisterIerFieldAwd3ieMask  = 0x200
)

// GetAwd3ie ADC analog watchdog 3 interrupt
func (r *RegisterIerType) GetAwd3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldAwd3ieMask) != 0
}

// SetAwd3ie ADC analog watchdog 3 interrupt
func (r *RegisterIerType) SetAwd3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldAwd3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldAwd3ieMask)
	}
}

const (
	RegisterIerFieldJqovfieShift = 10
	RegisterIerFieldJqovfieMask  = 0x400
)

// GetJqovfie ADC group injected contexts queue overflow interrupt
func (r *RegisterIerType) GetJqovfie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldJqovfieMask) != 0
}

// SetJqovfie ADC group injected contexts queue overflow interrupt
func (r *RegisterIerType) SetJqovfie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldJqovfieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldJqovfieMask)
	}
}

// RegisterCrType ADC control register
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
	RegisterCrFieldAdenShift = 0
	RegisterCrFieldAdenMask  = 0x1
)

// GetAden ADC enable
func (r *RegisterCrType) GetAden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAdenMask) != 0
}

// SetAden ADC enable
func (r *RegisterCrType) SetAden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAdenMask)
	}
}

const (
	RegisterCrFieldAddisShift = 1
	RegisterCrFieldAddisMask  = 0x2
)

// GetAddis ADC disable
func (r *RegisterCrType) GetAddis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAddisMask) != 0
}

// SetAddis ADC disable
func (r *RegisterCrType) SetAddis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAddisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAddisMask)
	}
}

const (
	RegisterCrFieldAdstartShift = 2
	RegisterCrFieldAdstartMask  = 0x4
)

// GetAdstart ADC group regular conversion start
func (r *RegisterCrType) GetAdstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAdstartMask) != 0
}

// SetAdstart ADC group regular conversion start
func (r *RegisterCrType) SetAdstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAdstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAdstartMask)
	}
}

const (
	RegisterCrFieldJadstartShift = 3
	RegisterCrFieldJadstartMask  = 0x8
)

// GetJadstart ADC group injected conversion start
func (r *RegisterCrType) GetJadstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldJadstartMask) != 0
}

// SetJadstart ADC group injected conversion start
func (r *RegisterCrType) SetJadstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldJadstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldJadstartMask)
	}
}

const (
	RegisterCrFieldAdstpShift = 4
	RegisterCrFieldAdstpMask  = 0x10
)

// GetAdstp ADC group regular conversion stop
func (r *RegisterCrType) GetAdstp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAdstpMask) != 0
}

// SetAdstp ADC group regular conversion stop
func (r *RegisterCrType) SetAdstp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAdstpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAdstpMask)
	}
}

const (
	RegisterCrFieldJadstpShift = 5
	RegisterCrFieldJadstpMask  = 0x20
)

// GetJadstp ADC group injected conversion stop
func (r *RegisterCrType) GetJadstp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldJadstpMask) != 0
}

// SetJadstp ADC group injected conversion stop
func (r *RegisterCrType) SetJadstp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldJadstpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldJadstpMask)
	}
}

const (
	RegisterCrFieldBoostShift = 8
	RegisterCrFieldBoostMask  = 0x300
)

// GetBoost Boost mode control
func (r *RegisterCrType) GetBoost() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldBoostMask) >> RegisterCrFieldBoostShift)
}

// SetBoost Boost mode control
func (r *RegisterCrType) SetBoost(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldBoostMask)|(uint32(value)<<RegisterCrFieldBoostShift))
}

const (
	RegisterCrFieldAdcallinShift = 16
	RegisterCrFieldAdcallinMask  = 0x10000
)

// GetAdcallin Linearity calibration
func (r *RegisterCrType) GetAdcallin() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAdcallinMask) != 0
}

// SetAdcallin Linearity calibration
func (r *RegisterCrType) SetAdcallin(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAdcallinMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAdcallinMask)
	}
}

const (
	RegisterCrFieldLincalrdyw1Shift = 22
	RegisterCrFieldLincalrdyw1Mask  = 0x400000
)

// GetLincalrdyw1 Linearity calibration ready Word 1
func (r *RegisterCrType) GetLincalrdyw1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLincalrdyw1Mask) != 0
}

// SetLincalrdyw1 Linearity calibration ready Word 1
func (r *RegisterCrType) SetLincalrdyw1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldLincalrdyw1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldLincalrdyw1Mask)
	}
}

const (
	RegisterCrFieldLincalrdyw2Shift = 23
	RegisterCrFieldLincalrdyw2Mask  = 0x800000
)

// GetLincalrdyw2 Linearity calibration ready Word 2
func (r *RegisterCrType) GetLincalrdyw2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLincalrdyw2Mask) != 0
}

// SetLincalrdyw2 Linearity calibration ready Word 2
func (r *RegisterCrType) SetLincalrdyw2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldLincalrdyw2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldLincalrdyw2Mask)
	}
}

const (
	RegisterCrFieldLincalrdyw3Shift = 24
	RegisterCrFieldLincalrdyw3Mask  = 0x1000000
)

// GetLincalrdyw3 Linearity calibration ready Word 3
func (r *RegisterCrType) GetLincalrdyw3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLincalrdyw3Mask) != 0
}

// SetLincalrdyw3 Linearity calibration ready Word 3
func (r *RegisterCrType) SetLincalrdyw3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldLincalrdyw3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldLincalrdyw3Mask)
	}
}

const (
	RegisterCrFieldLincalrdyw4Shift = 25
	RegisterCrFieldLincalrdyw4Mask  = 0x2000000
)

// GetLincalrdyw4 Linearity calibration ready Word 4
func (r *RegisterCrType) GetLincalrdyw4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLincalrdyw4Mask) != 0
}

// SetLincalrdyw4 Linearity calibration ready Word 4
func (r *RegisterCrType) SetLincalrdyw4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldLincalrdyw4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldLincalrdyw4Mask)
	}
}

const (
	RegisterCrFieldLincalrdyw5Shift = 26
	RegisterCrFieldLincalrdyw5Mask  = 0x4000000
)

// GetLincalrdyw5 Linearity calibration ready Word 5
func (r *RegisterCrType) GetLincalrdyw5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLincalrdyw5Mask) != 0
}

// SetLincalrdyw5 Linearity calibration ready Word 5
func (r *RegisterCrType) SetLincalrdyw5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldLincalrdyw5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldLincalrdyw5Mask)
	}
}

const (
	RegisterCrFieldLincalrdyw6Shift = 27
	RegisterCrFieldLincalrdyw6Mask  = 0x8000000
)

// GetLincalrdyw6 Linearity calibration ready Word 6
func (r *RegisterCrType) GetLincalrdyw6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLincalrdyw6Mask) != 0
}

// SetLincalrdyw6 Linearity calibration ready Word 6
func (r *RegisterCrType) SetLincalrdyw6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldLincalrdyw6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldLincalrdyw6Mask)
	}
}

const (
	RegisterCrFieldAdvregenShift = 28
	RegisterCrFieldAdvregenMask  = 0x10000000
)

// GetAdvregen ADC voltage regulator enable
func (r *RegisterCrType) GetAdvregen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAdvregenMask) != 0
}

// SetAdvregen ADC voltage regulator enable
func (r *RegisterCrType) SetAdvregen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAdvregenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAdvregenMask)
	}
}

const (
	RegisterCrFieldDeeppwdShift = 29
	RegisterCrFieldDeeppwdMask  = 0x20000000
)

// GetDeeppwd ADC deep power down enable
func (r *RegisterCrType) GetDeeppwd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDeeppwdMask) != 0
}

// SetDeeppwd ADC deep power down enable
func (r *RegisterCrType) SetDeeppwd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDeeppwdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDeeppwdMask)
	}
}

const (
	RegisterCrFieldAdcaldifShift = 30
	RegisterCrFieldAdcaldifMask  = 0x40000000
)

// GetAdcaldif ADC differential mode for calibration
func (r *RegisterCrType) GetAdcaldif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAdcaldifMask) != 0
}

// SetAdcaldif ADC differential mode for calibration
func (r *RegisterCrType) SetAdcaldif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAdcaldifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAdcaldifMask)
	}
}

const (
	RegisterCrFieldAdcalShift = 31
	RegisterCrFieldAdcalMask  = 0x80000000
)

// GetAdcal ADC calibration
func (r *RegisterCrType) GetAdcal() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAdcalMask) != 0
}

// SetAdcal ADC calibration
func (r *RegisterCrType) SetAdcal(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAdcalMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAdcalMask)
	}
}

// RegisterCfgrType ADC configuration register 1
type RegisterCfgrType uint32

func (r *RegisterCfgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCfgrFieldDmngtShift = 0
	RegisterCfgrFieldDmngtMask  = 0x3
)

// GetDmngt ADC DMA transfer enable
func (r *RegisterCfgrType) GetDmngt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldDmngtMask) >> RegisterCfgrFieldDmngtShift)
}

// SetDmngt ADC DMA transfer enable
func (r *RegisterCfgrType) SetDmngt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldDmngtMask)|(uint32(value)<<RegisterCfgrFieldDmngtShift))
}

type RegisterCfgrFieldResEnumType uint8

const (
	// RegisterCfgrFieldResEnumResolution16bit 16 bits
	RegisterCfgrFieldResEnumResolution16bit RegisterCfgrFieldResEnumType = 0x0

	// RegisterCfgrFieldResEnumResolution14bitl 14 bits in legacy mode (not optimized power consumption)
	RegisterCfgrFieldResEnumResolution14bitl RegisterCfgrFieldResEnumType = 0x1

	// RegisterCfgrFieldResEnumResolution12bitl 12 bits in legacy mode (not optimized power consumption)
	RegisterCfgrFieldResEnumResolution12bitl RegisterCfgrFieldResEnumType = 0x2

	// RegisterCfgrFieldResEnumResolution14bit 14 bits
	RegisterCfgrFieldResEnumResolution14bit RegisterCfgrFieldResEnumType = 0x5

	// RegisterCfgrFieldResEnumResolution12bit 12 bits
	RegisterCfgrFieldResEnumResolution12bit RegisterCfgrFieldResEnumType = 0x6

	// RegisterCfgrFieldResEnumResolution10bit 10 bits
	RegisterCfgrFieldResEnumResolution10bit RegisterCfgrFieldResEnumType = 0x3

	// RegisterCfgrFieldResEnumResolution8bit 8 bits
	RegisterCfgrFieldResEnumResolution8bit RegisterCfgrFieldResEnumType = 0x7

	RegisterCfgrFieldResShift = 2
	RegisterCfgrFieldResMask  = 0x1c
)

// GetRes ADC data resolution
func (r *RegisterCfgrType) GetRes() RegisterCfgrFieldResEnumType {
	return RegisterCfgrFieldResEnumType((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldResMask) >> RegisterCfgrFieldResShift)
}

// SetRes ADC data resolution
func (r *RegisterCfgrType) SetRes(value RegisterCfgrFieldResEnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldResMask)|(uint32(value)<<RegisterCfgrFieldResShift))
}

const (
	RegisterCfgrFieldExtselShift = 5
	RegisterCfgrFieldExtselMask  = 0x3e0
)

// GetExtsel ADC group regular external trigger source
func (r *RegisterCfgrType) GetExtsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldExtselMask) >> RegisterCfgrFieldExtselShift)
}

// SetExtsel ADC group regular external trigger source
func (r *RegisterCfgrType) SetExtsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldExtselMask)|(uint32(value)<<RegisterCfgrFieldExtselShift))
}

const (
	RegisterCfgrFieldExtenShift = 10
	RegisterCfgrFieldExtenMask  = 0xc00
)

// GetExten ADC group regular external trigger polarity
func (r *RegisterCfgrType) GetExten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldExtenMask) >> RegisterCfgrFieldExtenShift)
}

// SetExten ADC group regular external trigger polarity
func (r *RegisterCfgrType) SetExten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldExtenMask)|(uint32(value)<<RegisterCfgrFieldExtenShift))
}

const (
	RegisterCfgrFieldOvrmodShift = 12
	RegisterCfgrFieldOvrmodMask  = 0x1000
)

// GetOvrmod ADC group regular overrun configuration
func (r *RegisterCfgrType) GetOvrmod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldOvrmodMask) != 0
}

// SetOvrmod ADC group regular overrun configuration
func (r *RegisterCfgrType) SetOvrmod(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldOvrmodMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldOvrmodMask)
	}
}

const (
	RegisterCfgrFieldContShift = 13
	RegisterCfgrFieldContMask  = 0x2000
)

// GetCont ADC group regular continuous conversion mode
func (r *RegisterCfgrType) GetCont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldContMask) != 0
}

// SetCont ADC group regular continuous conversion mode
func (r *RegisterCfgrType) SetCont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldContMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldContMask)
	}
}

const (
	RegisterCfgrFieldAutdlyShift = 14
	RegisterCfgrFieldAutdlyMask  = 0x4000
)

// GetAutdly ADC low power auto wait
func (r *RegisterCfgrType) GetAutdly() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldAutdlyMask) != 0
}

// SetAutdly ADC low power auto wait
func (r *RegisterCfgrType) SetAutdly(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldAutdlyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldAutdlyMask)
	}
}

const (
	RegisterCfgrFieldDiscenShift = 16
	RegisterCfgrFieldDiscenMask  = 0x10000
)

// GetDiscen ADC group regular sequencer discontinuous mode
func (r *RegisterCfgrType) GetDiscen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldDiscenMask) != 0
}

// SetDiscen ADC group regular sequencer discontinuous mode
func (r *RegisterCfgrType) SetDiscen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldDiscenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldDiscenMask)
	}
}

const (
	RegisterCfgrFieldDiscnumShift = 17
	RegisterCfgrFieldDiscnumMask  = 0xe0000
)

// GetDiscnum ADC group regular sequencer discontinuous number of ranks
func (r *RegisterCfgrType) GetDiscnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldDiscnumMask) >> RegisterCfgrFieldDiscnumShift)
}

// SetDiscnum ADC group regular sequencer discontinuous number of ranks
func (r *RegisterCfgrType) SetDiscnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldDiscnumMask)|(uint32(value)<<RegisterCfgrFieldDiscnumShift))
}

const (
	RegisterCfgrFieldJdiscenShift = 20
	RegisterCfgrFieldJdiscenMask  = 0x100000
)

// GetJdiscen ADC group injected sequencer discontinuous mode
func (r *RegisterCfgrType) GetJdiscen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldJdiscenMask) != 0
}

// SetJdiscen ADC group injected sequencer discontinuous mode
func (r *RegisterCfgrType) SetJdiscen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldJdiscenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldJdiscenMask)
	}
}

const (
	RegisterCfgrFieldJqmShift = 21
	RegisterCfgrFieldJqmMask  = 0x200000
)

// GetJqm ADC group injected contexts queue mode
func (r *RegisterCfgrType) GetJqm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldJqmMask) != 0
}

// SetJqm ADC group injected contexts queue mode
func (r *RegisterCfgrType) SetJqm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldJqmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldJqmMask)
	}
}

const (
	RegisterCfgrFieldAwd1sglShift = 22
	RegisterCfgrFieldAwd1sglMask  = 0x400000
)

// GetAwd1sgl ADC analog watchdog 1 monitoring a single channel or all channels
func (r *RegisterCfgrType) GetAwd1sgl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldAwd1sglMask) != 0
}

// SetAwd1sgl ADC analog watchdog 1 monitoring a single channel or all channels
func (r *RegisterCfgrType) SetAwd1sgl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldAwd1sglMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldAwd1sglMask)
	}
}

const (
	RegisterCfgrFieldAwd1enShift = 23
	RegisterCfgrFieldAwd1enMask  = 0x800000
)

// GetAwd1en ADC analog watchdog 1 enable on scope ADC group regular
func (r *RegisterCfgrType) GetAwd1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldAwd1enMask) != 0
}

// SetAwd1en ADC analog watchdog 1 enable on scope ADC group regular
func (r *RegisterCfgrType) SetAwd1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldAwd1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldAwd1enMask)
	}
}

const (
	RegisterCfgrFieldJawd1enShift = 24
	RegisterCfgrFieldJawd1enMask  = 0x1000000
)

// GetJawd1en ADC analog watchdog 1 enable on scope ADC group injected
func (r *RegisterCfgrType) GetJawd1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldJawd1enMask) != 0
}

// SetJawd1en ADC analog watchdog 1 enable on scope ADC group injected
func (r *RegisterCfgrType) SetJawd1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldJawd1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldJawd1enMask)
	}
}

const (
	RegisterCfgrFieldJautoShift = 25
	RegisterCfgrFieldJautoMask  = 0x2000000
)

// GetJauto ADC group injected automatic trigger mode
func (r *RegisterCfgrType) GetJauto() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldJautoMask) != 0
}

// SetJauto ADC group injected automatic trigger mode
func (r *RegisterCfgrType) SetJauto(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldJautoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldJautoMask)
	}
}

const (
	RegisterCfgrFieldAwdch1chShift = 26
	RegisterCfgrFieldAwdch1chMask  = 0x7c000000
)

// GetAwdch1ch ADC analog watchdog 1 monitored channel selection
func (r *RegisterCfgrType) GetAwdch1ch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldAwdch1chMask) >> RegisterCfgrFieldAwdch1chShift)
}

// SetAwdch1ch ADC analog watchdog 1 monitored channel selection
func (r *RegisterCfgrType) SetAwdch1ch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldAwdch1chMask)|(uint32(value)<<RegisterCfgrFieldAwdch1chShift))
}

const (
	RegisterCfgrFieldJqdisShift = 31
	RegisterCfgrFieldJqdisMask  = 0x80000000
)

// GetJqdis ADC group injected contexts queue disable
func (r *RegisterCfgrType) GetJqdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldJqdisMask) != 0
}

// SetJqdis ADC group injected contexts queue disable
func (r *RegisterCfgrType) SetJqdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldJqdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldJqdisMask)
	}
}

// RegisterCfgr2Type ADC configuration register 2
type RegisterCfgr2Type uint32

func (r *RegisterCfgr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfgr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfgr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfgr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfgr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCfgr2FieldRovseShift = 0
	RegisterCfgr2FieldRovseMask  = 0x1
)

// GetRovse ADC oversampler enable on scope ADC group regular
func (r *RegisterCfgr2Type) GetRovse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldRovseMask) != 0
}

// SetRovse ADC oversampler enable on scope ADC group regular
func (r *RegisterCfgr2Type) SetRovse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldRovseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldRovseMask)
	}
}

const (
	RegisterCfgr2FieldJovseShift = 1
	RegisterCfgr2FieldJovseMask  = 0x2
)

// GetJovse ADC oversampler enable on scope ADC group injected
func (r *RegisterCfgr2Type) GetJovse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldJovseMask) != 0
}

// SetJovse ADC oversampler enable on scope ADC group injected
func (r *RegisterCfgr2Type) SetJovse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldJovseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldJovseMask)
	}
}

const (
	RegisterCfgr2FieldOvssShift = 5
	RegisterCfgr2FieldOvssMask  = 0x1e0
)

// GetOvss ADC oversampling shift
func (r *RegisterCfgr2Type) GetOvss() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldOvssMask) >> RegisterCfgr2FieldOvssShift)
}

// SetOvss ADC oversampling shift
func (r *RegisterCfgr2Type) SetOvss(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldOvssMask)|(uint32(value)<<RegisterCfgr2FieldOvssShift))
}

const (
	RegisterCfgr2FieldTrovsShift = 9
	RegisterCfgr2FieldTrovsMask  = 0x200
)

// GetTrovs ADC oversampling discontinuous mode (triggered mode) for ADC group regular
func (r *RegisterCfgr2Type) GetTrovs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldTrovsMask) != 0
}

// SetTrovs ADC oversampling discontinuous mode (triggered mode) for ADC group regular
func (r *RegisterCfgr2Type) SetTrovs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldTrovsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldTrovsMask)
	}
}

const (
	RegisterCfgr2FieldRovsmShift = 10
	RegisterCfgr2FieldRovsmMask  = 0x400
)

// GetRovsm Regular Oversampling mode
func (r *RegisterCfgr2Type) GetRovsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldRovsmMask) != 0
}

// SetRovsm Regular Oversampling mode
func (r *RegisterCfgr2Type) SetRovsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldRovsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldRovsmMask)
	}
}

const (
	RegisterCfgr2FieldRshift1Shift = 11
	RegisterCfgr2FieldRshift1Mask  = 0x800
)

// GetRshift1 Right-shift data after Offset 1 correction
func (r *RegisterCfgr2Type) GetRshift1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldRshift1Mask) != 0
}

// SetRshift1 Right-shift data after Offset 1 correction
func (r *RegisterCfgr2Type) SetRshift1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldRshift1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldRshift1Mask)
	}
}

const (
	RegisterCfgr2FieldRshift2Shift = 12
	RegisterCfgr2FieldRshift2Mask  = 0x1000
)

// GetRshift2 Right-shift data after Offset 2 correction
func (r *RegisterCfgr2Type) GetRshift2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldRshift2Mask) != 0
}

// SetRshift2 Right-shift data after Offset 2 correction
func (r *RegisterCfgr2Type) SetRshift2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldRshift2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldRshift2Mask)
	}
}

const (
	RegisterCfgr2FieldRshift3Shift = 13
	RegisterCfgr2FieldRshift3Mask  = 0x2000
)

// GetRshift3 Right-shift data after Offset 3 correction
func (r *RegisterCfgr2Type) GetRshift3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldRshift3Mask) != 0
}

// SetRshift3 Right-shift data after Offset 3 correction
func (r *RegisterCfgr2Type) SetRshift3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldRshift3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldRshift3Mask)
	}
}

const (
	RegisterCfgr2FieldRshift4Shift = 14
	RegisterCfgr2FieldRshift4Mask  = 0x4000
)

// GetRshift4 Right-shift data after Offset 4 correction
func (r *RegisterCfgr2Type) GetRshift4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldRshift4Mask) != 0
}

// SetRshift4 Right-shift data after Offset 4 correction
func (r *RegisterCfgr2Type) SetRshift4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldRshift4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldRshift4Mask)
	}
}

const (
	RegisterCfgr2FieldOsrShift = 16
	RegisterCfgr2FieldOsrMask  = 0x3ff0000
)

// GetOsr Oversampling ratio
func (r *RegisterCfgr2Type) GetOsr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldOsrMask) >> RegisterCfgr2FieldOsrShift)
}

// SetOsr Oversampling ratio
func (r *RegisterCfgr2Type) SetOsr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldOsrMask)|(uint32(value)<<RegisterCfgr2FieldOsrShift))
}

const (
	RegisterCfgr2FieldLshiftShift = 28
	RegisterCfgr2FieldLshiftMask  = 0xf0000000
)

// GetLshift Left shift factor
func (r *RegisterCfgr2Type) GetLshift() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldLshiftMask) >> RegisterCfgr2FieldLshiftShift)
}

// SetLshift Left shift factor
func (r *RegisterCfgr2Type) SetLshift(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldLshiftMask)|(uint32(value)<<RegisterCfgr2FieldLshiftShift))
}

// RegisterSmpr1Type ADC sampling time register 1
type RegisterSmpr1Type uint32

func (r *RegisterSmpr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSmpr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSmpr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSmpr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSmpr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSmpr1FieldSmp1Shift = 3
	RegisterSmpr1FieldSmp1Mask  = 0x38
)

// GetSmp1 ADC channel 1 sampling time selection
func (r *RegisterSmpr1Type) GetSmp1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr1FieldSmp1Mask) >> RegisterSmpr1FieldSmp1Shift)
}

// SetSmp1 ADC channel 1 sampling time selection
func (r *RegisterSmpr1Type) SetSmp1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr1FieldSmp1Mask)|(uint32(value)<<RegisterSmpr1FieldSmp1Shift))
}

const (
	RegisterSmpr1FieldSmp2Shift = 6
	RegisterSmpr1FieldSmp2Mask  = 0x1c0
)

// GetSmp2 ADC channel 2 sampling time selection
func (r *RegisterSmpr1Type) GetSmp2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr1FieldSmp2Mask) >> RegisterSmpr1FieldSmp2Shift)
}

// SetSmp2 ADC channel 2 sampling time selection
func (r *RegisterSmpr1Type) SetSmp2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr1FieldSmp2Mask)|(uint32(value)<<RegisterSmpr1FieldSmp2Shift))
}

const (
	RegisterSmpr1FieldSmp3Shift = 9
	RegisterSmpr1FieldSmp3Mask  = 0xe00
)

// GetSmp3 ADC channel 3 sampling time selection
func (r *RegisterSmpr1Type) GetSmp3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr1FieldSmp3Mask) >> RegisterSmpr1FieldSmp3Shift)
}

// SetSmp3 ADC channel 3 sampling time selection
func (r *RegisterSmpr1Type) SetSmp3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr1FieldSmp3Mask)|(uint32(value)<<RegisterSmpr1FieldSmp3Shift))
}

const (
	RegisterSmpr1FieldSmp4Shift = 12
	RegisterSmpr1FieldSmp4Mask  = 0x7000
)

// GetSmp4 ADC channel 4 sampling time selection
func (r *RegisterSmpr1Type) GetSmp4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr1FieldSmp4Mask) >> RegisterSmpr1FieldSmp4Shift)
}

// SetSmp4 ADC channel 4 sampling time selection
func (r *RegisterSmpr1Type) SetSmp4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr1FieldSmp4Mask)|(uint32(value)<<RegisterSmpr1FieldSmp4Shift))
}

const (
	RegisterSmpr1FieldSmp5Shift = 15
	RegisterSmpr1FieldSmp5Mask  = 0x38000
)

// GetSmp5 ADC channel 5 sampling time selection
func (r *RegisterSmpr1Type) GetSmp5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr1FieldSmp5Mask) >> RegisterSmpr1FieldSmp5Shift)
}

// SetSmp5 ADC channel 5 sampling time selection
func (r *RegisterSmpr1Type) SetSmp5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr1FieldSmp5Mask)|(uint32(value)<<RegisterSmpr1FieldSmp5Shift))
}

const (
	RegisterSmpr1FieldSmp6Shift = 18
	RegisterSmpr1FieldSmp6Mask  = 0x1c0000
)

// GetSmp6 ADC channel 6 sampling time selection
func (r *RegisterSmpr1Type) GetSmp6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr1FieldSmp6Mask) >> RegisterSmpr1FieldSmp6Shift)
}

// SetSmp6 ADC channel 6 sampling time selection
func (r *RegisterSmpr1Type) SetSmp6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr1FieldSmp6Mask)|(uint32(value)<<RegisterSmpr1FieldSmp6Shift))
}

const (
	RegisterSmpr1FieldSmp7Shift = 21
	RegisterSmpr1FieldSmp7Mask  = 0xe00000
)

// GetSmp7 ADC channel 7 sampling time selection
func (r *RegisterSmpr1Type) GetSmp7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr1FieldSmp7Mask) >> RegisterSmpr1FieldSmp7Shift)
}

// SetSmp7 ADC channel 7 sampling time selection
func (r *RegisterSmpr1Type) SetSmp7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr1FieldSmp7Mask)|(uint32(value)<<RegisterSmpr1FieldSmp7Shift))
}

const (
	RegisterSmpr1FieldSmp8Shift = 24
	RegisterSmpr1FieldSmp8Mask  = 0x7000000
)

// GetSmp8 ADC channel 8 sampling time selection
func (r *RegisterSmpr1Type) GetSmp8() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr1FieldSmp8Mask) >> RegisterSmpr1FieldSmp8Shift)
}

// SetSmp8 ADC channel 8 sampling time selection
func (r *RegisterSmpr1Type) SetSmp8(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr1FieldSmp8Mask)|(uint32(value)<<RegisterSmpr1FieldSmp8Shift))
}

const (
	RegisterSmpr1FieldSmp9Shift = 27
	RegisterSmpr1FieldSmp9Mask  = 0x38000000
)

// GetSmp9 ADC channel 9 sampling time selection
func (r *RegisterSmpr1Type) GetSmp9() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr1FieldSmp9Mask) >> RegisterSmpr1FieldSmp9Shift)
}

// SetSmp9 ADC channel 9 sampling time selection
func (r *RegisterSmpr1Type) SetSmp9(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr1FieldSmp9Mask)|(uint32(value)<<RegisterSmpr1FieldSmp9Shift))
}

// RegisterSmpr2Type ADC sampling time register 2
type RegisterSmpr2Type uint32

func (r *RegisterSmpr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSmpr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSmpr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSmpr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSmpr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSmpr2FieldSmp10Shift = 0
	RegisterSmpr2FieldSmp10Mask  = 0x7
)

// GetSmp10 ADC channel 10 sampling time selection
func (r *RegisterSmpr2Type) GetSmp10() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr2FieldSmp10Mask) >> RegisterSmpr2FieldSmp10Shift)
}

// SetSmp10 ADC channel 10 sampling time selection
func (r *RegisterSmpr2Type) SetSmp10(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr2FieldSmp10Mask)|(uint32(value)<<RegisterSmpr2FieldSmp10Shift))
}

const (
	RegisterSmpr2FieldSmp11Shift = 3
	RegisterSmpr2FieldSmp11Mask  = 0x38
)

// GetSmp11 ADC channel 11 sampling time selection
func (r *RegisterSmpr2Type) GetSmp11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr2FieldSmp11Mask) >> RegisterSmpr2FieldSmp11Shift)
}

// SetSmp11 ADC channel 11 sampling time selection
func (r *RegisterSmpr2Type) SetSmp11(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr2FieldSmp11Mask)|(uint32(value)<<RegisterSmpr2FieldSmp11Shift))
}

const (
	RegisterSmpr2FieldSmp12Shift = 6
	RegisterSmpr2FieldSmp12Mask  = 0x1c0
)

// GetSmp12 ADC channel 12 sampling time selection
func (r *RegisterSmpr2Type) GetSmp12() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr2FieldSmp12Mask) >> RegisterSmpr2FieldSmp12Shift)
}

// SetSmp12 ADC channel 12 sampling time selection
func (r *RegisterSmpr2Type) SetSmp12(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr2FieldSmp12Mask)|(uint32(value)<<RegisterSmpr2FieldSmp12Shift))
}

const (
	RegisterSmpr2FieldSmp13Shift = 9
	RegisterSmpr2FieldSmp13Mask  = 0xe00
)

// GetSmp13 ADC channel 13 sampling time selection
func (r *RegisterSmpr2Type) GetSmp13() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr2FieldSmp13Mask) >> RegisterSmpr2FieldSmp13Shift)
}

// SetSmp13 ADC channel 13 sampling time selection
func (r *RegisterSmpr2Type) SetSmp13(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr2FieldSmp13Mask)|(uint32(value)<<RegisterSmpr2FieldSmp13Shift))
}

const (
	RegisterSmpr2FieldSmp14Shift = 12
	RegisterSmpr2FieldSmp14Mask  = 0x7000
)

// GetSmp14 ADC channel 14 sampling time selection
func (r *RegisterSmpr2Type) GetSmp14() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr2FieldSmp14Mask) >> RegisterSmpr2FieldSmp14Shift)
}

// SetSmp14 ADC channel 14 sampling time selection
func (r *RegisterSmpr2Type) SetSmp14(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr2FieldSmp14Mask)|(uint32(value)<<RegisterSmpr2FieldSmp14Shift))
}

const (
	RegisterSmpr2FieldSmp15Shift = 15
	RegisterSmpr2FieldSmp15Mask  = 0x38000
)

// GetSmp15 ADC channel 15 sampling time selection
func (r *RegisterSmpr2Type) GetSmp15() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr2FieldSmp15Mask) >> RegisterSmpr2FieldSmp15Shift)
}

// SetSmp15 ADC channel 15 sampling time selection
func (r *RegisterSmpr2Type) SetSmp15(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr2FieldSmp15Mask)|(uint32(value)<<RegisterSmpr2FieldSmp15Shift))
}

const (
	RegisterSmpr2FieldSmp16Shift = 18
	RegisterSmpr2FieldSmp16Mask  = 0x1c0000
)

// GetSmp16 ADC channel 16 sampling time selection
func (r *RegisterSmpr2Type) GetSmp16() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr2FieldSmp16Mask) >> RegisterSmpr2FieldSmp16Shift)
}

// SetSmp16 ADC channel 16 sampling time selection
func (r *RegisterSmpr2Type) SetSmp16(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr2FieldSmp16Mask)|(uint32(value)<<RegisterSmpr2FieldSmp16Shift))
}

const (
	RegisterSmpr2FieldSmp17Shift = 21
	RegisterSmpr2FieldSmp17Mask  = 0xe00000
)

// GetSmp17 ADC channel 17 sampling time selection
func (r *RegisterSmpr2Type) GetSmp17() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr2FieldSmp17Mask) >> RegisterSmpr2FieldSmp17Shift)
}

// SetSmp17 ADC channel 17 sampling time selection
func (r *RegisterSmpr2Type) SetSmp17(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr2FieldSmp17Mask)|(uint32(value)<<RegisterSmpr2FieldSmp17Shift))
}

const (
	RegisterSmpr2FieldSmp18Shift = 24
	RegisterSmpr2FieldSmp18Mask  = 0x7000000
)

// GetSmp18 ADC channel 18 sampling time selection
func (r *RegisterSmpr2Type) GetSmp18() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr2FieldSmp18Mask) >> RegisterSmpr2FieldSmp18Shift)
}

// SetSmp18 ADC channel 18 sampling time selection
func (r *RegisterSmpr2Type) SetSmp18(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr2FieldSmp18Mask)|(uint32(value)<<RegisterSmpr2FieldSmp18Shift))
}

const (
	RegisterSmpr2FieldSmp19Shift = 27
	RegisterSmpr2FieldSmp19Mask  = 0x38000000
)

// GetSmp19 ADC channel 18 sampling time selection
func (r *RegisterSmpr2Type) GetSmp19() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSmpr2FieldSmp19Mask) >> RegisterSmpr2FieldSmp19Shift)
}

// SetSmp19 ADC channel 18 sampling time selection
func (r *RegisterSmpr2Type) SetSmp19(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSmpr2FieldSmp19Mask)|(uint32(value)<<RegisterSmpr2FieldSmp19Shift))
}

// RegisterPcselType ADC pre channel selection register
type RegisterPcselType uint32

func (r *RegisterPcselType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPcselType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPcselType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPcselType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPcselType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPcselFieldPcselShift = 0
	RegisterPcselFieldPcselMask  = 0xfffff
)

// GetPcsel Channel x (VINP[i]) pre selection
func (r *RegisterPcselType) GetPcsel() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterPcselFieldPcselMask) >> RegisterPcselFieldPcselShift)
}

// SetPcsel Channel x (VINP[i]) pre selection
func (r *RegisterPcselType) SetPcsel(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPcselFieldPcselMask)|(uint32(value)<<RegisterPcselFieldPcselShift))
}

// RegisterLtr1Type ADC analog watchdog 1 threshold register
type RegisterLtr1Type uint32

func (r *RegisterLtr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterLtr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterLtr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterLtr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterLtr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterLtr1FieldLtr1Shift = 0
	RegisterLtr1FieldLtr1Mask  = 0x3ffffff
)

// GetLtr1 ADC analog watchdog 1 threshold low
func (r *RegisterLtr1Type) GetLtr1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterLtr1FieldLtr1Mask) >> RegisterLtr1FieldLtr1Shift)
}

// SetLtr1 ADC analog watchdog 1 threshold low
func (r *RegisterLtr1Type) SetLtr1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterLtr1FieldLtr1Mask)|(uint32(value)<<RegisterLtr1FieldLtr1Shift))
}

// RegisterLhtr1Type ADC analog watchdog 2 threshold register
type RegisterLhtr1Type uint32

func (r *RegisterLhtr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterLhtr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterLhtr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterLhtr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterLhtr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterLhtr1FieldLhtr1Shift = 0
	RegisterLhtr1FieldLhtr1Mask  = 0x3ffffff
)

// GetLhtr1 ADC analog watchdog 2 threshold low
func (r *RegisterLhtr1Type) GetLhtr1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterLhtr1FieldLhtr1Mask) >> RegisterLhtr1FieldLhtr1Shift)
}

// SetLhtr1 ADC analog watchdog 2 threshold low
func (r *RegisterLhtr1Type) SetLhtr1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterLhtr1FieldLhtr1Mask)|(uint32(value)<<RegisterLhtr1FieldLhtr1Shift))
}

// RegisterSqr1Type ADC group regular sequencer ranks register 1
type RegisterSqr1Type uint32

func (r *RegisterSqr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSqr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSqr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSqr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSqr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSqr1FieldL3Shift = 0
	RegisterSqr1FieldL3Mask  = 0xf
)

// GetL3 L3
func (r *RegisterSqr1Type) GetL3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr1FieldL3Mask) >> RegisterSqr1FieldL3Shift)
}

// SetL3 L3
func (r *RegisterSqr1Type) SetL3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr1FieldL3Mask)|(uint32(value)<<RegisterSqr1FieldL3Shift))
}

const (
	RegisterSqr1FieldSq1Shift = 6
	RegisterSqr1FieldSq1Mask  = 0x7c0
)

// GetSq1 ADC group regular sequencer rank 1
func (r *RegisterSqr1Type) GetSq1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr1FieldSq1Mask) >> RegisterSqr1FieldSq1Shift)
}

// SetSq1 ADC group regular sequencer rank 1
func (r *RegisterSqr1Type) SetSq1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr1FieldSq1Mask)|(uint32(value)<<RegisterSqr1FieldSq1Shift))
}

const (
	RegisterSqr1FieldSq2Shift = 12
	RegisterSqr1FieldSq2Mask  = 0x1f000
)

// GetSq2 ADC group regular sequencer rank 2
func (r *RegisterSqr1Type) GetSq2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr1FieldSq2Mask) >> RegisterSqr1FieldSq2Shift)
}

// SetSq2 ADC group regular sequencer rank 2
func (r *RegisterSqr1Type) SetSq2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr1FieldSq2Mask)|(uint32(value)<<RegisterSqr1FieldSq2Shift))
}

const (
	RegisterSqr1FieldSq3Shift = 18
	RegisterSqr1FieldSq3Mask  = 0x7c0000
)

// GetSq3 ADC group regular sequencer rank 3
func (r *RegisterSqr1Type) GetSq3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr1FieldSq3Mask) >> RegisterSqr1FieldSq3Shift)
}

// SetSq3 ADC group regular sequencer rank 3
func (r *RegisterSqr1Type) SetSq3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr1FieldSq3Mask)|(uint32(value)<<RegisterSqr1FieldSq3Shift))
}

const (
	RegisterSqr1FieldSq4Shift = 24
	RegisterSqr1FieldSq4Mask  = 0x1f000000
)

// GetSq4 ADC group regular sequencer rank 4
func (r *RegisterSqr1Type) GetSq4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr1FieldSq4Mask) >> RegisterSqr1FieldSq4Shift)
}

// SetSq4 ADC group regular sequencer rank 4
func (r *RegisterSqr1Type) SetSq4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr1FieldSq4Mask)|(uint32(value)<<RegisterSqr1FieldSq4Shift))
}

// RegisterSqr2Type ADC group regular sequencer ranks register 2
type RegisterSqr2Type uint32

func (r *RegisterSqr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSqr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSqr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSqr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSqr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSqr2FieldSq5Shift = 0
	RegisterSqr2FieldSq5Mask  = 0x1f
)

// GetSq5 ADC group regular sequencer rank 5
func (r *RegisterSqr2Type) GetSq5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr2FieldSq5Mask) >> RegisterSqr2FieldSq5Shift)
}

// SetSq5 ADC group regular sequencer rank 5
func (r *RegisterSqr2Type) SetSq5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr2FieldSq5Mask)|(uint32(value)<<RegisterSqr2FieldSq5Shift))
}

const (
	RegisterSqr2FieldSq6Shift = 6
	RegisterSqr2FieldSq6Mask  = 0x7c0
)

// GetSq6 ADC group regular sequencer rank 6
func (r *RegisterSqr2Type) GetSq6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr2FieldSq6Mask) >> RegisterSqr2FieldSq6Shift)
}

// SetSq6 ADC group regular sequencer rank 6
func (r *RegisterSqr2Type) SetSq6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr2FieldSq6Mask)|(uint32(value)<<RegisterSqr2FieldSq6Shift))
}

const (
	RegisterSqr2FieldSq7Shift = 12
	RegisterSqr2FieldSq7Mask  = 0x1f000
)

// GetSq7 ADC group regular sequencer rank 7
func (r *RegisterSqr2Type) GetSq7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr2FieldSq7Mask) >> RegisterSqr2FieldSq7Shift)
}

// SetSq7 ADC group regular sequencer rank 7
func (r *RegisterSqr2Type) SetSq7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr2FieldSq7Mask)|(uint32(value)<<RegisterSqr2FieldSq7Shift))
}

const (
	RegisterSqr2FieldSq8Shift = 18
	RegisterSqr2FieldSq8Mask  = 0x7c0000
)

// GetSq8 ADC group regular sequencer rank 8
func (r *RegisterSqr2Type) GetSq8() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr2FieldSq8Mask) >> RegisterSqr2FieldSq8Shift)
}

// SetSq8 ADC group regular sequencer rank 8
func (r *RegisterSqr2Type) SetSq8(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr2FieldSq8Mask)|(uint32(value)<<RegisterSqr2FieldSq8Shift))
}

const (
	RegisterSqr2FieldSq9Shift = 24
	RegisterSqr2FieldSq9Mask  = 0x1f000000
)

// GetSq9 ADC group regular sequencer rank 9
func (r *RegisterSqr2Type) GetSq9() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr2FieldSq9Mask) >> RegisterSqr2FieldSq9Shift)
}

// SetSq9 ADC group regular sequencer rank 9
func (r *RegisterSqr2Type) SetSq9(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr2FieldSq9Mask)|(uint32(value)<<RegisterSqr2FieldSq9Shift))
}

// RegisterSqr3Type ADC group regular sequencer ranks register 3
type RegisterSqr3Type uint32

func (r *RegisterSqr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSqr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSqr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSqr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSqr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSqr3FieldSq10Shift = 0
	RegisterSqr3FieldSq10Mask  = 0x1f
)

// GetSq10 ADC group regular sequencer rank 10
func (r *RegisterSqr3Type) GetSq10() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr3FieldSq10Mask) >> RegisterSqr3FieldSq10Shift)
}

// SetSq10 ADC group regular sequencer rank 10
func (r *RegisterSqr3Type) SetSq10(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr3FieldSq10Mask)|(uint32(value)<<RegisterSqr3FieldSq10Shift))
}

const (
	RegisterSqr3FieldSq11Shift = 6
	RegisterSqr3FieldSq11Mask  = 0x7c0
)

// GetSq11 ADC group regular sequencer rank 11
func (r *RegisterSqr3Type) GetSq11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr3FieldSq11Mask) >> RegisterSqr3FieldSq11Shift)
}

// SetSq11 ADC group regular sequencer rank 11
func (r *RegisterSqr3Type) SetSq11(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr3FieldSq11Mask)|(uint32(value)<<RegisterSqr3FieldSq11Shift))
}

const (
	RegisterSqr3FieldSq12Shift = 12
	RegisterSqr3FieldSq12Mask  = 0x1f000
)

// GetSq12 ADC group regular sequencer rank 12
func (r *RegisterSqr3Type) GetSq12() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr3FieldSq12Mask) >> RegisterSqr3FieldSq12Shift)
}

// SetSq12 ADC group regular sequencer rank 12
func (r *RegisterSqr3Type) SetSq12(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr3FieldSq12Mask)|(uint32(value)<<RegisterSqr3FieldSq12Shift))
}

const (
	RegisterSqr3FieldSq13Shift = 18
	RegisterSqr3FieldSq13Mask  = 0x7c0000
)

// GetSq13 ADC group regular sequencer rank 13
func (r *RegisterSqr3Type) GetSq13() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr3FieldSq13Mask) >> RegisterSqr3FieldSq13Shift)
}

// SetSq13 ADC group regular sequencer rank 13
func (r *RegisterSqr3Type) SetSq13(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr3FieldSq13Mask)|(uint32(value)<<RegisterSqr3FieldSq13Shift))
}

const (
	RegisterSqr3FieldSq14Shift = 24
	RegisterSqr3FieldSq14Mask  = 0x1f000000
)

// GetSq14 ADC group regular sequencer rank 14
func (r *RegisterSqr3Type) GetSq14() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr3FieldSq14Mask) >> RegisterSqr3FieldSq14Shift)
}

// SetSq14 ADC group regular sequencer rank 14
func (r *RegisterSqr3Type) SetSq14(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr3FieldSq14Mask)|(uint32(value)<<RegisterSqr3FieldSq14Shift))
}

// RegisterSqr4Type ADC group regular sequencer ranks register 4
type RegisterSqr4Type uint32

func (r *RegisterSqr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSqr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSqr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSqr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSqr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSqr4FieldSq15Shift = 0
	RegisterSqr4FieldSq15Mask  = 0x1f
)

// GetSq15 ADC group regular sequencer rank 15
func (r *RegisterSqr4Type) GetSq15() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr4FieldSq15Mask) >> RegisterSqr4FieldSq15Shift)
}

// SetSq15 ADC group regular sequencer rank 15
func (r *RegisterSqr4Type) SetSq15(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr4FieldSq15Mask)|(uint32(value)<<RegisterSqr4FieldSq15Shift))
}

const (
	RegisterSqr4FieldSq16Shift = 6
	RegisterSqr4FieldSq16Mask  = 0x7c0
)

// GetSq16 ADC group regular sequencer rank 16
func (r *RegisterSqr4Type) GetSq16() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSqr4FieldSq16Mask) >> RegisterSqr4FieldSq16Shift)
}

// SetSq16 ADC group regular sequencer rank 16
func (r *RegisterSqr4Type) SetSq16(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSqr4FieldSq16Mask)|(uint32(value)<<RegisterSqr4FieldSq16Shift))
}

// RegisterDrType ADC group regular conversion data register
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
	RegisterDrFieldRdataShift = 0
	RegisterDrFieldRdataMask  = 0xffff
)

// GetRdata ADC group regular conversion data
func (r *RegisterDrType) GetRdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldRdataMask) >> RegisterDrFieldRdataShift)
}

// SetRdata ADC group regular conversion data
func (r *RegisterDrType) SetRdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldRdataMask)|(uint32(value)<<RegisterDrFieldRdataShift))
}

// RegisterJsqrType ADC group injected sequencer register
type RegisterJsqrType uint32

func (r *RegisterJsqrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterJsqrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterJsqrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterJsqrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterJsqrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterJsqrFieldJlShift = 0
	RegisterJsqrFieldJlMask  = 0x3
)

// GetJl ADC group injected sequencer scan length
func (r *RegisterJsqrType) GetJl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterJsqrFieldJlMask) >> RegisterJsqrFieldJlShift)
}

// SetJl ADC group injected sequencer scan length
func (r *RegisterJsqrType) SetJl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterJsqrFieldJlMask)|(uint32(value)<<RegisterJsqrFieldJlShift))
}

const (
	RegisterJsqrFieldJextselShift = 2
	RegisterJsqrFieldJextselMask  = 0x7c
)

// GetJextsel ADC group injected external trigger source
func (r *RegisterJsqrType) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterJsqrFieldJextselMask) >> RegisterJsqrFieldJextselShift)
}

// SetJextsel ADC group injected external trigger source
func (r *RegisterJsqrType) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterJsqrFieldJextselMask)|(uint32(value)<<RegisterJsqrFieldJextselShift))
}

const (
	RegisterJsqrFieldJextenShift = 7
	RegisterJsqrFieldJextenMask  = 0x180
)

// GetJexten ADC group injected external trigger polarity
func (r *RegisterJsqrType) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterJsqrFieldJextenMask) >> RegisterJsqrFieldJextenShift)
}

// SetJexten ADC group injected external trigger polarity
func (r *RegisterJsqrType) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterJsqrFieldJextenMask)|(uint32(value)<<RegisterJsqrFieldJextenShift))
}

const (
	RegisterJsqrFieldJsq1Shift = 9
	RegisterJsqrFieldJsq1Mask  = 0x3e00
)

// GetJsq1 ADC group injected sequencer rank 1
func (r *RegisterJsqrType) GetJsq1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterJsqrFieldJsq1Mask) >> RegisterJsqrFieldJsq1Shift)
}

// SetJsq1 ADC group injected sequencer rank 1
func (r *RegisterJsqrType) SetJsq1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterJsqrFieldJsq1Mask)|(uint32(value)<<RegisterJsqrFieldJsq1Shift))
}

const (
	RegisterJsqrFieldJsq2Shift = 15
	RegisterJsqrFieldJsq2Mask  = 0xf8000
)

// GetJsq2 ADC group injected sequencer rank 2
func (r *RegisterJsqrType) GetJsq2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterJsqrFieldJsq2Mask) >> RegisterJsqrFieldJsq2Shift)
}

// SetJsq2 ADC group injected sequencer rank 2
func (r *RegisterJsqrType) SetJsq2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterJsqrFieldJsq2Mask)|(uint32(value)<<RegisterJsqrFieldJsq2Shift))
}

const (
	RegisterJsqrFieldJsq3Shift = 21
	RegisterJsqrFieldJsq3Mask  = 0x3e00000
)

// GetJsq3 ADC group injected sequencer rank 3
func (r *RegisterJsqrType) GetJsq3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterJsqrFieldJsq3Mask) >> RegisterJsqrFieldJsq3Shift)
}

// SetJsq3 ADC group injected sequencer rank 3
func (r *RegisterJsqrType) SetJsq3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterJsqrFieldJsq3Mask)|(uint32(value)<<RegisterJsqrFieldJsq3Shift))
}

const (
	RegisterJsqrFieldJsq4Shift = 27
	RegisterJsqrFieldJsq4Mask  = 0xf8000000
)

// GetJsq4 ADC group injected sequencer rank 4
func (r *RegisterJsqrType) GetJsq4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterJsqrFieldJsq4Mask) >> RegisterJsqrFieldJsq4Shift)
}

// SetJsq4 ADC group injected sequencer rank 4
func (r *RegisterJsqrType) SetJsq4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterJsqrFieldJsq4Mask)|(uint32(value)<<RegisterJsqrFieldJsq4Shift))
}

// RegisterOfr1Type ADC offset number 1 register
type RegisterOfr1Type uint32

func (r *RegisterOfr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOfr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOfr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOfr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOfr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOfr1FieldOffset1Shift = 0
	RegisterOfr1FieldOffset1Mask  = 0x3ffffff
)

// GetOffset1 ADC offset number 1 offset level
func (r *RegisterOfr1Type) GetOffset1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOfr1FieldOffset1Mask) >> RegisterOfr1FieldOffset1Shift)
}

// SetOffset1 ADC offset number 1 offset level
func (r *RegisterOfr1Type) SetOffset1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOfr1FieldOffset1Mask)|(uint32(value)<<RegisterOfr1FieldOffset1Shift))
}

const (
	RegisterOfr1FieldOffset1chShift = 26
	RegisterOfr1FieldOffset1chMask  = 0x7c000000
)

// GetOffset1ch ADC offset number 1 channel selection
func (r *RegisterOfr1Type) GetOffset1ch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOfr1FieldOffset1chMask) >> RegisterOfr1FieldOffset1chShift)
}

// SetOffset1ch ADC offset number 1 channel selection
func (r *RegisterOfr1Type) SetOffset1ch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOfr1FieldOffset1chMask)|(uint32(value)<<RegisterOfr1FieldOffset1chShift))
}

const (
	RegisterOfr1FieldSsateShift = 31
	RegisterOfr1FieldSsateMask  = 0x80000000
)

// GetSsate ADC offset number 1 enable
func (r *RegisterOfr1Type) GetSsate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOfr1FieldSsateMask) != 0
}

// SetSsate ADC offset number 1 enable
func (r *RegisterOfr1Type) SetSsate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOfr1FieldSsateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOfr1FieldSsateMask)
	}
}

// RegisterOfr2Type ADC offset number 2 register
type RegisterOfr2Type uint32

func (r *RegisterOfr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOfr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOfr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOfr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOfr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOfr2FieldOffset1Shift = 0
	RegisterOfr2FieldOffset1Mask  = 0x3ffffff
)

// GetOffset1 ADC offset number 1 offset level
func (r *RegisterOfr2Type) GetOffset1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOfr2FieldOffset1Mask) >> RegisterOfr2FieldOffset1Shift)
}

// SetOffset1 ADC offset number 1 offset level
func (r *RegisterOfr2Type) SetOffset1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOfr2FieldOffset1Mask)|(uint32(value)<<RegisterOfr2FieldOffset1Shift))
}

const (
	RegisterOfr2FieldOffset1chShift = 26
	RegisterOfr2FieldOffset1chMask  = 0x7c000000
)

// GetOffset1ch ADC offset number 1 channel selection
func (r *RegisterOfr2Type) GetOffset1ch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOfr2FieldOffset1chMask) >> RegisterOfr2FieldOffset1chShift)
}

// SetOffset1ch ADC offset number 1 channel selection
func (r *RegisterOfr2Type) SetOffset1ch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOfr2FieldOffset1chMask)|(uint32(value)<<RegisterOfr2FieldOffset1chShift))
}

const (
	RegisterOfr2FieldSsateShift = 31
	RegisterOfr2FieldSsateMask  = 0x80000000
)

// GetSsate ADC offset number 1 enable
func (r *RegisterOfr2Type) GetSsate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOfr2FieldSsateMask) != 0
}

// SetSsate ADC offset number 1 enable
func (r *RegisterOfr2Type) SetSsate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOfr2FieldSsateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOfr2FieldSsateMask)
	}
}

// RegisterOfr3Type ADC offset number 3 register
type RegisterOfr3Type uint32

func (r *RegisterOfr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOfr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOfr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOfr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOfr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOfr3FieldOffset1Shift = 0
	RegisterOfr3FieldOffset1Mask  = 0x3ffffff
)

// GetOffset1 ADC offset number 1 offset level
func (r *RegisterOfr3Type) GetOffset1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOfr3FieldOffset1Mask) >> RegisterOfr3FieldOffset1Shift)
}

// SetOffset1 ADC offset number 1 offset level
func (r *RegisterOfr3Type) SetOffset1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOfr3FieldOffset1Mask)|(uint32(value)<<RegisterOfr3FieldOffset1Shift))
}

const (
	RegisterOfr3FieldOffset1chShift = 26
	RegisterOfr3FieldOffset1chMask  = 0x7c000000
)

// GetOffset1ch ADC offset number 1 channel selection
func (r *RegisterOfr3Type) GetOffset1ch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOfr3FieldOffset1chMask) >> RegisterOfr3FieldOffset1chShift)
}

// SetOffset1ch ADC offset number 1 channel selection
func (r *RegisterOfr3Type) SetOffset1ch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOfr3FieldOffset1chMask)|(uint32(value)<<RegisterOfr3FieldOffset1chShift))
}

const (
	RegisterOfr3FieldSsateShift = 31
	RegisterOfr3FieldSsateMask  = 0x80000000
)

// GetSsate ADC offset number 1 enable
func (r *RegisterOfr3Type) GetSsate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOfr3FieldSsateMask) != 0
}

// SetSsate ADC offset number 1 enable
func (r *RegisterOfr3Type) SetSsate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOfr3FieldSsateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOfr3FieldSsateMask)
	}
}

// RegisterOfr4Type ADC offset number 4 register
type RegisterOfr4Type uint32

func (r *RegisterOfr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOfr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOfr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOfr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOfr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOfr4FieldOffset1Shift = 0
	RegisterOfr4FieldOffset1Mask  = 0x3ffffff
)

// GetOffset1 ADC offset number 1 offset level
func (r *RegisterOfr4Type) GetOffset1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOfr4FieldOffset1Mask) >> RegisterOfr4FieldOffset1Shift)
}

// SetOffset1 ADC offset number 1 offset level
func (r *RegisterOfr4Type) SetOffset1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOfr4FieldOffset1Mask)|(uint32(value)<<RegisterOfr4FieldOffset1Shift))
}

const (
	RegisterOfr4FieldOffset1chShift = 26
	RegisterOfr4FieldOffset1chMask  = 0x7c000000
)

// GetOffset1ch ADC offset number 1 channel selection
func (r *RegisterOfr4Type) GetOffset1ch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOfr4FieldOffset1chMask) >> RegisterOfr4FieldOffset1chShift)
}

// SetOffset1ch ADC offset number 1 channel selection
func (r *RegisterOfr4Type) SetOffset1ch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOfr4FieldOffset1chMask)|(uint32(value)<<RegisterOfr4FieldOffset1chShift))
}

const (
	RegisterOfr4FieldSsateShift = 31
	RegisterOfr4FieldSsateMask  = 0x80000000
)

// GetSsate ADC offset number 1 enable
func (r *RegisterOfr4Type) GetSsate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOfr4FieldSsateMask) != 0
}

// SetSsate ADC offset number 1 enable
func (r *RegisterOfr4Type) SetSsate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOfr4FieldSsateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOfr4FieldSsateMask)
	}
}

// RegisterJdr1Type ADC group injected sequencer rank 1 register
type RegisterJdr1Type uint32

func (r *RegisterJdr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterJdr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterJdr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterJdr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterJdr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterJdr1FieldJdata1Shift = 0
	RegisterJdr1FieldJdata1Mask  = 0xffffffff
)

// GetJdata1 ADC group injected sequencer rank 1 conversion data
func (r *RegisterJdr1Type) GetJdata1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterJdr1FieldJdata1Mask) >> RegisterJdr1FieldJdata1Shift)
}

// SetJdata1 ADC group injected sequencer rank 1 conversion data
func (r *RegisterJdr1Type) SetJdata1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterJdr1FieldJdata1Mask)|(uint32(value)<<RegisterJdr1FieldJdata1Shift))
}

// RegisterJdr2Type ADC group injected sequencer rank 2 register
type RegisterJdr2Type uint32

func (r *RegisterJdr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterJdr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterJdr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterJdr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterJdr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterJdr2FieldJdata2Shift = 0
	RegisterJdr2FieldJdata2Mask  = 0xffffffff
)

// GetJdata2 ADC group injected sequencer rank 2 conversion data
func (r *RegisterJdr2Type) GetJdata2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterJdr2FieldJdata2Mask) >> RegisterJdr2FieldJdata2Shift)
}

// SetJdata2 ADC group injected sequencer rank 2 conversion data
func (r *RegisterJdr2Type) SetJdata2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterJdr2FieldJdata2Mask)|(uint32(value)<<RegisterJdr2FieldJdata2Shift))
}

// RegisterJdr3Type ADC group injected sequencer rank 3 register
type RegisterJdr3Type uint32

func (r *RegisterJdr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterJdr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterJdr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterJdr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterJdr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterJdr3FieldJdata3Shift = 0
	RegisterJdr3FieldJdata3Mask  = 0xffffffff
)

// GetJdata3 ADC group injected sequencer rank 3 conversion data
func (r *RegisterJdr3Type) GetJdata3() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterJdr3FieldJdata3Mask) >> RegisterJdr3FieldJdata3Shift)
}

// SetJdata3 ADC group injected sequencer rank 3 conversion data
func (r *RegisterJdr3Type) SetJdata3(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterJdr3FieldJdata3Mask)|(uint32(value)<<RegisterJdr3FieldJdata3Shift))
}

// RegisterJdr4Type ADC group injected sequencer rank 4 register
type RegisterJdr4Type uint32

func (r *RegisterJdr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterJdr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterJdr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterJdr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterJdr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterJdr4FieldJdata4Shift = 0
	RegisterJdr4FieldJdata4Mask  = 0xffffffff
)

// GetJdata4 ADC group injected sequencer rank 4 conversion data
func (r *RegisterJdr4Type) GetJdata4() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterJdr4FieldJdata4Mask) >> RegisterJdr4FieldJdata4Shift)
}

// SetJdata4 ADC group injected sequencer rank 4 conversion data
func (r *RegisterJdr4Type) SetJdata4(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterJdr4FieldJdata4Mask)|(uint32(value)<<RegisterJdr4FieldJdata4Shift))
}

// RegisterAwd2crType ADC analog watchdog 2 configuration register
type RegisterAwd2crType uint32

func (r *RegisterAwd2crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAwd2crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAwd2crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAwd2crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAwd2crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAwd2crFieldAwd2chShift = 0
	RegisterAwd2crFieldAwd2chMask  = 0xfffff
)

// GetAwd2ch ADC analog watchdog 2 monitored channel selection
func (r *RegisterAwd2crType) GetAwd2ch() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterAwd2crFieldAwd2chMask) >> RegisterAwd2crFieldAwd2chShift)
}

// SetAwd2ch ADC analog watchdog 2 monitored channel selection
func (r *RegisterAwd2crType) SetAwd2ch(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAwd2crFieldAwd2chMask)|(uint32(value)<<RegisterAwd2crFieldAwd2chShift))
}

// RegisterAwd3crType ADC analog watchdog 3 configuration register
type RegisterAwd3crType uint32

func (r *RegisterAwd3crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAwd3crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAwd3crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAwd3crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAwd3crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAwd3crFieldAwd3chShift = 1
	RegisterAwd3crFieldAwd3chMask  = 0x1ffffe
)

// GetAwd3ch ADC analog watchdog 3 monitored channel selection
func (r *RegisterAwd3crType) GetAwd3ch() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterAwd3crFieldAwd3chMask) >> RegisterAwd3crFieldAwd3chShift)
}

// SetAwd3ch ADC analog watchdog 3 monitored channel selection
func (r *RegisterAwd3crType) SetAwd3ch(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAwd3crFieldAwd3chMask)|(uint32(value)<<RegisterAwd3crFieldAwd3chShift))
}

// RegisterLtr2Type ADC watchdog lower threshold register 2
type RegisterLtr2Type uint32

func (r *RegisterLtr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterLtr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterLtr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterLtr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterLtr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterLtr2FieldLtr2Shift = 0
	RegisterLtr2FieldLtr2Mask  = 0x3ffffff
)

// GetLtr2 Analog watchdog 2 lower threshold
func (r *RegisterLtr2Type) GetLtr2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterLtr2FieldLtr2Mask) >> RegisterLtr2FieldLtr2Shift)
}

// SetLtr2 Analog watchdog 2 lower threshold
func (r *RegisterLtr2Type) SetLtr2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterLtr2FieldLtr2Mask)|(uint32(value)<<RegisterLtr2FieldLtr2Shift))
}

// RegisterHtr2Type ADC watchdog higher threshold register 2
type RegisterHtr2Type uint32

func (r *RegisterHtr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHtr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHtr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHtr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHtr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHtr2FieldHtr2Shift = 0
	RegisterHtr2FieldHtr2Mask  = 0x3ffffff
)

// GetHtr2 Analog watchdog 2 higher threshold
func (r *RegisterHtr2Type) GetHtr2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHtr2FieldHtr2Mask) >> RegisterHtr2FieldHtr2Shift)
}

// SetHtr2 Analog watchdog 2 higher threshold
func (r *RegisterHtr2Type) SetHtr2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHtr2FieldHtr2Mask)|(uint32(value)<<RegisterHtr2FieldHtr2Shift))
}

// RegisterLtr3Type ADC watchdog lower threshold register 3
type RegisterLtr3Type uint32

func (r *RegisterLtr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterLtr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterLtr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterLtr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterLtr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterLtr3FieldLtr3Shift = 0
	RegisterLtr3FieldLtr3Mask  = 0x3ffffff
)

// GetLtr3 Analog watchdog 3 lower threshold
func (r *RegisterLtr3Type) GetLtr3() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterLtr3FieldLtr3Mask) >> RegisterLtr3FieldLtr3Shift)
}

// SetLtr3 Analog watchdog 3 lower threshold
func (r *RegisterLtr3Type) SetLtr3(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterLtr3FieldLtr3Mask)|(uint32(value)<<RegisterLtr3FieldLtr3Shift))
}

// RegisterHtr3Type ADC watchdog higher threshold register 3
type RegisterHtr3Type uint32

func (r *RegisterHtr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHtr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHtr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHtr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHtr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHtr3FieldHtr3Shift = 0
	RegisterHtr3FieldHtr3Mask  = 0x3ffffff
)

// GetHtr3 Analog watchdog 3 higher threshold
func (r *RegisterHtr3Type) GetHtr3() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHtr3FieldHtr3Mask) >> RegisterHtr3FieldHtr3Shift)
}

// SetHtr3 Analog watchdog 3 higher threshold
func (r *RegisterHtr3Type) SetHtr3(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHtr3FieldHtr3Mask)|(uint32(value)<<RegisterHtr3FieldHtr3Shift))
}

// RegisterDifselType ADC channel differential or single-ended mode selection register
type RegisterDifselType uint32

func (r *RegisterDifselType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDifselType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDifselType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDifselType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDifselType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDifselFieldDifselShift = 0
	RegisterDifselFieldDifselMask  = 0xfffff
)

// GetDifsel ADC channel differential or single-ended mode for channel
func (r *RegisterDifselType) GetDifsel() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDifselFieldDifselMask) >> RegisterDifselFieldDifselShift)
}

// SetDifsel ADC channel differential or single-ended mode for channel
func (r *RegisterDifselType) SetDifsel(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDifselFieldDifselMask)|(uint32(value)<<RegisterDifselFieldDifselShift))
}

// RegisterCalfactType ADC calibration factors register
type RegisterCalfactType uint32

func (r *RegisterCalfactType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCalfactType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCalfactType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCalfactType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCalfactType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCalfactFieldCalfactsShift = 0
	RegisterCalfactFieldCalfactsMask  = 0x7ff
)

// GetCalfacts ADC calibration factor in single-ended mode
func (r *RegisterCalfactType) GetCalfacts() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCalfactFieldCalfactsMask) >> RegisterCalfactFieldCalfactsShift)
}

// SetCalfacts ADC calibration factor in single-ended mode
func (r *RegisterCalfactType) SetCalfacts(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCalfactFieldCalfactsMask)|(uint32(value)<<RegisterCalfactFieldCalfactsShift))
}

const (
	RegisterCalfactFieldCalfactdShift = 16
	RegisterCalfactFieldCalfactdMask  = 0x7ff0000
)

// GetCalfactd ADC calibration factor in differential mode
func (r *RegisterCalfactType) GetCalfactd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCalfactFieldCalfactdMask) >> RegisterCalfactFieldCalfactdShift)
}

// SetCalfactd ADC calibration factor in differential mode
func (r *RegisterCalfactType) SetCalfactd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCalfactFieldCalfactdMask)|(uint32(value)<<RegisterCalfactFieldCalfactdShift))
}

// RegisterCalfact2Type ADC Calibration Factor register 2
type RegisterCalfact2Type uint32

func (r *RegisterCalfact2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCalfact2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCalfact2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCalfact2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCalfact2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCalfact2FieldLincalfactShift = 0
	RegisterCalfact2FieldLincalfactMask  = 0x3fffffff
)

// GetLincalfact Linearity Calibration Factor
func (r *RegisterCalfact2Type) GetLincalfact() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCalfact2FieldLincalfactMask) >> RegisterCalfact2FieldLincalfactShift)
}

// SetLincalfact Linearity Calibration Factor
func (r *RegisterCalfact2Type) SetLincalfact(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCalfact2FieldLincalfactMask)|(uint32(value)<<RegisterCalfact2FieldLincalfactShift))
}
