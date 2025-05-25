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
	Ier    registerIerType
	_      [28]byte
	M1cr   registerM1crType
	M1sr   registerM1srType
	M1far  registerM1farType
	M1fdrl registerM1fdrlType
	M1fdrh registerM1fdrhType
	M1fecr registerM1fecrType
	_      [8]byte
	M2cr   registerM2crType
	M2sr   registerM2srType
	M2far  registerM2farType
	M2fdrl registerM2fdrlType
	M2fdrh registerM2fdrhType
	_      [4]byte
	M2fecr registerM2fecrType
	_      [4]byte
	M3cr   registerM3crType
	M3sr   registerM3srType
	M3far  registerM3farType
	M3fdrl registerM3fdrlType
	M3fdrh registerM3fdrhType
	_      [8]byte
	M3fecr registerM3fecrType
	M4cr   registerM4crType
	M4sr   registerM4srType
	M4far  registerM4farType
	M4fdrl registerM4fdrlType
	M4fdrh registerM4fdrhType
	M4fecr registerM4fecrType
	_      [12]byte
	M5cr   registerM5crType
	M5sr   registerM5srType
	M5far  registerM5farType
	M5fdrl registerM5fdrlType
	M5fdrh registerM5fdrhType
	M5fecr registerM5fecrType
}

// registerIerType RAMECC interrupt enable register
type registerIerType uint32

const (
	RegisterIerFieldGieShift = 0
	RegisterIerFieldGieMask  = 0x1
)

// GetGie Global interrupt enable
func (r *registerIerType) GetGie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldGieMask) != 0
}

// SetGie Global interrupt enable
func (r *registerIerType) SetGie(value bool) {
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
func (r *registerIerType) GetGeccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldGeccseieMask) != 0
}

// SetGeccseie Global ECC single error interrupt enable
func (r *registerIerType) SetGeccseie(value bool) {
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
func (r *registerIerType) GetGeccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldGeccdeieMask) != 0
}

// SetGeccdeie Global ECC double error interrupt enable
func (r *registerIerType) SetGeccdeie(value bool) {
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
func (r *registerIerType) GetGeccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldGeccdebwieMask) != 0
}

// SetGeccdebwie Global ECC double error on byte write (BW) interrupt enable
func (r *registerIerType) SetGeccdebwie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldGeccdebwieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldGeccdebwieMask)
	}
}

// registerM1crType RAMECC monitor x configuration register
type registerM1crType uint32

const (
	RegisterM1crFieldEccseieShift = 2
	RegisterM1crFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *registerM1crType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1crFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *registerM1crType) SetEccseie(value bool) {
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
func (r *registerM1crType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1crFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *registerM1crType) SetEccdeie(value bool) {
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
func (r *registerM1crType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1crFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *registerM1crType) SetEccdebwie(value bool) {
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
func (r *registerM1crType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1crFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *registerM1crType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1crFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1crFieldEccelenMask)
	}
}

// registerM1srType RAMECC monitor x status register
type registerM1srType uint32

const (
	RegisterM1srFieldSedcfShift = 0
	RegisterM1srFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *registerM1srType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *registerM1srType) SetSedcf(value bool) {
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
func (r *registerM1srType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *registerM1srType) SetDedf(value bool) {
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
func (r *registerM1srType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *registerM1srType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1srFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1srFieldDebwdfMask)
	}
}

// registerM1farType RAMECC monitor x failing address register
type registerM1farType uint32

const (
	RegisterM1farFieldFaddShift = 0
	RegisterM1farFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *registerM1farType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM1farFieldFaddMask) >> RegisterM1farFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *registerM1farType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM1farFieldFaddMask)|(uint32(value)<<RegisterM1farFieldFaddShift))
}

// registerM1fdrlType RAMECC monitor x failing data low register
type registerM1fdrlType uint32

const (
	RegisterM1fdrlFieldFdatalShift = 0
	RegisterM1fdrlFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *registerM1fdrlType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrlFieldFdatalMask) >> RegisterM1fdrlFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *registerM1fdrlType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrlFieldFdatalMask)|(uint32(value)<<RegisterM1fdrlFieldFdatalShift))
}

