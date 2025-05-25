package dma

import (
	"unsafe"
	"volatile"
)

var (
	Dma1 = (*_dma)(unsafe.Pointer(uintptr(0x40020000)))
	Dma2 = (*_dma)(unsafe.Pointer(uintptr(0x40020400)))

	Instances = [2]*_dma{
		Dma1,
		Dma2,
	}
)

type _dma struct {
	Lisr   registerLisrType
	Hisr   registerHisrType
	Lifcr  registerLifcrType
	Hifcr  registerHifcrType
	S0cr   registerS0crType
	S0ndtr registerS0ndtrType
	S0par  registerS0parType
	S0m0ar registerS0m0arType
	S0m1ar registerS0m1arType
	S0fcr  registerS0fcrType
	S1cr   registerS1crType
	S1ndtr registerS1ndtrType
	S1par  registerS1parType
	S1m0ar registerS1m0arType
	S1m1ar registerS1m1arType
	S1fcr  registerS1fcrType
	S2cr   registerS2crType
	S2ndtr registerS2ndtrType
	S2par  registerS2parType
	S2m0ar registerS2m0arType
	S2m1ar registerS2m1arType
	S2fcr  registerS2fcrType
	S3cr   registerS3crType
	S3ndtr registerS3ndtrType
	S3par  registerS3parType
	S3m0ar registerS3m0arType
	S3m1ar registerS3m1arType
	S3fcr  registerS3fcrType
	S4cr   registerS4crType
	S4ndtr registerS4ndtrType
	S4par  registerS4parType
	S4m0ar registerS4m0arType
	S4m1ar registerS4m1arType
	S4fcr  registerS4fcrType
	S5cr   registerS5crType
	S5ndtr registerS5ndtrType
	S5par  registerS5parType
	S5m0ar registerS5m0arType
	S5m1ar registerS5m1arType
	S5fcr  registerS5fcrType
	S6cr   registerS6crType
	S6ndtr registerS6ndtrType
	S6par  registerS6parType
	S6m0ar registerS6m0arType
	S6m1ar registerS6m1arType
	S6fcr  registerS6fcrType
	S7cr   registerS7crType
	S7ndtr registerS7ndtrType
	S7par  registerS7parType
	S7m0ar registerS7m0arType
	S7m1ar registerS7m1arType
	S7fcr  registerS7fcrType
}

// registerLisrType low interrupt status register
type registerLisrType uint32

const (
	RegisterLisrFieldFeif0Shift = 0
	RegisterLisrFieldFeif0Mask  = 0x1
)

// GetFeif0 Stream x FIFO error interrupt flag (x=3..0)
func (r *registerLisrType) GetFeif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldFeif0Mask) != 0
}

// SetFeif0 Stream x FIFO error interrupt flag (x=3..0)
func (r *registerLisrType) SetFeif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldFeif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldFeif0Mask)
	}
}

const (
	RegisterLisrFieldDmeif0Shift = 2
	RegisterLisrFieldDmeif0Mask  = 0x4
)

// GetDmeif0 Stream x direct mode error interrupt flag (x=3..0)
func (r *registerLisrType) GetDmeif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldDmeif0Mask) != 0
}

// SetDmeif0 Stream x direct mode error interrupt flag (x=3..0)
func (r *registerLisrType) SetDmeif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldDmeif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldDmeif0Mask)
	}
}

const (
	RegisterLisrFieldTeif0Shift = 3
	RegisterLisrFieldTeif0Mask  = 0x8
)

// GetTeif0 Stream x transfer error interrupt flag (x=3..0)
func (r *registerLisrType) GetTeif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldTeif0Mask) != 0
}

// SetTeif0 Stream x transfer error interrupt flag (x=3..0)
func (r *registerLisrType) SetTeif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldTeif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldTeif0Mask)
	}
}

const (
	RegisterLisrFieldHtif0Shift = 4
	RegisterLisrFieldHtif0Mask  = 0x10
)

// GetHtif0 Stream x half transfer interrupt flag (x=3..0)
func (r *registerLisrType) GetHtif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldHtif0Mask) != 0
}

// SetHtif0 Stream x half transfer interrupt flag (x=3..0)
func (r *registerLisrType) SetHtif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldHtif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldHtif0Mask)
	}
}

const (
	RegisterLisrFieldTcif0Shift = 5
	RegisterLisrFieldTcif0Mask  = 0x20
)

// GetTcif0 Stream x transfer complete interrupt flag (x = 3..0)
func (r *registerLisrType) GetTcif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldTcif0Mask) != 0
}

// SetTcif0 Stream x transfer complete interrupt flag (x = 3..0)
func (r *registerLisrType) SetTcif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldTcif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldTcif0Mask)
	}
}

const (
	RegisterLisrFieldFeif1Shift = 6
	RegisterLisrFieldFeif1Mask  = 0x40
)

// GetFeif1 Stream x FIFO error interrupt flag (x=3..0)
func (r *registerLisrType) GetFeif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldFeif1Mask) != 0
}

// SetFeif1 Stream x FIFO error interrupt flag (x=3..0)
func (r *registerLisrType) SetFeif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldFeif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldFeif1Mask)
	}
}

const (
	RegisterLisrFieldDmeif1Shift = 8
	RegisterLisrFieldDmeif1Mask  = 0x100
)

// GetDmeif1 Stream x direct mode error interrupt flag (x=3..0)
func (r *registerLisrType) GetDmeif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldDmeif1Mask) != 0
}

// SetDmeif1 Stream x direct mode error interrupt flag (x=3..0)
func (r *registerLisrType) SetDmeif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldDmeif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldDmeif1Mask)
	}
}

const (
	RegisterLisrFieldTeif1Shift = 9
	RegisterLisrFieldTeif1Mask  = 0x200
)

// GetTeif1 Stream x transfer error interrupt flag (x=3..0)
func (r *registerLisrType) GetTeif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldTeif1Mask) != 0
}

// SetTeif1 Stream x transfer error interrupt flag (x=3..0)
func (r *registerLisrType) SetTeif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldTeif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldTeif1Mask)
	}
}

const (
	RegisterLisrFieldHtif1Shift = 10
	RegisterLisrFieldHtif1Mask  = 0x400
)

// GetHtif1 Stream x half transfer interrupt flag (x=3..0)
func (r *registerLisrType) GetHtif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldHtif1Mask) != 0
}

// SetHtif1 Stream x half transfer interrupt flag (x=3..0)
func (r *registerLisrType) SetHtif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldHtif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldHtif1Mask)
	}
}

const (
	RegisterLisrFieldTcif1Shift = 11
	RegisterLisrFieldTcif1Mask  = 0x800
)

// GetTcif1 Stream x transfer complete interrupt flag (x = 3..0)
func (r *registerLisrType) GetTcif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldTcif1Mask) != 0
}

// SetTcif1 Stream x transfer complete interrupt flag (x = 3..0)
func (r *registerLisrType) SetTcif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldTcif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldTcif1Mask)
	}
}

const (
	RegisterLisrFieldFeif2Shift = 16
	RegisterLisrFieldFeif2Mask  = 0x10000
)

// GetFeif2 Stream x FIFO error interrupt flag (x=3..0)
func (r *registerLisrType) GetFeif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldFeif2Mask) != 0
}

// SetFeif2 Stream x FIFO error interrupt flag (x=3..0)
func (r *registerLisrType) SetFeif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldFeif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldFeif2Mask)
	}
}

const (
	RegisterLisrFieldDmeif2Shift = 18
	RegisterLisrFieldDmeif2Mask  = 0x40000
)

// GetDmeif2 Stream x direct mode error interrupt flag (x=3..0)
func (r *registerLisrType) GetDmeif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldDmeif2Mask) != 0
}

// SetDmeif2 Stream x direct mode error interrupt flag (x=3..0)
func (r *registerLisrType) SetDmeif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldDmeif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldDmeif2Mask)
	}
}

const (
	RegisterLisrFieldTeif2Shift = 19
	RegisterLisrFieldTeif2Mask  = 0x80000
)

// GetTeif2 Stream x transfer error interrupt flag (x=3..0)
func (r *registerLisrType) GetTeif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldTeif2Mask) != 0
}

// SetTeif2 Stream x transfer error interrupt flag (x=3..0)
func (r *registerLisrType) SetTeif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldTeif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldTeif2Mask)
	}
}

const (
	RegisterLisrFieldHtif2Shift = 20
	RegisterLisrFieldHtif2Mask  = 0x100000
)

// GetHtif2 Stream x half transfer interrupt flag (x=3..0)
func (r *registerLisrType) GetHtif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldHtif2Mask) != 0
}

// SetHtif2 Stream x half transfer interrupt flag (x=3..0)
func (r *registerLisrType) SetHtif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldHtif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldHtif2Mask)
	}
}

const (
	RegisterLisrFieldTcif2Shift = 21
	RegisterLisrFieldTcif2Mask  = 0x200000
)

// GetTcif2 Stream x transfer complete interrupt flag (x = 3..0)
func (r *registerLisrType) GetTcif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldTcif2Mask) != 0
}

// SetTcif2 Stream x transfer complete interrupt flag (x = 3..0)
func (r *registerLisrType) SetTcif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldTcif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldTcif2Mask)
	}
}

const (
	RegisterLisrFieldFeif3Shift = 22
	RegisterLisrFieldFeif3Mask  = 0x400000
)

// GetFeif3 Stream x FIFO error interrupt flag (x=3..0)
func (r *registerLisrType) GetFeif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldFeif3Mask) != 0
}

// SetFeif3 Stream x FIFO error interrupt flag (x=3..0)
func (r *registerLisrType) SetFeif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldFeif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldFeif3Mask)
	}
}

const (
	RegisterLisrFieldDmeif3Shift = 24
	RegisterLisrFieldDmeif3Mask  = 0x1000000
)

// GetDmeif3 Stream x direct mode error interrupt flag (x=3..0)
func (r *registerLisrType) GetDmeif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldDmeif3Mask) != 0
}

// SetDmeif3 Stream x direct mode error interrupt flag (x=3..0)
func (r *registerLisrType) SetDmeif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldDmeif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldDmeif3Mask)
	}
}

const (
	RegisterLisrFieldTeif3Shift = 25
	RegisterLisrFieldTeif3Mask  = 0x2000000
)

// GetTeif3 Stream x transfer error interrupt flag (x=3..0)
func (r *registerLisrType) GetTeif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldTeif3Mask) != 0
}

// SetTeif3 Stream x transfer error interrupt flag (x=3..0)
func (r *registerLisrType) SetTeif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldTeif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldTeif3Mask)
	}
}

const (
	RegisterLisrFieldHtif3Shift = 26
	RegisterLisrFieldHtif3Mask  = 0x4000000
)

// GetHtif3 Stream x half transfer interrupt flag (x=3..0)
func (r *registerLisrType) GetHtif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldHtif3Mask) != 0
}

// SetHtif3 Stream x half transfer interrupt flag (x=3..0)
func (r *registerLisrType) SetHtif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldHtif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldHtif3Mask)
	}
}

const (
	RegisterLisrFieldTcif3Shift = 27
	RegisterLisrFieldTcif3Mask  = 0x8000000
)

// GetTcif3 Stream x transfer complete interrupt flag (x = 3..0)
func (r *registerLisrType) GetTcif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLisrFieldTcif3Mask) != 0
}

// SetTcif3 Stream x transfer complete interrupt flag (x = 3..0)
func (r *registerLisrType) SetTcif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLisrFieldTcif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLisrFieldTcif3Mask)
	}
}

// registerHisrType high interrupt status register
type registerHisrType uint32

const (
	RegisterHisrFieldFeif4Shift = 0
	RegisterHisrFieldFeif4Mask  = 0x1
)

// GetFeif4 Stream x FIFO error interrupt flag (x=7..4)
func (r *registerHisrType) GetFeif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldFeif4Mask) != 0
}

// SetFeif4 Stream x FIFO error interrupt flag (x=7..4)
func (r *registerHisrType) SetFeif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldFeif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldFeif4Mask)
	}
}

const (
	RegisterHisrFieldDmeif4Shift = 2
	RegisterHisrFieldDmeif4Mask  = 0x4
)

// GetDmeif4 Stream x direct mode error interrupt flag (x=7..4)
func (r *registerHisrType) GetDmeif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldDmeif4Mask) != 0
}

// SetDmeif4 Stream x direct mode error interrupt flag (x=7..4)
func (r *registerHisrType) SetDmeif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldDmeif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldDmeif4Mask)
	}
}

const (
	RegisterHisrFieldTeif4Shift = 3
	RegisterHisrFieldTeif4Mask  = 0x8
)

// GetTeif4 Stream x transfer error interrupt flag (x=7..4)
func (r *registerHisrType) GetTeif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldTeif4Mask) != 0
}

// SetTeif4 Stream x transfer error interrupt flag (x=7..4)
func (r *registerHisrType) SetTeif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldTeif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldTeif4Mask)
	}
}

const (
	RegisterHisrFieldHtif4Shift = 4
	RegisterHisrFieldHtif4Mask  = 0x10
)

// GetHtif4 Stream x half transfer interrupt flag (x=7..4)
func (r *registerHisrType) GetHtif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldHtif4Mask) != 0
}

// SetHtif4 Stream x half transfer interrupt flag (x=7..4)
func (r *registerHisrType) SetHtif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldHtif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldHtif4Mask)
	}
}

const (
	RegisterHisrFieldTcif4Shift = 5
	RegisterHisrFieldTcif4Mask  = 0x20
)

// GetTcif4 Stream x transfer complete interrupt flag (x=7..4)
func (r *registerHisrType) GetTcif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldTcif4Mask) != 0
}

// SetTcif4 Stream x transfer complete interrupt flag (x=7..4)
func (r *registerHisrType) SetTcif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldTcif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldTcif4Mask)
	}
}

const (
	RegisterHisrFieldFeif5Shift = 6
	RegisterHisrFieldFeif5Mask  = 0x40
)

// GetFeif5 Stream x FIFO error interrupt flag (x=7..4)
func (r *registerHisrType) GetFeif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldFeif5Mask) != 0
}

// SetFeif5 Stream x FIFO error interrupt flag (x=7..4)
func (r *registerHisrType) SetFeif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldFeif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldFeif5Mask)
	}
}

const (
	RegisterHisrFieldDmeif5Shift = 8
	RegisterHisrFieldDmeif5Mask  = 0x100
)

// GetDmeif5 Stream x direct mode error interrupt flag (x=7..4)
func (r *registerHisrType) GetDmeif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldDmeif5Mask) != 0
}

// SetDmeif5 Stream x direct mode error interrupt flag (x=7..4)
func (r *registerHisrType) SetDmeif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldDmeif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldDmeif5Mask)
	}
}

const (
	RegisterHisrFieldTeif5Shift = 9
	RegisterHisrFieldTeif5Mask  = 0x200
)

// GetTeif5 Stream x transfer error interrupt flag (x=7..4)
func (r *registerHisrType) GetTeif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldTeif5Mask) != 0
}

// SetTeif5 Stream x transfer error interrupt flag (x=7..4)
func (r *registerHisrType) SetTeif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldTeif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldTeif5Mask)
	}
}

