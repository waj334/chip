package rtc

import (
	"unsafe"
	"volatile"
)

var (
	Rtc = (*_rtc)(unsafe.Pointer(uintptr(0x58004000)))
)

type _rtc struct {
	Rtc_tr       registerRtc_trType
	Rtc_dr       registerRtc_drType
	Rtc_cr       registerRtc_crType
	Rtc_isr      registerRtc_isrType
	Rtc_prer     registerRtc_prerType
	Rtc_wutr     registerRtc_wutrType
	_            [4]byte
	Rtc_alrmar   registerRtc_alrmarType
	Rtc_alrmbr   registerRtc_alrmbrType
	Rtc_wpr      registerRtc_wprType
	Rtc_ssr      registerRtc_ssrType
	Rtc_shiftr   registerRtc_shiftrType
	Rtc_tstr     registerRtc_tstrType
	Rtc_tsdr     registerRtc_tsdrType
	Rtc_tsssr    registerRtc_tsssrType
	Rtc_calr     registerRtc_calrType
	Rtc_tampcr   registerRtc_tampcrType
	Rtc_alrmassr registerRtc_alrmassrType
	Rtc_alrmbssr registerRtc_alrmbssrType
	Rtc_or       registerRtc_orType
	Rtc_bkp0r    registerRtc_bkp0rType
	Rtc_bkp1r    registerRtc_bkp1rType
	Rtc_bkp2r    registerRtc_bkp2rType
	Rtc_bkp3r    registerRtc_bkp3rType
	Rtc_bkp4r    registerRtc_bkp4rType
	Rtc_bkp5r    registerRtc_bkp5rType
	Rtc_bkp6r    registerRtc_bkp6rType
	Rtc_bkp7r    registerRtc_bkp7rType
	Rtc_bkp8r    registerRtc_bkp8rType
	Rtc_bkp9r    registerRtc_bkp9rType
	Rtc_bkp10r   registerRtc_bkp10rType
	Rtc_bkp11r   registerRtc_bkp11rType
	Rtc_bkp12r   registerRtc_bkp12rType
	Rtc_bkp13r   registerRtc_bkp13rType
	Rtc_bkp14r   registerRtc_bkp14rType
	Rtc_bkp15r   registerRtc_bkp15rType
	Rtc_bkp16r   registerRtc_bkp16rType
	Rtc_bkp17r   registerRtc_bkp17rType
	Rtc_bkp18r   registerRtc_bkp18rType
	Rtc_bkp19r   registerRtc_bkp19rType
	Rtc_bkp20r   registerRtc_bkp20rType
	Rtc_bkp21r   registerRtc_bkp21rType
	Rtc_bkp22r   registerRtc_bkp22rType
	Rtc_bkp23r   registerRtc_bkp23rType
	Rtc_bkp24r   registerRtc_bkp24rType
	Rtc_bkp25r   registerRtc_bkp25rType
	Rtc_bkp26r   registerRtc_bkp26rType
	Rtc_bkp27r   registerRtc_bkp27rType
	Rtc_bkp28r   registerRtc_bkp28rType
	Rtc_bkp29r   registerRtc_bkp29rType
	Rtc_bkp30r   registerRtc_bkp30rType
	Rtc_bkp31r   registerRtc_bkp31rType
}

// registerRtc_trType The RTC_TR is the calendar time shadow register. This register must be written in initialization mode only. Refer to Calendar initialization and configuration on page9 and Reading the calendar on page10.This register is write protected. The write access procedure is described in RTC register write protection on page9.
type registerRtc_trType uint32

const (
	RegisterRtc_trFieldSuShift = 0
	RegisterRtc_trFieldSuMask  = 0xf
)

// GetSu Second units in BCD format
func (r *registerRtc_trType) GetSu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_trFieldSuMask) >> RegisterRtc_trFieldSuShift)
}

// SetSu Second units in BCD format
func (r *registerRtc_trType) SetSu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_trFieldSuMask)|(uint32(value)<<RegisterRtc_trFieldSuShift))
}

const (
	RegisterRtc_trFieldStShift = 4
	RegisterRtc_trFieldStMask  = 0x70
)

// GetSt Second tens in BCD format
func (r *registerRtc_trType) GetSt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_trFieldStMask) >> RegisterRtc_trFieldStShift)
}

// SetSt Second tens in BCD format
func (r *registerRtc_trType) SetSt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_trFieldStMask)|(uint32(value)<<RegisterRtc_trFieldStShift))
}

const (
	RegisterRtc_trFieldMnuShift = 8
	RegisterRtc_trFieldMnuMask  = 0xf00
)

// GetMnu Minute units in BCD format
func (r *registerRtc_trType) GetMnu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_trFieldMnuMask) >> RegisterRtc_trFieldMnuShift)
}

// SetMnu Minute units in BCD format
func (r *registerRtc_trType) SetMnu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_trFieldMnuMask)|(uint32(value)<<RegisterRtc_trFieldMnuShift))
}

const (
	RegisterRtc_trFieldMntShift = 12
	RegisterRtc_trFieldMntMask  = 0x7000
)

// GetMnt Minute tens in BCD format
func (r *registerRtc_trType) GetMnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_trFieldMntMask) >> RegisterRtc_trFieldMntShift)
}

// SetMnt Minute tens in BCD format
func (r *registerRtc_trType) SetMnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_trFieldMntMask)|(uint32(value)<<RegisterRtc_trFieldMntShift))
}

const (
	RegisterRtc_trFieldHuShift = 16
	RegisterRtc_trFieldHuMask  = 0xf0000
)

// GetHu Hour units in BCD format
func (r *registerRtc_trType) GetHu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_trFieldHuMask) >> RegisterRtc_trFieldHuShift)
}

// SetHu Hour units in BCD format
func (r *registerRtc_trType) SetHu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_trFieldHuMask)|(uint32(value)<<RegisterRtc_trFieldHuShift))
}

const (
	RegisterRtc_trFieldHtShift = 20
	RegisterRtc_trFieldHtMask  = 0x300000
)

// GetHt Hour tens in BCD format
func (r *registerRtc_trType) GetHt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_trFieldHtMask) >> RegisterRtc_trFieldHtShift)
}

// SetHt Hour tens in BCD format
func (r *registerRtc_trType) SetHt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_trFieldHtMask)|(uint32(value)<<RegisterRtc_trFieldHtShift))
}

const (
	RegisterRtc_trFieldPmShift = 22
	RegisterRtc_trFieldPmMask  = 0x400000
)

// GetPm AM/PM notation
func (r *registerRtc_trType) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_trFieldPmMask) != 0
}

// SetPm AM/PM notation
func (r *registerRtc_trType) SetPm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_trFieldPmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_trFieldPmMask)
	}
}

// registerRtc_drType The RTC_DR is the calendar date shadow register. This register must be written in initialization mode only. Refer to Calendar initialization and configuration on page9 and Reading the calendar on page10.This register is write protected. The write access procedure is described in RTC register write protection on page9.
type registerRtc_drType uint32

const (
	RegisterRtc_drFieldDuShift = 0
	RegisterRtc_drFieldDuMask  = 0xf
)

// GetDu Date units in BCD format
func (r *registerRtc_drType) GetDu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_drFieldDuMask) >> RegisterRtc_drFieldDuShift)
}

// SetDu Date units in BCD format
func (r *registerRtc_drType) SetDu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_drFieldDuMask)|(uint32(value)<<RegisterRtc_drFieldDuShift))
}

const (
	RegisterRtc_drFieldDtShift = 4
	RegisterRtc_drFieldDtMask  = 0x30
)

// GetDt Date tens in BCD format
func (r *registerRtc_drType) GetDt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_drFieldDtMask) >> RegisterRtc_drFieldDtShift)
}

// SetDt Date tens in BCD format
func (r *registerRtc_drType) SetDt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_drFieldDtMask)|(uint32(value)<<RegisterRtc_drFieldDtShift))
}

const (
	RegisterRtc_drFieldMuShift = 8
	RegisterRtc_drFieldMuMask  = 0xf00
)

// GetMu Month units in BCD format
func (r *registerRtc_drType) GetMu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_drFieldMuMask) >> RegisterRtc_drFieldMuShift)
}

// SetMu Month units in BCD format
func (r *registerRtc_drType) SetMu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_drFieldMuMask)|(uint32(value)<<RegisterRtc_drFieldMuShift))
}

const (
	RegisterRtc_drFieldMtShift = 12
	RegisterRtc_drFieldMtMask  = 0x1000
)

// GetMt Month tens in BCD format
func (r *registerRtc_drType) GetMt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_drFieldMtMask) != 0
}

// SetMt Month tens in BCD format
func (r *registerRtc_drType) SetMt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_drFieldMtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_drFieldMtMask)
	}
}

const (
	RegisterRtc_drFieldWduShift = 13
	RegisterRtc_drFieldWduMask  = 0xe000
)

// GetWdu Week day units
func (r *registerRtc_drType) GetWdu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_drFieldWduMask) >> RegisterRtc_drFieldWduShift)
}

// SetWdu Week day units
func (r *registerRtc_drType) SetWdu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_drFieldWduMask)|(uint32(value)<<RegisterRtc_drFieldWduShift))
}

const (
	RegisterRtc_drFieldYuShift = 16
	RegisterRtc_drFieldYuMask  = 0xf0000
)

// GetYu Year units in BCD format
func (r *registerRtc_drType) GetYu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_drFieldYuMask) >> RegisterRtc_drFieldYuShift)
}

// SetYu Year units in BCD format
func (r *registerRtc_drType) SetYu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_drFieldYuMask)|(uint32(value)<<RegisterRtc_drFieldYuShift))
}

const (
	RegisterRtc_drFieldYtShift = 20
	RegisterRtc_drFieldYtMask  = 0xf00000
)

// GetYt Year tens in BCD format
func (r *registerRtc_drType) GetYt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_drFieldYtMask) >> RegisterRtc_drFieldYtShift)
}

// SetYt Year tens in BCD format
func (r *registerRtc_drType) SetYt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_drFieldYtMask)|(uint32(value)<<RegisterRtc_drFieldYtShift))
}

// registerRtc_crType RTC control register
type registerRtc_crType uint32

const (
	RegisterRtc_crFieldWuckselShift = 0
	RegisterRtc_crFieldWuckselMask  = 0x7
)

// GetWucksel Wakeup clock selection
func (r *registerRtc_crType) GetWucksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldWuckselMask) >> RegisterRtc_crFieldWuckselShift)
}

// SetWucksel Wakeup clock selection
func (r *registerRtc_crType) SetWucksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldWuckselMask)|(uint32(value)<<RegisterRtc_crFieldWuckselShift))
}

const (
	RegisterRtc_crFieldTsedgeShift = 3
	RegisterRtc_crFieldTsedgeMask  = 0x8
)

// GetTsedge Time-stamp event active edge TSE must be reset when TSEDGE is changed to avoid unwanted TSF setting.
func (r *registerRtc_crType) GetTsedge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldTsedgeMask) != 0
}

// SetTsedge Time-stamp event active edge TSE must be reset when TSEDGE is changed to avoid unwanted TSF setting.
func (r *registerRtc_crType) SetTsedge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldTsedgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldTsedgeMask)
	}
}

const (
	RegisterRtc_crFieldRefckonShift = 4
	RegisterRtc_crFieldRefckonMask  = 0x10
)

