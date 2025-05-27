//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package hrtim_common

import (
	"unsafe"
	"volatile"
)

var (
	Hrtim_common = (*_hrtim_common)(unsafe.Pointer(uintptr(0x40017780)))
)

type _hrtim_common struct {
	Cr1     registerCr1Type
	Cr2     registerCr2Type
	Isr     registerIsrType
	Icr     registerIcrType
	Ier     registerIerType
	Oenr    registerOenrType
	Disr    registerDisrType
	Odsr    registerOdsrType
	Bmcr    registerBmcrType
	Bmtrg   registerBmtrgType
	Bmcmpr6 registerBmcmpr6Type
	Bmper   registerBmperType
	Eecr1   registerEecr1Type
	Eecr2   registerEecr2Type
	Eecr3   registerEecr3Type
	Adc1r   registerAdc1rType
	Adc2r   registerAdc2rType
	Adc3r   registerAdc3rType
	Adc4r   registerAdc4rType
	Dllcr   registerDllcrType
	Fltinr1 registerFltinr1Type
	Fltinr2 registerFltinr2Type
	Bdmupdr registerBdmupdrType
	Bdtxupr registerBdtxuprType
	Bdmadr  registerBdmadrType
}

// registerCr1Type Control Register 1
type registerCr1Type uint32

const (
	RegisterCr1FieldMudisShift = 0
	RegisterCr1FieldMudisMask  = 0x1
)

// GetMudis Master Update Disable
func (r *registerCr1Type) GetMudis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldMudisMask) != 0
}

// SetMudis Master Update Disable
func (r *registerCr1Type) SetMudis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldMudisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldMudisMask)
	}
}

const (
	RegisterCr1FieldTaudisShift = 1
	RegisterCr1FieldTaudisMask  = 0x2
)

// GetTaudis Timer A Update Disable
func (r *registerCr1Type) GetTaudis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTaudisMask) != 0
}

// SetTaudis Timer A Update Disable
func (r *registerCr1Type) SetTaudis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTaudisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTaudisMask)
	}
}

const (
	RegisterCr1FieldTbudisShift = 2
	RegisterCr1FieldTbudisMask  = 0x4
)

// GetTbudis Timer B Update Disable
func (r *registerCr1Type) GetTbudis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTbudisMask) != 0
}

// SetTbudis Timer B Update Disable
func (r *registerCr1Type) SetTbudis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTbudisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTbudisMask)
	}
}

const (
	RegisterCr1FieldTcudisShift = 3
	RegisterCr1FieldTcudisMask  = 0x8
)

// GetTcudis Timer C Update Disable
func (r *registerCr1Type) GetTcudis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTcudisMask) != 0
}

// SetTcudis Timer C Update Disable
func (r *registerCr1Type) SetTcudis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTcudisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTcudisMask)
	}
}

const (
	RegisterCr1FieldTdudisShift = 4
	RegisterCr1FieldTdudisMask  = 0x10
)

// GetTdudis Timer D Update Disable
func (r *registerCr1Type) GetTdudis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTdudisMask) != 0
}

// SetTdudis Timer D Update Disable
func (r *registerCr1Type) SetTdudis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTdudisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTdudisMask)
	}
}

const (
	RegisterCr1FieldTeudisShift = 5
	RegisterCr1FieldTeudisMask  = 0x20
)

// GetTeudis Timer E Update Disable
func (r *registerCr1Type) GetTeudis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTeudisMask) != 0
}

// SetTeudis Timer E Update Disable
func (r *registerCr1Type) SetTeudis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTeudisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTeudisMask)
	}
}

const (
	RegisterCr1FieldAd1usrcShift = 16
	RegisterCr1FieldAd1usrcMask  = 0x70000
)

// GetAd1usrc ADC Trigger 1 Update Source
func (r *registerCr1Type) GetAd1usrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldAd1usrcMask) >> RegisterCr1FieldAd1usrcShift)
}

// SetAd1usrc ADC Trigger 1 Update Source
func (r *registerCr1Type) SetAd1usrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldAd1usrcMask)|(uint32(value)<<RegisterCr1FieldAd1usrcShift))
}

const (
	RegisterCr1FieldAd2usrcShift = 19
	RegisterCr1FieldAd2usrcMask  = 0x380000
)

// GetAd2usrc ADC Trigger 2 Update Source
func (r *registerCr1Type) GetAd2usrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldAd2usrcMask) >> RegisterCr1FieldAd2usrcShift)
}

// SetAd2usrc ADC Trigger 2 Update Source
func (r *registerCr1Type) SetAd2usrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldAd2usrcMask)|(uint32(value)<<RegisterCr1FieldAd2usrcShift))
}

const (
	RegisterCr1FieldAd3usrcShift = 22
	RegisterCr1FieldAd3usrcMask  = 0x1c00000
)

// GetAd3usrc ADC Trigger 3 Update Source
func (r *registerCr1Type) GetAd3usrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldAd3usrcMask) >> RegisterCr1FieldAd3usrcShift)
}

// SetAd3usrc ADC Trigger 3 Update Source
func (r *registerCr1Type) SetAd3usrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldAd3usrcMask)|(uint32(value)<<RegisterCr1FieldAd3usrcShift))
}

const (
	RegisterCr1FieldAd4usrcShift = 25
	RegisterCr1FieldAd4usrcMask  = 0xe000000
)

// GetAd4usrc ADC Trigger 4 Update Source
func (r *registerCr1Type) GetAd4usrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldAd4usrcMask) >> RegisterCr1FieldAd4usrcShift)
}

// SetAd4usrc ADC Trigger 4 Update Source
func (r *registerCr1Type) SetAd4usrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldAd4usrcMask)|(uint32(value)<<RegisterCr1FieldAd4usrcShift))
}

// registerCr2Type Control Register 2
type registerCr2Type uint32

const (
	RegisterCr2FieldMswuShift = 0
	RegisterCr2FieldMswuMask  = 0x1
)

// GetMswu Master Timer Software update
func (r *registerCr2Type) GetMswu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldMswuMask) != 0
}

// SetMswu Master Timer Software update
func (r *registerCr2Type) SetMswu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldMswuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldMswuMask)
	}
}

const (
	RegisterCr2FieldTaswuShift = 1
	RegisterCr2FieldTaswuMask  = 0x2
)

// GetTaswu Timer A Software update
func (r *registerCr2Type) GetTaswu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTaswuMask) != 0
}

// SetTaswu Timer A Software update
func (r *registerCr2Type) SetTaswu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTaswuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTaswuMask)
	}
}

const (
	RegisterCr2FieldTbswuShift = 2
	RegisterCr2FieldTbswuMask  = 0x4
)

// GetTbswu Timer B Software Update
func (r *registerCr2Type) GetTbswu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTbswuMask) != 0
}

// SetTbswu Timer B Software Update
func (r *registerCr2Type) SetTbswu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTbswuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTbswuMask)
	}
}

const (
	RegisterCr2FieldTcswuShift = 3
	RegisterCr2FieldTcswuMask  = 0x8
)

// GetTcswu Timer C Software Update
func (r *registerCr2Type) GetTcswu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTcswuMask) != 0
}

// SetTcswu Timer C Software Update
func (r *registerCr2Type) SetTcswu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTcswuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTcswuMask)
	}
}

const (
	RegisterCr2FieldTdswuShift = 4
	RegisterCr2FieldTdswuMask  = 0x10
)

// GetTdswu Timer D Software Update
func (r *registerCr2Type) GetTdswu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTdswuMask) != 0
}

// SetTdswu Timer D Software Update
func (r *registerCr2Type) SetTdswu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTdswuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTdswuMask)
	}
}

const (
	RegisterCr2FieldTeswuShift = 5
	RegisterCr2FieldTeswuMask  = 0x20
)

// GetTeswu Timer E Software Update
func (r *registerCr2Type) GetTeswu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTeswuMask) != 0
}

// SetTeswu Timer E Software Update
func (r *registerCr2Type) SetTeswu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTeswuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTeswuMask)
	}
}

const (
	RegisterCr2FieldMrstShift = 8
	RegisterCr2FieldMrstMask  = 0x100
)

// GetMrst Master Counter software reset
func (r *registerCr2Type) GetMrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldMrstMask) != 0
}

// SetMrst Master Counter software reset
func (r *registerCr2Type) SetMrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldMrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldMrstMask)
	}
}

const (
	RegisterCr2FieldTarstShift = 9
	RegisterCr2FieldTarstMask  = 0x200
)

// GetTarst Timer A counter software reset
func (r *registerCr2Type) GetTarst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTarstMask) != 0
}

// SetTarst Timer A counter software reset
func (r *registerCr2Type) SetTarst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTarstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTarstMask)
	}
}

const (
	RegisterCr2FieldTbrstShift = 10
	RegisterCr2FieldTbrstMask  = 0x400
)

// GetTbrst Timer B counter software reset
func (r *registerCr2Type) GetTbrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTbrstMask) != 0
}

// SetTbrst Timer B counter software reset
func (r *registerCr2Type) SetTbrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTbrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTbrstMask)
	}
}

const (
	RegisterCr2FieldTcrstShift = 11
	RegisterCr2FieldTcrstMask  = 0x800
)

// GetTcrst Timer C counter software reset
func (r *registerCr2Type) GetTcrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTcrstMask) != 0
}

// SetTcrst Timer C counter software reset
func (r *registerCr2Type) SetTcrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTcrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTcrstMask)
	}
}

const (
	RegisterCr2FieldTdrstShift = 12
	RegisterCr2FieldTdrstMask  = 0x1000
)

// GetTdrst Timer D counter software reset
func (r *registerCr2Type) GetTdrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTdrstMask) != 0
}

// SetTdrst Timer D counter software reset
func (r *registerCr2Type) SetTdrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTdrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTdrstMask)
	}
}

const (
	RegisterCr2FieldTerstShift = 13
	RegisterCr2FieldTerstMask  = 0x2000
)

// GetTerst Timer E counter software reset
func (r *registerCr2Type) GetTerst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTerstMask) != 0
}

// SetTerst Timer E counter software reset
func (r *registerCr2Type) SetTerst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTerstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTerstMask)
	}
}

// registerIsrType Interrupt Status Register
type registerIsrType uint32

const (
	RegisterIsrFieldFlt1Shift = 0
	RegisterIsrFieldFlt1Mask  = 0x1
)

// GetFlt1 Fault 1 Interrupt Flag
func (r *registerIsrType) GetFlt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldFlt1Mask) != 0
}

const (
	RegisterIsrFieldFlt2Shift = 1
	RegisterIsrFieldFlt2Mask  = 0x2
)

// GetFlt2 Fault 2 Interrupt Flag
func (r *registerIsrType) GetFlt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldFlt2Mask) != 0
}

const (
	RegisterIsrFieldFlt3Shift = 2
	RegisterIsrFieldFlt3Mask  = 0x4
)

// GetFlt3 Fault 3 Interrupt Flag
func (r *registerIsrType) GetFlt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldFlt3Mask) != 0
}

const (
	RegisterIsrFieldFlt4Shift = 3
	RegisterIsrFieldFlt4Mask  = 0x8
)

// GetFlt4 Fault 4 Interrupt Flag
func (r *registerIsrType) GetFlt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldFlt4Mask) != 0
}

const (
	RegisterIsrFieldFlt5Shift = 4
	RegisterIsrFieldFlt5Mask  = 0x10
)

// GetFlt5 Fault 5 Interrupt Flag
func (r *registerIsrType) GetFlt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldFlt5Mask) != 0
}

const (
	RegisterIsrFieldSysfltShift = 5
	RegisterIsrFieldSysfltMask  = 0x20
)

// GetSysflt System Fault Interrupt Flag
func (r *registerIsrType) GetSysflt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldSysfltMask) != 0
}

// SetSysflt System Fault Interrupt Flag
func (r *registerIsrType) SetSysflt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldSysfltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldSysfltMask)
	}
}

const (
	RegisterIsrFieldDllrdyShift = 16
	RegisterIsrFieldDllrdyMask  = 0x10000
)

// GetDllrdy DLL Ready Interrupt Flag
func (r *registerIsrType) GetDllrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldDllrdyMask) != 0
}

const (
	RegisterIsrFieldBmperShift = 17
	RegisterIsrFieldBmperMask  = 0x20000
)

// GetBmper Burst mode Period Interrupt Flag
func (r *registerIsrType) GetBmper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldBmperMask) != 0
}

// registerIcrType Interrupt Clear Register
type registerIcrType uint32

const (
	RegisterIcrFieldFlt1cShift = 0
	RegisterIcrFieldFlt1cMask  = 0x1
)

// SetFlt1c Fault 1 Interrupt Flag Clear
func (r *registerIcrType) SetFlt1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldFlt1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldFlt1cMask)
	}
}

const (
	RegisterIcrFieldFlt2cShift = 1
	RegisterIcrFieldFlt2cMask  = 0x2
)

// SetFlt2c Fault 2 Interrupt Flag Clear
func (r *registerIcrType) SetFlt2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldFlt2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldFlt2cMask)
	}
}

const (
	RegisterIcrFieldFlt3cShift = 2
	RegisterIcrFieldFlt3cMask  = 0x4
)

// SetFlt3c Fault 3 Interrupt Flag Clear
func (r *registerIcrType) SetFlt3c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldFlt3cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldFlt3cMask)
	}
}

const (
	RegisterIcrFieldFlt4cShift = 3
	RegisterIcrFieldFlt4cMask  = 0x8
)

// SetFlt4c Fault 4 Interrupt Flag Clear
func (r *registerIcrType) SetFlt4c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldFlt4cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldFlt4cMask)
	}
}

const (
	RegisterIcrFieldFlt5cShift = 4
	RegisterIcrFieldFlt5cMask  = 0x10
)

// SetFlt5c Fault 5 Interrupt Flag Clear
func (r *registerIcrType) SetFlt5c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldFlt5cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldFlt5cMask)
	}
}

const (
	RegisterIcrFieldSysfltcShift = 5
	RegisterIcrFieldSysfltcMask  = 0x20
)

// GetSysfltc System Fault Interrupt Flag Clear
func (r *registerIcrType) GetSysfltc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldSysfltcMask) != 0
}

// SetSysfltc System Fault Interrupt Flag Clear
func (r *registerIcrType) SetSysfltc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldSysfltcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldSysfltcMask)
	}
}

const (
	RegisterIcrFieldDllrdycShift = 16
	RegisterIcrFieldDllrdycMask  = 0x10000
)

// SetDllrdyc DLL Ready Interrupt flag Clear
func (r *registerIcrType) SetDllrdyc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldDllrdycMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldDllrdycMask)
	}
}

const (
	RegisterIcrFieldBmpercShift = 17
	RegisterIcrFieldBmpercMask  = 0x20000
)

// SetBmperc Burst mode period flag Clear
func (r *registerIcrType) SetBmperc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldBmpercMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldBmpercMask)
	}
}

// registerIerType Interrupt Enable Register
type registerIerType uint32

const (
	RegisterIerFieldFlt1ieShift = 0
	RegisterIerFieldFlt1ieMask  = 0x1
)

// GetFlt1ie Fault 1 Interrupt Enable
func (r *registerIerType) GetFlt1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldFlt1ieMask) != 0
}

// SetFlt1ie Fault 1 Interrupt Enable
func (r *registerIerType) SetFlt1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldFlt1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldFlt1ieMask)
	}
}

const (
	RegisterIerFieldFlt2ieShift = 1
	RegisterIerFieldFlt2ieMask  = 0x2
)

// GetFlt2ie Fault 2 Interrupt Enable
func (r *registerIerType) GetFlt2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldFlt2ieMask) != 0
}

// SetFlt2ie Fault 2 Interrupt Enable
func (r *registerIerType) SetFlt2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldFlt2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldFlt2ieMask)
	}
}

const (
	RegisterIerFieldFlt3ieShift = 2
	RegisterIerFieldFlt3ieMask  = 0x4
)

// GetFlt3ie Fault 3 Interrupt Enable
func (r *registerIerType) GetFlt3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldFlt3ieMask) != 0
}

// SetFlt3ie Fault 3 Interrupt Enable
func (r *registerIerType) SetFlt3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldFlt3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldFlt3ieMask)
	}
}

const (
	RegisterIerFieldFlt4ieShift = 3
	RegisterIerFieldFlt4ieMask  = 0x8
)

// GetFlt4ie Fault 4 Interrupt Enable
func (r *registerIerType) GetFlt4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldFlt4ieMask) != 0
}

// SetFlt4ie Fault 4 Interrupt Enable
func (r *registerIerType) SetFlt4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldFlt4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldFlt4ieMask)
	}
}

const (
	RegisterIerFieldFlt5ieShift = 4
	RegisterIerFieldFlt5ieMask  = 0x10
)

// GetFlt5ie Fault 5 Interrupt Enable
func (r *registerIerType) GetFlt5ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldFlt5ieMask) != 0
}

// SetFlt5ie Fault 5 Interrupt Enable
func (r *registerIerType) SetFlt5ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldFlt5ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldFlt5ieMask)
	}
}

const (
	RegisterIerFieldSysflteShift = 5
	RegisterIerFieldSysflteMask  = 0x20
)

// GetSysflte System Fault Interrupt Enable
func (r *registerIerType) GetSysflte() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldSysflteMask) != 0
}

// SetSysflte System Fault Interrupt Enable
func (r *registerIerType) SetSysflte(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldSysflteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldSysflteMask)
	}
}

const (
	RegisterIerFieldDllrdyieShift = 16
	RegisterIerFieldDllrdyieMask  = 0x10000
)

// GetDllrdyie DLL Ready Interrupt Enable
func (r *registerIerType) GetDllrdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldDllrdyieMask) != 0
}

// SetDllrdyie DLL Ready Interrupt Enable
func (r *registerIerType) SetDllrdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldDllrdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldDllrdyieMask)
	}
}

const (
	RegisterIerFieldBmperieShift = 17
	RegisterIerFieldBmperieMask  = 0x20000
)

// GetBmperie Burst mode period Interrupt Enable
func (r *registerIerType) GetBmperie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldBmperieMask) != 0
}

// SetBmperie Burst mode period Interrupt Enable
func (r *registerIerType) SetBmperie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldBmperieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldBmperieMask)
	}
}

// registerOenrType Output Enable Register
type registerOenrType uint32

const (
	RegisterOenrFieldTa1oenShift = 0
	RegisterOenrFieldTa1oenMask  = 0x1
)

// GetTa1oen Timer A Output 1 Enable
func (r *registerOenrType) GetTa1oen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOenrFieldTa1oenMask) != 0
}

// SetTa1oen Timer A Output 1 Enable
func (r *registerOenrType) SetTa1oen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOenrFieldTa1oenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOenrFieldTa1oenMask)
	}
}

const (
	RegisterOenrFieldTa2oenShift = 1
	RegisterOenrFieldTa2oenMask  = 0x2
)

// GetTa2oen Timer A Output 2 Enable
func (r *registerOenrType) GetTa2oen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOenrFieldTa2oenMask) != 0
}

