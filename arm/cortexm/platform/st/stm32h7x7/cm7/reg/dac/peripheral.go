//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package dac

import (
	"unsafe"
	"volatile"
)

var (
	Dac = (*_dac)(unsafe.Pointer(uintptr(0x40007400)))
)

type _dac struct {
	Cr      RegisterCrType
	Swtrgr  RegisterSwtrgrType
	Dhr12r1 RegisterDhr12r1Type
	Dhr12l1 RegisterDhr12l1Type
	Dhr8r1  RegisterDhr8r1Type
	Dhr12r2 RegisterDhr12r2Type
	Dhr12l2 RegisterDhr12l2Type
	Dhr8r2  RegisterDhr8r2Type
	Dhr12rd RegisterDhr12rdType
	Dhr12ld RegisterDhr12ldType
	Dhr8rd  RegisterDhr8rdType
	Dor1    RegisterDor1Type
	Dor2    RegisterDor2Type
	Sr      RegisterSrType
	Ccr     RegisterCcrType
	Mcr     RegisterMcrType
	Shsr1   RegisterShsr1Type
	Shsr2   RegisterShsr2Type
	Shhr    RegisterShhrType
	Shrr    RegisterShrrType
}

// RegisterCrType DAC control register
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
	RegisterCrFieldEn1Shift = 0
	RegisterCrFieldEn1Mask  = 0x1
)

// GetEn1 DAC channel1 enable This bit is set and cleared by software to enable/disable DAC channel1.
func (r *RegisterCrType) GetEn1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEn1Mask) != 0
}

// SetEn1 DAC channel1 enable This bit is set and cleared by software to enable/disable DAC channel1.
func (r *RegisterCrType) SetEn1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEn1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEn1Mask)
	}
}

const (
	RegisterCrFieldTen1Shift = 1
	RegisterCrFieldTen1Mask  = 0x2
)

// GetTen1 DAC channel1 trigger enable
func (r *RegisterCrType) GetTen1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTen1Mask) != 0
}

// SetTen1 DAC channel1 trigger enable
func (r *RegisterCrType) SetTen1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTen1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTen1Mask)
	}
}

const (
	RegisterCrFieldTsel1Shift = 2
	RegisterCrFieldTsel1Mask  = 0x1c
)

// GetTsel1 DAC channel1 trigger selection These bits select the external event used to trigger DAC channel1. Note: Only used if bit TEN1 = 1 (DAC channel1 trigger enabled).
func (r *RegisterCrType) GetTsel1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTsel1Mask) >> RegisterCrFieldTsel1Shift)
}

// SetTsel1 DAC channel1 trigger selection These bits select the external event used to trigger DAC channel1. Note: Only used if bit TEN1 = 1 (DAC channel1 trigger enabled).
func (r *RegisterCrType) SetTsel1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTsel1Mask)|(uint32(value)<<RegisterCrFieldTsel1Shift))
}

const (
	RegisterCrFieldWave1Shift = 6
	RegisterCrFieldWave1Mask  = 0xc0
)

// GetWave1 DAC channel1 noise/triangle wave generation enable These bits are set and cleared by software. Note: Only used if bit TEN1 = 1 (DAC channel1 trigger enabled).
func (r *RegisterCrType) GetWave1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldWave1Mask) >> RegisterCrFieldWave1Shift)
}

// SetWave1 DAC channel1 noise/triangle wave generation enable These bits are set and cleared by software. Note: Only used if bit TEN1 = 1 (DAC channel1 trigger enabled).
func (r *RegisterCrType) SetWave1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldWave1Mask)|(uint32(value)<<RegisterCrFieldWave1Shift))
}

const (
	RegisterCrFieldMamp1Shift = 8
	RegisterCrFieldMamp1Mask  = 0xf00
)

// GetMamp1 DAC channel1 mask/amplitude selector These bits are written by software to select mask in wave generation mode or amplitude in triangle generation mode. = 1011: Unmask bits[11:0] of LFSR/ triangle amplitude equal to 4095
func (r *RegisterCrType) GetMamp1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldMamp1Mask) >> RegisterCrFieldMamp1Shift)
}

// SetMamp1 DAC channel1 mask/amplitude selector These bits are written by software to select mask in wave generation mode or amplitude in triangle generation mode. = 1011: Unmask bits[11:0] of LFSR/ triangle amplitude equal to 4095
func (r *RegisterCrType) SetMamp1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldMamp1Mask)|(uint32(value)<<RegisterCrFieldMamp1Shift))
}

const (
	RegisterCrFieldDmaen1Shift = 12
	RegisterCrFieldDmaen1Mask  = 0x1000
)

