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
	Mdioscr      registerMdioscrType
	Mdioswrfr    registerMdioswrfrType
	Mdioscwrfr   registerMdioscwrfrType
	Mdiosrdfr    registerMdiosrdfrType
	Mdioscrdfr   registerMdioscrdfrType
	Mdiossr      registerMdiossrType
	Mdiosclrfr   registerMdiosclrfrType
	Mdiosdinr0   registerMdiosdinr0Type
	Mdiosdinr1   registerMdiosdinr1Type
	Mdiosdinr2   registerMdiosdinr2Type
	Mdiosdinr3   registerMdiosdinr3Type
	Mdiosdinr4   registerMdiosdinr4Type
	Mdiosdinr5   registerMdiosdinr5Type
	Mdiosdinr6   registerMdiosdinr6Type
	Mdiosdinr7   registerMdiosdinr7Type
	Mdiosdinr8   registerMdiosdinr8Type
	Mdiosdinr9   registerMdiosdinr9Type
	Mdiosdinr10  registerMdiosdinr10Type
	Mdiosdinr11  registerMdiosdinr11Type
	Mdiosdinr12  registerMdiosdinr12Type
	Mdiosdinr13  registerMdiosdinr13Type
	Mdiosdinr14  registerMdiosdinr14Type
	Mdiosdinr15  registerMdiosdinr15Type
	Mdiosdinr16  registerMdiosdinr16Type
	Mdiosdinr17  registerMdiosdinr17Type
	Mdiosdinr18  registerMdiosdinr18Type
	Mdiosdinr19  registerMdiosdinr19Type
	Mdiosdinr20  registerMdiosdinr20Type
	Mdiosdinr21  registerMdiosdinr21Type
	Mdiosdinr22  registerMdiosdinr22Type
	Mdiosdinr23  registerMdiosdinr23Type
	Mdiosdinr24  registerMdiosdinr24Type
	Mdiosdinr25  registerMdiosdinr25Type
	Mdiosdinr26  registerMdiosdinr26Type
	Mdiosdinr27  registerMdiosdinr27Type
	Mdiosdinr28  registerMdiosdinr28Type
	Mdiosdinr29  registerMdiosdinr29Type
	Mdiosdinr30  registerMdiosdinr30Type
	Mdiosdinr31  registerMdiosdinr31Type
	Mdiosdoutr0  registerMdiosdoutr0Type
	Mdiosdoutr1  registerMdiosdoutr1Type
	Mdiosdoutr2  registerMdiosdoutr2Type
	Mdiosdoutr3  registerMdiosdoutr3Type
	Mdiosdoutr4  registerMdiosdoutr4Type
	Mdiosdoutr5  registerMdiosdoutr5Type
	Mdiosdoutr6  registerMdiosdoutr6Type
	Mdiosdoutr7  registerMdiosdoutr7Type
	Mdiosdoutr8  registerMdiosdoutr8Type
	Mdiosdoutr9  registerMdiosdoutr9Type
	Mdiosdoutr10 registerMdiosdoutr10Type
	Mdiosdoutr11 registerMdiosdoutr11Type
	Mdiosdoutr12 registerMdiosdoutr12Type
	Mdiosdoutr13 registerMdiosdoutr13Type
	Mdiosdoutr14 registerMdiosdoutr14Type
	Mdiosdoutr15 registerMdiosdoutr15Type
	Mdiosdoutr16 registerMdiosdoutr16Type
	Mdiosdoutr17 registerMdiosdoutr17Type
	Mdiosdoutr18 registerMdiosdoutr18Type
	Mdiosdoutr19 registerMdiosdoutr19Type
	Mdiosdoutr20 registerMdiosdoutr20Type
	Mdiosdoutr21 registerMdiosdoutr21Type
	Mdiosdoutr22 registerMdiosdoutr22Type
	Mdiosdoutr23 registerMdiosdoutr23Type
	Mdiosdoutr24 registerMdiosdoutr24Type
	Mdiosdoutr25 registerMdiosdoutr25Type
	Mdiosdoutr26 registerMdiosdoutr26Type
	Mdiosdoutr27 registerMdiosdoutr27Type
	Mdiosdoutr28 registerMdiosdoutr28Type
	Mdiosdoutr29 registerMdiosdoutr29Type
	Mdiosdoutr30 registerMdiosdoutr30Type
	Mdiosdoutr31 registerMdiosdoutr31Type
}

// registerMdioscrType MDIOS configuration register
type registerMdioscrType uint32

const (
	RegisterMdioscrFieldEnShift = 0
	RegisterMdioscrFieldEnMask  = 0x1
)

// GetEn Peripheral enable
func (r *registerMdioscrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrFieldEnMask) != 0
}

// SetEn Peripheral enable
func (r *registerMdioscrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdioscrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdioscrFieldEnMask)
	}
}

const (
	RegisterMdioscrFieldWrieShift = 1
	RegisterMdioscrFieldWrieMask  = 0x2
)

// GetWrie Register write interrupt enable
func (r *registerMdioscrType) GetWrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrFieldWrieMask) != 0
}

// SetWrie Register write interrupt enable
func (r *registerMdioscrType) SetWrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdioscrFieldWrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdioscrFieldWrieMask)
	}
}

const (
	RegisterMdioscrFieldRdieShift = 2
	RegisterMdioscrFieldRdieMask  = 0x4
)

// GetRdie Register Read Interrupt Enable
func (r *registerMdioscrType) GetRdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrFieldRdieMask) != 0
}

// SetRdie Register Read Interrupt Enable
func (r *registerMdioscrType) SetRdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdioscrFieldRdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdioscrFieldRdieMask)
	}
}

const (
	RegisterMdioscrFieldEieShift = 3
	RegisterMdioscrFieldEieMask  = 0x8
)

// GetEie Error interrupt enable
func (r *registerMdioscrType) GetEie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrFieldEieMask) != 0
}

// SetEie Error interrupt enable
func (r *registerMdioscrType) SetEie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdioscrFieldEieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdioscrFieldEieMask)
	}
}

const (
	RegisterMdioscrFieldDpcShift = 7
	RegisterMdioscrFieldDpcMask  = 0x80
)

