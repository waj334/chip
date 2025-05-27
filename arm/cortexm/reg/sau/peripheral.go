//go:build arm && cortexm

package sau

import (
	"unsafe"
	"volatile"
)

var (
	Sau = (*_sau)(unsafe.Pointer(uintptr(0xe000edd0)))
)

type _sau struct {
	Ctrl registerCtrlType
	Type registerTypeType
	Rnr  registerRnrType
	Rbar registerRbarType
	Rlar registerRlarType
	Sfsr registerSfsrType
	Sfar registerSfarType
}

// registerCtrlType Allows enabling of the Security Attribution Unit (SAU)
type registerCtrlType uint32

const (
	RegisterCtrlFieldEnableShift = 0
	RegisterCtrlFieldEnableMask  = 0x1
)

// GetEnable Enables the (SAU)
func (r *registerCtrlType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldEnableMask) != 0
}

// SetEnable Enables the (SAU)
func (r *registerCtrlType) SetEnable(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrlFieldEnableMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrlFieldEnableMask)
	}
}

const (
	RegisterCtrlFieldAllnsShift = 1
	RegisterCtrlFieldAllnsMask  = 0x2
)

// GetAllns When SAU_CTRL.ENABLE is 0 this bit controls whether the memory is marked as Non-secure or Secure
func (r *registerCtrlType) GetAllns() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldAllnsMask) != 0
}

// SetAllns When SAU_CTRL.ENABLE is 0 this bit controls whether the memory is marked as Non-secure or Secure
func (r *registerCtrlType) SetAllns(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrlFieldAllnsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrlFieldAllnsMask)
	}
}

// registerTypeType Indicates the number of regions implemented by the Security Attribution Unit (SAU)
type registerTypeType uint32

const (
	RegisterTypeFieldSregionShift = 0
	RegisterTypeFieldSregionMask  = 0xff
)

// GetSregion The number of implemented (SAU) regions
func (r *registerTypeType) GetSregion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTypeFieldSregionMask) >> RegisterTypeFieldSregionShift)
}

// SetSregion The number of implemented (SAU) regions
func (r *registerTypeType) SetSregion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTypeFieldSregionMask)|(uint32(value)<<RegisterTypeFieldSregionShift))
}

// registerRnrType Selects the region currently accessed by SAU_RBAR and SAU_RLAR
type registerRnrType uint32

const (
	RegisterRnrFieldRegionShift = 0
	RegisterRnrFieldRegionMask  = 0xff
)

// GetRegion Indicates the Security Attribution Unit (SAU) region accessed by SAU_RBAR and SAU_RLAR
func (r *registerRnrType) GetRegion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRnrFieldRegionMask) >> RegisterRnrFieldRegionShift)
}

// SetRegion Indicates the Security Attribution Unit (SAU) region accessed by SAU_RBAR and SAU_RLAR
func (r *registerRnrType) SetRegion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRnrFieldRegionMask)|(uint32(value)<<RegisterRnrFieldRegionShift))
}

// registerRbarType Provides indirect read and write access to the base address of the currently selected Security Attribution Unit (SAU) region
type registerRbarType uint32

const (
	RegisterRbarFieldBaddrShift = 5
	RegisterRbarFieldBaddrMask  = 0xffffffe0
)

// GetBaddr Holds bits[31:5] of the base address for the selected (SAU) region
func (r *registerRbarType) GetBaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRbarFieldBaddrMask) >> RegisterRbarFieldBaddrShift)
}

// SetBaddr Holds bits[31:5] of the base address for the selected (SAU) region
func (r *registerRbarType) SetBaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRbarFieldBaddrMask)|(uint32(value)<<RegisterRbarFieldBaddrShift))
}

// registerRlarType Provides indirect read and write access to the limit address of the currently selected Security Attribution Unit (SAU) region
type registerRlarType uint32

const (
	RegisterRlarFieldLaddrShift = 5
	RegisterRlarFieldLaddrMask  = 0xffffffe0
)

// GetLaddr Holds bits[31:5] of the limit address for the selected (SAU) region
func (r *registerRlarType) GetLaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRlarFieldLaddrMask) >> RegisterRlarFieldLaddrShift)
}

// SetLaddr Holds bits[31:5] of the limit address for the selected (SAU) region
func (r *registerRlarType) SetLaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRlarFieldLaddrMask)|(uint32(value)<<RegisterRlarFieldLaddrShift))
}

const (
	RegisterRlarFieldNscShift = 5
	RegisterRlarFieldNscMask  = 0xffffffe0
)

// GetNsc Controls whether Non-secure state is permitted to execute an SG instruction from this region
func (r *registerRlarType) GetNsc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRlarFieldNscMask) >> RegisterRlarFieldNscShift)
}

// SetNsc Controls whether Non-secure state is permitted to execute an SG instruction from this region
func (r *registerRlarType) SetNsc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRlarFieldNscMask)|(uint32(value)<<RegisterRlarFieldNscShift))
}

const (
	RegisterRlarFieldEnableShift = 5
	RegisterRlarFieldEnableMask  = 0xffffffe0
)

// GetEnable (SAU) region enable
func (r *registerRlarType) GetEnable() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRlarFieldEnableMask) >> RegisterRlarFieldEnableShift)
}

