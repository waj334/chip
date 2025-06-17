//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package ramecc2

import (
	"unsafe"
	"volatile"
)

var (
	Ramecc2 = (*_ramecc2)(unsafe.Pointer(uintptr(0x48023000)))
)

type _ramecc2 struct {
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
	_      [4]byte
	M3cr   RegisterM3crType
	M3sr   RegisterM3srType
	M3far  RegisterM3farType
	M3fdrl RegisterM3fdrlType
	M3fdrh RegisterM3fdrhType
	_      [8]byte
	M3fecr RegisterM3fecrType
	M4cr   RegisterM4crType
	M4sr   RegisterM4srType
	M4far  RegisterM4farType
	M4fdrl RegisterM4fdrlType
	M4fdrh RegisterM4fdrhType
	M4fecr RegisterM4fecrType
	_      [12]byte
	M5cr   RegisterM5crType
	M5sr   RegisterM5srType
	M5far  RegisterM5farType
	M5fdrl RegisterM5fdrlType
	M5fdrh RegisterM5fdrhType
	M5fecr RegisterM5fecrType
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

// RegisterM3crType RAMECC monitor x configuration register
type RegisterM3crType uint32

func (r *RegisterM3crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM3crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM3crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM3crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM3crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM3crFieldEccseieShift = 2
	RegisterM3crFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *RegisterM3crType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3crFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *RegisterM3crType) SetEccseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM3crFieldEccseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM3crFieldEccseieMask)
	}
}

const (
	RegisterM3crFieldEccdeieShift = 3
	RegisterM3crFieldEccdeieMask  = 0x8
)

// GetEccdeie ECC double error interrupt enable
func (r *RegisterM3crType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3crFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *RegisterM3crType) SetEccdeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM3crFieldEccdeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM3crFieldEccdeieMask)
	}
}

const (
	RegisterM3crFieldEccdebwieShift = 4
	RegisterM3crFieldEccdebwieMask  = 0x10
)

// GetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM3crType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3crFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM3crType) SetEccdebwie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM3crFieldEccdebwieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM3crFieldEccdebwieMask)
	}
}

const (
	RegisterM3crFieldEccelenShift = 5
	RegisterM3crFieldEccelenMask  = 0x20
)

// GetEccelen ECC error latching enable
func (r *RegisterM3crType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3crFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *RegisterM3crType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM3crFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM3crFieldEccelenMask)
	}
}

// RegisterM3srType RAMECC monitor x status register
type RegisterM3srType uint32

func (r *RegisterM3srType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM3srType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM3srType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM3srType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM3srType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM3srFieldSedcfShift = 0
	RegisterM3srFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *RegisterM3srType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3srFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *RegisterM3srType) SetSedcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM3srFieldSedcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM3srFieldSedcfMask)
	}
}

const (
	RegisterM3srFieldDedfShift = 1
	RegisterM3srFieldDedfMask  = 0x2
)

// GetDedf ECC double error detected flag
func (r *RegisterM3srType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3srFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *RegisterM3srType) SetDedf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM3srFieldDedfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM3srFieldDedfMask)
	}
}

const (
	RegisterM3srFieldDebwdfShift = 2
	RegisterM3srFieldDebwdfMask  = 0x4
)

// GetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM3srType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3srFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM3srType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM3srFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM3srFieldDebwdfMask)
	}
}

// RegisterM3farType RAMECC monitor x failing address register
type RegisterM3farType uint32

func (r *RegisterM3farType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM3farType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM3farType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM3farType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM3farType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM3farFieldFaddShift = 0
	RegisterM3farFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *RegisterM3farType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3farFieldFaddMask) >> RegisterM3farFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *RegisterM3farType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3farFieldFaddMask)|(uint32(value)<<RegisterM3farFieldFaddShift))
}

// RegisterM3fdrlType RAMECC monitor x failing data low register
type RegisterM3fdrlType uint32

func (r *RegisterM3fdrlType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM3fdrlType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM3fdrlType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM3fdrlType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM3fdrlType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM3fdrlFieldFdatalShift = 0
	RegisterM3fdrlFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *RegisterM3fdrlType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3fdrlFieldFdatalMask) >> RegisterM3fdrlFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *RegisterM3fdrlType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3fdrlFieldFdatalMask)|(uint32(value)<<RegisterM3fdrlFieldFdatalShift))
}

// RegisterM3fdrhType RAMECC monitor x failing data high register
type RegisterM3fdrhType uint32

func (r *RegisterM3fdrhType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM3fdrhType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM3fdrhType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM3fdrhType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM3fdrhType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM3fdrhFieldFdatahShift = 0
	RegisterM3fdrhFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *RegisterM3fdrhType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3fdrhFieldFdatahMask) >> RegisterM3fdrhFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *RegisterM3fdrhType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3fdrhFieldFdatahMask)|(uint32(value)<<RegisterM3fdrhFieldFdatahShift))
}

