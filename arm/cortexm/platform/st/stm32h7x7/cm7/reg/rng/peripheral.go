//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package rng

import (
	"unsafe"
	"volatile"
)

var (
	Rng = (*_rng)(unsafe.Pointer(uintptr(0x48021800)))
)

type _rng struct {
	Cr registerCrType
	Sr registerSrType
	Dr registerDrType
}

// registerCrType RNG control register
type registerCrType uint32

const (
	RegisterCrFieldRngenShift = 2
	RegisterCrFieldRngenMask  = 0x4
)

// GetRngen Random number generator enable
func (r *registerCrType) GetRngen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRngenMask) != 0
}

// SetRngen Random number generator enable
func (r *registerCrType) SetRngen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRngenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRngenMask)
	}
}

const (
	RegisterCrFieldIeShift = 3
	RegisterCrFieldIeMask  = 0x8
)

// GetIe Interrupt enable
func (r *registerCrType) GetIe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldIeMask) != 0
}

// SetIe Interrupt enable
func (r *registerCrType) SetIe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldIeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldIeMask)
	}
}

const (
	RegisterCrFieldCedShift = 5
	RegisterCrFieldCedMask  = 0x20
)

// GetCed Clock error detection Note: The clock error detection can be used only when ck_rc48 or ck_pll1_q (ck_pll1_q = 48MHz) source is selected otherwise, CED bit must be equal to 1. The clock error detection cannot be enabled nor disabled on the fly when RNG peripheral is enabled, to enable or disable CED the RNG must be disabled.
func (r *registerCrType) GetCed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCedMask) != 0
}

// SetCed Clock error detection Note: The clock error detection can be used only when ck_rc48 or ck_pll1_q (ck_pll1_q = 48MHz) source is selected otherwise, CED bit must be equal to 1. The clock error detection cannot be enabled nor disabled on the fly when RNG peripheral is enabled, to enable or disable CED the RNG must be disabled.
func (r *registerCrType) SetCed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCedMask)
	}
}

// registerSrType RNG status register
type registerSrType uint32

const (
	RegisterSrFieldDrdyShift = 0
	RegisterSrFieldDrdyMask  = 0x1
)

// GetDrdy Data ready Note: If IE=1 in RNG_CR, an interrupt is generated when DRDY=1. It can rise when the peripheral is disabled. When the output buffer becomes empty (after reading RNG_DR), this bit returns to 0 until a new random value is generated.
func (r *registerSrType) GetDrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDrdyMask) != 0
}

const (
	RegisterSrFieldCecsShift = 1
	RegisterSrFieldCecsMask  = 0x2
)

// GetCecs Clock error current status Note: This bit is meaningless if CED (Clock error detection) bit in RNG_CR is equal to 1.
func (r *registerSrType) GetCecs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCecsMask) != 0
}

const (
	RegisterSrFieldSecsShift = 2
	RegisterSrFieldSecsMask  = 0x4
)

// GetSecs Seed error current status ** More than 64 consecutive bits at the same value (0 or 1) ** More than 32 consecutive alternances of 0 and 1 (0101010101...01)
func (r *registerSrType) GetSecs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSecsMask) != 0
}

const (
	RegisterSrFieldCeisShift = 5
	RegisterSrFieldCeisMask  = 0x20
)

// GetCeis Clock error interrupt status This bit is set at the same time as CECS. It is cleared by writing it to 0. An interrupt is pending if IE = 1 in the RNG_CR register. Note: This bit is meaningless if CED (Clock error detection) bit in RNG_CR is equal to 1.
func (r *registerSrType) GetCeis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCeisMask) != 0
}

// SetCeis Clock error interrupt status This bit is set at the same time as CECS. It is cleared by writing it to 0. An interrupt is pending if IE = 1 in the RNG_CR register. Note: This bit is meaningless if CED (Clock error detection) bit in RNG_CR is equal to 1.
func (r *registerSrType) SetCeis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCeisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCeisMask)
	}
}

const (
	RegisterSrFieldSeisShift = 6
	RegisterSrFieldSeisMask  = 0x40
)

// GetSeis Seed error interrupt status This bit is set at the same time as SECS. It is cleared by writing it to 0. ** More than 64 consecutive bits at the same value (0 or 1) ** More than 32 consecutive alternances of 0 and 1 (0101010101...01) An interrupt is pending if IE = 1 in the RNG_CR register.
func (r *registerSrType) GetSeis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSeisMask) != 0
}

// SetSeis Seed error interrupt status This bit is set at the same time as SECS. It is cleared by writing it to 0. ** More than 64 consecutive bits at the same value (0 or 1) ** More than 32 consecutive alternances of 0 and 1 (0101010101...01) An interrupt is pending if IE = 1 in the RNG_CR register.
func (r *registerSrType) SetSeis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldSeisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldSeisMask)
	}
}

// registerDrType The RNG_DR register is a read-only register that delivers a 32-bit random value when read. The content of this register is valid when DRDY= 1, even if RNGEN=0.
type registerDrType uint32

const (
	RegisterDrFieldRndataShift = 0
	RegisterDrFieldRndataMask  = 0xffffffff
)

// GetRndata Random data 32-bit random data which are valid when DRDY=1.
func (r *registerDrType) GetRndata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldRndataMask) >> RegisterDrFieldRndataShift)
}

// SetRndata Random data 32-bit random data which are valid when DRDY=1.
func (r *registerDrType) SetRndata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldRndataMask)|(uint32(value)<<RegisterDrFieldRndataShift))
}
