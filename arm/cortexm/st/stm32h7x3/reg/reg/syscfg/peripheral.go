package syscfg

import (
	"unsafe"
	"volatile"
)

var (
	Syscfg = (*_syscfg)(unsafe.Pointer(uintptr(0x58000400)))
)

type _syscfg struct {
	_       [4]byte
	Pmcr    registerPmcrType
	Exticr1 registerExticr1Type
	Exticr2 registerExticr2Type
	Exticr3 registerExticr3Type
	Exticr4 registerExticr4Type
	_       [8]byte
	Cccsr   registerCccsrType
	Ccvr    registerCcvrType
	Cccr    registerCccrType
	Pwrcr   registerPwrcrType
	_       [244]byte
	Pkgr    registerPkgrType
	_       [472]byte
	Ur0     registerUr0Type
	_       [4]byte
	Ur2     registerUr2Type
	Ur3     registerUr3Type
	Ur4     registerUr4Type
	Ur5     registerUr5Type
	Ur6     registerUr6Type
	Ur7     registerUr7Type
	Ur8     registerUr8Type
	Ur9     registerUr9Type
	Ur10    registerUr10Type
	Ur11    registerUr11Type
	Ur12    registerUr12Type
	Ur13    registerUr13Type
	Ur14    registerUr14Type
	Ur15    registerUr15Type
	Ur16    registerUr16Type
	Ur17    registerUr17Type
}

// registerPmcrType peripheral mode configuration register
type registerPmcrType uint32

const (
	RegisterPmcrFieldI2c1fmpShift = 0
	RegisterPmcrFieldI2c1fmpMask  = 0x1
)

// GetI2c1fmp I2C1 Fm+
func (r *registerPmcrType) GetI2c1fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldI2c1fmpMask) != 0
}

// SetI2c1fmp I2C1 Fm+
func (r *registerPmcrType) SetI2c1fmp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldI2c1fmpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldI2c1fmpMask)
	}
}

const (
	RegisterPmcrFieldI2c2fmpShift = 1
	RegisterPmcrFieldI2c2fmpMask  = 0x2
)

// GetI2c2fmp I2C2 Fm+
func (r *registerPmcrType) GetI2c2fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldI2c2fmpMask) != 0
}

// SetI2c2fmp I2C2 Fm+
func (r *registerPmcrType) SetI2c2fmp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldI2c2fmpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldI2c2fmpMask)
	}
}

const (
	RegisterPmcrFieldI2c3fmpShift = 2
	RegisterPmcrFieldI2c3fmpMask  = 0x4
)

// GetI2c3fmp I2C3 Fm+
func (r *registerPmcrType) GetI2c3fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldI2c3fmpMask) != 0
}

// SetI2c3fmp I2C3 Fm+
func (r *registerPmcrType) SetI2c3fmp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldI2c3fmpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldI2c3fmpMask)
	}
}

const (
	RegisterPmcrFieldI2c4fmpShift = 3
	RegisterPmcrFieldI2c4fmpMask  = 0x8
)

// GetI2c4fmp I2C4 Fm+
func (r *registerPmcrType) GetI2c4fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldI2c4fmpMask) != 0
}

// SetI2c4fmp I2C4 Fm+
func (r *registerPmcrType) SetI2c4fmp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldI2c4fmpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldI2c4fmpMask)
	}
}

const (
	RegisterPmcrFieldPb6fmpShift = 4
	RegisterPmcrFieldPb6fmpMask  = 0x10
)

// GetPb6fmp PB(6) Fm+
func (r *registerPmcrType) GetPb6fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPb6fmpMask) != 0
}

// SetPb6fmp PB(6) Fm+
func (r *registerPmcrType) SetPb6fmp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldPb6fmpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldPb6fmpMask)
	}
}

const (
	RegisterPmcrFieldPb7fmpShift = 5
	RegisterPmcrFieldPb7fmpMask  = 0x20
)

// GetPb7fmp PB(7) Fast Mode Plus
func (r *registerPmcrType) GetPb7fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPb7fmpMask) != 0
}

// SetPb7fmp PB(7) Fast Mode Plus
func (r *registerPmcrType) SetPb7fmp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldPb7fmpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldPb7fmpMask)
	}
}

const (
	RegisterPmcrFieldPb8fmpShift = 6
	RegisterPmcrFieldPb8fmpMask  = 0x40
)

// GetPb8fmp PB(8) Fast Mode Plus
func (r *registerPmcrType) GetPb8fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPb8fmpMask) != 0
}

// SetPb8fmp PB(8) Fast Mode Plus
func (r *registerPmcrType) SetPb8fmp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldPb8fmpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldPb8fmpMask)
	}
}

