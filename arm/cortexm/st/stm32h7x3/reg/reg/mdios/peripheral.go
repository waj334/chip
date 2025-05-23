package mdios

import (
	"unsafe"
	"volatile"
)

var (
	Mdios = (*_mdios)(unsafe.Pointer(uintptr(0x40009400)))
)

type _mdios struct {
	Cr      registerCrType
	Wrfr    registerWrfrType
	Cwrfr   registerCwrfrType
	Rdfr    registerRdfrType
	Crdfr   registerCrdfrType
	Sr      registerSrType
	Clrfr   registerClrfrType
	Dinr0   registerDinr0Type
	Dinr1   registerDinr1Type
	Dinr2   registerDinr2Type
	Dinr3   registerDinr3Type
	Dinr4   registerDinr4Type
	Dinr5   registerDinr5Type
	Dinr6   registerDinr6Type
	Dinr7   registerDinr7Type
	Dinr8   registerDinr8Type
	Dinr9   registerDinr9Type
	Dinr10  registerDinr10Type
	Dinr11  registerDinr11Type
	Dinr12  registerDinr12Type
	Dinr13  registerDinr13Type
	Dinr14  registerDinr14Type
	Dinr15  registerDinr15Type
	Dinr16  registerDinr16Type
	Dinr17  registerDinr17Type
	Dinr18  registerDinr18Type
	Dinr19  registerDinr19Type
	Dinr20  registerDinr20Type
	Dinr21  registerDinr21Type
	Dinr22  registerDinr22Type
	Dinr23  registerDinr23Type
	Dinr24  registerDinr24Type
	Dinr25  registerDinr25Type
	Dinr26  registerDinr26Type
	Dinr27  registerDinr27Type
	Dinr28  registerDinr28Type
	Dinr29  registerDinr29Type
	Dinr30  registerDinr30Type
	Dinr31  registerDinr31Type
	Doutr0  registerDoutr0Type
	Doutr1  registerDoutr1Type
	Doutr2  registerDoutr2Type
	Doutr3  registerDoutr3Type
	Doutr4  registerDoutr4Type
	Doutr5  registerDoutr5Type
	Doutr6  registerDoutr6Type
	Doutr7  registerDoutr7Type
	Doutr8  registerDoutr8Type
	Doutr9  registerDoutr9Type
	Doutr10 registerDoutr10Type
	Doutr11 registerDoutr11Type
	Doutr12 registerDoutr12Type
	Doutr13 registerDoutr13Type
	Doutr14 registerDoutr14Type
	Doutr15 registerDoutr15Type
	Doutr16 registerDoutr16Type
	Doutr17 registerDoutr17Type
	Doutr18 registerDoutr18Type
	Doutr19 registerDoutr19Type
	Doutr20 registerDoutr20Type
	Doutr21 registerDoutr21Type
	Doutr22 registerDoutr22Type
	Doutr23 registerDoutr23Type
	Doutr24 registerDoutr24Type
	Doutr25 registerDoutr25Type
	Doutr26 registerDoutr26Type
	Doutr27 registerDoutr27Type
	Doutr28 registerDoutr28Type
	Doutr29 registerDoutr29Type
	Doutr30 registerDoutr30Type
	Doutr31 registerDoutr31Type
}

// registerCrType MDIOS configuration register
type registerCrType uint32

const (
	RegisterCrFieldEnShift = 0
	RegisterCrFieldEnMask  = 0x1
)

// GetEn Peripheral enable
func (r *registerCrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEnMask) != 0
}

// SetEn Peripheral enable
func (r *registerCrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEnMask)
	}
}

const (
	RegisterCrFieldWrieShift = 1
	RegisterCrFieldWrieMask  = 0x2
)

// GetWrie Register write interrupt enable
func (r *registerCrType) GetWrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldWrieMask) != 0
}

// SetWrie Register write interrupt enable
func (r *registerCrType) SetWrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldWrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldWrieMask)
	}
}

const (
	RegisterCrFieldRdieShift = 2
	RegisterCrFieldRdieMask  = 0x4
)

// GetRdie Register Read Interrupt Enable
func (r *registerCrType) GetRdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRdieMask) != 0
}

// SetRdie Register Read Interrupt Enable
func (r *registerCrType) SetRdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRdieMask)
	}
}

const (
	RegisterCrFieldEieShift = 3
	RegisterCrFieldEieMask  = 0x8
)

// GetEie Error interrupt enable
func (r *registerCrType) GetEie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEieMask) != 0
}

// SetEie Error interrupt enable
func (r *registerCrType) SetEie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEieMask)
	}
}

const (
	RegisterCrFieldDpcShift = 7
	RegisterCrFieldDpcMask  = 0x80
)

// GetDpc Disable Preamble Check
func (r *registerCrType) GetDpc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDpcMask) != 0
}

// SetDpc Disable Preamble Check
func (r *registerCrType) SetDpc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDpcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDpcMask)
	}
}

const (
	RegisterCrFieldPort_addressShift = 8
	RegisterCrFieldPort_addressMask  = 0x1f00
)

// GetPort_address Slaves's address
func (r *registerCrType) GetPort_address() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPort_addressMask) >> RegisterCrFieldPort_addressShift)
}

// SetPort_address Slaves's address
func (r *registerCrType) SetPort_address(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPort_addressMask)|(uint32(value)<<RegisterCrFieldPort_addressShift))
}

// registerWrfrType MDIOS write flag register
type registerWrfrType uint32

const (
	RegisterWrfrFieldWrfShift = 0
	RegisterWrfrFieldWrfMask  = 0xffffffff
)

// GetWrf Write flags for MDIO registers 0 to 31
func (r *registerWrfrType) GetWrf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterWrfrFieldWrfMask) >> RegisterWrfrFieldWrfShift)
}

// SetWrf Write flags for MDIO registers 0 to 31
func (r *registerWrfrType) SetWrf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWrfrFieldWrfMask)|(uint32(value)<<RegisterWrfrFieldWrfShift))
}

// registerCwrfrType MDIOS clear write flag register
type registerCwrfrType uint32

const (
	RegisterCwrfrFieldCwrfShift = 0
	RegisterCwrfrFieldCwrfMask  = 0xffffffff
)

// GetCwrf Clear the write flag
func (r *registerCwrfrType) GetCwrf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCwrfrFieldCwrfMask) >> RegisterCwrfrFieldCwrfShift)
}

