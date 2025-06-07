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
	Mdioscr      RegisterMdioscrType
	Mdioswrfr    RegisterMdioswrfrType
	Mdioscwrfr   RegisterMdioscwrfrType
	Mdiosrdfr    RegisterMdiosrdfrType
	Mdioscrdfr   RegisterMdioscrdfrType
	Mdiossr      RegisterMdiossrType
	Mdiosclrfr   RegisterMdiosclrfrType
	Mdiosdinr0   RegisterMdiosdinr0Type
	Mdiosdinr1   RegisterMdiosdinr1Type
	Mdiosdinr2   RegisterMdiosdinr2Type
	Mdiosdinr3   RegisterMdiosdinr3Type
	Mdiosdinr4   RegisterMdiosdinr4Type
	Mdiosdinr5   RegisterMdiosdinr5Type
	Mdiosdinr6   RegisterMdiosdinr6Type
	Mdiosdinr7   RegisterMdiosdinr7Type
	Mdiosdinr8   RegisterMdiosdinr8Type
	Mdiosdinr9   RegisterMdiosdinr9Type
	Mdiosdinr10  RegisterMdiosdinr10Type
	Mdiosdinr11  RegisterMdiosdinr11Type
	Mdiosdinr12  RegisterMdiosdinr12Type
	Mdiosdinr13  RegisterMdiosdinr13Type
	Mdiosdinr14  RegisterMdiosdinr14Type
	Mdiosdinr15  RegisterMdiosdinr15Type
	Mdiosdinr16  RegisterMdiosdinr16Type
	Mdiosdinr17  RegisterMdiosdinr17Type
	Mdiosdinr18  RegisterMdiosdinr18Type
	Mdiosdinr19  RegisterMdiosdinr19Type
	Mdiosdinr20  RegisterMdiosdinr20Type
	Mdiosdinr21  RegisterMdiosdinr21Type
	Mdiosdinr22  RegisterMdiosdinr22Type
	Mdiosdinr23  RegisterMdiosdinr23Type
	Mdiosdinr24  RegisterMdiosdinr24Type
	Mdiosdinr25  RegisterMdiosdinr25Type
	Mdiosdinr26  RegisterMdiosdinr26Type
	Mdiosdinr27  RegisterMdiosdinr27Type
	Mdiosdinr28  RegisterMdiosdinr28Type
	Mdiosdinr29  RegisterMdiosdinr29Type
	Mdiosdinr30  RegisterMdiosdinr30Type
	Mdiosdinr31  RegisterMdiosdinr31Type
	Mdiosdoutr0  RegisterMdiosdoutr0Type
	Mdiosdoutr1  RegisterMdiosdoutr1Type
	Mdiosdoutr2  RegisterMdiosdoutr2Type
	Mdiosdoutr3  RegisterMdiosdoutr3Type
	Mdiosdoutr4  RegisterMdiosdoutr4Type
	Mdiosdoutr5  RegisterMdiosdoutr5Type
	Mdiosdoutr6  RegisterMdiosdoutr6Type
	Mdiosdoutr7  RegisterMdiosdoutr7Type
	Mdiosdoutr8  RegisterMdiosdoutr8Type
	Mdiosdoutr9  RegisterMdiosdoutr9Type
	Mdiosdoutr10 RegisterMdiosdoutr10Type
	Mdiosdoutr11 RegisterMdiosdoutr11Type
	Mdiosdoutr12 RegisterMdiosdoutr12Type
	Mdiosdoutr13 RegisterMdiosdoutr13Type
	Mdiosdoutr14 RegisterMdiosdoutr14Type
	Mdiosdoutr15 RegisterMdiosdoutr15Type
	Mdiosdoutr16 RegisterMdiosdoutr16Type
	Mdiosdoutr17 RegisterMdiosdoutr17Type
	Mdiosdoutr18 RegisterMdiosdoutr18Type
	Mdiosdoutr19 RegisterMdiosdoutr19Type
	Mdiosdoutr20 RegisterMdiosdoutr20Type
	Mdiosdoutr21 RegisterMdiosdoutr21Type
	Mdiosdoutr22 RegisterMdiosdoutr22Type
	Mdiosdoutr23 RegisterMdiosdoutr23Type
	Mdiosdoutr24 RegisterMdiosdoutr24Type
	Mdiosdoutr25 RegisterMdiosdoutr25Type
	Mdiosdoutr26 RegisterMdiosdoutr26Type
	Mdiosdoutr27 RegisterMdiosdoutr27Type
	Mdiosdoutr28 RegisterMdiosdoutr28Type
	Mdiosdoutr29 RegisterMdiosdoutr29Type
	Mdiosdoutr30 RegisterMdiosdoutr30Type
	Mdiosdoutr31 RegisterMdiosdoutr31Type
}

// RegisterMdioscrType MDIOS configuration register
type RegisterMdioscrType uint32

func (r *RegisterMdioscrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdioscrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdioscrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdioscrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdioscrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdioscrFieldEnShift = 0
	RegisterMdioscrFieldEnMask  = 0x1
)

// GetEn Peripheral enable
func (r *RegisterMdioscrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrFieldEnMask) != 0
}

// SetEn Peripheral enable
func (r *RegisterMdioscrType) SetEn(value bool) {
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
func (r *RegisterMdioscrType) GetWrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrFieldWrieMask) != 0
}

// SetWrie Register write interrupt enable
func (r *RegisterMdioscrType) SetWrie(value bool) {
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
func (r *RegisterMdioscrType) GetRdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrFieldRdieMask) != 0
}

// SetRdie Register Read Interrupt Enable
func (r *RegisterMdioscrType) SetRdie(value bool) {
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
func (r *RegisterMdioscrType) GetEie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrFieldEieMask) != 0
}

// SetEie Error interrupt enable
func (r *RegisterMdioscrType) SetEie(value bool) {
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
func (r *RegisterMdioscrType) GetDpc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrFieldDpcMask) != 0
}

// SetDpc Disable Preamble Check
func (r *RegisterMdioscrType) SetDpc(value bool) {
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
func (r *RegisterMdioscrType) GetPortaddress() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrFieldPortaddressMask) >> RegisterMdioscrFieldPortaddressShift)
}

// SetPortaddress Slaves's address
func (r *RegisterMdioscrType) SetPortaddress(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdioscrFieldPortaddressMask)|(uint32(value)<<RegisterMdioscrFieldPortaddressShift))
}