const (
	RegisterPmcrFieldPb9fmpShift = 7
	RegisterPmcrFieldPb9fmpMask  = 0x80
)

// GetPb9fmp PB(9) Fm+
func (r *registerPmcrType) GetPb9fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPb9fmpMask) != 0
}

// SetPb9fmp PB(9) Fm+
func (r *registerPmcrType) SetPb9fmp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldPb9fmpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldPb9fmpMask)
	}
}

const (
	RegisterPmcrFieldBoosteShift = 8
	RegisterPmcrFieldBoosteMask  = 0x100
)

// GetBooste Booster Enable
func (r *registerPmcrType) GetBooste() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldBoosteMask) != 0
}

// SetBooste Booster Enable
func (r *registerPmcrType) SetBooste(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldBoosteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldBoosteMask)
	}
}

const (
	RegisterPmcrFieldBoostvddselShift = 9
	RegisterPmcrFieldBoostvddselMask  = 0x200
)

// GetBoostvddsel Analog switch supply voltage selection
func (r *registerPmcrType) GetBoostvddsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldBoostvddselMask) != 0
}

// SetBoostvddsel Analog switch supply voltage selection
func (r *registerPmcrType) SetBoostvddsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldBoostvddselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldBoostvddselMask)
	}
}

const (
	RegisterPmcrFieldEpisShift = 21
	RegisterPmcrFieldEpisMask  = 0xe00000
)

// GetEpis Ethernet PHY Interface Selection
func (r *registerPmcrType) GetEpis() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldEpisMask) >> RegisterPmcrFieldEpisShift)
}

// SetEpis Ethernet PHY Interface Selection
func (r *registerPmcrType) SetEpis(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldEpisMask)|(uint32(value)<<RegisterPmcrFieldEpisShift))
}

const (
	RegisterPmcrFieldPa0soShift = 24
	RegisterPmcrFieldPa0soMask  = 0x1000000
)

// GetPa0so PA0 Switch Open
func (r *registerPmcrType) GetPa0so() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPa0soMask) != 0
}

// SetPa0so PA0 Switch Open
func (r *registerPmcrType) SetPa0so(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldPa0soMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldPa0soMask)
	}
}

const (
	RegisterPmcrFieldPa1soShift = 25
	RegisterPmcrFieldPa1soMask  = 0x2000000
)

// GetPa1so PA1 Switch Open
func (r *registerPmcrType) GetPa1so() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPa1soMask) != 0
}

// SetPa1so PA1 Switch Open
func (r *registerPmcrType) SetPa1so(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldPa1soMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldPa1soMask)
	}
}

const (
	RegisterPmcrFieldPc2soShift = 26
	RegisterPmcrFieldPc2soMask  = 0x4000000
)

// GetPc2so PC2 Switch Open
func (r *registerPmcrType) GetPc2so() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPc2soMask) != 0
}

// SetPc2so PC2 Switch Open
func (r *registerPmcrType) SetPc2so(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldPc2soMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldPc2soMask)
	}
}

const (
	RegisterPmcrFieldPc3soShift = 27
	RegisterPmcrFieldPc3soMask  = 0x8000000
)

// GetPc3so PC3 Switch Open
func (r *registerPmcrType) GetPc3so() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPc3soMask) != 0
}

// SetPc3so PC3 Switch Open
func (r *registerPmcrType) SetPc3so(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldPc3soMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldPc3soMask)
	}
}

// registerExticr1Type external interrupt configuration register 1
type registerExticr1Type uint32

const (
	RegisterExticr1FieldExti3Shift = 12
	RegisterExticr1FieldExti3Mask  = 0xf000
)

// GetExti3 EXTI x configuration (x = 0 to 3)
func (r *registerExticr1Type) GetExti3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr1FieldExti3Mask) >> RegisterExticr1FieldExti3Shift)
}

// SetExti3 EXTI x configuration (x = 0 to 3)
func (r *registerExticr1Type) SetExti3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr1FieldExti3Mask)|(uint32(value)<<RegisterExticr1FieldExti3Shift))
}

const (
	RegisterExticr1FieldExti2Shift = 8
	RegisterExticr1FieldExti2Mask  = 0xf00
)

// GetExti2 EXTI x configuration (x = 0 to 3)
func (r *registerExticr1Type) GetExti2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr1FieldExti2Mask) >> RegisterExticr1FieldExti2Shift)
}

// SetExti2 EXTI x configuration (x = 0 to 3)
func (r *registerExticr1Type) SetExti2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr1FieldExti2Mask)|(uint32(value)<<RegisterExticr1FieldExti2Shift))
}

const (
	RegisterExticr1FieldExti1Shift = 4
	RegisterExticr1FieldExti1Mask  = 0xf0
)