// registerM1fdrhType RAMECC monitor x failing data high register
type registerM1fdrhType uint32

const (
	RegisterM1fdrhFieldFdatahShift = 0
	RegisterM1fdrhFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *registerM1fdrhType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrhFieldFdatahMask) >> RegisterM1fdrhFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *registerM1fdrhType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrhFieldFdatahMask)|(uint32(value)<<RegisterM1fdrhFieldFdatahShift))
}

// registerM1fecrType RAMECC monitor x failing ECC error code register
type registerM1fecrType uint32

const (
	RegisterM1fecrFieldFecShift = 0
	RegisterM1fecrFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *registerM1fecrType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM1fecrFieldFecMask) >> RegisterM1fecrFieldFecShift)
}

// SetFec Failing error code
func (r *registerM1fecrType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM1fecrFieldFecMask)|(uint32(value)<<RegisterM1fecrFieldFecShift))
}

// registerM2crType RAMECC monitor x configuration register
type registerM2crType uint32

const (
	RegisterM2crFieldEccseieShift = 2
	RegisterM2crFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *registerM2crType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *registerM2crType) SetEccseie(value bool) {
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
func (r *registerM2crType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *registerM2crType) SetEccdeie(value bool) {
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
func (r *registerM2crType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *registerM2crType) SetEccdebwie(value bool) {
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
func (r *registerM2crType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *registerM2crType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2crFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2crFieldEccelenMask)
	}
}

// registerM2srType RAMECC monitor x status register
type registerM2srType uint32

const (
	RegisterM2srFieldSedcfShift = 0
	RegisterM2srFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *registerM2srType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2srFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *registerM2srType) SetSedcf(value bool) {
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
func (r *registerM2srType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2srFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *registerM2srType) SetDedf(value bool) {
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
func (r *registerM2srType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2srFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *registerM2srType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2srFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2srFieldDebwdfMask)
	}
}

// registerM2farType RAMECC monitor x failing address register
type registerM2farType uint32

const (
	RegisterM2farFieldFaddShift = 0
	RegisterM2farFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *registerM2farType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM2farFieldFaddMask) >> RegisterM2farFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *registerM2farType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM2farFieldFaddMask)|(uint32(value)<<RegisterM2farFieldFaddShift))
}

// registerM2fdrlType RAMECC monitor x failing data low register
type registerM2fdrlType uint32

const (
	RegisterM2fdrlFieldFdatalShift = 0
	RegisterM2fdrlFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *registerM2fdrlType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM2fdrlFieldFdatalMask) >> RegisterM2fdrlFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *registerM2fdrlType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM2fdrlFieldFdatalMask)|(uint32(value)<<RegisterM2fdrlFieldFdatalShift))
}

// registerM2fdrhType RAMECC monitor x failing data high register
type registerM2fdrhType uint32

const (
	RegisterM2fdrhFieldFdatahShift = 0
	RegisterM2fdrhFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *registerM2fdrhType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM2fdrhFieldFdatahMask) >> RegisterM2fdrhFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *registerM2fdrhType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM2fdrhFieldFdatahMask)|(uint32(value)<<RegisterM2fdrhFieldFdatahShift))
}

// registerM2fecrType RAMECC monitor x failing ECC error code register
type registerM2fecrType uint32

const (
	RegisterM2fecrFieldFecShift = 0
	RegisterM2fecrFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *registerM2fecrType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM2fecrFieldFecMask) >> RegisterM2fecrFieldFecShift)
}

// SetFec Failing error code
func (r *registerM2fecrType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM2fecrFieldFecMask)|(uint32(value)<<RegisterM2fecrFieldFecShift))
}

// registerM3crType RAMECC monitor x configuration register
type registerM3crType uint32

