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
	Revidr  registerRevidrType
	Cpuid   registerCpuidType
	Icsr    registerIcsrType
	Vtor    registerVtorType
	Aircr   registerAircrType
	Scr     registerScrType
	Ccr     registerCcrType
	Shpr1   registerShpr1Type
	Shpr2   registerShpr2Type
	Shpr3   registerShpr3Type
	Shcsr   registerShcsrType
	Cfsr    registerCfsrType
	Mmfsr   registerMmfsrType
	Bfsr    registerBfsrType
	Ufsr    registerUfsrType
	Hfsr    registerHfsrType
	_       [4]byte
	Mmfar   registerMmfarType
	Bfar    registerBfarType
	Afsr    registerAfsrType
	Idpfr0  registerIdpfr0Type
	Idpfr1  registerIdpfr1Type
	Iddfr0  registerIddfr0Type
	Idafr0  registerIdafr0Type
	Idmmfr0 registerIdmmfr0Type
	Idmmfr1 registerIdmmfr1Type
	Idmmfr2 registerIdmmfr2Type
	Idmmfr3 registerIdmmfr3Type
	Idisar0 registerIdisar0Type
	Idisar1 registerIdisar1Type
	Idisar2 registerIdisar2Type
	Idisar3 registerIdisar3Type
	Idisar4 registerIdisar4Type
	Idisar5 registerIdisar5Type
	Clidr   registerClidrType
	_       [4]byte
	Ctr     registerCtrType
	Ccsidr  registerCcsidrType
	Csselr  registerCsselrType
	Cpacr   registerCpacrType
	Nsacr   registerNsacrType
}

// registerRevidrType Provides additional implementation-specific minor revision that can be interpreted with the CPUID register
type registerRevidrType uint32

const (
	RegisterRevidrFieldImplementationspecificShift = 0
	RegisterRevidrFieldImplementationspecificMask  = 0xffffffff
)

// GetImplementationspecific Implementation-specific minor revision information that can be interpreted with the CPUID register.
func (r *registerRevidrType) GetImplementationspecific() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRevidrFieldImplementationspecificMask) >> RegisterRevidrFieldImplementationspecificShift)
}

// SetImplementationspecific Implementation-specific minor revision information that can be interpreted with the CPUID register.
func (r *registerRevidrType) SetImplementationspecific(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRevidrFieldImplementationspecificMask)|(uint32(value)<<RegisterRevidrFieldImplementationspecificShift))
}

// registerCpuidType Contains the processor part number, version, and implementation information
type registerCpuidType uint32

const (
	RegisterCpuidFieldRevisionShift = 0
	RegisterCpuidFieldRevisionMask  = 0xf
)

// GetRevision Revision number, the m value in the rnpm product revision identifier
func (r *registerCpuidType) GetRevision() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldRevisionMask) >> RegisterCpuidFieldRevisionShift)
}

// SetRevision Revision number, the m value in the rnpm product revision identifier
func (r *registerCpuidType) SetRevision(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpuidFieldRevisionMask)|(uint32(value)<<RegisterCpuidFieldRevisionShift))
}

const (
	RegisterCpuidFieldPartnoShift = 4
	RegisterCpuidFieldPartnoMask  = 0xfff0
)

// GetPartno Part number of the processor
func (r *registerCpuidType) GetPartno() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldPartnoMask) >> RegisterCpuidFieldPartnoShift)
}

// SetPartno Part number of the processor
func (r *registerCpuidType) SetPartno(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpuidFieldPartnoMask)|(uint32(value)<<RegisterCpuidFieldPartnoShift))
}

const (
	RegisterCpuidFieldArchitectureShift = 16
	RegisterCpuidFieldArchitectureMask  = 0xf0000
)

// GetArchitecture Reads as 0xF
func (r *registerCpuidType) GetArchitecture() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldArchitectureMask) >> RegisterCpuidFieldArchitectureShift)
}

// SetArchitecture Reads as 0xF
func (r *registerCpuidType) SetArchitecture(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpuidFieldArchitectureMask)|(uint32(value)<<RegisterCpuidFieldArchitectureShift))
}

const (
	RegisterCpuidFieldVariantShift = 20
	RegisterCpuidFieldVariantMask  = 0xf00000
)

// GetVariant Variant number, the n value in the rnpm product revision identifier
func (r *registerCpuidType) GetVariant() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldVariantMask) >> RegisterCpuidFieldVariantShift)
}

// SetVariant Variant number, the n value in the rnpm product revision identifier
func (r *registerCpuidType) SetVariant(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpuidFieldVariantMask)|(uint32(value)<<RegisterCpuidFieldVariantShift))
}

const (
	RegisterCpuidFieldImplementerShift = 24
	RegisterCpuidFieldImplementerMask  = 0xff000000
)

// GetImplementer Implementer code
func (r *registerCpuidType) GetImplementer() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldImplementerMask) >> RegisterCpuidFieldImplementerShift)
}

// SetImplementer Implementer code
func (r *registerCpuidType) SetImplementer(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpuidFieldImplementerMask)|(uint32(value)<<RegisterCpuidFieldImplementerShift))
}

// registerIcsrType Provides a set-pending bit for the non-maskable interrupt exception, and set-pending and clear-pending bits for the PendSV and SysTick exceptions
type registerIcsrType uint32

const (
	RegisterIcsrFieldVectactiveShift = 0
	RegisterIcsrFieldVectactiveMask  = 0x1ff
)

// GetVectactive Contains the active exception number
func (r *registerIcsrType) GetVectactive() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldVectactiveMask) >> RegisterIcsrFieldVectactiveShift)
}

// SetVectactive Contains the active exception number
func (r *registerIcsrType) SetVectactive(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldVectactiveMask)|(uint32(value)<<RegisterIcsrFieldVectactiveShift))
}

const (
	RegisterIcsrFieldRettobaseShift = 11
	RegisterIcsrFieldRettobaseMask  = 0x800
)

// GetRettobase Indicates whether there are pre-empted active exceptions
func (r *registerIcsrType) GetRettobase() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldRettobaseMask) != 0
}

// SetRettobase Indicates whether there are pre-empted active exceptions
func (r *registerIcsrType) SetRettobase(value bool) {
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
func (r *registerIcsrType) GetVectpending() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldVectpendingMask) >> RegisterIcsrFieldVectpendingShift)
}

// SetVectpending Indicates the exception number of the highest priority pending enabled exception
func (r *registerIcsrType) SetVectpending(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldVectpendingMask)|(uint32(value)<<RegisterIcsrFieldVectpendingShift))
}

const (
	RegisterIcsrFieldIsrpendingShift = 22
	RegisterIcsrFieldIsrpendingMask  = 0x400000
)

// GetIsrpending Interrupt pending flag, excluding NMI and Faults
func (r *registerIcsrType) GetIsrpending() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldIsrpendingMask) != 0
}

// SetIsrpending Interrupt pending flag, excluding NMI and Faults
func (r *registerIcsrType) SetIsrpending(value bool) {
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
func (r *registerIcsrType) GetIsrpreempt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldIsrpreemptMask) != 0
}

// SetIsrpreempt Indicates whether a pending exception is handled on exit from debug state.
func (r *registerIcsrType) SetIsrpreempt(value bool) {
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
func (r *registerIcsrType) GetPendstclr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendstclrMask) != 0
}

// SetPendstclr SysTick exception clear-pending bit
func (r *registerIcsrType) SetPendstclr(value bool) {
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
func (r *registerIcsrType) GetPendstset() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendstsetMask) != 0
}

// SetPendstset SysTick exception set-pending bit
func (r *registerIcsrType) SetPendstset(value bool) {
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
func (r *registerIcsrType) GetPendsvclr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendsvclrMask) != 0
}

// SetPendsvclr PendSV clear-pending bit (write-only)
func (r *registerIcsrType) SetPendsvclr(value bool) {
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
func (r *registerIcsrType) GetPendsvset() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendsvsetMask) != 0
}

// SetPendsvset PendSV set-pending bit
func (r *registerIcsrType) SetPendsvset(value bool) {
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
func (r *registerIcsrType) GetPendnmiclr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendnmiclrMask) != 0
}

// SetPendnmiclr Pend NMI clear bit (write-only)
func (r *registerIcsrType) SetPendnmiclr(value bool) {
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
func (r *registerIcsrType) GetPendnmiset() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendnmisetMask) != 0
}

// SetPendnmiset On writes, makes the NMI exception pending
func (r *registerIcsrType) SetPendnmiset(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldPendnmisetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldPendnmisetMask)
	}
}

// registerVtorType Indicates the offset of the vector table base address from memory address 0x00000000
type registerVtorType uint32

const (
	RegisterVtorFieldTbloffShift = 7
	RegisterVtorFieldTbloffMask  = 0xffffff80
)

// GetTbloff Bits[31:7] of the vector table address
func (r *registerVtorType) GetTbloff() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterVtorFieldTbloffMask) >> RegisterVtorFieldTbloffShift)
}

// SetTbloff Bits[31:7] of the vector table address
func (r *registerVtorType) SetTbloff(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterVtorFieldTbloffMask)|(uint32(value)<<RegisterVtorFieldTbloffShift))
}

// registerAircrType Sets or returns interrupt control data
type registerAircrType uint32

const (
	RegisterAircrFieldVectclractiveShift = 1
	RegisterAircrFieldVectclractiveMask  = 0x2
)

// GetVectclractive Reserved for Debug use. This bit reads as 0. When writing to the register you must write 0 to this bit, otherwise behavior is UNPREDICTABLE
func (r *registerAircrType) GetVectclractive() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldVectclractiveMask) != 0
}

// SetVectclractive Reserved for Debug use. This bit reads as 0. When writing to the register you must write 0 to this bit, otherwise behavior is UNPREDICTABLE
func (r *registerAircrType) SetVectclractive(value bool) {
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
func (r *registerAircrType) GetSysresetreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldSysresetreqMask) != 0
}

// SetSysresetreq System Reset Request
func (r *registerAircrType) SetSysresetreq(value bool) {
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
func (r *registerAircrType) GetSysresetreqs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldSysresetreqsMask) != 0
}

// SetSysresetreqs System reset request, Secure state only
func (r *registerAircrType) SetSysresetreqs(value bool) {
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
func (r *registerAircrType) GetPrigroup() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldPrigroupMask) >> RegisterAircrFieldPrigroupShift)
}

// SetPrigroup Interrupt priority grouping
func (r *registerAircrType) SetPrigroup(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldPrigroupMask)|(uint32(value)<<RegisterAircrFieldPrigroupShift))
}

const (
	RegisterAircrFieldBfhfnminsShift = 13
	RegisterAircrFieldBfhfnminsMask  = 0x2000
)