// GetExti1 EXTI x configuration (x = 0 to 3)
func (r *registerExticr1Type) GetExti1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr1FieldExti1Mask) >> RegisterExticr1FieldExti1Shift)
}

// SetExti1 EXTI x configuration (x = 0 to 3)
func (r *registerExticr1Type) SetExti1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr1FieldExti1Mask)|(uint32(value)<<RegisterExticr1FieldExti1Shift))
}

const (
	RegisterExticr1FieldExti0Shift = 0
	RegisterExticr1FieldExti0Mask  = 0xf
)

// GetExti0 EXTI x configuration (x = 0 to 3)
func (r *registerExticr1Type) GetExti0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr1FieldExti0Mask) >> RegisterExticr1FieldExti0Shift)
}

// SetExti0 EXTI x configuration (x = 0 to 3)
func (r *registerExticr1Type) SetExti0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr1FieldExti0Mask)|(uint32(value)<<RegisterExticr1FieldExti0Shift))
}

// registerExticr2Type external interrupt configuration register 2
type registerExticr2Type uint32

const (
	RegisterExticr2FieldExti7Shift = 12
	RegisterExticr2FieldExti7Mask  = 0xf000
)

// GetExti7 EXTI x configuration (x = 4 to 7)
func (r *registerExticr2Type) GetExti7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr2FieldExti7Mask) >> RegisterExticr2FieldExti7Shift)
}

// SetExti7 EXTI x configuration (x = 4 to 7)
func (r *registerExticr2Type) SetExti7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr2FieldExti7Mask)|(uint32(value)<<RegisterExticr2FieldExti7Shift))
}

const (
	RegisterExticr2FieldExti6Shift = 8
	RegisterExticr2FieldExti6Mask  = 0xf00
)

// GetExti6 EXTI x configuration (x = 4 to 7)
func (r *registerExticr2Type) GetExti6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr2FieldExti6Mask) >> RegisterExticr2FieldExti6Shift)
}

// SetExti6 EXTI x configuration (x = 4 to 7)
func (r *registerExticr2Type) SetExti6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr2FieldExti6Mask)|(uint32(value)<<RegisterExticr2FieldExti6Shift))
}

const (
	RegisterExticr2FieldExti5Shift = 4
	RegisterExticr2FieldExti5Mask  = 0xf0
)

// GetExti5 EXTI x configuration (x = 4 to 7)
func (r *registerExticr2Type) GetExti5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr2FieldExti5Mask) >> RegisterExticr2FieldExti5Shift)
}

// SetExti5 EXTI x configuration (x = 4 to 7)
func (r *registerExticr2Type) SetExti5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr2FieldExti5Mask)|(uint32(value)<<RegisterExticr2FieldExti5Shift))
}

const (
	RegisterExticr2FieldExti4Shift = 0
	RegisterExticr2FieldExti4Mask  = 0xf
)

// GetExti4 EXTI x configuration (x = 4 to 7)
func (r *registerExticr2Type) GetExti4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr2FieldExti4Mask) >> RegisterExticr2FieldExti4Shift)
}

// SetExti4 EXTI x configuration (x = 4 to 7)
func (r *registerExticr2Type) SetExti4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr2FieldExti4Mask)|(uint32(value)<<RegisterExticr2FieldExti4Shift))
}

// registerExticr3Type external interrupt configuration register 3
type registerExticr3Type uint32

const (
	RegisterExticr3FieldExti11Shift = 12
	RegisterExticr3FieldExti11Mask  = 0xf000
)

// GetExti11 EXTI x configuration (x = 8 to 11)
func (r *registerExticr3Type) GetExti11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr3FieldExti11Mask) >> RegisterExticr3FieldExti11Shift)
}

// SetExti11 EXTI x configuration (x = 8 to 11)
func (r *registerExticr3Type) SetExti11(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr3FieldExti11Mask)|(uint32(value)<<RegisterExticr3FieldExti11Shift))
}

const (
	RegisterExticr3FieldExti10Shift = 8
	RegisterExticr3FieldExti10Mask  = 0xf00
)

// GetExti10 EXTI10
func (r *registerExticr3Type) GetExti10() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr3FieldExti10Mask) >> RegisterExticr3FieldExti10Shift)
}

// SetExti10 EXTI10
func (r *registerExticr3Type) SetExti10(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr3FieldExti10Mask)|(uint32(value)<<RegisterExticr3FieldExti10Shift))
}

const (
	RegisterExticr3FieldExti9Shift = 4
	RegisterExticr3FieldExti9Mask  = 0xf0
)

// GetExti9 EXTI x configuration (x = 8 to 11)
func (r *registerExticr3Type) GetExti9() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr3FieldExti9Mask) >> RegisterExticr3FieldExti9Shift)
}

