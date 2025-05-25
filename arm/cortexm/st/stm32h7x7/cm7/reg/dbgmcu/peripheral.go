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
	RegisterIdcFieldDev_idShift = 0
	RegisterIdcFieldDev_idMask  = 0xfff
)

// GetDev_id Device ID
func (r *registerIdcType) GetDev_id() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterIdcFieldDev_idMask) >> RegisterIdcFieldDev_idShift)
}

// SetDev_id Device ID
func (r *registerIdcType) SetDev_id(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdcFieldDev_idMask)|(uint32(value)<<RegisterIdcFieldDev_idShift))
}

const (
	RegisterIdcFieldRev_idShift = 16
	RegisterIdcFieldRev_idMask  = 0xffff0000
)

// GetRev_id Revision
func (r *registerIdcType) GetRev_id() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterIdcFieldRev_idMask) >> RegisterIdcFieldRev_idShift)
}

// SetRev_id Revision
func (r *registerIdcType) SetRev_id(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdcFieldRev_idMask)|(uint32(value)<<RegisterIdcFieldRev_idShift))
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
	RegisterApb1lfz1FieldDbg_tim2Shift = 0
	RegisterApb1lfz1FieldDbg_tim2Mask  = 0x1
)

// GetDbg_tim2 TIM2 stop in debug
func (r *registerApb1lfz1Type) GetDbg_tim2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_tim2Mask) != 0
}

// SetDbg_tim2 TIM2 stop in debug
func (r *registerApb1lfz1Type) SetDbg_tim2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_tim2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_tim2Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_tim3Shift = 1
	RegisterApb1lfz1FieldDbg_tim3Mask  = 0x2
)

// GetDbg_tim3 TIM3 stop in debug
func (r *registerApb1lfz1Type) GetDbg_tim3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_tim3Mask) != 0
}

// SetDbg_tim3 TIM3 stop in debug
func (r *registerApb1lfz1Type) SetDbg_tim3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_tim3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_tim3Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_tim4Shift = 2
	RegisterApb1lfz1FieldDbg_tim4Mask  = 0x4
)

// GetDbg_tim4 TIM4 stop in debug
func (r *registerApb1lfz1Type) GetDbg_tim4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_tim4Mask) != 0
}

// SetDbg_tim4 TIM4 stop in debug
func (r *registerApb1lfz1Type) SetDbg_tim4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_tim4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_tim4Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_tim5Shift = 3
	RegisterApb1lfz1FieldDbg_tim5Mask  = 0x8
)

// GetDbg_tim5 TIM5 stop in debug
func (r *registerApb1lfz1Type) GetDbg_tim5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_tim5Mask) != 0
}

// SetDbg_tim5 TIM5 stop in debug
func (r *registerApb1lfz1Type) SetDbg_tim5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_tim5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_tim5Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_tim6Shift = 4
	RegisterApb1lfz1FieldDbg_tim6Mask  = 0x10
)

// GetDbg_tim6 TIM6 stop in debug
func (r *registerApb1lfz1Type) GetDbg_tim6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_tim6Mask) != 0
}

// SetDbg_tim6 TIM6 stop in debug
func (r *registerApb1lfz1Type) SetDbg_tim6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_tim6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_tim6Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_tim7Shift = 5
	RegisterApb1lfz1FieldDbg_tim7Mask  = 0x20
)

// GetDbg_tim7 TIM7 stop in debug
func (r *registerApb1lfz1Type) GetDbg_tim7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_tim7Mask) != 0
}

// SetDbg_tim7 TIM7 stop in debug
func (r *registerApb1lfz1Type) SetDbg_tim7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_tim7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_tim7Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_tim12Shift = 6
	RegisterApb1lfz1FieldDbg_tim12Mask  = 0x40
)

// GetDbg_tim12 TIM12 stop in debug
func (r *registerApb1lfz1Type) GetDbg_tim12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_tim12Mask) != 0
}

// SetDbg_tim12 TIM12 stop in debug
func (r *registerApb1lfz1Type) SetDbg_tim12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_tim12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_tim12Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_tim13Shift = 7
	RegisterApb1lfz1FieldDbg_tim13Mask  = 0x80
)

