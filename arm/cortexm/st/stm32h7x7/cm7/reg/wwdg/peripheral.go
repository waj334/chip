package wwdg

import (
	"unsafe"
	"volatile"
)

var (
	Wwdg1 = (*_wwdg)(unsafe.Pointer(uintptr(0x50003000)))
	Wwdg2 = (*_wwdg)(unsafe.Pointer(uintptr(0x40002c00)))

	Instances = [2]*_wwdg{
		Wwdg1,
		Wwdg2,
	}
)

type _wwdg struct {
	Cr  registerCrType
	Cfr registerCfrType
	Sr  registerSrType
}

// registerCrType Control register
type registerCrType uint32

const (
	RegisterCrFieldTShift = 0
	RegisterCrFieldTMask  = 0x7f
)

// GetT 7-bit counter (MSB to LSB) These bits contain the value of the watchdog counter. It is decremented every (4096 x 2WDGTB[1:0]) PCLK cycles. A reset is produced when it is decremented from 0x40 to 0x3F (T6 becomes cleared).
func (r *registerCrType) GetT() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTMask) >> RegisterCrFieldTShift)
}

// SetT 7-bit counter (MSB to LSB) These bits contain the value of the watchdog counter. It is decremented every (4096 x 2WDGTB[1:0]) PCLK cycles. A reset is produced when it is decremented from 0x40 to 0x3F (T6 becomes cleared).
func (r *registerCrType) SetT(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTMask)|(uint32(value)<<RegisterCrFieldTShift))
}

const (
	RegisterCrFieldWdgaShift = 7
	RegisterCrFieldWdgaMask  = 0x80
)

// GetWdga Activation bit This bit is set by software and only cleared by hardware after a reset. When WDGA=1, the watchdog can generate a reset.
func (r *registerCrType) GetWdga() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldWdgaMask) != 0
}

// SetWdga Activation bit This bit is set by software and only cleared by hardware after a reset. When WDGA=1, the watchdog can generate a reset.
func (r *registerCrType) SetWdga(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldWdgaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldWdgaMask)
	}
}

// registerCfrType Configuration register
type registerCfrType uint32

const (
	RegisterCfrFieldWShift = 0
	RegisterCfrFieldWMask  = 0x7f
)

// GetW 7-bit window value These bits contain the window value to be compared to the downcounter.
func (r *registerCfrType) GetW() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfrFieldWMask) >> RegisterCfrFieldWShift)
}

// SetW 7-bit window value These bits contain the window value to be compared to the downcounter.
func (r *registerCfrType) SetW(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfrFieldWMask)|(uint32(value)<<RegisterCfrFieldWShift))
}

const (
	RegisterCfrFieldEwiShift = 9
	RegisterCfrFieldEwiMask  = 0x200
)

// GetEwi Early wakeup interrupt When set, an interrupt occurs whenever the counter reaches the value 0x40. This interrupt is only cleared by hardware after a reset.
func (r *registerCfrType) GetEwi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfrFieldEwiMask) != 0
}

// SetEwi Early wakeup interrupt When set, an interrupt occurs whenever the counter reaches the value 0x40. This interrupt is only cleared by hardware after a reset.
func (r *registerCfrType) SetEwi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfrFieldEwiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfrFieldEwiMask)
	}
}

const (
	RegisterCfrFieldWdgtbShift = 11
	RegisterCfrFieldWdgtbMask  = 0x1800
)

// GetWdgtb Timer base The time base of the prescaler can be modified as follows:
func (r *registerCfrType) GetWdgtb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfrFieldWdgtbMask) >> RegisterCfrFieldWdgtbShift)
}

// SetWdgtb Timer base The time base of the prescaler can be modified as follows:
func (r *registerCfrType) SetWdgtb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfrFieldWdgtbMask)|(uint32(value)<<RegisterCfrFieldWdgtbShift))
}

// registerSrType Status register
type registerSrType uint32

const (
	RegisterSrFieldEwifShift = 0
	RegisterSrFieldEwifMask  = 0x1
)

// GetEwif Early wakeup interrupt flag This bit is set by hardware when the counter has reached the value 0x40. It must be cleared by software by writing 0. A write of 1 has no effect. This bit is also set if the interrupt is not enabled.
func (r *registerSrType) GetEwif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldEwifMask) != 0
}

// SetEwif Early wakeup interrupt flag This bit is set by hardware when the counter has reached the value 0x40. It must be cleared by software by writing 0. A write of 1 has no effect. This bit is also set if the interrupt is not enabled.
func (r *registerSrType) SetEwif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldEwifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldEwifMask)
	}
}
