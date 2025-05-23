package fpu_cpacr

import (
	"unsafe"
	"volatile"
)

var (
	Fpu_cpacr = (*_fpu_cpacr)(unsafe.Pointer(uintptr(0xe000ed88)))
)

type _fpu_cpacr struct {
	Cpacr registerCpacrType
}

// registerCpacrType Coprocessor access control register
type registerCpacrType uint32

const (
	RegisterCpacrFieldCpShift = 20
	RegisterCpacrFieldCpMask  = 0xf00000
)

// GetCp CP
func (r *registerCpacrType) GetCp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpacrFieldCpMask) >> RegisterCpacrFieldCpShift)
}

// SetCp CP
func (r *registerCpacrType) SetCp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpacrFieldCpMask)|(uint32(value)<<RegisterCpacrFieldCpShift))
}