// GetDmaen1 DAC channel1 DMA enable This bit is set and cleared by software.
func (r *RegisterCrType) GetDmaen1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDmaen1Mask) != 0
}

// SetDmaen1 DAC channel1 DMA enable This bit is set and cleared by software.
func (r *RegisterCrType) SetDmaen1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDmaen1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDmaen1Mask)
	}
}

const (
	RegisterCrFieldDmaudrie1Shift = 13
	RegisterCrFieldDmaudrie1Mask  = 0x2000
)

// GetDmaudrie1 DAC channel1 DMA Underrun Interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) GetDmaudrie1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDmaudrie1Mask) != 0
}

// SetDmaudrie1 DAC channel1 DMA Underrun Interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) SetDmaudrie1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDmaudrie1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDmaudrie1Mask)
	}
}

const (
	RegisterCrFieldCen1Shift = 14
	RegisterCrFieldCen1Mask  = 0x4000
)

// GetCen1 DAC Channel 1 calibration enable This bit is set and cleared by software to enable/disable DAC channel 1 calibration, it can be written only if bit EN1=0 into DAC_CR (the calibration mode can be entered/exit only when the DAC channel is disabled) Otherwise, the write operation is ignored.
func (r *RegisterCrType) GetCen1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCen1Mask) != 0
}

// SetCen1 DAC Channel 1 calibration enable This bit is set and cleared by software to enable/disable DAC channel 1 calibration, it can be written only if bit EN1=0 into DAC_CR (the calibration mode can be entered/exit only when the DAC channel is disabled) Otherwise, the write operation is ignored.
func (r *RegisterCrType) SetCen1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCen1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCen1Mask)
	}
}

const (
	RegisterCrFieldEn2Shift = 16
	RegisterCrFieldEn2Mask  = 0x10000
)

// GetEn2 DAC channel2 enable This bit is set and cleared by software to enable/disable DAC channel2.
func (r *RegisterCrType) GetEn2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEn2Mask) != 0
}

// SetEn2 DAC channel2 enable This bit is set and cleared by software to enable/disable DAC channel2.
func (r *RegisterCrType) SetEn2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEn2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEn2Mask)
	}
}

const (
	RegisterCrFieldTen2Shift = 17
	RegisterCrFieldTen2Mask  = 0x20000
)

// GetTen2 DAC channel2 trigger enable
func (r *RegisterCrType) GetTen2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTen2Mask) != 0
}

// SetTen2 DAC channel2 trigger enable
func (r *RegisterCrType) SetTen2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTen2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTen2Mask)
	}
}

const (
	RegisterCrFieldTsel2Shift = 18
	RegisterCrFieldTsel2Mask  = 0x1c0000
)

// GetTsel2 DAC channel2 trigger selection These bits select the external event used to trigger DAC channel2 Note: Only used if bit TEN2 = 1 (DAC channel2 trigger enabled).
func (r *RegisterCrType) GetTsel2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTsel2Mask) >> RegisterCrFieldTsel2Shift)
}

// SetTsel2 DAC channel2 trigger selection These bits select the external event used to trigger DAC channel2 Note: Only used if bit TEN2 = 1 (DAC channel2 trigger enabled).
func (r *RegisterCrType) SetTsel2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTsel2Mask)|(uint32(value)<<RegisterCrFieldTsel2Shift))
}

const (
	RegisterCrFieldWave2Shift = 22
	RegisterCrFieldWave2Mask  = 0xc00000
)

// GetWave2 DAC channel2 noise/triangle wave generation enable These bits are set/reset by software. 1x: Triangle wave generation enabled Note: Only used if bit TEN2 = 1 (DAC channel2 trigger enabled)
func (r *RegisterCrType) GetWave2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldWave2Mask) >> RegisterCrFieldWave2Shift)
}

// SetWave2 DAC channel2 noise/triangle wave generation enable These bits are set/reset by software. 1x: Triangle wave generation enabled Note: Only used if bit TEN2 = 1 (DAC channel2 trigger enabled)
func (r *RegisterCrType) SetWave2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldWave2Mask)|(uint32(value)<<RegisterCrFieldWave2Shift))
}

const (
	RegisterCrFieldMamp2Shift = 24
	RegisterCrFieldMamp2Mask  = 0xf000000
)

// GetMamp2 DAC channel2 mask/amplitude selector These bits are written by software to select mask in wave generation mode or amplitude in triangle generation mode. = 1011: Unmask bits[11:0] of LFSR/ triangle amplitude equal to 4095
func (r *RegisterCrType) GetMamp2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldMamp2Mask) >> RegisterCrFieldMamp2Shift)
}