// SetCwrf Clear the write flag
func (r *registerCwrfrType) SetCwrf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCwrfrFieldCwrfMask)|(uint32(value)<<RegisterCwrfrFieldCwrfShift))
}

// registerRdfrType MDIOS read flag register
type registerRdfrType uint32

const (
	RegisterRdfrFieldRdfShift = 0
	RegisterRdfrFieldRdfMask  = 0xffffffff
)

// GetRdf Read flags for MDIO registers 0 to 31
func (r *registerRdfrType) GetRdf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRdfrFieldRdfMask) >> RegisterRdfrFieldRdfShift)
}

// SetRdf Read flags for MDIO registers 0 to 31
func (r *registerRdfrType) SetRdf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRdfrFieldRdfMask)|(uint32(value)<<RegisterRdfrFieldRdfShift))
}

// registerCrdfrType MDIOS clear read flag register
type registerCrdfrType uint32

const (
	RegisterCrdfrFieldCrdfShift = 0
	RegisterCrdfrFieldCrdfMask  = 0xffffffff
)

// GetCrdf Clear the read flag
func (r *registerCrdfrType) GetCrdf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrdfrFieldCrdfMask) >> RegisterCrdfrFieldCrdfShift)
}

// SetCrdf Clear the read flag
func (r *registerCrdfrType) SetCrdf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrdfrFieldCrdfMask)|(uint32(value)<<RegisterCrdfrFieldCrdfShift))
}

// registerSrType MDIOS status register
type registerSrType uint32

const (
	RegisterSrFieldPerfShift = 0
	RegisterSrFieldPerfMask  = 0x1
)

// GetPerf Preamble error flag
func (r *registerSrType) GetPerf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldPerfMask) != 0
}

// SetPerf Preamble error flag
func (r *registerSrType) SetPerf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldPerfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldPerfMask)
	}
}

const (
	RegisterSrFieldSerfShift = 1
	RegisterSrFieldSerfMask  = 0x2
)

// GetSerf Start error flag
func (r *registerSrType) GetSerf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSerfMask) != 0
}

// SetSerf Start error flag
func (r *registerSrType) SetSerf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldSerfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldSerfMask)
	}
}

const (
	RegisterSrFieldTerfShift = 2
	RegisterSrFieldTerfMask  = 0x4
)

// GetTerf Turnaround error flag
func (r *registerSrType) GetTerf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTerfMask) != 0
}

// SetTerf Turnaround error flag
func (r *registerSrType) SetTerf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldTerfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldTerfMask)
	}
}

// registerClrfrType MDIOS clear flag register
type registerClrfrType uint32

const (
	RegisterClrfrFieldCperfShift = 0
	RegisterClrfrFieldCperfMask  = 0x1
)

// GetCperf Clear the preamble error flag
func (r *registerClrfrType) GetCperf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterClrfrFieldCperfMask) != 0
}

// SetCperf Clear the preamble error flag
func (r *registerClrfrType) SetCperf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterClrfrFieldCperfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterClrfrFieldCperfMask)
	}
}

const (
	RegisterClrfrFieldCserfShift = 1
	RegisterClrfrFieldCserfMask  = 0x2
)

// GetCserf Clear the start error flag
func (r *registerClrfrType) GetCserf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterClrfrFieldCserfMask) != 0
}

// SetCserf Clear the start error flag
func (r *registerClrfrType) SetCserf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterClrfrFieldCserfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterClrfrFieldCserfMask)
	}
}

const (
	RegisterClrfrFieldCterfShift = 2
	RegisterClrfrFieldCterfMask  = 0x4
)

// GetCterf Clear the turnaround error flag
func (r *registerClrfrType) GetCterf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterClrfrFieldCterfMask) != 0
}

// SetCterf Clear the turnaround error flag
func (r *registerClrfrType) SetCterf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterClrfrFieldCterfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterClrfrFieldCterfMask)
	}
}

// registerDinr0Type MDIOS input data register 0
type registerDinr0Type uint32

const (
	RegisterDinr0FieldDin0Shift = 0
	RegisterDinr0FieldDin0Mask  = 0xffff
)

// GetDin0 Input data received from MDIO Master during write frames
func (r *registerDinr0Type) GetDin0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr0FieldDin0Mask) >> RegisterDinr0FieldDin0Shift)
}

// SetDin0 Input data received from MDIO Master during write frames
func (r *registerDinr0Type) SetDin0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr0FieldDin0Mask)|(uint32(value)<<RegisterDinr0FieldDin0Shift))
}

// registerDinr1Type MDIOS input data register 1
type registerDinr1Type uint32

const (
	RegisterDinr1FieldDin1Shift = 0
	RegisterDinr1FieldDin1Mask  = 0xffff
)

// GetDin1 Input data received from MDIO Master during write frames
func (r *registerDinr1Type) GetDin1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr1FieldDin1Mask) >> RegisterDinr1FieldDin1Shift)
}

// SetDin1 Input data received from MDIO Master during write frames
func (r *registerDinr1Type) SetDin1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr1FieldDin1Mask)|(uint32(value)<<RegisterDinr1FieldDin1Shift))
}

// registerDinr2Type MDIOS input data register 2
type registerDinr2Type uint32

const (
	RegisterDinr2FieldDin2Shift = 0
	RegisterDinr2FieldDin2Mask  = 0xffff
)

// GetDin2 Input data received from MDIO Master during write frames
func (r *registerDinr2Type) GetDin2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr2FieldDin2Mask) >> RegisterDinr2FieldDin2Shift)
}

// SetDin2 Input data received from MDIO Master during write frames
func (r *registerDinr2Type) SetDin2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr2FieldDin2Mask)|(uint32(value)<<RegisterDinr2FieldDin2Shift))
}

// registerDinr3Type MDIOS input data register 3
type registerDinr3Type uint32

const (
	RegisterDinr3FieldDin3Shift = 0
	RegisterDinr3FieldDin3Mask  = 0xffff
)

// GetDin3 Input data received from MDIO Master during write frames
func (r *registerDinr3Type) GetDin3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr3FieldDin3Mask) >> RegisterDinr3FieldDin3Shift)
}

// SetDin3 Input data received from MDIO Master during write frames
func (r *registerDinr3Type) SetDin3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr3FieldDin3Mask)|(uint32(value)<<RegisterDinr3FieldDin3Shift))
}

