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
	Opamp1csr   RegisterOpamp1csrType
	Opamp1otr   RegisterOpamp1otrType
	Opamp1hsotr RegisterOpamp1hsotrType
	_           [4]byte
	Opamp2csr   RegisterOpamp2csrType
	Opamp2otr   RegisterOpamp2otrType
	Opamp2hsotr RegisterOpamp2hsotrType
}

// RegisterOpamp1csrType OPAMP1 control/status register
type RegisterOpamp1csrType uint32

func (r *RegisterOpamp1csrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOpamp1csrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOpamp1csrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOpamp1csrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOpamp1csrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOpamp1csrFieldOpaenShift = 0
	RegisterOpamp1csrFieldOpaenMask  = 0x1
)

// GetOpaen Operational amplifier Enable
func (r *RegisterOpamp1csrType) GetOpaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldOpaenMask) != 0
}

// SetOpaen Operational amplifier Enable
func (r *RegisterOpamp1csrType) SetOpaen(value bool) {
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
func (r *RegisterOpamp1csrType) GetForcevp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldForcevpMask) != 0
}

// SetForcevp Force internal reference on VP (reserved for test
func (r *RegisterOpamp1csrType) SetForcevp(value bool) {
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
func (r *RegisterOpamp1csrType) GetVpsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldVpselMask) >> RegisterOpamp1csrFieldVpselShift)
}

// SetVpsel Operational amplifier PGA mode
func (r *RegisterOpamp1csrType) SetVpsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldVpselMask)|(uint32(value)<<RegisterOpamp1csrFieldVpselShift))
}

const (
	RegisterOpamp1csrFieldVmselShift = 5
	RegisterOpamp1csrFieldVmselMask  = 0x60
)

// GetVmsel Inverting input selection
func (r *RegisterOpamp1csrType) GetVmsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldVmselMask) >> RegisterOpamp1csrFieldVmselShift)
}

// SetVmsel Inverting input selection
func (r *RegisterOpamp1csrType) SetVmsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldVmselMask)|(uint32(value)<<RegisterOpamp1csrFieldVmselShift))
}

const (
	RegisterOpamp1csrFieldOpahsmShift = 8
	RegisterOpamp1csrFieldOpahsmMask  = 0x100
)

// GetOpahsm Operational amplifier high-speed mode
func (r *RegisterOpamp1csrType) GetOpahsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldOpahsmMask) != 0
}

// SetOpahsm Operational amplifier high-speed mode
func (r *RegisterOpamp1csrType) SetOpahsm(value bool) {
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
func (r *RegisterOpamp1csrType) GetCalon() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldCalonMask) != 0
}

// SetCalon Calibration mode enabled
func (r *RegisterOpamp1csrType) SetCalon(value bool) {
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
func (r *RegisterOpamp1csrType) GetCalsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldCalselMask) >> RegisterOpamp1csrFieldCalselShift)
}

// SetCalsel Calibration selection
func (r *RegisterOpamp1csrType) SetCalsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldCalselMask)|(uint32(value)<<RegisterOpamp1csrFieldCalselShift))
}

const (
	RegisterOpamp1csrFieldPgagainShift = 14
	RegisterOpamp1csrFieldPgagainMask  = 0x3c000
)

// GetPgagain allows to switch from AOP offset trimmed values to AOP offset
func (r *RegisterOpamp1csrType) GetPgagain() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldPgagainMask) >> RegisterOpamp1csrFieldPgagainShift)
}

// SetPgagain allows to switch from AOP offset trimmed values to AOP offset
func (r *RegisterOpamp1csrType) SetPgagain(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldPgagainMask)|(uint32(value)<<RegisterOpamp1csrFieldPgagainShift))
}

const (
	RegisterOpamp1csrFieldUsertrimShift = 18
	RegisterOpamp1csrFieldUsertrimMask  = 0x40000
)

// GetUsertrim User trimming enable
func (r *RegisterOpamp1csrType) GetUsertrim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldUsertrimMask) != 0
}

// SetUsertrim User trimming enable
func (r *RegisterOpamp1csrType) SetUsertrim(value bool) {
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
func (r *RegisterOpamp1csrType) GetTstref() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldTstrefMask) != 0
}

// SetTstref OPAMP calibration reference voltage output control (reserved for test)
func (r *RegisterOpamp1csrType) SetTstref(value bool) {
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
func (r *RegisterOpamp1csrType) GetCalout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1csrFieldCaloutMask) != 0
}

// SetCalout Operational amplifier calibration output
func (r *RegisterOpamp1csrType) SetCalout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp1csrFieldCaloutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1csrFieldCaloutMask)
	}
}

// RegisterOpamp1otrType OPAMP1 offset trimming register in normal mode
type RegisterOpamp1otrType uint32

func (r *RegisterOpamp1otrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOpamp1otrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOpamp1otrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOpamp1otrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOpamp1otrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOpamp1otrFieldTrimoffsetnShift = 0
	RegisterOpamp1otrFieldTrimoffsetnMask  = 0x1f
)

