//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package mdma

import (
	"unsafe"
	"volatile"
)

var (
	Mdma = (*_mdma)(unsafe.Pointer(uintptr(0x52000000)))
)

type _mdma struct {
	Mdma_gisr0    registerMdma_gisr0Type
	_             [60]byte
	Mdma_c0isr    registerMdma_c0isrType
	Mdma_c0ifcr   registerMdma_c0ifcrType
	Mdma_c0esr    registerMdma_c0esrType
	Mdma_c0cr     registerMdma_c0crType
	Mdma_c0tcr    registerMdma_c0tcrType
	Mdma_c0bndtr  registerMdma_c0bndtrType
	Mdma_c0sar    registerMdma_c0sarType
	Mdma_c0dar    registerMdma_c0darType
	Mdma_c0brur   registerMdma_c0brurType
	Mdma_c0lar    registerMdma_c0larType
	Mdma_c0tbr    registerMdma_c0tbrType
	_             [4]byte
	Mdma_c0mar    registerMdma_c0marType
	Mdma_c0mdr    registerMdma_c0mdrType
	_             [8]byte
	Mdma_c1isr    registerMdma_c1isrType
	Mdma_c1ifcr   registerMdma_c1ifcrType
	Mdma_c1esr    registerMdma_c1esrType
	Mdma_c1cr     registerMdma_c1crType
	Mdma_c1tcr    registerMdma_c1tcrType
	Mdma_c1bndtr  registerMdma_c1bndtrType
	Mdma_c1sar    registerMdma_c1sarType
	Mdma_c1dar    registerMdma_c1darType
	Mdma_c1brur   registerMdma_c1brurType
	Mdma_c1lar    registerMdma_c1larType
	Mdma_c1tbr    registerMdma_c1tbrType
	_             [4]byte
	Mdma_c1mar    registerMdma_c1marType
	Mdma_c1mdr    registerMdma_c1mdrType
	_             [8]byte
	Mdma_c2isr    registerMdma_c2isrType
	Mdma_c2ifcr   registerMdma_c2ifcrType
	Mdma_c2esr    registerMdma_c2esrType
	Mdma_c2cr     registerMdma_c2crType
	Mdma_c2tcr    registerMdma_c2tcrType
	Mdma_c2bndtr  registerMdma_c2bndtrType
	Mdma_c2sar    registerMdma_c2sarType
	Mdma_c2dar    registerMdma_c2darType
	Mdma_c2brur   registerMdma_c2brurType
	Mdma_c2lar    registerMdma_c2larType
	Mdma_c2tbr    registerMdma_c2tbrType
	_             [4]byte
	Mdma_c2mar    registerMdma_c2marType
	Mdma_c2mdr    registerMdma_c2mdrType
	_             [8]byte
	Mdma_c3isr    registerMdma_c3isrType
	Mdma_c3ifcr   registerMdma_c3ifcrType
	Mdma_c3esr    registerMdma_c3esrType
	Mdma_c3cr     registerMdma_c3crType
	Mdma_c3tcr    registerMdma_c3tcrType
	Mdma_c3bndtr  registerMdma_c3bndtrType
	Mdma_c3sar    registerMdma_c3sarType
	Mdma_c3dar    registerMdma_c3darType
	Mdma_c3brur   registerMdma_c3brurType
	Mdma_c3lar    registerMdma_c3larType
	Mdma_c3tbr    registerMdma_c3tbrType
	_             [4]byte
	Mdma_c3mar    registerMdma_c3marType
	Mdma_c3mdr    registerMdma_c3mdrType
	_             [8]byte
	Mdma_c4isr    registerMdma_c4isrType
	Mdma_c4ifcr   registerMdma_c4ifcrType
	Mdma_c4esr    registerMdma_c4esrType
	Mdma_c4cr     registerMdma_c4crType
	Mdma_c4tcr    registerMdma_c4tcrType
	Mdma_c4bndtr  registerMdma_c4bndtrType
	Mdma_c4sar    registerMdma_c4sarType
	Mdma_c4dar    registerMdma_c4darType
	Mdma_c4brur   registerMdma_c4brurType
	Mdma_c4lar    registerMdma_c4larType
	Mdma_c4tbr    registerMdma_c4tbrType
	_             [4]byte
	Mdma_c4mar    registerMdma_c4marType
	Mdma_c4mdr    registerMdma_c4mdrType
	_             [8]byte
	Mdma_c5isr    registerMdma_c5isrType
	Mdma_c5ifcr   registerMdma_c5ifcrType
	Mdma_c5esr    registerMdma_c5esrType
	Mdma_c5cr     registerMdma_c5crType
	Mdma_c5tcr    registerMdma_c5tcrType
	Mdma_c5bndtr  registerMdma_c5bndtrType
	Mdma_c5sar    registerMdma_c5sarType
	Mdma_c5dar    registerMdma_c5darType
	Mdma_c5brur   registerMdma_c5brurType
	Mdma_c5lar    registerMdma_c5larType
	Mdma_c5tbr    registerMdma_c5tbrType
	_             [4]byte
	Mdma_c5mar    registerMdma_c5marType
	Mdma_c5mdr    registerMdma_c5mdrType
	_             [8]byte
	Mdma_c6isr    registerMdma_c6isrType
	Mdma_c6ifcr   registerMdma_c6ifcrType
	Mdma_c6esr    registerMdma_c6esrType
	Mdma_c6cr     registerMdma_c6crType
	Mdma_c6tcr    registerMdma_c6tcrType
	Mdma_c6bndtr  registerMdma_c6bndtrType
	Mdma_c6sar    registerMdma_c6sarType
	Mdma_c6dar    registerMdma_c6darType
	Mdma_c6brur   registerMdma_c6brurType
	Mdma_c6lar    registerMdma_c6larType
	Mdma_c6tbr    registerMdma_c6tbrType
	_             [4]byte
	Mdma_c6mar    registerMdma_c6marType
	Mdma_c6mdr    registerMdma_c6mdrType
	_             [8]byte
	Mdma_c7isr    registerMdma_c7isrType
	Mdma_c7ifcr   registerMdma_c7ifcrType
	Mdma_c7esr    registerMdma_c7esrType
	Mdma_c7cr     registerMdma_c7crType
	Mdma_c7tcr    registerMdma_c7tcrType
	Mdma_c7bndtr  registerMdma_c7bndtrType
	Mdma_c7sar    registerMdma_c7sarType
	Mdma_c7dar    registerMdma_c7darType
	Mdma_c7brur   registerMdma_c7brurType
	Mdma_c7lar    registerMdma_c7larType
	Mdma_c7tbr    registerMdma_c7tbrType
	_             [4]byte
	Mdma_c7mar    registerMdma_c7marType
	Mdma_c7mdr    registerMdma_c7mdrType
	_             [8]byte
	Mdma_c8isr    registerMdma_c8isrType
	Mdma_c8ifcr   registerMdma_c8ifcrType
	Mdma_c8esr    registerMdma_c8esrType
	Mdma_c8cr     registerMdma_c8crType
	Mdma_c8tcr    registerMdma_c8tcrType
	Mdma_c8bndtr  registerMdma_c8bndtrType
	Mdma_c8sar    registerMdma_c8sarType
	Mdma_c8dar    registerMdma_c8darType
	Mdma_c8brur   registerMdma_c8brurType
	Mdma_c8lar    registerMdma_c8larType
	Mdma_c8tbr    registerMdma_c8tbrType
	_             [4]byte
	Mdma_c8mar    registerMdma_c8marType
	Mdma_c8mdr    registerMdma_c8mdrType
	_             [8]byte
	Mdma_c9isr    registerMdma_c9isrType
	Mdma_c9ifcr   registerMdma_c9ifcrType
	Mdma_c9esr    registerMdma_c9esrType
	Mdma_c9cr     registerMdma_c9crType
	Mdma_c9tcr    registerMdma_c9tcrType
	Mdma_c9bndtr  registerMdma_c9bndtrType
	Mdma_c9sar    registerMdma_c9sarType
	Mdma_c9dar    registerMdma_c9darType
	Mdma_c9brur   registerMdma_c9brurType
	Mdma_c9lar    registerMdma_c9larType
	Mdma_c9tbr    registerMdma_c9tbrType
	_             [4]byte
	Mdma_c9mar    registerMdma_c9marType
	Mdma_c9mdr    registerMdma_c9mdrType
	_             [8]byte
	Mdma_c10isr   registerMdma_c10isrType
	Mdma_c10ifcr  registerMdma_c10ifcrType
	Mdma_c10esr   registerMdma_c10esrType
	Mdma_c10cr    registerMdma_c10crType
	Mdma_c10tcr   registerMdma_c10tcrType
	Mdma_c10bndtr registerMdma_c10bndtrType
	Mdma_c10sar   registerMdma_c10sarType
	Mdma_c10dar   registerMdma_c10darType
	Mdma_c10brur  registerMdma_c10brurType
	Mdma_c10lar   registerMdma_c10larType
	Mdma_c10tbr   registerMdma_c10tbrType
	_             [4]byte
	Mdma_c10mar   registerMdma_c10marType
	Mdma_c10mdr   registerMdma_c10mdrType
	_             [8]byte
	Mdma_c11isr   registerMdma_c11isrType
	Mdma_c11ifcr  registerMdma_c11ifcrType
	Mdma_c11esr   registerMdma_c11esrType
	Mdma_c11cr    registerMdma_c11crType
	Mdma_c11tcr   registerMdma_c11tcrType
	Mdma_c11bndtr registerMdma_c11bndtrType
	Mdma_c11sar   registerMdma_c11sarType
	Mdma_c11dar   registerMdma_c11darType
	Mdma_c11brur  registerMdma_c11brurType
	Mdma_c11lar   registerMdma_c11larType
	Mdma_c11tbr   registerMdma_c11tbrType
	_             [4]byte
	Mdma_c11mar   registerMdma_c11marType
	Mdma_c11mdr   registerMdma_c11mdrType
	_             [8]byte
	Mdma_c12isr   registerMdma_c12isrType
	Mdma_c12ifcr  registerMdma_c12ifcrType
	Mdma_c12esr   registerMdma_c12esrType
	Mdma_c12cr    registerMdma_c12crType
	Mdma_c12tcr   registerMdma_c12tcrType
	Mdma_c12bndtr registerMdma_c12bndtrType
	Mdma_c12sar   registerMdma_c12sarType
	Mdma_c12dar   registerMdma_c12darType
	Mdma_c12brur  registerMdma_c12brurType
	Mdma_c12lar   registerMdma_c12larType
	Mdma_c12tbr   registerMdma_c12tbrType
	_             [4]byte
	Mdma_c12mar   registerMdma_c12marType
	Mdma_c12mdr   registerMdma_c12mdrType
	_             [8]byte
	Mdma_c13isr   registerMdma_c13isrType
	Mdma_c13ifcr  registerMdma_c13ifcrType
	Mdma_c13esr   registerMdma_c13esrType
	Mdma_c13cr    registerMdma_c13crType
	Mdma_c13tcr   registerMdma_c13tcrType
	Mdma_c13bndtr registerMdma_c13bndtrType
	Mdma_c13sar   registerMdma_c13sarType
	Mdma_c13dar   registerMdma_c13darType
	Mdma_c13brur  registerMdma_c13brurType
	Mdma_c13lar   registerMdma_c13larType
	Mdma_c13tbr   registerMdma_c13tbrType
	_             [4]byte
	Mdma_c13mar   registerMdma_c13marType
	Mdma_c13mdr   registerMdma_c13mdrType
	_             [8]byte
	Mdma_c14isr   registerMdma_c14isrType
	Mdma_c14ifcr  registerMdma_c14ifcrType
	Mdma_c14esr   registerMdma_c14esrType
	Mdma_c14cr    registerMdma_c14crType
	Mdma_c14tcr   registerMdma_c14tcrType
	Mdma_c14bndtr registerMdma_c14bndtrType
	Mdma_c14sar   registerMdma_c14sarType
	Mdma_c14dar   registerMdma_c14darType
	Mdma_c14brur  registerMdma_c14brurType
	Mdma_c14lar   registerMdma_c14larType
	Mdma_c14tbr   registerMdma_c14tbrType
	_             [4]byte
	Mdma_c14mar   registerMdma_c14marType
	Mdma_c14mdr   registerMdma_c14mdrType
	_             [8]byte
	Mdma_c15isr   registerMdma_c15isrType
	Mdma_c15ifcr  registerMdma_c15ifcrType
	Mdma_c15esr   registerMdma_c15esrType
	Mdma_c15cr    registerMdma_c15crType
	Mdma_c15tcr   registerMdma_c15tcrType
	Mdma_c15bndtr registerMdma_c15bndtrType
	Mdma_c15sar   registerMdma_c15sarType
	Mdma_c15dar   registerMdma_c15darType
	Mdma_c15brur  registerMdma_c15brurType
	Mdma_c15lar   registerMdma_c15larType
	Mdma_c15tbr   registerMdma_c15tbrType
	_             [4]byte
	Mdma_c15mar   registerMdma_c15marType
	Mdma_c15mdr   registerMdma_c15mdrType
}

// registerMdma_gisr0Type MDMA Global Interrupt/Status Register
type registerMdma_gisr0Type uint32

const (
	RegisterMdma_gisr0FieldGif0Shift = 0
	RegisterMdma_gisr0FieldGif0Mask  = 0x1
)

// GetGif0 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif0Mask) != 0
}

// SetGif0 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif0Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif1Shift = 1
	RegisterMdma_gisr0FieldGif1Mask  = 0x2
)

// GetGif1 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif1Mask) != 0
}

// SetGif1 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif1Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif2Shift = 2
	RegisterMdma_gisr0FieldGif2Mask  = 0x4
)

// GetGif2 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif2Mask) != 0
}

// SetGif2 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif2Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif3Shift = 3
	RegisterMdma_gisr0FieldGif3Mask  = 0x8
)

// GetGif3 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif3Mask) != 0
}

// SetGif3 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif3Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif4Shift = 4
	RegisterMdma_gisr0FieldGif4Mask  = 0x10
)

// GetGif4 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif4Mask) != 0
}

// SetGif4 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif4Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif5Shift = 5
	RegisterMdma_gisr0FieldGif5Mask  = 0x20
)

// GetGif5 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif5Mask) != 0
}

// SetGif5 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif5Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif6Shift = 6
	RegisterMdma_gisr0FieldGif6Mask  = 0x40
)

// GetGif6 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif6Mask) != 0
}

// SetGif6 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif6Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif7Shift = 7
	RegisterMdma_gisr0FieldGif7Mask  = 0x80
)

// GetGif7 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif7Mask) != 0
}

// SetGif7 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif7Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif8Shift = 8
	RegisterMdma_gisr0FieldGif8Mask  = 0x100
)

// GetGif8 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif8Mask) != 0
}

// SetGif8 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif8Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif9Shift = 9
	RegisterMdma_gisr0FieldGif9Mask  = 0x200
)

// GetGif9 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif9Mask) != 0
}

// SetGif9 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif9Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif10Shift = 10
	RegisterMdma_gisr0FieldGif10Mask  = 0x400
)

// GetGif10 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif10Mask) != 0
}

// SetGif10 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif10Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif11Shift = 11
	RegisterMdma_gisr0FieldGif11Mask  = 0x800
)

// GetGif11 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif11Mask) != 0
}

// SetGif11 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif11Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif12Shift = 12
	RegisterMdma_gisr0FieldGif12Mask  = 0x1000
)

// GetGif12 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif12Mask) != 0
}

// SetGif12 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif12Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif13Shift = 13
	RegisterMdma_gisr0FieldGif13Mask  = 0x2000
)

// GetGif13 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif13Mask) != 0
}

// SetGif13 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif13Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif14Shift = 14
	RegisterMdma_gisr0FieldGif14Mask  = 0x4000
)

// GetGif14 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif14Mask) != 0
}

// SetGif14 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif14Mask)
	}
}

const (
	RegisterMdma_gisr0FieldGif15Shift = 15
	RegisterMdma_gisr0FieldGif15Mask  = 0x8000
)

// GetGif15 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) GetGif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_gisr0FieldGif15Mask) != 0
}

// SetGif15 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *registerMdma_gisr0Type) SetGif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_gisr0FieldGif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_gisr0FieldGif15Mask)
	}
}

// registerMdma_c0isrType MDMA channel x interrupt/status register
type registerMdma_c0isrType uint32

const (
	RegisterMdma_c0isrFieldTeif0Shift = 0
	RegisterMdma_c0isrFieldTeif0Mask  = 0x1
)

// GetTeif0 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c0isrType) GetTeif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0isrFieldTeif0Mask) != 0
}

// SetTeif0 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c0isrType) SetTeif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0isrFieldTeif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0isrFieldTeif0Mask)
	}
}

const (
	RegisterMdma_c0isrFieldCtcif0Shift = 1
	RegisterMdma_c0isrFieldCtcif0Mask  = 0x2
)

// GetCtcif0 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c0isrType) GetCtcif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0isrFieldCtcif0Mask) != 0
}

// SetCtcif0 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c0isrType) SetCtcif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0isrFieldCtcif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0isrFieldCtcif0Mask)
	}
}

const (
	RegisterMdma_c0isrFieldBrtif0Shift = 2
	RegisterMdma_c0isrFieldBrtif0Mask  = 0x4
)

// GetBrtif0 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c0isrType) GetBrtif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0isrFieldBrtif0Mask) != 0
}

// SetBrtif0 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c0isrType) SetBrtif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0isrFieldBrtif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0isrFieldBrtif0Mask)
	}
}

const (
	RegisterMdma_c0isrFieldBtif0Shift = 3
	RegisterMdma_c0isrFieldBtif0Mask  = 0x8
)

// GetBtif0 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c0isrType) GetBtif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0isrFieldBtif0Mask) != 0
}

// SetBtif0 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c0isrType) SetBtif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0isrFieldBtif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0isrFieldBtif0Mask)
	}
}

const (
	RegisterMdma_c0isrFieldTcif0Shift = 4
	RegisterMdma_c0isrFieldTcif0Mask  = 0x10
)

// GetTcif0 channel x buffer transfer complete
func (r *registerMdma_c0isrType) GetTcif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0isrFieldTcif0Mask) != 0
}

// SetTcif0 channel x buffer transfer complete
func (r *registerMdma_c0isrType) SetTcif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0isrFieldTcif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0isrFieldTcif0Mask)
	}
}

const (
	RegisterMdma_c0isrFieldCrqa0Shift = 16
	RegisterMdma_c0isrFieldCrqa0Mask  = 0x10000
)

// GetCrqa0 channel x request active flag
func (r *registerMdma_c0isrType) GetCrqa0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0isrFieldCrqa0Mask) != 0
}

// SetCrqa0 channel x request active flag
func (r *registerMdma_c0isrType) SetCrqa0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0isrFieldCrqa0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0isrFieldCrqa0Mask)
	}
}

// registerMdma_c0ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c0ifcrType uint32

const (
	RegisterMdma_c0ifcrFieldCteif0Shift = 0
	RegisterMdma_c0ifcrFieldCteif0Mask  = 0x1
)

// GetCteif0 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c0ifcrType) GetCteif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0ifcrFieldCteif0Mask) != 0
}

// SetCteif0 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c0ifcrType) SetCteif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0ifcrFieldCteif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0ifcrFieldCteif0Mask)
	}
}

const (
	RegisterMdma_c0ifcrFieldCctcif0Shift = 1
	RegisterMdma_c0ifcrFieldCctcif0Mask  = 0x2
)

// GetCctcif0 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c0ifcrType) GetCctcif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0ifcrFieldCctcif0Mask) != 0
}

// SetCctcif0 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c0ifcrType) SetCctcif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0ifcrFieldCctcif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0ifcrFieldCctcif0Mask)
	}
}

const (
	RegisterMdma_c0ifcrFieldCbrtif0Shift = 2
	RegisterMdma_c0ifcrFieldCbrtif0Mask  = 0x4
)

// GetCbrtif0 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c0ifcrType) GetCbrtif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0ifcrFieldCbrtif0Mask) != 0
}

// SetCbrtif0 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c0ifcrType) SetCbrtif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0ifcrFieldCbrtif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0ifcrFieldCbrtif0Mask)
	}
}

const (
	RegisterMdma_c0ifcrFieldCbtif0Shift = 3
	RegisterMdma_c0ifcrFieldCbtif0Mask  = 0x8
)

// GetCbtif0 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c0ifcrType) GetCbtif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0ifcrFieldCbtif0Mask) != 0
}

// SetCbtif0 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c0ifcrType) SetCbtif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0ifcrFieldCbtif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0ifcrFieldCbtif0Mask)
	}
}

const (
	RegisterMdma_c0ifcrFieldCltcif0Shift = 4
	RegisterMdma_c0ifcrFieldCltcif0Mask  = 0x10
)

// GetCltcif0 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c0ifcrType) GetCltcif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0ifcrFieldCltcif0Mask) != 0
}

// SetCltcif0 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c0ifcrType) SetCltcif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0ifcrFieldCltcif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0ifcrFieldCltcif0Mask)
	}
}

// registerMdma_c0esrType MDMA Channel x error status register
type registerMdma_c0esrType uint32

const (
	RegisterMdma_c0esrFieldTeaShift = 0
	RegisterMdma_c0esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c0esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0esrFieldTeaMask) >> RegisterMdma_c0esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c0esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c0esrFieldTeaShift))
}

const (
	RegisterMdma_c0esrFieldTedShift = 7
	RegisterMdma_c0esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c0esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c0esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0esrFieldTedMask)
	}
}

const (
	RegisterMdma_c0esrFieldTeldShift = 8
	RegisterMdma_c0esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c0esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c0esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c0esrFieldTemdShift = 9
	RegisterMdma_c0esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c0esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c0esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c0esrFieldAseShift = 10
	RegisterMdma_c0esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c0esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c0esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0esrFieldAseMask)
	}
}

const (
	RegisterMdma_c0esrFieldBseShift = 11
	RegisterMdma_c0esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c0esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c0esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0esrFieldBseMask)
	}
}

// registerMdma_c0crType This register is used to control the concerned channel.
type registerMdma_c0crType uint32

const (
	RegisterMdma_c0crFieldEnShift = 0
	RegisterMdma_c0crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c0crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c0crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0crFieldEnMask)
	}
}

const (
	RegisterMdma_c0crFieldTeieShift = 1
	RegisterMdma_c0crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c0crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c0crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0crFieldTeieMask)
	}
}

const (
	RegisterMdma_c0crFieldCtcieShift = 2
	RegisterMdma_c0crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c0crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c0crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c0crFieldBrtieShift = 3
	RegisterMdma_c0crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c0crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c0crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c0crFieldBtieShift = 4
	RegisterMdma_c0crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c0crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c0crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0crFieldBtieMask)
	}
}

const (
	RegisterMdma_c0crFieldTcieShift = 5
	RegisterMdma_c0crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c0crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c0crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0crFieldTcieMask)
	}
}

const (
	RegisterMdma_c0crFieldPlShift = 6
	RegisterMdma_c0crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c0crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0crFieldPlMask) >> RegisterMdma_c0crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c0crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0crFieldPlMask)|(uint32(value)<<RegisterMdma_c0crFieldPlShift))
}

const (
	RegisterMdma_c0crFieldBexShift = 12
	RegisterMdma_c0crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c0crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c0crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0crFieldBexMask)
	}
}

const (
	RegisterMdma_c0crFieldHexShift = 13
	RegisterMdma_c0crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c0crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c0crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0crFieldHexMask)
	}
}

const (
	RegisterMdma_c0crFieldWexShift = 14
	RegisterMdma_c0crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c0crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c0crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0crFieldWexMask)
	}
}

const (
	RegisterMdma_c0crFieldSwrqShift = 16
	RegisterMdma_c0crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c0crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0crFieldSwrqMask)
	}
}

// registerMdma_c0tcrType This register is used to configure the concerned channel.
type registerMdma_c0tcrType uint32

const (
	RegisterMdma_c0tcrFieldSincShift = 0
	RegisterMdma_c0tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c0tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldSincMask) >> RegisterMdma_c0tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c0tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c0tcrFieldSincShift))
}

const (
	RegisterMdma_c0tcrFieldDincShift = 2
	RegisterMdma_c0tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c0tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldDincMask) >> RegisterMdma_c0tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c0tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c0tcrFieldDincShift))
}

const (
	RegisterMdma_c0tcrFieldSsizeShift = 4
	RegisterMdma_c0tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c0tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldSsizeMask) >> RegisterMdma_c0tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c0tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c0tcrFieldSsizeShift))
}

const (
	RegisterMdma_c0tcrFieldDsizeShift = 6
	RegisterMdma_c0tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c0tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldDsizeMask) >> RegisterMdma_c0tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c0tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c0tcrFieldDsizeShift))
}

const (
	RegisterMdma_c0tcrFieldSincosShift = 8
	RegisterMdma_c0tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c0tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldSincosMask) >> RegisterMdma_c0tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c0tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c0tcrFieldSincosShift))
}

const (
	RegisterMdma_c0tcrFieldDincosShift = 10
	RegisterMdma_c0tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c0tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldDincosMask) >> RegisterMdma_c0tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c0tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c0tcrFieldDincosShift))
}

const (
	RegisterMdma_c0tcrFieldSburstShift = 12
	RegisterMdma_c0tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c0tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldSburstMask) >> RegisterMdma_c0tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c0tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c0tcrFieldSburstShift))
}

const (
	RegisterMdma_c0tcrFieldDburstShift = 15
	RegisterMdma_c0tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c0tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldDburstMask) >> RegisterMdma_c0tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c0tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c0tcrFieldDburstShift))
}

const (
	RegisterMdma_c0tcrFieldTlenShift = 18
	RegisterMdma_c0tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c0tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldTlenMask) >> RegisterMdma_c0tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c0tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c0tcrFieldTlenShift))
}

const (
	RegisterMdma_c0tcrFieldPkeShift = 25
	RegisterMdma_c0tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c0tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c0tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c0tcrFieldPamShift = 26
	RegisterMdma_c0tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c0tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldPamMask) >> RegisterMdma_c0tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c0tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c0tcrFieldPamShift))
}

const (
	RegisterMdma_c0tcrFieldTrgmShift = 28
	RegisterMdma_c0tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c0tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldTrgmMask) >> RegisterMdma_c0tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c0tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c0tcrFieldTrgmShift))
}

const (
	RegisterMdma_c0tcrFieldSwrmShift = 30
	RegisterMdma_c0tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c0tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c0tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c0tcrFieldBwmShift = 31
	RegisterMdma_c0tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c0tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c0tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tcrFieldBwmMask)
	}
}

// registerMdma_c0bndtrType MDMA Channel x block number of data register
type registerMdma_c0bndtrType uint32

const (
	RegisterMdma_c0bndtrFieldBndtShift = 0
	RegisterMdma_c0bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c0bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0bndtrFieldBndtMask) >> RegisterMdma_c0bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c0bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c0bndtrFieldBndtShift))
}

const (
	RegisterMdma_c0bndtrFieldBrsumShift = 18
	RegisterMdma_c0bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c0bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c0bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c0bndtrFieldBrdumShift = 19
	RegisterMdma_c0bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c0bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c0bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c0bndtrFieldBrcShift = 20
	RegisterMdma_c0bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c0bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0bndtrFieldBrcMask) >> RegisterMdma_c0bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c0bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c0bndtrFieldBrcShift))
}

// registerMdma_c0sarType MDMA channel x source address register
type registerMdma_c0sarType uint32

const (
	RegisterMdma_c0sarFieldSarShift = 0
	RegisterMdma_c0sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c0sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0sarFieldSarMask) >> RegisterMdma_c0sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c0sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0sarFieldSarMask)|(uint32(value)<<RegisterMdma_c0sarFieldSarShift))
}

// registerMdma_c0darType MDMA channel x destination address register
type registerMdma_c0darType uint32

const (
	RegisterMdma_c0darFieldDarShift = 0
	RegisterMdma_c0darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c0darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0darFieldDarMask) >> RegisterMdma_c0darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c0darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0darFieldDarMask)|(uint32(value)<<RegisterMdma_c0darFieldDarShift))
}

// registerMdma_c0brurType MDMA channel x Block Repeat address Update register
type registerMdma_c0brurType uint32

const (
	RegisterMdma_c0brurFieldSuvShift = 0
	RegisterMdma_c0brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c0brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0brurFieldSuvMask) >> RegisterMdma_c0brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c0brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c0brurFieldSuvShift))
}

const (
	RegisterMdma_c0brurFieldDuvShift = 16
	RegisterMdma_c0brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c0brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0brurFieldDuvMask) >> RegisterMdma_c0brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c0brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c0brurFieldDuvShift))
}

// registerMdma_c0larType MDMA channel x Link Address register
type registerMdma_c0larType uint32

const (
	RegisterMdma_c0larFieldLarShift = 0
	RegisterMdma_c0larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c0larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0larFieldLarMask) >> RegisterMdma_c0larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c0larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0larFieldLarMask)|(uint32(value)<<RegisterMdma_c0larFieldLarShift))
}

// registerMdma_c0tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c0tbrType uint32

const (
	RegisterMdma_c0tbrFieldTselShift = 0
	RegisterMdma_c0tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c0tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tbrFieldTselMask) >> RegisterMdma_c0tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c0tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c0tbrFieldTselShift))
}

const (
	RegisterMdma_c0tbrFieldSbusShift = 16
	RegisterMdma_c0tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c0tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c0tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c0tbrFieldDbusShift = 17
	RegisterMdma_c0tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c0tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c0tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c0tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0tbrFieldDbusMask)
	}
}

// registerMdma_c0marType MDMA channel x Mask address register
type registerMdma_c0marType uint32

const (
	RegisterMdma_c0marFieldMarShift = 0
	RegisterMdma_c0marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c0marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0marFieldMarMask) >> RegisterMdma_c0marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c0marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0marFieldMarMask)|(uint32(value)<<RegisterMdma_c0marFieldMarShift))
}

// registerMdma_c0mdrType MDMA channel x Mask Data register
type registerMdma_c0mdrType uint32

const (
	RegisterMdma_c0mdrFieldMdrShift = 0
	RegisterMdma_c0mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c0mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c0mdrFieldMdrMask) >> RegisterMdma_c0mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c0mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c0mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c0mdrFieldMdrShift))
}

// registerMdma_c1isrType MDMA channel x interrupt/status register
type registerMdma_c1isrType uint32

const (
	RegisterMdma_c1isrFieldTeif1Shift = 0
	RegisterMdma_c1isrFieldTeif1Mask  = 0x1
)

// GetTeif1 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c1isrType) GetTeif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1isrFieldTeif1Mask) != 0
}

// SetTeif1 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c1isrType) SetTeif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1isrFieldTeif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1isrFieldTeif1Mask)
	}
}