// RegisterMdioswrfrType MDIOS write flag register
type RegisterMdioswrfrType uint32

func (r *RegisterMdioswrfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdioswrfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdioswrfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdioswrfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdioswrfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdioswrfrFieldWrfShift = 0
	RegisterMdioswrfrFieldWrfMask  = 0xffffffff
)

// GetWrf Write flags for MDIO registers 0 to 31
func (r *RegisterMdioswrfrType) GetWrf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdioswrfrFieldWrfMask) >> RegisterMdioswrfrFieldWrfShift)
}

// SetWrf Write flags for MDIO registers 0 to 31
func (r *RegisterMdioswrfrType) SetWrf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdioswrfrFieldWrfMask)|(uint32(value)<<RegisterMdioswrfrFieldWrfShift))
}

// RegisterMdioscwrfrType MDIOS clear write flag register
type RegisterMdioscwrfrType uint32

func (r *RegisterMdioscwrfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdioscwrfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdioscwrfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdioscwrfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdioscwrfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdioscwrfrFieldCwrfShift = 0
	RegisterMdioscwrfrFieldCwrfMask  = 0xffffffff
)

// GetCwrf Clear the write flag
func (r *RegisterMdioscwrfrType) GetCwrf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdioscwrfrFieldCwrfMask) >> RegisterMdioscwrfrFieldCwrfShift)
}

// SetCwrf Clear the write flag
func (r *RegisterMdioscwrfrType) SetCwrf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdioscwrfrFieldCwrfMask)|(uint32(value)<<RegisterMdioscwrfrFieldCwrfShift))
}

// RegisterMdiosrdfrType MDIOS read flag register
type RegisterMdiosrdfrType uint32

func (r *RegisterMdiosrdfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosrdfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosrdfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosrdfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosrdfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosrdfrFieldRdfShift = 0
	RegisterMdiosrdfrFieldRdfMask  = 0xffffffff
)

// GetRdf Read flags for MDIO registers 0 to 31
func (r *RegisterMdiosrdfrType) GetRdf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosrdfrFieldRdfMask) >> RegisterMdiosrdfrFieldRdfShift)
}

// SetRdf Read flags for MDIO registers 0 to 31
func (r *RegisterMdiosrdfrType) SetRdf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosrdfrFieldRdfMask)|(uint32(value)<<RegisterMdiosrdfrFieldRdfShift))
}

// RegisterMdioscrdfrType MDIOS clear read flag register
type RegisterMdioscrdfrType uint32

func (r *RegisterMdioscrdfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdioscrdfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdioscrdfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdioscrdfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdioscrdfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdioscrdfrFieldCrdfShift = 0
	RegisterMdioscrdfrFieldCrdfMask  = 0xffffffff
)

// GetCrdf Clear the read flag
func (r *RegisterMdioscrdfrType) GetCrdf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdioscrdfrFieldCrdfMask) >> RegisterMdioscrdfrFieldCrdfShift)
}

// SetCrdf Clear the read flag
func (r *RegisterMdioscrdfrType) SetCrdf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdioscrdfrFieldCrdfMask)|(uint32(value)<<RegisterMdioscrdfrFieldCrdfShift))
}

// RegisterMdiossrType MDIOS status register
type RegisterMdiossrType uint32

func (r *RegisterMdiossrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiossrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiossrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiossrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiossrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiossrFieldPerfShift = 0
	RegisterMdiossrFieldPerfMask  = 0x1
)

// GetPerf Preamble error flag
func (r *RegisterMdiossrType) GetPerf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdiossrFieldPerfMask) != 0
}

// SetPerf Preamble error flag
func (r *RegisterMdiossrType) SetPerf(value bool) {
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
func (r *RegisterMdiossrType) GetSerf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdiossrFieldSerfMask) != 0
}

// SetSerf Start error flag
func (r *RegisterMdiossrType) SetSerf(value bool) {
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
func (r *RegisterMdiossrType) GetTerf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdiossrFieldTerfMask) != 0
}

// SetTerf Turnaround error flag
func (r *RegisterMdiossrType) SetTerf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdiossrFieldTerfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdiossrFieldTerfMask)
	}
}

// RegisterMdiosclrfrType MDIOS clear flag register
type RegisterMdiosclrfrType uint32

func (r *RegisterMdiosclrfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosclrfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosclrfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosclrfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosclrfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosclrfrFieldCperfShift = 0
	RegisterMdiosclrfrFieldCperfMask  = 0x1
)

// GetCperf Clear the preamble error flag
func (r *RegisterMdiosclrfrType) GetCperf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdiosclrfrFieldCperfMask) != 0
}

// SetCperf Clear the preamble error flag
func (r *RegisterMdiosclrfrType) SetCperf(value bool) {
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
func (r *RegisterMdiosclrfrType) GetCserf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdiosclrfrFieldCserfMask) != 0
}

// SetCserf Clear the start error flag
func (r *RegisterMdiosclrfrType) SetCserf(value bool) {
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
func (r *RegisterMdiosclrfrType) GetCterf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdiosclrfrFieldCterfMask) != 0
}

// SetCterf Clear the turnaround error flag
func (r *RegisterMdiosclrfrType) SetCterf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdiosclrfrFieldCterfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdiosclrfrFieldCterfMask)
	}
}

// RegisterMdiosdinr0Type MDIOS input data register 0
type RegisterMdiosdinr0Type uint32

func (r *RegisterMdiosdinr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr0FieldDin0Shift = 0
	RegisterMdiosdinr0FieldDin0Mask  = 0xffff
)

// GetDin0 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr0Type) GetDin0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr0FieldDin0Mask) >> RegisterMdiosdinr0FieldDin0Shift)
}

// SetDin0 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr0Type) SetDin0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr0FieldDin0Mask)|(uint32(value)<<RegisterMdiosdinr0FieldDin0Shift))
}

// RegisterMdiosdinr1Type MDIOS input data register 1
type RegisterMdiosdinr1Type uint32

func (r *RegisterMdiosdinr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr1FieldDin1Shift = 0
	RegisterMdiosdinr1FieldDin1Mask  = 0xffff
)

// GetDin1 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr1Type) GetDin1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr1FieldDin1Mask) >> RegisterMdiosdinr1FieldDin1Shift)
}

// SetDin1 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr1Type) SetDin1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr1FieldDin1Mask)|(uint32(value)<<RegisterMdiosdinr1FieldDin1Shift))
}

