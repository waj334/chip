//go:build arm && cortexm

package ewicis

import (
	"unsafe"
	"volatile"
)

var (
	Ewicis = (*_ewicis)(unsafe.Pointer(uintptr(0xe001e400)))
)

type _ewicis struct {
	Eventspr   RegisterEventsprType
	_          [124]byte
	Eventmaska RegisterEventmaskaType
	Eventmask  [32]RegisterEventmaskType
}

// RegisterEventsprType Set pending events at wakeup that cannot be directly set in the Nested Vectored Interrupt Controller (NVIC) using the architecture programming model
type RegisterEventsprType uint32

func (r *RegisterEventsprType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterEventsprType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterEventsprType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterEventsprType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterEventsprType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterEventsprFieldNmiShift = 1
	RegisterEventsprFieldNmiMask  = 0x2
)

// GetNmi Causes the processor to behave like a Non-maskable Interrupt (NMI) has occurred
func (r *RegisterEventsprType) GetNmi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEventsprFieldNmiMask) != 0
}

// SetNmi Causes the processor to behave like a Non-maskable Interrupt (NMI) has occurred
func (r *RegisterEventsprType) SetNmi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEventsprFieldNmiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEventsprFieldNmiMask)
	}
}

const (
	RegisterEventsprFieldEventShift = 1
	RegisterEventsprFieldEventMask  = 0x2
)

// GetEvent Causes the processor to behave like an RXEV event has occurred
func (r *RegisterEventsprType) GetEvent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEventsprFieldEventMask) != 0
}

// SetEvent Causes the processor to behave like an RXEV event has occurred
func (r *RegisterEventsprType) SetEvent(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEventsprFieldEventMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEventsprFieldEventMask)
	}
}

const (
	RegisterEventsprFieldEdbgreqShift = 2
	RegisterEventsprFieldEdbgreqMask  = 0x4
)

// GetEdbgreq Causes the processor to behave like an external debug request has occurred
func (r *RegisterEventsprType) GetEdbgreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEventsprFieldEdbgreqMask) != 0
}

// SetEdbgreq Causes the processor to behave like an external debug request has occurred
func (r *RegisterEventsprType) SetEdbgreq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEventsprFieldEdbgreqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEventsprFieldEdbgreqMask)
	}
}

// RegisterEventmaskaType Provides the events on sleep entry which cause the processor to wake up, including information about internal events
type RegisterEventmaskaType uint32

func (r *RegisterEventmaskaType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterEventmaskaType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterEventmaskaType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterEventmaskaType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterEventmaskaType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterEventmaskaFieldEventShift = 0
	RegisterEventmaskaFieldEventMask  = 0x1
)

// GetEvent Sensitive to RXEV when in WFE sleep
func (r *RegisterEventmaskaType) GetEvent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEventmaskaFieldEventMask) != 0
}

// SetEvent Sensitive to RXEV when in WFE sleep
func (r *RegisterEventmaskaType) SetEvent(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEventmaskaFieldEventMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEventmaskaFieldEventMask)
	}
}

const (
	RegisterEventmaskaFieldNmiShift = 1
	RegisterEventmaskaFieldNmiMask  = 0x2
)

// GetNmi Mask for Non-Maskable Interrupt (NMI)
func (r *RegisterEventmaskaType) GetNmi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEventmaskaFieldNmiMask) != 0
}

// SetNmi Mask for Non-Maskable Interrupt (NMI)
func (r *RegisterEventmaskaType) SetNmi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEventmaskaFieldNmiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEventmaskaFieldNmiMask)
	}
}

const (
	RegisterEventmaskaFieldEdbgreqShift = 2
	RegisterEventmaskaFieldEdbgreqMask  = 0x4
)

// GetEdbgreq Mask for external debug request
func (r *RegisterEventmaskaType) GetEdbgreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEventmaskaFieldEdbgreqMask) != 0
}

// SetEdbgreq Mask for external debug request
func (r *RegisterEventmaskaType) SetEdbgreq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEventmaskaFieldEdbgreqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEventmaskaFieldEdbgreqMask)
	}
}

// RegisterEventmaskType Provides the events on sleep entry which cause the processor to wake up, including information about external events
type RegisterEventmaskType uint32

func (r *RegisterEventmaskType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterEventmaskType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterEventmaskType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterEventmaskType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterEventmaskType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterEventmaskFieldIrqShift = 0
	RegisterEventmaskFieldIrqMask  = 0xffffffff
)

// GetIrq Masks for interrupts ((n-1)x32) to (nx32)-1.
func (r *RegisterEventmaskType) GetIrq() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterEventmaskFieldIrqMask) >> RegisterEventmaskFieldIrqShift)
}

// SetIrq Masks for interrupts ((n-1)x32) to (nx32)-1.
func (r *RegisterEventmaskType) SetIrq(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEventmaskFieldIrqMask)|(uint32(value)<<RegisterEventmaskFieldIrqShift))
}
