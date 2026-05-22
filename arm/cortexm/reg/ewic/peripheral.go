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
	Cr      RegisterCrType
	Ascr    RegisterAscrType
	Clrmask RegisterClrmaskType
	Numid   RegisterNumidType
	_       [496]byte
	Maska   RegisterMaskaType
	Mask    [15]RegisterMaskType
	_       [448]byte
	Penda   RegisterPendaType
	Pend    [15]RegisterPendType
	_       [448]byte
	Psr     RegisterPsrType
}

// RegisterCrType The main External Wakeup Interrupt Controller (EWIC) control register
type RegisterCrType uint32

func (r *RegisterCrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrFieldEnShift = 0
	RegisterCrFieldEnMask  = 0x1
)

// GetEn Enable
func (r *RegisterCrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEnMask) != 0
}

// SetEn Enable
func (r *RegisterCrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEnMask)
	}
}

// RegisterAscrType Determines whether the processor generates APB transactions on entry and exit from Wakeup Interrupt Controller (WIC) sleep to set up the wakeup state in the External Wakeup Interrupt Controller (EWIC)
type RegisterAscrType uint32

func (r *RegisterAscrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAscrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAscrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAscrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAscrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAscrFieldAspdShift = 0
	RegisterAscrFieldAspdMask  = 0x1
)

// GetAspd The processor must use this value to decide whether any automatic EWIC accesses must be performed on transitioning to a low-power state
func (r *RegisterAscrType) GetAspd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAscrFieldAspdMask) != 0
}

// SetAspd The processor must use this value to decide whether any automatic EWIC accesses must be performed on transitioning to a low-power state
func (r *RegisterAscrType) SetAspd(value bool) {
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
func (r *RegisterAscrType) GetAspu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAscrFieldAspuMask) != 0
}

// SetAspu The processor must use this value to decide whether any automatic EWIC accesses must be performed on transitioning from a low-power state
func (r *RegisterAscrType) SetAspu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAscrFieldAspuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAscrFieldAspuMask)
	}
}

// RegisterClrmaskType Causes EWIC_MASKA and all the EWIC_MASKn registers to be cleared. The write data is ignored. This register is RAZ.
type RegisterClrmaskType uint32

func (r *RegisterClrmaskType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterClrmaskType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterClrmaskType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterClrmaskType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterClrmaskType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterNumidType Returns the total number of events supported in the External Wakeup Interrupt Controller (EWIC) that have been configured during synthesis
type RegisterNumidType uint32

func (r *RegisterNumidType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterNumidType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterNumidType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterNumidType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterNumidType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterNumidFieldNumeventShift = 0
	RegisterNumidFieldNumeventMask  = 0xffff
)

// GetNumevent The number of events supported
func (r *RegisterNumidType) GetNumevent() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterNumidFieldNumeventMask) >> RegisterNumidFieldNumeventShift)
}

// SetNumevent The number of events supported
func (r *RegisterNumidType) SetNumevent(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterNumidFieldNumeventMask)|(uint32(value)<<RegisterNumidFieldNumeventShift))
}

// RegisterMaskaType Defines the mask for special events and the EWIC_MASKn registers for external interrupt (IRQ) events
type RegisterMaskaType uint32

func (r *RegisterMaskaType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaskaType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaskaType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaskaType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaskaType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaskaFieldEventShift = 0
	RegisterMaskaFieldEventMask  = 0x1
)

// GetEvent Mask for Wait For Exception (WFE) wakeup event
func (r *RegisterMaskaType) GetEvent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskaFieldEventMask) != 0
}

// SetEvent Mask for Wait For Exception (WFE) wakeup event
func (r *RegisterMaskaType) SetEvent(value bool) {
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
func (r *RegisterMaskaType) GetNmi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskaFieldNmiMask) != 0
}

