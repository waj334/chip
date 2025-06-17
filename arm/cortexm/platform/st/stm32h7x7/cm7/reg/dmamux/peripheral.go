//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package dmamux

import (
	"unsafe"
	"volatile"
)

var (
	Dmamux1 = (*_dmamux)(unsafe.Pointer(uintptr(0x40020800)))
	Dmamux2 = (*_dmamux)(unsafe.Pointer(uintptr(0x58025800)))
)

type _dmamux struct {
	C0cr  RegisterC0crType
	C1cr  RegisterC1crType
	C2cr  RegisterC2crType
	C3cr  RegisterC3crType
	C4cr  RegisterC4crType
	C5cr  RegisterC5crType
	C6cr  RegisterC6crType
	C7cr  RegisterC7crType
	C8cr  RegisterC8crType
	C9cr  RegisterC9crType
	C10cr RegisterC10crType
	C11cr RegisterC11crType
	C12cr RegisterC12crType
	C13cr RegisterC13crType
	C14cr RegisterC14crType
	C15cr RegisterC15crType
	_     [64]byte
	Csr   RegisterCsrType
	Cfr   RegisterCfrType
	_     [120]byte
	Rg0cr RegisterRg0crType
	Rg1cr RegisterRg1crType
	Rg2cr RegisterRg2crType
	Rg3cr RegisterRg3crType
	Rg4cr RegisterRg4crType
	Rg5cr RegisterRg5crType
	Rg6cr RegisterRg6crType
	Rg7cr RegisterRg7crType
	_     [32]byte
	Rgsr  RegisterRgsrType
	Rgcfr RegisterRgcfrType
}

// RegisterC0crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC0crType uint32

func (r *RegisterC0crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC0crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC0crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC0crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC0crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC0crFieldDmareqidShift = 0
	RegisterC0crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC0crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldDmareqidMask) >> RegisterC0crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC0crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC0crFieldDmareqidMask)|(uint32(value)<<RegisterC0crFieldDmareqidShift))
}

const (
	RegisterC0crFieldSoieShift = 8
	RegisterC0crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC0crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC0crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC0crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC0crFieldSoieMask)
	}
}

const (
	RegisterC0crFieldEgeShift = 9
	RegisterC0crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC0crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC0crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC0crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC0crFieldEgeMask)
	}
}

const (
	RegisterC0crFieldSeShift = 16
	RegisterC0crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC0crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC0crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC0crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC0crFieldSeMask)
	}
}

const (
	RegisterC0crFieldSpolShift = 17
	RegisterC0crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC0crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldSpolMask) >> RegisterC0crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC0crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC0crFieldSpolMask)|(uint32(value)<<RegisterC0crFieldSpolShift))
}

const (
	RegisterC0crFieldNbreqShift = 19
	RegisterC0crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC0crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldNbreqMask) >> RegisterC0crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC0crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC0crFieldNbreqMask)|(uint32(value)<<RegisterC0crFieldNbreqShift))
}

const (
	RegisterC0crFieldSyncidShift = 24
	RegisterC0crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC0crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldSyncidMask) >> RegisterC0crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC0crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC0crFieldSyncidMask)|(uint32(value)<<RegisterC0crFieldSyncidShift))
}

// RegisterC1crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC1crType uint32

func (r *RegisterC1crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1crFieldDmareqidShift = 0
	RegisterC1crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC1crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldDmareqidMask) >> RegisterC1crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC1crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC1crFieldDmareqidMask)|(uint32(value)<<RegisterC1crFieldDmareqidShift))
}

const (
	RegisterC1crFieldSoieShift = 8
	RegisterC1crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC1crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC1crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1crFieldSoieMask)
	}
}

const (
	RegisterC1crFieldEgeShift = 9
	RegisterC1crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC1crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC1crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1crFieldEgeMask)
	}
}

const (
	RegisterC1crFieldSeShift = 16
	RegisterC1crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC1crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC1crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1crFieldSeMask)
	}
}

const (
	RegisterC1crFieldSpolShift = 17
	RegisterC1crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC1crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldSpolMask) >> RegisterC1crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC1crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC1crFieldSpolMask)|(uint32(value)<<RegisterC1crFieldSpolShift))
}

const (
	RegisterC1crFieldNbreqShift = 19
	RegisterC1crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC1crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldNbreqMask) >> RegisterC1crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC1crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC1crFieldNbreqMask)|(uint32(value)<<RegisterC1crFieldNbreqShift))
}

const (
	RegisterC1crFieldSyncidShift = 24
	RegisterC1crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC1crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldSyncidMask) >> RegisterC1crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC1crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC1crFieldSyncidMask)|(uint32(value)<<RegisterC1crFieldSyncidShift))
}

// RegisterC2crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC2crType uint32

func (r *RegisterC2crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC2crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC2crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC2crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC2crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC2crFieldDmareqidShift = 0
	RegisterC2crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC2crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldDmareqidMask) >> RegisterC2crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC2crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC2crFieldDmareqidMask)|(uint32(value)<<RegisterC2crFieldDmareqidShift))
}

const (
	RegisterC2crFieldSoieShift = 8
	RegisterC2crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC2crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC2crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC2crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC2crFieldSoieMask)
	}
}

const (
	RegisterC2crFieldEgeShift = 9
	RegisterC2crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC2crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC2crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC2crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC2crFieldEgeMask)
	}
}

const (
	RegisterC2crFieldSeShift = 16
	RegisterC2crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC2crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC2crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC2crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC2crFieldSeMask)
	}
}

