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
	Fpccr  RegisterFpccrType
	Fpcar  RegisterFpcarType
	Fpdscr RegisterFpdscrType
	Mvfr0  RegisterMvfr0Type
	Mvfr1  RegisterMvfr1Type
	Mvfr2  RegisterMvfr2Type
}

// RegisterFpccrType Holds control data for the Floating Point Unit
type RegisterFpccrType uint32

func (r *RegisterFpccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterFpccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterFpccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterFpccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterFpccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterFpccrFieldLspactShift = 0
	RegisterFpccrFieldLspactMask  = 0x1
)

// GetLspact Lazy state preservation active
func (r *RegisterFpccrType) GetLspact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldLspactMask) != 0
}

// SetLspact Lazy state preservation active
func (r *RegisterFpccrType) SetLspact(value bool) {
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
func (r *RegisterFpccrType) GetUser() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldUserMask) != 0
}

// SetUser Indicates the privilege level of the software executing, when the processor allocated the floating point stack
func (r *RegisterFpccrType) SetUser(value bool) {
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
func (r *RegisterFpccrType) GetS() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldSMask) != 0
}

// SetS Security status of the floating point context
func (r *RegisterFpccrType) SetS(value bool) {
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
func (r *RegisterFpccrType) GetThread() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldThreadMask) != 0
}

// SetThread Thread mode
func (r *RegisterFpccrType) SetThread(value bool) {
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
func (r *RegisterFpccrType) GetHfrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldHfrdyMask) != 0
}

// SetHfrdy HardFault ready
func (r *RegisterFpccrType) SetHfrdy(value bool) {
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
func (r *RegisterFpccrType) GetMmrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldMmrdyMask) != 0
}

// SetMmrdy MemManage ready
func (r *RegisterFpccrType) SetMmrdy(value bool) {
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
func (r *RegisterFpccrType) GetBfrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldBfrdyMask) != 0
}

// SetBfrdy BusFault ready
func (r *RegisterFpccrType) SetBfrdy(value bool) {
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
func (r *RegisterFpccrType) GetSfrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldSfrdyMask) != 0
}

// SetSfrdy SecureFault ready
func (r *RegisterFpccrType) SetSfrdy(value bool) {
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
func (r *RegisterFpccrType) GetMonrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldMonrdyMask) != 0
}

// SetMonrdy DebugMonitor ready
func (r *RegisterFpccrType) SetMonrdy(value bool) {
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
func (r *RegisterFpccrType) GetSplimviol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldSplimviolMask) != 0
}

// SetSplimviol Stack pointer limit violation
func (r *RegisterFpccrType) SetSplimviol(value bool) {
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
func (r *RegisterFpccrType) GetUfrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldUfrdyMask) != 0
}

// SetUfrdy UsageFault ready
func (r *RegisterFpccrType) SetUfrdy(value bool) {
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
func (r *RegisterFpccrType) GetTs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldTsMask) != 0
}

// SetTs Treat as Secure
func (r *RegisterFpccrType) SetTs(value bool) {
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
func (r *RegisterFpccrType) GetClronrets() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldClronretsMask) != 0
}

// SetClronrets Clear on return Secure only
func (r *RegisterFpccrType) SetClronrets(value bool) {
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
func (r *RegisterFpccrType) GetClronret() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldClronretMask) != 0
}

// SetClronret Clear on return
func (r *RegisterFpccrType) SetClronret(value bool) {
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
func (r *RegisterFpccrType) GetLspens() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldLspensMask) != 0
}

// SetLspens Lazy state preservation enable Secure only
func (r *RegisterFpccrType) SetLspens(value bool) {
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
func (r *RegisterFpccrType) GetLspen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldLspenMask) != 0
}

// SetLspen Automatic state preservation enable
func (r *RegisterFpccrType) SetLspen(value bool) {
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
func (r *RegisterFpccrType) GetAspen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpccrFieldAspenMask) != 0
}

// SetAspen Automatic state preservation enable
func (r *RegisterFpccrType) SetAspen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpccrFieldAspenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpccrFieldAspenMask)
	}
}

// RegisterFpcarType Holds the location of the unpopulated floating-point register space allocated on an exception stack frame
type RegisterFpcarType uint32

func (r *RegisterFpcarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterFpcarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterFpcarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterFpcarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterFpcarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterFpdscrType Holds the default values for the floating-point status control data that the processor assigns to the FPSCR when it creates a new floating-point context
type RegisterFpdscrType uint32

func (r *RegisterFpdscrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterFpdscrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterFpdscrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterFpdscrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterFpdscrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterFpdscrFieldLtpsizeShift = 16
	RegisterFpdscrFieldLtpsizeMask  = 0x70000
)

// GetLtpsize Default value for FPSCR.LTPSIZE
func (r *RegisterFpdscrType) GetLtpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFpdscrFieldLtpsizeMask) >> RegisterFpdscrFieldLtpsizeShift)
}

