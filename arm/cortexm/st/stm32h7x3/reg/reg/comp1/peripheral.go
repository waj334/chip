package comp1

import (
	"unsafe"
	"volatile"
)

var (
	Comp1 = (*_comp1)(unsafe.Pointer(uintptr(0x58003800)))
)

type _comp1 struct {
	Sr    registerSrType
	Icfr  registerIcfrType
	Or    registerOrType
	Cfgr1 registerCfgr1Type
	Cfgr2 registerCfgr2Type
}

// registerSrType Comparator status register
type registerSrType uint32

const (
	RegisterSrFieldC1valShift = 0
	RegisterSrFieldC1valMask  = 0x1
)

// GetC1val COMP channel 1 output status bit
func (r *registerSrType) GetC1val() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldC1valMask) != 0
}

// SetC1val COMP channel 1 output status bit
func (r *registerSrType) SetC1val(value bool) {
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
func (r *registerSrType) GetC2val() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldC2valMask) != 0
}

// SetC2val COMP channel 2 output status bit
func (r *registerSrType) SetC2val(value bool) {
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
func (r *registerSrType) GetC1if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldC1ifMask) != 0
}

// SetC1if COMP channel 1 Interrupt Flag
func (r *registerSrType) SetC1if(value bool) {
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
func (r *registerSrType) GetC2if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldC2ifMask) != 0
}

// SetC2if COMP channel 2 Interrupt Flag
func (r *registerSrType) SetC2if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldC2ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldC2ifMask)
	}
}

// registerIcfrType Comparator interrupt clear flag register
type registerIcfrType uint32

const (
	RegisterIcfrFieldCc1ifShift = 16
	RegisterIcfrFieldCc1ifMask  = 0x10000
)

// GetCc1if Clear COMP channel 1 Interrupt Flag
func (r *registerIcfrType) GetCc1if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcfrFieldCc1ifMask) != 0
}

// SetCc1if Clear COMP channel 1 Interrupt Flag
func (r *registerIcfrType) SetCc1if(value bool) {
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
func (r *registerIcfrType) GetCc2if() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcfrFieldCc2ifMask) != 0
}

// SetCc2if Clear COMP channel 2 Interrupt Flag
func (r *registerIcfrType) SetCc2if(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcfrFieldCc2ifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcfrFieldCc2ifMask)
	}
}

// registerOrType Comparator option register
type registerOrType uint32

const (
	RegisterOrFieldAfopShift = 0
	RegisterOrFieldAfopMask  = 0x7ff
)

// GetAfop Selection of source for alternate function of output ports
func (r *registerOrType) GetAfop() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOrFieldAfopMask) >> RegisterOrFieldAfopShift)
}

// SetAfop Selection of source for alternate function of output ports
func (r *registerOrType) SetAfop(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOrFieldAfopMask)|(uint32(value)<<RegisterOrFieldAfopShift))
}

const (
	RegisterOrFieldOrShift = 11
	RegisterOrFieldOrMask  = 0xfffff800
)

// GetOr Option Register
func (r *registerOrType) GetOr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOrFieldOrMask) >> RegisterOrFieldOrShift)
}

// SetOr Option Register
func (r *registerOrType) SetOr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOrFieldOrMask)|(uint32(value)<<RegisterOrFieldOrShift))
}

// registerCfgr1Type Comparator configuration register 1
type registerCfgr1Type uint32

const (
	RegisterCfgr1FieldEnShift = 0
	RegisterCfgr1FieldEnMask  = 0x1
)

// GetEn COMP channel 1 enable bit
func (r *registerCfgr1Type) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldEnMask) != 0
}

// SetEn COMP channel 1 enable bit
func (r *registerCfgr1Type) SetEn(value bool) {
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
func (r *registerCfgr1Type) GetBrgen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldBrgenMask) != 0
}

// SetBrgen Scaler bridge enable
func (r *registerCfgr1Type) SetBrgen(value bool) {
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
func (r *registerCfgr1Type) GetScalen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldScalenMask) != 0
}

// SetScalen Voltage scaler enable bit
func (r *registerCfgr1Type) SetScalen(value bool) {
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
func (r *registerCfgr1Type) GetPolarity() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldPolarityMask) != 0
}

// SetPolarity COMP channel 1 polarity selection bit
func (r *registerCfgr1Type) SetPolarity(value bool) {
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
func (r *registerCfgr1Type) GetIten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldItenMask) != 0
}

// SetIten COMP channel 1 interrupt enable
func (r *registerCfgr1Type) SetIten(value bool) {
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
func (r *registerCfgr1Type) GetHyst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldHystMask) >> RegisterCfgr1FieldHystShift)
}

// SetHyst COMP channel 1 hysteresis selection bits
func (r *registerCfgr1Type) SetHyst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldHystMask)|(uint32(value)<<RegisterCfgr1FieldHystShift))
}

const (
	RegisterCfgr1FieldPwrmodeShift = 12
	RegisterCfgr1FieldPwrmodeMask  = 0x3000
)

// GetPwrmode Power Mode of the COMP channel 1
func (r *registerCfgr1Type) GetPwrmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldPwrmodeMask) >> RegisterCfgr1FieldPwrmodeShift)
}

// SetPwrmode Power Mode of the COMP channel 1
func (r *registerCfgr1Type) SetPwrmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldPwrmodeMask)|(uint32(value)<<RegisterCfgr1FieldPwrmodeShift))
}

const (
	RegisterCfgr1FieldInmselShift = 16
	RegisterCfgr1FieldInmselMask  = 0x70000
)

