//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package mdios

import (
	"unsafe"
	"volatile"
)

var (
	Mdios = (*_mdios)(unsafe.Pointer(uintptr(0x40009400)))
)

type _mdios struct {
	Mdios_cr      registerMdios_crType
	Mdios_wrfr    registerMdios_wrfrType
	Mdios_cwrfr   registerMdios_cwrfrType
	Mdios_rdfr    registerMdios_rdfrType
	Mdios_crdfr   registerMdios_crdfrType
	Mdios_sr      registerMdios_srType
	Mdios_clrfr   registerMdios_clrfrType
	Mdios_dinr0   registerMdios_dinr0Type
	Mdios_dinr1   registerMdios_dinr1Type
	Mdios_dinr2   registerMdios_dinr2Type
	Mdios_dinr3   registerMdios_dinr3Type
	Mdios_dinr4   registerMdios_dinr4Type
	Mdios_dinr5   registerMdios_dinr5Type
	Mdios_dinr6   registerMdios_dinr6Type
	Mdios_dinr7   registerMdios_dinr7Type
	Mdios_dinr8   registerMdios_dinr8Type
	Mdios_dinr9   registerMdios_dinr9Type
	Mdios_dinr10  registerMdios_dinr10Type
	Mdios_dinr11  registerMdios_dinr11Type
	Mdios_dinr12  registerMdios_dinr12Type
	Mdios_dinr13  registerMdios_dinr13Type
	Mdios_dinr14  registerMdios_dinr14Type
	Mdios_dinr15  registerMdios_dinr15Type
	Mdios_dinr16  registerMdios_dinr16Type
	Mdios_dinr17  registerMdios_dinr17Type
	Mdios_dinr18  registerMdios_dinr18Type
	Mdios_dinr19  registerMdios_dinr19Type
	Mdios_dinr20  registerMdios_dinr20Type
	Mdios_dinr21  registerMdios_dinr21Type
	Mdios_dinr22  registerMdios_dinr22Type
	Mdios_dinr23  registerMdios_dinr23Type
	Mdios_dinr24  registerMdios_dinr24Type
	Mdios_dinr25  registerMdios_dinr25Type
	Mdios_dinr26  registerMdios_dinr26Type
	Mdios_dinr27  registerMdios_dinr27Type
	Mdios_dinr28  registerMdios_dinr28Type
	Mdios_dinr29  registerMdios_dinr29Type
	Mdios_dinr30  registerMdios_dinr30Type
	Mdios_dinr31  registerMdios_dinr31Type
	Mdios_doutr0  registerMdios_doutr0Type
	Mdios_doutr1  registerMdios_doutr1Type
	Mdios_doutr2  registerMdios_doutr2Type
	Mdios_doutr3  registerMdios_doutr3Type
	Mdios_doutr4  registerMdios_doutr4Type
	Mdios_doutr5  registerMdios_doutr5Type
	Mdios_doutr6  registerMdios_doutr6Type
	Mdios_doutr7  registerMdios_doutr7Type
	Mdios_doutr8  registerMdios_doutr8Type
	Mdios_doutr9  registerMdios_doutr9Type
	Mdios_doutr10 registerMdios_doutr10Type
	Mdios_doutr11 registerMdios_doutr11Type
	Mdios_doutr12 registerMdios_doutr12Type
	Mdios_doutr13 registerMdios_doutr13Type
	Mdios_doutr14 registerMdios_doutr14Type
	Mdios_doutr15 registerMdios_doutr15Type
	Mdios_doutr16 registerMdios_doutr16Type
	Mdios_doutr17 registerMdios_doutr17Type
	Mdios_doutr18 registerMdios_doutr18Type
	Mdios_doutr19 registerMdios_doutr19Type
	Mdios_doutr20 registerMdios_doutr20Type
	Mdios_doutr21 registerMdios_doutr21Type
	Mdios_doutr22 registerMdios_doutr22Type
	Mdios_doutr23 registerMdios_doutr23Type
	Mdios_doutr24 registerMdios_doutr24Type
	Mdios_doutr25 registerMdios_doutr25Type
	Mdios_doutr26 registerMdios_doutr26Type
	Mdios_doutr27 registerMdios_doutr27Type
	Mdios_doutr28 registerMdios_doutr28Type
	Mdios_doutr29 registerMdios_doutr29Type
	Mdios_doutr30 registerMdios_doutr30Type
	Mdios_doutr31 registerMdios_doutr31Type
}

// registerMdios_crType MDIOS configuration register
type registerMdios_crType uint32

const (
	RegisterMdios_crFieldEnShift = 0
	RegisterMdios_crFieldEnMask  = 0x1
)

// GetEn Peripheral enable
func (r *registerMdios_crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdios_crFieldEnMask) != 0
}

// SetEn Peripheral enable
func (r *registerMdios_crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdios_crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdios_crFieldEnMask)
	}
}

const (
	RegisterMdios_crFieldWrieShift = 1
	RegisterMdios_crFieldWrieMask  = 0x2
)

// GetWrie Register write interrupt enable
func (r *registerMdios_crType) GetWrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdios_crFieldWrieMask) != 0
}

// SetWrie Register write interrupt enable
func (r *registerMdios_crType) SetWrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdios_crFieldWrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdios_crFieldWrieMask)
	}
}

const (
	RegisterMdios_crFieldRdieShift = 2
	RegisterMdios_crFieldRdieMask  = 0x4
)

// GetRdie Register Read Interrupt Enable
func (r *registerMdios_crType) GetRdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdios_crFieldRdieMask) != 0
}

// SetRdie Register Read Interrupt Enable
func (r *registerMdios_crType) SetRdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdios_crFieldRdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdios_crFieldRdieMask)
	}
}

const (
	RegisterMdios_crFieldEieShift = 3
	RegisterMdios_crFieldEieMask  = 0x8
)

// GetEie Error interrupt enable
func (r *registerMdios_crType) GetEie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdios_crFieldEieMask) != 0
}

// SetEie Error interrupt enable
func (r *registerMdios_crType) SetEie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdios_crFieldEieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdios_crFieldEieMask)
	}
}