// GetDbg_tim13 TIM13 stop in debug
func (r *registerApb1lfz1Type) GetDbg_tim13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_tim13Mask) != 0
}

// SetDbg_tim13 TIM13 stop in debug
func (r *registerApb1lfz1Type) SetDbg_tim13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_tim13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_tim13Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_tim14Shift = 8
	RegisterApb1lfz1FieldDbg_tim14Mask  = 0x100
)

// GetDbg_tim14 TIM14 stop in debug
func (r *registerApb1lfz1Type) GetDbg_tim14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_tim14Mask) != 0
}

// SetDbg_tim14 TIM14 stop in debug
func (r *registerApb1lfz1Type) SetDbg_tim14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_tim14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_tim14Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_lptim1Shift = 9
	RegisterApb1lfz1FieldDbg_lptim1Mask  = 0x200
)

// GetDbg_lptim1 LPTIM1 stop in debug
func (r *registerApb1lfz1Type) GetDbg_lptim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_lptim1Mask) != 0
}

// SetDbg_lptim1 LPTIM1 stop in debug
func (r *registerApb1lfz1Type) SetDbg_lptim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_lptim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_lptim1Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_wwdg2Shift = 11
	RegisterApb1lfz1FieldDbg_wwdg2Mask  = 0x800
)

// GetDbg_wwdg2 WWDG2 stop in debug
func (r *registerApb1lfz1Type) GetDbg_wwdg2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_wwdg2Mask) != 0
}

// SetDbg_wwdg2 WWDG2 stop in debug
func (r *registerApb1lfz1Type) SetDbg_wwdg2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_wwdg2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_wwdg2Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_i2c1Shift = 21
	RegisterApb1lfz1FieldDbg_i2c1Mask  = 0x200000
)

// GetDbg_i2c1 I2C1 SMBUS timeout stop in debug
func (r *registerApb1lfz1Type) GetDbg_i2c1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_i2c1Mask) != 0
}

// SetDbg_i2c1 I2C1 SMBUS timeout stop in debug
func (r *registerApb1lfz1Type) SetDbg_i2c1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_i2c1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_i2c1Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_i2c2Shift = 22
	RegisterApb1lfz1FieldDbg_i2c2Mask  = 0x400000
)

// GetDbg_i2c2 I2C2 SMBUS timeout stop in debug
func (r *registerApb1lfz1Type) GetDbg_i2c2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_i2c2Mask) != 0
}

// SetDbg_i2c2 I2C2 SMBUS timeout stop in debug
func (r *registerApb1lfz1Type) SetDbg_i2c2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_i2c2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_i2c2Mask)
	}
}

const (
	RegisterApb1lfz1FieldDbg_i2c3Shift = 23
	RegisterApb1lfz1FieldDbg_i2c3Mask  = 0x800000
)

// GetDbg_i2c3 I2C3 SMBUS timeout stop in debug
func (r *registerApb1lfz1Type) GetDbg_i2c3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz1FieldDbg_i2c3Mask) != 0
}

// SetDbg_i2c3 I2C3 SMBUS timeout stop in debug
func (r *registerApb1lfz1Type) SetDbg_i2c3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz1FieldDbg_i2c3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz1FieldDbg_i2c3Mask)
	}
}

// registerApb1lfz2Type DBGMCU APB1L peripheral freeze register CPU2
type registerApb1lfz2Type uint32

const (
	RegisterApb1lfz2FieldDbg_tim2Shift = 0
	RegisterApb1lfz2FieldDbg_tim2Mask  = 0x1
)

// GetDbg_tim2 TIM2 stop in debug
func (r *registerApb1lfz2Type) GetDbg_tim2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_tim2Mask) != 0
}

// SetDbg_tim2 TIM2 stop in debug
func (r *registerApb1lfz2Type) SetDbg_tim2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_tim2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_tim2Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_tim3Shift = 1
	RegisterApb1lfz2FieldDbg_tim3Mask  = 0x2
)

// GetDbg_tim3 TIM3 stop in debug
func (r *registerApb1lfz2Type) GetDbg_tim3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_tim3Mask) != 0
}

// SetDbg_tim3 TIM3 stop in debug
func (r *registerApb1lfz2Type) SetDbg_tim3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_tim3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_tim3Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_tim4Shift = 2
	RegisterApb1lfz2FieldDbg_tim4Mask  = 0x4
)

