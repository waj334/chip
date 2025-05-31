//go:build arm && cortexm

package cmo

import (
	"unsafe"
	"volatile"
)

var (
	Cmo = (*_cmo)(unsafe.Pointer(uintptr(0xe000ef50)))
)

type _cmo struct {
	Iciallu  registerIcialluType
	_        [4]byte
	Icimvau  registerIcimvauType
	Dcimvac  registerDcimvacType
	Dcisw    registerDciswType
	Dccmvau  registerDccmvauType
	Dccmvac  registerDccmvacType
	Dccsw    registerDccswType
	Dccimvac registerDccimvacType
	Dccisw   registerDcciswType
	Bpiall   registerBpiallType
}

// registerIcialluType Invalidates all instruction caches to Point of Unification (PoU)
type registerIcialluType uint32

const (
	RegisterIcialluFieldIgnoredShift = 0
	RegisterIcialluFieldIgnoredMask  = 0xffffffff
)

// GetIgnored The value written to this field is ignored
func (r *registerIcialluType) GetIgnored() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcialluFieldIgnoredMask) >> RegisterIcialluFieldIgnoredShift)
}

// SetIgnored The value written to this field is ignored
func (r *registerIcialluType) SetIgnored(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcialluFieldIgnoredMask)|(uint32(value)<<RegisterIcialluFieldIgnoredShift))
}

// registerIcimvauType Invalidates instruction cache lines by address to Point of Unification (PoU)
type registerIcimvauType uint32

const (
	RegisterIcimvauFieldAddressShift = 0
	RegisterIcimvauFieldAddressMask  = 0xffffffff
)

// GetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *registerIcimvauType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcimvauFieldAddressMask) >> RegisterIcimvauFieldAddressShift)
}

// SetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *registerIcimvauType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcimvauFieldAddressMask)|(uint32(value)<<RegisterIcimvauFieldAddressShift))
}

// registerDcimvacType Invalidates data or unified cache lines by address to Point of Coherency (PoC)
type registerDcimvacType uint32

const (
	RegisterDcimvacFieldAddressShift = 0
	RegisterDcimvacFieldAddressMask  = 0xffffffff
)

// GetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *registerDcimvacType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDcimvacFieldAddressMask) >> RegisterDcimvacFieldAddressShift)
}

// SetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *registerDcimvacType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcimvacFieldAddressMask)|(uint32(value)<<RegisterDcimvacFieldAddressShift))
}

// registerDciswType Invalidates instruction cache lines by address to Point of Unification (PoU)
type registerDciswType uint32

const (
	RegisterDciswFieldLevelShift = 1
	RegisterDciswFieldLevelMask  = 0xe
)

// GetLevel Cache level. This field is 0b000.
func (r *registerDciswType) GetLevel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDciswFieldLevelMask) >> RegisterDciswFieldLevelShift)
}

// SetLevel Cache level. This field is 0b000.
func (r *registerDciswType) SetLevel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDciswFieldLevelMask)|(uint32(value)<<RegisterDciswFieldLevelShift))
}

const (
	RegisterDciswFieldSetwayShift = 4
	RegisterDciswFieldSetwayMask  = 0xfffffff0
)

// GetSetway Cache set/way. Contains two fields: Way, bits[31:32-A], the number of the way to operate on. Set, bits[B-1:L], the number of the set to operate on.
func (r *registerDciswType) GetSetway() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDciswFieldSetwayMask) >> RegisterDciswFieldSetwayShift)
}

// SetSetway Cache set/way. Contains two fields: Way, bits[31:32-A], the number of the way to operate on. Set, bits[B-1:L], the number of the set to operate on.
func (r *registerDciswType) SetSetway(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDciswFieldSetwayMask)|(uint32(value)<<RegisterDciswFieldSetwayShift))
}

// registerDccmvauType Cleans data or unified cache lines by address to Point of Unification (PoU)
type registerDccmvauType uint32

const (
	RegisterDccmvauFieldAddressShift = 0
	RegisterDccmvauFieldAddressMask  = 0xffffffff
)

// GetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *registerDccmvauType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDccmvauFieldAddressMask) >> RegisterDccmvauFieldAddressShift)
}

// SetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *registerDccmvauType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDccmvauFieldAddressMask)|(uint32(value)<<RegisterDccmvauFieldAddressShift))
}