const (
	RegisterM3crFieldEccseieShift = 2
	RegisterM3crFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *registerM3crType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3crFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *registerM3crType) SetEccseie(value bool) {
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
func (r *registerM3crType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3crFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *registerM3crType) SetEccdeie(value bool) {
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
func (r *registerM3crType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3crFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *registerM3crType) SetEccdebwie(value bool) {
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
func (r *registerM3crType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3crFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *registerM3crType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM3crFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM3crFieldEccelenMask)
	}
}

// registerM3srType RAMECC monitor x status register
type registerM3srType uint32

const (
	RegisterM3srFieldSedcfShift = 0
	RegisterM3srFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *registerM3srType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3srFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *registerM3srType) SetSedcf(value bool) {
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
func (r *registerM3srType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3srFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *registerM3srType) SetDedf(value bool) {
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
func (r *registerM3srType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM3srFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *registerM3srType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM3srFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM3srFieldDebwdfMask)
	}
}

// registerM3farType RAMECC monitor x failing address register
type registerM3farType uint32

const (
	RegisterM3farFieldFaddShift = 0
	RegisterM3farFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *registerM3farType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3farFieldFaddMask) >> RegisterM3farFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *registerM3farType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3farFieldFaddMask)|(uint32(value)<<RegisterM3farFieldFaddShift))
}

// registerM3fdrlType RAMECC monitor x failing data low register
type registerM3fdrlType uint32

const (
	RegisterM3fdrlFieldFdatalShift = 0
	RegisterM3fdrlFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *registerM3fdrlType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3fdrlFieldFdatalMask) >> RegisterM3fdrlFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *registerM3fdrlType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3fdrlFieldFdatalMask)|(uint32(value)<<RegisterM3fdrlFieldFdatalShift))
}

// registerM3fdrhType RAMECC monitor x failing data high register
type registerM3fdrhType uint32

const (
	RegisterM3fdrhFieldFdatahShift = 0
	RegisterM3fdrhFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *registerM3fdrhType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3fdrhFieldFdatahMask) >> RegisterM3fdrhFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *registerM3fdrhType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3fdrhFieldFdatahMask)|(uint32(value)<<RegisterM3fdrhFieldFdatahShift))
}

// registerM3fecrType RAMECC monitor x failing ECC error code register
type registerM3fecrType uint32

const (
	RegisterM3fecrFieldFecShift = 0
	RegisterM3fecrFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *registerM3fecrType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3fecrFieldFecMask) >> RegisterM3fecrFieldFecShift)
}

// SetFec Failing error code
func (r *registerM3fecrType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3fecrFieldFecMask)|(uint32(value)<<RegisterM3fecrFieldFecShift))
}

// registerM4crType RAMECC monitor x configuration register
type registerM4crType uint32

const (
	RegisterM4crFieldEccseieShift = 2
	RegisterM4crFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *registerM4crType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4crFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *registerM4crType) SetEccseie(value bool) {
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
func (r *registerM4crType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4crFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *registerM4crType) SetEccdeie(value bool) {
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
func (r *registerM4crType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4crFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *registerM4crType) SetEccdebwie(value bool) {
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
func (r *registerM4crType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4crFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *registerM4crType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM4crFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM4crFieldEccelenMask)
	}
}

// registerM4srType RAMECC monitor x status register
type registerM4srType uint32

const (
	RegisterM4srFieldSedcfShift = 0
	RegisterM4srFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *registerM4srType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4srFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *registerM4srType) SetSedcf(value bool) {
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
func (r *registerM4srType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4srFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *registerM4srType) SetDedf(value bool) {
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
func (r *registerM4srType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM4srFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *registerM4srType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM4srFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM4srFieldDebwdfMask)
	}
}

// registerM4farType RAMECC monitor x failing address register
type registerM4farType uint32

const (
	RegisterM4farFieldFaddShift = 0
	RegisterM4farFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *registerM4farType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4farFieldFaddMask) >> RegisterM4farFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *registerM4farType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4farFieldFaddMask)|(uint32(value)<<RegisterM4farFieldFaddShift))
}

// registerM4fdrlType RAMECC monitor x failing data low register
type registerM4fdrlType uint32

