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
	Opamp1_csr   registerOpamp1_csrType
	Opamp1_otr   registerOpamp1_otrType
	Opamp1_hsotr registerOpamp1_hsotrType
	_            [4]byte
	Opamp2_csr   registerOpamp2_csrType
	Opamp2_otr   registerOpamp2_otrType
	Opamp2_hsotr registerOpamp2_hsotrType
}

// registerOpamp1_csrType OPAMP1 control/status register
type registerOpamp1_csrType uint32

const (
	RegisterOpamp1_csrFieldOpaenShift = 0
	RegisterOpamp1_csrFieldOpaenMask  = 0x1
)

// GetOpaen Operational amplifier Enable
func (r *registerOpamp1_csrType) GetOpaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_csrFieldOpaenMask) != 0
}

// SetOpaen Operational amplifier Enable
func (r *registerOpamp1_csrType) SetOpaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1_csrFieldOpaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_csrFieldOpaenMask)
	}
}

const (
	RegisterOpamp1_csrFieldForce_vpShift = 1
	RegisterOpamp1_csrFieldForce_vpMask  = 0x2
)

// GetForce_vp Force internal reference on VP (reserved for test
func (r *registerOpamp1_csrType) GetForce_vp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_csrFieldForce_vpMask) != 0
}

// SetForce_vp Force internal reference on VP (reserved for test
func (r *registerOpamp1_csrType) SetForce_vp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1_csrFieldForce_vpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_csrFieldForce_vpMask)
	}
}

const (
	RegisterOpamp1_csrFieldVp_selShift = 2
	RegisterOpamp1_csrFieldVp_selMask  = 0xc
)

// GetVp_sel Operational amplifier PGA mode
func (r *registerOpamp1_csrType) GetVp_sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_csrFieldVp_selMask) >> RegisterOpamp1_csrFieldVp_selShift)
}

// SetVp_sel Operational amplifier PGA mode
func (r *registerOpamp1_csrType) SetVp_sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_csrFieldVp_selMask)|(uint32(value)<<RegisterOpamp1_csrFieldVp_selShift))
}

const (
	RegisterOpamp1_csrFieldVm_selShift = 5
	RegisterOpamp1_csrFieldVm_selMask  = 0x60
)

// GetVm_sel Inverting input selection
func (r *registerOpamp1_csrType) GetVm_sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_csrFieldVm_selMask) >> RegisterOpamp1_csrFieldVm_selShift)
}

// SetVm_sel Inverting input selection
func (r *registerOpamp1_csrType) SetVm_sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_csrFieldVm_selMask)|(uint32(value)<<RegisterOpamp1_csrFieldVm_selShift))
}

const (
	RegisterOpamp1_csrFieldOpahsmShift = 8
	RegisterOpamp1_csrFieldOpahsmMask  = 0x100
)

// GetOpahsm Operational amplifier high-speed mode
func (r *registerOpamp1_csrType) GetOpahsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_csrFieldOpahsmMask) != 0
}

// SetOpahsm Operational amplifier high-speed mode
func (r *registerOpamp1_csrType) SetOpahsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1_csrFieldOpahsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_csrFieldOpahsmMask)
	}
}

const (
	RegisterOpamp1_csrFieldCalonShift = 11
	RegisterOpamp1_csrFieldCalonMask  = 0x800
)

// GetCalon Calibration mode enabled
func (r *registerOpamp1_csrType) GetCalon() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_csrFieldCalonMask) != 0
}

// SetCalon Calibration mode enabled
func (r *registerOpamp1_csrType) SetCalon(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1_csrFieldCalonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_csrFieldCalonMask)
	}
}

const (
	RegisterOpamp1_csrFieldCalselShift = 12
	RegisterOpamp1_csrFieldCalselMask  = 0x3000
)

// GetCalsel Calibration selection
func (r *registerOpamp1_csrType) GetCalsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_csrFieldCalselMask) >> RegisterOpamp1_csrFieldCalselShift)
}

// SetCalsel Calibration selection
func (r *registerOpamp1_csrType) SetCalsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_csrFieldCalselMask)|(uint32(value)<<RegisterOpamp1_csrFieldCalselShift))
}

const (
	RegisterOpamp1_csrFieldPga_gainShift = 14
	RegisterOpamp1_csrFieldPga_gainMask  = 0x3c000
)

// GetPga_gain allows to switch from AOP offset trimmed values to AOP offset
func (r *registerOpamp1_csrType) GetPga_gain() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_csrFieldPga_gainMask) >> RegisterOpamp1_csrFieldPga_gainShift)
}

