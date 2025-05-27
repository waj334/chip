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
	Errfr0     registerErrfr0Type
	_          [4]byte
	Errctrl0   registerErrctrl0Type
	_          [4]byte
	Errstatus0 registerErrstatus0Type
	_          [4]byte
	Erraddr0   registerErraddr0Type
	Erraddr20  registerErraddr20Type
	Errmisc00  registerErrmisc00Type
	Errmisc10  registerErrmisc10Type
	Errmisc20  registerErrmisc20Type
	Errmisc30  registerErrmisc30Type
	Errmisc40  registerErrmisc40Type
	Errmisc50  registerErrmisc50Type
	Errmisc60  registerErrmisc60Type
	Errmisc70  registerErrmisc70Type
	_          [3520]byte
	Errgsr0    registerErrgsr0Type
	_          [452]byte
	Errdevid   registerErrdevidType
	_          [36664]byte
	Rfsr       registerRfsrType
}

// registerErrfr0Type Describes the RAS features that are supported
type registerErrfr0Type uint32

const (
	RegisterErrfr0FieldEdShift = 0
	RegisterErrfr0FieldEdMask  = 0x3
)

// GetEd Error reporting and logging
func (r *registerErrfr0Type) GetEd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterErrfr0FieldEdMask) >> RegisterErrfr0FieldEdShift)
}

// SetEd Error reporting and logging
func (r *registerErrfr0Type) SetEd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrfr0FieldEdMask)|(uint32(value)<<RegisterErrfr0FieldEdShift))
}

const (
	RegisterErrfr0FieldUeShift = 8
	RegisterErrfr0FieldUeMask  = 0x300
)

// GetUe Enable Uncorrected error (UE) reporting as an external abort
func (r *registerErrfr0Type) GetUe() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterErrfr0FieldUeMask) >> RegisterErrfr0FieldUeShift)
}

// SetUe Enable Uncorrected error (UE) reporting as an external abort
func (r *registerErrfr0Type) SetUe(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrfr0FieldUeMask)|(uint32(value)<<RegisterErrfr0FieldUeShift))
}

// registerErrctrl0Type This register is RES0
type registerErrctrl0Type uint32

// registerErrstatus0Type Contains information about the Reliability, Availability, and Serviceability (RAS) event that is currently logged in record 0
type registerErrstatus0Type uint32

const (
	RegisterErrstatus0FieldSerrShift = 0
	RegisterErrstatus0FieldSerrMask  = 0xff
)

// GetSerr Architecturally-defined primary error code
func (r *registerErrstatus0Type) GetSerr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldSerrMask) >> RegisterErrstatus0FieldSerrShift)
}

// SetSerr Architecturally-defined primary error code
func (r *registerErrstatus0Type) SetSerr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldSerrMask)|(uint32(value)<<RegisterErrstatus0FieldSerrShift))
}

const (
	RegisterErrstatus0FieldUetShift = 20
	RegisterErrstatus0FieldUetMask  = 0x300000
)

// GetUet Uncorrectable error type
func (r *registerErrstatus0Type) GetUet() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldUetMask) >> RegisterErrstatus0FieldUetShift)
}

// SetUet Uncorrectable error type
func (r *registerErrstatus0Type) SetUet(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldUetMask)|(uint32(value)<<RegisterErrstatus0FieldUetShift))
}

const (
	RegisterErrstatus0FieldPoisonShift = 22
	RegisterErrstatus0FieldPoisonMask  = 0x400000
)

// GetPoison BusFault due to RPOISON or TEBRn.POISON
func (r *registerErrstatus0Type) GetPoison() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldPoisonMask) != 0
}

// SetPoison BusFault due to RPOISON or TEBRn.POISON
func (r *registerErrstatus0Type) SetPoison(value bool) {
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
func (r *registerErrstatus0Type) GetDe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldDeMask) != 0
}

// SetDe Deferred errors
func (r *registerErrstatus0Type) SetDe(value bool) {
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
func (r *registerErrstatus0Type) GetCe() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldCeMask) >> RegisterErrstatus0FieldCeShift)
}

// SetCe Corrected errors
func (r *registerErrstatus0Type) SetCe(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldCeMask)|(uint32(value)<<RegisterErrstatus0FieldCeShift))
}

const (
	RegisterErrstatus0FieldMvShift = 26
	RegisterErrstatus0FieldMvMask  = 0x4000000
)

// GetMv Miscellaneous registers valid
func (r *registerErrstatus0Type) GetMv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldMvMask) != 0
}

// SetMv Miscellaneous registers valid
func (r *registerErrstatus0Type) SetMv(value bool) {
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
func (r *registerErrstatus0Type) GetOf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldOfMask) != 0
}

// SetOf RAS events occurred since the last time ERRSTATUS0.V was cleared
func (r *registerErrstatus0Type) SetOf(value bool) {
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
func (r *registerErrstatus0Type) GetEr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldErMask) != 0
}

// SetEr BusFault caused by RAS event
func (r *registerErrstatus0Type) SetEr(value bool) {
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
func (r *registerErrstatus0Type) GetUe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldUeMask) != 0
}

// SetUe Uncorrected errors (UEs)
func (r *registerErrstatus0Type) SetUe(value bool) {
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
func (r *registerErrstatus0Type) GetV() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldVMask) != 0
}