// GetDpc Disable Preamble Check
func (r *registerMdioscrType) GetDpc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrFieldDpcMask) != 0
}

// SetDpc Disable Preamble Check
func (r *registerMdioscrType) SetDpc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdioscrFieldDpcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdioscrFieldDpcMask)
	}
}

const (
	RegisterMdioscrFieldPortaddressShift = 8
	RegisterMdioscrFieldPortaddressMask  = 0x1f00
)

// GetPortaddress Slaves's address
func (r *registerMdioscrType) GetPortaddress() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrFieldPortaddressMask) >> RegisterMdioscrFieldPortaddressShift)
}

// SetPortaddress Slaves's address
func (r *registerMdioscrType) SetPortaddress(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdioscrFieldPortaddressMask)|(uint32(value)<<RegisterMdioscrFieldPortaddressShift))
}

// registerMdioswrfrType MDIOS write flag register
type registerMdioswrfrType uint32

const (
	RegisterMdioswrfrFieldWrfShift = 0
	RegisterMdioswrfrFieldWrfMask  = 0xffffffff
)

// GetWrf Write flags for MDIO registers 0 to 31
func (r *registerMdioswrfrType) GetWrf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdioswrfrFieldWrfMask) >> RegisterMdioswrfrFieldWrfShift)
}

// SetWrf Write flags for MDIO registers 0 to 31
func (r *registerMdioswrfrType) SetWrf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdioswrfrFieldWrfMask)|(uint32(value)<<RegisterMdioswrfrFieldWrfShift))
}

// registerMdioscwrfrType MDIOS clear write flag register
type registerMdioscwrfrType uint32

const (
	RegisterMdioscwrfrFieldCwrfShift = 0
	RegisterMdioscwrfrFieldCwrfMask  = 0xffffffff
)

// GetCwrf Clear the write flag
func (r *registerMdioscwrfrType) GetCwrf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdioscwrfrFieldCwrfMask) >> RegisterMdioscwrfrFieldCwrfShift)
}

// SetCwrf Clear the write flag
func (r *registerMdioscwrfrType) SetCwrf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdioscwrfrFieldCwrfMask)|(uint32(value)<<RegisterMdioscwrfrFieldCwrfShift))
}

// registerMdiosrdfrType MDIOS read flag register
type registerMdiosrdfrType uint32

const (
	RegisterMdiosrdfrFieldRdfShift = 0
	RegisterMdiosrdfrFieldRdfMask  = 0xffffffff
)

// GetRdf Read flags for MDIO registers 0 to 31
func (r *registerMdiosrdfrType) GetRdf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosrdfrFieldRdfMask) >> RegisterMdiosrdfrFieldRdfShift)
}

// SetRdf Read flags for MDIO registers 0 to 31
func (r *registerMdiosrdfrType) SetRdf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosrdfrFieldRdfMask)|(uint32(value)<<RegisterMdiosrdfrFieldRdfShift))
}

// registerMdioscrdfrType MDIOS clear read flag register
type registerMdioscrdfrType uint32

const (
	RegisterMdioscrdfrFieldCrdfShift = 0
	RegisterMdioscrdfrFieldCrdfMask  = 0xffffffff
)

// GetCrdf Clear the read flag
func (r *registerMdioscrdfrType) GetCrdf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrdfrFieldCrdfMask) >> RegisterMdioscrdfrFieldCrdfShift)
}

// SetCrdf Clear the read flag
func (r *registerMdioscrdfrType) SetCrdf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdioscrdfrFieldCrdfMask)|(uint32(value)<<RegisterMdioscrdfrFieldCrdfShift))
}

// registerMdiossrType MDIOS status register
type registerMdiossrType uint32

const (
	RegisterMdiossrFieldPerfShift = 0
	RegisterMdiossrFieldPerfMask  = 0x1
)

// GetPerf Preamble error flag
func (r *registerMdiossrType) GetPerf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdiossrFieldPerfMask) != 0
}

// SetPerf Preamble error flag
func (r *registerMdiossrType) SetPerf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdiossrFieldPerfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdiossrFieldPerfMask)
	}
}

const (
	RegisterMdiossrFieldSerfShift = 1
	RegisterMdiossrFieldSerfMask  = 0x2
)

// GetSerf Start error flag
func (r *registerMdiossrType) GetSerf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdiossrFieldSerfMask) != 0
}

// SetSerf Start error flag
func (r *registerMdiossrType) SetSerf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdiossrFieldSerfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdiossrFieldSerfMask)
	}
}

const (
	RegisterMdiossrFieldTerfShift = 2
	RegisterMdiossrFieldTerfMask  = 0x4
)

// GetTerf Turnaround error flag
func (r *registerMdiossrType) GetTerf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdiossrFieldTerfMask) != 0
}

// SetTerf Turnaround error flag
func (r *registerMdiossrType) SetTerf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdiossrFieldTerfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdiossrFieldTerfMask)
	}
}

// registerMdiosclrfrType MDIOS clear flag register
type registerMdiosclrfrType uint32

const (
	RegisterMdiosclrfrFieldCperfShift = 0
	RegisterMdiosclrfrFieldCperfMask  = 0x1
)

// GetCperf Clear the preamble error flag
func (r *registerMdiosclrfrType) GetCperf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdiosclrfrFieldCperfMask) != 0
}

// SetCperf Clear the preamble error flag
func (r *registerMdiosclrfrType) SetCperf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdiosclrfrFieldCperfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdiosclrfrFieldCperfMask)
	}
}

const (
	RegisterMdiosclrfrFieldCserfShift = 1
	RegisterMdiosclrfrFieldCserfMask  = 0x2
)

// GetCserf Clear the start error flag
func (r *registerMdiosclrfrType) GetCserf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdiosclrfrFieldCserfMask) != 0
}

// SetCserf Clear the start error flag
func (r *registerMdiosclrfrType) SetCserf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdiosclrfrFieldCserfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdiosclrfrFieldCserfMask)
	}
}

const (
	RegisterMdiosclrfrFieldCterfShift = 2
	RegisterMdiosclrfrFieldCterfMask  = 0x4
)

// GetCterf Clear the turnaround error flag
func (r *registerMdiosclrfrType) GetCterf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdiosclrfrFieldCterfMask) != 0
}

