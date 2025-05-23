package gpioa

import (
	"unsafe"
	"volatile"
)

var (
	Gpioa = (*_gpioa)(unsafe.Pointer(uintptr(0x58020000)))
	Gpiob = (*_gpioa)(unsafe.Pointer(uintptr(0x58020400)))
	Gpioc = (*_gpioa)(unsafe.Pointer(uintptr(0x58020800)))
	Gpiod = (*_gpioa)(unsafe.Pointer(uintptr(0x58020c00)))
	Gpioe = (*_gpioa)(unsafe.Pointer(uintptr(0x58021000)))
	Gpiof = (*_gpioa)(unsafe.Pointer(uintptr(0x58021400)))
	Gpiog = (*_gpioa)(unsafe.Pointer(uintptr(0x58021800)))
	Gpioh = (*_gpioa)(unsafe.Pointer(uintptr(0x58021c00)))
	Gpioi = (*_gpioa)(unsafe.Pointer(uintptr(0x58022000)))
	Gpioj = (*_gpioa)(unsafe.Pointer(uintptr(0x58022400)))
	Gpiok = (*_gpioa)(unsafe.Pointer(uintptr(0x58022800)))
)

type _gpioa struct {
	Moder   registerModerType
	Otyper  registerOtyperType
	Ospeedr registerOspeedrType
	Pupdr   registerPupdrType
	Idr     registerIdrType
	Odr     registerOdrType
	Bsrr    registerBsrrType
	Lckr    registerLckrType
	Afrl    registerAfrlType
	Afrh    registerAfrhType
}

// registerModerType GPIO port mode register
type registerModerType uint32

const (
	RegisterModerFieldMode0Shift = 0
	RegisterModerFieldMode0Mask  = 0x3
)

// GetMode0 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode0Mask) >> RegisterModerFieldMode0Shift)
}

// SetMode0 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode0Mask)|(uint32(value)<<RegisterModerFieldMode0Shift))
}

const (
	RegisterModerFieldMode1Shift = 2
	RegisterModerFieldMode1Mask  = 0xc
)

// GetMode1 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode1Mask) >> RegisterModerFieldMode1Shift)
}

// SetMode1 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode1Mask)|(uint32(value)<<RegisterModerFieldMode1Shift))
}

const (
	RegisterModerFieldMode2Shift = 4
	RegisterModerFieldMode2Mask  = 0x30
)

// GetMode2 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode2Mask) >> RegisterModerFieldMode2Shift)
}

// SetMode2 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode2Mask)|(uint32(value)<<RegisterModerFieldMode2Shift))
}

const (
	RegisterModerFieldMode3Shift = 6
	RegisterModerFieldMode3Mask  = 0xc0
)

// GetMode3 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode3Mask) >> RegisterModerFieldMode3Shift)
}

// SetMode3 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode3Mask)|(uint32(value)<<RegisterModerFieldMode3Shift))
}

const (
	RegisterModerFieldMode4Shift = 8
	RegisterModerFieldMode4Mask  = 0x300
)

// GetMode4 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode4Mask) >> RegisterModerFieldMode4Shift)
}

// SetMode4 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode4Mask)|(uint32(value)<<RegisterModerFieldMode4Shift))
}

const (
	RegisterModerFieldMode5Shift = 10
	RegisterModerFieldMode5Mask  = 0xc00
)

// GetMode5 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode5Mask) >> RegisterModerFieldMode5Shift)
}

// SetMode5 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode5Mask)|(uint32(value)<<RegisterModerFieldMode5Shift))
}

const (
	RegisterModerFieldMode6Shift = 12
	RegisterModerFieldMode6Mask  = 0x3000
)

// GetMode6 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode6Mask) >> RegisterModerFieldMode6Shift)
}

// SetMode6 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode6Mask)|(uint32(value)<<RegisterModerFieldMode6Shift))
}

const (
	RegisterModerFieldMode7Shift = 14
	RegisterModerFieldMode7Mask  = 0xc000
)

// GetMode7 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode7Mask) >> RegisterModerFieldMode7Shift)
}

// SetMode7 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode7Mask)|(uint32(value)<<RegisterModerFieldMode7Shift))
}

const (
	RegisterModerFieldMode8Shift = 16
	RegisterModerFieldMode8Mask  = 0x30000
)

// GetMode8 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode8() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode8Mask) >> RegisterModerFieldMode8Shift)
}

// SetMode8 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode8(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode8Mask)|(uint32(value)<<RegisterModerFieldMode8Shift))
}

const (
	RegisterModerFieldMode9Shift = 18
	RegisterModerFieldMode9Mask  = 0xc0000
)

// GetMode9 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode9() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode9Mask) >> RegisterModerFieldMode9Shift)
}

// SetMode9 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode9(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode9Mask)|(uint32(value)<<RegisterModerFieldMode9Shift))
}

const (
	RegisterModerFieldMode10Shift = 20
	RegisterModerFieldMode10Mask  = 0x300000
)

// GetMode10 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode10() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode10Mask) >> RegisterModerFieldMode10Shift)
}

// SetMode10 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode10(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode10Mask)|(uint32(value)<<RegisterModerFieldMode10Shift))
}

const (
	RegisterModerFieldMode11Shift = 22
	RegisterModerFieldMode11Mask  = 0xc00000
)

// GetMode11 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode11Mask) >> RegisterModerFieldMode11Shift)
}

// SetMode11 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode11(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode11Mask)|(uint32(value)<<RegisterModerFieldMode11Shift))
}

const (
	RegisterModerFieldMode12Shift = 24
	RegisterModerFieldMode12Mask  = 0x3000000
)

// GetMode12 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode12() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode12Mask) >> RegisterModerFieldMode12Shift)
}

// SetMode12 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode12(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode12Mask)|(uint32(value)<<RegisterModerFieldMode12Shift))
}

const (
	RegisterModerFieldMode13Shift = 26
	RegisterModerFieldMode13Mask  = 0xc000000
)

// GetMode13 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode13() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode13Mask) >> RegisterModerFieldMode13Shift)
}

// SetMode13 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode13(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode13Mask)|(uint32(value)<<RegisterModerFieldMode13Shift))
}

const (
	RegisterModerFieldMode14Shift = 28
	RegisterModerFieldMode14Mask  = 0x30000000
)