// GetInmsel COMP channel 1 inverting input selection field
func (r *registerCfgr1Type) GetInmsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldInmselMask) >> RegisterCfgr1FieldInmselShift)
}

// SetInmsel COMP channel 1 inverting input selection field
func (r *registerCfgr1Type) SetInmsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldInmselMask)|(uint32(value)<<RegisterCfgr1FieldInmselShift))
}

const (
	RegisterCfgr1FieldInpselShift = 20
	RegisterCfgr1FieldInpselMask  = 0x100000
)

// GetInpsel COMP channel 1 non-inverting input selection bit
func (r *registerCfgr1Type) GetInpsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldInpselMask) != 0
}

// SetInpsel COMP channel 1 non-inverting input selection bit
func (r *registerCfgr1Type) SetInpsel(value bool) {
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
func (r *registerCfgr1Type) GetBlanking() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldBlankingMask) >> RegisterCfgr1FieldBlankingShift)
}

// SetBlanking COMP channel 1 blanking source selection bits
func (r *registerCfgr1Type) SetBlanking(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldBlankingMask)|(uint32(value)<<RegisterCfgr1FieldBlankingShift))
}

const (
	RegisterCfgr1FieldLockShift = 31
	RegisterCfgr1FieldLockMask  = 0x80000000
)

// GetLock Lock bit
func (r *registerCfgr1Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr1FieldLockMask) != 0
}

// SetLock Lock bit
func (r *registerCfgr1Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr1FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr1FieldLockMask)
	}
}

// registerCfgr2Type Comparator configuration register 2
type registerCfgr2Type uint32

const (
	RegisterCfgr2FieldEnShift = 0
	RegisterCfgr2FieldEnMask  = 0x1
)

// GetEn COMP channel 1 enable bit
func (r *registerCfgr2Type) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldEnMask) != 0
}

// SetEn COMP channel 1 enable bit
func (r *registerCfgr2Type) SetEn(value bool) {
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
func (r *registerCfgr2Type) GetBrgen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldBrgenMask) != 0
}

// SetBrgen Scaler bridge enable
func (r *registerCfgr2Type) SetBrgen(value bool) {
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
func (r *registerCfgr2Type) GetScalen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldScalenMask) != 0
}

// SetScalen Voltage scaler enable bit
func (r *registerCfgr2Type) SetScalen(value bool) {
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
func (r *registerCfgr2Type) GetPolarity() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldPolarityMask) != 0
}

// SetPolarity COMP channel 1 polarity selection bit
func (r *registerCfgr2Type) SetPolarity(value bool) {
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
func (r *registerCfgr2Type) GetWinmode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldWinmodeMask) != 0
}

// SetWinmode Window comparator mode selection bit
func (r *registerCfgr2Type) SetWinmode(value bool) {
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
func (r *registerCfgr2Type) GetIten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldItenMask) != 0
}

// SetIten COMP channel 1 interrupt enable
func (r *registerCfgr2Type) SetIten(value bool) {
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
func (r *registerCfgr2Type) GetHyst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldHystMask) >> RegisterCfgr2FieldHystShift)
}

// SetHyst COMP channel 1 hysteresis selection bits
func (r *registerCfgr2Type) SetHyst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldHystMask)|(uint32(value)<<RegisterCfgr2FieldHystShift))
}

const (
	RegisterCfgr2FieldPwrmodeShift = 12
	RegisterCfgr2FieldPwrmodeMask  = 0x3000
)

// GetPwrmode Power Mode of the COMP channel 1
func (r *registerCfgr2Type) GetPwrmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldPwrmodeMask) >> RegisterCfgr2FieldPwrmodeShift)
}

// SetPwrmode Power Mode of the COMP channel 1
func (r *registerCfgr2Type) SetPwrmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldPwrmodeMask)|(uint32(value)<<RegisterCfgr2FieldPwrmodeShift))
}

const (
	RegisterCfgr2FieldInmselShift = 16
	RegisterCfgr2FieldInmselMask  = 0x70000
)

// GetInmsel COMP channel 1 inverting input selection field
func (r *registerCfgr2Type) GetInmsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldInmselMask) >> RegisterCfgr2FieldInmselShift)
}

// SetInmsel COMP channel 1 inverting input selection field
func (r *registerCfgr2Type) SetInmsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldInmselMask)|(uint32(value)<<RegisterCfgr2FieldInmselShift))
}

const (
	RegisterCfgr2FieldInpselShift = 20
	RegisterCfgr2FieldInpselMask  = 0x100000
)

// GetInpsel COMP channel 1 non-inverting input selection bit
func (r *registerCfgr2Type) GetInpsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldInpselMask) != 0
}

// SetInpsel COMP channel 1 non-inverting input selection bit
func (r *registerCfgr2Type) SetInpsel(value bool) {
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
func (r *registerCfgr2Type) GetBlanking() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldBlankingMask) >> RegisterCfgr2FieldBlankingShift)
}

// SetBlanking COMP channel 1 blanking source selection bits
func (r *registerCfgr2Type) SetBlanking(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldBlankingMask)|(uint32(value)<<RegisterCfgr2FieldBlankingShift))
}

const (
	RegisterCfgr2FieldLockShift = 31
	RegisterCfgr2FieldLockMask  = 0x80000000
)

// GetLock Lock bit
func (r *registerCfgr2Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldLockMask) != 0
}

// SetLock Lock bit
func (r *registerCfgr2Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgr2FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldLockMask)
	}
}
