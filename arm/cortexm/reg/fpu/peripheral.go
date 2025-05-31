//go:build arm && cortexm

package fpu

import (
	"unsafe"
	"volatile"
)

var (
	Fpu = (*_fpu)(unsafe.Pointer(uintptr(0xe000ef34)))
)

type _fpu struct {
	Fpccr  registerFpccrType
	Fpcar  registerFpcarType
	Fpdscr registerFpdscrType
	Mvfr0  registerMvfr0Type
	Mvfr1  registerMvfr1Type
	Mvfr2  registerMvfr2Type
}

// registerFpccrType Holds control data for the Floating Point Unit
type registerFpccrType uint32

const (
	RegisterFpccrFieldLspactShift = 0
	RegisterFpccrFieldLspactMask  = 0x1
)

// GetLspact Lazy state preservation active
func (r *registerFpccrType) GetLspact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldLspactMask) != 0
}

// SetLspact Lazy state preservation active
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

// GetUser Indicates the privilege level of the software executing, when the processor allocated the floating point stack
func (r *registerFpccrType) GetUser() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldUserMask) != 0
}

// SetUser Indicates the privilege level of the software executing, when the processor allocated the floating point stack
func (r *registerFpccrType) SetUser(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldUserMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldUserMask)
	}
}

const (
	RegisterFpccrFieldSShift = 2
	RegisterFpccrFieldSMask  = 0x4
)

// GetS Security status of the floating point context
func (r *registerFpccrType) GetS() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldSMask) != 0
}

// SetS Security status of the floating point context
func (r *registerFpccrType) SetS(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldSMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldSMask)
	}
}

const (
	RegisterFpccrFieldThreadShift = 3
	RegisterFpccrFieldThreadMask  = 0x8
)

// GetThread Thread mode
func (r *registerFpccrType) GetThread() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldThreadMask) != 0
}

// SetThread Thread mode
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

// GetHfrdy HardFault ready
func (r *registerFpccrType) GetHfrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldHfrdyMask) != 0
}

// SetHfrdy HardFault ready
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

// GetMmrdy MemManage ready
func (r *registerFpccrType) GetMmrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldMmrdyMask) != 0
}

// SetMmrdy MemManage ready
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

// GetBfrdy BusFault ready
func (r *registerFpccrType) GetBfrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldBfrdyMask) != 0
}

// SetBfrdy BusFault ready
func (r *registerFpccrType) SetBfrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldBfrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldBfrdyMask)
	}
}

const (
	RegisterFpccrFieldSfrdyShift = 7
	RegisterFpccrFieldSfrdyMask  = 0x80
)

// GetSfrdy SecureFault ready
func (r *registerFpccrType) GetSfrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldSfrdyMask) != 0
}

// SetSfrdy SecureFault ready
func (r *registerFpccrType) SetSfrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldSfrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldSfrdyMask)
	}
}

const (
	RegisterFpccrFieldMonrdyShift = 8
	RegisterFpccrFieldMonrdyMask  = 0x100
)

// GetMonrdy DebugMonitor ready
func (r *registerFpccrType) GetMonrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldMonrdyMask) != 0
}

// SetMonrdy DebugMonitor ready
func (r *registerFpccrType) SetMonrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldMonrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldMonrdyMask)
	}
}

const (
	RegisterFpccrFieldSplimviolShift = 9
	RegisterFpccrFieldSplimviolMask  = 0x200
)

// GetSplimviol Stack pointer limit violation
func (r *registerFpccrType) GetSplimviol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldSplimviolMask) != 0
}

// SetSplimviol Stack pointer limit violation
func (r *registerFpccrType) SetSplimviol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldSplimviolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldSplimviolMask)
	}
}

const (
	RegisterFpccrFieldUfrdyShift = 10
	RegisterFpccrFieldUfrdyMask  = 0x400
)

// GetUfrdy UsageFault ready
func (r *registerFpccrType) GetUfrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldUfrdyMask) != 0
}

// SetUfrdy UsageFault ready
func (r *registerFpccrType) SetUfrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldUfrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldUfrdyMask)
	}
}

