//go:build arm && cortexm

package ras

import (
	"unsafe"
	"volatile"
)

var (
	Ras = (*_ras)(unsafe.Pointer(uintptr(0xe0005000)))
)

type _ras struct {
	Errfr0     RegisterErrfr0Type
	_          [4]byte
	Errctrl0   RegisterErrctrl0Type
	_          [4]byte
	Errstatus0 RegisterErrstatus0Type
	_          [4]byte
	Erraddr0   RegisterErraddr0Type
	Erraddr20  RegisterErraddr20Type
	Errmisc00  RegisterErrmisc00Type
	Errmisc10  RegisterErrmisc10Type
	Errmisc20  RegisterErrmisc20Type
	Errmisc30  RegisterErrmisc30Type
	Errmisc40  RegisterErrmisc40Type
	Errmisc50  RegisterErrmisc50Type
	Errmisc60  RegisterErrmisc60Type
	Errmisc70  RegisterErrmisc70Type
	_          [3520]byte
	Errgsr0    RegisterErrgsr0Type
	_          [452]byte
	Errdevid   RegisterErrdevidType
	_          [36664]byte
	Rfsr       RegisterRfsrType
}

// RegisterErrfr0Type Describes the RAS features that are supported
type RegisterErrfr0Type uint32

func (r *RegisterErrfr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrfr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrfr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrfr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrfr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterErrfr0FieldEdShift = 0
	RegisterErrfr0FieldEdMask  = 0x3
)

// GetEd Error reporting and logging
func (r *RegisterErrfr0Type) GetEd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterErrfr0FieldEdMask) >> RegisterErrfr0FieldEdShift)
}

// SetEd Error reporting and logging
func (r *RegisterErrfr0Type) SetEd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrfr0FieldEdMask)|(uint32(value)<<RegisterErrfr0FieldEdShift))
}

const (
	RegisterErrfr0FieldUeShift = 8
	RegisterErrfr0FieldUeMask  = 0x300
)

// GetUe Enable Uncorrected error (UE) reporting as an external abort
func (r *RegisterErrfr0Type) GetUe() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterErrfr0FieldUeMask) >> RegisterErrfr0FieldUeShift)
}

// SetUe Enable Uncorrected error (UE) reporting as an external abort
func (r *RegisterErrfr0Type) SetUe(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrfr0FieldUeMask)|(uint32(value)<<RegisterErrfr0FieldUeShift))
}

// RegisterErrctrl0Type This register is RES0
type RegisterErrctrl0Type uint32

func (r *RegisterErrctrl0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrctrl0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrctrl0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrctrl0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrctrl0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterErrstatus0Type Contains information about the Reliability, Availability, and Serviceability (RAS) event that is currently logged in record 0
type RegisterErrstatus0Type uint32

func (r *RegisterErrstatus0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrstatus0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrstatus0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrstatus0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrstatus0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterErrstatus0FieldSerrShift = 0
	RegisterErrstatus0FieldSerrMask  = 0xff
)

// GetSerr Architecturally-defined primary error code
func (r *RegisterErrstatus0Type) GetSerr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldSerrMask) >> RegisterErrstatus0FieldSerrShift)
}

// SetSerr Architecturally-defined primary error code
func (r *RegisterErrstatus0Type) SetSerr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldSerrMask)|(uint32(value)<<RegisterErrstatus0FieldSerrShift))
}

const (
	RegisterErrstatus0FieldUetShift = 20
	RegisterErrstatus0FieldUetMask  = 0x300000
)

// GetUet Uncorrectable error type
func (r *RegisterErrstatus0Type) GetUet() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldUetMask) >> RegisterErrstatus0FieldUetShift)
}

// SetUet Uncorrectable error type
func (r *RegisterErrstatus0Type) SetUet(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldUetMask)|(uint32(value)<<RegisterErrstatus0FieldUetShift))
}

const (
	RegisterErrstatus0FieldPoisonShift = 22
	RegisterErrstatus0FieldPoisonMask  = 0x400000
)

// GetPoison BusFault due to RPOISON or TEBRn.POISON
func (r *RegisterErrstatus0Type) GetPoison() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldPoisonMask) != 0
}

// SetPoison BusFault due to RPOISON or TEBRn.POISON
func (r *RegisterErrstatus0Type) SetPoison(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErrstatus0FieldPoisonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldPoisonMask)
	}
}

const (
	RegisterErrstatus0FieldDeShift = 23
	RegisterErrstatus0FieldDeMask  = 0x800000
)

// GetDe Deferred errors
func (r *RegisterErrstatus0Type) GetDe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldDeMask) != 0
}

// SetDe Deferred errors
func (r *RegisterErrstatus0Type) SetDe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErrstatus0FieldDeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldDeMask)
	}
}

