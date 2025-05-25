package art

import (
	"unsafe"
	"volatile"
)

var (
	Art = (*_art)(unsafe.Pointer(uintptr(0x40024400)))
)

type _art struct {
	Ctr registerCtrType
}

// registerCtrType control register
type registerCtrType uint32

const (
	RegisterCtrFieldEnShift = 0
	RegisterCtrFieldEnMask  = 0x1
)

// GetEn Cache enable
func (r *registerCtrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldEnMask) != 0
}

// SetEn Cache enable
func (r *registerCtrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldEnMask)
	}
}

const (
	RegisterCtrFieldPcacheaddrShift = 8
	RegisterCtrFieldPcacheaddrMask  = 0xfff00
)

// GetPcacheaddr Cacheable page index
func (r *registerCtrType) GetPcacheaddr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldPcacheaddrMask) >> RegisterCtrFieldPcacheaddrShift)
}

// SetPcacheaddr Cacheable page index
func (r *registerCtrType) SetPcacheaddr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldPcacheaddrMask)|(uint32(value)<<RegisterCtrFieldPcacheaddrShift))
}
