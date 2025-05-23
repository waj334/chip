package mpu

import (
	"unsafe"
	"volatile"
)

var (
	Mpu = (*_mpu)(unsafe.Pointer(uintptr(0xe000ed90)))
)

type _mpu struct {
	Mpu_typer registerMpu_typerType
	Mpu_ctrl  registerMpu_ctrlType
	Mpu_rnr   registerMpu_rnrType
	Mpu_rbar  registerMpu_rbarType
	Mpu_rasr  registerMpu_rasrType
}

// registerMpu_typerType MPU type register
type registerMpu_typerType uint32

const (
	RegisterMpu_typerFieldSeparateShift = 0
	RegisterMpu_typerFieldSeparateMask  = 0x1
)

// GetSeparate Separate flag
func (r *registerMpu_typerType) GetSeparate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMpu_typerFieldSeparateMask) != 0
}

// SetSeparate Separate flag
func (r *registerMpu_typerType) SetSeparate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMpu_typerFieldSeparateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMpu_typerFieldSeparateMask)
	}
}

const (
	RegisterMpu_typerFieldDregionShift = 8
	RegisterMpu_typerFieldDregionMask  = 0xff00
)

// GetDregion Number of MPU data regions
func (r *registerMpu_typerType) GetDregion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMpu_typerFieldDregionMask) >> RegisterMpu_typerFieldDregionShift)
}

// SetDregion Number of MPU data regions
func (r *registerMpu_typerType) SetDregion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMpu_typerFieldDregionMask)|(uint32(value)<<RegisterMpu_typerFieldDregionShift))
}

const (
	RegisterMpu_typerFieldIregionShift = 16
	RegisterMpu_typerFieldIregionMask  = 0xff0000
)

// GetIregion Number of MPU instruction regions
func (r *registerMpu_typerType) GetIregion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMpu_typerFieldIregionMask) >> RegisterMpu_typerFieldIregionShift)
}

// SetIregion Number of MPU instruction regions
func (r *registerMpu_typerType) SetIregion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMpu_typerFieldIregionMask)|(uint32(value)<<RegisterMpu_typerFieldIregionShift))
}

// registerMpu_ctrlType MPU control register
type registerMpu_ctrlType uint32

const (
	RegisterMpu_ctrlFieldEnableShift = 0
	RegisterMpu_ctrlFieldEnableMask  = 0x1
)

// GetEnable Enables the MPU
func (r *registerMpu_ctrlType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMpu_ctrlFieldEnableMask) != 0
}

// SetEnable Enables the MPU
func (r *registerMpu_ctrlType) SetEnable(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMpu_ctrlFieldEnableMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMpu_ctrlFieldEnableMask)
	}
}

const (
	RegisterMpu_ctrlFieldHfnmienaShift = 1
	RegisterMpu_ctrlFieldHfnmienaMask  = 0x2
)

// GetHfnmiena Enables the operation of MPU during hard fault
func (r *registerMpu_ctrlType) GetHfnmiena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMpu_ctrlFieldHfnmienaMask) != 0
}

// SetHfnmiena Enables the operation of MPU during hard fault
func (r *registerMpu_ctrlType) SetHfnmiena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMpu_ctrlFieldHfnmienaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMpu_ctrlFieldHfnmienaMask)
	}
}

const (
	RegisterMpu_ctrlFieldPrivdefenaShift = 2
	RegisterMpu_ctrlFieldPrivdefenaMask  = 0x4
)

// GetPrivdefena Enable priviliged software access to default memory map
func (r *registerMpu_ctrlType) GetPrivdefena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMpu_ctrlFieldPrivdefenaMask) != 0
}

// SetPrivdefena Enable priviliged software access to default memory map
func (r *registerMpu_ctrlType) SetPrivdefena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMpu_ctrlFieldPrivdefenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMpu_ctrlFieldPrivdefenaMask)
	}
}

// registerMpu_rnrType MPU region number register
type registerMpu_rnrType uint32

const (
	RegisterMpu_rnrFieldRegionShift = 0
	RegisterMpu_rnrFieldRegionMask  = 0xff
)

// GetRegion MPU region
func (r *registerMpu_rnrType) GetRegion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rnrFieldRegionMask) >> RegisterMpu_rnrFieldRegionShift)
}

// SetRegion MPU region
func (r *registerMpu_rnrType) SetRegion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rnrFieldRegionMask)|(uint32(value)<<RegisterMpu_rnrFieldRegionShift))
}

// registerMpu_rbarType MPU region base address register
type registerMpu_rbarType uint32

const (
	RegisterMpu_rbarFieldRegionShift = 0
	RegisterMpu_rbarFieldRegionMask  = 0xf
)

// GetRegion MPU region field
func (r *registerMpu_rbarType) GetRegion() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rbarFieldRegionMask) >> RegisterMpu_rbarFieldRegionShift)
}

// SetRegion MPU region field
func (r *registerMpu_rbarType) SetRegion(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rbarFieldRegionMask)|(uint32(value)<<RegisterMpu_rbarFieldRegionShift))
}

