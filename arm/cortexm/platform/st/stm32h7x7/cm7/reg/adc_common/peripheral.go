//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package adc_common

import (
	"unsafe"
	"volatile"
)

var (
	Adc3_common  = (*_adc_common)(unsafe.Pointer(uintptr(0x58026300)))
	Adc12_common = (*_adc_common)(unsafe.Pointer(uintptr(0x40022300)))

	Instances = [2]*_adc_common{
		Adc12_common,
		Adc3_common,
	}
)

type _adc_common struct {
	Csr  RegisterCsrType
	_    [4]byte
	Ccr  RegisterCcrType
	Cdr  RegisterCdrType
	Cdr2 RegisterCdr2Type
}

// RegisterCsrType ADC Common status register
type RegisterCsrType uint32

func (r *RegisterCsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsrFieldAdrdymstShift = 0
	RegisterCsrFieldAdrdymstMask  = 0x1
)

// GetAdrdymst Master ADC ready
func (r *RegisterCsrType) GetAdrdymst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAdrdymstMask) != 0
}

// SetAdrdymst Master ADC ready
func (r *RegisterCsrType) SetAdrdymst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAdrdymstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAdrdymstMask)
	}
}

const (
	RegisterCsrFieldEosmpmstShift = 1
	RegisterCsrFieldEosmpmstMask  = 0x2
)

// GetEosmpmst End of Sampling phase flag of the master ADC
func (r *RegisterCsrType) GetEosmpmst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEosmpmstMask) != 0
}

// SetEosmpmst End of Sampling phase flag of the master ADC
func (r *RegisterCsrType) SetEosmpmst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEosmpmstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEosmpmstMask)
	}
}

const (
	RegisterCsrFieldEocmstShift = 2
	RegisterCsrFieldEocmstMask  = 0x4
)

// GetEocmst End of regular conversion of the master ADC
func (r *RegisterCsrType) GetEocmst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEocmstMask) != 0
}

// SetEocmst End of regular conversion of the master ADC
func (r *RegisterCsrType) SetEocmst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEocmstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEocmstMask)
	}
}

const (
	RegisterCsrFieldEosmstShift = 3
	RegisterCsrFieldEosmstMask  = 0x8
)

// GetEosmst End of regular sequence flag of the master ADC
func (r *RegisterCsrType) GetEosmst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEosmstMask) != 0
}

// SetEosmst End of regular sequence flag of the master ADC
func (r *RegisterCsrType) SetEosmst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEosmstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEosmstMask)
	}
}

const (
	RegisterCsrFieldOvrmstShift = 4
	RegisterCsrFieldOvrmstMask  = 0x10
)

// GetOvrmst Overrun flag of the master ADC
func (r *RegisterCsrType) GetOvrmst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldOvrmstMask) != 0
}

// SetOvrmst Overrun flag of the master ADC
func (r *RegisterCsrType) SetOvrmst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldOvrmstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldOvrmstMask)
	}
}

const (
	RegisterCsrFieldJeocmstShift = 5
	RegisterCsrFieldJeocmstMask  = 0x20
)

// GetJeocmst End of injected conversion flag of the master ADC
func (r *RegisterCsrType) GetJeocmst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldJeocmstMask) != 0
}

// SetJeocmst End of injected conversion flag of the master ADC
func (r *RegisterCsrType) SetJeocmst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldJeocmstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldJeocmstMask)
	}
}

const (
	RegisterCsrFieldJeosmstShift = 6
	RegisterCsrFieldJeosmstMask  = 0x40
)

// GetJeosmst End of injected sequence flag of the master ADC
func (r *RegisterCsrType) GetJeosmst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldJeosmstMask) != 0
}

// SetJeosmst End of injected sequence flag of the master ADC
func (r *RegisterCsrType) SetJeosmst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldJeosmstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldJeosmstMask)
	}
}

const (
	RegisterCsrFieldAwd1mstShift = 7
	RegisterCsrFieldAwd1mstMask  = 0x80
)

// GetAwd1mst Analog watchdog 1 flag of the master ADC
func (r *RegisterCsrType) GetAwd1mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAwd1mstMask) != 0
}

