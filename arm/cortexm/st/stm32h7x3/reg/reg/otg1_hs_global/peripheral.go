package otg1_hs_global

import (
	"unsafe"
	"volatile"
)

var (
	Otg1_hs_global = (*_otg1_hs_global)(unsafe.Pointer(uintptr(0x40040000)))
	Otg2_hs_global = (*_otg1_hs_global)(unsafe.Pointer(uintptr(0x40080000)))
)

type _otg1_hs_global struct {
	Otg_hs_gotgctl         registerOtg_hs_gotgctlType
	Otg_hs_gotgint         registerOtg_hs_gotgintType
	Otg_hs_gahbcfg         registerOtg_hs_gahbcfgType
	Otg_hs_gusbcfg         registerOtg_hs_gusbcfgType
	Otg_hs_grstctl         registerOtg_hs_grstctlType
	Otg_hs_gintsts         registerOtg_hs_gintstsType
	Otg_hs_gintmsk         registerOtg_hs_gintmskType
	Otg_hs_grxstsr_host    registerOtg_hs_grxstsr_hostType
	Otg_hs_grxstsp_host    registerOtg_hs_grxstsp_hostType
	Otg_hs_grxfsiz         registerOtg_hs_grxfsizType
	Otg_hs_hnptxfsiz_host  registerOtg_hs_hnptxfsiz_hostType
	Otg_hs_dieptxf0_device registerOtg_hs_dieptxf0_deviceType
	Otg_hs_gnptxsts        registerOtg_hs_gnptxstsType
	_                      [8]byte
	Otg_hs_gccfg           registerOtg_hs_gccfgType
	Otg_hs_cid             registerOtg_hs_cidType
	_                      [192]byte
	Otg_hs_hptxfsiz        registerOtg_hs_hptxfsizType
	Otg_hs_dieptxf1        registerOtg_hs_dieptxf1Type
	Otg_hs_dieptxf2        registerOtg_hs_dieptxf2Type
	_                      [16]byte
	Otg_hs_dieptxf3        registerOtg_hs_dieptxf3Type
	Otg_hs_dieptxf4        registerOtg_hs_dieptxf4Type
	Otg_hs_dieptxf5        registerOtg_hs_dieptxf5Type
	Otg_hs_dieptxf6        registerOtg_hs_dieptxf6Type
	Otg_hs_dieptxf7        registerOtg_hs_dieptxf7Type
	Otg_hs_grxstsr_device  registerOtg_hs_grxstsr_deviceType
	Otg_hs_grxstsp_device  registerOtg_hs_grxstsp_deviceType
	_                      [48]byte
	Otg_hs_glpmcfg         registerOtg_hs_glpmcfgType
}

// registerOtg_hs_gotgctlType OTG_HS control and status register
type registerOtg_hs_gotgctlType uint32

const (
	RegisterOtg_hs_gotgctlFieldSrqscsShift = 0
	RegisterOtg_hs_gotgctlFieldSrqscsMask  = 0x1
)

// GetSrqscs Session request success
func (r *registerOtg_hs_gotgctlType) GetSrqscs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgctlFieldSrqscsMask) != 0
}

const (
	RegisterOtg_hs_gotgctlFieldSrqShift = 1
	RegisterOtg_hs_gotgctlFieldSrqMask  = 0x2
)

// GetSrq Session request
func (r *registerOtg_hs_gotgctlType) GetSrq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgctlFieldSrqMask) != 0
}

// SetSrq Session request
func (r *registerOtg_hs_gotgctlType) SetSrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gotgctlFieldSrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gotgctlFieldSrqMask)
	}
}

const (
	RegisterOtg_hs_gotgctlFieldHngscsShift = 8
	RegisterOtg_hs_gotgctlFieldHngscsMask  = 0x100
)

// GetHngscs Host negotiation success
func (r *registerOtg_hs_gotgctlType) GetHngscs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgctlFieldHngscsMask) != 0
}

const (
	RegisterOtg_hs_gotgctlFieldHnprqShift = 9
	RegisterOtg_hs_gotgctlFieldHnprqMask  = 0x200
)

// GetHnprq HNP request
func (r *registerOtg_hs_gotgctlType) GetHnprq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgctlFieldHnprqMask) != 0
}

// SetHnprq HNP request
func (r *registerOtg_hs_gotgctlType) SetHnprq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gotgctlFieldHnprqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gotgctlFieldHnprqMask)
	}
}

const (
	RegisterOtg_hs_gotgctlFieldHshnpenShift = 10
	RegisterOtg_hs_gotgctlFieldHshnpenMask  = 0x400
)

// GetHshnpen Host set HNP enable
func (r *registerOtg_hs_gotgctlType) GetHshnpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgctlFieldHshnpenMask) != 0
}

// SetHshnpen Host set HNP enable
func (r *registerOtg_hs_gotgctlType) SetHshnpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gotgctlFieldHshnpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gotgctlFieldHshnpenMask)
	}
}

const (
	RegisterOtg_hs_gotgctlFieldDhnpenShift = 11
	RegisterOtg_hs_gotgctlFieldDhnpenMask  = 0x800
)

// GetDhnpen Device HNP enabled
func (r *registerOtg_hs_gotgctlType) GetDhnpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgctlFieldDhnpenMask) != 0
}

// SetDhnpen Device HNP enabled
func (r *registerOtg_hs_gotgctlType) SetDhnpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gotgctlFieldDhnpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gotgctlFieldDhnpenMask)
	}
}

const (
	RegisterOtg_hs_gotgctlFieldCidstsShift = 16
	RegisterOtg_hs_gotgctlFieldCidstsMask  = 0x10000
)

// GetCidsts Connector ID status
func (r *registerOtg_hs_gotgctlType) GetCidsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgctlFieldCidstsMask) != 0
}

const (
	RegisterOtg_hs_gotgctlFieldDbctShift = 17
	RegisterOtg_hs_gotgctlFieldDbctMask  = 0x20000
)

// GetDbct Long/short debounce time
func (r *registerOtg_hs_gotgctlType) GetDbct() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgctlFieldDbctMask) != 0
}

const (
	RegisterOtg_hs_gotgctlFieldAsvldShift = 18
	RegisterOtg_hs_gotgctlFieldAsvldMask  = 0x40000
)

// GetAsvld A-session valid
func (r *registerOtg_hs_gotgctlType) GetAsvld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgctlFieldAsvldMask) != 0
}

const (
	RegisterOtg_hs_gotgctlFieldBsvldShift = 19
	RegisterOtg_hs_gotgctlFieldBsvldMask  = 0x80000
)

// GetBsvld B-session valid
func (r *registerOtg_hs_gotgctlType) GetBsvld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgctlFieldBsvldMask) != 0
}

const (
	RegisterOtg_hs_gotgctlFieldEhenShift = 12
	RegisterOtg_hs_gotgctlFieldEhenMask  = 0x1000
)

// GetEhen Embedded host enable
func (r *registerOtg_hs_gotgctlType) GetEhen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgctlFieldEhenMask) != 0
}

// SetEhen Embedded host enable
func (r *registerOtg_hs_gotgctlType) SetEhen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gotgctlFieldEhenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gotgctlFieldEhenMask)
	}
}

// registerOtg_hs_gotgintType OTG_HS interrupt register
type registerOtg_hs_gotgintType uint32

const (
	RegisterOtg_hs_gotgintFieldSedetShift = 2
	RegisterOtg_hs_gotgintFieldSedetMask  = 0x4
)

// GetSedet Session end detected
func (r *registerOtg_hs_gotgintType) GetSedet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgintFieldSedetMask) != 0
}

// SetSedet Session end detected
func (r *registerOtg_hs_gotgintType) SetSedet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gotgintFieldSedetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gotgintFieldSedetMask)
	}
}

const (
	RegisterOtg_hs_gotgintFieldSrsschgShift = 8
	RegisterOtg_hs_gotgintFieldSrsschgMask  = 0x100
)

// GetSrsschg Session request success status change
func (r *registerOtg_hs_gotgintType) GetSrsschg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgintFieldSrsschgMask) != 0
}

// SetSrsschg Session request success status change
func (r *registerOtg_hs_gotgintType) SetSrsschg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gotgintFieldSrsschgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gotgintFieldSrsschgMask)
	}
}

const (
	RegisterOtg_hs_gotgintFieldHnsschgShift = 9
	RegisterOtg_hs_gotgintFieldHnsschgMask  = 0x200
)

// GetHnsschg Host negotiation success status change
func (r *registerOtg_hs_gotgintType) GetHnsschg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgintFieldHnsschgMask) != 0
}

// SetHnsschg Host negotiation success status change
func (r *registerOtg_hs_gotgintType) SetHnsschg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gotgintFieldHnsschgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gotgintFieldHnsschgMask)
	}
}

const (
	RegisterOtg_hs_gotgintFieldHngdetShift = 17
	RegisterOtg_hs_gotgintFieldHngdetMask  = 0x20000
)

// GetHngdet Host negotiation detected
func (r *registerOtg_hs_gotgintType) GetHngdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgintFieldHngdetMask) != 0
}

// SetHngdet Host negotiation detected
func (r *registerOtg_hs_gotgintType) SetHngdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gotgintFieldHngdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gotgintFieldHngdetMask)
	}
}

const (
	RegisterOtg_hs_gotgintFieldAdtochgShift = 18
	RegisterOtg_hs_gotgintFieldAdtochgMask  = 0x40000
)

// GetAdtochg A-device timeout change
func (r *registerOtg_hs_gotgintType) GetAdtochg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgintFieldAdtochgMask) != 0
}

// SetAdtochg A-device timeout change
func (r *registerOtg_hs_gotgintType) SetAdtochg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gotgintFieldAdtochgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gotgintFieldAdtochgMask)
	}
}

const (
	RegisterOtg_hs_gotgintFieldDbcdneShift = 19
	RegisterOtg_hs_gotgintFieldDbcdneMask  = 0x80000
)

// GetDbcdne Debounce done
func (r *registerOtg_hs_gotgintType) GetDbcdne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgintFieldDbcdneMask) != 0
}

// SetDbcdne Debounce done
func (r *registerOtg_hs_gotgintType) SetDbcdne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gotgintFieldDbcdneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gotgintFieldDbcdneMask)
	}
}

const (
	RegisterOtg_hs_gotgintFieldIdchngShift = 20
	RegisterOtg_hs_gotgintFieldIdchngMask  = 0x100000
)

// GetIdchng ID input pin changed
func (r *registerOtg_hs_gotgintType) GetIdchng() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gotgintFieldIdchngMask) != 0
}

// SetIdchng ID input pin changed
func (r *registerOtg_hs_gotgintType) SetIdchng(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gotgintFieldIdchngMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gotgintFieldIdchngMask)
	}
}

// registerOtg_hs_gahbcfgType OTG_HS AHB configuration register
type registerOtg_hs_gahbcfgType uint32

