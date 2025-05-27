//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package opamp

import (
	"unsafe"
	"volatile"
)

var (
	Opamp = (*_opamp)(unsafe.Pointer(uintptr(0x40009000)))
)

type _opamp struct {
	Opamp1csr   registerOpamp1csrType
	Opamp1otr   registerOpamp1otrType
	Opamp1hsotr registerOpamp1hsotrType
	_           [4]byte
	Opamp2csr   registerOpamp2csrType
	Opamp2otr   registerOpamp2otrType
	Opamp2hsotr registerOpamp2hsotrType
}

// registerOpamp1csrType OPAMP1 control/status register
type registerOpamp1csrType uint32

const (
	RegisterOpamp1csrFieldOpaenShift = 0
	RegisterOpamp1csrFieldOpaenMask  = 0x1
)

// GetOpaen Operational amplifier Enable
func (r *registerOpamp1csrType) GetOpaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldOpaenMask) != 0
}

// SetOpaen Operational amplifier Enable
func (r *registerOpamp1csrType) SetOpaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1csrFieldOpaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldOpaenMask)
	}
}

const (
	RegisterOpamp1csrFieldForcevpShift = 1
	RegisterOpamp1csrFieldForcevpMask  = 0x2
)

// GetForcevp Force internal reference on VP (reserved for test
func (r *registerOpamp1csrType) GetForcevp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldForcevpMask) != 0
}

// SetForcevp Force internal reference on VP (reserved for test
func (r *registerOpamp1csrType) SetForcevp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1csrFieldForcevpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldForcevpMask)
	}
}

const (
	RegisterOpamp1csrFieldVpselShift = 2
	RegisterOpamp1csrFieldVpselMask  = 0xc
)

// GetVpsel Operational amplifier PGA mode
func (r *registerOpamp1csrType) GetVpsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldVpselMask) >> RegisterOpamp1csrFieldVpselShift)
}

// SetVpsel Operational amplifier PGA mode
func (r *registerOpamp1csrType) SetVpsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldVpselMask)|(uint32(value)<<RegisterOpamp1csrFieldVpselShift))
}

const (
	RegisterOpamp1csrFieldVmselShift = 5
	RegisterOpamp1csrFieldVmselMask  = 0x60
)

// GetVmsel Inverting input selection
func (r *registerOpamp1csrType) GetVmsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldVmselMask) >> RegisterOpamp1csrFieldVmselShift)
}

// SetVmsel Inverting input selection
func (r *registerOpamp1csrType) SetVmsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldVmselMask)|(uint32(value)<<RegisterOpamp1csrFieldVmselShift))
}

const (
	RegisterOpamp1csrFieldOpahsmShift = 8
	RegisterOpamp1csrFieldOpahsmMask  = 0x100
)

// GetOpahsm Operational amplifier high-speed mode
func (r *registerOpamp1csrType) GetOpahsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldOpahsmMask) != 0
}

// SetOpahsm Operational amplifier high-speed mode
func (r *registerOpamp1csrType) SetOpahsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1csrFieldOpahsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldOpahsmMask)
	}
}

const (
	RegisterOpamp1csrFieldCalonShift = 11
	RegisterOpamp1csrFieldCalonMask  = 0x800
)

// GetCalon Calibration mode enabled
func (r *registerOpamp1csrType) GetCalon() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldCalonMask) != 0
}

// SetCalon Calibration mode enabled
func (r *registerOpamp1csrType) SetCalon(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1csrFieldCalonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldCalonMask)
	}
}

const (
	RegisterOpamp1csrFieldCalselShift = 12
	RegisterOpamp1csrFieldCalselMask  = 0x3000
)

// GetCalsel Calibration selection
func (r *registerOpamp1csrType) GetCalsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldCalselMask) >> RegisterOpamp1csrFieldCalselShift)
}

// SetCalsel Calibration selection
func (r *registerOpamp1csrType) SetCalsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldCalselMask)|(uint32(value)<<RegisterOpamp1csrFieldCalselShift))
}

const (
	RegisterOpamp1csrFieldPgagainShift = 14
	RegisterOpamp1csrFieldPgagainMask  = 0x3c000
)

// GetPgagain allows to switch from AOP offset trimmed values to AOP offset
func (r *registerOpamp1csrType) GetPgagain() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldPgagainMask) >> RegisterOpamp1csrFieldPgagainShift)
}

