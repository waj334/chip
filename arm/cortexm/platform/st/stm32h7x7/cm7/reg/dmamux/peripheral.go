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
	C0cr  registerC0crType
	C1cr  registerC1crType
	C2cr  registerC2crType
	C3cr  registerC3crType
	C4cr  registerC4crType
	C5cr  registerC5crType
	C6cr  registerC6crType
	C7cr  registerC7crType
	C8cr  registerC8crType
	C9cr  registerC9crType
	C10cr registerC10crType
	C11cr registerC11crType
	C12cr registerC12crType
	C13cr registerC13crType
	C14cr registerC14crType
	C15cr registerC15crType
	_     [64]byte
	Csr   registerCsrType
	Cfr   registerCfrType
	_     [120]byte
	Rg0cr registerRg0crType
	Rg1cr registerRg1crType
	Rg2cr registerRg2crType
	Rg3cr registerRg3crType
	Rg4cr registerRg4crType
	Rg5cr registerRg5crType
	Rg6cr registerRg6crType
	Rg7cr registerRg7crType
	_     [32]byte
	Rgsr  registerRgsrType
	Rgcfr registerRgcfrType
}

// registerC0crType DMAMux - DMA request line multiplexer channel x control register
type registerC0crType uint32

const (
	RegisterC0crFieldDmareqidShift = 0
	RegisterC0crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC0crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldDmareqidMask) >> RegisterC0crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC0crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC0crFieldDmareqidMask)|(uint32(value)<<RegisterC0crFieldDmareqidShift))
}

const (
	RegisterC0crFieldSoieShift = 8
	RegisterC0crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC0crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC0crType) SetSoie(value bool) {
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
func (r *registerC0crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC0crType) SetEge(value bool) {
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
func (r *registerC0crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC0crType) SetSe(value bool) {
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
func (r *registerC0crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldSpolMask) >> RegisterC0crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC0crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC0crFieldSpolMask)|(uint32(value)<<RegisterC0crFieldSpolShift))
}

const (
	RegisterC0crFieldNbreqShift = 19
	RegisterC0crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC0crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldNbreqMask) >> RegisterC0crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC0crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC0crFieldNbreqMask)|(uint32(value)<<RegisterC0crFieldNbreqShift))
}

const (
	RegisterC0crFieldSyncidShift = 24
	RegisterC0crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC0crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC0crFieldSyncidMask) >> RegisterC0crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC0crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC0crFieldSyncidMask)|(uint32(value)<<RegisterC0crFieldSyncidShift))
}

// registerC1crType DMAMux - DMA request line multiplexer channel x control register
type registerC1crType uint32

const (
	RegisterC1crFieldDmareqidShift = 0
	RegisterC1crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC1crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldDmareqidMask) >> RegisterC1crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC1crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC1crFieldDmareqidMask)|(uint32(value)<<RegisterC1crFieldDmareqidShift))
}

const (
	RegisterC1crFieldSoieShift = 8
	RegisterC1crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC1crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC1crType) SetSoie(value bool) {
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
func (r *registerC1crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC1crType) SetEge(value bool) {
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
func (r *registerC1crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC1crType) SetSe(value bool) {
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
func (r *registerC1crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldSpolMask) >> RegisterC1crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC1crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC1crFieldSpolMask)|(uint32(value)<<RegisterC1crFieldSpolShift))
}

const (
	RegisterC1crFieldNbreqShift = 19
	RegisterC1crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC1crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldNbreqMask) >> RegisterC1crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC1crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC1crFieldNbreqMask)|(uint32(value)<<RegisterC1crFieldNbreqShift))
}

const (
	RegisterC1crFieldSyncidShift = 24
	RegisterC1crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC1crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC1crFieldSyncidMask) >> RegisterC1crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC1crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC1crFieldSyncidMask)|(uint32(value)<<RegisterC1crFieldSyncidShift))
}

// registerC2crType DMAMux - DMA request line multiplexer channel x control register
type registerC2crType uint32

const (
	RegisterC2crFieldDmareqidShift = 0
	RegisterC2crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC2crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldDmareqidMask) >> RegisterC2crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC2crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC2crFieldDmareqidMask)|(uint32(value)<<RegisterC2crFieldDmareqidShift))
}

const (
	RegisterC2crFieldSoieShift = 8
	RegisterC2crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC2crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC2crType) SetSoie(value bool) {
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
func (r *registerC2crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC2crType) SetEge(value bool) {
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
func (r *registerC2crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC2crType) SetSe(value bool) {
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
func (r *registerC2crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldSpolMask) >> RegisterC2crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC2crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC2crFieldSpolMask)|(uint32(value)<<RegisterC2crFieldSpolShift))
}

