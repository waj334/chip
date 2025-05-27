//go:build arm && cortexm

package icb

import (
	"unsafe"
	"volatile"
)

var (
	Icb = (*_icb)(unsafe.Pointer(uintptr(0xe000e004)))
)

type _icb struct {
	Ictr  registerIctrType
	Actlr registerActlrType
	Cppwr registerCppwrType
}

// registerIctrType Provides information about the interrupt controller
type registerIctrType uint32

const (
	RegisterIctrFieldIntlinesnumShift = 0
	RegisterIctrFieldIntlinesnumMask  = 0xf
)

// GetIntlinesnum The total number of interrupt lines supported by an implementation, defined in groups of 32
func (r *registerIctrType) GetIntlinesnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIctrFieldIntlinesnumMask) >> RegisterIctrFieldIntlinesnumShift)
}

// SetIntlinesnum The total number of interrupt lines supported by an implementation, defined in groups of 32
func (r *registerIctrType) SetIntlinesnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIctrFieldIntlinesnumMask)|(uint32(value)<<RegisterIctrFieldIntlinesnumShift))
}

// registerActlrType Allow software to control the processor features and functionality
type registerActlrType uint32

const (
	RegisterActlrFieldFpexcodisShift = 10
	RegisterActlrFieldFpexcodisMask  = 0x400
)

// GetFpexcodis Determines whether floating-point exception outputs are disabled
func (r *registerActlrType) GetFpexcodis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterActlrFieldFpexcodisMask) != 0
}

// SetFpexcodis Determines whether floating-point exception outputs are disabled
func (r *registerActlrType) SetFpexcodis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterActlrFieldFpexcodisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterActlrFieldFpexcodisMask)
	}
}

const (
	RegisterActlrFieldDisnwamodeShift = 11
	RegisterActlrFieldDisnwamodeMask  = 0x800
)

// GetDisnwamode Determines whether no write allocate mode is disabled
func (r *registerActlrType) GetDisnwamode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterActlrFieldDisnwamodeMask) != 0
}

// SetDisnwamode Determines whether no write allocate mode is disabled
func (r *registerActlrType) SetDisnwamode(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterActlrFieldDisnwamodeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterActlrFieldDisnwamodeMask)
	}
}

const (
	RegisterActlrFieldDisitmatbflushShift = 12
	RegisterActlrFieldDisitmatbflushMask  = 0x1000
)

// GetDisitmatbflush Determines whether Instrumentation Trace Macrocell (ITM) or Data Watchpoint and Trace (DWT) ATB flush is disabled
func (r *registerActlrType) GetDisitmatbflush() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterActlrFieldDisitmatbflushMask) != 0
}

// SetDisitmatbflush Determines whether Instrumentation Trace Macrocell (ITM) or Data Watchpoint and Trace (DWT) ATB flush is disabled
func (r *registerActlrType) SetDisitmatbflush(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterActlrFieldDisitmatbflushMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterActlrFieldDisitmatbflushMask)
	}
}

const (
	RegisterActlrFieldEventbusensShift = 13
	RegisterActlrFieldEventbusensMask  = 0x2000
)

// GetEventbusens Accessibility of EVENTBUSEN
func (r *registerActlrType) GetEventbusens() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterActlrFieldEventbusensMask) != 0
}

// SetEventbusens Accessibility of EVENTBUSEN
func (r *registerActlrType) SetEventbusens(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterActlrFieldEventbusensMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterActlrFieldEventbusensMask)
	}
}

const (
	RegisterActlrFieldEventbusenShift = 14
	RegisterActlrFieldEventbusenMask  = 0x4000
)

// GetEventbusen Activate EVENTBUS output
func (r *registerActlrType) GetEventbusen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterActlrFieldEventbusenMask) != 0
}

// SetEventbusen Activate EVENTBUS output
func (r *registerActlrType) SetEventbusen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterActlrFieldEventbusenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterActlrFieldEventbusenMask)
	}
}

const (
	RegisterActlrFieldDiscritaxirurShift = 15
	RegisterActlrFieldDiscritaxirurMask  = 0x8000
)

// GetDiscritaxirur Disable critical AXI Read-Under-Read
func (r *registerActlrType) GetDiscritaxirur() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterActlrFieldDiscritaxirurMask) != 0
}

// SetDiscritaxirur Disable critical AXI Read-Under-Read
func (r *registerActlrType) SetDiscritaxirur(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterActlrFieldDiscritaxirurMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterActlrFieldDiscritaxirurMask)
	}
}

const (
	RegisterActlrFieldDiscritaxiruwShift = 27
	RegisterActlrFieldDiscritaxiruwMask  = 0x8000000
)

// GetDiscritaxiruw Disable-Critical-AXI-Read-Under-Write
func (r *registerActlrType) GetDiscritaxiruw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterActlrFieldDiscritaxiruwMask) != 0
}

// SetDiscritaxiruw Disable-Critical-AXI-Read-Under-Write
func (r *registerActlrType) SetDiscritaxiruw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterActlrFieldDiscritaxiruwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterActlrFieldDiscritaxiruwMask)
	}
}

