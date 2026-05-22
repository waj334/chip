//go:build arm && cortexm

package scb

import (
	"unsafe"
	"volatile"
)

var (
	Scb = (*_scb)(unsafe.Pointer(uintptr(0xe000ecfc)))
)

type _scb struct {
	Revidr  RegisterRevidrType
	Cpuid   RegisterCpuidType
	Icsr    RegisterIcsrType
	Vtor    RegisterVtorType
	Aircr   RegisterAircrType
	Scr     RegisterScrType
	Ccr     RegisterCcrType
	Shpr1   RegisterShpr1Type
	Shpr2   RegisterShpr2Type
	Shpr3   RegisterShpr3Type
	Shcsr   RegisterShcsrType
	Cfsr    RegisterCfsrType
	Hfsr    RegisterHfsrType
	_       [4]byte
	Mmfar   RegisterMmfarType
	Bfar    RegisterBfarType
	Afsr    RegisterAfsrType
	Idpfr0  RegisterIdpfr0Type
	Idpfr1  RegisterIdpfr1Type
	Iddfr0  RegisterIddfr0Type
	Idafr0  RegisterIdafr0Type
	Idmmfr0 RegisterIdmmfr0Type
	Idmmfr1 RegisterIdmmfr1Type
	Idmmfr2 RegisterIdmmfr2Type
	Idmmfr3 RegisterIdmmfr3Type
	Idisar0 RegisterIdisar0Type
	Idisar1 RegisterIdisar1Type
	Idisar2 RegisterIdisar2Type
	Idisar3 RegisterIdisar3Type
	Idisar4 RegisterIdisar4Type
	Clidr   RegisterClidrType
	_       [4]byte
	Ctr     RegisterCtrType
	Ccsidr  RegisterCcsidrType
	Csselr  RegisterCsselrType
	Cpacr   RegisterCpacrType
	Nsacr   RegisterNsacrType
}

// RegisterRevidrType Provides additional implementation-specific minor revision that can be interpreted with the CPUID register
type RegisterRevidrType uint32

func (r *RegisterRevidrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRevidrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRevidrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRevidrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRevidrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRevidrFieldImplementationspecificShift = 0
	RegisterRevidrFieldImplementationspecificMask  = 0xffffffff
)

// GetImplementationspecific Implementation-specific minor revision information that can be interpreted with the CPUID register.
func (r *RegisterRevidrType) GetImplementationspecific() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRevidrFieldImplementationspecificMask) >> RegisterRevidrFieldImplementationspecificShift)
}

// SetImplementationspecific Implementation-specific minor revision information that can be interpreted with the CPUID register.
func (r *RegisterRevidrType) SetImplementationspecific(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRevidrFieldImplementationspecificMask)|(uint32(value)<<RegisterRevidrFieldImplementationspecificShift))
}

// RegisterCpuidType Contains the processor part number, version, and implementation information
type RegisterCpuidType uint32

func (r *RegisterCpuidType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCpuidType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCpuidType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCpuidType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCpuidType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCpuidFieldRevisionShift = 0
	RegisterCpuidFieldRevisionMask  = 0xf
)

// GetRevision Revision number, the m value in the rnpm product revision identifier
func (r *RegisterCpuidType) GetRevision() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldRevisionMask) >> RegisterCpuidFieldRevisionShift)
}

// SetRevision Revision number, the m value in the rnpm product revision identifier
func (r *RegisterCpuidType) SetRevision(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpuidFieldRevisionMask)|(uint32(value)<<RegisterCpuidFieldRevisionShift))
}

const (
	RegisterCpuidFieldPartnoShift = 4
	RegisterCpuidFieldPartnoMask  = 0xfff0
)

// GetPartno Part number of the processor
func (r *RegisterCpuidType) GetPartno() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldPartnoMask) >> RegisterCpuidFieldPartnoShift)
}

// SetPartno Part number of the processor
func (r *RegisterCpuidType) SetPartno(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpuidFieldPartnoMask)|(uint32(value)<<RegisterCpuidFieldPartnoShift))
}

const (
	RegisterCpuidFieldArchitectureShift = 16
	RegisterCpuidFieldArchitectureMask  = 0xf0000
)

// GetArchitecture Reads as 0xF
func (r *RegisterCpuidType) GetArchitecture() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldArchitectureMask) >> RegisterCpuidFieldArchitectureShift)
}

// SetArchitecture Reads as 0xF
func (r *RegisterCpuidType) SetArchitecture(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpuidFieldArchitectureMask)|(uint32(value)<<RegisterCpuidFieldArchitectureShift))
}

const (
	RegisterCpuidFieldVariantShift = 20
	RegisterCpuidFieldVariantMask  = 0xf00000
)

// GetVariant Variant number, the n value in the rnpm product revision identifier
func (r *RegisterCpuidType) GetVariant() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldVariantMask) >> RegisterCpuidFieldVariantShift)
}

// SetVariant Variant number, the n value in the rnpm product revision identifier
func (r *RegisterCpuidType) SetVariant(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpuidFieldVariantMask)|(uint32(value)<<RegisterCpuidFieldVariantShift))
}

const (
	RegisterCpuidFieldImplementerShift = 24
	RegisterCpuidFieldImplementerMask  = 0xff000000
)

// GetImplementer Implementer code
func (r *RegisterCpuidType) GetImplementer() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldImplementerMask) >> RegisterCpuidFieldImplementerShift)
}

// SetImplementer Implementer code
func (r *RegisterCpuidType) SetImplementer(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpuidFieldImplementerMask)|(uint32(value)<<RegisterCpuidFieldImplementerShift))
}

// RegisterIcsrType Provides a set-pending bit for the non-maskable interrupt exception, and set-pending and clear-pending bits for the PendSV and SysTick exceptions
type RegisterIcsrType uint32

func (r *RegisterIcsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIcsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIcsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIcsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIcsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIcsrFieldVectactiveShift = 0
	RegisterIcsrFieldVectactiveMask  = 0x1ff
)

// GetVectactive Contains the active exception number
func (r *RegisterIcsrType) GetVectactive() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldVectactiveMask) >> RegisterIcsrFieldVectactiveShift)
}

// SetVectactive Contains the active exception number
func (r *RegisterIcsrType) SetVectactive(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldVectactiveMask)|(uint32(value)<<RegisterIcsrFieldVectactiveShift))
}

const (
	RegisterIcsrFieldRettobaseShift = 11
	RegisterIcsrFieldRettobaseMask  = 0x800
)

// GetRettobase Indicates whether there are pre-empted active exceptions
func (r *RegisterIcsrType) GetRettobase() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldRettobaseMask) != 0
}

// SetRettobase Indicates whether there are pre-empted active exceptions
func (r *RegisterIcsrType) SetRettobase(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldRettobaseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldRettobaseMask)
	}
}

const (
	RegisterIcsrFieldVectpendingShift = 12
	RegisterIcsrFieldVectpendingMask  = 0x1ff000
)

// GetVectpending Indicates the exception number of the highest priority pending enabled exception
func (r *RegisterIcsrType) GetVectpending() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldVectpendingMask) >> RegisterIcsrFieldVectpendingShift)
}

// SetVectpending Indicates the exception number of the highest priority pending enabled exception
func (r *RegisterIcsrType) SetVectpending(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldVectpendingMask)|(uint32(value)<<RegisterIcsrFieldVectpendingShift))
}

const (
	RegisterIcsrFieldIsrpendingShift = 22
	RegisterIcsrFieldIsrpendingMask  = 0x400000
)

// GetIsrpending Interrupt pending flag, excluding NMI and Faults
func (r *RegisterIcsrType) GetIsrpending() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldIsrpendingMask) != 0
}

// SetIsrpending Interrupt pending flag, excluding NMI and Faults
func (r *RegisterIcsrType) SetIsrpending(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldIsrpendingMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldIsrpendingMask)
	}
}

const (
	RegisterIcsrFieldIsrpreemptShift = 23
	RegisterIcsrFieldIsrpreemptMask  = 0x800000
)

// GetIsrpreempt Indicates whether a pending exception is handled on exit from debug state.
func (r *RegisterIcsrType) GetIsrpreempt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldIsrpreemptMask) != 0
}

// SetIsrpreempt Indicates whether a pending exception is handled on exit from debug state.
func (r *RegisterIcsrType) SetIsrpreempt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldIsrpreemptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldIsrpreemptMask)
	}
}

const (
	RegisterIcsrFieldPendstclrShift = 25
	RegisterIcsrFieldPendstclrMask  = 0x2000000
)

// GetPendstclr SysTick exception clear-pending bit
func (r *RegisterIcsrType) GetPendstclr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendstclrMask) != 0
}

// SetPendstclr SysTick exception clear-pending bit
func (r *RegisterIcsrType) SetPendstclr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldPendstclrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldPendstclrMask)
	}
}

const (
	RegisterIcsrFieldPendstsetShift = 26
	RegisterIcsrFieldPendstsetMask  = 0x4000000
)

// GetPendstset SysTick exception set-pending bit
func (r *RegisterIcsrType) GetPendstset() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendstsetMask) != 0
}

// SetPendstset SysTick exception set-pending bit
func (r *RegisterIcsrType) SetPendstset(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldPendstsetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldPendstsetMask)
	}
}

const (
	RegisterIcsrFieldPendsvclrShift = 27
	RegisterIcsrFieldPendsvclrMask  = 0x8000000
)

// GetPendsvclr PendSV clear-pending bit (write-only)
func (r *RegisterIcsrType) GetPendsvclr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendsvclrMask) != 0
}

// SetPendsvclr PendSV clear-pending bit (write-only)
func (r *RegisterIcsrType) SetPendsvclr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldPendsvclrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldPendsvclrMask)
	}
}

const (
	RegisterIcsrFieldPendsvsetShift = 28
	RegisterIcsrFieldPendsvsetMask  = 0x10000000
)

// GetPendsvset PendSV set-pending bit
func (r *RegisterIcsrType) GetPendsvset() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendsvsetMask) != 0
}

// SetPendsvset PendSV set-pending bit
func (r *RegisterIcsrType) SetPendsvset(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldPendsvsetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldPendsvsetMask)
	}
}

const (
	RegisterIcsrFieldPendnmiclrShift = 30
	RegisterIcsrFieldPendnmiclrMask  = 0x40000000
)

// GetPendnmiclr Pend NMI clear bit (write-only)
func (r *RegisterIcsrType) GetPendnmiclr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendnmiclrMask) != 0
}

// SetPendnmiclr Pend NMI clear bit (write-only)
func (r *RegisterIcsrType) SetPendnmiclr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldPendnmiclrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldPendnmiclrMask)
	}
}

const (
	RegisterIcsrFieldPendnmisetShift = 31
	RegisterIcsrFieldPendnmisetMask  = 0x80000000
)

// GetPendnmiset On writes, makes the NMI exception pending
func (r *RegisterIcsrType) GetPendnmiset() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendnmisetMask) != 0
}

// SetPendnmiset On writes, makes the NMI exception pending
func (r *RegisterIcsrType) SetPendnmiset(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldPendnmisetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldPendnmisetMask)
	}
}

// RegisterVtorType Indicates the offset of the vector table base address from memory address 0x00000000
type RegisterVtorType uint32

func (r *RegisterVtorType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterVtorType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterVtorType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterVtorType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterVtorType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterVtorFieldTbloffShift = 7
	RegisterVtorFieldTbloffMask  = 0xffffff80
)

// GetTbloff Bits[31:7] of the vector table address
func (r *RegisterVtorType) GetTbloff() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterVtorFieldTbloffMask) >> RegisterVtorFieldTbloffShift)
}

// SetTbloff Bits[31:7] of the vector table address
func (r *RegisterVtorType) SetTbloff(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterVtorFieldTbloffMask)|(uint32(value)<<RegisterVtorFieldTbloffShift))
}

// RegisterAircrType Sets or returns interrupt control data
type RegisterAircrType uint32

func (r *RegisterAircrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAircrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAircrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAircrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAircrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAircrFieldVectclractiveShift = 1
	RegisterAircrFieldVectclractiveMask  = 0x2
)