// registerDinr4Type MDIOS input data register 4
type registerDinr4Type uint32

const (
	RegisterDinr4FieldDin4Shift = 0
	RegisterDinr4FieldDin4Mask  = 0xffff
)

// GetDin4 Input data received from MDIO Master during write frames
func (r *registerDinr4Type) GetDin4() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr4FieldDin4Mask) >> RegisterDinr4FieldDin4Shift)
}

// SetDin4 Input data received from MDIO Master during write frames
func (r *registerDinr4Type) SetDin4(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr4FieldDin4Mask)|(uint32(value)<<RegisterDinr4FieldDin4Shift))
}

// registerDinr5Type MDIOS input data register 5
type registerDinr5Type uint32

const (
	RegisterDinr5FieldDin5Shift = 0
	RegisterDinr5FieldDin5Mask  = 0xffff
)

// GetDin5 Input data received from MDIO Master during write frames
func (r *registerDinr5Type) GetDin5() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr5FieldDin5Mask) >> RegisterDinr5FieldDin5Shift)
}

// SetDin5 Input data received from MDIO Master during write frames
func (r *registerDinr5Type) SetDin5(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr5FieldDin5Mask)|(uint32(value)<<RegisterDinr5FieldDin5Shift))
}

// registerDinr6Type MDIOS input data register 6
type registerDinr6Type uint32

const (
	RegisterDinr6FieldDin6Shift = 0
	RegisterDinr6FieldDin6Mask  = 0xffff
)

// GetDin6 Input data received from MDIO Master during write frames
func (r *registerDinr6Type) GetDin6() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr6FieldDin6Mask) >> RegisterDinr6FieldDin6Shift)
}

// SetDin6 Input data received from MDIO Master during write frames
func (r *registerDinr6Type) SetDin6(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr6FieldDin6Mask)|(uint32(value)<<RegisterDinr6FieldDin6Shift))
}

// registerDinr7Type MDIOS input data register 7
type registerDinr7Type uint32

const (
	RegisterDinr7FieldDin7Shift = 0
	RegisterDinr7FieldDin7Mask  = 0xffff
)

// GetDin7 Input data received from MDIO Master during write frames
func (r *registerDinr7Type) GetDin7() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr7FieldDin7Mask) >> RegisterDinr7FieldDin7Shift)
}

// SetDin7 Input data received from MDIO Master during write frames
func (r *registerDinr7Type) SetDin7(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr7FieldDin7Mask)|(uint32(value)<<RegisterDinr7FieldDin7Shift))
}

// registerDinr8Type MDIOS input data register 8
type registerDinr8Type uint32

const (
	RegisterDinr8FieldDin8Shift = 0
	RegisterDinr8FieldDin8Mask  = 0xffff
)

// GetDin8 Input data received from MDIO Master during write frames
func (r *registerDinr8Type) GetDin8() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr8FieldDin8Mask) >> RegisterDinr8FieldDin8Shift)
}

// SetDin8 Input data received from MDIO Master during write frames
func (r *registerDinr8Type) SetDin8(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr8FieldDin8Mask)|(uint32(value)<<RegisterDinr8FieldDin8Shift))
}

// registerDinr9Type MDIOS input data register 9
type registerDinr9Type uint32

const (
	RegisterDinr9FieldDin9Shift = 0
	RegisterDinr9FieldDin9Mask  = 0xffff
)

// GetDin9 Input data received from MDIO Master during write frames
func (r *registerDinr9Type) GetDin9() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr9FieldDin9Mask) >> RegisterDinr9FieldDin9Shift)
}

// SetDin9 Input data received from MDIO Master during write frames
func (r *registerDinr9Type) SetDin9(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr9FieldDin9Mask)|(uint32(value)<<RegisterDinr9FieldDin9Shift))
}

// registerDinr10Type MDIOS input data register 10
type registerDinr10Type uint32

const (
	RegisterDinr10FieldDin10Shift = 0
	RegisterDinr10FieldDin10Mask  = 0xffff
)

// GetDin10 Input data received from MDIO Master during write frames
func (r *registerDinr10Type) GetDin10() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr10FieldDin10Mask) >> RegisterDinr10FieldDin10Shift)
}

// SetDin10 Input data received from MDIO Master during write frames
func (r *registerDinr10Type) SetDin10(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr10FieldDin10Mask)|(uint32(value)<<RegisterDinr10FieldDin10Shift))
}

// registerDinr11Type MDIOS input data register 11
type registerDinr11Type uint32

const (
	RegisterDinr11FieldDin11Shift = 0
	RegisterDinr11FieldDin11Mask  = 0xffff
)

// GetDin11 Input data received from MDIO Master during write frames
func (r *registerDinr11Type) GetDin11() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr11FieldDin11Mask) >> RegisterDinr11FieldDin11Shift)
}

// SetDin11 Input data received from MDIO Master during write frames
func (r *registerDinr11Type) SetDin11(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr11FieldDin11Mask)|(uint32(value)<<RegisterDinr11FieldDin11Shift))
}

// registerDinr12Type MDIOS input data register 12
type registerDinr12Type uint32

const (
	RegisterDinr12FieldDin12Shift = 0
	RegisterDinr12FieldDin12Mask  = 0xffff
)

// GetDin12 Input data received from MDIO Master during write frames
func (r *registerDinr12Type) GetDin12() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr12FieldDin12Mask) >> RegisterDinr12FieldDin12Shift)
}

// SetDin12 Input data received from MDIO Master during write frames
func (r *registerDinr12Type) SetDin12(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr12FieldDin12Mask)|(uint32(value)<<RegisterDinr12FieldDin12Shift))
}

// registerDinr13Type MDIOS input data register 13
type registerDinr13Type uint32

const (
	RegisterDinr13FieldDin13Shift = 0
	RegisterDinr13FieldDin13Mask  = 0xffff
)

// GetDin13 Input data received from MDIO Master during write frames
func (r *registerDinr13Type) GetDin13() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr13FieldDin13Mask) >> RegisterDinr13FieldDin13Shift)
}

// SetDin13 Input data received from MDIO Master during write frames
func (r *registerDinr13Type) SetDin13(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr13FieldDin13Mask)|(uint32(value)<<RegisterDinr13FieldDin13Shift))
}

// registerDinr14Type MDIOS input data register 14
type registerDinr14Type uint32