// SetTa2oen Timer A Output 2 Enable
func (r *registerOenrType) SetTa2oen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOenrFieldTa2oenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOenrFieldTa2oenMask)
	}
}

const (
	RegisterOenrFieldTb1oenShift = 2
	RegisterOenrFieldTb1oenMask  = 0x4
)

// GetTb1oen Timer B Output 1 Enable
func (r *registerOenrType) GetTb1oen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOenrFieldTb1oenMask) != 0
}

// SetTb1oen Timer B Output 1 Enable
func (r *registerOenrType) SetTb1oen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOenrFieldTb1oenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOenrFieldTb1oenMask)
	}
}

const (
	RegisterOenrFieldTb2oenShift = 3
	RegisterOenrFieldTb2oenMask  = 0x8
)

// GetTb2oen Timer B Output 2 Enable
func (r *registerOenrType) GetTb2oen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOenrFieldTb2oenMask) != 0
}

// SetTb2oen Timer B Output 2 Enable
func (r *registerOenrType) SetTb2oen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOenrFieldTb2oenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOenrFieldTb2oenMask)
	}
}

const (
	RegisterOenrFieldTc1oenShift = 4
	RegisterOenrFieldTc1oenMask  = 0x10
)

// GetTc1oen Timer C Output 1 Enable
func (r *registerOenrType) GetTc1oen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOenrFieldTc1oenMask) != 0
}

// SetTc1oen Timer C Output 1 Enable
func (r *registerOenrType) SetTc1oen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOenrFieldTc1oenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOenrFieldTc1oenMask)
	}
}

const (
	RegisterOenrFieldTc2oenShift = 5
	RegisterOenrFieldTc2oenMask  = 0x20
)

// GetTc2oen Timer C Output 2 Enable
func (r *registerOenrType) GetTc2oen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOenrFieldTc2oenMask) != 0
}

// SetTc2oen Timer C Output 2 Enable
func (r *registerOenrType) SetTc2oen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOenrFieldTc2oenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOenrFieldTc2oenMask)
	}
}

const (
	RegisterOenrFieldTd1oenShift = 6
	RegisterOenrFieldTd1oenMask  = 0x40
)

// GetTd1oen Timer D Output 1 Enable
func (r *registerOenrType) GetTd1oen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOenrFieldTd1oenMask) != 0
}

// SetTd1oen Timer D Output 1 Enable
func (r *registerOenrType) SetTd1oen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOenrFieldTd1oenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOenrFieldTd1oenMask)
	}
}

const (
	RegisterOenrFieldTd2oenShift = 7
	RegisterOenrFieldTd2oenMask  = 0x80
)

// GetTd2oen Timer D Output 2 Enable
func (r *registerOenrType) GetTd2oen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOenrFieldTd2oenMask) != 0
}

// SetTd2oen Timer D Output 2 Enable
func (r *registerOenrType) SetTd2oen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOenrFieldTd2oenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOenrFieldTd2oenMask)
	}
}

const (
	RegisterOenrFieldTe1oenShift = 8
	RegisterOenrFieldTe1oenMask  = 0x100
)

// GetTe1oen Timer E Output 1 Enable
func (r *registerOenrType) GetTe1oen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOenrFieldTe1oenMask) != 0
}

// SetTe1oen Timer E Output 1 Enable
func (r *registerOenrType) SetTe1oen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOenrFieldTe1oenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOenrFieldTe1oenMask)
	}
}

const (
	RegisterOenrFieldTe2oenShift = 9
	RegisterOenrFieldTe2oenMask  = 0x200
)

// GetTe2oen Timer E Output 2 Enable
func (r *registerOenrType) GetTe2oen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOenrFieldTe2oenMask) != 0
}

// SetTe2oen Timer E Output 2 Enable
func (r *registerOenrType) SetTe2oen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOenrFieldTe2oenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOenrFieldTe2oenMask)
	}
}

// registerDisrType DISR
type registerDisrType uint32

const (
	RegisterDisrFieldTa1odisShift = 0
	RegisterDisrFieldTa1odisMask  = 0x1
)

// GetTa1odis TA1ODIS
func (r *registerDisrType) GetTa1odis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDisrFieldTa1odisMask) != 0
}

// SetTa1odis TA1ODIS
func (r *registerDisrType) SetTa1odis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDisrFieldTa1odisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDisrFieldTa1odisMask)
	}
}

const (
	RegisterDisrFieldTa2odisShift = 1
	RegisterDisrFieldTa2odisMask  = 0x2
)

// GetTa2odis TA2ODIS
func (r *registerDisrType) GetTa2odis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDisrFieldTa2odisMask) != 0
}

// SetTa2odis TA2ODIS
func (r *registerDisrType) SetTa2odis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDisrFieldTa2odisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDisrFieldTa2odisMask)
	}
}

const (
	RegisterDisrFieldTb1odisShift = 2
	RegisterDisrFieldTb1odisMask  = 0x4
)

// GetTb1odis TB1ODIS
func (r *registerDisrType) GetTb1odis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDisrFieldTb1odisMask) != 0
}

// SetTb1odis TB1ODIS
func (r *registerDisrType) SetTb1odis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDisrFieldTb1odisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDisrFieldTb1odisMask)
	}
}

const (
	RegisterDisrFieldTb2odisShift = 3
	RegisterDisrFieldTb2odisMask  = 0x8
)

// GetTb2odis TB2ODIS
func (r *registerDisrType) GetTb2odis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDisrFieldTb2odisMask) != 0
}

// SetTb2odis TB2ODIS
func (r *registerDisrType) SetTb2odis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDisrFieldTb2odisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDisrFieldTb2odisMask)
	}
}

const (
	RegisterDisrFieldTc1odisShift = 4
	RegisterDisrFieldTc1odisMask  = 0x10
)

// GetTc1odis TC1ODIS
func (r *registerDisrType) GetTc1odis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDisrFieldTc1odisMask) != 0
}

// SetTc1odis TC1ODIS
func (r *registerDisrType) SetTc1odis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDisrFieldTc1odisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDisrFieldTc1odisMask)
	}
}

const (
	RegisterDisrFieldTc2odisShift = 5
	RegisterDisrFieldTc2odisMask  = 0x20
)

// GetTc2odis TC2ODIS
func (r *registerDisrType) GetTc2odis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDisrFieldTc2odisMask) != 0
}

// SetTc2odis TC2ODIS
func (r *registerDisrType) SetTc2odis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDisrFieldTc2odisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDisrFieldTc2odisMask)
	}
}

const (
	RegisterDisrFieldTd1odisShift = 6
	RegisterDisrFieldTd1odisMask  = 0x40
)

// GetTd1odis TD1ODIS
func (r *registerDisrType) GetTd1odis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDisrFieldTd1odisMask) != 0
}

// SetTd1odis TD1ODIS
func (r *registerDisrType) SetTd1odis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDisrFieldTd1odisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDisrFieldTd1odisMask)
	}
}

const (
	RegisterDisrFieldTd2odisShift = 7
	RegisterDisrFieldTd2odisMask  = 0x80
)

// GetTd2odis TD2ODIS
func (r *registerDisrType) GetTd2odis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDisrFieldTd2odisMask) != 0
}

// SetTd2odis TD2ODIS
func (r *registerDisrType) SetTd2odis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDisrFieldTd2odisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDisrFieldTd2odisMask)
	}
}

const (
	RegisterDisrFieldTe1odisShift = 8
	RegisterDisrFieldTe1odisMask  = 0x100
)

// GetTe1odis TE1ODIS
func (r *registerDisrType) GetTe1odis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDisrFieldTe1odisMask) != 0
}

// SetTe1odis TE1ODIS
func (r *registerDisrType) SetTe1odis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDisrFieldTe1odisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDisrFieldTe1odisMask)
	}
}

const (
	RegisterDisrFieldTe2odisShift = 9
	RegisterDisrFieldTe2odisMask  = 0x200
)

// GetTe2odis TE2ODIS
func (r *registerDisrType) GetTe2odis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDisrFieldTe2odisMask) != 0
}

// SetTe2odis TE2ODIS
func (r *registerDisrType) SetTe2odis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDisrFieldTe2odisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDisrFieldTe2odisMask)
	}
}

// registerOdsrType Output Disable Status Register
type registerOdsrType uint32

const (
	RegisterOdsrFieldTa1odsShift = 0
	RegisterOdsrFieldTa1odsMask  = 0x1
)

// GetTa1ods Timer A Output 1 disable status
func (r *registerOdsrType) GetTa1ods() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdsrFieldTa1odsMask) != 0
}

// SetTa1ods Timer A Output 1 disable status
func (r *registerOdsrType) SetTa1ods(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdsrFieldTa1odsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdsrFieldTa1odsMask)
	}
}

const (
	RegisterOdsrFieldTa2odsShift = 1
	RegisterOdsrFieldTa2odsMask  = 0x2
)

// GetTa2ods Timer A Output 2 disable status
func (r *registerOdsrType) GetTa2ods() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdsrFieldTa2odsMask) != 0
}

// SetTa2ods Timer A Output 2 disable status
func (r *registerOdsrType) SetTa2ods(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdsrFieldTa2odsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdsrFieldTa2odsMask)
	}
}

const (
	RegisterOdsrFieldTb1odsShift = 2
	RegisterOdsrFieldTb1odsMask  = 0x4
)

// GetTb1ods Timer B Output 1 disable status
func (r *registerOdsrType) GetTb1ods() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdsrFieldTb1odsMask) != 0
}

// SetTb1ods Timer B Output 1 disable status
func (r *registerOdsrType) SetTb1ods(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdsrFieldTb1odsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdsrFieldTb1odsMask)
	}
}

const (
	RegisterOdsrFieldTb2odsShift = 3
	RegisterOdsrFieldTb2odsMask  = 0x8
)

// GetTb2ods Timer B Output 2 disable status
func (r *registerOdsrType) GetTb2ods() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdsrFieldTb2odsMask) != 0
}

// SetTb2ods Timer B Output 2 disable status
func (r *registerOdsrType) SetTb2ods(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdsrFieldTb2odsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdsrFieldTb2odsMask)
	}
}

const (
	RegisterOdsrFieldTc1odsShift = 4
	RegisterOdsrFieldTc1odsMask  = 0x10
)

// GetTc1ods Timer C Output 1 disable status
func (r *registerOdsrType) GetTc1ods() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdsrFieldTc1odsMask) != 0
}

// SetTc1ods Timer C Output 1 disable status
func (r *registerOdsrType) SetTc1ods(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdsrFieldTc1odsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdsrFieldTc1odsMask)
	}
}

const (
	RegisterOdsrFieldTc2odsShift = 5
	RegisterOdsrFieldTc2odsMask  = 0x20
)

// GetTc2ods Timer C Output 2 disable status
func (r *registerOdsrType) GetTc2ods() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdsrFieldTc2odsMask) != 0
}

// SetTc2ods Timer C Output 2 disable status
func (r *registerOdsrType) SetTc2ods(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdsrFieldTc2odsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdsrFieldTc2odsMask)
	}
}

const (
	RegisterOdsrFieldTd1odsShift = 6
	RegisterOdsrFieldTd1odsMask  = 0x40
)

// GetTd1ods Timer D Output 1 disable status
func (r *registerOdsrType) GetTd1ods() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdsrFieldTd1odsMask) != 0
}

// SetTd1ods Timer D Output 1 disable status
func (r *registerOdsrType) SetTd1ods(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdsrFieldTd1odsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdsrFieldTd1odsMask)
	}
}

const (
	RegisterOdsrFieldTd2odsShift = 7
	RegisterOdsrFieldTd2odsMask  = 0x80
)

// GetTd2ods Timer D Output 2 disable status
func (r *registerOdsrType) GetTd2ods() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdsrFieldTd2odsMask) != 0
}

// SetTd2ods Timer D Output 2 disable status
func (r *registerOdsrType) SetTd2ods(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdsrFieldTd2odsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdsrFieldTd2odsMask)
	}
}

const (
	RegisterOdsrFieldTe1odsShift = 8
	RegisterOdsrFieldTe1odsMask  = 0x100
)

// GetTe1ods Timer E Output 1 disable status
func (r *registerOdsrType) GetTe1ods() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdsrFieldTe1odsMask) != 0
}

// SetTe1ods Timer E Output 1 disable status
func (r *registerOdsrType) SetTe1ods(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdsrFieldTe1odsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdsrFieldTe1odsMask)
	}
}

const (
	RegisterOdsrFieldTe2odsShift = 9
	RegisterOdsrFieldTe2odsMask  = 0x200
)

// GetTe2ods Timer E Output 2 disable status
func (r *registerOdsrType) GetTe2ods() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOdsrFieldTe2odsMask) != 0
}

// SetTe2ods Timer E Output 2 disable status
func (r *registerOdsrType) SetTe2ods(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOdsrFieldTe2odsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOdsrFieldTe2odsMask)
	}
}

// registerBmcrType Burst Mode Control Register
type registerBmcrType uint32

const (
	RegisterBmcrFieldBmeShift = 0
	RegisterBmcrFieldBmeMask  = 0x1
)

// GetBme Burst Mode enable
func (r *registerBmcrType) GetBme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmcrFieldBmeMask) != 0
}

// SetBme Burst Mode enable
func (r *registerBmcrType) SetBme(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmcrFieldBmeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmcrFieldBmeMask)
	}
}

const (
	RegisterBmcrFieldBmomShift = 1
	RegisterBmcrFieldBmomMask  = 0x2
)

// GetBmom Burst Mode operating mode
func (r *registerBmcrType) GetBmom() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmcrFieldBmomMask) != 0
}

// SetBmom Burst Mode operating mode
func (r *registerBmcrType) SetBmom(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmcrFieldBmomMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmcrFieldBmomMask)
	}
}

const (
	RegisterBmcrFieldBmclkShift = 2
	RegisterBmcrFieldBmclkMask  = 0x3c
)

// GetBmclk Burst Mode Clock source
func (r *registerBmcrType) GetBmclk() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBmcrFieldBmclkMask) >> RegisterBmcrFieldBmclkShift)
}

// SetBmclk Burst Mode Clock source
func (r *registerBmcrType) SetBmclk(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBmcrFieldBmclkMask)|(uint32(value)<<RegisterBmcrFieldBmclkShift))
}

const (
	RegisterBmcrFieldBmprscShift = 6
	RegisterBmcrFieldBmprscMask  = 0x3c0
)

// GetBmprsc Burst Mode Prescaler
func (r *registerBmcrType) GetBmprsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBmcrFieldBmprscMask) >> RegisterBmcrFieldBmprscShift)
}

// SetBmprsc Burst Mode Prescaler
func (r *registerBmcrType) SetBmprsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBmcrFieldBmprscMask)|(uint32(value)<<RegisterBmcrFieldBmprscShift))
}

const (
	RegisterBmcrFieldBmprenShift = 10
	RegisterBmcrFieldBmprenMask  = 0x400
)

// GetBmpren Burst Mode Preload Enable
func (r *registerBmcrType) GetBmpren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmcrFieldBmprenMask) != 0
}

// SetBmpren Burst Mode Preload Enable
func (r *registerBmcrType) SetBmpren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmcrFieldBmprenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmcrFieldBmprenMask)
	}
}

const (
	RegisterBmcrFieldMtbmShift = 16
	RegisterBmcrFieldMtbmMask  = 0x10000
)

// GetMtbm Master Timer Burst Mode
func (r *registerBmcrType) GetMtbm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmcrFieldMtbmMask) != 0
}

// SetMtbm Master Timer Burst Mode
func (r *registerBmcrType) SetMtbm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmcrFieldMtbmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmcrFieldMtbmMask)
	}
}

const (
	RegisterBmcrFieldTabmShift = 17
	RegisterBmcrFieldTabmMask  = 0x20000
)

// GetTabm Timer A Burst Mode
func (r *registerBmcrType) GetTabm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmcrFieldTabmMask) != 0
}

// SetTabm Timer A Burst Mode
func (r *registerBmcrType) SetTabm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmcrFieldTabmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmcrFieldTabmMask)
	}
}

const (
	RegisterBmcrFieldTbbmShift = 18
	RegisterBmcrFieldTbbmMask  = 0x40000
)

// GetTbbm Timer B Burst Mode
func (r *registerBmcrType) GetTbbm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmcrFieldTbbmMask) != 0
}

// SetTbbm Timer B Burst Mode
func (r *registerBmcrType) SetTbbm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmcrFieldTbbmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmcrFieldTbbmMask)
	}
}

const (
	RegisterBmcrFieldTcbmShift = 19
	RegisterBmcrFieldTcbmMask  = 0x80000
)

// GetTcbm Timer C Burst Mode
func (r *registerBmcrType) GetTcbm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmcrFieldTcbmMask) != 0
}

// SetTcbm Timer C Burst Mode
func (r *registerBmcrType) SetTcbm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmcrFieldTcbmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmcrFieldTcbmMask)
	}
}

const (
	RegisterBmcrFieldTdbmShift = 20
	RegisterBmcrFieldTdbmMask  = 0x100000
)

// GetTdbm Timer D Burst Mode
func (r *registerBmcrType) GetTdbm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmcrFieldTdbmMask) != 0
}

// SetTdbm Timer D Burst Mode
func (r *registerBmcrType) SetTdbm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmcrFieldTdbmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmcrFieldTdbmMask)
	}
}

const (
	RegisterBmcrFieldTebmShift = 21
	RegisterBmcrFieldTebmMask  = 0x200000
)

// GetTebm Timer E Burst Mode
func (r *registerBmcrType) GetTebm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmcrFieldTebmMask) != 0
}

// SetTebm Timer E Burst Mode
func (r *registerBmcrType) SetTebm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmcrFieldTebmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmcrFieldTebmMask)
	}
}

const (
	RegisterBmcrFieldBmstatShift = 31
	RegisterBmcrFieldBmstatMask  = 0x80000000
)

// GetBmstat Burst Mode Status
func (r *registerBmcrType) GetBmstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmcrFieldBmstatMask) != 0
}

// SetBmstat Burst Mode Status
func (r *registerBmcrType) SetBmstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmcrFieldBmstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmcrFieldBmstatMask)
	}
}

// registerBmtrgType BMTRG
type registerBmtrgType uint32

const (
	RegisterBmtrgFieldSwShift = 0
	RegisterBmtrgFieldSwMask  = 0x1
)

// GetSw SW
func (r *registerBmtrgType) GetSw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldSwMask) != 0
}

// SetSw SW
func (r *registerBmtrgType) SetSw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldSwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldSwMask)
	}
}

const (
	RegisterBmtrgFieldMstrstShift = 1
	RegisterBmtrgFieldMstrstMask  = 0x2
)

// GetMstrst MSTRST
func (r *registerBmtrgType) GetMstrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldMstrstMask) != 0
}

// SetMstrst MSTRST
func (r *registerBmtrgType) SetMstrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldMstrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldMstrstMask)
	}
}

const (
	RegisterBmtrgFieldMstrepShift = 2
	RegisterBmtrgFieldMstrepMask  = 0x4
)

// GetMstrep MSTREP
func (r *registerBmtrgType) GetMstrep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldMstrepMask) != 0
}

