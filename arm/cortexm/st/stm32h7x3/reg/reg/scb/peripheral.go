package scb

import (
	"unsafe"
	"volatile"
)

var (
	Scb = (*_scb)(unsafe.Pointer(uintptr(0xe000ed00)))
)

type _scb struct {
	Cpuid                registerCpuidType
	Icsr                 registerIcsrType
	Vtor                 registerVtorType
	Aircr                registerAircrType
	Scr                  registerScrType
	Ccr                  registerCcrType
	Shpr1                registerShpr1Type
	Shpr2                registerShpr2Type
	Shpr3                registerShpr3Type
	Shcrs                registerShcrsType
	Cfsr_ufsr_bfsr_mmfsr registerCfsr_ufsr_bfsr_mmfsrType
	Hfsr                 registerHfsrType
	_                    [4]byte
	Mmfar                registerMmfarType
	Bfar                 registerBfarType
}

// registerCpuidType CPUID base register
type registerCpuidType uint32

const (
	RegisterCpuidFieldRevisionShift = 0
	RegisterCpuidFieldRevisionMask  = 0xf
)

// GetRevision Revision number
func (r *registerCpuidType) GetRevision() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldRevisionMask) >> RegisterCpuidFieldRevisionShift)
}

// SetRevision Revision number
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
	RegisterCpuidFieldConstantShift = 16
	RegisterCpuidFieldConstantMask  = 0xf0000
)

// GetConstant Reads as 0xF
func (r *registerCpuidType) GetConstant() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldConstantMask) >> RegisterCpuidFieldConstantShift)
}

// SetConstant Reads as 0xF
func (r *registerCpuidType) SetConstant(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpuidFieldConstantMask)|(uint32(value)<<RegisterCpuidFieldConstantShift))
}

const (
	RegisterCpuidFieldVariantShift = 20
	RegisterCpuidFieldVariantMask  = 0xf00000
)

// GetVariant Variant number
func (r *registerCpuidType) GetVariant() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpuidFieldVariantMask) >> RegisterCpuidFieldVariantShift)
}

// SetVariant Variant number
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

// registerIcsrType Interrupt control and state register
type registerIcsrType uint32

const (
	RegisterIcsrFieldVectactiveShift = 0
	RegisterIcsrFieldVectactiveMask  = 0x1ff
)

// GetVectactive Active vector
func (r *registerIcsrType) GetVectactive() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldVectactiveMask) >> RegisterIcsrFieldVectactiveShift)
}

// SetVectactive Active vector
func (r *registerIcsrType) SetVectactive(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldVectactiveMask)|(uint32(value)<<RegisterIcsrFieldVectactiveShift))
}

const (
	RegisterIcsrFieldRettobaseShift = 11
	RegisterIcsrFieldRettobaseMask  = 0x800
)

// GetRettobase Return to base level
func (r *registerIcsrType) GetRettobase() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldRettobaseMask) != 0
}

// SetRettobase Return to base level
func (r *registerIcsrType) SetRettobase(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldRettobaseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldRettobaseMask)
	}
}

const (
	RegisterIcsrFieldVectpendingShift = 12
	RegisterIcsrFieldVectpendingMask  = 0x7f000
)

// GetVectpending Pending vector
func (r *registerIcsrType) GetVectpending() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldVectpendingMask) >> RegisterIcsrFieldVectpendingShift)
}

// SetVectpending Pending vector
func (r *registerIcsrType) SetVectpending(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldVectpendingMask)|(uint32(value)<<RegisterIcsrFieldVectpendingShift))
}

const (
	RegisterIcsrFieldIsrpendingShift = 22
	RegisterIcsrFieldIsrpendingMask  = 0x400000
)

// GetIsrpending Interrupt pending flag
func (r *registerIcsrType) GetIsrpending() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldIsrpendingMask) != 0
}

// SetIsrpending Interrupt pending flag
func (r *registerIcsrType) SetIsrpending(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldIsrpendingMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldIsrpendingMask)
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

