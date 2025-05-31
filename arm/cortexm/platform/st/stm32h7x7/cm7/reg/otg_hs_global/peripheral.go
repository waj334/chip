//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package otg_hs_global

import (
	"unsafe"
	"volatile"
)

var (
	Otg1_hs_global = (*_otg_hs_global)(unsafe.Pointer(uintptr(0x40040000)))
	Otg2_hs_global = (*_otg_hs_global)(unsafe.Pointer(uintptr(0x40080000)))

	Instances = [2]*_otg_hs_global{
		Otg1_hs_global,
		Otg2_hs_global,
	}
)

type _otg_hs_global struct {
	Otghsgotgctl        registerOtghsgotgctlType
	Otghsgotgint        registerOtghsgotgintType
	Otghsgahbcfg        registerOtghsgahbcfgType
	Otghsgusbcfg        registerOtghsgusbcfgType
	Otghsgrstctl        registerOtghsgrstctlType
	Otghsgintsts        registerOtghsgintstsType
	Otghsgintmsk        registerOtghsgintmskType
	Otghsgrxstsrhost    registerOtghsgrxstsrhostType
	Otghsgrxstsrdevice  registerOtghsgrxstsrdeviceType
	Otghsgrxstsphost    registerOtghsgrxstsphostType
	Otghsgrxstspdevice  registerOtghsgrxstspdeviceType
	Otghsgrxfsiz        registerOtghsgrxfsizType
	Otghshnptxfsizhost  registerOtghshnptxfsizhostType
	Otghsdieptxf0device registerOtghsdieptxf0deviceType
	Otghsgnptxsts       registerOtghsgnptxstsType
	_                   [8]byte
	Otghsgccfg          registerOtghsgccfgType
	Otghscid            registerOtghscidType
	_                   [20]byte
	Otghsglpmcfg        registerOtghsglpmcfgType
	_                   [168]byte
	Otghshptxfsiz       registerOtghshptxfsizType
	Otghsdieptxf1       registerOtghsdieptxf1Type
	Otghsdieptxf2       registerOtghsdieptxf2Type
	Otghsdieptxf3       registerOtghsdieptxf3Type
	Otghsdieptxf4       registerOtghsdieptxf4Type
	Otghsdieptxf5       registerOtghsdieptxf5Type
	Otghsdieptxf6       registerOtghsdieptxf6Type
	Otghsdieptxf7       registerOtghsdieptxf7Type
	Otgdieptxf8         registerOtgdieptxf8Type
}

// registerOtghsgotgctlType OTG_HS control and status register
type registerOtghsgotgctlType uint32

const (
	RegisterOtghsgotgctlFieldSrqscsShift = 0
	RegisterOtghsgotgctlFieldSrqscsMask  = 0x1
)

// GetSrqscs Session request success
func (r *registerOtghsgotgctlType) GetSrqscs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgctlFieldSrqscsMask) != 0
}

const (
	RegisterOtghsgotgctlFieldSrqShift = 1
	RegisterOtghsgotgctlFieldSrqMask  = 0x2
)

// GetSrq Session request
func (r *registerOtghsgotgctlType) GetSrq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgctlFieldSrqMask) != 0
}

// SetSrq Session request
func (r *registerOtghsgotgctlType) SetSrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgotgctlFieldSrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgotgctlFieldSrqMask)
	}
}

const (
	RegisterOtghsgotgctlFieldHngscsShift = 8
	RegisterOtghsgotgctlFieldHngscsMask  = 0x100
)

// GetHngscs Host negotiation success
func (r *registerOtghsgotgctlType) GetHngscs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgctlFieldHngscsMask) != 0
}

const (
	RegisterOtghsgotgctlFieldHnprqShift = 9
	RegisterOtghsgotgctlFieldHnprqMask  = 0x200
)

// GetHnprq HNP request
func (r *registerOtghsgotgctlType) GetHnprq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgctlFieldHnprqMask) != 0
}

// SetHnprq HNP request
func (r *registerOtghsgotgctlType) SetHnprq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgotgctlFieldHnprqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgotgctlFieldHnprqMask)
	}
}

const (
	RegisterOtghsgotgctlFieldHshnpenShift = 10
	RegisterOtghsgotgctlFieldHshnpenMask  = 0x400
)

// GetHshnpen Host set HNP enable
func (r *registerOtghsgotgctlType) GetHshnpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgctlFieldHshnpenMask) != 0
}

// SetHshnpen Host set HNP enable
func (r *registerOtghsgotgctlType) SetHshnpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgotgctlFieldHshnpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgotgctlFieldHshnpenMask)
	}
}

const (
	RegisterOtghsgotgctlFieldDhnpenShift = 11
	RegisterOtghsgotgctlFieldDhnpenMask  = 0x800
)

// GetDhnpen Device HNP enabled
func (r *registerOtghsgotgctlType) GetDhnpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgctlFieldDhnpenMask) != 0
}

// SetDhnpen Device HNP enabled
func (r *registerOtghsgotgctlType) SetDhnpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgotgctlFieldDhnpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgotgctlFieldDhnpenMask)
	}
}

const (
	RegisterOtghsgotgctlFieldEhenShift = 12
	RegisterOtghsgotgctlFieldEhenMask  = 0x1000
)

// GetEhen Embedded host enable
func (r *registerOtghsgotgctlType) GetEhen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgctlFieldEhenMask) != 0
}

// SetEhen Embedded host enable
func (r *registerOtghsgotgctlType) SetEhen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgotgctlFieldEhenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgotgctlFieldEhenMask)
	}
}

const (
	RegisterOtghsgotgctlFieldCidstsShift = 16
	RegisterOtghsgotgctlFieldCidstsMask  = 0x10000
)

// GetCidsts Connector ID status
func (r *registerOtghsgotgctlType) GetCidsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgctlFieldCidstsMask) != 0
}

const (
	RegisterOtghsgotgctlFieldDbctShift = 17
	RegisterOtghsgotgctlFieldDbctMask  = 0x20000
)

// GetDbct Long/short debounce time
func (r *registerOtghsgotgctlType) GetDbct() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgctlFieldDbctMask) != 0
}

const (
	RegisterOtghsgotgctlFieldAsvldShift = 18
	RegisterOtghsgotgctlFieldAsvldMask  = 0x40000
)

// GetAsvld A-session valid
func (r *registerOtghsgotgctlType) GetAsvld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgctlFieldAsvldMask) != 0
}

const (
	RegisterOtghsgotgctlFieldBsvldShift = 19
	RegisterOtghsgotgctlFieldBsvldMask  = 0x80000
)

// GetBsvld B-session valid
func (r *registerOtghsgotgctlType) GetBsvld() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgctlFieldBsvldMask) != 0
}

// registerOtghsgotgintType OTG_HS interrupt register
type registerOtghsgotgintType uint32

const (
	RegisterOtghsgotgintFieldSedetShift = 2
	RegisterOtghsgotgintFieldSedetMask  = 0x4
)

// GetSedet Session end detected
func (r *registerOtghsgotgintType) GetSedet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgintFieldSedetMask) != 0
}

// SetSedet Session end detected
func (r *registerOtghsgotgintType) SetSedet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgotgintFieldSedetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgotgintFieldSedetMask)
	}
}

const (
	RegisterOtghsgotgintFieldSrsschgShift = 8
	RegisterOtghsgotgintFieldSrsschgMask  = 0x100
)

// GetSrsschg Session request success status change
func (r *registerOtghsgotgintType) GetSrsschg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgintFieldSrsschgMask) != 0
}

// SetSrsschg Session request success status change
func (r *registerOtghsgotgintType) SetSrsschg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgotgintFieldSrsschgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgotgintFieldSrsschgMask)
	}
}

const (
	RegisterOtghsgotgintFieldHnsschgShift = 9
	RegisterOtghsgotgintFieldHnsschgMask  = 0x200
)

// GetHnsschg Host negotiation success status change
func (r *registerOtghsgotgintType) GetHnsschg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgintFieldHnsschgMask) != 0
}

// SetHnsschg Host negotiation success status change
func (r *registerOtghsgotgintType) SetHnsschg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgotgintFieldHnsschgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgotgintFieldHnsschgMask)
	}
}

const (
	RegisterOtghsgotgintFieldHngdetShift = 17
	RegisterOtghsgotgintFieldHngdetMask  = 0x20000
)

// GetHngdet Host negotiation detected
func (r *registerOtghsgotgintType) GetHngdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgintFieldHngdetMask) != 0
}

// SetHngdet Host negotiation detected
func (r *registerOtghsgotgintType) SetHngdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgotgintFieldHngdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgotgintFieldHngdetMask)
	}
}

const (
	RegisterOtghsgotgintFieldAdtochgShift = 18
	RegisterOtghsgotgintFieldAdtochgMask  = 0x40000
)

// GetAdtochg A-device timeout change
func (r *registerOtghsgotgintType) GetAdtochg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgintFieldAdtochgMask) != 0
}

// SetAdtochg A-device timeout change
func (r *registerOtghsgotgintType) SetAdtochg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgotgintFieldAdtochgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgotgintFieldAdtochgMask)
	}
}

const (
	RegisterOtghsgotgintFieldDbcdneShift = 19
	RegisterOtghsgotgintFieldDbcdneMask  = 0x80000
)

// GetDbcdne Debounce done
func (r *registerOtghsgotgintType) GetDbcdne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgintFieldDbcdneMask) != 0
}

// SetDbcdne Debounce done
func (r *registerOtghsgotgintType) SetDbcdne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgotgintFieldDbcdneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgotgintFieldDbcdneMask)
	}
}

const (
	RegisterOtghsgotgintFieldIdchngShift = 20
	RegisterOtghsgotgintFieldIdchngMask  = 0x100000
)

// GetIdchng ID input pin changed
func (r *registerOtghsgotgintType) GetIdchng() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgotgintFieldIdchngMask) != 0
}

// SetIdchng ID input pin changed
func (r *registerOtghsgotgintType) SetIdchng(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgotgintFieldIdchngMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgotgintFieldIdchngMask)
	}
}

// registerOtghsgahbcfgType OTG_HS AHB configuration register
type registerOtghsgahbcfgType uint32

const (
	RegisterOtghsgahbcfgFieldGintShift = 0
	RegisterOtghsgahbcfgFieldGintMask  = 0x1
)

// GetGint Global interrupt mask
func (r *registerOtghsgahbcfgType) GetGint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgahbcfgFieldGintMask) != 0
}

// SetGint Global interrupt mask
func (r *registerOtghsgahbcfgType) SetGint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgahbcfgFieldGintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgahbcfgFieldGintMask)
	}
}

const (
	RegisterOtghsgahbcfgFieldHbstlenShift = 1
	RegisterOtghsgahbcfgFieldHbstlenMask  = 0x1e
)

