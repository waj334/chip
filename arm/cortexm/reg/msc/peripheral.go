//go:build arm && cortexm

package msc

import (
	"unsafe"
	"volatile"
)

var (
	Msc = (*_msc)(unsafe.Pointer(uintptr(0xe001e000)))
)

type _msc struct {
	Mscr   RegisterMscrType
	Pfcr   RegisterPfcrType
	_      [8]byte
	Itcmcr RegisterItcmcrType
	Dtcmcr RegisterDtcmcrType
	Pahbcr RegisterPahbcrType
}

// RegisterMscrType Controls the memory system features specific to the processor
type RegisterMscrType uint32

func (r *RegisterMscrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMscrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMscrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMscrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMscrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMscrFieldEccenShift = 1
	RegisterMscrFieldEccenMask  = 0x2
)

// GetEccen Indicates whether Error Correcting Code (ECC) is present and enabled
func (r *RegisterMscrType) GetEccen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldEccenMask) != 0
}

// SetEccen Indicates whether Error Correcting Code (ECC) is present and enabled
func (r *RegisterMscrType) SetEccen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMscrFieldEccenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMscrFieldEccenMask)
	}
}

const (
	RegisterMscrFieldForcewtShift = 2
	RegisterMscrFieldForcewtMask  = 0x4
)

// GetForcewt Enables Forced Write-Through in the L1 data cache
func (r *RegisterMscrType) GetForcewt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldForcewtMask) != 0
}

// SetForcewt Enables Forced Write-Through in the L1 data cache
func (r *RegisterMscrType) SetForcewt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMscrFieldForcewtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMscrFieldForcewtMask)
	}
}

const (
	RegisterMscrFieldEveccfaultShift = 3
	RegisterMscrFieldEveccfaultMask  = 0x8
)

// GetEveccfault Enables asynchronous BusFault exceptions when data is lost on evictions
func (r *RegisterMscrType) GetEveccfault() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldEveccfaultMask) != 0
}

// SetEveccfault Enables asynchronous BusFault exceptions when data is lost on evictions
func (r *RegisterMscrType) SetEveccfault(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMscrFieldEveccfaultMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMscrFieldEveccfaultMask)
	}
}

const (
	RegisterMscrFieldDcactiveShift = 12
	RegisterMscrFieldDcactiveMask  = 0x1000
)

// GetDcactive Indicates whether the L1 data cache is active
func (r *RegisterMscrType) GetDcactive() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldDcactiveMask) != 0
}

// SetDcactive Indicates whether the L1 data cache is active
func (r *RegisterMscrType) SetDcactive(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMscrFieldDcactiveMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMscrFieldDcactiveMask)
	}
}

const (
	RegisterMscrFieldIcactiveShift = 13
	RegisterMscrFieldIcactiveMask  = 0x2000
)

// GetIcactive Indicates whether the L1 instruction cache is active
func (r *RegisterMscrType) GetIcactive() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldIcactiveMask) != 0
}

// SetIcactive Indicates whether the L1 instruction cache is active
func (r *RegisterMscrType) SetIcactive(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMscrFieldIcactiveMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMscrFieldIcactiveMask)
	}
}

const (
	RegisterMscrFieldDccleanShift = 16
	RegisterMscrFieldDccleanMask  = 0x10000
)

// GetDcclean Indicates whether the data cache contains any dirty lines
func (r *RegisterMscrType) GetDcclean() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldDccleanMask) != 0
}

// SetDcclean Indicates whether the data cache contains any dirty lines
func (r *RegisterMscrType) SetDcclean(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMscrFieldDccleanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMscrFieldDccleanMask)
	}
}

const (
	RegisterMscrFieldCpwrdnShift = 17
	RegisterMscrFieldCpwrdnMask  = 0x20000
)

// GetCpwrdn Indicates when the data and instruction caches are not accessible because they are either being powered down or being initialized using the automatic invalidation sequence
func (r *RegisterMscrType) GetCpwrdn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldCpwrdnMask) != 0
}

// SetCpwrdn Indicates when the data and instruction caches are not accessible because they are either being powered down or being initialized using the automatic invalidation sequence
func (r *RegisterMscrType) SetCpwrdn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMscrFieldCpwrdnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMscrFieldCpwrdnMask)
	}
}

// RegisterPfcrType Controls the prefetcher
type RegisterPfcrType uint32