// GetPendsvclr PendSV clear-pending bit
func (r *registerIcsrType) GetPendsvclr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldPendsvclrMask) != 0
}

// SetPendsvclr PendSV clear-pending bit
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
	RegisterIcsrFieldNmipendsetShift = 31
	RegisterIcsrFieldNmipendsetMask  = 0x80000000
)

// GetNmipendset NMI set-pending bit.
func (r *registerIcsrType) GetNmipendset() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcsrFieldNmipendsetMask) != 0
}

// SetNmipendset NMI set-pending bit.
func (r *registerIcsrType) SetNmipendset(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcsrFieldNmipendsetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcsrFieldNmipendsetMask)
	}
}

// registerVtorType Vector table offset register
type registerVtorType uint32

const (
	RegisterVtorFieldTbloffShift = 9
	RegisterVtorFieldTbloffMask  = 0x3ffffe00
)

// GetTbloff Vector table base offset field
func (r *registerVtorType) GetTbloff() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterVtorFieldTbloffMask) >> RegisterVtorFieldTbloffShift)
}

// SetTbloff Vector table base offset field
func (r *registerVtorType) SetTbloff(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterVtorFieldTbloffMask)|(uint32(value)<<RegisterVtorFieldTbloffShift))
}

// registerAircrType Application interrupt and reset control register
type registerAircrType uint32

const (
	RegisterAircrFieldVectresetShift = 0
	RegisterAircrFieldVectresetMask  = 0x1
)

// GetVectreset VECTRESET
func (r *registerAircrType) GetVectreset() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldVectresetMask) != 0
}

// SetVectreset VECTRESET
func (r *registerAircrType) SetVectreset(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAircrFieldVectresetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldVectresetMask)
	}
}

const (
	RegisterAircrFieldVectclractiveShift = 1
	RegisterAircrFieldVectclractiveMask  = 0x2
)

// GetVectclractive VECTCLRACTIVE
func (r *registerAircrType) GetVectclractive() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldVectclractiveMask) != 0
}

// SetVectclractive VECTCLRACTIVE
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

// GetSysresetreq SYSRESETREQ
func (r *registerAircrType) GetSysresetreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldSysresetreqMask) != 0
}

// SetSysresetreq SYSRESETREQ
func (r *registerAircrType) SetSysresetreq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAircrFieldSysresetreqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldSysresetreqMask)
	}
}

const (
	RegisterAircrFieldPrigroupShift = 8
	RegisterAircrFieldPrigroupMask  = 0x700
)

// GetPrigroup PRIGROUP
func (r *registerAircrType) GetPrigroup() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldPrigroupMask) >> RegisterAircrFieldPrigroupShift)
}

// SetPrigroup PRIGROUP
func (r *registerAircrType) SetPrigroup(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldPrigroupMask)|(uint32(value)<<RegisterAircrFieldPrigroupShift))
}

const (
	RegisterAircrFieldEndianessShift = 15
	RegisterAircrFieldEndianessMask  = 0x8000
)

// GetEndianess ENDIANESS
func (r *registerAircrType) GetEndianess() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldEndianessMask) != 0
}

// SetEndianess ENDIANESS
func (r *registerAircrType) SetEndianess(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAircrFieldEndianessMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldEndianessMask)
	}
}

const (
	RegisterAircrFieldVectkeystatShift = 16
	RegisterAircrFieldVectkeystatMask  = 0xffff0000
)

// GetVectkeystat Register key
func (r *registerAircrType) GetVectkeystat() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterAircrFieldVectkeystatMask) >> RegisterAircrFieldVectkeystatShift)
}

// SetVectkeystat Register key
func (r *registerAircrType) SetVectkeystat(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAircrFieldVectkeystatMask)|(uint32(value)<<RegisterAircrFieldVectkeystatShift))
}

// registerScrType System control register
type registerScrType uint32

const (
	RegisterScrFieldSleeponexitShift = 1
	RegisterScrFieldSleeponexitMask  = 0x2
)

// GetSleeponexit SLEEPONEXIT
func (r *registerScrType) GetSleeponexit() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScrFieldSleeponexitMask) != 0
}

