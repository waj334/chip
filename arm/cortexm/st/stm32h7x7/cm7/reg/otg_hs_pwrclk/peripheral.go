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
	Otg_hs_pcgcr registerOtg_hs_pcgcrType
}

// registerOtg_hs_pcgcrType Power and clock gating control register
type registerOtg_hs_pcgcrType uint32

const (
	RegisterOtg_hs_pcgcrFieldStppclkShift = 0
	RegisterOtg_hs_pcgcrFieldStppclkMask  = 0x1
)

// GetStppclk Stop PHY clock
func (r *registerOtg_hs_pcgcrType) GetStppclk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_pcgcrFieldStppclkMask) != 0
}

// SetStppclk Stop PHY clock
func (r *registerOtg_hs_pcgcrType) SetStppclk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_pcgcrFieldStppclkMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_pcgcrFieldStppclkMask)
	}
}

const (
	RegisterOtg_hs_pcgcrFieldGatehclkShift = 1
	RegisterOtg_hs_pcgcrFieldGatehclkMask  = 0x2
)

// GetGatehclk Gate HCLK
func (r *registerOtg_hs_pcgcrType) GetGatehclk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_pcgcrFieldGatehclkMask) != 0
}

// SetGatehclk Gate HCLK
func (r *registerOtg_hs_pcgcrType) SetGatehclk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_pcgcrFieldGatehclkMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_pcgcrFieldGatehclkMask)
	}
}

const (
	RegisterOtg_hs_pcgcrFieldPhysuspShift = 4
	RegisterOtg_hs_pcgcrFieldPhysuspMask  = 0x10
)

// GetPhysusp PHY suspended
func (r *registerOtg_hs_pcgcrType) GetPhysusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_pcgcrFieldPhysuspMask) != 0
}

// SetPhysusp PHY suspended
func (r *registerOtg_hs_pcgcrType) SetPhysusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_pcgcrFieldPhysuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_pcgcrFieldPhysuspMask)
	}
}