// GetDbg_tim4 TIM4 stop in debug
func (r *registerApb1lfz2Type) GetDbg_tim4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_tim4Mask) != 0
}

// SetDbg_tim4 TIM4 stop in debug
func (r *registerApb1lfz2Type) SetDbg_tim4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_tim4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_tim4Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_tim5Shift = 3
	RegisterApb1lfz2FieldDbg_tim5Mask  = 0x8
)

// GetDbg_tim5 TIM5 stop in debug
func (r *registerApb1lfz2Type) GetDbg_tim5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_tim5Mask) != 0
}

// SetDbg_tim5 TIM5 stop in debug
func (r *registerApb1lfz2Type) SetDbg_tim5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_tim5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_tim5Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_tim6Shift = 4
	RegisterApb1lfz2FieldDbg_tim6Mask  = 0x10
)

// GetDbg_tim6 TIM6 stop in debug
func (r *registerApb1lfz2Type) GetDbg_tim6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_tim6Mask) != 0
}

// SetDbg_tim6 TIM6 stop in debug
func (r *registerApb1lfz2Type) SetDbg_tim6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_tim6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_tim6Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_tim7Shift = 5
	RegisterApb1lfz2FieldDbg_tim7Mask  = 0x20
)

// GetDbg_tim7 TIM4 stop in debug
func (r *registerApb1lfz2Type) GetDbg_tim7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_tim7Mask) != 0
}

// SetDbg_tim7 TIM4 stop in debug
func (r *registerApb1lfz2Type) SetDbg_tim7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_tim7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_tim7Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_tim12Shift = 6
	RegisterApb1lfz2FieldDbg_tim12Mask  = 0x40
)

// GetDbg_tim12 TIM12 stop in debug
func (r *registerApb1lfz2Type) GetDbg_tim12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_tim12Mask) != 0
}

// SetDbg_tim12 TIM12 stop in debug
func (r *registerApb1lfz2Type) SetDbg_tim12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_tim12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_tim12Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_tim13Shift = 7
	RegisterApb1lfz2FieldDbg_tim13Mask  = 0x80
)

// GetDbg_tim13 TIM13 stop in debug
func (r *registerApb1lfz2Type) GetDbg_tim13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_tim13Mask) != 0
}

// SetDbg_tim13 TIM13 stop in debug
func (r *registerApb1lfz2Type) SetDbg_tim13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_tim13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_tim13Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_tim14Shift = 8
	RegisterApb1lfz2FieldDbg_tim14Mask  = 0x100
)

// GetDbg_tim14 TIM14 stop in debug
func (r *registerApb1lfz2Type) GetDbg_tim14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_tim14Mask) != 0
}

// SetDbg_tim14 TIM14 stop in debug
func (r *registerApb1lfz2Type) SetDbg_tim14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_tim14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_tim14Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_lptim1Shift = 9
	RegisterApb1lfz2FieldDbg_lptim1Mask  = 0x200
)

// GetDbg_lptim1 LPTIM1 stop in debug
func (r *registerApb1lfz2Type) GetDbg_lptim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_lptim1Mask) != 0
}

// SetDbg_lptim1 LPTIM1 stop in debug
func (r *registerApb1lfz2Type) SetDbg_lptim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_lptim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_lptim1Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_wwdg2Shift = 11
	RegisterApb1lfz2FieldDbg_wwdg2Mask  = 0x800
)

// GetDbg_wwdg2 WWDG2 stop in debug
func (r *registerApb1lfz2Type) GetDbg_wwdg2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_wwdg2Mask) != 0
}

// SetDbg_wwdg2 WWDG2 stop in debug
func (r *registerApb1lfz2Type) SetDbg_wwdg2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_wwdg2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_wwdg2Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_i2c1Shift = 21
	RegisterApb1lfz2FieldDbg_i2c1Mask  = 0x200000
)

// GetDbg_i2c1 I2C1 SMBUS timeout stop in debug
func (r *registerApb1lfz2Type) GetDbg_i2c1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_i2c1Mask) != 0
}

// SetDbg_i2c1 I2C1 SMBUS timeout stop in debug
func (r *registerApb1lfz2Type) SetDbg_i2c1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_i2c1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_i2c1Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_i2c2Shift = 22
	RegisterApb1lfz2FieldDbg_i2c2Mask  = 0x400000
)

