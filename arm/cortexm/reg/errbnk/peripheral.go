//go:build arm && cortexm

package errbnk

import (
	"unsafe"
	"volatile"
)

var (
	Errbnk = (*_errbnk)(unsafe.Pointer(uintptr(0xe001e100)))
)

type _errbnk struct {
	Iebr0 RegisterIebr0Type
	Iebr1 RegisterIebr1Type
	_     [8]byte
	Debr0 RegisterDebr0Type
	Debr1 RegisterDebr1Type
	_     [8]byte
	Tebr0 RegisterTebr0Type
	_     [4]byte
	Tebr1 RegisterTebr1Type
}

// RegisterIebr0Type Record errors that occur during memory accesses to the L1 instruction cache
type RegisterIebr0Type uint32

func (r *RegisterIebr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIebr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIebr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIebr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIebr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIebr0FieldValidShift = 0
	RegisterIebr0FieldValidMask  = 0x1
)

// GetValid Indicates whether the entry is valid or not
func (r *RegisterIebr0Type) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIebr0FieldValidMask) != 0
}

// SetValid Indicates whether the entry is valid or not
func (r *RegisterIebr0Type) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIebr0FieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIebr0FieldValidMask)
	}
}

const (
	RegisterIebr0FieldLockedShift = 1
	RegisterIebr0FieldLockedMask  = 0x2
)

// GetLocked Indicates whether the location is locked or not
func (r *RegisterIebr0Type) GetLocked() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIebr0FieldLockedMask) != 0
}

// SetLocked Indicates whether the location is locked or not
func (r *RegisterIebr0Type) SetLocked(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIebr0FieldLockedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIebr0FieldLockedMask)
	}
}

const (
	RegisterIebr0FieldLocationlwoShift = 2
	RegisterIebr0FieldLocationlwoMask  = 0x1c
)

// GetLocationlwo Indicates the location in the L1 instruction cache RAM, line word offset
func (r *RegisterIebr0Type) GetLocationlwo() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIebr0FieldLocationlwoMask) >> RegisterIebr0FieldLocationlwoShift)
}

// SetLocationlwo Indicates the location in the L1 instruction cache RAM, line word offset
func (r *RegisterIebr0Type) SetLocationlwo(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIebr0FieldLocationlwoMask)|(uint32(value)<<RegisterIebr0FieldLocationlwoShift))
}

const (
	RegisterIebr0FieldLocationindexShift = 5
	RegisterIebr0FieldLocationindexMask  = 0x7fe0
)

// GetLocationindex Indicates the location in the L1 instruction cache RAM, index
func (r *RegisterIebr0Type) GetLocationindex() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterIebr0FieldLocationindexMask) >> RegisterIebr0FieldLocationindexShift)
}

// SetLocationindex Indicates the location in the L1 instruction cache RAM, index
func (r *RegisterIebr0Type) SetLocationindex(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIebr0FieldLocationindexMask)|(uint32(value)<<RegisterIebr0FieldLocationindexShift))
}

const (
	RegisterIebr0FieldLocationwayShift = 15
	RegisterIebr0FieldLocationwayMask  = 0x8000
)

// GetLocationway Indicates the location in the L1 instruction cache RAM, way
func (r *RegisterIebr0Type) GetLocationway() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIebr0FieldLocationwayMask) != 0
}

// SetLocationway Indicates the location in the L1 instruction cache RAM, way
func (r *RegisterIebr0Type) SetLocationway(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIebr0FieldLocationwayMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIebr0FieldLocationwayMask)
	}
}

const (
	RegisterIebr0FieldBankShift = 16
	RegisterIebr0FieldBankMask  = 0x10000
)

// GetBank Indicates which RAM bank to use
func (r *RegisterIebr0Type) GetBank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIebr0FieldBankMask) != 0
}

// SetBank Indicates which RAM bank to use
func (r *RegisterIebr0Type) SetBank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIebr0FieldBankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIebr0FieldBankMask)
	}
}

const (
	RegisterIebr0FieldSwdefShift = 30
	RegisterIebr0FieldSwdefMask  = 0xc0000000
)

