//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package exti

import (
	"unsafe"
	"volatile"
)

var (
	Exti = (*_exti)(unsafe.Pointer(uintptr(0x58000000)))
)

type _exti struct {
	Rtsr1   registerRtsr1Type
	Ftsr1   registerFtsr1Type
	Swier1  registerSwier1Type
	D3pmr1  registerD3pmr1Type
	D3pcr1l registerD3pcr1lType
	D3pcr1h registerD3pcr1hType
	_       [8]byte
	Rtsr2   registerRtsr2Type
	Ftsr2   registerFtsr2Type
	Swier2  registerSwier2Type
	D3pmr2  registerD3pmr2Type
	D3pcr2l registerD3pcr2lType
	D3pcr2h registerD3pcr2hType
	_       [8]byte
	Rtsr3   registerRtsr3Type
	Ftsr3   registerFtsr3Type
	Swier3  registerSwier3Type
	D3pmr3  registerD3pmr3Type
	_       [4]byte
	D3pcr3h registerD3pcr3hType
	_       [40]byte
	Cpuimr1 registerCpuimr1Type
	Cpuemr1 registerCpuemr1Type
	Cpupr1  registerCpupr1Type
	_       [4]byte
	Cpuimr2 registerCpuimr2Type
	Cpuemr2 registerCpuemr2Type
	Cpupr2  registerCpupr2Type
	_       [4]byte
	Cpuimr3 registerCpuimr3Type
	Cpuemr3 registerCpuemr3Type
	Cpupr3  registerCpupr3Type
}

// registerRtsr1Type EXTI rising trigger selection register
type registerRtsr1Type uint32

const (
	RegisterRtsr1FieldTr0Shift = 0
	RegisterRtsr1FieldTr0Mask  = 0x1
)

// GetTr0 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr0Mask) != 0
}

// SetTr0 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr0Mask)
	}
}

const (
	RegisterRtsr1FieldTr1Shift = 1
	RegisterRtsr1FieldTr1Mask  = 0x2
)

// GetTr1 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr1Mask) != 0
}

// SetTr1 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr1Mask)
	}
}

const (
	RegisterRtsr1FieldTr2Shift = 2
	RegisterRtsr1FieldTr2Mask  = 0x4
)

// GetTr2 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr2Mask) != 0
}

// SetTr2 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr2Mask)
	}
}

const (
	RegisterRtsr1FieldTr3Shift = 3
	RegisterRtsr1FieldTr3Mask  = 0x8
)

// GetTr3 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr3Mask) != 0
}

// SetTr3 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr3Mask)
	}
}

const (
	RegisterRtsr1FieldTr4Shift = 4
	RegisterRtsr1FieldTr4Mask  = 0x10
)

// GetTr4 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr4Mask) != 0
}

// SetTr4 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr4Mask)
	}
}

const (
	RegisterRtsr1FieldTr5Shift = 5
	RegisterRtsr1FieldTr5Mask  = 0x20
)

// GetTr5 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr5Mask) != 0
}

// SetTr5 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr5Mask)
	}
}

const (
	RegisterRtsr1FieldTr6Shift = 6
	RegisterRtsr1FieldTr6Mask  = 0x40
)

// GetTr6 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr6Mask) != 0
}

// SetTr6 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr6Mask)
	}
}

const (
	RegisterRtsr1FieldTr7Shift = 7
	RegisterRtsr1FieldTr7Mask  = 0x80
)

// GetTr7 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr7Mask) != 0
}

// SetTr7 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr7Mask)
	}
}

const (
	RegisterRtsr1FieldTr8Shift = 8
	RegisterRtsr1FieldTr8Mask  = 0x100
)

// GetTr8 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr8Mask) != 0
}

// SetTr8 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr8Mask)
	}
}

const (
	RegisterRtsr1FieldTr9Shift = 9
	RegisterRtsr1FieldTr9Mask  = 0x200
)

// GetTr9 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr9Mask) != 0
}

// SetTr9 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr9Mask)
	}
}

const (
	RegisterRtsr1FieldTr10Shift = 10
	RegisterRtsr1FieldTr10Mask  = 0x400
)

// GetTr10 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr10Mask) != 0
}

// SetTr10 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr10Mask)
	}
}

const (
	RegisterRtsr1FieldTr11Shift = 11
	RegisterRtsr1FieldTr11Mask  = 0x800
)

// GetTr11 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr11Mask) != 0
}

// SetTr11 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr11Mask)
	}
}

const (
	RegisterRtsr1FieldTr12Shift = 12
	RegisterRtsr1FieldTr12Mask  = 0x1000
)

// GetTr12 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr12Mask) != 0
}

// SetTr12 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr12Mask)
	}
}

const (
	RegisterRtsr1FieldTr13Shift = 13
	RegisterRtsr1FieldTr13Mask  = 0x2000
)

// GetTr13 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr13Mask) != 0
}

// SetTr13 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr13Mask)
	}
}

const (
	RegisterRtsr1FieldTr14Shift = 14
	RegisterRtsr1FieldTr14Mask  = 0x4000
)

// GetTr14 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr14Mask) != 0
}

// SetTr14 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr14Mask)
	}
}

const (
	RegisterRtsr1FieldTr15Shift = 15
	RegisterRtsr1FieldTr15Mask  = 0x8000
)

// GetTr15 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr15Mask) != 0
}

// SetTr15 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr15Mask)
	}
}

const (
	RegisterRtsr1FieldTr16Shift = 16
	RegisterRtsr1FieldTr16Mask  = 0x10000
)

// GetTr16 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr16Mask) != 0
}

// SetTr16 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr16Mask)
	}
}

const (
	RegisterRtsr1FieldTr17Shift = 17
	RegisterRtsr1FieldTr17Mask  = 0x20000
)

// GetTr17 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr17Mask) != 0
}

// SetTr17 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr17Mask)
	}
}

const (
	RegisterRtsr1FieldTr18Shift = 18
	RegisterRtsr1FieldTr18Mask  = 0x40000
)

// GetTr18 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr18Mask) != 0
}

// SetTr18 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr18Mask)
	}
}

const (
	RegisterRtsr1FieldTr19Shift = 19
	RegisterRtsr1FieldTr19Mask  = 0x80000
)

// GetTr19 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr19Mask) != 0
}

// SetTr19 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr19Mask)
	}
}

const (
	RegisterRtsr1FieldTr20Shift = 20
	RegisterRtsr1FieldTr20Mask  = 0x100000
)

// GetTr20 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr20Mask) != 0
}

// SetTr20 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr20Mask)
	}
}

const (
	RegisterRtsr1FieldTr21Shift = 21
	RegisterRtsr1FieldTr21Mask  = 0x200000
)

// GetTr21 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) GetTr21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr1FieldTr21Mask) != 0
}

// SetTr21 Rising trigger event configuration bit of Configurable Event input
func (r *registerRtsr1Type) SetTr21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr1FieldTr21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr1FieldTr21Mask)
	}
}

// registerFtsr1Type EXTI falling trigger selection register
type registerFtsr1Type uint32

const (
	RegisterFtsr1FieldTr0Shift = 0
	RegisterFtsr1FieldTr0Mask  = 0x1
)

// GetTr0 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr0Mask) != 0
}

// SetTr0 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr0Mask)
	}
}

const (
	RegisterFtsr1FieldTr1Shift = 1
	RegisterFtsr1FieldTr1Mask  = 0x2
)

// GetTr1 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr1Mask) != 0
}

// SetTr1 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr1Mask)
	}
}

const (
	RegisterFtsr1FieldTr2Shift = 2
	RegisterFtsr1FieldTr2Mask  = 0x4
)

// GetTr2 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr2Mask) != 0
}

// SetTr2 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr2Mask)
	}
}

const (
	RegisterFtsr1FieldTr3Shift = 3
	RegisterFtsr1FieldTr3Mask  = 0x8
)

// GetTr3 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr3Mask) != 0
}

// SetTr3 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr3Mask)
	}
}

const (
	RegisterFtsr1FieldTr4Shift = 4
	RegisterFtsr1FieldTr4Mask  = 0x10
)

// GetTr4 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr4Mask) != 0
}

// SetTr4 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr4Mask)
	}
}

const (
	RegisterFtsr1FieldTr5Shift = 5
	RegisterFtsr1FieldTr5Mask  = 0x20
)

// GetTr5 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr5Mask) != 0
}

// SetTr5 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr5Mask)
	}
}

const (
	RegisterFtsr1FieldTr6Shift = 6
	RegisterFtsr1FieldTr6Mask  = 0x40
)

// GetTr6 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr6Mask) != 0
}

// SetTr6 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr6Mask)
	}
}

const (
	RegisterFtsr1FieldTr7Shift = 7
	RegisterFtsr1FieldTr7Mask  = 0x80
)

// GetTr7 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr7Mask) != 0
}

// SetTr7 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr7Mask)
	}
}

const (
	RegisterFtsr1FieldTr8Shift = 8
	RegisterFtsr1FieldTr8Mask  = 0x100
)

// GetTr8 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr8Mask) != 0
}

// SetTr8 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr8Mask)
	}
}

const (
	RegisterFtsr1FieldTr9Shift = 9
	RegisterFtsr1FieldTr9Mask  = 0x200
)

// GetTr9 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr9Mask) != 0
}

// SetTr9 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr9Mask)
	}
}

const (
	RegisterFtsr1FieldTr10Shift = 10
	RegisterFtsr1FieldTr10Mask  = 0x400
)

// GetTr10 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr10Mask) != 0
}

// SetTr10 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr10Mask)
	}
}

const (
	RegisterFtsr1FieldTr11Shift = 11
	RegisterFtsr1FieldTr11Mask  = 0x800
)

// GetTr11 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr11Mask) != 0
}

// SetTr11 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr11Mask)
	}
}

const (
	RegisterFtsr1FieldTr12Shift = 12
	RegisterFtsr1FieldTr12Mask  = 0x1000
)

// GetTr12 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr12Mask) != 0
}

// SetTr12 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr12Mask)
	}
}

const (
	RegisterFtsr1FieldTr13Shift = 13
	RegisterFtsr1FieldTr13Mask  = 0x2000
)

// GetTr13 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr13Mask) != 0
}

// SetTr13 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr13Mask)
	}
}

const (
	RegisterFtsr1FieldTr14Shift = 14
	RegisterFtsr1FieldTr14Mask  = 0x4000
)

// GetTr14 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr14Mask) != 0
}

// SetTr14 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr14Mask)
	}
}

const (
	RegisterFtsr1FieldTr15Shift = 15
	RegisterFtsr1FieldTr15Mask  = 0x8000
)

// GetTr15 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr15Mask) != 0
}

// SetTr15 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr15Mask)
	}
}

const (
	RegisterFtsr1FieldTr16Shift = 16
	RegisterFtsr1FieldTr16Mask  = 0x10000
)

// GetTr16 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr16Mask) != 0
}

// SetTr16 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr16Mask)
	}
}

const (
	RegisterFtsr1FieldTr17Shift = 17
	RegisterFtsr1FieldTr17Mask  = 0x20000
)

// GetTr17 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr17Mask) != 0
}

// SetTr17 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr17Mask)
	}
}

const (
	RegisterFtsr1FieldTr18Shift = 18
	RegisterFtsr1FieldTr18Mask  = 0x40000
)

// GetTr18 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr18Mask) != 0
}

// SetTr18 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr18Mask)
	}
}

const (
	RegisterFtsr1FieldTr19Shift = 19
	RegisterFtsr1FieldTr19Mask  = 0x80000
)

// GetTr19 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr19Mask) != 0
}

// SetTr19 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr19Mask)
	}
}

const (
	RegisterFtsr1FieldTr20Shift = 20
	RegisterFtsr1FieldTr20Mask  = 0x100000
)

// GetTr20 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr20Mask) != 0
}

// SetTr20 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr20Mask)
	}
}

const (
	RegisterFtsr1FieldTr21Shift = 21
	RegisterFtsr1FieldTr21Mask  = 0x200000
)

// GetTr21 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) GetTr21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr1FieldTr21Mask) != 0
}

// SetTr21 Rising trigger event configuration bit of Configurable Event input
func (r *registerFtsr1Type) SetTr21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr1FieldTr21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr1FieldTr21Mask)
	}
}

// registerSwier1Type EXTI software interrupt event register
type registerSwier1Type uint32

const (
	RegisterSwier1FieldSwier0Shift = 0
	RegisterSwier1FieldSwier0Mask  = 0x1
)

// GetSwier0 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier0Mask) != 0
}

// SetSwier0 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier0Mask)
	}
}

const (
	RegisterSwier1FieldSwier1Shift = 1
	RegisterSwier1FieldSwier1Mask  = 0x2
)

// GetSwier1 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier1Mask) != 0
}

// SetSwier1 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier1Mask)
	}
}

const (
	RegisterSwier1FieldSwier2Shift = 2
	RegisterSwier1FieldSwier2Mask  = 0x4
)

// GetSwier2 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier2Mask) != 0
}

// SetSwier2 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier2Mask)
	}
}

