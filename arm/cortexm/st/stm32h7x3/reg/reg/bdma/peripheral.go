package bdma

import (
	"unsafe"
	"volatile"
)

var (
	Bdma = (*_bdma)(unsafe.Pointer(uintptr(0x58025400)))
)

type _bdma struct {
	Isr    registerIsrType
	Ifcr   registerIfcrType
	Ccr1   registerCcr1Type
	Cndtr1 registerCndtr1Type
	Cpar1  registerCpar1Type
	Cmar1  registerCmar1Type
	_      [4]byte
	Ccr2   registerCcr2Type
	Cndtr2 registerCndtr2Type
	Cpar2  registerCpar2Type
	Cmar2  registerCmar2Type
	_      [4]byte
	Ccr3   registerCcr3Type
	Cndtr3 registerCndtr3Type
	Cpar3  registerCpar3Type
	Cmar3  registerCmar3Type
	_      [4]byte
	Ccr4   registerCcr4Type
	Cndtr4 registerCndtr4Type
	Cpar4  registerCpar4Type
	Cmar4  registerCmar4Type
	_      [4]byte
	Ccr5   registerCcr5Type
	Cndtr5 registerCndtr5Type
	Cpar5  registerCpar5Type
	Cmar5  registerCmar5Type
	_      [4]byte
	Ccr6   registerCcr6Type
	Cndtr6 registerCndtr6Type
	Cpar6  registerCpar6Type
	Cmar6  registerCmar6Type
	_      [4]byte
	Ccr7   registerCcr7Type
	Cndtr7 registerCndtr7Type
	Cpar7  registerCpar7Type
	Cmar7  registerCmar7Type
	_      [4]byte
	Ccr8   registerCcr8Type
	Cndtr8 registerCndtr8Type
	Cpar8  registerCpar8Type
	Cmar8  registerCmar8Type
}

// registerIsrType DMA interrupt status register
type registerIsrType uint32

const (
	RegisterIsrFieldGif1Shift = 0
	RegisterIsrFieldGif1Mask  = 0x1
)

// GetGif1 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetGif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldGif1Mask) != 0
}

// SetGif1 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetGif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldGif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldGif1Mask)
	}
}

const (
	RegisterIsrFieldTcif1Shift = 1
	RegisterIsrFieldTcif1Mask  = 0x2
)

// GetTcif1 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTcif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcif1Mask) != 0
}

// SetTcif1 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTcif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTcif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTcif1Mask)
	}
}

const (
	RegisterIsrFieldHtif1Shift = 2
	RegisterIsrFieldHtif1Mask  = 0x4
)

// GetHtif1 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetHtif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldHtif1Mask) != 0
}

// SetHtif1 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetHtif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldHtif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldHtif1Mask)
	}
}

const (
	RegisterIsrFieldTeif1Shift = 3
	RegisterIsrFieldTeif1Mask  = 0x8
)

// GetTeif1 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTeif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTeif1Mask) != 0
}

// SetTeif1 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTeif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTeif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTeif1Mask)
	}
}

const (
	RegisterIsrFieldGif2Shift = 4
	RegisterIsrFieldGif2Mask  = 0x10
)

// GetGif2 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetGif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldGif2Mask) != 0
}

// SetGif2 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetGif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldGif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldGif2Mask)
	}
}

const (
	RegisterIsrFieldTcif2Shift = 5
	RegisterIsrFieldTcif2Mask  = 0x20
)

// GetTcif2 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTcif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcif2Mask) != 0
}

// SetTcif2 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTcif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTcif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTcif2Mask)
	}
}

const (
	RegisterIsrFieldHtif2Shift = 6
	RegisterIsrFieldHtif2Mask  = 0x40
)

// GetHtif2 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetHtif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldHtif2Mask) != 0
}

// SetHtif2 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetHtif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldHtif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldHtif2Mask)
	}
}

const (
	RegisterIsrFieldTeif2Shift = 7
	RegisterIsrFieldTeif2Mask  = 0x80
)

// GetTeif2 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTeif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTeif2Mask) != 0
}

// SetTeif2 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTeif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTeif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTeif2Mask)
	}
}

const (
	RegisterIsrFieldGif3Shift = 8
	RegisterIsrFieldGif3Mask  = 0x100
)

// GetGif3 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetGif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldGif3Mask) != 0
}

// SetGif3 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetGif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldGif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldGif3Mask)
	}
}

const (
	RegisterIsrFieldTcif3Shift = 9
	RegisterIsrFieldTcif3Mask  = 0x200
)

// GetTcif3 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTcif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcif3Mask) != 0
}

// SetTcif3 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTcif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTcif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTcif3Mask)
	}
}

const (
	RegisterIsrFieldHtif3Shift = 10
	RegisterIsrFieldHtif3Mask  = 0x400
)

// GetHtif3 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetHtif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldHtif3Mask) != 0
}

// SetHtif3 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetHtif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldHtif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldHtif3Mask)
	}
}

const (
	RegisterIsrFieldTeif3Shift = 11
	RegisterIsrFieldTeif3Mask  = 0x800
)

// GetTeif3 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTeif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTeif3Mask) != 0
}

// SetTeif3 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTeif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTeif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTeif3Mask)
	}
}

const (
	RegisterIsrFieldGif4Shift = 12
	RegisterIsrFieldGif4Mask  = 0x1000
)

// GetGif4 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetGif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldGif4Mask) != 0
}

// SetGif4 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetGif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldGif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldGif4Mask)
	}
}

const (
	RegisterIsrFieldTcif4Shift = 13
	RegisterIsrFieldTcif4Mask  = 0x2000
)

// GetTcif4 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTcif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcif4Mask) != 0
}

// SetTcif4 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTcif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTcif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTcif4Mask)
	}
}

const (
	RegisterIsrFieldHtif4Shift = 14
	RegisterIsrFieldHtif4Mask  = 0x4000
)

// GetHtif4 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetHtif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldHtif4Mask) != 0
}

// SetHtif4 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetHtif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldHtif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldHtif4Mask)
	}
}

const (
	RegisterIsrFieldTeif4Shift = 15
	RegisterIsrFieldTeif4Mask  = 0x8000
)

// GetTeif4 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTeif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTeif4Mask) != 0
}

// SetTeif4 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTeif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTeif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTeif4Mask)
	}
}

const (
	RegisterIsrFieldGif5Shift = 16
	RegisterIsrFieldGif5Mask  = 0x10000
)

// GetGif5 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetGif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldGif5Mask) != 0
}

// SetGif5 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetGif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldGif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldGif5Mask)
	}
}

const (
	RegisterIsrFieldTcif5Shift = 17
	RegisterIsrFieldTcif5Mask  = 0x20000
)

// GetTcif5 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTcif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcif5Mask) != 0
}

// SetTcif5 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTcif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTcif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTcif5Mask)
	}
}

const (
	RegisterIsrFieldHtif5Shift = 18
	RegisterIsrFieldHtif5Mask  = 0x40000
)

// GetHtif5 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetHtif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldHtif5Mask) != 0
}

// SetHtif5 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetHtif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldHtif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldHtif5Mask)
	}
}

const (
	RegisterIsrFieldTeif5Shift = 19
	RegisterIsrFieldTeif5Mask  = 0x80000
)

// GetTeif5 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTeif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTeif5Mask) != 0
}

// SetTeif5 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTeif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTeif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTeif5Mask)
	}
}