// SetSleeponexit SLEEPONEXIT
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

// GetSleepdeep SLEEPDEEP
func (r *registerScrType) GetSleepdeep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScrFieldSleepdeepMask) != 0
}

// SetSleepdeep SLEEPDEEP
func (r *registerScrType) SetSleepdeep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScrFieldSleepdeepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScrFieldSleepdeepMask)
	}
}

const (
	RegisterScrFieldSeveonpendShift = 4
	RegisterScrFieldSeveonpendMask  = 0x10
)

// GetSeveonpend Send Event on Pending bit
func (r *registerScrType) GetSeveonpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScrFieldSeveonpendMask) != 0
}

// SetSeveonpend Send Event on Pending bit
func (r *registerScrType) SetSeveonpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScrFieldSeveonpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScrFieldSeveonpendMask)
	}
}

// registerCcrType Configuration and control register
type registerCcrType uint32

const (
	RegisterCcrFieldNonbasethrdenaShift = 0
	RegisterCcrFieldNonbasethrdenaMask  = 0x1
)

// GetNonbasethrdena Configures how the processor enters Thread mode
func (r *registerCcrType) GetNonbasethrdena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldNonbasethrdenaMask) != 0
}

// SetNonbasethrdena Configures how the processor enters Thread mode
func (r *registerCcrType) SetNonbasethrdena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldNonbasethrdenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldNonbasethrdenaMask)
	}
}

const (
	RegisterCcrFieldUsersetmpendShift = 1
	RegisterCcrFieldUsersetmpendMask  = 0x2
)

// GetUsersetmpend USERSETMPEND
func (r *registerCcrType) GetUsersetmpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldUsersetmpendMask) != 0
}

// SetUsersetmpend USERSETMPEND
func (r *registerCcrType) SetUsersetmpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldUsersetmpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldUsersetmpendMask)
	}
}

const (
	RegisterCcrFieldUnalign__trpShift = 3
	RegisterCcrFieldUnalign__trpMask  = 0x8
)

// GetUnalign__trp UNALIGN_ TRP
func (r *registerCcrType) GetUnalign__trp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldUnalign__trpMask) != 0
}

// SetUnalign__trp UNALIGN_ TRP
func (r *registerCcrType) SetUnalign__trp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldUnalign__trpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldUnalign__trpMask)
	}
}

const (
	RegisterCcrFieldDiv_0_trpShift = 4
	RegisterCcrFieldDiv_0_trpMask  = 0x10
)

// GetDiv_0_trp DIV_0_TRP
func (r *registerCcrType) GetDiv_0_trp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDiv_0_trpMask) != 0
}

// SetDiv_0_trp DIV_0_TRP
func (r *registerCcrType) SetDiv_0_trp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldDiv_0_trpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDiv_0_trpMask)
	}
}

const (
	RegisterCcrFieldBfhfnmignShift = 8
	RegisterCcrFieldBfhfnmignMask  = 0x100
)

// GetBfhfnmign BFHFNMIGN
func (r *registerCcrType) GetBfhfnmign() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldBfhfnmignMask) != 0
}

// SetBfhfnmign BFHFNMIGN
func (r *registerCcrType) SetBfhfnmign(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldBfhfnmignMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldBfhfnmignMask)
	}
}

const (
	RegisterCcrFieldStkalignShift = 9
	RegisterCcrFieldStkalignMask  = 0x200
)

// GetStkalign STKALIGN
func (r *registerCcrType) GetStkalign() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldStkalignMask) != 0
}

// SetStkalign STKALIGN
func (r *registerCcrType) SetStkalign(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldStkalignMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldStkalignMask)
	}
}

const (
	RegisterCcrFieldDcShift = 16
	RegisterCcrFieldDcMask  = 0x10000
)

// GetDc DC
func (r *registerCcrType) GetDc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDcMask) != 0
}

// SetDc DC
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

// GetIc IC
func (r *registerCcrType) GetIc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldIcMask) != 0
}

