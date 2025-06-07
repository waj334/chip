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
	Crel  RegisterCrelType
	Ccfg  RegisterCcfgType
	Cstat RegisterCstatType
	Cwd   RegisterCwdType
	Ir    RegisterIrType
	Ie    RegisterIeType
}

// RegisterCrelType Clock Calibration Unit Core Release Register
type RegisterCrelType uint32

func (r *RegisterCrelType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrelType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrelType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrelType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrelType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrelFieldDayShift = 0
	RegisterCrelFieldDayMask  = 0xff
)

// GetDay Time Stamp Day
func (r *RegisterCrelType) GetDay() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrelFieldDayMask) >> RegisterCrelFieldDayShift)
}

// SetDay Time Stamp Day
func (r *RegisterCrelType) SetDay(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrelFieldDayMask)|(uint32(value)<<RegisterCrelFieldDayShift))
}

const (
	RegisterCrelFieldMonShift = 8
	RegisterCrelFieldMonMask  = 0xff00
)

// GetMon Time Stamp Month
func (r *RegisterCrelType) GetMon() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrelFieldMonMask) >> RegisterCrelFieldMonShift)
}

// SetMon Time Stamp Month
func (r *RegisterCrelType) SetMon(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrelFieldMonMask)|(uint32(value)<<RegisterCrelFieldMonShift))
}

const (
	RegisterCrelFieldYearShift = 16
	RegisterCrelFieldYearMask  = 0xf0000
)

// GetYear Time Stamp Year
func (r *RegisterCrelType) GetYear() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrelFieldYearMask) >> RegisterCrelFieldYearShift)
}

// SetYear Time Stamp Year
func (r *RegisterCrelType) SetYear(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrelFieldYearMask)|(uint32(value)<<RegisterCrelFieldYearShift))
}

const (
	RegisterCrelFieldSubstepShift = 20
	RegisterCrelFieldSubstepMask  = 0xf00000
)

// GetSubstep Sub-step of Core Release
func (r *RegisterCrelType) GetSubstep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrelFieldSubstepMask) >> RegisterCrelFieldSubstepShift)
}

// SetSubstep Sub-step of Core Release
func (r *RegisterCrelType) SetSubstep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrelFieldSubstepMask)|(uint32(value)<<RegisterCrelFieldSubstepShift))
}

const (
	RegisterCrelFieldStepShift = 24
	RegisterCrelFieldStepMask  = 0xf000000
)

// GetStep Step of Core Release
func (r *RegisterCrelType) GetStep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrelFieldStepMask) >> RegisterCrelFieldStepShift)
}

// SetStep Step of Core Release
func (r *RegisterCrelType) SetStep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrelFieldStepMask)|(uint32(value)<<RegisterCrelFieldStepShift))
}

const (
	RegisterCrelFieldRelShift = 28
	RegisterCrelFieldRelMask  = 0xf0000000
)

// GetRel Core Release
func (r *RegisterCrelType) GetRel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrelFieldRelMask) >> RegisterCrelFieldRelShift)
}

// SetRel Core Release
func (r *RegisterCrelType) SetRel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrelFieldRelMask)|(uint32(value)<<RegisterCrelFieldRelShift))
}

// RegisterCcfgType Calibration Configuration Register
type RegisterCcfgType uint32

func (r *RegisterCcfgType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcfgType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcfgType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcfgType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcfgType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcfgFieldTqbtShift = 0
	RegisterCcfgFieldTqbtMask  = 0x1f
)

// GetTqbt Time Quanta per Bit Time
func (r *RegisterCcfgType) GetTqbt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcfgFieldTqbtMask) >> RegisterCcfgFieldTqbtShift)
}

// SetTqbt Time Quanta per Bit Time
func (r *RegisterCcfgType) SetTqbt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcfgFieldTqbtMask)|(uint32(value)<<RegisterCcfgFieldTqbtShift))
}

const (
	RegisterCcfgFieldBccShift = 6
	RegisterCcfgFieldBccMask  = 0x40
)

// GetBcc Bypass Clock Calibration
func (r *RegisterCcfgType) GetBcc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcfgFieldBccMask) != 0
}

// SetBcc Bypass Clock Calibration
func (r *RegisterCcfgType) SetBcc(value bool) {
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
func (r *RegisterCcfgType) GetCfl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcfgFieldCflMask) != 0
}

// SetCfl Calibration Field Length
func (r *RegisterCcfgType) SetCfl(value bool) {
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
func (r *RegisterCcfgType) GetOcpm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcfgFieldOcpmMask) >> RegisterCcfgFieldOcpmShift)
}

// SetOcpm Oscillator Clock Periods Minimum
func (r *RegisterCcfgType) SetOcpm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcfgFieldOcpmMask)|(uint32(value)<<RegisterCcfgFieldOcpmShift))
}

const (
	RegisterCcfgFieldCdivShift = 16
	RegisterCcfgFieldCdivMask  = 0xf0000
)

