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
	Otghsdcfg        RegisterOtghsdcfgType
	Otghsdctl        RegisterOtghsdctlType
	Otghsdsts        RegisterOtghsdstsType
	_                [4]byte
	Otghsdiepmsk     RegisterOtghsdiepmskType
	Otghsdoepmsk     RegisterOtghsdoepmskType
	Otghsdaint       RegisterOtghsdaintType
	Otghsdaintmsk    RegisterOtghsdaintmskType
	_                [8]byte
	Otghsdvbusdis    RegisterOtghsdvbusdisType
	Otghsdvbuspulse  RegisterOtghsdvbuspulseType
	Otghsdthrctl     RegisterOtghsdthrctlType
	Otghsdiepempmsk  RegisterOtghsdiepempmskType
	Otghsdeachint    RegisterOtghsdeachintType
	Otghsdeachintmsk RegisterOtghsdeachintmskType
	_                [192]byte
	Otghsdiepctl0    RegisterOtghsdiepctl0Type
	_                [4]byte
	Otghsdiepint0    RegisterOtghsdiepint0Type
	_                [4]byte
	Otghsdieptsiz0   RegisterOtghsdieptsiz0Type
	Otghsdiepdma1    RegisterOtghsdiepdma1Type
	Otghsdtxfsts0    RegisterOtghsdtxfsts0Type
	_                [4]byte
	Otghsdiepctl1    RegisterOtghsdiepctl1Type
	_                [4]byte
	Otghsdiepint1    RegisterOtghsdiepint1Type
	_                [4]byte
	Otghsdieptsiz1   RegisterOtghsdieptsiz1Type
	Otghsdiepdma2    RegisterOtghsdiepdma2Type
	Otghsdtxfsts1    RegisterOtghsdtxfsts1Type
	_                [4]byte
	Otghsdiepctl2    RegisterOtghsdiepctl2Type
	_                [4]byte
	Otghsdiepint2    RegisterOtghsdiepint2Type
	_                [4]byte
	Otghsdieptsiz2   RegisterOtghsdieptsiz2Type
	Otghsdiepdma3    RegisterOtghsdiepdma3Type
	Otghsdtxfsts2    RegisterOtghsdtxfsts2Type
	_                [4]byte
	Otghsdiepctl3    RegisterOtghsdiepctl3Type
	_                [4]byte
	Otghsdiepint3    RegisterOtghsdiepint3Type
	_                [4]byte
	Otghsdieptsiz3   RegisterOtghsdieptsiz3Type
	Otghsdiepdma4    RegisterOtghsdiepdma4Type
	Otghsdtxfsts3    RegisterOtghsdtxfsts3Type
	_                [4]byte
	Otghsdiepctl4    RegisterOtghsdiepctl4Type
	_                [4]byte
	Otghsdiepint4    RegisterOtghsdiepint4Type
	_                [4]byte
	Otghsdieptsiz4   RegisterOtghsdieptsiz4Type
	Otghsdiepdma5    RegisterOtghsdiepdma5Type
	Otghsdtxfsts4    RegisterOtghsdtxfsts4Type
	_                [4]byte
	Otghsdiepctl5    RegisterOtghsdiepctl5Type
	Otghsdieptsiz6   RegisterOtghsdieptsiz6Type
	Otghsdtxfsts6    RegisterOtghsdtxfsts6Type
	Otghsdieptsiz7   RegisterOtghsdieptsiz7Type
	Otghsdiepint5    RegisterOtghsdiepint5Type
	Otghsdtxfsts7    RegisterOtghsdtxfsts7Type
	Otghsdieptsiz5   RegisterOtghsdieptsiz5Type
	_                [4]byte
	Otghsdtxfsts5    RegisterOtghsdtxfsts5Type
	_                [4]byte
	Otghsdiepctl6    RegisterOtghsdiepctl6Type
	_                [4]byte
	Otghsdiepint6    RegisterOtghsdiepint6Type
	_                [20]byte
	Otghsdiepctl7    RegisterOtghsdiepctl7Type
	_                [4]byte
	Otghsdiepint7    RegisterOtghsdiepint7Type
	_                [276]byte
	Otghsdoepctl0    RegisterOtghsdoepctl0Type
	_                [4]byte
	Otghsdoepint0    RegisterOtghsdoepint0Type
	_                [4]byte
	Otghsdoeptsiz0   RegisterOtghsdoeptsiz0Type
	_                [12]byte
	Otghsdoepctl1    RegisterOtghsdoepctl1Type
	_                [4]byte
	Otghsdoepint1    RegisterOtghsdoepint1Type
	_                [4]byte
	Otghsdoeptsiz1   RegisterOtghsdoeptsiz1Type
	_                [12]byte
	Otghsdoepctl2    RegisterOtghsdoepctl2Type
	_                [4]byte
	Otghsdoepint2    RegisterOtghsdoepint2Type
	_                [4]byte
	Otghsdoeptsiz2   RegisterOtghsdoeptsiz2Type
	_                [12]byte
	Otghsdoepctl3    RegisterOtghsdoepctl3Type
	_                [4]byte
	Otghsdoepint3    RegisterOtghsdoepint3Type
	_                [4]byte
	Otghsdoeptsiz3   RegisterOtghsdoeptsiz3Type
	_                [12]byte
	Otghsdoepctl4    RegisterOtghsdoepctl4Type
	_                [4]byte
	Otghsdoepint4    RegisterOtghsdoepint4Type
	_                [4]byte
	Otghsdoeptsiz4   RegisterOtghsdoeptsiz4Type
	_                [12]byte
	Otghsdoepctl5    RegisterOtghsdoepctl5Type
	_                [4]byte
	Otghsdoepint5    RegisterOtghsdoepint5Type
	_                [4]byte
	Otghsdoeptsiz5   RegisterOtghsdoeptsiz5Type
	_                [12]byte
	Otghsdoepctl6    RegisterOtghsdoepctl6Type
	_                [4]byte
	Otghsdoepint6    RegisterOtghsdoepint6Type
	_                [4]byte
	Otghsdoeptsiz6   RegisterOtghsdoeptsiz6Type
	_                [12]byte
	Otghsdoepctl7    RegisterOtghsdoepctl7Type
	_                [4]byte
	Otghsdoepint7    RegisterOtghsdoepint7Type
	_                [4]byte
	Otghsdoeptsiz7   RegisterOtghsdoeptsiz7Type
}

// RegisterOtghsdcfgType OTG_HS device configuration register
type RegisterOtghsdcfgType uint32

func (r *RegisterOtghsdcfgType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdcfgType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdcfgType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdcfgType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdcfgType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdcfgFieldDspdShift = 0
	RegisterOtghsdcfgFieldDspdMask  = 0x3
)

// GetDspd Device speed
func (r *RegisterOtghsdcfgType) GetDspd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdcfgFieldDspdMask) >> RegisterOtghsdcfgFieldDspdShift)
}

// SetDspd Device speed
func (r *RegisterOtghsdcfgType) SetDspd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdcfgFieldDspdMask)|(uint32(value)<<RegisterOtghsdcfgFieldDspdShift))
}

const (
	RegisterOtghsdcfgFieldNzlsohskShift = 2
	RegisterOtghsdcfgFieldNzlsohskMask  = 0x4
)

// GetNzlsohsk Nonzero-length status OUT handshake
func (r *RegisterOtghsdcfgType) GetNzlsohsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdcfgFieldNzlsohskMask) != 0
}

// SetNzlsohsk Nonzero-length status OUT handshake
func (r *RegisterOtghsdcfgType) SetNzlsohsk(value bool) {
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
func (r *RegisterOtghsdcfgType) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdcfgFieldDadMask) >> RegisterOtghsdcfgFieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghsdcfgType) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdcfgFieldDadMask)|(uint32(value)<<RegisterOtghsdcfgFieldDadShift))
}

const (
	RegisterOtghsdcfgFieldPfivlShift = 11
	RegisterOtghsdcfgFieldPfivlMask  = 0x1800
)

// GetPfivl Periodic (micro)frame interval
func (r *RegisterOtghsdcfgType) GetPfivl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdcfgFieldPfivlMask) >> RegisterOtghsdcfgFieldPfivlShift)
}

// SetPfivl Periodic (micro)frame interval
func (r *RegisterOtghsdcfgType) SetPfivl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdcfgFieldPfivlMask)|(uint32(value)<<RegisterOtghsdcfgFieldPfivlShift))
}

const (
	RegisterOtghsdcfgFieldPerschivlShift = 24
	RegisterOtghsdcfgFieldPerschivlMask  = 0x3000000
)

// GetPerschivl Periodic scheduling interval
func (r *RegisterOtghsdcfgType) GetPerschivl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdcfgFieldPerschivlMask) >> RegisterOtghsdcfgFieldPerschivlShift)
}

// SetPerschivl Periodic scheduling interval
func (r *RegisterOtghsdcfgType) SetPerschivl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdcfgFieldPerschivlMask)|(uint32(value)<<RegisterOtghsdcfgFieldPerschivlShift))
}

// RegisterOtghsdctlType OTG_HS device control register
type RegisterOtghsdctlType uint32

func (r *RegisterOtghsdctlType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdctlType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdctlType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdctlType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdctlType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdctlFieldRwusigShift = 0
	RegisterOtghsdctlFieldRwusigMask  = 0x1
)

// GetRwusig Remote wakeup signaling
func (r *RegisterOtghsdctlType) GetRwusig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdctlFieldRwusigMask) != 0
}

// SetRwusig Remote wakeup signaling
func (r *RegisterOtghsdctlType) SetRwusig(value bool) {
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
func (r *RegisterOtghsdctlType) GetSdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdctlFieldSdisMask) != 0
}

// SetSdis Soft disconnect
func (r *RegisterOtghsdctlType) SetSdis(value bool) {
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
func (r *RegisterOtghsdctlType) GetGinsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdctlFieldGinstsMask) != 0
}

const (
	RegisterOtghsdctlFieldGonstsShift = 3
	RegisterOtghsdctlFieldGonstsMask  = 0x8
)

// GetGonsts Global OUT NAK status
func (r *RegisterOtghsdctlType) GetGonsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdctlFieldGonstsMask) != 0
}

const (
	RegisterOtghsdctlFieldTctlShift = 4
	RegisterOtghsdctlFieldTctlMask  = 0x70
)

// GetTctl Test control
func (r *RegisterOtghsdctlType) GetTctl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdctlFieldTctlMask) >> RegisterOtghsdctlFieldTctlShift)
}

// SetTctl Test control
func (r *RegisterOtghsdctlType) SetTctl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdctlFieldTctlMask)|(uint32(value)<<RegisterOtghsdctlFieldTctlShift))
}

const (
	RegisterOtghsdctlFieldSginakShift = 7
	RegisterOtghsdctlFieldSginakMask  = 0x80
)

// SetSginak Set global IN NAK
func (r *RegisterOtghsdctlType) SetSginak(value bool) {
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
func (r *RegisterOtghsdctlType) SetCginak(value bool) {
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
func (r *RegisterOtghsdctlType) SetSgonak(value bool) {
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
func (r *RegisterOtghsdctlType) SetCgonak(value bool) {
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
func (r *RegisterOtghsdctlType) GetPoprgdne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdctlFieldPoprgdneMask) != 0
}

// SetPoprgdne Power-on programming done
func (r *RegisterOtghsdctlType) SetPoprgdne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdctlFieldPoprgdneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdctlFieldPoprgdneMask)
	}
}

// RegisterOtghsdstsType OTG_HS device status register
type RegisterOtghsdstsType uint32

func (r *RegisterOtghsdstsType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdstsType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdstsType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdstsType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdstsType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdstsFieldSuspstsShift = 0
	RegisterOtghsdstsFieldSuspstsMask  = 0x1
)

// GetSuspsts Suspend status
func (r *RegisterOtghsdstsType) GetSuspsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdstsFieldSuspstsMask) != 0
}