// SetCterf Clear the turnaround error flag
func (r *registerMdiosclrfrType) SetCterf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdiosclrfrFieldCterfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdiosclrfrFieldCterfMask)
	}
}

// registerMdiosdinr0Type MDIOS input data register 0
type registerMdiosdinr0Type uint32

const (
	RegisterMdiosdinr0FieldDin0Shift = 0
	RegisterMdiosdinr0FieldDin0Mask  = 0xffff
)

// GetDin0 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr0Type) GetDin0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr0FieldDin0Mask) >> RegisterMdiosdinr0FieldDin0Shift)
}

// SetDin0 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr0Type) SetDin0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr0FieldDin0Mask)|(uint32(value)<<RegisterMdiosdinr0FieldDin0Shift))
}

// registerMdiosdinr1Type MDIOS input data register 1
type registerMdiosdinr1Type uint32

const (
	RegisterMdiosdinr1FieldDin1Shift = 0
	RegisterMdiosdinr1FieldDin1Mask  = 0xffff
)

// GetDin1 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr1Type) GetDin1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr1FieldDin1Mask) >> RegisterMdiosdinr1FieldDin1Shift)
}

// SetDin1 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr1Type) SetDin1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr1FieldDin1Mask)|(uint32(value)<<RegisterMdiosdinr1FieldDin1Shift))
}

// registerMdiosdinr2Type MDIOS input data register 2
type registerMdiosdinr2Type uint32

const (
	RegisterMdiosdinr2FieldDin2Shift = 0
	RegisterMdiosdinr2FieldDin2Mask  = 0xffff
)

// GetDin2 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr2Type) GetDin2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr2FieldDin2Mask) >> RegisterMdiosdinr2FieldDin2Shift)
}

// SetDin2 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr2Type) SetDin2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr2FieldDin2Mask)|(uint32(value)<<RegisterMdiosdinr2FieldDin2Shift))
}

// registerMdiosdinr3Type MDIOS input data register 3
type registerMdiosdinr3Type uint32

const (
	RegisterMdiosdinr3FieldDin3Shift = 0
	RegisterMdiosdinr3FieldDin3Mask  = 0xffff
)

// GetDin3 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr3Type) GetDin3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr3FieldDin3Mask) >> RegisterMdiosdinr3FieldDin3Shift)
}

// SetDin3 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr3Type) SetDin3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr3FieldDin3Mask)|(uint32(value)<<RegisterMdiosdinr3FieldDin3Shift))
}

// registerMdiosdinr4Type MDIOS input data register 4
type registerMdiosdinr4Type uint32

const (
	RegisterMdiosdinr4FieldDin4Shift = 0
	RegisterMdiosdinr4FieldDin4Mask  = 0xffff
)

// GetDin4 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr4Type) GetDin4() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr4FieldDin4Mask) >> RegisterMdiosdinr4FieldDin4Shift)
}

// SetDin4 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr4Type) SetDin4(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr4FieldDin4Mask)|(uint32(value)<<RegisterMdiosdinr4FieldDin4Shift))
}

// registerMdiosdinr5Type MDIOS input data register 5
type registerMdiosdinr5Type uint32

const (
	RegisterMdiosdinr5FieldDin5Shift = 0
	RegisterMdiosdinr5FieldDin5Mask  = 0xffff
)

// GetDin5 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr5Type) GetDin5() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr5FieldDin5Mask) >> RegisterMdiosdinr5FieldDin5Shift)
}

// SetDin5 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr5Type) SetDin5(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr5FieldDin5Mask)|(uint32(value)<<RegisterMdiosdinr5FieldDin5Shift))
}

// registerMdiosdinr6Type MDIOS input data register 6
type registerMdiosdinr6Type uint32

const (
	RegisterMdiosdinr6FieldDin6Shift = 0
	RegisterMdiosdinr6FieldDin6Mask  = 0xffff
)

// GetDin6 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr6Type) GetDin6() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr6FieldDin6Mask) >> RegisterMdiosdinr6FieldDin6Shift)
}

// SetDin6 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr6Type) SetDin6(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr6FieldDin6Mask)|(uint32(value)<<RegisterMdiosdinr6FieldDin6Shift))
}

// registerMdiosdinr7Type MDIOS input data register 7
type registerMdiosdinr7Type uint32

const (
	RegisterMdiosdinr7FieldDin7Shift = 0
	RegisterMdiosdinr7FieldDin7Mask  = 0xffff
)

// GetDin7 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr7Type) GetDin7() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr7FieldDin7Mask) >> RegisterMdiosdinr7FieldDin7Shift)
}

// SetDin7 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr7Type) SetDin7(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr7FieldDin7Mask)|(uint32(value)<<RegisterMdiosdinr7FieldDin7Shift))
}

// registerMdiosdinr8Type MDIOS input data register 8
type registerMdiosdinr8Type uint32

const (
	RegisterMdiosdinr8FieldDin8Shift = 0
	RegisterMdiosdinr8FieldDin8Mask  = 0xffff
)

// GetDin8 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr8Type) GetDin8() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr8FieldDin8Mask) >> RegisterMdiosdinr8FieldDin8Shift)
}

// SetDin8 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr8Type) SetDin8(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr8FieldDin8Mask)|(uint32(value)<<RegisterMdiosdinr8FieldDin8Shift))
}

// registerMdiosdinr9Type MDIOS input data register 9
type registerMdiosdinr9Type uint32

const (
	RegisterMdiosdinr9FieldDin9Shift = 0
	RegisterMdiosdinr9FieldDin9Mask  = 0xffff
)

// GetDin9 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr9Type) GetDin9() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr9FieldDin9Mask) >> RegisterMdiosdinr9FieldDin9Shift)
}

// SetDin9 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr9Type) SetDin9(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr9FieldDin9Mask)|(uint32(value)<<RegisterMdiosdinr9FieldDin9Shift))
}

// registerMdiosdinr10Type MDIOS input data register 10
type registerMdiosdinr10Type uint32

const (
	RegisterMdiosdinr10FieldDin10Shift = 0
	RegisterMdiosdinr10FieldDin10Mask  = 0xffff
)

// GetDin10 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr10Type) GetDin10() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr10FieldDin10Mask) >> RegisterMdiosdinr10FieldDin10Shift)
}

