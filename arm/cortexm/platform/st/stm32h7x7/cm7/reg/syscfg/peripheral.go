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
	Pmcr    RegisterPmcrType
	Exticr1 RegisterExticr1Type
	Exticr2 RegisterExticr2Type
	Exticr3 RegisterExticr3Type
	Exticr4 RegisterExticr4Type
	Cfgr    RegisterCfgrType
	_       [4]byte
	Cccsr   RegisterCccsrType
	Ccvr    RegisterCcvrType
	Cccr    RegisterCccrType
	Pwrcr   RegisterPwrcrType
	_       [244]byte
	Pkgr    RegisterPkgrType
	_       [472]byte
	Ur0     RegisterUr0Type
	Ur1     RegisterUr1Type
	Ur2     RegisterUr2Type
	Ur3     RegisterUr3Type
	Ur4     RegisterUr4Type
	Ur5     RegisterUr5Type
	Ur6     RegisterUr6Type
	Ur7     RegisterUr7Type
	Ur8     RegisterUr8Type
	Ur9     RegisterUr9Type
	Ur10    RegisterUr10Type
	Ur11    RegisterUr11Type
	Ur12    RegisterUr12Type
	Ur13    RegisterUr13Type
	Ur14    RegisterUr14Type
	Ur15    RegisterUr15Type
	Ur16    RegisterUr16Type
	Ur17    RegisterUr17Type
}

// RegisterPmcrType peripheral mode configuration register
type RegisterPmcrType uint32

func (r *RegisterPmcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPmcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPmcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPmcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPmcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPmcrFieldI2c1fmpShift = 0
	RegisterPmcrFieldI2c1fmpMask  = 0x1
)

// GetI2c1fmp I2C1 Fm+
func (r *RegisterPmcrType) GetI2c1fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldI2c1fmpMask) != 0
}

// SetI2c1fmp I2C1 Fm+
func (r *RegisterPmcrType) SetI2c1fmp(value bool) {
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
func (r *RegisterPmcrType) GetI2c2fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldI2c2fmpMask) != 0
}

// SetI2c2fmp I2C2 Fm+
func (r *RegisterPmcrType) SetI2c2fmp(value bool) {
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
func (r *RegisterPmcrType) GetI2c3fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldI2c3fmpMask) != 0
}

// SetI2c3fmp I2C3 Fm+
func (r *RegisterPmcrType) SetI2c3fmp(value bool) {
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
func (r *RegisterPmcrType) GetI2c4fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldI2c4fmpMask) != 0
}

// SetI2c4fmp I2C4 Fm+
func (r *RegisterPmcrType) SetI2c4fmp(value bool) {
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
func (r *RegisterPmcrType) GetPb6fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPb6fmpMask) != 0
}

// SetPb6fmp PB(6) Fm+
func (r *RegisterPmcrType) SetPb6fmp(value bool) {
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
func (r *RegisterPmcrType) GetPb7fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPb7fmpMask) != 0
}

// SetPb7fmp PB(7) Fast Mode Plus
func (r *RegisterPmcrType) SetPb7fmp(value bool) {
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
func (r *RegisterPmcrType) GetPb8fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPb8fmpMask) != 0
}

// SetPb8fmp PB(8) Fast Mode Plus
func (r *RegisterPmcrType) SetPb8fmp(value bool) {
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
func (r *RegisterPmcrType) GetPb9fmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPb9fmpMask) != 0
}

// SetPb9fmp PB(9) Fm+
func (r *RegisterPmcrType) SetPb9fmp(value bool) {
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
func (r *RegisterPmcrType) GetBooste() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldBoosteMask) != 0
}

// SetBooste Booster Enable
func (r *RegisterPmcrType) SetBooste(value bool) {
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
func (r *RegisterPmcrType) GetBoostvddsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldBoostvddselMask) != 0
}

// SetBoostvddsel Analog switch supply voltage selection
func (r *RegisterPmcrType) SetBoostvddsel(value bool) {
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
func (r *RegisterPmcrType) GetEpis() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldEpisMask) >> RegisterPmcrFieldEpisShift)
}

// SetEpis Ethernet PHY Interface Selection
func (r *RegisterPmcrType) SetEpis(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldEpisMask)|(uint32(value)<<RegisterPmcrFieldEpisShift))
}

const (
	RegisterPmcrFieldPa0soShift = 24
	RegisterPmcrFieldPa0soMask  = 0x1000000
)

// GetPa0so PA0 Switch Open
func (r *RegisterPmcrType) GetPa0so() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPa0soMask) != 0
}

// SetPa0so PA0 Switch Open
func (r *RegisterPmcrType) SetPa0so(value bool) {
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
func (r *RegisterPmcrType) GetPa1so() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPa1soMask) != 0
}

// SetPa1so PA1 Switch Open
func (r *RegisterPmcrType) SetPa1so(value bool) {
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
func (r *RegisterPmcrType) GetPc2so() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPc2soMask) != 0
}