// SetSuspsts Suspend status
func (r *RegisterOtghsdstsType) SetSuspsts(value bool) {
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
func (r *RegisterOtghsdstsType) GetEnumspd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdstsFieldEnumspdMask) >> RegisterOtghsdstsFieldEnumspdShift)
}

// SetEnumspd Enumerated speed
func (r *RegisterOtghsdstsType) SetEnumspd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdstsFieldEnumspdMask)|(uint32(value)<<RegisterOtghsdstsFieldEnumspdShift))
}

const (
	RegisterOtghsdstsFieldEerrShift = 3
	RegisterOtghsdstsFieldEerrMask  = 0x8
)

// GetEerr Erratic error
func (r *RegisterOtghsdstsType) GetEerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdstsFieldEerrMask) != 0
}

// SetEerr Erratic error
func (r *RegisterOtghsdstsType) SetEerr(value bool) {
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
func (r *RegisterOtghsdstsType) GetFnsof() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdstsFieldFnsofMask) >> RegisterOtghsdstsFieldFnsofShift)
}

// SetFnsof Frame number of the received SOF
func (r *RegisterOtghsdstsType) SetFnsof(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdstsFieldFnsofMask)|(uint32(value)<<RegisterOtghsdstsFieldFnsofShift))
}

// RegisterOtghsdiepmskType OTG_HS device IN endpoint common interrupt mask register
type RegisterOtghsdiepmskType uint32

func (r *RegisterOtghsdiepmskType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepmskType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepmskType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepmskType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepmskType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepmskFieldXfrcmShift = 0
	RegisterOtghsdiepmskFieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed interrupt mask
func (r *RegisterOtghsdiepmskType) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed interrupt mask
func (r *RegisterOtghsdiepmskType) SetXfrcm(value bool) {
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
func (r *RegisterOtghsdiepmskType) GetEpdm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldEpdmMask) != 0
}

// SetEpdm Endpoint disabled interrupt mask
func (r *RegisterOtghsdiepmskType) SetEpdm(value bool) {
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
func (r *RegisterOtghsdiepmskType) GetTom() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldTomMask) != 0
}

// SetTom Timeout condition mask (nonisochronous endpoints)
func (r *RegisterOtghsdiepmskType) SetTom(value bool) {
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
func (r *RegisterOtghsdiepmskType) GetIttxfemsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldIttxfemskMask) != 0
}

// SetIttxfemsk IN token received when TxFIFO empty mask
func (r *RegisterOtghsdiepmskType) SetIttxfemsk(value bool) {
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
func (r *RegisterOtghsdiepmskType) GetInepnmm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldInepnmmMask) != 0
}

// SetInepnmm IN token received with EP mismatch mask
func (r *RegisterOtghsdiepmskType) SetInepnmm(value bool) {
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
func (r *RegisterOtghsdiepmskType) GetInepnem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldInepnemMask) != 0
}

// SetInepnem IN endpoint NAK effective mask
func (r *RegisterOtghsdiepmskType) SetInepnem(value bool) {
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
func (r *RegisterOtghsdiepmskType) GetTxfurm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldTxfurmMask) != 0
}

// SetTxfurm FIFO underrun mask
func (r *RegisterOtghsdiepmskType) SetTxfurm(value bool) {
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
func (r *RegisterOtghsdiepmskType) GetBim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepmskFieldBimMask) != 0
}

// SetBim BNA interrupt mask
func (r *RegisterOtghsdiepmskType) SetBim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepmskFieldBimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepmskFieldBimMask)
	}
}

// RegisterOtghsdoepmskType OTG_HS device OUT endpoint common interrupt mask register
type RegisterOtghsdoepmskType uint32

func (r *RegisterOtghsdoepmskType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepmskType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepmskType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepmskType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepmskType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepmskFieldXfrcmShift = 0
	RegisterOtghsdoepmskFieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed interrupt mask
func (r *RegisterOtghsdoepmskType) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed interrupt mask
func (r *RegisterOtghsdoepmskType) SetXfrcm(value bool) {
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
func (r *RegisterOtghsdoepmskType) GetEpdm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldEpdmMask) != 0
}

// SetEpdm Endpoint disabled interrupt mask
func (r *RegisterOtghsdoepmskType) SetEpdm(value bool) {
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
func (r *RegisterOtghsdoepmskType) GetStupm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldStupmMask) != 0
}

// SetStupm SETUP phase done mask
func (r *RegisterOtghsdoepmskType) SetStupm(value bool) {
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
func (r *RegisterOtghsdoepmskType) GetOtepdm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldOtepdmMask) != 0
}

// SetOtepdm OUT token received when endpoint disabled mask
func (r *RegisterOtghsdoepmskType) SetOtepdm(value bool) {
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
func (r *RegisterOtghsdoepmskType) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received mask
func (r *RegisterOtghsdoepmskType) SetB2bstup(value bool) {
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
func (r *RegisterOtghsdoepmskType) GetOpem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldOpemMask) != 0
}

// SetOpem OUT packet error mask
func (r *RegisterOtghsdoepmskType) SetOpem(value bool) {
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
func (r *RegisterOtghsdoepmskType) GetBoim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepmskFieldBoimMask) != 0
}

// SetBoim BNA interrupt mask
func (r *RegisterOtghsdoepmskType) SetBoim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepmskFieldBoimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepmskFieldBoimMask)
	}
}

// RegisterOtghsdaintType OTG_HS device all endpoints interrupt register
type RegisterOtghsdaintType uint32

func (r *RegisterOtghsdaintType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdaintType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdaintType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdaintType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdaintType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdaintFieldIepintShift = 0
	RegisterOtghsdaintFieldIepintMask  = 0xffff
)

// GetIepint IN endpoint interrupt bits
func (r *RegisterOtghsdaintType) GetIepint() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdaintFieldIepintMask) >> RegisterOtghsdaintFieldIepintShift)
}

// SetIepint IN endpoint interrupt bits
func (r *RegisterOtghsdaintType) SetIepint(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdaintFieldIepintMask)|(uint32(value)<<RegisterOtghsdaintFieldIepintShift))
}

const (
	RegisterOtghsdaintFieldOepintShift = 16
	RegisterOtghsdaintFieldOepintMask  = 0xffff0000
)

// GetOepint OUT endpoint interrupt bits
func (r *RegisterOtghsdaintType) GetOepint() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdaintFieldOepintMask) >> RegisterOtghsdaintFieldOepintShift)
}

// SetOepint OUT endpoint interrupt bits
func (r *RegisterOtghsdaintType) SetOepint(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdaintFieldOepintMask)|(uint32(value)<<RegisterOtghsdaintFieldOepintShift))
}

// RegisterOtghsdaintmskType OTG_HS all endpoints interrupt mask register
type RegisterOtghsdaintmskType uint32

func (r *RegisterOtghsdaintmskType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdaintmskType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdaintmskType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdaintmskType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdaintmskType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdaintmskFieldIepmShift = 0
	RegisterOtghsdaintmskFieldIepmMask  = 0xffff
)

// GetIepm IN EP interrupt mask bits
func (r *RegisterOtghsdaintmskType) GetIepm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdaintmskFieldIepmMask) >> RegisterOtghsdaintmskFieldIepmShift)
}

// SetIepm IN EP interrupt mask bits
func (r *RegisterOtghsdaintmskType) SetIepm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdaintmskFieldIepmMask)|(uint32(value)<<RegisterOtghsdaintmskFieldIepmShift))
}

const (
	RegisterOtghsdaintmskFieldOepmShift = 16
	RegisterOtghsdaintmskFieldOepmMask  = 0xffff0000
)

// GetOepm OUT EP interrupt mask bits
func (r *RegisterOtghsdaintmskType) GetOepm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdaintmskFieldOepmMask) >> RegisterOtghsdaintmskFieldOepmShift)
}

// SetOepm OUT EP interrupt mask bits
func (r *RegisterOtghsdaintmskType) SetOepm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdaintmskFieldOepmMask)|(uint32(value)<<RegisterOtghsdaintmskFieldOepmShift))
}

// RegisterOtghsdvbusdisType OTG_HS device VBUS discharge time register
type RegisterOtghsdvbusdisType uint32

func (r *RegisterOtghsdvbusdisType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdvbusdisType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdvbusdisType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdvbusdisType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdvbusdisType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdvbusdisFieldVbusdtShift = 0
	RegisterOtghsdvbusdisFieldVbusdtMask  = 0xffff
)

// GetVbusdt Device VBUS discharge time
func (r *RegisterOtghsdvbusdisType) GetVbusdt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdvbusdisFieldVbusdtMask) >> RegisterOtghsdvbusdisFieldVbusdtShift)
}

// SetVbusdt Device VBUS discharge time
func (r *RegisterOtghsdvbusdisType) SetVbusdt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdvbusdisFieldVbusdtMask)|(uint32(value)<<RegisterOtghsdvbusdisFieldVbusdtShift))
}

// RegisterOtghsdvbuspulseType OTG_HS device VBUS pulsing time register
type RegisterOtghsdvbuspulseType uint32

func (r *RegisterOtghsdvbuspulseType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdvbuspulseType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdvbuspulseType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdvbuspulseType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdvbuspulseType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdvbuspulseFieldDvbuspShift = 0
	RegisterOtghsdvbuspulseFieldDvbuspMask  = 0xfff
)

// GetDvbusp Device VBUS pulsing time
func (r *RegisterOtghsdvbuspulseType) GetDvbusp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdvbuspulseFieldDvbuspMask) >> RegisterOtghsdvbuspulseFieldDvbuspShift)
}

// SetDvbusp Device VBUS pulsing time
func (r *RegisterOtghsdvbuspulseType) SetDvbusp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdvbuspulseFieldDvbuspMask)|(uint32(value)<<RegisterOtghsdvbuspulseFieldDvbuspShift))
}

// RegisterOtghsdthrctlType OTG_HS Device threshold control register
type RegisterOtghsdthrctlType uint32

func (r *RegisterOtghsdthrctlType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdthrctlType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdthrctlType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdthrctlType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdthrctlType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdthrctlFieldNonisothrenShift = 0
	RegisterOtghsdthrctlFieldNonisothrenMask  = 0x1
)

// GetNonisothren Nonisochronous IN endpoints threshold enable
func (r *RegisterOtghsdthrctlType) GetNonisothren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdthrctlFieldNonisothrenMask) != 0
}

// SetNonisothren Nonisochronous IN endpoints threshold enable
func (r *RegisterOtghsdthrctlType) SetNonisothren(value bool) {
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
func (r *RegisterOtghsdthrctlType) GetIsothren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdthrctlFieldIsothrenMask) != 0
}

// SetIsothren ISO IN endpoint threshold enable
func (r *RegisterOtghsdthrctlType) SetIsothren(value bool) {
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
func (r *RegisterOtghsdthrctlType) GetTxthrlen() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdthrctlFieldTxthrlenMask) >> RegisterOtghsdthrctlFieldTxthrlenShift)
}

// SetTxthrlen Transmit threshold length
func (r *RegisterOtghsdthrctlType) SetTxthrlen(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdthrctlFieldTxthrlenMask)|(uint32(value)<<RegisterOtghsdthrctlFieldTxthrlenShift))
}

const (
	RegisterOtghsdthrctlFieldRxthrenShift = 16
	RegisterOtghsdthrctlFieldRxthrenMask  = 0x10000
)

// GetRxthren Receive threshold enable
func (r *RegisterOtghsdthrctlType) GetRxthren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdthrctlFieldRxthrenMask) != 0
}

// SetRxthren Receive threshold enable
func (r *RegisterOtghsdthrctlType) SetRxthren(value bool) {
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
func (r *RegisterOtghsdthrctlType) GetRxthrlen() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdthrctlFieldRxthrlenMask) >> RegisterOtghsdthrctlFieldRxthrlenShift)
}