// GetTrimoffsetn Trim for NMOS differential pairs
func (r *RegisterOpamp1otrType) GetTrimoffsetn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1otrFieldTrimoffsetnMask) >> RegisterOpamp1otrFieldTrimoffsetnShift)
}

// SetTrimoffsetn Trim for NMOS differential pairs
func (r *RegisterOpamp1otrType) SetTrimoffsetn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1otrFieldTrimoffsetnMask)|(uint32(value)<<RegisterOpamp1otrFieldTrimoffsetnShift))
}

const (
	RegisterOpamp1otrFieldTrimoffsetpShift = 8
	RegisterOpamp1otrFieldTrimoffsetpMask  = 0x1f00
)

// GetTrimoffsetp Trim for PMOS differential pairs
func (r *RegisterOpamp1otrType) GetTrimoffsetp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1otrFieldTrimoffsetpMask) >> RegisterOpamp1otrFieldTrimoffsetpShift)
}

// SetTrimoffsetp Trim for PMOS differential pairs
func (r *RegisterOpamp1otrType) SetTrimoffsetp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1otrFieldTrimoffsetpMask)|(uint32(value)<<RegisterOpamp1otrFieldTrimoffsetpShift))
}

// RegisterOpamp1hsotrType OPAMP1 offset trimming register in low-power mode
type RegisterOpamp1hsotrType uint32

func (r *RegisterOpamp1hsotrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOpamp1hsotrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOpamp1hsotrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOpamp1hsotrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOpamp1hsotrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOpamp1hsotrFieldTrimlpoffsetnShift = 0
	RegisterOpamp1hsotrFieldTrimlpoffsetnMask  = 0x1f
)

// GetTrimlpoffsetn Trim for NMOS differential pairs
func (r *RegisterOpamp1hsotrType) GetTrimlpoffsetn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1hsotrFieldTrimlpoffsetnMask) >> RegisterOpamp1hsotrFieldTrimlpoffsetnShift)
}

// SetTrimlpoffsetn Trim for NMOS differential pairs
func (r *RegisterOpamp1hsotrType) SetTrimlpoffsetn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1hsotrFieldTrimlpoffsetnMask)|(uint32(value)<<RegisterOpamp1hsotrFieldTrimlpoffsetnShift))
}

const (
	RegisterOpamp1hsotrFieldTrimlpoffsetpShift = 8
	RegisterOpamp1hsotrFieldTrimlpoffsetpMask  = 0x1f00
)

// GetTrimlpoffsetp Trim for PMOS differential pairs
func (r *RegisterOpamp1hsotrType) GetTrimlpoffsetp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp1hsotrFieldTrimlpoffsetpMask) >> RegisterOpamp1hsotrFieldTrimlpoffsetpShift)
}

// SetTrimlpoffsetp Trim for PMOS differential pairs
func (r *RegisterOpamp1hsotrType) SetTrimlpoffsetp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp1hsotrFieldTrimlpoffsetpMask)|(uint32(value)<<RegisterOpamp1hsotrFieldTrimlpoffsetpShift))
}

// RegisterOpamp2csrType OPAMP2 control/status register
type RegisterOpamp2csrType uint32

func (r *RegisterOpamp2csrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOpamp2csrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOpamp2csrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOpamp2csrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOpamp2csrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOpamp2csrFieldOpaenShift = 0
	RegisterOpamp2csrFieldOpaenMask  = 0x1
)

// GetOpaen Operational amplifier Enable
func (r *RegisterOpamp2csrType) GetOpaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldOpaenMask) != 0
}

// SetOpaen Operational amplifier Enable
func (r *RegisterOpamp2csrType) SetOpaen(value bool) {
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
func (r *RegisterOpamp2csrType) GetForcevp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldForcevpMask) != 0
}

// SetForcevp Force internal reference on VP (reserved for test)
func (r *RegisterOpamp2csrType) SetForcevp(value bool) {
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
func (r *RegisterOpamp2csrType) GetVmsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldVmselMask) >> RegisterOpamp2csrFieldVmselShift)
}

// SetVmsel Inverting input selection
func (r *RegisterOpamp2csrType) SetVmsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldVmselMask)|(uint32(value)<<RegisterOpamp2csrFieldVmselShift))
}

const (
	RegisterOpamp2csrFieldOpahsmShift = 8
	RegisterOpamp2csrFieldOpahsmMask  = 0x100
)

// GetOpahsm Operational amplifier high-speed mode
func (r *RegisterOpamp2csrType) GetOpahsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldOpahsmMask) != 0
}

// SetOpahsm Operational amplifier high-speed mode
func (r *RegisterOpamp2csrType) SetOpahsm(value bool) {
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
func (r *RegisterOpamp2csrType) GetCalon() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldCalonMask) != 0
}

// SetCalon Calibration mode enabled
func (r *RegisterOpamp2csrType) SetCalon(value bool) {
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
func (r *RegisterOpamp2csrType) GetCalsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldCalselMask) >> RegisterOpamp2csrFieldCalselShift)
}

// SetCalsel Calibration selection
func (r *RegisterOpamp2csrType) SetCalsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldCalselMask)|(uint32(value)<<RegisterOpamp2csrFieldCalselShift))
}