// GetDbg_i2c2 I2C2 SMBUS timeout stop in debug
func (r *registerApb1lfz2Type) GetDbg_i2c2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_i2c2Mask) != 0
}

// SetDbg_i2c2 I2C2 SMBUS timeout stop in debug
func (r *registerApb1lfz2Type) SetDbg_i2c2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_i2c2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_i2c2Mask)
	}
}

const (
	RegisterApb1lfz2FieldDbg_i2c3Shift = 23
	RegisterApb1lfz2FieldDbg_i2c3Mask  = 0x800000
)

// GetDbg_i2c3 I2C3 SMBUS timeout stop in debug
func (r *registerApb1lfz2Type) GetDbg_i2c3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lfz2FieldDbg_i2c3Mask) != 0
}

// SetDbg_i2c3 I2C3 SMBUS timeout stop in debug
func (r *registerApb1lfz2Type) SetDbg_i2c3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lfz2FieldDbg_i2c3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lfz2FieldDbg_i2c3Mask)
	}
}

// registerApb2fz1Type DBGMCU APB2 peripheral freeze register
type registerApb2fz1Type uint32

const (
	RegisterApb2fz1FieldDbg_tim1Shift = 0
	RegisterApb2fz1FieldDbg_tim1Mask  = 0x1
)

// GetDbg_tim1 TIM1 stop in debug
func (r *registerApb2fz1Type) GetDbg_tim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz1FieldDbg_tim1Mask) != 0
}

// SetDbg_tim1 TIM1 stop in debug
func (r *registerApb2fz1Type) SetDbg_tim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz1FieldDbg_tim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz1FieldDbg_tim1Mask)
	}
}

const (
	RegisterApb2fz1FieldDbg_tim8Shift = 1
	RegisterApb2fz1FieldDbg_tim8Mask  = 0x2
)

// GetDbg_tim8 TIM8 stop in debug
func (r *registerApb2fz1Type) GetDbg_tim8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz1FieldDbg_tim8Mask) != 0
}

// SetDbg_tim8 TIM8 stop in debug
func (r *registerApb2fz1Type) SetDbg_tim8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz1FieldDbg_tim8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz1FieldDbg_tim8Mask)
	}
}

const (
	RegisterApb2fz1FieldDbg_tim15Shift = 16
	RegisterApb2fz1FieldDbg_tim15Mask  = 0x10000
)

// GetDbg_tim15 TIM15 stop in debug
func (r *registerApb2fz1Type) GetDbg_tim15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz1FieldDbg_tim15Mask) != 0
}

// SetDbg_tim15 TIM15 stop in debug
func (r *registerApb2fz1Type) SetDbg_tim15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz1FieldDbg_tim15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz1FieldDbg_tim15Mask)
	}
}

const (
	RegisterApb2fz1FieldDbg_tim16Shift = 17
	RegisterApb2fz1FieldDbg_tim16Mask  = 0x20000
)

// GetDbg_tim16 TIM16 stop in debug
func (r *registerApb2fz1Type) GetDbg_tim16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz1FieldDbg_tim16Mask) != 0
}

// SetDbg_tim16 TIM16 stop in debug
func (r *registerApb2fz1Type) SetDbg_tim16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz1FieldDbg_tim16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz1FieldDbg_tim16Mask)
	}
}

const (
	RegisterApb2fz1FieldDbg_tim17Shift = 18
	RegisterApb2fz1FieldDbg_tim17Mask  = 0x40000
)

// GetDbg_tim17 TIM17 stop in debug
func (r *registerApb2fz1Type) GetDbg_tim17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz1FieldDbg_tim17Mask) != 0
}

// SetDbg_tim17 TIM17 stop in debug
func (r *registerApb2fz1Type) SetDbg_tim17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz1FieldDbg_tim17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz1FieldDbg_tim17Mask)
	}
}

const (
	RegisterApb2fz1FieldDbg_hrtimShift = 29
	RegisterApb2fz1FieldDbg_hrtimMask  = 0x20000000
)

// GetDbg_hrtim HRTIM stop in debug
func (r *registerApb2fz1Type) GetDbg_hrtim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz1FieldDbg_hrtimMask) != 0
}