const (
	RegisterHisrFieldHtif5Shift = 10
	RegisterHisrFieldHtif5Mask  = 0x400
)

// GetHtif5 Stream x half transfer interrupt flag (x=7..4)
func (r *registerHisrType) GetHtif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldHtif5Mask) != 0
}

// SetHtif5 Stream x half transfer interrupt flag (x=7..4)
func (r *registerHisrType) SetHtif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldHtif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldHtif5Mask)
	}
}

const (
	RegisterHisrFieldTcif5Shift = 11
	RegisterHisrFieldTcif5Mask  = 0x800
)

// GetTcif5 Stream x transfer complete interrupt flag (x=7..4)
func (r *registerHisrType) GetTcif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldTcif5Mask) != 0
}

// SetTcif5 Stream x transfer complete interrupt flag (x=7..4)
func (r *registerHisrType) SetTcif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldTcif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldTcif5Mask)
	}
}

const (
	RegisterHisrFieldFeif6Shift = 16
	RegisterHisrFieldFeif6Mask  = 0x10000
)

// GetFeif6 Stream x FIFO error interrupt flag (x=7..4)
func (r *registerHisrType) GetFeif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldFeif6Mask) != 0
}

// SetFeif6 Stream x FIFO error interrupt flag (x=7..4)
func (r *registerHisrType) SetFeif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldFeif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldFeif6Mask)
	}
}

const (
	RegisterHisrFieldDmeif6Shift = 18
	RegisterHisrFieldDmeif6Mask  = 0x40000
)

// GetDmeif6 Stream x direct mode error interrupt flag (x=7..4)
func (r *registerHisrType) GetDmeif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldDmeif6Mask) != 0
}

// SetDmeif6 Stream x direct mode error interrupt flag (x=7..4)
func (r *registerHisrType) SetDmeif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldDmeif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldDmeif6Mask)
	}
}

const (
	RegisterHisrFieldTeif6Shift = 19
	RegisterHisrFieldTeif6Mask  = 0x80000
)

// GetTeif6 Stream x transfer error interrupt flag (x=7..4)
func (r *registerHisrType) GetTeif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldTeif6Mask) != 0
}

// SetTeif6 Stream x transfer error interrupt flag (x=7..4)
func (r *registerHisrType) SetTeif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldTeif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldTeif6Mask)
	}
}

const (
	RegisterHisrFieldHtif6Shift = 20
	RegisterHisrFieldHtif6Mask  = 0x100000
)

// GetHtif6 Stream x half transfer interrupt flag (x=7..4)
func (r *registerHisrType) GetHtif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldHtif6Mask) != 0
}

// SetHtif6 Stream x half transfer interrupt flag (x=7..4)
func (r *registerHisrType) SetHtif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldHtif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldHtif6Mask)
	}
}

const (
	RegisterHisrFieldTcif6Shift = 21
	RegisterHisrFieldTcif6Mask  = 0x200000
)

// GetTcif6 Stream x transfer complete interrupt flag (x=7..4)
func (r *registerHisrType) GetTcif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldTcif6Mask) != 0
}

// SetTcif6 Stream x transfer complete interrupt flag (x=7..4)
func (r *registerHisrType) SetTcif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldTcif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldTcif6Mask)
	}
}

const (
	RegisterHisrFieldFeif7Shift = 22
	RegisterHisrFieldFeif7Mask  = 0x400000
)

// GetFeif7 Stream x FIFO error interrupt flag (x=7..4)
func (r *registerHisrType) GetFeif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldFeif7Mask) != 0
}

// SetFeif7 Stream x FIFO error interrupt flag (x=7..4)
func (r *registerHisrType) SetFeif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldFeif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldFeif7Mask)
	}
}

const (
	RegisterHisrFieldDmeif7Shift = 24
	RegisterHisrFieldDmeif7Mask  = 0x1000000
)

// GetDmeif7 Stream x direct mode error interrupt flag (x=7..4)
func (r *registerHisrType) GetDmeif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldDmeif7Mask) != 0
}

// SetDmeif7 Stream x direct mode error interrupt flag (x=7..4)
func (r *registerHisrType) SetDmeif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldDmeif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldDmeif7Mask)
	}
}

const (
	RegisterHisrFieldTeif7Shift = 25
	RegisterHisrFieldTeif7Mask  = 0x2000000
)

// GetTeif7 Stream x transfer error interrupt flag (x=7..4)
func (r *registerHisrType) GetTeif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldTeif7Mask) != 0
}

// SetTeif7 Stream x transfer error interrupt flag (x=7..4)
func (r *registerHisrType) SetTeif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldTeif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldTeif7Mask)
	}
}

const (
	RegisterHisrFieldHtif7Shift = 26
	RegisterHisrFieldHtif7Mask  = 0x4000000
)

// GetHtif7 Stream x half transfer interrupt flag (x=7..4)
func (r *registerHisrType) GetHtif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldHtif7Mask) != 0
}

// SetHtif7 Stream x half transfer interrupt flag (x=7..4)
func (r *registerHisrType) SetHtif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldHtif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldHtif7Mask)
	}
}

const (
	RegisterHisrFieldTcif7Shift = 27
	RegisterHisrFieldTcif7Mask  = 0x8000000
)

// GetTcif7 Stream x transfer complete interrupt flag (x=7..4)
func (r *registerHisrType) GetTcif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHisrFieldTcif7Mask) != 0
}

// SetTcif7 Stream x transfer complete interrupt flag (x=7..4)
func (r *registerHisrType) SetTcif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHisrFieldTcif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHisrFieldTcif7Mask)
	}
}

// registerLifcrType low interrupt flag clear register
type registerLifcrType uint32

const (
	RegisterLifcrFieldCfeif0Shift = 0
	RegisterLifcrFieldCfeif0Mask  = 0x1
)

// GetCfeif0 Stream x clear FIFO error interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCfeif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCfeif0Mask) != 0
}

// SetCfeif0 Stream x clear FIFO error interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCfeif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCfeif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCfeif0Mask)
	}
}

const (
	RegisterLifcrFieldCdmeif0Shift = 2
	RegisterLifcrFieldCdmeif0Mask  = 0x4
)

// GetCdmeif0 Stream x clear direct mode error interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCdmeif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCdmeif0Mask) != 0
}

// SetCdmeif0 Stream x clear direct mode error interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCdmeif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCdmeif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCdmeif0Mask)
	}
}

const (
	RegisterLifcrFieldCteif0Shift = 3
	RegisterLifcrFieldCteif0Mask  = 0x8
)

// GetCteif0 Stream x clear transfer error interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCteif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCteif0Mask) != 0
}

// SetCteif0 Stream x clear transfer error interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCteif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCteif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCteif0Mask)
	}
}

const (
	RegisterLifcrFieldChtif0Shift = 4
	RegisterLifcrFieldChtif0Mask  = 0x10
)

// GetChtif0 Stream x clear half transfer interrupt flag (x = 3..0)
func (r *registerLifcrType) GetChtif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldChtif0Mask) != 0
}

// SetChtif0 Stream x clear half transfer interrupt flag (x = 3..0)
func (r *registerLifcrType) SetChtif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldChtif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldChtif0Mask)
	}
}

const (
	RegisterLifcrFieldCtcif0Shift = 5
	RegisterLifcrFieldCtcif0Mask  = 0x20
)

// GetCtcif0 Stream x clear transfer complete interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCtcif0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCtcif0Mask) != 0
}

// SetCtcif0 Stream x clear transfer complete interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCtcif0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCtcif0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCtcif0Mask)
	}
}

const (
	RegisterLifcrFieldCfeif1Shift = 6
	RegisterLifcrFieldCfeif1Mask  = 0x40
)

// GetCfeif1 Stream x clear FIFO error interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCfeif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCfeif1Mask) != 0
}

// SetCfeif1 Stream x clear FIFO error interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCfeif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCfeif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCfeif1Mask)
	}
}

const (
	RegisterLifcrFieldCdmeif1Shift = 8
	RegisterLifcrFieldCdmeif1Mask  = 0x100
)

// GetCdmeif1 Stream x clear direct mode error interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCdmeif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCdmeif1Mask) != 0
}

// SetCdmeif1 Stream x clear direct mode error interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCdmeif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCdmeif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCdmeif1Mask)
	}
}

const (
	RegisterLifcrFieldCteif1Shift = 9
	RegisterLifcrFieldCteif1Mask  = 0x200
)

// GetCteif1 Stream x clear transfer error interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCteif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCteif1Mask) != 0
}

// SetCteif1 Stream x clear transfer error interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCteif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCteif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCteif1Mask)
	}
}

const (
	RegisterLifcrFieldChtif1Shift = 10
	RegisterLifcrFieldChtif1Mask  = 0x400
)

// GetChtif1 Stream x clear half transfer interrupt flag (x = 3..0)
func (r *registerLifcrType) GetChtif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldChtif1Mask) != 0
}

// SetChtif1 Stream x clear half transfer interrupt flag (x = 3..0)
func (r *registerLifcrType) SetChtif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldChtif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldChtif1Mask)
	}
}

const (
	RegisterLifcrFieldCtcif1Shift = 11
	RegisterLifcrFieldCtcif1Mask  = 0x800
)

// GetCtcif1 Stream x clear transfer complete interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCtcif1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCtcif1Mask) != 0
}

// SetCtcif1 Stream x clear transfer complete interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCtcif1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCtcif1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCtcif1Mask)
	}
}

const (
	RegisterLifcrFieldCfeif2Shift = 16
	RegisterLifcrFieldCfeif2Mask  = 0x10000
)

// GetCfeif2 Stream x clear FIFO error interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCfeif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCfeif2Mask) != 0
}

// SetCfeif2 Stream x clear FIFO error interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCfeif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCfeif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCfeif2Mask)
	}
}

const (
	RegisterLifcrFieldCdmeif2Shift = 18
	RegisterLifcrFieldCdmeif2Mask  = 0x40000
)

// GetCdmeif2 Stream x clear direct mode error interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCdmeif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCdmeif2Mask) != 0
}

// SetCdmeif2 Stream x clear direct mode error interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCdmeif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCdmeif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCdmeif2Mask)
	}
}

const (
	RegisterLifcrFieldCteif2Shift = 19
	RegisterLifcrFieldCteif2Mask  = 0x80000
)

// GetCteif2 Stream x clear transfer error interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCteif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCteif2Mask) != 0
}

// SetCteif2 Stream x clear transfer error interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCteif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCteif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCteif2Mask)
	}
}

const (
	RegisterLifcrFieldChtif2Shift = 20
	RegisterLifcrFieldChtif2Mask  = 0x100000
)

// GetChtif2 Stream x clear half transfer interrupt flag (x = 3..0)
func (r *registerLifcrType) GetChtif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldChtif2Mask) != 0
}

// SetChtif2 Stream x clear half transfer interrupt flag (x = 3..0)
func (r *registerLifcrType) SetChtif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldChtif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldChtif2Mask)
	}
}

const (
	RegisterLifcrFieldCtcif2Shift = 21
	RegisterLifcrFieldCtcif2Mask  = 0x200000
)

// GetCtcif2 Stream x clear transfer complete interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCtcif2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCtcif2Mask) != 0
}

// SetCtcif2 Stream x clear transfer complete interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCtcif2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCtcif2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCtcif2Mask)
	}
}

const (
	RegisterLifcrFieldCfeif3Shift = 22
	RegisterLifcrFieldCfeif3Mask  = 0x400000
)

// GetCfeif3 Stream x clear FIFO error interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCfeif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCfeif3Mask) != 0
}

// SetCfeif3 Stream x clear FIFO error interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCfeif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCfeif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCfeif3Mask)
	}
}

const (
	RegisterLifcrFieldCdmeif3Shift = 24
	RegisterLifcrFieldCdmeif3Mask  = 0x1000000
)

// GetCdmeif3 Stream x clear direct mode error interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCdmeif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCdmeif3Mask) != 0
}

// SetCdmeif3 Stream x clear direct mode error interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCdmeif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCdmeif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCdmeif3Mask)
	}
}

const (
	RegisterLifcrFieldCteif3Shift = 25
	RegisterLifcrFieldCteif3Mask  = 0x2000000
)

// GetCteif3 Stream x clear transfer error interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCteif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCteif3Mask) != 0
}

// SetCteif3 Stream x clear transfer error interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCteif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCteif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCteif3Mask)
	}
}

const (
	RegisterLifcrFieldChtif3Shift = 26
	RegisterLifcrFieldChtif3Mask  = 0x4000000
)

// GetChtif3 Stream x clear half transfer interrupt flag (x = 3..0)
func (r *registerLifcrType) GetChtif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldChtif3Mask) != 0
}

// SetChtif3 Stream x clear half transfer interrupt flag (x = 3..0)
func (r *registerLifcrType) SetChtif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldChtif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldChtif3Mask)
	}
}

const (
	RegisterLifcrFieldCtcif3Shift = 27
	RegisterLifcrFieldCtcif3Mask  = 0x8000000
)

// GetCtcif3 Stream x clear transfer complete interrupt flag (x = 3..0)
func (r *registerLifcrType) GetCtcif3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterLifcrFieldCtcif3Mask) != 0
}

// SetCtcif3 Stream x clear transfer complete interrupt flag (x = 3..0)
func (r *registerLifcrType) SetCtcif3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterLifcrFieldCtcif3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterLifcrFieldCtcif3Mask)
	}
}

// registerHifcrType high interrupt flag clear register
type registerHifcrType uint32

const (
	RegisterHifcrFieldCfeif4Shift = 0
	RegisterHifcrFieldCfeif4Mask  = 0x1
)

// GetCfeif4 Stream x clear FIFO error interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCfeif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCfeif4Mask) != 0
}

// SetCfeif4 Stream x clear FIFO error interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCfeif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCfeif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCfeif4Mask)
	}
}

const (
	RegisterHifcrFieldCdmeif4Shift = 2
	RegisterHifcrFieldCdmeif4Mask  = 0x4
)

// GetCdmeif4 Stream x clear direct mode error interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCdmeif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCdmeif4Mask) != 0
}

// SetCdmeif4 Stream x clear direct mode error interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCdmeif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCdmeif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCdmeif4Mask)
	}
}

const (
	RegisterHifcrFieldCteif4Shift = 3
	RegisterHifcrFieldCteif4Mask  = 0x8
)

// GetCteif4 Stream x clear transfer error interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCteif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCteif4Mask) != 0
}

// SetCteif4 Stream x clear transfer error interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCteif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCteif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCteif4Mask)
	}
}

const (
	RegisterHifcrFieldChtif4Shift = 4
	RegisterHifcrFieldChtif4Mask  = 0x10
)

// GetChtif4 Stream x clear half transfer interrupt flag (x = 7..4)
func (r *registerHifcrType) GetChtif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldChtif4Mask) != 0
}

// SetChtif4 Stream x clear half transfer interrupt flag (x = 7..4)
func (r *registerHifcrType) SetChtif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldChtif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldChtif4Mask)
	}
}

