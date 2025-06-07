//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package comp

import (
	"unsafe"
	"volatile"
)

var (
	Comp = (*_comp)(unsafe.Pointer(uintptr(0x58003800)))
)

type _comp struct {
	Sr    RegisterSrType
	Icfr  RegisterIcfrType
	Or    RegisterOrType
	Cfgr1 RegisterCfgr1Type
	Cfgr2 RegisterCfgr2Type
}

// RegisterSrType Comparator status register
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
	RegisterSrFieldC1valShift = 0
	RegisterSrFieldC1valMask  = 0x1
)

// GetC1val COMP channel 1 output status bit
func (r *RegisterSrType) GetC1val() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldC1valMask) != 0
}

// SetC1val COMP channel 1 output status bit
func (r *RegisterSrType) SetC1val(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldC1valMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldC1valMask)
	}
}

const (
	RegisterSrFieldC2valShift = 1
	RegisterSrFieldC2valMask  = 0x2
)

// GetC2val COMP channel 2 output status bit
func (r *RegisterSrType) GetC2val() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldC2valMask) != 0
}

// SetC2val COMP channel 2 output status bit
func (r *RegisterSrType) SetC2val(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldC2valMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldC2valMask)
	}
}

const (
	RegisterSrFieldC1ifShift = 16
	RegisterSrFieldC1ifMask  = 0x10000
)

// GetC1if COMP channel 1 Interrupt Flag
func (r *RegisterSrType) GetC1if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldC1ifMask) != 0
}

// SetC1if COMP channel 1 Interrupt Flag
func (r *RegisterSrType) SetC1if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldC1ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldC1ifMask)
	}
}

const (
	RegisterSrFieldC2ifShift = 17
	RegisterSrFieldC2ifMask  = 0x20000
)

// GetC2if COMP channel 2 Interrupt Flag
func (r *RegisterSrType) GetC2if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldC2ifMask) != 0
}

// SetC2if COMP channel 2 Interrupt Flag
func (r *RegisterSrType) SetC2if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldC2ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldC2ifMask)
	}
}

// RegisterIcfrType Comparator interrupt clear flag register
type RegisterIcfrType uint32

func (r *RegisterIcfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIcfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIcfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIcfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIcfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIcfrFieldCc1ifShift = 16
	RegisterIcfrFieldCc1ifMask  = 0x10000
)

// GetCc1if Clear COMP channel 1 Interrupt Flag
func (r *RegisterIcfrType) GetCc1if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcfrFieldCc1ifMask) != 0
}

// SetCc1if Clear COMP channel 1 Interrupt Flag
func (r *RegisterIcfrType) SetCc1if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcfrFieldCc1ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcfrFieldCc1ifMask)
	}
}

const (
	RegisterIcfrFieldCc2ifShift = 17
	RegisterIcfrFieldCc2ifMask  = 0x20000
)

// GetCc2if Clear COMP channel 2 Interrupt Flag
func (r *RegisterIcfrType) GetCc2if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcfrFieldCc2ifMask) != 0
}

// SetCc2if Clear COMP channel 2 Interrupt Flag
func (r *RegisterIcfrType) SetCc2if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcfrFieldCc2ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcfrFieldCc2ifMask)
	}
}

// RegisterOrType Comparator option register
type RegisterOrType uint32

func (r *RegisterOrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOrFieldAfopShift = 0
	RegisterOrFieldAfopMask  = 0x7ff
)

// GetAfop Selection of source for alternate function of output ports
func (r *RegisterOrType) GetAfop() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOrFieldAfopMask) >> RegisterOrFieldAfopShift)
}

// SetAfop Selection of source for alternate function of output ports
func (r *RegisterOrType) SetAfop(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOrFieldAfopMask)|(uint32(value)<<RegisterOrFieldAfopShift))
}

const (
	RegisterOrFieldOrShift = 11
	RegisterOrFieldOrMask  = 0xfffff800
)

// GetOr Option Register
func (r *RegisterOrType) GetOr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOrFieldOrMask) >> RegisterOrFieldOrShift)
}

// SetOr Option Register
func (r *RegisterOrType) SetOr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOrFieldOrMask)|(uint32(value)<<RegisterOrFieldOrShift))
}

// RegisterCfgr1Type Comparator configuration register 1
type RegisterCfgr1Type uint32

func (r *RegisterCfgr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfgr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfgr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfgr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfgr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCfgr1FieldEnShift = 0
	RegisterCfgr1FieldEnMask  = 0x1
)

// GetEn COMP channel 1 enable bit
func (r *RegisterCfgr1Type) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldEnMask) != 0
}

// SetEn COMP channel 1 enable bit
func (r *RegisterCfgr1Type) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr1FieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldEnMask)
	}
}

const (
	RegisterCfgr1FieldBrgenShift = 1
	RegisterCfgr1FieldBrgenMask  = 0x2
)

// GetBrgen Scaler bridge enable
func (r *RegisterCfgr1Type) GetBrgen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldBrgenMask) != 0
}