const (
	RegisterSwier1FieldSwier3Shift = 3
	RegisterSwier1FieldSwier3Mask  = 0x8
)

// GetSwier3 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier3Mask) != 0
}

// SetSwier3 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier3Mask)
	}
}

const (
	RegisterSwier1FieldSwier4Shift = 4
	RegisterSwier1FieldSwier4Mask  = 0x10
)

// GetSwier4 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier4Mask) != 0
}

// SetSwier4 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier4Mask)
	}
}

const (
	RegisterSwier1FieldSwier5Shift = 5
	RegisterSwier1FieldSwier5Mask  = 0x20
)

// GetSwier5 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier5Mask) != 0
}

// SetSwier5 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier5Mask)
	}
}

const (
	RegisterSwier1FieldSwier6Shift = 6
	RegisterSwier1FieldSwier6Mask  = 0x40
)

// GetSwier6 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier6Mask) != 0
}

// SetSwier6 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier6Mask)
	}
}

const (
	RegisterSwier1FieldSwier7Shift = 7
	RegisterSwier1FieldSwier7Mask  = 0x80
)

// GetSwier7 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier7Mask) != 0
}

// SetSwier7 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier7Mask)
	}
}

const (
	RegisterSwier1FieldSwier8Shift = 8
	RegisterSwier1FieldSwier8Mask  = 0x100
)

// GetSwier8 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier8Mask) != 0
}

// SetSwier8 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier8Mask)
	}
}

const (
	RegisterSwier1FieldSwier9Shift = 9
	RegisterSwier1FieldSwier9Mask  = 0x200
)

// GetSwier9 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier9Mask) != 0
}

// SetSwier9 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier9Mask)
	}
}

const (
	RegisterSwier1FieldSwier10Shift = 10
	RegisterSwier1FieldSwier10Mask  = 0x400
)

// GetSwier10 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier10Mask) != 0
}

// SetSwier10 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier10Mask)
	}
}

const (
	RegisterSwier1FieldSwier11Shift = 11
	RegisterSwier1FieldSwier11Mask  = 0x800
)

// GetSwier11 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier11Mask) != 0
}

// SetSwier11 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier11Mask)
	}
}

const (
	RegisterSwier1FieldSwier12Shift = 12
	RegisterSwier1FieldSwier12Mask  = 0x1000
)

// GetSwier12 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier12Mask) != 0
}

// SetSwier12 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier12Mask)
	}
}

const (
	RegisterSwier1FieldSwier13Shift = 13
	RegisterSwier1FieldSwier13Mask  = 0x2000
)

// GetSwier13 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier13Mask) != 0
}

// SetSwier13 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier13Mask)
	}
}

const (
	RegisterSwier1FieldSwier14Shift = 14
	RegisterSwier1FieldSwier14Mask  = 0x4000
)

// GetSwier14 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier14Mask) != 0
}

// SetSwier14 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier14Mask)
	}
}

const (
	RegisterSwier1FieldSwier15Shift = 15
	RegisterSwier1FieldSwier15Mask  = 0x8000
)

// GetSwier15 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier15Mask) != 0
}

// SetSwier15 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier15Mask)
	}
}

const (
	RegisterSwier1FieldSwier16Shift = 16
	RegisterSwier1FieldSwier16Mask  = 0x10000
)

// GetSwier16 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier16Mask) != 0
}

// SetSwier16 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier16Mask)
	}
}

const (
	RegisterSwier1FieldSwier17Shift = 17
	RegisterSwier1FieldSwier17Mask  = 0x20000
)

// GetSwier17 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier17Mask) != 0
}

// SetSwier17 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier17Mask)
	}
}

const (
	RegisterSwier1FieldSwier18Shift = 18
	RegisterSwier1FieldSwier18Mask  = 0x40000
)

// GetSwier18 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier18Mask) != 0
}

// SetSwier18 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier18Mask)
	}
}

const (
	RegisterSwier1FieldSwier19Shift = 19
	RegisterSwier1FieldSwier19Mask  = 0x80000
)

// GetSwier19 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier19Mask) != 0
}

// SetSwier19 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier19Mask)
	}
}

const (
	RegisterSwier1FieldSwier20Shift = 20
	RegisterSwier1FieldSwier20Mask  = 0x100000
)

// GetSwier20 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier20Mask) != 0
}

// SetSwier20 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier20Mask)
	}
}

const (
	RegisterSwier1FieldSwier21Shift = 21
	RegisterSwier1FieldSwier21Mask  = 0x200000
)

// GetSwier21 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) GetSwier21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier1FieldSwier21Mask) != 0
}

// SetSwier21 Rising trigger event configuration bit of Configurable Event input
func (r *registerSwier1Type) SetSwier21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier1FieldSwier21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier1FieldSwier21Mask)
	}
}

// registerD3pmr1Type EXTI D3 pending mask register
type registerD3pmr1Type uint32

const (
	RegisterD3pmr1FieldMr0Shift = 0
	RegisterD3pmr1FieldMr0Mask  = 0x1
)

// GetMr0 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr0Mask) != 0
}

// SetMr0 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr0Mask)
	}
}

const (
	RegisterD3pmr1FieldMr1Shift = 1
	RegisterD3pmr1FieldMr1Mask  = 0x2
)

// GetMr1 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr1Mask) != 0
}

// SetMr1 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr1Mask)
	}
}

const (
	RegisterD3pmr1FieldMr2Shift = 2
	RegisterD3pmr1FieldMr2Mask  = 0x4
)

// GetMr2 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr2Mask) != 0
}

// SetMr2 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr2Mask)
	}
}

const (
	RegisterD3pmr1FieldMr3Shift = 3
	RegisterD3pmr1FieldMr3Mask  = 0x8
)

// GetMr3 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr3Mask) != 0
}

// SetMr3 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr3Mask)
	}
}

const (
	RegisterD3pmr1FieldMr4Shift = 4
	RegisterD3pmr1FieldMr4Mask  = 0x10
)

// GetMr4 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr4Mask) != 0
}

// SetMr4 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr4Mask)
	}
}

const (
	RegisterD3pmr1FieldMr5Shift = 5
	RegisterD3pmr1FieldMr5Mask  = 0x20
)

// GetMr5 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr5Mask) != 0
}

// SetMr5 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr5Mask)
	}
}

const (
	RegisterD3pmr1FieldMr6Shift = 6
	RegisterD3pmr1FieldMr6Mask  = 0x40
)

// GetMr6 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr6Mask) != 0
}

// SetMr6 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr6Mask)
	}
}

const (
	RegisterD3pmr1FieldMr7Shift = 7
	RegisterD3pmr1FieldMr7Mask  = 0x80
)

// GetMr7 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr7Mask) != 0
}

// SetMr7 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr7Mask)
	}
}

const (
	RegisterD3pmr1FieldMr8Shift = 8
	RegisterD3pmr1FieldMr8Mask  = 0x100
)

// GetMr8 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr8Mask) != 0
}

// SetMr8 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr8Mask)
	}
}

const (
	RegisterD3pmr1FieldMr9Shift = 9
	RegisterD3pmr1FieldMr9Mask  = 0x200
)

// GetMr9 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr9Mask) != 0
}

// SetMr9 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr9Mask)
	}
}

const (
	RegisterD3pmr1FieldMr10Shift = 10
	RegisterD3pmr1FieldMr10Mask  = 0x400
)

// GetMr10 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr10Mask) != 0
}

// SetMr10 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr10Mask)
	}
}

const (
	RegisterD3pmr1FieldMr11Shift = 11
	RegisterD3pmr1FieldMr11Mask  = 0x800
)

// GetMr11 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr11Mask) != 0
}

// SetMr11 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr11Mask)
	}
}

const (
	RegisterD3pmr1FieldMr12Shift = 12
	RegisterD3pmr1FieldMr12Mask  = 0x1000
)

// GetMr12 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr12Mask) != 0
}

// SetMr12 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr12Mask)
	}
}

const (
	RegisterD3pmr1FieldMr13Shift = 13
	RegisterD3pmr1FieldMr13Mask  = 0x2000
)

// GetMr13 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr13Mask) != 0
}

// SetMr13 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr13Mask)
	}
}

const (
	RegisterD3pmr1FieldMr14Shift = 14
	RegisterD3pmr1FieldMr14Mask  = 0x4000
)

// GetMr14 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr14Mask) != 0
}

// SetMr14 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr14Mask)
	}
}

const (
	RegisterD3pmr1FieldMr15Shift = 15
	RegisterD3pmr1FieldMr15Mask  = 0x8000
)

// GetMr15 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr15Mask) != 0
}

// SetMr15 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr15Mask)
	}
}

const (
	RegisterD3pmr1FieldMr19Shift = 19
	RegisterD3pmr1FieldMr19Mask  = 0x80000
)

// GetMr19 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr19Mask) != 0
}

// SetMr19 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr19Mask)
	}
}

const (
	RegisterD3pmr1FieldMr20Shift = 20
	RegisterD3pmr1FieldMr20Mask  = 0x100000
)

// GetMr20 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr20Mask) != 0
}

// SetMr20 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr20Mask)
	}
}

const (
	RegisterD3pmr1FieldMr21Shift = 21
	RegisterD3pmr1FieldMr21Mask  = 0x200000
)

// GetMr21 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr21Mask) != 0
}

// SetMr21 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr21Mask)
	}
}

const (
	RegisterD3pmr1FieldMr25Shift = 25
	RegisterD3pmr1FieldMr25Mask  = 0x2000000
)

// GetMr25 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) GetMr25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr1FieldMr25Mask) != 0
}

// SetMr25 Rising trigger event configuration bit of Configurable Event input
func (r *registerD3pmr1Type) SetMr25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr1FieldMr25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr1FieldMr25Mask)
	}
}

// registerD3pcr1lType EXTI D3 pending clear selection register low
type registerD3pcr1lType uint32

const (
	RegisterD3pcr1lFieldPcs0Shift = 0
	RegisterD3pcr1lFieldPcs0Mask  = 0x3
)

// GetPcs0 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs0Mask) >> RegisterD3pcr1lFieldPcs0Shift)
}

// SetPcs0 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs0Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs0Shift))
}

const (
	RegisterD3pcr1lFieldPcs1Shift = 2
	RegisterD3pcr1lFieldPcs1Mask  = 0xc
)

// GetPcs1 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs1Mask) >> RegisterD3pcr1lFieldPcs1Shift)
}

// SetPcs1 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs1Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs1Shift))
}

const (
	RegisterD3pcr1lFieldPcs2Shift = 4
	RegisterD3pcr1lFieldPcs2Mask  = 0x30
)

// GetPcs2 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs2Mask) >> RegisterD3pcr1lFieldPcs2Shift)
}

// SetPcs2 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs2Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs2Shift))
}

const (
	RegisterD3pcr1lFieldPcs3Shift = 6
	RegisterD3pcr1lFieldPcs3Mask  = 0xc0
)

// GetPcs3 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs3Mask) >> RegisterD3pcr1lFieldPcs3Shift)
}

// SetPcs3 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs3Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs3Shift))
}

const (
	RegisterD3pcr1lFieldPcs4Shift = 8
	RegisterD3pcr1lFieldPcs4Mask  = 0x300
)

// GetPcs4 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs4Mask) >> RegisterD3pcr1lFieldPcs4Shift)
}

// SetPcs4 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs4Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs4Shift))
}

const (
	RegisterD3pcr1lFieldPcs5Shift = 10
	RegisterD3pcr1lFieldPcs5Mask  = 0xc00
)

// GetPcs5 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs5Mask) >> RegisterD3pcr1lFieldPcs5Shift)
}

// SetPcs5 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs5Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs5Shift))
}

const (
	RegisterD3pcr1lFieldPcs6Shift = 12
	RegisterD3pcr1lFieldPcs6Mask  = 0x3000
)

// GetPcs6 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs6Mask) >> RegisterD3pcr1lFieldPcs6Shift)
}

// SetPcs6 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs6Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs6Shift))
}

const (
	RegisterD3pcr1lFieldPcs7Shift = 14
	RegisterD3pcr1lFieldPcs7Mask  = 0xc000
)

// GetPcs7 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs7Mask) >> RegisterD3pcr1lFieldPcs7Shift)
}

// SetPcs7 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs7Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs7Shift))
}

const (
	RegisterD3pcr1lFieldPcs8Shift = 16
	RegisterD3pcr1lFieldPcs8Mask  = 0x30000
)

// GetPcs8 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs8() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs8Mask) >> RegisterD3pcr1lFieldPcs8Shift)
}

// SetPcs8 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs8(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs8Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs8Shift))
}

const (
	RegisterD3pcr1lFieldPcs9Shift = 18
	RegisterD3pcr1lFieldPcs9Mask  = 0xc0000
)

// GetPcs9 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs9() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs9Mask) >> RegisterD3pcr1lFieldPcs9Shift)
}

// SetPcs9 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs9(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs9Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs9Shift))
}

const (
	RegisterD3pcr1lFieldPcs10Shift = 20
	RegisterD3pcr1lFieldPcs10Mask  = 0x300000
)