const (
	RegisterC2crFieldSpolShift = 17
	RegisterC2crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC2crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldSpolMask) >> RegisterC2crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC2crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC2crFieldSpolMask)|(uint32(value)<<RegisterC2crFieldSpolShift))
}

const (
	RegisterC2crFieldNbreqShift = 19
	RegisterC2crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC2crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldNbreqMask) >> RegisterC2crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC2crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC2crFieldNbreqMask)|(uint32(value)<<RegisterC2crFieldNbreqShift))
}

const (
	RegisterC2crFieldSyncidShift = 24
	RegisterC2crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC2crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldSyncidMask) >> RegisterC2crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC2crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC2crFieldSyncidMask)|(uint32(value)<<RegisterC2crFieldSyncidShift))
}

// RegisterC3crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC3crType uint32

func (r *RegisterC3crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC3crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC3crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC3crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC3crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC3crFieldDmareqidShift = 0
	RegisterC3crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC3crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldDmareqidMask) >> RegisterC3crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC3crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC3crFieldDmareqidMask)|(uint32(value)<<RegisterC3crFieldDmareqidShift))
}

const (
	RegisterC3crFieldSoieShift = 8
	RegisterC3crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC3crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC3crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC3crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC3crFieldSoieMask)
	}
}

const (
	RegisterC3crFieldEgeShift = 9
	RegisterC3crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC3crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC3crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC3crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC3crFieldEgeMask)
	}
}

const (
	RegisterC3crFieldSeShift = 16
	RegisterC3crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC3crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC3crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC3crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC3crFieldSeMask)
	}
}

const (
	RegisterC3crFieldSpolShift = 17
	RegisterC3crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC3crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldSpolMask) >> RegisterC3crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC3crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC3crFieldSpolMask)|(uint32(value)<<RegisterC3crFieldSpolShift))
}

const (
	RegisterC3crFieldNbreqShift = 19
	RegisterC3crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC3crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldNbreqMask) >> RegisterC3crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC3crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC3crFieldNbreqMask)|(uint32(value)<<RegisterC3crFieldNbreqShift))
}

const (
	RegisterC3crFieldSyncidShift = 24
	RegisterC3crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC3crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldSyncidMask) >> RegisterC3crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC3crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC3crFieldSyncidMask)|(uint32(value)<<RegisterC3crFieldSyncidShift))
}

// RegisterC4crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC4crType uint32

func (r *RegisterC4crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC4crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC4crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC4crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC4crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC4crFieldDmareqidShift = 0
	RegisterC4crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC4crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldDmareqidMask) >> RegisterC4crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC4crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC4crFieldDmareqidMask)|(uint32(value)<<RegisterC4crFieldDmareqidShift))
}

const (
	RegisterC4crFieldSoieShift = 8
	RegisterC4crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC4crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC4crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC4crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC4crFieldSoieMask)
	}
}

const (
	RegisterC4crFieldEgeShift = 9
	RegisterC4crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC4crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC4crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC4crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC4crFieldEgeMask)
	}
}

const (
	RegisterC4crFieldSeShift = 16
	RegisterC4crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC4crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC4crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC4crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC4crFieldSeMask)
	}
}

const (
	RegisterC4crFieldSpolShift = 17
	RegisterC4crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC4crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldSpolMask) >> RegisterC4crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC4crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC4crFieldSpolMask)|(uint32(value)<<RegisterC4crFieldSpolShift))
}

const (
	RegisterC4crFieldNbreqShift = 19
	RegisterC4crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC4crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldNbreqMask) >> RegisterC4crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC4crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC4crFieldNbreqMask)|(uint32(value)<<RegisterC4crFieldNbreqShift))
}

const (
	RegisterC4crFieldSyncidShift = 24
	RegisterC4crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC4crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldSyncidMask) >> RegisterC4crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC4crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC4crFieldSyncidMask)|(uint32(value)<<RegisterC4crFieldSyncidShift))
}

// RegisterC5crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC5crType uint32

func (r *RegisterC5crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC5crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC5crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC5crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC5crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC5crFieldDmareqidShift = 0
	RegisterC5crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC5crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldDmareqidMask) >> RegisterC5crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC5crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC5crFieldDmareqidMask)|(uint32(value)<<RegisterC5crFieldDmareqidShift))
}

const (
	RegisterC5crFieldSoieShift = 8
	RegisterC5crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC5crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC5crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC5crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC5crFieldSoieMask)
	}
}

const (
	RegisterC5crFieldEgeShift = 9
	RegisterC5crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC5crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC5crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC5crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC5crFieldEgeMask)
	}
}

const (
	RegisterC5crFieldSeShift = 16
	RegisterC5crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC5crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC5crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC5crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC5crFieldSeMask)
	}
}

const (
	RegisterC5crFieldSpolShift = 17
	RegisterC5crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC5crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldSpolMask) >> RegisterC5crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC5crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC5crFieldSpolMask)|(uint32(value)<<RegisterC5crFieldSpolShift))
}

const (
	RegisterC5crFieldNbreqShift = 19
	RegisterC5crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC5crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldNbreqMask) >> RegisterC5crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC5crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC5crFieldNbreqMask)|(uint32(value)<<RegisterC5crFieldNbreqShift))
}