// GetHbstlen Burst length/type
func (r *registerOtghsgahbcfgType) GetHbstlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgahbcfgFieldHbstlenMask) >> RegisterOtghsgahbcfgFieldHbstlenShift)
}

// SetHbstlen Burst length/type
func (r *registerOtghsgahbcfgType) SetHbstlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgahbcfgFieldHbstlenMask)|(uint32(value)<<RegisterOtghsgahbcfgFieldHbstlenShift))
}

const (
	RegisterOtghsgahbcfgFieldDmaenShift = 5
	RegisterOtghsgahbcfgFieldDmaenMask  = 0x20
)

// GetDmaen DMA enable
func (r *registerOtghsgahbcfgType) GetDmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgahbcfgFieldDmaenMask) != 0
}

// SetDmaen DMA enable
func (r *registerOtghsgahbcfgType) SetDmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgahbcfgFieldDmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgahbcfgFieldDmaenMask)
	}
}

const (
	RegisterOtghsgahbcfgFieldTxfelvlShift = 7
	RegisterOtghsgahbcfgFieldTxfelvlMask  = 0x80
)

// GetTxfelvl TxFIFO empty level
func (r *registerOtghsgahbcfgType) GetTxfelvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgahbcfgFieldTxfelvlMask) != 0
}

// SetTxfelvl TxFIFO empty level
func (r *registerOtghsgahbcfgType) SetTxfelvl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgahbcfgFieldTxfelvlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgahbcfgFieldTxfelvlMask)
	}
}

const (
	RegisterOtghsgahbcfgFieldPtxfelvlShift = 8
	RegisterOtghsgahbcfgFieldPtxfelvlMask  = 0x100
)

// GetPtxfelvl Periodic TxFIFO empty level
func (r *registerOtghsgahbcfgType) GetPtxfelvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgahbcfgFieldPtxfelvlMask) != 0
}

// SetPtxfelvl Periodic TxFIFO empty level
func (r *registerOtghsgahbcfgType) SetPtxfelvl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgahbcfgFieldPtxfelvlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgahbcfgFieldPtxfelvlMask)
	}
}

// registerOtghsgusbcfgType OTG_HS USB configuration register
type registerOtghsgusbcfgType uint32

const (
	RegisterOtghsgusbcfgFieldTocalShift = 0
	RegisterOtghsgusbcfgFieldTocalMask  = 0x7
)

// GetTocal FS timeout calibration
func (r *registerOtghsgusbcfgType) GetTocal() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldTocalMask) >> RegisterOtghsgusbcfgFieldTocalShift)
}

// SetTocal FS timeout calibration
func (r *registerOtghsgusbcfgType) SetTocal(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldTocalMask)|(uint32(value)<<RegisterOtghsgusbcfgFieldTocalShift))
}

const (
	RegisterOtghsgusbcfgFieldPhyselShift = 6
	RegisterOtghsgusbcfgFieldPhyselMask  = 0x40
)

// SetPhysel USB 2.0 high-speed ULPI PHY or USB 1.1 full-speed serial transceiver select
func (r *registerOtghsgusbcfgType) SetPhysel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldPhyselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldPhyselMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldSrpcapShift = 8
	RegisterOtghsgusbcfgFieldSrpcapMask  = 0x100
)

// GetSrpcap SRP-capable
func (r *registerOtghsgusbcfgType) GetSrpcap() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldSrpcapMask) != 0
}

// SetSrpcap SRP-capable
func (r *registerOtghsgusbcfgType) SetSrpcap(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldSrpcapMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldSrpcapMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldHnpcapShift = 9
	RegisterOtghsgusbcfgFieldHnpcapMask  = 0x200
)

// GetHnpcap HNP-capable
func (r *registerOtghsgusbcfgType) GetHnpcap() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldHnpcapMask) != 0
}

// SetHnpcap HNP-capable
func (r *registerOtghsgusbcfgType) SetHnpcap(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldHnpcapMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldHnpcapMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldTrdtShift = 10
	RegisterOtghsgusbcfgFieldTrdtMask  = 0x3c00
)

// GetTrdt USB turnaround time
func (r *registerOtghsgusbcfgType) GetTrdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldTrdtMask) >> RegisterOtghsgusbcfgFieldTrdtShift)
}

// SetTrdt USB turnaround time
func (r *registerOtghsgusbcfgType) SetTrdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldTrdtMask)|(uint32(value)<<RegisterOtghsgusbcfgFieldTrdtShift))
}

const (
	RegisterOtghsgusbcfgFieldPhylpcsShift = 15
	RegisterOtghsgusbcfgFieldPhylpcsMask  = 0x8000
)

// GetPhylpcs PHY Low-power clock select
func (r *registerOtghsgusbcfgType) GetPhylpcs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldPhylpcsMask) != 0
}

// SetPhylpcs PHY Low-power clock select
func (r *registerOtghsgusbcfgType) SetPhylpcs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldPhylpcsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldPhylpcsMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldUlpifslsShift = 17
	RegisterOtghsgusbcfgFieldUlpifslsMask  = 0x20000
)

// GetUlpifsls ULPI FS/LS select
func (r *registerOtghsgusbcfgType) GetUlpifsls() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldUlpifslsMask) != 0
}

// SetUlpifsls ULPI FS/LS select
func (r *registerOtghsgusbcfgType) SetUlpifsls(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldUlpifslsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldUlpifslsMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldUlpiarShift = 18
	RegisterOtghsgusbcfgFieldUlpiarMask  = 0x40000
)

// GetUlpiar ULPI Auto-resume
func (r *registerOtghsgusbcfgType) GetUlpiar() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldUlpiarMask) != 0
}

// SetUlpiar ULPI Auto-resume
func (r *registerOtghsgusbcfgType) SetUlpiar(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldUlpiarMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldUlpiarMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldUlpicsmShift = 19
	RegisterOtghsgusbcfgFieldUlpicsmMask  = 0x80000
)

// GetUlpicsm ULPI Clock SuspendM
func (r *registerOtghsgusbcfgType) GetUlpicsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldUlpicsmMask) != 0
}

// SetUlpicsm ULPI Clock SuspendM
func (r *registerOtghsgusbcfgType) SetUlpicsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldUlpicsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldUlpicsmMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldUlpievbusdShift = 20
	RegisterOtghsgusbcfgFieldUlpievbusdMask  = 0x100000
)

// GetUlpievbusd ULPI External VBUS Drive
func (r *registerOtghsgusbcfgType) GetUlpievbusd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldUlpievbusdMask) != 0
}

// SetUlpievbusd ULPI External VBUS Drive
func (r *registerOtghsgusbcfgType) SetUlpievbusd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldUlpievbusdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldUlpievbusdMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldUlpievbusiShift = 21
	RegisterOtghsgusbcfgFieldUlpievbusiMask  = 0x200000
)

// GetUlpievbusi ULPI external VBUS indicator
func (r *registerOtghsgusbcfgType) GetUlpievbusi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldUlpievbusiMask) != 0
}

// SetUlpievbusi ULPI external VBUS indicator
func (r *registerOtghsgusbcfgType) SetUlpievbusi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldUlpievbusiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldUlpievbusiMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldTsdpsShift = 22
	RegisterOtghsgusbcfgFieldTsdpsMask  = 0x400000
)

// GetTsdps TermSel DLine pulsing selection
func (r *registerOtghsgusbcfgType) GetTsdps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldTsdpsMask) != 0
}

// SetTsdps TermSel DLine pulsing selection
func (r *registerOtghsgusbcfgType) SetTsdps(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldTsdpsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldTsdpsMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldPcciShift = 23
	RegisterOtghsgusbcfgFieldPcciMask  = 0x800000
)

// GetPcci Indicator complement
func (r *registerOtghsgusbcfgType) GetPcci() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldPcciMask) != 0
}

// SetPcci Indicator complement
func (r *registerOtghsgusbcfgType) SetPcci(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldPcciMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldPcciMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldPtciShift = 24
	RegisterOtghsgusbcfgFieldPtciMask  = 0x1000000
)

// GetPtci Indicator pass through
func (r *registerOtghsgusbcfgType) GetPtci() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldPtciMask) != 0
}

// SetPtci Indicator pass through
func (r *registerOtghsgusbcfgType) SetPtci(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldPtciMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldPtciMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldUlpiipdShift = 25
	RegisterOtghsgusbcfgFieldUlpiipdMask  = 0x2000000
)

// GetUlpiipd ULPI interface protect disable
func (r *registerOtghsgusbcfgType) GetUlpiipd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldUlpiipdMask) != 0
}

// SetUlpiipd ULPI interface protect disable
func (r *registerOtghsgusbcfgType) SetUlpiipd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldUlpiipdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldUlpiipdMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldFhmodShift = 29
	RegisterOtghsgusbcfgFieldFhmodMask  = 0x20000000
)

// GetFhmod Forced host mode
func (r *registerOtghsgusbcfgType) GetFhmod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldFhmodMask) != 0
}

// SetFhmod Forced host mode
func (r *registerOtghsgusbcfgType) SetFhmod(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldFhmodMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldFhmodMask)
	}
}

const (
	RegisterOtghsgusbcfgFieldFdmodShift = 30
	RegisterOtghsgusbcfgFieldFdmodMask  = 0x40000000
)

// GetFdmod Forced peripheral mode
func (r *registerOtghsgusbcfgType) GetFdmod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgusbcfgFieldFdmodMask) != 0
}

// SetFdmod Forced peripheral mode
func (r *registerOtghsgusbcfgType) SetFdmod(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgusbcfgFieldFdmodMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgusbcfgFieldFdmodMask)
	}
}

// registerOtghsgrstctlType OTG_HS reset register
type registerOtghsgrstctlType uint32

const (
	RegisterOtghsgrstctlFieldCsrstShift = 0
	RegisterOtghsgrstctlFieldCsrstMask  = 0x1
)

// GetCsrst Core soft reset
func (r *registerOtghsgrstctlType) GetCsrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrstctlFieldCsrstMask) != 0
}

// SetCsrst Core soft reset
func (r *registerOtghsgrstctlType) SetCsrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgrstctlFieldCsrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrstctlFieldCsrstMask)
	}
}

const (
	RegisterOtghsgrstctlFieldHsrstShift = 1
	RegisterOtghsgrstctlFieldHsrstMask  = 0x2
)

// GetHsrst HCLK soft reset
func (r *registerOtghsgrstctlType) GetHsrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrstctlFieldHsrstMask) != 0
}