const (
	RegisterOpamp2csrFieldPgagainShift = 14
	RegisterOpamp2csrFieldPgagainMask  = 0x3c000
)

// GetPgagain Operational amplifier Programmable amplifier gain value
func (r *RegisterOpamp2csrType) GetPgagain() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldPgagainMask) >> RegisterOpamp2csrFieldPgagainShift)
}

// SetPgagain Operational amplifier Programmable amplifier gain value
func (r *RegisterOpamp2csrType) SetPgagain(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldPgagainMask)|(uint32(value)<<RegisterOpamp2csrFieldPgagainShift))
}

const (
	RegisterOpamp2csrFieldUsertrimShift = 18
	RegisterOpamp2csrFieldUsertrimMask  = 0x40000
)

// GetUsertrim User trimming enable
func (r *RegisterOpamp2csrType) GetUsertrim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldUsertrimMask) != 0
}

// SetUsertrim User trimming enable
func (r *RegisterOpamp2csrType) SetUsertrim(value bool) {
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
func (r *RegisterOpamp2csrType) GetTstref() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldTstrefMask) != 0
}

// SetTstref OPAMP calibration reference voltage output control (reserved for test)
func (r *RegisterOpamp2csrType) SetTstref(value bool) {
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
func (r *RegisterOpamp2csrType) GetCalout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2csrFieldCaloutMask) != 0
}

// SetCalout Operational amplifier calibration output
func (r *RegisterOpamp2csrType) SetCalout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOpamp2csrFieldCaloutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2csrFieldCaloutMask)
	}
}

// RegisterOpamp2otrType OPAMP2 offset trimming register in normal mode
type RegisterOpamp2otrType uint32

func (r *RegisterOpamp2otrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOpamp2otrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOpamp2otrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOpamp2otrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOpamp2otrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOpamp2otrFieldTrimoffsetnShift = 0
	RegisterOpamp2otrFieldTrimoffsetnMask  = 0x1f
)

// GetTrimoffsetn Trim for NMOS differential pairs
func (r *RegisterOpamp2otrType) GetTrimoffsetn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2otrFieldTrimoffsetnMask) >> RegisterOpamp2otrFieldTrimoffsetnShift)
}

// SetTrimoffsetn Trim for NMOS differential pairs
func (r *RegisterOpamp2otrType) SetTrimoffsetn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2otrFieldTrimoffsetnMask)|(uint32(value)<<RegisterOpamp2otrFieldTrimoffsetnShift))
}

const (
	RegisterOpamp2otrFieldTrimoffsetpShift = 8
	RegisterOpamp2otrFieldTrimoffsetpMask  = 0x1f00
)

// GetTrimoffsetp Trim for PMOS differential pairs
func (r *RegisterOpamp2otrType) GetTrimoffsetp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2otrFieldTrimoffsetpMask) >> RegisterOpamp2otrFieldTrimoffsetpShift)
}

// SetTrimoffsetp Trim for PMOS differential pairs
func (r *RegisterOpamp2otrType) SetTrimoffsetp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2otrFieldTrimoffsetpMask)|(uint32(value)<<RegisterOpamp2otrFieldTrimoffsetpShift))
}

// RegisterOpamp2hsotrType OPAMP2 offset trimming register in low-power mode
type RegisterOpamp2hsotrType uint32

func (r *RegisterOpamp2hsotrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOpamp2hsotrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOpamp2hsotrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOpamp2hsotrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOpamp2hsotrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOpamp2hsotrFieldTrimlpoffsetnShift = 0
	RegisterOpamp2hsotrFieldTrimlpoffsetnMask  = 0x1f
)

// GetTrimlpoffsetn Trim for NMOS differential pairs
func (r *RegisterOpamp2hsotrType) GetTrimlpoffsetn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2hsotrFieldTrimlpoffsetnMask) >> RegisterOpamp2hsotrFieldTrimlpoffsetnShift)
}

// SetTrimlpoffsetn Trim for NMOS differential pairs
func (r *RegisterOpamp2hsotrType) SetTrimlpoffsetn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2hsotrFieldTrimlpoffsetnMask)|(uint32(value)<<RegisterOpamp2hsotrFieldTrimlpoffsetnShift))
}

const (
	RegisterOpamp2hsotrFieldTrimlpoffsetpShift = 8
	RegisterOpamp2hsotrFieldTrimlpoffsetpMask  = 0x1f00
)

// GetTrimlpoffsetp Trim for PMOS differential pairs
func (r *RegisterOpamp2hsotrType) GetTrimlpoffsetp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOpamp2hsotrFieldTrimlpoffsetpMask) >> RegisterOpamp2hsotrFieldTrimlpoffsetpShift)
}

// SetTrimlpoffsetp Trim for PMOS differential pairs
func (r *RegisterOpamp2hsotrType) SetTrimlpoffsetp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOpamp2hsotrFieldTrimlpoffsetpMask)|(uint32(value)<<RegisterOpamp2hsotrFieldTrimlpoffsetpShift))
}