const (
	RegisterIsrFieldGif6Shift = 20
	RegisterIsrFieldGif6Mask  = 0x100000
)

// GetGif6 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetGif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldGif6Mask) != 0
}

// SetGif6 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetGif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldGif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldGif6Mask)
	}
}

const (
	RegisterIsrFieldTcif6Shift = 21
	RegisterIsrFieldTcif6Mask  = 0x200000
)

// GetTcif6 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTcif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcif6Mask) != 0
}

// SetTcif6 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTcif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTcif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTcif6Mask)
	}
}

const (
	RegisterIsrFieldHtif6Shift = 22
	RegisterIsrFieldHtif6Mask  = 0x400000
)

// GetHtif6 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetHtif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldHtif6Mask) != 0
}

// SetHtif6 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetHtif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldHtif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldHtif6Mask)
	}
}

const (
	RegisterIsrFieldTeif6Shift = 23
	RegisterIsrFieldTeif6Mask  = 0x800000
)

// GetTeif6 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTeif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTeif6Mask) != 0
}

// SetTeif6 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTeif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTeif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTeif6Mask)
	}
}

const (
	RegisterIsrFieldGif7Shift = 24
	RegisterIsrFieldGif7Mask  = 0x1000000
)

// GetGif7 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetGif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldGif7Mask) != 0
}

// SetGif7 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetGif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldGif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldGif7Mask)
	}
}

const (
	RegisterIsrFieldTcif7Shift = 25
	RegisterIsrFieldTcif7Mask  = 0x2000000
)

// GetTcif7 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTcif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcif7Mask) != 0
}

// SetTcif7 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTcif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTcif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTcif7Mask)
	}
}

const (
	RegisterIsrFieldHtif7Shift = 26
	RegisterIsrFieldHtif7Mask  = 0x4000000
)

// GetHtif7 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetHtif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldHtif7Mask) != 0
}

// SetHtif7 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetHtif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldHtif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldHtif7Mask)
	}
}

const (
	RegisterIsrFieldTeif7Shift = 27
	RegisterIsrFieldTeif7Mask  = 0x8000000
)

// GetTeif7 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTeif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTeif7Mask) != 0
}

// SetTeif7 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTeif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTeif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTeif7Mask)
	}
}

const (
	RegisterIsrFieldGif8Shift = 28
	RegisterIsrFieldGif8Mask  = 0x10000000
)

// GetGif8 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetGif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldGif8Mask) != 0
}

// SetGif8 Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetGif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldGif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldGif8Mask)
	}
}

const (
	RegisterIsrFieldTcif8Shift = 29
	RegisterIsrFieldTcif8Mask  = 0x20000000
)

// GetTcif8 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTcif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcif8Mask) != 0
}

// SetTcif8 Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTcif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTcif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTcif8Mask)
	}
}

const (
	RegisterIsrFieldHtif8Shift = 30
	RegisterIsrFieldHtif8Mask  = 0x40000000
)

// GetHtif8 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetHtif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldHtif8Mask) != 0
}

// SetHtif8 Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetHtif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldHtif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldHtif8Mask)
	}
}

const (
	RegisterIsrFieldTeif8Shift = 31
	RegisterIsrFieldTeif8Mask  = 0x80000000
)

// GetTeif8 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) GetTeif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTeif8Mask) != 0
}

// SetTeif8 Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
func (r *registerIsrType) SetTeif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTeif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTeif8Mask)
	}
}

// registerIfcrType DMA interrupt flag clear register
type registerIfcrType uint32

const (
	RegisterIfcrFieldCgif1Shift = 0
	RegisterIfcrFieldCgif1Mask  = 0x1
)

// GetCgif1 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCgif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCgif1Mask) != 0
}

// SetCgif1 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCgif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCgif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCgif1Mask)
	}
}

const (
	RegisterIfcrFieldCtcif1Shift = 1
	RegisterIfcrFieldCtcif1Mask  = 0x2
)

// GetCtcif1 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCtcif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCtcif1Mask) != 0
}

// SetCtcif1 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCtcif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCtcif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCtcif1Mask)
	}
}

const (
	RegisterIfcrFieldChtif1Shift = 2
	RegisterIfcrFieldChtif1Mask  = 0x4
)

// GetChtif1 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) GetChtif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldChtif1Mask) != 0
}

// SetChtif1 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) SetChtif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldChtif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldChtif1Mask)
	}
}

const (
	RegisterIfcrFieldCteif1Shift = 3
	RegisterIfcrFieldCteif1Mask  = 0x8
)

// GetCteif1 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCteif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCteif1Mask) != 0
}

// SetCteif1 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCteif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCteif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCteif1Mask)
	}
}

const (
	RegisterIfcrFieldCgif2Shift = 4
	RegisterIfcrFieldCgif2Mask  = 0x10
)

// GetCgif2 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCgif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCgif2Mask) != 0
}

// SetCgif2 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCgif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCgif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCgif2Mask)
	}
}

const (
	RegisterIfcrFieldCtcif2Shift = 5
	RegisterIfcrFieldCtcif2Mask  = 0x20
)

// GetCtcif2 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCtcif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCtcif2Mask) != 0
}

// SetCtcif2 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCtcif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCtcif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCtcif2Mask)
	}
}

const (
	RegisterIfcrFieldChtif2Shift = 6
	RegisterIfcrFieldChtif2Mask  = 0x40
)

// GetChtif2 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) GetChtif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldChtif2Mask) != 0
}

// SetChtif2 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) SetChtif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldChtif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldChtif2Mask)
	}
}

const (
	RegisterIfcrFieldCteif2Shift = 7
	RegisterIfcrFieldCteif2Mask  = 0x80
)

// GetCteif2 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCteif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCteif2Mask) != 0
}

// SetCteif2 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCteif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCteif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCteif2Mask)
	}
}

const (
	RegisterIfcrFieldCgif3Shift = 8
	RegisterIfcrFieldCgif3Mask  = 0x100
)

// GetCgif3 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCgif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCgif3Mask) != 0
}

// SetCgif3 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCgif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCgif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCgif3Mask)
	}
}

const (
	RegisterIfcrFieldCtcif3Shift = 9
	RegisterIfcrFieldCtcif3Mask  = 0x200
)

// GetCtcif3 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCtcif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCtcif3Mask) != 0
}

// SetCtcif3 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCtcif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCtcif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCtcif3Mask)
	}
}

const (
	RegisterIfcrFieldChtif3Shift = 10
	RegisterIfcrFieldChtif3Mask  = 0x400
)

// GetChtif3 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) GetChtif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldChtif3Mask) != 0
}

// SetChtif3 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) SetChtif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldChtif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldChtif3Mask)
	}
}

const (
	RegisterIfcrFieldCteif3Shift = 11
	RegisterIfcrFieldCteif3Mask  = 0x800
)

// GetCteif3 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCteif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCteif3Mask) != 0
}

// SetCteif3 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCteif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCteif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCteif3Mask)
	}
}

const (
	RegisterIfcrFieldCgif4Shift = 12
	RegisterIfcrFieldCgif4Mask  = 0x1000
)

// GetCgif4 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCgif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCgif4Mask) != 0
}

// SetCgif4 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCgif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCgif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCgif4Mask)
	}
}