const (
	RegisterMdios_crFieldDpcShift = 7
	RegisterMdios_crFieldDpcMask  = 0x80
)

// GetDpc Disable Preamble Check
func (r *registerMdios_crType) GetDpc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdios_crFieldDpcMask) != 0
}

// SetDpc Disable Preamble Check
func (r *registerMdios_crType) SetDpc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdios_crFieldDpcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdios_crFieldDpcMask)
	}
}

const (
	RegisterMdios_crFieldPort_addressShift = 8
	RegisterMdios_crFieldPort_addressMask  = 0x1f00
)

// GetPort_address Slaves's address
func (r *registerMdios_crType) GetPort_address() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_crFieldPort_addressMask) >> RegisterMdios_crFieldPort_addressShift)
}

// SetPort_address Slaves's address
func (r *registerMdios_crType) SetPort_address(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_crFieldPort_addressMask)|(uint32(value)<<RegisterMdios_crFieldPort_addressShift))
}

// registerMdios_wrfrType MDIOS write flag register
type registerMdios_wrfrType uint32

const (
	RegisterMdios_wrfrFieldWrfShift = 0
	RegisterMdios_wrfrFieldWrfMask  = 0xffffffff
)

// GetWrf Write flags for MDIO registers 0 to 31
func (r *registerMdios_wrfrType) GetWrf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_wrfrFieldWrfMask) >> RegisterMdios_wrfrFieldWrfShift)
}

// SetWrf Write flags for MDIO registers 0 to 31
func (r *registerMdios_wrfrType) SetWrf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_wrfrFieldWrfMask)|(uint32(value)<<RegisterMdios_wrfrFieldWrfShift))
}

// registerMdios_cwrfrType MDIOS clear write flag register
type registerMdios_cwrfrType uint32

const (
	RegisterMdios_cwrfrFieldCwrfShift = 0
	RegisterMdios_cwrfrFieldCwrfMask  = 0xffffffff
)

// GetCwrf Clear the write flag
func (r *registerMdios_cwrfrType) GetCwrf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_cwrfrFieldCwrfMask) >> RegisterMdios_cwrfrFieldCwrfShift)
}

// SetCwrf Clear the write flag
func (r *registerMdios_cwrfrType) SetCwrf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_cwrfrFieldCwrfMask)|(uint32(value)<<RegisterMdios_cwrfrFieldCwrfShift))
}

// registerMdios_rdfrType MDIOS read flag register
type registerMdios_rdfrType uint32

const (
	RegisterMdios_rdfrFieldRdfShift = 0
	RegisterMdios_rdfrFieldRdfMask  = 0xffffffff
)

// GetRdf Read flags for MDIO registers 0 to 31
func (r *registerMdios_rdfrType) GetRdf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_rdfrFieldRdfMask) >> RegisterMdios_rdfrFieldRdfShift)
}

// SetRdf Read flags for MDIO registers 0 to 31
func (r *registerMdios_rdfrType) SetRdf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_rdfrFieldRdfMask)|(uint32(value)<<RegisterMdios_rdfrFieldRdfShift))
}

// registerMdios_crdfrType MDIOS clear read flag register
type registerMdios_crdfrType uint32

const (
	RegisterMdios_crdfrFieldCrdfShift = 0
	RegisterMdios_crdfrFieldCrdfMask  = 0xffffffff
)

// GetCrdf Clear the read flag
func (r *registerMdios_crdfrType) GetCrdf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_crdfrFieldCrdfMask) >> RegisterMdios_crdfrFieldCrdfShift)
}

// SetCrdf Clear the read flag
func (r *registerMdios_crdfrType) SetCrdf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_crdfrFieldCrdfMask)|(uint32(value)<<RegisterMdios_crdfrFieldCrdfShift))
}

// registerMdios_srType MDIOS status register
type registerMdios_srType uint32

const (
	RegisterMdios_srFieldPerfShift = 0
	RegisterMdios_srFieldPerfMask  = 0x1
)

// GetPerf Preamble error flag
func (r *registerMdios_srType) GetPerf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdios_srFieldPerfMask) != 0
}

// SetPerf Preamble error flag
func (r *registerMdios_srType) SetPerf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdios_srFieldPerfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdios_srFieldPerfMask)
	}
}

const (
	RegisterMdios_srFieldSerfShift = 1
	RegisterMdios_srFieldSerfMask  = 0x2
)

// GetSerf Start error flag
func (r *registerMdios_srType) GetSerf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdios_srFieldSerfMask) != 0
}

// SetSerf Start error flag
func (r *registerMdios_srType) SetSerf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdios_srFieldSerfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdios_srFieldSerfMask)
	}
}

const (
	RegisterMdios_srFieldTerfShift = 2
	RegisterMdios_srFieldTerfMask  = 0x4
)

// GetTerf Turnaround error flag
func (r *registerMdios_srType) GetTerf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdios_srFieldTerfMask) != 0
}

// SetTerf Turnaround error flag
func (r *registerMdios_srType) SetTerf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdios_srFieldTerfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdios_srFieldTerfMask)
	}
}

// registerMdios_clrfrType MDIOS clear flag register
type registerMdios_clrfrType uint32

const (
	RegisterMdios_clrfrFieldCperfShift = 0
	RegisterMdios_clrfrFieldCperfMask  = 0x1
)

// GetCperf Clear the preamble error flag
func (r *registerMdios_clrfrType) GetCperf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdios_clrfrFieldCperfMask) != 0
}

// SetCperf Clear the preamble error flag
func (r *registerMdios_clrfrType) SetCperf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdios_clrfrFieldCperfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdios_clrfrFieldCperfMask)
	}
}

const (
	RegisterMdios_clrfrFieldCserfShift = 1
	RegisterMdios_clrfrFieldCserfMask  = 0x2
)

// GetCserf Clear the start error flag
func (r *registerMdios_clrfrType) GetCserf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdios_clrfrFieldCserfMask) != 0
}

// SetCserf Clear the start error flag
func (r *registerMdios_clrfrType) SetCserf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdios_clrfrFieldCserfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdios_clrfrFieldCserfMask)
	}
}

const (
	RegisterMdios_clrfrFieldCterfShift = 2
	RegisterMdios_clrfrFieldCterfMask  = 0x4
)