// GetMode14 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode14() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode14Mask) >> RegisterModerFieldMode14Shift)
}

// SetMode14 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode14(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode14Mask)|(uint32(value)<<RegisterModerFieldMode14Shift))
}

const (
	RegisterModerFieldMode15Shift = 30
	RegisterModerFieldMode15Mask  = 0xc0000000
)

// GetMode15 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) GetMode15() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterModerFieldMode15Mask) >> RegisterModerFieldMode15Shift)
}

// SetMode15 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
func (r *registerModerType) SetMode15(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterModerFieldMode15Mask)|(uint32(value)<<RegisterModerFieldMode15Shift))
}

// registerOtyperType GPIO port output type register
type registerOtyperType uint32

const (
	RegisterOtyperFieldOt0Shift = 0
	RegisterOtyperFieldOt0Mask  = 0x1
)

// GetOt0 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt0Mask) != 0
}

// SetOt0 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt0Mask)
	}
}

const (
	RegisterOtyperFieldOt1Shift = 1
	RegisterOtyperFieldOt1Mask  = 0x2
)

// GetOt1 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt1Mask) != 0
}

// SetOt1 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt1Mask)
	}
}

const (
	RegisterOtyperFieldOt2Shift = 2
	RegisterOtyperFieldOt2Mask  = 0x4
)

// GetOt2 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt2Mask) != 0
}

// SetOt2 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt2Mask)
	}
}

const (
	RegisterOtyperFieldOt3Shift = 3
	RegisterOtyperFieldOt3Mask  = 0x8
)

// GetOt3 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt3Mask) != 0
}

// SetOt3 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt3Mask)
	}
}

const (
	RegisterOtyperFieldOt4Shift = 4
	RegisterOtyperFieldOt4Mask  = 0x10
)

// GetOt4 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt4Mask) != 0
}

// SetOt4 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt4Mask)
	}
}

const (
	RegisterOtyperFieldOt5Shift = 5
	RegisterOtyperFieldOt5Mask  = 0x20
)

// GetOt5 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt5Mask) != 0
}

// SetOt5 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt5Mask)
	}
}

const (
	RegisterOtyperFieldOt6Shift = 6
	RegisterOtyperFieldOt6Mask  = 0x40
)

// GetOt6 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt6Mask) != 0
}

// SetOt6 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt6Mask)
	}
}

const (
	RegisterOtyperFieldOt7Shift = 7
	RegisterOtyperFieldOt7Mask  = 0x80
)

// GetOt7 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt7Mask) != 0
}

// SetOt7 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt7Mask)
	}
}

const (
	RegisterOtyperFieldOt8Shift = 8
	RegisterOtyperFieldOt8Mask  = 0x100
)

// GetOt8 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt8Mask) != 0
}

// SetOt8 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt8Mask)
	}
}

const (
	RegisterOtyperFieldOt9Shift = 9
	RegisterOtyperFieldOt9Mask  = 0x200
)

// GetOt9 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt9Mask) != 0
}

// SetOt9 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt9Mask)
	}
}

const (
	RegisterOtyperFieldOt10Shift = 10
	RegisterOtyperFieldOt10Mask  = 0x400
)

// GetOt10 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt10Mask) != 0
}

// SetOt10 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt10Mask)
	}
}

const (
	RegisterOtyperFieldOt11Shift = 11
	RegisterOtyperFieldOt11Mask  = 0x800
)

// GetOt11 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt11Mask) != 0
}

// SetOt11 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt11Mask)
	}
}

const (
	RegisterOtyperFieldOt12Shift = 12
	RegisterOtyperFieldOt12Mask  = 0x1000
)

// GetOt12 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt12Mask) != 0
}

// SetOt12 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt12Mask)
	}
}

const (
	RegisterOtyperFieldOt13Shift = 13
	RegisterOtyperFieldOt13Mask  = 0x2000
)

// GetOt13 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt13Mask) != 0
}

// SetOt13 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt13Mask)
	}
}

const (
	RegisterOtyperFieldOt14Shift = 14
	RegisterOtyperFieldOt14Mask  = 0x4000
)

// GetOt14 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt14Mask) != 0
}

// SetOt14 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt14Mask)
	}
}

const (
	RegisterOtyperFieldOt15Shift = 15
	RegisterOtyperFieldOt15Mask  = 0x8000
)

// GetOt15 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) GetOt15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtyperFieldOt15Mask) != 0
}

// SetOt15 Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
func (r *registerOtyperType) SetOt15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtyperFieldOt15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtyperFieldOt15Mask)
	}
}

// registerOspeedrType GPIO port output speed register
type registerOspeedrType uint32

const (
	RegisterOspeedrFieldOspeed0Shift = 0
	RegisterOspeedrFieldOspeed0Mask  = 0x3
)

// GetOspeed0 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed0Mask) >> RegisterOspeedrFieldOspeed0Shift)
}

// SetOspeed0 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed0Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed0Shift))
}

const (
	RegisterOspeedrFieldOspeed1Shift = 2
	RegisterOspeedrFieldOspeed1Mask  = 0xc
)

// GetOspeed1 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed1Mask) >> RegisterOspeedrFieldOspeed1Shift)
}

// SetOspeed1 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed1Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed1Shift))
}

const (
	RegisterOspeedrFieldOspeed2Shift = 4
	RegisterOspeedrFieldOspeed2Mask  = 0x30
)

// GetOspeed2 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed2Mask) >> RegisterOspeedrFieldOspeed2Shift)
}

// SetOspeed2 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed2Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed2Shift))
}

const (
	RegisterOspeedrFieldOspeed3Shift = 6
	RegisterOspeedrFieldOspeed3Mask  = 0xc0
)

// GetOspeed3 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed3Mask) >> RegisterOspeedrFieldOspeed3Shift)
}

// SetOspeed3 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed3Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed3Shift))
}

const (
	RegisterOspeedrFieldOspeed4Shift = 8
	RegisterOspeedrFieldOspeed4Mask  = 0x300
)

// GetOspeed4 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed4Mask) >> RegisterOspeedrFieldOspeed4Shift)
}

// SetOspeed4 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed4Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed4Shift))
}

const (
	RegisterOspeedrFieldOspeed5Shift = 10
	RegisterOspeedrFieldOspeed5Mask  = 0xc00
)