const (
	RegisterDinr14FieldDin14Shift = 0
	RegisterDinr14FieldDin14Mask  = 0xffff
)

// GetDin14 Input data received from MDIO Master during write frames
func (r *registerDinr14Type) GetDin14() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr14FieldDin14Mask) >> RegisterDinr14FieldDin14Shift)
}

// SetDin14 Input data received from MDIO Master during write frames
func (r *registerDinr14Type) SetDin14(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr14FieldDin14Mask)|(uint32(value)<<RegisterDinr14FieldDin14Shift))
}

// registerDinr15Type MDIOS input data register 15
type registerDinr15Type uint32

const (
	RegisterDinr15FieldDin15Shift = 0
	RegisterDinr15FieldDin15Mask  = 0xffff
)

// GetDin15 Input data received from MDIO Master during write frames
func (r *registerDinr15Type) GetDin15() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr15FieldDin15Mask) >> RegisterDinr15FieldDin15Shift)
}

// SetDin15 Input data received from MDIO Master during write frames
func (r *registerDinr15Type) SetDin15(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr15FieldDin15Mask)|(uint32(value)<<RegisterDinr15FieldDin15Shift))
}

// registerDinr16Type MDIOS input data register 16
type registerDinr16Type uint32

const (
	RegisterDinr16FieldDin16Shift = 0
	RegisterDinr16FieldDin16Mask  = 0xffff
)

// GetDin16 Input data received from MDIO Master during write frames
func (r *registerDinr16Type) GetDin16() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr16FieldDin16Mask) >> RegisterDinr16FieldDin16Shift)
}

// SetDin16 Input data received from MDIO Master during write frames
func (r *registerDinr16Type) SetDin16(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr16FieldDin16Mask)|(uint32(value)<<RegisterDinr16FieldDin16Shift))
}

// registerDinr17Type MDIOS input data register 17
type registerDinr17Type uint32

const (
	RegisterDinr17FieldDin17Shift = 0
	RegisterDinr17FieldDin17Mask  = 0xffff
)

// GetDin17 Input data received from MDIO Master during write frames
func (r *registerDinr17Type) GetDin17() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr17FieldDin17Mask) >> RegisterDinr17FieldDin17Shift)
}

// SetDin17 Input data received from MDIO Master during write frames
func (r *registerDinr17Type) SetDin17(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr17FieldDin17Mask)|(uint32(value)<<RegisterDinr17FieldDin17Shift))
}

// registerDinr18Type MDIOS input data register 18
type registerDinr18Type uint32

const (
	RegisterDinr18FieldDin18Shift = 0
	RegisterDinr18FieldDin18Mask  = 0xffff
)

// GetDin18 Input data received from MDIO Master during write frames
func (r *registerDinr18Type) GetDin18() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr18FieldDin18Mask) >> RegisterDinr18FieldDin18Shift)
}

// SetDin18 Input data received from MDIO Master during write frames
func (r *registerDinr18Type) SetDin18(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr18FieldDin18Mask)|(uint32(value)<<RegisterDinr18FieldDin18Shift))
}

// registerDinr19Type MDIOS input data register 19
type registerDinr19Type uint32

const (
	RegisterDinr19FieldDin19Shift = 0
	RegisterDinr19FieldDin19Mask  = 0xffff
)

// GetDin19 Input data received from MDIO Master during write frames
func (r *registerDinr19Type) GetDin19() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr19FieldDin19Mask) >> RegisterDinr19FieldDin19Shift)
}

// SetDin19 Input data received from MDIO Master during write frames
func (r *registerDinr19Type) SetDin19(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr19FieldDin19Mask)|(uint32(value)<<RegisterDinr19FieldDin19Shift))
}

// registerDinr20Type MDIOS input data register 20
type registerDinr20Type uint32

const (
	RegisterDinr20FieldDin20Shift = 0
	RegisterDinr20FieldDin20Mask  = 0xffff
)

// GetDin20 Input data received from MDIO Master during write frames
func (r *registerDinr20Type) GetDin20() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr20FieldDin20Mask) >> RegisterDinr20FieldDin20Shift)
}

// SetDin20 Input data received from MDIO Master during write frames
func (r *registerDinr20Type) SetDin20(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr20FieldDin20Mask)|(uint32(value)<<RegisterDinr20FieldDin20Shift))
}

// registerDinr21Type MDIOS input data register 21
type registerDinr21Type uint32

const (
	RegisterDinr21FieldDin21Shift = 0
	RegisterDinr21FieldDin21Mask  = 0xffff
)

// GetDin21 Input data received from MDIO Master during write frames
func (r *registerDinr21Type) GetDin21() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr21FieldDin21Mask) >> RegisterDinr21FieldDin21Shift)
}

// SetDin21 Input data received from MDIO Master during write frames
func (r *registerDinr21Type) SetDin21(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr21FieldDin21Mask)|(uint32(value)<<RegisterDinr21FieldDin21Shift))
}

// registerDinr22Type MDIOS input data register 22
type registerDinr22Type uint32

const (
	RegisterDinr22FieldDin22Shift = 0
	RegisterDinr22FieldDin22Mask  = 0xffff
)

// GetDin22 Input data received from MDIO Master during write frames
func (r *registerDinr22Type) GetDin22() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr22FieldDin22Mask) >> RegisterDinr22FieldDin22Shift)
}

// SetDin22 Input data received from MDIO Master during write frames
func (r *registerDinr22Type) SetDin22(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr22FieldDin22Mask)|(uint32(value)<<RegisterDinr22FieldDin22Shift))
}

// registerDinr23Type MDIOS input data register 23
type registerDinr23Type uint32

const (
	RegisterDinr23FieldDin23Shift = 0
	RegisterDinr23FieldDin23Mask  = 0xffff
)

// GetDin23 Input data received from MDIO Master during write frames
func (r *registerDinr23Type) GetDin23() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr23FieldDin23Mask) >> RegisterDinr23FieldDin23Shift)
}

// SetDin23 Input data received from MDIO Master during write frames
func (r *registerDinr23Type) SetDin23(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr23FieldDin23Mask)|(uint32(value)<<RegisterDinr23FieldDin23Shift))
}

// registerDinr24Type MDIOS input data register 24
type registerDinr24Type uint32

const (
	RegisterDinr24FieldDin24Shift = 0
	RegisterDinr24FieldDin24Mask  = 0xffff
)

