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
	Mscr   registerMscrType
	Pfcr   registerPfcrType
	_      [8]byte
	Itcmcr registerItcmcrType
	Dtcmcr registerDtcmcrType
	Pahbcr registerPahbcrType
}

// registerMscrType Controls the memory system features specific to the processor
type registerMscrType uint32

const (
	RegisterMscrFieldEccenShift = 1
	RegisterMscrFieldEccenMask  = 0x2
)

// GetEccen Indicates whether Error Correcting Code (ECC) is present and enabled
func (r *registerMscrType) GetEccen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldEccenMask) != 0
}

// SetEccen Indicates whether Error Correcting Code (ECC) is present and enabled
func (r *registerMscrType) SetEccen(value bool) {
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
func (r *registerMscrType) GetForcewt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldForcewtMask) != 0
}

// SetForcewt Enables Forced Write-Through in the L1 data cache
func (r *registerMscrType) SetForcewt(value bool) {
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
func (r *registerMscrType) GetEveccfault() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldEveccfaultMask) != 0
}

// SetEveccfault Enables asynchronous BusFault exceptions when data is lost on evictions
func (r *registerMscrType) SetEveccfault(value bool) {
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
func (r *registerMscrType) GetDcactive() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldDcactiveMask) != 0
}

// SetDcactive Indicates whether the L1 data cache is active
func (r *registerMscrType) SetDcactive(value bool) {
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
func (r *registerMscrType) GetIcactive() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldIcactiveMask) != 0
}

// SetIcactive Indicates whether the L1 instruction cache is active
func (r *registerMscrType) SetIcactive(value bool) {
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
func (r *registerMscrType) GetDcclean() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldDccleanMask) != 0
}

// SetDcclean Indicates whether the data cache contains any dirty lines
func (r *registerMscrType) SetDcclean(value bool) {
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
func (r *registerMscrType) GetCpwrdn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMscrFieldCpwrdnMask) != 0
}

// SetCpwrdn Indicates when the data and instruction caches are not accessible because they are either being powered down or being initialized using the automatic invalidation sequence
func (r *registerMscrType) SetCpwrdn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMscrFieldCpwrdnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMscrFieldCpwrdnMask)
	}
}

// registerPfcrType Controls the prefetcher
type registerPfcrType uint32

const (
	RegisterPfcrFieldEnableShift = 0
	RegisterPfcrFieldEnableMask  = 0x1
)

// GetEnable Prefetcher enable
func (r *registerPfcrType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPfcrFieldEnableMask) != 0
}

// SetEnable Prefetcher enable
func (r *registerPfcrType) SetEnable(value bool) {
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
func (r *registerPfcrType) GetDisnlp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPfcrFieldDisnlpMask) != 0
}

// SetDisnlp Disables Next Line Prefetch mode
func (r *registerPfcrType) SetDisnlp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPfcrFieldDisnlpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPfcrFieldDisnlpMask)
	}
}

// registerItcmcrType Enable access to the Tightly Coupled Memories (TCMs) by software running on the processor
type registerItcmcrType uint32

const (
	RegisterItcmcrFieldEnShift = 0
	RegisterItcmcrFieldEnMask  = 0x1
)

// GetEn TCM enable
func (r *registerItcmcrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterItcmcrFieldEnMask) != 0
}

// SetEn TCM enable
func (r *registerItcmcrType) SetEn(value bool) {
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
func (r *registerItcmcrType) GetSz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterItcmcrFieldSzMask) >> RegisterItcmcrFieldSzShift)
}

// SetSz Indicates the size of the relevant TCM
func (r *registerItcmcrType) SetSz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterItcmcrFieldSzMask)|(uint32(value)<<RegisterItcmcrFieldSzShift))
}

// registerDtcmcrType Enable access to the Tightly Coupled Memories (TCMs) by software running on the processor
type registerDtcmcrType uint32

const (
	RegisterDtcmcrFieldEnShift = 0
	RegisterDtcmcrFieldEnMask  = 0x1
)

// GetEn TCM enable
func (r *registerDtcmcrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtcmcrFieldEnMask) != 0
}

// SetEn TCM enable
func (r *registerDtcmcrType) SetEn(value bool) {
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
func (r *registerDtcmcrType) GetSz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDtcmcrFieldSzMask) >> RegisterDtcmcrFieldSzShift)
}

// SetSz Indicates the size of the relevant TCM
func (r *registerDtcmcrType) SetSz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtcmcrFieldSzMask)|(uint32(value)<<RegisterDtcmcrFieldSzShift))
}

// registerPahbcrType Enables accesses to Peripheral AHB (P-AHB) interface from software running on the processor
type registerPahbcrType uint32

const (
	RegisterPahbcrFieldEnShift = 0
	RegisterPahbcrFieldEnMask  = 0x1
)

// GetEn P-AHB enable
func (r *registerPahbcrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPahbcrFieldEnMask) != 0
}

// SetEn P-AHB enable
func (r *registerPahbcrType) SetEn(value bool) {
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
func (r *registerPahbcrType) GetSz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPahbcrFieldSzMask) >> RegisterPahbcrFieldSzShift)
}

// SetSz P-AHB size
func (r *registerPahbcrType) SetSz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPahbcrFieldSzMask)|(uint32(value)<<RegisterPahbcrFieldSzShift))
}
