package nvic_stir

import (
	"unsafe"
	"volatile"
)

var (
	Nvic_stir = (*_nvic_stir)(unsafe.Pointer(uintptr(0xe000ef00)))
)

type _nvic_stir struct {
	Stir registerStirType
}

// registerStirType Software trigger interrupt register
type registerStirType uint32

const (
	RegisterStirFieldIntidShift = 0
	RegisterStirFieldIntidMask  = 0x1ff
)

// GetIntid Software generated interrupt ID
func (r *registerStirType) GetIntid() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterStirFieldIntidMask) >> RegisterStirFieldIntidShift)
}

// SetIntid Software generated interrupt ID
func (r *registerStirType) SetIntid(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterStirFieldIntidMask)|(uint32(value)<<RegisterStirFieldIntidShift))
}