// SetIc IC
func (r *registerCcrType) SetIc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldIcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldIcMask)
	}
}

const (
	RegisterCcrFieldBpShift = 18
	RegisterCcrFieldBpMask  = 0x40000
)

// GetBp BP
func (r *registerCcrType) GetBp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldBpMask) != 0
}

// SetBp BP
func (r *registerCcrType) SetBp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldBpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldBpMask)
	}
}

// registerShpr1Type System handler priority registers
type registerShpr1Type uint32

const (
	RegisterShpr1FieldPri_4Shift = 0
	RegisterShpr1FieldPri_4Mask  = 0xff
)

// GetPri_4 Priority of system handler 4
func (r *registerShpr1Type) GetPri_4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr1FieldPri_4Mask) >> RegisterShpr1FieldPri_4Shift)
}

// SetPri_4 Priority of system handler 4
func (r *registerShpr1Type) SetPri_4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr1FieldPri_4Mask)|(uint32(value)<<RegisterShpr1FieldPri_4Shift))
}

const (
	RegisterShpr1FieldPri_5Shift = 8
	RegisterShpr1FieldPri_5Mask  = 0xff00
)

// GetPri_5 Priority of system handler 5
func (r *registerShpr1Type) GetPri_5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr1FieldPri_5Mask) >> RegisterShpr1FieldPri_5Shift)
}

// SetPri_5 Priority of system handler 5
func (r *registerShpr1Type) SetPri_5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr1FieldPri_5Mask)|(uint32(value)<<RegisterShpr1FieldPri_5Shift))
}

const (
	RegisterShpr1FieldPri_6Shift = 16
	RegisterShpr1FieldPri_6Mask  = 0xff0000
)

// GetPri_6 Priority of system handler 6
func (r *registerShpr1Type) GetPri_6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr1FieldPri_6Mask) >> RegisterShpr1FieldPri_6Shift)
}

// SetPri_6 Priority of system handler 6
func (r *registerShpr1Type) SetPri_6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr1FieldPri_6Mask)|(uint32(value)<<RegisterShpr1FieldPri_6Shift))
}

// registerShpr2Type System handler priority registers
type registerShpr2Type uint32

const (
	RegisterShpr2FieldPri_11Shift = 24
	RegisterShpr2FieldPri_11Mask  = 0xff000000
)

// GetPri_11 Priority of system handler 11
func (r *registerShpr2Type) GetPri_11() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr2FieldPri_11Mask) >> RegisterShpr2FieldPri_11Shift)
}

// SetPri_11 Priority of system handler 11
func (r *registerShpr2Type) SetPri_11(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr2FieldPri_11Mask)|(uint32(value)<<RegisterShpr2FieldPri_11Shift))
}

// registerShpr3Type System handler priority registers
type registerShpr3Type uint32

const (
	RegisterShpr3FieldPri_14Shift = 16
	RegisterShpr3FieldPri_14Mask  = 0xff0000
)

// GetPri_14 Priority of system handler 14
func (r *registerShpr3Type) GetPri_14() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr3FieldPri_14Mask) >> RegisterShpr3FieldPri_14Shift)
}

// SetPri_14 Priority of system handler 14
func (r *registerShpr3Type) SetPri_14(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr3FieldPri_14Mask)|(uint32(value)<<RegisterShpr3FieldPri_14Shift))
}

const (
	RegisterShpr3FieldPri_15Shift = 24
	RegisterShpr3FieldPri_15Mask  = 0xff000000
)

// GetPri_15 Priority of system handler 15
func (r *registerShpr3Type) GetPri_15() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShpr3FieldPri_15Mask) >> RegisterShpr3FieldPri_15Shift)
}

// SetPri_15 Priority of system handler 15
func (r *registerShpr3Type) SetPri_15(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShpr3FieldPri_15Mask)|(uint32(value)<<RegisterShpr3FieldPri_15Shift))
}

// registerShcrsType System handler control and state register
type registerShcrsType uint32