// GetBfhfnmins BusFault, HardFault, and NMI Non-secure enable
func (r *registerAircrType) GetBfhfnmins() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldBfhfnminsMask) != 0
}

// SetBfhfnmins BusFault, HardFault, and NMI Non-secure enable
func (r *registerAircrType) SetBfhfnmins(value bool) {
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
func (r *registerAircrType) GetPris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldPrisMask) != 0
}

// SetPris Prioritize Secure exceptions
func (r *registerAircrType) SetPris(value bool) {
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
func (r *registerAircrType) GetEndianness() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldEndiannessMask) != 0
}

// SetEndianness Data endianness
func (r *registerAircrType) SetEndianness(value bool) {
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
func (r *registerAircrType) GetVectkey() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldVectkeyMask) >> RegisterAircrFieldVectkeyShift)
}

// SetVectkey Register key. Reads as 0x0FA05, On writes, write 0x5FA to VECTKEY, otherwise the write is ignored.
func (r *registerAircrType) SetVectkey(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldVectkeyMask)|(uint32(value)<<RegisterAircrFieldVectkeyShift))
}

// registerScrType Controls features of entry to and exit from low-power state
type registerScrType uint32

const (
	RegisterScrFieldSleeponexitShift = 1
	RegisterScrFieldSleeponexitMask  = 0x2
)

// GetSleeponexit Indicates sleep-on-exit when returning from Handler mode to Thread mode
func (r *registerScrType) GetSleeponexit() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScrFieldSleeponexitMask) != 0
}

// SetSleeponexit Indicates sleep-on-exit when returning from Handler mode to Thread mode
func (r *registerScrType) SetSleeponexit(value bool) {
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
func (r *registerScrType) GetSleepdeep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScrFieldSleepdeepMask) != 0
}

// SetSleepdeep Indicates whether the processor uses sleep or deep sleep as its low-power mode
func (r *registerScrType) SetSleepdeep(value bool) {
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
func (r *registerScrType) GetSleepdeeps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScrFieldSleepdeepsMask) != 0
}

// SetSleepdeeps Controls whether the SLEEPDEEP bit is only accessible from the Secure state
func (r *registerScrType) SetSleepdeeps(value bool) {
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
func (r *registerScrType) GetSevonpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScrFieldSevonpendMask) != 0
}

// SetSevonpend Determines whether an interrupt assigned to the same Security state as the SEVONPEND bit transitioning from inactive state to pending state generates a wakeup event
func (r *registerScrType) SetSevonpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScrFieldSevonpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScrFieldSevonpendMask)
	}
}

// registerCcrType indicates some aspects of the behavior of the processor
type registerCcrType uint32

const (
	RegisterCcrFieldUsersetmpendShift = 1
	RegisterCcrFieldUsersetmpendMask  = 0x2
)

// GetUsersetmpend User set main pending. Determines whether unprivileged accesses are permitted to pend interrupts from the STIR.
func (r *registerCcrType) GetUsersetmpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldUsersetmpendMask) != 0
}

// SetUsersetmpend User set main pending. Determines whether unprivileged accesses are permitted to pend interrupts from the STIR.
func (r *registerCcrType) SetUsersetmpend(value bool) {
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
func (r *registerCcrType) GetUnaligntrp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldUnaligntrpMask) != 0
}

// SetUnaligntrp Controls the trapping of unaligned word or halfword accesses
func (r *registerCcrType) SetUnaligntrp(value bool) {
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
func (r *registerCcrType) GetDiv0trp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDiv0trpMask) != 0
}

// SetDiv0trp Divide by zero trap. Controls the generation of a DIVBYZERO UsageFault when attempting to perform integer division by zero.
func (r *registerCcrType) SetDiv0trp(value bool) {
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
func (r *registerCcrType) GetBfhfnmign() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldBfhfnmignMask) != 0
}

// SetBfhfnmign Determines the effect of precise bus faults on handlers running at a requested priority less than 0
func (r *registerCcrType) SetBfhfnmign(value bool) {
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
func (r *registerCcrType) GetStkofhfnmign() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldStkofhfnmignMask) != 0
}

// SetStkofhfnmign Controls the effect of a stack limit violation while executing at a requested priority less than 0
func (r *registerCcrType) SetStkofhfnmign(value bool) {
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
func (r *registerCcrType) GetDc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDcMask) != 0
}

// SetDc Data cache enable
func (r *registerCcrType) SetDc(value bool) {
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
func (r *registerCcrType) GetIc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldIcMask) != 0
}

// SetIc Instruction cache enable
func (r *registerCcrType) SetIc(value bool) {
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
func (r *registerCcrType) GetLob() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldLobMask) != 0
}

// SetLob Loop and branch info cache enable
func (r *registerCcrType) SetLob(value bool) {
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
func (r *registerCcrType) GetTrd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldTrdMask) != 0
}

// SetTrd Thread reentrancy disabled
func (r *registerCcrType) SetTrd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldTrdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldTrdMask)
	}
}

// registerShpr1Type Sets or returns priority for system handlers 4-7
type registerShpr1Type uint32

const (
	RegisterShpr1FieldPri4Shift = 0
	RegisterShpr1FieldPri4Mask  = 0xff
)

// GetPri4 Priority of system handler 4, MemManage
func (r *registerShpr1Type) GetPri4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr1FieldPri4Mask) >> RegisterShpr1FieldPri4Shift)
}

// SetPri4 Priority of system handler 4, MemManage
func (r *registerShpr1Type) SetPri4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr1FieldPri4Mask)|(uint32(value)<<RegisterShpr1FieldPri4Shift))
}

const (
	RegisterShpr1FieldPri5Shift = 8
	RegisterShpr1FieldPri5Mask  = 0xff00
)

// GetPri5 Priority of system handler 5, BusFault
func (r *registerShpr1Type) GetPri5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr1FieldPri5Mask) >> RegisterShpr1FieldPri5Shift)
}

// SetPri5 Priority of system handler 5, BusFault
func (r *registerShpr1Type) SetPri5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr1FieldPri5Mask)|(uint32(value)<<RegisterShpr1FieldPri5Shift))
}

const (
	RegisterShpr1FieldPri6Shift = 16
	RegisterShpr1FieldPri6Mask  = 0xff0000
)

// GetPri6 Priority of system handler 6, UsageFault
func (r *registerShpr1Type) GetPri6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr1FieldPri6Mask) >> RegisterShpr1FieldPri6Shift)
}

// SetPri6 Priority of system handler 6, UsageFault
func (r *registerShpr1Type) SetPri6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr1FieldPri6Mask)|(uint32(value)<<RegisterShpr1FieldPri6Shift))
}

const (
	RegisterShpr1FieldPri7Shift = 24
	RegisterShpr1FieldPri7Mask  = 0xff000000
)

// GetPri7 Priority of system handler 7, SecureFault
func (r *registerShpr1Type) GetPri7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr1FieldPri7Mask) >> RegisterShpr1FieldPri7Shift)
}

// SetPri7 Priority of system handler 7, SecureFault
func (r *registerShpr1Type) SetPri7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr1FieldPri7Mask)|(uint32(value)<<RegisterShpr1FieldPri7Shift))
}

// registerShpr2Type Sets or returns priority for system handler 8-11
type registerShpr2Type uint32

const (
	RegisterShpr2FieldPri11Shift = 24
	RegisterShpr2FieldPri11Mask  = 0xff000000
)

// GetPri11 Priority of system handler 11, SVCall
func (r *registerShpr2Type) GetPri11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr2FieldPri11Mask) >> RegisterShpr2FieldPri11Shift)
}

// SetPri11 Priority of system handler 11, SVCall
func (r *registerShpr2Type) SetPri11(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr2FieldPri11Mask)|(uint32(value)<<RegisterShpr2FieldPri11Shift))
}

// registerShpr3Type Sets or returns priority for system handlers 12-15
type registerShpr3Type uint32

const (
	RegisterShpr3FieldPri14Shift = 16
	RegisterShpr3FieldPri14Mask  = 0xff0000
)

// GetPri14 Priority of system handler 14, PendSV
func (r *registerShpr3Type) GetPri14() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr3FieldPri14Mask) >> RegisterShpr3FieldPri14Shift)
}

// SetPri14 Priority of system handler 14, PendSV
func (r *registerShpr3Type) SetPri14(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr3FieldPri14Mask)|(uint32(value)<<RegisterShpr3FieldPri14Shift))
}

const (
	RegisterShpr3FieldPri15Shift = 24
	RegisterShpr3FieldPri15Mask  = 0xff000000
)

// GetPri15 Priority of system handler 15, SysTick
func (r *registerShpr3Type) GetPri15() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr3FieldPri15Mask) >> RegisterShpr3FieldPri15Shift)
}

// SetPri15 Priority of system handler 15, SysTick
func (r *registerShpr3Type) SetPri15(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr3FieldPri15Mask)|(uint32(value)<<RegisterShpr3FieldPri15Shift))
}

// registerShcsrType Enables the system handlers and indicates the pending status of the BusFault, MemManage fault, and SVC exceptions, and indicates the active status of the system handlers
type registerShcsrType uint32

const (
	RegisterShcsrFieldMemfaultactShift = 0
	RegisterShcsrFieldMemfaultactMask  = 0x1
)

// GetMemfaultact MemManage exception active
func (r *registerShcsrType) GetMemfaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldMemfaultactMask) != 0
}

// SetMemfaultact MemManage exception active
func (r *registerShcsrType) SetMemfaultact(value bool) {
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
func (r *registerShcsrType) GetBusfaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldBusfaultactMask) != 0
}

// SetBusfaultact BusFault exception active
func (r *registerShcsrType) SetBusfaultact(value bool) {
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
func (r *registerShcsrType) GetHardfaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldHardfaultactMask) != 0
}

// SetHardfaultact HardFault exception active
func (r *registerShcsrType) SetHardfaultact(value bool) {
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
func (r *registerShcsrType) GetUsgfaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldUsgfaultactMask) != 0
}

// SetUsgfaultact UsageFault exception active
func (r *registerShcsrType) SetUsgfaultact(value bool) {
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
func (r *registerShcsrType) GetSecurefaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldSecurefaultactMask) != 0
}

// SetSecurefaultact SecureFault exception active state
func (r *registerShcsrType) SetSecurefaultact(value bool) {
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
func (r *registerShcsrType) GetNmiact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldNmiactMask) != 0
}

// SetNmiact NMI exception active
func (r *registerShcsrType) SetNmiact(value bool) {
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
func (r *registerShcsrType) GetSvcallact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldSvcallactMask) != 0
}