const (
	RegisterMdma_c1isrFieldCtcif1Shift = 1
	RegisterMdma_c1isrFieldCtcif1Mask  = 0x2
)

// GetCtcif1 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c1isrType) GetCtcif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1isrFieldCtcif1Mask) != 0
}

// SetCtcif1 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c1isrType) SetCtcif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1isrFieldCtcif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1isrFieldCtcif1Mask)
	}
}

const (
	RegisterMdma_c1isrFieldBrtif1Shift = 2
	RegisterMdma_c1isrFieldBrtif1Mask  = 0x4
)

// GetBrtif1 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c1isrType) GetBrtif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1isrFieldBrtif1Mask) != 0
}

// SetBrtif1 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c1isrType) SetBrtif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1isrFieldBrtif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1isrFieldBrtif1Mask)
	}
}

const (
	RegisterMdma_c1isrFieldBtif1Shift = 3
	RegisterMdma_c1isrFieldBtif1Mask  = 0x8
)

// GetBtif1 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c1isrType) GetBtif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1isrFieldBtif1Mask) != 0
}

// SetBtif1 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c1isrType) SetBtif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1isrFieldBtif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1isrFieldBtif1Mask)
	}
}

const (
	RegisterMdma_c1isrFieldTcif1Shift = 4
	RegisterMdma_c1isrFieldTcif1Mask  = 0x10
)

// GetTcif1 channel x buffer transfer complete
func (r *registerMdma_c1isrType) GetTcif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1isrFieldTcif1Mask) != 0
}

// SetTcif1 channel x buffer transfer complete
func (r *registerMdma_c1isrType) SetTcif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1isrFieldTcif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1isrFieldTcif1Mask)
	}
}

const (
	RegisterMdma_c1isrFieldCrqa1Shift = 16
	RegisterMdma_c1isrFieldCrqa1Mask  = 0x10000
)

// GetCrqa1 channel x request active flag
func (r *registerMdma_c1isrType) GetCrqa1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1isrFieldCrqa1Mask) != 0
}

// SetCrqa1 channel x request active flag
func (r *registerMdma_c1isrType) SetCrqa1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1isrFieldCrqa1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1isrFieldCrqa1Mask)
	}
}

// registerMdma_c1ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c1ifcrType uint32

const (
	RegisterMdma_c1ifcrFieldCteif1Shift = 0
	RegisterMdma_c1ifcrFieldCteif1Mask  = 0x1
)

// GetCteif1 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c1ifcrType) GetCteif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1ifcrFieldCteif1Mask) != 0
}

// SetCteif1 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c1ifcrType) SetCteif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1ifcrFieldCteif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1ifcrFieldCteif1Mask)
	}
}

const (
	RegisterMdma_c1ifcrFieldCctcif1Shift = 1
	RegisterMdma_c1ifcrFieldCctcif1Mask  = 0x2
)

// GetCctcif1 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c1ifcrType) GetCctcif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1ifcrFieldCctcif1Mask) != 0
}

// SetCctcif1 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c1ifcrType) SetCctcif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1ifcrFieldCctcif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1ifcrFieldCctcif1Mask)
	}
}

const (
	RegisterMdma_c1ifcrFieldCbrtif1Shift = 2
	RegisterMdma_c1ifcrFieldCbrtif1Mask  = 0x4
)

// GetCbrtif1 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c1ifcrType) GetCbrtif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1ifcrFieldCbrtif1Mask) != 0
}

// SetCbrtif1 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c1ifcrType) SetCbrtif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1ifcrFieldCbrtif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1ifcrFieldCbrtif1Mask)
	}
}

const (
	RegisterMdma_c1ifcrFieldCbtif1Shift = 3
	RegisterMdma_c1ifcrFieldCbtif1Mask  = 0x8
)

// GetCbtif1 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c1ifcrType) GetCbtif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1ifcrFieldCbtif1Mask) != 0
}

// SetCbtif1 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c1ifcrType) SetCbtif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1ifcrFieldCbtif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1ifcrFieldCbtif1Mask)
	}
}

const (
	RegisterMdma_c1ifcrFieldCltcif1Shift = 4
	RegisterMdma_c1ifcrFieldCltcif1Mask  = 0x10
)

// GetCltcif1 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c1ifcrType) GetCltcif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1ifcrFieldCltcif1Mask) != 0
}

// SetCltcif1 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c1ifcrType) SetCltcif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1ifcrFieldCltcif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1ifcrFieldCltcif1Mask)
	}
}

// registerMdma_c1esrType MDMA Channel x error status register
type registerMdma_c1esrType uint32

const (
	RegisterMdma_c1esrFieldTeaShift = 0
	RegisterMdma_c1esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c1esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1esrFieldTeaMask) >> RegisterMdma_c1esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c1esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c1esrFieldTeaShift))
}

const (
	RegisterMdma_c1esrFieldTedShift = 7
	RegisterMdma_c1esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c1esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c1esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1esrFieldTedMask)
	}
}

const (
	RegisterMdma_c1esrFieldTeldShift = 8
	RegisterMdma_c1esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c1esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c1esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c1esrFieldTemdShift = 9
	RegisterMdma_c1esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c1esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c1esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c1esrFieldAseShift = 10
	RegisterMdma_c1esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c1esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c1esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1esrFieldAseMask)
	}
}

const (
	RegisterMdma_c1esrFieldBseShift = 11
	RegisterMdma_c1esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c1esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c1esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1esrFieldBseMask)
	}
}

// registerMdma_c1crType This register is used to control the concerned channel.
type registerMdma_c1crType uint32

const (
	RegisterMdma_c1crFieldEnShift = 0
	RegisterMdma_c1crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c1crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c1crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1crFieldEnMask)
	}
}

const (
	RegisterMdma_c1crFieldTeieShift = 1
	RegisterMdma_c1crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c1crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c1crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1crFieldTeieMask)
	}
}

const (
	RegisterMdma_c1crFieldCtcieShift = 2
	RegisterMdma_c1crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c1crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c1crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c1crFieldBrtieShift = 3
	RegisterMdma_c1crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c1crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c1crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c1crFieldBtieShift = 4
	RegisterMdma_c1crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c1crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c1crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1crFieldBtieMask)
	}
}

const (
	RegisterMdma_c1crFieldTcieShift = 5
	RegisterMdma_c1crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c1crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c1crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1crFieldTcieMask)
	}
}

const (
	RegisterMdma_c1crFieldPlShift = 6
	RegisterMdma_c1crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c1crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1crFieldPlMask) >> RegisterMdma_c1crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c1crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1crFieldPlMask)|(uint32(value)<<RegisterMdma_c1crFieldPlShift))
}

const (
	RegisterMdma_c1crFieldBexShift = 12
	RegisterMdma_c1crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c1crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c1crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1crFieldBexMask)
	}
}

const (
	RegisterMdma_c1crFieldHexShift = 13
	RegisterMdma_c1crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c1crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c1crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1crFieldHexMask)
	}
}

const (
	RegisterMdma_c1crFieldWexShift = 14
	RegisterMdma_c1crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c1crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c1crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1crFieldWexMask)
	}
}

const (
	RegisterMdma_c1crFieldSwrqShift = 16
	RegisterMdma_c1crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c1crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1crFieldSwrqMask)
	}
}

// registerMdma_c1tcrType This register is used to configure the concerned channel.
type registerMdma_c1tcrType uint32

const (
	RegisterMdma_c1tcrFieldSincShift = 0
	RegisterMdma_c1tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c1tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldSincMask) >> RegisterMdma_c1tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c1tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c1tcrFieldSincShift))
}

const (
	RegisterMdma_c1tcrFieldDincShift = 2
	RegisterMdma_c1tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c1tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldDincMask) >> RegisterMdma_c1tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c1tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c1tcrFieldDincShift))
}

const (
	RegisterMdma_c1tcrFieldSsizeShift = 4
	RegisterMdma_c1tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c1tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldSsizeMask) >> RegisterMdma_c1tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c1tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c1tcrFieldSsizeShift))
}

const (
	RegisterMdma_c1tcrFieldDsizeShift = 6
	RegisterMdma_c1tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c1tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldDsizeMask) >> RegisterMdma_c1tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c1tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c1tcrFieldDsizeShift))
}

const (
	RegisterMdma_c1tcrFieldSincosShift = 8
	RegisterMdma_c1tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c1tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldSincosMask) >> RegisterMdma_c1tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c1tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c1tcrFieldSincosShift))
}

const (
	RegisterMdma_c1tcrFieldDincosShift = 10
	RegisterMdma_c1tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c1tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldDincosMask) >> RegisterMdma_c1tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c1tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c1tcrFieldDincosShift))
}

const (
	RegisterMdma_c1tcrFieldSburstShift = 12
	RegisterMdma_c1tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c1tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldSburstMask) >> RegisterMdma_c1tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c1tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c1tcrFieldSburstShift))
}

const (
	RegisterMdma_c1tcrFieldDburstShift = 15
	RegisterMdma_c1tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c1tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldDburstMask) >> RegisterMdma_c1tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c1tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c1tcrFieldDburstShift))
}

const (
	RegisterMdma_c1tcrFieldTlenShift = 18
	RegisterMdma_c1tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c1tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldTlenMask) >> RegisterMdma_c1tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c1tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c1tcrFieldTlenShift))
}

const (
	RegisterMdma_c1tcrFieldPkeShift = 25
	RegisterMdma_c1tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c1tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c1tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c1tcrFieldPamShift = 26
	RegisterMdma_c1tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c1tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldPamMask) >> RegisterMdma_c1tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c1tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c1tcrFieldPamShift))
}

const (
	RegisterMdma_c1tcrFieldTrgmShift = 28
	RegisterMdma_c1tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c1tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldTrgmMask) >> RegisterMdma_c1tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c1tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c1tcrFieldTrgmShift))
}

const (
	RegisterMdma_c1tcrFieldSwrmShift = 30
	RegisterMdma_c1tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c1tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c1tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c1tcrFieldBwmShift = 31
	RegisterMdma_c1tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c1tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c1tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tcrFieldBwmMask)
	}
}

// registerMdma_c1bndtrType MDMA Channel x block number of data register
type registerMdma_c1bndtrType uint32

const (
	RegisterMdma_c1bndtrFieldBndtShift = 0
	RegisterMdma_c1bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c1bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1bndtrFieldBndtMask) >> RegisterMdma_c1bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c1bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c1bndtrFieldBndtShift))
}

const (
	RegisterMdma_c1bndtrFieldBrsumShift = 18
	RegisterMdma_c1bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c1bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c1bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c1bndtrFieldBrdumShift = 19
	RegisterMdma_c1bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c1bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c1bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c1bndtrFieldBrcShift = 20
	RegisterMdma_c1bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c1bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1bndtrFieldBrcMask) >> RegisterMdma_c1bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c1bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c1bndtrFieldBrcShift))
}

// registerMdma_c1sarType MDMA channel x source address register
type registerMdma_c1sarType uint32

const (
	RegisterMdma_c1sarFieldSarShift = 0
	RegisterMdma_c1sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c1sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1sarFieldSarMask) >> RegisterMdma_c1sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c1sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1sarFieldSarMask)|(uint32(value)<<RegisterMdma_c1sarFieldSarShift))
}

// registerMdma_c1darType MDMA channel x destination address register
type registerMdma_c1darType uint32

const (
	RegisterMdma_c1darFieldDarShift = 0
	RegisterMdma_c1darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c1darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1darFieldDarMask) >> RegisterMdma_c1darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c1darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1darFieldDarMask)|(uint32(value)<<RegisterMdma_c1darFieldDarShift))
}

// registerMdma_c1brurType MDMA channel x Block Repeat address Update register
type registerMdma_c1brurType uint32

const (
	RegisterMdma_c1brurFieldSuvShift = 0
	RegisterMdma_c1brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c1brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1brurFieldSuvMask) >> RegisterMdma_c1brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c1brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c1brurFieldSuvShift))
}

const (
	RegisterMdma_c1brurFieldDuvShift = 16
	RegisterMdma_c1brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c1brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1brurFieldDuvMask) >> RegisterMdma_c1brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c1brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c1brurFieldDuvShift))
}

// registerMdma_c1larType MDMA channel x Link Address register
type registerMdma_c1larType uint32

const (
	RegisterMdma_c1larFieldLarShift = 0
	RegisterMdma_c1larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c1larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1larFieldLarMask) >> RegisterMdma_c1larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c1larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1larFieldLarMask)|(uint32(value)<<RegisterMdma_c1larFieldLarShift))
}

// registerMdma_c1tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c1tbrType uint32

const (
	RegisterMdma_c1tbrFieldTselShift = 0
	RegisterMdma_c1tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c1tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tbrFieldTselMask) >> RegisterMdma_c1tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c1tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c1tbrFieldTselShift))
}

const (
	RegisterMdma_c1tbrFieldSbusShift = 16
	RegisterMdma_c1tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c1tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c1tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c1tbrFieldDbusShift = 17
	RegisterMdma_c1tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c1tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c1tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c1tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1tbrFieldDbusMask)
	}
}

// registerMdma_c1marType MDMA channel x Mask address register
type registerMdma_c1marType uint32

const (
	RegisterMdma_c1marFieldMarShift = 0
	RegisterMdma_c1marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c1marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1marFieldMarMask) >> RegisterMdma_c1marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c1marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1marFieldMarMask)|(uint32(value)<<RegisterMdma_c1marFieldMarShift))
}

// registerMdma_c1mdrType MDMA channel x Mask Data register
type registerMdma_c1mdrType uint32

const (
	RegisterMdma_c1mdrFieldMdrShift = 0
	RegisterMdma_c1mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c1mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c1mdrFieldMdrMask) >> RegisterMdma_c1mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c1mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c1mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c1mdrFieldMdrShift))
}

// registerMdma_c2isrType MDMA channel x interrupt/status register
type registerMdma_c2isrType uint32

const (
	RegisterMdma_c2isrFieldTeif2Shift = 0
	RegisterMdma_c2isrFieldTeif2Mask  = 0x1
)

// GetTeif2 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c2isrType) GetTeif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2isrFieldTeif2Mask) != 0
}

// SetTeif2 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c2isrType) SetTeif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2isrFieldTeif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2isrFieldTeif2Mask)
	}
}

const (
	RegisterMdma_c2isrFieldCtcif2Shift = 1
	RegisterMdma_c2isrFieldCtcif2Mask  = 0x2
)

// GetCtcif2 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c2isrType) GetCtcif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2isrFieldCtcif2Mask) != 0
}

// SetCtcif2 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c2isrType) SetCtcif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2isrFieldCtcif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2isrFieldCtcif2Mask)
	}
}

const (
	RegisterMdma_c2isrFieldBrtif2Shift = 2
	RegisterMdma_c2isrFieldBrtif2Mask  = 0x4
)

// GetBrtif2 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c2isrType) GetBrtif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2isrFieldBrtif2Mask) != 0
}

// SetBrtif2 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c2isrType) SetBrtif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2isrFieldBrtif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2isrFieldBrtif2Mask)
	}
}

const (
	RegisterMdma_c2isrFieldBtif2Shift = 3
	RegisterMdma_c2isrFieldBtif2Mask  = 0x8
)

// GetBtif2 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c2isrType) GetBtif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2isrFieldBtif2Mask) != 0
}

// SetBtif2 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c2isrType) SetBtif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2isrFieldBtif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2isrFieldBtif2Mask)
	}
}

const (
	RegisterMdma_c2isrFieldTcif2Shift = 4
	RegisterMdma_c2isrFieldTcif2Mask  = 0x10
)

// GetTcif2 channel x buffer transfer complete
func (r *registerMdma_c2isrType) GetTcif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2isrFieldTcif2Mask) != 0
}

// SetTcif2 channel x buffer transfer complete
func (r *registerMdma_c2isrType) SetTcif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2isrFieldTcif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2isrFieldTcif2Mask)
	}
}

const (
	RegisterMdma_c2isrFieldCrqa2Shift = 16
	RegisterMdma_c2isrFieldCrqa2Mask  = 0x10000
)

// GetCrqa2 channel x request active flag
func (r *registerMdma_c2isrType) GetCrqa2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2isrFieldCrqa2Mask) != 0
}

// SetCrqa2 channel x request active flag
func (r *registerMdma_c2isrType) SetCrqa2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2isrFieldCrqa2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2isrFieldCrqa2Mask)
	}
}

// registerMdma_c2ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c2ifcrType uint32

const (
	RegisterMdma_c2ifcrFieldCteif2Shift = 0
	RegisterMdma_c2ifcrFieldCteif2Mask  = 0x1
)

// GetCteif2 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c2ifcrType) GetCteif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2ifcrFieldCteif2Mask) != 0
}

// SetCteif2 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c2ifcrType) SetCteif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2ifcrFieldCteif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2ifcrFieldCteif2Mask)
	}
}

const (
	RegisterMdma_c2ifcrFieldCctcif2Shift = 1
	RegisterMdma_c2ifcrFieldCctcif2Mask  = 0x2
)

// GetCctcif2 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c2ifcrType) GetCctcif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2ifcrFieldCctcif2Mask) != 0
}

// SetCctcif2 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c2ifcrType) SetCctcif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2ifcrFieldCctcif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2ifcrFieldCctcif2Mask)
	}
}

const (
	RegisterMdma_c2ifcrFieldCbrtif2Shift = 2
	RegisterMdma_c2ifcrFieldCbrtif2Mask  = 0x4
)

// GetCbrtif2 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c2ifcrType) GetCbrtif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2ifcrFieldCbrtif2Mask) != 0
}

// SetCbrtif2 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c2ifcrType) SetCbrtif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2ifcrFieldCbrtif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2ifcrFieldCbrtif2Mask)
	}
}

const (
	RegisterMdma_c2ifcrFieldCbtif2Shift = 3
	RegisterMdma_c2ifcrFieldCbtif2Mask  = 0x8
)

// GetCbtif2 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c2ifcrType) GetCbtif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2ifcrFieldCbtif2Mask) != 0
}

// SetCbtif2 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c2ifcrType) SetCbtif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2ifcrFieldCbtif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2ifcrFieldCbtif2Mask)
	}
}

const (
	RegisterMdma_c2ifcrFieldCltcif2Shift = 4
	RegisterMdma_c2ifcrFieldCltcif2Mask  = 0x10
)

// GetCltcif2 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c2ifcrType) GetCltcif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2ifcrFieldCltcif2Mask) != 0
}

// SetCltcif2 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c2ifcrType) SetCltcif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2ifcrFieldCltcif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2ifcrFieldCltcif2Mask)
	}
}

// registerMdma_c2esrType MDMA Channel x error status register
type registerMdma_c2esrType uint32

const (
	RegisterMdma_c2esrFieldTeaShift = 0
	RegisterMdma_c2esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c2esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2esrFieldTeaMask) >> RegisterMdma_c2esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c2esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c2esrFieldTeaShift))
}

const (
	RegisterMdma_c2esrFieldTedShift = 7
	RegisterMdma_c2esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c2esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c2esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2esrFieldTedMask)
	}
}

const (
	RegisterMdma_c2esrFieldTeldShift = 8
	RegisterMdma_c2esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c2esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c2esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c2esrFieldTemdShift = 9
	RegisterMdma_c2esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c2esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c2esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c2esrFieldAseShift = 10
	RegisterMdma_c2esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c2esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c2esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2esrFieldAseMask)
	}
}

const (
	RegisterMdma_c2esrFieldBseShift = 11
	RegisterMdma_c2esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c2esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c2esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2esrFieldBseMask)
	}
}

// registerMdma_c2crType This register is used to control the concerned channel.
type registerMdma_c2crType uint32

const (
	RegisterMdma_c2crFieldEnShift = 0
	RegisterMdma_c2crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c2crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c2crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2crFieldEnMask)
	}
}

const (
	RegisterMdma_c2crFieldTeieShift = 1
	RegisterMdma_c2crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c2crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c2crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2crFieldTeieMask)
	}
}

const (
	RegisterMdma_c2crFieldCtcieShift = 2
	RegisterMdma_c2crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c2crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c2crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c2crFieldBrtieShift = 3
	RegisterMdma_c2crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c2crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c2crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c2crFieldBtieShift = 4
	RegisterMdma_c2crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c2crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c2crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2crFieldBtieMask)
	}
}

const (
	RegisterMdma_c2crFieldTcieShift = 5
	RegisterMdma_c2crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c2crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c2crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2crFieldTcieMask)
	}
}

const (
	RegisterMdma_c2crFieldPlShift = 6
	RegisterMdma_c2crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c2crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2crFieldPlMask) >> RegisterMdma_c2crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c2crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2crFieldPlMask)|(uint32(value)<<RegisterMdma_c2crFieldPlShift))
}

const (
	RegisterMdma_c2crFieldBexShift = 12
	RegisterMdma_c2crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c2crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c2crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2crFieldBexMask)
	}
}

const (
	RegisterMdma_c2crFieldHexShift = 13
	RegisterMdma_c2crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c2crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c2crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2crFieldHexMask)
	}
}

const (
	RegisterMdma_c2crFieldWexShift = 14
	RegisterMdma_c2crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c2crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c2crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2crFieldWexMask)
	}
}

const (
	RegisterMdma_c2crFieldSwrqShift = 16
	RegisterMdma_c2crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c2crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2crFieldSwrqMask)
	}
}

// registerMdma_c2tcrType This register is used to configure the concerned channel.
type registerMdma_c2tcrType uint32

const (
	RegisterMdma_c2tcrFieldSincShift = 0
	RegisterMdma_c2tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c2tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldSincMask) >> RegisterMdma_c2tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c2tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c2tcrFieldSincShift))
}

const (
	RegisterMdma_c2tcrFieldDincShift = 2
	RegisterMdma_c2tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c2tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldDincMask) >> RegisterMdma_c2tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c2tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c2tcrFieldDincShift))
}

const (
	RegisterMdma_c2tcrFieldSsizeShift = 4
	RegisterMdma_c2tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c2tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldSsizeMask) >> RegisterMdma_c2tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c2tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c2tcrFieldSsizeShift))
}

const (
	RegisterMdma_c2tcrFieldDsizeShift = 6
	RegisterMdma_c2tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c2tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldDsizeMask) >> RegisterMdma_c2tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c2tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c2tcrFieldDsizeShift))
}

const (
	RegisterMdma_c2tcrFieldSincosShift = 8
	RegisterMdma_c2tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c2tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldSincosMask) >> RegisterMdma_c2tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c2tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c2tcrFieldSincosShift))
}

const (
	RegisterMdma_c2tcrFieldDincosShift = 10
	RegisterMdma_c2tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c2tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldDincosMask) >> RegisterMdma_c2tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c2tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c2tcrFieldDincosShift))
}

const (
	RegisterMdma_c2tcrFieldSburstShift = 12
	RegisterMdma_c2tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c2tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldSburstMask) >> RegisterMdma_c2tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c2tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c2tcrFieldSburstShift))
}

const (
	RegisterMdma_c2tcrFieldDburstShift = 15
	RegisterMdma_c2tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c2tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldDburstMask) >> RegisterMdma_c2tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c2tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c2tcrFieldDburstShift))
}

const (
	RegisterMdma_c2tcrFieldTlenShift = 18
	RegisterMdma_c2tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c2tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldTlenMask) >> RegisterMdma_c2tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c2tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c2tcrFieldTlenShift))
}

const (
	RegisterMdma_c2tcrFieldPkeShift = 25
	RegisterMdma_c2tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c2tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c2tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c2tcrFieldPamShift = 26
	RegisterMdma_c2tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c2tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldPamMask) >> RegisterMdma_c2tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c2tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c2tcrFieldPamShift))
}

const (
	RegisterMdma_c2tcrFieldTrgmShift = 28
	RegisterMdma_c2tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c2tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldTrgmMask) >> RegisterMdma_c2tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c2tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c2tcrFieldTrgmShift))
}

const (
	RegisterMdma_c2tcrFieldSwrmShift = 30
	RegisterMdma_c2tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c2tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c2tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c2tcrFieldBwmShift = 31
	RegisterMdma_c2tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c2tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c2tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tcrFieldBwmMask)
	}
}

// registerMdma_c2bndtrType MDMA Channel x block number of data register
type registerMdma_c2bndtrType uint32

const (
	RegisterMdma_c2bndtrFieldBndtShift = 0
	RegisterMdma_c2bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c2bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2bndtrFieldBndtMask) >> RegisterMdma_c2bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c2bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c2bndtrFieldBndtShift))
}

const (
	RegisterMdma_c2bndtrFieldBrsumShift = 18
	RegisterMdma_c2bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c2bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c2bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c2bndtrFieldBrdumShift = 19
	RegisterMdma_c2bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c2bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c2bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c2bndtrFieldBrcShift = 20
	RegisterMdma_c2bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c2bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2bndtrFieldBrcMask) >> RegisterMdma_c2bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c2bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c2bndtrFieldBrcShift))
}

// registerMdma_c2sarType MDMA channel x source address register
type registerMdma_c2sarType uint32

const (
	RegisterMdma_c2sarFieldSarShift = 0
	RegisterMdma_c2sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c2sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2sarFieldSarMask) >> RegisterMdma_c2sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c2sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2sarFieldSarMask)|(uint32(value)<<RegisterMdma_c2sarFieldSarShift))
}

// registerMdma_c2darType MDMA channel x destination address register
type registerMdma_c2darType uint32

const (
	RegisterMdma_c2darFieldDarShift = 0
	RegisterMdma_c2darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c2darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2darFieldDarMask) >> RegisterMdma_c2darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c2darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2darFieldDarMask)|(uint32(value)<<RegisterMdma_c2darFieldDarShift))
}

// registerMdma_c2brurType MDMA channel x Block Repeat address Update register
type registerMdma_c2brurType uint32

const (
	RegisterMdma_c2brurFieldSuvShift = 0
	RegisterMdma_c2brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c2brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2brurFieldSuvMask) >> RegisterMdma_c2brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c2brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c2brurFieldSuvShift))
}

const (
	RegisterMdma_c2brurFieldDuvShift = 16
	RegisterMdma_c2brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c2brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2brurFieldDuvMask) >> RegisterMdma_c2brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c2brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c2brurFieldDuvShift))
}

// registerMdma_c2larType MDMA channel x Link Address register
type registerMdma_c2larType uint32

const (
	RegisterMdma_c2larFieldLarShift = 0
	RegisterMdma_c2larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c2larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2larFieldLarMask) >> RegisterMdma_c2larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c2larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2larFieldLarMask)|(uint32(value)<<RegisterMdma_c2larFieldLarShift))
}

// registerMdma_c2tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c2tbrType uint32

const (
	RegisterMdma_c2tbrFieldTselShift = 0
	RegisterMdma_c2tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c2tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tbrFieldTselMask) >> RegisterMdma_c2tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c2tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c2tbrFieldTselShift))
}

const (
	RegisterMdma_c2tbrFieldSbusShift = 16
	RegisterMdma_c2tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c2tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c2tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c2tbrFieldDbusShift = 17
	RegisterMdma_c2tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c2tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c2tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c2tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2tbrFieldDbusMask)
	}
}

// registerMdma_c2marType MDMA channel x Mask address register
type registerMdma_c2marType uint32

const (
	RegisterMdma_c2marFieldMarShift = 0
	RegisterMdma_c2marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c2marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2marFieldMarMask) >> RegisterMdma_c2marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c2marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2marFieldMarMask)|(uint32(value)<<RegisterMdma_c2marFieldMarShift))
}

// registerMdma_c2mdrType MDMA channel x Mask Data register
type registerMdma_c2mdrType uint32

const (
	RegisterMdma_c2mdrFieldMdrShift = 0
	RegisterMdma_c2mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c2mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c2mdrFieldMdrMask) >> RegisterMdma_c2mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c2mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c2mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c2mdrFieldMdrShift))
}

// registerMdma_c3isrType MDMA channel x interrupt/status register
type registerMdma_c3isrType uint32

const (
	RegisterMdma_c3isrFieldTeif3Shift = 0
	RegisterMdma_c3isrFieldTeif3Mask  = 0x1
)

// GetTeif3 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c3isrType) GetTeif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3isrFieldTeif3Mask) != 0
}

// SetTeif3 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c3isrType) SetTeif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3isrFieldTeif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3isrFieldTeif3Mask)
	}
}

const (
	RegisterMdma_c3isrFieldCtcif3Shift = 1
	RegisterMdma_c3isrFieldCtcif3Mask  = 0x2
)

// GetCtcif3 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c3isrType) GetCtcif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3isrFieldCtcif3Mask) != 0
}

// SetCtcif3 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c3isrType) SetCtcif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3isrFieldCtcif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3isrFieldCtcif3Mask)
	}
}

const (
	RegisterMdma_c3isrFieldBrtif3Shift = 2
	RegisterMdma_c3isrFieldBrtif3Mask  = 0x4
)

// GetBrtif3 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c3isrType) GetBrtif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3isrFieldBrtif3Mask) != 0
}

// SetBrtif3 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c3isrType) SetBrtif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3isrFieldBrtif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3isrFieldBrtif3Mask)
	}
}

const (
	RegisterMdma_c3isrFieldBtif3Shift = 3
	RegisterMdma_c3isrFieldBtif3Mask  = 0x8
)

// GetBtif3 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c3isrType) GetBtif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3isrFieldBtif3Mask) != 0
}

// SetBtif3 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c3isrType) SetBtif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3isrFieldBtif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3isrFieldBtif3Mask)
	}
}

const (
	RegisterMdma_c3isrFieldTcif3Shift = 4
	RegisterMdma_c3isrFieldTcif3Mask  = 0x10
)

// GetTcif3 channel x buffer transfer complete
func (r *registerMdma_c3isrType) GetTcif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3isrFieldTcif3Mask) != 0
}

// SetTcif3 channel x buffer transfer complete
func (r *registerMdma_c3isrType) SetTcif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3isrFieldTcif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3isrFieldTcif3Mask)
	}
}

const (
	RegisterMdma_c3isrFieldCrqa3Shift = 16
	RegisterMdma_c3isrFieldCrqa3Mask  = 0x10000
)

// GetCrqa3 channel x request active flag
func (r *registerMdma_c3isrType) GetCrqa3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3isrFieldCrqa3Mask) != 0
}

// SetCrqa3 channel x request active flag
func (r *registerMdma_c3isrType) SetCrqa3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3isrFieldCrqa3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3isrFieldCrqa3Mask)
	}
}

// registerMdma_c3ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c3ifcrType uint32

const (
	RegisterMdma_c3ifcrFieldCteif3Shift = 0
	RegisterMdma_c3ifcrFieldCteif3Mask  = 0x1
)

// GetCteif3 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c3ifcrType) GetCteif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3ifcrFieldCteif3Mask) != 0
}