const (
	RegisterShcrsFieldMemfaultactShift = 0
	RegisterShcrsFieldMemfaultactMask  = 0x1
)

// GetMemfaultact Memory management fault exception active bit
func (r *registerShcrsType) GetMemfaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldMemfaultactMask) != 0
}

// SetMemfaultact Memory management fault exception active bit
func (r *registerShcrsType) SetMemfaultact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldMemfaultactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldMemfaultactMask)
	}
}

const (
	RegisterShcrsFieldBusfaultactShift = 1
	RegisterShcrsFieldBusfaultactMask  = 0x2
)

// GetBusfaultact Bus fault exception active bit
func (r *registerShcrsType) GetBusfaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldBusfaultactMask) != 0
}

// SetBusfaultact Bus fault exception active bit
func (r *registerShcrsType) SetBusfaultact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldBusfaultactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldBusfaultactMask)
	}
}

const (
	RegisterShcrsFieldUsgfaultactShift = 3
	RegisterShcrsFieldUsgfaultactMask  = 0x8
)

// GetUsgfaultact Usage fault exception active bit
func (r *registerShcrsType) GetUsgfaultact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldUsgfaultactMask) != 0
}

// SetUsgfaultact Usage fault exception active bit
func (r *registerShcrsType) SetUsgfaultact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldUsgfaultactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldUsgfaultactMask)
	}
}

const (
	RegisterShcrsFieldSvcallactShift = 7
	RegisterShcrsFieldSvcallactMask  = 0x80
)

// GetSvcallact SVC call active bit
func (r *registerShcrsType) GetSvcallact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldSvcallactMask) != 0
}

// SetSvcallact SVC call active bit
func (r *registerShcrsType) SetSvcallact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldSvcallactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldSvcallactMask)
	}
}

const (
	RegisterShcrsFieldMonitoractShift = 8
	RegisterShcrsFieldMonitoractMask  = 0x100
)

// GetMonitoract Debug monitor active bit
func (r *registerShcrsType) GetMonitoract() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldMonitoractMask) != 0
}

// SetMonitoract Debug monitor active bit
func (r *registerShcrsType) SetMonitoract(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldMonitoractMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldMonitoractMask)
	}
}

const (
	RegisterShcrsFieldPendsvactShift = 10
	RegisterShcrsFieldPendsvactMask  = 0x400
)

// GetPendsvact PendSV exception active bit
func (r *registerShcrsType) GetPendsvact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldPendsvactMask) != 0
}

// SetPendsvact PendSV exception active bit
func (r *registerShcrsType) SetPendsvact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldPendsvactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldPendsvactMask)
	}
}

const (
	RegisterShcrsFieldSystickactShift = 11
	RegisterShcrsFieldSystickactMask  = 0x800
)

// GetSystickact SysTick exception active bit
func (r *registerShcrsType) GetSystickact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldSystickactMask) != 0
}

// SetSystickact SysTick exception active bit
func (r *registerShcrsType) SetSystickact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldSystickactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldSystickactMask)
	}
}

const (
	RegisterShcrsFieldUsgfaultpendedShift = 12
	RegisterShcrsFieldUsgfaultpendedMask  = 0x1000
)

// GetUsgfaultpended Usage fault exception pending bit
func (r *registerShcrsType) GetUsgfaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldUsgfaultpendedMask) != 0
}

// SetUsgfaultpended Usage fault exception pending bit
func (r *registerShcrsType) SetUsgfaultpended(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldUsgfaultpendedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldUsgfaultpendedMask)
	}
}

const (
	RegisterShcrsFieldMemfaultpendedShift = 13
	RegisterShcrsFieldMemfaultpendedMask  = 0x2000
)

// GetMemfaultpended Memory management fault exception pending bit
func (r *registerShcrsType) GetMemfaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldMemfaultpendedMask) != 0
}

// SetMemfaultpended Memory management fault exception pending bit
func (r *registerShcrsType) SetMemfaultpended(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldMemfaultpendedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldMemfaultpendedMask)
	}
}

