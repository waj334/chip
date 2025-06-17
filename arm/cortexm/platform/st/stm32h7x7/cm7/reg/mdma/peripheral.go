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
	Mdmagisr0    RegisterMdmagisr0Type
	_            [60]byte
	Mdmac0isr    RegisterMdmac0isrType
	Mdmac0ifcr   RegisterMdmac0ifcrType
	Mdmac0esr    RegisterMdmac0esrType
	Mdmac0cr     RegisterMdmac0crType
	Mdmac0tcr    RegisterMdmac0tcrType
	Mdmac0bndtr  RegisterMdmac0bndtrType
	Mdmac0sar    RegisterMdmac0sarType
	Mdmac0dar    RegisterMdmac0darType
	Mdmac0brur   RegisterMdmac0brurType
	Mdmac0lar    RegisterMdmac0larType
	Mdmac0tbr    RegisterMdmac0tbrType
	_            [4]byte
	Mdmac0mar    RegisterMdmac0marType
	Mdmac0mdr    RegisterMdmac0mdrType
	_            [8]byte
	Mdmac1isr    RegisterMdmac1isrType
	Mdmac1ifcr   RegisterMdmac1ifcrType
	Mdmac1esr    RegisterMdmac1esrType
	Mdmac1cr     RegisterMdmac1crType
	Mdmac1tcr    RegisterMdmac1tcrType
	Mdmac1bndtr  RegisterMdmac1bndtrType
	Mdmac1sar    RegisterMdmac1sarType
	Mdmac1dar    RegisterMdmac1darType
	Mdmac1brur   RegisterMdmac1brurType
	Mdmac1lar    RegisterMdmac1larType
	Mdmac1tbr    RegisterMdmac1tbrType
	_            [4]byte
	Mdmac1mar    RegisterMdmac1marType
	Mdmac1mdr    RegisterMdmac1mdrType
	_            [8]byte
	Mdmac2isr    RegisterMdmac2isrType
	Mdmac2ifcr   RegisterMdmac2ifcrType
	Mdmac2esr    RegisterMdmac2esrType
	Mdmac2cr     RegisterMdmac2crType
	Mdmac2tcr    RegisterMdmac2tcrType
	Mdmac2bndtr  RegisterMdmac2bndtrType
	Mdmac2sar    RegisterMdmac2sarType
	Mdmac2dar    RegisterMdmac2darType
	Mdmac2brur   RegisterMdmac2brurType
	Mdmac2lar    RegisterMdmac2larType
	Mdmac2tbr    RegisterMdmac2tbrType
	_            [4]byte
	Mdmac2mar    RegisterMdmac2marType
	Mdmac2mdr    RegisterMdmac2mdrType
	_            [8]byte
	Mdmac3isr    RegisterMdmac3isrType
	Mdmac3ifcr   RegisterMdmac3ifcrType
	Mdmac3esr    RegisterMdmac3esrType
	Mdmac3cr     RegisterMdmac3crType
	Mdmac3tcr    RegisterMdmac3tcrType
	Mdmac3bndtr  RegisterMdmac3bndtrType
	Mdmac3sar    RegisterMdmac3sarType
	Mdmac3dar    RegisterMdmac3darType
	Mdmac3brur   RegisterMdmac3brurType
	Mdmac3lar    RegisterMdmac3larType
	Mdmac3tbr    RegisterMdmac3tbrType
	_            [4]byte
	Mdmac3mar    RegisterMdmac3marType
	Mdmac3mdr    RegisterMdmac3mdrType
	_            [8]byte
	Mdmac4isr    RegisterMdmac4isrType
	Mdmac4ifcr   RegisterMdmac4ifcrType
	Mdmac4esr    RegisterMdmac4esrType
	Mdmac4cr     RegisterMdmac4crType
	Mdmac4tcr    RegisterMdmac4tcrType
	Mdmac4bndtr  RegisterMdmac4bndtrType
	Mdmac4sar    RegisterMdmac4sarType
	Mdmac4dar    RegisterMdmac4darType
	Mdmac4brur   RegisterMdmac4brurType
	Mdmac4lar    RegisterMdmac4larType
	Mdmac4tbr    RegisterMdmac4tbrType
	_            [4]byte
	Mdmac4mar    RegisterMdmac4marType
	Mdmac4mdr    RegisterMdmac4mdrType
	_            [8]byte
	Mdmac5isr    RegisterMdmac5isrType
	Mdmac5ifcr   RegisterMdmac5ifcrType
	Mdmac5esr    RegisterMdmac5esrType
	Mdmac5cr     RegisterMdmac5crType
	Mdmac5tcr    RegisterMdmac5tcrType
	Mdmac5bndtr  RegisterMdmac5bndtrType
	Mdmac5sar    RegisterMdmac5sarType
	Mdmac5dar    RegisterMdmac5darType
	Mdmac5brur   RegisterMdmac5brurType
	Mdmac5lar    RegisterMdmac5larType
	Mdmac5tbr    RegisterMdmac5tbrType
	_            [4]byte
	Mdmac5mar    RegisterMdmac5marType
	Mdmac5mdr    RegisterMdmac5mdrType
	_            [8]byte
	Mdmac6isr    RegisterMdmac6isrType
	Mdmac6ifcr   RegisterMdmac6ifcrType
	Mdmac6esr    RegisterMdmac6esrType
	Mdmac6cr     RegisterMdmac6crType
	Mdmac6tcr    RegisterMdmac6tcrType
	Mdmac6bndtr  RegisterMdmac6bndtrType
	Mdmac6sar    RegisterMdmac6sarType
	Mdmac6dar    RegisterMdmac6darType
	Mdmac6brur   RegisterMdmac6brurType
	Mdmac6lar    RegisterMdmac6larType
	Mdmac6tbr    RegisterMdmac6tbrType
	_            [4]byte
	Mdmac6mar    RegisterMdmac6marType
	Mdmac6mdr    RegisterMdmac6mdrType
	_            [8]byte
	Mdmac7isr    RegisterMdmac7isrType
	Mdmac7ifcr   RegisterMdmac7ifcrType
	Mdmac7esr    RegisterMdmac7esrType
	Mdmac7cr     RegisterMdmac7crType
	Mdmac7tcr    RegisterMdmac7tcrType
	Mdmac7bndtr  RegisterMdmac7bndtrType
	Mdmac7sar    RegisterMdmac7sarType
	Mdmac7dar    RegisterMdmac7darType
	Mdmac7brur   RegisterMdmac7brurType
	Mdmac7lar    RegisterMdmac7larType
	Mdmac7tbr    RegisterMdmac7tbrType
	_            [4]byte
	Mdmac7mar    RegisterMdmac7marType
	Mdmac7mdr    RegisterMdmac7mdrType
	_            [8]byte
	Mdmac8isr    RegisterMdmac8isrType
	Mdmac8ifcr   RegisterMdmac8ifcrType
	Mdmac8esr    RegisterMdmac8esrType
	Mdmac8cr     RegisterMdmac8crType
	Mdmac8tcr    RegisterMdmac8tcrType
	Mdmac8bndtr  RegisterMdmac8bndtrType
	Mdmac8sar    RegisterMdmac8sarType
	Mdmac8dar    RegisterMdmac8darType
	Mdmac8brur   RegisterMdmac8brurType
	Mdmac8lar    RegisterMdmac8larType
	Mdmac8tbr    RegisterMdmac8tbrType
	_            [4]byte
	Mdmac8mar    RegisterMdmac8marType
	Mdmac8mdr    RegisterMdmac8mdrType
	_            [8]byte
	Mdmac9isr    RegisterMdmac9isrType
	Mdmac9ifcr   RegisterMdmac9ifcrType
	Mdmac9esr    RegisterMdmac9esrType
	Mdmac9cr     RegisterMdmac9crType
	Mdmac9tcr    RegisterMdmac9tcrType
	Mdmac9bndtr  RegisterMdmac9bndtrType
	Mdmac9sar    RegisterMdmac9sarType
	Mdmac9dar    RegisterMdmac9darType
	Mdmac9brur   RegisterMdmac9brurType
	Mdmac9lar    RegisterMdmac9larType
	Mdmac9tbr    RegisterMdmac9tbrType
	_            [4]byte
	Mdmac9mar    RegisterMdmac9marType
	Mdmac9mdr    RegisterMdmac9mdrType
	_            [8]byte
	Mdmac10isr   RegisterMdmac10isrType
	Mdmac10ifcr  RegisterMdmac10ifcrType
	Mdmac10esr   RegisterMdmac10esrType
	Mdmac10cr    RegisterMdmac10crType
	Mdmac10tcr   RegisterMdmac10tcrType
	Mdmac10bndtr RegisterMdmac10bndtrType
	Mdmac10sar   RegisterMdmac10sarType
	Mdmac10dar   RegisterMdmac10darType
	Mdmac10brur  RegisterMdmac10brurType
	Mdmac10lar   RegisterMdmac10larType
	Mdmac10tbr   RegisterMdmac10tbrType
	_            [4]byte
	Mdmac10mar   RegisterMdmac10marType
	Mdmac10mdr   RegisterMdmac10mdrType
	_            [8]byte
	Mdmac11isr   RegisterMdmac11isrType
	Mdmac11ifcr  RegisterMdmac11ifcrType
	Mdmac11esr   RegisterMdmac11esrType
	Mdmac11cr    RegisterMdmac11crType
	Mdmac11tcr   RegisterMdmac11tcrType
	Mdmac11bndtr RegisterMdmac11bndtrType
	Mdmac11sar   RegisterMdmac11sarType
	Mdmac11dar   RegisterMdmac11darType
	Mdmac11brur  RegisterMdmac11brurType
	Mdmac11lar   RegisterMdmac11larType
	Mdmac11tbr   RegisterMdmac11tbrType
	_            [4]byte
	Mdmac11mar   RegisterMdmac11marType
	Mdmac11mdr   RegisterMdmac11mdrType
	_            [8]byte
	Mdmac12isr   RegisterMdmac12isrType
	Mdmac12ifcr  RegisterMdmac12ifcrType
	Mdmac12esr   RegisterMdmac12esrType
	Mdmac12cr    RegisterMdmac12crType
	Mdmac12tcr   RegisterMdmac12tcrType
	Mdmac12bndtr RegisterMdmac12bndtrType
	Mdmac12sar   RegisterMdmac12sarType
	Mdmac12dar   RegisterMdmac12darType
	Mdmac12brur  RegisterMdmac12brurType
	Mdmac12lar   RegisterMdmac12larType
	Mdmac12tbr   RegisterMdmac12tbrType
	_            [4]byte
	Mdmac12mar   RegisterMdmac12marType
	Mdmac12mdr   RegisterMdmac12mdrType
	_            [8]byte
	Mdmac13isr   RegisterMdmac13isrType
	Mdmac13ifcr  RegisterMdmac13ifcrType
	Mdmac13esr   RegisterMdmac13esrType
	Mdmac13cr    RegisterMdmac13crType
	Mdmac13tcr   RegisterMdmac13tcrType
	Mdmac13bndtr RegisterMdmac13bndtrType
	Mdmac13sar   RegisterMdmac13sarType
	Mdmac13dar   RegisterMdmac13darType
	Mdmac13brur  RegisterMdmac13brurType
	Mdmac13lar   RegisterMdmac13larType
	Mdmac13tbr   RegisterMdmac13tbrType
	_            [4]byte
	Mdmac13mar   RegisterMdmac13marType
	Mdmac13mdr   RegisterMdmac13mdrType
	_            [8]byte
	Mdmac14isr   RegisterMdmac14isrType
	Mdmac14ifcr  RegisterMdmac14ifcrType
	Mdmac14esr   RegisterMdmac14esrType
	Mdmac14cr    RegisterMdmac14crType
	Mdmac14tcr   RegisterMdmac14tcrType
	Mdmac14bndtr RegisterMdmac14bndtrType
	Mdmac14sar   RegisterMdmac14sarType
	Mdmac14dar   RegisterMdmac14darType
	Mdmac14brur  RegisterMdmac14brurType
	Mdmac14lar   RegisterMdmac14larType
	Mdmac14tbr   RegisterMdmac14tbrType
	_            [4]byte
	Mdmac14mar   RegisterMdmac14marType
	Mdmac14mdr   RegisterMdmac14mdrType
	_            [8]byte
	Mdmac15isr   RegisterMdmac15isrType
	Mdmac15ifcr  RegisterMdmac15ifcrType
	Mdmac15esr   RegisterMdmac15esrType
	Mdmac15cr    RegisterMdmac15crType
	Mdmac15tcr   RegisterMdmac15tcrType
	Mdmac15bndtr RegisterMdmac15bndtrType
	Mdmac15sar   RegisterMdmac15sarType
	Mdmac15dar   RegisterMdmac15darType
	Mdmac15brur  RegisterMdmac15brurType
	Mdmac15lar   RegisterMdmac15larType
	Mdmac15tbr   RegisterMdmac15tbrType
	_            [4]byte
	Mdmac15mar   RegisterMdmac15marType
	Mdmac15mdr   RegisterMdmac15mdrType
}

// RegisterMdmagisr0Type MDMA Global Interrupt/Status Register
type RegisterMdmagisr0Type uint32

func (r *RegisterMdmagisr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmagisr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmagisr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmagisr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmagisr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmagisr0FieldGif0Shift = 0
	RegisterMdmagisr0FieldGif0Mask  = 0x1
)

// GetGif0 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif0Mask) != 0
}

// SetGif0 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif0Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif1Shift = 1
	RegisterMdmagisr0FieldGif1Mask  = 0x2
)

// GetGif1 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif1Mask) != 0
}

// SetGif1 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif1Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif2Shift = 2
	RegisterMdmagisr0FieldGif2Mask  = 0x4
)

// GetGif2 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif2Mask) != 0
}

// SetGif2 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif2Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif3Shift = 3
	RegisterMdmagisr0FieldGif3Mask  = 0x8
)

// GetGif3 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif3Mask) != 0
}

// SetGif3 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif3Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif4Shift = 4
	RegisterMdmagisr0FieldGif4Mask  = 0x10
)

// GetGif4 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif4Mask) != 0
}

// SetGif4 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif4Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif5Shift = 5
	RegisterMdmagisr0FieldGif5Mask  = 0x20
)

// GetGif5 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif5Mask) != 0
}

// SetGif5 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif5Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif6Shift = 6
	RegisterMdmagisr0FieldGif6Mask  = 0x40
)

// GetGif6 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif6Mask) != 0
}

// SetGif6 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif6Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif7Shift = 7
	RegisterMdmagisr0FieldGif7Mask  = 0x80
)

// GetGif7 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif7Mask) != 0
}

// SetGif7 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif7Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif8Shift = 8
	RegisterMdmagisr0FieldGif8Mask  = 0x100
)

// GetGif8 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif8Mask) != 0
}

// SetGif8 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif8Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif9Shift = 9
	RegisterMdmagisr0FieldGif9Mask  = 0x200
)

// GetGif9 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif9Mask) != 0
}

// SetGif9 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif9Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif10Shift = 10
	RegisterMdmagisr0FieldGif10Mask  = 0x400
)

// GetGif10 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif10Mask) != 0
}

// SetGif10 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif10Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif11Shift = 11
	RegisterMdmagisr0FieldGif11Mask  = 0x800
)

// GetGif11 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif11Mask) != 0
}

// SetGif11 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif11Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif12Shift = 12
	RegisterMdmagisr0FieldGif12Mask  = 0x1000
)

// GetGif12 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif12Mask) != 0
}

// SetGif12 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif12Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif13Shift = 13
	RegisterMdmagisr0FieldGif13Mask  = 0x2000
)

// GetGif13 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif13Mask) != 0
}

// SetGif13 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif13Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif14Shift = 14
	RegisterMdmagisr0FieldGif14Mask  = 0x4000
)

// GetGif14 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif14Mask) != 0
}

// SetGif14 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif14Mask)
	}
}

const (
	RegisterMdmagisr0FieldGif15Shift = 15
	RegisterMdmagisr0FieldGif15Mask  = 0x8000
)

// GetGif15 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) GetGif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmagisr0FieldGif15Mask) != 0
}

// SetGif15 Channel x global interrupt flag (x=...) This bit is set and reset by hardware. It is a logical OR of all the Channel x interrupt flags (CTCIFx, BTIFx, BRTIFx, TEIFx) which are enabled in the interrupt mask register (CTCIEx, BTIEx, BRTIEx, TEIEx)
func (r *RegisterMdmagisr0Type) SetGif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmagisr0FieldGif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmagisr0FieldGif15Mask)
	}
}

// RegisterMdmac0isrType MDMA channel x interrupt/status register
type RegisterMdmac0isrType uint32

func (r *RegisterMdmac0isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0isrFieldTeif0Shift = 0
	RegisterMdmac0isrFieldTeif0Mask  = 0x1
)

// GetTeif0 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac0isrType) GetTeif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0isrFieldTeif0Mask) != 0
}

// SetTeif0 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac0isrType) SetTeif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0isrFieldTeif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0isrFieldTeif0Mask)
	}
}

const (
	RegisterMdmac0isrFieldCtcif0Shift = 1
	RegisterMdmac0isrFieldCtcif0Mask  = 0x2
)

// GetCtcif0 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac0isrType) GetCtcif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0isrFieldCtcif0Mask) != 0
}

// SetCtcif0 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac0isrType) SetCtcif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0isrFieldCtcif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0isrFieldCtcif0Mask)
	}
}

const (
	RegisterMdmac0isrFieldBrtif0Shift = 2
	RegisterMdmac0isrFieldBrtif0Mask  = 0x4
)

// GetBrtif0 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac0isrType) GetBrtif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0isrFieldBrtif0Mask) != 0
}

// SetBrtif0 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac0isrType) SetBrtif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0isrFieldBrtif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0isrFieldBrtif0Mask)
	}
}

const (
	RegisterMdmac0isrFieldBtif0Shift = 3
	RegisterMdmac0isrFieldBtif0Mask  = 0x8
)

// GetBtif0 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac0isrType) GetBtif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0isrFieldBtif0Mask) != 0
}

// SetBtif0 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac0isrType) SetBtif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0isrFieldBtif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0isrFieldBtif0Mask)
	}
}

const (
	RegisterMdmac0isrFieldTcif0Shift = 4
	RegisterMdmac0isrFieldTcif0Mask  = 0x10
)

// GetTcif0 channel x buffer transfer complete
func (r *RegisterMdmac0isrType) GetTcif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0isrFieldTcif0Mask) != 0
}

// SetTcif0 channel x buffer transfer complete
func (r *RegisterMdmac0isrType) SetTcif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0isrFieldTcif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0isrFieldTcif0Mask)
	}
}

const (
	RegisterMdmac0isrFieldCrqa0Shift = 16
	RegisterMdmac0isrFieldCrqa0Mask  = 0x10000
)

// GetCrqa0 channel x request active flag
func (r *RegisterMdmac0isrType) GetCrqa0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0isrFieldCrqa0Mask) != 0
}

// SetCrqa0 channel x request active flag
func (r *RegisterMdmac0isrType) SetCrqa0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0isrFieldCrqa0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0isrFieldCrqa0Mask)
	}
}

// RegisterMdmac0ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac0ifcrType uint32

func (r *RegisterMdmac0ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0ifcrFieldCteif0Shift = 0
	RegisterMdmac0ifcrFieldCteif0Mask  = 0x1
)

// GetCteif0 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac0ifcrType) GetCteif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0ifcrFieldCteif0Mask) != 0
}

// SetCteif0 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac0ifcrType) SetCteif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0ifcrFieldCteif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0ifcrFieldCteif0Mask)
	}
}

const (
	RegisterMdmac0ifcrFieldCctcif0Shift = 1
	RegisterMdmac0ifcrFieldCctcif0Mask  = 0x2
)

// GetCctcif0 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac0ifcrType) GetCctcif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0ifcrFieldCctcif0Mask) != 0
}

// SetCctcif0 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac0ifcrType) SetCctcif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0ifcrFieldCctcif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0ifcrFieldCctcif0Mask)
	}
}

const (
	RegisterMdmac0ifcrFieldCbrtif0Shift = 2
	RegisterMdmac0ifcrFieldCbrtif0Mask  = 0x4
)

// GetCbrtif0 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac0ifcrType) GetCbrtif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0ifcrFieldCbrtif0Mask) != 0
}

// SetCbrtif0 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac0ifcrType) SetCbrtif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0ifcrFieldCbrtif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0ifcrFieldCbrtif0Mask)
	}
}

const (
	RegisterMdmac0ifcrFieldCbtif0Shift = 3
	RegisterMdmac0ifcrFieldCbtif0Mask  = 0x8
)

// GetCbtif0 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac0ifcrType) GetCbtif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0ifcrFieldCbtif0Mask) != 0
}

// SetCbtif0 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac0ifcrType) SetCbtif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0ifcrFieldCbtif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0ifcrFieldCbtif0Mask)
	}
}

const (
	RegisterMdmac0ifcrFieldCltcif0Shift = 4
	RegisterMdmac0ifcrFieldCltcif0Mask  = 0x10
)

// GetCltcif0 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac0ifcrType) GetCltcif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0ifcrFieldCltcif0Mask) != 0
}

// SetCltcif0 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac0ifcrType) SetCltcif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0ifcrFieldCltcif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0ifcrFieldCltcif0Mask)
	}
}

// RegisterMdmac0esrType MDMA Channel x error status register
type RegisterMdmac0esrType uint32

func (r *RegisterMdmac0esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0esrFieldTeaShift = 0
	RegisterMdmac0esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac0esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0esrFieldTeaMask) >> RegisterMdmac0esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac0esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0esrFieldTeaMask)|(uint32(value)<<RegisterMdmac0esrFieldTeaShift))
}

const (
	RegisterMdmac0esrFieldTedShift = 7
	RegisterMdmac0esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac0esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac0esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0esrFieldTedMask)
	}
}

const (
	RegisterMdmac0esrFieldTeldShift = 8
	RegisterMdmac0esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac0esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac0esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0esrFieldTeldMask)
	}
}

const (
	RegisterMdmac0esrFieldTemdShift = 9
	RegisterMdmac0esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac0esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac0esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0esrFieldTemdMask)
	}
}

const (
	RegisterMdmac0esrFieldAseShift = 10
	RegisterMdmac0esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac0esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac0esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0esrFieldAseMask)
	}
}

const (
	RegisterMdmac0esrFieldBseShift = 11
	RegisterMdmac0esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac0esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac0esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0esrFieldBseMask)
	}
}

// RegisterMdmac0crType This register is used to control the concerned channel.
type RegisterMdmac0crType uint32

func (r *RegisterMdmac0crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0crFieldEnShift = 0
	RegisterMdmac0crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac0crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac0crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0crFieldEnMask)
	}
}

const (
	RegisterMdmac0crFieldTeieShift = 1
	RegisterMdmac0crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac0crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac0crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0crFieldTeieMask)
	}
}

const (
	RegisterMdmac0crFieldCtcieShift = 2
	RegisterMdmac0crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac0crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac0crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0crFieldCtcieMask)
	}
}

const (
	RegisterMdmac0crFieldBrtieShift = 3
	RegisterMdmac0crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac0crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac0crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0crFieldBrtieMask)
	}
}

const (
	RegisterMdmac0crFieldBtieShift = 4
	RegisterMdmac0crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac0crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac0crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0crFieldBtieMask)
	}
}

const (
	RegisterMdmac0crFieldTcieShift = 5
	RegisterMdmac0crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac0crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac0crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0crFieldTcieMask)
	}
}

const (
	RegisterMdmac0crFieldPlShift = 6
	RegisterMdmac0crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac0crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0crFieldPlMask) >> RegisterMdmac0crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac0crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0crFieldPlMask)|(uint32(value)<<RegisterMdmac0crFieldPlShift))
}

const (
	RegisterMdmac0crFieldBexShift = 12
	RegisterMdmac0crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac0crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac0crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0crFieldBexMask)
	}
}

const (
	RegisterMdmac0crFieldHexShift = 13
	RegisterMdmac0crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac0crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac0crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0crFieldHexMask)
	}
}

const (
	RegisterMdmac0crFieldWexShift = 14
	RegisterMdmac0crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac0crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac0crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0crFieldWexMask)
	}
}

const (
	RegisterMdmac0crFieldSwrqShift = 16
	RegisterMdmac0crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac0crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0crFieldSwrqMask)
	}
}

// RegisterMdmac0tcrType This register is used to configure the concerned channel.
type RegisterMdmac0tcrType uint32

func (r *RegisterMdmac0tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0tcrFieldSincShift = 0
	RegisterMdmac0tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac0tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldSincMask) >> RegisterMdmac0tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac0tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldSincMask)|(uint32(value)<<RegisterMdmac0tcrFieldSincShift))
}

const (
	RegisterMdmac0tcrFieldDincShift = 2
	RegisterMdmac0tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac0tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldDincMask) >> RegisterMdmac0tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac0tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldDincMask)|(uint32(value)<<RegisterMdmac0tcrFieldDincShift))
}

const (
	RegisterMdmac0tcrFieldSsizeShift = 4
	RegisterMdmac0tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac0tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldSsizeMask) >> RegisterMdmac0tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac0tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac0tcrFieldSsizeShift))
}

const (
	RegisterMdmac0tcrFieldDsizeShift = 6
	RegisterMdmac0tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac0tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldDsizeMask) >> RegisterMdmac0tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac0tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac0tcrFieldDsizeShift))
}

const (
	RegisterMdmac0tcrFieldSincosShift = 8
	RegisterMdmac0tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac0tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldSincosMask) >> RegisterMdmac0tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac0tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac0tcrFieldSincosShift))
}

const (
	RegisterMdmac0tcrFieldDincosShift = 10
	RegisterMdmac0tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac0tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldDincosMask) >> RegisterMdmac0tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac0tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac0tcrFieldDincosShift))
}

const (
	RegisterMdmac0tcrFieldSburstShift = 12
	RegisterMdmac0tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac0tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldSburstMask) >> RegisterMdmac0tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac0tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac0tcrFieldSburstShift))
}

const (
	RegisterMdmac0tcrFieldDburstShift = 15
	RegisterMdmac0tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac0tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldDburstMask) >> RegisterMdmac0tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac0tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac0tcrFieldDburstShift))
}

const (
	RegisterMdmac0tcrFieldTlenShift = 18
	RegisterMdmac0tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac0tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldTlenMask) >> RegisterMdmac0tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac0tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac0tcrFieldTlenShift))
}

const (
	RegisterMdmac0tcrFieldPkeShift = 25
	RegisterMdmac0tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac0tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac0tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac0tcrFieldPamShift = 26
	RegisterMdmac0tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac0tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldPamMask) >> RegisterMdmac0tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac0tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldPamMask)|(uint32(value)<<RegisterMdmac0tcrFieldPamShift))
}

const (
	RegisterMdmac0tcrFieldTrgmShift = 28
	RegisterMdmac0tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac0tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldTrgmMask) >> RegisterMdmac0tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac0tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac0tcrFieldTrgmShift))
}

const (
	RegisterMdmac0tcrFieldSwrmShift = 30
	RegisterMdmac0tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac0tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac0tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac0tcrFieldBwmShift = 31
	RegisterMdmac0tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac0tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac0tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tcrFieldBwmMask)
	}
}

// RegisterMdmac0bndtrType MDMA Channel x block number of data register
type RegisterMdmac0bndtrType uint32

func (r *RegisterMdmac0bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0bndtrFieldBndtShift = 0
	RegisterMdmac0bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac0bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0bndtrFieldBndtMask) >> RegisterMdmac0bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac0bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac0bndtrFieldBndtShift))
}

const (
	RegisterMdmac0bndtrFieldBrsumShift = 18
	RegisterMdmac0bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac0bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac0bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac0bndtrFieldBrdumShift = 19
	RegisterMdmac0bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac0bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac0bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac0bndtrFieldBrcShift = 20
	RegisterMdmac0bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac0bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0bndtrFieldBrcMask) >> RegisterMdmac0bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac0bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac0bndtrFieldBrcShift))
}

// RegisterMdmac0sarType MDMA channel x source address register
type RegisterMdmac0sarType uint32

func (r *RegisterMdmac0sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0sarFieldSarShift = 0
	RegisterMdmac0sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac0sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0sarFieldSarMask) >> RegisterMdmac0sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac0sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0sarFieldSarMask)|(uint32(value)<<RegisterMdmac0sarFieldSarShift))
}

// RegisterMdmac0darType MDMA channel x destination address register
type RegisterMdmac0darType uint32

func (r *RegisterMdmac0darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0darFieldDarShift = 0
	RegisterMdmac0darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac0darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0darFieldDarMask) >> RegisterMdmac0darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac0darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0darFieldDarMask)|(uint32(value)<<RegisterMdmac0darFieldDarShift))
}

// RegisterMdmac0brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac0brurType uint32

func (r *RegisterMdmac0brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0brurFieldSuvShift = 0
	RegisterMdmac0brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac0brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0brurFieldSuvMask) >> RegisterMdmac0brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac0brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0brurFieldSuvMask)|(uint32(value)<<RegisterMdmac0brurFieldSuvShift))
}

const (
	RegisterMdmac0brurFieldDuvShift = 16
	RegisterMdmac0brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac0brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0brurFieldDuvMask) >> RegisterMdmac0brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac0brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0brurFieldDuvMask)|(uint32(value)<<RegisterMdmac0brurFieldDuvShift))
}

// RegisterMdmac0larType MDMA channel x Link Address register
type RegisterMdmac0larType uint32

func (r *RegisterMdmac0larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0larFieldLarShift = 0
	RegisterMdmac0larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac0larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0larFieldLarMask) >> RegisterMdmac0larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac0larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0larFieldLarMask)|(uint32(value)<<RegisterMdmac0larFieldLarShift))
}

// RegisterMdmac0tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac0tbrType uint32

func (r *RegisterMdmac0tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0tbrFieldTselShift = 0
	RegisterMdmac0tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac0tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tbrFieldTselMask) >> RegisterMdmac0tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac0tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tbrFieldTselMask)|(uint32(value)<<RegisterMdmac0tbrFieldTselShift))
}

const (
	RegisterMdmac0tbrFieldSbusShift = 16
	RegisterMdmac0tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac0tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac0tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac0tbrFieldDbusShift = 17
	RegisterMdmac0tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac0tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac0tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac0tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0tbrFieldDbusMask)
	}
}

// RegisterMdmac0marType MDMA channel x Mask address register
type RegisterMdmac0marType uint32

func (r *RegisterMdmac0marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0marFieldMarShift = 0
	RegisterMdmac0marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac0marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0marFieldMarMask) >> RegisterMdmac0marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac0marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0marFieldMarMask)|(uint32(value)<<RegisterMdmac0marFieldMarShift))
}

// RegisterMdmac0mdrType MDMA channel x Mask Data register
type RegisterMdmac0mdrType uint32

func (r *RegisterMdmac0mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac0mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac0mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac0mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac0mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac0mdrFieldMdrShift = 0
	RegisterMdmac0mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac0mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac0mdrFieldMdrMask) >> RegisterMdmac0mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac0mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac0mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac0mdrFieldMdrShift))
}

// RegisterMdmac1isrType MDMA channel x interrupt/status register
type RegisterMdmac1isrType uint32

func (r *RegisterMdmac1isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1isrFieldTeif1Shift = 0
	RegisterMdmac1isrFieldTeif1Mask  = 0x1
)

// GetTeif1 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac1isrType) GetTeif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1isrFieldTeif1Mask) != 0
}

// SetTeif1 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac1isrType) SetTeif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1isrFieldTeif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1isrFieldTeif1Mask)
	}
}

const (
	RegisterMdmac1isrFieldCtcif1Shift = 1
	RegisterMdmac1isrFieldCtcif1Mask  = 0x2
)

// GetCtcif1 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac1isrType) GetCtcif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1isrFieldCtcif1Mask) != 0
}

// SetCtcif1 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac1isrType) SetCtcif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1isrFieldCtcif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1isrFieldCtcif1Mask)
	}
}

const (
	RegisterMdmac1isrFieldBrtif1Shift = 2
	RegisterMdmac1isrFieldBrtif1Mask  = 0x4
)

// GetBrtif1 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac1isrType) GetBrtif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1isrFieldBrtif1Mask) != 0
}

// SetBrtif1 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac1isrType) SetBrtif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1isrFieldBrtif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1isrFieldBrtif1Mask)
	}
}

const (
	RegisterMdmac1isrFieldBtif1Shift = 3
	RegisterMdmac1isrFieldBtif1Mask  = 0x8
)

// GetBtif1 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac1isrType) GetBtif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1isrFieldBtif1Mask) != 0
}

// SetBtif1 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac1isrType) SetBtif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1isrFieldBtif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1isrFieldBtif1Mask)
	}
}

const (
	RegisterMdmac1isrFieldTcif1Shift = 4
	RegisterMdmac1isrFieldTcif1Mask  = 0x10
)

// GetTcif1 channel x buffer transfer complete
func (r *RegisterMdmac1isrType) GetTcif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1isrFieldTcif1Mask) != 0
}

// SetTcif1 channel x buffer transfer complete
func (r *RegisterMdmac1isrType) SetTcif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1isrFieldTcif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1isrFieldTcif1Mask)
	}
}

const (
	RegisterMdmac1isrFieldCrqa1Shift = 16
	RegisterMdmac1isrFieldCrqa1Mask  = 0x10000
)

// GetCrqa1 channel x request active flag
func (r *RegisterMdmac1isrType) GetCrqa1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1isrFieldCrqa1Mask) != 0
}

// SetCrqa1 channel x request active flag
func (r *RegisterMdmac1isrType) SetCrqa1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1isrFieldCrqa1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1isrFieldCrqa1Mask)
	}
}

// RegisterMdmac1ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac1ifcrType uint32

