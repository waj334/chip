//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package delay_block_sdmmc1

import (
	"unsafe"
	"volatile"
)

var (
	Delay_block_quadspi = (*_delay_block_sdmmc1)(unsafe.Pointer(uintptr(0x52006000)))
	Delay_block_sdmmc1  = (*_delay_block_sdmmc1)(unsafe.Pointer(uintptr(0x52008000)))
	Delay_block_sdmmc2  = (*_delay_block_sdmmc1)(unsafe.Pointer(uintptr(0x48022800)))
)

type _delay_block_sdmmc1 struct {
	Cr   registerCrType
	Cfgr registerCfgrType
}

// registerCrType DLYB control register
type registerCrType uint32

const (
	RegisterCrFieldDenShift = 0
	RegisterCrFieldDenMask  = 0x1
)

// GetDen Delay block enable bit
func (r *registerCrType) GetDen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDenMask) != 0
}

// SetDen Delay block enable bit
func (r *registerCrType) SetDen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDenMask)
	}
}

const (
	RegisterCrFieldSenShift = 1
	RegisterCrFieldSenMask  = 0x2
)

// GetSen Sampler length enable bit
func (r *registerCrType) GetSen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSenMask) != 0
}

// SetSen Sampler length enable bit
func (r *registerCrType) SetSen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldSenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSenMask)
	}
}

// registerCfgrType DLYB configuration register
type registerCfgrType uint32

const (
	RegisterCfgrFieldSelShift = 0
	RegisterCfgrFieldSelMask  = 0xf
)

// GetSel Select the phase for the Output clock
func (r *registerCfgrType) GetSel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSelMask) >> RegisterCfgrFieldSelShift)
}

// SetSel Select the phase for the Output clock
func (r *registerCfgrType) SetSel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSelMask)|(uint32(value)<<RegisterCfgrFieldSelShift))
}

const (
	RegisterCfgrFieldUnitShift = 8
	RegisterCfgrFieldUnitMask  = 0x7f00
)

// GetUnit Delay Defines the delay of a Unit delay cell
func (r *registerCfgrType) GetUnit() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldUnitMask) >> RegisterCfgrFieldUnitShift)
}

// SetUnit Delay Defines the delay of a Unit delay cell
func (r *registerCfgrType) SetUnit(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldUnitMask)|(uint32(value)<<RegisterCfgrFieldUnitShift))
}

const (
	RegisterCfgrFieldLngShift = 16
	RegisterCfgrFieldLngMask  = 0xfff0000
)

// GetLng Delay line length value
func (r *registerCfgrType) GetLng() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldLngMask) >> RegisterCfgrFieldLngShift)
}

// SetLng Delay line length value
func (r *registerCfgrType) SetLng(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldLngMask)|(uint32(value)<<RegisterCfgrFieldLngShift))
}

const (
	RegisterCfgrFieldLngfShift = 31
	RegisterCfgrFieldLngfMask  = 0x80000000
)

// GetLngf Length valid flag
func (r *registerCfgrType) GetLngf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldLngfMask) != 0
}

// SetLngf Length valid flag
func (r *registerCfgrType) SetLngf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldLngfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldLngfMask)
	}
}