const (
	RegisterFpccrFieldTsShift = 26
	RegisterFpccrFieldTsMask  = 0x4000000
)

// GetTs Treat as Secure
func (r *registerFpccrType) GetTs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldTsMask) != 0
}

// SetTs Treat as Secure
func (r *registerFpccrType) SetTs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldTsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldTsMask)
	}
}

const (
	RegisterFpccrFieldClronretsShift = 27
	RegisterFpccrFieldClronretsMask  = 0x8000000
)

// GetClronrets Clear on return Secure only
func (r *registerFpccrType) GetClronrets() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldClronretsMask) != 0
}

// SetClronrets Clear on return Secure only
func (r *registerFpccrType) SetClronrets(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldClronretsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldClronretsMask)
	}
}

const (
	RegisterFpccrFieldClronretShift = 28
	RegisterFpccrFieldClronretMask  = 0x10000000
)

// GetClronret Clear on return
func (r *registerFpccrType) GetClronret() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldClronretMask) != 0
}

// SetClronret Clear on return
func (r *registerFpccrType) SetClronret(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldClronretMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldClronretMask)
	}
}

const (
	RegisterFpccrFieldLspensShift = 29
	RegisterFpccrFieldLspensMask  = 0x20000000
)

// GetLspens Lazy state preservation enable Secure only
func (r *registerFpccrType) GetLspens() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldLspensMask) != 0
}

// SetLspens Lazy state preservation enable Secure only
func (r *registerFpccrType) SetLspens(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldLspensMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldLspensMask)
	}
}

const (
	RegisterFpccrFieldLspenShift = 30
	RegisterFpccrFieldLspenMask  = 0x40000000
)

// GetLspen Automatic state preservation enable
func (r *registerFpccrType) GetLspen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldLspenMask) != 0
}

// SetLspen Automatic state preservation enable
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

// GetAspen Automatic state preservation enable
func (r *registerFpccrType) GetAspen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldAspenMask) != 0
}

// SetAspen Automatic state preservation enable
func (r *registerFpccrType) SetAspen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldAspenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldAspenMask)
	}
}

// registerFpcarType Holds the location of the unpopulated floating-point register space allocated on an exception stack frame
type registerFpcarType uint32

// registerFpdscrType Holds the default values for the floating-point status control data that the processor assigns to the FPSCR when it creates a new floating-point context
type registerFpdscrType uint32

const (
	RegisterFpdscrFieldLtpsizeShift = 16
	RegisterFpdscrFieldLtpsizeMask  = 0x70000
)

// GetLtpsize Default value for FPSCR.LTPSIZE
func (r *registerFpdscrType) GetLtpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFpdscrFieldLtpsizeMask) >> RegisterFpdscrFieldLtpsizeShift)
}

// SetLtpsize Default value for FPSCR.LTPSIZE
func (r *registerFpdscrType) SetLtpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFpdscrFieldLtpsizeMask)|(uint32(value)<<RegisterFpdscrFieldLtpsizeShift))
}

const (
	RegisterFpdscrFieldFz16Shift = 19
	RegisterFpdscrFieldFz16Mask  = 0x80000
)

// GetFz16 Default value for FPSCR.FZ16
func (r *registerFpdscrType) GetFz16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpdscrFieldFz16Mask) != 0
}

// SetFz16 Default value for FPSCR.FZ16
func (r *registerFpdscrType) SetFz16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpdscrFieldFz16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpdscrFieldFz16Mask)
	}
}

const (
	RegisterFpdscrFieldRmodeShift = 22
	RegisterFpdscrFieldRmodeMask  = 0xc00000
)

// GetRmode Default value for FPSCR.RMode
func (r *registerFpdscrType) GetRmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFpdscrFieldRmodeMask) >> RegisterFpdscrFieldRmodeShift)
}

// SetRmode Default value for FPSCR.RMode
func (r *registerFpdscrType) SetRmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFpdscrFieldRmodeMask)|(uint32(value)<<RegisterFpdscrFieldRmodeShift))
}

const (
	RegisterFpdscrFieldFzShift = 24
	RegisterFpdscrFieldFzMask  = 0x1000000
)

// GetFz Default value for FPSCR.FZ
func (r *registerFpdscrType) GetFz() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpdscrFieldFzMask) != 0
}