// GetVectclractive Reserved for Debug use. This bit reads as 0. When writing to the register you must write 0 to this bit, otherwise behavior is UNPREDICTABLE
func (r *RegisterAircrType) GetVectclractive() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldVectclractiveMask) != 0
}

// SetVectclractive Reserved for Debug use. This bit reads as 0. When writing to the register you must write 0 to this bit, otherwise behavior is UNPREDICTABLE
func (r *RegisterAircrType) SetVectclractive(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAircrFieldVectclractiveMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldVectclractiveMask)
	}
}

const (
	RegisterAircrFieldSysresetreqShift = 2
	RegisterAircrFieldSysresetreqMask  = 0x4
)

// GetSysresetreq System Reset Request
func (r *RegisterAircrType) GetSysresetreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldSysresetreqMask) != 0
}

// SetSysresetreq System Reset Request
func (r *RegisterAircrType) SetSysresetreq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAircrFieldSysresetreqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldSysresetreqMask)
	}
}

const (
	RegisterAircrFieldSysresetreqsShift = 3
	RegisterAircrFieldSysresetreqsMask  = 0x8
)

// GetSysresetreqs System reset request, Secure state only
func (r *RegisterAircrType) GetSysresetreqs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldSysresetreqsMask) != 0
}

// SetSysresetreqs System reset request, Secure state only
func (r *RegisterAircrType) SetSysresetreqs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAircrFieldSysresetreqsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldSysresetreqsMask)
	}
}

const (
	RegisterAircrFieldPrigroupShift = 8
	RegisterAircrFieldPrigroupMask  = 0x700
)

// GetPrigroup Interrupt priority grouping
func (r *RegisterAircrType) GetPrigroup() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldPrigroupMask) >> RegisterAircrFieldPrigroupShift)
}

// SetPrigroup Interrupt priority grouping
func (r *RegisterAircrType) SetPrigroup(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldPrigroupMask)|(uint32(value)<<RegisterAircrFieldPrigroupShift))
}

const (
	RegisterAircrFieldBfhfnminsShift = 13
	RegisterAircrFieldBfhfnminsMask  = 0x2000
)

// GetBfhfnmins BusFault, HardFault, and NMI Non-secure enable
func (r *RegisterAircrType) GetBfhfnmins() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldBfhfnminsMask) != 0
}

// SetBfhfnmins BusFault, HardFault, and NMI Non-secure enable
func (r *RegisterAircrType) SetBfhfnmins(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAircrFieldBfhfnminsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldBfhfnminsMask)
	}
}

const (
	RegisterAircrFieldPrisShift = 14
	RegisterAircrFieldPrisMask  = 0x4000
)

// GetPris Prioritize Secure exceptions
func (r *RegisterAircrType) GetPris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldPrisMask) != 0
}

// SetPris Prioritize Secure exceptions
func (r *RegisterAircrType) SetPris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAircrFieldPrisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldPrisMask)
	}
}

const (
	RegisterAircrFieldEndiannessShift = 15
	RegisterAircrFieldEndiannessMask  = 0x8000
)

// GetEndianness Data endianness
func (r *RegisterAircrType) GetEndianness() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldEndiannessMask) != 0
}

// SetEndianness Data endianness
func (r *RegisterAircrType) SetEndianness(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAircrFieldEndiannessMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldEndiannessMask)
	}
}

const (
	RegisterAircrFieldVectkeyShift = 16
	RegisterAircrFieldVectkeyMask  = 0xffff0000
)

// GetVectkey Register key. Reads as 0x0FA05, On writes, write 0x5FA to VECTKEY, otherwise the write is ignored.
func (r *RegisterAircrType) GetVectkey() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldVectkeyMask) >> RegisterAircrFieldVectkeyShift)
}

// SetVectkey Register key. Reads as 0x0FA05, On writes, write 0x5FA to VECTKEY, otherwise the write is ignored.
func (r *RegisterAircrType) SetVectkey(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldVectkeyMask)|(uint32(value)<<RegisterAircrFieldVectkeyShift))
}

// RegisterScrType Controls features of entry to and exit from low-power state
type RegisterScrType uint32

func (r *RegisterScrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterScrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterScrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterScrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterScrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterScrFieldSleeponexitShift = 1
	RegisterScrFieldSleeponexitMask  = 0x2
)

// GetSleeponexit Indicates sleep-on-exit when returning from Handler mode to Thread mode
func (r *RegisterScrType) GetSleeponexit() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScrFieldSleeponexitMask) != 0
}

// SetSleeponexit Indicates sleep-on-exit when returning from Handler mode to Thread mode
func (r *RegisterScrType) SetSleeponexit(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScrFieldSleeponexitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScrFieldSleeponexitMask)
	}
}

const (
	RegisterScrFieldSleepdeepShift = 2
	RegisterScrFieldSleepdeepMask  = 0x4
)

// GetSleepdeep Indicates whether the processor uses sleep or deep sleep as its low-power mode
func (r *RegisterScrType) GetSleepdeep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScrFieldSleepdeepMask) != 0
}

// SetSleepdeep Indicates whether the processor uses sleep or deep sleep as its low-power mode
func (r *RegisterScrType) SetSleepdeep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScrFieldSleepdeepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScrFieldSleepdeepMask)
	}
}

const (
	RegisterScrFieldSleepdeepsShift = 3
	RegisterScrFieldSleepdeepsMask  = 0x8
)

// GetSleepdeeps Controls whether the SLEEPDEEP bit is only accessible from the Secure state
func (r *RegisterScrType) GetSleepdeeps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScrFieldSleepdeepsMask) != 0
}

// SetSleepdeeps Controls whether the SLEEPDEEP bit is only accessible from the Secure state
func (r *RegisterScrType) SetSleepdeeps(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScrFieldSleepdeepsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScrFieldSleepdeepsMask)
	}
}

const (
	RegisterScrFieldSevonpendShift = 4
	RegisterScrFieldSevonpendMask  = 0x10
)

// GetSevonpend Determines whether an interrupt assigned to the same Security state as the SEVONPEND bit transitioning from inactive state to pending state generates a wakeup event
func (r *RegisterScrType) GetSevonpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScrFieldSevonpendMask) != 0
}

// SetSevonpend Determines whether an interrupt assigned to the same Security state as the SEVONPEND bit transitioning from inactive state to pending state generates a wakeup event
func (r *RegisterScrType) SetSevonpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScrFieldSevonpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScrFieldSevonpendMask)
	}
}

// RegisterCcrType indicates some aspects of the behavior of the processor
type RegisterCcrType uint32

func (r *RegisterCcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcrFieldUsersetmpendShift = 1
	RegisterCcrFieldUsersetmpendMask  = 0x2
)

// GetUsersetmpend User set main pending. Determines whether unprivileged accesses are permitted to pend interrupts from the STIR.
func (r *RegisterCcrType) GetUsersetmpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldUsersetmpendMask) != 0
}

// SetUsersetmpend User set main pending. Determines whether unprivileged accesses are permitted to pend interrupts from the STIR.
func (r *RegisterCcrType) SetUsersetmpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldUsersetmpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldUsersetmpendMask)
	}
}

const (
	RegisterCcrFieldUnaligntrpShift = 3
	RegisterCcrFieldUnaligntrpMask  = 0x8
)

// GetUnaligntrp Controls the trapping of unaligned word or halfword accesses
func (r *RegisterCcrType) GetUnaligntrp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldUnaligntrpMask) != 0
}

// SetUnaligntrp Controls the trapping of unaligned word or halfword accesses
func (r *RegisterCcrType) SetUnaligntrp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldUnaligntrpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldUnaligntrpMask)
	}
}

const (
	RegisterCcrFieldDiv0trpShift = 4
	RegisterCcrFieldDiv0trpMask  = 0x10
)

// GetDiv0trp Divide by zero trap. Controls the generation of a DIVBYZERO UsageFault when attempting to perform integer division by zero.
func (r *RegisterCcrType) GetDiv0trp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDiv0trpMask) != 0
}

// SetDiv0trp Divide by zero trap. Controls the generation of a DIVBYZERO UsageFault when attempting to perform integer division by zero.
func (r *RegisterCcrType) SetDiv0trp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldDiv0trpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDiv0trpMask)
	}
}

const (
	RegisterCcrFieldBfhfnmignShift = 8
	RegisterCcrFieldBfhfnmignMask  = 0x100
)

// GetBfhfnmign Determines the effect of precise bus faults on handlers running at a requested priority less than 0
func (r *RegisterCcrType) GetBfhfnmign() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldBfhfnmignMask) != 0
}

// SetBfhfnmign Determines the effect of precise bus faults on handlers running at a requested priority less than 0
func (r *RegisterCcrType) SetBfhfnmign(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldBfhfnmignMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldBfhfnmignMask)
	}
}

const (
	RegisterCcrFieldStkofhfnmignShift = 10
	RegisterCcrFieldStkofhfnmignMask  = 0x400
)

// GetStkofhfnmign Controls the effect of a stack limit violation while executing at a requested priority less than 0
func (r *RegisterCcrType) GetStkofhfnmign() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldStkofhfnmignMask) != 0
}

// SetStkofhfnmign Controls the effect of a stack limit violation while executing at a requested priority less than 0
func (r *RegisterCcrType) SetStkofhfnmign(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldStkofhfnmignMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldStkofhfnmignMask)
	}
}

const (
	RegisterCcrFieldDcShift = 16
	RegisterCcrFieldDcMask  = 0x10000
)

// GetDc Data cache enable
func (r *RegisterCcrType) GetDc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDcMask) != 0
}

// SetDc Data cache enable
func (r *RegisterCcrType) SetDc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldDcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDcMask)
	}
}

const (
	RegisterCcrFieldIcShift = 17
	RegisterCcrFieldIcMask  = 0x20000
)

// GetIc Instruction cache enable
func (r *RegisterCcrType) GetIc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldIcMask) != 0
}

// SetIc Instruction cache enable
func (r *RegisterCcrType) SetIc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldIcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldIcMask)
	}
}

const (
	RegisterCcrFieldLobShift = 19
	RegisterCcrFieldLobMask  = 0x80000
)

// GetLob Loop and branch info cache enable
func (r *RegisterCcrType) GetLob() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldLobMask) != 0
}

// SetLob Loop and branch info cache enable
func (r *RegisterCcrType) SetLob(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldLobMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldLobMask)
	}
}

const (
	RegisterCcrFieldTrdShift = 20
	RegisterCcrFieldTrdMask  = 0x100000
)

// GetTrd Thread reentrancy disabled
func (r *RegisterCcrType) GetTrd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldTrdMask) != 0
}

// SetTrd Thread reentrancy disabled
func (r *RegisterCcrType) SetTrd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldTrdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldTrdMask)
	}
}

// RegisterShpr1Type Sets or returns priority for system handlers 4-7
type RegisterShpr1Type uint32

func (r *RegisterShpr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterShpr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterShpr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterShpr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterShpr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterShpr1FieldPri4Shift = 0
	RegisterShpr1FieldPri4Mask  = 0xff
)

// GetPri4 Priority of system handler 4, MemManage
func (r *RegisterShpr1Type) GetPri4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr1FieldPri4Mask) >> RegisterShpr1FieldPri4Shift)
}

// SetPri4 Priority of system handler 4, MemManage
func (r *RegisterShpr1Type) SetPri4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr1FieldPri4Mask)|(uint32(value)<<RegisterShpr1FieldPri4Shift))
}

const (
	RegisterShpr1FieldPri5Shift = 8
	RegisterShpr1FieldPri5Mask  = 0xff00
)

// GetPri5 Priority of system handler 5, BusFault
func (r *RegisterShpr1Type) GetPri5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr1FieldPri5Mask) >> RegisterShpr1FieldPri5Shift)
}

// SetPri5 Priority of system handler 5, BusFault
func (r *RegisterShpr1Type) SetPri5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr1FieldPri5Mask)|(uint32(value)<<RegisterShpr1FieldPri5Shift))
}

const (
	RegisterShpr1FieldPri6Shift = 16
	RegisterShpr1FieldPri6Mask  = 0xff0000
)