// SetMstrep MSTREP
func (r *registerBmtrgType) SetMstrep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldMstrepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldMstrepMask)
	}
}

const (
	RegisterBmtrgFieldMstcmp1Shift = 3
	RegisterBmtrgFieldMstcmp1Mask  = 0x8
)

// GetMstcmp1 MSTCMP1
func (r *registerBmtrgType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerBmtrgType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldMstcmp1Mask)
	}
}

const (
	RegisterBmtrgFieldMstcmp2Shift = 4
	RegisterBmtrgFieldMstcmp2Mask  = 0x10
)

// GetMstcmp2 MSTCMP2
func (r *registerBmtrgType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerBmtrgType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldMstcmp2Mask)
	}
}

const (
	RegisterBmtrgFieldMstcmp3Shift = 5
	RegisterBmtrgFieldMstcmp3Mask  = 0x20
)

// GetMstcmp3 MSTCMP3
func (r *registerBmtrgType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerBmtrgType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldMstcmp3Mask)
	}
}

const (
	RegisterBmtrgFieldMstcmp4Shift = 6
	RegisterBmtrgFieldMstcmp4Mask  = 0x40
)

// GetMstcmp4 MSTCMP4
func (r *registerBmtrgType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerBmtrgType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldMstcmp4Mask)
	}
}

const (
	RegisterBmtrgFieldTarstShift = 7
	RegisterBmtrgFieldTarstMask  = 0x80
)

// GetTarst TARST
func (r *registerBmtrgType) GetTarst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTarstMask) != 0
}

// SetTarst TARST
func (r *registerBmtrgType) SetTarst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTarstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTarstMask)
	}
}

const (
	RegisterBmtrgFieldTarepShift = 8
	RegisterBmtrgFieldTarepMask  = 0x100
)

// GetTarep TAREP
func (r *registerBmtrgType) GetTarep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTarepMask) != 0
}

// SetTarep TAREP
func (r *registerBmtrgType) SetTarep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTarepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTarepMask)
	}
}

const (
	RegisterBmtrgFieldTacmp1Shift = 9
	RegisterBmtrgFieldTacmp1Mask  = 0x200
)

// GetTacmp1 TACMP1
func (r *registerBmtrgType) GetTacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTacmp1Mask) != 0
}

// SetTacmp1 TACMP1
func (r *registerBmtrgType) SetTacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTacmp1Mask)
	}
}

const (
	RegisterBmtrgFieldTacmp2Shift = 10
	RegisterBmtrgFieldTacmp2Mask  = 0x400
)

// GetTacmp2 TACMP2
func (r *registerBmtrgType) GetTacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTacmp2Mask) != 0
}

// SetTacmp2 TACMP2
func (r *registerBmtrgType) SetTacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTacmp2Mask)
	}
}

const (
	RegisterBmtrgFieldTbrstShift = 11
	RegisterBmtrgFieldTbrstMask  = 0x800
)

// GetTbrst TBRST
func (r *registerBmtrgType) GetTbrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTbrstMask) != 0
}

// SetTbrst TBRST
func (r *registerBmtrgType) SetTbrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTbrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTbrstMask)
	}
}

const (
	RegisterBmtrgFieldTbrepShift = 12
	RegisterBmtrgFieldTbrepMask  = 0x1000
)

// GetTbrep TBREP
func (r *registerBmtrgType) GetTbrep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTbrepMask) != 0
}

// SetTbrep TBREP
func (r *registerBmtrgType) SetTbrep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTbrepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTbrepMask)
	}
}

const (
	RegisterBmtrgFieldTbcmp1Shift = 13
	RegisterBmtrgFieldTbcmp1Mask  = 0x2000
)

// GetTbcmp1 TBCMP1
func (r *registerBmtrgType) GetTbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTbcmp1Mask) != 0
}

// SetTbcmp1 TBCMP1
func (r *registerBmtrgType) SetTbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTbcmp1Mask)
	}
}

const (
	RegisterBmtrgFieldTbcmp2Shift = 14
	RegisterBmtrgFieldTbcmp2Mask  = 0x4000
)

// GetTbcmp2 TBCMP2
func (r *registerBmtrgType) GetTbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTbcmp2Mask) != 0
}

// SetTbcmp2 TBCMP2
func (r *registerBmtrgType) SetTbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTbcmp2Mask)
	}
}

const (
	RegisterBmtrgFieldTcrstShift = 15
	RegisterBmtrgFieldTcrstMask  = 0x8000
)

// GetTcrst TCRST
func (r *registerBmtrgType) GetTcrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTcrstMask) != 0
}

// SetTcrst TCRST
func (r *registerBmtrgType) SetTcrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTcrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTcrstMask)
	}
}

const (
	RegisterBmtrgFieldTcrepShift = 16
	RegisterBmtrgFieldTcrepMask  = 0x10000
)

// GetTcrep TCREP
func (r *registerBmtrgType) GetTcrep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTcrepMask) != 0
}

// SetTcrep TCREP
func (r *registerBmtrgType) SetTcrep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTcrepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTcrepMask)
	}
}

const (
	RegisterBmtrgFieldTccmp1Shift = 17
	RegisterBmtrgFieldTccmp1Mask  = 0x20000
)

// GetTccmp1 TCCMP1
func (r *registerBmtrgType) GetTccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTccmp1Mask) != 0
}

// SetTccmp1 TCCMP1
func (r *registerBmtrgType) SetTccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTccmp1Mask)
	}
}

const (
	RegisterBmtrgFieldTccmp2Shift = 18
	RegisterBmtrgFieldTccmp2Mask  = 0x40000
)

// GetTccmp2 TCCMP2
func (r *registerBmtrgType) GetTccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTccmp2Mask) != 0
}

// SetTccmp2 TCCMP2
func (r *registerBmtrgType) SetTccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTccmp2Mask)
	}
}

const (
	RegisterBmtrgFieldTdrstShift = 19
	RegisterBmtrgFieldTdrstMask  = 0x80000
)

// GetTdrst TDRST
func (r *registerBmtrgType) GetTdrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTdrstMask) != 0
}

// SetTdrst TDRST
func (r *registerBmtrgType) SetTdrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTdrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTdrstMask)
	}
}

const (
	RegisterBmtrgFieldTdrepShift = 20
	RegisterBmtrgFieldTdrepMask  = 0x100000
)

// GetTdrep TDREP
func (r *registerBmtrgType) GetTdrep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTdrepMask) != 0
}

// SetTdrep TDREP
func (r *registerBmtrgType) SetTdrep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTdrepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTdrepMask)
	}
}

const (
	RegisterBmtrgFieldTdcmp1Shift = 21
	RegisterBmtrgFieldTdcmp1Mask  = 0x200000
)

// GetTdcmp1 TDCMP1
func (r *registerBmtrgType) GetTdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTdcmp1Mask) != 0
}

// SetTdcmp1 TDCMP1
func (r *registerBmtrgType) SetTdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTdcmp1Mask)
	}
}

const (
	RegisterBmtrgFieldTdcmp2Shift = 22
	RegisterBmtrgFieldTdcmp2Mask  = 0x400000
)

// GetTdcmp2 TDCMP2
func (r *registerBmtrgType) GetTdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTdcmp2Mask) != 0
}

// SetTdcmp2 TDCMP2
func (r *registerBmtrgType) SetTdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTdcmp2Mask)
	}
}

const (
	RegisterBmtrgFieldTerstShift = 23
	RegisterBmtrgFieldTerstMask  = 0x800000
)

// GetTerst TERST
func (r *registerBmtrgType) GetTerst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTerstMask) != 0
}

// SetTerst TERST
func (r *registerBmtrgType) SetTerst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTerstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTerstMask)
	}
}

const (
	RegisterBmtrgFieldTerepShift = 24
	RegisterBmtrgFieldTerepMask  = 0x1000000
)

// GetTerep TEREP
func (r *registerBmtrgType) GetTerep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTerepMask) != 0
}

// SetTerep TEREP
func (r *registerBmtrgType) SetTerep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTerepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTerepMask)
	}
}

const (
	RegisterBmtrgFieldTecmp1Shift = 25
	RegisterBmtrgFieldTecmp1Mask  = 0x2000000
)

// GetTecmp1 TECMP1
func (r *registerBmtrgType) GetTecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTecmp1Mask) != 0
}

// SetTecmp1 TECMP1
func (r *registerBmtrgType) SetTecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTecmp1Mask)
	}
}

const (
	RegisterBmtrgFieldTecmp2Shift = 26
	RegisterBmtrgFieldTecmp2Mask  = 0x4000000
)

// GetTecmp2 TECMP2
func (r *registerBmtrgType) GetTecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldTecmp2Mask) != 0
}

// SetTecmp2 TECMP2
func (r *registerBmtrgType) SetTecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldTecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldTecmp2Mask)
	}
}

const (
	RegisterBmtrgFieldOchpevShift = 31
	RegisterBmtrgFieldOchpevMask  = 0x80000000
)

// GetOchpev OCHPEV
func (r *registerBmtrgType) GetOchpev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBmtrgFieldOchpevMask) != 0
}

// SetOchpev OCHPEV
func (r *registerBmtrgType) SetOchpev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBmtrgFieldOchpevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBmtrgFieldOchpevMask)
	}
}

// registerBmcmpr6Type BMCMPR6
type registerBmcmpr6Type uint32

const (
	RegisterBmcmpr6FieldBmcmpShift = 0
	RegisterBmcmpr6FieldBmcmpMask  = 0xffff
)

// GetBmcmp BMCMP
func (r *registerBmcmpr6Type) GetBmcmp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBmcmpr6FieldBmcmpMask) >> RegisterBmcmpr6FieldBmcmpShift)
}

// SetBmcmp BMCMP
func (r *registerBmcmpr6Type) SetBmcmp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBmcmpr6FieldBmcmpMask)|(uint32(value)<<RegisterBmcmpr6FieldBmcmpShift))
}

// registerBmperType Burst Mode Period Register
type registerBmperType uint32

const (
	RegisterBmperFieldBmperShift = 0
	RegisterBmperFieldBmperMask  = 0xffff
)

// GetBmper Burst mode Period
func (r *registerBmperType) GetBmper() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBmperFieldBmperMask) >> RegisterBmperFieldBmperShift)
}

// SetBmper Burst mode Period
func (r *registerBmperType) SetBmper(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBmperFieldBmperMask)|(uint32(value)<<RegisterBmperFieldBmperShift))
}

// registerEecr1Type Timer External Event Control Register 1
type registerEecr1Type uint32

const (
	RegisterEecr1FieldEe1srcShift = 0
	RegisterEecr1FieldEe1srcMask  = 0x3
)

// GetEe1src External Event 1 Source
func (r *registerEecr1Type) GetEe1src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe1srcMask) >> RegisterEecr1FieldEe1srcShift)
}

// SetEe1src External Event 1 Source
func (r *registerEecr1Type) SetEe1src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe1srcMask)|(uint32(value)<<RegisterEecr1FieldEe1srcShift))
}

const (
	RegisterEecr1FieldEe1polShift = 2
	RegisterEecr1FieldEe1polMask  = 0x4
)

// GetEe1pol External Event 1 Polarity
func (r *registerEecr1Type) GetEe1pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe1polMask) != 0
}

// SetEe1pol External Event 1 Polarity
func (r *registerEecr1Type) SetEe1pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr1FieldEe1polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe1polMask)
	}
}

const (
	RegisterEecr1FieldEe1snsShift = 3
	RegisterEecr1FieldEe1snsMask  = 0x18
)

// GetEe1sns External Event 1 Sensitivity
func (r *registerEecr1Type) GetEe1sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe1snsMask) >> RegisterEecr1FieldEe1snsShift)
}

// SetEe1sns External Event 1 Sensitivity
func (r *registerEecr1Type) SetEe1sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe1snsMask)|(uint32(value)<<RegisterEecr1FieldEe1snsShift))
}

const (
	RegisterEecr1FieldEe1fastShift = 5
	RegisterEecr1FieldEe1fastMask  = 0x20
)

// GetEe1fast External Event 1 Fast mode
func (r *registerEecr1Type) GetEe1fast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe1fastMask) != 0
}

// SetEe1fast External Event 1 Fast mode
func (r *registerEecr1Type) SetEe1fast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr1FieldEe1fastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe1fastMask)
	}
}

const (
	RegisterEecr1FieldEe2srcShift = 6
	RegisterEecr1FieldEe2srcMask  = 0xc0
)

// GetEe2src External Event 2 Source
func (r *registerEecr1Type) GetEe2src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe2srcMask) >> RegisterEecr1FieldEe2srcShift)
}

// SetEe2src External Event 2 Source
func (r *registerEecr1Type) SetEe2src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe2srcMask)|(uint32(value)<<RegisterEecr1FieldEe2srcShift))
}

const (
	RegisterEecr1FieldEe2polShift = 8
	RegisterEecr1FieldEe2polMask  = 0x100
)

// GetEe2pol External Event 2 Polarity
func (r *registerEecr1Type) GetEe2pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe2polMask) != 0
}

// SetEe2pol External Event 2 Polarity
func (r *registerEecr1Type) SetEe2pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr1FieldEe2polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe2polMask)
	}
}

const (
	RegisterEecr1FieldEe2snsShift = 9
	RegisterEecr1FieldEe2snsMask  = 0x600
)

// GetEe2sns External Event 2 Sensitivity
func (r *registerEecr1Type) GetEe2sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe2snsMask) >> RegisterEecr1FieldEe2snsShift)
}

// SetEe2sns External Event 2 Sensitivity
func (r *registerEecr1Type) SetEe2sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe2snsMask)|(uint32(value)<<RegisterEecr1FieldEe2snsShift))
}

const (
	RegisterEecr1FieldEe2fastShift = 11
	RegisterEecr1FieldEe2fastMask  = 0x800
)

// GetEe2fast External Event 2 Fast mode
func (r *registerEecr1Type) GetEe2fast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe2fastMask) != 0
}

// SetEe2fast External Event 2 Fast mode
func (r *registerEecr1Type) SetEe2fast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr1FieldEe2fastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe2fastMask)
	}
}

const (
	RegisterEecr1FieldEe3srcShift = 12
	RegisterEecr1FieldEe3srcMask  = 0x3000
)

// GetEe3src External Event 3 Source
func (r *registerEecr1Type) GetEe3src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe3srcMask) >> RegisterEecr1FieldEe3srcShift)
}

// SetEe3src External Event 3 Source
func (r *registerEecr1Type) SetEe3src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe3srcMask)|(uint32(value)<<RegisterEecr1FieldEe3srcShift))
}

const (
	RegisterEecr1FieldEe3polShift = 14
	RegisterEecr1FieldEe3polMask  = 0x4000
)

// GetEe3pol External Event 3 Polarity
func (r *registerEecr1Type) GetEe3pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe3polMask) != 0
}

// SetEe3pol External Event 3 Polarity
func (r *registerEecr1Type) SetEe3pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr1FieldEe3polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe3polMask)
	}
}

const (
	RegisterEecr1FieldEe3snsShift = 15
	RegisterEecr1FieldEe3snsMask  = 0x18000
)

// GetEe3sns External Event 3 Sensitivity
func (r *registerEecr1Type) GetEe3sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe3snsMask) >> RegisterEecr1FieldEe3snsShift)
}

// SetEe3sns External Event 3 Sensitivity
func (r *registerEecr1Type) SetEe3sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe3snsMask)|(uint32(value)<<RegisterEecr1FieldEe3snsShift))
}

const (
	RegisterEecr1FieldEe3fastShift = 17
	RegisterEecr1FieldEe3fastMask  = 0x20000
)

// GetEe3fast External Event 3 Fast mode
func (r *registerEecr1Type) GetEe3fast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe3fastMask) != 0
}

// SetEe3fast External Event 3 Fast mode
func (r *registerEecr1Type) SetEe3fast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr1FieldEe3fastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe3fastMask)
	}
}

const (
	RegisterEecr1FieldEe4srcShift = 18
	RegisterEecr1FieldEe4srcMask  = 0xc0000
)

// GetEe4src External Event 4 Source
func (r *registerEecr1Type) GetEe4src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe4srcMask) >> RegisterEecr1FieldEe4srcShift)
}

// SetEe4src External Event 4 Source
func (r *registerEecr1Type) SetEe4src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe4srcMask)|(uint32(value)<<RegisterEecr1FieldEe4srcShift))
}

const (
	RegisterEecr1FieldEe4polShift = 20
	RegisterEecr1FieldEe4polMask  = 0x100000
)

// GetEe4pol External Event 4 Polarity
func (r *registerEecr1Type) GetEe4pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe4polMask) != 0
}

// SetEe4pol External Event 4 Polarity
func (r *registerEecr1Type) SetEe4pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr1FieldEe4polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe4polMask)
	}
}

const (
	RegisterEecr1FieldEe4snsShift = 21
	RegisterEecr1FieldEe4snsMask  = 0x600000
)

// GetEe4sns External Event 4 Sensitivity
func (r *registerEecr1Type) GetEe4sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe4snsMask) >> RegisterEecr1FieldEe4snsShift)
}

// SetEe4sns External Event 4 Sensitivity
func (r *registerEecr1Type) SetEe4sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe4snsMask)|(uint32(value)<<RegisterEecr1FieldEe4snsShift))
}

const (
	RegisterEecr1FieldEe4fastShift = 23
	RegisterEecr1FieldEe4fastMask  = 0x800000
)

// GetEe4fast External Event 4 Fast mode
func (r *registerEecr1Type) GetEe4fast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe4fastMask) != 0
}

// SetEe4fast External Event 4 Fast mode
func (r *registerEecr1Type) SetEe4fast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr1FieldEe4fastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe4fastMask)
	}
}

const (
	RegisterEecr1FieldEe5srcShift = 24
	RegisterEecr1FieldEe5srcMask  = 0x3000000
)

// GetEe5src External Event 5 Source
func (r *registerEecr1Type) GetEe5src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe5srcMask) >> RegisterEecr1FieldEe5srcShift)
}

// SetEe5src External Event 5 Source
func (r *registerEecr1Type) SetEe5src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe5srcMask)|(uint32(value)<<RegisterEecr1FieldEe5srcShift))
}

const (
	RegisterEecr1FieldEe5polShift = 26
	RegisterEecr1FieldEe5polMask  = 0x4000000
)

// GetEe5pol External Event 5 Polarity
func (r *registerEecr1Type) GetEe5pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe5polMask) != 0
}

// SetEe5pol External Event 5 Polarity
func (r *registerEecr1Type) SetEe5pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr1FieldEe5polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe5polMask)
	}
}

const (
	RegisterEecr1FieldEe5snsShift = 27
	RegisterEecr1FieldEe5snsMask  = 0x18000000
)

// GetEe5sns External Event 5 Sensitivity
func (r *registerEecr1Type) GetEe5sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe5snsMask) >> RegisterEecr1FieldEe5snsShift)
}

// SetEe5sns External Event 5 Sensitivity
func (r *registerEecr1Type) SetEe5sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe5snsMask)|(uint32(value)<<RegisterEecr1FieldEe5snsShift))
}