// SetPc2so PC2 Switch Open
func (r *RegisterPmcrType) SetPc2so(value bool) {
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
func (r *RegisterPmcrType) GetPc3so() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPmcrFieldPc3soMask) != 0
}

// SetPc3so PC3 Switch Open
func (r *RegisterPmcrType) SetPc3so(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPmcrFieldPc3soMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPmcrFieldPc3soMask)
	}
}

// RegisterExticr1Type external interrupt configuration register 1
type RegisterExticr1Type uint32

func (r *RegisterExticr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterExticr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterExticr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterExticr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterExticr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterExticr1FieldExti0Shift = 0
	RegisterExticr1FieldExti0Mask  = 0xf
)

// GetExti0 EXTI x configuration (x = 0 to 3)
func (r *RegisterExticr1Type) GetExti0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr1FieldExti0Mask) >> RegisterExticr1FieldExti0Shift)
}

// SetExti0 EXTI x configuration (x = 0 to 3)
func (r *RegisterExticr1Type) SetExti0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr1FieldExti0Mask)|(uint32(value)<<RegisterExticr1FieldExti0Shift))
}

const (
	RegisterExticr1FieldExti1Shift = 4
	RegisterExticr1FieldExti1Mask  = 0xf0
)

// GetExti1 EXTI x configuration (x = 0 to 3)
func (r *RegisterExticr1Type) GetExti1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr1FieldExti1Mask) >> RegisterExticr1FieldExti1Shift)
}

// SetExti1 EXTI x configuration (x = 0 to 3)
func (r *RegisterExticr1Type) SetExti1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr1FieldExti1Mask)|(uint32(value)<<RegisterExticr1FieldExti1Shift))
}

const (
	RegisterExticr1FieldExti2Shift = 8
	RegisterExticr1FieldExti2Mask  = 0xf00
)

// GetExti2 EXTI x configuration (x = 0 to 3)
func (r *RegisterExticr1Type) GetExti2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr1FieldExti2Mask) >> RegisterExticr1FieldExti2Shift)
}

// SetExti2 EXTI x configuration (x = 0 to 3)
func (r *RegisterExticr1Type) SetExti2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr1FieldExti2Mask)|(uint32(value)<<RegisterExticr1FieldExti2Shift))
}

const (
	RegisterExticr1FieldExti3Shift = 12
	RegisterExticr1FieldExti3Mask  = 0xf000
)

// GetExti3 EXTI x configuration (x = 0 to 3)
func (r *RegisterExticr1Type) GetExti3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr1FieldExti3Mask) >> RegisterExticr1FieldExti3Shift)
}

// SetExti3 EXTI x configuration (x = 0 to 3)
func (r *RegisterExticr1Type) SetExti3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr1FieldExti3Mask)|(uint32(value)<<RegisterExticr1FieldExti3Shift))
}

// RegisterExticr2Type external interrupt configuration register 2
type RegisterExticr2Type uint32

func (r *RegisterExticr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterExticr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterExticr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterExticr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterExticr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterExticr2FieldExti4Shift = 0
	RegisterExticr2FieldExti4Mask  = 0xf
)

// GetExti4 EXTI x configuration (x = 4 to 7)
func (r *RegisterExticr2Type) GetExti4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr2FieldExti4Mask) >> RegisterExticr2FieldExti4Shift)
}

// SetExti4 EXTI x configuration (x = 4 to 7)
func (r *RegisterExticr2Type) SetExti4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr2FieldExti4Mask)|(uint32(value)<<RegisterExticr2FieldExti4Shift))
}

const (
	RegisterExticr2FieldExti5Shift = 4
	RegisterExticr2FieldExti5Mask  = 0xf0
)

// GetExti5 EXTI x configuration (x = 4 to 7)
func (r *RegisterExticr2Type) GetExti5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr2FieldExti5Mask) >> RegisterExticr2FieldExti5Shift)
}

// SetExti5 EXTI x configuration (x = 4 to 7)
func (r *RegisterExticr2Type) SetExti5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr2FieldExti5Mask)|(uint32(value)<<RegisterExticr2FieldExti5Shift))
}

const (
	RegisterExticr2FieldExti6Shift = 8
	RegisterExticr2FieldExti6Mask  = 0xf00
)

// GetExti6 EXTI x configuration (x = 4 to 7)
func (r *RegisterExticr2Type) GetExti6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr2FieldExti6Mask) >> RegisterExticr2FieldExti6Shift)
}

// SetExti6 EXTI x configuration (x = 4 to 7)
func (r *RegisterExticr2Type) SetExti6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr2FieldExti6Mask)|(uint32(value)<<RegisterExticr2FieldExti6Shift))
}

const (
	RegisterExticr2FieldExti7Shift = 12
	RegisterExticr2FieldExti7Mask  = 0xf000
)

// GetExti7 EXTI x configuration (x = 4 to 7)
func (r *RegisterExticr2Type) GetExti7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr2FieldExti7Mask) >> RegisterExticr2FieldExti7Shift)
}