// SetDin10 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr10Type) SetDin10(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr10FieldDin10Mask)|(uint32(value)<<RegisterMdiosdinr10FieldDin10Shift))
}

// registerMdiosdinr11Type MDIOS input data register 11
type registerMdiosdinr11Type uint32

const (
	RegisterMdiosdinr11FieldDin11Shift = 0
	RegisterMdiosdinr11FieldDin11Mask  = 0xffff
)

// GetDin11 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr11Type) GetDin11() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr11FieldDin11Mask) >> RegisterMdiosdinr11FieldDin11Shift)
}

// SetDin11 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr11Type) SetDin11(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr11FieldDin11Mask)|(uint32(value)<<RegisterMdiosdinr11FieldDin11Shift))
}

// registerMdiosdinr12Type MDIOS input data register 12
type registerMdiosdinr12Type uint32

const (
	RegisterMdiosdinr12FieldDin12Shift = 0
	RegisterMdiosdinr12FieldDin12Mask  = 0xffff
)

// GetDin12 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr12Type) GetDin12() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr12FieldDin12Mask) >> RegisterMdiosdinr12FieldDin12Shift)
}

// SetDin12 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr12Type) SetDin12(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr12FieldDin12Mask)|(uint32(value)<<RegisterMdiosdinr12FieldDin12Shift))
}

// registerMdiosdinr13Type MDIOS input data register 13
type registerMdiosdinr13Type uint32

const (
	RegisterMdiosdinr13FieldDin13Shift = 0
	RegisterMdiosdinr13FieldDin13Mask  = 0xffff
)

// GetDin13 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr13Type) GetDin13() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr13FieldDin13Mask) >> RegisterMdiosdinr13FieldDin13Shift)
}

// SetDin13 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr13Type) SetDin13(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr13FieldDin13Mask)|(uint32(value)<<RegisterMdiosdinr13FieldDin13Shift))
}

// registerMdiosdinr14Type MDIOS input data register 14
type registerMdiosdinr14Type uint32

const (
	RegisterMdiosdinr14FieldDin14Shift = 0
	RegisterMdiosdinr14FieldDin14Mask  = 0xffff
)

// GetDin14 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr14Type) GetDin14() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr14FieldDin14Mask) >> RegisterMdiosdinr14FieldDin14Shift)
}

// SetDin14 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr14Type) SetDin14(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr14FieldDin14Mask)|(uint32(value)<<RegisterMdiosdinr14FieldDin14Shift))
}

// registerMdiosdinr15Type MDIOS input data register 15
type registerMdiosdinr15Type uint32

const (
	RegisterMdiosdinr15FieldDin15Shift = 0
	RegisterMdiosdinr15FieldDin15Mask  = 0xffff
)

// GetDin15 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr15Type) GetDin15() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr15FieldDin15Mask) >> RegisterMdiosdinr15FieldDin15Shift)
}

// SetDin15 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr15Type) SetDin15(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr15FieldDin15Mask)|(uint32(value)<<RegisterMdiosdinr15FieldDin15Shift))
}

// registerMdiosdinr16Type MDIOS input data register 16
type registerMdiosdinr16Type uint32

const (
	RegisterMdiosdinr16FieldDin16Shift = 0
	RegisterMdiosdinr16FieldDin16Mask  = 0xffff
)

// GetDin16 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr16Type) GetDin16() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr16FieldDin16Mask) >> RegisterMdiosdinr16FieldDin16Shift)
}

// SetDin16 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr16Type) SetDin16(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr16FieldDin16Mask)|(uint32(value)<<RegisterMdiosdinr16FieldDin16Shift))
}

// registerMdiosdinr17Type MDIOS input data register 17
type registerMdiosdinr17Type uint32

const (
	RegisterMdiosdinr17FieldDin17Shift = 0
	RegisterMdiosdinr17FieldDin17Mask  = 0xffff
)

// GetDin17 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr17Type) GetDin17() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr17FieldDin17Mask) >> RegisterMdiosdinr17FieldDin17Shift)
}

// SetDin17 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr17Type) SetDin17(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr17FieldDin17Mask)|(uint32(value)<<RegisterMdiosdinr17FieldDin17Shift))
}

// registerMdiosdinr18Type MDIOS input data register 18
type registerMdiosdinr18Type uint32

const (
	RegisterMdiosdinr18FieldDin18Shift = 0
	RegisterMdiosdinr18FieldDin18Mask  = 0xffff
)

// GetDin18 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr18Type) GetDin18() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr18FieldDin18Mask) >> RegisterMdiosdinr18FieldDin18Shift)
}

// SetDin18 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr18Type) SetDin18(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr18FieldDin18Mask)|(uint32(value)<<RegisterMdiosdinr18FieldDin18Shift))
}

// registerMdiosdinr19Type MDIOS input data register 19
type registerMdiosdinr19Type uint32

const (
	RegisterMdiosdinr19FieldDin19Shift = 0
	RegisterMdiosdinr19FieldDin19Mask  = 0xffff
)

// GetDin19 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr19Type) GetDin19() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr19FieldDin19Mask) >> RegisterMdiosdinr19FieldDin19Shift)
}

// SetDin19 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr19Type) SetDin19(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr19FieldDin19Mask)|(uint32(value)<<RegisterMdiosdinr19FieldDin19Shift))
}

// registerMdiosdinr20Type MDIOS input data register 20
type registerMdiosdinr20Type uint32

const (
	RegisterMdiosdinr20FieldDin20Shift = 0
	RegisterMdiosdinr20FieldDin20Mask  = 0xffff
)

// GetDin20 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr20Type) GetDin20() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr20FieldDin20Mask) >> RegisterMdiosdinr20FieldDin20Shift)
}

// SetDin20 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr20Type) SetDin20(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr20FieldDin20Mask)|(uint32(value)<<RegisterMdiosdinr20FieldDin20Shift))
}

// registerMdiosdinr21Type MDIOS input data register 21
type registerMdiosdinr21Type uint32

const (
	RegisterMdiosdinr21FieldDin21Shift = 0
	RegisterMdiosdinr21FieldDin21Mask  = 0xffff
)

