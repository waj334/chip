package otg1_hs_device

import (
	"unsafe"
	"volatile"
)

var (
	Otg1_hs_device = (*_otg1_hs_device)(unsafe.Pointer(uintptr(0x40040800)))
	Otg2_hs_device = (*_otg1_hs_device)(unsafe.Pointer(uintptr(0x40080800)))
)

type _otg1_hs_device struct {
	Otg_hs_dcfg        registerOtg_hs_dcfgType
	Otg_hs_dctl        registerOtg_hs_dctlType
	Otg_hs_dsts        registerOtg_hs_dstsType
	_                  [4]byte
	Otg_hs_diepmsk     registerOtg_hs_diepmskType
	Otg_hs_doepmsk     registerOtg_hs_doepmskType
	Otg_hs_daint       registerOtg_hs_daintType
	Otg_hs_daintmsk    registerOtg_hs_daintmskType
	_                  [8]byte
	Otg_hs_dvbusdis    registerOtg_hs_dvbusdisType
	Otg_hs_dvbuspulse  registerOtg_hs_dvbuspulseType
	Otg_hs_dthrctl     registerOtg_hs_dthrctlType
	Otg_hs_diepempmsk  registerOtg_hs_diepempmskType
	Otg_hs_deachint    registerOtg_hs_deachintType
	Otg_hs_deachintmsk registerOtg_hs_deachintmskType
	_                  [192]byte
	Otg_hs_diepctl0    registerOtg_hs_diepctl0Type
	_                  [28]byte
	Otg_hs_diepctl1    registerOtg_hs_diepctl1Type
	_                  [28]byte
	Otg_hs_diepctl2    registerOtg_hs_diepctl2Type
	_                  [28]byte
	Otg_hs_diepctl3    registerOtg_hs_diepctl3Type
	_                  [28]byte
	Otg_hs_diepctl4    registerOtg_hs_diepctl4Type
	_                  [28]byte
	Otg_hs_diepctl5    registerOtg_hs_diepctl5Type
	_                  [28]byte
	Otg_hs_diepctl6    registerOtg_hs_diepctl6Type
	_                  [28]byte
	Otg_hs_diepctl7    registerOtg_hs_diepctl7Type
	Otg_hs_diepint0    registerOtg_hs_diepint0Type
	_                  [28]byte
	Otg_hs_diepint1    registerOtg_hs_diepint1Type
	_                  [28]byte
	Otg_hs_diepint2    registerOtg_hs_diepint2Type
	_                  [28]byte
	Otg_hs_diepint3    registerOtg_hs_diepint3Type
	_                  [28]byte
	Otg_hs_diepint4    registerOtg_hs_diepint4Type
	_                  [28]byte
	Otg_hs_diepint5    registerOtg_hs_diepint5Type
	_                  [28]byte
	Otg_hs_diepint6    registerOtg_hs_diepint6Type
	_                  [28]byte
	Otg_hs_diepint7    registerOtg_hs_diepint7Type
	Otg_hs_dieptsiz0   registerOtg_hs_dieptsiz0Type
	Otg_hs_diepdma1    registerOtg_hs_diepdma1Type
	_                  [28]byte
	Otg_hs_diepdma2    registerOtg_hs_diepdma2Type
	_                  [28]byte
	Otg_hs_diepdma3    registerOtg_hs_diepdma3Type
	_                  [28]byte
	Otg_hs_diepdma4    registerOtg_hs_diepdma4Type
	_                  [28]byte
	Otg_hs_diepdma5    registerOtg_hs_diepdma5Type
	Otg_hs_dtxfsts0    registerOtg_hs_dtxfsts0Type
	_                  [28]byte
	Otg_hs_dtxfsts1    registerOtg_hs_dtxfsts1Type
	_                  [28]byte
	Otg_hs_dtxfsts2    registerOtg_hs_dtxfsts2Type
	_                  [28]byte
	Otg_hs_dtxfsts3    registerOtg_hs_dtxfsts3Type
	_                  [28]byte
	Otg_hs_dtxfsts4    registerOtg_hs_dtxfsts4Type
	_                  [28]byte
	Otg_hs_dtxfsts5    registerOtg_hs_dtxfsts5Type
	Otg_hs_dieptsiz1   registerOtg_hs_dieptsiz1Type
	_                  [28]byte
	Otg_hs_dieptsiz2   registerOtg_hs_dieptsiz2Type
	_                  [28]byte
	Otg_hs_dieptsiz3   registerOtg_hs_dieptsiz3Type
	_                  [28]byte
	Otg_hs_dieptsiz4   registerOtg_hs_dieptsiz4Type
	_                  [28]byte
	Otg_hs_dieptsiz5   registerOtg_hs_dieptsiz5Type
	_                  [332]byte
	Otg_hs_doepctl0    registerOtg_hs_doepctl0Type
	_                  [28]byte
	Otg_hs_doepctl1    registerOtg_hs_doepctl1Type
	_                  [28]byte
	Otg_hs_doepctl2    registerOtg_hs_doepctl2Type
	_                  [28]byte
	Otg_hs_doepctl3    registerOtg_hs_doepctl3Type
	Otg_hs_doepint0    registerOtg_hs_doepint0Type
	_                  [28]byte
	Otg_hs_doepint1    registerOtg_hs_doepint1Type
	_                  [28]byte
	Otg_hs_doepint2    registerOtg_hs_doepint2Type
	_                  [28]byte
	Otg_hs_doepint3    registerOtg_hs_doepint3Type
	_                  [28]byte
	Otg_hs_doepint4    registerOtg_hs_doepint4Type
	_                  [28]byte
	Otg_hs_doepint5    registerOtg_hs_doepint5Type
	_                  [28]byte
	Otg_hs_doepint6    registerOtg_hs_doepint6Type
	_                  [28]byte
	Otg_hs_doepint7    registerOtg_hs_doepint7Type
	Otg_hs_doeptsiz0   registerOtg_hs_doeptsiz0Type
	_                  [28]byte
	Otg_hs_doeptsiz1   registerOtg_hs_doeptsiz1Type
	_                  [28]byte
	Otg_hs_doeptsiz2   registerOtg_hs_doeptsiz2Type
	_                  [28]byte
	Otg_hs_doeptsiz3   registerOtg_hs_doeptsiz3Type
	_                  [28]byte
	Otg_hs_doeptsiz4   registerOtg_hs_doeptsiz4Type
	Otg_hs_dieptsiz6   registerOtg_hs_dieptsiz6Type
	Otg_hs_dtxfsts6    registerOtg_hs_dtxfsts6Type
	Otg_hs_dieptsiz7   registerOtg_hs_dieptsiz7Type
	Otg_hs_dtxfsts7    registerOtg_hs_dtxfsts7Type
	_                  [464]byte
	Otg_hs_doepctl4    registerOtg_hs_doepctl4Type
	_                  [28]byte
	Otg_hs_doepctl5    registerOtg_hs_doepctl5Type
	_                  [28]byte
	Otg_hs_doepctl6    registerOtg_hs_doepctl6Type
	_                  [28]byte
	Otg_hs_doepctl7    registerOtg_hs_doepctl7Type
	Otg_hs_doeptsiz5   registerOtg_hs_doeptsiz5Type
	_                  [28]byte
	Otg_hs_doeptsiz6   registerOtg_hs_doeptsiz6Type
	_                  [28]byte
	Otg_hs_doeptsiz7   registerOtg_hs_doeptsiz7Type
}

// registerOtg_hs_dcfgType OTG_HS device configuration register
type registerOtg_hs_dcfgType uint32

const (
	RegisterOtg_hs_dcfgFieldDspdShift = 0
	RegisterOtg_hs_dcfgFieldDspdMask  = 0x3
)

// GetDspd Device speed
func (r *registerOtg_hs_dcfgType) GetDspd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dcfgFieldDspdMask) >> RegisterOtg_hs_dcfgFieldDspdShift)
}

// SetDspd Device speed
func (r *registerOtg_hs_dcfgType) SetDspd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dcfgFieldDspdMask)|(uint32(value)<<RegisterOtg_hs_dcfgFieldDspdShift))
}

const (
	RegisterOtg_hs_dcfgFieldNzlsohskShift = 2
	RegisterOtg_hs_dcfgFieldNzlsohskMask  = 0x4
)

// GetNzlsohsk Nonzero-length status OUT handshake
func (r *registerOtg_hs_dcfgType) GetNzlsohsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dcfgFieldNzlsohskMask) != 0
}

// SetNzlsohsk Nonzero-length status OUT handshake
func (r *registerOtg_hs_dcfgType) SetNzlsohsk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dcfgFieldNzlsohskMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dcfgFieldNzlsohskMask)
	}
}

const (
	RegisterOtg_hs_dcfgFieldDadShift = 4
	RegisterOtg_hs_dcfgFieldDadMask  = 0x7f0
)

// GetDad Device address
func (r *registerOtg_hs_dcfgType) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dcfgFieldDadMask) >> RegisterOtg_hs_dcfgFieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_dcfgType) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dcfgFieldDadMask)|(uint32(value)<<RegisterOtg_hs_dcfgFieldDadShift))
}

const (
	RegisterOtg_hs_dcfgFieldPfivlShift = 11
	RegisterOtg_hs_dcfgFieldPfivlMask  = 0x1800
)

// GetPfivl Periodic (micro)frame interval
func (r *registerOtg_hs_dcfgType) GetPfivl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dcfgFieldPfivlMask) >> RegisterOtg_hs_dcfgFieldPfivlShift)
}

// SetPfivl Periodic (micro)frame interval
func (r *registerOtg_hs_dcfgType) SetPfivl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dcfgFieldPfivlMask)|(uint32(value)<<RegisterOtg_hs_dcfgFieldPfivlShift))
}

const (
	RegisterOtg_hs_dcfgFieldPerschivlShift = 24
	RegisterOtg_hs_dcfgFieldPerschivlMask  = 0x3000000
)

// GetPerschivl Periodic scheduling interval
func (r *registerOtg_hs_dcfgType) GetPerschivl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dcfgFieldPerschivlMask) >> RegisterOtg_hs_dcfgFieldPerschivlShift)
}

// SetPerschivl Periodic scheduling interval
func (r *registerOtg_hs_dcfgType) SetPerschivl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dcfgFieldPerschivlMask)|(uint32(value)<<RegisterOtg_hs_dcfgFieldPerschivlShift))
}

// registerOtg_hs_dctlType OTG_HS device control register
type registerOtg_hs_dctlType uint32

const (
	RegisterOtg_hs_dctlFieldRwusigShift = 0
	RegisterOtg_hs_dctlFieldRwusigMask  = 0x1
)

// GetRwusig Remote wakeup signaling
func (r *registerOtg_hs_dctlType) GetRwusig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dctlFieldRwusigMask) != 0
}

// SetRwusig Remote wakeup signaling
func (r *registerOtg_hs_dctlType) SetRwusig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dctlFieldRwusigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dctlFieldRwusigMask)
	}
}

const (
	RegisterOtg_hs_dctlFieldSdisShift = 1
	RegisterOtg_hs_dctlFieldSdisMask  = 0x2
)

// GetSdis Soft disconnect
func (r *registerOtg_hs_dctlType) GetSdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dctlFieldSdisMask) != 0
}

// SetSdis Soft disconnect
func (r *registerOtg_hs_dctlType) SetSdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dctlFieldSdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dctlFieldSdisMask)
	}
}

const (
	RegisterOtg_hs_dctlFieldGinstsShift = 2
	RegisterOtg_hs_dctlFieldGinstsMask  = 0x4
)

// GetGinsts Global IN NAK status
func (r *registerOtg_hs_dctlType) GetGinsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dctlFieldGinstsMask) != 0
}

const (
	RegisterOtg_hs_dctlFieldGonstsShift = 3
	RegisterOtg_hs_dctlFieldGonstsMask  = 0x8
)

// GetGonsts Global OUT NAK status
func (r *registerOtg_hs_dctlType) GetGonsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dctlFieldGonstsMask) != 0
}

const (
	RegisterOtg_hs_dctlFieldTctlShift = 4
	RegisterOtg_hs_dctlFieldTctlMask  = 0x70
)

// GetTctl Test control
func (r *registerOtg_hs_dctlType) GetTctl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dctlFieldTctlMask) >> RegisterOtg_hs_dctlFieldTctlShift)
}

// SetTctl Test control
func (r *registerOtg_hs_dctlType) SetTctl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dctlFieldTctlMask)|(uint32(value)<<RegisterOtg_hs_dctlFieldTctlShift))
}

const (
	RegisterOtg_hs_dctlFieldSginakShift = 7
	RegisterOtg_hs_dctlFieldSginakMask  = 0x80
)

// SetSginak Set global IN NAK
func (r *registerOtg_hs_dctlType) SetSginak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dctlFieldSginakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dctlFieldSginakMask)
	}
}

const (
	RegisterOtg_hs_dctlFieldCginakShift = 8
	RegisterOtg_hs_dctlFieldCginakMask  = 0x100
)

// SetCginak Clear global IN NAK
func (r *registerOtg_hs_dctlType) SetCginak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dctlFieldCginakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dctlFieldCginakMask)
	}
}

const (
	RegisterOtg_hs_dctlFieldSgonakShift = 9
	RegisterOtg_hs_dctlFieldSgonakMask  = 0x200
)

// SetSgonak Set global OUT NAK
func (r *registerOtg_hs_dctlType) SetSgonak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dctlFieldSgonakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dctlFieldSgonakMask)
	}
}

const (
	RegisterOtg_hs_dctlFieldCgonakShift = 10
	RegisterOtg_hs_dctlFieldCgonakMask  = 0x400
)

// SetCgonak Clear global OUT NAK
func (r *registerOtg_hs_dctlType) SetCgonak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dctlFieldCgonakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dctlFieldCgonakMask)
	}
}

const (
	RegisterOtg_hs_dctlFieldPoprgdneShift = 11
	RegisterOtg_hs_dctlFieldPoprgdneMask  = 0x800
)

// GetPoprgdne Power-on programming done
func (r *registerOtg_hs_dctlType) GetPoprgdne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dctlFieldPoprgdneMask) != 0
}

// SetPoprgdne Power-on programming done
func (r *registerOtg_hs_dctlType) SetPoprgdne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dctlFieldPoprgdneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dctlFieldPoprgdneMask)
	}
}

// registerOtg_hs_dstsType OTG_HS device status register
type registerOtg_hs_dstsType uint32

const (
	RegisterOtg_hs_dstsFieldSuspstsShift = 0
	RegisterOtg_hs_dstsFieldSuspstsMask  = 0x1
)

// GetSuspsts Suspend status
func (r *registerOtg_hs_dstsType) GetSuspsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dstsFieldSuspstsMask) != 0
}

// SetSuspsts Suspend status
func (r *registerOtg_hs_dstsType) SetSuspsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dstsFieldSuspstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dstsFieldSuspstsMask)
	}
}

const (
	RegisterOtg_hs_dstsFieldEnumspdShift = 1
	RegisterOtg_hs_dstsFieldEnumspdMask  = 0x6
)

// GetEnumspd Enumerated speed
func (r *registerOtg_hs_dstsType) GetEnumspd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dstsFieldEnumspdMask) >> RegisterOtg_hs_dstsFieldEnumspdShift)
}

// SetEnumspd Enumerated speed
func (r *registerOtg_hs_dstsType) SetEnumspd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dstsFieldEnumspdMask)|(uint32(value)<<RegisterOtg_hs_dstsFieldEnumspdShift))
}

const (
	RegisterOtg_hs_dstsFieldEerrShift = 3
	RegisterOtg_hs_dstsFieldEerrMask  = 0x8
)

// GetEerr Erratic error
func (r *registerOtg_hs_dstsType) GetEerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dstsFieldEerrMask) != 0
}

// SetEerr Erratic error
func (r *registerOtg_hs_dstsType) SetEerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dstsFieldEerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dstsFieldEerrMask)
	}
}

const (
	RegisterOtg_hs_dstsFieldFnsofShift = 8
	RegisterOtg_hs_dstsFieldFnsofMask  = 0x3fff00
)

// GetFnsof Frame number of the received SOF
func (r *registerOtg_hs_dstsType) GetFnsof() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dstsFieldFnsofMask) >> RegisterOtg_hs_dstsFieldFnsofShift)
}

