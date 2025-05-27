//go:build arm && cortexm

package mpu

import (
	"unsafe"
	"volatile"
)

var (
	Mpu = (*_mpu)(unsafe.Pointer(uintptr(0xe000ed90)))
)

type _mpu struct {
	Type   registerTypeType
	Ctrl   registerCtrlType
	Rnr    registerRnrType
	Rbar   registerRbarType
	Rlar   registerRlarType
	Rbara1 registerRbara1Type
	Rbara2 registerRbara2Type
	Rbara3 registerRbara3Type
	Rlara1 registerRlara1Type
	Rlara2 registerRlara2Type
	Rlara3 registerRlara3Type
	_      [4]byte
	Mair0  registerMair0Type
	Mair1  registerMair1Type
}

// registerTypeType Indicates whether the MPU is present, and if so, how many regions it supports
type registerTypeType uint32

const (
	RegisterTypeFieldSeparateShift = 0
	RegisterTypeFieldSeparateMask  = 0x1
)

// GetSeparate Indicates support for unified or separate instructions and data address regions
func (r *registerTypeType) GetSeparate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTypeFieldSeparateMask) != 0
}

// SetSeparate Indicates support for unified or separate instructions and data address regions
func (r *registerTypeType) SetSeparate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTypeFieldSeparateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTypeFieldSeparateMask)
	}
}

const (
	RegisterTypeFieldDregionShift = 8
	RegisterTypeFieldDregionMask  = 0xff00
)

// GetDregion Number of regions supported by the MPU
func (r *registerTypeType) GetDregion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTypeFieldDregionMask) >> RegisterTypeFieldDregionShift)
}

// SetDregion Number of regions supported by the MPU
func (r *registerTypeType) SetDregion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTypeFieldDregionMask)|(uint32(value)<<RegisterTypeFieldDregionShift))
}

// registerCtrlType Enables the MPU
type registerCtrlType uint32

const (
	RegisterCtrlFieldEnableShift = 0
	RegisterCtrlFieldEnableMask  = 0x1
)

// GetEnable Enables the MPU
func (r *registerCtrlType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldEnableMask) != 0
}

// SetEnable Enables the MPU
func (r *registerCtrlType) SetEnable(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrlFieldEnableMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrlFieldEnableMask)
	}
}

const (
	RegisterCtrlFieldHfnmienaShift = 1
	RegisterCtrlFieldHfnmienaMask  = 0x2
)

// GetHfnmiena Enables the operation of MPU during HardFault and NMI handlers
func (r *registerCtrlType) GetHfnmiena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldHfnmienaMask) != 0
}

// SetHfnmiena Enables the operation of MPU during HardFault and NMI handlers
func (r *registerCtrlType) SetHfnmiena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrlFieldHfnmienaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrlFieldHfnmienaMask)
	}
}

const (
	RegisterCtrlFieldPrivdefenaShift = 2
	RegisterCtrlFieldPrivdefenaMask  = 0x4
)

// GetPrivdefena Enables privileged software access to the default memory map
func (r *registerCtrlType) GetPrivdefena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldPrivdefenaMask) != 0
}

// SetPrivdefena Enables privileged software access to the default memory map
func (r *registerCtrlType) SetPrivdefena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrlFieldPrivdefenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrlFieldPrivdefenaMask)
	}
}

// registerRnrType Selects the region currently accessed by MPU_RBAR and MPU_RLAR
type registerRnrType uint32

const (
	RegisterRnrFieldRegionShift = 0
	RegisterRnrFieldRegionMask  = 0xff
)

// GetRegion Indicates the memory region accessed by MPU_RBAR and PMU_RLAR
func (r *registerRnrType) GetRegion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRnrFieldRegionMask) >> RegisterRnrFieldRegionShift)
}

// SetRegion Indicates the memory region accessed by MPU_RBAR and PMU_RLAR
func (r *registerRnrType) SetRegion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRnrFieldRegionMask)|(uint32(value)<<RegisterRnrFieldRegionShift))
}

// registerRbarType Defines the base address of the MPU region selected by the MPU_RNR
type registerRbarType uint32

const (
	RegisterRbarFieldXnShift = 0
	RegisterRbarFieldXnMask  = 0x1
)

// GetXn Defines whether code can be executed from this region
func (r *registerRbarType) GetXn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRbarFieldXnMask) != 0
}

// SetXn Defines whether code can be executed from this region
func (r *registerRbarType) SetXn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRbarFieldXnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRbarFieldXnMask)
	}
}

const (
	RegisterRbarFieldApShift = 1
	RegisterRbarFieldApMask  = 0x6
)

// GetAp Access permissions
func (r *registerRbarType) GetAp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRbarFieldApMask) >> RegisterRbarFieldApShift)
}