// SetExti7 EXTI x configuration (x = 4 to 7)
func (r *RegisterExticr2Type) SetExti7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr2FieldExti7Mask)|(uint32(value)<<RegisterExticr2FieldExti7Shift))
}

// RegisterExticr3Type external interrupt configuration register 3
type RegisterExticr3Type uint32

func (r *RegisterExticr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterExticr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterExticr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterExticr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterExticr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterExticr3FieldExti8Shift = 0
	RegisterExticr3FieldExti8Mask  = 0xf
)

// GetExti8 EXTI x configuration (x = 8 to 11)
func (r *RegisterExticr3Type) GetExti8() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr3FieldExti8Mask) >> RegisterExticr3FieldExti8Shift)
}

// SetExti8 EXTI x configuration (x = 8 to 11)
func (r *RegisterExticr3Type) SetExti8(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr3FieldExti8Mask)|(uint32(value)<<RegisterExticr3FieldExti8Shift))
}

const (
	RegisterExticr3FieldExti9Shift = 4
	RegisterExticr3FieldExti9Mask  = 0xf0
)

// GetExti9 EXTI x configuration (x = 8 to 11)
func (r *RegisterExticr3Type) GetExti9() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr3FieldExti9Mask) >> RegisterExticr3FieldExti9Shift)
}

// SetExti9 EXTI x configuration (x = 8 to 11)
func (r *RegisterExticr3Type) SetExti9(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr3FieldExti9Mask)|(uint32(value)<<RegisterExticr3FieldExti9Shift))
}

const (
	RegisterExticr3FieldExti10Shift = 8
	RegisterExticr3FieldExti10Mask  = 0xf00
)

// GetExti10 EXTI10
func (r *RegisterExticr3Type) GetExti10() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr3FieldExti10Mask) >> RegisterExticr3FieldExti10Shift)
}

// SetExti10 EXTI10
func (r *RegisterExticr3Type) SetExti10(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr3FieldExti10Mask)|(uint32(value)<<RegisterExticr3FieldExti10Shift))
}

const (
	RegisterExticr3FieldExti11Shift = 12
	RegisterExticr3FieldExti11Mask  = 0xf000
)

// GetExti11 EXTI x configuration (x = 8 to 11)
func (r *RegisterExticr3Type) GetExti11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr3FieldExti11Mask) >> RegisterExticr3FieldExti11Shift)
}

// SetExti11 EXTI x configuration (x = 8 to 11)
func (r *RegisterExticr3Type) SetExti11(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr3FieldExti11Mask)|(uint32(value)<<RegisterExticr3FieldExti11Shift))
}

// RegisterExticr4Type external interrupt configuration register 4
type RegisterExticr4Type uint32

func (r *RegisterExticr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterExticr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterExticr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterExticr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterExticr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterExticr4FieldExti12Shift = 0
	RegisterExticr4FieldExti12Mask  = 0xf
)

// GetExti12 EXTI x configuration (x = 12 to 15)
func (r *RegisterExticr4Type) GetExti12() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr4FieldExti12Mask) >> RegisterExticr4FieldExti12Shift)
}

// SetExti12 EXTI x configuration (x = 12 to 15)
func (r *RegisterExticr4Type) SetExti12(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr4FieldExti12Mask)|(uint32(value)<<RegisterExticr4FieldExti12Shift))
}

const (
	RegisterExticr4FieldExti13Shift = 4
	RegisterExticr4FieldExti13Mask  = 0xf0
)

// GetExti13 EXTI x configuration (x = 12 to 15)
func (r *RegisterExticr4Type) GetExti13() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr4FieldExti13Mask) >> RegisterExticr4FieldExti13Shift)
}

// SetExti13 EXTI x configuration (x = 12 to 15)
func (r *RegisterExticr4Type) SetExti13(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr4FieldExti13Mask)|(uint32(value)<<RegisterExticr4FieldExti13Shift))
}

const (
	RegisterExticr4FieldExti14Shift = 8
	RegisterExticr4FieldExti14Mask  = 0xf00
)

// GetExti14 EXTI x configuration (x = 12 to 15)
func (r *RegisterExticr4Type) GetExti14() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr4FieldExti14Mask) >> RegisterExticr4FieldExti14Shift)
}

// SetExti14 EXTI x configuration (x = 12 to 15)
func (r *RegisterExticr4Type) SetExti14(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr4FieldExti14Mask)|(uint32(value)<<RegisterExticr4FieldExti14Shift))
}

const (
	RegisterExticr4FieldExti15Shift = 12
	RegisterExticr4FieldExti15Mask  = 0xf000
)

// GetExti15 EXTI x configuration (x = 12 to 15)
func (r *RegisterExticr4Type) GetExti15() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterExticr4FieldExti15Mask) >> RegisterExticr4FieldExti15Shift)
}

// SetExti15 EXTI x configuration (x = 12 to 15)
func (r *RegisterExticr4Type) SetExti15(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterExticr4FieldExti15Mask)|(uint32(value)<<RegisterExticr4FieldExti15Shift))
}