// GetDin24 Input data received from MDIO Master during write frames
func (r *registerDinr24Type) GetDin24() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr24FieldDin24Mask) >> RegisterDinr24FieldDin24Shift)
}

// SetDin24 Input data received from MDIO Master during write frames
func (r *registerDinr24Type) SetDin24(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr24FieldDin24Mask)|(uint32(value)<<RegisterDinr24FieldDin24Shift))
}

// registerDinr25Type MDIOS input data register 25
type registerDinr25Type uint32

const (
	RegisterDinr25FieldDin25Shift = 0
	RegisterDinr25FieldDin25Mask  = 0xffff
)

// GetDin25 Input data received from MDIO Master during write frames
func (r *registerDinr25Type) GetDin25() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr25FieldDin25Mask) >> RegisterDinr25FieldDin25Shift)
}

// SetDin25 Input data received from MDIO Master during write frames
func (r *registerDinr25Type) SetDin25(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr25FieldDin25Mask)|(uint32(value)<<RegisterDinr25FieldDin25Shift))
}

// registerDinr26Type MDIOS input data register 26
type registerDinr26Type uint32

const (
	RegisterDinr26FieldDin26Shift = 0
	RegisterDinr26FieldDin26Mask  = 0xffff
)

// GetDin26 Input data received from MDIO Master during write frames
func (r *registerDinr26Type) GetDin26() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr26FieldDin26Mask) >> RegisterDinr26FieldDin26Shift)
}

// SetDin26 Input data received from MDIO Master during write frames
func (r *registerDinr26Type) SetDin26(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr26FieldDin26Mask)|(uint32(value)<<RegisterDinr26FieldDin26Shift))
}

// registerDinr27Type MDIOS input data register 27
type registerDinr27Type uint32

const (
	RegisterDinr27FieldDin27Shift = 0
	RegisterDinr27FieldDin27Mask  = 0xffff
)

// GetDin27 Input data received from MDIO Master during write frames
func (r *registerDinr27Type) GetDin27() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr27FieldDin27Mask) >> RegisterDinr27FieldDin27Shift)
}

// SetDin27 Input data received from MDIO Master during write frames
func (r *registerDinr27Type) SetDin27(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr27FieldDin27Mask)|(uint32(value)<<RegisterDinr27FieldDin27Shift))
}

// registerDinr28Type MDIOS input data register 28
type registerDinr28Type uint32

const (
	RegisterDinr28FieldDin28Shift = 0
	RegisterDinr28FieldDin28Mask  = 0xffff
)

// GetDin28 Input data received from MDIO Master during write frames
func (r *registerDinr28Type) GetDin28() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr28FieldDin28Mask) >> RegisterDinr28FieldDin28Shift)
}

// SetDin28 Input data received from MDIO Master during write frames
func (r *registerDinr28Type) SetDin28(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr28FieldDin28Mask)|(uint32(value)<<RegisterDinr28FieldDin28Shift))
}

// registerDinr29Type MDIOS input data register 29
type registerDinr29Type uint32

const (
	RegisterDinr29FieldDin29Shift = 0
	RegisterDinr29FieldDin29Mask  = 0xffff
)

// GetDin29 Input data received from MDIO Master during write frames
func (r *registerDinr29Type) GetDin29() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr29FieldDin29Mask) >> RegisterDinr29FieldDin29Shift)
}

// SetDin29 Input data received from MDIO Master during write frames
func (r *registerDinr29Type) SetDin29(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr29FieldDin29Mask)|(uint32(value)<<RegisterDinr29FieldDin29Shift))
}

// registerDinr30Type MDIOS input data register 30
type registerDinr30Type uint32

const (
	RegisterDinr30FieldDin30Shift = 0
	RegisterDinr30FieldDin30Mask  = 0xffff
)

// GetDin30 Input data received from MDIO Master during write frames
func (r *registerDinr30Type) GetDin30() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr30FieldDin30Mask) >> RegisterDinr30FieldDin30Shift)
}

// SetDin30 Input data received from MDIO Master during write frames
func (r *registerDinr30Type) SetDin30(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr30FieldDin30Mask)|(uint32(value)<<RegisterDinr30FieldDin30Shift))
}

// registerDinr31Type MDIOS input data register 31
type registerDinr31Type uint32

const (
	RegisterDinr31FieldDin31Shift = 0
	RegisterDinr31FieldDin31Mask  = 0xffff
)

// GetDin31 Input data received from MDIO Master during write frames
func (r *registerDinr31Type) GetDin31() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDinr31FieldDin31Mask) >> RegisterDinr31FieldDin31Shift)
}

// SetDin31 Input data received from MDIO Master during write frames
func (r *registerDinr31Type) SetDin31(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinr31FieldDin31Mask)|(uint32(value)<<RegisterDinr31FieldDin31Shift))
}

// registerDoutr0Type MDIOS output data register 0
type registerDoutr0Type uint32

const (
	RegisterDoutr0FieldDout0Shift = 0
	RegisterDoutr0FieldDout0Mask  = 0xffff
)

// GetDout0 Output data sent to MDIO Master during read frames
func (r *registerDoutr0Type) GetDout0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr0FieldDout0Mask) >> RegisterDoutr0FieldDout0Shift)
}

// SetDout0 Output data sent to MDIO Master during read frames
func (r *registerDoutr0Type) SetDout0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr0FieldDout0Mask)|(uint32(value)<<RegisterDoutr0FieldDout0Shift))
}

// registerDoutr1Type MDIOS output data register 1
type registerDoutr1Type uint32

const (
	RegisterDoutr1FieldDout1Shift = 0
	RegisterDoutr1FieldDout1Mask  = 0xffff
)

// GetDout1 Output data sent to MDIO Master during read frames
func (r *registerDoutr1Type) GetDout1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr1FieldDout1Mask) >> RegisterDoutr1FieldDout1Shift)
}

// SetDout1 Output data sent to MDIO Master during read frames
func (r *registerDoutr1Type) SetDout1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr1FieldDout1Mask)|(uint32(value)<<RegisterDoutr1FieldDout1Shift))
}

// registerDoutr2Type MDIOS output data register 2
type registerDoutr2Type uint32

const (
	RegisterDoutr2FieldDout2Shift = 0
	RegisterDoutr2FieldDout2Mask  = 0xffff
)

// GetDout2 Output data sent to MDIO Master during read frames
func (r *registerDoutr2Type) GetDout2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr2FieldDout2Mask) >> RegisterDoutr2FieldDout2Shift)
}

