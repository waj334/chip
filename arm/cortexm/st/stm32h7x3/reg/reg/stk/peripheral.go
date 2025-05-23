package stk

import (
	"unsafe"
	"volatile"
)

var (
	Stk = (*_stk)(unsafe.Pointer(uintptr(0xe000e010)))
)

type _stk struct {
	Csr   registerCsrType
	Rvr   registerRvrType
	Cvr   registerCvrType
	Calib registerCalibType
}

// registerCsrType SysTick control and status register
type registerCsrType uint32

const (
	RegisterCsrFieldEnableShift = 0
	RegisterCsrFieldEnableMask  = 0x1
)

// GetEnable Counter enable
func (r *registerCsrType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEnableMask) != 0
}

// SetEnable Counter enable
func (r *registerCsrType) SetEnable(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEnableMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEnableMask)
	}
}

const (
	RegisterCsrFieldTickintShift = 1
	RegisterCsrFieldTickintMask  = 0x2
)

// GetTickint SysTick exception request enable
func (r *registerCsrType) GetTickint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldTickintMask) != 0
}

// SetTickint SysTick exception request enable
func (r *registerCsrType) SetTickint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldTickintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldTickintMask)
	}
}

const (
	RegisterCsrFieldClksourceShift = 2
	RegisterCsrFieldClksourceMask  = 0x4
)

// GetClksource Clock source selection
func (r *registerCsrType) GetClksource() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldClksourceMask) != 0
}

// SetClksource Clock source selection
func (r *registerCsrType) SetClksource(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldClksourceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldClksourceMask)
	}
}

const (
	RegisterCsrFieldCountflagShift = 16
	RegisterCsrFieldCountflagMask  = 0x10000
)

// GetCountflag COUNTFLAG
func (r *registerCsrType) GetCountflag() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldCountflagMask) != 0
}

// SetCountflag COUNTFLAG
func (r *registerCsrType) SetCountflag(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldCountflagMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldCountflagMask)
	}
}

// registerRvrType SysTick reload value register
type registerRvrType uint32

const (
	RegisterRvrFieldReloadShift = 0
	RegisterRvrFieldReloadMask  = 0xffffff
)

// GetReload RELOAD value
func (r *registerRvrType) GetReload() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRvrFieldReloadMask) >> RegisterRvrFieldReloadShift)
}

// SetReload RELOAD value
func (r *registerRvrType) SetReload(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRvrFieldReloadMask)|(uint32(value)<<RegisterRvrFieldReloadShift))
}

// registerCvrType SysTick current value register
type registerCvrType uint32

const (
	RegisterCvrFieldCurrentShift = 0
	RegisterCvrFieldCurrentMask  = 0xffffff
)

// GetCurrent Current counter value
func (r *registerCvrType) GetCurrent() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCvrFieldCurrentMask) >> RegisterCvrFieldCurrentShift)
}

// SetCurrent Current counter value
func (r *registerCvrType) SetCurrent(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCvrFieldCurrentMask)|(uint32(value)<<RegisterCvrFieldCurrentShift))
}

// registerCalibType SysTick calibration value register
type registerCalibType uint32

const (
	RegisterCalibFieldTenmsShift = 0
	RegisterCalibFieldTenmsMask  = 0xffffff
)

// GetTenms Calibration value
func (r *registerCalibType) GetTenms() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCalibFieldTenmsMask) >> RegisterCalibFieldTenmsShift)
}

// SetTenms Calibration value
func (r *registerCalibType) SetTenms(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCalibFieldTenmsMask)|(uint32(value)<<RegisterCalibFieldTenmsShift))
}

const (
	RegisterCalibFieldSkewShift = 30
	RegisterCalibFieldSkewMask  = 0x40000000
)

// GetSkew SKEW flag: Indicates whether the TENMS value is exact
func (r *registerCalibType) GetSkew() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCalibFieldSkewMask) != 0
}

// SetSkew SKEW flag: Indicates whether the TENMS value is exact
func (r *registerCalibType) SetSkew(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCalibFieldSkewMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCalibFieldSkewMask)
	}
}

const (
	RegisterCalibFieldNorefShift = 31
	RegisterCalibFieldNorefMask  = 0x80000000
)

// GetNoref NOREF flag. Reads as zero
func (r *registerCalibType) GetNoref() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCalibFieldNorefMask) != 0
}

// SetNoref NOREF flag. Reads as zero
func (r *registerCalibType) SetNoref(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCalibFieldNorefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCalibFieldNorefMask)
	}
}