func (r *RegisterMdmac1ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1ifcrFieldCteif1Shift = 0
	RegisterMdmac1ifcrFieldCteif1Mask  = 0x1
)

// GetCteif1 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac1ifcrType) GetCteif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1ifcrFieldCteif1Mask) != 0
}

// SetCteif1 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac1ifcrType) SetCteif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1ifcrFieldCteif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1ifcrFieldCteif1Mask)
	}
}

const (
	RegisterMdmac1ifcrFieldCctcif1Shift = 1
	RegisterMdmac1ifcrFieldCctcif1Mask  = 0x2
)

// GetCctcif1 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac1ifcrType) GetCctcif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1ifcrFieldCctcif1Mask) != 0
}

// SetCctcif1 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac1ifcrType) SetCctcif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1ifcrFieldCctcif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1ifcrFieldCctcif1Mask)
	}
}

const (
	RegisterMdmac1ifcrFieldCbrtif1Shift = 2
	RegisterMdmac1ifcrFieldCbrtif1Mask  = 0x4
)

// GetCbrtif1 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac1ifcrType) GetCbrtif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1ifcrFieldCbrtif1Mask) != 0
}

// SetCbrtif1 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac1ifcrType) SetCbrtif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1ifcrFieldCbrtif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1ifcrFieldCbrtif1Mask)
	}
}

const (
	RegisterMdmac1ifcrFieldCbtif1Shift = 3
	RegisterMdmac1ifcrFieldCbtif1Mask  = 0x8
)

// GetCbtif1 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac1ifcrType) GetCbtif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1ifcrFieldCbtif1Mask) != 0
}

// SetCbtif1 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac1ifcrType) SetCbtif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1ifcrFieldCbtif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1ifcrFieldCbtif1Mask)
	}
}

const (
	RegisterMdmac1ifcrFieldCltcif1Shift = 4
	RegisterMdmac1ifcrFieldCltcif1Mask  = 0x10
)

// GetCltcif1 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac1ifcrType) GetCltcif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1ifcrFieldCltcif1Mask) != 0
}

// SetCltcif1 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac1ifcrType) SetCltcif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1ifcrFieldCltcif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1ifcrFieldCltcif1Mask)
	}
}

// RegisterMdmac1esrType MDMA Channel x error status register
type RegisterMdmac1esrType uint32

func (r *RegisterMdmac1esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1esrFieldTeaShift = 0
	RegisterMdmac1esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac1esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1esrFieldTeaMask) >> RegisterMdmac1esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac1esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1esrFieldTeaMask)|(uint32(value)<<RegisterMdmac1esrFieldTeaShift))
}

const (
	RegisterMdmac1esrFieldTedShift = 7
	RegisterMdmac1esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac1esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac1esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1esrFieldTedMask)
	}
}

const (
	RegisterMdmac1esrFieldTeldShift = 8
	RegisterMdmac1esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac1esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac1esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1esrFieldTeldMask)
	}
}

const (
	RegisterMdmac1esrFieldTemdShift = 9
	RegisterMdmac1esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac1esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac1esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1esrFieldTemdMask)
	}
}

const (
	RegisterMdmac1esrFieldAseShift = 10
	RegisterMdmac1esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac1esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac1esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1esrFieldAseMask)
	}
}

const (
	RegisterMdmac1esrFieldBseShift = 11
	RegisterMdmac1esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac1esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac1esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1esrFieldBseMask)
	}
}

// RegisterMdmac1crType This register is used to control the concerned channel.
type RegisterMdmac1crType uint32

func (r *RegisterMdmac1crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1crFieldEnShift = 0
	RegisterMdmac1crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac1crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac1crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1crFieldEnMask)
	}
}

const (
	RegisterMdmac1crFieldTeieShift = 1
	RegisterMdmac1crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac1crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac1crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1crFieldTeieMask)
	}
}

const (
	RegisterMdmac1crFieldCtcieShift = 2
	RegisterMdmac1crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac1crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac1crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1crFieldCtcieMask)
	}
}

const (
	RegisterMdmac1crFieldBrtieShift = 3
	RegisterMdmac1crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac1crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac1crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1crFieldBrtieMask)
	}
}

const (
	RegisterMdmac1crFieldBtieShift = 4
	RegisterMdmac1crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac1crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac1crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1crFieldBtieMask)
	}
}

const (
	RegisterMdmac1crFieldTcieShift = 5
	RegisterMdmac1crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac1crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac1crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1crFieldTcieMask)
	}
}

const (
	RegisterMdmac1crFieldPlShift = 6
	RegisterMdmac1crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac1crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1crFieldPlMask) >> RegisterMdmac1crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac1crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1crFieldPlMask)|(uint32(value)<<RegisterMdmac1crFieldPlShift))
}

const (
	RegisterMdmac1crFieldBexShift = 12
	RegisterMdmac1crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac1crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac1crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1crFieldBexMask)
	}
}

const (
	RegisterMdmac1crFieldHexShift = 13
	RegisterMdmac1crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac1crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac1crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1crFieldHexMask)
	}
}

const (
	RegisterMdmac1crFieldWexShift = 14
	RegisterMdmac1crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac1crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac1crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1crFieldWexMask)
	}
}

const (
	RegisterMdmac1crFieldSwrqShift = 16
	RegisterMdmac1crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac1crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1crFieldSwrqMask)
	}
}

// RegisterMdmac1tcrType This register is used to configure the concerned channel.
type RegisterMdmac1tcrType uint32

func (r *RegisterMdmac1tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1tcrFieldSincShift = 0
	RegisterMdmac1tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac1tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldSincMask) >> RegisterMdmac1tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac1tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldSincMask)|(uint32(value)<<RegisterMdmac1tcrFieldSincShift))
}

const (
	RegisterMdmac1tcrFieldDincShift = 2
	RegisterMdmac1tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac1tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldDincMask) >> RegisterMdmac1tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac1tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldDincMask)|(uint32(value)<<RegisterMdmac1tcrFieldDincShift))
}

const (
	RegisterMdmac1tcrFieldSsizeShift = 4
	RegisterMdmac1tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac1tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldSsizeMask) >> RegisterMdmac1tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac1tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac1tcrFieldSsizeShift))
}

const (
	RegisterMdmac1tcrFieldDsizeShift = 6
	RegisterMdmac1tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac1tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldDsizeMask) >> RegisterMdmac1tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac1tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac1tcrFieldDsizeShift))
}

const (
	RegisterMdmac1tcrFieldSincosShift = 8
	RegisterMdmac1tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac1tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldSincosMask) >> RegisterMdmac1tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac1tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac1tcrFieldSincosShift))
}

const (
	RegisterMdmac1tcrFieldDincosShift = 10
	RegisterMdmac1tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac1tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldDincosMask) >> RegisterMdmac1tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac1tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac1tcrFieldDincosShift))
}

const (
	RegisterMdmac1tcrFieldSburstShift = 12
	RegisterMdmac1tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac1tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldSburstMask) >> RegisterMdmac1tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac1tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac1tcrFieldSburstShift))
}

const (
	RegisterMdmac1tcrFieldDburstShift = 15
	RegisterMdmac1tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac1tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldDburstMask) >> RegisterMdmac1tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac1tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac1tcrFieldDburstShift))
}

const (
	RegisterMdmac1tcrFieldTlenShift = 18
	RegisterMdmac1tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac1tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldTlenMask) >> RegisterMdmac1tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac1tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac1tcrFieldTlenShift))
}

const (
	RegisterMdmac1tcrFieldPkeShift = 25
	RegisterMdmac1tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac1tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac1tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac1tcrFieldPamShift = 26
	RegisterMdmac1tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac1tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldPamMask) >> RegisterMdmac1tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac1tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldPamMask)|(uint32(value)<<RegisterMdmac1tcrFieldPamShift))
}

const (
	RegisterMdmac1tcrFieldTrgmShift = 28
	RegisterMdmac1tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac1tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldTrgmMask) >> RegisterMdmac1tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac1tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac1tcrFieldTrgmShift))
}

const (
	RegisterMdmac1tcrFieldSwrmShift = 30
	RegisterMdmac1tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac1tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac1tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac1tcrFieldBwmShift = 31
	RegisterMdmac1tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac1tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac1tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tcrFieldBwmMask)
	}
}

// RegisterMdmac1bndtrType MDMA Channel x block number of data register
type RegisterMdmac1bndtrType uint32

func (r *RegisterMdmac1bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1bndtrFieldBndtShift = 0
	RegisterMdmac1bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac1bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1bndtrFieldBndtMask) >> RegisterMdmac1bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac1bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac1bndtrFieldBndtShift))
}

const (
	RegisterMdmac1bndtrFieldBrsumShift = 18
	RegisterMdmac1bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac1bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac1bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac1bndtrFieldBrdumShift = 19
	RegisterMdmac1bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac1bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac1bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac1bndtrFieldBrcShift = 20
	RegisterMdmac1bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac1bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1bndtrFieldBrcMask) >> RegisterMdmac1bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac1bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac1bndtrFieldBrcShift))
}

// RegisterMdmac1sarType MDMA channel x source address register
type RegisterMdmac1sarType uint32

func (r *RegisterMdmac1sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1sarFieldSarShift = 0
	RegisterMdmac1sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac1sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1sarFieldSarMask) >> RegisterMdmac1sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac1sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1sarFieldSarMask)|(uint32(value)<<RegisterMdmac1sarFieldSarShift))
}

// RegisterMdmac1darType MDMA channel x destination address register
type RegisterMdmac1darType uint32

func (r *RegisterMdmac1darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1darFieldDarShift = 0
	RegisterMdmac1darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac1darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1darFieldDarMask) >> RegisterMdmac1darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac1darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1darFieldDarMask)|(uint32(value)<<RegisterMdmac1darFieldDarShift))
}

// RegisterMdmac1brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac1brurType uint32

func (r *RegisterMdmac1brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1brurFieldSuvShift = 0
	RegisterMdmac1brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac1brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1brurFieldSuvMask) >> RegisterMdmac1brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac1brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1brurFieldSuvMask)|(uint32(value)<<RegisterMdmac1brurFieldSuvShift))
}

const (
	RegisterMdmac1brurFieldDuvShift = 16
	RegisterMdmac1brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac1brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1brurFieldDuvMask) >> RegisterMdmac1brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac1brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1brurFieldDuvMask)|(uint32(value)<<RegisterMdmac1brurFieldDuvShift))
}

// RegisterMdmac1larType MDMA channel x Link Address register
type RegisterMdmac1larType uint32

func (r *RegisterMdmac1larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1larFieldLarShift = 0
	RegisterMdmac1larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac1larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1larFieldLarMask) >> RegisterMdmac1larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac1larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1larFieldLarMask)|(uint32(value)<<RegisterMdmac1larFieldLarShift))
}

// RegisterMdmac1tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac1tbrType uint32

func (r *RegisterMdmac1tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1tbrFieldTselShift = 0
	RegisterMdmac1tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac1tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tbrFieldTselMask) >> RegisterMdmac1tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac1tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tbrFieldTselMask)|(uint32(value)<<RegisterMdmac1tbrFieldTselShift))
}

const (
	RegisterMdmac1tbrFieldSbusShift = 16
	RegisterMdmac1tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac1tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac1tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac1tbrFieldDbusShift = 17
	RegisterMdmac1tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac1tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac1tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac1tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1tbrFieldDbusMask)
	}
}

// RegisterMdmac1marType MDMA channel x Mask address register
type RegisterMdmac1marType uint32

func (r *RegisterMdmac1marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1marFieldMarShift = 0
	RegisterMdmac1marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac1marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1marFieldMarMask) >> RegisterMdmac1marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac1marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1marFieldMarMask)|(uint32(value)<<RegisterMdmac1marFieldMarShift))
}

// RegisterMdmac1mdrType MDMA channel x Mask Data register
type RegisterMdmac1mdrType uint32

func (r *RegisterMdmac1mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac1mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac1mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac1mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac1mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac1mdrFieldMdrShift = 0
	RegisterMdmac1mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac1mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac1mdrFieldMdrMask) >> RegisterMdmac1mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac1mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac1mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac1mdrFieldMdrShift))
}

// RegisterMdmac2isrType MDMA channel x interrupt/status register
type RegisterMdmac2isrType uint32

func (r *RegisterMdmac2isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2isrFieldTeif2Shift = 0
	RegisterMdmac2isrFieldTeif2Mask  = 0x1
)

// GetTeif2 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac2isrType) GetTeif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2isrFieldTeif2Mask) != 0
}

// SetTeif2 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac2isrType) SetTeif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2isrFieldTeif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2isrFieldTeif2Mask)
	}
}

const (
	RegisterMdmac2isrFieldCtcif2Shift = 1
	RegisterMdmac2isrFieldCtcif2Mask  = 0x2
)

// GetCtcif2 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac2isrType) GetCtcif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2isrFieldCtcif2Mask) != 0
}

// SetCtcif2 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac2isrType) SetCtcif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2isrFieldCtcif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2isrFieldCtcif2Mask)
	}
}

const (
	RegisterMdmac2isrFieldBrtif2Shift = 2
	RegisterMdmac2isrFieldBrtif2Mask  = 0x4
)

// GetBrtif2 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac2isrType) GetBrtif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2isrFieldBrtif2Mask) != 0
}

// SetBrtif2 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac2isrType) SetBrtif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2isrFieldBrtif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2isrFieldBrtif2Mask)
	}
}

const (
	RegisterMdmac2isrFieldBtif2Shift = 3
	RegisterMdmac2isrFieldBtif2Mask  = 0x8
)

// GetBtif2 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac2isrType) GetBtif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2isrFieldBtif2Mask) != 0
}

// SetBtif2 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac2isrType) SetBtif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2isrFieldBtif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2isrFieldBtif2Mask)
	}
}

const (
	RegisterMdmac2isrFieldTcif2Shift = 4
	RegisterMdmac2isrFieldTcif2Mask  = 0x10
)

// GetTcif2 channel x buffer transfer complete
func (r *RegisterMdmac2isrType) GetTcif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2isrFieldTcif2Mask) != 0
}

// SetTcif2 channel x buffer transfer complete
func (r *RegisterMdmac2isrType) SetTcif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2isrFieldTcif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2isrFieldTcif2Mask)
	}
}

const (
	RegisterMdmac2isrFieldCrqa2Shift = 16
	RegisterMdmac2isrFieldCrqa2Mask  = 0x10000
)

// GetCrqa2 channel x request active flag
func (r *RegisterMdmac2isrType) GetCrqa2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2isrFieldCrqa2Mask) != 0
}

// SetCrqa2 channel x request active flag
func (r *RegisterMdmac2isrType) SetCrqa2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2isrFieldCrqa2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2isrFieldCrqa2Mask)
	}
}

// RegisterMdmac2ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac2ifcrType uint32

func (r *RegisterMdmac2ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2ifcrFieldCteif2Shift = 0
	RegisterMdmac2ifcrFieldCteif2Mask  = 0x1
)

// GetCteif2 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac2ifcrType) GetCteif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2ifcrFieldCteif2Mask) != 0
}

// SetCteif2 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac2ifcrType) SetCteif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2ifcrFieldCteif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2ifcrFieldCteif2Mask)
	}
}

const (
	RegisterMdmac2ifcrFieldCctcif2Shift = 1
	RegisterMdmac2ifcrFieldCctcif2Mask  = 0x2
)

// GetCctcif2 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac2ifcrType) GetCctcif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2ifcrFieldCctcif2Mask) != 0
}

// SetCctcif2 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac2ifcrType) SetCctcif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2ifcrFieldCctcif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2ifcrFieldCctcif2Mask)
	}
}

const (
	RegisterMdmac2ifcrFieldCbrtif2Shift = 2
	RegisterMdmac2ifcrFieldCbrtif2Mask  = 0x4
)

// GetCbrtif2 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac2ifcrType) GetCbrtif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2ifcrFieldCbrtif2Mask) != 0
}

// SetCbrtif2 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac2ifcrType) SetCbrtif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2ifcrFieldCbrtif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2ifcrFieldCbrtif2Mask)
	}
}

const (
	RegisterMdmac2ifcrFieldCbtif2Shift = 3
	RegisterMdmac2ifcrFieldCbtif2Mask  = 0x8
)

// GetCbtif2 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac2ifcrType) GetCbtif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2ifcrFieldCbtif2Mask) != 0
}

// SetCbtif2 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac2ifcrType) SetCbtif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2ifcrFieldCbtif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2ifcrFieldCbtif2Mask)
	}
}

const (
	RegisterMdmac2ifcrFieldCltcif2Shift = 4
	RegisterMdmac2ifcrFieldCltcif2Mask  = 0x10
)

// GetCltcif2 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac2ifcrType) GetCltcif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2ifcrFieldCltcif2Mask) != 0
}

// SetCltcif2 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac2ifcrType) SetCltcif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2ifcrFieldCltcif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2ifcrFieldCltcif2Mask)
	}
}

// RegisterMdmac2esrType MDMA Channel x error status register
type RegisterMdmac2esrType uint32

func (r *RegisterMdmac2esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2esrFieldTeaShift = 0
	RegisterMdmac2esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac2esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2esrFieldTeaMask) >> RegisterMdmac2esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac2esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2esrFieldTeaMask)|(uint32(value)<<RegisterMdmac2esrFieldTeaShift))
}

const (
	RegisterMdmac2esrFieldTedShift = 7
	RegisterMdmac2esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac2esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac2esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2esrFieldTedMask)
	}
}

const (
	RegisterMdmac2esrFieldTeldShift = 8
	RegisterMdmac2esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac2esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac2esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2esrFieldTeldMask)
	}
}

const (
	RegisterMdmac2esrFieldTemdShift = 9
	RegisterMdmac2esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac2esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac2esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2esrFieldTemdMask)
	}
}

const (
	RegisterMdmac2esrFieldAseShift = 10
	RegisterMdmac2esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac2esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac2esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2esrFieldAseMask)
	}
}

const (
	RegisterMdmac2esrFieldBseShift = 11
	RegisterMdmac2esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac2esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac2esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2esrFieldBseMask)
	}
}

// RegisterMdmac2crType This register is used to control the concerned channel.
type RegisterMdmac2crType uint32

func (r *RegisterMdmac2crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2crFieldEnShift = 0
	RegisterMdmac2crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac2crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac2crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2crFieldEnMask)
	}
}

const (
	RegisterMdmac2crFieldTeieShift = 1
	RegisterMdmac2crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac2crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac2crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2crFieldTeieMask)
	}
}

const (
	RegisterMdmac2crFieldCtcieShift = 2
	RegisterMdmac2crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac2crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac2crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2crFieldCtcieMask)
	}
}

const (
	RegisterMdmac2crFieldBrtieShift = 3
	RegisterMdmac2crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac2crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac2crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2crFieldBrtieMask)
	}
}

const (
	RegisterMdmac2crFieldBtieShift = 4
	RegisterMdmac2crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac2crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac2crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2crFieldBtieMask)
	}
}

const (
	RegisterMdmac2crFieldTcieShift = 5
	RegisterMdmac2crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac2crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac2crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2crFieldTcieMask)
	}
}

const (
	RegisterMdmac2crFieldPlShift = 6
	RegisterMdmac2crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac2crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2crFieldPlMask) >> RegisterMdmac2crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac2crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2crFieldPlMask)|(uint32(value)<<RegisterMdmac2crFieldPlShift))
}

const (
	RegisterMdmac2crFieldBexShift = 12
	RegisterMdmac2crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac2crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac2crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2crFieldBexMask)
	}
}

const (
	RegisterMdmac2crFieldHexShift = 13
	RegisterMdmac2crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac2crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac2crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2crFieldHexMask)
	}
}

const (
	RegisterMdmac2crFieldWexShift = 14
	RegisterMdmac2crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac2crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac2crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2crFieldWexMask)
	}
}

const (
	RegisterMdmac2crFieldSwrqShift = 16
	RegisterMdmac2crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac2crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2crFieldSwrqMask)
	}
}

// RegisterMdmac2tcrType This register is used to configure the concerned channel.
type RegisterMdmac2tcrType uint32

func (r *RegisterMdmac2tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2tcrFieldSincShift = 0
	RegisterMdmac2tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac2tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldSincMask) >> RegisterMdmac2tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac2tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldSincMask)|(uint32(value)<<RegisterMdmac2tcrFieldSincShift))
}

const (
	RegisterMdmac2tcrFieldDincShift = 2
	RegisterMdmac2tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac2tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldDincMask) >> RegisterMdmac2tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac2tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldDincMask)|(uint32(value)<<RegisterMdmac2tcrFieldDincShift))
}

const (
	RegisterMdmac2tcrFieldSsizeShift = 4
	RegisterMdmac2tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac2tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldSsizeMask) >> RegisterMdmac2tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac2tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac2tcrFieldSsizeShift))
}

const (
	RegisterMdmac2tcrFieldDsizeShift = 6
	RegisterMdmac2tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac2tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldDsizeMask) >> RegisterMdmac2tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac2tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac2tcrFieldDsizeShift))
}

const (
	RegisterMdmac2tcrFieldSincosShift = 8
	RegisterMdmac2tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac2tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldSincosMask) >> RegisterMdmac2tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac2tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac2tcrFieldSincosShift))
}

const (
	RegisterMdmac2tcrFieldDincosShift = 10
	RegisterMdmac2tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac2tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldDincosMask) >> RegisterMdmac2tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac2tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac2tcrFieldDincosShift))
}

const (
	RegisterMdmac2tcrFieldSburstShift = 12
	RegisterMdmac2tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac2tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldSburstMask) >> RegisterMdmac2tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac2tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac2tcrFieldSburstShift))
}

const (
	RegisterMdmac2tcrFieldDburstShift = 15
	RegisterMdmac2tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac2tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldDburstMask) >> RegisterMdmac2tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac2tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac2tcrFieldDburstShift))
}

const (
	RegisterMdmac2tcrFieldTlenShift = 18
	RegisterMdmac2tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac2tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldTlenMask) >> RegisterMdmac2tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac2tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac2tcrFieldTlenShift))
}

const (
	RegisterMdmac2tcrFieldPkeShift = 25
	RegisterMdmac2tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac2tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac2tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac2tcrFieldPamShift = 26
	RegisterMdmac2tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac2tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldPamMask) >> RegisterMdmac2tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac2tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldPamMask)|(uint32(value)<<RegisterMdmac2tcrFieldPamShift))
}

const (
	RegisterMdmac2tcrFieldTrgmShift = 28
	RegisterMdmac2tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac2tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldTrgmMask) >> RegisterMdmac2tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac2tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac2tcrFieldTrgmShift))
}

const (
	RegisterMdmac2tcrFieldSwrmShift = 30
	RegisterMdmac2tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac2tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac2tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac2tcrFieldBwmShift = 31
	RegisterMdmac2tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac2tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac2tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tcrFieldBwmMask)
	}
}

// RegisterMdmac2bndtrType MDMA Channel x block number of data register
type RegisterMdmac2bndtrType uint32

func (r *RegisterMdmac2bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2bndtrFieldBndtShift = 0
	RegisterMdmac2bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac2bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2bndtrFieldBndtMask) >> RegisterMdmac2bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac2bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac2bndtrFieldBndtShift))
}

const (
	RegisterMdmac2bndtrFieldBrsumShift = 18
	RegisterMdmac2bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac2bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac2bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac2bndtrFieldBrdumShift = 19
	RegisterMdmac2bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac2bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac2bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac2bndtrFieldBrcShift = 20
	RegisterMdmac2bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac2bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2bndtrFieldBrcMask) >> RegisterMdmac2bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac2bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac2bndtrFieldBrcShift))
}

// RegisterMdmac2sarType MDMA channel x source address register
type RegisterMdmac2sarType uint32

func (r *RegisterMdmac2sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2sarFieldSarShift = 0
	RegisterMdmac2sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac2sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2sarFieldSarMask) >> RegisterMdmac2sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac2sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2sarFieldSarMask)|(uint32(value)<<RegisterMdmac2sarFieldSarShift))
}

// RegisterMdmac2darType MDMA channel x destination address register
type RegisterMdmac2darType uint32

func (r *RegisterMdmac2darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2darFieldDarShift = 0
	RegisterMdmac2darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac2darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2darFieldDarMask) >> RegisterMdmac2darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac2darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2darFieldDarMask)|(uint32(value)<<RegisterMdmac2darFieldDarShift))
}

// RegisterMdmac2brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac2brurType uint32

func (r *RegisterMdmac2brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2brurFieldSuvShift = 0
	RegisterMdmac2brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac2brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2brurFieldSuvMask) >> RegisterMdmac2brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac2brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2brurFieldSuvMask)|(uint32(value)<<RegisterMdmac2brurFieldSuvShift))
}

const (
	RegisterMdmac2brurFieldDuvShift = 16
	RegisterMdmac2brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac2brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2brurFieldDuvMask) >> RegisterMdmac2brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac2brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2brurFieldDuvMask)|(uint32(value)<<RegisterMdmac2brurFieldDuvShift))
}

// RegisterMdmac2larType MDMA channel x Link Address register
type RegisterMdmac2larType uint32

func (r *RegisterMdmac2larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2larFieldLarShift = 0
	RegisterMdmac2larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac2larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2larFieldLarMask) >> RegisterMdmac2larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac2larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2larFieldLarMask)|(uint32(value)<<RegisterMdmac2larFieldLarShift))
}

// RegisterMdmac2tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac2tbrType uint32

func (r *RegisterMdmac2tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2tbrFieldTselShift = 0
	RegisterMdmac2tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac2tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tbrFieldTselMask) >> RegisterMdmac2tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac2tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tbrFieldTselMask)|(uint32(value)<<RegisterMdmac2tbrFieldTselShift))
}

const (
	RegisterMdmac2tbrFieldSbusShift = 16
	RegisterMdmac2tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac2tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac2tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac2tbrFieldDbusShift = 17
	RegisterMdmac2tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac2tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac2tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac2tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2tbrFieldDbusMask)
	}
}

// RegisterMdmac2marType MDMA channel x Mask address register
type RegisterMdmac2marType uint32

func (r *RegisterMdmac2marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2marFieldMarShift = 0
	RegisterMdmac2marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac2marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2marFieldMarMask) >> RegisterMdmac2marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac2marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2marFieldMarMask)|(uint32(value)<<RegisterMdmac2marFieldMarShift))
}

// RegisterMdmac2mdrType MDMA channel x Mask Data register
type RegisterMdmac2mdrType uint32

func (r *RegisterMdmac2mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac2mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac2mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac2mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac2mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac2mdrFieldMdrShift = 0
	RegisterMdmac2mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac2mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac2mdrFieldMdrMask) >> RegisterMdmac2mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac2mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac2mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac2mdrFieldMdrShift))
}

// RegisterMdmac3isrType MDMA channel x interrupt/status register
type RegisterMdmac3isrType uint32

func (r *RegisterMdmac3isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3isrFieldTeif3Shift = 0
	RegisterMdmac3isrFieldTeif3Mask  = 0x1
)

// GetTeif3 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac3isrType) GetTeif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3isrFieldTeif3Mask) != 0
}

// SetTeif3 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac3isrType) SetTeif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3isrFieldTeif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3isrFieldTeif3Mask)
	}
}

const (
	RegisterMdmac3isrFieldCtcif3Shift = 1
	RegisterMdmac3isrFieldCtcif3Mask  = 0x2
)

// GetCtcif3 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac3isrType) GetCtcif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3isrFieldCtcif3Mask) != 0
}

// SetCtcif3 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac3isrType) SetCtcif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3isrFieldCtcif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3isrFieldCtcif3Mask)
	}
}

const (
	RegisterMdmac3isrFieldBrtif3Shift = 2
	RegisterMdmac3isrFieldBrtif3Mask  = 0x4
)

// GetBrtif3 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac3isrType) GetBrtif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3isrFieldBrtif3Mask) != 0
}

// SetBrtif3 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac3isrType) SetBrtif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3isrFieldBrtif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3isrFieldBrtif3Mask)
	}
}

const (
	RegisterMdmac3isrFieldBtif3Shift = 3
	RegisterMdmac3isrFieldBtif3Mask  = 0x8
)

// GetBtif3 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac3isrType) GetBtif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3isrFieldBtif3Mask) != 0
}

// SetBtif3 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac3isrType) SetBtif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3isrFieldBtif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3isrFieldBtif3Mask)
	}
}

const (
	RegisterMdmac3isrFieldTcif3Shift = 4
	RegisterMdmac3isrFieldTcif3Mask  = 0x10
)

// GetTcif3 channel x buffer transfer complete
func (r *RegisterMdmac3isrType) GetTcif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3isrFieldTcif3Mask) != 0
}

// SetTcif3 channel x buffer transfer complete
func (r *RegisterMdmac3isrType) SetTcif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3isrFieldTcif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3isrFieldTcif3Mask)
	}
}

const (
	RegisterMdmac3isrFieldCrqa3Shift = 16
	RegisterMdmac3isrFieldCrqa3Mask  = 0x10000
)

// GetCrqa3 channel x request active flag
func (r *RegisterMdmac3isrType) GetCrqa3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3isrFieldCrqa3Mask) != 0
}

// SetCrqa3 channel x request active flag
func (r *RegisterMdmac3isrType) SetCrqa3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3isrFieldCrqa3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3isrFieldCrqa3Mask)
	}
}

// RegisterMdmac3ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac3ifcrType uint32

func (r *RegisterMdmac3ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3ifcrFieldCteif3Shift = 0
	RegisterMdmac3ifcrFieldCteif3Mask  = 0x1
)

// GetCteif3 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac3ifcrType) GetCteif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3ifcrFieldCteif3Mask) != 0
}

// SetCteif3 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac3ifcrType) SetCteif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3ifcrFieldCteif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3ifcrFieldCteif3Mask)
	}
}

const (
	RegisterMdmac3ifcrFieldCctcif3Shift = 1
	RegisterMdmac3ifcrFieldCctcif3Mask  = 0x2
)

// GetCctcif3 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac3ifcrType) GetCctcif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3ifcrFieldCctcif3Mask) != 0
}

// SetCctcif3 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac3ifcrType) SetCctcif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3ifcrFieldCctcif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3ifcrFieldCctcif3Mask)
	}
}

const (
	RegisterMdmac3ifcrFieldCbrtif3Shift = 2
	RegisterMdmac3ifcrFieldCbrtif3Mask  = 0x4
)

// GetCbrtif3 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac3ifcrType) GetCbrtif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3ifcrFieldCbrtif3Mask) != 0
}

// SetCbrtif3 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac3ifcrType) SetCbrtif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3ifcrFieldCbrtif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3ifcrFieldCbrtif3Mask)
	}
}

const (
	RegisterMdmac3ifcrFieldCbtif3Shift = 3
	RegisterMdmac3ifcrFieldCbtif3Mask  = 0x8
)

// GetCbtif3 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac3ifcrType) GetCbtif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3ifcrFieldCbtif3Mask) != 0
}

// SetCbtif3 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac3ifcrType) SetCbtif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3ifcrFieldCbtif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3ifcrFieldCbtif3Mask)
	}
}

const (
	RegisterMdmac3ifcrFieldCltcif3Shift = 4
	RegisterMdmac3ifcrFieldCltcif3Mask  = 0x10
)

// GetCltcif3 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac3ifcrType) GetCltcif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3ifcrFieldCltcif3Mask) != 0
}

// SetCltcif3 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac3ifcrType) SetCltcif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3ifcrFieldCltcif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3ifcrFieldCltcif3Mask)
	}
}

// RegisterMdmac3esrType MDMA Channel x error status register
type RegisterMdmac3esrType uint32

func (r *RegisterMdmac3esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3esrFieldTeaShift = 0
	RegisterMdmac3esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac3esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3esrFieldTeaMask) >> RegisterMdmac3esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac3esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3esrFieldTeaMask)|(uint32(value)<<RegisterMdmac3esrFieldTeaShift))
}

const (
	RegisterMdmac3esrFieldTedShift = 7
	RegisterMdmac3esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac3esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac3esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3esrFieldTedMask)
	}
}

const (
	RegisterMdmac3esrFieldTeldShift = 8
	RegisterMdmac3esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac3esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac3esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3esrFieldTeldMask)
	}
}

const (
	RegisterMdmac3esrFieldTemdShift = 9
	RegisterMdmac3esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac3esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac3esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3esrFieldTemdMask)
	}
}

const (
	RegisterMdmac3esrFieldAseShift = 10
	RegisterMdmac3esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac3esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac3esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3esrFieldAseMask)
	}
}

const (
	RegisterMdmac3esrFieldBseShift = 11
	RegisterMdmac3esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac3esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac3esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3esrFieldBseMask)
	}
}

// RegisterMdmac3crType This register is used to control the concerned channel.
type RegisterMdmac3crType uint32

func (r *RegisterMdmac3crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3crFieldEnShift = 0
	RegisterMdmac3crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac3crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac3crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3crFieldEnMask)
	}
}

const (
	RegisterMdmac3crFieldTeieShift = 1
	RegisterMdmac3crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac3crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac3crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3crFieldTeieMask)
	}
}

const (
	RegisterMdmac3crFieldCtcieShift = 2
	RegisterMdmac3crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac3crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac3crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3crFieldCtcieMask)
	}
}