const (
	RegisterC2crFieldNbreqShift = 19
	RegisterC2crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC2crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldNbreqMask) >> RegisterC2crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC2crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC2crFieldNbreqMask)|(uint32(value)<<RegisterC2crFieldNbreqShift))
}

const (
	RegisterC2crFieldSyncidShift = 24
	RegisterC2crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC2crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC2crFieldSyncidMask) >> RegisterC2crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC2crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC2crFieldSyncidMask)|(uint32(value)<<RegisterC2crFieldSyncidShift))
}

// registerC3crType DMAMux - DMA request line multiplexer channel x control register
type registerC3crType uint32

const (
	RegisterC3crFieldDmareqidShift = 0
	RegisterC3crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC3crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldDmareqidMask) >> RegisterC3crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC3crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC3crFieldDmareqidMask)|(uint32(value)<<RegisterC3crFieldDmareqidShift))
}

const (
	RegisterC3crFieldSoieShift = 8
	RegisterC3crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC3crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC3crType) SetSoie(value bool) {
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
func (r *registerC3crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC3crType) SetEge(value bool) {
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
func (r *registerC3crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC3crType) SetSe(value bool) {
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
func (r *registerC3crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldSpolMask) >> RegisterC3crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC3crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC3crFieldSpolMask)|(uint32(value)<<RegisterC3crFieldSpolShift))
}

const (
	RegisterC3crFieldNbreqShift = 19
	RegisterC3crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC3crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldNbreqMask) >> RegisterC3crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC3crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC3crFieldNbreqMask)|(uint32(value)<<RegisterC3crFieldNbreqShift))
}

const (
	RegisterC3crFieldSyncidShift = 24
	RegisterC3crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC3crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC3crFieldSyncidMask) >> RegisterC3crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC3crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC3crFieldSyncidMask)|(uint32(value)<<RegisterC3crFieldSyncidShift))
}

// registerC4crType DMAMux - DMA request line multiplexer channel x control register
type registerC4crType uint32

const (
	RegisterC4crFieldDmareqidShift = 0
	RegisterC4crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC4crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldDmareqidMask) >> RegisterC4crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC4crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC4crFieldDmareqidMask)|(uint32(value)<<RegisterC4crFieldDmareqidShift))
}

const (
	RegisterC4crFieldSoieShift = 8
	RegisterC4crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC4crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC4crType) SetSoie(value bool) {
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
func (r *registerC4crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC4crType) SetEge(value bool) {
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
func (r *registerC4crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC4crType) SetSe(value bool) {
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
func (r *registerC4crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldSpolMask) >> RegisterC4crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC4crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC4crFieldSpolMask)|(uint32(value)<<RegisterC4crFieldSpolShift))
}

const (
	RegisterC4crFieldNbreqShift = 19
	RegisterC4crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC4crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldNbreqMask) >> RegisterC4crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC4crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC4crFieldNbreqMask)|(uint32(value)<<RegisterC4crFieldNbreqShift))
}

const (
	RegisterC4crFieldSyncidShift = 24
	RegisterC4crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC4crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC4crFieldSyncidMask) >> RegisterC4crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC4crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC4crFieldSyncidMask)|(uint32(value)<<RegisterC4crFieldSyncidShift))
}

// registerC5crType DMAMux - DMA request line multiplexer channel x control register
type registerC5crType uint32

const (
	RegisterC5crFieldDmareqidShift = 0
	RegisterC5crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC5crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldDmareqidMask) >> RegisterC5crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC5crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC5crFieldDmareqidMask)|(uint32(value)<<RegisterC5crFieldDmareqidShift))
}

const (
	RegisterC5crFieldSoieShift = 8
	RegisterC5crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC5crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC5crType) SetSoie(value bool) {
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
func (r *registerC5crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC5crType) SetEge(value bool) {
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
func (r *registerC5crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC5crType) SetSe(value bool) {
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
func (r *registerC5crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldSpolMask) >> RegisterC5crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC5crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC5crFieldSpolMask)|(uint32(value)<<RegisterC5crFieldSpolShift))
}

const (
	RegisterC5crFieldNbreqShift = 19
	RegisterC5crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC5crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldNbreqMask) >> RegisterC5crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC5crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC5crFieldNbreqMask)|(uint32(value)<<RegisterC5crFieldNbreqShift))
}

const (
	RegisterC5crFieldSyncidShift = 24
	RegisterC5crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC5crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC5crFieldSyncidMask) >> RegisterC5crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC5crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC5crFieldSyncidMask)|(uint32(value)<<RegisterC5crFieldSyncidShift))
}

