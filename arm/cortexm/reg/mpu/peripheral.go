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
	Type   RegisterTypeType
	Ctrl   RegisterCtrlType
	Rnr    RegisterRnrType
	Rbar   RegisterRbarType
	Rasr   RegisterRasrType
	Rbara1 RegisterRbara1Type
	Rasra1 RegisterRasra1Type
	Rbara2 RegisterRbara2Type
	Rasra2 RegisterRasra2Type
	Rbara3 RegisterRbara3Type
	Rasra3 RegisterRasra3Type
}

// RegisterTypeType Indicates whether the MPU is present, and how many regions it supports
type RegisterTypeType uint32

func (r *RegisterTypeType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTypeType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTypeType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTypeType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTypeType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTypeFieldSeparateShift = 0
	RegisterTypeFieldSeparateMask  = 0x1
)

// GetSeparate Indicates support for unified or separate instruction and data address regions
func (r *RegisterTypeType) GetSeparate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTypeFieldSeparateMask) != 0
}

const (
	RegisterTypeFieldDregionShift = 8
	RegisterTypeFieldDregionMask  = 0xff00
)

// GetDregion Number of supported MPU data regions
func (r *RegisterTypeType) GetDregion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTypeFieldDregionMask) >> RegisterTypeFieldDregionShift)
}

// RegisterCtrlType MPU control register
type RegisterCtrlType uint32

func (r *RegisterCtrlType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCtrlType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCtrlType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCtrlType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCtrlType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCtrlFieldEnableShift = 0
	RegisterCtrlFieldEnableMask  = 0x1
)

// GetEnable Enables the MPU
func (r *RegisterCtrlType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldEnableMask) != 0
}

// SetEnable Enables the MPU
func (r *RegisterCtrlType) SetEnable(value bool) {
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

// GetHfnmiena Enables MPU operation during HardFault and NMI handlers
func (r *RegisterCtrlType) GetHfnmiena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldHfnmienaMask) != 0
}

// SetHfnmiena Enables MPU operation during HardFault and NMI handlers
func (r *RegisterCtrlType) SetHfnmiena(value bool) {
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
func (r *RegisterCtrlType) GetPrivdefena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldPrivdefenaMask) != 0
}

// SetPrivdefena Enables privileged software access to the default memory map
func (r *RegisterCtrlType) SetPrivdefena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrlFieldPrivdefenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrlFieldPrivdefenaMask)
	}
}

// RegisterRnrType MPU region number register
type RegisterRnrType uint32

func (r *RegisterRnrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRnrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRnrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRnrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRnrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRnrFieldRegionShift = 0
	RegisterRnrFieldRegionMask  = 0xff
)

// GetRegion Selects the MPU region accessed by MPU_RBAR and MPU_RASR
func (r *RegisterRnrType) GetRegion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRnrFieldRegionMask) >> RegisterRnrFieldRegionShift)
}

// SetRegion Selects the MPU region accessed by MPU_RBAR and MPU_RASR
func (r *RegisterRnrType) SetRegion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRnrFieldRegionMask)|(uint32(value)<<RegisterRnrFieldRegionShift))
}

// RegisterRbarType MPU region base address register
type RegisterRbarType uint32

func (r *RegisterRbarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRbarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRbarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRbarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRbarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRbarFieldRegionShift = 0
	RegisterRbarFieldRegionMask  = 0xf
)

// GetRegion MPU region number field used when VALID is set
func (r *RegisterRbarType) GetRegion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRbarFieldRegionMask) >> RegisterRbarFieldRegionShift)
}

// SetRegion MPU region number field used when VALID is set
func (r *RegisterRbarType) SetRegion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRbarFieldRegionMask)|(uint32(value)<<RegisterRbarFieldRegionShift))
}

const (
	RegisterRbarFieldValidShift = 4
	RegisterRbarFieldValidMask  = 0x10
)

// GetValid Updates MPU_RNR with REGION when this bit is set
func (r *RegisterRbarType) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRbarFieldValidMask) != 0
}

// SetValid Updates MPU_RNR with REGION when this bit is set
func (r *RegisterRbarType) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRbarFieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRbarFieldValidMask)
	}
}

const (
	RegisterRbarFieldAddrShift = 5
	RegisterRbarFieldAddrMask  = 0xffffffe0
)