// SetFnsof Frame number of the received SOF
func (r *registerOtg_hs_dstsType) SetFnsof(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dstsFieldFnsofMask)|(uint32(value)<<RegisterOtg_hs_dstsFieldFnsofShift))
}

// registerOtg_hs_diepmskType OTG_HS device IN endpoint common interrupt mask register
type registerOtg_hs_diepmskType uint32

const (
	RegisterOtg_hs_diepmskFieldXfrcmShift = 0
	RegisterOtg_hs_diepmskFieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed interrupt mask
func (r *registerOtg_hs_diepmskType) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepmskFieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed interrupt mask
func (r *registerOtg_hs_diepmskType) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepmskFieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepmskFieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_diepmskFieldEpdmShift = 1
	RegisterOtg_hs_diepmskFieldEpdmMask  = 0x2
)

// GetEpdm Endpoint disabled interrupt mask
func (r *registerOtg_hs_diepmskType) GetEpdm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepmskFieldEpdmMask) != 0
}

// SetEpdm Endpoint disabled interrupt mask
func (r *registerOtg_hs_diepmskType) SetEpdm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepmskFieldEpdmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepmskFieldEpdmMask)
	}
}

const (
	RegisterOtg_hs_diepmskFieldTomShift = 3
	RegisterOtg_hs_diepmskFieldTomMask  = 0x8
)

// GetTom Timeout condition mask (nonisochronous endpoints)
func (r *registerOtg_hs_diepmskType) GetTom() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepmskFieldTomMask) != 0
}

// SetTom Timeout condition mask (nonisochronous endpoints)
func (r *registerOtg_hs_diepmskType) SetTom(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepmskFieldTomMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepmskFieldTomMask)
	}
}

const (
	RegisterOtg_hs_diepmskFieldIttxfemskShift = 4
	RegisterOtg_hs_diepmskFieldIttxfemskMask  = 0x10
)

// GetIttxfemsk IN token received when TxFIFO empty mask
func (r *registerOtg_hs_diepmskType) GetIttxfemsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepmskFieldIttxfemskMask) != 0
}

// SetIttxfemsk IN token received when TxFIFO empty mask
func (r *registerOtg_hs_diepmskType) SetIttxfemsk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepmskFieldIttxfemskMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepmskFieldIttxfemskMask)
	}
}

const (
	RegisterOtg_hs_diepmskFieldInepnmmShift = 5
	RegisterOtg_hs_diepmskFieldInepnmmMask  = 0x20
)

// GetInepnmm IN token received with EP mismatch mask
func (r *registerOtg_hs_diepmskType) GetInepnmm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepmskFieldInepnmmMask) != 0
}

// SetInepnmm IN token received with EP mismatch mask
func (r *registerOtg_hs_diepmskType) SetInepnmm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepmskFieldInepnmmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepmskFieldInepnmmMask)
	}
}

const (
	RegisterOtg_hs_diepmskFieldInepnemShift = 6
	RegisterOtg_hs_diepmskFieldInepnemMask  = 0x40
)

// GetInepnem IN endpoint NAK effective mask
func (r *registerOtg_hs_diepmskType) GetInepnem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepmskFieldInepnemMask) != 0
}

// SetInepnem IN endpoint NAK effective mask
func (r *registerOtg_hs_diepmskType) SetInepnem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepmskFieldInepnemMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepmskFieldInepnemMask)
	}
}

const (
	RegisterOtg_hs_diepmskFieldTxfurmShift = 8
	RegisterOtg_hs_diepmskFieldTxfurmMask  = 0x100
)

// GetTxfurm FIFO underrun mask
func (r *registerOtg_hs_diepmskType) GetTxfurm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepmskFieldTxfurmMask) != 0
}

// SetTxfurm FIFO underrun mask
func (r *registerOtg_hs_diepmskType) SetTxfurm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepmskFieldTxfurmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepmskFieldTxfurmMask)
	}
}

const (
	RegisterOtg_hs_diepmskFieldBimShift = 9
	RegisterOtg_hs_diepmskFieldBimMask  = 0x200
)

// GetBim BNA interrupt mask
func (r *registerOtg_hs_diepmskType) GetBim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepmskFieldBimMask) != 0
}

// SetBim BNA interrupt mask
func (r *registerOtg_hs_diepmskType) SetBim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepmskFieldBimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepmskFieldBimMask)
	}
}

// registerOtg_hs_doepmskType OTG_HS device OUT endpoint common interrupt mask register
type registerOtg_hs_doepmskType uint32

const (
	RegisterOtg_hs_doepmskFieldXfrcmShift = 0
	RegisterOtg_hs_doepmskFieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed interrupt mask
func (r *registerOtg_hs_doepmskType) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepmskFieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed interrupt mask
func (r *registerOtg_hs_doepmskType) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepmskFieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepmskFieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_doepmskFieldEpdmShift = 1
	RegisterOtg_hs_doepmskFieldEpdmMask  = 0x2
)

// GetEpdm Endpoint disabled interrupt mask
func (r *registerOtg_hs_doepmskType) GetEpdm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepmskFieldEpdmMask) != 0
}

// SetEpdm Endpoint disabled interrupt mask
func (r *registerOtg_hs_doepmskType) SetEpdm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepmskFieldEpdmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepmskFieldEpdmMask)
	}
}

const (
	RegisterOtg_hs_doepmskFieldStupmShift = 3
	RegisterOtg_hs_doepmskFieldStupmMask  = 0x8
)

// GetStupm SETUP phase done mask
func (r *registerOtg_hs_doepmskType) GetStupm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepmskFieldStupmMask) != 0
}

// SetStupm SETUP phase done mask
func (r *registerOtg_hs_doepmskType) SetStupm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepmskFieldStupmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepmskFieldStupmMask)
	}
}

const (
	RegisterOtg_hs_doepmskFieldOtepdmShift = 4
	RegisterOtg_hs_doepmskFieldOtepdmMask  = 0x10
)

// GetOtepdm OUT token received when endpoint disabled mask
func (r *registerOtg_hs_doepmskType) GetOtepdm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepmskFieldOtepdmMask) != 0
}

// SetOtepdm OUT token received when endpoint disabled mask
func (r *registerOtg_hs_doepmskType) SetOtepdm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepmskFieldOtepdmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepmskFieldOtepdmMask)
	}
}

const (
	RegisterOtg_hs_doepmskFieldB2bstupShift = 6
	RegisterOtg_hs_doepmskFieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received mask
func (r *registerOtg_hs_doepmskType) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepmskFieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received mask
func (r *registerOtg_hs_doepmskType) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepmskFieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepmskFieldB2bstupMask)
	}
}

const (
	RegisterOtg_hs_doepmskFieldOpemShift = 8
	RegisterOtg_hs_doepmskFieldOpemMask  = 0x100
)

// GetOpem OUT packet error mask
func (r *registerOtg_hs_doepmskType) GetOpem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepmskFieldOpemMask) != 0
}

// SetOpem OUT packet error mask
func (r *registerOtg_hs_doepmskType) SetOpem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepmskFieldOpemMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepmskFieldOpemMask)
	}
}

const (
	RegisterOtg_hs_doepmskFieldBoimShift = 9
	RegisterOtg_hs_doepmskFieldBoimMask  = 0x200
)

// GetBoim BNA interrupt mask
func (r *registerOtg_hs_doepmskType) GetBoim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepmskFieldBoimMask) != 0
}

// SetBoim BNA interrupt mask
func (r *registerOtg_hs_doepmskType) SetBoim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepmskFieldBoimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepmskFieldBoimMask)
	}
}

// registerOtg_hs_daintType OTG_HS device all endpoints interrupt register
type registerOtg_hs_daintType uint32

const (
	RegisterOtg_hs_daintFieldIepintShift = 0
	RegisterOtg_hs_daintFieldIepintMask  = 0xffff
)

// GetIepint IN endpoint interrupt bits
func (r *registerOtg_hs_daintType) GetIepint() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_daintFieldIepintMask) >> RegisterOtg_hs_daintFieldIepintShift)
}

// SetIepint IN endpoint interrupt bits
func (r *registerOtg_hs_daintType) SetIepint(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_daintFieldIepintMask)|(uint32(value)<<RegisterOtg_hs_daintFieldIepintShift))
}

const (
	RegisterOtg_hs_daintFieldOepintShift = 16
	RegisterOtg_hs_daintFieldOepintMask  = 0xffff0000
)

// GetOepint OUT endpoint interrupt bits
func (r *registerOtg_hs_daintType) GetOepint() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_daintFieldOepintMask) >> RegisterOtg_hs_daintFieldOepintShift)
}

// SetOepint OUT endpoint interrupt bits
func (r *registerOtg_hs_daintType) SetOepint(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_daintFieldOepintMask)|(uint32(value)<<RegisterOtg_hs_daintFieldOepintShift))
}

// registerOtg_hs_daintmskType OTG_HS all endpoints interrupt mask register
type registerOtg_hs_daintmskType uint32

const (
	RegisterOtg_hs_daintmskFieldIepmShift = 0
	RegisterOtg_hs_daintmskFieldIepmMask  = 0xffff
)

// GetIepm IN EP interrupt mask bits
func (r *registerOtg_hs_daintmskType) GetIepm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_daintmskFieldIepmMask) >> RegisterOtg_hs_daintmskFieldIepmShift)
}

// SetIepm IN EP interrupt mask bits
func (r *registerOtg_hs_daintmskType) SetIepm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_daintmskFieldIepmMask)|(uint32(value)<<RegisterOtg_hs_daintmskFieldIepmShift))
}

const (
	RegisterOtg_hs_daintmskFieldOepmShift = 16
	RegisterOtg_hs_daintmskFieldOepmMask  = 0xffff0000
)

// GetOepm OUT EP interrupt mask bits
func (r *registerOtg_hs_daintmskType) GetOepm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_daintmskFieldOepmMask) >> RegisterOtg_hs_daintmskFieldOepmShift)
}

// SetOepm OUT EP interrupt mask bits
func (r *registerOtg_hs_daintmskType) SetOepm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_daintmskFieldOepmMask)|(uint32(value)<<RegisterOtg_hs_daintmskFieldOepmShift))
}

// registerOtg_hs_dvbusdisType OTG_HS device VBUS discharge time register
type registerOtg_hs_dvbusdisType uint32

const (
	RegisterOtg_hs_dvbusdisFieldVbusdtShift = 0
	RegisterOtg_hs_dvbusdisFieldVbusdtMask  = 0xffff
)

// GetVbusdt Device VBUS discharge time
func (r *registerOtg_hs_dvbusdisType) GetVbusdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dvbusdisFieldVbusdtMask) >> RegisterOtg_hs_dvbusdisFieldVbusdtShift)
}

// SetVbusdt Device VBUS discharge time
func (r *registerOtg_hs_dvbusdisType) SetVbusdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dvbusdisFieldVbusdtMask)|(uint32(value)<<RegisterOtg_hs_dvbusdisFieldVbusdtShift))
}

// registerOtg_hs_dvbuspulseType OTG_HS device VBUS pulsing time register
type registerOtg_hs_dvbuspulseType uint32

const (
	RegisterOtg_hs_dvbuspulseFieldDvbuspShift = 0
	RegisterOtg_hs_dvbuspulseFieldDvbuspMask  = 0xfff
)

// GetDvbusp Device VBUS pulsing time
func (r *registerOtg_hs_dvbuspulseType) GetDvbusp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dvbuspulseFieldDvbuspMask) >> RegisterOtg_hs_dvbuspulseFieldDvbuspShift)
}

// SetDvbusp Device VBUS pulsing time
func (r *registerOtg_hs_dvbuspulseType) SetDvbusp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dvbuspulseFieldDvbuspMask)|(uint32(value)<<RegisterOtg_hs_dvbuspulseFieldDvbuspShift))
}

// registerOtg_hs_dthrctlType OTG_HS Device threshold control register
type registerOtg_hs_dthrctlType uint32

const (
	RegisterOtg_hs_dthrctlFieldNonisothrenShift = 0
	RegisterOtg_hs_dthrctlFieldNonisothrenMask  = 0x1
)

// GetNonisothren Nonisochronous IN endpoints threshold enable
func (r *registerOtg_hs_dthrctlType) GetNonisothren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dthrctlFieldNonisothrenMask) != 0
}

// SetNonisothren Nonisochronous IN endpoints threshold enable
func (r *registerOtg_hs_dthrctlType) SetNonisothren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dthrctlFieldNonisothrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dthrctlFieldNonisothrenMask)
	}
}

const (
	RegisterOtg_hs_dthrctlFieldIsothrenShift = 1
	RegisterOtg_hs_dthrctlFieldIsothrenMask  = 0x2
)

// GetIsothren ISO IN endpoint threshold enable
func (r *registerOtg_hs_dthrctlType) GetIsothren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dthrctlFieldIsothrenMask) != 0
}

// SetIsothren ISO IN endpoint threshold enable
func (r *registerOtg_hs_dthrctlType) SetIsothren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dthrctlFieldIsothrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dthrctlFieldIsothrenMask)
	}
}

const (
	RegisterOtg_hs_dthrctlFieldTxthrlenShift = 2
	RegisterOtg_hs_dthrctlFieldTxthrlenMask  = 0x7fc
)

// GetTxthrlen Transmit threshold length
func (r *registerOtg_hs_dthrctlType) GetTxthrlen() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dthrctlFieldTxthrlenMask) >> RegisterOtg_hs_dthrctlFieldTxthrlenShift)
}

// SetTxthrlen Transmit threshold length
func (r *registerOtg_hs_dthrctlType) SetTxthrlen(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dthrctlFieldTxthrlenMask)|(uint32(value)<<RegisterOtg_hs_dthrctlFieldTxthrlenShift))
}

const (
	RegisterOtg_hs_dthrctlFieldRxthrenShift = 16
	RegisterOtg_hs_dthrctlFieldRxthrenMask  = 0x10000
)

// GetRxthren Receive threshold enable
func (r *registerOtg_hs_dthrctlType) GetRxthren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dthrctlFieldRxthrenMask) != 0
}

// SetRxthren Receive threshold enable
func (r *registerOtg_hs_dthrctlType) SetRxthren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dthrctlFieldRxthrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dthrctlFieldRxthrenMask)
	}
}

const (
	RegisterOtg_hs_dthrctlFieldRxthrlenShift = 17
	RegisterOtg_hs_dthrctlFieldRxthrlenMask  = 0x3fe0000
)

// GetRxthrlen Receive threshold length
func (r *registerOtg_hs_dthrctlType) GetRxthrlen() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dthrctlFieldRxthrlenMask) >> RegisterOtg_hs_dthrctlFieldRxthrlenShift)
}

// SetRxthrlen Receive threshold length
func (r *registerOtg_hs_dthrctlType) SetRxthrlen(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dthrctlFieldRxthrlenMask)|(uint32(value)<<RegisterOtg_hs_dthrctlFieldRxthrlenShift))
}

const (
	RegisterOtg_hs_dthrctlFieldArpenShift = 27
	RegisterOtg_hs_dthrctlFieldArpenMask  = 0x8000000
)

// GetArpen Arbiter parking enable
func (r *registerOtg_hs_dthrctlType) GetArpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dthrctlFieldArpenMask) != 0
}

// SetArpen Arbiter parking enable
func (r *registerOtg_hs_dthrctlType) SetArpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_dthrctlFieldArpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dthrctlFieldArpenMask)
	}
}

// registerOtg_hs_diepempmskType OTG_HS device IN endpoint FIFO empty interrupt mask register
type registerOtg_hs_diepempmskType uint32

const (
	RegisterOtg_hs_diepempmskFieldIneptxfemShift = 0
	RegisterOtg_hs_diepempmskFieldIneptxfemMask  = 0xffff
)

// GetIneptxfem IN EP Tx FIFO empty interrupt mask bits
func (r *registerOtg_hs_diepempmskType) GetIneptxfem() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepempmskFieldIneptxfemMask) >> RegisterOtg_hs_diepempmskFieldIneptxfemShift)
}