const (
	RegisterEecr1FieldEe5fastShift = 29
	RegisterEecr1FieldEe5fastMask  = 0x20000000
)

// GetEe5fast External Event 5 Fast mode
func (r *registerEecr1Type) GetEe5fast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr1FieldEe5fastMask) != 0
}

// SetEe5fast External Event 5 Fast mode
func (r *registerEecr1Type) SetEe5fast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr1FieldEe5fastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr1FieldEe5fastMask)
	}
}

// registerEecr2Type Timer External Event Control Register 2
type registerEecr2Type uint32

const (
	RegisterEecr2FieldEe6srcShift = 0
	RegisterEecr2FieldEe6srcMask  = 0x3
)

// GetEe6src External Event 6 Source
func (r *registerEecr2Type) GetEe6src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe6srcMask) >> RegisterEecr2FieldEe6srcShift)
}

// SetEe6src External Event 6 Source
func (r *registerEecr2Type) SetEe6src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe6srcMask)|(uint32(value)<<RegisterEecr2FieldEe6srcShift))
}

const (
	RegisterEecr2FieldEe6polShift = 2
	RegisterEecr2FieldEe6polMask  = 0x4
)

// GetEe6pol External Event 6 Polarity
func (r *registerEecr2Type) GetEe6pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe6polMask) != 0
}

// SetEe6pol External Event 6 Polarity
func (r *registerEecr2Type) SetEe6pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr2FieldEe6polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe6polMask)
	}
}

const (
	RegisterEecr2FieldEe6snsShift = 3
	RegisterEecr2FieldEe6snsMask  = 0x18
)

// GetEe6sns External Event 6 Sensitivity
func (r *registerEecr2Type) GetEe6sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe6snsMask) >> RegisterEecr2FieldEe6snsShift)
}

// SetEe6sns External Event 6 Sensitivity
func (r *registerEecr2Type) SetEe6sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe6snsMask)|(uint32(value)<<RegisterEecr2FieldEe6snsShift))
}

const (
	RegisterEecr2FieldEe7srcShift = 6
	RegisterEecr2FieldEe7srcMask  = 0xc0
)

// GetEe7src External Event 7 Source
func (r *registerEecr2Type) GetEe7src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe7srcMask) >> RegisterEecr2FieldEe7srcShift)
}

// SetEe7src External Event 7 Source
func (r *registerEecr2Type) SetEe7src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe7srcMask)|(uint32(value)<<RegisterEecr2FieldEe7srcShift))
}

const (
	RegisterEecr2FieldEe7polShift = 8
	RegisterEecr2FieldEe7polMask  = 0x100
)

// GetEe7pol External Event 7 Polarity
func (r *registerEecr2Type) GetEe7pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe7polMask) != 0
}

// SetEe7pol External Event 7 Polarity
func (r *registerEecr2Type) SetEe7pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr2FieldEe7polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe7polMask)
	}
}

const (
	RegisterEecr2FieldEe7snsShift = 9
	RegisterEecr2FieldEe7snsMask  = 0x600
)

// GetEe7sns External Event 7 Sensitivity
func (r *registerEecr2Type) GetEe7sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe7snsMask) >> RegisterEecr2FieldEe7snsShift)
}

// SetEe7sns External Event 7 Sensitivity
func (r *registerEecr2Type) SetEe7sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe7snsMask)|(uint32(value)<<RegisterEecr2FieldEe7snsShift))
}

const (
	RegisterEecr2FieldEe8srcShift = 12
	RegisterEecr2FieldEe8srcMask  = 0x3000
)

// GetEe8src External Event 8 Source
func (r *registerEecr2Type) GetEe8src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe8srcMask) >> RegisterEecr2FieldEe8srcShift)
}

// SetEe8src External Event 8 Source
func (r *registerEecr2Type) SetEe8src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe8srcMask)|(uint32(value)<<RegisterEecr2FieldEe8srcShift))
}

const (
	RegisterEecr2FieldEe8polShift = 14
	RegisterEecr2FieldEe8polMask  = 0x4000
)

// GetEe8pol External Event 8 Polarity
func (r *registerEecr2Type) GetEe8pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe8polMask) != 0
}

// SetEe8pol External Event 8 Polarity
func (r *registerEecr2Type) SetEe8pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr2FieldEe8polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe8polMask)
	}
}

const (
	RegisterEecr2FieldEe8snsShift = 15
	RegisterEecr2FieldEe8snsMask  = 0x18000
)

// GetEe8sns External Event 8 Sensitivity
func (r *registerEecr2Type) GetEe8sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe8snsMask) >> RegisterEecr2FieldEe8snsShift)
}

// SetEe8sns External Event 8 Sensitivity
func (r *registerEecr2Type) SetEe8sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe8snsMask)|(uint32(value)<<RegisterEecr2FieldEe8snsShift))
}

const (
	RegisterEecr2FieldEe9srcShift = 18
	RegisterEecr2FieldEe9srcMask  = 0xc0000
)

// GetEe9src External Event 9 Source
func (r *registerEecr2Type) GetEe9src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe9srcMask) >> RegisterEecr2FieldEe9srcShift)
}

// SetEe9src External Event 9 Source
func (r *registerEecr2Type) SetEe9src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe9srcMask)|(uint32(value)<<RegisterEecr2FieldEe9srcShift))
}

const (
	RegisterEecr2FieldEe9polShift = 20
	RegisterEecr2FieldEe9polMask  = 0x100000
)

// GetEe9pol External Event 9 Polarity
func (r *registerEecr2Type) GetEe9pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe9polMask) != 0
}

// SetEe9pol External Event 9 Polarity
func (r *registerEecr2Type) SetEe9pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr2FieldEe9polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe9polMask)
	}
}

const (
	RegisterEecr2FieldEe9snsShift = 21
	RegisterEecr2FieldEe9snsMask  = 0x600000
)

// GetEe9sns External Event 9 Sensitivity
func (r *registerEecr2Type) GetEe9sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe9snsMask) >> RegisterEecr2FieldEe9snsShift)
}

// SetEe9sns External Event 9 Sensitivity
func (r *registerEecr2Type) SetEe9sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe9snsMask)|(uint32(value)<<RegisterEecr2FieldEe9snsShift))
}

const (
	RegisterEecr2FieldEe10srcShift = 24
	RegisterEecr2FieldEe10srcMask  = 0x3000000
)

// GetEe10src External Event 10 Source
func (r *registerEecr2Type) GetEe10src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe10srcMask) >> RegisterEecr2FieldEe10srcShift)
}

// SetEe10src External Event 10 Source
func (r *registerEecr2Type) SetEe10src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe10srcMask)|(uint32(value)<<RegisterEecr2FieldEe10srcShift))
}

const (
	RegisterEecr2FieldEe10polShift = 26
	RegisterEecr2FieldEe10polMask  = 0x4000000
)

// GetEe10pol External Event 10 Polarity
func (r *registerEecr2Type) GetEe10pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe10polMask) != 0
}

// SetEe10pol External Event 10 Polarity
func (r *registerEecr2Type) SetEe10pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr2FieldEe10polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe10polMask)
	}
}

const (
	RegisterEecr2FieldEe10snsShift = 27
	RegisterEecr2FieldEe10snsMask  = 0x18000000
)

// GetEe10sns External Event 10 Sensitivity
func (r *registerEecr2Type) GetEe10sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr2FieldEe10snsMask) >> RegisterEecr2FieldEe10snsShift)
}

// SetEe10sns External Event 10 Sensitivity
func (r *registerEecr2Type) SetEe10sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr2FieldEe10snsMask)|(uint32(value)<<RegisterEecr2FieldEe10snsShift))
}

// registerEecr3Type Timer External Event Control Register 3
type registerEecr3Type uint32

const (
	RegisterEecr3FieldEe6srcShift = 0
	RegisterEecr3FieldEe6srcMask  = 0x3
)

// GetEe6src EE6SRC
func (r *registerEecr3Type) GetEe6src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe6srcMask) >> RegisterEecr3FieldEe6srcShift)
}

// SetEe6src EE6SRC
func (r *registerEecr3Type) SetEe6src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe6srcMask)|(uint32(value)<<RegisterEecr3FieldEe6srcShift))
}

const (
	RegisterEecr3FieldEe6polShift = 2
	RegisterEecr3FieldEe6polMask  = 0x4
)

// GetEe6pol EE6POL
func (r *registerEecr3Type) GetEe6pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe6polMask) != 0
}

// SetEe6pol EE6POL
func (r *registerEecr3Type) SetEe6pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr3FieldEe6polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe6polMask)
	}
}

const (
	RegisterEecr3FieldEe6snsShift = 3
	RegisterEecr3FieldEe6snsMask  = 0x18
)

// GetEe6sns EE6SNS
func (r *registerEecr3Type) GetEe6sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe6snsMask) >> RegisterEecr3FieldEe6snsShift)
}

// SetEe6sns EE6SNS
func (r *registerEecr3Type) SetEe6sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe6snsMask)|(uint32(value)<<RegisterEecr3FieldEe6snsShift))
}

const (
	RegisterEecr3FieldEe7srcShift = 6
	RegisterEecr3FieldEe7srcMask  = 0xc0
)

// GetEe7src EE7SRC
func (r *registerEecr3Type) GetEe7src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe7srcMask) >> RegisterEecr3FieldEe7srcShift)
}

// SetEe7src EE7SRC
func (r *registerEecr3Type) SetEe7src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe7srcMask)|(uint32(value)<<RegisterEecr3FieldEe7srcShift))
}

const (
	RegisterEecr3FieldEe7polShift = 8
	RegisterEecr3FieldEe7polMask  = 0x100
)

// GetEe7pol EE7POL
func (r *registerEecr3Type) GetEe7pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe7polMask) != 0
}

// SetEe7pol EE7POL
func (r *registerEecr3Type) SetEe7pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr3FieldEe7polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe7polMask)
	}
}

const (
	RegisterEecr3FieldEe7snsShift = 9
	RegisterEecr3FieldEe7snsMask  = 0x600
)

// GetEe7sns EE7SNS
func (r *registerEecr3Type) GetEe7sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe7snsMask) >> RegisterEecr3FieldEe7snsShift)
}

// SetEe7sns EE7SNS
func (r *registerEecr3Type) SetEe7sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe7snsMask)|(uint32(value)<<RegisterEecr3FieldEe7snsShift))
}

const (
	RegisterEecr3FieldEe8srcShift = 12
	RegisterEecr3FieldEe8srcMask  = 0x3000
)

// GetEe8src EE8SRC
func (r *registerEecr3Type) GetEe8src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe8srcMask) >> RegisterEecr3FieldEe8srcShift)
}

// SetEe8src EE8SRC
func (r *registerEecr3Type) SetEe8src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe8srcMask)|(uint32(value)<<RegisterEecr3FieldEe8srcShift))
}

const (
	RegisterEecr3FieldEe8polShift = 14
	RegisterEecr3FieldEe8polMask  = 0x4000
)

// GetEe8pol EE8POL
func (r *registerEecr3Type) GetEe8pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe8polMask) != 0
}

// SetEe8pol EE8POL
func (r *registerEecr3Type) SetEe8pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr3FieldEe8polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe8polMask)
	}
}

const (
	RegisterEecr3FieldEe8snsShift = 15
	RegisterEecr3FieldEe8snsMask  = 0x18000
)

// GetEe8sns EE8SNS
func (r *registerEecr3Type) GetEe8sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe8snsMask) >> RegisterEecr3FieldEe8snsShift)
}

// SetEe8sns EE8SNS
func (r *registerEecr3Type) SetEe8sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe8snsMask)|(uint32(value)<<RegisterEecr3FieldEe8snsShift))
}

const (
	RegisterEecr3FieldEe9srcShift = 18
	RegisterEecr3FieldEe9srcMask  = 0xc0000
)

// GetEe9src EE9SRC
func (r *registerEecr3Type) GetEe9src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe9srcMask) >> RegisterEecr3FieldEe9srcShift)
}

// SetEe9src EE9SRC
func (r *registerEecr3Type) SetEe9src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe9srcMask)|(uint32(value)<<RegisterEecr3FieldEe9srcShift))
}

const (
	RegisterEecr3FieldEe9polShift = 20
	RegisterEecr3FieldEe9polMask  = 0x100000
)

// GetEe9pol EE9POL
func (r *registerEecr3Type) GetEe9pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe9polMask) != 0
}

// SetEe9pol EE9POL
func (r *registerEecr3Type) SetEe9pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr3FieldEe9polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe9polMask)
	}
}

const (
	RegisterEecr3FieldEe9snsShift = 21
	RegisterEecr3FieldEe9snsMask  = 0x600000
)

// GetEe9sns EE9SNS
func (r *registerEecr3Type) GetEe9sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe9snsMask) >> RegisterEecr3FieldEe9snsShift)
}

// SetEe9sns EE9SNS
func (r *registerEecr3Type) SetEe9sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe9snsMask)|(uint32(value)<<RegisterEecr3FieldEe9snsShift))
}

const (
	RegisterEecr3FieldEe10srcShift = 24
	RegisterEecr3FieldEe10srcMask  = 0x3000000
)

// GetEe10src EE10SRC
func (r *registerEecr3Type) GetEe10src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe10srcMask) >> RegisterEecr3FieldEe10srcShift)
}

// SetEe10src EE10SRC
func (r *registerEecr3Type) SetEe10src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe10srcMask)|(uint32(value)<<RegisterEecr3FieldEe10srcShift))
}

const (
	RegisterEecr3FieldEe10polShift = 26
	RegisterEecr3FieldEe10polMask  = 0x4000000
)

// GetEe10pol EE10POL
func (r *registerEecr3Type) GetEe10pol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe10polMask) != 0
}

// SetEe10pol EE10POL
func (r *registerEecr3Type) SetEe10pol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEecr3FieldEe10polMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe10polMask)
	}
}

const (
	RegisterEecr3FieldEe10snsShift = 27
	RegisterEecr3FieldEe10snsMask  = 0x18000000
)

// GetEe10sns EE10SNS
func (r *registerEecr3Type) GetEe10sns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEecr3FieldEe10snsMask) >> RegisterEecr3FieldEe10snsShift)
}

// SetEe10sns EE10SNS
func (r *registerEecr3Type) SetEe10sns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEecr3FieldEe10snsMask)|(uint32(value)<<RegisterEecr3FieldEe10snsShift))
}

// registerAdc1rType ADC Trigger 1 Register
type registerAdc1rType uint32

const (
	RegisterAdc1rFieldAd1mc1Shift = 0
	RegisterAdc1rFieldAd1mc1Mask  = 0x1
)

// GetAd1mc1 ADC trigger 1 on Master Compare 1
func (r *registerAdc1rType) GetAd1mc1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1mc1Mask) != 0
}

// SetAd1mc1 ADC trigger 1 on Master Compare 1
func (r *registerAdc1rType) SetAd1mc1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1mc1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1mc1Mask)
	}
}

const (
	RegisterAdc1rFieldAd1mc2Shift = 1
	RegisterAdc1rFieldAd1mc2Mask  = 0x2
)

// GetAd1mc2 ADC trigger 1 on Master Compare 2
func (r *registerAdc1rType) GetAd1mc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1mc2Mask) != 0
}

// SetAd1mc2 ADC trigger 1 on Master Compare 2
func (r *registerAdc1rType) SetAd1mc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1mc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1mc2Mask)
	}
}

const (
	RegisterAdc1rFieldAd1mc3Shift = 2
	RegisterAdc1rFieldAd1mc3Mask  = 0x4
)

// GetAd1mc3 ADC trigger 1 on Master Compare 3
func (r *registerAdc1rType) GetAd1mc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1mc3Mask) != 0
}

// SetAd1mc3 ADC trigger 1 on Master Compare 3
func (r *registerAdc1rType) SetAd1mc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1mc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1mc3Mask)
	}
}

const (
	RegisterAdc1rFieldAd1mc4Shift = 3
	RegisterAdc1rFieldAd1mc4Mask  = 0x8
)

// GetAd1mc4 ADC trigger 1 on Master Compare 4
func (r *registerAdc1rType) GetAd1mc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1mc4Mask) != 0
}

// SetAd1mc4 ADC trigger 1 on Master Compare 4
func (r *registerAdc1rType) SetAd1mc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1mc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1mc4Mask)
	}
}

const (
	RegisterAdc1rFieldAd1mperShift = 4
	RegisterAdc1rFieldAd1mperMask  = 0x10
)

// GetAd1mper ADC trigger 1 on Master Period
func (r *registerAdc1rType) GetAd1mper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1mperMask) != 0
}

// SetAd1mper ADC trigger 1 on Master Period
func (r *registerAdc1rType) SetAd1mper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1mperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1mperMask)
	}
}

const (
	RegisterAdc1rFieldAd1eev1Shift = 5
	RegisterAdc1rFieldAd1eev1Mask  = 0x20
)

// GetAd1eev1 ADC trigger 1 on External Event 1
func (r *registerAdc1rType) GetAd1eev1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1eev1Mask) != 0
}

// SetAd1eev1 ADC trigger 1 on External Event 1
func (r *registerAdc1rType) SetAd1eev1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1eev1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1eev1Mask)
	}
}

const (
	RegisterAdc1rFieldAd1eev2Shift = 6
	RegisterAdc1rFieldAd1eev2Mask  = 0x40
)

// GetAd1eev2 ADC trigger 1 on External Event 2
func (r *registerAdc1rType) GetAd1eev2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1eev2Mask) != 0
}

// SetAd1eev2 ADC trigger 1 on External Event 2
func (r *registerAdc1rType) SetAd1eev2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1eev2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1eev2Mask)
	}
}

const (
	RegisterAdc1rFieldAd1eev3Shift = 7
	RegisterAdc1rFieldAd1eev3Mask  = 0x80
)

// GetAd1eev3 ADC trigger 1 on External Event 3
func (r *registerAdc1rType) GetAd1eev3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1eev3Mask) != 0
}

// SetAd1eev3 ADC trigger 1 on External Event 3
func (r *registerAdc1rType) SetAd1eev3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1eev3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1eev3Mask)
	}
}

const (
	RegisterAdc1rFieldAd1eev4Shift = 8
	RegisterAdc1rFieldAd1eev4Mask  = 0x100
)

// GetAd1eev4 ADC trigger 1 on External Event 4
func (r *registerAdc1rType) GetAd1eev4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1eev4Mask) != 0
}

// SetAd1eev4 ADC trigger 1 on External Event 4
func (r *registerAdc1rType) SetAd1eev4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1eev4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1eev4Mask)
	}
}

const (
	RegisterAdc1rFieldAd1eev5Shift = 9
	RegisterAdc1rFieldAd1eev5Mask  = 0x200
)

// GetAd1eev5 ADC trigger 1 on External Event 5
func (r *registerAdc1rType) GetAd1eev5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1eev5Mask) != 0
}

// SetAd1eev5 ADC trigger 1 on External Event 5
func (r *registerAdc1rType) SetAd1eev5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1eev5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1eev5Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tac2Shift = 10
	RegisterAdc1rFieldAd1tac2Mask  = 0x400
)

