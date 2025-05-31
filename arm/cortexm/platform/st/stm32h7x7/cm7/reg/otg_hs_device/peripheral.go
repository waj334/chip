//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package otg_hs_device

import (
	"unsafe"
	"volatile"
)

var (
	Otg1_hs_device = (*_otg_hs_device)(unsafe.Pointer(uintptr(0x40040800)))
	Otg2_hs_device = (*_otg_hs_device)(unsafe.Pointer(uintptr(0x40080800)))

	Instances = [2]*_otg_hs_device{
		Otg1_hs_device,
		Otg2_hs_device,
	}
)

type _otg_hs_device struct {
	Otghsdcfg        registerOtghsdcfgType
	Otghsdctl        registerOtghsdctlType
	Otghsdsts        registerOtghsdstsType
	_                [4]byte
	Otghsdiepmsk     registerOtghsdiepmskType
	Otghsdoepmsk     registerOtghsdoepmskType
	Otghsdaint       registerOtghsdaintType
	Otghsdaintmsk    registerOtghsdaintmskType
	_                [8]byte
	Otghsdvbusdis    registerOtghsdvbusdisType
	Otghsdvbuspulse  registerOtghsdvbuspulseType
	Otghsdthrctl     registerOtghsdthrctlType
	Otghsdiepempmsk  registerOtghsdiepempmskType
	Otghsdeachint    registerOtghsdeachintType
	Otghsdeachintmsk registerOtghsdeachintmskType
	_                [192]byte
	Otghsdiepctl0    registerOtghsdiepctl0Type
	_                [4]byte
	Otghsdiepint0    registerOtghsdiepint0Type
	_                [4]byte
	Otghsdieptsiz0   registerOtghsdieptsiz0Type
	Otghsdiepdma1    registerOtghsdiepdma1Type
	Otghsdtxfsts0    registerOtghsdtxfsts0Type
	_                [4]byte
	Otghsdiepctl1    registerOtghsdiepctl1Type
	_                [4]byte
	Otghsdiepint1    registerOtghsdiepint1Type
	_                [4]byte
	Otghsdieptsiz1   registerOtghsdieptsiz1Type
	Otghsdiepdma2    registerOtghsdiepdma2Type
	Otghsdtxfsts1    registerOtghsdtxfsts1Type
	_                [4]byte
	Otghsdiepctl2    registerOtghsdiepctl2Type
	_                [4]byte
	Otghsdiepint2    registerOtghsdiepint2Type
	_                [4]byte
	Otghsdieptsiz2   registerOtghsdieptsiz2Type
	Otghsdiepdma3    registerOtghsdiepdma3Type
	Otghsdtxfsts2    registerOtghsdtxfsts2Type
	_                [4]byte
	Otghsdiepctl3    registerOtghsdiepctl3Type
	_                [4]byte
	Otghsdiepint3    registerOtghsdiepint3Type
	_                [4]byte
	Otghsdieptsiz3   registerOtghsdieptsiz3Type
	Otghsdiepdma4    registerOtghsdiepdma4Type
	Otghsdtxfsts3    registerOtghsdtxfsts3Type
	_                [4]byte
	Otghsdiepctl4    registerOtghsdiepctl4Type
	_                [4]byte
	Otghsdiepint4    registerOtghsdiepint4Type
	_                [4]byte
	Otghsdieptsiz4   registerOtghsdieptsiz4Type
	Otghsdiepdma5    registerOtghsdiepdma5Type
	Otghsdtxfsts4    registerOtghsdtxfsts4Type
	_                [4]byte
	Otghsdiepctl5    registerOtghsdiepctl5Type
	Otghsdieptsiz6   registerOtghsdieptsiz6Type
	Otghsdtxfsts6    registerOtghsdtxfsts6Type
	Otghsdieptsiz7   registerOtghsdieptsiz7Type
	Otghsdiepint5    registerOtghsdiepint5Type
	Otghsdtxfsts7    registerOtghsdtxfsts7Type
	Otghsdieptsiz5   registerOtghsdieptsiz5Type
	_                [4]byte
	Otghsdtxfsts5    registerOtghsdtxfsts5Type
	_                [4]byte
	Otghsdiepctl6    registerOtghsdiepctl6Type
	_                [4]byte
	Otghsdiepint6    registerOtghsdiepint6Type
	_                [20]byte
	Otghsdiepctl7    registerOtghsdiepctl7Type
	_                [4]byte
	Otghsdiepint7    registerOtghsdiepint7Type
	_                [276]byte
	Otghsdoepctl0    registerOtghsdoepctl0Type
	_                [4]byte
	Otghsdoepint0    registerOtghsdoepint0Type
	_                [4]byte
	Otghsdoeptsiz0   registerOtghsdoeptsiz0Type
	_                [12]byte
	Otghsdoepctl1    registerOtghsdoepctl1Type
	_                [4]byte
	Otghsdoepint1    registerOtghsdoepint1Type
	_                [4]byte
	Otghsdoeptsiz1   registerOtghsdoeptsiz1Type
	_                [12]byte
	Otghsdoepctl2    registerOtghsdoepctl2Type
	_                [4]byte
	Otghsdoepint2    registerOtghsdoepint2Type
	_                [4]byte
	Otghsdoeptsiz2   registerOtghsdoeptsiz2Type
	_                [12]byte
	Otghsdoepctl3    registerOtghsdoepctl3Type
	_                [4]byte
	Otghsdoepint3    registerOtghsdoepint3Type
	_                [4]byte
	Otghsdoeptsiz3   registerOtghsdoeptsiz3Type
	_                [12]byte
	Otghsdoepctl4    registerOtghsdoepctl4Type
	_                [4]byte
	Otghsdoepint4    registerOtghsdoepint4Type
	_                [4]byte
	Otghsdoeptsiz4   registerOtghsdoeptsiz4Type
	_                [12]byte
	Otghsdoepctl5    registerOtghsdoepctl5Type
	_                [4]byte
	Otghsdoepint5    registerOtghsdoepint5Type
	_                [4]byte
	Otghsdoeptsiz5   registerOtghsdoeptsiz5Type
	_                [12]byte
	Otghsdoepctl6    registerOtghsdoepctl6Type
	_                [4]byte
	Otghsdoepint6    registerOtghsdoepint6Type
	_                [4]byte
	Otghsdoeptsiz6   registerOtghsdoeptsiz6Type
	_                [12]byte
	Otghsdoepctl7    registerOtghsdoepctl7Type
	_                [4]byte
	Otghsdoepint7    registerOtghsdoepint7Type
	_                [4]byte
	Otghsdoeptsiz7   registerOtghsdoeptsiz7Type
}

// registerOtghsdcfgType OTG_HS device configuration register
type registerOtghsdcfgType uint32

const (
	RegisterOtghsdcfgFieldDspdShift = 0
	RegisterOtghsdcfgFieldDspdMask  = 0x3
)

// GetDspd Device speed
func (r *registerOtghsdcfgType) GetDspd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdcfgFieldDspdMask) >> RegisterOtghsdcfgFieldDspdShift)
}

// SetDspd Device speed
func (r *registerOtghsdcfgType) SetDspd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdcfgFieldDspdMask)|(uint32(value)<<RegisterOtghsdcfgFieldDspdShift))
}

const (
	RegisterOtghsdcfgFieldNzlsohskShift = 2
	RegisterOtghsdcfgFieldNzlsohskMask  = 0x4
)

// GetNzlsohsk Nonzero-length status OUT handshake
func (r *registerOtghsdcfgType) GetNzlsohsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdcfgFieldNzlsohskMask) != 0
}

// SetNzlsohsk Nonzero-length status OUT handshake
func (r *registerOtghsdcfgType) SetNzlsohsk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdcfgFieldNzlsohskMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdcfgFieldNzlsohskMask)
	}
}

const (
	RegisterOtghsdcfgFieldDadShift = 4
	RegisterOtghsdcfgFieldDadMask  = 0x7f0
)

// GetDad Device address
func (r *registerOtghsdcfgType) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdcfgFieldDadMask) >> RegisterOtghsdcfgFieldDadShift)
}

// SetDad Device address
func (r *registerOtghsdcfgType) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdcfgFieldDadMask)|(uint32(value)<<RegisterOtghsdcfgFieldDadShift))
}

const (
	RegisterOtghsdcfgFieldPfivlShift = 11
	RegisterOtghsdcfgFieldPfivlMask  = 0x1800
)

// GetPfivl Periodic (micro)frame interval
func (r *registerOtghsdcfgType) GetPfivl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdcfgFieldPfivlMask) >> RegisterOtghsdcfgFieldPfivlShift)
}

// SetPfivl Periodic (micro)frame interval
func (r *registerOtghsdcfgType) SetPfivl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdcfgFieldPfivlMask)|(uint32(value)<<RegisterOtghsdcfgFieldPfivlShift))
}

const (
	RegisterOtghsdcfgFieldPerschivlShift = 24
	RegisterOtghsdcfgFieldPerschivlMask  = 0x3000000
)

// GetPerschivl Periodic scheduling interval
func (r *registerOtghsdcfgType) GetPerschivl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdcfgFieldPerschivlMask) >> RegisterOtghsdcfgFieldPerschivlShift)
}

// SetPerschivl Periodic scheduling interval
func (r *registerOtghsdcfgType) SetPerschivl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdcfgFieldPerschivlMask)|(uint32(value)<<RegisterOtghsdcfgFieldPerschivlShift))
}

// registerOtghsdctlType OTG_HS device control register
type registerOtghsdctlType uint32

const (
	RegisterOtghsdctlFieldRwusigShift = 0
	RegisterOtghsdctlFieldRwusigMask  = 0x1
)

// GetRwusig Remote wakeup signaling
func (r *registerOtghsdctlType) GetRwusig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdctlFieldRwusigMask) != 0
}

// SetRwusig Remote wakeup signaling
func (r *registerOtghsdctlType) SetRwusig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdctlFieldRwusigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdctlFieldRwusigMask)
	}
}

const (
	RegisterOtghsdctlFieldSdisShift = 1
	RegisterOtghsdctlFieldSdisMask  = 0x2
)

// GetSdis Soft disconnect
func (r *registerOtghsdctlType) GetSdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdctlFieldSdisMask) != 0
}

// SetSdis Soft disconnect
func (r *registerOtghsdctlType) SetSdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdctlFieldSdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdctlFieldSdisMask)
	}
}

const (
	RegisterOtghsdctlFieldGinstsShift = 2
	RegisterOtghsdctlFieldGinstsMask  = 0x4
)

// GetGinsts Global IN NAK status
func (r *registerOtghsdctlType) GetGinsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdctlFieldGinstsMask) != 0
}

const (
	RegisterOtghsdctlFieldGonstsShift = 3
	RegisterOtghsdctlFieldGonstsMask  = 0x8
)

// GetGonsts Global OUT NAK status
func (r *registerOtghsdctlType) GetGonsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdctlFieldGonstsMask) != 0
}

const (
	RegisterOtghsdctlFieldTctlShift = 4
	RegisterOtghsdctlFieldTctlMask  = 0x70
)

// GetTctl Test control
func (r *registerOtghsdctlType) GetTctl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdctlFieldTctlMask) >> RegisterOtghsdctlFieldTctlShift)
}

// SetTctl Test control
func (r *registerOtghsdctlType) SetTctl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdctlFieldTctlMask)|(uint32(value)<<RegisterOtghsdctlFieldTctlShift))
}

const (
	RegisterOtghsdctlFieldSginakShift = 7
	RegisterOtghsdctlFieldSginakMask  = 0x80
)

// SetSginak Set global IN NAK
func (r *registerOtghsdctlType) SetSginak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdctlFieldSginakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdctlFieldSginakMask)
	}
}

const (
	RegisterOtghsdctlFieldCginakShift = 8
	RegisterOtghsdctlFieldCginakMask  = 0x100
)

// SetCginak Clear global IN NAK
func (r *registerOtghsdctlType) SetCginak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdctlFieldCginakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdctlFieldCginakMask)
	}
}

const (
	RegisterOtghsdctlFieldSgonakShift = 9
	RegisterOtghsdctlFieldSgonakMask  = 0x200
)

// SetSgonak Set global OUT NAK
func (r *registerOtghsdctlType) SetSgonak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdctlFieldSgonakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdctlFieldSgonakMask)
	}
}

const (
	RegisterOtghsdctlFieldCgonakShift = 10
	RegisterOtghsdctlFieldCgonakMask  = 0x400
)

// SetCgonak Clear global OUT NAK
func (r *registerOtghsdctlType) SetCgonak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdctlFieldCgonakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdctlFieldCgonakMask)
	}
}

const (
	RegisterOtghsdctlFieldPoprgdneShift = 11
	RegisterOtghsdctlFieldPoprgdneMask  = 0x800
)

// GetPoprgdne Power-on programming done
func (r *registerOtghsdctlType) GetPoprgdne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdctlFieldPoprgdneMask) != 0
}

// SetPoprgdne Power-on programming done
func (r *registerOtghsdctlType) SetPoprgdne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdctlFieldPoprgdneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdctlFieldPoprgdneMask)
	}
}

// registerOtghsdstsType OTG_HS device status register
type registerOtghsdstsType uint32

const (
	RegisterOtghsdstsFieldSuspstsShift = 0
	RegisterOtghsdstsFieldSuspstsMask  = 0x1
)

// GetSuspsts Suspend status
func (r *registerOtghsdstsType) GetSuspsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdstsFieldSuspstsMask) != 0
}

// SetSuspsts Suspend status
func (r *registerOtghsdstsType) SetSuspsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdstsFieldSuspstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdstsFieldSuspstsMask)
	}
}

const (
	RegisterOtghsdstsFieldEnumspdShift = 1
	RegisterOtghsdstsFieldEnumspdMask  = 0x6
)

// GetEnumspd Enumerated speed
func (r *registerOtghsdstsType) GetEnumspd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdstsFieldEnumspdMask) >> RegisterOtghsdstsFieldEnumspdShift)
}

// SetEnumspd Enumerated speed
func (r *registerOtghsdstsType) SetEnumspd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdstsFieldEnumspdMask)|(uint32(value)<<RegisterOtghsdstsFieldEnumspdShift))
}

const (
	RegisterOtghsdstsFieldEerrShift = 3
	RegisterOtghsdstsFieldEerrMask  = 0x8
)

// GetEerr Erratic error
func (r *registerOtghsdstsType) GetEerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdstsFieldEerrMask) != 0
}

// SetEerr Erratic error
func (r *registerOtghsdstsType) SetEerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdstsFieldEerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdstsFieldEerrMask)
	}
}