// RegisterM3fecrType RAMECC monitor x failing ECC error code register
type RegisterM3fecrType uint32

func (r *RegisterM3fecrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM3fecrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM3fecrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM3fecrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM3fecrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM3fecrFieldFecShift = 0
	RegisterM3fecrFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *RegisterM3fecrType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3fecrFieldFecMask) >> RegisterM3fecrFieldFecShift)
}

// SetFec Failing error code
func (r *RegisterM3fecrType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3fecrFieldFecMask)|(uint32(value)<<RegisterM3fecrFieldFecShift))
}

// RegisterM4crType RAMECC monitor x configuration register
type RegisterM4crType uint32

func (r *RegisterM4crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM4crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM4crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM4crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM4crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM4crFieldEccseieShift = 2
	RegisterM4crFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *RegisterM4crType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4crFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *RegisterM4crType) SetEccseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM4crFieldEccseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM4crFieldEccseieMask)
	}
}

const (
	RegisterM4crFieldEccdeieShift = 3
	RegisterM4crFieldEccdeieMask  = 0x8
)

// GetEccdeie ECC double error interrupt enable
func (r *RegisterM4crType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4crFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *RegisterM4crType) SetEccdeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM4crFieldEccdeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM4crFieldEccdeieMask)
	}
}

const (
	RegisterM4crFieldEccdebwieShift = 4
	RegisterM4crFieldEccdebwieMask  = 0x10
)

// GetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM4crType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4crFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM4crType) SetEccdebwie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM4crFieldEccdebwieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM4crFieldEccdebwieMask)
	}
}

const (
	RegisterM4crFieldEccelenShift = 5
	RegisterM4crFieldEccelenMask  = 0x20
)

// GetEccelen ECC error latching enable
func (r *RegisterM4crType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4crFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *RegisterM4crType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM4crFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM4crFieldEccelenMask)
	}
}

// RegisterM4srType RAMECC monitor x status register
type RegisterM4srType uint32

func (r *RegisterM4srType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM4srType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM4srType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM4srType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM4srType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM4srFieldSedcfShift = 0
	RegisterM4srFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *RegisterM4srType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4srFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *RegisterM4srType) SetSedcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM4srFieldSedcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM4srFieldSedcfMask)
	}
}

const (
	RegisterM4srFieldDedfShift = 1
	RegisterM4srFieldDedfMask  = 0x2
)

// GetDedf ECC double error detected flag
func (r *RegisterM4srType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4srFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *RegisterM4srType) SetDedf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM4srFieldDedfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM4srFieldDedfMask)
	}
}

const (
	RegisterM4srFieldDebwdfShift = 2
	RegisterM4srFieldDebwdfMask  = 0x4
)

// GetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM4srType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4srFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM4srType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM4srFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM4srFieldDebwdfMask)
	}
}

// RegisterM4farType RAMECC monitor x failing address register
type RegisterM4farType uint32

func (r *RegisterM4farType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM4farType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM4farType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM4farType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM4farType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM4farFieldFaddShift = 0
	RegisterM4farFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *RegisterM4farType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4farFieldFaddMask) >> RegisterM4farFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *RegisterM4farType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4farFieldFaddMask)|(uint32(value)<<RegisterM4farFieldFaddShift))
}

// RegisterM4fdrlType RAMECC monitor x failing data low register
type RegisterM4fdrlType uint32

func (r *RegisterM4fdrlType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM4fdrlType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM4fdrlType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM4fdrlType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM4fdrlType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM4fdrlFieldFdatalShift = 0
	RegisterM4fdrlFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *RegisterM4fdrlType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4fdrlFieldFdatalMask) >> RegisterM4fdrlFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *RegisterM4fdrlType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4fdrlFieldFdatalMask)|(uint32(value)<<RegisterM4fdrlFieldFdatalShift))
}

// RegisterM4fdrhType RAMECC monitor x failing data high register
type RegisterM4fdrhType uint32

func (r *RegisterM4fdrhType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM4fdrhType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM4fdrhType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM4fdrhType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM4fdrhType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM4fdrhFieldFdatahShift = 0
	RegisterM4fdrhFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *RegisterM4fdrhType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4fdrhFieldFdatahMask) >> RegisterM4fdrhFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *RegisterM4fdrhType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4fdrhFieldFdatahMask)|(uint32(value)<<RegisterM4fdrhFieldFdatahShift))
}

// RegisterM4fecrType RAMECC monitor x failing ECC error code register
type RegisterM4fecrType uint32