// GetAd1tac2 ADC trigger 1 on Timer A compare 2
func (r *registerAdc1rType) GetAd1tac2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tac2Mask) != 0
}

// SetAd1tac2 ADC trigger 1 on Timer A compare 2
func (r *registerAdc1rType) SetAd1tac2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tac2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tac2Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tac3Shift = 11
	RegisterAdc1rFieldAd1tac3Mask  = 0x800
)

// GetAd1tac3 ADC trigger 1 on Timer A compare 3
func (r *registerAdc1rType) GetAd1tac3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tac3Mask) != 0
}

// SetAd1tac3 ADC trigger 1 on Timer A compare 3
func (r *registerAdc1rType) SetAd1tac3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tac3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tac3Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tac4Shift = 12
	RegisterAdc1rFieldAd1tac4Mask  = 0x1000
)

// GetAd1tac4 ADC trigger 1 on Timer A compare 4
func (r *registerAdc1rType) GetAd1tac4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tac4Mask) != 0
}

// SetAd1tac4 ADC trigger 1 on Timer A compare 4
func (r *registerAdc1rType) SetAd1tac4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tac4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tac4Mask)
	}
}

const (
	RegisterAdc1rFieldAd1taperShift = 13
	RegisterAdc1rFieldAd1taperMask  = 0x2000
)

// GetAd1taper ADC trigger 1 on Timer A Period
func (r *registerAdc1rType) GetAd1taper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1taperMask) != 0
}

// SetAd1taper ADC trigger 1 on Timer A Period
func (r *registerAdc1rType) SetAd1taper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1taperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1taperMask)
	}
}

const (
	RegisterAdc1rFieldAd1tarstShift = 14
	RegisterAdc1rFieldAd1tarstMask  = 0x4000
)

// GetAd1tarst ADC trigger 1 on Timer A Reset
func (r *registerAdc1rType) GetAd1tarst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tarstMask) != 0
}

// SetAd1tarst ADC trigger 1 on Timer A Reset
func (r *registerAdc1rType) SetAd1tarst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tarstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tarstMask)
	}
}

const (
	RegisterAdc1rFieldAd1tbc2Shift = 15
	RegisterAdc1rFieldAd1tbc2Mask  = 0x8000
)

// GetAd1tbc2 ADC trigger 1 on Timer B compare 2
func (r *registerAdc1rType) GetAd1tbc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tbc2Mask) != 0
}

// SetAd1tbc2 ADC trigger 1 on Timer B compare 2
func (r *registerAdc1rType) SetAd1tbc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tbc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tbc2Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tbc3Shift = 16
	RegisterAdc1rFieldAd1tbc3Mask  = 0x10000
)

// GetAd1tbc3 ADC trigger 1 on Timer B compare 3
func (r *registerAdc1rType) GetAd1tbc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tbc3Mask) != 0
}

// SetAd1tbc3 ADC trigger 1 on Timer B compare 3
func (r *registerAdc1rType) SetAd1tbc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tbc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tbc3Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tbc4Shift = 17
	RegisterAdc1rFieldAd1tbc4Mask  = 0x20000
)

// GetAd1tbc4 ADC trigger 1 on Timer B compare 4
func (r *registerAdc1rType) GetAd1tbc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tbc4Mask) != 0
}

// SetAd1tbc4 ADC trigger 1 on Timer B compare 4
func (r *registerAdc1rType) SetAd1tbc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tbc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tbc4Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tbperShift = 18
	RegisterAdc1rFieldAd1tbperMask  = 0x40000
)

// GetAd1tbper ADC trigger 1 on Timer B Period
func (r *registerAdc1rType) GetAd1tbper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tbperMask) != 0
}

// SetAd1tbper ADC trigger 1 on Timer B Period
func (r *registerAdc1rType) SetAd1tbper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tbperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tbperMask)
	}
}

const (
	RegisterAdc1rFieldAd1tbrstShift = 19
	RegisterAdc1rFieldAd1tbrstMask  = 0x80000
)

// GetAd1tbrst ADC trigger 1 on Timer B Reset
func (r *registerAdc1rType) GetAd1tbrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tbrstMask) != 0
}

// SetAd1tbrst ADC trigger 1 on Timer B Reset
func (r *registerAdc1rType) SetAd1tbrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tbrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tbrstMask)
	}
}

const (
	RegisterAdc1rFieldAd1tcc2Shift = 20
	RegisterAdc1rFieldAd1tcc2Mask  = 0x100000
)

// GetAd1tcc2 ADC trigger 1 on Timer C compare 2
func (r *registerAdc1rType) GetAd1tcc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tcc2Mask) != 0
}

// SetAd1tcc2 ADC trigger 1 on Timer C compare 2
func (r *registerAdc1rType) SetAd1tcc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tcc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tcc2Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tcc3Shift = 21
	RegisterAdc1rFieldAd1tcc3Mask  = 0x200000
)

// GetAd1tcc3 ADC trigger 1 on Timer C compare 3
func (r *registerAdc1rType) GetAd1tcc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tcc3Mask) != 0
}

// SetAd1tcc3 ADC trigger 1 on Timer C compare 3
func (r *registerAdc1rType) SetAd1tcc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tcc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tcc3Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tcc4Shift = 22
	RegisterAdc1rFieldAd1tcc4Mask  = 0x400000
)

// GetAd1tcc4 ADC trigger 1 on Timer C compare 4
func (r *registerAdc1rType) GetAd1tcc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tcc4Mask) != 0
}

// SetAd1tcc4 ADC trigger 1 on Timer C compare 4
func (r *registerAdc1rType) SetAd1tcc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tcc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tcc4Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tcperShift = 23
	RegisterAdc1rFieldAd1tcperMask  = 0x800000
)

// GetAd1tcper ADC trigger 1 on Timer C Period
func (r *registerAdc1rType) GetAd1tcper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tcperMask) != 0
}

// SetAd1tcper ADC trigger 1 on Timer C Period
func (r *registerAdc1rType) SetAd1tcper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tcperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tcperMask)
	}
}

const (
	RegisterAdc1rFieldAd1tdc2Shift = 24
	RegisterAdc1rFieldAd1tdc2Mask  = 0x1000000
)

// GetAd1tdc2 ADC trigger 1 on Timer D compare 2
func (r *registerAdc1rType) GetAd1tdc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tdc2Mask) != 0
}

// SetAd1tdc2 ADC trigger 1 on Timer D compare 2
func (r *registerAdc1rType) SetAd1tdc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tdc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tdc2Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tdc3Shift = 25
	RegisterAdc1rFieldAd1tdc3Mask  = 0x2000000
)

// GetAd1tdc3 ADC trigger 1 on Timer D compare 3
func (r *registerAdc1rType) GetAd1tdc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tdc3Mask) != 0
}

// SetAd1tdc3 ADC trigger 1 on Timer D compare 3
func (r *registerAdc1rType) SetAd1tdc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tdc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tdc3Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tdc4Shift = 26
	RegisterAdc1rFieldAd1tdc4Mask  = 0x4000000
)

// GetAd1tdc4 ADC trigger 1 on Timer D compare 4
func (r *registerAdc1rType) GetAd1tdc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tdc4Mask) != 0
}

// SetAd1tdc4 ADC trigger 1 on Timer D compare 4
func (r *registerAdc1rType) SetAd1tdc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tdc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tdc4Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tdperShift = 27
	RegisterAdc1rFieldAd1tdperMask  = 0x8000000
)

// GetAd1tdper ADC trigger 1 on Timer D Period
func (r *registerAdc1rType) GetAd1tdper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tdperMask) != 0
}

// SetAd1tdper ADC trigger 1 on Timer D Period
func (r *registerAdc1rType) SetAd1tdper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tdperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tdperMask)
	}
}

const (
	RegisterAdc1rFieldAd1tec2Shift = 28
	RegisterAdc1rFieldAd1tec2Mask  = 0x10000000
)

// GetAd1tec2 ADC trigger 1 on Timer E compare 2
func (r *registerAdc1rType) GetAd1tec2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tec2Mask) != 0
}

// SetAd1tec2 ADC trigger 1 on Timer E compare 2
func (r *registerAdc1rType) SetAd1tec2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tec2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tec2Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tec3Shift = 29
	RegisterAdc1rFieldAd1tec3Mask  = 0x20000000
)

// GetAd1tec3 ADC trigger 1 on Timer E compare 3
func (r *registerAdc1rType) GetAd1tec3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tec3Mask) != 0
}

// SetAd1tec3 ADC trigger 1 on Timer E compare 3
func (r *registerAdc1rType) SetAd1tec3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tec3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tec3Mask)
	}
}

const (
	RegisterAdc1rFieldAd1tec4Shift = 30
	RegisterAdc1rFieldAd1tec4Mask  = 0x40000000
)

// GetAd1tec4 ADC trigger 1 on Timer E compare 4
func (r *registerAdc1rType) GetAd1tec4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1tec4Mask) != 0
}

// SetAd1tec4 ADC trigger 1 on Timer E compare 4
func (r *registerAdc1rType) SetAd1tec4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1tec4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1tec4Mask)
	}
}

const (
	RegisterAdc1rFieldAd1teperShift = 31
	RegisterAdc1rFieldAd1teperMask  = 0x80000000
)

// GetAd1teper ADC trigger 1 on Timer E Period
func (r *registerAdc1rType) GetAd1teper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc1rFieldAd1teperMask) != 0
}

// SetAd1teper ADC trigger 1 on Timer E Period
func (r *registerAdc1rType) SetAd1teper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc1rFieldAd1teperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc1rFieldAd1teperMask)
	}
}

// registerAdc2rType ADC Trigger 2 Register
type registerAdc2rType uint32

const (
	RegisterAdc2rFieldAd2mc1Shift = 0
	RegisterAdc2rFieldAd2mc1Mask  = 0x1
)

// GetAd2mc1 ADC trigger 2 on Master Compare 1
func (r *registerAdc2rType) GetAd2mc1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2mc1Mask) != 0
}

// SetAd2mc1 ADC trigger 2 on Master Compare 1
func (r *registerAdc2rType) SetAd2mc1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2mc1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2mc1Mask)
	}
}

const (
	RegisterAdc2rFieldAd2mc2Shift = 1
	RegisterAdc2rFieldAd2mc2Mask  = 0x2
)

// GetAd2mc2 ADC trigger 2 on Master Compare 2
func (r *registerAdc2rType) GetAd2mc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2mc2Mask) != 0
}

// SetAd2mc2 ADC trigger 2 on Master Compare 2
func (r *registerAdc2rType) SetAd2mc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2mc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2mc2Mask)
	}
}

const (
	RegisterAdc2rFieldAd2mc3Shift = 2
	RegisterAdc2rFieldAd2mc3Mask  = 0x4
)

// GetAd2mc3 ADC trigger 2 on Master Compare 3
func (r *registerAdc2rType) GetAd2mc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2mc3Mask) != 0
}

// SetAd2mc3 ADC trigger 2 on Master Compare 3
func (r *registerAdc2rType) SetAd2mc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2mc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2mc3Mask)
	}
}

const (
	RegisterAdc2rFieldAd2mc4Shift = 3
	RegisterAdc2rFieldAd2mc4Mask  = 0x8
)

// GetAd2mc4 ADC trigger 2 on Master Compare 4
func (r *registerAdc2rType) GetAd2mc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2mc4Mask) != 0
}

// SetAd2mc4 ADC trigger 2 on Master Compare 4
func (r *registerAdc2rType) SetAd2mc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2mc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2mc4Mask)
	}
}

const (
	RegisterAdc2rFieldAd2mperShift = 4
	RegisterAdc2rFieldAd2mperMask  = 0x10
)

// GetAd2mper ADC trigger 2 on Master Period
func (r *registerAdc2rType) GetAd2mper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2mperMask) != 0
}

// SetAd2mper ADC trigger 2 on Master Period
func (r *registerAdc2rType) SetAd2mper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2mperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2mperMask)
	}
}

const (
	RegisterAdc2rFieldAd2eev6Shift = 5
	RegisterAdc2rFieldAd2eev6Mask  = 0x20
)

// GetAd2eev6 ADC trigger 2 on External Event 6
func (r *registerAdc2rType) GetAd2eev6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2eev6Mask) != 0
}

// SetAd2eev6 ADC trigger 2 on External Event 6
func (r *registerAdc2rType) SetAd2eev6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2eev6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2eev6Mask)
	}
}

const (
	RegisterAdc2rFieldAd2eev7Shift = 6
	RegisterAdc2rFieldAd2eev7Mask  = 0x40
)

// GetAd2eev7 ADC trigger 2 on External Event 7
func (r *registerAdc2rType) GetAd2eev7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2eev7Mask) != 0
}

// SetAd2eev7 ADC trigger 2 on External Event 7
func (r *registerAdc2rType) SetAd2eev7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2eev7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2eev7Mask)
	}
}

const (
	RegisterAdc2rFieldAd2eev8Shift = 7
	RegisterAdc2rFieldAd2eev8Mask  = 0x80
)

// GetAd2eev8 ADC trigger 2 on External Event 8
func (r *registerAdc2rType) GetAd2eev8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2eev8Mask) != 0
}

// SetAd2eev8 ADC trigger 2 on External Event 8
func (r *registerAdc2rType) SetAd2eev8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2eev8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2eev8Mask)
	}
}

const (
	RegisterAdc2rFieldAd2eev9Shift = 8
	RegisterAdc2rFieldAd2eev9Mask  = 0x100
)

// GetAd2eev9 ADC trigger 2 on External Event 9
func (r *registerAdc2rType) GetAd2eev9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2eev9Mask) != 0
}

// SetAd2eev9 ADC trigger 2 on External Event 9
func (r *registerAdc2rType) SetAd2eev9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2eev9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2eev9Mask)
	}
}

const (
	RegisterAdc2rFieldAd2eev10Shift = 9
	RegisterAdc2rFieldAd2eev10Mask  = 0x200
)

// GetAd2eev10 ADC trigger 2 on External Event 10
func (r *registerAdc2rType) GetAd2eev10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2eev10Mask) != 0
}

// SetAd2eev10 ADC trigger 2 on External Event 10
func (r *registerAdc2rType) SetAd2eev10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2eev10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2eev10Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tac2Shift = 10
	RegisterAdc2rFieldAd2tac2Mask  = 0x400
)

// GetAd2tac2 ADC trigger 2 on Timer A compare 2
func (r *registerAdc2rType) GetAd2tac2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tac2Mask) != 0
}

// SetAd2tac2 ADC trigger 2 on Timer A compare 2
func (r *registerAdc2rType) SetAd2tac2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tac2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tac2Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tac3Shift = 11
	RegisterAdc2rFieldAd2tac3Mask  = 0x800
)

// GetAd2tac3 ADC trigger 2 on Timer A compare 3
func (r *registerAdc2rType) GetAd2tac3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tac3Mask) != 0
}

// SetAd2tac3 ADC trigger 2 on Timer A compare 3
func (r *registerAdc2rType) SetAd2tac3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tac3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tac3Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tac4Shift = 12
	RegisterAdc2rFieldAd2tac4Mask  = 0x1000
)

// GetAd2tac4 ADC trigger 2 on Timer A compare 4
func (r *registerAdc2rType) GetAd2tac4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tac4Mask) != 0
}

// SetAd2tac4 ADC trigger 2 on Timer A compare 4
func (r *registerAdc2rType) SetAd2tac4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tac4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tac4Mask)
	}
}

const (
	RegisterAdc2rFieldAd2taperShift = 13
	RegisterAdc2rFieldAd2taperMask  = 0x2000
)

// GetAd2taper ADC trigger 2 on Timer A Period
func (r *registerAdc2rType) GetAd2taper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2taperMask) != 0
}

// SetAd2taper ADC trigger 2 on Timer A Period
func (r *registerAdc2rType) SetAd2taper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2taperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2taperMask)
	}
}

const (
	RegisterAdc2rFieldAd2tbc2Shift = 14
	RegisterAdc2rFieldAd2tbc2Mask  = 0x4000
)

// GetAd2tbc2 ADC trigger 2 on Timer B compare 2
func (r *registerAdc2rType) GetAd2tbc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tbc2Mask) != 0
}

// SetAd2tbc2 ADC trigger 2 on Timer B compare 2
func (r *registerAdc2rType) SetAd2tbc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tbc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tbc2Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tbc3Shift = 15
	RegisterAdc2rFieldAd2tbc3Mask  = 0x8000
)

// GetAd2tbc3 ADC trigger 2 on Timer B compare 3
func (r *registerAdc2rType) GetAd2tbc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tbc3Mask) != 0
}

// SetAd2tbc3 ADC trigger 2 on Timer B compare 3
func (r *registerAdc2rType) SetAd2tbc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tbc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tbc3Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tbc4Shift = 16
	RegisterAdc2rFieldAd2tbc4Mask  = 0x10000
)

// GetAd2tbc4 ADC trigger 2 on Timer B compare 4
func (r *registerAdc2rType) GetAd2tbc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tbc4Mask) != 0
}

// SetAd2tbc4 ADC trigger 2 on Timer B compare 4
func (r *registerAdc2rType) SetAd2tbc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tbc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tbc4Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tbperShift = 17
	RegisterAdc2rFieldAd2tbperMask  = 0x20000
)

// GetAd2tbper ADC trigger 2 on Timer B Period
func (r *registerAdc2rType) GetAd2tbper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tbperMask) != 0
}

// SetAd2tbper ADC trigger 2 on Timer B Period
func (r *registerAdc2rType) SetAd2tbper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tbperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tbperMask)
	}
}

const (
	RegisterAdc2rFieldAd2tcc2Shift = 18
	RegisterAdc2rFieldAd2tcc2Mask  = 0x40000
)

// GetAd2tcc2 ADC trigger 2 on Timer C compare 2
func (r *registerAdc2rType) GetAd2tcc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tcc2Mask) != 0
}

// SetAd2tcc2 ADC trigger 2 on Timer C compare 2
func (r *registerAdc2rType) SetAd2tcc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tcc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tcc2Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tcc3Shift = 19
	RegisterAdc2rFieldAd2tcc3Mask  = 0x80000
)

// GetAd2tcc3 ADC trigger 2 on Timer C compare 3
func (r *registerAdc2rType) GetAd2tcc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tcc3Mask) != 0
}

// SetAd2tcc3 ADC trigger 2 on Timer C compare 3
func (r *registerAdc2rType) SetAd2tcc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tcc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tcc3Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tcc4Shift = 20
	RegisterAdc2rFieldAd2tcc4Mask  = 0x100000
)

// GetAd2tcc4 ADC trigger 2 on Timer C compare 4
func (r *registerAdc2rType) GetAd2tcc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tcc4Mask) != 0
}

// SetAd2tcc4 ADC trigger 2 on Timer C compare 4
func (r *registerAdc2rType) SetAd2tcc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tcc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tcc4Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tcperShift = 21
	RegisterAdc2rFieldAd2tcperMask  = 0x200000
)

// GetAd2tcper ADC trigger 2 on Timer C Period
func (r *registerAdc2rType) GetAd2tcper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tcperMask) != 0
}

// SetAd2tcper ADC trigger 2 on Timer C Period
func (r *registerAdc2rType) SetAd2tcper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tcperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tcperMask)
	}
}