const (
	RegisterC5crFieldSyncidShift = 24
	RegisterC5crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC5crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldSyncidMask) >> RegisterC5crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC5crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC5crFieldSyncidMask)|(uint32(value)<<RegisterC5crFieldSyncidShift))
}

// RegisterC6crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC6crType uint32

func (r *RegisterC6crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC6crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC6crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC6crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC6crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC6crFieldDmareqidShift = 0
	RegisterC6crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC6crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldDmareqidMask) >> RegisterC6crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC6crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC6crFieldDmareqidMask)|(uint32(value)<<RegisterC6crFieldDmareqidShift))
}

const (
	RegisterC6crFieldSoieShift = 8
	RegisterC6crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC6crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC6crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC6crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC6crFieldSoieMask)
	}
}

const (
	RegisterC6crFieldEgeShift = 9
	RegisterC6crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC6crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC6crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC6crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC6crFieldEgeMask)
	}
}

const (
	RegisterC6crFieldSeShift = 16
	RegisterC6crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC6crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC6crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC6crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC6crFieldSeMask)
	}
}

const (
	RegisterC6crFieldSpolShift = 17
	RegisterC6crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC6crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldSpolMask) >> RegisterC6crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC6crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC6crFieldSpolMask)|(uint32(value)<<RegisterC6crFieldSpolShift))
}

const (
	RegisterC6crFieldNbreqShift = 19
	RegisterC6crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC6crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldNbreqMask) >> RegisterC6crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC6crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC6crFieldNbreqMask)|(uint32(value)<<RegisterC6crFieldNbreqShift))
}

const (
	RegisterC6crFieldSyncidShift = 24
	RegisterC6crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC6crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldSyncidMask) >> RegisterC6crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC6crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC6crFieldSyncidMask)|(uint32(value)<<RegisterC6crFieldSyncidShift))
}

// RegisterC7crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC7crType uint32

func (r *RegisterC7crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC7crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC7crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC7crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC7crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC7crFieldDmareqidShift = 0
	RegisterC7crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC7crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldDmareqidMask) >> RegisterC7crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC7crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC7crFieldDmareqidMask)|(uint32(value)<<RegisterC7crFieldDmareqidShift))
}

const (
	RegisterC7crFieldSoieShift = 8
	RegisterC7crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC7crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC7crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC7crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC7crFieldSoieMask)
	}
}

const (
	RegisterC7crFieldEgeShift = 9
	RegisterC7crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC7crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC7crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC7crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC7crFieldEgeMask)
	}
}

const (
	RegisterC7crFieldSeShift = 16
	RegisterC7crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC7crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC7crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC7crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC7crFieldSeMask)
	}
}

const (
	RegisterC7crFieldSpolShift = 17
	RegisterC7crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC7crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldSpolMask) >> RegisterC7crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC7crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC7crFieldSpolMask)|(uint32(value)<<RegisterC7crFieldSpolShift))
}

const (
	RegisterC7crFieldNbreqShift = 19
	RegisterC7crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC7crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldNbreqMask) >> RegisterC7crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC7crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC7crFieldNbreqMask)|(uint32(value)<<RegisterC7crFieldNbreqShift))
}

const (
	RegisterC7crFieldSyncidShift = 24
	RegisterC7crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC7crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldSyncidMask) >> RegisterC7crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC7crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC7crFieldSyncidMask)|(uint32(value)<<RegisterC7crFieldSyncidShift))
}

// RegisterC8crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC8crType uint32

func (r *RegisterC8crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC8crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC8crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC8crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC8crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC8crFieldDmareqidShift = 0
	RegisterC8crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC8crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldDmareqidMask) >> RegisterC8crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC8crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC8crFieldDmareqidMask)|(uint32(value)<<RegisterC8crFieldDmareqidShift))
}

const (
	RegisterC8crFieldSoieShift = 8
	RegisterC8crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC8crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC8crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC8crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC8crFieldSoieMask)
	}
}

const (
	RegisterC8crFieldEgeShift = 9
	RegisterC8crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC8crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC8crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC8crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC8crFieldEgeMask)
	}
}

const (
	RegisterC8crFieldSeShift = 16
	RegisterC8crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC8crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC8crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC8crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC8crFieldSeMask)
	}
}

const (
	RegisterC8crFieldSpolShift = 17
	RegisterC8crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC8crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldSpolMask) >> RegisterC8crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC8crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC8crFieldSpolMask)|(uint32(value)<<RegisterC8crFieldSpolShift))
}

const (
	RegisterC8crFieldNbreqShift = 19
	RegisterC8crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC8crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldNbreqMask) >> RegisterC8crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC8crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC8crFieldNbreqMask)|(uint32(value)<<RegisterC8crFieldNbreqShift))
}

const (
	RegisterC8crFieldSyncidShift = 24
	RegisterC8crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC8crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldSyncidMask) >> RegisterC8crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC8crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC8crFieldSyncidMask)|(uint32(value)<<RegisterC8crFieldSyncidShift))
}

// RegisterC9crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC9crType uint32

func (r *RegisterC9crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC9crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC9crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC9crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC9crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC9crFieldDmareqidShift = 0
	RegisterC9crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC9crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldDmareqidMask) >> RegisterC9crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC9crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC9crFieldDmareqidMask)|(uint32(value)<<RegisterC9crFieldDmareqidShift))
}