// GetRefckon RTC_REFIN reference clock detection enable (50 or 60Hz) Note: PREDIV_S must be 0x00FF.
func (r *registerRtc_crType) GetRefckon() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldRefckonMask) != 0
}

// SetRefckon RTC_REFIN reference clock detection enable (50 or 60Hz) Note: PREDIV_S must be 0x00FF.
func (r *registerRtc_crType) SetRefckon(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldRefckonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldRefckonMask)
	}
}

const (
	RegisterRtc_crFieldBypshadShift = 5
	RegisterRtc_crFieldBypshadMask  = 0x20
)

// GetBypshad Bypass the shadow registers Note: If the frequency of the APB clock is less than seven times the frequency of RTCCLK, BYPSHAD must be set to 1.
func (r *registerRtc_crType) GetBypshad() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldBypshadMask) != 0
}

// SetBypshad Bypass the shadow registers Note: If the frequency of the APB clock is less than seven times the frequency of RTCCLK, BYPSHAD must be set to 1.
func (r *registerRtc_crType) SetBypshad(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldBypshadMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldBypshadMask)
	}
}

const (
	RegisterRtc_crFieldFmtShift = 6
	RegisterRtc_crFieldFmtMask  = 0x40
)

// GetFmt Hour format
func (r *registerRtc_crType) GetFmt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldFmtMask) != 0
}

// SetFmt Hour format
func (r *registerRtc_crType) SetFmt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldFmtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldFmtMask)
	}
}

const (
	RegisterRtc_crFieldAlraeShift = 8
	RegisterRtc_crFieldAlraeMask  = 0x100
)

// GetAlrae Alarm A enable
func (r *registerRtc_crType) GetAlrae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldAlraeMask) != 0
}

// SetAlrae Alarm A enable
func (r *registerRtc_crType) SetAlrae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldAlraeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldAlraeMask)
	}
}

const (
	RegisterRtc_crFieldAlrbeShift = 9
	RegisterRtc_crFieldAlrbeMask  = 0x200
)

// GetAlrbe Alarm B enable
func (r *registerRtc_crType) GetAlrbe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldAlrbeMask) != 0
}

// SetAlrbe Alarm B enable
func (r *registerRtc_crType) SetAlrbe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldAlrbeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldAlrbeMask)
	}
}

const (
	RegisterRtc_crFieldWuteShift = 10
	RegisterRtc_crFieldWuteMask  = 0x400
)

// GetWute Wakeup timer enable
func (r *registerRtc_crType) GetWute() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldWuteMask) != 0
}

// SetWute Wakeup timer enable
func (r *registerRtc_crType) SetWute(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldWuteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldWuteMask)
	}
}

const (
	RegisterRtc_crFieldTseShift = 11
	RegisterRtc_crFieldTseMask  = 0x800
)

// GetTse timestamp enable
func (r *registerRtc_crType) GetTse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldTseMask) != 0
}

// SetTse timestamp enable
func (r *registerRtc_crType) SetTse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldTseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldTseMask)
	}
}

const (
	RegisterRtc_crFieldAlraieShift = 12
	RegisterRtc_crFieldAlraieMask  = 0x1000
)

// GetAlraie Alarm A interrupt enable
func (r *registerRtc_crType) GetAlraie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldAlraieMask) != 0
}

// SetAlraie Alarm A interrupt enable
func (r *registerRtc_crType) SetAlraie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldAlraieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldAlraieMask)
	}
}

const (
	RegisterRtc_crFieldAlrbieShift = 13
	RegisterRtc_crFieldAlrbieMask  = 0x2000
)

// GetAlrbie Alarm B interrupt enable
func (r *registerRtc_crType) GetAlrbie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldAlrbieMask) != 0
}

// SetAlrbie Alarm B interrupt enable
func (r *registerRtc_crType) SetAlrbie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldAlrbieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldAlrbieMask)
	}
}

const (
	RegisterRtc_crFieldWutieShift = 14
	RegisterRtc_crFieldWutieMask  = 0x4000
)

// GetWutie Wakeup timer interrupt enable
func (r *registerRtc_crType) GetWutie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldWutieMask) != 0
}

// SetWutie Wakeup timer interrupt enable
func (r *registerRtc_crType) SetWutie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldWutieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldWutieMask)
	}
}

const (
	RegisterRtc_crFieldTsieShift = 15
	RegisterRtc_crFieldTsieMask  = 0x8000
)

// GetTsie Time-stamp interrupt enable
func (r *registerRtc_crType) GetTsie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldTsieMask) != 0
}

// SetTsie Time-stamp interrupt enable
func (r *registerRtc_crType) SetTsie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldTsieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldTsieMask)
	}
}

const (
	RegisterRtc_crFieldAdd1hShift = 16
	RegisterRtc_crFieldAdd1hMask  = 0x10000
)

// SetAdd1h Add 1 hour (summer time change) When this bit is set outside initialization mode, 1 hour is added to the calendar time. This bit is always read as 0.
func (r *registerRtc_crType) SetAdd1h(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldAdd1hMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldAdd1hMask)
	}
}

const (
	RegisterRtc_crFieldSub1hShift = 17
	RegisterRtc_crFieldSub1hMask  = 0x20000
)

// SetSub1h Subtract 1 hour (winter time change) When this bit is set outside initialization mode, 1 hour is subtracted to the calendar time if the current hour is not 0. This bit is always read as 0. Setting this bit has no effect when current hour is 0.
func (r *registerRtc_crType) SetSub1h(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldSub1hMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldSub1hMask)
	}
}

const (
	RegisterRtc_crFieldBkpShift = 18
	RegisterRtc_crFieldBkpMask  = 0x40000
)

// GetBkp Backup This bit can be written by the user to memorize whether the daylight saving time change has been performed or not.
func (r *registerRtc_crType) GetBkp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldBkpMask) != 0
}

// SetBkp Backup This bit can be written by the user to memorize whether the daylight saving time change has been performed or not.
func (r *registerRtc_crType) SetBkp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldBkpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldBkpMask)
	}
}

const (
	RegisterRtc_crFieldCoselShift = 19
	RegisterRtc_crFieldCoselMask  = 0x80000
)

// GetCosel Calibration output selection When COE=1, this bit selects which signal is output on RTC_CALIB. These frequencies are valid for RTCCLK at 32.768 kHz and prescalers at their default values (PREDIV_A=127 and PREDIV_S=255). Refer to Section24.3.15: Calibration clock output
func (r *registerRtc_crType) GetCosel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldCoselMask) != 0
}

// SetCosel Calibration output selection When COE=1, this bit selects which signal is output on RTC_CALIB. These frequencies are valid for RTCCLK at 32.768 kHz and prescalers at their default values (PREDIV_A=127 and PREDIV_S=255). Refer to Section24.3.15: Calibration clock output
func (r *registerRtc_crType) SetCosel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldCoselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldCoselMask)
	}
}

const (
	RegisterRtc_crFieldPolShift = 20
	RegisterRtc_crFieldPolMask  = 0x100000
)

// GetPol Output polarity This bit is used to configure the polarity of RTC_ALARM output
func (r *registerRtc_crType) GetPol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldPolMask) != 0
}

// SetPol Output polarity This bit is used to configure the polarity of RTC_ALARM output
func (r *registerRtc_crType) SetPol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldPolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldPolMask)
	}
}

const (
	RegisterRtc_crFieldOselShift = 21
	RegisterRtc_crFieldOselMask  = 0x600000
)

// GetOsel Output selection These bits are used to select the flag to be routed to RTC_ALARM output
func (r *registerRtc_crType) GetOsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldOselMask) >> RegisterRtc_crFieldOselShift)
}

// SetOsel Output selection These bits are used to select the flag to be routed to RTC_ALARM output
func (r *registerRtc_crType) SetOsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldOselMask)|(uint32(value)<<RegisterRtc_crFieldOselShift))
}

const (
	RegisterRtc_crFieldCoeShift = 23
	RegisterRtc_crFieldCoeMask  = 0x800000
)

// GetCoe Calibration output enable This bit enables the RTC_CALIB output
func (r *registerRtc_crType) GetCoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldCoeMask) != 0
}

// SetCoe Calibration output enable This bit enables the RTC_CALIB output
func (r *registerRtc_crType) SetCoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldCoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldCoeMask)
	}
}

const (
	RegisterRtc_crFieldItseShift = 24
	RegisterRtc_crFieldItseMask  = 0x1000000
)

// GetItse timestamp on internal event enable
func (r *registerRtc_crType) GetItse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_crFieldItseMask) != 0
}

// SetItse timestamp on internal event enable
func (r *registerRtc_crType) SetItse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_crFieldItseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_crFieldItseMask)
	}
}

// registerRtc_isrType This register is write protected (except for RTC_ISR[13:8] bits). The write access procedure is described in RTC register write protection on page9.
type registerRtc_isrType uint32

const (
	RegisterRtc_isrFieldAlrawfShift = 0
	RegisterRtc_isrFieldAlrawfMask  = 0x1
)

// GetAlrawf Alarm A write flag This bit is set by hardware when Alarm A values can be changed, after the ALRAE bit has been set to 0 in RTC_CR. It is cleared by hardware in initialization mode.
func (r *registerRtc_isrType) GetAlrawf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldAlrawfMask) != 0
}

const (
	RegisterRtc_isrFieldAlrbwfShift = 1
	RegisterRtc_isrFieldAlrbwfMask  = 0x2
)

// GetAlrbwf Alarm B write flag This bit is set by hardware when Alarm B values can be changed, after the ALRBE bit has been set to 0 in RTC_CR. It is cleared by hardware in initialization mode.
func (r *registerRtc_isrType) GetAlrbwf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldAlrbwfMask) != 0
}

const (
	RegisterRtc_isrFieldWutwfShift = 2
	RegisterRtc_isrFieldWutwfMask  = 0x4
)

// GetWutwf Wakeup timer write flag This bit is set by hardware up to 2 RTCCLK cycles after the WUTE bit has been set to 0 in RTC_CR, and is cleared up to 2 RTCCLK cycles after the WUTE bit has been set to 1. The wakeup timer values can be changed when WUTE bit is cleared and WUTWF is set.
func (r *registerRtc_isrType) GetWutwf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldWutwfMask) != 0
}

const (
	RegisterRtc_isrFieldShpfShift = 3
	RegisterRtc_isrFieldShpfMask  = 0x8
)

// GetShpf Shift operation pending This flag is set by hardware as soon as a shift operation is initiated by a write to the RTC_SHIFTR register. It is cleared by hardware when the corresponding shift operation has been executed. Writing to the SHPF bit has no effect.
func (r *registerRtc_isrType) GetShpf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldShpfMask) != 0
}

const (
	RegisterRtc_isrFieldInitsShift = 4
	RegisterRtc_isrFieldInitsMask  = 0x10
)

// GetInits Initialization status flag This bit is set by hardware when the calendar year field is different from 0 (Backup domain reset state).
func (r *registerRtc_isrType) GetInits() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldInitsMask) != 0
}

const (
	RegisterRtc_isrFieldRsfShift = 5
	RegisterRtc_isrFieldRsfMask  = 0x20
)

// GetRsf Registers synchronization flag This bit is set by hardware each time the calendar registers are copied into the shadow registers (RTC_SSRx, RTC_TRx and RTC_DRx). This bit is cleared by hardware in initialization mode, while a shift operation is pending (SHPF=1), or when in bypass shadow register mode (BYPSHAD=1). This bit can also be cleared by software. It is cleared either by software or by hardware in initialization mode.
func (r *registerRtc_isrType) GetRsf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldRsfMask) != 0
}

