//go:build arm && cortexm

package ewic

import (
	"unsafe"
	"volatile"
)

var (
	Ewic = (*_ewic)(unsafe.Pointer(uintptr(0xe0047000)))
)

type _ewic struct {
	Cr      registerCrType
	Ascr    registerAscrType
	Clrmask registerClrmaskType
	Numid   registerNumidType
	_       [496]byte
	Maska   registerMaskaType
	Mask    [15]registerMaskType
	_       [448]byte
	Penda   registerPendaType
	Pend    [15]registerPendType
	_       [448]byte
	Psr     registerPsrType
}

// registerCrType The main External Wakeup Interrupt Controller (EWIC) control register
type registerCrType uint32

const (
	RegisterCrFieldEnShift = 0
	RegisterCrFieldEnMask  = 0x1
)

// GetEn Enable
func (r *registerCrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEnMask) != 0
}

// SetEn Enable
func (r *registerCrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEnMask)
	}
}

// registerAscrType Determines whether the processor generates APB transactions on entry and exit from Wakeup Interrupt Controller (WIC) sleep to set up the wakeup state in the External Wakeup Interrupt Controller (EWIC)
type registerAscrType uint32

const (
	RegisterAscrFieldAspdShift = 0
	RegisterAscrFieldAspdMask  = 0x1
)

// GetAspd The processor must use this value to decide whether any automatic EWIC accesses must be performed on transitioning to a low-power state
func (r *registerAscrType) GetAspd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAscrFieldAspdMask) != 0
}

// SetAspd The processor must use this value to decide whether any automatic EWIC accesses must be performed on transitioning to a low-power state
func (r *registerAscrType) SetAspd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAscrFieldAspdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAscrFieldAspdMask)
	}
}

const (
	RegisterAscrFieldAspuShift = 1
	RegisterAscrFieldAspuMask  = 0x2
)

// GetAspu The processor must use this value to decide whether any automatic EWIC accesses must be performed on transitioning from a low-power state
func (r *registerAscrType) GetAspu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAscrFieldAspuMask) != 0
}

// SetAspu The processor must use this value to decide whether any automatic EWIC accesses must be performed on transitioning from a low-power state
func (r *registerAscrType) SetAspu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAscrFieldAspuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAscrFieldAspuMask)
	}
}

// registerClrmaskType Causes EWIC_MASKA and all the EWIC_MASKn registers to be cleared. The write data is ignored. This register is RAZ.
type registerClrmaskType uint32

// registerNumidType Returns the total number of events supported in the External Wakeup Interrupt Controller (EWIC) that have been configured during synthesis
type registerNumidType uint32

const (
	RegisterNumidFieldNumeventShift = 0
	RegisterNumidFieldNumeventMask  = 0xffff
)

// GetNumevent The number of events supported
func (r *registerNumidType) GetNumevent() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterNumidFieldNumeventMask) >> RegisterNumidFieldNumeventShift)
}

// SetNumevent The number of events supported
func (r *registerNumidType) SetNumevent(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterNumidFieldNumeventMask)|(uint32(value)<<RegisterNumidFieldNumeventShift))
}

// registerMaskaType Defines the mask for special events and the EWIC_MASKn registers for external interrupt (IRQ) events
type registerMaskaType uint32

const (
	RegisterMaskaFieldEventShift = 0
	RegisterMaskaFieldEventMask  = 0x1
)

// GetEvent Mask for Wait For Exception (WFE) wakeup event
func (r *registerMaskaType) GetEvent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskaFieldEventMask) != 0
}

// SetEvent Mask for Wait For Exception (WFE) wakeup event
func (r *registerMaskaType) SetEvent(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskaFieldEventMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskaFieldEventMask)
	}
}

const (
	RegisterMaskaFieldNmiShift = 1
	RegisterMaskaFieldNmiMask  = 0x2
)

// GetNmi Mask for Non-Maskable Interrupt (NMI)
func (r *registerMaskaType) GetNmi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskaFieldNmiMask) != 0
}

// SetNmi Mask for Non-Maskable Interrupt (NMI)
func (r *registerMaskaType) SetNmi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskaFieldNmiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskaFieldNmiMask)
	}
}

const (
	RegisterMaskaFieldEdbgreqShift = 2
	RegisterMaskaFieldEdbgreqMask  = 0x4
)