// SetV Status valid
func (r *registerErrstatus0Type) SetV(value bool) {
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
func (r *registerErrstatus0Type) GetAv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrstatus0FieldAvMask) != 0
}

// SetAv Address valid
func (r *registerErrstatus0Type) SetAv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErrstatus0FieldAvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErrstatus0FieldAvMask)
	}
}

// registerErraddr0Type Contains information about the address of the Reliability, Availability, and Serviceability (RAS) event in record 0
type registerErraddr0Type uint32

// registerErraddr20Type Contains information about the address of the Reliability, Availability, and Serviceability (RAS) event in record 0
type registerErraddr20Type uint32

const (
	RegisterErraddr20FieldAiShift = 29
	RegisterErraddr20FieldAiMask  = 0x20000000
)

// GetAi Address incorrect
func (r *registerErraddr20Type) GetAi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErraddr20FieldAiMask) != 0
}

// SetAi Address incorrect
func (r *registerErraddr20Type) SetAi(value bool) {
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
func (r *registerErraddr20Type) GetSi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErraddr20FieldSiMask) != 0
}

// SetSi Security information incorrect
func (r *registerErraddr20Type) SetSi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErraddr20FieldSiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErraddr20FieldSiMask)
	}
}

// registerErrmisc00Type This register is RES0
type registerErrmisc00Type uint32

// registerErrmisc10Type Error syndrome register for the event in record 0
type registerErrmisc10Type uint32

const (
	RegisterErrmisc10FieldTypeShift = 0
	RegisterErrmisc10FieldTypeMask  = 0x3
)

// GetType type of Reliability, Availability, and Serviceability (RAS) event logged
func (r *registerErrmisc10Type) GetType() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterErrmisc10FieldTypeMask) >> RegisterErrmisc10FieldTypeShift)
}

// SetType type of Reliability, Availability, and Serviceability (RAS) event logged
func (r *registerErrmisc10Type) SetType(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrmisc10FieldTypeMask)|(uint32(value)<<RegisterErrmisc10FieldTypeShift))
}

// registerErrmisc20Type This register is RES0
type registerErrmisc20Type uint32

// registerErrmisc30Type This register is RES0
type registerErrmisc30Type uint32

// registerErrmisc40Type This register is RES0
type registerErrmisc40Type uint32

// registerErrmisc50Type This register is RES0
type registerErrmisc50Type uint32

// registerErrmisc60Type This register is RES0
type registerErrmisc60Type uint32

// registerErrmisc70Type This register is RES0
type registerErrmisc70Type uint32

// registerErrgsr0Type Summarizes the valid error records
type registerErrgsr0Type uint32

const (
	RegisterErrgsr0FieldErr0Shift = 0
	RegisterErrgsr0FieldErr0Mask  = 0x1
)

// GetErr0 Error record 0 is valid
func (r *registerErrgsr0Type) GetErr0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterErrgsr0FieldErr0Mask) != 0
}

// SetErr0 Error record 0 is valid
func (r *registerErrgsr0Type) SetErr0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterErrgsr0FieldErr0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterErrgsr0FieldErr0Mask)
	}
}

// registerErrdevidType Contains the number of error records that an implementation supports
type registerErrdevidType uint32

const (
	RegisterErrdevidFieldNumShift = 0
	RegisterErrdevidFieldNumMask  = 0xffff
)

// GetNum Maximum Error Record Index+1
func (r *registerErrdevidType) GetNum() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterErrdevidFieldNumMask) >> RegisterErrdevidFieldNumShift)
}

// SetNum Maximum Error Record Index+1
func (r *registerErrdevidType) SetNum(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterErrdevidFieldNumMask)|(uint32(value)<<RegisterErrdevidFieldNumShift))
}

// registerRfsrType Reports the fault status of Reliability, Availability, and Serviceability (RAS) related faults from Error Correcting Code (ECC) errors that are detected in the Level 1 (L1) instruction cache, data cache, and TCM
type registerRfsrType uint32

const (
	RegisterRfsrFieldUetShift = 0
	RegisterRfsrFieldUetMask  = 0x3
)

// GetUet Error type
func (r *registerRfsrType) GetUet() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRfsrFieldUetMask) >> RegisterRfsrFieldUetShift)
}

// SetUet Error type
func (r *registerRfsrType) SetUet(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRfsrFieldUetMask)|(uint32(value)<<RegisterRfsrFieldUetShift))
}

const (
	RegisterRfsrFieldIsShift = 16
	RegisterRfsrFieldIsMask  = 0x7fff0000
)

// GetIs Indicates the type of RAS exception that has occurred
func (r *registerRfsrType) GetIs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRfsrFieldIsMask) >> RegisterRfsrFieldIsShift)
}

// SetIs Indicates the type of RAS exception that has occurred
func (r *registerRfsrType) SetIs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRfsrFieldIsMask)|(uint32(value)<<RegisterRfsrFieldIsShift))
}

const (
	RegisterRfsrFieldValidShift = 31
	RegisterRfsrFieldValidMask  = 0x80000000
)

// GetValid Indicates whether the register is valid
func (r *registerRfsrType) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRfsrFieldValidMask) != 0
}

// SetValid Indicates whether the register is valid
func (r *registerRfsrType) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRfsrFieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRfsrFieldValidMask)
	}
}
