//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package ramecc1

import (
	"unsafe"
	"volatile"
)

var (
	Ramecc1 = (*_ramecc1)(unsafe.Pointer(uintptr(0x52009000)))
)

type _ramecc1 struct {
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
	RegisterM1srFieldEccseieShift = 2
	RegisterM1srFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *RegisterM1srType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *RegisterM1srType) SetEccseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1srFieldEccseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1srFieldEccseieMask)
	}
}

const (
	RegisterM1srFieldEccdeieShift = 3
	RegisterM1srFieldEccdeieMask  = 0x8
)

// GetEccdeie ECC double error interrupt enable
func (r *RegisterM1srType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *RegisterM1srType) SetEccdeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1srFieldEccdeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1srFieldEccdeieMask)
	}
}

const (
	RegisterM1srFieldEccdebwieShift = 4
	RegisterM1srFieldEccdebwieMask  = 0x10
)

// GetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM1srType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM1srType) SetEccdebwie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1srFieldEccdebwieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1srFieldEccdebwieMask)
	}
}

const (
	RegisterM1srFieldEccelenShift = 5
	RegisterM1srFieldEccelenMask  = 0x20
)

// GetEccelen ECC error latching enable
func (r *RegisterM1srType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *RegisterM1srType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1srFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1srFieldEccelenMask)
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
	RegisterM1farFieldEccseieShift = 2
	RegisterM1farFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *RegisterM1farType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1farFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *RegisterM1farType) SetEccseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1farFieldEccseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1farFieldEccseieMask)
	}
}

const (
	RegisterM1farFieldEccdeieShift = 3
	RegisterM1farFieldEccdeieMask  = 0x8
)

// GetEccdeie ECC double error interrupt enable
func (r *RegisterM1farType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1farFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *RegisterM1farType) SetEccdeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1farFieldEccdeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1farFieldEccdeieMask)
	}
}

const (
	RegisterM1farFieldEccdebwieShift = 4
	RegisterM1farFieldEccdebwieMask  = 0x10
)

// GetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM1farType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1farFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM1farType) SetEccdebwie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1farFieldEccdebwieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1farFieldEccdebwieMask)
	}
}

const (
	RegisterM1farFieldEccelenShift = 5
	RegisterM1farFieldEccelenMask  = 0x20
)

// GetEccelen ECC error latching enable
func (r *RegisterM1farType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1farFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *RegisterM1farType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1farFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1farFieldEccelenMask)
	}
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
	RegisterM1fdrlFieldEccseieShift = 2
	RegisterM1fdrlFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *RegisterM1fdrlType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrlFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *RegisterM1fdrlType) SetEccseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fdrlFieldEccseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrlFieldEccseieMask)
	}
}

const (
	RegisterM1fdrlFieldEccdeieShift = 3
	RegisterM1fdrlFieldEccdeieMask  = 0x8
)

// GetEccdeie ECC double error interrupt enable
func (r *RegisterM1fdrlType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrlFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *RegisterM1fdrlType) SetEccdeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fdrlFieldEccdeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrlFieldEccdeieMask)
	}
}

const (
	RegisterM1fdrlFieldEccdebwieShift = 4
	RegisterM1fdrlFieldEccdebwieMask  = 0x10
)

// GetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM1fdrlType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrlFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM1fdrlType) SetEccdebwie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fdrlFieldEccdebwieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrlFieldEccdebwieMask)
	}
}

const (
	RegisterM1fdrlFieldEccelenShift = 5
	RegisterM1fdrlFieldEccelenMask  = 0x20
)

// GetEccelen ECC error latching enable
func (r *RegisterM1fdrlType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrlFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *RegisterM1fdrlType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fdrlFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrlFieldEccelenMask)
	}
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
	RegisterM1fdrhFieldEccseieShift = 2
	RegisterM1fdrhFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *RegisterM1fdrhType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrhFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *RegisterM1fdrhType) SetEccseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fdrhFieldEccseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrhFieldEccseieMask)
	}
}

const (
	RegisterM1fdrhFieldEccdeieShift = 3
	RegisterM1fdrhFieldEccdeieMask  = 0x8
)

// GetEccdeie ECC double error interrupt enable
func (r *RegisterM1fdrhType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrhFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *RegisterM1fdrhType) SetEccdeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fdrhFieldEccdeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrhFieldEccdeieMask)
	}
}