// RegisterMdiosdinr2Type MDIOS input data register 2
type RegisterMdiosdinr2Type uint32

func (r *RegisterMdiosdinr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr2FieldDin2Shift = 0
	RegisterMdiosdinr2FieldDin2Mask  = 0xffff
)

// GetDin2 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr2Type) GetDin2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr2FieldDin2Mask) >> RegisterMdiosdinr2FieldDin2Shift)
}

// SetDin2 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr2Type) SetDin2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr2FieldDin2Mask)|(uint32(value)<<RegisterMdiosdinr2FieldDin2Shift))
}

// RegisterMdiosdinr3Type MDIOS input data register 3
type RegisterMdiosdinr3Type uint32

func (r *RegisterMdiosdinr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr3FieldDin3Shift = 0
	RegisterMdiosdinr3FieldDin3Mask  = 0xffff
)

// GetDin3 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr3Type) GetDin3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr3FieldDin3Mask) >> RegisterMdiosdinr3FieldDin3Shift)
}

// SetDin3 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr3Type) SetDin3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr3FieldDin3Mask)|(uint32(value)<<RegisterMdiosdinr3FieldDin3Shift))
}

// RegisterMdiosdinr4Type MDIOS input data register 4
type RegisterMdiosdinr4Type uint32

func (r *RegisterMdiosdinr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr4FieldDin4Shift = 0
	RegisterMdiosdinr4FieldDin4Mask  = 0xffff
)

// GetDin4 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr4Type) GetDin4() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr4FieldDin4Mask) >> RegisterMdiosdinr4FieldDin4Shift)
}

// SetDin4 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr4Type) SetDin4(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr4FieldDin4Mask)|(uint32(value)<<RegisterMdiosdinr4FieldDin4Shift))
}

// RegisterMdiosdinr5Type MDIOS input data register 5
type RegisterMdiosdinr5Type uint32

func (r *RegisterMdiosdinr5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr5FieldDin5Shift = 0
	RegisterMdiosdinr5FieldDin5Mask  = 0xffff
)

// GetDin5 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr5Type) GetDin5() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr5FieldDin5Mask) >> RegisterMdiosdinr5FieldDin5Shift)
}

// SetDin5 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr5Type) SetDin5(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr5FieldDin5Mask)|(uint32(value)<<RegisterMdiosdinr5FieldDin5Shift))
}

// RegisterMdiosdinr6Type MDIOS input data register 6
type RegisterMdiosdinr6Type uint32

func (r *RegisterMdiosdinr6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr6FieldDin6Shift = 0
	RegisterMdiosdinr6FieldDin6Mask  = 0xffff
)

// GetDin6 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr6Type) GetDin6() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr6FieldDin6Mask) >> RegisterMdiosdinr6FieldDin6Shift)
}

// SetDin6 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr6Type) SetDin6(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr6FieldDin6Mask)|(uint32(value)<<RegisterMdiosdinr6FieldDin6Shift))
}

// RegisterMdiosdinr7Type MDIOS input data register 7
type RegisterMdiosdinr7Type uint32

func (r *RegisterMdiosdinr7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr7FieldDin7Shift = 0
	RegisterMdiosdinr7FieldDin7Mask  = 0xffff
)

// GetDin7 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr7Type) GetDin7() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr7FieldDin7Mask) >> RegisterMdiosdinr7FieldDin7Shift)
}

// SetDin7 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr7Type) SetDin7(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr7FieldDin7Mask)|(uint32(value)<<RegisterMdiosdinr7FieldDin7Shift))
}

// RegisterMdiosdinr8Type MDIOS input data register 8
type RegisterMdiosdinr8Type uint32

func (r *RegisterMdiosdinr8Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr8Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr8Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr8Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr8Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr8FieldDin8Shift = 0
	RegisterMdiosdinr8FieldDin8Mask  = 0xffff
)

// GetDin8 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr8Type) GetDin8() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr8FieldDin8Mask) >> RegisterMdiosdinr8FieldDin8Shift)
}

// SetDin8 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr8Type) SetDin8(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr8FieldDin8Mask)|(uint32(value)<<RegisterMdiosdinr8FieldDin8Shift))
}

// RegisterMdiosdinr9Type MDIOS input data register 9
type RegisterMdiosdinr9Type uint32

func (r *RegisterMdiosdinr9Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr9Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr9Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr9Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr9Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr9FieldDin9Shift = 0
	RegisterMdiosdinr9FieldDin9Mask  = 0xffff
)

// GetDin9 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr9Type) GetDin9() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr9FieldDin9Mask) >> RegisterMdiosdinr9FieldDin9Shift)
}

// SetDin9 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr9Type) SetDin9(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr9FieldDin9Mask)|(uint32(value)<<RegisterMdiosdinr9FieldDin9Shift))
}

// RegisterMdiosdinr10Type MDIOS input data register 10
type RegisterMdiosdinr10Type uint32

func (r *RegisterMdiosdinr10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr10FieldDin10Shift = 0
	RegisterMdiosdinr10FieldDin10Mask  = 0xffff
)

// GetDin10 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr10Type) GetDin10() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr10FieldDin10Mask) >> RegisterMdiosdinr10FieldDin10Shift)
}

// SetDin10 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr10Type) SetDin10(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr10FieldDin10Mask)|(uint32(value)<<RegisterMdiosdinr10FieldDin10Shift))
}

// RegisterMdiosdinr11Type MDIOS input data register 11
type RegisterMdiosdinr11Type uint32

func (r *RegisterMdiosdinr11Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr11Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr11Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr11Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr11Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr11FieldDin11Shift = 0
	RegisterMdiosdinr11FieldDin11Mask  = 0xffff
)

// GetDin11 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr11Type) GetDin11() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr11FieldDin11Mask) >> RegisterMdiosdinr11FieldDin11Shift)
}

// SetDin11 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr11Type) SetDin11(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr11FieldDin11Mask)|(uint32(value)<<RegisterMdiosdinr11FieldDin11Shift))
}

// RegisterMdiosdinr12Type MDIOS input data register 12
type RegisterMdiosdinr12Type uint32

func (r *RegisterMdiosdinr12Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr12Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr12Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr12Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr12Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr12FieldDin12Shift = 0
	RegisterMdiosdinr12FieldDin12Mask  = 0xffff
)

// GetDin12 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr12Type) GetDin12() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr12FieldDin12Mask) >> RegisterMdiosdinr12FieldDin12Shift)
}