// SetAwd1mst Analog watchdog 1 flag of the master ADC
func (r *RegisterCsrType) SetAwd1mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAwd1mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAwd1mstMask)
	}
}

const (
	RegisterCsrFieldAwd2mstShift = 8
	RegisterCsrFieldAwd2mstMask  = 0x100
)

// GetAwd2mst Analog watchdog 2 flag of the master ADC
func (r *RegisterCsrType) GetAwd2mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAwd2mstMask) != 0
}

// SetAwd2mst Analog watchdog 2 flag of the master ADC
func (r *RegisterCsrType) SetAwd2mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAwd2mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAwd2mstMask)
	}
}

const (
	RegisterCsrFieldAwd3mstShift = 9
	RegisterCsrFieldAwd3mstMask  = 0x200
)

// GetAwd3mst Analog watchdog 3 flag of the master ADC
func (r *RegisterCsrType) GetAwd3mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAwd3mstMask) != 0
}

// SetAwd3mst Analog watchdog 3 flag of the master ADC
func (r *RegisterCsrType) SetAwd3mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAwd3mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAwd3mstMask)
	}
}

const (
	RegisterCsrFieldJqovfmstShift = 10
	RegisterCsrFieldJqovfmstMask  = 0x400
)

// GetJqovfmst Injected Context Queue Overflow flag of the master ADC
func (r *RegisterCsrType) GetJqovfmst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldJqovfmstMask) != 0
}

// SetJqovfmst Injected Context Queue Overflow flag of the master ADC
func (r *RegisterCsrType) SetJqovfmst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldJqovfmstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldJqovfmstMask)
	}
}

const (
	RegisterCsrFieldAdrdyslvShift = 16
	RegisterCsrFieldAdrdyslvMask  = 0x10000
)

// GetAdrdyslv Slave ADC ready
func (r *RegisterCsrType) GetAdrdyslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAdrdyslvMask) != 0
}

// SetAdrdyslv Slave ADC ready
func (r *RegisterCsrType) SetAdrdyslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAdrdyslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAdrdyslvMask)
	}
}

const (
	RegisterCsrFieldEosmpslvShift = 17
	RegisterCsrFieldEosmpslvMask  = 0x20000
)

// GetEosmpslv End of Sampling phase flag of the slave ADC
func (r *RegisterCsrType) GetEosmpslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEosmpslvMask) != 0
}

// SetEosmpslv End of Sampling phase flag of the slave ADC
func (r *RegisterCsrType) SetEosmpslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEosmpslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEosmpslvMask)
	}
}

const (
	RegisterCsrFieldEocslvShift = 18
	RegisterCsrFieldEocslvMask  = 0x40000
)

// GetEocslv End of regular conversion of the slave ADC
func (r *RegisterCsrType) GetEocslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEocslvMask) != 0
}

// SetEocslv End of regular conversion of the slave ADC
func (r *RegisterCsrType) SetEocslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEocslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEocslvMask)
	}
}

const (
	RegisterCsrFieldEosslvShift = 19
	RegisterCsrFieldEosslvMask  = 0x80000
)

// GetEosslv End of regular sequence flag of the slave ADC
func (r *RegisterCsrType) GetEosslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEosslvMask) != 0
}

// SetEosslv End of regular sequence flag of the slave ADC
func (r *RegisterCsrType) SetEosslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEosslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEosslvMask)
	}
}

const (
	RegisterCsrFieldOvrslvShift = 20
	RegisterCsrFieldOvrslvMask  = 0x100000
)

// GetOvrslv Overrun flag of the slave ADC
func (r *RegisterCsrType) GetOvrslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldOvrslvMask) != 0
}

// SetOvrslv Overrun flag of the slave ADC
func (r *RegisterCsrType) SetOvrslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldOvrslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldOvrslvMask)
	}
}

const (
	RegisterCsrFieldJeocslvShift = 21
	RegisterCsrFieldJeocslvMask  = 0x200000
)

// GetJeocslv End of injected conversion flag of the slave ADC
func (r *RegisterCsrType) GetJeocslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldJeocslvMask) != 0
}