// SetSvcallact SVCall active
func (r *registerShcsrType) SetSvcallact(value bool) {
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
func (r *registerShcsrType) GetMonitoract() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldMonitoractMask) != 0
}

// SetMonitoract Debug monitor active
func (r *registerShcsrType) SetMonitoract(value bool) {
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
func (r *registerShcsrType) GetPendsvact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldPendsvactMask) != 0
}

// SetPendsvact PendSV exception active
func (r *registerShcsrType) SetPendsvact(value bool) {
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
func (r *registerShcsrType) GetSystickact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldSystickactMask) != 0
}

// SetSystickact SysTick exception active
func (r *registerShcsrType) SetSystickact(value bool) {
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
func (r *registerShcsrType) GetUsgfaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldUsgfaultpendedMask) != 0
}

// SetUsgfaultpended UsageFault exception pending
func (r *registerShcsrType) SetUsgfaultpended(value bool) {
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
func (r *registerShcsrType) GetMemfaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldMemfaultpendedMask) != 0
}

// SetMemfaultpended MemManage exception pending
func (r *registerShcsrType) SetMemfaultpended(value bool) {
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
func (r *registerShcsrType) GetBusfaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldBusfaultpendedMask) != 0
}

// SetBusfaultpended BusFault pending
func (r *registerShcsrType) SetBusfaultpended(value bool) {
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
func (r *registerShcsrType) GetSvcallpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldSvcallpendedMask) != 0
}

// SetSvcallpended SVCall pending
func (r *registerShcsrType) SetSvcallpended(value bool) {
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
func (r *registerShcsrType) GetMemfaultena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldMemfaultenaMask) != 0
}

// SetMemfaultena MemManage enable
func (r *registerShcsrType) SetMemfaultena(value bool) {
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
func (r *registerShcsrType) GetBusfaultena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldBusfaultenaMask) != 0
}

// SetBusfaultena BusFault enable
func (r *registerShcsrType) SetBusfaultena(value bool) {
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
func (r *registerShcsrType) GetUsgfaultena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldUsgfaultenaMask) != 0
}

// SetUsgfaultena UsageFault enable
func (r *registerShcsrType) SetUsgfaultena(value bool) {
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
func (r *registerShcsrType) GetSecurefaultena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldSecurefaultenaMask) != 0
}

// SetSecurefaultena RES0
func (r *registerShcsrType) SetSecurefaultena(value bool) {
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
func (r *registerShcsrType) GetSecurefaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldSecurefaultpendedMask) != 0
}

// SetSecurefaultpended SecureFault exception pended state bit, set to 1 to allow exception modification
func (r *registerShcsrType) SetSecurefaultpended(value bool) {
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
func (r *registerShcsrType) GetHardfaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcsrFieldHardfaultpendedMask) != 0
}

// SetHardfaultpended HardFault exception pended state bit, set to 1 to allow exception modification
func (r *registerShcsrType) SetHardfaultpended(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcsrFieldHardfaultpendedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcsrFieldHardfaultpendedMask)
	}
}

// registerCfsrType Indicates the cause of a MemManage fault, BusFault, or UsageFault
type registerCfsrType uint32

const (
	RegisterCfsrFieldMemmanageShift = 0
	RegisterCfsrFieldMemmanageMask  = 0xff
)

// GetMemmanage Provides information on MemManage exceptions
func (r *registerCfsrType) GetMemmanage() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfsrFieldMemmanageMask) >> RegisterCfsrFieldMemmanageShift)
}

// SetMemmanage Provides information on MemManage exceptions
func (r *registerCfsrType) SetMemmanage(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfsrFieldMemmanageMask)|(uint32(value)<<RegisterCfsrFieldMemmanageShift))
}

const (
	RegisterCfsrFieldBusfaultShift = 8
	RegisterCfsrFieldBusfaultMask  = 0xff00
)

// GetBusfault Provides information on BusFault exceptions
func (r *registerCfsrType) GetBusfault() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfsrFieldBusfaultMask) >> RegisterCfsrFieldBusfaultShift)
}

// SetBusfault Provides information on BusFault exceptions
func (r *registerCfsrType) SetBusfault(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfsrFieldBusfaultMask)|(uint32(value)<<RegisterCfsrFieldBusfaultShift))
}

const (
	RegisterCfsrFieldUsagefaultShift = 16
	RegisterCfsrFieldUsagefaultMask  = 0xffff0000
)

// GetUsagefault Provides information on UsageFault exceptions
func (r *registerCfsrType) GetUsagefault() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCfsrFieldUsagefaultMask) >> RegisterCfsrFieldUsagefaultShift)
}

// SetUsagefault Provides information on UsageFault exceptions
func (r *registerCfsrType) SetUsagefault(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfsrFieldUsagefaultMask)|(uint32(value)<<RegisterCfsrFieldUsagefaultShift))
}

// registerMmfsrType Indicates the cause of memory access faults
type registerMmfsrType uint32

const (
	RegisterMmfsrFieldIaccviolShift = 0
	RegisterMmfsrFieldIaccviolMask  = 0x1
)

// GetIaccviol Instruction access violation flag
func (r *registerMmfsrType) GetIaccviol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmfsrFieldIaccviolMask) != 0
}

// SetIaccviol Instruction access violation flag
func (r *registerMmfsrType) SetIaccviol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmfsrFieldIaccviolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmfsrFieldIaccviolMask)
	}
}

const (
	RegisterMmfsrFieldDaccviolShift = 1
	RegisterMmfsrFieldDaccviolMask  = 0x2
)

// GetDaccviol Data access violation flag
func (r *registerMmfsrType) GetDaccviol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmfsrFieldDaccviolMask) != 0
}

// SetDaccviol Data access violation flag
func (r *registerMmfsrType) SetDaccviol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmfsrFieldDaccviolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmfsrFieldDaccviolMask)
	}
}

const (
	RegisterMmfsrFieldMunstkerrShift = 3
	RegisterMmfsrFieldMunstkerrMask  = 0x8
)

// GetMunstkerr MemManage fault on unstacking for a return from exception
func (r *registerMmfsrType) GetMunstkerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmfsrFieldMunstkerrMask) != 0
}

// SetMunstkerr MemManage fault on unstacking for a return from exception
func (r *registerMmfsrType) SetMunstkerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmfsrFieldMunstkerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmfsrFieldMunstkerrMask)
	}
}

const (
	RegisterMmfsrFieldMstkerrShift = 4
	RegisterMmfsrFieldMstkerrMask  = 0x10
)

// GetMstkerr MemManage fault on stacking for exception entry
func (r *registerMmfsrType) GetMstkerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmfsrFieldMstkerrMask) != 0
}

// SetMstkerr MemManage fault on stacking for exception entry
func (r *registerMmfsrType) SetMstkerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmfsrFieldMstkerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmfsrFieldMstkerrMask)
	}
}

const (
	RegisterMmfsrFieldMlsperrShift = 5
	RegisterMmfsrFieldMlsperrMask  = 0x20
)

// GetMlsperr MemManage lazy state preservation error flag
func (r *registerMmfsrType) GetMlsperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmfsrFieldMlsperrMask) != 0
}

// SetMlsperr MemManage lazy state preservation error flag
func (r *registerMmfsrType) SetMlsperr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmfsrFieldMlsperrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmfsrFieldMlsperrMask)
	}
}

const (
	RegisterMmfsrFieldMmarvalidShift = 7
	RegisterMmfsrFieldMmarvalidMask  = 0x80
)

// GetMmarvalid MemManage Fault Address Register (MMFAR) valid flag
func (r *registerMmfsrType) GetMmarvalid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmfsrFieldMmarvalidMask) != 0
}

// SetMmarvalid MemManage Fault Address Register (MMFAR) valid flag
func (r *registerMmfsrType) SetMmarvalid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmfsrFieldMmarvalidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmfsrFieldMmarvalidMask)
	}
}

// registerBfsrType Indicate the cause of a bus access fault
type registerBfsrType uint32

const (
	RegisterBfsrFieldIbuserrShift = 0
	RegisterBfsrFieldIbuserrMask  = 0x1
)

// GetIbuserr Instruction bus error
func (r *registerBfsrType) GetIbuserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBfsrFieldIbuserrMask) != 0
}

// SetIbuserr Instruction bus error
func (r *registerBfsrType) SetIbuserr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBfsrFieldIbuserrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBfsrFieldIbuserrMask)
	}
}

const (
	RegisterBfsrFieldPreciserrShift = 1
	RegisterBfsrFieldPreciserrMask  = 0x2
)

// GetPreciserr Precise data bus error
func (r *registerBfsrType) GetPreciserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBfsrFieldPreciserrMask) != 0
}

// SetPreciserr Precise data bus error
func (r *registerBfsrType) SetPreciserr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBfsrFieldPreciserrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBfsrFieldPreciserrMask)
	}
}

const (
	RegisterBfsrFieldUnstkerrShift = 3
	RegisterBfsrFieldUnstkerrMask  = 0x8
)

// GetUnstkerr BusFault on unstacking for a return from exception
func (r *registerBfsrType) GetUnstkerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBfsrFieldUnstkerrMask) != 0
}

// SetUnstkerr BusFault on unstacking for a return from exception
func (r *registerBfsrType) SetUnstkerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBfsrFieldUnstkerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBfsrFieldUnstkerrMask)
	}
}

const (
	RegisterBfsrFieldStkerrShift = 4
	RegisterBfsrFieldStkerrMask  = 0x10
)

// GetStkerr BusFault on stacking for exception entry
func (r *registerBfsrType) GetStkerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBfsrFieldStkerrMask) != 0
}

// SetStkerr BusFault on stacking for exception entry
func (r *registerBfsrType) SetStkerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBfsrFieldStkerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBfsrFieldStkerrMask)
	}
}

const (
	RegisterBfsrFieldLsperrShift = 5
	RegisterBfsrFieldLsperrMask  = 0x20
)

// GetLsperr Lazy state preservation error
func (r *registerBfsrType) GetLsperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBfsrFieldLsperrMask) != 0
}

// SetLsperr Lazy state preservation error
func (r *registerBfsrType) SetLsperr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBfsrFieldLsperrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBfsrFieldLsperrMask)
	}
}

const (
	RegisterBfsrFieldBfarvalidShift = 7
	RegisterBfsrFieldBfarvalidMask  = 0x80
)

// GetBfarvalid BusFault Address Register (BFAR) valid flag.
func (r *registerBfsrType) GetBfarvalid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBfsrFieldBfarvalidMask) != 0
}

// SetBfarvalid BusFault Address Register (BFAR) valid flag.
func (r *registerBfsrType) SetBfarvalid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBfsrFieldBfarvalidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBfsrFieldBfarvalidMask)
	}
}