const (
	RegisterHifcrFieldCtcif4Shift = 5
	RegisterHifcrFieldCtcif4Mask  = 0x20
)

// GetCtcif4 Stream x clear transfer complete interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCtcif4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCtcif4Mask) != 0
}

// SetCtcif4 Stream x clear transfer complete interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCtcif4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCtcif4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCtcif4Mask)
	}
}

const (
	RegisterHifcrFieldCfeif5Shift = 6
	RegisterHifcrFieldCfeif5Mask  = 0x40
)

// GetCfeif5 Stream x clear FIFO error interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCfeif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCfeif5Mask) != 0
}

// SetCfeif5 Stream x clear FIFO error interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCfeif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCfeif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCfeif5Mask)
	}
}

const (
	RegisterHifcrFieldCdmeif5Shift = 8
	RegisterHifcrFieldCdmeif5Mask  = 0x100
)

// GetCdmeif5 Stream x clear direct mode error interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCdmeif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCdmeif5Mask) != 0
}

// SetCdmeif5 Stream x clear direct mode error interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCdmeif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCdmeif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCdmeif5Mask)
	}
}

const (
	RegisterHifcrFieldCteif5Shift = 9
	RegisterHifcrFieldCteif5Mask  = 0x200
)

// GetCteif5 Stream x clear transfer error interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCteif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCteif5Mask) != 0
}

// SetCteif5 Stream x clear transfer error interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCteif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCteif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCteif5Mask)
	}
}

const (
	RegisterHifcrFieldChtif5Shift = 10
	RegisterHifcrFieldChtif5Mask  = 0x400
)

// GetChtif5 Stream x clear half transfer interrupt flag (x = 7..4)
func (r *registerHifcrType) GetChtif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldChtif5Mask) != 0
}

// SetChtif5 Stream x clear half transfer interrupt flag (x = 7..4)
func (r *registerHifcrType) SetChtif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldChtif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldChtif5Mask)
	}
}

const (
	RegisterHifcrFieldCtcif5Shift = 11
	RegisterHifcrFieldCtcif5Mask  = 0x800
)

// GetCtcif5 Stream x clear transfer complete interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCtcif5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCtcif5Mask) != 0
}

// SetCtcif5 Stream x clear transfer complete interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCtcif5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCtcif5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCtcif5Mask)
	}
}

const (
	RegisterHifcrFieldCfeif6Shift = 16
	RegisterHifcrFieldCfeif6Mask  = 0x10000
)

// GetCfeif6 Stream x clear FIFO error interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCfeif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCfeif6Mask) != 0
}

// SetCfeif6 Stream x clear FIFO error interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCfeif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCfeif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCfeif6Mask)
	}
}

const (
	RegisterHifcrFieldCdmeif6Shift = 18
	RegisterHifcrFieldCdmeif6Mask  = 0x40000
)

// GetCdmeif6 Stream x clear direct mode error interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCdmeif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCdmeif6Mask) != 0
}

// SetCdmeif6 Stream x clear direct mode error interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCdmeif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCdmeif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCdmeif6Mask)
	}
}

const (
	RegisterHifcrFieldCteif6Shift = 19
	RegisterHifcrFieldCteif6Mask  = 0x80000
)

// GetCteif6 Stream x clear transfer error interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCteif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCteif6Mask) != 0
}

// SetCteif6 Stream x clear transfer error interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCteif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCteif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCteif6Mask)
	}
}

const (
	RegisterHifcrFieldChtif6Shift = 20
	RegisterHifcrFieldChtif6Mask  = 0x100000
)

// GetChtif6 Stream x clear half transfer interrupt flag (x = 7..4)
func (r *registerHifcrType) GetChtif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldChtif6Mask) != 0
}

// SetChtif6 Stream x clear half transfer interrupt flag (x = 7..4)
func (r *registerHifcrType) SetChtif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldChtif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldChtif6Mask)
	}
}

const (
	RegisterHifcrFieldCtcif6Shift = 21
	RegisterHifcrFieldCtcif6Mask  = 0x200000
)

// GetCtcif6 Stream x clear transfer complete interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCtcif6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCtcif6Mask) != 0
}

// SetCtcif6 Stream x clear transfer complete interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCtcif6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCtcif6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCtcif6Mask)
	}
}

const (
	RegisterHifcrFieldCfeif7Shift = 22
	RegisterHifcrFieldCfeif7Mask  = 0x400000
)

// GetCfeif7 Stream x clear FIFO error interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCfeif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCfeif7Mask) != 0
}

// SetCfeif7 Stream x clear FIFO error interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCfeif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCfeif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCfeif7Mask)
	}
}

const (
	RegisterHifcrFieldCdmeif7Shift = 24
	RegisterHifcrFieldCdmeif7Mask  = 0x1000000
)

// GetCdmeif7 Stream x clear direct mode error interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCdmeif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCdmeif7Mask) != 0
}

// SetCdmeif7 Stream x clear direct mode error interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCdmeif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCdmeif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCdmeif7Mask)
	}
}

const (
	RegisterHifcrFieldCteif7Shift = 25
	RegisterHifcrFieldCteif7Mask  = 0x2000000
)

// GetCteif7 Stream x clear transfer error interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCteif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCteif7Mask) != 0
}

// SetCteif7 Stream x clear transfer error interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCteif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCteif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCteif7Mask)
	}
}

const (
	RegisterHifcrFieldChtif7Shift = 26
	RegisterHifcrFieldChtif7Mask  = 0x4000000
)

// GetChtif7 Stream x clear half transfer interrupt flag (x = 7..4)
func (r *registerHifcrType) GetChtif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldChtif7Mask) != 0
}

// SetChtif7 Stream x clear half transfer interrupt flag (x = 7..4)
func (r *registerHifcrType) SetChtif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldChtif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldChtif7Mask)
	}
}

const (
	RegisterHifcrFieldCtcif7Shift = 27
	RegisterHifcrFieldCtcif7Mask  = 0x8000000
)

// GetCtcif7 Stream x clear transfer complete interrupt flag (x = 7..4)
func (r *registerHifcrType) GetCtcif7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHifcrFieldCtcif7Mask) != 0
}

// SetCtcif7 Stream x clear transfer complete interrupt flag (x = 7..4)
func (r *registerHifcrType) SetCtcif7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHifcrFieldCtcif7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHifcrFieldCtcif7Mask)
	}
}

// registerS0crType stream x configuration register
type registerS0crType uint32

const (
	RegisterS0crFieldEnShift = 0
	RegisterS0crFieldEnMask  = 0x1
)

// GetEn Stream enable / flag stream ready when read low
func (r *registerS0crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldEnMask) != 0
}

// SetEn Stream enable / flag stream ready when read low
func (r *registerS0crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldEnMask)
	}
}

const (
	RegisterS0crFieldDmeieShift = 1
	RegisterS0crFieldDmeieMask  = 0x2
)

// GetDmeie Direct mode error interrupt enable
func (r *registerS0crType) GetDmeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldDmeieMask) != 0
}

// SetDmeie Direct mode error interrupt enable
func (r *registerS0crType) SetDmeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0crFieldDmeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldDmeieMask)
	}
}

const (
	RegisterS0crFieldTeieShift = 2
	RegisterS0crFieldTeieMask  = 0x4
)

// GetTeie Transfer error interrupt enable
func (r *registerS0crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable
func (r *registerS0crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldTeieMask)
	}
}

const (
	RegisterS0crFieldHtieShift = 3
	RegisterS0crFieldHtieMask  = 0x8
)

// GetHtie Half transfer interrupt enable
func (r *registerS0crType) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable
func (r *registerS0crType) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0crFieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldHtieMask)
	}
}

const (
	RegisterS0crFieldTcieShift = 4
	RegisterS0crFieldTcieMask  = 0x10
)

// GetTcie Transfer complete interrupt enable
func (r *registerS0crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable
func (r *registerS0crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldTcieMask)
	}
}

const (
	RegisterS0crFieldPfctrlShift = 5
	RegisterS0crFieldPfctrlMask  = 0x20
)

// GetPfctrl Peripheral flow controller
func (r *registerS0crType) GetPfctrl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldPfctrlMask) != 0
}

// SetPfctrl Peripheral flow controller
func (r *registerS0crType) SetPfctrl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0crFieldPfctrlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldPfctrlMask)
	}
}

const (
	RegisterS0crFieldDirShift = 6
	RegisterS0crFieldDirMask  = 0xc0
)

// GetDir Data transfer direction
func (r *registerS0crType) GetDir() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldDirMask) >> RegisterS0crFieldDirShift)
}

// SetDir Data transfer direction
func (r *registerS0crType) SetDir(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldDirMask)|(uint32(value)<<RegisterS0crFieldDirShift))
}

const (
	RegisterS0crFieldCircShift = 8
	RegisterS0crFieldCircMask  = 0x100
)

// GetCirc Circular mode
func (r *registerS0crType) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldCircMask) != 0
}

// SetCirc Circular mode
func (r *registerS0crType) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0crFieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldCircMask)
	}
}

const (
	RegisterS0crFieldPincShift = 9
	RegisterS0crFieldPincMask  = 0x200
)

// GetPinc Peripheral increment mode
func (r *registerS0crType) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldPincMask) != 0
}

// SetPinc Peripheral increment mode
func (r *registerS0crType) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0crFieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldPincMask)
	}
}

const (
	RegisterS0crFieldMincShift = 10
	RegisterS0crFieldMincMask  = 0x400
)

// GetMinc Memory increment mode
func (r *registerS0crType) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldMincMask) != 0
}

// SetMinc Memory increment mode
func (r *registerS0crType) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0crFieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldMincMask)
	}
}

const (
	RegisterS0crFieldPsizeShift = 11
	RegisterS0crFieldPsizeMask  = 0x1800
)

// GetPsize Peripheral data size
func (r *registerS0crType) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldPsizeMask) >> RegisterS0crFieldPsizeShift)
}

// SetPsize Peripheral data size
func (r *registerS0crType) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldPsizeMask)|(uint32(value)<<RegisterS0crFieldPsizeShift))
}

const (
	RegisterS0crFieldMsizeShift = 13
	RegisterS0crFieldMsizeMask  = 0x6000
)

// GetMsize Memory data size
func (r *registerS0crType) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldMsizeMask) >> RegisterS0crFieldMsizeShift)
}

// SetMsize Memory data size
func (r *registerS0crType) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldMsizeMask)|(uint32(value)<<RegisterS0crFieldMsizeShift))
}

const (
	RegisterS0crFieldPincosShift = 15
	RegisterS0crFieldPincosMask  = 0x8000
)

// GetPincos Peripheral increment offset size
func (r *registerS0crType) GetPincos() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldPincosMask) != 0
}

// SetPincos Peripheral increment offset size
func (r *registerS0crType) SetPincos(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0crFieldPincosMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldPincosMask)
	}
}

const (
	RegisterS0crFieldPlShift = 16
	RegisterS0crFieldPlMask  = 0x30000
)

// GetPl Priority level
func (r *registerS0crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldPlMask) >> RegisterS0crFieldPlShift)
}

// SetPl Priority level
func (r *registerS0crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldPlMask)|(uint32(value)<<RegisterS0crFieldPlShift))
}

const (
	RegisterS0crFieldDbmShift = 18
	RegisterS0crFieldDbmMask  = 0x40000
)

// GetDbm Double buffer mode
func (r *registerS0crType) GetDbm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldDbmMask) != 0
}

// SetDbm Double buffer mode
func (r *registerS0crType) SetDbm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0crFieldDbmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldDbmMask)
	}
}

const (
	RegisterS0crFieldCtShift = 19
	RegisterS0crFieldCtMask  = 0x80000
)

// GetCt Current target (only in double buffer mode)
func (r *registerS0crType) GetCt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldCtMask) != 0
}

// SetCt Current target (only in double buffer mode)
func (r *registerS0crType) SetCt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0crFieldCtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldCtMask)
	}
}

const (
	RegisterS0crFieldPburstShift = 21
	RegisterS0crFieldPburstMask  = 0x600000
)

// GetPburst Peripheral burst transfer configuration
func (r *registerS0crType) GetPburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldPburstMask) >> RegisterS0crFieldPburstShift)
}

// SetPburst Peripheral burst transfer configuration
func (r *registerS0crType) SetPburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldPburstMask)|(uint32(value)<<RegisterS0crFieldPburstShift))
}

const (
	RegisterS0crFieldMburstShift = 23
	RegisterS0crFieldMburstMask  = 0x1800000
)

// GetMburst Memory burst transfer configuration
func (r *registerS0crType) GetMburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS0crFieldMburstMask) >> RegisterS0crFieldMburstShift)
}

// SetMburst Memory burst transfer configuration
func (r *registerS0crType) SetMburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS0crFieldMburstMask)|(uint32(value)<<RegisterS0crFieldMburstShift))
}

// registerS0ndtrType stream x number of data register
type registerS0ndtrType uint32

const (
	RegisterS0ndtrFieldNdtShift = 0
	RegisterS0ndtrFieldNdtMask  = 0xffff
)

// GetNdt Number of data items to transfer
func (r *registerS0ndtrType) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterS0ndtrFieldNdtMask) >> RegisterS0ndtrFieldNdtShift)
}

// SetNdt Number of data items to transfer
func (r *registerS0ndtrType) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS0ndtrFieldNdtMask)|(uint32(value)<<RegisterS0ndtrFieldNdtShift))
}

// registerS0parType stream x peripheral address register
type registerS0parType uint32

const (
	RegisterS0parFieldPaShift = 0
	RegisterS0parFieldPaMask  = 0xffffffff
)

// GetPa Peripheral address
func (r *registerS0parType) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS0parFieldPaMask) >> RegisterS0parFieldPaShift)
}

// SetPa Peripheral address
func (r *registerS0parType) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS0parFieldPaMask)|(uint32(value)<<RegisterS0parFieldPaShift))
}

// registerS0m0arType stream x memory 0 address register
type registerS0m0arType uint32

const (
	RegisterS0m0arFieldM0aShift = 0
	RegisterS0m0arFieldM0aMask  = 0xffffffff
)

// GetM0a Memory 0 address
func (r *registerS0m0arType) GetM0a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS0m0arFieldM0aMask) >> RegisterS0m0arFieldM0aShift)
}

// SetM0a Memory 0 address
func (r *registerS0m0arType) SetM0a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS0m0arFieldM0aMask)|(uint32(value)<<RegisterS0m0arFieldM0aShift))
}

// registerS0m1arType stream x memory 1 address register
type registerS0m1arType uint32

const (
	RegisterS0m1arFieldM1aShift = 0
	RegisterS0m1arFieldM1aMask  = 0xffffffff
)

// GetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS0m1arType) GetM1a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS0m1arFieldM1aMask) >> RegisterS0m1arFieldM1aShift)
}

// SetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS0m1arType) SetM1a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS0m1arFieldM1aMask)|(uint32(value)<<RegisterS0m1arFieldM1aShift))
}

