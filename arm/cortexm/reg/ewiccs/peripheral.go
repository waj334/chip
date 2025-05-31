//go:build arm && cortexm

package ewiccs

import (
	"unsafe"
	"volatile"
)

var (
	Ewiccs = (*_ewiccs)(unsafe.Pointer(uintptr(0xe0047f00)))
)

type _ewiccs struct {
	Itctrl     registerItctrlType
	_          [156]byte
	Claimset   registerClaimsetType
	Claimclr   registerClaimclrType
	Devaff0    registerDevaff0Type
	Devaff1    registerDevaff1Type
	Lar        registerLarType
	Lsr        registerLsrType
	Authstatus registerAuthstatusType
	Devarch    registerDevarchType
	Devid2     registerDevid2Type
	Devid1     registerDevid1Type
	Devid      registerDevidType
	Devtype    registerDevtypeType
	Pidr4      registerPidr4Type
	Pidr5      registerPidr5Type
	Pidr6      registerPidr6Type
	Pidr7      registerPidr7Type
	Pidr0      registerPidr0Type
	Pidr1      registerPidr1Type
	Pidr2      registerPidr2Type
	Pidr3      registerPidr3Type
	Cidr0      registerCidr0Type
	Cidr1      registerCidr1Type
	Cidr2      registerCidr2Type
	Cidr3      registerCidr3Type
}

// registerItctrlType Used to dynamically switch between functional mode and integration mode. The EWIC does not support integration mode, and this register is RAZ.
type registerItctrlType uint32

// registerClaimsetType The EWIC_CLAIMSET register is used to set whether functionality is in use by a debug agent. The EWIC does not have any associated debug functionality, and this register is 0xF.
type registerClaimsetType uint32

// registerClaimclrType The EWIC_CLAIMCLR register is used to set whether functionality is in use by a debug agent. The EWIC does not have any associated debug functionality, and this register is 0x0.
type registerClaimclrType uint32

// registerDevaff0Type The EWIC_DEVAFF0 register enables a debugger to determine whether the EWIC and the processor have an affinity with each other. The EWIC does not have any associated debug functionality, and this register is 0x80000000.
type registerDevaff0Type uint32

// registerDevaff1Type The EWIC_DEVAFF1 register enables a debugger to determine whether the EWIC and the processor have an affinity with each other. The EWIC does not have any associated debug functionality, and this register is 0x0.
type registerDevaff1Type uint32

// registerLarType The EWIC_LAR register controls software access to CoreSight components to reduce the likelihood of accidental access to the EWIC. The EWIC does not support software locking, and writing to this register has no affect.
type registerLarType uint32

// registerLsrType The EWIC_LSR register controls software access to CoreSight components to reduce the likelihood of accidental access to the EWIC. The EWIC does not support software locking, and this register is RAZ.
type registerLsrType uint32

// registerAuthstatusType The EWIC_AUTHSTATUS register indicates whether the EWIC includes debug functionality. The EWIC is not a debug component, therefore, the EWIC_AUTHSTATIS register is RAZ.
type registerAuthstatusType uint32

// registerDevarchType Identifies the architecture of the EWIC
type registerDevarchType uint32

const (
	RegisterDevarchFieldArchidShift = 0
	RegisterDevarchFieldArchidMask  = 0xffff
)

// GetArchid Returns a value that identifies the architecture of the component. This value is 0x0A07 corresponding to EWIC architecture.
func (r *registerDevarchType) GetArchid() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDevarchFieldArchidMask) >> RegisterDevarchFieldArchidShift)
}

// SetArchid Returns a value that identifies the architecture of the component. This value is 0x0A07 corresponding to EWIC architecture.
func (r *registerDevarchType) SetArchid(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDevarchFieldArchidMask)|(uint32(value)<<RegisterDevarchFieldArchidShift))
}

const (
	RegisterDevarchFieldRevisionShift = 16
	RegisterDevarchFieldRevisionMask  = 0xf0000
)