// SetRsf Registers synchronization flag This bit is set by hardware each time the calendar registers are copied into the shadow registers (RTC_SSRx, RTC_TRx and RTC_DRx). This bit is cleared by hardware in initialization mode, while a shift operation is pending (SHPF=1), or when in bypass shadow register mode (BYPSHAD=1). This bit can also be cleared by software. It is cleared either by software or by hardware in initialization mode.
func (r *registerRtc_isrType) SetRsf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_isrFieldRsfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_isrFieldRsfMask)
	}
}

const (
	RegisterRtc_isrFieldInitfShift = 6
	RegisterRtc_isrFieldInitfMask  = 0x40
)

// GetInitf Initialization flag When this bit is set to 1, the RTC is in initialization state, and the time, date and prescaler registers can be updated.
func (r *registerRtc_isrType) GetInitf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldInitfMask) != 0
}

const (
	RegisterRtc_isrFieldInitShift = 7
	RegisterRtc_isrFieldInitMask  = 0x80
)

// GetInit Initialization mode
func (r *registerRtc_isrType) GetInit() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldInitMask) != 0
}

// SetInit Initialization mode
func (r *registerRtc_isrType) SetInit(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_isrFieldInitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_isrFieldInitMask)
	}
}

const (
	RegisterRtc_isrFieldAlrafShift = 8
	RegisterRtc_isrFieldAlrafMask  = 0x100
)

// GetAlraf Alarm A flag This flag is set by hardware when the time/date registers (RTC_TR and RTC_DR) match the Alarm A register (RTC_ALRMAR). This flag is cleared by software by writing 0.
func (r *registerRtc_isrType) GetAlraf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldAlrafMask) != 0
}

// SetAlraf Alarm A flag This flag is set by hardware when the time/date registers (RTC_TR and RTC_DR) match the Alarm A register (RTC_ALRMAR). This flag is cleared by software by writing 0.
func (r *registerRtc_isrType) SetAlraf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_isrFieldAlrafMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_isrFieldAlrafMask)
	}
}

const (
	RegisterRtc_isrFieldAlrbfShift = 9
	RegisterRtc_isrFieldAlrbfMask  = 0x200
)

// GetAlrbf Alarm B flag This flag is set by hardware when the time/date registers (RTC_TR and RTC_DR) match the Alarm B register (RTC_ALRMBR). This flag is cleared by software by writing 0.
func (r *registerRtc_isrType) GetAlrbf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldAlrbfMask) != 0
}

// SetAlrbf Alarm B flag This flag is set by hardware when the time/date registers (RTC_TR and RTC_DR) match the Alarm B register (RTC_ALRMBR). This flag is cleared by software by writing 0.
func (r *registerRtc_isrType) SetAlrbf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_isrFieldAlrbfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_isrFieldAlrbfMask)
	}
}

const (
	RegisterRtc_isrFieldWutfShift = 10
	RegisterRtc_isrFieldWutfMask  = 0x400
)

// GetWutf Wakeup timer flag This flag is set by hardware when the wakeup auto-reload counter reaches 0. This flag is cleared by software by writing 0. This flag must be cleared by software at least 1.5 RTCCLK periods before WUTF is set to 1 again.
func (r *registerRtc_isrType) GetWutf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldWutfMask) != 0
}

// SetWutf Wakeup timer flag This flag is set by hardware when the wakeup auto-reload counter reaches 0. This flag is cleared by software by writing 0. This flag must be cleared by software at least 1.5 RTCCLK periods before WUTF is set to 1 again.
func (r *registerRtc_isrType) SetWutf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_isrFieldWutfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_isrFieldWutfMask)
	}
}

const (
	RegisterRtc_isrFieldTsfShift = 11
	RegisterRtc_isrFieldTsfMask  = 0x800
)

// GetTsf Time-stamp flag This flag is set by hardware when a time-stamp event occurs. This flag is cleared by software by writing 0.
func (r *registerRtc_isrType) GetTsf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldTsfMask) != 0
}

// SetTsf Time-stamp flag This flag is set by hardware when a time-stamp event occurs. This flag is cleared by software by writing 0.
func (r *registerRtc_isrType) SetTsf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_isrFieldTsfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_isrFieldTsfMask)
	}
}

const (
	RegisterRtc_isrFieldTsovfShift = 12
	RegisterRtc_isrFieldTsovfMask  = 0x1000
)

// GetTsovf Time-stamp overflow flag This flag is set by hardware when a time-stamp event occurs while TSF is already set. This flag is cleared by software by writing 0. It is recommended to check and then clear TSOVF only after clearing the TSF bit. Otherwise, an overflow might not be noticed if a time-stamp event occurs immediately before the TSF bit is cleared.
func (r *registerRtc_isrType) GetTsovf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldTsovfMask) != 0
}

// SetTsovf Time-stamp overflow flag This flag is set by hardware when a time-stamp event occurs while TSF is already set. This flag is cleared by software by writing 0. It is recommended to check and then clear TSOVF only after clearing the TSF bit. Otherwise, an overflow might not be noticed if a time-stamp event occurs immediately before the TSF bit is cleared.
func (r *registerRtc_isrType) SetTsovf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_isrFieldTsovfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_isrFieldTsovfMask)
	}
}

const (
	RegisterRtc_isrFieldTamp1fShift = 13
	RegisterRtc_isrFieldTamp1fMask  = 0x2000
)

// GetTamp1f RTC_TAMP1 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP1 input. It is cleared by software writing 0
func (r *registerRtc_isrType) GetTamp1f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldTamp1fMask) != 0
}

// SetTamp1f RTC_TAMP1 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP1 input. It is cleared by software writing 0
func (r *registerRtc_isrType) SetTamp1f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_isrFieldTamp1fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_isrFieldTamp1fMask)
	}
}

const (
	RegisterRtc_isrFieldTamp2fShift = 14
	RegisterRtc_isrFieldTamp2fMask  = 0x4000
)

// GetTamp2f RTC_TAMP2 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP2 input. It is cleared by software writing 0
func (r *registerRtc_isrType) GetTamp2f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldTamp2fMask) != 0
}

// SetTamp2f RTC_TAMP2 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP2 input. It is cleared by software writing 0
func (r *registerRtc_isrType) SetTamp2f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_isrFieldTamp2fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_isrFieldTamp2fMask)
	}
}

const (
	RegisterRtc_isrFieldTamp3fShift = 15
	RegisterRtc_isrFieldTamp3fMask  = 0x8000
)

// GetTamp3f RTC_TAMP3 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP3 input. It is cleared by software writing 0
func (r *registerRtc_isrType) GetTamp3f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldTamp3fMask) != 0
}

// SetTamp3f RTC_TAMP3 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP3 input. It is cleared by software writing 0
func (r *registerRtc_isrType) SetTamp3f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_isrFieldTamp3fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_isrFieldTamp3fMask)
	}
}

const (
	RegisterRtc_isrFieldRecalpfShift = 16
	RegisterRtc_isrFieldRecalpfMask  = 0x10000
)

// GetRecalpf Recalibration pending Flag The RECALPF status flag is automatically set to 1 when software writes to the RTC_CALR register, indicating that the RTC_CALR register is blocked. When the new calibration settings are taken into account, this bit returns to 0. Refer to Re-calibration on-the-fly.
func (r *registerRtc_isrType) GetRecalpf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldRecalpfMask) != 0
}

const (
	RegisterRtc_isrFieldItsfShift = 17
	RegisterRtc_isrFieldItsfMask  = 0x20000
)

// GetItsf Internal tTime-stamp flag
func (r *registerRtc_isrType) GetItsf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_isrFieldItsfMask) != 0
}

// SetItsf Internal tTime-stamp flag
func (r *registerRtc_isrType) SetItsf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_isrFieldItsfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_isrFieldItsfMask)
	}
}

// registerRtc_prerType This register must be written in initialization mode only. The initialization must be performed in two separate write accesses. Refer to Calendar initialization and configuration on page9.This register is write protected. The write access procedure is described in RTC register write protection on page9.
type registerRtc_prerType uint32

const (
	RegisterRtc_prerFieldPrediv_sShift = 0
	RegisterRtc_prerFieldPrediv_sMask  = 0x7fff
)

// GetPrediv_s Synchronous prescaler factor This is the synchronous division factor: ck_spre frequency = ck_apre frequency/(PREDIV_S+1)
func (r *registerRtc_prerType) GetPrediv_s() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_prerFieldPrediv_sMask) >> RegisterRtc_prerFieldPrediv_sShift)
}

// SetPrediv_s Synchronous prescaler factor This is the synchronous division factor: ck_spre frequency = ck_apre frequency/(PREDIV_S+1)
func (r *registerRtc_prerType) SetPrediv_s(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_prerFieldPrediv_sMask)|(uint32(value)<<RegisterRtc_prerFieldPrediv_sShift))
}

const (
	RegisterRtc_prerFieldPrediv_aShift = 16
	RegisterRtc_prerFieldPrediv_aMask  = 0x7f0000
)

// GetPrediv_a Asynchronous prescaler factor This is the asynchronous division factor: ck_apre frequency = RTCCLK frequency/(PREDIV_A+1)
func (r *registerRtc_prerType) GetPrediv_a() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_prerFieldPrediv_aMask) >> RegisterRtc_prerFieldPrediv_aShift)
}

// SetPrediv_a Asynchronous prescaler factor This is the asynchronous division factor: ck_apre frequency = RTCCLK frequency/(PREDIV_A+1)
func (r *registerRtc_prerType) SetPrediv_a(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_prerFieldPrediv_aMask)|(uint32(value)<<RegisterRtc_prerFieldPrediv_aShift))
}

// registerRtc_wutrType This register can be written only when WUTWF is set to 1 in RTC_ISR.This register is write protected. The write access procedure is described in RTC register write protection on page9.
type registerRtc_wutrType uint32

const (
	RegisterRtc_wutrFieldWutShift = 0
	RegisterRtc_wutrFieldWutMask  = 0xffff
)

// GetWut Wakeup auto-reload value bits When the wakeup timer is enabled (WUTE set to 1), the WUTF flag is set every (WUT[15:0] + 1) ck_wut cycles. The ck_wut period is selected through WUCKSEL[2:0] bits of the RTC_CR register When WUCKSEL[2] = 1, the wakeup timer becomes 17-bits and WUCKSEL[1] effectively becomes WUT[16] the most-significant bit to be reloaded into the timer. The first assertion of WUTF occurs (WUT+1) ck_wut cycles after WUTE is set. Setting WUT[15:0] to 0x0000 with WUCKSEL[2:0] =011 (RTCCLK/2) is forbidden.
func (r *registerRtc_wutrType) GetWut() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_wutrFieldWutMask) >> RegisterRtc_wutrFieldWutShift)
}

// SetWut Wakeup auto-reload value bits When the wakeup timer is enabled (WUTE set to 1), the WUTF flag is set every (WUT[15:0] + 1) ck_wut cycles. The ck_wut period is selected through WUCKSEL[2:0] bits of the RTC_CR register When WUCKSEL[2] = 1, the wakeup timer becomes 17-bits and WUCKSEL[1] effectively becomes WUT[16] the most-significant bit to be reloaded into the timer. The first assertion of WUTF occurs (WUT+1) ck_wut cycles after WUTE is set. Setting WUT[15:0] to 0x0000 with WUCKSEL[2:0] =011 (RTCCLK/2) is forbidden.
func (r *registerRtc_wutrType) SetWut(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_wutrFieldWutMask)|(uint32(value)<<RegisterRtc_wutrFieldWutShift))
}