// SetRxthrlen Receive threshold length
func (r *RegisterOtghsdthrctlType) SetRxthrlen(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdthrctlFieldRxthrlenMask)|(uint32(value)<<RegisterOtghsdthrctlFieldRxthrlenShift))
}

const (
	RegisterOtghsdthrctlFieldArpenShift = 27
	RegisterOtghsdthrctlFieldArpenMask  = 0x8000000
)

// GetArpen Arbiter parking enable
func (r *RegisterOtghsdthrctlType) GetArpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdthrctlFieldArpenMask) != 0
}

// SetArpen Arbiter parking enable
func (r *RegisterOtghsdthrctlType) SetArpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdthrctlFieldArpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdthrctlFieldArpenMask)
	}
}

// RegisterOtghsdiepempmskType OTG_HS device IN endpoint FIFO empty interrupt mask register
type RegisterOtghsdiepempmskType uint32

func (r *RegisterOtghsdiepempmskType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepempmskType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepempmskType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepempmskType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepempmskType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepempmskFieldIneptxfemShift = 0
	RegisterOtghsdiepempmskFieldIneptxfemMask  = 0xffff
)

// GetIneptxfem IN EP Tx FIFO empty interrupt mask bits
func (r *RegisterOtghsdiepempmskType) GetIneptxfem() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepempmskFieldIneptxfemMask) >> RegisterOtghsdiepempmskFieldIneptxfemShift)
}

// SetIneptxfem IN EP Tx FIFO empty interrupt mask bits
func (r *RegisterOtghsdiepempmskType) SetIneptxfem(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepempmskFieldIneptxfemMask)|(uint32(value)<<RegisterOtghsdiepempmskFieldIneptxfemShift))
}

// RegisterOtghsdeachintType OTG_HS device each endpoint interrupt register
type RegisterOtghsdeachintType uint32

func (r *RegisterOtghsdeachintType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdeachintType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdeachintType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdeachintType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdeachintType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdeachintFieldIep1intShift = 1
	RegisterOtghsdeachintFieldIep1intMask  = 0x2
)

// GetIep1int IN endpoint 1interrupt bit
func (r *RegisterOtghsdeachintType) GetIep1int() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdeachintFieldIep1intMask) != 0
}

// SetIep1int IN endpoint 1interrupt bit
func (r *RegisterOtghsdeachintType) SetIep1int(value bool) {
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
func (r *RegisterOtghsdeachintType) GetOep1int() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdeachintFieldOep1intMask) != 0
}

// SetOep1int OUT endpoint 1 interrupt bit
func (r *RegisterOtghsdeachintType) SetOep1int(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdeachintFieldOep1intMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdeachintFieldOep1intMask)
	}
}

// RegisterOtghsdeachintmskType OTG_HS device each endpoint interrupt register mask
type RegisterOtghsdeachintmskType uint32

func (r *RegisterOtghsdeachintmskType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdeachintmskType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdeachintmskType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdeachintmskType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdeachintmskType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdeachintmskFieldIep1intmShift = 1
	RegisterOtghsdeachintmskFieldIep1intmMask  = 0x2
)

// GetIep1intm IN Endpoint 1 interrupt mask bit
func (r *RegisterOtghsdeachintmskType) GetIep1intm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdeachintmskFieldIep1intmMask) != 0
}

// SetIep1intm IN Endpoint 1 interrupt mask bit
func (r *RegisterOtghsdeachintmskType) SetIep1intm(value bool) {
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
func (r *RegisterOtghsdeachintmskType) GetOep1intm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdeachintmskFieldOep1intmMask) != 0
}

// SetOep1intm OUT Endpoint 1 interrupt mask bit
func (r *RegisterOtghsdeachintmskType) SetOep1intm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdeachintmskFieldOep1intmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdeachintmskFieldOep1intmMask)
	}
}

// RegisterOtghsdiepctl0Type OTG device endpoint-0 control register
type RegisterOtghsdiepctl0Type uint32

func (r *RegisterOtghsdiepctl0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepctl0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepctl0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepctl0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepctl0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepctl0FieldMpsizShift = 0
	RegisterOtghsdiepctl0FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl0Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldMpsizMask) >> RegisterOtghsdiepctl0FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl0Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl0FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl0FieldUsbaepShift = 15
	RegisterOtghsdiepctl0FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl0Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl0Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdiepctl0Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl0FieldNakstsShift = 17
	RegisterOtghsdiepctl0FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdiepctl0Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl0FieldEptypShift = 18
	RegisterOtghsdiepctl0FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdiepctl0Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldEptypMask) >> RegisterOtghsdiepctl0FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdiepctl0Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl0FieldEptypShift))
}

const (
	RegisterOtghsdiepctl0FieldStallShift = 21
	RegisterOtghsdiepctl0FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *RegisterOtghsdiepctl0Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdiepctl0Type) SetStall(value bool) {
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
func (r *RegisterOtghsdiepctl0Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldTxfnumMask) >> RegisterOtghsdiepctl0FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *RegisterOtghsdiepctl0Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl0FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl0FieldCnakShift = 26
	RegisterOtghsdiepctl0FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *RegisterOtghsdiepctl0Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdiepctl0Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdiepctl0Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdiepctl0Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdiepctl0Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdiepctl0Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdiepctl0Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl0FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdiepctl0Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl0FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl0FieldEpenaMask)
	}
}

// RegisterOtghsdiepint0Type OTG device endpoint-0 interrupt register
type RegisterOtghsdiepint0Type uint32

func (r *RegisterOtghsdiepint0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepint0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepint0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepint0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepint0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepint0FieldXfrcShift = 0
	RegisterOtghsdiepint0FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint0Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint0Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdiepint0Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdiepint0Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdiepint0Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *RegisterOtghsdiepint0Type) SetToc(value bool) {
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
func (r *RegisterOtghsdiepint0Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *RegisterOtghsdiepint0Type) SetIttxfe(value bool) {
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
func (r *RegisterOtghsdiepint0Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *RegisterOtghsdiepint0Type) SetInepne(value bool) {
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
func (r *RegisterOtghsdiepint0Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint0FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint0FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint0Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint0Type) SetTxfifoudrn(value bool) {
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
func (r *RegisterOtghsdiepint0Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *RegisterOtghsdiepint0Type) SetBna(value bool) {
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
func (r *RegisterOtghsdiepint0Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *RegisterOtghsdiepint0Type) SetPktdrpsts(value bool) {
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
func (r *RegisterOtghsdiepint0Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *RegisterOtghsdiepint0Type) SetBerr(value bool) {
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
func (r *RegisterOtghsdiepint0Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint0FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *RegisterOtghsdiepint0Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint0FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint0FieldNakMask)
	}
}

// RegisterOtghsdieptsiz0Type OTG_HS device IN endpoint 0 transfer size register
type RegisterOtghsdieptsiz0Type uint32

func (r *RegisterOtghsdieptsiz0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdieptsiz0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdieptsiz0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdieptsiz0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdieptsiz0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdieptsiz0FieldXfrsizShift = 0
	RegisterOtghsdieptsiz0FieldXfrsizMask  = 0x7f
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz0Type) GetXfrsiz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz0FieldXfrsizMask) >> RegisterOtghsdieptsiz0FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz0Type) SetXfrsiz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz0FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz0FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz0FieldPktcntShift = 19
	RegisterOtghsdieptsiz0FieldPktcntMask  = 0x180000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdieptsiz0Type) GetPktcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz0FieldPktcntMask) >> RegisterOtghsdieptsiz0FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdieptsiz0Type) SetPktcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz0FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz0FieldPktcntShift))
}

// RegisterOtghsdiepdma1Type OTG_HS device endpoint-1 DMA address register
type RegisterOtghsdiepdma1Type uint32

func (r *RegisterOtghsdiepdma1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepdma1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepdma1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepdma1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepdma1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepdma1FieldDmaaddrShift = 0
	RegisterOtghsdiepdma1FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghsdiepdma1Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepdma1FieldDmaaddrMask) >> RegisterOtghsdiepdma1FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghsdiepdma1Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepdma1FieldDmaaddrMask)|(uint32(value)<<RegisterOtghsdiepdma1FieldDmaaddrShift))
}

// RegisterOtghsdtxfsts0Type OTG_HS device IN endpoint transmit FIFO status register
type RegisterOtghsdtxfsts0Type uint32

func (r *RegisterOtghsdtxfsts0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdtxfsts0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdtxfsts0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdtxfsts0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdtxfsts0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdtxfsts0FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts0FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts0Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts0FieldIneptfsavMask) >> RegisterOtghsdtxfsts0FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts0Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts0FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts0FieldIneptfsavShift))
}

// RegisterOtghsdiepctl1Type OTG device endpoint-1 control register
type RegisterOtghsdiepctl1Type uint32

func (r *RegisterOtghsdiepctl1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepctl1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepctl1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepctl1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepctl1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepctl1FieldMpsizShift = 0
	RegisterOtghsdiepctl1FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl1Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldMpsizMask) >> RegisterOtghsdiepctl1FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl1Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl1FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl1FieldUsbaepShift = 15
	RegisterOtghsdiepctl1FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl1Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl1Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdiepctl1Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl1FieldNakstsShift = 17
	RegisterOtghsdiepctl1FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdiepctl1Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl1FieldEptypShift = 18
	RegisterOtghsdiepctl1FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdiepctl1Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldEptypMask) >> RegisterOtghsdiepctl1FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdiepctl1Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl1FieldEptypShift))
}

const (
	RegisterOtghsdiepctl1FieldStallShift = 21
	RegisterOtghsdiepctl1FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *RegisterOtghsdiepctl1Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdiepctl1Type) SetStall(value bool) {
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
func (r *RegisterOtghsdiepctl1Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldTxfnumMask) >> RegisterOtghsdiepctl1FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *RegisterOtghsdiepctl1Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl1FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl1FieldCnakShift = 26
	RegisterOtghsdiepctl1FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *RegisterOtghsdiepctl1Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdiepctl1Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdiepctl1Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdiepctl1Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdiepctl1Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdiepctl1Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdiepctl1Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl1FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdiepctl1Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl1FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl1FieldEpenaMask)
	}
}

// RegisterOtghsdiepint1Type OTG device endpoint-1 interrupt register
type RegisterOtghsdiepint1Type uint32

func (r *RegisterOtghsdiepint1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepint1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepint1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepint1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepint1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepint1FieldXfrcShift = 0
	RegisterOtghsdiepint1FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint1Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint1Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdiepint1Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdiepint1Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdiepint1Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *RegisterOtghsdiepint1Type) SetToc(value bool) {
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
func (r *RegisterOtghsdiepint1Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *RegisterOtghsdiepint1Type) SetIttxfe(value bool) {
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
func (r *RegisterOtghsdiepint1Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *RegisterOtghsdiepint1Type) SetInepne(value bool) {
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
func (r *RegisterOtghsdiepint1Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint1FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint1FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint1Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint1Type) SetTxfifoudrn(value bool) {
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
func (r *RegisterOtghsdiepint1Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *RegisterOtghsdiepint1Type) SetBna(value bool) {
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
func (r *RegisterOtghsdiepint1Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *RegisterOtghsdiepint1Type) SetPktdrpsts(value bool) {
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
func (r *RegisterOtghsdiepint1Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *RegisterOtghsdiepint1Type) SetBerr(value bool) {
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
func (r *RegisterOtghsdiepint1Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint1FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *RegisterOtghsdiepint1Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint1FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint1FieldNakMask)
	}
}

// RegisterOtghsdieptsiz1Type OTG_HS device endpoint transfer size register
type RegisterOtghsdieptsiz1Type uint32

func (r *RegisterOtghsdieptsiz1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdieptsiz1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdieptsiz1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdieptsiz1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdieptsiz1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdieptsiz1FieldXfrsizShift = 0
	RegisterOtghsdieptsiz1FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz1Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz1FieldXfrsizMask) >> RegisterOtghsdieptsiz1FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz1Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz1FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz1FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz1FieldPktcntShift = 19
	RegisterOtghsdieptsiz1FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdieptsiz1Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz1FieldPktcntMask) >> RegisterOtghsdieptsiz1FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdieptsiz1Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz1FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz1FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz1FieldMcntShift = 29
	RegisterOtghsdieptsiz1FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *RegisterOtghsdieptsiz1Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz1FieldMcntMask) >> RegisterOtghsdieptsiz1FieldMcntShift)
}

// SetMcnt Multi count
func (r *RegisterOtghsdieptsiz1Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz1FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz1FieldMcntShift))
}

