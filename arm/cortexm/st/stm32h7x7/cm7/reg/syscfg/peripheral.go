//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

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
	Cfgr    registerCfgrType
	_       [4]byte
	Cccsr   registerCccsrType
	Ccvr    registerCcvrType
	Cccr    registerCccrType
	Pwrcr   registerPwrcrType
	_       [244]byte
	Pkgr    registerPkgrType
	_       [472]byte
	Ur0     registerUr0Type
	Ur1     registerUr1Type
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

// registerExticr2Type external interrupt configuration register 2
type registerExticr2Type uint32

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

// registerExticr3Type external interrupt configuration register 3
type registerExticr3Type uint32

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

// registerExticr4Type external interrupt configuration register 4
type registerExticr4Type uint32

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

// registerCfgrType configuration register
type registerCfgrType uint32

const (
	RegisterCfgrFieldCm4lShift = 0
	RegisterCfgrFieldCm4lMask  = 0x1
)

// GetCm4l CM4L
func (r *registerCfgrType) GetCm4l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldCm4lMask) != 0
}

// SetCm4l CM4L
func (r *registerCfgrType) SetCm4l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldCm4lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldCm4lMask)
	}
}

const (
	RegisterCfgrFieldPvdlShift = 2
	RegisterCfgrFieldPvdlMask  = 0x4
)

// GetPvdl PVDL
func (r *registerCfgrType) GetPvdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldPvdlMask) != 0
}

// SetPvdl PVDL
func (r *registerCfgrType) SetPvdl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldPvdlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldPvdlMask)
	}
}

const (
	RegisterCfgrFieldFlashlShift = 3
	RegisterCfgrFieldFlashlMask  = 0x8
)

// GetFlashl FLASHL
func (r *registerCfgrType) GetFlashl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldFlashlMask) != 0
}

// SetFlashl FLASHL
func (r *registerCfgrType) SetFlashl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldFlashlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldFlashlMask)
	}
}

const (
	RegisterCfgrFieldCm7lShift = 6
	RegisterCfgrFieldCm7lMask  = 0x40
)

// GetCm7l CM7L
func (r *registerCfgrType) GetCm7l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldCm7lMask) != 0
}

// SetCm7l CM7L
func (r *registerCfgrType) SetCm7l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldCm7lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldCm7lMask)
	}
}

const (
	RegisterCfgrFieldBkramlShift = 7
	RegisterCfgrFieldBkramlMask  = 0x80
)

// GetBkraml BKRAML
func (r *registerCfgrType) GetBkraml() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldBkramlMask) != 0
}

// SetBkraml BKRAML
func (r *registerCfgrType) SetBkraml(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldBkramlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldBkramlMask)
	}
}

const (
	RegisterCfgrFieldSram4lShift = 9
	RegisterCfgrFieldSram4lMask  = 0x200
)

// GetSram4l SRAM4L
func (r *registerCfgrType) GetSram4l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSram4lMask) != 0
}

// SetSram4l SRAM4L
func (r *registerCfgrType) SetSram4l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldSram4lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSram4lMask)
	}
}

const (
	RegisterCfgrFieldSram3lShift = 10
	RegisterCfgrFieldSram3lMask  = 0x400
)

// GetSram3l SRAM3L
func (r *registerCfgrType) GetSram3l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSram3lMask) != 0
}

// SetSram3l SRAM3L
func (r *registerCfgrType) SetSram3l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldSram3lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSram3lMask)
	}
}

const (
	RegisterCfgrFieldSram2lShift = 11
	RegisterCfgrFieldSram2lMask  = 0x800
)

// GetSram2l SRAM2L
func (r *registerCfgrType) GetSram2l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSram2lMask) != 0
}

// SetSram2l SRAM2L
func (r *registerCfgrType) SetSram2l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldSram2lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSram2lMask)
	}
}

const (
	RegisterCfgrFieldSram1lShift = 12
	RegisterCfgrFieldSram1lMask  = 0x1000
)

// GetSram1l SRAM1L
func (r *registerCfgrType) GetSram1l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSram1lMask) != 0
}

// SetSram1l SRAM1L
func (r *registerCfgrType) SetSram1l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldSram1lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSram1lMask)
	}
}