// SetLtpsize Default value for FPSCR.LTPSIZE
func (r *RegisterFpdscrType) SetLtpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFpdscrFieldLtpsizeMask)|(uint32(value)<<RegisterFpdscrFieldLtpsizeShift))
}

const (
	RegisterFpdscrFieldFz16Shift = 19
	RegisterFpdscrFieldFz16Mask  = 0x80000
)

// GetFz16 Default value for FPSCR.FZ16
func (r *RegisterFpdscrType) GetFz16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpdscrFieldFz16Mask) != 0
}

// SetFz16 Default value for FPSCR.FZ16
func (r *RegisterFpdscrType) SetFz16(value bool) {
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
func (r *RegisterFpdscrType) GetRmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFpdscrFieldRmodeMask) >> RegisterFpdscrFieldRmodeShift)
}

// SetRmode Default value for FPSCR.RMode
func (r *RegisterFpdscrType) SetRmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFpdscrFieldRmodeMask)|(uint32(value)<<RegisterFpdscrFieldRmodeShift))
}

const (
	RegisterFpdscrFieldFzShift = 24
	RegisterFpdscrFieldFzMask  = 0x1000000
)

// GetFz Default value for FPSCR.FZ
func (r *RegisterFpdscrType) GetFz() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpdscrFieldFzMask) != 0
}

// SetFz Default value for FPSCR.FZ
func (r *RegisterFpdscrType) SetFz(value bool) {
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
func (r *RegisterFpdscrType) GetDn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpdscrFieldDnMask) != 0
}

// SetDn Default value for FPSCR.DN
func (r *RegisterFpdscrType) SetDn(value bool) {
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
func (r *RegisterFpdscrType) GetAhp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFpdscrFieldAhpMask) != 0
}

// SetAhp Default value for FPSCR.AHP
func (r *RegisterFpdscrType) SetAhp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFpdscrFieldAhpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFpdscrFieldAhpMask)
	}
}

// RegisterMvfr0Type Describes the features provided by the Floating-point extension
type RegisterMvfr0Type uint32

func (r *RegisterMvfr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMvfr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMvfr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMvfr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMvfr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMvfr0FieldSimdregShift = 0
	RegisterMvfr0FieldSimdregMask  = 0xf
)

// GetSimdreg SIMD registers
func (r *RegisterMvfr0Type) GetSimdreg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr0FieldSimdregMask) >> RegisterMvfr0FieldSimdregShift)
}

// SetSimdreg SIMD registers
func (r *RegisterMvfr0Type) SetSimdreg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr0FieldSimdregMask)|(uint32(value)<<RegisterMvfr0FieldSimdregShift))
}

const (
	RegisterMvfr0FieldFpspShift = 4
	RegisterMvfr0FieldFpspMask  = 0xf0
)

// GetFpsp Floating-point single-precision
func (r *RegisterMvfr0Type) GetFpsp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr0FieldFpspMask) >> RegisterMvfr0FieldFpspShift)
}

// SetFpsp Floating-point single-precision
func (r *RegisterMvfr0Type) SetFpsp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr0FieldFpspMask)|(uint32(value)<<RegisterMvfr0FieldFpspShift))
}

const (
	RegisterMvfr0FieldFpdpShift = 8
	RegisterMvfr0FieldFpdpMask  = 0xf00
)

// GetFpdp Floating-point double-precision
func (r *RegisterMvfr0Type) GetFpdp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr0FieldFpdpMask) >> RegisterMvfr0FieldFpdpShift)
}

// SetFpdp Floating-point double-precision
func (r *RegisterMvfr0Type) SetFpdp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr0FieldFpdpMask)|(uint32(value)<<RegisterMvfr0FieldFpdpShift))
}

const (
	RegisterMvfr0FieldFpdivideShift = 16
	RegisterMvfr0FieldFpdivideMask  = 0xf0000
)

// GetFpdivide Floating-point divide
func (r *RegisterMvfr0Type) GetFpdivide() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr0FieldFpdivideMask) >> RegisterMvfr0FieldFpdivideShift)
}

// SetFpdivide Floating-point divide
func (r *RegisterMvfr0Type) SetFpdivide(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr0FieldFpdivideMask)|(uint32(value)<<RegisterMvfr0FieldFpdivideShift))
}

const (
	RegisterMvfr0FieldFpsqrtShift = 20
	RegisterMvfr0FieldFpsqrtMask  = 0xf00000
)

// GetFpsqrt Floating-point square root
func (r *RegisterMvfr0Type) GetFpsqrt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr0FieldFpsqrtMask) >> RegisterMvfr0FieldFpsqrtShift)
}

// SetFpsqrt Floating-point square root
func (r *RegisterMvfr0Type) SetFpsqrt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr0FieldFpsqrtMask)|(uint32(value)<<RegisterMvfr0FieldFpsqrtShift))
}