// GetSwdef Error detection logic sets this field to 0b00 on a new allocation and on Cold reset
func (r *RegisterIebr0Type) GetSwdef() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIebr0FieldSwdefMask) >> RegisterIebr0FieldSwdefShift)
}

// SetSwdef Error detection logic sets this field to 0b00 on a new allocation and on Cold reset
func (r *RegisterIebr0Type) SetSwdef(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIebr0FieldSwdefMask)|(uint32(value)<<RegisterIebr0FieldSwdefShift))
}

// RegisterIebr1Type Record errors that occur during memory accesses to the L1 instruction cache
type RegisterIebr1Type uint32

func (r *RegisterIebr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIebr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIebr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIebr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIebr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIebr1FieldValidShift = 0
	RegisterIebr1FieldValidMask  = 0x1
)

// GetValid Indicates whether the entry is valid or not
func (r *RegisterIebr1Type) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIebr1FieldValidMask) != 0
}

// SetValid Indicates whether the entry is valid or not
func (r *RegisterIebr1Type) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIebr1FieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIebr1FieldValidMask)
	}
}

const (
	RegisterIebr1FieldLockedShift = 1
	RegisterIebr1FieldLockedMask  = 0x2
)

// GetLocked Indicates whether the location is locked or not
func (r *RegisterIebr1Type) GetLocked() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIebr1FieldLockedMask) != 0
}

// SetLocked Indicates whether the location is locked or not
func (r *RegisterIebr1Type) SetLocked(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIebr1FieldLockedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIebr1FieldLockedMask)
	}
}

const (
	RegisterIebr1FieldLocationlwoShift = 2
	RegisterIebr1FieldLocationlwoMask  = 0x1c
)

// GetLocationlwo Indicates the location in the L1 instruction cache RAM, line word offset
func (r *RegisterIebr1Type) GetLocationlwo() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIebr1FieldLocationlwoMask) >> RegisterIebr1FieldLocationlwoShift)
}

// SetLocationlwo Indicates the location in the L1 instruction cache RAM, line word offset
func (r *RegisterIebr1Type) SetLocationlwo(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIebr1FieldLocationlwoMask)|(uint32(value)<<RegisterIebr1FieldLocationlwoShift))
}

const (
	RegisterIebr1FieldLocationindexShift = 5
	RegisterIebr1FieldLocationindexMask  = 0x7fe0
)

// GetLocationindex Indicates the location in the L1 instruction cache RAM, index
func (r *RegisterIebr1Type) GetLocationindex() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterIebr1FieldLocationindexMask) >> RegisterIebr1FieldLocationindexShift)
}

// SetLocationindex Indicates the location in the L1 instruction cache RAM, index
func (r *RegisterIebr1Type) SetLocationindex(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIebr1FieldLocationindexMask)|(uint32(value)<<RegisterIebr1FieldLocationindexShift))
}

const (
	RegisterIebr1FieldLocationwayShift = 15
	RegisterIebr1FieldLocationwayMask  = 0x8000
)

// GetLocationway Indicates the location in the L1 instruction cache RAM, way
func (r *RegisterIebr1Type) GetLocationway() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIebr1FieldLocationwayMask) != 0
}

// SetLocationway Indicates the location in the L1 instruction cache RAM, way
func (r *RegisterIebr1Type) SetLocationway(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIebr1FieldLocationwayMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIebr1FieldLocationwayMask)
	}
}

const (
	RegisterIebr1FieldBankShift = 16
	RegisterIebr1FieldBankMask  = 0x10000
)

// GetBank Indicates which RAM bank to use
func (r *RegisterIebr1Type) GetBank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIebr1FieldBankMask) != 0
}

// SetBank Indicates which RAM bank to use
func (r *RegisterIebr1Type) SetBank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIebr1FieldBankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIebr1FieldBankMask)
	}
}

const (
	RegisterIebr1FieldSwdefShift = 30
	RegisterIebr1FieldSwdefMask  = 0xc0000000
)

// GetSwdef Error detection logic sets this field to 0b00 on a new allocation and on Cold reset
func (r *RegisterIebr1Type) GetSwdef() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIebr1FieldSwdefMask) >> RegisterIebr1FieldSwdefShift)
}