const (
	RegisterOtg_hs_gahbcfgFieldGintShift = 0
	RegisterOtg_hs_gahbcfgFieldGintMask  = 0x1
)

// GetGint Global interrupt mask
func (r *registerOtg_hs_gahbcfgType) GetGint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gahbcfgFieldGintMask) != 0
}

// SetGint Global interrupt mask
func (r *registerOtg_hs_gahbcfgType) SetGint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gahbcfgFieldGintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gahbcfgFieldGintMask)
	}
}

const (
	RegisterOtg_hs_gahbcfgFieldHbstlenShift = 1
	RegisterOtg_hs_gahbcfgFieldHbstlenMask  = 0x1e
)

// GetHbstlen Burst length/type
func (r *registerOtg_hs_gahbcfgType) GetHbstlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gahbcfgFieldHbstlenMask) >> RegisterOtg_hs_gahbcfgFieldHbstlenShift)
}

// SetHbstlen Burst length/type
func (r *registerOtg_hs_gahbcfgType) SetHbstlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gahbcfgFieldHbstlenMask)|(uint32(value)<<RegisterOtg_hs_gahbcfgFieldHbstlenShift))
}

const (
	RegisterOtg_hs_gahbcfgFieldDmaenShift = 5
	RegisterOtg_hs_gahbcfgFieldDmaenMask  = 0x20
)

// GetDmaen DMA enable
func (r *registerOtg_hs_gahbcfgType) GetDmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gahbcfgFieldDmaenMask) != 0
}

// SetDmaen DMA enable
func (r *registerOtg_hs_gahbcfgType) SetDmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gahbcfgFieldDmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gahbcfgFieldDmaenMask)
	}
}

const (
	RegisterOtg_hs_gahbcfgFieldTxfelvlShift = 7
	RegisterOtg_hs_gahbcfgFieldTxfelvlMask  = 0x80
)

// GetTxfelvl TxFIFO empty level
func (r *registerOtg_hs_gahbcfgType) GetTxfelvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gahbcfgFieldTxfelvlMask) != 0
}

// SetTxfelvl TxFIFO empty level
func (r *registerOtg_hs_gahbcfgType) SetTxfelvl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gahbcfgFieldTxfelvlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gahbcfgFieldTxfelvlMask)
	}
}

const (
	RegisterOtg_hs_gahbcfgFieldPtxfelvlShift = 8
	RegisterOtg_hs_gahbcfgFieldPtxfelvlMask  = 0x100
)

// GetPtxfelvl Periodic TxFIFO empty level
func (r *registerOtg_hs_gahbcfgType) GetPtxfelvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gahbcfgFieldPtxfelvlMask) != 0
}

// SetPtxfelvl Periodic TxFIFO empty level
func (r *registerOtg_hs_gahbcfgType) SetPtxfelvl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gahbcfgFieldPtxfelvlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gahbcfgFieldPtxfelvlMask)
	}
}

// registerOtg_hs_gusbcfgType OTG_HS USB configuration register
type registerOtg_hs_gusbcfgType uint32

const (
	RegisterOtg_hs_gusbcfgFieldTocalShift = 0
	RegisterOtg_hs_gusbcfgFieldTocalMask  = 0x7
)

// GetTocal FS timeout calibration
func (r *registerOtg_hs_gusbcfgType) GetTocal() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldTocalMask) >> RegisterOtg_hs_gusbcfgFieldTocalShift)
}

// SetTocal FS timeout calibration
func (r *registerOtg_hs_gusbcfgType) SetTocal(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldTocalMask)|(uint32(value)<<RegisterOtg_hs_gusbcfgFieldTocalShift))
}

const (
	RegisterOtg_hs_gusbcfgFieldPhyselShift = 6
	RegisterOtg_hs_gusbcfgFieldPhyselMask  = 0x40
)

// SetPhysel USB 2.0 high-speed ULPI PHY or USB 1.1 full-speed serial transceiver select
func (r *registerOtg_hs_gusbcfgType) SetPhysel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldPhyselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldPhyselMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldSrpcapShift = 8
	RegisterOtg_hs_gusbcfgFieldSrpcapMask  = 0x100
)

// GetSrpcap SRP-capable
func (r *registerOtg_hs_gusbcfgType) GetSrpcap() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldSrpcapMask) != 0
}

// SetSrpcap SRP-capable
func (r *registerOtg_hs_gusbcfgType) SetSrpcap(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldSrpcapMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldSrpcapMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldHnpcapShift = 9
	RegisterOtg_hs_gusbcfgFieldHnpcapMask  = 0x200
)

// GetHnpcap HNP-capable
func (r *registerOtg_hs_gusbcfgType) GetHnpcap() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldHnpcapMask) != 0
}

// SetHnpcap HNP-capable
func (r *registerOtg_hs_gusbcfgType) SetHnpcap(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldHnpcapMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldHnpcapMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldTrdtShift = 10
	RegisterOtg_hs_gusbcfgFieldTrdtMask  = 0x3c00
)

// GetTrdt USB turnaround time
func (r *registerOtg_hs_gusbcfgType) GetTrdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldTrdtMask) >> RegisterOtg_hs_gusbcfgFieldTrdtShift)
}

// SetTrdt USB turnaround time
func (r *registerOtg_hs_gusbcfgType) SetTrdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldTrdtMask)|(uint32(value)<<RegisterOtg_hs_gusbcfgFieldTrdtShift))
}

const (
	RegisterOtg_hs_gusbcfgFieldPhylpcsShift = 15
	RegisterOtg_hs_gusbcfgFieldPhylpcsMask  = 0x8000
)

// GetPhylpcs PHY Low-power clock select
func (r *registerOtg_hs_gusbcfgType) GetPhylpcs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldPhylpcsMask) != 0
}

// SetPhylpcs PHY Low-power clock select
func (r *registerOtg_hs_gusbcfgType) SetPhylpcs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldPhylpcsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldPhylpcsMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldUlpifslsShift = 17
	RegisterOtg_hs_gusbcfgFieldUlpifslsMask  = 0x20000
)

// GetUlpifsls ULPI FS/LS select
func (r *registerOtg_hs_gusbcfgType) GetUlpifsls() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldUlpifslsMask) != 0
}

// SetUlpifsls ULPI FS/LS select
func (r *registerOtg_hs_gusbcfgType) SetUlpifsls(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldUlpifslsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldUlpifslsMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldUlpiarShift = 18
	RegisterOtg_hs_gusbcfgFieldUlpiarMask  = 0x40000
)

// GetUlpiar ULPI Auto-resume
func (r *registerOtg_hs_gusbcfgType) GetUlpiar() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldUlpiarMask) != 0
}

// SetUlpiar ULPI Auto-resume
func (r *registerOtg_hs_gusbcfgType) SetUlpiar(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldUlpiarMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldUlpiarMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldUlpicsmShift = 19
	RegisterOtg_hs_gusbcfgFieldUlpicsmMask  = 0x80000
)

// GetUlpicsm ULPI Clock SuspendM
func (r *registerOtg_hs_gusbcfgType) GetUlpicsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldUlpicsmMask) != 0
}

// SetUlpicsm ULPI Clock SuspendM
func (r *registerOtg_hs_gusbcfgType) SetUlpicsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldUlpicsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldUlpicsmMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldUlpievbusdShift = 20
	RegisterOtg_hs_gusbcfgFieldUlpievbusdMask  = 0x100000
)

// GetUlpievbusd ULPI External VBUS Drive
func (r *registerOtg_hs_gusbcfgType) GetUlpievbusd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldUlpievbusdMask) != 0
}

// SetUlpievbusd ULPI External VBUS Drive
func (r *registerOtg_hs_gusbcfgType) SetUlpievbusd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldUlpievbusdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldUlpievbusdMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldUlpievbusiShift = 21
	RegisterOtg_hs_gusbcfgFieldUlpievbusiMask  = 0x200000
)

// GetUlpievbusi ULPI external VBUS indicator
func (r *registerOtg_hs_gusbcfgType) GetUlpievbusi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldUlpievbusiMask) != 0
}

// SetUlpievbusi ULPI external VBUS indicator
func (r *registerOtg_hs_gusbcfgType) SetUlpievbusi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldUlpievbusiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldUlpievbusiMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldTsdpsShift = 22
	RegisterOtg_hs_gusbcfgFieldTsdpsMask  = 0x400000
)

// GetTsdps TermSel DLine pulsing selection
func (r *registerOtg_hs_gusbcfgType) GetTsdps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldTsdpsMask) != 0
}

// SetTsdps TermSel DLine pulsing selection
func (r *registerOtg_hs_gusbcfgType) SetTsdps(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldTsdpsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldTsdpsMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldPcciShift = 23
	RegisterOtg_hs_gusbcfgFieldPcciMask  = 0x800000
)

// GetPcci Indicator complement
func (r *registerOtg_hs_gusbcfgType) GetPcci() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldPcciMask) != 0
}

// SetPcci Indicator complement
func (r *registerOtg_hs_gusbcfgType) SetPcci(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldPcciMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldPcciMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldPtciShift = 24
	RegisterOtg_hs_gusbcfgFieldPtciMask  = 0x1000000
)

// GetPtci Indicator pass through
func (r *registerOtg_hs_gusbcfgType) GetPtci() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldPtciMask) != 0
}

// SetPtci Indicator pass through
func (r *registerOtg_hs_gusbcfgType) SetPtci(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldPtciMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldPtciMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldUlpiipdShift = 25
	RegisterOtg_hs_gusbcfgFieldUlpiipdMask  = 0x2000000
)

// GetUlpiipd ULPI interface protect disable
func (r *registerOtg_hs_gusbcfgType) GetUlpiipd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldUlpiipdMask) != 0
}

// SetUlpiipd ULPI interface protect disable
func (r *registerOtg_hs_gusbcfgType) SetUlpiipd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldUlpiipdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldUlpiipdMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldFhmodShift = 29
	RegisterOtg_hs_gusbcfgFieldFhmodMask  = 0x20000000
)

// GetFhmod Forced host mode
func (r *registerOtg_hs_gusbcfgType) GetFhmod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldFhmodMask) != 0
}

// SetFhmod Forced host mode
func (r *registerOtg_hs_gusbcfgType) SetFhmod(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldFhmodMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldFhmodMask)
	}
}

const (
	RegisterOtg_hs_gusbcfgFieldFdmodShift = 30
	RegisterOtg_hs_gusbcfgFieldFdmodMask  = 0x40000000
)

// GetFdmod Forced peripheral mode
func (r *registerOtg_hs_gusbcfgType) GetFdmod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gusbcfgFieldFdmodMask) != 0
}