const (
	RegisterC9crFieldSoieShift = 8
	RegisterC9crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC9crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC9crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC9crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC9crFieldSoieMask)
	}
}

const (
	RegisterC9crFieldEgeShift = 9
	RegisterC9crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC9crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC9crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC9crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC9crFieldEgeMask)
	}
}

const (
	RegisterC9crFieldSeShift = 16
	RegisterC9crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC9crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC9crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC9crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC9crFieldSeMask)
	}
}

const (
	RegisterC9crFieldSpolShift = 17
	RegisterC9crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC9crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldSpolMask) >> RegisterC9crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC9crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC9crFieldSpolMask)|(uint32(value)<<RegisterC9crFieldSpolShift))
}

const (
	RegisterC9crFieldNbreqShift = 19
	RegisterC9crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC9crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldNbreqMask) >> RegisterC9crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC9crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC9crFieldNbreqMask)|(uint32(value)<<RegisterC9crFieldNbreqShift))
}

const (
	RegisterC9crFieldSyncidShift = 24
	RegisterC9crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC9crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldSyncidMask) >> RegisterC9crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC9crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC9crFieldSyncidMask)|(uint32(value)<<RegisterC9crFieldSyncidShift))
}

// RegisterC10crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC10crType uint32

func (r *RegisterC10crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC10crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC10crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC10crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC10crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC10crFieldDmareqidShift = 0
	RegisterC10crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC10crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldDmareqidMask) >> RegisterC10crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC10crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC10crFieldDmareqidMask)|(uint32(value)<<RegisterC10crFieldDmareqidShift))
}

const (
	RegisterC10crFieldSoieShift = 8
	RegisterC10crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC10crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC10crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC10crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC10crFieldSoieMask)
	}
}

const (
	RegisterC10crFieldEgeShift = 9
	RegisterC10crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC10crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC10crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC10crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC10crFieldEgeMask)
	}
}

const (
	RegisterC10crFieldSeShift = 16
	RegisterC10crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC10crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC10crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC10crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC10crFieldSeMask)
	}
}

const (
	RegisterC10crFieldSpolShift = 17
	RegisterC10crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC10crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldSpolMask) >> RegisterC10crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC10crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC10crFieldSpolMask)|(uint32(value)<<RegisterC10crFieldSpolShift))
}

const (
	RegisterC10crFieldNbreqShift = 19
	RegisterC10crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC10crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldNbreqMask) >> RegisterC10crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC10crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC10crFieldNbreqMask)|(uint32(value)<<RegisterC10crFieldNbreqShift))
}

const (
	RegisterC10crFieldSyncidShift = 24
	RegisterC10crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC10crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldSyncidMask) >> RegisterC10crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC10crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC10crFieldSyncidMask)|(uint32(value)<<RegisterC10crFieldSyncidShift))
}

// RegisterC11crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC11crType uint32

func (r *RegisterC11crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC11crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC11crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC11crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC11crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC11crFieldDmareqidShift = 0
	RegisterC11crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC11crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldDmareqidMask) >> RegisterC11crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC11crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC11crFieldDmareqidMask)|(uint32(value)<<RegisterC11crFieldDmareqidShift))
}

const (
	RegisterC11crFieldSoieShift = 8
	RegisterC11crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC11crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC11crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC11crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC11crFieldSoieMask)
	}
}

const (
	RegisterC11crFieldEgeShift = 9
	RegisterC11crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC11crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC11crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC11crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC11crFieldEgeMask)
	}
}

const (
	RegisterC11crFieldSeShift = 16
	RegisterC11crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC11crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC11crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC11crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC11crFieldSeMask)
	}
}

const (
	RegisterC11crFieldSpolShift = 17
	RegisterC11crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC11crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldSpolMask) >> RegisterC11crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC11crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC11crFieldSpolMask)|(uint32(value)<<RegisterC11crFieldSpolShift))
}

const (
	RegisterC11crFieldNbreqShift = 19
	RegisterC11crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC11crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldNbreqMask) >> RegisterC11crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC11crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC11crFieldNbreqMask)|(uint32(value)<<RegisterC11crFieldNbreqShift))
}

const (
	RegisterC11crFieldSyncidShift = 24
	RegisterC11crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC11crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldSyncidMask) >> RegisterC11crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC11crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC11crFieldSyncidMask)|(uint32(value)<<RegisterC11crFieldSyncidShift))
}

// RegisterC12crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC12crType uint32

func (r *RegisterC12crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC12crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC12crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC12crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC12crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC12crFieldDmareqidShift = 0
	RegisterC12crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC12crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldDmareqidMask) >> RegisterC12crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC12crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC12crFieldDmareqidMask)|(uint32(value)<<RegisterC12crFieldDmareqidShift))
}

const (
	RegisterC12crFieldSoieShift = 8
	RegisterC12crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC12crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC12crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC12crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC12crFieldSoieMask)
	}
}

const (
	RegisterC12crFieldEgeShift = 9
	RegisterC12crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC12crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC12crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC12crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC12crFieldEgeMask)
	}
}

const (
	RegisterC12crFieldSeShift = 16
	RegisterC12crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC12crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC12crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC12crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC12crFieldSeMask)
	}
}

const (
	RegisterC12crFieldSpolShift = 17
	RegisterC12crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC12crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldSpolMask) >> RegisterC12crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC12crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC12crFieldSpolMask)|(uint32(value)<<RegisterC12crFieldSpolShift))
}

