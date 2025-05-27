//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package dbgmcu

import (
	"unsafe"
	"volatile"
)

var (
	Dbgmcu = (*_dbgmcu)(unsafe.Pointer(uintptr(0x5c001000)))
)

type _dbgmcu struct {
	Idc      registerIdcType
	Cr       registerCrType
	_        [44]byte
	Apb3fz1  registerApb3fz1Type
	Apb3fz2  registerApb3fz2Type
	Apb1lfz1 registerApb1lfz1Type
	Apb1lfz2 registerApb1lfz2Type
	_        [8]byte
	Apb2fz1  registerApb2fz1Type
	Apb2fz2  registerApb2fz2Type
	Apb4fz1  registerApb4fz1Type
	Apb4fz2  registerApb4fz2Type
}

// registerIdcType DBGMCU Identity Code Register
type registerIdcType uint32

const (
	RegisterIdcFieldDevidShift = 0
	RegisterIdcFieldDevidMask  = 0xfff
)

// GetDevid Device ID
func (r *registerIdcType) GetDevid() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterIdcFieldDevidMask) >> RegisterIdcFieldDevidShift)
}

// SetDevid Device ID
func (r *registerIdcType) SetDevid(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdcFieldDevidMask)|(uint32(value)<<RegisterIdcFieldDevidShift))
}

const (
	RegisterIdcFieldRevidShift = 16
	RegisterIdcFieldRevidMask  = 0xffff0000
)

// GetRevid Revision
func (r *registerIdcType) GetRevid() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterIdcFieldRevidMask) >> RegisterIdcFieldRevidShift)
}

// SetRevid Revision
func (r *registerIdcType) SetRevid(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdcFieldRevidMask)|(uint32(value)<<RegisterIdcFieldRevidShift))
}

// registerCrType DBGMCU Configuration Register
type registerCrType uint32

const (
	RegisterCrFieldDbgslpd1Shift = 0
	RegisterCrFieldDbgslpd1Mask  = 0x1
)

// GetDbgslpd1 Allow D1 domain debug in Sleep mode
func (r *registerCrType) GetDbgslpd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDbgslpd1Mask) != 0
}

// SetDbgslpd1 Allow D1 domain debug in Sleep mode
func (r *registerCrType) SetDbgslpd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDbgslpd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDbgslpd1Mask)
	}
}

const (
	RegisterCrFieldDbgstpd1Shift = 1
	RegisterCrFieldDbgstpd1Mask  = 0x2
)

// GetDbgstpd1 Allow D1 domain debug in Stop mode
func (r *registerCrType) GetDbgstpd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDbgstpd1Mask) != 0
}

// SetDbgstpd1 Allow D1 domain debug in Stop mode
func (r *registerCrType) SetDbgstpd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDbgstpd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDbgstpd1Mask)
	}
}

const (
	RegisterCrFieldDbgstbd1Shift = 2
	RegisterCrFieldDbgstbd1Mask  = 0x4
)

// GetDbgstbd1 Allow D1 domain debug in Standby mode
func (r *registerCrType) GetDbgstbd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDbgstbd1Mask) != 0
}

// SetDbgstbd1 Allow D1 domain debug in Standby mode
func (r *registerCrType) SetDbgstbd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDbgstbd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDbgstbd1Mask)
	}
}

const (
	RegisterCrFieldDbgslpd2Shift = 3
	RegisterCrFieldDbgslpd2Mask  = 0x8
)

// GetDbgslpd2 Allow D2 domain debug in Sleep mode
func (r *registerCrType) GetDbgslpd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDbgslpd2Mask) != 0
}

// SetDbgslpd2 Allow D2 domain debug in Sleep mode
func (r *registerCrType) SetDbgslpd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDbgslpd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDbgslpd2Mask)
	}
}

const (
	RegisterCrFieldDbgstpd2Shift = 4
	RegisterCrFieldDbgstpd2Mask  = 0x10
)

// GetDbgstpd2 Allow D2 domain debug in Stop mode
func (r *registerCrType) GetDbgstpd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDbgstpd2Mask) != 0
}

// SetDbgstpd2 Allow D2 domain debug in Stop mode
func (r *registerCrType) SetDbgstpd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDbgstpd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDbgstpd2Mask)
	}
}

const (
	RegisterCrFieldDbgstbd2Shift = 5
	RegisterCrFieldDbgstbd2Mask  = 0x20
)