const (
	RegisterIfcrFieldCtcif4Shift = 13
	RegisterIfcrFieldCtcif4Mask  = 0x2000
)

// GetCtcif4 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCtcif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCtcif4Mask) != 0
}

// SetCtcif4 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCtcif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCtcif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCtcif4Mask)
	}
}

const (
	RegisterIfcrFieldChtif4Shift = 14
	RegisterIfcrFieldChtif4Mask  = 0x4000
)

// GetChtif4 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) GetChtif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldChtif4Mask) != 0
}

// SetChtif4 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) SetChtif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldChtif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldChtif4Mask)
	}
}

const (
	RegisterIfcrFieldCteif4Shift = 15
	RegisterIfcrFieldCteif4Mask  = 0x8000
)

// GetCteif4 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCteif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCteif4Mask) != 0
}

// SetCteif4 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCteif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCteif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCteif4Mask)
	}
}

const (
	RegisterIfcrFieldCgif5Shift = 16
	RegisterIfcrFieldCgif5Mask  = 0x10000
)

// GetCgif5 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCgif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCgif5Mask) != 0
}

// SetCgif5 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCgif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCgif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCgif5Mask)
	}
}

const (
	RegisterIfcrFieldCtcif5Shift = 17
	RegisterIfcrFieldCtcif5Mask  = 0x20000
)

// GetCtcif5 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCtcif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCtcif5Mask) != 0
}

// SetCtcif5 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCtcif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCtcif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCtcif5Mask)
	}
}

const (
	RegisterIfcrFieldChtif5Shift = 18
	RegisterIfcrFieldChtif5Mask  = 0x40000
)

// GetChtif5 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) GetChtif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldChtif5Mask) != 0
}

// SetChtif5 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) SetChtif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldChtif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldChtif5Mask)
	}
}

const (
	RegisterIfcrFieldCteif5Shift = 19
	RegisterIfcrFieldCteif5Mask  = 0x80000
)

// GetCteif5 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCteif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCteif5Mask) != 0
}

// SetCteif5 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCteif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCteif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCteif5Mask)
	}
}

const (
	RegisterIfcrFieldCgif6Shift = 20
	RegisterIfcrFieldCgif6Mask  = 0x100000
)

// GetCgif6 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCgif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCgif6Mask) != 0
}

// SetCgif6 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCgif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCgif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCgif6Mask)
	}
}

const (
	RegisterIfcrFieldCtcif6Shift = 21
	RegisterIfcrFieldCtcif6Mask  = 0x200000
)

// GetCtcif6 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCtcif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCtcif6Mask) != 0
}

// SetCtcif6 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCtcif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCtcif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCtcif6Mask)
	}
}

const (
	RegisterIfcrFieldChtif6Shift = 22
	RegisterIfcrFieldChtif6Mask  = 0x400000
)

// GetChtif6 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) GetChtif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldChtif6Mask) != 0
}

// SetChtif6 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) SetChtif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldChtif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldChtif6Mask)
	}
}

const (
	RegisterIfcrFieldCteif6Shift = 23
	RegisterIfcrFieldCteif6Mask  = 0x800000
)

// GetCteif6 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCteif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCteif6Mask) != 0
}

// SetCteif6 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCteif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCteif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCteif6Mask)
	}
}

const (
	RegisterIfcrFieldCgif7Shift = 24
	RegisterIfcrFieldCgif7Mask  = 0x1000000
)

// GetCgif7 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCgif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCgif7Mask) != 0
}

// SetCgif7 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCgif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCgif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCgif7Mask)
	}
}

const (
	RegisterIfcrFieldCtcif7Shift = 25
	RegisterIfcrFieldCtcif7Mask  = 0x2000000
)

// GetCtcif7 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCtcif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCtcif7Mask) != 0
}

// SetCtcif7 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCtcif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCtcif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCtcif7Mask)
	}
}

const (
	RegisterIfcrFieldChtif7Shift = 26
	RegisterIfcrFieldChtif7Mask  = 0x4000000
)

// GetChtif7 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) GetChtif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldChtif7Mask) != 0
}

// SetChtif7 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) SetChtif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldChtif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldChtif7Mask)
	}
}

const (
	RegisterIfcrFieldCteif7Shift = 27
	RegisterIfcrFieldCteif7Mask  = 0x8000000
)

// GetCteif7 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCteif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCteif7Mask) != 0
}

// SetCteif7 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCteif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCteif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCteif7Mask)
	}
}

const (
	RegisterIfcrFieldCgif8Shift = 28
	RegisterIfcrFieldCgif8Mask  = 0x10000000
)

// GetCgif8 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCgif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCgif8Mask) != 0
}

// SetCgif8 Channel x global interrupt clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCgif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCgif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCgif8Mask)
	}
}

const (
	RegisterIfcrFieldCtcif8Shift = 29
	RegisterIfcrFieldCtcif8Mask  = 0x20000000
)

// GetCtcif8 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCtcif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCtcif8Mask) != 0
}

// SetCtcif8 Channel x transfer complete clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCtcif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCtcif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCtcif8Mask)
	}
}

const (
	RegisterIfcrFieldChtif8Shift = 30
	RegisterIfcrFieldChtif8Mask  = 0x40000000
)

// GetChtif8 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) GetChtif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldChtif8Mask) != 0
}

// SetChtif8 Channel x half transfer clear This bit is set and cleared by software.
func (r *registerIfcrType) SetChtif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldChtif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldChtif8Mask)
	}
}

const (
	RegisterIfcrFieldCteif8Shift = 31
	RegisterIfcrFieldCteif8Mask  = 0x80000000
)

// GetCteif8 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) GetCteif8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCteif8Mask) != 0
}

// SetCteif8 Channel x transfer error clear This bit is set and cleared by software.
func (r *registerIfcrType) SetCteif8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCteif8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCteif8Mask)
	}
}

// registerCcr1Type DMA channel x configuration register
type registerCcr1Type uint32

const (
	RegisterCcr1FieldEnShift = 0
	RegisterCcr1FieldEnMask  = 0x1
)

// GetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr1Type) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldEnMask) != 0
}

// SetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr1Type) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldEnMask)
	}
}

const (
	RegisterCcr1FieldTcieShift = 1
	RegisterCcr1FieldTcieMask  = 0x2
)

// GetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr1Type) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr1Type) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldTcieMask)
	}
}

const (
	RegisterCcr1FieldHtieShift = 2
	RegisterCcr1FieldHtieMask  = 0x4
)

// GetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr1Type) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr1Type) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldHtieMask)
	}
}

const (
	RegisterCcr1FieldTeieShift = 3
	RegisterCcr1FieldTeieMask  = 0x8
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr1Type) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr1Type) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldTeieMask)
	}
}

const (
	RegisterCcr1FieldDirShift = 4
	RegisterCcr1FieldDirMask  = 0x10
)

// GetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr1Type) GetDir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldDirMask) != 0
}

// SetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr1Type) SetDir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldDirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldDirMask)
	}
}

const (
	RegisterCcr1FieldCircShift = 5
	RegisterCcr1FieldCircMask  = 0x20
)

// GetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr1Type) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldCircMask) != 0
}

// SetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr1Type) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldCircMask)
	}
}

const (
	RegisterCcr1FieldPincShift = 6
	RegisterCcr1FieldPincMask  = 0x40
)

// GetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr1Type) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldPincMask) != 0
}

// SetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr1Type) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldPincMask)
	}
}

const (
	RegisterCcr1FieldMincShift = 7
	RegisterCcr1FieldMincMask  = 0x80
)

// GetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr1Type) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldMincMask) != 0
}

// SetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr1Type) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldMincMask)
	}
}

const (
	RegisterCcr1FieldPsizeShift = 8
	RegisterCcr1FieldPsizeMask  = 0x300
)

// GetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr1Type) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldPsizeMask) >> RegisterCcr1FieldPsizeShift)
}

// SetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr1Type) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldPsizeMask)|(uint32(value)<<RegisterCcr1FieldPsizeShift))
}

const (
	RegisterCcr1FieldMsizeShift = 10
	RegisterCcr1FieldMsizeMask  = 0xc00
)

// GetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr1Type) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldMsizeMask) >> RegisterCcr1FieldMsizeShift)
}

// SetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr1Type) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldMsizeMask)|(uint32(value)<<RegisterCcr1FieldMsizeShift))
}

const (
	RegisterCcr1FieldPlShift = 12
	RegisterCcr1FieldPlMask  = 0x3000
)

// GetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr1Type) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldPlMask) >> RegisterCcr1FieldPlShift)
}

// SetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr1Type) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldPlMask)|(uint32(value)<<RegisterCcr1FieldPlShift))
}

const (
	RegisterCcr1FieldMem2memShift = 14
	RegisterCcr1FieldMem2memMask  = 0x4000
)

// GetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr1Type) GetMem2mem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldMem2memMask) != 0
}

// SetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr1Type) SetMem2mem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldMem2memMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldMem2memMask)
	}
}

// registerCndtr1Type DMA channel x number of data register
type registerCndtr1Type uint32

const (
	RegisterCndtr1FieldNdtShift = 0
	RegisterCndtr1FieldNdtMask  = 0xffff
)

// GetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr1Type) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCndtr1FieldNdtMask) >> RegisterCndtr1FieldNdtShift)
}

// SetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr1Type) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCndtr1FieldNdtMask)|(uint32(value)<<RegisterCndtr1FieldNdtShift))
}

// registerCpar1Type This register must not be written when the channel is enabled.
type registerCpar1Type uint32

const (
	RegisterCpar1FieldPaShift = 0
	RegisterCpar1FieldPaMask  = 0xffffffff
)

// GetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar1Type) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCpar1FieldPaMask) >> RegisterCpar1FieldPaShift)
}

// SetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar1Type) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpar1FieldPaMask)|(uint32(value)<<RegisterCpar1FieldPaShift))
}

// registerCmar1Type This register must not be written when the channel is enabled.
type registerCmar1Type uint32

const (
	RegisterCmar1FieldMaShift = 0
	RegisterCmar1FieldMaMask  = 0xffffffff
)

// GetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar1Type) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCmar1FieldMaMask) >> RegisterCmar1FieldMaShift)
}

// SetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar1Type) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmar1FieldMaMask)|(uint32(value)<<RegisterCmar1FieldMaShift))
}

// registerCcr2Type DMA channel x configuration register
type registerCcr2Type uint32

const (
	RegisterCcr2FieldEnShift = 0
	RegisterCcr2FieldEnMask  = 0x1
)

// GetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr2Type) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldEnMask) != 0
}

// SetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr2Type) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldEnMask)
	}
}

const (
	RegisterCcr2FieldTcieShift = 1
	RegisterCcr2FieldTcieMask  = 0x2
)

// GetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr2Type) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr2Type) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldTcieMask)
	}
}

const (
	RegisterCcr2FieldHtieShift = 2
	RegisterCcr2FieldHtieMask  = 0x4
)

// GetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr2Type) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr2Type) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldHtieMask)
	}
}

const (
	RegisterCcr2FieldTeieShift = 3
	RegisterCcr2FieldTeieMask  = 0x8
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr2Type) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr2Type) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldTeieMask)
	}
}

const (
	RegisterCcr2FieldDirShift = 4
	RegisterCcr2FieldDirMask  = 0x10
)

// GetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr2Type) GetDir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldDirMask) != 0
}

// SetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr2Type) SetDir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldDirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldDirMask)
	}
}

const (
	RegisterCcr2FieldCircShift = 5
	RegisterCcr2FieldCircMask  = 0x20
)

// GetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr2Type) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldCircMask) != 0
}

// SetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr2Type) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldCircMask)
	}
}

const (
	RegisterCcr2FieldPincShift = 6
	RegisterCcr2FieldPincMask  = 0x40
)

// GetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr2Type) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldPincMask) != 0
}

// SetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr2Type) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldPincMask)
	}
}

const (
	RegisterCcr2FieldMincShift = 7
	RegisterCcr2FieldMincMask  = 0x80
)

// GetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr2Type) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldMincMask) != 0
}

// SetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr2Type) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldMincMask)
	}
}

const (
	RegisterCcr2FieldPsizeShift = 8
	RegisterCcr2FieldPsizeMask  = 0x300
)

// GetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr2Type) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldPsizeMask) >> RegisterCcr2FieldPsizeShift)
}

// SetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr2Type) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldPsizeMask)|(uint32(value)<<RegisterCcr2FieldPsizeShift))
}

const (
	RegisterCcr2FieldMsizeShift = 10
	RegisterCcr2FieldMsizeMask  = 0xc00
)

// GetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr2Type) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldMsizeMask) >> RegisterCcr2FieldMsizeShift)
}

// SetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr2Type) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldMsizeMask)|(uint32(value)<<RegisterCcr2FieldMsizeShift))
}

const (
	RegisterCcr2FieldPlShift = 12
	RegisterCcr2FieldPlMask  = 0x3000
)

// GetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr2Type) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldPlMask) >> RegisterCcr2FieldPlShift)
}

// SetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr2Type) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldPlMask)|(uint32(value)<<RegisterCcr2FieldPlShift))
}

const (
	RegisterCcr2FieldMem2memShift = 14
	RegisterCcr2FieldMem2memMask  = 0x4000
)

// GetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr2Type) GetMem2mem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldMem2memMask) != 0
}

// SetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr2Type) SetMem2mem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldMem2memMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldMem2memMask)
	}
}

// registerCndtr2Type DMA channel x number of data register
type registerCndtr2Type uint32

const (
	RegisterCndtr2FieldNdtShift = 0
	RegisterCndtr2FieldNdtMask  = 0xffff
)

// GetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr2Type) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCndtr2FieldNdtMask) >> RegisterCndtr2FieldNdtShift)
}

// SetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr2Type) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCndtr2FieldNdtMask)|(uint32(value)<<RegisterCndtr2FieldNdtShift))
}

// registerCpar2Type This register must not be written when the channel is enabled.
type registerCpar2Type uint32

const (
	RegisterCpar2FieldPaShift = 0
	RegisterCpar2FieldPaMask  = 0xffffffff
)

// GetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar2Type) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCpar2FieldPaMask) >> RegisterCpar2FieldPaShift)
}

// SetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar2Type) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpar2FieldPaMask)|(uint32(value)<<RegisterCpar2FieldPaShift))
}

// registerCmar2Type This register must not be written when the channel is enabled.
type registerCmar2Type uint32

const (
	RegisterCmar2FieldMaShift = 0
	RegisterCmar2FieldMaMask  = 0xffffffff
)

// GetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar2Type) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCmar2FieldMaMask) >> RegisterCmar2FieldMaShift)
}

// SetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar2Type) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmar2FieldMaMask)|(uint32(value)<<RegisterCmar2FieldMaShift))
}

// registerCcr3Type DMA channel x configuration register
type registerCcr3Type uint32

const (
	RegisterCcr3FieldEnShift = 0
	RegisterCcr3FieldEnMask  = 0x1
)

// GetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr3Type) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldEnMask) != 0
}

// SetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr3Type) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr3FieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldEnMask)
	}
}

const (
	RegisterCcr3FieldTcieShift = 1
	RegisterCcr3FieldTcieMask  = 0x2
)

// GetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr3Type) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr3Type) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr3FieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldTcieMask)
	}
}

const (
	RegisterCcr3FieldHtieShift = 2
	RegisterCcr3FieldHtieMask  = 0x4
)

// GetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr3Type) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr3Type) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr3FieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldHtieMask)
	}
}

const (
	RegisterCcr3FieldTeieShift = 3
	RegisterCcr3FieldTeieMask  = 0x8
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr3Type) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr3Type) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr3FieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldTeieMask)
	}
}

const (
	RegisterCcr3FieldDirShift = 4
	RegisterCcr3FieldDirMask  = 0x10
)

// GetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr3Type) GetDir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldDirMask) != 0
}

// SetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr3Type) SetDir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr3FieldDirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldDirMask)
	}
}

const (
	RegisterCcr3FieldCircShift = 5
	RegisterCcr3FieldCircMask  = 0x20
)

// GetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr3Type) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldCircMask) != 0
}

// SetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr3Type) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr3FieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldCircMask)
	}
}

const (
	RegisterCcr3FieldPincShift = 6
	RegisterCcr3FieldPincMask  = 0x40
)

// GetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr3Type) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldPincMask) != 0
}

// SetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr3Type) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr3FieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldPincMask)
	}
}

const (
	RegisterCcr3FieldMincShift = 7
	RegisterCcr3FieldMincMask  = 0x80
)

// GetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr3Type) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldMincMask) != 0
}

// SetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr3Type) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr3FieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldMincMask)
	}
}

const (
	RegisterCcr3FieldPsizeShift = 8
	RegisterCcr3FieldPsizeMask  = 0x300
)

// GetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr3Type) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldPsizeMask) >> RegisterCcr3FieldPsizeShift)
}

// SetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr3Type) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldPsizeMask)|(uint32(value)<<RegisterCcr3FieldPsizeShift))
}

const (
	RegisterCcr3FieldMsizeShift = 10
	RegisterCcr3FieldMsizeMask  = 0xc00
)

// GetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr3Type) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldMsizeMask) >> RegisterCcr3FieldMsizeShift)
}

// SetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr3Type) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldMsizeMask)|(uint32(value)<<RegisterCcr3FieldMsizeShift))
}

const (
	RegisterCcr3FieldPlShift = 12
	RegisterCcr3FieldPlMask  = 0x3000
)

// GetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr3Type) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldPlMask) >> RegisterCcr3FieldPlShift)
}

// SetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr3Type) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldPlMask)|(uint32(value)<<RegisterCcr3FieldPlShift))
}

const (
	RegisterCcr3FieldMem2memShift = 14
	RegisterCcr3FieldMem2memMask  = 0x4000
)

// GetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr3Type) GetMem2mem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr3FieldMem2memMask) != 0
}

// SetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr3Type) SetMem2mem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr3FieldMem2memMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr3FieldMem2memMask)
	}
}

// registerCndtr3Type DMA channel x number of data register
type registerCndtr3Type uint32

const (
	RegisterCndtr3FieldNdtShift = 0
	RegisterCndtr3FieldNdtMask  = 0xffff
)

// GetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr3Type) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCndtr3FieldNdtMask) >> RegisterCndtr3FieldNdtShift)
}

// SetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr3Type) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCndtr3FieldNdtMask)|(uint32(value)<<RegisterCndtr3FieldNdtShift))
}

// registerCpar3Type This register must not be written when the channel is enabled.
type registerCpar3Type uint32

const (
	RegisterCpar3FieldPaShift = 0
	RegisterCpar3FieldPaMask  = 0xffffffff
)

// GetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar3Type) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCpar3FieldPaMask) >> RegisterCpar3FieldPaShift)
}

// SetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar3Type) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpar3FieldPaMask)|(uint32(value)<<RegisterCpar3FieldPaShift))
}

// registerCmar3Type This register must not be written when the channel is enabled.
type registerCmar3Type uint32

const (
	RegisterCmar3FieldMaShift = 0
	RegisterCmar3FieldMaMask  = 0xffffffff
)

// GetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar3Type) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCmar3FieldMaMask) >> RegisterCmar3FieldMaShift)
}

// SetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar3Type) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmar3FieldMaMask)|(uint32(value)<<RegisterCmar3FieldMaShift))
}

// registerCcr4Type DMA channel x configuration register
type registerCcr4Type uint32

const (
	RegisterCcr4FieldEnShift = 0
	RegisterCcr4FieldEnMask  = 0x1
)

// GetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr4Type) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldEnMask) != 0
}

// SetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr4Type) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr4FieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldEnMask)
	}
}

const (
	RegisterCcr4FieldTcieShift = 1
	RegisterCcr4FieldTcieMask  = 0x2
)

// GetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr4Type) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr4Type) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr4FieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldTcieMask)
	}
}

const (
	RegisterCcr4FieldHtieShift = 2
	RegisterCcr4FieldHtieMask  = 0x4
)

// GetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr4Type) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr4Type) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr4FieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldHtieMask)
	}
}

const (
	RegisterCcr4FieldTeieShift = 3
	RegisterCcr4FieldTeieMask  = 0x8
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr4Type) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr4Type) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr4FieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldTeieMask)
	}
}

const (
	RegisterCcr4FieldDirShift = 4
	RegisterCcr4FieldDirMask  = 0x10
)

// GetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr4Type) GetDir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldDirMask) != 0
}

// SetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr4Type) SetDir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr4FieldDirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldDirMask)
	}
}

const (
	RegisterCcr4FieldCircShift = 5
	RegisterCcr4FieldCircMask  = 0x20
)

// GetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr4Type) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldCircMask) != 0
}

// SetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr4Type) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr4FieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldCircMask)
	}
}

const (
	RegisterCcr4FieldPincShift = 6
	RegisterCcr4FieldPincMask  = 0x40
)

// GetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr4Type) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldPincMask) != 0
}

// SetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr4Type) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr4FieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldPincMask)
	}
}

const (
	RegisterCcr4FieldMincShift = 7
	RegisterCcr4FieldMincMask  = 0x80
)

// GetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr4Type) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldMincMask) != 0
}

// SetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr4Type) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr4FieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldMincMask)
	}
}

const (
	RegisterCcr4FieldPsizeShift = 8
	RegisterCcr4FieldPsizeMask  = 0x300
)

// GetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr4Type) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldPsizeMask) >> RegisterCcr4FieldPsizeShift)
}

// SetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr4Type) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldPsizeMask)|(uint32(value)<<RegisterCcr4FieldPsizeShift))
}

const (
	RegisterCcr4FieldMsizeShift = 10
	RegisterCcr4FieldMsizeMask  = 0xc00
)

// GetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr4Type) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldMsizeMask) >> RegisterCcr4FieldMsizeShift)
}

// SetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr4Type) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldMsizeMask)|(uint32(value)<<RegisterCcr4FieldMsizeShift))
}

const (
	RegisterCcr4FieldPlShift = 12
	RegisterCcr4FieldPlMask  = 0x3000
)

// GetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr4Type) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldPlMask) >> RegisterCcr4FieldPlShift)
}

// SetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr4Type) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldPlMask)|(uint32(value)<<RegisterCcr4FieldPlShift))
}

const (
	RegisterCcr4FieldMem2memShift = 14
	RegisterCcr4FieldMem2memMask  = 0x4000
)

// GetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr4Type) GetMem2mem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr4FieldMem2memMask) != 0
}

// SetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr4Type) SetMem2mem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr4FieldMem2memMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr4FieldMem2memMask)
	}
}

// registerCndtr4Type DMA channel x number of data register
type registerCndtr4Type uint32

const (
	RegisterCndtr4FieldNdtShift = 0
	RegisterCndtr4FieldNdtMask  = 0xffff
)

// GetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr4Type) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCndtr4FieldNdtMask) >> RegisterCndtr4FieldNdtShift)
}

// SetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr4Type) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCndtr4FieldNdtMask)|(uint32(value)<<RegisterCndtr4FieldNdtShift))
}

// registerCpar4Type This register must not be written when the channel is enabled.
type registerCpar4Type uint32

const (
	RegisterCpar4FieldPaShift = 0
	RegisterCpar4FieldPaMask  = 0xffffffff
)

// GetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar4Type) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCpar4FieldPaMask) >> RegisterCpar4FieldPaShift)
}

// SetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar4Type) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpar4FieldPaMask)|(uint32(value)<<RegisterCpar4FieldPaShift))
}

// registerCmar4Type This register must not be written when the channel is enabled.
type registerCmar4Type uint32

const (
	RegisterCmar4FieldMaShift = 0
	RegisterCmar4FieldMaMask  = 0xffffffff
)

// GetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar4Type) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCmar4FieldMaMask) >> RegisterCmar4FieldMaShift)
}

// SetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar4Type) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmar4FieldMaMask)|(uint32(value)<<RegisterCmar4FieldMaShift))
}

// registerCcr5Type DMA channel x configuration register
type registerCcr5Type uint32

const (
	RegisterCcr5FieldEnShift = 0
	RegisterCcr5FieldEnMask  = 0x1
)

// GetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr5Type) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldEnMask) != 0
}

// SetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr5Type) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldEnMask)
	}
}

const (
	RegisterCcr5FieldTcieShift = 1
	RegisterCcr5FieldTcieMask  = 0x2
)

// GetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr5Type) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr5Type) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldTcieMask)
	}
}

const (
	RegisterCcr5FieldHtieShift = 2
	RegisterCcr5FieldHtieMask  = 0x4
)

// GetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr5Type) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr5Type) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldHtieMask)
	}
}

const (
	RegisterCcr5FieldTeieShift = 3
	RegisterCcr5FieldTeieMask  = 0x8
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr5Type) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr5Type) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldTeieMask)
	}
}

const (
	RegisterCcr5FieldDirShift = 4
	RegisterCcr5FieldDirMask  = 0x10
)

// GetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr5Type) GetDir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldDirMask) != 0
}

// SetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr5Type) SetDir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldDirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldDirMask)
	}
}

const (
	RegisterCcr5FieldCircShift = 5
	RegisterCcr5FieldCircMask  = 0x20
)

// GetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr5Type) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldCircMask) != 0
}

// SetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr5Type) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldCircMask)
	}
}

const (
	RegisterCcr5FieldPincShift = 6
	RegisterCcr5FieldPincMask  = 0x40
)

// GetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr5Type) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldPincMask) != 0
}

// SetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr5Type) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldPincMask)
	}
}

const (
	RegisterCcr5FieldMincShift = 7
	RegisterCcr5FieldMincMask  = 0x80
)

// GetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr5Type) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldMincMask) != 0
}

// SetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr5Type) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldMincMask)
	}
}

const (
	RegisterCcr5FieldPsizeShift = 8
	RegisterCcr5FieldPsizeMask  = 0x300
)

// GetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr5Type) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldPsizeMask) >> RegisterCcr5FieldPsizeShift)
}

// SetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr5Type) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldPsizeMask)|(uint32(value)<<RegisterCcr5FieldPsizeShift))
}

const (
	RegisterCcr5FieldMsizeShift = 10
	RegisterCcr5FieldMsizeMask  = 0xc00
)

// GetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr5Type) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldMsizeMask) >> RegisterCcr5FieldMsizeShift)
}

// SetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr5Type) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldMsizeMask)|(uint32(value)<<RegisterCcr5FieldMsizeShift))
}

const (
	RegisterCcr5FieldPlShift = 12
	RegisterCcr5FieldPlMask  = 0x3000
)

// GetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr5Type) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldPlMask) >> RegisterCcr5FieldPlShift)
}

// SetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr5Type) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldPlMask)|(uint32(value)<<RegisterCcr5FieldPlShift))
}

const (
	RegisterCcr5FieldMem2memShift = 14
	RegisterCcr5FieldMem2memMask  = 0x4000
)

// GetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr5Type) GetMem2mem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr5FieldMem2memMask) != 0
}

// SetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr5Type) SetMem2mem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr5FieldMem2memMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr5FieldMem2memMask)
	}
}

// registerCndtr5Type DMA channel x number of data register
type registerCndtr5Type uint32

const (
	RegisterCndtr5FieldNdtShift = 0
	RegisterCndtr5FieldNdtMask  = 0xffff
)

// GetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr5Type) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCndtr5FieldNdtMask) >> RegisterCndtr5FieldNdtShift)
}

// SetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr5Type) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCndtr5FieldNdtMask)|(uint32(value)<<RegisterCndtr5FieldNdtShift))
}

// registerCpar5Type This register must not be written when the channel is enabled.
type registerCpar5Type uint32

const (
	RegisterCpar5FieldPaShift = 0
	RegisterCpar5FieldPaMask  = 0xffffffff
)

// GetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar5Type) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCpar5FieldPaMask) >> RegisterCpar5FieldPaShift)
}

// SetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar5Type) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpar5FieldPaMask)|(uint32(value)<<RegisterCpar5FieldPaShift))
}

// registerCmar5Type This register must not be written when the channel is enabled.
type registerCmar5Type uint32

const (
	RegisterCmar5FieldMaShift = 0
	RegisterCmar5FieldMaMask  = 0xffffffff
)

// GetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar5Type) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCmar5FieldMaMask) >> RegisterCmar5FieldMaShift)
}

// SetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar5Type) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmar5FieldMaMask)|(uint32(value)<<RegisterCmar5FieldMaShift))
}

// registerCcr6Type DMA channel x configuration register
type registerCcr6Type uint32

const (
	RegisterCcr6FieldEnShift = 0
	RegisterCcr6FieldEnMask  = 0x1
)

// GetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr6Type) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldEnMask) != 0
}

// SetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr6Type) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr6FieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldEnMask)
	}
}

const (
	RegisterCcr6FieldTcieShift = 1
	RegisterCcr6FieldTcieMask  = 0x2
)

// GetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr6Type) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr6Type) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr6FieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldTcieMask)
	}
}