const (
	RegisterCfgrFieldDtcmlShift = 13
	RegisterCfgrFieldDtcmlMask  = 0x2000
)

// GetDtcml DTCML
func (r *registerCfgrType) GetDtcml() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldDtcmlMask) != 0
}

// SetDtcml DTCML
func (r *registerCfgrType) SetDtcml(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldDtcmlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldDtcmlMask)
	}
}

const (
	RegisterCfgrFieldItcmlShift = 14
	RegisterCfgrFieldItcmlMask  = 0x4000
)

// GetItcml ITCML
func (r *registerCfgrType) GetItcml() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldItcmlMask) != 0
}

// SetItcml ITCML
func (r *registerCfgrType) SetItcml(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldItcmlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldItcmlMask)
	}
}

const (
	RegisterCfgrFieldAxisramlShift = 15
	RegisterCfgrFieldAxisramlMask  = 0x8000
)

// GetAxisraml AXISRAML
func (r *registerCfgrType) GetAxisraml() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldAxisramlMask) != 0
}

// SetAxisraml AXISRAML
func (r *registerCfgrType) SetAxisraml(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldAxisramlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldAxisramlMask)
	}
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
	RegisterPwrcrFieldOdenMask  = 0x1
)

// GetOden Overdrive enable
func (r *registerPwrcrType) GetOden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPwrcrFieldOdenMask) != 0
}

// SetOden Overdrive enable
func (r *registerPwrcrType) SetOden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPwrcrFieldOdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPwrcrFieldOdenMask)
	}
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

// registerUr1Type SYSCFG user register 1
type registerUr1Type uint32

const (
	RegisterUr1FieldBcm4Shift = 0
	RegisterUr1FieldBcm4Mask  = 0x1
)

// GetBcm4 Boot Cortex-M4
func (r *registerUr1Type) GetBcm4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr1FieldBcm4Mask) != 0
}

// SetBcm4 Boot Cortex-M4
func (r *registerUr1Type) SetBcm4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr1FieldBcm4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr1FieldBcm4Mask)
	}
}

const (
	RegisterUr1FieldBcm7Shift = 16
	RegisterUr1FieldBcm7Mask  = 0x10000
)

// GetBcm7 Boot Cortex-M7
func (r *registerUr1Type) GetBcm7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr1FieldBcm7Mask) != 0
}

// SetBcm7 Boot Cortex-M7
func (r *registerUr1Type) SetBcm7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr1FieldBcm7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr1FieldBcm7Mask)
	}
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

const (
	RegisterUr2FieldBcm7_add0Shift = 16
	RegisterUr2FieldBcm7_add0Mask  = 0xffff0000
)

// GetBcm7_add0 Cortex-M7 Boot Address 0
func (r *registerUr2Type) GetBcm7_add0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr2FieldBcm7_add0Mask) >> RegisterUr2FieldBcm7_add0Shift)
}

// SetBcm7_add0 Cortex-M7 Boot Address 0
func (r *registerUr2Type) SetBcm7_add0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr2FieldBcm7_add0Mask)|(uint32(value)<<RegisterUr2FieldBcm7_add0Shift))
}

// registerUr3Type SYSCFG user register 3
type registerUr3Type uint32

const (
	RegisterUr3FieldBcm4_add1Shift = 0
	RegisterUr3FieldBcm4_add1Mask  = 0xffff
)

// GetBcm4_add1 Cortex-M4 Boot Address 0
func (r *registerUr3Type) GetBcm4_add1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr3FieldBcm4_add1Mask) >> RegisterUr3FieldBcm4_add1Shift)
}

// SetBcm4_add1 Cortex-M4 Boot Address 0
func (r *registerUr3Type) SetBcm4_add1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr3FieldBcm4_add1Mask)|(uint32(value)<<RegisterUr3FieldBcm4_add1Shift))
}

const (
	RegisterUr3FieldBcm7_add1Shift = 16
	RegisterUr3FieldBcm7_add1Mask  = 0xffff0000
)

// GetBcm7_add1 Cortex-M7 Boot Address 1
func (r *registerUr3Type) GetBcm7_add1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr3FieldBcm7_add1Mask) >> RegisterUr3FieldBcm7_add1Shift)
}