// SetPga_gain allows to switch from AOP offset trimmed values to AOP offset
func (r *registerOpamp1_csrType) SetPga_gain(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_csrFieldPga_gainMask)|(uint32(value)<<RegisterOpamp1_csrFieldPga_gainShift))
}

const (
	RegisterOpamp1_csrFieldUsertrimShift = 18
	RegisterOpamp1_csrFieldUsertrimMask  = 0x40000
)

// GetUsertrim User trimming enable
func (r *registerOpamp1_csrType) GetUsertrim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_csrFieldUsertrimMask) != 0
}

// SetUsertrim User trimming enable
func (r *registerOpamp1_csrType) SetUsertrim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1_csrFieldUsertrimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_csrFieldUsertrimMask)
	}
}

const (
	RegisterOpamp1_csrFieldTstrefShift = 29
	RegisterOpamp1_csrFieldTstrefMask  = 0x20000000
)

// GetTstref OPAMP calibration reference voltage output control (reserved for test)
func (r *registerOpamp1_csrType) GetTstref() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_csrFieldTstrefMask) != 0
}

// SetTstref OPAMP calibration reference voltage output control (reserved for test)
func (r *registerOpamp1_csrType) SetTstref(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1_csrFieldTstrefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_csrFieldTstrefMask)
	}
}

const (
	RegisterOpamp1_csrFieldCaloutShift = 30
	RegisterOpamp1_csrFieldCaloutMask  = 0x40000000
)

// GetCalout Operational amplifier calibration output
func (r *registerOpamp1_csrType) GetCalout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_csrFieldCaloutMask) != 0
}

// SetCalout Operational amplifier calibration output
func (r *registerOpamp1_csrType) SetCalout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1_csrFieldCaloutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_csrFieldCaloutMask)
	}
}

// registerOpamp1_otrType OPAMP1 offset trimming register in normal mode
type registerOpamp1_otrType uint32

const (
	RegisterOpamp1_otrFieldTrimoffsetnShift = 0
	RegisterOpamp1_otrFieldTrimoffsetnMask  = 0x1f
)

// GetTrimoffsetn Trim for NMOS differential pairs
func (r *registerOpamp1_otrType) GetTrimoffsetn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_otrFieldTrimoffsetnMask) >> RegisterOpamp1_otrFieldTrimoffsetnShift)
}

// SetTrimoffsetn Trim for NMOS differential pairs
func (r *registerOpamp1_otrType) SetTrimoffsetn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_otrFieldTrimoffsetnMask)|(uint32(value)<<RegisterOpamp1_otrFieldTrimoffsetnShift))
}

const (
	RegisterOpamp1_otrFieldTrimoffsetpShift = 8
	RegisterOpamp1_otrFieldTrimoffsetpMask  = 0x1f00
)

// GetTrimoffsetp Trim for PMOS differential pairs
func (r *registerOpamp1_otrType) GetTrimoffsetp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_otrFieldTrimoffsetpMask) >> RegisterOpamp1_otrFieldTrimoffsetpShift)
}

// SetTrimoffsetp Trim for PMOS differential pairs
func (r *registerOpamp1_otrType) SetTrimoffsetp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_otrFieldTrimoffsetpMask)|(uint32(value)<<RegisterOpamp1_otrFieldTrimoffsetpShift))
}

// registerOpamp1_hsotrType OPAMP1 offset trimming register in low-power mode
type registerOpamp1_hsotrType uint32

const (
	RegisterOpamp1_hsotrFieldTrimlpoffsetnShift = 0
	RegisterOpamp1_hsotrFieldTrimlpoffsetnMask  = 0x1f
)

// GetTrimlpoffsetn Trim for NMOS differential pairs
func (r *registerOpamp1_hsotrType) GetTrimlpoffsetn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_hsotrFieldTrimlpoffsetnMask) >> RegisterOpamp1_hsotrFieldTrimlpoffsetnShift)
}

// SetTrimlpoffsetn Trim for NMOS differential pairs
func (r *registerOpamp1_hsotrType) SetTrimlpoffsetn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_hsotrFieldTrimlpoffsetnMask)|(uint32(value)<<RegisterOpamp1_hsotrFieldTrimlpoffsetnShift))
}

const (
	RegisterOpamp1_hsotrFieldTrimlpoffsetpShift = 8
	RegisterOpamp1_hsotrFieldTrimlpoffsetpMask  = 0x1f00
)