// SetPgagain allows to switch from AOP offset trimmed values to AOP offset
func (r *registerOpamp1csrType) SetPgagain(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldPgagainMask)|(uint32(value)<<RegisterOpamp1csrFieldPgagainShift))
}

const (
	RegisterOpamp1csrFieldUsertrimShift = 18
	RegisterOpamp1csrFieldUsertrimMask  = 0x40000
)

// GetUsertrim User trimming enable
func (r *registerOpamp1csrType) GetUsertrim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldUsertrimMask) != 0
}

// SetUsertrim User trimming enable
func (r *registerOpamp1csrType) SetUsertrim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1csrFieldUsertrimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldUsertrimMask)
	}
}

const (
	RegisterOpamp1csrFieldTstrefShift = 29
	RegisterOpamp1csrFieldTstrefMask  = 0x20000000
)

// GetTstref OPAMP calibration reference voltage output control (reserved for test)
func (r *registerOpamp1csrType) GetTstref() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldTstrefMask) != 0
}

// SetTstref OPAMP calibration reference voltage output control (reserved for test)
func (r *registerOpamp1csrType) SetTstref(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1csrFieldTstrefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldTstrefMask)
	}
}

const (
	RegisterOpamp1csrFieldCaloutShift = 30
	RegisterOpamp1csrFieldCaloutMask  = 0x40000000
)

// GetCalout Operational amplifier calibration output
func (r *registerOpamp1csrType) GetCalout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldCaloutMask) != 0
}

// SetCalout Operational amplifier calibration output
func (r *registerOpamp1csrType) SetCalout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1csrFieldCaloutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldCaloutMask)
	}
}

// registerOpamp1otrType OPAMP1 offset trimming register in normal mode
type registerOpamp1otrType uint32

const (
	RegisterOpamp1otrFieldTrimoffsetnShift = 0
	RegisterOpamp1otrFieldTrimoffsetnMask  = 0x1f
)

// GetTrimoffsetn Trim for NMOS differential pairs
func (r *registerOpamp1otrType) GetTrimoffsetn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1otrFieldTrimoffsetnMask) >> RegisterOpamp1otrFieldTrimoffsetnShift)
}

// SetTrimoffsetn Trim for NMOS differential pairs
func (r *registerOpamp1otrType) SetTrimoffsetn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1otrFieldTrimoffsetnMask)|(uint32(value)<<RegisterOpamp1otrFieldTrimoffsetnShift))
}

const (
	RegisterOpamp1otrFieldTrimoffsetpShift = 8
	RegisterOpamp1otrFieldTrimoffsetpMask  = 0x1f00
)

// GetTrimoffsetp Trim for PMOS differential pairs
func (r *registerOpamp1otrType) GetTrimoffsetp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1otrFieldTrimoffsetpMask) >> RegisterOpamp1otrFieldTrimoffsetpShift)
}

// SetTrimoffsetp Trim for PMOS differential pairs
func (r *registerOpamp1otrType) SetTrimoffsetp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1otrFieldTrimoffsetpMask)|(uint32(value)<<RegisterOpamp1otrFieldTrimoffsetpShift))
}

// registerOpamp1hsotrType OPAMP1 offset trimming register in low-power mode
type registerOpamp1hsotrType uint32

const (
	RegisterOpamp1hsotrFieldTrimlpoffsetnShift = 0
	RegisterOpamp1hsotrFieldTrimlpoffsetnMask  = 0x1f
)

// GetTrimlpoffsetn Trim for NMOS differential pairs
func (r *registerOpamp1hsotrType) GetTrimlpoffsetn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1hsotrFieldTrimlpoffsetnMask) >> RegisterOpamp1hsotrFieldTrimlpoffsetnShift)
}

// SetTrimlpoffsetn Trim for NMOS differential pairs
func (r *registerOpamp1hsotrType) SetTrimlpoffsetn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1hsotrFieldTrimlpoffsetnMask)|(uint32(value)<<RegisterOpamp1hsotrFieldTrimlpoffsetnShift))
}

const (
	RegisterOpamp1hsotrFieldTrimlpoffsetpShift = 8
	RegisterOpamp1hsotrFieldTrimlpoffsetpMask  = 0x1f00
)