// registerUfsrType Indicates the cause of a UsageFault
type registerUfsrType uint32

const (
	RegisterUfsrFieldUndefinstrShift = 0
	RegisterUfsrFieldUndefinstrMask  = 0x1
)

// GetUndefinstr Undefined instruction flag. Sticky flag indicating whether an undefined instruction error has occurred.
func (r *registerUfsrType) GetUndefinstr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUfsrFieldUndefinstrMask) != 0
}

// SetUndefinstr Undefined instruction flag. Sticky flag indicating whether an undefined instruction error has occurred.
func (r *registerUfsrType) SetUndefinstr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUfsrFieldUndefinstrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUfsrFieldUndefinstrMask)
	}
}

const (
	RegisterUfsrFieldInvstateShift = 1
	RegisterUfsrFieldInvstateMask  = 0x2
)

// GetInvstate Invalid state flag. Sticky flag indicating whether an EPSR.T or EPSR.IT validity error has occurred.
func (r *registerUfsrType) GetInvstate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUfsrFieldInvstateMask) != 0
}

// SetInvstate Invalid state flag. Sticky flag indicating whether an EPSR.T or EPSR.IT validity error has occurred.
func (r *registerUfsrType) SetInvstate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUfsrFieldInvstateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUfsrFieldInvstateMask)
	}
}

const (
	RegisterUfsrFieldInvpcShift = 2
	RegisterUfsrFieldInvpcMask  = 0x4
)

// GetInvpc Invalid PC flag. Sticky flag indicating whether an integrity check error has occurred.
func (r *registerUfsrType) GetInvpc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUfsrFieldInvpcMask) != 0
}

// SetInvpc Invalid PC flag. Sticky flag indicating whether an integrity check error has occurred.
func (r *registerUfsrType) SetInvpc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUfsrFieldInvpcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUfsrFieldInvpcMask)
	}
}

const (
	RegisterUfsrFieldNocpShift = 3
	RegisterUfsrFieldNocpMask  = 0x8
)

// GetNocp No coprocessor flag. Sticky flag indicating whether a coprocessor disabled or not present error has occurred.
func (r *registerUfsrType) GetNocp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUfsrFieldNocpMask) != 0
}

// SetNocp No coprocessor flag. Sticky flag indicating whether a coprocessor disabled or not present error has occurred.
func (r *registerUfsrType) SetNocp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUfsrFieldNocpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUfsrFieldNocpMask)
	}
}

const (
	RegisterUfsrFieldStkofShift = 4
	RegisterUfsrFieldStkofMask  = 0x10
)

// GetStkof Stack overflow flag. Sticky flag indicating whether a stack overflow error has occurred.
func (r *registerUfsrType) GetStkof() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUfsrFieldStkofMask) != 0
}

// SetStkof Stack overflow flag. Sticky flag indicating whether a stack overflow error has occurred.
func (r *registerUfsrType) SetStkof(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUfsrFieldStkofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUfsrFieldStkofMask)
	}
}

const (
	RegisterUfsrFieldUnalignedShift = 8
	RegisterUfsrFieldUnalignedMask  = 0x100
)

// GetUnaligned Unaligned access flag. Sticky flag indicating whether an unaligned access error has occurred.
func (r *registerUfsrType) GetUnaligned() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUfsrFieldUnalignedMask) != 0
}

// SetUnaligned Unaligned access flag. Sticky flag indicating whether an unaligned access error has occurred.
func (r *registerUfsrType) SetUnaligned(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUfsrFieldUnalignedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUfsrFieldUnalignedMask)
	}
}

const (
	RegisterUfsrFieldDivbyzeroShift = 9
	RegisterUfsrFieldDivbyzeroMask  = 0x200
)

// GetDivbyzero Divide by zero flag. Sticky flag indicating whether an integer division by zero error has occurred.
func (r *registerUfsrType) GetDivbyzero() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterUfsrFieldDivbyzeroMask) != 0
}

// SetDivbyzero Divide by zero flag. Sticky flag indicating whether an integer division by zero error has occurred.
func (r *registerUfsrType) SetDivbyzero(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterUfsrFieldDivbyzeroMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterUfsrFieldDivbyzeroMask)
	}
}

// registerHfsrType Gives information about events that activate the HardFault handler
type registerHfsrType uint32

const (
	RegisterHfsrFieldVecttblShift = 1
	RegisterHfsrFieldVecttblMask  = 0x2
)

// GetVecttbl Indicates a HardFault on a vector table read during exception processing
func (r *registerHfsrType) GetVecttbl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHfsrFieldVecttblMask) != 0
}

// SetVecttbl Indicates a HardFault on a vector table read during exception processing
func (r *registerHfsrType) SetVecttbl(value bool) {
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
func (r *registerHfsrType) GetForced() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHfsrFieldForcedMask) != 0
}

// SetForced Indicates a forced HardFault, generated by escalation of a fault with configurable priority that cannot be handled, either because of priority or because it is disabled
func (r *registerHfsrType) SetForced(value bool) {
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
func (r *registerHfsrType) GetDebugevt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHfsrFieldDebugevtMask) != 0
}

// SetDebugevt Reserved for Debug use. When writing to the register you must write 1 to this bit, otherwise behavior is unpredictable.
func (r *registerHfsrType) SetDebugevt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHfsrFieldDebugevtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHfsrFieldDebugevtMask)
	}
}

// registerMmfarType Shows the address of the memory location that caused a Memory Protection Unit (MPU) fault
type registerMmfarType uint32

// registerBfarType Shows the address associated with a precise data access BusFault
type registerBfarType uint32

// registerAfsrType provides fault status information
type registerAfsrType uint32

const (
	RegisterAfsrFieldIitcmShift = 0
	RegisterAfsrFieldIitcmMask  = 0x1
)

// GetIitcm Imprecise fault on ITCM interface
func (r *registerAfsrType) GetIitcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldIitcmMask) != 0
}

// SetIitcm Imprecise fault on ITCM interface
func (r *registerAfsrType) SetIitcm(value bool) {
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
func (r *registerAfsrType) GetIdtcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldIdtcmMask) != 0
}

// SetIdtcm Imprecise fault on DTCM interface
func (r *registerAfsrType) SetIdtcm(value bool) {
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
func (r *registerAfsrType) GetIpahb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldIpahbMask) != 0
}

// SetIpahb Imprecise fault on P-AHB interface
func (r *registerAfsrType) SetIpahb(value bool) {
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
func (r *registerAfsrType) GetImaxi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldImaxiMask) != 0
}

// SetImaxi Imprecise fault on M-AXI interface
func (r *registerAfsrType) SetImaxi(value bool) {
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
func (r *registerAfsrType) GetIeppb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldIeppbMask) != 0
}

// SetIeppb Imprecise fault on EPPB interface
func (r *registerAfsrType) SetIeppb(value bool) {
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
func (r *registerAfsrType) GetImaxitype() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldImaxitypeMask) != 0
}

// SetImaxitype AXI response that caused the imprecise fault. Only valid when AFSR.IMAXI is 1.
func (r *registerAfsrType) SetImaxitype(value bool) {
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
func (r *registerAfsrType) GetIecc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldIeccMask) != 0
}

// SetIecc Imprecise fault caused by uncorrectable ECC error
func (r *registerAfsrType) SetIecc(value bool) {
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
func (r *registerAfsrType) GetIpoison() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldIpoisonMask) != 0
}

// SetIpoison Imprecise BusFault because of RPOISON
func (r *registerAfsrType) SetIpoison(value bool) {
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
func (r *registerAfsrType) GetPitcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPitcmMask) != 0
}

// SetPitcm Precise fault on ITCM interface
func (r *registerAfsrType) SetPitcm(value bool) {
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
func (r *registerAfsrType) GetPdtcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPdtcmMask) != 0
}

// SetPdtcm Precise fault on DTCM interface
func (r *registerAfsrType) SetPdtcm(value bool) {
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
func (r *registerAfsrType) GetPpahb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPpahbMask) != 0
}

// SetPpahb Precise fault on Peripheral AHB (P-AHB) interface
func (r *registerAfsrType) SetPpahb(value bool) {
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
func (r *registerAfsrType) GetPmaxi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPmaxiMask) != 0
}

// SetPmaxi Precise fault on M-AXI interface
func (r *registerAfsrType) SetPmaxi(value bool) {
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
func (r *registerAfsrType) GetPeppb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPeppbMask) != 0
}

// SetPeppb Precise fault on External Private Peripheral Bus (EPPB) interface
func (r *registerAfsrType) SetPeppb(value bool) {
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
func (r *registerAfsrType) GetPippb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPippbMask) != 0
}

// SetPippb Precise fault on Internal Private Peripheral Bus (EPPB) interface
func (r *registerAfsrType) SetPippb(value bool) {
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
func (r *registerAfsrType) GetPmaxitype() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPmaxitypeMask) != 0
}

// SetPmaxitype AXI response that caused the precise fault. Only valid when AFSR.PMAXI is 1.
func (r *registerAfsrType) SetPmaxitype(value bool) {
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
func (r *registerAfsrType) GetPecc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPeccMask) != 0
}

// SetPecc Precise fault caused by uncorrectable ECC error
func (r *registerAfsrType) SetPecc(value bool) {
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
func (r *registerAfsrType) GetPtgu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPtguMask) != 0
}

// SetPtgu Precise fault caused by TGU security violation
func (r *registerAfsrType) SetPtgu(value bool) {
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
func (r *registerAfsrType) GetPpoison() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldPpoisonMask) != 0
}

// SetPpoison Precise fault caused by RPOISON or TEBRx.POISON
func (r *registerAfsrType) SetPpoison(value bool) {
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
func (r *registerAfsrType) GetFitcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFitcmMask) != 0
}

// SetFitcm Fetch fault on Instruction Tightly Coupled Memory (ITCM) interface
func (r *registerAfsrType) SetFitcm(value bool) {
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
func (r *registerAfsrType) GetFdtcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFdtcmMask) != 0
}

// SetFdtcm Fetch fault on Data Tightly Coupled Memory (DTCM) interface
func (r *registerAfsrType) SetFdtcm(value bool) {
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
func (r *registerAfsrType) GetFmaxi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFmaxiMask) != 0
}

// SetFmaxi Fetch fault on Master AXI (M-AXI) interface
func (r *registerAfsrType) SetFmaxi(value bool) {
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
func (r *registerAfsrType) GetFmaxitype() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFmaxitypeMask) != 0
}

// SetFmaxitype AXI response that caused the fetch fault. Only valid when AFSR.FMAXI is 1.
func (r *registerAfsrType) SetFmaxitype(value bool) {
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
func (r *registerAfsrType) GetFecc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFeccMask) != 0
}

