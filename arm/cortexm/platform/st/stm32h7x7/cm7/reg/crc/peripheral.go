//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package crc

import (
	"unsafe"
	"volatile"
)

var (
	Crc = (*_crc)(unsafe.Pointer(uintptr(0x58024c00)))
)

type _crc struct {
	Dr   RegisterDrType
	Idr  RegisterIdrType
	Cr   RegisterCrType
	_    [4]byte
	Init RegisterInitType
	Pol  RegisterPolType
}

// RegisterDrType Data register
type RegisterDrType uint32

func (r *RegisterDrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDrFieldDrShift = 0
	RegisterDrFieldDrMask  = 0xffffffff
)

// GetDr Data Register
func (r *RegisterDrType) GetDr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldDrMask) >> RegisterDrFieldDrShift)
}

// SetDr Data Register
func (r *RegisterDrType) SetDr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldDrMask)|(uint32(value)<<RegisterDrFieldDrShift))
}

// RegisterIdrType Independent Data register
type RegisterIdrType uint32

func (r *RegisterIdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIdrFieldIdrShift = 0
	RegisterIdrFieldIdrMask  = 0xffffffff
)

// GetIdr Independent Data register
func (r *RegisterIdrType) GetIdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldIdrMask) >> RegisterIdrFieldIdrShift)
}

// SetIdr Independent Data register
func (r *RegisterIdrType) SetIdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldIdrMask)|(uint32(value)<<RegisterIdrFieldIdrShift))
}

// RegisterCrType Control register
type RegisterCrType uint32

func (r *RegisterCrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrFieldResetShift = 0
	RegisterCrFieldResetMask  = 0x1
)

// GetReset RESET bit
func (r *RegisterCrType) GetReset() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldResetMask) != 0
}

// SetReset RESET bit
func (r *RegisterCrType) SetReset(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldResetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldResetMask)
	}
}

const (
	RegisterCrFieldPolysizeShift = 3
	RegisterCrFieldPolysizeMask  = 0x18
)

// GetPolysize Polynomial size
func (r *RegisterCrType) GetPolysize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPolysizeMask) >> RegisterCrFieldPolysizeShift)
}

// SetPolysize Polynomial size
func (r *RegisterCrType) SetPolysize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPolysizeMask)|(uint32(value)<<RegisterCrFieldPolysizeShift))
}

const (
	RegisterCrFieldRevinShift = 5
	RegisterCrFieldRevinMask  = 0x60
)

// GetRevin Reverse input data
func (r *RegisterCrType) GetRevin() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRevinMask) >> RegisterCrFieldRevinShift)
}

// SetRevin Reverse input data
func (r *RegisterCrType) SetRevin(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRevinMask)|(uint32(value)<<RegisterCrFieldRevinShift))
}

const (
	RegisterCrFieldRevoutShift = 7
	RegisterCrFieldRevoutMask  = 0x80
)

// GetRevout Reverse output data
func (r *RegisterCrType) GetRevout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRevoutMask) != 0
}

// SetRevout Reverse output data
func (r *RegisterCrType) SetRevout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRevoutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRevoutMask)
	}
}

// RegisterInitType Initial CRC value
type RegisterInitType uint32

func (r *RegisterInitType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterInitType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterInitType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterInitType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterInitType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterInitFieldCrcinitShift = 0
	RegisterInitFieldCrcinitMask  = 0xffffffff
)

// GetCrcinit Programmable initial CRC value
func (r *RegisterInitType) GetCrcinit() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterInitFieldCrcinitMask) >> RegisterInitFieldCrcinitShift)
}

// SetCrcinit Programmable initial CRC value
func (r *RegisterInitType) SetCrcinit(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterInitFieldCrcinitMask)|(uint32(value)<<RegisterInitFieldCrcinitShift))
}

// RegisterPolType CRC polynomial
type RegisterPolType uint32

func (r *RegisterPolType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPolType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPolType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPolType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPolType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPolFieldPolShift = 0
	RegisterPolFieldPolMask  = 0xffffffff
)

// GetPol Programmable polynomial
func (r *RegisterPolType) GetPol() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterPolFieldPolMask) >> RegisterPolFieldPolShift)
}

// SetPol Programmable polynomial
func (r *RegisterPolType) SetPol(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPolFieldPolMask)|(uint32(value)<<RegisterPolFieldPolShift))
}