const (
	RegisterOtghsdstsFieldFnsofShift = 8
	RegisterOtghsdstsFieldFnsofMask  = 0x3fff00
)

// GetFnsof Frame number of the received SOF
func (r *registerOtghsdstsType) GetFnsof() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdstsFieldFnsofMask) >> RegisterOtghsdstsFieldFnsofShift)
}

// SetFnsof Frame number of the received SOF
func (r *registerOtghsdstsType) SetFnsof(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdstsFieldFnsofMask)|(uint32(value)<<RegisterOtghsdstsFieldFnsofShift))
}

// registerOtghsdiepmskType OTG_HS device IN endpoint common interrupt mask register
type registerOtghsdiepmskType uint32

const (
	RegisterOtghsdiepmskFieldXfrcmShift = 0
	RegisterOtghsdiepmskFieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed interrupt mask
func (r *registerOtghsdiepmskType) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed interrupt mask
func (r *registerOtghsdiepmskType) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepmskFieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepmskFieldXfrcmMask)
	}
}

const (
	RegisterOtghsdiepmskFieldEpdmShift = 1
	RegisterOtghsdiepmskFieldEpdmMask  = 0x2
)

// GetEpdm Endpoint disabled interrupt mask
func (r *registerOtghsdiepmskType) GetEpdm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldEpdmMask) != 0
}

// SetEpdm Endpoint disabled interrupt mask
func (r *registerOtghsdiepmskType) SetEpdm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepmskFieldEpdmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepmskFieldEpdmMask)
	}
}

const (
	RegisterOtghsdiepmskFieldTomShift = 3
	RegisterOtghsdiepmskFieldTomMask  = 0x8
)

// GetTom Timeout condition mask (nonisochronous endpoints)
func (r *registerOtghsdiepmskType) GetTom() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldTomMask) != 0
}

// SetTom Timeout condition mask (nonisochronous endpoints)
func (r *registerOtghsdiepmskType) SetTom(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepmskFieldTomMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepmskFieldTomMask)
	}
}

const (
	RegisterOtghsdiepmskFieldIttxfemskShift = 4
	RegisterOtghsdiepmskFieldIttxfemskMask  = 0x10
)

// GetIttxfemsk IN token received when TxFIFO empty mask
func (r *registerOtghsdiepmskType) GetIttxfemsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldIttxfemskMask) != 0
}

// SetIttxfemsk IN token received when TxFIFO empty mask
func (r *registerOtghsdiepmskType) SetIttxfemsk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepmskFieldIttxfemskMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepmskFieldIttxfemskMask)
	}
}

const (
	RegisterOtghsdiepmskFieldInepnmmShift = 5
	RegisterOtghsdiepmskFieldInepnmmMask  = 0x20
)

// GetInepnmm IN token received with EP mismatch mask
func (r *registerOtghsdiepmskType) GetInepnmm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldInepnmmMask) != 0
}

// SetInepnmm IN token received with EP mismatch mask
func (r *registerOtghsdiepmskType) SetInepnmm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepmskFieldInepnmmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepmskFieldInepnmmMask)
	}
}

const (
	RegisterOtghsdiepmskFieldInepnemShift = 6
	RegisterOtghsdiepmskFieldInepnemMask  = 0x40
)

// GetInepnem IN endpoint NAK effective mask
func (r *registerOtghsdiepmskType) GetInepnem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldInepnemMask) != 0
}

// SetInepnem IN endpoint NAK effective mask
func (r *registerOtghsdiepmskType) SetInepnem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepmskFieldInepnemMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepmskFieldInepnemMask)
	}
}

const (
	RegisterOtghsdiepmskFieldTxfurmShift = 8
	RegisterOtghsdiepmskFieldTxfurmMask  = 0x100
)

// GetTxfurm FIFO underrun mask
func (r *registerOtghsdiepmskType) GetTxfurm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldTxfurmMask) != 0
}

// SetTxfurm FIFO underrun mask
func (r *registerOtghsdiepmskType) SetTxfurm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepmskFieldTxfurmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepmskFieldTxfurmMask)
	}
}

const (
	RegisterOtghsdiepmskFieldBimShift = 9
	RegisterOtghsdiepmskFieldBimMask  = 0x200
)

// GetBim BNA interrupt mask
func (r *registerOtghsdiepmskType) GetBim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldBimMask) != 0
}

// SetBim BNA interrupt mask
func (r *registerOtghsdiepmskType) SetBim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepmskFieldBimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepmskFieldBimMask)
	}
}

// registerOtghsdoepmskType OTG_HS device OUT endpoint common interrupt mask register
type registerOtghsdoepmskType uint32

const (
	RegisterOtghsdoepmskFieldXfrcmShift = 0
	RegisterOtghsdoepmskFieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed interrupt mask
func (r *registerOtghsdoepmskType) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed interrupt mask
func (r *registerOtghsdoepmskType) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepmskFieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepmskFieldXfrcmMask)
	}
}

const (
	RegisterOtghsdoepmskFieldEpdmShift = 1
	RegisterOtghsdoepmskFieldEpdmMask  = 0x2
)

// GetEpdm Endpoint disabled interrupt mask
func (r *registerOtghsdoepmskType) GetEpdm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldEpdmMask) != 0
}

// SetEpdm Endpoint disabled interrupt mask
func (r *registerOtghsdoepmskType) SetEpdm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepmskFieldEpdmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepmskFieldEpdmMask)
	}
}

const (
	RegisterOtghsdoepmskFieldStupmShift = 3
	RegisterOtghsdoepmskFieldStupmMask  = 0x8
)

// GetStupm SETUP phase done mask
func (r *registerOtghsdoepmskType) GetStupm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldStupmMask) != 0
}

// SetStupm SETUP phase done mask
func (r *registerOtghsdoepmskType) SetStupm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepmskFieldStupmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepmskFieldStupmMask)
	}
}

const (
	RegisterOtghsdoepmskFieldOtepdmShift = 4
	RegisterOtghsdoepmskFieldOtepdmMask  = 0x10
)

// GetOtepdm OUT token received when endpoint disabled mask
func (r *registerOtghsdoepmskType) GetOtepdm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldOtepdmMask) != 0
}

// SetOtepdm OUT token received when endpoint disabled mask
func (r *registerOtghsdoepmskType) SetOtepdm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepmskFieldOtepdmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepmskFieldOtepdmMask)
	}
}

const (
	RegisterOtghsdoepmskFieldB2bstupShift = 6
	RegisterOtghsdoepmskFieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received mask
func (r *registerOtghsdoepmskType) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received mask
func (r *registerOtghsdoepmskType) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepmskFieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepmskFieldB2bstupMask)
	}
}

const (
	RegisterOtghsdoepmskFieldOpemShift = 8
	RegisterOtghsdoepmskFieldOpemMask  = 0x100
)

// GetOpem OUT packet error mask
func (r *registerOtghsdoepmskType) GetOpem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldOpemMask) != 0
}

// SetOpem OUT packet error mask
func (r *registerOtghsdoepmskType) SetOpem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepmskFieldOpemMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepmskFieldOpemMask)
	}
}

const (
	RegisterOtghsdoepmskFieldBoimShift = 9
	RegisterOtghsdoepmskFieldBoimMask  = 0x200
)

// GetBoim BNA interrupt mask
func (r *registerOtghsdoepmskType) GetBoim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldBoimMask) != 0
}

// SetBoim BNA interrupt mask
func (r *registerOtghsdoepmskType) SetBoim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepmskFieldBoimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepmskFieldBoimMask)
	}
}

// registerOtghsdaintType OTG_HS device all endpoints interrupt register
type registerOtghsdaintType uint32

const (
	RegisterOtghsdaintFieldIepintShift = 0
	RegisterOtghsdaintFieldIepintMask  = 0xffff
)

// GetIepint IN endpoint interrupt bits
func (r *registerOtghsdaintType) GetIepint() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdaintFieldIepintMask) >> RegisterOtghsdaintFieldIepintShift)
}

// SetIepint IN endpoint interrupt bits
func (r *registerOtghsdaintType) SetIepint(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdaintFieldIepintMask)|(uint32(value)<<RegisterOtghsdaintFieldIepintShift))
}

const (
	RegisterOtghsdaintFieldOepintShift = 16
	RegisterOtghsdaintFieldOepintMask  = 0xffff0000
)

// GetOepint OUT endpoint interrupt bits
func (r *registerOtghsdaintType) GetOepint() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdaintFieldOepintMask) >> RegisterOtghsdaintFieldOepintShift)
}

// SetOepint OUT endpoint interrupt bits
func (r *registerOtghsdaintType) SetOepint(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdaintFieldOepintMask)|(uint32(value)<<RegisterOtghsdaintFieldOepintShift))
}

// registerOtghsdaintmskType OTG_HS all endpoints interrupt mask register
type registerOtghsdaintmskType uint32

const (
	RegisterOtghsdaintmskFieldIepmShift = 0
	RegisterOtghsdaintmskFieldIepmMask  = 0xffff
)

// GetIepm IN EP interrupt mask bits
func (r *registerOtghsdaintmskType) GetIepm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdaintmskFieldIepmMask) >> RegisterOtghsdaintmskFieldIepmShift)
}

// SetIepm IN EP interrupt mask bits
func (r *registerOtghsdaintmskType) SetIepm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdaintmskFieldIepmMask)|(uint32(value)<<RegisterOtghsdaintmskFieldIepmShift))
}

const (
	RegisterOtghsdaintmskFieldOepmShift = 16
	RegisterOtghsdaintmskFieldOepmMask  = 0xffff0000
)

// GetOepm OUT EP interrupt mask bits
func (r *registerOtghsdaintmskType) GetOepm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdaintmskFieldOepmMask) >> RegisterOtghsdaintmskFieldOepmShift)
}

// SetOepm OUT EP interrupt mask bits
func (r *registerOtghsdaintmskType) SetOepm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdaintmskFieldOepmMask)|(uint32(value)<<RegisterOtghsdaintmskFieldOepmShift))
}

// registerOtghsdvbusdisType OTG_HS device VBUS discharge time register
type registerOtghsdvbusdisType uint32

const (
	RegisterOtghsdvbusdisFieldVbusdtShift = 0
	RegisterOtghsdvbusdisFieldVbusdtMask  = 0xffff
)

// GetVbusdt Device VBUS discharge time
func (r *registerOtghsdvbusdisType) GetVbusdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdvbusdisFieldVbusdtMask) >> RegisterOtghsdvbusdisFieldVbusdtShift)
}

// SetVbusdt Device VBUS discharge time
func (r *registerOtghsdvbusdisType) SetVbusdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdvbusdisFieldVbusdtMask)|(uint32(value)<<RegisterOtghsdvbusdisFieldVbusdtShift))
}

// registerOtghsdvbuspulseType OTG_HS device VBUS pulsing time register
type registerOtghsdvbuspulseType uint32

const (
	RegisterOtghsdvbuspulseFieldDvbuspShift = 0
	RegisterOtghsdvbuspulseFieldDvbuspMask  = 0xfff
)

// GetDvbusp Device VBUS pulsing time
func (r *registerOtghsdvbuspulseType) GetDvbusp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdvbuspulseFieldDvbuspMask) >> RegisterOtghsdvbuspulseFieldDvbuspShift)
}

// SetDvbusp Device VBUS pulsing time
func (r *registerOtghsdvbuspulseType) SetDvbusp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdvbuspulseFieldDvbuspMask)|(uint32(value)<<RegisterOtghsdvbuspulseFieldDvbuspShift))
}

// registerOtghsdthrctlType OTG_HS Device threshold control register
type registerOtghsdthrctlType uint32

const (
	RegisterOtghsdthrctlFieldNonisothrenShift = 0
	RegisterOtghsdthrctlFieldNonisothrenMask  = 0x1
)

// GetNonisothren Nonisochronous IN endpoints threshold enable
func (r *registerOtghsdthrctlType) GetNonisothren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdthrctlFieldNonisothrenMask) != 0
}

// SetNonisothren Nonisochronous IN endpoints threshold enable
func (r *registerOtghsdthrctlType) SetNonisothren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdthrctlFieldNonisothrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdthrctlFieldNonisothrenMask)
	}
}

const (
	RegisterOtghsdthrctlFieldIsothrenShift = 1
	RegisterOtghsdthrctlFieldIsothrenMask  = 0x2
)

// GetIsothren ISO IN endpoint threshold enable
func (r *registerOtghsdthrctlType) GetIsothren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdthrctlFieldIsothrenMask) != 0
}

// SetIsothren ISO IN endpoint threshold enable
func (r *registerOtghsdthrctlType) SetIsothren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdthrctlFieldIsothrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdthrctlFieldIsothrenMask)
	}
}

const (
	RegisterOtghsdthrctlFieldTxthrlenShift = 2
	RegisterOtghsdthrctlFieldTxthrlenMask  = 0x7fc
)

// GetTxthrlen Transmit threshold length
func (r *registerOtghsdthrctlType) GetTxthrlen() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdthrctlFieldTxthrlenMask) >> RegisterOtghsdthrctlFieldTxthrlenShift)
}

// SetTxthrlen Transmit threshold length
func (r *registerOtghsdthrctlType) SetTxthrlen(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdthrctlFieldTxthrlenMask)|(uint32(value)<<RegisterOtghsdthrctlFieldTxthrlenShift))
}

const (
	RegisterOtghsdthrctlFieldRxthrenShift = 16
	RegisterOtghsdthrctlFieldRxthrenMask  = 0x10000
)

// GetRxthren Receive threshold enable
func (r *registerOtghsdthrctlType) GetRxthren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdthrctlFieldRxthrenMask) != 0
}

// SetRxthren Receive threshold enable
func (r *registerOtghsdthrctlType) SetRxthren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdthrctlFieldRxthrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdthrctlFieldRxthrenMask)
	}
}

const (
	RegisterOtghsdthrctlFieldRxthrlenShift = 17
	RegisterOtghsdthrctlFieldRxthrlenMask  = 0x3fe0000
)

// GetRxthrlen Receive threshold length
func (r *registerOtghsdthrctlType) GetRxthrlen() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdthrctlFieldRxthrlenMask) >> RegisterOtghsdthrctlFieldRxthrlenShift)
}

// SetRxthrlen Receive threshold length
func (r *registerOtghsdthrctlType) SetRxthrlen(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdthrctlFieldRxthrlenMask)|(uint32(value)<<RegisterOtghsdthrctlFieldRxthrlenShift))
}