// GetTrimlpoffsetp Trim for PMOS differential pairs
func (r *registerOpamp1hsotrType) GetTrimlpoffsetp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1hsotrFieldTrimlpoffsetpMask) >> RegisterOpamp1hsotrFieldTrimlpoffsetpShift)
}

// SetTrimlpoffsetp Trim for PMOS differential pairs
func (r *registerOpamp1hsotrType) SetTrimlpoffsetp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1hsotrFieldTrimlpoffsetpMask)|(uint32(value)<<RegisterOpamp1hsotrFieldTrimlpoffsetpShift))
}

// registerOpamp2csrType OPAMP2 control/status register
type registerOpamp2csrType uint32

const (
	RegisterOpamp2csrFieldOpaenShift = 0
	RegisterOpamp2csrFieldOpaenMask  = 0x1
)

// GetOpaen Operational amplifier Enable
func (r *registerOpamp2csrType) GetOpaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldOpaenMask) != 0
}

// SetOpaen Operational amplifier Enable
func (r *registerOpamp2csrType) SetOpaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2csrFieldOpaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldOpaenMask)
	}
}

const (
	RegisterOpamp2csrFieldForcevpShift = 1
	RegisterOpamp2csrFieldForcevpMask  = 0x2
)

// GetForcevp Force internal reference on VP (reserved for test)
func (r *registerOpamp2csrType) GetForcevp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldForcevpMask) != 0
}

// SetForcevp Force internal reference on VP (reserved for test)
func (r *registerOpamp2csrType) SetForcevp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2csrFieldForcevpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldForcevpMask)
	}
}

const (
	RegisterOpamp2csrFieldVmselShift = 5
	RegisterOpamp2csrFieldVmselMask  = 0x60
)

// GetVmsel Inverting input selection
func (r *registerOpamp2csrType) GetVmsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldVmselMask) >> RegisterOpamp2csrFieldVmselShift)
}

// SetVmsel Inverting input selection
func (r *registerOpamp2csrType) SetVmsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldVmselMask)|(uint32(value)<<RegisterOpamp2csrFieldVmselShift))
}

const (
	RegisterOpamp2csrFieldOpahsmShift = 8
	RegisterOpamp2csrFieldOpahsmMask  = 0x100
)

// GetOpahsm Operational amplifier high-speed mode
func (r *registerOpamp2csrType) GetOpahsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldOpahsmMask) != 0
}

// SetOpahsm Operational amplifier high-speed mode
func (r *registerOpamp2csrType) SetOpahsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2csrFieldOpahsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldOpahsmMask)
	}
}

const (
	RegisterOpamp2csrFieldCalonShift = 11
	RegisterOpamp2csrFieldCalonMask  = 0x800
)

// GetCalon Calibration mode enabled
func (r *registerOpamp2csrType) GetCalon() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldCalonMask) != 0
}

// SetCalon Calibration mode enabled
func (r *registerOpamp2csrType) SetCalon(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2csrFieldCalonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldCalonMask)
	}
}

const (
	RegisterOpamp2csrFieldCalselShift = 12
	RegisterOpamp2csrFieldCalselMask  = 0x3000
)

// GetCalsel Calibration selection
func (r *registerOpamp2csrType) GetCalsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldCalselMask) >> RegisterOpamp2csrFieldCalselShift)
}

// SetCalsel Calibration selection
func (r *registerOpamp2csrType) SetCalsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldCalselMask)|(uint32(value)<<RegisterOpamp2csrFieldCalselShift))
}

const (
	RegisterOpamp2csrFieldPgagainShift = 14
	RegisterOpamp2csrFieldPgagainMask  = 0x3c000
)

// GetPgagain Operational amplifier Programmable amplifier gain value
func (r *registerOpamp2csrType) GetPgagain() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldPgagainMask) >> RegisterOpamp2csrFieldPgagainShift)
}

// SetPgagain Operational amplifier Programmable amplifier gain value
func (r *registerOpamp2csrType) SetPgagain(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldPgagainMask)|(uint32(value)<<RegisterOpamp2csrFieldPgagainShift))
}

const (
	RegisterOpamp2csrFieldUsertrimShift = 18
	RegisterOpamp2csrFieldUsertrimMask  = 0x40000
)

// GetUsertrim User trimming enable
func (r *registerOpamp2csrType) GetUsertrim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldUsertrimMask) != 0
}