// SetDin12 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr12Type) SetDin12(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr12FieldDin12Mask)|(uint32(value)<<RegisterMdiosdinr12FieldDin12Shift))
}

// RegisterMdiosdinr13Type MDIOS input data register 13
type RegisterMdiosdinr13Type uint32

func (r *RegisterMdiosdinr13Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr13Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr13Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr13Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr13Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr13FieldDin13Shift = 0
	RegisterMdiosdinr13FieldDin13Mask  = 0xffff
)

// GetDin13 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr13Type) GetDin13() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr13FieldDin13Mask) >> RegisterMdiosdinr13FieldDin13Shift)
}

// SetDin13 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr13Type) SetDin13(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr13FieldDin13Mask)|(uint32(value)<<RegisterMdiosdinr13FieldDin13Shift))
}

// RegisterMdiosdinr14Type MDIOS input data register 14
type RegisterMdiosdinr14Type uint32

func (r *RegisterMdiosdinr14Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr14Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr14Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr14Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr14Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr14FieldDin14Shift = 0
	RegisterMdiosdinr14FieldDin14Mask  = 0xffff
)

// GetDin14 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr14Type) GetDin14() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr14FieldDin14Mask) >> RegisterMdiosdinr14FieldDin14Shift)
}

// SetDin14 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr14Type) SetDin14(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr14FieldDin14Mask)|(uint32(value)<<RegisterMdiosdinr14FieldDin14Shift))
}

// RegisterMdiosdinr15Type MDIOS input data register 15
type RegisterMdiosdinr15Type uint32

func (r *RegisterMdiosdinr15Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr15Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr15Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr15Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr15Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr15FieldDin15Shift = 0
	RegisterMdiosdinr15FieldDin15Mask  = 0xffff
)

// GetDin15 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr15Type) GetDin15() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr15FieldDin15Mask) >> RegisterMdiosdinr15FieldDin15Shift)
}

// SetDin15 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr15Type) SetDin15(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr15FieldDin15Mask)|(uint32(value)<<RegisterMdiosdinr15FieldDin15Shift))
}

// RegisterMdiosdinr16Type MDIOS input data register 16
type RegisterMdiosdinr16Type uint32

func (r *RegisterMdiosdinr16Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr16Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr16Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr16Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr16Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr16FieldDin16Shift = 0
	RegisterMdiosdinr16FieldDin16Mask  = 0xffff
)

// GetDin16 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr16Type) GetDin16() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr16FieldDin16Mask) >> RegisterMdiosdinr16FieldDin16Shift)
}

// SetDin16 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr16Type) SetDin16(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr16FieldDin16Mask)|(uint32(value)<<RegisterMdiosdinr16FieldDin16Shift))
}

// RegisterMdiosdinr17Type MDIOS input data register 17
type RegisterMdiosdinr17Type uint32

func (r *RegisterMdiosdinr17Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr17Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr17Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr17Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr17Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr17FieldDin17Shift = 0
	RegisterMdiosdinr17FieldDin17Mask  = 0xffff
)

// GetDin17 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr17Type) GetDin17() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr17FieldDin17Mask) >> RegisterMdiosdinr17FieldDin17Shift)
}

// SetDin17 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr17Type) SetDin17(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr17FieldDin17Mask)|(uint32(value)<<RegisterMdiosdinr17FieldDin17Shift))
}

// RegisterMdiosdinr18Type MDIOS input data register 18
type RegisterMdiosdinr18Type uint32

func (r *RegisterMdiosdinr18Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr18Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr18Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr18Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr18Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr18FieldDin18Shift = 0
	RegisterMdiosdinr18FieldDin18Mask  = 0xffff
)

// GetDin18 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr18Type) GetDin18() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr18FieldDin18Mask) >> RegisterMdiosdinr18FieldDin18Shift)
}

// SetDin18 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr18Type) SetDin18(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr18FieldDin18Mask)|(uint32(value)<<RegisterMdiosdinr18FieldDin18Shift))
}

// RegisterMdiosdinr19Type MDIOS input data register 19
type RegisterMdiosdinr19Type uint32

func (r *RegisterMdiosdinr19Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr19Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr19Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr19Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr19Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr19FieldDin19Shift = 0
	RegisterMdiosdinr19FieldDin19Mask  = 0xffff
)

// GetDin19 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr19Type) GetDin19() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr19FieldDin19Mask) >> RegisterMdiosdinr19FieldDin19Shift)
}

// SetDin19 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr19Type) SetDin19(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr19FieldDin19Mask)|(uint32(value)<<RegisterMdiosdinr19FieldDin19Shift))
}

// RegisterMdiosdinr20Type MDIOS input data register 20
type RegisterMdiosdinr20Type uint32

func (r *RegisterMdiosdinr20Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr20Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr20Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr20Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr20Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr20FieldDin20Shift = 0
	RegisterMdiosdinr20FieldDin20Mask  = 0xffff
)

// GetDin20 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr20Type) GetDin20() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr20FieldDin20Mask) >> RegisterMdiosdinr20FieldDin20Shift)
}

// SetDin20 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr20Type) SetDin20(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr20FieldDin20Mask)|(uint32(value)<<RegisterMdiosdinr20FieldDin20Shift))
}

// RegisterMdiosdinr21Type MDIOS input data register 21
type RegisterMdiosdinr21Type uint32

func (r *RegisterMdiosdinr21Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr21Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr21Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr21Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr21Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr21FieldDin21Shift = 0
	RegisterMdiosdinr21FieldDin21Mask  = 0xffff
)

// GetDin21 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr21Type) GetDin21() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr21FieldDin21Mask) >> RegisterMdiosdinr21FieldDin21Shift)
}

// SetDin21 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr21Type) SetDin21(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr21FieldDin21Mask)|(uint32(value)<<RegisterMdiosdinr21FieldDin21Shift))
}

// RegisterMdiosdinr22Type MDIOS input data register 22
type RegisterMdiosdinr22Type uint32

func (r *RegisterMdiosdinr22Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr22Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr22Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr22Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr22Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr22FieldDin22Shift = 0
	RegisterMdiosdinr22FieldDin22Mask  = 0xffff
)

// GetDin22 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr22Type) GetDin22() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr22FieldDin22Mask) >> RegisterMdiosdinr22FieldDin22Shift)
}

// SetDin22 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr22Type) SetDin22(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr22FieldDin22Mask)|(uint32(value)<<RegisterMdiosdinr22FieldDin22Shift))
}

// RegisterMdiosdinr23Type MDIOS input data register 23
type RegisterMdiosdinr23Type uint32