// SetJeocslv End of injected conversion flag of the slave ADC
func (r *RegisterCsrType) SetJeocslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldJeocslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldJeocslvMask)
	}
}

const (
	RegisterCsrFieldJeosslvShift = 22
	RegisterCsrFieldJeosslvMask  = 0x400000
)

// GetJeosslv End of injected sequence flag of the slave ADC
func (r *RegisterCsrType) GetJeosslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldJeosslvMask) != 0
}

// SetJeosslv End of injected sequence flag of the slave ADC
func (r *RegisterCsrType) SetJeosslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldJeosslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldJeosslvMask)
	}
}

const (
	RegisterCsrFieldAwd1slvShift = 23
	RegisterCsrFieldAwd1slvMask  = 0x800000
)

// GetAwd1slv Analog watchdog 1 flag of the slave ADC
func (r *RegisterCsrType) GetAwd1slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAwd1slvMask) != 0
}

// SetAwd1slv Analog watchdog 1 flag of the slave ADC
func (r *RegisterCsrType) SetAwd1slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAwd1slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAwd1slvMask)
	}
}

const (
	RegisterCsrFieldAwd2slvShift = 24
	RegisterCsrFieldAwd2slvMask  = 0x1000000
)

// GetAwd2slv Analog watchdog 2 flag of the slave ADC
func (r *RegisterCsrType) GetAwd2slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAwd2slvMask) != 0
}

// SetAwd2slv Analog watchdog 2 flag of the slave ADC
func (r *RegisterCsrType) SetAwd2slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAwd2slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAwd2slvMask)
	}
}

const (
	RegisterCsrFieldAwd3slvShift = 25
	RegisterCsrFieldAwd3slvMask  = 0x2000000
)

// GetAwd3slv Analog watchdog 3 flag of the slave ADC
func (r *RegisterCsrType) GetAwd3slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAwd3slvMask) != 0
}

// SetAwd3slv Analog watchdog 3 flag of the slave ADC
func (r *RegisterCsrType) SetAwd3slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAwd3slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAwd3slvMask)
	}
}

const (
	RegisterCsrFieldJqovfslvShift = 26
	RegisterCsrFieldJqovfslvMask  = 0x4000000
)

// GetJqovfslv Injected Context Queue Overflow flag of the slave ADC
func (r *RegisterCsrType) GetJqovfslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldJqovfslvMask) != 0
}

// SetJqovfslv Injected Context Queue Overflow flag of the slave ADC
func (r *RegisterCsrType) SetJqovfslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldJqovfslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldJqovfslvMask)
	}
}

// RegisterCcrType ADC common control register
type RegisterCcrType uint32

func (r *RegisterCcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcrFieldDualShift = 0
	RegisterCcrFieldDualMask  = 0x1f
)

// GetDual Dual ADC mode selection
func (r *RegisterCcrType) GetDual() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDualMask) >> RegisterCcrFieldDualShift)
}

// SetDual Dual ADC mode selection
func (r *RegisterCcrType) SetDual(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDualMask)|(uint32(value)<<RegisterCcrFieldDualShift))
}

const (
	RegisterCcrFieldDelayShift = 8
	RegisterCcrFieldDelayMask  = 0xf00
)

// GetDelay Delay between 2 sampling phases
func (r *RegisterCcrType) GetDelay() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDelayMask) >> RegisterCcrFieldDelayShift)
}

// SetDelay Delay between 2 sampling phases
func (r *RegisterCcrType) SetDelay(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDelayMask)|(uint32(value)<<RegisterCcrFieldDelayShift))
}

const (
	RegisterCcrFieldDamdfShift = 14
	RegisterCcrFieldDamdfMask  = 0xc000
)

// GetDamdf Dual ADC Mode Data Format
func (r *RegisterCcrType) GetDamdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDamdfMask) >> RegisterCcrFieldDamdfShift)
}

// SetDamdf Dual ADC Mode Data Format
func (r *RegisterCcrType) SetDamdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDamdfMask)|(uint32(value)<<RegisterCcrFieldDamdfShift))
}