// SetMamp2 DAC channel2 mask/amplitude selector These bits are written by software to select mask in wave generation mode or amplitude in triangle generation mode. = 1011: Unmask bits[11:0] of LFSR/ triangle amplitude equal to 4095
func (r *RegisterCrType) SetMamp2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldMamp2Mask)|(uint32(value)<<RegisterCrFieldMamp2Shift))
}

const (
	RegisterCrFieldDmaen2Shift = 28
	RegisterCrFieldDmaen2Mask  = 0x10000000
)

// GetDmaen2 DAC channel2 DMA enable This bit is set and cleared by software.
func (r *RegisterCrType) GetDmaen2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDmaen2Mask) != 0
}

// SetDmaen2 DAC channel2 DMA enable This bit is set and cleared by software.
func (r *RegisterCrType) SetDmaen2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDmaen2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDmaen2Mask)
	}
}

const (
	RegisterCrFieldDmaudrie2Shift = 29
	RegisterCrFieldDmaudrie2Mask  = 0x20000000
)

// GetDmaudrie2 DAC channel2 DMA underrun interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) GetDmaudrie2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDmaudrie2Mask) != 0
}

// SetDmaudrie2 DAC channel2 DMA underrun interrupt enable This bit is set and cleared by software.
func (r *RegisterCrType) SetDmaudrie2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDmaudrie2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDmaudrie2Mask)
	}
}

const (
	RegisterCrFieldCen2Shift = 30
	RegisterCrFieldCen2Mask  = 0x40000000
)

// GetCen2 DAC Channel 2 calibration enable This bit is set and cleared by software to enable/disable DAC channel 2 calibration, it can be written only if bit EN2=0 into DAC_CR (the calibration mode can be entered/exit only when the DAC channel is disabled) Otherwise, the write operation is ignored.
func (r *RegisterCrType) GetCen2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCen2Mask) != 0
}

// SetCen2 DAC Channel 2 calibration enable This bit is set and cleared by software to enable/disable DAC channel 2 calibration, it can be written only if bit EN2=0 into DAC_CR (the calibration mode can be entered/exit only when the DAC channel is disabled) Otherwise, the write operation is ignored.
func (r *RegisterCrType) SetCen2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCen2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCen2Mask)
	}
}

// RegisterSwtrgrType DAC software trigger register
type RegisterSwtrgrType uint32

func (r *RegisterSwtrgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSwtrgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSwtrgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSwtrgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSwtrgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSwtrgrFieldSwtrig1Shift = 0
	RegisterSwtrgrFieldSwtrig1Mask  = 0x1
)

// GetSwtrig1 DAC channel1 software trigger This bit is set by software to trigger the DAC in software trigger mode. Note: This bit is cleared by hardware (one APB1 clock cycle later) once the DAC_DHR1 register value has been loaded into the DAC_DOR1 register.
func (r *RegisterSwtrgrType) GetSwtrig1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwtrgrFieldSwtrig1Mask) != 0
}

// SetSwtrig1 DAC channel1 software trigger This bit is set by software to trigger the DAC in software trigger mode. Note: This bit is cleared by hardware (one APB1 clock cycle later) once the DAC_DHR1 register value has been loaded into the DAC_DOR1 register.
func (r *RegisterSwtrgrType) SetSwtrig1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwtrgrFieldSwtrig1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwtrgrFieldSwtrig1Mask)
	}
}

const (
	RegisterSwtrgrFieldSwtrig2Shift = 1
	RegisterSwtrgrFieldSwtrig2Mask  = 0x2
)

// GetSwtrig2 DAC channel2 software trigger This bit is set by software to trigger the DAC in software trigger mode. Note: This bit is cleared by hardware (one APB1 clock cycle later) once the DAC_DHR2 register value has been loaded into the DAC_DOR2 register.
func (r *RegisterSwtrgrType) GetSwtrig2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSwtrgrFieldSwtrig2Mask) != 0
}

// SetSwtrig2 DAC channel2 software trigger This bit is set by software to trigger the DAC in software trigger mode. Note: This bit is cleared by hardware (one APB1 clock cycle later) once the DAC_DHR2 register value has been loaded into the DAC_DOR2 register.
func (r *RegisterSwtrgrType) SetSwtrig2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSwtrgrFieldSwtrig2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSwtrgrFieldSwtrig2Mask)
	}
}

// RegisterDhr12r1Type DAC channel1 12-bit right-aligned data holding register
type RegisterDhr12r1Type uint32

func (r *RegisterDhr12r1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDhr12r1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDhr12r1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDhr12r1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDhr12r1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDhr12r1FieldDacc1dhrShift = 0
	RegisterDhr12r1FieldDacc1dhrMask  = 0xfff
)