// SetAp Access permissions
func (r *registerRbarType) SetAp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRbarFieldApMask)|(uint32(value)<<RegisterRbarFieldApShift))
}

const (
	RegisterRbarFieldShShift = 3
	RegisterRbarFieldShMask  = 0x18
)

// GetSh Defines the shareability domain of this region for Normal memory
func (r *registerRbarType) GetSh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRbarFieldShMask) >> RegisterRbarFieldShShift)
}

// SetSh Defines the shareability domain of this region for Normal memory
func (r *registerRbarType) SetSh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRbarFieldShMask)|(uint32(value)<<RegisterRbarFieldShShift))
}

const (
	RegisterRbarFieldBaseShift = 5
	RegisterRbarFieldBaseMask  = 0xffffffe0
)

// GetBase Contains bits[31:5] of the lower inclusive limit of the selected MPU memory region. This value is zero extended to provide the base address to be checked against.
func (r *registerRbarType) GetBase() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRbarFieldBaseMask) >> RegisterRbarFieldBaseShift)
}

// SetBase Contains bits[31:5] of the lower inclusive limit of the selected MPU memory region. This value is zero extended to provide the base address to be checked against.
func (r *registerRbarType) SetBase(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRbarFieldBaseMask)|(uint32(value)<<RegisterRbarFieldBaseShift))
}

// registerRlarType Defines the limit address of the MPU region selected by the MPU_RNR
type registerRlarType uint32

const (
	RegisterRlarFieldEnShift = 0
	RegisterRlarFieldEnMask  = 0x1
)

// GetEn Region enable
func (r *registerRlarType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRlarFieldEnMask) != 0
}

// SetEn Region enable
func (r *registerRlarType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRlarFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRlarFieldEnMask)
	}
}

const (
	RegisterRlarFieldAttrindxShift = 1
	RegisterRlarFieldAttrindxMask  = 0xe
)

// GetAttrindx Associates a set of attributes in the MPU_MAIR0 and MPU_MAIR1 fields
func (r *registerRlarType) GetAttrindx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRlarFieldAttrindxMask) >> RegisterRlarFieldAttrindxShift)
}

// SetAttrindx Associates a set of attributes in the MPU_MAIR0 and MPU_MAIR1 fields
func (r *registerRlarType) SetAttrindx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRlarFieldAttrindxMask)|(uint32(value)<<RegisterRlarFieldAttrindxShift))
}

const (
	RegisterRlarFieldPxnShift = 4
	RegisterRlarFieldPxnMask  = 0x10
)

// GetPxn Defines whether code can be executed from this privileged region
func (r *registerRlarType) GetPxn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRlarFieldPxnMask) != 0
}

// SetPxn Defines whether code can be executed from this privileged region
func (r *registerRlarType) SetPxn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRlarFieldPxnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRlarFieldPxnMask)
	}
}

const (
	RegisterRlarFieldLimitShift = 5
	RegisterRlarFieldLimitMask  = 0xffffffe0
)

// GetLimit Contains bits[31:5] of the upper inclusive limit of the selected MPU memory region. This value is postfixed with 0x1F to provide the limit address to be checked against.
func (r *registerRlarType) GetLimit() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRlarFieldLimitMask) >> RegisterRlarFieldLimitShift)
}

// SetLimit Contains bits[31:5] of the upper inclusive limit of the selected MPU memory region. This value is postfixed with 0x1F to provide the limit address to be checked against.
func (r *registerRlarType) SetLimit(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRlarFieldLimitMask)|(uint32(value)<<RegisterRlarFieldLimitShift))
}

// registerRbara1Type Provides indirect read and write access to the MPU base address register
type registerRbara1Type uint32

// registerRbara2Type Provides indirect read and write access to the MPU base address register
type registerRbara2Type uint32

// registerRbara3Type Provides indirect read and write access to the MPU base address register
type registerRbara3Type uint32

// registerRlara1Type Provides indirect read and write access to the MPU limit address register
type registerRlara1Type uint32

// registerRlara2Type Provides indirect read and write access to the MPU limit address register
type registerRlara2Type uint32

// registerRlara3Type Provides indirect read and write access to the MPU limit address register
type registerRlara3Type uint32

// registerMair0Type Provides the memory attribute encodings corresponding to the AttrIndex values
type registerMair0Type uint32

const (
	RegisterMair0FieldAttr0Shift = 0
	RegisterMair0FieldAttr0Mask  = 0xff
)

// GetAttr0 Memory attribute encoding for MPU regions with an AttrIndex of 0
func (r *registerMair0Type) GetAttr0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMair0FieldAttr0Mask) >> RegisterMair0FieldAttr0Shift)
}