const (
	RegisterMpu_rbarFieldValidShift = 4
	RegisterMpu_rbarFieldValidMask  = 0x10
)

// GetValid MPU region number valid
func (r *registerMpu_rbarType) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rbarFieldValidMask) != 0
}

// SetValid MPU region number valid
func (r *registerMpu_rbarType) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMpu_rbarFieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rbarFieldValidMask)
	}
}

const (
	RegisterMpu_rbarFieldAddrShift = 5
	RegisterMpu_rbarFieldAddrMask  = 0xffffffe0
)

// GetAddr Region base address field
func (r *registerMpu_rbarType) GetAddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rbarFieldAddrMask) >> RegisterMpu_rbarFieldAddrShift)
}

// SetAddr Region base address field
func (r *registerMpu_rbarType) SetAddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rbarFieldAddrMask)|(uint32(value)<<RegisterMpu_rbarFieldAddrShift))
}

// registerMpu_rasrType MPU region attribute and size register
type registerMpu_rasrType uint32

const (
	RegisterMpu_rasrFieldEnableShift = 0
	RegisterMpu_rasrFieldEnableMask  = 0x1
)

// GetEnable Region enable bit.
func (r *registerMpu_rasrType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rasrFieldEnableMask) != 0
}

// SetEnable Region enable bit.
func (r *registerMpu_rasrType) SetEnable(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMpu_rasrFieldEnableMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rasrFieldEnableMask)
	}
}

const (
	RegisterMpu_rasrFieldSizeShift = 1
	RegisterMpu_rasrFieldSizeMask  = 0x3e
)

// GetSize Size of the MPU protection region
func (r *registerMpu_rasrType) GetSize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rasrFieldSizeMask) >> RegisterMpu_rasrFieldSizeShift)
}

// SetSize Size of the MPU protection region
func (r *registerMpu_rasrType) SetSize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rasrFieldSizeMask)|(uint32(value)<<RegisterMpu_rasrFieldSizeShift))
}

const (
	RegisterMpu_rasrFieldSrdShift = 8
	RegisterMpu_rasrFieldSrdMask  = 0xff00
)

// GetSrd Subregion disable bits
func (r *registerMpu_rasrType) GetSrd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rasrFieldSrdMask) >> RegisterMpu_rasrFieldSrdShift)
}

// SetSrd Subregion disable bits
func (r *registerMpu_rasrType) SetSrd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rasrFieldSrdMask)|(uint32(value)<<RegisterMpu_rasrFieldSrdShift))
}

const (
	RegisterMpu_rasrFieldBShift = 16
	RegisterMpu_rasrFieldBMask  = 0x10000
)

// GetB memory attribute
func (r *registerMpu_rasrType) GetB() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rasrFieldBMask) != 0
}

// SetB memory attribute
func (r *registerMpu_rasrType) SetB(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMpu_rasrFieldBMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rasrFieldBMask)
	}
}

const (
	RegisterMpu_rasrFieldCShift = 17
	RegisterMpu_rasrFieldCMask  = 0x20000
)

// GetC memory attribute
func (r *registerMpu_rasrType) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rasrFieldCMask) != 0
}

// SetC memory attribute
func (r *registerMpu_rasrType) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMpu_rasrFieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rasrFieldCMask)
	}
}

const (
	RegisterMpu_rasrFieldSShift = 18
	RegisterMpu_rasrFieldSMask  = 0x40000
)

// GetS Shareable memory attribute
func (r *registerMpu_rasrType) GetS() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rasrFieldSMask) != 0
}

// SetS Shareable memory attribute
func (r *registerMpu_rasrType) SetS(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMpu_rasrFieldSMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rasrFieldSMask)
	}
}

const (
	RegisterMpu_rasrFieldTexShift = 19
	RegisterMpu_rasrFieldTexMask  = 0x380000
)

// GetTex memory attribute
func (r *registerMpu_rasrType) GetTex() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rasrFieldTexMask) >> RegisterMpu_rasrFieldTexShift)
}

// SetTex memory attribute
func (r *registerMpu_rasrType) SetTex(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rasrFieldTexMask)|(uint32(value)<<RegisterMpu_rasrFieldTexShift))
}

const (
	RegisterMpu_rasrFieldApShift = 24
	RegisterMpu_rasrFieldApMask  = 0x7000000
)

// GetAp Access permission
func (r *registerMpu_rasrType) GetAp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rasrFieldApMask) >> RegisterMpu_rasrFieldApShift)
}

// SetAp Access permission
func (r *registerMpu_rasrType) SetAp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rasrFieldApMask)|(uint32(value)<<RegisterMpu_rasrFieldApShift))
}

const (
	RegisterMpu_rasrFieldXnShift = 28
	RegisterMpu_rasrFieldXnMask  = 0x10000000
)

// GetXn Instruction access disable bit
func (r *registerMpu_rasrType) GetXn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMpu_rasrFieldXnMask) != 0
}

// SetXn Instruction access disable bit
func (r *registerMpu_rasrType) SetXn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMpu_rasrFieldXnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMpu_rasrFieldXnMask)
	}
}