// SetEnable (SAU) region enable
func (r *registerRlarType) SetEnable(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRlarFieldEnableMask)|(uint32(value)<<RegisterRlarFieldEnableShift))
}

// registerSfsrType Provides information about any security related faults
type registerSfsrType uint32

const (
	RegisterSfsrFieldInvepShift = 0
	RegisterSfsrFieldInvepMask  = 0x1
)

// GetInvep Inidcates a function call from the Non-secure state or exception targets a non-SG instruction in the Secure state
func (r *registerSfsrType) GetInvep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSfsrFieldInvepMask) != 0
}

// SetInvep Inidcates a function call from the Non-secure state or exception targets a non-SG instruction in the Secure state
func (r *registerSfsrType) SetInvep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSfsrFieldInvepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSfsrFieldInvepMask)
	}
}

const (
	RegisterSfsrFieldInvisShift = 1
	RegisterSfsrFieldInvisMask  = 0x2
)

// GetInvis Indicates the integrity signature in an exception stack frame is found to be invalid during the unstacking operation
func (r *registerSfsrType) GetInvis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSfsrFieldInvisMask) != 0
}

// SetInvis Indicates the integrity signature in an exception stack frame is found to be invalid during the unstacking operation
func (r *registerSfsrType) SetInvis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSfsrFieldInvisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSfsrFieldInvisMask)
	}
}

const (
	RegisterSfsrFieldInverShift = 2
	RegisterSfsrFieldInverMask  = 0x4
)

// GetInver Indicates EXC_RETURN.DCRS being set to 0 when returning from an exception in the Non-secure state, or EXC_RETURN.ES being set to 1 when returning from an exception in the Non-secure state
func (r *registerSfsrType) GetInver() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSfsrFieldInverMask) != 0
}

// SetInver Indicates EXC_RETURN.DCRS being set to 0 when returning from an exception in the Non-secure state, or EXC_RETURN.ES being set to 1 when returning from an exception in the Non-secure state
func (r *registerSfsrType) SetInver(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSfsrFieldInverMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSfsrFieldInverMask)
	}
}

const (
	RegisterSfsrFieldAuviolShift = 3
	RegisterSfsrFieldAuviolMask  = 0x8
)

// GetAuviol Sticky flag indicating that an attempt was made to access parts of the address space that are marked as Secure with NS-Req for the transaction set to Non-secure
func (r *registerSfsrType) GetAuviol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSfsrFieldAuviolMask) != 0
}

// SetAuviol Sticky flag indicating that an attempt was made to access parts of the address space that are marked as Secure with NS-Req for the transaction set to Non-secure
func (r *registerSfsrType) SetAuviol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSfsrFieldAuviolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSfsrFieldAuviolMask)
	}
}

const (
	RegisterSfsrFieldInvtranShift = 4
	RegisterSfsrFieldInvtranMask  = 0x10
)

// GetInvtran Sticky flag indicating that an exception was raised due to a branch that was not flagged as being domain crossing causing a transition from Secure to Non-secure memory
func (r *registerSfsrType) GetInvtran() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSfsrFieldInvtranMask) != 0
}

// SetInvtran Sticky flag indicating that an exception was raised due to a branch that was not flagged as being domain crossing causing a transition from Secure to Non-secure memory
func (r *registerSfsrType) SetInvtran(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSfsrFieldInvtranMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSfsrFieldInvtranMask)
	}
}

const (
	RegisterSfsrFieldLsperrShift = 5
	RegisterSfsrFieldLsperrMask  = 0x20
)

// GetLsperr Sticky flag indicating that an Security Attribution Unit (SAU) or Implementation Defined Attribution Unit (IDAU) violation occurred during the lazy preservation of floating-point state
func (r *registerSfsrType) GetLsperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSfsrFieldLsperrMask) != 0
}

// SetLsperr Sticky flag indicating that an Security Attribution Unit (SAU) or Implementation Defined Attribution Unit (IDAU) violation occurred during the lazy preservation of floating-point state
func (r *registerSfsrType) SetLsperr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSfsrFieldLsperrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSfsrFieldLsperrMask)
	}
}

const (
	RegisterSfsrFieldSfarvalidShift = 6
	RegisterSfsrFieldSfarvalidMask  = 0x40
)

// GetSfarvalid This bit is set when the SFAR register contains a valid value
func (r *registerSfsrType) GetSfarvalid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSfsrFieldSfarvalidMask) != 0
}

// SetSfarvalid This bit is set when the SFAR register contains a valid value
func (r *registerSfsrType) SetSfarvalid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSfsrFieldSfarvalidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSfsrFieldSfarvalidMask)
	}
}

const (
	RegisterSfsrFieldLserrShift = 7
	RegisterSfsrFieldLserrMask  = 0x80
)

// GetLserr Sticky flag indicating that an error occurred during lazy state activation or deactivation
func (r *registerSfsrType) GetLserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSfsrFieldLserrMask) != 0
}

// SetLserr Sticky flag indicating that an error occurred during lazy state activation or deactivation
func (r *registerSfsrType) SetLserr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSfsrFieldLserrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSfsrFieldLserrMask)
	}
}

// registerSfarType Shows the address of the memory location that caused a security violation
type registerSfarType uint32