// SetDbg_hrtim HRTIM stop in debug
func (r *registerApb2fz1Type) SetDbg_hrtim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz1FieldDbg_hrtimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz1FieldDbg_hrtimMask)
	}
}

// registerApb2fz2Type DBGMCU APB2 peripheral freeze register CPU2
type registerApb2fz2Type uint32

const (
	RegisterApb2fz2FieldDbg_tim1Shift = 0
	RegisterApb2fz2FieldDbg_tim1Mask  = 0x1
)

// GetDbg_tim1 TIM1 stop in debug
func (r *registerApb2fz2Type) GetDbg_tim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz2FieldDbg_tim1Mask) != 0
}

// SetDbg_tim1 TIM1 stop in debug
func (r *registerApb2fz2Type) SetDbg_tim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz2FieldDbg_tim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz2FieldDbg_tim1Mask)
	}
}

const (
	RegisterApb2fz2FieldDbg_tim8Shift = 1
	RegisterApb2fz2FieldDbg_tim8Mask  = 0x2
)

// GetDbg_tim8 TIM8 stop in debug
func (r *registerApb2fz2Type) GetDbg_tim8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz2FieldDbg_tim8Mask) != 0
}

// SetDbg_tim8 TIM8 stop in debug
func (r *registerApb2fz2Type) SetDbg_tim8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz2FieldDbg_tim8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz2FieldDbg_tim8Mask)
	}
}

const (
	RegisterApb2fz2FieldDbg_tim15Shift = 16
	RegisterApb2fz2FieldDbg_tim15Mask  = 0x10000
)

// GetDbg_tim15 TIM15 stop in debug
func (r *registerApb2fz2Type) GetDbg_tim15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz2FieldDbg_tim15Mask) != 0
}

// SetDbg_tim15 TIM15 stop in debug
func (r *registerApb2fz2Type) SetDbg_tim15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz2FieldDbg_tim15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz2FieldDbg_tim15Mask)
	}
}

const (
	RegisterApb2fz2FieldDbg_tim16Shift = 17
	RegisterApb2fz2FieldDbg_tim16Mask  = 0x20000
)

// GetDbg_tim16 TIM16 stop in debug
func (r *registerApb2fz2Type) GetDbg_tim16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz2FieldDbg_tim16Mask) != 0
}

// SetDbg_tim16 TIM16 stop in debug
func (r *registerApb2fz2Type) SetDbg_tim16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz2FieldDbg_tim16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz2FieldDbg_tim16Mask)
	}
}

const (
	RegisterApb2fz2FieldDbg_tim17Shift = 18
	RegisterApb2fz2FieldDbg_tim17Mask  = 0x40000
)

// GetDbg_tim17 TIM17 stop in debug
func (r *registerApb2fz2Type) GetDbg_tim17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz2FieldDbg_tim17Mask) != 0
}

// SetDbg_tim17 TIM17 stop in debug
func (r *registerApb2fz2Type) SetDbg_tim17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz2FieldDbg_tim17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz2FieldDbg_tim17Mask)
	}
}

const (
	RegisterApb2fz2FieldDbg_hrtimShift = 29
	RegisterApb2fz2FieldDbg_hrtimMask  = 0x20000000
)

// GetDbg_hrtim HRTIM stop in debug
func (r *registerApb2fz2Type) GetDbg_hrtim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2fz2FieldDbg_hrtimMask) != 0
}

// SetDbg_hrtim HRTIM stop in debug
func (r *registerApb2fz2Type) SetDbg_hrtim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2fz2FieldDbg_hrtimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2fz2FieldDbg_hrtimMask)
	}
}

// registerApb4fz1Type DBGMCU APB4 peripheral freeze register
type registerApb4fz1Type uint32

const (
	RegisterApb4fz1FieldDbg_i2c4Shift = 7
	RegisterApb4fz1FieldDbg_i2c4Mask  = 0x80
)

// GetDbg_i2c4 I2C4 SMBUS timeout stop in debug
func (r *registerApb4fz1Type) GetDbg_i2c4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbg_i2c4Mask) != 0
}