// registerDccmvacType Cleans data or unified cache lines by address to Point of Unification (PoC)
type registerDccmvacType uint32

const (
	RegisterDccmvacFieldAddressShift = 0
	RegisterDccmvacFieldAddressMask  = 0xffffffff
)

// GetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *registerDccmvacType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDccmvacFieldAddressMask) >> RegisterDccmvacFieldAddressShift)
}

// SetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *registerDccmvacType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDccmvacFieldAddressMask)|(uint32(value)<<RegisterDccmvacFieldAddressShift))
}

// registerDccswType Cleans data or unified cache line by set/way
type registerDccswType uint32

const (
	RegisterDccswFieldLevelShift = 1
	RegisterDccswFieldLevelMask  = 0xe
)

// GetLevel Cache level. This field is 0b000
func (r *registerDccswType) GetLevel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDccswFieldLevelMask) >> RegisterDccswFieldLevelShift)
}

// SetLevel Cache level. This field is 0b000
func (r *registerDccswType) SetLevel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDccswFieldLevelMask)|(uint32(value)<<RegisterDccswFieldLevelShift))
}

const (
	RegisterDccswFieldSetwayShift = 4
	RegisterDccswFieldSetwayMask  = 0xfffffff0
)

// GetSetway Contains two fields: Way, bits [31:32-A], the number of the way to operate on. Set, bits [B-1:L], the number of the set to operate on.
func (r *registerDccswType) GetSetway() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDccswFieldSetwayMask) >> RegisterDccswFieldSetwayShift)
}

// SetSetway Contains two fields: Way, bits [31:32-A], the number of the way to operate on. Set, bits [B-1:L], the number of the set to operate on.
func (r *registerDccswType) SetSetway(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDccswFieldSetwayMask)|(uint32(value)<<RegisterDccswFieldSetwayShift))
}

// registerDccimvacType Cleans data or unified cache lines by address to Point of Coherency (PoC)
type registerDccimvacType uint32

const (
	RegisterDccimvacFieldAddressShift = 0
	RegisterDccimvacFieldAddressMask  = 0xffffffff
)

// GetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *registerDccimvacType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDccimvacFieldAddressMask) >> RegisterDccimvacFieldAddressShift)
}

// SetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *registerDccimvacType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDccimvacFieldAddressMask)|(uint32(value)<<RegisterDccimvacFieldAddressShift))
}

// registerDcciswType Cleans and invalidates data or unified cache line by set/way
type registerDcciswType uint32

const (
	RegisterDcciswFieldLevelShift = 1
	RegisterDcciswFieldLevelMask  = 0xe
)

// GetLevel Cache level. This field is 0b000.
func (r *registerDcciswType) GetLevel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDcciswFieldLevelMask) >> RegisterDcciswFieldLevelShift)
}

// SetLevel Cache level. This field is 0b000.
func (r *registerDcciswType) SetLevel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcciswFieldLevelMask)|(uint32(value)<<RegisterDcciswFieldLevelShift))
}

const (
	RegisterDcciswFieldSetwayShift = 4
	RegisterDcciswFieldSetwayMask  = 0xfffffff0
)

// GetSetway Contains two fields: Way, bits[31:32-A], the number of the way to operate on. Set, bits[B-1:L], the number of the set to operate on.
func (r *registerDcciswType) GetSetway() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDcciswFieldSetwayMask) >> RegisterDcciswFieldSetwayShift)
}

// SetSetway Contains two fields: Way, bits[31:32-A], the number of the way to operate on. Set, bits[B-1:L], the number of the set to operate on.
func (r *registerDcciswType) SetSetway(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcciswFieldSetwayMask)|(uint32(value)<<RegisterDcciswFieldSetwayShift))
}

// registerBpiallType Invalidates all entries from branch predictors
type registerBpiallType uint32

const (
	RegisterBpiallFieldIgnoredShift = 0
	RegisterBpiallFieldIgnoredMask  = 0xffffffff
)

// GetIgnored The value written to this field is ignored
func (r *registerBpiallType) GetIgnored() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterBpiallFieldIgnoredMask) >> RegisterBpiallFieldIgnoredShift)
}

// SetIgnored The value written to this field is ignored
func (r *registerBpiallType) SetIgnored(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBpiallFieldIgnoredMask)|(uint32(value)<<RegisterBpiallFieldIgnoredShift))
}