// SetAttr0 Memory attribute encoding for MPU regions with an AttrIndex of 0
func (r *registerMair0Type) SetAttr0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMair0FieldAttr0Mask)|(uint32(value)<<RegisterMair0FieldAttr0Shift))
}

const (
	RegisterMair0FieldAttr1Shift = 8
	RegisterMair0FieldAttr1Mask  = 0xff00
)

// GetAttr1 Memory attribute encoding for MPU regions with an AttrIndex of 1
func (r *registerMair0Type) GetAttr1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMair0FieldAttr1Mask) >> RegisterMair0FieldAttr1Shift)
}

// SetAttr1 Memory attribute encoding for MPU regions with an AttrIndex of 1
func (r *registerMair0Type) SetAttr1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMair0FieldAttr1Mask)|(uint32(value)<<RegisterMair0FieldAttr1Shift))
}

const (
	RegisterMair0FieldAttr2Shift = 16
	RegisterMair0FieldAttr2Mask  = 0xff0000
)

// GetAttr2 Memory attribute encoding for MPU regions with an AttrIndex of 2
func (r *registerMair0Type) GetAttr2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMair0FieldAttr2Mask) >> RegisterMair0FieldAttr2Shift)
}

// SetAttr2 Memory attribute encoding for MPU regions with an AttrIndex of 2
func (r *registerMair0Type) SetAttr2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMair0FieldAttr2Mask)|(uint32(value)<<RegisterMair0FieldAttr2Shift))
}

const (
	RegisterMair0FieldAttr3Shift = 24
	RegisterMair0FieldAttr3Mask  = 0xff000000
)

// GetAttr3 Memory attribute encoding for MPU regions with an AttrIndex of 3
func (r *registerMair0Type) GetAttr3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMair0FieldAttr3Mask) >> RegisterMair0FieldAttr3Shift)
}

// SetAttr3 Memory attribute encoding for MPU regions with an AttrIndex of 3
func (r *registerMair0Type) SetAttr3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMair0FieldAttr3Mask)|(uint32(value)<<RegisterMair0FieldAttr3Shift))
}

// registerMair1Type Provides the memory attribute encodings corresponding to the AttrIndex values
type registerMair1Type uint32

const (
	RegisterMair1FieldAttr4Shift = 0
	RegisterMair1FieldAttr4Mask  = 0xff
)

// GetAttr4 Memory attribute encoding for MPU regions with an AttrIndex of 4
func (r *registerMair1Type) GetAttr4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMair1FieldAttr4Mask) >> RegisterMair1FieldAttr4Shift)
}

// SetAttr4 Memory attribute encoding for MPU regions with an AttrIndex of 4
func (r *registerMair1Type) SetAttr4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMair1FieldAttr4Mask)|(uint32(value)<<RegisterMair1FieldAttr4Shift))
}

const (
	RegisterMair1FieldAttr5Shift = 8
	RegisterMair1FieldAttr5Mask  = 0xff00
)

// GetAttr5 Memory attribute encoding for MPU regions with an AttrIndex of 5
func (r *registerMair1Type) GetAttr5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMair1FieldAttr5Mask) >> RegisterMair1FieldAttr5Shift)
}

// SetAttr5 Memory attribute encoding for MPU regions with an AttrIndex of 5
func (r *registerMair1Type) SetAttr5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMair1FieldAttr5Mask)|(uint32(value)<<RegisterMair1FieldAttr5Shift))
}

const (
	RegisterMair1FieldAttr6Shift = 16
	RegisterMair1FieldAttr6Mask  = 0xff0000
)

// GetAttr6 Memory attribute encoding for MPU regions with an AttrIndex of 6
func (r *registerMair1Type) GetAttr6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMair1FieldAttr6Mask) >> RegisterMair1FieldAttr6Shift)
}

// SetAttr6 Memory attribute encoding for MPU regions with an AttrIndex of 6
func (r *registerMair1Type) SetAttr6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMair1FieldAttr6Mask)|(uint32(value)<<RegisterMair1FieldAttr6Shift))
}

const (
	RegisterMair1FieldAttr7Shift = 24
	RegisterMair1FieldAttr7Mask  = 0xff000000
)

// GetAttr7 Memory attribute encoding for MPU regions with an AttrIndex of 7
func (r *registerMair1Type) GetAttr7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMair1FieldAttr7Mask) >> RegisterMair1FieldAttr7Shift)
}

// SetAttr7 Memory attribute encoding for MPU regions with an AttrIndex of 7
func (r *registerMair1Type) SetAttr7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMair1FieldAttr7Mask)|(uint32(value)<<RegisterMair1FieldAttr7Shift))
}
