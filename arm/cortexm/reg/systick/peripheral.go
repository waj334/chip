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
	Csr   RegisterCsrType
	Rvr   RegisterRvrType
	Cvr   RegisterCvrType
	Calib RegisterCalibType
}

// RegisterCsrType Controls and provides status date for the SysTick timer
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
	RegisterCsrFieldEnableShift = 0
	RegisterCsrFieldEnableMask  = 0x1
)

// GetEnable Enables the counter
func (r *RegisterCsrType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEnableMask) != 0
}

// SetEnable Enables the counter
func (r *RegisterCsrType) SetEnable(value bool) {
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
func (r *RegisterCsrType) GetTickint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldTickintMask) != 0
}

// SetTickint Enables SysTick exception request
func (r *RegisterCsrType) SetTickint(value bool) {
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
func (r *RegisterCsrType) GetClksource() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldClksourceMask) != 0
}

// SetClksource Selects the SysTick timer clock source
func (r *RegisterCsrType) SetClksource(value bool) {
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
func (r *RegisterCsrType) GetCountflag() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldCountflagMask) != 0
}

// SetCountflag Returns 1 if timer counted to 0 since the last read of this register
func (r *RegisterCsrType) SetCountflag(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldCountflagMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldCountflagMask)
	}
}

// RegisterRvrType Specifies the SysTick timer counter reload value
type RegisterRvrType uint32

func (r *RegisterRvrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRvrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRvrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRvrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRvrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRvrFieldReloadShift = 0
	RegisterRvrFieldReloadMask  = 0xffffff
)

// GetReload Value to load into the SYST_CVR when the counter is enabled and when it reaches 0
func (r *RegisterRvrType) GetReload() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRvrFieldReloadMask) >> RegisterRvrFieldReloadShift)
}

// SetReload Value to load into the SYST_CVR when the counter is enabled and when it reaches 0
func (r *RegisterRvrType) SetReload(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRvrFieldReloadMask)|(uint32(value)<<RegisterRvrFieldReloadShift))
}

// RegisterCvrType Contains the current value of the SysTick counter
type RegisterCvrType uint32

func (r *RegisterCvrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCvrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCvrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCvrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCvrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCvrFieldCurrentShift = 0
	RegisterCvrFieldCurrentMask  = 0xffffff
)

// GetCurrent Current value of the SysTick counter. A write of any value clears the field to 0, and also clears the SYST_CSR.COUNTFLAG bit to 0.
func (r *RegisterCvrType) GetCurrent() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCvrFieldCurrentMask) >> RegisterCvrFieldCurrentShift)
}

// SetCurrent Current value of the SysTick counter. A write of any value clears the field to 0, and also clears the SYST_CSR.COUNTFLAG bit to 0.
func (r *RegisterCvrType) SetCurrent(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCvrFieldCurrentMask)|(uint32(value)<<RegisterCvrFieldCurrentShift))
}

// RegisterCalibType Indicates the SysTick calibration value and parameters for the selected Security state
type RegisterCalibType uint32

func (r *RegisterCalibType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCalibType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCalibType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCalibType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCalibType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCalibFieldTenmsShift = 0
	RegisterCalibFieldTenmsMask  = 0xffffff
)

// GetTenms Reload value for 10ms (100Hz) timing, subject to system clock skew errors. If the value reads as zero, the calibration value is not known.
func (r *RegisterCalibType) GetTenms() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCalibFieldTenmsMask) >> RegisterCalibFieldTenmsShift)
}

// SetTenms Reload value for 10ms (100Hz) timing, subject to system clock skew errors. If the value reads as zero, the calibration value is not known.
func (r *RegisterCalibType) SetTenms(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCalibFieldTenmsMask)|(uint32(value)<<RegisterCalibFieldTenmsShift))
}

const (
	RegisterCalibFieldSkewShift = 30
	RegisterCalibFieldSkewMask  = 0x40000000
)

// GetSkew Indicates whether the TENMS value is exact
func (r *RegisterCalibType) GetSkew() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCalibFieldSkewMask) != 0
}

// SetSkew Indicates whether the TENMS value is exact
func (r *RegisterCalibType) SetSkew(value bool) {
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
func (r *RegisterCalibType) GetNoref() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCalibFieldNorefMask) != 0
}

// SetNoref Indicates whether the device provides a reference clock to the processor
func (r *RegisterCalibType) SetNoref(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCalibFieldNorefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCalibFieldNorefMask)
	}
}
