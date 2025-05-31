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
	Dr   registerDrType
	Idr  registerIdrType
	Cr   registerCrType
	_    [4]byte
	Init registerInitType
	Pol  registerPolType
}

// registerDrType Data register
type registerDrType uint32

const (
	RegisterDrFieldDrShift = 0
	RegisterDrFieldDrMask  = 0xffffffff
)

// GetDr Data Register
func (r *registerDrType) GetDr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldDrMask) >> RegisterDrFieldDrShift)
}

// SetDr Data Register
func (r *registerDrType) SetDr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldDrMask)|(uint32(value)<<RegisterDrFieldDrShift))
}

// registerIdrType Independent Data register
type registerIdrType uint32

const (
	RegisterIdrFieldIdrShift = 0
	RegisterIdrFieldIdrMask  = 0xffffffff
)

// GetIdr Independent Data register
func (r *registerIdrType) GetIdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldIdrMask) >> RegisterIdrFieldIdrShift)
}

// SetIdr Independent Data register
func (r *registerIdrType) SetIdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldIdrMask)|(uint32(value)<<RegisterIdrFieldIdrShift))
}

// registerCrType Control register
type registerCrType uint32

const (
	RegisterCrFieldResetShift = 0
	RegisterCrFieldResetMask  = 0x1
)

// GetReset RESET bit
func (r *registerCrType) GetReset() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldResetMask) != 0
}

// SetReset RESET bit
func (r *registerCrType) SetReset(value bool) {
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
func (r *registerCrType) GetPolysize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPolysizeMask) >> RegisterCrFieldPolysizeShift)
}

// SetPolysize Polynomial size
func (r *registerCrType) SetPolysize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPolysizeMask)|(uint32(value)<<RegisterCrFieldPolysizeShift))
}

const (
	RegisterCrFieldRevinShift = 5
	RegisterCrFieldRevinMask  = 0x60
)

// GetRevin Reverse input data
func (r *registerCrType) GetRevin() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRevinMask) >> RegisterCrFieldRevinShift)
}

// SetRevin Reverse input data
func (r *registerCrType) SetRevin(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRevinMask)|(uint32(value)<<RegisterCrFieldRevinShift))
}

const (
	RegisterCrFieldRevoutShift = 7
	RegisterCrFieldRevoutMask  = 0x80
)

// GetRevout Reverse output data
func (r *registerCrType) GetRevout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRevoutMask) != 0
}

// SetRevout Reverse output data
func (r *registerCrType) SetRevout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRevoutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRevoutMask)
	}
}

// registerInitType Initial CRC value
type registerInitType uint32

const (
	RegisterInitFieldCrcinitShift = 0
	RegisterInitFieldCrcinitMask  = 0xffffffff
)

// GetCrcinit Programmable initial CRC value
func (r *registerInitType) GetCrcinit() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterInitFieldCrcinitMask) >> RegisterInitFieldCrcinitShift)
}

// SetCrcinit Programmable initial CRC value
func (r *registerInitType) SetCrcinit(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterInitFieldCrcinitMask)|(uint32(value)<<RegisterInitFieldCrcinitShift))
}

// registerPolType CRC polynomial
type registerPolType uint32

const (
	RegisterPolFieldPolShift = 0
	RegisterPolFieldPolMask  = 0xffffffff
)

// GetPol Programmable polynomial
func (r *registerPolType) GetPol() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterPolFieldPolMask) >> RegisterPolFieldPolShift)
}

// SetPol Programmable polynomial
func (r *registerPolType) SetPol(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPolFieldPolMask)|(uint32(value)<<RegisterPolFieldPolShift))
}