// GetOspeed5 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed5Mask) >> RegisterOspeedrFieldOspeed5Shift)
}

// SetOspeed5 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed5Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed5Shift))
}

const (
	RegisterOspeedrFieldOspeed6Shift = 12
	RegisterOspeedrFieldOspeed6Mask  = 0x3000
)

// GetOspeed6 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed6Mask) >> RegisterOspeedrFieldOspeed6Shift)
}

// SetOspeed6 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed6Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed6Shift))
}

const (
	RegisterOspeedrFieldOspeed7Shift = 14
	RegisterOspeedrFieldOspeed7Mask  = 0xc000
)

// GetOspeed7 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed7Mask) >> RegisterOspeedrFieldOspeed7Shift)
}

// SetOspeed7 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed7Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed7Shift))
}

const (
	RegisterOspeedrFieldOspeed8Shift = 16
	RegisterOspeedrFieldOspeed8Mask  = 0x30000
)

// GetOspeed8 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed8() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed8Mask) >> RegisterOspeedrFieldOspeed8Shift)
}

// SetOspeed8 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed8(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed8Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed8Shift))
}

const (
	RegisterOspeedrFieldOspeed9Shift = 18
	RegisterOspeedrFieldOspeed9Mask  = 0xc0000
)

// GetOspeed9 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed9() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed9Mask) >> RegisterOspeedrFieldOspeed9Shift)
}

// SetOspeed9 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed9(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed9Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed9Shift))
}

const (
	RegisterOspeedrFieldOspeed10Shift = 20
	RegisterOspeedrFieldOspeed10Mask  = 0x300000
)

// GetOspeed10 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed10() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed10Mask) >> RegisterOspeedrFieldOspeed10Shift)
}

// SetOspeed10 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed10(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed10Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed10Shift))
}

const (
	RegisterOspeedrFieldOspeed11Shift = 22
	RegisterOspeedrFieldOspeed11Mask  = 0xc00000
)

// GetOspeed11 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed11Mask) >> RegisterOspeedrFieldOspeed11Shift)
}

// SetOspeed11 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed11(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed11Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed11Shift))
}

const (
	RegisterOspeedrFieldOspeed12Shift = 24
	RegisterOspeedrFieldOspeed12Mask  = 0x3000000
)

// GetOspeed12 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed12() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed12Mask) >> RegisterOspeedrFieldOspeed12Shift)
}

// SetOspeed12 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed12(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed12Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed12Shift))
}

const (
	RegisterOspeedrFieldOspeed13Shift = 26
	RegisterOspeedrFieldOspeed13Mask  = 0xc000000
)

// GetOspeed13 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed13() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed13Mask) >> RegisterOspeedrFieldOspeed13Shift)
}

// SetOspeed13 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed13(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed13Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed13Shift))
}

const (
	RegisterOspeedrFieldOspeed14Shift = 28
	RegisterOspeedrFieldOspeed14Mask  = 0x30000000
)

// GetOspeed14 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed14() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed14Mask) >> RegisterOspeedrFieldOspeed14Shift)
}

// SetOspeed14 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed14(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed14Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed14Shift))
}

const (
	RegisterOspeedrFieldOspeed15Shift = 30
	RegisterOspeedrFieldOspeed15Mask  = 0xc0000000
)

// GetOspeed15 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) GetOspeed15() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOspeedrFieldOspeed15Mask) >> RegisterOspeedrFieldOspeed15Shift)
}

// SetOspeed15 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
func (r *registerOspeedrType) SetOspeed15(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOspeedrFieldOspeed15Mask)|(uint32(value)<<RegisterOspeedrFieldOspeed15Shift))
}

// registerPupdrType GPIO port pull-up/pull-down register
type registerPupdrType uint32

const (
	RegisterPupdrFieldPupd0Shift = 0
	RegisterPupdrFieldPupd0Mask  = 0x3
)

// GetPupd0 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd0Mask) >> RegisterPupdrFieldPupd0Shift)
}

// SetPupd0 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd0Mask)|(uint32(value)<<RegisterPupdrFieldPupd0Shift))
}

const (
	RegisterPupdrFieldPupd1Shift = 2
	RegisterPupdrFieldPupd1Mask  = 0xc
)

// GetPupd1 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd1Mask) >> RegisterPupdrFieldPupd1Shift)
}

// SetPupd1 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd1Mask)|(uint32(value)<<RegisterPupdrFieldPupd1Shift))
}

const (
	RegisterPupdrFieldPupd2Shift = 4
	RegisterPupdrFieldPupd2Mask  = 0x30
)

// GetPupd2 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd2Mask) >> RegisterPupdrFieldPupd2Shift)
}

// SetPupd2 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd2Mask)|(uint32(value)<<RegisterPupdrFieldPupd2Shift))
}

const (
	RegisterPupdrFieldPupd3Shift = 6
	RegisterPupdrFieldPupd3Mask  = 0xc0
)

// GetPupd3 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd3Mask) >> RegisterPupdrFieldPupd3Shift)
}

// SetPupd3 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd3Mask)|(uint32(value)<<RegisterPupdrFieldPupd3Shift))
}

const (
	RegisterPupdrFieldPupd4Shift = 8
	RegisterPupdrFieldPupd4Mask  = 0x300
)

// GetPupd4 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd4Mask) >> RegisterPupdrFieldPupd4Shift)
}

// SetPupd4 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd4Mask)|(uint32(value)<<RegisterPupdrFieldPupd4Shift))
}

const (
	RegisterPupdrFieldPupd5Shift = 10
	RegisterPupdrFieldPupd5Mask  = 0xc00
)

// GetPupd5 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd5Mask) >> RegisterPupdrFieldPupd5Shift)
}

// SetPupd5 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd5Mask)|(uint32(value)<<RegisterPupdrFieldPupd5Shift))
}

const (
	RegisterPupdrFieldPupd6Shift = 12
	RegisterPupdrFieldPupd6Mask  = 0x3000
)

// GetPupd6 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd6Mask) >> RegisterPupdrFieldPupd6Shift)
}

// SetPupd6 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd6Mask)|(uint32(value)<<RegisterPupdrFieldPupd6Shift))
}

const (
	RegisterPupdrFieldPupd7Shift = 14
	RegisterPupdrFieldPupd7Mask  = 0xc000
)