// GetDacc1dhr DAC channel1 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
func (r *RegisterDhr12r1Type) GetDacc1dhr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDhr12r1FieldDacc1dhrMask) >> RegisterDhr12r1FieldDacc1dhrShift)
}

// SetDacc1dhr DAC channel1 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
func (r *RegisterDhr12r1Type) SetDacc1dhr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDhr12r1FieldDacc1dhrMask)|(uint32(value)<<RegisterDhr12r1FieldDacc1dhrShift))
}

// RegisterDhr12l1Type DAC channel1 12-bit left aligned data holding register
type RegisterDhr12l1Type uint32

func (r *RegisterDhr12l1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDhr12l1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDhr12l1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDhr12l1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDhr12l1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDhr12l1FieldDacc1dhrShift = 4
	RegisterDhr12l1FieldDacc1dhrMask  = 0xfff0
)

// GetDacc1dhr DAC channel1 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
func (r *RegisterDhr12l1Type) GetDacc1dhr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDhr12l1FieldDacc1dhrMask) >> RegisterDhr12l1FieldDacc1dhrShift)
}

// SetDacc1dhr DAC channel1 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
func (r *RegisterDhr12l1Type) SetDacc1dhr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDhr12l1FieldDacc1dhrMask)|(uint32(value)<<RegisterDhr12l1FieldDacc1dhrShift))
}

// RegisterDhr8r1Type DAC channel1 8-bit right aligned data holding register
type RegisterDhr8r1Type uint32

func (r *RegisterDhr8r1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDhr8r1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDhr8r1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDhr8r1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDhr8r1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDhr8r1FieldDacc1dhrShift = 0
	RegisterDhr8r1FieldDacc1dhrMask  = 0xff
)

// GetDacc1dhr DAC channel1 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel1.
func (r *RegisterDhr8r1Type) GetDacc1dhr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDhr8r1FieldDacc1dhrMask) >> RegisterDhr8r1FieldDacc1dhrShift)
}

// SetDacc1dhr DAC channel1 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel1.
func (r *RegisterDhr8r1Type) SetDacc1dhr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDhr8r1FieldDacc1dhrMask)|(uint32(value)<<RegisterDhr8r1FieldDacc1dhrShift))
}

// RegisterDhr12r2Type DAC channel2 12-bit right aligned data holding register
type RegisterDhr12r2Type uint32

func (r *RegisterDhr12r2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDhr12r2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDhr12r2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDhr12r2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDhr12r2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDhr12r2FieldDacc2dhrShift = 0
	RegisterDhr12r2FieldDacc2dhrMask  = 0xfff
)

// GetDacc2dhr DAC channel2 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.
func (r *RegisterDhr12r2Type) GetDacc2dhr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDhr12r2FieldDacc2dhrMask) >> RegisterDhr12r2FieldDacc2dhrShift)
}

// SetDacc2dhr DAC channel2 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.
func (r *RegisterDhr12r2Type) SetDacc2dhr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDhr12r2FieldDacc2dhrMask)|(uint32(value)<<RegisterDhr12r2FieldDacc2dhrShift))
}

// RegisterDhr12l2Type DAC channel2 12-bit left aligned data holding register
type RegisterDhr12l2Type uint32

func (r *RegisterDhr12l2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDhr12l2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDhr12l2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDhr12l2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDhr12l2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDhr12l2FieldDacc2dhrShift = 4
	RegisterDhr12l2FieldDacc2dhrMask  = 0xfff0
)

// GetDacc2dhr DAC channel2 12-bit left-aligned data These bits are written by software which specify 12-bit data for DAC channel2.
func (r *RegisterDhr12l2Type) GetDacc2dhr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDhr12l2FieldDacc2dhrMask) >> RegisterDhr12l2FieldDacc2dhrShift)
}

// SetDacc2dhr DAC channel2 12-bit left-aligned data These bits are written by software which specify 12-bit data for DAC channel2.
func (r *RegisterDhr12l2Type) SetDacc2dhr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDhr12l2FieldDacc2dhrMask)|(uint32(value)<<RegisterDhr12l2FieldDacc2dhrShift))
}

// RegisterDhr8r2Type DAC channel2 8-bit right-aligned data holding register
type RegisterDhr8r2Type uint32

func (r *RegisterDhr8r2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDhr8r2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDhr8r2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDhr8r2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDhr8r2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDhr8r2FieldDacc2dhrShift = 0
	RegisterDhr8r2FieldDacc2dhrMask  = 0xff
)

// GetDacc2dhr DAC channel2 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel2.
func (r *RegisterDhr8r2Type) GetDacc2dhr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDhr8r2FieldDacc2dhrMask) >> RegisterDhr8r2FieldDacc2dhrShift)
}