// GetPcs10 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs10() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs10Mask) >> RegisterD3pcr1lFieldPcs10Shift)
}

// SetPcs10 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs10(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs10Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs10Shift))
}

const (
	RegisterD3pcr1lFieldPcs11Shift = 22
	RegisterD3pcr1lFieldPcs11Mask  = 0xc00000
)

// GetPcs11 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs11Mask) >> RegisterD3pcr1lFieldPcs11Shift)
}

// SetPcs11 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs11(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs11Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs11Shift))
}

const (
	RegisterD3pcr1lFieldPcs12Shift = 24
	RegisterD3pcr1lFieldPcs12Mask  = 0x3000000
)

// GetPcs12 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs12() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs12Mask) >> RegisterD3pcr1lFieldPcs12Shift)
}

// SetPcs12 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs12(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs12Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs12Shift))
}

const (
	RegisterD3pcr1lFieldPcs13Shift = 26
	RegisterD3pcr1lFieldPcs13Mask  = 0xc000000
)

// GetPcs13 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs13() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs13Mask) >> RegisterD3pcr1lFieldPcs13Shift)
}

// SetPcs13 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs13(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs13Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs13Shift))
}

const (
	RegisterD3pcr1lFieldPcs14Shift = 28
	RegisterD3pcr1lFieldPcs14Mask  = 0x30000000
)

// GetPcs14 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs14() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs14Mask) >> RegisterD3pcr1lFieldPcs14Shift)
}

// SetPcs14 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs14(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs14Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs14Shift))
}

const (
	RegisterD3pcr1lFieldPcs15Shift = 30
	RegisterD3pcr1lFieldPcs15Mask  = 0xc0000000
)

// GetPcs15 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) GetPcs15() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1lFieldPcs15Mask) >> RegisterD3pcr1lFieldPcs15Shift)
}

// SetPcs15 D3 Pending request clear input signal selection on Event input x = truncate (n/2)
func (r *registerD3pcr1lType) SetPcs15(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1lFieldPcs15Mask)|(uint32(value)<<RegisterD3pcr1lFieldPcs15Shift))
}

// registerD3pcr1hType EXTI D3 pending clear selection register high
type registerD3pcr1hType uint32

const (
	RegisterD3pcr1hFieldPcs19Shift = 6
	RegisterD3pcr1hFieldPcs19Mask  = 0xc0
)

// GetPcs19 D3 Pending request clear input signal selection on Event input x = truncate ((n+32)/2)
func (r *registerD3pcr1hType) GetPcs19() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1hFieldPcs19Mask) >> RegisterD3pcr1hFieldPcs19Shift)
}

// SetPcs19 D3 Pending request clear input signal selection on Event input x = truncate ((n+32)/2)
func (r *registerD3pcr1hType) SetPcs19(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1hFieldPcs19Mask)|(uint32(value)<<RegisterD3pcr1hFieldPcs19Shift))
}

const (
	RegisterD3pcr1hFieldPcs20Shift = 8
	RegisterD3pcr1hFieldPcs20Mask  = 0x300
)

// GetPcs20 D3 Pending request clear input signal selection on Event input x = truncate ((n+32)/2)
func (r *registerD3pcr1hType) GetPcs20() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1hFieldPcs20Mask) >> RegisterD3pcr1hFieldPcs20Shift)
}

// SetPcs20 D3 Pending request clear input signal selection on Event input x = truncate ((n+32)/2)
func (r *registerD3pcr1hType) SetPcs20(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1hFieldPcs20Mask)|(uint32(value)<<RegisterD3pcr1hFieldPcs20Shift))
}

const (
	RegisterD3pcr1hFieldPcs21Shift = 10
	RegisterD3pcr1hFieldPcs21Mask  = 0xc00
)

// GetPcs21 D3 Pending request clear input signal selection on Event input x = truncate ((n+32)/2)
func (r *registerD3pcr1hType) GetPcs21() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1hFieldPcs21Mask) >> RegisterD3pcr1hFieldPcs21Shift)
}

// SetPcs21 D3 Pending request clear input signal selection on Event input x = truncate ((n+32)/2)
func (r *registerD3pcr1hType) SetPcs21(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1hFieldPcs21Mask)|(uint32(value)<<RegisterD3pcr1hFieldPcs21Shift))
}

const (
	RegisterD3pcr1hFieldPcs25Shift = 18
	RegisterD3pcr1hFieldPcs25Mask  = 0xc0000
)

// GetPcs25 D3 Pending request clear input signal selection on Event input x = truncate ((n+32)/2)
func (r *registerD3pcr1hType) GetPcs25() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr1hFieldPcs25Mask) >> RegisterD3pcr1hFieldPcs25Shift)
}

// SetPcs25 D3 Pending request clear input signal selection on Event input x = truncate ((n+32)/2)
func (r *registerD3pcr1hType) SetPcs25(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr1hFieldPcs25Mask)|(uint32(value)<<RegisterD3pcr1hFieldPcs25Shift))
}

// registerRtsr2Type EXTI rising trigger selection register
type registerRtsr2Type uint32

const (
	RegisterRtsr2FieldTr49Shift = 17
	RegisterRtsr2FieldTr49Mask  = 0x20000
)

// GetTr49 Rising trigger event configuration bit of Configurable Event input x+32
func (r *registerRtsr2Type) GetTr49() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr2FieldTr49Mask) != 0
}

// SetTr49 Rising trigger event configuration bit of Configurable Event input x+32
func (r *registerRtsr2Type) SetTr49(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr2FieldTr49Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr2FieldTr49Mask)
	}
}

const (
	RegisterRtsr2FieldTr51Shift = 19
	RegisterRtsr2FieldTr51Mask  = 0x80000
)

// GetTr51 Rising trigger event configuration bit of Configurable Event input x+32
func (r *registerRtsr2Type) GetTr51() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr2FieldTr51Mask) != 0
}

// SetTr51 Rising trigger event configuration bit of Configurable Event input x+32
func (r *registerRtsr2Type) SetTr51(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr2FieldTr51Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr2FieldTr51Mask)
	}
}

// registerFtsr2Type EXTI falling trigger selection register
type registerFtsr2Type uint32

const (
	RegisterFtsr2FieldTr49Shift = 17
	RegisterFtsr2FieldTr49Mask  = 0x20000
)

// GetTr49 Falling trigger event configuration bit of Configurable Event input x+32
func (r *registerFtsr2Type) GetTr49() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr2FieldTr49Mask) != 0
}

// SetTr49 Falling trigger event configuration bit of Configurable Event input x+32
func (r *registerFtsr2Type) SetTr49(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr2FieldTr49Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr2FieldTr49Mask)
	}
}

const (
	RegisterFtsr2FieldTr51Shift = 19
	RegisterFtsr2FieldTr51Mask  = 0x80000
)

// GetTr51 Falling trigger event configuration bit of Configurable Event input x+32
func (r *registerFtsr2Type) GetTr51() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr2FieldTr51Mask) != 0
}

// SetTr51 Falling trigger event configuration bit of Configurable Event input x+32
func (r *registerFtsr2Type) SetTr51(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr2FieldTr51Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr2FieldTr51Mask)
	}
}

// registerSwier2Type EXTI software interrupt event register
type registerSwier2Type uint32

const (
	RegisterSwier2FieldSwier49Shift = 17
	RegisterSwier2FieldSwier49Mask  = 0x20000
)

// GetSwier49 Software interrupt on line x+32
func (r *registerSwier2Type) GetSwier49() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier2FieldSwier49Mask) != 0
}

// SetSwier49 Software interrupt on line x+32
func (r *registerSwier2Type) SetSwier49(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier2FieldSwier49Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier2FieldSwier49Mask)
	}
}

const (
	RegisterSwier2FieldSwier51Shift = 19
	RegisterSwier2FieldSwier51Mask  = 0x80000
)

// GetSwier51 Software interrupt on line x+32
func (r *registerSwier2Type) GetSwier51() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier2FieldSwier51Mask) != 0
}

// SetSwier51 Software interrupt on line x+32
func (r *registerSwier2Type) SetSwier51(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier2FieldSwier51Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier2FieldSwier51Mask)
	}
}

// registerD3pmr2Type EXTI D3 pending mask register
type registerD3pmr2Type uint32

const (
	RegisterD3pmr2FieldMr34Shift = 2
	RegisterD3pmr2FieldMr34Mask  = 0x4
)

// GetMr34 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) GetMr34() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr2FieldMr34Mask) != 0
}

// SetMr34 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) SetMr34(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr2FieldMr34Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr2FieldMr34Mask)
	}
}

const (
	RegisterD3pmr2FieldMr35Shift = 3
	RegisterD3pmr2FieldMr35Mask  = 0x8
)

// GetMr35 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) GetMr35() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr2FieldMr35Mask) != 0
}

// SetMr35 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) SetMr35(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr2FieldMr35Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr2FieldMr35Mask)
	}
}

const (
	RegisterD3pmr2FieldMr41Shift = 9
	RegisterD3pmr2FieldMr41Mask  = 0x200
)

// GetMr41 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) GetMr41() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr2FieldMr41Mask) != 0
}

// SetMr41 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) SetMr41(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr2FieldMr41Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr2FieldMr41Mask)
	}
}

const (
	RegisterD3pmr2FieldMr48Shift = 16
	RegisterD3pmr2FieldMr48Mask  = 0x10000
)

// GetMr48 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) GetMr48() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr2FieldMr48Mask) != 0
}

// SetMr48 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) SetMr48(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr2FieldMr48Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr2FieldMr48Mask)
	}
}

const (
	RegisterD3pmr2FieldMr49Shift = 17
	RegisterD3pmr2FieldMr49Mask  = 0x20000
)

// GetMr49 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) GetMr49() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr2FieldMr49Mask) != 0
}

// SetMr49 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) SetMr49(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr2FieldMr49Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr2FieldMr49Mask)
	}
}

const (
	RegisterD3pmr2FieldMr50Shift = 18
	RegisterD3pmr2FieldMr50Mask  = 0x40000
)

// GetMr50 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) GetMr50() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr2FieldMr50Mask) != 0
}

// SetMr50 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) SetMr50(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr2FieldMr50Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr2FieldMr50Mask)
	}
}

const (
	RegisterD3pmr2FieldMr51Shift = 19
	RegisterD3pmr2FieldMr51Mask  = 0x80000
)

// GetMr51 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) GetMr51() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr2FieldMr51Mask) != 0
}

// SetMr51 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) SetMr51(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr2FieldMr51Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr2FieldMr51Mask)
	}
}

const (
	RegisterD3pmr2FieldMr52Shift = 20
	RegisterD3pmr2FieldMr52Mask  = 0x100000
)

// GetMr52 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) GetMr52() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr2FieldMr52Mask) != 0
}

// SetMr52 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) SetMr52(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr2FieldMr52Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr2FieldMr52Mask)
	}
}

const (
	RegisterD3pmr2FieldMr53Shift = 21
	RegisterD3pmr2FieldMr53Mask  = 0x200000
)

// GetMr53 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) GetMr53() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr2FieldMr53Mask) != 0
}

// SetMr53 D3 Pending Mask on Event input x+32
func (r *registerD3pmr2Type) SetMr53(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr2FieldMr53Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr2FieldMr53Mask)
	}
}

// registerD3pcr2lType EXTI D3 pending clear selection register low
type registerD3pcr2lType uint32

const (
	RegisterD3pcr2lFieldPcs34Shift = 4
	RegisterD3pcr2lFieldPcs34Mask  = 0x30
)

// GetPcs34 D3 Pending request clear input signal selection on Event input x = truncate ((n+64)/2)
func (r *registerD3pcr2lType) GetPcs34() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr2lFieldPcs34Mask) >> RegisterD3pcr2lFieldPcs34Shift)
}

// SetPcs34 D3 Pending request clear input signal selection on Event input x = truncate ((n+64)/2)
func (r *registerD3pcr2lType) SetPcs34(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr2lFieldPcs34Mask)|(uint32(value)<<RegisterD3pcr2lFieldPcs34Shift))
}

const (
	RegisterD3pcr2lFieldPcs35Shift = 6
	RegisterD3pcr2lFieldPcs35Mask  = 0xc0
)

// GetPcs35 D3 Pending request clear input signal selection on Event input x = truncate ((n+64)/2)
func (r *registerD3pcr2lType) GetPcs35() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr2lFieldPcs35Mask) >> RegisterD3pcr2lFieldPcs35Shift)
}

// SetPcs35 D3 Pending request clear input signal selection on Event input x = truncate ((n+64)/2)
func (r *registerD3pcr2lType) SetPcs35(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr2lFieldPcs35Mask)|(uint32(value)<<RegisterD3pcr2lFieldPcs35Shift))
}

const (
	RegisterD3pcr2lFieldPcs41Shift = 18
	RegisterD3pcr2lFieldPcs41Mask  = 0xc0000
)

// GetPcs41 D3 Pending request clear input signal selection on Event input x = truncate ((n+64)/2)
func (r *registerD3pcr2lType) GetPcs41() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr2lFieldPcs41Mask) >> RegisterD3pcr2lFieldPcs41Shift)
}

