//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package ramecc3

import (
	"unsafe"
	"volatile"
)

var (
	Ramecc3 = (*_ramecc3)(unsafe.Pointer(uintptr(0x58027000)))
)

type _ramecc3 struct {
	Ier    RegisterIerType
	_      [28]byte
	M1cr   RegisterM1crType
	M1sr   RegisterM1srType
	M1far  RegisterM1farType
	M1fdrl RegisterM1fdrlType
	M1fdrh RegisterM1fdrhType
	M1fecr RegisterM1fecrType
	_      [8]byte
	M2cr   RegisterM2crType
	M2sr   RegisterM2srType
	M2far  RegisterM2farType
	M2fdrl RegisterM2fdrlType
	M2fdrh RegisterM2fdrhType
	_      [4]byte
	M2fecr RegisterM2fecrType
}

// RegisterIerType RAMECC interrupt enable register
type RegisterIerType uint32

func (r *RegisterIerType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIerType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIerType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIerType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIerType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIerFieldGieShift = 0
	RegisterIerFieldGieMask  = 0x1
)

// GetGie Global interrupt enable
func (r *RegisterIerType) GetGie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldGieMask) != 0
}

// SetGie Global interrupt enable
func (r *RegisterIerType) SetGie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldGieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldGieMask)
	}
}

const (
	RegisterIerFieldGeccseieShift = 1
	RegisterIerFieldGeccseieMask  = 0x2
)

// GetGeccseie Global ECC single error interrupt enable
func (r *RegisterIerType) GetGeccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldGeccseieMask) != 0
}

// SetGeccseie Global ECC single error interrupt enable
func (r *RegisterIerType) SetGeccseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldGeccseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldGeccseieMask)
	}
}

const (
	RegisterIerFieldGeccdeieShift = 2
	RegisterIerFieldGeccdeieMask  = 0x4
)

// GetGeccdeie Global ECC double error interrupt enable
func (r *RegisterIerType) GetGeccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldGeccdeieMask) != 0
}

// SetGeccdeie Global ECC double error interrupt enable
func (r *RegisterIerType) SetGeccdeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldGeccdeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldGeccdeieMask)
	}
}

const (
	RegisterIerFieldGeccdebwieShift = 3
	RegisterIerFieldGeccdebwieMask  = 0x8
)

// GetGeccdebwie Global ECC double error on byte write (BW) interrupt enable
func (r *RegisterIerType) GetGeccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldGeccdebwieMask) != 0
}

// SetGeccdebwie Global ECC double error on byte write (BW) interrupt enable
func (r *RegisterIerType) SetGeccdebwie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldGeccdebwieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldGeccdebwieMask)
	}
}

// RegisterM1crType RAMECC monitor x configuration register
type RegisterM1crType uint32

func (r *RegisterM1crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM1crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM1crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM1crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM1crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM1crFieldEccseieShift = 2
	RegisterM1crFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *RegisterM1crType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1crFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *RegisterM1crType) SetEccseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1crFieldEccseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1crFieldEccseieMask)
	}
}

const (
	RegisterM1crFieldEccdeieShift = 3
	RegisterM1crFieldEccdeieMask  = 0x8
)

// GetEccdeie ECC double error interrupt enable
func (r *RegisterM1crType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1crFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *RegisterM1crType) SetEccdeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1crFieldEccdeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1crFieldEccdeieMask)
	}
}

const (
	RegisterM1crFieldEccdebwieShift = 4
	RegisterM1crFieldEccdebwieMask  = 0x10
)

// GetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM1crType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1crFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM1crType) SetEccdebwie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1crFieldEccdebwieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1crFieldEccdebwieMask)
	}
}

const (
	RegisterM1crFieldEccelenShift = 5
	RegisterM1crFieldEccelenMask  = 0x20
)

// GetEccelen ECC error latching enable
func (r *RegisterM1crType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1crFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *RegisterM1crType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1crFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1crFieldEccelenMask)
	}
}

// RegisterM1srType RAMECC monitor x status register
type RegisterM1srType uint32

func (r *RegisterM1srType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM1srType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM1srType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM1srType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM1srType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM1srFieldSedcfShift = 0
	RegisterM1srFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *RegisterM1srType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *RegisterM1srType) SetSedcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1srFieldSedcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1srFieldSedcfMask)
	}
}

const (
	RegisterM1srFieldDedfShift = 1
	RegisterM1srFieldDedfMask  = 0x2
)

// GetDedf ECC double error detected flag
func (r *RegisterM1srType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *RegisterM1srType) SetDedf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1srFieldDedfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1srFieldDedfMask)
	}
}

const (
	RegisterM1srFieldDebwdfShift = 2
	RegisterM1srFieldDebwdfMask  = 0x4
)

// GetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM1srType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM1srType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1srFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1srFieldDebwdfMask)
	}
}

// RegisterM1farType RAMECC monitor x failing address register
type RegisterM1farType uint32

func (r *RegisterM1farType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM1farType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM1farType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM1farType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM1farType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM1farFieldFaddShift = 0
	RegisterM1farFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *RegisterM1farType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM1farFieldFaddMask) >> RegisterM1farFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *RegisterM1farType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM1farFieldFaddMask)|(uint32(value)<<RegisterM1farFieldFaddShift))
}

// RegisterM1fdrlType RAMECC monitor x failing data low register
type RegisterM1fdrlType uint32

func (r *RegisterM1fdrlType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM1fdrlType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM1fdrlType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM1fdrlType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM1fdrlType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM1fdrlFieldFdatalShift = 0
	RegisterM1fdrlFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *RegisterM1fdrlType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrlFieldFdatalMask) >> RegisterM1fdrlFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *RegisterM1fdrlType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrlFieldFdatalMask)|(uint32(value)<<RegisterM1fdrlFieldFdatalShift))
}