// SetDacc2dhr DAC channel2 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel2.
func (r *RegisterDhr8r2Type) SetDacc2dhr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDhr8r2FieldDacc2dhrMask)|(uint32(value)<<RegisterDhr8r2FieldDacc2dhrShift))
}

// RegisterDhr12rdType Dual DAC 12-bit right-aligned data holding register
type RegisterDhr12rdType uint32

func (r *RegisterDhr12rdType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDhr12rdType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDhr12rdType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDhr12rdType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDhr12rdType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDhr12rdFieldDacc1dhrShift = 0
	RegisterDhr12rdFieldDacc1dhrMask  = 0xfff
)

// GetDacc1dhr DAC channel1 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
func (r *RegisterDhr12rdType) GetDacc1dhr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDhr12rdFieldDacc1dhrMask) >> RegisterDhr12rdFieldDacc1dhrShift)
}

// SetDacc1dhr DAC channel1 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
func (r *RegisterDhr12rdType) SetDacc1dhr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDhr12rdFieldDacc1dhrMask)|(uint32(value)<<RegisterDhr12rdFieldDacc1dhrShift))
}

const (
	RegisterDhr12rdFieldDacc2dhrShift = 16
	RegisterDhr12rdFieldDacc2dhrMask  = 0xfff0000
)

// GetDacc2dhr DAC channel2 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.
func (r *RegisterDhr12rdType) GetDacc2dhr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDhr12rdFieldDacc2dhrMask) >> RegisterDhr12rdFieldDacc2dhrShift)
}

// SetDacc2dhr DAC channel2 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.
func (r *RegisterDhr12rdType) SetDacc2dhr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDhr12rdFieldDacc2dhrMask)|(uint32(value)<<RegisterDhr12rdFieldDacc2dhrShift))
}

// RegisterDhr12ldType DUAL DAC 12-bit left aligned data holding register
type RegisterDhr12ldType uint32

func (r *RegisterDhr12ldType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDhr12ldType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDhr12ldType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDhr12ldType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDhr12ldType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDhr12ldFieldDacc1dhrShift = 4
	RegisterDhr12ldFieldDacc1dhrMask  = 0xfff0
)

// GetDacc1dhr DAC channel1 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
func (r *RegisterDhr12ldType) GetDacc1dhr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDhr12ldFieldDacc1dhrMask) >> RegisterDhr12ldFieldDacc1dhrShift)
}

// SetDacc1dhr DAC channel1 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
func (r *RegisterDhr12ldType) SetDacc1dhr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDhr12ldFieldDacc1dhrMask)|(uint32(value)<<RegisterDhr12ldFieldDacc1dhrShift))
}

const (
	RegisterDhr12ldFieldDacc2dhrShift = 20
	RegisterDhr12ldFieldDacc2dhrMask  = 0xfff00000
)

// GetDacc2dhr DAC channel2 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.
func (r *RegisterDhr12ldType) GetDacc2dhr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDhr12ldFieldDacc2dhrMask) >> RegisterDhr12ldFieldDacc2dhrShift)
}

// SetDacc2dhr DAC channel2 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.
func (r *RegisterDhr12ldType) SetDacc2dhr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDhr12ldFieldDacc2dhrMask)|(uint32(value)<<RegisterDhr12ldFieldDacc2dhrShift))
}

// RegisterDhr8rdType DUAL DAC 8-bit right aligned data holding register
type RegisterDhr8rdType uint32

func (r *RegisterDhr8rdType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDhr8rdType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDhr8rdType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDhr8rdType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDhr8rdType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDhr8rdFieldDacc1dhrShift = 0
	RegisterDhr8rdFieldDacc1dhrMask  = 0xff
)

// GetDacc1dhr DAC channel1 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel1.
func (r *RegisterDhr8rdType) GetDacc1dhr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDhr8rdFieldDacc1dhrMask) >> RegisterDhr8rdFieldDacc1dhrShift)
}

// SetDacc1dhr DAC channel1 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel1.
func (r *RegisterDhr8rdType) SetDacc1dhr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDhr8rdFieldDacc1dhrMask)|(uint32(value)<<RegisterDhr8rdFieldDacc1dhrShift))
}

const (
	RegisterDhr8rdFieldDacc2dhrShift = 8
	RegisterDhr8rdFieldDacc2dhrMask  = 0xff00
)

// GetDacc2dhr DAC channel2 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel2.
func (r *RegisterDhr8rdType) GetDacc2dhr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDhr8rdFieldDacc2dhrMask) >> RegisterDhr8rdFieldDacc2dhrShift)
}