// SetCteif3 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c3ifcrType) SetCteif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3ifcrFieldCteif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3ifcrFieldCteif3Mask)
	}
}

const (
	RegisterMdma_c3ifcrFieldCctcif3Shift = 1
	RegisterMdma_c3ifcrFieldCctcif3Mask  = 0x2
)

// GetCctcif3 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c3ifcrType) GetCctcif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3ifcrFieldCctcif3Mask) != 0
}

// SetCctcif3 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c3ifcrType) SetCctcif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3ifcrFieldCctcif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3ifcrFieldCctcif3Mask)
	}
}

const (
	RegisterMdma_c3ifcrFieldCbrtif3Shift = 2
	RegisterMdma_c3ifcrFieldCbrtif3Mask  = 0x4
)

// GetCbrtif3 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c3ifcrType) GetCbrtif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3ifcrFieldCbrtif3Mask) != 0
}

// SetCbrtif3 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c3ifcrType) SetCbrtif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3ifcrFieldCbrtif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3ifcrFieldCbrtif3Mask)
	}
}

const (
	RegisterMdma_c3ifcrFieldCbtif3Shift = 3
	RegisterMdma_c3ifcrFieldCbtif3Mask  = 0x8
)

// GetCbtif3 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c3ifcrType) GetCbtif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3ifcrFieldCbtif3Mask) != 0
}

// SetCbtif3 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c3ifcrType) SetCbtif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3ifcrFieldCbtif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3ifcrFieldCbtif3Mask)
	}
}

const (
	RegisterMdma_c3ifcrFieldCltcif3Shift = 4
	RegisterMdma_c3ifcrFieldCltcif3Mask  = 0x10
)

// GetCltcif3 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c3ifcrType) GetCltcif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3ifcrFieldCltcif3Mask) != 0
}

// SetCltcif3 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c3ifcrType) SetCltcif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3ifcrFieldCltcif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3ifcrFieldCltcif3Mask)
	}
}

// registerMdma_c3esrType MDMA Channel x error status register
type registerMdma_c3esrType uint32

const (
	RegisterMdma_c3esrFieldTeaShift = 0
	RegisterMdma_c3esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c3esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3esrFieldTeaMask) >> RegisterMdma_c3esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c3esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c3esrFieldTeaShift))
}

const (
	RegisterMdma_c3esrFieldTedShift = 7
	RegisterMdma_c3esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c3esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c3esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3esrFieldTedMask)
	}
}

const (
	RegisterMdma_c3esrFieldTeldShift = 8
	RegisterMdma_c3esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c3esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c3esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c3esrFieldTemdShift = 9
	RegisterMdma_c3esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c3esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c3esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c3esrFieldAseShift = 10
	RegisterMdma_c3esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c3esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c3esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3esrFieldAseMask)
	}
}

const (
	RegisterMdma_c3esrFieldBseShift = 11
	RegisterMdma_c3esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c3esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c3esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3esrFieldBseMask)
	}
}

// registerMdma_c3crType This register is used to control the concerned channel.
type registerMdma_c3crType uint32

const (
	RegisterMdma_c3crFieldEnShift = 0
	RegisterMdma_c3crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c3crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c3crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3crFieldEnMask)
	}
}

const (
	RegisterMdma_c3crFieldTeieShift = 1
	RegisterMdma_c3crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c3crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c3crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3crFieldTeieMask)
	}
}

const (
	RegisterMdma_c3crFieldCtcieShift = 2
	RegisterMdma_c3crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c3crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c3crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c3crFieldBrtieShift = 3
	RegisterMdma_c3crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c3crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c3crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c3crFieldBtieShift = 4
	RegisterMdma_c3crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c3crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c3crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3crFieldBtieMask)
	}
}

const (
	RegisterMdma_c3crFieldTcieShift = 5
	RegisterMdma_c3crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c3crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c3crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3crFieldTcieMask)
	}
}

const (
	RegisterMdma_c3crFieldPlShift = 6
	RegisterMdma_c3crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c3crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3crFieldPlMask) >> RegisterMdma_c3crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c3crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3crFieldPlMask)|(uint32(value)<<RegisterMdma_c3crFieldPlShift))
}

const (
	RegisterMdma_c3crFieldBexShift = 12
	RegisterMdma_c3crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c3crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c3crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3crFieldBexMask)
	}
}

const (
	RegisterMdma_c3crFieldHexShift = 13
	RegisterMdma_c3crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c3crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c3crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3crFieldHexMask)
	}
}

const (
	RegisterMdma_c3crFieldWexShift = 14
	RegisterMdma_c3crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c3crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c3crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3crFieldWexMask)
	}
}

const (
	RegisterMdma_c3crFieldSwrqShift = 16
	RegisterMdma_c3crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c3crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3crFieldSwrqMask)
	}
}

// registerMdma_c3tcrType This register is used to configure the concerned channel.
type registerMdma_c3tcrType uint32

const (
	RegisterMdma_c3tcrFieldSincShift = 0
	RegisterMdma_c3tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c3tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldSincMask) >> RegisterMdma_c3tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c3tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c3tcrFieldSincShift))
}

const (
	RegisterMdma_c3tcrFieldDincShift = 2
	RegisterMdma_c3tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c3tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldDincMask) >> RegisterMdma_c3tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c3tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c3tcrFieldDincShift))
}

const (
	RegisterMdma_c3tcrFieldSsizeShift = 4
	RegisterMdma_c3tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c3tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldSsizeMask) >> RegisterMdma_c3tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c3tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c3tcrFieldSsizeShift))
}

const (
	RegisterMdma_c3tcrFieldDsizeShift = 6
	RegisterMdma_c3tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c3tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldDsizeMask) >> RegisterMdma_c3tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c3tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c3tcrFieldDsizeShift))
}

const (
	RegisterMdma_c3tcrFieldSincosShift = 8
	RegisterMdma_c3tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c3tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldSincosMask) >> RegisterMdma_c3tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c3tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c3tcrFieldSincosShift))
}

const (
	RegisterMdma_c3tcrFieldDincosShift = 10
	RegisterMdma_c3tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c3tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldDincosMask) >> RegisterMdma_c3tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c3tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c3tcrFieldDincosShift))
}

const (
	RegisterMdma_c3tcrFieldSburstShift = 12
	RegisterMdma_c3tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c3tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldSburstMask) >> RegisterMdma_c3tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c3tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c3tcrFieldSburstShift))
}

const (
	RegisterMdma_c3tcrFieldDburstShift = 15
	RegisterMdma_c3tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c3tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldDburstMask) >> RegisterMdma_c3tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c3tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c3tcrFieldDburstShift))
}

const (
	RegisterMdma_c3tcrFieldTlenShift = 18
	RegisterMdma_c3tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c3tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldTlenMask) >> RegisterMdma_c3tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c3tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c3tcrFieldTlenShift))
}

const (
	RegisterMdma_c3tcrFieldPkeShift = 25
	RegisterMdma_c3tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c3tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c3tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c3tcrFieldPamShift = 26
	RegisterMdma_c3tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c3tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldPamMask) >> RegisterMdma_c3tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c3tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c3tcrFieldPamShift))
}

const (
	RegisterMdma_c3tcrFieldTrgmShift = 28
	RegisterMdma_c3tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c3tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldTrgmMask) >> RegisterMdma_c3tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c3tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c3tcrFieldTrgmShift))
}

const (
	RegisterMdma_c3tcrFieldSwrmShift = 30
	RegisterMdma_c3tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c3tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c3tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c3tcrFieldBwmShift = 31
	RegisterMdma_c3tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c3tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c3tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tcrFieldBwmMask)
	}
}

// registerMdma_c3bndtrType MDMA Channel x block number of data register
type registerMdma_c3bndtrType uint32

const (
	RegisterMdma_c3bndtrFieldBndtShift = 0
	RegisterMdma_c3bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c3bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3bndtrFieldBndtMask) >> RegisterMdma_c3bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c3bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c3bndtrFieldBndtShift))
}

const (
	RegisterMdma_c3bndtrFieldBrsumShift = 18
	RegisterMdma_c3bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c3bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c3bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c3bndtrFieldBrdumShift = 19
	RegisterMdma_c3bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c3bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c3bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c3bndtrFieldBrcShift = 20
	RegisterMdma_c3bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c3bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3bndtrFieldBrcMask) >> RegisterMdma_c3bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c3bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c3bndtrFieldBrcShift))
}

// registerMdma_c3sarType MDMA channel x source address register
type registerMdma_c3sarType uint32

const (
	RegisterMdma_c3sarFieldSarShift = 0
	RegisterMdma_c3sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c3sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3sarFieldSarMask) >> RegisterMdma_c3sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c3sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3sarFieldSarMask)|(uint32(value)<<RegisterMdma_c3sarFieldSarShift))
}

// registerMdma_c3darType MDMA channel x destination address register
type registerMdma_c3darType uint32

const (
	RegisterMdma_c3darFieldDarShift = 0
	RegisterMdma_c3darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c3darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3darFieldDarMask) >> RegisterMdma_c3darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c3darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3darFieldDarMask)|(uint32(value)<<RegisterMdma_c3darFieldDarShift))
}

// registerMdma_c3brurType MDMA channel x Block Repeat address Update register
type registerMdma_c3brurType uint32

const (
	RegisterMdma_c3brurFieldSuvShift = 0
	RegisterMdma_c3brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c3brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3brurFieldSuvMask) >> RegisterMdma_c3brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c3brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c3brurFieldSuvShift))
}

const (
	RegisterMdma_c3brurFieldDuvShift = 16
	RegisterMdma_c3brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c3brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3brurFieldDuvMask) >> RegisterMdma_c3brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c3brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c3brurFieldDuvShift))
}

// registerMdma_c3larType MDMA channel x Link Address register
type registerMdma_c3larType uint32

const (
	RegisterMdma_c3larFieldLarShift = 0
	RegisterMdma_c3larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c3larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3larFieldLarMask) >> RegisterMdma_c3larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c3larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3larFieldLarMask)|(uint32(value)<<RegisterMdma_c3larFieldLarShift))
}

// registerMdma_c3tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c3tbrType uint32

const (
	RegisterMdma_c3tbrFieldTselShift = 0
	RegisterMdma_c3tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c3tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tbrFieldTselMask) >> RegisterMdma_c3tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c3tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c3tbrFieldTselShift))
}

const (
	RegisterMdma_c3tbrFieldSbusShift = 16
	RegisterMdma_c3tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c3tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c3tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c3tbrFieldDbusShift = 17
	RegisterMdma_c3tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c3tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c3tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c3tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3tbrFieldDbusMask)
	}
}

// registerMdma_c3marType MDMA channel x Mask address register
type registerMdma_c3marType uint32

const (
	RegisterMdma_c3marFieldMarShift = 0
	RegisterMdma_c3marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c3marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3marFieldMarMask) >> RegisterMdma_c3marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c3marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3marFieldMarMask)|(uint32(value)<<RegisterMdma_c3marFieldMarShift))
}

// registerMdma_c3mdrType MDMA channel x Mask Data register
type registerMdma_c3mdrType uint32

const (
	RegisterMdma_c3mdrFieldMdrShift = 0
	RegisterMdma_c3mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c3mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c3mdrFieldMdrMask) >> RegisterMdma_c3mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c3mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c3mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c3mdrFieldMdrShift))
}

// registerMdma_c4isrType MDMA channel x interrupt/status register
type registerMdma_c4isrType uint32

const (
	RegisterMdma_c4isrFieldTeif4Shift = 0
	RegisterMdma_c4isrFieldTeif4Mask  = 0x1
)

// GetTeif4 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c4isrType) GetTeif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4isrFieldTeif4Mask) != 0
}

// SetTeif4 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c4isrType) SetTeif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4isrFieldTeif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4isrFieldTeif4Mask)
	}
}

const (
	RegisterMdma_c4isrFieldCtcif4Shift = 1
	RegisterMdma_c4isrFieldCtcif4Mask  = 0x2
)

// GetCtcif4 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c4isrType) GetCtcif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4isrFieldCtcif4Mask) != 0
}

// SetCtcif4 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c4isrType) SetCtcif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4isrFieldCtcif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4isrFieldCtcif4Mask)
	}
}

const (
	RegisterMdma_c4isrFieldBrtif4Shift = 2
	RegisterMdma_c4isrFieldBrtif4Mask  = 0x4
)

// GetBrtif4 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c4isrType) GetBrtif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4isrFieldBrtif4Mask) != 0
}

// SetBrtif4 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c4isrType) SetBrtif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4isrFieldBrtif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4isrFieldBrtif4Mask)
	}
}

const (
	RegisterMdma_c4isrFieldBtif4Shift = 3
	RegisterMdma_c4isrFieldBtif4Mask  = 0x8
)

// GetBtif4 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c4isrType) GetBtif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4isrFieldBtif4Mask) != 0
}

// SetBtif4 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c4isrType) SetBtif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4isrFieldBtif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4isrFieldBtif4Mask)
	}
}

const (
	RegisterMdma_c4isrFieldTcif4Shift = 4
	RegisterMdma_c4isrFieldTcif4Mask  = 0x10
)

// GetTcif4 channel x buffer transfer complete
func (r *registerMdma_c4isrType) GetTcif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4isrFieldTcif4Mask) != 0
}

// SetTcif4 channel x buffer transfer complete
func (r *registerMdma_c4isrType) SetTcif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4isrFieldTcif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4isrFieldTcif4Mask)
	}
}

const (
	RegisterMdma_c4isrFieldCrqa4Shift = 16
	RegisterMdma_c4isrFieldCrqa4Mask  = 0x10000
)

// GetCrqa4 channel x request active flag
func (r *registerMdma_c4isrType) GetCrqa4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4isrFieldCrqa4Mask) != 0
}

// SetCrqa4 channel x request active flag
func (r *registerMdma_c4isrType) SetCrqa4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4isrFieldCrqa4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4isrFieldCrqa4Mask)
	}
}

// registerMdma_c4ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c4ifcrType uint32

const (
	RegisterMdma_c4ifcrFieldCteif4Shift = 0
	RegisterMdma_c4ifcrFieldCteif4Mask  = 0x1
)

// GetCteif4 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c4ifcrType) GetCteif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4ifcrFieldCteif4Mask) != 0
}

// SetCteif4 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c4ifcrType) SetCteif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4ifcrFieldCteif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4ifcrFieldCteif4Mask)
	}
}

const (
	RegisterMdma_c4ifcrFieldCctcif4Shift = 1
	RegisterMdma_c4ifcrFieldCctcif4Mask  = 0x2
)

// GetCctcif4 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c4ifcrType) GetCctcif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4ifcrFieldCctcif4Mask) != 0
}

// SetCctcif4 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c4ifcrType) SetCctcif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4ifcrFieldCctcif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4ifcrFieldCctcif4Mask)
	}
}

const (
	RegisterMdma_c4ifcrFieldCbrtif4Shift = 2
	RegisterMdma_c4ifcrFieldCbrtif4Mask  = 0x4
)

// GetCbrtif4 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c4ifcrType) GetCbrtif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4ifcrFieldCbrtif4Mask) != 0
}

// SetCbrtif4 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c4ifcrType) SetCbrtif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4ifcrFieldCbrtif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4ifcrFieldCbrtif4Mask)
	}
}

const (
	RegisterMdma_c4ifcrFieldCbtif4Shift = 3
	RegisterMdma_c4ifcrFieldCbtif4Mask  = 0x8
)

// GetCbtif4 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c4ifcrType) GetCbtif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4ifcrFieldCbtif4Mask) != 0
}

// SetCbtif4 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c4ifcrType) SetCbtif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4ifcrFieldCbtif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4ifcrFieldCbtif4Mask)
	}
}

const (
	RegisterMdma_c4ifcrFieldCltcif4Shift = 4
	RegisterMdma_c4ifcrFieldCltcif4Mask  = 0x10
)

// GetCltcif4 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c4ifcrType) GetCltcif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4ifcrFieldCltcif4Mask) != 0
}

// SetCltcif4 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c4ifcrType) SetCltcif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4ifcrFieldCltcif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4ifcrFieldCltcif4Mask)
	}
}

// registerMdma_c4esrType MDMA Channel x error status register
type registerMdma_c4esrType uint32

const (
	RegisterMdma_c4esrFieldTeaShift = 0
	RegisterMdma_c4esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c4esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4esrFieldTeaMask) >> RegisterMdma_c4esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c4esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c4esrFieldTeaShift))
}

const (
	RegisterMdma_c4esrFieldTedShift = 7
	RegisterMdma_c4esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c4esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c4esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4esrFieldTedMask)
	}
}

const (
	RegisterMdma_c4esrFieldTeldShift = 8
	RegisterMdma_c4esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c4esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c4esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c4esrFieldTemdShift = 9
	RegisterMdma_c4esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c4esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c4esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c4esrFieldAseShift = 10
	RegisterMdma_c4esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c4esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c4esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4esrFieldAseMask)
	}
}

const (
	RegisterMdma_c4esrFieldBseShift = 11
	RegisterMdma_c4esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c4esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c4esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4esrFieldBseMask)
	}
}

// registerMdma_c4crType This register is used to control the concerned channel.
type registerMdma_c4crType uint32

const (
	RegisterMdma_c4crFieldEnShift = 0
	RegisterMdma_c4crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c4crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c4crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4crFieldEnMask)
	}
}

const (
	RegisterMdma_c4crFieldTeieShift = 1
	RegisterMdma_c4crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c4crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c4crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4crFieldTeieMask)
	}
}

const (
	RegisterMdma_c4crFieldCtcieShift = 2
	RegisterMdma_c4crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c4crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c4crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c4crFieldBrtieShift = 3
	RegisterMdma_c4crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c4crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c4crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c4crFieldBtieShift = 4
	RegisterMdma_c4crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c4crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c4crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4crFieldBtieMask)
	}
}

const (
	RegisterMdma_c4crFieldTcieShift = 5
	RegisterMdma_c4crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c4crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c4crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4crFieldTcieMask)
	}
}

const (
	RegisterMdma_c4crFieldPlShift = 6
	RegisterMdma_c4crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c4crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4crFieldPlMask) >> RegisterMdma_c4crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c4crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4crFieldPlMask)|(uint32(value)<<RegisterMdma_c4crFieldPlShift))
}

const (
	RegisterMdma_c4crFieldBexShift = 12
	RegisterMdma_c4crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c4crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c4crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4crFieldBexMask)
	}
}

const (
	RegisterMdma_c4crFieldHexShift = 13
	RegisterMdma_c4crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c4crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c4crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4crFieldHexMask)
	}
}

const (
	RegisterMdma_c4crFieldWexShift = 14
	RegisterMdma_c4crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c4crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c4crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4crFieldWexMask)
	}
}

const (
	RegisterMdma_c4crFieldSwrqShift = 16
	RegisterMdma_c4crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c4crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4crFieldSwrqMask)
	}
}

// registerMdma_c4tcrType This register is used to configure the concerned channel.
type registerMdma_c4tcrType uint32

const (
	RegisterMdma_c4tcrFieldSincShift = 0
	RegisterMdma_c4tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c4tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldSincMask) >> RegisterMdma_c4tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c4tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c4tcrFieldSincShift))
}

const (
	RegisterMdma_c4tcrFieldDincShift = 2
	RegisterMdma_c4tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c4tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldDincMask) >> RegisterMdma_c4tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c4tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c4tcrFieldDincShift))
}

const (
	RegisterMdma_c4tcrFieldSsizeShift = 4
	RegisterMdma_c4tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c4tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldSsizeMask) >> RegisterMdma_c4tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c4tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c4tcrFieldSsizeShift))
}

const (
	RegisterMdma_c4tcrFieldDsizeShift = 6
	RegisterMdma_c4tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c4tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldDsizeMask) >> RegisterMdma_c4tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c4tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c4tcrFieldDsizeShift))
}

const (
	RegisterMdma_c4tcrFieldSincosShift = 8
	RegisterMdma_c4tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c4tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldSincosMask) >> RegisterMdma_c4tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c4tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c4tcrFieldSincosShift))
}

const (
	RegisterMdma_c4tcrFieldDincosShift = 10
	RegisterMdma_c4tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c4tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldDincosMask) >> RegisterMdma_c4tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c4tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c4tcrFieldDincosShift))
}

const (
	RegisterMdma_c4tcrFieldSburstShift = 12
	RegisterMdma_c4tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c4tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldSburstMask) >> RegisterMdma_c4tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c4tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c4tcrFieldSburstShift))
}

const (
	RegisterMdma_c4tcrFieldDburstShift = 15
	RegisterMdma_c4tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c4tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldDburstMask) >> RegisterMdma_c4tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c4tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c4tcrFieldDburstShift))
}

const (
	RegisterMdma_c4tcrFieldTlenShift = 18
	RegisterMdma_c4tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c4tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldTlenMask) >> RegisterMdma_c4tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c4tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c4tcrFieldTlenShift))
}

const (
	RegisterMdma_c4tcrFieldPkeShift = 25
	RegisterMdma_c4tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c4tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c4tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c4tcrFieldPamShift = 26
	RegisterMdma_c4tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c4tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldPamMask) >> RegisterMdma_c4tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c4tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c4tcrFieldPamShift))
}

const (
	RegisterMdma_c4tcrFieldTrgmShift = 28
	RegisterMdma_c4tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c4tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldTrgmMask) >> RegisterMdma_c4tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c4tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c4tcrFieldTrgmShift))
}

const (
	RegisterMdma_c4tcrFieldSwrmShift = 30
	RegisterMdma_c4tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c4tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c4tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c4tcrFieldBwmShift = 31
	RegisterMdma_c4tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c4tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c4tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tcrFieldBwmMask)
	}
}

// registerMdma_c4bndtrType MDMA Channel x block number of data register
type registerMdma_c4bndtrType uint32

const (
	RegisterMdma_c4bndtrFieldBndtShift = 0
	RegisterMdma_c4bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c4bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4bndtrFieldBndtMask) >> RegisterMdma_c4bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c4bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c4bndtrFieldBndtShift))
}

const (
	RegisterMdma_c4bndtrFieldBrsumShift = 18
	RegisterMdma_c4bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c4bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c4bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c4bndtrFieldBrdumShift = 19
	RegisterMdma_c4bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c4bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c4bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c4bndtrFieldBrcShift = 20
	RegisterMdma_c4bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c4bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4bndtrFieldBrcMask) >> RegisterMdma_c4bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c4bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c4bndtrFieldBrcShift))
}

// registerMdma_c4sarType MDMA channel x source address register
type registerMdma_c4sarType uint32

const (
	RegisterMdma_c4sarFieldSarShift = 0
	RegisterMdma_c4sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c4sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4sarFieldSarMask) >> RegisterMdma_c4sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c4sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4sarFieldSarMask)|(uint32(value)<<RegisterMdma_c4sarFieldSarShift))
}

// registerMdma_c4darType MDMA channel x destination address register
type registerMdma_c4darType uint32

const (
	RegisterMdma_c4darFieldDarShift = 0
	RegisterMdma_c4darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c4darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4darFieldDarMask) >> RegisterMdma_c4darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c4darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4darFieldDarMask)|(uint32(value)<<RegisterMdma_c4darFieldDarShift))
}

// registerMdma_c4brurType MDMA channel x Block Repeat address Update register
type registerMdma_c4brurType uint32

const (
	RegisterMdma_c4brurFieldSuvShift = 0
	RegisterMdma_c4brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c4brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4brurFieldSuvMask) >> RegisterMdma_c4brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c4brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c4brurFieldSuvShift))
}

const (
	RegisterMdma_c4brurFieldDuvShift = 16
	RegisterMdma_c4brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c4brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4brurFieldDuvMask) >> RegisterMdma_c4brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c4brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c4brurFieldDuvShift))
}

// registerMdma_c4larType MDMA channel x Link Address register
type registerMdma_c4larType uint32

const (
	RegisterMdma_c4larFieldLarShift = 0
	RegisterMdma_c4larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c4larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4larFieldLarMask) >> RegisterMdma_c4larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c4larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4larFieldLarMask)|(uint32(value)<<RegisterMdma_c4larFieldLarShift))
}

// registerMdma_c4tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c4tbrType uint32

const (
	RegisterMdma_c4tbrFieldTselShift = 0
	RegisterMdma_c4tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c4tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tbrFieldTselMask) >> RegisterMdma_c4tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c4tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c4tbrFieldTselShift))
}

const (
	RegisterMdma_c4tbrFieldSbusShift = 16
	RegisterMdma_c4tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c4tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c4tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c4tbrFieldDbusShift = 17
	RegisterMdma_c4tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c4tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c4tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c4tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4tbrFieldDbusMask)
	}
}

// registerMdma_c4marType MDMA channel x Mask address register
type registerMdma_c4marType uint32

const (
	RegisterMdma_c4marFieldMarShift = 0
	RegisterMdma_c4marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c4marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4marFieldMarMask) >> RegisterMdma_c4marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c4marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4marFieldMarMask)|(uint32(value)<<RegisterMdma_c4marFieldMarShift))
}

// registerMdma_c4mdrType MDMA channel x Mask Data register
type registerMdma_c4mdrType uint32

const (
	RegisterMdma_c4mdrFieldMdrShift = 0
	RegisterMdma_c4mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c4mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c4mdrFieldMdrMask) >> RegisterMdma_c4mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c4mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c4mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c4mdrFieldMdrShift))
}

// registerMdma_c5isrType MDMA channel x interrupt/status register
type registerMdma_c5isrType uint32

const (
	RegisterMdma_c5isrFieldTeif5Shift = 0
	RegisterMdma_c5isrFieldTeif5Mask  = 0x1
)

// GetTeif5 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c5isrType) GetTeif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5isrFieldTeif5Mask) != 0
}

// SetTeif5 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c5isrType) SetTeif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5isrFieldTeif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5isrFieldTeif5Mask)
	}
}

const (
	RegisterMdma_c5isrFieldCtcif5Shift = 1
	RegisterMdma_c5isrFieldCtcif5Mask  = 0x2
)

// GetCtcif5 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c5isrType) GetCtcif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5isrFieldCtcif5Mask) != 0
}

// SetCtcif5 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c5isrType) SetCtcif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5isrFieldCtcif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5isrFieldCtcif5Mask)
	}
}

const (
	RegisterMdma_c5isrFieldBrtif5Shift = 2
	RegisterMdma_c5isrFieldBrtif5Mask  = 0x4
)

// GetBrtif5 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c5isrType) GetBrtif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5isrFieldBrtif5Mask) != 0
}

// SetBrtif5 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c5isrType) SetBrtif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5isrFieldBrtif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5isrFieldBrtif5Mask)
	}
}

const (
	RegisterMdma_c5isrFieldBtif5Shift = 3
	RegisterMdma_c5isrFieldBtif5Mask  = 0x8
)

// GetBtif5 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c5isrType) GetBtif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5isrFieldBtif5Mask) != 0
}

// SetBtif5 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c5isrType) SetBtif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5isrFieldBtif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5isrFieldBtif5Mask)
	}
}

const (
	RegisterMdma_c5isrFieldTcif5Shift = 4
	RegisterMdma_c5isrFieldTcif5Mask  = 0x10
)

// GetTcif5 channel x buffer transfer complete
func (r *registerMdma_c5isrType) GetTcif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5isrFieldTcif5Mask) != 0
}

// SetTcif5 channel x buffer transfer complete
func (r *registerMdma_c5isrType) SetTcif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5isrFieldTcif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5isrFieldTcif5Mask)
	}
}

const (
	RegisterMdma_c5isrFieldCrqa5Shift = 16
	RegisterMdma_c5isrFieldCrqa5Mask  = 0x10000
)

// GetCrqa5 channel x request active flag
func (r *registerMdma_c5isrType) GetCrqa5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5isrFieldCrqa5Mask) != 0
}

// SetCrqa5 channel x request active flag
func (r *registerMdma_c5isrType) SetCrqa5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5isrFieldCrqa5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5isrFieldCrqa5Mask)
	}
}

// registerMdma_c5ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c5ifcrType uint32

const (
	RegisterMdma_c5ifcrFieldCteif5Shift = 0
	RegisterMdma_c5ifcrFieldCteif5Mask  = 0x1
)

// GetCteif5 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c5ifcrType) GetCteif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5ifcrFieldCteif5Mask) != 0
}

// SetCteif5 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c5ifcrType) SetCteif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5ifcrFieldCteif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5ifcrFieldCteif5Mask)
	}
}

const (
	RegisterMdma_c5ifcrFieldCctcif5Shift = 1
	RegisterMdma_c5ifcrFieldCctcif5Mask  = 0x2
)

// GetCctcif5 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c5ifcrType) GetCctcif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5ifcrFieldCctcif5Mask) != 0
}

// SetCctcif5 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c5ifcrType) SetCctcif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5ifcrFieldCctcif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5ifcrFieldCctcif5Mask)
	}
}

const (
	RegisterMdma_c5ifcrFieldCbrtif5Shift = 2
	RegisterMdma_c5ifcrFieldCbrtif5Mask  = 0x4
)

// GetCbrtif5 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c5ifcrType) GetCbrtif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5ifcrFieldCbrtif5Mask) != 0
}

// SetCbrtif5 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c5ifcrType) SetCbrtif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5ifcrFieldCbrtif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5ifcrFieldCbrtif5Mask)
	}
}

const (
	RegisterMdma_c5ifcrFieldCbtif5Shift = 3
	RegisterMdma_c5ifcrFieldCbtif5Mask  = 0x8
)

// GetCbtif5 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c5ifcrType) GetCbtif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5ifcrFieldCbtif5Mask) != 0
}

// SetCbtif5 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c5ifcrType) SetCbtif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5ifcrFieldCbtif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5ifcrFieldCbtif5Mask)
	}
}

const (
	RegisterMdma_c5ifcrFieldCltcif5Shift = 4
	RegisterMdma_c5ifcrFieldCltcif5Mask  = 0x10
)

// GetCltcif5 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c5ifcrType) GetCltcif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5ifcrFieldCltcif5Mask) != 0
}

// SetCltcif5 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c5ifcrType) SetCltcif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5ifcrFieldCltcif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5ifcrFieldCltcif5Mask)
	}
}

// registerMdma_c5esrType MDMA Channel x error status register
type registerMdma_c5esrType uint32

const (
	RegisterMdma_c5esrFieldTeaShift = 0
	RegisterMdma_c5esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c5esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5esrFieldTeaMask) >> RegisterMdma_c5esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c5esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c5esrFieldTeaShift))
}

const (
	RegisterMdma_c5esrFieldTedShift = 7
	RegisterMdma_c5esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c5esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c5esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5esrFieldTedMask)
	}
}