// registerC6crType DMAMux - DMA request line multiplexer channel x control register
type registerC6crType uint32

const (
	RegisterC6crFieldDmareqidShift = 0
	RegisterC6crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC6crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldDmareqidMask) >> RegisterC6crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC6crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC6crFieldDmareqidMask)|(uint32(value)<<RegisterC6crFieldDmareqidShift))
}

const (
	RegisterC6crFieldSoieShift = 8
	RegisterC6crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC6crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC6crType) SetSoie(value bool) {
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
func (r *registerC6crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC6crType) SetEge(value bool) {
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
func (r *registerC6crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC6crType) SetSe(value bool) {
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
func (r *registerC6crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldSpolMask) >> RegisterC6crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC6crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC6crFieldSpolMask)|(uint32(value)<<RegisterC6crFieldSpolShift))
}

const (
	RegisterC6crFieldNbreqShift = 19
	RegisterC6crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC6crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldNbreqMask) >> RegisterC6crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC6crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC6crFieldNbreqMask)|(uint32(value)<<RegisterC6crFieldNbreqShift))
}

const (
	RegisterC6crFieldSyncidShift = 24
	RegisterC6crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC6crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC6crFieldSyncidMask) >> RegisterC6crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC6crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC6crFieldSyncidMask)|(uint32(value)<<RegisterC6crFieldSyncidShift))
}

// registerC7crType DMAMux - DMA request line multiplexer channel x control register
type registerC7crType uint32

const (
	RegisterC7crFieldDmareqidShift = 0
	RegisterC7crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC7crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldDmareqidMask) >> RegisterC7crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC7crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC7crFieldDmareqidMask)|(uint32(value)<<RegisterC7crFieldDmareqidShift))
}

const (
	RegisterC7crFieldSoieShift = 8
	RegisterC7crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC7crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC7crType) SetSoie(value bool) {
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
func (r *registerC7crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC7crType) SetEge(value bool) {
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
func (r *registerC7crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC7crType) SetSe(value bool) {
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
func (r *registerC7crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldSpolMask) >> RegisterC7crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC7crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC7crFieldSpolMask)|(uint32(value)<<RegisterC7crFieldSpolShift))
}

const (
	RegisterC7crFieldNbreqShift = 19
	RegisterC7crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC7crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldNbreqMask) >> RegisterC7crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC7crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC7crFieldNbreqMask)|(uint32(value)<<RegisterC7crFieldNbreqShift))
}

const (
	RegisterC7crFieldSyncidShift = 24
	RegisterC7crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC7crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC7crFieldSyncidMask) >> RegisterC7crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC7crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC7crFieldSyncidMask)|(uint32(value)<<RegisterC7crFieldSyncidShift))
}

// registerC8crType DMAMux - DMA request line multiplexer channel x control register
type registerC8crType uint32

const (
	RegisterC8crFieldDmareqidShift = 0
	RegisterC8crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC8crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldDmareqidMask) >> RegisterC8crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC8crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC8crFieldDmareqidMask)|(uint32(value)<<RegisterC8crFieldDmareqidShift))
}

const (
	RegisterC8crFieldSoieShift = 8
	RegisterC8crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC8crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC8crType) SetSoie(value bool) {
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
func (r *registerC8crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC8crType) SetEge(value bool) {
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
func (r *registerC8crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC8crType) SetSe(value bool) {
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
func (r *registerC8crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldSpolMask) >> RegisterC8crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC8crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC8crFieldSpolMask)|(uint32(value)<<RegisterC8crFieldSpolShift))
}

const (
	RegisterC8crFieldNbreqShift = 19
	RegisterC8crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC8crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldNbreqMask) >> RegisterC8crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC8crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC8crFieldNbreqMask)|(uint32(value)<<RegisterC8crFieldNbreqShift))
}

const (
	RegisterC8crFieldSyncidShift = 24
	RegisterC8crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC8crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC8crFieldSyncidMask) >> RegisterC8crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC8crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC8crFieldSyncidMask)|(uint32(value)<<RegisterC8crFieldSyncidShift))
}

// registerC9crType DMAMux - DMA request line multiplexer channel x control register
type registerC9crType uint32

const (
	RegisterC9crFieldDmareqidShift = 0
	RegisterC9crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC9crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldDmareqidMask) >> RegisterC9crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC9crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC9crFieldDmareqidMask)|(uint32(value)<<RegisterC9crFieldDmareqidShift))
}