// RegisterCfgrType configuration register
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
	RegisterCfgrFieldCm4lShift = 0
	RegisterCfgrFieldCm4lMask  = 0x1
)

// GetCm4l CM4L
func (r *RegisterCfgrType) GetCm4l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldCm4lMask) != 0
}

// SetCm4l CM4L
func (r *RegisterCfgrType) SetCm4l(value bool) {
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
func (r *RegisterCfgrType) GetPvdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldPvdlMask) != 0
}

// SetPvdl PVDL
func (r *RegisterCfgrType) SetPvdl(value bool) {
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
func (r *RegisterCfgrType) GetFlashl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldFlashlMask) != 0
}

// SetFlashl FLASHL
func (r *RegisterCfgrType) SetFlashl(value bool) {
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
func (r *RegisterCfgrType) GetCm7l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldCm7lMask) != 0
}

// SetCm7l CM7L
func (r *RegisterCfgrType) SetCm7l(value bool) {
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
func (r *RegisterCfgrType) GetBkraml() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldBkramlMask) != 0
}

// SetBkraml BKRAML
func (r *RegisterCfgrType) SetBkraml(value bool) {
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
func (r *RegisterCfgrType) GetSram4l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSram4lMask) != 0
}

// SetSram4l SRAM4L
func (r *RegisterCfgrType) SetSram4l(value bool) {
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
func (r *RegisterCfgrType) GetSram3l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSram3lMask) != 0
}

// SetSram3l SRAM3L
func (r *RegisterCfgrType) SetSram3l(value bool) {
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
func (r *RegisterCfgrType) GetSram2l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSram2lMask) != 0
}

// SetSram2l SRAM2L
func (r *RegisterCfgrType) SetSram2l(value bool) {
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
func (r *RegisterCfgrType) GetSram1l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSram1lMask) != 0
}

// SetSram1l SRAM1L
func (r *RegisterCfgrType) SetSram1l(value bool) {
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
func (r *RegisterCfgrType) GetDtcml() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldDtcmlMask) != 0
}

// SetDtcml DTCML
func (r *RegisterCfgrType) SetDtcml(value bool) {
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
func (r *RegisterCfgrType) GetItcml() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldItcmlMask) != 0
}

// SetItcml ITCML
func (r *RegisterCfgrType) SetItcml(value bool) {
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
func (r *RegisterCfgrType) GetAxisraml() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldAxisramlMask) != 0
}

// SetAxisraml AXISRAML
func (r *RegisterCfgrType) SetAxisraml(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldAxisramlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldAxisramlMask)
	}
}

// RegisterCccsrType compensation cell control/status register
type RegisterCccsrType uint32

func (r *RegisterCccsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCccsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCccsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCccsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCccsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCccsrFieldEnShift = 0
	RegisterCccsrFieldEnMask  = 0x1
)

// GetEn enable
func (r *RegisterCccsrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCccsrFieldEnMask) != 0
}

// SetEn enable
func (r *RegisterCccsrType) SetEn(value bool) {
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
func (r *RegisterCccsrType) GetCs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCccsrFieldCsMask) != 0
}

// SetCs Code selection
func (r *RegisterCccsrType) SetCs(value bool) {
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
func (r *RegisterCccsrType) GetReady() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCccsrFieldReadyMask) != 0
}

// SetReady Compensation cell ready flag
func (r *RegisterCccsrType) SetReady(value bool) {
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
func (r *RegisterCccsrType) GetHslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCccsrFieldHslvMask) != 0
}

// SetHslv High-speed at low-voltage
func (r *RegisterCccsrType) SetHslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCccsrFieldHslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCccsrFieldHslvMask)
	}
}

// RegisterCcvrType SYSCFG compensation cell value register
type RegisterCcvrType uint32

func (r *RegisterCcvrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcvrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcvrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcvrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcvrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcvrFieldNcvShift = 0
	RegisterCcvrFieldNcvMask  = 0xf
)

// GetNcv NMOS compensation value
func (r *RegisterCcvrType) GetNcv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcvrFieldNcvMask) >> RegisterCcvrFieldNcvShift)
}

// SetNcv NMOS compensation value
func (r *RegisterCcvrType) SetNcv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcvrFieldNcvMask)|(uint32(value)<<RegisterCcvrFieldNcvShift))
}

const (
	RegisterCcvrFieldPcvShift = 4
	RegisterCcvrFieldPcvMask  = 0xf0
)

// GetPcv PMOS compensation value
func (r *RegisterCcvrType) GetPcv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcvrFieldPcvMask) >> RegisterCcvrFieldPcvShift)
}

// SetPcv PMOS compensation value
func (r *RegisterCcvrType) SetPcv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcvrFieldPcvMask)|(uint32(value)<<RegisterCcvrFieldPcvShift))
}

// RegisterCccrType SYSCFG compensation cell code register
type RegisterCccrType uint32