// SetPcs41 D3 Pending request clear input signal selection on Event input x = truncate ((n+64)/2)
func (r *registerD3pcr2lType) SetPcs41(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr2lFieldPcs41Mask)|(uint32(value)<<RegisterD3pcr2lFieldPcs41Shift))
}

// registerD3pcr2hType EXTI D3 pending clear selection register high
type registerD3pcr2hType uint32

const (
	RegisterD3pcr2hFieldPcs48Shift = 0
	RegisterD3pcr2hFieldPcs48Mask  = 0x3
)

// GetPcs48 Pending request clear input signal selection on Event input x= truncate ((n+96)/2)
func (r *registerD3pcr2hType) GetPcs48() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr2hFieldPcs48Mask) >> RegisterD3pcr2hFieldPcs48Shift)
}

// SetPcs48 Pending request clear input signal selection on Event input x= truncate ((n+96)/2)
func (r *registerD3pcr2hType) SetPcs48(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr2hFieldPcs48Mask)|(uint32(value)<<RegisterD3pcr2hFieldPcs48Shift))
}

const (
	RegisterD3pcr2hFieldPcs49Shift = 2
	RegisterD3pcr2hFieldPcs49Mask  = 0xc
)

// GetPcs49 Pending request clear input signal selection on Event input x= truncate ((n+96)/2)
func (r *registerD3pcr2hType) GetPcs49() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr2hFieldPcs49Mask) >> RegisterD3pcr2hFieldPcs49Shift)
}

// SetPcs49 Pending request clear input signal selection on Event input x= truncate ((n+96)/2)
func (r *registerD3pcr2hType) SetPcs49(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr2hFieldPcs49Mask)|(uint32(value)<<RegisterD3pcr2hFieldPcs49Shift))
}

const (
	RegisterD3pcr2hFieldPcs50Shift = 4
	RegisterD3pcr2hFieldPcs50Mask  = 0x30
)

// GetPcs50 Pending request clear input signal selection on Event input x= truncate ((n+96)/2)
func (r *registerD3pcr2hType) GetPcs50() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr2hFieldPcs50Mask) >> RegisterD3pcr2hFieldPcs50Shift)
}

// SetPcs50 Pending request clear input signal selection on Event input x= truncate ((n+96)/2)
func (r *registerD3pcr2hType) SetPcs50(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr2hFieldPcs50Mask)|(uint32(value)<<RegisterD3pcr2hFieldPcs50Shift))
}

const (
	RegisterD3pcr2hFieldPcs51Shift = 6
	RegisterD3pcr2hFieldPcs51Mask  = 0xc0
)

// GetPcs51 Pending request clear input signal selection on Event input x= truncate ((n+96)/2)
func (r *registerD3pcr2hType) GetPcs51() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr2hFieldPcs51Mask) >> RegisterD3pcr2hFieldPcs51Shift)
}

// SetPcs51 Pending request clear input signal selection on Event input x= truncate ((n+96)/2)
func (r *registerD3pcr2hType) SetPcs51(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr2hFieldPcs51Mask)|(uint32(value)<<RegisterD3pcr2hFieldPcs51Shift))
}

const (
	RegisterD3pcr2hFieldPcs52Shift = 8
	RegisterD3pcr2hFieldPcs52Mask  = 0x300
)

// GetPcs52 Pending request clear input signal selection on Event input x= truncate ((n+96)/2)
func (r *registerD3pcr2hType) GetPcs52() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr2hFieldPcs52Mask) >> RegisterD3pcr2hFieldPcs52Shift)
}

// SetPcs52 Pending request clear input signal selection on Event input x= truncate ((n+96)/2)
func (r *registerD3pcr2hType) SetPcs52(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr2hFieldPcs52Mask)|(uint32(value)<<RegisterD3pcr2hFieldPcs52Shift))
}

const (
	RegisterD3pcr2hFieldPcs53Shift = 10
	RegisterD3pcr2hFieldPcs53Mask  = 0xc00
)

// GetPcs53 Pending request clear input signal selection on Event input x= truncate ((n+96)/2)
func (r *registerD3pcr2hType) GetPcs53() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr2hFieldPcs53Mask) >> RegisterD3pcr2hFieldPcs53Shift)
}

// SetPcs53 Pending request clear input signal selection on Event input x= truncate ((n+96)/2)
func (r *registerD3pcr2hType) SetPcs53(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr2hFieldPcs53Mask)|(uint32(value)<<RegisterD3pcr2hFieldPcs53Shift))
}

// registerRtsr3Type EXTI rising trigger selection register
type registerRtsr3Type uint32

const (
	RegisterRtsr3FieldTr82Shift = 18
	RegisterRtsr3FieldTr82Mask  = 0x40000
)

// GetTr82 Rising trigger event configuration bit of Configurable Event input x+64
func (r *registerRtsr3Type) GetTr82() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr3FieldTr82Mask) != 0
}

// SetTr82 Rising trigger event configuration bit of Configurable Event input x+64
func (r *registerRtsr3Type) SetTr82(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr3FieldTr82Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr3FieldTr82Mask)
	}
}

const (
	RegisterRtsr3FieldTr84Shift = 20
	RegisterRtsr3FieldTr84Mask  = 0x100000
)

// GetTr84 Rising trigger event configuration bit of Configurable Event input x+64
func (r *registerRtsr3Type) GetTr84() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr3FieldTr84Mask) != 0
}

// SetTr84 Rising trigger event configuration bit of Configurable Event input x+64
func (r *registerRtsr3Type) SetTr84(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr3FieldTr84Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr3FieldTr84Mask)
	}
}

const (
	RegisterRtsr3FieldTr85Shift = 21
	RegisterRtsr3FieldTr85Mask  = 0x200000
)

// GetTr85 Rising trigger event configuration bit of Configurable Event input x+64
func (r *registerRtsr3Type) GetTr85() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr3FieldTr85Mask) != 0
}

// SetTr85 Rising trigger event configuration bit of Configurable Event input x+64
func (r *registerRtsr3Type) SetTr85(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr3FieldTr85Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr3FieldTr85Mask)
	}
}

const (
	RegisterRtsr3FieldTr86Shift = 22
	RegisterRtsr3FieldTr86Mask  = 0x400000
)

// GetTr86 Rising trigger event configuration bit of Configurable Event input x+64
func (r *registerRtsr3Type) GetTr86() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtsr3FieldTr86Mask) != 0
}

// SetTr86 Rising trigger event configuration bit of Configurable Event input x+64
func (r *registerRtsr3Type) SetTr86(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtsr3FieldTr86Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtsr3FieldTr86Mask)
	}
}

// registerFtsr3Type EXTI falling trigger selection register
type registerFtsr3Type uint32

const (
	RegisterFtsr3FieldTr82Shift = 18
	RegisterFtsr3FieldTr82Mask  = 0x40000
)

// GetTr82 Falling trigger event configuration bit of Configurable Event input x+64
func (r *registerFtsr3Type) GetTr82() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr3FieldTr82Mask) != 0
}

// SetTr82 Falling trigger event configuration bit of Configurable Event input x+64
func (r *registerFtsr3Type) SetTr82(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr3FieldTr82Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr3FieldTr82Mask)
	}
}

const (
	RegisterFtsr3FieldTr84Shift = 20
	RegisterFtsr3FieldTr84Mask  = 0x100000
)

// GetTr84 Falling trigger event configuration bit of Configurable Event input x+64
func (r *registerFtsr3Type) GetTr84() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr3FieldTr84Mask) != 0
}

// SetTr84 Falling trigger event configuration bit of Configurable Event input x+64
func (r *registerFtsr3Type) SetTr84(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr3FieldTr84Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr3FieldTr84Mask)
	}
}

const (
	RegisterFtsr3FieldTr85Shift = 21
	RegisterFtsr3FieldTr85Mask  = 0x200000
)

// GetTr85 Falling trigger event configuration bit of Configurable Event input x+64
func (r *registerFtsr3Type) GetTr85() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr3FieldTr85Mask) != 0
}

// SetTr85 Falling trigger event configuration bit of Configurable Event input x+64
func (r *registerFtsr3Type) SetTr85(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr3FieldTr85Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr3FieldTr85Mask)
	}
}

const (
	RegisterFtsr3FieldTr86Shift = 22
	RegisterFtsr3FieldTr86Mask  = 0x400000
)

// GetTr86 Falling trigger event configuration bit of Configurable Event input x+64
func (r *registerFtsr3Type) GetTr86() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFtsr3FieldTr86Mask) != 0
}

// SetTr86 Falling trigger event configuration bit of Configurable Event input x+64
func (r *registerFtsr3Type) SetTr86(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFtsr3FieldTr86Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFtsr3FieldTr86Mask)
	}
}

// registerSwier3Type EXTI software interrupt event register
type registerSwier3Type uint32

const (
	RegisterSwier3FieldSwier82Shift = 18
	RegisterSwier3FieldSwier82Mask  = 0x40000
)

// GetSwier82 Software interrupt on line x+64
func (r *registerSwier3Type) GetSwier82() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier3FieldSwier82Mask) != 0
}

// SetSwier82 Software interrupt on line x+64
func (r *registerSwier3Type) SetSwier82(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier3FieldSwier82Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier3FieldSwier82Mask)
	}
}

const (
	RegisterSwier3FieldSwier84Shift = 20
	RegisterSwier3FieldSwier84Mask  = 0x100000
)

// GetSwier84 Software interrupt on line x+64
func (r *registerSwier3Type) GetSwier84() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier3FieldSwier84Mask) != 0
}

// SetSwier84 Software interrupt on line x+64
func (r *registerSwier3Type) SetSwier84(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier3FieldSwier84Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier3FieldSwier84Mask)
	}
}

const (
	RegisterSwier3FieldSwier85Shift = 21
	RegisterSwier3FieldSwier85Mask  = 0x200000
)

// GetSwier85 Software interrupt on line x+64
func (r *registerSwier3Type) GetSwier85() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier3FieldSwier85Mask) != 0
}

// SetSwier85 Software interrupt on line x+64
func (r *registerSwier3Type) SetSwier85(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier3FieldSwier85Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier3FieldSwier85Mask)
	}
}

const (
	RegisterSwier3FieldSwier86Shift = 22
	RegisterSwier3FieldSwier86Mask  = 0x400000
)

// GetSwier86 Software interrupt on line x+64
func (r *registerSwier3Type) GetSwier86() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwier3FieldSwier86Mask) != 0
}

// SetSwier86 Software interrupt on line x+64
func (r *registerSwier3Type) SetSwier86(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwier3FieldSwier86Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwier3FieldSwier86Mask)
	}
}

// registerD3pmr3Type EXTI D3 pending mask register
type registerD3pmr3Type uint32

const (
	RegisterD3pmr3FieldMr88Shift = 24
	RegisterD3pmr3FieldMr88Mask  = 0x1000000
)

// GetMr88 D3 Pending Mask on Event input x+64
func (r *registerD3pmr3Type) GetMr88() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3pmr3FieldMr88Mask) != 0
}

// SetMr88 D3 Pending Mask on Event input x+64
func (r *registerD3pmr3Type) SetMr88(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3pmr3FieldMr88Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3pmr3FieldMr88Mask)
	}
}

// registerD3pcr3hType EXTI D3 pending clear selection register high
type registerD3pcr3hType uint32

const (
	RegisterD3pcr3hFieldPcs88Shift = 18
	RegisterD3pcr3hFieldPcs88Mask  = 0xc0000
)

// GetPcs88 D3 Pending request clear input signal selection on Event input x= truncate N+160/2
func (r *registerD3pcr3hType) GetPcs88() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3pcr3hFieldPcs88Mask) >> RegisterD3pcr3hFieldPcs88Shift)
}

// SetPcs88 D3 Pending request clear input signal selection on Event input x= truncate N+160/2
func (r *registerD3pcr3hType) SetPcs88(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3pcr3hFieldPcs88Mask)|(uint32(value)<<RegisterD3pcr3hFieldPcs88Shift))
}

// registerCpuimr1Type EXTI interrupt mask register
type registerCpuimr1Type uint32

const (
	RegisterCpuimr1FieldMr0Shift = 0
	RegisterCpuimr1FieldMr0Mask  = 0x1
)

// GetMr0 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr0Mask) != 0
}

// SetMr0 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr0Mask)
	}
}

const (
	RegisterCpuimr1FieldMr1Shift = 1
	RegisterCpuimr1FieldMr1Mask  = 0x2
)

// GetMr1 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr1Mask) != 0
}

// SetMr1 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr1Mask)
	}
}

const (
	RegisterCpuimr1FieldMr2Shift = 2
	RegisterCpuimr1FieldMr2Mask  = 0x4
)

// GetMr2 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr2Mask) != 0
}

// SetMr2 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr2Mask)
	}
}

const (
	RegisterCpuimr1FieldMr3Shift = 3
	RegisterCpuimr1FieldMr3Mask  = 0x8
)

// GetMr3 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr3Mask) != 0
}

// SetMr3 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr3Mask)
	}
}

const (
	RegisterCpuimr1FieldMr4Shift = 4
	RegisterCpuimr1FieldMr4Mask  = 0x10
)