// RegisterOtghsdiepdma2Type OTG_HS device endpoint-2 DMA address register
type RegisterOtghsdiepdma2Type uint32

func (r *RegisterOtghsdiepdma2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepdma2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepdma2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepdma2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepdma2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepdma2FieldDmaaddrShift = 0
	RegisterOtghsdiepdma2FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghsdiepdma2Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepdma2FieldDmaaddrMask) >> RegisterOtghsdiepdma2FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghsdiepdma2Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepdma2FieldDmaaddrMask)|(uint32(value)<<RegisterOtghsdiepdma2FieldDmaaddrShift))
}

// RegisterOtghsdtxfsts1Type OTG_HS device IN endpoint transmit FIFO status register
type RegisterOtghsdtxfsts1Type uint32

func (r *RegisterOtghsdtxfsts1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdtxfsts1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdtxfsts1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdtxfsts1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdtxfsts1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdtxfsts1FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts1FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts1Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts1FieldIneptfsavMask) >> RegisterOtghsdtxfsts1FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts1Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts1FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts1FieldIneptfsavShift))
}

// RegisterOtghsdiepctl2Type OTG device endpoint-2 control register
type RegisterOtghsdiepctl2Type uint32

func (r *RegisterOtghsdiepctl2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepctl2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepctl2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepctl2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepctl2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepctl2FieldMpsizShift = 0
	RegisterOtghsdiepctl2FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl2Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldMpsizMask) >> RegisterOtghsdiepctl2FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl2Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl2FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl2FieldUsbaepShift = 15
	RegisterOtghsdiepctl2FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl2Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl2Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdiepctl2Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl2FieldNakstsShift = 17
	RegisterOtghsdiepctl2FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdiepctl2Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl2FieldEptypShift = 18
	RegisterOtghsdiepctl2FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdiepctl2Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldEptypMask) >> RegisterOtghsdiepctl2FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdiepctl2Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl2FieldEptypShift))
}

const (
	RegisterOtghsdiepctl2FieldStallShift = 21
	RegisterOtghsdiepctl2FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *RegisterOtghsdiepctl2Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdiepctl2Type) SetStall(value bool) {
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
func (r *RegisterOtghsdiepctl2Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldTxfnumMask) >> RegisterOtghsdiepctl2FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *RegisterOtghsdiepctl2Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl2FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl2FieldCnakShift = 26
	RegisterOtghsdiepctl2FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *RegisterOtghsdiepctl2Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdiepctl2Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdiepctl2Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdiepctl2Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdiepctl2Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdiepctl2Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdiepctl2Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl2FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdiepctl2Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl2FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl2FieldEpenaMask)
	}
}

// RegisterOtghsdiepint2Type OTG device endpoint-2 interrupt register
type RegisterOtghsdiepint2Type uint32

func (r *RegisterOtghsdiepint2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepint2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepint2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepint2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepint2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepint2FieldXfrcShift = 0
	RegisterOtghsdiepint2FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint2Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint2Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdiepint2Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdiepint2Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdiepint2Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *RegisterOtghsdiepint2Type) SetToc(value bool) {
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
func (r *RegisterOtghsdiepint2Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *RegisterOtghsdiepint2Type) SetIttxfe(value bool) {
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
func (r *RegisterOtghsdiepint2Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *RegisterOtghsdiepint2Type) SetInepne(value bool) {
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
func (r *RegisterOtghsdiepint2Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint2FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint2FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint2Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint2Type) SetTxfifoudrn(value bool) {
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
func (r *RegisterOtghsdiepint2Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *RegisterOtghsdiepint2Type) SetBna(value bool) {
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
func (r *RegisterOtghsdiepint2Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *RegisterOtghsdiepint2Type) SetPktdrpsts(value bool) {
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
func (r *RegisterOtghsdiepint2Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *RegisterOtghsdiepint2Type) SetBerr(value bool) {
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
func (r *RegisterOtghsdiepint2Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint2FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *RegisterOtghsdiepint2Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint2FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint2FieldNakMask)
	}
}

// RegisterOtghsdieptsiz2Type OTG_HS device endpoint transfer size register
type RegisterOtghsdieptsiz2Type uint32

func (r *RegisterOtghsdieptsiz2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdieptsiz2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdieptsiz2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdieptsiz2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdieptsiz2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdieptsiz2FieldXfrsizShift = 0
	RegisterOtghsdieptsiz2FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz2Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz2FieldXfrsizMask) >> RegisterOtghsdieptsiz2FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz2Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz2FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz2FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz2FieldPktcntShift = 19
	RegisterOtghsdieptsiz2FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdieptsiz2Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz2FieldPktcntMask) >> RegisterOtghsdieptsiz2FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdieptsiz2Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz2FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz2FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz2FieldMcntShift = 29
	RegisterOtghsdieptsiz2FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *RegisterOtghsdieptsiz2Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz2FieldMcntMask) >> RegisterOtghsdieptsiz2FieldMcntShift)
}

// SetMcnt Multi count
func (r *RegisterOtghsdieptsiz2Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz2FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz2FieldMcntShift))
}

// RegisterOtghsdiepdma3Type OTG_HS device endpoint-3 DMA address register
type RegisterOtghsdiepdma3Type uint32

func (r *RegisterOtghsdiepdma3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepdma3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepdma3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepdma3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepdma3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepdma3FieldDmaaddrShift = 0
	RegisterOtghsdiepdma3FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghsdiepdma3Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepdma3FieldDmaaddrMask) >> RegisterOtghsdiepdma3FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghsdiepdma3Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepdma3FieldDmaaddrMask)|(uint32(value)<<RegisterOtghsdiepdma3FieldDmaaddrShift))
}

// RegisterOtghsdtxfsts2Type OTG_HS device IN endpoint transmit FIFO status register
type RegisterOtghsdtxfsts2Type uint32

func (r *RegisterOtghsdtxfsts2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdtxfsts2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdtxfsts2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdtxfsts2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdtxfsts2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdtxfsts2FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts2FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts2Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts2FieldIneptfsavMask) >> RegisterOtghsdtxfsts2FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts2Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts2FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts2FieldIneptfsavShift))
}

// RegisterOtghsdiepctl3Type OTG device endpoint-3 control register
type RegisterOtghsdiepctl3Type uint32

func (r *RegisterOtghsdiepctl3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepctl3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepctl3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepctl3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepctl3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepctl3FieldMpsizShift = 0
	RegisterOtghsdiepctl3FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl3Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldMpsizMask) >> RegisterOtghsdiepctl3FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl3Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl3FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl3FieldUsbaepShift = 15
	RegisterOtghsdiepctl3FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl3Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl3Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdiepctl3Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl3FieldNakstsShift = 17
	RegisterOtghsdiepctl3FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdiepctl3Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl3FieldEptypShift = 18
	RegisterOtghsdiepctl3FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdiepctl3Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldEptypMask) >> RegisterOtghsdiepctl3FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdiepctl3Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl3FieldEptypShift))
}

const (
	RegisterOtghsdiepctl3FieldStallShift = 21
	RegisterOtghsdiepctl3FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *RegisterOtghsdiepctl3Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdiepctl3Type) SetStall(value bool) {
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
func (r *RegisterOtghsdiepctl3Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldTxfnumMask) >> RegisterOtghsdiepctl3FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *RegisterOtghsdiepctl3Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl3FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl3FieldCnakShift = 26
	RegisterOtghsdiepctl3FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *RegisterOtghsdiepctl3Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdiepctl3Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdiepctl3Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdiepctl3Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdiepctl3Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdiepctl3Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdiepctl3Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl3FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdiepctl3Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl3FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl3FieldEpenaMask)
	}
}

// RegisterOtghsdiepint3Type OTG device endpoint-3 interrupt register
type RegisterOtghsdiepint3Type uint32

func (r *RegisterOtghsdiepint3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepint3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepint3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepint3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepint3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepint3FieldXfrcShift = 0
	RegisterOtghsdiepint3FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint3Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint3Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdiepint3Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdiepint3Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdiepint3Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *RegisterOtghsdiepint3Type) SetToc(value bool) {
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
func (r *RegisterOtghsdiepint3Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *RegisterOtghsdiepint3Type) SetIttxfe(value bool) {
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
func (r *RegisterOtghsdiepint3Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *RegisterOtghsdiepint3Type) SetInepne(value bool) {
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
func (r *RegisterOtghsdiepint3Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint3FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint3FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint3Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint3Type) SetTxfifoudrn(value bool) {
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
func (r *RegisterOtghsdiepint3Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *RegisterOtghsdiepint3Type) SetBna(value bool) {
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
func (r *RegisterOtghsdiepint3Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *RegisterOtghsdiepint3Type) SetPktdrpsts(value bool) {
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
func (r *RegisterOtghsdiepint3Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *RegisterOtghsdiepint3Type) SetBerr(value bool) {
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
func (r *RegisterOtghsdiepint3Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint3FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *RegisterOtghsdiepint3Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint3FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint3FieldNakMask)
	}
}

// RegisterOtghsdieptsiz3Type OTG_HS device endpoint transfer size register
type RegisterOtghsdieptsiz3Type uint32

func (r *RegisterOtghsdieptsiz3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdieptsiz3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdieptsiz3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdieptsiz3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdieptsiz3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdieptsiz3FieldXfrsizShift = 0
	RegisterOtghsdieptsiz3FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz3Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz3FieldXfrsizMask) >> RegisterOtghsdieptsiz3FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz3Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz3FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz3FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz3FieldPktcntShift = 19
	RegisterOtghsdieptsiz3FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdieptsiz3Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz3FieldPktcntMask) >> RegisterOtghsdieptsiz3FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdieptsiz3Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz3FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz3FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz3FieldMcntShift = 29
	RegisterOtghsdieptsiz3FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *RegisterOtghsdieptsiz3Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz3FieldMcntMask) >> RegisterOtghsdieptsiz3FieldMcntShift)
}

// SetMcnt Multi count
func (r *RegisterOtghsdieptsiz3Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz3FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz3FieldMcntShift))
}

// RegisterOtghsdiepdma4Type OTG_HS device endpoint-4 DMA address register
type RegisterOtghsdiepdma4Type uint32

func (r *RegisterOtghsdiepdma4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepdma4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepdma4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepdma4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepdma4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepdma4FieldDmaaddrShift = 0
	RegisterOtghsdiepdma4FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghsdiepdma4Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepdma4FieldDmaaddrMask) >> RegisterOtghsdiepdma4FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghsdiepdma4Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepdma4FieldDmaaddrMask)|(uint32(value)<<RegisterOtghsdiepdma4FieldDmaaddrShift))
}

// RegisterOtghsdtxfsts3Type OTG_HS device IN endpoint transmit FIFO status register
type RegisterOtghsdtxfsts3Type uint32

func (r *RegisterOtghsdtxfsts3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdtxfsts3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdtxfsts3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdtxfsts3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdtxfsts3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdtxfsts3FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts3FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts3Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts3FieldIneptfsavMask) >> RegisterOtghsdtxfsts3FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts3Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts3FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts3FieldIneptfsavShift))
}

// RegisterOtghsdiepctl4Type OTG device endpoint-4 control register
type RegisterOtghsdiepctl4Type uint32

func (r *RegisterOtghsdiepctl4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepctl4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepctl4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepctl4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepctl4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepctl4FieldMpsizShift = 0
	RegisterOtghsdiepctl4FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl4Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldMpsizMask) >> RegisterOtghsdiepctl4FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl4Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl4FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl4FieldUsbaepShift = 15
	RegisterOtghsdiepctl4FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl4Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl4Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdiepctl4Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl4FieldNakstsShift = 17
	RegisterOtghsdiepctl4FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdiepctl4Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl4FieldEptypShift = 18
	RegisterOtghsdiepctl4FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdiepctl4Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldEptypMask) >> RegisterOtghsdiepctl4FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdiepctl4Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl4FieldEptypShift))
}