// GetPri6 Priority of system handler 6, UsageFault
func (r *RegisterShpr1Type) GetPri6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr1FieldPri6Mask) >> RegisterShpr1FieldPri6Shift)
}

// SetPri6 Priority of system handler 6, UsageFault
func (r *RegisterShpr1Type) SetPri6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr1FieldPri6Mask)|(uint32(value)<<RegisterShpr1FieldPri6Shift))
}

const (
	RegisterShpr1FieldPri7Shift = 24
	RegisterShpr1FieldPri7Mask  = 0xff000000
)

// GetPri7 Priority of system handler 7, SecureFault
func (r *RegisterShpr1Type) GetPri7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr1FieldPri7Mask) >> RegisterShpr1FieldPri7Shift)
}

// SetPri7 Priority of system handler 7, SecureFault
func (r *RegisterShpr1Type) SetPri7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr1FieldPri7Mask)|(uint32(value)<<RegisterShpr1FieldPri7Shift))
}

// RegisterShpr2Type Sets or returns priority for system handler 8-11
type RegisterShpr2Type uint32

func (r *RegisterShpr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterShpr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterShpr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterShpr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterShpr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterShpr2FieldPri11Shift = 24
	RegisterShpr2FieldPri11Mask  = 0xff000000
)

// GetPri11 Priority of system handler 11, SVCall
func (r *RegisterShpr2Type) GetPri11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr2FieldPri11Mask) >> RegisterShpr2FieldPri11Shift)
}

// SetPri11 Priority of system handler 11, SVCall
func (r *RegisterShpr2Type) SetPri11(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr2FieldPri11Mask)|(uint32(value)<<RegisterShpr2FieldPri11Shift))
}

// RegisterShpr3Type Sets or returns priority for system handlers 12-15
type RegisterShpr3Type uint32

func (r *RegisterShpr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterShpr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterShpr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterShpr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterShpr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterShpr3FieldPri14Shift = 16
	RegisterShpr3FieldPri14Mask  = 0xff0000
)

// GetPri14 Priority of system handler 14, PendSV
func (r *RegisterShpr3Type) GetPri14() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr3FieldPri14Mask) >> RegisterShpr3FieldPri14Shift)
}

// SetPri14 Priority of system handler 14, PendSV
func (r *RegisterShpr3Type) SetPri14(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr3FieldPri14Mask)|(uint32(value)<<RegisterShpr3FieldPri14Shift))
}

const (
	RegisterShpr3FieldPri15Shift = 24
	RegisterShpr3FieldPri15Mask  = 0xff000000
)

// GetPri15 Priority of system handler 15, SysTick
func (r *RegisterShpr3Type) GetPri15() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr3FieldPri15Mask) >> RegisterShpr3FieldPri15Shift)
}

// SetPri15 Priority of system handler 15, SysTick
func (r *RegisterShpr3Type) SetPri15(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr3FieldPri15Mask)|(uint32(value)<<RegisterShpr3FieldPri15Shift))
}

// RegisterShcsrType Enables the system handlers and indicates the pending status of the BusFault, MemManage fault, and SVC exceptions, and indicates the active status of the system handlers
type RegisterShcsrType uint32

func (r *RegisterShcsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterShcsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterShcsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterShcsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterShcsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterShcsrFieldMemfaultactShift = 0
	RegisterShcsrFieldMemfaultactMask  = 0x1
)

// GetMemfaultact MemManage exception active
func (r *RegisterShcsrType) GetMemfaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldMemfaultactMask) != 0
}

// SetMemfaultact MemManage exception active
func (r *RegisterShcsrType) SetMemfaultact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldMemfaultactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldMemfaultactMask)
	}
}

const (
	RegisterShcsrFieldBusfaultactShift = 1
	RegisterShcsrFieldBusfaultactMask  = 0x2
)

// GetBusfaultact BusFault exception active
func (r *RegisterShcsrType) GetBusfaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldBusfaultactMask) != 0
}

// SetBusfaultact BusFault exception active
func (r *RegisterShcsrType) SetBusfaultact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldBusfaultactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldBusfaultactMask)
	}
}

const (
	RegisterShcsrFieldHardfaultactShift = 2
	RegisterShcsrFieldHardfaultactMask  = 0x4
)

// GetHardfaultact HardFault exception active
func (r *RegisterShcsrType) GetHardfaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldHardfaultactMask) != 0
}

// SetHardfaultact HardFault exception active
func (r *RegisterShcsrType) SetHardfaultact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldHardfaultactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldHardfaultactMask)
	}
}

const (
	RegisterShcsrFieldUsgfaultactShift = 3
	RegisterShcsrFieldUsgfaultactMask  = 0x8
)

// GetUsgfaultact UsageFault exception active
func (r *RegisterShcsrType) GetUsgfaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldUsgfaultactMask) != 0
}

// SetUsgfaultact UsageFault exception active
func (r *RegisterShcsrType) SetUsgfaultact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldUsgfaultactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldUsgfaultactMask)
	}
}

const (
	RegisterShcsrFieldSecurefaultactShift = 4
	RegisterShcsrFieldSecurefaultactMask  = 0x10
)

// GetSecurefaultact SecureFault exception active state
func (r *RegisterShcsrType) GetSecurefaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldSecurefaultactMask) != 0
}

// SetSecurefaultact SecureFault exception active state
func (r *RegisterShcsrType) SetSecurefaultact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldSecurefaultactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldSecurefaultactMask)
	}
}

const (
	RegisterShcsrFieldNmiactShift = 5
	RegisterShcsrFieldNmiactMask  = 0x20
)

// GetNmiact NMI exception active
func (r *RegisterShcsrType) GetNmiact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldNmiactMask) != 0
}

// SetNmiact NMI exception active
func (r *RegisterShcsrType) SetNmiact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldNmiactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldNmiactMask)
	}
}

const (
	RegisterShcsrFieldSvcallactShift = 7
	RegisterShcsrFieldSvcallactMask  = 0x80
)

// GetSvcallact SVCall active
func (r *RegisterShcsrType) GetSvcallact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldSvcallactMask) != 0
}

// SetSvcallact SVCall active
func (r *RegisterShcsrType) SetSvcallact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldSvcallactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldSvcallactMask)
	}
}

const (
	RegisterShcsrFieldMonitoractShift = 8
	RegisterShcsrFieldMonitoractMask  = 0x100
)

// GetMonitoract Debug monitor active
func (r *RegisterShcsrType) GetMonitoract() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldMonitoractMask) != 0
}

// SetMonitoract Debug monitor active
func (r *RegisterShcsrType) SetMonitoract(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldMonitoractMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldMonitoractMask)
	}
}

const (
	RegisterShcsrFieldPendsvactShift = 10
	RegisterShcsrFieldPendsvactMask  = 0x400
)

// GetPendsvact PendSV exception active
func (r *RegisterShcsrType) GetPendsvact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldPendsvactMask) != 0
}

// SetPendsvact PendSV exception active
func (r *RegisterShcsrType) SetPendsvact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldPendsvactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldPendsvactMask)
	}
}

const (
	RegisterShcsrFieldSystickactShift = 11
	RegisterShcsrFieldSystickactMask  = 0x800
)

// GetSystickact SysTick exception active
func (r *RegisterShcsrType) GetSystickact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldSystickactMask) != 0
}

// SetSystickact SysTick exception active
func (r *RegisterShcsrType) SetSystickact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldSystickactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldSystickactMask)
	}
}

const (
	RegisterShcsrFieldUsgfaultpendedShift = 12
	RegisterShcsrFieldUsgfaultpendedMask  = 0x1000
)

// GetUsgfaultpended UsageFault exception pending
func (r *RegisterShcsrType) GetUsgfaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldUsgfaultpendedMask) != 0
}

// SetUsgfaultpended UsageFault exception pending
func (r *RegisterShcsrType) SetUsgfaultpended(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldUsgfaultpendedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldUsgfaultpendedMask)
	}
}

const (
	RegisterShcsrFieldMemfaultpendedShift = 13
	RegisterShcsrFieldMemfaultpendedMask  = 0x2000
)

// GetMemfaultpended MemManage exception pending
func (r *RegisterShcsrType) GetMemfaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldMemfaultpendedMask) != 0
}

// SetMemfaultpended MemManage exception pending
func (r *RegisterShcsrType) SetMemfaultpended(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldMemfaultpendedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldMemfaultpendedMask)
	}
}

const (
	RegisterShcsrFieldBusfaultpendedShift = 14
	RegisterShcsrFieldBusfaultpendedMask  = 0x4000
)

// GetBusfaultpended BusFault pending
func (r *RegisterShcsrType) GetBusfaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldBusfaultpendedMask) != 0
}

// SetBusfaultpended BusFault pending
func (r *RegisterShcsrType) SetBusfaultpended(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldBusfaultpendedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldBusfaultpendedMask)
	}
}

const (
	RegisterShcsrFieldSvcallpendedShift = 15
	RegisterShcsrFieldSvcallpendedMask  = 0x8000
)

// GetSvcallpended SVCall pending
func (r *RegisterShcsrType) GetSvcallpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldSvcallpendedMask) != 0
}

// SetSvcallpended SVCall pending
func (r *RegisterShcsrType) SetSvcallpended(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldSvcallpendedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldSvcallpendedMask)
	}
}

const (
	RegisterShcsrFieldMemfaultenaShift = 16
	RegisterShcsrFieldMemfaultenaMask  = 0x10000
)

// GetMemfaultena MemManage enable
func (r *RegisterShcsrType) GetMemfaultena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldMemfaultenaMask) != 0
}

// SetMemfaultena MemManage enable
func (r *RegisterShcsrType) SetMemfaultena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldMemfaultenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldMemfaultenaMask)
	}
}

const (
	RegisterShcsrFieldBusfaultenaShift = 17
	RegisterShcsrFieldBusfaultenaMask  = 0x20000
)

// GetBusfaultena BusFault enable
func (r *RegisterShcsrType) GetBusfaultena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldBusfaultenaMask) != 0
}

// SetBusfaultena BusFault enable
func (r *RegisterShcsrType) SetBusfaultena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldBusfaultenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldBusfaultenaMask)
	}
}

const (
	RegisterShcsrFieldUsgfaultenaShift = 18
	RegisterShcsrFieldUsgfaultenaMask  = 0x40000
)

// GetUsgfaultena UsageFault enable
func (r *RegisterShcsrType) GetUsgfaultena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldUsgfaultenaMask) != 0
}

// SetUsgfaultena UsageFault enable
func (r *RegisterShcsrType) SetUsgfaultena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldUsgfaultenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldUsgfaultenaMask)
	}
}

const (
	RegisterShcsrFieldSecurefaultenaShift = 19
	RegisterShcsrFieldSecurefaultenaMask  = 0x80000
)

// GetSecurefaultena RES0
func (r *RegisterShcsrType) GetSecurefaultena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldSecurefaultenaMask) != 0
}

// SetSecurefaultena RES0
func (r *RegisterShcsrType) SetSecurefaultena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldSecurefaultenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldSecurefaultenaMask)
	}
}

const (
	RegisterShcsrFieldSecurefaultpendedShift = 20
	RegisterShcsrFieldSecurefaultpendedMask  = 0x100000
)

// GetSecurefaultpended SecureFault exception pended state bit, set to 1 to allow exception modification
func (r *RegisterShcsrType) GetSecurefaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldSecurefaultpendedMask) != 0
}

// SetSecurefaultpended SecureFault exception pended state bit, set to 1 to allow exception modification
func (r *RegisterShcsrType) SetSecurefaultpended(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldSecurefaultpendedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldSecurefaultpendedMask)
	}
}

const (
	RegisterShcsrFieldHardfaultpendedShift = 21
	RegisterShcsrFieldHardfaultpendedMask  = 0x200000
)

// GetHardfaultpended HardFault exception pended state bit, set to 1 to allow exception modification
func (r *RegisterShcsrType) GetHardfaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldHardfaultpendedMask) != 0
}

// SetHardfaultpended HardFault exception pended state bit, set to 1 to allow exception modification
func (r *RegisterShcsrType) SetHardfaultpended(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldHardfaultpendedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldHardfaultpendedMask)
	}
}