// GetAddr Region base address bits[31:5]
func (r *RegisterRbarType) GetAddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRbarFieldAddrMask) >> RegisterRbarFieldAddrShift)
}

// SetAddr Region base address bits[31:5]
func (r *RegisterRbarType) SetAddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRbarFieldAddrMask)|(uint32(value)<<RegisterRbarFieldAddrShift))
}

// RegisterRasrType MPU region attribute and size register
type RegisterRasrType uint32

func (r *RegisterRasrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRasrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRasrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRasrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRasrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRasrFieldEnableShift = 0
	RegisterRasrFieldEnableMask  = 0x1
)

// GetEnable Enables the selected MPU region
func (r *RegisterRasrType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasrFieldEnableMask) != 0
}

// SetEnable Enables the selected MPU region
func (r *RegisterRasrType) SetEnable(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasrFieldEnableMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasrFieldEnableMask)
	}
}

const (
	RegisterRasrFieldSizeShift = 1
	RegisterRasrFieldSizeMask  = 0x3e
)

// GetSize Region size encoding; region size is 2^(SIZE + 1)
func (r *RegisterRasrType) GetSize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasrFieldSizeMask) >> RegisterRasrFieldSizeShift)
}

// SetSize Region size encoding; region size is 2^(SIZE + 1)
func (r *RegisterRasrType) SetSize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasrFieldSizeMask)|(uint32(value)<<RegisterRasrFieldSizeShift))
}

const (
	RegisterRasrFieldSrdShift = 8
	RegisterRasrFieldSrdMask  = 0xff00
)

// GetSrd Subregion disable bits
func (r *RegisterRasrType) GetSrd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasrFieldSrdMask) >> RegisterRasrFieldSrdShift)
}

// SetSrd Subregion disable bits
func (r *RegisterRasrType) SetSrd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasrFieldSrdMask)|(uint32(value)<<RegisterRasrFieldSrdShift))
}

const (
	RegisterRasrFieldBShift = 16
	RegisterRasrFieldBMask  = 0x10000
)

// GetB Bufferable memory attribute
func (r *RegisterRasrType) GetB() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasrFieldBMask) != 0
}

// SetB Bufferable memory attribute
func (r *RegisterRasrType) SetB(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasrFieldBMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasrFieldBMask)
	}
}

const (
	RegisterRasrFieldCShift = 17
	RegisterRasrFieldCMask  = 0x20000
)

// GetC Cacheable memory attribute
func (r *RegisterRasrType) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasrFieldCMask) != 0
}

// SetC Cacheable memory attribute
func (r *RegisterRasrType) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasrFieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasrFieldCMask)
	}
}

const (
	RegisterRasrFieldSShift = 18
	RegisterRasrFieldSMask  = 0x40000
)

// GetS Shareable memory attribute
func (r *RegisterRasrType) GetS() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasrFieldSMask) != 0
}

// SetS Shareable memory attribute
func (r *RegisterRasrType) SetS(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasrFieldSMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasrFieldSMask)
	}
}

const (
	RegisterRasrFieldTexShift = 19
	RegisterRasrFieldTexMask  = 0x380000
)

// GetTex Type extension memory attribute
func (r *RegisterRasrType) GetTex() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasrFieldTexMask) >> RegisterRasrFieldTexShift)
}

// SetTex Type extension memory attribute
func (r *RegisterRasrType) SetTex(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasrFieldTexMask)|(uint32(value)<<RegisterRasrFieldTexShift))
}

const (
	RegisterRasrFieldApShift = 24
	RegisterRasrFieldApMask  = 0x7000000
)

// GetAp Access permission field
func (r *RegisterRasrType) GetAp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasrFieldApMask) >> RegisterRasrFieldApShift)
}

// SetAp Access permission field
func (r *RegisterRasrType) SetAp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasrFieldApMask)|(uint32(value)<<RegisterRasrFieldApShift))
}

const (
	RegisterRasrFieldXnShift = 28
	RegisterRasrFieldXnMask  = 0x10000000
)

// GetXn Execute-never attribute
func (r *RegisterRasrType) GetXn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasrFieldXnMask) != 0
}

// SetXn Execute-never attribute
func (r *RegisterRasrType) SetXn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasrFieldXnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasrFieldXnMask)
	}
}

// RegisterRbara1Type MPU region base address alias register 1
type RegisterRbara1Type uint32