const (
	RegisterM4fdrlFieldFdatalShift = 0
	RegisterM4fdrlFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *registerM4fdrlType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4fdrlFieldFdatalMask) >> RegisterM4fdrlFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *registerM4fdrlType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4fdrlFieldFdatalMask)|(uint32(value)<<RegisterM4fdrlFieldFdatalShift))
}

// registerM4fdrhType RAMECC monitor x failing data high register
type registerM4fdrhType uint32

const (
	RegisterM4fdrhFieldFdatahShift = 0
	RegisterM4fdrhFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *registerM4fdrhType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4fdrhFieldFdatahMask) >> RegisterM4fdrhFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *registerM4fdrhType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4fdrhFieldFdatahMask)|(uint32(value)<<RegisterM4fdrhFieldFdatahShift))
}

// registerM4fecrType RAMECC monitor x failing ECC error code register
type registerM4fecrType uint32

const (
	RegisterM4fecrFieldFecShift = 0
	RegisterM4fecrFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *registerM4fecrType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4fecrFieldFecMask) >> RegisterM4fecrFieldFecShift)
}

// SetFec Failing error code
func (r *registerM4fecrType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4fecrFieldFecMask)|(uint32(value)<<RegisterM4fecrFieldFecShift))
}

// registerM5crType RAMECC monitor x configuration register
type registerM5crType uint32

const (
	RegisterM5crFieldEccseieShift = 2
	RegisterM5crFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *registerM5crType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5crFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *registerM5crType) SetEccseie(value bool) {
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
func (r *registerM5crType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5crFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *registerM5crType) SetEccdeie(value bool) {
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
func (r *registerM5crType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5crFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *registerM5crType) SetEccdebwie(value bool) {
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
func (r *registerM5crType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5crFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *registerM5crType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM5crFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM5crFieldEccelenMask)
	}
}

// registerM5srType RAMECC monitor x status register
type registerM5srType uint32

const (
	RegisterM5srFieldSedcfShift = 0
	RegisterM5srFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *registerM5srType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5srFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *registerM5srType) SetSedcf(value bool) {
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
func (r *registerM5srType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5srFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *registerM5srType) SetDedf(value bool) {
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
func (r *registerM5srType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM5srFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *registerM5srType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM5srFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM5srFieldDebwdfMask)
	}
}

// registerM5farType RAMECC monitor x failing address register
type registerM5farType uint32

const (
	RegisterM5farFieldFaddShift = 0
	RegisterM5farFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *registerM5farType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5farFieldFaddMask) >> RegisterM5farFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *registerM5farType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5farFieldFaddMask)|(uint32(value)<<RegisterM5farFieldFaddShift))
}

// registerM5fdrlType RAMECC monitor x failing data low register
type registerM5fdrlType uint32

const (
	RegisterM5fdrlFieldFdatalShift = 0
	RegisterM5fdrlFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *registerM5fdrlType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5fdrlFieldFdatalMask) >> RegisterM5fdrlFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *registerM5fdrlType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5fdrlFieldFdatalMask)|(uint32(value)<<RegisterM5fdrlFieldFdatalShift))
}

// registerM5fdrhType RAMECC monitor x failing data high register
type registerM5fdrhType uint32

const (
	RegisterM5fdrhFieldFecShift = 0
	RegisterM5fdrhFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *registerM5fdrhType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5fdrhFieldFecMask) >> RegisterM5fdrhFieldFecShift)
}

// SetFec Failing error code
func (r *registerM5fdrhType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5fdrhFieldFecMask)|(uint32(value)<<RegisterM5fdrhFieldFecShift))
}

// registerM5fecrType RAMECC monitor x failing ECC error code register
type registerM5fecrType uint32

const (
	RegisterM5fecrFieldFecShift = 0
	RegisterM5fecrFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *registerM5fecrType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5fecrFieldFecMask) >> RegisterM5fecrFieldFecShift)
}

// SetFec Failing error code
func (r *registerM5fecrType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5fecrFieldFecMask)|(uint32(value)<<RegisterM5fecrFieldFecShift))
}