// SetFz Default value for FPSCR.FZ
func (r *registerFpdscrType) SetFz(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpdscrFieldFzMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpdscrFieldFzMask)
	}
}

const (
	RegisterFpdscrFieldDnShift = 25
	RegisterFpdscrFieldDnMask  = 0x2000000
)

// GetDn Default value for FPSCR.DN
func (r *registerFpdscrType) GetDn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpdscrFieldDnMask) != 0
}

// SetDn Default value for FPSCR.DN
func (r *registerFpdscrType) SetDn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpdscrFieldDnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpdscrFieldDnMask)
	}
}

const (
	RegisterFpdscrFieldAhpShift = 26
	RegisterFpdscrFieldAhpMask  = 0x4000000
)

// GetAhp Default value for FPSCR.AHP
func (r *registerFpdscrType) GetAhp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpdscrFieldAhpMask) != 0
}

// SetAhp Default value for FPSCR.AHP
func (r *registerFpdscrType) SetAhp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpdscrFieldAhpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpdscrFieldAhpMask)
	}
}

// registerMvfr0Type Describes the features provided by the Floating-point extension
type registerMvfr0Type uint32

const (
	RegisterMvfr0FieldSimdregShift = 0
	RegisterMvfr0FieldSimdregMask  = 0xf
)

// GetSimdreg SIMD registers
func (r *registerMvfr0Type) GetSimdreg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr0FieldSimdregMask) >> RegisterMvfr0FieldSimdregShift)
}

// SetSimdreg SIMD registers
func (r *registerMvfr0Type) SetSimdreg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr0FieldSimdregMask)|(uint32(value)<<RegisterMvfr0FieldSimdregShift))
}

const (
	RegisterMvfr0FieldFpspShift = 4
	RegisterMvfr0FieldFpspMask  = 0xf0
)

// GetFpsp Floating-point single-precision
func (r *registerMvfr0Type) GetFpsp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr0FieldFpspMask) >> RegisterMvfr0FieldFpspShift)
}

// SetFpsp Floating-point single-precision
func (r *registerMvfr0Type) SetFpsp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr0FieldFpspMask)|(uint32(value)<<RegisterMvfr0FieldFpspShift))
}

const (
	RegisterMvfr0FieldFpdpShift = 8
	RegisterMvfr0FieldFpdpMask  = 0xf00
)

// GetFpdp Floating-point double-precision
func (r *registerMvfr0Type) GetFpdp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr0FieldFpdpMask) >> RegisterMvfr0FieldFpdpShift)
}

// SetFpdp Floating-point double-precision
func (r *registerMvfr0Type) SetFpdp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr0FieldFpdpMask)|(uint32(value)<<RegisterMvfr0FieldFpdpShift))
}

const (
	RegisterMvfr0FieldFpdivideShift = 16
	RegisterMvfr0FieldFpdivideMask  = 0xf0000
)

// GetFpdivide Floating-point divide
func (r *registerMvfr0Type) GetFpdivide() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr0FieldFpdivideMask) >> RegisterMvfr0FieldFpdivideShift)
}

// SetFpdivide Floating-point divide
func (r *registerMvfr0Type) SetFpdivide(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr0FieldFpdivideMask)|(uint32(value)<<RegisterMvfr0FieldFpdivideShift))
}

const (
	RegisterMvfr0FieldFpsqrtShift = 20
	RegisterMvfr0FieldFpsqrtMask  = 0xf00000
)

// GetFpsqrt Floating-point square root
func (r *registerMvfr0Type) GetFpsqrt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr0FieldFpsqrtMask) >> RegisterMvfr0FieldFpsqrtShift)
}

// SetFpsqrt Floating-point square root
func (r *registerMvfr0Type) SetFpsqrt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr0FieldFpsqrtMask)|(uint32(value)<<RegisterMvfr0FieldFpsqrtShift))
}

const (
	RegisterMvfr0FieldFproundShift = 28
	RegisterMvfr0FieldFproundMask  = 0xf0000000
)

// GetFpround Floating-point rounding modes
func (r *registerMvfr0Type) GetFpround() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr0FieldFproundMask) >> RegisterMvfr0FieldFproundShift)
}

