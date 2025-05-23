package ac

import (
	"unsafe"
	"volatile"
)

var (
	Ac = (*_ac)(unsafe.Pointer(uintptr(0xe000ef90)))
)

type _ac struct {
	Itcmcr registerItcmcrType
	Dtcmcr registerDtcmcrType
	Ahbpcr registerAhbpcrType
	Cacr   registerCacrType
	Ahbscr registerAhbscrType
	_      [4]byte
	Abfsr  registerAbfsrType
}

// registerItcmcrType Instruction and Data Tightly-Coupled Memory Control Registers
type registerItcmcrType uint32

const (
	RegisterItcmcrFieldEnShift = 0
	RegisterItcmcrFieldEnMask  = 0x1
)

// GetEn EN
func (r *registerItcmcrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterItcmcrFieldEnMask) != 0
}

// SetEn EN
func (r *registerItcmcrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterItcmcrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterItcmcrFieldEnMask)
	}
}

const (
	RegisterItcmcrFieldRmwShift = 1
	RegisterItcmcrFieldRmwMask  = 0x2
)

// GetRmw RMW
func (r *registerItcmcrType) GetRmw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterItcmcrFieldRmwMask) != 0
}

// SetRmw RMW
func (r *registerItcmcrType) SetRmw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterItcmcrFieldRmwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterItcmcrFieldRmwMask)
	}
}

const (
	RegisterItcmcrFieldRetenShift = 2
	RegisterItcmcrFieldRetenMask  = 0x4
)

// GetReten RETEN
func (r *registerItcmcrType) GetReten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterItcmcrFieldRetenMask) != 0
}

// SetReten RETEN
func (r *registerItcmcrType) SetReten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterItcmcrFieldRetenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterItcmcrFieldRetenMask)
	}
}

const (
	RegisterItcmcrFieldSzShift = 3
	RegisterItcmcrFieldSzMask  = 0x78
)

// GetSz SZ
func (r *registerItcmcrType) GetSz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterItcmcrFieldSzMask) >> RegisterItcmcrFieldSzShift)
}

// SetSz SZ
func (r *registerItcmcrType) SetSz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterItcmcrFieldSzMask)|(uint32(value)<<RegisterItcmcrFieldSzShift))
}

// registerDtcmcrType Instruction and Data Tightly-Coupled Memory Control Registers
type registerDtcmcrType uint32

const (
	RegisterDtcmcrFieldEnShift = 0
	RegisterDtcmcrFieldEnMask  = 0x1
)

// GetEn EN
func (r *registerDtcmcrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtcmcrFieldEnMask) != 0
}

// SetEn EN
func (r *registerDtcmcrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtcmcrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtcmcrFieldEnMask)
	}
}

const (
	RegisterDtcmcrFieldRmwShift = 1
	RegisterDtcmcrFieldRmwMask  = 0x2
)

// GetRmw RMW
func (r *registerDtcmcrType) GetRmw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtcmcrFieldRmwMask) != 0
}

// SetRmw RMW
func (r *registerDtcmcrType) SetRmw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtcmcrFieldRmwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtcmcrFieldRmwMask)
	}
}

const (
	RegisterDtcmcrFieldRetenShift = 2
	RegisterDtcmcrFieldRetenMask  = 0x4
)

// GetReten RETEN
func (r *registerDtcmcrType) GetReten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtcmcrFieldRetenMask) != 0
}

// SetReten RETEN
func (r *registerDtcmcrType) SetReten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtcmcrFieldRetenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtcmcrFieldRetenMask)
	}
}

const (
	RegisterDtcmcrFieldSzShift = 3
	RegisterDtcmcrFieldSzMask  = 0x78
)

// GetSz SZ
func (r *registerDtcmcrType) GetSz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDtcmcrFieldSzMask) >> RegisterDtcmcrFieldSzShift)
}

// SetSz SZ
func (r *registerDtcmcrType) SetSz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtcmcrFieldSzMask)|(uint32(value)<<RegisterDtcmcrFieldSzShift))
}

// registerAhbpcrType AHBP Control register
type registerAhbpcrType uint32

const (
	RegisterAhbpcrFieldEnShift = 0
	RegisterAhbpcrFieldEnMask  = 0x1
)

// GetEn EN
func (r *registerAhbpcrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhbpcrFieldEnMask) != 0
}

// SetEn EN
func (r *registerAhbpcrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhbpcrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhbpcrFieldEnMask)
	}
}

const (
	RegisterAhbpcrFieldSzShift = 1
	RegisterAhbpcrFieldSzMask  = 0xe
)

// GetSz SZ
func (r *registerAhbpcrType) GetSz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAhbpcrFieldSzMask) >> RegisterAhbpcrFieldSzShift)
}

// SetSz SZ
func (r *registerAhbpcrType) SetSz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAhbpcrFieldSzMask)|(uint32(value)<<RegisterAhbpcrFieldSzShift))
}

// registerCacrType Auxiliary Cache Control register
type registerCacrType uint32

const (
	RegisterCacrFieldSiwtShift = 0
	RegisterCacrFieldSiwtMask  = 0x1
)

// GetSiwt SIWT
func (r *registerCacrType) GetSiwt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCacrFieldSiwtMask) != 0
}