// SetIneptxfem IN EP Tx FIFO empty interrupt mask bits
func (r *registerOtg_hs_diepempmskType) SetIneptxfem(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepempmskFieldIneptxfemMask)|(uint32(value)<<RegisterOtg_hs_diepempmskFieldIneptxfemShift))
}

// registerOtg_hs_deachintType OTG_HS device each endpoint interrupt register
type registerOtg_hs_deachintType uint32

const (
	RegisterOtg_hs_deachintFieldIep1intShift = 1
	RegisterOtg_hs_deachintFieldIep1intMask  = 0x2
)

// GetIep1int IN endpoint 1interrupt bit
func (r *registerOtg_hs_deachintType) GetIep1int() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_deachintFieldIep1intMask) != 0
}

// SetIep1int IN endpoint 1interrupt bit
func (r *registerOtg_hs_deachintType) SetIep1int(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_deachintFieldIep1intMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_deachintFieldIep1intMask)
	}
}

const (
	RegisterOtg_hs_deachintFieldOep1intShift = 17
	RegisterOtg_hs_deachintFieldOep1intMask  = 0x20000
)

// GetOep1int OUT endpoint 1 interrupt bit
func (r *registerOtg_hs_deachintType) GetOep1int() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_deachintFieldOep1intMask) != 0
}

// SetOep1int OUT endpoint 1 interrupt bit
func (r *registerOtg_hs_deachintType) SetOep1int(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_deachintFieldOep1intMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_deachintFieldOep1intMask)
	}
}

// registerOtg_hs_deachintmskType OTG_HS device each endpoint interrupt register mask
type registerOtg_hs_deachintmskType uint32

const (
	RegisterOtg_hs_deachintmskFieldIep1intmShift = 1
	RegisterOtg_hs_deachintmskFieldIep1intmMask  = 0x2
)

// GetIep1intm IN Endpoint 1 interrupt mask bit
func (r *registerOtg_hs_deachintmskType) GetIep1intm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_deachintmskFieldIep1intmMask) != 0
}

// SetIep1intm IN Endpoint 1 interrupt mask bit
func (r *registerOtg_hs_deachintmskType) SetIep1intm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_deachintmskFieldIep1intmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_deachintmskFieldIep1intmMask)
	}
}

const (
	RegisterOtg_hs_deachintmskFieldOep1intmShift = 17
	RegisterOtg_hs_deachintmskFieldOep1intmMask  = 0x20000
)

// GetOep1intm OUT Endpoint 1 interrupt mask bit
func (r *registerOtg_hs_deachintmskType) GetOep1intm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_deachintmskFieldOep1intmMask) != 0
}

// SetOep1intm OUT Endpoint 1 interrupt mask bit
func (r *registerOtg_hs_deachintmskType) SetOep1intm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_deachintmskFieldOep1intmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_deachintmskFieldOep1intmMask)
	}
}

// registerOtg_hs_diepctl0Type OTG device endpoint-0 control register
type registerOtg_hs_diepctl0Type uint32

const (
	RegisterOtg_hs_diepctl0FieldMpsizShift = 0
	RegisterOtg_hs_diepctl0FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl0Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl0FieldMpsizMask) >> RegisterOtg_hs_diepctl0FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl0Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl0FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_diepctl0FieldMpsizShift))
}

const (
	RegisterOtg_hs_diepctl0FieldUsbaepShift = 15
	RegisterOtg_hs_diepctl0FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl0Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl0FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl0Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl0FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl0FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_diepctl0FieldEonum_dpidShift = 16
	RegisterOtg_hs_diepctl0FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even/odd frame
func (r *registerOtg_hs_diepctl0Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl0FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_diepctl0FieldNakstsShift = 17
	RegisterOtg_hs_diepctl0FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_diepctl0Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl0FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_diepctl0FieldEptypShift = 18
	RegisterOtg_hs_diepctl0FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_diepctl0Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl0FieldEptypMask) >> RegisterOtg_hs_diepctl0FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_diepctl0Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl0FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_diepctl0FieldEptypShift))
}

const (
	RegisterOtg_hs_diepctl0FieldStallShift = 21
	RegisterOtg_hs_diepctl0FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_diepctl0Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl0FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_diepctl0Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl0FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl0FieldStallMask)
	}
}

const (
	RegisterOtg_hs_diepctl0FieldTxfnumShift = 22
	RegisterOtg_hs_diepctl0FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl0Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl0FieldTxfnumMask) >> RegisterOtg_hs_diepctl0FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl0Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl0FieldTxfnumMask)|(uint32(value)<<RegisterOtg_hs_diepctl0FieldTxfnumShift))
}

const (
	RegisterOtg_hs_diepctl0FieldCnakShift = 26
	RegisterOtg_hs_diepctl0FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_diepctl0Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl0FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl0FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl0FieldSnakShift = 27
	RegisterOtg_hs_diepctl0FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_diepctl0Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl0FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl0FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl0FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_diepctl0FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID
func (r *registerOtg_hs_diepctl0Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl0FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl0FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl0FieldSoddfrmShift = 29
	RegisterOtg_hs_diepctl0FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_diepctl0Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl0FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl0FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl0FieldEpdisShift = 30
	RegisterOtg_hs_diepctl0FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl0Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl0FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl0Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl0FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl0FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_diepctl0FieldEpenaShift = 31
	RegisterOtg_hs_diepctl0FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_diepctl0Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl0FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_diepctl0Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl0FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl0FieldEpenaMask)
	}
}

// registerOtg_hs_diepctl1Type OTG device endpoint-1 control register
type registerOtg_hs_diepctl1Type uint32

const (
	RegisterOtg_hs_diepctl1FieldMpsizShift = 0
	RegisterOtg_hs_diepctl1FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl1Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl1FieldMpsizMask) >> RegisterOtg_hs_diepctl1FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl1Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl1FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_diepctl1FieldMpsizShift))
}

const (
	RegisterOtg_hs_diepctl1FieldUsbaepShift = 15
	RegisterOtg_hs_diepctl1FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl1Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl1FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl1Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl1FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl1FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_diepctl1FieldEonum_dpidShift = 16
	RegisterOtg_hs_diepctl1FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even/odd frame
func (r *registerOtg_hs_diepctl1Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl1FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_diepctl1FieldNakstsShift = 17
	RegisterOtg_hs_diepctl1FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_diepctl1Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl1FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_diepctl1FieldEptypShift = 18
	RegisterOtg_hs_diepctl1FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_diepctl1Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl1FieldEptypMask) >> RegisterOtg_hs_diepctl1FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_diepctl1Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl1FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_diepctl1FieldEptypShift))
}

const (
	RegisterOtg_hs_diepctl1FieldStallShift = 21
	RegisterOtg_hs_diepctl1FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_diepctl1Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl1FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_diepctl1Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl1FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl1FieldStallMask)
	}
}

const (
	RegisterOtg_hs_diepctl1FieldTxfnumShift = 22
	RegisterOtg_hs_diepctl1FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl1Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl1FieldTxfnumMask) >> RegisterOtg_hs_diepctl1FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl1Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl1FieldTxfnumMask)|(uint32(value)<<RegisterOtg_hs_diepctl1FieldTxfnumShift))
}

const (
	RegisterOtg_hs_diepctl1FieldCnakShift = 26
	RegisterOtg_hs_diepctl1FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_diepctl1Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl1FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl1FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl1FieldSnakShift = 27
	RegisterOtg_hs_diepctl1FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_diepctl1Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl1FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl1FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl1FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_diepctl1FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID
func (r *registerOtg_hs_diepctl1Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl1FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl1FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl1FieldSoddfrmShift = 29
	RegisterOtg_hs_diepctl1FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_diepctl1Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl1FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl1FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl1FieldEpdisShift = 30
	RegisterOtg_hs_diepctl1FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl1Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl1FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl1Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl1FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl1FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_diepctl1FieldEpenaShift = 31
	RegisterOtg_hs_diepctl1FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_diepctl1Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl1FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_diepctl1Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl1FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl1FieldEpenaMask)
	}
}

// registerOtg_hs_diepctl2Type OTG device endpoint-2 control register
type registerOtg_hs_diepctl2Type uint32

const (
	RegisterOtg_hs_diepctl2FieldMpsizShift = 0
	RegisterOtg_hs_diepctl2FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl2Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl2FieldMpsizMask) >> RegisterOtg_hs_diepctl2FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl2Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl2FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_diepctl2FieldMpsizShift))
}

const (
	RegisterOtg_hs_diepctl2FieldUsbaepShift = 15
	RegisterOtg_hs_diepctl2FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl2Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl2FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl2Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl2FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl2FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_diepctl2FieldEonum_dpidShift = 16
	RegisterOtg_hs_diepctl2FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even/odd frame
func (r *registerOtg_hs_diepctl2Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl2FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_diepctl2FieldNakstsShift = 17
	RegisterOtg_hs_diepctl2FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_diepctl2Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl2FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_diepctl2FieldEptypShift = 18
	RegisterOtg_hs_diepctl2FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_diepctl2Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl2FieldEptypMask) >> RegisterOtg_hs_diepctl2FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_diepctl2Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl2FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_diepctl2FieldEptypShift))
}

const (
	RegisterOtg_hs_diepctl2FieldStallShift = 21
	RegisterOtg_hs_diepctl2FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_diepctl2Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl2FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_diepctl2Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl2FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl2FieldStallMask)
	}
}

const (
	RegisterOtg_hs_diepctl2FieldTxfnumShift = 22
	RegisterOtg_hs_diepctl2FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl2Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl2FieldTxfnumMask) >> RegisterOtg_hs_diepctl2FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl2Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl2FieldTxfnumMask)|(uint32(value)<<RegisterOtg_hs_diepctl2FieldTxfnumShift))
}

const (
	RegisterOtg_hs_diepctl2FieldCnakShift = 26
	RegisterOtg_hs_diepctl2FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_diepctl2Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl2FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl2FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl2FieldSnakShift = 27
	RegisterOtg_hs_diepctl2FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_diepctl2Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl2FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl2FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl2FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_diepctl2FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID
func (r *registerOtg_hs_diepctl2Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl2FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl2FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl2FieldSoddfrmShift = 29
	RegisterOtg_hs_diepctl2FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_diepctl2Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl2FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl2FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl2FieldEpdisShift = 30
	RegisterOtg_hs_diepctl2FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl2Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl2FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl2Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl2FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl2FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_diepctl2FieldEpenaShift = 31
	RegisterOtg_hs_diepctl2FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_diepctl2Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl2FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_diepctl2Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl2FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl2FieldEpenaMask)
	}
}

// registerOtg_hs_diepctl3Type OTG device endpoint-3 control register
type registerOtg_hs_diepctl3Type uint32

const (
	RegisterOtg_hs_diepctl3FieldMpsizShift = 0
	RegisterOtg_hs_diepctl3FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl3Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl3FieldMpsizMask) >> RegisterOtg_hs_diepctl3FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl3Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl3FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_diepctl3FieldMpsizShift))
}

const (
	RegisterOtg_hs_diepctl3FieldUsbaepShift = 15
	RegisterOtg_hs_diepctl3FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl3Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl3FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl3Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl3FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl3FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_diepctl3FieldEonum_dpidShift = 16
	RegisterOtg_hs_diepctl3FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even/odd frame
func (r *registerOtg_hs_diepctl3Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl3FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_diepctl3FieldNakstsShift = 17
	RegisterOtg_hs_diepctl3FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_diepctl3Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl3FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_diepctl3FieldEptypShift = 18
	RegisterOtg_hs_diepctl3FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_diepctl3Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl3FieldEptypMask) >> RegisterOtg_hs_diepctl3FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_diepctl3Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl3FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_diepctl3FieldEptypShift))
}

const (
	RegisterOtg_hs_diepctl3FieldStallShift = 21
	RegisterOtg_hs_diepctl3FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_diepctl3Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl3FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_diepctl3Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl3FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl3FieldStallMask)
	}
}

const (
	RegisterOtg_hs_diepctl3FieldTxfnumShift = 22
	RegisterOtg_hs_diepctl3FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl3Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl3FieldTxfnumMask) >> RegisterOtg_hs_diepctl3FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl3Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl3FieldTxfnumMask)|(uint32(value)<<RegisterOtg_hs_diepctl3FieldTxfnumShift))
}

const (
	RegisterOtg_hs_diepctl3FieldCnakShift = 26
	RegisterOtg_hs_diepctl3FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_diepctl3Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl3FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl3FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl3FieldSnakShift = 27
	RegisterOtg_hs_diepctl3FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_diepctl3Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl3FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl3FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl3FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_diepctl3FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID
func (r *registerOtg_hs_diepctl3Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl3FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl3FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl3FieldSoddfrmShift = 29
	RegisterOtg_hs_diepctl3FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_diepctl3Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl3FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl3FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl3FieldEpdisShift = 30
	RegisterOtg_hs_diepctl3FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl3Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl3FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl3Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl3FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl3FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_diepctl3FieldEpenaShift = 31
	RegisterOtg_hs_diepctl3FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_diepctl3Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl3FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_diepctl3Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl3FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl3FieldEpenaMask)
	}
}

// registerOtg_hs_diepctl4Type OTG device endpoint-4 control register
type registerOtg_hs_diepctl4Type uint32

const (
	RegisterOtg_hs_diepctl4FieldMpsizShift = 0
	RegisterOtg_hs_diepctl4FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl4Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl4FieldMpsizMask) >> RegisterOtg_hs_diepctl4FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl4Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl4FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_diepctl4FieldMpsizShift))
}

const (
	RegisterOtg_hs_diepctl4FieldUsbaepShift = 15
	RegisterOtg_hs_diepctl4FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl4Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl4FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl4Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl4FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl4FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_diepctl4FieldEonum_dpidShift = 16
	RegisterOtg_hs_diepctl4FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even/odd frame
func (r *registerOtg_hs_diepctl4Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl4FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_diepctl4FieldNakstsShift = 17
	RegisterOtg_hs_diepctl4FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_diepctl4Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl4FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_diepctl4FieldEptypShift = 18
	RegisterOtg_hs_diepctl4FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_diepctl4Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl4FieldEptypMask) >> RegisterOtg_hs_diepctl4FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_diepctl4Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl4FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_diepctl4FieldEptypShift))
}

const (
	RegisterOtg_hs_diepctl4FieldStallShift = 21
	RegisterOtg_hs_diepctl4FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_diepctl4Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl4FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_diepctl4Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl4FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl4FieldStallMask)
	}
}

const (
	RegisterOtg_hs_diepctl4FieldTxfnumShift = 22
	RegisterOtg_hs_diepctl4FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl4Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl4FieldTxfnumMask) >> RegisterOtg_hs_diepctl4FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl4Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl4FieldTxfnumMask)|(uint32(value)<<RegisterOtg_hs_diepctl4FieldTxfnumShift))
}

const (
	RegisterOtg_hs_diepctl4FieldCnakShift = 26
	RegisterOtg_hs_diepctl4FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_diepctl4Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl4FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl4FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl4FieldSnakShift = 27
	RegisterOtg_hs_diepctl4FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_diepctl4Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl4FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl4FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl4FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_diepctl4FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID
func (r *registerOtg_hs_diepctl4Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl4FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl4FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl4FieldSoddfrmShift = 29
	RegisterOtg_hs_diepctl4FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_diepctl4Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl4FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl4FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl4FieldEpdisShift = 30
	RegisterOtg_hs_diepctl4FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl4Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl4FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl4Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl4FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl4FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_diepctl4FieldEpenaShift = 31
	RegisterOtg_hs_diepctl4FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_diepctl4Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl4FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_diepctl4Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl4FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl4FieldEpenaMask)
	}
}

// registerOtg_hs_diepctl5Type OTG device endpoint-5 control register
type registerOtg_hs_diepctl5Type uint32

const (
	RegisterOtg_hs_diepctl5FieldMpsizShift = 0
	RegisterOtg_hs_diepctl5FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl5Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl5FieldMpsizMask) >> RegisterOtg_hs_diepctl5FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl5Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl5FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_diepctl5FieldMpsizShift))
}

const (
	RegisterOtg_hs_diepctl5FieldUsbaepShift = 15
	RegisterOtg_hs_diepctl5FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl5Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl5FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl5Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl5FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl5FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_diepctl5FieldEonum_dpidShift = 16
	RegisterOtg_hs_diepctl5FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even/odd frame
func (r *registerOtg_hs_diepctl5Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl5FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_diepctl5FieldNakstsShift = 17
	RegisterOtg_hs_diepctl5FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_diepctl5Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl5FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_diepctl5FieldEptypShift = 18
	RegisterOtg_hs_diepctl5FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_diepctl5Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl5FieldEptypMask) >> RegisterOtg_hs_diepctl5FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_diepctl5Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl5FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_diepctl5FieldEptypShift))
}