// SetBrgen Scaler bridge enable
func (r *RegisterCfgr1Type) SetBrgen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr1FieldBrgenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldBrgenMask)
	}
}

const (
	RegisterCfgr1FieldScalenShift = 2
	RegisterCfgr1FieldScalenMask  = 0x4
)

// GetScalen Voltage scaler enable bit
func (r *RegisterCfgr1Type) GetScalen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldScalenMask) != 0
}

// SetScalen Voltage scaler enable bit
func (r *RegisterCfgr1Type) SetScalen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr1FieldScalenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldScalenMask)
	}
}

const (
	RegisterCfgr1FieldPolarityShift = 3
	RegisterCfgr1FieldPolarityMask  = 0x8
)

// GetPolarity COMP channel 1 polarity selection bit
func (r *RegisterCfgr1Type) GetPolarity() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldPolarityMask) != 0
}

// SetPolarity COMP channel 1 polarity selection bit
func (r *RegisterCfgr1Type) SetPolarity(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr1FieldPolarityMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldPolarityMask)
	}
}

const (
	RegisterCfgr1FieldItenShift = 6
	RegisterCfgr1FieldItenMask  = 0x40
)

// GetIten COMP channel 1 interrupt enable
func (r *RegisterCfgr1Type) GetIten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldItenMask) != 0
}

// SetIten COMP channel 1 interrupt enable
func (r *RegisterCfgr1Type) SetIten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr1FieldItenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldItenMask)
	}
}

const (
	RegisterCfgr1FieldHystShift = 8
	RegisterCfgr1FieldHystMask  = 0x300
)

// GetHyst COMP channel 1 hysteresis selection bits
func (r *RegisterCfgr1Type) GetHyst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldHystMask) >> RegisterCfgr1FieldHystShift)
}

// SetHyst COMP channel 1 hysteresis selection bits
func (r *RegisterCfgr1Type) SetHyst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldHystMask)|(uint32(value)<<RegisterCfgr1FieldHystShift))
}

const (
	RegisterCfgr1FieldPwrmodeShift = 12
	RegisterCfgr1FieldPwrmodeMask  = 0x3000
)

// GetPwrmode Power Mode of the COMP channel 1
func (r *RegisterCfgr1Type) GetPwrmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldPwrmodeMask) >> RegisterCfgr1FieldPwrmodeShift)
}

// SetPwrmode Power Mode of the COMP channel 1
func (r *RegisterCfgr1Type) SetPwrmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldPwrmodeMask)|(uint32(value)<<RegisterCfgr1FieldPwrmodeShift))
}

const (
	RegisterCfgr1FieldInmselShift = 16
	RegisterCfgr1FieldInmselMask  = 0x70000
)

// GetInmsel COMP channel 1 inverting input selection field
func (r *RegisterCfgr1Type) GetInmsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldInmselMask) >> RegisterCfgr1FieldInmselShift)
}

// SetInmsel COMP channel 1 inverting input selection field
func (r *RegisterCfgr1Type) SetInmsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldInmselMask)|(uint32(value)<<RegisterCfgr1FieldInmselShift))
}

const (
	RegisterCfgr1FieldInpselShift = 20
	RegisterCfgr1FieldInpselMask  = 0x100000
)

// GetInpsel COMP channel 1 non-inverting input selection bit
func (r *RegisterCfgr1Type) GetInpsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldInpselMask) != 0
}

// SetInpsel COMP channel 1 non-inverting input selection bit
func (r *RegisterCfgr1Type) SetInpsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr1FieldInpselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldInpselMask)
	}
}

const (
	RegisterCfgr1FieldBlankingShift = 24
	RegisterCfgr1FieldBlankingMask  = 0xf000000
)

// GetBlanking COMP channel 1 blanking source selection bits
func (r *RegisterCfgr1Type) GetBlanking() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldBlankingMask) >> RegisterCfgr1FieldBlankingShift)
}

// SetBlanking COMP channel 1 blanking source selection bits
func (r *RegisterCfgr1Type) SetBlanking(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldBlankingMask)|(uint32(value)<<RegisterCfgr1FieldBlankingShift))
}

const (
	RegisterCfgr1FieldLockShift = 31
	RegisterCfgr1FieldLockMask  = 0x80000000
)

// GetLock Lock bit
func (r *RegisterCfgr1Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldLockMask) != 0
}

// SetLock Lock bit
func (r *RegisterCfgr1Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr1FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldLockMask)
	}
}

// RegisterCfgr2Type Comparator configuration register 2
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
	RegisterCfgr2FieldEnShift = 0
	RegisterCfgr2FieldEnMask  = 0x1
)

// GetEn COMP channel 1 enable bit
func (r *RegisterCfgr2Type) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldEnMask) != 0
}

// SetEn COMP channel 1 enable bit
func (r *RegisterCfgr2Type) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldEnMask)
	}
}

const (
	RegisterCfgr2FieldBrgenShift = 1
	RegisterCfgr2FieldBrgenMask  = 0x2
)