// GetCterf Clear the turnaround error flag
func (r *registerMdios_clrfrType) GetCterf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdios_clrfrFieldCterfMask) != 0
}

// SetCterf Clear the turnaround error flag
func (r *registerMdios_clrfrType) SetCterf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdios_clrfrFieldCterfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdios_clrfrFieldCterfMask)
	}
}

// registerMdios_dinr0Type MDIOS input data register 0
type registerMdios_dinr0Type uint32

const (
	RegisterMdios_dinr0FieldDin0Shift = 0
	RegisterMdios_dinr0FieldDin0Mask  = 0xffff
)

// GetDin0 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr0Type) GetDin0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr0FieldDin0Mask) >> RegisterMdios_dinr0FieldDin0Shift)
}

// SetDin0 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr0Type) SetDin0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr0FieldDin0Mask)|(uint32(value)<<RegisterMdios_dinr0FieldDin0Shift))
}

// registerMdios_dinr1Type MDIOS input data register 1
type registerMdios_dinr1Type uint32

const (
	RegisterMdios_dinr1FieldDin1Shift = 0
	RegisterMdios_dinr1FieldDin1Mask  = 0xffff
)

// GetDin1 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr1Type) GetDin1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr1FieldDin1Mask) >> RegisterMdios_dinr1FieldDin1Shift)
}

// SetDin1 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr1Type) SetDin1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr1FieldDin1Mask)|(uint32(value)<<RegisterMdios_dinr1FieldDin1Shift))
}

// registerMdios_dinr2Type MDIOS input data register 2
type registerMdios_dinr2Type uint32

const (
	RegisterMdios_dinr2FieldDin2Shift = 0
	RegisterMdios_dinr2FieldDin2Mask  = 0xffff
)

// GetDin2 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr2Type) GetDin2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr2FieldDin2Mask) >> RegisterMdios_dinr2FieldDin2Shift)
}

// SetDin2 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr2Type) SetDin2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr2FieldDin2Mask)|(uint32(value)<<RegisterMdios_dinr2FieldDin2Shift))
}

// registerMdios_dinr3Type MDIOS input data register 3
type registerMdios_dinr3Type uint32

const (
	RegisterMdios_dinr3FieldDin3Shift = 0
	RegisterMdios_dinr3FieldDin3Mask  = 0xffff
)

// GetDin3 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr3Type) GetDin3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr3FieldDin3Mask) >> RegisterMdios_dinr3FieldDin3Shift)
}

// SetDin3 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr3Type) SetDin3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr3FieldDin3Mask)|(uint32(value)<<RegisterMdios_dinr3FieldDin3Shift))
}

// registerMdios_dinr4Type MDIOS input data register 4
type registerMdios_dinr4Type uint32

const (
	RegisterMdios_dinr4FieldDin4Shift = 0
	RegisterMdios_dinr4FieldDin4Mask  = 0xffff
)

// GetDin4 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr4Type) GetDin4() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr4FieldDin4Mask) >> RegisterMdios_dinr4FieldDin4Shift)
}

// SetDin4 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr4Type) SetDin4(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr4FieldDin4Mask)|(uint32(value)<<RegisterMdios_dinr4FieldDin4Shift))
}

// registerMdios_dinr5Type MDIOS input data register 5
type registerMdios_dinr5Type uint32

const (
	RegisterMdios_dinr5FieldDin5Shift = 0
	RegisterMdios_dinr5FieldDin5Mask  = 0xffff
)

// GetDin5 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr5Type) GetDin5() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr5FieldDin5Mask) >> RegisterMdios_dinr5FieldDin5Shift)
}

// SetDin5 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr5Type) SetDin5(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr5FieldDin5Mask)|(uint32(value)<<RegisterMdios_dinr5FieldDin5Shift))
}

// registerMdios_dinr6Type MDIOS input data register 6
type registerMdios_dinr6Type uint32

const (
	RegisterMdios_dinr6FieldDin6Shift = 0
	RegisterMdios_dinr6FieldDin6Mask  = 0xffff
)

// GetDin6 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr6Type) GetDin6() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr6FieldDin6Mask) >> RegisterMdios_dinr6FieldDin6Shift)
}

// SetDin6 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr6Type) SetDin6(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr6FieldDin6Mask)|(uint32(value)<<RegisterMdios_dinr6FieldDin6Shift))
}

// registerMdios_dinr7Type MDIOS input data register 7
type registerMdios_dinr7Type uint32

const (
	RegisterMdios_dinr7FieldDin7Shift = 0
	RegisterMdios_dinr7FieldDin7Mask  = 0xffff
)

// GetDin7 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr7Type) GetDin7() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr7FieldDin7Mask) >> RegisterMdios_dinr7FieldDin7Shift)
}

// SetDin7 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr7Type) SetDin7(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr7FieldDin7Mask)|(uint32(value)<<RegisterMdios_dinr7FieldDin7Shift))
}

// registerMdios_dinr8Type MDIOS input data register 8
type registerMdios_dinr8Type uint32

const (
	RegisterMdios_dinr8FieldDin8Shift = 0
	RegisterMdios_dinr8FieldDin8Mask  = 0xffff
)

// GetDin8 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr8Type) GetDin8() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr8FieldDin8Mask) >> RegisterMdios_dinr8FieldDin8Shift)
}

// SetDin8 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr8Type) SetDin8(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr8FieldDin8Mask)|(uint32(value)<<RegisterMdios_dinr8FieldDin8Shift))
}

// registerMdios_dinr9Type MDIOS input data register 9
type registerMdios_dinr9Type uint32

const (
	RegisterMdios_dinr9FieldDin9Shift = 0
	RegisterMdios_dinr9FieldDin9Mask  = 0xffff
)

// GetDin9 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr9Type) GetDin9() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr9FieldDin9Mask) >> RegisterMdios_dinr9FieldDin9Shift)
}

// SetDin9 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr9Type) SetDin9(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr9FieldDin9Mask)|(uint32(value)<<RegisterMdios_dinr9FieldDin9Shift))
}

// registerMdios_dinr10Type MDIOS input data register 10
type registerMdios_dinr10Type uint32