// GetDbgstbd2 Allow D2 domain debug in Standby mode
func (r *registerCrType) GetDbgstbd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDbgstbd2Mask) != 0
}

// SetDbgstbd2 Allow D2 domain debug in Standby mode
func (r *registerCrType) SetDbgstbd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDbgstbd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDbgstbd2Mask)
	}
}

const (
	RegisterCrFieldDbgstpd3Shift = 7
	RegisterCrFieldDbgstpd3Mask  = 0x80
)

// GetDbgstpd3 Allow debug in D3 Stop mode
func (r *registerCrType) GetDbgstpd3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDbgstpd3Mask) != 0
}

// SetDbgstpd3 Allow debug in D3 Stop mode
func (r *registerCrType) SetDbgstpd3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDbgstpd3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDbgstpd3Mask)
	}
}

const (
	RegisterCrFieldDbgstbd3Shift = 8
	RegisterCrFieldDbgstbd3Mask  = 0x100
)

// GetDbgstbd3 Allow debug in D3 Standby mode
func (r *registerCrType) GetDbgstbd3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDbgstbd3Mask) != 0
}

// SetDbgstbd3 Allow debug in D3 Standby mode
func (r *registerCrType) SetDbgstbd3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDbgstbd3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDbgstbd3Mask)
	}
}

const (
	RegisterCrFieldTraceclkenShift = 20
	RegisterCrFieldTraceclkenMask  = 0x100000
)

// GetTraceclken Trace port clock enable
func (r *registerCrType) GetTraceclken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTraceclkenMask) != 0
}

// SetTraceclken Trace port clock enable
func (r *registerCrType) SetTraceclken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTraceclkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTraceclkenMask)
	}
}

const (
	RegisterCrFieldD1dbgckenShift = 21
	RegisterCrFieldD1dbgckenMask  = 0x200000
)

// GetD1dbgcken D1 debug clock enable
func (r *registerCrType) GetD1dbgcken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldD1dbgckenMask) != 0
}

// SetD1dbgcken D1 debug clock enable
func (r *registerCrType) SetD1dbgcken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldD1dbgckenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldD1dbgckenMask)
	}
}

const (
	RegisterCrFieldD3dbgckenShift = 22
	RegisterCrFieldD3dbgckenMask  = 0x400000
)

// GetD3dbgcken D3 debug clock enable
func (r *registerCrType) GetD3dbgcken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldD3dbgckenMask) != 0
}

// SetD3dbgcken D3 debug clock enable
func (r *registerCrType) SetD3dbgcken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldD3dbgckenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldD3dbgckenMask)
	}
}

const (
	RegisterCrFieldTrgoenShift = 28
	RegisterCrFieldTrgoenMask  = 0x10000000
)

// GetTrgoen External trigger output enable
func (r *registerCrType) GetTrgoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTrgoenMask) != 0
}

// SetTrgoen External trigger output enable
func (r *registerCrType) SetTrgoen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTrgoenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTrgoenMask)
	}
}

// registerApb3fz1Type DBGMCU APB3 peripheral freeze register CPU1
type registerApb3fz1Type uint32

const (
	RegisterApb3fz1FieldWwdg1Shift = 6
	RegisterApb3fz1FieldWwdg1Mask  = 0x40
)

// GetWwdg1 WWDG1 stop in debug
func (r *registerApb3fz1Type) GetWwdg1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb3fz1FieldWwdg1Mask) != 0
}

// SetWwdg1 WWDG1 stop in debug
func (r *registerApb3fz1Type) SetWwdg1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb3fz1FieldWwdg1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb3fz1FieldWwdg1Mask)
	}
}

// registerApb3fz2Type DBGMCU APB3 peripheral freeze register CPU2
type registerApb3fz2Type uint32

const (
	RegisterApb3fz2FieldWwdg1Shift = 6
	RegisterApb3fz2FieldWwdg1Mask  = 0x40
)

// GetWwdg1 WWDG1 stop in debug
func (r *registerApb3fz2Type) GetWwdg1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb3fz2FieldWwdg1Mask) != 0
}

// SetWwdg1 WWDG1 stop in debug
func (r *registerApb3fz2Type) SetWwdg1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb3fz2FieldWwdg1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb3fz2FieldWwdg1Mask)
	}
}

// registerApb1lfz1Type DBGMCU APB1L peripheral freeze register
type registerApb1lfz1Type uint32

const (
	RegisterApb1lfz1FieldDbgtim2Shift = 0
	RegisterApb1lfz1FieldDbgtim2Mask  = 0x1
)