// GetRevision Architecture revision. This value is 0b0000.
func (r *registerDevarchType) GetRevision() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDevarchFieldRevisionMask) >> RegisterDevarchFieldRevisionShift)
}

// SetRevision Architecture revision. This value is 0b0000.
func (r *registerDevarchType) SetRevision(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDevarchFieldRevisionMask)|(uint32(value)<<RegisterDevarchFieldRevisionShift))
}

const (
	RegisterDevarchFieldPresentShift = 20
	RegisterDevarchFieldPresentMask  = 0x100000
)

// GetPresent Indicates the presence of this register. This value is 0b1.
func (r *registerDevarchType) GetPresent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDevarchFieldPresentMask) != 0
}

// SetPresent Indicates the presence of this register. This value is 0b1.
func (r *registerDevarchType) SetPresent(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDevarchFieldPresentMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDevarchFieldPresentMask)
	}
}

const (
	RegisterDevarchFieldArchitectjep106iShift = 21
	RegisterDevarchFieldArchitectjep106iMask  = 0xfe00000
)

// GetArchitectjep106i Defines the architect of the component, JEP106 identification code
func (r *registerDevarchType) GetArchitectjep106i() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDevarchFieldArchitectjep106iMask) >> RegisterDevarchFieldArchitectjep106iShift)
}

// SetArchitectjep106i Defines the architect of the component, JEP106 identification code
func (r *registerDevarchType) SetArchitectjep106i(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDevarchFieldArchitectjep106iMask)|(uint32(value)<<RegisterDevarchFieldArchitectjep106iShift))
}

const (
	RegisterDevarchFieldArchitectjep106cShift = 28
	RegisterDevarchFieldArchitectjep106cMask  = 0xf0000000
)

// GetArchitectjep106c Defines the architect of the component, JEP106 continuation code
func (r *registerDevarchType) GetArchitectjep106c() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDevarchFieldArchitectjep106cMask) >> RegisterDevarchFieldArchitectjep106cShift)
}

// SetArchitectjep106c Defines the architect of the component, JEP106 continuation code
func (r *registerDevarchType) SetArchitectjep106c(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDevarchFieldArchitectjep106cMask)|(uint32(value)<<RegisterDevarchFieldArchitectjep106cShift))
}

// registerDevid2Type The EWIC_DEVID2 indicates some capabilities of the EWIC. This register is RAZ/WI.
type registerDevid2Type uint32

// registerDevid1Type The EWIC_DEVID1 indicates some capabilities of the EWIC. This register is RAZ/WI.
type registerDevid1Type uint32

// registerDevidType The EWIC_DEVID indicates some capabilities of the EWIC. This register is RAZ/WI.
type registerDevidType uint32

// registerDevtypeType Provides part number information about the EWIC component to the debugger.
type registerDevtypeType uint32

const (
	RegisterDevtypeFieldMajorShift = 0
	RegisterDevtypeFieldMajorMask  = 0xf
)

// GetMajor Major type for the component device type. This field is set to 0b0000.
func (r *registerDevtypeType) GetMajor() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDevtypeFieldMajorMask) >> RegisterDevtypeFieldMajorShift)
}

// SetMajor Major type for the component device type. This field is set to 0b0000.
func (r *registerDevtypeType) SetMajor(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDevtypeFieldMajorMask)|(uint32(value)<<RegisterDevtypeFieldMajorShift))
}

const (
	RegisterDevtypeFieldSubShift = 4
	RegisterDevtypeFieldSubMask  = 0xf0
)

// GetSub Sub type for the component device type. This field is set to 0b0000.
func (r *registerDevtypeType) GetSub() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDevtypeFieldSubMask) >> RegisterDevtypeFieldSubShift)
}

// SetSub Sub type for the component device type. This field is set to 0b0000.
func (r *registerDevtypeType) SetSub(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDevtypeFieldSubMask)|(uint32(value)<<RegisterDevtypeFieldSubShift))
}