const (
	RegisterMdmac3crFieldBrtieShift = 3
	RegisterMdmac3crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac3crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac3crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3crFieldBrtieMask)
	}
}

const (
	RegisterMdmac3crFieldBtieShift = 4
	RegisterMdmac3crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac3crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac3crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3crFieldBtieMask)
	}
}

const (
	RegisterMdmac3crFieldTcieShift = 5
	RegisterMdmac3crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac3crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac3crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3crFieldTcieMask)
	}
}

const (
	RegisterMdmac3crFieldPlShift = 6
	RegisterMdmac3crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac3crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3crFieldPlMask) >> RegisterMdmac3crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac3crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3crFieldPlMask)|(uint32(value)<<RegisterMdmac3crFieldPlShift))
}

const (
	RegisterMdmac3crFieldBexShift = 12
	RegisterMdmac3crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac3crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac3crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3crFieldBexMask)
	}
}

const (
	RegisterMdmac3crFieldHexShift = 13
	RegisterMdmac3crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac3crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac3crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3crFieldHexMask)
	}
}

const (
	RegisterMdmac3crFieldWexShift = 14
	RegisterMdmac3crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac3crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac3crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3crFieldWexMask)
	}
}

const (
	RegisterMdmac3crFieldSwrqShift = 16
	RegisterMdmac3crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac3crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3crFieldSwrqMask)
	}
}

// RegisterMdmac3tcrType This register is used to configure the concerned channel.
type RegisterMdmac3tcrType uint32

func (r *RegisterMdmac3tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3tcrFieldSincShift = 0
	RegisterMdmac3tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac3tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldSincMask) >> RegisterMdmac3tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac3tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldSincMask)|(uint32(value)<<RegisterMdmac3tcrFieldSincShift))
}

const (
	RegisterMdmac3tcrFieldDincShift = 2
	RegisterMdmac3tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac3tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldDincMask) >> RegisterMdmac3tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac3tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldDincMask)|(uint32(value)<<RegisterMdmac3tcrFieldDincShift))
}

const (
	RegisterMdmac3tcrFieldSsizeShift = 4
	RegisterMdmac3tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac3tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldSsizeMask) >> RegisterMdmac3tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac3tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac3tcrFieldSsizeShift))
}

const (
	RegisterMdmac3tcrFieldDsizeShift = 6
	RegisterMdmac3tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac3tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldDsizeMask) >> RegisterMdmac3tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac3tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac3tcrFieldDsizeShift))
}

const (
	RegisterMdmac3tcrFieldSincosShift = 8
	RegisterMdmac3tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac3tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldSincosMask) >> RegisterMdmac3tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac3tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac3tcrFieldSincosShift))
}

const (
	RegisterMdmac3tcrFieldDincosShift = 10
	RegisterMdmac3tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac3tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldDincosMask) >> RegisterMdmac3tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac3tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac3tcrFieldDincosShift))
}

const (
	RegisterMdmac3tcrFieldSburstShift = 12
	RegisterMdmac3tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac3tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldSburstMask) >> RegisterMdmac3tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac3tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac3tcrFieldSburstShift))
}

const (
	RegisterMdmac3tcrFieldDburstShift = 15
	RegisterMdmac3tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac3tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldDburstMask) >> RegisterMdmac3tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac3tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac3tcrFieldDburstShift))
}

const (
	RegisterMdmac3tcrFieldTlenShift = 18
	RegisterMdmac3tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac3tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldTlenMask) >> RegisterMdmac3tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac3tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac3tcrFieldTlenShift))
}

const (
	RegisterMdmac3tcrFieldPkeShift = 25
	RegisterMdmac3tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac3tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac3tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac3tcrFieldPamShift = 26
	RegisterMdmac3tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac3tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldPamMask) >> RegisterMdmac3tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac3tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldPamMask)|(uint32(value)<<RegisterMdmac3tcrFieldPamShift))
}

const (
	RegisterMdmac3tcrFieldTrgmShift = 28
	RegisterMdmac3tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac3tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldTrgmMask) >> RegisterMdmac3tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac3tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac3tcrFieldTrgmShift))
}

const (
	RegisterMdmac3tcrFieldSwrmShift = 30
	RegisterMdmac3tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac3tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac3tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac3tcrFieldBwmShift = 31
	RegisterMdmac3tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac3tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac3tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tcrFieldBwmMask)
	}
}

// RegisterMdmac3bndtrType MDMA Channel x block number of data register
type RegisterMdmac3bndtrType uint32

func (r *RegisterMdmac3bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3bndtrFieldBndtShift = 0
	RegisterMdmac3bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac3bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3bndtrFieldBndtMask) >> RegisterMdmac3bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac3bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac3bndtrFieldBndtShift))
}

const (
	RegisterMdmac3bndtrFieldBrsumShift = 18
	RegisterMdmac3bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac3bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac3bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac3bndtrFieldBrdumShift = 19
	RegisterMdmac3bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac3bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac3bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac3bndtrFieldBrcShift = 20
	RegisterMdmac3bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac3bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3bndtrFieldBrcMask) >> RegisterMdmac3bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac3bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac3bndtrFieldBrcShift))
}

// RegisterMdmac3sarType MDMA channel x source address register
type RegisterMdmac3sarType uint32

func (r *RegisterMdmac3sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3sarFieldSarShift = 0
	RegisterMdmac3sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac3sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3sarFieldSarMask) >> RegisterMdmac3sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac3sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3sarFieldSarMask)|(uint32(value)<<RegisterMdmac3sarFieldSarShift))
}

// RegisterMdmac3darType MDMA channel x destination address register
type RegisterMdmac3darType uint32

func (r *RegisterMdmac3darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3darFieldDarShift = 0
	RegisterMdmac3darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac3darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3darFieldDarMask) >> RegisterMdmac3darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac3darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3darFieldDarMask)|(uint32(value)<<RegisterMdmac3darFieldDarShift))
}

// RegisterMdmac3brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac3brurType uint32

func (r *RegisterMdmac3brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3brurFieldSuvShift = 0
	RegisterMdmac3brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac3brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3brurFieldSuvMask) >> RegisterMdmac3brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac3brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3brurFieldSuvMask)|(uint32(value)<<RegisterMdmac3brurFieldSuvShift))
}

const (
	RegisterMdmac3brurFieldDuvShift = 16
	RegisterMdmac3brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac3brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3brurFieldDuvMask) >> RegisterMdmac3brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac3brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3brurFieldDuvMask)|(uint32(value)<<RegisterMdmac3brurFieldDuvShift))
}

// RegisterMdmac3larType MDMA channel x Link Address register
type RegisterMdmac3larType uint32

func (r *RegisterMdmac3larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3larFieldLarShift = 0
	RegisterMdmac3larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac3larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3larFieldLarMask) >> RegisterMdmac3larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac3larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3larFieldLarMask)|(uint32(value)<<RegisterMdmac3larFieldLarShift))
}

// RegisterMdmac3tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac3tbrType uint32

func (r *RegisterMdmac3tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3tbrFieldTselShift = 0
	RegisterMdmac3tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac3tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tbrFieldTselMask) >> RegisterMdmac3tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac3tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tbrFieldTselMask)|(uint32(value)<<RegisterMdmac3tbrFieldTselShift))
}

const (
	RegisterMdmac3tbrFieldSbusShift = 16
	RegisterMdmac3tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac3tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac3tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac3tbrFieldDbusShift = 17
	RegisterMdmac3tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac3tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac3tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac3tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3tbrFieldDbusMask)
	}
}

// RegisterMdmac3marType MDMA channel x Mask address register
type RegisterMdmac3marType uint32

func (r *RegisterMdmac3marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3marFieldMarShift = 0
	RegisterMdmac3marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac3marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3marFieldMarMask) >> RegisterMdmac3marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac3marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3marFieldMarMask)|(uint32(value)<<RegisterMdmac3marFieldMarShift))
}

// RegisterMdmac3mdrType MDMA channel x Mask Data register
type RegisterMdmac3mdrType uint32

func (r *RegisterMdmac3mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac3mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac3mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac3mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac3mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac3mdrFieldMdrShift = 0
	RegisterMdmac3mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac3mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac3mdrFieldMdrMask) >> RegisterMdmac3mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac3mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac3mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac3mdrFieldMdrShift))
}

// RegisterMdmac4isrType MDMA channel x interrupt/status register
type RegisterMdmac4isrType uint32

func (r *RegisterMdmac4isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4isrFieldTeif4Shift = 0
	RegisterMdmac4isrFieldTeif4Mask  = 0x1
)

// GetTeif4 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac4isrType) GetTeif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4isrFieldTeif4Mask) != 0
}

// SetTeif4 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac4isrType) SetTeif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4isrFieldTeif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4isrFieldTeif4Mask)
	}
}

const (
	RegisterMdmac4isrFieldCtcif4Shift = 1
	RegisterMdmac4isrFieldCtcif4Mask  = 0x2
)

// GetCtcif4 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac4isrType) GetCtcif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4isrFieldCtcif4Mask) != 0
}

// SetCtcif4 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac4isrType) SetCtcif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4isrFieldCtcif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4isrFieldCtcif4Mask)
	}
}

const (
	RegisterMdmac4isrFieldBrtif4Shift = 2
	RegisterMdmac4isrFieldBrtif4Mask  = 0x4
)

// GetBrtif4 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac4isrType) GetBrtif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4isrFieldBrtif4Mask) != 0
}

// SetBrtif4 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac4isrType) SetBrtif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4isrFieldBrtif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4isrFieldBrtif4Mask)
	}
}

const (
	RegisterMdmac4isrFieldBtif4Shift = 3
	RegisterMdmac4isrFieldBtif4Mask  = 0x8
)

// GetBtif4 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac4isrType) GetBtif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4isrFieldBtif4Mask) != 0
}

// SetBtif4 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac4isrType) SetBtif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4isrFieldBtif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4isrFieldBtif4Mask)
	}
}

const (
	RegisterMdmac4isrFieldTcif4Shift = 4
	RegisterMdmac4isrFieldTcif4Mask  = 0x10
)

// GetTcif4 channel x buffer transfer complete
func (r *RegisterMdmac4isrType) GetTcif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4isrFieldTcif4Mask) != 0
}

// SetTcif4 channel x buffer transfer complete
func (r *RegisterMdmac4isrType) SetTcif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4isrFieldTcif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4isrFieldTcif4Mask)
	}
}

const (
	RegisterMdmac4isrFieldCrqa4Shift = 16
	RegisterMdmac4isrFieldCrqa4Mask  = 0x10000
)

// GetCrqa4 channel x request active flag
func (r *RegisterMdmac4isrType) GetCrqa4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4isrFieldCrqa4Mask) != 0
}

// SetCrqa4 channel x request active flag
func (r *RegisterMdmac4isrType) SetCrqa4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4isrFieldCrqa4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4isrFieldCrqa4Mask)
	}
}

// RegisterMdmac4ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac4ifcrType uint32

func (r *RegisterMdmac4ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4ifcrFieldCteif4Shift = 0
	RegisterMdmac4ifcrFieldCteif4Mask  = 0x1
)

// GetCteif4 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac4ifcrType) GetCteif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4ifcrFieldCteif4Mask) != 0
}

// SetCteif4 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac4ifcrType) SetCteif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4ifcrFieldCteif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4ifcrFieldCteif4Mask)
	}
}

const (
	RegisterMdmac4ifcrFieldCctcif4Shift = 1
	RegisterMdmac4ifcrFieldCctcif4Mask  = 0x2
)

// GetCctcif4 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac4ifcrType) GetCctcif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4ifcrFieldCctcif4Mask) != 0
}

// SetCctcif4 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac4ifcrType) SetCctcif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4ifcrFieldCctcif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4ifcrFieldCctcif4Mask)
	}
}

const (
	RegisterMdmac4ifcrFieldCbrtif4Shift = 2
	RegisterMdmac4ifcrFieldCbrtif4Mask  = 0x4
)

// GetCbrtif4 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac4ifcrType) GetCbrtif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4ifcrFieldCbrtif4Mask) != 0
}

// SetCbrtif4 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac4ifcrType) SetCbrtif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4ifcrFieldCbrtif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4ifcrFieldCbrtif4Mask)
	}
}

const (
	RegisterMdmac4ifcrFieldCbtif4Shift = 3
	RegisterMdmac4ifcrFieldCbtif4Mask  = 0x8
)

// GetCbtif4 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac4ifcrType) GetCbtif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4ifcrFieldCbtif4Mask) != 0
}

// SetCbtif4 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac4ifcrType) SetCbtif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4ifcrFieldCbtif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4ifcrFieldCbtif4Mask)
	}
}

const (
	RegisterMdmac4ifcrFieldCltcif4Shift = 4
	RegisterMdmac4ifcrFieldCltcif4Mask  = 0x10
)

// GetCltcif4 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac4ifcrType) GetCltcif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4ifcrFieldCltcif4Mask) != 0
}

// SetCltcif4 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac4ifcrType) SetCltcif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4ifcrFieldCltcif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4ifcrFieldCltcif4Mask)
	}
}

// RegisterMdmac4esrType MDMA Channel x error status register
type RegisterMdmac4esrType uint32

func (r *RegisterMdmac4esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4esrFieldTeaShift = 0
	RegisterMdmac4esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac4esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4esrFieldTeaMask) >> RegisterMdmac4esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac4esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4esrFieldTeaMask)|(uint32(value)<<RegisterMdmac4esrFieldTeaShift))
}

const (
	RegisterMdmac4esrFieldTedShift = 7
	RegisterMdmac4esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac4esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac4esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4esrFieldTedMask)
	}
}

const (
	RegisterMdmac4esrFieldTeldShift = 8
	RegisterMdmac4esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac4esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac4esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4esrFieldTeldMask)
	}
}

const (
	RegisterMdmac4esrFieldTemdShift = 9
	RegisterMdmac4esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac4esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac4esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4esrFieldTemdMask)
	}
}

const (
	RegisterMdmac4esrFieldAseShift = 10
	RegisterMdmac4esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac4esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac4esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4esrFieldAseMask)
	}
}

const (
	RegisterMdmac4esrFieldBseShift = 11
	RegisterMdmac4esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac4esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac4esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4esrFieldBseMask)
	}
}

// RegisterMdmac4crType This register is used to control the concerned channel.
type RegisterMdmac4crType uint32

func (r *RegisterMdmac4crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4crFieldEnShift = 0
	RegisterMdmac4crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac4crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac4crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4crFieldEnMask)
	}
}

const (
	RegisterMdmac4crFieldTeieShift = 1
	RegisterMdmac4crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac4crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac4crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4crFieldTeieMask)
	}
}

const (
	RegisterMdmac4crFieldCtcieShift = 2
	RegisterMdmac4crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac4crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac4crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4crFieldCtcieMask)
	}
}

const (
	RegisterMdmac4crFieldBrtieShift = 3
	RegisterMdmac4crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac4crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac4crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4crFieldBrtieMask)
	}
}

const (
	RegisterMdmac4crFieldBtieShift = 4
	RegisterMdmac4crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac4crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac4crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4crFieldBtieMask)
	}
}

const (
	RegisterMdmac4crFieldTcieShift = 5
	RegisterMdmac4crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac4crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac4crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4crFieldTcieMask)
	}
}

const (
	RegisterMdmac4crFieldPlShift = 6
	RegisterMdmac4crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac4crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4crFieldPlMask) >> RegisterMdmac4crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac4crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4crFieldPlMask)|(uint32(value)<<RegisterMdmac4crFieldPlShift))
}

const (
	RegisterMdmac4crFieldBexShift = 12
	RegisterMdmac4crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac4crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac4crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4crFieldBexMask)
	}
}

const (
	RegisterMdmac4crFieldHexShift = 13
	RegisterMdmac4crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac4crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac4crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4crFieldHexMask)
	}
}

const (
	RegisterMdmac4crFieldWexShift = 14
	RegisterMdmac4crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac4crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac4crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4crFieldWexMask)
	}
}

const (
	RegisterMdmac4crFieldSwrqShift = 16
	RegisterMdmac4crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac4crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4crFieldSwrqMask)
	}
}

// RegisterMdmac4tcrType This register is used to configure the concerned channel.
type RegisterMdmac4tcrType uint32

func (r *RegisterMdmac4tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4tcrFieldSincShift = 0
	RegisterMdmac4tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac4tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldSincMask) >> RegisterMdmac4tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac4tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldSincMask)|(uint32(value)<<RegisterMdmac4tcrFieldSincShift))
}

const (
	RegisterMdmac4tcrFieldDincShift = 2
	RegisterMdmac4tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac4tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldDincMask) >> RegisterMdmac4tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac4tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldDincMask)|(uint32(value)<<RegisterMdmac4tcrFieldDincShift))
}

const (
	RegisterMdmac4tcrFieldSsizeShift = 4
	RegisterMdmac4tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac4tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldSsizeMask) >> RegisterMdmac4tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac4tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac4tcrFieldSsizeShift))
}

const (
	RegisterMdmac4tcrFieldDsizeShift = 6
	RegisterMdmac4tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac4tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldDsizeMask) >> RegisterMdmac4tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac4tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac4tcrFieldDsizeShift))
}

const (
	RegisterMdmac4tcrFieldSincosShift = 8
	RegisterMdmac4tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac4tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldSincosMask) >> RegisterMdmac4tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac4tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac4tcrFieldSincosShift))
}

const (
	RegisterMdmac4tcrFieldDincosShift = 10
	RegisterMdmac4tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac4tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldDincosMask) >> RegisterMdmac4tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac4tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac4tcrFieldDincosShift))
}

const (
	RegisterMdmac4tcrFieldSburstShift = 12
	RegisterMdmac4tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac4tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldSburstMask) >> RegisterMdmac4tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac4tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac4tcrFieldSburstShift))
}

const (
	RegisterMdmac4tcrFieldDburstShift = 15
	RegisterMdmac4tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac4tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldDburstMask) >> RegisterMdmac4tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac4tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac4tcrFieldDburstShift))
}

const (
	RegisterMdmac4tcrFieldTlenShift = 18
	RegisterMdmac4tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac4tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldTlenMask) >> RegisterMdmac4tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac4tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac4tcrFieldTlenShift))
}

const (
	RegisterMdmac4tcrFieldPkeShift = 25
	RegisterMdmac4tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac4tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac4tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac4tcrFieldPamShift = 26
	RegisterMdmac4tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac4tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldPamMask) >> RegisterMdmac4tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac4tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldPamMask)|(uint32(value)<<RegisterMdmac4tcrFieldPamShift))
}

const (
	RegisterMdmac4tcrFieldTrgmShift = 28
	RegisterMdmac4tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac4tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldTrgmMask) >> RegisterMdmac4tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac4tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac4tcrFieldTrgmShift))
}

const (
	RegisterMdmac4tcrFieldSwrmShift = 30
	RegisterMdmac4tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac4tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac4tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac4tcrFieldBwmShift = 31
	RegisterMdmac4tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac4tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac4tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tcrFieldBwmMask)
	}
}

// RegisterMdmac4bndtrType MDMA Channel x block number of data register
type RegisterMdmac4bndtrType uint32

func (r *RegisterMdmac4bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4bndtrFieldBndtShift = 0
	RegisterMdmac4bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac4bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4bndtrFieldBndtMask) >> RegisterMdmac4bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac4bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac4bndtrFieldBndtShift))
}

const (
	RegisterMdmac4bndtrFieldBrsumShift = 18
	RegisterMdmac4bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac4bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac4bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac4bndtrFieldBrdumShift = 19
	RegisterMdmac4bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac4bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac4bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac4bndtrFieldBrcShift = 20
	RegisterMdmac4bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac4bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4bndtrFieldBrcMask) >> RegisterMdmac4bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac4bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac4bndtrFieldBrcShift))
}

// RegisterMdmac4sarType MDMA channel x source address register
type RegisterMdmac4sarType uint32

func (r *RegisterMdmac4sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4sarFieldSarShift = 0
	RegisterMdmac4sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac4sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4sarFieldSarMask) >> RegisterMdmac4sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac4sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4sarFieldSarMask)|(uint32(value)<<RegisterMdmac4sarFieldSarShift))
}

// RegisterMdmac4darType MDMA channel x destination address register
type RegisterMdmac4darType uint32

func (r *RegisterMdmac4darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4darFieldDarShift = 0
	RegisterMdmac4darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac4darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4darFieldDarMask) >> RegisterMdmac4darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac4darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4darFieldDarMask)|(uint32(value)<<RegisterMdmac4darFieldDarShift))
}

// RegisterMdmac4brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac4brurType uint32

func (r *RegisterMdmac4brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4brurFieldSuvShift = 0
	RegisterMdmac4brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac4brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4brurFieldSuvMask) >> RegisterMdmac4brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac4brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4brurFieldSuvMask)|(uint32(value)<<RegisterMdmac4brurFieldSuvShift))
}

const (
	RegisterMdmac4brurFieldDuvShift = 16
	RegisterMdmac4brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac4brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4brurFieldDuvMask) >> RegisterMdmac4brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac4brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4brurFieldDuvMask)|(uint32(value)<<RegisterMdmac4brurFieldDuvShift))
}

// RegisterMdmac4larType MDMA channel x Link Address register
type RegisterMdmac4larType uint32

func (r *RegisterMdmac4larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4larFieldLarShift = 0
	RegisterMdmac4larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac4larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4larFieldLarMask) >> RegisterMdmac4larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac4larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4larFieldLarMask)|(uint32(value)<<RegisterMdmac4larFieldLarShift))
}

// RegisterMdmac4tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac4tbrType uint32

func (r *RegisterMdmac4tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4tbrFieldTselShift = 0
	RegisterMdmac4tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac4tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tbrFieldTselMask) >> RegisterMdmac4tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac4tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tbrFieldTselMask)|(uint32(value)<<RegisterMdmac4tbrFieldTselShift))
}

const (
	RegisterMdmac4tbrFieldSbusShift = 16
	RegisterMdmac4tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac4tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac4tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac4tbrFieldDbusShift = 17
	RegisterMdmac4tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac4tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac4tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac4tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4tbrFieldDbusMask)
	}
}

// RegisterMdmac4marType MDMA channel x Mask address register
type RegisterMdmac4marType uint32

func (r *RegisterMdmac4marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4marFieldMarShift = 0
	RegisterMdmac4marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac4marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4marFieldMarMask) >> RegisterMdmac4marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac4marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4marFieldMarMask)|(uint32(value)<<RegisterMdmac4marFieldMarShift))
}

// RegisterMdmac4mdrType MDMA channel x Mask Data register
type RegisterMdmac4mdrType uint32

func (r *RegisterMdmac4mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac4mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac4mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac4mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac4mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac4mdrFieldMdrShift = 0
	RegisterMdmac4mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac4mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac4mdrFieldMdrMask) >> RegisterMdmac4mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac4mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac4mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac4mdrFieldMdrShift))
}

// RegisterMdmac5isrType MDMA channel x interrupt/status register
type RegisterMdmac5isrType uint32

func (r *RegisterMdmac5isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5isrFieldTeif5Shift = 0
	RegisterMdmac5isrFieldTeif5Mask  = 0x1
)

// GetTeif5 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac5isrType) GetTeif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5isrFieldTeif5Mask) != 0
}

// SetTeif5 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac5isrType) SetTeif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5isrFieldTeif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5isrFieldTeif5Mask)
	}
}

const (
	RegisterMdmac5isrFieldCtcif5Shift = 1
	RegisterMdmac5isrFieldCtcif5Mask  = 0x2
)

// GetCtcif5 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac5isrType) GetCtcif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5isrFieldCtcif5Mask) != 0
}

// SetCtcif5 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac5isrType) SetCtcif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5isrFieldCtcif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5isrFieldCtcif5Mask)
	}
}

const (
	RegisterMdmac5isrFieldBrtif5Shift = 2
	RegisterMdmac5isrFieldBrtif5Mask  = 0x4
)

// GetBrtif5 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac5isrType) GetBrtif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5isrFieldBrtif5Mask) != 0
}

// SetBrtif5 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac5isrType) SetBrtif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5isrFieldBrtif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5isrFieldBrtif5Mask)
	}
}

const (
	RegisterMdmac5isrFieldBtif5Shift = 3
	RegisterMdmac5isrFieldBtif5Mask  = 0x8
)

// GetBtif5 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac5isrType) GetBtif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5isrFieldBtif5Mask) != 0
}

// SetBtif5 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac5isrType) SetBtif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5isrFieldBtif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5isrFieldBtif5Mask)
	}
}

const (
	RegisterMdmac5isrFieldTcif5Shift = 4
	RegisterMdmac5isrFieldTcif5Mask  = 0x10
)

// GetTcif5 channel x buffer transfer complete
func (r *RegisterMdmac5isrType) GetTcif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5isrFieldTcif5Mask) != 0
}

// SetTcif5 channel x buffer transfer complete
func (r *RegisterMdmac5isrType) SetTcif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5isrFieldTcif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5isrFieldTcif5Mask)
	}
}

const (
	RegisterMdmac5isrFieldCrqa5Shift = 16
	RegisterMdmac5isrFieldCrqa5Mask  = 0x10000
)

// GetCrqa5 channel x request active flag
func (r *RegisterMdmac5isrType) GetCrqa5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5isrFieldCrqa5Mask) != 0
}

// SetCrqa5 channel x request active flag
func (r *RegisterMdmac5isrType) SetCrqa5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5isrFieldCrqa5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5isrFieldCrqa5Mask)
	}
}

// RegisterMdmac5ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac5ifcrType uint32

func (r *RegisterMdmac5ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5ifcrFieldCteif5Shift = 0
	RegisterMdmac5ifcrFieldCteif5Mask  = 0x1
)

// GetCteif5 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac5ifcrType) GetCteif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5ifcrFieldCteif5Mask) != 0
}

// SetCteif5 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac5ifcrType) SetCteif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5ifcrFieldCteif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5ifcrFieldCteif5Mask)
	}
}

const (
	RegisterMdmac5ifcrFieldCctcif5Shift = 1
	RegisterMdmac5ifcrFieldCctcif5Mask  = 0x2
)

// GetCctcif5 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac5ifcrType) GetCctcif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5ifcrFieldCctcif5Mask) != 0
}

// SetCctcif5 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac5ifcrType) SetCctcif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5ifcrFieldCctcif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5ifcrFieldCctcif5Mask)
	}
}

const (
	RegisterMdmac5ifcrFieldCbrtif5Shift = 2
	RegisterMdmac5ifcrFieldCbrtif5Mask  = 0x4
)

// GetCbrtif5 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac5ifcrType) GetCbrtif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5ifcrFieldCbrtif5Mask) != 0
}

// SetCbrtif5 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac5ifcrType) SetCbrtif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5ifcrFieldCbrtif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5ifcrFieldCbrtif5Mask)
	}
}

const (
	RegisterMdmac5ifcrFieldCbtif5Shift = 3
	RegisterMdmac5ifcrFieldCbtif5Mask  = 0x8
)

// GetCbtif5 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac5ifcrType) GetCbtif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5ifcrFieldCbtif5Mask) != 0
}

// SetCbtif5 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac5ifcrType) SetCbtif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5ifcrFieldCbtif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5ifcrFieldCbtif5Mask)
	}
}

const (
	RegisterMdmac5ifcrFieldCltcif5Shift = 4
	RegisterMdmac5ifcrFieldCltcif5Mask  = 0x10
)

// GetCltcif5 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac5ifcrType) GetCltcif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5ifcrFieldCltcif5Mask) != 0
}

// SetCltcif5 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac5ifcrType) SetCltcif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5ifcrFieldCltcif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5ifcrFieldCltcif5Mask)
	}
}

// RegisterMdmac5esrType MDMA Channel x error status register
type RegisterMdmac5esrType uint32

func (r *RegisterMdmac5esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5esrFieldTeaShift = 0
	RegisterMdmac5esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac5esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5esrFieldTeaMask) >> RegisterMdmac5esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac5esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5esrFieldTeaMask)|(uint32(value)<<RegisterMdmac5esrFieldTeaShift))
}

const (
	RegisterMdmac5esrFieldTedShift = 7
	RegisterMdmac5esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac5esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac5esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5esrFieldTedMask)
	}
}

const (
	RegisterMdmac5esrFieldTeldShift = 8
	RegisterMdmac5esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac5esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac5esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5esrFieldTeldMask)
	}
}

const (
	RegisterMdmac5esrFieldTemdShift = 9
	RegisterMdmac5esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac5esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac5esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5esrFieldTemdMask)
	}
}

const (
	RegisterMdmac5esrFieldAseShift = 10
	RegisterMdmac5esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac5esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac5esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5esrFieldAseMask)
	}
}

const (
	RegisterMdmac5esrFieldBseShift = 11
	RegisterMdmac5esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac5esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac5esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5esrFieldBseMask)
	}
}

// RegisterMdmac5crType This register is used to control the concerned channel.
type RegisterMdmac5crType uint32

func (r *RegisterMdmac5crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5crFieldEnShift = 0
	RegisterMdmac5crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac5crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac5crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5crFieldEnMask)
	}
}

const (
	RegisterMdmac5crFieldTeieShift = 1
	RegisterMdmac5crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac5crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac5crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5crFieldTeieMask)
	}
}

const (
	RegisterMdmac5crFieldCtcieShift = 2
	RegisterMdmac5crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac5crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac5crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5crFieldCtcieMask)
	}
}

const (
	RegisterMdmac5crFieldBrtieShift = 3
	RegisterMdmac5crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac5crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac5crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5crFieldBrtieMask)
	}
}

const (
	RegisterMdmac5crFieldBtieShift = 4
	RegisterMdmac5crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac5crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac5crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5crFieldBtieMask)
	}
}

const (
	RegisterMdmac5crFieldTcieShift = 5
	RegisterMdmac5crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac5crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac5crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5crFieldTcieMask)
	}
}

const (
	RegisterMdmac5crFieldPlShift = 6
	RegisterMdmac5crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac5crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5crFieldPlMask) >> RegisterMdmac5crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac5crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5crFieldPlMask)|(uint32(value)<<RegisterMdmac5crFieldPlShift))
}

const (
	RegisterMdmac5crFieldBexShift = 12
	RegisterMdmac5crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac5crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac5crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5crFieldBexMask)
	}
}

const (
	RegisterMdmac5crFieldHexShift = 13
	RegisterMdmac5crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac5crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac5crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5crFieldHexMask)
	}
}

const (
	RegisterMdmac5crFieldWexShift = 14
	RegisterMdmac5crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac5crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac5crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5crFieldWexMask)
	}
}

const (
	RegisterMdmac5crFieldSwrqShift = 16
	RegisterMdmac5crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac5crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5crFieldSwrqMask)
	}
}

// RegisterMdmac5tcrType This register is used to configure the concerned channel.
type RegisterMdmac5tcrType uint32

func (r *RegisterMdmac5tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5tcrFieldSincShift = 0
	RegisterMdmac5tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac5tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldSincMask) >> RegisterMdmac5tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac5tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldSincMask)|(uint32(value)<<RegisterMdmac5tcrFieldSincShift))
}

const (
	RegisterMdmac5tcrFieldDincShift = 2
	RegisterMdmac5tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac5tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldDincMask) >> RegisterMdmac5tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac5tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldDincMask)|(uint32(value)<<RegisterMdmac5tcrFieldDincShift))
}

const (
	RegisterMdmac5tcrFieldSsizeShift = 4
	RegisterMdmac5tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac5tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldSsizeMask) >> RegisterMdmac5tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac5tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac5tcrFieldSsizeShift))
}

const (
	RegisterMdmac5tcrFieldDsizeShift = 6
	RegisterMdmac5tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac5tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldDsizeMask) >> RegisterMdmac5tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac5tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac5tcrFieldDsizeShift))
}

const (
	RegisterMdmac5tcrFieldSincosShift = 8
	RegisterMdmac5tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac5tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldSincosMask) >> RegisterMdmac5tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac5tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac5tcrFieldSincosShift))
}

const (
	RegisterMdmac5tcrFieldDincosShift = 10
	RegisterMdmac5tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac5tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldDincosMask) >> RegisterMdmac5tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac5tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac5tcrFieldDincosShift))
}

const (
	RegisterMdmac5tcrFieldSburstShift = 12
	RegisterMdmac5tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac5tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldSburstMask) >> RegisterMdmac5tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac5tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac5tcrFieldSburstShift))
}

const (
	RegisterMdmac5tcrFieldDburstShift = 15
	RegisterMdmac5tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac5tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldDburstMask) >> RegisterMdmac5tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac5tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac5tcrFieldDburstShift))
}

const (
	RegisterMdmac5tcrFieldTlenShift = 18
	RegisterMdmac5tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac5tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldTlenMask) >> RegisterMdmac5tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac5tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac5tcrFieldTlenShift))
}

const (
	RegisterMdmac5tcrFieldPkeShift = 25
	RegisterMdmac5tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac5tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac5tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac5tcrFieldPamShift = 26
	RegisterMdmac5tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac5tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldPamMask) >> RegisterMdmac5tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac5tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldPamMask)|(uint32(value)<<RegisterMdmac5tcrFieldPamShift))
}