// GetDbgtim2 TIM2 stop in debug
func (r *registerApb1lfz1Type) GetDbgtim2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgtim2Mask) != 0
}

// SetDbgtim2 TIM2 stop in debug
func (r *registerApb1lfz1Type) SetDbgtim2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgtim2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgtim2Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbgtim3Shift = 1
	RegisterApb1lfz1FieldDbgtim3Mask  = 0x2
)

// GetDbgtim3 TIM3 stop in debug
func (r *registerApb1lfz1Type) GetDbgtim3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgtim3Mask) != 0
}

// SetDbgtim3 TIM3 stop in debug
func (r *registerApb1lfz1Type) SetDbgtim3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgtim3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgtim3Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbgtim4Shift = 2
	RegisterApb1lfz1FieldDbgtim4Mask  = 0x4
)

// GetDbgtim4 TIM4 stop in debug
func (r *registerApb1lfz1Type) GetDbgtim4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgtim4Mask) != 0
}

// SetDbgtim4 TIM4 stop in debug
func (r *registerApb1lfz1Type) SetDbgtim4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgtim4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgtim4Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbgtim5Shift = 3
	RegisterApb1lfz1FieldDbgtim5Mask  = 0x8
)

// GetDbgtim5 TIM5 stop in debug
func (r *registerApb1lfz1Type) GetDbgtim5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgtim5Mask) != 0
}

// SetDbgtim5 TIM5 stop in debug
func (r *registerApb1lfz1Type) SetDbgtim5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgtim5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgtim5Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbgtim6Shift = 4
	RegisterApb1lfz1FieldDbgtim6Mask  = 0x10
)

// GetDbgtim6 TIM6 stop in debug
func (r *registerApb1lfz1Type) GetDbgtim6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgtim6Mask) != 0
}

// SetDbgtim6 TIM6 stop in debug
func (r *registerApb1lfz1Type) SetDbgtim6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgtim6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgtim6Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbgtim7Shift = 5
	RegisterApb1lfz1FieldDbgtim7Mask  = 0x20
)

// GetDbgtim7 TIM7 stop in debug
func (r *registerApb1lfz1Type) GetDbgtim7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgtim7Mask) != 0
}

// SetDbgtim7 TIM7 stop in debug
func (r *registerApb1lfz1Type) SetDbgtim7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgtim7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgtim7Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbgtim12Shift = 6
	RegisterApb1lfz1FieldDbgtim12Mask  = 0x40
)

// GetDbgtim12 TIM12 stop in debug
func (r *registerApb1lfz1Type) GetDbgtim12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgtim12Mask) != 0
}

// SetDbgtim12 TIM12 stop in debug
func (r *registerApb1lfz1Type) SetDbgtim12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgtim12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgtim12Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbgtim13Shift = 7
	RegisterApb1lfz1FieldDbgtim13Mask  = 0x80
)

// GetDbgtim13 TIM13 stop in debug
func (r *registerApb1lfz1Type) GetDbgtim13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgtim13Mask) != 0
}

// SetDbgtim13 TIM13 stop in debug
func (r *registerApb1lfz1Type) SetDbgtim13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgtim13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgtim13Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbgtim14Shift = 8
	RegisterApb1lfz1FieldDbgtim14Mask  = 0x100
)

// GetDbgtim14 TIM14 stop in debug
func (r *registerApb1lfz1Type) GetDbgtim14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgtim14Mask) != 0
}

// SetDbgtim14 TIM14 stop in debug
func (r *registerApb1lfz1Type) SetDbgtim14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgtim14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgtim14Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbglptim1Shift = 9
	RegisterApb1lfz1FieldDbglptim1Mask  = 0x200
)

// GetDbglptim1 LPTIM1 stop in debug
func (r *registerApb1lfz1Type) GetDbglptim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbglptim1Mask) != 0
}

// SetDbglptim1 LPTIM1 stop in debug
func (r *registerApb1lfz1Type) SetDbglptim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbglptim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbglptim1Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbgwwdg2Shift = 11
	RegisterApb1lfz1FieldDbgwwdg2Mask  = 0x800
)

// GetDbgwwdg2 WWDG2 stop in debug
func (r *registerApb1lfz1Type) GetDbgwwdg2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgwwdg2Mask) != 0
}

// SetDbgwwdg2 WWDG2 stop in debug
func (r *registerApb1lfz1Type) SetDbgwwdg2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgwwdg2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgwwdg2Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbgi2c1Shift = 21
	RegisterApb1lfz1FieldDbgi2c1Mask  = 0x200000
)