// SetSwdef Error detection logic sets this field to 0b00 on a new allocation and on Cold reset
func (r *RegisterIebr1Type) SetSwdef(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIebr1FieldSwdefMask)|(uint32(value)<<RegisterIebr1FieldSwdefShift))
}

// RegisterDebr0Type Record errors that occur during memory accesses to the L1 data cache
type RegisterDebr0Type uint32

func (r *RegisterDebr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDebr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDebr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDebr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDebr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDebr0FieldValidShift = 0
	RegisterDebr0FieldValidMask  = 0x1
)

// GetValid Indicates whether the entry is valid or not
func (r *RegisterDebr0Type) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDebr0FieldValidMask) != 0
}

// SetValid Indicates whether the entry is valid or not
func (r *RegisterDebr0Type) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDebr0FieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDebr0FieldValidMask)
	}
}

const (
	RegisterDebr0FieldLockedShift = 1
	RegisterDebr0FieldLockedMask  = 0x2
)

// GetLocked Indicates whether the location is locked or not
func (r *RegisterDebr0Type) GetLocked() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDebr0FieldLockedMask) != 0
}

// SetLocked Indicates whether the location is locked or not
func (r *RegisterDebr0Type) SetLocked(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDebr0FieldLockedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDebr0FieldLockedMask)
	}
}

const (
	RegisterDebr0FieldLocationlwoShift = 2
	RegisterDebr0FieldLocationlwoMask  = 0x1c
)

// GetLocationlwo Indicates the location in the L1 instruction cache RAM, line doubleword offset
func (r *RegisterDebr0Type) GetLocationlwo() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDebr0FieldLocationlwoMask) >> RegisterDebr0FieldLocationlwoShift)
}

// SetLocationlwo Indicates the location in the L1 instruction cache RAM, line doubleword offset
func (r *RegisterDebr0Type) SetLocationlwo(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDebr0FieldLocationlwoMask)|(uint32(value)<<RegisterDebr0FieldLocationlwoShift))
}

const (
	RegisterDebr0FieldLocationindexShift = 5
	RegisterDebr0FieldLocationindexMask  = 0x3fe0
)

// GetLocationindex Indicates the location in the L1 instruction cache RAM, index
func (r *RegisterDebr0Type) GetLocationindex() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDebr0FieldLocationindexMask) >> RegisterDebr0FieldLocationindexShift)
}

// SetLocationindex Indicates the location in the L1 instruction cache RAM, index
func (r *RegisterDebr0Type) SetLocationindex(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDebr0FieldLocationindexMask)|(uint32(value)<<RegisterDebr0FieldLocationindexShift))
}

const (
	RegisterDebr0FieldLocationwayShift = 14
	RegisterDebr0FieldLocationwayMask  = 0xc000
)

// GetLocationway Indicates the location in the L1 instruction cache RAM, way
func (r *RegisterDebr0Type) GetLocationway() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDebr0FieldLocationwayMask) >> RegisterDebr0FieldLocationwayShift)
}

// SetLocationway Indicates the location in the L1 instruction cache RAM, way
func (r *RegisterDebr0Type) SetLocationway(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDebr0FieldLocationwayMask)|(uint32(value)<<RegisterDebr0FieldLocationwayShift))
}

const (
	RegisterDebr0FieldBankShift = 16
	RegisterDebr0FieldBankMask  = 0x10000
)

// GetBank Indicates which RAM bank to use
func (r *RegisterDebr0Type) GetBank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDebr0FieldBankMask) != 0
}

// SetBank Indicates which RAM bank to use
func (r *RegisterDebr0Type) SetBank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDebr0FieldBankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDebr0FieldBankMask)
	}
}

const (
	RegisterDebr0FieldTypeShift = 17
	RegisterDebr0FieldTypeMask  = 0x20000
)

// GetType Indicates the error type
func (r *RegisterDebr0Type) GetType() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDebr0FieldTypeMask) != 0
}

// SetType Indicates the error type
func (r *RegisterDebr0Type) SetType(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDebr0FieldTypeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDebr0FieldTypeMask)
	}
}