func (r *RegisterCccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCccrFieldNccShift = 0
	RegisterCccrFieldNccMask  = 0xf
)

// GetNcc NMOS compensation code
func (r *RegisterCccrType) GetNcc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCccrFieldNccMask) >> RegisterCccrFieldNccShift)
}

// SetNcc NMOS compensation code
func (r *RegisterCccrType) SetNcc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCccrFieldNccMask)|(uint32(value)<<RegisterCccrFieldNccShift))
}

const (
	RegisterCccrFieldPccShift = 4
	RegisterCccrFieldPccMask  = 0xf0
)

// GetPcc PMOS compensation code
func (r *RegisterCccrType) GetPcc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCccrFieldPccMask) >> RegisterCccrFieldPccShift)
}

// SetPcc PMOS compensation code
func (r *RegisterCccrType) SetPcc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCccrFieldPccMask)|(uint32(value)<<RegisterCccrFieldPccShift))
}

// RegisterPwrcrType SYSCFG power control register
type RegisterPwrcrType uint32

func (r *RegisterPwrcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPwrcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPwrcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPwrcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPwrcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPwrcrFieldOdenShift = 0
	RegisterPwrcrFieldOdenMask  = 0x1
)

// GetOden Overdrive enable
func (r *RegisterPwrcrType) GetOden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPwrcrFieldOdenMask) != 0
}

// SetOden Overdrive enable
func (r *RegisterPwrcrType) SetOden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPwrcrFieldOdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPwrcrFieldOdenMask)
	}
}

// RegisterPkgrType SYSCFG package register
type RegisterPkgrType uint32

func (r *RegisterPkgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPkgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPkgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPkgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPkgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPkgrFieldPkgShift = 0
	RegisterPkgrFieldPkgMask  = 0xf
)

// GetPkg Package
func (r *RegisterPkgrType) GetPkg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPkgrFieldPkgMask) >> RegisterPkgrFieldPkgShift)
}

// SetPkg Package
func (r *RegisterPkgrType) SetPkg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPkgrFieldPkgMask)|(uint32(value)<<RegisterPkgrFieldPkgShift))
}

// RegisterUr0Type SYSCFG user register 0
type RegisterUr0Type uint32

func (r *RegisterUr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr0FieldBksShift = 0
	RegisterUr0FieldBksMask  = 0x1
)

// GetBks Bank Swap
func (r *RegisterUr0Type) GetBks() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr0FieldBksMask) != 0
}

// SetBks Bank Swap
func (r *RegisterUr0Type) SetBks(value bool) {
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
func (r *RegisterUr0Type) GetRdp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterUr0FieldRdpMask) >> RegisterUr0FieldRdpShift)
}

// SetRdp Readout protection
func (r *RegisterUr0Type) SetRdp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr0FieldRdpMask)|(uint32(value)<<RegisterUr0FieldRdpShift))
}

// RegisterUr1Type SYSCFG user register 1
type RegisterUr1Type uint32

func (r *RegisterUr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr1FieldBcm4Shift = 0
	RegisterUr1FieldBcm4Mask  = 0x1
)

// GetBcm4 Boot Cortex-M4
func (r *RegisterUr1Type) GetBcm4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr1FieldBcm4Mask) != 0
}

// SetBcm4 Boot Cortex-M4
func (r *RegisterUr1Type) SetBcm4(value bool) {
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
func (r *RegisterUr1Type) GetBcm7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr1FieldBcm7Mask) != 0
}

// SetBcm7 Boot Cortex-M7
func (r *RegisterUr1Type) SetBcm7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr1FieldBcm7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr1FieldBcm7Mask)
	}
}

// RegisterUr2Type SYSCFG user register 2
type RegisterUr2Type uint32

func (r *RegisterUr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr2FieldBorhShift = 0
	RegisterUr2FieldBorhMask  = 0x3
)

// GetBorh BOR_LVL Brownout Reset Threshold Level
func (r *RegisterUr2Type) GetBorh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterUr2FieldBorhMask) >> RegisterUr2FieldBorhShift)
}

const (
	RegisterUr2FieldBcm7add0Shift = 16
	RegisterUr2FieldBcm7add0Mask  = 0xffff0000
)

// GetBcm7add0 Cortex-M7 Boot Address 0
func (r *RegisterUr2Type) GetBcm7add0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr2FieldBcm7add0Mask) >> RegisterUr2FieldBcm7add0Shift)
}

// SetBcm7add0 Cortex-M7 Boot Address 0
func (r *RegisterUr2Type) SetBcm7add0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr2FieldBcm7add0Mask)|(uint32(value)<<RegisterUr2FieldBcm7add0Shift))
}

// RegisterUr3Type SYSCFG user register 3
type RegisterUr3Type uint32

func (r *RegisterUr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr3FieldBcm4add1Shift = 0
	RegisterUr3FieldBcm4add1Mask  = 0xffff
)

// GetBcm4add1 Cortex-M4 Boot Address 0
func (r *RegisterUr3Type) GetBcm4add1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr3FieldBcm4add1Mask) >> RegisterUr3FieldBcm4add1Shift)
}

