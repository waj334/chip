//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package can_ccu

import (
	"unsafe"
	"volatile"
)

var (
	Can_ccu = (*_can_ccu)(unsafe.Pointer(uintptr(0x4000a800)))
)

type _can_ccu struct {
	Crel  registerCrelType
	Ccfg  registerCcfgType
	Cstat registerCstatType
	Cwd   registerCwdType
	Ir    registerIrType
	Ie    registerIeType
}

// registerCrelType Clock Calibration Unit Core Release Register
type registerCrelType uint32

const (
	RegisterCrelFieldDayShift = 0
	RegisterCrelFieldDayMask  = 0xff
)

// GetDay Time Stamp Day
func (r *registerCrelType) GetDay() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrelFieldDayMask) >> RegisterCrelFieldDayShift)
}

// SetDay Time Stamp Day
func (r *registerCrelType) SetDay(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrelFieldDayMask)|(uint32(value)<<RegisterCrelFieldDayShift))
}

const (
	RegisterCrelFieldMonShift = 8
	RegisterCrelFieldMonMask  = 0xff00
)

// GetMon Time Stamp Month
func (r *registerCrelType) GetMon() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrelFieldMonMask) >> RegisterCrelFieldMonShift)
}

// SetMon Time Stamp Month
func (r *registerCrelType) SetMon(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrelFieldMonMask)|(uint32(value)<<RegisterCrelFieldMonShift))
}

const (
	RegisterCrelFieldYearShift = 16
	RegisterCrelFieldYearMask  = 0xf0000
)

// GetYear Time Stamp Year
func (r *registerCrelType) GetYear() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrelFieldYearMask) >> RegisterCrelFieldYearShift)
}

// SetYear Time Stamp Year
func (r *registerCrelType) SetYear(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrelFieldYearMask)|(uint32(value)<<RegisterCrelFieldYearShift))
}

const (
	RegisterCrelFieldSubstepShift = 20
	RegisterCrelFieldSubstepMask  = 0xf00000
)

// GetSubstep Sub-step of Core Release
func (r *registerCrelType) GetSubstep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrelFieldSubstepMask) >> RegisterCrelFieldSubstepShift)
}

// SetSubstep Sub-step of Core Release
func (r *registerCrelType) SetSubstep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrelFieldSubstepMask)|(uint32(value)<<RegisterCrelFieldSubstepShift))
}

const (
	RegisterCrelFieldStepShift = 24
	RegisterCrelFieldStepMask  = 0xf000000
)

// GetStep Step of Core Release
func (r *registerCrelType) GetStep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrelFieldStepMask) >> RegisterCrelFieldStepShift)
}

// SetStep Step of Core Release
func (r *registerCrelType) SetStep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrelFieldStepMask)|(uint32(value)<<RegisterCrelFieldStepShift))
}

const (
	RegisterCrelFieldRelShift = 28
	RegisterCrelFieldRelMask  = 0xf0000000
)

// GetRel Core Release
func (r *registerCrelType) GetRel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrelFieldRelMask) >> RegisterCrelFieldRelShift)
}

// SetRel Core Release
func (r *registerCrelType) SetRel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrelFieldRelMask)|(uint32(value)<<RegisterCrelFieldRelShift))
}

// registerCcfgType Calibration Configuration Register
type registerCcfgType uint32

const (
	RegisterCcfgFieldTqbtShift = 0
	RegisterCcfgFieldTqbtMask  = 0x1f
)

// GetTqbt Time Quanta per Bit Time
func (r *registerCcfgType) GetTqbt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcfgFieldTqbtMask) >> RegisterCcfgFieldTqbtShift)
}

// SetTqbt Time Quanta per Bit Time
func (r *registerCcfgType) SetTqbt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcfgFieldTqbtMask)|(uint32(value)<<RegisterCcfgFieldTqbtShift))
}

const (
	RegisterCcfgFieldBccShift = 6
	RegisterCcfgFieldBccMask  = 0x40
)

// GetBcc Bypass Clock Calibration
func (r *registerCcfgType) GetBcc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcfgFieldBccMask) != 0
}

// SetBcc Bypass Clock Calibration
func (r *registerCcfgType) SetBcc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcfgFieldBccMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcfgFieldBccMask)
	}
}

const (
	RegisterCcfgFieldCflShift = 7
	RegisterCcfgFieldCflMask  = 0x80
)

// GetCfl Calibration Field Length
func (r *registerCcfgType) GetCfl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcfgFieldCflMask) != 0
}

// SetCfl Calibration Field Length
func (r *registerCcfgType) SetCfl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcfgFieldCflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcfgFieldCflMask)
	}
}

const (
	RegisterCcfgFieldOcpmShift = 8
	RegisterCcfgFieldOcpmMask  = 0xff00
)

// GetOcpm Oscillator Clock Periods Minimum
func (r *registerCcfgType) GetOcpm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcfgFieldOcpmMask) >> RegisterCcfgFieldOcpmShift)
}

// SetOcpm Oscillator Clock Periods Minimum
func (r *registerCcfgType) SetOcpm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcfgFieldOcpmMask)|(uint32(value)<<RegisterCcfgFieldOcpmShift))
}