// GetPupd7 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd7Mask) >> RegisterPupdrFieldPupd7Shift)
}

// SetPupd7 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd7Mask)|(uint32(value)<<RegisterPupdrFieldPupd7Shift))
}

const (
	RegisterPupdrFieldPupd8Shift = 16
	RegisterPupdrFieldPupd8Mask  = 0x30000
)

// GetPupd8 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd8() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd8Mask) >> RegisterPupdrFieldPupd8Shift)
}

// SetPupd8 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd8(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd8Mask)|(uint32(value)<<RegisterPupdrFieldPupd8Shift))
}

const (
	RegisterPupdrFieldPupd9Shift = 18
	RegisterPupdrFieldPupd9Mask  = 0xc0000
)

// GetPupd9 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd9() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd9Mask) >> RegisterPupdrFieldPupd9Shift)
}

// SetPupd9 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd9(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd9Mask)|(uint32(value)<<RegisterPupdrFieldPupd9Shift))
}

const (
	RegisterPupdrFieldPupd10Shift = 20
	RegisterPupdrFieldPupd10Mask  = 0x300000
)

// GetPupd10 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd10() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd10Mask) >> RegisterPupdrFieldPupd10Shift)
}

// SetPupd10 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd10(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd10Mask)|(uint32(value)<<RegisterPupdrFieldPupd10Shift))
}

const (
	RegisterPupdrFieldPupd11Shift = 22
	RegisterPupdrFieldPupd11Mask  = 0xc00000
)

// GetPupd11 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd11Mask) >> RegisterPupdrFieldPupd11Shift)
}

// SetPupd11 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd11(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd11Mask)|(uint32(value)<<RegisterPupdrFieldPupd11Shift))
}

const (
	RegisterPupdrFieldPupd12Shift = 24
	RegisterPupdrFieldPupd12Mask  = 0x3000000
)

// GetPupd12 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd12() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd12Mask) >> RegisterPupdrFieldPupd12Shift)
}

// SetPupd12 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd12(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd12Mask)|(uint32(value)<<RegisterPupdrFieldPupd12Shift))
}

const (
	RegisterPupdrFieldPupd13Shift = 26
	RegisterPupdrFieldPupd13Mask  = 0xc000000
)

// GetPupd13 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd13() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd13Mask) >> RegisterPupdrFieldPupd13Shift)
}

// SetPupd13 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd13(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd13Mask)|(uint32(value)<<RegisterPupdrFieldPupd13Shift))
}

const (
	RegisterPupdrFieldPupd14Shift = 28
	RegisterPupdrFieldPupd14Mask  = 0x30000000
)

// GetPupd14 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd14() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd14Mask) >> RegisterPupdrFieldPupd14Shift)
}

// SetPupd14 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd14(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd14Mask)|(uint32(value)<<RegisterPupdrFieldPupd14Shift))
}

const (
	RegisterPupdrFieldPupd15Shift = 30
	RegisterPupdrFieldPupd15Mask  = 0xc0000000
)

// GetPupd15 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) GetPupd15() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPupdrFieldPupd15Mask) >> RegisterPupdrFieldPupd15Shift)
}

// SetPupd15 [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
func (r *registerPupdrType) SetPupd15(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPupdrFieldPupd15Mask)|(uint32(value)<<RegisterPupdrFieldPupd15Shift))
}

// registerIdrType GPIO port input data register
type registerIdrType uint32

const (
	RegisterIdrFieldId0Shift = 0
	RegisterIdrFieldId0Mask  = 0x1
)

// GetId0 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId0Mask) != 0
}

// SetId0 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId0Mask)
	}
}

const (
	RegisterIdrFieldId1Shift = 1
	RegisterIdrFieldId1Mask  = 0x2
)

// GetId1 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId1Mask) != 0
}

// SetId1 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId1Mask)
	}
}

const (
	RegisterIdrFieldId2Shift = 2
	RegisterIdrFieldId2Mask  = 0x4
)

// GetId2 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId2Mask) != 0
}

// SetId2 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId2Mask)
	}
}

const (
	RegisterIdrFieldId3Shift = 3
	RegisterIdrFieldId3Mask  = 0x8
)

// GetId3 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId3Mask) != 0
}

// SetId3 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId3Mask)
	}
}

const (
	RegisterIdrFieldId4Shift = 4
	RegisterIdrFieldId4Mask  = 0x10
)

// GetId4 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId4Mask) != 0
}

// SetId4 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId4Mask)
	}
}

const (
	RegisterIdrFieldId5Shift = 5
	RegisterIdrFieldId5Mask  = 0x20
)

// GetId5 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId5Mask) != 0
}

// SetId5 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId5Mask)
	}
}

const (
	RegisterIdrFieldId6Shift = 6
	RegisterIdrFieldId6Mask  = 0x40
)

// GetId6 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId6Mask) != 0
}

// SetId6 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId6Mask)
	}
}

const (
	RegisterIdrFieldId7Shift = 7
	RegisterIdrFieldId7Mask  = 0x80
)

// GetId7 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId7Mask) != 0
}

// SetId7 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId7Mask)
	}
}

const (
	RegisterIdrFieldId8Shift = 8
	RegisterIdrFieldId8Mask  = 0x100
)

// GetId8 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId8Mask) != 0
}

// SetId8 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId8Mask)
	}
}

const (
	RegisterIdrFieldId9Shift = 9
	RegisterIdrFieldId9Mask  = 0x200
)

// GetId9 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId9Mask) != 0
}

// SetId9 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId9Mask)
	}
}

const (
	RegisterIdrFieldId10Shift = 10
	RegisterIdrFieldId10Mask  = 0x400
)

// GetId10 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId10Mask) != 0
}

// SetId10 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId10Mask)
	}
}

const (
	RegisterIdrFieldId11Shift = 11
	RegisterIdrFieldId11Mask  = 0x800
)

// GetId11 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId11Mask) != 0
}

// SetId11 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId11Mask)
	}
}

const (
	RegisterIdrFieldId12Shift = 12
	RegisterIdrFieldId12Mask  = 0x1000
)

// GetId12 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId12Mask) != 0
}

// SetId12 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId12Mask)
	}
}

const (
	RegisterIdrFieldId13Shift = 13
	RegisterIdrFieldId13Mask  = 0x2000
)

// GetId13 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId13Mask) != 0
}