// RegisterCfsrType Indicates the cause of a MemManage fault, BusFault, or UsageFault
type RegisterCfsrType uint32

func (r *RegisterCfsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCfsrFieldMemmanageShift = 0
	RegisterCfsrFieldMemmanageMask  = 0xff
)

// GetMemmanage Provides information on MemManage exceptions
func (r *RegisterCfsrType) GetMemmanage() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfsrFieldMemmanageMask) >> RegisterCfsrFieldMemmanageShift)
}

// SetMemmanage Provides information on MemManage exceptions
func (r *RegisterCfsrType) SetMemmanage(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfsrFieldMemmanageMask)|(uint32(value)<<RegisterCfsrFieldMemmanageShift))
}

const (
	RegisterCfsrFieldBusfaultShift = 8
	RegisterCfsrFieldBusfaultMask  = 0xff00
)

// GetBusfault Provides information on BusFault exceptions
func (r *RegisterCfsrType) GetBusfault() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfsrFieldBusfaultMask) >> RegisterCfsrFieldBusfaultShift)
}

// SetBusfault Provides information on BusFault exceptions
func (r *RegisterCfsrType) SetBusfault(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfsrFieldBusfaultMask)|(uint32(value)<<RegisterCfsrFieldBusfaultShift))
}

const (
	RegisterCfsrFieldUsagefaultShift = 16
	RegisterCfsrFieldUsagefaultMask  = 0xffff0000
)

// GetUsagefault Provides information on UsageFault exceptions
func (r *RegisterCfsrType) GetUsagefault() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCfsrFieldUsagefaultMask) >> RegisterCfsrFieldUsagefaultShift)
}

// SetUsagefault Provides information on UsageFault exceptions
func (r *RegisterCfsrType) SetUsagefault(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfsrFieldUsagefaultMask)|(uint32(value)<<RegisterCfsrFieldUsagefaultShift))
}

// RegisterHfsrType Gives information about events that activate the HardFault handler
type RegisterHfsrType uint32

func (r *RegisterHfsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHfsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHfsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHfsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHfsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHfsrFieldVecttblShift = 1
	RegisterHfsrFieldVecttblMask  = 0x2
)

// GetVecttbl Indicates a HardFault on a vector table read during exception processing
func (r *RegisterHfsrType) GetVecttbl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHfsrFieldVecttblMask) != 0
}

// SetVecttbl Indicates a HardFault on a vector table read during exception processing
func (r *RegisterHfsrType) SetVecttbl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHfsrFieldVecttblMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHfsrFieldVecttblMask)
	}
}

const (
	RegisterHfsrFieldForcedShift = 30
	RegisterHfsrFieldForcedMask  = 0x40000000
)

// GetForced Indicates a forced HardFault, generated by escalation of a fault with configurable priority that cannot be handled, either because of priority or because it is disabled
func (r *RegisterHfsrType) GetForced() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHfsrFieldForcedMask) != 0
}

// SetForced Indicates a forced HardFault, generated by escalation of a fault with configurable priority that cannot be handled, either because of priority or because it is disabled
func (r *RegisterHfsrType) SetForced(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHfsrFieldForcedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHfsrFieldForcedMask)
	}
}

const (
	RegisterHfsrFieldDebugevtShift = 31
	RegisterHfsrFieldDebugevtMask  = 0x80000000
)

// GetDebugevt Reserved for Debug use. When writing to the register you must write 1 to this bit, otherwise behavior is unpredictable.
func (r *RegisterHfsrType) GetDebugevt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHfsrFieldDebugevtMask) != 0
}

// SetDebugevt Reserved for Debug use. When writing to the register you must write 1 to this bit, otherwise behavior is unpredictable.
func (r *RegisterHfsrType) SetDebugevt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHfsrFieldDebugevtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHfsrFieldDebugevtMask)
	}
}

// RegisterMmfarType Shows the address of the memory location that caused a Memory Protection Unit (MPU) fault
type RegisterMmfarType uint32

func (r *RegisterMmfarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMmfarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMmfarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMmfarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMmfarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterBfarType Shows the address associated with a precise data access BusFault
type RegisterBfarType uint32

func (r *RegisterBfarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBfarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBfarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBfarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBfarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterAfsrType provides fault status information
type RegisterAfsrType uint32

func (r *RegisterAfsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAfsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAfsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAfsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAfsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAfsrFieldIitcmShift = 0
	RegisterAfsrFieldIitcmMask  = 0x1
)

// GetIitcm Imprecise fault on ITCM interface
func (r *RegisterAfsrType) GetIitcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldIitcmMask) != 0
}

// SetIitcm Imprecise fault on ITCM interface
func (r *RegisterAfsrType) SetIitcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldIitcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldIitcmMask)
	}
}

const (
	RegisterAfsrFieldIdtcmShift = 1
	RegisterAfsrFieldIdtcmMask  = 0x2
)

// GetIdtcm Imprecise fault on DTCM interface
func (r *RegisterAfsrType) GetIdtcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldIdtcmMask) != 0
}

// SetIdtcm Imprecise fault on DTCM interface
func (r *RegisterAfsrType) SetIdtcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldIdtcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldIdtcmMask)
	}
}

const (
	RegisterAfsrFieldIpahbShift = 2
	RegisterAfsrFieldIpahbMask  = 0x4
)

// GetIpahb Imprecise fault on P-AHB interface
func (r *RegisterAfsrType) GetIpahb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldIpahbMask) != 0
}

// SetIpahb Imprecise fault on P-AHB interface
func (r *RegisterAfsrType) SetIpahb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldIpahbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldIpahbMask)
	}
}

const (
	RegisterAfsrFieldImaxiShift = 3
	RegisterAfsrFieldImaxiMask  = 0x8
)

// GetImaxi Imprecise fault on M-AXI interface
func (r *RegisterAfsrType) GetImaxi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldImaxiMask) != 0
}

// SetImaxi Imprecise fault on M-AXI interface
func (r *RegisterAfsrType) SetImaxi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldImaxiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldImaxiMask)
	}
}

const (
	RegisterAfsrFieldIeppbShift = 4
	RegisterAfsrFieldIeppbMask  = 0x10
)

// GetIeppb Imprecise fault on EPPB interface
func (r *RegisterAfsrType) GetIeppb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldIeppbMask) != 0
}

// SetIeppb Imprecise fault on EPPB interface
func (r *RegisterAfsrType) SetIeppb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldIeppbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldIeppbMask)
	}
}

const (
	RegisterAfsrFieldImaxitypeShift = 6
	RegisterAfsrFieldImaxitypeMask  = 0x40
)

// GetImaxitype AXI response that caused the imprecise fault. Only valid when AFSR.IMAXI is 1.
func (r *RegisterAfsrType) GetImaxitype() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldImaxitypeMask) != 0
}

// SetImaxitype AXI response that caused the imprecise fault. Only valid when AFSR.IMAXI is 1.
func (r *RegisterAfsrType) SetImaxitype(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldImaxitypeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldImaxitypeMask)
	}
}

const (
	RegisterAfsrFieldIeccShift = 7
	RegisterAfsrFieldIeccMask  = 0x80
)

// GetIecc Imprecise fault caused by uncorrectable ECC error
func (r *RegisterAfsrType) GetIecc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldIeccMask) != 0
}

// SetIecc Imprecise fault caused by uncorrectable ECC error
func (r *RegisterAfsrType) SetIecc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldIeccMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldIeccMask)
	}
}

const (
	RegisterAfsrFieldIpoisonShift = 9
	RegisterAfsrFieldIpoisonMask  = 0x200
)

// GetIpoison Imprecise BusFault because of RPOISON
func (r *RegisterAfsrType) GetIpoison() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldIpoisonMask) != 0
}

// SetIpoison Imprecise BusFault because of RPOISON
func (r *RegisterAfsrType) SetIpoison(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldIpoisonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldIpoisonMask)
	}
}

const (
	RegisterAfsrFieldPitcmShift = 10
	RegisterAfsrFieldPitcmMask  = 0x400
)

// GetPitcm Precise fault on ITCM interface
func (r *RegisterAfsrType) GetPitcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPitcmMask) != 0
}

// SetPitcm Precise fault on ITCM interface
func (r *RegisterAfsrType) SetPitcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldPitcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldPitcmMask)
	}
}

const (
	RegisterAfsrFieldPdtcmShift = 11
	RegisterAfsrFieldPdtcmMask  = 0x800
)

// GetPdtcm Precise fault on DTCM interface
func (r *RegisterAfsrType) GetPdtcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPdtcmMask) != 0
}

// SetPdtcm Precise fault on DTCM interface
func (r *RegisterAfsrType) SetPdtcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldPdtcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldPdtcmMask)
	}
}

const (
	RegisterAfsrFieldPpahbShift = 12
	RegisterAfsrFieldPpahbMask  = 0x1000
)

// GetPpahb Precise fault on Peripheral AHB (P-AHB) interface
func (r *RegisterAfsrType) GetPpahb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPpahbMask) != 0
}

// SetPpahb Precise fault on Peripheral AHB (P-AHB) interface
func (r *RegisterAfsrType) SetPpahb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldPpahbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldPpahbMask)
	}
}

const (
	RegisterAfsrFieldPmaxiShift = 13
	RegisterAfsrFieldPmaxiMask  = 0x2000
)

// GetPmaxi Precise fault on M-AXI interface
func (r *RegisterAfsrType) GetPmaxi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPmaxiMask) != 0
}

// SetPmaxi Precise fault on M-AXI interface
func (r *RegisterAfsrType) SetPmaxi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldPmaxiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldPmaxiMask)
	}
}

const (
	RegisterAfsrFieldPeppbShift = 14
	RegisterAfsrFieldPeppbMask  = 0x4000
)

// GetPeppb Precise fault on External Private Peripheral Bus (EPPB) interface
func (r *RegisterAfsrType) GetPeppb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPeppbMask) != 0
}

// SetPeppb Precise fault on External Private Peripheral Bus (EPPB) interface
func (r *RegisterAfsrType) SetPeppb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldPeppbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldPeppbMask)
	}
}

const (
	RegisterAfsrFieldPippbShift = 15
	RegisterAfsrFieldPippbMask  = 0x8000
)

// GetPippb Precise fault on Internal Private Peripheral Bus (EPPB) interface
func (r *RegisterAfsrType) GetPippb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPippbMask) != 0
}

// SetPippb Precise fault on Internal Private Peripheral Bus (EPPB) interface
func (r *RegisterAfsrType) SetPippb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldPippbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldPippbMask)
	}
}

const (
	RegisterAfsrFieldPmaxitypeShift = 16
	RegisterAfsrFieldPmaxitypeMask  = 0x10000
)

// GetPmaxitype AXI response that caused the precise fault. Only valid when AFSR.PMAXI is 1.
func (r *RegisterAfsrType) GetPmaxitype() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPmaxitypeMask) != 0
}

// SetPmaxitype AXI response that caused the precise fault. Only valid when AFSR.PMAXI is 1.
func (r *RegisterAfsrType) SetPmaxitype(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldPmaxitypeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldPmaxitypeMask)
	}
}

const (
	RegisterAfsrFieldPeccShift = 17
	RegisterAfsrFieldPeccMask  = 0x20000
)

// GetPecc Precise fault caused by uncorrectable ECC error
func (r *RegisterAfsrType) GetPecc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPeccMask) != 0
}

// SetPecc Precise fault caused by uncorrectable ECC error
func (r *RegisterAfsrType) SetPecc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldPeccMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldPeccMask)
	}
}

const (
	RegisterAfsrFieldPtguShift = 18
	RegisterAfsrFieldPtguMask  = 0x40000
)

// GetPtgu Precise fault caused by TGU security violation
func (r *RegisterAfsrType) GetPtgu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPtguMask) != 0
}

// SetPtgu Precise fault caused by TGU security violation
func (r *RegisterAfsrType) SetPtgu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldPtguMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldPtguMask)
	}
}

const (
	RegisterAfsrFieldPpoisonShift = 19
	RegisterAfsrFieldPpoisonMask  = 0x80000
)