const (
	RegisterC9crFieldSoieShift = 8
	RegisterC9crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC9crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC9crType) SetSoie(value bool) {
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
func (r *registerC9crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC9crType) SetEge(value bool) {
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
func (r *registerC9crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC9crType) SetSe(value bool) {
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
func (r *registerC9crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldSpolMask) >> RegisterC9crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC9crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC9crFieldSpolMask)|(uint32(value)<<RegisterC9crFieldSpolShift))
}

const (
	RegisterC9crFieldNbreqShift = 19
	RegisterC9crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC9crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldNbreqMask) >> RegisterC9crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC9crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC9crFieldNbreqMask)|(uint32(value)<<RegisterC9crFieldNbreqShift))
}

const (
	RegisterC9crFieldSyncidShift = 24
	RegisterC9crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC9crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC9crFieldSyncidMask) >> RegisterC9crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC9crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC9crFieldSyncidMask)|(uint32(value)<<RegisterC9crFieldSyncidShift))
}

// registerC10crType DMAMux - DMA request line multiplexer channel x control register
type registerC10crType uint32

const (
	RegisterC10crFieldDmareqidShift = 0
	RegisterC10crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC10crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldDmareqidMask) >> RegisterC10crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC10crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC10crFieldDmareqidMask)|(uint32(value)<<RegisterC10crFieldDmareqidShift))
}

const (
	RegisterC10crFieldSoieShift = 8
	RegisterC10crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC10crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC10crType) SetSoie(value bool) {
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
func (r *registerC10crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC10crType) SetEge(value bool) {
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
func (r *registerC10crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC10crType) SetSe(value bool) {
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
func (r *registerC10crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldSpolMask) >> RegisterC10crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC10crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC10crFieldSpolMask)|(uint32(value)<<RegisterC10crFieldSpolShift))
}

const (
	RegisterC10crFieldNbreqShift = 19
	RegisterC10crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC10crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldNbreqMask) >> RegisterC10crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC10crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC10crFieldNbreqMask)|(uint32(value)<<RegisterC10crFieldNbreqShift))
}

const (
	RegisterC10crFieldSyncidShift = 24
	RegisterC10crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC10crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC10crFieldSyncidMask) >> RegisterC10crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC10crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC10crFieldSyncidMask)|(uint32(value)<<RegisterC10crFieldSyncidShift))
}

// registerC11crType DMAMux - DMA request line multiplexer channel x control register
type registerC11crType uint32

const (
	RegisterC11crFieldDmareqidShift = 0
	RegisterC11crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC11crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldDmareqidMask) >> RegisterC11crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC11crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC11crFieldDmareqidMask)|(uint32(value)<<RegisterC11crFieldDmareqidShift))
}

const (
	RegisterC11crFieldSoieShift = 8
	RegisterC11crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC11crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC11crType) SetSoie(value bool) {
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
func (r *registerC11crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC11crType) SetEge(value bool) {
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
func (r *registerC11crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC11crType) SetSe(value bool) {
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
func (r *registerC11crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldSpolMask) >> RegisterC11crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC11crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC11crFieldSpolMask)|(uint32(value)<<RegisterC11crFieldSpolShift))
}

const (
	RegisterC11crFieldNbreqShift = 19
	RegisterC11crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC11crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldNbreqMask) >> RegisterC11crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC11crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC11crFieldNbreqMask)|(uint32(value)<<RegisterC11crFieldNbreqShift))
}

const (
	RegisterC11crFieldSyncidShift = 24
	RegisterC11crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC11crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC11crFieldSyncidMask) >> RegisterC11crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC11crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC11crFieldSyncidMask)|(uint32(value)<<RegisterC11crFieldSyncidShift))
}

// registerC12crType DMAMux - DMA request line multiplexer channel x control register
type registerC12crType uint32

const (
	RegisterC12crFieldDmareqidShift = 0
	RegisterC12crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC12crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldDmareqidMask) >> RegisterC12crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC12crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC12crFieldDmareqidMask)|(uint32(value)<<RegisterC12crFieldDmareqidShift))
}

const (
	RegisterC12crFieldSoieShift = 8
	RegisterC12crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC12crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC12crType) SetSoie(value bool) {
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
func (r *registerC12crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC12crType) SetEge(value bool) {
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
func (r *registerC12crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC12crType) SetSe(value bool) {
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
func (r *registerC12crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldSpolMask) >> RegisterC12crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC12crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC12crFieldSpolMask)|(uint32(value)<<RegisterC12crFieldSpolShift))
}

