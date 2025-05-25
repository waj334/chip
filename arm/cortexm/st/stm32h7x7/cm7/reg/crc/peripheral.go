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
	RegisterCrFieldRev_inShift = 5
	RegisterCrFieldRev_inMask  = 0x60
)

// GetRev_in Reverse input data
func (r *registerCrType) GetRev_in() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRev_inMask) >> RegisterCrFieldRev_inShift)
}

// SetRev_in Reverse input data
func (r *registerCrType) SetRev_in(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRev_inMask)|(uint32(value)<<RegisterCrFieldRev_inShift))
}

const (
	RegisterCrFieldRev_outShift = 7
	RegisterCrFieldRev_outMask  = 0x80
)

// GetRev_out Reverse output data
func (r *registerCrType) GetRev_out() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRev_outMask) != 0
}

// SetRev_out Reverse output data
func (r *registerCrType) SetRev_out(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRev_outMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRev_outMask)
	}
}

// registerInitType Initial CRC value
type registerInitType uint32

const (
	RegisterInitFieldCrc_initShift = 0
	RegisterInitFieldCrc_initMask  = 0xffffffff
)

// GetCrc_init Programmable initial CRC value
func (r *registerInitType) GetCrc_init() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterInitFieldCrc_initMask) >> RegisterInitFieldCrc_initShift)
}

// SetCrc_init Programmable initial CRC value
func (r *registerInitType) SetCrc_init(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterInitFieldCrc_initMask)|(uint32(value)<<RegisterInitFieldCrc_initShift))
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