const (
	RegisterOtghsdiepctl4FieldStallShift = 21
	RegisterOtghsdiepctl4FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *RegisterOtghsdiepctl4Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdiepctl4Type) SetStall(value bool) {
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
func (r *RegisterOtghsdiepctl4Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldTxfnumMask) >> RegisterOtghsdiepctl4FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *RegisterOtghsdiepctl4Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl4FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl4FieldCnakShift = 26
	RegisterOtghsdiepctl4FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *RegisterOtghsdiepctl4Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdiepctl4Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdiepctl4Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdiepctl4Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdiepctl4Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdiepctl4Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdiepctl4Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl4FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdiepctl4Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl4FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl4FieldEpenaMask)
	}
}

// RegisterOtghsdiepint4Type OTG device endpoint-4 interrupt register
type RegisterOtghsdiepint4Type uint32

func (r *RegisterOtghsdiepint4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepint4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepint4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepint4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepint4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepint4FieldXfrcShift = 0
	RegisterOtghsdiepint4FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint4Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint4Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdiepint4Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdiepint4Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdiepint4Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *RegisterOtghsdiepint4Type) SetToc(value bool) {
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
func (r *RegisterOtghsdiepint4Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *RegisterOtghsdiepint4Type) SetIttxfe(value bool) {
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
func (r *RegisterOtghsdiepint4Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *RegisterOtghsdiepint4Type) SetInepne(value bool) {
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
func (r *RegisterOtghsdiepint4Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint4FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint4FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint4Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint4Type) SetTxfifoudrn(value bool) {
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
func (r *RegisterOtghsdiepint4Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *RegisterOtghsdiepint4Type) SetBna(value bool) {
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
func (r *RegisterOtghsdiepint4Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *RegisterOtghsdiepint4Type) SetPktdrpsts(value bool) {
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
func (r *RegisterOtghsdiepint4Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *RegisterOtghsdiepint4Type) SetBerr(value bool) {
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
func (r *RegisterOtghsdiepint4Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint4FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *RegisterOtghsdiepint4Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint4FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint4FieldNakMask)
	}
}

// RegisterOtghsdieptsiz4Type OTG_HS device endpoint transfer size register
type RegisterOtghsdieptsiz4Type uint32

func (r *RegisterOtghsdieptsiz4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdieptsiz4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdieptsiz4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdieptsiz4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdieptsiz4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdieptsiz4FieldXfrsizShift = 0
	RegisterOtghsdieptsiz4FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz4Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz4FieldXfrsizMask) >> RegisterOtghsdieptsiz4FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz4Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz4FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz4FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz4FieldPktcntShift = 19
	RegisterOtghsdieptsiz4FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdieptsiz4Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz4FieldPktcntMask) >> RegisterOtghsdieptsiz4FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdieptsiz4Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz4FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz4FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz4FieldMcntShift = 29
	RegisterOtghsdieptsiz4FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *RegisterOtghsdieptsiz4Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz4FieldMcntMask) >> RegisterOtghsdieptsiz4FieldMcntShift)
}

// SetMcnt Multi count
func (r *RegisterOtghsdieptsiz4Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz4FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz4FieldMcntShift))
}

// RegisterOtghsdiepdma5Type OTG_HS device endpoint-5 DMA address register
type RegisterOtghsdiepdma5Type uint32

func (r *RegisterOtghsdiepdma5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepdma5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepdma5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepdma5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepdma5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepdma5FieldDmaaddrShift = 0
	RegisterOtghsdiepdma5FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghsdiepdma5Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepdma5FieldDmaaddrMask) >> RegisterOtghsdiepdma5FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghsdiepdma5Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepdma5FieldDmaaddrMask)|(uint32(value)<<RegisterOtghsdiepdma5FieldDmaaddrShift))
}

// RegisterOtghsdtxfsts4Type OTG_HS device IN endpoint transmit FIFO status register
type RegisterOtghsdtxfsts4Type uint32

func (r *RegisterOtghsdtxfsts4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdtxfsts4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdtxfsts4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdtxfsts4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdtxfsts4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdtxfsts4FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts4FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts4Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts4FieldIneptfsavMask) >> RegisterOtghsdtxfsts4FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts4Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts4FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts4FieldIneptfsavShift))
}

// RegisterOtghsdiepctl5Type OTG device endpoint-5 control register
type RegisterOtghsdiepctl5Type uint32

func (r *RegisterOtghsdiepctl5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepctl5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepctl5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepctl5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepctl5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepctl5FieldMpsizShift = 0
	RegisterOtghsdiepctl5FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl5Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldMpsizMask) >> RegisterOtghsdiepctl5FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl5Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl5FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl5FieldUsbaepShift = 15
	RegisterOtghsdiepctl5FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl5Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl5Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdiepctl5Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl5FieldNakstsShift = 17
	RegisterOtghsdiepctl5FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdiepctl5Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl5FieldEptypShift = 18
	RegisterOtghsdiepctl5FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdiepctl5Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldEptypMask) >> RegisterOtghsdiepctl5FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdiepctl5Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl5FieldEptypShift))
}

const (
	RegisterOtghsdiepctl5FieldStallShift = 21
	RegisterOtghsdiepctl5FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *RegisterOtghsdiepctl5Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdiepctl5Type) SetStall(value bool) {
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
func (r *RegisterOtghsdiepctl5Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldTxfnumMask) >> RegisterOtghsdiepctl5FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *RegisterOtghsdiepctl5Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl5FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl5FieldCnakShift = 26
	RegisterOtghsdiepctl5FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *RegisterOtghsdiepctl5Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdiepctl5Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdiepctl5Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdiepctl5Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdiepctl5Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdiepctl5Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdiepctl5Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl5FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdiepctl5Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl5FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl5FieldEpenaMask)
	}
}

// RegisterOtghsdieptsiz6Type OTG_HS device endpoint transfer size register
type RegisterOtghsdieptsiz6Type uint32

func (r *RegisterOtghsdieptsiz6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdieptsiz6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdieptsiz6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdieptsiz6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdieptsiz6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdieptsiz6FieldXfrsizShift = 0
	RegisterOtghsdieptsiz6FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz6Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz6FieldXfrsizMask) >> RegisterOtghsdieptsiz6FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz6Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz6FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz6FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz6FieldPktcntShift = 19
	RegisterOtghsdieptsiz6FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdieptsiz6Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz6FieldPktcntMask) >> RegisterOtghsdieptsiz6FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdieptsiz6Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz6FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz6FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz6FieldMcntShift = 29
	RegisterOtghsdieptsiz6FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *RegisterOtghsdieptsiz6Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz6FieldMcntMask) >> RegisterOtghsdieptsiz6FieldMcntShift)
}

// SetMcnt Multi count
func (r *RegisterOtghsdieptsiz6Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz6FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz6FieldMcntShift))
}

// RegisterOtghsdtxfsts6Type OTG_HS device IN endpoint transmit FIFO status register
type RegisterOtghsdtxfsts6Type uint32

func (r *RegisterOtghsdtxfsts6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdtxfsts6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdtxfsts6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdtxfsts6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdtxfsts6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdtxfsts6FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts6FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts6Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts6FieldIneptfsavMask) >> RegisterOtghsdtxfsts6FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts6Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts6FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts6FieldIneptfsavShift))
}

// RegisterOtghsdieptsiz7Type OTG_HS device endpoint transfer size register
type RegisterOtghsdieptsiz7Type uint32

func (r *RegisterOtghsdieptsiz7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdieptsiz7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdieptsiz7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdieptsiz7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdieptsiz7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdieptsiz7FieldXfrsizShift = 0
	RegisterOtghsdieptsiz7FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz7Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz7FieldXfrsizMask) >> RegisterOtghsdieptsiz7FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz7Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz7FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz7FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz7FieldPktcntShift = 19
	RegisterOtghsdieptsiz7FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdieptsiz7Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz7FieldPktcntMask) >> RegisterOtghsdieptsiz7FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdieptsiz7Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz7FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz7FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz7FieldMcntShift = 29
	RegisterOtghsdieptsiz7FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *RegisterOtghsdieptsiz7Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz7FieldMcntMask) >> RegisterOtghsdieptsiz7FieldMcntShift)
}

// SetMcnt Multi count
func (r *RegisterOtghsdieptsiz7Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz7FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz7FieldMcntShift))
}

// RegisterOtghsdiepint5Type OTG device endpoint-5 interrupt register
type RegisterOtghsdiepint5Type uint32

func (r *RegisterOtghsdiepint5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepint5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepint5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepint5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepint5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepint5FieldXfrcShift = 0
	RegisterOtghsdiepint5FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint5Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint5Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdiepint5Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdiepint5Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdiepint5Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *RegisterOtghsdiepint5Type) SetToc(value bool) {
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
func (r *RegisterOtghsdiepint5Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *RegisterOtghsdiepint5Type) SetIttxfe(value bool) {
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
func (r *RegisterOtghsdiepint5Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *RegisterOtghsdiepint5Type) SetInepne(value bool) {
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
func (r *RegisterOtghsdiepint5Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint5FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint5FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint5Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint5Type) SetTxfifoudrn(value bool) {
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
func (r *RegisterOtghsdiepint5Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *RegisterOtghsdiepint5Type) SetBna(value bool) {
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
func (r *RegisterOtghsdiepint5Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *RegisterOtghsdiepint5Type) SetPktdrpsts(value bool) {
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
func (r *RegisterOtghsdiepint5Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *RegisterOtghsdiepint5Type) SetBerr(value bool) {
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
func (r *RegisterOtghsdiepint5Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint5FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *RegisterOtghsdiepint5Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint5FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint5FieldNakMask)
	}
}

// RegisterOtghsdtxfsts7Type OTG_HS device IN endpoint transmit FIFO status register
type RegisterOtghsdtxfsts7Type uint32

func (r *RegisterOtghsdtxfsts7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdtxfsts7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdtxfsts7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdtxfsts7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdtxfsts7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdtxfsts7FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts7FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts7Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts7FieldIneptfsavMask) >> RegisterOtghsdtxfsts7FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts7Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts7FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts7FieldIneptfsavShift))
}

// RegisterOtghsdieptsiz5Type OTG_HS device endpoint transfer size register
type RegisterOtghsdieptsiz5Type uint32

func (r *RegisterOtghsdieptsiz5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdieptsiz5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdieptsiz5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdieptsiz5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdieptsiz5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdieptsiz5FieldXfrsizShift = 0
	RegisterOtghsdieptsiz5FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz5Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz5FieldXfrsizMask) >> RegisterOtghsdieptsiz5FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdieptsiz5Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz5FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdieptsiz5FieldXfrsizShift))
}

const (
	RegisterOtghsdieptsiz5FieldPktcntShift = 19
	RegisterOtghsdieptsiz5FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdieptsiz5Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz5FieldPktcntMask) >> RegisterOtghsdieptsiz5FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdieptsiz5Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz5FieldPktcntMask)|(uint32(value)<<RegisterOtghsdieptsiz5FieldPktcntShift))
}

const (
	RegisterOtghsdieptsiz5FieldMcntShift = 29
	RegisterOtghsdieptsiz5FieldMcntMask  = 0x60000000
)

// GetMcnt Multi count
func (r *RegisterOtghsdieptsiz5Type) GetMcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptsiz5FieldMcntMask) >> RegisterOtghsdieptsiz5FieldMcntShift)
}

// SetMcnt Multi count
func (r *RegisterOtghsdieptsiz5Type) SetMcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptsiz5FieldMcntMask)|(uint32(value)<<RegisterOtghsdieptsiz5FieldMcntShift))
}

// RegisterOtghsdtxfsts5Type OTG_HS device IN endpoint transmit FIFO status register
type RegisterOtghsdtxfsts5Type uint32

func (r *RegisterOtghsdtxfsts5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdtxfsts5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdtxfsts5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdtxfsts5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdtxfsts5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdtxfsts5FieldIneptfsavShift = 0
	RegisterOtghsdtxfsts5FieldIneptfsavMask  = 0xffff
)

// GetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts5Type) GetIneptfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdtxfsts5FieldIneptfsavMask) >> RegisterOtghsdtxfsts5FieldIneptfsavShift)
}