// GetCdiv Clock Divider
func (r *RegisterCcfgType) GetCdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcfgFieldCdivMask) >> RegisterCcfgFieldCdivShift)
}

// SetCdiv Clock Divider
func (r *RegisterCcfgType) SetCdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcfgFieldCdivMask)|(uint32(value)<<RegisterCcfgFieldCdivShift))
}

const (
	RegisterCcfgFieldSwrShift = 31
	RegisterCcfgFieldSwrMask  = 0x80000000
)

// GetSwr Software Reset
func (r *RegisterCcfgType) GetSwr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcfgFieldSwrMask) != 0
}

// SetSwr Software Reset
func (r *RegisterCcfgType) SetSwr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcfgFieldSwrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcfgFieldSwrMask)
	}
}

// RegisterCstatType Calibration Status Register
type RegisterCstatType uint32

func (r *RegisterCstatType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCstatType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCstatType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCstatType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCstatType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCstatFieldOcpcShift = 0
	RegisterCstatFieldOcpcMask  = 0x3ffff
)

// GetOcpc Oscillator Clock Period Counter
func (r *RegisterCstatType) GetOcpc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCstatFieldOcpcMask) >> RegisterCstatFieldOcpcShift)
}

// SetOcpc Oscillator Clock Period Counter
func (r *RegisterCstatType) SetOcpc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCstatFieldOcpcMask)|(uint32(value)<<RegisterCstatFieldOcpcShift))
}

const (
	RegisterCstatFieldTqcShift = 18
	RegisterCstatFieldTqcMask  = 0x1ffc0000
)

// GetTqc Time Quanta Counter
func (r *RegisterCstatType) GetTqc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCstatFieldTqcMask) >> RegisterCstatFieldTqcShift)
}

// SetTqc Time Quanta Counter
func (r *RegisterCstatType) SetTqc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCstatFieldTqcMask)|(uint32(value)<<RegisterCstatFieldTqcShift))
}

const (
	RegisterCstatFieldCalsShift = 30
	RegisterCstatFieldCalsMask  = 0xc0000000
)

// GetCals Calibration State
func (r *RegisterCstatType) GetCals() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCstatFieldCalsMask) >> RegisterCstatFieldCalsShift)
}

// SetCals Calibration State
func (r *RegisterCstatType) SetCals(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCstatFieldCalsMask)|(uint32(value)<<RegisterCstatFieldCalsShift))
}

// RegisterCwdType Calibration Watchdog Register
type RegisterCwdType uint32

func (r *RegisterCwdType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCwdType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCwdType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCwdType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCwdType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCwdFieldWdcShift = 0
	RegisterCwdFieldWdcMask  = 0xffff
)

// GetWdc WDC
func (r *RegisterCwdType) GetWdc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCwdFieldWdcMask) >> RegisterCwdFieldWdcShift)
}

// SetWdc WDC
func (r *RegisterCwdType) SetWdc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwdFieldWdcMask)|(uint32(value)<<RegisterCwdFieldWdcShift))
}

const (
	RegisterCwdFieldWdvShift = 16
	RegisterCwdFieldWdvMask  = 0xffff0000
)

// GetWdv WDV
func (r *RegisterCwdType) GetWdv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCwdFieldWdvMask) >> RegisterCwdFieldWdvShift)
}

// SetWdv WDV
func (r *RegisterCwdType) SetWdv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwdFieldWdvMask)|(uint32(value)<<RegisterCwdFieldWdvShift))
}

// RegisterIrType Clock Calibration Unit Interrupt Register
type RegisterIrType uint32

func (r *RegisterIrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIrFieldCweShift = 0
	RegisterIrFieldCweMask  = 0x1
)

// GetCwe Calibration Watchdog Event
func (r *RegisterIrType) GetCwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIrFieldCweMask) != 0
}

// SetCwe Calibration Watchdog Event
func (r *RegisterIrType) SetCwe(value bool) {
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
func (r *RegisterIrType) GetCsc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIrFieldCscMask) != 0
}

// SetCsc Calibration State Changed
func (r *RegisterIrType) SetCsc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIrFieldCscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIrFieldCscMask)
	}
}

// RegisterIeType Clock Calibration Unit Interrupt Enable Register
type RegisterIeType uint32

func (r *RegisterIeType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIeType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIeType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIeType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIeType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIeFieldCweeShift = 0
	RegisterIeFieldCweeMask  = 0x1
)

// GetCwee Calibration Watchdog Event Enable
func (r *RegisterIeType) GetCwee() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIeFieldCweeMask) != 0
}

// SetCwee Calibration Watchdog Event Enable
func (r *RegisterIeType) SetCwee(value bool) {
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
func (r *RegisterIeType) GetCsce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIeFieldCsceMask) != 0
}

// SetCsce Calibration State Changed Enable
func (r *RegisterIeType) SetCsce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIeFieldCsceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIeFieldCsceMask)
	}
}