// GetPpoison Precise fault caused by RPOISON or TEBRx.POISON
func (r *RegisterAfsrType) GetPpoison() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPpoisonMask) != 0
}

// SetPpoison Precise fault caused by RPOISON or TEBRx.POISON
func (r *RegisterAfsrType) SetPpoison(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldPpoisonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldPpoisonMask)
	}
}

const (
	RegisterAfsrFieldFitcmShift = 21
	RegisterAfsrFieldFitcmMask  = 0x200000
)

// GetFitcm Fetch fault on Instruction Tightly Coupled Memory (ITCM) interface
func (r *RegisterAfsrType) GetFitcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFitcmMask) != 0
}

// SetFitcm Fetch fault on Instruction Tightly Coupled Memory (ITCM) interface
func (r *RegisterAfsrType) SetFitcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldFitcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldFitcmMask)
	}
}

const (
	RegisterAfsrFieldFdtcmShift = 22
	RegisterAfsrFieldFdtcmMask  = 0x400000
)

// GetFdtcm Fetch fault on Data Tightly Coupled Memory (DTCM) interface
func (r *RegisterAfsrType) GetFdtcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFdtcmMask) != 0
}

// SetFdtcm Fetch fault on Data Tightly Coupled Memory (DTCM) interface
func (r *RegisterAfsrType) SetFdtcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldFdtcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldFdtcmMask)
	}
}

const (
	RegisterAfsrFieldFmaxiShift = 24
	RegisterAfsrFieldFmaxiMask  = 0x1000000
)

// GetFmaxi Fetch fault on Master AXI (M-AXI) interface
func (r *RegisterAfsrType) GetFmaxi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFmaxiMask) != 0
}

// SetFmaxi Fetch fault on Master AXI (M-AXI) interface
func (r *RegisterAfsrType) SetFmaxi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldFmaxiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldFmaxiMask)
	}
}

const (
	RegisterAfsrFieldFmaxitypeShift = 27
	RegisterAfsrFieldFmaxitypeMask  = 0x8000000
)

// GetFmaxitype AXI response that caused the fetch fault. Only valid when AFSR.FMAXI is 1.
func (r *RegisterAfsrType) GetFmaxitype() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFmaxitypeMask) != 0
}

// SetFmaxitype AXI response that caused the fetch fault. Only valid when AFSR.FMAXI is 1.
func (r *RegisterAfsrType) SetFmaxitype(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldFmaxitypeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldFmaxitypeMask)
	}
}

const (
	RegisterAfsrFieldFeccShift = 28
	RegisterAfsrFieldFeccMask  = 0x10000000
)

// GetFecc Fetch fault caused by uncorrectable Error Correcting Code (ECC) error
func (r *RegisterAfsrType) GetFecc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFeccMask) != 0
}

// SetFecc Fetch fault caused by uncorrectable Error Correcting Code (ECC) error
func (r *RegisterAfsrType) SetFecc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldFeccMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldFeccMask)
	}
}

const (
	RegisterAfsrFieldFtguShift = 29
	RegisterAfsrFieldFtguMask  = 0x20000000
)

// GetFtgu Fetch fault caused by TCM Gate Unit (TGU) security violation
func (r *RegisterAfsrType) GetFtgu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFtguMask) != 0
}

// SetFtgu Fetch fault caused by TCM Gate Unit (TGU) security violation
func (r *RegisterAfsrType) SetFtgu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldFtguMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldFtguMask)
	}
}

const (
	RegisterAfsrFieldFpoisonShift = 30
	RegisterAfsrFieldFpoisonMask  = 0x40000000
)

// GetFpoison Fetch fault caused by RPOISON or TEBRx.POISON
func (r *RegisterAfsrType) GetFpoison() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFpoisonMask) != 0
}

// SetFpoison Fetch fault caused by RPOISON or TEBRx.POISON
func (r *RegisterAfsrType) SetFpoison(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldFpoisonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldFpoisonMask)
	}
}

// RegisterIdpfr0Type Indicates the version of the Reliability, Availability, and Serviceability (RAS) extension supported
type RegisterIdpfr0Type uint32

func (r *RegisterIdpfr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdpfr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdpfr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdpfr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdpfr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIdpfr0FieldState0Shift = 0
	RegisterIdpfr0FieldState0Mask  = 0xf
)

// GetState0 A32 instruction set support
func (r *RegisterIdpfr0Type) GetState0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdpfr0FieldState0Mask) >> RegisterIdpfr0FieldState0Shift)
}

// SetState0 A32 instruction set support
func (r *RegisterIdpfr0Type) SetState0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdpfr0FieldState0Mask)|(uint32(value)<<RegisterIdpfr0FieldState0Shift))
}

const (
	RegisterIdpfr0FieldState1Shift = 4
	RegisterIdpfr0FieldState1Mask  = 0xf0
)

// GetState1 T32 instruction set support
func (r *RegisterIdpfr0Type) GetState1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdpfr0FieldState1Mask) >> RegisterIdpfr0FieldState1Shift)
}

// SetState1 T32 instruction set support
func (r *RegisterIdpfr0Type) SetState1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdpfr0FieldState1Mask)|(uint32(value)<<RegisterIdpfr0FieldState1Shift))
}

const (
	RegisterIdpfr0FieldRasShift = 28
	RegisterIdpfr0FieldRasMask  = 0xf0000000
)

// GetRas Identifies which version of the RAS architecture is implemented
func (r *RegisterIdpfr0Type) GetRas() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdpfr0FieldRasMask) >> RegisterIdpfr0FieldRasShift)
}

// SetRas Identifies which version of the RAS architecture is implemented
func (r *RegisterIdpfr0Type) SetRas(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdpfr0FieldRasMask)|(uint32(value)<<RegisterIdpfr0FieldRasShift))
}

// RegisterIdpfr1Type Gives information about the programmers' model and Extensions support
type RegisterIdpfr1Type uint32

func (r *RegisterIdpfr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdpfr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdpfr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdpfr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdpfr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIdpfr1FieldSecurityShift = 4
	RegisterIdpfr1FieldSecurityMask  = 0xf0
)

// GetSecurity Identifies whether the Security Extension in implemented
func (r *RegisterIdpfr1Type) GetSecurity() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdpfr1FieldSecurityMask) >> RegisterIdpfr1FieldSecurityShift)
}

// SetSecurity Identifies whether the Security Extension in implemented
func (r *RegisterIdpfr1Type) SetSecurity(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdpfr1FieldSecurityMask)|(uint32(value)<<RegisterIdpfr1FieldSecurityShift))
}

const (
	RegisterIdpfr1FieldMprogmodShift = 8
	RegisterIdpfr1FieldMprogmodMask  = 0xf00
)

// GetMprogmod M-profile programmers' model
func (r *RegisterIdpfr1Type) GetMprogmod() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdpfr1FieldMprogmodMask) >> RegisterIdpfr1FieldMprogmodShift)
}

// SetMprogmod M-profile programmers' model
func (r *RegisterIdpfr1Type) SetMprogmod(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdpfr1FieldMprogmodMask)|(uint32(value)<<RegisterIdpfr1FieldMprogmodShift))
}

// RegisterIddfr0Type no information available
type RegisterIddfr0Type uint32

func (r *RegisterIddfr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIddfr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIddfr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIddfr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIddfr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIddfr0FieldMprofdbgShift = 20
	RegisterIddfr0FieldMprofdbgMask  = 0xf00000
)

// GetMprofdbg Indicates the supported M-Proﬁle debug architecture
func (r *RegisterIddfr0Type) GetMprofdbg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIddfr0FieldMprofdbgMask) >> RegisterIddfr0FieldMprofdbgShift)
}

// SetMprofdbg Indicates the supported M-Proﬁle debug architecture
func (r *RegisterIddfr0Type) SetMprofdbg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIddfr0FieldMprofdbgMask)|(uint32(value)<<RegisterIddfr0FieldMprofdbgShift))
}

const (
	RegisterIddfr0FieldUdeShift = 28
	RegisterIddfr0FieldUdeMask  = 0xf0000000
)

// GetUde Indicates support for the Unprivileged Debug Extension
func (r *RegisterIddfr0Type) GetUde() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIddfr0FieldUdeMask) >> RegisterIddfr0FieldUdeShift)
}

// SetUde Indicates support for the Unprivileged Debug Extension
func (r *RegisterIddfr0Type) SetUde(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIddfr0FieldUdeMask)|(uint32(value)<<RegisterIddfr0FieldUdeShift))
}

// RegisterIdafr0Type The ID_AFR0 is fully Reserved, RES0
type RegisterIdafr0Type uint32

func (r *RegisterIdafr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdafr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdafr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdafr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdafr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterIdmmfr0Type Gives information about the implemented memory model and memory management support
type RegisterIdmmfr0Type uint32

func (r *RegisterIdmmfr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdmmfr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdmmfr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdmmfr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdmmfr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIdmmfr0FieldPmsaShift = 4
	RegisterIdmmfr0FieldPmsaMask  = 0xf0
)

// GetPmsa Indicates support for the protected memory system architecture (PMSA)
func (r *RegisterIdmmfr0Type) GetPmsa() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr0FieldPmsaMask) >> RegisterIdmmfr0FieldPmsaShift)
}

// SetPmsa Indicates support for the protected memory system architecture (PMSA)
func (r *RegisterIdmmfr0Type) SetPmsa(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr0FieldPmsaMask)|(uint32(value)<<RegisterIdmmfr0FieldPmsaShift))
}

const (
	RegisterIdmmfr0FieldOutershrShift = 8
	RegisterIdmmfr0FieldOutershrMask  = 0xf00
)

// GetOutershr Indicates the outermost Shareability domain implemented
func (r *RegisterIdmmfr0Type) GetOutershr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr0FieldOutershrMask) >> RegisterIdmmfr0FieldOutershrShift)
}

// SetOutershr Indicates the outermost Shareability domain implemented
func (r *RegisterIdmmfr0Type) SetOutershr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr0FieldOutershrMask)|(uint32(value)<<RegisterIdmmfr0FieldOutershrShift))
}

const (
	RegisterIdmmfr0FieldSharelvlShift = 12
	RegisterIdmmfr0FieldSharelvlMask  = 0xf000
)

// GetSharelvl Indicates the number of Shareability levels implemented
func (r *RegisterIdmmfr0Type) GetSharelvl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr0FieldSharelvlMask) >> RegisterIdmmfr0FieldSharelvlShift)
}

// SetSharelvl Indicates the number of Shareability levels implemented
func (r *RegisterIdmmfr0Type) SetSharelvl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr0FieldSharelvlMask)|(uint32(value)<<RegisterIdmmfr0FieldSharelvlShift))
}

const (
	RegisterIdmmfr0FieldTcmShift = 16
	RegisterIdmmfr0FieldTcmMask  = 0xf0000
)

// GetTcm Indicates support for TCMs
func (r *RegisterIdmmfr0Type) GetTcm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr0FieldTcmMask) >> RegisterIdmmfr0FieldTcmShift)
}

// SetTcm Indicates support for TCMs
func (r *RegisterIdmmfr0Type) SetTcm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr0FieldTcmMask)|(uint32(value)<<RegisterIdmmfr0FieldTcmShift))
}

const (
	RegisterIdmmfr0FieldAuxregShift = 20
	RegisterIdmmfr0FieldAuxregMask  = 0xf00000
)

// GetAuxreg Indicates support for Auxiliary Control registers
func (r *RegisterIdmmfr0Type) GetAuxreg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr0FieldAuxregMask) >> RegisterIdmmfr0FieldAuxregShift)
}

// SetAuxreg Indicates support for Auxiliary Control registers
func (r *RegisterIdmmfr0Type) SetAuxreg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr0FieldAuxregMask)|(uint32(value)<<RegisterIdmmfr0FieldAuxregShift))
}

// RegisterIdmmfr1Type Gives information about the IMPLEMENTATION DEFINED memory model and memory management support. This register is Reserved, RES0.
type RegisterIdmmfr1Type uint32