const (
	RegisterOtg_hs_diepctl5FieldStallShift = 21
	RegisterOtg_hs_diepctl5FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_diepctl5Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl5FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_diepctl5Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl5FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl5FieldStallMask)
	}
}

const (
	RegisterOtg_hs_diepctl5FieldTxfnumShift = 22
	RegisterOtg_hs_diepctl5FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl5Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl5FieldTxfnumMask) >> RegisterOtg_hs_diepctl5FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl5Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl5FieldTxfnumMask)|(uint32(value)<<RegisterOtg_hs_diepctl5FieldTxfnumShift))
}

const (
	RegisterOtg_hs_diepctl5FieldCnakShift = 26
	RegisterOtg_hs_diepctl5FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_diepctl5Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl5FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl5FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl5FieldSnakShift = 27
	RegisterOtg_hs_diepctl5FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_diepctl5Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl5FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl5FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl5FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_diepctl5FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID
func (r *registerOtg_hs_diepctl5Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl5FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl5FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl5FieldSoddfrmShift = 29
	RegisterOtg_hs_diepctl5FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_diepctl5Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl5FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl5FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl5FieldEpdisShift = 30
	RegisterOtg_hs_diepctl5FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl5Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl5FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl5Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl5FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl5FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_diepctl5FieldEpenaShift = 31
	RegisterOtg_hs_diepctl5FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_diepctl5Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl5FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_diepctl5Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl5FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl5FieldEpenaMask)
	}
}

// registerOtg_hs_diepctl6Type OTG device endpoint-6 control register
type registerOtg_hs_diepctl6Type uint32

const (
	RegisterOtg_hs_diepctl6FieldMpsizShift = 0
	RegisterOtg_hs_diepctl6FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl6Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl6FieldMpsizMask) >> RegisterOtg_hs_diepctl6FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl6Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl6FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_diepctl6FieldMpsizShift))
}

const (
	RegisterOtg_hs_diepctl6FieldUsbaepShift = 15
	RegisterOtg_hs_diepctl6FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl6Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl6FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl6Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl6FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl6FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_diepctl6FieldEonum_dpidShift = 16
	RegisterOtg_hs_diepctl6FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even/odd frame
func (r *registerOtg_hs_diepctl6Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl6FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_diepctl6FieldNakstsShift = 17
	RegisterOtg_hs_diepctl6FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_diepctl6Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl6FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_diepctl6FieldEptypShift = 18
	RegisterOtg_hs_diepctl6FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_diepctl6Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl6FieldEptypMask) >> RegisterOtg_hs_diepctl6FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_diepctl6Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl6FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_diepctl6FieldEptypShift))
}

const (
	RegisterOtg_hs_diepctl6FieldStallShift = 21
	RegisterOtg_hs_diepctl6FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_diepctl6Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl6FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_diepctl6Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl6FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl6FieldStallMask)
	}
}

const (
	RegisterOtg_hs_diepctl6FieldTxfnumShift = 22
	RegisterOtg_hs_diepctl6FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl6Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl6FieldTxfnumMask) >> RegisterOtg_hs_diepctl6FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl6Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl6FieldTxfnumMask)|(uint32(value)<<RegisterOtg_hs_diepctl6FieldTxfnumShift))
}

const (
	RegisterOtg_hs_diepctl6FieldCnakShift = 26
	RegisterOtg_hs_diepctl6FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_diepctl6Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl6FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl6FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl6FieldSnakShift = 27
	RegisterOtg_hs_diepctl6FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_diepctl6Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl6FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl6FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl6FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_diepctl6FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID
func (r *registerOtg_hs_diepctl6Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl6FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl6FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl6FieldSoddfrmShift = 29
	RegisterOtg_hs_diepctl6FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_diepctl6Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl6FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl6FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl6FieldEpdisShift = 30
	RegisterOtg_hs_diepctl6FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl6Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl6FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl6Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl6FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl6FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_diepctl6FieldEpenaShift = 31
	RegisterOtg_hs_diepctl6FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_diepctl6Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl6FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_diepctl6Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl6FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl6FieldEpenaMask)
	}
}

// registerOtg_hs_diepctl7Type OTG device endpoint-7 control register
type registerOtg_hs_diepctl7Type uint32

const (
	RegisterOtg_hs_diepctl7FieldMpsizShift = 0
	RegisterOtg_hs_diepctl7FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl7Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl7FieldMpsizMask) >> RegisterOtg_hs_diepctl7FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_diepctl7Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl7FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_diepctl7FieldMpsizShift))
}

const (
	RegisterOtg_hs_diepctl7FieldUsbaepShift = 15
	RegisterOtg_hs_diepctl7FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl7Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl7FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_diepctl7Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl7FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl7FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_diepctl7FieldEonum_dpidShift = 16
	RegisterOtg_hs_diepctl7FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even/odd frame
func (r *registerOtg_hs_diepctl7Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl7FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_diepctl7FieldNakstsShift = 17
	RegisterOtg_hs_diepctl7FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_diepctl7Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl7FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_diepctl7FieldEptypShift = 18
	RegisterOtg_hs_diepctl7FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_diepctl7Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl7FieldEptypMask) >> RegisterOtg_hs_diepctl7FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_diepctl7Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl7FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_diepctl7FieldEptypShift))
}

const (
	RegisterOtg_hs_diepctl7FieldStallShift = 21
	RegisterOtg_hs_diepctl7FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_diepctl7Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl7FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_diepctl7Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl7FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl7FieldStallMask)
	}
}

const (
	RegisterOtg_hs_diepctl7FieldTxfnumShift = 22
	RegisterOtg_hs_diepctl7FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl7Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl7FieldTxfnumMask) >> RegisterOtg_hs_diepctl7FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtg_hs_diepctl7Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl7FieldTxfnumMask)|(uint32(value)<<RegisterOtg_hs_diepctl7FieldTxfnumShift))
}

const (
	RegisterOtg_hs_diepctl7FieldCnakShift = 26
	RegisterOtg_hs_diepctl7FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_diepctl7Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl7FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl7FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl7FieldSnakShift = 27
	RegisterOtg_hs_diepctl7FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_diepctl7Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl7FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl7FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_diepctl7FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_diepctl7FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID
func (r *registerOtg_hs_diepctl7Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl7FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl7FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl7FieldSoddfrmShift = 29
	RegisterOtg_hs_diepctl7FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_diepctl7Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl7FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl7FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_diepctl7FieldEpdisShift = 30
	RegisterOtg_hs_diepctl7FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl7Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl7FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_diepctl7Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl7FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl7FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_diepctl7FieldEpenaShift = 31
	RegisterOtg_hs_diepctl7FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_diepctl7Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepctl7FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_diepctl7Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepctl7FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepctl7FieldEpenaMask)
	}
}

// registerOtg_hs_diepint0Type OTG device endpoint-0 interrupt register
type registerOtg_hs_diepint0Type uint32

const (
	RegisterOtg_hs_diepint0FieldXfrcShift = 0
	RegisterOtg_hs_diepint0FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint0Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint0FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint0Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint0FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint0FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_diepint0FieldEpdisdShift = 1
	RegisterOtg_hs_diepint0FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint0Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint0FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint0Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint0FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint0FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_diepint0FieldTocShift = 3
	RegisterOtg_hs_diepint0FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtg_hs_diepint0Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint0FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtg_hs_diepint0Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint0FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint0FieldTocMask)
	}
}

const (
	RegisterOtg_hs_diepint0FieldIttxfeShift = 4
	RegisterOtg_hs_diepint0FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint0Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint0FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint0Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint0FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint0FieldIttxfeMask)
	}
}

const (
	RegisterOtg_hs_diepint0FieldInepneShift = 6
	RegisterOtg_hs_diepint0FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint0Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint0FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint0Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint0FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint0FieldInepneMask)
	}
}

const (
	RegisterOtg_hs_diepint0FieldTxfeShift = 7
	RegisterOtg_hs_diepint0FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtg_hs_diepint0Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint0FieldTxfeMask) != 0
}

const (
	RegisterOtg_hs_diepint0FieldTxfifoudrnShift = 8
	RegisterOtg_hs_diepint0FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint0Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint0FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint0Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint0FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint0FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtg_hs_diepint0FieldBnaShift = 9
	RegisterOtg_hs_diepint0FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint0Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint0FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint0Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint0FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint0FieldBnaMask)
	}
}

const (
	RegisterOtg_hs_diepint0FieldPktdrpstsShift = 11
	RegisterOtg_hs_diepint0FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint0Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint0FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint0Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint0FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint0FieldPktdrpstsMask)
	}
}

const (
	RegisterOtg_hs_diepint0FieldBerrShift = 12
	RegisterOtg_hs_diepint0FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtg_hs_diepint0Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint0FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtg_hs_diepint0Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint0FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint0FieldBerrMask)
	}
}

const (
	RegisterOtg_hs_diepint0FieldNakShift = 13
	RegisterOtg_hs_diepint0FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtg_hs_diepint0Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint0FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtg_hs_diepint0Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint0FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint0FieldNakMask)
	}
}

// registerOtg_hs_diepint1Type OTG device endpoint-1 interrupt register
type registerOtg_hs_diepint1Type uint32

const (
	RegisterOtg_hs_diepint1FieldXfrcShift = 0
	RegisterOtg_hs_diepint1FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint1Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint1FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint1Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint1FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint1FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_diepint1FieldEpdisdShift = 1
	RegisterOtg_hs_diepint1FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint1Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint1FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint1Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint1FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint1FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_diepint1FieldTocShift = 3
	RegisterOtg_hs_diepint1FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtg_hs_diepint1Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint1FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtg_hs_diepint1Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint1FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint1FieldTocMask)
	}
}

const (
	RegisterOtg_hs_diepint1FieldIttxfeShift = 4
	RegisterOtg_hs_diepint1FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint1Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint1FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint1Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint1FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint1FieldIttxfeMask)
	}
}

const (
	RegisterOtg_hs_diepint1FieldInepneShift = 6
	RegisterOtg_hs_diepint1FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint1Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint1FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint1Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint1FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint1FieldInepneMask)
	}
}

const (
	RegisterOtg_hs_diepint1FieldTxfeShift = 7
	RegisterOtg_hs_diepint1FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtg_hs_diepint1Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint1FieldTxfeMask) != 0
}

const (
	RegisterOtg_hs_diepint1FieldTxfifoudrnShift = 8
	RegisterOtg_hs_diepint1FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint1Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint1FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint1Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint1FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint1FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtg_hs_diepint1FieldBnaShift = 9
	RegisterOtg_hs_diepint1FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint1Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint1FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint1Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint1FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint1FieldBnaMask)
	}
}

const (
	RegisterOtg_hs_diepint1FieldPktdrpstsShift = 11
	RegisterOtg_hs_diepint1FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint1Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint1FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint1Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint1FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint1FieldPktdrpstsMask)
	}
}

const (
	RegisterOtg_hs_diepint1FieldBerrShift = 12
	RegisterOtg_hs_diepint1FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtg_hs_diepint1Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint1FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtg_hs_diepint1Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint1FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint1FieldBerrMask)
	}
}

const (
	RegisterOtg_hs_diepint1FieldNakShift = 13
	RegisterOtg_hs_diepint1FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtg_hs_diepint1Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint1FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtg_hs_diepint1Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint1FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint1FieldNakMask)
	}
}

// registerOtg_hs_diepint2Type OTG device endpoint-2 interrupt register
type registerOtg_hs_diepint2Type uint32

const (
	RegisterOtg_hs_diepint2FieldXfrcShift = 0
	RegisterOtg_hs_diepint2FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint2Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint2FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint2Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint2FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint2FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_diepint2FieldEpdisdShift = 1
	RegisterOtg_hs_diepint2FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint2Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint2FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint2Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint2FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint2FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_diepint2FieldTocShift = 3
	RegisterOtg_hs_diepint2FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtg_hs_diepint2Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint2FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtg_hs_diepint2Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint2FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint2FieldTocMask)
	}
}

const (
	RegisterOtg_hs_diepint2FieldIttxfeShift = 4
	RegisterOtg_hs_diepint2FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint2Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint2FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint2Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint2FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint2FieldIttxfeMask)
	}
}

const (
	RegisterOtg_hs_diepint2FieldInepneShift = 6
	RegisterOtg_hs_diepint2FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint2Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint2FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint2Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint2FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint2FieldInepneMask)
	}
}

const (
	RegisterOtg_hs_diepint2FieldTxfeShift = 7
	RegisterOtg_hs_diepint2FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtg_hs_diepint2Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint2FieldTxfeMask) != 0
}

const (
	RegisterOtg_hs_diepint2FieldTxfifoudrnShift = 8
	RegisterOtg_hs_diepint2FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint2Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint2FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint2Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint2FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint2FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtg_hs_diepint2FieldBnaShift = 9
	RegisterOtg_hs_diepint2FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint2Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint2FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint2Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint2FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint2FieldBnaMask)
	}
}

const (
	RegisterOtg_hs_diepint2FieldPktdrpstsShift = 11
	RegisterOtg_hs_diepint2FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint2Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint2FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint2Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint2FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint2FieldPktdrpstsMask)
	}
}

const (
	RegisterOtg_hs_diepint2FieldBerrShift = 12
	RegisterOtg_hs_diepint2FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtg_hs_diepint2Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint2FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtg_hs_diepint2Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint2FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint2FieldBerrMask)
	}
}

const (
	RegisterOtg_hs_diepint2FieldNakShift = 13
	RegisterOtg_hs_diepint2FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtg_hs_diepint2Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint2FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtg_hs_diepint2Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint2FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint2FieldNakMask)
	}
}

// registerOtg_hs_diepint3Type OTG device endpoint-3 interrupt register
type registerOtg_hs_diepint3Type uint32

const (
	RegisterOtg_hs_diepint3FieldXfrcShift = 0
	RegisterOtg_hs_diepint3FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint3Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint3FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint3Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint3FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint3FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_diepint3FieldEpdisdShift = 1
	RegisterOtg_hs_diepint3FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint3Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint3FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint3Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint3FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint3FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_diepint3FieldTocShift = 3
	RegisterOtg_hs_diepint3FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtg_hs_diepint3Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint3FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtg_hs_diepint3Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint3FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint3FieldTocMask)
	}
}

const (
	RegisterOtg_hs_diepint3FieldIttxfeShift = 4
	RegisterOtg_hs_diepint3FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint3Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint3FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint3Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint3FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint3FieldIttxfeMask)
	}
}

const (
	RegisterOtg_hs_diepint3FieldInepneShift = 6
	RegisterOtg_hs_diepint3FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint3Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint3FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint3Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint3FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint3FieldInepneMask)
	}
}

const (
	RegisterOtg_hs_diepint3FieldTxfeShift = 7
	RegisterOtg_hs_diepint3FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtg_hs_diepint3Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint3FieldTxfeMask) != 0
}

const (
	RegisterOtg_hs_diepint3FieldTxfifoudrnShift = 8
	RegisterOtg_hs_diepint3FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint3Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint3FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint3Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint3FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint3FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtg_hs_diepint3FieldBnaShift = 9
	RegisterOtg_hs_diepint3FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint3Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint3FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint3Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint3FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint3FieldBnaMask)
	}
}

const (
	RegisterOtg_hs_diepint3FieldPktdrpstsShift = 11
	RegisterOtg_hs_diepint3FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint3Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint3FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint3Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint3FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint3FieldPktdrpstsMask)
	}
}

const (
	RegisterOtg_hs_diepint3FieldBerrShift = 12
	RegisterOtg_hs_diepint3FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtg_hs_diepint3Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint3FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtg_hs_diepint3Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint3FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint3FieldBerrMask)
	}
}

const (
	RegisterOtg_hs_diepint3FieldNakShift = 13
	RegisterOtg_hs_diepint3FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtg_hs_diepint3Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint3FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtg_hs_diepint3Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint3FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint3FieldNakMask)
	}
}

// registerOtg_hs_diepint4Type OTG device endpoint-4 interrupt register
type registerOtg_hs_diepint4Type uint32

const (
	RegisterOtg_hs_diepint4FieldXfrcShift = 0
	RegisterOtg_hs_diepint4FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint4Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint4FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint4Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint4FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint4FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_diepint4FieldEpdisdShift = 1
	RegisterOtg_hs_diepint4FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint4Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint4FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint4Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint4FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint4FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_diepint4FieldTocShift = 3
	RegisterOtg_hs_diepint4FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtg_hs_diepint4Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint4FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtg_hs_diepint4Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint4FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint4FieldTocMask)
	}
}