// SetBcm4add1 Cortex-M4 Boot Address 0
func (r *RegisterUr3Type) SetBcm4add1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr3FieldBcm4add1Mask)|(uint32(value)<<RegisterUr3FieldBcm4add1Shift))
}

const (
	RegisterUr3FieldBcm7add1Shift = 16
	RegisterUr3FieldBcm7add1Mask  = 0xffff0000
)

// GetBcm7add1 Cortex-M7 Boot Address 1
func (r *RegisterUr3Type) GetBcm7add1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr3FieldBcm7add1Mask) >> RegisterUr3FieldBcm7add1Shift)
}

// SetBcm7add1 Cortex-M7 Boot Address 1
func (r *RegisterUr3Type) SetBcm7add1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr3FieldBcm7add1Mask)|(uint32(value)<<RegisterUr3FieldBcm7add1Shift))
}

// RegisterUr4Type SYSCFG user register 4
type RegisterUr4Type uint32

func (r *RegisterUr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr4FieldBcm4add1Shift = 0
	RegisterUr4FieldBcm4add1Mask  = 0xffff
)

// GetBcm4add1 Mass Erase Protected Area Disabled for bank 1
func (r *RegisterUr4Type) GetBcm4add1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr4FieldBcm4add1Mask) >> RegisterUr4FieldBcm4add1Shift)
}

// SetBcm4add1 Mass Erase Protected Area Disabled for bank 1
func (r *RegisterUr4Type) SetBcm4add1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr4FieldBcm4add1Mask)|(uint32(value)<<RegisterUr4FieldBcm4add1Shift))
}

const (
	RegisterUr4FieldMepad1Shift = 16
	RegisterUr4FieldMepad1Mask  = 0x10000
)

// GetMepad1 Boot Cortex-M4 Address 1
func (r *RegisterUr4Type) GetMepad1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr4FieldMepad1Mask) != 0
}

// RegisterUr5Type SYSCFG user register 5
type RegisterUr5Type uint32

func (r *RegisterUr5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr5FieldMesad1Shift = 0
	RegisterUr5FieldMesad1Mask  = 0x1
)

// GetMesad1 Mass erase secured area disabled for bank 1
func (r *RegisterUr5Type) GetMesad1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr5FieldMesad1Mask) != 0
}

// SetMesad1 Mass erase secured area disabled for bank 1
func (r *RegisterUr5Type) SetMesad1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr5FieldMesad1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr5FieldMesad1Mask)
	}
}

const (
	RegisterUr5FieldWrps1Shift = 16
	RegisterUr5FieldWrps1Mask  = 0xff0000
)

// GetWrps1 Write protection for flash bank 1
func (r *RegisterUr5Type) GetWrps1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterUr5FieldWrps1Mask) >> RegisterUr5FieldWrps1Shift)
}

// SetWrps1 Write protection for flash bank 1
func (r *RegisterUr5Type) SetWrps1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr5FieldWrps1Mask)|(uint32(value)<<RegisterUr5FieldWrps1Shift))
}

// RegisterUr6Type SYSCFG user register 6
type RegisterUr6Type uint32

func (r *RegisterUr6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr6FieldPabeg1Shift = 0
	RegisterUr6FieldPabeg1Mask  = 0xfff
)

// GetPabeg1 Protected area start address for bank 1
func (r *RegisterUr6Type) GetPabeg1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr6FieldPabeg1Mask) >> RegisterUr6FieldPabeg1Shift)
}

// SetPabeg1 Protected area start address for bank 1
func (r *RegisterUr6Type) SetPabeg1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr6FieldPabeg1Mask)|(uint32(value)<<RegisterUr6FieldPabeg1Shift))
}

const (
	RegisterUr6FieldPaend1Shift = 16
	RegisterUr6FieldPaend1Mask  = 0xfff0000
)

// GetPaend1 Protected area end address for bank 1
func (r *RegisterUr6Type) GetPaend1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr6FieldPaend1Mask) >> RegisterUr6FieldPaend1Shift)
}

// SetPaend1 Protected area end address for bank 1
func (r *RegisterUr6Type) SetPaend1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr6FieldPaend1Mask)|(uint32(value)<<RegisterUr6FieldPaend1Shift))
}

// RegisterUr7Type SYSCFG user register 7
type RegisterUr7Type uint32

func (r *RegisterUr7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr7FieldSabeg1Shift = 0
	RegisterUr7FieldSabeg1Mask  = 0xfff
)

// GetSabeg1 Secured area start address for bank 1
func (r *RegisterUr7Type) GetSabeg1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr7FieldSabeg1Mask) >> RegisterUr7FieldSabeg1Shift)
}

// SetSabeg1 Secured area start address for bank 1
func (r *RegisterUr7Type) SetSabeg1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr7FieldSabeg1Mask)|(uint32(value)<<RegisterUr7FieldSabeg1Shift))
}

const (
	RegisterUr7FieldSaend1Shift = 16
	RegisterUr7FieldSaend1Mask  = 0xfff0000
)