const (
	RegisterMdma_c5esrFieldTeldShift = 8
	RegisterMdma_c5esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c5esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c5esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c5esrFieldTemdShift = 9
	RegisterMdma_c5esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c5esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c5esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c5esrFieldAseShift = 10
	RegisterMdma_c5esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c5esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c5esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5esrFieldAseMask)
	}
}

const (
	RegisterMdma_c5esrFieldBseShift = 11
	RegisterMdma_c5esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c5esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c5esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5esrFieldBseMask)
	}
}

// registerMdma_c5crType This register is used to control the concerned channel.
type registerMdma_c5crType uint32

const (
	RegisterMdma_c5crFieldEnShift = 0
	RegisterMdma_c5crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c5crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c5crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5crFieldEnMask)
	}
}

const (
	RegisterMdma_c5crFieldTeieShift = 1
	RegisterMdma_c5crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c5crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c5crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5crFieldTeieMask)
	}
}

const (
	RegisterMdma_c5crFieldCtcieShift = 2
	RegisterMdma_c5crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c5crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c5crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c5crFieldBrtieShift = 3
	RegisterMdma_c5crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c5crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c5crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c5crFieldBtieShift = 4
	RegisterMdma_c5crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c5crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c5crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5crFieldBtieMask)
	}
}

const (
	RegisterMdma_c5crFieldTcieShift = 5
	RegisterMdma_c5crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c5crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c5crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5crFieldTcieMask)
	}
}

const (
	RegisterMdma_c5crFieldPlShift = 6
	RegisterMdma_c5crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c5crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5crFieldPlMask) >> RegisterMdma_c5crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c5crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5crFieldPlMask)|(uint32(value)<<RegisterMdma_c5crFieldPlShift))
}

const (
	RegisterMdma_c5crFieldBexShift = 12
	RegisterMdma_c5crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c5crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c5crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5crFieldBexMask)
	}
}

const (
	RegisterMdma_c5crFieldHexShift = 13
	RegisterMdma_c5crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c5crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c5crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5crFieldHexMask)
	}
}

const (
	RegisterMdma_c5crFieldWexShift = 14
	RegisterMdma_c5crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c5crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c5crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5crFieldWexMask)
	}
}

const (
	RegisterMdma_c5crFieldSwrqShift = 16
	RegisterMdma_c5crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c5crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5crFieldSwrqMask)
	}
}

// registerMdma_c5tcrType This register is used to configure the concerned channel.
type registerMdma_c5tcrType uint32

const (
	RegisterMdma_c5tcrFieldSincShift = 0
	RegisterMdma_c5tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c5tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldSincMask) >> RegisterMdma_c5tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c5tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c5tcrFieldSincShift))
}

const (
	RegisterMdma_c5tcrFieldDincShift = 2
	RegisterMdma_c5tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c5tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldDincMask) >> RegisterMdma_c5tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c5tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c5tcrFieldDincShift))
}

const (
	RegisterMdma_c5tcrFieldSsizeShift = 4
	RegisterMdma_c5tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c5tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldSsizeMask) >> RegisterMdma_c5tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c5tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c5tcrFieldSsizeShift))
}

const (
	RegisterMdma_c5tcrFieldDsizeShift = 6
	RegisterMdma_c5tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c5tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldDsizeMask) >> RegisterMdma_c5tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c5tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c5tcrFieldDsizeShift))
}

const (
	RegisterMdma_c5tcrFieldSincosShift = 8
	RegisterMdma_c5tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c5tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldSincosMask) >> RegisterMdma_c5tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c5tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c5tcrFieldSincosShift))
}

const (
	RegisterMdma_c5tcrFieldDincosShift = 10
	RegisterMdma_c5tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c5tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldDincosMask) >> RegisterMdma_c5tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c5tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c5tcrFieldDincosShift))
}

const (
	RegisterMdma_c5tcrFieldSburstShift = 12
	RegisterMdma_c5tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c5tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldSburstMask) >> RegisterMdma_c5tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c5tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c5tcrFieldSburstShift))
}

const (
	RegisterMdma_c5tcrFieldDburstShift = 15
	RegisterMdma_c5tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c5tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldDburstMask) >> RegisterMdma_c5tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c5tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c5tcrFieldDburstShift))
}

const (
	RegisterMdma_c5tcrFieldTlenShift = 18
	RegisterMdma_c5tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c5tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldTlenMask) >> RegisterMdma_c5tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c5tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c5tcrFieldTlenShift))
}

const (
	RegisterMdma_c5tcrFieldPkeShift = 25
	RegisterMdma_c5tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c5tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c5tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c5tcrFieldPamShift = 26
	RegisterMdma_c5tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c5tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldPamMask) >> RegisterMdma_c5tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c5tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c5tcrFieldPamShift))
}

const (
	RegisterMdma_c5tcrFieldTrgmShift = 28
	RegisterMdma_c5tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c5tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldTrgmMask) >> RegisterMdma_c5tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c5tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c5tcrFieldTrgmShift))
}

const (
	RegisterMdma_c5tcrFieldSwrmShift = 30
	RegisterMdma_c5tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c5tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c5tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c5tcrFieldBwmShift = 31
	RegisterMdma_c5tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c5tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c5tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tcrFieldBwmMask)
	}
}

// registerMdma_c5bndtrType MDMA Channel x block number of data register
type registerMdma_c5bndtrType uint32

const (
	RegisterMdma_c5bndtrFieldBndtShift = 0
	RegisterMdma_c5bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c5bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5bndtrFieldBndtMask) >> RegisterMdma_c5bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c5bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c5bndtrFieldBndtShift))
}

const (
	RegisterMdma_c5bndtrFieldBrsumShift = 18
	RegisterMdma_c5bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c5bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c5bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c5bndtrFieldBrdumShift = 19
	RegisterMdma_c5bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c5bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c5bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c5bndtrFieldBrcShift = 20
	RegisterMdma_c5bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c5bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5bndtrFieldBrcMask) >> RegisterMdma_c5bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c5bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c5bndtrFieldBrcShift))
}

// registerMdma_c5sarType MDMA channel x source address register
type registerMdma_c5sarType uint32

const (
	RegisterMdma_c5sarFieldSarShift = 0
	RegisterMdma_c5sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c5sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5sarFieldSarMask) >> RegisterMdma_c5sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c5sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5sarFieldSarMask)|(uint32(value)<<RegisterMdma_c5sarFieldSarShift))
}

// registerMdma_c5darType MDMA channel x destination address register
type registerMdma_c5darType uint32

const (
	RegisterMdma_c5darFieldDarShift = 0
	RegisterMdma_c5darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c5darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5darFieldDarMask) >> RegisterMdma_c5darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c5darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5darFieldDarMask)|(uint32(value)<<RegisterMdma_c5darFieldDarShift))
}

// registerMdma_c5brurType MDMA channel x Block Repeat address Update register
type registerMdma_c5brurType uint32

const (
	RegisterMdma_c5brurFieldSuvShift = 0
	RegisterMdma_c5brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c5brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5brurFieldSuvMask) >> RegisterMdma_c5brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c5brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c5brurFieldSuvShift))
}

const (
	RegisterMdma_c5brurFieldDuvShift = 16
	RegisterMdma_c5brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c5brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5brurFieldDuvMask) >> RegisterMdma_c5brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c5brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c5brurFieldDuvShift))
}

// registerMdma_c5larType MDMA channel x Link Address register
type registerMdma_c5larType uint32

const (
	RegisterMdma_c5larFieldLarShift = 0
	RegisterMdma_c5larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c5larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5larFieldLarMask) >> RegisterMdma_c5larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c5larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5larFieldLarMask)|(uint32(value)<<RegisterMdma_c5larFieldLarShift))
}

// registerMdma_c5tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c5tbrType uint32

const (
	RegisterMdma_c5tbrFieldTselShift = 0
	RegisterMdma_c5tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c5tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tbrFieldTselMask) >> RegisterMdma_c5tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c5tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c5tbrFieldTselShift))
}

const (
	RegisterMdma_c5tbrFieldSbusShift = 16
	RegisterMdma_c5tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c5tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c5tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c5tbrFieldDbusShift = 17
	RegisterMdma_c5tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c5tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c5tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c5tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5tbrFieldDbusMask)
	}
}

// registerMdma_c5marType MDMA channel x Mask address register
type registerMdma_c5marType uint32

const (
	RegisterMdma_c5marFieldMarShift = 0
	RegisterMdma_c5marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c5marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5marFieldMarMask) >> RegisterMdma_c5marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c5marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5marFieldMarMask)|(uint32(value)<<RegisterMdma_c5marFieldMarShift))
}

// registerMdma_c5mdrType MDMA channel x Mask Data register
type registerMdma_c5mdrType uint32

const (
	RegisterMdma_c5mdrFieldMdrShift = 0
	RegisterMdma_c5mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c5mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c5mdrFieldMdrMask) >> RegisterMdma_c5mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c5mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c5mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c5mdrFieldMdrShift))
}

// registerMdma_c6isrType MDMA channel x interrupt/status register
type registerMdma_c6isrType uint32

const (
	RegisterMdma_c6isrFieldTeif6Shift = 0
	RegisterMdma_c6isrFieldTeif6Mask  = 0x1
)

// GetTeif6 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c6isrType) GetTeif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6isrFieldTeif6Mask) != 0
}

// SetTeif6 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c6isrType) SetTeif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6isrFieldTeif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6isrFieldTeif6Mask)
	}
}

const (
	RegisterMdma_c6isrFieldCtcif6Shift = 1
	RegisterMdma_c6isrFieldCtcif6Mask  = 0x2
)

// GetCtcif6 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c6isrType) GetCtcif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6isrFieldCtcif6Mask) != 0
}

// SetCtcif6 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c6isrType) SetCtcif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6isrFieldCtcif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6isrFieldCtcif6Mask)
	}
}

const (
	RegisterMdma_c6isrFieldBrtif6Shift = 2
	RegisterMdma_c6isrFieldBrtif6Mask  = 0x4
)

// GetBrtif6 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c6isrType) GetBrtif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6isrFieldBrtif6Mask) != 0
}

// SetBrtif6 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c6isrType) SetBrtif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6isrFieldBrtif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6isrFieldBrtif6Mask)
	}
}

const (
	RegisterMdma_c6isrFieldBtif6Shift = 3
	RegisterMdma_c6isrFieldBtif6Mask  = 0x8
)

// GetBtif6 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c6isrType) GetBtif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6isrFieldBtif6Mask) != 0
}

// SetBtif6 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c6isrType) SetBtif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6isrFieldBtif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6isrFieldBtif6Mask)
	}
}

const (
	RegisterMdma_c6isrFieldTcif6Shift = 4
	RegisterMdma_c6isrFieldTcif6Mask  = 0x10
)

// GetTcif6 channel x buffer transfer complete
func (r *registerMdma_c6isrType) GetTcif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6isrFieldTcif6Mask) != 0
}

// SetTcif6 channel x buffer transfer complete
func (r *registerMdma_c6isrType) SetTcif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6isrFieldTcif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6isrFieldTcif6Mask)
	}
}

const (
	RegisterMdma_c6isrFieldCrqa6Shift = 16
	RegisterMdma_c6isrFieldCrqa6Mask  = 0x10000
)

// GetCrqa6 channel x request active flag
func (r *registerMdma_c6isrType) GetCrqa6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6isrFieldCrqa6Mask) != 0
}

// SetCrqa6 channel x request active flag
func (r *registerMdma_c6isrType) SetCrqa6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6isrFieldCrqa6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6isrFieldCrqa6Mask)
	}
}

// registerMdma_c6ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c6ifcrType uint32

const (
	RegisterMdma_c6ifcrFieldCteif6Shift = 0
	RegisterMdma_c6ifcrFieldCteif6Mask  = 0x1
)

// GetCteif6 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c6ifcrType) GetCteif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6ifcrFieldCteif6Mask) != 0
}

// SetCteif6 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c6ifcrType) SetCteif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6ifcrFieldCteif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6ifcrFieldCteif6Mask)
	}
}

const (
	RegisterMdma_c6ifcrFieldCctcif6Shift = 1
	RegisterMdma_c6ifcrFieldCctcif6Mask  = 0x2
)

// GetCctcif6 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c6ifcrType) GetCctcif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6ifcrFieldCctcif6Mask) != 0
}

// SetCctcif6 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c6ifcrType) SetCctcif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6ifcrFieldCctcif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6ifcrFieldCctcif6Mask)
	}
}

const (
	RegisterMdma_c6ifcrFieldCbrtif6Shift = 2
	RegisterMdma_c6ifcrFieldCbrtif6Mask  = 0x4
)

// GetCbrtif6 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c6ifcrType) GetCbrtif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6ifcrFieldCbrtif6Mask) != 0
}

// SetCbrtif6 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c6ifcrType) SetCbrtif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6ifcrFieldCbrtif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6ifcrFieldCbrtif6Mask)
	}
}

const (
	RegisterMdma_c6ifcrFieldCbtif6Shift = 3
	RegisterMdma_c6ifcrFieldCbtif6Mask  = 0x8
)

// GetCbtif6 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c6ifcrType) GetCbtif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6ifcrFieldCbtif6Mask) != 0
}

// SetCbtif6 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c6ifcrType) SetCbtif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6ifcrFieldCbtif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6ifcrFieldCbtif6Mask)
	}
}

const (
	RegisterMdma_c6ifcrFieldCltcif6Shift = 4
	RegisterMdma_c6ifcrFieldCltcif6Mask  = 0x10
)

// GetCltcif6 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c6ifcrType) GetCltcif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6ifcrFieldCltcif6Mask) != 0
}

// SetCltcif6 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c6ifcrType) SetCltcif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6ifcrFieldCltcif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6ifcrFieldCltcif6Mask)
	}
}

// registerMdma_c6esrType MDMA Channel x error status register
type registerMdma_c6esrType uint32

const (
	RegisterMdma_c6esrFieldTeaShift = 0
	RegisterMdma_c6esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c6esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6esrFieldTeaMask) >> RegisterMdma_c6esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c6esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c6esrFieldTeaShift))
}

const (
	RegisterMdma_c6esrFieldTedShift = 7
	RegisterMdma_c6esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c6esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c6esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6esrFieldTedMask)
	}
}

const (
	RegisterMdma_c6esrFieldTeldShift = 8
	RegisterMdma_c6esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c6esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c6esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c6esrFieldTemdShift = 9
	RegisterMdma_c6esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c6esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c6esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c6esrFieldAseShift = 10
	RegisterMdma_c6esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c6esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c6esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6esrFieldAseMask)
	}
}

const (
	RegisterMdma_c6esrFieldBseShift = 11
	RegisterMdma_c6esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c6esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c6esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6esrFieldBseMask)
	}
}

// registerMdma_c6crType This register is used to control the concerned channel.
type registerMdma_c6crType uint32

const (
	RegisterMdma_c6crFieldEnShift = 0
	RegisterMdma_c6crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c6crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c6crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6crFieldEnMask)
	}
}

const (
	RegisterMdma_c6crFieldTeieShift = 1
	RegisterMdma_c6crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c6crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c6crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6crFieldTeieMask)
	}
}

const (
	RegisterMdma_c6crFieldCtcieShift = 2
	RegisterMdma_c6crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c6crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c6crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c6crFieldBrtieShift = 3
	RegisterMdma_c6crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c6crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c6crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c6crFieldBtieShift = 4
	RegisterMdma_c6crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c6crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c6crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6crFieldBtieMask)
	}
}

const (
	RegisterMdma_c6crFieldTcieShift = 5
	RegisterMdma_c6crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c6crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c6crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6crFieldTcieMask)
	}
}

const (
	RegisterMdma_c6crFieldPlShift = 6
	RegisterMdma_c6crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c6crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6crFieldPlMask) >> RegisterMdma_c6crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c6crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6crFieldPlMask)|(uint32(value)<<RegisterMdma_c6crFieldPlShift))
}

const (
	RegisterMdma_c6crFieldBexShift = 12
	RegisterMdma_c6crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c6crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c6crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6crFieldBexMask)
	}
}

const (
	RegisterMdma_c6crFieldHexShift = 13
	RegisterMdma_c6crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c6crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c6crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6crFieldHexMask)
	}
}

const (
	RegisterMdma_c6crFieldWexShift = 14
	RegisterMdma_c6crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c6crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c6crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6crFieldWexMask)
	}
}

const (
	RegisterMdma_c6crFieldSwrqShift = 16
	RegisterMdma_c6crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c6crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6crFieldSwrqMask)
	}
}

// registerMdma_c6tcrType This register is used to configure the concerned channel.
type registerMdma_c6tcrType uint32

const (
	RegisterMdma_c6tcrFieldSincShift = 0
	RegisterMdma_c6tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c6tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldSincMask) >> RegisterMdma_c6tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c6tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c6tcrFieldSincShift))
}

const (
	RegisterMdma_c6tcrFieldDincShift = 2
	RegisterMdma_c6tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c6tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldDincMask) >> RegisterMdma_c6tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c6tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c6tcrFieldDincShift))
}

const (
	RegisterMdma_c6tcrFieldSsizeShift = 4
	RegisterMdma_c6tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c6tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldSsizeMask) >> RegisterMdma_c6tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c6tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c6tcrFieldSsizeShift))
}

const (
	RegisterMdma_c6tcrFieldDsizeShift = 6
	RegisterMdma_c6tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c6tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldDsizeMask) >> RegisterMdma_c6tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c6tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c6tcrFieldDsizeShift))
}

const (
	RegisterMdma_c6tcrFieldSincosShift = 8
	RegisterMdma_c6tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c6tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldSincosMask) >> RegisterMdma_c6tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c6tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c6tcrFieldSincosShift))
}

const (
	RegisterMdma_c6tcrFieldDincosShift = 10
	RegisterMdma_c6tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c6tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldDincosMask) >> RegisterMdma_c6tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c6tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c6tcrFieldDincosShift))
}

const (
	RegisterMdma_c6tcrFieldSburstShift = 12
	RegisterMdma_c6tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c6tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldSburstMask) >> RegisterMdma_c6tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c6tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c6tcrFieldSburstShift))
}

const (
	RegisterMdma_c6tcrFieldDburstShift = 15
	RegisterMdma_c6tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c6tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldDburstMask) >> RegisterMdma_c6tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c6tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c6tcrFieldDburstShift))
}

const (
	RegisterMdma_c6tcrFieldTlenShift = 18
	RegisterMdma_c6tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c6tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldTlenMask) >> RegisterMdma_c6tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c6tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c6tcrFieldTlenShift))
}

const (
	RegisterMdma_c6tcrFieldPkeShift = 25
	RegisterMdma_c6tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c6tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c6tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c6tcrFieldPamShift = 26
	RegisterMdma_c6tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c6tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldPamMask) >> RegisterMdma_c6tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c6tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c6tcrFieldPamShift))
}

const (
	RegisterMdma_c6tcrFieldTrgmShift = 28
	RegisterMdma_c6tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c6tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldTrgmMask) >> RegisterMdma_c6tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c6tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c6tcrFieldTrgmShift))
}

const (
	RegisterMdma_c6tcrFieldSwrmShift = 30
	RegisterMdma_c6tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c6tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c6tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c6tcrFieldBwmShift = 31
	RegisterMdma_c6tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c6tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c6tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tcrFieldBwmMask)
	}
}

// registerMdma_c6bndtrType MDMA Channel x block number of data register
type registerMdma_c6bndtrType uint32

const (
	RegisterMdma_c6bndtrFieldBndtShift = 0
	RegisterMdma_c6bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c6bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6bndtrFieldBndtMask) >> RegisterMdma_c6bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c6bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c6bndtrFieldBndtShift))
}

const (
	RegisterMdma_c6bndtrFieldBrsumShift = 18
	RegisterMdma_c6bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c6bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c6bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c6bndtrFieldBrdumShift = 19
	RegisterMdma_c6bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c6bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c6bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c6bndtrFieldBrcShift = 20
	RegisterMdma_c6bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0
func (r *registerMdma_c6bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6bndtrFieldBrcMask) >> RegisterMdma_c6bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0
func (r *registerMdma_c6bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c6bndtrFieldBrcShift))
}

// registerMdma_c6sarType MDMA channel x source address register
type registerMdma_c6sarType uint32

const (
	RegisterMdma_c6sarFieldSarShift = 0
	RegisterMdma_c6sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c6sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6sarFieldSarMask) >> RegisterMdma_c6sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c6sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6sarFieldSarMask)|(uint32(value)<<RegisterMdma_c6sarFieldSarShift))
}

// registerMdma_c6darType MDMA channel x destination address register
type registerMdma_c6darType uint32

const (
	RegisterMdma_c6darFieldDarShift = 0
	RegisterMdma_c6darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c6darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6darFieldDarMask) >> RegisterMdma_c6darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c6darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6darFieldDarMask)|(uint32(value)<<RegisterMdma_c6darFieldDarShift))
}

// registerMdma_c6brurType MDMA channel x Block Repeat address Update register
type registerMdma_c6brurType uint32

const (
	RegisterMdma_c6brurFieldSuvShift = 0
	RegisterMdma_c6brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c6brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6brurFieldSuvMask) >> RegisterMdma_c6brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c6brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c6brurFieldSuvShift))
}

const (
	RegisterMdma_c6brurFieldDuvShift = 16
	RegisterMdma_c6brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c6brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6brurFieldDuvMask) >> RegisterMdma_c6brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c6brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c6brurFieldDuvShift))
}

// registerMdma_c6larType MDMA channel x Link Address register
type registerMdma_c6larType uint32

const (
	RegisterMdma_c6larFieldLarShift = 0
	RegisterMdma_c6larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c6larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6larFieldLarMask) >> RegisterMdma_c6larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c6larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6larFieldLarMask)|(uint32(value)<<RegisterMdma_c6larFieldLarShift))
}

// registerMdma_c6tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c6tbrType uint32

const (
	RegisterMdma_c6tbrFieldTselShift = 0
	RegisterMdma_c6tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c6tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tbrFieldTselMask) >> RegisterMdma_c6tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c6tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c6tbrFieldTselShift))
}

const (
	RegisterMdma_c6tbrFieldSbusShift = 16
	RegisterMdma_c6tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c6tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c6tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c6tbrFieldDbusShift = 17
	RegisterMdma_c6tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c6tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c6tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c6tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6tbrFieldDbusMask)
	}
}

// registerMdma_c6marType MDMA channel x Mask address register
type registerMdma_c6marType uint32

const (
	RegisterMdma_c6marFieldMarShift = 0
	RegisterMdma_c6marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c6marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6marFieldMarMask) >> RegisterMdma_c6marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c6marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6marFieldMarMask)|(uint32(value)<<RegisterMdma_c6marFieldMarShift))
}

// registerMdma_c6mdrType MDMA channel x Mask Data register
type registerMdma_c6mdrType uint32

const (
	RegisterMdma_c6mdrFieldMdrShift = 0
	RegisterMdma_c6mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c6mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c6mdrFieldMdrMask) >> RegisterMdma_c6mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c6mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c6mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c6mdrFieldMdrShift))
}

// registerMdma_c7isrType MDMA channel x interrupt/status register
type registerMdma_c7isrType uint32

const (
	RegisterMdma_c7isrFieldTeif7Shift = 0
	RegisterMdma_c7isrFieldTeif7Mask  = 0x1
)

// GetTeif7 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c7isrType) GetTeif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7isrFieldTeif7Mask) != 0
}

// SetTeif7 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c7isrType) SetTeif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7isrFieldTeif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7isrFieldTeif7Mask)
	}
}

const (
	RegisterMdma_c7isrFieldCtcif7Shift = 1
	RegisterMdma_c7isrFieldCtcif7Mask  = 0x2
)

// GetCtcif7 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c7isrType) GetCtcif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7isrFieldCtcif7Mask) != 0
}

// SetCtcif7 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c7isrType) SetCtcif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7isrFieldCtcif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7isrFieldCtcif7Mask)
	}
}

const (
	RegisterMdma_c7isrFieldBrtif7Shift = 2
	RegisterMdma_c7isrFieldBrtif7Mask  = 0x4
)

// GetBrtif7 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c7isrType) GetBrtif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7isrFieldBrtif7Mask) != 0
}

// SetBrtif7 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c7isrType) SetBrtif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7isrFieldBrtif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7isrFieldBrtif7Mask)
	}
}

const (
	RegisterMdma_c7isrFieldBtif7Shift = 3
	RegisterMdma_c7isrFieldBtif7Mask  = 0x8
)

// GetBtif7 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c7isrType) GetBtif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7isrFieldBtif7Mask) != 0
}

// SetBtif7 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c7isrType) SetBtif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7isrFieldBtif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7isrFieldBtif7Mask)
	}
}

const (
	RegisterMdma_c7isrFieldTcif7Shift = 4
	RegisterMdma_c7isrFieldTcif7Mask  = 0x10
)

// GetTcif7 channel x buffer transfer complete
func (r *registerMdma_c7isrType) GetTcif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7isrFieldTcif7Mask) != 0
}

// SetTcif7 channel x buffer transfer complete
func (r *registerMdma_c7isrType) SetTcif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7isrFieldTcif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7isrFieldTcif7Mask)
	}
}

const (
	RegisterMdma_c7isrFieldCrqa7Shift = 16
	RegisterMdma_c7isrFieldCrqa7Mask  = 0x10000
)

// GetCrqa7 channel x request active flag
func (r *registerMdma_c7isrType) GetCrqa7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7isrFieldCrqa7Mask) != 0
}

// SetCrqa7 channel x request active flag
func (r *registerMdma_c7isrType) SetCrqa7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7isrFieldCrqa7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7isrFieldCrqa7Mask)
	}
}

// registerMdma_c7ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c7ifcrType uint32

const (
	RegisterMdma_c7ifcrFieldCteif7Shift = 0
	RegisterMdma_c7ifcrFieldCteif7Mask  = 0x1
)

// GetCteif7 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c7ifcrType) GetCteif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7ifcrFieldCteif7Mask) != 0
}

// SetCteif7 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c7ifcrType) SetCteif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7ifcrFieldCteif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7ifcrFieldCteif7Mask)
	}
}

const (
	RegisterMdma_c7ifcrFieldCctcif7Shift = 1
	RegisterMdma_c7ifcrFieldCctcif7Mask  = 0x2
)

// GetCctcif7 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c7ifcrType) GetCctcif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7ifcrFieldCctcif7Mask) != 0
}

// SetCctcif7 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c7ifcrType) SetCctcif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7ifcrFieldCctcif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7ifcrFieldCctcif7Mask)
	}
}

const (
	RegisterMdma_c7ifcrFieldCbrtif7Shift = 2
	RegisterMdma_c7ifcrFieldCbrtif7Mask  = 0x4
)

// GetCbrtif7 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c7ifcrType) GetCbrtif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7ifcrFieldCbrtif7Mask) != 0
}

// SetCbrtif7 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c7ifcrType) SetCbrtif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7ifcrFieldCbrtif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7ifcrFieldCbrtif7Mask)
	}
}

const (
	RegisterMdma_c7ifcrFieldCbtif7Shift = 3
	RegisterMdma_c7ifcrFieldCbtif7Mask  = 0x8
)

// GetCbtif7 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c7ifcrType) GetCbtif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7ifcrFieldCbtif7Mask) != 0
}

// SetCbtif7 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c7ifcrType) SetCbtif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7ifcrFieldCbtif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7ifcrFieldCbtif7Mask)
	}
}

const (
	RegisterMdma_c7ifcrFieldCltcif7Shift = 4
	RegisterMdma_c7ifcrFieldCltcif7Mask  = 0x10
)

// GetCltcif7 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c7ifcrType) GetCltcif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7ifcrFieldCltcif7Mask) != 0
}

// SetCltcif7 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c7ifcrType) SetCltcif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7ifcrFieldCltcif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7ifcrFieldCltcif7Mask)
	}
}

// registerMdma_c7esrType MDMA Channel x error status register
type registerMdma_c7esrType uint32

const (
	RegisterMdma_c7esrFieldTeaShift = 0
	RegisterMdma_c7esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c7esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7esrFieldTeaMask) >> RegisterMdma_c7esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c7esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c7esrFieldTeaShift))
}

const (
	RegisterMdma_c7esrFieldTedShift = 7
	RegisterMdma_c7esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c7esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c7esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7esrFieldTedMask)
	}
}

const (
	RegisterMdma_c7esrFieldTeldShift = 8
	RegisterMdma_c7esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c7esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c7esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c7esrFieldTemdShift = 9
	RegisterMdma_c7esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c7esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c7esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c7esrFieldAseShift = 10
	RegisterMdma_c7esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c7esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c7esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7esrFieldAseMask)
	}
}

const (
	RegisterMdma_c7esrFieldBseShift = 11
	RegisterMdma_c7esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c7esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c7esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7esrFieldBseMask)
	}
}

// registerMdma_c7crType This register is used to control the concerned channel.
type registerMdma_c7crType uint32

const (
	RegisterMdma_c7crFieldEnShift = 0
	RegisterMdma_c7crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c7crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c7crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7crFieldEnMask)
	}
}

const (
	RegisterMdma_c7crFieldTeieShift = 1
	RegisterMdma_c7crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c7crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c7crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7crFieldTeieMask)
	}
}

const (
	RegisterMdma_c7crFieldCtcieShift = 2
	RegisterMdma_c7crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c7crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c7crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c7crFieldBrtieShift = 3
	RegisterMdma_c7crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c7crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c7crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c7crFieldBtieShift = 4
	RegisterMdma_c7crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c7crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c7crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7crFieldBtieMask)
	}
}

const (
	RegisterMdma_c7crFieldTcieShift = 5
	RegisterMdma_c7crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c7crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c7crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7crFieldTcieMask)
	}
}

const (
	RegisterMdma_c7crFieldPlShift = 6
	RegisterMdma_c7crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c7crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7crFieldPlMask) >> RegisterMdma_c7crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c7crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7crFieldPlMask)|(uint32(value)<<RegisterMdma_c7crFieldPlShift))
}

const (
	RegisterMdma_c7crFieldBexShift = 12
	RegisterMdma_c7crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c7crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c7crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7crFieldBexMask)
	}
}

const (
	RegisterMdma_c7crFieldHexShift = 13
	RegisterMdma_c7crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c7crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c7crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7crFieldHexMask)
	}
}

const (
	RegisterMdma_c7crFieldWexShift = 14
	RegisterMdma_c7crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c7crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c7crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7crFieldWexMask)
	}
}

const (
	RegisterMdma_c7crFieldSwrqShift = 16
	RegisterMdma_c7crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c7crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7crFieldSwrqMask)
	}
}

// registerMdma_c7tcrType This register is used to configure the concerned channel.
type registerMdma_c7tcrType uint32