// SetSiwt SIWT
func (r *registerCacrType) SetSiwt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCacrFieldSiwtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCacrFieldSiwtMask)
	}
}

const (
	RegisterCacrFieldEccenShift = 1
	RegisterCacrFieldEccenMask  = 0x2
)

// GetEccen ECCEN
func (r *registerCacrType) GetEccen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCacrFieldEccenMask) != 0
}

// SetEccen ECCEN
func (r *registerCacrType) SetEccen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCacrFieldEccenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCacrFieldEccenMask)
	}
}

const (
	RegisterCacrFieldForcewtShift = 2
	RegisterCacrFieldForcewtMask  = 0x4
)

// GetForcewt FORCEWT
func (r *registerCacrType) GetForcewt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCacrFieldForcewtMask) != 0
}

// SetForcewt FORCEWT
func (r *registerCacrType) SetForcewt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCacrFieldForcewtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCacrFieldForcewtMask)
	}
}

// registerAhbscrType AHB Slave Control register
type registerAhbscrType uint32

const (
	RegisterAhbscrFieldCtlShift = 0
	RegisterAhbscrFieldCtlMask  = 0x3
)

// GetCtl CTL
func (r *registerAhbscrType) GetCtl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAhbscrFieldCtlMask) >> RegisterAhbscrFieldCtlShift)
}

// SetCtl CTL
func (r *registerAhbscrType) SetCtl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAhbscrFieldCtlMask)|(uint32(value)<<RegisterAhbscrFieldCtlShift))
}

const (
	RegisterAhbscrFieldTpriShift = 2
	RegisterAhbscrFieldTpriMask  = 0x7fc
)

// GetTpri TPRI
func (r *registerAhbscrType) GetTpri() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterAhbscrFieldTpriMask) >> RegisterAhbscrFieldTpriShift)
}

// SetTpri TPRI
func (r *registerAhbscrType) SetTpri(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAhbscrFieldTpriMask)|(uint32(value)<<RegisterAhbscrFieldTpriShift))
}

const (
	RegisterAhbscrFieldInitcountShift = 11
	RegisterAhbscrFieldInitcountMask  = 0xf800
)

// GetInitcount INITCOUNT
func (r *registerAhbscrType) GetInitcount() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAhbscrFieldInitcountMask) >> RegisterAhbscrFieldInitcountShift)
}

// SetInitcount INITCOUNT
func (r *registerAhbscrType) SetInitcount(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAhbscrFieldInitcountMask)|(uint32(value)<<RegisterAhbscrFieldInitcountShift))
}

// registerAbfsrType Auxiliary Bus Fault Status register
type registerAbfsrType uint32

const (
	RegisterAbfsrFieldItcmShift = 0
	RegisterAbfsrFieldItcmMask  = 0x1
)

// GetItcm ITCM
func (r *registerAbfsrType) GetItcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAbfsrFieldItcmMask) != 0
}

// SetItcm ITCM
func (r *registerAbfsrType) SetItcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAbfsrFieldItcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAbfsrFieldItcmMask)
	}
}

const (
	RegisterAbfsrFieldDtcmShift = 1
	RegisterAbfsrFieldDtcmMask  = 0x2
)

// GetDtcm DTCM
func (r *registerAbfsrType) GetDtcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAbfsrFieldDtcmMask) != 0
}

// SetDtcm DTCM
func (r *registerAbfsrType) SetDtcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAbfsrFieldDtcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAbfsrFieldDtcmMask)
	}
}

const (
	RegisterAbfsrFieldAhbpShift = 2
	RegisterAbfsrFieldAhbpMask  = 0x4
)

// GetAhbp AHBP
func (r *registerAbfsrType) GetAhbp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAbfsrFieldAhbpMask) != 0
}

// SetAhbp AHBP
func (r *registerAbfsrType) SetAhbp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAbfsrFieldAhbpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAbfsrFieldAhbpMask)
	}
}

const (
	RegisterAbfsrFieldAximShift = 3
	RegisterAbfsrFieldAximMask  = 0x8
)

// GetAxim AXIM
func (r *registerAbfsrType) GetAxim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAbfsrFieldAximMask) != 0
}

// SetAxim AXIM
func (r *registerAbfsrType) SetAxim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAbfsrFieldAximMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAbfsrFieldAximMask)
	}
}

const (
	RegisterAbfsrFieldEppbShift = 4
	RegisterAbfsrFieldEppbMask  = 0x10
)

// GetEppb EPPB
func (r *registerAbfsrType) GetEppb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAbfsrFieldEppbMask) != 0
}

// SetEppb EPPB
func (r *registerAbfsrType) SetEppb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAbfsrFieldEppbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAbfsrFieldEppbMask)
	}
}

const (
	RegisterAbfsrFieldAximtypeShift = 8
	RegisterAbfsrFieldAximtypeMask  = 0x300
)

// GetAximtype AXIMTYPE
func (r *registerAbfsrType) GetAximtype() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAbfsrFieldAximtypeMask) >> RegisterAbfsrFieldAximtypeShift)
}

// SetAximtype AXIMTYPE
func (r *registerAbfsrType) SetAximtype(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAbfsrFieldAximtypeMask)|(uint32(value)<<RegisterAbfsrFieldAximtypeShift))
}