const (
	RegisterOtg_hs_diepint4FieldIttxfeShift = 4
	RegisterOtg_hs_diepint4FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint4Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint4FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint4Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint4FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint4FieldIttxfeMask)
	}
}

const (
	RegisterOtg_hs_diepint4FieldInepneShift = 6
	RegisterOtg_hs_diepint4FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint4Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint4FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint4Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint4FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint4FieldInepneMask)
	}
}

const (
	RegisterOtg_hs_diepint4FieldTxfeShift = 7
	RegisterOtg_hs_diepint4FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtg_hs_diepint4Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint4FieldTxfeMask) != 0
}

const (
	RegisterOtg_hs_diepint4FieldTxfifoudrnShift = 8
	RegisterOtg_hs_diepint4FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint4Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint4FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint4Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint4FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint4FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtg_hs_diepint4FieldBnaShift = 9
	RegisterOtg_hs_diepint4FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint4Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint4FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint4Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint4FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint4FieldBnaMask)
	}
}

const (
	RegisterOtg_hs_diepint4FieldPktdrpstsShift = 11
	RegisterOtg_hs_diepint4FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint4Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint4FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint4Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint4FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint4FieldPktdrpstsMask)
	}
}

const (
	RegisterOtg_hs_diepint4FieldBerrShift = 12
	RegisterOtg_hs_diepint4FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtg_hs_diepint4Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint4FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtg_hs_diepint4Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint4FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint4FieldBerrMask)
	}
}

const (
	RegisterOtg_hs_diepint4FieldNakShift = 13
	RegisterOtg_hs_diepint4FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtg_hs_diepint4Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint4FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtg_hs_diepint4Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint4FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint4FieldNakMask)
	}
}

// registerOtg_hs_diepint5Type OTG device endpoint-5 interrupt register
type registerOtg_hs_diepint5Type uint32

const (
	RegisterOtg_hs_diepint5FieldXfrcShift = 0
	RegisterOtg_hs_diepint5FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint5Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint5FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint5Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint5FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint5FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_diepint5FieldEpdisdShift = 1
	RegisterOtg_hs_diepint5FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint5Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint5FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint5Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint5FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint5FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_diepint5FieldTocShift = 3
	RegisterOtg_hs_diepint5FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtg_hs_diepint5Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint5FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtg_hs_diepint5Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint5FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint5FieldTocMask)
	}
}

const (
	RegisterOtg_hs_diepint5FieldIttxfeShift = 4
	RegisterOtg_hs_diepint5FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint5Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint5FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint5Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint5FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint5FieldIttxfeMask)
	}
}

const (
	RegisterOtg_hs_diepint5FieldInepneShift = 6
	RegisterOtg_hs_diepint5FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint5Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint5FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint5Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint5FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint5FieldInepneMask)
	}
}

const (
	RegisterOtg_hs_diepint5FieldTxfeShift = 7
	RegisterOtg_hs_diepint5FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtg_hs_diepint5Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint5FieldTxfeMask) != 0
}

const (
	RegisterOtg_hs_diepint5FieldTxfifoudrnShift = 8
	RegisterOtg_hs_diepint5FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint5Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint5FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint5Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint5FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint5FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtg_hs_diepint5FieldBnaShift = 9
	RegisterOtg_hs_diepint5FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint5Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint5FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint5Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint5FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint5FieldBnaMask)
	}
}

const (
	RegisterOtg_hs_diepint5FieldPktdrpstsShift = 11
	RegisterOtg_hs_diepint5FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint5Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint5FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint5Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint5FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint5FieldPktdrpstsMask)
	}
}

const (
	RegisterOtg_hs_diepint5FieldBerrShift = 12
	RegisterOtg_hs_diepint5FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtg_hs_diepint5Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint5FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtg_hs_diepint5Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint5FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint5FieldBerrMask)
	}
}

const (
	RegisterOtg_hs_diepint5FieldNakShift = 13
	RegisterOtg_hs_diepint5FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtg_hs_diepint5Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint5FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtg_hs_diepint5Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint5FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint5FieldNakMask)
	}
}

// registerOtg_hs_diepint6Type OTG device endpoint-6 interrupt register
type registerOtg_hs_diepint6Type uint32

const (
	RegisterOtg_hs_diepint6FieldXfrcShift = 0
	RegisterOtg_hs_diepint6FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint6Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint6FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint6Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint6FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint6FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_diepint6FieldEpdisdShift = 1
	RegisterOtg_hs_diepint6FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint6Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint6FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint6Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint6FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint6FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_diepint6FieldTocShift = 3
	RegisterOtg_hs_diepint6FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtg_hs_diepint6Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint6FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtg_hs_diepint6Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint6FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint6FieldTocMask)
	}
}

const (
	RegisterOtg_hs_diepint6FieldIttxfeShift = 4
	RegisterOtg_hs_diepint6FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint6Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint6FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint6Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint6FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint6FieldIttxfeMask)
	}
}

const (
	RegisterOtg_hs_diepint6FieldInepneShift = 6
	RegisterOtg_hs_diepint6FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint6Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint6FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint6Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint6FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint6FieldInepneMask)
	}
}

const (
	RegisterOtg_hs_diepint6FieldTxfeShift = 7
	RegisterOtg_hs_diepint6FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtg_hs_diepint6Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint6FieldTxfeMask) != 0
}

const (
	RegisterOtg_hs_diepint6FieldTxfifoudrnShift = 8
	RegisterOtg_hs_diepint6FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint6Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint6FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint6Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint6FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint6FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtg_hs_diepint6FieldBnaShift = 9
	RegisterOtg_hs_diepint6FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint6Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint6FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint6Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint6FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint6FieldBnaMask)
	}
}

const (
	RegisterOtg_hs_diepint6FieldPktdrpstsShift = 11
	RegisterOtg_hs_diepint6FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint6Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint6FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint6Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint6FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint6FieldPktdrpstsMask)
	}
}

const (
	RegisterOtg_hs_diepint6FieldBerrShift = 12
	RegisterOtg_hs_diepint6FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtg_hs_diepint6Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint6FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtg_hs_diepint6Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint6FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint6FieldBerrMask)
	}
}

const (
	RegisterOtg_hs_diepint6FieldNakShift = 13
	RegisterOtg_hs_diepint6FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtg_hs_diepint6Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint6FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtg_hs_diepint6Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint6FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint6FieldNakMask)
	}
}

// registerOtg_hs_diepint7Type OTG device endpoint-7 interrupt register
type registerOtg_hs_diepint7Type uint32

const (
	RegisterOtg_hs_diepint7FieldXfrcShift = 0
	RegisterOtg_hs_diepint7FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint7Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint7FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_diepint7Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint7FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint7FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_diepint7FieldEpdisdShift = 1
	RegisterOtg_hs_diepint7FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint7Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint7FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_diepint7Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint7FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint7FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_diepint7FieldTocShift = 3
	RegisterOtg_hs_diepint7FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtg_hs_diepint7Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint7FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtg_hs_diepint7Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint7FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint7FieldTocMask)
	}
}

const (
	RegisterOtg_hs_diepint7FieldIttxfeShift = 4
	RegisterOtg_hs_diepint7FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint7Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint7FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtg_hs_diepint7Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint7FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint7FieldIttxfeMask)
	}
}

const (
	RegisterOtg_hs_diepint7FieldInepneShift = 6
	RegisterOtg_hs_diepint7FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint7Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint7FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtg_hs_diepint7Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint7FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint7FieldInepneMask)
	}
}

const (
	RegisterOtg_hs_diepint7FieldTxfeShift = 7
	RegisterOtg_hs_diepint7FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtg_hs_diepint7Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint7FieldTxfeMask) != 0
}

const (
	RegisterOtg_hs_diepint7FieldTxfifoudrnShift = 8
	RegisterOtg_hs_diepint7FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint7Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint7FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtg_hs_diepint7Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint7FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint7FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtg_hs_diepint7FieldBnaShift = 9
	RegisterOtg_hs_diepint7FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint7Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint7FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtg_hs_diepint7Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint7FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint7FieldBnaMask)
	}
}

const (
	RegisterOtg_hs_diepint7FieldPktdrpstsShift = 11
	RegisterOtg_hs_diepint7FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint7Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint7FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtg_hs_diepint7Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint7FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint7FieldPktdrpstsMask)
	}
}

const (
	RegisterOtg_hs_diepint7FieldBerrShift = 12
	RegisterOtg_hs_diepint7FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtg_hs_diepint7Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint7FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtg_hs_diepint7Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint7FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint7FieldBerrMask)
	}
}

const (
	RegisterOtg_hs_diepint7FieldNakShift = 13
	RegisterOtg_hs_diepint7FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtg_hs_diepint7Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepint7FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtg_hs_diepint7Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_diepint7FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepint7FieldNakMask)
	}
}

// registerOtg_hs_dieptsiz0Type OTG_HS device IN endpoint 0 transfer size register
type registerOtg_hs_dieptsiz0Type uint32

const (
	RegisterOtg_hs_dieptsiz0FieldXfrsizShift = 0
	RegisterOtg_hs_dieptsiz0FieldXfrsizMask  = 0x7f
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz0Type) GetXfrsiz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz0FieldXfrsizMask) >> RegisterOtg_hs_dieptsiz0FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz0Type) SetXfrsiz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz0FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz0FieldXfrsizShift))
}

const (
	RegisterOtg_hs_dieptsiz0FieldPktcntShift = 19
	RegisterOtg_hs_dieptsiz0FieldPktcntMask  = 0x180000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz0Type) GetPktcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz0FieldPktcntMask) >> RegisterOtg_hs_dieptsiz0FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz0Type) SetPktcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz0FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz0FieldPktcntShift))
}

// registerOtg_hs_diepdma1Type OTG_HS device endpoint-1 DMA address register
type registerOtg_hs_diepdma1Type uint32

const (
	RegisterOtg_hs_diepdma1FieldDmaaddrShift = 0
	RegisterOtg_hs_diepdma1FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_diepdma1Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepdma1FieldDmaaddrMask) >> RegisterOtg_hs_diepdma1FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_diepdma1Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepdma1FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_diepdma1FieldDmaaddrShift))
}

// registerOtg_hs_diepdma2Type OTG_HS device endpoint-2 DMA address register
type registerOtg_hs_diepdma2Type uint32

const (
	RegisterOtg_hs_diepdma2FieldDmaaddrShift = 0
	RegisterOtg_hs_diepdma2FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_diepdma2Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepdma2FieldDmaaddrMask) >> RegisterOtg_hs_diepdma2FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_diepdma2Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepdma2FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_diepdma2FieldDmaaddrShift))
}

// registerOtg_hs_diepdma3Type OTG_HS device endpoint-3 DMA address register
type registerOtg_hs_diepdma3Type uint32

const (
	RegisterOtg_hs_diepdma3FieldDmaaddrShift = 0
	RegisterOtg_hs_diepdma3FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_diepdma3Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepdma3FieldDmaaddrMask) >> RegisterOtg_hs_diepdma3FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_diepdma3Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepdma3FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_diepdma3FieldDmaaddrShift))
}

// registerOtg_hs_diepdma4Type OTG_HS device endpoint-4 DMA address register
type registerOtg_hs_diepdma4Type uint32

const (
	RegisterOtg_hs_diepdma4FieldDmaaddrShift = 0
	RegisterOtg_hs_diepdma4FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_diepdma4Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepdma4FieldDmaaddrMask) >> RegisterOtg_hs_diepdma4FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_diepdma4Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepdma4FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_diepdma4FieldDmaaddrShift))
}

// registerOtg_hs_diepdma5Type OTG_HS device endpoint-5 DMA address register
type registerOtg_hs_diepdma5Type uint32

const (
	RegisterOtg_hs_diepdma5FieldDmaaddrShift = 0
	RegisterOtg_hs_diepdma5FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_diepdma5Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_diepdma5FieldDmaaddrMask) >> RegisterOtg_hs_diepdma5FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_diepdma5Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_diepdma5FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_diepdma5FieldDmaaddrShift))
}

// registerOtg_hs_dtxfsts0Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtg_hs_dtxfsts0Type uint32

const (
	RegisterOtg_hs_dtxfsts0FieldIneptfsavShift = 0
	RegisterOtg_hs_dtxfsts0FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts0Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dtxfsts0FieldIneptfsavMask) >> RegisterOtg_hs_dtxfsts0FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts0Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dtxfsts0FieldIneptfsavMask)|(uint32(value)<<RegisterOtg_hs_dtxfsts0FieldIneptfsavShift))
}

// registerOtg_hs_dtxfsts1Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtg_hs_dtxfsts1Type uint32

const (
	RegisterOtg_hs_dtxfsts1FieldIneptfsavShift = 0
	RegisterOtg_hs_dtxfsts1FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts1Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dtxfsts1FieldIneptfsavMask) >> RegisterOtg_hs_dtxfsts1FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts1Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dtxfsts1FieldIneptfsavMask)|(uint32(value)<<RegisterOtg_hs_dtxfsts1FieldIneptfsavShift))
}

// registerOtg_hs_dtxfsts2Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtg_hs_dtxfsts2Type uint32

const (
	RegisterOtg_hs_dtxfsts2FieldIneptfsavShift = 0
	RegisterOtg_hs_dtxfsts2FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts2Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dtxfsts2FieldIneptfsavMask) >> RegisterOtg_hs_dtxfsts2FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts2Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dtxfsts2FieldIneptfsavMask)|(uint32(value)<<RegisterOtg_hs_dtxfsts2FieldIneptfsavShift))
}

// registerOtg_hs_dtxfsts3Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtg_hs_dtxfsts3Type uint32

const (
	RegisterOtg_hs_dtxfsts3FieldIneptfsavShift = 0
	RegisterOtg_hs_dtxfsts3FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts3Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dtxfsts3FieldIneptfsavMask) >> RegisterOtg_hs_dtxfsts3FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts3Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dtxfsts3FieldIneptfsavMask)|(uint32(value)<<RegisterOtg_hs_dtxfsts3FieldIneptfsavShift))
}

// registerOtg_hs_dtxfsts4Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtg_hs_dtxfsts4Type uint32

const (
	RegisterOtg_hs_dtxfsts4FieldIneptfsavShift = 0
	RegisterOtg_hs_dtxfsts4FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts4Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dtxfsts4FieldIneptfsavMask) >> RegisterOtg_hs_dtxfsts4FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts4Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dtxfsts4FieldIneptfsavMask)|(uint32(value)<<RegisterOtg_hs_dtxfsts4FieldIneptfsavShift))
}

// registerOtg_hs_dtxfsts5Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtg_hs_dtxfsts5Type uint32

const (
	RegisterOtg_hs_dtxfsts5FieldIneptfsavShift = 0
	RegisterOtg_hs_dtxfsts5FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts5Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dtxfsts5FieldIneptfsavMask) >> RegisterOtg_hs_dtxfsts5FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts5Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dtxfsts5FieldIneptfsavMask)|(uint32(value)<<RegisterOtg_hs_dtxfsts5FieldIneptfsavShift))
}

// registerOtg_hs_dieptsiz1Type OTG_HS device endpoint transfer size register
type registerOtg_hs_dieptsiz1Type uint32

const (
	RegisterOtg_hs_dieptsiz1FieldXfrsizShift = 0
	RegisterOtg_hs_dieptsiz1FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz1Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz1FieldXfrsizMask) >> RegisterOtg_hs_dieptsiz1FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz1Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz1FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz1FieldXfrsizShift))
}

const (
	RegisterOtg_hs_dieptsiz1FieldPktcntShift = 19
	RegisterOtg_hs_dieptsiz1FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz1Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz1FieldPktcntMask) >> RegisterOtg_hs_dieptsiz1FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz1Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz1FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz1FieldPktcntShift))
}

const (
	RegisterOtg_hs_dieptsiz1FieldMcntShift = 29
	RegisterOtg_hs_dieptsiz1FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtg_hs_dieptsiz1Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz1FieldMcntMask) >> RegisterOtg_hs_dieptsiz1FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtg_hs_dieptsiz1Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz1FieldMcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz1FieldMcntShift))
}

// registerOtg_hs_dieptsiz2Type OTG_HS device endpoint transfer size register
type registerOtg_hs_dieptsiz2Type uint32