const (
	RegisterMdma_c7tcrFieldSincShift = 0
	RegisterMdma_c7tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c7tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldSincMask) >> RegisterMdma_c7tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c7tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c7tcrFieldSincShift))
}

const (
	RegisterMdma_c7tcrFieldDincShift = 2
	RegisterMdma_c7tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c7tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldDincMask) >> RegisterMdma_c7tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c7tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c7tcrFieldDincShift))
}

const (
	RegisterMdma_c7tcrFieldSsizeShift = 4
	RegisterMdma_c7tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c7tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldSsizeMask) >> RegisterMdma_c7tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c7tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c7tcrFieldSsizeShift))
}

const (
	RegisterMdma_c7tcrFieldDsizeShift = 6
	RegisterMdma_c7tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c7tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldDsizeMask) >> RegisterMdma_c7tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c7tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c7tcrFieldDsizeShift))
}

const (
	RegisterMdma_c7tcrFieldSincosShift = 8
	RegisterMdma_c7tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c7tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldSincosMask) >> RegisterMdma_c7tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c7tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c7tcrFieldSincosShift))
}

const (
	RegisterMdma_c7tcrFieldDincosShift = 10
	RegisterMdma_c7tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c7tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldDincosMask) >> RegisterMdma_c7tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c7tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c7tcrFieldDincosShift))
}

const (
	RegisterMdma_c7tcrFieldSburstShift = 12
	RegisterMdma_c7tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c7tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldSburstMask) >> RegisterMdma_c7tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c7tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c7tcrFieldSburstShift))
}

const (
	RegisterMdma_c7tcrFieldDburstShift = 15
	RegisterMdma_c7tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c7tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldDburstMask) >> RegisterMdma_c7tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c7tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c7tcrFieldDburstShift))
}

const (
	RegisterMdma_c7tcrFieldTlenShift = 18
	RegisterMdma_c7tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c7tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldTlenMask) >> RegisterMdma_c7tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c7tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c7tcrFieldTlenShift))
}

const (
	RegisterMdma_c7tcrFieldPkeShift = 25
	RegisterMdma_c7tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c7tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c7tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c7tcrFieldPamShift = 26
	RegisterMdma_c7tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c7tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldPamMask) >> RegisterMdma_c7tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c7tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c7tcrFieldPamShift))
}

const (
	RegisterMdma_c7tcrFieldTrgmShift = 28
	RegisterMdma_c7tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c7tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldTrgmMask) >> RegisterMdma_c7tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c7tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c7tcrFieldTrgmShift))
}

const (
	RegisterMdma_c7tcrFieldSwrmShift = 30
	RegisterMdma_c7tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c7tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c7tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c7tcrFieldBwmShift = 31
	RegisterMdma_c7tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c7tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c7tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tcrFieldBwmMask)
	}
}

// registerMdma_c7bndtrType MDMA Channel x block number of data register
type registerMdma_c7bndtrType uint32

const (
	RegisterMdma_c7bndtrFieldBndtShift = 0
	RegisterMdma_c7bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c7bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7bndtrFieldBndtMask) >> RegisterMdma_c7bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c7bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c7bndtrFieldBndtShift))
}

const (
	RegisterMdma_c7bndtrFieldBrsumShift = 18
	RegisterMdma_c7bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c7bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c7bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c7bndtrFieldBrdumShift = 19
	RegisterMdma_c7bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c7bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c7bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c7bndtrFieldBrcShift = 20
	RegisterMdma_c7bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c7bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7bndtrFieldBrcMask) >> RegisterMdma_c7bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c7bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c7bndtrFieldBrcShift))
}

// registerMdma_c7sarType MDMA channel x source address register
type registerMdma_c7sarType uint32

const (
	RegisterMdma_c7sarFieldSarShift = 0
	RegisterMdma_c7sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c7sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7sarFieldSarMask) >> RegisterMdma_c7sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c7sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7sarFieldSarMask)|(uint32(value)<<RegisterMdma_c7sarFieldSarShift))
}

// registerMdma_c7darType MDMA channel x destination address register
type registerMdma_c7darType uint32

const (
	RegisterMdma_c7darFieldDarShift = 0
	RegisterMdma_c7darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c7darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7darFieldDarMask) >> RegisterMdma_c7darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c7darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7darFieldDarMask)|(uint32(value)<<RegisterMdma_c7darFieldDarShift))
}

// registerMdma_c7brurType MDMA channel x Block Repeat address Update register
type registerMdma_c7brurType uint32

const (
	RegisterMdma_c7brurFieldSuvShift = 0
	RegisterMdma_c7brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c7brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7brurFieldSuvMask) >> RegisterMdma_c7brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c7brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c7brurFieldSuvShift))
}

const (
	RegisterMdma_c7brurFieldDuvShift = 16
	RegisterMdma_c7brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c7brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7brurFieldDuvMask) >> RegisterMdma_c7brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c7brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c7brurFieldDuvShift))
}

// registerMdma_c7larType MDMA channel x Link Address register
type registerMdma_c7larType uint32

const (
	RegisterMdma_c7larFieldLarShift = 0
	RegisterMdma_c7larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c7larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7larFieldLarMask) >> RegisterMdma_c7larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c7larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7larFieldLarMask)|(uint32(value)<<RegisterMdma_c7larFieldLarShift))
}

// registerMdma_c7tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c7tbrType uint32

const (
	RegisterMdma_c7tbrFieldTselShift = 0
	RegisterMdma_c7tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c7tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tbrFieldTselMask) >> RegisterMdma_c7tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c7tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c7tbrFieldTselShift))
}

const (
	RegisterMdma_c7tbrFieldSbusShift = 16
	RegisterMdma_c7tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c7tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c7tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c7tbrFieldDbusShift = 17
	RegisterMdma_c7tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c7tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c7tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c7tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7tbrFieldDbusMask)
	}
}

// registerMdma_c7marType MDMA channel x Mask address register
type registerMdma_c7marType uint32

const (
	RegisterMdma_c7marFieldMarShift = 0
	RegisterMdma_c7marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c7marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7marFieldMarMask) >> RegisterMdma_c7marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c7marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7marFieldMarMask)|(uint32(value)<<RegisterMdma_c7marFieldMarShift))
}

// registerMdma_c7mdrType MDMA channel x Mask Data register
type registerMdma_c7mdrType uint32

const (
	RegisterMdma_c7mdrFieldMdrShift = 0
	RegisterMdma_c7mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c7mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c7mdrFieldMdrMask) >> RegisterMdma_c7mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c7mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c7mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c7mdrFieldMdrShift))
}

// registerMdma_c8isrType MDMA channel x interrupt/status register
type registerMdma_c8isrType uint32

const (
	RegisterMdma_c8isrFieldTeif8Shift = 0
	RegisterMdma_c8isrFieldTeif8Mask  = 0x1
)

// GetTeif8 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c8isrType) GetTeif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8isrFieldTeif8Mask) != 0
}

// SetTeif8 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c8isrType) SetTeif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8isrFieldTeif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8isrFieldTeif8Mask)
	}
}

const (
	RegisterMdma_c8isrFieldCtcif8Shift = 1
	RegisterMdma_c8isrFieldCtcif8Mask  = 0x2
)

// GetCtcif8 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c8isrType) GetCtcif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8isrFieldCtcif8Mask) != 0
}

// SetCtcif8 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c8isrType) SetCtcif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8isrFieldCtcif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8isrFieldCtcif8Mask)
	}
}

const (
	RegisterMdma_c8isrFieldBrtif8Shift = 2
	RegisterMdma_c8isrFieldBrtif8Mask  = 0x4
)

// GetBrtif8 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c8isrType) GetBrtif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8isrFieldBrtif8Mask) != 0
}

// SetBrtif8 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c8isrType) SetBrtif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8isrFieldBrtif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8isrFieldBrtif8Mask)
	}
}

const (
	RegisterMdma_c8isrFieldBtif8Shift = 3
	RegisterMdma_c8isrFieldBtif8Mask  = 0x8
)

// GetBtif8 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c8isrType) GetBtif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8isrFieldBtif8Mask) != 0
}

// SetBtif8 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c8isrType) SetBtif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8isrFieldBtif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8isrFieldBtif8Mask)
	}
}

const (
	RegisterMdma_c8isrFieldTcif8Shift = 4
	RegisterMdma_c8isrFieldTcif8Mask  = 0x10
)

// GetTcif8 channel x buffer transfer complete
func (r *registerMdma_c8isrType) GetTcif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8isrFieldTcif8Mask) != 0
}

// SetTcif8 channel x buffer transfer complete
func (r *registerMdma_c8isrType) SetTcif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8isrFieldTcif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8isrFieldTcif8Mask)
	}
}

const (
	RegisterMdma_c8isrFieldCrqa8Shift = 16
	RegisterMdma_c8isrFieldCrqa8Mask  = 0x10000
)

// GetCrqa8 channel x request active flag
func (r *registerMdma_c8isrType) GetCrqa8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8isrFieldCrqa8Mask) != 0
}

// SetCrqa8 channel x request active flag
func (r *registerMdma_c8isrType) SetCrqa8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8isrFieldCrqa8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8isrFieldCrqa8Mask)
	}
}

// registerMdma_c8ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c8ifcrType uint32

const (
	RegisterMdma_c8ifcrFieldCteif8Shift = 0
	RegisterMdma_c8ifcrFieldCteif8Mask  = 0x1
)

// GetCteif8 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c8ifcrType) GetCteif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8ifcrFieldCteif8Mask) != 0
}

// SetCteif8 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c8ifcrType) SetCteif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8ifcrFieldCteif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8ifcrFieldCteif8Mask)
	}
}

const (
	RegisterMdma_c8ifcrFieldCctcif8Shift = 1
	RegisterMdma_c8ifcrFieldCctcif8Mask  = 0x2
)

// GetCctcif8 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c8ifcrType) GetCctcif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8ifcrFieldCctcif8Mask) != 0
}

// SetCctcif8 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c8ifcrType) SetCctcif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8ifcrFieldCctcif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8ifcrFieldCctcif8Mask)
	}
}

const (
	RegisterMdma_c8ifcrFieldCbrtif8Shift = 2
	RegisterMdma_c8ifcrFieldCbrtif8Mask  = 0x4
)

// GetCbrtif8 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c8ifcrType) GetCbrtif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8ifcrFieldCbrtif8Mask) != 0
}

// SetCbrtif8 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c8ifcrType) SetCbrtif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8ifcrFieldCbrtif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8ifcrFieldCbrtif8Mask)
	}
}

const (
	RegisterMdma_c8ifcrFieldCbtif8Shift = 3
	RegisterMdma_c8ifcrFieldCbtif8Mask  = 0x8
)

// GetCbtif8 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c8ifcrType) GetCbtif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8ifcrFieldCbtif8Mask) != 0
}

// SetCbtif8 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c8ifcrType) SetCbtif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8ifcrFieldCbtif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8ifcrFieldCbtif8Mask)
	}
}

const (
	RegisterMdma_c8ifcrFieldCltcif8Shift = 4
	RegisterMdma_c8ifcrFieldCltcif8Mask  = 0x10
)

// GetCltcif8 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c8ifcrType) GetCltcif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8ifcrFieldCltcif8Mask) != 0
}

// SetCltcif8 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c8ifcrType) SetCltcif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8ifcrFieldCltcif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8ifcrFieldCltcif8Mask)
	}
}

// registerMdma_c8esrType MDMA Channel x error status register
type registerMdma_c8esrType uint32

const (
	RegisterMdma_c8esrFieldTeaShift = 0
	RegisterMdma_c8esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c8esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8esrFieldTeaMask) >> RegisterMdma_c8esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c8esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c8esrFieldTeaShift))
}

const (
	RegisterMdma_c8esrFieldTedShift = 7
	RegisterMdma_c8esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c8esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c8esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8esrFieldTedMask)
	}
}

const (
	RegisterMdma_c8esrFieldTeldShift = 8
	RegisterMdma_c8esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c8esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c8esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c8esrFieldTemdShift = 9
	RegisterMdma_c8esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c8esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c8esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c8esrFieldAseShift = 10
	RegisterMdma_c8esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c8esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c8esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8esrFieldAseMask)
	}
}

const (
	RegisterMdma_c8esrFieldBseShift = 11
	RegisterMdma_c8esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c8esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c8esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8esrFieldBseMask)
	}
}

// registerMdma_c8crType This register is used to control the concerned channel.
type registerMdma_c8crType uint32

const (
	RegisterMdma_c8crFieldEnShift = 0
	RegisterMdma_c8crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c8crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c8crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8crFieldEnMask)
	}
}

const (
	RegisterMdma_c8crFieldTeieShift = 1
	RegisterMdma_c8crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c8crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c8crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8crFieldTeieMask)
	}
}

const (
	RegisterMdma_c8crFieldCtcieShift = 2
	RegisterMdma_c8crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c8crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c8crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c8crFieldBrtieShift = 3
	RegisterMdma_c8crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c8crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c8crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c8crFieldBtieShift = 4
	RegisterMdma_c8crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c8crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c8crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8crFieldBtieMask)
	}
}

const (
	RegisterMdma_c8crFieldTcieShift = 5
	RegisterMdma_c8crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c8crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c8crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8crFieldTcieMask)
	}
}

const (
	RegisterMdma_c8crFieldPlShift = 6
	RegisterMdma_c8crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c8crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8crFieldPlMask) >> RegisterMdma_c8crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c8crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8crFieldPlMask)|(uint32(value)<<RegisterMdma_c8crFieldPlShift))
}

const (
	RegisterMdma_c8crFieldBexShift = 12
	RegisterMdma_c8crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c8crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c8crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8crFieldBexMask)
	}
}

const (
	RegisterMdma_c8crFieldHexShift = 13
	RegisterMdma_c8crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c8crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c8crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8crFieldHexMask)
	}
}

const (
	RegisterMdma_c8crFieldWexShift = 14
	RegisterMdma_c8crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c8crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c8crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8crFieldWexMask)
	}
}

const (
	RegisterMdma_c8crFieldSwrqShift = 16
	RegisterMdma_c8crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c8crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8crFieldSwrqMask)
	}
}

// registerMdma_c8tcrType This register is used to configure the concerned channel.
type registerMdma_c8tcrType uint32

const (
	RegisterMdma_c8tcrFieldSincShift = 0
	RegisterMdma_c8tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c8tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldSincMask) >> RegisterMdma_c8tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c8tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c8tcrFieldSincShift))
}

const (
	RegisterMdma_c8tcrFieldDincShift = 2
	RegisterMdma_c8tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c8tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldDincMask) >> RegisterMdma_c8tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c8tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c8tcrFieldDincShift))
}

const (
	RegisterMdma_c8tcrFieldSsizeShift = 4
	RegisterMdma_c8tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c8tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldSsizeMask) >> RegisterMdma_c8tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c8tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c8tcrFieldSsizeShift))
}

const (
	RegisterMdma_c8tcrFieldDsizeShift = 6
	RegisterMdma_c8tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c8tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldDsizeMask) >> RegisterMdma_c8tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c8tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c8tcrFieldDsizeShift))
}

const (
	RegisterMdma_c8tcrFieldSincosShift = 8
	RegisterMdma_c8tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c8tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldSincosMask) >> RegisterMdma_c8tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c8tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c8tcrFieldSincosShift))
}

const (
	RegisterMdma_c8tcrFieldDincosShift = 10
	RegisterMdma_c8tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c8tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldDincosMask) >> RegisterMdma_c8tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c8tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c8tcrFieldDincosShift))
}

const (
	RegisterMdma_c8tcrFieldSburstShift = 12
	RegisterMdma_c8tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c8tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldSburstMask) >> RegisterMdma_c8tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c8tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c8tcrFieldSburstShift))
}

const (
	RegisterMdma_c8tcrFieldDburstShift = 15
	RegisterMdma_c8tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c8tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldDburstMask) >> RegisterMdma_c8tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c8tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c8tcrFieldDburstShift))
}

const (
	RegisterMdma_c8tcrFieldTlenShift = 18
	RegisterMdma_c8tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c8tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldTlenMask) >> RegisterMdma_c8tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c8tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c8tcrFieldTlenShift))
}

const (
	RegisterMdma_c8tcrFieldPkeShift = 25
	RegisterMdma_c8tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c8tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c8tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c8tcrFieldPamShift = 26
	RegisterMdma_c8tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c8tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldPamMask) >> RegisterMdma_c8tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c8tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c8tcrFieldPamShift))
}

const (
	RegisterMdma_c8tcrFieldTrgmShift = 28
	RegisterMdma_c8tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c8tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldTrgmMask) >> RegisterMdma_c8tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c8tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c8tcrFieldTrgmShift))
}

const (
	RegisterMdma_c8tcrFieldSwrmShift = 30
	RegisterMdma_c8tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c8tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c8tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c8tcrFieldBwmShift = 31
	RegisterMdma_c8tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c8tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c8tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tcrFieldBwmMask)
	}
}

// registerMdma_c8bndtrType MDMA Channel x block number of data register
type registerMdma_c8bndtrType uint32

const (
	RegisterMdma_c8bndtrFieldBndtShift = 0
	RegisterMdma_c8bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c8bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8bndtrFieldBndtMask) >> RegisterMdma_c8bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c8bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c8bndtrFieldBndtShift))
}

const (
	RegisterMdma_c8bndtrFieldBrsumShift = 18
	RegisterMdma_c8bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c8bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c8bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c8bndtrFieldBrdumShift = 19
	RegisterMdma_c8bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c8bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c8bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c8bndtrFieldBrcShift = 20
	RegisterMdma_c8bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c8bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8bndtrFieldBrcMask) >> RegisterMdma_c8bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c8bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c8bndtrFieldBrcShift))
}

// registerMdma_c8sarType MDMA channel x source address register
type registerMdma_c8sarType uint32

const (
	RegisterMdma_c8sarFieldSarShift = 0
	RegisterMdma_c8sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c8sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8sarFieldSarMask) >> RegisterMdma_c8sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c8sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8sarFieldSarMask)|(uint32(value)<<RegisterMdma_c8sarFieldSarShift))
}

// registerMdma_c8darType MDMA channel x destination address register
type registerMdma_c8darType uint32

const (
	RegisterMdma_c8darFieldDarShift = 0
	RegisterMdma_c8darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c8darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8darFieldDarMask) >> RegisterMdma_c8darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c8darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8darFieldDarMask)|(uint32(value)<<RegisterMdma_c8darFieldDarShift))
}

// registerMdma_c8brurType MDMA channel x Block Repeat address Update register
type registerMdma_c8brurType uint32

const (
	RegisterMdma_c8brurFieldSuvShift = 0
	RegisterMdma_c8brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c8brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8brurFieldSuvMask) >> RegisterMdma_c8brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c8brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c8brurFieldSuvShift))
}

const (
	RegisterMdma_c8brurFieldDuvShift = 16
	RegisterMdma_c8brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c8brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8brurFieldDuvMask) >> RegisterMdma_c8brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c8brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c8brurFieldDuvShift))
}

// registerMdma_c8larType MDMA channel x Link Address register
type registerMdma_c8larType uint32

const (
	RegisterMdma_c8larFieldLarShift = 0
	RegisterMdma_c8larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c8larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8larFieldLarMask) >> RegisterMdma_c8larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c8larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8larFieldLarMask)|(uint32(value)<<RegisterMdma_c8larFieldLarShift))
}

// registerMdma_c8tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c8tbrType uint32

const (
	RegisterMdma_c8tbrFieldTselShift = 0
	RegisterMdma_c8tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c8tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tbrFieldTselMask) >> RegisterMdma_c8tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c8tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c8tbrFieldTselShift))
}

const (
	RegisterMdma_c8tbrFieldSbusShift = 16
	RegisterMdma_c8tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c8tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c8tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c8tbrFieldDbusShift = 17
	RegisterMdma_c8tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c8tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c8tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c8tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8tbrFieldDbusMask)
	}
}

// registerMdma_c8marType MDMA channel x Mask address register
type registerMdma_c8marType uint32

const (
	RegisterMdma_c8marFieldMarShift = 0
	RegisterMdma_c8marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c8marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8marFieldMarMask) >> RegisterMdma_c8marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c8marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8marFieldMarMask)|(uint32(value)<<RegisterMdma_c8marFieldMarShift))
}

// registerMdma_c8mdrType MDMA channel x Mask Data register
type registerMdma_c8mdrType uint32

const (
	RegisterMdma_c8mdrFieldMdrShift = 0
	RegisterMdma_c8mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c8mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c8mdrFieldMdrMask) >> RegisterMdma_c8mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c8mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c8mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c8mdrFieldMdrShift))
}

// registerMdma_c9isrType MDMA channel x interrupt/status register
type registerMdma_c9isrType uint32

const (
	RegisterMdma_c9isrFieldTeif9Shift = 0
	RegisterMdma_c9isrFieldTeif9Mask  = 0x1
)

// GetTeif9 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c9isrType) GetTeif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9isrFieldTeif9Mask) != 0
}

// SetTeif9 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c9isrType) SetTeif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9isrFieldTeif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9isrFieldTeif9Mask)
	}
}

const (
	RegisterMdma_c9isrFieldCtcif9Shift = 1
	RegisterMdma_c9isrFieldCtcif9Mask  = 0x2
)

// GetCtcif9 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c9isrType) GetCtcif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9isrFieldCtcif9Mask) != 0
}

// SetCtcif9 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c9isrType) SetCtcif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9isrFieldCtcif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9isrFieldCtcif9Mask)
	}
}

const (
	RegisterMdma_c9isrFieldBrtif9Shift = 2
	RegisterMdma_c9isrFieldBrtif9Mask  = 0x4
)

// GetBrtif9 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c9isrType) GetBrtif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9isrFieldBrtif9Mask) != 0
}

// SetBrtif9 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c9isrType) SetBrtif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9isrFieldBrtif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9isrFieldBrtif9Mask)
	}
}

const (
	RegisterMdma_c9isrFieldBtif9Shift = 3
	RegisterMdma_c9isrFieldBtif9Mask  = 0x8
)

// GetBtif9 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c9isrType) GetBtif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9isrFieldBtif9Mask) != 0
}

// SetBtif9 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c9isrType) SetBtif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9isrFieldBtif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9isrFieldBtif9Mask)
	}
}

const (
	RegisterMdma_c9isrFieldTcif9Shift = 4
	RegisterMdma_c9isrFieldTcif9Mask  = 0x10
)

// GetTcif9 channel x buffer transfer complete
func (r *registerMdma_c9isrType) GetTcif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9isrFieldTcif9Mask) != 0
}

// SetTcif9 channel x buffer transfer complete
func (r *registerMdma_c9isrType) SetTcif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9isrFieldTcif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9isrFieldTcif9Mask)
	}
}

const (
	RegisterMdma_c9isrFieldCrqa9Shift = 16
	RegisterMdma_c9isrFieldCrqa9Mask  = 0x10000
)

// GetCrqa9 channel x request active flag
func (r *registerMdma_c9isrType) GetCrqa9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9isrFieldCrqa9Mask) != 0
}

// SetCrqa9 channel x request active flag
func (r *registerMdma_c9isrType) SetCrqa9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9isrFieldCrqa9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9isrFieldCrqa9Mask)
	}
}

// registerMdma_c9ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c9ifcrType uint32

const (
	RegisterMdma_c9ifcrFieldCteif9Shift = 0
	RegisterMdma_c9ifcrFieldCteif9Mask  = 0x1
)

// GetCteif9 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c9ifcrType) GetCteif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9ifcrFieldCteif9Mask) != 0
}

// SetCteif9 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c9ifcrType) SetCteif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9ifcrFieldCteif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9ifcrFieldCteif9Mask)
	}
}

const (
	RegisterMdma_c9ifcrFieldCctcif9Shift = 1
	RegisterMdma_c9ifcrFieldCctcif9Mask  = 0x2
)

// GetCctcif9 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c9ifcrType) GetCctcif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9ifcrFieldCctcif9Mask) != 0
}

// SetCctcif9 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c9ifcrType) SetCctcif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9ifcrFieldCctcif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9ifcrFieldCctcif9Mask)
	}
}

const (
	RegisterMdma_c9ifcrFieldCbrtif9Shift = 2
	RegisterMdma_c9ifcrFieldCbrtif9Mask  = 0x4
)

// GetCbrtif9 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c9ifcrType) GetCbrtif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9ifcrFieldCbrtif9Mask) != 0
}

// SetCbrtif9 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c9ifcrType) SetCbrtif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9ifcrFieldCbrtif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9ifcrFieldCbrtif9Mask)
	}
}

const (
	RegisterMdma_c9ifcrFieldCbtif9Shift = 3
	RegisterMdma_c9ifcrFieldCbtif9Mask  = 0x8
)

// GetCbtif9 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c9ifcrType) GetCbtif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9ifcrFieldCbtif9Mask) != 0
}

// SetCbtif9 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c9ifcrType) SetCbtif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9ifcrFieldCbtif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9ifcrFieldCbtif9Mask)
	}
}

const (
	RegisterMdma_c9ifcrFieldCltcif9Shift = 4
	RegisterMdma_c9ifcrFieldCltcif9Mask  = 0x10
)

// GetCltcif9 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c9ifcrType) GetCltcif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9ifcrFieldCltcif9Mask) != 0
}

// SetCltcif9 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c9ifcrType) SetCltcif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9ifcrFieldCltcif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9ifcrFieldCltcif9Mask)
	}
}

// registerMdma_c9esrType MDMA Channel x error status register
type registerMdma_c9esrType uint32

const (
	RegisterMdma_c9esrFieldTeaShift = 0
	RegisterMdma_c9esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c9esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9esrFieldTeaMask) >> RegisterMdma_c9esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c9esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c9esrFieldTeaShift))
}

const (
	RegisterMdma_c9esrFieldTedShift = 7
	RegisterMdma_c9esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c9esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c9esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9esrFieldTedMask)
	}
}

const (
	RegisterMdma_c9esrFieldTeldShift = 8
	RegisterMdma_c9esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c9esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c9esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c9esrFieldTemdShift = 9
	RegisterMdma_c9esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c9esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c9esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c9esrFieldAseShift = 10
	RegisterMdma_c9esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c9esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c9esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9esrFieldAseMask)
	}
}

const (
	RegisterMdma_c9esrFieldBseShift = 11
	RegisterMdma_c9esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c9esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c9esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9esrFieldBseMask)
	}
}

// registerMdma_c9crType This register is used to control the concerned channel.
type registerMdma_c9crType uint32

const (
	RegisterMdma_c9crFieldEnShift = 0
	RegisterMdma_c9crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c9crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c9crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9crFieldEnMask)
	}
}

const (
	RegisterMdma_c9crFieldTeieShift = 1
	RegisterMdma_c9crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c9crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c9crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9crFieldTeieMask)
	}
}

const (
	RegisterMdma_c9crFieldCtcieShift = 2
	RegisterMdma_c9crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c9crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c9crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c9crFieldBrtieShift = 3
	RegisterMdma_c9crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c9crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c9crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c9crFieldBtieShift = 4
	RegisterMdma_c9crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c9crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c9crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9crFieldBtieMask)
	}
}

const (
	RegisterMdma_c9crFieldTcieShift = 5
	RegisterMdma_c9crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c9crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c9crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9crFieldTcieMask)
	}
}

const (
	RegisterMdma_c9crFieldPlShift = 6
	RegisterMdma_c9crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c9crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9crFieldPlMask) >> RegisterMdma_c9crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c9crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9crFieldPlMask)|(uint32(value)<<RegisterMdma_c9crFieldPlShift))
}

const (
	RegisterMdma_c9crFieldBexShift = 12
	RegisterMdma_c9crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c9crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c9crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9crFieldBexMask)
	}
}

const (
	RegisterMdma_c9crFieldHexShift = 13
	RegisterMdma_c9crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c9crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c9crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9crFieldHexMask)
	}
}

const (
	RegisterMdma_c9crFieldWexShift = 14
	RegisterMdma_c9crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c9crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c9crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9crFieldWexMask)
	}
}

const (
	RegisterMdma_c9crFieldSwrqShift = 16
	RegisterMdma_c9crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c9crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9crFieldSwrqMask)
	}
}

// registerMdma_c9tcrType This register is used to configure the concerned channel.
type registerMdma_c9tcrType uint32

const (
	RegisterMdma_c9tcrFieldSincShift = 0
	RegisterMdma_c9tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c9tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldSincMask) >> RegisterMdma_c9tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c9tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c9tcrFieldSincShift))
}

const (
	RegisterMdma_c9tcrFieldDincShift = 2
	RegisterMdma_c9tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c9tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldDincMask) >> RegisterMdma_c9tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c9tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c9tcrFieldDincShift))
}

const (
	RegisterMdma_c9tcrFieldSsizeShift = 4
	RegisterMdma_c9tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c9tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldSsizeMask) >> RegisterMdma_c9tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c9tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c9tcrFieldSsizeShift))
}

const (
	RegisterMdma_c9tcrFieldDsizeShift = 6
	RegisterMdma_c9tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c9tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldDsizeMask) >> RegisterMdma_c9tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c9tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c9tcrFieldDsizeShift))
}

const (
	RegisterMdma_c9tcrFieldSincosShift = 8
	RegisterMdma_c9tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c9tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldSincosMask) >> RegisterMdma_c9tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c9tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c9tcrFieldSincosShift))
}

const (
	RegisterMdma_c9tcrFieldDincosShift = 10
	RegisterMdma_c9tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c9tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldDincosMask) >> RegisterMdma_c9tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c9tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c9tcrFieldDincosShift))
}

const (
	RegisterMdma_c9tcrFieldSburstShift = 12
	RegisterMdma_c9tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c9tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldSburstMask) >> RegisterMdma_c9tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c9tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c9tcrFieldSburstShift))
}

const (
	RegisterMdma_c9tcrFieldDburstShift = 15
	RegisterMdma_c9tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c9tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldDburstMask) >> RegisterMdma_c9tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c9tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c9tcrFieldDburstShift))
}

const (
	RegisterMdma_c9tcrFieldTlenShift = 18
	RegisterMdma_c9tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c9tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldTlenMask) >> RegisterMdma_c9tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c9tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c9tcrFieldTlenShift))
}

const (
	RegisterMdma_c9tcrFieldPkeShift = 25
	RegisterMdma_c9tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c9tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c9tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c9tcrFieldPamShift = 26
	RegisterMdma_c9tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c9tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldPamMask) >> RegisterMdma_c9tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c9tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c9tcrFieldPamShift))
}

const (
	RegisterMdma_c9tcrFieldTrgmShift = 28
	RegisterMdma_c9tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c9tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldTrgmMask) >> RegisterMdma_c9tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c9tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c9tcrFieldTrgmShift))
}