// SetId13 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId13Mask)
	}
}

const (
	RegisterIdrFieldId14Shift = 14
	RegisterIdrFieldId14Mask  = 0x4000
)

// GetId14 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId14Mask) != 0
}

// SetId14 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId14Mask)
	}
}

const (
	RegisterIdrFieldId15Shift = 15
	RegisterIdrFieldId15Mask  = 0x8000
)

// GetId15 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) GetId15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldId15Mask) != 0
}

// SetId15 Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
func (r *registerIdrType) SetId15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdrFieldId15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldId15Mask)
	}
}

// registerOdrType GPIO port output data register
type registerOdrType uint32

const (
	RegisterOdrFieldOd0Shift = 0
	RegisterOdrFieldOd0Mask  = 0x1
)

// GetOd0 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd0Mask) != 0
}

// SetOd0 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd0Mask)
	}
}

const (
	RegisterOdrFieldOd1Shift = 1
	RegisterOdrFieldOd1Mask  = 0x2
)

// GetOd1 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd1Mask) != 0
}

// SetOd1 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd1Mask)
	}
}

const (
	RegisterOdrFieldOd2Shift = 2
	RegisterOdrFieldOd2Mask  = 0x4
)

// GetOd2 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd2Mask) != 0
}

// SetOd2 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd2Mask)
	}
}

const (
	RegisterOdrFieldOd3Shift = 3
	RegisterOdrFieldOd3Mask  = 0x8
)

// GetOd3 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd3Mask) != 0
}

// SetOd3 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd3Mask)
	}
}

const (
	RegisterOdrFieldOd4Shift = 4
	RegisterOdrFieldOd4Mask  = 0x10
)

// GetOd4 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd4Mask) != 0
}

// SetOd4 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd4Mask)
	}
}

const (
	RegisterOdrFieldOd5Shift = 5
	RegisterOdrFieldOd5Mask  = 0x20
)

// GetOd5 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd5Mask) != 0
}

// SetOd5 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd5Mask)
	}
}

const (
	RegisterOdrFieldOd6Shift = 6
	RegisterOdrFieldOd6Mask  = 0x40
)

// GetOd6 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd6Mask) != 0
}

// SetOd6 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd6Mask)
	}
}

const (
	RegisterOdrFieldOd7Shift = 7
	RegisterOdrFieldOd7Mask  = 0x80
)

// GetOd7 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd7Mask) != 0
}

// SetOd7 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd7Mask)
	}
}

const (
	RegisterOdrFieldOd8Shift = 8
	RegisterOdrFieldOd8Mask  = 0x100
)

// GetOd8 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd8Mask) != 0
}

// SetOd8 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd8Mask)
	}
}

const (
	RegisterOdrFieldOd9Shift = 9
	RegisterOdrFieldOd9Mask  = 0x200
)

// GetOd9 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd9Mask) != 0
}

// SetOd9 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd9Mask)
	}
}

const (
	RegisterOdrFieldOd10Shift = 10
	RegisterOdrFieldOd10Mask  = 0x400
)

// GetOd10 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd10Mask) != 0
}

// SetOd10 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd10Mask)
	}
}

const (
	RegisterOdrFieldOd11Shift = 11
	RegisterOdrFieldOd11Mask  = 0x800
)

// GetOd11 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd11Mask) != 0
}

// SetOd11 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd11Mask)
	}
}

const (
	RegisterOdrFieldOd12Shift = 12
	RegisterOdrFieldOd12Mask  = 0x1000
)

// GetOd12 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd12Mask) != 0
}

// SetOd12 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd12Mask)
	}
}

const (
	RegisterOdrFieldOd13Shift = 13
	RegisterOdrFieldOd13Mask  = 0x2000
)

// GetOd13 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd13Mask) != 0
}

// SetOd13 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd13Mask)
	}
}

const (
	RegisterOdrFieldOd14Shift = 14
	RegisterOdrFieldOd14Mask  = 0x4000
)

// GetOd14 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd14Mask) != 0
}

// SetOd14 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd14Mask)
	}
}

const (
	RegisterOdrFieldOd15Shift = 15
	RegisterOdrFieldOd15Mask  = 0x8000
)

// GetOd15 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) GetOd15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdrFieldOd15Mask) != 0
}

// SetOd15 Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
func (r *registerOdrType) SetOd15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdrFieldOd15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdrFieldOd15Mask)
	}
}

// registerBsrrType GPIO port bit set/reset register
type registerBsrrType uint32

const (
	RegisterBsrrFieldBs0Shift = 0
	RegisterBsrrFieldBs0Mask  = 0x1
)

// GetBs0 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs0Mask) != 0
}

// SetBs0 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs0Mask)
	}
}

const (
	RegisterBsrrFieldBs1Shift = 1
	RegisterBsrrFieldBs1Mask  = 0x2
)

// GetBs1 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs1Mask) != 0
}

// SetBs1 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs1Mask)
	}
}

const (
	RegisterBsrrFieldBs2Shift = 2
	RegisterBsrrFieldBs2Mask  = 0x4
)

// GetBs2 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs2Mask) != 0
}

// SetBs2 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs2Mask)
	}
}

const (
	RegisterBsrrFieldBs3Shift = 3
	RegisterBsrrFieldBs3Mask  = 0x8
)

// GetBs3 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs3Mask) != 0
}

// SetBs3 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs3Mask)
	}
}

const (
	RegisterBsrrFieldBs4Shift = 4
	RegisterBsrrFieldBs4Mask  = 0x10
)

// GetBs4 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs4Mask) != 0
}

// SetBs4 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs4Mask)
	}
}

const (
	RegisterBsrrFieldBs5Shift = 5
	RegisterBsrrFieldBs5Mask  = 0x20
)

// GetBs5 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs5Mask) != 0
}

// SetBs5 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs5Mask)
	}
}

const (
	RegisterBsrrFieldBs6Shift = 6
	RegisterBsrrFieldBs6Mask  = 0x40
)

// GetBs6 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs6Mask) != 0
}

// SetBs6 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs6Mask)
	}
}

const (
	RegisterBsrrFieldBs7Shift = 7
	RegisterBsrrFieldBs7Mask  = 0x80
)

// GetBs7 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs7Mask) != 0
}

// SetBs7 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs7Mask)
	}
}

const (
	RegisterBsrrFieldBs8Shift = 8
	RegisterBsrrFieldBs8Mask  = 0x100
)