// SetFdmod Forced peripheral mode
func (r *registerOtg_hs_gusbcfgType) SetFdmod(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gusbcfgFieldFdmodMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gusbcfgFieldFdmodMask)
	}
}

// registerOtg_hs_grstctlType OTG_HS reset register
type registerOtg_hs_grstctlType uint32

const (
	RegisterOtg_hs_grstctlFieldCsrstShift = 0
	RegisterOtg_hs_grstctlFieldCsrstMask  = 0x1
)

// GetCsrst Core soft reset
func (r *registerOtg_hs_grstctlType) GetCsrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grstctlFieldCsrstMask) != 0
}

// SetCsrst Core soft reset
func (r *registerOtg_hs_grstctlType) SetCsrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_grstctlFieldCsrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grstctlFieldCsrstMask)
	}
}

const (
	RegisterOtg_hs_grstctlFieldHsrstShift = 1
	RegisterOtg_hs_grstctlFieldHsrstMask  = 0x2
)

// GetHsrst HCLK soft reset
func (r *registerOtg_hs_grstctlType) GetHsrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grstctlFieldHsrstMask) != 0
}

// SetHsrst HCLK soft reset
func (r *registerOtg_hs_grstctlType) SetHsrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_grstctlFieldHsrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grstctlFieldHsrstMask)
	}
}

const (
	RegisterOtg_hs_grstctlFieldFcrstShift = 2
	RegisterOtg_hs_grstctlFieldFcrstMask  = 0x4
)

// GetFcrst Host frame counter reset
func (r *registerOtg_hs_grstctlType) GetFcrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grstctlFieldFcrstMask) != 0
}

// SetFcrst Host frame counter reset
func (r *registerOtg_hs_grstctlType) SetFcrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_grstctlFieldFcrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grstctlFieldFcrstMask)
	}
}

const (
	RegisterOtg_hs_grstctlFieldRxfflshShift = 4
	RegisterOtg_hs_grstctlFieldRxfflshMask  = 0x10
)

// GetRxfflsh RxFIFO flush
func (r *registerOtg_hs_grstctlType) GetRxfflsh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grstctlFieldRxfflshMask) != 0
}

// SetRxfflsh RxFIFO flush
func (r *registerOtg_hs_grstctlType) SetRxfflsh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_grstctlFieldRxfflshMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grstctlFieldRxfflshMask)
	}
}

const (
	RegisterOtg_hs_grstctlFieldTxfflshShift = 5
	RegisterOtg_hs_grstctlFieldTxfflshMask  = 0x20
)

// GetTxfflsh TxFIFO flush
func (r *registerOtg_hs_grstctlType) GetTxfflsh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grstctlFieldTxfflshMask) != 0
}

// SetTxfflsh TxFIFO flush
func (r *registerOtg_hs_grstctlType) SetTxfflsh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_grstctlFieldTxfflshMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grstctlFieldTxfflshMask)
	}
}

const (
	RegisterOtg_hs_grstctlFieldTxfnumShift = 6
	RegisterOtg_hs_grstctlFieldTxfnumMask  = 0x7c0
)

// GetTxfnum TxFIFO number
func (r *registerOtg_hs_grstctlType) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grstctlFieldTxfnumMask) >> RegisterOtg_hs_grstctlFieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtg_hs_grstctlType) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grstctlFieldTxfnumMask)|(uint32(value)<<RegisterOtg_hs_grstctlFieldTxfnumShift))
}

const (
	RegisterOtg_hs_grstctlFieldAhbidlShift = 31
	RegisterOtg_hs_grstctlFieldAhbidlMask  = 0x80000000
)

// GetAhbidl AHB master idle
func (r *registerOtg_hs_grstctlType) GetAhbidl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grstctlFieldAhbidlMask) != 0
}

const (
	RegisterOtg_hs_grstctlFieldDmareqShift = 30
	RegisterOtg_hs_grstctlFieldDmareqMask  = 0x40000000
)

// GetDmareq DMA request signal enabled for USB OTG HS
func (r *registerOtg_hs_grstctlType) GetDmareq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grstctlFieldDmareqMask) != 0
}

// registerOtg_hs_gintstsType OTG_HS core interrupt register
type registerOtg_hs_gintstsType uint32

const (
	RegisterOtg_hs_gintstsFieldCmodShift = 0
	RegisterOtg_hs_gintstsFieldCmodMask  = 0x1
)

// GetCmod Current mode of operation
func (r *registerOtg_hs_gintstsType) GetCmod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldCmodMask) != 0
}

const (
	RegisterOtg_hs_gintstsFieldMmisShift = 1
	RegisterOtg_hs_gintstsFieldMmisMask  = 0x2
)

// GetMmis Mode mismatch interrupt
func (r *registerOtg_hs_gintstsType) GetMmis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldMmisMask) != 0
}

// SetMmis Mode mismatch interrupt
func (r *registerOtg_hs_gintstsType) SetMmis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldMmisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldMmisMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldOtgintShift = 2
	RegisterOtg_hs_gintstsFieldOtgintMask  = 0x4
)

// GetOtgint OTG interrupt
func (r *registerOtg_hs_gintstsType) GetOtgint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldOtgintMask) != 0
}

const (
	RegisterOtg_hs_gintstsFieldSofShift = 3
	RegisterOtg_hs_gintstsFieldSofMask  = 0x8
)

// GetSof Start of frame
func (r *registerOtg_hs_gintstsType) GetSof() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldSofMask) != 0
}

// SetSof Start of frame
func (r *registerOtg_hs_gintstsType) SetSof(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldSofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldSofMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldRxflvlShift = 4
	RegisterOtg_hs_gintstsFieldRxflvlMask  = 0x10
)

// GetRxflvl RxFIFO nonempty
func (r *registerOtg_hs_gintstsType) GetRxflvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldRxflvlMask) != 0
}

const (
	RegisterOtg_hs_gintstsFieldNptxfeShift = 5
	RegisterOtg_hs_gintstsFieldNptxfeMask  = 0x20
)

// GetNptxfe Nonperiodic TxFIFO empty
func (r *registerOtg_hs_gintstsType) GetNptxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldNptxfeMask) != 0
}

const (
	RegisterOtg_hs_gintstsFieldGinakeffShift = 6
	RegisterOtg_hs_gintstsFieldGinakeffMask  = 0x40
)

// GetGinakeff Global IN nonperiodic NAK effective
func (r *registerOtg_hs_gintstsType) GetGinakeff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldGinakeffMask) != 0
}

const (
	RegisterOtg_hs_gintstsFieldBoutnakeffShift = 7
	RegisterOtg_hs_gintstsFieldBoutnakeffMask  = 0x80
)

// GetBoutnakeff Global OUT NAK effective
func (r *registerOtg_hs_gintstsType) GetBoutnakeff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldBoutnakeffMask) != 0
}

const (
	RegisterOtg_hs_gintstsFieldEsuspShift = 10
	RegisterOtg_hs_gintstsFieldEsuspMask  = 0x400
)

// GetEsusp Early suspend
func (r *registerOtg_hs_gintstsType) GetEsusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldEsuspMask) != 0
}

// SetEsusp Early suspend
func (r *registerOtg_hs_gintstsType) SetEsusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldEsuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldEsuspMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldUsbsuspShift = 11
	RegisterOtg_hs_gintstsFieldUsbsuspMask  = 0x800
)

// GetUsbsusp USB suspend
func (r *registerOtg_hs_gintstsType) GetUsbsusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldUsbsuspMask) != 0
}

// SetUsbsusp USB suspend
func (r *registerOtg_hs_gintstsType) SetUsbsusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldUsbsuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldUsbsuspMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldUsbrstShift = 12
	RegisterOtg_hs_gintstsFieldUsbrstMask  = 0x1000
)

// GetUsbrst USB reset
func (r *registerOtg_hs_gintstsType) GetUsbrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldUsbrstMask) != 0
}

// SetUsbrst USB reset
func (r *registerOtg_hs_gintstsType) SetUsbrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldUsbrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldUsbrstMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldEnumdneShift = 13
	RegisterOtg_hs_gintstsFieldEnumdneMask  = 0x2000
)

// GetEnumdne Enumeration done
func (r *registerOtg_hs_gintstsType) GetEnumdne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldEnumdneMask) != 0
}

// SetEnumdne Enumeration done
func (r *registerOtg_hs_gintstsType) SetEnumdne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldEnumdneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldEnumdneMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldIsoodrpShift = 14
	RegisterOtg_hs_gintstsFieldIsoodrpMask  = 0x4000
)

// GetIsoodrp Isochronous OUT packet dropped interrupt
func (r *registerOtg_hs_gintstsType) GetIsoodrp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldIsoodrpMask) != 0
}

// SetIsoodrp Isochronous OUT packet dropped interrupt
func (r *registerOtg_hs_gintstsType) SetIsoodrp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldIsoodrpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldIsoodrpMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldEopfShift = 15
	RegisterOtg_hs_gintstsFieldEopfMask  = 0x8000
)

// GetEopf End of periodic frame interrupt
func (r *registerOtg_hs_gintstsType) GetEopf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldEopfMask) != 0
}

// SetEopf End of periodic frame interrupt
func (r *registerOtg_hs_gintstsType) SetEopf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldEopfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldEopfMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldIepintShift = 18
	RegisterOtg_hs_gintstsFieldIepintMask  = 0x40000
)

// GetIepint IN endpoint interrupt
func (r *registerOtg_hs_gintstsType) GetIepint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldIepintMask) != 0
}

const (
	RegisterOtg_hs_gintstsFieldOepintShift = 19
	RegisterOtg_hs_gintstsFieldOepintMask  = 0x80000
)

// GetOepint OUT endpoint interrupt
func (r *registerOtg_hs_gintstsType) GetOepint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldOepintMask) != 0
}

const (
	RegisterOtg_hs_gintstsFieldIisoixfrShift = 20
	RegisterOtg_hs_gintstsFieldIisoixfrMask  = 0x100000
)

// GetIisoixfr Incomplete isochronous IN transfer
func (r *registerOtg_hs_gintstsType) GetIisoixfr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldIisoixfrMask) != 0
}

// SetIisoixfr Incomplete isochronous IN transfer
func (r *registerOtg_hs_gintstsType) SetIisoixfr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldIisoixfrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldIisoixfrMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldPxfr_incompisooutShift = 21
	RegisterOtg_hs_gintstsFieldPxfr_incompisooutMask  = 0x200000
)

// GetPxfr_incompisoout Incomplete periodic transfer
func (r *registerOtg_hs_gintstsType) GetPxfr_incompisoout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldPxfr_incompisooutMask) != 0
}

// SetPxfr_incompisoout Incomplete periodic transfer
func (r *registerOtg_hs_gintstsType) SetPxfr_incompisoout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldPxfr_incompisooutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldPxfr_incompisooutMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldDatafsuspShift = 22
	RegisterOtg_hs_gintstsFieldDatafsuspMask  = 0x400000
)