const (
	RegisterMdmac5tcrFieldTrgmShift = 28
	RegisterMdmac5tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac5tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldTrgmMask) >> RegisterMdmac5tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac5tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac5tcrFieldTrgmShift))
}

const (
	RegisterMdmac5tcrFieldSwrmShift = 30
	RegisterMdmac5tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac5tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac5tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac5tcrFieldBwmShift = 31
	RegisterMdmac5tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac5tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac5tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tcrFieldBwmMask)
	}
}

// RegisterMdmac5bndtrType MDMA Channel x block number of data register
type RegisterMdmac5bndtrType uint32

func (r *RegisterMdmac5bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5bndtrFieldBndtShift = 0
	RegisterMdmac5bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac5bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5bndtrFieldBndtMask) >> RegisterMdmac5bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac5bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac5bndtrFieldBndtShift))
}

const (
	RegisterMdmac5bndtrFieldBrsumShift = 18
	RegisterMdmac5bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac5bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac5bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac5bndtrFieldBrdumShift = 19
	RegisterMdmac5bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac5bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac5bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac5bndtrFieldBrcShift = 20
	RegisterMdmac5bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac5bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5bndtrFieldBrcMask) >> RegisterMdmac5bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac5bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac5bndtrFieldBrcShift))
}

// RegisterMdmac5sarType MDMA channel x source address register
type RegisterMdmac5sarType uint32

func (r *RegisterMdmac5sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5sarFieldSarShift = 0
	RegisterMdmac5sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac5sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5sarFieldSarMask) >> RegisterMdmac5sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac5sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5sarFieldSarMask)|(uint32(value)<<RegisterMdmac5sarFieldSarShift))
}

// RegisterMdmac5darType MDMA channel x destination address register
type RegisterMdmac5darType uint32

func (r *RegisterMdmac5darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5darFieldDarShift = 0
	RegisterMdmac5darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac5darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5darFieldDarMask) >> RegisterMdmac5darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac5darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5darFieldDarMask)|(uint32(value)<<RegisterMdmac5darFieldDarShift))
}

// RegisterMdmac5brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac5brurType uint32

func (r *RegisterMdmac5brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5brurFieldSuvShift = 0
	RegisterMdmac5brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac5brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5brurFieldSuvMask) >> RegisterMdmac5brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac5brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5brurFieldSuvMask)|(uint32(value)<<RegisterMdmac5brurFieldSuvShift))
}

const (
	RegisterMdmac5brurFieldDuvShift = 16
	RegisterMdmac5brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac5brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5brurFieldDuvMask) >> RegisterMdmac5brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac5brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5brurFieldDuvMask)|(uint32(value)<<RegisterMdmac5brurFieldDuvShift))
}

// RegisterMdmac5larType MDMA channel x Link Address register
type RegisterMdmac5larType uint32

func (r *RegisterMdmac5larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5larFieldLarShift = 0
	RegisterMdmac5larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac5larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5larFieldLarMask) >> RegisterMdmac5larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac5larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5larFieldLarMask)|(uint32(value)<<RegisterMdmac5larFieldLarShift))
}

// RegisterMdmac5tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac5tbrType uint32

func (r *RegisterMdmac5tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5tbrFieldTselShift = 0
	RegisterMdmac5tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac5tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tbrFieldTselMask) >> RegisterMdmac5tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac5tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tbrFieldTselMask)|(uint32(value)<<RegisterMdmac5tbrFieldTselShift))
}

const (
	RegisterMdmac5tbrFieldSbusShift = 16
	RegisterMdmac5tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac5tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac5tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac5tbrFieldDbusShift = 17
	RegisterMdmac5tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac5tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac5tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac5tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5tbrFieldDbusMask)
	}
}

// RegisterMdmac5marType MDMA channel x Mask address register
type RegisterMdmac5marType uint32

func (r *RegisterMdmac5marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5marFieldMarShift = 0
	RegisterMdmac5marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac5marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5marFieldMarMask) >> RegisterMdmac5marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac5marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5marFieldMarMask)|(uint32(value)<<RegisterMdmac5marFieldMarShift))
}

// RegisterMdmac5mdrType MDMA channel x Mask Data register
type RegisterMdmac5mdrType uint32

func (r *RegisterMdmac5mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac5mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac5mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac5mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac5mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac5mdrFieldMdrShift = 0
	RegisterMdmac5mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac5mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac5mdrFieldMdrMask) >> RegisterMdmac5mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac5mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac5mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac5mdrFieldMdrShift))
}

// RegisterMdmac6isrType MDMA channel x interrupt/status register
type RegisterMdmac6isrType uint32

func (r *RegisterMdmac6isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6isrFieldTeif6Shift = 0
	RegisterMdmac6isrFieldTeif6Mask  = 0x1
)

// GetTeif6 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac6isrType) GetTeif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6isrFieldTeif6Mask) != 0
}

// SetTeif6 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac6isrType) SetTeif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6isrFieldTeif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6isrFieldTeif6Mask)
	}
}

const (
	RegisterMdmac6isrFieldCtcif6Shift = 1
	RegisterMdmac6isrFieldCtcif6Mask  = 0x2
)

// GetCtcif6 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac6isrType) GetCtcif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6isrFieldCtcif6Mask) != 0
}

// SetCtcif6 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac6isrType) SetCtcif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6isrFieldCtcif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6isrFieldCtcif6Mask)
	}
}

const (
	RegisterMdmac6isrFieldBrtif6Shift = 2
	RegisterMdmac6isrFieldBrtif6Mask  = 0x4
)

// GetBrtif6 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac6isrType) GetBrtif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6isrFieldBrtif6Mask) != 0
}

// SetBrtif6 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac6isrType) SetBrtif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6isrFieldBrtif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6isrFieldBrtif6Mask)
	}
}

const (
	RegisterMdmac6isrFieldBtif6Shift = 3
	RegisterMdmac6isrFieldBtif6Mask  = 0x8
)

// GetBtif6 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac6isrType) GetBtif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6isrFieldBtif6Mask) != 0
}

// SetBtif6 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac6isrType) SetBtif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6isrFieldBtif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6isrFieldBtif6Mask)
	}
}

const (
	RegisterMdmac6isrFieldTcif6Shift = 4
	RegisterMdmac6isrFieldTcif6Mask  = 0x10
)

// GetTcif6 channel x buffer transfer complete
func (r *RegisterMdmac6isrType) GetTcif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6isrFieldTcif6Mask) != 0
}

// SetTcif6 channel x buffer transfer complete
func (r *RegisterMdmac6isrType) SetTcif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6isrFieldTcif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6isrFieldTcif6Mask)
	}
}

const (
	RegisterMdmac6isrFieldCrqa6Shift = 16
	RegisterMdmac6isrFieldCrqa6Mask  = 0x10000
)

// GetCrqa6 channel x request active flag
func (r *RegisterMdmac6isrType) GetCrqa6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6isrFieldCrqa6Mask) != 0
}

// SetCrqa6 channel x request active flag
func (r *RegisterMdmac6isrType) SetCrqa6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6isrFieldCrqa6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6isrFieldCrqa6Mask)
	}
}

// RegisterMdmac6ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac6ifcrType uint32

func (r *RegisterMdmac6ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6ifcrFieldCteif6Shift = 0
	RegisterMdmac6ifcrFieldCteif6Mask  = 0x1
)

// GetCteif6 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac6ifcrType) GetCteif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6ifcrFieldCteif6Mask) != 0
}

// SetCteif6 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac6ifcrType) SetCteif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6ifcrFieldCteif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6ifcrFieldCteif6Mask)
	}
}

const (
	RegisterMdmac6ifcrFieldCctcif6Shift = 1
	RegisterMdmac6ifcrFieldCctcif6Mask  = 0x2
)

// GetCctcif6 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac6ifcrType) GetCctcif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6ifcrFieldCctcif6Mask) != 0
}

// SetCctcif6 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac6ifcrType) SetCctcif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6ifcrFieldCctcif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6ifcrFieldCctcif6Mask)
	}
}

const (
	RegisterMdmac6ifcrFieldCbrtif6Shift = 2
	RegisterMdmac6ifcrFieldCbrtif6Mask  = 0x4
)

// GetCbrtif6 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac6ifcrType) GetCbrtif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6ifcrFieldCbrtif6Mask) != 0
}

// SetCbrtif6 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac6ifcrType) SetCbrtif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6ifcrFieldCbrtif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6ifcrFieldCbrtif6Mask)
	}
}

const (
	RegisterMdmac6ifcrFieldCbtif6Shift = 3
	RegisterMdmac6ifcrFieldCbtif6Mask  = 0x8
)

// GetCbtif6 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac6ifcrType) GetCbtif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6ifcrFieldCbtif6Mask) != 0
}

// SetCbtif6 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac6ifcrType) SetCbtif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6ifcrFieldCbtif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6ifcrFieldCbtif6Mask)
	}
}

const (
	RegisterMdmac6ifcrFieldCltcif6Shift = 4
	RegisterMdmac6ifcrFieldCltcif6Mask  = 0x10
)

// GetCltcif6 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac6ifcrType) GetCltcif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6ifcrFieldCltcif6Mask) != 0
}

// SetCltcif6 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac6ifcrType) SetCltcif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6ifcrFieldCltcif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6ifcrFieldCltcif6Mask)
	}
}

// RegisterMdmac6esrType MDMA Channel x error status register
type RegisterMdmac6esrType uint32

func (r *RegisterMdmac6esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6esrFieldTeaShift = 0
	RegisterMdmac6esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac6esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6esrFieldTeaMask) >> RegisterMdmac6esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac6esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6esrFieldTeaMask)|(uint32(value)<<RegisterMdmac6esrFieldTeaShift))
}

const (
	RegisterMdmac6esrFieldTedShift = 7
	RegisterMdmac6esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac6esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac6esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6esrFieldTedMask)
	}
}

const (
	RegisterMdmac6esrFieldTeldShift = 8
	RegisterMdmac6esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac6esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac6esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6esrFieldTeldMask)
	}
}

const (
	RegisterMdmac6esrFieldTemdShift = 9
	RegisterMdmac6esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac6esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac6esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6esrFieldTemdMask)
	}
}

const (
	RegisterMdmac6esrFieldAseShift = 10
	RegisterMdmac6esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac6esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac6esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6esrFieldAseMask)
	}
}

const (
	RegisterMdmac6esrFieldBseShift = 11
	RegisterMdmac6esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac6esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac6esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6esrFieldBseMask)
	}
}

// RegisterMdmac6crType This register is used to control the concerned channel.
type RegisterMdmac6crType uint32

func (r *RegisterMdmac6crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6crFieldEnShift = 0
	RegisterMdmac6crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac6crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac6crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6crFieldEnMask)
	}
}

const (
	RegisterMdmac6crFieldTeieShift = 1
	RegisterMdmac6crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac6crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac6crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6crFieldTeieMask)
	}
}

const (
	RegisterMdmac6crFieldCtcieShift = 2
	RegisterMdmac6crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac6crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac6crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6crFieldCtcieMask)
	}
}

const (
	RegisterMdmac6crFieldBrtieShift = 3
	RegisterMdmac6crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac6crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac6crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6crFieldBrtieMask)
	}
}

const (
	RegisterMdmac6crFieldBtieShift = 4
	RegisterMdmac6crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac6crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac6crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6crFieldBtieMask)
	}
}

const (
	RegisterMdmac6crFieldTcieShift = 5
	RegisterMdmac6crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac6crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac6crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6crFieldTcieMask)
	}
}

const (
	RegisterMdmac6crFieldPlShift = 6
	RegisterMdmac6crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac6crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6crFieldPlMask) >> RegisterMdmac6crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac6crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6crFieldPlMask)|(uint32(value)<<RegisterMdmac6crFieldPlShift))
}

const (
	RegisterMdmac6crFieldBexShift = 12
	RegisterMdmac6crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac6crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac6crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6crFieldBexMask)
	}
}

const (
	RegisterMdmac6crFieldHexShift = 13
	RegisterMdmac6crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac6crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac6crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6crFieldHexMask)
	}
}

const (
	RegisterMdmac6crFieldWexShift = 14
	RegisterMdmac6crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac6crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac6crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6crFieldWexMask)
	}
}

const (
	RegisterMdmac6crFieldSwrqShift = 16
	RegisterMdmac6crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac6crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6crFieldSwrqMask)
	}
}

// RegisterMdmac6tcrType This register is used to configure the concerned channel.
type RegisterMdmac6tcrType uint32

func (r *RegisterMdmac6tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6tcrFieldSincShift = 0
	RegisterMdmac6tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac6tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldSincMask) >> RegisterMdmac6tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac6tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldSincMask)|(uint32(value)<<RegisterMdmac6tcrFieldSincShift))
}

const (
	RegisterMdmac6tcrFieldDincShift = 2
	RegisterMdmac6tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac6tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldDincMask) >> RegisterMdmac6tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac6tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldDincMask)|(uint32(value)<<RegisterMdmac6tcrFieldDincShift))
}

const (
	RegisterMdmac6tcrFieldSsizeShift = 4
	RegisterMdmac6tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac6tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldSsizeMask) >> RegisterMdmac6tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac6tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac6tcrFieldSsizeShift))
}

const (
	RegisterMdmac6tcrFieldDsizeShift = 6
	RegisterMdmac6tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac6tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldDsizeMask) >> RegisterMdmac6tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac6tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac6tcrFieldDsizeShift))
}

const (
	RegisterMdmac6tcrFieldSincosShift = 8
	RegisterMdmac6tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac6tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldSincosMask) >> RegisterMdmac6tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac6tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac6tcrFieldSincosShift))
}

const (
	RegisterMdmac6tcrFieldDincosShift = 10
	RegisterMdmac6tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac6tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldDincosMask) >> RegisterMdmac6tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac6tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac6tcrFieldDincosShift))
}

const (
	RegisterMdmac6tcrFieldSburstShift = 12
	RegisterMdmac6tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac6tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldSburstMask) >> RegisterMdmac6tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac6tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac6tcrFieldSburstShift))
}

const (
	RegisterMdmac6tcrFieldDburstShift = 15
	RegisterMdmac6tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac6tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldDburstMask) >> RegisterMdmac6tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac6tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac6tcrFieldDburstShift))
}

const (
	RegisterMdmac6tcrFieldTlenShift = 18
	RegisterMdmac6tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac6tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldTlenMask) >> RegisterMdmac6tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac6tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac6tcrFieldTlenShift))
}

const (
	RegisterMdmac6tcrFieldPkeShift = 25
	RegisterMdmac6tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac6tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac6tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac6tcrFieldPamShift = 26
	RegisterMdmac6tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac6tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldPamMask) >> RegisterMdmac6tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac6tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldPamMask)|(uint32(value)<<RegisterMdmac6tcrFieldPamShift))
}

const (
	RegisterMdmac6tcrFieldTrgmShift = 28
	RegisterMdmac6tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac6tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldTrgmMask) >> RegisterMdmac6tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac6tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac6tcrFieldTrgmShift))
}

const (
	RegisterMdmac6tcrFieldSwrmShift = 30
	RegisterMdmac6tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac6tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac6tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac6tcrFieldBwmShift = 31
	RegisterMdmac6tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac6tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac6tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tcrFieldBwmMask)
	}
}

// RegisterMdmac6bndtrType MDMA Channel x block number of data register
type RegisterMdmac6bndtrType uint32

func (r *RegisterMdmac6bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6bndtrFieldBndtShift = 0
	RegisterMdmac6bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac6bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6bndtrFieldBndtMask) >> RegisterMdmac6bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac6bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac6bndtrFieldBndtShift))
}

const (
	RegisterMdmac6bndtrFieldBrsumShift = 18
	RegisterMdmac6bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac6bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac6bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac6bndtrFieldBrdumShift = 19
	RegisterMdmac6bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac6bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac6bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac6bndtrFieldBrcShift = 20
	RegisterMdmac6bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac6bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6bndtrFieldBrcMask) >> RegisterMdmac6bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac6bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac6bndtrFieldBrcShift))
}

// RegisterMdmac6sarType MDMA channel x source address register
type RegisterMdmac6sarType uint32

func (r *RegisterMdmac6sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6sarFieldSarShift = 0
	RegisterMdmac6sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac6sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6sarFieldSarMask) >> RegisterMdmac6sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac6sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6sarFieldSarMask)|(uint32(value)<<RegisterMdmac6sarFieldSarShift))
}

// RegisterMdmac6darType MDMA channel x destination address register
type RegisterMdmac6darType uint32

func (r *RegisterMdmac6darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6darFieldDarShift = 0
	RegisterMdmac6darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac6darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6darFieldDarMask) >> RegisterMdmac6darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac6darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6darFieldDarMask)|(uint32(value)<<RegisterMdmac6darFieldDarShift))
}

// RegisterMdmac6brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac6brurType uint32

func (r *RegisterMdmac6brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6brurFieldSuvShift = 0
	RegisterMdmac6brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac6brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6brurFieldSuvMask) >> RegisterMdmac6brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac6brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6brurFieldSuvMask)|(uint32(value)<<RegisterMdmac6brurFieldSuvShift))
}

const (
	RegisterMdmac6brurFieldDuvShift = 16
	RegisterMdmac6brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac6brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6brurFieldDuvMask) >> RegisterMdmac6brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac6brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6brurFieldDuvMask)|(uint32(value)<<RegisterMdmac6brurFieldDuvShift))
}

// RegisterMdmac6larType MDMA channel x Link Address register
type RegisterMdmac6larType uint32

func (r *RegisterMdmac6larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6larFieldLarShift = 0
	RegisterMdmac6larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac6larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6larFieldLarMask) >> RegisterMdmac6larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac6larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6larFieldLarMask)|(uint32(value)<<RegisterMdmac6larFieldLarShift))
}

// RegisterMdmac6tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac6tbrType uint32

func (r *RegisterMdmac6tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6tbrFieldTselShift = 0
	RegisterMdmac6tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac6tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tbrFieldTselMask) >> RegisterMdmac6tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac6tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tbrFieldTselMask)|(uint32(value)<<RegisterMdmac6tbrFieldTselShift))
}

const (
	RegisterMdmac6tbrFieldSbusShift = 16
	RegisterMdmac6tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac6tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac6tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac6tbrFieldDbusShift = 17
	RegisterMdmac6tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac6tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac6tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac6tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6tbrFieldDbusMask)
	}
}

// RegisterMdmac6marType MDMA channel x Mask address register
type RegisterMdmac6marType uint32

func (r *RegisterMdmac6marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6marFieldMarShift = 0
	RegisterMdmac6marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac6marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6marFieldMarMask) >> RegisterMdmac6marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac6marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6marFieldMarMask)|(uint32(value)<<RegisterMdmac6marFieldMarShift))
}

// RegisterMdmac6mdrType MDMA channel x Mask Data register
type RegisterMdmac6mdrType uint32

func (r *RegisterMdmac6mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac6mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac6mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac6mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac6mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac6mdrFieldMdrShift = 0
	RegisterMdmac6mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac6mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac6mdrFieldMdrMask) >> RegisterMdmac6mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac6mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac6mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac6mdrFieldMdrShift))
}

// RegisterMdmac7isrType MDMA channel x interrupt/status register
type RegisterMdmac7isrType uint32

func (r *RegisterMdmac7isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7isrFieldTeif7Shift = 0
	RegisterMdmac7isrFieldTeif7Mask  = 0x1
)

// GetTeif7 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac7isrType) GetTeif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7isrFieldTeif7Mask) != 0
}

// SetTeif7 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac7isrType) SetTeif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7isrFieldTeif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7isrFieldTeif7Mask)
	}
}

const (
	RegisterMdmac7isrFieldCtcif7Shift = 1
	RegisterMdmac7isrFieldCtcif7Mask  = 0x2
)

// GetCtcif7 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac7isrType) GetCtcif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7isrFieldCtcif7Mask) != 0
}

// SetCtcif7 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac7isrType) SetCtcif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7isrFieldCtcif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7isrFieldCtcif7Mask)
	}
}

const (
	RegisterMdmac7isrFieldBrtif7Shift = 2
	RegisterMdmac7isrFieldBrtif7Mask  = 0x4
)

// GetBrtif7 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac7isrType) GetBrtif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7isrFieldBrtif7Mask) != 0
}

// SetBrtif7 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac7isrType) SetBrtif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7isrFieldBrtif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7isrFieldBrtif7Mask)
	}
}

const (
	RegisterMdmac7isrFieldBtif7Shift = 3
	RegisterMdmac7isrFieldBtif7Mask  = 0x8
)

// GetBtif7 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac7isrType) GetBtif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7isrFieldBtif7Mask) != 0
}

// SetBtif7 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac7isrType) SetBtif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7isrFieldBtif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7isrFieldBtif7Mask)
	}
}

const (
	RegisterMdmac7isrFieldTcif7Shift = 4
	RegisterMdmac7isrFieldTcif7Mask  = 0x10
)

// GetTcif7 channel x buffer transfer complete
func (r *RegisterMdmac7isrType) GetTcif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7isrFieldTcif7Mask) != 0
}

// SetTcif7 channel x buffer transfer complete
func (r *RegisterMdmac7isrType) SetTcif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7isrFieldTcif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7isrFieldTcif7Mask)
	}
}

const (
	RegisterMdmac7isrFieldCrqa7Shift = 16
	RegisterMdmac7isrFieldCrqa7Mask  = 0x10000
)

// GetCrqa7 channel x request active flag
func (r *RegisterMdmac7isrType) GetCrqa7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7isrFieldCrqa7Mask) != 0
}

// SetCrqa7 channel x request active flag
func (r *RegisterMdmac7isrType) SetCrqa7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7isrFieldCrqa7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7isrFieldCrqa7Mask)
	}
}

// RegisterMdmac7ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac7ifcrType uint32

func (r *RegisterMdmac7ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7ifcrFieldCteif7Shift = 0
	RegisterMdmac7ifcrFieldCteif7Mask  = 0x1
)

// GetCteif7 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac7ifcrType) GetCteif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7ifcrFieldCteif7Mask) != 0
}

// SetCteif7 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac7ifcrType) SetCteif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7ifcrFieldCteif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7ifcrFieldCteif7Mask)
	}
}

const (
	RegisterMdmac7ifcrFieldCctcif7Shift = 1
	RegisterMdmac7ifcrFieldCctcif7Mask  = 0x2
)

// GetCctcif7 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac7ifcrType) GetCctcif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7ifcrFieldCctcif7Mask) != 0
}

// SetCctcif7 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac7ifcrType) SetCctcif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7ifcrFieldCctcif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7ifcrFieldCctcif7Mask)
	}
}

const (
	RegisterMdmac7ifcrFieldCbrtif7Shift = 2
	RegisterMdmac7ifcrFieldCbrtif7Mask  = 0x4
)

// GetCbrtif7 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac7ifcrType) GetCbrtif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7ifcrFieldCbrtif7Mask) != 0
}

// SetCbrtif7 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac7ifcrType) SetCbrtif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7ifcrFieldCbrtif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7ifcrFieldCbrtif7Mask)
	}
}

const (
	RegisterMdmac7ifcrFieldCbtif7Shift = 3
	RegisterMdmac7ifcrFieldCbtif7Mask  = 0x8
)

// GetCbtif7 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac7ifcrType) GetCbtif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7ifcrFieldCbtif7Mask) != 0
}

// SetCbtif7 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac7ifcrType) SetCbtif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7ifcrFieldCbtif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7ifcrFieldCbtif7Mask)
	}
}

const (
	RegisterMdmac7ifcrFieldCltcif7Shift = 4
	RegisterMdmac7ifcrFieldCltcif7Mask  = 0x10
)

// GetCltcif7 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac7ifcrType) GetCltcif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7ifcrFieldCltcif7Mask) != 0
}

// SetCltcif7 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac7ifcrType) SetCltcif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7ifcrFieldCltcif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7ifcrFieldCltcif7Mask)
	}
}

// RegisterMdmac7esrType MDMA Channel x error status register
type RegisterMdmac7esrType uint32

func (r *RegisterMdmac7esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7esrFieldTeaShift = 0
	RegisterMdmac7esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac7esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7esrFieldTeaMask) >> RegisterMdmac7esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac7esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7esrFieldTeaMask)|(uint32(value)<<RegisterMdmac7esrFieldTeaShift))
}

const (
	RegisterMdmac7esrFieldTedShift = 7
	RegisterMdmac7esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac7esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac7esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7esrFieldTedMask)
	}
}

const (
	RegisterMdmac7esrFieldTeldShift = 8
	RegisterMdmac7esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac7esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac7esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7esrFieldTeldMask)
	}
}

const (
	RegisterMdmac7esrFieldTemdShift = 9
	RegisterMdmac7esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac7esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac7esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7esrFieldTemdMask)
	}
}

const (
	RegisterMdmac7esrFieldAseShift = 10
	RegisterMdmac7esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac7esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac7esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7esrFieldAseMask)
	}
}

const (
	RegisterMdmac7esrFieldBseShift = 11
	RegisterMdmac7esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac7esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac7esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7esrFieldBseMask)
	}
}

// RegisterMdmac7crType This register is used to control the concerned channel.
type RegisterMdmac7crType uint32

func (r *RegisterMdmac7crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7crFieldEnShift = 0
	RegisterMdmac7crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac7crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac7crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7crFieldEnMask)
	}
}

const (
	RegisterMdmac7crFieldTeieShift = 1
	RegisterMdmac7crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac7crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac7crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7crFieldTeieMask)
	}
}

const (
	RegisterMdmac7crFieldCtcieShift = 2
	RegisterMdmac7crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac7crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac7crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7crFieldCtcieMask)
	}
}

const (
	RegisterMdmac7crFieldBrtieShift = 3
	RegisterMdmac7crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac7crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac7crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7crFieldBrtieMask)
	}
}

const (
	RegisterMdmac7crFieldBtieShift = 4
	RegisterMdmac7crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac7crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac7crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7crFieldBtieMask)
	}
}

const (
	RegisterMdmac7crFieldTcieShift = 5
	RegisterMdmac7crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac7crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac7crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7crFieldTcieMask)
	}
}

const (
	RegisterMdmac7crFieldPlShift = 6
	RegisterMdmac7crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac7crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7crFieldPlMask) >> RegisterMdmac7crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac7crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7crFieldPlMask)|(uint32(value)<<RegisterMdmac7crFieldPlShift))
}

const (
	RegisterMdmac7crFieldBexShift = 12
	RegisterMdmac7crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac7crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac7crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7crFieldBexMask)
	}
}

const (
	RegisterMdmac7crFieldHexShift = 13
	RegisterMdmac7crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac7crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac7crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7crFieldHexMask)
	}
}

const (
	RegisterMdmac7crFieldWexShift = 14
	RegisterMdmac7crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac7crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac7crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7crFieldWexMask)
	}
}

const (
	RegisterMdmac7crFieldSwrqShift = 16
	RegisterMdmac7crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac7crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7crFieldSwrqMask)
	}
}

// RegisterMdmac7tcrType This register is used to configure the concerned channel.
type RegisterMdmac7tcrType uint32

func (r *RegisterMdmac7tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7tcrFieldSincShift = 0
	RegisterMdmac7tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac7tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldSincMask) >> RegisterMdmac7tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac7tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldSincMask)|(uint32(value)<<RegisterMdmac7tcrFieldSincShift))
}

const (
	RegisterMdmac7tcrFieldDincShift = 2
	RegisterMdmac7tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac7tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldDincMask) >> RegisterMdmac7tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac7tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldDincMask)|(uint32(value)<<RegisterMdmac7tcrFieldDincShift))
}

const (
	RegisterMdmac7tcrFieldSsizeShift = 4
	RegisterMdmac7tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac7tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldSsizeMask) >> RegisterMdmac7tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac7tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac7tcrFieldSsizeShift))
}

const (
	RegisterMdmac7tcrFieldDsizeShift = 6
	RegisterMdmac7tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac7tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldDsizeMask) >> RegisterMdmac7tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac7tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac7tcrFieldDsizeShift))
}

const (
	RegisterMdmac7tcrFieldSincosShift = 8
	RegisterMdmac7tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac7tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldSincosMask) >> RegisterMdmac7tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac7tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac7tcrFieldSincosShift))
}

const (
	RegisterMdmac7tcrFieldDincosShift = 10
	RegisterMdmac7tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac7tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldDincosMask) >> RegisterMdmac7tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac7tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac7tcrFieldDincosShift))
}

const (
	RegisterMdmac7tcrFieldSburstShift = 12
	RegisterMdmac7tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac7tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldSburstMask) >> RegisterMdmac7tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac7tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac7tcrFieldSburstShift))
}

const (
	RegisterMdmac7tcrFieldDburstShift = 15
	RegisterMdmac7tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac7tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldDburstMask) >> RegisterMdmac7tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac7tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac7tcrFieldDburstShift))
}

const (
	RegisterMdmac7tcrFieldTlenShift = 18
	RegisterMdmac7tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac7tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldTlenMask) >> RegisterMdmac7tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac7tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac7tcrFieldTlenShift))
}

const (
	RegisterMdmac7tcrFieldPkeShift = 25
	RegisterMdmac7tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac7tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac7tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac7tcrFieldPamShift = 26
	RegisterMdmac7tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac7tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldPamMask) >> RegisterMdmac7tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac7tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldPamMask)|(uint32(value)<<RegisterMdmac7tcrFieldPamShift))
}

const (
	RegisterMdmac7tcrFieldTrgmShift = 28
	RegisterMdmac7tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac7tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldTrgmMask) >> RegisterMdmac7tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac7tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac7tcrFieldTrgmShift))
}

const (
	RegisterMdmac7tcrFieldSwrmShift = 30
	RegisterMdmac7tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac7tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac7tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac7tcrFieldBwmShift = 31
	RegisterMdmac7tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac7tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac7tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tcrFieldBwmMask)
	}
}

// RegisterMdmac7bndtrType MDMA Channel x block number of data register
type RegisterMdmac7bndtrType uint32

func (r *RegisterMdmac7bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7bndtrFieldBndtShift = 0
	RegisterMdmac7bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac7bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7bndtrFieldBndtMask) >> RegisterMdmac7bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac7bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac7bndtrFieldBndtShift))
}

const (
	RegisterMdmac7bndtrFieldBrsumShift = 18
	RegisterMdmac7bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac7bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac7bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac7bndtrFieldBrdumShift = 19
	RegisterMdmac7bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac7bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac7bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac7bndtrFieldBrcShift = 20
	RegisterMdmac7bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac7bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7bndtrFieldBrcMask) >> RegisterMdmac7bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac7bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac7bndtrFieldBrcShift))
}

// RegisterMdmac7sarType MDMA channel x source address register
type RegisterMdmac7sarType uint32

func (r *RegisterMdmac7sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7sarFieldSarShift = 0
	RegisterMdmac7sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac7sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7sarFieldSarMask) >> RegisterMdmac7sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac7sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7sarFieldSarMask)|(uint32(value)<<RegisterMdmac7sarFieldSarShift))
}

// RegisterMdmac7darType MDMA channel x destination address register
type RegisterMdmac7darType uint32

func (r *RegisterMdmac7darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7darFieldDarShift = 0
	RegisterMdmac7darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac7darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7darFieldDarMask) >> RegisterMdmac7darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac7darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7darFieldDarMask)|(uint32(value)<<RegisterMdmac7darFieldDarShift))
}

// RegisterMdmac7brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac7brurType uint32

func (r *RegisterMdmac7brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7brurFieldSuvShift = 0
	RegisterMdmac7brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac7brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7brurFieldSuvMask) >> RegisterMdmac7brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac7brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7brurFieldSuvMask)|(uint32(value)<<RegisterMdmac7brurFieldSuvShift))
}

const (
	RegisterMdmac7brurFieldDuvShift = 16
	RegisterMdmac7brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac7brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7brurFieldDuvMask) >> RegisterMdmac7brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac7brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7brurFieldDuvMask)|(uint32(value)<<RegisterMdmac7brurFieldDuvShift))
}

// RegisterMdmac7larType MDMA channel x Link Address register
type RegisterMdmac7larType uint32

func (r *RegisterMdmac7larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7larFieldLarShift = 0
	RegisterMdmac7larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac7larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7larFieldLarMask) >> RegisterMdmac7larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac7larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7larFieldLarMask)|(uint32(value)<<RegisterMdmac7larFieldLarShift))
}

// RegisterMdmac7tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac7tbrType uint32

func (r *RegisterMdmac7tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7tbrFieldTselShift = 0
	RegisterMdmac7tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac7tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tbrFieldTselMask) >> RegisterMdmac7tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac7tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tbrFieldTselMask)|(uint32(value)<<RegisterMdmac7tbrFieldTselShift))
}

const (
	RegisterMdmac7tbrFieldSbusShift = 16
	RegisterMdmac7tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac7tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac7tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac7tbrFieldDbusShift = 17
	RegisterMdmac7tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac7tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac7tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac7tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7tbrFieldDbusMask)
	}
}

// RegisterMdmac7marType MDMA channel x Mask address register
type RegisterMdmac7marType uint32

func (r *RegisterMdmac7marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7marFieldMarShift = 0
	RegisterMdmac7marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac7marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7marFieldMarMask) >> RegisterMdmac7marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac7marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7marFieldMarMask)|(uint32(value)<<RegisterMdmac7marFieldMarShift))
}