// registerPidr4Type provides information about the memory size and JEP106 continuation code that the External Wakeup Interrupt Controller (EWIC) component uses
type registerPidr4Type uint32

const (
	RegisterPidr4FieldDes2Shift = 0
	RegisterPidr4FieldDes2Mask  = 0xf
)

// GetDes2 JEP106 continuation code. Together with EWIC_PIDR2.DES_1 and EWIC_PIDR1.DES_0, they indicate the designer of the component, not the implementer, except where the two are the same. The reset value of this field is 0x4.
func (r *registerPidr4Type) GetDes2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr4FieldDes2Mask) >> RegisterPidr4FieldDes2Shift)
}

// SetDes2 JEP106 continuation code. Together with EWIC_PIDR2.DES_1 and EWIC_PIDR1.DES_0, they indicate the designer of the component, not the implementer, except where the two are the same. The reset value of this field is 0x4.
func (r *registerPidr4Type) SetDes2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr4FieldDes2Mask)|(uint32(value)<<RegisterPidr4FieldDes2Shift))
}

const (
	RegisterPidr4FieldSizeShift = 4
	RegisterPidr4FieldSizeMask  = 0xf0
)

// GetSize This field indicates the memory size that the EWIC uses. This field returns 0x0 indicating that the component uses an UNKNOWN number of 4KB blocks.
func (r *registerPidr4Type) GetSize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr4FieldSizeMask) >> RegisterPidr4FieldSizeShift)
}

// SetSize This field indicates the memory size that the EWIC uses. This field returns 0x0 indicating that the component uses an UNKNOWN number of 4KB blocks.
func (r *registerPidr4Type) SetSize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr4FieldSizeMask)|(uint32(value)<<RegisterPidr4FieldSizeShift))
}

// registerPidr5Type The register is reserved, RES0
type registerPidr5Type uint32

// registerPidr6Type The register is reserved, RES0
type registerPidr6Type uint32

// registerPidr7Type The register is reserved, RES0
type registerPidr7Type uint32

// registerPidr0Type Indicates the EWIC component part number
type registerPidr0Type uint32

const (
	RegisterPidr0FieldPart0Shift = 0
	RegisterPidr0FieldPart0Mask  = 0xff
)

// GetPart0 This field indicates the part number. When taken together with EWIC_PIDR1.PART_1, it indicates the component. The part number is selected by the designer of the component.
func (r *registerPidr0Type) GetPart0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr0FieldPart0Mask) >> RegisterPidr0FieldPart0Shift)
}

// SetPart0 This field indicates the part number. When taken together with EWIC_PIDR1.PART_1, it indicates the component. The part number is selected by the designer of the component.
func (r *registerPidr0Type) SetPart0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr0FieldPart0Mask)|(uint32(value)<<RegisterPidr0FieldPart0Shift))
}

// registerPidr1Type Indicates the EWIC component JEP106 continuation code and part number
type registerPidr1Type uint32

const (
	RegisterPidr1FieldPart1Shift = 0
	RegisterPidr1FieldPart1Mask  = 0xf
)

// GetPart1 This field indicates the part number, bits[11:8]. Taken together with EWIC_PIDR0.PART_0 it indicates the component. The part number is selected by the designer of the component. The reset value is 0xD.
func (r *registerPidr1Type) GetPart1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr1FieldPart1Mask) >> RegisterPidr1FieldPart1Shift)
}

// SetPart1 This field indicates the part number, bits[11:8]. Taken together with EWIC_PIDR0.PART_0 it indicates the component. The part number is selected by the designer of the component. The reset value is 0xD.
func (r *registerPidr1Type) SetPart1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr1FieldPart1Mask)|(uint32(value)<<RegisterPidr1FieldPart1Shift))
}

const (
	RegisterPidr1FieldDes0Shift = 4
	RegisterPidr1FieldDes0Mask  = 0xf0
)