// GetBs8 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs8Mask) != 0
}

// SetBs8 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs8Mask)
	}
}

const (
	RegisterBsrrFieldBs9Shift = 9
	RegisterBsrrFieldBs9Mask  = 0x200
)

// GetBs9 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs9Mask) != 0
}

// SetBs9 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs9Mask)
	}
}

const (
	RegisterBsrrFieldBs10Shift = 10
	RegisterBsrrFieldBs10Mask  = 0x400
)

// GetBs10 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs10Mask) != 0
}

// SetBs10 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs10Mask)
	}
}

const (
	RegisterBsrrFieldBs11Shift = 11
	RegisterBsrrFieldBs11Mask  = 0x800
)

// GetBs11 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs11Mask) != 0
}

// SetBs11 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs11Mask)
	}
}

const (
	RegisterBsrrFieldBs12Shift = 12
	RegisterBsrrFieldBs12Mask  = 0x1000
)

// GetBs12 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs12Mask) != 0
}

// SetBs12 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs12Mask)
	}
}

const (
	RegisterBsrrFieldBs13Shift = 13
	RegisterBsrrFieldBs13Mask  = 0x2000
)

// GetBs13 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs13Mask) != 0
}

// SetBs13 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs13Mask)
	}
}

const (
	RegisterBsrrFieldBs14Shift = 14
	RegisterBsrrFieldBs14Mask  = 0x4000
)

// GetBs14 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs14Mask) != 0
}

// SetBs14 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs14Mask)
	}
}

const (
	RegisterBsrrFieldBs15Shift = 15
	RegisterBsrrFieldBs15Mask  = 0x8000
)

// GetBs15 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) GetBs15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBs15Mask) != 0
}

// SetBs15 Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
func (r *registerBsrrType) SetBs15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBs15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBs15Mask)
	}
}

const (
	RegisterBsrrFieldBr0Shift = 16
	RegisterBsrrFieldBr0Mask  = 0x10000
)

// GetBr0 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr0Mask) != 0
}

// SetBr0 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr0Mask)
	}
}

const (
	RegisterBsrrFieldBr1Shift = 17
	RegisterBsrrFieldBr1Mask  = 0x20000
)

// GetBr1 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr1Mask) != 0
}

// SetBr1 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr1Mask)
	}
}

const (
	RegisterBsrrFieldBr2Shift = 18
	RegisterBsrrFieldBr2Mask  = 0x40000
)

// GetBr2 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr2Mask) != 0
}

// SetBr2 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr2Mask)
	}
}

const (
	RegisterBsrrFieldBr3Shift = 19
	RegisterBsrrFieldBr3Mask  = 0x80000
)

// GetBr3 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr3Mask) != 0
}

// SetBr3 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr3Mask)
	}
}

const (
	RegisterBsrrFieldBr4Shift = 20
	RegisterBsrrFieldBr4Mask  = 0x100000
)

// GetBr4 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr4Mask) != 0
}

// SetBr4 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr4Mask)
	}
}

const (
	RegisterBsrrFieldBr5Shift = 21
	RegisterBsrrFieldBr5Mask  = 0x200000
)

// GetBr5 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr5Mask) != 0
}

// SetBr5 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr5Mask)
	}
}

const (
	RegisterBsrrFieldBr6Shift = 22
	RegisterBsrrFieldBr6Mask  = 0x400000
)

// GetBr6 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr6Mask) != 0
}

// SetBr6 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr6Mask)
	}
}

const (
	RegisterBsrrFieldBr7Shift = 23
	RegisterBsrrFieldBr7Mask  = 0x800000
)

// GetBr7 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr7Mask) != 0
}

// SetBr7 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr7Mask)
	}
}

const (
	RegisterBsrrFieldBr8Shift = 24
	RegisterBsrrFieldBr8Mask  = 0x1000000
)

// GetBr8 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr8Mask) != 0
}

// SetBr8 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr8Mask)
	}
}

const (
	RegisterBsrrFieldBr9Shift = 25
	RegisterBsrrFieldBr9Mask  = 0x2000000
)

// GetBr9 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr9Mask) != 0
}

// SetBr9 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr9Mask)
	}
}

const (
	RegisterBsrrFieldBr10Shift = 26
	RegisterBsrrFieldBr10Mask  = 0x4000000
)

// GetBr10 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr10Mask) != 0
}

// SetBr10 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr10Mask)
	}
}

const (
	RegisterBsrrFieldBr11Shift = 27
	RegisterBsrrFieldBr11Mask  = 0x8000000
)

// GetBr11 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr11Mask) != 0
}

// SetBr11 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr11Mask)
	}
}

const (
	RegisterBsrrFieldBr12Shift = 28
	RegisterBsrrFieldBr12Mask  = 0x10000000
)

// GetBr12 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr12Mask) != 0
}

// SetBr12 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr12Mask)
	}
}

const (
	RegisterBsrrFieldBr13Shift = 29
	RegisterBsrrFieldBr13Mask  = 0x20000000
)

// GetBr13 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr13Mask) != 0
}

// SetBr13 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr13Mask)
	}
}

const (
	RegisterBsrrFieldBr14Shift = 30
	RegisterBsrrFieldBr14Mask  = 0x40000000
)

// GetBr14 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr14Mask) != 0
}

// SetBr14 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr14Mask)
	}
}

const (
	RegisterBsrrFieldBr15Shift = 31
	RegisterBsrrFieldBr15Mask  = 0x80000000
)

// GetBr15 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) GetBr15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBsrrFieldBr15Mask) != 0
}

// SetBr15 Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
func (r *registerBsrrType) SetBr15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBsrrFieldBr15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBsrrFieldBr15Mask)
	}
}

// registerLckrType This register is used to lock the configuration of the port bits when a correct write sequence is applied to bit 16 (LCKK). The value of bits [15:0] is used to lock the configuration of the GPIO. During the write sequence, the value of LCKR[15:0] must not change. When the LOCK sequence has been applied on a port bit, the value of this port bit can no longer be modified until the next MCU reset or peripheral reset.A specific write sequence is used to write to the GPIOx_LCKR register. Only word access (32-bit long) is allowed during this locking sequence.Each lock bit freezes a specific configuration register (control and alternate function registers).
type registerLckrType uint32

const (
	RegisterLckrFieldLck0Shift = 0
	RegisterLckrFieldLck0Mask  = 0x1
)