const (
	RegisterOtg_hs_dieptsiz2FieldXfrsizShift = 0
	RegisterOtg_hs_dieptsiz2FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz2Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz2FieldXfrsizMask) >> RegisterOtg_hs_dieptsiz2FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz2Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz2FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz2FieldXfrsizShift))
}

const (
	RegisterOtg_hs_dieptsiz2FieldPktcntShift = 19
	RegisterOtg_hs_dieptsiz2FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz2Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz2FieldPktcntMask) >> RegisterOtg_hs_dieptsiz2FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz2Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz2FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz2FieldPktcntShift))
}

const (
	RegisterOtg_hs_dieptsiz2FieldMcntShift = 29
	RegisterOtg_hs_dieptsiz2FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtg_hs_dieptsiz2Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz2FieldMcntMask) >> RegisterOtg_hs_dieptsiz2FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtg_hs_dieptsiz2Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz2FieldMcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz2FieldMcntShift))
}

// registerOtg_hs_dieptsiz3Type OTG_HS device endpoint transfer size register
type registerOtg_hs_dieptsiz3Type uint32

const (
	RegisterOtg_hs_dieptsiz3FieldXfrsizShift = 0
	RegisterOtg_hs_dieptsiz3FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz3Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz3FieldXfrsizMask) >> RegisterOtg_hs_dieptsiz3FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz3Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz3FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz3FieldXfrsizShift))
}

const (
	RegisterOtg_hs_dieptsiz3FieldPktcntShift = 19
	RegisterOtg_hs_dieptsiz3FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz3Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz3FieldPktcntMask) >> RegisterOtg_hs_dieptsiz3FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz3Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz3FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz3FieldPktcntShift))
}

const (
	RegisterOtg_hs_dieptsiz3FieldMcntShift = 29
	RegisterOtg_hs_dieptsiz3FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtg_hs_dieptsiz3Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz3FieldMcntMask) >> RegisterOtg_hs_dieptsiz3FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtg_hs_dieptsiz3Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz3FieldMcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz3FieldMcntShift))
}

// registerOtg_hs_dieptsiz4Type OTG_HS device endpoint transfer size register
type registerOtg_hs_dieptsiz4Type uint32

const (
	RegisterOtg_hs_dieptsiz4FieldXfrsizShift = 0
	RegisterOtg_hs_dieptsiz4FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz4Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz4FieldXfrsizMask) >> RegisterOtg_hs_dieptsiz4FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz4Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz4FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz4FieldXfrsizShift))
}

const (
	RegisterOtg_hs_dieptsiz4FieldPktcntShift = 19
	RegisterOtg_hs_dieptsiz4FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz4Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz4FieldPktcntMask) >> RegisterOtg_hs_dieptsiz4FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz4Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz4FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz4FieldPktcntShift))
}

const (
	RegisterOtg_hs_dieptsiz4FieldMcntShift = 29
	RegisterOtg_hs_dieptsiz4FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtg_hs_dieptsiz4Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz4FieldMcntMask) >> RegisterOtg_hs_dieptsiz4FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtg_hs_dieptsiz4Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz4FieldMcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz4FieldMcntShift))
}

// registerOtg_hs_dieptsiz5Type OTG_HS device endpoint transfer size register
type registerOtg_hs_dieptsiz5Type uint32

const (
	RegisterOtg_hs_dieptsiz5FieldXfrsizShift = 0
	RegisterOtg_hs_dieptsiz5FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz5Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz5FieldXfrsizMask) >> RegisterOtg_hs_dieptsiz5FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz5Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz5FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz5FieldXfrsizShift))
}

const (
	RegisterOtg_hs_dieptsiz5FieldPktcntShift = 19
	RegisterOtg_hs_dieptsiz5FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz5Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz5FieldPktcntMask) >> RegisterOtg_hs_dieptsiz5FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz5Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz5FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz5FieldPktcntShift))
}

const (
	RegisterOtg_hs_dieptsiz5FieldMcntShift = 29
	RegisterOtg_hs_dieptsiz5FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtg_hs_dieptsiz5Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz5FieldMcntMask) >> RegisterOtg_hs_dieptsiz5FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtg_hs_dieptsiz5Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz5FieldMcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz5FieldMcntShift))
}

// registerOtg_hs_doepctl0Type OTG_HS device control OUT endpoint 0 control register
type registerOtg_hs_doepctl0Type uint32

const (
	RegisterOtg_hs_doepctl0FieldMpsizShift = 0
	RegisterOtg_hs_doepctl0FieldMpsizMask  = 0x3
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl0Type) GetMpsiz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl0FieldMpsizMask) >> RegisterOtg_hs_doepctl0FieldMpsizShift)
}

const (
	RegisterOtg_hs_doepctl0FieldUsbaepShift = 15
	RegisterOtg_hs_doepctl0FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl0Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl0FieldUsbaepMask) != 0
}

const (
	RegisterOtg_hs_doepctl0FieldNakstsShift = 17
	RegisterOtg_hs_doepctl0FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_doepctl0Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl0FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_doepctl0FieldEptypShift = 18
	RegisterOtg_hs_doepctl0FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_doepctl0Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl0FieldEptypMask) >> RegisterOtg_hs_doepctl0FieldEptypShift)
}

const (
	RegisterOtg_hs_doepctl0FieldSnpmShift = 20
	RegisterOtg_hs_doepctl0FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtg_hs_doepctl0Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl0FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtg_hs_doepctl0Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl0FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl0FieldSnpmMask)
	}
}

const (
	RegisterOtg_hs_doepctl0FieldStallShift = 21
	RegisterOtg_hs_doepctl0FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_doepctl0Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl0FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_doepctl0Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl0FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl0FieldStallMask)
	}
}

const (
	RegisterOtg_hs_doepctl0FieldCnakShift = 26
	RegisterOtg_hs_doepctl0FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_doepctl0Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl0FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl0FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl0FieldSnakShift = 27
	RegisterOtg_hs_doepctl0FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_doepctl0Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl0FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl0FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl0FieldEpdisShift = 30
	RegisterOtg_hs_doepctl0FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl0Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl0FieldEpdisMask) != 0
}

const (
	RegisterOtg_hs_doepctl0FieldEpenaShift = 31
	RegisterOtg_hs_doepctl0FieldEpenaMask  = 0x80000000
)

// SetEpena Endpoint enable
func (r *registerOtg_hs_doepctl0Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl0FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl0FieldEpenaMask)
	}
}

// registerOtg_hs_doepctl1Type OTG device endpoint-1 control register
type registerOtg_hs_doepctl1Type uint32

const (
	RegisterOtg_hs_doepctl1FieldMpsizShift = 0
	RegisterOtg_hs_doepctl1FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl1Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl1FieldMpsizMask) >> RegisterOtg_hs_doepctl1FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl1Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl1FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_doepctl1FieldMpsizShift))
}

const (
	RegisterOtg_hs_doepctl1FieldUsbaepShift = 15
	RegisterOtg_hs_doepctl1FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl1Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl1FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl1Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl1FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl1FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_doepctl1FieldEonum_dpidShift = 16
	RegisterOtg_hs_doepctl1FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even odd frame/Endpoint data PID
func (r *registerOtg_hs_doepctl1Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl1FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_doepctl1FieldNakstsShift = 17
	RegisterOtg_hs_doepctl1FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_doepctl1Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl1FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_doepctl1FieldEptypShift = 18
	RegisterOtg_hs_doepctl1FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_doepctl1Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl1FieldEptypMask) >> RegisterOtg_hs_doepctl1FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_doepctl1Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl1FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_doepctl1FieldEptypShift))
}

const (
	RegisterOtg_hs_doepctl1FieldSnpmShift = 20
	RegisterOtg_hs_doepctl1FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtg_hs_doepctl1Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl1FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtg_hs_doepctl1Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl1FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl1FieldSnpmMask)
	}
}

const (
	RegisterOtg_hs_doepctl1FieldStallShift = 21
	RegisterOtg_hs_doepctl1FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_doepctl1Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl1FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_doepctl1Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl1FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl1FieldStallMask)
	}
}

const (
	RegisterOtg_hs_doepctl1FieldCnakShift = 26
	RegisterOtg_hs_doepctl1FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_doepctl1Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl1FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl1FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl1FieldSnakShift = 27
	RegisterOtg_hs_doepctl1FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_doepctl1Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl1FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl1FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl1FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_doepctl1FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID/Set even frame
func (r *registerOtg_hs_doepctl1Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl1FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl1FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl1FieldSoddfrmShift = 29
	RegisterOtg_hs_doepctl1FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_doepctl1Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl1FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl1FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl1FieldEpdisShift = 30
	RegisterOtg_hs_doepctl1FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl1Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl1FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl1Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl1FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl1FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_doepctl1FieldEpenaShift = 31
	RegisterOtg_hs_doepctl1FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_doepctl1Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl1FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_doepctl1Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl1FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl1FieldEpenaMask)
	}
}

// registerOtg_hs_doepctl2Type OTG device endpoint-2 control register
type registerOtg_hs_doepctl2Type uint32

const (
	RegisterOtg_hs_doepctl2FieldMpsizShift = 0
	RegisterOtg_hs_doepctl2FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl2Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl2FieldMpsizMask) >> RegisterOtg_hs_doepctl2FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl2Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl2FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_doepctl2FieldMpsizShift))
}

const (
	RegisterOtg_hs_doepctl2FieldUsbaepShift = 15
	RegisterOtg_hs_doepctl2FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl2Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl2FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl2Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl2FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl2FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_doepctl2FieldEonum_dpidShift = 16
	RegisterOtg_hs_doepctl2FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even odd frame/Endpoint data PID
func (r *registerOtg_hs_doepctl2Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl2FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_doepctl2FieldNakstsShift = 17
	RegisterOtg_hs_doepctl2FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_doepctl2Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl2FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_doepctl2FieldEptypShift = 18
	RegisterOtg_hs_doepctl2FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_doepctl2Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl2FieldEptypMask) >> RegisterOtg_hs_doepctl2FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_doepctl2Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl2FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_doepctl2FieldEptypShift))
}

const (
	RegisterOtg_hs_doepctl2FieldSnpmShift = 20
	RegisterOtg_hs_doepctl2FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtg_hs_doepctl2Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl2FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtg_hs_doepctl2Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl2FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl2FieldSnpmMask)
	}
}

const (
	RegisterOtg_hs_doepctl2FieldStallShift = 21
	RegisterOtg_hs_doepctl2FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_doepctl2Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl2FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_doepctl2Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl2FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl2FieldStallMask)
	}
}

const (
	RegisterOtg_hs_doepctl2FieldCnakShift = 26
	RegisterOtg_hs_doepctl2FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_doepctl2Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl2FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl2FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl2FieldSnakShift = 27
	RegisterOtg_hs_doepctl2FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_doepctl2Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl2FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl2FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl2FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_doepctl2FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID/Set even frame
func (r *registerOtg_hs_doepctl2Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl2FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl2FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl2FieldSoddfrmShift = 29
	RegisterOtg_hs_doepctl2FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_doepctl2Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl2FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl2FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl2FieldEpdisShift = 30
	RegisterOtg_hs_doepctl2FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl2Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl2FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl2Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl2FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl2FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_doepctl2FieldEpenaShift = 31
	RegisterOtg_hs_doepctl2FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_doepctl2Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl2FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_doepctl2Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl2FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl2FieldEpenaMask)
	}
}

// registerOtg_hs_doepctl3Type OTG device endpoint-3 control register
type registerOtg_hs_doepctl3Type uint32

const (
	RegisterOtg_hs_doepctl3FieldMpsizShift = 0
	RegisterOtg_hs_doepctl3FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl3Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl3FieldMpsizMask) >> RegisterOtg_hs_doepctl3FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl3Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl3FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_doepctl3FieldMpsizShift))
}

const (
	RegisterOtg_hs_doepctl3FieldUsbaepShift = 15
	RegisterOtg_hs_doepctl3FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl3Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl3FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl3Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl3FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl3FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_doepctl3FieldEonum_dpidShift = 16
	RegisterOtg_hs_doepctl3FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even odd frame/Endpoint data PID
func (r *registerOtg_hs_doepctl3Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl3FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_doepctl3FieldNakstsShift = 17
	RegisterOtg_hs_doepctl3FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_doepctl3Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl3FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_doepctl3FieldEptypShift = 18
	RegisterOtg_hs_doepctl3FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_doepctl3Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl3FieldEptypMask) >> RegisterOtg_hs_doepctl3FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_doepctl3Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl3FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_doepctl3FieldEptypShift))
}

const (
	RegisterOtg_hs_doepctl3FieldSnpmShift = 20
	RegisterOtg_hs_doepctl3FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtg_hs_doepctl3Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl3FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtg_hs_doepctl3Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl3FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl3FieldSnpmMask)
	}
}

const (
	RegisterOtg_hs_doepctl3FieldStallShift = 21
	RegisterOtg_hs_doepctl3FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_doepctl3Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl3FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_doepctl3Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl3FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl3FieldStallMask)
	}
}

const (
	RegisterOtg_hs_doepctl3FieldCnakShift = 26
	RegisterOtg_hs_doepctl3FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_doepctl3Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl3FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl3FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl3FieldSnakShift = 27
	RegisterOtg_hs_doepctl3FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_doepctl3Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl3FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl3FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl3FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_doepctl3FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID/Set even frame
func (r *registerOtg_hs_doepctl3Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl3FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl3FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl3FieldSoddfrmShift = 29
	RegisterOtg_hs_doepctl3FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_doepctl3Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl3FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl3FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl3FieldEpdisShift = 30
	RegisterOtg_hs_doepctl3FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl3Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl3FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl3Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl3FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl3FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_doepctl3FieldEpenaShift = 31
	RegisterOtg_hs_doepctl3FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_doepctl3Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl3FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_doepctl3Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl3FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl3FieldEpenaMask)
	}
}

// registerOtg_hs_doepint0Type OTG_HS device endpoint-0 interrupt register
type registerOtg_hs_doepint0Type uint32

const (
	RegisterOtg_hs_doepint0FieldXfrcShift = 0
	RegisterOtg_hs_doepint0FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint0Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint0FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint0Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint0FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint0FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_doepint0FieldEpdisdShift = 1
	RegisterOtg_hs_doepint0FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint0Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint0FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint0Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint0FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint0FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_doepint0FieldStupShift = 3
	RegisterOtg_hs_doepint0FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtg_hs_doepint0Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint0FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtg_hs_doepint0Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint0FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint0FieldStupMask)
	}
}

const (
	RegisterOtg_hs_doepint0FieldOtepdisShift = 4
	RegisterOtg_hs_doepint0FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint0Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint0FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint0Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint0FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint0FieldOtepdisMask)
	}
}

const (
	RegisterOtg_hs_doepint0FieldB2bstupShift = 6
	RegisterOtg_hs_doepint0FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint0Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint0FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint0Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint0FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint0FieldB2bstupMask)
	}
}

const (
	RegisterOtg_hs_doepint0FieldNyetShift = 14
	RegisterOtg_hs_doepint0FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtg_hs_doepint0Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint0FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtg_hs_doepint0Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint0FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint0FieldNyetMask)
	}
}

// registerOtg_hs_doepint1Type OTG_HS device endpoint-1 interrupt register
type registerOtg_hs_doepint1Type uint32

const (
	RegisterOtg_hs_doepint1FieldXfrcShift = 0
	RegisterOtg_hs_doepint1FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint1Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint1FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint1Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint1FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint1FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_doepint1FieldEpdisdShift = 1
	RegisterOtg_hs_doepint1FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint1Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint1FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint1Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint1FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint1FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_doepint1FieldStupShift = 3
	RegisterOtg_hs_doepint1FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtg_hs_doepint1Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint1FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtg_hs_doepint1Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint1FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint1FieldStupMask)
	}
}

const (
	RegisterOtg_hs_doepint1FieldOtepdisShift = 4
	RegisterOtg_hs_doepint1FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint1Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint1FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint1Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint1FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint1FieldOtepdisMask)
	}
}

const (
	RegisterOtg_hs_doepint1FieldB2bstupShift = 6
	RegisterOtg_hs_doepint1FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint1Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint1FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint1Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint1FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint1FieldB2bstupMask)
	}
}

const (
	RegisterOtg_hs_doepint1FieldNyetShift = 14
	RegisterOtg_hs_doepint1FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtg_hs_doepint1Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint1FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtg_hs_doepint1Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint1FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint1FieldNyetMask)
	}
}