const (
	RegisterC12crFieldNbreqShift = 19
	RegisterC12crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC12crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldNbreqMask) >> RegisterC12crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC12crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC12crFieldNbreqMask)|(uint32(value)<<RegisterC12crFieldNbreqShift))
}

const (
	RegisterC12crFieldSyncidShift = 24
	RegisterC12crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC12crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldSyncidMask) >> RegisterC12crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC12crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC12crFieldSyncidMask)|(uint32(value)<<RegisterC12crFieldSyncidShift))
}

// RegisterC13crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC13crType uint32

func (r *RegisterC13crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC13crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC13crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC13crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC13crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC13crFieldDmareqidShift = 0
	RegisterC13crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC13crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldDmareqidMask) >> RegisterC13crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC13crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC13crFieldDmareqidMask)|(uint32(value)<<RegisterC13crFieldDmareqidShift))
}

const (
	RegisterC13crFieldSoieShift = 8
	RegisterC13crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC13crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC13crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC13crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC13crFieldSoieMask)
	}
}

const (
	RegisterC13crFieldEgeShift = 9
	RegisterC13crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC13crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC13crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC13crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC13crFieldEgeMask)
	}
}

const (
	RegisterC13crFieldSeShift = 16
	RegisterC13crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC13crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC13crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC13crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC13crFieldSeMask)
	}
}

const (
	RegisterC13crFieldSpolShift = 17
	RegisterC13crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC13crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldSpolMask) >> RegisterC13crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC13crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC13crFieldSpolMask)|(uint32(value)<<RegisterC13crFieldSpolShift))
}

const (
	RegisterC13crFieldNbreqShift = 19
	RegisterC13crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC13crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldNbreqMask) >> RegisterC13crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC13crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC13crFieldNbreqMask)|(uint32(value)<<RegisterC13crFieldNbreqShift))
}

const (
	RegisterC13crFieldSyncidShift = 24
	RegisterC13crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC13crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldSyncidMask) >> RegisterC13crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC13crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC13crFieldSyncidMask)|(uint32(value)<<RegisterC13crFieldSyncidShift))
}

// RegisterC14crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC14crType uint32

func (r *RegisterC14crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC14crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC14crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC14crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC14crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC14crFieldDmareqidShift = 0
	RegisterC14crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC14crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldDmareqidMask) >> RegisterC14crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC14crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC14crFieldDmareqidMask)|(uint32(value)<<RegisterC14crFieldDmareqidShift))
}

const (
	RegisterC14crFieldSoieShift = 8
	RegisterC14crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC14crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC14crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC14crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC14crFieldSoieMask)
	}
}

const (
	RegisterC14crFieldEgeShift = 9
	RegisterC14crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC14crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC14crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC14crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC14crFieldEgeMask)
	}
}

const (
	RegisterC14crFieldSeShift = 16
	RegisterC14crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC14crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC14crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC14crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC14crFieldSeMask)
	}
}

const (
	RegisterC14crFieldSpolShift = 17
	RegisterC14crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC14crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldSpolMask) >> RegisterC14crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC14crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC14crFieldSpolMask)|(uint32(value)<<RegisterC14crFieldSpolShift))
}

const (
	RegisterC14crFieldNbreqShift = 19
	RegisterC14crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC14crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldNbreqMask) >> RegisterC14crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC14crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC14crFieldNbreqMask)|(uint32(value)<<RegisterC14crFieldNbreqShift))
}

const (
	RegisterC14crFieldSyncidShift = 24
	RegisterC14crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC14crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldSyncidMask) >> RegisterC14crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC14crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC14crFieldSyncidMask)|(uint32(value)<<RegisterC14crFieldSyncidShift))
}

// RegisterC15crType DMAMux - DMA request line multiplexer channel x control register
type RegisterC15crType uint32

func (r *RegisterC15crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC15crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC15crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC15crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC15crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC15crFieldDmareqidShift = 0
	RegisterC15crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *RegisterC15crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldDmareqidMask) >> RegisterC15crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *RegisterC15crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC15crFieldDmareqidMask)|(uint32(value)<<RegisterC15crFieldDmareqidShift))
}

const (
	RegisterC15crFieldSoieShift = 8
	RegisterC15crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC15crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *RegisterC15crType) SetSoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC15crFieldSoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC15crFieldSoieMask)
	}
}

const (
	RegisterC15crFieldEgeShift = 9
	RegisterC15crFieldEgeMask  = 0x200
)

// GetEge Event generation enable/disable
func (r *RegisterC15crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *RegisterC15crType) SetEge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC15crFieldEgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC15crFieldEgeMask)
	}
}

const (
	RegisterC15crFieldSeShift = 16
	RegisterC15crFieldSeMask  = 0x10000
)

// GetSe Synchronous operating mode enable/disable
func (r *RegisterC15crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *RegisterC15crType) SetSe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC15crFieldSeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC15crFieldSeMask)
	}
}

const (
	RegisterC15crFieldSpolShift = 17
	RegisterC15crFieldSpolMask  = 0x60000
)

// GetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC15crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldSpolMask) >> RegisterC15crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *RegisterC15crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC15crFieldSpolMask)|(uint32(value)<<RegisterC15crFieldSpolShift))
}

const (
	RegisterC15crFieldNbreqShift = 19
	RegisterC15crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC15crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldNbreqMask) >> RegisterC15crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *RegisterC15crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC15crFieldNbreqMask)|(uint32(value)<<RegisterC15crFieldNbreqShift))
}