func (r *RegisterRbara1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRbara1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRbara1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRbara1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRbara1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRbara1FieldRegionShift = 0
	RegisterRbara1FieldRegionMask  = 0xf
)

// GetRegion MPU region number field used when VALID is set
func (r *RegisterRbara1Type) GetRegion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRbara1FieldRegionMask) >> RegisterRbara1FieldRegionShift)
}

// SetRegion MPU region number field used when VALID is set
func (r *RegisterRbara1Type) SetRegion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRbara1FieldRegionMask)|(uint32(value)<<RegisterRbara1FieldRegionShift))
}

const (
	RegisterRbara1FieldValidShift = 4
	RegisterRbara1FieldValidMask  = 0x10
)

// GetValid Updates MPU_RNR with REGION when this bit is set
func (r *RegisterRbara1Type) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRbara1FieldValidMask) != 0
}

// SetValid Updates MPU_RNR with REGION when this bit is set
func (r *RegisterRbara1Type) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRbara1FieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRbara1FieldValidMask)
	}
}

const (
	RegisterRbara1FieldAddrShift = 5
	RegisterRbara1FieldAddrMask  = 0xffffffe0
)

// GetAddr Region base address bits[31:5]
func (r *RegisterRbara1Type) GetAddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRbara1FieldAddrMask) >> RegisterRbara1FieldAddrShift)
}

// SetAddr Region base address bits[31:5]
func (r *RegisterRbara1Type) SetAddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRbara1FieldAddrMask)|(uint32(value)<<RegisterRbara1FieldAddrShift))
}

// RegisterRasra1Type MPU region attribute and size alias register 1
type RegisterRasra1Type uint32

func (r *RegisterRasra1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRasra1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRasra1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRasra1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRasra1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRasra1FieldEnableShift = 0
	RegisterRasra1FieldEnableMask  = 0x1
)

// GetEnable Enables the selected MPU region
func (r *RegisterRasra1Type) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra1FieldEnableMask) != 0
}

// SetEnable Enables the selected MPU region
func (r *RegisterRasra1Type) SetEnable(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra1FieldEnableMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra1FieldEnableMask)
	}
}

const (
	RegisterRasra1FieldSizeShift = 1
	RegisterRasra1FieldSizeMask  = 0x3e
)

// GetSize Region size encoding; region size is 2^(SIZE + 1)
func (r *RegisterRasra1Type) GetSize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasra1FieldSizeMask) >> RegisterRasra1FieldSizeShift)
}

// SetSize Region size encoding; region size is 2^(SIZE + 1)
func (r *RegisterRasra1Type) SetSize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasra1FieldSizeMask)|(uint32(value)<<RegisterRasra1FieldSizeShift))
}

const (
	RegisterRasra1FieldSrdShift = 8
	RegisterRasra1FieldSrdMask  = 0xff00
)

// GetSrd Subregion disable bits
func (r *RegisterRasra1Type) GetSrd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasra1FieldSrdMask) >> RegisterRasra1FieldSrdShift)
}

// SetSrd Subregion disable bits
func (r *RegisterRasra1Type) SetSrd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasra1FieldSrdMask)|(uint32(value)<<RegisterRasra1FieldSrdShift))
}

const (
	RegisterRasra1FieldBShift = 16
	RegisterRasra1FieldBMask  = 0x10000
)

// GetB Bufferable memory attribute
func (r *RegisterRasra1Type) GetB() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra1FieldBMask) != 0
}

// SetB Bufferable memory attribute
func (r *RegisterRasra1Type) SetB(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra1FieldBMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra1FieldBMask)
	}
}

const (
	RegisterRasra1FieldCShift = 17
	RegisterRasra1FieldCMask  = 0x20000
)

// GetC Cacheable memory attribute
func (r *RegisterRasra1Type) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra1FieldCMask) != 0
}

// SetC Cacheable memory attribute
func (r *RegisterRasra1Type) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra1FieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra1FieldCMask)
	}
}

const (
	RegisterRasra1FieldSShift = 18
	RegisterRasra1FieldSMask  = 0x40000
)

// GetS Shareable memory attribute
func (r *RegisterRasra1Type) GetS() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra1FieldSMask) != 0
}

// SetS Shareable memory attribute
func (r *RegisterRasra1Type) SetS(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra1FieldSMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra1FieldSMask)
	}
}