// GetDatafsusp Data fetch suspended
func (r *registerOtg_hs_gintstsType) GetDatafsusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldDatafsuspMask) != 0
}

// SetDatafsusp Data fetch suspended
func (r *registerOtg_hs_gintstsType) SetDatafsusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldDatafsuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldDatafsuspMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldHprtintShift = 24
	RegisterOtg_hs_gintstsFieldHprtintMask  = 0x1000000
)

// GetHprtint Host port interrupt
func (r *registerOtg_hs_gintstsType) GetHprtint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldHprtintMask) != 0
}

const (
	RegisterOtg_hs_gintstsFieldHcintShift = 25
	RegisterOtg_hs_gintstsFieldHcintMask  = 0x2000000
)

// GetHcint Host channels interrupt
func (r *registerOtg_hs_gintstsType) GetHcint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldHcintMask) != 0
}

const (
	RegisterOtg_hs_gintstsFieldPtxfeShift = 26
	RegisterOtg_hs_gintstsFieldPtxfeMask  = 0x4000000
)

// GetPtxfe Periodic TxFIFO empty
func (r *registerOtg_hs_gintstsType) GetPtxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldPtxfeMask) != 0
}

const (
	RegisterOtg_hs_gintstsFieldCidschgShift = 28
	RegisterOtg_hs_gintstsFieldCidschgMask  = 0x10000000
)

// GetCidschg Connector ID status change
func (r *registerOtg_hs_gintstsType) GetCidschg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldCidschgMask) != 0
}

// SetCidschg Connector ID status change
func (r *registerOtg_hs_gintstsType) SetCidschg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldCidschgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldCidschgMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldDiscintShift = 29
	RegisterOtg_hs_gintstsFieldDiscintMask  = 0x20000000
)

// GetDiscint Disconnect detected interrupt
func (r *registerOtg_hs_gintstsType) GetDiscint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldDiscintMask) != 0
}

// SetDiscint Disconnect detected interrupt
func (r *registerOtg_hs_gintstsType) SetDiscint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldDiscintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldDiscintMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldSrqintShift = 30
	RegisterOtg_hs_gintstsFieldSrqintMask  = 0x40000000
)

// GetSrqint Session request/new session detected interrupt
func (r *registerOtg_hs_gintstsType) GetSrqint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldSrqintMask) != 0
}

// SetSrqint Session request/new session detected interrupt
func (r *registerOtg_hs_gintstsType) SetSrqint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldSrqintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldSrqintMask)
	}
}

const (
	RegisterOtg_hs_gintstsFieldWkuintShift = 31
	RegisterOtg_hs_gintstsFieldWkuintMask  = 0x80000000
)

// GetWkuint Resume/remote wakeup detected interrupt
func (r *registerOtg_hs_gintstsType) GetWkuint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintstsFieldWkuintMask) != 0
}

// SetWkuint Resume/remote wakeup detected interrupt
func (r *registerOtg_hs_gintstsType) SetWkuint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintstsFieldWkuintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintstsFieldWkuintMask)
	}
}

// registerOtg_hs_gintmskType OTG_HS interrupt mask register
type registerOtg_hs_gintmskType uint32

const (
	RegisterOtg_hs_gintmskFieldMmismShift = 1
	RegisterOtg_hs_gintmskFieldMmismMask  = 0x2
)

// GetMmism Mode mismatch interrupt mask
func (r *registerOtg_hs_gintmskType) GetMmism() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldMmismMask) != 0
}

// SetMmism Mode mismatch interrupt mask
func (r *registerOtg_hs_gintmskType) SetMmism(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldMmismMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldMmismMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldOtgintShift = 2
	RegisterOtg_hs_gintmskFieldOtgintMask  = 0x4
)

// GetOtgint OTG interrupt mask
func (r *registerOtg_hs_gintmskType) GetOtgint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldOtgintMask) != 0
}

// SetOtgint OTG interrupt mask
func (r *registerOtg_hs_gintmskType) SetOtgint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldOtgintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldOtgintMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldSofmShift = 3
	RegisterOtg_hs_gintmskFieldSofmMask  = 0x8
)

// GetSofm Start of frame mask
func (r *registerOtg_hs_gintmskType) GetSofm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldSofmMask) != 0
}

// SetSofm Start of frame mask
func (r *registerOtg_hs_gintmskType) SetSofm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldSofmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldSofmMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldRxflvlmShift = 4
	RegisterOtg_hs_gintmskFieldRxflvlmMask  = 0x10
)

// GetRxflvlm Receive FIFO nonempty mask
func (r *registerOtg_hs_gintmskType) GetRxflvlm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldRxflvlmMask) != 0
}

// SetRxflvlm Receive FIFO nonempty mask
func (r *registerOtg_hs_gintmskType) SetRxflvlm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldRxflvlmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldRxflvlmMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldNptxfemShift = 5
	RegisterOtg_hs_gintmskFieldNptxfemMask  = 0x20
)

// GetNptxfem Nonperiodic TxFIFO empty mask
func (r *registerOtg_hs_gintmskType) GetNptxfem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldNptxfemMask) != 0
}

// SetNptxfem Nonperiodic TxFIFO empty mask
func (r *registerOtg_hs_gintmskType) SetNptxfem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldNptxfemMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldNptxfemMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldGinakeffmShift = 6
	RegisterOtg_hs_gintmskFieldGinakeffmMask  = 0x40
)

// GetGinakeffm Global nonperiodic IN NAK effective mask
func (r *registerOtg_hs_gintmskType) GetGinakeffm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldGinakeffmMask) != 0
}

// SetGinakeffm Global nonperiodic IN NAK effective mask
func (r *registerOtg_hs_gintmskType) SetGinakeffm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldGinakeffmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldGinakeffmMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldGonakeffmShift = 7
	RegisterOtg_hs_gintmskFieldGonakeffmMask  = 0x80
)

// GetGonakeffm Global OUT NAK effective mask
func (r *registerOtg_hs_gintmskType) GetGonakeffm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldGonakeffmMask) != 0
}

// SetGonakeffm Global OUT NAK effective mask
func (r *registerOtg_hs_gintmskType) SetGonakeffm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldGonakeffmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldGonakeffmMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldEsuspmShift = 10
	RegisterOtg_hs_gintmskFieldEsuspmMask  = 0x400
)

// GetEsuspm Early suspend mask
func (r *registerOtg_hs_gintmskType) GetEsuspm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldEsuspmMask) != 0
}

// SetEsuspm Early suspend mask
func (r *registerOtg_hs_gintmskType) SetEsuspm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldEsuspmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldEsuspmMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldUsbsuspmShift = 11
	RegisterOtg_hs_gintmskFieldUsbsuspmMask  = 0x800
)

// GetUsbsuspm USB suspend mask
func (r *registerOtg_hs_gintmskType) GetUsbsuspm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldUsbsuspmMask) != 0
}

// SetUsbsuspm USB suspend mask
func (r *registerOtg_hs_gintmskType) SetUsbsuspm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldUsbsuspmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldUsbsuspmMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldUsbrstShift = 12
	RegisterOtg_hs_gintmskFieldUsbrstMask  = 0x1000
)

// GetUsbrst USB reset mask
func (r *registerOtg_hs_gintmskType) GetUsbrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldUsbrstMask) != 0
}

// SetUsbrst USB reset mask
func (r *registerOtg_hs_gintmskType) SetUsbrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldUsbrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldUsbrstMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldEnumdnemShift = 13
	RegisterOtg_hs_gintmskFieldEnumdnemMask  = 0x2000
)

// GetEnumdnem Enumeration done mask
func (r *registerOtg_hs_gintmskType) GetEnumdnem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldEnumdnemMask) != 0
}

// SetEnumdnem Enumeration done mask
func (r *registerOtg_hs_gintmskType) SetEnumdnem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldEnumdnemMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldEnumdnemMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldIsoodrpmShift = 14
	RegisterOtg_hs_gintmskFieldIsoodrpmMask  = 0x4000
)

// GetIsoodrpm Isochronous OUT packet dropped interrupt mask
func (r *registerOtg_hs_gintmskType) GetIsoodrpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldIsoodrpmMask) != 0
}

// SetIsoodrpm Isochronous OUT packet dropped interrupt mask
func (r *registerOtg_hs_gintmskType) SetIsoodrpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldIsoodrpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldIsoodrpmMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldEopfmShift = 15
	RegisterOtg_hs_gintmskFieldEopfmMask  = 0x8000
)

// GetEopfm End of periodic frame interrupt mask
func (r *registerOtg_hs_gintmskType) GetEopfm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldEopfmMask) != 0
}

// SetEopfm End of periodic frame interrupt mask
func (r *registerOtg_hs_gintmskType) SetEopfm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldEopfmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldEopfmMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldIepintShift = 18
	RegisterOtg_hs_gintmskFieldIepintMask  = 0x40000
)

// GetIepint IN endpoints interrupt mask
func (r *registerOtg_hs_gintmskType) GetIepint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldIepintMask) != 0
}

// SetIepint IN endpoints interrupt mask
func (r *registerOtg_hs_gintmskType) SetIepint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldIepintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldIepintMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldOepintShift = 19
	RegisterOtg_hs_gintmskFieldOepintMask  = 0x80000
)

// GetOepint OUT endpoints interrupt mask
func (r *registerOtg_hs_gintmskType) GetOepint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldOepintMask) != 0
}

// SetOepint OUT endpoints interrupt mask
func (r *registerOtg_hs_gintmskType) SetOepint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldOepintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldOepintMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldIisoixfrmShift = 20
	RegisterOtg_hs_gintmskFieldIisoixfrmMask  = 0x100000
)

// GetIisoixfrm Incomplete isochronous IN transfer mask
func (r *registerOtg_hs_gintmskType) GetIisoixfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldIisoixfrmMask) != 0
}

// SetIisoixfrm Incomplete isochronous IN transfer mask
func (r *registerOtg_hs_gintmskType) SetIisoixfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldIisoixfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldIisoixfrmMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldPxfrm_iisooxfrmShift = 21
	RegisterOtg_hs_gintmskFieldPxfrm_iisooxfrmMask  = 0x200000
)

// GetPxfrm_iisooxfrm Incomplete periodic transfer mask
func (r *registerOtg_hs_gintmskType) GetPxfrm_iisooxfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldPxfrm_iisooxfrmMask) != 0
}

// SetPxfrm_iisooxfrm Incomplete periodic transfer mask
func (r *registerOtg_hs_gintmskType) SetPxfrm_iisooxfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldPxfrm_iisooxfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldPxfrm_iisooxfrmMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldFsuspmShift = 22
	RegisterOtg_hs_gintmskFieldFsuspmMask  = 0x400000
)