// SetDout2 Output data sent to MDIO Master during read frames
func (r *registerDoutr2Type) SetDout2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr2FieldDout2Mask)|(uint32(value)<<RegisterDoutr2FieldDout2Shift))
}

// registerDoutr3Type MDIOS output data register 3
type registerDoutr3Type uint32

const (
	RegisterDoutr3FieldDout3Shift = 0
	RegisterDoutr3FieldDout3Mask  = 0xffff
)

// GetDout3 Output data sent to MDIO Master during read frames
func (r *registerDoutr3Type) GetDout3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr3FieldDout3Mask) >> RegisterDoutr3FieldDout3Shift)
}

// SetDout3 Output data sent to MDIO Master during read frames
func (r *registerDoutr3Type) SetDout3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr3FieldDout3Mask)|(uint32(value)<<RegisterDoutr3FieldDout3Shift))
}

// registerDoutr4Type MDIOS output data register 4
type registerDoutr4Type uint32

const (
	RegisterDoutr4FieldDout4Shift = 0
	RegisterDoutr4FieldDout4Mask  = 0xffff
)

// GetDout4 Output data sent to MDIO Master during read frames
func (r *registerDoutr4Type) GetDout4() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr4FieldDout4Mask) >> RegisterDoutr4FieldDout4Shift)
}

// SetDout4 Output data sent to MDIO Master during read frames
func (r *registerDoutr4Type) SetDout4(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr4FieldDout4Mask)|(uint32(value)<<RegisterDoutr4FieldDout4Shift))
}

// registerDoutr5Type MDIOS output data register 5
type registerDoutr5Type uint32

const (
	RegisterDoutr5FieldDout5Shift = 0
	RegisterDoutr5FieldDout5Mask  = 0xffff
)

// GetDout5 Output data sent to MDIO Master during read frames
func (r *registerDoutr5Type) GetDout5() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr5FieldDout5Mask) >> RegisterDoutr5FieldDout5Shift)
}

// SetDout5 Output data sent to MDIO Master during read frames
func (r *registerDoutr5Type) SetDout5(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr5FieldDout5Mask)|(uint32(value)<<RegisterDoutr5FieldDout5Shift))
}

// registerDoutr6Type MDIOS output data register 6
type registerDoutr6Type uint32

const (
	RegisterDoutr6FieldDout6Shift = 0
	RegisterDoutr6FieldDout6Mask  = 0xffff
)

// GetDout6 Output data sent to MDIO Master during read frames
func (r *registerDoutr6Type) GetDout6() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr6FieldDout6Mask) >> RegisterDoutr6FieldDout6Shift)
}

// SetDout6 Output data sent to MDIO Master during read frames
func (r *registerDoutr6Type) SetDout6(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr6FieldDout6Mask)|(uint32(value)<<RegisterDoutr6FieldDout6Shift))
}

// registerDoutr7Type MDIOS output data register 7
type registerDoutr7Type uint32

const (
	RegisterDoutr7FieldDout7Shift = 0
	RegisterDoutr7FieldDout7Mask  = 0xffff
)

// GetDout7 Output data sent to MDIO Master during read frames
func (r *registerDoutr7Type) GetDout7() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr7FieldDout7Mask) >> RegisterDoutr7FieldDout7Shift)
}

// SetDout7 Output data sent to MDIO Master during read frames
func (r *registerDoutr7Type) SetDout7(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr7FieldDout7Mask)|(uint32(value)<<RegisterDoutr7FieldDout7Shift))
}

// registerDoutr8Type MDIOS output data register 8
type registerDoutr8Type uint32

const (
	RegisterDoutr8FieldDout8Shift = 0
	RegisterDoutr8FieldDout8Mask  = 0xffff
)

// GetDout8 Output data sent to MDIO Master during read frames
func (r *registerDoutr8Type) GetDout8() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr8FieldDout8Mask) >> RegisterDoutr8FieldDout8Shift)
}

// SetDout8 Output data sent to MDIO Master during read frames
func (r *registerDoutr8Type) SetDout8(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr8FieldDout8Mask)|(uint32(value)<<RegisterDoutr8FieldDout8Shift))
}

// registerDoutr9Type MDIOS output data register 9
type registerDoutr9Type uint32

const (
	RegisterDoutr9FieldDout9Shift = 0
	RegisterDoutr9FieldDout9Mask  = 0xffff
)

// GetDout9 Output data sent to MDIO Master during read frames
func (r *registerDoutr9Type) GetDout9() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr9FieldDout9Mask) >> RegisterDoutr9FieldDout9Shift)
}

// SetDout9 Output data sent to MDIO Master during read frames
func (r *registerDoutr9Type) SetDout9(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr9FieldDout9Mask)|(uint32(value)<<RegisterDoutr9FieldDout9Shift))
}

// registerDoutr10Type MDIOS output data register 10
type registerDoutr10Type uint32

const (
	RegisterDoutr10FieldDout10Shift = 0
	RegisterDoutr10FieldDout10Mask  = 0xffff
)

// GetDout10 Output data sent to MDIO Master during read frames
func (r *registerDoutr10Type) GetDout10() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr10FieldDout10Mask) >> RegisterDoutr10FieldDout10Shift)
}

// SetDout10 Output data sent to MDIO Master during read frames
func (r *registerDoutr10Type) SetDout10(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr10FieldDout10Mask)|(uint32(value)<<RegisterDoutr10FieldDout10Shift))
}

// registerDoutr11Type MDIOS output data register 11
type registerDoutr11Type uint32

const (
	RegisterDoutr11FieldDout11Shift = 0
	RegisterDoutr11FieldDout11Mask  = 0xffff
)

// GetDout11 Output data sent to MDIO Master during read frames
func (r *registerDoutr11Type) GetDout11() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr11FieldDout11Mask) >> RegisterDoutr11FieldDout11Shift)
}

// SetDout11 Output data sent to MDIO Master during read frames
func (r *registerDoutr11Type) SetDout11(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr11FieldDout11Mask)|(uint32(value)<<RegisterDoutr11FieldDout11Shift))
}

// registerDoutr12Type MDIOS output data register 12
type registerDoutr12Type uint32

const (
	RegisterDoutr12FieldDout12Shift = 0
	RegisterDoutr12FieldDout12Mask  = 0xffff
)