const (
	RegisterMdma_c9tcrFieldSwrmShift = 30
	RegisterMdma_c9tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c9tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c9tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c9tcrFieldBwmShift = 31
	RegisterMdma_c9tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c9tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c9tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tcrFieldBwmMask)
	}
}

// registerMdma_c9bndtrType MDMA Channel x block number of data register
type registerMdma_c9bndtrType uint32

const (
	RegisterMdma_c9bndtrFieldBndtShift = 0
	RegisterMdma_c9bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c9bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9bndtrFieldBndtMask) >> RegisterMdma_c9bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c9bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c9bndtrFieldBndtShift))
}

const (
	RegisterMdma_c9bndtrFieldBrsumShift = 18
	RegisterMdma_c9bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c9bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c9bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c9bndtrFieldBrdumShift = 19
	RegisterMdma_c9bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c9bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c9bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c9bndtrFieldBrcShift = 20
	RegisterMdma_c9bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c9bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9bndtrFieldBrcMask) >> RegisterMdma_c9bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c9bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c9bndtrFieldBrcShift))
}

// registerMdma_c9sarType MDMA channel x source address register
type registerMdma_c9sarType uint32

const (
	RegisterMdma_c9sarFieldSarShift = 0
	RegisterMdma_c9sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c9sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9sarFieldSarMask) >> RegisterMdma_c9sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c9sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9sarFieldSarMask)|(uint32(value)<<RegisterMdma_c9sarFieldSarShift))
}

// registerMdma_c9darType MDMA channel x destination address register
type registerMdma_c9darType uint32

const (
	RegisterMdma_c9darFieldDarShift = 0
	RegisterMdma_c9darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c9darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9darFieldDarMask) >> RegisterMdma_c9darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c9darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9darFieldDarMask)|(uint32(value)<<RegisterMdma_c9darFieldDarShift))
}

// registerMdma_c9brurType MDMA channel x Block Repeat address Update register
type registerMdma_c9brurType uint32

const (
	RegisterMdma_c9brurFieldSuvShift = 0
	RegisterMdma_c9brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c9brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9brurFieldSuvMask) >> RegisterMdma_c9brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c9brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c9brurFieldSuvShift))
}

const (
	RegisterMdma_c9brurFieldDuvShift = 16
	RegisterMdma_c9brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c9brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9brurFieldDuvMask) >> RegisterMdma_c9brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c9brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c9brurFieldDuvShift))
}

// registerMdma_c9larType MDMA channel x Link Address register
type registerMdma_c9larType uint32

const (
	RegisterMdma_c9larFieldLarShift = 0
	RegisterMdma_c9larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c9larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9larFieldLarMask) >> RegisterMdma_c9larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c9larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9larFieldLarMask)|(uint32(value)<<RegisterMdma_c9larFieldLarShift))
}

// registerMdma_c9tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c9tbrType uint32

const (
	RegisterMdma_c9tbrFieldTselShift = 0
	RegisterMdma_c9tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c9tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tbrFieldTselMask) >> RegisterMdma_c9tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c9tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c9tbrFieldTselShift))
}

const (
	RegisterMdma_c9tbrFieldSbusShift = 16
	RegisterMdma_c9tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c9tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c9tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c9tbrFieldDbusShift = 17
	RegisterMdma_c9tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c9tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c9tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c9tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9tbrFieldDbusMask)
	}
}

// registerMdma_c9marType MDMA channel x Mask address register
type registerMdma_c9marType uint32

const (
	RegisterMdma_c9marFieldMarShift = 0
	RegisterMdma_c9marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c9marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9marFieldMarMask) >> RegisterMdma_c9marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c9marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9marFieldMarMask)|(uint32(value)<<RegisterMdma_c9marFieldMarShift))
}

// registerMdma_c9mdrType MDMA channel x Mask Data register
type registerMdma_c9mdrType uint32

const (
	RegisterMdma_c9mdrFieldMdrShift = 0
	RegisterMdma_c9mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c9mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c9mdrFieldMdrMask) >> RegisterMdma_c9mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c9mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c9mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c9mdrFieldMdrShift))
}

// registerMdma_c10isrType MDMA channel x interrupt/status register
type registerMdma_c10isrType uint32

const (
	RegisterMdma_c10isrFieldTeif10Shift = 0
	RegisterMdma_c10isrFieldTeif10Mask  = 0x1
)

// GetTeif10 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c10isrType) GetTeif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10isrFieldTeif10Mask) != 0
}

// SetTeif10 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c10isrType) SetTeif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10isrFieldTeif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10isrFieldTeif10Mask)
	}
}

const (
	RegisterMdma_c10isrFieldCtcif10Shift = 1
	RegisterMdma_c10isrFieldCtcif10Mask  = 0x2
)

// GetCtcif10 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c10isrType) GetCtcif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10isrFieldCtcif10Mask) != 0
}

// SetCtcif10 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c10isrType) SetCtcif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10isrFieldCtcif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10isrFieldCtcif10Mask)
	}
}

const (
	RegisterMdma_c10isrFieldBrtif10Shift = 2
	RegisterMdma_c10isrFieldBrtif10Mask  = 0x4
)

// GetBrtif10 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c10isrType) GetBrtif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10isrFieldBrtif10Mask) != 0
}

// SetBrtif10 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c10isrType) SetBrtif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10isrFieldBrtif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10isrFieldBrtif10Mask)
	}
}

const (
	RegisterMdma_c10isrFieldBtif10Shift = 3
	RegisterMdma_c10isrFieldBtif10Mask  = 0x8
)

// GetBtif10 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c10isrType) GetBtif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10isrFieldBtif10Mask) != 0
}

// SetBtif10 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c10isrType) SetBtif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10isrFieldBtif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10isrFieldBtif10Mask)
	}
}

const (
	RegisterMdma_c10isrFieldTcif10Shift = 4
	RegisterMdma_c10isrFieldTcif10Mask  = 0x10
)

// GetTcif10 channel x buffer transfer complete
func (r *registerMdma_c10isrType) GetTcif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10isrFieldTcif10Mask) != 0
}

// SetTcif10 channel x buffer transfer complete
func (r *registerMdma_c10isrType) SetTcif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10isrFieldTcif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10isrFieldTcif10Mask)
	}
}

const (
	RegisterMdma_c10isrFieldCrqa10Shift = 16
	RegisterMdma_c10isrFieldCrqa10Mask  = 0x10000
)

// GetCrqa10 channel x request active flag
func (r *registerMdma_c10isrType) GetCrqa10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10isrFieldCrqa10Mask) != 0
}

// SetCrqa10 channel x request active flag
func (r *registerMdma_c10isrType) SetCrqa10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10isrFieldCrqa10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10isrFieldCrqa10Mask)
	}
}

// registerMdma_c10ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c10ifcrType uint32

const (
	RegisterMdma_c10ifcrFieldCteif10Shift = 0
	RegisterMdma_c10ifcrFieldCteif10Mask  = 0x1
)

// GetCteif10 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c10ifcrType) GetCteif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10ifcrFieldCteif10Mask) != 0
}

// SetCteif10 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c10ifcrType) SetCteif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10ifcrFieldCteif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10ifcrFieldCteif10Mask)
	}
}

const (
	RegisterMdma_c10ifcrFieldCctcif10Shift = 1
	RegisterMdma_c10ifcrFieldCctcif10Mask  = 0x2
)

// GetCctcif10 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c10ifcrType) GetCctcif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10ifcrFieldCctcif10Mask) != 0
}

// SetCctcif10 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c10ifcrType) SetCctcif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10ifcrFieldCctcif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10ifcrFieldCctcif10Mask)
	}
}

const (
	RegisterMdma_c10ifcrFieldCbrtif10Shift = 2
	RegisterMdma_c10ifcrFieldCbrtif10Mask  = 0x4
)

// GetCbrtif10 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c10ifcrType) GetCbrtif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10ifcrFieldCbrtif10Mask) != 0
}

// SetCbrtif10 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c10ifcrType) SetCbrtif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10ifcrFieldCbrtif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10ifcrFieldCbrtif10Mask)
	}
}

const (
	RegisterMdma_c10ifcrFieldCbtif10Shift = 3
	RegisterMdma_c10ifcrFieldCbtif10Mask  = 0x8
)

// GetCbtif10 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c10ifcrType) GetCbtif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10ifcrFieldCbtif10Mask) != 0
}

// SetCbtif10 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c10ifcrType) SetCbtif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10ifcrFieldCbtif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10ifcrFieldCbtif10Mask)
	}
}

const (
	RegisterMdma_c10ifcrFieldCltcif10Shift = 4
	RegisterMdma_c10ifcrFieldCltcif10Mask  = 0x10
)

// GetCltcif10 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c10ifcrType) GetCltcif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10ifcrFieldCltcif10Mask) != 0
}

// SetCltcif10 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c10ifcrType) SetCltcif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10ifcrFieldCltcif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10ifcrFieldCltcif10Mask)
	}
}

// registerMdma_c10esrType MDMA Channel x error status register
type registerMdma_c10esrType uint32

const (
	RegisterMdma_c10esrFieldTeaShift = 0
	RegisterMdma_c10esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c10esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10esrFieldTeaMask) >> RegisterMdma_c10esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c10esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c10esrFieldTeaShift))
}

const (
	RegisterMdma_c10esrFieldTedShift = 7
	RegisterMdma_c10esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c10esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c10esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10esrFieldTedMask)
	}
}

const (
	RegisterMdma_c10esrFieldTeldShift = 8
	RegisterMdma_c10esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c10esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c10esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c10esrFieldTemdShift = 9
	RegisterMdma_c10esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c10esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c10esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c10esrFieldAseShift = 10
	RegisterMdma_c10esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c10esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c10esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10esrFieldAseMask)
	}
}

const (
	RegisterMdma_c10esrFieldBseShift = 11
	RegisterMdma_c10esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c10esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c10esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10esrFieldBseMask)
	}
}

// registerMdma_c10crType This register is used to control the concerned channel.
type registerMdma_c10crType uint32

const (
	RegisterMdma_c10crFieldEnShift = 0
	RegisterMdma_c10crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c10crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c10crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10crFieldEnMask)
	}
}

const (
	RegisterMdma_c10crFieldTeieShift = 1
	RegisterMdma_c10crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c10crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c10crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10crFieldTeieMask)
	}
}

const (
	RegisterMdma_c10crFieldCtcieShift = 2
	RegisterMdma_c10crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c10crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c10crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c10crFieldBrtieShift = 3
	RegisterMdma_c10crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c10crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c10crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c10crFieldBtieShift = 4
	RegisterMdma_c10crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c10crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c10crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10crFieldBtieMask)
	}
}

const (
	RegisterMdma_c10crFieldTcieShift = 5
	RegisterMdma_c10crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c10crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c10crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10crFieldTcieMask)
	}
}

const (
	RegisterMdma_c10crFieldPlShift = 6
	RegisterMdma_c10crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c10crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10crFieldPlMask) >> RegisterMdma_c10crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c10crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10crFieldPlMask)|(uint32(value)<<RegisterMdma_c10crFieldPlShift))
}

const (
	RegisterMdma_c10crFieldBexShift = 12
	RegisterMdma_c10crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c10crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c10crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10crFieldBexMask)
	}
}

const (
	RegisterMdma_c10crFieldHexShift = 13
	RegisterMdma_c10crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c10crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c10crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10crFieldHexMask)
	}
}

const (
	RegisterMdma_c10crFieldWexShift = 14
	RegisterMdma_c10crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c10crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c10crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10crFieldWexMask)
	}
}

const (
	RegisterMdma_c10crFieldSwrqShift = 16
	RegisterMdma_c10crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c10crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10crFieldSwrqMask)
	}
}

// registerMdma_c10tcrType This register is used to configure the concerned channel.
type registerMdma_c10tcrType uint32

const (
	RegisterMdma_c10tcrFieldSincShift = 0
	RegisterMdma_c10tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c10tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldSincMask) >> RegisterMdma_c10tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c10tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c10tcrFieldSincShift))
}

const (
	RegisterMdma_c10tcrFieldDincShift = 2
	RegisterMdma_c10tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c10tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldDincMask) >> RegisterMdma_c10tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c10tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c10tcrFieldDincShift))
}

const (
	RegisterMdma_c10tcrFieldSsizeShift = 4
	RegisterMdma_c10tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c10tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldSsizeMask) >> RegisterMdma_c10tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c10tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c10tcrFieldSsizeShift))
}

const (
	RegisterMdma_c10tcrFieldDsizeShift = 6
	RegisterMdma_c10tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c10tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldDsizeMask) >> RegisterMdma_c10tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c10tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c10tcrFieldDsizeShift))
}

const (
	RegisterMdma_c10tcrFieldSincosShift = 8
	RegisterMdma_c10tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c10tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldSincosMask) >> RegisterMdma_c10tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c10tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c10tcrFieldSincosShift))
}

const (
	RegisterMdma_c10tcrFieldDincosShift = 10
	RegisterMdma_c10tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c10tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldDincosMask) >> RegisterMdma_c10tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c10tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c10tcrFieldDincosShift))
}

const (
	RegisterMdma_c10tcrFieldSburstShift = 12
	RegisterMdma_c10tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c10tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldSburstMask) >> RegisterMdma_c10tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c10tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c10tcrFieldSburstShift))
}

const (
	RegisterMdma_c10tcrFieldDburstShift = 15
	RegisterMdma_c10tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c10tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldDburstMask) >> RegisterMdma_c10tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c10tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c10tcrFieldDburstShift))
}

const (
	RegisterMdma_c10tcrFieldTlenShift = 18
	RegisterMdma_c10tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c10tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldTlenMask) >> RegisterMdma_c10tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c10tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c10tcrFieldTlenShift))
}

const (
	RegisterMdma_c10tcrFieldPkeShift = 25
	RegisterMdma_c10tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c10tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c10tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c10tcrFieldPamShift = 26
	RegisterMdma_c10tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c10tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldPamMask) >> RegisterMdma_c10tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c10tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c10tcrFieldPamShift))
}

const (
	RegisterMdma_c10tcrFieldTrgmShift = 28
	RegisterMdma_c10tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c10tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldTrgmMask) >> RegisterMdma_c10tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c10tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c10tcrFieldTrgmShift))
}

const (
	RegisterMdma_c10tcrFieldSwrmShift = 30
	RegisterMdma_c10tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c10tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c10tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c10tcrFieldBwmShift = 31
	RegisterMdma_c10tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c10tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c10tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tcrFieldBwmMask)
	}
}

// registerMdma_c10bndtrType MDMA Channel x block number of data register
type registerMdma_c10bndtrType uint32

const (
	RegisterMdma_c10bndtrFieldBndtShift = 0
	RegisterMdma_c10bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c10bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10bndtrFieldBndtMask) >> RegisterMdma_c10bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c10bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c10bndtrFieldBndtShift))
}

const (
	RegisterMdma_c10bndtrFieldBrsumShift = 18
	RegisterMdma_c10bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c10bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c10bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c10bndtrFieldBrdumShift = 19
	RegisterMdma_c10bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c10bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c10bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c10bndtrFieldBrcShift = 20
	RegisterMdma_c10bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c10bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10bndtrFieldBrcMask) >> RegisterMdma_c10bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c10bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c10bndtrFieldBrcShift))
}

// registerMdma_c10sarType MDMA channel x source address register
type registerMdma_c10sarType uint32

const (
	RegisterMdma_c10sarFieldSarShift = 0
	RegisterMdma_c10sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c10sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10sarFieldSarMask) >> RegisterMdma_c10sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c10sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10sarFieldSarMask)|(uint32(value)<<RegisterMdma_c10sarFieldSarShift))
}

// registerMdma_c10darType MDMA channel x destination address register
type registerMdma_c10darType uint32

const (
	RegisterMdma_c10darFieldDarShift = 0
	RegisterMdma_c10darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c10darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10darFieldDarMask) >> RegisterMdma_c10darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c10darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10darFieldDarMask)|(uint32(value)<<RegisterMdma_c10darFieldDarShift))
}

// registerMdma_c10brurType MDMA channel x Block Repeat address Update register
type registerMdma_c10brurType uint32

const (
	RegisterMdma_c10brurFieldSuvShift = 0
	RegisterMdma_c10brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c10brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10brurFieldSuvMask) >> RegisterMdma_c10brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c10brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c10brurFieldSuvShift))
}

const (
	RegisterMdma_c10brurFieldDuvShift = 16
	RegisterMdma_c10brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c10brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10brurFieldDuvMask) >> RegisterMdma_c10brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c10brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c10brurFieldDuvShift))
}

// registerMdma_c10larType MDMA channel x Link Address register
type registerMdma_c10larType uint32

const (
	RegisterMdma_c10larFieldLarShift = 0
	RegisterMdma_c10larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c10larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10larFieldLarMask) >> RegisterMdma_c10larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c10larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10larFieldLarMask)|(uint32(value)<<RegisterMdma_c10larFieldLarShift))
}

// registerMdma_c10tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c10tbrType uint32

const (
	RegisterMdma_c10tbrFieldTselShift = 0
	RegisterMdma_c10tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c10tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tbrFieldTselMask) >> RegisterMdma_c10tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c10tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c10tbrFieldTselShift))
}

const (
	RegisterMdma_c10tbrFieldSbusShift = 16
	RegisterMdma_c10tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c10tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c10tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c10tbrFieldDbusShift = 17
	RegisterMdma_c10tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c10tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c10tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c10tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10tbrFieldDbusMask)
	}
}

// registerMdma_c10marType MDMA channel x Mask address register
type registerMdma_c10marType uint32

const (
	RegisterMdma_c10marFieldMarShift = 0
	RegisterMdma_c10marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c10marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10marFieldMarMask) >> RegisterMdma_c10marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c10marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10marFieldMarMask)|(uint32(value)<<RegisterMdma_c10marFieldMarShift))
}

// registerMdma_c10mdrType MDMA channel x Mask Data register
type registerMdma_c10mdrType uint32

const (
	RegisterMdma_c10mdrFieldMdrShift = 0
	RegisterMdma_c10mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c10mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c10mdrFieldMdrMask) >> RegisterMdma_c10mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c10mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c10mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c10mdrFieldMdrShift))
}

// registerMdma_c11isrType MDMA channel x interrupt/status register
type registerMdma_c11isrType uint32

const (
	RegisterMdma_c11isrFieldTeif11Shift = 0
	RegisterMdma_c11isrFieldTeif11Mask  = 0x1
)

// GetTeif11 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c11isrType) GetTeif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11isrFieldTeif11Mask) != 0
}

// SetTeif11 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c11isrType) SetTeif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11isrFieldTeif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11isrFieldTeif11Mask)
	}
}

const (
	RegisterMdma_c11isrFieldCtcif11Shift = 1
	RegisterMdma_c11isrFieldCtcif11Mask  = 0x2
)

// GetCtcif11 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c11isrType) GetCtcif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11isrFieldCtcif11Mask) != 0
}

// SetCtcif11 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c11isrType) SetCtcif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11isrFieldCtcif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11isrFieldCtcif11Mask)
	}
}

const (
	RegisterMdma_c11isrFieldBrtif11Shift = 2
	RegisterMdma_c11isrFieldBrtif11Mask  = 0x4
)

// GetBrtif11 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c11isrType) GetBrtif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11isrFieldBrtif11Mask) != 0
}

// SetBrtif11 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c11isrType) SetBrtif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11isrFieldBrtif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11isrFieldBrtif11Mask)
	}
}

const (
	RegisterMdma_c11isrFieldBtif11Shift = 3
	RegisterMdma_c11isrFieldBtif11Mask  = 0x8
)

// GetBtif11 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c11isrType) GetBtif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11isrFieldBtif11Mask) != 0
}

// SetBtif11 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c11isrType) SetBtif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11isrFieldBtif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11isrFieldBtif11Mask)
	}
}

const (
	RegisterMdma_c11isrFieldTcif11Shift = 4
	RegisterMdma_c11isrFieldTcif11Mask  = 0x10
)

// GetTcif11 channel x buffer transfer complete
func (r *registerMdma_c11isrType) GetTcif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11isrFieldTcif11Mask) != 0
}

// SetTcif11 channel x buffer transfer complete
func (r *registerMdma_c11isrType) SetTcif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11isrFieldTcif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11isrFieldTcif11Mask)
	}
}

const (
	RegisterMdma_c11isrFieldCrqa11Shift = 16
	RegisterMdma_c11isrFieldCrqa11Mask  = 0x10000
)

// GetCrqa11 channel x request active flag
func (r *registerMdma_c11isrType) GetCrqa11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11isrFieldCrqa11Mask) != 0
}

// SetCrqa11 channel x request active flag
func (r *registerMdma_c11isrType) SetCrqa11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11isrFieldCrqa11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11isrFieldCrqa11Mask)
	}
}

// registerMdma_c11ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c11ifcrType uint32

const (
	RegisterMdma_c11ifcrFieldCteif11Shift = 0
	RegisterMdma_c11ifcrFieldCteif11Mask  = 0x1
)

// GetCteif11 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c11ifcrType) GetCteif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11ifcrFieldCteif11Mask) != 0
}

// SetCteif11 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c11ifcrType) SetCteif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11ifcrFieldCteif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11ifcrFieldCteif11Mask)
	}
}

const (
	RegisterMdma_c11ifcrFieldCctcif11Shift = 1
	RegisterMdma_c11ifcrFieldCctcif11Mask  = 0x2
)

// GetCctcif11 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c11ifcrType) GetCctcif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11ifcrFieldCctcif11Mask) != 0
}

// SetCctcif11 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c11ifcrType) SetCctcif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11ifcrFieldCctcif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11ifcrFieldCctcif11Mask)
	}
}

const (
	RegisterMdma_c11ifcrFieldCbrtif11Shift = 2
	RegisterMdma_c11ifcrFieldCbrtif11Mask  = 0x4
)

// GetCbrtif11 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c11ifcrType) GetCbrtif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11ifcrFieldCbrtif11Mask) != 0
}

// SetCbrtif11 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c11ifcrType) SetCbrtif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11ifcrFieldCbrtif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11ifcrFieldCbrtif11Mask)
	}
}

const (
	RegisterMdma_c11ifcrFieldCbtif11Shift = 3
	RegisterMdma_c11ifcrFieldCbtif11Mask  = 0x8
)

// GetCbtif11 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c11ifcrType) GetCbtif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11ifcrFieldCbtif11Mask) != 0
}

// SetCbtif11 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c11ifcrType) SetCbtif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11ifcrFieldCbtif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11ifcrFieldCbtif11Mask)
	}
}

const (
	RegisterMdma_c11ifcrFieldCltcif11Shift = 4
	RegisterMdma_c11ifcrFieldCltcif11Mask  = 0x10
)

// GetCltcif11 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c11ifcrType) GetCltcif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11ifcrFieldCltcif11Mask) != 0
}

// SetCltcif11 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c11ifcrType) SetCltcif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11ifcrFieldCltcif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11ifcrFieldCltcif11Mask)
	}
}

// registerMdma_c11esrType MDMA Channel x error status register
type registerMdma_c11esrType uint32

const (
	RegisterMdma_c11esrFieldTeaShift = 0
	RegisterMdma_c11esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c11esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11esrFieldTeaMask) >> RegisterMdma_c11esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c11esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c11esrFieldTeaShift))
}

const (
	RegisterMdma_c11esrFieldTedShift = 7
	RegisterMdma_c11esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c11esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c11esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11esrFieldTedMask)
	}
}

const (
	RegisterMdma_c11esrFieldTeldShift = 8
	RegisterMdma_c11esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c11esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c11esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c11esrFieldTemdShift = 9
	RegisterMdma_c11esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c11esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c11esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c11esrFieldAseShift = 10
	RegisterMdma_c11esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c11esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c11esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11esrFieldAseMask)
	}
}

const (
	RegisterMdma_c11esrFieldBseShift = 11
	RegisterMdma_c11esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c11esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c11esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11esrFieldBseMask)
	}
}

// registerMdma_c11crType This register is used to control the concerned channel.
type registerMdma_c11crType uint32

const (
	RegisterMdma_c11crFieldEnShift = 0
	RegisterMdma_c11crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c11crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c11crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11crFieldEnMask)
	}
}

const (
	RegisterMdma_c11crFieldTeieShift = 1
	RegisterMdma_c11crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c11crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c11crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11crFieldTeieMask)
	}
}

const (
	RegisterMdma_c11crFieldCtcieShift = 2
	RegisterMdma_c11crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c11crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c11crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c11crFieldBrtieShift = 3
	RegisterMdma_c11crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c11crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c11crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c11crFieldBtieShift = 4
	RegisterMdma_c11crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c11crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c11crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11crFieldBtieMask)
	}
}

const (
	RegisterMdma_c11crFieldTcieShift = 5
	RegisterMdma_c11crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c11crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c11crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11crFieldTcieMask)
	}
}

const (
	RegisterMdma_c11crFieldPlShift = 6
	RegisterMdma_c11crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c11crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11crFieldPlMask) >> RegisterMdma_c11crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c11crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11crFieldPlMask)|(uint32(value)<<RegisterMdma_c11crFieldPlShift))
}

const (
	RegisterMdma_c11crFieldBexShift = 12
	RegisterMdma_c11crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c11crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c11crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11crFieldBexMask)
	}
}

const (
	RegisterMdma_c11crFieldHexShift = 13
	RegisterMdma_c11crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c11crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c11crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11crFieldHexMask)
	}
}

const (
	RegisterMdma_c11crFieldWexShift = 14
	RegisterMdma_c11crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c11crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c11crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11crFieldWexMask)
	}
}

const (
	RegisterMdma_c11crFieldSwrqShift = 16
	RegisterMdma_c11crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c11crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11crFieldSwrqMask)
	}
}

// registerMdma_c11tcrType This register is used to configure the concerned channel.
type registerMdma_c11tcrType uint32

const (
	RegisterMdma_c11tcrFieldSincShift = 0
	RegisterMdma_c11tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c11tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldSincMask) >> RegisterMdma_c11tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c11tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c11tcrFieldSincShift))
}

const (
	RegisterMdma_c11tcrFieldDincShift = 2
	RegisterMdma_c11tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c11tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldDincMask) >> RegisterMdma_c11tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c11tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c11tcrFieldDincShift))
}

const (
	RegisterMdma_c11tcrFieldSsizeShift = 4
	RegisterMdma_c11tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c11tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldSsizeMask) >> RegisterMdma_c11tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c11tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c11tcrFieldSsizeShift))
}

const (
	RegisterMdma_c11tcrFieldDsizeShift = 6
	RegisterMdma_c11tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c11tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldDsizeMask) >> RegisterMdma_c11tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c11tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c11tcrFieldDsizeShift))
}

const (
	RegisterMdma_c11tcrFieldSincosShift = 8
	RegisterMdma_c11tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c11tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldSincosMask) >> RegisterMdma_c11tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c11tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c11tcrFieldSincosShift))
}

const (
	RegisterMdma_c11tcrFieldDincosShift = 10
	RegisterMdma_c11tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c11tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldDincosMask) >> RegisterMdma_c11tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c11tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c11tcrFieldDincosShift))
}

const (
	RegisterMdma_c11tcrFieldSburstShift = 12
	RegisterMdma_c11tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c11tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldSburstMask) >> RegisterMdma_c11tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c11tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c11tcrFieldSburstShift))
}

const (
	RegisterMdma_c11tcrFieldDburstShift = 15
	RegisterMdma_c11tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c11tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldDburstMask) >> RegisterMdma_c11tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c11tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c11tcrFieldDburstShift))
}

const (
	RegisterMdma_c11tcrFieldTlenShift = 18
	RegisterMdma_c11tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c11tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldTlenMask) >> RegisterMdma_c11tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c11tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c11tcrFieldTlenShift))
}

const (
	RegisterMdma_c11tcrFieldPkeShift = 25
	RegisterMdma_c11tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c11tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c11tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c11tcrFieldPamShift = 26
	RegisterMdma_c11tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c11tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldPamMask) >> RegisterMdma_c11tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c11tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c11tcrFieldPamShift))
}

const (
	RegisterMdma_c11tcrFieldTrgmShift = 28
	RegisterMdma_c11tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c11tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldTrgmMask) >> RegisterMdma_c11tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c11tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c11tcrFieldTrgmShift))
}

const (
	RegisterMdma_c11tcrFieldSwrmShift = 30
	RegisterMdma_c11tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c11tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c11tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c11tcrFieldBwmShift = 31
	RegisterMdma_c11tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c11tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c11tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tcrFieldBwmMask)
	}
}

// registerMdma_c11bndtrType MDMA Channel x block number of data register
type registerMdma_c11bndtrType uint32

const (
	RegisterMdma_c11bndtrFieldBndtShift = 0
	RegisterMdma_c11bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c11bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11bndtrFieldBndtMask) >> RegisterMdma_c11bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c11bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c11bndtrFieldBndtShift))
}

const (
	RegisterMdma_c11bndtrFieldBrsumShift = 18
	RegisterMdma_c11bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c11bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c11bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c11bndtrFieldBrdumShift = 19
	RegisterMdma_c11bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c11bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c11bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c11bndtrFieldBrcShift = 20
	RegisterMdma_c11bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c11bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11bndtrFieldBrcMask) >> RegisterMdma_c11bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c11bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c11bndtrFieldBrcShift))
}

// registerMdma_c11sarType MDMA channel x source address register
type registerMdma_c11sarType uint32

const (
	RegisterMdma_c11sarFieldSarShift = 0
	RegisterMdma_c11sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c11sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11sarFieldSarMask) >> RegisterMdma_c11sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c11sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11sarFieldSarMask)|(uint32(value)<<RegisterMdma_c11sarFieldSarShift))
}

// registerMdma_c11darType MDMA channel x destination address register
type registerMdma_c11darType uint32

const (
	RegisterMdma_c11darFieldDarShift = 0
	RegisterMdma_c11darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c11darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11darFieldDarMask) >> RegisterMdma_c11darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c11darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11darFieldDarMask)|(uint32(value)<<RegisterMdma_c11darFieldDarShift))
}

// registerMdma_c11brurType MDMA channel x Block Repeat address Update register
type registerMdma_c11brurType uint32

const (
	RegisterMdma_c11brurFieldSuvShift = 0
	RegisterMdma_c11brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c11brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11brurFieldSuvMask) >> RegisterMdma_c11brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c11brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c11brurFieldSuvShift))
}

const (
	RegisterMdma_c11brurFieldDuvShift = 16
	RegisterMdma_c11brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c11brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11brurFieldDuvMask) >> RegisterMdma_c11brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c11brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c11brurFieldDuvShift))
}