const (
	RegisterRasra1FieldTexShift = 19
	RegisterRasra1FieldTexMask  = 0x380000
)

// GetTex Type extension memory attribute
func (r *RegisterRasra1Type) GetTex() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasra1FieldTexMask) >> RegisterRasra1FieldTexShift)
}

// SetTex Type extension memory attribute
func (r *RegisterRasra1Type) SetTex(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasra1FieldTexMask)|(uint32(value)<<RegisterRasra1FieldTexShift))
}

const (
	RegisterRasra1FieldApShift = 24
	RegisterRasra1FieldApMask  = 0x7000000
)

// GetAp Access permission field
func (r *RegisterRasra1Type) GetAp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasra1FieldApMask) >> RegisterRasra1FieldApShift)
}

// SetAp Access permission field
func (r *RegisterRasra1Type) SetAp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasra1FieldApMask)|(uint32(value)<<RegisterRasra1FieldApShift))
}

const (
	RegisterRasra1FieldXnShift = 28
	RegisterRasra1FieldXnMask  = 0x10000000
)

// GetXn Execute-never attribute
func (r *RegisterRasra1Type) GetXn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra1FieldXnMask) != 0
}

// SetXn Execute-never attribute
func (r *RegisterRasra1Type) SetXn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra1FieldXnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra1FieldXnMask)
	}
}

// RegisterRbara2Type MPU region base address alias register 2
type RegisterRbara2Type uint32

func (r *RegisterRbara2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRbara2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRbara2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRbara2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRbara2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRbara2FieldRegionShift = 0
	RegisterRbara2FieldRegionMask  = 0xf
)

// GetRegion MPU region number field used when VALID is set
func (r *RegisterRbara2Type) GetRegion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRbara2FieldRegionMask) >> RegisterRbara2FieldRegionShift)
}

// SetRegion MPU region number field used when VALID is set
func (r *RegisterRbara2Type) SetRegion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRbara2FieldRegionMask)|(uint32(value)<<RegisterRbara2FieldRegionShift))
}

const (
	RegisterRbara2FieldValidShift = 4
	RegisterRbara2FieldValidMask  = 0x10
)

// GetValid Updates MPU_RNR with REGION when this bit is set
func (r *RegisterRbara2Type) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRbara2FieldValidMask) != 0
}

// SetValid Updates MPU_RNR with REGION when this bit is set
func (r *RegisterRbara2Type) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRbara2FieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRbara2FieldValidMask)
	}
}

const (
	RegisterRbara2FieldAddrShift = 5
	RegisterRbara2FieldAddrMask  = 0xffffffe0
)

// GetAddr Region base address bits[31:5]
func (r *RegisterRbara2Type) GetAddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRbara2FieldAddrMask) >> RegisterRbara2FieldAddrShift)
}

// SetAddr Region base address bits[31:5]
func (r *RegisterRbara2Type) SetAddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRbara2FieldAddrMask)|(uint32(value)<<RegisterRbara2FieldAddrShift))
}

// RegisterRasra2Type MPU region attribute and size alias register 2
type RegisterRasra2Type uint32

func (r *RegisterRasra2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRasra2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRasra2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRasra2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRasra2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRasra2FieldEnableShift = 0
	RegisterRasra2FieldEnableMask  = 0x1
)

// GetEnable Enables the selected MPU region
func (r *RegisterRasra2Type) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra2FieldEnableMask) != 0
}

// SetEnable Enables the selected MPU region
func (r *RegisterRasra2Type) SetEnable(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra2FieldEnableMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra2FieldEnableMask)
	}
}

const (
	RegisterRasra2FieldSizeShift = 1
	RegisterRasra2FieldSizeMask  = 0x3e
)

// GetSize Region size encoding; region size is 2^(SIZE + 1)
func (r *RegisterRasra2Type) GetSize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasra2FieldSizeMask) >> RegisterRasra2FieldSizeShift)
}

// SetSize Region size encoding; region size is 2^(SIZE + 1)
func (r *RegisterRasra2Type) SetSize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasra2FieldSizeMask)|(uint32(value)<<RegisterRasra2FieldSizeShift))
}

const (
	RegisterRasra2FieldSrdShift = 8
	RegisterRasra2FieldSrdMask  = 0xff00
)