const (
	RegisterMdios_dinr10FieldDin10Shift = 0
	RegisterMdios_dinr10FieldDin10Mask  = 0xffff
)

// GetDin10 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr10Type) GetDin10() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr10FieldDin10Mask) >> RegisterMdios_dinr10FieldDin10Shift)
}

// SetDin10 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr10Type) SetDin10(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr10FieldDin10Mask)|(uint32(value)<<RegisterMdios_dinr10FieldDin10Shift))
}

// registerMdios_dinr11Type MDIOS input data register 11
type registerMdios_dinr11Type uint32

const (
	RegisterMdios_dinr11FieldDin11Shift = 0
	RegisterMdios_dinr11FieldDin11Mask  = 0xffff
)

// GetDin11 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr11Type) GetDin11() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr11FieldDin11Mask) >> RegisterMdios_dinr11FieldDin11Shift)
}

// SetDin11 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr11Type) SetDin11(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr11FieldDin11Mask)|(uint32(value)<<RegisterMdios_dinr11FieldDin11Shift))
}

// registerMdios_dinr12Type MDIOS input data register 12
type registerMdios_dinr12Type uint32

const (
	RegisterMdios_dinr12FieldDin12Shift = 0
	RegisterMdios_dinr12FieldDin12Mask  = 0xffff
)

// GetDin12 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr12Type) GetDin12() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr12FieldDin12Mask) >> RegisterMdios_dinr12FieldDin12Shift)
}

// SetDin12 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr12Type) SetDin12(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr12FieldDin12Mask)|(uint32(value)<<RegisterMdios_dinr12FieldDin12Shift))
}

// registerMdios_dinr13Type MDIOS input data register 13
type registerMdios_dinr13Type uint32

const (
	RegisterMdios_dinr13FieldDin13Shift = 0
	RegisterMdios_dinr13FieldDin13Mask  = 0xffff
)

// GetDin13 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr13Type) GetDin13() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr13FieldDin13Mask) >> RegisterMdios_dinr13FieldDin13Shift)
}

// SetDin13 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr13Type) SetDin13(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr13FieldDin13Mask)|(uint32(value)<<RegisterMdios_dinr13FieldDin13Shift))
}

// registerMdios_dinr14Type MDIOS input data register 14
type registerMdios_dinr14Type uint32

const (
	RegisterMdios_dinr14FieldDin14Shift = 0
	RegisterMdios_dinr14FieldDin14Mask  = 0xffff
)

// GetDin14 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr14Type) GetDin14() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr14FieldDin14Mask) >> RegisterMdios_dinr14FieldDin14Shift)
}

// SetDin14 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr14Type) SetDin14(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr14FieldDin14Mask)|(uint32(value)<<RegisterMdios_dinr14FieldDin14Shift))
}

// registerMdios_dinr15Type MDIOS input data register 15
type registerMdios_dinr15Type uint32

const (
	RegisterMdios_dinr15FieldDin15Shift = 0
	RegisterMdios_dinr15FieldDin15Mask  = 0xffff
)

// GetDin15 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr15Type) GetDin15() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr15FieldDin15Mask) >> RegisterMdios_dinr15FieldDin15Shift)
}

// SetDin15 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr15Type) SetDin15(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr15FieldDin15Mask)|(uint32(value)<<RegisterMdios_dinr15FieldDin15Shift))
}

// registerMdios_dinr16Type MDIOS input data register 16
type registerMdios_dinr16Type uint32

const (
	RegisterMdios_dinr16FieldDin16Shift = 0
	RegisterMdios_dinr16FieldDin16Mask  = 0xffff
)

// GetDin16 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr16Type) GetDin16() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr16FieldDin16Mask) >> RegisterMdios_dinr16FieldDin16Shift)
}

// SetDin16 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr16Type) SetDin16(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr16FieldDin16Mask)|(uint32(value)<<RegisterMdios_dinr16FieldDin16Shift))
}

// registerMdios_dinr17Type MDIOS input data register 17
type registerMdios_dinr17Type uint32

const (
	RegisterMdios_dinr17FieldDin17Shift = 0
	RegisterMdios_dinr17FieldDin17Mask  = 0xffff
)

// GetDin17 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr17Type) GetDin17() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr17FieldDin17Mask) >> RegisterMdios_dinr17FieldDin17Shift)
}

// SetDin17 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr17Type) SetDin17(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr17FieldDin17Mask)|(uint32(value)<<RegisterMdios_dinr17FieldDin17Shift))
}

// registerMdios_dinr18Type MDIOS input data register 18
type registerMdios_dinr18Type uint32

const (
	RegisterMdios_dinr18FieldDin18Shift = 0
	RegisterMdios_dinr18FieldDin18Mask  = 0xffff
)

// GetDin18 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr18Type) GetDin18() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr18FieldDin18Mask) >> RegisterMdios_dinr18FieldDin18Shift)
}

// SetDin18 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr18Type) SetDin18(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr18FieldDin18Mask)|(uint32(value)<<RegisterMdios_dinr18FieldDin18Shift))
}

// registerMdios_dinr19Type MDIOS input data register 19
type registerMdios_dinr19Type uint32

const (
	RegisterMdios_dinr19FieldDin19Shift = 0
	RegisterMdios_dinr19FieldDin19Mask  = 0xffff
)

// GetDin19 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr19Type) GetDin19() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr19FieldDin19Mask) >> RegisterMdios_dinr19FieldDin19Shift)
}

// SetDin19 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr19Type) SetDin19(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr19FieldDin19Mask)|(uint32(value)<<RegisterMdios_dinr19FieldDin19Shift))
}

// registerMdios_dinr20Type MDIOS input data register 20
type registerMdios_dinr20Type uint32

const (
	RegisterMdios_dinr20FieldDin20Shift = 0
	RegisterMdios_dinr20FieldDin20Mask  = 0xffff
)

// GetDin20 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr20Type) GetDin20() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr20FieldDin20Mask) >> RegisterMdios_dinr20FieldDin20Shift)
}

// SetDin20 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr20Type) SetDin20(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr20FieldDin20Mask)|(uint32(value)<<RegisterMdios_dinr20FieldDin20Shift))
}