// SetDacc2dhr DAC channel2 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel2.
func (r *RegisterDhr8rdType) SetDacc2dhr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDhr8rdFieldDacc2dhrMask)|(uint32(value)<<RegisterDhr8rdFieldDacc2dhrShift))
}

// RegisterDor1Type DAC channel1 data output register
type RegisterDor1Type uint32

func (r *RegisterDor1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDor1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDor1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDor1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDor1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDor1FieldDacc1dorShift = 0
	RegisterDor1FieldDacc1dorMask  = 0xfff
)

// GetDacc1dor DAC channel1 data output These bits are read-only, they contain data output for DAC channel1.
func (r *RegisterDor1Type) GetDacc1dor() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDor1FieldDacc1dorMask) >> RegisterDor1FieldDacc1dorShift)
}

// SetDacc1dor DAC channel1 data output These bits are read-only, they contain data output for DAC channel1.
func (r *RegisterDor1Type) SetDacc1dor(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDor1FieldDacc1dorMask)|(uint32(value)<<RegisterDor1FieldDacc1dorShift))
}

// RegisterDor2Type DAC channel2 data output register
type RegisterDor2Type uint32

func (r *RegisterDor2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDor2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDor2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDor2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDor2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDor2FieldDacc2dorShift = 0
	RegisterDor2FieldDacc2dorMask  = 0xfff
)

// GetDacc2dor DAC channel2 data output These bits are read-only, they contain data output for DAC channel2.
func (r *RegisterDor2Type) GetDacc2dor() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDor2FieldDacc2dorMask) >> RegisterDor2FieldDacc2dorShift)
}

// SetDacc2dor DAC channel2 data output These bits are read-only, they contain data output for DAC channel2.
func (r *RegisterDor2Type) SetDacc2dor(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDor2FieldDacc2dorMask)|(uint32(value)<<RegisterDor2FieldDacc2dorShift))
}

// RegisterSrType DAC status register
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
	RegisterSrFieldDmaudr1Shift = 13
	RegisterSrFieldDmaudr1Mask  = 0x2000
)

// GetDmaudr1 DAC channel1 DMA underrun flag This bit is set by hardware and cleared by software (by writing it to 1).
func (r *RegisterSrType) GetDmaudr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDmaudr1Mask) != 0
}

// SetDmaudr1 DAC channel1 DMA underrun flag This bit is set by hardware and cleared by software (by writing it to 1).
func (r *RegisterSrType) SetDmaudr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldDmaudr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldDmaudr1Mask)
	}
}

const (
	RegisterSrFieldCalflag1Shift = 14
	RegisterSrFieldCalflag1Mask  = 0x4000
)

// GetCalflag1 DAC Channel 1 calibration offset status This bit is set and cleared by hardware
func (r *RegisterSrType) GetCalflag1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCalflag1Mask) != 0
}

const (
	RegisterSrFieldBwst1Shift = 15
	RegisterSrFieldBwst1Mask  = 0x8000
)

// GetBwst1 DAC Channel 1 busy writing sample time flag This bit is systematically set just after Sample & Hold mode enable and is set each time the software writes the register DAC_SHSR1, It is cleared by hardware when the write operation of DAC_SHSR1 is complete. (It takes about 3LSI periods of synchronization).
func (r *RegisterSrType) GetBwst1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldBwst1Mask) != 0
}

const (
	RegisterSrFieldDmaudr2Shift = 29
	RegisterSrFieldDmaudr2Mask  = 0x20000000
)

// GetDmaudr2 DAC channel2 DMA underrun flag This bit is set by hardware and cleared by software (by writing it to 1).
func (r *RegisterSrType) GetDmaudr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDmaudr2Mask) != 0
}

// SetDmaudr2 DAC channel2 DMA underrun flag This bit is set by hardware and cleared by software (by writing it to 1).
func (r *RegisterSrType) SetDmaudr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldDmaudr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldDmaudr2Mask)
	}
}

const (
	RegisterSrFieldCalflag2Shift = 30
	RegisterSrFieldCalflag2Mask  = 0x40000000
)

// GetCalflag2 DAC Channel 2 calibration offset status This bit is set and cleared by hardware
func (r *RegisterSrType) GetCalflag2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCalflag2Mask) != 0
}

const (
	RegisterSrFieldBwst2Shift = 31
	RegisterSrFieldBwst2Mask  = 0x80000000
)

// GetBwst2 DAC Channel 2 busy writing sample time flag This bit is systematically set just after Sample & Hold mode enable and is set each time the software writes the register DAC_SHSR2, It is cleared by hardware when the write operation of DAC_SHSR2 is complete. (It takes about 3 LSI periods of synchronization).
func (r *RegisterSrType) GetBwst2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldBwst2Mask) != 0
}