// GetDbgi2c1 I2C1 SMBUS timeout stop in debug
func (r *registerApb1lfz1Type) GetDbgi2c1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgi2c1Mask) != 0
}

// SetDbgi2c1 I2C1 SMBUS timeout stop in debug
func (r *registerApb1lfz1Type) SetDbgi2c1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgi2c1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgi2c1Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbgi2c2Shift = 22
	RegisterApb1lfz1FieldDbgi2c2Mask  = 0x400000
)

// GetDbgi2c2 I2C2 SMBUS timeout stop in debug
func (r *registerApb1lfz1Type) GetDbgi2c2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgi2c2Mask) != 0
}

// SetDbgi2c2 I2C2 SMBUS timeout stop in debug
func (r *registerApb1lfz1Type) SetDbgi2c2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgi2c2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgi2c2Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbgi2c3Shift = 23
	RegisterApb1lfz1FieldDbgi2c3Mask  = 0x800000
)

// GetDbgi2c3 I2C3 SMBUS timeout stop in debug
func (r *registerApb1lfz1Type) GetDbgi2c3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbgi2c3Mask) != 0
}

// SetDbgi2c3 I2C3 SMBUS timeout stop in debug
func (r *registerApb1lfz1Type) SetDbgi2c3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbgi2c3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbgi2c3Mask)
	}
}

// registerApb1lfz2Type DBGMCU APB1L peripheral freeze register CPU2
type registerApb1lfz2Type uint32

const (
	RegisterApb1lfz2FieldDbgtim2Shift = 0
	RegisterApb1lfz2FieldDbgtim2Mask  = 0x1
)

// GetDbgtim2 TIM2 stop in debug
func (r *registerApb1lfz2Type) GetDbgtim2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgtim2Mask) != 0
}

// SetDbgtim2 TIM2 stop in debug
func (r *registerApb1lfz2Type) SetDbgtim2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgtim2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgtim2Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbgtim3Shift = 1
	RegisterApb1lfz2FieldDbgtim3Mask  = 0x2
)

// GetDbgtim3 TIM3 stop in debug
func (r *registerApb1lfz2Type) GetDbgtim3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgtim3Mask) != 0
}

// SetDbgtim3 TIM3 stop in debug
func (r *registerApb1lfz2Type) SetDbgtim3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgtim3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgtim3Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbgtim4Shift = 2
	RegisterApb1lfz2FieldDbgtim4Mask  = 0x4
)

// GetDbgtim4 TIM4 stop in debug
func (r *registerApb1lfz2Type) GetDbgtim4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgtim4Mask) != 0
}

// SetDbgtim4 TIM4 stop in debug
func (r *registerApb1lfz2Type) SetDbgtim4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgtim4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgtim4Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbgtim5Shift = 3
	RegisterApb1lfz2FieldDbgtim5Mask  = 0x8
)

// GetDbgtim5 TIM5 stop in debug
func (r *registerApb1lfz2Type) GetDbgtim5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgtim5Mask) != 0
}

// SetDbgtim5 TIM5 stop in debug
func (r *registerApb1lfz2Type) SetDbgtim5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgtim5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgtim5Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbgtim6Shift = 4
	RegisterApb1lfz2FieldDbgtim6Mask  = 0x10
)

// GetDbgtim6 TIM6 stop in debug
func (r *registerApb1lfz2Type) GetDbgtim6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgtim6Mask) != 0
}

// SetDbgtim6 TIM6 stop in debug
func (r *registerApb1lfz2Type) SetDbgtim6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgtim6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgtim6Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbgtim7Shift = 5
	RegisterApb1lfz2FieldDbgtim7Mask  = 0x20
)

// GetDbgtim7 TIM4 stop in debug
func (r *registerApb1lfz2Type) GetDbgtim7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgtim7Mask) != 0
}

// SetDbgtim7 TIM4 stop in debug
func (r *registerApb1lfz2Type) SetDbgtim7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgtim7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgtim7Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbgtim12Shift = 6
	RegisterApb1lfz2FieldDbgtim12Mask  = 0x40
)

// GetDbgtim12 TIM12 stop in debug
func (r *registerApb1lfz2Type) GetDbgtim12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgtim12Mask) != 0
}

// SetDbgtim12 TIM12 stop in debug
func (r *registerApb1lfz2Type) SetDbgtim12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgtim12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgtim12Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbgtim13Shift = 7
	RegisterApb1lfz2FieldDbgtim13Mask  = 0x80
)

