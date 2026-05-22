//go:build arm && cortexm

package pci

import (
	"unsafe"
	"volatile"
)

var (
	Pci = (*_pci)(unsafe.Pointer(uintptr(0xe001e700)))
)

type _pci struct {
	Cfginfosel RegisterCfginfoselType
	Cfginford  RegisterCfginfordType
}

// RegisterCfginfoselType Selects the configuration information which can then be read back using CFGINFORD
type RegisterCfginfoselType uint32

func (r *RegisterCfginfoselType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfginfoselType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfginfoselType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfginfoselType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfginfoselType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterCfginfordType Displays the configuration information that the CFGINFOSEL register selects
type RegisterCfginfordType uint32

func (r *RegisterCfginfordType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfginfordType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfginfordType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfginfordType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfginfordType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}