// registerMdios_dinr21Type MDIOS input data register 21
type registerMdios_dinr21Type uint32

const (
	RegisterMdios_dinr21FieldDin21Shift = 0
	RegisterMdios_dinr21FieldDin21Mask  = 0xffff
)

// GetDin21 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr21Type) GetDin21() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr21FieldDin21Mask) >> RegisterMdios_dinr21FieldDin21Shift)
}

// SetDin21 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr21Type) SetDin21(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr21FieldDin21Mask)|(uint32(value)<<RegisterMdios_dinr21FieldDin21Shift))
}

// registerMdios_dinr22Type MDIOS input data register 22
type registerMdios_dinr22Type uint32

const (
	RegisterMdios_dinr22FieldDin22Shift = 0
	RegisterMdios_dinr22FieldDin22Mask  = 0xffff
)

// GetDin22 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr22Type) GetDin22() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr22FieldDin22Mask) >> RegisterMdios_dinr22FieldDin22Shift)
}

// SetDin22 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr22Type) SetDin22(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr22FieldDin22Mask)|(uint32(value)<<RegisterMdios_dinr22FieldDin22Shift))
}

// registerMdios_dinr23Type MDIOS input data register 23
type registerMdios_dinr23Type uint32

const (
	RegisterMdios_dinr23FieldDin23Shift = 0
	RegisterMdios_dinr23FieldDin23Mask  = 0xffff
)

// GetDin23 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr23Type) GetDin23() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr23FieldDin23Mask) >> RegisterMdios_dinr23FieldDin23Shift)
}

// SetDin23 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr23Type) SetDin23(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr23FieldDin23Mask)|(uint32(value)<<RegisterMdios_dinr23FieldDin23Shift))
}

// registerMdios_dinr24Type MDIOS input data register 24
type registerMdios_dinr24Type uint32

const (
	RegisterMdios_dinr24FieldDin24Shift = 0
	RegisterMdios_dinr24FieldDin24Mask  = 0xffff
)

// GetDin24 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr24Type) GetDin24() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr24FieldDin24Mask) >> RegisterMdios_dinr24FieldDin24Shift)
}

// SetDin24 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr24Type) SetDin24(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr24FieldDin24Mask)|(uint32(value)<<RegisterMdios_dinr24FieldDin24Shift))
}

// registerMdios_dinr25Type MDIOS input data register 25
type registerMdios_dinr25Type uint32

const (
	RegisterMdios_dinr25FieldDin25Shift = 0
	RegisterMdios_dinr25FieldDin25Mask  = 0xffff
)

// GetDin25 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr25Type) GetDin25() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr25FieldDin25Mask) >> RegisterMdios_dinr25FieldDin25Shift)
}

// SetDin25 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr25Type) SetDin25(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr25FieldDin25Mask)|(uint32(value)<<RegisterMdios_dinr25FieldDin25Shift))
}

// registerMdios_dinr26Type MDIOS input data register 26
type registerMdios_dinr26Type uint32

const (
	RegisterMdios_dinr26FieldDin26Shift = 0
	RegisterMdios_dinr26FieldDin26Mask  = 0xffff
)

// GetDin26 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr26Type) GetDin26() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr26FieldDin26Mask) >> RegisterMdios_dinr26FieldDin26Shift)
}

// SetDin26 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr26Type) SetDin26(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr26FieldDin26Mask)|(uint32(value)<<RegisterMdios_dinr26FieldDin26Shift))
}

// registerMdios_dinr27Type MDIOS input data register 27
type registerMdios_dinr27Type uint32

const (
	RegisterMdios_dinr27FieldDin27Shift = 0
	RegisterMdios_dinr27FieldDin27Mask  = 0xffff
)

// GetDin27 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr27Type) GetDin27() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr27FieldDin27Mask) >> RegisterMdios_dinr27FieldDin27Shift)
}

// SetDin27 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr27Type) SetDin27(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr27FieldDin27Mask)|(uint32(value)<<RegisterMdios_dinr27FieldDin27Shift))
}

// registerMdios_dinr28Type MDIOS input data register 28
type registerMdios_dinr28Type uint32

const (
	RegisterMdios_dinr28FieldDin28Shift = 0
	RegisterMdios_dinr28FieldDin28Mask  = 0xffff
)

// GetDin28 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr28Type) GetDin28() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr28FieldDin28Mask) >> RegisterMdios_dinr28FieldDin28Shift)
}

// SetDin28 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr28Type) SetDin28(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr28FieldDin28Mask)|(uint32(value)<<RegisterMdios_dinr28FieldDin28Shift))
}

// registerMdios_dinr29Type MDIOS input data register 29
type registerMdios_dinr29Type uint32

const (
	RegisterMdios_dinr29FieldDin29Shift = 0
	RegisterMdios_dinr29FieldDin29Mask  = 0xffff
)

// GetDin29 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr29Type) GetDin29() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr29FieldDin29Mask) >> RegisterMdios_dinr29FieldDin29Shift)
}

// SetDin29 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr29Type) SetDin29(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr29FieldDin29Mask)|(uint32(value)<<RegisterMdios_dinr29FieldDin29Shift))
}

// registerMdios_dinr30Type MDIOS input data register 30
type registerMdios_dinr30Type uint32

const (
	RegisterMdios_dinr30FieldDin30Shift = 0
	RegisterMdios_dinr30FieldDin30Mask  = 0xffff
)

// GetDin30 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr30Type) GetDin30() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr30FieldDin30Mask) >> RegisterMdios_dinr30FieldDin30Shift)
}

// SetDin30 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr30Type) SetDin30(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr30FieldDin30Mask)|(uint32(value)<<RegisterMdios_dinr30FieldDin30Shift))
}

// registerMdios_dinr31Type MDIOS input data register 31
type registerMdios_dinr31Type uint32

const (
	RegisterMdios_dinr31FieldDin31Shift = 0
	RegisterMdios_dinr31FieldDin31Mask  = 0xffff
)

// GetDin31 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr31Type) GetDin31() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_dinr31FieldDin31Mask) >> RegisterMdios_dinr31FieldDin31Shift)
}