func (r *RegisterIdmmfr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdmmfr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdmmfr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdmmfr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdmmfr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterIdmmfr2Type Gives information about the implemented memory model and memory management support
type RegisterIdmmfr2Type uint32

func (r *RegisterIdmmfr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdmmfr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdmmfr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdmmfr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdmmfr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIdmmfr2FieldWfistallShift = 24
	RegisterIdmmfr2FieldWfistallMask  = 0xf000000
)

// GetWfistall Indicates the support for Wait For Interrupt (WFI) stalling
func (r *RegisterIdmmfr2Type) GetWfistall() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr2FieldWfistallMask) >> RegisterIdmmfr2FieldWfistallShift)
}

// SetWfistall Indicates the support for Wait For Interrupt (WFI) stalling
func (r *RegisterIdmmfr2Type) SetWfistall(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr2FieldWfistallMask)|(uint32(value)<<RegisterIdmmfr2FieldWfistallShift))
}

// RegisterIdmmfr3Type Gives information about the implemented memory model and memory management support
type RegisterIdmmfr3Type uint32

func (r *RegisterIdmmfr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdmmfr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdmmfr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdmmfr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdmmfr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIdmmfr3FieldCmaintvaShift = 0
	RegisterIdmmfr3FieldCmaintvaMask  = 0xf
)

// GetCmaintva Indicates the supported cache maintenance operations by address
func (r *RegisterIdmmfr3Type) GetCmaintva() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr3FieldCmaintvaMask) >> RegisterIdmmfr3FieldCmaintvaShift)
}

// SetCmaintva Indicates the supported cache maintenance operations by address
func (r *RegisterIdmmfr3Type) SetCmaintva(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr3FieldCmaintvaMask)|(uint32(value)<<RegisterIdmmfr3FieldCmaintvaShift))
}

const (
	RegisterIdmmfr3FieldCmaintswShift = 4
	RegisterIdmmfr3FieldCmaintswMask  = 0xf0
)

// GetCmaintsw Indicates the supported cache maintenance operations by set/way
func (r *RegisterIdmmfr3Type) GetCmaintsw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr3FieldCmaintswMask) >> RegisterIdmmfr3FieldCmaintswShift)
}

// SetCmaintsw Indicates the supported cache maintenance operations by set/way
func (r *RegisterIdmmfr3Type) SetCmaintsw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr3FieldCmaintswMask)|(uint32(value)<<RegisterIdmmfr3FieldCmaintswShift))
}

const (
	RegisterIdmmfr3FieldBpmaintShift = 8
	RegisterIdmmfr3FieldBpmaintMask  = 0xf00
)

// GetBpmaint Indicates the supported branch predictor maintenance
func (r *RegisterIdmmfr3Type) GetBpmaint() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr3FieldBpmaintMask) >> RegisterIdmmfr3FieldBpmaintShift)
}

// SetBpmaint Indicates the supported branch predictor maintenance
func (r *RegisterIdmmfr3Type) SetBpmaint(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr3FieldBpmaintMask)|(uint32(value)<<RegisterIdmmfr3FieldBpmaintShift))
}

// RegisterIdisar0Type Gives information about the implemented instruction set
type RegisterIdisar0Type uint32

func (r *RegisterIdisar0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdisar0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdisar0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdisar0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdisar0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIdisar0FieldBitcountShift = 4
	RegisterIdisar0FieldBitcountMask  = 0xf0
)

// GetBitcount Indicates the supported bit count instructions
func (r *RegisterIdisar0Type) GetBitcount() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar0FieldBitcountMask) >> RegisterIdisar0FieldBitcountShift)
}

// SetBitcount Indicates the supported bit count instructions
func (r *RegisterIdisar0Type) SetBitcount(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar0FieldBitcountMask)|(uint32(value)<<RegisterIdisar0FieldBitcountShift))
}

const (
	RegisterIdisar0FieldBitfieldShift = 8
	RegisterIdisar0FieldBitfieldMask  = 0xf00
)

// GetBitfield Indicates the supported bit field instructions
func (r *RegisterIdisar0Type) GetBitfield() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar0FieldBitfieldMask) >> RegisterIdisar0FieldBitfieldShift)
}

// SetBitfield Indicates the supported bit field instructions
func (r *RegisterIdisar0Type) SetBitfield(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar0FieldBitfieldMask)|(uint32(value)<<RegisterIdisar0FieldBitfieldShift))
}

const (
	RegisterIdisar0FieldCmpbranchShift = 12
	RegisterIdisar0FieldCmpbranchMask  = 0xf000
)

// GetCmpbranch Indicates the supported combined Compare and Branch instructions
func (r *RegisterIdisar0Type) GetCmpbranch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar0FieldCmpbranchMask) >> RegisterIdisar0FieldCmpbranchShift)
}

// SetCmpbranch Indicates the supported combined Compare and Branch instructions
func (r *RegisterIdisar0Type) SetCmpbranch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar0FieldCmpbranchMask)|(uint32(value)<<RegisterIdisar0FieldCmpbranchShift))
}

const (
	RegisterIdisar0FieldCoprocShift = 16
	RegisterIdisar0FieldCoprocMask  = 0xf0000
)

// GetCoproc Indicates the supported coprocessor instructions
func (r *RegisterIdisar0Type) GetCoproc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar0FieldCoprocMask) >> RegisterIdisar0FieldCoprocShift)
}

// SetCoproc Indicates the supported coprocessor instructions
func (r *RegisterIdisar0Type) SetCoproc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar0FieldCoprocMask)|(uint32(value)<<RegisterIdisar0FieldCoprocShift))
}

const (
	RegisterIdisar0FieldDebugShift = 20
	RegisterIdisar0FieldDebugMask  = 0xf00000
)

// GetDebug Indicates the implemented debug instructions
func (r *RegisterIdisar0Type) GetDebug() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar0FieldDebugMask) >> RegisterIdisar0FieldDebugShift)
}

// SetDebug Indicates the implemented debug instructions
func (r *RegisterIdisar0Type) SetDebug(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar0FieldDebugMask)|(uint32(value)<<RegisterIdisar0FieldDebugShift))
}

const (
	RegisterIdisar0FieldDivideShift = 24
	RegisterIdisar0FieldDivideMask  = 0xf000000
)

// GetDivide Indicates the supported divide instructions
func (r *RegisterIdisar0Type) GetDivide() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar0FieldDivideMask) >> RegisterIdisar0FieldDivideShift)
}

// SetDivide Indicates the supported divide instructions
func (r *RegisterIdisar0Type) SetDivide(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar0FieldDivideMask)|(uint32(value)<<RegisterIdisar0FieldDivideShift))
}

// RegisterIdisar1Type Gives information about the implemented instruction set
type RegisterIdisar1Type uint32

func (r *RegisterIdisar1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdisar1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdisar1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdisar1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdisar1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIdisar1FieldExtendShift = 12
	RegisterIdisar1FieldExtendMask  = 0xf000
)

// GetExtend Indicates the implemented Extend instructions
func (r *RegisterIdisar1Type) GetExtend() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar1FieldExtendMask) >> RegisterIdisar1FieldExtendShift)
}

// SetExtend Indicates the implemented Extend instructions
func (r *RegisterIdisar1Type) SetExtend(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar1FieldExtendMask)|(uint32(value)<<RegisterIdisar1FieldExtendShift))
}

const (
	RegisterIdisar1FieldIfthenShift = 16
	RegisterIdisar1FieldIfthenMask  = 0xf0000
)

// GetIfthen Indicates the implemented If-Then instructions
func (r *RegisterIdisar1Type) GetIfthen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar1FieldIfthenMask) >> RegisterIdisar1FieldIfthenShift)
}

// SetIfthen Indicates the implemented If-Then instructions
func (r *RegisterIdisar1Type) SetIfthen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar1FieldIfthenMask)|(uint32(value)<<RegisterIdisar1FieldIfthenShift))
}

const (
	RegisterIdisar1FieldImmediateShift = 20
	RegisterIdisar1FieldImmediateMask  = 0xf00000
)

// GetImmediate Indicates the implemented data-processing instructions with long immediates
func (r *RegisterIdisar1Type) GetImmediate() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar1FieldImmediateMask) >> RegisterIdisar1FieldImmediateShift)
}

// SetImmediate Indicates the implemented data-processing instructions with long immediates
func (r *RegisterIdisar1Type) SetImmediate(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar1FieldImmediateMask)|(uint32(value)<<RegisterIdisar1FieldImmediateShift))
}

const (
	RegisterIdisar1FieldInterworkShift = 24
	RegisterIdisar1FieldInterworkMask  = 0xf000000
)

// GetInterwork Indicates the implemented interworking instructions
func (r *RegisterIdisar1Type) GetInterwork() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar1FieldInterworkMask) >> RegisterIdisar1FieldInterworkShift)
}

// SetInterwork Indicates the implemented interworking instructions
func (r *RegisterIdisar1Type) SetInterwork(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar1FieldInterworkMask)|(uint32(value)<<RegisterIdisar1FieldInterworkShift))
}

// RegisterIdisar2Type Gives information about the implemented instruction set
type RegisterIdisar2Type uint32

func (r *RegisterIdisar2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdisar2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdisar2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdisar2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdisar2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIdisar2FieldLoadstoreShift = 0
	RegisterIdisar2FieldLoadstoreMask  = 0xf
)

// GetLoadstore Indicates the implemented additional load/store instructions
func (r *RegisterIdisar2Type) GetLoadstore() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldLoadstoreMask) >> RegisterIdisar2FieldLoadstoreShift)
}

// SetLoadstore Indicates the implemented additional load/store instructions
func (r *RegisterIdisar2Type) SetLoadstore(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldLoadstoreMask)|(uint32(value)<<RegisterIdisar2FieldLoadstoreShift))
}

const (
	RegisterIdisar2FieldMemhintShift = 4
	RegisterIdisar2FieldMemhintMask  = 0xf0
)

// GetMemhint Memory hint. Indicates the implemented memory hint instructions
func (r *RegisterIdisar2Type) GetMemhint() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldMemhintMask) >> RegisterIdisar2FieldMemhintShift)
}

// SetMemhint Memory hint. Indicates the implemented memory hint instructions
func (r *RegisterIdisar2Type) SetMemhint(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldMemhintMask)|(uint32(value)<<RegisterIdisar2FieldMemhintShift))
}

const (
	RegisterIdisar2FieldMultiaccessintShift = 8
	RegisterIdisar2FieldMultiaccessintMask  = 0xf00
)

// GetMultiaccessint Multi-access instructions. Indicates the support for interruptible multi-access instructions
func (r *RegisterIdisar2Type) GetMultiaccessint() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldMultiaccessintMask) >> RegisterIdisar2FieldMultiaccessintShift)
}

// SetMultiaccessint Multi-access instructions. Indicates the support for interruptible multi-access instructions
func (r *RegisterIdisar2Type) SetMultiaccessint(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldMultiaccessintMask)|(uint32(value)<<RegisterIdisar2FieldMultiaccessintShift))
}

const (
	RegisterIdisar2FieldMultShift = 12
	RegisterIdisar2FieldMultMask  = 0xf000
)

// GetMult Indicates the implemented additional Multiply instructions
func (r *RegisterIdisar2Type) GetMult() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldMultMask) >> RegisterIdisar2FieldMultShift)
}

// SetMult Indicates the implemented additional Multiply instructions
func (r *RegisterIdisar2Type) SetMult(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldMultMask)|(uint32(value)<<RegisterIdisar2FieldMultShift))
}

const (
	RegisterIdisar2FieldMultsShift = 16
	RegisterIdisar2FieldMultsMask  = 0xf0000
)

// GetMults Indicates the implemented Advanced Signed Multiple instructions
func (r *RegisterIdisar2Type) GetMults() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldMultsMask) >> RegisterIdisar2FieldMultsShift)
}

// SetMults Indicates the implemented Advanced Signed Multiple instructions
func (r *RegisterIdisar2Type) SetMults(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldMultsMask)|(uint32(value)<<RegisterIdisar2FieldMultsShift))
}

const (
	RegisterIdisar2FieldMultuShift = 20
	RegisterIdisar2FieldMultuMask  = 0xf00000
)

// GetMultu Indicates the implemented Advanced Unsigned Multiple instructions
func (r *RegisterIdisar2Type) GetMultu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldMultuMask) >> RegisterIdisar2FieldMultuShift)
}