// registerS0fcrType stream x FIFO control register
type registerS0fcrType uint32

const (
	RegisterS0fcrFieldFthShift = 0
	RegisterS0fcrFieldFthMask  = 0x3
)

// GetFth FIFO threshold selection
func (r *registerS0fcrType) GetFth() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS0fcrFieldFthMask) >> RegisterS0fcrFieldFthShift)
}

// SetFth FIFO threshold selection
func (r *registerS0fcrType) SetFth(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS0fcrFieldFthMask)|(uint32(value)<<RegisterS0fcrFieldFthShift))
}

const (
	RegisterS0fcrFieldDmdisShift = 2
	RegisterS0fcrFieldDmdisMask  = 0x4
)

// GetDmdis Direct mode disable
func (r *registerS0fcrType) GetDmdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0fcrFieldDmdisMask) != 0
}

// SetDmdis Direct mode disable
func (r *registerS0fcrType) SetDmdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0fcrFieldDmdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0fcrFieldDmdisMask)
	}
}

const (
	RegisterS0fcrFieldFsShift = 3
	RegisterS0fcrFieldFsMask  = 0x38
)

// GetFs FIFO status
func (r *registerS0fcrType) GetFs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS0fcrFieldFsMask) >> RegisterS0fcrFieldFsShift)
}

const (
	RegisterS0fcrFieldFeieShift = 7
	RegisterS0fcrFieldFeieMask  = 0x80
)

// GetFeie FIFO error interrupt enable
func (r *registerS0fcrType) GetFeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS0fcrFieldFeieMask) != 0
}

// SetFeie FIFO error interrupt enable
func (r *registerS0fcrType) SetFeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS0fcrFieldFeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS0fcrFieldFeieMask)
	}
}

// registerS1crType stream x configuration register
type registerS1crType uint32

const (
	RegisterS1crFieldEnShift = 0
	RegisterS1crFieldEnMask  = 0x1
)

// GetEn Stream enable / flag stream ready when read low
func (r *registerS1crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldEnMask) != 0
}

// SetEn Stream enable / flag stream ready when read low
func (r *registerS1crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldEnMask)
	}
}

const (
	RegisterS1crFieldDmeieShift = 1
	RegisterS1crFieldDmeieMask  = 0x2
)

// GetDmeie Direct mode error interrupt enable
func (r *registerS1crType) GetDmeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldDmeieMask) != 0
}

// SetDmeie Direct mode error interrupt enable
func (r *registerS1crType) SetDmeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldDmeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldDmeieMask)
	}
}

const (
	RegisterS1crFieldTeieShift = 2
	RegisterS1crFieldTeieMask  = 0x4
)

// GetTeie Transfer error interrupt enable
func (r *registerS1crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable
func (r *registerS1crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldTeieMask)
	}
}

const (
	RegisterS1crFieldHtieShift = 3
	RegisterS1crFieldHtieMask  = 0x8
)

// GetHtie Half transfer interrupt enable
func (r *registerS1crType) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable
func (r *registerS1crType) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldHtieMask)
	}
}

const (
	RegisterS1crFieldTcieShift = 4
	RegisterS1crFieldTcieMask  = 0x10
)

// GetTcie Transfer complete interrupt enable
func (r *registerS1crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable
func (r *registerS1crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldTcieMask)
	}
}

const (
	RegisterS1crFieldPfctrlShift = 5
	RegisterS1crFieldPfctrlMask  = 0x20
)

// GetPfctrl Peripheral flow controller
func (r *registerS1crType) GetPfctrl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldPfctrlMask) != 0
}

// SetPfctrl Peripheral flow controller
func (r *registerS1crType) SetPfctrl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldPfctrlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldPfctrlMask)
	}
}

const (
	RegisterS1crFieldDirShift = 6
	RegisterS1crFieldDirMask  = 0xc0
)

// GetDir Data transfer direction
func (r *registerS1crType) GetDir() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldDirMask) >> RegisterS1crFieldDirShift)
}

// SetDir Data transfer direction
func (r *registerS1crType) SetDir(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldDirMask)|(uint32(value)<<RegisterS1crFieldDirShift))
}

const (
	RegisterS1crFieldCircShift = 8
	RegisterS1crFieldCircMask  = 0x100
)

// GetCirc Circular mode
func (r *registerS1crType) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldCircMask) != 0
}

// SetCirc Circular mode
func (r *registerS1crType) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldCircMask)
	}
}

const (
	RegisterS1crFieldPincShift = 9
	RegisterS1crFieldPincMask  = 0x200
)

// GetPinc Peripheral increment mode
func (r *registerS1crType) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldPincMask) != 0
}

// SetPinc Peripheral increment mode
func (r *registerS1crType) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldPincMask)
	}
}

const (
	RegisterS1crFieldMincShift = 10
	RegisterS1crFieldMincMask  = 0x400
)

// GetMinc Memory increment mode
func (r *registerS1crType) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldMincMask) != 0
}

// SetMinc Memory increment mode
func (r *registerS1crType) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldMincMask)
	}
}

const (
	RegisterS1crFieldPsizeShift = 11
	RegisterS1crFieldPsizeMask  = 0x1800
)

// GetPsize Peripheral data size
func (r *registerS1crType) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldPsizeMask) >> RegisterS1crFieldPsizeShift)
}

// SetPsize Peripheral data size
func (r *registerS1crType) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldPsizeMask)|(uint32(value)<<RegisterS1crFieldPsizeShift))
}

const (
	RegisterS1crFieldMsizeShift = 13
	RegisterS1crFieldMsizeMask  = 0x6000
)

// GetMsize Memory data size
func (r *registerS1crType) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldMsizeMask) >> RegisterS1crFieldMsizeShift)
}

// SetMsize Memory data size
func (r *registerS1crType) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldMsizeMask)|(uint32(value)<<RegisterS1crFieldMsizeShift))
}

const (
	RegisterS1crFieldPincosShift = 15
	RegisterS1crFieldPincosMask  = 0x8000
)

// GetPincos Peripheral increment offset size
func (r *registerS1crType) GetPincos() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldPincosMask) != 0
}

// SetPincos Peripheral increment offset size
func (r *registerS1crType) SetPincos(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldPincosMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldPincosMask)
	}
}

const (
	RegisterS1crFieldPlShift = 16
	RegisterS1crFieldPlMask  = 0x30000
)

// GetPl Priority level
func (r *registerS1crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldPlMask) >> RegisterS1crFieldPlShift)
}

// SetPl Priority level
func (r *registerS1crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldPlMask)|(uint32(value)<<RegisterS1crFieldPlShift))
}

const (
	RegisterS1crFieldDbmShift = 18
	RegisterS1crFieldDbmMask  = 0x40000
)

// GetDbm Double buffer mode
func (r *registerS1crType) GetDbm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldDbmMask) != 0
}

// SetDbm Double buffer mode
func (r *registerS1crType) SetDbm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldDbmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldDbmMask)
	}
}

const (
	RegisterS1crFieldCtShift = 19
	RegisterS1crFieldCtMask  = 0x80000
)

// GetCt Current target (only in double buffer mode)
func (r *registerS1crType) GetCt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldCtMask) != 0
}

// SetCt Current target (only in double buffer mode)
func (r *registerS1crType) SetCt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldCtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldCtMask)
	}
}

const (
	RegisterS1crFieldAckShift = 20
	RegisterS1crFieldAckMask  = 0x100000
)

// GetAck ACK
func (r *registerS1crType) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldAckMask) != 0
}

// SetAck ACK
func (r *registerS1crType) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1crFieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldAckMask)
	}
}

const (
	RegisterS1crFieldPburstShift = 21
	RegisterS1crFieldPburstMask  = 0x600000
)

// GetPburst Peripheral burst transfer configuration
func (r *registerS1crType) GetPburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldPburstMask) >> RegisterS1crFieldPburstShift)
}

// SetPburst Peripheral burst transfer configuration
func (r *registerS1crType) SetPburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldPburstMask)|(uint32(value)<<RegisterS1crFieldPburstShift))
}

const (
	RegisterS1crFieldMburstShift = 23
	RegisterS1crFieldMburstMask  = 0x1800000
)

// GetMburst Memory burst transfer configuration
func (r *registerS1crType) GetMburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS1crFieldMburstMask) >> RegisterS1crFieldMburstShift)
}

// SetMburst Memory burst transfer configuration
func (r *registerS1crType) SetMburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS1crFieldMburstMask)|(uint32(value)<<RegisterS1crFieldMburstShift))
}

// registerS1ndtrType stream x number of data register
type registerS1ndtrType uint32

const (
	RegisterS1ndtrFieldNdtShift = 0
	RegisterS1ndtrFieldNdtMask  = 0xffff
)

// GetNdt Number of data items to transfer
func (r *registerS1ndtrType) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterS1ndtrFieldNdtMask) >> RegisterS1ndtrFieldNdtShift)
}

// SetNdt Number of data items to transfer
func (r *registerS1ndtrType) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS1ndtrFieldNdtMask)|(uint32(value)<<RegisterS1ndtrFieldNdtShift))
}

// registerS1parType stream x peripheral address register
type registerS1parType uint32

const (
	RegisterS1parFieldPaShift = 0
	RegisterS1parFieldPaMask  = 0xffffffff
)

// GetPa Peripheral address
func (r *registerS1parType) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS1parFieldPaMask) >> RegisterS1parFieldPaShift)
}

// SetPa Peripheral address
func (r *registerS1parType) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS1parFieldPaMask)|(uint32(value)<<RegisterS1parFieldPaShift))
}

// registerS1m0arType stream x memory 0 address register
type registerS1m0arType uint32

const (
	RegisterS1m0arFieldM0aShift = 0
	RegisterS1m0arFieldM0aMask  = 0xffffffff
)

// GetM0a Memory 0 address
func (r *registerS1m0arType) GetM0a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS1m0arFieldM0aMask) >> RegisterS1m0arFieldM0aShift)
}

// SetM0a Memory 0 address
func (r *registerS1m0arType) SetM0a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS1m0arFieldM0aMask)|(uint32(value)<<RegisterS1m0arFieldM0aShift))
}

// registerS1m1arType stream x memory 1 address register
type registerS1m1arType uint32

const (
	RegisterS1m1arFieldM1aShift = 0
	RegisterS1m1arFieldM1aMask  = 0xffffffff
)

// GetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS1m1arType) GetM1a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS1m1arFieldM1aMask) >> RegisterS1m1arFieldM1aShift)
}

// SetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS1m1arType) SetM1a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS1m1arFieldM1aMask)|(uint32(value)<<RegisterS1m1arFieldM1aShift))
}

// registerS1fcrType stream x FIFO control register
type registerS1fcrType uint32

const (
	RegisterS1fcrFieldFthShift = 0
	RegisterS1fcrFieldFthMask  = 0x3
)

// GetFth FIFO threshold selection
func (r *registerS1fcrType) GetFth() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS1fcrFieldFthMask) >> RegisterS1fcrFieldFthShift)
}

// SetFth FIFO threshold selection
func (r *registerS1fcrType) SetFth(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS1fcrFieldFthMask)|(uint32(value)<<RegisterS1fcrFieldFthShift))
}

const (
	RegisterS1fcrFieldDmdisShift = 2
	RegisterS1fcrFieldDmdisMask  = 0x4
)

// GetDmdis Direct mode disable
func (r *registerS1fcrType) GetDmdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1fcrFieldDmdisMask) != 0
}

// SetDmdis Direct mode disable
func (r *registerS1fcrType) SetDmdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1fcrFieldDmdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1fcrFieldDmdisMask)
	}
}

const (
	RegisterS1fcrFieldFsShift = 3
	RegisterS1fcrFieldFsMask  = 0x38
)

// GetFs FIFO status
func (r *registerS1fcrType) GetFs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS1fcrFieldFsMask) >> RegisterS1fcrFieldFsShift)
}

const (
	RegisterS1fcrFieldFeieShift = 7
	RegisterS1fcrFieldFeieMask  = 0x80
)

// GetFeie FIFO error interrupt enable
func (r *registerS1fcrType) GetFeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS1fcrFieldFeieMask) != 0
}

// SetFeie FIFO error interrupt enable
func (r *registerS1fcrType) SetFeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS1fcrFieldFeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS1fcrFieldFeieMask)
	}
}

// registerS2crType stream x configuration register
type registerS2crType uint32

const (
	RegisterS2crFieldEnShift = 0
	RegisterS2crFieldEnMask  = 0x1
)

// GetEn Stream enable / flag stream ready when read low
func (r *registerS2crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldEnMask) != 0
}

// SetEn Stream enable / flag stream ready when read low
func (r *registerS2crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldEnMask)
	}
}

const (
	RegisterS2crFieldDmeieShift = 1
	RegisterS2crFieldDmeieMask  = 0x2
)

// GetDmeie Direct mode error interrupt enable
func (r *registerS2crType) GetDmeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldDmeieMask) != 0
}

// SetDmeie Direct mode error interrupt enable
func (r *registerS2crType) SetDmeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldDmeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldDmeieMask)
	}
}

const (
	RegisterS2crFieldTeieShift = 2
	RegisterS2crFieldTeieMask  = 0x4
)

// GetTeie Transfer error interrupt enable
func (r *registerS2crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable
func (r *registerS2crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldTeieMask)
	}
}

const (
	RegisterS2crFieldHtieShift = 3
	RegisterS2crFieldHtieMask  = 0x8
)

// GetHtie Half transfer interrupt enable
func (r *registerS2crType) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable
func (r *registerS2crType) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldHtieMask)
	}
}

const (
	RegisterS2crFieldTcieShift = 4
	RegisterS2crFieldTcieMask  = 0x10
)

// GetTcie Transfer complete interrupt enable
func (r *registerS2crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable
func (r *registerS2crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldTcieMask)
	}
}

const (
	RegisterS2crFieldPfctrlShift = 5
	RegisterS2crFieldPfctrlMask  = 0x20
)

// GetPfctrl Peripheral flow controller
func (r *registerS2crType) GetPfctrl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldPfctrlMask) != 0
}

// SetPfctrl Peripheral flow controller
func (r *registerS2crType) SetPfctrl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldPfctrlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldPfctrlMask)
	}
}

const (
	RegisterS2crFieldDirShift = 6
	RegisterS2crFieldDirMask  = 0xc0
)

// GetDir Data transfer direction
func (r *registerS2crType) GetDir() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldDirMask) >> RegisterS2crFieldDirShift)
}

// SetDir Data transfer direction
func (r *registerS2crType) SetDir(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldDirMask)|(uint32(value)<<RegisterS2crFieldDirShift))
}

const (
	RegisterS2crFieldCircShift = 8
	RegisterS2crFieldCircMask  = 0x100
)

// GetCirc Circular mode
func (r *registerS2crType) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldCircMask) != 0
}

// SetCirc Circular mode
func (r *registerS2crType) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldCircMask)
	}
}

const (
	RegisterS2crFieldPincShift = 9
	RegisterS2crFieldPincMask  = 0x200
)

// GetPinc Peripheral increment mode
func (r *registerS2crType) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldPincMask) != 0
}

// SetPinc Peripheral increment mode
func (r *registerS2crType) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldPincMask)
	}
}