// SetIneptfsav IN endpoint TxFIFO space avail
func (r *RegisterOtghsdtxfsts5Type) SetIneptfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdtxfsts5FieldIneptfsavMask)|(uint32(value)<<RegisterOtghsdtxfsts5FieldIneptfsavShift))
}

// RegisterOtghsdiepctl6Type OTG device endpoint-6 control register
type RegisterOtghsdiepctl6Type uint32

func (r *RegisterOtghsdiepctl6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepctl6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepctl6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepctl6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepctl6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepctl6FieldMpsizShift = 0
	RegisterOtghsdiepctl6FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl6Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldMpsizMask) >> RegisterOtghsdiepctl6FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl6Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl6FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl6FieldUsbaepShift = 15
	RegisterOtghsdiepctl6FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl6Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl6Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdiepctl6Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl6FieldNakstsShift = 17
	RegisterOtghsdiepctl6FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdiepctl6Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl6FieldEptypShift = 18
	RegisterOtghsdiepctl6FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdiepctl6Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldEptypMask) >> RegisterOtghsdiepctl6FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdiepctl6Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl6FieldEptypShift))
}

const (
	RegisterOtghsdiepctl6FieldStallShift = 21
	RegisterOtghsdiepctl6FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *RegisterOtghsdiepctl6Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdiepctl6Type) SetStall(value bool) {
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
func (r *RegisterOtghsdiepctl6Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldTxfnumMask) >> RegisterOtghsdiepctl6FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *RegisterOtghsdiepctl6Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl6FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl6FieldCnakShift = 26
	RegisterOtghsdiepctl6FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *RegisterOtghsdiepctl6Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdiepctl6Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdiepctl6Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdiepctl6Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdiepctl6Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdiepctl6Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdiepctl6Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl6FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdiepctl6Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl6FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl6FieldEpenaMask)
	}
}

// RegisterOtghsdiepint6Type OTG device endpoint-6 interrupt register
type RegisterOtghsdiepint6Type uint32

func (r *RegisterOtghsdiepint6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepint6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepint6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepint6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepint6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepint6FieldXfrcShift = 0
	RegisterOtghsdiepint6FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint6Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint6Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdiepint6Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdiepint6Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdiepint6Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *RegisterOtghsdiepint6Type) SetToc(value bool) {
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
func (r *RegisterOtghsdiepint6Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *RegisterOtghsdiepint6Type) SetIttxfe(value bool) {
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
func (r *RegisterOtghsdiepint6Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *RegisterOtghsdiepint6Type) SetInepne(value bool) {
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
func (r *RegisterOtghsdiepint6Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint6FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint6FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint6Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint6Type) SetTxfifoudrn(value bool) {
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
func (r *RegisterOtghsdiepint6Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *RegisterOtghsdiepint6Type) SetBna(value bool) {
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
func (r *RegisterOtghsdiepint6Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *RegisterOtghsdiepint6Type) SetPktdrpsts(value bool) {
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
func (r *RegisterOtghsdiepint6Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *RegisterOtghsdiepint6Type) SetBerr(value bool) {
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
func (r *RegisterOtghsdiepint6Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint6FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *RegisterOtghsdiepint6Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint6FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint6FieldNakMask)
	}
}

// RegisterOtghsdiepctl7Type OTG device endpoint-7 control register
type RegisterOtghsdiepctl7Type uint32

func (r *RegisterOtghsdiepctl7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepctl7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepctl7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepctl7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepctl7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepctl7FieldMpsizShift = 0
	RegisterOtghsdiepctl7FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl7Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldMpsizMask) >> RegisterOtghsdiepctl7FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdiepctl7Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldMpsizMask)|(uint32(value)<<RegisterOtghsdiepctl7FieldMpsizShift))
}

const (
	RegisterOtghsdiepctl7FieldUsbaepShift = 15
	RegisterOtghsdiepctl7FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl7Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdiepctl7Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdiepctl7Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdiepctl7FieldNakstsShift = 17
	RegisterOtghsdiepctl7FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdiepctl7Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldNakstsMask) != 0
}

const (
	RegisterOtghsdiepctl7FieldEptypShift = 18
	RegisterOtghsdiepctl7FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdiepctl7Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldEptypMask) >> RegisterOtghsdiepctl7FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdiepctl7Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldEptypMask)|(uint32(value)<<RegisterOtghsdiepctl7FieldEptypShift))
}

const (
	RegisterOtghsdiepctl7FieldStallShift = 21
	RegisterOtghsdiepctl7FieldStallMask  = 0x200000
)

// GetStall STALL handshake
func (r *RegisterOtghsdiepctl7Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdiepctl7Type) SetStall(value bool) {
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
func (r *RegisterOtghsdiepctl7Type) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldTxfnumMask) >> RegisterOtghsdiepctl7FieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *RegisterOtghsdiepctl7Type) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldTxfnumMask)|(uint32(value)<<RegisterOtghsdiepctl7FieldTxfnumShift))
}

const (
	RegisterOtghsdiepctl7FieldCnakShift = 26
	RegisterOtghsdiepctl7FieldCnakMask  = 0x4000000
)

// SetCnak Clear NAK
func (r *RegisterOtghsdiepctl7Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdiepctl7Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdiepctl7Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdiepctl7Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdiepctl7Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdiepctl7Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdiepctl7Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepctl7FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdiepctl7Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepctl7FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepctl7FieldEpenaMask)
	}
}

// RegisterOtghsdiepint7Type OTG device endpoint-7 interrupt register
type RegisterOtghsdiepint7Type uint32

func (r *RegisterOtghsdiepint7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdiepint7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdiepint7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdiepint7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdiepint7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdiepint7FieldXfrcShift = 0
	RegisterOtghsdiepint7FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint7Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdiepint7Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdiepint7Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdiepint7Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdiepint7Type) GetToc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldTocMask) != 0
}

// SetToc Timeout condition
func (r *RegisterOtghsdiepint7Type) SetToc(value bool) {
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
func (r *RegisterOtghsdiepint7Type) GetIttxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldIttxfeMask) != 0
}

// SetIttxfe IN token received when TxFIFO is empty
func (r *RegisterOtghsdiepint7Type) SetIttxfe(value bool) {
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
func (r *RegisterOtghsdiepint7Type) GetInepne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldInepneMask) != 0
}

// SetInepne IN endpoint NAK effective
func (r *RegisterOtghsdiepint7Type) SetInepne(value bool) {
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
func (r *RegisterOtghsdiepint7Type) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldTxfeMask) != 0
}

const (
	RegisterOtghsdiepint7FieldTxfifoudrnShift = 8
	RegisterOtghsdiepint7FieldTxfifoudrnMask  = 0x100
)

// GetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint7Type) GetTxfifoudrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldTxfifoudrnMask) != 0
}

// SetTxfifoudrn Transmit Fifo Underrun
func (r *RegisterOtghsdiepint7Type) SetTxfifoudrn(value bool) {
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
func (r *RegisterOtghsdiepint7Type) GetBna() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldBnaMask) != 0
}

// SetBna Buffer not available interrupt
func (r *RegisterOtghsdiepint7Type) SetBna(value bool) {
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
func (r *RegisterOtghsdiepint7Type) GetPktdrpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldPktdrpstsMask) != 0
}

// SetPktdrpsts Packet dropped status
func (r *RegisterOtghsdiepint7Type) SetPktdrpsts(value bool) {
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
func (r *RegisterOtghsdiepint7Type) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldBerrMask) != 0
}

// SetBerr Babble error interrupt
func (r *RegisterOtghsdiepint7Type) SetBerr(value bool) {
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
func (r *RegisterOtghsdiepint7Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdiepint7FieldNakMask) != 0
}

// SetNak NAK interrupt
func (r *RegisterOtghsdiepint7Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdiepint7FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdiepint7FieldNakMask)
	}
}

// RegisterOtghsdoepctl0Type OTG_HS device control OUT endpoint 0 control register
type RegisterOtghsdoepctl0Type uint32

func (r *RegisterOtghsdoepctl0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepctl0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepctl0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepctl0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepctl0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepctl0FieldMpsizShift = 0
	RegisterOtghsdoepctl0FieldMpsizMask  = 0x3
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl0Type) GetMpsiz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldMpsizMask) >> RegisterOtghsdoepctl0FieldMpsizShift)
}

const (
	RegisterOtghsdoepctl0FieldUsbaepShift = 15
	RegisterOtghsdoepctl0FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl0Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldUsbaepMask) != 0
}

const (
	RegisterOtghsdoepctl0FieldNakstsShift = 17
	RegisterOtghsdoepctl0FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdoepctl0Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl0FieldEptypShift = 18
	RegisterOtghsdoepctl0FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdoepctl0Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldEptypMask) >> RegisterOtghsdoepctl0FieldEptypShift)
}

const (
	RegisterOtghsdoepctl0FieldSnpmShift = 20
	RegisterOtghsdoepctl0FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *RegisterOtghsdoepctl0Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *RegisterOtghsdoepctl0Type) SetSnpm(value bool) {
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
func (r *RegisterOtghsdoepctl0Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdoepctl0Type) SetStall(value bool) {
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
func (r *RegisterOtghsdoepctl0Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdoepctl0Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdoepctl0Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl0FieldEpdisMask) != 0
}

const (
	RegisterOtghsdoepctl0FieldEpenaShift = 31
	RegisterOtghsdoepctl0FieldEpenaMask  = 0x80000000
)

// SetEpena Endpoint enable
func (r *RegisterOtghsdoepctl0Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl0FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl0FieldEpenaMask)
	}
}

// RegisterOtghsdoepint0Type OTG_HS device endpoint-0 interrupt register
type RegisterOtghsdoepint0Type uint32

func (r *RegisterOtghsdoepint0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepint0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepint0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepint0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepint0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepint0FieldXfrcShift = 0
	RegisterOtghsdoepint0FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint0Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint0FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint0Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdoepint0Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint0FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdoepint0Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdoepint0Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint0FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *RegisterOtghsdoepint0Type) SetStup(value bool) {
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
func (r *RegisterOtghsdoepint0Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint0FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *RegisterOtghsdoepint0Type) SetOtepdis(value bool) {
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
func (r *RegisterOtghsdoepint0Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint0FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *RegisterOtghsdoepint0Type) SetB2bstup(value bool) {
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
func (r *RegisterOtghsdoepint0Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint0FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *RegisterOtghsdoepint0Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint0FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint0FieldNyetMask)
	}
}

// RegisterOtghsdoeptsiz0Type OTG_HS device endpoint-0 transfer size register
type RegisterOtghsdoeptsiz0Type uint32

func (r *RegisterOtghsdoeptsiz0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoeptsiz0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoeptsiz0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoeptsiz0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoeptsiz0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoeptsiz0FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz0FieldXfrsizMask  = 0x7f
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz0Type) GetXfrsiz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz0FieldXfrsizMask) >> RegisterOtghsdoeptsiz0FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz0Type) SetXfrsiz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz0FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz0FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz0FieldPktcntShift = 19
	RegisterOtghsdoeptsiz0FieldPktcntMask  = 0x80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz0Type) GetPktcnt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz0FieldPktcntMask) != 0
}

// SetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz0Type) SetPktcnt(value bool) {
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
func (r *RegisterOtghsdoeptsiz0Type) GetStupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz0FieldStupcntMask) >> RegisterOtghsdoeptsiz0FieldStupcntShift)
}

// SetStupcnt SETUP packet count
func (r *RegisterOtghsdoeptsiz0Type) SetStupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz0FieldStupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz0FieldStupcntShift))
}

// RegisterOtghsdoepctl1Type OTG device endpoint-1 control register
type RegisterOtghsdoepctl1Type uint32

func (r *RegisterOtghsdoepctl1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepctl1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepctl1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepctl1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepctl1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepctl1FieldMpsizShift = 0
	RegisterOtghsdoepctl1FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl1Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldMpsizMask) >> RegisterOtghsdoepctl1FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl1Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl1FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl1FieldUsbaepShift = 15
	RegisterOtghsdoepctl1FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl1Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl1Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdoepctl1Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl1FieldNakstsShift = 17
	RegisterOtghsdoepctl1FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdoepctl1Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl1FieldEptypShift = 18
	RegisterOtghsdoepctl1FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdoepctl1Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldEptypMask) >> RegisterOtghsdoepctl1FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdoepctl1Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl1FieldEptypShift))
}

