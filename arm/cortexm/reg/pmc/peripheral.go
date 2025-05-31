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
	Cpdlpstate registerCpdlpstateType
	Dpdlpstate registerDpdlpstateType
}

// registerCpdlpstateType Specifies the desired low-power states for core (PDCORE), Extension Processing Unit (PDEPU), and RAM (PDRAMS) power domains
type registerCpdlpstateType uint32

const (
	RegisterCpdlpstateFieldClpstateShift = 0
	RegisterCpdlpstateFieldClpstateMask  = 0x3
)

// GetClpstate Type of low-power state for PDCORE
func (r *registerCpdlpstateType) GetClpstate() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpdlpstateFieldClpstateMask) >> RegisterCpdlpstateFieldClpstateShift)
}

// SetClpstate Type of low-power state for PDCORE
func (r *registerCpdlpstateType) SetClpstate(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpdlpstateFieldClpstateMask)|(uint32(value)<<RegisterCpdlpstateFieldClpstateShift))
}

const (
	RegisterCpdlpstateFieldElpstateShift = 4
	RegisterCpdlpstateFieldElpstateMask  = 0x30
)

// GetElpstate Type of low-power state for PDEPU
func (r *registerCpdlpstateType) GetElpstate() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpdlpstateFieldElpstateMask) >> RegisterCpdlpstateFieldElpstateShift)
}

// SetElpstate Type of low-power state for PDEPU
func (r *registerCpdlpstateType) SetElpstate(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpdlpstateFieldElpstateMask)|(uint32(value)<<RegisterCpdlpstateFieldElpstateShift))
}

const (
	RegisterCpdlpstateFieldRlpstateShift = 8
	RegisterCpdlpstateFieldRlpstateMask  = 0x300
)

// GetRlpstate Power-on state for PDRAMS power domain
func (r *registerCpdlpstateType) GetRlpstate() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCpdlpstateFieldRlpstateMask) >> RegisterCpdlpstateFieldRlpstateShift)
}

// SetRlpstate Power-on state for PDRAMS power domain
func (r *registerCpdlpstateType) SetRlpstate(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpdlpstateFieldRlpstateMask)|(uint32(value)<<RegisterCpdlpstateFieldRlpstateShift))
}

// registerDpdlpstateType Specifies the desired low-power states for the debug (PDDEBUG) power domain
type registerDpdlpstateType uint32

const (
	RegisterDpdlpstateFieldDlpstateShift = 0
	RegisterDpdlpstateFieldDlpstateMask  = 0x3
)

// GetDlpstate Type of low-power state for PDDEBUG
func (r *registerDpdlpstateType) GetDlpstate() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDpdlpstateFieldDlpstateMask) >> RegisterDpdlpstateFieldDlpstateShift)
}

// SetDlpstate Type of low-power state for PDDEBUG
func (r *registerDpdlpstateType) SetDlpstate(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDpdlpstateFieldDlpstateMask)|(uint32(value)<<RegisterDpdlpstateFieldDlpstateShift))
}