// GetDout12 Output data sent to MDIO Master during read frames
func (r *registerDoutr12Type) GetDout12() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr12FieldDout12Mask) >> RegisterDoutr12FieldDout12Shift)
}

// SetDout12 Output data sent to MDIO Master during read frames
func (r *registerDoutr12Type) SetDout12(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr12FieldDout12Mask)|(uint32(value)<<RegisterDoutr12FieldDout12Shift))
}

// registerDoutr13Type MDIOS output data register 13
type registerDoutr13Type uint32

const (
	RegisterDoutr13FieldDout13Shift = 0
	RegisterDoutr13FieldDout13Mask  = 0xffff
)

// GetDout13 Output data sent to MDIO Master during read frames
func (r *registerDoutr13Type) GetDout13() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr13FieldDout13Mask) >> RegisterDoutr13FieldDout13Shift)
}

// SetDout13 Output data sent to MDIO Master during read frames
func (r *registerDoutr13Type) SetDout13(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr13FieldDout13Mask)|(uint32(value)<<RegisterDoutr13FieldDout13Shift))
}

// registerDoutr14Type MDIOS output data register 14
type registerDoutr14Type uint32

const (
	RegisterDoutr14FieldDout14Shift = 0
	RegisterDoutr14FieldDout14Mask  = 0xffff
)

// GetDout14 Output data sent to MDIO Master during read frames
func (r *registerDoutr14Type) GetDout14() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr14FieldDout14Mask) >> RegisterDoutr14FieldDout14Shift)
}

// SetDout14 Output data sent to MDIO Master during read frames
func (r *registerDoutr14Type) SetDout14(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr14FieldDout14Mask)|(uint32(value)<<RegisterDoutr14FieldDout14Shift))
}

// registerDoutr15Type MDIOS output data register 15
type registerDoutr15Type uint32

const (
	RegisterDoutr15FieldDout15Shift = 0
	RegisterDoutr15FieldDout15Mask  = 0xffff
)

// GetDout15 Output data sent to MDIO Master during read frames
func (r *registerDoutr15Type) GetDout15() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr15FieldDout15Mask) >> RegisterDoutr15FieldDout15Shift)
}

// SetDout15 Output data sent to MDIO Master during read frames
func (r *registerDoutr15Type) SetDout15(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr15FieldDout15Mask)|(uint32(value)<<RegisterDoutr15FieldDout15Shift))
}

// registerDoutr16Type MDIOS output data register 16
type registerDoutr16Type uint32

const (
	RegisterDoutr16FieldDout16Shift = 0
	RegisterDoutr16FieldDout16Mask  = 0xffff
)

// GetDout16 Output data sent to MDIO Master during read frames
func (r *registerDoutr16Type) GetDout16() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr16FieldDout16Mask) >> RegisterDoutr16FieldDout16Shift)
}

// SetDout16 Output data sent to MDIO Master during read frames
func (r *registerDoutr16Type) SetDout16(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr16FieldDout16Mask)|(uint32(value)<<RegisterDoutr16FieldDout16Shift))
}

// registerDoutr17Type MDIOS output data register 17
type registerDoutr17Type uint32

const (
	RegisterDoutr17FieldDout17Shift = 0
	RegisterDoutr17FieldDout17Mask  = 0xffff
)

// GetDout17 Output data sent to MDIO Master during read frames
func (r *registerDoutr17Type) GetDout17() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr17FieldDout17Mask) >> RegisterDoutr17FieldDout17Shift)
}

// SetDout17 Output data sent to MDIO Master during read frames
func (r *registerDoutr17Type) SetDout17(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr17FieldDout17Mask)|(uint32(value)<<RegisterDoutr17FieldDout17Shift))
}

// registerDoutr18Type MDIOS output data register 18
type registerDoutr18Type uint32

const (
	RegisterDoutr18FieldDout18Shift = 0
	RegisterDoutr18FieldDout18Mask  = 0xffff
)

// GetDout18 Output data sent to MDIO Master during read frames
func (r *registerDoutr18Type) GetDout18() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr18FieldDout18Mask) >> RegisterDoutr18FieldDout18Shift)
}

// SetDout18 Output data sent to MDIO Master during read frames
func (r *registerDoutr18Type) SetDout18(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr18FieldDout18Mask)|(uint32(value)<<RegisterDoutr18FieldDout18Shift))
}

// registerDoutr19Type MDIOS output data register 19
type registerDoutr19Type uint32

const (
	RegisterDoutr19FieldDout19Shift = 0
	RegisterDoutr19FieldDout19Mask  = 0xffff
)

// GetDout19 Output data sent to MDIO Master during read frames
func (r *registerDoutr19Type) GetDout19() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr19FieldDout19Mask) >> RegisterDoutr19FieldDout19Shift)
}

// SetDout19 Output data sent to MDIO Master during read frames
func (r *registerDoutr19Type) SetDout19(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr19FieldDout19Mask)|(uint32(value)<<RegisterDoutr19FieldDout19Shift))
}

// registerDoutr20Type MDIOS output data register 20
type registerDoutr20Type uint32

const (
	RegisterDoutr20FieldDout20Shift = 0
	RegisterDoutr20FieldDout20Mask  = 0xffff
)

// GetDout20 Output data sent to MDIO Master during read frames
func (r *registerDoutr20Type) GetDout20() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr20FieldDout20Mask) >> RegisterDoutr20FieldDout20Shift)
}

// SetDout20 Output data sent to MDIO Master during read frames
func (r *registerDoutr20Type) SetDout20(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr20FieldDout20Mask)|(uint32(value)<<RegisterDoutr20FieldDout20Shift))
}

// registerDoutr21Type MDIOS output data register 21
type registerDoutr21Type uint32

const (
	RegisterDoutr21FieldDout21Shift = 0
	RegisterDoutr21FieldDout21Mask  = 0xffff
)

// GetDout21 Output data sent to MDIO Master during read frames
func (r *registerDoutr21Type) GetDout21() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr21FieldDout21Mask) >> RegisterDoutr21FieldDout21Shift)
}

// SetDout21 Output data sent to MDIO Master during read frames
func (r *registerDoutr21Type) SetDout21(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr21FieldDout21Mask)|(uint32(value)<<RegisterDoutr21FieldDout21Shift))
}

// registerDoutr22Type MDIOS output data register 22
type registerDoutr22Type uint32