// SetDin31 Input data received from MDIO Master during write frames
func (r *registerMdios_dinr31Type) SetDin31(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_dinr31FieldDin31Mask)|(uint32(value)<<RegisterMdios_dinr31FieldDin31Shift))
}

// registerMdios_doutr0Type MDIOS output data register 0
type registerMdios_doutr0Type uint32

const (
	RegisterMdios_doutr0FieldDout0Shift = 0
	RegisterMdios_doutr0FieldDout0Mask  = 0xffff
)

// GetDout0 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr0Type) GetDout0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr0FieldDout0Mask) >> RegisterMdios_doutr0FieldDout0Shift)
}

// SetDout0 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr0Type) SetDout0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr0FieldDout0Mask)|(uint32(value)<<RegisterMdios_doutr0FieldDout0Shift))
}

// registerMdios_doutr1Type MDIOS output data register 1
type registerMdios_doutr1Type uint32

const (
	RegisterMdios_doutr1FieldDout1Shift = 0
	RegisterMdios_doutr1FieldDout1Mask  = 0xffff
)

// GetDout1 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr1Type) GetDout1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr1FieldDout1Mask) >> RegisterMdios_doutr1FieldDout1Shift)
}

// SetDout1 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr1Type) SetDout1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr1FieldDout1Mask)|(uint32(value)<<RegisterMdios_doutr1FieldDout1Shift))
}

// registerMdios_doutr2Type MDIOS output data register 2
type registerMdios_doutr2Type uint32

const (
	RegisterMdios_doutr2FieldDout2Shift = 0
	RegisterMdios_doutr2FieldDout2Mask  = 0xffff
)

// GetDout2 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr2Type) GetDout2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr2FieldDout2Mask) >> RegisterMdios_doutr2FieldDout2Shift)
}

// SetDout2 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr2Type) SetDout2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr2FieldDout2Mask)|(uint32(value)<<RegisterMdios_doutr2FieldDout2Shift))
}

// registerMdios_doutr3Type MDIOS output data register 3
type registerMdios_doutr3Type uint32

const (
	RegisterMdios_doutr3FieldDout3Shift = 0
	RegisterMdios_doutr3FieldDout3Mask  = 0xffff
)

// GetDout3 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr3Type) GetDout3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr3FieldDout3Mask) >> RegisterMdios_doutr3FieldDout3Shift)
}

// SetDout3 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr3Type) SetDout3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr3FieldDout3Mask)|(uint32(value)<<RegisterMdios_doutr3FieldDout3Shift))
}

// registerMdios_doutr4Type MDIOS output data register 4
type registerMdios_doutr4Type uint32

const (
	RegisterMdios_doutr4FieldDout4Shift = 0
	RegisterMdios_doutr4FieldDout4Mask  = 0xffff
)

// GetDout4 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr4Type) GetDout4() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr4FieldDout4Mask) >> RegisterMdios_doutr4FieldDout4Shift)
}

// SetDout4 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr4Type) SetDout4(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr4FieldDout4Mask)|(uint32(value)<<RegisterMdios_doutr4FieldDout4Shift))
}

// registerMdios_doutr5Type MDIOS output data register 5
type registerMdios_doutr5Type uint32

const (
	RegisterMdios_doutr5FieldDout5Shift = 0
	RegisterMdios_doutr5FieldDout5Mask  = 0xffff
)

// GetDout5 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr5Type) GetDout5() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr5FieldDout5Mask) >> RegisterMdios_doutr5FieldDout5Shift)
}

// SetDout5 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr5Type) SetDout5(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr5FieldDout5Mask)|(uint32(value)<<RegisterMdios_doutr5FieldDout5Shift))
}

// registerMdios_doutr6Type MDIOS output data register 6
type registerMdios_doutr6Type uint32

const (
	RegisterMdios_doutr6FieldDout6Shift = 0
	RegisterMdios_doutr6FieldDout6Mask  = 0xffff
)

// GetDout6 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr6Type) GetDout6() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr6FieldDout6Mask) >> RegisterMdios_doutr6FieldDout6Shift)
}

// SetDout6 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr6Type) SetDout6(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr6FieldDout6Mask)|(uint32(value)<<RegisterMdios_doutr6FieldDout6Shift))
}

// registerMdios_doutr7Type MDIOS output data register 7
type registerMdios_doutr7Type uint32

const (
	RegisterMdios_doutr7FieldDout7Shift = 0
	RegisterMdios_doutr7FieldDout7Mask  = 0xffff
)

// GetDout7 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr7Type) GetDout7() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr7FieldDout7Mask) >> RegisterMdios_doutr7FieldDout7Shift)
}

// SetDout7 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr7Type) SetDout7(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr7FieldDout7Mask)|(uint32(value)<<RegisterMdios_doutr7FieldDout7Shift))
}

// registerMdios_doutr8Type MDIOS output data register 8
type registerMdios_doutr8Type uint32

const (
	RegisterMdios_doutr8FieldDout8Shift = 0
	RegisterMdios_doutr8FieldDout8Mask  = 0xffff
)

// GetDout8 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr8Type) GetDout8() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr8FieldDout8Mask) >> RegisterMdios_doutr8FieldDout8Shift)
}

// SetDout8 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr8Type) SetDout8(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr8FieldDout8Mask)|(uint32(value)<<RegisterMdios_doutr8FieldDout8Shift))
}

// registerMdios_doutr9Type MDIOS output data register 9
type registerMdios_doutr9Type uint32

const (
	RegisterMdios_doutr9FieldDout9Shift = 0
	RegisterMdios_doutr9FieldDout9Mask  = 0xffff
)

// GetDout9 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr9Type) GetDout9() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr9FieldDout9Mask) >> RegisterMdios_doutr9FieldDout9Shift)
}

// SetDout9 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr9Type) SetDout9(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr9FieldDout9Mask)|(uint32(value)<<RegisterMdios_doutr9FieldDout9Shift))
}

// registerMdios_doutr10Type MDIOS output data register 10
type registerMdios_doutr10Type uint32

const (
	RegisterMdios_doutr10FieldDout10Shift = 0
	RegisterMdios_doutr10FieldDout10Mask  = 0xffff
)