const (
	RegisterShcrsFieldBusfaultpendedShift = 14
	RegisterShcrsFieldBusfaultpendedMask  = 0x4000
)

// GetBusfaultpended Bus fault exception pending bit
func (r *registerShcrsType) GetBusfaultpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldBusfaultpendedMask) != 0
}

// SetBusfaultpended Bus fault exception pending bit
func (r *registerShcrsType) SetBusfaultpended(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldBusfaultpendedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldBusfaultpendedMask)
	}
}

const (
	RegisterShcrsFieldSvcallpendedShift = 15
	RegisterShcrsFieldSvcallpendedMask  = 0x8000
)

// GetSvcallpended SVC call pending bit
func (r *registerShcrsType) GetSvcallpended() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldSvcallpendedMask) != 0
}

// SetSvcallpended SVC call pending bit
func (r *registerShcrsType) SetSvcallpended(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldSvcallpendedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldSvcallpendedMask)
	}
}

const (
	RegisterShcrsFieldMemfaultenaShift = 16
	RegisterShcrsFieldMemfaultenaMask  = 0x10000
)

// GetMemfaultena Memory management fault enable bit
func (r *registerShcrsType) GetMemfaultena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldMemfaultenaMask) != 0
}

// SetMemfaultena Memory management fault enable bit
func (r *registerShcrsType) SetMemfaultena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldMemfaultenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldMemfaultenaMask)
	}
}

const (
	RegisterShcrsFieldBusfaultenaShift = 17
	RegisterShcrsFieldBusfaultenaMask  = 0x20000
)

// GetBusfaultena Bus fault enable bit
func (r *registerShcrsType) GetBusfaultena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldBusfaultenaMask) != 0
}

// SetBusfaultena Bus fault enable bit
func (r *registerShcrsType) SetBusfaultena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldBusfaultenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldBusfaultenaMask)
	}
}

const (
	RegisterShcrsFieldUsgfaultenaShift = 18
	RegisterShcrsFieldUsgfaultenaMask  = 0x40000
)

// GetUsgfaultena Usage fault enable bit
func (r *registerShcrsType) GetUsgfaultena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterShcrsFieldUsgfaultenaMask) != 0
}

// SetUsgfaultena Usage fault enable bit
func (r *registerShcrsType) SetUsgfaultena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterShcrsFieldUsgfaultenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterShcrsFieldUsgfaultenaMask)
	}
}

// registerCfsr_ufsr_bfsr_mmfsrType Configurable fault status register
type registerCfsr_ufsr_bfsr_mmfsrType uint32

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldIaccviolShift = 0
	RegisterCfsr_ufsr_bfsr_mmfsrFieldIaccviolMask  = 0x1
)

// GetIaccviol IACCVIOL
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetIaccviol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldIaccviolMask) != 0
}

// SetIaccviol IACCVIOL
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetIaccviol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldIaccviolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldIaccviolMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldDaccviolShift = 1
	RegisterCfsr_ufsr_bfsr_mmfsrFieldDaccviolMask  = 0x2
)

// GetDaccviol DACCVIOL
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetDaccviol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldDaccviolMask) != 0
}

// SetDaccviol DACCVIOL
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetDaccviol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldDaccviolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldDaccviolMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldMunstkerrShift = 3
	RegisterCfsr_ufsr_bfsr_mmfsrFieldMunstkerrMask  = 0x8
)

// GetMunstkerr MUNSTKERR
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetMunstkerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldMunstkerrMask) != 0
}

// SetMunstkerr MUNSTKERR
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetMunstkerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldMunstkerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldMunstkerrMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldMstkerrShift = 4
	RegisterCfsr_ufsr_bfsr_mmfsrFieldMstkerrMask  = 0x10
)

// GetMstkerr MSTKERR
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetMstkerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldMstkerrMask) != 0
}

// SetMstkerr MSTKERR
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetMstkerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldMstkerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldMstkerrMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldMlsperrShift = 5
	RegisterCfsr_ufsr_bfsr_mmfsrFieldMlsperrMask  = 0x20
)