// GetSrd Subregion disable bits
func (r *RegisterRasra2Type) GetSrd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasra2FieldSrdMask) >> RegisterRasra2FieldSrdShift)
}

// SetSrd Subregion disable bits
func (r *RegisterRasra2Type) SetSrd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasra2FieldSrdMask)|(uint32(value)<<RegisterRasra2FieldSrdShift))
}

const (
	RegisterRasra2FieldBShift = 16
	RegisterRasra2FieldBMask  = 0x10000
)

// GetB Bufferable memory attribute
func (r *RegisterRasra2Type) GetB() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra2FieldBMask) != 0
}

// SetB Bufferable memory attribute
func (r *RegisterRasra2Type) SetB(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra2FieldBMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra2FieldBMask)
	}
}

const (
	RegisterRasra2FieldCShift = 17
	RegisterRasra2FieldCMask  = 0x20000
)

// GetC Cacheable memory attribute
func (r *RegisterRasra2Type) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra2FieldCMask) != 0
}

// SetC Cacheable memory attribute
func (r *RegisterRasra2Type) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra2FieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra2FieldCMask)
	}
}

const (
	RegisterRasra2FieldSShift = 18
	RegisterRasra2FieldSMask  = 0x40000
)

// GetS Shareable memory attribute
func (r *RegisterRasra2Type) GetS() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra2FieldSMask) != 0
}

// SetS Shareable memory attribute
func (r *RegisterRasra2Type) SetS(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra2FieldSMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra2FieldSMask)
	}
}

const (
	RegisterRasra2FieldTexShift = 19
	RegisterRasra2FieldTexMask  = 0x380000
)

// GetTex Type extension memory attribute
func (r *RegisterRasra2Type) GetTex() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasra2FieldTexMask) >> RegisterRasra2FieldTexShift)
}

// SetTex Type extension memory attribute
func (r *RegisterRasra2Type) SetTex(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasra2FieldTexMask)|(uint32(value)<<RegisterRasra2FieldTexShift))
}

const (
	RegisterRasra2FieldApShift = 24
	RegisterRasra2FieldApMask  = 0x7000000
)

// GetAp Access permission field
func (r *RegisterRasra2Type) GetAp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasra2FieldApMask) >> RegisterRasra2FieldApShift)
}

// SetAp Access permission field
func (r *RegisterRasra2Type) SetAp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasra2FieldApMask)|(uint32(value)<<RegisterRasra2FieldApShift))
}

const (
	RegisterRasra2FieldXnShift = 28
	RegisterRasra2FieldXnMask  = 0x10000000
)

// GetXn Execute-never attribute
func (r *RegisterRasra2Type) GetXn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra2FieldXnMask) != 0
}

// SetXn Execute-never attribute
func (r *RegisterRasra2Type) SetXn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra2FieldXnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra2FieldXnMask)
	}
}

// RegisterRbara3Type MPU region base address alias register 3
type RegisterRbara3Type uint32

func (r *RegisterRbara3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRbara3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRbara3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRbara3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRbara3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRbara3FieldRegionShift = 0
	RegisterRbara3FieldRegionMask  = 0xf
)

// GetRegion MPU region number field used when VALID is set
func (r *RegisterRbara3Type) GetRegion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRbara3FieldRegionMask) >> RegisterRbara3FieldRegionShift)
}

// SetRegion MPU region number field used when VALID is set
func (r *RegisterRbara3Type) SetRegion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRbara3FieldRegionMask)|(uint32(value)<<RegisterRbara3FieldRegionShift))
}

const (
	RegisterRbara3FieldValidShift = 4
	RegisterRbara3FieldValidMask  = 0x10
)

// GetValid Updates MPU_RNR with REGION when this bit is set
func (r *RegisterRbara3Type) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRbara3FieldValidMask) != 0
}

// SetValid Updates MPU_RNR with REGION when this bit is set
func (r *RegisterRbara3Type) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRbara3FieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRbara3FieldValidMask)
	}
}

const (
	RegisterRbara3FieldAddrShift = 5
	RegisterRbara3FieldAddrMask  = 0xffffffe0
)

// GetAddr Region base address bits[31:5]
func (r *RegisterRbara3Type) GetAddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRbara3FieldAddrMask) >> RegisterRbara3FieldAddrShift)
}