// registerOtg_hs_doepint2Type OTG_HS device endpoint-2 interrupt register
type registerOtg_hs_doepint2Type uint32

const (
	RegisterOtg_hs_doepint2FieldXfrcShift = 0
	RegisterOtg_hs_doepint2FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint2Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint2FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint2Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint2FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint2FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_doepint2FieldEpdisdShift = 1
	RegisterOtg_hs_doepint2FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint2Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint2FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint2Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint2FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint2FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_doepint2FieldStupShift = 3
	RegisterOtg_hs_doepint2FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtg_hs_doepint2Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint2FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtg_hs_doepint2Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint2FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint2FieldStupMask)
	}
}

const (
	RegisterOtg_hs_doepint2FieldOtepdisShift = 4
	RegisterOtg_hs_doepint2FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint2Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint2FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint2Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint2FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint2FieldOtepdisMask)
	}
}

const (
	RegisterOtg_hs_doepint2FieldB2bstupShift = 6
	RegisterOtg_hs_doepint2FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint2Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint2FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint2Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint2FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint2FieldB2bstupMask)
	}
}

const (
	RegisterOtg_hs_doepint2FieldNyetShift = 14
	RegisterOtg_hs_doepint2FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtg_hs_doepint2Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint2FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtg_hs_doepint2Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint2FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint2FieldNyetMask)
	}
}

// registerOtg_hs_doepint3Type OTG_HS device endpoint-3 interrupt register
type registerOtg_hs_doepint3Type uint32

const (
	RegisterOtg_hs_doepint3FieldXfrcShift = 0
	RegisterOtg_hs_doepint3FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint3Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint3FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint3Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint3FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint3FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_doepint3FieldEpdisdShift = 1
	RegisterOtg_hs_doepint3FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint3Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint3FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint3Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint3FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint3FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_doepint3FieldStupShift = 3
	RegisterOtg_hs_doepint3FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtg_hs_doepint3Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint3FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtg_hs_doepint3Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint3FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint3FieldStupMask)
	}
}

const (
	RegisterOtg_hs_doepint3FieldOtepdisShift = 4
	RegisterOtg_hs_doepint3FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint3Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint3FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint3Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint3FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint3FieldOtepdisMask)
	}
}

const (
	RegisterOtg_hs_doepint3FieldB2bstupShift = 6
	RegisterOtg_hs_doepint3FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint3Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint3FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint3Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint3FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint3FieldB2bstupMask)
	}
}

const (
	RegisterOtg_hs_doepint3FieldNyetShift = 14
	RegisterOtg_hs_doepint3FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtg_hs_doepint3Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint3FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtg_hs_doepint3Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint3FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint3FieldNyetMask)
	}
}

// registerOtg_hs_doepint4Type OTG_HS device endpoint-4 interrupt register
type registerOtg_hs_doepint4Type uint32

const (
	RegisterOtg_hs_doepint4FieldXfrcShift = 0
	RegisterOtg_hs_doepint4FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint4Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint4FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint4Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint4FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint4FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_doepint4FieldEpdisdShift = 1
	RegisterOtg_hs_doepint4FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint4Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint4FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint4Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint4FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint4FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_doepint4FieldStupShift = 3
	RegisterOtg_hs_doepint4FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtg_hs_doepint4Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint4FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtg_hs_doepint4Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint4FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint4FieldStupMask)
	}
}

const (
	RegisterOtg_hs_doepint4FieldOtepdisShift = 4
	RegisterOtg_hs_doepint4FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint4Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint4FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint4Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint4FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint4FieldOtepdisMask)
	}
}

const (
	RegisterOtg_hs_doepint4FieldB2bstupShift = 6
	RegisterOtg_hs_doepint4FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint4Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint4FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint4Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint4FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint4FieldB2bstupMask)
	}
}

const (
	RegisterOtg_hs_doepint4FieldNyetShift = 14
	RegisterOtg_hs_doepint4FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtg_hs_doepint4Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint4FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtg_hs_doepint4Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint4FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint4FieldNyetMask)
	}
}

// registerOtg_hs_doepint5Type OTG_HS device endpoint-5 interrupt register
type registerOtg_hs_doepint5Type uint32

const (
	RegisterOtg_hs_doepint5FieldXfrcShift = 0
	RegisterOtg_hs_doepint5FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint5Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint5FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint5Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint5FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint5FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_doepint5FieldEpdisdShift = 1
	RegisterOtg_hs_doepint5FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint5Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint5FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint5Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint5FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint5FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_doepint5FieldStupShift = 3
	RegisterOtg_hs_doepint5FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtg_hs_doepint5Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint5FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtg_hs_doepint5Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint5FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint5FieldStupMask)
	}
}

const (
	RegisterOtg_hs_doepint5FieldOtepdisShift = 4
	RegisterOtg_hs_doepint5FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint5Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint5FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint5Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint5FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint5FieldOtepdisMask)
	}
}

const (
	RegisterOtg_hs_doepint5FieldB2bstupShift = 6
	RegisterOtg_hs_doepint5FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint5Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint5FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint5Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint5FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint5FieldB2bstupMask)
	}
}

const (
	RegisterOtg_hs_doepint5FieldNyetShift = 14
	RegisterOtg_hs_doepint5FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtg_hs_doepint5Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint5FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtg_hs_doepint5Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint5FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint5FieldNyetMask)
	}
}

// registerOtg_hs_doepint6Type OTG_HS device endpoint-6 interrupt register
type registerOtg_hs_doepint6Type uint32

const (
	RegisterOtg_hs_doepint6FieldXfrcShift = 0
	RegisterOtg_hs_doepint6FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint6Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint6FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint6Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint6FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint6FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_doepint6FieldEpdisdShift = 1
	RegisterOtg_hs_doepint6FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint6Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint6FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint6Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint6FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint6FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_doepint6FieldStupShift = 3
	RegisterOtg_hs_doepint6FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtg_hs_doepint6Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint6FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtg_hs_doepint6Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint6FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint6FieldStupMask)
	}
}

const (
	RegisterOtg_hs_doepint6FieldOtepdisShift = 4
	RegisterOtg_hs_doepint6FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint6Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint6FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint6Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint6FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint6FieldOtepdisMask)
	}
}

const (
	RegisterOtg_hs_doepint6FieldB2bstupShift = 6
	RegisterOtg_hs_doepint6FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint6Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint6FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint6Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint6FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint6FieldB2bstupMask)
	}
}

const (
	RegisterOtg_hs_doepint6FieldNyetShift = 14
	RegisterOtg_hs_doepint6FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtg_hs_doepint6Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint6FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtg_hs_doepint6Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint6FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint6FieldNyetMask)
	}
}

// registerOtg_hs_doepint7Type OTG_HS device endpoint-7 interrupt register
type registerOtg_hs_doepint7Type uint32

const (
	RegisterOtg_hs_doepint7FieldXfrcShift = 0
	RegisterOtg_hs_doepint7FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint7Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint7FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtg_hs_doepint7Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint7FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint7FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_doepint7FieldEpdisdShift = 1
	RegisterOtg_hs_doepint7FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint7Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint7FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtg_hs_doepint7Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint7FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint7FieldEpdisdMask)
	}
}

const (
	RegisterOtg_hs_doepint7FieldStupShift = 3
	RegisterOtg_hs_doepint7FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtg_hs_doepint7Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint7FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtg_hs_doepint7Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint7FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint7FieldStupMask)
	}
}

const (
	RegisterOtg_hs_doepint7FieldOtepdisShift = 4
	RegisterOtg_hs_doepint7FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint7Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint7FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtg_hs_doepint7Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint7FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint7FieldOtepdisMask)
	}
}

const (
	RegisterOtg_hs_doepint7FieldB2bstupShift = 6
	RegisterOtg_hs_doepint7FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint7Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint7FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtg_hs_doepint7Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint7FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint7FieldB2bstupMask)
	}
}

const (
	RegisterOtg_hs_doepint7FieldNyetShift = 14
	RegisterOtg_hs_doepint7FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtg_hs_doepint7Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepint7FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtg_hs_doepint7Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepint7FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepint7FieldNyetMask)
	}
}

// registerOtg_hs_doeptsiz0Type OTG_HS device endpoint-0 transfer size register
type registerOtg_hs_doeptsiz0Type uint32

const (
	RegisterOtg_hs_doeptsiz0FieldXfrsizShift = 0
	RegisterOtg_hs_doeptsiz0FieldXfrsizMask  = 0x7f
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz0Type) GetXfrsiz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz0FieldXfrsizMask) >> RegisterOtg_hs_doeptsiz0FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz0Type) SetXfrsiz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz0FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz0FieldXfrsizShift))
}

const (
	RegisterOtg_hs_doeptsiz0FieldPktcntShift = 19
	RegisterOtg_hs_doeptsiz0FieldPktcntMask  = 0x80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz0Type) GetPktcnt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz0FieldPktcntMask) != 0
}

// SetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz0Type) SetPktcnt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doeptsiz0FieldPktcntMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz0FieldPktcntMask)
	}
}

const (
	RegisterOtg_hs_doeptsiz0FieldStupcntShift = 29
	RegisterOtg_hs_doeptsiz0FieldStupcntMask  = 0x60000000
)

// GetStupcnt SETUP packet count
func (r *registerOtg_hs_doeptsiz0Type) GetStupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz0FieldStupcntMask) >> RegisterOtg_hs_doeptsiz0FieldStupcntShift)
}

// SetStupcnt SETUP packet count
func (r *registerOtg_hs_doeptsiz0Type) SetStupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz0FieldStupcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz0FieldStupcntShift))
}

// registerOtg_hs_doeptsiz1Type OTG_HS device endpoint-1 transfer size register
type registerOtg_hs_doeptsiz1Type uint32

const (
	RegisterOtg_hs_doeptsiz1FieldXfrsizShift = 0
	RegisterOtg_hs_doeptsiz1FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz1Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz1FieldXfrsizMask) >> RegisterOtg_hs_doeptsiz1FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz1Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz1FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz1FieldXfrsizShift))
}

const (
	RegisterOtg_hs_doeptsiz1FieldPktcntShift = 19
	RegisterOtg_hs_doeptsiz1FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz1Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz1FieldPktcntMask) >> RegisterOtg_hs_doeptsiz1FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz1Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz1FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz1FieldPktcntShift))
}

const (
	RegisterOtg_hs_doeptsiz1FieldRxdpid_stupcntShift = 29
	RegisterOtg_hs_doeptsiz1FieldRxdpid_stupcntMask  = 0x60000000
)

// GetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz1Type) GetRxdpid_stupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz1FieldRxdpid_stupcntMask) >> RegisterOtg_hs_doeptsiz1FieldRxdpid_stupcntShift)
}

// SetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz1Type) SetRxdpid_stupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz1FieldRxdpid_stupcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz1FieldRxdpid_stupcntShift))
}

// registerOtg_hs_doeptsiz2Type OTG_HS device endpoint-2 transfer size register
type registerOtg_hs_doeptsiz2Type uint32

const (
	RegisterOtg_hs_doeptsiz2FieldXfrsizShift = 0
	RegisterOtg_hs_doeptsiz2FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz2Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz2FieldXfrsizMask) >> RegisterOtg_hs_doeptsiz2FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz2Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz2FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz2FieldXfrsizShift))
}

const (
	RegisterOtg_hs_doeptsiz2FieldPktcntShift = 19
	RegisterOtg_hs_doeptsiz2FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz2Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz2FieldPktcntMask) >> RegisterOtg_hs_doeptsiz2FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz2Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz2FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz2FieldPktcntShift))
}

const (
	RegisterOtg_hs_doeptsiz2FieldRxdpid_stupcntShift = 29
	RegisterOtg_hs_doeptsiz2FieldRxdpid_stupcntMask  = 0x60000000
)

// GetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz2Type) GetRxdpid_stupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz2FieldRxdpid_stupcntMask) >> RegisterOtg_hs_doeptsiz2FieldRxdpid_stupcntShift)
}

// SetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz2Type) SetRxdpid_stupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz2FieldRxdpid_stupcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz2FieldRxdpid_stupcntShift))
}

// registerOtg_hs_doeptsiz3Type OTG_HS device endpoint-3 transfer size register
type registerOtg_hs_doeptsiz3Type uint32

const (
	RegisterOtg_hs_doeptsiz3FieldXfrsizShift = 0
	RegisterOtg_hs_doeptsiz3FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz3Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz3FieldXfrsizMask) >> RegisterOtg_hs_doeptsiz3FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz3Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz3FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz3FieldXfrsizShift))
}

const (
	RegisterOtg_hs_doeptsiz3FieldPktcntShift = 19
	RegisterOtg_hs_doeptsiz3FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz3Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz3FieldPktcntMask) >> RegisterOtg_hs_doeptsiz3FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz3Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz3FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz3FieldPktcntShift))
}

const (
	RegisterOtg_hs_doeptsiz3FieldRxdpid_stupcntShift = 29
	RegisterOtg_hs_doeptsiz3FieldRxdpid_stupcntMask  = 0x60000000
)

// GetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz3Type) GetRxdpid_stupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz3FieldRxdpid_stupcntMask) >> RegisterOtg_hs_doeptsiz3FieldRxdpid_stupcntShift)
}

// SetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz3Type) SetRxdpid_stupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz3FieldRxdpid_stupcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz3FieldRxdpid_stupcntShift))
}

// registerOtg_hs_doeptsiz4Type OTG_HS device endpoint-4 transfer size register
type registerOtg_hs_doeptsiz4Type uint32

const (
	RegisterOtg_hs_doeptsiz4FieldXfrsizShift = 0
	RegisterOtg_hs_doeptsiz4FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz4Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz4FieldXfrsizMask) >> RegisterOtg_hs_doeptsiz4FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz4Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz4FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz4FieldXfrsizShift))
}

const (
	RegisterOtg_hs_doeptsiz4FieldPktcntShift = 19
	RegisterOtg_hs_doeptsiz4FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz4Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz4FieldPktcntMask) >> RegisterOtg_hs_doeptsiz4FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz4Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz4FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz4FieldPktcntShift))
}

const (
	RegisterOtg_hs_doeptsiz4FieldRxdpid_stupcntShift = 29
	RegisterOtg_hs_doeptsiz4FieldRxdpid_stupcntMask  = 0x60000000
)

// GetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz4Type) GetRxdpid_stupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz4FieldRxdpid_stupcntMask) >> RegisterOtg_hs_doeptsiz4FieldRxdpid_stupcntShift)
}

// SetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz4Type) SetRxdpid_stupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz4FieldRxdpid_stupcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz4FieldRxdpid_stupcntShift))
}

// registerOtg_hs_dieptsiz6Type OTG_HS device endpoint transfer size register
type registerOtg_hs_dieptsiz6Type uint32

const (
	RegisterOtg_hs_dieptsiz6FieldXfrsizShift = 0
	RegisterOtg_hs_dieptsiz6FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz6Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz6FieldXfrsizMask) >> RegisterOtg_hs_dieptsiz6FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz6Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz6FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz6FieldXfrsizShift))
}

const (
	RegisterOtg_hs_dieptsiz6FieldPktcntShift = 19
	RegisterOtg_hs_dieptsiz6FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz6Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz6FieldPktcntMask) >> RegisterOtg_hs_dieptsiz6FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz6Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz6FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz6FieldPktcntShift))
}

const (
	RegisterOtg_hs_dieptsiz6FieldMcntShift = 29
	RegisterOtg_hs_dieptsiz6FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtg_hs_dieptsiz6Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz6FieldMcntMask) >> RegisterOtg_hs_dieptsiz6FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtg_hs_dieptsiz6Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz6FieldMcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz6FieldMcntShift))
}

// registerOtg_hs_dtxfsts6Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtg_hs_dtxfsts6Type uint32

const (
	RegisterOtg_hs_dtxfsts6FieldIneptfsavShift = 0
	RegisterOtg_hs_dtxfsts6FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts6Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dtxfsts6FieldIneptfsavMask) >> RegisterOtg_hs_dtxfsts6FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts6Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dtxfsts6FieldIneptfsavMask)|(uint32(value)<<RegisterOtg_hs_dtxfsts6FieldIneptfsavShift))
}

// registerOtg_hs_dieptsiz7Type OTG_HS device endpoint transfer size register
type registerOtg_hs_dieptsiz7Type uint32

const (
	RegisterOtg_hs_dieptsiz7FieldXfrsizShift = 0
	RegisterOtg_hs_dieptsiz7FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz7Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz7FieldXfrsizMask) >> RegisterOtg_hs_dieptsiz7FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_dieptsiz7Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz7FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz7FieldXfrsizShift))
}