// SetFecc Fetch fault caused by uncorrectable Error Correcting Code (ECC) error
func (r *registerAfsrType) SetFecc(value bool) {
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
func (r *registerAfsrType) GetFtgu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFtguMask) != 0
}

// SetFtgu Fetch fault caused by TCM Gate Unit (TGU) security violation
func (r *registerAfsrType) SetFtgu(value bool) {
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
func (r *registerAfsrType) GetFpoison() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAfsrFieldFpoisonMask) != 0
}

// SetFpoison Fetch fault caused by RPOISON or TEBRx.POISON
func (r *registerAfsrType) SetFpoison(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAfsrFieldFpoisonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAfsrFieldFpoisonMask)
	}
}

// registerIdpfr0Type Indicates the version of the Reliability, Availability, and Serviceability (RAS) extension supported
type registerIdpfr0Type uint32

const (
	RegisterIdpfr0FieldState0Shift = 0
	RegisterIdpfr0FieldState0Mask  = 0xf
)

// GetState0 A32 instruction set support
func (r *registerIdpfr0Type) GetState0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdpfr0FieldState0Mask) >> RegisterIdpfr0FieldState0Shift)
}

// SetState0 A32 instruction set support
func (r *registerIdpfr0Type) SetState0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdpfr0FieldState0Mask)|(uint32(value)<<RegisterIdpfr0FieldState0Shift))
}

const (
	RegisterIdpfr0FieldState1Shift = 4
	RegisterIdpfr0FieldState1Mask  = 0xf0
)

// GetState1 T32 instruction set support
func (r *registerIdpfr0Type) GetState1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdpfr0FieldState1Mask) >> RegisterIdpfr0FieldState1Shift)
}

// SetState1 T32 instruction set support
func (r *registerIdpfr0Type) SetState1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdpfr0FieldState1Mask)|(uint32(value)<<RegisterIdpfr0FieldState1Shift))
}

const (
	RegisterIdpfr0FieldRasShift = 28
	RegisterIdpfr0FieldRasMask  = 0xf0000000
)

// GetRas Identifies which version of the RAS architecture is implemented
func (r *registerIdpfr0Type) GetRas() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdpfr0FieldRasMask) >> RegisterIdpfr0FieldRasShift)
}

// SetRas Identifies which version of the RAS architecture is implemented
func (r *registerIdpfr0Type) SetRas(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdpfr0FieldRasMask)|(uint32(value)<<RegisterIdpfr0FieldRasShift))
}

// registerIdpfr1Type Gives information about the programmers' model and Extensions support
type registerIdpfr1Type uint32

const (
	RegisterIdpfr1FieldSecurityShift = 4
	RegisterIdpfr1FieldSecurityMask  = 0xf0
)

// GetSecurity Identifies whether the Security Extension in implemented
func (r *registerIdpfr1Type) GetSecurity() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdpfr1FieldSecurityMask) >> RegisterIdpfr1FieldSecurityShift)
}

// SetSecurity Identifies whether the Security Extension in implemented
func (r *registerIdpfr1Type) SetSecurity(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdpfr1FieldSecurityMask)|(uint32(value)<<RegisterIdpfr1FieldSecurityShift))
}

const (
	RegisterIdpfr1FieldMprogmodShift = 8
	RegisterIdpfr1FieldMprogmodMask  = 0xf00
)

// GetMprogmod M-profile programmers' model
func (r *registerIdpfr1Type) GetMprogmod() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdpfr1FieldMprogmodMask) >> RegisterIdpfr1FieldMprogmodShift)
}

// SetMprogmod M-profile programmers' model
func (r *registerIdpfr1Type) SetMprogmod(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdpfr1FieldMprogmodMask)|(uint32(value)<<RegisterIdpfr1FieldMprogmodShift))
}

// registerIddfr0Type no information available
type registerIddfr0Type uint32

const (
	RegisterIddfr0FieldMprofdbgShift = 20
	RegisterIddfr0FieldMprofdbgMask  = 0xf00000
)

// GetMprofdbg Indicates the supported M-Proﬁle debug architecture
func (r *registerIddfr0Type) GetMprofdbg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIddfr0FieldMprofdbgMask) >> RegisterIddfr0FieldMprofdbgShift)
}

// SetMprofdbg Indicates the supported M-Proﬁle debug architecture
func (r *registerIddfr0Type) SetMprofdbg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIddfr0FieldMprofdbgMask)|(uint32(value)<<RegisterIddfr0FieldMprofdbgShift))
}

const (
	RegisterIddfr0FieldUdeShift = 28
	RegisterIddfr0FieldUdeMask  = 0xf0000000
)

// GetUde Indicates support for the Unprivileged Debug Extension
func (r *registerIddfr0Type) GetUde() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIddfr0FieldUdeMask) >> RegisterIddfr0FieldUdeShift)
}

// SetUde Indicates support for the Unprivileged Debug Extension
func (r *registerIddfr0Type) SetUde(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIddfr0FieldUdeMask)|(uint32(value)<<RegisterIddfr0FieldUdeShift))
}

// registerIdafr0Type The ID_AFR0 is fully Reserved, RES0
type registerIdafr0Type uint32

// registerIdmmfr0Type Gives information about the implemented memory model and memory management support
type registerIdmmfr0Type uint32

const (
	RegisterIdmmfr0FieldPmsaShift = 4
	RegisterIdmmfr0FieldPmsaMask  = 0xf0
)

// GetPmsa Indicates support for the protected memory system architecture (PMSA)
func (r *registerIdmmfr0Type) GetPmsa() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr0FieldPmsaMask) >> RegisterIdmmfr0FieldPmsaShift)
}

// SetPmsa Indicates support for the protected memory system architecture (PMSA)
func (r *registerIdmmfr0Type) SetPmsa(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr0FieldPmsaMask)|(uint32(value)<<RegisterIdmmfr0FieldPmsaShift))
}

const (
	RegisterIdmmfr0FieldOutershrShift = 8
	RegisterIdmmfr0FieldOutershrMask  = 0xf00
)

// GetOutershr Indicates the outermost Shareability domain implemented
func (r *registerIdmmfr0Type) GetOutershr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr0FieldOutershrMask) >> RegisterIdmmfr0FieldOutershrShift)
}

// SetOutershr Indicates the outermost Shareability domain implemented
func (r *registerIdmmfr0Type) SetOutershr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr0FieldOutershrMask)|(uint32(value)<<RegisterIdmmfr0FieldOutershrShift))
}

const (
	RegisterIdmmfr0FieldSharelvlShift = 12
	RegisterIdmmfr0FieldSharelvlMask  = 0xf000
)

// GetSharelvl Indicates the number of Shareability levels implemented
func (r *registerIdmmfr0Type) GetSharelvl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr0FieldSharelvlMask) >> RegisterIdmmfr0FieldSharelvlShift)
}

// SetSharelvl Indicates the number of Shareability levels implemented
func (r *registerIdmmfr0Type) SetSharelvl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr0FieldSharelvlMask)|(uint32(value)<<RegisterIdmmfr0FieldSharelvlShift))
}

const (
	RegisterIdmmfr0FieldTcmShift = 16
	RegisterIdmmfr0FieldTcmMask  = 0xf0000
)

// GetTcm Indicates support for TCMs
func (r *registerIdmmfr0Type) GetTcm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr0FieldTcmMask) >> RegisterIdmmfr0FieldTcmShift)
}

// SetTcm Indicates support for TCMs
func (r *registerIdmmfr0Type) SetTcm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr0FieldTcmMask)|(uint32(value)<<RegisterIdmmfr0FieldTcmShift))
}

const (
	RegisterIdmmfr0FieldAuxregShift = 20
	RegisterIdmmfr0FieldAuxregMask  = 0xf00000
)

// GetAuxreg Indicates support for Auxiliary Control registers
func (r *registerIdmmfr0Type) GetAuxreg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr0FieldAuxregMask) >> RegisterIdmmfr0FieldAuxregShift)
}

// SetAuxreg Indicates support for Auxiliary Control registers
func (r *registerIdmmfr0Type) SetAuxreg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr0FieldAuxregMask)|(uint32(value)<<RegisterIdmmfr0FieldAuxregShift))
}

// registerIdmmfr1Type Gives information about the IMPLEMENTATION DEFINED memory model and memory management support. This register is Reserved, RES0.
type registerIdmmfr1Type uint32

// registerIdmmfr2Type Gives information about the implemented memory model and memory management support
type registerIdmmfr2Type uint32

const (
	RegisterIdmmfr2FieldWfistallShift = 24
	RegisterIdmmfr2FieldWfistallMask  = 0xf000000
)

// GetWfistall Indicates the support for Wait For Interrupt (WFI) stalling
func (r *registerIdmmfr2Type) GetWfistall() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr2FieldWfistallMask) >> RegisterIdmmfr2FieldWfistallShift)
}

// SetWfistall Indicates the support for Wait For Interrupt (WFI) stalling
func (r *registerIdmmfr2Type) SetWfistall(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr2FieldWfistallMask)|(uint32(value)<<RegisterIdmmfr2FieldWfistallShift))
}

// registerIdmmfr3Type Gives information about the implemented memory model and memory management support
type registerIdmmfr3Type uint32

const (
	RegisterIdmmfr3FieldCmaintvaShift = 0
	RegisterIdmmfr3FieldCmaintvaMask  = 0xf
)

// GetCmaintva Indicates the supported cache maintenance operations by address
func (r *registerIdmmfr3Type) GetCmaintva() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr3FieldCmaintvaMask) >> RegisterIdmmfr3FieldCmaintvaShift)
}

// SetCmaintva Indicates the supported cache maintenance operations by address
func (r *registerIdmmfr3Type) SetCmaintva(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr3FieldCmaintvaMask)|(uint32(value)<<RegisterIdmmfr3FieldCmaintvaShift))
}

const (
	RegisterIdmmfr3FieldCmaintswShift = 4
	RegisterIdmmfr3FieldCmaintswMask  = 0xf0
)

// GetCmaintsw Indicates the supported cache maintenance operations by set/way
func (r *registerIdmmfr3Type) GetCmaintsw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr3FieldCmaintswMask) >> RegisterIdmmfr3FieldCmaintswShift)
}

// SetCmaintsw Indicates the supported cache maintenance operations by set/way
func (r *registerIdmmfr3Type) SetCmaintsw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr3FieldCmaintswMask)|(uint32(value)<<RegisterIdmmfr3FieldCmaintswShift))
}

const (
	RegisterIdmmfr3FieldBpmaintShift = 8
	RegisterIdmmfr3FieldBpmaintMask  = 0xf00
)