// SetDbg_i2c4 I2C4 SMBUS timeout stop in debug
func (r *registerApb4fz1Type) SetDbg_i2c4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbg_i2c4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbg_i2c4Mask)
	}
}

const (
	RegisterApb4fz1FieldDbg_lptim2Shift = 9
	RegisterApb4fz1FieldDbg_lptim2Mask  = 0x200
)

// GetDbg_lptim2 LPTIM2 stop in debug
func (r *registerApb4fz1Type) GetDbg_lptim2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbg_lptim2Mask) != 0
}

// SetDbg_lptim2 LPTIM2 stop in debug
func (r *registerApb4fz1Type) SetDbg_lptim2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbg_lptim2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbg_lptim2Mask)
	}
}

const (
	RegisterApb4fz1FieldDbg_lptim3Shift = 10
	RegisterApb4fz1FieldDbg_lptim3Mask  = 0x400
)

// GetDbg_lptim3 LPTIM2 stop in debug
func (r *registerApb4fz1Type) GetDbg_lptim3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbg_lptim3Mask) != 0
}

// SetDbg_lptim3 LPTIM2 stop in debug
func (r *registerApb4fz1Type) SetDbg_lptim3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbg_lptim3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbg_lptim3Mask)
	}
}

const (
	RegisterApb4fz1FieldDbg_lptim4Shift = 11
	RegisterApb4fz1FieldDbg_lptim4Mask  = 0x800
)

// GetDbg_lptim4 LPTIM4 stop in debug
func (r *registerApb4fz1Type) GetDbg_lptim4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbg_lptim4Mask) != 0
}

// SetDbg_lptim4 LPTIM4 stop in debug
func (r *registerApb4fz1Type) SetDbg_lptim4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbg_lptim4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbg_lptim4Mask)
	}
}

const (
	RegisterApb4fz1FieldDbg_lptim5Shift = 12
	RegisterApb4fz1FieldDbg_lptim5Mask  = 0x1000
)

// GetDbg_lptim5 LPTIM5 stop in debug
func (r *registerApb4fz1Type) GetDbg_lptim5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbg_lptim5Mask) != 0
}

// SetDbg_lptim5 LPTIM5 stop in debug
func (r *registerApb4fz1Type) SetDbg_lptim5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbg_lptim5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbg_lptim5Mask)
	}
}

const (
	RegisterApb4fz1FieldDbg_rtcShift = 16
	RegisterApb4fz1FieldDbg_rtcMask  = 0x10000
)

// GetDbg_rtc RTC stop in debug
func (r *registerApb4fz1Type) GetDbg_rtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbg_rtcMask) != 0
}

// SetDbg_rtc RTC stop in debug
func (r *registerApb4fz1Type) SetDbg_rtc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbg_rtcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbg_rtcMask)
	}
}

const (
	RegisterApb4fz1FieldDbg_wdglsd1Shift = 18
	RegisterApb4fz1FieldDbg_wdglsd1Mask  = 0x40000
)

// GetDbg_wdglsd1 Independent watchdog for D1 stop in debug
func (r *registerApb4fz1Type) GetDbg_wdglsd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbg_wdglsd1Mask) != 0
}

// SetDbg_wdglsd1 Independent watchdog for D1 stop in debug
func (r *registerApb4fz1Type) SetDbg_wdglsd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbg_wdglsd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbg_wdglsd1Mask)
	}
}

const (
	RegisterApb4fz1FieldDbg_wdglsd2Shift = 19
	RegisterApb4fz1FieldDbg_wdglsd2Mask  = 0x80000
)

// GetDbg_wdglsd2 Independent watchdog for D2 stop in debug
func (r *registerApb4fz1Type) GetDbg_wdglsd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz1FieldDbg_wdglsd2Mask) != 0
}

// SetDbg_wdglsd2 Independent watchdog for D2 stop in debug
func (r *registerApb4fz1Type) SetDbg_wdglsd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz1FieldDbg_wdglsd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz1FieldDbg_wdglsd2Mask)
	}
}

// registerApb4fz2Type DBGMCU APB4 peripheral freeze register CPU2
type registerApb4fz2Type uint32

const (
	RegisterApb4fz2FieldDbg_i2c4Shift = 7
	RegisterApb4fz2FieldDbg_i2c4Mask  = 0x80
)