// SetAddr Region base address bits[31:5]
func (r *RegisterRbara3Type) SetAddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRbara3FieldAddrMask)|(uint32(value)<<RegisterRbara3FieldAddrShift))
}

// RegisterRasra3Type MPU region attribute and size alias register 3
type RegisterRasra3Type uint32

func (r *RegisterRasra3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRasra3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRasra3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRasra3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRasra3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRasra3FieldEnableShift = 0
	RegisterRasra3FieldEnableMask  = 0x1
)

// GetEnable Enables the selected MPU region
func (r *RegisterRasra3Type) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra3FieldEnableMask) != 0
}

// SetEnable Enables the selected MPU region
func (r *RegisterRasra3Type) SetEnable(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra3FieldEnableMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra3FieldEnableMask)
	}
}

const (
	RegisterRasra3FieldSizeShift = 1
	RegisterRasra3FieldSizeMask  = 0x3e
)

// GetSize Region size encoding; region size is 2^(SIZE + 1)
func (r *RegisterRasra3Type) GetSize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasra3FieldSizeMask) >> RegisterRasra3FieldSizeShift)
}

// SetSize Region size encoding; region size is 2^(SIZE + 1)
func (r *RegisterRasra3Type) SetSize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasra3FieldSizeMask)|(uint32(value)<<RegisterRasra3FieldSizeShift))
}

const (
	RegisterRasra3FieldSrdShift = 8
	RegisterRasra3FieldSrdMask  = 0xff00
)

// GetSrd Subregion disable bits
func (r *RegisterRasra3Type) GetSrd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasra3FieldSrdMask) >> RegisterRasra3FieldSrdShift)
}

// SetSrd Subregion disable bits
func (r *RegisterRasra3Type) SetSrd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasra3FieldSrdMask)|(uint32(value)<<RegisterRasra3FieldSrdShift))
}

const (
	RegisterRasra3FieldBShift = 16
	RegisterRasra3FieldBMask  = 0x10000
)

// GetB Bufferable memory attribute
func (r *RegisterRasra3Type) GetB() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra3FieldBMask) != 0
}

// SetB Bufferable memory attribute
func (r *RegisterRasra3Type) SetB(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra3FieldBMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra3FieldBMask)
	}
}

const (
	RegisterRasra3FieldCShift = 17
	RegisterRasra3FieldCMask  = 0x20000
)

// GetC Cacheable memory attribute
func (r *RegisterRasra3Type) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra3FieldCMask) != 0
}

// SetC Cacheable memory attribute
func (r *RegisterRasra3Type) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra3FieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra3FieldCMask)
	}
}

const (
	RegisterRasra3FieldSShift = 18
	RegisterRasra3FieldSMask  = 0x40000
)

// GetS Shareable memory attribute
func (r *RegisterRasra3Type) GetS() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra3FieldSMask) != 0
}

// SetS Shareable memory attribute
func (r *RegisterRasra3Type) SetS(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra3FieldSMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra3FieldSMask)
	}
}

const (
	RegisterRasra3FieldTexShift = 19
	RegisterRasra3FieldTexMask  = 0x380000
)

// GetTex Type extension memory attribute
func (r *RegisterRasra3Type) GetTex() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasra3FieldTexMask) >> RegisterRasra3FieldTexShift)
}

// SetTex Type extension memory attribute
func (r *RegisterRasra3Type) SetTex(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasra3FieldTexMask)|(uint32(value)<<RegisterRasra3FieldTexShift))
}

const (
	RegisterRasra3FieldApShift = 24
	RegisterRasra3FieldApMask  = 0x7000000
)

// GetAp Access permission field
func (r *RegisterRasra3Type) GetAp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRasra3FieldApMask) >> RegisterRasra3FieldApShift)
}

// SetAp Access permission field
func (r *RegisterRasra3Type) SetAp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRasra3FieldApMask)|(uint32(value)<<RegisterRasra3FieldApShift))
}

const (
	RegisterRasra3FieldXnShift = 28
	RegisterRasra3FieldXnMask  = 0x10000000
)

// GetXn Execute-never attribute
func (r *RegisterRasra3Type) GetXn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRasra3FieldXnMask) != 0
}

// SetXn Execute-never attribute
func (r *RegisterRasra3Type) SetXn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRasra3FieldXnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRasra3FieldXnMask)
	}
}