func (r *RegisterMdiosdinr23Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr23Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr23Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr23Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr23Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr23FieldDin23Shift = 0
	RegisterMdiosdinr23FieldDin23Mask  = 0xffff
)

// GetDin23 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr23Type) GetDin23() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr23FieldDin23Mask) >> RegisterMdiosdinr23FieldDin23Shift)
}

// SetDin23 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr23Type) SetDin23(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr23FieldDin23Mask)|(uint32(value)<<RegisterMdiosdinr23FieldDin23Shift))
}

// RegisterMdiosdinr24Type MDIOS input data register 24
type RegisterMdiosdinr24Type uint32

func (r *RegisterMdiosdinr24Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr24Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr24Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr24Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr24Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr24FieldDin24Shift = 0
	RegisterMdiosdinr24FieldDin24Mask  = 0xffff
)

// GetDin24 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr24Type) GetDin24() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr24FieldDin24Mask) >> RegisterMdiosdinr24FieldDin24Shift)
}

// SetDin24 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr24Type) SetDin24(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr24FieldDin24Mask)|(uint32(value)<<RegisterMdiosdinr24FieldDin24Shift))
}

// RegisterMdiosdinr25Type MDIOS input data register 25
type RegisterMdiosdinr25Type uint32

func (r *RegisterMdiosdinr25Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr25Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr25Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr25Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr25Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr25FieldDin25Shift = 0
	RegisterMdiosdinr25FieldDin25Mask  = 0xffff
)

// GetDin25 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr25Type) GetDin25() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr25FieldDin25Mask) >> RegisterMdiosdinr25FieldDin25Shift)
}

// SetDin25 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr25Type) SetDin25(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr25FieldDin25Mask)|(uint32(value)<<RegisterMdiosdinr25FieldDin25Shift))
}

// RegisterMdiosdinr26Type MDIOS input data register 26
type RegisterMdiosdinr26Type uint32

func (r *RegisterMdiosdinr26Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr26Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr26Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr26Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr26Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr26FieldDin26Shift = 0
	RegisterMdiosdinr26FieldDin26Mask  = 0xffff
)

// GetDin26 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr26Type) GetDin26() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr26FieldDin26Mask) >> RegisterMdiosdinr26FieldDin26Shift)
}

// SetDin26 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr26Type) SetDin26(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr26FieldDin26Mask)|(uint32(value)<<RegisterMdiosdinr26FieldDin26Shift))
}

// RegisterMdiosdinr27Type MDIOS input data register 27
type RegisterMdiosdinr27Type uint32

func (r *RegisterMdiosdinr27Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr27Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr27Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr27Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr27Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr27FieldDin27Shift = 0
	RegisterMdiosdinr27FieldDin27Mask  = 0xffff
)

// GetDin27 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr27Type) GetDin27() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr27FieldDin27Mask) >> RegisterMdiosdinr27FieldDin27Shift)
}

// SetDin27 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr27Type) SetDin27(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr27FieldDin27Mask)|(uint32(value)<<RegisterMdiosdinr27FieldDin27Shift))
}

// RegisterMdiosdinr28Type MDIOS input data register 28
type RegisterMdiosdinr28Type uint32

func (r *RegisterMdiosdinr28Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr28Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr28Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr28Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr28Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr28FieldDin28Shift = 0
	RegisterMdiosdinr28FieldDin28Mask  = 0xffff
)

// GetDin28 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr28Type) GetDin28() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr28FieldDin28Mask) >> RegisterMdiosdinr28FieldDin28Shift)
}

// SetDin28 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr28Type) SetDin28(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr28FieldDin28Mask)|(uint32(value)<<RegisterMdiosdinr28FieldDin28Shift))
}

// RegisterMdiosdinr29Type MDIOS input data register 29
type RegisterMdiosdinr29Type uint32

func (r *RegisterMdiosdinr29Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr29Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr29Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr29Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr29Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr29FieldDin29Shift = 0
	RegisterMdiosdinr29FieldDin29Mask  = 0xffff
)

// GetDin29 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr29Type) GetDin29() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr29FieldDin29Mask) >> RegisterMdiosdinr29FieldDin29Shift)
}

// SetDin29 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr29Type) SetDin29(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr29FieldDin29Mask)|(uint32(value)<<RegisterMdiosdinr29FieldDin29Shift))
}

// RegisterMdiosdinr30Type MDIOS input data register 30
type RegisterMdiosdinr30Type uint32

func (r *RegisterMdiosdinr30Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr30Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr30Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr30Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr30Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr30FieldDin30Shift = 0
	RegisterMdiosdinr30FieldDin30Mask  = 0xffff
)

// GetDin30 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr30Type) GetDin30() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr30FieldDin30Mask) >> RegisterMdiosdinr30FieldDin30Shift)
}

// SetDin30 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr30Type) SetDin30(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr30FieldDin30Mask)|(uint32(value)<<RegisterMdiosdinr30FieldDin30Shift))
}

// RegisterMdiosdinr31Type MDIOS input data register 31
type RegisterMdiosdinr31Type uint32

func (r *RegisterMdiosdinr31Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdinr31Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdinr31Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdinr31Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdinr31Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdinr31FieldDin31Shift = 0
	RegisterMdiosdinr31FieldDin31Mask  = 0xffff
)

// GetDin31 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr31Type) GetDin31() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdinr31FieldDin31Mask) >> RegisterMdiosdinr31FieldDin31Shift)
}

// SetDin31 Input data received from MDIO Master during write frames
func (r *RegisterMdiosdinr31Type) SetDin31(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdinr31FieldDin31Mask)|(uint32(value)<<RegisterMdiosdinr31FieldDin31Shift))
}

// RegisterMdiosdoutr0Type MDIOS output data register 0
type RegisterMdiosdoutr0Type uint32

func (r *RegisterMdiosdoutr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr0FieldDout0Shift = 0
	RegisterMdiosdoutr0FieldDout0Mask  = 0xffff
)

// GetDout0 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr0Type) GetDout0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr0FieldDout0Mask) >> RegisterMdiosdoutr0FieldDout0Shift)
}

// SetDout0 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr0Type) SetDout0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr0FieldDout0Mask)|(uint32(value)<<RegisterMdiosdoutr0FieldDout0Shift))
}

// RegisterMdiosdoutr1Type MDIOS output data register 1
type RegisterMdiosdoutr1Type uint32

func (r *RegisterMdiosdoutr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr1FieldDout1Shift = 0
	RegisterMdiosdoutr1FieldDout1Mask  = 0xffff
)