// GetDbg_i2c4 I2C4 SMBUS timeout stop in debug
func (r *registerApb4fz2Type) GetDbg_i2c4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbg_i2c4Mask) != 0
}

// SetDbg_i2c4 I2C4 SMBUS timeout stop in debug
func (r *registerApb4fz2Type) SetDbg_i2c4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbg_i2c4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbg_i2c4Mask)
	}
}

const (
	RegisterApb4fz2FieldDbg_lptim2Shift = 9
	RegisterApb4fz2FieldDbg_lptim2Mask  = 0x200
)

// GetDbg_lptim2 LPTIM2 stop in debug
func (r *registerApb4fz2Type) GetDbg_lptim2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbg_lptim2Mask) != 0
}

// SetDbg_lptim2 LPTIM2 stop in debug
func (r *registerApb4fz2Type) SetDbg_lptim2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbg_lptim2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbg_lptim2Mask)
	}
}

const (
	RegisterApb4fz2FieldDbg_lptim3Shift = 10
	RegisterApb4fz2FieldDbg_lptim3Mask  = 0x400
)

// GetDbg_lptim3 LPTIM2 stop in debug
func (r *registerApb4fz2Type) GetDbg_lptim3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbg_lptim3Mask) != 0
}

// SetDbg_lptim3 LPTIM2 stop in debug
func (r *registerApb4fz2Type) SetDbg_lptim3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbg_lptim3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbg_lptim3Mask)
	}
}

const (
	RegisterApb4fz2FieldDbg_lptim4Shift = 11
	RegisterApb4fz2FieldDbg_lptim4Mask  = 0x800
)

// GetDbg_lptim4 LPTIM4 stop in debug
func (r *registerApb4fz2Type) GetDbg_lptim4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbg_lptim4Mask) != 0
}

// SetDbg_lptim4 LPTIM4 stop in debug
func (r *registerApb4fz2Type) SetDbg_lptim4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbg_lptim4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbg_lptim4Mask)
	}
}

const (
	RegisterApb4fz2FieldDbg_lptim5Shift = 12
	RegisterApb4fz2FieldDbg_lptim5Mask  = 0x1000
)

// GetDbg_lptim5 LPTIM5 stop in debug
func (r *registerApb4fz2Type) GetDbg_lptim5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbg_lptim5Mask) != 0
}

// SetDbg_lptim5 LPTIM5 stop in debug
func (r *registerApb4fz2Type) SetDbg_lptim5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbg_lptim5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbg_lptim5Mask)
	}
}

const (
	RegisterApb4fz2FieldDbg_rtcShift = 16
	RegisterApb4fz2FieldDbg_rtcMask  = 0x10000
)

// GetDbg_rtc RTC stop in debug
func (r *registerApb4fz2Type) GetDbg_rtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbg_rtcMask) != 0
}

// SetDbg_rtc RTC stop in debug
func (r *registerApb4fz2Type) SetDbg_rtc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbg_rtcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbg_rtcMask)
	}
}

const (
	RegisterApb4fz2FieldDbg_wdglsd1Shift = 18
	RegisterApb4fz2FieldDbg_wdglsd1Mask  = 0x40000
)

// GetDbg_wdglsd1 LS watchdog for D1 stop in debug
func (r *registerApb4fz2Type) GetDbg_wdglsd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbg_wdglsd1Mask) != 0
}

// SetDbg_wdglsd1 LS watchdog for D1 stop in debug
func (r *registerApb4fz2Type) SetDbg_wdglsd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbg_wdglsd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbg_wdglsd1Mask)
	}
}

const (
	RegisterApb4fz2FieldDbg_wdglsd2Shift = 19
	RegisterApb4fz2FieldDbg_wdglsd2Mask  = 0x80000
)

// GetDbg_wdglsd2 LS watchdog for D2 stop in debug
func (r *registerApb4fz2Type) GetDbg_wdglsd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4fz2FieldDbg_wdglsd2Mask) != 0
}

// SetDbg_wdglsd2 LS watchdog for D2 stop in debug
func (r *registerApb4fz2Type) SetDbg_wdglsd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4fz2FieldDbg_wdglsd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4fz2FieldDbg_wdglsd2Mask)
	}
}
