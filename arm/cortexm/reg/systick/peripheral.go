//go:build arm && cortexm

package systick

import (
	"unsafe"
	"volatile"
)

var (
	Systick = (*_systick)(unsafe.Pointer(uintptr(0xe000e010)))
)

type _systick struct {
	Csr   registerCsrType
	Rvr   registerRvrType
	Cvr   registerCvrType
	Calib registerCalibType
}

// registerCsrType Controls and provides status date for the SysTick timer
type registerCsrType uint32

const (
	RegisterCsrFieldEnableShift = 0
	RegisterCsrFieldEnableMask  = 0x1
)

// GetEnable Enables the counter
func (r *registerCsrType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEnableMask) != 0
}

// SetEnable Enables the counter
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

// GetTickint Enables SysTick exception request
func (r *registerCsrType) GetTickint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldTickintMask) != 0
}

// SetTickint Enables SysTick exception request
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

// GetClksource Selects the SysTick timer clock source
func (r *registerCsrType) GetClksource() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldClksourceMask) != 0
}

// SetClksource Selects the SysTick timer clock source
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

// GetCountflag Returns 1 if timer counted to 0 since the last read of this register
func (r *registerCsrType) GetCountflag() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldCountflagMask) != 0
}

// SetCountflag Returns 1 if timer counted to 0 since the last read of this register
func (r *registerCsrType) SetCountflag(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldCountflagMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldCountflagMask)
	}
}

// registerRvrType Specifies the SysTick timer counter reload value
type registerRvrType uint32

const (
	RegisterRvrFieldReloadShift = 0
	RegisterRvrFieldReloadMask  = 0xffffff
)

// GetReload Value to load into the SYST_CVR when the counter is enabled and when it reaches 0
func (r *registerRvrType) GetReload() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRvrFieldReloadMask) >> RegisterRvrFieldReloadShift)
}

// SetReload Value to load into the SYST_CVR when the counter is enabled and when it reaches 0
func (r *registerRvrType) SetReload(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRvrFieldReloadMask)|(uint32(value)<<RegisterRvrFieldReloadShift))
}

// registerCvrType Contains the current value of the SysTick counter
type registerCvrType uint32

const (
	RegisterCvrFieldCurrentShift = 0
	RegisterCvrFieldCurrentMask  = 0xffffff
)

// GetCurrent Current value of the SysTick counter. A write of any value clears the field to 0, and also clears the SYST_CSR.COUNTFLAG bit to 0.
func (r *registerCvrType) GetCurrent() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCvrFieldCurrentMask) >> RegisterCvrFieldCurrentShift)
}

// SetCurrent Current value of the SysTick counter. A write of any value clears the field to 0, and also clears the SYST_CSR.COUNTFLAG bit to 0.
func (r *registerCvrType) SetCurrent(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCvrFieldCurrentMask)|(uint32(value)<<RegisterCvrFieldCurrentShift))
}

// registerCalibType Indicates the SysTick calibration value and parameters for the selected Security state
type registerCalibType uint32

const (
	RegisterCalibFieldTenmsShift = 0
	RegisterCalibFieldTenmsMask  = 0xffffff
)

// GetTenms Reload value for 10ms (100Hz) timing, subject to system clock skew errors. If the value reads as zero, the calibration value is not known.
func (r *registerCalibType) GetTenms() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCalibFieldTenmsMask) >> RegisterCalibFieldTenmsShift)
}

// SetTenms Reload value for 10ms (100Hz) timing, subject to system clock skew errors. If the value reads as zero, the calibration value is not known.
func (r *registerCalibType) SetTenms(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCalibFieldTenmsMask)|(uint32(value)<<RegisterCalibFieldTenmsShift))
}

const (
	RegisterCalibFieldSkewShift = 30
	RegisterCalibFieldSkewMask  = 0x40000000
)

// GetSkew Indicates whether the TENMS value is exact
func (r *registerCalibType) GetSkew() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCalibFieldSkewMask) != 0
}

// SetSkew Indicates whether the TENMS value is exact
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

// GetNoref Indicates whether the device provides a reference clock to the processor
func (r *registerCalibType) GetNoref() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCalibFieldNorefMask) != 0
}

// SetNoref Indicates whether the device provides a reference clock to the processor
func (r *registerCalibType) SetNoref(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCalibFieldNorefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCalibFieldNorefMask)
	}
}