// GetFsuspm Data fetch suspended mask
func (r *registerOtg_hs_gintmskType) GetFsuspm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldFsuspmMask) != 0
}

// SetFsuspm Data fetch suspended mask
func (r *registerOtg_hs_gintmskType) SetFsuspm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldFsuspmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldFsuspmMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldPrtimShift = 24
	RegisterOtg_hs_gintmskFieldPrtimMask  = 0x1000000
)

// GetPrtim Host port interrupt mask
func (r *registerOtg_hs_gintmskType) GetPrtim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldPrtimMask) != 0
}

const (
	RegisterOtg_hs_gintmskFieldHcimShift = 25
	RegisterOtg_hs_gintmskFieldHcimMask  = 0x2000000
)

// GetHcim Host channels interrupt mask
func (r *registerOtg_hs_gintmskType) GetHcim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldHcimMask) != 0
}

// SetHcim Host channels interrupt mask
func (r *registerOtg_hs_gintmskType) SetHcim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldHcimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldHcimMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldPtxfemShift = 26
	RegisterOtg_hs_gintmskFieldPtxfemMask  = 0x4000000
)

// GetPtxfem Periodic TxFIFO empty mask
func (r *registerOtg_hs_gintmskType) GetPtxfem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldPtxfemMask) != 0
}

// SetPtxfem Periodic TxFIFO empty mask
func (r *registerOtg_hs_gintmskType) SetPtxfem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldPtxfemMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldPtxfemMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldCidschgmShift = 28
	RegisterOtg_hs_gintmskFieldCidschgmMask  = 0x10000000
)

// GetCidschgm Connector ID status change mask
func (r *registerOtg_hs_gintmskType) GetCidschgm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldCidschgmMask) != 0
}

// SetCidschgm Connector ID status change mask
func (r *registerOtg_hs_gintmskType) SetCidschgm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldCidschgmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldCidschgmMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldDiscintShift = 29
	RegisterOtg_hs_gintmskFieldDiscintMask  = 0x20000000
)

// GetDiscint Disconnect detected interrupt mask
func (r *registerOtg_hs_gintmskType) GetDiscint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldDiscintMask) != 0
}

// SetDiscint Disconnect detected interrupt mask
func (r *registerOtg_hs_gintmskType) SetDiscint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldDiscintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldDiscintMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldSrqimShift = 30
	RegisterOtg_hs_gintmskFieldSrqimMask  = 0x40000000
)

// GetSrqim Session request/new session detected interrupt mask
func (r *registerOtg_hs_gintmskType) GetSrqim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldSrqimMask) != 0
}

// SetSrqim Session request/new session detected interrupt mask
func (r *registerOtg_hs_gintmskType) SetSrqim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldSrqimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldSrqimMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldWuimShift = 31
	RegisterOtg_hs_gintmskFieldWuimMask  = 0x80000000
)

// GetWuim Resume/remote wakeup detected interrupt mask
func (r *registerOtg_hs_gintmskType) GetWuim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldWuimMask) != 0
}

// SetWuim Resume/remote wakeup detected interrupt mask
func (r *registerOtg_hs_gintmskType) SetWuim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldWuimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldWuimMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldRstdeShift = 23
	RegisterOtg_hs_gintmskFieldRstdeMask  = 0x800000
)

// GetRstde Reset detected interrupt mask
func (r *registerOtg_hs_gintmskType) GetRstde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldRstdeMask) != 0
}

// SetRstde Reset detected interrupt mask
func (r *registerOtg_hs_gintmskType) SetRstde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldRstdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldRstdeMask)
	}
}

const (
	RegisterOtg_hs_gintmskFieldLpmintmShift = 27
	RegisterOtg_hs_gintmskFieldLpmintmMask  = 0x8000000
)

// GetLpmintm LPM interrupt mask
func (r *registerOtg_hs_gintmskType) GetLpmintm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gintmskFieldLpmintmMask) != 0
}

// SetLpmintm LPM interrupt mask
func (r *registerOtg_hs_gintmskType) SetLpmintm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gintmskFieldLpmintmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gintmskFieldLpmintmMask)
	}
}

// registerOtg_hs_grxstsr_hostType OTG_HS Receive status debug read register (host mode)
type registerOtg_hs_grxstsr_hostType uint32

const (
	RegisterOtg_hs_grxstsr_hostFieldChnumShift = 0
	RegisterOtg_hs_grxstsr_hostFieldChnumMask  = 0xf
)

// GetChnum Channel number
func (r *registerOtg_hs_grxstsr_hostType) GetChnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsr_hostFieldChnumMask) >> RegisterOtg_hs_grxstsr_hostFieldChnumShift)
}

// SetChnum Channel number
func (r *registerOtg_hs_grxstsr_hostType) SetChnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsr_hostFieldChnumMask)|(uint32(value)<<RegisterOtg_hs_grxstsr_hostFieldChnumShift))
}

const (
	RegisterOtg_hs_grxstsr_hostFieldBcntShift = 4
	RegisterOtg_hs_grxstsr_hostFieldBcntMask  = 0x7ff0
)

// GetBcnt Byte count
func (r *registerOtg_hs_grxstsr_hostType) GetBcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsr_hostFieldBcntMask) >> RegisterOtg_hs_grxstsr_hostFieldBcntShift)
}

// SetBcnt Byte count
func (r *registerOtg_hs_grxstsr_hostType) SetBcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsr_hostFieldBcntMask)|(uint32(value)<<RegisterOtg_hs_grxstsr_hostFieldBcntShift))
}

const (
	RegisterOtg_hs_grxstsr_hostFieldDpidShift = 15
	RegisterOtg_hs_grxstsr_hostFieldDpidMask  = 0x18000
)

// GetDpid Data PID
func (r *registerOtg_hs_grxstsr_hostType) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsr_hostFieldDpidMask) >> RegisterOtg_hs_grxstsr_hostFieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_grxstsr_hostType) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsr_hostFieldDpidMask)|(uint32(value)<<RegisterOtg_hs_grxstsr_hostFieldDpidShift))
}

const (
	RegisterOtg_hs_grxstsr_hostFieldPktstsShift = 17
	RegisterOtg_hs_grxstsr_hostFieldPktstsMask  = 0x1e0000
)

// GetPktsts Packet status
func (r *registerOtg_hs_grxstsr_hostType) GetPktsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsr_hostFieldPktstsMask) >> RegisterOtg_hs_grxstsr_hostFieldPktstsShift)
}

// SetPktsts Packet status
func (r *registerOtg_hs_grxstsr_hostType) SetPktsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsr_hostFieldPktstsMask)|(uint32(value)<<RegisterOtg_hs_grxstsr_hostFieldPktstsShift))
}

// registerOtg_hs_grxstsp_hostType OTG_HS status read and pop register (host mode)
type registerOtg_hs_grxstsp_hostType uint32

const (
	RegisterOtg_hs_grxstsp_hostFieldChnumShift = 0
	RegisterOtg_hs_grxstsp_hostFieldChnumMask  = 0xf
)

// GetChnum Channel number
func (r *registerOtg_hs_grxstsp_hostType) GetChnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsp_hostFieldChnumMask) >> RegisterOtg_hs_grxstsp_hostFieldChnumShift)
}

// SetChnum Channel number
func (r *registerOtg_hs_grxstsp_hostType) SetChnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsp_hostFieldChnumMask)|(uint32(value)<<RegisterOtg_hs_grxstsp_hostFieldChnumShift))
}

const (
	RegisterOtg_hs_grxstsp_hostFieldBcntShift = 4
	RegisterOtg_hs_grxstsp_hostFieldBcntMask  = 0x7ff0
)

// GetBcnt Byte count
func (r *registerOtg_hs_grxstsp_hostType) GetBcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsp_hostFieldBcntMask) >> RegisterOtg_hs_grxstsp_hostFieldBcntShift)
}

// SetBcnt Byte count
func (r *registerOtg_hs_grxstsp_hostType) SetBcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsp_hostFieldBcntMask)|(uint32(value)<<RegisterOtg_hs_grxstsp_hostFieldBcntShift))
}

const (
	RegisterOtg_hs_grxstsp_hostFieldDpidShift = 15
	RegisterOtg_hs_grxstsp_hostFieldDpidMask  = 0x18000
)

// GetDpid Data PID
func (r *registerOtg_hs_grxstsp_hostType) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsp_hostFieldDpidMask) >> RegisterOtg_hs_grxstsp_hostFieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_grxstsp_hostType) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsp_hostFieldDpidMask)|(uint32(value)<<RegisterOtg_hs_grxstsp_hostFieldDpidShift))
}

const (
	RegisterOtg_hs_grxstsp_hostFieldPktstsShift = 17
	RegisterOtg_hs_grxstsp_hostFieldPktstsMask  = 0x1e0000
)

// GetPktsts Packet status
func (r *registerOtg_hs_grxstsp_hostType) GetPktsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsp_hostFieldPktstsMask) >> RegisterOtg_hs_grxstsp_hostFieldPktstsShift)
}

// SetPktsts Packet status
func (r *registerOtg_hs_grxstsp_hostType) SetPktsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsp_hostFieldPktstsMask)|(uint32(value)<<RegisterOtg_hs_grxstsp_hostFieldPktstsShift))
}

// registerOtg_hs_grxfsizType OTG_HS Receive FIFO size register
type registerOtg_hs_grxfsizType uint32

const (
	RegisterOtg_hs_grxfsizFieldRxfdShift = 0
	RegisterOtg_hs_grxfsizFieldRxfdMask  = 0xffff
)

// GetRxfd RxFIFO depth
func (r *registerOtg_hs_grxfsizType) GetRxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxfsizFieldRxfdMask) >> RegisterOtg_hs_grxfsizFieldRxfdShift)
}

// SetRxfd RxFIFO depth
func (r *registerOtg_hs_grxfsizType) SetRxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxfsizFieldRxfdMask)|(uint32(value)<<RegisterOtg_hs_grxfsizFieldRxfdShift))
}

// registerOtg_hs_hnptxfsiz_hostType OTG_HS nonperiodic transmit FIFO size register (host mode)
type registerOtg_hs_hnptxfsiz_hostType uint32

const (
	RegisterOtg_hs_hnptxfsiz_hostFieldNptxfsaShift = 0
	RegisterOtg_hs_hnptxfsiz_hostFieldNptxfsaMask  = 0xffff
)

// GetNptxfsa Nonperiodic transmit RAM start address
func (r *registerOtg_hs_hnptxfsiz_hostType) GetNptxfsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hnptxfsiz_hostFieldNptxfsaMask) >> RegisterOtg_hs_hnptxfsiz_hostFieldNptxfsaShift)
}