// GetLck0 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck0Mask) != 0
}

// SetLck0 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck0Mask)
	}
}

const (
	RegisterLckrFieldLck1Shift = 1
	RegisterLckrFieldLck1Mask  = 0x2
)

// GetLck1 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck1Mask) != 0
}

// SetLck1 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck1Mask)
	}
}

const (
	RegisterLckrFieldLck2Shift = 2
	RegisterLckrFieldLck2Mask  = 0x4
)

// GetLck2 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck2Mask) != 0
}

// SetLck2 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck2Mask)
	}
}

const (
	RegisterLckrFieldLck3Shift = 3
	RegisterLckrFieldLck3Mask  = 0x8
)

// GetLck3 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck3Mask) != 0
}

// SetLck3 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck3Mask)
	}
}

const (
	RegisterLckrFieldLck4Shift = 4
	RegisterLckrFieldLck4Mask  = 0x10
)

// GetLck4 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck4Mask) != 0
}

// SetLck4 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck4Mask)
	}
}

const (
	RegisterLckrFieldLck5Shift = 5
	RegisterLckrFieldLck5Mask  = 0x20
)

// GetLck5 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck5Mask) != 0
}

// SetLck5 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck5Mask)
	}
}

const (
	RegisterLckrFieldLck6Shift = 6
	RegisterLckrFieldLck6Mask  = 0x40
)

// GetLck6 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck6Mask) != 0
}

// SetLck6 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck6Mask)
	}
}

const (
	RegisterLckrFieldLck7Shift = 7
	RegisterLckrFieldLck7Mask  = 0x80
)

// GetLck7 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck7Mask) != 0
}

// SetLck7 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck7Mask)
	}
}

const (
	RegisterLckrFieldLck8Shift = 8
	RegisterLckrFieldLck8Mask  = 0x100
)

// GetLck8 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck8Mask) != 0
}

// SetLck8 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck8Mask)
	}
}

const (
	RegisterLckrFieldLck9Shift = 9
	RegisterLckrFieldLck9Mask  = 0x200
)

// GetLck9 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck9Mask) != 0
}

// SetLck9 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck9Mask)
	}
}

const (
	RegisterLckrFieldLck10Shift = 10
	RegisterLckrFieldLck10Mask  = 0x400
)

// GetLck10 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck10Mask) != 0
}

// SetLck10 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck10Mask)
	}
}

const (
	RegisterLckrFieldLck11Shift = 11
	RegisterLckrFieldLck11Mask  = 0x800
)

// GetLck11 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck11Mask) != 0
}

// SetLck11 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck11Mask)
	}
}

const (
	RegisterLckrFieldLck12Shift = 12
	RegisterLckrFieldLck12Mask  = 0x1000
)

// GetLck12 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck12Mask) != 0
}

// SetLck12 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck12Mask)
	}
}

const (
	RegisterLckrFieldLck13Shift = 13
	RegisterLckrFieldLck13Mask  = 0x2000
)

// GetLck13 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck13Mask) != 0
}

// SetLck13 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck13Mask)
	}
}

const (
	RegisterLckrFieldLck14Shift = 14
	RegisterLckrFieldLck14Mask  = 0x4000
)

// GetLck14 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck14Mask) != 0
}

// SetLck14 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck14Mask)
	}
}

const (
	RegisterLckrFieldLck15Shift = 15
	RegisterLckrFieldLck15Mask  = 0x8000
)

// GetLck15 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) GetLck15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLck15Mask) != 0
}

// SetLck15 Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
func (r *registerLckrType) SetLck15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLck15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLck15Mask)
	}
}

const (
	RegisterLckrFieldLckkShift = 16
	RegisterLckrFieldLckkMask  = 0x10000
)

// GetLckk Lock key This bit can be read any time. It can only be modified using the lock key write sequence. LOCK key write sequence: WR LCKR[16] = 1 + LCKR[15:0] WR LCKR[16] = 0 + LCKR[15:0] WR LCKR[16] = 1 + LCKR[15:0] RD LCKR RD LCKR[16] = 1 (this read operation is optional but it confirms that the lock is active) Note: During the LOCK key write sequence, the value of LCK[15:0] must not change. Any error in the lock sequence aborts the lock. After the first lock sequence on any bit of the port, any read access on the LCKK bit will return 1 until the next MCU reset or peripheral reset.
func (r *registerLckrType) GetLckk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLckrFieldLckkMask) != 0
}

// SetLckk Lock key This bit can be read any time. It can only be modified using the lock key write sequence. LOCK key write sequence: WR LCKR[16] = 1 + LCKR[15:0] WR LCKR[16] = 0 + LCKR[15:0] WR LCKR[16] = 1 + LCKR[15:0] RD LCKR RD LCKR[16] = 1 (this read operation is optional but it confirms that the lock is active) Note: During the LOCK key write sequence, the value of LCK[15:0] must not change. Any error in the lock sequence aborts the lock. After the first lock sequence on any bit of the port, any read access on the LCKK bit will return 1 until the next MCU reset or peripheral reset.
func (r *registerLckrType) SetLckk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLckrFieldLckkMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLckrFieldLckkMask)
	}
}

// registerAfrlType GPIO alternate function low register
type registerAfrlType uint32

const (
	RegisterAfrlFieldAfsel0Shift = 0
	RegisterAfrlFieldAfsel0Mask  = 0xf
)

// GetAfsel0 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) GetAfsel0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrlFieldAfsel0Mask) >> RegisterAfrlFieldAfsel0Shift)
}

// SetAfsel0 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) SetAfsel0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrlFieldAfsel0Mask)|(uint32(value)<<RegisterAfrlFieldAfsel0Shift))
}

const (
	RegisterAfrlFieldAfsel1Shift = 4
	RegisterAfrlFieldAfsel1Mask  = 0xf0
)

// GetAfsel1 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) GetAfsel1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrlFieldAfsel1Mask) >> RegisterAfrlFieldAfsel1Shift)
}

// SetAfsel1 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) SetAfsel1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrlFieldAfsel1Mask)|(uint32(value)<<RegisterAfrlFieldAfsel1Shift))
}

const (
	RegisterAfrlFieldAfsel2Shift = 8
	RegisterAfrlFieldAfsel2Mask  = 0xf00
)