// registerRtc_alrmarType This register can be written only when ALRAWF is set to 1 in RTC_ISR, or in initialization mode.This register is write protected. The write access procedure is described in RTC register write protection on page9.
type registerRtc_alrmarType uint32

const (
	RegisterRtc_alrmarFieldSuShift = 0
	RegisterRtc_alrmarFieldSuMask  = 0xf
)

// GetSu Second units in BCD format.
func (r *registerRtc_alrmarType) GetSu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldSuMask) >> RegisterRtc_alrmarFieldSuShift)
}

// SetSu Second units in BCD format.
func (r *registerRtc_alrmarType) SetSu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldSuMask)|(uint32(value)<<RegisterRtc_alrmarFieldSuShift))
}

const (
	RegisterRtc_alrmarFieldStShift = 4
	RegisterRtc_alrmarFieldStMask  = 0x70
)

// GetSt Second tens in BCD format.
func (r *registerRtc_alrmarType) GetSt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldStMask) >> RegisterRtc_alrmarFieldStShift)
}

// SetSt Second tens in BCD format.
func (r *registerRtc_alrmarType) SetSt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldStMask)|(uint32(value)<<RegisterRtc_alrmarFieldStShift))
}

const (
	RegisterRtc_alrmarFieldMsk1Shift = 7
	RegisterRtc_alrmarFieldMsk1Mask  = 0x80
)

// GetMsk1 Alarm A seconds mask
func (r *registerRtc_alrmarType) GetMsk1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldMsk1Mask) != 0
}

// SetMsk1 Alarm A seconds mask
func (r *registerRtc_alrmarType) SetMsk1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_alrmarFieldMsk1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldMsk1Mask)
	}
}

const (
	RegisterRtc_alrmarFieldMnuShift = 8
	RegisterRtc_alrmarFieldMnuMask  = 0xf00
)

// GetMnu Minute units in BCD format.
func (r *registerRtc_alrmarType) GetMnu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldMnuMask) >> RegisterRtc_alrmarFieldMnuShift)
}

// SetMnu Minute units in BCD format.
func (r *registerRtc_alrmarType) SetMnu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldMnuMask)|(uint32(value)<<RegisterRtc_alrmarFieldMnuShift))
}

const (
	RegisterRtc_alrmarFieldMntShift = 12
	RegisterRtc_alrmarFieldMntMask  = 0x7000
)

// GetMnt Minute tens in BCD format.
func (r *registerRtc_alrmarType) GetMnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldMntMask) >> RegisterRtc_alrmarFieldMntShift)
}

// SetMnt Minute tens in BCD format.
func (r *registerRtc_alrmarType) SetMnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldMntMask)|(uint32(value)<<RegisterRtc_alrmarFieldMntShift))
}

const (
	RegisterRtc_alrmarFieldMsk2Shift = 15
	RegisterRtc_alrmarFieldMsk2Mask  = 0x8000
)

// GetMsk2 Alarm A minutes mask
func (r *registerRtc_alrmarType) GetMsk2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldMsk2Mask) != 0
}

// SetMsk2 Alarm A minutes mask
func (r *registerRtc_alrmarType) SetMsk2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_alrmarFieldMsk2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldMsk2Mask)
	}
}

const (
	RegisterRtc_alrmarFieldHuShift = 16
	RegisterRtc_alrmarFieldHuMask  = 0xf0000
)

// GetHu Hour units in BCD format.
func (r *registerRtc_alrmarType) GetHu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldHuMask) >> RegisterRtc_alrmarFieldHuShift)
}

// SetHu Hour units in BCD format.
func (r *registerRtc_alrmarType) SetHu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldHuMask)|(uint32(value)<<RegisterRtc_alrmarFieldHuShift))
}

const (
	RegisterRtc_alrmarFieldHtShift = 20
	RegisterRtc_alrmarFieldHtMask  = 0x300000
)

// GetHt Hour tens in BCD format.
func (r *registerRtc_alrmarType) GetHt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldHtMask) >> RegisterRtc_alrmarFieldHtShift)
}

// SetHt Hour tens in BCD format.
func (r *registerRtc_alrmarType) SetHt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldHtMask)|(uint32(value)<<RegisterRtc_alrmarFieldHtShift))
}

const (
	RegisterRtc_alrmarFieldPmShift = 22
	RegisterRtc_alrmarFieldPmMask  = 0x400000
)

// GetPm AM/PM notation
func (r *registerRtc_alrmarType) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldPmMask) != 0
}

// SetPm AM/PM notation
func (r *registerRtc_alrmarType) SetPm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_alrmarFieldPmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldPmMask)
	}
}

const (
	RegisterRtc_alrmarFieldMsk3Shift = 23
	RegisterRtc_alrmarFieldMsk3Mask  = 0x800000
)

// GetMsk3 Alarm A hours mask
func (r *registerRtc_alrmarType) GetMsk3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldMsk3Mask) != 0
}

// SetMsk3 Alarm A hours mask
func (r *registerRtc_alrmarType) SetMsk3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_alrmarFieldMsk3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldMsk3Mask)
	}
}

const (
	RegisterRtc_alrmarFieldDuShift = 24
	RegisterRtc_alrmarFieldDuMask  = 0xf000000
)

// GetDu Date units or day in BCD format.
func (r *registerRtc_alrmarType) GetDu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldDuMask) >> RegisterRtc_alrmarFieldDuShift)
}

// SetDu Date units or day in BCD format.
func (r *registerRtc_alrmarType) SetDu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldDuMask)|(uint32(value)<<RegisterRtc_alrmarFieldDuShift))
}

const (
	RegisterRtc_alrmarFieldDtShift = 28
	RegisterRtc_alrmarFieldDtMask  = 0x30000000
)

// GetDt Date tens in BCD format.
func (r *registerRtc_alrmarType) GetDt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldDtMask) >> RegisterRtc_alrmarFieldDtShift)
}

// SetDt Date tens in BCD format.
func (r *registerRtc_alrmarType) SetDt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldDtMask)|(uint32(value)<<RegisterRtc_alrmarFieldDtShift))
}

const (
	RegisterRtc_alrmarFieldWdselShift = 30
	RegisterRtc_alrmarFieldWdselMask  = 0x40000000
)

// GetWdsel Week day selection
func (r *registerRtc_alrmarType) GetWdsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldWdselMask) != 0
}

// SetWdsel Week day selection
func (r *registerRtc_alrmarType) SetWdsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_alrmarFieldWdselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldWdselMask)
	}
}

const (
	RegisterRtc_alrmarFieldMsk4Shift = 31
	RegisterRtc_alrmarFieldMsk4Mask  = 0x80000000
)

// GetMsk4 Alarm A date mask
func (r *registerRtc_alrmarType) GetMsk4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmarFieldMsk4Mask) != 0
}

// SetMsk4 Alarm A date mask
func (r *registerRtc_alrmarType) SetMsk4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_alrmarFieldMsk4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmarFieldMsk4Mask)
	}
}

// registerRtc_alrmbrType This register can be written only when ALRBWF is set to 1 in RTC_ISR, or in initialization mode.This register is write protected. The write access procedure is described in RTC register write protection on page9.
type registerRtc_alrmbrType uint32

const (
	RegisterRtc_alrmbrFieldSuShift = 0
	RegisterRtc_alrmbrFieldSuMask  = 0xf
)

// GetSu Second units in BCD format
func (r *registerRtc_alrmbrType) GetSu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldSuMask) >> RegisterRtc_alrmbrFieldSuShift)
}

// SetSu Second units in BCD format
func (r *registerRtc_alrmbrType) SetSu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldSuMask)|(uint32(value)<<RegisterRtc_alrmbrFieldSuShift))
}

const (
	RegisterRtc_alrmbrFieldStShift = 4
	RegisterRtc_alrmbrFieldStMask  = 0x70
)

// GetSt Second tens in BCD format
func (r *registerRtc_alrmbrType) GetSt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldStMask) >> RegisterRtc_alrmbrFieldStShift)
}

// SetSt Second tens in BCD format
func (r *registerRtc_alrmbrType) SetSt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldStMask)|(uint32(value)<<RegisterRtc_alrmbrFieldStShift))
}

const (
	RegisterRtc_alrmbrFieldMsk1Shift = 7
	RegisterRtc_alrmbrFieldMsk1Mask  = 0x80
)

// GetMsk1 Alarm B seconds mask
func (r *registerRtc_alrmbrType) GetMsk1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldMsk1Mask) != 0
}

// SetMsk1 Alarm B seconds mask
func (r *registerRtc_alrmbrType) SetMsk1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_alrmbrFieldMsk1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldMsk1Mask)
	}
}

const (
	RegisterRtc_alrmbrFieldMnuShift = 8
	RegisterRtc_alrmbrFieldMnuMask  = 0xf00
)

// GetMnu Minute units in BCD format
func (r *registerRtc_alrmbrType) GetMnu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldMnuMask) >> RegisterRtc_alrmbrFieldMnuShift)
}

// SetMnu Minute units in BCD format
func (r *registerRtc_alrmbrType) SetMnu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldMnuMask)|(uint32(value)<<RegisterRtc_alrmbrFieldMnuShift))
}

const (
	RegisterRtc_alrmbrFieldMntShift = 12
	RegisterRtc_alrmbrFieldMntMask  = 0x7000
)

// GetMnt Minute tens in BCD format
func (r *registerRtc_alrmbrType) GetMnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldMntMask) >> RegisterRtc_alrmbrFieldMntShift)
}

// SetMnt Minute tens in BCD format
func (r *registerRtc_alrmbrType) SetMnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldMntMask)|(uint32(value)<<RegisterRtc_alrmbrFieldMntShift))
}

const (
	RegisterRtc_alrmbrFieldMsk2Shift = 15
	RegisterRtc_alrmbrFieldMsk2Mask  = 0x8000
)

// GetMsk2 Alarm B minutes mask
func (r *registerRtc_alrmbrType) GetMsk2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldMsk2Mask) != 0
}

// SetMsk2 Alarm B minutes mask
func (r *registerRtc_alrmbrType) SetMsk2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_alrmbrFieldMsk2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldMsk2Mask)
	}
}

const (
	RegisterRtc_alrmbrFieldHuShift = 16
	RegisterRtc_alrmbrFieldHuMask  = 0xf0000
)

// GetHu Hour units in BCD format
func (r *registerRtc_alrmbrType) GetHu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldHuMask) >> RegisterRtc_alrmbrFieldHuShift)
}

// SetHu Hour units in BCD format
func (r *registerRtc_alrmbrType) SetHu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldHuMask)|(uint32(value)<<RegisterRtc_alrmbrFieldHuShift))
}

const (
	RegisterRtc_alrmbrFieldHtShift = 20
	RegisterRtc_alrmbrFieldHtMask  = 0x300000
)

// GetHt Hour tens in BCD format
func (r *registerRtc_alrmbrType) GetHt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldHtMask) >> RegisterRtc_alrmbrFieldHtShift)
}

// SetHt Hour tens in BCD format
func (r *registerRtc_alrmbrType) SetHt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldHtMask)|(uint32(value)<<RegisterRtc_alrmbrFieldHtShift))
}