const (
	RegisterAdc2rFieldAd2tcrstShift = 22
	RegisterAdc2rFieldAd2tcrstMask  = 0x400000
)

// GetAd2tcrst ADC trigger 2 on Timer C Reset
func (r *registerAdc2rType) GetAd2tcrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tcrstMask) != 0
}

// SetAd2tcrst ADC trigger 2 on Timer C Reset
func (r *registerAdc2rType) SetAd2tcrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tcrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tcrstMask)
	}
}

const (
	RegisterAdc2rFieldAd2tdc2Shift = 23
	RegisterAdc2rFieldAd2tdc2Mask  = 0x800000
)

// GetAd2tdc2 ADC trigger 2 on Timer D compare 2
func (r *registerAdc2rType) GetAd2tdc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tdc2Mask) != 0
}

// SetAd2tdc2 ADC trigger 2 on Timer D compare 2
func (r *registerAdc2rType) SetAd2tdc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tdc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tdc2Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tdc3Shift = 24
	RegisterAdc2rFieldAd2tdc3Mask  = 0x1000000
)

// GetAd2tdc3 ADC trigger 2 on Timer D compare 3
func (r *registerAdc2rType) GetAd2tdc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tdc3Mask) != 0
}

// SetAd2tdc3 ADC trigger 2 on Timer D compare 3
func (r *registerAdc2rType) SetAd2tdc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tdc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tdc3Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tdc4Shift = 25
	RegisterAdc2rFieldAd2tdc4Mask  = 0x2000000
)

// GetAd2tdc4 ADC trigger 2 on Timer D compare 4
func (r *registerAdc2rType) GetAd2tdc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tdc4Mask) != 0
}

// SetAd2tdc4 ADC trigger 2 on Timer D compare 4
func (r *registerAdc2rType) SetAd2tdc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tdc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tdc4Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tdperShift = 26
	RegisterAdc2rFieldAd2tdperMask  = 0x4000000
)

// GetAd2tdper ADC trigger 2 on Timer D Period
func (r *registerAdc2rType) GetAd2tdper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tdperMask) != 0
}

// SetAd2tdper ADC trigger 2 on Timer D Period
func (r *registerAdc2rType) SetAd2tdper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tdperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tdperMask)
	}
}

const (
	RegisterAdc2rFieldAd2tdrstShift = 27
	RegisterAdc2rFieldAd2tdrstMask  = 0x8000000
)

// GetAd2tdrst ADC trigger 2 on Timer D Reset
func (r *registerAdc2rType) GetAd2tdrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tdrstMask) != 0
}

// SetAd2tdrst ADC trigger 2 on Timer D Reset
func (r *registerAdc2rType) SetAd2tdrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tdrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tdrstMask)
	}
}

const (
	RegisterAdc2rFieldAd2tec2Shift = 28
	RegisterAdc2rFieldAd2tec2Mask  = 0x10000000
)

// GetAd2tec2 ADC trigger 2 on Timer E compare 2
func (r *registerAdc2rType) GetAd2tec2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tec2Mask) != 0
}

// SetAd2tec2 ADC trigger 2 on Timer E compare 2
func (r *registerAdc2rType) SetAd2tec2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tec2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tec2Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tec3Shift = 29
	RegisterAdc2rFieldAd2tec3Mask  = 0x20000000
)

// GetAd2tec3 ADC trigger 2 on Timer E compare 3
func (r *registerAdc2rType) GetAd2tec3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tec3Mask) != 0
}

// SetAd2tec3 ADC trigger 2 on Timer E compare 3
func (r *registerAdc2rType) SetAd2tec3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tec3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tec3Mask)
	}
}

const (
	RegisterAdc2rFieldAd2tec4Shift = 30
	RegisterAdc2rFieldAd2tec4Mask  = 0x40000000
)

// GetAd2tec4 ADC trigger 2 on Timer E compare 4
func (r *registerAdc2rType) GetAd2tec4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2tec4Mask) != 0
}

// SetAd2tec4 ADC trigger 2 on Timer E compare 4
func (r *registerAdc2rType) SetAd2tec4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2tec4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2tec4Mask)
	}
}

const (
	RegisterAdc2rFieldAd2terstShift = 31
	RegisterAdc2rFieldAd2terstMask  = 0x80000000
)

// GetAd2terst ADC trigger 2 on Timer E Reset
func (r *registerAdc2rType) GetAd2terst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc2rFieldAd2terstMask) != 0
}

// SetAd2terst ADC trigger 2 on Timer E Reset
func (r *registerAdc2rType) SetAd2terst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc2rFieldAd2terstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc2rFieldAd2terstMask)
	}
}

// registerAdc3rType ADC Trigger 3 Register
type registerAdc3rType uint32

const (
	RegisterAdc3rFieldAd1mc1Shift = 0
	RegisterAdc3rFieldAd1mc1Mask  = 0x1
)

// GetAd1mc1 AD1MC1
func (r *registerAdc3rType) GetAd1mc1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1mc1Mask) != 0
}

// SetAd1mc1 AD1MC1
func (r *registerAdc3rType) SetAd1mc1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1mc1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1mc1Mask)
	}
}

const (
	RegisterAdc3rFieldAd1mc2Shift = 1
	RegisterAdc3rFieldAd1mc2Mask  = 0x2
)

// GetAd1mc2 AD1MC2
func (r *registerAdc3rType) GetAd1mc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1mc2Mask) != 0
}

// SetAd1mc2 AD1MC2
func (r *registerAdc3rType) SetAd1mc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1mc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1mc2Mask)
	}
}

const (
	RegisterAdc3rFieldAd1mc3Shift = 2
	RegisterAdc3rFieldAd1mc3Mask  = 0x4
)

// GetAd1mc3 AD1MC3
func (r *registerAdc3rType) GetAd1mc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1mc3Mask) != 0
}

// SetAd1mc3 AD1MC3
func (r *registerAdc3rType) SetAd1mc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1mc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1mc3Mask)
	}
}

const (
	RegisterAdc3rFieldAd1mc4Shift = 3
	RegisterAdc3rFieldAd1mc4Mask  = 0x8
)

// GetAd1mc4 AD1MC4
func (r *registerAdc3rType) GetAd1mc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1mc4Mask) != 0
}

// SetAd1mc4 AD1MC4
func (r *registerAdc3rType) SetAd1mc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1mc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1mc4Mask)
	}
}

const (
	RegisterAdc3rFieldAd1mperShift = 4
	RegisterAdc3rFieldAd1mperMask  = 0x10
)

// GetAd1mper AD1MPER
func (r *registerAdc3rType) GetAd1mper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1mperMask) != 0
}

// SetAd1mper AD1MPER
func (r *registerAdc3rType) SetAd1mper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1mperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1mperMask)
	}
}

const (
	RegisterAdc3rFieldAd1eev1Shift = 5
	RegisterAdc3rFieldAd1eev1Mask  = 0x20
)

// GetAd1eev1 AD1EEV1
func (r *registerAdc3rType) GetAd1eev1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1eev1Mask) != 0
}

// SetAd1eev1 AD1EEV1
func (r *registerAdc3rType) SetAd1eev1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1eev1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1eev1Mask)
	}
}

const (
	RegisterAdc3rFieldAd1eev2Shift = 6
	RegisterAdc3rFieldAd1eev2Mask  = 0x40
)

// GetAd1eev2 AD1EEV2
func (r *registerAdc3rType) GetAd1eev2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1eev2Mask) != 0
}

// SetAd1eev2 AD1EEV2
func (r *registerAdc3rType) SetAd1eev2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1eev2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1eev2Mask)
	}
}

const (
	RegisterAdc3rFieldAd1eev3Shift = 7
	RegisterAdc3rFieldAd1eev3Mask  = 0x80
)

// GetAd1eev3 AD1EEV3
func (r *registerAdc3rType) GetAd1eev3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1eev3Mask) != 0
}

// SetAd1eev3 AD1EEV3
func (r *registerAdc3rType) SetAd1eev3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1eev3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1eev3Mask)
	}
}

const (
	RegisterAdc3rFieldAd1eev4Shift = 8
	RegisterAdc3rFieldAd1eev4Mask  = 0x100
)

// GetAd1eev4 AD1EEV4
func (r *registerAdc3rType) GetAd1eev4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1eev4Mask) != 0
}

// SetAd1eev4 AD1EEV4
func (r *registerAdc3rType) SetAd1eev4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1eev4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1eev4Mask)
	}
}

const (
	RegisterAdc3rFieldAd1eev5Shift = 9
	RegisterAdc3rFieldAd1eev5Mask  = 0x200
)

// GetAd1eev5 AD1EEV5
func (r *registerAdc3rType) GetAd1eev5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1eev5Mask) != 0
}

// SetAd1eev5 AD1EEV5
func (r *registerAdc3rType) SetAd1eev5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1eev5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1eev5Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tac2Shift = 10
	RegisterAdc3rFieldAd1tac2Mask  = 0x400
)

// GetAd1tac2 AD1TAC2
func (r *registerAdc3rType) GetAd1tac2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tac2Mask) != 0
}

// SetAd1tac2 AD1TAC2
func (r *registerAdc3rType) SetAd1tac2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tac2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tac2Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tac3Shift = 11
	RegisterAdc3rFieldAd1tac3Mask  = 0x800
)

// GetAd1tac3 AD1TAC3
func (r *registerAdc3rType) GetAd1tac3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tac3Mask) != 0
}

// SetAd1tac3 AD1TAC3
func (r *registerAdc3rType) SetAd1tac3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tac3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tac3Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tac4Shift = 12
	RegisterAdc3rFieldAd1tac4Mask  = 0x1000
)

// GetAd1tac4 AD1TAC4
func (r *registerAdc3rType) GetAd1tac4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tac4Mask) != 0
}

// SetAd1tac4 AD1TAC4
func (r *registerAdc3rType) SetAd1tac4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tac4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tac4Mask)
	}
}

const (
	RegisterAdc3rFieldAd1taperShift = 13
	RegisterAdc3rFieldAd1taperMask  = 0x2000
)

// GetAd1taper AD1TAPER
func (r *registerAdc3rType) GetAd1taper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1taperMask) != 0
}

// SetAd1taper AD1TAPER
func (r *registerAdc3rType) SetAd1taper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1taperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1taperMask)
	}
}

const (
	RegisterAdc3rFieldAd1tarstShift = 14
	RegisterAdc3rFieldAd1tarstMask  = 0x4000
)

// GetAd1tarst AD1TARST
func (r *registerAdc3rType) GetAd1tarst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tarstMask) != 0
}

// SetAd1tarst AD1TARST
func (r *registerAdc3rType) SetAd1tarst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tarstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tarstMask)
	}
}

const (
	RegisterAdc3rFieldAd1tbc2Shift = 15
	RegisterAdc3rFieldAd1tbc2Mask  = 0x8000
)

// GetAd1tbc2 AD1TBC2
func (r *registerAdc3rType) GetAd1tbc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tbc2Mask) != 0
}

// SetAd1tbc2 AD1TBC2
func (r *registerAdc3rType) SetAd1tbc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tbc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tbc2Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tbc3Shift = 16
	RegisterAdc3rFieldAd1tbc3Mask  = 0x10000
)

// GetAd1tbc3 AD1TBC3
func (r *registerAdc3rType) GetAd1tbc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tbc3Mask) != 0
}

// SetAd1tbc3 AD1TBC3
func (r *registerAdc3rType) SetAd1tbc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tbc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tbc3Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tbc4Shift = 17
	RegisterAdc3rFieldAd1tbc4Mask  = 0x20000
)

// GetAd1tbc4 AD1TBC4
func (r *registerAdc3rType) GetAd1tbc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tbc4Mask) != 0
}

// SetAd1tbc4 AD1TBC4
func (r *registerAdc3rType) SetAd1tbc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tbc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tbc4Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tbperShift = 18
	RegisterAdc3rFieldAd1tbperMask  = 0x40000
)

// GetAd1tbper AD1TBPER
func (r *registerAdc3rType) GetAd1tbper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tbperMask) != 0
}

// SetAd1tbper AD1TBPER
func (r *registerAdc3rType) SetAd1tbper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tbperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tbperMask)
	}
}

const (
	RegisterAdc3rFieldAd1tbrstShift = 19
	RegisterAdc3rFieldAd1tbrstMask  = 0x80000
)

// GetAd1tbrst AD1TBRST
func (r *registerAdc3rType) GetAd1tbrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tbrstMask) != 0
}

// SetAd1tbrst AD1TBRST
func (r *registerAdc3rType) SetAd1tbrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tbrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tbrstMask)
	}
}

const (
	RegisterAdc3rFieldAd1tcc2Shift = 20
	RegisterAdc3rFieldAd1tcc2Mask  = 0x100000
)

// GetAd1tcc2 AD1TCC2
func (r *registerAdc3rType) GetAd1tcc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tcc2Mask) != 0
}

// SetAd1tcc2 AD1TCC2
func (r *registerAdc3rType) SetAd1tcc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tcc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tcc2Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tcc3Shift = 21
	RegisterAdc3rFieldAd1tcc3Mask  = 0x200000
)

// GetAd1tcc3 AD1TCC3
func (r *registerAdc3rType) GetAd1tcc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tcc3Mask) != 0
}

// SetAd1tcc3 AD1TCC3
func (r *registerAdc3rType) SetAd1tcc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tcc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tcc3Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tcc4Shift = 22
	RegisterAdc3rFieldAd1tcc4Mask  = 0x400000
)

// GetAd1tcc4 AD1TCC4
func (r *registerAdc3rType) GetAd1tcc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tcc4Mask) != 0
}

// SetAd1tcc4 AD1TCC4
func (r *registerAdc3rType) SetAd1tcc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tcc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tcc4Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tcperShift = 23
	RegisterAdc3rFieldAd1tcperMask  = 0x800000
)

// GetAd1tcper AD1TCPER
func (r *registerAdc3rType) GetAd1tcper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tcperMask) != 0
}

// SetAd1tcper AD1TCPER
func (r *registerAdc3rType) SetAd1tcper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tcperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tcperMask)
	}
}

const (
	RegisterAdc3rFieldAd1tdc2Shift = 24
	RegisterAdc3rFieldAd1tdc2Mask  = 0x1000000
)

// GetAd1tdc2 AD1TDC2
func (r *registerAdc3rType) GetAd1tdc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tdc2Mask) != 0
}

// SetAd1tdc2 AD1TDC2
func (r *registerAdc3rType) SetAd1tdc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tdc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tdc2Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tdc3Shift = 25
	RegisterAdc3rFieldAd1tdc3Mask  = 0x2000000
)

// GetAd1tdc3 AD1TDC3
func (r *registerAdc3rType) GetAd1tdc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tdc3Mask) != 0
}

// SetAd1tdc3 AD1TDC3
func (r *registerAdc3rType) SetAd1tdc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tdc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tdc3Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tdc4Shift = 26
	RegisterAdc3rFieldAd1tdc4Mask  = 0x4000000
)

// GetAd1tdc4 AD1TDC4
func (r *registerAdc3rType) GetAd1tdc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tdc4Mask) != 0
}

// SetAd1tdc4 AD1TDC4
func (r *registerAdc3rType) SetAd1tdc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tdc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tdc4Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tdperShift = 27
	RegisterAdc3rFieldAd1tdperMask  = 0x8000000
)

// GetAd1tdper AD1TDPER
func (r *registerAdc3rType) GetAd1tdper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tdperMask) != 0
}

// SetAd1tdper AD1TDPER
func (r *registerAdc3rType) SetAd1tdper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tdperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tdperMask)
	}
}

const (
	RegisterAdc3rFieldAd1tec2Shift = 28
	RegisterAdc3rFieldAd1tec2Mask  = 0x10000000
)

// GetAd1tec2 AD1TEC2
func (r *registerAdc3rType) GetAd1tec2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tec2Mask) != 0
}

// SetAd1tec2 AD1TEC2
func (r *registerAdc3rType) SetAd1tec2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tec2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tec2Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tec3Shift = 29
	RegisterAdc3rFieldAd1tec3Mask  = 0x20000000
)

// GetAd1tec3 AD1TEC3
func (r *registerAdc3rType) GetAd1tec3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tec3Mask) != 0
}

// SetAd1tec3 AD1TEC3
func (r *registerAdc3rType) SetAd1tec3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tec3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tec3Mask)
	}
}

const (
	RegisterAdc3rFieldAd1tec4Shift = 30
	RegisterAdc3rFieldAd1tec4Mask  = 0x40000000
)

// GetAd1tec4 AD1TEC4
func (r *registerAdc3rType) GetAd1tec4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1tec4Mask) != 0
}

// SetAd1tec4 AD1TEC4
func (r *registerAdc3rType) SetAd1tec4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1tec4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1tec4Mask)
	}
}

const (
	RegisterAdc3rFieldAd1teperShift = 31
	RegisterAdc3rFieldAd1teperMask  = 0x80000000
)

// GetAd1teper AD1TEPER
func (r *registerAdc3rType) GetAd1teper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc3rFieldAd1teperMask) != 0
}

// SetAd1teper AD1TEPER
func (r *registerAdc3rType) SetAd1teper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc3rFieldAd1teperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc3rFieldAd1teperMask)
	}
}

// registerAdc4rType ADC Trigger 4 Register
type registerAdc4rType uint32

const (
	RegisterAdc4rFieldAd2mc1Shift = 0
	RegisterAdc4rFieldAd2mc1Mask  = 0x1
)

// GetAd2mc1 AD2MC1
func (r *registerAdc4rType) GetAd2mc1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2mc1Mask) != 0
}

// SetAd2mc1 AD2MC1
func (r *registerAdc4rType) SetAd2mc1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2mc1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2mc1Mask)
	}
}

const (
	RegisterAdc4rFieldAd2mc2Shift = 1
	RegisterAdc4rFieldAd2mc2Mask  = 0x2
)

// GetAd2mc2 AD2MC2
func (r *registerAdc4rType) GetAd2mc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2mc2Mask) != 0
}

// SetAd2mc2 AD2MC2
func (r *registerAdc4rType) SetAd2mc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2mc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2mc2Mask)
	}
}

const (
	RegisterAdc4rFieldAd2mc3Shift = 2
	RegisterAdc4rFieldAd2mc3Mask  = 0x4
)

// GetAd2mc3 AD2MC3
func (r *registerAdc4rType) GetAd2mc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2mc3Mask) != 0
}

// SetAd2mc3 AD2MC3
func (r *registerAdc4rType) SetAd2mc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2mc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2mc3Mask)
	}
}

const (
	RegisterAdc4rFieldAd2mc4Shift = 3
	RegisterAdc4rFieldAd2mc4Mask  = 0x8
)

// GetAd2mc4 AD2MC4
func (r *registerAdc4rType) GetAd2mc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2mc4Mask) != 0
}

// SetAd2mc4 AD2MC4
func (r *registerAdc4rType) SetAd2mc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2mc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2mc4Mask)
	}
}

const (
	RegisterAdc4rFieldAd2mperShift = 4
	RegisterAdc4rFieldAd2mperMask  = 0x10
)

// GetAd2mper AD2MPER
func (r *registerAdc4rType) GetAd2mper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2mperMask) != 0
}