const (
	RegisterOtg_hs_dieptsiz7FieldPktcntShift = 19
	RegisterOtg_hs_dieptsiz7FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz7Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz7FieldPktcntMask) >> RegisterOtg_hs_dieptsiz7FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_dieptsiz7Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz7FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz7FieldPktcntShift))
}

const (
	RegisterOtg_hs_dieptsiz7FieldMcntShift = 29
	RegisterOtg_hs_dieptsiz7FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtg_hs_dieptsiz7Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptsiz7FieldMcntMask) >> RegisterOtg_hs_dieptsiz7FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtg_hs_dieptsiz7Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptsiz7FieldMcntMask)|(uint32(value)<<RegisterOtg_hs_dieptsiz7FieldMcntShift))
}

// registerOtg_hs_dtxfsts7Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtg_hs_dtxfsts7Type uint32

const (
	RegisterOtg_hs_dtxfsts7FieldIneptfsavShift = 0
	RegisterOtg_hs_dtxfsts7FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts7Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dtxfsts7FieldIneptfsavMask) >> RegisterOtg_hs_dtxfsts7FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtg_hs_dtxfsts7Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dtxfsts7FieldIneptfsavMask)|(uint32(value)<<RegisterOtg_hs_dtxfsts7FieldIneptfsavShift))
}

// registerOtg_hs_doepctl4Type OTG device endpoint-4 control register
type registerOtg_hs_doepctl4Type uint32

const (
	RegisterOtg_hs_doepctl4FieldMpsizShift = 0
	RegisterOtg_hs_doepctl4FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl4Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl4FieldMpsizMask) >> RegisterOtg_hs_doepctl4FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl4Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl4FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_doepctl4FieldMpsizShift))
}

const (
	RegisterOtg_hs_doepctl4FieldUsbaepShift = 15
	RegisterOtg_hs_doepctl4FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl4Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl4FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl4Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl4FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl4FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_doepctl4FieldEonum_dpidShift = 16
	RegisterOtg_hs_doepctl4FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even odd frame/Endpoint data PID
func (r *registerOtg_hs_doepctl4Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl4FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_doepctl4FieldNakstsShift = 17
	RegisterOtg_hs_doepctl4FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_doepctl4Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl4FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_doepctl4FieldEptypShift = 18
	RegisterOtg_hs_doepctl4FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_doepctl4Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl4FieldEptypMask) >> RegisterOtg_hs_doepctl4FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_doepctl4Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl4FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_doepctl4FieldEptypShift))
}

const (
	RegisterOtg_hs_doepctl4FieldSnpmShift = 20
	RegisterOtg_hs_doepctl4FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtg_hs_doepctl4Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl4FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtg_hs_doepctl4Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl4FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl4FieldSnpmMask)
	}
}

const (
	RegisterOtg_hs_doepctl4FieldStallShift = 21
	RegisterOtg_hs_doepctl4FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_doepctl4Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl4FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_doepctl4Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl4FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl4FieldStallMask)
	}
}

const (
	RegisterOtg_hs_doepctl4FieldCnakShift = 26
	RegisterOtg_hs_doepctl4FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_doepctl4Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl4FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl4FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl4FieldSnakShift = 27
	RegisterOtg_hs_doepctl4FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_doepctl4Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl4FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl4FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl4FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_doepctl4FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID/Set even frame
func (r *registerOtg_hs_doepctl4Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl4FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl4FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl4FieldSoddfrmShift = 29
	RegisterOtg_hs_doepctl4FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_doepctl4Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl4FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl4FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl4FieldEpdisShift = 30
	RegisterOtg_hs_doepctl4FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl4Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl4FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl4Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl4FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl4FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_doepctl4FieldEpenaShift = 31
	RegisterOtg_hs_doepctl4FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_doepctl4Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl4FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_doepctl4Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl4FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl4FieldEpenaMask)
	}
}

// registerOtg_hs_doepctl5Type OTG device endpoint-5 control register
type registerOtg_hs_doepctl5Type uint32

const (
	RegisterOtg_hs_doepctl5FieldMpsizShift = 0
	RegisterOtg_hs_doepctl5FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl5Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl5FieldMpsizMask) >> RegisterOtg_hs_doepctl5FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl5Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl5FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_doepctl5FieldMpsizShift))
}

const (
	RegisterOtg_hs_doepctl5FieldUsbaepShift = 15
	RegisterOtg_hs_doepctl5FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl5Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl5FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl5Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl5FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl5FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_doepctl5FieldEonum_dpidShift = 16
	RegisterOtg_hs_doepctl5FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even odd frame/Endpoint data PID
func (r *registerOtg_hs_doepctl5Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl5FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_doepctl5FieldNakstsShift = 17
	RegisterOtg_hs_doepctl5FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_doepctl5Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl5FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_doepctl5FieldEptypShift = 18
	RegisterOtg_hs_doepctl5FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_doepctl5Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl5FieldEptypMask) >> RegisterOtg_hs_doepctl5FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_doepctl5Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl5FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_doepctl5FieldEptypShift))
}

const (
	RegisterOtg_hs_doepctl5FieldSnpmShift = 20
	RegisterOtg_hs_doepctl5FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtg_hs_doepctl5Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl5FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtg_hs_doepctl5Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl5FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl5FieldSnpmMask)
	}
}

const (
	RegisterOtg_hs_doepctl5FieldStallShift = 21
	RegisterOtg_hs_doepctl5FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_doepctl5Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl5FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_doepctl5Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl5FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl5FieldStallMask)
	}
}

const (
	RegisterOtg_hs_doepctl5FieldCnakShift = 26
	RegisterOtg_hs_doepctl5FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_doepctl5Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl5FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl5FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl5FieldSnakShift = 27
	RegisterOtg_hs_doepctl5FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_doepctl5Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl5FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl5FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl5FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_doepctl5FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID/Set even frame
func (r *registerOtg_hs_doepctl5Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl5FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl5FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl5FieldSoddfrmShift = 29
	RegisterOtg_hs_doepctl5FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_doepctl5Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl5FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl5FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl5FieldEpdisShift = 30
	RegisterOtg_hs_doepctl5FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl5Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl5FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl5Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl5FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl5FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_doepctl5FieldEpenaShift = 31
	RegisterOtg_hs_doepctl5FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_doepctl5Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl5FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_doepctl5Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl5FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl5FieldEpenaMask)
	}
}

// registerOtg_hs_doepctl6Type OTG device endpoint-6 control register
type registerOtg_hs_doepctl6Type uint32

const (
	RegisterOtg_hs_doepctl6FieldMpsizShift = 0
	RegisterOtg_hs_doepctl6FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl6Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl6FieldMpsizMask) >> RegisterOtg_hs_doepctl6FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl6Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl6FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_doepctl6FieldMpsizShift))
}

const (
	RegisterOtg_hs_doepctl6FieldUsbaepShift = 15
	RegisterOtg_hs_doepctl6FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl6Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl6FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl6Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl6FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl6FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_doepctl6FieldEonum_dpidShift = 16
	RegisterOtg_hs_doepctl6FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even odd frame/Endpoint data PID
func (r *registerOtg_hs_doepctl6Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl6FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_doepctl6FieldNakstsShift = 17
	RegisterOtg_hs_doepctl6FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_doepctl6Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl6FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_doepctl6FieldEptypShift = 18
	RegisterOtg_hs_doepctl6FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_doepctl6Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl6FieldEptypMask) >> RegisterOtg_hs_doepctl6FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_doepctl6Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl6FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_doepctl6FieldEptypShift))
}

const (
	RegisterOtg_hs_doepctl6FieldSnpmShift = 20
	RegisterOtg_hs_doepctl6FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtg_hs_doepctl6Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl6FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtg_hs_doepctl6Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl6FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl6FieldSnpmMask)
	}
}

const (
	RegisterOtg_hs_doepctl6FieldStallShift = 21
	RegisterOtg_hs_doepctl6FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_doepctl6Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl6FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_doepctl6Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl6FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl6FieldStallMask)
	}
}

const (
	RegisterOtg_hs_doepctl6FieldCnakShift = 26
	RegisterOtg_hs_doepctl6FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_doepctl6Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl6FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl6FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl6FieldSnakShift = 27
	RegisterOtg_hs_doepctl6FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_doepctl6Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl6FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl6FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl6FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_doepctl6FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID/Set even frame
func (r *registerOtg_hs_doepctl6Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl6FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl6FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl6FieldSoddfrmShift = 29
	RegisterOtg_hs_doepctl6FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_doepctl6Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl6FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl6FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl6FieldEpdisShift = 30
	RegisterOtg_hs_doepctl6FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl6Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl6FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl6Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl6FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl6FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_doepctl6FieldEpenaShift = 31
	RegisterOtg_hs_doepctl6FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_doepctl6Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl6FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_doepctl6Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl6FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl6FieldEpenaMask)
	}
}

// registerOtg_hs_doepctl7Type OTG device endpoint-7 control register
type registerOtg_hs_doepctl7Type uint32

const (
	RegisterOtg_hs_doepctl7FieldMpsizShift = 0
	RegisterOtg_hs_doepctl7FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl7Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl7FieldMpsizMask) >> RegisterOtg_hs_doepctl7FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_doepctl7Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl7FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_doepctl7FieldMpsizShift))
}

const (
	RegisterOtg_hs_doepctl7FieldUsbaepShift = 15
	RegisterOtg_hs_doepctl7FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl7Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl7FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtg_hs_doepctl7Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl7FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl7FieldUsbaepMask)
	}
}

const (
	RegisterOtg_hs_doepctl7FieldEonum_dpidShift = 16
	RegisterOtg_hs_doepctl7FieldEonum_dpidMask  = 0x10000
)

// GetEonum_dpid Even odd frame/Endpoint data PID
func (r *registerOtg_hs_doepctl7Type) GetEonum_dpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl7FieldEonum_dpidMask) != 0
}

const (
	RegisterOtg_hs_doepctl7FieldNakstsShift = 17
	RegisterOtg_hs_doepctl7FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtg_hs_doepctl7Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl7FieldNakstsMask) != 0
}

const (
	RegisterOtg_hs_doepctl7FieldEptypShift = 18
	RegisterOtg_hs_doepctl7FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_doepctl7Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl7FieldEptypMask) >> RegisterOtg_hs_doepctl7FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_doepctl7Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl7FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_doepctl7FieldEptypShift))
}

const (
	RegisterOtg_hs_doepctl7FieldSnpmShift = 20
	RegisterOtg_hs_doepctl7FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtg_hs_doepctl7Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl7FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtg_hs_doepctl7Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl7FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl7FieldSnpmMask)
	}
}

const (
	RegisterOtg_hs_doepctl7FieldStallShift = 21
	RegisterOtg_hs_doepctl7FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtg_hs_doepctl7Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl7FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtg_hs_doepctl7Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl7FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl7FieldStallMask)
	}
}

const (
	RegisterOtg_hs_doepctl7FieldCnakShift = 26
	RegisterOtg_hs_doepctl7FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtg_hs_doepctl7Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl7FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl7FieldCnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl7FieldSnakShift = 27
	RegisterOtg_hs_doepctl7FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtg_hs_doepctl7Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl7FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl7FieldSnakMask)
	}
}

const (
	RegisterOtg_hs_doepctl7FieldSd0pid_sevnfrmShift = 28
	RegisterOtg_hs_doepctl7FieldSd0pid_sevnfrmMask  = 0x10000000
)

// SetSd0pid_sevnfrm Set DATA0 PID/Set even frame
func (r *registerOtg_hs_doepctl7Type) SetSd0pid_sevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl7FieldSd0pid_sevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl7FieldSd0pid_sevnfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl7FieldSoddfrmShift = 29
	RegisterOtg_hs_doepctl7FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtg_hs_doepctl7Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl7FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl7FieldSoddfrmMask)
	}
}

const (
	RegisterOtg_hs_doepctl7FieldEpdisShift = 30
	RegisterOtg_hs_doepctl7FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl7Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl7FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtg_hs_doepctl7Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl7FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl7FieldEpdisMask)
	}
}

const (
	RegisterOtg_hs_doepctl7FieldEpenaShift = 31
	RegisterOtg_hs_doepctl7FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtg_hs_doepctl7Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doepctl7FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtg_hs_doepctl7Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_doepctl7FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doepctl7FieldEpenaMask)
	}
}

// registerOtg_hs_doeptsiz5Type OTG_HS device endpoint-5 transfer size register
type registerOtg_hs_doeptsiz5Type uint32

const (
	RegisterOtg_hs_doeptsiz5FieldXfrsizShift = 0
	RegisterOtg_hs_doeptsiz5FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz5Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz5FieldXfrsizMask) >> RegisterOtg_hs_doeptsiz5FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz5Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz5FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz5FieldXfrsizShift))
}

const (
	RegisterOtg_hs_doeptsiz5FieldPktcntShift = 19
	RegisterOtg_hs_doeptsiz5FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz5Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz5FieldPktcntMask) >> RegisterOtg_hs_doeptsiz5FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz5Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz5FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz5FieldPktcntShift))
}

const (
	RegisterOtg_hs_doeptsiz5FieldRxdpid_stupcntShift = 29
	RegisterOtg_hs_doeptsiz5FieldRxdpid_stupcntMask  = 0x60000000
)

// GetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz5Type) GetRxdpid_stupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz5FieldRxdpid_stupcntMask) >> RegisterOtg_hs_doeptsiz5FieldRxdpid_stupcntShift)
}

// SetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz5Type) SetRxdpid_stupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz5FieldRxdpid_stupcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz5FieldRxdpid_stupcntShift))
}

// registerOtg_hs_doeptsiz6Type OTG_HS device endpoint-6 transfer size register
type registerOtg_hs_doeptsiz6Type uint32

const (
	RegisterOtg_hs_doeptsiz6FieldXfrsizShift = 0
	RegisterOtg_hs_doeptsiz6FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz6Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz6FieldXfrsizMask) >> RegisterOtg_hs_doeptsiz6FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz6Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz6FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz6FieldXfrsizShift))
}

const (
	RegisterOtg_hs_doeptsiz6FieldPktcntShift = 19
	RegisterOtg_hs_doeptsiz6FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz6Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz6FieldPktcntMask) >> RegisterOtg_hs_doeptsiz6FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz6Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz6FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz6FieldPktcntShift))
}

const (
	RegisterOtg_hs_doeptsiz6FieldRxdpid_stupcntShift = 29
	RegisterOtg_hs_doeptsiz6FieldRxdpid_stupcntMask  = 0x60000000
)

// GetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz6Type) GetRxdpid_stupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz6FieldRxdpid_stupcntMask) >> RegisterOtg_hs_doeptsiz6FieldRxdpid_stupcntShift)
}

// SetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz6Type) SetRxdpid_stupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz6FieldRxdpid_stupcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz6FieldRxdpid_stupcntShift))
}

// registerOtg_hs_doeptsiz7Type OTG_HS device endpoint-7 transfer size register
type registerOtg_hs_doeptsiz7Type uint32

const (
	RegisterOtg_hs_doeptsiz7FieldXfrsizShift = 0
	RegisterOtg_hs_doeptsiz7FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz7Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz7FieldXfrsizMask) >> RegisterOtg_hs_doeptsiz7FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_doeptsiz7Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz7FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz7FieldXfrsizShift))
}

const (
	RegisterOtg_hs_doeptsiz7FieldPktcntShift = 19
	RegisterOtg_hs_doeptsiz7FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz7Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz7FieldPktcntMask) >> RegisterOtg_hs_doeptsiz7FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_doeptsiz7Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz7FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz7FieldPktcntShift))
}

const (
	RegisterOtg_hs_doeptsiz7FieldRxdpid_stupcntShift = 29
	RegisterOtg_hs_doeptsiz7FieldRxdpid_stupcntMask  = 0x60000000
)

// GetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz7Type) GetRxdpid_stupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_doeptsiz7FieldRxdpid_stupcntMask) >> RegisterOtg_hs_doeptsiz7FieldRxdpid_stupcntShift)
}

// SetRxdpid_stupcnt Received data PID/SETUP packet count
func (r *registerOtg_hs_doeptsiz7Type) SetRxdpid_stupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_doeptsiz7FieldRxdpid_stupcntMask)|(uint32(value)<<RegisterOtg_hs_doeptsiz7FieldRxdpid_stupcntShift))
}