// SetBcm7_add1 Cortex-M7 Boot Address 1
func (r *registerUr3Type) SetBcm7_add1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr3FieldBcm7_add1Mask)|(uint32(value)<<RegisterUr3FieldBcm7_add1Shift))
}

// registerUr4Type SYSCFG user register 4
type registerUr4Type uint32

const (
	RegisterUr4FieldBcm4_add1Shift = 0
	RegisterUr4FieldBcm4_add1Mask  = 0xffff
)

// GetBcm4_add1 Mass Erase Protected Area Disabled for bank 1
func (r *registerUr4Type) GetBcm4_add1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr4FieldBcm4_add1Mask) >> RegisterUr4FieldBcm4_add1Shift)
}

// SetBcm4_add1 Mass Erase Protected Area Disabled for bank 1
func (r *registerUr4Type) SetBcm4_add1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr4FieldBcm4_add1Mask)|(uint32(value)<<RegisterUr4FieldBcm4_add1Shift))
}

const (
	RegisterUr4FieldMepad_1Shift = 16
	RegisterUr4FieldMepad_1Mask  = 0x10000
)

// GetMepad_1 Boot Cortex-M4 Address 1
func (r *registerUr4Type) GetMepad_1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr4FieldMepad_1Mask) != 0
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
	RegisterUr5FieldWrps_1Shift = 16
	RegisterUr5FieldWrps_1Mask  = 0xff0000
)

// GetWrps_1 Write protection for flash bank 1
func (r *registerUr5Type) GetWrps_1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterUr5FieldWrps_1Mask) >> RegisterUr5FieldWrps_1Shift)
}

// SetWrps_1 Write protection for flash bank 1
func (r *registerUr5Type) SetWrps_1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr5FieldWrps_1Mask)|(uint32(value)<<RegisterUr5FieldWrps_1Shift))
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
	RegisterUr9FieldWrps_2Shift = 0
	RegisterUr9FieldWrps_2Mask  = 0xff
)

// GetWrps_2 Write protection for flash bank 2
func (r *registerUr9Type) GetWrps_2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterUr9FieldWrps_2Mask) >> RegisterUr9FieldWrps_2Shift)
}

// SetWrps_2 Write protection for flash bank 2
func (r *registerUr9Type) SetWrps_2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr9FieldWrps_2Mask)|(uint32(value)<<RegisterUr9FieldWrps_2Shift))
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
	RegisterUr12FieldIwdg2mShift = 0
	RegisterUr12FieldIwdg2mMask  = 0x1
)

// GetIwdg2m Independent Watchdog 2 mode
func (r *registerUr12Type) GetIwdg2m() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr12FieldIwdg2mMask) != 0
}

// SetIwdg2m Independent Watchdog 2 mode
func (r *registerUr12Type) SetIwdg2m(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr12FieldIwdg2mMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr12FieldIwdg2mMask)
	}
}

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

const (
	RegisterUr14FieldD2sbrstShift = 16
	RegisterUr14FieldD2sbrstMask  = 0x10000
)

// GetD2sbrst D2 Standby Reset
func (r *registerUr14Type) GetD2sbrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr14FieldD2sbrstMask) != 0
}

// SetD2sbrst D2 Standby Reset
func (r *registerUr14Type) SetD2sbrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr14FieldD2sbrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr14FieldD2sbrstMask)
	}
}

// registerUr15Type SYSCFG user register 15
type registerUr15Type uint32

const (
	RegisterUr15FieldD2stprstShift = 0
	RegisterUr15FieldD2stprstMask  = 0x1
)

// GetD2stprst D2 Stop Reset
func (r *registerUr15Type) GetD2stprst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr15FieldD2stprstMask) != 0
}

// SetD2stprst D2 Stop Reset
func (r *registerUr15Type) SetD2stprst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr15FieldD2stprstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr15FieldD2stprstMask)
	}
}

const (
	RegisterUr15FieldFziwdgstbShift = 16
	RegisterUr15FieldFziwdgstbMask  = 0x10000
)

// GetFziwdgstb Freeze independent watchdog in Standby mode
func (r *registerUr15Type) GetFziwdgstb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr15FieldFziwdgstbMask) != 0
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