const (
	RegisterOtghsdoepctl1FieldSnpmShift = 20
	RegisterOtghsdoepctl1FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *RegisterOtghsdoepctl1Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *RegisterOtghsdoepctl1Type) SetSnpm(value bool) {
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
func (r *RegisterOtghsdoepctl1Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdoepctl1Type) SetStall(value bool) {
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
func (r *RegisterOtghsdoepctl1Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdoepctl1Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdoepctl1Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdoepctl1Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdoepctl1Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdoepctl1Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdoepctl1Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl1FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdoepctl1Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl1FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl1FieldEpenaMask)
	}
}

// RegisterOtghsdoepint1Type OTG_HS device endpoint-1 interrupt register
type RegisterOtghsdoepint1Type uint32

func (r *RegisterOtghsdoepint1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepint1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepint1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepint1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepint1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepint1FieldXfrcShift = 0
	RegisterOtghsdoepint1FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint1Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint1FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint1Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdoepint1Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint1FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdoepint1Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdoepint1Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint1FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *RegisterOtghsdoepint1Type) SetStup(value bool) {
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
func (r *RegisterOtghsdoepint1Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint1FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *RegisterOtghsdoepint1Type) SetOtepdis(value bool) {
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
func (r *RegisterOtghsdoepint1Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint1FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *RegisterOtghsdoepint1Type) SetB2bstup(value bool) {
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
func (r *RegisterOtghsdoepint1Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint1FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *RegisterOtghsdoepint1Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint1FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint1FieldNyetMask)
	}
}

// RegisterOtghsdoeptsiz1Type OTG_HS device endpoint-1 transfer size register
type RegisterOtghsdoeptsiz1Type uint32

func (r *RegisterOtghsdoeptsiz1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoeptsiz1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoeptsiz1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoeptsiz1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoeptsiz1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoeptsiz1FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz1FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz1Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz1FieldXfrsizMask) >> RegisterOtghsdoeptsiz1FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz1Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz1FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz1FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz1FieldPktcntShift = 19
	RegisterOtghsdoeptsiz1FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz1Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz1FieldPktcntMask) >> RegisterOtghsdoeptsiz1FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz1Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz1FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz1FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz1FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz1FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz1Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz1FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz1FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz1Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz1FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz1FieldRxdpidstupcntShift))
}

// RegisterOtghsdoepctl2Type OTG device endpoint-2 control register
type RegisterOtghsdoepctl2Type uint32

func (r *RegisterOtghsdoepctl2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepctl2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepctl2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepctl2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepctl2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepctl2FieldMpsizShift = 0
	RegisterOtghsdoepctl2FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl2Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldMpsizMask) >> RegisterOtghsdoepctl2FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl2Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl2FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl2FieldUsbaepShift = 15
	RegisterOtghsdoepctl2FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl2Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl2Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdoepctl2Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl2FieldNakstsShift = 17
	RegisterOtghsdoepctl2FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdoepctl2Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl2FieldEptypShift = 18
	RegisterOtghsdoepctl2FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdoepctl2Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldEptypMask) >> RegisterOtghsdoepctl2FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdoepctl2Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl2FieldEptypShift))
}

const (
	RegisterOtghsdoepctl2FieldSnpmShift = 20
	RegisterOtghsdoepctl2FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *RegisterOtghsdoepctl2Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *RegisterOtghsdoepctl2Type) SetSnpm(value bool) {
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
func (r *RegisterOtghsdoepctl2Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdoepctl2Type) SetStall(value bool) {
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
func (r *RegisterOtghsdoepctl2Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdoepctl2Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdoepctl2Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdoepctl2Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdoepctl2Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdoepctl2Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdoepctl2Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl2FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdoepctl2Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl2FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl2FieldEpenaMask)
	}
}

// RegisterOtghsdoepint2Type OTG_HS device endpoint-2 interrupt register
type RegisterOtghsdoepint2Type uint32

func (r *RegisterOtghsdoepint2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepint2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepint2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepint2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepint2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepint2FieldXfrcShift = 0
	RegisterOtghsdoepint2FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint2Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint2FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint2Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdoepint2Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint2FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdoepint2Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdoepint2Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint2FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *RegisterOtghsdoepint2Type) SetStup(value bool) {
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
func (r *RegisterOtghsdoepint2Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint2FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *RegisterOtghsdoepint2Type) SetOtepdis(value bool) {
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
func (r *RegisterOtghsdoepint2Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint2FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *RegisterOtghsdoepint2Type) SetB2bstup(value bool) {
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
func (r *RegisterOtghsdoepint2Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint2FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *RegisterOtghsdoepint2Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint2FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint2FieldNyetMask)
	}
}

// RegisterOtghsdoeptsiz2Type OTG_HS device endpoint-2 transfer size register
type RegisterOtghsdoeptsiz2Type uint32

func (r *RegisterOtghsdoeptsiz2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoeptsiz2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoeptsiz2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoeptsiz2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoeptsiz2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoeptsiz2FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz2FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz2Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz2FieldXfrsizMask) >> RegisterOtghsdoeptsiz2FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz2Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz2FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz2FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz2FieldPktcntShift = 19
	RegisterOtghsdoeptsiz2FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz2Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz2FieldPktcntMask) >> RegisterOtghsdoeptsiz2FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz2Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz2FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz2FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz2FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz2FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz2Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz2FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz2FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz2Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz2FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz2FieldRxdpidstupcntShift))
}

// RegisterOtghsdoepctl3Type OTG device endpoint-3 control register
type RegisterOtghsdoepctl3Type uint32

func (r *RegisterOtghsdoepctl3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepctl3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepctl3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepctl3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepctl3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepctl3FieldMpsizShift = 0
	RegisterOtghsdoepctl3FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl3Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldMpsizMask) >> RegisterOtghsdoepctl3FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl3Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl3FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl3FieldUsbaepShift = 15
	RegisterOtghsdoepctl3FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl3Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl3Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdoepctl3Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl3FieldNakstsShift = 17
	RegisterOtghsdoepctl3FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdoepctl3Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl3FieldEptypShift = 18
	RegisterOtghsdoepctl3FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdoepctl3Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldEptypMask) >> RegisterOtghsdoepctl3FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdoepctl3Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl3FieldEptypShift))
}

const (
	RegisterOtghsdoepctl3FieldSnpmShift = 20
	RegisterOtghsdoepctl3FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *RegisterOtghsdoepctl3Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *RegisterOtghsdoepctl3Type) SetSnpm(value bool) {
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
func (r *RegisterOtghsdoepctl3Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdoepctl3Type) SetStall(value bool) {
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
func (r *RegisterOtghsdoepctl3Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdoepctl3Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdoepctl3Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdoepctl3Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdoepctl3Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdoepctl3Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdoepctl3Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl3FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdoepctl3Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl3FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl3FieldEpenaMask)
	}
}

// RegisterOtghsdoepint3Type OTG_HS device endpoint-3 interrupt register
type RegisterOtghsdoepint3Type uint32

func (r *RegisterOtghsdoepint3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepint3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepint3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepint3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepint3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepint3FieldXfrcShift = 0
	RegisterOtghsdoepint3FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint3Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint3FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint3Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdoepint3Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint3FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdoepint3Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdoepint3Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint3FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *RegisterOtghsdoepint3Type) SetStup(value bool) {
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
func (r *RegisterOtghsdoepint3Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint3FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *RegisterOtghsdoepint3Type) SetOtepdis(value bool) {
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
func (r *RegisterOtghsdoepint3Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint3FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *RegisterOtghsdoepint3Type) SetB2bstup(value bool) {
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
func (r *RegisterOtghsdoepint3Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint3FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *RegisterOtghsdoepint3Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint3FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint3FieldNyetMask)
	}
}

// RegisterOtghsdoeptsiz3Type OTG_HS device endpoint-3 transfer size register
type RegisterOtghsdoeptsiz3Type uint32

func (r *RegisterOtghsdoeptsiz3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoeptsiz3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoeptsiz3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoeptsiz3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoeptsiz3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoeptsiz3FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz3FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz3Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz3FieldXfrsizMask) >> RegisterOtghsdoeptsiz3FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz3Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz3FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz3FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz3FieldPktcntShift = 19
	RegisterOtghsdoeptsiz3FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz3Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz3FieldPktcntMask) >> RegisterOtghsdoeptsiz3FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz3Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz3FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz3FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz3FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz3FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz3Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz3FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz3FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz3Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz3FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz3FieldRxdpidstupcntShift))
}

// RegisterOtghsdoepctl4Type OTG device endpoint-4 control register
type RegisterOtghsdoepctl4Type uint32

func (r *RegisterOtghsdoepctl4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepctl4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepctl4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepctl4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepctl4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepctl4FieldMpsizShift = 0
	RegisterOtghsdoepctl4FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl4Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldMpsizMask) >> RegisterOtghsdoepctl4FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl4Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl4FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl4FieldUsbaepShift = 15
	RegisterOtghsdoepctl4FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl4Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl4Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdoepctl4Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl4FieldNakstsShift = 17
	RegisterOtghsdoepctl4FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdoepctl4Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl4FieldEptypShift = 18
	RegisterOtghsdoepctl4FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdoepctl4Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldEptypMask) >> RegisterOtghsdoepctl4FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdoepctl4Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl4FieldEptypShift))
}

const (
	RegisterOtghsdoepctl4FieldSnpmShift = 20
	RegisterOtghsdoepctl4FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *RegisterOtghsdoepctl4Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *RegisterOtghsdoepctl4Type) SetSnpm(value bool) {
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
func (r *RegisterOtghsdoepctl4Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdoepctl4Type) SetStall(value bool) {
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
func (r *RegisterOtghsdoepctl4Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdoepctl4Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdoepctl4Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdoepctl4Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdoepctl4Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdoepctl4Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdoepctl4Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl4FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdoepctl4Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl4FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl4FieldEpenaMask)
	}
}

// RegisterOtghsdoepint4Type OTG_HS device endpoint-4 interrupt register
type RegisterOtghsdoepint4Type uint32

func (r *RegisterOtghsdoepint4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepint4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepint4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepint4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepint4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepint4FieldXfrcShift = 0
	RegisterOtghsdoepint4FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint4Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint4FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint4Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdoepint4Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint4FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdoepint4Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdoepint4Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint4FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *RegisterOtghsdoepint4Type) SetStup(value bool) {
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
func (r *RegisterOtghsdoepint4Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint4FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *RegisterOtghsdoepint4Type) SetOtepdis(value bool) {
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
func (r *RegisterOtghsdoepint4Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint4FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *RegisterOtghsdoepint4Type) SetB2bstup(value bool) {
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
func (r *RegisterOtghsdoepint4Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint4FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *RegisterOtghsdoepint4Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint4FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint4FieldNyetMask)
	}
}

// RegisterOtghsdoeptsiz4Type OTG_HS device endpoint-4 transfer size register
type RegisterOtghsdoeptsiz4Type uint32

func (r *RegisterOtghsdoeptsiz4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoeptsiz4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoeptsiz4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoeptsiz4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoeptsiz4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoeptsiz4FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz4FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz4Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz4FieldXfrsizMask) >> RegisterOtghsdoeptsiz4FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz4Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz4FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz4FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz4FieldPktcntShift = 19
	RegisterOtghsdoeptsiz4FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz4Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz4FieldPktcntMask) >> RegisterOtghsdoeptsiz4FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz4Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz4FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz4FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz4FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz4FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz4Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz4FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz4FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz4Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz4FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz4FieldRxdpidstupcntShift))
}

// RegisterOtghsdoepctl5Type OTG device endpoint-5 control register
type RegisterOtghsdoepctl5Type uint32

func (r *RegisterOtghsdoepctl5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepctl5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepctl5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepctl5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepctl5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepctl5FieldMpsizShift = 0
	RegisterOtghsdoepctl5FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl5Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldMpsizMask) >> RegisterOtghsdoepctl5FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl5Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl5FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl5FieldUsbaepShift = 15
	RegisterOtghsdoepctl5FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl5Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl5Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdoepctl5Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl5FieldNakstsShift = 17
	RegisterOtghsdoepctl5FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdoepctl5Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl5FieldEptypShift = 18
	RegisterOtghsdoepctl5FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdoepctl5Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldEptypMask) >> RegisterOtghsdoepctl5FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdoepctl5Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl5FieldEptypShift))
}