const (
	RegisterDebr0FieldSwdefShift = 30
	RegisterDebr0FieldSwdefMask  = 0xc0000000
)

// GetSwdef Error detection logic sets this field to 0b00 on a new allocation and on Cold reset
func (r *RegisterDebr0Type) GetSwdef() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDebr0FieldSwdefMask) >> RegisterDebr0FieldSwdefShift)
}

// SetSwdef Error detection logic sets this field to 0b00 on a new allocation and on Cold reset
func (r *RegisterDebr0Type) SetSwdef(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDebr0FieldSwdefMask)|(uint32(value)<<RegisterDebr0FieldSwdefShift))
}

// RegisterDebr1Type Record errors that occur during memory accesses to the L1 data cache
type RegisterDebr1Type uint32

func (r *RegisterDebr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDebr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDebr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDebr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDebr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDebr1FieldValidShift = 0
	RegisterDebr1FieldValidMask  = 0x1
)

// GetValid Indicates whether the entry is valid or not
func (r *RegisterDebr1Type) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDebr1FieldValidMask) != 0
}

// SetValid Indicates whether the entry is valid or not
func (r *RegisterDebr1Type) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDebr1FieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDebr1FieldValidMask)
	}
}

const (
	RegisterDebr1FieldLockedShift = 1
	RegisterDebr1FieldLockedMask  = 0x2
)

// GetLocked Indicates whether the location is locked or not
func (r *RegisterDebr1Type) GetLocked() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDebr1FieldLockedMask) != 0
}

// SetLocked Indicates whether the location is locked or not
func (r *RegisterDebr1Type) SetLocked(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDebr1FieldLockedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDebr1FieldLockedMask)
	}
}

const (
	RegisterDebr1FieldLocationlwoShift = 2
	RegisterDebr1FieldLocationlwoMask  = 0x1c
)

// GetLocationlwo Indicates the location in the L1 instruction cache RAM, line doubleword offset
func (r *RegisterDebr1Type) GetLocationlwo() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDebr1FieldLocationlwoMask) >> RegisterDebr1FieldLocationlwoShift)
}

// SetLocationlwo Indicates the location in the L1 instruction cache RAM, line doubleword offset
func (r *RegisterDebr1Type) SetLocationlwo(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDebr1FieldLocationlwoMask)|(uint32(value)<<RegisterDebr1FieldLocationlwoShift))
}

const (
	RegisterDebr1FieldLocationindexShift = 5
	RegisterDebr1FieldLocationindexMask  = 0x3fe0
)

// GetLocationindex Indicates the location in the L1 instruction cache RAM, index
func (r *RegisterDebr1Type) GetLocationindex() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDebr1FieldLocationindexMask) >> RegisterDebr1FieldLocationindexShift)
}

// SetLocationindex Indicates the location in the L1 instruction cache RAM, index
func (r *RegisterDebr1Type) SetLocationindex(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDebr1FieldLocationindexMask)|(uint32(value)<<RegisterDebr1FieldLocationindexShift))
}

const (
	RegisterDebr1FieldLocationwayShift = 14
	RegisterDebr1FieldLocationwayMask  = 0xc000
)

// GetLocationway Indicates the location in the L1 instruction cache RAM, way
func (r *RegisterDebr1Type) GetLocationway() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDebr1FieldLocationwayMask) >> RegisterDebr1FieldLocationwayShift)
}

// SetLocationway Indicates the location in the L1 instruction cache RAM, way
func (r *RegisterDebr1Type) SetLocationway(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDebr1FieldLocationwayMask)|(uint32(value)<<RegisterDebr1FieldLocationwayShift))
}

const (
	RegisterDebr1FieldBankShift = 16
	RegisterDebr1FieldBankMask  = 0x10000
)

// GetBank Indicates which RAM bank to use
func (r *RegisterDebr1Type) GetBank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDebr1FieldBankMask) != 0
}

// SetBank Indicates which RAM bank to use
func (r *RegisterDebr1Type) SetBank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDebr1FieldBankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDebr1FieldBankMask)
	}
}

const (
	RegisterDebr1FieldTypeShift = 17
	RegisterDebr1FieldTypeMask  = 0x20000
)