// GetDin21 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr21Type) GetDin21() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr21FieldDin21Mask) >> RegisterMdiosdinr21FieldDin21Shift)
}

// SetDin21 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr21Type) SetDin21(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr21FieldDin21Mask)|(uint32(value)<<RegisterMdiosdinr21FieldDin21Shift))
}

// registerMdiosdinr22Type MDIOS input data register 22
type registerMdiosdinr22Type uint32

const (
	RegisterMdiosdinr22FieldDin22Shift = 0
	RegisterMdiosdinr22FieldDin22Mask  = 0xffff
)

// GetDin22 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr22Type) GetDin22() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr22FieldDin22Mask) >> RegisterMdiosdinr22FieldDin22Shift)
}

// SetDin22 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr22Type) SetDin22(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr22FieldDin22Mask)|(uint32(value)<<RegisterMdiosdinr22FieldDin22Shift))
}

// registerMdiosdinr23Type MDIOS input data register 23
type registerMdiosdinr23Type uint32

const (
	RegisterMdiosdinr23FieldDin23Shift = 0
	RegisterMdiosdinr23FieldDin23Mask  = 0xffff
)

// GetDin23 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr23Type) GetDin23() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr23FieldDin23Mask) >> RegisterMdiosdinr23FieldDin23Shift)
}

// SetDin23 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr23Type) SetDin23(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr23FieldDin23Mask)|(uint32(value)<<RegisterMdiosdinr23FieldDin23Shift))
}

// registerMdiosdinr24Type MDIOS input data register 24
type registerMdiosdinr24Type uint32

const (
	RegisterMdiosdinr24FieldDin24Shift = 0
	RegisterMdiosdinr24FieldDin24Mask  = 0xffff
)

// GetDin24 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr24Type) GetDin24() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr24FieldDin24Mask) >> RegisterMdiosdinr24FieldDin24Shift)
}

// SetDin24 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr24Type) SetDin24(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr24FieldDin24Mask)|(uint32(value)<<RegisterMdiosdinr24FieldDin24Shift))
}

// registerMdiosdinr25Type MDIOS input data register 25
type registerMdiosdinr25Type uint32

const (
	RegisterMdiosdinr25FieldDin25Shift = 0
	RegisterMdiosdinr25FieldDin25Mask  = 0xffff
)

// GetDin25 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr25Type) GetDin25() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr25FieldDin25Mask) >> RegisterMdiosdinr25FieldDin25Shift)
}

// SetDin25 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr25Type) SetDin25(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr25FieldDin25Mask)|(uint32(value)<<RegisterMdiosdinr25FieldDin25Shift))
}

// registerMdiosdinr26Type MDIOS input data register 26
type registerMdiosdinr26Type uint32

const (
	RegisterMdiosdinr26FieldDin26Shift = 0
	RegisterMdiosdinr26FieldDin26Mask  = 0xffff
)

// GetDin26 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr26Type) GetDin26() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr26FieldDin26Mask) >> RegisterMdiosdinr26FieldDin26Shift)
}

// SetDin26 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr26Type) SetDin26(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr26FieldDin26Mask)|(uint32(value)<<RegisterMdiosdinr26FieldDin26Shift))
}

// registerMdiosdinr27Type MDIOS input data register 27
type registerMdiosdinr27Type uint32

const (
	RegisterMdiosdinr27FieldDin27Shift = 0
	RegisterMdiosdinr27FieldDin27Mask  = 0xffff
)

// GetDin27 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr27Type) GetDin27() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr27FieldDin27Mask) >> RegisterMdiosdinr27FieldDin27Shift)
}

// SetDin27 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr27Type) SetDin27(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr27FieldDin27Mask)|(uint32(value)<<RegisterMdiosdinr27FieldDin27Shift))
}

// registerMdiosdinr28Type MDIOS input data register 28
type registerMdiosdinr28Type uint32

const (
	RegisterMdiosdinr28FieldDin28Shift = 0
	RegisterMdiosdinr28FieldDin28Mask  = 0xffff
)

// GetDin28 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr28Type) GetDin28() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr28FieldDin28Mask) >> RegisterMdiosdinr28FieldDin28Shift)
}

// SetDin28 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr28Type) SetDin28(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr28FieldDin28Mask)|(uint32(value)<<RegisterMdiosdinr28FieldDin28Shift))
}

// registerMdiosdinr29Type MDIOS input data register 29
type registerMdiosdinr29Type uint32

const (
	RegisterMdiosdinr29FieldDin29Shift = 0
	RegisterMdiosdinr29FieldDin29Mask  = 0xffff
)

// GetDin29 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr29Type) GetDin29() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr29FieldDin29Mask) >> RegisterMdiosdinr29FieldDin29Shift)
}

// SetDin29 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr29Type) SetDin29(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr29FieldDin29Mask)|(uint32(value)<<RegisterMdiosdinr29FieldDin29Shift))
}

// registerMdiosdinr30Type MDIOS input data register 30
type registerMdiosdinr30Type uint32

const (
	RegisterMdiosdinr30FieldDin30Shift = 0
	RegisterMdiosdinr30FieldDin30Mask  = 0xffff
)

// GetDin30 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr30Type) GetDin30() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr30FieldDin30Mask) >> RegisterMdiosdinr30FieldDin30Shift)
}

// SetDin30 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr30Type) SetDin30(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr30FieldDin30Mask)|(uint32(value)<<RegisterMdiosdinr30FieldDin30Shift))
}

// registerMdiosdinr31Type MDIOS input data register 31
type registerMdiosdinr31Type uint32

const (
	RegisterMdiosdinr31FieldDin31Shift = 0
	RegisterMdiosdinr31FieldDin31Mask  = 0xffff
)

// GetDin31 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr31Type) GetDin31() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr31FieldDin31Mask) >> RegisterMdiosdinr31FieldDin31Shift)
}

// SetDin31 Input data received from MDIO Master during write frames
func (r *registerMdiosdinr31Type) SetDin31(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr31FieldDin31Mask)|(uint32(value)<<RegisterMdiosdinr31FieldDin31Shift))
}

// registerMdiosdoutr0Type MDIOS output data register 0
type registerMdiosdoutr0Type uint32