// registerMdma_c11larType MDMA channel x Link Address register
type registerMdma_c11larType uint32

const (
	RegisterMdma_c11larFieldLarShift = 0
	RegisterMdma_c11larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c11larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11larFieldLarMask) >> RegisterMdma_c11larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c11larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11larFieldLarMask)|(uint32(value)<<RegisterMdma_c11larFieldLarShift))
}

// registerMdma_c11tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c11tbrType uint32

const (
	RegisterMdma_c11tbrFieldTselShift = 0
	RegisterMdma_c11tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c11tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tbrFieldTselMask) >> RegisterMdma_c11tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c11tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c11tbrFieldTselShift))
}

const (
	RegisterMdma_c11tbrFieldSbusShift = 16
	RegisterMdma_c11tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c11tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c11tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c11tbrFieldDbusShift = 17
	RegisterMdma_c11tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c11tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c11tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c11tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11tbrFieldDbusMask)
	}
}

// registerMdma_c11marType MDMA channel x Mask address register
type registerMdma_c11marType uint32

const (
	RegisterMdma_c11marFieldMarShift = 0
	RegisterMdma_c11marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c11marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11marFieldMarMask) >> RegisterMdma_c11marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c11marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11marFieldMarMask)|(uint32(value)<<RegisterMdma_c11marFieldMarShift))
}

// registerMdma_c11mdrType MDMA channel x Mask Data register
type registerMdma_c11mdrType uint32

const (
	RegisterMdma_c11mdrFieldMdrShift = 0
	RegisterMdma_c11mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c11mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c11mdrFieldMdrMask) >> RegisterMdma_c11mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c11mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c11mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c11mdrFieldMdrShift))
}

// registerMdma_c12isrType MDMA channel x interrupt/status register
type registerMdma_c12isrType uint32

const (
	RegisterMdma_c12isrFieldTeif12Shift = 0
	RegisterMdma_c12isrFieldTeif12Mask  = 0x1
)

// GetTeif12 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c12isrType) GetTeif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12isrFieldTeif12Mask) != 0
}

// SetTeif12 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c12isrType) SetTeif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12isrFieldTeif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12isrFieldTeif12Mask)
	}
}

const (
	RegisterMdma_c12isrFieldCtcif12Shift = 1
	RegisterMdma_c12isrFieldCtcif12Mask  = 0x2
)

// GetCtcif12 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c12isrType) GetCtcif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12isrFieldCtcif12Mask) != 0
}

// SetCtcif12 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c12isrType) SetCtcif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12isrFieldCtcif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12isrFieldCtcif12Mask)
	}
}

const (
	RegisterMdma_c12isrFieldBrtif12Shift = 2
	RegisterMdma_c12isrFieldBrtif12Mask  = 0x4
)

// GetBrtif12 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c12isrType) GetBrtif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12isrFieldBrtif12Mask) != 0
}

// SetBrtif12 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c12isrType) SetBrtif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12isrFieldBrtif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12isrFieldBrtif12Mask)
	}
}

const (
	RegisterMdma_c12isrFieldBtif12Shift = 3
	RegisterMdma_c12isrFieldBtif12Mask  = 0x8
)

// GetBtif12 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c12isrType) GetBtif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12isrFieldBtif12Mask) != 0
}

// SetBtif12 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c12isrType) SetBtif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12isrFieldBtif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12isrFieldBtif12Mask)
	}
}

const (
	RegisterMdma_c12isrFieldTcif12Shift = 4
	RegisterMdma_c12isrFieldTcif12Mask  = 0x10
)

// GetTcif12 channel x buffer transfer complete
func (r *registerMdma_c12isrType) GetTcif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12isrFieldTcif12Mask) != 0
}

// SetTcif12 channel x buffer transfer complete
func (r *registerMdma_c12isrType) SetTcif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12isrFieldTcif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12isrFieldTcif12Mask)
	}
}

const (
	RegisterMdma_c12isrFieldCrqa12Shift = 16
	RegisterMdma_c12isrFieldCrqa12Mask  = 0x10000
)

// GetCrqa12 channel x request active flag
func (r *registerMdma_c12isrType) GetCrqa12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12isrFieldCrqa12Mask) != 0
}

// SetCrqa12 channel x request active flag
func (r *registerMdma_c12isrType) SetCrqa12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12isrFieldCrqa12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12isrFieldCrqa12Mask)
	}
}

// registerMdma_c12ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c12ifcrType uint32

const (
	RegisterMdma_c12ifcrFieldCteif12Shift = 0
	RegisterMdma_c12ifcrFieldCteif12Mask  = 0x1
)

// GetCteif12 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c12ifcrType) GetCteif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12ifcrFieldCteif12Mask) != 0
}

// SetCteif12 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c12ifcrType) SetCteif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12ifcrFieldCteif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12ifcrFieldCteif12Mask)
	}
}

const (
	RegisterMdma_c12ifcrFieldCctcif12Shift = 1
	RegisterMdma_c12ifcrFieldCctcif12Mask  = 0x2
)

// GetCctcif12 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c12ifcrType) GetCctcif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12ifcrFieldCctcif12Mask) != 0
}

// SetCctcif12 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c12ifcrType) SetCctcif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12ifcrFieldCctcif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12ifcrFieldCctcif12Mask)
	}
}

const (
	RegisterMdma_c12ifcrFieldCbrtif12Shift = 2
	RegisterMdma_c12ifcrFieldCbrtif12Mask  = 0x4
)

// GetCbrtif12 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c12ifcrType) GetCbrtif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12ifcrFieldCbrtif12Mask) != 0
}

// SetCbrtif12 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c12ifcrType) SetCbrtif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12ifcrFieldCbrtif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12ifcrFieldCbrtif12Mask)
	}
}

const (
	RegisterMdma_c12ifcrFieldCbtif12Shift = 3
	RegisterMdma_c12ifcrFieldCbtif12Mask  = 0x8
)

// GetCbtif12 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c12ifcrType) GetCbtif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12ifcrFieldCbtif12Mask) != 0
}

// SetCbtif12 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c12ifcrType) SetCbtif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12ifcrFieldCbtif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12ifcrFieldCbtif12Mask)
	}
}

const (
	RegisterMdma_c12ifcrFieldCltcif12Shift = 4
	RegisterMdma_c12ifcrFieldCltcif12Mask  = 0x10
)

// GetCltcif12 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c12ifcrType) GetCltcif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12ifcrFieldCltcif12Mask) != 0
}

// SetCltcif12 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c12ifcrType) SetCltcif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12ifcrFieldCltcif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12ifcrFieldCltcif12Mask)
	}
}

// registerMdma_c12esrType MDMA Channel x error status register
type registerMdma_c12esrType uint32

const (
	RegisterMdma_c12esrFieldTeaShift = 0
	RegisterMdma_c12esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c12esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12esrFieldTeaMask) >> RegisterMdma_c12esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c12esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c12esrFieldTeaShift))
}

const (
	RegisterMdma_c12esrFieldTedShift = 7
	RegisterMdma_c12esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c12esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c12esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12esrFieldTedMask)
	}
}

const (
	RegisterMdma_c12esrFieldTeldShift = 8
	RegisterMdma_c12esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c12esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c12esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c12esrFieldTemdShift = 9
	RegisterMdma_c12esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c12esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c12esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c12esrFieldAseShift = 10
	RegisterMdma_c12esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c12esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c12esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12esrFieldAseMask)
	}
}

const (
	RegisterMdma_c12esrFieldBseShift = 11
	RegisterMdma_c12esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c12esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c12esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12esrFieldBseMask)
	}
}

// registerMdma_c12crType This register is used to control the concerned channel.
type registerMdma_c12crType uint32

const (
	RegisterMdma_c12crFieldEnShift = 0
	RegisterMdma_c12crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c12crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c12crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12crFieldEnMask)
	}
}

const (
	RegisterMdma_c12crFieldTeieShift = 1
	RegisterMdma_c12crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c12crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c12crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12crFieldTeieMask)
	}
}

const (
	RegisterMdma_c12crFieldCtcieShift = 2
	RegisterMdma_c12crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c12crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c12crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c12crFieldBrtieShift = 3
	RegisterMdma_c12crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c12crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c12crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c12crFieldBtieShift = 4
	RegisterMdma_c12crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c12crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c12crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12crFieldBtieMask)
	}
}

const (
	RegisterMdma_c12crFieldTcieShift = 5
	RegisterMdma_c12crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c12crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c12crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12crFieldTcieMask)
	}
}

const (
	RegisterMdma_c12crFieldPlShift = 6
	RegisterMdma_c12crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c12crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12crFieldPlMask) >> RegisterMdma_c12crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c12crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12crFieldPlMask)|(uint32(value)<<RegisterMdma_c12crFieldPlShift))
}

const (
	RegisterMdma_c12crFieldBexShift = 12
	RegisterMdma_c12crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c12crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c12crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12crFieldBexMask)
	}
}

const (
	RegisterMdma_c12crFieldHexShift = 13
	RegisterMdma_c12crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c12crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c12crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12crFieldHexMask)
	}
}

const (
	RegisterMdma_c12crFieldWexShift = 14
	RegisterMdma_c12crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c12crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c12crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12crFieldWexMask)
	}
}

const (
	RegisterMdma_c12crFieldSwrqShift = 16
	RegisterMdma_c12crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c12crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12crFieldSwrqMask)
	}
}

// registerMdma_c12tcrType This register is used to configure the concerned channel.
type registerMdma_c12tcrType uint32

const (
	RegisterMdma_c12tcrFieldSincShift = 0
	RegisterMdma_c12tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c12tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldSincMask) >> RegisterMdma_c12tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c12tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c12tcrFieldSincShift))
}

const (
	RegisterMdma_c12tcrFieldDincShift = 2
	RegisterMdma_c12tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c12tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldDincMask) >> RegisterMdma_c12tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c12tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c12tcrFieldDincShift))
}

const (
	RegisterMdma_c12tcrFieldSsizeShift = 4
	RegisterMdma_c12tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c12tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldSsizeMask) >> RegisterMdma_c12tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c12tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c12tcrFieldSsizeShift))
}

const (
	RegisterMdma_c12tcrFieldDsizeShift = 6
	RegisterMdma_c12tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c12tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldDsizeMask) >> RegisterMdma_c12tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c12tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c12tcrFieldDsizeShift))
}

const (
	RegisterMdma_c12tcrFieldSincosShift = 8
	RegisterMdma_c12tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c12tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldSincosMask) >> RegisterMdma_c12tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c12tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c12tcrFieldSincosShift))
}

const (
	RegisterMdma_c12tcrFieldDincosShift = 10
	RegisterMdma_c12tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c12tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldDincosMask) >> RegisterMdma_c12tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c12tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c12tcrFieldDincosShift))
}

const (
	RegisterMdma_c12tcrFieldSburstShift = 12
	RegisterMdma_c12tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c12tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldSburstMask) >> RegisterMdma_c12tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c12tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c12tcrFieldSburstShift))
}

const (
	RegisterMdma_c12tcrFieldDburstShift = 15
	RegisterMdma_c12tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c12tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldDburstMask) >> RegisterMdma_c12tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c12tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c12tcrFieldDburstShift))
}

const (
	RegisterMdma_c12tcrFieldTlenShift = 18
	RegisterMdma_c12tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c12tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldTlenMask) >> RegisterMdma_c12tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c12tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c12tcrFieldTlenShift))
}

const (
	RegisterMdma_c12tcrFieldPkeShift = 25
	RegisterMdma_c12tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c12tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c12tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c12tcrFieldPamShift = 26
	RegisterMdma_c12tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c12tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldPamMask) >> RegisterMdma_c12tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c12tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c12tcrFieldPamShift))
}

const (
	RegisterMdma_c12tcrFieldTrgmShift = 28
	RegisterMdma_c12tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c12tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldTrgmMask) >> RegisterMdma_c12tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c12tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c12tcrFieldTrgmShift))
}

const (
	RegisterMdma_c12tcrFieldSwrmShift = 30
	RegisterMdma_c12tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c12tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c12tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c12tcrFieldBwmShift = 31
	RegisterMdma_c12tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c12tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c12tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tcrFieldBwmMask)
	}
}

// registerMdma_c12bndtrType MDMA Channel x block number of data register
type registerMdma_c12bndtrType uint32

const (
	RegisterMdma_c12bndtrFieldBndtShift = 0
	RegisterMdma_c12bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c12bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12bndtrFieldBndtMask) >> RegisterMdma_c12bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c12bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c12bndtrFieldBndtShift))
}

const (
	RegisterMdma_c12bndtrFieldBrsumShift = 18
	RegisterMdma_c12bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c12bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c12bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c12bndtrFieldBrdumShift = 19
	RegisterMdma_c12bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c12bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c12bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c12bndtrFieldBrcShift = 20
	RegisterMdma_c12bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c12bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12bndtrFieldBrcMask) >> RegisterMdma_c12bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c12bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c12bndtrFieldBrcShift))
}

// registerMdma_c12sarType MDMA channel x source address register
type registerMdma_c12sarType uint32

const (
	RegisterMdma_c12sarFieldSarShift = 0
	RegisterMdma_c12sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c12sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12sarFieldSarMask) >> RegisterMdma_c12sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c12sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12sarFieldSarMask)|(uint32(value)<<RegisterMdma_c12sarFieldSarShift))
}

// registerMdma_c12darType MDMA channel x destination address register
type registerMdma_c12darType uint32

const (
	RegisterMdma_c12darFieldDarShift = 0
	RegisterMdma_c12darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c12darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12darFieldDarMask) >> RegisterMdma_c12darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c12darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12darFieldDarMask)|(uint32(value)<<RegisterMdma_c12darFieldDarShift))
}

// registerMdma_c12brurType MDMA channel x Block Repeat address Update register
type registerMdma_c12brurType uint32

const (
	RegisterMdma_c12brurFieldSuvShift = 0
	RegisterMdma_c12brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c12brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12brurFieldSuvMask) >> RegisterMdma_c12brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c12brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c12brurFieldSuvShift))
}

const (
	RegisterMdma_c12brurFieldDuvShift = 16
	RegisterMdma_c12brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c12brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12brurFieldDuvMask) >> RegisterMdma_c12brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c12brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c12brurFieldDuvShift))
}

// registerMdma_c12larType MDMA channel x Link Address register
type registerMdma_c12larType uint32

const (
	RegisterMdma_c12larFieldLarShift = 0
	RegisterMdma_c12larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c12larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12larFieldLarMask) >> RegisterMdma_c12larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c12larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12larFieldLarMask)|(uint32(value)<<RegisterMdma_c12larFieldLarShift))
}

// registerMdma_c12tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c12tbrType uint32

const (
	RegisterMdma_c12tbrFieldTselShift = 0
	RegisterMdma_c12tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c12tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tbrFieldTselMask) >> RegisterMdma_c12tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c12tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c12tbrFieldTselShift))
}

const (
	RegisterMdma_c12tbrFieldSbusShift = 16
	RegisterMdma_c12tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c12tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c12tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c12tbrFieldDbusShift = 17
	RegisterMdma_c12tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c12tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c12tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c12tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12tbrFieldDbusMask)
	}
}

// registerMdma_c12marType MDMA channel x Mask address register
type registerMdma_c12marType uint32

const (
	RegisterMdma_c12marFieldMarShift = 0
	RegisterMdma_c12marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c12marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12marFieldMarMask) >> RegisterMdma_c12marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c12marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12marFieldMarMask)|(uint32(value)<<RegisterMdma_c12marFieldMarShift))
}

// registerMdma_c12mdrType MDMA channel x Mask Data register
type registerMdma_c12mdrType uint32

const (
	RegisterMdma_c12mdrFieldMdrShift = 0
	RegisterMdma_c12mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c12mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c12mdrFieldMdrMask) >> RegisterMdma_c12mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c12mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c12mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c12mdrFieldMdrShift))
}

// registerMdma_c13isrType MDMA channel x interrupt/status register
type registerMdma_c13isrType uint32

const (
	RegisterMdma_c13isrFieldTeif13Shift = 0
	RegisterMdma_c13isrFieldTeif13Mask  = 0x1
)

// GetTeif13 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c13isrType) GetTeif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13isrFieldTeif13Mask) != 0
}

// SetTeif13 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c13isrType) SetTeif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13isrFieldTeif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13isrFieldTeif13Mask)
	}
}

const (
	RegisterMdma_c13isrFieldCtcif13Shift = 1
	RegisterMdma_c13isrFieldCtcif13Mask  = 0x2
)

// GetCtcif13 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c13isrType) GetCtcif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13isrFieldCtcif13Mask) != 0
}

// SetCtcif13 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c13isrType) SetCtcif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13isrFieldCtcif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13isrFieldCtcif13Mask)
	}
}

const (
	RegisterMdma_c13isrFieldBrtif13Shift = 2
	RegisterMdma_c13isrFieldBrtif13Mask  = 0x4
)

// GetBrtif13 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c13isrType) GetBrtif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13isrFieldBrtif13Mask) != 0
}

// SetBrtif13 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c13isrType) SetBrtif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13isrFieldBrtif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13isrFieldBrtif13Mask)
	}
}

const (
	RegisterMdma_c13isrFieldBtif13Shift = 3
	RegisterMdma_c13isrFieldBtif13Mask  = 0x8
)

// GetBtif13 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c13isrType) GetBtif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13isrFieldBtif13Mask) != 0
}

// SetBtif13 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c13isrType) SetBtif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13isrFieldBtif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13isrFieldBtif13Mask)
	}
}

const (
	RegisterMdma_c13isrFieldTcif13Shift = 4
	RegisterMdma_c13isrFieldTcif13Mask  = 0x10
)

// GetTcif13 channel x buffer transfer complete
func (r *registerMdma_c13isrType) GetTcif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13isrFieldTcif13Mask) != 0
}

// SetTcif13 channel x buffer transfer complete
func (r *registerMdma_c13isrType) SetTcif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13isrFieldTcif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13isrFieldTcif13Mask)
	}
}

const (
	RegisterMdma_c13isrFieldCrqa13Shift = 16
	RegisterMdma_c13isrFieldCrqa13Mask  = 0x10000
)

// GetCrqa13 channel x request active flag
func (r *registerMdma_c13isrType) GetCrqa13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13isrFieldCrqa13Mask) != 0
}

// SetCrqa13 channel x request active flag
func (r *registerMdma_c13isrType) SetCrqa13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13isrFieldCrqa13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13isrFieldCrqa13Mask)
	}
}

// registerMdma_c13ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c13ifcrType uint32

const (
	RegisterMdma_c13ifcrFieldCteif13Shift = 0
	RegisterMdma_c13ifcrFieldCteif13Mask  = 0x1
)

// GetCteif13 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c13ifcrType) GetCteif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13ifcrFieldCteif13Mask) != 0
}

// SetCteif13 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c13ifcrType) SetCteif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13ifcrFieldCteif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13ifcrFieldCteif13Mask)
	}
}

const (
	RegisterMdma_c13ifcrFieldCctcif13Shift = 1
	RegisterMdma_c13ifcrFieldCctcif13Mask  = 0x2
)

// GetCctcif13 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c13ifcrType) GetCctcif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13ifcrFieldCctcif13Mask) != 0
}

// SetCctcif13 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c13ifcrType) SetCctcif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13ifcrFieldCctcif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13ifcrFieldCctcif13Mask)
	}
}

const (
	RegisterMdma_c13ifcrFieldCbrtif13Shift = 2
	RegisterMdma_c13ifcrFieldCbrtif13Mask  = 0x4
)

// GetCbrtif13 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c13ifcrType) GetCbrtif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13ifcrFieldCbrtif13Mask) != 0
}

// SetCbrtif13 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c13ifcrType) SetCbrtif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13ifcrFieldCbrtif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13ifcrFieldCbrtif13Mask)
	}
}

const (
	RegisterMdma_c13ifcrFieldCbtif13Shift = 3
	RegisterMdma_c13ifcrFieldCbtif13Mask  = 0x8
)

// GetCbtif13 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c13ifcrType) GetCbtif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13ifcrFieldCbtif13Mask) != 0
}

// SetCbtif13 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c13ifcrType) SetCbtif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13ifcrFieldCbtif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13ifcrFieldCbtif13Mask)
	}
}

const (
	RegisterMdma_c13ifcrFieldCltcif13Shift = 4
	RegisterMdma_c13ifcrFieldCltcif13Mask  = 0x10
)

// GetCltcif13 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c13ifcrType) GetCltcif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13ifcrFieldCltcif13Mask) != 0
}

// SetCltcif13 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c13ifcrType) SetCltcif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13ifcrFieldCltcif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13ifcrFieldCltcif13Mask)
	}
}

// registerMdma_c13esrType MDMA Channel x error status register
type registerMdma_c13esrType uint32

const (
	RegisterMdma_c13esrFieldTeaShift = 0
	RegisterMdma_c13esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c13esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13esrFieldTeaMask) >> RegisterMdma_c13esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c13esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c13esrFieldTeaShift))
}

const (
	RegisterMdma_c13esrFieldTedShift = 7
	RegisterMdma_c13esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c13esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c13esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13esrFieldTedMask)
	}
}

const (
	RegisterMdma_c13esrFieldTeldShift = 8
	RegisterMdma_c13esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c13esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c13esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c13esrFieldTemdShift = 9
	RegisterMdma_c13esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c13esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c13esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c13esrFieldAseShift = 10
	RegisterMdma_c13esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c13esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c13esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13esrFieldAseMask)
	}
}

const (
	RegisterMdma_c13esrFieldBseShift = 11
	RegisterMdma_c13esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c13esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c13esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13esrFieldBseMask)
	}
}

// registerMdma_c13crType This register is used to control the concerned channel.
type registerMdma_c13crType uint32

const (
	RegisterMdma_c13crFieldEnShift = 0
	RegisterMdma_c13crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c13crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c13crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13crFieldEnMask)
	}
}

const (
	RegisterMdma_c13crFieldTeieShift = 1
	RegisterMdma_c13crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c13crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c13crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13crFieldTeieMask)
	}
}

const (
	RegisterMdma_c13crFieldCtcieShift = 2
	RegisterMdma_c13crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c13crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c13crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c13crFieldBrtieShift = 3
	RegisterMdma_c13crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c13crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c13crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c13crFieldBtieShift = 4
	RegisterMdma_c13crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c13crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c13crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13crFieldBtieMask)
	}
}

const (
	RegisterMdma_c13crFieldTcieShift = 5
	RegisterMdma_c13crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c13crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c13crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13crFieldTcieMask)
	}
}

const (
	RegisterMdma_c13crFieldPlShift = 6
	RegisterMdma_c13crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c13crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13crFieldPlMask) >> RegisterMdma_c13crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c13crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13crFieldPlMask)|(uint32(value)<<RegisterMdma_c13crFieldPlShift))
}

const (
	RegisterMdma_c13crFieldBexShift = 12
	RegisterMdma_c13crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c13crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c13crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13crFieldBexMask)
	}
}

const (
	RegisterMdma_c13crFieldHexShift = 13
	RegisterMdma_c13crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c13crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c13crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13crFieldHexMask)
	}
}

const (
	RegisterMdma_c13crFieldWexShift = 14
	RegisterMdma_c13crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c13crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c13crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13crFieldWexMask)
	}
}

const (
	RegisterMdma_c13crFieldSwrqShift = 16
	RegisterMdma_c13crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c13crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13crFieldSwrqMask)
	}
}

// registerMdma_c13tcrType This register is used to configure the concerned channel.
type registerMdma_c13tcrType uint32

const (
	RegisterMdma_c13tcrFieldSincShift = 0
	RegisterMdma_c13tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c13tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldSincMask) >> RegisterMdma_c13tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c13tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c13tcrFieldSincShift))
}

const (
	RegisterMdma_c13tcrFieldDincShift = 2
	RegisterMdma_c13tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c13tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldDincMask) >> RegisterMdma_c13tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c13tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c13tcrFieldDincShift))
}

const (
	RegisterMdma_c13tcrFieldSsizeShift = 4
	RegisterMdma_c13tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c13tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldSsizeMask) >> RegisterMdma_c13tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c13tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c13tcrFieldSsizeShift))
}

const (
	RegisterMdma_c13tcrFieldDsizeShift = 6
	RegisterMdma_c13tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c13tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldDsizeMask) >> RegisterMdma_c13tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c13tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c13tcrFieldDsizeShift))
}

const (
	RegisterMdma_c13tcrFieldSincosShift = 8
	RegisterMdma_c13tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c13tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldSincosMask) >> RegisterMdma_c13tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c13tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c13tcrFieldSincosShift))
}

const (
	RegisterMdma_c13tcrFieldDincosShift = 10
	RegisterMdma_c13tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c13tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldDincosMask) >> RegisterMdma_c13tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c13tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c13tcrFieldDincosShift))
}

const (
	RegisterMdma_c13tcrFieldSburstShift = 12
	RegisterMdma_c13tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c13tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldSburstMask) >> RegisterMdma_c13tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c13tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c13tcrFieldSburstShift))
}

const (
	RegisterMdma_c13tcrFieldDburstShift = 15
	RegisterMdma_c13tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c13tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldDburstMask) >> RegisterMdma_c13tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c13tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c13tcrFieldDburstShift))
}

const (
	RegisterMdma_c13tcrFieldTlenShift = 18
	RegisterMdma_c13tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c13tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldTlenMask) >> RegisterMdma_c13tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c13tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c13tcrFieldTlenShift))
}

const (
	RegisterMdma_c13tcrFieldPkeShift = 25
	RegisterMdma_c13tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c13tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c13tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c13tcrFieldPamShift = 26
	RegisterMdma_c13tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c13tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldPamMask) >> RegisterMdma_c13tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c13tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c13tcrFieldPamShift))
}

const (
	RegisterMdma_c13tcrFieldTrgmShift = 28
	RegisterMdma_c13tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c13tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldTrgmMask) >> RegisterMdma_c13tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c13tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c13tcrFieldTrgmShift))
}

const (
	RegisterMdma_c13tcrFieldSwrmShift = 30
	RegisterMdma_c13tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c13tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c13tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c13tcrFieldBwmShift = 31
	RegisterMdma_c13tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c13tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c13tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tcrFieldBwmMask)
	}
}

// registerMdma_c13bndtrType MDMA Channel x block number of data register
type registerMdma_c13bndtrType uint32

const (
	RegisterMdma_c13bndtrFieldBndtShift = 0
	RegisterMdma_c13bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c13bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13bndtrFieldBndtMask) >> RegisterMdma_c13bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c13bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c13bndtrFieldBndtShift))
}

const (
	RegisterMdma_c13bndtrFieldBrsumShift = 18
	RegisterMdma_c13bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c13bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c13bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c13bndtrFieldBrdumShift = 19
	RegisterMdma_c13bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c13bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c13bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c13bndtrFieldBrcShift = 20
	RegisterMdma_c13bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c13bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13bndtrFieldBrcMask) >> RegisterMdma_c13bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c13bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c13bndtrFieldBrcShift))
}

// registerMdma_c13sarType MDMA channel x source address register
type registerMdma_c13sarType uint32

const (
	RegisterMdma_c13sarFieldSarShift = 0
	RegisterMdma_c13sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c13sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13sarFieldSarMask) >> RegisterMdma_c13sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c13sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13sarFieldSarMask)|(uint32(value)<<RegisterMdma_c13sarFieldSarShift))
}

// registerMdma_c13darType MDMA channel x destination address register
type registerMdma_c13darType uint32

const (
	RegisterMdma_c13darFieldDarShift = 0
	RegisterMdma_c13darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c13darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13darFieldDarMask) >> RegisterMdma_c13darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c13darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13darFieldDarMask)|(uint32(value)<<RegisterMdma_c13darFieldDarShift))
}

// registerMdma_c13brurType MDMA channel x Block Repeat address Update register
type registerMdma_c13brurType uint32

const (
	RegisterMdma_c13brurFieldSuvShift = 0
	RegisterMdma_c13brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c13brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13brurFieldSuvMask) >> RegisterMdma_c13brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c13brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c13brurFieldSuvShift))
}

const (
	RegisterMdma_c13brurFieldDuvShift = 16
	RegisterMdma_c13brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c13brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13brurFieldDuvMask) >> RegisterMdma_c13brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c13brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c13brurFieldDuvShift))
}

// registerMdma_c13larType MDMA channel x Link Address register
type registerMdma_c13larType uint32

const (
	RegisterMdma_c13larFieldLarShift = 0
	RegisterMdma_c13larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c13larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13larFieldLarMask) >> RegisterMdma_c13larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c13larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13larFieldLarMask)|(uint32(value)<<RegisterMdma_c13larFieldLarShift))
}

// registerMdma_c13tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c13tbrType uint32

const (
	RegisterMdma_c13tbrFieldTselShift = 0
	RegisterMdma_c13tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c13tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tbrFieldTselMask) >> RegisterMdma_c13tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c13tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c13tbrFieldTselShift))
}

const (
	RegisterMdma_c13tbrFieldSbusShift = 16
	RegisterMdma_c13tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c13tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c13tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c13tbrFieldDbusShift = 17
	RegisterMdma_c13tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c13tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c13tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c13tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13tbrFieldDbusMask)
	}
}

// registerMdma_c13marType MDMA channel x Mask address register
type registerMdma_c13marType uint32

const (
	RegisterMdma_c13marFieldMarShift = 0
	RegisterMdma_c13marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c13marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13marFieldMarMask) >> RegisterMdma_c13marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c13marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13marFieldMarMask)|(uint32(value)<<RegisterMdma_c13marFieldMarShift))
}

// registerMdma_c13mdrType MDMA channel x Mask Data register
type registerMdma_c13mdrType uint32

const (
	RegisterMdma_c13mdrFieldMdrShift = 0
	RegisterMdma_c13mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c13mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c13mdrFieldMdrMask) >> RegisterMdma_c13mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c13mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c13mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c13mdrFieldMdrShift))
}

// registerMdma_c14isrType MDMA channel x interrupt/status register
type registerMdma_c14isrType uint32

const (
	RegisterMdma_c14isrFieldTeif14Shift = 0
	RegisterMdma_c14isrFieldTeif14Mask  = 0x1
)

// GetTeif14 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c14isrType) GetTeif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14isrFieldTeif14Mask) != 0
}

// SetTeif14 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c14isrType) SetTeif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14isrFieldTeif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14isrFieldTeif14Mask)
	}
}

const (
	RegisterMdma_c14isrFieldCtcif14Shift = 1
	RegisterMdma_c14isrFieldCtcif14Mask  = 0x2
)

// GetCtcif14 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c14isrType) GetCtcif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14isrFieldCtcif14Mask) != 0
}