// registerCppwrType Specifies whether coprocessors are permitted to enter a non-retentive power state
type registerCppwrType uint32

const (
	RegisterCppwrFieldSu0Shift = 0
	RegisterCppwrFieldSu0Mask  = 0x1
)

// GetSu0 This field indicates and allows modification of whether the state associated with coprocessor 0 is permitted to become unknown
func (r *registerCppwrType) GetSu0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSu0Mask) != 0
}

// SetSu0 This field indicates and allows modification of whether the state associated with coprocessor 0 is permitted to become unknown
func (r *registerCppwrType) SetSu0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSu0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSu0Mask)
	}
}

const (
	RegisterCppwrFieldSus0Shift = 1
	RegisterCppwrFieldSus0Mask  = 0x2
)

// GetSus0 This field indicates and allows modification of whether the SU0 field can be modified from Non-secure state
func (r *registerCppwrType) GetSus0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSus0Mask) != 0
}

// SetSus0 This field indicates and allows modification of whether the SU0 field can be modified from Non-secure state
func (r *registerCppwrType) SetSus0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSus0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSus0Mask)
	}
}

const (
	RegisterCppwrFieldSu1Shift = 2
	RegisterCppwrFieldSu1Mask  = 0x4
)

// GetSu1 This field indicates and allows modification of whether the state associated with coprocessor 1 is permitted to become unknown
func (r *registerCppwrType) GetSu1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSu1Mask) != 0
}

// SetSu1 This field indicates and allows modification of whether the state associated with coprocessor 1 is permitted to become unknown
func (r *registerCppwrType) SetSu1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSu1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSu1Mask)
	}
}

const (
	RegisterCppwrFieldSus1Shift = 3
	RegisterCppwrFieldSus1Mask  = 0x8
)

// GetSus1 This field indicates and allows modification of whether the SU1 field can be modified from Non-secure state
func (r *registerCppwrType) GetSus1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSus1Mask) != 0
}

// SetSus1 This field indicates and allows modification of whether the SU1 field can be modified from Non-secure state
func (r *registerCppwrType) SetSus1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSus1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSus1Mask)
	}
}

const (
	RegisterCppwrFieldSu2Shift = 4
	RegisterCppwrFieldSu2Mask  = 0x10
)

// GetSu2 This field indicates and allows modification of whether the state associated with coprocessor 2 is permitted to become unknown
func (r *registerCppwrType) GetSu2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSu2Mask) != 0
}

// SetSu2 This field indicates and allows modification of whether the state associated with coprocessor 2 is permitted to become unknown
func (r *registerCppwrType) SetSu2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSu2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSu2Mask)
	}
}

const (
	RegisterCppwrFieldSus2Shift = 5
	RegisterCppwrFieldSus2Mask  = 0x20
)

// GetSus2 This field indicates and allows modification of whether the SU2 field can be modified from Non-secure state
func (r *registerCppwrType) GetSus2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSus2Mask) != 0
}

// SetSus2 This field indicates and allows modification of whether the SU2 field can be modified from Non-secure state
func (r *registerCppwrType) SetSus2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSus2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSus2Mask)
	}
}

const (
	RegisterCppwrFieldSu3Shift = 6
	RegisterCppwrFieldSu3Mask  = 0x40
)

// GetSu3 This field indicates and allows modification of whether the state associated with coprocessor 3 is permitted to become unknown
func (r *registerCppwrType) GetSu3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSu3Mask) != 0
}

// SetSu3 This field indicates and allows modification of whether the state associated with coprocessor 3 is permitted to become unknown
func (r *registerCppwrType) SetSu3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSu3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSu3Mask)
	}
}

const (
	RegisterCppwrFieldSus3Shift = 7
	RegisterCppwrFieldSus3Mask  = 0x80
)

// GetSus3 This field indicates and allows modification of whether the SU3 field can be modified from Non-secure state
func (r *registerCppwrType) GetSus3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSus3Mask) != 0
}

// SetSus3 This field indicates and allows modification of whether the SU3 field can be modified from Non-secure state
func (r *registerCppwrType) SetSus3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSus3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSus3Mask)
	}
}

const (
	RegisterCppwrFieldSu4Shift = 8
	RegisterCppwrFieldSu4Mask  = 0x100
)

// GetSu4 This field indicates and allows modification of whether the state associated with coprocessor 4 is permitted to become unknown
func (r *registerCppwrType) GetSu4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSu4Mask) != 0
}

// SetSu4 This field indicates and allows modification of whether the state associated with coprocessor 4 is permitted to become unknown
func (r *registerCppwrType) SetSu4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSu4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSu4Mask)
	}
}

const (
	RegisterCppwrFieldSus4Shift = 9
	RegisterCppwrFieldSus4Mask  = 0x200
)

// GetSus4 This field indicates and allows modification of whether the SU4 field can be modified from Non-secure state
func (r *registerCppwrType) GetSus4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSus4Mask) != 0
}

// SetSus4 This field indicates and allows modification of whether the SU4 field can be modified from Non-secure state
func (r *registerCppwrType) SetSus4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSus4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSus4Mask)
	}
}