const (
	RegisterS2crFieldMincShift = 10
	RegisterS2crFieldMincMask  = 0x400
)

// GetMinc Memory increment mode
func (r *registerS2crType) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldMincMask) != 0
}

// SetMinc Memory increment mode
func (r *registerS2crType) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldMincMask)
	}
}

const (
	RegisterS2crFieldPsizeShift = 11
	RegisterS2crFieldPsizeMask  = 0x1800
)

// GetPsize Peripheral data size
func (r *registerS2crType) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldPsizeMask) >> RegisterS2crFieldPsizeShift)
}

// SetPsize Peripheral data size
func (r *registerS2crType) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldPsizeMask)|(uint32(value)<<RegisterS2crFieldPsizeShift))
}

const (
	RegisterS2crFieldMsizeShift = 13
	RegisterS2crFieldMsizeMask  = 0x6000
)

// GetMsize Memory data size
func (r *registerS2crType) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldMsizeMask) >> RegisterS2crFieldMsizeShift)
}

// SetMsize Memory data size
func (r *registerS2crType) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldMsizeMask)|(uint32(value)<<RegisterS2crFieldMsizeShift))
}

const (
	RegisterS2crFieldPincosShift = 15
	RegisterS2crFieldPincosMask  = 0x8000
)

// GetPincos Peripheral increment offset size
func (r *registerS2crType) GetPincos() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldPincosMask) != 0
}

// SetPincos Peripheral increment offset size
func (r *registerS2crType) SetPincos(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldPincosMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldPincosMask)
	}
}

const (
	RegisterS2crFieldPlShift = 16
	RegisterS2crFieldPlMask  = 0x30000
)

// GetPl Priority level
func (r *registerS2crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldPlMask) >> RegisterS2crFieldPlShift)
}

// SetPl Priority level
func (r *registerS2crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldPlMask)|(uint32(value)<<RegisterS2crFieldPlShift))
}

const (
	RegisterS2crFieldDbmShift = 18
	RegisterS2crFieldDbmMask  = 0x40000
)

// GetDbm Double buffer mode
func (r *registerS2crType) GetDbm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldDbmMask) != 0
}

// SetDbm Double buffer mode
func (r *registerS2crType) SetDbm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldDbmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldDbmMask)
	}
}

const (
	RegisterS2crFieldCtShift = 19
	RegisterS2crFieldCtMask  = 0x80000
)

// GetCt Current target (only in double buffer mode)
func (r *registerS2crType) GetCt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldCtMask) != 0
}

// SetCt Current target (only in double buffer mode)
func (r *registerS2crType) SetCt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldCtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldCtMask)
	}
}

const (
	RegisterS2crFieldAckShift = 20
	RegisterS2crFieldAckMask  = 0x100000
)

// GetAck ACK
func (r *registerS2crType) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldAckMask) != 0
}

// SetAck ACK
func (r *registerS2crType) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2crFieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldAckMask)
	}
}

const (
	RegisterS2crFieldPburstShift = 21
	RegisterS2crFieldPburstMask  = 0x600000
)

// GetPburst Peripheral burst transfer configuration
func (r *registerS2crType) GetPburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldPburstMask) >> RegisterS2crFieldPburstShift)
}

// SetPburst Peripheral burst transfer configuration
func (r *registerS2crType) SetPburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldPburstMask)|(uint32(value)<<RegisterS2crFieldPburstShift))
}

const (
	RegisterS2crFieldMburstShift = 23
	RegisterS2crFieldMburstMask  = 0x1800000
)

// GetMburst Memory burst transfer configuration
func (r *registerS2crType) GetMburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS2crFieldMburstMask) >> RegisterS2crFieldMburstShift)
}

// SetMburst Memory burst transfer configuration
func (r *registerS2crType) SetMburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS2crFieldMburstMask)|(uint32(value)<<RegisterS2crFieldMburstShift))
}

// registerS2ndtrType stream x number of data register
type registerS2ndtrType uint32

const (
	RegisterS2ndtrFieldNdtShift = 0
	RegisterS2ndtrFieldNdtMask  = 0xffff
)

// GetNdt Number of data items to transfer
func (r *registerS2ndtrType) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterS2ndtrFieldNdtMask) >> RegisterS2ndtrFieldNdtShift)
}

// SetNdt Number of data items to transfer
func (r *registerS2ndtrType) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS2ndtrFieldNdtMask)|(uint32(value)<<RegisterS2ndtrFieldNdtShift))
}

// registerS2parType stream x peripheral address register
type registerS2parType uint32

const (
	RegisterS2parFieldPaShift = 0
	RegisterS2parFieldPaMask  = 0xffffffff
)

// GetPa Peripheral address
func (r *registerS2parType) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS2parFieldPaMask) >> RegisterS2parFieldPaShift)
}

// SetPa Peripheral address
func (r *registerS2parType) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS2parFieldPaMask)|(uint32(value)<<RegisterS2parFieldPaShift))
}

// registerS2m0arType stream x memory 0 address register
type registerS2m0arType uint32

const (
	RegisterS2m0arFieldM0aShift = 0
	RegisterS2m0arFieldM0aMask  = 0xffffffff
)

// GetM0a Memory 0 address
func (r *registerS2m0arType) GetM0a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS2m0arFieldM0aMask) >> RegisterS2m0arFieldM0aShift)
}

// SetM0a Memory 0 address
func (r *registerS2m0arType) SetM0a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS2m0arFieldM0aMask)|(uint32(value)<<RegisterS2m0arFieldM0aShift))
}

// registerS2m1arType stream x memory 1 address register
type registerS2m1arType uint32

const (
	RegisterS2m1arFieldM1aShift = 0
	RegisterS2m1arFieldM1aMask  = 0xffffffff
)

// GetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS2m1arType) GetM1a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS2m1arFieldM1aMask) >> RegisterS2m1arFieldM1aShift)
}

// SetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS2m1arType) SetM1a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS2m1arFieldM1aMask)|(uint32(value)<<RegisterS2m1arFieldM1aShift))
}

// registerS2fcrType stream x FIFO control register
type registerS2fcrType uint32

const (
	RegisterS2fcrFieldFthShift = 0
	RegisterS2fcrFieldFthMask  = 0x3
)

// GetFth FIFO threshold selection
func (r *registerS2fcrType) GetFth() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS2fcrFieldFthMask) >> RegisterS2fcrFieldFthShift)
}

// SetFth FIFO threshold selection
func (r *registerS2fcrType) SetFth(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS2fcrFieldFthMask)|(uint32(value)<<RegisterS2fcrFieldFthShift))
}

const (
	RegisterS2fcrFieldDmdisShift = 2
	RegisterS2fcrFieldDmdisMask  = 0x4
)

// GetDmdis Direct mode disable
func (r *registerS2fcrType) GetDmdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2fcrFieldDmdisMask) != 0
}

// SetDmdis Direct mode disable
func (r *registerS2fcrType) SetDmdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2fcrFieldDmdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2fcrFieldDmdisMask)
	}
}

const (
	RegisterS2fcrFieldFsShift = 3
	RegisterS2fcrFieldFsMask  = 0x38
)

// GetFs FIFO status
func (r *registerS2fcrType) GetFs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS2fcrFieldFsMask) >> RegisterS2fcrFieldFsShift)
}

const (
	RegisterS2fcrFieldFeieShift = 7
	RegisterS2fcrFieldFeieMask  = 0x80
)

// GetFeie FIFO error interrupt enable
func (r *registerS2fcrType) GetFeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS2fcrFieldFeieMask) != 0
}

// SetFeie FIFO error interrupt enable
func (r *registerS2fcrType) SetFeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS2fcrFieldFeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS2fcrFieldFeieMask)
	}
}

// registerS3crType stream x configuration register
type registerS3crType uint32

const (
	RegisterS3crFieldEnShift = 0
	RegisterS3crFieldEnMask  = 0x1
)

// GetEn Stream enable / flag stream ready when read low
func (r *registerS3crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldEnMask) != 0
}

// SetEn Stream enable / flag stream ready when read low
func (r *registerS3crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldEnMask)
	}
}

const (
	RegisterS3crFieldDmeieShift = 1
	RegisterS3crFieldDmeieMask  = 0x2
)

// GetDmeie Direct mode error interrupt enable
func (r *registerS3crType) GetDmeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldDmeieMask) != 0
}

// SetDmeie Direct mode error interrupt enable
func (r *registerS3crType) SetDmeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldDmeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldDmeieMask)
	}
}

const (
	RegisterS3crFieldTeieShift = 2
	RegisterS3crFieldTeieMask  = 0x4
)

// GetTeie Transfer error interrupt enable
func (r *registerS3crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable
func (r *registerS3crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldTeieMask)
	}
}

const (
	RegisterS3crFieldHtieShift = 3
	RegisterS3crFieldHtieMask  = 0x8
)

// GetHtie Half transfer interrupt enable
func (r *registerS3crType) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable
func (r *registerS3crType) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldHtieMask)
	}
}

const (
	RegisterS3crFieldTcieShift = 4
	RegisterS3crFieldTcieMask  = 0x10
)

// GetTcie Transfer complete interrupt enable
func (r *registerS3crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable
func (r *registerS3crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldTcieMask)
	}
}

const (
	RegisterS3crFieldPfctrlShift = 5
	RegisterS3crFieldPfctrlMask  = 0x20
)

// GetPfctrl Peripheral flow controller
func (r *registerS3crType) GetPfctrl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldPfctrlMask) != 0
}

// SetPfctrl Peripheral flow controller
func (r *registerS3crType) SetPfctrl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldPfctrlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldPfctrlMask)
	}
}

const (
	RegisterS3crFieldDirShift = 6
	RegisterS3crFieldDirMask  = 0xc0
)

// GetDir Data transfer direction
func (r *registerS3crType) GetDir() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldDirMask) >> RegisterS3crFieldDirShift)
}

// SetDir Data transfer direction
func (r *registerS3crType) SetDir(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldDirMask)|(uint32(value)<<RegisterS3crFieldDirShift))
}

const (
	RegisterS3crFieldCircShift = 8
	RegisterS3crFieldCircMask  = 0x100
)

// GetCirc Circular mode
func (r *registerS3crType) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldCircMask) != 0
}

// SetCirc Circular mode
func (r *registerS3crType) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldCircMask)
	}
}

const (
	RegisterS3crFieldPincShift = 9
	RegisterS3crFieldPincMask  = 0x200
)

// GetPinc Peripheral increment mode
func (r *registerS3crType) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldPincMask) != 0
}

// SetPinc Peripheral increment mode
func (r *registerS3crType) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldPincMask)
	}
}

const (
	RegisterS3crFieldMincShift = 10
	RegisterS3crFieldMincMask  = 0x400
)

// GetMinc Memory increment mode
func (r *registerS3crType) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldMincMask) != 0
}

// SetMinc Memory increment mode
func (r *registerS3crType) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldMincMask)
	}
}

const (
	RegisterS3crFieldPsizeShift = 11
	RegisterS3crFieldPsizeMask  = 0x1800
)

// GetPsize Peripheral data size
func (r *registerS3crType) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldPsizeMask) >> RegisterS3crFieldPsizeShift)
}

// SetPsize Peripheral data size
func (r *registerS3crType) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldPsizeMask)|(uint32(value)<<RegisterS3crFieldPsizeShift))
}

const (
	RegisterS3crFieldMsizeShift = 13
	RegisterS3crFieldMsizeMask  = 0x6000
)

// GetMsize Memory data size
func (r *registerS3crType) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldMsizeMask) >> RegisterS3crFieldMsizeShift)
}

// SetMsize Memory data size
func (r *registerS3crType) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldMsizeMask)|(uint32(value)<<RegisterS3crFieldMsizeShift))
}

const (
	RegisterS3crFieldPincosShift = 15
	RegisterS3crFieldPincosMask  = 0x8000
)

// GetPincos Peripheral increment offset size
func (r *registerS3crType) GetPincos() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldPincosMask) != 0
}

// SetPincos Peripheral increment offset size
func (r *registerS3crType) SetPincos(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldPincosMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldPincosMask)
	}
}

const (
	RegisterS3crFieldPlShift = 16
	RegisterS3crFieldPlMask  = 0x30000
)

// GetPl Priority level
func (r *registerS3crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldPlMask) >> RegisterS3crFieldPlShift)
}

// SetPl Priority level
func (r *registerS3crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldPlMask)|(uint32(value)<<RegisterS3crFieldPlShift))
}

const (
	RegisterS3crFieldDbmShift = 18
	RegisterS3crFieldDbmMask  = 0x40000
)

// GetDbm Double buffer mode
func (r *registerS3crType) GetDbm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldDbmMask) != 0
}

// SetDbm Double buffer mode
func (r *registerS3crType) SetDbm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldDbmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldDbmMask)
	}
}

const (
	RegisterS3crFieldCtShift = 19
	RegisterS3crFieldCtMask  = 0x80000
)

// GetCt Current target (only in double buffer mode)
func (r *registerS3crType) GetCt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldCtMask) != 0
}

// SetCt Current target (only in double buffer mode)
func (r *registerS3crType) SetCt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldCtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldCtMask)
	}
}

const (
	RegisterS3crFieldAckShift = 20
	RegisterS3crFieldAckMask  = 0x100000
)

// GetAck ACK
func (r *registerS3crType) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldAckMask) != 0
}

// SetAck ACK
func (r *registerS3crType) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3crFieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldAckMask)
	}
}

const (
	RegisterS3crFieldPburstShift = 21
	RegisterS3crFieldPburstMask  = 0x600000
)

// GetPburst Peripheral burst transfer configuration
func (r *registerS3crType) GetPburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldPburstMask) >> RegisterS3crFieldPburstShift)
}

// SetPburst Peripheral burst transfer configuration
func (r *registerS3crType) SetPburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldPburstMask)|(uint32(value)<<RegisterS3crFieldPburstShift))
}

const (
	RegisterS3crFieldMburstShift = 23
	RegisterS3crFieldMburstMask  = 0x1800000
)

// GetMburst Memory burst transfer configuration
func (r *registerS3crType) GetMburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS3crFieldMburstMask) >> RegisterS3crFieldMburstShift)
}

// SetMburst Memory burst transfer configuration
func (r *registerS3crType) SetMburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS3crFieldMburstMask)|(uint32(value)<<RegisterS3crFieldMburstShift))
}

// registerS3ndtrType stream x number of data register
type registerS3ndtrType uint32

const (
	RegisterS3ndtrFieldNdtShift = 0
	RegisterS3ndtrFieldNdtMask  = 0xffff
)

// GetNdt Number of data items to transfer
func (r *registerS3ndtrType) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterS3ndtrFieldNdtMask) >> RegisterS3ndtrFieldNdtShift)
}

// SetNdt Number of data items to transfer
func (r *registerS3ndtrType) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS3ndtrFieldNdtMask)|(uint32(value)<<RegisterS3ndtrFieldNdtShift))
}

// registerS3parType stream x peripheral address register
type registerS3parType uint32

const (
	RegisterS3parFieldPaShift = 0
	RegisterS3parFieldPaMask  = 0xffffffff
)

// GetPa Peripheral address
func (r *registerS3parType) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS3parFieldPaMask) >> RegisterS3parFieldPaShift)
}