// GetBrgen Scaler bridge enable
func (r *RegisterCfgr2Type) GetBrgen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldBrgenMask) != 0
}

// SetBrgen Scaler bridge enable
func (r *RegisterCfgr2Type) SetBrgen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldBrgenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldBrgenMask)
	}
}

const (
	RegisterCfgr2FieldScalenShift = 2
	RegisterCfgr2FieldScalenMask  = 0x4
)

// GetScalen Voltage scaler enable bit
func (r *RegisterCfgr2Type) GetScalen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldScalenMask) != 0
}

// SetScalen Voltage scaler enable bit
func (r *RegisterCfgr2Type) SetScalen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldScalenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldScalenMask)
	}
}

const (
	RegisterCfgr2FieldPolarityShift = 3
	RegisterCfgr2FieldPolarityMask  = 0x8
)

// GetPolarity COMP channel 1 polarity selection bit
func (r *RegisterCfgr2Type) GetPolarity() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldPolarityMask) != 0
}

// SetPolarity COMP channel 1 polarity selection bit
func (r *RegisterCfgr2Type) SetPolarity(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldPolarityMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldPolarityMask)
	}
}

const (
	RegisterCfgr2FieldWinmodeShift = 4
	RegisterCfgr2FieldWinmodeMask  = 0x10
)

// GetWinmode Window comparator mode selection bit
func (r *RegisterCfgr2Type) GetWinmode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldWinmodeMask) != 0
}

// SetWinmode Window comparator mode selection bit
func (r *RegisterCfgr2Type) SetWinmode(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldWinmodeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldWinmodeMask)
	}
}

const (
	RegisterCfgr2FieldItenShift = 6
	RegisterCfgr2FieldItenMask  = 0x40
)

// GetIten COMP channel 1 interrupt enable
func (r *RegisterCfgr2Type) GetIten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldItenMask) != 0
}

// SetIten COMP channel 1 interrupt enable
func (r *RegisterCfgr2Type) SetIten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldItenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldItenMask)
	}
}

const (
	RegisterCfgr2FieldHystShift = 8
	RegisterCfgr2FieldHystMask  = 0x300
)

// GetHyst COMP channel 1 hysteresis selection bits
func (r *RegisterCfgr2Type) GetHyst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldHystMask) >> RegisterCfgr2FieldHystShift)
}

// SetHyst COMP channel 1 hysteresis selection bits
func (r *RegisterCfgr2Type) SetHyst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldHystMask)|(uint32(value)<<RegisterCfgr2FieldHystShift))
}

const (
	RegisterCfgr2FieldPwrmodeShift = 12
	RegisterCfgr2FieldPwrmodeMask  = 0x3000
)

// GetPwrmode Power Mode of the COMP channel 1
func (r *RegisterCfgr2Type) GetPwrmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldPwrmodeMask) >> RegisterCfgr2FieldPwrmodeShift)
}

// SetPwrmode Power Mode of the COMP channel 1
func (r *RegisterCfgr2Type) SetPwrmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldPwrmodeMask)|(uint32(value)<<RegisterCfgr2FieldPwrmodeShift))
}

const (
	RegisterCfgr2FieldInmselShift = 16
	RegisterCfgr2FieldInmselMask  = 0x70000
)

// GetInmsel COMP channel 1 inverting input selection field
func (r *RegisterCfgr2Type) GetInmsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldInmselMask) >> RegisterCfgr2FieldInmselShift)
}

// SetInmsel COMP channel 1 inverting input selection field
func (r *RegisterCfgr2Type) SetInmsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldInmselMask)|(uint32(value)<<RegisterCfgr2FieldInmselShift))
}

const (
	RegisterCfgr2FieldInpselShift = 20
	RegisterCfgr2FieldInpselMask  = 0x100000
)

// GetInpsel COMP channel 1 non-inverting input selection bit
func (r *RegisterCfgr2Type) GetInpsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldInpselMask) != 0
}

// SetInpsel COMP channel 1 non-inverting input selection bit
func (r *RegisterCfgr2Type) SetInpsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldInpselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldInpselMask)
	}
}

const (
	RegisterCfgr2FieldBlankingShift = 24
	RegisterCfgr2FieldBlankingMask  = 0xf000000
)

// GetBlanking COMP channel 1 blanking source selection bits
func (r *RegisterCfgr2Type) GetBlanking() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldBlankingMask) >> RegisterCfgr2FieldBlankingShift)
}

// SetBlanking COMP channel 1 blanking source selection bits
func (r *RegisterCfgr2Type) SetBlanking(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldBlankingMask)|(uint32(value)<<RegisterCfgr2FieldBlankingShift))
}

const (
	RegisterCfgr2FieldLockShift = 31
	RegisterCfgr2FieldLockMask  = 0x80000000
)

// GetLock Lock bit
func (r *RegisterCfgr2Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldLockMask) != 0
}

// SetLock Lock bit
func (r *RegisterCfgr2Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldLockMask)
	}
}
