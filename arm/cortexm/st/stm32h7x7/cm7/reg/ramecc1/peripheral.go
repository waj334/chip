package ramecc1

import (
	"unsafe"
	"volatile"
)

var (
	Ramecc1 = (*_ramecc1)(unsafe.Pointer(uintptr(0x52009000)))
)

type _ramecc1 struct {
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
	RegisterM1srFieldEccseieShift = 2
	RegisterM1srFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *registerM1srType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *registerM1srType) SetEccseie(value bool) {
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
func (r *registerM1srType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *registerM1srType) SetEccdeie(value bool) {
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
func (r *registerM1srType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *registerM1srType) SetEccdebwie(value bool) {
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
func (r *registerM1srType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1srFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *registerM1srType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1srFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1srFieldEccelenMask)
	}
}

// registerM1farType RAMECC monitor x failing address register
type registerM1farType uint32

const (
	RegisterM1farFieldEccseieShift = 2
	RegisterM1farFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *registerM1farType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1farFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *registerM1farType) SetEccseie(value bool) {
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
func (r *registerM1farType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1farFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *registerM1farType) SetEccdeie(value bool) {
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
func (r *registerM1farType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1farFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *registerM1farType) SetEccdebwie(value bool) {
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
func (r *registerM1farType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1farFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *registerM1farType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1farFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1farFieldEccelenMask)
	}
}

// registerM1fdrlType RAMECC monitor x failing data low register
type registerM1fdrlType uint32

const (
	RegisterM1fdrlFieldEccseieShift = 2
	RegisterM1fdrlFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *registerM1fdrlType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrlFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *registerM1fdrlType) SetEccseie(value bool) {
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
func (r *registerM1fdrlType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrlFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *registerM1fdrlType) SetEccdeie(value bool) {
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
func (r *registerM1fdrlType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrlFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *registerM1fdrlType) SetEccdebwie(value bool) {
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
func (r *registerM1fdrlType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrlFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *registerM1fdrlType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fdrlFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrlFieldEccelenMask)
	}
}

// registerM1fdrhType RAMECC monitor x failing data high register
type registerM1fdrhType uint32

const (
	RegisterM1fdrhFieldEccseieShift = 2
	RegisterM1fdrhFieldEccseieMask  = 0x4
)

// GetEccseie ECC single error interrupt enable
func (r *registerM1fdrhType) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrhFieldEccseieMask) != 0
}

// SetEccseie ECC single error interrupt enable
func (r *registerM1fdrhType) SetEccseie(value bool) {
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
func (r *registerM1fdrhType) GetEccdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrhFieldEccdeieMask) != 0
}

// SetEccdeie ECC double error interrupt enable
func (r *registerM1fdrhType) SetEccdeie(value bool) {
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
func (r *registerM1fdrhType) GetEccdebwie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrhFieldEccdebwieMask) != 0
}

// SetEccdebwie ECC double error on byte write (BW) interrupt enable
func (r *registerM1fdrhType) SetEccdebwie(value bool) {
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
func (r *registerM1fdrhType) GetEccelen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fdrhFieldEccelenMask) != 0
}

// SetEccelen ECC error latching enable
func (r *registerM1fdrhType) SetEccelen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fdrhFieldEccelenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fdrhFieldEccelenMask)
	}
}

// registerM1fecrType RAMECC monitor x failing ECC error code register
type registerM1fecrType uint32

const (
	RegisterM1fecrFieldSedcfShift = 0
	RegisterM1fecrFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *registerM1fecrType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fecrFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *registerM1fecrType) SetSedcf(value bool) {
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
func (r *registerM1fecrType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fecrFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *registerM1fecrType) SetDedf(value bool) {
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
func (r *registerM1fecrType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM1fecrFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *registerM1fecrType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM1fecrFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM1fecrFieldDebwdfMask)
	}
}

// registerM2crType RAMECC monitor x configuration register
type registerM2crType uint32