// GetMlsperr MLSPERR
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetMlsperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldMlsperrMask) != 0
}

// SetMlsperr MLSPERR
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetMlsperr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldMlsperrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldMlsperrMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldMmarvalidShift = 7
	RegisterCfsr_ufsr_bfsr_mmfsrFieldMmarvalidMask  = 0x80
)

// GetMmarvalid MMARVALID
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetMmarvalid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldMmarvalidMask) != 0
}

// SetMmarvalid MMARVALID
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetMmarvalid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldMmarvalidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldMmarvalidMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldIbuserrShift = 8
	RegisterCfsr_ufsr_bfsr_mmfsrFieldIbuserrMask  = 0x100
)

// GetIbuserr Instruction bus error
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetIbuserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldIbuserrMask) != 0
}

// SetIbuserr Instruction bus error
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetIbuserr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldIbuserrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldIbuserrMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldPreciserrShift = 9
	RegisterCfsr_ufsr_bfsr_mmfsrFieldPreciserrMask  = 0x200
)

// GetPreciserr Precise data bus error
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetPreciserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldPreciserrMask) != 0
}

// SetPreciserr Precise data bus error
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetPreciserr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldPreciserrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldPreciserrMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldImpreciserrShift = 10
	RegisterCfsr_ufsr_bfsr_mmfsrFieldImpreciserrMask  = 0x400
)

// GetImpreciserr Imprecise data bus error
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetImpreciserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldImpreciserrMask) != 0
}

// SetImpreciserr Imprecise data bus error
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetImpreciserr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldImpreciserrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldImpreciserrMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldUnstkerrShift = 11
	RegisterCfsr_ufsr_bfsr_mmfsrFieldUnstkerrMask  = 0x800
)

// GetUnstkerr Bus fault on unstacking for a return from exception
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetUnstkerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldUnstkerrMask) != 0
}

// SetUnstkerr Bus fault on unstacking for a return from exception
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetUnstkerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldUnstkerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldUnstkerrMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldStkerrShift = 12
	RegisterCfsr_ufsr_bfsr_mmfsrFieldStkerrMask  = 0x1000
)

// GetStkerr Bus fault on stacking for exception entry
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetStkerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldStkerrMask) != 0
}

// SetStkerr Bus fault on stacking for exception entry
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetStkerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldStkerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldStkerrMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldLsperrShift = 13
	RegisterCfsr_ufsr_bfsr_mmfsrFieldLsperrMask  = 0x2000
)

// GetLsperr Bus fault on floating-point lazy state preservation
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetLsperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldLsperrMask) != 0
}

// SetLsperr Bus fault on floating-point lazy state preservation
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetLsperr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldLsperrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldLsperrMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldBfarvalidShift = 15
	RegisterCfsr_ufsr_bfsr_mmfsrFieldBfarvalidMask  = 0x8000
)

// GetBfarvalid Bus Fault Address Register (BFAR) valid flag
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetBfarvalid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldBfarvalidMask) != 0
}

// SetBfarvalid Bus Fault Address Register (BFAR) valid flag
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetBfarvalid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldBfarvalidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldBfarvalidMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldUndefinstrShift = 16
	RegisterCfsr_ufsr_bfsr_mmfsrFieldUndefinstrMask  = 0x10000
)

// GetUndefinstr Undefined instruction usage fault
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetUndefinstr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldUndefinstrMask) != 0
}

// SetUndefinstr Undefined instruction usage fault
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetUndefinstr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldUndefinstrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldUndefinstrMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldInvstateShift = 17
	RegisterCfsr_ufsr_bfsr_mmfsrFieldInvstateMask  = 0x20000
)

// GetInvstate Invalid state usage fault
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetInvstate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldInvstateMask) != 0
}