// GetType Indicates the error type
func (r *RegisterDebr1Type) GetType() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDebr1FieldTypeMask) != 0
}

// SetType Indicates the error type
func (r *RegisterDebr1Type) SetType(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDebr1FieldTypeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDebr1FieldTypeMask)
	}
}

const (
	RegisterDebr1FieldSwdefShift = 30
	RegisterDebr1FieldSwdefMask  = 0xc0000000
)

// GetSwdef Error detection logic sets this field to 0b00 on a new allocation and on Cold reset
func (r *RegisterDebr1Type) GetSwdef() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDebr1FieldSwdefMask) >> RegisterDebr1FieldSwdefShift)
}

// SetSwdef Error detection logic sets this field to 0b00 on a new allocation and on Cold reset
func (r *RegisterDebr1Type) SetSwdef(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDebr1FieldSwdefMask)|(uint32(value)<<RegisterDebr1FieldSwdefShift))
}

// RegisterTebr0Type Record the location of errors in the TCM
type RegisterTebr0Type uint32

func (r *RegisterTebr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTebr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTebr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTebr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTebr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTebr0FieldValidShift = 0
	RegisterTebr0FieldValidMask  = 0x1
)

// GetValid Indicates whether the entry is valid or not
func (r *RegisterTebr0Type) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTebr0FieldValidMask) != 0
}

// SetValid Indicates whether the entry is valid or not
func (r *RegisterTebr0Type) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTebr0FieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTebr0FieldValidMask)
	}
}

const (
	RegisterTebr0FieldLockedShift = 1
	RegisterTebr0FieldLockedMask  = 0x2
)

// GetLocked Indicates whether the location is locked or not
func (r *RegisterTebr0Type) GetLocked() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTebr0FieldLockedMask) != 0
}

// SetLocked Indicates whether the location is locked or not
func (r *RegisterTebr0Type) SetLocked(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTebr0FieldLockedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTebr0FieldLockedMask)
	}
}

const (
	RegisterTebr0FieldLocationShift = 2
	RegisterTebr0FieldLocationMask  = 0xfffffc
)

// GetLocation Indicates the physical location in the data cache RAM
func (r *RegisterTebr0Type) GetLocation() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTebr0FieldLocationMask) >> RegisterTebr0FieldLocationShift)
}

// SetLocation Indicates the physical location in the data cache RAM
func (r *RegisterTebr0Type) SetLocation(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTebr0FieldLocationMask)|(uint32(value)<<RegisterTebr0FieldLocationShift))
}

const (
	RegisterTebr0FieldBankShift = 24
	RegisterTebr0FieldBankMask  = 0x7000000
)

// GetBank Indicates which RAM bank to use
func (r *RegisterTebr0Type) GetBank() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTebr0FieldBankMask) >> RegisterTebr0FieldBankShift)
}

// SetBank Indicates which RAM bank to use
func (r *RegisterTebr0Type) SetBank(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTebr0FieldBankMask)|(uint32(value)<<RegisterTebr0FieldBankShift))
}

const (
	RegisterTebr0FieldTypeShift = 27
	RegisterTebr0FieldTypeMask  = 0x8000000
)

// GetType Indicates the error type
func (r *RegisterTebr0Type) GetType() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTebr0FieldTypeMask) != 0
}

// SetType Indicates the error type
func (r *RegisterTebr0Type) SetType(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTebr0FieldTypeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTebr0FieldTypeMask)
	}
}

const (
	RegisterTebr0FieldPoisonShift = 28
	RegisterTebr0FieldPoisonMask  = 0x10000000
)

// GetPoison Indicates whether a BusFault is generated or not
func (r *RegisterTebr0Type) GetPoison() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTebr0FieldPoisonMask) != 0
}

// SetPoison Indicates whether a BusFault is generated or not
func (r *RegisterTebr0Type) SetPoison(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTebr0FieldPoisonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTebr0FieldPoisonMask)
	}
}

const (
	RegisterTebr0FieldSwdefShift = 30
	RegisterTebr0FieldSwdefMask  = 0xc0000000
)