// GetMr4 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr4Mask) != 0
}

// SetMr4 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr4Mask)
	}
}

const (
	RegisterCpuimr1FieldMr5Shift = 5
	RegisterCpuimr1FieldMr5Mask  = 0x20
)

// GetMr5 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr5Mask) != 0
}

// SetMr5 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr5Mask)
	}
}

const (
	RegisterCpuimr1FieldMr6Shift = 6
	RegisterCpuimr1FieldMr6Mask  = 0x40
)

// GetMr6 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr6Mask) != 0
}

// SetMr6 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr6Mask)
	}
}

const (
	RegisterCpuimr1FieldMr7Shift = 7
	RegisterCpuimr1FieldMr7Mask  = 0x80
)

// GetMr7 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr7Mask) != 0
}

// SetMr7 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr7Mask)
	}
}

const (
	RegisterCpuimr1FieldMr8Shift = 8
	RegisterCpuimr1FieldMr8Mask  = 0x100
)

// GetMr8 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr8Mask) != 0
}

// SetMr8 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr8Mask)
	}
}

const (
	RegisterCpuimr1FieldMr9Shift = 9
	RegisterCpuimr1FieldMr9Mask  = 0x200
)

// GetMr9 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr9Mask) != 0
}

// SetMr9 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr9Mask)
	}
}

const (
	RegisterCpuimr1FieldMr10Shift = 10
	RegisterCpuimr1FieldMr10Mask  = 0x400
)

// GetMr10 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr10Mask) != 0
}

// SetMr10 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr10Mask)
	}
}

const (
	RegisterCpuimr1FieldMr11Shift = 11
	RegisterCpuimr1FieldMr11Mask  = 0x800
)

// GetMr11 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr11Mask) != 0
}

// SetMr11 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr11Mask)
	}
}

const (
	RegisterCpuimr1FieldMr12Shift = 12
	RegisterCpuimr1FieldMr12Mask  = 0x1000
)

// GetMr12 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr12Mask) != 0
}

// SetMr12 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr12Mask)
	}
}

const (
	RegisterCpuimr1FieldMr13Shift = 13
	RegisterCpuimr1FieldMr13Mask  = 0x2000
)

// GetMr13 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr13Mask) != 0
}

// SetMr13 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr13Mask)
	}
}

const (
	RegisterCpuimr1FieldMr14Shift = 14
	RegisterCpuimr1FieldMr14Mask  = 0x4000
)

// GetMr14 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr14Mask) != 0
}

// SetMr14 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr14Mask)
	}
}

const (
	RegisterCpuimr1FieldMr15Shift = 15
	RegisterCpuimr1FieldMr15Mask  = 0x8000
)

// GetMr15 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr15Mask) != 0
}

// SetMr15 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr15Mask)
	}
}

const (
	RegisterCpuimr1FieldMr16Shift = 16
	RegisterCpuimr1FieldMr16Mask  = 0x10000
)

// GetMr16 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr16Mask) != 0
}

// SetMr16 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr16Mask)
	}
}

const (
	RegisterCpuimr1FieldMr17Shift = 17
	RegisterCpuimr1FieldMr17Mask  = 0x20000
)

// GetMr17 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr17Mask) != 0
}

// SetMr17 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr17Mask)
	}
}

const (
	RegisterCpuimr1FieldMr18Shift = 18
	RegisterCpuimr1FieldMr18Mask  = 0x40000
)

// GetMr18 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr18Mask) != 0
}

// SetMr18 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr18Mask)
	}
}

const (
	RegisterCpuimr1FieldMr19Shift = 19
	RegisterCpuimr1FieldMr19Mask  = 0x80000
)

// GetMr19 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr19Mask) != 0
}

// SetMr19 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr19Mask)
	}
}

const (
	RegisterCpuimr1FieldMr20Shift = 20
	RegisterCpuimr1FieldMr20Mask  = 0x100000
)

// GetMr20 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr20Mask) != 0
}

// SetMr20 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr20Mask)
	}
}

const (
	RegisterCpuimr1FieldMr21Shift = 21
	RegisterCpuimr1FieldMr21Mask  = 0x200000
)

// GetMr21 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr21Mask) != 0
}

// SetMr21 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr21Mask)
	}
}

const (
	RegisterCpuimr1FieldMr22Shift = 22
	RegisterCpuimr1FieldMr22Mask  = 0x400000
)

// GetMr22 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr22() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr22Mask) != 0
}

// SetMr22 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr22(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr22Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr22Mask)
	}
}

const (
	RegisterCpuimr1FieldMr23Shift = 23
	RegisterCpuimr1FieldMr23Mask  = 0x800000
)

// GetMr23 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr23() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr23Mask) != 0
}

// SetMr23 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr23(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr23Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr23Mask)
	}
}

const (
	RegisterCpuimr1FieldMr24Shift = 24
	RegisterCpuimr1FieldMr24Mask  = 0x1000000
)

// GetMr24 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr24() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr24Mask) != 0
}

// SetMr24 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr24(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr24Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr24Mask)
	}
}

const (
	RegisterCpuimr1FieldMr25Shift = 25
	RegisterCpuimr1FieldMr25Mask  = 0x2000000
)

// GetMr25 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr25Mask) != 0
}

// SetMr25 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr25Mask)
	}
}

const (
	RegisterCpuimr1FieldMr26Shift = 26
	RegisterCpuimr1FieldMr26Mask  = 0x4000000
)

// GetMr26 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr26() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr26Mask) != 0
}

// SetMr26 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr26(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr26Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr26Mask)
	}
}

const (
	RegisterCpuimr1FieldMr27Shift = 27
	RegisterCpuimr1FieldMr27Mask  = 0x8000000
)

// GetMr27 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr27() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr27Mask) != 0
}

// SetMr27 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr27(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr27Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr27Mask)
	}
}

const (
	RegisterCpuimr1FieldMr28Shift = 28
	RegisterCpuimr1FieldMr28Mask  = 0x10000000
)

// GetMr28 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr28() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr28Mask) != 0
}

// SetMr28 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr28(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr28Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr28Mask)
	}
}

const (
	RegisterCpuimr1FieldMr29Shift = 29
	RegisterCpuimr1FieldMr29Mask  = 0x20000000
)

// GetMr29 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr29() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr29Mask) != 0
}

// SetMr29 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr29(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr29Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr29Mask)
	}
}

const (
	RegisterCpuimr1FieldMr30Shift = 30
	RegisterCpuimr1FieldMr30Mask  = 0x40000000
)

// GetMr30 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr30() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr30Mask) != 0
}

// SetMr30 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr30(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr30Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr30Mask)
	}
}

const (
	RegisterCpuimr1FieldMr31Shift = 31
	RegisterCpuimr1FieldMr31Mask  = 0x80000000
)

// GetMr31 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) GetMr31() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr1FieldMr31Mask) != 0
}

// SetMr31 Rising trigger event configuration bit of Configurable Event input
func (r *registerCpuimr1Type) SetMr31(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr1FieldMr31Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr1FieldMr31Mask)
	}
}

// registerCpuemr1Type EXTI event mask register
type registerCpuemr1Type uint32

const (
	RegisterCpuemr1FieldMr0Shift = 0
	RegisterCpuemr1FieldMr0Mask  = 0x1
)

// GetMr0 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr0Mask) != 0
}

// SetMr0 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr0Mask)
	}
}

const (
	RegisterCpuemr1FieldMr1Shift = 1
	RegisterCpuemr1FieldMr1Mask  = 0x2
)

// GetMr1 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr1Mask) != 0
}

// SetMr1 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr1Mask)
	}
}

const (
	RegisterCpuemr1FieldMr2Shift = 2
	RegisterCpuemr1FieldMr2Mask  = 0x4
)

// GetMr2 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr2Mask) != 0
}

// SetMr2 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr2Mask)
	}
}

const (
	RegisterCpuemr1FieldMr3Shift = 3
	RegisterCpuemr1FieldMr3Mask  = 0x8
)

// GetMr3 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr3Mask) != 0
}

// SetMr3 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr3Mask)
	}
}

const (
	RegisterCpuemr1FieldMr4Shift = 4
	RegisterCpuemr1FieldMr4Mask  = 0x10
)

// GetMr4 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr4Mask) != 0
}

// SetMr4 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr4Mask)
	}
}

const (
	RegisterCpuemr1FieldMr5Shift = 5
	RegisterCpuemr1FieldMr5Mask  = 0x20
)

// GetMr5 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr5Mask) != 0
}

// SetMr5 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr5Mask)
	}
}

const (
	RegisterCpuemr1FieldMr6Shift = 6
	RegisterCpuemr1FieldMr6Mask  = 0x40
)

// GetMr6 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr6Mask) != 0
}

// SetMr6 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr6Mask)
	}
}

const (
	RegisterCpuemr1FieldMr7Shift = 7
	RegisterCpuemr1FieldMr7Mask  = 0x80
)

// GetMr7 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr7Mask) != 0
}

// SetMr7 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr7Mask)
	}
}

const (
	RegisterCpuemr1FieldMr8Shift = 8
	RegisterCpuemr1FieldMr8Mask  = 0x100
)

// GetMr8 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr8Mask) != 0
}

// SetMr8 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr8Mask)
	}
}

const (
	RegisterCpuemr1FieldMr9Shift = 9
	RegisterCpuemr1FieldMr9Mask  = 0x200
)

// GetMr9 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr9Mask) != 0
}

// SetMr9 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr9Mask)
	}
}

const (
	RegisterCpuemr1FieldMr10Shift = 10
	RegisterCpuemr1FieldMr10Mask  = 0x400
)

// GetMr10 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr10Mask) != 0
}

// SetMr10 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr10Mask)
	}
}

const (
	RegisterCpuemr1FieldMr11Shift = 11
	RegisterCpuemr1FieldMr11Mask  = 0x800
)

// GetMr11 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr11Mask) != 0
}

// SetMr11 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr11Mask)
	}
}

const (
	RegisterCpuemr1FieldMr12Shift = 12
	RegisterCpuemr1FieldMr12Mask  = 0x1000
)

// GetMr12 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr12Mask) != 0
}

// SetMr12 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr12Mask)
	}
}

const (
	RegisterCpuemr1FieldMr13Shift = 13
	RegisterCpuemr1FieldMr13Mask  = 0x2000
)

// GetMr13 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr13Mask) != 0
}

// SetMr13 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr13Mask)
	}
}

const (
	RegisterCpuemr1FieldMr14Shift = 14
	RegisterCpuemr1FieldMr14Mask  = 0x4000
)

// GetMr14 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr14Mask) != 0
}

// SetMr14 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr14Mask)
	}
}

const (
	RegisterCpuemr1FieldMr15Shift = 15
	RegisterCpuemr1FieldMr15Mask  = 0x8000
)

// GetMr15 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr15Mask) != 0
}

// SetMr15 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr15Mask)
	}
}

const (
	RegisterCpuemr1FieldMr16Shift = 16
	RegisterCpuemr1FieldMr16Mask  = 0x10000
)

// GetMr16 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr16Mask) != 0
}

// SetMr16 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr16Mask)
	}
}

const (
	RegisterCpuemr1FieldMr17Shift = 17
	RegisterCpuemr1FieldMr17Mask  = 0x20000
)

// GetMr17 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr17Mask) != 0
}

// SetMr17 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr17Mask)
	}
}

const (
	RegisterCpuemr1FieldMr18Shift = 18
	RegisterCpuemr1FieldMr18Mask  = 0x40000
)

// GetMr18 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr18Mask) != 0
}

// SetMr18 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr18Mask)
	}
}

const (
	RegisterCpuemr1FieldMr19Shift = 19
	RegisterCpuemr1FieldMr19Mask  = 0x80000
)

// GetMr19 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr19Mask) != 0
}

// SetMr19 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr19Mask)
	}
}

const (
	RegisterCpuemr1FieldMr20Shift = 20
	RegisterCpuemr1FieldMr20Mask  = 0x100000
)

// GetMr20 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr20Mask) != 0
}

// SetMr20 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr20Mask)
	}
}

const (
	RegisterCpuemr1FieldMr21Shift = 21
	RegisterCpuemr1FieldMr21Mask  = 0x200000
)

// GetMr21 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr21Mask) != 0
}

// SetMr21 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr21Mask)
	}
}

const (
	RegisterCpuemr1FieldMr22Shift = 22
	RegisterCpuemr1FieldMr22Mask  = 0x400000
)

// GetMr22 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr22() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr22Mask) != 0
}

// SetMr22 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr22(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr22Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr22Mask)
	}
}

const (
	RegisterCpuemr1FieldMr23Shift = 23
	RegisterCpuemr1FieldMr23Mask  = 0x800000
)

// GetMr23 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr23() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr23Mask) != 0
}

// SetMr23 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr23(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr23Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr23Mask)
	}
}

const (
	RegisterCpuemr1FieldMr24Shift = 24
	RegisterCpuemr1FieldMr24Mask  = 0x1000000
)

// GetMr24 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr24() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr24Mask) != 0
}

// SetMr24 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr24(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr24Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr24Mask)
	}
}