const (
	RegisterDoutr22FieldDout22Shift = 0
	RegisterDoutr22FieldDout22Mask  = 0xffff
)

// GetDout22 Output data sent to MDIO Master during read frames
func (r *registerDoutr22Type) GetDout22() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr22FieldDout22Mask) >> RegisterDoutr22FieldDout22Shift)
}

// SetDout22 Output data sent to MDIO Master during read frames
func (r *registerDoutr22Type) SetDout22(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr22FieldDout22Mask)|(uint32(value)<<RegisterDoutr22FieldDout22Shift))
}

// registerDoutr23Type MDIOS output data register 23
type registerDoutr23Type uint32

const (
	RegisterDoutr23FieldDout23Shift = 0
	RegisterDoutr23FieldDout23Mask  = 0xffff
)

// GetDout23 Output data sent to MDIO Master during read frames
func (r *registerDoutr23Type) GetDout23() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr23FieldDout23Mask) >> RegisterDoutr23FieldDout23Shift)
}

// SetDout23 Output data sent to MDIO Master during read frames
func (r *registerDoutr23Type) SetDout23(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr23FieldDout23Mask)|(uint32(value)<<RegisterDoutr23FieldDout23Shift))
}

// registerDoutr24Type MDIOS output data register 24
type registerDoutr24Type uint32

const (
	RegisterDoutr24FieldDout24Shift = 0
	RegisterDoutr24FieldDout24Mask  = 0xffff
)

// GetDout24 Output data sent to MDIO Master during read frames
func (r *registerDoutr24Type) GetDout24() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr24FieldDout24Mask) >> RegisterDoutr24FieldDout24Shift)
}

// SetDout24 Output data sent to MDIO Master during read frames
func (r *registerDoutr24Type) SetDout24(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr24FieldDout24Mask)|(uint32(value)<<RegisterDoutr24FieldDout24Shift))
}

// registerDoutr25Type MDIOS output data register 25
type registerDoutr25Type uint32

const (
	RegisterDoutr25FieldDout25Shift = 0
	RegisterDoutr25FieldDout25Mask  = 0xffff
)

// GetDout25 Output data sent to MDIO Master during read frames
func (r *registerDoutr25Type) GetDout25() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr25FieldDout25Mask) >> RegisterDoutr25FieldDout25Shift)
}

// SetDout25 Output data sent to MDIO Master during read frames
func (r *registerDoutr25Type) SetDout25(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr25FieldDout25Mask)|(uint32(value)<<RegisterDoutr25FieldDout25Shift))
}

// registerDoutr26Type MDIOS output data register 26
type registerDoutr26Type uint32

const (
	RegisterDoutr26FieldDout26Shift = 0
	RegisterDoutr26FieldDout26Mask  = 0xffff
)

// GetDout26 Output data sent to MDIO Master during read frames
func (r *registerDoutr26Type) GetDout26() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr26FieldDout26Mask) >> RegisterDoutr26FieldDout26Shift)
}

// SetDout26 Output data sent to MDIO Master during read frames
func (r *registerDoutr26Type) SetDout26(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr26FieldDout26Mask)|(uint32(value)<<RegisterDoutr26FieldDout26Shift))
}

// registerDoutr27Type MDIOS output data register 27
type registerDoutr27Type uint32

const (
	RegisterDoutr27FieldDout27Shift = 0
	RegisterDoutr27FieldDout27Mask  = 0xffff
)

// GetDout27 Output data sent to MDIO Master during read frames
func (r *registerDoutr27Type) GetDout27() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr27FieldDout27Mask) >> RegisterDoutr27FieldDout27Shift)
}

// SetDout27 Output data sent to MDIO Master during read frames
func (r *registerDoutr27Type) SetDout27(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr27FieldDout27Mask)|(uint32(value)<<RegisterDoutr27FieldDout27Shift))
}

// registerDoutr28Type MDIOS output data register 28
type registerDoutr28Type uint32

const (
	RegisterDoutr28FieldDout28Shift = 0
	RegisterDoutr28FieldDout28Mask  = 0xffff
)

// GetDout28 Output data sent to MDIO Master during read frames
func (r *registerDoutr28Type) GetDout28() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr28FieldDout28Mask) >> RegisterDoutr28FieldDout28Shift)
}

// SetDout28 Output data sent to MDIO Master during read frames
func (r *registerDoutr28Type) SetDout28(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr28FieldDout28Mask)|(uint32(value)<<RegisterDoutr28FieldDout28Shift))
}

// registerDoutr29Type MDIOS output data register 29
type registerDoutr29Type uint32

const (
	RegisterDoutr29FieldDout29Shift = 0
	RegisterDoutr29FieldDout29Mask  = 0xffff
)

// GetDout29 Output data sent to MDIO Master during read frames
func (r *registerDoutr29Type) GetDout29() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr29FieldDout29Mask) >> RegisterDoutr29FieldDout29Shift)
}

// SetDout29 Output data sent to MDIO Master during read frames
func (r *registerDoutr29Type) SetDout29(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr29FieldDout29Mask)|(uint32(value)<<RegisterDoutr29FieldDout29Shift))
}

// registerDoutr30Type MDIOS output data register 30
type registerDoutr30Type uint32

const (
	RegisterDoutr30FieldDout30Shift = 0
	RegisterDoutr30FieldDout30Mask  = 0xffff
)

// GetDout30 Output data sent to MDIO Master during read frames
func (r *registerDoutr30Type) GetDout30() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr30FieldDout30Mask) >> RegisterDoutr30FieldDout30Shift)
}

// SetDout30 Output data sent to MDIO Master during read frames
func (r *registerDoutr30Type) SetDout30(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr30FieldDout30Mask)|(uint32(value)<<RegisterDoutr30FieldDout30Shift))
}

// registerDoutr31Type MDIOS output data register 31
type registerDoutr31Type uint32

const (
	RegisterDoutr31FieldDout31Shift = 0
	RegisterDoutr31FieldDout31Mask  = 0xffff
)

// GetDout31 Output data sent to MDIO Master during read frames
func (r *registerDoutr31Type) GetDout31() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDoutr31FieldDout31Mask) >> RegisterDoutr31FieldDout31Shift)
}

// SetDout31 Output data sent to MDIO Master during read frames
func (r *registerDoutr31Type) SetDout31(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutr31FieldDout31Mask)|(uint32(value)<<RegisterDoutr31FieldDout31Shift))
}
