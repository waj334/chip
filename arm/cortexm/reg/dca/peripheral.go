//go:build arm && cortexm

package dca

import (
	"unsafe"
	"volatile"
)

var (
	Dca = (*_dca)(unsafe.Pointer(uintptr(0xe001e200)))
)

type _dca struct {
	Dcadcrr RegisterDcadcrrType
	Dcaicrr RegisterDcaicrrType
	_       [8]byte
	Dcadclr RegisterDcadclrType
	Dcaiclr RegisterDcaiclrType
}

// RegisterDcadcrrType Read the data from the Level 1 (L1) data cache from the location that is determined by the DCADCLR registers
type RegisterDcadcrrType uint32

func (r *RegisterDcadcrrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDcadcrrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDcadcrrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDcadcrrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDcadcrrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDcadcrrFieldTagShift = 0
	RegisterDcadcrrFieldTagMask  = 0x1fffff
)

// GetTag Tag address. The number of significant bits of TAG depends on the cache size.
func (r *RegisterDcadcrrType) GetTag() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDcadcrrFieldTagMask) >> RegisterDcadcrrFieldTagShift)
}

// SetTag Tag address. The number of significant bits of TAG depends on the cache size.
func (r *RegisterDcadcrrType) SetTag(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcadcrrFieldTagMask)|(uint32(value)<<RegisterDcadcrrFieldTagShift))
}

const (
	RegisterDcadcrrFieldValidShift = 21
	RegisterDcadcrrFieldValidMask  = 0x200000
)

// GetValid Valid state of the cache line
func (r *RegisterDcadcrrType) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcadcrrFieldValidMask) != 0
}

// SetValid Valid state of the cache line
func (r *RegisterDcadcrrType) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDcadcrrFieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDcadcrrFieldValidMask)
	}
}

// RegisterDcaicrrType Read the data from the Level 1 (L1) instruction cache from the location that is determined by the DCAICLR registers
type RegisterDcaicrrType uint32

func (r *RegisterDcaicrrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDcaicrrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDcaicrrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDcaicrrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDcaicrrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDcaicrrFieldTagShift = 0
	RegisterDcaicrrFieldTagMask  = 0x1fffff
)

// GetTag Tag address. The number of significant bits of TAG depends on the cache size.
func (r *RegisterDcaicrrType) GetTag() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDcaicrrFieldTagMask) >> RegisterDcaicrrFieldTagShift)
}

// SetTag Tag address. The number of significant bits of TAG depends on the cache size.
func (r *RegisterDcaicrrType) SetTag(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcaicrrFieldTagMask)|(uint32(value)<<RegisterDcaicrrFieldTagShift))
}

const (
	RegisterDcaicrrFieldValidShift = 21
	RegisterDcaicrrFieldValidMask  = 0x200000
)

// GetValid Valid state of the cache line
func (r *RegisterDcaicrrType) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcaicrrFieldValidMask) != 0
}

// SetValid Valid state of the cache line
func (r *RegisterDcaicrrType) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDcaicrrFieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDcaicrrFieldValidMask)
	}
}

// RegisterDcadclrType Set the location to be read from the Level 1 (L1) data cache
type RegisterDcadclrType uint32

func (r *RegisterDcadclrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDcadclrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDcadclrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDcadclrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDcadclrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDcadclrFieldRamtypeShift = 0
	RegisterDcadclrFieldRamtypeMask  = 0x1
)

// GetRamtype RAM type
func (r *RegisterDcadclrType) GetRamtype() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcadclrFieldRamtypeMask) != 0
}

// SetRamtype RAM type
func (r *RegisterDcadclrType) SetRamtype(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDcadclrFieldRamtypeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDcadclrFieldRamtypeMask)
	}
}

const (
	RegisterDcadclrFieldOffsetShift = 2
	RegisterDcadclrFieldOffsetMask  = 0x1c
)

// GetOffset Data offset
func (r *RegisterDcadclrType) GetOffset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDcadclrFieldOffsetMask) >> RegisterDcadclrFieldOffsetShift)
}

// SetOffset Data offset
func (r *RegisterDcadclrType) SetOffset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcadclrFieldOffsetMask)|(uint32(value)<<RegisterDcadclrFieldOffsetShift))
}

const (
	RegisterDcadclrFieldSetShift = 5
	RegisterDcadclrFieldSetMask  = 0x7fe0
)

// GetSet Set index. The number of significant bits of SET depends on the cache size.
func (r *RegisterDcadclrType) GetSet() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDcadclrFieldSetMask) >> RegisterDcadclrFieldSetShift)
}

// SetSet Set index. The number of significant bits of SET depends on the cache size.
func (r *RegisterDcadclrType) SetSet(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcadclrFieldSetMask)|(uint32(value)<<RegisterDcadclrFieldSetShift))
}

const (
	RegisterDcadclrFieldWayShift = 30
	RegisterDcadclrFieldWayMask  = 0x40000000
)

// GetWay Cache way
func (r *RegisterDcadclrType) GetWay() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcadclrFieldWayMask) != 0
}

// SetWay Cache way
func (r *RegisterDcadclrType) SetWay(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDcadclrFieldWayMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDcadclrFieldWayMask)
	}
}

// RegisterDcaiclrType Set the location to be read from the Level 1 (L1) instruction cache
type RegisterDcaiclrType uint32

func (r *RegisterDcaiclrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDcaiclrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDcaiclrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDcaiclrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDcaiclrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDcaiclrFieldRamtypeShift = 0
	RegisterDcaiclrFieldRamtypeMask  = 0x1
)

// GetRamtype RAM type
func (r *RegisterDcaiclrType) GetRamtype() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcaiclrFieldRamtypeMask) != 0
}

// SetRamtype RAM type
func (r *RegisterDcaiclrType) SetRamtype(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDcaiclrFieldRamtypeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDcaiclrFieldRamtypeMask)
	}
}

const (
	RegisterDcaiclrFieldOffsetShift = 2
	RegisterDcaiclrFieldOffsetMask  = 0x1c
)

// GetOffset Data offset
func (r *RegisterDcaiclrType) GetOffset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDcaiclrFieldOffsetMask) >> RegisterDcaiclrFieldOffsetShift)
}

// SetOffset Data offset
func (r *RegisterDcaiclrType) SetOffset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcaiclrFieldOffsetMask)|(uint32(value)<<RegisterDcaiclrFieldOffsetShift))
}

const (
	RegisterDcaiclrFieldSetShift = 5
	RegisterDcaiclrFieldSetMask  = 0x7fe0
)

// GetSet Set index. The number of significant bits of SET depends on the cache size.
func (r *RegisterDcaiclrType) GetSet() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDcaiclrFieldSetMask) >> RegisterDcaiclrFieldSetShift)
}

// SetSet Set index. The number of significant bits of SET depends on the cache size.
func (r *RegisterDcaiclrType) SetSet(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcaiclrFieldSetMask)|(uint32(value)<<RegisterDcaiclrFieldSetShift))
}

const (
	RegisterDcaiclrFieldWayShift = 30
	RegisterDcaiclrFieldWayMask  = 0x40000000
)

// GetWay Cache way
func (r *RegisterDcaiclrType) GetWay() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcaiclrFieldWayMask) != 0
}

// SetWay Cache way
func (r *RegisterDcaiclrType) SetWay(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDcaiclrFieldWayMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDcaiclrFieldWayMask)
	}
}