const (
	RegisterC12crFieldNbreqShift = 19
	RegisterC12crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC12crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldNbreqMask) >> RegisterC12crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC12crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC12crFieldNbreqMask)|(uint32(value)<<RegisterC12crFieldNbreqShift))
}

const (
	RegisterC12crFieldSyncidShift = 24
	RegisterC12crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC12crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC12crFieldSyncidMask) >> RegisterC12crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC12crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC12crFieldSyncidMask)|(uint32(value)<<RegisterC12crFieldSyncidShift))
}

// registerC13crType DMAMux - DMA request line multiplexer channel x control register
type registerC13crType uint32

const (
	RegisterC13crFieldDmareqidShift = 0
	RegisterC13crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC13crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldDmareqidMask) >> RegisterC13crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC13crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC13crFieldDmareqidMask)|(uint32(value)<<RegisterC13crFieldDmareqidShift))
}

const (
	RegisterC13crFieldSoieShift = 8
	RegisterC13crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC13crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC13crType) SetSoie(value bool) {
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
func (r *registerC13crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC13crType) SetEge(value bool) {
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
func (r *registerC13crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC13crType) SetSe(value bool) {
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
func (r *registerC13crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldSpolMask) >> RegisterC13crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC13crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC13crFieldSpolMask)|(uint32(value)<<RegisterC13crFieldSpolShift))
}

const (
	RegisterC13crFieldNbreqShift = 19
	RegisterC13crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC13crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldNbreqMask) >> RegisterC13crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC13crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC13crFieldNbreqMask)|(uint32(value)<<RegisterC13crFieldNbreqShift))
}

const (
	RegisterC13crFieldSyncidShift = 24
	RegisterC13crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC13crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC13crFieldSyncidMask) >> RegisterC13crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC13crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC13crFieldSyncidMask)|(uint32(value)<<RegisterC13crFieldSyncidShift))
}

// registerC14crType DMAMux - DMA request line multiplexer channel x control register
type registerC14crType uint32

const (
	RegisterC14crFieldDmareqidShift = 0
	RegisterC14crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC14crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldDmareqidMask) >> RegisterC14crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC14crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC14crFieldDmareqidMask)|(uint32(value)<<RegisterC14crFieldDmareqidShift))
}

const (
	RegisterC14crFieldSoieShift = 8
	RegisterC14crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC14crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC14crType) SetSoie(value bool) {
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
func (r *registerC14crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC14crType) SetEge(value bool) {
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
func (r *registerC14crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC14crType) SetSe(value bool) {
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
func (r *registerC14crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldSpolMask) >> RegisterC14crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC14crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC14crFieldSpolMask)|(uint32(value)<<RegisterC14crFieldSpolShift))
}

const (
	RegisterC14crFieldNbreqShift = 19
	RegisterC14crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC14crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldNbreqMask) >> RegisterC14crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC14crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC14crFieldNbreqMask)|(uint32(value)<<RegisterC14crFieldNbreqShift))
}

const (
	RegisterC14crFieldSyncidShift = 24
	RegisterC14crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC14crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC14crFieldSyncidMask) >> RegisterC14crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC14crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC14crFieldSyncidMask)|(uint32(value)<<RegisterC14crFieldSyncidShift))
}

// registerC15crType DMAMux - DMA request line multiplexer channel x control register
type registerC15crType uint32

const (
	RegisterC15crFieldDmareqidShift = 0
	RegisterC15crFieldDmareqidMask  = 0xff
)

// GetDmareqid Input DMA request line selected
func (r *registerC15crType) GetDmareqid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldDmareqidMask) >> RegisterC15crFieldDmareqidShift)
}

// SetDmareqid Input DMA request line selected
func (r *registerC15crType) SetDmareqid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC15crFieldDmareqidMask)|(uint32(value)<<RegisterC15crFieldDmareqidShift))
}

const (
	RegisterC15crFieldSoieShift = 8
	RegisterC15crFieldSoieMask  = 0x100
)

// GetSoie Interrupt enable at synchronization event overrun
func (r *registerC15crType) GetSoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldSoieMask) != 0
}

// SetSoie Interrupt enable at synchronization event overrun
func (r *registerC15crType) SetSoie(value bool) {
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
func (r *registerC15crType) GetEge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldEgeMask) != 0
}

// SetEge Event generation enable/disable
func (r *registerC15crType) SetEge(value bool) {
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
func (r *registerC15crType) GetSe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldSeMask) != 0
}

// SetSe Synchronous operating mode enable/disable
func (r *registerC15crType) SetSe(value bool) {
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
func (r *registerC15crType) GetSpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldSpolMask) >> RegisterC15crFieldSpolShift)
}