// RegisterCcrType DAC calibration control register
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
	RegisterCcrFieldOtrim1Shift = 0
	RegisterCcrFieldOtrim1Mask  = 0x1f
)

// GetOtrim1 DAC Channel 1 offset trimming value
func (r *RegisterCcrType) GetOtrim1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldOtrim1Mask) >> RegisterCcrFieldOtrim1Shift)
}

// SetOtrim1 DAC Channel 1 offset trimming value
func (r *RegisterCcrType) SetOtrim1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldOtrim1Mask)|(uint32(value)<<RegisterCcrFieldOtrim1Shift))
}

const (
	RegisterCcrFieldOtrim2Shift = 16
	RegisterCcrFieldOtrim2Mask  = 0x1f0000
)

// GetOtrim2 DAC Channel 2 offset trimming value
func (r *RegisterCcrType) GetOtrim2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldOtrim2Mask) >> RegisterCcrFieldOtrim2Shift)
}

// SetOtrim2 DAC Channel 2 offset trimming value
func (r *RegisterCcrType) SetOtrim2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldOtrim2Mask)|(uint32(value)<<RegisterCcrFieldOtrim2Shift))
}

// RegisterMcrType DAC mode control register
type RegisterMcrType uint32

func (r *RegisterMcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMcrFieldMode1Shift = 0
	RegisterMcrFieldMode1Mask  = 0x7
)

// GetMode1 DAC Channel 1 mode These bits can be written only when the DAC is disabled and not in the calibration mode (when bit EN1=0 and bit CEN1 =0 in the DAC_CR register). If EN1=1 or CEN1 =1 the write operation is ignored. They can be set and cleared by software to select the DAC Channel 1 mode: DAC Channel 1 in normal Mode DAC Channel 1 in sample &amp; hold mode
func (r *RegisterMcrType) GetMode1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldMode1Mask) >> RegisterMcrFieldMode1Shift)
}

// SetMode1 DAC Channel 1 mode These bits can be written only when the DAC is disabled and not in the calibration mode (when bit EN1=0 and bit CEN1 =0 in the DAC_CR register). If EN1=1 or CEN1 =1 the write operation is ignored. They can be set and cleared by software to select the DAC Channel 1 mode: DAC Channel 1 in normal Mode DAC Channel 1 in sample &amp; hold mode
func (r *RegisterMcrType) SetMode1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldMode1Mask)|(uint32(value)<<RegisterMcrFieldMode1Shift))
}

const (
	RegisterMcrFieldMode2Shift = 16
	RegisterMcrFieldMode2Mask  = 0x70000
)

// GetMode2 DAC Channel 2 mode These bits can be written only when the DAC is disabled and not in the calibration mode (when bit EN2=0 and bit CEN2 =0 in the DAC_CR register). If EN2=1 or CEN2 =1 the write operation is ignored. They can be set and cleared by software to select the DAC Channel 2 mode: DAC Channel 2 in normal Mode DAC Channel 2 in sample &amp; hold mode
func (r *RegisterMcrType) GetMode2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldMode2Mask) >> RegisterMcrFieldMode2Shift)
}

// SetMode2 DAC Channel 2 mode These bits can be written only when the DAC is disabled and not in the calibration mode (when bit EN2=0 and bit CEN2 =0 in the DAC_CR register). If EN2=1 or CEN2 =1 the write operation is ignored. They can be set and cleared by software to select the DAC Channel 2 mode: DAC Channel 2 in normal Mode DAC Channel 2 in sample &amp; hold mode
func (r *RegisterMcrType) SetMode2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldMode2Mask)|(uint32(value)<<RegisterMcrFieldMode2Shift))
}

// RegisterShsr1Type DAC Sample and Hold sample time register 1
type RegisterShsr1Type uint32

func (r *RegisterShsr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterShsr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterShsr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterShsr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterShsr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterShsr1FieldTsample1Shift = 0
	RegisterShsr1FieldTsample1Mask  = 0x3ff
)

// GetTsample1 DAC Channel 1 sample Time (only valid in sample &amp; hold mode) These bits can be written when the DAC channel1 is disabled or also during normal operation. in the latter case, the write can be done only when BWSTx of DAC_SR register is low, If BWSTx=1, the write operation is ignored.
func (r *RegisterShsr1Type) GetTsample1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterShsr1FieldTsample1Mask) >> RegisterShsr1FieldTsample1Shift)
}

// SetTsample1 DAC Channel 1 sample Time (only valid in sample &amp; hold mode) These bits can be written when the DAC channel1 is disabled or also during normal operation. in the latter case, the write can be done only when BWSTx of DAC_SR register is low, If BWSTx=1, the write operation is ignored.
func (r *RegisterShsr1Type) SetTsample1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShsr1FieldTsample1Mask)|(uint32(value)<<RegisterShsr1FieldTsample1Shift))
}

