//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package otg_hs_pwrclk

import (
	"unsafe"
	"volatile"
)

var (
	Otg1_hs_pwrclk = (*_otg_hs_pwrclk)(unsafe.Pointer(uintptr(0x40040e00)))
	Otg2_hs_pwrclk = (*_otg_hs_pwrclk)(unsafe.Pointer(uintptr(0x40080e00)))

	Instances = [2]*_otg_hs_pwrclk{
		Otg1_hs_pwrclk,
		Otg2_hs_pwrclk,
	}
)

type _otg_hs_pwrclk struct {
	Otghspcgcr RegisterOtghspcgcrType
}

// RegisterOtghspcgcrType Power and clock gating control register
type RegisterOtghspcgcrType uint32

func (r *RegisterOtghspcgcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghspcgcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghspcgcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghspcgcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghspcgcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghspcgcrFieldStppclkShift = 0
	RegisterOtghspcgcrFieldStppclkMask  = 0x1
)

// GetStppclk Stop PHY clock
func (r *RegisterOtghspcgcrType) GetStppclk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghspcgcrFieldStppclkMask) != 0
}

// SetStppclk Stop PHY clock
func (r *RegisterOtghspcgcrType) SetStppclk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghspcgcrFieldStppclkMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghspcgcrFieldStppclkMask)
	}
}

const (
	RegisterOtghspcgcrFieldGatehclkShift = 1
	RegisterOtghspcgcrFieldGatehclkMask  = 0x2
)

// GetGatehclk Gate HCLK
func (r *RegisterOtghspcgcrType) GetGatehclk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghspcgcrFieldGatehclkMask) != 0
}

// SetGatehclk Gate HCLK
func (r *RegisterOtghspcgcrType) SetGatehclk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghspcgcrFieldGatehclkMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghspcgcrFieldGatehclkMask)
	}
}

const (
	RegisterOtghspcgcrFieldPhysuspShift = 4
	RegisterOtghspcgcrFieldPhysuspMask  = 0x10
)

// GetPhysusp PHY suspended
func (r *RegisterOtghspcgcrType) GetPhysusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghspcgcrFieldPhysuspMask) != 0
}

// SetPhysusp PHY suspended
func (r *RegisterOtghspcgcrType) SetPhysusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghspcgcrFieldPhysuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghspcgcrFieldPhysuspMask)
	}
}