// SetHsrst HCLK soft reset
func (r *registerOtghsgrstctlType) SetHsrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgrstctlFieldHsrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrstctlFieldHsrstMask)
	}
}

const (
	RegisterOtghsgrstctlFieldFcrstShift = 2
	RegisterOtghsgrstctlFieldFcrstMask  = 0x4
)

// GetFcrst Host frame counter reset
func (r *registerOtghsgrstctlType) GetFcrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrstctlFieldFcrstMask) != 0
}

// SetFcrst Host frame counter reset
func (r *registerOtghsgrstctlType) SetFcrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgrstctlFieldFcrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrstctlFieldFcrstMask)
	}
}

const (
	RegisterOtghsgrstctlFieldRxfflshShift = 4
	RegisterOtghsgrstctlFieldRxfflshMask  = 0x10
)

// GetRxfflsh RxFIFO flush
func (r *registerOtghsgrstctlType) GetRxfflsh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrstctlFieldRxfflshMask) != 0
}

// SetRxfflsh RxFIFO flush
func (r *registerOtghsgrstctlType) SetRxfflsh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgrstctlFieldRxfflshMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrstctlFieldRxfflshMask)
	}
}

const (
	RegisterOtghsgrstctlFieldTxfflshShift = 5
	RegisterOtghsgrstctlFieldTxfflshMask  = 0x20
)

// GetTxfflsh TxFIFO flush
func (r *registerOtghsgrstctlType) GetTxfflsh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrstctlFieldTxfflshMask) != 0
}

// SetTxfflsh TxFIFO flush
func (r *registerOtghsgrstctlType) SetTxfflsh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgrstctlFieldTxfflshMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrstctlFieldTxfflshMask)
	}
}

const (
	RegisterOtghsgrstctlFieldTxfnumShift = 6
	RegisterOtghsgrstctlFieldTxfnumMask  = 0x7c0
)

// GetTxfnum TxFIFO number
func (r *registerOtghsgrstctlType) GetTxfnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrstctlFieldTxfnumMask) >> RegisterOtghsgrstctlFieldTxfnumShift)
}

// SetTxfnum TxFIFO number
func (r *registerOtghsgrstctlType) SetTxfnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrstctlFieldTxfnumMask)|(uint32(value)<<RegisterOtghsgrstctlFieldTxfnumShift))
}

const (
	RegisterOtghsgrstctlFieldDmareqShift = 30
	RegisterOtghsgrstctlFieldDmareqMask  = 0x40000000
)

// GetDmareq DMA request signal enabled for USB OTG HS
func (r *registerOtghsgrstctlType) GetDmareq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrstctlFieldDmareqMask) != 0
}

const (
	RegisterOtghsgrstctlFieldAhbidlShift = 31
	RegisterOtghsgrstctlFieldAhbidlMask  = 0x80000000
)

// GetAhbidl AHB master idle
func (r *registerOtghsgrstctlType) GetAhbidl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrstctlFieldAhbidlMask) != 0
}

// registerOtghsgintstsType OTG_HS core interrupt register
type registerOtghsgintstsType uint32

const (
	RegisterOtghsgintstsFieldCmodShift = 0
	RegisterOtghsgintstsFieldCmodMask  = 0x1
)

// GetCmod Current mode of operation
func (r *registerOtghsgintstsType) GetCmod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldCmodMask) != 0
}

const (
	RegisterOtghsgintstsFieldMmisShift = 1
	RegisterOtghsgintstsFieldMmisMask  = 0x2
)

// GetMmis Mode mismatch interrupt
func (r *registerOtghsgintstsType) GetMmis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldMmisMask) != 0
}

// SetMmis Mode mismatch interrupt
func (r *registerOtghsgintstsType) SetMmis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldMmisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldMmisMask)
	}
}

const (
	RegisterOtghsgintstsFieldOtgintShift = 2
	RegisterOtghsgintstsFieldOtgintMask  = 0x4
)

// GetOtgint OTG interrupt
func (r *registerOtghsgintstsType) GetOtgint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldOtgintMask) != 0
}

const (
	RegisterOtghsgintstsFieldSofShift = 3
	RegisterOtghsgintstsFieldSofMask  = 0x8
)

// GetSof Start of frame
func (r *registerOtghsgintstsType) GetSof() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldSofMask) != 0
}

// SetSof Start of frame
func (r *registerOtghsgintstsType) SetSof(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldSofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldSofMask)
	}
}

const (
	RegisterOtghsgintstsFieldRxflvlShift = 4
	RegisterOtghsgintstsFieldRxflvlMask  = 0x10
)

// GetRxflvl RxFIFO nonempty
func (r *registerOtghsgintstsType) GetRxflvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldRxflvlMask) != 0
}

const (
	RegisterOtghsgintstsFieldNptxfeShift = 5
	RegisterOtghsgintstsFieldNptxfeMask  = 0x20
)

// GetNptxfe Nonperiodic TxFIFO empty
func (r *registerOtghsgintstsType) GetNptxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldNptxfeMask) != 0
}

const (
	RegisterOtghsgintstsFieldGinakeffShift = 6
	RegisterOtghsgintstsFieldGinakeffMask  = 0x40
)

// GetGinakeff Global IN nonperiodic NAK effective
func (r *registerOtghsgintstsType) GetGinakeff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldGinakeffMask) != 0
}

const (
	RegisterOtghsgintstsFieldBoutnakeffShift = 7
	RegisterOtghsgintstsFieldBoutnakeffMask  = 0x80
)

// GetBoutnakeff Global OUT NAK effective
func (r *registerOtghsgintstsType) GetBoutnakeff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldBoutnakeffMask) != 0
}

const (
	RegisterOtghsgintstsFieldEsuspShift = 10
	RegisterOtghsgintstsFieldEsuspMask  = 0x400
)

// GetEsusp Early suspend
func (r *registerOtghsgintstsType) GetEsusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldEsuspMask) != 0
}

// SetEsusp Early suspend
func (r *registerOtghsgintstsType) SetEsusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldEsuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldEsuspMask)
	}
}

const (
	RegisterOtghsgintstsFieldUsbsuspShift = 11
	RegisterOtghsgintstsFieldUsbsuspMask  = 0x800
)

// GetUsbsusp USB suspend
func (r *registerOtghsgintstsType) GetUsbsusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldUsbsuspMask) != 0
}

// SetUsbsusp USB suspend
func (r *registerOtghsgintstsType) SetUsbsusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldUsbsuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldUsbsuspMask)
	}
}

const (
	RegisterOtghsgintstsFieldUsbrstShift = 12
	RegisterOtghsgintstsFieldUsbrstMask  = 0x1000
)

// GetUsbrst USB reset
func (r *registerOtghsgintstsType) GetUsbrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldUsbrstMask) != 0
}

// SetUsbrst USB reset
func (r *registerOtghsgintstsType) SetUsbrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldUsbrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldUsbrstMask)
	}
}

const (
	RegisterOtghsgintstsFieldEnumdneShift = 13
	RegisterOtghsgintstsFieldEnumdneMask  = 0x2000
)

// GetEnumdne Enumeration done
func (r *registerOtghsgintstsType) GetEnumdne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldEnumdneMask) != 0
}

// SetEnumdne Enumeration done
func (r *registerOtghsgintstsType) SetEnumdne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldEnumdneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldEnumdneMask)
	}
}

const (
	RegisterOtghsgintstsFieldIsoodrpShift = 14
	RegisterOtghsgintstsFieldIsoodrpMask  = 0x4000
)

// GetIsoodrp Isochronous OUT packet dropped interrupt
func (r *registerOtghsgintstsType) GetIsoodrp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldIsoodrpMask) != 0
}

// SetIsoodrp Isochronous OUT packet dropped interrupt
func (r *registerOtghsgintstsType) SetIsoodrp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldIsoodrpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldIsoodrpMask)
	}
}

const (
	RegisterOtghsgintstsFieldEopfShift = 15
	RegisterOtghsgintstsFieldEopfMask  = 0x8000
)

// GetEopf End of periodic frame interrupt
func (r *registerOtghsgintstsType) GetEopf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldEopfMask) != 0
}

// SetEopf End of periodic frame interrupt
func (r *registerOtghsgintstsType) SetEopf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldEopfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldEopfMask)
	}
}

const (
	RegisterOtghsgintstsFieldIepintShift = 18
	RegisterOtghsgintstsFieldIepintMask  = 0x40000
)

// GetIepint IN endpoint interrupt
func (r *registerOtghsgintstsType) GetIepint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldIepintMask) != 0
}

const (
	RegisterOtghsgintstsFieldOepintShift = 19
	RegisterOtghsgintstsFieldOepintMask  = 0x80000
)

// GetOepint OUT endpoint interrupt
func (r *registerOtghsgintstsType) GetOepint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldOepintMask) != 0
}

const (
	RegisterOtghsgintstsFieldIisoixfrShift = 20
	RegisterOtghsgintstsFieldIisoixfrMask  = 0x100000
)

// GetIisoixfr Incomplete isochronous IN transfer
func (r *registerOtghsgintstsType) GetIisoixfr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldIisoixfrMask) != 0
}

// SetIisoixfr Incomplete isochronous IN transfer
func (r *registerOtghsgintstsType) SetIisoixfr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldIisoixfrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldIisoixfrMask)
	}
}

const (
	RegisterOtghsgintstsFieldPxfrincompisooutShift = 21
	RegisterOtghsgintstsFieldPxfrincompisooutMask  = 0x200000
)

// GetPxfrincompisoout Incomplete periodic transfer
func (r *registerOtghsgintstsType) GetPxfrincompisoout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldPxfrincompisooutMask) != 0
}

// SetPxfrincompisoout Incomplete periodic transfer
func (r *registerOtghsgintstsType) SetPxfrincompisoout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldPxfrincompisooutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldPxfrincompisooutMask)
	}
}

const (
	RegisterOtghsgintstsFieldDatafsuspShift = 22
	RegisterOtghsgintstsFieldDatafsuspMask  = 0x400000
)

// GetDatafsusp Data fetch suspended
func (r *registerOtghsgintstsType) GetDatafsusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldDatafsuspMask) != 0
}

// SetDatafsusp Data fetch suspended
func (r *registerOtghsgintstsType) SetDatafsusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldDatafsuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldDatafsuspMask)
	}
}

const (
	RegisterOtghsgintstsFieldHprtintShift = 24
	RegisterOtghsgintstsFieldHprtintMask  = 0x1000000
)

// GetHprtint Host port interrupt
func (r *registerOtghsgintstsType) GetHprtint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldHprtintMask) != 0
}