// RegisterShsr2Type DAC Sample and Hold sample time register 2
type RegisterShsr2Type uint32

func (r *RegisterShsr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterShsr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterShsr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterShsr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterShsr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterShsr2FieldTsample2Shift = 0
	RegisterShsr2FieldTsample2Mask  = 0x3ff
)

// GetTsample2 DAC Channel 2 sample Time (only valid in sample &amp; hold mode) These bits can be written when the DAC channel2 is disabled or also during normal operation. in the latter case, the write can be done only when BWSTx of DAC_SR register is low, if BWSTx=1, the write operation is ignored.
func (r *RegisterShsr2Type) GetTsample2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterShsr2FieldTsample2Mask) >> RegisterShsr2FieldTsample2Shift)
}

// SetTsample2 DAC Channel 2 sample Time (only valid in sample &amp; hold mode) These bits can be written when the DAC channel2 is disabled or also during normal operation. in the latter case, the write can be done only when BWSTx of DAC_SR register is low, if BWSTx=1, the write operation is ignored.
func (r *RegisterShsr2Type) SetTsample2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShsr2FieldTsample2Mask)|(uint32(value)<<RegisterShsr2FieldTsample2Shift))
}

// RegisterShhrType DAC Sample and Hold hold time register
type RegisterShhrType uint32

func (r *RegisterShhrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterShhrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterShhrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterShhrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterShhrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterShhrFieldThold1Shift = 0
	RegisterShhrFieldThold1Mask  = 0x3ff
)

// GetThold1 DAC Channel 1 hold Time (only valid in sample &amp; hold mode) Hold time= (THOLD[9:0]) x T LSI
func (r *RegisterShhrType) GetThold1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterShhrFieldThold1Mask) >> RegisterShhrFieldThold1Shift)
}

// SetThold1 DAC Channel 1 hold Time (only valid in sample &amp; hold mode) Hold time= (THOLD[9:0]) x T LSI
func (r *RegisterShhrType) SetThold1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShhrFieldThold1Mask)|(uint32(value)<<RegisterShhrFieldThold1Shift))
}

const (
	RegisterShhrFieldThold2Shift = 16
	RegisterShhrFieldThold2Mask  = 0x3ff0000
)

// GetThold2 DAC Channel 2 hold time (only valid in sample &amp; hold mode). Hold time= (THOLD[9:0]) x T LSI
func (r *RegisterShhrType) GetThold2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterShhrFieldThold2Mask) >> RegisterShhrFieldThold2Shift)
}

// SetThold2 DAC Channel 2 hold time (only valid in sample &amp; hold mode). Hold time= (THOLD[9:0]) x T LSI
func (r *RegisterShhrType) SetThold2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShhrFieldThold2Mask)|(uint32(value)<<RegisterShhrFieldThold2Shift))
}

// RegisterShrrType DAC Sample and Hold refresh time register
type RegisterShrrType uint32

func (r *RegisterShrrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterShrrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterShrrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterShrrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterShrrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterShrrFieldTrefresh1Shift = 0
	RegisterShrrFieldTrefresh1Mask  = 0xff
)

// GetTrefresh1 DAC Channel 1 refresh Time (only valid in sample &amp; hold mode) Refresh time= (TREFRESH[7:0]) x T LSI
func (r *RegisterShrrType) GetTrefresh1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShrrFieldTrefresh1Mask) >> RegisterShrrFieldTrefresh1Shift)
}

// SetTrefresh1 DAC Channel 1 refresh Time (only valid in sample &amp; hold mode) Refresh time= (TREFRESH[7:0]) x T LSI
func (r *RegisterShrrType) SetTrefresh1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShrrFieldTrefresh1Mask)|(uint32(value)<<RegisterShrrFieldTrefresh1Shift))
}

const (
	RegisterShrrFieldTrefresh2Shift = 16
	RegisterShrrFieldTrefresh2Mask  = 0xff0000
)

// GetTrefresh2 DAC Channel 2 refresh Time (only valid in sample &amp; hold mode) Refresh time= (TREFRESH[7:0]) x T LSI
func (r *RegisterShrrType) GetTrefresh2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterShrrFieldTrefresh2Mask) >> RegisterShrrFieldTrefresh2Shift)
}

// SetTrefresh2 DAC Channel 2 refresh Time (only valid in sample &amp; hold mode) Refresh time= (TREFRESH[7:0]) x T LSI
func (r *RegisterShrrType) SetTrefresh2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterShrrFieldTrefresh2Mask)|(uint32(value)<<RegisterShrrFieldTrefresh2Shift))
}