// GetDout1 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr1Type) GetDout1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr1FieldDout1Mask) >> RegisterMdiosdoutr1FieldDout1Shift)
}

// SetDout1 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr1Type) SetDout1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr1FieldDout1Mask)|(uint32(value)<<RegisterMdiosdoutr1FieldDout1Shift))
}

// RegisterMdiosdoutr2Type MDIOS output data register 2
type RegisterMdiosdoutr2Type uint32

func (r *RegisterMdiosdoutr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr2FieldDout2Shift = 0
	RegisterMdiosdoutr2FieldDout2Mask  = 0xffff
)

// GetDout2 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr2Type) GetDout2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr2FieldDout2Mask) >> RegisterMdiosdoutr2FieldDout2Shift)
}

// SetDout2 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr2Type) SetDout2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr2FieldDout2Mask)|(uint32(value)<<RegisterMdiosdoutr2FieldDout2Shift))
}

// RegisterMdiosdoutr3Type MDIOS output data register 3
type RegisterMdiosdoutr3Type uint32

func (r *RegisterMdiosdoutr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr3FieldDout3Shift = 0
	RegisterMdiosdoutr3FieldDout3Mask  = 0xffff
)

// GetDout3 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr3Type) GetDout3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr3FieldDout3Mask) >> RegisterMdiosdoutr3FieldDout3Shift)
}

// SetDout3 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr3Type) SetDout3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr3FieldDout3Mask)|(uint32(value)<<RegisterMdiosdoutr3FieldDout3Shift))
}

// RegisterMdiosdoutr4Type MDIOS output data register 4
type RegisterMdiosdoutr4Type uint32

func (r *RegisterMdiosdoutr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr4FieldDout4Shift = 0
	RegisterMdiosdoutr4FieldDout4Mask  = 0xffff
)

// GetDout4 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr4Type) GetDout4() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr4FieldDout4Mask) >> RegisterMdiosdoutr4FieldDout4Shift)
}

// SetDout4 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr4Type) SetDout4(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr4FieldDout4Mask)|(uint32(value)<<RegisterMdiosdoutr4FieldDout4Shift))
}

// RegisterMdiosdoutr5Type MDIOS output data register 5
type RegisterMdiosdoutr5Type uint32

func (r *RegisterMdiosdoutr5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr5FieldDout5Shift = 0
	RegisterMdiosdoutr5FieldDout5Mask  = 0xffff
)

// GetDout5 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr5Type) GetDout5() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr5FieldDout5Mask) >> RegisterMdiosdoutr5FieldDout5Shift)
}

// SetDout5 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr5Type) SetDout5(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr5FieldDout5Mask)|(uint32(value)<<RegisterMdiosdoutr5FieldDout5Shift))
}

// RegisterMdiosdoutr6Type MDIOS output data register 6
type RegisterMdiosdoutr6Type uint32

func (r *RegisterMdiosdoutr6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr6FieldDout6Shift = 0
	RegisterMdiosdoutr6FieldDout6Mask  = 0xffff
)

// GetDout6 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr6Type) GetDout6() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr6FieldDout6Mask) >> RegisterMdiosdoutr6FieldDout6Shift)
}

// SetDout6 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr6Type) SetDout6(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr6FieldDout6Mask)|(uint32(value)<<RegisterMdiosdoutr6FieldDout6Shift))
}

// RegisterMdiosdoutr7Type MDIOS output data register 7
type RegisterMdiosdoutr7Type uint32

func (r *RegisterMdiosdoutr7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr7FieldDout7Shift = 0
	RegisterMdiosdoutr7FieldDout7Mask  = 0xffff
)

// GetDout7 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr7Type) GetDout7() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr7FieldDout7Mask) >> RegisterMdiosdoutr7FieldDout7Shift)
}

// SetDout7 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr7Type) SetDout7(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr7FieldDout7Mask)|(uint32(value)<<RegisterMdiosdoutr7FieldDout7Shift))
}

// RegisterMdiosdoutr8Type MDIOS output data register 8
type RegisterMdiosdoutr8Type uint32

func (r *RegisterMdiosdoutr8Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr8Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr8Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr8Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr8Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr8FieldDout8Shift = 0
	RegisterMdiosdoutr8FieldDout8Mask  = 0xffff
)

// GetDout8 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr8Type) GetDout8() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr8FieldDout8Mask) >> RegisterMdiosdoutr8FieldDout8Shift)
}

// SetDout8 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr8Type) SetDout8(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr8FieldDout8Mask)|(uint32(value)<<RegisterMdiosdoutr8FieldDout8Shift))
}

// RegisterMdiosdoutr9Type MDIOS output data register 9
type RegisterMdiosdoutr9Type uint32

func (r *RegisterMdiosdoutr9Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr9Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr9Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr9Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr9Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr9FieldDout9Shift = 0
	RegisterMdiosdoutr9FieldDout9Mask  = 0xffff
)

// GetDout9 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr9Type) GetDout9() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr9FieldDout9Mask) >> RegisterMdiosdoutr9FieldDout9Shift)
}

// SetDout9 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr9Type) SetDout9(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr9FieldDout9Mask)|(uint32(value)<<RegisterMdiosdoutr9FieldDout9Shift))
}

// RegisterMdiosdoutr10Type MDIOS output data register 10
type RegisterMdiosdoutr10Type uint32

func (r *RegisterMdiosdoutr10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr10FieldDout10Shift = 0
	RegisterMdiosdoutr10FieldDout10Mask  = 0xffff
)

// GetDout10 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr10Type) GetDout10() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr10FieldDout10Mask) >> RegisterMdiosdoutr10FieldDout10Shift)
}

// SetDout10 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr10Type) SetDout10(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr10FieldDout10Mask)|(uint32(value)<<RegisterMdiosdoutr10FieldDout10Shift))
}

// RegisterMdiosdoutr11Type MDIOS output data register 11
type RegisterMdiosdoutr11Type uint32

func (r *RegisterMdiosdoutr11Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr11Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr11Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr11Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr11Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr11FieldDout11Shift = 0
	RegisterMdiosdoutr11FieldDout11Mask  = 0xffff
)

// GetDout11 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr11Type) GetDout11() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr11FieldDout11Mask) >> RegisterMdiosdoutr11FieldDout11Shift)
}

// SetDout11 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr11Type) SetDout11(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr11FieldDout11Mask)|(uint32(value)<<RegisterMdiosdoutr11FieldDout11Shift))
}