// GetTrimlpoffsetp Trim for PMOS differential pairs
func (r *registerOpamp1_hsotrType) GetTrimlpoffsetp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1_hsotrFieldTrimlpoffsetpMask) >> RegisterOpamp1_hsotrFieldTrimlpoffsetpShift)
}

// SetTrimlpoffsetp Trim for PMOS differential pairs
func (r *registerOpamp1_hsotrType) SetTrimlpoffsetp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1_hsotrFieldTrimlpoffsetpMask)|(uint32(value)<<RegisterOpamp1_hsotrFieldTrimlpoffsetpShift))
}

// registerOpamp2_csrType OPAMP2 control/status register
type registerOpamp2_csrType uint32

const (
	RegisterOpamp2_csrFieldOpaenShift = 0
	RegisterOpamp2_csrFieldOpaenMask  = 0x1
)

// GetOpaen Operational amplifier Enable
func (r *registerOpamp2_csrType) GetOpaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_csrFieldOpaenMask) != 0
}

// SetOpaen Operational amplifier Enable
func (r *registerOpamp2_csrType) SetOpaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2_csrFieldOpaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_csrFieldOpaenMask)
	}
}

const (
	RegisterOpamp2_csrFieldForce_vpShift = 1
	RegisterOpamp2_csrFieldForce_vpMask  = 0x2
)

// GetForce_vp Force internal reference on VP (reserved for test)
func (r *registerOpamp2_csrType) GetForce_vp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_csrFieldForce_vpMask) != 0
}

// SetForce_vp Force internal reference on VP (reserved for test)
func (r *registerOpamp2_csrType) SetForce_vp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2_csrFieldForce_vpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_csrFieldForce_vpMask)
	}
}

const (
	RegisterOpamp2_csrFieldVm_selShift = 5
	RegisterOpamp2_csrFieldVm_selMask  = 0x60
)

// GetVm_sel Inverting input selection
func (r *registerOpamp2_csrType) GetVm_sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_csrFieldVm_selMask) >> RegisterOpamp2_csrFieldVm_selShift)
}

// SetVm_sel Inverting input selection
func (r *registerOpamp2_csrType) SetVm_sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_csrFieldVm_selMask)|(uint32(value)<<RegisterOpamp2_csrFieldVm_selShift))
}

const (
	RegisterOpamp2_csrFieldOpahsmShift = 8
	RegisterOpamp2_csrFieldOpahsmMask  = 0x100
)

// GetOpahsm Operational amplifier high-speed mode
func (r *registerOpamp2_csrType) GetOpahsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_csrFieldOpahsmMask) != 0
}

// SetOpahsm Operational amplifier high-speed mode
func (r *registerOpamp2_csrType) SetOpahsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2_csrFieldOpahsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_csrFieldOpahsmMask)
	}
}

const (
	RegisterOpamp2_csrFieldCalonShift = 11
	RegisterOpamp2_csrFieldCalonMask  = 0x800
)

// GetCalon Calibration mode enabled
func (r *registerOpamp2_csrType) GetCalon() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_csrFieldCalonMask) != 0
}

// SetCalon Calibration mode enabled
func (r *registerOpamp2_csrType) SetCalon(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2_csrFieldCalonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_csrFieldCalonMask)
	}
}

const (
	RegisterOpamp2_csrFieldCalselShift = 12
	RegisterOpamp2_csrFieldCalselMask  = 0x3000
)

// GetCalsel Calibration selection
func (r *registerOpamp2_csrType) GetCalsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_csrFieldCalselMask) >> RegisterOpamp2_csrFieldCalselShift)
}

// SetCalsel Calibration selection
func (r *registerOpamp2_csrType) SetCalsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_csrFieldCalselMask)|(uint32(value)<<RegisterOpamp2_csrFieldCalselShift))
}

const (
	RegisterOpamp2_csrFieldPga_gainShift = 14
	RegisterOpamp2_csrFieldPga_gainMask  = 0x3c000
)

// GetPga_gain Operational amplifier Programmable amplifier gain value
func (r *registerOpamp2_csrType) GetPga_gain() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_csrFieldPga_gainMask) >> RegisterOpamp2_csrFieldPga_gainShift)
}

// SetPga_gain Operational amplifier Programmable amplifier gain value
func (r *registerOpamp2_csrType) SetPga_gain(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_csrFieldPga_gainMask)|(uint32(value)<<RegisterOpamp2_csrFieldPga_gainShift))
}

const (
	RegisterOpamp2_csrFieldUsertrimShift = 18
	RegisterOpamp2_csrFieldUsertrimMask  = 0x40000
)

