//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package art

import (
	"unsafe"
	"volatile"
)

var (
	Art = (*_art)(unsafe.Pointer(uintptr(0x40024400)))
)

type _art struct {
	Ctr RegisterCtrType
}

// RegisterCtrType control register
type RegisterCtrType uint32

func (r *RegisterCtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCtrFieldEnShift = 0
	RegisterCtrFieldEnMask  = 0x1
)

// GetEn Cache enable
func (r *RegisterCtrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldEnMask) != 0
}

// SetEn Cache enable
func (r *RegisterCtrType) SetEn(value bool) {
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
func (r *RegisterCtrType) GetPcacheaddr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldPcacheaddrMask) >> RegisterCtrFieldPcacheaddrShift)
}

// SetPcacheaddr Cacheable page index
func (r *RegisterCtrType) SetPcacheaddr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldPcacheaddrMask)|(uint32(value)<<RegisterCtrFieldPcacheaddrShift))
}