const (
	RegisterM2crFieldSedcfShift = 0
	RegisterM2crFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *registerM2crType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *registerM2crType) SetSedcf(value bool) {
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
func (r *registerM2crType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *registerM2crType) SetDedf(value bool) {
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
func (r *registerM2crType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2crFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *registerM2crType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2crFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2crFieldDebwdfMask)
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
	RegisterM2farFieldSedcfShift = 0
	RegisterM2farFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *registerM2farType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2farFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *registerM2farType) SetSedcf(value bool) {
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
func (r *registerM2farType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2farFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *registerM2farType) SetDedf(value bool) {
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
func (r *registerM2farType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2farFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *registerM2farType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2farFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2farFieldDebwdfMask)
	}
}

// registerM2fdrlType RAMECC monitor x failing data low register
type registerM2fdrlType uint32

const (
	RegisterM2fdrlFieldSedcfShift = 0
	RegisterM2fdrlFieldSedcfMask  = 0x1
)

// GetSedcf ECC single error detected and corrected flag
func (r *registerM2fdrlType) GetSedcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2fdrlFieldSedcfMask) != 0
}

// SetSedcf ECC single error detected and corrected flag
func (r *registerM2fdrlType) SetSedcf(value bool) {
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
func (r *registerM2fdrlType) GetDedf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2fdrlFieldDedfMask) != 0
}

// SetDedf ECC double error detected flag
func (r *registerM2fdrlType) SetDedf(value bool) {
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
func (r *registerM2fdrlType) GetDebwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterM2fdrlFieldDebwdfMask) != 0
}

// SetDebwdf ECC double error on byte write (BW) detected flag
func (r *registerM2fdrlType) SetDebwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterM2fdrlFieldDebwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterM2fdrlFieldDebwdfMask)
	}
}

// registerM2fdrhType RAMECC monitor x failing data high register
type registerM2fdrhType uint32

const (
	RegisterM2fdrhFieldFaddShift = 0
	RegisterM2fdrhFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *registerM2fdrhType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM2fdrhFieldFaddMask) >> RegisterM2fdrhFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *registerM2fdrhType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM2fdrhFieldFaddMask)|(uint32(value)<<RegisterM2fdrhFieldFaddShift))
}

// registerM2fecrType RAMECC monitor x failing ECC error code register
type registerM2fecrType uint32

const (
	RegisterM2fecrFieldFaddShift = 0
	RegisterM2fecrFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *registerM2fecrType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM2fecrFieldFaddMask) >> RegisterM2fecrFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *registerM2fecrType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM2fecrFieldFaddMask)|(uint32(value)<<RegisterM2fecrFieldFaddShift))
}

// registerM3crType RAMECC monitor x configuration register
type registerM3crType uint32

const (
	RegisterM3crFieldFaddShift = 0
	RegisterM3crFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *registerM3crType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3crFieldFaddMask) >> RegisterM3crFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *registerM3crType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3crFieldFaddMask)|(uint32(value)<<RegisterM3crFieldFaddShift))
}

// registerM3srType RAMECC monitor x status register
type registerM3srType uint32

const (
	RegisterM3srFieldFaddShift = 0
	RegisterM3srFieldFaddMask  = 0xffffffff
)

// GetFadd ECC error failing address
func (r *registerM3srType) GetFadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3srFieldFaddMask) >> RegisterM3srFieldFaddShift)
}

// SetFadd ECC error failing address
func (r *registerM3srType) SetFadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3srFieldFaddMask)|(uint32(value)<<RegisterM3srFieldFaddShift))
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
	RegisterM3fdrhFieldFdatalShift = 0
	RegisterM3fdrhFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *registerM3fdrhType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3fdrhFieldFdatalMask) >> RegisterM3fdrhFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *registerM3fdrhType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3fdrhFieldFdatalMask)|(uint32(value)<<RegisterM3fdrhFieldFdatalShift))
}

// registerM3fecrType RAMECC monitor x failing ECC error code register
type registerM3fecrType uint32

const (
	RegisterM3fecrFieldFdatalShift = 0
	RegisterM3fecrFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *registerM3fecrType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM3fecrFieldFdatalMask) >> RegisterM3fecrFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *registerM3fecrType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM3fecrFieldFdatalMask)|(uint32(value)<<RegisterM3fecrFieldFdatalShift))
}