const (
	RegisterC15crFieldSyncidShift = 24
	RegisterC15crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *RegisterC15crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldSyncidMask) >> RegisterC15crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *RegisterC15crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC15crFieldSyncidMask)|(uint32(value)<<RegisterC15crFieldSyncidShift))
}

// RegisterCsrType DMAMUX request line multiplexer interrupt channel status register
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
	RegisterCsrFieldSofShift = 0
	RegisterCsrFieldSofMask  = 0xffff
)

// GetSof Synchronization overrun event flag
func (r *RegisterCsrType) GetSof() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldSofMask) >> RegisterCsrFieldSofShift)
}

// SetSof Synchronization overrun event flag
func (r *RegisterCsrType) SetSof(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldSofMask)|(uint32(value)<<RegisterCsrFieldSofShift))
}

// RegisterCfrType DMAMUX request line multiplexer interrupt clear flag register
type RegisterCfrType uint32

func (r *RegisterCfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCfrFieldCsofShift = 0
	RegisterCfrFieldCsofMask  = 0xffff
)

// GetCsof Clear synchronization overrun event flag
func (r *RegisterCfrType) GetCsof() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCfrFieldCsofMask) >> RegisterCfrFieldCsofShift)
}

// SetCsof Clear synchronization overrun event flag
func (r *RegisterCfrType) SetCsof(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfrFieldCsofMask)|(uint32(value)<<RegisterCfrFieldCsofShift))
}

// RegisterRg0crType DMAMux - DMA request generator channel x control register
type RegisterRg0crType uint32

func (r *RegisterRg0crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRg0crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRg0crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRg0crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRg0crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRg0crFieldSigidShift = 0
	RegisterRg0crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *RegisterRg0crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg0crFieldSigidMask) >> RegisterRg0crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *RegisterRg0crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg0crFieldSigidMask)|(uint32(value)<<RegisterRg0crFieldSigidShift))
}

const (
	RegisterRg0crFieldOieShift = 8
	RegisterRg0crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *RegisterRg0crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg0crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *RegisterRg0crType) SetOie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg0crFieldOieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg0crFieldOieMask)
	}
}

const (
	RegisterRg0crFieldGeShift = 16
	RegisterRg0crFieldGeMask  = 0x10000
)

// GetGe DMA request generator channel enable/disable
func (r *RegisterRg0crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg0crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *RegisterRg0crType) SetGe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg0crFieldGeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg0crFieldGeMask)
	}
}

const (
	RegisterRg0crFieldGpolShift = 17
	RegisterRg0crFieldGpolMask  = 0x60000
)

// GetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg0crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg0crFieldGpolMask) >> RegisterRg0crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg0crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg0crFieldGpolMask)|(uint32(value)<<RegisterRg0crFieldGpolShift))
}

const (
	RegisterRg0crFieldGnbreqShift = 19
	RegisterRg0crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg0crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg0crFieldGnbreqMask) >> RegisterRg0crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg0crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg0crFieldGnbreqMask)|(uint32(value)<<RegisterRg0crFieldGnbreqShift))
}

// RegisterRg1crType DMAMux - DMA request generator channel x control register
type RegisterRg1crType uint32

func (r *RegisterRg1crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRg1crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRg1crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRg1crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRg1crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRg1crFieldSigidShift = 0
	RegisterRg1crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *RegisterRg1crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg1crFieldSigidMask) >> RegisterRg1crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *RegisterRg1crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg1crFieldSigidMask)|(uint32(value)<<RegisterRg1crFieldSigidShift))
}

const (
	RegisterRg1crFieldOieShift = 8
	RegisterRg1crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *RegisterRg1crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg1crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *RegisterRg1crType) SetOie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg1crFieldOieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg1crFieldOieMask)
	}
}

const (
	RegisterRg1crFieldGeShift = 16
	RegisterRg1crFieldGeMask  = 0x10000
)

// GetGe DMA request generator channel enable/disable
func (r *RegisterRg1crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg1crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *RegisterRg1crType) SetGe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg1crFieldGeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg1crFieldGeMask)
	}
}

const (
	RegisterRg1crFieldGpolShift = 17
	RegisterRg1crFieldGpolMask  = 0x60000
)

// GetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg1crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg1crFieldGpolMask) >> RegisterRg1crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg1crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg1crFieldGpolMask)|(uint32(value)<<RegisterRg1crFieldGpolShift))
}

const (
	RegisterRg1crFieldGnbreqShift = 19
	RegisterRg1crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg1crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg1crFieldGnbreqMask) >> RegisterRg1crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg1crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg1crFieldGnbreqMask)|(uint32(value)<<RegisterRg1crFieldGnbreqShift))
}

// RegisterRg2crType DMAMux - DMA request generator channel x control register
type RegisterRg2crType uint32

func (r *RegisterRg2crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRg2crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRg2crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRg2crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRg2crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRg2crFieldSigidShift = 0
	RegisterRg2crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *RegisterRg2crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg2crFieldSigidMask) >> RegisterRg2crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *RegisterRg2crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg2crFieldSigidMask)|(uint32(value)<<RegisterRg2crFieldSigidShift))
}

const (
	RegisterRg2crFieldOieShift = 8
	RegisterRg2crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *RegisterRg2crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg2crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *RegisterRg2crType) SetOie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg2crFieldOieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg2crFieldOieMask)
	}
}