// RegisterMdmac7mdrType MDMA channel x Mask Data register
type RegisterMdmac7mdrType uint32

func (r *RegisterMdmac7mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac7mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac7mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac7mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac7mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac7mdrFieldMdrShift = 0
	RegisterMdmac7mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac7mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac7mdrFieldMdrMask) >> RegisterMdmac7mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac7mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac7mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac7mdrFieldMdrShift))
}

// RegisterMdmac8isrType MDMA channel x interrupt/status register
type RegisterMdmac8isrType uint32

func (r *RegisterMdmac8isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8isrFieldTeif8Shift = 0
	RegisterMdmac8isrFieldTeif8Mask  = 0x1
)

// GetTeif8 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac8isrType) GetTeif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8isrFieldTeif8Mask) != 0
}

// SetTeif8 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac8isrType) SetTeif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8isrFieldTeif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8isrFieldTeif8Mask)
	}
}

const (
	RegisterMdmac8isrFieldCtcif8Shift = 1
	RegisterMdmac8isrFieldCtcif8Mask  = 0x2
)

// GetCtcif8 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac8isrType) GetCtcif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8isrFieldCtcif8Mask) != 0
}

// SetCtcif8 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac8isrType) SetCtcif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8isrFieldCtcif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8isrFieldCtcif8Mask)
	}
}

const (
	RegisterMdmac8isrFieldBrtif8Shift = 2
	RegisterMdmac8isrFieldBrtif8Mask  = 0x4
)

// GetBrtif8 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac8isrType) GetBrtif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8isrFieldBrtif8Mask) != 0
}

// SetBrtif8 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac8isrType) SetBrtif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8isrFieldBrtif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8isrFieldBrtif8Mask)
	}
}

const (
	RegisterMdmac8isrFieldBtif8Shift = 3
	RegisterMdmac8isrFieldBtif8Mask  = 0x8
)

// GetBtif8 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac8isrType) GetBtif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8isrFieldBtif8Mask) != 0
}

// SetBtif8 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac8isrType) SetBtif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8isrFieldBtif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8isrFieldBtif8Mask)
	}
}

const (
	RegisterMdmac8isrFieldTcif8Shift = 4
	RegisterMdmac8isrFieldTcif8Mask  = 0x10
)

// GetTcif8 channel x buffer transfer complete
func (r *RegisterMdmac8isrType) GetTcif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8isrFieldTcif8Mask) != 0
}

// SetTcif8 channel x buffer transfer complete
func (r *RegisterMdmac8isrType) SetTcif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8isrFieldTcif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8isrFieldTcif8Mask)
	}
}

const (
	RegisterMdmac8isrFieldCrqa8Shift = 16
	RegisterMdmac8isrFieldCrqa8Mask  = 0x10000
)

// GetCrqa8 channel x request active flag
func (r *RegisterMdmac8isrType) GetCrqa8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8isrFieldCrqa8Mask) != 0
}

// SetCrqa8 channel x request active flag
func (r *RegisterMdmac8isrType) SetCrqa8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8isrFieldCrqa8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8isrFieldCrqa8Mask)
	}
}

// RegisterMdmac8ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac8ifcrType uint32

func (r *RegisterMdmac8ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8ifcrFieldCteif8Shift = 0
	RegisterMdmac8ifcrFieldCteif8Mask  = 0x1
)

// GetCteif8 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac8ifcrType) GetCteif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8ifcrFieldCteif8Mask) != 0
}

// SetCteif8 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac8ifcrType) SetCteif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8ifcrFieldCteif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8ifcrFieldCteif8Mask)
	}
}

const (
	RegisterMdmac8ifcrFieldCctcif8Shift = 1
	RegisterMdmac8ifcrFieldCctcif8Mask  = 0x2
)

// GetCctcif8 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac8ifcrType) GetCctcif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8ifcrFieldCctcif8Mask) != 0
}

// SetCctcif8 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac8ifcrType) SetCctcif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8ifcrFieldCctcif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8ifcrFieldCctcif8Mask)
	}
}

const (
	RegisterMdmac8ifcrFieldCbrtif8Shift = 2
	RegisterMdmac8ifcrFieldCbrtif8Mask  = 0x4
)

// GetCbrtif8 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac8ifcrType) GetCbrtif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8ifcrFieldCbrtif8Mask) != 0
}

// SetCbrtif8 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac8ifcrType) SetCbrtif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8ifcrFieldCbrtif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8ifcrFieldCbrtif8Mask)
	}
}

const (
	RegisterMdmac8ifcrFieldCbtif8Shift = 3
	RegisterMdmac8ifcrFieldCbtif8Mask  = 0x8
)

// GetCbtif8 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac8ifcrType) GetCbtif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8ifcrFieldCbtif8Mask) != 0
}

// SetCbtif8 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac8ifcrType) SetCbtif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8ifcrFieldCbtif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8ifcrFieldCbtif8Mask)
	}
}

const (
	RegisterMdmac8ifcrFieldCltcif8Shift = 4
	RegisterMdmac8ifcrFieldCltcif8Mask  = 0x10
)

// GetCltcif8 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac8ifcrType) GetCltcif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8ifcrFieldCltcif8Mask) != 0
}

// SetCltcif8 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac8ifcrType) SetCltcif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8ifcrFieldCltcif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8ifcrFieldCltcif8Mask)
	}
}

// RegisterMdmac8esrType MDMA Channel x error status register
type RegisterMdmac8esrType uint32

func (r *RegisterMdmac8esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8esrFieldTeaShift = 0
	RegisterMdmac8esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac8esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8esrFieldTeaMask) >> RegisterMdmac8esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac8esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8esrFieldTeaMask)|(uint32(value)<<RegisterMdmac8esrFieldTeaShift))
}

const (
	RegisterMdmac8esrFieldTedShift = 7
	RegisterMdmac8esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac8esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac8esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8esrFieldTedMask)
	}
}

const (
	RegisterMdmac8esrFieldTeldShift = 8
	RegisterMdmac8esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac8esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac8esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8esrFieldTeldMask)
	}
}

const (
	RegisterMdmac8esrFieldTemdShift = 9
	RegisterMdmac8esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac8esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac8esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8esrFieldTemdMask)
	}
}

const (
	RegisterMdmac8esrFieldAseShift = 10
	RegisterMdmac8esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac8esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac8esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8esrFieldAseMask)
	}
}

const (
	RegisterMdmac8esrFieldBseShift = 11
	RegisterMdmac8esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac8esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac8esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8esrFieldBseMask)
	}
}

// RegisterMdmac8crType This register is used to control the concerned channel.
type RegisterMdmac8crType uint32

func (r *RegisterMdmac8crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8crFieldEnShift = 0
	RegisterMdmac8crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac8crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac8crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8crFieldEnMask)
	}
}

const (
	RegisterMdmac8crFieldTeieShift = 1
	RegisterMdmac8crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac8crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac8crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8crFieldTeieMask)
	}
}

const (
	RegisterMdmac8crFieldCtcieShift = 2
	RegisterMdmac8crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac8crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac8crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8crFieldCtcieMask)
	}
}

const (
	RegisterMdmac8crFieldBrtieShift = 3
	RegisterMdmac8crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac8crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac8crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8crFieldBrtieMask)
	}
}

const (
	RegisterMdmac8crFieldBtieShift = 4
	RegisterMdmac8crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac8crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac8crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8crFieldBtieMask)
	}
}

const (
	RegisterMdmac8crFieldTcieShift = 5
	RegisterMdmac8crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac8crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac8crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8crFieldTcieMask)
	}
}

const (
	RegisterMdmac8crFieldPlShift = 6
	RegisterMdmac8crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac8crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8crFieldPlMask) >> RegisterMdmac8crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac8crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8crFieldPlMask)|(uint32(value)<<RegisterMdmac8crFieldPlShift))
}

const (
	RegisterMdmac8crFieldBexShift = 12
	RegisterMdmac8crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac8crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac8crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8crFieldBexMask)
	}
}

const (
	RegisterMdmac8crFieldHexShift = 13
	RegisterMdmac8crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac8crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac8crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8crFieldHexMask)
	}
}

const (
	RegisterMdmac8crFieldWexShift = 14
	RegisterMdmac8crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac8crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac8crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8crFieldWexMask)
	}
}

const (
	RegisterMdmac8crFieldSwrqShift = 16
	RegisterMdmac8crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac8crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8crFieldSwrqMask)
	}
}

// RegisterMdmac8tcrType This register is used to configure the concerned channel.
type RegisterMdmac8tcrType uint32

func (r *RegisterMdmac8tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8tcrFieldSincShift = 0
	RegisterMdmac8tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac8tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldSincMask) >> RegisterMdmac8tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac8tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldSincMask)|(uint32(value)<<RegisterMdmac8tcrFieldSincShift))
}

const (
	RegisterMdmac8tcrFieldDincShift = 2
	RegisterMdmac8tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac8tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldDincMask) >> RegisterMdmac8tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac8tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldDincMask)|(uint32(value)<<RegisterMdmac8tcrFieldDincShift))
}

const (
	RegisterMdmac8tcrFieldSsizeShift = 4
	RegisterMdmac8tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac8tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldSsizeMask) >> RegisterMdmac8tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac8tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac8tcrFieldSsizeShift))
}

const (
	RegisterMdmac8tcrFieldDsizeShift = 6
	RegisterMdmac8tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac8tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldDsizeMask) >> RegisterMdmac8tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac8tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac8tcrFieldDsizeShift))
}

const (
	RegisterMdmac8tcrFieldSincosShift = 8
	RegisterMdmac8tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac8tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldSincosMask) >> RegisterMdmac8tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac8tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac8tcrFieldSincosShift))
}

const (
	RegisterMdmac8tcrFieldDincosShift = 10
	RegisterMdmac8tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac8tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldDincosMask) >> RegisterMdmac8tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac8tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac8tcrFieldDincosShift))
}

const (
	RegisterMdmac8tcrFieldSburstShift = 12
	RegisterMdmac8tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac8tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldSburstMask) >> RegisterMdmac8tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac8tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac8tcrFieldSburstShift))
}

const (
	RegisterMdmac8tcrFieldDburstShift = 15
	RegisterMdmac8tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac8tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldDburstMask) >> RegisterMdmac8tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac8tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac8tcrFieldDburstShift))
}

const (
	RegisterMdmac8tcrFieldTlenShift = 18
	RegisterMdmac8tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac8tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldTlenMask) >> RegisterMdmac8tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac8tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac8tcrFieldTlenShift))
}

const (
	RegisterMdmac8tcrFieldPkeShift = 25
	RegisterMdmac8tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac8tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac8tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac8tcrFieldPamShift = 26
	RegisterMdmac8tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac8tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldPamMask) >> RegisterMdmac8tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac8tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldPamMask)|(uint32(value)<<RegisterMdmac8tcrFieldPamShift))
}

const (
	RegisterMdmac8tcrFieldTrgmShift = 28
	RegisterMdmac8tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac8tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldTrgmMask) >> RegisterMdmac8tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac8tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac8tcrFieldTrgmShift))
}

const (
	RegisterMdmac8tcrFieldSwrmShift = 30
	RegisterMdmac8tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac8tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac8tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac8tcrFieldBwmShift = 31
	RegisterMdmac8tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac8tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac8tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tcrFieldBwmMask)
	}
}

// RegisterMdmac8bndtrType MDMA Channel x block number of data register
type RegisterMdmac8bndtrType uint32

func (r *RegisterMdmac8bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8bndtrFieldBndtShift = 0
	RegisterMdmac8bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac8bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8bndtrFieldBndtMask) >> RegisterMdmac8bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac8bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac8bndtrFieldBndtShift))
}

const (
	RegisterMdmac8bndtrFieldBrsumShift = 18
	RegisterMdmac8bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac8bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac8bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac8bndtrFieldBrdumShift = 19
	RegisterMdmac8bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac8bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac8bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac8bndtrFieldBrcShift = 20
	RegisterMdmac8bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac8bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8bndtrFieldBrcMask) >> RegisterMdmac8bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac8bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac8bndtrFieldBrcShift))
}

// RegisterMdmac8sarType MDMA channel x source address register
type RegisterMdmac8sarType uint32

func (r *RegisterMdmac8sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8sarFieldSarShift = 0
	RegisterMdmac8sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac8sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8sarFieldSarMask) >> RegisterMdmac8sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac8sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8sarFieldSarMask)|(uint32(value)<<RegisterMdmac8sarFieldSarShift))
}

// RegisterMdmac8darType MDMA channel x destination address register
type RegisterMdmac8darType uint32

func (r *RegisterMdmac8darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8darFieldDarShift = 0
	RegisterMdmac8darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac8darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8darFieldDarMask) >> RegisterMdmac8darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac8darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8darFieldDarMask)|(uint32(value)<<RegisterMdmac8darFieldDarShift))
}

// RegisterMdmac8brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac8brurType uint32

func (r *RegisterMdmac8brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8brurFieldSuvShift = 0
	RegisterMdmac8brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac8brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8brurFieldSuvMask) >> RegisterMdmac8brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac8brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8brurFieldSuvMask)|(uint32(value)<<RegisterMdmac8brurFieldSuvShift))
}

const (
	RegisterMdmac8brurFieldDuvShift = 16
	RegisterMdmac8brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac8brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8brurFieldDuvMask) >> RegisterMdmac8brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac8brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8brurFieldDuvMask)|(uint32(value)<<RegisterMdmac8brurFieldDuvShift))
}

// RegisterMdmac8larType MDMA channel x Link Address register
type RegisterMdmac8larType uint32

func (r *RegisterMdmac8larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8larFieldLarShift = 0
	RegisterMdmac8larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac8larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8larFieldLarMask) >> RegisterMdmac8larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac8larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8larFieldLarMask)|(uint32(value)<<RegisterMdmac8larFieldLarShift))
}

// RegisterMdmac8tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac8tbrType uint32

func (r *RegisterMdmac8tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8tbrFieldTselShift = 0
	RegisterMdmac8tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac8tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tbrFieldTselMask) >> RegisterMdmac8tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac8tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tbrFieldTselMask)|(uint32(value)<<RegisterMdmac8tbrFieldTselShift))
}

const (
	RegisterMdmac8tbrFieldSbusShift = 16
	RegisterMdmac8tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac8tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac8tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac8tbrFieldDbusShift = 17
	RegisterMdmac8tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac8tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac8tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac8tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8tbrFieldDbusMask)
	}
}

// RegisterMdmac8marType MDMA channel x Mask address register
type RegisterMdmac8marType uint32

func (r *RegisterMdmac8marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8marFieldMarShift = 0
	RegisterMdmac8marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac8marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8marFieldMarMask) >> RegisterMdmac8marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac8marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8marFieldMarMask)|(uint32(value)<<RegisterMdmac8marFieldMarShift))
}

// RegisterMdmac8mdrType MDMA channel x Mask Data register
type RegisterMdmac8mdrType uint32

func (r *RegisterMdmac8mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac8mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac8mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac8mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac8mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac8mdrFieldMdrShift = 0
	RegisterMdmac8mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac8mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac8mdrFieldMdrMask) >> RegisterMdmac8mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac8mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac8mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac8mdrFieldMdrShift))
}

// RegisterMdmac9isrType MDMA channel x interrupt/status register
type RegisterMdmac9isrType uint32

func (r *RegisterMdmac9isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9isrFieldTeif9Shift = 0
	RegisterMdmac9isrFieldTeif9Mask  = 0x1
)

// GetTeif9 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac9isrType) GetTeif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9isrFieldTeif9Mask) != 0
}

// SetTeif9 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac9isrType) SetTeif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9isrFieldTeif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9isrFieldTeif9Mask)
	}
}

const (
	RegisterMdmac9isrFieldCtcif9Shift = 1
	RegisterMdmac9isrFieldCtcif9Mask  = 0x2
)

// GetCtcif9 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac9isrType) GetCtcif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9isrFieldCtcif9Mask) != 0
}

// SetCtcif9 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac9isrType) SetCtcif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9isrFieldCtcif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9isrFieldCtcif9Mask)
	}
}

const (
	RegisterMdmac9isrFieldBrtif9Shift = 2
	RegisterMdmac9isrFieldBrtif9Mask  = 0x4
)

// GetBrtif9 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac9isrType) GetBrtif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9isrFieldBrtif9Mask) != 0
}

// SetBrtif9 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac9isrType) SetBrtif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9isrFieldBrtif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9isrFieldBrtif9Mask)
	}
}

const (
	RegisterMdmac9isrFieldBtif9Shift = 3
	RegisterMdmac9isrFieldBtif9Mask  = 0x8
)

// GetBtif9 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac9isrType) GetBtif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9isrFieldBtif9Mask) != 0
}

// SetBtif9 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac9isrType) SetBtif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9isrFieldBtif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9isrFieldBtif9Mask)
	}
}

const (
	RegisterMdmac9isrFieldTcif9Shift = 4
	RegisterMdmac9isrFieldTcif9Mask  = 0x10
)

// GetTcif9 channel x buffer transfer complete
func (r *RegisterMdmac9isrType) GetTcif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9isrFieldTcif9Mask) != 0
}

// SetTcif9 channel x buffer transfer complete
func (r *RegisterMdmac9isrType) SetTcif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9isrFieldTcif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9isrFieldTcif9Mask)
	}
}

const (
	RegisterMdmac9isrFieldCrqa9Shift = 16
	RegisterMdmac9isrFieldCrqa9Mask  = 0x10000
)

// GetCrqa9 channel x request active flag
func (r *RegisterMdmac9isrType) GetCrqa9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9isrFieldCrqa9Mask) != 0
}

// SetCrqa9 channel x request active flag
func (r *RegisterMdmac9isrType) SetCrqa9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9isrFieldCrqa9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9isrFieldCrqa9Mask)
	}
}

// RegisterMdmac9ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac9ifcrType uint32

func (r *RegisterMdmac9ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9ifcrFieldCteif9Shift = 0
	RegisterMdmac9ifcrFieldCteif9Mask  = 0x1
)

// GetCteif9 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac9ifcrType) GetCteif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9ifcrFieldCteif9Mask) != 0
}

// SetCteif9 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac9ifcrType) SetCteif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9ifcrFieldCteif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9ifcrFieldCteif9Mask)
	}
}

const (
	RegisterMdmac9ifcrFieldCctcif9Shift = 1
	RegisterMdmac9ifcrFieldCctcif9Mask  = 0x2
)

// GetCctcif9 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac9ifcrType) GetCctcif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9ifcrFieldCctcif9Mask) != 0
}

// SetCctcif9 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac9ifcrType) SetCctcif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9ifcrFieldCctcif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9ifcrFieldCctcif9Mask)
	}
}

const (
	RegisterMdmac9ifcrFieldCbrtif9Shift = 2
	RegisterMdmac9ifcrFieldCbrtif9Mask  = 0x4
)

// GetCbrtif9 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac9ifcrType) GetCbrtif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9ifcrFieldCbrtif9Mask) != 0
}

// SetCbrtif9 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac9ifcrType) SetCbrtif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9ifcrFieldCbrtif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9ifcrFieldCbrtif9Mask)
	}
}

const (
	RegisterMdmac9ifcrFieldCbtif9Shift = 3
	RegisterMdmac9ifcrFieldCbtif9Mask  = 0x8
)

// GetCbtif9 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac9ifcrType) GetCbtif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9ifcrFieldCbtif9Mask) != 0
}

// SetCbtif9 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac9ifcrType) SetCbtif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9ifcrFieldCbtif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9ifcrFieldCbtif9Mask)
	}
}

const (
	RegisterMdmac9ifcrFieldCltcif9Shift = 4
	RegisterMdmac9ifcrFieldCltcif9Mask  = 0x10
)

// GetCltcif9 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac9ifcrType) GetCltcif9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9ifcrFieldCltcif9Mask) != 0
}

// SetCltcif9 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac9ifcrType) SetCltcif9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9ifcrFieldCltcif9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9ifcrFieldCltcif9Mask)
	}
}

// RegisterMdmac9esrType MDMA Channel x error status register
type RegisterMdmac9esrType uint32

func (r *RegisterMdmac9esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9esrFieldTeaShift = 0
	RegisterMdmac9esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac9esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9esrFieldTeaMask) >> RegisterMdmac9esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac9esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9esrFieldTeaMask)|(uint32(value)<<RegisterMdmac9esrFieldTeaShift))
}

const (
	RegisterMdmac9esrFieldTedShift = 7
	RegisterMdmac9esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac9esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac9esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9esrFieldTedMask)
	}
}

const (
	RegisterMdmac9esrFieldTeldShift = 8
	RegisterMdmac9esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac9esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac9esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9esrFieldTeldMask)
	}
}

const (
	RegisterMdmac9esrFieldTemdShift = 9
	RegisterMdmac9esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac9esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac9esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9esrFieldTemdMask)
	}
}

const (
	RegisterMdmac9esrFieldAseShift = 10
	RegisterMdmac9esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac9esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac9esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9esrFieldAseMask)
	}
}

const (
	RegisterMdmac9esrFieldBseShift = 11
	RegisterMdmac9esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac9esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac9esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9esrFieldBseMask)
	}
}

// RegisterMdmac9crType This register is used to control the concerned channel.
type RegisterMdmac9crType uint32

func (r *RegisterMdmac9crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9crFieldEnShift = 0
	RegisterMdmac9crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac9crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac9crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9crFieldEnMask)
	}
}

const (
	RegisterMdmac9crFieldTeieShift = 1
	RegisterMdmac9crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac9crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac9crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9crFieldTeieMask)
	}
}

const (
	RegisterMdmac9crFieldCtcieShift = 2
	RegisterMdmac9crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac9crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac9crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9crFieldCtcieMask)
	}
}

const (
	RegisterMdmac9crFieldBrtieShift = 3
	RegisterMdmac9crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac9crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac9crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9crFieldBrtieMask)
	}
}

const (
	RegisterMdmac9crFieldBtieShift = 4
	RegisterMdmac9crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac9crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac9crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9crFieldBtieMask)
	}
}

const (
	RegisterMdmac9crFieldTcieShift = 5
	RegisterMdmac9crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac9crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac9crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9crFieldTcieMask)
	}
}

const (
	RegisterMdmac9crFieldPlShift = 6
	RegisterMdmac9crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac9crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9crFieldPlMask) >> RegisterMdmac9crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac9crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9crFieldPlMask)|(uint32(value)<<RegisterMdmac9crFieldPlShift))
}

const (
	RegisterMdmac9crFieldBexShift = 12
	RegisterMdmac9crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac9crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac9crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9crFieldBexMask)
	}
}

const (
	RegisterMdmac9crFieldHexShift = 13
	RegisterMdmac9crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac9crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac9crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9crFieldHexMask)
	}
}

const (
	RegisterMdmac9crFieldWexShift = 14
	RegisterMdmac9crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac9crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac9crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9crFieldWexMask)
	}
}

const (
	RegisterMdmac9crFieldSwrqShift = 16
	RegisterMdmac9crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac9crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9crFieldSwrqMask)
	}
}

// RegisterMdmac9tcrType This register is used to configure the concerned channel.
type RegisterMdmac9tcrType uint32

func (r *RegisterMdmac9tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9tcrFieldSincShift = 0
	RegisterMdmac9tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac9tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldSincMask) >> RegisterMdmac9tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac9tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldSincMask)|(uint32(value)<<RegisterMdmac9tcrFieldSincShift))
}

const (
	RegisterMdmac9tcrFieldDincShift = 2
	RegisterMdmac9tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac9tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldDincMask) >> RegisterMdmac9tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac9tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldDincMask)|(uint32(value)<<RegisterMdmac9tcrFieldDincShift))
}

const (
	RegisterMdmac9tcrFieldSsizeShift = 4
	RegisterMdmac9tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac9tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldSsizeMask) >> RegisterMdmac9tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac9tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac9tcrFieldSsizeShift))
}

const (
	RegisterMdmac9tcrFieldDsizeShift = 6
	RegisterMdmac9tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac9tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldDsizeMask) >> RegisterMdmac9tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac9tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac9tcrFieldDsizeShift))
}

const (
	RegisterMdmac9tcrFieldSincosShift = 8
	RegisterMdmac9tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac9tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldSincosMask) >> RegisterMdmac9tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac9tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac9tcrFieldSincosShift))
}

const (
	RegisterMdmac9tcrFieldDincosShift = 10
	RegisterMdmac9tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac9tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldDincosMask) >> RegisterMdmac9tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac9tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac9tcrFieldDincosShift))
}

const (
	RegisterMdmac9tcrFieldSburstShift = 12
	RegisterMdmac9tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac9tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldSburstMask) >> RegisterMdmac9tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac9tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac9tcrFieldSburstShift))
}

const (
	RegisterMdmac9tcrFieldDburstShift = 15
	RegisterMdmac9tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac9tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldDburstMask) >> RegisterMdmac9tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac9tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac9tcrFieldDburstShift))
}

const (
	RegisterMdmac9tcrFieldTlenShift = 18
	RegisterMdmac9tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac9tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldTlenMask) >> RegisterMdmac9tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac9tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac9tcrFieldTlenShift))
}

const (
	RegisterMdmac9tcrFieldPkeShift = 25
	RegisterMdmac9tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac9tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac9tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac9tcrFieldPamShift = 26
	RegisterMdmac9tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac9tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldPamMask) >> RegisterMdmac9tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac9tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldPamMask)|(uint32(value)<<RegisterMdmac9tcrFieldPamShift))
}

const (
	RegisterMdmac9tcrFieldTrgmShift = 28
	RegisterMdmac9tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac9tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldTrgmMask) >> RegisterMdmac9tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac9tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac9tcrFieldTrgmShift))
}

const (
	RegisterMdmac9tcrFieldSwrmShift = 30
	RegisterMdmac9tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac9tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac9tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac9tcrFieldBwmShift = 31
	RegisterMdmac9tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac9tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac9tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tcrFieldBwmMask)
	}
}

// RegisterMdmac9bndtrType MDMA Channel x block number of data register
type RegisterMdmac9bndtrType uint32

func (r *RegisterMdmac9bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9bndtrFieldBndtShift = 0
	RegisterMdmac9bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac9bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9bndtrFieldBndtMask) >> RegisterMdmac9bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac9bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac9bndtrFieldBndtShift))
}

const (
	RegisterMdmac9bndtrFieldBrsumShift = 18
	RegisterMdmac9bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac9bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac9bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac9bndtrFieldBrdumShift = 19
	RegisterMdmac9bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac9bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac9bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac9bndtrFieldBrcShift = 20
	RegisterMdmac9bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac9bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9bndtrFieldBrcMask) >> RegisterMdmac9bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac9bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac9bndtrFieldBrcShift))
}

// RegisterMdmac9sarType MDMA channel x source address register
type RegisterMdmac9sarType uint32

func (r *RegisterMdmac9sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9sarFieldSarShift = 0
	RegisterMdmac9sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac9sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9sarFieldSarMask) >> RegisterMdmac9sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac9sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9sarFieldSarMask)|(uint32(value)<<RegisterMdmac9sarFieldSarShift))
}

// RegisterMdmac9darType MDMA channel x destination address register
type RegisterMdmac9darType uint32

func (r *RegisterMdmac9darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9darFieldDarShift = 0
	RegisterMdmac9darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac9darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9darFieldDarMask) >> RegisterMdmac9darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac9darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9darFieldDarMask)|(uint32(value)<<RegisterMdmac9darFieldDarShift))
}

// RegisterMdmac9brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac9brurType uint32

func (r *RegisterMdmac9brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9brurFieldSuvShift = 0
	RegisterMdmac9brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac9brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9brurFieldSuvMask) >> RegisterMdmac9brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac9brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9brurFieldSuvMask)|(uint32(value)<<RegisterMdmac9brurFieldSuvShift))
}

const (
	RegisterMdmac9brurFieldDuvShift = 16
	RegisterMdmac9brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac9brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9brurFieldDuvMask) >> RegisterMdmac9brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac9brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9brurFieldDuvMask)|(uint32(value)<<RegisterMdmac9brurFieldDuvShift))
}

// RegisterMdmac9larType MDMA channel x Link Address register
type RegisterMdmac9larType uint32

func (r *RegisterMdmac9larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9larFieldLarShift = 0
	RegisterMdmac9larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac9larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9larFieldLarMask) >> RegisterMdmac9larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac9larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9larFieldLarMask)|(uint32(value)<<RegisterMdmac9larFieldLarShift))
}

// RegisterMdmac9tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac9tbrType uint32

func (r *RegisterMdmac9tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9tbrFieldTselShift = 0
	RegisterMdmac9tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac9tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tbrFieldTselMask) >> RegisterMdmac9tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac9tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tbrFieldTselMask)|(uint32(value)<<RegisterMdmac9tbrFieldTselShift))
}

const (
	RegisterMdmac9tbrFieldSbusShift = 16
	RegisterMdmac9tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac9tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac9tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac9tbrFieldDbusShift = 17
	RegisterMdmac9tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac9tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac9tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac9tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9tbrFieldDbusMask)
	}
}

// RegisterMdmac9marType MDMA channel x Mask address register
type RegisterMdmac9marType uint32

func (r *RegisterMdmac9marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9marFieldMarShift = 0
	RegisterMdmac9marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac9marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9marFieldMarMask) >> RegisterMdmac9marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac9marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9marFieldMarMask)|(uint32(value)<<RegisterMdmac9marFieldMarShift))
}

// RegisterMdmac9mdrType MDMA channel x Mask Data register
type RegisterMdmac9mdrType uint32

func (r *RegisterMdmac9mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac9mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac9mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac9mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac9mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac9mdrFieldMdrShift = 0
	RegisterMdmac9mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac9mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac9mdrFieldMdrMask) >> RegisterMdmac9mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac9mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac9mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac9mdrFieldMdrShift))
}

// RegisterMdmac10isrType MDMA channel x interrupt/status register
type RegisterMdmac10isrType uint32

func (r *RegisterMdmac10isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10isrFieldTeif10Shift = 0
	RegisterMdmac10isrFieldTeif10Mask  = 0x1
)

// GetTeif10 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac10isrType) GetTeif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10isrFieldTeif10Mask) != 0
}

// SetTeif10 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac10isrType) SetTeif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10isrFieldTeif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10isrFieldTeif10Mask)
	}
}

const (
	RegisterMdmac10isrFieldCtcif10Shift = 1
	RegisterMdmac10isrFieldCtcif10Mask  = 0x2
)

// GetCtcif10 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac10isrType) GetCtcif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10isrFieldCtcif10Mask) != 0
}

// SetCtcif10 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac10isrType) SetCtcif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10isrFieldCtcif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10isrFieldCtcif10Mask)
	}
}

const (
	RegisterMdmac10isrFieldBrtif10Shift = 2
	RegisterMdmac10isrFieldBrtif10Mask  = 0x4
)

// GetBrtif10 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac10isrType) GetBrtif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10isrFieldBrtif10Mask) != 0
}

// SetBrtif10 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac10isrType) SetBrtif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10isrFieldBrtif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10isrFieldBrtif10Mask)
	}
}

const (
	RegisterMdmac10isrFieldBtif10Shift = 3
	RegisterMdmac10isrFieldBtif10Mask  = 0x8
)

// GetBtif10 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac10isrType) GetBtif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10isrFieldBtif10Mask) != 0
}

// SetBtif10 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac10isrType) SetBtif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10isrFieldBtif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10isrFieldBtif10Mask)
	}
}

const (
	RegisterMdmac10isrFieldTcif10Shift = 4
	RegisterMdmac10isrFieldTcif10Mask  = 0x10
)

// GetTcif10 channel x buffer transfer complete
func (r *RegisterMdmac10isrType) GetTcif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10isrFieldTcif10Mask) != 0
}

// SetTcif10 channel x buffer transfer complete
func (r *RegisterMdmac10isrType) SetTcif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10isrFieldTcif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10isrFieldTcif10Mask)
	}
}

const (
	RegisterMdmac10isrFieldCrqa10Shift = 16
	RegisterMdmac10isrFieldCrqa10Mask  = 0x10000
)

// GetCrqa10 channel x request active flag
func (r *RegisterMdmac10isrType) GetCrqa10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10isrFieldCrqa10Mask) != 0
}

// SetCrqa10 channel x request active flag
func (r *RegisterMdmac10isrType) SetCrqa10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10isrFieldCrqa10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10isrFieldCrqa10Mask)
	}
}

// RegisterMdmac10ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac10ifcrType uint32

func (r *RegisterMdmac10ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10ifcrFieldCteif10Shift = 0
	RegisterMdmac10ifcrFieldCteif10Mask  = 0x1
)

// GetCteif10 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac10ifcrType) GetCteif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10ifcrFieldCteif10Mask) != 0
}

// SetCteif10 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac10ifcrType) SetCteif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10ifcrFieldCteif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10ifcrFieldCteif10Mask)
	}
}

const (
	RegisterMdmac10ifcrFieldCctcif10Shift = 1
	RegisterMdmac10ifcrFieldCctcif10Mask  = 0x2
)

// GetCctcif10 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac10ifcrType) GetCctcif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10ifcrFieldCctcif10Mask) != 0
}

// SetCctcif10 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac10ifcrType) SetCctcif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10ifcrFieldCctcif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10ifcrFieldCctcif10Mask)
	}
}