const (
	RegisterCpuemr1FieldMr25Shift = 25
	RegisterCpuemr1FieldMr25Mask  = 0x2000000
)

// GetMr25 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr25Mask) != 0
}

// SetMr25 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr25Mask)
	}
}

const (
	RegisterCpuemr1FieldMr26Shift = 26
	RegisterCpuemr1FieldMr26Mask  = 0x4000000
)

// GetMr26 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr26() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr26Mask) != 0
}

// SetMr26 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr26(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr26Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr26Mask)
	}
}

const (
	RegisterCpuemr1FieldMr27Shift = 27
	RegisterCpuemr1FieldMr27Mask  = 0x8000000
)

// GetMr27 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr27() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr27Mask) != 0
}

// SetMr27 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr27(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr27Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr27Mask)
	}
}

const (
	RegisterCpuemr1FieldMr28Shift = 28
	RegisterCpuemr1FieldMr28Mask  = 0x10000000
)

// GetMr28 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr28() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr28Mask) != 0
}

// SetMr28 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr28(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr28Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr28Mask)
	}
}

const (
	RegisterCpuemr1FieldMr29Shift = 29
	RegisterCpuemr1FieldMr29Mask  = 0x20000000
)

// GetMr29 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr29() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr29Mask) != 0
}

// SetMr29 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr29(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr29Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr29Mask)
	}
}

const (
	RegisterCpuemr1FieldMr30Shift = 30
	RegisterCpuemr1FieldMr30Mask  = 0x40000000
)

// GetMr30 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr30() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr30Mask) != 0
}

// SetMr30 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr30(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr30Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr30Mask)
	}
}

const (
	RegisterCpuemr1FieldMr31Shift = 31
	RegisterCpuemr1FieldMr31Mask  = 0x80000000
)

// GetMr31 CPU Event mask on Event input x
func (r *registerCpuemr1Type) GetMr31() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr1FieldMr31Mask) != 0
}

// SetMr31 CPU Event mask on Event input x
func (r *registerCpuemr1Type) SetMr31(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr1FieldMr31Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr1FieldMr31Mask)
	}
}

// registerCpupr1Type EXTI pending register
type registerCpupr1Type uint32

const (
	RegisterCpupr1FieldPr0Shift = 0
	RegisterCpupr1FieldPr0Mask  = 0x1
)

// GetPr0 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr0Mask) != 0
}

// SetPr0 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr0Mask)
	}
}

const (
	RegisterCpupr1FieldPr1Shift = 1
	RegisterCpupr1FieldPr1Mask  = 0x2
)

// GetPr1 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr1Mask) != 0
}

// SetPr1 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr1Mask)
	}
}

const (
	RegisterCpupr1FieldPr2Shift = 2
	RegisterCpupr1FieldPr2Mask  = 0x4
)

// GetPr2 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr2Mask) != 0
}

// SetPr2 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr2Mask)
	}
}

const (
	RegisterCpupr1FieldPr3Shift = 3
	RegisterCpupr1FieldPr3Mask  = 0x8
)

// GetPr3 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr3Mask) != 0
}

// SetPr3 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr3Mask)
	}
}

const (
	RegisterCpupr1FieldPr4Shift = 4
	RegisterCpupr1FieldPr4Mask  = 0x10
)

// GetPr4 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr4Mask) != 0
}

// SetPr4 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr4Mask)
	}
}

const (
	RegisterCpupr1FieldPr5Shift = 5
	RegisterCpupr1FieldPr5Mask  = 0x20
)

// GetPr5 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr5Mask) != 0
}

// SetPr5 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr5Mask)
	}
}

const (
	RegisterCpupr1FieldPr6Shift = 6
	RegisterCpupr1FieldPr6Mask  = 0x40
)

// GetPr6 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr6Mask) != 0
}

// SetPr6 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr6Mask)
	}
}

const (
	RegisterCpupr1FieldPr7Shift = 7
	RegisterCpupr1FieldPr7Mask  = 0x80
)

// GetPr7 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr7Mask) != 0
}

// SetPr7 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr7Mask)
	}
}

const (
	RegisterCpupr1FieldPr8Shift = 8
	RegisterCpupr1FieldPr8Mask  = 0x100
)

// GetPr8 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr8Mask) != 0
}

// SetPr8 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr8Mask)
	}
}

const (
	RegisterCpupr1FieldPr9Shift = 9
	RegisterCpupr1FieldPr9Mask  = 0x200
)

// GetPr9 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr9Mask) != 0
}

// SetPr9 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr9Mask)
	}
}

const (
	RegisterCpupr1FieldPr10Shift = 10
	RegisterCpupr1FieldPr10Mask  = 0x400
)

// GetPr10 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr10Mask) != 0
}

// SetPr10 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr10Mask)
	}
}

const (
	RegisterCpupr1FieldPr11Shift = 11
	RegisterCpupr1FieldPr11Mask  = 0x800
)

// GetPr11 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr11Mask) != 0
}

// SetPr11 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr11Mask)
	}
}

const (
	RegisterCpupr1FieldPr12Shift = 12
	RegisterCpupr1FieldPr12Mask  = 0x1000
)

// GetPr12 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr12Mask) != 0
}

// SetPr12 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr12Mask)
	}
}

const (
	RegisterCpupr1FieldPr13Shift = 13
	RegisterCpupr1FieldPr13Mask  = 0x2000
)

// GetPr13 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr13Mask) != 0
}

// SetPr13 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr13Mask)
	}
}

const (
	RegisterCpupr1FieldPr14Shift = 14
	RegisterCpupr1FieldPr14Mask  = 0x4000
)

// GetPr14 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr14Mask) != 0
}

// SetPr14 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr14Mask)
	}
}

const (
	RegisterCpupr1FieldPr15Shift = 15
	RegisterCpupr1FieldPr15Mask  = 0x8000
)

// GetPr15 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr15Mask) != 0
}

// SetPr15 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr15Mask)
	}
}

const (
	RegisterCpupr1FieldPr16Shift = 16
	RegisterCpupr1FieldPr16Mask  = 0x10000
)

// GetPr16 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr16Mask) != 0
}

// SetPr16 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr16Mask)
	}
}

const (
	RegisterCpupr1FieldPr17Shift = 17
	RegisterCpupr1FieldPr17Mask  = 0x20000
)

// GetPr17 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr17Mask) != 0
}

// SetPr17 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr17Mask)
	}
}

const (
	RegisterCpupr1FieldPr18Shift = 18
	RegisterCpupr1FieldPr18Mask  = 0x40000
)

// GetPr18 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr18Mask) != 0
}

// SetPr18 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr18Mask)
	}
}

const (
	RegisterCpupr1FieldPr19Shift = 19
	RegisterCpupr1FieldPr19Mask  = 0x80000
)

// GetPr19 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr19Mask) != 0
}

// SetPr19 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr19Mask)
	}
}

const (
	RegisterCpupr1FieldPr20Shift = 20
	RegisterCpupr1FieldPr20Mask  = 0x100000
)

// GetPr20 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr20Mask) != 0
}

// SetPr20 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr20Mask)
	}
}

const (
	RegisterCpupr1FieldPr21Shift = 21
	RegisterCpupr1FieldPr21Mask  = 0x200000
)

// GetPr21 CPU Event mask on Event input x
func (r *registerCpupr1Type) GetPr21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr1FieldPr21Mask) != 0
}

// SetPr21 CPU Event mask on Event input x
func (r *registerCpupr1Type) SetPr21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr1FieldPr21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr1FieldPr21Mask)
	}
}

// registerCpuimr2Type EXTI interrupt mask register
type registerCpuimr2Type uint32

const (
	RegisterCpuimr2FieldMr0Shift = 0
	RegisterCpuimr2FieldMr0Mask  = 0x1
)

// GetMr0 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr0Mask) != 0
}

// SetMr0 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr0Mask)
	}
}

const (
	RegisterCpuimr2FieldMr1Shift = 1
	RegisterCpuimr2FieldMr1Mask  = 0x2
)

// GetMr1 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr1Mask) != 0
}

// SetMr1 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr1Mask)
	}
}

const (
	RegisterCpuimr2FieldMr2Shift = 2
	RegisterCpuimr2FieldMr2Mask  = 0x4
)

// GetMr2 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr2Mask) != 0
}

// SetMr2 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr2Mask)
	}
}

const (
	RegisterCpuimr2FieldMr3Shift = 3
	RegisterCpuimr2FieldMr3Mask  = 0x8
)

// GetMr3 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr3Mask) != 0
}

// SetMr3 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr3Mask)
	}
}

const (
	RegisterCpuimr2FieldMr4Shift = 4
	RegisterCpuimr2FieldMr4Mask  = 0x10
)

// GetMr4 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr4Mask) != 0
}

// SetMr4 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr4Mask)
	}
}

const (
	RegisterCpuimr2FieldMr5Shift = 5
	RegisterCpuimr2FieldMr5Mask  = 0x20
)

// GetMr5 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr5Mask) != 0
}

// SetMr5 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr5Mask)
	}
}

const (
	RegisterCpuimr2FieldMr6Shift = 6
	RegisterCpuimr2FieldMr6Mask  = 0x40
)

// GetMr6 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr6Mask) != 0
}

// SetMr6 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr6Mask)
	}
}

const (
	RegisterCpuimr2FieldMr7Shift = 7
	RegisterCpuimr2FieldMr7Mask  = 0x80
)

// GetMr7 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr7Mask) != 0
}

// SetMr7 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr7Mask)
	}
}

const (
	RegisterCpuimr2FieldMr8Shift = 8
	RegisterCpuimr2FieldMr8Mask  = 0x100
)

// GetMr8 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr8Mask) != 0
}

// SetMr8 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr8Mask)
	}
}

const (
	RegisterCpuimr2FieldMr9Shift = 9
	RegisterCpuimr2FieldMr9Mask  = 0x200
)

// GetMr9 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr9Mask) != 0
}

// SetMr9 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr9Mask)
	}
}

const (
	RegisterCpuimr2FieldMr10Shift = 10
	RegisterCpuimr2FieldMr10Mask  = 0x400
)

// GetMr10 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr10Mask) != 0
}

// SetMr10 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr10Mask)
	}
}

const (
	RegisterCpuimr2FieldMr11Shift = 11
	RegisterCpuimr2FieldMr11Mask  = 0x800
)

// GetMr11 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr11Mask) != 0
}

// SetMr11 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr11Mask)
	}
}

const (
	RegisterCpuimr2FieldMr12Shift = 12
	RegisterCpuimr2FieldMr12Mask  = 0x1000
)

// GetMr12 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr12Mask) != 0
}

// SetMr12 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr12Mask)
	}
}

const (
	RegisterCpuimr2FieldMr14Shift = 14
	RegisterCpuimr2FieldMr14Mask  = 0x4000
)

// GetMr14 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr14Mask) != 0
}

// SetMr14 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr14Mask)
	}
}

const (
	RegisterCpuimr2FieldMr15Shift = 15
	RegisterCpuimr2FieldMr15Mask  = 0x8000
)

// GetMr15 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr15Mask) != 0
}

// SetMr15 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr15Mask)
	}
}

const (
	RegisterCpuimr2FieldMr16Shift = 16
	RegisterCpuimr2FieldMr16Mask  = 0x10000
)

// GetMr16 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr16Mask) != 0
}

// SetMr16 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr16Mask)
	}
}

const (
	RegisterCpuimr2FieldMr17Shift = 17
	RegisterCpuimr2FieldMr17Mask  = 0x20000
)

// GetMr17 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr17Mask) != 0
}

// SetMr17 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr17Mask)
	}
}

const (
	RegisterCpuimr2FieldMr18Shift = 18
	RegisterCpuimr2FieldMr18Mask  = 0x40000
)

// GetMr18 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr18Mask) != 0
}

// SetMr18 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr18Mask)
	}
}

const (
	RegisterCpuimr2FieldMr19Shift = 19
	RegisterCpuimr2FieldMr19Mask  = 0x80000
)

// GetMr19 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr19Mask) != 0
}

// SetMr19 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr19Mask)
	}
}

const (
	RegisterCpuimr2FieldMr20Shift = 20
	RegisterCpuimr2FieldMr20Mask  = 0x100000
)

// GetMr20 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr20Mask) != 0
}

// SetMr20 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr20Mask)
	}
}

const (
	RegisterCpuimr2FieldMr21Shift = 21
	RegisterCpuimr2FieldMr21Mask  = 0x200000
)

// GetMr21 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr21Mask) != 0
}

// SetMr21 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr21Mask)
	}
}

const (
	RegisterCpuimr2FieldMr22Shift = 22
	RegisterCpuimr2FieldMr22Mask  = 0x400000
)

// GetMr22 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr22() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr22Mask) != 0
}

// SetMr22 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr22(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr22Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr22Mask)
	}
}

const (
	RegisterCpuimr2FieldMr23Shift = 23
	RegisterCpuimr2FieldMr23Mask  = 0x800000
)

// GetMr23 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr23() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr23Mask) != 0
}

// SetMr23 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr23(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr23Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr23Mask)
	}
}

const (
	RegisterCpuimr2FieldMr24Shift = 24
	RegisterCpuimr2FieldMr24Mask  = 0x1000000
)