const (
	RegisterMdiosdoutr0FieldDout0Shift = 0
	RegisterMdiosdoutr0FieldDout0Mask  = 0xffff
)

// GetDout0 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr0Type) GetDout0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr0FieldDout0Mask) >> RegisterMdiosdoutr0FieldDout0Shift)
}

// SetDout0 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr0Type) SetDout0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr0FieldDout0Mask)|(uint32(value)<<RegisterMdiosdoutr0FieldDout0Shift))
}

// registerMdiosdoutr1Type MDIOS output data register 1
type registerMdiosdoutr1Type uint32

const (
	RegisterMdiosdoutr1FieldDout1Shift = 0
	RegisterMdiosdoutr1FieldDout1Mask  = 0xffff
)

// GetDout1 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr1Type) GetDout1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr1FieldDout1Mask) >> RegisterMdiosdoutr1FieldDout1Shift)
}

// SetDout1 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr1Type) SetDout1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr1FieldDout1Mask)|(uint32(value)<<RegisterMdiosdoutr1FieldDout1Shift))
}

// registerMdiosdoutr2Type MDIOS output data register 2
type registerMdiosdoutr2Type uint32

const (
	RegisterMdiosdoutr2FieldDout2Shift = 0
	RegisterMdiosdoutr2FieldDout2Mask  = 0xffff
)

// GetDout2 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr2Type) GetDout2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr2FieldDout2Mask) >> RegisterMdiosdoutr2FieldDout2Shift)
}

// SetDout2 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr2Type) SetDout2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr2FieldDout2Mask)|(uint32(value)<<RegisterMdiosdoutr2FieldDout2Shift))
}

// registerMdiosdoutr3Type MDIOS output data register 3
type registerMdiosdoutr3Type uint32

const (
	RegisterMdiosdoutr3FieldDout3Shift = 0
	RegisterMdiosdoutr3FieldDout3Mask  = 0xffff
)

// GetDout3 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr3Type) GetDout3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr3FieldDout3Mask) >> RegisterMdiosdoutr3FieldDout3Shift)
}

// SetDout3 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr3Type) SetDout3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr3FieldDout3Mask)|(uint32(value)<<RegisterMdiosdoutr3FieldDout3Shift))
}

// registerMdiosdoutr4Type MDIOS output data register 4
type registerMdiosdoutr4Type uint32

const (
	RegisterMdiosdoutr4FieldDout4Shift = 0
	RegisterMdiosdoutr4FieldDout4Mask  = 0xffff
)

// GetDout4 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr4Type) GetDout4() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr4FieldDout4Mask) >> RegisterMdiosdoutr4FieldDout4Shift)
}

// SetDout4 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr4Type) SetDout4(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr4FieldDout4Mask)|(uint32(value)<<RegisterMdiosdoutr4FieldDout4Shift))
}

// registerMdiosdoutr5Type MDIOS output data register 5
type registerMdiosdoutr5Type uint32

const (
	RegisterMdiosdoutr5FieldDout5Shift = 0
	RegisterMdiosdoutr5FieldDout5Mask  = 0xffff
)

// GetDout5 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr5Type) GetDout5() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr5FieldDout5Mask) >> RegisterMdiosdoutr5FieldDout5Shift)
}

// SetDout5 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr5Type) SetDout5(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr5FieldDout5Mask)|(uint32(value)<<RegisterMdiosdoutr5FieldDout5Shift))
}

// registerMdiosdoutr6Type MDIOS output data register 6
type registerMdiosdoutr6Type uint32

const (
	RegisterMdiosdoutr6FieldDout6Shift = 0
	RegisterMdiosdoutr6FieldDout6Mask  = 0xffff
)

// GetDout6 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr6Type) GetDout6() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr6FieldDout6Mask) >> RegisterMdiosdoutr6FieldDout6Shift)
}

// SetDout6 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr6Type) SetDout6(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr6FieldDout6Mask)|(uint32(value)<<RegisterMdiosdoutr6FieldDout6Shift))
}

// registerMdiosdoutr7Type MDIOS output data register 7
type registerMdiosdoutr7Type uint32

const (
	RegisterMdiosdoutr7FieldDout7Shift = 0
	RegisterMdiosdoutr7FieldDout7Mask  = 0xffff
)

// GetDout7 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr7Type) GetDout7() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr7FieldDout7Mask) >> RegisterMdiosdoutr7FieldDout7Shift)
}

// SetDout7 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr7Type) SetDout7(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr7FieldDout7Mask)|(uint32(value)<<RegisterMdiosdoutr7FieldDout7Shift))
}

// registerMdiosdoutr8Type MDIOS output data register 8
type registerMdiosdoutr8Type uint32

const (
	RegisterMdiosdoutr8FieldDout8Shift = 0
	RegisterMdiosdoutr8FieldDout8Mask  = 0xffff
)

// GetDout8 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr8Type) GetDout8() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr8FieldDout8Mask) >> RegisterMdiosdoutr8FieldDout8Shift)
}

// SetDout8 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr8Type) SetDout8(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr8FieldDout8Mask)|(uint32(value)<<RegisterMdiosdoutr8FieldDout8Shift))
}

// registerMdiosdoutr9Type MDIOS output data register 9
type registerMdiosdoutr9Type uint32

const (
	RegisterMdiosdoutr9FieldDout9Shift = 0
	RegisterMdiosdoutr9FieldDout9Mask  = 0xffff
)

// GetDout9 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr9Type) GetDout9() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr9FieldDout9Mask) >> RegisterMdiosdoutr9FieldDout9Shift)
}

// SetDout9 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr9Type) SetDout9(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr9FieldDout9Mask)|(uint32(value)<<RegisterMdiosdoutr9FieldDout9Shift))
}

// registerMdiosdoutr10Type MDIOS output data register 10
type registerMdiosdoutr10Type uint32

const (
	RegisterMdiosdoutr10FieldDout10Shift = 0
	RegisterMdiosdoutr10FieldDout10Mask  = 0xffff
)

// GetDout10 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr10Type) GetDout10() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr10FieldDout10Mask) >> RegisterMdiosdoutr10FieldDout10Shift)
}

// SetDout10 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr10Type) SetDout10(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr10FieldDout10Mask)|(uint32(value)<<RegisterMdiosdoutr10FieldDout10Shift))
}