const (
	RegisterCppwrFieldSu5Shift = 10
	RegisterCppwrFieldSu5Mask  = 0x400
)

// GetSu5 This field indicates and allows modification of whether the state associated with coprocessor 5 is permitted to become unknown
func (r *registerCppwrType) GetSu5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSu5Mask) != 0
}

// SetSu5 This field indicates and allows modification of whether the state associated with coprocessor 5 is permitted to become unknown
func (r *registerCppwrType) SetSu5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSu5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSu5Mask)
	}
}

const (
	RegisterCppwrFieldSus5Shift = 11
	RegisterCppwrFieldSus5Mask  = 0x800
)

// GetSus5 This field indicates and allows modification of whether the SU5 field can be modified from Non-secure state
func (r *registerCppwrType) GetSus5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSus5Mask) != 0
}

// SetSus5 This field indicates and allows modification of whether the SU5 field can be modified from Non-secure state
func (r *registerCppwrType) SetSus5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSus5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSus5Mask)
	}
}

const (
	RegisterCppwrFieldSu6Shift = 12
	RegisterCppwrFieldSu6Mask  = 0x1000
)

// GetSu6 This field indicates and allows modification of whether the state associated with coprocessor 6 is permitted to become unknown
func (r *registerCppwrType) GetSu6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSu6Mask) != 0
}

// SetSu6 This field indicates and allows modification of whether the state associated with coprocessor 6 is permitted to become unknown
func (r *registerCppwrType) SetSu6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSu6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSu6Mask)
	}
}

const (
	RegisterCppwrFieldSus6Shift = 13
	RegisterCppwrFieldSus6Mask  = 0x2000
)

// GetSus6 This field indicates and allows modification of whether the SU6 field can be modified from Non-secure state
func (r *registerCppwrType) GetSus6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSus6Mask) != 0
}

// SetSus6 This field indicates and allows modification of whether the SU6 field can be modified from Non-secure state
func (r *registerCppwrType) SetSus6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSus6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSus6Mask)
	}
}

const (
	RegisterCppwrFieldSu7Shift = 14
	RegisterCppwrFieldSu7Mask  = 0x4000
)

// GetSu7 This field indicates and allows modification of whether the state associated with coprocessor 7 is permitted to become unknown
func (r *registerCppwrType) GetSu7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSu7Mask) != 0
}

// SetSu7 This field indicates and allows modification of whether the state associated with coprocessor 7 is permitted to become unknown
func (r *registerCppwrType) SetSu7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSu7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSu7Mask)
	}
}

const (
	RegisterCppwrFieldSus7Shift = 15
	RegisterCppwrFieldSus7Mask  = 0x8000
)

// GetSus7 This field indicates and allows modification of whether the SU7 field can be modified from Non-secure state
func (r *registerCppwrType) GetSus7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSus7Mask) != 0
}

// SetSus7 This field indicates and allows modification of whether the SU7 field can be modified from Non-secure state
func (r *registerCppwrType) SetSus7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSus7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSus7Mask)
	}
}

const (
	RegisterCppwrFieldSu10Shift = 20
	RegisterCppwrFieldSu10Mask  = 0x100000
)

// GetSu10 This bit indicates and allows modification of whether the state associated with the floating-point and M-profile Vector Extension (MVE) functionality is permitted to become UNKNOWN.
func (r *registerCppwrType) GetSu10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSu10Mask) != 0
}

// SetSu10 This bit indicates and allows modification of whether the state associated with the floating-point and M-profile Vector Extension (MVE) functionality is permitted to become UNKNOWN.
func (r *registerCppwrType) SetSu10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSu10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSu10Mask)
	}
}

const (
	RegisterCppwrFieldSus10Shift = 21
	RegisterCppwrFieldSus10Mask  = 0x200000
)

// GetSus10 This bit indicates and allows modification of whether the SU10 field can be modified from Non-secure state.
func (r *registerCppwrType) GetSus10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSus10Mask) != 0
}

// SetSus10 This bit indicates and allows modification of whether the SU10 field can be modified from Non-secure state.
func (r *registerCppwrType) SetSus10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSus10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSus10Mask)
	}
}

const (
	RegisterCppwrFieldSu11Shift = 22
	RegisterCppwrFieldSu11Mask  = 0x400000
)

// GetSu11 If the value of this bit is not programmed to the same value as the SU10 field, then the value is unknown.
func (r *registerCppwrType) GetSu11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSu11Mask) != 0
}

// SetSu11 If the value of this bit is not programmed to the same value as the SU10 field, then the value is unknown.
func (r *registerCppwrType) SetSu11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSu11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSu11Mask)
	}
}

const (
	RegisterCppwrFieldSus11Shift = 23
	RegisterCppwrFieldSus11Mask  = 0x800000
)

// GetSus11 If the value of this bit is not programmed to the same value as the SUS10 field, then the value is unknown.
func (r *registerCppwrType) GetSus11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCppwrFieldSus11Mask) != 0
}

// SetSus11 If the value of this bit is not programmed to the same value as the SUS10 field, then the value is unknown.
func (r *registerCppwrType) SetSus11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCppwrFieldSus11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCppwrFieldSus11Mask)
	}
}