// SetExti9 EXTI x configuration (x = 8 to 11)
func (r *registerExticr3Type) SetExti9(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr3FieldExti9Mask)|(uint32(value)<<RegisterExticr3FieldExti9Shift))
}

const (
	RegisterExticr3FieldExti8Shift = 0
	RegisterExticr3FieldExti8Mask  = 0xf
)

// GetExti8 EXTI x configuration (x = 8 to 11)
func (r *registerExticr3Type) GetExti8() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr3FieldExti8Mask) >> RegisterExticr3FieldExti8Shift)
}

// SetExti8 EXTI x configuration (x = 8 to 11)
func (r *registerExticr3Type) SetExti8(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr3FieldExti8Mask)|(uint32(value)<<RegisterExticr3FieldExti8Shift))
}

// registerExticr4Type external interrupt configuration register 4
type registerExticr4Type uint32

const (
	RegisterExticr4FieldExti15Shift = 12
	RegisterExticr4FieldExti15Mask  = 0xf000
)

// GetExti15 EXTI x configuration (x = 12 to 15)
func (r *registerExticr4Type) GetExti15() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr4FieldExti15Mask) >> RegisterExticr4FieldExti15Shift)
}

// SetExti15 EXTI x configuration (x = 12 to 15)
func (r *registerExticr4Type) SetExti15(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr4FieldExti15Mask)|(uint32(value)<<RegisterExticr4FieldExti15Shift))
}

const (
	RegisterExticr4FieldExti14Shift = 8
	RegisterExticr4FieldExti14Mask  = 0xf00
)

// GetExti14 EXTI x configuration (x = 12 to 15)
func (r *registerExticr4Type) GetExti14() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr4FieldExti14Mask) >> RegisterExticr4FieldExti14Shift)
}

// SetExti14 EXTI x configuration (x = 12 to 15)
func (r *registerExticr4Type) SetExti14(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr4FieldExti14Mask)|(uint32(value)<<RegisterExticr4FieldExti14Shift))
}

const (
	RegisterExticr4FieldExti13Shift = 4
	RegisterExticr4FieldExti13Mask  = 0xf0
)

// GetExti13 EXTI x configuration (x = 12 to 15)
func (r *registerExticr4Type) GetExti13() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr4FieldExti13Mask) >> RegisterExticr4FieldExti13Shift)
}

// SetExti13 EXTI x configuration (x = 12 to 15)
func (r *registerExticr4Type) SetExti13(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr4FieldExti13Mask)|(uint32(value)<<RegisterExticr4FieldExti13Shift))
}

const (
	RegisterExticr4FieldExti12Shift = 0
	RegisterExticr4FieldExti12Mask  = 0xf
)

// GetExti12 EXTI x configuration (x = 12 to 15)
func (r *registerExticr4Type) GetExti12() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr4FieldExti12Mask) >> RegisterExticr4FieldExti12Shift)
}

// SetExti12 EXTI x configuration (x = 12 to 15)
func (r *registerExticr4Type) SetExti12(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr4FieldExti12Mask)|(uint32(value)<<RegisterExticr4FieldExti12Shift))
}

// registerCccsrType compensation cell control/status register
type registerCccsrType uint32

const (
	RegisterCccsrFieldEnShift = 0
	RegisterCccsrFieldEnMask  = 0x1
)

// GetEn enable
func (r *registerCccsrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCccsrFieldEnMask) != 0
}

// SetEn enable
func (r *registerCccsrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCccsrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCccsrFieldEnMask)
	}
}

const (
	RegisterCccsrFieldCsShift = 1
	RegisterCccsrFieldCsMask  = 0x2
)

// GetCs Code selection
func (r *registerCccsrType) GetCs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCccsrFieldCsMask) != 0
}

// SetCs Code selection
func (r *registerCccsrType) SetCs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCccsrFieldCsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCccsrFieldCsMask)
	}
}

const (
	RegisterCccsrFieldReadyShift = 8
	RegisterCccsrFieldReadyMask  = 0x100
)

// GetReady Compensation cell ready flag
func (r *registerCccsrType) GetReady() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCccsrFieldReadyMask) != 0
}

// SetReady Compensation cell ready flag
func (r *registerCccsrType) SetReady(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCccsrFieldReadyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCccsrFieldReadyMask)
	}
}

const (
	RegisterCccsrFieldHslvShift = 16
	RegisterCccsrFieldHslvMask  = 0x10000
)

// GetHslv High-speed at low-voltage
func (r *registerCccsrType) GetHslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCccsrFieldHslvMask) != 0
}

// SetHslv High-speed at low-voltage
func (r *registerCccsrType) SetHslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCccsrFieldHslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCccsrFieldHslvMask)
	}
}

// registerCcvrType SYSCFG compensation cell value register
type registerCcvrType uint32