// GetDbgtim13 TIM13 stop in debug
func (r *registerApb1lfz2Type) GetDbgtim13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgtim13Mask) != 0
}

// SetDbgtim13 TIM13 stop in debug
func (r *registerApb1lfz2Type) SetDbgtim13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgtim13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgtim13Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbgtim14Shift = 8
	RegisterApb1lfz2FieldDbgtim14Mask  = 0x100
)

// GetDbgtim14 TIM14 stop in debug
func (r *registerApb1lfz2Type) GetDbgtim14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgtim14Mask) != 0
}

// SetDbgtim14 TIM14 stop in debug
func (r *registerApb1lfz2Type) SetDbgtim14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgtim14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgtim14Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbglptim1Shift = 9
	RegisterApb1lfz2FieldDbglptim1Mask  = 0x200
)

// GetDbglptim1 LPTIM1 stop in debug
func (r *registerApb1lfz2Type) GetDbglptim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbglptim1Mask) != 0
}

// SetDbglptim1 LPTIM1 stop in debug
func (r *registerApb1lfz2Type) SetDbglptim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbglptim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbglptim1Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbgwwdg2Shift = 11
	RegisterApb1lfz2FieldDbgwwdg2Mask  = 0x800
)

// GetDbgwwdg2 WWDG2 stop in debug
func (r *registerApb1lfz2Type) GetDbgwwdg2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgwwdg2Mask) != 0
}

// SetDbgwwdg2 WWDG2 stop in debug
func (r *registerApb1lfz2Type) SetDbgwwdg2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgwwdg2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgwwdg2Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbgi2c1Shift = 21
	RegisterApb1lfz2FieldDbgi2c1Mask  = 0x200000
)

// GetDbgi2c1 I2C1 SMBUS timeout stop in debug
func (r *registerApb1lfz2Type) GetDbgi2c1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgi2c1Mask) != 0
}

// SetDbgi2c1 I2C1 SMBUS timeout stop in debug
func (r *registerApb1lfz2Type) SetDbgi2c1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgi2c1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgi2c1Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbgi2c2Shift = 22
	RegisterApb1lfz2FieldDbgi2c2Mask  = 0x400000
)

// GetDbgi2c2 I2C2 SMBUS timeout stop in debug
func (r *registerApb1lfz2Type) GetDbgi2c2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgi2c2Mask) != 0
}

// SetDbgi2c2 I2C2 SMBUS timeout stop in debug
func (r *registerApb1lfz2Type) SetDbgi2c2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgi2c2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgi2c2Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbgi2c3Shift = 23
	RegisterApb1lfz2FieldDbgi2c3Mask  = 0x800000
)

// GetDbgi2c3 I2C3 SMBUS timeout stop in debug
func (r *registerApb1lfz2Type) GetDbgi2c3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbgi2c3Mask) != 0
}

// SetDbgi2c3 I2C3 SMBUS timeout stop in debug
func (r *registerApb1lfz2Type) SetDbgi2c3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbgi2c3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbgi2c3Mask)
	}
}

// registerApb2fz1Type DBGMCU APB2 peripheral freeze register
type registerApb2fz1Type uint32

const (
	RegisterApb2fz1FieldDbgtim1Shift = 0
	RegisterApb2fz1FieldDbgtim1Mask  = 0x1
)

// GetDbgtim1 TIM1 stop in debug
func (r *registerApb2fz1Type) GetDbgtim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz1FieldDbgtim1Mask) != 0
}

// SetDbgtim1 TIM1 stop in debug
func (r *registerApb2fz1Type) SetDbgtim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz1FieldDbgtim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz1FieldDbgtim1Mask)
	}
}

const (
	RegisterApb2fz1FieldDbgtim8Shift = 1
	RegisterApb2fz1FieldDbgtim8Mask  = 0x2
)

// GetDbgtim8 TIM8 stop in debug
func (r *registerApb2fz1Type) GetDbgtim8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz1FieldDbgtim8Mask) != 0
}

// SetDbgtim8 TIM8 stop in debug
func (r *registerApb2fz1Type) SetDbgtim8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz1FieldDbgtim8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz1FieldDbgtim8Mask)
	}
}

const (
	RegisterApb2fz1FieldDbgtim15Shift = 16
	RegisterApb2fz1FieldDbgtim15Mask  = 0x10000
)