const (
	RegisterOtghsgintstsFieldHcintShift = 25
	RegisterOtghsgintstsFieldHcintMask  = 0x2000000
)

// GetHcint Host channels interrupt
func (r *registerOtghsgintstsType) GetHcint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldHcintMask) != 0
}

const (
	RegisterOtghsgintstsFieldPtxfeShift = 26
	RegisterOtghsgintstsFieldPtxfeMask  = 0x4000000
)

// GetPtxfe Periodic TxFIFO empty
func (r *registerOtghsgintstsType) GetPtxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldPtxfeMask) != 0
}

const (
	RegisterOtghsgintstsFieldCidschgShift = 28
	RegisterOtghsgintstsFieldCidschgMask  = 0x10000000
)

// GetCidschg Connector ID status change
func (r *registerOtghsgintstsType) GetCidschg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldCidschgMask) != 0
}

// SetCidschg Connector ID status change
func (r *registerOtghsgintstsType) SetCidschg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldCidschgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldCidschgMask)
	}
}

const (
	RegisterOtghsgintstsFieldDiscintShift = 29
	RegisterOtghsgintstsFieldDiscintMask  = 0x20000000
)

// GetDiscint Disconnect detected interrupt
func (r *registerOtghsgintstsType) GetDiscint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldDiscintMask) != 0
}

// SetDiscint Disconnect detected interrupt
func (r *registerOtghsgintstsType) SetDiscint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldDiscintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldDiscintMask)
	}
}

const (
	RegisterOtghsgintstsFieldSrqintShift = 30
	RegisterOtghsgintstsFieldSrqintMask  = 0x40000000
)

// GetSrqint Session request/new session detected interrupt
func (r *registerOtghsgintstsType) GetSrqint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldSrqintMask) != 0
}

// SetSrqint Session request/new session detected interrupt
func (r *registerOtghsgintstsType) SetSrqint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldSrqintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldSrqintMask)
	}
}

const (
	RegisterOtghsgintstsFieldWkuintShift = 31
	RegisterOtghsgintstsFieldWkuintMask  = 0x80000000
)

// GetWkuint Resume/remote wakeup detected interrupt
func (r *registerOtghsgintstsType) GetWkuint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintstsFieldWkuintMask) != 0
}

// SetWkuint Resume/remote wakeup detected interrupt
func (r *registerOtghsgintstsType) SetWkuint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintstsFieldWkuintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintstsFieldWkuintMask)
	}
}

// registerOtghsgintmskType OTG_HS interrupt mask register
type registerOtghsgintmskType uint32

const (
	RegisterOtghsgintmskFieldMmismShift = 1
	RegisterOtghsgintmskFieldMmismMask  = 0x2
)

// GetMmism Mode mismatch interrupt mask
func (r *registerOtghsgintmskType) GetMmism() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldMmismMask) != 0
}

// SetMmism Mode mismatch interrupt mask
func (r *registerOtghsgintmskType) SetMmism(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldMmismMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldMmismMask)
	}
}

const (
	RegisterOtghsgintmskFieldOtgintShift = 2
	RegisterOtghsgintmskFieldOtgintMask  = 0x4
)

// GetOtgint OTG interrupt mask
func (r *registerOtghsgintmskType) GetOtgint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldOtgintMask) != 0
}

// SetOtgint OTG interrupt mask
func (r *registerOtghsgintmskType) SetOtgint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldOtgintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldOtgintMask)
	}
}

const (
	RegisterOtghsgintmskFieldSofmShift = 3
	RegisterOtghsgintmskFieldSofmMask  = 0x8
)

// GetSofm Start of frame mask
func (r *registerOtghsgintmskType) GetSofm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldSofmMask) != 0
}

// SetSofm Start of frame mask
func (r *registerOtghsgintmskType) SetSofm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldSofmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldSofmMask)
	}
}

const (
	RegisterOtghsgintmskFieldRxflvlmShift = 4
	RegisterOtghsgintmskFieldRxflvlmMask  = 0x10
)

// GetRxflvlm Receive FIFO nonempty mask
func (r *registerOtghsgintmskType) GetRxflvlm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldRxflvlmMask) != 0
}

// SetRxflvlm Receive FIFO nonempty mask
func (r *registerOtghsgintmskType) SetRxflvlm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldRxflvlmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldRxflvlmMask)
	}
}

const (
	RegisterOtghsgintmskFieldNptxfemShift = 5
	RegisterOtghsgintmskFieldNptxfemMask  = 0x20
)

// GetNptxfem Nonperiodic TxFIFO empty mask
func (r *registerOtghsgintmskType) GetNptxfem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldNptxfemMask) != 0
}

// SetNptxfem Nonperiodic TxFIFO empty mask
func (r *registerOtghsgintmskType) SetNptxfem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldNptxfemMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldNptxfemMask)
	}
}

const (
	RegisterOtghsgintmskFieldGinakeffmShift = 6
	RegisterOtghsgintmskFieldGinakeffmMask  = 0x40
)

// GetGinakeffm Global nonperiodic IN NAK effective mask
func (r *registerOtghsgintmskType) GetGinakeffm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldGinakeffmMask) != 0
}

// SetGinakeffm Global nonperiodic IN NAK effective mask
func (r *registerOtghsgintmskType) SetGinakeffm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldGinakeffmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldGinakeffmMask)
	}
}

const (
	RegisterOtghsgintmskFieldGonakeffmShift = 7
	RegisterOtghsgintmskFieldGonakeffmMask  = 0x80
)

// GetGonakeffm Global OUT NAK effective mask
func (r *registerOtghsgintmskType) GetGonakeffm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldGonakeffmMask) != 0
}

// SetGonakeffm Global OUT NAK effective mask
func (r *registerOtghsgintmskType) SetGonakeffm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldGonakeffmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldGonakeffmMask)
	}
}

const (
	RegisterOtghsgintmskFieldEsuspmShift = 10
	RegisterOtghsgintmskFieldEsuspmMask  = 0x400
)

// GetEsuspm Early suspend mask
func (r *registerOtghsgintmskType) GetEsuspm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldEsuspmMask) != 0
}

// SetEsuspm Early suspend mask
func (r *registerOtghsgintmskType) SetEsuspm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldEsuspmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldEsuspmMask)
	}
}

const (
	RegisterOtghsgintmskFieldUsbsuspmShift = 11
	RegisterOtghsgintmskFieldUsbsuspmMask  = 0x800
)

// GetUsbsuspm USB suspend mask
func (r *registerOtghsgintmskType) GetUsbsuspm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldUsbsuspmMask) != 0
}

// SetUsbsuspm USB suspend mask
func (r *registerOtghsgintmskType) SetUsbsuspm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldUsbsuspmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldUsbsuspmMask)
	}
}

const (
	RegisterOtghsgintmskFieldUsbrstShift = 12
	RegisterOtghsgintmskFieldUsbrstMask  = 0x1000
)

// GetUsbrst USB reset mask
func (r *registerOtghsgintmskType) GetUsbrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldUsbrstMask) != 0
}

// SetUsbrst USB reset mask
func (r *registerOtghsgintmskType) SetUsbrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldUsbrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldUsbrstMask)
	}
}

const (
	RegisterOtghsgintmskFieldEnumdnemShift = 13
	RegisterOtghsgintmskFieldEnumdnemMask  = 0x2000
)

// GetEnumdnem Enumeration done mask
func (r *registerOtghsgintmskType) GetEnumdnem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldEnumdnemMask) != 0
}

// SetEnumdnem Enumeration done mask
func (r *registerOtghsgintmskType) SetEnumdnem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldEnumdnemMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldEnumdnemMask)
	}
}

const (
	RegisterOtghsgintmskFieldIsoodrpmShift = 14
	RegisterOtghsgintmskFieldIsoodrpmMask  = 0x4000
)

// GetIsoodrpm Isochronous OUT packet dropped interrupt mask
func (r *registerOtghsgintmskType) GetIsoodrpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldIsoodrpmMask) != 0
}

// SetIsoodrpm Isochronous OUT packet dropped interrupt mask
func (r *registerOtghsgintmskType) SetIsoodrpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldIsoodrpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldIsoodrpmMask)
	}
}

const (
	RegisterOtghsgintmskFieldEopfmShift = 15
	RegisterOtghsgintmskFieldEopfmMask  = 0x8000
)

// GetEopfm End of periodic frame interrupt mask
func (r *registerOtghsgintmskType) GetEopfm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldEopfmMask) != 0
}

// SetEopfm End of periodic frame interrupt mask
func (r *registerOtghsgintmskType) SetEopfm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldEopfmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldEopfmMask)
	}
}

const (
	RegisterOtghsgintmskFieldIepintShift = 18
	RegisterOtghsgintmskFieldIepintMask  = 0x40000
)

// GetIepint IN endpoints interrupt mask
func (r *registerOtghsgintmskType) GetIepint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldIepintMask) != 0
}

// SetIepint IN endpoints interrupt mask
func (r *registerOtghsgintmskType) SetIepint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldIepintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldIepintMask)
	}
}

const (
	RegisterOtghsgintmskFieldOepintShift = 19
	RegisterOtghsgintmskFieldOepintMask  = 0x80000
)

// GetOepint OUT endpoints interrupt mask
func (r *registerOtghsgintmskType) GetOepint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldOepintMask) != 0
}

// SetOepint OUT endpoints interrupt mask
func (r *registerOtghsgintmskType) SetOepint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldOepintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldOepintMask)
	}
}

const (
	RegisterOtghsgintmskFieldIisoixfrmShift = 20
	RegisterOtghsgintmskFieldIisoixfrmMask  = 0x100000
)

// GetIisoixfrm Incomplete isochronous IN transfer mask
func (r *registerOtghsgintmskType) GetIisoixfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldIisoixfrmMask) != 0
}

// SetIisoixfrm Incomplete isochronous IN transfer mask
func (r *registerOtghsgintmskType) SetIisoixfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldIisoixfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldIisoixfrmMask)
	}
}

const (
	RegisterOtghsgintmskFieldPxfrmiisooxfrmShift = 21
	RegisterOtghsgintmskFieldPxfrmiisooxfrmMask  = 0x200000
)

// GetPxfrmiisooxfrm Incomplete periodic transfer mask
func (r *registerOtghsgintmskType) GetPxfrmiisooxfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldPxfrmiisooxfrmMask) != 0
}

// SetPxfrmiisooxfrm Incomplete periodic transfer mask
func (r *registerOtghsgintmskType) SetPxfrmiisooxfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldPxfrmiisooxfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldPxfrmiisooxfrmMask)
	}
}