// GetBpmaint Indicates the supported branch predictor maintenance
func (r *registerIdmmfr3Type) GetBpmaint() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmmfr3FieldBpmaintMask) >> RegisterIdmmfr3FieldBpmaintShift)
}

// SetBpmaint Indicates the supported branch predictor maintenance
func (r *registerIdmmfr3Type) SetBpmaint(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmmfr3FieldBpmaintMask)|(uint32(value)<<RegisterIdmmfr3FieldBpmaintShift))
}

// registerIdisar0Type Gives information about the implemented instruction set
type registerIdisar0Type uint32

const (
	RegisterIdisar0FieldBitcountShift = 4
	RegisterIdisar0FieldBitcountMask  = 0xf0
)

// GetBitcount Indicates the supported bit count instructions
func (r *registerIdisar0Type) GetBitcount() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar0FieldBitcountMask) >> RegisterIdisar0FieldBitcountShift)
}

// SetBitcount Indicates the supported bit count instructions
func (r *registerIdisar0Type) SetBitcount(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar0FieldBitcountMask)|(uint32(value)<<RegisterIdisar0FieldBitcountShift))
}

const (
	RegisterIdisar0FieldBitfieldShift = 8
	RegisterIdisar0FieldBitfieldMask  = 0xf00
)

// GetBitfield Indicates the supported bit field instructions
func (r *registerIdisar0Type) GetBitfield() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar0FieldBitfieldMask) >> RegisterIdisar0FieldBitfieldShift)
}

// SetBitfield Indicates the supported bit field instructions
func (r *registerIdisar0Type) SetBitfield(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar0FieldBitfieldMask)|(uint32(value)<<RegisterIdisar0FieldBitfieldShift))
}

const (
	RegisterIdisar0FieldCmpbranchShift = 12
	RegisterIdisar0FieldCmpbranchMask  = 0xf000
)

// GetCmpbranch Indicates the supported combined Compare and Branch instructions
func (r *registerIdisar0Type) GetCmpbranch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar0FieldCmpbranchMask) >> RegisterIdisar0FieldCmpbranchShift)
}

// SetCmpbranch Indicates the supported combined Compare and Branch instructions
func (r *registerIdisar0Type) SetCmpbranch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar0FieldCmpbranchMask)|(uint32(value)<<RegisterIdisar0FieldCmpbranchShift))
}

const (
	RegisterIdisar0FieldCoprocShift = 16
	RegisterIdisar0FieldCoprocMask  = 0xf0000
)

// GetCoproc Indicates the supported coprocessor instructions
func (r *registerIdisar0Type) GetCoproc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar0FieldCoprocMask) >> RegisterIdisar0FieldCoprocShift)
}

// SetCoproc Indicates the supported coprocessor instructions
func (r *registerIdisar0Type) SetCoproc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar0FieldCoprocMask)|(uint32(value)<<RegisterIdisar0FieldCoprocShift))
}

const (
	RegisterIdisar0FieldDebugShift = 20
	RegisterIdisar0FieldDebugMask  = 0xf00000
)

// GetDebug Indicates the implemented debug instructions
func (r *registerIdisar0Type) GetDebug() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar0FieldDebugMask) >> RegisterIdisar0FieldDebugShift)
}

// SetDebug Indicates the implemented debug instructions
func (r *registerIdisar0Type) SetDebug(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar0FieldDebugMask)|(uint32(value)<<RegisterIdisar0FieldDebugShift))
}

const (
	RegisterIdisar0FieldDivideShift = 24
	RegisterIdisar0FieldDivideMask  = 0xf000000
)

// GetDivide Indicates the supported divide instructions
func (r *registerIdisar0Type) GetDivide() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar0FieldDivideMask) >> RegisterIdisar0FieldDivideShift)
}

// SetDivide Indicates the supported divide instructions
func (r *registerIdisar0Type) SetDivide(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar0FieldDivideMask)|(uint32(value)<<RegisterIdisar0FieldDivideShift))
}

// registerIdisar1Type Gives information about the implemented instruction set
type registerIdisar1Type uint32

const (
	RegisterIdisar1FieldExtendShift = 12
	RegisterIdisar1FieldExtendMask  = 0xf000
)

// GetExtend Indicates the implemented Extend instructions
func (r *registerIdisar1Type) GetExtend() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar1FieldExtendMask) >> RegisterIdisar1FieldExtendShift)
}

// SetExtend Indicates the implemented Extend instructions
func (r *registerIdisar1Type) SetExtend(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar1FieldExtendMask)|(uint32(value)<<RegisterIdisar1FieldExtendShift))
}

const (
	RegisterIdisar1FieldIfthenShift = 16
	RegisterIdisar1FieldIfthenMask  = 0xf0000
)

// GetIfthen Indicates the implemented If-Then instructions
func (r *registerIdisar1Type) GetIfthen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar1FieldIfthenMask) >> RegisterIdisar1FieldIfthenShift)
}

// SetIfthen Indicates the implemented If-Then instructions
func (r *registerIdisar1Type) SetIfthen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar1FieldIfthenMask)|(uint32(value)<<RegisterIdisar1FieldIfthenShift))
}

const (
	RegisterIdisar1FieldImmediateShift = 20
	RegisterIdisar1FieldImmediateMask  = 0xf00000
)

// GetImmediate Indicates the implemented data-processing instructions with long immediates
func (r *registerIdisar1Type) GetImmediate() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar1FieldImmediateMask) >> RegisterIdisar1FieldImmediateShift)
}

// SetImmediate Indicates the implemented data-processing instructions with long immediates
func (r *registerIdisar1Type) SetImmediate(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar1FieldImmediateMask)|(uint32(value)<<RegisterIdisar1FieldImmediateShift))
}

const (
	RegisterIdisar1FieldInterworkShift = 24
	RegisterIdisar1FieldInterworkMask  = 0xf000000
)

// GetInterwork Indicates the implemented interworking instructions
func (r *registerIdisar1Type) GetInterwork() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar1FieldInterworkMask) >> RegisterIdisar1FieldInterworkShift)
}

// SetInterwork Indicates the implemented interworking instructions
func (r *registerIdisar1Type) SetInterwork(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar1FieldInterworkMask)|(uint32(value)<<RegisterIdisar1FieldInterworkShift))
}

// registerIdisar2Type Gives information about the implemented instruction set
type registerIdisar2Type uint32

const (
	RegisterIdisar2FieldLoadstoreShift = 0
	RegisterIdisar2FieldLoadstoreMask  = 0xf
)

// GetLoadstore Indicates the implemented additional load/store instructions
func (r *registerIdisar2Type) GetLoadstore() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldLoadstoreMask) >> RegisterIdisar2FieldLoadstoreShift)
}

// SetLoadstore Indicates the implemented additional load/store instructions
func (r *registerIdisar2Type) SetLoadstore(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldLoadstoreMask)|(uint32(value)<<RegisterIdisar2FieldLoadstoreShift))
}

const (
	RegisterIdisar2FieldMemhintShift = 4
	RegisterIdisar2FieldMemhintMask  = 0xf0
)

// GetMemhint Memory hint. Indicates the implemented memory hint instructions
func (r *registerIdisar2Type) GetMemhint() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldMemhintMask) >> RegisterIdisar2FieldMemhintShift)
}

// SetMemhint Memory hint. Indicates the implemented memory hint instructions
func (r *registerIdisar2Type) SetMemhint(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldMemhintMask)|(uint32(value)<<RegisterIdisar2FieldMemhintShift))
}

const (
	RegisterIdisar2FieldMultiaccessintShift = 8
	RegisterIdisar2FieldMultiaccessintMask  = 0xf00
)

// GetMultiaccessint Multi-access instructions. Indicates the support for interruptible multi-access instructions
func (r *registerIdisar2Type) GetMultiaccessint() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldMultiaccessintMask) >> RegisterIdisar2FieldMultiaccessintShift)
}

// SetMultiaccessint Multi-access instructions. Indicates the support for interruptible multi-access instructions
func (r *registerIdisar2Type) SetMultiaccessint(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldMultiaccessintMask)|(uint32(value)<<RegisterIdisar2FieldMultiaccessintShift))
}

const (
	RegisterIdisar2FieldMultShift = 12
	RegisterIdisar2FieldMultMask  = 0xf000
)

// GetMult Indicates the implemented additional Multiply instructions
func (r *registerIdisar2Type) GetMult() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldMultMask) >> RegisterIdisar2FieldMultShift)
}

// SetMult Indicates the implemented additional Multiply instructions
func (r *registerIdisar2Type) SetMult(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldMultMask)|(uint32(value)<<RegisterIdisar2FieldMultShift))
}

const (
	RegisterIdisar2FieldMultsShift = 16
	RegisterIdisar2FieldMultsMask  = 0xf0000
)

// GetMults Indicates the implemented Advanced Signed Multiple instructions
func (r *registerIdisar2Type) GetMults() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldMultsMask) >> RegisterIdisar2FieldMultsShift)
}

// SetMults Indicates the implemented Advanced Signed Multiple instructions
func (r *registerIdisar2Type) SetMults(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldMultsMask)|(uint32(value)<<RegisterIdisar2FieldMultsShift))
}

const (
	RegisterIdisar2FieldMultuShift = 20
	RegisterIdisar2FieldMultuMask  = 0xf00000
)

// GetMultu Indicates the implemented Advanced Unsigned Multiple instructions
func (r *registerIdisar2Type) GetMultu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldMultuMask) >> RegisterIdisar2FieldMultuShift)
}

// SetMultu Indicates the implemented Advanced Unsigned Multiple instructions
func (r *registerIdisar2Type) SetMultu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldMultuMask)|(uint32(value)<<RegisterIdisar2FieldMultuShift))
}

const (
	RegisterIdisar2FieldReversalShift = 28
	RegisterIdisar2FieldReversalMask  = 0xf0000000
)

// GetReversal Indicates the implemented reversal instructions
func (r *registerIdisar2Type) GetReversal() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar2FieldReversalMask) >> RegisterIdisar2FieldReversalShift)
}

// SetReversal Indicates the implemented reversal instructions
func (r *registerIdisar2Type) SetReversal(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar2FieldReversalMask)|(uint32(value)<<RegisterIdisar2FieldReversalShift))
}

// registerIdisar3Type Gives information about the implemented instruction set
type registerIdisar3Type uint32

const (
	RegisterIdisar3FieldSaturateShift = 0
	RegisterIdisar3FieldSaturateMask  = 0xf
)

// GetSaturate Indicates the implemented saturating instructions
func (r *registerIdisar3Type) GetSaturate() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldSaturateMask) >> RegisterIdisar3FieldSaturateShift)
}

