//go:build arm && cortexm

package pmc

import (
	"unsafe"
	"volatile"
)

var (
	Pmc = (*_pmc)(unsafe.Pointer(uintptr(0xe001e300)))
)

type _pmc struct {
	Cpdlpstate RegisterCpdlpstateType
	Dpdlpstate RegisterDpdlpstateType
}

// RegisterCpdlpstateType Specifies the desired low-power states for core (PDCORE), Extension Processing Unit (PDEPU), and RAM (PDRAMS) power domains
type RegisterCpdlpstateType uint32

func (r *RegisterCpdlpstateType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCpdlpstateType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCpdlpstateType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCpdlpstateType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCpdlpstateType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCpdlpstateFieldClpstateShift = 0
	RegisterCpdlpstateFieldClpstateMask  = 0x3
)

// GetClpstate Type of low-power state for PDCORE
func (r *RegisterCpdlpstateType) GetClpstate() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpdlpstateFieldClpstateMask) >> RegisterCpdlpstateFieldClpstateShift)
}

// SetClpstate Type of low-power state for PDCORE
func (r *RegisterCpdlpstateType) SetClpstate(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpdlpstateFieldClpstateMask)|(uint32(value)<<RegisterCpdlpstateFieldClpstateShift))
}

const (
	RegisterCpdlpstateFieldElpstateShift = 4
	RegisterCpdlpstateFieldElpstateMask  = 0x30
)

// GetElpstate Type of low-power state for PDEPU
func (r *RegisterCpdlpstateType) GetElpstate() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpdlpstateFieldElpstateMask) >> RegisterCpdlpstateFieldElpstateShift)
}

// SetElpstate Type of low-power state for PDEPU
func (r *RegisterCpdlpstateType) SetElpstate(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpdlpstateFieldElpstateMask)|(uint32(value)<<RegisterCpdlpstateFieldElpstateShift))
}

const (
	RegisterCpdlpstateFieldRlpstateShift = 8
	RegisterCpdlpstateFieldRlpstateMask  = 0x300
)

// GetRlpstate Power-on state for PDRAMS power domain
func (r *RegisterCpdlpstateType) GetRlpstate() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpdlpstateFieldRlpstateMask) >> RegisterCpdlpstateFieldRlpstateShift)
}

// SetRlpstate Power-on state for PDRAMS power domain
func (r *RegisterCpdlpstateType) SetRlpstate(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpdlpstateFieldRlpstateMask)|(uint32(value)<<RegisterCpdlpstateFieldRlpstateShift))
}

// RegisterDpdlpstateType Specifies the desired low-power states for the debug (PDDEBUG) power domain
type RegisterDpdlpstateType uint32

func (r *RegisterDpdlpstateType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDpdlpstateType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDpdlpstateType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDpdlpstateType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDpdlpstateType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDpdlpstateFieldDlpstateShift = 0
	RegisterDpdlpstateFieldDlpstateMask  = 0x3
)

// GetDlpstate Type of low-power state for PDDEBUG
func (r *RegisterDpdlpstateType) GetDlpstate() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDpdlpstateFieldDlpstateMask) >> RegisterDpdlpstateFieldDlpstateShift)
}

// SetDlpstate Type of low-power state for PDDEBUG
func (r *RegisterDpdlpstateType) SetDlpstate(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDpdlpstateFieldDlpstateMask)|(uint32(value)<<RegisterDpdlpstateFieldDlpstateShift))
}