// GetEdbgreq Mask for external debug request
func (r *registerMaskaType) GetEdbgreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskaFieldEdbgreqMask) != 0
}

// SetEdbgreq Mask for external debug request
func (r *registerMaskaType) SetEdbgreq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskaFieldEdbgreqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskaFieldEdbgreqMask)
	}
}

// registerMaskType Defines the mask for special events and the EWIC_MASKn registers for external interrupt (IRQ) events
type registerMaskType uint32

const (
	RegisterMaskFieldIrqShift = 0
	RegisterMaskFieldIrqMask  = 0xffffffff
)

// GetIrq Masks for interrupts (nx32) to ((n+1)x32)-1
func (r *registerMaskType) GetIrq() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMaskFieldIrqMask) >> RegisterMaskFieldIrqShift)
}

// SetIrq Masks for interrupts (nx32) to ((n+1)x32)-1
func (r *registerMaskType) SetIrq(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaskFieldIrqMask)|(uint32(value)<<RegisterMaskFieldIrqShift))
}

// registerPendaType Indicate which events have been pended
type registerPendaType uint32

const (
	RegisterPendaFieldEventShift = 0
	RegisterPendaFieldEventMask  = 0x1
)

// GetEvent Wait For Exception (WFE) wakeup event is pended
func (r *registerPendaType) GetEvent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPendaFieldEventMask) != 0
}

// SetEvent Wait For Exception (WFE) wakeup event is pended
func (r *registerPendaType) SetEvent(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPendaFieldEventMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPendaFieldEventMask)
	}
}

const (
	RegisterPendaFieldNmiShift = 1
	RegisterPendaFieldNmiMask  = 0x2
)

// GetNmi Non-Maskable Interrupt (NMI) is pended
func (r *registerPendaType) GetNmi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPendaFieldNmiMask) != 0
}

// SetNmi Non-Maskable Interrupt (NMI) is pended
func (r *registerPendaType) SetNmi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPendaFieldNmiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPendaFieldNmiMask)
	}
}

const (
	RegisterPendaFieldEdbgreqShift = 2
	RegisterPendaFieldEdbgreqMask  = 0x4
)

// GetEdbgreq External debug request is pended
func (r *registerPendaType) GetEdbgreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPendaFieldEdbgreqMask) != 0
}

// SetEdbgreq External debug request is pended
func (r *registerPendaType) SetEdbgreq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPendaFieldEdbgreqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPendaFieldEdbgreqMask)
	}
}

// registerPendType Indicate which events have been pended
type registerPendType uint32

const (
	RegisterPendFieldIrqShift = 0
	RegisterPendFieldIrqMask  = 0xffffffff
)

// GetIrq Interrupts (nx32) to ((n+1)x32)-1 are pended. A write of zero to this field is ignored.
func (r *registerPendType) GetIrq() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterPendFieldIrqMask) >> RegisterPendFieldIrqShift)
}

// SetIrq Interrupts (nx32) to ((n+1)x32)-1 are pended. A write of zero to this field is ignored.
func (r *registerPendType) SetIrq(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPendFieldIrqMask)|(uint32(value)<<RegisterPendFieldIrqShift))
}

// registerPsrType Indicates which EWIC_PENDn registers are non-zero
type registerPsrType uint32

const (
	RegisterPsrFieldNzaShift = 0
	RegisterPsrFieldNzaMask  = 0x1
)

// GetNza If EWIC_PSR.NZA set, then EWIC_PENDA is non-zero
func (r *registerPsrType) GetNza() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPsrFieldNzaMask) != 0
}

// SetNza If EWIC_PSR.NZA set, then EWIC_PENDA is non-zero
func (r *registerPsrType) SetNza(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPsrFieldNzaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPsrFieldNzaMask)
	}
}

const (
	RegisterPsrFieldNzShift = 1
	RegisterPsrFieldNzMask  = 0xfffe
)

// GetNz If EWIC_PSR.NZ[n+1] is set, then EWIC_PENDn is non-zero
func (r *registerPsrType) GetNz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPsrFieldNzMask) >> RegisterPsrFieldNzShift)
}

// SetNz If EWIC_PSR.NZ[n+1] is set, then EWIC_PENDn is non-zero
func (r *registerPsrType) SetNz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPsrFieldNzMask)|(uint32(value)<<RegisterPsrFieldNzShift))
}