// SetPa Peripheral address
func (r *registerS3parType) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS3parFieldPaMask)|(uint32(value)<<RegisterS3parFieldPaShift))
}

// registerS3m0arType stream x memory 0 address register
type registerS3m0arType uint32

const (
	RegisterS3m0arFieldM0aShift = 0
	RegisterS3m0arFieldM0aMask  = 0xffffffff
)

// GetM0a Memory 0 address
func (r *registerS3m0arType) GetM0a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS3m0arFieldM0aMask) >> RegisterS3m0arFieldM0aShift)
}

// SetM0a Memory 0 address
func (r *registerS3m0arType) SetM0a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS3m0arFieldM0aMask)|(uint32(value)<<RegisterS3m0arFieldM0aShift))
}

// registerS3m1arType stream x memory 1 address register
type registerS3m1arType uint32

const (
	RegisterS3m1arFieldM1aShift = 0
	RegisterS3m1arFieldM1aMask  = 0xffffffff
)

// GetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS3m1arType) GetM1a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS3m1arFieldM1aMask) >> RegisterS3m1arFieldM1aShift)
}

// SetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS3m1arType) SetM1a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS3m1arFieldM1aMask)|(uint32(value)<<RegisterS3m1arFieldM1aShift))
}

// registerS3fcrType stream x FIFO control register
type registerS3fcrType uint32

const (
	RegisterS3fcrFieldFthShift = 0
	RegisterS3fcrFieldFthMask  = 0x3
)

// GetFth FIFO threshold selection
func (r *registerS3fcrType) GetFth() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS3fcrFieldFthMask) >> RegisterS3fcrFieldFthShift)
}

// SetFth FIFO threshold selection
func (r *registerS3fcrType) SetFth(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS3fcrFieldFthMask)|(uint32(value)<<RegisterS3fcrFieldFthShift))
}

const (
	RegisterS3fcrFieldDmdisShift = 2
	RegisterS3fcrFieldDmdisMask  = 0x4
)

// GetDmdis Direct mode disable
func (r *registerS3fcrType) GetDmdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3fcrFieldDmdisMask) != 0
}

// SetDmdis Direct mode disable
func (r *registerS3fcrType) SetDmdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3fcrFieldDmdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3fcrFieldDmdisMask)
	}
}

const (
	RegisterS3fcrFieldFsShift = 3
	RegisterS3fcrFieldFsMask  = 0x38
)

// GetFs FIFO status
func (r *registerS3fcrType) GetFs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS3fcrFieldFsMask) >> RegisterS3fcrFieldFsShift)
}

const (
	RegisterS3fcrFieldFeieShift = 7
	RegisterS3fcrFieldFeieMask  = 0x80
)

// GetFeie FIFO error interrupt enable
func (r *registerS3fcrType) GetFeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS3fcrFieldFeieMask) != 0
}

// SetFeie FIFO error interrupt enable
func (r *registerS3fcrType) SetFeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS3fcrFieldFeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS3fcrFieldFeieMask)
	}
}

// registerS4crType stream x configuration register
type registerS4crType uint32

const (
	RegisterS4crFieldEnShift = 0
	RegisterS4crFieldEnMask  = 0x1
)

// GetEn Stream enable / flag stream ready when read low
func (r *registerS4crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldEnMask) != 0
}

// SetEn Stream enable / flag stream ready when read low
func (r *registerS4crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldEnMask)
	}
}

const (
	RegisterS4crFieldDmeieShift = 1
	RegisterS4crFieldDmeieMask  = 0x2
)

// GetDmeie Direct mode error interrupt enable
func (r *registerS4crType) GetDmeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldDmeieMask) != 0
}

// SetDmeie Direct mode error interrupt enable
func (r *registerS4crType) SetDmeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldDmeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldDmeieMask)
	}
}

const (
	RegisterS4crFieldTeieShift = 2
	RegisterS4crFieldTeieMask  = 0x4
)

// GetTeie Transfer error interrupt enable
func (r *registerS4crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable
func (r *registerS4crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldTeieMask)
	}
}

const (
	RegisterS4crFieldHtieShift = 3
	RegisterS4crFieldHtieMask  = 0x8
)

// GetHtie Half transfer interrupt enable
func (r *registerS4crType) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable
func (r *registerS4crType) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldHtieMask)
	}
}

const (
	RegisterS4crFieldTcieShift = 4
	RegisterS4crFieldTcieMask  = 0x10
)

// GetTcie Transfer complete interrupt enable
func (r *registerS4crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable
func (r *registerS4crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldTcieMask)
	}
}

const (
	RegisterS4crFieldPfctrlShift = 5
	RegisterS4crFieldPfctrlMask  = 0x20
)

// GetPfctrl Peripheral flow controller
func (r *registerS4crType) GetPfctrl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldPfctrlMask) != 0
}

// SetPfctrl Peripheral flow controller
func (r *registerS4crType) SetPfctrl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldPfctrlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldPfctrlMask)
	}
}

const (
	RegisterS4crFieldDirShift = 6
	RegisterS4crFieldDirMask  = 0xc0
)

// GetDir Data transfer direction
func (r *registerS4crType) GetDir() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldDirMask) >> RegisterS4crFieldDirShift)
}

// SetDir Data transfer direction
func (r *registerS4crType) SetDir(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldDirMask)|(uint32(value)<<RegisterS4crFieldDirShift))
}

const (
	RegisterS4crFieldCircShift = 8
	RegisterS4crFieldCircMask  = 0x100
)

// GetCirc Circular mode
func (r *registerS4crType) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldCircMask) != 0
}

// SetCirc Circular mode
func (r *registerS4crType) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldCircMask)
	}
}

const (
	RegisterS4crFieldPincShift = 9
	RegisterS4crFieldPincMask  = 0x200
)

// GetPinc Peripheral increment mode
func (r *registerS4crType) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldPincMask) != 0
}

// SetPinc Peripheral increment mode
func (r *registerS4crType) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldPincMask)
	}
}

const (
	RegisterS4crFieldMincShift = 10
	RegisterS4crFieldMincMask  = 0x400
)

// GetMinc Memory increment mode
func (r *registerS4crType) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldMincMask) != 0
}

// SetMinc Memory increment mode
func (r *registerS4crType) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldMincMask)
	}
}

const (
	RegisterS4crFieldPsizeShift = 11
	RegisterS4crFieldPsizeMask  = 0x1800
)

// GetPsize Peripheral data size
func (r *registerS4crType) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldPsizeMask) >> RegisterS4crFieldPsizeShift)
}

// SetPsize Peripheral data size
func (r *registerS4crType) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldPsizeMask)|(uint32(value)<<RegisterS4crFieldPsizeShift))
}

const (
	RegisterS4crFieldMsizeShift = 13
	RegisterS4crFieldMsizeMask  = 0x6000
)

// GetMsize Memory data size
func (r *registerS4crType) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldMsizeMask) >> RegisterS4crFieldMsizeShift)
}

// SetMsize Memory data size
func (r *registerS4crType) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldMsizeMask)|(uint32(value)<<RegisterS4crFieldMsizeShift))
}

const (
	RegisterS4crFieldPincosShift = 15
	RegisterS4crFieldPincosMask  = 0x8000
)

// GetPincos Peripheral increment offset size
func (r *registerS4crType) GetPincos() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldPincosMask) != 0
}

// SetPincos Peripheral increment offset size
func (r *registerS4crType) SetPincos(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldPincosMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldPincosMask)
	}
}

const (
	RegisterS4crFieldPlShift = 16
	RegisterS4crFieldPlMask  = 0x30000
)

// GetPl Priority level
func (r *registerS4crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldPlMask) >> RegisterS4crFieldPlShift)
}

// SetPl Priority level
func (r *registerS4crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldPlMask)|(uint32(value)<<RegisterS4crFieldPlShift))
}

const (
	RegisterS4crFieldDbmShift = 18
	RegisterS4crFieldDbmMask  = 0x40000
)

// GetDbm Double buffer mode
func (r *registerS4crType) GetDbm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldDbmMask) != 0
}

// SetDbm Double buffer mode
func (r *registerS4crType) SetDbm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldDbmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldDbmMask)
	}
}

const (
	RegisterS4crFieldCtShift = 19
	RegisterS4crFieldCtMask  = 0x80000
)

// GetCt Current target (only in double buffer mode)
func (r *registerS4crType) GetCt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldCtMask) != 0
}

// SetCt Current target (only in double buffer mode)
func (r *registerS4crType) SetCt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldCtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldCtMask)
	}
}

const (
	RegisterS4crFieldAckShift = 20
	RegisterS4crFieldAckMask  = 0x100000
)

// GetAck ACK
func (r *registerS4crType) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldAckMask) != 0
}

// SetAck ACK
func (r *registerS4crType) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4crFieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldAckMask)
	}
}

const (
	RegisterS4crFieldPburstShift = 21
	RegisterS4crFieldPburstMask  = 0x600000
)

// GetPburst Peripheral burst transfer configuration
func (r *registerS4crType) GetPburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldPburstMask) >> RegisterS4crFieldPburstShift)
}

// SetPburst Peripheral burst transfer configuration
func (r *registerS4crType) SetPburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldPburstMask)|(uint32(value)<<RegisterS4crFieldPburstShift))
}

const (
	RegisterS4crFieldMburstShift = 23
	RegisterS4crFieldMburstMask  = 0x1800000
)

// GetMburst Memory burst transfer configuration
func (r *registerS4crType) GetMburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS4crFieldMburstMask) >> RegisterS4crFieldMburstShift)
}

// SetMburst Memory burst transfer configuration
func (r *registerS4crType) SetMburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS4crFieldMburstMask)|(uint32(value)<<RegisterS4crFieldMburstShift))
}

// registerS4ndtrType stream x number of data register
type registerS4ndtrType uint32

const (
	RegisterS4ndtrFieldNdtShift = 0
	RegisterS4ndtrFieldNdtMask  = 0xffff
)

// GetNdt Number of data items to transfer
func (r *registerS4ndtrType) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterS4ndtrFieldNdtMask) >> RegisterS4ndtrFieldNdtShift)
}

// SetNdt Number of data items to transfer
func (r *registerS4ndtrType) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS4ndtrFieldNdtMask)|(uint32(value)<<RegisterS4ndtrFieldNdtShift))
}

// registerS4parType stream x peripheral address register
type registerS4parType uint32

const (
	RegisterS4parFieldPaShift = 0
	RegisterS4parFieldPaMask  = 0xffffffff
)

// GetPa Peripheral address
func (r *registerS4parType) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS4parFieldPaMask) >> RegisterS4parFieldPaShift)
}

// SetPa Peripheral address
func (r *registerS4parType) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS4parFieldPaMask)|(uint32(value)<<RegisterS4parFieldPaShift))
}

// registerS4m0arType stream x memory 0 address register
type registerS4m0arType uint32

const (
	RegisterS4m0arFieldM0aShift = 0
	RegisterS4m0arFieldM0aMask  = 0xffffffff
)

// GetM0a Memory 0 address
func (r *registerS4m0arType) GetM0a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS4m0arFieldM0aMask) >> RegisterS4m0arFieldM0aShift)
}

// SetM0a Memory 0 address
func (r *registerS4m0arType) SetM0a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS4m0arFieldM0aMask)|(uint32(value)<<RegisterS4m0arFieldM0aShift))
}

// registerS4m1arType stream x memory 1 address register
type registerS4m1arType uint32

const (
	RegisterS4m1arFieldM1aShift = 0
	RegisterS4m1arFieldM1aMask  = 0xffffffff
)

// GetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS4m1arType) GetM1a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS4m1arFieldM1aMask) >> RegisterS4m1arFieldM1aShift)
}

// SetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS4m1arType) SetM1a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS4m1arFieldM1aMask)|(uint32(value)<<RegisterS4m1arFieldM1aShift))
}

// registerS4fcrType stream x FIFO control register
type registerS4fcrType uint32

const (
	RegisterS4fcrFieldFthShift = 0
	RegisterS4fcrFieldFthMask  = 0x3
)

// GetFth FIFO threshold selection
func (r *registerS4fcrType) GetFth() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS4fcrFieldFthMask) >> RegisterS4fcrFieldFthShift)
}

// SetFth FIFO threshold selection
func (r *registerS4fcrType) SetFth(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS4fcrFieldFthMask)|(uint32(value)<<RegisterS4fcrFieldFthShift))
}

const (
	RegisterS4fcrFieldDmdisShift = 2
	RegisterS4fcrFieldDmdisMask  = 0x4
)

// GetDmdis Direct mode disable
func (r *registerS4fcrType) GetDmdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4fcrFieldDmdisMask) != 0
}

// SetDmdis Direct mode disable
func (r *registerS4fcrType) SetDmdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4fcrFieldDmdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4fcrFieldDmdisMask)
	}
}

const (
	RegisterS4fcrFieldFsShift = 3
	RegisterS4fcrFieldFsMask  = 0x38
)

// GetFs FIFO status
func (r *registerS4fcrType) GetFs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS4fcrFieldFsMask) >> RegisterS4fcrFieldFsShift)
}

const (
	RegisterS4fcrFieldFeieShift = 7
	RegisterS4fcrFieldFeieMask  = 0x80
)

// GetFeie FIFO error interrupt enable
func (r *registerS4fcrType) GetFeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS4fcrFieldFeieMask) != 0
}

// SetFeie FIFO error interrupt enable
func (r *registerS4fcrType) SetFeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS4fcrFieldFeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS4fcrFieldFeieMask)
	}
}

// registerS5crType stream x configuration register
type registerS5crType uint32

const (
	RegisterS5crFieldEnShift = 0
	RegisterS5crFieldEnMask  = 0x1
)

// GetEn Stream enable / flag stream ready when read low
func (r *registerS5crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldEnMask) != 0
}

// SetEn Stream enable / flag stream ready when read low
func (r *registerS5crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldEnMask)
	}
}

const (
	RegisterS5crFieldDmeieShift = 1
	RegisterS5crFieldDmeieMask  = 0x2
)

// GetDmeie Direct mode error interrupt enable
func (r *registerS5crType) GetDmeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldDmeieMask) != 0
}

// SetDmeie Direct mode error interrupt enable
func (r *registerS5crType) SetDmeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldDmeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldDmeieMask)
	}
}

const (
	RegisterS5crFieldTeieShift = 2
	RegisterS5crFieldTeieMask  = 0x4
)

// GetTeie Transfer error interrupt enable
func (r *registerS5crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable
func (r *registerS5crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldTeieMask)
	}
}

const (
	RegisterS5crFieldHtieShift = 3
	RegisterS5crFieldHtieMask  = 0x8
)

// GetHtie Half transfer interrupt enable
func (r *registerS5crType) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable
func (r *registerS5crType) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldHtieMask)
	}
}

const (
	RegisterS5crFieldTcieShift = 4
	RegisterS5crFieldTcieMask  = 0x10
)

// GetTcie Transfer complete interrupt enable
func (r *registerS5crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable
func (r *registerS5crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldTcieMask)
	}
}

const (
	RegisterS5crFieldPfctrlShift = 5
	RegisterS5crFieldPfctrlMask  = 0x20
)