// GetDout10 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr10Type) GetDout10() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr10FieldDout10Mask) >> RegisterMdios_doutr10FieldDout10Shift)
}

// SetDout10 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr10Type) SetDout10(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr10FieldDout10Mask)|(uint32(value)<<RegisterMdios_doutr10FieldDout10Shift))
}

// registerMdios_doutr11Type MDIOS output data register 11
type registerMdios_doutr11Type uint32

const (
	RegisterMdios_doutr11FieldDout11Shift = 0
	RegisterMdios_doutr11FieldDout11Mask  = 0xffff
)

// GetDout11 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr11Type) GetDout11() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr11FieldDout11Mask) >> RegisterMdios_doutr11FieldDout11Shift)
}

// SetDout11 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr11Type) SetDout11(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr11FieldDout11Mask)|(uint32(value)<<RegisterMdios_doutr11FieldDout11Shift))
}

// registerMdios_doutr12Type MDIOS output data register 12
type registerMdios_doutr12Type uint32

const (
	RegisterMdios_doutr12FieldDout12Shift = 0
	RegisterMdios_doutr12FieldDout12Mask  = 0xffff
)

// GetDout12 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr12Type) GetDout12() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr12FieldDout12Mask) >> RegisterMdios_doutr12FieldDout12Shift)
}

// SetDout12 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr12Type) SetDout12(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr12FieldDout12Mask)|(uint32(value)<<RegisterMdios_doutr12FieldDout12Shift))
}

// registerMdios_doutr13Type MDIOS output data register 13
type registerMdios_doutr13Type uint32

const (
	RegisterMdios_doutr13FieldDout13Shift = 0
	RegisterMdios_doutr13FieldDout13Mask  = 0xffff
)

// GetDout13 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr13Type) GetDout13() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr13FieldDout13Mask) >> RegisterMdios_doutr13FieldDout13Shift)
}

// SetDout13 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr13Type) SetDout13(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr13FieldDout13Mask)|(uint32(value)<<RegisterMdios_doutr13FieldDout13Shift))
}

// registerMdios_doutr14Type MDIOS output data register 14
type registerMdios_doutr14Type uint32

const (
	RegisterMdios_doutr14FieldDout14Shift = 0
	RegisterMdios_doutr14FieldDout14Mask  = 0xffff
)

// GetDout14 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr14Type) GetDout14() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr14FieldDout14Mask) >> RegisterMdios_doutr14FieldDout14Shift)
}

// SetDout14 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr14Type) SetDout14(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr14FieldDout14Mask)|(uint32(value)<<RegisterMdios_doutr14FieldDout14Shift))
}

// registerMdios_doutr15Type MDIOS output data register 15
type registerMdios_doutr15Type uint32

const (
	RegisterMdios_doutr15FieldDout15Shift = 0
	RegisterMdios_doutr15FieldDout15Mask  = 0xffff
)

// GetDout15 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr15Type) GetDout15() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr15FieldDout15Mask) >> RegisterMdios_doutr15FieldDout15Shift)
}

// SetDout15 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr15Type) SetDout15(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr15FieldDout15Mask)|(uint32(value)<<RegisterMdios_doutr15FieldDout15Shift))
}

// registerMdios_doutr16Type MDIOS output data register 16
type registerMdios_doutr16Type uint32

const (
	RegisterMdios_doutr16FieldDout16Shift = 0
	RegisterMdios_doutr16FieldDout16Mask  = 0xffff
)

// GetDout16 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr16Type) GetDout16() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr16FieldDout16Mask) >> RegisterMdios_doutr16FieldDout16Shift)
}

// SetDout16 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr16Type) SetDout16(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr16FieldDout16Mask)|(uint32(value)<<RegisterMdios_doutr16FieldDout16Shift))
}

// registerMdios_doutr17Type MDIOS output data register 17
type registerMdios_doutr17Type uint32

const (
	RegisterMdios_doutr17FieldDout17Shift = 0
	RegisterMdios_doutr17FieldDout17Mask  = 0xffff
)

// GetDout17 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr17Type) GetDout17() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr17FieldDout17Mask) >> RegisterMdios_doutr17FieldDout17Shift)
}

// SetDout17 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr17Type) SetDout17(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr17FieldDout17Mask)|(uint32(value)<<RegisterMdios_doutr17FieldDout17Shift))
}

// registerMdios_doutr18Type MDIOS output data register 18
type registerMdios_doutr18Type uint32

const (
	RegisterMdios_doutr18FieldDout18Shift = 0
	RegisterMdios_doutr18FieldDout18Mask  = 0xffff
)

// GetDout18 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr18Type) GetDout18() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr18FieldDout18Mask) >> RegisterMdios_doutr18FieldDout18Shift)
}

// SetDout18 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr18Type) SetDout18(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr18FieldDout18Mask)|(uint32(value)<<RegisterMdios_doutr18FieldDout18Shift))
}

// registerMdios_doutr19Type MDIOS output data register 19
type registerMdios_doutr19Type uint32

const (
	RegisterMdios_doutr19FieldDout19Shift = 0
	RegisterMdios_doutr19FieldDout19Mask  = 0xffff
)

// GetDout19 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr19Type) GetDout19() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr19FieldDout19Mask) >> RegisterMdios_doutr19FieldDout19Shift)
}

// SetDout19 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr19Type) SetDout19(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr19FieldDout19Mask)|(uint32(value)<<RegisterMdios_doutr19FieldDout19Shift))
}

// registerMdios_doutr20Type MDIOS output data register 20
type registerMdios_doutr20Type uint32

const (
	RegisterMdios_doutr20FieldDout20Shift = 0
	RegisterMdios_doutr20FieldDout20Mask  = 0xffff
)

// GetDout20 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr20Type) GetDout20() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr20FieldDout20Mask) >> RegisterMdios_doutr20FieldDout20Shift)
}

// SetDout20 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr20Type) SetDout20(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr20FieldDout20Mask)|(uint32(value)<<RegisterMdios_doutr20FieldDout20Shift))
}

// registerMdios_doutr21Type MDIOS output data register 21
type registerMdios_doutr21Type uint32

const (
	RegisterMdios_doutr21FieldDout21Shift = 0
	RegisterMdios_doutr21FieldDout21Mask  = 0xffff
)