const (
	RegisterRtc_alrmbrFieldPmShift = 22
	RegisterRtc_alrmbrFieldPmMask  = 0x400000
)

// GetPm AM/PM notation
func (r *registerRtc_alrmbrType) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldPmMask) != 0
}

// SetPm AM/PM notation
func (r *registerRtc_alrmbrType) SetPm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_alrmbrFieldPmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldPmMask)
	}
}

const (
	RegisterRtc_alrmbrFieldMsk3Shift = 23
	RegisterRtc_alrmbrFieldMsk3Mask  = 0x800000
)

// GetMsk3 Alarm B hours mask
func (r *registerRtc_alrmbrType) GetMsk3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldMsk3Mask) != 0
}

// SetMsk3 Alarm B hours mask
func (r *registerRtc_alrmbrType) SetMsk3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_alrmbrFieldMsk3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldMsk3Mask)
	}
}

const (
	RegisterRtc_alrmbrFieldDuShift = 24
	RegisterRtc_alrmbrFieldDuMask  = 0xf000000
)

// GetDu Date units or day in BCD format
func (r *registerRtc_alrmbrType) GetDu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldDuMask) >> RegisterRtc_alrmbrFieldDuShift)
}

// SetDu Date units or day in BCD format
func (r *registerRtc_alrmbrType) SetDu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldDuMask)|(uint32(value)<<RegisterRtc_alrmbrFieldDuShift))
}

const (
	RegisterRtc_alrmbrFieldDtShift = 28
	RegisterRtc_alrmbrFieldDtMask  = 0x30000000
)

// GetDt Date tens in BCD format
func (r *registerRtc_alrmbrType) GetDt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldDtMask) >> RegisterRtc_alrmbrFieldDtShift)
}

// SetDt Date tens in BCD format
func (r *registerRtc_alrmbrType) SetDt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldDtMask)|(uint32(value)<<RegisterRtc_alrmbrFieldDtShift))
}

const (
	RegisterRtc_alrmbrFieldWdselShift = 30
	RegisterRtc_alrmbrFieldWdselMask  = 0x40000000
)

// GetWdsel Week day selection
func (r *registerRtc_alrmbrType) GetWdsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldWdselMask) != 0
}

// SetWdsel Week day selection
func (r *registerRtc_alrmbrType) SetWdsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_alrmbrFieldWdselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldWdselMask)
	}
}

const (
	RegisterRtc_alrmbrFieldMsk4Shift = 31
	RegisterRtc_alrmbrFieldMsk4Mask  = 0x80000000
)

// GetMsk4 Alarm B date mask
func (r *registerRtc_alrmbrType) GetMsk4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbrFieldMsk4Mask) != 0
}

// SetMsk4 Alarm B date mask
func (r *registerRtc_alrmbrType) SetMsk4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_alrmbrFieldMsk4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbrFieldMsk4Mask)
	}
}

// registerRtc_wprType RTC write protection register
type registerRtc_wprType uint32

const (
	RegisterRtc_wprFieldKeyShift = 0
	RegisterRtc_wprFieldKeyMask  = 0xff
)

// GetKey Write protection key This byte is written by software. Reading this byte always returns 0x00. Refer to RTC register write protection for a description of how to unlock RTC register write protection.
func (r *registerRtc_wprType) GetKey() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_wprFieldKeyMask) >> RegisterRtc_wprFieldKeyShift)
}

// SetKey Write protection key This byte is written by software. Reading this byte always returns 0x00. Refer to RTC register write protection for a description of how to unlock RTC register write protection.
func (r *registerRtc_wprType) SetKey(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_wprFieldKeyMask)|(uint32(value)<<RegisterRtc_wprFieldKeyShift))
}

// registerRtc_ssrType RTC sub second register
type registerRtc_ssrType uint32

const (
	RegisterRtc_ssrFieldSsShift = 0
	RegisterRtc_ssrFieldSsMask  = 0xffff
)

// GetSs Sub second value SS[15:0] is the value in the synchronous prescaler counter. The fraction of a second is given by the formula below: Second fraction = (PREDIV_S - SS) / (PREDIV_S + 1) Note: SS can be larger than PREDIV_S only after a shift operation. In that case, the correct time/date is one second less than as indicated by RTC_TR/RTC_DR.
func (r *registerRtc_ssrType) GetSs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_ssrFieldSsMask) >> RegisterRtc_ssrFieldSsShift)
}

// SetSs Sub second value SS[15:0] is the value in the synchronous prescaler counter. The fraction of a second is given by the formula below: Second fraction = (PREDIV_S - SS) / (PREDIV_S + 1) Note: SS can be larger than PREDIV_S only after a shift operation. In that case, the correct time/date is one second less than as indicated by RTC_TR/RTC_DR.
func (r *registerRtc_ssrType) SetSs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_ssrFieldSsMask)|(uint32(value)<<RegisterRtc_ssrFieldSsShift))
}

// registerRtc_shiftrType This register is write protected. The write access procedure is described in RTC register write protection on page9.
type registerRtc_shiftrType uint32

const (
	RegisterRtc_shiftrFieldSubfsShift = 0
	RegisterRtc_shiftrFieldSubfsMask  = 0x7fff
)

// GetSubfs Subtract a fraction of a second These bits are write only and is always read as zero. Writing to this bit has no effect when a shift operation is pending (when SHPF=1, in RTC_ISR). The value which is written to SUBFS is added to the synchronous prescaler counter. Since this counter counts down, this operation effectively subtracts from (delays) the clock by: Delay (seconds) = SUBFS / (PREDIV_S + 1) A fraction of a second can effectively be added to the clock (advancing the clock) when the ADD1S function is used in conjunction with SUBFS, effectively advancing the clock by: Advance (seconds) = (1 - (SUBFS / (PREDIV_S + 1))). Note: Writing to SUBFS causes RSF to be cleared. Software can then wait until RSF=1 to be sure that the shadow registers have been updated with the shifted time.
func (r *registerRtc_shiftrType) GetSubfs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_shiftrFieldSubfsMask) >> RegisterRtc_shiftrFieldSubfsShift)
}

// SetSubfs Subtract a fraction of a second These bits are write only and is always read as zero. Writing to this bit has no effect when a shift operation is pending (when SHPF=1, in RTC_ISR). The value which is written to SUBFS is added to the synchronous prescaler counter. Since this counter counts down, this operation effectively subtracts from (delays) the clock by: Delay (seconds) = SUBFS / (PREDIV_S + 1) A fraction of a second can effectively be added to the clock (advancing the clock) when the ADD1S function is used in conjunction with SUBFS, effectively advancing the clock by: Advance (seconds) = (1 - (SUBFS / (PREDIV_S + 1))). Note: Writing to SUBFS causes RSF to be cleared. Software can then wait until RSF=1 to be sure that the shadow registers have been updated with the shifted time.
func (r *registerRtc_shiftrType) SetSubfs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_shiftrFieldSubfsMask)|(uint32(value)<<RegisterRtc_shiftrFieldSubfsShift))
}

const (
	RegisterRtc_shiftrFieldAdd1sShift = 31
	RegisterRtc_shiftrFieldAdd1sMask  = 0x80000000
)

// GetAdd1s Add one second This bit is write only and is always read as zero. Writing to this bit has no effect when a shift operation is pending (when SHPF=1, in RTC_ISR). This function is intended to be used with SUBFS (see description below) in order to effectively add a fraction of a second to the clock in an atomic operation.
func (r *registerRtc_shiftrType) GetAdd1s() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_shiftrFieldAdd1sMask) != 0
}

// SetAdd1s Add one second This bit is write only and is always read as zero. Writing to this bit has no effect when a shift operation is pending (when SHPF=1, in RTC_ISR). This function is intended to be used with SUBFS (see description below) in order to effectively add a fraction of a second to the clock in an atomic operation.
func (r *registerRtc_shiftrType) SetAdd1s(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_shiftrFieldAdd1sMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_shiftrFieldAdd1sMask)
	}
}

// registerRtc_tstrType The content of this register is valid only when TSF is set to 1 in RTC_ISR. It is cleared when TSF bit is reset.
type registerRtc_tstrType uint32

const (
	RegisterRtc_tstrFieldSuShift = 0
	RegisterRtc_tstrFieldSuMask  = 0xf
)

// GetSu Second units in BCD format.
func (r *registerRtc_tstrType) GetSu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tstrFieldSuMask) >> RegisterRtc_tstrFieldSuShift)
}

// SetSu Second units in BCD format.
func (r *registerRtc_tstrType) SetSu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tstrFieldSuMask)|(uint32(value)<<RegisterRtc_tstrFieldSuShift))
}

const (
	RegisterRtc_tstrFieldStShift = 4
	RegisterRtc_tstrFieldStMask  = 0x70
)

// GetSt Second tens in BCD format.
func (r *registerRtc_tstrType) GetSt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tstrFieldStMask) >> RegisterRtc_tstrFieldStShift)
}

// SetSt Second tens in BCD format.
func (r *registerRtc_tstrType) SetSt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tstrFieldStMask)|(uint32(value)<<RegisterRtc_tstrFieldStShift))
}

const (
	RegisterRtc_tstrFieldMnuShift = 8
	RegisterRtc_tstrFieldMnuMask  = 0xf00
)

// GetMnu Minute units in BCD format.
func (r *registerRtc_tstrType) GetMnu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tstrFieldMnuMask) >> RegisterRtc_tstrFieldMnuShift)
}

// SetMnu Minute units in BCD format.
func (r *registerRtc_tstrType) SetMnu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tstrFieldMnuMask)|(uint32(value)<<RegisterRtc_tstrFieldMnuShift))
}

const (
	RegisterRtc_tstrFieldMntShift = 12
	RegisterRtc_tstrFieldMntMask  = 0x7000
)

// GetMnt Minute tens in BCD format.
func (r *registerRtc_tstrType) GetMnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tstrFieldMntMask) >> RegisterRtc_tstrFieldMntShift)
}

// SetMnt Minute tens in BCD format.
func (r *registerRtc_tstrType) SetMnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tstrFieldMntMask)|(uint32(value)<<RegisterRtc_tstrFieldMntShift))
}

const (
	RegisterRtc_tstrFieldHuShift = 16
	RegisterRtc_tstrFieldHuMask  = 0xf0000
)

// GetHu Hour units in BCD format.
func (r *registerRtc_tstrType) GetHu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tstrFieldHuMask) >> RegisterRtc_tstrFieldHuShift)
}

// SetHu Hour units in BCD format.
func (r *registerRtc_tstrType) SetHu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tstrFieldHuMask)|(uint32(value)<<RegisterRtc_tstrFieldHuShift))
}

const (
	RegisterRtc_tstrFieldHtShift = 20
	RegisterRtc_tstrFieldHtMask  = 0x300000
)

// GetHt Hour tens in BCD format.
func (r *registerRtc_tstrType) GetHt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tstrFieldHtMask) >> RegisterRtc_tstrFieldHtShift)
}

// SetHt Hour tens in BCD format.
func (r *registerRtc_tstrType) SetHt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tstrFieldHtMask)|(uint32(value)<<RegisterRtc_tstrFieldHtShift))
}

const (
	RegisterRtc_tstrFieldPmShift = 22
	RegisterRtc_tstrFieldPmMask  = 0x400000
)

// GetPm AM/PM notation
func (r *registerRtc_tstrType) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tstrFieldPmMask) != 0
}