// GetSaend1 Secured area end address for bank 1
func (r *RegisterUr7Type) GetSaend1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr7FieldSaend1Mask) >> RegisterUr7FieldSaend1Shift)
}

// SetSaend1 Secured area end address for bank 1
func (r *RegisterUr7Type) SetSaend1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr7FieldSaend1Mask)|(uint32(value)<<RegisterUr7FieldSaend1Shift))
}

// RegisterUr8Type SYSCFG user register 8
type RegisterUr8Type uint32

func (r *RegisterUr8Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr8Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr8Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr8Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr8Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr8FieldMepad2Shift = 0
	RegisterUr8FieldMepad2Mask  = 0x1
)

// GetMepad2 Mass erase protected area disabled for bank 2
func (r *RegisterUr8Type) GetMepad2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr8FieldMepad2Mask) != 0
}

// SetMepad2 Mass erase protected area disabled for bank 2
func (r *RegisterUr8Type) SetMepad2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr8FieldMepad2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr8FieldMepad2Mask)
	}
}

const (
	RegisterUr8FieldMesad2Shift = 16
	RegisterUr8FieldMesad2Mask  = 0x10000
)

// GetMesad2 Mass erase secured area disabled for bank 2
func (r *RegisterUr8Type) GetMesad2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr8FieldMesad2Mask) != 0
}

// SetMesad2 Mass erase secured area disabled for bank 2
func (r *RegisterUr8Type) SetMesad2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr8FieldMesad2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr8FieldMesad2Mask)
	}
}

// RegisterUr9Type SYSCFG user register 9
type RegisterUr9Type uint32

func (r *RegisterUr9Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr9Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr9Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr9Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr9Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr9FieldWrps2Shift = 0
	RegisterUr9FieldWrps2Mask  = 0xff
)

// GetWrps2 Write protection for flash bank 2
func (r *RegisterUr9Type) GetWrps2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterUr9FieldWrps2Mask) >> RegisterUr9FieldWrps2Shift)
}

// SetWrps2 Write protection for flash bank 2
func (r *RegisterUr9Type) SetWrps2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr9FieldWrps2Mask)|(uint32(value)<<RegisterUr9FieldWrps2Shift))
}

const (
	RegisterUr9FieldPabeg2Shift = 16
	RegisterUr9FieldPabeg2Mask  = 0xfff0000
)

// GetPabeg2 Protected area start address for bank 2
func (r *RegisterUr9Type) GetPabeg2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr9FieldPabeg2Mask) >> RegisterUr9FieldPabeg2Shift)
}

// SetPabeg2 Protected area start address for bank 2
func (r *RegisterUr9Type) SetPabeg2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr9FieldPabeg2Mask)|(uint32(value)<<RegisterUr9FieldPabeg2Shift))
}

// RegisterUr10Type SYSCFG user register 10
type RegisterUr10Type uint32

func (r *RegisterUr10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr10FieldPaend2Shift = 0
	RegisterUr10FieldPaend2Mask  = 0xfff
)

// GetPaend2 Protected area end address for bank 2
func (r *RegisterUr10Type) GetPaend2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr10FieldPaend2Mask) >> RegisterUr10FieldPaend2Shift)
}

// SetPaend2 Protected area end address for bank 2
func (r *RegisterUr10Type) SetPaend2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr10FieldPaend2Mask)|(uint32(value)<<RegisterUr10FieldPaend2Shift))
}

const (
	RegisterUr10FieldSabeg2Shift = 16
	RegisterUr10FieldSabeg2Mask  = 0xfff0000
)

// GetSabeg2 Secured area start address for bank 2
func (r *RegisterUr10Type) GetSabeg2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr10FieldSabeg2Mask) >> RegisterUr10FieldSabeg2Shift)
}

// SetSabeg2 Secured area start address for bank 2
func (r *RegisterUr10Type) SetSabeg2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr10FieldSabeg2Mask)|(uint32(value)<<RegisterUr10FieldSabeg2Shift))
}

// RegisterUr11Type SYSCFG user register 11
type RegisterUr11Type uint32

func (r *RegisterUr11Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr11Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr11Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr11Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr11Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr11FieldSaend2Shift = 0
	RegisterUr11FieldSaend2Mask  = 0xfff
)

// GetSaend2 Secured area end address for bank 2
func (r *RegisterUr11Type) GetSaend2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterUr11FieldSaend2Mask) >> RegisterUr11FieldSaend2Shift)
}

// SetSaend2 Secured area end address for bank 2
func (r *RegisterUr11Type) SetSaend2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr11FieldSaend2Mask)|(uint32(value)<<RegisterUr11FieldSaend2Shift))
}

const (
	RegisterUr11FieldIwdg1mShift = 16
	RegisterUr11FieldIwdg1mMask  = 0x10000
)

// GetIwdg1m Independent Watchdog 1 mode
func (r *RegisterUr11Type) GetIwdg1m() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr11FieldIwdg1mMask) != 0
}