const (
	RegisterOtghsdthrctlFieldArpenShift = 27
	RegisterOtghsdthrctlFieldArpenMask  = 0x8000000
)

// GetArpen Arbiter parking enable
func (r *registerOtghsdthrctlType) GetArpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdthrctlFieldArpenMask) != 0
}

// SetArpen Arbiter parking enable
func (r *registerOtghsdthrctlType) SetArpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdthrctlFieldArpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdthrctlFieldArpenMask)
	}
}

// registerOtghsdiepempmskType OTG_HS device IN endpoint FIFO empty interrupt mask register
type registerOtghsdiepempmskType uint32

const (
	RegisterOtghsdiepempmskFieldIneptxfemShift = 0
	RegisterOtghsdiepempmskFieldIneptxfemMask  = 0xffff
)

// GetIneptxfem IN EP Tx FIFO empty interrupt mask bits
func (r *registerOtghsdiepempmskType) GetIneptxfem() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepempmskFieldIneptxfemMask) >> RegisterOtghsdiepempmskFieldIneptxfemShift)
}

// SetIneptxfem IN EP Tx FIFO empty interrupt mask bits
func (r *registerOtghsdiepempmskType) SetIneptxfem(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepempmskFieldIneptxfemMask)|(uint32(value)<<RegisterOtghsdiepempmskFieldIneptxfemShift))
}

// registerOtghsdeachintType OTG_HS device each endpoint interrupt register
type registerOtghsdeachintType uint32

const (
	RegisterOtghsdeachintFieldIep1intShift = 1
	RegisterOtghsdeachintFieldIep1intMask  = 0x2
)

// GetIep1int IN endpoint 1interrupt bit
func (r *registerOtghsdeachintType) GetIep1int() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdeachintFieldIep1intMask) != 0
}

// SetIep1int IN endpoint 1interrupt bit
func (r *registerOtghsdeachintType) SetIep1int(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdeachintFieldIep1intMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdeachintFieldIep1intMask)
	}
}

const (
	RegisterOtghsdeachintFieldOep1intShift = 17
	RegisterOtghsdeachintFieldOep1intMask  = 0x20000
)

// GetOep1int OUT endpoint 1 interrupt bit
func (r *registerOtghsdeachintType) GetOep1int() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdeachintFieldOep1intMask) != 0
}

// SetOep1int OUT endpoint 1 interrupt bit
func (r *registerOtghsdeachintType) SetOep1int(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdeachintFieldOep1intMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdeachintFieldOep1intMask)
	}
}

// registerOtghsdeachintmskType OTG_HS device each endpoint interrupt register mask
type registerOtghsdeachintmskType uint32

const (
	RegisterOtghsdeachintmskFieldIep1intmShift = 1
	RegisterOtghsdeachintmskFieldIep1intmMask  = 0x2
)

// GetIep1intm IN Endpoint 1 interrupt mask bit
func (r *registerOtghsdeachintmskType) GetIep1intm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdeachintmskFieldIep1intmMask) != 0
}

// SetIep1intm IN Endpoint 1 interrupt mask bit
func (r *registerOtghsdeachintmskType) SetIep1intm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdeachintmskFieldIep1intmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdeachintmskFieldIep1intmMask)
	}
}

const (
	RegisterOtghsdeachintmskFieldOep1intmShift = 17
	RegisterOtghsdeachintmskFieldOep1intmMask  = 0x20000
)

// GetOep1intm OUT Endpoint 1 interrupt mask bit
func (r *registerOtghsdeachintmskType) GetOep1intm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdeachintmskFieldOep1intmMask) != 0
}

// SetOep1intm OUT Endpoint 1 interrupt mask bit
func (r *registerOtghsdeachintmskType) SetOep1intm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdeachintmskFieldOep1intmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdeachintmskFieldOep1intmMask)
	}
}

// registerOtghsdiepctl0Type OTG device endpoint-0 control register
type registerOtghsdiepctl0Type uint32

const (
	RegisterOtghsdiepctl0FieldMpsizShift = 0
	RegisterOtghsdiepctl0FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdiepctl0Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldMpsizMask) >> RegisterOtghsdiepctl0FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdiepctl0Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl0FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl0FieldUsbaepShift = 15
	RegisterOtghsdiepctl0FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdiepctl0Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdiepctl0Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl0FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdiepctl0FieldEonumdpidShift = 16
	RegisterOtghsdiepctl0FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even/odd frame
func (r *registerOtghsdiepctl0Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl0FieldNakstsShift = 17
	RegisterOtghsdiepctl0FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdiepctl0Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl0FieldEptypShift = 18
	RegisterOtghsdiepctl0FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdiepctl0Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldEptypMask) >> RegisterOtghsdiepctl0FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdiepctl0Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl0FieldEptypShift))
}

const (
	RegisterOtghsdiepctl0FieldStallShift = 21
	RegisterOtghsdiepctl0FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdiepctl0Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdiepctl0Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl0FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldStallMask)
	}
}

const (
	RegisterOtghsdiepctl0FieldTxfnumShift = 22
	RegisterOtghsdiepctl0FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtghsdiepctl0Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldTxfnumMask) >> RegisterOtghsdiepctl0FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtghsdiepctl0Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl0FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl0FieldCnakShift = 26
	RegisterOtghsdiepctl0FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdiepctl0Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl0FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldCnakMask)
	}
}

const (
	RegisterOtghsdiepctl0FieldSnakShift = 27
	RegisterOtghsdiepctl0FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdiepctl0Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl0FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldSnakMask)
	}
}

const (
	RegisterOtghsdiepctl0FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdiepctl0FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID
func (r *registerOtghsdiepctl0Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl0FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdiepctl0FieldSoddfrmShift = 29
	RegisterOtghsdiepctl0FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdiepctl0Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl0FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdiepctl0FieldEpdisShift = 30
	RegisterOtghsdiepctl0FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdiepctl0Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdiepctl0Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl0FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldEpdisMask)
	}
}

const (
	RegisterOtghsdiepctl0FieldEpenaShift = 31
	RegisterOtghsdiepctl0FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdiepctl0Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdiepctl0Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl0FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldEpenaMask)
	}
}

// registerOtghsdiepint0Type OTG device endpoint-0 interrupt register
type registerOtghsdiepint0Type uint32

const (
	RegisterOtghsdiepint0FieldXfrcShift = 0
	RegisterOtghsdiepint0FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint0Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint0Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint0FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint0FieldXfrcMask)
	}
}

const (
	RegisterOtghsdiepint0FieldEpdisdShift = 1
	RegisterOtghsdiepint0FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint0Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint0Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint0FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint0FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdiepint0FieldTocShift = 3
	RegisterOtghsdiepint0FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtghsdiepint0Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtghsdiepint0Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint0FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint0FieldTocMask)
	}
}

const (
	RegisterOtghsdiepint0FieldIttxfeShift = 4
	RegisterOtghsdiepint0FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint0Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint0Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint0FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint0FieldIttxfeMask)
	}
}

const (
	RegisterOtghsdiepint0FieldInepneShift = 6
	RegisterOtghsdiepint0FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint0Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint0Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint0FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint0FieldInepneMask)
	}
}

const (
	RegisterOtghsdiepint0FieldTxfeShift = 7
	RegisterOtghsdiepint0FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtghsdiepint0Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint0FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint0FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint0Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint0Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint0FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint0FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtghsdiepint0FieldBnaShift = 9
	RegisterOtghsdiepint0FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtghsdiepint0Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtghsdiepint0Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint0FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint0FieldBnaMask)
	}
}

const (
	RegisterOtghsdiepint0FieldPktdrpstsShift = 11
	RegisterOtghsdiepint0FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint0Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint0Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint0FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint0FieldPktdrpstsMask)
	}
}

const (
	RegisterOtghsdiepint0FieldBerrShift = 12
	RegisterOtghsdiepint0FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtghsdiepint0Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtghsdiepint0Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint0FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint0FieldBerrMask)
	}
}

const (
	RegisterOtghsdiepint0FieldNakShift = 13
	RegisterOtghsdiepint0FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtghsdiepint0Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtghsdiepint0Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint0FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint0FieldNakMask)
	}
}

// registerOtghsdieptsiz0Type OTG_HS device IN endpoint 0 transfer size register
type registerOtghsdieptsiz0Type uint32

const (
	RegisterOtghsdieptsiz0FieldXfrsizShift = 0
	RegisterOtghsdieptsiz0FieldXfrsizMask  = 0x7f
)

// GetXfrsiz Transfer size
func (r *registerOtghsdieptsiz0Type) GetXfrsiz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz0FieldXfrsizMask) >> RegisterOtghsdieptsiz0FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdieptsiz0Type) SetXfrsiz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz0FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz0FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz0FieldPktcntShift = 19
	RegisterOtghsdieptsiz0FieldPktcntMask  = 0x180000
)

// GetPktcnt Packet count
func (r *registerOtghsdieptsiz0Type) GetPktcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz0FieldPktcntMask) >> RegisterOtghsdieptsiz0FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdieptsiz0Type) SetPktcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz0FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz0FieldPktcntShift))
}

// registerOtghsdiepdma1Type OTG_HS device endpoint-1 DMA address register
type registerOtghsdiepdma1Type uint32

const (
	RegisterOtghsdiepdma1FieldDmaaddrShift = 0
	RegisterOtghsdiepdma1FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghsdiepdma1Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepdma1FieldDmaaddrMask) >> RegisterOtghsdiepdma1FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghsdiepdma1Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepdma1FieldDmaaddrMask)|(uint32(value)<<RegisterOtghsdiepdma1FieldDmaaddrShift))
}

// registerOtghsdtxfsts0Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtghsdtxfsts0Type uint32

const (
	RegisterOtghsdtxfsts0FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts0FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts0Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts0FieldIneptfsavMask) >> RegisterOtghsdtxfsts0FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts0Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts0FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts0FieldIneptfsavShift))
}

// registerOtghsdiepctl1Type OTG device endpoint-1 control register
type registerOtghsdiepctl1Type uint32

const (
	RegisterOtghsdiepctl1FieldMpsizShift = 0
	RegisterOtghsdiepctl1FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdiepctl1Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldMpsizMask) >> RegisterOtghsdiepctl1FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdiepctl1Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl1FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl1FieldUsbaepShift = 15
	RegisterOtghsdiepctl1FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdiepctl1Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdiepctl1Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl1FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdiepctl1FieldEonumdpidShift = 16
	RegisterOtghsdiepctl1FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even/odd frame
func (r *registerOtghsdiepctl1Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl1FieldNakstsShift = 17
	RegisterOtghsdiepctl1FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdiepctl1Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl1FieldEptypShift = 18
	RegisterOtghsdiepctl1FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdiepctl1Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldEptypMask) >> RegisterOtghsdiepctl1FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdiepctl1Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl1FieldEptypShift))
}

const (
	RegisterOtghsdiepctl1FieldStallShift = 21
	RegisterOtghsdiepctl1FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdiepctl1Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdiepctl1Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl1FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldStallMask)
	}
}

const (
	RegisterOtghsdiepctl1FieldTxfnumShift = 22
	RegisterOtghsdiepctl1FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtghsdiepctl1Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldTxfnumMask) >> RegisterOtghsdiepctl1FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtghsdiepctl1Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl1FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl1FieldCnakShift = 26
	RegisterOtghsdiepctl1FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdiepctl1Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl1FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldCnakMask)
	}
}

const (
	RegisterOtghsdiepctl1FieldSnakShift = 27
	RegisterOtghsdiepctl1FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdiepctl1Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl1FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldSnakMask)
	}
}

const (
	RegisterOtghsdiepctl1FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdiepctl1FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID
func (r *registerOtghsdiepctl1Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl1FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdiepctl1FieldSoddfrmShift = 29
	RegisterOtghsdiepctl1FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdiepctl1Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl1FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdiepctl1FieldEpdisShift = 30
	RegisterOtghsdiepctl1FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdiepctl1Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdiepctl1Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl1FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldEpdisMask)
	}
}

const (
	RegisterOtghsdiepctl1FieldEpenaShift = 31
	RegisterOtghsdiepctl1FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdiepctl1Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdiepctl1Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl1FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldEpenaMask)
	}
}

// registerOtghsdiepint1Type OTG device endpoint-1 interrupt register
type registerOtghsdiepint1Type uint32

const (
	RegisterOtghsdiepint1FieldXfrcShift = 0
	RegisterOtghsdiepint1FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint1Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint1Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint1FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint1FieldXfrcMask)
	}
}

const (
	RegisterOtghsdiepint1FieldEpdisdShift = 1
	RegisterOtghsdiepint1FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint1Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint1Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint1FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint1FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdiepint1FieldTocShift = 3
	RegisterOtghsdiepint1FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtghsdiepint1Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtghsdiepint1Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint1FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint1FieldTocMask)
	}
}

const (
	RegisterOtghsdiepint1FieldIttxfeShift = 4
	RegisterOtghsdiepint1FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint1Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint1Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint1FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint1FieldIttxfeMask)
	}
}

const (
	RegisterOtghsdiepint1FieldInepneShift = 6
	RegisterOtghsdiepint1FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint1Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint1Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint1FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint1FieldInepneMask)
	}
}

const (
	RegisterOtghsdiepint1FieldTxfeShift = 7
	RegisterOtghsdiepint1FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtghsdiepint1Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint1FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint1FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint1Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint1Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint1FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint1FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtghsdiepint1FieldBnaShift = 9
	RegisterOtghsdiepint1FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtghsdiepint1Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtghsdiepint1Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint1FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint1FieldBnaMask)
	}
}

const (
	RegisterOtghsdiepint1FieldPktdrpstsShift = 11
	RegisterOtghsdiepint1FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint1Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint1Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint1FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint1FieldPktdrpstsMask)
	}
}

const (
	RegisterOtghsdiepint1FieldBerrShift = 12
	RegisterOtghsdiepint1FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtghsdiepint1Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtghsdiepint1Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint1FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint1FieldBerrMask)
	}
}

const (
	RegisterOtghsdiepint1FieldNakShift = 13
	RegisterOtghsdiepint1FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtghsdiepint1Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtghsdiepint1Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint1FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint1FieldNakMask)
	}
}

// registerOtghsdieptsiz1Type OTG_HS device endpoint transfer size register
type registerOtghsdieptsiz1Type uint32

const (
	RegisterOtghsdieptsiz1FieldXfrsizShift = 0
	RegisterOtghsdieptsiz1FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdieptsiz1Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz1FieldXfrsizMask) >> RegisterOtghsdieptsiz1FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdieptsiz1Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz1FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz1FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz1FieldPktcntShift = 19
	RegisterOtghsdieptsiz1FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdieptsiz1Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz1FieldPktcntMask) >> RegisterOtghsdieptsiz1FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdieptsiz1Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz1FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz1FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz1FieldMcntShift = 29
	RegisterOtghsdieptsiz1FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtghsdieptsiz1Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz1FieldMcntMask) >> RegisterOtghsdieptsiz1FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtghsdieptsiz1Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz1FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz1FieldMcntShift))
}