const (
	RegisterCcvrFieldNcvShift = 0
	RegisterCcvrFieldNcvMask  = 0xf
)

// GetNcv NMOS compensation value
func (r *registerCcvrType) GetNcv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcvrFieldNcvMask) >> RegisterCcvrFieldNcvShift)
}

// SetNcv NMOS compensation value
func (r *registerCcvrType) SetNcv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcvrFieldNcvMask)|(uint32(value)<<RegisterCcvrFieldNcvShift))
}

const (
	RegisterCcvrFieldPcvShift = 4
	RegisterCcvrFieldPcvMask  = 0xf0
)

// GetPcv PMOS compensation value
func (r *registerCcvrType) GetPcv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcvrFieldPcvMask) >> RegisterCcvrFieldPcvShift)
}

// SetPcv PMOS compensation value
func (r *registerCcvrType) SetPcv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcvrFieldPcvMask)|(uint32(value)<<RegisterCcvrFieldPcvShift))
}

// registerCccrType SYSCFG compensation cell code register
type registerCccrType uint32

const (
	RegisterCccrFieldNccShift = 0
	RegisterCccrFieldNccMask  = 0xf
)

// GetNcc NMOS compensation code
func (r *registerCccrType) GetNcc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCccrFieldNccMask) >> RegisterCccrFieldNccShift)
}

// SetNcc NMOS compensation code
func (r *registerCccrType) SetNcc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCccrFieldNccMask)|(uint32(value)<<RegisterCccrFieldNccShift))
}

const (
	RegisterCccrFieldPccShift = 4
	RegisterCccrFieldPccMask  = 0xf0
)

// GetPcc PMOS compensation code
func (r *registerCccrType) GetPcc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCccrFieldPccMask) >> RegisterCccrFieldPccShift)
}

// SetPcc PMOS compensation code
func (r *registerCccrType) SetPcc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCccrFieldPccMask)|(uint32(value)<<RegisterCccrFieldPccShift))
}

// registerPwrcrType SYSCFG power control register
type registerPwrcrType uint32

const (
	RegisterPwrcrFieldOdenShift = 0
	RegisterPwrcrFieldOdenMask  = 0xf
)

// GetOden Overdrive enable
func (r *registerPwrcrType) GetOden() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPwrcrFieldOdenMask) >> RegisterPwrcrFieldOdenShift)
}

// SetOden Overdrive enable
func (r *registerPwrcrType) SetOden(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPwrcrFieldOdenMask)|(uint32(value)<<RegisterPwrcrFieldOdenShift))
}

// registerPkgrType SYSCFG package register
type registerPkgrType uint32

const (
	RegisterPkgrFieldPkgShift = 0
	RegisterPkgrFieldPkgMask  = 0xf
)

// GetPkg Package
func (r *registerPkgrType) GetPkg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPkgrFieldPkgMask) >> RegisterPkgrFieldPkgShift)
}

// SetPkg Package
func (r *registerPkgrType) SetPkg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPkgrFieldPkgMask)|(uint32(value)<<RegisterPkgrFieldPkgShift))
}

// registerUr0Type SYSCFG user register 0
type registerUr0Type uint32

const (
	RegisterUr0FieldBksShift = 0
	RegisterUr0FieldBksMask  = 0x1
)

// GetBks Bank Swap
func (r *registerUr0Type) GetBks() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr0FieldBksMask) != 0
}

// SetBks Bank Swap
func (r *registerUr0Type) SetBks(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr0FieldBksMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr0FieldBksMask)
	}
}

const (
	RegisterUr0FieldRdpShift = 16
	RegisterUr0FieldRdpMask  = 0xff0000
)

// GetRdp Readout protection
func (r *registerUr0Type) GetRdp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterUr0FieldRdpMask) >> RegisterUr0FieldRdpShift)
}

// SetRdp Readout protection
func (r *registerUr0Type) SetRdp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr0FieldRdpMask)|(uint32(value)<<RegisterUr0FieldRdpShift))
}

// registerUr2Type SYSCFG user register 2
type registerUr2Type uint32

const (
	RegisterUr2FieldBorhShift = 0
	RegisterUr2FieldBorhMask  = 0x3
)

// GetBorh BOR_LVL Brownout Reset Threshold Level
func (r *registerUr2Type) GetBorh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterUr2FieldBorhMask) >> RegisterUr2FieldBorhShift)
}

// SetBorh BOR_LVL Brownout Reset Threshold Level
func (r *registerUr2Type) SetBorh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr2FieldBorhMask)|(uint32(value)<<RegisterUr2FieldBorhShift))
}

const (
	RegisterUr2FieldBoot_add0Shift = 16
	RegisterUr2FieldBoot_add0Mask  = 0xffff0000
)