// SetInvstate Invalid state usage fault
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetInvstate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldInvstateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldInvstateMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldInvpcShift = 18
	RegisterCfsr_ufsr_bfsr_mmfsrFieldInvpcMask  = 0x40000
)

// GetInvpc Invalid PC load usage fault
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetInvpc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldInvpcMask) != 0
}

// SetInvpc Invalid PC load usage fault
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetInvpc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldInvpcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldInvpcMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldNocpShift = 19
	RegisterCfsr_ufsr_bfsr_mmfsrFieldNocpMask  = 0x80000
)

// GetNocp No coprocessor usage fault.
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetNocp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldNocpMask) != 0
}

// SetNocp No coprocessor usage fault.
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetNocp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldNocpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldNocpMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldUnalignedShift = 24
	RegisterCfsr_ufsr_bfsr_mmfsrFieldUnalignedMask  = 0x1000000
)

// GetUnaligned Unaligned access usage fault
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetUnaligned() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldUnalignedMask) != 0
}

// SetUnaligned Unaligned access usage fault
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetUnaligned(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldUnalignedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldUnalignedMask)
	}
}

const (
	RegisterCfsr_ufsr_bfsr_mmfsrFieldDivbyzeroShift = 25
	RegisterCfsr_ufsr_bfsr_mmfsrFieldDivbyzeroMask  = 0x2000000
)

// GetDivbyzero Divide by zero usage fault
func (r *registerCfsr_ufsr_bfsr_mmfsrType) GetDivbyzero() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfsr_ufsr_bfsr_mmfsrFieldDivbyzeroMask) != 0
}

// SetDivbyzero Divide by zero usage fault
func (r *registerCfsr_ufsr_bfsr_mmfsrType) SetDivbyzero(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfsr_ufsr_bfsr_mmfsrFieldDivbyzeroMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfsr_ufsr_bfsr_mmfsrFieldDivbyzeroMask)
	}
}

// registerHfsrType Hard fault status register
type registerHfsrType uint32

const (
	RegisterHfsrFieldVecttblShift = 1
	RegisterHfsrFieldVecttblMask  = 0x2
)

// GetVecttbl Vector table hard fault
func (r *registerHfsrType) GetVecttbl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHfsrFieldVecttblMask) != 0
}

// SetVecttbl Vector table hard fault
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

// GetForced Forced hard fault
func (r *registerHfsrType) GetForced() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHfsrFieldForcedMask) != 0
}

// SetForced Forced hard fault
func (r *registerHfsrType) SetForced(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHfsrFieldForcedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHfsrFieldForcedMask)
	}
}

const (
	RegisterHfsrFieldDebug_vtShift = 31
	RegisterHfsrFieldDebug_vtMask  = 0x80000000
)

// GetDebug_vt Reserved for Debug use
func (r *registerHfsrType) GetDebug_vt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHfsrFieldDebug_vtMask) != 0
}

// SetDebug_vt Reserved for Debug use
func (r *registerHfsrType) SetDebug_vt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHfsrFieldDebug_vtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHfsrFieldDebug_vtMask)
	}
}

// registerMmfarType Memory management fault address register
type registerMmfarType uint32

const (
	RegisterMmfarFieldAddressShift = 0
	RegisterMmfarFieldAddressMask  = 0xffffffff
)

// GetAddress Memory management fault address
func (r *registerMmfarType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMmfarFieldAddressMask) >> RegisterMmfarFieldAddressShift)
}

// SetAddress Memory management fault address
func (r *registerMmfarType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMmfarFieldAddressMask)|(uint32(value)<<RegisterMmfarFieldAddressShift))
}

// registerBfarType Bus fault address register
type registerBfarType uint32

const (
	RegisterBfarFieldAddressShift = 0
	RegisterBfarFieldAddressMask  = 0xffffffff
)

// GetAddress Bus fault address
func (r *registerBfarType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterBfarFieldAddressMask) >> RegisterBfarFieldAddressShift)
}

// SetAddress Bus fault address
func (r *registerBfarType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBfarFieldAddressMask)|(uint32(value)<<RegisterBfarFieldAddressShift))
}