// SetSpol Synchronization event type selector Defines the synchronization event on the selected synchronization input:
func (r *registerC15crType) SetSpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC15crFieldSpolMask)|(uint32(value)<<RegisterC15crFieldSpolShift))
}

const (
	RegisterC15crFieldNbreqShift = 19
	RegisterC15crFieldNbreqMask  = 0xf80000
)

// GetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC15crType) GetNbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldNbreqMask) >> RegisterC15crFieldNbreqShift)
}

// SetNbreq Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
func (r *registerC15crType) SetNbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC15crFieldNbreqMask)|(uint32(value)<<RegisterC15crFieldNbreqShift))
}

const (
	RegisterC15crFieldSyncidShift = 24
	RegisterC15crFieldSyncidMask  = 0x1f000000
)

// GetSyncid Synchronization input selected
func (r *registerC15crType) GetSyncid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterC15crFieldSyncidMask) >> RegisterC15crFieldSyncidShift)
}

// SetSyncid Synchronization input selected
func (r *registerC15crType) SetSyncid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterC15crFieldSyncidMask)|(uint32(value)<<RegisterC15crFieldSyncidShift))
}

// registerCsrType DMAMUX request line multiplexer interrupt channel status register
type registerCsrType uint32

const (
	RegisterCsrFieldSofShift = 0
	RegisterCsrFieldSofMask  = 0xffff
)

// GetSof Synchronization overrun event flag
func (r *registerCsrType) GetSof() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldSofMask) >> RegisterCsrFieldSofShift)
}

// SetSof Synchronization overrun event flag
func (r *registerCsrType) SetSof(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldSofMask)|(uint32(value)<<RegisterCsrFieldSofShift))
}

// registerCfrType DMAMUX request line multiplexer interrupt clear flag register
type registerCfrType uint32

const (
	RegisterCfrFieldCsofShift = 0
	RegisterCfrFieldCsofMask  = 0xffff
)

// GetCsof Clear synchronization overrun event flag
func (r *registerCfrType) GetCsof() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCfrFieldCsofMask) >> RegisterCfrFieldCsofShift)
}

// SetCsof Clear synchronization overrun event flag
func (r *registerCfrType) SetCsof(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfrFieldCsofMask)|(uint32(value)<<RegisterCfrFieldCsofShift))
}

// registerRg0crType DMAMux - DMA request generator channel x control register
type registerRg0crType uint32

const (
	RegisterRg0crFieldSigidShift = 0
	RegisterRg0crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *registerRg0crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg0crFieldSigidMask) >> RegisterRg0crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *registerRg0crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg0crFieldSigidMask)|(uint32(value)<<RegisterRg0crFieldSigidShift))
}

const (
	RegisterRg0crFieldOieShift = 8
	RegisterRg0crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *registerRg0crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg0crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *registerRg0crType) SetOie(value bool) {
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
func (r *registerRg0crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg0crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *registerRg0crType) SetGe(value bool) {
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
func (r *registerRg0crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg0crFieldGpolMask) >> RegisterRg0crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *registerRg0crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg0crFieldGpolMask)|(uint32(value)<<RegisterRg0crFieldGpolShift))
}

const (
	RegisterRg0crFieldGnbreqShift = 19
	RegisterRg0crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg0crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg0crFieldGnbreqMask) >> RegisterRg0crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg0crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg0crFieldGnbreqMask)|(uint32(value)<<RegisterRg0crFieldGnbreqShift))
}

// registerRg1crType DMAMux - DMA request generator channel x control register
type registerRg1crType uint32

const (
	RegisterRg1crFieldSigidShift = 0
	RegisterRg1crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *registerRg1crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg1crFieldSigidMask) >> RegisterRg1crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *registerRg1crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg1crFieldSigidMask)|(uint32(value)<<RegisterRg1crFieldSigidShift))
}

const (
	RegisterRg1crFieldOieShift = 8
	RegisterRg1crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *registerRg1crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg1crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *registerRg1crType) SetOie(value bool) {
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
func (r *registerRg1crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg1crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *registerRg1crType) SetGe(value bool) {
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
func (r *registerRg1crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg1crFieldGpolMask) >> RegisterRg1crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *registerRg1crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg1crFieldGpolMask)|(uint32(value)<<RegisterRg1crFieldGpolShift))
}

const (
	RegisterRg1crFieldGnbreqShift = 19
	RegisterRg1crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg1crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg1crFieldGnbreqMask) >> RegisterRg1crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg1crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg1crFieldGnbreqMask)|(uint32(value)<<RegisterRg1crFieldGnbreqShift))
}

// registerRg2crType DMAMux - DMA request generator channel x control register
type registerRg2crType uint32