// GetDout21 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr21Type) GetDout21() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr21FieldDout21Mask) >> RegisterMdios_doutr21FieldDout21Shift)
}

// SetDout21 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr21Type) SetDout21(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr21FieldDout21Mask)|(uint32(value)<<RegisterMdios_doutr21FieldDout21Shift))
}

// registerMdios_doutr22Type MDIOS output data register 22
type registerMdios_doutr22Type uint32

const (
	RegisterMdios_doutr22FieldDout22Shift = 0
	RegisterMdios_doutr22FieldDout22Mask  = 0xffff
)

// GetDout22 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr22Type) GetDout22() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr22FieldDout22Mask) >> RegisterMdios_doutr22FieldDout22Shift)
}

// SetDout22 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr22Type) SetDout22(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr22FieldDout22Mask)|(uint32(value)<<RegisterMdios_doutr22FieldDout22Shift))
}

// registerMdios_doutr23Type MDIOS output data register 23
type registerMdios_doutr23Type uint32

const (
	RegisterMdios_doutr23FieldDout23Shift = 0
	RegisterMdios_doutr23FieldDout23Mask  = 0xffff
)

// GetDout23 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr23Type) GetDout23() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr23FieldDout23Mask) >> RegisterMdios_doutr23FieldDout23Shift)
}

// SetDout23 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr23Type) SetDout23(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr23FieldDout23Mask)|(uint32(value)<<RegisterMdios_doutr23FieldDout23Shift))
}

// registerMdios_doutr24Type MDIOS output data register 24
type registerMdios_doutr24Type uint32

const (
	RegisterMdios_doutr24FieldDout24Shift = 0
	RegisterMdios_doutr24FieldDout24Mask  = 0xffff
)

// GetDout24 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr24Type) GetDout24() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr24FieldDout24Mask) >> RegisterMdios_doutr24FieldDout24Shift)
}

// SetDout24 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr24Type) SetDout24(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr24FieldDout24Mask)|(uint32(value)<<RegisterMdios_doutr24FieldDout24Shift))
}

// registerMdios_doutr25Type MDIOS output data register 25
type registerMdios_doutr25Type uint32

const (
	RegisterMdios_doutr25FieldDout25Shift = 0
	RegisterMdios_doutr25FieldDout25Mask  = 0xffff
)

// GetDout25 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr25Type) GetDout25() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr25FieldDout25Mask) >> RegisterMdios_doutr25FieldDout25Shift)
}

// SetDout25 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr25Type) SetDout25(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr25FieldDout25Mask)|(uint32(value)<<RegisterMdios_doutr25FieldDout25Shift))
}

// registerMdios_doutr26Type MDIOS output data register 26
type registerMdios_doutr26Type uint32

const (
	RegisterMdios_doutr26FieldDout26Shift = 0
	RegisterMdios_doutr26FieldDout26Mask  = 0xffff
)

// GetDout26 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr26Type) GetDout26() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr26FieldDout26Mask) >> RegisterMdios_doutr26FieldDout26Shift)
}

// SetDout26 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr26Type) SetDout26(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr26FieldDout26Mask)|(uint32(value)<<RegisterMdios_doutr26FieldDout26Shift))
}

// registerMdios_doutr27Type MDIOS output data register 27
type registerMdios_doutr27Type uint32

const (
	RegisterMdios_doutr27FieldDout27Shift = 0
	RegisterMdios_doutr27FieldDout27Mask  = 0xffff
)

// GetDout27 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr27Type) GetDout27() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr27FieldDout27Mask) >> RegisterMdios_doutr27FieldDout27Shift)
}

// SetDout27 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr27Type) SetDout27(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr27FieldDout27Mask)|(uint32(value)<<RegisterMdios_doutr27FieldDout27Shift))
}

// registerMdios_doutr28Type MDIOS output data register 28
type registerMdios_doutr28Type uint32

const (
	RegisterMdios_doutr28FieldDout28Shift = 0
	RegisterMdios_doutr28FieldDout28Mask  = 0xffff
)

// GetDout28 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr28Type) GetDout28() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr28FieldDout28Mask) >> RegisterMdios_doutr28FieldDout28Shift)
}

// SetDout28 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr28Type) SetDout28(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr28FieldDout28Mask)|(uint32(value)<<RegisterMdios_doutr28FieldDout28Shift))
}

// registerMdios_doutr29Type MDIOS output data register 29
type registerMdios_doutr29Type uint32

const (
	RegisterMdios_doutr29FieldDout29Shift = 0
	RegisterMdios_doutr29FieldDout29Mask  = 0xffff
)

// GetDout29 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr29Type) GetDout29() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr29FieldDout29Mask) >> RegisterMdios_doutr29FieldDout29Shift)
}

// SetDout29 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr29Type) SetDout29(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr29FieldDout29Mask)|(uint32(value)<<RegisterMdios_doutr29FieldDout29Shift))
}

// registerMdios_doutr30Type MDIOS output data register 30
type registerMdios_doutr30Type uint32

const (
	RegisterMdios_doutr30FieldDout30Shift = 0
	RegisterMdios_doutr30FieldDout30Mask  = 0xffff
)

// GetDout30 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr30Type) GetDout30() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr30FieldDout30Mask) >> RegisterMdios_doutr30FieldDout30Shift)
}

// SetDout30 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr30Type) SetDout30(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr30FieldDout30Mask)|(uint32(value)<<RegisterMdios_doutr30FieldDout30Shift))
}

// registerMdios_doutr31Type MDIOS output data register 31
type registerMdios_doutr31Type uint32

const (
	RegisterMdios_doutr31FieldDout31Shift = 0
	RegisterMdios_doutr31FieldDout31Mask  = 0xffff
)

// GetDout31 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr31Type) GetDout31() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdios_doutr31FieldDout31Mask) >> RegisterMdios_doutr31FieldDout31Shift)
}

// SetDout31 Output data sent to MDIO Master during read frames
func (r *registerMdios_doutr31Type) SetDout31(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdios_doutr31FieldDout31Mask)|(uint32(value)<<RegisterMdios_doutr31FieldDout31Shift))
}