// registerM4crType RAMECC monitor x configuration register
type registerM4crType uint32

const (
	RegisterM4crFieldFdatalShift = 0
	RegisterM4crFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *registerM4crType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4crFieldFdatalMask) >> RegisterM4crFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *registerM4crType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4crFieldFdatalMask)|(uint32(value)<<RegisterM4crFieldFdatalShift))
}

// registerM4srType RAMECC monitor x status register
type registerM4srType uint32

const (
	RegisterM4srFieldFdatalShift = 0
	RegisterM4srFieldFdatalMask  = 0xffffffff
)

// GetFdatal Failing data low
func (r *registerM4srType) GetFdatal() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4srFieldFdatalMask) >> RegisterM4srFieldFdatalShift)
}

// SetFdatal Failing data low
func (r *registerM4srType) SetFdatal(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4srFieldFdatalMask)|(uint32(value)<<RegisterM4srFieldFdatalShift))
}

// registerM4farType RAMECC monitor x failing address register
type registerM4farType uint32

const (
	RegisterM4farFieldFdatahShift = 0
	RegisterM4farFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *registerM4farType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4farFieldFdatahMask) >> RegisterM4farFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *registerM4farType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4farFieldFdatahMask)|(uint32(value)<<RegisterM4farFieldFdatahShift))
}

// registerM4fdrlType RAMECC monitor x failing data low register
type registerM4fdrlType uint32

const (
	RegisterM4fdrlFieldFdatahShift = 0
	RegisterM4fdrlFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *registerM4fdrlType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4fdrlFieldFdatahMask) >> RegisterM4fdrlFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *registerM4fdrlType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4fdrlFieldFdatahMask)|(uint32(value)<<RegisterM4fdrlFieldFdatahShift))
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
	RegisterM4fecrFieldFdatahShift = 0
	RegisterM4fecrFieldFdatahMask  = 0xffffffff
)

// GetFdatah Failing data high (64-bit memory)
func (r *registerM4fecrType) GetFdatah() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM4fecrFieldFdatahMask) >> RegisterM4fecrFieldFdatahShift)
}

// SetFdatah Failing data high (64-bit memory)
func (r *registerM4fecrType) SetFdatah(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM4fecrFieldFdatahMask)|(uint32(value)<<RegisterM4fecrFieldFdatahShift))
}

// registerM5crType RAMECC monitor x configuration register
type registerM5crType uint32

const (
	RegisterM5crFieldFecShift = 0
	RegisterM5crFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *registerM5crType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5crFieldFecMask) >> RegisterM5crFieldFecShift)
}

// SetFec Failing error code
func (r *registerM5crType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5crFieldFecMask)|(uint32(value)<<RegisterM5crFieldFecShift))
}

// registerM5srType RAMECC monitor x status register
type registerM5srType uint32

const (
	RegisterM5srFieldFecShift = 0
	RegisterM5srFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *registerM5srType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5srFieldFecMask) >> RegisterM5srFieldFecShift)
}

// SetFec Failing error code
func (r *registerM5srType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5srFieldFecMask)|(uint32(value)<<RegisterM5srFieldFecShift))
}

// registerM5farType RAMECC monitor x failing address register
type registerM5farType uint32

const (
	RegisterM5farFieldFecShift = 0
	RegisterM5farFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *registerM5farType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5farFieldFecMask) >> RegisterM5farFieldFecShift)
}

// SetFec Failing error code
func (r *registerM5farType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5farFieldFecMask)|(uint32(value)<<RegisterM5farFieldFecShift))
}

// registerM5fdrlType RAMECC monitor x failing data low register
type registerM5fdrlType uint32

const (
	RegisterM5fdrlFieldFecShift = 0
	RegisterM5fdrlFieldFecMask  = 0xffffffff
)

// GetFec Failing error code
func (r *registerM5fdrlType) GetFec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterM5fdrlFieldFecMask) >> RegisterM5fdrlFieldFecShift)
}

// SetFec Failing error code
func (r *registerM5fdrlType) SetFec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterM5fdrlFieldFecMask)|(uint32(value)<<RegisterM5fdrlFieldFecShift))
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