const (
	RegisterMvfr0FieldFproundShift = 28
	RegisterMvfr0FieldFproundMask  = 0xf0000000
)

// GetFpround Floating-point rounding modes
func (r *RegisterMvfr0Type) GetFpround() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr0FieldFproundMask) >> RegisterMvfr0FieldFproundShift)
}

// SetFpround Floating-point rounding modes
func (r *RegisterMvfr0Type) SetFpround(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr0FieldFproundMask)|(uint32(value)<<RegisterMvfr0FieldFproundShift))
}

// RegisterMvfr1Type Describes the features provided by the Floating-point extension
type RegisterMvfr1Type uint32

func (r *RegisterMvfr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMvfr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMvfr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMvfr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMvfr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMvfr1FieldFpftzShift = 0
	RegisterMvfr1FieldFpftzMask  = 0xf
)

// GetFpftz Floating-point flush-to-zero
func (r *RegisterMvfr1Type) GetFpftz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr1FieldFpftzMask) >> RegisterMvfr1FieldFpftzShift)
}

// SetFpftz Floating-point flush-to-zero
func (r *RegisterMvfr1Type) SetFpftz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr1FieldFpftzMask)|(uint32(value)<<RegisterMvfr1FieldFpftzShift))
}

const (
	RegisterMvfr1FieldFpdnanShift = 4
	RegisterMvfr1FieldFpdnanMask  = 0xf0
)

// GetFpdnan Floating-point default NaN
func (r *RegisterMvfr1Type) GetFpdnan() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr1FieldFpdnanMask) >> RegisterMvfr1FieldFpdnanShift)
}

// SetFpdnan Floating-point default NaN
func (r *RegisterMvfr1Type) SetFpdnan(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr1FieldFpdnanMask)|(uint32(value)<<RegisterMvfr1FieldFpdnanShift))
}

const (
	RegisterMvfr1FieldMveShift = 8
	RegisterMvfr1FieldMveMask  = 0xf00
)

// GetMve Indicates support for MVE
func (r *RegisterMvfr1Type) GetMve() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr1FieldMveMask) >> RegisterMvfr1FieldMveShift)
}

// SetMve Indicates support for MVE
func (r *RegisterMvfr1Type) SetMve(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr1FieldMveMask)|(uint32(value)<<RegisterMvfr1FieldMveShift))
}

const (
	RegisterMvfr1FieldFp16Shift = 20
	RegisterMvfr1FieldFp16Mask  = 0xf00000
)

// GetFp16 Floating-point half-precision data processing
func (r *RegisterMvfr1Type) GetFp16() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr1FieldFp16Mask) >> RegisterMvfr1FieldFp16Shift)
}

// SetFp16 Floating-point half-precision data processing
func (r *RegisterMvfr1Type) SetFp16(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr1FieldFp16Mask)|(uint32(value)<<RegisterMvfr1FieldFp16Shift))
}

const (
	RegisterMvfr1FieldFphpShift = 24
	RegisterMvfr1FieldFphpMask  = 0xf000000
)

// GetFphp Floating-point half-precision conversion
func (r *RegisterMvfr1Type) GetFphp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr1FieldFphpMask) >> RegisterMvfr1FieldFphpShift)
}

// SetFphp Floating-point half-precision conversion
func (r *RegisterMvfr1Type) SetFphp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr1FieldFphpMask)|(uint32(value)<<RegisterMvfr1FieldFphpShift))
}

const (
	RegisterMvfr1FieldFmacShift = 28
	RegisterMvfr1FieldFmacMask  = 0xf0000000
)

// GetFmac Fused multiply accumulate instructions
func (r *RegisterMvfr1Type) GetFmac() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr1FieldFmacMask) >> RegisterMvfr1FieldFmacShift)
}

// SetFmac Fused multiply accumulate instructions
func (r *RegisterMvfr1Type) SetFmac(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr1FieldFmacMask)|(uint32(value)<<RegisterMvfr1FieldFmacShift))
}

// RegisterMvfr2Type Describes the features provided by the floating-point extension
type RegisterMvfr2Type uint32

func (r *RegisterMvfr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMvfr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMvfr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMvfr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMvfr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMvfr2FieldFpmiscShift = 4
	RegisterMvfr2FieldFpmiscMask  = 0xf0
)

// GetFpmisc Floating-point miscellaneous
func (r *RegisterMvfr2Type) GetFpmisc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMvfr2FieldFpmiscMask) >> RegisterMvfr2FieldFpmiscShift)
}

// SetFpmisc Floating-point miscellaneous
func (r *RegisterMvfr2Type) SetFpmisc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMvfr2FieldFpmiscMask)|(uint32(value)<<RegisterMvfr2FieldFpmiscShift))
}