// SetNptxfsa Nonperiodic transmit RAM start address
func (r *registerOtg_hs_hnptxfsiz_hostType) SetNptxfsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hnptxfsiz_hostFieldNptxfsaMask)|(uint32(value)<<RegisterOtg_hs_hnptxfsiz_hostFieldNptxfsaShift))
}

const (
	RegisterOtg_hs_hnptxfsiz_hostFieldNptxfdShift = 16
	RegisterOtg_hs_hnptxfsiz_hostFieldNptxfdMask  = 0xffff0000
)

// GetNptxfd Nonperiodic TxFIFO depth
func (r *registerOtg_hs_hnptxfsiz_hostType) GetNptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hnptxfsiz_hostFieldNptxfdMask) >> RegisterOtg_hs_hnptxfsiz_hostFieldNptxfdShift)
}

// SetNptxfd Nonperiodic TxFIFO depth
func (r *registerOtg_hs_hnptxfsiz_hostType) SetNptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hnptxfsiz_hostFieldNptxfdMask)|(uint32(value)<<RegisterOtg_hs_hnptxfsiz_hostFieldNptxfdShift))
}

// registerOtg_hs_dieptxf0_deviceType Endpoint 0 transmit FIFO size (peripheral mode)
type registerOtg_hs_dieptxf0_deviceType uint32

const (
	RegisterOtg_hs_dieptxf0_deviceFieldTx0fsaShift = 0
	RegisterOtg_hs_dieptxf0_deviceFieldTx0fsaMask  = 0xffff
)

// GetTx0fsa Endpoint 0 transmit RAM start address
func (r *registerOtg_hs_dieptxf0_deviceType) GetTx0fsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf0_deviceFieldTx0fsaMask) >> RegisterOtg_hs_dieptxf0_deviceFieldTx0fsaShift)
}

// SetTx0fsa Endpoint 0 transmit RAM start address
func (r *registerOtg_hs_dieptxf0_deviceType) SetTx0fsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf0_deviceFieldTx0fsaMask)|(uint32(value)<<RegisterOtg_hs_dieptxf0_deviceFieldTx0fsaShift))
}

const (
	RegisterOtg_hs_dieptxf0_deviceFieldTx0fdShift = 16
	RegisterOtg_hs_dieptxf0_deviceFieldTx0fdMask  = 0xffff0000
)

// GetTx0fd Endpoint 0 TxFIFO depth
func (r *registerOtg_hs_dieptxf0_deviceType) GetTx0fd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf0_deviceFieldTx0fdMask) >> RegisterOtg_hs_dieptxf0_deviceFieldTx0fdShift)
}

// SetTx0fd Endpoint 0 TxFIFO depth
func (r *registerOtg_hs_dieptxf0_deviceType) SetTx0fd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf0_deviceFieldTx0fdMask)|(uint32(value)<<RegisterOtg_hs_dieptxf0_deviceFieldTx0fdShift))
}

// registerOtg_hs_gnptxstsType OTG_HS nonperiodic transmit FIFO/queue status register
type registerOtg_hs_gnptxstsType uint32

const (
	RegisterOtg_hs_gnptxstsFieldNptxfsavShift = 0
	RegisterOtg_hs_gnptxstsFieldNptxfsavMask  = 0xffff
)

// GetNptxfsav Nonperiodic TxFIFO space available
func (r *registerOtg_hs_gnptxstsType) GetNptxfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gnptxstsFieldNptxfsavMask) >> RegisterOtg_hs_gnptxstsFieldNptxfsavShift)
}

// SetNptxfsav Nonperiodic TxFIFO space available
func (r *registerOtg_hs_gnptxstsType) SetNptxfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gnptxstsFieldNptxfsavMask)|(uint32(value)<<RegisterOtg_hs_gnptxstsFieldNptxfsavShift))
}

const (
	RegisterOtg_hs_gnptxstsFieldNptqxsavShift = 16
	RegisterOtg_hs_gnptxstsFieldNptqxsavMask  = 0xff0000
)

// GetNptqxsav Nonperiodic transmit request queue space available
func (r *registerOtg_hs_gnptxstsType) GetNptqxsav() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gnptxstsFieldNptqxsavMask) >> RegisterOtg_hs_gnptxstsFieldNptqxsavShift)
}

// SetNptqxsav Nonperiodic transmit request queue space available
func (r *registerOtg_hs_gnptxstsType) SetNptqxsav(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gnptxstsFieldNptqxsavMask)|(uint32(value)<<RegisterOtg_hs_gnptxstsFieldNptqxsavShift))
}

const (
	RegisterOtg_hs_gnptxstsFieldNptxqtopShift = 24
	RegisterOtg_hs_gnptxstsFieldNptxqtopMask  = 0x7f000000
)

// GetNptxqtop Top of the nonperiodic transmit request queue
func (r *registerOtg_hs_gnptxstsType) GetNptxqtop() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gnptxstsFieldNptxqtopMask) >> RegisterOtg_hs_gnptxstsFieldNptxqtopShift)
}

// SetNptxqtop Top of the nonperiodic transmit request queue
func (r *registerOtg_hs_gnptxstsType) SetNptxqtop(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gnptxstsFieldNptxqtopMask)|(uint32(value)<<RegisterOtg_hs_gnptxstsFieldNptxqtopShift))
}

// registerOtg_hs_gccfgType OTG_HS general core configuration register
type registerOtg_hs_gccfgType uint32

const (
	RegisterOtg_hs_gccfgFieldPwrdwnShift = 16
	RegisterOtg_hs_gccfgFieldPwrdwnMask  = 0x10000
)

// GetPwrdwn Power down
func (r *registerOtg_hs_gccfgType) GetPwrdwn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gccfgFieldPwrdwnMask) != 0
}

// SetPwrdwn Power down
func (r *registerOtg_hs_gccfgType) SetPwrdwn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gccfgFieldPwrdwnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gccfgFieldPwrdwnMask)
	}
}

const (
	RegisterOtg_hs_gccfgFieldBcdenShift = 17
	RegisterOtg_hs_gccfgFieldBcdenMask  = 0x20000
)

// GetBcden Battery charging detector (BCD) enable
func (r *registerOtg_hs_gccfgType) GetBcden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gccfgFieldBcdenMask) != 0
}

// SetBcden Battery charging detector (BCD) enable
func (r *registerOtg_hs_gccfgType) SetBcden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gccfgFieldBcdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gccfgFieldBcdenMask)
	}
}

const (
	RegisterOtg_hs_gccfgFieldDcdenShift = 18
	RegisterOtg_hs_gccfgFieldDcdenMask  = 0x40000
)

// GetDcden Data contact detection (DCD) mode enable
func (r *registerOtg_hs_gccfgType) GetDcden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gccfgFieldDcdenMask) != 0
}

// SetDcden Data contact detection (DCD) mode enable
func (r *registerOtg_hs_gccfgType) SetDcden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gccfgFieldDcdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gccfgFieldDcdenMask)
	}
}

const (
	RegisterOtg_hs_gccfgFieldPdenShift = 19
	RegisterOtg_hs_gccfgFieldPdenMask  = 0x80000
)

// GetPden Primary detection (PD) mode enable
func (r *registerOtg_hs_gccfgType) GetPden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gccfgFieldPdenMask) != 0
}

// SetPden Primary detection (PD) mode enable
func (r *registerOtg_hs_gccfgType) SetPden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gccfgFieldPdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gccfgFieldPdenMask)
	}
}

const (
	RegisterOtg_hs_gccfgFieldSdenShift = 20
	RegisterOtg_hs_gccfgFieldSdenMask  = 0x100000
)

// GetSden Secondary detection (SD) mode enable
func (r *registerOtg_hs_gccfgType) GetSden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gccfgFieldSdenMask) != 0
}

// SetSden Secondary detection (SD) mode enable
func (r *registerOtg_hs_gccfgType) SetSden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gccfgFieldSdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gccfgFieldSdenMask)
	}
}

const (
	RegisterOtg_hs_gccfgFieldVbdenShift = 21
	RegisterOtg_hs_gccfgFieldVbdenMask  = 0x200000
)

// GetVbden USB VBUS detection enable
func (r *registerOtg_hs_gccfgType) GetVbden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gccfgFieldVbdenMask) != 0
}

// SetVbden USB VBUS detection enable
func (r *registerOtg_hs_gccfgType) SetVbden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gccfgFieldVbdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gccfgFieldVbdenMask)
	}
}

const (
	RegisterOtg_hs_gccfgFieldDcdetShift = 0
	RegisterOtg_hs_gccfgFieldDcdetMask  = 0x1
)

// GetDcdet Data contact detection (DCD) status
func (r *registerOtg_hs_gccfgType) GetDcdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gccfgFieldDcdetMask) != 0
}

// SetDcdet Data contact detection (DCD) status
func (r *registerOtg_hs_gccfgType) SetDcdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gccfgFieldDcdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gccfgFieldDcdetMask)
	}
}

const (
	RegisterOtg_hs_gccfgFieldPdetShift = 1
	RegisterOtg_hs_gccfgFieldPdetMask  = 0x2
)

// GetPdet Primary detection (PD) status
func (r *registerOtg_hs_gccfgType) GetPdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gccfgFieldPdetMask) != 0
}

// SetPdet Primary detection (PD) status
func (r *registerOtg_hs_gccfgType) SetPdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gccfgFieldPdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gccfgFieldPdetMask)
	}
}

const (
	RegisterOtg_hs_gccfgFieldSdetShift = 2
	RegisterOtg_hs_gccfgFieldSdetMask  = 0x4
)

// GetSdet Secondary detection (SD) status
func (r *registerOtg_hs_gccfgType) GetSdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gccfgFieldSdetMask) != 0
}

// SetSdet Secondary detection (SD) status
func (r *registerOtg_hs_gccfgType) SetSdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gccfgFieldSdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gccfgFieldSdetMask)
	}
}

const (
	RegisterOtg_hs_gccfgFieldPs2detShift = 3
	RegisterOtg_hs_gccfgFieldPs2detMask  = 0x8
)

// GetPs2det DM pull-up detection status
func (r *registerOtg_hs_gccfgType) GetPs2det() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_gccfgFieldPs2detMask) != 0
}

// SetPs2det DM pull-up detection status
func (r *registerOtg_hs_gccfgType) SetPs2det(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_gccfgFieldPs2detMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_gccfgFieldPs2detMask)
	}
}

// registerOtg_hs_cidType OTG_HS core ID register
type registerOtg_hs_cidType uint32

const (
	RegisterOtg_hs_cidFieldProduct_idShift = 0
	RegisterOtg_hs_cidFieldProduct_idMask  = 0xffffffff
)

// GetProduct_id Product ID field
func (r *registerOtg_hs_cidType) GetProduct_id() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_cidFieldProduct_idMask) >> RegisterOtg_hs_cidFieldProduct_idShift)
}