// registerOtghsdiepdma2Type OTG_HS device endpoint-2 DMA address register
type registerOtghsdiepdma2Type uint32

const (
	RegisterOtghsdiepdma2FieldDmaaddrShift = 0
	RegisterOtghsdiepdma2FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghsdiepdma2Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepdma2FieldDmaaddrMask) >> RegisterOtghsdiepdma2FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghsdiepdma2Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepdma2FieldDmaaddrMask)|(uint32(value)<<RegisterOtghsdiepdma2FieldDmaaddrShift))
}

// registerOtghsdtxfsts1Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtghsdtxfsts1Type uint32

const (
	RegisterOtghsdtxfsts1FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts1FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts1Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts1FieldIneptfsavMask) >> RegisterOtghsdtxfsts1FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts1Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts1FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts1FieldIneptfsavShift))
}

// registerOtghsdiepctl2Type OTG device endpoint-2 control register
type registerOtghsdiepctl2Type uint32

const (
	RegisterOtghsdiepctl2FieldMpsizShift = 0
	RegisterOtghsdiepctl2FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdiepctl2Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldMpsizMask) >> RegisterOtghsdiepctl2FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdiepctl2Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl2FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl2FieldUsbaepShift = 15
	RegisterOtghsdiepctl2FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdiepctl2Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdiepctl2Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl2FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdiepctl2FieldEonumdpidShift = 16
	RegisterOtghsdiepctl2FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even/odd frame
func (r *registerOtghsdiepctl2Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl2FieldNakstsShift = 17
	RegisterOtghsdiepctl2FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdiepctl2Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl2FieldEptypShift = 18
	RegisterOtghsdiepctl2FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdiepctl2Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldEptypMask) >> RegisterOtghsdiepctl2FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdiepctl2Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl2FieldEptypShift))
}

const (
	RegisterOtghsdiepctl2FieldStallShift = 21
	RegisterOtghsdiepctl2FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdiepctl2Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdiepctl2Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl2FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldStallMask)
	}
}

const (
	RegisterOtghsdiepctl2FieldTxfnumShift = 22
	RegisterOtghsdiepctl2FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtghsdiepctl2Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldTxfnumMask) >> RegisterOtghsdiepctl2FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtghsdiepctl2Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl2FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl2FieldCnakShift = 26
	RegisterOtghsdiepctl2FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdiepctl2Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl2FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldCnakMask)
	}
}

const (
	RegisterOtghsdiepctl2FieldSnakShift = 27
	RegisterOtghsdiepctl2FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdiepctl2Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl2FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldSnakMask)
	}
}

const (
	RegisterOtghsdiepctl2FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdiepctl2FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID
func (r *registerOtghsdiepctl2Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl2FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdiepctl2FieldSoddfrmShift = 29
	RegisterOtghsdiepctl2FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdiepctl2Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl2FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdiepctl2FieldEpdisShift = 30
	RegisterOtghsdiepctl2FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdiepctl2Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdiepctl2Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl2FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldEpdisMask)
	}
}

const (
	RegisterOtghsdiepctl2FieldEpenaShift = 31
	RegisterOtghsdiepctl2FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdiepctl2Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdiepctl2Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl2FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldEpenaMask)
	}
}

// registerOtghsdiepint2Type OTG device endpoint-2 interrupt register
type registerOtghsdiepint2Type uint32

const (
	RegisterOtghsdiepint2FieldXfrcShift = 0
	RegisterOtghsdiepint2FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint2Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint2Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint2FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint2FieldXfrcMask)
	}
}

const (
	RegisterOtghsdiepint2FieldEpdisdShift = 1
	RegisterOtghsdiepint2FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint2Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint2Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint2FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint2FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdiepint2FieldTocShift = 3
	RegisterOtghsdiepint2FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtghsdiepint2Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtghsdiepint2Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint2FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint2FieldTocMask)
	}
}

const (
	RegisterOtghsdiepint2FieldIttxfeShift = 4
	RegisterOtghsdiepint2FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint2Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint2Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint2FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint2FieldIttxfeMask)
	}
}

const (
	RegisterOtghsdiepint2FieldInepneShift = 6
	RegisterOtghsdiepint2FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint2Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint2Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint2FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint2FieldInepneMask)
	}
}

const (
	RegisterOtghsdiepint2FieldTxfeShift = 7
	RegisterOtghsdiepint2FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtghsdiepint2Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint2FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint2FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint2Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint2Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint2FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint2FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtghsdiepint2FieldBnaShift = 9
	RegisterOtghsdiepint2FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtghsdiepint2Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtghsdiepint2Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint2FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint2FieldBnaMask)
	}
}

const (
	RegisterOtghsdiepint2FieldPktdrpstsShift = 11
	RegisterOtghsdiepint2FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint2Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint2Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint2FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint2FieldPktdrpstsMask)
	}
}

const (
	RegisterOtghsdiepint2FieldBerrShift = 12
	RegisterOtghsdiepint2FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtghsdiepint2Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtghsdiepint2Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint2FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint2FieldBerrMask)
	}
}

const (
	RegisterOtghsdiepint2FieldNakShift = 13
	RegisterOtghsdiepint2FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtghsdiepint2Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtghsdiepint2Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint2FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint2FieldNakMask)
	}
}

// registerOtghsdieptsiz2Type OTG_HS device endpoint transfer size register
type registerOtghsdieptsiz2Type uint32

const (
	RegisterOtghsdieptsiz2FieldXfrsizShift = 0
	RegisterOtghsdieptsiz2FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdieptsiz2Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz2FieldXfrsizMask) >> RegisterOtghsdieptsiz2FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdieptsiz2Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz2FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz2FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz2FieldPktcntShift = 19
	RegisterOtghsdieptsiz2FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdieptsiz2Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz2FieldPktcntMask) >> RegisterOtghsdieptsiz2FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdieptsiz2Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz2FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz2FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz2FieldMcntShift = 29
	RegisterOtghsdieptsiz2FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtghsdieptsiz2Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz2FieldMcntMask) >> RegisterOtghsdieptsiz2FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtghsdieptsiz2Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz2FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz2FieldMcntShift))
}

// registerOtghsdiepdma3Type OTG_HS device endpoint-3 DMA address register
type registerOtghsdiepdma3Type uint32

const (
	RegisterOtghsdiepdma3FieldDmaaddrShift = 0
	RegisterOtghsdiepdma3FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghsdiepdma3Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepdma3FieldDmaaddrMask) >> RegisterOtghsdiepdma3FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghsdiepdma3Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepdma3FieldDmaaddrMask)|(uint32(value)<<RegisterOtghsdiepdma3FieldDmaaddrShift))
}

// registerOtghsdtxfsts2Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtghsdtxfsts2Type uint32

const (
	RegisterOtghsdtxfsts2FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts2FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts2Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts2FieldIneptfsavMask) >> RegisterOtghsdtxfsts2FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts2Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts2FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts2FieldIneptfsavShift))
}

// registerOtghsdiepctl3Type OTG device endpoint-3 control register
type registerOtghsdiepctl3Type uint32

const (
	RegisterOtghsdiepctl3FieldMpsizShift = 0
	RegisterOtghsdiepctl3FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdiepctl3Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldMpsizMask) >> RegisterOtghsdiepctl3FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdiepctl3Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl3FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl3FieldUsbaepShift = 15
	RegisterOtghsdiepctl3FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdiepctl3Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdiepctl3Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl3FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdiepctl3FieldEonumdpidShift = 16
	RegisterOtghsdiepctl3FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even/odd frame
func (r *registerOtghsdiepctl3Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl3FieldNakstsShift = 17
	RegisterOtghsdiepctl3FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdiepctl3Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl3FieldEptypShift = 18
	RegisterOtghsdiepctl3FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdiepctl3Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldEptypMask) >> RegisterOtghsdiepctl3FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdiepctl3Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl3FieldEptypShift))
}

const (
	RegisterOtghsdiepctl3FieldStallShift = 21
	RegisterOtghsdiepctl3FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdiepctl3Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdiepctl3Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl3FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldStallMask)
	}
}

const (
	RegisterOtghsdiepctl3FieldTxfnumShift = 22
	RegisterOtghsdiepctl3FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtghsdiepctl3Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldTxfnumMask) >> RegisterOtghsdiepctl3FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtghsdiepctl3Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl3FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl3FieldCnakShift = 26
	RegisterOtghsdiepctl3FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdiepctl3Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl3FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldCnakMask)
	}
}

const (
	RegisterOtghsdiepctl3FieldSnakShift = 27
	RegisterOtghsdiepctl3FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdiepctl3Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl3FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldSnakMask)
	}
}

const (
	RegisterOtghsdiepctl3FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdiepctl3FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID
func (r *registerOtghsdiepctl3Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl3FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdiepctl3FieldSoddfrmShift = 29
	RegisterOtghsdiepctl3FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdiepctl3Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl3FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdiepctl3FieldEpdisShift = 30
	RegisterOtghsdiepctl3FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdiepctl3Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdiepctl3Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl3FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldEpdisMask)
	}
}

const (
	RegisterOtghsdiepctl3FieldEpenaShift = 31
	RegisterOtghsdiepctl3FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdiepctl3Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdiepctl3Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl3FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldEpenaMask)
	}
}

// registerOtghsdiepint3Type OTG device endpoint-3 interrupt register
type registerOtghsdiepint3Type uint32

const (
	RegisterOtghsdiepint3FieldXfrcShift = 0
	RegisterOtghsdiepint3FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint3Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint3Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint3FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint3FieldXfrcMask)
	}
}

const (
	RegisterOtghsdiepint3FieldEpdisdShift = 1
	RegisterOtghsdiepint3FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint3Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint3Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint3FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint3FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdiepint3FieldTocShift = 3
	RegisterOtghsdiepint3FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtghsdiepint3Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtghsdiepint3Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint3FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint3FieldTocMask)
	}
}

const (
	RegisterOtghsdiepint3FieldIttxfeShift = 4
	RegisterOtghsdiepint3FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint3Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint3Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint3FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint3FieldIttxfeMask)
	}
}

const (
	RegisterOtghsdiepint3FieldInepneShift = 6
	RegisterOtghsdiepint3FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint3Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint3Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint3FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint3FieldInepneMask)
	}
}

const (
	RegisterOtghsdiepint3FieldTxfeShift = 7
	RegisterOtghsdiepint3FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtghsdiepint3Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint3FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint3FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint3Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint3Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint3FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint3FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtghsdiepint3FieldBnaShift = 9
	RegisterOtghsdiepint3FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtghsdiepint3Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtghsdiepint3Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint3FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint3FieldBnaMask)
	}
}

const (
	RegisterOtghsdiepint3FieldPktdrpstsShift = 11
	RegisterOtghsdiepint3FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint3Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint3Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint3FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint3FieldPktdrpstsMask)
	}
}

const (
	RegisterOtghsdiepint3FieldBerrShift = 12
	RegisterOtghsdiepint3FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtghsdiepint3Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtghsdiepint3Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint3FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint3FieldBerrMask)
	}
}

const (
	RegisterOtghsdiepint3FieldNakShift = 13
	RegisterOtghsdiepint3FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtghsdiepint3Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtghsdiepint3Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint3FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint3FieldNakMask)
	}
}

// registerOtghsdieptsiz3Type OTG_HS device endpoint transfer size register
type registerOtghsdieptsiz3Type uint32

const (
	RegisterOtghsdieptsiz3FieldXfrsizShift = 0
	RegisterOtghsdieptsiz3FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdieptsiz3Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz3FieldXfrsizMask) >> RegisterOtghsdieptsiz3FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdieptsiz3Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz3FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz3FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz3FieldPktcntShift = 19
	RegisterOtghsdieptsiz3FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdieptsiz3Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz3FieldPktcntMask) >> RegisterOtghsdieptsiz3FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdieptsiz3Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz3FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz3FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz3FieldMcntShift = 29
	RegisterOtghsdieptsiz3FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtghsdieptsiz3Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz3FieldMcntMask) >> RegisterOtghsdieptsiz3FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtghsdieptsiz3Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz3FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz3FieldMcntShift))
}

// registerOtghsdiepdma4Type OTG_HS device endpoint-4 DMA address register
type registerOtghsdiepdma4Type uint32

const (
	RegisterOtghsdiepdma4FieldDmaaddrShift = 0
	RegisterOtghsdiepdma4FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghsdiepdma4Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepdma4FieldDmaaddrMask) >> RegisterOtghsdiepdma4FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghsdiepdma4Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepdma4FieldDmaaddrMask)|(uint32(value)<<RegisterOtghsdiepdma4FieldDmaaddrShift))
}

// registerOtghsdtxfsts3Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtghsdtxfsts3Type uint32

const (
	RegisterOtghsdtxfsts3FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts3FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts3Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts3FieldIneptfsavMask) >> RegisterOtghsdtxfsts3FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts3Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts3FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts3FieldIneptfsavShift))
}

// registerOtghsdiepctl4Type OTG device endpoint-4 control register
type registerOtghsdiepctl4Type uint32

const (
	RegisterOtghsdiepctl4FieldMpsizShift = 0
	RegisterOtghsdiepctl4FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdiepctl4Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldMpsizMask) >> RegisterOtghsdiepctl4FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdiepctl4Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl4FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl4FieldUsbaepShift = 15
	RegisterOtghsdiepctl4FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdiepctl4Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdiepctl4Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl4FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdiepctl4FieldEonumdpidShift = 16
	RegisterOtghsdiepctl4FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even/odd frame
func (r *registerOtghsdiepctl4Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl4FieldNakstsShift = 17
	RegisterOtghsdiepctl4FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdiepctl4Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl4FieldEptypShift = 18
	RegisterOtghsdiepctl4FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdiepctl4Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldEptypMask) >> RegisterOtghsdiepctl4FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdiepctl4Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl4FieldEptypShift))
}

const (
	RegisterOtghsdiepctl4FieldStallShift = 21
	RegisterOtghsdiepctl4FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdiepctl4Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdiepctl4Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl4FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldStallMask)
	}
}

const (
	RegisterOtghsdiepctl4FieldTxfnumShift = 22
	RegisterOtghsdiepctl4FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtghsdiepctl4Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldTxfnumMask) >> RegisterOtghsdiepctl4FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtghsdiepctl4Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl4FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl4FieldCnakShift = 26
	RegisterOtghsdiepctl4FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdiepctl4Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl4FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldCnakMask)
	}
}

const (
	RegisterOtghsdiepctl4FieldSnakShift = 27
	RegisterOtghsdiepctl4FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdiepctl4Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl4FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldSnakMask)
	}
}