// SetSaturate Indicates the implemented saturating instructions
func (r *registerIdisar3Type) SetSaturate(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldSaturateMask)|(uint32(value)<<RegisterIdisar3FieldSaturateShift))
}

const (
	RegisterIdisar3FieldSimdShift = 4
	RegisterIdisar3FieldSimdMask  = 0xf0
)

// GetSimd Indicates the implemented SIMD instructions
func (r *registerIdisar3Type) GetSimd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldSimdMask) >> RegisterIdisar3FieldSimdShift)
}

// SetSimd Indicates the implemented SIMD instructions
func (r *registerIdisar3Type) SetSimd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldSimdMask)|(uint32(value)<<RegisterIdisar3FieldSimdShift))
}

const (
	RegisterIdisar3FieldSvcShift = 8
	RegisterIdisar3FieldSvcMask  = 0xf00
)

// GetSvc Indicates the implemented SVC instructions
func (r *registerIdisar3Type) GetSvc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldSvcMask) >> RegisterIdisar3FieldSvcShift)
}

// SetSvc Indicates the implemented SVC instructions
func (r *registerIdisar3Type) SetSvc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldSvcMask)|(uint32(value)<<RegisterIdisar3FieldSvcShift))
}

const (
	RegisterIdisar3FieldSynchprimShift = 12
	RegisterIdisar3FieldSynchprimMask  = 0xf000
)

// GetSynchprim Used with ID_ISAR4.SynchPrim_frac to indicate the implemented synchronization primitive instructions
func (r *registerIdisar3Type) GetSynchprim() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldSynchprimMask) >> RegisterIdisar3FieldSynchprimShift)
}

// SetSynchprim Used with ID_ISAR4.SynchPrim_frac to indicate the implemented synchronization primitive instructions
func (r *registerIdisar3Type) SetSynchprim(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldSynchprimMask)|(uint32(value)<<RegisterIdisar3FieldSynchprimShift))
}

const (
	RegisterIdisar3FieldTabbranchShift = 16
	RegisterIdisar3FieldTabbranchMask  = 0xf0000
)

// GetTabbranch Indicates the implemented table branch instructions
func (r *registerIdisar3Type) GetTabbranch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldTabbranchMask) >> RegisterIdisar3FieldTabbranchShift)
}

// SetTabbranch Indicates the implemented table branch instructions
func (r *registerIdisar3Type) SetTabbranch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldTabbranchMask)|(uint32(value)<<RegisterIdisar3FieldTabbranchShift))
}

const (
	RegisterIdisar3FieldT32copyShift = 20
	RegisterIdisar3FieldT32copyMask  = 0xf00000
)

// GetT32copy Indicates the support for T32 non-flag setting MOV instructions
func (r *registerIdisar3Type) GetT32copy() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldT32copyMask) >> RegisterIdisar3FieldT32copyShift)
}

// SetT32copy Indicates the support for T32 non-flag setting MOV instructions
func (r *registerIdisar3Type) SetT32copy(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldT32copyMask)|(uint32(value)<<RegisterIdisar3FieldT32copyShift))
}

const (
	RegisterIdisar3FieldTruenopShift = 24
	RegisterIdisar3FieldTruenopMask  = 0xf000000
)

// GetTruenop Indicates the implemented true NOP instructions
func (r *registerIdisar3Type) GetTruenop() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar3FieldTruenopMask) >> RegisterIdisar3FieldTruenopShift)
}

// SetTruenop Indicates the implemented true NOP instructions
func (r *registerIdisar3Type) SetTruenop(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar3FieldTruenopMask)|(uint32(value)<<RegisterIdisar3FieldTruenopShift))
}

// registerIdisar4Type Gives information about the implemented instruction set
type registerIdisar4Type uint32

const (
	RegisterIdisar4FieldUnprivShift = 0
	RegisterIdisar4FieldUnprivMask  = 0xf
)

// GetUnpriv Indicates the implemented unprivileged instructions
func (r *registerIdisar4Type) GetUnpriv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar4FieldUnprivMask) >> RegisterIdisar4FieldUnprivShift)
}

// SetUnpriv Indicates the implemented unprivileged instructions
func (r *registerIdisar4Type) SetUnpriv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar4FieldUnprivMask)|(uint32(value)<<RegisterIdisar4FieldUnprivShift))
}

const (
	RegisterIdisar4FieldWithshiftsShift = 4
	RegisterIdisar4FieldWithshiftsMask  = 0xf0
)

// GetWithshifts Indicates the support for writeback addressing modes
func (r *registerIdisar4Type) GetWithshifts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar4FieldWithshiftsMask) >> RegisterIdisar4FieldWithshiftsShift)
}

// SetWithshifts Indicates the support for writeback addressing modes
func (r *registerIdisar4Type) SetWithshifts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar4FieldWithshiftsMask)|(uint32(value)<<RegisterIdisar4FieldWithshiftsShift))
}

const (
	RegisterIdisar4FieldWritebackShift = 8
	RegisterIdisar4FieldWritebackMask  = 0xf00
)

// GetWriteback Indicates the support for writeback addressing modes
func (r *registerIdisar4Type) GetWriteback() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar4FieldWritebackMask) >> RegisterIdisar4FieldWritebackShift)
}

// SetWriteback Indicates the support for writeback addressing modes
func (r *registerIdisar4Type) SetWriteback(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar4FieldWritebackMask)|(uint32(value)<<RegisterIdisar4FieldWritebackShift))
}

const (
	RegisterIdisar4FieldBarrierShift = 16
	RegisterIdisar4FieldBarrierMask  = 0xf0000
)

// GetBarrier Indicates the implemented Barrier instructions
func (r *registerIdisar4Type) GetBarrier() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar4FieldBarrierMask) >> RegisterIdisar4FieldBarrierShift)
}

// SetBarrier Indicates the implemented Barrier instructions
func (r *registerIdisar4Type) SetBarrier(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar4FieldBarrierMask)|(uint32(value)<<RegisterIdisar4FieldBarrierShift))
}

const (
	RegisterIdisar4FieldSynchprimfracShift = 20
	RegisterIdisar4FieldSynchprimfracMask  = 0xf00000
)

// GetSynchprimfrac Used in conjunction with ID_ISAR3.SynchPrim to indicate the implemented synchronization primitive instructions
func (r *registerIdisar4Type) GetSynchprimfrac() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar4FieldSynchprimfracMask) >> RegisterIdisar4FieldSynchprimfracShift)
}

// SetSynchprimfrac Used in conjunction with ID_ISAR3.SynchPrim to indicate the implemented synchronization primitive instructions
func (r *registerIdisar4Type) SetSynchprimfrac(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar4FieldSynchprimfracMask)|(uint32(value)<<RegisterIdisar4FieldSynchprimfracShift))
}

const (
	RegisterIdisar4FieldPsrmShift = 24
	RegisterIdisar4FieldPsrmMask  = 0xf000000
)

// GetPsrm Indicates the implemented M-profile instructions to modify the PSRs
func (r *registerIdisar4Type) GetPsrm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdisar4FieldPsrmMask) >> RegisterIdisar4FieldPsrmShift)
}

// SetPsrm Indicates the implemented M-profile instructions to modify the PSRs
func (r *registerIdisar4Type) SetPsrm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdisar4FieldPsrmMask)|(uint32(value)<<RegisterIdisar4FieldPsrmShift))
}

// registerIdisar5Type Gives information about the implemented instruction set. This register is Reserved, RES0
type registerIdisar5Type uint32

// registerClidrType The CLIDR identifies the type of caches implemented and the level of coherency and unification
type registerClidrType uint32

const (
	RegisterClidrFieldCtype1Shift = 0
	RegisterClidrFieldCtype1Mask  = 0x7
)

// GetCtype1 L1 cache type
func (r *registerClidrType) GetCtype1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldCtype1Mask) >> RegisterClidrFieldCtype1Shift)
}

// SetCtype1 L1 cache type
func (r *registerClidrType) SetCtype1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldCtype1Mask)|(uint32(value)<<RegisterClidrFieldCtype1Shift))
}

const (
	RegisterClidrFieldLouisShift = 21
	RegisterClidrFieldLouisMask  = 0xe00000
)

// GetLouis Level of Unification Inner Shareable
func (r *registerClidrType) GetLouis() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldLouisMask) >> RegisterClidrFieldLouisShift)
}

// SetLouis Level of Unification Inner Shareable
func (r *registerClidrType) SetLouis(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldLouisMask)|(uint32(value)<<RegisterClidrFieldLouisShift))
}

const (
	RegisterClidrFieldLocShift = 24
	RegisterClidrFieldLocMask  = 0x7000000
)

// GetLoc Level of Coherency
func (r *registerClidrType) GetLoc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldLocMask) >> RegisterClidrFieldLocShift)
}

// SetLoc Level of Coherency
func (r *registerClidrType) SetLoc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldLocMask)|(uint32(value)<<RegisterClidrFieldLocShift))
}

const (
	RegisterClidrFieldLouuShift = 27
	RegisterClidrFieldLouuMask  = 0x38000000
)

// GetLouu Level of Unification Uniprocessor
func (r *registerClidrType) GetLouu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldLouuMask) >> RegisterClidrFieldLouuShift)
}

// SetLouu Level of Unification Uniprocessor
func (r *registerClidrType) SetLouu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldLouuMask)|(uint32(value)<<RegisterClidrFieldLouuShift))
}

const (
	RegisterClidrFieldIcbShift = 30
	RegisterClidrFieldIcbMask  = 0xc0000000
)

// GetIcb Inner cache boundary. The Cortex-M55 processor supports inner cacheability on the bus. Therefore, this field cannot disclose any information.
func (r *registerClidrType) GetIcb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldIcbMask) >> RegisterClidrFieldIcbShift)
}

// SetIcb Inner cache boundary. The Cortex-M55 processor supports inner cacheability on the bus. Therefore, this field cannot disclose any information.
func (r *registerClidrType) SetIcb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldIcbMask)|(uint32(value)<<RegisterClidrFieldIcbShift))
}

// registerCtrType The CTR provides information about the architecture of the caches
type registerCtrType uint32

const (
	RegisterCtrFieldIminlineShift = 0
	RegisterCtrFieldIminlineMask  = 0xf
)

// GetIminline Instruction cache minimum line length. Log2 of the number of words in the smallest cache line of all the instruction caches that are controlled by the processor.
func (r *registerCtrType) GetIminline() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldIminlineMask) >> RegisterCtrFieldIminlineShift)
}