const (
	RegisterMdmac10ifcrFieldCbrtif10Shift = 2
	RegisterMdmac10ifcrFieldCbrtif10Mask  = 0x4
)

// GetCbrtif10 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac10ifcrType) GetCbrtif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10ifcrFieldCbrtif10Mask) != 0
}

// SetCbrtif10 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac10ifcrType) SetCbrtif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10ifcrFieldCbrtif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10ifcrFieldCbrtif10Mask)
	}
}

const (
	RegisterMdmac10ifcrFieldCbtif10Shift = 3
	RegisterMdmac10ifcrFieldCbtif10Mask  = 0x8
)

// GetCbtif10 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac10ifcrType) GetCbtif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10ifcrFieldCbtif10Mask) != 0
}

// SetCbtif10 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac10ifcrType) SetCbtif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10ifcrFieldCbtif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10ifcrFieldCbtif10Mask)
	}
}

const (
	RegisterMdmac10ifcrFieldCltcif10Shift = 4
	RegisterMdmac10ifcrFieldCltcif10Mask  = 0x10
)

// GetCltcif10 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac10ifcrType) GetCltcif10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10ifcrFieldCltcif10Mask) != 0
}

// SetCltcif10 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac10ifcrType) SetCltcif10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10ifcrFieldCltcif10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10ifcrFieldCltcif10Mask)
	}
}

// RegisterMdmac10esrType MDMA Channel x error status register
type RegisterMdmac10esrType uint32

func (r *RegisterMdmac10esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10esrFieldTeaShift = 0
	RegisterMdmac10esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac10esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10esrFieldTeaMask) >> RegisterMdmac10esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac10esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10esrFieldTeaMask)|(uint32(value)<<RegisterMdmac10esrFieldTeaShift))
}

const (
	RegisterMdmac10esrFieldTedShift = 7
	RegisterMdmac10esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac10esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac10esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10esrFieldTedMask)
	}
}

const (
	RegisterMdmac10esrFieldTeldShift = 8
	RegisterMdmac10esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac10esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac10esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10esrFieldTeldMask)
	}
}

const (
	RegisterMdmac10esrFieldTemdShift = 9
	RegisterMdmac10esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac10esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac10esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10esrFieldTemdMask)
	}
}

const (
	RegisterMdmac10esrFieldAseShift = 10
	RegisterMdmac10esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac10esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac10esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10esrFieldAseMask)
	}
}

const (
	RegisterMdmac10esrFieldBseShift = 11
	RegisterMdmac10esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac10esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac10esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10esrFieldBseMask)
	}
}

// RegisterMdmac10crType This register is used to control the concerned channel.
type RegisterMdmac10crType uint32

func (r *RegisterMdmac10crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10crFieldEnShift = 0
	RegisterMdmac10crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac10crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac10crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10crFieldEnMask)
	}
}

const (
	RegisterMdmac10crFieldTeieShift = 1
	RegisterMdmac10crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac10crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac10crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10crFieldTeieMask)
	}
}

const (
	RegisterMdmac10crFieldCtcieShift = 2
	RegisterMdmac10crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac10crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac10crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10crFieldCtcieMask)
	}
}

const (
	RegisterMdmac10crFieldBrtieShift = 3
	RegisterMdmac10crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac10crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac10crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10crFieldBrtieMask)
	}
}

const (
	RegisterMdmac10crFieldBtieShift = 4
	RegisterMdmac10crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac10crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac10crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10crFieldBtieMask)
	}
}

const (
	RegisterMdmac10crFieldTcieShift = 5
	RegisterMdmac10crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac10crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac10crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10crFieldTcieMask)
	}
}

const (
	RegisterMdmac10crFieldPlShift = 6
	RegisterMdmac10crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac10crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10crFieldPlMask) >> RegisterMdmac10crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac10crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10crFieldPlMask)|(uint32(value)<<RegisterMdmac10crFieldPlShift))
}

const (
	RegisterMdmac10crFieldBexShift = 12
	RegisterMdmac10crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac10crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac10crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10crFieldBexMask)
	}
}

const (
	RegisterMdmac10crFieldHexShift = 13
	RegisterMdmac10crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac10crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac10crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10crFieldHexMask)
	}
}

const (
	RegisterMdmac10crFieldWexShift = 14
	RegisterMdmac10crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac10crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac10crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10crFieldWexMask)
	}
}

const (
	RegisterMdmac10crFieldSwrqShift = 16
	RegisterMdmac10crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac10crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10crFieldSwrqMask)
	}
}

// RegisterMdmac10tcrType This register is used to configure the concerned channel.
type RegisterMdmac10tcrType uint32

func (r *RegisterMdmac10tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10tcrFieldSincShift = 0
	RegisterMdmac10tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac10tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldSincMask) >> RegisterMdmac10tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac10tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldSincMask)|(uint32(value)<<RegisterMdmac10tcrFieldSincShift))
}

const (
	RegisterMdmac10tcrFieldDincShift = 2
	RegisterMdmac10tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac10tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldDincMask) >> RegisterMdmac10tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac10tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldDincMask)|(uint32(value)<<RegisterMdmac10tcrFieldDincShift))
}

const (
	RegisterMdmac10tcrFieldSsizeShift = 4
	RegisterMdmac10tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac10tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldSsizeMask) >> RegisterMdmac10tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac10tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac10tcrFieldSsizeShift))
}

const (
	RegisterMdmac10tcrFieldDsizeShift = 6
	RegisterMdmac10tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac10tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldDsizeMask) >> RegisterMdmac10tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac10tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac10tcrFieldDsizeShift))
}

const (
	RegisterMdmac10tcrFieldSincosShift = 8
	RegisterMdmac10tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac10tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldSincosMask) >> RegisterMdmac10tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac10tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac10tcrFieldSincosShift))
}

const (
	RegisterMdmac10tcrFieldDincosShift = 10
	RegisterMdmac10tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac10tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldDincosMask) >> RegisterMdmac10tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac10tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac10tcrFieldDincosShift))
}

const (
	RegisterMdmac10tcrFieldSburstShift = 12
	RegisterMdmac10tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac10tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldSburstMask) >> RegisterMdmac10tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac10tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac10tcrFieldSburstShift))
}

const (
	RegisterMdmac10tcrFieldDburstShift = 15
	RegisterMdmac10tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac10tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldDburstMask) >> RegisterMdmac10tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac10tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac10tcrFieldDburstShift))
}

const (
	RegisterMdmac10tcrFieldTlenShift = 18
	RegisterMdmac10tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac10tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldTlenMask) >> RegisterMdmac10tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac10tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac10tcrFieldTlenShift))
}

const (
	RegisterMdmac10tcrFieldPkeShift = 25
	RegisterMdmac10tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac10tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac10tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac10tcrFieldPamShift = 26
	RegisterMdmac10tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac10tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldPamMask) >> RegisterMdmac10tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac10tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldPamMask)|(uint32(value)<<RegisterMdmac10tcrFieldPamShift))
}

const (
	RegisterMdmac10tcrFieldTrgmShift = 28
	RegisterMdmac10tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac10tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldTrgmMask) >> RegisterMdmac10tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac10tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac10tcrFieldTrgmShift))
}

const (
	RegisterMdmac10tcrFieldSwrmShift = 30
	RegisterMdmac10tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac10tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac10tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac10tcrFieldBwmShift = 31
	RegisterMdmac10tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac10tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac10tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tcrFieldBwmMask)
	}
}

// RegisterMdmac10bndtrType MDMA Channel x block number of data register
type RegisterMdmac10bndtrType uint32

func (r *RegisterMdmac10bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10bndtrFieldBndtShift = 0
	RegisterMdmac10bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac10bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10bndtrFieldBndtMask) >> RegisterMdmac10bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac10bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac10bndtrFieldBndtShift))
}

const (
	RegisterMdmac10bndtrFieldBrsumShift = 18
	RegisterMdmac10bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac10bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac10bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac10bndtrFieldBrdumShift = 19
	RegisterMdmac10bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac10bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac10bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac10bndtrFieldBrcShift = 20
	RegisterMdmac10bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac10bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10bndtrFieldBrcMask) >> RegisterMdmac10bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac10bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac10bndtrFieldBrcShift))
}

// RegisterMdmac10sarType MDMA channel x source address register
type RegisterMdmac10sarType uint32

func (r *RegisterMdmac10sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10sarFieldSarShift = 0
	RegisterMdmac10sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac10sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10sarFieldSarMask) >> RegisterMdmac10sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac10sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10sarFieldSarMask)|(uint32(value)<<RegisterMdmac10sarFieldSarShift))
}

// RegisterMdmac10darType MDMA channel x destination address register
type RegisterMdmac10darType uint32

func (r *RegisterMdmac10darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10darFieldDarShift = 0
	RegisterMdmac10darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac10darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10darFieldDarMask) >> RegisterMdmac10darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac10darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10darFieldDarMask)|(uint32(value)<<RegisterMdmac10darFieldDarShift))
}

// RegisterMdmac10brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac10brurType uint32

func (r *RegisterMdmac10brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10brurFieldSuvShift = 0
	RegisterMdmac10brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac10brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10brurFieldSuvMask) >> RegisterMdmac10brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac10brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10brurFieldSuvMask)|(uint32(value)<<RegisterMdmac10brurFieldSuvShift))
}

const (
	RegisterMdmac10brurFieldDuvShift = 16
	RegisterMdmac10brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac10brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10brurFieldDuvMask) >> RegisterMdmac10brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac10brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10brurFieldDuvMask)|(uint32(value)<<RegisterMdmac10brurFieldDuvShift))
}

// RegisterMdmac10larType MDMA channel x Link Address register
type RegisterMdmac10larType uint32

func (r *RegisterMdmac10larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10larFieldLarShift = 0
	RegisterMdmac10larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac10larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10larFieldLarMask) >> RegisterMdmac10larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac10larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10larFieldLarMask)|(uint32(value)<<RegisterMdmac10larFieldLarShift))
}

// RegisterMdmac10tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac10tbrType uint32

func (r *RegisterMdmac10tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10tbrFieldTselShift = 0
	RegisterMdmac10tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac10tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tbrFieldTselMask) >> RegisterMdmac10tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac10tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tbrFieldTselMask)|(uint32(value)<<RegisterMdmac10tbrFieldTselShift))
}

const (
	RegisterMdmac10tbrFieldSbusShift = 16
	RegisterMdmac10tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac10tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac10tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac10tbrFieldDbusShift = 17
	RegisterMdmac10tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac10tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac10tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac10tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10tbrFieldDbusMask)
	}
}

// RegisterMdmac10marType MDMA channel x Mask address register
type RegisterMdmac10marType uint32

func (r *RegisterMdmac10marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10marFieldMarShift = 0
	RegisterMdmac10marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac10marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10marFieldMarMask) >> RegisterMdmac10marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac10marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10marFieldMarMask)|(uint32(value)<<RegisterMdmac10marFieldMarShift))
}

// RegisterMdmac10mdrType MDMA channel x Mask Data register
type RegisterMdmac10mdrType uint32

func (r *RegisterMdmac10mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac10mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac10mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac10mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac10mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac10mdrFieldMdrShift = 0
	RegisterMdmac10mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac10mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac10mdrFieldMdrMask) >> RegisterMdmac10mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac10mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac10mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac10mdrFieldMdrShift))
}

// RegisterMdmac11isrType MDMA channel x interrupt/status register
type RegisterMdmac11isrType uint32

func (r *RegisterMdmac11isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11isrFieldTeif11Shift = 0
	RegisterMdmac11isrFieldTeif11Mask  = 0x1
)

// GetTeif11 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac11isrType) GetTeif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11isrFieldTeif11Mask) != 0
}

// SetTeif11 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac11isrType) SetTeif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11isrFieldTeif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11isrFieldTeif11Mask)
	}
}

const (
	RegisterMdmac11isrFieldCtcif11Shift = 1
	RegisterMdmac11isrFieldCtcif11Mask  = 0x2
)

// GetCtcif11 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac11isrType) GetCtcif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11isrFieldCtcif11Mask) != 0
}

// SetCtcif11 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac11isrType) SetCtcif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11isrFieldCtcif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11isrFieldCtcif11Mask)
	}
}

const (
	RegisterMdmac11isrFieldBrtif11Shift = 2
	RegisterMdmac11isrFieldBrtif11Mask  = 0x4
)

// GetBrtif11 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac11isrType) GetBrtif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11isrFieldBrtif11Mask) != 0
}

// SetBrtif11 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac11isrType) SetBrtif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11isrFieldBrtif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11isrFieldBrtif11Mask)
	}
}

const (
	RegisterMdmac11isrFieldBtif11Shift = 3
	RegisterMdmac11isrFieldBtif11Mask  = 0x8
)

// GetBtif11 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac11isrType) GetBtif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11isrFieldBtif11Mask) != 0
}

// SetBtif11 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac11isrType) SetBtif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11isrFieldBtif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11isrFieldBtif11Mask)
	}
}

const (
	RegisterMdmac11isrFieldTcif11Shift = 4
	RegisterMdmac11isrFieldTcif11Mask  = 0x10
)

// GetTcif11 channel x buffer transfer complete
func (r *RegisterMdmac11isrType) GetTcif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11isrFieldTcif11Mask) != 0
}

// SetTcif11 channel x buffer transfer complete
func (r *RegisterMdmac11isrType) SetTcif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11isrFieldTcif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11isrFieldTcif11Mask)
	}
}

const (
	RegisterMdmac11isrFieldCrqa11Shift = 16
	RegisterMdmac11isrFieldCrqa11Mask  = 0x10000
)

// GetCrqa11 channel x request active flag
func (r *RegisterMdmac11isrType) GetCrqa11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11isrFieldCrqa11Mask) != 0
}

// SetCrqa11 channel x request active flag
func (r *RegisterMdmac11isrType) SetCrqa11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11isrFieldCrqa11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11isrFieldCrqa11Mask)
	}
}

// RegisterMdmac11ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac11ifcrType uint32

func (r *RegisterMdmac11ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11ifcrFieldCteif11Shift = 0
	RegisterMdmac11ifcrFieldCteif11Mask  = 0x1
)

// GetCteif11 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac11ifcrType) GetCteif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11ifcrFieldCteif11Mask) != 0
}

// SetCteif11 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac11ifcrType) SetCteif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11ifcrFieldCteif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11ifcrFieldCteif11Mask)
	}
}

const (
	RegisterMdmac11ifcrFieldCctcif11Shift = 1
	RegisterMdmac11ifcrFieldCctcif11Mask  = 0x2
)

// GetCctcif11 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac11ifcrType) GetCctcif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11ifcrFieldCctcif11Mask) != 0
}

// SetCctcif11 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac11ifcrType) SetCctcif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11ifcrFieldCctcif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11ifcrFieldCctcif11Mask)
	}
}

const (
	RegisterMdmac11ifcrFieldCbrtif11Shift = 2
	RegisterMdmac11ifcrFieldCbrtif11Mask  = 0x4
)

// GetCbrtif11 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac11ifcrType) GetCbrtif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11ifcrFieldCbrtif11Mask) != 0
}

// SetCbrtif11 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac11ifcrType) SetCbrtif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11ifcrFieldCbrtif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11ifcrFieldCbrtif11Mask)
	}
}

const (
	RegisterMdmac11ifcrFieldCbtif11Shift = 3
	RegisterMdmac11ifcrFieldCbtif11Mask  = 0x8
)

// GetCbtif11 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac11ifcrType) GetCbtif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11ifcrFieldCbtif11Mask) != 0
}

// SetCbtif11 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac11ifcrType) SetCbtif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11ifcrFieldCbtif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11ifcrFieldCbtif11Mask)
	}
}

const (
	RegisterMdmac11ifcrFieldCltcif11Shift = 4
	RegisterMdmac11ifcrFieldCltcif11Mask  = 0x10
)

// GetCltcif11 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac11ifcrType) GetCltcif11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11ifcrFieldCltcif11Mask) != 0
}

// SetCltcif11 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac11ifcrType) SetCltcif11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11ifcrFieldCltcif11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11ifcrFieldCltcif11Mask)
	}
}

// RegisterMdmac11esrType MDMA Channel x error status register
type RegisterMdmac11esrType uint32

func (r *RegisterMdmac11esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11esrFieldTeaShift = 0
	RegisterMdmac11esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac11esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11esrFieldTeaMask) >> RegisterMdmac11esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac11esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11esrFieldTeaMask)|(uint32(value)<<RegisterMdmac11esrFieldTeaShift))
}

const (
	RegisterMdmac11esrFieldTedShift = 7
	RegisterMdmac11esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac11esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac11esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11esrFieldTedMask)
	}
}

const (
	RegisterMdmac11esrFieldTeldShift = 8
	RegisterMdmac11esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac11esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac11esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11esrFieldTeldMask)
	}
}

const (
	RegisterMdmac11esrFieldTemdShift = 9
	RegisterMdmac11esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac11esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac11esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11esrFieldTemdMask)
	}
}

const (
	RegisterMdmac11esrFieldAseShift = 10
	RegisterMdmac11esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac11esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac11esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11esrFieldAseMask)
	}
}

const (
	RegisterMdmac11esrFieldBseShift = 11
	RegisterMdmac11esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac11esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac11esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11esrFieldBseMask)
	}
}

// RegisterMdmac11crType This register is used to control the concerned channel.
type RegisterMdmac11crType uint32

func (r *RegisterMdmac11crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11crFieldEnShift = 0
	RegisterMdmac11crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac11crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac11crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11crFieldEnMask)
	}
}

const (
	RegisterMdmac11crFieldTeieShift = 1
	RegisterMdmac11crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac11crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac11crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11crFieldTeieMask)
	}
}

const (
	RegisterMdmac11crFieldCtcieShift = 2
	RegisterMdmac11crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac11crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac11crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11crFieldCtcieMask)
	}
}

const (
	RegisterMdmac11crFieldBrtieShift = 3
	RegisterMdmac11crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac11crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac11crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11crFieldBrtieMask)
	}
}

const (
	RegisterMdmac11crFieldBtieShift = 4
	RegisterMdmac11crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac11crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac11crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11crFieldBtieMask)
	}
}

const (
	RegisterMdmac11crFieldTcieShift = 5
	RegisterMdmac11crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac11crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac11crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11crFieldTcieMask)
	}
}

const (
	RegisterMdmac11crFieldPlShift = 6
	RegisterMdmac11crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac11crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11crFieldPlMask) >> RegisterMdmac11crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac11crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11crFieldPlMask)|(uint32(value)<<RegisterMdmac11crFieldPlShift))
}

const (
	RegisterMdmac11crFieldBexShift = 12
	RegisterMdmac11crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac11crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac11crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11crFieldBexMask)
	}
}

const (
	RegisterMdmac11crFieldHexShift = 13
	RegisterMdmac11crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac11crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac11crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11crFieldHexMask)
	}
}

const (
	RegisterMdmac11crFieldWexShift = 14
	RegisterMdmac11crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac11crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac11crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11crFieldWexMask)
	}
}

const (
	RegisterMdmac11crFieldSwrqShift = 16
	RegisterMdmac11crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac11crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11crFieldSwrqMask)
	}
}

// RegisterMdmac11tcrType This register is used to configure the concerned channel.
type RegisterMdmac11tcrType uint32

func (r *RegisterMdmac11tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11tcrFieldSincShift = 0
	RegisterMdmac11tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac11tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldSincMask) >> RegisterMdmac11tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac11tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldSincMask)|(uint32(value)<<RegisterMdmac11tcrFieldSincShift))
}

const (
	RegisterMdmac11tcrFieldDincShift = 2
	RegisterMdmac11tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac11tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldDincMask) >> RegisterMdmac11tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac11tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldDincMask)|(uint32(value)<<RegisterMdmac11tcrFieldDincShift))
}

const (
	RegisterMdmac11tcrFieldSsizeShift = 4
	RegisterMdmac11tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac11tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldSsizeMask) >> RegisterMdmac11tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac11tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac11tcrFieldSsizeShift))
}

const (
	RegisterMdmac11tcrFieldDsizeShift = 6
	RegisterMdmac11tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac11tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldDsizeMask) >> RegisterMdmac11tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac11tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac11tcrFieldDsizeShift))
}

const (
	RegisterMdmac11tcrFieldSincosShift = 8
	RegisterMdmac11tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac11tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldSincosMask) >> RegisterMdmac11tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac11tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac11tcrFieldSincosShift))
}

const (
	RegisterMdmac11tcrFieldDincosShift = 10
	RegisterMdmac11tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac11tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldDincosMask) >> RegisterMdmac11tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac11tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac11tcrFieldDincosShift))
}

const (
	RegisterMdmac11tcrFieldSburstShift = 12
	RegisterMdmac11tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac11tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldSburstMask) >> RegisterMdmac11tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac11tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac11tcrFieldSburstShift))
}

const (
	RegisterMdmac11tcrFieldDburstShift = 15
	RegisterMdmac11tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac11tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldDburstMask) >> RegisterMdmac11tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac11tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac11tcrFieldDburstShift))
}

const (
	RegisterMdmac11tcrFieldTlenShift = 18
	RegisterMdmac11tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac11tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldTlenMask) >> RegisterMdmac11tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac11tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac11tcrFieldTlenShift))
}

const (
	RegisterMdmac11tcrFieldPkeShift = 25
	RegisterMdmac11tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac11tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac11tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac11tcrFieldPamShift = 26
	RegisterMdmac11tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac11tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldPamMask) >> RegisterMdmac11tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac11tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldPamMask)|(uint32(value)<<RegisterMdmac11tcrFieldPamShift))
}

const (
	RegisterMdmac11tcrFieldTrgmShift = 28
	RegisterMdmac11tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac11tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldTrgmMask) >> RegisterMdmac11tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac11tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac11tcrFieldTrgmShift))
}

const (
	RegisterMdmac11tcrFieldSwrmShift = 30
	RegisterMdmac11tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac11tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac11tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac11tcrFieldBwmShift = 31
	RegisterMdmac11tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac11tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac11tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tcrFieldBwmMask)
	}
}

// RegisterMdmac11bndtrType MDMA Channel x block number of data register
type RegisterMdmac11bndtrType uint32

func (r *RegisterMdmac11bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11bndtrFieldBndtShift = 0
	RegisterMdmac11bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac11bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11bndtrFieldBndtMask) >> RegisterMdmac11bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac11bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac11bndtrFieldBndtShift))
}

const (
	RegisterMdmac11bndtrFieldBrsumShift = 18
	RegisterMdmac11bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac11bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac11bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac11bndtrFieldBrdumShift = 19
	RegisterMdmac11bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac11bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac11bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac11bndtrFieldBrcShift = 20
	RegisterMdmac11bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac11bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11bndtrFieldBrcMask) >> RegisterMdmac11bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac11bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac11bndtrFieldBrcShift))
}

// RegisterMdmac11sarType MDMA channel x source address register
type RegisterMdmac11sarType uint32

func (r *RegisterMdmac11sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11sarFieldSarShift = 0
	RegisterMdmac11sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac11sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11sarFieldSarMask) >> RegisterMdmac11sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac11sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11sarFieldSarMask)|(uint32(value)<<RegisterMdmac11sarFieldSarShift))
}

// RegisterMdmac11darType MDMA channel x destination address register
type RegisterMdmac11darType uint32

func (r *RegisterMdmac11darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11darFieldDarShift = 0
	RegisterMdmac11darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac11darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11darFieldDarMask) >> RegisterMdmac11darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac11darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11darFieldDarMask)|(uint32(value)<<RegisterMdmac11darFieldDarShift))
}

// RegisterMdmac11brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac11brurType uint32

func (r *RegisterMdmac11brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11brurFieldSuvShift = 0
	RegisterMdmac11brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac11brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11brurFieldSuvMask) >> RegisterMdmac11brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac11brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11brurFieldSuvMask)|(uint32(value)<<RegisterMdmac11brurFieldSuvShift))
}

const (
	RegisterMdmac11brurFieldDuvShift = 16
	RegisterMdmac11brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac11brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11brurFieldDuvMask) >> RegisterMdmac11brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac11brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11brurFieldDuvMask)|(uint32(value)<<RegisterMdmac11brurFieldDuvShift))
}

// RegisterMdmac11larType MDMA channel x Link Address register
type RegisterMdmac11larType uint32

func (r *RegisterMdmac11larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11larFieldLarShift = 0
	RegisterMdmac11larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac11larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11larFieldLarMask) >> RegisterMdmac11larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac11larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11larFieldLarMask)|(uint32(value)<<RegisterMdmac11larFieldLarShift))
}

// RegisterMdmac11tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac11tbrType uint32

func (r *RegisterMdmac11tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11tbrFieldTselShift = 0
	RegisterMdmac11tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac11tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tbrFieldTselMask) >> RegisterMdmac11tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac11tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tbrFieldTselMask)|(uint32(value)<<RegisterMdmac11tbrFieldTselShift))
}

const (
	RegisterMdmac11tbrFieldSbusShift = 16
	RegisterMdmac11tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac11tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac11tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac11tbrFieldDbusShift = 17
	RegisterMdmac11tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac11tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac11tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac11tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11tbrFieldDbusMask)
	}
}

// RegisterMdmac11marType MDMA channel x Mask address register
type RegisterMdmac11marType uint32

func (r *RegisterMdmac11marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11marFieldMarShift = 0
	RegisterMdmac11marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac11marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11marFieldMarMask) >> RegisterMdmac11marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac11marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11marFieldMarMask)|(uint32(value)<<RegisterMdmac11marFieldMarShift))
}

// RegisterMdmac11mdrType MDMA channel x Mask Data register
type RegisterMdmac11mdrType uint32

func (r *RegisterMdmac11mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac11mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac11mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac11mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac11mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac11mdrFieldMdrShift = 0
	RegisterMdmac11mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac11mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac11mdrFieldMdrMask) >> RegisterMdmac11mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac11mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac11mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac11mdrFieldMdrShift))
}

// RegisterMdmac12isrType MDMA channel x interrupt/status register
type RegisterMdmac12isrType uint32

func (r *RegisterMdmac12isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12isrFieldTeif12Shift = 0
	RegisterMdmac12isrFieldTeif12Mask  = 0x1
)

// GetTeif12 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac12isrType) GetTeif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12isrFieldTeif12Mask) != 0
}

// SetTeif12 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac12isrType) SetTeif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12isrFieldTeif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12isrFieldTeif12Mask)
	}
}

const (
	RegisterMdmac12isrFieldCtcif12Shift = 1
	RegisterMdmac12isrFieldCtcif12Mask  = 0x2
)

// GetCtcif12 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac12isrType) GetCtcif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12isrFieldCtcif12Mask) != 0
}

// SetCtcif12 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac12isrType) SetCtcif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12isrFieldCtcif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12isrFieldCtcif12Mask)
	}
}

const (
	RegisterMdmac12isrFieldBrtif12Shift = 2
	RegisterMdmac12isrFieldBrtif12Mask  = 0x4
)

// GetBrtif12 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac12isrType) GetBrtif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12isrFieldBrtif12Mask) != 0
}

// SetBrtif12 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac12isrType) SetBrtif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12isrFieldBrtif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12isrFieldBrtif12Mask)
	}
}

const (
	RegisterMdmac12isrFieldBtif12Shift = 3
	RegisterMdmac12isrFieldBtif12Mask  = 0x8
)

// GetBtif12 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac12isrType) GetBtif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12isrFieldBtif12Mask) != 0
}

// SetBtif12 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac12isrType) SetBtif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12isrFieldBtif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12isrFieldBtif12Mask)
	}
}

const (
	RegisterMdmac12isrFieldTcif12Shift = 4
	RegisterMdmac12isrFieldTcif12Mask  = 0x10
)

// GetTcif12 channel x buffer transfer complete
func (r *RegisterMdmac12isrType) GetTcif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12isrFieldTcif12Mask) != 0
}

// SetTcif12 channel x buffer transfer complete
func (r *RegisterMdmac12isrType) SetTcif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12isrFieldTcif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12isrFieldTcif12Mask)
	}
}

const (
	RegisterMdmac12isrFieldCrqa12Shift = 16
	RegisterMdmac12isrFieldCrqa12Mask  = 0x10000
)

// GetCrqa12 channel x request active flag
func (r *RegisterMdmac12isrType) GetCrqa12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12isrFieldCrqa12Mask) != 0
}

// SetCrqa12 channel x request active flag
func (r *RegisterMdmac12isrType) SetCrqa12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12isrFieldCrqa12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12isrFieldCrqa12Mask)
	}
}

// RegisterMdmac12ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac12ifcrType uint32

func (r *RegisterMdmac12ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12ifcrFieldCteif12Shift = 0
	RegisterMdmac12ifcrFieldCteif12Mask  = 0x1
)

// GetCteif12 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac12ifcrType) GetCteif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12ifcrFieldCteif12Mask) != 0
}

// SetCteif12 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac12ifcrType) SetCteif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12ifcrFieldCteif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12ifcrFieldCteif12Mask)
	}
}

const (
	RegisterMdmac12ifcrFieldCctcif12Shift = 1
	RegisterMdmac12ifcrFieldCctcif12Mask  = 0x2
)

// GetCctcif12 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac12ifcrType) GetCctcif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12ifcrFieldCctcif12Mask) != 0
}

// SetCctcif12 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac12ifcrType) SetCctcif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12ifcrFieldCctcif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12ifcrFieldCctcif12Mask)
	}
}

const (
	RegisterMdmac12ifcrFieldCbrtif12Shift = 2
	RegisterMdmac12ifcrFieldCbrtif12Mask  = 0x4
)

// GetCbrtif12 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac12ifcrType) GetCbrtif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12ifcrFieldCbrtif12Mask) != 0
}

// SetCbrtif12 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac12ifcrType) SetCbrtif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12ifcrFieldCbrtif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12ifcrFieldCbrtif12Mask)
	}
}

const (
	RegisterMdmac12ifcrFieldCbtif12Shift = 3
	RegisterMdmac12ifcrFieldCbtif12Mask  = 0x8
)

// GetCbtif12 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac12ifcrType) GetCbtif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12ifcrFieldCbtif12Mask) != 0
}

// SetCbtif12 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac12ifcrType) SetCbtif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12ifcrFieldCbtif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12ifcrFieldCbtif12Mask)
	}
}

const (
	RegisterMdmac12ifcrFieldCltcif12Shift = 4
	RegisterMdmac12ifcrFieldCltcif12Mask  = 0x10
)

// GetCltcif12 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac12ifcrType) GetCltcif12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12ifcrFieldCltcif12Mask) != 0
}

// SetCltcif12 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac12ifcrType) SetCltcif12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12ifcrFieldCltcif12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12ifcrFieldCltcif12Mask)
	}
}

// RegisterMdmac12esrType MDMA Channel x error status register
type RegisterMdmac12esrType uint32

func (r *RegisterMdmac12esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12esrFieldTeaShift = 0
	RegisterMdmac12esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac12esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12esrFieldTeaMask) >> RegisterMdmac12esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac12esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12esrFieldTeaMask)|(uint32(value)<<RegisterMdmac12esrFieldTeaShift))
}

const (
	RegisterMdmac12esrFieldTedShift = 7
	RegisterMdmac12esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac12esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac12esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12esrFieldTedMask)
	}
}

const (
	RegisterMdmac12esrFieldTeldShift = 8
	RegisterMdmac12esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac12esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac12esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12esrFieldTeldMask)
	}
}

const (
	RegisterMdmac12esrFieldTemdShift = 9
	RegisterMdmac12esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac12esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac12esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12esrFieldTemdMask)
	}
}

const (
	RegisterMdmac12esrFieldAseShift = 10
	RegisterMdmac12esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac12esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac12esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12esrFieldAseMask)
	}
}

const (
	RegisterMdmac12esrFieldBseShift = 11
	RegisterMdmac12esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac12esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac12esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12esrFieldBseMask)
	}
}

// RegisterMdmac12crType This register is used to control the concerned channel.
type RegisterMdmac12crType uint32

func (r *RegisterMdmac12crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12crFieldEnShift = 0
	RegisterMdmac12crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac12crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac12crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12crFieldEnMask)
	}
}

const (
	RegisterMdmac12crFieldTeieShift = 1
	RegisterMdmac12crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac12crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac12crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12crFieldTeieMask)
	}
}

const (
	RegisterMdmac12crFieldCtcieShift = 2
	RegisterMdmac12crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac12crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac12crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12crFieldCtcieMask)
	}
}

const (
	RegisterMdmac12crFieldBrtieShift = 3
	RegisterMdmac12crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac12crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac12crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12crFieldBrtieMask)
	}
}

const (
	RegisterMdmac12crFieldBtieShift = 4
	RegisterMdmac12crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac12crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac12crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12crFieldBtieMask)
	}
}

const (
	RegisterMdmac12crFieldTcieShift = 5
	RegisterMdmac12crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac12crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac12crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12crFieldTcieMask)
	}
}

const (
	RegisterMdmac12crFieldPlShift = 6
	RegisterMdmac12crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac12crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12crFieldPlMask) >> RegisterMdmac12crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac12crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12crFieldPlMask)|(uint32(value)<<RegisterMdmac12crFieldPlShift))
}