const (
	RegisterOtghsgintmskFieldFsuspmShift = 22
	RegisterOtghsgintmskFieldFsuspmMask  = 0x400000
)

// GetFsuspm Data fetch suspended mask
func (r *registerOtghsgintmskType) GetFsuspm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldFsuspmMask) != 0
}

// SetFsuspm Data fetch suspended mask
func (r *registerOtghsgintmskType) SetFsuspm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldFsuspmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldFsuspmMask)
	}
}

const (
	RegisterOtghsgintmskFieldRstdeShift = 23
	RegisterOtghsgintmskFieldRstdeMask  = 0x800000
)

// GetRstde Reset detected interrupt mask
func (r *registerOtghsgintmskType) GetRstde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldRstdeMask) != 0
}

// SetRstde Reset detected interrupt mask
func (r *registerOtghsgintmskType) SetRstde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldRstdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldRstdeMask)
	}
}

const (
	RegisterOtghsgintmskFieldPrtimShift = 24
	RegisterOtghsgintmskFieldPrtimMask  = 0x1000000
)

// GetPrtim Host port interrupt mask
func (r *registerOtghsgintmskType) GetPrtim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldPrtimMask) != 0
}

const (
	RegisterOtghsgintmskFieldHcimShift = 25
	RegisterOtghsgintmskFieldHcimMask  = 0x2000000
)

// GetHcim Host channels interrupt mask
func (r *registerOtghsgintmskType) GetHcim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldHcimMask) != 0
}

// SetHcim Host channels interrupt mask
func (r *registerOtghsgintmskType) SetHcim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldHcimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldHcimMask)
	}
}

const (
	RegisterOtghsgintmskFieldPtxfemShift = 26
	RegisterOtghsgintmskFieldPtxfemMask  = 0x4000000
)

// GetPtxfem Periodic TxFIFO empty mask
func (r *registerOtghsgintmskType) GetPtxfem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldPtxfemMask) != 0
}

// SetPtxfem Periodic TxFIFO empty mask
func (r *registerOtghsgintmskType) SetPtxfem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldPtxfemMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldPtxfemMask)
	}
}

const (
	RegisterOtghsgintmskFieldLpmintmShift = 27
	RegisterOtghsgintmskFieldLpmintmMask  = 0x8000000
)

// GetLpmintm LPM interrupt mask
func (r *registerOtghsgintmskType) GetLpmintm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldLpmintmMask) != 0
}

// SetLpmintm LPM interrupt mask
func (r *registerOtghsgintmskType) SetLpmintm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldLpmintmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldLpmintmMask)
	}
}

const (
	RegisterOtghsgintmskFieldCidschgmShift = 28
	RegisterOtghsgintmskFieldCidschgmMask  = 0x10000000
)

// GetCidschgm Connector ID status change mask
func (r *registerOtghsgintmskType) GetCidschgm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldCidschgmMask) != 0
}

// SetCidschgm Connector ID status change mask
func (r *registerOtghsgintmskType) SetCidschgm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldCidschgmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldCidschgmMask)
	}
}

const (
	RegisterOtghsgintmskFieldDiscintShift = 29
	RegisterOtghsgintmskFieldDiscintMask  = 0x20000000
)

// GetDiscint Disconnect detected interrupt mask
func (r *registerOtghsgintmskType) GetDiscint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldDiscintMask) != 0
}

// SetDiscint Disconnect detected interrupt mask
func (r *registerOtghsgintmskType) SetDiscint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldDiscintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldDiscintMask)
	}
}

const (
	RegisterOtghsgintmskFieldSrqimShift = 30
	RegisterOtghsgintmskFieldSrqimMask  = 0x40000000
)

// GetSrqim Session request/new session detected interrupt mask
func (r *registerOtghsgintmskType) GetSrqim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldSrqimMask) != 0
}

// SetSrqim Session request/new session detected interrupt mask
func (r *registerOtghsgintmskType) SetSrqim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldSrqimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldSrqimMask)
	}
}

const (
	RegisterOtghsgintmskFieldWuimShift = 31
	RegisterOtghsgintmskFieldWuimMask  = 0x80000000
)

// GetWuim Resume/remote wakeup detected interrupt mask
func (r *registerOtghsgintmskType) GetWuim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgintmskFieldWuimMask) != 0
}

// SetWuim Resume/remote wakeup detected interrupt mask
func (r *registerOtghsgintmskType) SetWuim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgintmskFieldWuimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgintmskFieldWuimMask)
	}
}

// registerOtghsgrxstsrhostType OTG_HS Receive status debug read register (host mode)
type registerOtghsgrxstsrhostType uint32

const (
	RegisterOtghsgrxstsrhostFieldChnumShift = 0
	RegisterOtghsgrxstsrhostFieldChnumMask  = 0xf
)

// GetChnum Channel number
func (r *registerOtghsgrxstsrhostType) GetChnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsrhostFieldChnumMask) >> RegisterOtghsgrxstsrhostFieldChnumShift)
}

// SetChnum Channel number
func (r *registerOtghsgrxstsrhostType) SetChnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsrhostFieldChnumMask)|(uint32(value)<<RegisterOtghsgrxstsrhostFieldChnumShift))
}

const (
	RegisterOtghsgrxstsrhostFieldBcntShift = 4
	RegisterOtghsgrxstsrhostFieldBcntMask  = 0x7ff0
)

// GetBcnt Byte count
func (r *registerOtghsgrxstsrhostType) GetBcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsrhostFieldBcntMask) >> RegisterOtghsgrxstsrhostFieldBcntShift)
}

// SetBcnt Byte count
func (r *registerOtghsgrxstsrhostType) SetBcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsrhostFieldBcntMask)|(uint32(value)<<RegisterOtghsgrxstsrhostFieldBcntShift))
}

const (
	RegisterOtghsgrxstsrhostFieldDpidShift = 15
	RegisterOtghsgrxstsrhostFieldDpidMask  = 0x18000
)

// GetDpid Data PID
func (r *registerOtghsgrxstsrhostType) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsrhostFieldDpidMask) >> RegisterOtghsgrxstsrhostFieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghsgrxstsrhostType) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsrhostFieldDpidMask)|(uint32(value)<<RegisterOtghsgrxstsrhostFieldDpidShift))
}

const (
	RegisterOtghsgrxstsrhostFieldPktstsShift = 17
	RegisterOtghsgrxstsrhostFieldPktstsMask  = 0x1e0000
)

// GetPktsts Packet status
func (r *registerOtghsgrxstsrhostType) GetPktsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsrhostFieldPktstsMask) >> RegisterOtghsgrxstsrhostFieldPktstsShift)
}

// SetPktsts Packet status
func (r *registerOtghsgrxstsrhostType) SetPktsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsrhostFieldPktstsMask)|(uint32(value)<<RegisterOtghsgrxstsrhostFieldPktstsShift))
}

// registerOtghsgrxstsrdeviceType OTG_HS Receive status debug read register (peripheral mode mode)
type registerOtghsgrxstsrdeviceType uint32

const (
	RegisterOtghsgrxstsrdeviceFieldEpnumShift = 0
	RegisterOtghsgrxstsrdeviceFieldEpnumMask  = 0xf
)

// GetEpnum Endpoint number
func (r *registerOtghsgrxstsrdeviceType) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsrdeviceFieldEpnumMask) >> RegisterOtghsgrxstsrdeviceFieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghsgrxstsrdeviceType) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsrdeviceFieldEpnumMask)|(uint32(value)<<RegisterOtghsgrxstsrdeviceFieldEpnumShift))
}

const (
	RegisterOtghsgrxstsrdeviceFieldBcntShift = 4
	RegisterOtghsgrxstsrdeviceFieldBcntMask  = 0x7ff0
)

// GetBcnt Byte count
func (r *registerOtghsgrxstsrdeviceType) GetBcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsrdeviceFieldBcntMask) >> RegisterOtghsgrxstsrdeviceFieldBcntShift)
}

// SetBcnt Byte count
func (r *registerOtghsgrxstsrdeviceType) SetBcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsrdeviceFieldBcntMask)|(uint32(value)<<RegisterOtghsgrxstsrdeviceFieldBcntShift))
}

const (
	RegisterOtghsgrxstsrdeviceFieldDpidShift = 15
	RegisterOtghsgrxstsrdeviceFieldDpidMask  = 0x18000
)

// GetDpid Data PID
func (r *registerOtghsgrxstsrdeviceType) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsrdeviceFieldDpidMask) >> RegisterOtghsgrxstsrdeviceFieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghsgrxstsrdeviceType) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsrdeviceFieldDpidMask)|(uint32(value)<<RegisterOtghsgrxstsrdeviceFieldDpidShift))
}

const (
	RegisterOtghsgrxstsrdeviceFieldPktstsShift = 17
	RegisterOtghsgrxstsrdeviceFieldPktstsMask  = 0x1e0000
)

// GetPktsts Packet status
func (r *registerOtghsgrxstsrdeviceType) GetPktsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsrdeviceFieldPktstsMask) >> RegisterOtghsgrxstsrdeviceFieldPktstsShift)
}

// SetPktsts Packet status
func (r *registerOtghsgrxstsrdeviceType) SetPktsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsrdeviceFieldPktstsMask)|(uint32(value)<<RegisterOtghsgrxstsrdeviceFieldPktstsShift))
}

const (
	RegisterOtghsgrxstsrdeviceFieldFrmnumShift = 21
	RegisterOtghsgrxstsrdeviceFieldFrmnumMask  = 0x1e00000
)

// GetFrmnum Frame number
func (r *registerOtghsgrxstsrdeviceType) GetFrmnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsrdeviceFieldFrmnumMask) >> RegisterOtghsgrxstsrdeviceFieldFrmnumShift)
}

// SetFrmnum Frame number
func (r *registerOtghsgrxstsrdeviceType) SetFrmnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsrdeviceFieldFrmnumMask)|(uint32(value)<<RegisterOtghsgrxstsrdeviceFieldFrmnumShift))
}

// registerOtghsgrxstsphostType OTG_HS status read and pop register (host mode)
type registerOtghsgrxstsphostType uint32

const (
	RegisterOtghsgrxstsphostFieldChnumShift = 0
	RegisterOtghsgrxstsphostFieldChnumMask  = 0xf
)

// GetChnum Channel number
func (r *registerOtghsgrxstsphostType) GetChnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsphostFieldChnumMask) >> RegisterOtghsgrxstsphostFieldChnumShift)
}

// SetChnum Channel number
func (r *registerOtghsgrxstsphostType) SetChnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsphostFieldChnumMask)|(uint32(value)<<RegisterOtghsgrxstsphostFieldChnumShift))
}