// SetProduct_id Product ID field
func (r *registerOtg_hs_cidType) SetProduct_id(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_cidFieldProduct_idMask)|(uint32(value)<<RegisterOtg_hs_cidFieldProduct_idShift))
}

// registerOtg_hs_hptxfsizType OTG_HS Host periodic transmit FIFO size register
type registerOtg_hs_hptxfsizType uint32

const (
	RegisterOtg_hs_hptxfsizFieldPtxsaShift = 0
	RegisterOtg_hs_hptxfsizFieldPtxsaMask  = 0xffff
)

// GetPtxsa Host periodic TxFIFO start address
func (r *registerOtg_hs_hptxfsizType) GetPtxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hptxfsizFieldPtxsaMask) >> RegisterOtg_hs_hptxfsizFieldPtxsaShift)
}

// SetPtxsa Host periodic TxFIFO start address
func (r *registerOtg_hs_hptxfsizType) SetPtxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hptxfsizFieldPtxsaMask)|(uint32(value)<<RegisterOtg_hs_hptxfsizFieldPtxsaShift))
}

const (
	RegisterOtg_hs_hptxfsizFieldPtxfdShift = 16
	RegisterOtg_hs_hptxfsizFieldPtxfdMask  = 0xffff0000
)

// GetPtxfd Host periodic TxFIFO depth
func (r *registerOtg_hs_hptxfsizType) GetPtxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hptxfsizFieldPtxfdMask) >> RegisterOtg_hs_hptxfsizFieldPtxfdShift)
}

// SetPtxfd Host periodic TxFIFO depth
func (r *registerOtg_hs_hptxfsizType) SetPtxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hptxfsizFieldPtxfdMask)|(uint32(value)<<RegisterOtg_hs_hptxfsizFieldPtxfdShift))
}

// registerOtg_hs_dieptxf1Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtg_hs_dieptxf1Type uint32

const (
	RegisterOtg_hs_dieptxf1FieldIneptxsaShift = 0
	RegisterOtg_hs_dieptxf1FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf1Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf1FieldIneptxsaMask) >> RegisterOtg_hs_dieptxf1FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf1Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf1FieldIneptxsaMask)|(uint32(value)<<RegisterOtg_hs_dieptxf1FieldIneptxsaShift))
}

const (
	RegisterOtg_hs_dieptxf1FieldIneptxfdShift = 16
	RegisterOtg_hs_dieptxf1FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf1Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf1FieldIneptxfdMask) >> RegisterOtg_hs_dieptxf1FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf1Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf1FieldIneptxfdMask)|(uint32(value)<<RegisterOtg_hs_dieptxf1FieldIneptxfdShift))
}

// registerOtg_hs_dieptxf2Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtg_hs_dieptxf2Type uint32

const (
	RegisterOtg_hs_dieptxf2FieldIneptxsaShift = 0
	RegisterOtg_hs_dieptxf2FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf2Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf2FieldIneptxsaMask) >> RegisterOtg_hs_dieptxf2FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf2Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf2FieldIneptxsaMask)|(uint32(value)<<RegisterOtg_hs_dieptxf2FieldIneptxsaShift))
}

const (
	RegisterOtg_hs_dieptxf2FieldIneptxfdShift = 16
	RegisterOtg_hs_dieptxf2FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf2Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf2FieldIneptxfdMask) >> RegisterOtg_hs_dieptxf2FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf2Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf2FieldIneptxfdMask)|(uint32(value)<<RegisterOtg_hs_dieptxf2FieldIneptxfdShift))
}

// registerOtg_hs_dieptxf3Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtg_hs_dieptxf3Type uint32

const (
	RegisterOtg_hs_dieptxf3FieldIneptxsaShift = 0
	RegisterOtg_hs_dieptxf3FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf3Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf3FieldIneptxsaMask) >> RegisterOtg_hs_dieptxf3FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf3Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf3FieldIneptxsaMask)|(uint32(value)<<RegisterOtg_hs_dieptxf3FieldIneptxsaShift))
}

const (
	RegisterOtg_hs_dieptxf3FieldIneptxfdShift = 16
	RegisterOtg_hs_dieptxf3FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf3Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf3FieldIneptxfdMask) >> RegisterOtg_hs_dieptxf3FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf3Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf3FieldIneptxfdMask)|(uint32(value)<<RegisterOtg_hs_dieptxf3FieldIneptxfdShift))
}

// registerOtg_hs_dieptxf4Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtg_hs_dieptxf4Type uint32

const (
	RegisterOtg_hs_dieptxf4FieldIneptxsaShift = 0
	RegisterOtg_hs_dieptxf4FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf4Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf4FieldIneptxsaMask) >> RegisterOtg_hs_dieptxf4FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf4Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf4FieldIneptxsaMask)|(uint32(value)<<RegisterOtg_hs_dieptxf4FieldIneptxsaShift))
}

const (
	RegisterOtg_hs_dieptxf4FieldIneptxfdShift = 16
	RegisterOtg_hs_dieptxf4FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf4Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf4FieldIneptxfdMask) >> RegisterOtg_hs_dieptxf4FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf4Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf4FieldIneptxfdMask)|(uint32(value)<<RegisterOtg_hs_dieptxf4FieldIneptxfdShift))
}

// registerOtg_hs_dieptxf5Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtg_hs_dieptxf5Type uint32

const (
	RegisterOtg_hs_dieptxf5FieldIneptxsaShift = 0
	RegisterOtg_hs_dieptxf5FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf5Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf5FieldIneptxsaMask) >> RegisterOtg_hs_dieptxf5FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf5Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf5FieldIneptxsaMask)|(uint32(value)<<RegisterOtg_hs_dieptxf5FieldIneptxsaShift))
}

const (
	RegisterOtg_hs_dieptxf5FieldIneptxfdShift = 16
	RegisterOtg_hs_dieptxf5FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf5Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf5FieldIneptxfdMask) >> RegisterOtg_hs_dieptxf5FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf5Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf5FieldIneptxfdMask)|(uint32(value)<<RegisterOtg_hs_dieptxf5FieldIneptxfdShift))
}

// registerOtg_hs_dieptxf6Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtg_hs_dieptxf6Type uint32

const (
	RegisterOtg_hs_dieptxf6FieldIneptxsaShift = 0
	RegisterOtg_hs_dieptxf6FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf6Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf6FieldIneptxsaMask) >> RegisterOtg_hs_dieptxf6FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf6Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf6FieldIneptxsaMask)|(uint32(value)<<RegisterOtg_hs_dieptxf6FieldIneptxsaShift))
}

const (
	RegisterOtg_hs_dieptxf6FieldIneptxfdShift = 16
	RegisterOtg_hs_dieptxf6FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf6Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf6FieldIneptxfdMask) >> RegisterOtg_hs_dieptxf6FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf6Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf6FieldIneptxfdMask)|(uint32(value)<<RegisterOtg_hs_dieptxf6FieldIneptxfdShift))
}

// registerOtg_hs_dieptxf7Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtg_hs_dieptxf7Type uint32

const (
	RegisterOtg_hs_dieptxf7FieldIneptxsaShift = 0
	RegisterOtg_hs_dieptxf7FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf7Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf7FieldIneptxsaMask) >> RegisterOtg_hs_dieptxf7FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtg_hs_dieptxf7Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf7FieldIneptxsaMask)|(uint32(value)<<RegisterOtg_hs_dieptxf7FieldIneptxsaShift))
}

const (
	RegisterOtg_hs_dieptxf7FieldIneptxfdShift = 16
	RegisterOtg_hs_dieptxf7FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf7Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_dieptxf7FieldIneptxfdMask) >> RegisterOtg_hs_dieptxf7FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtg_hs_dieptxf7Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_dieptxf7FieldIneptxfdMask)|(uint32(value)<<RegisterOtg_hs_dieptxf7FieldIneptxfdShift))
}

// registerOtg_hs_grxstsr_deviceType OTG_HS Receive status debug read register (peripheral mode mode)
type registerOtg_hs_grxstsr_deviceType uint32

const (
	RegisterOtg_hs_grxstsr_deviceFieldEpnumShift = 0
	RegisterOtg_hs_grxstsr_deviceFieldEpnumMask  = 0xf
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_grxstsr_deviceType) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsr_deviceFieldEpnumMask) >> RegisterOtg_hs_grxstsr_deviceFieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_grxstsr_deviceType) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsr_deviceFieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_grxstsr_deviceFieldEpnumShift))
}

const (
	RegisterOtg_hs_grxstsr_deviceFieldBcntShift = 4
	RegisterOtg_hs_grxstsr_deviceFieldBcntMask  = 0x7ff0
)

// GetBcnt Byte count
func (r *registerOtg_hs_grxstsr_deviceType) GetBcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsr_deviceFieldBcntMask) >> RegisterOtg_hs_grxstsr_deviceFieldBcntShift)
}

// SetBcnt Byte count
func (r *registerOtg_hs_grxstsr_deviceType) SetBcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsr_deviceFieldBcntMask)|(uint32(value)<<RegisterOtg_hs_grxstsr_deviceFieldBcntShift))
}

const (
	RegisterOtg_hs_grxstsr_deviceFieldDpidShift = 15
	RegisterOtg_hs_grxstsr_deviceFieldDpidMask  = 0x18000
)

// GetDpid Data PID
func (r *registerOtg_hs_grxstsr_deviceType) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsr_deviceFieldDpidMask) >> RegisterOtg_hs_grxstsr_deviceFieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_grxstsr_deviceType) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsr_deviceFieldDpidMask)|(uint32(value)<<RegisterOtg_hs_grxstsr_deviceFieldDpidShift))
}

const (
	RegisterOtg_hs_grxstsr_deviceFieldPktstsShift = 17
	RegisterOtg_hs_grxstsr_deviceFieldPktstsMask  = 0x1e0000
)

// GetPktsts Packet status
func (r *registerOtg_hs_grxstsr_deviceType) GetPktsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsr_deviceFieldPktstsMask) >> RegisterOtg_hs_grxstsr_deviceFieldPktstsShift)
}

// SetPktsts Packet status
func (r *registerOtg_hs_grxstsr_deviceType) SetPktsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsr_deviceFieldPktstsMask)|(uint32(value)<<RegisterOtg_hs_grxstsr_deviceFieldPktstsShift))
}

const (
	RegisterOtg_hs_grxstsr_deviceFieldFrmnumShift = 21
	RegisterOtg_hs_grxstsr_deviceFieldFrmnumMask  = 0x1e00000
)

// GetFrmnum Frame number
func (r *registerOtg_hs_grxstsr_deviceType) GetFrmnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsr_deviceFieldFrmnumMask) >> RegisterOtg_hs_grxstsr_deviceFieldFrmnumShift)
}

// SetFrmnum Frame number
func (r *registerOtg_hs_grxstsr_deviceType) SetFrmnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsr_deviceFieldFrmnumMask)|(uint32(value)<<RegisterOtg_hs_grxstsr_deviceFieldFrmnumShift))
}