// SetIminline Instruction cache minimum line length. Log2 of the number of words in the smallest cache line of all the instruction caches that are controlled by the processor.
func (r *registerCtrType) SetIminline(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldIminlineMask)|(uint32(value)<<RegisterCtrFieldIminlineShift))
}

const (
	RegisterCtrFieldDminlineShift = 16
	RegisterCtrFieldDminlineMask  = 0xf0000
)

// GetDminline Data cache minimum line length. Log2 of the number of words in the smallest cache line of all the data caches and unified caches that are controlled by the processor.
func (r *registerCtrType) GetDminline() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldDminlineMask) >> RegisterCtrFieldDminlineShift)
}

// SetDminline Data cache minimum line length. Log2 of the number of words in the smallest cache line of all the data caches and unified caches that are controlled by the processor.
func (r *registerCtrType) SetDminline(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldDminlineMask)|(uint32(value)<<RegisterCtrFieldDminlineShift))
}

const (
	RegisterCtrFieldErgShift = 20
	RegisterCtrFieldErgMask  = 0xf00000
)

// GetErg Exclusives Reservation Granule
func (r *registerCtrType) GetErg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldErgMask) >> RegisterCtrFieldErgShift)
}

// SetErg Exclusives Reservation Granule
func (r *registerCtrType) SetErg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldErgMask)|(uint32(value)<<RegisterCtrFieldErgShift))
}

const (
	RegisterCtrFieldCwgShift = 24
	RegisterCtrFieldCwgMask  = 0xf000000
)

// GetCwg Cache Write-Back Granule
func (r *registerCtrType) GetCwg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldCwgMask) >> RegisterCtrFieldCwgShift)
}

// SetCwg Cache Write-Back Granule
func (r *registerCtrType) SetCwg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldCwgMask)|(uint32(value)<<RegisterCtrFieldCwgShift))
}

const (
	RegisterCtrFieldFormatShift = 29
	RegisterCtrFieldFormatMask  = 0xe0000000
)

// GetFormat Cache type register format
func (r *registerCtrType) GetFormat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldFormatMask) >> RegisterCtrFieldFormatShift)
}

// SetFormat Cache type register format
func (r *registerCtrType) SetFormat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldFormatMask)|(uint32(value)<<RegisterCtrFieldFormatShift))
}

// registerCcsidrType Provides information about the architecture of the Level 1 (L1) instruction or data cache that the CSSELR selects.
type registerCcsidrType uint32

const (
	RegisterCcsidrFieldWtShift = 0
	RegisterCcsidrFieldWtMask  = 0x0
)

// GetWt Indicates support available for Write-Through
func (r *registerCcsidrType) GetWt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldWtMask) >> RegisterCcsidrFieldWtShift)
}

// SetWt Indicates support available for Write-Through
func (r *registerCcsidrType) SetWt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldWtMask)|(uint32(value)<<RegisterCcsidrFieldWtShift))
}

const (
	RegisterCcsidrFieldLinesizeShift = 0
	RegisterCcsidrFieldLinesizeMask  = 0x7
)

// GetLinesize Indicates the number of words in each cache line
func (r *registerCcsidrType) GetLinesize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldLinesizeMask) >> RegisterCcsidrFieldLinesizeShift)
}

// SetLinesize Indicates the number of words in each cache line
func (r *registerCcsidrType) SetLinesize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldLinesizeMask)|(uint32(value)<<RegisterCcsidrFieldLinesizeShift))
}

const (
	RegisterCcsidrFieldAssociativityShift = 3
	RegisterCcsidrFieldAssociativityMask  = 0x1ff8
)

// GetAssociativity Indicates associativity
func (r *registerCcsidrType) GetAssociativity() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldAssociativityMask) >> RegisterCcsidrFieldAssociativityShift)
}

// SetAssociativity Indicates associativity
func (r *registerCcsidrType) SetAssociativity(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldAssociativityMask)|(uint32(value)<<RegisterCcsidrFieldAssociativityShift))
}

const (
	RegisterCcsidrFieldNumsetShift = 13
	RegisterCcsidrFieldNumsetMask  = 0xfffe000
)

// GetNumset Indicates the number of sets
func (r *registerCcsidrType) GetNumset() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldNumsetMask) >> RegisterCcsidrFieldNumsetShift)
}

// SetNumset Indicates the number of sets
func (r *registerCcsidrType) SetNumset(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldNumsetMask)|(uint32(value)<<RegisterCcsidrFieldNumsetShift))
}

const (
	RegisterCcsidrFieldWaShift = 28
	RegisterCcsidrFieldWaMask  = 0x10000000
)

// GetWa Indicates support available for write allocation
func (r *registerCcsidrType) GetWa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldWaMask) != 0
}

// SetWa Indicates support available for write allocation
func (r *registerCcsidrType) SetWa(value bool) {
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
func (r *registerCcsidrType) GetRa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldRaMask) != 0
}

// SetRa Indicates support available for read allocation
func (r *registerCcsidrType) SetRa(value bool) {
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
func (r *registerCcsidrType) GetWb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldWbMask) != 0
}

// SetWb Indicates support available for Write-Back
func (r *registerCcsidrType) SetWb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcsidrFieldWbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldWbMask)
	}
}

// registerCsselrType The CSSELR selects the current CCSIDR by specifying the cache level and the type of cache
type registerCsselrType uint32

const (
	RegisterCsselrFieldIndShift = 0
	RegisterCsselrFieldIndMask  = 0x1
)

// GetInd Instruction not data bit
func (r *registerCsselrType) GetInd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsselrFieldIndMask) != 0
}

// SetInd Instruction not data bit
func (r *registerCsselrType) SetInd(value bool) {
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
func (r *registerCsselrType) GetLevel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCsselrFieldLevelMask) >> RegisterCsselrFieldLevelShift)
}

// SetLevel Cache level of required cache
func (r *registerCsselrType) SetLevel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsselrFieldLevelMask)|(uint32(value)<<RegisterCsselrFieldLevelShift))
}

// registerCpacrType Specifies the access privileges for coprocessors
type registerCpacrType uint32

const (
	RegisterCpacrFieldCpShift = 0
	RegisterCpacrFieldCpMask  = 0x3
)

// GetCp Coprocessor %s privilege
func (r *registerCpacrType) GetCp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpacrFieldCpMask) >> RegisterCpacrFieldCpShift)
}

// SetCp Coprocessor %s privilege
func (r *registerCpacrType) SetCp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpacrFieldCpMask)|(uint32(value)<<RegisterCpacrFieldCpShift))
}

type RegisterCpacrFieldCp_10EnumType uint8

const (
	// RegisterCpacrFieldCp10EnumDisabled All accesses to the FP Extension result in NOCP UsageFault.
	RegisterCpacrFieldCp10EnumDisabled RegisterCpacrFieldCp_10EnumType = 0x0

	// RegisterCpacrFieldCp10EnumUnprivileged Unprivileged accesses to the FP Extension result in NOCP UsageFault.
	RegisterCpacrFieldCp10EnumUnprivileged RegisterCpacrFieldCp_10EnumType = 0x1

	// RegisterCpacrFieldCp10EnumFull Full access to the FP Extension.
	RegisterCpacrFieldCp10EnumFull RegisterCpacrFieldCp_10EnumType = 0x2

	RegisterCpacrFieldCp10Shift = 20
	RegisterCpacrFieldCp10Mask  = 0x300000
)

// GetCp10 CP10 Privilege. Defines the access rights for the floating-point functionality.
func (r *registerCpacrType) GetCp10() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpacrFieldCp10Mask) >> RegisterCpacrFieldCp10Shift)
}

// SetCp10 CP10 Privilege. Defines the access rights for the floating-point functionality.
func (r *registerCpacrType) SetCp10(value RegisterCpacrFieldCp_10EnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpacrFieldCp10Mask)|(uint32(value)<<RegisterCpacrFieldCp10Shift))
}

type RegisterCpacrFieldCp_11EnumType uint8

const (
	// RegisterCpacrFieldCp11EnumDisabled All accesses to the FP Extension result in NOCP UsageFault.
	RegisterCpacrFieldCp11EnumDisabled RegisterCpacrFieldCp_11EnumType = 0x0

	// RegisterCpacrFieldCp11EnumUnprivileged Unprivileged accesses to the FP Extension result in NOCP UsageFault.
	RegisterCpacrFieldCp11EnumUnprivileged RegisterCpacrFieldCp_11EnumType = 0x1

	// RegisterCpacrFieldCp11EnumFull Full access to the FP Extension.
	RegisterCpacrFieldCp11EnumFull RegisterCpacrFieldCp_11EnumType = 0x2

	RegisterCpacrFieldCp11Shift = 22
	RegisterCpacrFieldCp11Mask  = 0xc00000
)

// GetCp11 CP11 Privilege. The value in this field is ignored.
func (r *registerCpacrType) GetCp11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpacrFieldCp11Mask) >> RegisterCpacrFieldCp11Shift)
}

// SetCp11 CP11 Privilege. The value in this field is ignored.
func (r *registerCpacrType) SetCp11(value RegisterCpacrFieldCp_11EnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpacrFieldCp11Mask)|(uint32(value)<<RegisterCpacrFieldCp11Shift))
}

// registerNsacrType In an implementation with the Security Extension, the NSACR register defines the Non-secure access permissions for both the Floating-point and MVE and coprocessors CP m, bit[ m], for m = 0-7. If MVE is implemented, this register specifies the Non-secure permissions for MVE
type registerNsacrType uint32

const (
	RegisterNsacrFieldCpShift = 0
	RegisterNsacrFieldCpMask  = 0x1
)

// GetCp Enables Non-secure access to this coprocessor
func (r *registerNsacrType) GetCp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterNsacrFieldCpMask) != 0
}

// SetCp Enables Non-secure access to this coprocessor
func (r *registerNsacrType) SetCp(value bool) {
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
func (r *registerNsacrType) GetCp11() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterNsacrFieldCp11Mask) >> RegisterNsacrFieldCp11Shift)
}

// SetCp11 Enables non-secure access to the Floating-point and MVE. Programming with a different value other than that used for CP10 is unpredictable.
func (r *registerNsacrType) SetCp11(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterNsacrFieldCp11Mask)|(uint32(value)<<RegisterNsacrFieldCp11Shift))
}

const (
	RegisterNsacrFieldCp10Shift = 10
	RegisterNsacrFieldCp10Mask  = 0x400
)

// GetCp10 Enables non-secure access to the Floating-point and MVE
func (r *registerNsacrType) GetCp10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterNsacrFieldCp10Mask) != 0
}

// SetCp10 Enables non-secure access to the Floating-point and MVE
func (r *registerNsacrType) SetCp10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterNsacrFieldCp10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterNsacrFieldCp10Mask)
	}
}