// GetMr24 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr24() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr24Mask) != 0
}

// SetMr24 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr24(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr24Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr24Mask)
	}
}

const (
	RegisterCpuimr2FieldMr25Shift = 25
	RegisterCpuimr2FieldMr25Mask  = 0x2000000
)

// GetMr25 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr25Mask) != 0
}

// SetMr25 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr25Mask)
	}
}

const (
	RegisterCpuimr2FieldMr26Shift = 26
	RegisterCpuimr2FieldMr26Mask  = 0x4000000
)

// GetMr26 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr26() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr26Mask) != 0
}

// SetMr26 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr26(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr26Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr26Mask)
	}
}

const (
	RegisterCpuimr2FieldMr27Shift = 27
	RegisterCpuimr2FieldMr27Mask  = 0x8000000
)

// GetMr27 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr27() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr27Mask) != 0
}

// SetMr27 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr27(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr27Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr27Mask)
	}
}

const (
	RegisterCpuimr2FieldMr28Shift = 28
	RegisterCpuimr2FieldMr28Mask  = 0x10000000
)

// GetMr28 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr28() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr28Mask) != 0
}

// SetMr28 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr28(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr28Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr28Mask)
	}
}

const (
	RegisterCpuimr2FieldMr29Shift = 29
	RegisterCpuimr2FieldMr29Mask  = 0x20000000
)

// GetMr29 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr29() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr29Mask) != 0
}

// SetMr29 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr29(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr29Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr29Mask)
	}
}

const (
	RegisterCpuimr2FieldMr30Shift = 30
	RegisterCpuimr2FieldMr30Mask  = 0x40000000
)

// GetMr30 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr30() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr30Mask) != 0
}

// SetMr30 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr30(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr30Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr30Mask)
	}
}

const (
	RegisterCpuimr2FieldMr31Shift = 31
	RegisterCpuimr2FieldMr31Mask  = 0x80000000
)

// GetMr31 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) GetMr31() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr2FieldMr31Mask) != 0
}

// SetMr31 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuimr2Type) SetMr31(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr2FieldMr31Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr2FieldMr31Mask)
	}
}

// registerCpuemr2Type EXTI event mask register
type registerCpuemr2Type uint32

const (
	RegisterCpuemr2FieldMr32Shift = 0
	RegisterCpuemr2FieldMr32Mask  = 0x1
)

// GetMr32 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr32() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr32Mask) != 0
}

// SetMr32 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr32(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr32Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr32Mask)
	}
}

const (
	RegisterCpuemr2FieldMr33Shift = 1
	RegisterCpuemr2FieldMr33Mask  = 0x2
)

// GetMr33 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr33() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr33Mask) != 0
}

// SetMr33 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr33(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr33Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr33Mask)
	}
}

const (
	RegisterCpuemr2FieldMr34Shift = 2
	RegisterCpuemr2FieldMr34Mask  = 0x4
)

// GetMr34 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr34() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr34Mask) != 0
}

// SetMr34 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr34(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr34Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr34Mask)
	}
}

const (
	RegisterCpuemr2FieldMr35Shift = 3
	RegisterCpuemr2FieldMr35Mask  = 0x8
)

// GetMr35 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr35() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr35Mask) != 0
}

// SetMr35 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr35(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr35Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr35Mask)
	}
}

const (
	RegisterCpuemr2FieldMr36Shift = 4
	RegisterCpuemr2FieldMr36Mask  = 0x10
)

// GetMr36 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr36() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr36Mask) != 0
}

// SetMr36 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr36(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr36Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr36Mask)
	}
}

const (
	RegisterCpuemr2FieldMr37Shift = 5
	RegisterCpuemr2FieldMr37Mask  = 0x20
)

// GetMr37 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr37() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr37Mask) != 0
}

// SetMr37 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr37(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr37Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr37Mask)
	}
}

const (
	RegisterCpuemr2FieldMr38Shift = 6
	RegisterCpuemr2FieldMr38Mask  = 0x40
)

// GetMr38 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr38() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr38Mask) != 0
}

// SetMr38 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr38(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr38Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr38Mask)
	}
}

const (
	RegisterCpuemr2FieldMr39Shift = 7
	RegisterCpuemr2FieldMr39Mask  = 0x80
)

// GetMr39 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr39() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr39Mask) != 0
}

// SetMr39 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr39(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr39Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr39Mask)
	}
}

const (
	RegisterCpuemr2FieldMr40Shift = 8
	RegisterCpuemr2FieldMr40Mask  = 0x100
)

// GetMr40 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr40() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr40Mask) != 0
}

// SetMr40 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr40(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr40Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr40Mask)
	}
}

const (
	RegisterCpuemr2FieldMr41Shift = 9
	RegisterCpuemr2FieldMr41Mask  = 0x200
)

// GetMr41 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr41() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr41Mask) != 0
}

// SetMr41 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr41(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr41Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr41Mask)
	}
}

const (
	RegisterCpuemr2FieldMr42Shift = 10
	RegisterCpuemr2FieldMr42Mask  = 0x400
)

// GetMr42 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr42() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr42Mask) != 0
}

// SetMr42 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr42(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr42Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr42Mask)
	}
}

const (
	RegisterCpuemr2FieldMr43Shift = 11
	RegisterCpuemr2FieldMr43Mask  = 0x800
)

// GetMr43 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr43() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr43Mask) != 0
}

// SetMr43 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr43(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr43Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr43Mask)
	}
}

const (
	RegisterCpuemr2FieldMr44Shift = 12
	RegisterCpuemr2FieldMr44Mask  = 0x1000
)

// GetMr44 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr44() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr44Mask) != 0
}

// SetMr44 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr44(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr44Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr44Mask)
	}
}

const (
	RegisterCpuemr2FieldMr46Shift = 14
	RegisterCpuemr2FieldMr46Mask  = 0x4000
)

// GetMr46 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr46() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr46Mask) != 0
}

// SetMr46 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr46(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr46Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr46Mask)
	}
}

const (
	RegisterCpuemr2FieldMr47Shift = 15
	RegisterCpuemr2FieldMr47Mask  = 0x8000
)

// GetMr47 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr47() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr47Mask) != 0
}

// SetMr47 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr47(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr47Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr47Mask)
	}
}

const (
	RegisterCpuemr2FieldMr48Shift = 16
	RegisterCpuemr2FieldMr48Mask  = 0x10000
)

// GetMr48 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr48() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr48Mask) != 0
}

// SetMr48 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr48(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr48Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr48Mask)
	}
}

const (
	RegisterCpuemr2FieldMr49Shift = 17
	RegisterCpuemr2FieldMr49Mask  = 0x20000
)

// GetMr49 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr49() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr49Mask) != 0
}

// SetMr49 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr49(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr49Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr49Mask)
	}
}

const (
	RegisterCpuemr2FieldMr50Shift = 18
	RegisterCpuemr2FieldMr50Mask  = 0x40000
)

// GetMr50 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr50() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr50Mask) != 0
}

// SetMr50 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr50(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr50Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr50Mask)
	}
}

const (
	RegisterCpuemr2FieldMr51Shift = 19
	RegisterCpuemr2FieldMr51Mask  = 0x80000
)

// GetMr51 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr51() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr51Mask) != 0
}

// SetMr51 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr51(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr51Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr51Mask)
	}
}

const (
	RegisterCpuemr2FieldMr52Shift = 20
	RegisterCpuemr2FieldMr52Mask  = 0x100000
)

// GetMr52 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr52() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr52Mask) != 0
}

// SetMr52 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr52(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr52Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr52Mask)
	}
}

const (
	RegisterCpuemr2FieldMr53Shift = 21
	RegisterCpuemr2FieldMr53Mask  = 0x200000
)

// GetMr53 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr53() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr53Mask) != 0
}

// SetMr53 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr53(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr53Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr53Mask)
	}
}

const (
	RegisterCpuemr2FieldMr54Shift = 22
	RegisterCpuemr2FieldMr54Mask  = 0x400000
)

// GetMr54 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr54() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr54Mask) != 0
}

// SetMr54 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr54(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr54Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr54Mask)
	}
}

const (
	RegisterCpuemr2FieldMr55Shift = 23
	RegisterCpuemr2FieldMr55Mask  = 0x800000
)

// GetMr55 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr55() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr55Mask) != 0
}

// SetMr55 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr55(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr55Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr55Mask)
	}
}

const (
	RegisterCpuemr2FieldMr56Shift = 24
	RegisterCpuemr2FieldMr56Mask  = 0x1000000
)

// GetMr56 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr56() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr56Mask) != 0
}

// SetMr56 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr56(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr56Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr56Mask)
	}
}

const (
	RegisterCpuemr2FieldMr57Shift = 25
	RegisterCpuemr2FieldMr57Mask  = 0x2000000
)

// GetMr57 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr57() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr57Mask) != 0
}

// SetMr57 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr57(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr57Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr57Mask)
	}
}

const (
	RegisterCpuemr2FieldMr58Shift = 26
	RegisterCpuemr2FieldMr58Mask  = 0x4000000
)

// GetMr58 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr58() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr58Mask) != 0
}

// SetMr58 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr58(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr58Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr58Mask)
	}
}

const (
	RegisterCpuemr2FieldMr59Shift = 27
	RegisterCpuemr2FieldMr59Mask  = 0x8000000
)

// GetMr59 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr59() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr59Mask) != 0
}

// SetMr59 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr59(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr59Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr59Mask)
	}
}

const (
	RegisterCpuemr2FieldMr60Shift = 28
	RegisterCpuemr2FieldMr60Mask  = 0x10000000
)

// GetMr60 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr60() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr60Mask) != 0
}

// SetMr60 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr60(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr60Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr60Mask)
	}
}

const (
	RegisterCpuemr2FieldMr61Shift = 29
	RegisterCpuemr2FieldMr61Mask  = 0x20000000
)

// GetMr61 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr61() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr61Mask) != 0
}

// SetMr61 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr61(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr61Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr61Mask)
	}
}

const (
	RegisterCpuemr2FieldMr62Shift = 30
	RegisterCpuemr2FieldMr62Mask  = 0x40000000
)

// GetMr62 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr62() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr62Mask) != 0
}

// SetMr62 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr62(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr62Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr62Mask)
	}
}

const (
	RegisterCpuemr2FieldMr63Shift = 31
	RegisterCpuemr2FieldMr63Mask  = 0x80000000
)

// GetMr63 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) GetMr63() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr2FieldMr63Mask) != 0
}

// SetMr63 CPU Interrupt Mask on Direct Event input x+32
func (r *registerCpuemr2Type) SetMr63(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr2FieldMr63Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr2FieldMr63Mask)
	}
}

// registerCpupr2Type EXTI pending register
type registerCpupr2Type uint32

const (
	RegisterCpupr2FieldPr49Shift = 17
	RegisterCpupr2FieldPr49Mask  = 0x20000
)

// GetPr49 Configurable event inputs x+32 Pending bit
func (r *registerCpupr2Type) GetPr49() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr2FieldPr49Mask) != 0
}

// SetPr49 Configurable event inputs x+32 Pending bit
func (r *registerCpupr2Type) SetPr49(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr2FieldPr49Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr2FieldPr49Mask)
	}
}

const (
	RegisterCpupr2FieldPr51Shift = 19
	RegisterCpupr2FieldPr51Mask  = 0x80000
)

// GetPr51 Configurable event inputs x+32 Pending bit
func (r *registerCpupr2Type) GetPr51() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr2FieldPr51Mask) != 0
}

// SetPr51 Configurable event inputs x+32 Pending bit
func (r *registerCpupr2Type) SetPr51(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr2FieldPr51Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr2FieldPr51Mask)
	}
}

// registerCpuimr3Type EXTI interrupt mask register
type registerCpuimr3Type uint32

const (
	RegisterCpuimr3FieldMr64Shift = 0
	RegisterCpuimr3FieldMr64Mask  = 0x1
)

// GetMr64 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr64() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr64Mask) != 0
}

// SetMr64 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr64(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr64Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr64Mask)
	}
}

const (
	RegisterCpuimr3FieldMr65Shift = 1
	RegisterCpuimr3FieldMr65Mask  = 0x2
)

// GetMr65 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr65() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr65Mask) != 0
}

// SetMr65 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr65(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr65Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr65Mask)
	}
}

const (
	RegisterCpuimr3FieldMr66Shift = 2
	RegisterCpuimr3FieldMr66Mask  = 0x4
)

// GetMr66 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr66() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr66Mask) != 0
}

// SetMr66 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr66(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr66Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr66Mask)
	}
}

const (
	RegisterCpuimr3FieldMr67Shift = 3
	RegisterCpuimr3FieldMr67Mask  = 0x8
)

// GetMr67 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr67() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr67Mask) != 0
}

// SetMr67 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr67(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr67Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr67Mask)
	}
}

const (
	RegisterCpuimr3FieldMr68Shift = 4
	RegisterCpuimr3FieldMr68Mask  = 0x10
)

// GetMr68 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr68() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr68Mask) != 0
}

// SetMr68 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr68(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr68Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr68Mask)
	}
}

const (
	RegisterCpuimr3FieldMr69Shift = 5
	RegisterCpuimr3FieldMr69Mask  = 0x20
)