const (
	RegisterErrstatus0FieldCeShift = 24
	RegisterErrstatus0FieldCeMask  = 0x3000000
)

// GetCe Corrected errors
func (r *RegisterErrstatus0Type) GetCe() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldCeMask) >> RegisterErrstatus0FieldCeShift)
}

// SetCe Corrected errors
func (r *RegisterErrstatus0Type) SetCe(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldCeMask)|(uint32(value)<<RegisterErrstatus0FieldCeShift))
}

const (
	RegisterErrstatus0FieldMvShift = 26
	RegisterErrstatus0FieldMvMask  = 0x4000000
)

// GetMv Miscellaneous registers valid
func (r *RegisterErrstatus0Type) GetMv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldMvMask) != 0
}

// SetMv Miscellaneous registers valid
func (r *RegisterErrstatus0Type) SetMv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErrstatus0FieldMvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldMvMask)
	}
}

const (
	RegisterErrstatus0FieldOfShift = 27
	RegisterErrstatus0FieldOfMask  = 0x8000000
)

// GetOf RAS events occurred since the last time ERRSTATUS0.V was cleared
func (r *RegisterErrstatus0Type) GetOf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldOfMask) != 0
}

// SetOf RAS events occurred since the last time ERRSTATUS0.V was cleared
func (r *RegisterErrstatus0Type) SetOf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErrstatus0FieldOfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldOfMask)
	}
}

const (
	RegisterErrstatus0FieldErShift = 28
	RegisterErrstatus0FieldErMask  = 0x10000000
)

// GetEr BusFault caused by RAS event
func (r *RegisterErrstatus0Type) GetEr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldErMask) != 0
}

// SetEr BusFault caused by RAS event
func (r *RegisterErrstatus0Type) SetEr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErrstatus0FieldErMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldErMask)
	}
}

const (
	RegisterErrstatus0FieldUeShift = 29
	RegisterErrstatus0FieldUeMask  = 0x20000000
)

// GetUe Uncorrected errors (UEs)
func (r *RegisterErrstatus0Type) GetUe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldUeMask) != 0
}

// SetUe Uncorrected errors (UEs)
func (r *RegisterErrstatus0Type) SetUe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErrstatus0FieldUeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldUeMask)
	}
}

const (
	RegisterErrstatus0FieldVShift = 30
	RegisterErrstatus0FieldVMask  = 0x40000000
)

// GetV Status valid
func (r *RegisterErrstatus0Type) GetV() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldVMask) != 0
}

// SetV Status valid
func (r *RegisterErrstatus0Type) SetV(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErrstatus0FieldVMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldVMask)
	}
}

const (
	RegisterErrstatus0FieldAvShift = 31
	RegisterErrstatus0FieldAvMask  = 0x80000000
)

// GetAv Address valid
func (r *RegisterErrstatus0Type) GetAv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldAvMask) != 0
}

// SetAv Address valid
func (r *RegisterErrstatus0Type) SetAv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErrstatus0FieldAvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldAvMask)
	}
}

// RegisterErraddr0Type Contains information about the address of the Reliability, Availability, and Serviceability (RAS) event in record 0
type RegisterErraddr0Type uint32

func (r *RegisterErraddr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErraddr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErraddr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErraddr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErraddr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterErraddr20Type Contains information about the address of the Reliability, Availability, and Serviceability (RAS) event in record 0
type RegisterErraddr20Type uint32

func (r *RegisterErraddr20Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErraddr20Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErraddr20Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErraddr20Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErraddr20Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterErraddr20FieldAiShift = 29
	RegisterErraddr20FieldAiMask  = 0x20000000
)

// GetAi Address incorrect
func (r *RegisterErraddr20Type) GetAi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErraddr20FieldAiMask) != 0
}

// SetAi Address incorrect
func (r *RegisterErraddr20Type) SetAi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErraddr20FieldAiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErraddr20FieldAiMask)
	}
}

const (
	RegisterErraddr20FieldSiShift = 30
	RegisterErraddr20FieldSiMask  = 0x40000000
)

// GetSi Security information incorrect
func (r *RegisterErraddr20Type) GetSi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErraddr20FieldSiMask) != 0
}

// SetSi Security information incorrect
func (r *RegisterErraddr20Type) SetSi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErraddr20FieldSiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErraddr20FieldSiMask)
	}
}

// RegisterErrmisc00Type This register is RES0
type RegisterErrmisc00Type uint32