// GetBoot_add0 Boot Address 0
func (r *registerUr2Type) GetBoot_add0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr2FieldBoot_add0Mask) >> RegisterUr2FieldBoot_add0Shift)
}

// SetBoot_add0 Boot Address 0
func (r *registerUr2Type) SetBoot_add0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr2FieldBoot_add0Mask)|(uint32(value)<<RegisterUr2FieldBoot_add0Shift))
}

// registerUr3Type SYSCFG user register 3
type registerUr3Type uint32

const (
	RegisterUr3FieldBoot_add1Shift = 16
	RegisterUr3FieldBoot_add1Mask  = 0xffff0000
)

// GetBoot_add1 Boot Address 1
func (r *registerUr3Type) GetBoot_add1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr3FieldBoot_add1Mask) >> RegisterUr3FieldBoot_add1Shift)
}

// SetBoot_add1 Boot Address 1
func (r *registerUr3Type) SetBoot_add1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr3FieldBoot_add1Mask)|(uint32(value)<<RegisterUr3FieldBoot_add1Shift))
}

// registerUr4Type SYSCFG user register 4
type registerUr4Type uint32

const (
	RegisterUr4FieldMepad_1Shift = 16
	RegisterUr4FieldMepad_1Mask  = 0x10000
)

// GetMepad_1 Mass Erase Protected Area Disabled for bank 1
func (r *registerUr4Type) GetMepad_1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr4FieldMepad_1Mask) != 0
}

// SetMepad_1 Mass Erase Protected Area Disabled for bank 1
func (r *registerUr4Type) SetMepad_1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr4FieldMepad_1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr4FieldMepad_1Mask)
	}
}

// registerUr5Type SYSCFG user register 5
type registerUr5Type uint32

const (
	RegisterUr5FieldMesad_1Shift = 0
	RegisterUr5FieldMesad_1Mask  = 0x1
)

// GetMesad_1 Mass erase secured area disabled for bank 1
func (r *registerUr5Type) GetMesad_1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr5FieldMesad_1Mask) != 0
}

// SetMesad_1 Mass erase secured area disabled for bank 1
func (r *registerUr5Type) SetMesad_1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr5FieldMesad_1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr5FieldMesad_1Mask)
	}
}

const (
	RegisterUr5FieldWrpn_1Shift = 16
	RegisterUr5FieldWrpn_1Mask  = 0xff0000
)

// GetWrpn_1 Write protection for flash bank 1
func (r *registerUr5Type) GetWrpn_1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterUr5FieldWrpn_1Mask) >> RegisterUr5FieldWrpn_1Shift)
}

// SetWrpn_1 Write protection for flash bank 1
func (r *registerUr5Type) SetWrpn_1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr5FieldWrpn_1Mask)|(uint32(value)<<RegisterUr5FieldWrpn_1Shift))
}

// registerUr6Type SYSCFG user register 6
type registerUr6Type uint32

const (
	RegisterUr6FieldPa_beg_1Shift = 0
	RegisterUr6FieldPa_beg_1Mask  = 0xfff
)

// GetPa_beg_1 Protected area start address for bank 1
func (r *registerUr6Type) GetPa_beg_1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr6FieldPa_beg_1Mask) >> RegisterUr6FieldPa_beg_1Shift)
}

// SetPa_beg_1 Protected area start address for bank 1
func (r *registerUr6Type) SetPa_beg_1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr6FieldPa_beg_1Mask)|(uint32(value)<<RegisterUr6FieldPa_beg_1Shift))
}

const (
	RegisterUr6FieldPa_end_1Shift = 16
	RegisterUr6FieldPa_end_1Mask  = 0xfff0000
)

// GetPa_end_1 Protected area end address for bank 1
func (r *registerUr6Type) GetPa_end_1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr6FieldPa_end_1Mask) >> RegisterUr6FieldPa_end_1Shift)
}

// SetPa_end_1 Protected area end address for bank 1
func (r *registerUr6Type) SetPa_end_1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr6FieldPa_end_1Mask)|(uint32(value)<<RegisterUr6FieldPa_end_1Shift))
}

// registerUr7Type SYSCFG user register 7
type registerUr7Type uint32

const (
	RegisterUr7FieldSa_beg_1Shift = 0
	RegisterUr7FieldSa_beg_1Mask  = 0xfff
)

// GetSa_beg_1 Secured area start address for bank 1
func (r *registerUr7Type) GetSa_beg_1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr7FieldSa_beg_1Mask) >> RegisterUr7FieldSa_beg_1Shift)
}

// SetSa_beg_1 Secured area start address for bank 1
func (r *registerUr7Type) SetSa_beg_1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr7FieldSa_beg_1Mask)|(uint32(value)<<RegisterUr7FieldSa_beg_1Shift))
}