// SetCtcif14 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c14isrType) SetCtcif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14isrFieldCtcif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14isrFieldCtcif14Mask)
	}
}

const (
	RegisterMdma_c14isrFieldBrtif14Shift = 2
	RegisterMdma_c14isrFieldBrtif14Mask  = 0x4
)

// GetBrtif14 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c14isrType) GetBrtif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14isrFieldBrtif14Mask) != 0
}

// SetBrtif14 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c14isrType) SetBrtif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14isrFieldBrtif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14isrFieldBrtif14Mask)
	}
}

const (
	RegisterMdma_c14isrFieldBtif14Shift = 3
	RegisterMdma_c14isrFieldBtif14Mask  = 0x8
)

// GetBtif14 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c14isrType) GetBtif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14isrFieldBtif14Mask) != 0
}

// SetBtif14 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c14isrType) SetBtif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14isrFieldBtif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14isrFieldBtif14Mask)
	}
}

const (
	RegisterMdma_c14isrFieldTcif14Shift = 4
	RegisterMdma_c14isrFieldTcif14Mask  = 0x10
)

// GetTcif14 channel x buffer transfer complete
func (r *registerMdma_c14isrType) GetTcif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14isrFieldTcif14Mask) != 0
}

// SetTcif14 channel x buffer transfer complete
func (r *registerMdma_c14isrType) SetTcif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14isrFieldTcif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14isrFieldTcif14Mask)
	}
}

const (
	RegisterMdma_c14isrFieldCrqa14Shift = 16
	RegisterMdma_c14isrFieldCrqa14Mask  = 0x10000
)

// GetCrqa14 channel x request active flag
func (r *registerMdma_c14isrType) GetCrqa14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14isrFieldCrqa14Mask) != 0
}

// SetCrqa14 channel x request active flag
func (r *registerMdma_c14isrType) SetCrqa14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14isrFieldCrqa14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14isrFieldCrqa14Mask)
	}
}

// registerMdma_c14ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c14ifcrType uint32

const (
	RegisterMdma_c14ifcrFieldCteif14Shift = 0
	RegisterMdma_c14ifcrFieldCteif14Mask  = 0x1
)

// GetCteif14 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c14ifcrType) GetCteif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14ifcrFieldCteif14Mask) != 0
}

// SetCteif14 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c14ifcrType) SetCteif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14ifcrFieldCteif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14ifcrFieldCteif14Mask)
	}
}

const (
	RegisterMdma_c14ifcrFieldCctcif14Shift = 1
	RegisterMdma_c14ifcrFieldCctcif14Mask  = 0x2
)

// GetCctcif14 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c14ifcrType) GetCctcif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14ifcrFieldCctcif14Mask) != 0
}

// SetCctcif14 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c14ifcrType) SetCctcif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14ifcrFieldCctcif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14ifcrFieldCctcif14Mask)
	}
}

const (
	RegisterMdma_c14ifcrFieldCbrtif14Shift = 2
	RegisterMdma_c14ifcrFieldCbrtif14Mask  = 0x4
)

// GetCbrtif14 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c14ifcrType) GetCbrtif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14ifcrFieldCbrtif14Mask) != 0
}

// SetCbrtif14 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c14ifcrType) SetCbrtif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14ifcrFieldCbrtif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14ifcrFieldCbrtif14Mask)
	}
}

const (
	RegisterMdma_c14ifcrFieldCbtif14Shift = 3
	RegisterMdma_c14ifcrFieldCbtif14Mask  = 0x8
)

// GetCbtif14 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c14ifcrType) GetCbtif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14ifcrFieldCbtif14Mask) != 0
}

// SetCbtif14 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c14ifcrType) SetCbtif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14ifcrFieldCbtif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14ifcrFieldCbtif14Mask)
	}
}

const (
	RegisterMdma_c14ifcrFieldCltcif14Shift = 4
	RegisterMdma_c14ifcrFieldCltcif14Mask  = 0x10
)

// GetCltcif14 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c14ifcrType) GetCltcif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14ifcrFieldCltcif14Mask) != 0
}

// SetCltcif14 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c14ifcrType) SetCltcif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14ifcrFieldCltcif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14ifcrFieldCltcif14Mask)
	}
}

// registerMdma_c14esrType MDMA Channel x error status register
type registerMdma_c14esrType uint32

const (
	RegisterMdma_c14esrFieldTeaShift = 0
	RegisterMdma_c14esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c14esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14esrFieldTeaMask) >> RegisterMdma_c14esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c14esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c14esrFieldTeaShift))
}

const (
	RegisterMdma_c14esrFieldTedShift = 7
	RegisterMdma_c14esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c14esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c14esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14esrFieldTedMask)
	}
}

const (
	RegisterMdma_c14esrFieldTeldShift = 8
	RegisterMdma_c14esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c14esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c14esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c14esrFieldTemdShift = 9
	RegisterMdma_c14esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c14esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c14esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c14esrFieldAseShift = 10
	RegisterMdma_c14esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c14esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c14esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14esrFieldAseMask)
	}
}

const (
	RegisterMdma_c14esrFieldBseShift = 11
	RegisterMdma_c14esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c14esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c14esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14esrFieldBseMask)
	}
}

// registerMdma_c14crType This register is used to control the concerned channel.
type registerMdma_c14crType uint32

const (
	RegisterMdma_c14crFieldEnShift = 0
	RegisterMdma_c14crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c14crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c14crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14crFieldEnMask)
	}
}

const (
	RegisterMdma_c14crFieldTeieShift = 1
	RegisterMdma_c14crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c14crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c14crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14crFieldTeieMask)
	}
}

const (
	RegisterMdma_c14crFieldCtcieShift = 2
	RegisterMdma_c14crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c14crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c14crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c14crFieldBrtieShift = 3
	RegisterMdma_c14crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c14crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c14crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c14crFieldBtieShift = 4
	RegisterMdma_c14crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c14crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c14crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14crFieldBtieMask)
	}
}

const (
	RegisterMdma_c14crFieldTcieShift = 5
	RegisterMdma_c14crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c14crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c14crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14crFieldTcieMask)
	}
}

const (
	RegisterMdma_c14crFieldPlShift = 6
	RegisterMdma_c14crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c14crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14crFieldPlMask) >> RegisterMdma_c14crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c14crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14crFieldPlMask)|(uint32(value)<<RegisterMdma_c14crFieldPlShift))
}

const (
	RegisterMdma_c14crFieldBexShift = 12
	RegisterMdma_c14crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c14crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c14crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14crFieldBexMask)
	}
}

const (
	RegisterMdma_c14crFieldHexShift = 13
	RegisterMdma_c14crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c14crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c14crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14crFieldHexMask)
	}
}

const (
	RegisterMdma_c14crFieldWexShift = 14
	RegisterMdma_c14crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c14crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c14crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14crFieldWexMask)
	}
}

const (
	RegisterMdma_c14crFieldSwrqShift = 16
	RegisterMdma_c14crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c14crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14crFieldSwrqMask)
	}
}

// registerMdma_c14tcrType This register is used to configure the concerned channel.
type registerMdma_c14tcrType uint32

const (
	RegisterMdma_c14tcrFieldSincShift = 0
	RegisterMdma_c14tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c14tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldSincMask) >> RegisterMdma_c14tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c14tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c14tcrFieldSincShift))
}

const (
	RegisterMdma_c14tcrFieldDincShift = 2
	RegisterMdma_c14tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c14tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldDincMask) >> RegisterMdma_c14tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c14tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c14tcrFieldDincShift))
}

const (
	RegisterMdma_c14tcrFieldSsizeShift = 4
	RegisterMdma_c14tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c14tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldSsizeMask) >> RegisterMdma_c14tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c14tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c14tcrFieldSsizeShift))
}

const (
	RegisterMdma_c14tcrFieldDsizeShift = 6
	RegisterMdma_c14tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c14tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldDsizeMask) >> RegisterMdma_c14tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c14tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c14tcrFieldDsizeShift))
}

const (
	RegisterMdma_c14tcrFieldSincosShift = 8
	RegisterMdma_c14tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c14tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldSincosMask) >> RegisterMdma_c14tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c14tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c14tcrFieldSincosShift))
}

const (
	RegisterMdma_c14tcrFieldDincosShift = 10
	RegisterMdma_c14tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c14tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldDincosMask) >> RegisterMdma_c14tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c14tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c14tcrFieldDincosShift))
}

const (
	RegisterMdma_c14tcrFieldSburstShift = 12
	RegisterMdma_c14tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c14tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldSburstMask) >> RegisterMdma_c14tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c14tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c14tcrFieldSburstShift))
}

const (
	RegisterMdma_c14tcrFieldDburstShift = 15
	RegisterMdma_c14tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c14tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldDburstMask) >> RegisterMdma_c14tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c14tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c14tcrFieldDburstShift))
}

const (
	RegisterMdma_c14tcrFieldTlenShift = 18
	RegisterMdma_c14tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c14tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldTlenMask) >> RegisterMdma_c14tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c14tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c14tcrFieldTlenShift))
}

const (
	RegisterMdma_c14tcrFieldPkeShift = 25
	RegisterMdma_c14tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c14tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c14tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c14tcrFieldPamShift = 26
	RegisterMdma_c14tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c14tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldPamMask) >> RegisterMdma_c14tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c14tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c14tcrFieldPamShift))
}

const (
	RegisterMdma_c14tcrFieldTrgmShift = 28
	RegisterMdma_c14tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c14tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldTrgmMask) >> RegisterMdma_c14tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c14tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c14tcrFieldTrgmShift))
}

const (
	RegisterMdma_c14tcrFieldSwrmShift = 30
	RegisterMdma_c14tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c14tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c14tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c14tcrFieldBwmShift = 31
	RegisterMdma_c14tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c14tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c14tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tcrFieldBwmMask)
	}
}

// registerMdma_c14bndtrType MDMA Channel x block number of data register
type registerMdma_c14bndtrType uint32

const (
	RegisterMdma_c14bndtrFieldBndtShift = 0
	RegisterMdma_c14bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c14bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14bndtrFieldBndtMask) >> RegisterMdma_c14bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c14bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c14bndtrFieldBndtShift))
}

const (
	RegisterMdma_c14bndtrFieldBrsumShift = 18
	RegisterMdma_c14bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c14bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c14bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c14bndtrFieldBrdumShift = 19
	RegisterMdma_c14bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c14bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c14bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c14bndtrFieldBrcShift = 20
	RegisterMdma_c14bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c14bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14bndtrFieldBrcMask) >> RegisterMdma_c14bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c14bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c14bndtrFieldBrcShift))
}

// registerMdma_c14sarType MDMA channel x source address register
type registerMdma_c14sarType uint32

const (
	RegisterMdma_c14sarFieldSarShift = 0
	RegisterMdma_c14sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c14sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14sarFieldSarMask) >> RegisterMdma_c14sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c14sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14sarFieldSarMask)|(uint32(value)<<RegisterMdma_c14sarFieldSarShift))
}

// registerMdma_c14darType MDMA channel x destination address register
type registerMdma_c14darType uint32

const (
	RegisterMdma_c14darFieldDarShift = 0
	RegisterMdma_c14darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c14darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14darFieldDarMask) >> RegisterMdma_c14darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c14darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14darFieldDarMask)|(uint32(value)<<RegisterMdma_c14darFieldDarShift))
}

// registerMdma_c14brurType MDMA channel x Block Repeat address Update register
type registerMdma_c14brurType uint32

const (
	RegisterMdma_c14brurFieldSuvShift = 0
	RegisterMdma_c14brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c14brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14brurFieldSuvMask) >> RegisterMdma_c14brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c14brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c14brurFieldSuvShift))
}

const (
	RegisterMdma_c14brurFieldDuvShift = 16
	RegisterMdma_c14brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c14brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14brurFieldDuvMask) >> RegisterMdma_c14brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c14brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c14brurFieldDuvShift))
}

// registerMdma_c14larType MDMA channel x Link Address register
type registerMdma_c14larType uint32

const (
	RegisterMdma_c14larFieldLarShift = 0
	RegisterMdma_c14larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c14larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14larFieldLarMask) >> RegisterMdma_c14larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c14larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14larFieldLarMask)|(uint32(value)<<RegisterMdma_c14larFieldLarShift))
}

// registerMdma_c14tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c14tbrType uint32

const (
	RegisterMdma_c14tbrFieldTselShift = 0
	RegisterMdma_c14tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c14tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tbrFieldTselMask) >> RegisterMdma_c14tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c14tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c14tbrFieldTselShift))
}

const (
	RegisterMdma_c14tbrFieldSbusShift = 16
	RegisterMdma_c14tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c14tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c14tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c14tbrFieldDbusShift = 17
	RegisterMdma_c14tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c14tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c14tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c14tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14tbrFieldDbusMask)
	}
}

// registerMdma_c14marType MDMA channel x Mask address register
type registerMdma_c14marType uint32

const (
	RegisterMdma_c14marFieldMarShift = 0
	RegisterMdma_c14marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c14marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14marFieldMarMask) >> RegisterMdma_c14marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c14marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14marFieldMarMask)|(uint32(value)<<RegisterMdma_c14marFieldMarShift))
}

// registerMdma_c14mdrType MDMA channel x Mask Data register
type registerMdma_c14mdrType uint32

const (
	RegisterMdma_c14mdrFieldMdrShift = 0
	RegisterMdma_c14mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c14mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c14mdrFieldMdrMask) >> RegisterMdma_c14mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c14mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c14mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c14mdrFieldMdrShift))
}

// registerMdma_c15isrType MDMA channel x interrupt/status register
type registerMdma_c15isrType uint32

const (
	RegisterMdma_c15isrFieldTeif15Shift = 0
	RegisterMdma_c15isrFieldTeif15Mask  = 0x1
)

// GetTeif15 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c15isrType) GetTeif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15isrFieldTeif15Mask) != 0
}

// SetTeif15 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c15isrType) SetTeif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15isrFieldTeif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15isrFieldTeif15Mask)
	}
}

const (
	RegisterMdma_c15isrFieldCtcif15Shift = 1
	RegisterMdma_c15isrFieldCtcif15Mask  = 0x2
)

// GetCtcif15 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c15isrType) GetCtcif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15isrFieldCtcif15Mask) != 0
}

// SetCtcif15 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *registerMdma_c15isrType) SetCtcif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15isrFieldCtcif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15isrFieldCtcif15Mask)
	}
}

const (
	RegisterMdma_c15isrFieldBrtif15Shift = 2
	RegisterMdma_c15isrFieldBrtif15Mask  = 0x4
)

// GetBrtif15 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c15isrType) GetBrtif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15isrFieldBrtif15Mask) != 0
}

// SetBrtif15 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c15isrType) SetBrtif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15isrFieldBrtif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15isrFieldBrtif15Mask)
	}
}

const (
	RegisterMdma_c15isrFieldBtif15Shift = 3
	RegisterMdma_c15isrFieldBtif15Mask  = 0x8
)

// GetBtif15 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c15isrType) GetBtif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15isrFieldBtif15Mask) != 0
}

// SetBtif15 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *registerMdma_c15isrType) SetBtif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15isrFieldBtif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15isrFieldBtif15Mask)
	}
}

const (
	RegisterMdma_c15isrFieldTcif15Shift = 4
	RegisterMdma_c15isrFieldTcif15Mask  = 0x10
)

// GetTcif15 channel x buffer transfer complete
func (r *registerMdma_c15isrType) GetTcif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15isrFieldTcif15Mask) != 0
}

// SetTcif15 channel x buffer transfer complete
func (r *registerMdma_c15isrType) SetTcif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15isrFieldTcif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15isrFieldTcif15Mask)
	}
}

const (
	RegisterMdma_c15isrFieldCrqa15Shift = 16
	RegisterMdma_c15isrFieldCrqa15Mask  = 0x10000
)

// GetCrqa15 channel x request active flag
func (r *registerMdma_c15isrType) GetCrqa15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15isrFieldCrqa15Mask) != 0
}

// SetCrqa15 channel x request active flag
func (r *registerMdma_c15isrType) SetCrqa15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15isrFieldCrqa15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15isrFieldCrqa15Mask)
	}
}

// registerMdma_c15ifcrType MDMA channel x interrupt flag clear register
type registerMdma_c15ifcrType uint32

const (
	RegisterMdma_c15ifcrFieldCteif15Shift = 0
	RegisterMdma_c15ifcrFieldCteif15Mask  = 0x1
)

// GetCteif15 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c15ifcrType) GetCteif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15ifcrFieldCteif15Mask) != 0
}

// SetCteif15 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *registerMdma_c15ifcrType) SetCteif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15ifcrFieldCteif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15ifcrFieldCteif15Mask)
	}
}

const (
	RegisterMdma_c15ifcrFieldCctcif15Shift = 1
	RegisterMdma_c15ifcrFieldCctcif15Mask  = 0x2
)

// GetCctcif15 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c15ifcrType) GetCctcif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15ifcrFieldCctcif15Mask) != 0
}

// SetCctcif15 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *registerMdma_c15ifcrType) SetCctcif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15ifcrFieldCctcif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15ifcrFieldCctcif15Mask)
	}
}

const (
	RegisterMdma_c15ifcrFieldCbrtif15Shift = 2
	RegisterMdma_c15ifcrFieldCbrtif15Mask  = 0x4
)

// GetCbrtif15 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c15ifcrType) GetCbrtif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15ifcrFieldCbrtif15Mask) != 0
}

// SetCbrtif15 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *registerMdma_c15ifcrType) SetCbrtif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15ifcrFieldCbrtif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15ifcrFieldCbrtif15Mask)
	}
}

const (
	RegisterMdma_c15ifcrFieldCbtif15Shift = 3
	RegisterMdma_c15ifcrFieldCbtif15Mask  = 0x8
)

// GetCbtif15 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c15ifcrType) GetCbtif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15ifcrFieldCbtif15Mask) != 0
}

// SetCbtif15 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *registerMdma_c15ifcrType) SetCbtif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15ifcrFieldCbtif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15ifcrFieldCbtif15Mask)
	}
}

const (
	RegisterMdma_c15ifcrFieldCltcif15Shift = 4
	RegisterMdma_c15ifcrFieldCltcif15Mask  = 0x10
)

// GetCltcif15 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c15ifcrType) GetCltcif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15ifcrFieldCltcif15Mask) != 0
}

// SetCltcif15 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *registerMdma_c15ifcrType) SetCltcif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15ifcrFieldCltcif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15ifcrFieldCltcif15Mask)
	}
}

// registerMdma_c15esrType MDMA Channel x error status register
type registerMdma_c15esrType uint32

const (
	RegisterMdma_c15esrFieldTeaShift = 0
	RegisterMdma_c15esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c15esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15esrFieldTeaMask) >> RegisterMdma_c15esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *registerMdma_c15esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15esrFieldTeaMask)|(uint32(value)<<RegisterMdma_c15esrFieldTeaShift))
}

const (
	RegisterMdma_c15esrFieldTedShift = 7
	RegisterMdma_c15esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c15esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *registerMdma_c15esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15esrFieldTedMask)
	}
}

const (
	RegisterMdma_c15esrFieldTeldShift = 8
	RegisterMdma_c15esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c15esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c15esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15esrFieldTeldMask)
	}
}

const (
	RegisterMdma_c15esrFieldTemdShift = 9
	RegisterMdma_c15esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c15esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c15esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15esrFieldTemdMask)
	}
}

const (
	RegisterMdma_c15esrFieldAseShift = 10
	RegisterMdma_c15esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c15esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c15esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15esrFieldAseMask)
	}
}

const (
	RegisterMdma_c15esrFieldBseShift = 11
	RegisterMdma_c15esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c15esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *registerMdma_c15esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15esrFieldBseMask)
	}
}

// registerMdma_c15crType This register is used to control the concerned channel.
type registerMdma_c15crType uint32

const (
	RegisterMdma_c15crFieldEnShift = 0
	RegisterMdma_c15crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *registerMdma_c15crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15crFieldEnMask) != 0
}

// SetEn channel enable
func (r *registerMdma_c15crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15crFieldEnMask)
	}
}

const (
	RegisterMdma_c15crFieldTeieShift = 1
	RegisterMdma_c15crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c15crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c15crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15crFieldTeieMask)
	}
}

const (
	RegisterMdma_c15crFieldCtcieShift = 2
	RegisterMdma_c15crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c15crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c15crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15crFieldCtcieMask)
	}
}

const (
	RegisterMdma_c15crFieldBrtieShift = 3
	RegisterMdma_c15crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c15crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c15crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15crFieldBrtieMask)
	}
}

const (
	RegisterMdma_c15crFieldBtieShift = 4
	RegisterMdma_c15crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c15crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c15crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15crFieldBtieMask)
	}
}

const (
	RegisterMdma_c15crFieldTcieShift = 5
	RegisterMdma_c15crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c15crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *registerMdma_c15crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15crFieldTcieMask)
	}
}

const (
	RegisterMdma_c15crFieldPlShift = 6
	RegisterMdma_c15crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c15crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15crFieldPlMask) >> RegisterMdma_c15crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c15crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15crFieldPlMask)|(uint32(value)<<RegisterMdma_c15crFieldPlShift))
}

const (
	RegisterMdma_c15crFieldBexShift = 12
	RegisterMdma_c15crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *registerMdma_c15crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *registerMdma_c15crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15crFieldBexMask)
	}
}

const (
	RegisterMdma_c15crFieldHexShift = 13
	RegisterMdma_c15crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *registerMdma_c15crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *registerMdma_c15crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15crFieldHexMask)
	}
}

const (
	RegisterMdma_c15crFieldWexShift = 14
	RegisterMdma_c15crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *registerMdma_c15crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *registerMdma_c15crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15crFieldWexMask)
	}
}

const (
	RegisterMdma_c15crFieldSwrqShift = 16
	RegisterMdma_c15crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *registerMdma_c15crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15crFieldSwrqMask)
	}
}

// registerMdma_c15tcrType This register is used to configure the concerned channel.
type registerMdma_c15tcrType uint32

const (
	RegisterMdma_c15tcrFieldSincShift = 0
	RegisterMdma_c15tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c15tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldSincMask) >> RegisterMdma_c15tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *registerMdma_c15tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldSincMask)|(uint32(value)<<RegisterMdma_c15tcrFieldSincShift))
}

const (
	RegisterMdma_c15tcrFieldDincShift = 2
	RegisterMdma_c15tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c15tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldDincMask) >> RegisterMdma_c15tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *registerMdma_c15tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldDincMask)|(uint32(value)<<RegisterMdma_c15tcrFieldDincShift))
}

const (
	RegisterMdma_c15tcrFieldSsizeShift = 4
	RegisterMdma_c15tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c15tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldSsizeMask) >> RegisterMdma_c15tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *registerMdma_c15tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldSsizeMask)|(uint32(value)<<RegisterMdma_c15tcrFieldSsizeShift))
}

const (
	RegisterMdma_c15tcrFieldDsizeShift = 6
	RegisterMdma_c15tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c15tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldDsizeMask) >> RegisterMdma_c15tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *registerMdma_c15tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldDsizeMask)|(uint32(value)<<RegisterMdma_c15tcrFieldDsizeShift))
}

const (
	RegisterMdma_c15tcrFieldSincosShift = 8
	RegisterMdma_c15tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *registerMdma_c15tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldSincosMask) >> RegisterMdma_c15tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *registerMdma_c15tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldSincosMask)|(uint32(value)<<RegisterMdma_c15tcrFieldSincosShift))
}

const (
	RegisterMdma_c15tcrFieldDincosShift = 10
	RegisterMdma_c15tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *registerMdma_c15tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldDincosMask) >> RegisterMdma_c15tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *registerMdma_c15tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldDincosMask)|(uint32(value)<<RegisterMdma_c15tcrFieldDincosShift))
}

const (
	RegisterMdma_c15tcrFieldSburstShift = 12
	RegisterMdma_c15tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *registerMdma_c15tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldSburstMask) >> RegisterMdma_c15tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *registerMdma_c15tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldSburstMask)|(uint32(value)<<RegisterMdma_c15tcrFieldSburstShift))
}

const (
	RegisterMdma_c15tcrFieldDburstShift = 15
	RegisterMdma_c15tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *registerMdma_c15tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldDburstMask) >> RegisterMdma_c15tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *registerMdma_c15tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldDburstMask)|(uint32(value)<<RegisterMdma_c15tcrFieldDburstShift))
}

const (
	RegisterMdma_c15tcrFieldTlenShift = 18
	RegisterMdma_c15tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *registerMdma_c15tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldTlenMask) >> RegisterMdma_c15tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *registerMdma_c15tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldTlenMask)|(uint32(value)<<RegisterMdma_c15tcrFieldTlenShift))
}

const (
	RegisterMdma_c15tcrFieldPkeShift = 25
	RegisterMdma_c15tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c15tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *registerMdma_c15tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldPkeMask)
	}
}

const (
	RegisterMdma_c15tcrFieldPamShift = 26
	RegisterMdma_c15tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c15tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldPamMask) >> RegisterMdma_c15tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *registerMdma_c15tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldPamMask)|(uint32(value)<<RegisterMdma_c15tcrFieldPamShift))
}

const (
	RegisterMdma_c15tcrFieldTrgmShift = 28
	RegisterMdma_c15tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c15tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldTrgmMask) >> RegisterMdma_c15tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c15tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldTrgmMask)|(uint32(value)<<RegisterMdma_c15tcrFieldTrgmShift))
}

const (
	RegisterMdma_c15tcrFieldSwrmShift = 30
	RegisterMdma_c15tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c15tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c15tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldSwrmMask)
	}
}

const (
	RegisterMdma_c15tcrFieldBwmShift = 31
	RegisterMdma_c15tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c15tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *registerMdma_c15tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tcrFieldBwmMask)
	}
}

// registerMdma_c15bndtrType MDMA Channel x block number of data register
type registerMdma_c15bndtrType uint32

const (
	RegisterMdma_c15bndtrFieldBndtShift = 0
	RegisterMdma_c15bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *registerMdma_c15bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15bndtrFieldBndtMask) >> RegisterMdma_c15bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *registerMdma_c15bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15bndtrFieldBndtMask)|(uint32(value)<<RegisterMdma_c15bndtrFieldBndtShift))
}

const (
	RegisterMdma_c15bndtrFieldBrsumShift = 18
	RegisterMdma_c15bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c15bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c15bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdma_c15bndtrFieldBrdumShift = 19
	RegisterMdma_c15bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c15bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c15bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdma_c15bndtrFieldBrcShift = 20
	RegisterMdma_c15bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c15bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15bndtrFieldBrcMask) >> RegisterMdma_c15bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *registerMdma_c15bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15bndtrFieldBrcMask)|(uint32(value)<<RegisterMdma_c15bndtrFieldBrcShift))
}

// registerMdma_c15sarType MDMA channel x source address register
type registerMdma_c15sarType uint32

const (
	RegisterMdma_c15sarFieldSarShift = 0
	RegisterMdma_c15sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *registerMdma_c15sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15sarFieldSarMask) >> RegisterMdma_c15sarFieldSarShift)
}

// SetSar source adr base
func (r *registerMdma_c15sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15sarFieldSarMask)|(uint32(value)<<RegisterMdma_c15sarFieldSarShift))
}

// registerMdma_c15darType MDMA channel x destination address register
type registerMdma_c15darType uint32

const (
	RegisterMdma_c15darFieldDarShift = 0
	RegisterMdma_c15darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *registerMdma_c15darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15darFieldDarMask) >> RegisterMdma_c15darFieldDarShift)
}

// SetDar Destination adr base
func (r *registerMdma_c15darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15darFieldDarMask)|(uint32(value)<<RegisterMdma_c15darFieldDarShift))
}

// registerMdma_c15brurType MDMA channel x Block Repeat address Update register
type registerMdma_c15brurType uint32

const (
	RegisterMdma_c15brurFieldSuvShift = 0
	RegisterMdma_c15brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *registerMdma_c15brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15brurFieldSuvMask) >> RegisterMdma_c15brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *registerMdma_c15brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15brurFieldSuvMask)|(uint32(value)<<RegisterMdma_c15brurFieldSuvShift))
}

const (
	RegisterMdma_c15brurFieldDuvShift = 16
	RegisterMdma_c15brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *registerMdma_c15brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15brurFieldDuvMask) >> RegisterMdma_c15brurFieldDuvShift)
}

// SetDuv destination address update
func (r *registerMdma_c15brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15brurFieldDuvMask)|(uint32(value)<<RegisterMdma_c15brurFieldDuvShift))
}

// registerMdma_c15larType MDMA channel x Link Address register
type registerMdma_c15larType uint32

const (
	RegisterMdma_c15larFieldLarShift = 0
	RegisterMdma_c15larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *registerMdma_c15larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15larFieldLarMask) >> RegisterMdma_c15larFieldLarShift)
}

// SetLar Link address register
func (r *registerMdma_c15larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15larFieldLarMask)|(uint32(value)<<RegisterMdma_c15larFieldLarShift))
}

// registerMdma_c15tbrType MDMA channel x Trigger and Bus selection Register
type registerMdma_c15tbrType uint32

const (
	RegisterMdma_c15tbrFieldTselShift = 0
	RegisterMdma_c15tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *registerMdma_c15tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tbrFieldTselMask) >> RegisterMdma_c15tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *registerMdma_c15tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tbrFieldTselMask)|(uint32(value)<<RegisterMdma_c15tbrFieldTselShift))
}

const (
	RegisterMdma_c15tbrFieldSbusShift = 16
	RegisterMdma_c15tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c15tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c15tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tbrFieldSbusMask)
	}
}

const (
	RegisterMdma_c15tbrFieldDbusShift = 17
	RegisterMdma_c15tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c15tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *registerMdma_c15tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdma_c15tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15tbrFieldDbusMask)
	}
}

// registerMdma_c15marType MDMA channel x Mask address register
type registerMdma_c15marType uint32

const (
	RegisterMdma_c15marFieldMarShift = 0
	RegisterMdma_c15marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *registerMdma_c15marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15marFieldMarMask) >> RegisterMdma_c15marFieldMarShift)
}

// SetMar Mask address
func (r *registerMdma_c15marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15marFieldMarMask)|(uint32(value)<<RegisterMdma_c15marFieldMarShift))
}

// registerMdma_c15mdrType MDMA channel x Mask Data register
type registerMdma_c15mdrType uint32

const (
	RegisterMdma_c15mdrFieldMdrShift = 0
	RegisterMdma_c15mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *registerMdma_c15mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdma_c15mdrFieldMdrMask) >> RegisterMdma_c15mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *registerMdma_c15mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdma_c15mdrFieldMdrMask)|(uint32(value)<<RegisterMdma_c15mdrFieldMdrShift))
}