// GetUsertrim User trimming enable
func (r *registerOpamp2_csrType) GetUsertrim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_csrFieldUsertrimMask) != 0
}

// SetUsertrim User trimming enable
func (r *registerOpamp2_csrType) SetUsertrim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2_csrFieldUsertrimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_csrFieldUsertrimMask)
	}
}

const (
	RegisterOpamp2_csrFieldTstrefShift = 29
	RegisterOpamp2_csrFieldTstrefMask  = 0x20000000
)

// GetTstref OPAMP calibration reference voltage output control (reserved for test)
func (r *registerOpamp2_csrType) GetTstref() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_csrFieldTstrefMask) != 0
}

// SetTstref OPAMP calibration reference voltage output control (reserved for test)
func (r *registerOpamp2_csrType) SetTstref(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2_csrFieldTstrefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_csrFieldTstrefMask)
	}
}

const (
	RegisterOpamp2_csrFieldCaloutShift = 30
	RegisterOpamp2_csrFieldCaloutMask  = 0x40000000
)

// GetCalout Operational amplifier calibration output
func (r *registerOpamp2_csrType) GetCalout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_csrFieldCaloutMask) != 0
}

// SetCalout Operational amplifier calibration output
func (r *registerOpamp2_csrType) SetCalout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2_csrFieldCaloutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_csrFieldCaloutMask)
	}
}

// registerOpamp2_otrType OPAMP2 offset trimming register in normal mode
type registerOpamp2_otrType uint32

const (
	RegisterOpamp2_otrFieldTrimoffsetnShift = 0
	RegisterOpamp2_otrFieldTrimoffsetnMask  = 0x1f
)

// GetTrimoffsetn Trim for NMOS differential pairs
func (r *registerOpamp2_otrType) GetTrimoffsetn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_otrFieldTrimoffsetnMask) >> RegisterOpamp2_otrFieldTrimoffsetnShift)
}

// SetTrimoffsetn Trim for NMOS differential pairs
func (r *registerOpamp2_otrType) SetTrimoffsetn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_otrFieldTrimoffsetnMask)|(uint32(value)<<RegisterOpamp2_otrFieldTrimoffsetnShift))
}

const (
	RegisterOpamp2_otrFieldTrimoffsetpShift = 8
	RegisterOpamp2_otrFieldTrimoffsetpMask  = 0x1f00
)

// GetTrimoffsetp Trim for PMOS differential pairs
func (r *registerOpamp2_otrType) GetTrimoffsetp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_otrFieldTrimoffsetpMask) >> RegisterOpamp2_otrFieldTrimoffsetpShift)
}

// SetTrimoffsetp Trim for PMOS differential pairs
func (r *registerOpamp2_otrType) SetTrimoffsetp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_otrFieldTrimoffsetpMask)|(uint32(value)<<RegisterOpamp2_otrFieldTrimoffsetpShift))
}

// registerOpamp2_hsotrType OPAMP2 offset trimming register in low-power mode
type registerOpamp2_hsotrType uint32

const (
	RegisterOpamp2_hsotrFieldTrimlpoffsetnShift = 0
	RegisterOpamp2_hsotrFieldTrimlpoffsetnMask  = 0x1f
)

// GetTrimlpoffsetn Trim for NMOS differential pairs
func (r *registerOpamp2_hsotrType) GetTrimlpoffsetn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_hsotrFieldTrimlpoffsetnMask) >> RegisterOpamp2_hsotrFieldTrimlpoffsetnShift)
}

// SetTrimlpoffsetn Trim for NMOS differential pairs
func (r *registerOpamp2_hsotrType) SetTrimlpoffsetn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_hsotrFieldTrimlpoffsetnMask)|(uint32(value)<<RegisterOpamp2_hsotrFieldTrimlpoffsetnShift))
}

const (
	RegisterOpamp2_hsotrFieldTrimlpoffsetpShift = 8
	RegisterOpamp2_hsotrFieldTrimlpoffsetpMask  = 0x1f00
)

// GetTrimlpoffsetp Trim for PMOS differential pairs
func (r *registerOpamp2_hsotrType) GetTrimlpoffsetp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2_hsotrFieldTrimlpoffsetpMask) >> RegisterOpamp2_hsotrFieldTrimlpoffsetpShift)
}

// SetTrimlpoffsetp Trim for PMOS differential pairs
func (r *registerOpamp2_hsotrType) SetTrimlpoffsetp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2_hsotrFieldTrimlpoffsetpMask)|(uint32(value)<<RegisterOpamp2_hsotrFieldTrimlpoffsetpShift))
}