const (
	RegisterOtghsgrxstsphostFieldBcntShift = 4
	RegisterOtghsgrxstsphostFieldBcntMask  = 0x7ff0
)

// GetBcnt Byte count
func (r *registerOtghsgrxstsphostType) GetBcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsphostFieldBcntMask) >> RegisterOtghsgrxstsphostFieldBcntShift)
}

// SetBcnt Byte count
func (r *registerOtghsgrxstsphostType) SetBcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsphostFieldBcntMask)|(uint32(value)<<RegisterOtghsgrxstsphostFieldBcntShift))
}

const (
	RegisterOtghsgrxstsphostFieldDpidShift = 15
	RegisterOtghsgrxstsphostFieldDpidMask  = 0x18000
)

// GetDpid Data PID
func (r *registerOtghsgrxstsphostType) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsphostFieldDpidMask) >> RegisterOtghsgrxstsphostFieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghsgrxstsphostType) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsphostFieldDpidMask)|(uint32(value)<<RegisterOtghsgrxstsphostFieldDpidShift))
}

const (
	RegisterOtghsgrxstsphostFieldPktstsShift = 17
	RegisterOtghsgrxstsphostFieldPktstsMask  = 0x1e0000
)

// GetPktsts Packet status
func (r *registerOtghsgrxstsphostType) GetPktsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstsphostFieldPktstsMask) >> RegisterOtghsgrxstsphostFieldPktstsShift)
}

// SetPktsts Packet status
func (r *registerOtghsgrxstsphostType) SetPktsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstsphostFieldPktstsMask)|(uint32(value)<<RegisterOtghsgrxstsphostFieldPktstsShift))
}

// registerOtghsgrxstspdeviceType OTG_HS status read and pop register (peripheral mode)
type registerOtghsgrxstspdeviceType uint32

const (
	RegisterOtghsgrxstspdeviceFieldEpnumShift = 0
	RegisterOtghsgrxstspdeviceFieldEpnumMask  = 0xf
)

// GetEpnum Endpoint number
func (r *registerOtghsgrxstspdeviceType) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstspdeviceFieldEpnumMask) >> RegisterOtghsgrxstspdeviceFieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghsgrxstspdeviceType) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstspdeviceFieldEpnumMask)|(uint32(value)<<RegisterOtghsgrxstspdeviceFieldEpnumShift))
}

const (
	RegisterOtghsgrxstspdeviceFieldBcntShift = 4
	RegisterOtghsgrxstspdeviceFieldBcntMask  = 0x7ff0
)

// GetBcnt Byte count
func (r *registerOtghsgrxstspdeviceType) GetBcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstspdeviceFieldBcntMask) >> RegisterOtghsgrxstspdeviceFieldBcntShift)
}

// SetBcnt Byte count
func (r *registerOtghsgrxstspdeviceType) SetBcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstspdeviceFieldBcntMask)|(uint32(value)<<RegisterOtghsgrxstspdeviceFieldBcntShift))
}

const (
	RegisterOtghsgrxstspdeviceFieldDpidShift = 15
	RegisterOtghsgrxstspdeviceFieldDpidMask  = 0x18000
)

// GetDpid Data PID
func (r *registerOtghsgrxstspdeviceType) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstspdeviceFieldDpidMask) >> RegisterOtghsgrxstspdeviceFieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghsgrxstspdeviceType) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstspdeviceFieldDpidMask)|(uint32(value)<<RegisterOtghsgrxstspdeviceFieldDpidShift))
}

const (
	RegisterOtghsgrxstspdeviceFieldPktstsShift = 17
	RegisterOtghsgrxstspdeviceFieldPktstsMask  = 0x1e0000
)

// GetPktsts Packet status
func (r *registerOtghsgrxstspdeviceType) GetPktsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstspdeviceFieldPktstsMask) >> RegisterOtghsgrxstspdeviceFieldPktstsShift)
}

// SetPktsts Packet status
func (r *registerOtghsgrxstspdeviceType) SetPktsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstspdeviceFieldPktstsMask)|(uint32(value)<<RegisterOtghsgrxstspdeviceFieldPktstsShift))
}

const (
	RegisterOtghsgrxstspdeviceFieldFrmnumShift = 21
	RegisterOtghsgrxstspdeviceFieldFrmnumMask  = 0x1e00000
)

// GetFrmnum Frame number
func (r *registerOtghsgrxstspdeviceType) GetFrmnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxstspdeviceFieldFrmnumMask) >> RegisterOtghsgrxstspdeviceFieldFrmnumShift)
}

// SetFrmnum Frame number
func (r *registerOtghsgrxstspdeviceType) SetFrmnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxstspdeviceFieldFrmnumMask)|(uint32(value)<<RegisterOtghsgrxstspdeviceFieldFrmnumShift))
}

// registerOtghsgrxfsizType OTG_HS Receive FIFO size register
type registerOtghsgrxfsizType uint32

const (
	RegisterOtghsgrxfsizFieldRxfdShift = 0
	RegisterOtghsgrxfsizFieldRxfdMask  = 0xffff
)

// GetRxfd RxFIFO depth
func (r *registerOtghsgrxfsizType) GetRxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgrxfsizFieldRxfdMask) >> RegisterOtghsgrxfsizFieldRxfdShift)
}

// SetRxfd RxFIFO depth
func (r *registerOtghsgrxfsizType) SetRxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgrxfsizFieldRxfdMask)|(uint32(value)<<RegisterOtghsgrxfsizFieldRxfdShift))
}

// registerOtghshnptxfsizhostType OTG_HS nonperiodic transmit FIFO size register (host mode)
type registerOtghshnptxfsizhostType uint32

const (
	RegisterOtghshnptxfsizhostFieldNptxfsaShift = 0
	RegisterOtghshnptxfsizhostFieldNptxfsaMask  = 0xffff
)

// GetNptxfsa Nonperiodic transmit RAM start address
func (r *registerOtghshnptxfsizhostType) GetNptxfsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshnptxfsizhostFieldNptxfsaMask) >> RegisterOtghshnptxfsizhostFieldNptxfsaShift)
}

// SetNptxfsa Nonperiodic transmit RAM start address
func (r *registerOtghshnptxfsizhostType) SetNptxfsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshnptxfsizhostFieldNptxfsaMask)|(uint32(value)<<RegisterOtghshnptxfsizhostFieldNptxfsaShift))
}

const (
	RegisterOtghshnptxfsizhostFieldNptxfdShift = 16
	RegisterOtghshnptxfsizhostFieldNptxfdMask  = 0xffff0000
)

// GetNptxfd Nonperiodic TxFIFO depth
func (r *registerOtghshnptxfsizhostType) GetNptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshnptxfsizhostFieldNptxfdMask) >> RegisterOtghshnptxfsizhostFieldNptxfdShift)
}

// SetNptxfd Nonperiodic TxFIFO depth
func (r *registerOtghshnptxfsizhostType) SetNptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshnptxfsizhostFieldNptxfdMask)|(uint32(value)<<RegisterOtghshnptxfsizhostFieldNptxfdShift))
}

// registerOtghsdieptxf0deviceType Endpoint 0 transmit FIFO size (peripheral mode)
type registerOtghsdieptxf0deviceType uint32

const (
	RegisterOtghsdieptxf0deviceFieldTx0fsaShift = 0
	RegisterOtghsdieptxf0deviceFieldTx0fsaMask  = 0xffff
)

// GetTx0fsa Endpoint 0 transmit RAM start address
func (r *registerOtghsdieptxf0deviceType) GetTx0fsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf0deviceFieldTx0fsaMask) >> RegisterOtghsdieptxf0deviceFieldTx0fsaShift)
}

// SetTx0fsa Endpoint 0 transmit RAM start address
func (r *registerOtghsdieptxf0deviceType) SetTx0fsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf0deviceFieldTx0fsaMask)|(uint32(value)<<RegisterOtghsdieptxf0deviceFieldTx0fsaShift))
}

const (
	RegisterOtghsdieptxf0deviceFieldTx0fdShift = 16
	RegisterOtghsdieptxf0deviceFieldTx0fdMask  = 0xffff0000
)

// GetTx0fd Endpoint 0 TxFIFO depth
func (r *registerOtghsdieptxf0deviceType) GetTx0fd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf0deviceFieldTx0fdMask) >> RegisterOtghsdieptxf0deviceFieldTx0fdShift)
}

// SetTx0fd Endpoint 0 TxFIFO depth
func (r *registerOtghsdieptxf0deviceType) SetTx0fd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf0deviceFieldTx0fdMask)|(uint32(value)<<RegisterOtghsdieptxf0deviceFieldTx0fdShift))
}

// registerOtghsgnptxstsType OTG_HS nonperiodic transmit FIFO/queue status register
type registerOtghsgnptxstsType uint32

const (
	RegisterOtghsgnptxstsFieldNptxfsavShift = 0
	RegisterOtghsgnptxstsFieldNptxfsavMask  = 0xffff
)

// GetNptxfsav Nonperiodic TxFIFO space available
func (r *registerOtghsgnptxstsType) GetNptxfsav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgnptxstsFieldNptxfsavMask) >> RegisterOtghsgnptxstsFieldNptxfsavShift)
}

// SetNptxfsav Nonperiodic TxFIFO space available
func (r *registerOtghsgnptxstsType) SetNptxfsav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgnptxstsFieldNptxfsavMask)|(uint32(value)<<RegisterOtghsgnptxstsFieldNptxfsavShift))
}

const (
	RegisterOtghsgnptxstsFieldNptqxsavShift = 16
	RegisterOtghsgnptxstsFieldNptqxsavMask  = 0xff0000
)

// GetNptqxsav Nonperiodic transmit request queue space available
func (r *registerOtghsgnptxstsType) GetNptqxsav() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgnptxstsFieldNptqxsavMask) >> RegisterOtghsgnptxstsFieldNptqxsavShift)
}

// SetNptqxsav Nonperiodic transmit request queue space available
func (r *registerOtghsgnptxstsType) SetNptqxsav(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgnptxstsFieldNptqxsavMask)|(uint32(value)<<RegisterOtghsgnptxstsFieldNptqxsavShift))
}

const (
	RegisterOtghsgnptxstsFieldNptxqtopShift = 24
	RegisterOtghsgnptxstsFieldNptxqtopMask  = 0x7f000000
)

// GetNptxqtop Top of the nonperiodic transmit request queue
func (r *registerOtghsgnptxstsType) GetNptxqtop() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgnptxstsFieldNptxqtopMask) >> RegisterOtghsgnptxstsFieldNptxqtopShift)
}

// SetNptxqtop Top of the nonperiodic transmit request queue
func (r *registerOtghsgnptxstsType) SetNptxqtop(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgnptxstsFieldNptxqtopMask)|(uint32(value)<<RegisterOtghsgnptxstsFieldNptxqtopShift))
}