const (
	RegisterRg2crFieldSigidShift = 0
	RegisterRg2crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *registerRg2crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg2crFieldSigidMask) >> RegisterRg2crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *registerRg2crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg2crFieldSigidMask)|(uint32(value)<<RegisterRg2crFieldSigidShift))
}

const (
	RegisterRg2crFieldOieShift = 8
	RegisterRg2crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *registerRg2crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg2crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *registerRg2crType) SetOie(value bool) {
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
func (r *registerRg2crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg2crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *registerRg2crType) SetGe(value bool) {
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
func (r *registerRg2crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg2crFieldGpolMask) >> RegisterRg2crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *registerRg2crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg2crFieldGpolMask)|(uint32(value)<<RegisterRg2crFieldGpolShift))
}

const (
	RegisterRg2crFieldGnbreqShift = 19
	RegisterRg2crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg2crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg2crFieldGnbreqMask) >> RegisterRg2crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg2crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg2crFieldGnbreqMask)|(uint32(value)<<RegisterRg2crFieldGnbreqShift))
}

// registerRg3crType DMAMux - DMA request generator channel x control register
type registerRg3crType uint32

const (
	RegisterRg3crFieldSigidShift = 0
	RegisterRg3crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *registerRg3crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg3crFieldSigidMask) >> RegisterRg3crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *registerRg3crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg3crFieldSigidMask)|(uint32(value)<<RegisterRg3crFieldSigidShift))
}

const (
	RegisterRg3crFieldOieShift = 8
	RegisterRg3crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *registerRg3crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg3crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *registerRg3crType) SetOie(value bool) {
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
func (r *registerRg3crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg3crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *registerRg3crType) SetGe(value bool) {
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
func (r *registerRg3crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg3crFieldGpolMask) >> RegisterRg3crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *registerRg3crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg3crFieldGpolMask)|(uint32(value)<<RegisterRg3crFieldGpolShift))
}

const (
	RegisterRg3crFieldGnbreqShift = 19
	RegisterRg3crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg3crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg3crFieldGnbreqMask) >> RegisterRg3crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg3crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg3crFieldGnbreqMask)|(uint32(value)<<RegisterRg3crFieldGnbreqShift))
}

// registerRg4crType DMAMux - DMA request generator channel x control register
type registerRg4crType uint32

const (
	RegisterRg4crFieldSigidShift = 0
	RegisterRg4crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *registerRg4crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg4crFieldSigidMask) >> RegisterRg4crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *registerRg4crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg4crFieldSigidMask)|(uint32(value)<<RegisterRg4crFieldSigidShift))
}

const (
	RegisterRg4crFieldOieShift = 8
	RegisterRg4crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *registerRg4crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg4crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *registerRg4crType) SetOie(value bool) {
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
func (r *registerRg4crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg4crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *registerRg4crType) SetGe(value bool) {
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
func (r *registerRg4crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg4crFieldGpolMask) >> RegisterRg4crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *registerRg4crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg4crFieldGpolMask)|(uint32(value)<<RegisterRg4crFieldGpolShift))
}

const (
	RegisterRg4crFieldGnbreqShift = 19
	RegisterRg4crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg4crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg4crFieldGnbreqMask) >> RegisterRg4crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg4crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg4crFieldGnbreqMask)|(uint32(value)<<RegisterRg4crFieldGnbreqShift))
}

// registerRg5crType DMAMux - DMA request generator channel x control register
type registerRg5crType uint32

const (
	RegisterRg5crFieldSigidShift = 0
	RegisterRg5crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *registerRg5crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg5crFieldSigidMask) >> RegisterRg5crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *registerRg5crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg5crFieldSigidMask)|(uint32(value)<<RegisterRg5crFieldSigidShift))
}

const (
	RegisterRg5crFieldOieShift = 8
	RegisterRg5crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *registerRg5crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg5crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *registerRg5crType) SetOie(value bool) {
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
func (r *registerRg5crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg5crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *registerRg5crType) SetGe(value bool) {
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
func (r *registerRg5crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg5crFieldGpolMask) >> RegisterRg5crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *registerRg5crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg5crFieldGpolMask)|(uint32(value)<<RegisterRg5crFieldGpolShift))
}

const (
	RegisterRg5crFieldGnbreqShift = 19
	RegisterRg5crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg5crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg5crFieldGnbreqMask) >> RegisterRg5crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg5crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg5crFieldGnbreqMask)|(uint32(value)<<RegisterRg5crFieldGnbreqShift))
}

// registerRg6crType DMAMux - DMA request generator channel x control register
type registerRg6crType uint32