// GetDes0 This field indicates the JEP106 identification code, bits[3:0]. Together, with EWIC_PIDR4.DES_2 and EWIC_PIDR2.DES_1, they indicate the designer of the component and not the implementer, except where the two are the same. The reset value is 0xB.
func (r *registerPidr1Type) GetDes0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr1FieldDes0Mask) >> RegisterPidr1FieldDes0Shift)
}

// SetDes0 This field indicates the JEP106 identification code, bits[3:0]. Together, with EWIC_PIDR4.DES_2 and EWIC_PIDR2.DES_1, they indicate the designer of the component and not the implementer, except where the two are the same. The reset value is 0xB.
func (r *registerPidr1Type) SetDes0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr1FieldDes0Mask)|(uint32(value)<<RegisterPidr1FieldDes0Shift))
}

// registerPidr2Type Indicates the EWIC component revision number, JEDEC value, and part of the JEP106 continuation code
type registerPidr2Type uint32

const (
	RegisterPidr2FieldDes1Shift = 0
	RegisterPidr2FieldDes1Mask  = 0x7
)

// GetDes1 This field is the JEP106 identification code, bits[6:4]. Together, with CTI_PIDR4.DES_2 and CTI_PIDR1.DES_0, they indicate the designer of the component and not the implementer, except where the two are the same. The reset value is 0b011.
func (r *registerPidr2Type) GetDes1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr2FieldDes1Mask) >> RegisterPidr2FieldDes1Shift)
}

// SetDes1 This field is the JEP106 identification code, bits[6:4]. Together, with CTI_PIDR4.DES_2 and CTI_PIDR1.DES_0, they indicate the designer of the component and not the implementer, except where the two are the same. The reset value is 0b011.
func (r *registerPidr2Type) SetDes1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr2FieldDes1Mask)|(uint32(value)<<RegisterPidr2FieldDes1Shift))
}

const (
	RegisterPidr2FieldJedecShift = 3
	RegisterPidr2FieldJedecMask  = 0x8
)

// GetJedec This field is always 1, indicating that a JEDEC assigned value is used.
func (r *registerPidr2Type) GetJedec() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPidr2FieldJedecMask) != 0
}

// SetJedec This field is always 1, indicating that a JEDEC assigned value is used.
func (r *registerPidr2Type) SetJedec(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPidr2FieldJedecMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPidr2FieldJedecMask)
	}
}

const (
	RegisterPidr2FieldRevisionShift = 4
	RegisterPidr2FieldRevisionMask  = 0xf0
)

// GetRevision This field indicates the revision number of the EWIC component. It is an incremental value starting at 0x0 for the first design. The reset value is 0x0.
func (r *registerPidr2Type) GetRevision() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr2FieldRevisionMask) >> RegisterPidr2FieldRevisionShift)
}

// SetRevision This field indicates the revision number of the EWIC component. It is an incremental value starting at 0x0 for the first design. The reset value is 0x0.
func (r *registerPidr2Type) SetRevision(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr2FieldRevisionMask)|(uint32(value)<<RegisterPidr2FieldRevisionShift))
}

// registerPidr3Type Indicates minor errata fixes of the Cross Trigger Interface (CTI) component and if you have modified the behavior of the component
type registerPidr3Type uint32

const (
	RegisterPidr3FieldCmodShift = 0
	RegisterPidr3FieldCmodMask  = 0xf
)

// GetCmod Customer modified. Where the component is reusable IP, this value indicates whether you have modified the behavior of the component. In most cases, this field is 0x0.
func (r *registerPidr3Type) GetCmod() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr3FieldCmodMask) >> RegisterPidr3FieldCmodShift)
}

// SetCmod Customer modified. Where the component is reusable IP, this value indicates whether you have modified the behavior of the component. In most cases, this field is 0x0.
func (r *registerPidr3Type) SetCmod(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr3FieldCmodMask)|(uint32(value)<<RegisterPidr3FieldCmodShift))
}

const (
	RegisterPidr3FieldRevandShift = 4
	RegisterPidr3FieldRevandMask  = 0xf0
)