// registerOtghsgccfgType OTG_HS general core configuration register
type registerOtghsgccfgType uint32

const (
	RegisterOtghsgccfgFieldDcdetShift = 0
	RegisterOtghsgccfgFieldDcdetMask  = 0x1
)

// GetDcdet Data contact detection (DCD) status
func (r *registerOtghsgccfgType) GetDcdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgccfgFieldDcdetMask) != 0
}

// SetDcdet Data contact detection (DCD) status
func (r *registerOtghsgccfgType) SetDcdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgccfgFieldDcdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgccfgFieldDcdetMask)
	}
}

const (
	RegisterOtghsgccfgFieldPdetShift = 1
	RegisterOtghsgccfgFieldPdetMask  = 0x2
)

// GetPdet Primary detection (PD) status
func (r *registerOtghsgccfgType) GetPdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgccfgFieldPdetMask) != 0
}

// SetPdet Primary detection (PD) status
func (r *registerOtghsgccfgType) SetPdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgccfgFieldPdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgccfgFieldPdetMask)
	}
}

const (
	RegisterOtghsgccfgFieldSdetShift = 2
	RegisterOtghsgccfgFieldSdetMask  = 0x4
)

// GetSdet Secondary detection (SD) status
func (r *registerOtghsgccfgType) GetSdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgccfgFieldSdetMask) != 0
}

// SetSdet Secondary detection (SD) status
func (r *registerOtghsgccfgType) SetSdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgccfgFieldSdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgccfgFieldSdetMask)
	}
}

const (
	RegisterOtghsgccfgFieldPs2detShift = 3
	RegisterOtghsgccfgFieldPs2detMask  = 0x8
)

// GetPs2det DM pull-up detection status
func (r *registerOtghsgccfgType) GetPs2det() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgccfgFieldPs2detMask) != 0
}

// SetPs2det DM pull-up detection status
func (r *registerOtghsgccfgType) SetPs2det(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgccfgFieldPs2detMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgccfgFieldPs2detMask)
	}
}

const (
	RegisterOtghsgccfgFieldPwrdwnShift = 16
	RegisterOtghsgccfgFieldPwrdwnMask  = 0x10000
)

// GetPwrdwn Power down
func (r *registerOtghsgccfgType) GetPwrdwn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgccfgFieldPwrdwnMask) != 0
}

// SetPwrdwn Power down
func (r *registerOtghsgccfgType) SetPwrdwn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgccfgFieldPwrdwnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgccfgFieldPwrdwnMask)
	}
}

const (
	RegisterOtghsgccfgFieldBcdenShift = 17
	RegisterOtghsgccfgFieldBcdenMask  = 0x20000
)

// GetBcden Battery charging detector (BCD) enable
func (r *registerOtghsgccfgType) GetBcden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgccfgFieldBcdenMask) != 0
}

// SetBcden Battery charging detector (BCD) enable
func (r *registerOtghsgccfgType) SetBcden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgccfgFieldBcdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgccfgFieldBcdenMask)
	}
}

const (
	RegisterOtghsgccfgFieldDcdenShift = 18
	RegisterOtghsgccfgFieldDcdenMask  = 0x40000
)

// GetDcden Data contact detection (DCD) mode enable
func (r *registerOtghsgccfgType) GetDcden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgccfgFieldDcdenMask) != 0
}

// SetDcden Data contact detection (DCD) mode enable
func (r *registerOtghsgccfgType) SetDcden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgccfgFieldDcdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgccfgFieldDcdenMask)
	}
}

const (
	RegisterOtghsgccfgFieldPdenShift = 19
	RegisterOtghsgccfgFieldPdenMask  = 0x80000
)

// GetPden Primary detection (PD) mode enable
func (r *registerOtghsgccfgType) GetPden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgccfgFieldPdenMask) != 0
}

// SetPden Primary detection (PD) mode enable
func (r *registerOtghsgccfgType) SetPden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgccfgFieldPdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgccfgFieldPdenMask)
	}
}

const (
	RegisterOtghsgccfgFieldSdenShift = 20
	RegisterOtghsgccfgFieldSdenMask  = 0x100000
)

// GetSden Secondary detection (SD) mode enable
func (r *registerOtghsgccfgType) GetSden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgccfgFieldSdenMask) != 0
}

// SetSden Secondary detection (SD) mode enable
func (r *registerOtghsgccfgType) SetSden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgccfgFieldSdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgccfgFieldSdenMask)
	}
}

const (
	RegisterOtghsgccfgFieldVbdenShift = 21
	RegisterOtghsgccfgFieldVbdenMask  = 0x200000
)

// GetVbden USB VBUS detection enable
func (r *registerOtghsgccfgType) GetVbden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsgccfgFieldVbdenMask) != 0
}

// SetVbden USB VBUS detection enable
func (r *registerOtghsgccfgType) SetVbden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsgccfgFieldVbdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsgccfgFieldVbdenMask)
	}
}

// registerOtghscidType OTG_HS core ID register
type registerOtghscidType uint32

const (
	RegisterOtghscidFieldProductidShift = 0
	RegisterOtghscidFieldProductidMask  = 0xffffffff
)

// GetProductid Product ID field
func (r *registerOtghscidType) GetProductid() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghscidFieldProductidMask) >> RegisterOtghscidFieldProductidShift)
}

// SetProductid Product ID field
func (r *registerOtghscidType) SetProductid(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghscidFieldProductidMask)|(uint32(value)<<RegisterOtghscidFieldProductidShift))
}

// registerOtghsglpmcfgType OTG core LPM configuration register
type registerOtghsglpmcfgType uint32

const (
	RegisterOtghsglpmcfgFieldLpmenShift = 0
	RegisterOtghsglpmcfgFieldLpmenMask  = 0x1
)

// GetLpmen LPM support enable
func (r *registerOtghsglpmcfgType) GetLpmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldLpmenMask) != 0
}

// SetLpmen LPM support enable
func (r *registerOtghsglpmcfgType) SetLpmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsglpmcfgFieldLpmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsglpmcfgFieldLpmenMask)
	}
}

const (
	RegisterOtghsglpmcfgFieldLpmackShift = 1
	RegisterOtghsglpmcfgFieldLpmackMask  = 0x2
)

// GetLpmack LPM token acknowledge enable
func (r *registerOtghsglpmcfgType) GetLpmack() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldLpmackMask) != 0
}

// SetLpmack LPM token acknowledge enable
func (r *registerOtghsglpmcfgType) SetLpmack(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsglpmcfgFieldLpmackMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsglpmcfgFieldLpmackMask)
	}
}

const (
	RegisterOtghsglpmcfgFieldBeslShift = 2
	RegisterOtghsglpmcfgFieldBeslMask  = 0x3c
)

// GetBesl Best effort service latency
func (r *registerOtghsglpmcfgType) GetBesl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldBeslMask) >> RegisterOtghsglpmcfgFieldBeslShift)
}

const (
	RegisterOtghsglpmcfgFieldRemwakeShift = 6
	RegisterOtghsglpmcfgFieldRemwakeMask  = 0x40
)

// GetRemwake bRemoteWake value
func (r *registerOtghsglpmcfgType) GetRemwake() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldRemwakeMask) != 0
}

const (
	RegisterOtghsglpmcfgFieldL1ssenShift = 7
	RegisterOtghsglpmcfgFieldL1ssenMask  = 0x80
)

// GetL1ssen L1 Shallow Sleep enable
func (r *registerOtghsglpmcfgType) GetL1ssen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldL1ssenMask) != 0
}

// SetL1ssen L1 Shallow Sleep enable
func (r *registerOtghsglpmcfgType) SetL1ssen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsglpmcfgFieldL1ssenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsglpmcfgFieldL1ssenMask)
	}
}

const (
	RegisterOtghsglpmcfgFieldBeslthrsShift = 8
	RegisterOtghsglpmcfgFieldBeslthrsMask  = 0xf00
)

// GetBeslthrs BESL threshold
func (r *registerOtghsglpmcfgType) GetBeslthrs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldBeslthrsMask) >> RegisterOtghsglpmcfgFieldBeslthrsShift)
}

// SetBeslthrs BESL threshold
func (r *registerOtghsglpmcfgType) SetBeslthrs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsglpmcfgFieldBeslthrsMask)|(uint32(value)<<RegisterOtghsglpmcfgFieldBeslthrsShift))
}

const (
	RegisterOtghsglpmcfgFieldL1dsenShift = 12
	RegisterOtghsglpmcfgFieldL1dsenMask  = 0x1000
)

// GetL1dsen L1 deep sleep enable
func (r *registerOtghsglpmcfgType) GetL1dsen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldL1dsenMask) != 0
}

// SetL1dsen L1 deep sleep enable
func (r *registerOtghsglpmcfgType) SetL1dsen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsglpmcfgFieldL1dsenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsglpmcfgFieldL1dsenMask)
	}
}

const (
	RegisterOtghsglpmcfgFieldLpmrstShift = 13
	RegisterOtghsglpmcfgFieldLpmrstMask  = 0x6000
)

// GetLpmrst LPM response
func (r *registerOtghsglpmcfgType) GetLpmrst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldLpmrstMask) >> RegisterOtghsglpmcfgFieldLpmrstShift)
}

const (
	RegisterOtghsglpmcfgFieldSlpstsShift = 15
	RegisterOtghsglpmcfgFieldSlpstsMask  = 0x8000
)

// GetSlpsts Port sleep status
func (r *registerOtghsglpmcfgType) GetSlpsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldSlpstsMask) != 0
}

const (
	RegisterOtghsglpmcfgFieldL1rsmokShift = 16
	RegisterOtghsglpmcfgFieldL1rsmokMask  = 0x10000
)

// GetL1rsmok Sleep State Resume OK
func (r *registerOtghsglpmcfgType) GetL1rsmok() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldL1rsmokMask) != 0
}

const (
	RegisterOtghsglpmcfgFieldLpmchidxShift = 17
	RegisterOtghsglpmcfgFieldLpmchidxMask  = 0x1e0000
)

// GetLpmchidx LPM Channel Index
func (r *registerOtghsglpmcfgType) GetLpmchidx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldLpmchidxMask) >> RegisterOtghsglpmcfgFieldLpmchidxShift)
}

// SetLpmchidx LPM Channel Index
func (r *registerOtghsglpmcfgType) SetLpmchidx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsglpmcfgFieldLpmchidxMask)|(uint32(value)<<RegisterOtghsglpmcfgFieldLpmchidxShift))
}

const (
	RegisterOtghsglpmcfgFieldLpmrcntShift = 21
	RegisterOtghsglpmcfgFieldLpmrcntMask  = 0xe00000
)