const (
	RegisterOtghsdiepctl4FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdiepctl4FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID
func (r *registerOtghsdiepctl4Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl4FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdiepctl4FieldSoddfrmShift = 29
	RegisterOtghsdiepctl4FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdiepctl4Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl4FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdiepctl4FieldEpdisShift = 30
	RegisterOtghsdiepctl4FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdiepctl4Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdiepctl4Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl4FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldEpdisMask)
	}
}

const (
	RegisterOtghsdiepctl4FieldEpenaShift = 31
	RegisterOtghsdiepctl4FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdiepctl4Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdiepctl4Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl4FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldEpenaMask)
	}
}

// registerOtghsdiepint4Type OTG device endpoint-4 interrupt register
type registerOtghsdiepint4Type uint32

const (
	RegisterOtghsdiepint4FieldXfrcShift = 0
	RegisterOtghsdiepint4FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint4Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint4Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint4FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint4FieldXfrcMask)
	}
}

const (
	RegisterOtghsdiepint4FieldEpdisdShift = 1
	RegisterOtghsdiepint4FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint4Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint4Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint4FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint4FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdiepint4FieldTocShift = 3
	RegisterOtghsdiepint4FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtghsdiepint4Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtghsdiepint4Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint4FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint4FieldTocMask)
	}
}

const (
	RegisterOtghsdiepint4FieldIttxfeShift = 4
	RegisterOtghsdiepint4FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint4Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint4Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint4FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint4FieldIttxfeMask)
	}
}

const (
	RegisterOtghsdiepint4FieldInepneShift = 6
	RegisterOtghsdiepint4FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint4Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint4Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint4FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint4FieldInepneMask)
	}
}

const (
	RegisterOtghsdiepint4FieldTxfeShift = 7
	RegisterOtghsdiepint4FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtghsdiepint4Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint4FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint4FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint4Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint4Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint4FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint4FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtghsdiepint4FieldBnaShift = 9
	RegisterOtghsdiepint4FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtghsdiepint4Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtghsdiepint4Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint4FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint4FieldBnaMask)
	}
}

const (
	RegisterOtghsdiepint4FieldPktdrpstsShift = 11
	RegisterOtghsdiepint4FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint4Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint4Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint4FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint4FieldPktdrpstsMask)
	}
}

const (
	RegisterOtghsdiepint4FieldBerrShift = 12
	RegisterOtghsdiepint4FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtghsdiepint4Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtghsdiepint4Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint4FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint4FieldBerrMask)
	}
}

const (
	RegisterOtghsdiepint4FieldNakShift = 13
	RegisterOtghsdiepint4FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtghsdiepint4Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtghsdiepint4Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint4FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint4FieldNakMask)
	}
}

// registerOtghsdieptsiz4Type OTG_HS device endpoint transfer size register
type registerOtghsdieptsiz4Type uint32

const (
	RegisterOtghsdieptsiz4FieldXfrsizShift = 0
	RegisterOtghsdieptsiz4FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdieptsiz4Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz4FieldXfrsizMask) >> RegisterOtghsdieptsiz4FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdieptsiz4Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz4FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz4FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz4FieldPktcntShift = 19
	RegisterOtghsdieptsiz4FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdieptsiz4Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz4FieldPktcntMask) >> RegisterOtghsdieptsiz4FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdieptsiz4Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz4FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz4FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz4FieldMcntShift = 29
	RegisterOtghsdieptsiz4FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtghsdieptsiz4Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz4FieldMcntMask) >> RegisterOtghsdieptsiz4FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtghsdieptsiz4Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz4FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz4FieldMcntShift))
}

// registerOtghsdiepdma5Type OTG_HS device endpoint-5 DMA address register
type registerOtghsdiepdma5Type uint32

const (
	RegisterOtghsdiepdma5FieldDmaaddrShift = 0
	RegisterOtghsdiepdma5FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghsdiepdma5Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepdma5FieldDmaaddrMask) >> RegisterOtghsdiepdma5FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghsdiepdma5Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepdma5FieldDmaaddrMask)|(uint32(value)<<RegisterOtghsdiepdma5FieldDmaaddrShift))
}

// registerOtghsdtxfsts4Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtghsdtxfsts4Type uint32

const (
	RegisterOtghsdtxfsts4FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts4FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts4Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts4FieldIneptfsavMask) >> RegisterOtghsdtxfsts4FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts4Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts4FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts4FieldIneptfsavShift))
}

// registerOtghsdiepctl5Type OTG device endpoint-5 control register
type registerOtghsdiepctl5Type uint32

const (
	RegisterOtghsdiepctl5FieldMpsizShift = 0
	RegisterOtghsdiepctl5FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdiepctl5Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldMpsizMask) >> RegisterOtghsdiepctl5FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdiepctl5Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl5FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl5FieldUsbaepShift = 15
	RegisterOtghsdiepctl5FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdiepctl5Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdiepctl5Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl5FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdiepctl5FieldEonumdpidShift = 16
	RegisterOtghsdiepctl5FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even/odd frame
func (r *registerOtghsdiepctl5Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl5FieldNakstsShift = 17
	RegisterOtghsdiepctl5FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdiepctl5Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl5FieldEptypShift = 18
	RegisterOtghsdiepctl5FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdiepctl5Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldEptypMask) >> RegisterOtghsdiepctl5FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdiepctl5Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl5FieldEptypShift))
}

const (
	RegisterOtghsdiepctl5FieldStallShift = 21
	RegisterOtghsdiepctl5FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdiepctl5Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdiepctl5Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl5FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldStallMask)
	}
}

const (
	RegisterOtghsdiepctl5FieldTxfnumShift = 22
	RegisterOtghsdiepctl5FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtghsdiepctl5Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldTxfnumMask) >> RegisterOtghsdiepctl5FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtghsdiepctl5Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl5FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl5FieldCnakShift = 26
	RegisterOtghsdiepctl5FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdiepctl5Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl5FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldCnakMask)
	}
}

const (
	RegisterOtghsdiepctl5FieldSnakShift = 27
	RegisterOtghsdiepctl5FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdiepctl5Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl5FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldSnakMask)
	}
}

const (
	RegisterOtghsdiepctl5FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdiepctl5FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID
func (r *registerOtghsdiepctl5Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl5FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdiepctl5FieldSoddfrmShift = 29
	RegisterOtghsdiepctl5FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdiepctl5Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl5FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdiepctl5FieldEpdisShift = 30
	RegisterOtghsdiepctl5FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdiepctl5Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdiepctl5Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl5FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldEpdisMask)
	}
}

const (
	RegisterOtghsdiepctl5FieldEpenaShift = 31
	RegisterOtghsdiepctl5FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdiepctl5Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdiepctl5Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl5FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldEpenaMask)
	}
}

// registerOtghsdieptsiz6Type OTG_HS device endpoint transfer size register
type registerOtghsdieptsiz6Type uint32

const (
	RegisterOtghsdieptsiz6FieldXfrsizShift = 0
	RegisterOtghsdieptsiz6FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdieptsiz6Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz6FieldXfrsizMask) >> RegisterOtghsdieptsiz6FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdieptsiz6Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz6FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz6FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz6FieldPktcntShift = 19
	RegisterOtghsdieptsiz6FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdieptsiz6Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz6FieldPktcntMask) >> RegisterOtghsdieptsiz6FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdieptsiz6Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz6FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz6FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz6FieldMcntShift = 29
	RegisterOtghsdieptsiz6FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtghsdieptsiz6Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz6FieldMcntMask) >> RegisterOtghsdieptsiz6FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtghsdieptsiz6Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz6FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz6FieldMcntShift))
}

// registerOtghsdtxfsts6Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtghsdtxfsts6Type uint32

const (
	RegisterOtghsdtxfsts6FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts6FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts6Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts6FieldIneptfsavMask) >> RegisterOtghsdtxfsts6FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts6Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts6FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts6FieldIneptfsavShift))
}

// registerOtghsdieptsiz7Type OTG_HS device endpoint transfer size register
type registerOtghsdieptsiz7Type uint32

const (
	RegisterOtghsdieptsiz7FieldXfrsizShift = 0
	RegisterOtghsdieptsiz7FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdieptsiz7Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz7FieldXfrsizMask) >> RegisterOtghsdieptsiz7FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdieptsiz7Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz7FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz7FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz7FieldPktcntShift = 19
	RegisterOtghsdieptsiz7FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdieptsiz7Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz7FieldPktcntMask) >> RegisterOtghsdieptsiz7FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdieptsiz7Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz7FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz7FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz7FieldMcntShift = 29
	RegisterOtghsdieptsiz7FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtghsdieptsiz7Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz7FieldMcntMask) >> RegisterOtghsdieptsiz7FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtghsdieptsiz7Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz7FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz7FieldMcntShift))
}

// registerOtghsdiepint5Type OTG device endpoint-5 interrupt register
type registerOtghsdiepint5Type uint32

const (
	RegisterOtghsdiepint5FieldXfrcShift = 0
	RegisterOtghsdiepint5FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint5Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint5Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint5FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint5FieldXfrcMask)
	}
}

const (
	RegisterOtghsdiepint5FieldEpdisdShift = 1
	RegisterOtghsdiepint5FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint5Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint5Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint5FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint5FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdiepint5FieldTocShift = 3
	RegisterOtghsdiepint5FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtghsdiepint5Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtghsdiepint5Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint5FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint5FieldTocMask)
	}
}

const (
	RegisterOtghsdiepint5FieldIttxfeShift = 4
	RegisterOtghsdiepint5FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint5Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint5Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint5FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint5FieldIttxfeMask)
	}
}

const (
	RegisterOtghsdiepint5FieldInepneShift = 6
	RegisterOtghsdiepint5FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint5Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint5Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint5FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint5FieldInepneMask)
	}
}

const (
	RegisterOtghsdiepint5FieldTxfeShift = 7
	RegisterOtghsdiepint5FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtghsdiepint5Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint5FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint5FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint5Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint5Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint5FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint5FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtghsdiepint5FieldBnaShift = 9
	RegisterOtghsdiepint5FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtghsdiepint5Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtghsdiepint5Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint5FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint5FieldBnaMask)
	}
}

const (
	RegisterOtghsdiepint5FieldPktdrpstsShift = 11
	RegisterOtghsdiepint5FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint5Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint5Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint5FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint5FieldPktdrpstsMask)
	}
}

const (
	RegisterOtghsdiepint5FieldBerrShift = 12
	RegisterOtghsdiepint5FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtghsdiepint5Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtghsdiepint5Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint5FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint5FieldBerrMask)
	}
}

const (
	RegisterOtghsdiepint5FieldNakShift = 13
	RegisterOtghsdiepint5FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtghsdiepint5Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtghsdiepint5Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint5FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint5FieldNakMask)
	}
}

// registerOtghsdtxfsts7Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtghsdtxfsts7Type uint32

const (
	RegisterOtghsdtxfsts7FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts7FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts7Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts7FieldIneptfsavMask) >> RegisterOtghsdtxfsts7FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts7Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts7FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts7FieldIneptfsavShift))
}

// registerOtghsdieptsiz5Type OTG_HS device endpoint transfer size register
type registerOtghsdieptsiz5Type uint32

const (
	RegisterOtghsdieptsiz5FieldXfrsizShift = 0
	RegisterOtghsdieptsiz5FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdieptsiz5Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz5FieldXfrsizMask) >> RegisterOtghsdieptsiz5FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdieptsiz5Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz5FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz5FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz5FieldPktcntShift = 19
	RegisterOtghsdieptsiz5FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdieptsiz5Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz5FieldPktcntMask) >> RegisterOtghsdieptsiz5FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdieptsiz5Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz5FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz5FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz5FieldMcntShift = 29
	RegisterOtghsdieptsiz5FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *registerOtghsdieptsiz5Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz5FieldMcntMask) >> RegisterOtghsdieptsiz5FieldMcntShift)
}

// SetMcnt Multi count
func (r *registerOtghsdieptsiz5Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz5FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz5FieldMcntShift))
}

// registerOtghsdtxfsts5Type OTG_HS device IN endpoint transmit FIFO status register
type registerOtghsdtxfsts5Type uint32

const (
	RegisterOtghsdtxfsts5FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts5FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts5Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts5FieldIneptfsavMask) >> RegisterOtghsdtxfsts5FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *registerOtghsdtxfsts5Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts5FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts5FieldIneptfsavShift))
}

// registerOtghsdiepctl6Type OTG device endpoint-6 control register
type registerOtghsdiepctl6Type uint32

const (
	RegisterOtghsdiepctl6FieldMpsizShift = 0
	RegisterOtghsdiepctl6FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdiepctl6Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldMpsizMask) >> RegisterOtghsdiepctl6FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdiepctl6Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl6FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl6FieldUsbaepShift = 15
	RegisterOtghsdiepctl6FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdiepctl6Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdiepctl6Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl6FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdiepctl6FieldEonumdpidShift = 16
	RegisterOtghsdiepctl6FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even/odd frame
func (r *registerOtghsdiepctl6Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl6FieldNakstsShift = 17
	RegisterOtghsdiepctl6FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdiepctl6Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl6FieldEptypShift = 18
	RegisterOtghsdiepctl6FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdiepctl6Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldEptypMask) >> RegisterOtghsdiepctl6FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdiepctl6Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl6FieldEptypShift))
}

const (
	RegisterOtghsdiepctl6FieldStallShift = 21
	RegisterOtghsdiepctl6FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdiepctl6Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdiepctl6Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl6FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldStallMask)
	}
}

const (
	RegisterOtghsdiepctl6FieldTxfnumShift = 22
	RegisterOtghsdiepctl6FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtghsdiepctl6Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldTxfnumMask) >> RegisterOtghsdiepctl6FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtghsdiepctl6Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl6FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl6FieldCnakShift = 26
	RegisterOtghsdiepctl6FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdiepctl6Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl6FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldCnakMask)
	}
}

const (
	RegisterOtghsdiepctl6FieldSnakShift = 27
	RegisterOtghsdiepctl6FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdiepctl6Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl6FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldSnakMask)
	}
}

const (
	RegisterOtghsdiepctl6FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdiepctl6FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID
func (r *registerOtghsdiepctl6Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl6FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdiepctl6FieldSoddfrmShift = 29
	RegisterOtghsdiepctl6FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdiepctl6Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl6FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdiepctl6FieldEpdisShift = 30
	RegisterOtghsdiepctl6FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdiepctl6Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdiepctl6Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl6FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldEpdisMask)
	}
}

const (
	RegisterOtghsdiepctl6FieldEpenaShift = 31
	RegisterOtghsdiepctl6FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdiepctl6Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdiepctl6Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl6FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldEpenaMask)
	}
}

// registerOtghsdiepint6Type OTG device endpoint-6 interrupt register
type registerOtghsdiepint6Type uint32

const (
	RegisterOtghsdiepint6FieldXfrcShift = 0
	RegisterOtghsdiepint6FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint6Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint6Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint6FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint6FieldXfrcMask)
	}
}