// GetRevand This field indicates minor errata fixes specific to this design, for example metal fixes after implementation. In most cases this field is 0x0.
func (r *registerPidr3Type) GetRevand() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr3FieldRevandMask) >> RegisterPidr3FieldRevandShift)
}

// SetRevand This field indicates minor errata fixes specific to this design, for example metal fixes after implementation. In most cases this field is 0x0.
func (r *registerPidr3Type) SetRevand(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr3FieldRevandMask)|(uint32(value)<<RegisterPidr3FieldRevandShift))
}

// registerCidr0Type Indicates the preamble
type registerCidr0Type uint32

const (
	RegisterCidr0FieldPrmbl0Shift = 0
	RegisterCidr0FieldPrmbl0Mask  = 0xff
)

// GetPrmbl0 Preamble. This field returns 0x0D.
func (r *registerCidr0Type) GetPrmbl0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCidr0FieldPrmbl0Mask) >> RegisterCidr0FieldPrmbl0Shift)
}

// SetPrmbl0 Preamble. This field returns 0x0D.
func (r *registerCidr0Type) SetPrmbl0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCidr0FieldPrmbl0Mask)|(uint32(value)<<RegisterCidr0FieldPrmbl0Shift))
}

// registerCidr1Type indicates the component class and preamble
type registerCidr1Type uint32

const (
	RegisterCidr1FieldClassShift = 0
	RegisterCidr1FieldClassMask  = 0xff
)

// GetClass Component class. Returns 0x9, indicating this is a CoreSight component.
func (r *registerCidr1Type) GetClass() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCidr1FieldClassMask) >> RegisterCidr1FieldClassShift)
}

// SetClass Component class. Returns 0x9, indicating this is a CoreSight component.
func (r *registerCidr1Type) SetClass(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCidr1FieldClassMask)|(uint32(value)<<RegisterCidr1FieldClassShift))
}

const (
	RegisterCidr1FieldPrmbl1Shift = 0
	RegisterCidr1FieldPrmbl1Mask  = 0xff
)

// GetPrmbl1 Preamble. This field returns 0x0.
func (r *registerCidr1Type) GetPrmbl1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCidr1FieldPrmbl1Mask) >> RegisterCidr1FieldPrmbl1Shift)
}

// SetPrmbl1 Preamble. This field returns 0x0.
func (r *registerCidr1Type) SetPrmbl1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCidr1FieldPrmbl1Mask)|(uint32(value)<<RegisterCidr1FieldPrmbl1Shift))
}

// registerCidr2Type Indicates the preamble
type registerCidr2Type uint32

const (
	RegisterCidr2FieldPrmbl2Shift = 0
	RegisterCidr2FieldPrmbl2Mask  = 0xff
)

// GetPrmbl2 Preamble. This field returns 0x05.
func (r *registerCidr2Type) GetPrmbl2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCidr2FieldPrmbl2Mask) >> RegisterCidr2FieldPrmbl2Shift)
}

// SetPrmbl2 Preamble. This field returns 0x05.
func (r *registerCidr2Type) SetPrmbl2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCidr2FieldPrmbl2Mask)|(uint32(value)<<RegisterCidr2FieldPrmbl2Shift))
}

// registerCidr3Type Indicates the preamble
type registerCidr3Type uint32

const (
	RegisterCidr3FieldPrmbl3Shift = 0
	RegisterCidr3FieldPrmbl3Mask  = 0xff
)

// GetPrmbl3 Preamble. This field returns 0xB1.
func (r *registerCidr3Type) GetPrmbl3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCidr3FieldPrmbl3Mask) >> RegisterCidr3FieldPrmbl3Shift)
}

// SetPrmbl3 Preamble. This field returns 0xB1.
func (r *registerCidr3Type) SetPrmbl3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCidr3FieldPrmbl3Mask)|(uint32(value)<<RegisterCidr3FieldPrmbl3Shift))
}