const (
	RegisterMdmac12crFieldBexShift = 12
	RegisterMdmac12crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac12crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac12crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12crFieldBexMask)
	}
}

const (
	RegisterMdmac12crFieldHexShift = 13
	RegisterMdmac12crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac12crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac12crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12crFieldHexMask)
	}
}

const (
	RegisterMdmac12crFieldWexShift = 14
	RegisterMdmac12crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac12crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac12crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12crFieldWexMask)
	}
}

const (
	RegisterMdmac12crFieldSwrqShift = 16
	RegisterMdmac12crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac12crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12crFieldSwrqMask)
	}
}

// RegisterMdmac12tcrType This register is used to configure the concerned channel.
type RegisterMdmac12tcrType uint32

func (r *RegisterMdmac12tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12tcrFieldSincShift = 0
	RegisterMdmac12tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac12tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldSincMask) >> RegisterMdmac12tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac12tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldSincMask)|(uint32(value)<<RegisterMdmac12tcrFieldSincShift))
}

const (
	RegisterMdmac12tcrFieldDincShift = 2
	RegisterMdmac12tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac12tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldDincMask) >> RegisterMdmac12tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac12tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldDincMask)|(uint32(value)<<RegisterMdmac12tcrFieldDincShift))
}

const (
	RegisterMdmac12tcrFieldSsizeShift = 4
	RegisterMdmac12tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac12tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldSsizeMask) >> RegisterMdmac12tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac12tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac12tcrFieldSsizeShift))
}

const (
	RegisterMdmac12tcrFieldDsizeShift = 6
	RegisterMdmac12tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac12tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldDsizeMask) >> RegisterMdmac12tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac12tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac12tcrFieldDsizeShift))
}

const (
	RegisterMdmac12tcrFieldSincosShift = 8
	RegisterMdmac12tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac12tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldSincosMask) >> RegisterMdmac12tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac12tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac12tcrFieldSincosShift))
}

const (
	RegisterMdmac12tcrFieldDincosShift = 10
	RegisterMdmac12tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac12tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldDincosMask) >> RegisterMdmac12tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac12tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac12tcrFieldDincosShift))
}

const (
	RegisterMdmac12tcrFieldSburstShift = 12
	RegisterMdmac12tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac12tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldSburstMask) >> RegisterMdmac12tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac12tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac12tcrFieldSburstShift))
}

const (
	RegisterMdmac12tcrFieldDburstShift = 15
	RegisterMdmac12tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac12tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldDburstMask) >> RegisterMdmac12tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac12tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac12tcrFieldDburstShift))
}

const (
	RegisterMdmac12tcrFieldTlenShift = 18
	RegisterMdmac12tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac12tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldTlenMask) >> RegisterMdmac12tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac12tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac12tcrFieldTlenShift))
}

const (
	RegisterMdmac12tcrFieldPkeShift = 25
	RegisterMdmac12tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac12tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac12tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac12tcrFieldPamShift = 26
	RegisterMdmac12tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac12tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldPamMask) >> RegisterMdmac12tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac12tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldPamMask)|(uint32(value)<<RegisterMdmac12tcrFieldPamShift))
}

const (
	RegisterMdmac12tcrFieldTrgmShift = 28
	RegisterMdmac12tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac12tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldTrgmMask) >> RegisterMdmac12tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac12tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac12tcrFieldTrgmShift))
}

const (
	RegisterMdmac12tcrFieldSwrmShift = 30
	RegisterMdmac12tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac12tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac12tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac12tcrFieldBwmShift = 31
	RegisterMdmac12tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac12tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac12tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tcrFieldBwmMask)
	}
}

// RegisterMdmac12bndtrType MDMA Channel x block number of data register
type RegisterMdmac12bndtrType uint32

func (r *RegisterMdmac12bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12bndtrFieldBndtShift = 0
	RegisterMdmac12bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac12bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12bndtrFieldBndtMask) >> RegisterMdmac12bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac12bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac12bndtrFieldBndtShift))
}

const (
	RegisterMdmac12bndtrFieldBrsumShift = 18
	RegisterMdmac12bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac12bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac12bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac12bndtrFieldBrdumShift = 19
	RegisterMdmac12bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac12bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac12bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac12bndtrFieldBrcShift = 20
	RegisterMdmac12bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac12bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12bndtrFieldBrcMask) >> RegisterMdmac12bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac12bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac12bndtrFieldBrcShift))
}

// RegisterMdmac12sarType MDMA channel x source address register
type RegisterMdmac12sarType uint32

func (r *RegisterMdmac12sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12sarFieldSarShift = 0
	RegisterMdmac12sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac12sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12sarFieldSarMask) >> RegisterMdmac12sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac12sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12sarFieldSarMask)|(uint32(value)<<RegisterMdmac12sarFieldSarShift))
}

// RegisterMdmac12darType MDMA channel x destination address register
type RegisterMdmac12darType uint32

func (r *RegisterMdmac12darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12darFieldDarShift = 0
	RegisterMdmac12darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac12darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12darFieldDarMask) >> RegisterMdmac12darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac12darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12darFieldDarMask)|(uint32(value)<<RegisterMdmac12darFieldDarShift))
}

// RegisterMdmac12brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac12brurType uint32

func (r *RegisterMdmac12brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12brurFieldSuvShift = 0
	RegisterMdmac12brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac12brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12brurFieldSuvMask) >> RegisterMdmac12brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac12brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12brurFieldSuvMask)|(uint32(value)<<RegisterMdmac12brurFieldSuvShift))
}

const (
	RegisterMdmac12brurFieldDuvShift = 16
	RegisterMdmac12brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac12brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12brurFieldDuvMask) >> RegisterMdmac12brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac12brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12brurFieldDuvMask)|(uint32(value)<<RegisterMdmac12brurFieldDuvShift))
}

// RegisterMdmac12larType MDMA channel x Link Address register
type RegisterMdmac12larType uint32

func (r *RegisterMdmac12larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12larFieldLarShift = 0
	RegisterMdmac12larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac12larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12larFieldLarMask) >> RegisterMdmac12larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac12larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12larFieldLarMask)|(uint32(value)<<RegisterMdmac12larFieldLarShift))
}

// RegisterMdmac12tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac12tbrType uint32

func (r *RegisterMdmac12tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12tbrFieldTselShift = 0
	RegisterMdmac12tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac12tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tbrFieldTselMask) >> RegisterMdmac12tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac12tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tbrFieldTselMask)|(uint32(value)<<RegisterMdmac12tbrFieldTselShift))
}

const (
	RegisterMdmac12tbrFieldSbusShift = 16
	RegisterMdmac12tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac12tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac12tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac12tbrFieldDbusShift = 17
	RegisterMdmac12tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac12tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac12tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac12tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12tbrFieldDbusMask)
	}
}

// RegisterMdmac12marType MDMA channel x Mask address register
type RegisterMdmac12marType uint32

func (r *RegisterMdmac12marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12marFieldMarShift = 0
	RegisterMdmac12marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac12marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12marFieldMarMask) >> RegisterMdmac12marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac12marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12marFieldMarMask)|(uint32(value)<<RegisterMdmac12marFieldMarShift))
}

// RegisterMdmac12mdrType MDMA channel x Mask Data register
type RegisterMdmac12mdrType uint32

func (r *RegisterMdmac12mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac12mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac12mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac12mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac12mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac12mdrFieldMdrShift = 0
	RegisterMdmac12mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac12mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac12mdrFieldMdrMask) >> RegisterMdmac12mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac12mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac12mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac12mdrFieldMdrShift))
}

// RegisterMdmac13isrType MDMA channel x interrupt/status register
type RegisterMdmac13isrType uint32

func (r *RegisterMdmac13isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13isrFieldTeif13Shift = 0
	RegisterMdmac13isrFieldTeif13Mask  = 0x1
)

// GetTeif13 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac13isrType) GetTeif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13isrFieldTeif13Mask) != 0
}

// SetTeif13 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac13isrType) SetTeif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13isrFieldTeif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13isrFieldTeif13Mask)
	}
}

const (
	RegisterMdmac13isrFieldCtcif13Shift = 1
	RegisterMdmac13isrFieldCtcif13Mask  = 0x2
)

// GetCtcif13 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac13isrType) GetCtcif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13isrFieldCtcif13Mask) != 0
}

// SetCtcif13 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac13isrType) SetCtcif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13isrFieldCtcif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13isrFieldCtcif13Mask)
	}
}

const (
	RegisterMdmac13isrFieldBrtif13Shift = 2
	RegisterMdmac13isrFieldBrtif13Mask  = 0x4
)

// GetBrtif13 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac13isrType) GetBrtif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13isrFieldBrtif13Mask) != 0
}

// SetBrtif13 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac13isrType) SetBrtif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13isrFieldBrtif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13isrFieldBrtif13Mask)
	}
}

const (
	RegisterMdmac13isrFieldBtif13Shift = 3
	RegisterMdmac13isrFieldBtif13Mask  = 0x8
)

// GetBtif13 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac13isrType) GetBtif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13isrFieldBtif13Mask) != 0
}

// SetBtif13 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac13isrType) SetBtif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13isrFieldBtif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13isrFieldBtif13Mask)
	}
}

const (
	RegisterMdmac13isrFieldTcif13Shift = 4
	RegisterMdmac13isrFieldTcif13Mask  = 0x10
)

// GetTcif13 channel x buffer transfer complete
func (r *RegisterMdmac13isrType) GetTcif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13isrFieldTcif13Mask) != 0
}

// SetTcif13 channel x buffer transfer complete
func (r *RegisterMdmac13isrType) SetTcif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13isrFieldTcif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13isrFieldTcif13Mask)
	}
}

const (
	RegisterMdmac13isrFieldCrqa13Shift = 16
	RegisterMdmac13isrFieldCrqa13Mask  = 0x10000
)

// GetCrqa13 channel x request active flag
func (r *RegisterMdmac13isrType) GetCrqa13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13isrFieldCrqa13Mask) != 0
}

// SetCrqa13 channel x request active flag
func (r *RegisterMdmac13isrType) SetCrqa13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13isrFieldCrqa13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13isrFieldCrqa13Mask)
	}
}

// RegisterMdmac13ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac13ifcrType uint32

func (r *RegisterMdmac13ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13ifcrFieldCteif13Shift = 0
	RegisterMdmac13ifcrFieldCteif13Mask  = 0x1
)

// GetCteif13 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac13ifcrType) GetCteif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13ifcrFieldCteif13Mask) != 0
}

// SetCteif13 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac13ifcrType) SetCteif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13ifcrFieldCteif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13ifcrFieldCteif13Mask)
	}
}

const (
	RegisterMdmac13ifcrFieldCctcif13Shift = 1
	RegisterMdmac13ifcrFieldCctcif13Mask  = 0x2
)

// GetCctcif13 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac13ifcrType) GetCctcif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13ifcrFieldCctcif13Mask) != 0
}

// SetCctcif13 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac13ifcrType) SetCctcif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13ifcrFieldCctcif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13ifcrFieldCctcif13Mask)
	}
}

const (
	RegisterMdmac13ifcrFieldCbrtif13Shift = 2
	RegisterMdmac13ifcrFieldCbrtif13Mask  = 0x4
)

// GetCbrtif13 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac13ifcrType) GetCbrtif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13ifcrFieldCbrtif13Mask) != 0
}

// SetCbrtif13 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac13ifcrType) SetCbrtif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13ifcrFieldCbrtif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13ifcrFieldCbrtif13Mask)
	}
}

const (
	RegisterMdmac13ifcrFieldCbtif13Shift = 3
	RegisterMdmac13ifcrFieldCbtif13Mask  = 0x8
)

// GetCbtif13 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac13ifcrType) GetCbtif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13ifcrFieldCbtif13Mask) != 0
}

// SetCbtif13 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac13ifcrType) SetCbtif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13ifcrFieldCbtif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13ifcrFieldCbtif13Mask)
	}
}

const (
	RegisterMdmac13ifcrFieldCltcif13Shift = 4
	RegisterMdmac13ifcrFieldCltcif13Mask  = 0x10
)

// GetCltcif13 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac13ifcrType) GetCltcif13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13ifcrFieldCltcif13Mask) != 0
}

// SetCltcif13 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac13ifcrType) SetCltcif13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13ifcrFieldCltcif13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13ifcrFieldCltcif13Mask)
	}
}

// RegisterMdmac13esrType MDMA Channel x error status register
type RegisterMdmac13esrType uint32

func (r *RegisterMdmac13esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13esrFieldTeaShift = 0
	RegisterMdmac13esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac13esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13esrFieldTeaMask) >> RegisterMdmac13esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac13esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13esrFieldTeaMask)|(uint32(value)<<RegisterMdmac13esrFieldTeaShift))
}

const (
	RegisterMdmac13esrFieldTedShift = 7
	RegisterMdmac13esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac13esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac13esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13esrFieldTedMask)
	}
}

const (
	RegisterMdmac13esrFieldTeldShift = 8
	RegisterMdmac13esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac13esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac13esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13esrFieldTeldMask)
	}
}

const (
	RegisterMdmac13esrFieldTemdShift = 9
	RegisterMdmac13esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac13esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac13esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13esrFieldTemdMask)
	}
}

const (
	RegisterMdmac13esrFieldAseShift = 10
	RegisterMdmac13esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac13esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac13esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13esrFieldAseMask)
	}
}

const (
	RegisterMdmac13esrFieldBseShift = 11
	RegisterMdmac13esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac13esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac13esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13esrFieldBseMask)
	}
}

// RegisterMdmac13crType This register is used to control the concerned channel.
type RegisterMdmac13crType uint32

func (r *RegisterMdmac13crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13crFieldEnShift = 0
	RegisterMdmac13crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac13crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac13crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13crFieldEnMask)
	}
}

const (
	RegisterMdmac13crFieldTeieShift = 1
	RegisterMdmac13crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac13crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac13crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13crFieldTeieMask)
	}
}

const (
	RegisterMdmac13crFieldCtcieShift = 2
	RegisterMdmac13crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac13crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac13crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13crFieldCtcieMask)
	}
}

const (
	RegisterMdmac13crFieldBrtieShift = 3
	RegisterMdmac13crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac13crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac13crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13crFieldBrtieMask)
	}
}

const (
	RegisterMdmac13crFieldBtieShift = 4
	RegisterMdmac13crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac13crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac13crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13crFieldBtieMask)
	}
}

const (
	RegisterMdmac13crFieldTcieShift = 5
	RegisterMdmac13crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac13crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac13crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13crFieldTcieMask)
	}
}

const (
	RegisterMdmac13crFieldPlShift = 6
	RegisterMdmac13crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac13crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13crFieldPlMask) >> RegisterMdmac13crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac13crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13crFieldPlMask)|(uint32(value)<<RegisterMdmac13crFieldPlShift))
}

const (
	RegisterMdmac13crFieldBexShift = 12
	RegisterMdmac13crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac13crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac13crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13crFieldBexMask)
	}
}

const (
	RegisterMdmac13crFieldHexShift = 13
	RegisterMdmac13crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac13crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac13crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13crFieldHexMask)
	}
}

const (
	RegisterMdmac13crFieldWexShift = 14
	RegisterMdmac13crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac13crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac13crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13crFieldWexMask)
	}
}

const (
	RegisterMdmac13crFieldSwrqShift = 16
	RegisterMdmac13crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac13crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13crFieldSwrqMask)
	}
}

// RegisterMdmac13tcrType This register is used to configure the concerned channel.
type RegisterMdmac13tcrType uint32

func (r *RegisterMdmac13tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13tcrFieldSincShift = 0
	RegisterMdmac13tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac13tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldSincMask) >> RegisterMdmac13tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac13tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldSincMask)|(uint32(value)<<RegisterMdmac13tcrFieldSincShift))
}

const (
	RegisterMdmac13tcrFieldDincShift = 2
	RegisterMdmac13tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac13tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldDincMask) >> RegisterMdmac13tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac13tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldDincMask)|(uint32(value)<<RegisterMdmac13tcrFieldDincShift))
}

const (
	RegisterMdmac13tcrFieldSsizeShift = 4
	RegisterMdmac13tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac13tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldSsizeMask) >> RegisterMdmac13tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac13tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac13tcrFieldSsizeShift))
}

const (
	RegisterMdmac13tcrFieldDsizeShift = 6
	RegisterMdmac13tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac13tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldDsizeMask) >> RegisterMdmac13tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac13tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac13tcrFieldDsizeShift))
}

const (
	RegisterMdmac13tcrFieldSincosShift = 8
	RegisterMdmac13tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac13tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldSincosMask) >> RegisterMdmac13tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac13tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac13tcrFieldSincosShift))
}

const (
	RegisterMdmac13tcrFieldDincosShift = 10
	RegisterMdmac13tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac13tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldDincosMask) >> RegisterMdmac13tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac13tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac13tcrFieldDincosShift))
}

const (
	RegisterMdmac13tcrFieldSburstShift = 12
	RegisterMdmac13tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac13tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldSburstMask) >> RegisterMdmac13tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac13tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac13tcrFieldSburstShift))
}

const (
	RegisterMdmac13tcrFieldDburstShift = 15
	RegisterMdmac13tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac13tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldDburstMask) >> RegisterMdmac13tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac13tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac13tcrFieldDburstShift))
}

const (
	RegisterMdmac13tcrFieldTlenShift = 18
	RegisterMdmac13tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac13tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldTlenMask) >> RegisterMdmac13tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac13tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac13tcrFieldTlenShift))
}

const (
	RegisterMdmac13tcrFieldPkeShift = 25
	RegisterMdmac13tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac13tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac13tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac13tcrFieldPamShift = 26
	RegisterMdmac13tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac13tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldPamMask) >> RegisterMdmac13tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac13tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldPamMask)|(uint32(value)<<RegisterMdmac13tcrFieldPamShift))
}

const (
	RegisterMdmac13tcrFieldTrgmShift = 28
	RegisterMdmac13tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac13tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldTrgmMask) >> RegisterMdmac13tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac13tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac13tcrFieldTrgmShift))
}

const (
	RegisterMdmac13tcrFieldSwrmShift = 30
	RegisterMdmac13tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac13tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac13tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac13tcrFieldBwmShift = 31
	RegisterMdmac13tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac13tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac13tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tcrFieldBwmMask)
	}
}

// RegisterMdmac13bndtrType MDMA Channel x block number of data register
type RegisterMdmac13bndtrType uint32

func (r *RegisterMdmac13bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13bndtrFieldBndtShift = 0
	RegisterMdmac13bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac13bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13bndtrFieldBndtMask) >> RegisterMdmac13bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac13bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac13bndtrFieldBndtShift))
}

const (
	RegisterMdmac13bndtrFieldBrsumShift = 18
	RegisterMdmac13bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac13bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac13bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac13bndtrFieldBrdumShift = 19
	RegisterMdmac13bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac13bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac13bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac13bndtrFieldBrcShift = 20
	RegisterMdmac13bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac13bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13bndtrFieldBrcMask) >> RegisterMdmac13bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac13bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac13bndtrFieldBrcShift))
}

// RegisterMdmac13sarType MDMA channel x source address register
type RegisterMdmac13sarType uint32

func (r *RegisterMdmac13sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13sarFieldSarShift = 0
	RegisterMdmac13sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac13sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13sarFieldSarMask) >> RegisterMdmac13sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac13sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13sarFieldSarMask)|(uint32(value)<<RegisterMdmac13sarFieldSarShift))
}

// RegisterMdmac13darType MDMA channel x destination address register
type RegisterMdmac13darType uint32

func (r *RegisterMdmac13darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13darFieldDarShift = 0
	RegisterMdmac13darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac13darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13darFieldDarMask) >> RegisterMdmac13darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac13darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13darFieldDarMask)|(uint32(value)<<RegisterMdmac13darFieldDarShift))
}

// RegisterMdmac13brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac13brurType uint32

func (r *RegisterMdmac13brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13brurFieldSuvShift = 0
	RegisterMdmac13brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac13brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13brurFieldSuvMask) >> RegisterMdmac13brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac13brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13brurFieldSuvMask)|(uint32(value)<<RegisterMdmac13brurFieldSuvShift))
}

const (
	RegisterMdmac13brurFieldDuvShift = 16
	RegisterMdmac13brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac13brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13brurFieldDuvMask) >> RegisterMdmac13brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac13brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13brurFieldDuvMask)|(uint32(value)<<RegisterMdmac13brurFieldDuvShift))
}

// RegisterMdmac13larType MDMA channel x Link Address register
type RegisterMdmac13larType uint32

func (r *RegisterMdmac13larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13larFieldLarShift = 0
	RegisterMdmac13larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac13larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13larFieldLarMask) >> RegisterMdmac13larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac13larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13larFieldLarMask)|(uint32(value)<<RegisterMdmac13larFieldLarShift))
}

// RegisterMdmac13tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac13tbrType uint32

func (r *RegisterMdmac13tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13tbrFieldTselShift = 0
	RegisterMdmac13tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac13tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tbrFieldTselMask) >> RegisterMdmac13tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac13tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tbrFieldTselMask)|(uint32(value)<<RegisterMdmac13tbrFieldTselShift))
}

const (
	RegisterMdmac13tbrFieldSbusShift = 16
	RegisterMdmac13tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac13tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac13tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac13tbrFieldDbusShift = 17
	RegisterMdmac13tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac13tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac13tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac13tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13tbrFieldDbusMask)
	}
}

// RegisterMdmac13marType MDMA channel x Mask address register
type RegisterMdmac13marType uint32

func (r *RegisterMdmac13marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13marFieldMarShift = 0
	RegisterMdmac13marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac13marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13marFieldMarMask) >> RegisterMdmac13marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac13marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13marFieldMarMask)|(uint32(value)<<RegisterMdmac13marFieldMarShift))
}

// RegisterMdmac13mdrType MDMA channel x Mask Data register
type RegisterMdmac13mdrType uint32

func (r *RegisterMdmac13mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac13mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac13mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac13mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac13mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac13mdrFieldMdrShift = 0
	RegisterMdmac13mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac13mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac13mdrFieldMdrMask) >> RegisterMdmac13mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac13mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac13mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac13mdrFieldMdrShift))
}

// RegisterMdmac14isrType MDMA channel x interrupt/status register
type RegisterMdmac14isrType uint32

func (r *RegisterMdmac14isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14isrFieldTeif14Shift = 0
	RegisterMdmac14isrFieldTeif14Mask  = 0x1
)

// GetTeif14 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac14isrType) GetTeif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14isrFieldTeif14Mask) != 0
}

// SetTeif14 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac14isrType) SetTeif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14isrFieldTeif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14isrFieldTeif14Mask)
	}
}

const (
	RegisterMdmac14isrFieldCtcif14Shift = 1
	RegisterMdmac14isrFieldCtcif14Mask  = 0x2
)

// GetCtcif14 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac14isrType) GetCtcif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14isrFieldCtcif14Mask) != 0
}

// SetCtcif14 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac14isrType) SetCtcif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14isrFieldCtcif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14isrFieldCtcif14Mask)
	}
}

const (
	RegisterMdmac14isrFieldBrtif14Shift = 2
	RegisterMdmac14isrFieldBrtif14Mask  = 0x4
)

// GetBrtif14 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac14isrType) GetBrtif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14isrFieldBrtif14Mask) != 0
}

// SetBrtif14 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac14isrType) SetBrtif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14isrFieldBrtif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14isrFieldBrtif14Mask)
	}
}

const (
	RegisterMdmac14isrFieldBtif14Shift = 3
	RegisterMdmac14isrFieldBtif14Mask  = 0x8
)

// GetBtif14 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac14isrType) GetBtif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14isrFieldBtif14Mask) != 0
}

// SetBtif14 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac14isrType) SetBtif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14isrFieldBtif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14isrFieldBtif14Mask)
	}
}

const (
	RegisterMdmac14isrFieldTcif14Shift = 4
	RegisterMdmac14isrFieldTcif14Mask  = 0x10
)

// GetTcif14 channel x buffer transfer complete
func (r *RegisterMdmac14isrType) GetTcif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14isrFieldTcif14Mask) != 0
}

// SetTcif14 channel x buffer transfer complete
func (r *RegisterMdmac14isrType) SetTcif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14isrFieldTcif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14isrFieldTcif14Mask)
	}
}

const (
	RegisterMdmac14isrFieldCrqa14Shift = 16
	RegisterMdmac14isrFieldCrqa14Mask  = 0x10000
)

// GetCrqa14 channel x request active flag
func (r *RegisterMdmac14isrType) GetCrqa14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14isrFieldCrqa14Mask) != 0
}

// SetCrqa14 channel x request active flag
func (r *RegisterMdmac14isrType) SetCrqa14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14isrFieldCrqa14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14isrFieldCrqa14Mask)
	}
}

// RegisterMdmac14ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac14ifcrType uint32

func (r *RegisterMdmac14ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14ifcrFieldCteif14Shift = 0
	RegisterMdmac14ifcrFieldCteif14Mask  = 0x1
)

// GetCteif14 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac14ifcrType) GetCteif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14ifcrFieldCteif14Mask) != 0
}

// SetCteif14 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac14ifcrType) SetCteif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14ifcrFieldCteif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14ifcrFieldCteif14Mask)
	}
}

const (
	RegisterMdmac14ifcrFieldCctcif14Shift = 1
	RegisterMdmac14ifcrFieldCctcif14Mask  = 0x2
)

// GetCctcif14 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac14ifcrType) GetCctcif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14ifcrFieldCctcif14Mask) != 0
}

// SetCctcif14 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac14ifcrType) SetCctcif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14ifcrFieldCctcif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14ifcrFieldCctcif14Mask)
	}
}

const (
	RegisterMdmac14ifcrFieldCbrtif14Shift = 2
	RegisterMdmac14ifcrFieldCbrtif14Mask  = 0x4
)

// GetCbrtif14 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac14ifcrType) GetCbrtif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14ifcrFieldCbrtif14Mask) != 0
}

// SetCbrtif14 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac14ifcrType) SetCbrtif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14ifcrFieldCbrtif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14ifcrFieldCbrtif14Mask)
	}
}

const (
	RegisterMdmac14ifcrFieldCbtif14Shift = 3
	RegisterMdmac14ifcrFieldCbtif14Mask  = 0x8
)

// GetCbtif14 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac14ifcrType) GetCbtif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14ifcrFieldCbtif14Mask) != 0
}

// SetCbtif14 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac14ifcrType) SetCbtif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14ifcrFieldCbtif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14ifcrFieldCbtif14Mask)
	}
}

const (
	RegisterMdmac14ifcrFieldCltcif14Shift = 4
	RegisterMdmac14ifcrFieldCltcif14Mask  = 0x10
)

// GetCltcif14 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac14ifcrType) GetCltcif14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14ifcrFieldCltcif14Mask) != 0
}

// SetCltcif14 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac14ifcrType) SetCltcif14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14ifcrFieldCltcif14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14ifcrFieldCltcif14Mask)
	}
}

// RegisterMdmac14esrType MDMA Channel x error status register
type RegisterMdmac14esrType uint32

func (r *RegisterMdmac14esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14esrFieldTeaShift = 0
	RegisterMdmac14esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac14esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14esrFieldTeaMask) >> RegisterMdmac14esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac14esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14esrFieldTeaMask)|(uint32(value)<<RegisterMdmac14esrFieldTeaShift))
}

const (
	RegisterMdmac14esrFieldTedShift = 7
	RegisterMdmac14esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac14esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac14esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14esrFieldTedMask)
	}
}

const (
	RegisterMdmac14esrFieldTeldShift = 8
	RegisterMdmac14esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac14esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac14esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14esrFieldTeldMask)
	}
}

const (
	RegisterMdmac14esrFieldTemdShift = 9
	RegisterMdmac14esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac14esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac14esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14esrFieldTemdMask)
	}
}

const (
	RegisterMdmac14esrFieldAseShift = 10
	RegisterMdmac14esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac14esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac14esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14esrFieldAseMask)
	}
}

const (
	RegisterMdmac14esrFieldBseShift = 11
	RegisterMdmac14esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac14esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac14esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14esrFieldBseMask)
	}
}

// RegisterMdmac14crType This register is used to control the concerned channel.
type RegisterMdmac14crType uint32

func (r *RegisterMdmac14crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14crFieldEnShift = 0
	RegisterMdmac14crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac14crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac14crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14crFieldEnMask)
	}
}

const (
	RegisterMdmac14crFieldTeieShift = 1
	RegisterMdmac14crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac14crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac14crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14crFieldTeieMask)
	}
}

const (
	RegisterMdmac14crFieldCtcieShift = 2
	RegisterMdmac14crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac14crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac14crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14crFieldCtcieMask)
	}
}

const (
	RegisterMdmac14crFieldBrtieShift = 3
	RegisterMdmac14crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac14crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac14crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14crFieldBrtieMask)
	}
}

const (
	RegisterMdmac14crFieldBtieShift = 4
	RegisterMdmac14crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac14crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac14crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14crFieldBtieMask)
	}
}

const (
	RegisterMdmac14crFieldTcieShift = 5
	RegisterMdmac14crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac14crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac14crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14crFieldTcieMask)
	}
}

const (
	RegisterMdmac14crFieldPlShift = 6
	RegisterMdmac14crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac14crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14crFieldPlMask) >> RegisterMdmac14crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac14crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14crFieldPlMask)|(uint32(value)<<RegisterMdmac14crFieldPlShift))
}

const (
	RegisterMdmac14crFieldBexShift = 12
	RegisterMdmac14crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac14crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac14crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14crFieldBexMask)
	}
}

const (
	RegisterMdmac14crFieldHexShift = 13
	RegisterMdmac14crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac14crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac14crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14crFieldHexMask)
	}
}

const (
	RegisterMdmac14crFieldWexShift = 14
	RegisterMdmac14crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac14crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac14crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14crFieldWexMask)
	}
}

const (
	RegisterMdmac14crFieldSwrqShift = 16
	RegisterMdmac14crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac14crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14crFieldSwrqMask)
	}
}

// RegisterMdmac14tcrType This register is used to configure the concerned channel.
type RegisterMdmac14tcrType uint32

func (r *RegisterMdmac14tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14tcrFieldSincShift = 0
	RegisterMdmac14tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac14tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldSincMask) >> RegisterMdmac14tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac14tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldSincMask)|(uint32(value)<<RegisterMdmac14tcrFieldSincShift))
}

const (
	RegisterMdmac14tcrFieldDincShift = 2
	RegisterMdmac14tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac14tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldDincMask) >> RegisterMdmac14tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac14tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldDincMask)|(uint32(value)<<RegisterMdmac14tcrFieldDincShift))
}

const (
	RegisterMdmac14tcrFieldSsizeShift = 4
	RegisterMdmac14tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac14tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldSsizeMask) >> RegisterMdmac14tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac14tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac14tcrFieldSsizeShift))
}

const (
	RegisterMdmac14tcrFieldDsizeShift = 6
	RegisterMdmac14tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac14tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldDsizeMask) >> RegisterMdmac14tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac14tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac14tcrFieldDsizeShift))
}

const (
	RegisterMdmac14tcrFieldSincosShift = 8
	RegisterMdmac14tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac14tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldSincosMask) >> RegisterMdmac14tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac14tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac14tcrFieldSincosShift))
}

const (
	RegisterMdmac14tcrFieldDincosShift = 10
	RegisterMdmac14tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac14tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldDincosMask) >> RegisterMdmac14tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac14tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac14tcrFieldDincosShift))
}

const (
	RegisterMdmac14tcrFieldSburstShift = 12
	RegisterMdmac14tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac14tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldSburstMask) >> RegisterMdmac14tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac14tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac14tcrFieldSburstShift))
}

const (
	RegisterMdmac14tcrFieldDburstShift = 15
	RegisterMdmac14tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac14tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldDburstMask) >> RegisterMdmac14tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac14tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac14tcrFieldDburstShift))
}

const (
	RegisterMdmac14tcrFieldTlenShift = 18
	RegisterMdmac14tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac14tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldTlenMask) >> RegisterMdmac14tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac14tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac14tcrFieldTlenShift))
}

const (
	RegisterMdmac14tcrFieldPkeShift = 25
	RegisterMdmac14tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac14tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac14tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac14tcrFieldPamShift = 26
	RegisterMdmac14tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac14tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldPamMask) >> RegisterMdmac14tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac14tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldPamMask)|(uint32(value)<<RegisterMdmac14tcrFieldPamShift))
}

const (
	RegisterMdmac14tcrFieldTrgmShift = 28
	RegisterMdmac14tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac14tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldTrgmMask) >> RegisterMdmac14tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac14tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac14tcrFieldTrgmShift))
}

const (
	RegisterMdmac14tcrFieldSwrmShift = 30
	RegisterMdmac14tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac14tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac14tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac14tcrFieldBwmShift = 31
	RegisterMdmac14tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac14tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac14tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tcrFieldBwmMask)
	}
}

// RegisterMdmac14bndtrType MDMA Channel x block number of data register
type RegisterMdmac14bndtrType uint32

func (r *RegisterMdmac14bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14bndtrFieldBndtShift = 0
	RegisterMdmac14bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac14bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14bndtrFieldBndtMask) >> RegisterMdmac14bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac14bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac14bndtrFieldBndtShift))
}

