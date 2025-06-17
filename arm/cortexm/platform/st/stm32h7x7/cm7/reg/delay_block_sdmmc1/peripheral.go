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
	Cr   RegisterCrType
	Cfgr RegisterCfgrType
}

// RegisterCrType DLYB control register
type RegisterCrType uint32

func (r *RegisterCrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrFieldDenShift = 0
	RegisterCrFieldDenMask  = 0x1
)

// GetDen Delay block enable bit
func (r *RegisterCrType) GetDen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDenMask) != 0
}

// SetDen Delay block enable bit
func (r *RegisterCrType) SetDen(value bool) {
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
func (r *RegisterCrType) GetSen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSenMask) != 0
}

// SetSen Sampler length enable bit
func (r *RegisterCrType) SetSen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldSenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSenMask)
	}
}

// RegisterCfgrType DLYB configuration register
type RegisterCfgrType uint32

func (r *RegisterCfgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCfgrFieldSelShift = 0
	RegisterCfgrFieldSelMask  = 0xf
)

// GetSel Select the phase for the Output clock
func (r *RegisterCfgrType) GetSel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSelMask) >> RegisterCfgrFieldSelShift)
}

// SetSel Select the phase for the Output clock
func (r *RegisterCfgrType) SetSel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSelMask)|(uint32(value)<<RegisterCfgrFieldSelShift))
}

const (
	RegisterCfgrFieldUnitShift = 8
	RegisterCfgrFieldUnitMask  = 0x7f00
)

// GetUnit Delay Defines the delay of a Unit delay cell
func (r *RegisterCfgrType) GetUnit() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldUnitMask) >> RegisterCfgrFieldUnitShift)
}

// SetUnit Delay Defines the delay of a Unit delay cell
func (r *RegisterCfgrType) SetUnit(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldUnitMask)|(uint32(value)<<RegisterCfgrFieldUnitShift))
}

const (
	RegisterCfgrFieldLngShift = 16
	RegisterCfgrFieldLngMask  = 0xfff0000
)

// GetLng Delay line length value
func (r *RegisterCfgrType) GetLng() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldLngMask) >> RegisterCfgrFieldLngShift)
}

// SetLng Delay line length value
func (r *RegisterCfgrType) SetLng(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldLngMask)|(uint32(value)<<RegisterCfgrFieldLngShift))
}

const (
	RegisterCfgrFieldLngfShift = 31
	RegisterCfgrFieldLngfMask  = 0x80000000
)

// GetLngf Length valid flag
func (r *RegisterCfgrType) GetLngf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldLngfMask) != 0
}

// SetLngf Length valid flag
func (r *RegisterCfgrType) SetLngf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldLngfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldLngfMask)
	}
}