// SetMultu Indicates the implemented Advanced Unsigned Multiple instructions
func (r *RegisterIdisar2Type) SetMultu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldMultuMask)|(uint32(value)<<RegisterIdisar2FieldMultuShift))
}

const (
	RegisterIdisar2FieldReversalShift = 28
	RegisterIdisar2FieldReversalMask  = 0xf0000000
)

// GetReversal Indicates the implemented reversal instructions
func (r *RegisterIdisar2Type) GetReversal() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldReversalMask) >> RegisterIdisar2FieldReversalShift)
}

// SetReversal Indicates the implemented reversal instructions
func (r *RegisterIdisar2Type) SetReversal(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldReversalMask)|(uint32(value)<<RegisterIdisar2FieldReversalShift))
}

// RegisterIdisar3Type Gives information about the implemented instruction set
type RegisterIdisar3Type uint32

func (r *RegisterIdisar3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdisar3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdisar3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdisar3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdisar3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIdisar3FieldSaturateShift = 0
	RegisterIdisar3FieldSaturateMask  = 0xf
)

// GetSaturate Indicates the implemented saturating instructions
func (r *RegisterIdisar3Type) GetSaturate() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldSaturateMask) >> RegisterIdisar3FieldSaturateShift)
}

// SetSaturate Indicates the implemented saturating instructions
func (r *RegisterIdisar3Type) SetSaturate(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldSaturateMask)|(uint32(value)<<RegisterIdisar3FieldSaturateShift))
}

const (
	RegisterIdisar3FieldSimdShift = 4
	RegisterIdisar3FieldSimdMask  = 0xf0
)

// GetSimd Indicates the implemented SIMD instructions
func (r *RegisterIdisar3Type) GetSimd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldSimdMask) >> RegisterIdisar3FieldSimdShift)
}

// SetSimd Indicates the implemented SIMD instructions
func (r *RegisterIdisar3Type) SetSimd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldSimdMask)|(uint32(value)<<RegisterIdisar3FieldSimdShift))
}

const (
	RegisterIdisar3FieldSvcShift = 8
	RegisterIdisar3FieldSvcMask  = 0xf00
)

// GetSvc Indicates the implemented SVC instructions
func (r *RegisterIdisar3Type) GetSvc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldSvcMask) >> RegisterIdisar3FieldSvcShift)
}

// SetSvc Indicates the implemented SVC instructions
func (r *RegisterIdisar3Type) SetSvc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldSvcMask)|(uint32(value)<<RegisterIdisar3FieldSvcShift))
}

const (
	RegisterIdisar3FieldSynchprimShift = 12
	RegisterIdisar3FieldSynchprimMask  = 0xf000
)

// GetSynchprim Used with ID_ISAR4.SynchPrim_frac to indicate the implemented synchronization primitive instructions
func (r *RegisterIdisar3Type) GetSynchprim() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldSynchprimMask) >> RegisterIdisar3FieldSynchprimShift)
}

// SetSynchprim Used with ID_ISAR4.SynchPrim_frac to indicate the implemented synchronization primitive instructions
func (r *RegisterIdisar3Type) SetSynchprim(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldSynchprimMask)|(uint32(value)<<RegisterIdisar3FieldSynchprimShift))
}

const (
	RegisterIdisar3FieldTabbranchShift = 16
	RegisterIdisar3FieldTabbranchMask  = 0xf0000
)

// GetTabbranch Indicates the implemented table branch instructions
func (r *RegisterIdisar3Type) GetTabbranch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldTabbranchMask) >> RegisterIdisar3FieldTabbranchShift)
}

// SetTabbranch Indicates the implemented table branch instructions
func (r *RegisterIdisar3Type) SetTabbranch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldTabbranchMask)|(uint32(value)<<RegisterIdisar3FieldTabbranchShift))
}

const (
	RegisterIdisar3FieldT32copyShift = 20
	RegisterIdisar3FieldT32copyMask  = 0xf00000
)

// GetT32copy Indicates the support for T32 non-flag setting MOV instructions
func (r *RegisterIdisar3Type) GetT32copy() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldT32copyMask) >> RegisterIdisar3FieldT32copyShift)
}

// SetT32copy Indicates the support for T32 non-flag setting MOV instructions
func (r *RegisterIdisar3Type) SetT32copy(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldT32copyMask)|(uint32(value)<<RegisterIdisar3FieldT32copyShift))
}

const (
	RegisterIdisar3FieldTruenopShift = 24
	RegisterIdisar3FieldTruenopMask  = 0xf000000
)

// GetTruenop Indicates the implemented true NOP instructions
func (r *RegisterIdisar3Type) GetTruenop() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldTruenopMask) >> RegisterIdisar3FieldTruenopShift)
}

// SetTruenop Indicates the implemented true NOP instructions
func (r *RegisterIdisar3Type) SetTruenop(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldTruenopMask)|(uint32(value)<<RegisterIdisar3FieldTruenopShift))
}

// RegisterIdisar4Type Gives information about the implemented instruction set
type RegisterIdisar4Type uint32

func (r *RegisterIdisar4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdisar4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdisar4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdisar4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdisar4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIdisar4FieldUnprivShift = 0
	RegisterIdisar4FieldUnprivMask  = 0xf
)

// GetUnpriv Indicates the implemented unprivileged instructions
func (r *RegisterIdisar4Type) GetUnpriv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar4FieldUnprivMask) >> RegisterIdisar4FieldUnprivShift)
}

// SetUnpriv Indicates the implemented unprivileged instructions
func (r *RegisterIdisar4Type) SetUnpriv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar4FieldUnprivMask)|(uint32(value)<<RegisterIdisar4FieldUnprivShift))
}

const (
	RegisterIdisar4FieldWithshiftsShift = 4
	RegisterIdisar4FieldWithshiftsMask  = 0xf0
)

// GetWithshifts Indicates the support for writeback addressing modes
func (r *RegisterIdisar4Type) GetWithshifts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar4FieldWithshiftsMask) >> RegisterIdisar4FieldWithshiftsShift)
}

// SetWithshifts Indicates the support for writeback addressing modes
func (r *RegisterIdisar4Type) SetWithshifts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar4FieldWithshiftsMask)|(uint32(value)<<RegisterIdisar4FieldWithshiftsShift))
}

const (
	RegisterIdisar4FieldWritebackShift = 8
	RegisterIdisar4FieldWritebackMask  = 0xf00
)

// GetWriteback Indicates the support for writeback addressing modes
func (r *RegisterIdisar4Type) GetWriteback() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar4FieldWritebackMask) >> RegisterIdisar4FieldWritebackShift)
}

// SetWriteback Indicates the support for writeback addressing modes
func (r *RegisterIdisar4Type) SetWriteback(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar4FieldWritebackMask)|(uint32(value)<<RegisterIdisar4FieldWritebackShift))
}

const (
	RegisterIdisar4FieldBarrierShift = 16
	RegisterIdisar4FieldBarrierMask  = 0xf0000
)

// GetBarrier Indicates the implemented Barrier instructions
func (r *RegisterIdisar4Type) GetBarrier() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar4FieldBarrierMask) >> RegisterIdisar4FieldBarrierShift)
}

// SetBarrier Indicates the implemented Barrier instructions
func (r *RegisterIdisar4Type) SetBarrier(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar4FieldBarrierMask)|(uint32(value)<<RegisterIdisar4FieldBarrierShift))
}

const (
	RegisterIdisar4FieldSynchprimfracShift = 20
	RegisterIdisar4FieldSynchprimfracMask  = 0xf00000
)

// GetSynchprimfrac Used in conjunction with ID_ISAR3.SynchPrim to indicate the implemented synchronization primitive instructions
func (r *RegisterIdisar4Type) GetSynchprimfrac() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar4FieldSynchprimfracMask) >> RegisterIdisar4FieldSynchprimfracShift)
}

// SetSynchprimfrac Used in conjunction with ID_ISAR3.SynchPrim to indicate the implemented synchronization primitive instructions
func (r *RegisterIdisar4Type) SetSynchprimfrac(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar4FieldSynchprimfracMask)|(uint32(value)<<RegisterIdisar4FieldSynchprimfracShift))
}

const (
	RegisterIdisar4FieldPsrmShift = 24
	RegisterIdisar4FieldPsrmMask  = 0xf000000
)

// GetPsrm Indicates the implemented M-profile instructions to modify the PSRs
func (r *RegisterIdisar4Type) GetPsrm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar4FieldPsrmMask) >> RegisterIdisar4FieldPsrmShift)
}

// SetPsrm Indicates the implemented M-profile instructions to modify the PSRs
func (r *RegisterIdisar4Type) SetPsrm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar4FieldPsrmMask)|(uint32(value)<<RegisterIdisar4FieldPsrmShift))
}

// RegisterClidrType The CLIDR identifies the type of caches implemented and the level of coherency and unification
type RegisterClidrType uint32

func (r *RegisterClidrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterClidrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterClidrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterClidrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterClidrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterClidrFieldCtype1Shift = 0
	RegisterClidrFieldCtype1Mask  = 0x7
)

// GetCtype1 L1 cache type
func (r *RegisterClidrType) GetCtype1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldCtype1Mask) >> RegisterClidrFieldCtype1Shift)
}

// SetCtype1 L1 cache type
func (r *RegisterClidrType) SetCtype1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldCtype1Mask)|(uint32(value)<<RegisterClidrFieldCtype1Shift))
}

const (
	RegisterClidrFieldLouisShift = 21
	RegisterClidrFieldLouisMask  = 0xe00000
)

// GetLouis Level of Unification Inner Shareable
func (r *RegisterClidrType) GetLouis() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldLouisMask) >> RegisterClidrFieldLouisShift)
}

// SetLouis Level of Unification Inner Shareable
func (r *RegisterClidrType) SetLouis(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldLouisMask)|(uint32(value)<<RegisterClidrFieldLouisShift))
}

const (
	RegisterClidrFieldLocShift = 24
	RegisterClidrFieldLocMask  = 0x7000000
)

// GetLoc Level of Coherency
func (r *RegisterClidrType) GetLoc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldLocMask) >> RegisterClidrFieldLocShift)
}

// SetLoc Level of Coherency
func (r *RegisterClidrType) SetLoc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldLocMask)|(uint32(value)<<RegisterClidrFieldLocShift))
}

const (
	RegisterClidrFieldLouuShift = 27
	RegisterClidrFieldLouuMask  = 0x38000000
)

// GetLouu Level of Unification Uniprocessor
func (r *RegisterClidrType) GetLouu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldLouuMask) >> RegisterClidrFieldLouuShift)
}

// SetLouu Level of Unification Uniprocessor
func (r *RegisterClidrType) SetLouu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldLouuMask)|(uint32(value)<<RegisterClidrFieldLouuShift))
}

const (
	RegisterClidrFieldIcbShift = 30
	RegisterClidrFieldIcbMask  = 0xc0000000
)

// GetIcb Inner cache boundary. The Cortex-M55 processor supports inner cacheability on the bus. Therefore, this field cannot disclose any information.
func (r *RegisterClidrType) GetIcb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldIcbMask) >> RegisterClidrFieldIcbShift)
}

// SetIcb Inner cache boundary. The Cortex-M55 processor supports inner cacheability on the bus. Therefore, this field cannot disclose any information.
func (r *RegisterClidrType) SetIcb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldIcbMask)|(uint32(value)<<RegisterClidrFieldIcbShift))
}

// RegisterCtrType The CTR provides information about the architecture of the caches
type RegisterCtrType uint32

func (r *RegisterCtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCtrFieldIminlineShift = 0
	RegisterCtrFieldIminlineMask  = 0xf
)

// GetIminline Instruction cache minimum line length. Log2 of the number of words in the smallest cache line of all the instruction caches that are controlled by the processor.
func (r *RegisterCtrType) GetIminline() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldIminlineMask) >> RegisterCtrFieldIminlineShift)
}

// SetIminline Instruction cache minimum line length. Log2 of the number of words in the smallest cache line of all the instruction caches that are controlled by the processor.
func (r *RegisterCtrType) SetIminline(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldIminlineMask)|(uint32(value)<<RegisterCtrFieldIminlineShift))
}