func (r *RegisterErrmisc00Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrmisc00Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrmisc00Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrmisc00Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrmisc00Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterErrmisc10Type Error syndrome register for the event in record 0
type RegisterErrmisc10Type uint32

func (r *RegisterErrmisc10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrmisc10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrmisc10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrmisc10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrmisc10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterErrmisc10FieldTypeShift = 0
	RegisterErrmisc10FieldTypeMask  = 0x3
)

// GetType type of Reliability, Availability, and Serviceability (RAS) event logged
func (r *RegisterErrmisc10Type) GetType() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterErrmisc10FieldTypeMask) >> RegisterErrmisc10FieldTypeShift)
}

// SetType type of Reliability, Availability, and Serviceability (RAS) event logged
func (r *RegisterErrmisc10Type) SetType(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrmisc10FieldTypeMask)|(uint32(value)<<RegisterErrmisc10FieldTypeShift))
}

// RegisterErrmisc20Type This register is RES0
type RegisterErrmisc20Type uint32

func (r *RegisterErrmisc20Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrmisc20Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrmisc20Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrmisc20Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrmisc20Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterErrmisc30Type This register is RES0
type RegisterErrmisc30Type uint32

func (r *RegisterErrmisc30Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrmisc30Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrmisc30Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrmisc30Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrmisc30Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterErrmisc40Type This register is RES0
type RegisterErrmisc40Type uint32

func (r *RegisterErrmisc40Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrmisc40Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrmisc40Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrmisc40Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrmisc40Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterErrmisc50Type This register is RES0
type RegisterErrmisc50Type uint32

func (r *RegisterErrmisc50Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrmisc50Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrmisc50Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrmisc50Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrmisc50Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterErrmisc60Type This register is RES0
type RegisterErrmisc60Type uint32

func (r *RegisterErrmisc60Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrmisc60Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrmisc60Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrmisc60Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrmisc60Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterErrmisc70Type This register is RES0
type RegisterErrmisc70Type uint32

func (r *RegisterErrmisc70Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrmisc70Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrmisc70Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrmisc70Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrmisc70Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterErrgsr0Type Summarizes the valid error records
type RegisterErrgsr0Type uint32

func (r *RegisterErrgsr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrgsr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrgsr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrgsr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrgsr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterErrgsr0FieldErr0Shift = 0
	RegisterErrgsr0FieldErr0Mask  = 0x1
)

// GetErr0 Error record 0 is valid
func (r *RegisterErrgsr0Type) GetErr0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrgsr0FieldErr0Mask) != 0
}

// SetErr0 Error record 0 is valid
func (r *RegisterErrgsr0Type) SetErr0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErrgsr0FieldErr0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErrgsr0FieldErr0Mask)
	}
}

// RegisterErrdevidType Contains the number of error records that an implementation supports
type RegisterErrdevidType uint32

func (r *RegisterErrdevidType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterErrdevidType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterErrdevidType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterErrdevidType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterErrdevidType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterErrdevidFieldNumShift = 0
	RegisterErrdevidFieldNumMask  = 0xffff
)

// GetNum Maximum Error Record Index+1
func (r *RegisterErrdevidType) GetNum() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterErrdevidFieldNumMask) >> RegisterErrdevidFieldNumShift)
}

// SetNum Maximum Error Record Index+1
func (r *RegisterErrdevidType) SetNum(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrdevidFieldNumMask)|(uint32(value)<<RegisterErrdevidFieldNumShift))
}

// RegisterRfsrType Reports the fault status of Reliability, Availability, and Serviceability (RAS) related faults from Error Correcting Code (ECC) errors that are detected in the Level 1 (L1) instruction cache, data cache, and TCM
type RegisterRfsrType uint32

func (r *RegisterRfsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRfsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRfsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRfsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRfsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRfsrFieldUetShift = 0
	RegisterRfsrFieldUetMask  = 0x3
)

// GetUet Error type
func (r *RegisterRfsrType) GetUet() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRfsrFieldUetMask) >> RegisterRfsrFieldUetShift)
}

// SetUet Error type
func (r *RegisterRfsrType) SetUet(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRfsrFieldUetMask)|(uint32(value)<<RegisterRfsrFieldUetShift))
}

const (
	RegisterRfsrFieldIsShift = 16
	RegisterRfsrFieldIsMask  = 0x7fff0000
)

// GetIs Indicates the type of RAS exception that has occurred
func (r *RegisterRfsrType) GetIs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRfsrFieldIsMask) >> RegisterRfsrFieldIsShift)
}

// SetIs Indicates the type of RAS exception that has occurred
func (r *RegisterRfsrType) SetIs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRfsrFieldIsMask)|(uint32(value)<<RegisterRfsrFieldIsShift))
}

const (
	RegisterRfsrFieldValidShift = 31
	RegisterRfsrFieldValidMask  = 0x80000000
)

// GetValid Indicates whether the register is valid
func (r *RegisterRfsrType) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRfsrFieldValidMask) != 0
}

// SetValid Indicates whether the register is valid
func (r *RegisterRfsrType) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRfsrFieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRfsrFieldValidMask)
	}
}