// SetUsertrim User trimming enable
func (r *registerOpamp2csrType) SetUsertrim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2csrFieldUsertrimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldUsertrimMask)
	}
}

const (
	RegisterOpamp2csrFieldTstrefShift = 29
	RegisterOpamp2csrFieldTstrefMask  = 0x20000000
)

// GetTstref OPAMP calibration reference voltage output control (reserved for test)
func (r *registerOpamp2csrType) GetTstref() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldTstrefMask) != 0
}

// SetTstref OPAMP calibration reference voltage output control (reserved for test)
func (r *registerOpamp2csrType) SetTstref(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2csrFieldTstrefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldTstrefMask)
	}
}

const (
	RegisterOpamp2csrFieldCaloutShift = 30
	RegisterOpamp2csrFieldCaloutMask  = 0x40000000
)

// GetCalout Operational amplifier calibration output
func (r *registerOpamp2csrType) GetCalout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldCaloutMask) != 0
}

// SetCalout Operational amplifier calibration output
func (r *registerOpamp2csrType) SetCalout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2csrFieldCaloutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldCaloutMask)
	}
}

// registerOpamp2otrType OPAMP2 offset trimming register in normal mode
type registerOpamp2otrType uint32

const (
	RegisterOpamp2otrFieldTrimoffsetnShift = 0
	RegisterOpamp2otrFieldTrimoffsetnMask  = 0x1f
)

// GetTrimoffsetn Trim for NMOS differential pairs
func (r *registerOpamp2otrType) GetTrimoffsetn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2otrFieldTrimoffsetnMask) >> RegisterOpamp2otrFieldTrimoffsetnShift)
}

// SetTrimoffsetn Trim for NMOS differential pairs
func (r *registerOpamp2otrType) SetTrimoffsetn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2otrFieldTrimoffsetnMask)|(uint32(value)<<RegisterOpamp2otrFieldTrimoffsetnShift))
}

const (
	RegisterOpamp2otrFieldTrimoffsetpShift = 8
	RegisterOpamp2otrFieldTrimoffsetpMask  = 0x1f00
)

// GetTrimoffsetp Trim for PMOS differential pairs
func (r *registerOpamp2otrType) GetTrimoffsetp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2otrFieldTrimoffsetpMask) >> RegisterOpamp2otrFieldTrimoffsetpShift)
}

// SetTrimoffsetp Trim for PMOS differential pairs
func (r *registerOpamp2otrType) SetTrimoffsetp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2otrFieldTrimoffsetpMask)|(uint32(value)<<RegisterOpamp2otrFieldTrimoffsetpShift))
}

// registerOpamp2hsotrType OPAMP2 offset trimming register in low-power mode
type registerOpamp2hsotrType uint32

const (
	RegisterOpamp2hsotrFieldTrimlpoffsetnShift = 0
	RegisterOpamp2hsotrFieldTrimlpoffsetnMask  = 0x1f
)

// GetTrimlpoffsetn Trim for NMOS differential pairs
func (r *registerOpamp2hsotrType) GetTrimlpoffsetn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2hsotrFieldTrimlpoffsetnMask) >> RegisterOpamp2hsotrFieldTrimlpoffsetnShift)
}

// SetTrimlpoffsetn Trim for NMOS differential pairs
func (r *registerOpamp2hsotrType) SetTrimlpoffsetn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2hsotrFieldTrimlpoffsetnMask)|(uint32(value)<<RegisterOpamp2hsotrFieldTrimlpoffsetnShift))
}

const (
	RegisterOpamp2hsotrFieldTrimlpoffsetpShift = 8
	RegisterOpamp2hsotrFieldTrimlpoffsetpMask  = 0x1f00
)

// GetTrimlpoffsetp Trim for PMOS differential pairs
func (r *registerOpamp2hsotrType) GetTrimlpoffsetp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2hsotrFieldTrimlpoffsetpMask) >> RegisterOpamp2hsotrFieldTrimlpoffsetpShift)
}

// SetTrimlpoffsetp Trim for PMOS differential pairs
func (r *registerOpamp2hsotrType) SetTrimlpoffsetp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2hsotrFieldTrimlpoffsetpMask)|(uint32(value)<<RegisterOpamp2hsotrFieldTrimlpoffsetpShift))
}