// SetIwdg1m Independent Watchdog 1 mode
func (r *RegisterUr11Type) SetIwdg1m(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr11FieldIwdg1mMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr11FieldIwdg1mMask)
	}
}

// RegisterUr12Type SYSCFG user register 12
type RegisterUr12Type uint32

func (r *RegisterUr12Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr12Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr12Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr12Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr12Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr12FieldIwdg2mShift = 0
	RegisterUr12FieldIwdg2mMask  = 0x1
)

// GetIwdg2m Independent Watchdog 2 mode
func (r *RegisterUr12Type) GetIwdg2m() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr12FieldIwdg2mMask) != 0
}

// SetIwdg2m Independent Watchdog 2 mode
func (r *RegisterUr12Type) SetIwdg2m(value bool) {
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
func (r *RegisterUr12Type) GetSecure() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr12FieldSecureMask) != 0
}

// SetSecure Secure mode
func (r *RegisterUr12Type) SetSecure(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr12FieldSecureMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr12FieldSecureMask)
	}
}

// RegisterUr13Type SYSCFG user register 13
type RegisterUr13Type uint32

func (r *RegisterUr13Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr13Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr13Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr13Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr13Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr13FieldSdrsShift = 0
	RegisterUr13FieldSdrsMask  = 0x3
)

// GetSdrs Secured DTCM RAM Size
func (r *RegisterUr13Type) GetSdrs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterUr13FieldSdrsMask) >> RegisterUr13FieldSdrsShift)
}

// SetSdrs Secured DTCM RAM Size
func (r *RegisterUr13Type) SetSdrs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUr13FieldSdrsMask)|(uint32(value)<<RegisterUr13FieldSdrsShift))
}

const (
	RegisterUr13FieldD1sbrstShift = 16
	RegisterUr13FieldD1sbrstMask  = 0x10000
)

// GetD1sbrst D1 Standby reset
func (r *RegisterUr13Type) GetD1sbrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr13FieldD1sbrstMask) != 0
}

// SetD1sbrst D1 Standby reset
func (r *RegisterUr13Type) SetD1sbrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr13FieldD1sbrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr13FieldD1sbrstMask)
	}
}

// RegisterUr14Type SYSCFG user register 14
type RegisterUr14Type uint32

func (r *RegisterUr14Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr14Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr14Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr14Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr14Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr14FieldD1stprstShift = 0
	RegisterUr14FieldD1stprstMask  = 0x1
)

// GetD1stprst D1 Stop Reset
func (r *RegisterUr14Type) GetD1stprst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr14FieldD1stprstMask) != 0
}

// SetD1stprst D1 Stop Reset
func (r *RegisterUr14Type) SetD1stprst(value bool) {
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
func (r *RegisterUr14Type) GetD2sbrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr14FieldD2sbrstMask) != 0
}

// SetD2sbrst D2 Standby Reset
func (r *RegisterUr14Type) SetD2sbrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr14FieldD2sbrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr14FieldD2sbrstMask)
	}
}

// RegisterUr15Type SYSCFG user register 15
type RegisterUr15Type uint32

func (r *RegisterUr15Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr15Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr15Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr15Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr15Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr15FieldD2stprstShift = 0
	RegisterUr15FieldD2stprstMask  = 0x1
)

// GetD2stprst D2 Stop Reset
func (r *RegisterUr15Type) GetD2stprst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr15FieldD2stprstMask) != 0
}

// SetD2stprst D2 Stop Reset
func (r *RegisterUr15Type) SetD2stprst(value bool) {
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
func (r *RegisterUr15Type) GetFziwdgstb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr15FieldFziwdgstbMask) != 0
}

// RegisterUr16Type SYSCFG user register 16
type RegisterUr16Type uint32

func (r *RegisterUr16Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr16Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr16Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr16Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr16Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr16FieldFziwdgstpShift = 0
	RegisterUr16FieldFziwdgstpMask  = 0x1
)

// GetFziwdgstp Freeze independent watchdog in Stop mode
func (r *RegisterUr16Type) GetFziwdgstp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr16FieldFziwdgstpMask) != 0
}

// SetFziwdgstp Freeze independent watchdog in Stop mode
func (r *RegisterUr16Type) SetFziwdgstp(value bool) {
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
func (r *RegisterUr16Type) GetPkp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr16FieldPkpMask) != 0
}

// SetPkp Private key programmed
func (r *RegisterUr16Type) SetPkp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr16FieldPkpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr16FieldPkpMask)
	}
}

// RegisterUr17Type SYSCFG user register 17
type RegisterUr17Type uint32

func (r *RegisterUr17Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUr17Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUr17Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUr17Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUr17Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUr17FieldIohslvShift = 0
	RegisterUr17FieldIohslvMask  = 0x1
)

// GetIohslv I/O high speed / low voltage
func (r *RegisterUr17Type) GetIohslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUr17FieldIohslvMask) != 0
}

// SetIohslv I/O high speed / low voltage
func (r *RegisterUr17Type) SetIohslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUr17FieldIohslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUr17FieldIohslvMask)
	}
}