const (
	RegisterRg2crFieldGeShift = 16
	RegisterRg2crFieldGeMask  = 0x10000
)

// GetGe DMA request generator channel enable/disable
func (r *RegisterRg2crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg2crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *RegisterRg2crType) SetGe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg2crFieldGeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg2crFieldGeMask)
	}
}

const (
	RegisterRg2crFieldGpolShift = 17
	RegisterRg2crFieldGpolMask  = 0x60000
)

// GetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg2crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg2crFieldGpolMask) >> RegisterRg2crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg2crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg2crFieldGpolMask)|(uint32(value)<<RegisterRg2crFieldGpolShift))
}

const (
	RegisterRg2crFieldGnbreqShift = 19
	RegisterRg2crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg2crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg2crFieldGnbreqMask) >> RegisterRg2crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg2crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg2crFieldGnbreqMask)|(uint32(value)<<RegisterRg2crFieldGnbreqShift))
}

// RegisterRg3crType DMAMux - DMA request generator channel x control register
type RegisterRg3crType uint32

func (r *RegisterRg3crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRg3crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRg3crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRg3crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRg3crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRg3crFieldSigidShift = 0
	RegisterRg3crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *RegisterRg3crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg3crFieldSigidMask) >> RegisterRg3crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *RegisterRg3crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg3crFieldSigidMask)|(uint32(value)<<RegisterRg3crFieldSigidShift))
}

const (
	RegisterRg3crFieldOieShift = 8
	RegisterRg3crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *RegisterRg3crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg3crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *RegisterRg3crType) SetOie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg3crFieldOieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg3crFieldOieMask)
	}
}

const (
	RegisterRg3crFieldGeShift = 16
	RegisterRg3crFieldGeMask  = 0x10000
)

// GetGe DMA request generator channel enable/disable
func (r *RegisterRg3crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg3crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *RegisterRg3crType) SetGe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg3crFieldGeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg3crFieldGeMask)
	}
}

const (
	RegisterRg3crFieldGpolShift = 17
	RegisterRg3crFieldGpolMask  = 0x60000
)

// GetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg3crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg3crFieldGpolMask) >> RegisterRg3crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg3crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg3crFieldGpolMask)|(uint32(value)<<RegisterRg3crFieldGpolShift))
}

const (
	RegisterRg3crFieldGnbreqShift = 19
	RegisterRg3crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg3crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg3crFieldGnbreqMask) >> RegisterRg3crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg3crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg3crFieldGnbreqMask)|(uint32(value)<<RegisterRg3crFieldGnbreqShift))
}

// RegisterRg4crType DMAMux - DMA request generator channel x control register
type RegisterRg4crType uint32

func (r *RegisterRg4crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRg4crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRg4crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRg4crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRg4crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRg4crFieldSigidShift = 0
	RegisterRg4crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *RegisterRg4crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg4crFieldSigidMask) >> RegisterRg4crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *RegisterRg4crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg4crFieldSigidMask)|(uint32(value)<<RegisterRg4crFieldSigidShift))
}

const (
	RegisterRg4crFieldOieShift = 8
	RegisterRg4crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *RegisterRg4crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg4crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *RegisterRg4crType) SetOie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg4crFieldOieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg4crFieldOieMask)
	}
}

const (
	RegisterRg4crFieldGeShift = 16
	RegisterRg4crFieldGeMask  = 0x10000
)

// GetGe DMA request generator channel enable/disable
func (r *RegisterRg4crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg4crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *RegisterRg4crType) SetGe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg4crFieldGeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg4crFieldGeMask)
	}
}

const (
	RegisterRg4crFieldGpolShift = 17
	RegisterRg4crFieldGpolMask  = 0x60000
)

// GetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg4crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg4crFieldGpolMask) >> RegisterRg4crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg4crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg4crFieldGpolMask)|(uint32(value)<<RegisterRg4crFieldGpolShift))
}

const (
	RegisterRg4crFieldGnbreqShift = 19
	RegisterRg4crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg4crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg4crFieldGnbreqMask) >> RegisterRg4crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg4crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg4crFieldGnbreqMask)|(uint32(value)<<RegisterRg4crFieldGnbreqShift))
}

// RegisterRg5crType DMAMux - DMA request generator channel x control register
type RegisterRg5crType uint32

func (r *RegisterRg5crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRg5crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRg5crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRg5crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRg5crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRg5crFieldSigidShift = 0
	RegisterRg5crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *RegisterRg5crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg5crFieldSigidMask) >> RegisterRg5crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *RegisterRg5crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg5crFieldSigidMask)|(uint32(value)<<RegisterRg5crFieldSigidShift))
}

const (
	RegisterRg5crFieldOieShift = 8
	RegisterRg5crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *RegisterRg5crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg5crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *RegisterRg5crType) SetOie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg5crFieldOieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg5crFieldOieMask)
	}
}

const (
	RegisterRg5crFieldGeShift = 16
	RegisterRg5crFieldGeMask  = 0x10000
)

// GetGe DMA request generator channel enable/disable
func (r *RegisterRg5crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg5crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *RegisterRg5crType) SetGe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg5crFieldGeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg5crFieldGeMask)
	}
}

const (
	RegisterRg5crFieldGpolShift = 17
	RegisterRg5crFieldGpolMask  = 0x60000
)

// GetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg5crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg5crFieldGpolMask) >> RegisterRg5crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg5crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg5crFieldGpolMask)|(uint32(value)<<RegisterRg5crFieldGpolShift))
}