const (
	RegisterCcfgFieldCdivShift = 16
	RegisterCcfgFieldCdivMask  = 0xf0000
)

// GetCdiv Clock Divider
func (r *registerCcfgType) GetCdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcfgFieldCdivMask) >> RegisterCcfgFieldCdivShift)
}

// SetCdiv Clock Divider
func (r *registerCcfgType) SetCdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcfgFieldCdivMask)|(uint32(value)<<RegisterCcfgFieldCdivShift))
}

const (
	RegisterCcfgFieldSwrShift = 31
	RegisterCcfgFieldSwrMask  = 0x80000000
)

// GetSwr Software Reset
func (r *registerCcfgType) GetSwr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcfgFieldSwrMask) != 0
}

// SetSwr Software Reset
func (r *registerCcfgType) SetSwr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcfgFieldSwrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcfgFieldSwrMask)
	}
}

// registerCstatType Calibration Status Register
type registerCstatType uint32

const (
	RegisterCstatFieldOcpcShift = 0
	RegisterCstatFieldOcpcMask  = 0x3ffff
)

// GetOcpc Oscillator Clock Period Counter
func (r *registerCstatType) GetOcpc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCstatFieldOcpcMask) >> RegisterCstatFieldOcpcShift)
}

// SetOcpc Oscillator Clock Period Counter
func (r *registerCstatType) SetOcpc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCstatFieldOcpcMask)|(uint32(value)<<RegisterCstatFieldOcpcShift))
}

const (
	RegisterCstatFieldTqcShift = 18
	RegisterCstatFieldTqcMask  = 0x1ffc0000
)

// GetTqc Time Quanta Counter
func (r *registerCstatType) GetTqc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCstatFieldTqcMask) >> RegisterCstatFieldTqcShift)
}

// SetTqc Time Quanta Counter
func (r *registerCstatType) SetTqc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCstatFieldTqcMask)|(uint32(value)<<RegisterCstatFieldTqcShift))
}

const (
	RegisterCstatFieldCalsShift = 30
	RegisterCstatFieldCalsMask  = 0xc0000000
)

// GetCals Calibration State
func (r *registerCstatType) GetCals() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCstatFieldCalsMask) >> RegisterCstatFieldCalsShift)
}

// SetCals Calibration State
func (r *registerCstatType) SetCals(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCstatFieldCalsMask)|(uint32(value)<<RegisterCstatFieldCalsShift))
}

// registerCwdType Calibration Watchdog Register
type registerCwdType uint32

const (
	RegisterCwdFieldWdcShift = 0
	RegisterCwdFieldWdcMask  = 0xffff
)

// GetWdc WDC
func (r *registerCwdType) GetWdc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCwdFieldWdcMask) >> RegisterCwdFieldWdcShift)
}

// SetWdc WDC
func (r *registerCwdType) SetWdc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwdFieldWdcMask)|(uint32(value)<<RegisterCwdFieldWdcShift))
}

const (
	RegisterCwdFieldWdvShift = 16
	RegisterCwdFieldWdvMask  = 0xffff0000
)

// GetWdv WDV
func (r *registerCwdType) GetWdv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCwdFieldWdvMask) >> RegisterCwdFieldWdvShift)
}

// SetWdv WDV
func (r *registerCwdType) SetWdv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwdFieldWdvMask)|(uint32(value)<<RegisterCwdFieldWdvShift))
}

// registerIrType Clock Calibration Unit Interrupt Register
type registerIrType uint32

const (
	RegisterIrFieldCweShift = 0
	RegisterIrFieldCweMask  = 0x1
)

// GetCwe Calibration Watchdog Event
func (r *registerIrType) GetCwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIrFieldCweMask) != 0
}

// SetCwe Calibration Watchdog Event
func (r *registerIrType) SetCwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIrFieldCweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIrFieldCweMask)
	}
}

const (
	RegisterIrFieldCscShift = 1
	RegisterIrFieldCscMask  = 0x2
)

// GetCsc Calibration State Changed
func (r *registerIrType) GetCsc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIrFieldCscMask) != 0
}

// SetCsc Calibration State Changed
func (r *registerIrType) SetCsc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIrFieldCscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIrFieldCscMask)
	}
}

// registerIeType Clock Calibration Unit Interrupt Enable Register
type registerIeType uint32

const (
	RegisterIeFieldCweeShift = 0
	RegisterIeFieldCweeMask  = 0x1
)

// GetCwee Calibration Watchdog Event Enable
func (r *registerIeType) GetCwee() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIeFieldCweeMask) != 0
}

// SetCwee Calibration Watchdog Event Enable
func (r *registerIeType) SetCwee(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIeFieldCweeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIeFieldCweeMask)
	}
}

const (
	RegisterIeFieldCsceShift = 1
	RegisterIeFieldCsceMask  = 0x2
)

// GetCsce Calibration State Changed Enable
func (r *registerIeType) GetCsce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIeFieldCsceMask) != 0
}

// SetCsce Calibration State Changed Enable
func (r *registerIeType) SetCsce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIeFieldCsceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIeFieldCsceMask)
	}
}