const (
	RegisterCtrFieldDminlineShift = 16
	RegisterCtrFieldDminlineMask  = 0xf0000
)

// GetDminline Data cache minimum line length. Log2 of the number of words in the smallest cache line of all the data caches and unified caches that are controlled by the processor.
func (r *RegisterCtrType) GetDminline() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldDminlineMask) >> RegisterCtrFieldDminlineShift)
}

// SetDminline Data cache minimum line length. Log2 of the number of words in the smallest cache line of all the data caches and unified caches that are controlled by the processor.
func (r *RegisterCtrType) SetDminline(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldDminlineMask)|(uint32(value)<<RegisterCtrFieldDminlineShift))
}

const (
	RegisterCtrFieldErgShift = 20
	RegisterCtrFieldErgMask  = 0xf00000
)

// GetErg Exclusives Reservation Granule
func (r *RegisterCtrType) GetErg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldErgMask) >> RegisterCtrFieldErgShift)
}

// SetErg Exclusives Reservation Granule
func (r *RegisterCtrType) SetErg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldErgMask)|(uint32(value)<<RegisterCtrFieldErgShift))
}

const (
	RegisterCtrFieldCwgShift = 24
	RegisterCtrFieldCwgMask  = 0xf000000
)

// GetCwg Cache Write-Back Granule
func (r *RegisterCtrType) GetCwg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldCwgMask) >> RegisterCtrFieldCwgShift)
}

// SetCwg Cache Write-Back Granule
func (r *RegisterCtrType) SetCwg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldCwgMask)|(uint32(value)<<RegisterCtrFieldCwgShift))
}

const (
	RegisterCtrFieldFormatShift = 29
	RegisterCtrFieldFormatMask  = 0xe0000000
)

// GetFormat Cache type register format
func (r *RegisterCtrType) GetFormat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldFormatMask) >> RegisterCtrFieldFormatShift)
}

// SetFormat Cache type register format
func (r *RegisterCtrType) SetFormat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldFormatMask)|(uint32(value)<<RegisterCtrFieldFormatShift))
}

// RegisterCcsidrType Provides information about the architecture of the Level 1 (L1) instruction or data cache that the CSSELR selects.
type RegisterCcsidrType uint32

func (r *RegisterCcsidrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcsidrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcsidrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcsidrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcsidrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcsidrFieldWtShift = 0
	RegisterCcsidrFieldWtMask  = 0x0
)

// GetWt Indicates support available for Write-Through
func (r *RegisterCcsidrType) GetWt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldWtMask) >> RegisterCcsidrFieldWtShift)
}

// SetWt Indicates support available for Write-Through
func (r *RegisterCcsidrType) SetWt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldWtMask)|(uint32(value)<<RegisterCcsidrFieldWtShift))
}

const (
	RegisterCcsidrFieldLinesizeShift = 0
	RegisterCcsidrFieldLinesizeMask  = 0x7
)

// GetLinesize Indicates the number of words in each cache line
func (r *RegisterCcsidrType) GetLinesize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldLinesizeMask) >> RegisterCcsidrFieldLinesizeShift)
}

// SetLinesize Indicates the number of words in each cache line
func (r *RegisterCcsidrType) SetLinesize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldLinesizeMask)|(uint32(value)<<RegisterCcsidrFieldLinesizeShift))
}

const (
	RegisterCcsidrFieldAssociativityShift = 3
	RegisterCcsidrFieldAssociativityMask  = 0x1ff8
)

// GetAssociativity Indicates associativity
func (r *RegisterCcsidrType) GetAssociativity() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldAssociativityMask) >> RegisterCcsidrFieldAssociativityShift)
}

// SetAssociativity Indicates associativity
func (r *RegisterCcsidrType) SetAssociativity(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldAssociativityMask)|(uint32(value)<<RegisterCcsidrFieldAssociativityShift))
}

const (
	RegisterCcsidrFieldNumsetShift = 13
	RegisterCcsidrFieldNumsetMask  = 0xfffe000
)

// GetNumset Indicates the number of sets
func (r *RegisterCcsidrType) GetNumset() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldNumsetMask) >> RegisterCcsidrFieldNumsetShift)
}

// SetNumset Indicates the number of sets
func (r *RegisterCcsidrType) SetNumset(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldNumsetMask)|(uint32(value)<<RegisterCcsidrFieldNumsetShift))
}

const (
	RegisterCcsidrFieldWaShift = 28
	RegisterCcsidrFieldWaMask  = 0x10000000
)

// GetWa Indicates support available for write allocation
func (r *RegisterCcsidrType) GetWa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldWaMask) != 0
}

// SetWa Indicates support available for write allocation
func (r *RegisterCcsidrType) SetWa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcsidrFieldWaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldWaMask)
	}
}

const (
	RegisterCcsidrFieldRaShift = 29
	RegisterCcsidrFieldRaMask  = 0x20000000
)

// GetRa Indicates support available for read allocation
func (r *RegisterCcsidrType) GetRa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldRaMask) != 0
}

// SetRa Indicates support available for read allocation
func (r *RegisterCcsidrType) SetRa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcsidrFieldRaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldRaMask)
	}
}

const (
	RegisterCcsidrFieldWbShift = 30
	RegisterCcsidrFieldWbMask  = 0x40000000
)

// GetWb Indicates support available for Write-Back
func (r *RegisterCcsidrType) GetWb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldWbMask) != 0
}

// SetWb Indicates support available for Write-Back
func (r *RegisterCcsidrType) SetWb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcsidrFieldWbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldWbMask)
	}
}

// RegisterCsselrType The CSSELR selects the current CCSIDR by specifying the cache level and the type of cache
type RegisterCsselrType uint32

func (r *RegisterCsselrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsselrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsselrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsselrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsselrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsselrFieldIndShift = 0
	RegisterCsselrFieldIndMask  = 0x1
)

// GetInd Instruction not data bit
func (r *RegisterCsselrType) GetInd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsselrFieldIndMask) != 0
}

// SetInd Instruction not data bit
func (r *RegisterCsselrType) SetInd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsselrFieldIndMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsselrFieldIndMask)
	}
}

const (
	RegisterCsselrFieldLevelShift = 1
	RegisterCsselrFieldLevelMask  = 0xe
)

// GetLevel Cache level of required cache
func (r *RegisterCsselrType) GetLevel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCsselrFieldLevelMask) >> RegisterCsselrFieldLevelShift)
}

// SetLevel Cache level of required cache
func (r *RegisterCsselrType) SetLevel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsselrFieldLevelMask)|(uint32(value)<<RegisterCsselrFieldLevelShift))
}

// RegisterCpacrType Specifies the access privileges for coprocessors
type RegisterCpacrType uint32

func (r *RegisterCpacrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCpacrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCpacrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCpacrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCpacrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCpacrFieldCpShift = 0
	RegisterCpacrFieldCpMask  = 0x3
)

// GetCp Coprocessor %s privilege
func (r *RegisterCpacrType) GetCp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpacrFieldCpMask) >> RegisterCpacrFieldCpShift)
}

// SetCp Coprocessor %s privilege
func (r *RegisterCpacrType) SetCp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpacrFieldCpMask)|(uint32(value)<<RegisterCpacrFieldCpShift))
}

type RegisterCpacrFieldCp_10EnumType uint8

const (
	// RegisterCpacrFieldCp10EnumDisabled All accesses to the FP Extension result in NOCP UsageFault.
	RegisterCpacrFieldCp10EnumDisabled RegisterCpacrFieldCp_10EnumType = 0x0

	// RegisterCpacrFieldCp10EnumUnprivileged Unprivileged accesses to the FP Extension result in NOCP UsageFault.
	RegisterCpacrFieldCp10EnumUnprivileged RegisterCpacrFieldCp_10EnumType = 0x1

	// RegisterCpacrFieldCp10EnumFull Full access to the FP Extension.
	RegisterCpacrFieldCp10EnumFull RegisterCpacrFieldCp_10EnumType = 0x3

	RegisterCpacrFieldCp10Shift = 20
	RegisterCpacrFieldCp10Mask  = 0x300000
)

// GetCp10 CP10 Privilege. Defines the access rights for the floating-point functionality.
func (r *RegisterCpacrType) GetCp10() RegisterCpacrFieldCp_10EnumType {
	return RegisterCpacrFieldCp_10EnumType((volatile.LoadUint32((*uint32)(r)) & RegisterCpacrFieldCp10Mask) >> RegisterCpacrFieldCp10Shift)
}

// SetCp10 CP10 Privilege. Defines the access rights for the floating-point functionality.
func (r *RegisterCpacrType) SetCp10(value RegisterCpacrFieldCp_10EnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpacrFieldCp10Mask)|(uint32(value)<<RegisterCpacrFieldCp10Shift))
}

type RegisterCpacrFieldCp_11EnumType uint8

const (
	// RegisterCpacrFieldCp11EnumDisabled All accesses to the FP Extension result in NOCP UsageFault.
	RegisterCpacrFieldCp11EnumDisabled RegisterCpacrFieldCp_11EnumType = 0x0

	// RegisterCpacrFieldCp11EnumUnprivileged Unprivileged accesses to the FP Extension result in NOCP UsageFault.
	RegisterCpacrFieldCp11EnumUnprivileged RegisterCpacrFieldCp_11EnumType = 0x1

	// RegisterCpacrFieldCp11EnumFull Full access to the FP Extension.
	RegisterCpacrFieldCp11EnumFull RegisterCpacrFieldCp_11EnumType = 0x3

	RegisterCpacrFieldCp11Shift = 22
	RegisterCpacrFieldCp11Mask  = 0xc00000
)

// GetCp11 CP11 Privilege. The value in this field is ignored.
func (r *RegisterCpacrType) GetCp11() RegisterCpacrFieldCp_11EnumType {
	return RegisterCpacrFieldCp_11EnumType((volatile.LoadUint32((*uint32)(r)) & RegisterCpacrFieldCp11Mask) >> RegisterCpacrFieldCp11Shift)
}

// SetCp11 CP11 Privilege. The value in this field is ignored.
func (r *RegisterCpacrType) SetCp11(value RegisterCpacrFieldCp_11EnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpacrFieldCp11Mask)|(uint32(value)<<RegisterCpacrFieldCp11Shift))
}

// RegisterNsacrType In an implementation with the Security Extension, the NSACR register defines the Non-secure access permissions for both the Floating-point and MVE and coprocessors CP m, bit[ m], for m = 0-7. If MVE is implemented, this register specifies the Non-secure permissions for MVE
type RegisterNsacrType uint32

func (r *RegisterNsacrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterNsacrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterNsacrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterNsacrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterNsacrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterNsacrFieldCpShift = 0
	RegisterNsacrFieldCpMask  = 0x1
)

// GetCp Enables Non-secure access to this coprocessor
func (r *RegisterNsacrType) GetCp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterNsacrFieldCpMask) != 0
}

// SetCp Enables Non-secure access to this coprocessor
func (r *RegisterNsacrType) SetCp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterNsacrFieldCpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterNsacrFieldCpMask)
	}
}

const (
	RegisterNsacrFieldCp11Shift = 1
	RegisterNsacrFieldCp11Mask  = 0xffe
)

// GetCp11 Enables non-secure access to the Floating-point and MVE. Programming with a different value other than that used for CP10 is unpredictable.
func (r *RegisterNsacrType) GetCp11() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterNsacrFieldCp11Mask) >> RegisterNsacrFieldCp11Shift)
}

// SetCp11 Enables non-secure access to the Floating-point and MVE. Programming with a different value other than that used for CP10 is unpredictable.
func (r *RegisterNsacrType) SetCp11(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterNsacrFieldCp11Mask)|(uint32(value)<<RegisterNsacrFieldCp11Shift))
}

const (
	RegisterNsacrFieldCp10Shift = 10
	RegisterNsacrFieldCp10Mask  = 0x400
)

// GetCp10 Enables non-secure access to the Floating-point and MVE
func (r *RegisterNsacrType) GetCp10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterNsacrFieldCp10Mask) != 0
}

// SetCp10 Enables non-secure access to the Floating-point and MVE
func (r *RegisterNsacrType) SetCp10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterNsacrFieldCp10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterNsacrFieldCp10Mask)
	}
}
