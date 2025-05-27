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
	Dcadcrr registerDcadcrrType
	Dcaicrr registerDcaicrrType
	_       [8]byte
	Dcadclr registerDcadclrType
	Dcaiclr registerDcaiclrType
}

// registerDcadcrrType Read the data from the Level 1 (L1) data cache from the location that is determined by the DCADCLR registers
type registerDcadcrrType uint32

const (
	RegisterDcadcrrFieldTagShift = 0
	RegisterDcadcrrFieldTagMask  = 0x1fffff
)

// GetTag Tag address. The number of significant bits of TAG depends on the cache size.
func (r *registerDcadcrrType) GetTag() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDcadcrrFieldTagMask) >> RegisterDcadcrrFieldTagShift)
}

// SetTag Tag address. The number of significant bits of TAG depends on the cache size.
func (r *registerDcadcrrType) SetTag(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcadcrrFieldTagMask)|(uint32(value)<<RegisterDcadcrrFieldTagShift))
}

const (
	RegisterDcadcrrFieldValidShift = 21
	RegisterDcadcrrFieldValidMask  = 0x200000
)

// GetValid Valid state of the cache line
func (r *registerDcadcrrType) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcadcrrFieldValidMask) != 0
}

// SetValid Valid state of the cache line
func (r *registerDcadcrrType) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDcadcrrFieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDcadcrrFieldValidMask)
	}
}

// registerDcaicrrType Read the data from the Level 1 (L1) instruction cache from the location that is determined by the DCAICLR registers
type registerDcaicrrType uint32

const (
	RegisterDcaicrrFieldTagShift = 0
	RegisterDcaicrrFieldTagMask  = 0x1fffff
)

// GetTag Tag address. The number of significant bits of TAG depends on the cache size.
func (r *registerDcaicrrType) GetTag() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDcaicrrFieldTagMask) >> RegisterDcaicrrFieldTagShift)
}

// SetTag Tag address. The number of significant bits of TAG depends on the cache size.
func (r *registerDcaicrrType) SetTag(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcaicrrFieldTagMask)|(uint32(value)<<RegisterDcaicrrFieldTagShift))
}

const (
	RegisterDcaicrrFieldValidShift = 21
	RegisterDcaicrrFieldValidMask  = 0x200000
)

// GetValid Valid state of the cache line
func (r *registerDcaicrrType) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcaicrrFieldValidMask) != 0
}

// SetValid Valid state of the cache line
func (r *registerDcaicrrType) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDcaicrrFieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDcaicrrFieldValidMask)
	}
}

// registerDcadclrType Set the location to be read from the Level 1 (L1) data cache
type registerDcadclrType uint32

const (
	RegisterDcadclrFieldRamtypeShift = 0
	RegisterDcadclrFieldRamtypeMask  = 0x1
)

// GetRamtype RAM type
func (r *registerDcadclrType) GetRamtype() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcadclrFieldRamtypeMask) != 0
}

// SetRamtype RAM type
func (r *registerDcadclrType) SetRamtype(value bool) {
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
func (r *registerDcadclrType) GetOffset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDcadclrFieldOffsetMask) >> RegisterDcadclrFieldOffsetShift)
}

// SetOffset Data offset
func (r *registerDcadclrType) SetOffset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcadclrFieldOffsetMask)|(uint32(value)<<RegisterDcadclrFieldOffsetShift))
}

const (
	RegisterDcadclrFieldSetShift = 5
	RegisterDcadclrFieldSetMask  = 0x7fe0
)

// GetSet Set index. The number of significant bits of SET depends on the cache size.
func (r *registerDcadclrType) GetSet() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDcadclrFieldSetMask) >> RegisterDcadclrFieldSetShift)
}

// SetSet Set index. The number of significant bits of SET depends on the cache size.
func (r *registerDcadclrType) SetSet(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcadclrFieldSetMask)|(uint32(value)<<RegisterDcadclrFieldSetShift))
}

const (
	RegisterDcadclrFieldWayShift = 30
	RegisterDcadclrFieldWayMask  = 0x40000000
)

// GetWay Cache way
func (r *registerDcadclrType) GetWay() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcadclrFieldWayMask) != 0
}

// SetWay Cache way
func (r *registerDcadclrType) SetWay(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDcadclrFieldWayMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDcadclrFieldWayMask)
	}
}

// registerDcaiclrType Set the location to be read from the Level 1 (L1) instruction cache
type registerDcaiclrType uint32

const (
	RegisterDcaiclrFieldRamtypeShift = 0
	RegisterDcaiclrFieldRamtypeMask  = 0x1
)

// GetRamtype RAM type
func (r *registerDcaiclrType) GetRamtype() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcaiclrFieldRamtypeMask) != 0
}

// SetRamtype RAM type
func (r *registerDcaiclrType) SetRamtype(value bool) {
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
func (r *registerDcaiclrType) GetOffset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDcaiclrFieldOffsetMask) >> RegisterDcaiclrFieldOffsetShift)
}

// SetOffset Data offset
func (r *registerDcaiclrType) SetOffset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcaiclrFieldOffsetMask)|(uint32(value)<<RegisterDcaiclrFieldOffsetShift))
}

const (
	RegisterDcaiclrFieldSetShift = 5
	RegisterDcaiclrFieldSetMask  = 0x7fe0
)

// GetSet Set index. The number of significant bits of SET depends on the cache size.
func (r *registerDcaiclrType) GetSet() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDcaiclrFieldSetMask) >> RegisterDcaiclrFieldSetShift)
}

// SetSet Set index. The number of significant bits of SET depends on the cache size.
func (r *registerDcaiclrType) SetSet(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcaiclrFieldSetMask)|(uint32(value)<<RegisterDcaiclrFieldSetShift))
}

const (
	RegisterDcaiclrFieldWayShift = 30
	RegisterDcaiclrFieldWayMask  = 0x40000000
)

// GetWay Cache way
func (r *registerDcaiclrType) GetWay() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcaiclrFieldWayMask) != 0
}

// SetWay Cache way
func (r *registerDcaiclrType) SetWay(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDcaiclrFieldWayMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDcaiclrFieldWayMask)
	}
}