const (
	RegisterRg6crFieldSigidShift = 0
	RegisterRg6crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *registerRg6crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg6crFieldSigidMask) >> RegisterRg6crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *registerRg6crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg6crFieldSigidMask)|(uint32(value)<<RegisterRg6crFieldSigidShift))
}

const (
	RegisterRg6crFieldOieShift = 8
	RegisterRg6crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *registerRg6crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg6crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *registerRg6crType) SetOie(value bool) {
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
func (r *registerRg6crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg6crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *registerRg6crType) SetGe(value bool) {
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
func (r *registerRg6crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg6crFieldGpolMask) >> RegisterRg6crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *registerRg6crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg6crFieldGpolMask)|(uint32(value)<<RegisterRg6crFieldGpolShift))
}

const (
	RegisterRg6crFieldGnbreqShift = 19
	RegisterRg6crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg6crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg6crFieldGnbreqMask) >> RegisterRg6crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg6crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg6crFieldGnbreqMask)|(uint32(value)<<RegisterRg6crFieldGnbreqShift))
}

// registerRg7crType DMAMux - DMA request generator channel x control register
type registerRg7crType uint32

const (
	RegisterRg7crFieldSigidShift = 0
	RegisterRg7crFieldSigidMask  = 0x1f
)

// GetSigid DMA request trigger input selected
func (r *registerRg7crType) GetSigid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg7crFieldSigidMask) >> RegisterRg7crFieldSigidShift)
}

// SetSigid DMA request trigger input selected
func (r *registerRg7crType) SetSigid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg7crFieldSigidMask)|(uint32(value)<<RegisterRg7crFieldSigidShift))
}

const (
	RegisterRg7crFieldOieShift = 8
	RegisterRg7crFieldOieMask  = 0x100
)

// GetOie Interrupt enable at trigger event overrun
func (r *registerRg7crType) GetOie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg7crFieldOieMask) != 0
}

// SetOie Interrupt enable at trigger event overrun
func (r *registerRg7crType) SetOie(value bool) {
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
func (r *registerRg7crType) GetGe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRg7crFieldGeMask) != 0
}

// SetGe DMA request generator channel enable/disable
func (r *registerRg7crType) SetGe(value bool) {
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
func (r *registerRg7crType) GetGpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg7crFieldGpolMask) >> RegisterRg7crFieldGpolShift)
}

// SetGpol DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
func (r *registerRg7crType) SetGpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg7crFieldGpolMask)|(uint32(value)<<RegisterRg7crFieldGpolShift))
}

const (
	RegisterRg7crFieldGnbreqShift = 19
	RegisterRg7crFieldGnbreqMask  = 0xf80000
)

// GetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg7crType) GetGnbreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRg7crFieldGnbreqMask) >> RegisterRg7crFieldGnbreqShift)
}

// SetGnbreq Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
func (r *registerRg7crType) SetGnbreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRg7crFieldGnbreqMask)|(uint32(value)<<RegisterRg7crFieldGnbreqShift))
}

// registerRgsrType DMAMux - DMA request generator status register
type registerRgsrType uint32

const (
	RegisterRgsrFieldOfShift = 0
	RegisterRgsrFieldOfMask  = 0xff
)

// GetOf Trigger event overrun flag The flag is set when a trigger event occurs on DMA request generator channel x, while the DMA request generator counter value is lower than GNBREQ. The flag is cleared by writing 1 to the corresponding COFx bit in DMAMUX_RGCFR register.
func (r *registerRgsrType) GetOf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRgsrFieldOfMask) >> RegisterRgsrFieldOfShift)
}

// SetOf Trigger event overrun flag The flag is set when a trigger event occurs on DMA request generator channel x, while the DMA request generator counter value is lower than GNBREQ. The flag is cleared by writing 1 to the corresponding COFx bit in DMAMUX_RGCFR register.
func (r *registerRgsrType) SetOf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRgsrFieldOfMask)|(uint32(value)<<RegisterRgsrFieldOfShift))
}

// registerRgcfrType DMAMux - DMA request generator clear flag register
type registerRgcfrType uint32

const (
	RegisterRgcfrFieldCofShift = 0
	RegisterRgcfrFieldCofMask  = 0xff
)

// GetCof Clear trigger event overrun flag Upon setting, this bit clears the corresponding overrun flag OFx in the DMAMUX_RGCSR register.
func (r *registerRgcfrType) GetCof() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRgcfrFieldCofMask) >> RegisterRgcfrFieldCofShift)
}

// SetCof Clear trigger event overrun flag Upon setting, this bit clears the corresponding overrun flag OFx in the DMAMUX_RGCSR register.
func (r *registerRgcfrType) SetCof(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRgcfrFieldCofMask)|(uint32(value)<<RegisterRgcfrFieldCofShift))
}