const (
	RegisterUr7FieldSa_end_1Shift = 16
	RegisterUr7FieldSa_end_1Mask  = 0xfff0000
)

// GetSa_end_1 Secured area end address for bank 1
func (r *registerUr7Type) GetSa_end_1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr7FieldSa_end_1Mask) >> RegisterUr7FieldSa_end_1Shift)
}

// SetSa_end_1 Secured area end address for bank 1
func (r *registerUr7Type) SetSa_end_1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr7FieldSa_end_1Mask)|(uint32(value)<<RegisterUr7FieldSa_end_1Shift))
}

// registerUr8Type SYSCFG user register 8
type registerUr8Type uint32

const (
	RegisterUr8FieldMepad_2Shift = 0
	RegisterUr8FieldMepad_2Mask  = 0x1
)

// GetMepad_2 Mass erase protected area disabled for bank 2
func (r *registerUr8Type) GetMepad_2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr8FieldMepad_2Mask) != 0
}

// SetMepad_2 Mass erase protected area disabled for bank 2
func (r *registerUr8Type) SetMepad_2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr8FieldMepad_2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr8FieldMepad_2Mask)
	}
}

const (
	RegisterUr8FieldMesad_2Shift = 16
	RegisterUr8FieldMesad_2Mask  = 0x10000
)

// GetMesad_2 Mass erase secured area disabled for bank 2
func (r *registerUr8Type) GetMesad_2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr8FieldMesad_2Mask) != 0
}

// SetMesad_2 Mass erase secured area disabled for bank 2
func (r *registerUr8Type) SetMesad_2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr8FieldMesad_2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr8FieldMesad_2Mask)
	}
}

// registerUr9Type SYSCFG user register 9
type registerUr9Type uint32

const (
	RegisterUr9FieldWrpn_2Shift = 0
	RegisterUr9FieldWrpn_2Mask  = 0xff
)

// GetWrpn_2 Write protection for flash bank 2
func (r *registerUr9Type) GetWrpn_2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterUr9FieldWrpn_2Mask) >> RegisterUr9FieldWrpn_2Shift)
}

// SetWrpn_2 Write protection for flash bank 2
func (r *registerUr9Type) SetWrpn_2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr9FieldWrpn_2Mask)|(uint32(value)<<RegisterUr9FieldWrpn_2Shift))
}

const (
	RegisterUr9FieldPa_beg_2Shift = 16
	RegisterUr9FieldPa_beg_2Mask  = 0xfff0000
)

// GetPa_beg_2 Protected area start address for bank 2
func (r *registerUr9Type) GetPa_beg_2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr9FieldPa_beg_2Mask) >> RegisterUr9FieldPa_beg_2Shift)
}

// SetPa_beg_2 Protected area start address for bank 2
func (r *registerUr9Type) SetPa_beg_2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr9FieldPa_beg_2Mask)|(uint32(value)<<RegisterUr9FieldPa_beg_2Shift))
}

// registerUr10Type SYSCFG user register 10
type registerUr10Type uint32

const (
	RegisterUr10FieldPa_end_2Shift = 0
	RegisterUr10FieldPa_end_2Mask  = 0xfff
)

// GetPa_end_2 Protected area end address for bank 2
func (r *registerUr10Type) GetPa_end_2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr10FieldPa_end_2Mask) >> RegisterUr10FieldPa_end_2Shift)
}

// SetPa_end_2 Protected area end address for bank 2
func (r *registerUr10Type) SetPa_end_2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr10FieldPa_end_2Mask)|(uint32(value)<<RegisterUr10FieldPa_end_2Shift))
}

const (
	RegisterUr10FieldSa_beg_2Shift = 16
	RegisterUr10FieldSa_beg_2Mask  = 0xfff0000
)

// GetSa_beg_2 Secured area start address for bank 2
func (r *registerUr10Type) GetSa_beg_2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr10FieldSa_beg_2Mask) >> RegisterUr10FieldSa_beg_2Shift)
}

// SetSa_beg_2 Secured area start address for bank 2
func (r *registerUr10Type) SetSa_beg_2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr10FieldSa_beg_2Mask)|(uint32(value)<<RegisterUr10FieldSa_beg_2Shift))
}

// registerUr11Type SYSCFG user register 11
type registerUr11Type uint32

const (
	RegisterUr11FieldSa_end_2Shift = 0
	RegisterUr11FieldSa_end_2Mask  = 0xfff
)

// GetSa_end_2 Secured area end address for bank 2
func (r *registerUr11Type) GetSa_end_2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr11FieldSa_end_2Mask) >> RegisterUr11FieldSa_end_2Shift)
}

// SetSa_end_2 Secured area end address for bank 2
func (r *registerUr11Type) SetSa_end_2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr11FieldSa_end_2Mask)|(uint32(value)<<RegisterUr11FieldSa_end_2Shift))
}