// GetDbgtim15 TIM15 stop in debug
func (r *registerApb2fz1Type) GetDbgtim15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz1FieldDbgtim15Mask) != 0
}

// SetDbgtim15 TIM15 stop in debug
func (r *registerApb2fz1Type) SetDbgtim15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz1FieldDbgtim15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz1FieldDbgtim15Mask)
	}
}

const (
	RegisterApb2fz1FieldDbgtim16Shift = 17
	RegisterApb2fz1FieldDbgtim16Mask  = 0x20000
)

// GetDbgtim16 TIM16 stop in debug
func (r *registerApb2fz1Type) GetDbgtim16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz1FieldDbgtim16Mask) != 0
}

// SetDbgtim16 TIM16 stop in debug
func (r *registerApb2fz1Type) SetDbgtim16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz1FieldDbgtim16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz1FieldDbgtim16Mask)
	}
}

const (
	RegisterApb2fz1FieldDbgtim17Shift = 18
	RegisterApb2fz1FieldDbgtim17Mask  = 0x40000
)

// GetDbgtim17 TIM17 stop in debug
func (r *registerApb2fz1Type) GetDbgtim17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz1FieldDbgtim17Mask) != 0
}

// SetDbgtim17 TIM17 stop in debug
func (r *registerApb2fz1Type) SetDbgtim17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz1FieldDbgtim17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz1FieldDbgtim17Mask)
	}
}

const (
	RegisterApb2fz1FieldDbghrtimShift = 29
	RegisterApb2fz1FieldDbghrtimMask  = 0x20000000
)

// GetDbghrtim HRTIM stop in debug
func (r *registerApb2fz1Type) GetDbghrtim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz1FieldDbghrtimMask) != 0
}

// SetDbghrtim HRTIM stop in debug
func (r *registerApb2fz1Type) SetDbghrtim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz1FieldDbghrtimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz1FieldDbghrtimMask)
	}
}

// registerApb2fz2Type DBGMCU APB2 peripheral freeze register CPU2
type registerApb2fz2Type uint32

const (
	RegisterApb2fz2FieldDbgtim1Shift = 0
	RegisterApb2fz2FieldDbgtim1Mask  = 0x1
)

// GetDbgtim1 TIM1 stop in debug
func (r *registerApb2fz2Type) GetDbgtim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz2FieldDbgtim1Mask) != 0
}

// SetDbgtim1 TIM1 stop in debug
func (r *registerApb2fz2Type) SetDbgtim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz2FieldDbgtim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz2FieldDbgtim1Mask)
	}
}

const (
	RegisterApb2fz2FieldDbgtim8Shift = 1
	RegisterApb2fz2FieldDbgtim8Mask  = 0x2
)

// GetDbgtim8 TIM8 stop in debug
func (r *registerApb2fz2Type) GetDbgtim8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz2FieldDbgtim8Mask) != 0
}

// SetDbgtim8 TIM8 stop in debug
func (r *registerApb2fz2Type) SetDbgtim8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz2FieldDbgtim8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz2FieldDbgtim8Mask)
	}
}

const (
	RegisterApb2fz2FieldDbgtim15Shift = 16
	RegisterApb2fz2FieldDbgtim15Mask  = 0x10000
)

// GetDbgtim15 TIM15 stop in debug
func (r *registerApb2fz2Type) GetDbgtim15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz2FieldDbgtim15Mask) != 0
}

// SetDbgtim15 TIM15 stop in debug
func (r *registerApb2fz2Type) SetDbgtim15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz2FieldDbgtim15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz2FieldDbgtim15Mask)
	}
}

const (
	RegisterApb2fz2FieldDbgtim16Shift = 17
	RegisterApb2fz2FieldDbgtim16Mask  = 0x20000
)

// GetDbgtim16 TIM16 stop in debug
func (r *registerApb2fz2Type) GetDbgtim16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz2FieldDbgtim16Mask) != 0
}

// SetDbgtim16 TIM16 stop in debug
func (r *registerApb2fz2Type) SetDbgtim16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz2FieldDbgtim16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz2FieldDbgtim16Mask)
	}
}

const (
	RegisterApb2fz2FieldDbgtim17Shift = 18
	RegisterApb2fz2FieldDbgtim17Mask  = 0x40000
)

// GetDbgtim17 TIM17 stop in debug
func (r *registerApb2fz2Type) GetDbgtim17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz2FieldDbgtim17Mask) != 0
}