const (
	RegisterOtghsdoepctl5FieldSnpmShift = 20
	RegisterOtghsdoepctl5FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *RegisterOtghsdoepctl5Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *RegisterOtghsdoepctl5Type) SetSnpm(value bool) {
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
func (r *RegisterOtghsdoepctl5Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdoepctl5Type) SetStall(value bool) {
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
func (r *RegisterOtghsdoepctl5Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdoepctl5Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdoepctl5Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdoepctl5Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdoepctl5Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdoepctl5Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdoepctl5Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl5FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdoepctl5Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl5FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl5FieldEpenaMask)
	}
}

// RegisterOtghsdoepint5Type OTG_HS device endpoint-5 interrupt register
type RegisterOtghsdoepint5Type uint32

func (r *RegisterOtghsdoepint5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepint5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepint5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepint5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepint5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepint5FieldXfrcShift = 0
	RegisterOtghsdoepint5FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint5Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint5FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint5Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdoepint5Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint5FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdoepint5Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdoepint5Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint5FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *RegisterOtghsdoepint5Type) SetStup(value bool) {
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
func (r *RegisterOtghsdoepint5Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint5FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *RegisterOtghsdoepint5Type) SetOtepdis(value bool) {
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
func (r *RegisterOtghsdoepint5Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint5FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *RegisterOtghsdoepint5Type) SetB2bstup(value bool) {
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
func (r *RegisterOtghsdoepint5Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint5FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *RegisterOtghsdoepint5Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint5FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint5FieldNyetMask)
	}
}

// RegisterOtghsdoeptsiz5Type OTG_HS device endpoint-5 transfer size register
type RegisterOtghsdoeptsiz5Type uint32

func (r *RegisterOtghsdoeptsiz5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoeptsiz5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoeptsiz5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoeptsiz5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoeptsiz5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoeptsiz5FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz5FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz5Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz5FieldXfrsizMask) >> RegisterOtghsdoeptsiz5FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz5Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz5FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz5FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz5FieldPktcntShift = 19
	RegisterOtghsdoeptsiz5FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz5Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz5FieldPktcntMask) >> RegisterOtghsdoeptsiz5FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz5Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz5FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz5FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz5FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz5FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz5Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz5FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz5FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz5Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz5FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz5FieldRxdpidstupcntShift))
}

// RegisterOtghsdoepctl6Type OTG device endpoint-6 control register
type RegisterOtghsdoepctl6Type uint32

func (r *RegisterOtghsdoepctl6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepctl6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepctl6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepctl6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepctl6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepctl6FieldMpsizShift = 0
	RegisterOtghsdoepctl6FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl6Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldMpsizMask) >> RegisterOtghsdoepctl6FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl6Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl6FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl6FieldUsbaepShift = 15
	RegisterOtghsdoepctl6FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl6Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl6Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdoepctl6Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl6FieldNakstsShift = 17
	RegisterOtghsdoepctl6FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdoepctl6Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl6FieldEptypShift = 18
	RegisterOtghsdoepctl6FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdoepctl6Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldEptypMask) >> RegisterOtghsdoepctl6FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdoepctl6Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl6FieldEptypShift))
}

const (
	RegisterOtghsdoepctl6FieldSnpmShift = 20
	RegisterOtghsdoepctl6FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *RegisterOtghsdoepctl6Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *RegisterOtghsdoepctl6Type) SetSnpm(value bool) {
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
func (r *RegisterOtghsdoepctl6Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdoepctl6Type) SetStall(value bool) {
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
func (r *RegisterOtghsdoepctl6Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdoepctl6Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdoepctl6Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdoepctl6Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdoepctl6Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdoepctl6Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdoepctl6Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl6FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdoepctl6Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl6FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl6FieldEpenaMask)
	}
}

// RegisterOtghsdoepint6Type OTG_HS device endpoint-6 interrupt register
type RegisterOtghsdoepint6Type uint32

func (r *RegisterOtghsdoepint6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepint6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepint6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepint6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepint6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepint6FieldXfrcShift = 0
	RegisterOtghsdoepint6FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint6Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint6FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint6Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdoepint6Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint6FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdoepint6Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdoepint6Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint6FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *RegisterOtghsdoepint6Type) SetStup(value bool) {
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
func (r *RegisterOtghsdoepint6Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint6FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *RegisterOtghsdoepint6Type) SetOtepdis(value bool) {
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
func (r *RegisterOtghsdoepint6Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint6FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *RegisterOtghsdoepint6Type) SetB2bstup(value bool) {
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
func (r *RegisterOtghsdoepint6Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint6FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *RegisterOtghsdoepint6Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint6FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint6FieldNyetMask)
	}
}

// RegisterOtghsdoeptsiz6Type OTG_HS device endpoint-6 transfer size register
type RegisterOtghsdoeptsiz6Type uint32

func (r *RegisterOtghsdoeptsiz6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoeptsiz6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoeptsiz6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoeptsiz6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoeptsiz6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoeptsiz6FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz6FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz6Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz6FieldXfrsizMask) >> RegisterOtghsdoeptsiz6FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz6Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz6FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz6FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz6FieldPktcntShift = 19
	RegisterOtghsdoeptsiz6FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz6Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz6FieldPktcntMask) >> RegisterOtghsdoeptsiz6FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz6Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz6FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz6FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz6FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz6FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz6Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz6FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz6FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz6Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz6FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz6FieldRxdpidstupcntShift))
}

// RegisterOtghsdoepctl7Type OTG device endpoint-7 control register
type RegisterOtghsdoepctl7Type uint32

func (r *RegisterOtghsdoepctl7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepctl7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepctl7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepctl7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepctl7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepctl7FieldMpsizShift = 0
	RegisterOtghsdoepctl7FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl7Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldMpsizMask) >> RegisterOtghsdoepctl7FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghsdoepctl7Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldMpsizMask)|(uint32(value)<<RegisterOtghsdoepctl7FieldMpsizShift))
}

const (
	RegisterOtghsdoepctl7FieldUsbaepShift = 15
	RegisterOtghsdoepctl7FieldUsbaepMask  = 0x8000
)

// GetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl7Type) GetUsbaep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldUsbaepMask) != 0
}

// SetUsbaep USB active endpoint
func (r *RegisterOtghsdoepctl7Type) SetUsbaep(value bool) {
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
func (r *RegisterOtghsdoepctl7Type) GetEonumdpid() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldEonumdpidMask) != 0
}

const (
	RegisterOtghsdoepctl7FieldNakstsShift = 17
	RegisterOtghsdoepctl7FieldNakstsMask  = 0x20000
)

// GetNaksts NAK status
func (r *RegisterOtghsdoepctl7Type) GetNaksts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldNakstsMask) != 0
}

const (
	RegisterOtghsdoepctl7FieldEptypShift = 18
	RegisterOtghsdoepctl7FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *RegisterOtghsdoepctl7Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldEptypMask) >> RegisterOtghsdoepctl7FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghsdoepctl7Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldEptypMask)|(uint32(value)<<RegisterOtghsdoepctl7FieldEptypShift))
}

const (
	RegisterOtghsdoepctl7FieldSnpmShift = 20
	RegisterOtghsdoepctl7FieldSnpmMask  = 0x100000
)

// GetSnpm Snoop mode
func (r *RegisterOtghsdoepctl7Type) GetSnpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldSnpmMask) != 0
}

// SetSnpm Snoop mode
func (r *RegisterOtghsdoepctl7Type) SetSnpm(value bool) {
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
func (r *RegisterOtghsdoepctl7Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldStallMask) != 0
}

// SetStall STALL handshake
func (r *RegisterOtghsdoepctl7Type) SetStall(value bool) {
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
func (r *RegisterOtghsdoepctl7Type) SetCnak(value bool) {
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
func (r *RegisterOtghsdoepctl7Type) SetSnak(value bool) {
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
func (r *RegisterOtghsdoepctl7Type) SetSd0pidsevnfrm(value bool) {
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
func (r *RegisterOtghsdoepctl7Type) SetSoddfrm(value bool) {
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
func (r *RegisterOtghsdoepctl7Type) GetEpdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldEpdisMask) != 0
}

// SetEpdis Endpoint disable
func (r *RegisterOtghsdoepctl7Type) SetEpdis(value bool) {
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
func (r *RegisterOtghsdoepctl7Type) GetEpena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepctl7FieldEpenaMask) != 0
}

// SetEpena Endpoint enable
func (r *RegisterOtghsdoepctl7Type) SetEpena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepctl7FieldEpenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepctl7FieldEpenaMask)
	}
}

// RegisterOtghsdoepint7Type OTG_HS device endpoint-7 interrupt register
type RegisterOtghsdoepint7Type uint32

func (r *RegisterOtghsdoepint7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoepint7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoepint7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoepint7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoepint7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoepint7FieldXfrcShift = 0
	RegisterOtghsdoepint7FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint7Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint7FieldXfrcMask) != 0
}

// SetXfrc Transfer completed interrupt
func (r *RegisterOtghsdoepint7Type) SetXfrc(value bool) {
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
func (r *RegisterOtghsdoepint7Type) GetEpdisd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint7FieldEpdisdMask) != 0
}

// SetEpdisd Endpoint disabled interrupt
func (r *RegisterOtghsdoepint7Type) SetEpdisd(value bool) {
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
func (r *RegisterOtghsdoepint7Type) GetStup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint7FieldStupMask) != 0
}

// SetStup SETUP phase done
func (r *RegisterOtghsdoepint7Type) SetStup(value bool) {
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
func (r *RegisterOtghsdoepint7Type) GetOtepdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint7FieldOtepdisMask) != 0
}

// SetOtepdis OUT token received when endpoint disabled
func (r *RegisterOtghsdoepint7Type) SetOtepdis(value bool) {
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
func (r *RegisterOtghsdoepint7Type) GetB2bstup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint7FieldB2bstupMask) != 0
}

// SetB2bstup Back-to-back SETUP packets received
func (r *RegisterOtghsdoepint7Type) SetB2bstup(value bool) {
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
func (r *RegisterOtghsdoepint7Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoepint7FieldNyetMask) != 0
}

// SetNyet NYET interrupt
func (r *RegisterOtghsdoepint7Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsdoepint7FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoepint7FieldNyetMask)
	}
}

// RegisterOtghsdoeptsiz7Type OTG_HS device endpoint-7 transfer size register
type RegisterOtghsdoeptsiz7Type uint32

func (r *RegisterOtghsdoeptsiz7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghsdoeptsiz7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghsdoeptsiz7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghsdoeptsiz7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghsdoeptsiz7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghsdoeptsiz7FieldXfrsizShift = 0
	RegisterOtghsdoeptsiz7FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz7Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz7FieldXfrsizMask) >> RegisterOtghsdoeptsiz7FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghsdoeptsiz7Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz7FieldXfrsizMask)|(uint32(value)<<RegisterOtghsdoeptsiz7FieldXfrsizShift))
}

const (
	RegisterOtghsdoeptsiz7FieldPktcntShift = 19
	RegisterOtghsdoeptsiz7FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz7Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz7FieldPktcntMask) >> RegisterOtghsdoeptsiz7FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghsdoeptsiz7Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz7FieldPktcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz7FieldPktcntShift))
}

const (
	RegisterOtghsdoeptsiz7FieldRxdpidstupcntShift = 29
	RegisterOtghsdoeptsiz7FieldRxdpidstupcntMask  = 0x60000000
)

// GetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz7Type) GetRxdpidstupcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdoeptsiz7FieldRxdpidstupcntMask) >> RegisterOtghsdoeptsiz7FieldRxdpidstupcntShift)
}

// SetRxdpidstupcnt Received data PID/SETUP packet count
func (r *RegisterOtghsdoeptsiz7Type) SetRxdpidstupcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdoeptsiz7FieldRxdpidstupcntMask)|(uint32(value)<<RegisterOtghsdoeptsiz7FieldRxdpidstupcntShift))
}
