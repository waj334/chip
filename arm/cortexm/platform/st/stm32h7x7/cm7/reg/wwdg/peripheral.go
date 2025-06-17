//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

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
	Cr  RegisterCrType
	Cfr RegisterCfrType
	Sr  RegisterSrType
}

// RegisterCrType Control register
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
	RegisterCrFieldTShift = 0
	RegisterCrFieldTMask  = 0x7f
)

// GetT 7-bit counter (MSB to LSB) These bits contain the value of the watchdog counter. It is decremented every (4096 x 2WDGTB[1:0]) PCLK cycles. A reset is produced when it is decremented from 0x40 to 0x3F (T6 becomes cleared).
func (r *RegisterCrType) GetT() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTMask) >> RegisterCrFieldTShift)
}

// SetT 7-bit counter (MSB to LSB) These bits contain the value of the watchdog counter. It is decremented every (4096 x 2WDGTB[1:0]) PCLK cycles. A reset is produced when it is decremented from 0x40 to 0x3F (T6 becomes cleared).
func (r *RegisterCrType) SetT(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTMask)|(uint32(value)<<RegisterCrFieldTShift))
}

const (
	RegisterCrFieldWdgaShift = 7
	RegisterCrFieldWdgaMask  = 0x80
)

// GetWdga Activation bit This bit is set by software and only cleared by hardware after a reset. When WDGA=1, the watchdog can generate a reset.
func (r *RegisterCrType) GetWdga() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldWdgaMask) != 0
}

// SetWdga Activation bit This bit is set by software and only cleared by hardware after a reset. When WDGA=1, the watchdog can generate a reset.
func (r *RegisterCrType) SetWdga(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldWdgaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldWdgaMask)
	}
}

// RegisterCfrType Configuration register
type RegisterCfrType uint32

func (r *RegisterCfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCfrFieldWShift = 0
	RegisterCfrFieldWMask  = 0x7f
)

// GetW 7-bit window value These bits contain the window value to be compared to the downcounter.
func (r *RegisterCfrType) GetW() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfrFieldWMask) >> RegisterCfrFieldWShift)
}

// SetW 7-bit window value These bits contain the window value to be compared to the downcounter.
func (r *RegisterCfrType) SetW(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfrFieldWMask)|(uint32(value)<<RegisterCfrFieldWShift))
}

const (
	RegisterCfrFieldEwiShift = 9
	RegisterCfrFieldEwiMask  = 0x200
)

// GetEwi Early wakeup interrupt When set, an interrupt occurs whenever the counter reaches the value 0x40. This interrupt is only cleared by hardware after a reset.
func (r *RegisterCfrType) GetEwi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfrFieldEwiMask) != 0
}

// SetEwi Early wakeup interrupt When set, an interrupt occurs whenever the counter reaches the value 0x40. This interrupt is only cleared by hardware after a reset.
func (r *RegisterCfrType) SetEwi(value bool) {
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
func (r *RegisterCfrType) GetWdgtb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfrFieldWdgtbMask) >> RegisterCfrFieldWdgtbShift)
}

// SetWdgtb Timer base The time base of the prescaler can be modified as follows:
func (r *RegisterCfrType) SetWdgtb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfrFieldWdgtbMask)|(uint32(value)<<RegisterCfrFieldWdgtbShift))
}

// RegisterSrType Status register
type RegisterSrType uint32

func (r *RegisterSrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSrFieldEwifShift = 0
	RegisterSrFieldEwifMask  = 0x1
)

// GetEwif Early wakeup interrupt flag This bit is set by hardware when the counter has reached the value 0x40. It must be cleared by software by writing 0. A write of 1 has no effect. This bit is also set if the interrupt is not enabled.
func (r *RegisterSrType) GetEwif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldEwifMask) != 0
}

// SetEwif Early wakeup interrupt flag This bit is set by hardware when the counter has reached the value 0x40. It must be cleared by software by writing 0. A write of 1 has no effect. This bit is also set if the interrupt is not enabled.
func (r *RegisterSrType) SetEwif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldEwifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldEwifMask)
	}
}