const (
	RegisterCcrFieldCkmodeShift = 16
	RegisterCcrFieldCkmodeMask  = 0x30000
)

// GetCkmode ADC clock mode
func (r *RegisterCcrType) GetCkmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldCkmodeMask) >> RegisterCcrFieldCkmodeShift)
}

// SetCkmode ADC clock mode
func (r *RegisterCcrType) SetCkmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldCkmodeMask)|(uint32(value)<<RegisterCcrFieldCkmodeShift))
}

const (
	RegisterCcrFieldPrescShift = 18
	RegisterCcrFieldPrescMask  = 0x3c0000
)

// GetPresc ADC prescaler
func (r *RegisterCcrType) GetPresc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldPrescMask) >> RegisterCcrFieldPrescShift)
}

// SetPresc ADC prescaler
func (r *RegisterCcrType) SetPresc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldPrescMask)|(uint32(value)<<RegisterCcrFieldPrescShift))
}

const (
	RegisterCcrFieldVrefenShift = 22
	RegisterCcrFieldVrefenMask  = 0x400000
)

// GetVrefen VREFINT enable
func (r *RegisterCcrType) GetVrefen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldVrefenMask) != 0
}

// SetVrefen VREFINT enable
func (r *RegisterCcrType) SetVrefen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldVrefenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldVrefenMask)
	}
}

const (
	RegisterCcrFieldVsenseenShift = 23
	RegisterCcrFieldVsenseenMask  = 0x800000
)

// GetVsenseen Temperature sensor enable
func (r *RegisterCcrType) GetVsenseen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldVsenseenMask) != 0
}

// SetVsenseen Temperature sensor enable
func (r *RegisterCcrType) SetVsenseen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldVsenseenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldVsenseenMask)
	}
}

const (
	RegisterCcrFieldVbatenShift = 24
	RegisterCcrFieldVbatenMask  = 0x1000000
)

// GetVbaten VBAT enable
func (r *RegisterCcrType) GetVbaten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldVbatenMask) != 0
}

// SetVbaten VBAT enable
func (r *RegisterCcrType) SetVbaten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldVbatenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldVbatenMask)
	}
}

// RegisterCdrType ADC common regular data register for dual and triple modes
type RegisterCdrType uint32

func (r *RegisterCdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCdrFieldRdatamstShift = 0
	RegisterCdrFieldRdatamstMask  = 0xffff
)

// GetRdatamst Regular data of the master ADC
func (r *RegisterCdrType) GetRdatamst() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCdrFieldRdatamstMask) >> RegisterCdrFieldRdatamstShift)
}

// SetRdatamst Regular data of the master ADC
func (r *RegisterCdrType) SetRdatamst(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCdrFieldRdatamstMask)|(uint32(value)<<RegisterCdrFieldRdatamstShift))
}

const (
	RegisterCdrFieldRdataslvShift = 16
	RegisterCdrFieldRdataslvMask  = 0xffff0000
)

// GetRdataslv Regular data of the slave ADC
func (r *RegisterCdrType) GetRdataslv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCdrFieldRdataslvMask) >> RegisterCdrFieldRdataslvShift)
}

// SetRdataslv Regular data of the slave ADC
func (r *RegisterCdrType) SetRdataslv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCdrFieldRdataslvMask)|(uint32(value)<<RegisterCdrFieldRdataslvShift))
}

// RegisterCdr2Type ADC x common regular data register for 32-bit dual mode
type RegisterCdr2Type uint32

func (r *RegisterCdr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCdr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCdr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCdr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCdr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCdr2FieldRdataaltShift = 0
	RegisterCdr2FieldRdataaltMask  = 0xffffffff
)

// GetRdataalt Regular data of the master/slave alternated ADCs
func (r *RegisterCdr2Type) GetRdataalt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCdr2FieldRdataaltMask) >> RegisterCdr2FieldRdataaltShift)
}

// SetRdataalt Regular data of the master/slave alternated ADCs
func (r *RegisterCdr2Type) SetRdataalt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCdr2FieldRdataaltMask)|(uint32(value)<<RegisterCdr2FieldRdataaltShift))
}