const (
	RegisterRg5crFieldGnbreqShift = 19
	RegisterRg5crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg5crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg5crFieldGnbreqMask) >> RegisterRg5crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg5crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg5crFieldGnbreqMask)|(uint32(value)<<RegisterRg5crFieldGnbreqShift))
}

// RegisterRg6crType DMAMux - DMA request generator channel x control register
type RegisterRg6crType uint32

func (r *RegisterRg6crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRg6crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRg6crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRg6crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRg6crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRg6crFieldSigidShift = 0
	RegisterRg6crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *RegisterRg6crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg6crFieldSigidMask) >> RegisterRg6crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *RegisterRg6crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg6crFieldSigidMask)|(uint32(value)<<RegisterRg6crFieldSigidShift))
}

const (
	RegisterRg6crFieldOieShift = 8
	RegisterRg6crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *RegisterRg6crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg6crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *RegisterRg6crType) SetOie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg6crFieldOieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg6crFieldOieMask)
	}
}

const (
	RegisterRg6crFieldGeShift = 16
	RegisterRg6crFieldGeMask  = 0x10000
)

// GetGe DMA request generator channel enable/disable
func (r *RegisterRg6crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg6crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *RegisterRg6crType) SetGe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg6crFieldGeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg6crFieldGeMask)
	}
}

const (
	RegisterRg6crFieldGpolShift = 17
	RegisterRg6crFieldGpolMask  = 0x60000
)

// GetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg6crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg6crFieldGpolMask) >> RegisterRg6crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg6crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg6crFieldGpolMask)|(uint32(value)<<RegisterRg6crFieldGpolShift))
}

const (
	RegisterRg6crFieldGnbreqShift = 19
	RegisterRg6crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg6crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg6crFieldGnbreqMask) >> RegisterRg6crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg6crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg6crFieldGnbreqMask)|(uint32(value)<<RegisterRg6crFieldGnbreqShift))
}

// RegisterRg7crType DMAMux - DMA request generator channel x control register
type RegisterRg7crType uint32

func (r *RegisterRg7crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRg7crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRg7crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRg7crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRg7crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRg7crFieldSigidShift = 0
	RegisterRg7crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *RegisterRg7crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg7crFieldSigidMask) >> RegisterRg7crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *RegisterRg7crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg7crFieldSigidMask)|(uint32(value)<<RegisterRg7crFieldSigidShift))
}

const (
	RegisterRg7crFieldOieShift = 8
	RegisterRg7crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *RegisterRg7crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg7crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *RegisterRg7crType) SetOie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg7crFieldOieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg7crFieldOieMask)
	}
}

const (
	RegisterRg7crFieldGeShift = 16
	RegisterRg7crFieldGeMask  = 0x10000
)

// GetGe DMA request generator channel enable/disable
func (r *RegisterRg7crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg7crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *RegisterRg7crType) SetGe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRg7crFieldGeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRg7crFieldGeMask)
	}
}

const (
	RegisterRg7crFieldGpolShift = 17
	RegisterRg7crFieldGpolMask  = 0x60000
)

// GetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg7crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg7crFieldGpolMask) >> RegisterRg7crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *RegisterRg7crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg7crFieldGpolMask)|(uint32(value)<<RegisterRg7crFieldGpolShift))
}

const (
	RegisterRg7crFieldGnbreqShift = 19
	RegisterRg7crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg7crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg7crFieldGnbreqMask) >> RegisterRg7crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *RegisterRg7crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg7crFieldGnbreqMask)|(uint32(value)<<RegisterRg7crFieldGnbreqShift))
}

// RegisterRgsrType DMAMux - DMA request generator status register
type RegisterRgsrType uint32

func (r *RegisterRgsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRgsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRgsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRgsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRgsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRgsrFieldOfShift = 0
	RegisterRgsrFieldOfMask  = 0xff
)

// GetOf Trigger event overrun flag The flag is set when a trigger event occurs on DMA request generator channel x, while the DMA request generator counter value is lower than GNBREQ. The flag is cleared by writing 1 to the corresponding COFx bit in DMAMUX_RGCFR register.
func (r *RegisterRgsrType) GetOf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRgsrFieldOfMask) >> RegisterRgsrFieldOfShift)
}

// SetOf Trigger event overrun flag The flag is set when a trigger event occurs on DMA request generator channel x, while the DMA request generator counter value is lower than GNBREQ. The flag is cleared by writing 1 to the corresponding COFx bit in DMAMUX_RGCFR register.
func (r *RegisterRgsrType) SetOf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRgsrFieldOfMask)|(uint32(value)<<RegisterRgsrFieldOfShift))
}

// RegisterRgcfrType DMAMux - DMA request generator clear flag register
type RegisterRgcfrType uint32

func (r *RegisterRgcfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRgcfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRgcfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRgcfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRgcfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRgcfrFieldCofShift = 0
	RegisterRgcfrFieldCofMask  = 0xff
)

// GetCof Clear trigger event overrun flag Upon setting, this bit clears the corresponding overrun flag OFx in the DMAMUX_RGCSR register.
func (r *RegisterRgcfrType) GetCof() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRgcfrFieldCofMask) >> RegisterRgcfrFieldCofShift)
}

// SetCof Clear trigger event overrun flag Upon setting, this bit clears the corresponding overrun flag OFx in the DMAMUX_RGCSR register.
func (r *RegisterRgcfrType) SetCof(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRgcfrFieldCofMask)|(uint32(value)<<RegisterRgcfrFieldCofShift))
}