// RegisterMdiosdoutr12Type MDIOS output data register 12
type RegisterMdiosdoutr12Type uint32

func (r *RegisterMdiosdoutr12Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr12Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr12Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr12Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr12Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr12FieldDout12Shift = 0
	RegisterMdiosdoutr12FieldDout12Mask  = 0xffff
)

// GetDout12 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr12Type) GetDout12() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr12FieldDout12Mask) >> RegisterMdiosdoutr12FieldDout12Shift)
}

// SetDout12 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr12Type) SetDout12(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr12FieldDout12Mask)|(uint32(value)<<RegisterMdiosdoutr12FieldDout12Shift))
}

// RegisterMdiosdoutr13Type MDIOS output data register 13
type RegisterMdiosdoutr13Type uint32

func (r *RegisterMdiosdoutr13Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr13Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr13Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr13Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr13Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr13FieldDout13Shift = 0
	RegisterMdiosdoutr13FieldDout13Mask  = 0xffff
)

// GetDout13 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr13Type) GetDout13() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr13FieldDout13Mask) >> RegisterMdiosdoutr13FieldDout13Shift)
}

// SetDout13 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr13Type) SetDout13(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr13FieldDout13Mask)|(uint32(value)<<RegisterMdiosdoutr13FieldDout13Shift))
}

// RegisterMdiosdoutr14Type MDIOS output data register 14
type RegisterMdiosdoutr14Type uint32

func (r *RegisterMdiosdoutr14Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr14Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr14Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr14Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr14Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr14FieldDout14Shift = 0
	RegisterMdiosdoutr14FieldDout14Mask  = 0xffff
)

// GetDout14 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr14Type) GetDout14() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr14FieldDout14Mask) >> RegisterMdiosdoutr14FieldDout14Shift)
}

// SetDout14 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr14Type) SetDout14(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr14FieldDout14Mask)|(uint32(value)<<RegisterMdiosdoutr14FieldDout14Shift))
}

// RegisterMdiosdoutr15Type MDIOS output data register 15
type RegisterMdiosdoutr15Type uint32

func (r *RegisterMdiosdoutr15Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr15Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr15Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr15Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr15Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr15FieldDout15Shift = 0
	RegisterMdiosdoutr15FieldDout15Mask  = 0xffff
)

// GetDout15 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr15Type) GetDout15() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr15FieldDout15Mask) >> RegisterMdiosdoutr15FieldDout15Shift)
}

// SetDout15 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr15Type) SetDout15(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr15FieldDout15Mask)|(uint32(value)<<RegisterMdiosdoutr15FieldDout15Shift))
}

// RegisterMdiosdoutr16Type MDIOS output data register 16
type RegisterMdiosdoutr16Type uint32

func (r *RegisterMdiosdoutr16Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr16Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr16Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr16Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr16Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr16FieldDout16Shift = 0
	RegisterMdiosdoutr16FieldDout16Mask  = 0xffff
)

// GetDout16 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr16Type) GetDout16() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr16FieldDout16Mask) >> RegisterMdiosdoutr16FieldDout16Shift)
}

// SetDout16 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr16Type) SetDout16(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr16FieldDout16Mask)|(uint32(value)<<RegisterMdiosdoutr16FieldDout16Shift))
}

// RegisterMdiosdoutr17Type MDIOS output data register 17
type RegisterMdiosdoutr17Type uint32

func (r *RegisterMdiosdoutr17Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr17Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr17Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr17Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr17Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr17FieldDout17Shift = 0
	RegisterMdiosdoutr17FieldDout17Mask  = 0xffff
)

// GetDout17 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr17Type) GetDout17() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr17FieldDout17Mask) >> RegisterMdiosdoutr17FieldDout17Shift)
}

// SetDout17 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr17Type) SetDout17(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr17FieldDout17Mask)|(uint32(value)<<RegisterMdiosdoutr17FieldDout17Shift))
}

// RegisterMdiosdoutr18Type MDIOS output data register 18
type RegisterMdiosdoutr18Type uint32

func (r *RegisterMdiosdoutr18Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr18Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr18Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr18Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr18Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr18FieldDout18Shift = 0
	RegisterMdiosdoutr18FieldDout18Mask  = 0xffff
)

// GetDout18 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr18Type) GetDout18() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr18FieldDout18Mask) >> RegisterMdiosdoutr18FieldDout18Shift)
}

// SetDout18 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr18Type) SetDout18(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr18FieldDout18Mask)|(uint32(value)<<RegisterMdiosdoutr18FieldDout18Shift))
}

// RegisterMdiosdoutr19Type MDIOS output data register 19
type RegisterMdiosdoutr19Type uint32

func (r *RegisterMdiosdoutr19Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr19Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr19Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr19Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr19Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr19FieldDout19Shift = 0
	RegisterMdiosdoutr19FieldDout19Mask  = 0xffff
)

// GetDout19 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr19Type) GetDout19() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr19FieldDout19Mask) >> RegisterMdiosdoutr19FieldDout19Shift)
}

// SetDout19 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr19Type) SetDout19(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr19FieldDout19Mask)|(uint32(value)<<RegisterMdiosdoutr19FieldDout19Shift))
}

// RegisterMdiosdoutr20Type MDIOS output data register 20
type RegisterMdiosdoutr20Type uint32

func (r *RegisterMdiosdoutr20Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr20Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr20Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr20Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr20Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr20FieldDout20Shift = 0
	RegisterMdiosdoutr20FieldDout20Mask  = 0xffff
)

// GetDout20 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr20Type) GetDout20() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr20FieldDout20Mask) >> RegisterMdiosdoutr20FieldDout20Shift)
}

// SetDout20 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr20Type) SetDout20(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr20FieldDout20Mask)|(uint32(value)<<RegisterMdiosdoutr20FieldDout20Shift))
}

// RegisterMdiosdoutr21Type MDIOS output data register 21
type RegisterMdiosdoutr21Type uint32

func (r *RegisterMdiosdoutr21Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr21Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr21Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr21Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr21Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr21FieldDout21Shift = 0
	RegisterMdiosdoutr21FieldDout21Mask  = 0xffff
)

// GetDout21 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr21Type) GetDout21() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr21FieldDout21Mask) >> RegisterMdiosdoutr21FieldDout21Shift)
}

// SetDout21 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr21Type) SetDout21(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr21FieldDout21Mask)|(uint32(value)<<RegisterMdiosdoutr21FieldDout21Shift))
}

// RegisterMdiosdoutr22Type MDIOS output data register 22
type RegisterMdiosdoutr22Type uint32