// SetPm AM/PM notation
func (r *registerRtc_tstrType) SetPm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tstrFieldPmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tstrFieldPmMask)
	}
}

// registerRtc_tsdrType The content of this register is valid only when TSF is set to 1 in RTC_ISR. It is cleared when TSF bit is reset.
type registerRtc_tsdrType uint32

const (
	RegisterRtc_tsdrFieldDuShift = 0
	RegisterRtc_tsdrFieldDuMask  = 0xf
)

// GetDu Date units in BCD format
func (r *registerRtc_tsdrType) GetDu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tsdrFieldDuMask) >> RegisterRtc_tsdrFieldDuShift)
}

// SetDu Date units in BCD format
func (r *registerRtc_tsdrType) SetDu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tsdrFieldDuMask)|(uint32(value)<<RegisterRtc_tsdrFieldDuShift))
}

const (
	RegisterRtc_tsdrFieldDtShift = 4
	RegisterRtc_tsdrFieldDtMask  = 0x30
)

// GetDt Date tens in BCD format
func (r *registerRtc_tsdrType) GetDt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tsdrFieldDtMask) >> RegisterRtc_tsdrFieldDtShift)
}

// SetDt Date tens in BCD format
func (r *registerRtc_tsdrType) SetDt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tsdrFieldDtMask)|(uint32(value)<<RegisterRtc_tsdrFieldDtShift))
}

const (
	RegisterRtc_tsdrFieldMuShift = 8
	RegisterRtc_tsdrFieldMuMask  = 0xf00
)

// GetMu Month units in BCD format
func (r *registerRtc_tsdrType) GetMu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tsdrFieldMuMask) >> RegisterRtc_tsdrFieldMuShift)
}

// SetMu Month units in BCD format
func (r *registerRtc_tsdrType) SetMu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tsdrFieldMuMask)|(uint32(value)<<RegisterRtc_tsdrFieldMuShift))
}

const (
	RegisterRtc_tsdrFieldMtShift = 12
	RegisterRtc_tsdrFieldMtMask  = 0x1000
)

// GetMt Month tens in BCD format
func (r *registerRtc_tsdrType) GetMt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tsdrFieldMtMask) != 0
}

// SetMt Month tens in BCD format
func (r *registerRtc_tsdrType) SetMt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tsdrFieldMtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tsdrFieldMtMask)
	}
}

const (
	RegisterRtc_tsdrFieldWduShift = 13
	RegisterRtc_tsdrFieldWduMask  = 0xe000
)

// GetWdu Week day units
func (r *registerRtc_tsdrType) GetWdu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tsdrFieldWduMask) >> RegisterRtc_tsdrFieldWduShift)
}

// SetWdu Week day units
func (r *registerRtc_tsdrType) SetWdu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tsdrFieldWduMask)|(uint32(value)<<RegisterRtc_tsdrFieldWduShift))
}

// registerRtc_tsssrType The content of this register is valid only when RTC_ISR/TSF is set. It is cleared when the RTC_ISR/TSF bit is reset.
type registerRtc_tsssrType uint32

const (
	RegisterRtc_tsssrFieldSsShift = 0
	RegisterRtc_tsssrFieldSsMask  = 0xffff
)

// GetSs Sub second value SS[15:0] is the value of the synchronous prescaler counter when the timestamp event occurred.
func (r *registerRtc_tsssrType) GetSs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tsssrFieldSsMask) >> RegisterRtc_tsssrFieldSsShift)
}

// SetSs Sub second value SS[15:0] is the value of the synchronous prescaler counter when the timestamp event occurred.
func (r *registerRtc_tsssrType) SetSs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tsssrFieldSsMask)|(uint32(value)<<RegisterRtc_tsssrFieldSsShift))
}

// registerRtc_calrType This register is write protected. The write access procedure is described in RTC register write protection on page9.
type registerRtc_calrType uint32

const (
	RegisterRtc_calrFieldCalmShift = 0
	RegisterRtc_calrFieldCalmMask  = 0x1ff
)

// GetCalm Calibration minus The frequency of the calendar is reduced by masking CALM out of 220 RTCCLK pulses (32 seconds if the input frequency is 32768 Hz). This decreases the frequency of the calendar with a resolution of 0.9537 ppm. To increase the frequency of the calendar, this feature should be used in conjunction with CALP. See Section24.3.12: RTC smooth digital calibration on page13.
func (r *registerRtc_calrType) GetCalm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_calrFieldCalmMask) >> RegisterRtc_calrFieldCalmShift)
}

// SetCalm Calibration minus The frequency of the calendar is reduced by masking CALM out of 220 RTCCLK pulses (32 seconds if the input frequency is 32768 Hz). This decreases the frequency of the calendar with a resolution of 0.9537 ppm. To increase the frequency of the calendar, this feature should be used in conjunction with CALP. See Section24.3.12: RTC smooth digital calibration on page13.
func (r *registerRtc_calrType) SetCalm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_calrFieldCalmMask)|(uint32(value)<<RegisterRtc_calrFieldCalmShift))
}

const (
	RegisterRtc_calrFieldCalw16Shift = 13
	RegisterRtc_calrFieldCalw16Mask  = 0x2000
)

// GetCalw16 Use a 16-second calibration cycle period When CALW16 is set to 1, the 16-second calibration cycle period is selected.This bit must not be set to 1 if CALW8=1. Note: CALM[0] is stuck at 0 when CALW16= 1. Refer to Section24.3.12: RTC smooth digital calibration.
func (r *registerRtc_calrType) GetCalw16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_calrFieldCalw16Mask) != 0
}

// SetCalw16 Use a 16-second calibration cycle period When CALW16 is set to 1, the 16-second calibration cycle period is selected.This bit must not be set to 1 if CALW8=1. Note: CALM[0] is stuck at 0 when CALW16= 1. Refer to Section24.3.12: RTC smooth digital calibration.
func (r *registerRtc_calrType) SetCalw16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_calrFieldCalw16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_calrFieldCalw16Mask)
	}
}

const (
	RegisterRtc_calrFieldCalw8Shift = 14
	RegisterRtc_calrFieldCalw8Mask  = 0x4000
)

// GetCalw8 Use an 8-second calibration cycle period When CALW8 is set to 1, the 8-second calibration cycle period is selected. Note: CALM[1:0] are stuck at 00; when CALW8= 1. Refer to Section24.3.12: RTC smooth digital calibration.
func (r *registerRtc_calrType) GetCalw8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_calrFieldCalw8Mask) != 0
}

// SetCalw8 Use an 8-second calibration cycle period When CALW8 is set to 1, the 8-second calibration cycle period is selected. Note: CALM[1:0] are stuck at 00; when CALW8= 1. Refer to Section24.3.12: RTC smooth digital calibration.
func (r *registerRtc_calrType) SetCalw8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_calrFieldCalw8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_calrFieldCalw8Mask)
	}
}

const (
	RegisterRtc_calrFieldCalpShift = 15
	RegisterRtc_calrFieldCalpMask  = 0x8000
)

// GetCalp Increase frequency of RTC by 488.5 ppm This feature is intended to be used in conjunction with CALM, which lowers the frequency of the calendar with a fine resolution. if the input frequency is 32768 Hz, the number of RTCCLK pulses added during a 32-second window is calculated as follows: (512 * CALP) - CALM. Refer to Section24.3.12: RTC smooth digital calibration.
func (r *registerRtc_calrType) GetCalp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_calrFieldCalpMask) != 0
}

// SetCalp Increase frequency of RTC by 488.5 ppm This feature is intended to be used in conjunction with CALM, which lowers the frequency of the calendar with a fine resolution. if the input frequency is 32768 Hz, the number of RTCCLK pulses added during a 32-second window is calculated as follows: (512 * CALP) - CALM. Refer to Section24.3.12: RTC smooth digital calibration.
func (r *registerRtc_calrType) SetCalp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_calrFieldCalpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_calrFieldCalpMask)
	}
}

// registerRtc_tampcrType RTC tamper and alternate function configuration register
type registerRtc_tampcrType uint32

const (
	RegisterRtc_tampcrFieldTamp1eShift = 0
	RegisterRtc_tampcrFieldTamp1eMask  = 0x1
)

// GetTamp1e RTC_TAMP1 input detection enable
func (r *registerRtc_tampcrType) GetTamp1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp1eMask) != 0
}

// SetTamp1e RTC_TAMP1 input detection enable
func (r *registerRtc_tampcrType) SetTamp1e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp1eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp1eMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp1trgShift = 1
	RegisterRtc_tampcrFieldTamp1trgMask  = 0x2
)

// GetTamp1trg Active level for RTC_TAMP1 input If TAMPFLT != 00 if TAMPFLT = 00:
func (r *registerRtc_tampcrType) GetTamp1trg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp1trgMask) != 0
}

// SetTamp1trg Active level for RTC_TAMP1 input If TAMPFLT != 00 if TAMPFLT = 00:
func (r *registerRtc_tampcrType) SetTamp1trg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp1trgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp1trgMask)
	}
}

const (
	RegisterRtc_tampcrFieldTampieShift = 2
	RegisterRtc_tampcrFieldTampieMask  = 0x4
)

// GetTampie Tamper interrupt enable
func (r *registerRtc_tampcrType) GetTampie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTampieMask) != 0
}

// SetTampie Tamper interrupt enable
func (r *registerRtc_tampcrType) SetTampie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTampieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTampieMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp2eShift = 3
	RegisterRtc_tampcrFieldTamp2eMask  = 0x8
)

// GetTamp2e RTC_TAMP2 input detection enable
func (r *registerRtc_tampcrType) GetTamp2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp2eMask) != 0
}

// SetTamp2e RTC_TAMP2 input detection enable
func (r *registerRtc_tampcrType) SetTamp2e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp2eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp2eMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp2trgShift = 4
	RegisterRtc_tampcrFieldTamp2trgMask  = 0x10
)

// GetTamp2trg Active level for RTC_TAMP2 input if TAMPFLT != 00: if TAMPFLT = 00:
func (r *registerRtc_tampcrType) GetTamp2trg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp2trgMask) != 0
}

// SetTamp2trg Active level for RTC_TAMP2 input if TAMPFLT != 00: if TAMPFLT = 00:
func (r *registerRtc_tampcrType) SetTamp2trg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp2trgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp2trgMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp3eShift = 5
	RegisterRtc_tampcrFieldTamp3eMask  = 0x20
)

// GetTamp3e RTC_TAMP3 detection enable
func (r *registerRtc_tampcrType) GetTamp3e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp3eMask) != 0
}

// SetTamp3e RTC_TAMP3 detection enable
func (r *registerRtc_tampcrType) SetTamp3e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp3eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp3eMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp3trgShift = 6
	RegisterRtc_tampcrFieldTamp3trgMask  = 0x40
)

// GetTamp3trg Active level for RTC_TAMP3 input if TAMPFLT != 00: if TAMPFLT = 00:
func (r *registerRtc_tampcrType) GetTamp3trg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp3trgMask) != 0
}

// SetTamp3trg Active level for RTC_TAMP3 input if TAMPFLT != 00: if TAMPFLT = 00:
func (r *registerRtc_tampcrType) SetTamp3trg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp3trgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp3trgMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamptsShift = 7
	RegisterRtc_tampcrFieldTamptsMask  = 0x80
)

// GetTampts Activate timestamp on tamper detection event TAMPTS is valid even if TSE=0 in the RTC_CR register.
func (r *registerRtc_tampcrType) GetTampts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamptsMask) != 0
}