func (r *RegisterM4fecrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM4fecrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM4fecrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM4fecrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM4fecrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM4fecrFieldFecShift = 0
	RegisterM4fecrFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *RegisterM4fecrType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4fecrFieldFecMask) >> RegisterM4fecrFieldFecShift)
}

// SetFec Failing error code
func (r *RegisterM4fecrType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4fecrFieldFecMask)|(uint32(value)<<RegisterM4fecrFieldFecShift))
}

// RegisterM5crType RAMECC monitor x configuration register
type RegisterM5crType uint32

func (r *RegisterM5crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM5crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM5crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM5crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM5crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM5crFieldEccseieShift = 2
	RegisterM5crFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *RegisterM5crType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5crFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *RegisterM5crType) SetEccseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM5crFieldEccseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM5crFieldEccseieMask)
	}
}

const (
	RegisterM5crFieldEccdeieShift = 3
	RegisterM5crFieldEccdeieMask  = 0x8
)

// GetEccdeie ECC double error interrupt enable
func (r *RegisterM5crType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5crFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *RegisterM5crType) SetEccdeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM5crFieldEccdeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM5crFieldEccdeieMask)
	}
}

const (
	RegisterM5crFieldEccdebwieShift = 4
	RegisterM5crFieldEccdebwieMask  = 0x10
)

// GetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM5crType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5crFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM5crType) SetEccdebwie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM5crFieldEccdebwieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM5crFieldEccdebwieMask)
	}
}

const (
	RegisterM5crFieldEccelenShift = 5
	RegisterM5crFieldEccelenMask  = 0x20
)

// GetEccelen ECC error latching enable
func (r *RegisterM5crType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5crFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *RegisterM5crType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM5crFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM5crFieldEccelenMask)
	}
}

// RegisterM5srType RAMECC monitor x status register
type RegisterM5srType uint32

func (r *RegisterM5srType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM5srType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM5srType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM5srType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM5srType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM5srFieldSedcfShift = 0
	RegisterM5srFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *RegisterM5srType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5srFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *RegisterM5srType) SetSedcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM5srFieldSedcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM5srFieldSedcfMask)
	}
}

const (
	RegisterM5srFieldDedfShift = 1
	RegisterM5srFieldDedfMask  = 0x2
)

// GetDedf ECC double error detected flag
func (r *RegisterM5srType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5srFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *RegisterM5srType) SetDedf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM5srFieldDedfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM5srFieldDedfMask)
	}
}

const (
	RegisterM5srFieldDebwdfShift = 2
	RegisterM5srFieldDebwdfMask  = 0x4
)

// GetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM5srType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5srFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM5srType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM5srFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM5srFieldDebwdfMask)
	}
}

// RegisterM5farType RAMECC monitor x failing address register
type RegisterM5farType uint32

func (r *RegisterM5farType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM5farType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM5farType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM5farType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM5farType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM5farFieldFaddShift = 0
	RegisterM5farFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *RegisterM5farType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5farFieldFaddMask) >> RegisterM5farFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *RegisterM5farType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5farFieldFaddMask)|(uint32(value)<<RegisterM5farFieldFaddShift))
}

// RegisterM5fdrlType RAMECC monitor x failing data low register
type RegisterM5fdrlType uint32

func (r *RegisterM5fdrlType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM5fdrlType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM5fdrlType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM5fdrlType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM5fdrlType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM5fdrlFieldFdatalShift = 0
	RegisterM5fdrlFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *RegisterM5fdrlType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5fdrlFieldFdatalMask) >> RegisterM5fdrlFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *RegisterM5fdrlType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5fdrlFieldFdatalMask)|(uint32(value)<<RegisterM5fdrlFieldFdatalShift))
}

// RegisterM5fdrhType RAMECC monitor x failing data high register
type RegisterM5fdrhType uint32

func (r *RegisterM5fdrhType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM5fdrhType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM5fdrhType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM5fdrhType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM5fdrhType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM5fdrhFieldFecShift = 0
	RegisterM5fdrhFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *RegisterM5fdrhType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5fdrhFieldFecMask) >> RegisterM5fdrhFieldFecShift)
}

// SetFec Failing error code
func (r *RegisterM5fdrhType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5fdrhFieldFecMask)|(uint32(value)<<RegisterM5fdrhFieldFecShift))
}

// RegisterM5fecrType RAMECC monitor x failing ECC error code register
type RegisterM5fecrType uint32

func (r *RegisterM5fecrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterM5fecrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterM5fecrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterM5fecrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterM5fecrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterM5fecrFieldFecShift = 0
	RegisterM5fecrFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *RegisterM5fecrType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5fecrFieldFecMask) >> RegisterM5fecrFieldFecShift)
}

// SetFec Failing error code
func (r *RegisterM5fecrType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5fecrFieldFecMask)|(uint32(value)<<RegisterM5fecrFieldFecShift))
}