const (
	RegisterUr11FieldIwdg1mShift = 16
	RegisterUr11FieldIwdg1mMask  = 0x10000
)

// GetIwdg1m Independent Watchdog 1 mode
func (r *registerUr11Type) GetIwdg1m() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr11FieldIwdg1mMask) != 0
}

// SetIwdg1m Independent Watchdog 1 mode
func (r *registerUr11Type) SetIwdg1m(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr11FieldIwdg1mMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr11FieldIwdg1mMask)
	}
}

// registerUr12Type SYSCFG user register 12
type registerUr12Type uint32

const (
	RegisterUr12FieldSecureShift = 16
	RegisterUr12FieldSecureMask  = 0x10000
)

// GetSecure Secure mode
func (r *registerUr12Type) GetSecure() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr12FieldSecureMask) != 0
}

// SetSecure Secure mode
func (r *registerUr12Type) SetSecure(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr12FieldSecureMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr12FieldSecureMask)
	}
}

// registerUr13Type SYSCFG user register 13
type registerUr13Type uint32

const (
	RegisterUr13FieldSdrsShift = 0
	RegisterUr13FieldSdrsMask  = 0x3
)

// GetSdrs Secured DTCM RAM Size
func (r *registerUr13Type) GetSdrs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterUr13FieldSdrsMask) >> RegisterUr13FieldSdrsShift)
}

// SetSdrs Secured DTCM RAM Size
func (r *registerUr13Type) SetSdrs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr13FieldSdrsMask)|(uint32(value)<<RegisterUr13FieldSdrsShift))
}

const (
	RegisterUr13FieldD1sbrstShift = 16
	RegisterUr13FieldD1sbrstMask  = 0x10000
)

// GetD1sbrst D1 Standby reset
func (r *registerUr13Type) GetD1sbrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr13FieldD1sbrstMask) != 0
}

// SetD1sbrst D1 Standby reset
func (r *registerUr13Type) SetD1sbrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr13FieldD1sbrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr13FieldD1sbrstMask)
	}
}

// registerUr14Type SYSCFG user register 14
type registerUr14Type uint32

const (
	RegisterUr14FieldD1stprstShift = 0
	RegisterUr14FieldD1stprstMask  = 0x1
)

// GetD1stprst D1 Stop Reset
func (r *registerUr14Type) GetD1stprst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr14FieldD1stprstMask) != 0
}

// SetD1stprst D1 Stop Reset
func (r *registerUr14Type) SetD1stprst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr14FieldD1stprstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr14FieldD1stprstMask)
	}
}

// registerUr15Type SYSCFG user register 15
type registerUr15Type uint32

const (
	RegisterUr15FieldFziwdgstbShift = 16
	RegisterUr15FieldFziwdgstbMask  = 0x10000
)

// GetFziwdgstb Freeze independent watchdog in Standby mode
func (r *registerUr15Type) GetFziwdgstb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr15FieldFziwdgstbMask) != 0
}

// SetFziwdgstb Freeze independent watchdog in Standby mode
func (r *registerUr15Type) SetFziwdgstb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr15FieldFziwdgstbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr15FieldFziwdgstbMask)
	}
}

// registerUr16Type SYSCFG user register 16
type registerUr16Type uint32

const (
	RegisterUr16FieldFziwdgstpShift = 0
	RegisterUr16FieldFziwdgstpMask  = 0x1
)

// GetFziwdgstp Freeze independent watchdog in Stop mode
func (r *registerUr16Type) GetFziwdgstp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr16FieldFziwdgstpMask) != 0
}

// SetFziwdgstp Freeze independent watchdog in Stop mode
func (r *registerUr16Type) SetFziwdgstp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr16FieldFziwdgstpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr16FieldFziwdgstpMask)
	}
}

const (
	RegisterUr16FieldPkpShift = 16
	RegisterUr16FieldPkpMask  = 0x10000
)

// GetPkp Private key programmed
func (r *registerUr16Type) GetPkp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr16FieldPkpMask) != 0
}

// SetPkp Private key programmed
func (r *registerUr16Type) SetPkp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr16FieldPkpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr16FieldPkpMask)
	}
}

// registerUr17Type SYSCFG user register 17
type registerUr17Type uint32

const (
	RegisterUr17FieldIo_hslvShift = 0
	RegisterUr17FieldIo_hslvMask  = 0x1
)

// GetIo_hslv I/O high speed / low voltage
func (r *registerUr17Type) GetIo_hslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr17FieldIo_hslvMask) != 0
}

// SetIo_hslv I/O high speed / low voltage
func (r *registerUr17Type) SetIo_hslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr17FieldIo_hslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr17FieldIo_hslvMask)
	}
}
