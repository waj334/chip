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
	Iciallu  RegisterIcialluType
	_        [4]byte
	Icimvau  RegisterIcimvauType
	Dcimvac  RegisterDcimvacType
	Dcisw    RegisterDciswType
	Dccmvau  RegisterDccmvauType
	Dccmvac  RegisterDccmvacType
	Dccsw    RegisterDccswType
	Dccimvac RegisterDccimvacType
	Dccisw   RegisterDcciswType
	Bpiall   RegisterBpiallType
}

// RegisterIcialluType Invalidates all instruction caches to Point of Unification (PoU)
type RegisterIcialluType uint32

func (r *RegisterIcialluType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIcialluType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIcialluType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIcialluType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIcialluType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIcialluFieldIgnoredShift = 0
	RegisterIcialluFieldIgnoredMask  = 0xffffffff
)

// GetIgnored The value written to this field is ignored
func (r *RegisterIcialluType) GetIgnored() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcialluFieldIgnoredMask) >> RegisterIcialluFieldIgnoredShift)
}

// SetIgnored The value written to this field is ignored
func (r *RegisterIcialluType) SetIgnored(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcialluFieldIgnoredMask)|(uint32(value)<<RegisterIcialluFieldIgnoredShift))
}

// RegisterIcimvauType Invalidates instruction cache lines by address to Point of Unification (PoU)
type RegisterIcimvauType uint32

func (r *RegisterIcimvauType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIcimvauType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIcimvauType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIcimvauType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIcimvauType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIcimvauFieldAddressShift = 0
	RegisterIcimvauFieldAddressMask  = 0xffffffff
)

// GetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *RegisterIcimvauType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcimvauFieldAddressMask) >> RegisterIcimvauFieldAddressShift)
}

// SetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *RegisterIcimvauType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcimvauFieldAddressMask)|(uint32(value)<<RegisterIcimvauFieldAddressShift))
}

// RegisterDcimvacType Invalidates data or unified cache lines by address to Point of Coherency (PoC)
type RegisterDcimvacType uint32

func (r *RegisterDcimvacType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDcimvacType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDcimvacType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDcimvacType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDcimvacType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDcimvacFieldAddressShift = 0
	RegisterDcimvacFieldAddressMask  = 0xffffffff
)

// GetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *RegisterDcimvacType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDcimvacFieldAddressMask) >> RegisterDcimvacFieldAddressShift)
}

// SetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *RegisterDcimvacType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcimvacFieldAddressMask)|(uint32(value)<<RegisterDcimvacFieldAddressShift))
}

// RegisterDciswType Invalidates instruction cache lines by address to Point of Unification (PoU)
type RegisterDciswType uint32

func (r *RegisterDciswType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDciswType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDciswType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDciswType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDciswType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDciswFieldLevelShift = 1
	RegisterDciswFieldLevelMask  = 0xe
)

// GetLevel Cache level. This field is 0b000.
func (r *RegisterDciswType) GetLevel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDciswFieldLevelMask) >> RegisterDciswFieldLevelShift)
}

// SetLevel Cache level. This field is 0b000.
func (r *RegisterDciswType) SetLevel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDciswFieldLevelMask)|(uint32(value)<<RegisterDciswFieldLevelShift))
}

const (
	RegisterDciswFieldSetwayShift = 4
	RegisterDciswFieldSetwayMask  = 0xfffffff0
)

// GetSetway Cache set/way. Contains two fields: Way, bits[31:32-A], the number of the way to operate on. Set, bits[B-1:L], the number of the set to operate on.
func (r *RegisterDciswType) GetSetway() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDciswFieldSetwayMask) >> RegisterDciswFieldSetwayShift)
}

// SetSetway Cache set/way. Contains two fields: Way, bits[31:32-A], the number of the way to operate on. Set, bits[B-1:L], the number of the set to operate on.
func (r *RegisterDciswType) SetSetway(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDciswFieldSetwayMask)|(uint32(value)<<RegisterDciswFieldSetwayShift))
}

// RegisterDccmvauType Cleans data or unified cache lines by address to Point of Unification (PoU)
type RegisterDccmvauType uint32

func (r *RegisterDccmvauType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDccmvauType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDccmvauType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDccmvauType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDccmvauType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDccmvauFieldAddressShift = 0
	RegisterDccmvauFieldAddressMask  = 0xffffffff
)

// GetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *RegisterDccmvauType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDccmvauFieldAddressMask) >> RegisterDccmvauFieldAddressShift)
}

// SetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *RegisterDccmvauType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDccmvauFieldAddressMask)|(uint32(value)<<RegisterDccmvauFieldAddressShift))
}

// RegisterDccmvacType Cleans data or unified cache lines by address to Point of Unification (PoC)
type RegisterDccmvacType uint32