const (
	RegisterM1fdrhFieldEccdebwieShift = 4
	RegisterM1fdrhFieldEccdebwieMask  = 0x10
)

// GetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM1fdrhType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrhFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *RegisterM1fdrhType) SetEccdebwie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fdrhFieldEccdebwieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrhFieldEccdebwieMask)
	}
}

const (
	RegisterM1fdrhFieldEccelenShift = 5
	RegisterM1fdrhFieldEccelenMask  = 0x20
)

// GetEccelen ECC error latching enable
func (r *RegisterM1fdrhType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrhFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *RegisterM1fdrhType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fdrhFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrhFieldEccelenMask)
	}
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
	RegisterM1fecrFieldSedcfShift = 0
	RegisterM1fecrFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *RegisterM1fecrType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fecrFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *RegisterM1fecrType) SetSedcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fecrFieldSedcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fecrFieldSedcfMask)
	}
}

const (
	RegisterM1fecrFieldDedfShift = 1
	RegisterM1fecrFieldDedfMask  = 0x2
)

// GetDedf ECC double error detected flag
func (r *RegisterM1fecrType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fecrFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *RegisterM1fecrType) SetDedf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fecrFieldDedfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fecrFieldDedfMask)
	}
}

const (
	RegisterM1fecrFieldDebwdfShift = 2
	RegisterM1fecrFieldDebwdfMask  = 0x4
)

// GetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM1fecrType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fecrFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM1fecrType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fecrFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fecrFieldDebwdfMask)
	}
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
	RegisterM2crFieldSedcfShift = 0
	RegisterM2crFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *RegisterM2crType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *RegisterM2crType) SetSedcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2crFieldSedcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2crFieldSedcfMask)
	}
}

const (
	RegisterM2crFieldDedfShift = 1
	RegisterM2crFieldDedfMask  = 0x2
)

// GetDedf ECC double error detected flag
func (r *RegisterM2crType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *RegisterM2crType) SetDedf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2crFieldDedfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2crFieldDedfMask)
	}
}

const (
	RegisterM2crFieldDebwdfShift = 2
	RegisterM2crFieldDebwdfMask  = 0x4
)

// GetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM2crType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM2crType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2crFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2crFieldDebwdfMask)
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
	RegisterM2farFieldSedcfShift = 0
	RegisterM2farFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *RegisterM2farType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2farFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *RegisterM2farType) SetSedcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2farFieldSedcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2farFieldSedcfMask)
	}
}

const (
	RegisterM2farFieldDedfShift = 1
	RegisterM2farFieldDedfMask  = 0x2
)

// GetDedf ECC double error detected flag
func (r *RegisterM2farType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2farFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *RegisterM2farType) SetDedf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2farFieldDedfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2farFieldDedfMask)
	}
}

const (
	RegisterM2farFieldDebwdfShift = 2
	RegisterM2farFieldDebwdfMask  = 0x4
)

// GetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM2farType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2farFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM2farType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2farFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2farFieldDebwdfMask)
	}
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
	RegisterM2fdrlFieldSedcfShift = 0
	RegisterM2fdrlFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *RegisterM2fdrlType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2fdrlFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *RegisterM2fdrlType) SetSedcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2fdrlFieldSedcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2fdrlFieldSedcfMask)
	}
}

const (
	RegisterM2fdrlFieldDedfShift = 1
	RegisterM2fdrlFieldDedfMask  = 0x2
)

// GetDedf ECC double error detected flag
func (r *RegisterM2fdrlType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2fdrlFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *RegisterM2fdrlType) SetDedf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2fdrlFieldDedfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2fdrlFieldDedfMask)
	}
}

const (
	RegisterM2fdrlFieldDebwdfShift = 2
	RegisterM2fdrlFieldDebwdfMask  = 0x4
)

// GetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM2fdrlType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2fdrlFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *RegisterM2fdrlType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2fdrlFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2fdrlFieldDebwdfMask)
	}
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
	RegisterM2fdrhFieldFaddShift = 0
	RegisterM2fdrhFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *RegisterM2fdrhType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM2fdrhFieldFaddMask) >> RegisterM2fdrhFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *RegisterM2fdrhType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM2fdrhFieldFaddMask)|(uint32(value)<<RegisterM2fdrhFieldFaddShift))
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
	RegisterM2fecrFieldFaddShift = 0
	RegisterM2fecrFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *RegisterM2fecrType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM2fecrFieldFaddMask) >> RegisterM2fecrFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *RegisterM2fecrType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM2fecrFieldFaddMask)|(uint32(value)<<RegisterM2fecrFieldFaddShift))
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
	RegisterM3crFieldFaddShift = 0
	RegisterM3crFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *RegisterM3crType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3crFieldFaddMask) >> RegisterM3crFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *RegisterM3crType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3crFieldFaddMask)|(uint32(value)<<RegisterM3crFieldFaddShift))
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
	RegisterM3srFieldFaddShift = 0
	RegisterM3srFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *RegisterM3srType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3srFieldFaddMask) >> RegisterM3srFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *RegisterM3srType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3srFieldFaddMask)|(uint32(value)<<RegisterM3srFieldFaddShift))
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
	RegisterM3fdrhFieldFdatalShift = 0
	RegisterM3fdrhFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *RegisterM3fdrhType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3fdrhFieldFdatalMask) >> RegisterM3fdrhFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *RegisterM3fdrhType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3fdrhFieldFdatalMask)|(uint32(value)<<RegisterM3fdrhFieldFdatalShift))
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
	RegisterM3fecrFieldFdatalShift = 0
	RegisterM3fecrFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *RegisterM3fecrType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3fecrFieldFdatalMask) >> RegisterM3fecrFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *RegisterM3fecrType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3fecrFieldFdatalMask)|(uint32(value)<<RegisterM3fecrFieldFdatalShift))
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
	RegisterM4crFieldFdatalShift = 0
	RegisterM4crFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *RegisterM4crType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4crFieldFdatalMask) >> RegisterM4crFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *RegisterM4crType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4crFieldFdatalMask)|(uint32(value)<<RegisterM4crFieldFdatalShift))
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
	RegisterM4srFieldFdatalShift = 0
	RegisterM4srFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *RegisterM4srType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4srFieldFdatalMask) >> RegisterM4srFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *RegisterM4srType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4srFieldFdatalMask)|(uint32(value)<<RegisterM4srFieldFdatalShift))
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
	RegisterM4farFieldFdatahShift = 0
	RegisterM4farFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *RegisterM4farType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4farFieldFdatahMask) >> RegisterM4farFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *RegisterM4farType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4farFieldFdatahMask)|(uint32(value)<<RegisterM4farFieldFdatahShift))
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
	RegisterM4fdrlFieldFdatahShift = 0
	RegisterM4fdrlFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *RegisterM4fdrlType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4fdrlFieldFdatahMask) >> RegisterM4fdrlFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *RegisterM4fdrlType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4fdrlFieldFdatahMask)|(uint32(value)<<RegisterM4fdrlFieldFdatahShift))
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
	RegisterM4fecrFieldFdatahShift = 0
	RegisterM4fecrFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *RegisterM4fecrType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4fecrFieldFdatahMask) >> RegisterM4fecrFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *RegisterM4fecrType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4fecrFieldFdatahMask)|(uint32(value)<<RegisterM4fecrFieldFdatahShift))
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
	RegisterM5crFieldFecShift = 0
	RegisterM5crFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *RegisterM5crType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5crFieldFecMask) >> RegisterM5crFieldFecShift)
}

// SetFec Failing error code
func (r *RegisterM5crType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5crFieldFecMask)|(uint32(value)<<RegisterM5crFieldFecShift))
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
	RegisterM5srFieldFecShift = 0
	RegisterM5srFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *RegisterM5srType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5srFieldFecMask) >> RegisterM5srFieldFecShift)
}

// SetFec Failing error code
func (r *RegisterM5srType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5srFieldFecMask)|(uint32(value)<<RegisterM5srFieldFecShift))
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
	RegisterM5farFieldFecShift = 0
	RegisterM5farFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *RegisterM5farType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5farFieldFecMask) >> RegisterM5farFieldFecShift)
}

// SetFec Failing error code
func (r *RegisterM5farType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5farFieldFecMask)|(uint32(value)<<RegisterM5farFieldFecShift))
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
	RegisterM5fdrlFieldFecShift = 0
	RegisterM5fdrlFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *RegisterM5fdrlType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5fdrlFieldFecMask) >> RegisterM5fdrlFieldFecShift)
}

// SetFec Failing error code
func (r *RegisterM5fdrlType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5fdrlFieldFecMask)|(uint32(value)<<RegisterM5fdrlFieldFecShift))
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