// GetPfctrl Peripheral flow controller
func (r *registerS5crType) GetPfctrl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldPfctrlMask) != 0
}

// SetPfctrl Peripheral flow controller
func (r *registerS5crType) SetPfctrl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldPfctrlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldPfctrlMask)
	}
}

const (
	RegisterS5crFieldDirShift = 6
	RegisterS5crFieldDirMask  = 0xc0
)

// GetDir Data transfer direction
func (r *registerS5crType) GetDir() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldDirMask) >> RegisterS5crFieldDirShift)
}

// SetDir Data transfer direction
func (r *registerS5crType) SetDir(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldDirMask)|(uint32(value)<<RegisterS5crFieldDirShift))
}

const (
	RegisterS5crFieldCircShift = 8
	RegisterS5crFieldCircMask  = 0x100
)

// GetCirc Circular mode
func (r *registerS5crType) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldCircMask) != 0
}

// SetCirc Circular mode
func (r *registerS5crType) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldCircMask)
	}
}

const (
	RegisterS5crFieldPincShift = 9
	RegisterS5crFieldPincMask  = 0x200
)

// GetPinc Peripheral increment mode
func (r *registerS5crType) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldPincMask) != 0
}

// SetPinc Peripheral increment mode
func (r *registerS5crType) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldPincMask)
	}
}

const (
	RegisterS5crFieldMincShift = 10
	RegisterS5crFieldMincMask  = 0x400
)

// GetMinc Memory increment mode
func (r *registerS5crType) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldMincMask) != 0
}

// SetMinc Memory increment mode
func (r *registerS5crType) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldMincMask)
	}
}

const (
	RegisterS5crFieldPsizeShift = 11
	RegisterS5crFieldPsizeMask  = 0x1800
)

// GetPsize Peripheral data size
func (r *registerS5crType) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldPsizeMask) >> RegisterS5crFieldPsizeShift)
}

// SetPsize Peripheral data size
func (r *registerS5crType) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldPsizeMask)|(uint32(value)<<RegisterS5crFieldPsizeShift))
}

const (
	RegisterS5crFieldMsizeShift = 13
	RegisterS5crFieldMsizeMask  = 0x6000
)

// GetMsize Memory data size
func (r *registerS5crType) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldMsizeMask) >> RegisterS5crFieldMsizeShift)
}

// SetMsize Memory data size
func (r *registerS5crType) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldMsizeMask)|(uint32(value)<<RegisterS5crFieldMsizeShift))
}

const (
	RegisterS5crFieldPincosShift = 15
	RegisterS5crFieldPincosMask  = 0x8000
)

// GetPincos Peripheral increment offset size
func (r *registerS5crType) GetPincos() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldPincosMask) != 0
}

// SetPincos Peripheral increment offset size
func (r *registerS5crType) SetPincos(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldPincosMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldPincosMask)
	}
}

const (
	RegisterS5crFieldPlShift = 16
	RegisterS5crFieldPlMask  = 0x30000
)

// GetPl Priority level
func (r *registerS5crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldPlMask) >> RegisterS5crFieldPlShift)
}

// SetPl Priority level
func (r *registerS5crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldPlMask)|(uint32(value)<<RegisterS5crFieldPlShift))
}

const (
	RegisterS5crFieldDbmShift = 18
	RegisterS5crFieldDbmMask  = 0x40000
)

// GetDbm Double buffer mode
func (r *registerS5crType) GetDbm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldDbmMask) != 0
}

// SetDbm Double buffer mode
func (r *registerS5crType) SetDbm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldDbmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldDbmMask)
	}
}

const (
	RegisterS5crFieldCtShift = 19
	RegisterS5crFieldCtMask  = 0x80000
)

// GetCt Current target (only in double buffer mode)
func (r *registerS5crType) GetCt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldCtMask) != 0
}

// SetCt Current target (only in double buffer mode)
func (r *registerS5crType) SetCt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldCtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldCtMask)
	}
}

const (
	RegisterS5crFieldAckShift = 20
	RegisterS5crFieldAckMask  = 0x100000
)

// GetAck ACK
func (r *registerS5crType) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldAckMask) != 0
}

// SetAck ACK
func (r *registerS5crType) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5crFieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldAckMask)
	}
}

const (
	RegisterS5crFieldPburstShift = 21
	RegisterS5crFieldPburstMask  = 0x600000
)

// GetPburst Peripheral burst transfer configuration
func (r *registerS5crType) GetPburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldPburstMask) >> RegisterS5crFieldPburstShift)
}

// SetPburst Peripheral burst transfer configuration
func (r *registerS5crType) SetPburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldPburstMask)|(uint32(value)<<RegisterS5crFieldPburstShift))
}

const (
	RegisterS5crFieldMburstShift = 23
	RegisterS5crFieldMburstMask  = 0x1800000
)

// GetMburst Memory burst transfer configuration
func (r *registerS5crType) GetMburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS5crFieldMburstMask) >> RegisterS5crFieldMburstShift)
}

// SetMburst Memory burst transfer configuration
func (r *registerS5crType) SetMburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS5crFieldMburstMask)|(uint32(value)<<RegisterS5crFieldMburstShift))
}

// registerS5ndtrType stream x number of data register
type registerS5ndtrType uint32

const (
	RegisterS5ndtrFieldNdtShift = 0
	RegisterS5ndtrFieldNdtMask  = 0xffff
)

// GetNdt Number of data items to transfer
func (r *registerS5ndtrType) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterS5ndtrFieldNdtMask) >> RegisterS5ndtrFieldNdtShift)
}

// SetNdt Number of data items to transfer
func (r *registerS5ndtrType) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS5ndtrFieldNdtMask)|(uint32(value)<<RegisterS5ndtrFieldNdtShift))
}

// registerS5parType stream x peripheral address register
type registerS5parType uint32

const (
	RegisterS5parFieldPaShift = 0
	RegisterS5parFieldPaMask  = 0xffffffff
)

// GetPa Peripheral address
func (r *registerS5parType) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS5parFieldPaMask) >> RegisterS5parFieldPaShift)
}

// SetPa Peripheral address
func (r *registerS5parType) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS5parFieldPaMask)|(uint32(value)<<RegisterS5parFieldPaShift))
}

// registerS5m0arType stream x memory 0 address register
type registerS5m0arType uint32

const (
	RegisterS5m0arFieldM0aShift = 0
	RegisterS5m0arFieldM0aMask  = 0xffffffff
)

// GetM0a Memory 0 address
func (r *registerS5m0arType) GetM0a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS5m0arFieldM0aMask) >> RegisterS5m0arFieldM0aShift)
}

// SetM0a Memory 0 address
func (r *registerS5m0arType) SetM0a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS5m0arFieldM0aMask)|(uint32(value)<<RegisterS5m0arFieldM0aShift))
}

// registerS5m1arType stream x memory 1 address register
type registerS5m1arType uint32

const (
	RegisterS5m1arFieldM1aShift = 0
	RegisterS5m1arFieldM1aMask  = 0xffffffff
)

// GetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS5m1arType) GetM1a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS5m1arFieldM1aMask) >> RegisterS5m1arFieldM1aShift)
}

// SetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS5m1arType) SetM1a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS5m1arFieldM1aMask)|(uint32(value)<<RegisterS5m1arFieldM1aShift))
}

// registerS5fcrType stream x FIFO control register
type registerS5fcrType uint32

const (
	RegisterS5fcrFieldFthShift = 0
	RegisterS5fcrFieldFthMask  = 0x3
)

// GetFth FIFO threshold selection
func (r *registerS5fcrType) GetFth() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS5fcrFieldFthMask) >> RegisterS5fcrFieldFthShift)
}

// SetFth FIFO threshold selection
func (r *registerS5fcrType) SetFth(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS5fcrFieldFthMask)|(uint32(value)<<RegisterS5fcrFieldFthShift))
}

const (
	RegisterS5fcrFieldDmdisShift = 2
	RegisterS5fcrFieldDmdisMask  = 0x4
)

// GetDmdis Direct mode disable
func (r *registerS5fcrType) GetDmdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5fcrFieldDmdisMask) != 0
}

// SetDmdis Direct mode disable
func (r *registerS5fcrType) SetDmdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5fcrFieldDmdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5fcrFieldDmdisMask)
	}
}

const (
	RegisterS5fcrFieldFsShift = 3
	RegisterS5fcrFieldFsMask  = 0x38
)

// GetFs FIFO status
func (r *registerS5fcrType) GetFs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS5fcrFieldFsMask) >> RegisterS5fcrFieldFsShift)
}

const (
	RegisterS5fcrFieldFeieShift = 7
	RegisterS5fcrFieldFeieMask  = 0x80
)

// GetFeie FIFO error interrupt enable
func (r *registerS5fcrType) GetFeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS5fcrFieldFeieMask) != 0
}

// SetFeie FIFO error interrupt enable
func (r *registerS5fcrType) SetFeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS5fcrFieldFeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS5fcrFieldFeieMask)
	}
}

// registerS6crType stream x configuration register
type registerS6crType uint32

const (
	RegisterS6crFieldEnShift = 0
	RegisterS6crFieldEnMask  = 0x1
)

// GetEn Stream enable / flag stream ready when read low
func (r *registerS6crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldEnMask) != 0
}

// SetEn Stream enable / flag stream ready when read low
func (r *registerS6crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldEnMask)
	}
}

const (
	RegisterS6crFieldDmeieShift = 1
	RegisterS6crFieldDmeieMask  = 0x2
)

// GetDmeie Direct mode error interrupt enable
func (r *registerS6crType) GetDmeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldDmeieMask) != 0
}

// SetDmeie Direct mode error interrupt enable
func (r *registerS6crType) SetDmeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldDmeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldDmeieMask)
	}
}

const (
	RegisterS6crFieldTeieShift = 2
	RegisterS6crFieldTeieMask  = 0x4
)

// GetTeie Transfer error interrupt enable
func (r *registerS6crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable
func (r *registerS6crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldTeieMask)
	}
}

const (
	RegisterS6crFieldHtieShift = 3
	RegisterS6crFieldHtieMask  = 0x8
)

// GetHtie Half transfer interrupt enable
func (r *registerS6crType) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable
func (r *registerS6crType) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldHtieMask)
	}
}

const (
	RegisterS6crFieldTcieShift = 4
	RegisterS6crFieldTcieMask  = 0x10
)

// GetTcie Transfer complete interrupt enable
func (r *registerS6crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable
func (r *registerS6crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldTcieMask)
	}
}

const (
	RegisterS6crFieldPfctrlShift = 5
	RegisterS6crFieldPfctrlMask  = 0x20
)

// GetPfctrl Peripheral flow controller
func (r *registerS6crType) GetPfctrl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldPfctrlMask) != 0
}

// SetPfctrl Peripheral flow controller
func (r *registerS6crType) SetPfctrl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldPfctrlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldPfctrlMask)
	}
}

const (
	RegisterS6crFieldDirShift = 6
	RegisterS6crFieldDirMask  = 0xc0
)

// GetDir Data transfer direction
func (r *registerS6crType) GetDir() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldDirMask) >> RegisterS6crFieldDirShift)
}

// SetDir Data transfer direction
func (r *registerS6crType) SetDir(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldDirMask)|(uint32(value)<<RegisterS6crFieldDirShift))
}

const (
	RegisterS6crFieldCircShift = 8
	RegisterS6crFieldCircMask  = 0x100
)

// GetCirc Circular mode
func (r *registerS6crType) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldCircMask) != 0
}

// SetCirc Circular mode
func (r *registerS6crType) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldCircMask)
	}
}

const (
	RegisterS6crFieldPincShift = 9
	RegisterS6crFieldPincMask  = 0x200
)

// GetPinc Peripheral increment mode
func (r *registerS6crType) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldPincMask) != 0
}

// SetPinc Peripheral increment mode
func (r *registerS6crType) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldPincMask)
	}
}

const (
	RegisterS6crFieldMincShift = 10
	RegisterS6crFieldMincMask  = 0x400
)

// GetMinc Memory increment mode
func (r *registerS6crType) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldMincMask) != 0
}

// SetMinc Memory increment mode
func (r *registerS6crType) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldMincMask)
	}
}

const (
	RegisterS6crFieldPsizeShift = 11
	RegisterS6crFieldPsizeMask  = 0x1800
)

// GetPsize Peripheral data size
func (r *registerS6crType) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldPsizeMask) >> RegisterS6crFieldPsizeShift)
}

// SetPsize Peripheral data size
func (r *registerS6crType) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldPsizeMask)|(uint32(value)<<RegisterS6crFieldPsizeShift))
}

const (
	RegisterS6crFieldMsizeShift = 13
	RegisterS6crFieldMsizeMask  = 0x6000
)

// GetMsize Memory data size
func (r *registerS6crType) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldMsizeMask) >> RegisterS6crFieldMsizeShift)
}

// SetMsize Memory data size
func (r *registerS6crType) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldMsizeMask)|(uint32(value)<<RegisterS6crFieldMsizeShift))
}

const (
	RegisterS6crFieldPincosShift = 15
	RegisterS6crFieldPincosMask  = 0x8000
)

// GetPincos Peripheral increment offset size
func (r *registerS6crType) GetPincos() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldPincosMask) != 0
}

// SetPincos Peripheral increment offset size
func (r *registerS6crType) SetPincos(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldPincosMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldPincosMask)
	}
}

const (
	RegisterS6crFieldPlShift = 16
	RegisterS6crFieldPlMask  = 0x30000
)

// GetPl Priority level
func (r *registerS6crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldPlMask) >> RegisterS6crFieldPlShift)
}

// SetPl Priority level
func (r *registerS6crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldPlMask)|(uint32(value)<<RegisterS6crFieldPlShift))
}

const (
	RegisterS6crFieldDbmShift = 18
	RegisterS6crFieldDbmMask  = 0x40000
)

// GetDbm Double buffer mode
func (r *registerS6crType) GetDbm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldDbmMask) != 0
}

// SetDbm Double buffer mode
func (r *registerS6crType) SetDbm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldDbmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldDbmMask)
	}
}

const (
	RegisterS6crFieldCtShift = 19
	RegisterS6crFieldCtMask  = 0x80000
)

// GetCt Current target (only in double buffer mode)
func (r *registerS6crType) GetCt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldCtMask) != 0
}

// SetCt Current target (only in double buffer mode)
func (r *registerS6crType) SetCt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldCtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldCtMask)
	}
}

const (
	RegisterS6crFieldAckShift = 20
	RegisterS6crFieldAckMask  = 0x100000
)

// GetAck ACK
func (r *registerS6crType) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldAckMask) != 0
}

// SetAck ACK
func (r *registerS6crType) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6crFieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldAckMask)
	}
}

const (
	RegisterS6crFieldPburstShift = 21
	RegisterS6crFieldPburstMask  = 0x600000
)

// GetPburst Peripheral burst transfer configuration
func (r *registerS6crType) GetPburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldPburstMask) >> RegisterS6crFieldPburstShift)
}

// SetPburst Peripheral burst transfer configuration
func (r *registerS6crType) SetPburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldPburstMask)|(uint32(value)<<RegisterS6crFieldPburstShift))
}