// registerOtg_hs_grxstsp_deviceType OTG_HS status read and pop register (peripheral mode)
type registerOtg_hs_grxstsp_deviceType uint32

const (
	RegisterOtg_hs_grxstsp_deviceFieldEpnumShift = 0
	RegisterOtg_hs_grxstsp_deviceFieldEpnumMask  = 0xf
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_grxstsp_deviceType) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsp_deviceFieldEpnumMask) >> RegisterOtg_hs_grxstsp_deviceFieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_grxstsp_deviceType) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsp_deviceFieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_grxstsp_deviceFieldEpnumShift))
}

const (
	RegisterOtg_hs_grxstsp_deviceFieldBcntShift = 4
	RegisterOtg_hs_grxstsp_deviceFieldBcntMask  = 0x7ff0
)

// GetBcnt Byte count
func (r *registerOtg_hs_grxstsp_deviceType) GetBcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsp_deviceFieldBcntMask) >> RegisterOtg_hs_grxstsp_deviceFieldBcntShift)
}

// SetBcnt Byte count
func (r *registerOtg_hs_grxstsp_deviceType) SetBcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsp_deviceFieldBcntMask)|(uint32(value)<<RegisterOtg_hs_grxstsp_deviceFieldBcntShift))
}

const (
	RegisterOtg_hs_grxstsp_deviceFieldDpidShift = 15
	RegisterOtg_hs_grxstsp_deviceFieldDpidMask  = 0x18000
)

// GetDpid Data PID
func (r *registerOtg_hs_grxstsp_deviceType) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsp_deviceFieldDpidMask) >> RegisterOtg_hs_grxstsp_deviceFieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_grxstsp_deviceType) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsp_deviceFieldDpidMask)|(uint32(value)<<RegisterOtg_hs_grxstsp_deviceFieldDpidShift))
}

const (
	RegisterOtg_hs_grxstsp_deviceFieldPktstsShift = 17
	RegisterOtg_hs_grxstsp_deviceFieldPktstsMask  = 0x1e0000
)

// GetPktsts Packet status
func (r *registerOtg_hs_grxstsp_deviceType) GetPktsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsp_deviceFieldPktstsMask) >> RegisterOtg_hs_grxstsp_deviceFieldPktstsShift)
}

// SetPktsts Packet status
func (r *registerOtg_hs_grxstsp_deviceType) SetPktsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsp_deviceFieldPktstsMask)|(uint32(value)<<RegisterOtg_hs_grxstsp_deviceFieldPktstsShift))
}

const (
	RegisterOtg_hs_grxstsp_deviceFieldFrmnumShift = 21
	RegisterOtg_hs_grxstsp_deviceFieldFrmnumMask  = 0x1e00000
)

// GetFrmnum Frame number
func (r *registerOtg_hs_grxstsp_deviceType) GetFrmnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_grxstsp_deviceFieldFrmnumMask) >> RegisterOtg_hs_grxstsp_deviceFieldFrmnumShift)
}

// SetFrmnum Frame number
func (r *registerOtg_hs_grxstsp_deviceType) SetFrmnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_grxstsp_deviceFieldFrmnumMask)|(uint32(value)<<RegisterOtg_hs_grxstsp_deviceFieldFrmnumShift))
}

// registerOtg_hs_glpmcfgType OTG core LPM configuration register
type registerOtg_hs_glpmcfgType uint32

const (
	RegisterOtg_hs_glpmcfgFieldLpmenShift = 0
	RegisterOtg_hs_glpmcfgFieldLpmenMask  = 0x1
)

// GetLpmen LPM support enable
func (r *registerOtg_hs_glpmcfgType) GetLpmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldLpmenMask) != 0
}

// SetLpmen LPM support enable
func (r *registerOtg_hs_glpmcfgType) SetLpmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_glpmcfgFieldLpmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_glpmcfgFieldLpmenMask)
	}
}

const (
	RegisterOtg_hs_glpmcfgFieldLpmackShift = 1
	RegisterOtg_hs_glpmcfgFieldLpmackMask  = 0x2
)

// GetLpmack LPM token acknowledge enable
func (r *registerOtg_hs_glpmcfgType) GetLpmack() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldLpmackMask) != 0
}

// SetLpmack LPM token acknowledge enable
func (r *registerOtg_hs_glpmcfgType) SetLpmack(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_glpmcfgFieldLpmackMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_glpmcfgFieldLpmackMask)
	}
}

const (
	RegisterOtg_hs_glpmcfgFieldBeslShift = 2
	RegisterOtg_hs_glpmcfgFieldBeslMask  = 0x3c
)

// GetBesl Best effort service latency
func (r *registerOtg_hs_glpmcfgType) GetBesl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldBeslMask) >> RegisterOtg_hs_glpmcfgFieldBeslShift)
}

const (
	RegisterOtg_hs_glpmcfgFieldRemwakeShift = 6
	RegisterOtg_hs_glpmcfgFieldRemwakeMask  = 0x40
)

// GetRemwake bRemoteWake value
func (r *registerOtg_hs_glpmcfgType) GetRemwake() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldRemwakeMask) != 0
}

const (
	RegisterOtg_hs_glpmcfgFieldL1ssenShift = 7
	RegisterOtg_hs_glpmcfgFieldL1ssenMask  = 0x80
)

// GetL1ssen L1 Shallow Sleep enable
func (r *registerOtg_hs_glpmcfgType) GetL1ssen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldL1ssenMask) != 0
}

// SetL1ssen L1 Shallow Sleep enable
func (r *registerOtg_hs_glpmcfgType) SetL1ssen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_glpmcfgFieldL1ssenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_glpmcfgFieldL1ssenMask)
	}
}

const (
	RegisterOtg_hs_glpmcfgFieldBeslthrsShift = 8
	RegisterOtg_hs_glpmcfgFieldBeslthrsMask  = 0xf00
)

// GetBeslthrs BESL threshold
func (r *registerOtg_hs_glpmcfgType) GetBeslthrs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldBeslthrsMask) >> RegisterOtg_hs_glpmcfgFieldBeslthrsShift)
}

// SetBeslthrs BESL threshold
func (r *registerOtg_hs_glpmcfgType) SetBeslthrs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_glpmcfgFieldBeslthrsMask)|(uint32(value)<<RegisterOtg_hs_glpmcfgFieldBeslthrsShift))
}

const (
	RegisterOtg_hs_glpmcfgFieldL1dsenShift = 12
	RegisterOtg_hs_glpmcfgFieldL1dsenMask  = 0x1000
)

// GetL1dsen L1 deep sleep enable
func (r *registerOtg_hs_glpmcfgType) GetL1dsen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldL1dsenMask) != 0
}

// SetL1dsen L1 deep sleep enable
func (r *registerOtg_hs_glpmcfgType) SetL1dsen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_glpmcfgFieldL1dsenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_glpmcfgFieldL1dsenMask)
	}
}

const (
	RegisterOtg_hs_glpmcfgFieldLpmrstShift = 13
	RegisterOtg_hs_glpmcfgFieldLpmrstMask  = 0x6000
)

// GetLpmrst LPM response
func (r *registerOtg_hs_glpmcfgType) GetLpmrst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldLpmrstMask) >> RegisterOtg_hs_glpmcfgFieldLpmrstShift)
}

const (
	RegisterOtg_hs_glpmcfgFieldSlpstsShift = 15
	RegisterOtg_hs_glpmcfgFieldSlpstsMask  = 0x8000
)

// GetSlpsts Port sleep status
func (r *registerOtg_hs_glpmcfgType) GetSlpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldSlpstsMask) != 0
}

const (
	RegisterOtg_hs_glpmcfgFieldL1rsmokShift = 16
	RegisterOtg_hs_glpmcfgFieldL1rsmokMask  = 0x10000
)

// GetL1rsmok Sleep State Resume OK
func (r *registerOtg_hs_glpmcfgType) GetL1rsmok() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldL1rsmokMask) != 0
}

const (
	RegisterOtg_hs_glpmcfgFieldLpmchidxShift = 17
	RegisterOtg_hs_glpmcfgFieldLpmchidxMask  = 0x1e0000
)

// GetLpmchidx LPM Channel Index
func (r *registerOtg_hs_glpmcfgType) GetLpmchidx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldLpmchidxMask) >> RegisterOtg_hs_glpmcfgFieldLpmchidxShift)
}

// SetLpmchidx LPM Channel Index
func (r *registerOtg_hs_glpmcfgType) SetLpmchidx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_glpmcfgFieldLpmchidxMask)|(uint32(value)<<RegisterOtg_hs_glpmcfgFieldLpmchidxShift))
}

const (
	RegisterOtg_hs_glpmcfgFieldLpmrcntShift = 21
	RegisterOtg_hs_glpmcfgFieldLpmrcntMask  = 0xe00000
)

// GetLpmrcnt LPM retry count
func (r *registerOtg_hs_glpmcfgType) GetLpmrcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldLpmrcntMask) >> RegisterOtg_hs_glpmcfgFieldLpmrcntShift)
}

// SetLpmrcnt LPM retry count
func (r *registerOtg_hs_glpmcfgType) SetLpmrcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_glpmcfgFieldLpmrcntMask)|(uint32(value)<<RegisterOtg_hs_glpmcfgFieldLpmrcntShift))
}

const (
	RegisterOtg_hs_glpmcfgFieldSndlpmShift = 24
	RegisterOtg_hs_glpmcfgFieldSndlpmMask  = 0x1000000
)

// GetSndlpm Send LPM transaction
func (r *registerOtg_hs_glpmcfgType) GetSndlpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldSndlpmMask) != 0
}

// SetSndlpm Send LPM transaction
func (r *registerOtg_hs_glpmcfgType) SetSndlpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_glpmcfgFieldSndlpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_glpmcfgFieldSndlpmMask)
	}
}

const (
	RegisterOtg_hs_glpmcfgFieldLpmrcntstsShift = 25
	RegisterOtg_hs_glpmcfgFieldLpmrcntstsMask  = 0xe000000
)

// GetLpmrcntsts LPM retry count status
func (r *registerOtg_hs_glpmcfgType) GetLpmrcntsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldLpmrcntstsMask) >> RegisterOtg_hs_glpmcfgFieldLpmrcntstsShift)
}

const (
	RegisterOtg_hs_glpmcfgFieldEnbeslShift = 28
	RegisterOtg_hs_glpmcfgFieldEnbeslMask  = 0x10000000
)

// GetEnbesl Enable best effort service latency
func (r *registerOtg_hs_glpmcfgType) GetEnbesl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_glpmcfgFieldEnbeslMask) != 0
}

// SetEnbesl Enable best effort service latency
func (r *registerOtg_hs_glpmcfgType) SetEnbesl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_glpmcfgFieldEnbeslMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_glpmcfgFieldEnbeslMask)
	}
}