const (
	RegisterOtghsdiepint6FieldEpdisdShift = 1
	RegisterOtghsdiepint6FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint6Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint6Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint6FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint6FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdiepint6FieldTocShift = 3
	RegisterOtghsdiepint6FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtghsdiepint6Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtghsdiepint6Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint6FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint6FieldTocMask)
	}
}

const (
	RegisterOtghsdiepint6FieldIttxfeShift = 4
	RegisterOtghsdiepint6FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint6Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint6Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint6FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint6FieldIttxfeMask)
	}
}

const (
	RegisterOtghsdiepint6FieldInepneShift = 6
	RegisterOtghsdiepint6FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint6Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint6Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint6FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint6FieldInepneMask)
	}
}

const (
	RegisterOtghsdiepint6FieldTxfeShift = 7
	RegisterOtghsdiepint6FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtghsdiepint6Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint6FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint6FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint6Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint6Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint6FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint6FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtghsdiepint6FieldBnaShift = 9
	RegisterOtghsdiepint6FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtghsdiepint6Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtghsdiepint6Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint6FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint6FieldBnaMask)
	}
}

const (
	RegisterOtghsdiepint6FieldPktdrpstsShift = 11
	RegisterOtghsdiepint6FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint6Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint6Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint6FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint6FieldPktdrpstsMask)
	}
}

const (
	RegisterOtghsdiepint6FieldBerrShift = 12
	RegisterOtghsdiepint6FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtghsdiepint6Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtghsdiepint6Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint6FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint6FieldBerrMask)
	}
}

const (
	RegisterOtghsdiepint6FieldNakShift = 13
	RegisterOtghsdiepint6FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtghsdiepint6Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtghsdiepint6Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint6FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint6FieldNakMask)
	}
}

// registerOtghsdiepctl7Type OTG device endpoint-7 control register
type registerOtghsdiepctl7Type uint32

const (
	RegisterOtghsdiepctl7FieldMpsizShift = 0
	RegisterOtghsdiepctl7FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdiepctl7Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldMpsizMask) >> RegisterOtghsdiepctl7FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdiepctl7Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl7FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl7FieldUsbaepShift = 15
	RegisterOtghsdiepctl7FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdiepctl7Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdiepctl7Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl7FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdiepctl7FieldEonumdpidShift = 16
	RegisterOtghsdiepctl7FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even/odd frame
func (r *registerOtghsdiepctl7Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl7FieldNakstsShift = 17
	RegisterOtghsdiepctl7FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdiepctl7Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl7FieldEptypShift = 18
	RegisterOtghsdiepctl7FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdiepctl7Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldEptypMask) >> RegisterOtghsdiepctl7FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdiepctl7Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl7FieldEptypShift))
}

const (
	RegisterOtghsdiepctl7FieldStallShift = 21
	RegisterOtghsdiepctl7FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdiepctl7Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdiepctl7Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl7FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldStallMask)
	}
}

const (
	RegisterOtghsdiepctl7FieldTxfnumShift = 22
	RegisterOtghsdiepctl7FieldTxfnumMask  = 0x3c00000
)

// GetTxfnum TxFIFO number
func (r *registerOtghsdiepctl7Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldTxfnumMask) >> RegisterOtghsdiepctl7FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtghsdiepctl7Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl7FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl7FieldCnakShift = 26
	RegisterOtghsdiepctl7FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdiepctl7Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl7FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldCnakMask)
	}
}

const (
	RegisterOtghsdiepctl7FieldSnakShift = 27
	RegisterOtghsdiepctl7FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdiepctl7Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl7FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldSnakMask)
	}
}

const (
	RegisterOtghsdiepctl7FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdiepctl7FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID
func (r *registerOtghsdiepctl7Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl7FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdiepctl7FieldSoddfrmShift = 29
	RegisterOtghsdiepctl7FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdiepctl7Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl7FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdiepctl7FieldEpdisShift = 30
	RegisterOtghsdiepctl7FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdiepctl7Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdiepctl7Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl7FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldEpdisMask)
	}
}

const (
	RegisterOtghsdiepctl7FieldEpenaShift = 31
	RegisterOtghsdiepctl7FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdiepctl7Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdiepctl7Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl7FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldEpenaMask)
	}
}

// registerOtghsdiepint7Type OTG device endpoint-7 interrupt register
type registerOtghsdiepint7Type uint32

const (
	RegisterOtghsdiepint7FieldXfrcShift = 0
	RegisterOtghsdiepint7FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint7Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdiepint7Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint7FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint7FieldXfrcMask)
	}
}

const (
	RegisterOtghsdiepint7FieldEpdisdShift = 1
	RegisterOtghsdiepint7FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint7Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdiepint7Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint7FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint7FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdiepint7FieldTocShift = 3
	RegisterOtghsdiepint7FieldTocMask  = 0x8
)

// GetToc Timeout condition
func (r *registerOtghsdiepint7Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *registerOtghsdiepint7Type) SetToc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint7FieldTocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint7FieldTocMask)
	}
}

const (
	RegisterOtghsdiepint7FieldIttxfeShift = 4
	RegisterOtghsdiepint7FieldIttxfeMask  = 0x10
)

// GetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint7Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *registerOtghsdiepint7Type) SetIttxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint7FieldIttxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint7FieldIttxfeMask)
	}
}

const (
	RegisterOtghsdiepint7FieldInepneShift = 6
	RegisterOtghsdiepint7FieldInepneMask  = 0x40
)

// GetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint7Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *registerOtghsdiepint7Type) SetInepne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint7FieldInepneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint7FieldInepneMask)
	}
}

const (
	RegisterOtghsdiepint7FieldTxfeShift = 7
	RegisterOtghsdiepint7FieldTxfeMask  = 0x80
)

// GetTxfe Transmit FIFO empty
func (r *registerOtghsdiepint7Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint7FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint7FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint7Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *registerOtghsdiepint7Type) SetTxfifoudrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint7FieldTxfifoudrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint7FieldTxfifoudrnMask)
	}
}

const (
	RegisterOtghsdiepint7FieldBnaShift = 9
	RegisterOtghsdiepint7FieldBnaMask  = 0x200
)

// GetBna Buffer not available interrupt
func (r *registerOtghsdiepint7Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *registerOtghsdiepint7Type) SetBna(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint7FieldBnaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint7FieldBnaMask)
	}
}

const (
	RegisterOtghsdiepint7FieldPktdrpstsShift = 11
	RegisterOtghsdiepint7FieldPktdrpstsMask  = 0x800
)

// GetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint7Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *registerOtghsdiepint7Type) SetPktdrpsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint7FieldPktdrpstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint7FieldPktdrpstsMask)
	}
}

const (
	RegisterOtghsdiepint7FieldBerrShift = 12
	RegisterOtghsdiepint7FieldBerrMask  = 0x1000
)

// GetBerr Babble error interrupt
func (r *registerOtghsdiepint7Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *registerOtghsdiepint7Type) SetBerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint7FieldBerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint7FieldBerrMask)
	}
}

const (
	RegisterOtghsdiepint7FieldNakShift = 13
	RegisterOtghsdiepint7FieldNakMask  = 0x2000
)

// GetNak NAK interrupt
func (r *registerOtghsdiepint7Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *registerOtghsdiepint7Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint7FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint7FieldNakMask)
	}
}

// registerOtghsdoepctl0Type OTG_HS device control OUT endpoint 0 control register
type registerOtghsdoepctl0Type uint32

const (
	RegisterOtghsdoepctl0FieldMpsizShift = 0
	RegisterOtghsdoepctl0FieldMpsizMask  = 0x3
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdoepctl0Type) GetMpsiz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldMpsizMask) >> RegisterOtghsdoepctl0FieldMpsizShift)
}

const (
	RegisterOtghsdoepctl0FieldUsbaepShift = 15
	RegisterOtghsdoepctl0FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdoepctl0Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldUsbaepMask) != 0
}

const (
	RegisterOtghsdoepctl0FieldNakstsShift = 17
	RegisterOtghsdoepctl0FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdoepctl0Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl0FieldEptypShift = 18
	RegisterOtghsdoepctl0FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdoepctl0Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldEptypMask) >> RegisterOtghsdoepctl0FieldEptypShift)
}

const (
	RegisterOtghsdoepctl0FieldSnpmShift = 20
	RegisterOtghsdoepctl0FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtghsdoepctl0Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtghsdoepctl0Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl0FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl0FieldSnpmMask)
	}
}

const (
	RegisterOtghsdoepctl0FieldStallShift = 21
	RegisterOtghsdoepctl0FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdoepctl0Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdoepctl0Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl0FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl0FieldStallMask)
	}
}

const (
	RegisterOtghsdoepctl0FieldCnakShift = 26
	RegisterOtghsdoepctl0FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdoepctl0Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl0FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl0FieldCnakMask)
	}
}

const (
	RegisterOtghsdoepctl0FieldSnakShift = 27
	RegisterOtghsdoepctl0FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdoepctl0Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl0FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl0FieldSnakMask)
	}
}

const (
	RegisterOtghsdoepctl0FieldEpdisShift = 30
	RegisterOtghsdoepctl0FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdoepctl0Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldEpdisMask) != 0
}

const (
	RegisterOtghsdoepctl0FieldEpenaShift = 31
	RegisterOtghsdoepctl0FieldEpenaMask  = 0x80000000
)

// SetEpena Endpoint enable
func (r *registerOtghsdoepctl0Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl0FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl0FieldEpenaMask)
	}
}

// registerOtghsdoepint0Type OTG_HS device endpoint-0 interrupt register
type registerOtghsdoepint0Type uint32

const (
	RegisterOtghsdoepint0FieldXfrcShift = 0
	RegisterOtghsdoepint0FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint0Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint0FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint0Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint0FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint0FieldXfrcMask)
	}
}

const (
	RegisterOtghsdoepint0FieldEpdisdShift = 1
	RegisterOtghsdoepint0FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint0Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint0FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint0Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint0FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint0FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdoepint0FieldStupShift = 3
	RegisterOtghsdoepint0FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtghsdoepint0Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint0FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtghsdoepint0Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint0FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint0FieldStupMask)
	}
}

const (
	RegisterOtghsdoepint0FieldOtepdisShift = 4
	RegisterOtghsdoepint0FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint0Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint0FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint0Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint0FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint0FieldOtepdisMask)
	}
}

const (
	RegisterOtghsdoepint0FieldB2bstupShift = 6
	RegisterOtghsdoepint0FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint0Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint0FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint0Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint0FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint0FieldB2bstupMask)
	}
}

const (
	RegisterOtghsdoepint0FieldNyetShift = 14
	RegisterOtghsdoepint0FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtghsdoepint0Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint0FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtghsdoepint0Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint0FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint0FieldNyetMask)
	}
}

// registerOtghsdoeptsiz0Type OTG_HS device endpoint-0 transfer size register
type registerOtghsdoeptsiz0Type uint32

const (
	RegisterOtghsdoeptsiz0FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz0FieldXfrsizMask  = 0x7f
)

// GetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz0Type) GetXfrsiz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz0FieldXfrsizMask) >> RegisterOtghsdoeptsiz0FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz0Type) SetXfrsiz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz0FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz0FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz0FieldPktcntShift = 19
	RegisterOtghsdoeptsiz0FieldPktcntMask  = 0x80000
)

// GetPktcnt Packet count
func (r *registerOtghsdoeptsiz0Type) GetPktcnt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz0FieldPktcntMask) != 0
}

// SetPktcnt Packet count
func (r *registerOtghsdoeptsiz0Type) SetPktcnt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoeptsiz0FieldPktcntMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz0FieldPktcntMask)
	}
}

const (
	RegisterOtghsdoeptsiz0FieldStupcntShift = 29
	RegisterOtghsdoeptsiz0FieldStupcntMask  = 0x60000000
)

// GetStupcnt SETUP packet count
func (r *registerOtghsdoeptsiz0Type) GetStupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz0FieldStupcntMask) >> RegisterOtghsdoeptsiz0FieldStupcntShift)
}

// SetStupcnt SETUP packet count
func (r *registerOtghsdoeptsiz0Type) SetStupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz0FieldStupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz0FieldStupcntShift))
}

// registerOtghsdoepctl1Type OTG device endpoint-1 control register
type registerOtghsdoepctl1Type uint32

const (
	RegisterOtghsdoepctl1FieldMpsizShift = 0
	RegisterOtghsdoepctl1FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdoepctl1Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldMpsizMask) >> RegisterOtghsdoepctl1FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdoepctl1Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl1FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl1FieldUsbaepShift = 15
	RegisterOtghsdoepctl1FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdoepctl1Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdoepctl1Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl1FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdoepctl1FieldEonumdpidShift = 16
	RegisterOtghsdoepctl1FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even odd frame/Endpoint data PID
func (r *registerOtghsdoepctl1Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl1FieldNakstsShift = 17
	RegisterOtghsdoepctl1FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdoepctl1Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl1FieldEptypShift = 18
	RegisterOtghsdoepctl1FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdoepctl1Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldEptypMask) >> RegisterOtghsdoepctl1FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdoepctl1Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl1FieldEptypShift))
}

const (
	RegisterOtghsdoepctl1FieldSnpmShift = 20
	RegisterOtghsdoepctl1FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtghsdoepctl1Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtghsdoepctl1Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl1FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldSnpmMask)
	}
}

const (
	RegisterOtghsdoepctl1FieldStallShift = 21
	RegisterOtghsdoepctl1FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdoepctl1Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdoepctl1Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl1FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldStallMask)
	}
}

const (
	RegisterOtghsdoepctl1FieldCnakShift = 26
	RegisterOtghsdoepctl1FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdoepctl1Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl1FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldCnakMask)
	}
}

const (
	RegisterOtghsdoepctl1FieldSnakShift = 27
	RegisterOtghsdoepctl1FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdoepctl1Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl1FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldSnakMask)
	}
}

const (
	RegisterOtghsdoepctl1FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdoepctl1FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID/Set even frame
func (r *registerOtghsdoepctl1Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl1FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdoepctl1FieldSoddfrmShift = 29
	RegisterOtghsdoepctl1FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdoepctl1Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl1FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdoepctl1FieldEpdisShift = 30
	RegisterOtghsdoepctl1FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdoepctl1Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdoepctl1Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl1FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldEpdisMask)
	}
}

const (
	RegisterOtghsdoepctl1FieldEpenaShift = 31
	RegisterOtghsdoepctl1FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdoepctl1Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdoepctl1Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl1FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldEpenaMask)
	}
}

// registerOtghsdoepint1Type OTG_HS device endpoint-1 interrupt register
type registerOtghsdoepint1Type uint32

const (
	RegisterOtghsdoepint1FieldXfrcShift = 0
	RegisterOtghsdoepint1FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint1Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint1FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint1Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint1FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint1FieldXfrcMask)
	}
}