// SetTampts Activate timestamp on tamper detection event TAMPTS is valid even if TSE=0 in the RTC_CR register.
func (r *registerRtc_tampcrType) SetTampts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamptsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamptsMask)
	}
}

const (
	RegisterRtc_tampcrFieldTampfreqShift = 8
	RegisterRtc_tampcrFieldTampfreqMask  = 0x700
)

// GetTampfreq Tamper sampling frequency Determines the frequency at which each of the RTC_TAMPx inputs are sampled.
func (r *registerRtc_tampcrType) GetTampfreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTampfreqMask) >> RegisterRtc_tampcrFieldTampfreqShift)
}

// SetTampfreq Tamper sampling frequency Determines the frequency at which each of the RTC_TAMPx inputs are sampled.
func (r *registerRtc_tampcrType) SetTampfreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTampfreqMask)|(uint32(value)<<RegisterRtc_tampcrFieldTampfreqShift))
}

const (
	RegisterRtc_tampcrFieldTampfltShift = 11
	RegisterRtc_tampcrFieldTampfltMask  = 0x1800
)

// GetTampflt RTC_TAMPx filter count These bits determines the number of consecutive samples at the specified level (TAMP*TRG) needed to activate a Tamper event. TAMPFLT is valid for each of the RTC_TAMPx inputs.
func (r *registerRtc_tampcrType) GetTampflt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTampfltMask) >> RegisterRtc_tampcrFieldTampfltShift)
}

// SetTampflt RTC_TAMPx filter count These bits determines the number of consecutive samples at the specified level (TAMP*TRG) needed to activate a Tamper event. TAMPFLT is valid for each of the RTC_TAMPx inputs.
func (r *registerRtc_tampcrType) SetTampflt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTampfltMask)|(uint32(value)<<RegisterRtc_tampcrFieldTampfltShift))
}

const (
	RegisterRtc_tampcrFieldTampprchShift = 13
	RegisterRtc_tampcrFieldTampprchMask  = 0x6000
)

// GetTampprch RTC_TAMPx precharge duration These bit determines the duration of time during which the pull-up/is activated before each sample. TAMPPRCH is valid for each of the RTC_TAMPx inputs.
func (r *registerRtc_tampcrType) GetTampprch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTampprchMask) >> RegisterRtc_tampcrFieldTampprchShift)
}

// SetTampprch RTC_TAMPx precharge duration These bit determines the duration of time during which the pull-up/is activated before each sample. TAMPPRCH is valid for each of the RTC_TAMPx inputs.
func (r *registerRtc_tampcrType) SetTampprch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTampprchMask)|(uint32(value)<<RegisterRtc_tampcrFieldTampprchShift))
}

const (
	RegisterRtc_tampcrFieldTamppudisShift = 15
	RegisterRtc_tampcrFieldTamppudisMask  = 0x8000
)

// GetTamppudis RTC_TAMPx pull-up disable This bit determines if each of the RTC_TAMPx pins are pre-charged before each sample.
func (r *registerRtc_tampcrType) GetTamppudis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamppudisMask) != 0
}

// SetTamppudis RTC_TAMPx pull-up disable This bit determines if each of the RTC_TAMPx pins are pre-charged before each sample.
func (r *registerRtc_tampcrType) SetTamppudis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamppudisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamppudisMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp1ieShift = 16
	RegisterRtc_tampcrFieldTamp1ieMask  = 0x10000
)

// GetTamp1ie Tamper 1 interrupt enable
func (r *registerRtc_tampcrType) GetTamp1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp1ieMask) != 0
}

// SetTamp1ie Tamper 1 interrupt enable
func (r *registerRtc_tampcrType) SetTamp1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp1ieMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp1noeraseShift = 17
	RegisterRtc_tampcrFieldTamp1noeraseMask  = 0x20000
)

// GetTamp1noerase Tamper 1 no erase
func (r *registerRtc_tampcrType) GetTamp1noerase() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp1noeraseMask) != 0
}

// SetTamp1noerase Tamper 1 no erase
func (r *registerRtc_tampcrType) SetTamp1noerase(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp1noeraseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp1noeraseMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp1mfShift = 18
	RegisterRtc_tampcrFieldTamp1mfMask  = 0x40000
)

// GetTamp1mf Tamper 1 mask flag
func (r *registerRtc_tampcrType) GetTamp1mf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp1mfMask) != 0
}

// SetTamp1mf Tamper 1 mask flag
func (r *registerRtc_tampcrType) SetTamp1mf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp1mfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp1mfMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp2ieShift = 19
	RegisterRtc_tampcrFieldTamp2ieMask  = 0x80000
)

// GetTamp2ie Tamper 2 interrupt enable
func (r *registerRtc_tampcrType) GetTamp2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp2ieMask) != 0
}

// SetTamp2ie Tamper 2 interrupt enable
func (r *registerRtc_tampcrType) SetTamp2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp2ieMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp2noeraseShift = 20
	RegisterRtc_tampcrFieldTamp2noeraseMask  = 0x100000
)

// GetTamp2noerase Tamper 2 no erase
func (r *registerRtc_tampcrType) GetTamp2noerase() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp2noeraseMask) != 0
}

// SetTamp2noerase Tamper 2 no erase
func (r *registerRtc_tampcrType) SetTamp2noerase(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp2noeraseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp2noeraseMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp2mfShift = 21
	RegisterRtc_tampcrFieldTamp2mfMask  = 0x200000
)

// GetTamp2mf Tamper 2 mask flag
func (r *registerRtc_tampcrType) GetTamp2mf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp2mfMask) != 0
}

// SetTamp2mf Tamper 2 mask flag
func (r *registerRtc_tampcrType) SetTamp2mf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp2mfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp2mfMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp3ieShift = 22
	RegisterRtc_tampcrFieldTamp3ieMask  = 0x400000
)

// GetTamp3ie Tamper 3 interrupt enable
func (r *registerRtc_tampcrType) GetTamp3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp3ieMask) != 0
}

// SetTamp3ie Tamper 3 interrupt enable
func (r *registerRtc_tampcrType) SetTamp3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp3ieMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp3noeraseShift = 23
	RegisterRtc_tampcrFieldTamp3noeraseMask  = 0x800000
)

// GetTamp3noerase Tamper 3 no erase
func (r *registerRtc_tampcrType) GetTamp3noerase() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp3noeraseMask) != 0
}

// SetTamp3noerase Tamper 3 no erase
func (r *registerRtc_tampcrType) SetTamp3noerase(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp3noeraseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp3noeraseMask)
	}
}

const (
	RegisterRtc_tampcrFieldTamp3mfShift = 24
	RegisterRtc_tampcrFieldTamp3mfMask  = 0x1000000
)

// GetTamp3mf Tamper 3 mask flag
func (r *registerRtc_tampcrType) GetTamp3mf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_tampcrFieldTamp3mfMask) != 0
}

// SetTamp3mf Tamper 3 mask flag
func (r *registerRtc_tampcrType) SetTamp3mf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_tampcrFieldTamp3mfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_tampcrFieldTamp3mfMask)
	}
}

// registerRtc_alrmassrType This register can be written only when ALRAE is reset in RTC_CR register, or in initialization mode.This register is write protected. The write access procedure is described in RTC register write protection on page9
type registerRtc_alrmassrType uint32

const (
	RegisterRtc_alrmassrFieldSsShift = 0
	RegisterRtc_alrmassrFieldSsMask  = 0x7fff
)

// GetSs Sub seconds value This value is compared with the contents of the synchronous prescaler counter to determine if Alarm A is to be activated. Only bits 0 up MASKSS-1 are compared.
func (r *registerRtc_alrmassrType) GetSs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmassrFieldSsMask) >> RegisterRtc_alrmassrFieldSsShift)
}

// SetSs Sub seconds value This value is compared with the contents of the synchronous prescaler counter to determine if Alarm A is to be activated. Only bits 0 up MASKSS-1 are compared.
func (r *registerRtc_alrmassrType) SetSs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmassrFieldSsMask)|(uint32(value)<<RegisterRtc_alrmassrFieldSsShift))
}

const (
	RegisterRtc_alrmassrFieldMaskssShift = 24
	RegisterRtc_alrmassrFieldMaskssMask  = 0xf000000
)

// GetMaskss Mask the most-significant bits starting at this bit ... The overflow bits of the synchronous counter (bits 15) is never compared. This bit can be different from 0 only after a shift operation.
func (r *registerRtc_alrmassrType) GetMaskss() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmassrFieldMaskssMask) >> RegisterRtc_alrmassrFieldMaskssShift)
}

// SetMaskss Mask the most-significant bits starting at this bit ... The overflow bits of the synchronous counter (bits 15) is never compared. This bit can be different from 0 only after a shift operation.
func (r *registerRtc_alrmassrType) SetMaskss(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmassrFieldMaskssMask)|(uint32(value)<<RegisterRtc_alrmassrFieldMaskssShift))
}

// registerRtc_alrmbssrType This register can be written only when ALRBE is reset in RTC_CR register, or in initialization mode.This register is write protected.The write access procedure is described in Section: RTC register write protection.
type registerRtc_alrmbssrType uint32

const (
	RegisterRtc_alrmbssrFieldSsShift = 0
	RegisterRtc_alrmbssrFieldSsMask  = 0x7fff
)

// GetSs Sub seconds value This value is compared with the contents of the synchronous prescaler counter to determine if Alarm B is to be activated. Only bits 0 up to MASKSS-1 are compared.
func (r *registerRtc_alrmbssrType) GetSs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbssrFieldSsMask) >> RegisterRtc_alrmbssrFieldSsShift)
}

// SetSs Sub seconds value This value is compared with the contents of the synchronous prescaler counter to determine if Alarm B is to be activated. Only bits 0 up to MASKSS-1 are compared.
func (r *registerRtc_alrmbssrType) SetSs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbssrFieldSsMask)|(uint32(value)<<RegisterRtc_alrmbssrFieldSsShift))
}

const (
	RegisterRtc_alrmbssrFieldMaskssShift = 24
	RegisterRtc_alrmbssrFieldMaskssMask  = 0xf000000
)

// GetMaskss Mask the most-significant bits starting at this bit ... The overflow bits of the synchronous counter (bits 15) is never compared. This bit can be different from 0 only after a shift operation.
func (r *registerRtc_alrmbssrType) GetMaskss() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_alrmbssrFieldMaskssMask) >> RegisterRtc_alrmbssrFieldMaskssShift)
}

// SetMaskss Mask the most-significant bits starting at this bit ... The overflow bits of the synchronous counter (bits 15) is never compared. This bit can be different from 0 only after a shift operation.
func (r *registerRtc_alrmbssrType) SetMaskss(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_alrmbssrFieldMaskssMask)|(uint32(value)<<RegisterRtc_alrmbssrFieldMaskssShift))
}

// registerRtc_orType RTC option register
type registerRtc_orType uint32

const (
	RegisterRtc_orFieldRtc_alarm_typeShift = 0
	RegisterRtc_orFieldRtc_alarm_typeMask  = 0x1
)

// GetRtc_alarm_type RTC_ALARM output type on PC13
func (r *registerRtc_orType) GetRtc_alarm_type() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_orFieldRtc_alarm_typeMask) != 0
}

// SetRtc_alarm_type RTC_ALARM output type on PC13
func (r *registerRtc_orType) SetRtc_alarm_type(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_orFieldRtc_alarm_typeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_orFieldRtc_alarm_typeMask)
	}
}