const (
	RegisterMdmac14bndtrFieldBrsumShift = 18
	RegisterMdmac14bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac14bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac14bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac14bndtrFieldBrdumShift = 19
	RegisterMdmac14bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac14bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac14bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac14bndtrFieldBrcShift = 20
	RegisterMdmac14bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac14bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14bndtrFieldBrcMask) >> RegisterMdmac14bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac14bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac14bndtrFieldBrcShift))
}

// RegisterMdmac14sarType MDMA channel x source address register
type RegisterMdmac14sarType uint32

func (r *RegisterMdmac14sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14sarFieldSarShift = 0
	RegisterMdmac14sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac14sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14sarFieldSarMask) >> RegisterMdmac14sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac14sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14sarFieldSarMask)|(uint32(value)<<RegisterMdmac14sarFieldSarShift))
}

// RegisterMdmac14darType MDMA channel x destination address register
type RegisterMdmac14darType uint32

func (r *RegisterMdmac14darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14darFieldDarShift = 0
	RegisterMdmac14darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac14darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14darFieldDarMask) >> RegisterMdmac14darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac14darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14darFieldDarMask)|(uint32(value)<<RegisterMdmac14darFieldDarShift))
}

// RegisterMdmac14brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac14brurType uint32

func (r *RegisterMdmac14brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14brurFieldSuvShift = 0
	RegisterMdmac14brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac14brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14brurFieldSuvMask) >> RegisterMdmac14brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac14brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14brurFieldSuvMask)|(uint32(value)<<RegisterMdmac14brurFieldSuvShift))
}

const (
	RegisterMdmac14brurFieldDuvShift = 16
	RegisterMdmac14brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac14brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14brurFieldDuvMask) >> RegisterMdmac14brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac14brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14brurFieldDuvMask)|(uint32(value)<<RegisterMdmac14brurFieldDuvShift))
}

// RegisterMdmac14larType MDMA channel x Link Address register
type RegisterMdmac14larType uint32

func (r *RegisterMdmac14larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14larFieldLarShift = 0
	RegisterMdmac14larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac14larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14larFieldLarMask) >> RegisterMdmac14larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac14larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14larFieldLarMask)|(uint32(value)<<RegisterMdmac14larFieldLarShift))
}

// RegisterMdmac14tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac14tbrType uint32

func (r *RegisterMdmac14tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14tbrFieldTselShift = 0
	RegisterMdmac14tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac14tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tbrFieldTselMask) >> RegisterMdmac14tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac14tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tbrFieldTselMask)|(uint32(value)<<RegisterMdmac14tbrFieldTselShift))
}

const (
	RegisterMdmac14tbrFieldSbusShift = 16
	RegisterMdmac14tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac14tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac14tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac14tbrFieldDbusShift = 17
	RegisterMdmac14tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac14tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac14tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac14tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14tbrFieldDbusMask)
	}
}

// RegisterMdmac14marType MDMA channel x Mask address register
type RegisterMdmac14marType uint32

func (r *RegisterMdmac14marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14marFieldMarShift = 0
	RegisterMdmac14marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac14marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14marFieldMarMask) >> RegisterMdmac14marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac14marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14marFieldMarMask)|(uint32(value)<<RegisterMdmac14marFieldMarShift))
}

// RegisterMdmac14mdrType MDMA channel x Mask Data register
type RegisterMdmac14mdrType uint32

func (r *RegisterMdmac14mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac14mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac14mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac14mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac14mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac14mdrFieldMdrShift = 0
	RegisterMdmac14mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac14mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac14mdrFieldMdrMask) >> RegisterMdmac14mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac14mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac14mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac14mdrFieldMdrShift))
}

// RegisterMdmac15isrType MDMA channel x interrupt/status register
type RegisterMdmac15isrType uint32

func (r *RegisterMdmac15isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15isrFieldTeif15Shift = 0
	RegisterMdmac15isrFieldTeif15Mask  = 0x1
)

// GetTeif15 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac15isrType) GetTeif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15isrFieldTeif15Mask) != 0
}

// SetTeif15 Channel x transfer error interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac15isrType) SetTeif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15isrFieldTeif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15isrFieldTeif15Mask)
	}
}

const (
	RegisterMdmac15isrFieldCtcif15Shift = 1
	RegisterMdmac15isrFieldCtcif15Mask  = 0x2
)

// GetCtcif15 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac15isrType) GetCtcif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15isrFieldCtcif15Mask) != 0
}

// SetCtcif15 Channel x Channel Transfer Complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register. CTC is set when the last block was transferred and the channel has been automatically disabled. CTC is also set when the channel is suspended, as a result of writing EN bit to 0.
func (r *RegisterMdmac15isrType) SetCtcif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15isrFieldCtcif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15isrFieldCtcif15Mask)
	}
}

const (
	RegisterMdmac15isrFieldBrtif15Shift = 2
	RegisterMdmac15isrFieldBrtif15Mask  = 0x4
)

// GetBrtif15 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac15isrType) GetBrtif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15isrFieldBrtif15Mask) != 0
}

// SetBrtif15 Channel x block repeat transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac15isrType) SetBrtif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15isrFieldBrtif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15isrFieldBrtif15Mask)
	}
}

const (
	RegisterMdmac15isrFieldBtif15Shift = 3
	RegisterMdmac15isrFieldBtif15Mask  = 0x8
)

// GetBtif15 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac15isrType) GetBtif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15isrFieldBtif15Mask) != 0
}

// SetBtif15 Channel x block transfer complete interrupt flag This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCRy register.
func (r *RegisterMdmac15isrType) SetBtif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15isrFieldBtif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15isrFieldBtif15Mask)
	}
}

const (
	RegisterMdmac15isrFieldTcif15Shift = 4
	RegisterMdmac15isrFieldTcif15Mask  = 0x10
)

// GetTcif15 channel x buffer transfer complete
func (r *RegisterMdmac15isrType) GetTcif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15isrFieldTcif15Mask) != 0
}

// SetTcif15 channel x buffer transfer complete
func (r *RegisterMdmac15isrType) SetTcif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15isrFieldTcif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15isrFieldTcif15Mask)
	}
}

const (
	RegisterMdmac15isrFieldCrqa15Shift = 16
	RegisterMdmac15isrFieldCrqa15Mask  = 0x10000
)

// GetCrqa15 channel x request active flag
func (r *RegisterMdmac15isrType) GetCrqa15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15isrFieldCrqa15Mask) != 0
}

// SetCrqa15 channel x request active flag
func (r *RegisterMdmac15isrType) SetCrqa15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15isrFieldCrqa15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15isrFieldCrqa15Mask)
	}
}

// RegisterMdmac15ifcrType MDMA channel x interrupt flag clear register
type RegisterMdmac15ifcrType uint32

func (r *RegisterMdmac15ifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15ifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15ifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15ifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15ifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15ifcrFieldCteif15Shift = 0
	RegisterMdmac15ifcrFieldCteif15Mask  = 0x1
)

// GetCteif15 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac15ifcrType) GetCteif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15ifcrFieldCteif15Mask) != 0
}

// SetCteif15 Channel x clear transfer error interrupt flag Writing a 1 into this bit clears TEIFx in the MDMA_ISRy register
func (r *RegisterMdmac15ifcrType) SetCteif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15ifcrFieldCteif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15ifcrFieldCteif15Mask)
	}
}

const (
	RegisterMdmac15ifcrFieldCctcif15Shift = 1
	RegisterMdmac15ifcrFieldCctcif15Mask  = 0x2
)

// GetCctcif15 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac15ifcrType) GetCctcif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15ifcrFieldCctcif15Mask) != 0
}

// SetCctcif15 Clear Channel transfer complete interrupt flag for channel x Writing a 1 into this bit clears CTCIFx in the MDMA_ISRy register
func (r *RegisterMdmac15ifcrType) SetCctcif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15ifcrFieldCctcif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15ifcrFieldCctcif15Mask)
	}
}

const (
	RegisterMdmac15ifcrFieldCbrtif15Shift = 2
	RegisterMdmac15ifcrFieldCbrtif15Mask  = 0x4
)

// GetCbrtif15 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac15ifcrType) GetCbrtif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15ifcrFieldCbrtif15Mask) != 0
}

// SetCbrtif15 Channel x clear block repeat transfer complete interrupt flag Writing a 1 into this bit clears BRTIFx in the MDMA_ISRy register
func (r *RegisterMdmac15ifcrType) SetCbrtif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15ifcrFieldCbrtif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15ifcrFieldCbrtif15Mask)
	}
}

const (
	RegisterMdmac15ifcrFieldCbtif15Shift = 3
	RegisterMdmac15ifcrFieldCbtif15Mask  = 0x8
)

// GetCbtif15 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac15ifcrType) GetCbtif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15ifcrFieldCbtif15Mask) != 0
}

// SetCbtif15 Channel x Clear block transfer complete interrupt flag Writing a 1 into this bit clears BTIFx in the MDMA_ISRy register
func (r *RegisterMdmac15ifcrType) SetCbtif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15ifcrFieldCbtif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15ifcrFieldCbtif15Mask)
	}
}

const (
	RegisterMdmac15ifcrFieldCltcif15Shift = 4
	RegisterMdmac15ifcrFieldCltcif15Mask  = 0x10
)

// GetCltcif15 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac15ifcrType) GetCltcif15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15ifcrFieldCltcif15Mask) != 0
}

// SetCltcif15 CLear buffer Transfer Complete Interrupt Flag for channel x Writing a 1 into this bit clears TCIFx in the MDMA_ISRy register
func (r *RegisterMdmac15ifcrType) SetCltcif15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15ifcrFieldCltcif15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15ifcrFieldCltcif15Mask)
	}
}

// RegisterMdmac15esrType MDMA Channel x error status register
type RegisterMdmac15esrType uint32

func (r *RegisterMdmac15esrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15esrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15esrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15esrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15esrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15esrFieldTeaShift = 0
	RegisterMdmac15esrFieldTeaMask  = 0x7f
)

// GetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac15esrType) GetTea() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15esrFieldTeaMask) >> RegisterMdmac15esrFieldTeaShift)
}

// SetTea Transfer Error Address These bits are set and cleared by HW, in case of an MDMA data transfer error. It is used in conjunction with TED. This field indicates the 7 LSBits of the address which generated a transfer/access error. It may be used by SW to retrieve the failing address, by adding this value (truncated to the buffer transfer length size) to the current SAR/DAR value. Note: The SAR/DAR current value doesnt reflect this last address due to the FIFO management system. The SAR/DAR are only updated at the end of a (buffer) transfer (of TLEN+1 bytes). Note: It is not set in case of a link data error.
func (r *RegisterMdmac15esrType) SetTea(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15esrFieldTeaMask)|(uint32(value)<<RegisterMdmac15esrFieldTeaShift))
}

const (
	RegisterMdmac15esrFieldTedShift = 7
	RegisterMdmac15esrFieldTedMask  = 0x80
)

// GetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac15esrType) GetTed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15esrFieldTedMask) != 0
}

// SetTed Transfer Error Direction These bit is set and cleared by HW, in case of an MDMA data transfer error.
func (r *RegisterMdmac15esrType) SetTed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15esrFieldTedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15esrFieldTedMask)
	}
}

const (
	RegisterMdmac15esrFieldTeldShift = 8
	RegisterMdmac15esrFieldTeldMask  = 0x100
)

// GetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac15esrType) GetTeld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15esrFieldTeldMask) != 0
}

// SetTeld Transfer Error Link Data These bit is set by HW, in case of a transfer error while reading the block link data structure. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac15esrType) SetTeld(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15esrFieldTeldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15esrFieldTeldMask)
	}
}

const (
	RegisterMdmac15esrFieldTemdShift = 9
	RegisterMdmac15esrFieldTemdMask  = 0x200
)

// GetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac15esrType) GetTemd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15esrFieldTemdMask) != 0
}

// SetTemd Transfer Error Mask Data These bit is set by HW, in case of a transfer error while writing the Mask Data. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac15esrType) SetTemd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15esrFieldTemdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15esrFieldTemdMask)
	}
}

const (
	RegisterMdmac15esrFieldAseShift = 10
	RegisterMdmac15esrFieldAseMask  = 0x400
)

// GetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac15esrType) GetAse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15esrFieldAseMask) != 0
}

// SetAse Address/Size Error These bit is set by HW, when the programmed address is not aligned with the data size. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac15esrType) SetAse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15esrFieldAseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15esrFieldAseMask)
	}
}

const (
	RegisterMdmac15esrFieldBseShift = 11
	RegisterMdmac15esrFieldBseMask  = 0x800
)

// GetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac15esrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15esrFieldBseMask) != 0
}

// SetBse Block Size Error These bit is set by HW, when the block size is not an integer multiple of the data size either for source or destination. TED will indicate whether the problem is on the source or destination. It is cleared by software writing 1 to the CTEIFx bit in the DMA_IFCRy register.
func (r *RegisterMdmac15esrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15esrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15esrFieldBseMask)
	}
}

// RegisterMdmac15crType This register is used to control the concerned channel.
type RegisterMdmac15crType uint32

func (r *RegisterMdmac15crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15crFieldEnShift = 0
	RegisterMdmac15crFieldEnMask  = 0x1
)

// GetEn channel enable
func (r *RegisterMdmac15crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15crFieldEnMask) != 0
}

// SetEn channel enable
func (r *RegisterMdmac15crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15crFieldEnMask)
	}
}

const (
	RegisterMdmac15crFieldTeieShift = 1
	RegisterMdmac15crFieldTeieMask  = 0x2
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac15crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac15crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15crFieldTeieMask)
	}
}

const (
	RegisterMdmac15crFieldCtcieShift = 2
	RegisterMdmac15crFieldCtcieMask  = 0x4
)

// GetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac15crType) GetCtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15crFieldCtcieMask) != 0
}

// SetCtcie Channel Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac15crType) SetCtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15crFieldCtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15crFieldCtcieMask)
	}
}

const (
	RegisterMdmac15crFieldBrtieShift = 3
	RegisterMdmac15crFieldBrtieMask  = 0x8
)

// GetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac15crType) GetBrtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15crFieldBrtieMask) != 0
}

// SetBrtie Block Repeat transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac15crType) SetBrtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15crFieldBrtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15crFieldBrtieMask)
	}
}

const (
	RegisterMdmac15crFieldBtieShift = 4
	RegisterMdmac15crFieldBtieMask  = 0x10
)

// GetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac15crType) GetBtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15crFieldBtieMask) != 0
}

// SetBtie Block Transfer interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac15crType) SetBtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15crFieldBtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15crFieldBtieMask)
	}
}

const (
	RegisterMdmac15crFieldTcieShift = 5
	RegisterMdmac15crFieldTcieMask  = 0x20
)

// GetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac15crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15crFieldTcieMask) != 0
}

// SetTcie buffer Transfer Complete interrupt enable This bit is set and cleared by software.
func (r *RegisterMdmac15crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15crFieldTcieMask)
	}
}

const (
	RegisterMdmac15crFieldPlShift = 6
	RegisterMdmac15crFieldPlMask  = 0xc0
)

// GetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac15crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15crFieldPlMask) >> RegisterMdmac15crFieldPlShift)
}

// SetPl Priority level These bits are set and cleared by software. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac15crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15crFieldPlMask)|(uint32(value)<<RegisterMdmac15crFieldPlShift))
}

const (
	RegisterMdmac15crFieldBexShift = 12
	RegisterMdmac15crFieldBexMask  = 0x1000
)

// GetBex byte Endianness exchange
func (r *RegisterMdmac15crType) GetBex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15crFieldBexMask) != 0
}

// SetBex byte Endianness exchange
func (r *RegisterMdmac15crType) SetBex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15crFieldBexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15crFieldBexMask)
	}
}

const (
	RegisterMdmac15crFieldHexShift = 13
	RegisterMdmac15crFieldHexMask  = 0x2000
)

// GetHex Half word Endianes exchange
func (r *RegisterMdmac15crType) GetHex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15crFieldHexMask) != 0
}

// SetHex Half word Endianes exchange
func (r *RegisterMdmac15crType) SetHex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15crFieldHexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15crFieldHexMask)
	}
}

const (
	RegisterMdmac15crFieldWexShift = 14
	RegisterMdmac15crFieldWexMask  = 0x4000
)

// GetWex Word Endianness exchange
func (r *RegisterMdmac15crType) GetWex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15crFieldWexMask) != 0
}

// SetWex Word Endianness exchange
func (r *RegisterMdmac15crType) SetWex(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15crFieldWexMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15crFieldWexMask)
	}
}

const (
	RegisterMdmac15crFieldSwrqShift = 16
	RegisterMdmac15crFieldSwrqMask  = 0x10000
)

// SetSwrq SW ReQuest Writing a 1 into this bit sets the CRQAx in MDMA_ISRy register, activating the request on Channel x Note: Either the whole CxCR register or the 8-bit/16-bit register @ Address offset: 0x4E + 0x40 chn may be used for SWRQ activation. In case of a SW request, acknowledge is not generated (neither HW signal, nor CxMAR write access).
func (r *RegisterMdmac15crType) SetSwrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15crFieldSwrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15crFieldSwrqMask)
	}
}

// RegisterMdmac15tcrType This register is used to configure the concerned channel.
type RegisterMdmac15tcrType uint32

func (r *RegisterMdmac15tcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15tcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15tcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15tcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15tcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15tcrFieldSincShift = 0
	RegisterMdmac15tcrFieldSincMask  = 0x3
)

// GetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac15tcrType) GetSinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldSincMask) >> RegisterMdmac15tcrFieldSincShift)
}

// SetSinc Source increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When source is AHB (SBUS=1), SINC = 00 is forbidden. In Linked List Mode, at the end of a block (single or last block in repeated block transfer mode), this register will be loaded from memory (from address given by current LAR[31:0] + 0x00).
func (r *RegisterMdmac15tcrType) SetSinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldSincMask)|(uint32(value)<<RegisterMdmac15tcrFieldSincShift))
}

const (
	RegisterMdmac15tcrFieldDincShift = 2
	RegisterMdmac15tcrFieldDincMask  = 0xc
)

// GetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac15tcrType) GetDinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldDincMask) >> RegisterMdmac15tcrFieldDincShift)
}

// SetDinc Destination increment mode These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: When destination is AHB (DBUS=1), DINC = 00 is forbidden.
func (r *RegisterMdmac15tcrType) SetDinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldDincMask)|(uint32(value)<<RegisterMdmac15tcrFieldDincShift))
}

const (
	RegisterMdmac15tcrFieldSsizeShift = 4
	RegisterMdmac15tcrFieldSsizeMask  = 0x30
)

// GetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac15tcrType) GetSsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldSsizeMask) >> RegisterMdmac15tcrFieldSsizeShift)
}

// SetSsize Source data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0 Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If SINCOS &lt; SSIZE and SINC &#8800; 00, the result will be unpredictable. Note: SSIZE = 11 (double-word) is forbidden when source is TCM/AHB bus (SBUS=1).
func (r *RegisterMdmac15tcrType) SetSsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldSsizeMask)|(uint32(value)<<RegisterMdmac15tcrFieldSsizeShift))
}

const (
	RegisterMdmac15tcrFieldDsizeShift = 6
	RegisterMdmac15tcrFieldDsizeMask  = 0xc0
)

// GetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac15tcrType) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldDsizeMask) >> RegisterMdmac15tcrFieldDsizeShift)
}

// SetDsize Destination data size These bits are set and cleared by software. These bits are protected and can be written only if EN is 0. Note: If a value of 11 is programmed for the TCM access/AHB port, a transfer error will occur (TEIF bit set) If DINCOS &lt; DSIZE and DINC &#8800; 00, the result will be unpredictable. Note: DSIZE = 11 (double-word) is forbidden when destination is TCM/AHB bus (DBUS=1).
func (r *RegisterMdmac15tcrType) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldDsizeMask)|(uint32(value)<<RegisterMdmac15tcrFieldDsizeShift))
}

const (
	RegisterMdmac15tcrFieldSincosShift = 8
	RegisterMdmac15tcrFieldSincosMask  = 0x300
)

// GetSincos source increment offset size
func (r *RegisterMdmac15tcrType) GetSincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldSincosMask) >> RegisterMdmac15tcrFieldSincosShift)
}

// SetSincos source increment offset size
func (r *RegisterMdmac15tcrType) SetSincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldSincosMask)|(uint32(value)<<RegisterMdmac15tcrFieldSincosShift))
}

const (
	RegisterMdmac15tcrFieldDincosShift = 10
	RegisterMdmac15tcrFieldDincosMask  = 0xc00
)

// GetDincos Destination increment offset
func (r *RegisterMdmac15tcrType) GetDincos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldDincosMask) >> RegisterMdmac15tcrFieldDincosShift)
}

// SetDincos Destination increment offset
func (r *RegisterMdmac15tcrType) SetDincos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldDincosMask)|(uint32(value)<<RegisterMdmac15tcrFieldDincosShift))
}

const (
	RegisterMdmac15tcrFieldSburstShift = 12
	RegisterMdmac15tcrFieldSburstMask  = 0x7000
)

// GetSburst source burst transfer configuration
func (r *RegisterMdmac15tcrType) GetSburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldSburstMask) >> RegisterMdmac15tcrFieldSburstShift)
}

// SetSburst source burst transfer configuration
func (r *RegisterMdmac15tcrType) SetSburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldSburstMask)|(uint32(value)<<RegisterMdmac15tcrFieldSburstShift))
}

const (
	RegisterMdmac15tcrFieldDburstShift = 15
	RegisterMdmac15tcrFieldDburstMask  = 0x38000
)

// GetDburst Destination burst transfer configuration
func (r *RegisterMdmac15tcrType) GetDburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldDburstMask) >> RegisterMdmac15tcrFieldDburstShift)
}

// SetDburst Destination burst transfer configuration
func (r *RegisterMdmac15tcrType) SetDburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldDburstMask)|(uint32(value)<<RegisterMdmac15tcrFieldDburstShift))
}

const (
	RegisterMdmac15tcrFieldTlenShift = 18
	RegisterMdmac15tcrFieldTlenMask  = 0x1fc0000
)

// GetTlen buffer transfer lengh
func (r *RegisterMdmac15tcrType) GetTlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldTlenMask) >> RegisterMdmac15tcrFieldTlenShift)
}

// SetTlen buffer transfer lengh
func (r *RegisterMdmac15tcrType) SetTlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldTlenMask)|(uint32(value)<<RegisterMdmac15tcrFieldTlenShift))
}

const (
	RegisterMdmac15tcrFieldPkeShift = 25
	RegisterMdmac15tcrFieldPkeMask  = 0x2000000
)

// GetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac15tcrType) GetPke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldPkeMask) != 0
}

// SetPke PacK Enable These bit is set and cleared by software. If the Source Size is smaller than the destination, it will be padded according to the PAM value. If the Source data size is larger than the destination one, it will be truncated. The alignment will be done according to the PAM[0] value. This bit is protected and can be written only if EN is 0
func (r *RegisterMdmac15tcrType) SetPke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15tcrFieldPkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldPkeMask)
	}
}

const (
	RegisterMdmac15tcrFieldPamShift = 26
	RegisterMdmac15tcrFieldPamMask  = 0xc000000
)

// GetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac15tcrType) GetPam() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldPamMask) >> RegisterMdmac15tcrFieldPamShift)
}

// SetPam Padding/Alignement Mode These bits are set and cleared by software. Case 1: Source data size smaller than destination data size - 3 options are valid. Case 2: Source data size larger than destination data size. The remainder part is discarded. When PKE = 1 or DSIZE=SSIZE, these bits are ignored. These bits are protected and can be written only if EN is 0
func (r *RegisterMdmac15tcrType) SetPam(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldPamMask)|(uint32(value)<<RegisterMdmac15tcrFieldPamShift))
}

const (
	RegisterMdmac15tcrFieldTrgmShift = 28
	RegisterMdmac15tcrFieldTrgmMask  = 0x30000000
)

// GetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac15tcrType) GetTrgm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldTrgmMask) >> RegisterMdmac15tcrFieldTrgmShift)
}

// SetTrgm Trigger Mode These bits are set and cleared by software. Note: If TRGM is 11 for the current block, all the values loaded at the end of the current block through the linked list mechanism must keep the same value (TRGM=11) and the same SWRM value, otherwise the result is undefined. These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac15tcrType) SetTrgm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldTrgmMask)|(uint32(value)<<RegisterMdmac15tcrFieldTrgmShift))
}

const (
	RegisterMdmac15tcrFieldSwrmShift = 30
	RegisterMdmac15tcrFieldSwrmMask  = 0x40000000
)

// GetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac15tcrType) GetSwrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldSwrmMask) != 0
}

// SetSwrm SW Request Mode This bit is set and cleared by software. If a HW or SW request is currently active, the bit change will be delayed until the current transfer is completed. If the CxMAR contains a valid address, the CxMDR value will also be written @ CxMAR address. This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac15tcrType) SetSwrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15tcrFieldSwrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldSwrmMask)
	}
}

const (
	RegisterMdmac15tcrFieldBwmShift = 31
	RegisterMdmac15tcrFieldBwmMask  = 0x80000000
)

// GetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac15tcrType) GetBwm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tcrFieldBwmMask) != 0
}

// SetBwm Bufferable Write Mode This bit is set and cleared by software. This bit is protected and can be written only if EN is 0. Note: All MDMA destination accesses are non-cacheable.
func (r *RegisterMdmac15tcrType) SetBwm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15tcrFieldBwmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tcrFieldBwmMask)
	}
}

// RegisterMdmac15bndtrType MDMA Channel x block number of data register
type RegisterMdmac15bndtrType uint32

func (r *RegisterMdmac15bndtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15bndtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15bndtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15bndtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15bndtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15bndtrFieldBndtShift = 0
	RegisterMdmac15bndtrFieldBndtMask  = 0x1ffff
)

// GetBndt block number of data to transfer
func (r *RegisterMdmac15bndtrType) GetBndt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15bndtrFieldBndtMask) >> RegisterMdmac15bndtrFieldBndtShift)
}

// SetBndt block number of data to transfer
func (r *RegisterMdmac15bndtrType) SetBndt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15bndtrFieldBndtMask)|(uint32(value)<<RegisterMdmac15bndtrFieldBndtShift))
}

const (
	RegisterMdmac15bndtrFieldBrsumShift = 18
	RegisterMdmac15bndtrFieldBrsumMask  = 0x40000
)

// GetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac15bndtrType) GetBrsum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15bndtrFieldBrsumMask) != 0
}

// SetBrsum Block Repeat Source address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac15bndtrType) SetBrsum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15bndtrFieldBrsumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15bndtrFieldBrsumMask)
	}
}

const (
	RegisterMdmac15bndtrFieldBrdumShift = 19
	RegisterMdmac15bndtrFieldBrdumMask  = 0x80000
)

// GetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac15bndtrType) GetBrdum() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15bndtrFieldBrdumMask) != 0
}

// SetBrdum Block Repeat Destination address Update Mode These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac15bndtrType) SetBrdum(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15bndtrFieldBrdumMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15bndtrFieldBrdumMask)
	}
}

const (
	RegisterMdmac15bndtrFieldBrcShift = 20
	RegisterMdmac15bndtrFieldBrcMask  = 0xfff00000
)

// GetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac15bndtrType) GetBrc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15bndtrFieldBrcMask) >> RegisterMdmac15bndtrFieldBrcShift)
}

// SetBrc Block Repeat Count This field contains the number of repetitions of the current block (0 to 4095). When the channel is enabled, this register is read-only, indicating the remaining number of blocks, excluding the current one. This register decrements after each complete block transfer. Once the last block transfer has completed, this register can either stay at zero or be reloaded automatically from memory (in Linked List mode - i.e. Link Address valid). These bits are protected and can be written only if EN is 0.
func (r *RegisterMdmac15bndtrType) SetBrc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15bndtrFieldBrcMask)|(uint32(value)<<RegisterMdmac15bndtrFieldBrcShift))
}

// RegisterMdmac15sarType MDMA channel x source address register
type RegisterMdmac15sarType uint32

func (r *RegisterMdmac15sarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15sarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15sarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15sarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15sarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15sarFieldSarShift = 0
	RegisterMdmac15sarFieldSarMask  = 0xffffffff
)

// GetSar source adr base
func (r *RegisterMdmac15sarType) GetSar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15sarFieldSarMask) >> RegisterMdmac15sarFieldSarShift)
}

// SetSar source adr base
func (r *RegisterMdmac15sarType) SetSar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15sarFieldSarMask)|(uint32(value)<<RegisterMdmac15sarFieldSarShift))
}

// RegisterMdmac15darType MDMA channel x destination address register
type RegisterMdmac15darType uint32

func (r *RegisterMdmac15darType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15darType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15darType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15darType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15darType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15darFieldDarShift = 0
	RegisterMdmac15darFieldDarMask  = 0xffffffff
)

// GetDar Destination adr base
func (r *RegisterMdmac15darType) GetDar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15darFieldDarMask) >> RegisterMdmac15darFieldDarShift)
}

// SetDar Destination adr base
func (r *RegisterMdmac15darType) SetDar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15darFieldDarMask)|(uint32(value)<<RegisterMdmac15darFieldDarShift))
}

// RegisterMdmac15brurType MDMA channel x Block Repeat address Update register
type RegisterMdmac15brurType uint32

func (r *RegisterMdmac15brurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15brurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15brurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15brurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15brurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15brurFieldSuvShift = 0
	RegisterMdmac15brurFieldSuvMask  = 0xffff
)

// GetSuv source adresse update value
func (r *RegisterMdmac15brurType) GetSuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15brurFieldSuvMask) >> RegisterMdmac15brurFieldSuvShift)
}

// SetSuv source adresse update value
func (r *RegisterMdmac15brurType) SetSuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15brurFieldSuvMask)|(uint32(value)<<RegisterMdmac15brurFieldSuvShift))
}

const (
	RegisterMdmac15brurFieldDuvShift = 16
	RegisterMdmac15brurFieldDuvMask  = 0xffff0000
)

// GetDuv destination address update
func (r *RegisterMdmac15brurType) GetDuv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15brurFieldDuvMask) >> RegisterMdmac15brurFieldDuvShift)
}

// SetDuv destination address update
func (r *RegisterMdmac15brurType) SetDuv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15brurFieldDuvMask)|(uint32(value)<<RegisterMdmac15brurFieldDuvShift))
}

// RegisterMdmac15larType MDMA channel x Link Address register
type RegisterMdmac15larType uint32

func (r *RegisterMdmac15larType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15larType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15larType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15larType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15larType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15larFieldLarShift = 0
	RegisterMdmac15larFieldLarMask  = 0xffffffff
)

// GetLar Link address register
func (r *RegisterMdmac15larType) GetLar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15larFieldLarMask) >> RegisterMdmac15larFieldLarShift)
}

// SetLar Link address register
func (r *RegisterMdmac15larType) SetLar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15larFieldLarMask)|(uint32(value)<<RegisterMdmac15larFieldLarShift))
}

// RegisterMdmac15tbrType MDMA channel x Trigger and Bus selection Register
type RegisterMdmac15tbrType uint32

func (r *RegisterMdmac15tbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15tbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15tbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15tbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15tbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15tbrFieldTselShift = 0
	RegisterMdmac15tbrFieldTselMask  = 0x3f
)

// GetTsel Trigger selection
func (r *RegisterMdmac15tbrType) GetTsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tbrFieldTselMask) >> RegisterMdmac15tbrFieldTselShift)
}

// SetTsel Trigger selection
func (r *RegisterMdmac15tbrType) SetTsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tbrFieldTselMask)|(uint32(value)<<RegisterMdmac15tbrFieldTselShift))
}

const (
	RegisterMdmac15tbrFieldSbusShift = 16
	RegisterMdmac15tbrFieldSbusMask  = 0x10000
)

// GetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac15tbrType) GetSbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tbrFieldSbusMask) != 0
}

// SetSbus Source BUS select This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac15tbrType) SetSbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15tbrFieldSbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tbrFieldSbusMask)
	}
}

const (
	RegisterMdmac15tbrFieldDbusShift = 17
	RegisterMdmac15tbrFieldDbusMask  = 0x20000
)

// GetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac15tbrType) GetDbus() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15tbrFieldDbusMask) != 0
}

// SetDbus Destination BUS slect This bit is protected and can be written only if EN is 0.
func (r *RegisterMdmac15tbrType) SetDbus(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdmac15tbrFieldDbusMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15tbrFieldDbusMask)
	}
}

// RegisterMdmac15marType MDMA channel x Mask address register
type RegisterMdmac15marType uint32

func (r *RegisterMdmac15marType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15marType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15marType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15marType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15marType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15marFieldMarShift = 0
	RegisterMdmac15marFieldMarMask  = 0xffffffff
)

// GetMar Mask address
func (r *RegisterMdmac15marType) GetMar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15marFieldMarMask) >> RegisterMdmac15marFieldMarShift)
}

// SetMar Mask address
func (r *RegisterMdmac15marType) SetMar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15marFieldMarMask)|(uint32(value)<<RegisterMdmac15marFieldMarShift))
}

// RegisterMdmac15mdrType MDMA channel x Mask Data register
type RegisterMdmac15mdrType uint32

func (r *RegisterMdmac15mdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdmac15mdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdmac15mdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdmac15mdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdmac15mdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdmac15mdrFieldMdrShift = 0
	RegisterMdmac15mdrFieldMdrMask  = 0xffffffff
)

// GetMdr Mask data
func (r *RegisterMdmac15mdrType) GetMdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMdmac15mdrFieldMdrMask) >> RegisterMdmac15mdrFieldMdrShift)
}

// SetMdr Mask data
func (r *RegisterMdmac15mdrType) SetMdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMdmac15mdrFieldMdrMask)|(uint32(value)<<RegisterMdmac15mdrFieldMdrShift))
}
