package fpu

import (
	"unsafe"
	"volatile"
)

var (
	Fpu = (*_fpu)(unsafe.Pointer(uintptr(0xe000ef34)))
)

type _fpu struct {
	Fpccr registerFpccrType
	Fpcar registerFpcarType
	Fpscr registerFpscrType
}

// registerFpccrType Floating-point context control register
type registerFpccrType uint32

const (
	RegisterFpccrFieldLspactShift = 0
	RegisterFpccrFieldLspactMask  = 0x1
)

// GetLspact LSPACT
func (r *registerFpccrType) GetLspact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldLspactMask) != 0
}

// SetLspact LSPACT
func (r *registerFpccrType) SetLspact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldLspactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldLspactMask)
	}
}

const (
	RegisterFpccrFieldUserShift = 1
	RegisterFpccrFieldUserMask  = 0x2
)

// GetUser USER
func (r *registerFpccrType) GetUser() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldUserMask) != 0
}

// SetUser USER
func (r *registerFpccrType) SetUser(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldUserMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldUserMask)
	}
}

const (
	RegisterFpccrFieldThreadShift = 3
	RegisterFpccrFieldThreadMask  = 0x8
)

// GetThread THREAD
func (r *registerFpccrType) GetThread() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldThreadMask) != 0
}

// SetThread THREAD
func (r *registerFpccrType) SetThread(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldThreadMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldThreadMask)
	}
}

const (
	RegisterFpccrFieldHfrdyShift = 4
	RegisterFpccrFieldHfrdyMask  = 0x10
)

// GetHfrdy HFRDY
func (r *registerFpccrType) GetHfrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldHfrdyMask) != 0
}

// SetHfrdy HFRDY
func (r *registerFpccrType) SetHfrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldHfrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldHfrdyMask)
	}
}

const (
	RegisterFpccrFieldMmrdyShift = 5
	RegisterFpccrFieldMmrdyMask  = 0x20
)

// GetMmrdy MMRDY
func (r *registerFpccrType) GetMmrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldMmrdyMask) != 0
}

// SetMmrdy MMRDY
func (r *registerFpccrType) SetMmrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldMmrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldMmrdyMask)
	}
}

const (
	RegisterFpccrFieldBfrdyShift = 6
	RegisterFpccrFieldBfrdyMask  = 0x40
)

// GetBfrdy BFRDY
func (r *registerFpccrType) GetBfrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldBfrdyMask) != 0
}

// SetBfrdy BFRDY
func (r *registerFpccrType) SetBfrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldBfrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldBfrdyMask)
	}
}

const (
	RegisterFpccrFieldMonrdyShift = 8
	RegisterFpccrFieldMonrdyMask  = 0x100
)

// GetMonrdy MONRDY
func (r *registerFpccrType) GetMonrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldMonrdyMask) != 0
}

// SetMonrdy MONRDY
func (r *registerFpccrType) SetMonrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldMonrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldMonrdyMask)
	}
}

const (
	RegisterFpccrFieldLspenShift = 30
	RegisterFpccrFieldLspenMask  = 0x40000000
)

// GetLspen LSPEN
func (r *registerFpccrType) GetLspen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldLspenMask) != 0
}

// SetLspen LSPEN
func (r *registerFpccrType) SetLspen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldLspenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldLspenMask)
	}
}

const (
	RegisterFpccrFieldAspenShift = 31
	RegisterFpccrFieldAspenMask  = 0x80000000
)

// GetAspen ASPEN
func (r *registerFpccrType) GetAspen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldAspenMask) != 0
}

// SetAspen ASPEN
func (r *registerFpccrType) SetAspen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldAspenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldAspenMask)
	}
}

// registerFpcarType Floating-point context address register
type registerFpcarType uint32

const (
	RegisterFpcarFieldAddressShift = 3
	RegisterFpcarFieldAddressMask  = 0xfffffff8
)

// GetAddress Location of unpopulated floating-point
func (r *registerFpcarType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFpcarFieldAddressMask) >> RegisterFpcarFieldAddressShift)
}

// SetAddress Location of unpopulated floating-point
func (r *registerFpcarType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFpcarFieldAddressMask)|(uint32(value)<<RegisterFpcarFieldAddressShift))
}

// registerFpscrType Floating-point status control register
type registerFpscrType uint32

const (
	RegisterFpscrFieldIocShift = 0
	RegisterFpscrFieldIocMask  = 0x1
)

// GetIoc Invalid operation cumulative exception bit
func (r *registerFpscrType) GetIoc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldIocMask) != 0
}

// SetIoc Invalid operation cumulative exception bit
func (r *registerFpscrType) SetIoc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldIocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldIocMask)
	}
}

const (
	RegisterFpscrFieldDzcShift = 1
	RegisterFpscrFieldDzcMask  = 0x2
)