// RegisterM1fdrhType RAMECC monitor x failing data high register
type RegisterM1fdrhType uint32

func (r *RegisterM1fdrhType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM1fdrhType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM1fdrhType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM1fdrhType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM1fdrhType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM1fdrhFieldFdatahShift = 0
	RegisterM1fdrhFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *RegisterM1fdrhType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrhFieldFdatahMask) >> RegisterM1fdrhFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *RegisterM1fdrhType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrhFieldFdatahMask)|(uint32(value)<<RegisterM1fdrhFieldFdatahShift))
}

// RegisterM1fecrType RAMECC monitor x failing ECC error code register
type RegisterM1fecrType uint32

func (r *RegisterM1fecrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM1fecrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM1fecrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM1fecrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM1fecrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM1fecrFieldFecShift = 0
	RegisterM1fecrFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *RegisterM1fecrType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM1fecrFieldFecMask) >> RegisterM1fecrFieldFecShift)
}

// SetFec Failing error code
func (r *RegisterM1fecrType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM1fecrFieldFecMask)|(uint32(value)<<RegisterM1fecrFieldFecShift))
}

// RegisterM2crType RAMECC monitor x configuration register
type RegisterM2crType uint32

func (r *RegisterM2crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM2crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM2crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM2crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM2crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM2crFieldEccseieShift = 2
	RegisterM2crFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *RegisterM2crType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *RegisterM2crType) SetEccseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2crFieldEccseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2crFieldEccseieMask)
	}
}

const (
	RegisterM2crFieldEccdeieShift = 3
	RegisterM2crFieldEccdeieMask  = 0x8
)

// GetEccdeie ECC double error interrupt enable
func (r *RegisterM2crType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *RegisterM2crType) SetEccdeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2crFieldEccdeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2crFieldEccdeieMask)
	}
}

const (
	RegisterM2crFieldEccdebwieShift = 4
	RegisterM2crFieldEccdebwieMask  = 0x10
)

// GetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM2crType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM2crType) SetEccdebwie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2crFieldEccdebwieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2crFieldEccdebwieMask)
	}
}

const (
	RegisterM2crFieldEccelenShift = 5
	RegisterM2crFieldEccelenMask  = 0x20
)

// GetEccelen ECC error latching enable
func (r *RegisterM2crType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *RegisterM2crType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2crFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2crFieldEccelenMask)
	}
}

// RegisterM2srType RAMECC monitor x status register
type RegisterM2srType uint32

func (r *RegisterM2srType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM2srType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM2srType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM2srType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM2srType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM2srFieldSedcfShift = 0
	RegisterM2srFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *RegisterM2srType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2srFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *RegisterM2srType) SetSedcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2srFieldSedcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2srFieldSedcfMask)
	}
}

const (
	RegisterM2srFieldDedfShift = 1
	RegisterM2srFieldDedfMask  = 0x2
)

// GetDedf ECC double error detected flag
func (r *RegisterM2srType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2srFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *RegisterM2srType) SetDedf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2srFieldDedfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2srFieldDedfMask)
	}
}

const (
	RegisterM2srFieldDebwdfShift = 2
	RegisterM2srFieldDebwdfMask  = 0x4
)

// GetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM2srType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2srFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM2srType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2srFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2srFieldDebwdfMask)
	}
}

// RegisterM2farType RAMECC monitor x failing address register
type RegisterM2farType uint32

func (r *RegisterM2farType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM2farType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM2farType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM2farType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM2farType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM2farFieldFaddShift = 0
	RegisterM2farFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *RegisterM2farType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM2farFieldFaddMask) >> RegisterM2farFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *RegisterM2farType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM2farFieldFaddMask)|(uint32(value)<<RegisterM2farFieldFaddShift))
}

// RegisterM2fdrlType RAMECC monitor x failing data low register
type RegisterM2fdrlType uint32

func (r *RegisterM2fdrlType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM2fdrlType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM2fdrlType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM2fdrlType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM2fdrlType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM2fdrlFieldFdatalShift = 0
	RegisterM2fdrlFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *RegisterM2fdrlType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM2fdrlFieldFdatalMask) >> RegisterM2fdrlFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *RegisterM2fdrlType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM2fdrlFieldFdatalMask)|(uint32(value)<<RegisterM2fdrlFieldFdatalShift))
}

// RegisterM2fdrhType RAMECC monitor x failing data high register
type RegisterM2fdrhType uint32

func (r *RegisterM2fdrhType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM2fdrhType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM2fdrhType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM2fdrhType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM2fdrhType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM2fdrhFieldFdatahShift = 0
	RegisterM2fdrhFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *RegisterM2fdrhType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM2fdrhFieldFdatahMask) >> RegisterM2fdrhFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *RegisterM2fdrhType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM2fdrhFieldFdatahMask)|(uint32(value)<<RegisterM2fdrhFieldFdatahShift))
}

// RegisterM2fecrType RAMECC monitor x failing ECC error code register
type RegisterM2fecrType uint32

func (r *RegisterM2fecrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM2fecrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM2fecrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM2fecrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM2fecrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM2fecrFieldFecShift = 0
	RegisterM2fecrFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *RegisterM2fecrType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM2fecrFieldFecMask) >> RegisterM2fecrFieldFecShift)
}

// SetFec Failing error code
func (r *RegisterM2fecrType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM2fecrFieldFecMask)|(uint32(value)<<RegisterM2fecrFieldFecShift))
}