// GetLpmrcnt LPM retry count
func (r *registerOtghsglpmcfgType) GetLpmrcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldLpmrcntMask) >> RegisterOtghsglpmcfgFieldLpmrcntShift)
}

// SetLpmrcnt LPM retry count
func (r *registerOtghsglpmcfgType) SetLpmrcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsglpmcfgFieldLpmrcntMask)|(uint32(value)<<RegisterOtghsglpmcfgFieldLpmrcntShift))
}

const (
	RegisterOtghsglpmcfgFieldSndlpmShift = 24
	RegisterOtghsglpmcfgFieldSndlpmMask  = 0x1000000
)

// GetSndlpm Send LPM transaction
func (r *registerOtghsglpmcfgType) GetSndlpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldSndlpmMask) != 0
}

// SetSndlpm Send LPM transaction
func (r *registerOtghsglpmcfgType) SetSndlpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsglpmcfgFieldSndlpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsglpmcfgFieldSndlpmMask)
	}
}

const (
	RegisterOtghsglpmcfgFieldLpmrcntstsShift = 25
	RegisterOtghsglpmcfgFieldLpmrcntstsMask  = 0xe000000
)

// GetLpmrcntsts LPM retry count status
func (r *registerOtghsglpmcfgType) GetLpmrcntsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldLpmrcntstsMask) >> RegisterOtghsglpmcfgFieldLpmrcntstsShift)
}

const (
	RegisterOtghsglpmcfgFieldEnbeslShift = 28
	RegisterOtghsglpmcfgFieldEnbeslMask  = 0x10000000
)

// GetEnbesl Enable best effort service latency
func (r *registerOtghsglpmcfgType) GetEnbesl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghsglpmcfgFieldEnbeslMask) != 0
}

// SetEnbesl Enable best effort service latency
func (r *registerOtghsglpmcfgType) SetEnbesl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghsglpmcfgFieldEnbeslMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghsglpmcfgFieldEnbeslMask)
	}
}

// registerOtghshptxfsizType OTG_HS Host periodic transmit FIFO size register
type registerOtghshptxfsizType uint32

const (
	RegisterOtghshptxfsizFieldPtxsaShift = 0
	RegisterOtghshptxfsizFieldPtxsaMask  = 0xffff
)

// GetPtxsa Host periodic TxFIFO start address
func (r *registerOtghshptxfsizType) GetPtxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshptxfsizFieldPtxsaMask) >> RegisterOtghshptxfsizFieldPtxsaShift)
}

// SetPtxsa Host periodic TxFIFO start address
func (r *registerOtghshptxfsizType) SetPtxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshptxfsizFieldPtxsaMask)|(uint32(value)<<RegisterOtghshptxfsizFieldPtxsaShift))
}

const (
	RegisterOtghshptxfsizFieldPtxfdShift = 16
	RegisterOtghshptxfsizFieldPtxfdMask  = 0xffff0000
)

// GetPtxfd Host periodic TxFIFO depth
func (r *registerOtghshptxfsizType) GetPtxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshptxfsizFieldPtxfdMask) >> RegisterOtghshptxfsizFieldPtxfdShift)
}

// SetPtxfd Host periodic TxFIFO depth
func (r *registerOtghshptxfsizType) SetPtxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshptxfsizFieldPtxfdMask)|(uint32(value)<<RegisterOtghshptxfsizFieldPtxfdShift))
}

// registerOtghsdieptxf1Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtghsdieptxf1Type uint32

const (
	RegisterOtghsdieptxf1FieldIneptxsaShift = 0
	RegisterOtghsdieptxf1FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf1Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf1FieldIneptxsaMask) >> RegisterOtghsdieptxf1FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf1Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf1FieldIneptxsaMask)|(uint32(value)<<RegisterOtghsdieptxf1FieldIneptxsaShift))
}

const (
	RegisterOtghsdieptxf1FieldIneptxfdShift = 16
	RegisterOtghsdieptxf1FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf1Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf1FieldIneptxfdMask) >> RegisterOtghsdieptxf1FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf1Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf1FieldIneptxfdMask)|(uint32(value)<<RegisterOtghsdieptxf1FieldIneptxfdShift))
}

// registerOtghsdieptxf2Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtghsdieptxf2Type uint32

const (
	RegisterOtghsdieptxf2FieldIneptxsaShift = 0
	RegisterOtghsdieptxf2FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf2Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf2FieldIneptxsaMask) >> RegisterOtghsdieptxf2FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf2Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf2FieldIneptxsaMask)|(uint32(value)<<RegisterOtghsdieptxf2FieldIneptxsaShift))
}

const (
	RegisterOtghsdieptxf2FieldIneptxfdShift = 16
	RegisterOtghsdieptxf2FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf2Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf2FieldIneptxfdMask) >> RegisterOtghsdieptxf2FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf2Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf2FieldIneptxfdMask)|(uint32(value)<<RegisterOtghsdieptxf2FieldIneptxfdShift))
}

// registerOtghsdieptxf3Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtghsdieptxf3Type uint32

const (
	RegisterOtghsdieptxf3FieldIneptxsaShift = 0
	RegisterOtghsdieptxf3FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf3Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf3FieldIneptxsaMask) >> RegisterOtghsdieptxf3FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf3Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf3FieldIneptxsaMask)|(uint32(value)<<RegisterOtghsdieptxf3FieldIneptxsaShift))
}

const (
	RegisterOtghsdieptxf3FieldIneptxfdShift = 16
	RegisterOtghsdieptxf3FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf3Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf3FieldIneptxfdMask) >> RegisterOtghsdieptxf3FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf3Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf3FieldIneptxfdMask)|(uint32(value)<<RegisterOtghsdieptxf3FieldIneptxfdShift))
}

// registerOtghsdieptxf4Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtghsdieptxf4Type uint32

const (
	RegisterOtghsdieptxf4FieldIneptxsaShift = 0
	RegisterOtghsdieptxf4FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf4Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf4FieldIneptxsaMask) >> RegisterOtghsdieptxf4FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf4Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf4FieldIneptxsaMask)|(uint32(value)<<RegisterOtghsdieptxf4FieldIneptxsaShift))
}

const (
	RegisterOtghsdieptxf4FieldIneptxfdShift = 16
	RegisterOtghsdieptxf4FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf4Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf4FieldIneptxfdMask) >> RegisterOtghsdieptxf4FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf4Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf4FieldIneptxfdMask)|(uint32(value)<<RegisterOtghsdieptxf4FieldIneptxfdShift))
}

// registerOtghsdieptxf5Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtghsdieptxf5Type uint32

const (
	RegisterOtghsdieptxf5FieldIneptxsaShift = 0
	RegisterOtghsdieptxf5FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf5Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf5FieldIneptxsaMask) >> RegisterOtghsdieptxf5FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf5Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf5FieldIneptxsaMask)|(uint32(value)<<RegisterOtghsdieptxf5FieldIneptxsaShift))
}

const (
	RegisterOtghsdieptxf5FieldIneptxfdShift = 16
	RegisterOtghsdieptxf5FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf5Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf5FieldIneptxfdMask) >> RegisterOtghsdieptxf5FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf5Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf5FieldIneptxfdMask)|(uint32(value)<<RegisterOtghsdieptxf5FieldIneptxfdShift))
}

// registerOtghsdieptxf6Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtghsdieptxf6Type uint32

const (
	RegisterOtghsdieptxf6FieldIneptxsaShift = 0
	RegisterOtghsdieptxf6FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf6Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf6FieldIneptxsaMask) >> RegisterOtghsdieptxf6FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf6Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf6FieldIneptxsaMask)|(uint32(value)<<RegisterOtghsdieptxf6FieldIneptxsaShift))
}

const (
	RegisterOtghsdieptxf6FieldIneptxfdShift = 16
	RegisterOtghsdieptxf6FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf6Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf6FieldIneptxfdMask) >> RegisterOtghsdieptxf6FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf6Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf6FieldIneptxfdMask)|(uint32(value)<<RegisterOtghsdieptxf6FieldIneptxfdShift))
}

// registerOtghsdieptxf7Type OTG_HS device IN endpoint transmit FIFO size register
type registerOtghsdieptxf7Type uint32

const (
	RegisterOtghsdieptxf7FieldIneptxsaShift = 0
	RegisterOtghsdieptxf7FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf7Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf7FieldIneptxsaMask) >> RegisterOtghsdieptxf7FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address
func (r *registerOtghsdieptxf7Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf7FieldIneptxsaMask)|(uint32(value)<<RegisterOtghsdieptxf7FieldIneptxsaShift))
}

const (
	RegisterOtghsdieptxf7FieldIneptxfdShift = 16
	RegisterOtghsdieptxf7FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf7Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghsdieptxf7FieldIneptxfdMask) >> RegisterOtghsdieptxf7FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint TxFIFO depth
func (r *registerOtghsdieptxf7Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghsdieptxf7FieldIneptxfdMask)|(uint32(value)<<RegisterOtghsdieptxf7FieldIneptxfdShift))
}

type registerOtgdieptxf8Type uint32

const (
	RegisterOtgdieptxf8FieldIneptxsaShift = 0
	RegisterOtgdieptxf8FieldIneptxsaMask  = 0xffff
)

// GetIneptxsa IN endpoint FIFOx transmit RAM start address This field contains the memory start address for IN endpoint transmit FIFOx. The address must be aligned with a 32-bit memory location.
func (r *registerOtgdieptxf8Type) GetIneptxsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtgdieptxf8FieldIneptxsaMask) >> RegisterOtgdieptxf8FieldIneptxsaShift)
}

// SetIneptxsa IN endpoint FIFOx transmit RAM start address This field contains the memory start address for IN endpoint transmit FIFOx. The address must be aligned with a 32-bit memory location.
func (r *registerOtgdieptxf8Type) SetIneptxsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtgdieptxf8FieldIneptxsaMask)|(uint32(value)<<RegisterOtgdieptxf8FieldIneptxsaShift))
}

const (
	RegisterOtgdieptxf8FieldIneptxfdShift = 16
	RegisterOtgdieptxf8FieldIneptxfdMask  = 0xffff0000
)

// GetIneptxfd IN endpoint Tx FIFO depth This value is in terms of 32-bit words. Minimum value is 16
func (r *registerOtgdieptxf8Type) GetIneptxfd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtgdieptxf8FieldIneptxfdMask) >> RegisterOtgdieptxf8FieldIneptxfdShift)
}

// SetIneptxfd IN endpoint Tx FIFO depth This value is in terms of 32-bit words. Minimum value is 16
func (r *registerOtgdieptxf8Type) SetIneptxfd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtgdieptxf8FieldIneptxfdMask)|(uint32(value)<<RegisterOtgdieptxf8FieldIneptxfdShift))
}