// SetAd2mper AD2MPER
func (r *registerAdc4rType) SetAd2mper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2mperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2mperMask)
	}
}

const (
	RegisterAdc4rFieldAd2eev6Shift = 5
	RegisterAdc4rFieldAd2eev6Mask  = 0x20
)

// GetAd2eev6 AD2EEV6
func (r *registerAdc4rType) GetAd2eev6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2eev6Mask) != 0
}

// SetAd2eev6 AD2EEV6
func (r *registerAdc4rType) SetAd2eev6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2eev6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2eev6Mask)
	}
}

const (
	RegisterAdc4rFieldAd2eev7Shift = 6
	RegisterAdc4rFieldAd2eev7Mask  = 0x40
)

// GetAd2eev7 AD2EEV7
func (r *registerAdc4rType) GetAd2eev7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2eev7Mask) != 0
}

// SetAd2eev7 AD2EEV7
func (r *registerAdc4rType) SetAd2eev7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2eev7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2eev7Mask)
	}
}

const (
	RegisterAdc4rFieldAd2eev8Shift = 7
	RegisterAdc4rFieldAd2eev8Mask  = 0x80
)

// GetAd2eev8 AD2EEV8
func (r *registerAdc4rType) GetAd2eev8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2eev8Mask) != 0
}

// SetAd2eev8 AD2EEV8
func (r *registerAdc4rType) SetAd2eev8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2eev8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2eev8Mask)
	}
}

const (
	RegisterAdc4rFieldAd2eev9Shift = 8
	RegisterAdc4rFieldAd2eev9Mask  = 0x100
)

// GetAd2eev9 AD2EEV9
func (r *registerAdc4rType) GetAd2eev9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2eev9Mask) != 0
}

// SetAd2eev9 AD2EEV9
func (r *registerAdc4rType) SetAd2eev9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2eev9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2eev9Mask)
	}
}

const (
	RegisterAdc4rFieldAd2eev10Shift = 9
	RegisterAdc4rFieldAd2eev10Mask  = 0x200
)

// GetAd2eev10 AD2EEV10
func (r *registerAdc4rType) GetAd2eev10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2eev10Mask) != 0
}

// SetAd2eev10 AD2EEV10
func (r *registerAdc4rType) SetAd2eev10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2eev10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2eev10Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tac2Shift = 10
	RegisterAdc4rFieldAd2tac2Mask  = 0x400
)

// GetAd2tac2 AD2TAC2
func (r *registerAdc4rType) GetAd2tac2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tac2Mask) != 0
}

// SetAd2tac2 AD2TAC2
func (r *registerAdc4rType) SetAd2tac2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tac2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tac2Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tac3Shift = 11
	RegisterAdc4rFieldAd2tac3Mask  = 0x800
)

// GetAd2tac3 AD2TAC3
func (r *registerAdc4rType) GetAd2tac3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tac3Mask) != 0
}

// SetAd2tac3 AD2TAC3
func (r *registerAdc4rType) SetAd2tac3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tac3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tac3Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tac4Shift = 12
	RegisterAdc4rFieldAd2tac4Mask  = 0x1000
)

// GetAd2tac4 AD2TAC4
func (r *registerAdc4rType) GetAd2tac4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tac4Mask) != 0
}

// SetAd2tac4 AD2TAC4
func (r *registerAdc4rType) SetAd2tac4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tac4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tac4Mask)
	}
}

const (
	RegisterAdc4rFieldAd2taperShift = 13
	RegisterAdc4rFieldAd2taperMask  = 0x2000
)

// GetAd2taper AD2TAPER
func (r *registerAdc4rType) GetAd2taper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2taperMask) != 0
}

// SetAd2taper AD2TAPER
func (r *registerAdc4rType) SetAd2taper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2taperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2taperMask)
	}
}

const (
	RegisterAdc4rFieldAd2tbc2Shift = 14
	RegisterAdc4rFieldAd2tbc2Mask  = 0x4000
)

// GetAd2tbc2 AD2TBC2
func (r *registerAdc4rType) GetAd2tbc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tbc2Mask) != 0
}

// SetAd2tbc2 AD2TBC2
func (r *registerAdc4rType) SetAd2tbc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tbc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tbc2Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tbc3Shift = 15
	RegisterAdc4rFieldAd2tbc3Mask  = 0x8000
)

// GetAd2tbc3 AD2TBC3
func (r *registerAdc4rType) GetAd2tbc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tbc3Mask) != 0
}

// SetAd2tbc3 AD2TBC3
func (r *registerAdc4rType) SetAd2tbc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tbc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tbc3Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tbc4Shift = 16
	RegisterAdc4rFieldAd2tbc4Mask  = 0x10000
)

// GetAd2tbc4 AD2TBC4
func (r *registerAdc4rType) GetAd2tbc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tbc4Mask) != 0
}

// SetAd2tbc4 AD2TBC4
func (r *registerAdc4rType) SetAd2tbc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tbc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tbc4Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tbperShift = 17
	RegisterAdc4rFieldAd2tbperMask  = 0x20000
)

// GetAd2tbper AD2TBPER
func (r *registerAdc4rType) GetAd2tbper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tbperMask) != 0
}

// SetAd2tbper AD2TBPER
func (r *registerAdc4rType) SetAd2tbper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tbperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tbperMask)
	}
}

const (
	RegisterAdc4rFieldAd2tcc2Shift = 18
	RegisterAdc4rFieldAd2tcc2Mask  = 0x40000
)

// GetAd2tcc2 AD2TCC2
func (r *registerAdc4rType) GetAd2tcc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tcc2Mask) != 0
}

// SetAd2tcc2 AD2TCC2
func (r *registerAdc4rType) SetAd2tcc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tcc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tcc2Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tcc3Shift = 19
	RegisterAdc4rFieldAd2tcc3Mask  = 0x80000
)

// GetAd2tcc3 AD2TCC3
func (r *registerAdc4rType) GetAd2tcc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tcc3Mask) != 0
}

// SetAd2tcc3 AD2TCC3
func (r *registerAdc4rType) SetAd2tcc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tcc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tcc3Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tcc4Shift = 20
	RegisterAdc4rFieldAd2tcc4Mask  = 0x100000
)

// GetAd2tcc4 AD2TCC4
func (r *registerAdc4rType) GetAd2tcc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tcc4Mask) != 0
}

// SetAd2tcc4 AD2TCC4
func (r *registerAdc4rType) SetAd2tcc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tcc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tcc4Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tcperShift = 21
	RegisterAdc4rFieldAd2tcperMask  = 0x200000
)

// GetAd2tcper AD2TCPER
func (r *registerAdc4rType) GetAd2tcper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tcperMask) != 0
}

// SetAd2tcper AD2TCPER
func (r *registerAdc4rType) SetAd2tcper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tcperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tcperMask)
	}
}

const (
	RegisterAdc4rFieldAd2tcrstShift = 22
	RegisterAdc4rFieldAd2tcrstMask  = 0x400000
)

// GetAd2tcrst AD2TCRST
func (r *registerAdc4rType) GetAd2tcrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tcrstMask) != 0
}

// SetAd2tcrst AD2TCRST
func (r *registerAdc4rType) SetAd2tcrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tcrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tcrstMask)
	}
}

const (
	RegisterAdc4rFieldAd2tdc2Shift = 23
	RegisterAdc4rFieldAd2tdc2Mask  = 0x800000
)

// GetAd2tdc2 AD2TDC2
func (r *registerAdc4rType) GetAd2tdc2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tdc2Mask) != 0
}

// SetAd2tdc2 AD2TDC2
func (r *registerAdc4rType) SetAd2tdc2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tdc2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tdc2Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tdc3Shift = 24
	RegisterAdc4rFieldAd2tdc3Mask  = 0x1000000
)

// GetAd2tdc3 AD2TDC3
func (r *registerAdc4rType) GetAd2tdc3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tdc3Mask) != 0
}

// SetAd2tdc3 AD2TDC3
func (r *registerAdc4rType) SetAd2tdc3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tdc3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tdc3Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tdc4Shift = 25
	RegisterAdc4rFieldAd2tdc4Mask  = 0x2000000
)

// GetAd2tdc4 AD2TDC4
func (r *registerAdc4rType) GetAd2tdc4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tdc4Mask) != 0
}

// SetAd2tdc4 AD2TDC4
func (r *registerAdc4rType) SetAd2tdc4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tdc4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tdc4Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tdperShift = 26
	RegisterAdc4rFieldAd2tdperMask  = 0x4000000
)

// GetAd2tdper AD2TDPER
func (r *registerAdc4rType) GetAd2tdper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tdperMask) != 0
}

// SetAd2tdper AD2TDPER
func (r *registerAdc4rType) SetAd2tdper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tdperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tdperMask)
	}
}

const (
	RegisterAdc4rFieldAd2tdrstShift = 27
	RegisterAdc4rFieldAd2tdrstMask  = 0x8000000
)

// GetAd2tdrst AD2TDRST
func (r *registerAdc4rType) GetAd2tdrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tdrstMask) != 0
}

// SetAd2tdrst AD2TDRST
func (r *registerAdc4rType) SetAd2tdrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tdrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tdrstMask)
	}
}

const (
	RegisterAdc4rFieldAd2tec2Shift = 28
	RegisterAdc4rFieldAd2tec2Mask  = 0x10000000
)

// GetAd2tec2 AD2TEC2
func (r *registerAdc4rType) GetAd2tec2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tec2Mask) != 0
}

// SetAd2tec2 AD2TEC2
func (r *registerAdc4rType) SetAd2tec2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tec2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tec2Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tec3Shift = 29
	RegisterAdc4rFieldAd2tec3Mask  = 0x20000000
)

// GetAd2tec3 AD2TEC3
func (r *registerAdc4rType) GetAd2tec3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tec3Mask) != 0
}

// SetAd2tec3 AD2TEC3
func (r *registerAdc4rType) SetAd2tec3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tec3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tec3Mask)
	}
}

const (
	RegisterAdc4rFieldAd2tec4Shift = 30
	RegisterAdc4rFieldAd2tec4Mask  = 0x40000000
)

// GetAd2tec4 AD2TEC4
func (r *registerAdc4rType) GetAd2tec4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2tec4Mask) != 0
}

// SetAd2tec4 AD2TEC4
func (r *registerAdc4rType) SetAd2tec4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2tec4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2tec4Mask)
	}
}

const (
	RegisterAdc4rFieldAd2terstShift = 31
	RegisterAdc4rFieldAd2terstMask  = 0x80000000
)

// GetAd2terst AD2TERST
func (r *registerAdc4rType) GetAd2terst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAdc4rFieldAd2terstMask) != 0
}

// SetAd2terst AD2TERST
func (r *registerAdc4rType) SetAd2terst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAdc4rFieldAd2terstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAdc4rFieldAd2terstMask)
	}
}

// registerDllcrType DLL Control Register
type registerDllcrType uint32

const (
	RegisterDllcrFieldCalShift = 0
	RegisterDllcrFieldCalMask  = 0x1
)

// GetCal DLL Calibration Start
func (r *registerDllcrType) GetCal() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDllcrFieldCalMask) != 0
}

// SetCal DLL Calibration Start
func (r *registerDllcrType) SetCal(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDllcrFieldCalMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDllcrFieldCalMask)
	}
}

const (
	RegisterDllcrFieldCalenShift = 1
	RegisterDllcrFieldCalenMask  = 0x2
)

// GetCalen DLL Calibration Enable
func (r *registerDllcrType) GetCalen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDllcrFieldCalenMask) != 0
}

// SetCalen DLL Calibration Enable
func (r *registerDllcrType) SetCalen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDllcrFieldCalenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDllcrFieldCalenMask)
	}
}

const (
	RegisterDllcrFieldCalrteShift = 2
	RegisterDllcrFieldCalrteMask  = 0xc
)

// GetCalrte DLL Calibration rate
func (r *registerDllcrType) GetCalrte() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDllcrFieldCalrteMask) >> RegisterDllcrFieldCalrteShift)
}

// SetCalrte DLL Calibration rate
func (r *registerDllcrType) SetCalrte(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDllcrFieldCalrteMask)|(uint32(value)<<RegisterDllcrFieldCalrteShift))
}

// registerFltinr1Type HRTIM Fault Input Register 1
type registerFltinr1Type uint32

const (
	RegisterFltinr1FieldFlt1eShift = 0
	RegisterFltinr1FieldFlt1eMask  = 0x1
)

// GetFlt1e FLT1E
func (r *registerFltinr1Type) GetFlt1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt1eMask) != 0
}

// SetFlt1e FLT1E
func (r *registerFltinr1Type) SetFlt1e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt1eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt1eMask)
	}
}

const (
	RegisterFltinr1FieldFlt1pShift = 1
	RegisterFltinr1FieldFlt1pMask  = 0x2
)

// GetFlt1p FLT1P
func (r *registerFltinr1Type) GetFlt1p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt1pMask) != 0
}

// SetFlt1p FLT1P
func (r *registerFltinr1Type) SetFlt1p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt1pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt1pMask)
	}
}

const (
	RegisterFltinr1FieldFlt1srcShift = 2
	RegisterFltinr1FieldFlt1srcMask  = 0x4
)

// GetFlt1src FLT1SRC
func (r *registerFltinr1Type) GetFlt1src() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt1srcMask) != 0
}

// SetFlt1src FLT1SRC
func (r *registerFltinr1Type) SetFlt1src(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt1srcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt1srcMask)
	}
}

const (
	RegisterFltinr1FieldFlt1fShift = 3
	RegisterFltinr1FieldFlt1fMask  = 0x78
)

// GetFlt1f FLT1F
func (r *registerFltinr1Type) GetFlt1f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt1fMask) >> RegisterFltinr1FieldFlt1fShift)
}

// SetFlt1f FLT1F
func (r *registerFltinr1Type) SetFlt1f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt1fMask)|(uint32(value)<<RegisterFltinr1FieldFlt1fShift))
}

const (
	RegisterFltinr1FieldFlt1lckShift = 7
	RegisterFltinr1FieldFlt1lckMask  = 0x80
)

// GetFlt1lck FLT1LCK
func (r *registerFltinr1Type) GetFlt1lck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt1lckMask) != 0
}

// SetFlt1lck FLT1LCK
func (r *registerFltinr1Type) SetFlt1lck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt1lckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt1lckMask)
	}
}

const (
	RegisterFltinr1FieldFlt2eShift = 8
	RegisterFltinr1FieldFlt2eMask  = 0x100
)

// GetFlt2e FLT2E
func (r *registerFltinr1Type) GetFlt2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt2eMask) != 0
}

// SetFlt2e FLT2E
func (r *registerFltinr1Type) SetFlt2e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt2eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt2eMask)
	}
}

const (
	RegisterFltinr1FieldFlt2pShift = 9
	RegisterFltinr1FieldFlt2pMask  = 0x200
)

// GetFlt2p FLT2P
func (r *registerFltinr1Type) GetFlt2p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt2pMask) != 0
}

// SetFlt2p FLT2P
func (r *registerFltinr1Type) SetFlt2p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt2pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt2pMask)
	}
}

const (
	RegisterFltinr1FieldFlt2srcShift = 10
	RegisterFltinr1FieldFlt2srcMask  = 0x400
)

// GetFlt2src FLT2SRC
func (r *registerFltinr1Type) GetFlt2src() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt2srcMask) != 0
}

// SetFlt2src FLT2SRC
func (r *registerFltinr1Type) SetFlt2src(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt2srcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt2srcMask)
	}
}

const (
	RegisterFltinr1FieldFlt2fShift = 11
	RegisterFltinr1FieldFlt2fMask  = 0x7800
)

// GetFlt2f FLT2F
func (r *registerFltinr1Type) GetFlt2f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt2fMask) >> RegisterFltinr1FieldFlt2fShift)
}

// SetFlt2f FLT2F
func (r *registerFltinr1Type) SetFlt2f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt2fMask)|(uint32(value)<<RegisterFltinr1FieldFlt2fShift))
}

const (
	RegisterFltinr1FieldFlt2lckShift = 15
	RegisterFltinr1FieldFlt2lckMask  = 0x8000
)

// GetFlt2lck FLT2LCK
func (r *registerFltinr1Type) GetFlt2lck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt2lckMask) != 0
}

// SetFlt2lck FLT2LCK
func (r *registerFltinr1Type) SetFlt2lck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt2lckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt2lckMask)
	}
}

const (
	RegisterFltinr1FieldFlt3eShift = 16
	RegisterFltinr1FieldFlt3eMask  = 0x10000
)

// GetFlt3e FLT3E
func (r *registerFltinr1Type) GetFlt3e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt3eMask) != 0
}

// SetFlt3e FLT3E
func (r *registerFltinr1Type) SetFlt3e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt3eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt3eMask)
	}
}

const (
	RegisterFltinr1FieldFlt3pShift = 17
	RegisterFltinr1FieldFlt3pMask  = 0x20000
)

// GetFlt3p FLT3P
func (r *registerFltinr1Type) GetFlt3p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt3pMask) != 0
}

// SetFlt3p FLT3P
func (r *registerFltinr1Type) SetFlt3p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt3pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt3pMask)
	}
}

const (
	RegisterFltinr1FieldFlt3srcShift = 18
	RegisterFltinr1FieldFlt3srcMask  = 0x40000
)

// GetFlt3src FLT3SRC
func (r *registerFltinr1Type) GetFlt3src() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt3srcMask) != 0
}

// SetFlt3src FLT3SRC
func (r *registerFltinr1Type) SetFlt3src(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt3srcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt3srcMask)
	}
}

const (
	RegisterFltinr1FieldFlt3fShift = 19
	RegisterFltinr1FieldFlt3fMask  = 0x780000
)

// GetFlt3f FLT3F
func (r *registerFltinr1Type) GetFlt3f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt3fMask) >> RegisterFltinr1FieldFlt3fShift)
}

// SetFlt3f FLT3F
func (r *registerFltinr1Type) SetFlt3f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt3fMask)|(uint32(value)<<RegisterFltinr1FieldFlt3fShift))
}

const (
	RegisterFltinr1FieldFlt3lckShift = 23
	RegisterFltinr1FieldFlt3lckMask  = 0x800000
)

// GetFlt3lck FLT3LCK
func (r *registerFltinr1Type) GetFlt3lck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt3lckMask) != 0
}

// SetFlt3lck FLT3LCK
func (r *registerFltinr1Type) SetFlt3lck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt3lckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt3lckMask)
	}
}

const (
	RegisterFltinr1FieldFlt4eShift = 24
	RegisterFltinr1FieldFlt4eMask  = 0x1000000
)

// GetFlt4e FLT4E
func (r *registerFltinr1Type) GetFlt4e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt4eMask) != 0
}

// SetFlt4e FLT4E
func (r *registerFltinr1Type) SetFlt4e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt4eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt4eMask)
	}
}

const (
	RegisterFltinr1FieldFlt4pShift = 25
	RegisterFltinr1FieldFlt4pMask  = 0x2000000
)

// GetFlt4p FLT4P
func (r *registerFltinr1Type) GetFlt4p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt4pMask) != 0
}

// SetFlt4p FLT4P
func (r *registerFltinr1Type) SetFlt4p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt4pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt4pMask)
	}
}