// GetDzc Division by zero cumulative exception bit.
func (r *registerFpscrType) GetDzc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldDzcMask) != 0
}

// SetDzc Division by zero cumulative exception bit.
func (r *registerFpscrType) SetDzc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldDzcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldDzcMask)
	}
}

const (
	RegisterFpscrFieldOfcShift = 2
	RegisterFpscrFieldOfcMask  = 0x4
)

// GetOfc Overflow cumulative exception bit
func (r *registerFpscrType) GetOfc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldOfcMask) != 0
}

// SetOfc Overflow cumulative exception bit
func (r *registerFpscrType) SetOfc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldOfcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldOfcMask)
	}
}

const (
	RegisterFpscrFieldUfcShift = 3
	RegisterFpscrFieldUfcMask  = 0x8
)

// GetUfc Underflow cumulative exception bit
func (r *registerFpscrType) GetUfc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldUfcMask) != 0
}

// SetUfc Underflow cumulative exception bit
func (r *registerFpscrType) SetUfc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldUfcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldUfcMask)
	}
}

const (
	RegisterFpscrFieldIxcShift = 4
	RegisterFpscrFieldIxcMask  = 0x10
)

// GetIxc Inexact cumulative exception bit
func (r *registerFpscrType) GetIxc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldIxcMask) != 0
}

// SetIxc Inexact cumulative exception bit
func (r *registerFpscrType) SetIxc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldIxcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldIxcMask)
	}
}

const (
	RegisterFpscrFieldIdcShift = 7
	RegisterFpscrFieldIdcMask  = 0x80
)

// GetIdc Input denormal cumulative exception bit.
func (r *registerFpscrType) GetIdc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldIdcMask) != 0
}

// SetIdc Input denormal cumulative exception bit.
func (r *registerFpscrType) SetIdc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldIdcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldIdcMask)
	}
}

const (
	RegisterFpscrFieldRmodeShift = 22
	RegisterFpscrFieldRmodeMask  = 0xc00000
)

// GetRmode Rounding Mode control field
func (r *registerFpscrType) GetRmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldRmodeMask) >> RegisterFpscrFieldRmodeShift)
}

// SetRmode Rounding Mode control field
func (r *registerFpscrType) SetRmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldRmodeMask)|(uint32(value)<<RegisterFpscrFieldRmodeShift))
}

const (
	RegisterFpscrFieldFzShift = 24
	RegisterFpscrFieldFzMask  = 0x1000000
)

// GetFz Flush-to-zero mode control bit:
func (r *registerFpscrType) GetFz() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldFzMask) != 0
}

// SetFz Flush-to-zero mode control bit:
func (r *registerFpscrType) SetFz(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldFzMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldFzMask)
	}
}

const (
	RegisterFpscrFieldDnShift = 25
	RegisterFpscrFieldDnMask  = 0x2000000
)

// GetDn Default NaN mode control bit
func (r *registerFpscrType) GetDn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldDnMask) != 0
}

// SetDn Default NaN mode control bit
func (r *registerFpscrType) SetDn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldDnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldDnMask)
	}
}

const (
	RegisterFpscrFieldAhpShift = 26
	RegisterFpscrFieldAhpMask  = 0x4000000
)

// GetAhp Alternative half-precision control bit
func (r *registerFpscrType) GetAhp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldAhpMask) != 0
}

// SetAhp Alternative half-precision control bit
func (r *registerFpscrType) SetAhp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldAhpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldAhpMask)
	}
}

const (
	RegisterFpscrFieldVShift = 28
	RegisterFpscrFieldVMask  = 0x10000000
)

// GetV Overflow condition code flag
func (r *registerFpscrType) GetV() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldVMask) != 0
}

// SetV Overflow condition code flag
func (r *registerFpscrType) SetV(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldVMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldVMask)
	}
}

const (
	RegisterFpscrFieldCShift = 29
	RegisterFpscrFieldCMask  = 0x20000000
)

// GetC Carry condition code flag
func (r *registerFpscrType) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldCMask) != 0
}

// SetC Carry condition code flag
func (r *registerFpscrType) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldCMask)
	}
}

const (
	RegisterFpscrFieldZShift = 30
	RegisterFpscrFieldZMask  = 0x40000000
)

// GetZ Zero condition code flag
func (r *registerFpscrType) GetZ() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldZMask) != 0
}

// SetZ Zero condition code flag
func (r *registerFpscrType) SetZ(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldZMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldZMask)
	}
}

const (
	RegisterFpscrFieldNShift = 31
	RegisterFpscrFieldNMask  = 0x80000000
)

// GetN Negative condition code flag
func (r *registerFpscrType) GetN() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpscrFieldNMask) != 0
}

// SetN Negative condition code flag
func (r *registerFpscrType) SetN(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpscrFieldNMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpscrFieldNMask)
	}
}