// GetAfsel2 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) GetAfsel2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrlFieldAfsel2Mask) >> RegisterAfrlFieldAfsel2Shift)
}

// SetAfsel2 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) SetAfsel2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrlFieldAfsel2Mask)|(uint32(value)<<RegisterAfrlFieldAfsel2Shift))
}

const (
	RegisterAfrlFieldAfsel3Shift = 12
	RegisterAfrlFieldAfsel3Mask  = 0xf000
)

// GetAfsel3 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) GetAfsel3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrlFieldAfsel3Mask) >> RegisterAfrlFieldAfsel3Shift)
}

// SetAfsel3 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) SetAfsel3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrlFieldAfsel3Mask)|(uint32(value)<<RegisterAfrlFieldAfsel3Shift))
}

const (
	RegisterAfrlFieldAfsel4Shift = 16
	RegisterAfrlFieldAfsel4Mask  = 0xf0000
)

// GetAfsel4 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) GetAfsel4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrlFieldAfsel4Mask) >> RegisterAfrlFieldAfsel4Shift)
}

// SetAfsel4 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) SetAfsel4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrlFieldAfsel4Mask)|(uint32(value)<<RegisterAfrlFieldAfsel4Shift))
}

const (
	RegisterAfrlFieldAfsel5Shift = 20
	RegisterAfrlFieldAfsel5Mask  = 0xf00000
)

// GetAfsel5 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) GetAfsel5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrlFieldAfsel5Mask) >> RegisterAfrlFieldAfsel5Shift)
}

// SetAfsel5 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) SetAfsel5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrlFieldAfsel5Mask)|(uint32(value)<<RegisterAfrlFieldAfsel5Shift))
}

const (
	RegisterAfrlFieldAfsel6Shift = 24
	RegisterAfrlFieldAfsel6Mask  = 0xf000000
)

// GetAfsel6 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) GetAfsel6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrlFieldAfsel6Mask) >> RegisterAfrlFieldAfsel6Shift)
}

// SetAfsel6 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) SetAfsel6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrlFieldAfsel6Mask)|(uint32(value)<<RegisterAfrlFieldAfsel6Shift))
}

const (
	RegisterAfrlFieldAfsel7Shift = 28
	RegisterAfrlFieldAfsel7Mask  = 0xf0000000
)

// GetAfsel7 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) GetAfsel7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrlFieldAfsel7Mask) >> RegisterAfrlFieldAfsel7Shift)
}

// SetAfsel7 [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
func (r *registerAfrlType) SetAfsel7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrlFieldAfsel7Mask)|(uint32(value)<<RegisterAfrlFieldAfsel7Shift))
}

// registerAfrhType GPIO alternate function high register
type registerAfrhType uint32

const (
	RegisterAfrhFieldAfsel8Shift = 0
	RegisterAfrhFieldAfsel8Mask  = 0xf
)

// GetAfsel8 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) GetAfsel8() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrhFieldAfsel8Mask) >> RegisterAfrhFieldAfsel8Shift)
}

// SetAfsel8 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) SetAfsel8(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrhFieldAfsel8Mask)|(uint32(value)<<RegisterAfrhFieldAfsel8Shift))
}

const (
	RegisterAfrhFieldAfsel9Shift = 4
	RegisterAfrhFieldAfsel9Mask  = 0xf0
)

// GetAfsel9 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) GetAfsel9() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrhFieldAfsel9Mask) >> RegisterAfrhFieldAfsel9Shift)
}

// SetAfsel9 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) SetAfsel9(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrhFieldAfsel9Mask)|(uint32(value)<<RegisterAfrhFieldAfsel9Shift))
}

const (
	RegisterAfrhFieldAfsel10Shift = 8
	RegisterAfrhFieldAfsel10Mask  = 0xf00
)

// GetAfsel10 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) GetAfsel10() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrhFieldAfsel10Mask) >> RegisterAfrhFieldAfsel10Shift)
}

// SetAfsel10 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) SetAfsel10(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrhFieldAfsel10Mask)|(uint32(value)<<RegisterAfrhFieldAfsel10Shift))
}

const (
	RegisterAfrhFieldAfsel11Shift = 12
	RegisterAfrhFieldAfsel11Mask  = 0xf000
)

// GetAfsel11 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) GetAfsel11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrhFieldAfsel11Mask) >> RegisterAfrhFieldAfsel11Shift)
}

// SetAfsel11 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) SetAfsel11(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrhFieldAfsel11Mask)|(uint32(value)<<RegisterAfrhFieldAfsel11Shift))
}

const (
	RegisterAfrhFieldAfsel12Shift = 16
	RegisterAfrhFieldAfsel12Mask  = 0xf0000
)

// GetAfsel12 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) GetAfsel12() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrhFieldAfsel12Mask) >> RegisterAfrhFieldAfsel12Shift)
}

// SetAfsel12 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) SetAfsel12(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrhFieldAfsel12Mask)|(uint32(value)<<RegisterAfrhFieldAfsel12Shift))
}

const (
	RegisterAfrhFieldAfsel13Shift = 20
	RegisterAfrhFieldAfsel13Mask  = 0xf00000
)

// GetAfsel13 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) GetAfsel13() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrhFieldAfsel13Mask) >> RegisterAfrhFieldAfsel13Shift)
}

// SetAfsel13 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) SetAfsel13(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrhFieldAfsel13Mask)|(uint32(value)<<RegisterAfrhFieldAfsel13Shift))
}

const (
	RegisterAfrhFieldAfsel14Shift = 24
	RegisterAfrhFieldAfsel14Mask  = 0xf000000
)

// GetAfsel14 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) GetAfsel14() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrhFieldAfsel14Mask) >> RegisterAfrhFieldAfsel14Shift)
}

// SetAfsel14 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) SetAfsel14(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrhFieldAfsel14Mask)|(uint32(value)<<RegisterAfrhFieldAfsel14Shift))
}

const (
	RegisterAfrhFieldAfsel15Shift = 28
	RegisterAfrhFieldAfsel15Mask  = 0xf0000000
)

// GetAfsel15 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) GetAfsel15() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAfrhFieldAfsel15Mask) >> RegisterAfrhFieldAfsel15Shift)
}

// SetAfsel15 [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
func (r *registerAfrhType) SetAfsel15(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAfrhFieldAfsel15Mask)|(uint32(value)<<RegisterAfrhFieldAfsel15Shift))
}