const (
	RegisterFltinr1FieldFlt4srcShift = 26
	RegisterFltinr1FieldFlt4srcMask  = 0x4000000
)

// GetFlt4src FLT4SRC
func (r *registerFltinr1Type) GetFlt4src() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt4srcMask) != 0
}

// SetFlt4src FLT4SRC
func (r *registerFltinr1Type) SetFlt4src(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt4srcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt4srcMask)
	}
}

const (
	RegisterFltinr1FieldFlt4fShift = 27
	RegisterFltinr1FieldFlt4fMask  = 0x78000000
)

// GetFlt4f FLT4F
func (r *registerFltinr1Type) GetFlt4f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt4fMask) >> RegisterFltinr1FieldFlt4fShift)
}

// SetFlt4f FLT4F
func (r *registerFltinr1Type) SetFlt4f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt4fMask)|(uint32(value)<<RegisterFltinr1FieldFlt4fShift))
}

const (
	RegisterFltinr1FieldFlt4lckShift = 31
	RegisterFltinr1FieldFlt4lckMask  = 0x80000000
)

// GetFlt4lck FLT4LCK
func (r *registerFltinr1Type) GetFlt4lck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr1FieldFlt4lckMask) != 0
}

// SetFlt4lck FLT4LCK
func (r *registerFltinr1Type) SetFlt4lck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr1FieldFlt4lckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr1FieldFlt4lckMask)
	}
}

// registerFltinr2Type HRTIM Fault Input Register 2
type registerFltinr2Type uint32

const (
	RegisterFltinr2FieldFlt5eShift = 0
	RegisterFltinr2FieldFlt5eMask  = 0x1
)

// GetFlt5e FLT5E
func (r *registerFltinr2Type) GetFlt5e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr2FieldFlt5eMask) != 0
}

// SetFlt5e FLT5E
func (r *registerFltinr2Type) SetFlt5e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr2FieldFlt5eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr2FieldFlt5eMask)
	}
}

const (
	RegisterFltinr2FieldFlt5pShift = 1
	RegisterFltinr2FieldFlt5pMask  = 0x2
)

// GetFlt5p FLT5P
func (r *registerFltinr2Type) GetFlt5p() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr2FieldFlt5pMask) != 0
}

// SetFlt5p FLT5P
func (r *registerFltinr2Type) SetFlt5p(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr2FieldFlt5pMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr2FieldFlt5pMask)
	}
}

const (
	RegisterFltinr2FieldFlt5srcShift = 2
	RegisterFltinr2FieldFlt5srcMask  = 0x4
)

// GetFlt5src FLT5SRC
func (r *registerFltinr2Type) GetFlt5src() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr2FieldFlt5srcMask) != 0
}

// SetFlt5src FLT5SRC
func (r *registerFltinr2Type) SetFlt5src(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr2FieldFlt5srcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr2FieldFlt5srcMask)
	}
}

const (
	RegisterFltinr2FieldFlt5fShift = 3
	RegisterFltinr2FieldFlt5fMask  = 0x78
)

// GetFlt5f FLT5F
func (r *registerFltinr2Type) GetFlt5f() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFltinr2FieldFlt5fMask) >> RegisterFltinr2FieldFlt5fShift)
}

// SetFlt5f FLT5F
func (r *registerFltinr2Type) SetFlt5f(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFltinr2FieldFlt5fMask)|(uint32(value)<<RegisterFltinr2FieldFlt5fShift))
}

const (
	RegisterFltinr2FieldFlt5lckShift = 7
	RegisterFltinr2FieldFlt5lckMask  = 0x80
)

// GetFlt5lck FLT5LCK
func (r *registerFltinr2Type) GetFlt5lck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltinr2FieldFlt5lckMask) != 0
}

// SetFlt5lck FLT5LCK
func (r *registerFltinr2Type) SetFlt5lck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltinr2FieldFlt5lckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltinr2FieldFlt5lckMask)
	}
}

const (
	RegisterFltinr2FieldFltsdShift = 24
	RegisterFltinr2FieldFltsdMask  = 0x3000000
)

// GetFltsd FLTSD
func (r *registerFltinr2Type) GetFltsd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFltinr2FieldFltsdMask) >> RegisterFltinr2FieldFltsdShift)
}

// SetFltsd FLTSD
func (r *registerFltinr2Type) SetFltsd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFltinr2FieldFltsdMask)|(uint32(value)<<RegisterFltinr2FieldFltsdShift))
}

// registerBdmupdrType BDMUPDR
type registerBdmupdrType uint32

const (
	RegisterBdmupdrFieldMcrShift = 0
	RegisterBdmupdrFieldMcrMask  = 0x1
)

// GetMcr MCR
func (r *registerBdmupdrType) GetMcr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdmupdrFieldMcrMask) != 0
}

// SetMcr MCR
func (r *registerBdmupdrType) SetMcr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdmupdrFieldMcrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdmupdrFieldMcrMask)
	}
}

const (
	RegisterBdmupdrFieldMicrShift = 1
	RegisterBdmupdrFieldMicrMask  = 0x2
)

// GetMicr MICR
func (r *registerBdmupdrType) GetMicr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdmupdrFieldMicrMask) != 0
}

// SetMicr MICR
func (r *registerBdmupdrType) SetMicr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdmupdrFieldMicrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdmupdrFieldMicrMask)
	}
}

const (
	RegisterBdmupdrFieldMdierShift = 2
	RegisterBdmupdrFieldMdierMask  = 0x4
)

// GetMdier MDIER
func (r *registerBdmupdrType) GetMdier() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdmupdrFieldMdierMask) != 0
}

// SetMdier MDIER
func (r *registerBdmupdrType) SetMdier(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdmupdrFieldMdierMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdmupdrFieldMdierMask)
	}
}

const (
	RegisterBdmupdrFieldMcntShift = 3
	RegisterBdmupdrFieldMcntMask  = 0x8
)

// GetMcnt MCNT
func (r *registerBdmupdrType) GetMcnt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdmupdrFieldMcntMask) != 0
}

// SetMcnt MCNT
func (r *registerBdmupdrType) SetMcnt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdmupdrFieldMcntMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdmupdrFieldMcntMask)
	}
}

const (
	RegisterBdmupdrFieldMperShift = 4
	RegisterBdmupdrFieldMperMask  = 0x10
)

// GetMper MPER
func (r *registerBdmupdrType) GetMper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdmupdrFieldMperMask) != 0
}

// SetMper MPER
func (r *registerBdmupdrType) SetMper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdmupdrFieldMperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdmupdrFieldMperMask)
	}
}

const (
	RegisterBdmupdrFieldMrepShift = 5
	RegisterBdmupdrFieldMrepMask  = 0x20
)

// GetMrep MREP
func (r *registerBdmupdrType) GetMrep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdmupdrFieldMrepMask) != 0
}

// SetMrep MREP
func (r *registerBdmupdrType) SetMrep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdmupdrFieldMrepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdmupdrFieldMrepMask)
	}
}

const (
	RegisterBdmupdrFieldMcmp1Shift = 6
	RegisterBdmupdrFieldMcmp1Mask  = 0x40
)

// GetMcmp1 MCMP1
func (r *registerBdmupdrType) GetMcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdmupdrFieldMcmp1Mask) != 0
}

// SetMcmp1 MCMP1
func (r *registerBdmupdrType) SetMcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdmupdrFieldMcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdmupdrFieldMcmp1Mask)
	}
}

const (
	RegisterBdmupdrFieldMcmp2Shift = 7
	RegisterBdmupdrFieldMcmp2Mask  = 0x80
)

// GetMcmp2 MCMP2
func (r *registerBdmupdrType) GetMcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdmupdrFieldMcmp2Mask) != 0
}

// SetMcmp2 MCMP2
func (r *registerBdmupdrType) SetMcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdmupdrFieldMcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdmupdrFieldMcmp2Mask)
	}
}

const (
	RegisterBdmupdrFieldMcmp3Shift = 8
	RegisterBdmupdrFieldMcmp3Mask  = 0x100
)

// GetMcmp3 MCMP3
func (r *registerBdmupdrType) GetMcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdmupdrFieldMcmp3Mask) != 0
}

// SetMcmp3 MCMP3
func (r *registerBdmupdrType) SetMcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdmupdrFieldMcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdmupdrFieldMcmp3Mask)
	}
}

const (
	RegisterBdmupdrFieldMcmp4Shift = 9
	RegisterBdmupdrFieldMcmp4Mask  = 0x200
)

// GetMcmp4 MCMP4
func (r *registerBdmupdrType) GetMcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdmupdrFieldMcmp4Mask) != 0
}

// SetMcmp4 MCMP4
func (r *registerBdmupdrType) SetMcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdmupdrFieldMcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdmupdrFieldMcmp4Mask)
	}
}

// registerBdtxuprType Burst DMA Timerx update Register
type registerBdtxuprType uint32

const (
	RegisterBdtxuprFieldTimxcrShift = 0
	RegisterBdtxuprFieldTimxcrMask  = 0x1
)

// GetTimxcr HRTIM_TIMxCR register update enable
func (r *registerBdtxuprType) GetTimxcr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxcrMask) != 0
}

// SetTimxcr HRTIM_TIMxCR register update enable
func (r *registerBdtxuprType) SetTimxcr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxcrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxcrMask)
	}
}

const (
	RegisterBdtxuprFieldTimxicrShift = 1
	RegisterBdtxuprFieldTimxicrMask  = 0x2
)

// GetTimxicr HRTIM_TIMxICR register update enable
func (r *registerBdtxuprType) GetTimxicr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxicrMask) != 0
}

// SetTimxicr HRTIM_TIMxICR register update enable
func (r *registerBdtxuprType) SetTimxicr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxicrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxicrMask)
	}
}

const (
	RegisterBdtxuprFieldTimxdierShift = 2
	RegisterBdtxuprFieldTimxdierMask  = 0x4
)

// GetTimxdier HRTIM_TIMxDIER register update enable
func (r *registerBdtxuprType) GetTimxdier() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxdierMask) != 0
}

// SetTimxdier HRTIM_TIMxDIER register update enable
func (r *registerBdtxuprType) SetTimxdier(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxdierMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxdierMask)
	}
}

const (
	RegisterBdtxuprFieldTimxcntShift = 3
	RegisterBdtxuprFieldTimxcntMask  = 0x8
)

// GetTimxcnt HRTIM_CNTxR register update enable
func (r *registerBdtxuprType) GetTimxcnt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxcntMask) != 0
}

// SetTimxcnt HRTIM_CNTxR register update enable
func (r *registerBdtxuprType) SetTimxcnt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxcntMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxcntMask)
	}
}

const (
	RegisterBdtxuprFieldTimxperShift = 4
	RegisterBdtxuprFieldTimxperMask  = 0x10
)

// GetTimxper HRTIM_PERxR register update enable
func (r *registerBdtxuprType) GetTimxper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxperMask) != 0
}

// SetTimxper HRTIM_PERxR register update enable
func (r *registerBdtxuprType) SetTimxper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxperMask)
	}
}

const (
	RegisterBdtxuprFieldTimxrepShift = 5
	RegisterBdtxuprFieldTimxrepMask  = 0x20
)

// GetTimxrep HRTIM_REPxR register update enable
func (r *registerBdtxuprType) GetTimxrep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxrepMask) != 0
}

// SetTimxrep HRTIM_REPxR register update enable
func (r *registerBdtxuprType) SetTimxrep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxrepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxrepMask)
	}
}

const (
	RegisterBdtxuprFieldTimxcmp1Shift = 6
	RegisterBdtxuprFieldTimxcmp1Mask  = 0x40
)

// GetTimxcmp1 HRTIM_CMP1xR register update enable
func (r *registerBdtxuprType) GetTimxcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxcmp1Mask) != 0
}

// SetTimxcmp1 HRTIM_CMP1xR register update enable
func (r *registerBdtxuprType) SetTimxcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxcmp1Mask)
	}
}

const (
	RegisterBdtxuprFieldTimxcmp2Shift = 7
	RegisterBdtxuprFieldTimxcmp2Mask  = 0x80
)

// GetTimxcmp2 HRTIM_CMP2xR register update enable
func (r *registerBdtxuprType) GetTimxcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxcmp2Mask) != 0
}

// SetTimxcmp2 HRTIM_CMP2xR register update enable
func (r *registerBdtxuprType) SetTimxcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxcmp2Mask)
	}
}

const (
	RegisterBdtxuprFieldTimxcmp3Shift = 8
	RegisterBdtxuprFieldTimxcmp3Mask  = 0x100
)

// GetTimxcmp3 HRTIM_CMP3xR register update enable
func (r *registerBdtxuprType) GetTimxcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxcmp3Mask) != 0
}

// SetTimxcmp3 HRTIM_CMP3xR register update enable
func (r *registerBdtxuprType) SetTimxcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxcmp3Mask)
	}
}

const (
	RegisterBdtxuprFieldTimxcmp4Shift = 9
	RegisterBdtxuprFieldTimxcmp4Mask  = 0x200
)

// GetTimxcmp4 HRTIM_CMP4xR register update enable
func (r *registerBdtxuprType) GetTimxcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxcmp4Mask) != 0
}

// SetTimxcmp4 HRTIM_CMP4xR register update enable
func (r *registerBdtxuprType) SetTimxcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxcmp4Mask)
	}
}

const (
	RegisterBdtxuprFieldTimxdtxrShift = 10
	RegisterBdtxuprFieldTimxdtxrMask  = 0x400
)

// GetTimxdtxr HRTIM_DTxR register update enable
func (r *registerBdtxuprType) GetTimxdtxr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxdtxrMask) != 0
}

// SetTimxdtxr HRTIM_DTxR register update enable
func (r *registerBdtxuprType) SetTimxdtxr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxdtxrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxdtxrMask)
	}
}

const (
	RegisterBdtxuprFieldTimxset1rShift = 11
	RegisterBdtxuprFieldTimxset1rMask  = 0x800
)

// GetTimxset1r HRTIM_SET1xR register update enable
func (r *registerBdtxuprType) GetTimxset1r() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxset1rMask) != 0
}

// SetTimxset1r HRTIM_SET1xR register update enable
func (r *registerBdtxuprType) SetTimxset1r(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxset1rMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxset1rMask)
	}
}

const (
	RegisterBdtxuprFieldTimxrst1rShift = 12
	RegisterBdtxuprFieldTimxrst1rMask  = 0x1000
)

// GetTimxrst1r HRTIM_RST1xR register update enable
func (r *registerBdtxuprType) GetTimxrst1r() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxrst1rMask) != 0
}

// SetTimxrst1r HRTIM_RST1xR register update enable
func (r *registerBdtxuprType) SetTimxrst1r(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxrst1rMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxrst1rMask)
	}
}

const (
	RegisterBdtxuprFieldTimxset2rShift = 13
	RegisterBdtxuprFieldTimxset2rMask  = 0x2000
)

// GetTimxset2r HRTIM_SET2xR register update enable
func (r *registerBdtxuprType) GetTimxset2r() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxset2rMask) != 0
}

// SetTimxset2r HRTIM_SET2xR register update enable
func (r *registerBdtxuprType) SetTimxset2r(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxset2rMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxset2rMask)
	}
}

const (
	RegisterBdtxuprFieldTimxrst2rShift = 14
	RegisterBdtxuprFieldTimxrst2rMask  = 0x4000
)

// GetTimxrst2r HRTIM_RST2xR register update enable
func (r *registerBdtxuprType) GetTimxrst2r() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxrst2rMask) != 0
}

// SetTimxrst2r HRTIM_RST2xR register update enable
func (r *registerBdtxuprType) SetTimxrst2r(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxrst2rMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxrst2rMask)
	}
}

const (
	RegisterBdtxuprFieldTimxeefr1Shift = 15
	RegisterBdtxuprFieldTimxeefr1Mask  = 0x8000
)

// GetTimxeefr1 HRTIM_EEFxR1 register update enable
func (r *registerBdtxuprType) GetTimxeefr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxeefr1Mask) != 0
}

// SetTimxeefr1 HRTIM_EEFxR1 register update enable
func (r *registerBdtxuprType) SetTimxeefr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxeefr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxeefr1Mask)
	}
}

const (
	RegisterBdtxuprFieldTimxeefr2Shift = 16
	RegisterBdtxuprFieldTimxeefr2Mask  = 0x10000
)

// GetTimxeefr2 HRTIM_EEFxR2 register update enable
func (r *registerBdtxuprType) GetTimxeefr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxeefr2Mask) != 0
}

// SetTimxeefr2 HRTIM_EEFxR2 register update enable
func (r *registerBdtxuprType) SetTimxeefr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxeefr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxeefr2Mask)
	}
}

const (
	RegisterBdtxuprFieldTimxrstrShift = 17
	RegisterBdtxuprFieldTimxrstrMask  = 0x20000
)

// GetTimxrstr HRTIM_RSTxR register update enable
func (r *registerBdtxuprType) GetTimxrstr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxrstrMask) != 0
}

// SetTimxrstr HRTIM_RSTxR register update enable
func (r *registerBdtxuprType) SetTimxrstr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxrstrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxrstrMask)
	}
}

const (
	RegisterBdtxuprFieldTimxchprShift = 18
	RegisterBdtxuprFieldTimxchprMask  = 0x40000
)

// GetTimxchpr HRTIM_CHPxR register update enable
func (r *registerBdtxuprType) GetTimxchpr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxchprMask) != 0
}

// SetTimxchpr HRTIM_CHPxR register update enable
func (r *registerBdtxuprType) SetTimxchpr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxchprMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxchprMask)
	}
}

const (
	RegisterBdtxuprFieldTimxoutrShift = 19
	RegisterBdtxuprFieldTimxoutrMask  = 0x80000
)

// GetTimxoutr HRTIM_OUTxR register update enable
func (r *registerBdtxuprType) GetTimxoutr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxoutrMask) != 0
}

// SetTimxoutr HRTIM_OUTxR register update enable
func (r *registerBdtxuprType) SetTimxoutr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxoutrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxoutrMask)
	}
}

const (
	RegisterBdtxuprFieldTimxfltrShift = 20
	RegisterBdtxuprFieldTimxfltrMask  = 0x100000
)

// GetTimxfltr HRTIM_FLTxR register update enable
func (r *registerBdtxuprType) GetTimxfltr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdtxuprFieldTimxfltrMask) != 0
}

// SetTimxfltr HRTIM_FLTxR register update enable
func (r *registerBdtxuprType) SetTimxfltr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdtxuprFieldTimxfltrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdtxuprFieldTimxfltrMask)
	}
}

// registerBdmadrType Burst DMA Data Register
type registerBdmadrType uint32

const (
	RegisterBdmadrFieldBdmadrShift = 0
	RegisterBdmadrFieldBdmadrMask  = 0xffffffff
)

// GetBdmadr Burst DMA Data register
func (r *registerBdmadrType) GetBdmadr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterBdmadrFieldBdmadrMask) >> RegisterBdmadrFieldBdmadrShift)
}

// SetBdmadr Burst DMA Data register
func (r *registerBdmadrType) SetBdmadr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBdmadrFieldBdmadrMask)|(uint32(value)<<RegisterBdmadrFieldBdmadrShift))
}