// SetNmi Mask for Non-Maskable Interrupt (NMI)
func (r *RegisterMaskaType) SetNmi(value bool) {
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
func (r *RegisterMaskaType) GetEdbgreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskaFieldEdbgreqMask) != 0
}

// SetEdbgreq Mask for external debug request
func (r *RegisterMaskaType) SetEdbgreq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskaFieldEdbgreqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskaFieldEdbgreqMask)
	}
}

// RegisterMaskType Defines the mask for special events and the EWIC_MASKn registers for external interrupt (IRQ) events
type RegisterMaskType uint32

func (r *RegisterMaskType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaskType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaskType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaskType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaskType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaskFieldIrqShift = 0
	RegisterMaskFieldIrqMask  = 0xffffffff
)

// GetIrq Masks for interrupts (nx32) to ((n+1)x32)-1
func (r *RegisterMaskType) GetIrq() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMaskFieldIrqMask) >> RegisterMaskFieldIrqShift)
}

// SetIrq Masks for interrupts (nx32) to ((n+1)x32)-1
func (r *RegisterMaskType) SetIrq(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaskFieldIrqMask)|(uint32(value)<<RegisterMaskFieldIrqShift))
}

// RegisterPendaType Indicate which events have been pended
type RegisterPendaType uint32

func (r *RegisterPendaType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPendaType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPendaType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPendaType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPendaType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPendaFieldEventShift = 0
	RegisterPendaFieldEventMask  = 0x1
)

// GetEvent Wait For Exception (WFE) wakeup event is pended
func (r *RegisterPendaType) GetEvent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPendaFieldEventMask) != 0
}

// SetEvent Wait For Exception (WFE) wakeup event is pended
func (r *RegisterPendaType) SetEvent(value bool) {
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
func (r *RegisterPendaType) GetNmi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPendaFieldNmiMask) != 0
}

// SetNmi Non-Maskable Interrupt (NMI) is pended
func (r *RegisterPendaType) SetNmi(value bool) {
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
func (r *RegisterPendaType) GetEdbgreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPendaFieldEdbgreqMask) != 0
}

// SetEdbgreq External debug request is pended
func (r *RegisterPendaType) SetEdbgreq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPendaFieldEdbgreqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPendaFieldEdbgreqMask)
	}
}

// RegisterPendType Indicate which events have been pended
type RegisterPendType uint32

func (r *RegisterPendType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPendType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPendType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPendType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPendType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPendFieldIrqShift = 0
	RegisterPendFieldIrqMask  = 0xffffffff
)

// GetIrq Interrupts (nx32) to ((n+1)x32)-1 are pended. A write of zero to this field is ignored.
func (r *RegisterPendType) GetIrq() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterPendFieldIrqMask) >> RegisterPendFieldIrqShift)
}

// SetIrq Interrupts (nx32) to ((n+1)x32)-1 are pended. A write of zero to this field is ignored.
func (r *RegisterPendType) SetIrq(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPendFieldIrqMask)|(uint32(value)<<RegisterPendFieldIrqShift))
}

// RegisterPsrType Indicates which EWIC_PENDn registers are non-zero
type RegisterPsrType uint32

func (r *RegisterPsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPsrFieldNzaShift = 0
	RegisterPsrFieldNzaMask  = 0x1
)

// GetNza If EWIC_PSR.NZA set, then EWIC_PENDA is non-zero
func (r *RegisterPsrType) GetNza() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPsrFieldNzaMask) != 0
}

// SetNza If EWIC_PSR.NZA set, then EWIC_PENDA is non-zero
func (r *RegisterPsrType) SetNza(value bool) {
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
func (r *RegisterPsrType) GetNz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPsrFieldNzMask) >> RegisterPsrFieldNzShift)
}

// SetNz If EWIC_PSR.NZ[n+1] is set, then EWIC_PENDn is non-zero
func (r *RegisterPsrType) SetNz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPsrFieldNzMask)|(uint32(value)<<RegisterPsrFieldNzShift))
}