const (
	RegisterOtghsdoepint1FieldEpdisdShift = 1
	RegisterOtghsdoepint1FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint1Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint1FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint1Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint1FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint1FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdoepint1FieldStupShift = 3
	RegisterOtghsdoepint1FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtghsdoepint1Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint1FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtghsdoepint1Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint1FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint1FieldStupMask)
	}
}

const (
	RegisterOtghsdoepint1FieldOtepdisShift = 4
	RegisterOtghsdoepint1FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint1Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint1FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint1Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint1FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint1FieldOtepdisMask)
	}
}

const (
	RegisterOtghsdoepint1FieldB2bstupShift = 6
	RegisterOtghsdoepint1FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint1Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint1FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint1Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint1FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint1FieldB2bstupMask)
	}
}

const (
	RegisterOtghsdoepint1FieldNyetShift = 14
	RegisterOtghsdoepint1FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtghsdoepint1Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint1FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtghsdoepint1Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint1FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint1FieldNyetMask)
	}
}

// registerOtghsdoeptsiz1Type OTG_HS device endpoint-1 transfer size register
type registerOtghsdoeptsiz1Type uint32

const (
	RegisterOtghsdoeptsiz1FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz1FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz1Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz1FieldXfrsizMask) >> RegisterOtghsdoeptsiz1FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz1Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz1FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz1FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz1FieldPktcntShift = 19
	RegisterOtghsdoeptsiz1FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdoeptsiz1Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz1FieldPktcntMask) >> RegisterOtghsdoeptsiz1FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdoeptsiz1Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz1FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz1FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz1FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz1FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz1Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz1FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz1FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz1Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz1FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz1FieldRxdpidstupcntShift))
}

// registerOtghsdoepctl2Type OTG device endpoint-2 control register
type registerOtghsdoepctl2Type uint32

const (
	RegisterOtghsdoepctl2FieldMpsizShift = 0
	RegisterOtghsdoepctl2FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdoepctl2Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldMpsizMask) >> RegisterOtghsdoepctl2FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdoepctl2Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl2FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl2FieldUsbaepShift = 15
	RegisterOtghsdoepctl2FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdoepctl2Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdoepctl2Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl2FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdoepctl2FieldEonumdpidShift = 16
	RegisterOtghsdoepctl2FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even odd frame/Endpoint data PID
func (r *registerOtghsdoepctl2Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl2FieldNakstsShift = 17
	RegisterOtghsdoepctl2FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdoepctl2Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl2FieldEptypShift = 18
	RegisterOtghsdoepctl2FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdoepctl2Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldEptypMask) >> RegisterOtghsdoepctl2FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdoepctl2Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl2FieldEptypShift))
}

const (
	RegisterOtghsdoepctl2FieldSnpmShift = 20
	RegisterOtghsdoepctl2FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtghsdoepctl2Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtghsdoepctl2Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl2FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldSnpmMask)
	}
}

const (
	RegisterOtghsdoepctl2FieldStallShift = 21
	RegisterOtghsdoepctl2FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdoepctl2Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdoepctl2Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl2FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldStallMask)
	}
}

const (
	RegisterOtghsdoepctl2FieldCnakShift = 26
	RegisterOtghsdoepctl2FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdoepctl2Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl2FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldCnakMask)
	}
}

const (
	RegisterOtghsdoepctl2FieldSnakShift = 27
	RegisterOtghsdoepctl2FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdoepctl2Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl2FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldSnakMask)
	}
}

const (
	RegisterOtghsdoepctl2FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdoepctl2FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID/Set even frame
func (r *registerOtghsdoepctl2Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl2FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdoepctl2FieldSoddfrmShift = 29
	RegisterOtghsdoepctl2FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdoepctl2Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl2FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdoepctl2FieldEpdisShift = 30
	RegisterOtghsdoepctl2FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdoepctl2Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdoepctl2Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl2FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldEpdisMask)
	}
}

const (
	RegisterOtghsdoepctl2FieldEpenaShift = 31
	RegisterOtghsdoepctl2FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdoepctl2Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdoepctl2Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl2FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldEpenaMask)
	}
}

// registerOtghsdoepint2Type OTG_HS device endpoint-2 interrupt register
type registerOtghsdoepint2Type uint32

const (
	RegisterOtghsdoepint2FieldXfrcShift = 0
	RegisterOtghsdoepint2FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint2Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint2FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint2Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint2FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint2FieldXfrcMask)
	}
}

const (
	RegisterOtghsdoepint2FieldEpdisdShift = 1
	RegisterOtghsdoepint2FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint2Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint2FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint2Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint2FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint2FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdoepint2FieldStupShift = 3
	RegisterOtghsdoepint2FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtghsdoepint2Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint2FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtghsdoepint2Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint2FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint2FieldStupMask)
	}
}

const (
	RegisterOtghsdoepint2FieldOtepdisShift = 4
	RegisterOtghsdoepint2FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint2Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint2FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint2Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint2FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint2FieldOtepdisMask)
	}
}

const (
	RegisterOtghsdoepint2FieldB2bstupShift = 6
	RegisterOtghsdoepint2FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint2Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint2FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint2Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint2FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint2FieldB2bstupMask)
	}
}

const (
	RegisterOtghsdoepint2FieldNyetShift = 14
	RegisterOtghsdoepint2FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtghsdoepint2Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint2FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtghsdoepint2Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint2FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint2FieldNyetMask)
	}
}

// registerOtghsdoeptsiz2Type OTG_HS device endpoint-2 transfer size register
type registerOtghsdoeptsiz2Type uint32

const (
	RegisterOtghsdoeptsiz2FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz2FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz2Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz2FieldXfrsizMask) >> RegisterOtghsdoeptsiz2FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz2Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz2FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz2FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz2FieldPktcntShift = 19
	RegisterOtghsdoeptsiz2FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdoeptsiz2Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz2FieldPktcntMask) >> RegisterOtghsdoeptsiz2FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdoeptsiz2Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz2FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz2FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz2FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz2FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz2Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz2FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz2FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz2Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz2FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz2FieldRxdpidstupcntShift))
}

// registerOtghsdoepctl3Type OTG device endpoint-3 control register
type registerOtghsdoepctl3Type uint32

const (
	RegisterOtghsdoepctl3FieldMpsizShift = 0
	RegisterOtghsdoepctl3FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdoepctl3Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldMpsizMask) >> RegisterOtghsdoepctl3FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdoepctl3Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl3FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl3FieldUsbaepShift = 15
	RegisterOtghsdoepctl3FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdoepctl3Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdoepctl3Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl3FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdoepctl3FieldEonumdpidShift = 16
	RegisterOtghsdoepctl3FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even odd frame/Endpoint data PID
func (r *registerOtghsdoepctl3Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl3FieldNakstsShift = 17
	RegisterOtghsdoepctl3FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdoepctl3Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl3FieldEptypShift = 18
	RegisterOtghsdoepctl3FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdoepctl3Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldEptypMask) >> RegisterOtghsdoepctl3FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdoepctl3Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl3FieldEptypShift))
}

const (
	RegisterOtghsdoepctl3FieldSnpmShift = 20
	RegisterOtghsdoepctl3FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtghsdoepctl3Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtghsdoepctl3Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl3FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldSnpmMask)
	}
}

const (
	RegisterOtghsdoepctl3FieldStallShift = 21
	RegisterOtghsdoepctl3FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdoepctl3Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdoepctl3Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl3FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldStallMask)
	}
}

const (
	RegisterOtghsdoepctl3FieldCnakShift = 26
	RegisterOtghsdoepctl3FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdoepctl3Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl3FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldCnakMask)
	}
}

const (
	RegisterOtghsdoepctl3FieldSnakShift = 27
	RegisterOtghsdoepctl3FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdoepctl3Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl3FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldSnakMask)
	}
}

const (
	RegisterOtghsdoepctl3FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdoepctl3FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID/Set even frame
func (r *registerOtghsdoepctl3Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl3FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdoepctl3FieldSoddfrmShift = 29
	RegisterOtghsdoepctl3FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdoepctl3Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl3FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdoepctl3FieldEpdisShift = 30
	RegisterOtghsdoepctl3FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdoepctl3Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdoepctl3Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl3FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldEpdisMask)
	}
}

const (
	RegisterOtghsdoepctl3FieldEpenaShift = 31
	RegisterOtghsdoepctl3FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdoepctl3Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdoepctl3Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl3FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldEpenaMask)
	}
}

// registerOtghsdoepint3Type OTG_HS device endpoint-3 interrupt register
type registerOtghsdoepint3Type uint32

const (
	RegisterOtghsdoepint3FieldXfrcShift = 0
	RegisterOtghsdoepint3FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint3Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint3FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint3Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint3FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint3FieldXfrcMask)
	}
}

const (
	RegisterOtghsdoepint3FieldEpdisdShift = 1
	RegisterOtghsdoepint3FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint3Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint3FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint3Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint3FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint3FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdoepint3FieldStupShift = 3
	RegisterOtghsdoepint3FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtghsdoepint3Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint3FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtghsdoepint3Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint3FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint3FieldStupMask)
	}
}

const (
	RegisterOtghsdoepint3FieldOtepdisShift = 4
	RegisterOtghsdoepint3FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint3Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint3FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint3Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint3FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint3FieldOtepdisMask)
	}
}

const (
	RegisterOtghsdoepint3FieldB2bstupShift = 6
	RegisterOtghsdoepint3FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint3Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint3FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint3Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint3FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint3FieldB2bstupMask)
	}
}

const (
	RegisterOtghsdoepint3FieldNyetShift = 14
	RegisterOtghsdoepint3FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtghsdoepint3Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint3FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtghsdoepint3Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint3FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint3FieldNyetMask)
	}
}

// registerOtghsdoeptsiz3Type OTG_HS device endpoint-3 transfer size register
type registerOtghsdoeptsiz3Type uint32

const (
	RegisterOtghsdoeptsiz3FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz3FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz3Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz3FieldXfrsizMask) >> RegisterOtghsdoeptsiz3FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz3Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz3FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz3FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz3FieldPktcntShift = 19
	RegisterOtghsdoeptsiz3FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdoeptsiz3Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz3FieldPktcntMask) >> RegisterOtghsdoeptsiz3FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdoeptsiz3Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz3FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz3FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz3FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz3FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz3Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz3FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz3FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz3Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz3FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz3FieldRxdpidstupcntShift))
}

// registerOtghsdoepctl4Type OTG device endpoint-4 control register
type registerOtghsdoepctl4Type uint32

const (
	RegisterOtghsdoepctl4FieldMpsizShift = 0
	RegisterOtghsdoepctl4FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdoepctl4Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldMpsizMask) >> RegisterOtghsdoepctl4FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdoepctl4Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl4FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl4FieldUsbaepShift = 15
	RegisterOtghsdoepctl4FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdoepctl4Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdoepctl4Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl4FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdoepctl4FieldEonumdpidShift = 16
	RegisterOtghsdoepctl4FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even odd frame/Endpoint data PID
func (r *registerOtghsdoepctl4Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl4FieldNakstsShift = 17
	RegisterOtghsdoepctl4FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdoepctl4Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl4FieldEptypShift = 18
	RegisterOtghsdoepctl4FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdoepctl4Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldEptypMask) >> RegisterOtghsdoepctl4FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdoepctl4Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl4FieldEptypShift))
}

const (
	RegisterOtghsdoepctl4FieldSnpmShift = 20
	RegisterOtghsdoepctl4FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtghsdoepctl4Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtghsdoepctl4Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl4FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldSnpmMask)
	}
}

const (
	RegisterOtghsdoepctl4FieldStallShift = 21
	RegisterOtghsdoepctl4FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdoepctl4Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdoepctl4Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl4FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldStallMask)
	}
}

const (
	RegisterOtghsdoepctl4FieldCnakShift = 26
	RegisterOtghsdoepctl4FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdoepctl4Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl4FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldCnakMask)
	}
}

const (
	RegisterOtghsdoepctl4FieldSnakShift = 27
	RegisterOtghsdoepctl4FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdoepctl4Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl4FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldSnakMask)
	}
}

const (
	RegisterOtghsdoepctl4FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdoepctl4FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID/Set even frame
func (r *registerOtghsdoepctl4Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl4FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdoepctl4FieldSoddfrmShift = 29
	RegisterOtghsdoepctl4FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdoepctl4Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl4FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdoepctl4FieldEpdisShift = 30
	RegisterOtghsdoepctl4FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdoepctl4Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdoepctl4Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl4FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldEpdisMask)
	}
}

const (
	RegisterOtghsdoepctl4FieldEpenaShift = 31
	RegisterOtghsdoepctl4FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdoepctl4Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdoepctl4Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl4FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldEpenaMask)
	}
}

// registerOtghsdoepint4Type OTG_HS device endpoint-4 interrupt register
type registerOtghsdoepint4Type uint32

const (
	RegisterOtghsdoepint4FieldXfrcShift = 0
	RegisterOtghsdoepint4FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint4Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint4FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint4Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint4FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint4FieldXfrcMask)
	}
}

const (
	RegisterOtghsdoepint4FieldEpdisdShift = 1
	RegisterOtghsdoepint4FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint4Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint4FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint4Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint4FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint4FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdoepint4FieldStupShift = 3
	RegisterOtghsdoepint4FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtghsdoepint4Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint4FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtghsdoepint4Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint4FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint4FieldStupMask)
	}
}

const (
	RegisterOtghsdoepint4FieldOtepdisShift = 4
	RegisterOtghsdoepint4FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint4Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint4FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint4Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint4FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint4FieldOtepdisMask)
	}
}

const (
	RegisterOtghsdoepint4FieldB2bstupShift = 6
	RegisterOtghsdoepint4FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint4Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint4FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint4Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint4FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint4FieldB2bstupMask)
	}
}

const (
	RegisterOtghsdoepint4FieldNyetShift = 14
	RegisterOtghsdoepint4FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtghsdoepint4Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint4FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtghsdoepint4Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint4FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint4FieldNyetMask)
	}
}

// registerOtghsdoeptsiz4Type OTG_HS device endpoint-4 transfer size register
type registerOtghsdoeptsiz4Type uint32

const (
	RegisterOtghsdoeptsiz4FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz4FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz4Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz4FieldXfrsizMask) >> RegisterOtghsdoeptsiz4FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz4Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz4FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz4FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz4FieldPktcntShift = 19
	RegisterOtghsdoeptsiz4FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdoeptsiz4Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz4FieldPktcntMask) >> RegisterOtghsdoeptsiz4FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdoeptsiz4Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz4FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz4FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz4FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz4FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz4Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz4FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz4FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz4Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz4FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz4FieldRxdpidstupcntShift))
}

// registerOtghsdoepctl5Type OTG device endpoint-5 control register
type registerOtghsdoepctl5Type uint32