// registerMdiosdoutr11Type MDIOS output data register 11
type registerMdiosdoutr11Type uint32

const (
	RegisterMdiosdoutr11FieldDout11Shift = 0
	RegisterMdiosdoutr11FieldDout11Mask  = 0xffff
)

// GetDout11 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr11Type) GetDout11() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr11FieldDout11Mask) >> RegisterMdiosdoutr11FieldDout11Shift)
}

// SetDout11 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr11Type) SetDout11(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr11FieldDout11Mask)|(uint32(value)<<RegisterMdiosdoutr11FieldDout11Shift))
}

// registerMdiosdoutr12Type MDIOS output data register 12
type registerMdiosdoutr12Type uint32

const (
	RegisterMdiosdoutr12FieldDout12Shift = 0
	RegisterMdiosdoutr12FieldDout12Mask  = 0xffff
)

// GetDout12 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr12Type) GetDout12() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr12FieldDout12Mask) >> RegisterMdiosdoutr12FieldDout12Shift)
}

// SetDout12 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr12Type) SetDout12(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr12FieldDout12Mask)|(uint32(value)<<RegisterMdiosdoutr12FieldDout12Shift))
}

// registerMdiosdoutr13Type MDIOS output data register 13
type registerMdiosdoutr13Type uint32

const (
	RegisterMdiosdoutr13FieldDout13Shift = 0
	RegisterMdiosdoutr13FieldDout13Mask  = 0xffff
)

// GetDout13 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr13Type) GetDout13() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr13FieldDout13Mask) >> RegisterMdiosdoutr13FieldDout13Shift)
}

// SetDout13 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr13Type) SetDout13(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr13FieldDout13Mask)|(uint32(value)<<RegisterMdiosdoutr13FieldDout13Shift))
}

// registerMdiosdoutr14Type MDIOS output data register 14
type registerMdiosdoutr14Type uint32

const (
	RegisterMdiosdoutr14FieldDout14Shift = 0
	RegisterMdiosdoutr14FieldDout14Mask  = 0xffff
)

// GetDout14 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr14Type) GetDout14() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr14FieldDout14Mask) >> RegisterMdiosdoutr14FieldDout14Shift)
}

// SetDout14 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr14Type) SetDout14(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr14FieldDout14Mask)|(uint32(value)<<RegisterMdiosdoutr14FieldDout14Shift))
}

// registerMdiosdoutr15Type MDIOS output data register 15
type registerMdiosdoutr15Type uint32

const (
	RegisterMdiosdoutr15FieldDout15Shift = 0
	RegisterMdiosdoutr15FieldDout15Mask  = 0xffff
)

// GetDout15 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr15Type) GetDout15() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr15FieldDout15Mask) >> RegisterMdiosdoutr15FieldDout15Shift)
}

// SetDout15 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr15Type) SetDout15(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr15FieldDout15Mask)|(uint32(value)<<RegisterMdiosdoutr15FieldDout15Shift))
}

// registerMdiosdoutr16Type MDIOS output data register 16
type registerMdiosdoutr16Type uint32

const (
	RegisterMdiosdoutr16FieldDout16Shift = 0
	RegisterMdiosdoutr16FieldDout16Mask  = 0xffff
)

// GetDout16 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr16Type) GetDout16() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr16FieldDout16Mask) >> RegisterMdiosdoutr16FieldDout16Shift)
}

// SetDout16 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr16Type) SetDout16(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr16FieldDout16Mask)|(uint32(value)<<RegisterMdiosdoutr16FieldDout16Shift))
}

// registerMdiosdoutr17Type MDIOS output data register 17
type registerMdiosdoutr17Type uint32

const (
	RegisterMdiosdoutr17FieldDout17Shift = 0
	RegisterMdiosdoutr17FieldDout17Mask  = 0xffff
)

// GetDout17 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr17Type) GetDout17() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr17FieldDout17Mask) >> RegisterMdiosdoutr17FieldDout17Shift)
}

// SetDout17 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr17Type) SetDout17(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr17FieldDout17Mask)|(uint32(value)<<RegisterMdiosdoutr17FieldDout17Shift))
}

// registerMdiosdoutr18Type MDIOS output data register 18
type registerMdiosdoutr18Type uint32

const (
	RegisterMdiosdoutr18FieldDout18Shift = 0
	RegisterMdiosdoutr18FieldDout18Mask  = 0xffff
)

// GetDout18 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr18Type) GetDout18() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr18FieldDout18Mask) >> RegisterMdiosdoutr18FieldDout18Shift)
}

// SetDout18 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr18Type) SetDout18(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr18FieldDout18Mask)|(uint32(value)<<RegisterMdiosdoutr18FieldDout18Shift))
}

// registerMdiosdoutr19Type MDIOS output data register 19
type registerMdiosdoutr19Type uint32

const (
	RegisterMdiosdoutr19FieldDout19Shift = 0
	RegisterMdiosdoutr19FieldDout19Mask  = 0xffff
)

// GetDout19 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr19Type) GetDout19() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr19FieldDout19Mask) >> RegisterMdiosdoutr19FieldDout19Shift)
}

// SetDout19 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr19Type) SetDout19(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr19FieldDout19Mask)|(uint32(value)<<RegisterMdiosdoutr19FieldDout19Shift))
}

// registerMdiosdoutr20Type MDIOS output data register 20
type registerMdiosdoutr20Type uint32

const (
	RegisterMdiosdoutr20FieldDout20Shift = 0
	RegisterMdiosdoutr20FieldDout20Mask  = 0xffff
)

// GetDout20 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr20Type) GetDout20() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr20FieldDout20Mask) >> RegisterMdiosdoutr20FieldDout20Shift)
}

// SetDout20 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr20Type) SetDout20(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr20FieldDout20Mask)|(uint32(value)<<RegisterMdiosdoutr20FieldDout20Shift))
}

// registerMdiosdoutr21Type MDIOS output data register 21
type registerMdiosdoutr21Type uint32

const (
	RegisterMdiosdoutr21FieldDout21Shift = 0
	RegisterMdiosdoutr21FieldDout21Mask  = 0xffff
)