// SetDbgtim17 TIM17 stop in debug
func (r *registerApb2fz2Type) SetDbgtim17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz2FieldDbgtim17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz2FieldDbgtim17Mask)
	}
}

const (
	RegisterApb2fz2FieldDbghrtimShift = 29
	RegisterApb2fz2FieldDbghrtimMask  = 0x20000000
)

// GetDbghrtim HRTIM stop in debug
func (r *registerApb2fz2Type) GetDbghrtim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz2FieldDbghrtimMask) != 0
}

// SetDbghrtim HRTIM stop in debug
func (r *registerApb2fz2Type) SetDbghrtim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz2FieldDbghrtimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz2FieldDbghrtimMask)
	}
}

// registerApb4fz1Type DBGMCU APB4 peripheral freeze register
type registerApb4fz1Type uint32

const (
	RegisterApb4fz1FieldDbgi2c4Shift = 7
	RegisterApb4fz1FieldDbgi2c4Mask  = 0x80
)

// GetDbgi2c4 I2C4 SMBUS timeout stop in debug
func (r *registerApb4fz1Type) GetDbgi2c4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbgi2c4Mask) != 0
}

// SetDbgi2c4 I2C4 SMBUS timeout stop in debug
func (r *registerApb4fz1Type) SetDbgi2c4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbgi2c4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbgi2c4Mask)
	}
}

const (
	RegisterApb4fz1FieldDbglptim2Shift = 9
	RegisterApb4fz1FieldDbglptim2Mask  = 0x200
)

// GetDbglptim2 LPTIM2 stop in debug
func (r *registerApb4fz1Type) GetDbglptim2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbglptim2Mask) != 0
}

// SetDbglptim2 LPTIM2 stop in debug
func (r *registerApb4fz1Type) SetDbglptim2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbglptim2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbglptim2Mask)
	}
}

const (
	RegisterApb4fz1FieldDbglptim3Shift = 10
	RegisterApb4fz1FieldDbglptim3Mask  = 0x400
)

// GetDbglptim3 LPTIM2 stop in debug
func (r *registerApb4fz1Type) GetDbglptim3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbglptim3Mask) != 0
}

// SetDbglptim3 LPTIM2 stop in debug
func (r *registerApb4fz1Type) SetDbglptim3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbglptim3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbglptim3Mask)
	}
}

const (
	RegisterApb4fz1FieldDbglptim4Shift = 11
	RegisterApb4fz1FieldDbglptim4Mask  = 0x800
)

// GetDbglptim4 LPTIM4 stop in debug
func (r *registerApb4fz1Type) GetDbglptim4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbglptim4Mask) != 0
}

// SetDbglptim4 LPTIM4 stop in debug
func (r *registerApb4fz1Type) SetDbglptim4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbglptim4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbglptim4Mask)
	}
}

const (
	RegisterApb4fz1FieldDbglptim5Shift = 12
	RegisterApb4fz1FieldDbglptim5Mask  = 0x1000
)

// GetDbglptim5 LPTIM5 stop in debug
func (r *registerApb4fz1Type) GetDbglptim5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbglptim5Mask) != 0
}

// SetDbglptim5 LPTIM5 stop in debug
func (r *registerApb4fz1Type) SetDbglptim5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbglptim5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbglptim5Mask)
	}
}

const (
	RegisterApb4fz1FieldDbgrtcShift = 16
	RegisterApb4fz1FieldDbgrtcMask  = 0x10000
)

// GetDbgrtc RTC stop in debug
func (r *registerApb4fz1Type) GetDbgrtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbgrtcMask) != 0
}

// SetDbgrtc RTC stop in debug
func (r *registerApb4fz1Type) SetDbgrtc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbgrtcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbgrtcMask)
	}
}

const (
	RegisterApb4fz1FieldDbgwdglsd1Shift = 18
	RegisterApb4fz1FieldDbgwdglsd1Mask  = 0x40000
)

// GetDbgwdglsd1 Independent watchdog for D1 stop in debug
func (r *registerApb4fz1Type) GetDbgwdglsd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbgwdglsd1Mask) != 0
}

// SetDbgwdglsd1 Independent watchdog for D1 stop in debug
func (r *registerApb4fz1Type) SetDbgwdglsd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbgwdglsd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbgwdglsd1Mask)
	}
}

const (
	RegisterApb4fz1FieldDbgwdglsd2Shift = 19
	RegisterApb4fz1FieldDbgwdglsd2Mask  = 0x80000
)

