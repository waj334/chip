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
	Cfginfosel registerCfginfoselType
	Cfginford  registerCfginfordType
}

// registerCfginfoselType Selects the configuration information which can then be read back using CFGINFORD
type registerCfginfoselType uint32

// registerCfginfordType Displays the configuration information that the CFGINFOSEL register selects
type registerCfginfordType uint32