// GetSwdef Error detection logic sets this field to 0b00 on a new allocation and on Cold reset
func (r *RegisterTebr0Type) GetSwdef() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTebr0FieldSwdefMask) >> RegisterTebr0FieldSwdefShift)
}

// SetSwdef Error detection logic sets this field to 0b00 on a new allocation and on Cold reset
func (r *RegisterTebr0Type) SetSwdef(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTebr0FieldSwdefMask)|(uint32(value)<<RegisterTebr0FieldSwdefShift))
}

// RegisterTebr1Type Record the location of errors in the TCM
type RegisterTebr1Type uint32

func (r *RegisterTebr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTebr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTebr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTebr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTebr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTebr1FieldValidShift = 0
	RegisterTebr1FieldValidMask  = 0x1
)

// GetValid Indicates whether the entry is valid or not
func (r *RegisterTebr1Type) GetValid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTebr1FieldValidMask) != 0
}

// SetValid Indicates whether the entry is valid or not
func (r *RegisterTebr1Type) SetValid(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTebr1FieldValidMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTebr1FieldValidMask)
	}
}

const (
	RegisterTebr1FieldLockedShift = 1
	RegisterTebr1FieldLockedMask  = 0x2
)

// GetLocked Indicates whether the location is locked or not
func (r *RegisterTebr1Type) GetLocked() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTebr1FieldLockedMask) != 0
}

// SetLocked Indicates whether the location is locked or not
func (r *RegisterTebr1Type) SetLocked(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTebr1FieldLockedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTebr1FieldLockedMask)
	}
}

const (
	RegisterTebr1FieldLocationShift = 2
	RegisterTebr1FieldLocationMask  = 0xfffffc
)

// GetLocation Indicates the physical location in the data cache RAM
func (r *RegisterTebr1Type) GetLocation() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTebr1FieldLocationMask) >> RegisterTebr1FieldLocationShift)
}

// SetLocation Indicates the physical location in the data cache RAM
func (r *RegisterTebr1Type) SetLocation(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTebr1FieldLocationMask)|(uint32(value)<<RegisterTebr1FieldLocationShift))
}

const (
	RegisterTebr1FieldBankShift = 24
	RegisterTebr1FieldBankMask  = 0x7000000
)

// GetBank Indicates which RAM bank to use
func (r *RegisterTebr1Type) GetBank() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTebr1FieldBankMask) >> RegisterTebr1FieldBankShift)
}

// SetBank Indicates which RAM bank to use
func (r *RegisterTebr1Type) SetBank(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTebr1FieldBankMask)|(uint32(value)<<RegisterTebr1FieldBankShift))
}

const (
	RegisterTebr1FieldTypeShift = 27
	RegisterTebr1FieldTypeMask  = 0x8000000
)

// GetType Indicates the error type
func (r *RegisterTebr1Type) GetType() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTebr1FieldTypeMask) != 0
}

// SetType Indicates the error type
func (r *RegisterTebr1Type) SetType(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTebr1FieldTypeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTebr1FieldTypeMask)
	}
}

const (
	RegisterTebr1FieldPoisonShift = 28
	RegisterTebr1FieldPoisonMask  = 0x10000000
)

// GetPoison Indicates whether a BusFault is generated or not
func (r *RegisterTebr1Type) GetPoison() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTebr1FieldPoisonMask) != 0
}

// SetPoison Indicates whether a BusFault is generated or not
func (r *RegisterTebr1Type) SetPoison(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTebr1FieldPoisonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTebr1FieldPoisonMask)
	}
}

const (
	RegisterTebr1FieldSwdefShift = 30
	RegisterTebr1FieldSwdefMask  = 0xc0000000
)

// GetSwdef Error detection logic sets this field to 0b00 on a new allocation and on Cold reset
func (r *RegisterTebr1Type) GetSwdef() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTebr1FieldSwdefMask) >> RegisterTebr1FieldSwdefShift)
}

// SetSwdef Error detection logic sets this field to 0b00 on a new allocation and on Cold reset
func (r *RegisterTebr1Type) SetSwdef(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTebr1FieldSwdefMask)|(uint32(value)<<RegisterTebr1FieldSwdefShift))
}