func (r *RegisterDccmvacType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDccmvacType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDccmvacType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDccmvacType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDccmvacType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDccmvacFieldAddressShift = 0
	RegisterDccmvacFieldAddressMask  = 0xffffffff
)

// GetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *RegisterDccmvacType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDccmvacFieldAddressMask) >> RegisterDccmvacFieldAddressShift)
}

// SetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *RegisterDccmvacType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDccmvacFieldAddressMask)|(uint32(value)<<RegisterDccmvacFieldAddressShift))
}

// RegisterDccswType Cleans data or unified cache line by set/way
type RegisterDccswType uint32

func (r *RegisterDccswType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDccswType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDccswType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDccswType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDccswType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDccswFieldLevelShift = 1
	RegisterDccswFieldLevelMask  = 0xe
)

// GetLevel Cache level. This field is 0b000
func (r *RegisterDccswType) GetLevel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDccswFieldLevelMask) >> RegisterDccswFieldLevelShift)
}

// SetLevel Cache level. This field is 0b000
func (r *RegisterDccswType) SetLevel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDccswFieldLevelMask)|(uint32(value)<<RegisterDccswFieldLevelShift))
}

const (
	RegisterDccswFieldSetwayShift = 4
	RegisterDccswFieldSetwayMask  = 0xfffffff0
)

// GetSetway Contains two fields: Way, bits [31:32-A], the number of the way to operate on. Set, bits [B-1:L], the number of the set to operate on.
func (r *RegisterDccswType) GetSetway() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDccswFieldSetwayMask) >> RegisterDccswFieldSetwayShift)
}

// SetSetway Contains two fields: Way, bits [31:32-A], the number of the way to operate on. Set, bits [B-1:L], the number of the set to operate on.
func (r *RegisterDccswType) SetSetway(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDccswFieldSetwayMask)|(uint32(value)<<RegisterDccswFieldSetwayShift))
}

// RegisterDccimvacType Cleans data or unified cache lines by address to Point of Coherency (PoC)
type RegisterDccimvacType uint32

func (r *RegisterDccimvacType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDccimvacType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDccimvacType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDccimvacType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDccimvacType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDccimvacFieldAddressShift = 0
	RegisterDccimvacFieldAddressMask  = 0xffffffff
)

// GetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *RegisterDccimvacType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDccimvacFieldAddressMask) >> RegisterDccimvacFieldAddressShift)
}

// SetAddress Writing to this field initiates the maintenance operation for the address that is written
func (r *RegisterDccimvacType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDccimvacFieldAddressMask)|(uint32(value)<<RegisterDccimvacFieldAddressShift))
}

// RegisterDcciswType Cleans and invalidates data or unified cache line by set/way
type RegisterDcciswType uint32

func (r *RegisterDcciswType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDcciswType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDcciswType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDcciswType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDcciswType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDcciswFieldLevelShift = 1
	RegisterDcciswFieldLevelMask  = 0xe
)

// GetLevel Cache level. This field is 0b000.
func (r *RegisterDcciswType) GetLevel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDcciswFieldLevelMask) >> RegisterDcciswFieldLevelShift)
}

// SetLevel Cache level. This field is 0b000.
func (r *RegisterDcciswType) SetLevel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcciswFieldLevelMask)|(uint32(value)<<RegisterDcciswFieldLevelShift))
}

const (
	RegisterDcciswFieldSetwayShift = 4
	RegisterDcciswFieldSetwayMask  = 0xfffffff0
)

// GetSetway Contains two fields: Way, bits[31:32-A], the number of the way to operate on. Set, bits[B-1:L], the number of the set to operate on.
func (r *RegisterDcciswType) GetSetway() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDcciswFieldSetwayMask) >> RegisterDcciswFieldSetwayShift)
}

// SetSetway Contains two fields: Way, bits[31:32-A], the number of the way to operate on. Set, bits[B-1:L], the number of the set to operate on.
func (r *RegisterDcciswType) SetSetway(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcciswFieldSetwayMask)|(uint32(value)<<RegisterDcciswFieldSetwayShift))
}

// RegisterBpiallType Invalidates all entries from branch predictors
type RegisterBpiallType uint32

func (r *RegisterBpiallType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBpiallType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBpiallType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBpiallType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBpiallType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBpiallFieldIgnoredShift = 0
	RegisterBpiallFieldIgnoredMask  = 0xffffffff
)

// GetIgnored The value written to this field is ignored
func (r *RegisterBpiallType) GetIgnored() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterBpiallFieldIgnoredMask) >> RegisterBpiallFieldIgnoredShift)
}

// SetIgnored The value written to this field is ignored
func (r *RegisterBpiallType) SetIgnored(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBpiallFieldIgnoredMask)|(uint32(value)<<RegisterBpiallFieldIgnoredShift))
}