const (
	RegisterCcr6FieldHtieShift = 2
	RegisterCcr6FieldHtieMask  = 0x4
)

// GetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr6Type) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr6Type) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr6FieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldHtieMask)
	}
}

const (
	RegisterCcr6FieldTeieShift = 3
	RegisterCcr6FieldTeieMask  = 0x8
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr6Type) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr6Type) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr6FieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldTeieMask)
	}
}

const (
	RegisterCcr6FieldDirShift = 4
	RegisterCcr6FieldDirMask  = 0x10
)

// GetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr6Type) GetDir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldDirMask) != 0
}

// SetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr6Type) SetDir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr6FieldDirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldDirMask)
	}
}

const (
	RegisterCcr6FieldCircShift = 5
	RegisterCcr6FieldCircMask  = 0x20
)

// GetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr6Type) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldCircMask) != 0
}

// SetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr6Type) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr6FieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldCircMask)
	}
}

const (
	RegisterCcr6FieldPincShift = 6
	RegisterCcr6FieldPincMask  = 0x40
)

// GetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr6Type) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldPincMask) != 0
}

// SetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr6Type) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr6FieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldPincMask)
	}
}

const (
	RegisterCcr6FieldMincShift = 7
	RegisterCcr6FieldMincMask  = 0x80
)

// GetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr6Type) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldMincMask) != 0
}

// SetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr6Type) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr6FieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldMincMask)
	}
}

const (
	RegisterCcr6FieldPsizeShift = 8
	RegisterCcr6FieldPsizeMask  = 0x300
)

// GetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr6Type) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldPsizeMask) >> RegisterCcr6FieldPsizeShift)
}

// SetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr6Type) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldPsizeMask)|(uint32(value)<<RegisterCcr6FieldPsizeShift))
}

const (
	RegisterCcr6FieldMsizeShift = 10
	RegisterCcr6FieldMsizeMask  = 0xc00
)

// GetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr6Type) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldMsizeMask) >> RegisterCcr6FieldMsizeShift)
}

// SetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr6Type) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldMsizeMask)|(uint32(value)<<RegisterCcr6FieldMsizeShift))
}

const (
	RegisterCcr6FieldPlShift = 12
	RegisterCcr6FieldPlMask  = 0x3000
)

// GetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr6Type) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldPlMask) >> RegisterCcr6FieldPlShift)
}

// SetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr6Type) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldPlMask)|(uint32(value)<<RegisterCcr6FieldPlShift))
}

const (
	RegisterCcr6FieldMem2memShift = 14
	RegisterCcr6FieldMem2memMask  = 0x4000
)

// GetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr6Type) GetMem2mem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr6FieldMem2memMask) != 0
}

// SetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr6Type) SetMem2mem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr6FieldMem2memMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr6FieldMem2memMask)
	}
}

// registerCndtr6Type DMA channel x number of data register
type registerCndtr6Type uint32

const (
	RegisterCndtr6FieldNdtShift = 0
	RegisterCndtr6FieldNdtMask  = 0xffff
)

// GetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr6Type) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCndtr6FieldNdtMask) >> RegisterCndtr6FieldNdtShift)
}

// SetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr6Type) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCndtr6FieldNdtMask)|(uint32(value)<<RegisterCndtr6FieldNdtShift))
}

// registerCpar6Type This register must not be written when the channel is enabled.
type registerCpar6Type uint32

const (
	RegisterCpar6FieldPaShift = 0
	RegisterCpar6FieldPaMask  = 0xffffffff
)

// GetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar6Type) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCpar6FieldPaMask) >> RegisterCpar6FieldPaShift)
}

// SetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar6Type) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpar6FieldPaMask)|(uint32(value)<<RegisterCpar6FieldPaShift))
}

// registerCmar6Type This register must not be written when the channel is enabled.
type registerCmar6Type uint32

const (
	RegisterCmar6FieldMaShift = 0
	RegisterCmar6FieldMaMask  = 0xffffffff
)

// GetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar6Type) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCmar6FieldMaMask) >> RegisterCmar6FieldMaShift)
}

// SetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar6Type) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmar6FieldMaMask)|(uint32(value)<<RegisterCmar6FieldMaShift))
}

// registerCcr7Type DMA channel x configuration register
type registerCcr7Type uint32

const (
	RegisterCcr7FieldEnShift = 0
	RegisterCcr7FieldEnMask  = 0x1
)

// GetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr7Type) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr7FieldEnMask) != 0
}

// SetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr7Type) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr7FieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr7FieldEnMask)
	}
}

const (
	RegisterCcr7FieldTcieShift = 1
	RegisterCcr7FieldTcieMask  = 0x2
)

// GetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr7Type) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr7FieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr7Type) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr7FieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr7FieldTcieMask)
	}
}

const (
	RegisterCcr7FieldHtieShift = 2
	RegisterCcr7FieldHtieMask  = 0x4
)

// GetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr7Type) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr7FieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr7Type) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr7FieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr7FieldHtieMask)
	}
}

const (
	RegisterCcr7FieldTeieShift = 3
	RegisterCcr7FieldTeieMask  = 0x8
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr7Type) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr7FieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr7Type) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr7FieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr7FieldTeieMask)
	}
}

const (
	RegisterCcr7FieldDirShift = 4
	RegisterCcr7FieldDirMask  = 0x10
)

// GetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr7Type) GetDir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr7FieldDirMask) != 0
}

// SetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr7Type) SetDir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr7FieldDirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr7FieldDirMask)
	}
}

const (
	RegisterCcr7FieldCircShift = 5
	RegisterCcr7FieldCircMask  = 0x20
)

// GetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr7Type) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr7FieldCircMask) != 0
}

// SetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr7Type) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr7FieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr7FieldCircMask)
	}
}

const (
	RegisterCcr7FieldPincShift = 6
	RegisterCcr7FieldPincMask  = 0x40
)

// GetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr7Type) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr7FieldPincMask) != 0
}

// SetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr7Type) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr7FieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr7FieldPincMask)
	}
}

const (
	RegisterCcr7FieldMincShift = 7
	RegisterCcr7FieldMincMask  = 0x80
)

// GetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr7Type) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr7FieldMincMask) != 0
}

// SetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr7Type) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr7FieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr7FieldMincMask)
	}
}

const (
	RegisterCcr7FieldPsizeShift = 8
	RegisterCcr7FieldPsizeMask  = 0x300
)

// GetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr7Type) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr7FieldPsizeMask) >> RegisterCcr7FieldPsizeShift)
}

// SetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr7Type) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr7FieldPsizeMask)|(uint32(value)<<RegisterCcr7FieldPsizeShift))
}

const (
	RegisterCcr7FieldMsizeShift = 10
	RegisterCcr7FieldMsizeMask  = 0xc00
)

// GetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr7Type) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr7FieldMsizeMask) >> RegisterCcr7FieldMsizeShift)
}

// SetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr7Type) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr7FieldMsizeMask)|(uint32(value)<<RegisterCcr7FieldMsizeShift))
}

const (
	RegisterCcr7FieldPlShift = 12
	RegisterCcr7FieldPlMask  = 0x3000
)

// GetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr7Type) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr7FieldPlMask) >> RegisterCcr7FieldPlShift)
}

// SetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr7Type) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr7FieldPlMask)|(uint32(value)<<RegisterCcr7FieldPlShift))
}

const (
	RegisterCcr7FieldMem2memShift = 14
	RegisterCcr7FieldMem2memMask  = 0x4000
)

// GetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr7Type) GetMem2mem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr7FieldMem2memMask) != 0
}

// SetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr7Type) SetMem2mem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr7FieldMem2memMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr7FieldMem2memMask)
	}
}

// registerCndtr7Type DMA channel x number of data register
type registerCndtr7Type uint32

const (
	RegisterCndtr7FieldNdtShift = 0
	RegisterCndtr7FieldNdtMask  = 0xffff
)

// GetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr7Type) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCndtr7FieldNdtMask) >> RegisterCndtr7FieldNdtShift)
}

// SetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr7Type) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCndtr7FieldNdtMask)|(uint32(value)<<RegisterCndtr7FieldNdtShift))
}

// registerCpar7Type This register must not be written when the channel is enabled.
type registerCpar7Type uint32

const (
	RegisterCpar7FieldPaShift = 0
	RegisterCpar7FieldPaMask  = 0xffffffff
)

// GetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar7Type) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCpar7FieldPaMask) >> RegisterCpar7FieldPaShift)
}

// SetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar7Type) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpar7FieldPaMask)|(uint32(value)<<RegisterCpar7FieldPaShift))
}

// registerCmar7Type This register must not be written when the channel is enabled.
type registerCmar7Type uint32

const (
	RegisterCmar7FieldMaShift = 0
	RegisterCmar7FieldMaMask  = 0xffffffff
)

// GetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar7Type) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCmar7FieldMaMask) >> RegisterCmar7FieldMaShift)
}

// SetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar7Type) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmar7FieldMaMask)|(uint32(value)<<RegisterCmar7FieldMaShift))
}

// registerCcr8Type DMA channel x configuration register
type registerCcr8Type uint32

const (
	RegisterCcr8FieldEnShift = 0
	RegisterCcr8FieldEnMask  = 0x1
)

// GetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr8Type) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr8FieldEnMask) != 0
}

// SetEn Channel enable This bit is set and cleared by software.
func (r *registerCcr8Type) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr8FieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr8FieldEnMask)
	}
}

const (
	RegisterCcr8FieldTcieShift = 1
	RegisterCcr8FieldTcieMask  = 0x2
)

// GetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr8Type) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr8FieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable This bit is set and cleared by software.
func (r *registerCcr8Type) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr8FieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr8FieldTcieMask)
	}
}

const (
	RegisterCcr8FieldHtieShift = 2
	RegisterCcr8FieldHtieMask  = 0x4
)

// GetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr8Type) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr8FieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable This bit is set and cleared by software.
func (r *registerCcr8Type) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr8FieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr8FieldHtieMask)
	}
}

const (
	RegisterCcr8FieldTeieShift = 3
	RegisterCcr8FieldTeieMask  = 0x8
)

// GetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr8Type) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr8FieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit is set and cleared by software.
func (r *registerCcr8Type) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr8FieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr8FieldTeieMask)
	}
}

const (
	RegisterCcr8FieldDirShift = 4
	RegisterCcr8FieldDirMask  = 0x10
)

// GetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr8Type) GetDir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr8FieldDirMask) != 0
}

// SetDir Data transfer direction This bit is set and cleared by software.
func (r *registerCcr8Type) SetDir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr8FieldDirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr8FieldDirMask)
	}
}

const (
	RegisterCcr8FieldCircShift = 5
	RegisterCcr8FieldCircMask  = 0x20
)

// GetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr8Type) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr8FieldCircMask) != 0
}

// SetCirc Circular mode This bit is set and cleared by software.
func (r *registerCcr8Type) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr8FieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr8FieldCircMask)
	}
}

const (
	RegisterCcr8FieldPincShift = 6
	RegisterCcr8FieldPincMask  = 0x40
)

// GetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr8Type) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr8FieldPincMask) != 0
}

// SetPinc Peripheral increment mode This bit is set and cleared by software.
func (r *registerCcr8Type) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr8FieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr8FieldPincMask)
	}
}

const (
	RegisterCcr8FieldMincShift = 7
	RegisterCcr8FieldMincMask  = 0x80
)

// GetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr8Type) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr8FieldMincMask) != 0
}

// SetMinc Memory increment mode This bit is set and cleared by software.
func (r *registerCcr8Type) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr8FieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr8FieldMincMask)
	}
}

const (
	RegisterCcr8FieldPsizeShift = 8
	RegisterCcr8FieldPsizeMask  = 0x300
)

// GetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr8Type) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr8FieldPsizeMask) >> RegisterCcr8FieldPsizeShift)
}

// SetPsize Peripheral size These bits are set and cleared by software.
func (r *registerCcr8Type) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr8FieldPsizeMask)|(uint32(value)<<RegisterCcr8FieldPsizeShift))
}

const (
	RegisterCcr8FieldMsizeShift = 10
	RegisterCcr8FieldMsizeMask  = 0xc00
)

// GetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr8Type) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr8FieldMsizeMask) >> RegisterCcr8FieldMsizeShift)
}

// SetMsize Memory size These bits are set and cleared by software.
func (r *registerCcr8Type) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr8FieldMsizeMask)|(uint32(value)<<RegisterCcr8FieldMsizeShift))
}

const (
	RegisterCcr8FieldPlShift = 12
	RegisterCcr8FieldPlMask  = 0x3000
)

// GetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr8Type) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcr8FieldPlMask) >> RegisterCcr8FieldPlShift)
}

// SetPl Channel priority level These bits are set and cleared by software.
func (r *registerCcr8Type) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcr8FieldPlMask)|(uint32(value)<<RegisterCcr8FieldPlShift))
}

const (
	RegisterCcr8FieldMem2memShift = 14
	RegisterCcr8FieldMem2memMask  = 0x4000
)

// GetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr8Type) GetMem2mem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr8FieldMem2memMask) != 0
}

// SetMem2mem Memory to memory mode This bit is set and cleared by software.
func (r *registerCcr8Type) SetMem2mem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr8FieldMem2memMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr8FieldMem2memMask)
	}
}

// registerCndtr8Type DMA channel x number of data register
type registerCndtr8Type uint32

const (
	RegisterCndtr8FieldNdtShift = 0
	RegisterCndtr8FieldNdtMask  = 0xffff
)

// GetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr8Type) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCndtr8FieldNdtMask) >> RegisterCndtr8FieldNdtShift)
}

// SetNdt Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
func (r *registerCndtr8Type) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCndtr8FieldNdtMask)|(uint32(value)<<RegisterCndtr8FieldNdtShift))
}

// registerCpar8Type This register must not be written when the channel is enabled.
type registerCpar8Type uint32

const (
	RegisterCpar8FieldPaShift = 0
	RegisterCpar8FieldPaMask  = 0xffffffff
)

// GetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar8Type) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCpar8FieldPaMask) >> RegisterCpar8FieldPaShift)
}

// SetPa Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCpar8Type) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpar8FieldPaMask)|(uint32(value)<<RegisterCpar8FieldPaShift))
}

// registerCmar8Type This register must not be written when the channel is enabled.
type registerCmar8Type uint32

const (
	RegisterCmar8FieldMaShift = 0
	RegisterCmar8FieldMaMask  = 0xffffffff
)

// GetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar8Type) GetMa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCmar8FieldMaMask) >> RegisterCmar8FieldMaShift)
}

// SetMa Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
func (r *registerCmar8Type) SetMa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmar8FieldMaMask)|(uint32(value)<<RegisterCmar8FieldMaShift))
}