func (r *RegisterMdiosdoutr22Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr22Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr22Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr22Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr22Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr22FieldDout22Shift = 0
	RegisterMdiosdoutr22FieldDout22Mask  = 0xffff
)

// GetDout22 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr22Type) GetDout22() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr22FieldDout22Mask) >> RegisterMdiosdoutr22FieldDout22Shift)
}

// SetDout22 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr22Type) SetDout22(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr22FieldDout22Mask)|(uint32(value)<<RegisterMdiosdoutr22FieldDout22Shift))
}

// RegisterMdiosdoutr23Type MDIOS output data register 23
type RegisterMdiosdoutr23Type uint32

func (r *RegisterMdiosdoutr23Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr23Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr23Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr23Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr23Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr23FieldDout23Shift = 0
	RegisterMdiosdoutr23FieldDout23Mask  = 0xffff
)

// GetDout23 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr23Type) GetDout23() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr23FieldDout23Mask) >> RegisterMdiosdoutr23FieldDout23Shift)
}

// SetDout23 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr23Type) SetDout23(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr23FieldDout23Mask)|(uint32(value)<<RegisterMdiosdoutr23FieldDout23Shift))
}

// RegisterMdiosdoutr24Type MDIOS output data register 24
type RegisterMdiosdoutr24Type uint32

func (r *RegisterMdiosdoutr24Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr24Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr24Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr24Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr24Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr24FieldDout24Shift = 0
	RegisterMdiosdoutr24FieldDout24Mask  = 0xffff
)

// GetDout24 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr24Type) GetDout24() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr24FieldDout24Mask) >> RegisterMdiosdoutr24FieldDout24Shift)
}

// SetDout24 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr24Type) SetDout24(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr24FieldDout24Mask)|(uint32(value)<<RegisterMdiosdoutr24FieldDout24Shift))
}

// RegisterMdiosdoutr25Type MDIOS output data register 25
type RegisterMdiosdoutr25Type uint32

func (r *RegisterMdiosdoutr25Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr25Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr25Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr25Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr25Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr25FieldDout25Shift = 0
	RegisterMdiosdoutr25FieldDout25Mask  = 0xffff
)

// GetDout25 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr25Type) GetDout25() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr25FieldDout25Mask) >> RegisterMdiosdoutr25FieldDout25Shift)
}

// SetDout25 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr25Type) SetDout25(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr25FieldDout25Mask)|(uint32(value)<<RegisterMdiosdoutr25FieldDout25Shift))
}

// RegisterMdiosdoutr26Type MDIOS output data register 26
type RegisterMdiosdoutr26Type uint32

func (r *RegisterMdiosdoutr26Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr26Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr26Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr26Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr26Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr26FieldDout26Shift = 0
	RegisterMdiosdoutr26FieldDout26Mask  = 0xffff
)

// GetDout26 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr26Type) GetDout26() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr26FieldDout26Mask) >> RegisterMdiosdoutr26FieldDout26Shift)
}

// SetDout26 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr26Type) SetDout26(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr26FieldDout26Mask)|(uint32(value)<<RegisterMdiosdoutr26FieldDout26Shift))
}

// RegisterMdiosdoutr27Type MDIOS output data register 27
type RegisterMdiosdoutr27Type uint32

func (r *RegisterMdiosdoutr27Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr27Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr27Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr27Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr27Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr27FieldDout27Shift = 0
	RegisterMdiosdoutr27FieldDout27Mask  = 0xffff
)

// GetDout27 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr27Type) GetDout27() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr27FieldDout27Mask) >> RegisterMdiosdoutr27FieldDout27Shift)
}

// SetDout27 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr27Type) SetDout27(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr27FieldDout27Mask)|(uint32(value)<<RegisterMdiosdoutr27FieldDout27Shift))
}

// RegisterMdiosdoutr28Type MDIOS output data register 28
type RegisterMdiosdoutr28Type uint32

func (r *RegisterMdiosdoutr28Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr28Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr28Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr28Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr28Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr28FieldDout28Shift = 0
	RegisterMdiosdoutr28FieldDout28Mask  = 0xffff
)

// GetDout28 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr28Type) GetDout28() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr28FieldDout28Mask) >> RegisterMdiosdoutr28FieldDout28Shift)
}

// SetDout28 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr28Type) SetDout28(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr28FieldDout28Mask)|(uint32(value)<<RegisterMdiosdoutr28FieldDout28Shift))
}

// RegisterMdiosdoutr29Type MDIOS output data register 29
type RegisterMdiosdoutr29Type uint32

func (r *RegisterMdiosdoutr29Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr29Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr29Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr29Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr29Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr29FieldDout29Shift = 0
	RegisterMdiosdoutr29FieldDout29Mask  = 0xffff
)

// GetDout29 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr29Type) GetDout29() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr29FieldDout29Mask) >> RegisterMdiosdoutr29FieldDout29Shift)
}

// SetDout29 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr29Type) SetDout29(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr29FieldDout29Mask)|(uint32(value)<<RegisterMdiosdoutr29FieldDout29Shift))
}

// RegisterMdiosdoutr30Type MDIOS output data register 30
type RegisterMdiosdoutr30Type uint32

func (r *RegisterMdiosdoutr30Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr30Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr30Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr30Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr30Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr30FieldDout30Shift = 0
	RegisterMdiosdoutr30FieldDout30Mask  = 0xffff
)

// GetDout30 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr30Type) GetDout30() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr30FieldDout30Mask) >> RegisterMdiosdoutr30FieldDout30Shift)
}

// SetDout30 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr30Type) SetDout30(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr30FieldDout30Mask)|(uint32(value)<<RegisterMdiosdoutr30FieldDout30Shift))
}

// RegisterMdiosdoutr31Type MDIOS output data register 31
type RegisterMdiosdoutr31Type uint32

func (r *RegisterMdiosdoutr31Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdiosdoutr31Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdiosdoutr31Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdiosdoutr31Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdiosdoutr31Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdiosdoutr31FieldDout31Shift = 0
	RegisterMdiosdoutr31FieldDout31Mask  = 0xffff
)

// GetDout31 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr31Type) GetDout31() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdiosdoutr31FieldDout31Mask) >> RegisterMdiosdoutr31FieldDout31Shift)
}

// SetDout31 Output data sent to MDIO Master during read frames
func (r *RegisterMdiosdoutr31Type) SetDout31(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdiosdoutr31FieldDout31Mask)|(uint32(value)<<RegisterMdiosdoutr31FieldDout31Shift))
}