// GetMr69 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr69() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr69Mask) != 0
}

// SetMr69 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr69(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr69Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr69Mask)
	}
}

const (
	RegisterCpuimr3FieldMr70Shift = 6
	RegisterCpuimr3FieldMr70Mask  = 0x40
)

// GetMr70 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr70() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr70Mask) != 0
}

// SetMr70 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr70(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr70Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr70Mask)
	}
}

const (
	RegisterCpuimr3FieldMr71Shift = 7
	RegisterCpuimr3FieldMr71Mask  = 0x80
)

// GetMr71 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr71() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr71Mask) != 0
}

// SetMr71 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr71(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr71Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr71Mask)
	}
}

const (
	RegisterCpuimr3FieldMr72Shift = 8
	RegisterCpuimr3FieldMr72Mask  = 0x100
)

// GetMr72 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr72() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr72Mask) != 0
}

// SetMr72 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr72(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr72Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr72Mask)
	}
}

const (
	RegisterCpuimr3FieldMr73Shift = 9
	RegisterCpuimr3FieldMr73Mask  = 0x200
)

// GetMr73 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr73() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr73Mask) != 0
}

// SetMr73 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr73(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr73Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr73Mask)
	}
}

const (
	RegisterCpuimr3FieldMr74Shift = 10
	RegisterCpuimr3FieldMr74Mask  = 0x400
)

// GetMr74 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr74() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr74Mask) != 0
}

// SetMr74 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr74(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr74Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr74Mask)
	}
}

const (
	RegisterCpuimr3FieldMr75Shift = 11
	RegisterCpuimr3FieldMr75Mask  = 0x800
)

// GetMr75 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr75() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr75Mask) != 0
}

// SetMr75 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr75(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr75Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr75Mask)
	}
}

const (
	RegisterCpuimr3FieldMr76Shift = 12
	RegisterCpuimr3FieldMr76Mask  = 0x1000
)

// GetMr76 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr76() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr76Mask) != 0
}

// SetMr76 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr76(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr76Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr76Mask)
	}
}

const (
	RegisterCpuimr3FieldMr77Shift = 13
	RegisterCpuimr3FieldMr77Mask  = 0x2000
)

// GetMr77 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr77() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr77Mask) != 0
}

// SetMr77 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr77(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr77Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr77Mask)
	}
}

const (
	RegisterCpuimr3FieldMr78Shift = 14
	RegisterCpuimr3FieldMr78Mask  = 0x4000
)

// GetMr78 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr78() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr78Mask) != 0
}

// SetMr78 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr78(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr78Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr78Mask)
	}
}

const (
	RegisterCpuimr3FieldMr79Shift = 15
	RegisterCpuimr3FieldMr79Mask  = 0x8000
)

// GetMr79 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr79() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr79Mask) != 0
}

// SetMr79 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr79(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr79Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr79Mask)
	}
}

const (
	RegisterCpuimr3FieldMr80Shift = 16
	RegisterCpuimr3FieldMr80Mask  = 0x10000
)

// GetMr80 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr80() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr80Mask) != 0
}

// SetMr80 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr80(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr80Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr80Mask)
	}
}

const (
	RegisterCpuimr3FieldMr82Shift = 18
	RegisterCpuimr3FieldMr82Mask  = 0x40000
)

// GetMr82 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr82() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr82Mask) != 0
}

// SetMr82 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr82(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr82Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr82Mask)
	}
}

const (
	RegisterCpuimr3FieldMr84Shift = 20
	RegisterCpuimr3FieldMr84Mask  = 0x100000
)

// GetMr84 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr84() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr84Mask) != 0
}

// SetMr84 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr84(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr84Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr84Mask)
	}
}

const (
	RegisterCpuimr3FieldMr85Shift = 21
	RegisterCpuimr3FieldMr85Mask  = 0x200000
)

// GetMr85 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr85() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr85Mask) != 0
}

// SetMr85 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr85(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr85Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr85Mask)
	}
}

const (
	RegisterCpuimr3FieldMr86Shift = 22
	RegisterCpuimr3FieldMr86Mask  = 0x400000
)

// GetMr86 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr86() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr86Mask) != 0
}

// SetMr86 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr86(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr86Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr86Mask)
	}
}

const (
	RegisterCpuimr3FieldMr87Shift = 23
	RegisterCpuimr3FieldMr87Mask  = 0x800000
)

// GetMr87 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr87() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr87Mask) != 0
}

// SetMr87 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr87(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr87Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr87Mask)
	}
}

const (
	RegisterCpuimr3FieldMr88Shift = 24
	RegisterCpuimr3FieldMr88Mask  = 0x1000000
)

// GetMr88 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) GetMr88() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuimr3FieldMr88Mask) != 0
}

// SetMr88 CPU Interrupt Mask on Direct Event input x+64
func (r *registerCpuimr3Type) SetMr88(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuimr3FieldMr88Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuimr3FieldMr88Mask)
	}
}

// registerCpuemr3Type EXTI event mask register
type registerCpuemr3Type uint32

const (
	RegisterCpuemr3FieldMr64Shift = 0
	RegisterCpuemr3FieldMr64Mask  = 0x1
)

// GetMr64 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr64() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr64Mask) != 0
}

// SetMr64 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr64(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr64Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr64Mask)
	}
}

const (
	RegisterCpuemr3FieldMr65Shift = 1
	RegisterCpuemr3FieldMr65Mask  = 0x2
)

// GetMr65 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr65() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr65Mask) != 0
}

// SetMr65 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr65(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr65Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr65Mask)
	}
}

const (
	RegisterCpuemr3FieldMr66Shift = 2
	RegisterCpuemr3FieldMr66Mask  = 0x4
)

// GetMr66 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr66() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr66Mask) != 0
}

// SetMr66 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr66(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr66Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr66Mask)
	}
}

const (
	RegisterCpuemr3FieldMr67Shift = 3
	RegisterCpuemr3FieldMr67Mask  = 0x8
)

// GetMr67 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr67() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr67Mask) != 0
}

// SetMr67 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr67(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr67Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr67Mask)
	}
}

const (
	RegisterCpuemr3FieldMr68Shift = 4
	RegisterCpuemr3FieldMr68Mask  = 0x10
)

// GetMr68 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr68() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr68Mask) != 0
}

// SetMr68 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr68(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr68Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr68Mask)
	}
}

const (
	RegisterCpuemr3FieldMr69Shift = 5
	RegisterCpuemr3FieldMr69Mask  = 0x20
)

// GetMr69 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr69() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr69Mask) != 0
}

// SetMr69 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr69(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr69Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr69Mask)
	}
}

const (
	RegisterCpuemr3FieldMr70Shift = 6
	RegisterCpuemr3FieldMr70Mask  = 0x40
)

// GetMr70 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr70() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr70Mask) != 0
}

// SetMr70 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr70(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr70Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr70Mask)
	}
}

const (
	RegisterCpuemr3FieldMr71Shift = 7
	RegisterCpuemr3FieldMr71Mask  = 0x80
)

// GetMr71 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr71() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr71Mask) != 0
}

// SetMr71 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr71(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr71Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr71Mask)
	}
}

const (
	RegisterCpuemr3FieldMr72Shift = 8
	RegisterCpuemr3FieldMr72Mask  = 0x100
)

// GetMr72 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr72() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr72Mask) != 0
}

// SetMr72 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr72(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr72Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr72Mask)
	}
}

const (
	RegisterCpuemr3FieldMr73Shift = 9
	RegisterCpuemr3FieldMr73Mask  = 0x200
)

// GetMr73 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr73() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr73Mask) != 0
}

// SetMr73 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr73(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr73Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr73Mask)
	}
}

const (
	RegisterCpuemr3FieldMr74Shift = 10
	RegisterCpuemr3FieldMr74Mask  = 0x400
)

// GetMr74 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr74() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr74Mask) != 0
}

// SetMr74 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr74(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr74Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr74Mask)
	}
}

const (
	RegisterCpuemr3FieldMr75Shift = 11
	RegisterCpuemr3FieldMr75Mask  = 0x800
)

// GetMr75 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr75() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr75Mask) != 0
}

// SetMr75 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr75(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr75Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr75Mask)
	}
}

const (
	RegisterCpuemr3FieldMr76Shift = 12
	RegisterCpuemr3FieldMr76Mask  = 0x1000
)

// GetMr76 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr76() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr76Mask) != 0
}

// SetMr76 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr76(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr76Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr76Mask)
	}
}

const (
	RegisterCpuemr3FieldMr77Shift = 13
	RegisterCpuemr3FieldMr77Mask  = 0x2000
)

// GetMr77 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr77() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr77Mask) != 0
}

// SetMr77 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr77(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr77Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr77Mask)
	}
}

const (
	RegisterCpuemr3FieldMr78Shift = 14
	RegisterCpuemr3FieldMr78Mask  = 0x4000
)

// GetMr78 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr78() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr78Mask) != 0
}

// SetMr78 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr78(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr78Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr78Mask)
	}
}

const (
	RegisterCpuemr3FieldMr79Shift = 15
	RegisterCpuemr3FieldMr79Mask  = 0x8000
)

// GetMr79 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr79() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr79Mask) != 0
}

// SetMr79 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr79(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr79Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr79Mask)
	}
}

const (
	RegisterCpuemr3FieldMr80Shift = 16
	RegisterCpuemr3FieldMr80Mask  = 0x10000
)

// GetMr80 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr80() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr80Mask) != 0
}

// SetMr80 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr80(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr80Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr80Mask)
	}
}

const (
	RegisterCpuemr3FieldMr82Shift = 18
	RegisterCpuemr3FieldMr82Mask  = 0x40000
)

// GetMr82 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr82() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr82Mask) != 0
}

// SetMr82 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr82(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr82Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr82Mask)
	}
}

const (
	RegisterCpuemr3FieldMr84Shift = 20
	RegisterCpuemr3FieldMr84Mask  = 0x100000
)

// GetMr84 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr84() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr84Mask) != 0
}

// SetMr84 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr84(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr84Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr84Mask)
	}
}

const (
	RegisterCpuemr3FieldMr85Shift = 21
	RegisterCpuemr3FieldMr85Mask  = 0x200000
)

// GetMr85 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr85() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr85Mask) != 0
}

// SetMr85 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr85(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr85Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr85Mask)
	}
}

const (
	RegisterCpuemr3FieldMr86Shift = 22
	RegisterCpuemr3FieldMr86Mask  = 0x400000
)

// GetMr86 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr86() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr86Mask) != 0
}

// SetMr86 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr86(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr86Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr86Mask)
	}
}

const (
	RegisterCpuemr3FieldMr87Shift = 23
	RegisterCpuemr3FieldMr87Mask  = 0x800000
)

// GetMr87 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr87() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr87Mask) != 0
}

// SetMr87 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr87(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr87Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr87Mask)
	}
}

const (
	RegisterCpuemr3FieldMr88Shift = 24
	RegisterCpuemr3FieldMr88Mask  = 0x1000000
)

// GetMr88 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) GetMr88() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpuemr3FieldMr88Mask) != 0
}

// SetMr88 CPU Event mask on Event input x+64
func (r *registerCpuemr3Type) SetMr88(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpuemr3FieldMr88Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpuemr3FieldMr88Mask)
	}
}

// registerCpupr3Type EXTI pending register
type registerCpupr3Type uint32

const (
	RegisterCpupr3FieldPr82Shift = 18
	RegisterCpupr3FieldPr82Mask  = 0x40000
)

// GetPr82 Configurable event inputs x+64 Pending bit
func (r *registerCpupr3Type) GetPr82() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr3FieldPr82Mask) != 0
}

// SetPr82 Configurable event inputs x+64 Pending bit
func (r *registerCpupr3Type) SetPr82(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr3FieldPr82Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr3FieldPr82Mask)
	}
}

const (
	RegisterCpupr3FieldPr84Shift = 20
	RegisterCpupr3FieldPr84Mask  = 0x100000
)

// GetPr84 Configurable event inputs x+64 Pending bit
func (r *registerCpupr3Type) GetPr84() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr3FieldPr84Mask) != 0
}

// SetPr84 Configurable event inputs x+64 Pending bit
func (r *registerCpupr3Type) SetPr84(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr3FieldPr84Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr3FieldPr84Mask)
	}
}

const (
	RegisterCpupr3FieldPr85Shift = 21
	RegisterCpupr3FieldPr85Mask  = 0x200000
)

// GetPr85 Configurable event inputs x+64 Pending bit
func (r *registerCpupr3Type) GetPr85() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr3FieldPr85Mask) != 0
}

// SetPr85 Configurable event inputs x+64 Pending bit
func (r *registerCpupr3Type) SetPr85(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr3FieldPr85Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr3FieldPr85Mask)
	}
}

const (
	RegisterCpupr3FieldPr86Shift = 22
	RegisterCpupr3FieldPr86Mask  = 0x400000
)

// GetPr86 Configurable event inputs x+64 Pending bit
func (r *registerCpupr3Type) GetPr86() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpupr3FieldPr86Mask) != 0
}

// SetPr86 Configurable event inputs x+64 Pending bit
func (r *registerCpupr3Type) SetPr86(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpupr3FieldPr86Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpupr3FieldPr86Mask)
	}
}