// SetFpround Floating-point rounding modes
func (r *registerMvfr0Type) SetFpround(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr0FieldFproundMask)|(uint32(value)<<RegisterMvfr0FieldFproundShift))
}

// registerMvfr1Type Describes the features provided by the Floating-point extension
type registerMvfr1Type uint32

const (
	RegisterMvfr1FieldFpftzShift = 0
	RegisterMvfr1FieldFpftzMask  = 0xf
)

// GetFpftz Floating-point flush-to-zero
func (r *registerMvfr1Type) GetFpftz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr1FieldFpftzMask) >> RegisterMvfr1FieldFpftzShift)
}

// SetFpftz Floating-point flush-to-zero
func (r *registerMvfr1Type) SetFpftz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr1FieldFpftzMask)|(uint32(value)<<RegisterMvfr1FieldFpftzShift))
}

const (
	RegisterMvfr1FieldFpdnanShift = 4
	RegisterMvfr1FieldFpdnanMask  = 0xf0
)

// GetFpdnan Floating-point default NaN
func (r *registerMvfr1Type) GetFpdnan() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr1FieldFpdnanMask) >> RegisterMvfr1FieldFpdnanShift)
}

// SetFpdnan Floating-point default NaN
func (r *registerMvfr1Type) SetFpdnan(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr1FieldFpdnanMask)|(uint32(value)<<RegisterMvfr1FieldFpdnanShift))
}

const (
	RegisterMvfr1FieldMveShift = 8
	RegisterMvfr1FieldMveMask  = 0xf00
)

// GetMve Indicates support for MVE
func (r *registerMvfr1Type) GetMve() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr1FieldMveMask) >> RegisterMvfr1FieldMveShift)
}

// SetMve Indicates support for MVE
func (r *registerMvfr1Type) SetMve(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr1FieldMveMask)|(uint32(value)<<RegisterMvfr1FieldMveShift))
}

const (
	RegisterMvfr1FieldFp16Shift = 20
	RegisterMvfr1FieldFp16Mask  = 0xf00000
)

// GetFp16 Floating-point half-precision data processing
func (r *registerMvfr1Type) GetFp16() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr1FieldFp16Mask) >> RegisterMvfr1FieldFp16Shift)
}

// SetFp16 Floating-point half-precision data processing
func (r *registerMvfr1Type) SetFp16(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr1FieldFp16Mask)|(uint32(value)<<RegisterMvfr1FieldFp16Shift))
}

const (
	RegisterMvfr1FieldFphpShift = 24
	RegisterMvfr1FieldFphpMask  = 0xf000000
)

// GetFphp Floating-point half-precision conversion
func (r *registerMvfr1Type) GetFphp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr1FieldFphpMask) >> RegisterMvfr1FieldFphpShift)
}

// SetFphp Floating-point half-precision conversion
func (r *registerMvfr1Type) SetFphp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr1FieldFphpMask)|(uint32(value)<<RegisterMvfr1FieldFphpShift))
}

const (
	RegisterMvfr1FieldFmacShift = 28
	RegisterMvfr1FieldFmacMask  = 0xf0000000
)

// GetFmac Fused multiply accumulate instructions
func (r *registerMvfr1Type) GetFmac() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr1FieldFmacMask) >> RegisterMvfr1FieldFmacShift)
}

// SetFmac Fused multiply accumulate instructions
func (r *registerMvfr1Type) SetFmac(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr1FieldFmacMask)|(uint32(value)<<RegisterMvfr1FieldFmacShift))
}

// registerMvfr2Type Describes the features provided by the floating-point extension
type registerMvfr2Type uint32

const (
	RegisterMvfr2FieldFpmiscShift = 4
	RegisterMvfr2FieldFpmiscMask  = 0xf0
)

// GetFpmisc Floating-point miscellaneous
func (r *registerMvfr2Type) GetFpmisc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr2FieldFpmiscMask) >> RegisterMvfr2FieldFpmiscShift)
}

// SetFpmisc Floating-point miscellaneous
func (r *registerMvfr2Type) SetFpmisc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr2FieldFpmiscMask)|(uint32(value)<<RegisterMvfr2FieldFpmiscShift))
}