// GetDout21 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr21Type) GetDout21() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr21FieldDout21Mask) >> RegisterMdiosdoutr21FieldDout21Shift)
}

// SetDout21 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr21Type) SetDout21(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr21FieldDout21Mask)|(uint32(value)<<RegisterMdiosdoutr21FieldDout21Shift))
}

// registerMdiosdoutr22Type MDIOS output data register 22
type registerMdiosdoutr22Type uint32

const (
	RegisterMdiosdoutr22FieldDout22Shift = 0
	RegisterMdiosdoutr22FieldDout22Mask  = 0xffff
)

// GetDout22 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr22Type) GetDout22() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr22FieldDout22Mask) >> RegisterMdiosdoutr22FieldDout22Shift)
}

// SetDout22 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr22Type) SetDout22(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr22FieldDout22Mask)|(uint32(value)<<RegisterMdiosdoutr22FieldDout22Shift))
}

// registerMdiosdoutr23Type MDIOS output data register 23
type registerMdiosdoutr23Type uint32

const (
	RegisterMdiosdoutr23FieldDout23Shift = 0
	RegisterMdiosdoutr23FieldDout23Mask  = 0xffff
)

// GetDout23 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr23Type) GetDout23() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr23FieldDout23Mask) >> RegisterMdiosdoutr23FieldDout23Shift)
}

// SetDout23 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr23Type) SetDout23(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr23FieldDout23Mask)|(uint32(value)<<RegisterMdiosdoutr23FieldDout23Shift))
}

// registerMdiosdoutr24Type MDIOS output data register 24
type registerMdiosdoutr24Type uint32

const (
	RegisterMdiosdoutr24FieldDout24Shift = 0
	RegisterMdiosdoutr24FieldDout24Mask  = 0xffff
)

// GetDout24 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr24Type) GetDout24() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr24FieldDout24Mask) >> RegisterMdiosdoutr24FieldDout24Shift)
}

// SetDout24 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr24Type) SetDout24(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr24FieldDout24Mask)|(uint32(value)<<RegisterMdiosdoutr24FieldDout24Shift))
}

// registerMdiosdoutr25Type MDIOS output data register 25
type registerMdiosdoutr25Type uint32

const (
	RegisterMdiosdoutr25FieldDout25Shift = 0
	RegisterMdiosdoutr25FieldDout25Mask  = 0xffff
)

// GetDout25 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr25Type) GetDout25() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr25FieldDout25Mask) >> RegisterMdiosdoutr25FieldDout25Shift)
}

// SetDout25 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr25Type) SetDout25(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr25FieldDout25Mask)|(uint32(value)<<RegisterMdiosdoutr25FieldDout25Shift))
}

// registerMdiosdoutr26Type MDIOS output data register 26
type registerMdiosdoutr26Type uint32

const (
	RegisterMdiosdoutr26FieldDout26Shift = 0
	RegisterMdiosdoutr26FieldDout26Mask  = 0xffff
)

// GetDout26 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr26Type) GetDout26() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr26FieldDout26Mask) >> RegisterMdiosdoutr26FieldDout26Shift)
}

// SetDout26 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr26Type) SetDout26(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr26FieldDout26Mask)|(uint32(value)<<RegisterMdiosdoutr26FieldDout26Shift))
}

// registerMdiosdoutr27Type MDIOS output data register 27
type registerMdiosdoutr27Type uint32

const (
	RegisterMdiosdoutr27FieldDout27Shift = 0
	RegisterMdiosdoutr27FieldDout27Mask  = 0xffff
)

// GetDout27 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr27Type) GetDout27() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr27FieldDout27Mask) >> RegisterMdiosdoutr27FieldDout27Shift)
}

// SetDout27 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr27Type) SetDout27(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr27FieldDout27Mask)|(uint32(value)<<RegisterMdiosdoutr27FieldDout27Shift))
}

// registerMdiosdoutr28Type MDIOS output data register 28
type registerMdiosdoutr28Type uint32

const (
	RegisterMdiosdoutr28FieldDout28Shift = 0
	RegisterMdiosdoutr28FieldDout28Mask  = 0xffff
)

// GetDout28 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr28Type) GetDout28() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr28FieldDout28Mask) >> RegisterMdiosdoutr28FieldDout28Shift)
}

// SetDout28 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr28Type) SetDout28(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr28FieldDout28Mask)|(uint32(value)<<RegisterMdiosdoutr28FieldDout28Shift))
}

// registerMdiosdoutr29Type MDIOS output data register 29
type registerMdiosdoutr29Type uint32

const (
	RegisterMdiosdoutr29FieldDout29Shift = 0
	RegisterMdiosdoutr29FieldDout29Mask  = 0xffff
)

// GetDout29 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr29Type) GetDout29() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr29FieldDout29Mask) >> RegisterMdiosdoutr29FieldDout29Shift)
}

// SetDout29 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr29Type) SetDout29(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr29FieldDout29Mask)|(uint32(value)<<RegisterMdiosdoutr29FieldDout29Shift))
}

// registerMdiosdoutr30Type MDIOS output data register 30
type registerMdiosdoutr30Type uint32

const (
	RegisterMdiosdoutr30FieldDout30Shift = 0
	RegisterMdiosdoutr30FieldDout30Mask  = 0xffff
)

// GetDout30 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr30Type) GetDout30() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr30FieldDout30Mask) >> RegisterMdiosdoutr30FieldDout30Shift)
}

// SetDout30 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr30Type) SetDout30(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr30FieldDout30Mask)|(uint32(value)<<RegisterMdiosdoutr30FieldDout30Shift))
}

// registerMdiosdoutr31Type MDIOS output data register 31
type registerMdiosdoutr31Type uint32

const (
	RegisterMdiosdoutr31FieldDout31Shift = 0
	RegisterMdiosdoutr31FieldDout31Mask  = 0xffff
)

// GetDout31 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr31Type) GetDout31() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr31FieldDout31Mask) >> RegisterMdiosdoutr31FieldDout31Shift)
}

// SetDout31 Output data sent to MDIO Master during read frames
func (r *registerMdiosdoutr31Type) SetDout31(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr31FieldDout31Mask)|(uint32(value)<<RegisterMdiosdoutr31FieldDout31Shift))
}