func (r *RegisterPfcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPfcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPfcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPfcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPfcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPfcrFieldEnableShift = 0
	RegisterPfcrFieldEnableMask  = 0x1
)

// GetEnable Prefetcher enable
func (r *RegisterPfcrType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPfcrFieldEnableMask) != 0
}

// SetEnable Prefetcher enable
func (r *RegisterPfcrType) SetEnable(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPfcrFieldEnableMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPfcrFieldEnableMask)
	}
}

const (
	RegisterPfcrFieldDisnlpShift = 7
	RegisterPfcrFieldDisnlpMask  = 0x80
)

// GetDisnlp Disables Next Line Prefetch mode
func (r *RegisterPfcrType) GetDisnlp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPfcrFieldDisnlpMask) != 0
}

// SetDisnlp Disables Next Line Prefetch mode
func (r *RegisterPfcrType) SetDisnlp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPfcrFieldDisnlpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPfcrFieldDisnlpMask)
	}
}

// RegisterItcmcrType Enable access to the Tightly Coupled Memories (TCMs) by software running on the processor
type RegisterItcmcrType uint32

func (r *RegisterItcmcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterItcmcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterItcmcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterItcmcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterItcmcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterItcmcrFieldEnShift = 0
	RegisterItcmcrFieldEnMask  = 0x1
)

// GetEn TCM enable
func (r *RegisterItcmcrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterItcmcrFieldEnMask) != 0
}

// SetEn TCM enable
func (r *RegisterItcmcrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterItcmcrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterItcmcrFieldEnMask)
	}
}

const (
	RegisterItcmcrFieldSzShift = 3
	RegisterItcmcrFieldSzMask  = 0x78
)

// GetSz Indicates the size of the relevant TCM
func (r *RegisterItcmcrType) GetSz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterItcmcrFieldSzMask) >> RegisterItcmcrFieldSzShift)
}

// SetSz Indicates the size of the relevant TCM
func (r *RegisterItcmcrType) SetSz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterItcmcrFieldSzMask)|(uint32(value)<<RegisterItcmcrFieldSzShift))
}

// RegisterDtcmcrType Enable access to the Tightly Coupled Memories (TCMs) by software running on the processor
type RegisterDtcmcrType uint32

func (r *RegisterDtcmcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDtcmcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDtcmcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDtcmcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDtcmcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDtcmcrFieldEnShift = 0
	RegisterDtcmcrFieldEnMask  = 0x1
)

// GetEn TCM enable
func (r *RegisterDtcmcrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtcmcrFieldEnMask) != 0
}

// SetEn TCM enable
func (r *RegisterDtcmcrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtcmcrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtcmcrFieldEnMask)
	}
}

const (
	RegisterDtcmcrFieldSzShift = 3
	RegisterDtcmcrFieldSzMask  = 0x78
)

// GetSz Indicates the size of the relevant TCM
func (r *RegisterDtcmcrType) GetSz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDtcmcrFieldSzMask) >> RegisterDtcmcrFieldSzShift)
}

// SetSz Indicates the size of the relevant TCM
func (r *RegisterDtcmcrType) SetSz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtcmcrFieldSzMask)|(uint32(value)<<RegisterDtcmcrFieldSzShift))
}

// RegisterPahbcrType Enables accesses to Peripheral AHB (P-AHB) interface from software running on the processor
type RegisterPahbcrType uint32

func (r *RegisterPahbcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPahbcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPahbcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPahbcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPahbcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPahbcrFieldEnShift = 0
	RegisterPahbcrFieldEnMask  = 0x1
)

// GetEn P-AHB enable
func (r *RegisterPahbcrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPahbcrFieldEnMask) != 0
}

// SetEn P-AHB enable
func (r *RegisterPahbcrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPahbcrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPahbcrFieldEnMask)
	}
}

const (
	RegisterPahbcrFieldSzShift = 1
	RegisterPahbcrFieldSzMask  = 0xe
)

// GetSz P-AHB size
func (r *RegisterPahbcrType) GetSz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPahbcrFieldSzMask) >> RegisterPahbcrFieldSzShift)
}

// SetSz P-AHB size
func (r *RegisterPahbcrType) SetSz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPahbcrFieldSzMask)|(uint32(value)<<RegisterPahbcrFieldSzShift))
}