const (
	RegisterRtc_orFieldRtc_out_rmpShift = 1
	RegisterRtc_orFieldRtc_out_rmpMask  = 0x2
)

// GetRtc_out_rmp RTC_OUT remap
func (r *registerRtc_orType) GetRtc_out_rmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtc_orFieldRtc_out_rmpMask) != 0
}

// SetRtc_out_rmp RTC_OUT remap
func (r *registerRtc_orType) SetRtc_out_rmp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtc_orFieldRtc_out_rmpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtc_orFieldRtc_out_rmpMask)
	}
}

// registerRtc_bkp0rType RTC backup registers
type registerRtc_bkp0rType uint32

const (
	RegisterRtc_bkp0rFieldBkpShift = 0
	RegisterRtc_bkp0rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp0rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp0rFieldBkpMask) >> RegisterRtc_bkp0rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp0rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp0rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp0rFieldBkpShift))
}

// registerRtc_bkp1rType RTC backup registers
type registerRtc_bkp1rType uint32

const (
	RegisterRtc_bkp1rFieldBkpShift = 0
	RegisterRtc_bkp1rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp1rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp1rFieldBkpMask) >> RegisterRtc_bkp1rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp1rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp1rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp1rFieldBkpShift))
}

// registerRtc_bkp2rType RTC backup registers
type registerRtc_bkp2rType uint32

const (
	RegisterRtc_bkp2rFieldBkpShift = 0
	RegisterRtc_bkp2rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp2rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp2rFieldBkpMask) >> RegisterRtc_bkp2rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp2rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp2rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp2rFieldBkpShift))
}

// registerRtc_bkp3rType RTC backup registers
type registerRtc_bkp3rType uint32

const (
	RegisterRtc_bkp3rFieldBkpShift = 0
	RegisterRtc_bkp3rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp3rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp3rFieldBkpMask) >> RegisterRtc_bkp3rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp3rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp3rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp3rFieldBkpShift))
}

// registerRtc_bkp4rType RTC backup registers
type registerRtc_bkp4rType uint32

const (
	RegisterRtc_bkp4rFieldBkpShift = 0
	RegisterRtc_bkp4rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp4rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp4rFieldBkpMask) >> RegisterRtc_bkp4rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp4rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp4rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp4rFieldBkpShift))
}

// registerRtc_bkp5rType RTC backup registers
type registerRtc_bkp5rType uint32

const (
	RegisterRtc_bkp5rFieldBkpShift = 0
	RegisterRtc_bkp5rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp5rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp5rFieldBkpMask) >> RegisterRtc_bkp5rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp5rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp5rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp5rFieldBkpShift))
}

// registerRtc_bkp6rType RTC backup registers
type registerRtc_bkp6rType uint32

const (
	RegisterRtc_bkp6rFieldBkpShift = 0
	RegisterRtc_bkp6rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp6rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp6rFieldBkpMask) >> RegisterRtc_bkp6rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp6rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp6rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp6rFieldBkpShift))
}

// registerRtc_bkp7rType RTC backup registers
type registerRtc_bkp7rType uint32

const (
	RegisterRtc_bkp7rFieldBkpShift = 0
	RegisterRtc_bkp7rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp7rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp7rFieldBkpMask) >> RegisterRtc_bkp7rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp7rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp7rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp7rFieldBkpShift))
}

// registerRtc_bkp8rType RTC backup registers
type registerRtc_bkp8rType uint32

const (
	RegisterRtc_bkp8rFieldBkpShift = 0
	RegisterRtc_bkp8rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp8rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp8rFieldBkpMask) >> RegisterRtc_bkp8rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp8rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp8rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp8rFieldBkpShift))
}

// registerRtc_bkp9rType RTC backup registers
type registerRtc_bkp9rType uint32

const (
	RegisterRtc_bkp9rFieldBkpShift = 0
	RegisterRtc_bkp9rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp9rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp9rFieldBkpMask) >> RegisterRtc_bkp9rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp9rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp9rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp9rFieldBkpShift))
}

// registerRtc_bkp10rType RTC backup registers
type registerRtc_bkp10rType uint32

const (
	RegisterRtc_bkp10rFieldBkpShift = 0
	RegisterRtc_bkp10rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp10rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp10rFieldBkpMask) >> RegisterRtc_bkp10rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp10rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp10rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp10rFieldBkpShift))
}

// registerRtc_bkp11rType RTC backup registers
type registerRtc_bkp11rType uint32

const (
	RegisterRtc_bkp11rFieldBkpShift = 0
	RegisterRtc_bkp11rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp11rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp11rFieldBkpMask) >> RegisterRtc_bkp11rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp11rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp11rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp11rFieldBkpShift))
}

// registerRtc_bkp12rType RTC backup registers
type registerRtc_bkp12rType uint32

const (
	RegisterRtc_bkp12rFieldBkpShift = 0
	RegisterRtc_bkp12rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp12rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp12rFieldBkpMask) >> RegisterRtc_bkp12rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp12rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp12rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp12rFieldBkpShift))
}

// registerRtc_bkp13rType RTC backup registers
type registerRtc_bkp13rType uint32

const (
	RegisterRtc_bkp13rFieldBkpShift = 0
	RegisterRtc_bkp13rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp13rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp13rFieldBkpMask) >> RegisterRtc_bkp13rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp13rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp13rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp13rFieldBkpShift))
}

// registerRtc_bkp14rType RTC backup registers
type registerRtc_bkp14rType uint32

const (
	RegisterRtc_bkp14rFieldBkpShift = 0
	RegisterRtc_bkp14rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp14rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp14rFieldBkpMask) >> RegisterRtc_bkp14rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp14rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp14rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp14rFieldBkpShift))
}

// registerRtc_bkp15rType RTC backup registers
type registerRtc_bkp15rType uint32

const (
	RegisterRtc_bkp15rFieldBkpShift = 0
	RegisterRtc_bkp15rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp15rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp15rFieldBkpMask) >> RegisterRtc_bkp15rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp15rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp15rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp15rFieldBkpShift))
}

// registerRtc_bkp16rType RTC backup registers
type registerRtc_bkp16rType uint32

const (
	RegisterRtc_bkp16rFieldBkpShift = 0
	RegisterRtc_bkp16rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp16rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp16rFieldBkpMask) >> RegisterRtc_bkp16rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp16rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp16rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp16rFieldBkpShift))
}

// registerRtc_bkp17rType RTC backup registers
type registerRtc_bkp17rType uint32

const (
	RegisterRtc_bkp17rFieldBkpShift = 0
	RegisterRtc_bkp17rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp17rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp17rFieldBkpMask) >> RegisterRtc_bkp17rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp17rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp17rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp17rFieldBkpShift))
}

// registerRtc_bkp18rType RTC backup registers
type registerRtc_bkp18rType uint32

const (
	RegisterRtc_bkp18rFieldBkpShift = 0
	RegisterRtc_bkp18rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp18rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp18rFieldBkpMask) >> RegisterRtc_bkp18rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp18rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp18rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp18rFieldBkpShift))
}

// registerRtc_bkp19rType RTC backup registers
type registerRtc_bkp19rType uint32

const (
	RegisterRtc_bkp19rFieldBkpShift = 0
	RegisterRtc_bkp19rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp19rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp19rFieldBkpMask) >> RegisterRtc_bkp19rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp19rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp19rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp19rFieldBkpShift))
}

// registerRtc_bkp20rType RTC backup registers
type registerRtc_bkp20rType uint32

const (
	RegisterRtc_bkp20rFieldBkpShift = 0
	RegisterRtc_bkp20rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp20rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp20rFieldBkpMask) >> RegisterRtc_bkp20rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp20rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp20rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp20rFieldBkpShift))
}

// registerRtc_bkp21rType RTC backup registers
type registerRtc_bkp21rType uint32

const (
	RegisterRtc_bkp21rFieldBkpShift = 0
	RegisterRtc_bkp21rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp21rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp21rFieldBkpMask) >> RegisterRtc_bkp21rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp21rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp21rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp21rFieldBkpShift))
}

// registerRtc_bkp22rType RTC backup registers
type registerRtc_bkp22rType uint32

const (
	RegisterRtc_bkp22rFieldBkpShift = 0
	RegisterRtc_bkp22rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp22rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp22rFieldBkpMask) >> RegisterRtc_bkp22rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp22rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp22rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp22rFieldBkpShift))
}

// registerRtc_bkp23rType RTC backup registers
type registerRtc_bkp23rType uint32

const (
	RegisterRtc_bkp23rFieldBkpShift = 0
	RegisterRtc_bkp23rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp23rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp23rFieldBkpMask) >> RegisterRtc_bkp23rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp23rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp23rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp23rFieldBkpShift))
}

// registerRtc_bkp24rType RTC backup registers
type registerRtc_bkp24rType uint32

const (
	RegisterRtc_bkp24rFieldBkpShift = 0
	RegisterRtc_bkp24rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp24rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp24rFieldBkpMask) >> RegisterRtc_bkp24rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp24rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp24rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp24rFieldBkpShift))
}

// registerRtc_bkp25rType RTC backup registers
type registerRtc_bkp25rType uint32

const (
	RegisterRtc_bkp25rFieldBkpShift = 0
	RegisterRtc_bkp25rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp25rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp25rFieldBkpMask) >> RegisterRtc_bkp25rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp25rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp25rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp25rFieldBkpShift))
}

// registerRtc_bkp26rType RTC backup registers
type registerRtc_bkp26rType uint32

const (
	RegisterRtc_bkp26rFieldBkpShift = 0
	RegisterRtc_bkp26rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp26rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp26rFieldBkpMask) >> RegisterRtc_bkp26rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp26rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp26rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp26rFieldBkpShift))
}

// registerRtc_bkp27rType RTC backup registers
type registerRtc_bkp27rType uint32

const (
	RegisterRtc_bkp27rFieldBkpShift = 0
	RegisterRtc_bkp27rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp27rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp27rFieldBkpMask) >> RegisterRtc_bkp27rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp27rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp27rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp27rFieldBkpShift))
}

// registerRtc_bkp28rType RTC backup registers
type registerRtc_bkp28rType uint32

const (
	RegisterRtc_bkp28rFieldBkpShift = 0
	RegisterRtc_bkp28rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp28rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp28rFieldBkpMask) >> RegisterRtc_bkp28rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp28rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp28rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp28rFieldBkpShift))
}

// registerRtc_bkp29rType RTC backup registers
type registerRtc_bkp29rType uint32

const (
	RegisterRtc_bkp29rFieldBkpShift = 0
	RegisterRtc_bkp29rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp29rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp29rFieldBkpMask) >> RegisterRtc_bkp29rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp29rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp29rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp29rFieldBkpShift))
}

// registerRtc_bkp30rType RTC backup registers
type registerRtc_bkp30rType uint32

const (
	RegisterRtc_bkp30rFieldBkpShift = 0
	RegisterRtc_bkp30rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp30rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp30rFieldBkpMask) >> RegisterRtc_bkp30rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp30rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp30rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp30rFieldBkpShift))
}

// registerRtc_bkp31rType RTC backup registers
type registerRtc_bkp31rType uint32

const (
	RegisterRtc_bkp31rFieldBkpShift = 0
	RegisterRtc_bkp31rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp31rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtc_bkp31rFieldBkpMask) >> RegisterRtc_bkp31rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *registerRtc_bkp31rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtc_bkp31rFieldBkpMask)|(uint32(value)<<RegisterRtc_bkp31rFieldBkpShift))
}