const (
	RegisterS6crFieldMburstShift = 23
	RegisterS6crFieldMburstMask  = 0x1800000
)

// GetMburst Memory burst transfer configuration
func (r *registerS6crType) GetMburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS6crFieldMburstMask) >> RegisterS6crFieldMburstShift)
}

// SetMburst Memory burst transfer configuration
func (r *registerS6crType) SetMburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS6crFieldMburstMask)|(uint32(value)<<RegisterS6crFieldMburstShift))
}

// registerS6ndtrType stream x number of data register
type registerS6ndtrType uint32

const (
	RegisterS6ndtrFieldNdtShift = 0
	RegisterS6ndtrFieldNdtMask  = 0xffff
)

// GetNdt Number of data items to transfer
func (r *registerS6ndtrType) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterS6ndtrFieldNdtMask) >> RegisterS6ndtrFieldNdtShift)
}

// SetNdt Number of data items to transfer
func (r *registerS6ndtrType) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS6ndtrFieldNdtMask)|(uint32(value)<<RegisterS6ndtrFieldNdtShift))
}

// registerS6parType stream x peripheral address register
type registerS6parType uint32

const (
	RegisterS6parFieldPaShift = 0
	RegisterS6parFieldPaMask  = 0xffffffff
)

// GetPa Peripheral address
func (r *registerS6parType) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS6parFieldPaMask) >> RegisterS6parFieldPaShift)
}

// SetPa Peripheral address
func (r *registerS6parType) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS6parFieldPaMask)|(uint32(value)<<RegisterS6parFieldPaShift))
}

// registerS6m0arType stream x memory 0 address register
type registerS6m0arType uint32

const (
	RegisterS6m0arFieldM0aShift = 0
	RegisterS6m0arFieldM0aMask  = 0xffffffff
)

// GetM0a Memory 0 address
func (r *registerS6m0arType) GetM0a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS6m0arFieldM0aMask) >> RegisterS6m0arFieldM0aShift)
}

// SetM0a Memory 0 address
func (r *registerS6m0arType) SetM0a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS6m0arFieldM0aMask)|(uint32(value)<<RegisterS6m0arFieldM0aShift))
}

// registerS6m1arType stream x memory 1 address register
type registerS6m1arType uint32

const (
	RegisterS6m1arFieldM1aShift = 0
	RegisterS6m1arFieldM1aMask  = 0xffffffff
)

// GetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS6m1arType) GetM1a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS6m1arFieldM1aMask) >> RegisterS6m1arFieldM1aShift)
}

// SetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS6m1arType) SetM1a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS6m1arFieldM1aMask)|(uint32(value)<<RegisterS6m1arFieldM1aShift))
}

// registerS6fcrType stream x FIFO control register
type registerS6fcrType uint32

const (
	RegisterS6fcrFieldFthShift = 0
	RegisterS6fcrFieldFthMask  = 0x3
)

// GetFth FIFO threshold selection
func (r *registerS6fcrType) GetFth() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS6fcrFieldFthMask) >> RegisterS6fcrFieldFthShift)
}

// SetFth FIFO threshold selection
func (r *registerS6fcrType) SetFth(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS6fcrFieldFthMask)|(uint32(value)<<RegisterS6fcrFieldFthShift))
}

const (
	RegisterS6fcrFieldDmdisShift = 2
	RegisterS6fcrFieldDmdisMask  = 0x4
)

// GetDmdis Direct mode disable
func (r *registerS6fcrType) GetDmdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6fcrFieldDmdisMask) != 0
}

// SetDmdis Direct mode disable
func (r *registerS6fcrType) SetDmdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6fcrFieldDmdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6fcrFieldDmdisMask)
	}
}

const (
	RegisterS6fcrFieldFsShift = 3
	RegisterS6fcrFieldFsMask  = 0x38
)

// GetFs FIFO status
func (r *registerS6fcrType) GetFs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS6fcrFieldFsMask) >> RegisterS6fcrFieldFsShift)
}

const (
	RegisterS6fcrFieldFeieShift = 7
	RegisterS6fcrFieldFeieMask  = 0x80
)

// GetFeie FIFO error interrupt enable
func (r *registerS6fcrType) GetFeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS6fcrFieldFeieMask) != 0
}

// SetFeie FIFO error interrupt enable
func (r *registerS6fcrType) SetFeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS6fcrFieldFeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS6fcrFieldFeieMask)
	}
}

// registerS7crType stream x configuration register
type registerS7crType uint32

const (
	RegisterS7crFieldEnShift = 0
	RegisterS7crFieldEnMask  = 0x1
)

// GetEn Stream enable / flag stream ready when read low
func (r *registerS7crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldEnMask) != 0
}

// SetEn Stream enable / flag stream ready when read low
func (r *registerS7crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldEnMask)
	}
}

const (
	RegisterS7crFieldDmeieShift = 1
	RegisterS7crFieldDmeieMask  = 0x2
)

// GetDmeie Direct mode error interrupt enable
func (r *registerS7crType) GetDmeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldDmeieMask) != 0
}

// SetDmeie Direct mode error interrupt enable
func (r *registerS7crType) SetDmeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldDmeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldDmeieMask)
	}
}

const (
	RegisterS7crFieldTeieShift = 2
	RegisterS7crFieldTeieMask  = 0x4
)

// GetTeie Transfer error interrupt enable
func (r *registerS7crType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable
func (r *registerS7crType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldTeieMask)
	}
}

const (
	RegisterS7crFieldHtieShift = 3
	RegisterS7crFieldHtieMask  = 0x8
)

// GetHtie Half transfer interrupt enable
func (r *registerS7crType) GetHtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldHtieMask) != 0
}

// SetHtie Half transfer interrupt enable
func (r *registerS7crType) SetHtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldHtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldHtieMask)
	}
}

const (
	RegisterS7crFieldTcieShift = 4
	RegisterS7crFieldTcieMask  = 0x10
)

// GetTcie Transfer complete interrupt enable
func (r *registerS7crType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable
func (r *registerS7crType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldTcieMask)
	}
}

const (
	RegisterS7crFieldPfctrlShift = 5
	RegisterS7crFieldPfctrlMask  = 0x20
)

// GetPfctrl Peripheral flow controller
func (r *registerS7crType) GetPfctrl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldPfctrlMask) != 0
}

// SetPfctrl Peripheral flow controller
func (r *registerS7crType) SetPfctrl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldPfctrlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldPfctrlMask)
	}
}

const (
	RegisterS7crFieldDirShift = 6
	RegisterS7crFieldDirMask  = 0xc0
)

// GetDir Data transfer direction
func (r *registerS7crType) GetDir() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldDirMask) >> RegisterS7crFieldDirShift)
}

// SetDir Data transfer direction
func (r *registerS7crType) SetDir(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldDirMask)|(uint32(value)<<RegisterS7crFieldDirShift))
}

const (
	RegisterS7crFieldCircShift = 8
	RegisterS7crFieldCircMask  = 0x100
)

// GetCirc Circular mode
func (r *registerS7crType) GetCirc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldCircMask) != 0
}

// SetCirc Circular mode
func (r *registerS7crType) SetCirc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldCircMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldCircMask)
	}
}

const (
	RegisterS7crFieldPincShift = 9
	RegisterS7crFieldPincMask  = 0x200
)

// GetPinc Peripheral increment mode
func (r *registerS7crType) GetPinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldPincMask) != 0
}

// SetPinc Peripheral increment mode
func (r *registerS7crType) SetPinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldPincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldPincMask)
	}
}

const (
	RegisterS7crFieldMincShift = 10
	RegisterS7crFieldMincMask  = 0x400
)

// GetMinc Memory increment mode
func (r *registerS7crType) GetMinc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldMincMask) != 0
}

// SetMinc Memory increment mode
func (r *registerS7crType) SetMinc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldMincMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldMincMask)
	}
}

const (
	RegisterS7crFieldPsizeShift = 11
	RegisterS7crFieldPsizeMask  = 0x1800
)

// GetPsize Peripheral data size
func (r *registerS7crType) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldPsizeMask) >> RegisterS7crFieldPsizeShift)
}

// SetPsize Peripheral data size
func (r *registerS7crType) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldPsizeMask)|(uint32(value)<<RegisterS7crFieldPsizeShift))
}

const (
	RegisterS7crFieldMsizeShift = 13
	RegisterS7crFieldMsizeMask  = 0x6000
)

// GetMsize Memory data size
func (r *registerS7crType) GetMsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldMsizeMask) >> RegisterS7crFieldMsizeShift)
}

// SetMsize Memory data size
func (r *registerS7crType) SetMsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldMsizeMask)|(uint32(value)<<RegisterS7crFieldMsizeShift))
}

const (
	RegisterS7crFieldPincosShift = 15
	RegisterS7crFieldPincosMask  = 0x8000
)

// GetPincos Peripheral increment offset size
func (r *registerS7crType) GetPincos() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldPincosMask) != 0
}

// SetPincos Peripheral increment offset size
func (r *registerS7crType) SetPincos(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldPincosMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldPincosMask)
	}
}

const (
	RegisterS7crFieldPlShift = 16
	RegisterS7crFieldPlMask  = 0x30000
)

// GetPl Priority level
func (r *registerS7crType) GetPl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldPlMask) >> RegisterS7crFieldPlShift)
}

// SetPl Priority level
func (r *registerS7crType) SetPl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldPlMask)|(uint32(value)<<RegisterS7crFieldPlShift))
}

const (
	RegisterS7crFieldDbmShift = 18
	RegisterS7crFieldDbmMask  = 0x40000
)

// GetDbm Double buffer mode
func (r *registerS7crType) GetDbm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldDbmMask) != 0
}

// SetDbm Double buffer mode
func (r *registerS7crType) SetDbm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldDbmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldDbmMask)
	}
}

const (
	RegisterS7crFieldCtShift = 19
	RegisterS7crFieldCtMask  = 0x80000
)

// GetCt Current target (only in double buffer mode)
func (r *registerS7crType) GetCt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldCtMask) != 0
}

// SetCt Current target (only in double buffer mode)
func (r *registerS7crType) SetCt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldCtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldCtMask)
	}
}

const (
	RegisterS7crFieldAckShift = 20
	RegisterS7crFieldAckMask  = 0x100000
)

// GetAck ACK
func (r *registerS7crType) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldAckMask) != 0
}

// SetAck ACK
func (r *registerS7crType) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7crFieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldAckMask)
	}
}

const (
	RegisterS7crFieldPburstShift = 21
	RegisterS7crFieldPburstMask  = 0x600000
)

// GetPburst Peripheral burst transfer configuration
func (r *registerS7crType) GetPburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldPburstMask) >> RegisterS7crFieldPburstShift)
}

// SetPburst Peripheral burst transfer configuration
func (r *registerS7crType) SetPburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldPburstMask)|(uint32(value)<<RegisterS7crFieldPburstShift))
}

const (
	RegisterS7crFieldMburstShift = 23
	RegisterS7crFieldMburstMask  = 0x1800000
)

// GetMburst Memory burst transfer configuration
func (r *registerS7crType) GetMburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS7crFieldMburstMask) >> RegisterS7crFieldMburstShift)
}

// SetMburst Memory burst transfer configuration
func (r *registerS7crType) SetMburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS7crFieldMburstMask)|(uint32(value)<<RegisterS7crFieldMburstShift))
}

// registerS7ndtrType stream x number of data register
type registerS7ndtrType uint32

const (
	RegisterS7ndtrFieldNdtShift = 0
	RegisterS7ndtrFieldNdtMask  = 0xffff
)

// GetNdt Number of data items to transfer
func (r *registerS7ndtrType) GetNdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterS7ndtrFieldNdtMask) >> RegisterS7ndtrFieldNdtShift)
}

// SetNdt Number of data items to transfer
func (r *registerS7ndtrType) SetNdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS7ndtrFieldNdtMask)|(uint32(value)<<RegisterS7ndtrFieldNdtShift))
}

// registerS7parType stream x peripheral address register
type registerS7parType uint32

const (
	RegisterS7parFieldPaShift = 0
	RegisterS7parFieldPaMask  = 0xffffffff
)

// GetPa Peripheral address
func (r *registerS7parType) GetPa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS7parFieldPaMask) >> RegisterS7parFieldPaShift)
}

// SetPa Peripheral address
func (r *registerS7parType) SetPa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS7parFieldPaMask)|(uint32(value)<<RegisterS7parFieldPaShift))
}

// registerS7m0arType stream x memory 0 address register
type registerS7m0arType uint32

const (
	RegisterS7m0arFieldM0aShift = 0
	RegisterS7m0arFieldM0aMask  = 0xffffffff
)

// GetM0a Memory 0 address
func (r *registerS7m0arType) GetM0a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS7m0arFieldM0aMask) >> RegisterS7m0arFieldM0aShift)
}

// SetM0a Memory 0 address
func (r *registerS7m0arType) SetM0a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS7m0arFieldM0aMask)|(uint32(value)<<RegisterS7m0arFieldM0aShift))
}

// registerS7m1arType stream x memory 1 address register
type registerS7m1arType uint32

const (
	RegisterS7m1arFieldM1aShift = 0
	RegisterS7m1arFieldM1aMask  = 0xffffffff
)

// GetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS7m1arType) GetM1a() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterS7m1arFieldM1aMask) >> RegisterS7m1arFieldM1aShift)
}

// SetM1a Memory 1 address (used in case of Double buffer mode)
func (r *registerS7m1arType) SetM1a(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS7m1arFieldM1aMask)|(uint32(value)<<RegisterS7m1arFieldM1aShift))
}

// registerS7fcrType stream x FIFO control register
type registerS7fcrType uint32

const (
	RegisterS7fcrFieldFthShift = 0
	RegisterS7fcrFieldFthMask  = 0x3
)

// GetFth FIFO threshold selection
func (r *registerS7fcrType) GetFth() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS7fcrFieldFthMask) >> RegisterS7fcrFieldFthShift)
}

// SetFth FIFO threshold selection
func (r *registerS7fcrType) SetFth(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterS7fcrFieldFthMask)|(uint32(value)<<RegisterS7fcrFieldFthShift))
}

const (
	RegisterS7fcrFieldDmdisShift = 2
	RegisterS7fcrFieldDmdisMask  = 0x4
)

// GetDmdis Direct mode disable
func (r *registerS7fcrType) GetDmdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7fcrFieldDmdisMask) != 0
}

// SetDmdis Direct mode disable
func (r *registerS7fcrType) SetDmdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7fcrFieldDmdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7fcrFieldDmdisMask)
	}
}

const (
	RegisterS7fcrFieldFsShift = 3
	RegisterS7fcrFieldFsMask  = 0x38
)

// GetFs FIFO status
func (r *registerS7fcrType) GetFs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterS7fcrFieldFsMask) >> RegisterS7fcrFieldFsShift)
}

const (
	RegisterS7fcrFieldFeieShift = 7
	RegisterS7fcrFieldFeieMask  = 0x80
)

// GetFeie FIFO error interrupt enable
func (r *registerS7fcrType) GetFeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterS7fcrFieldFeieMask) != 0
}

// SetFeie FIFO error interrupt enable
func (r *registerS7fcrType) SetFeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterS7fcrFieldFeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterS7fcrFieldFeieMask)
	}
}