// GetDbgwdglsd2 Independent watchdog for D2 stop in debug
func (r *registerApb4fz1Type) GetDbgwdglsd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbgwdglsd2Mask) != 0
}

// SetDbgwdglsd2 Independent watchdog for D2 stop in debug
func (r *registerApb4fz1Type) SetDbgwdglsd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbgwdglsd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbgwdglsd2Mask)
	}
}

// registerApb4fz2Type DBGMCU APB4 peripheral freeze register CPU2
type registerApb4fz2Type uint32

const (
	RegisterApb4fz2FieldDbgi2c4Shift = 7
	RegisterApb4fz2FieldDbgi2c4Mask  = 0x80
)

// GetDbgi2c4 I2C4 SMBUS timeout stop in debug
func (r *registerApb4fz2Type) GetDbgi2c4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbgi2c4Mask) != 0
}

// SetDbgi2c4 I2C4 SMBUS timeout stop in debug
func (r *registerApb4fz2Type) SetDbgi2c4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbgi2c4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbgi2c4Mask)
	}
}

const (
	RegisterApb4fz2FieldDbglptim2Shift = 9
	RegisterApb4fz2FieldDbglptim2Mask  = 0x200
)

// GetDbglptim2 LPTIM2 stop in debug
func (r *registerApb4fz2Type) GetDbglptim2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbglptim2Mask) != 0
}

// SetDbglptim2 LPTIM2 stop in debug
func (r *registerApb4fz2Type) SetDbglptim2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbglptim2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbglptim2Mask)
	}
}

const (
	RegisterApb4fz2FieldDbglptim3Shift = 10
	RegisterApb4fz2FieldDbglptim3Mask  = 0x400
)

// GetDbglptim3 LPTIM2 stop in debug
func (r *registerApb4fz2Type) GetDbglptim3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbglptim3Mask) != 0
}

// SetDbglptim3 LPTIM2 stop in debug
func (r *registerApb4fz2Type) SetDbglptim3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbglptim3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbglptim3Mask)
	}
}

const (
	RegisterApb4fz2FieldDbglptim4Shift = 11
	RegisterApb4fz2FieldDbglptim4Mask  = 0x800
)

// GetDbglptim4 LPTIM4 stop in debug
func (r *registerApb4fz2Type) GetDbglptim4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbglptim4Mask) != 0
}

// SetDbglptim4 LPTIM4 stop in debug
func (r *registerApb4fz2Type) SetDbglptim4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbglptim4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbglptim4Mask)
	}
}

const (
	RegisterApb4fz2FieldDbglptim5Shift = 12
	RegisterApb4fz2FieldDbglptim5Mask  = 0x1000
)

// GetDbglptim5 LPTIM5 stop in debug
func (r *registerApb4fz2Type) GetDbglptim5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbglptim5Mask) != 0
}

// SetDbglptim5 LPTIM5 stop in debug
func (r *registerApb4fz2Type) SetDbglptim5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbglptim5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbglptim5Mask)
	}
}

const (
	RegisterApb4fz2FieldDbgrtcShift = 16
	RegisterApb4fz2FieldDbgrtcMask  = 0x10000
)

// GetDbgrtc RTC stop in debug
func (r *registerApb4fz2Type) GetDbgrtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbgrtcMask) != 0
}

// SetDbgrtc RTC stop in debug
func (r *registerApb4fz2Type) SetDbgrtc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbgrtcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbgrtcMask)
	}
}

const (
	RegisterApb4fz2FieldDbgwdglsd1Shift = 18
	RegisterApb4fz2FieldDbgwdglsd1Mask  = 0x40000
)

// GetDbgwdglsd1 LS watchdog for D1 stop in debug
func (r *registerApb4fz2Type) GetDbgwdglsd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbgwdglsd1Mask) != 0
}

// SetDbgwdglsd1 LS watchdog for D1 stop in debug
func (r *registerApb4fz2Type) SetDbgwdglsd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbgwdglsd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbgwdglsd1Mask)
	}
}

const (
	RegisterApb4fz2FieldDbgwdglsd2Shift = 19
	RegisterApb4fz2FieldDbgwdglsd2Mask  = 0x80000
)

// GetDbgwdglsd2 LS watchdog for D2 stop in debug
func (r *registerApb4fz2Type) GetDbgwdglsd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbgwdglsd2Mask) != 0
}

// SetDbgwdglsd2 LS watchdog for D2 stop in debug
func (r *registerApb4fz2Type) SetDbgwdglsd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbgwdglsd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbgwdglsd2Mask)
	}
}