const (
	RegisterOtghsdoepctl5FieldMpsizShift = 0
	RegisterOtghsdoepctl5FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdoepctl5Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldMpsizMask) >> RegisterOtghsdoepctl5FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdoepctl5Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl5FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl5FieldUsbaepShift = 15
	RegisterOtghsdoepctl5FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdoepctl5Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdoepctl5Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl5FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdoepctl5FieldEonumdpidShift = 16
	RegisterOtghsdoepctl5FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even odd frame/Endpoint data PID
func (r *registerOtghsdoepctl5Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl5FieldNakstsShift = 17
	RegisterOtghsdoepctl5FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdoepctl5Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl5FieldEptypShift = 18
	RegisterOtghsdoepctl5FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdoepctl5Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldEptypMask) >> RegisterOtghsdoepctl5FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdoepctl5Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl5FieldEptypShift))
}

const (
	RegisterOtghsdoepctl5FieldSnpmShift = 20
	RegisterOtghsdoepctl5FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtghsdoepctl5Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtghsdoepctl5Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl5FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldSnpmMask)
	}
}

const (
	RegisterOtghsdoepctl5FieldStallShift = 21
	RegisterOtghsdoepctl5FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdoepctl5Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdoepctl5Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl5FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldStallMask)
	}
}

const (
	RegisterOtghsdoepctl5FieldCnakShift = 26
	RegisterOtghsdoepctl5FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdoepctl5Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl5FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldCnakMask)
	}
}

const (
	RegisterOtghsdoepctl5FieldSnakShift = 27
	RegisterOtghsdoepctl5FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdoepctl5Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl5FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldSnakMask)
	}
}

const (
	RegisterOtghsdoepctl5FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdoepctl5FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID/Set even frame
func (r *registerOtghsdoepctl5Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl5FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdoepctl5FieldSoddfrmShift = 29
	RegisterOtghsdoepctl5FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdoepctl5Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl5FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdoepctl5FieldEpdisShift = 30
	RegisterOtghsdoepctl5FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdoepctl5Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdoepctl5Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl5FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldEpdisMask)
	}
}

const (
	RegisterOtghsdoepctl5FieldEpenaShift = 31
	RegisterOtghsdoepctl5FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdoepctl5Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdoepctl5Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl5FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldEpenaMask)
	}
}

// registerOtghsdoepint5Type OTG_HS device endpoint-5 interrupt register
type registerOtghsdoepint5Type uint32

const (
	RegisterOtghsdoepint5FieldXfrcShift = 0
	RegisterOtghsdoepint5FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint5Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint5FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint5Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint5FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint5FieldXfrcMask)
	}
}

const (
	RegisterOtghsdoepint5FieldEpdisdShift = 1
	RegisterOtghsdoepint5FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint5Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint5FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint5Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint5FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint5FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdoepint5FieldStupShift = 3
	RegisterOtghsdoepint5FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtghsdoepint5Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint5FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtghsdoepint5Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint5FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint5FieldStupMask)
	}
}

const (
	RegisterOtghsdoepint5FieldOtepdisShift = 4
	RegisterOtghsdoepint5FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint5Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint5FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint5Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint5FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint5FieldOtepdisMask)
	}
}

const (
	RegisterOtghsdoepint5FieldB2bstupShift = 6
	RegisterOtghsdoepint5FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint5Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint5FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint5Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint5FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint5FieldB2bstupMask)
	}
}

const (
	RegisterOtghsdoepint5FieldNyetShift = 14
	RegisterOtghsdoepint5FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtghsdoepint5Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint5FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtghsdoepint5Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint5FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint5FieldNyetMask)
	}
}

// registerOtghsdoeptsiz5Type OTG_HS device endpoint-5 transfer size register
type registerOtghsdoeptsiz5Type uint32

const (
	RegisterOtghsdoeptsiz5FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz5FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz5Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz5FieldXfrsizMask) >> RegisterOtghsdoeptsiz5FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz5Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz5FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz5FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz5FieldPktcntShift = 19
	RegisterOtghsdoeptsiz5FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdoeptsiz5Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz5FieldPktcntMask) >> RegisterOtghsdoeptsiz5FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdoeptsiz5Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz5FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz5FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz5FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz5FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz5Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz5FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz5FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz5Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz5FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz5FieldRxdpidstupcntShift))
}

// registerOtghsdoepctl6Type OTG device endpoint-6 control register
type registerOtghsdoepctl6Type uint32

const (
	RegisterOtghsdoepctl6FieldMpsizShift = 0
	RegisterOtghsdoepctl6FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdoepctl6Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldMpsizMask) >> RegisterOtghsdoepctl6FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdoepctl6Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl6FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl6FieldUsbaepShift = 15
	RegisterOtghsdoepctl6FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdoepctl6Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdoepctl6Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl6FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdoepctl6FieldEonumdpidShift = 16
	RegisterOtghsdoepctl6FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even odd frame/Endpoint data PID
func (r *registerOtghsdoepctl6Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl6FieldNakstsShift = 17
	RegisterOtghsdoepctl6FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdoepctl6Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl6FieldEptypShift = 18
	RegisterOtghsdoepctl6FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdoepctl6Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldEptypMask) >> RegisterOtghsdoepctl6FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdoepctl6Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl6FieldEptypShift))
}

const (
	RegisterOtghsdoepctl6FieldSnpmShift = 20
	RegisterOtghsdoepctl6FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtghsdoepctl6Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtghsdoepctl6Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl6FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldSnpmMask)
	}
}

const (
	RegisterOtghsdoepctl6FieldStallShift = 21
	RegisterOtghsdoepctl6FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdoepctl6Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdoepctl6Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl6FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldStallMask)
	}
}

const (
	RegisterOtghsdoepctl6FieldCnakShift = 26
	RegisterOtghsdoepctl6FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdoepctl6Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl6FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldCnakMask)
	}
}

const (
	RegisterOtghsdoepctl6FieldSnakShift = 27
	RegisterOtghsdoepctl6FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdoepctl6Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl6FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldSnakMask)
	}
}

const (
	RegisterOtghsdoepctl6FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdoepctl6FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID/Set even frame
func (r *registerOtghsdoepctl6Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl6FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdoepctl6FieldSoddfrmShift = 29
	RegisterOtghsdoepctl6FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdoepctl6Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl6FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdoepctl6FieldEpdisShift = 30
	RegisterOtghsdoepctl6FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdoepctl6Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdoepctl6Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl6FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldEpdisMask)
	}
}

const (
	RegisterOtghsdoepctl6FieldEpenaShift = 31
	RegisterOtghsdoepctl6FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdoepctl6Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdoepctl6Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl6FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldEpenaMask)
	}
}

// registerOtghsdoepint6Type OTG_HS device endpoint-6 interrupt register
type registerOtghsdoepint6Type uint32

const (
	RegisterOtghsdoepint6FieldXfrcShift = 0
	RegisterOtghsdoepint6FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint6Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint6FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint6Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint6FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint6FieldXfrcMask)
	}
}

const (
	RegisterOtghsdoepint6FieldEpdisdShift = 1
	RegisterOtghsdoepint6FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint6Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint6FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint6Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint6FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint6FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdoepint6FieldStupShift = 3
	RegisterOtghsdoepint6FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtghsdoepint6Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint6FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtghsdoepint6Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint6FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint6FieldStupMask)
	}
}

const (
	RegisterOtghsdoepint6FieldOtepdisShift = 4
	RegisterOtghsdoepint6FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint6Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint6FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint6Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint6FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint6FieldOtepdisMask)
	}
}

const (
	RegisterOtghsdoepint6FieldB2bstupShift = 6
	RegisterOtghsdoepint6FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint6Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint6FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint6Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint6FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint6FieldB2bstupMask)
	}
}

const (
	RegisterOtghsdoepint6FieldNyetShift = 14
	RegisterOtghsdoepint6FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtghsdoepint6Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint6FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtghsdoepint6Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint6FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint6FieldNyetMask)
	}
}

// registerOtghsdoeptsiz6Type OTG_HS device endpoint-6 transfer size register
type registerOtghsdoeptsiz6Type uint32

const (
	RegisterOtghsdoeptsiz6FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz6FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz6Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz6FieldXfrsizMask) >> RegisterOtghsdoeptsiz6FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz6Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz6FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz6FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz6FieldPktcntShift = 19
	RegisterOtghsdoeptsiz6FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdoeptsiz6Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz6FieldPktcntMask) >> RegisterOtghsdoeptsiz6FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdoeptsiz6Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz6FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz6FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz6FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz6FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz6Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz6FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz6FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz6Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz6FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz6FieldRxdpidstupcntShift))
}

// registerOtghsdoepctl7Type OTG device endpoint-7 control register
type registerOtghsdoepctl7Type uint32

const (
	RegisterOtghsdoepctl7FieldMpsizShift = 0
	RegisterOtghsdoepctl7FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghsdoepctl7Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldMpsizMask) >> RegisterOtghsdoepctl7FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghsdoepctl7Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl7FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl7FieldUsbaepShift = 15
	RegisterOtghsdoepctl7FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *registerOtghsdoepctl7Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *registerOtghsdoepctl7Type) SetUsbaep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl7FieldUsbaepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldUsbaepMask)
	}
}

const (
	RegisterOtghsdoepctl7FieldEonumdpidShift = 16
	RegisterOtghsdoepctl7FieldEonumdpidMask  = 0x10000
)

// GetEonumdpid Even odd frame/Endpoint data PID
func (r *registerOtghsdoepctl7Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl7FieldNakstsShift = 17
	RegisterOtghsdoepctl7FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *registerOtghsdoepctl7Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl7FieldEptypShift = 18
	RegisterOtghsdoepctl7FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghsdoepctl7Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldEptypMask) >> RegisterOtghsdoepctl7FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghsdoepctl7Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl7FieldEptypShift))
}

const (
	RegisterOtghsdoepctl7FieldSnpmShift = 20
	RegisterOtghsdoepctl7FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *registerOtghsdoepctl7Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *registerOtghsdoepctl7Type) SetSnpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl7FieldSnpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldSnpmMask)
	}
}

const (
	RegisterOtghsdoepctl7FieldStallShift = 21
	RegisterOtghsdoepctl7FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *registerOtghsdoepctl7Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *registerOtghsdoepctl7Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl7FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldStallMask)
	}
}

const (
	RegisterOtghsdoepctl7FieldCnakShift = 26
	RegisterOtghsdoepctl7FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *registerOtghsdoepctl7Type) SetCnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl7FieldCnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldCnakMask)
	}
}

const (
	RegisterOtghsdoepctl7FieldSnakShift = 27
	RegisterOtghsdoepctl7FieldSnakMask  = 0x8000000
)

// SetSnak Set NAK
func (r *registerOtghsdoepctl7Type) SetSnak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl7FieldSnakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldSnakMask)
	}
}

const (
	RegisterOtghsdoepctl7FieldSd0pidsevnfrmShift = 28
	RegisterOtghsdoepctl7FieldSd0pidsevnfrmMask  = 0x10000000
)

// SetSd0pidsevnfrm Set DATA0 PID/Set even frame
func (r *registerOtghsdoepctl7Type) SetSd0pidsevnfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl7FieldSd0pidsevnfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldSd0pidsevnfrmMask)
	}
}

const (
	RegisterOtghsdoepctl7FieldSoddfrmShift = 29
	RegisterOtghsdoepctl7FieldSoddfrmMask  = 0x20000000
)

// SetSoddfrm Set odd frame
func (r *registerOtghsdoepctl7Type) SetSoddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl7FieldSoddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldSoddfrmMask)
	}
}

const (
	RegisterOtghsdoepctl7FieldEpdisShift = 30
	RegisterOtghsdoepctl7FieldEpdisMask  = 0x40000000
)

// GetEpdis Endpoint disable
func (r *registerOtghsdoepctl7Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *registerOtghsdoepctl7Type) SetEpdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl7FieldEpdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldEpdisMask)
	}
}

const (
	RegisterOtghsdoepctl7FieldEpenaShift = 31
	RegisterOtghsdoepctl7FieldEpenaMask  = 0x80000000
)

// GetEpena Endpoint enable
func (r *registerOtghsdoepctl7Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *registerOtghsdoepctl7Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl7FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldEpenaMask)
	}
}

// registerOtghsdoepint7Type OTG_HS device endpoint-7 interrupt register
type registerOtghsdoepint7Type uint32

const (
	RegisterOtghsdoepint7FieldXfrcShift = 0
	RegisterOtghsdoepint7FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint7Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint7FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *registerOtghsdoepint7Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint7FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint7FieldXfrcMask)
	}
}

const (
	RegisterOtghsdoepint7FieldEpdisdShift = 1
	RegisterOtghsdoepint7FieldEpdisdMask  = 0x2
)

// GetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint7Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint7FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *registerOtghsdoepint7Type) SetEpdisd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint7FieldEpdisdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint7FieldEpdisdMask)
	}
}

const (
	RegisterOtghsdoepint7FieldStupShift = 3
	RegisterOtghsdoepint7FieldStupMask  = 0x8
)

// GetStup SETUP phase done
func (r *registerOtghsdoepint7Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint7FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *registerOtghsdoepint7Type) SetStup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint7FieldStupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint7FieldStupMask)
	}
}

const (
	RegisterOtghsdoepint7FieldOtepdisShift = 4
	RegisterOtghsdoepint7FieldOtepdisMask  = 0x10
)

// GetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint7Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint7FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *registerOtghsdoepint7Type) SetOtepdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint7FieldOtepdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint7FieldOtepdisMask)
	}
}

const (
	RegisterOtghsdoepint7FieldB2bstupShift = 6
	RegisterOtghsdoepint7FieldB2bstupMask  = 0x40
)

// GetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint7Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint7FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *registerOtghsdoepint7Type) SetB2bstup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint7FieldB2bstupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint7FieldB2bstupMask)
	}
}

const (
	RegisterOtghsdoepint7FieldNyetShift = 14
	RegisterOtghsdoepint7FieldNyetMask  = 0x4000
)

// GetNyet NYET interrupt
func (r *registerOtghsdoepint7Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint7FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *registerOtghsdoepint7Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint7FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint7FieldNyetMask)
	}
}

// registerOtghsdoeptsiz7Type OTG_HS device endpoint-7 transfer size register
type registerOtghsdoeptsiz7Type uint32

const (
	RegisterOtghsdoeptsiz7FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz7FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz7Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz7FieldXfrsizMask) >> RegisterOtghsdoeptsiz7FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghsdoeptsiz7Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz7FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz7FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz7FieldPktcntShift = 19
	RegisterOtghsdoeptsiz7FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghsdoeptsiz7Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz7FieldPktcntMask) >> RegisterOtghsdoeptsiz7FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghsdoeptsiz7Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz7FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz7FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz7FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz7FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz7Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz7FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz7FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *registerOtghsdoeptsiz7Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz7FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz7FieldRxdpidstupcntShift))
}
