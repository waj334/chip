//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package rtc

import (
	"unsafe"
	"volatile"
)

var (
	Rtc = (*_rtc)(unsafe.Pointer(uintptr(0x58004000)))
)

type _rtc struct {
	Rtctr       RegisterRtctrType
	Rtcdr       RegisterRtcdrType
	Rtccr       RegisterRtccrType
	Rtcisr      RegisterRtcisrType
	Rtcprer     RegisterRtcprerType
	Rtcwutr     RegisterRtcwutrType
	_           [4]byte
	Rtcalrmar   RegisterRtcalrmarType
	Rtcalrmbr   RegisterRtcalrmbrType
	Rtcwpr      RegisterRtcwprType
	Rtcssr      RegisterRtcssrType
	Rtcshiftr   RegisterRtcshiftrType
	Rtctstr     RegisterRtctstrType
	Rtctsdr     RegisterRtctsdrType
	Rtctsssr    RegisterRtctsssrType
	Rtccalr     RegisterRtccalrType
	Rtctampcr   RegisterRtctampcrType
	Rtcalrmassr RegisterRtcalrmassrType
	Rtcalrmbssr RegisterRtcalrmbssrType
	Rtcor       RegisterRtcorType
	Rtcbkp0r    RegisterRtcbkp0rType
	Rtcbkp1r    RegisterRtcbkp1rType
	Rtcbkp2r    RegisterRtcbkp2rType
	Rtcbkp3r    RegisterRtcbkp3rType
	Rtcbkp4r    RegisterRtcbkp4rType
	Rtcbkp5r    RegisterRtcbkp5rType
	Rtcbkp6r    RegisterRtcbkp6rType
	Rtcbkp7r    RegisterRtcbkp7rType
	Rtcbkp8r    RegisterRtcbkp8rType
	Rtcbkp9r    RegisterRtcbkp9rType
	Rtcbkp10r   RegisterRtcbkp10rType
	Rtcbkp11r   RegisterRtcbkp11rType
	Rtcbkp12r   RegisterRtcbkp12rType
	Rtcbkp13r   RegisterRtcbkp13rType
	Rtcbkp14r   RegisterRtcbkp14rType
	Rtcbkp15r   RegisterRtcbkp15rType
	Rtcbkp16r   RegisterRtcbkp16rType
	Rtcbkp17r   RegisterRtcbkp17rType
	Rtcbkp18r   RegisterRtcbkp18rType
	Rtcbkp19r   RegisterRtcbkp19rType
	Rtcbkp20r   RegisterRtcbkp20rType
	Rtcbkp21r   RegisterRtcbkp21rType
	Rtcbkp22r   RegisterRtcbkp22rType
	Rtcbkp23r   RegisterRtcbkp23rType
	Rtcbkp24r   RegisterRtcbkp24rType
	Rtcbkp25r   RegisterRtcbkp25rType
	Rtcbkp26r   RegisterRtcbkp26rType
	Rtcbkp27r   RegisterRtcbkp27rType
	Rtcbkp28r   RegisterRtcbkp28rType
	Rtcbkp29r   RegisterRtcbkp29rType
	Rtcbkp30r   RegisterRtcbkp30rType
	Rtcbkp31r   RegisterRtcbkp31rType
}

// RegisterRtctrType The RTC_TR is the calendar time shadow register. This register must be written in initialization mode only. Refer to Calendar initialization and configuration on page9 and Reading the calendar on page10.This register is write protected. The write access procedure is described in RTC register write protection on page9.
type RegisterRtctrType uint32

func (r *RegisterRtctrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtctrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtctrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtctrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtctrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtctrFieldSuShift = 0
	RegisterRtctrFieldSuMask  = 0xf
)

// GetSu Second units in BCD format
func (r *RegisterRtctrType) GetSu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctrFieldSuMask) >> RegisterRtctrFieldSuShift)
}

// SetSu Second units in BCD format
func (r *RegisterRtctrType) SetSu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctrFieldSuMask)|(uint32(value)<<RegisterRtctrFieldSuShift))
}

const (
	RegisterRtctrFieldStShift = 4
	RegisterRtctrFieldStMask  = 0x70
)

// GetSt Second tens in BCD format
func (r *RegisterRtctrType) GetSt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctrFieldStMask) >> RegisterRtctrFieldStShift)
}

// SetSt Second tens in BCD format
func (r *RegisterRtctrType) SetSt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctrFieldStMask)|(uint32(value)<<RegisterRtctrFieldStShift))
}

const (
	RegisterRtctrFieldMnuShift = 8
	RegisterRtctrFieldMnuMask  = 0xf00
)

// GetMnu Minute units in BCD format
func (r *RegisterRtctrType) GetMnu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctrFieldMnuMask) >> RegisterRtctrFieldMnuShift)
}

// SetMnu Minute units in BCD format
func (r *RegisterRtctrType) SetMnu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctrFieldMnuMask)|(uint32(value)<<RegisterRtctrFieldMnuShift))
}

const (
	RegisterRtctrFieldMntShift = 12
	RegisterRtctrFieldMntMask  = 0x7000
)

// GetMnt Minute tens in BCD format
func (r *RegisterRtctrType) GetMnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctrFieldMntMask) >> RegisterRtctrFieldMntShift)
}

// SetMnt Minute tens in BCD format
func (r *RegisterRtctrType) SetMnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctrFieldMntMask)|(uint32(value)<<RegisterRtctrFieldMntShift))
}

const (
	RegisterRtctrFieldHuShift = 16
	RegisterRtctrFieldHuMask  = 0xf0000
)

// GetHu Hour units in BCD format
func (r *RegisterRtctrType) GetHu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctrFieldHuMask) >> RegisterRtctrFieldHuShift)
}

// SetHu Hour units in BCD format
func (r *RegisterRtctrType) SetHu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctrFieldHuMask)|(uint32(value)<<RegisterRtctrFieldHuShift))
}

const (
	RegisterRtctrFieldHtShift = 20
	RegisterRtctrFieldHtMask  = 0x300000
)

// GetHt Hour tens in BCD format
func (r *RegisterRtctrType) GetHt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctrFieldHtMask) >> RegisterRtctrFieldHtShift)
}

// SetHt Hour tens in BCD format
func (r *RegisterRtctrType) SetHt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctrFieldHtMask)|(uint32(value)<<RegisterRtctrFieldHtShift))
}

const (
	RegisterRtctrFieldPmShift = 22
	RegisterRtctrFieldPmMask  = 0x400000
)

// GetPm AM/PM notation
func (r *RegisterRtctrType) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctrFieldPmMask) != 0
}

// SetPm AM/PM notation
func (r *RegisterRtctrType) SetPm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctrFieldPmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctrFieldPmMask)
	}
}

// RegisterRtcdrType The RTC_DR is the calendar date shadow register. This register must be written in initialization mode only. Refer to Calendar initialization and configuration on page9 and Reading the calendar on page10.This register is write protected. The write access procedure is described in RTC register write protection on page9.
type RegisterRtcdrType uint32

func (r *RegisterRtcdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcdrFieldDuShift = 0
	RegisterRtcdrFieldDuMask  = 0xf
)

// GetDu Date units in BCD format
func (r *RegisterRtcdrType) GetDu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcdrFieldDuMask) >> RegisterRtcdrFieldDuShift)
}

// SetDu Date units in BCD format
func (r *RegisterRtcdrType) SetDu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcdrFieldDuMask)|(uint32(value)<<RegisterRtcdrFieldDuShift))
}

const (
	RegisterRtcdrFieldDtShift = 4
	RegisterRtcdrFieldDtMask  = 0x30
)

// GetDt Date tens in BCD format
func (r *RegisterRtcdrType) GetDt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcdrFieldDtMask) >> RegisterRtcdrFieldDtShift)
}

// SetDt Date tens in BCD format
func (r *RegisterRtcdrType) SetDt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcdrFieldDtMask)|(uint32(value)<<RegisterRtcdrFieldDtShift))
}

const (
	RegisterRtcdrFieldMuShift = 8
	RegisterRtcdrFieldMuMask  = 0xf00
)

// GetMu Month units in BCD format
func (r *RegisterRtcdrType) GetMu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcdrFieldMuMask) >> RegisterRtcdrFieldMuShift)
}

// SetMu Month units in BCD format
func (r *RegisterRtcdrType) SetMu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcdrFieldMuMask)|(uint32(value)<<RegisterRtcdrFieldMuShift))
}

const (
	RegisterRtcdrFieldMtShift = 12
	RegisterRtcdrFieldMtMask  = 0x1000
)

// GetMt Month tens in BCD format
func (r *RegisterRtcdrType) GetMt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcdrFieldMtMask) != 0
}

// SetMt Month tens in BCD format
func (r *RegisterRtcdrType) SetMt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcdrFieldMtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcdrFieldMtMask)
	}
}

const (
	RegisterRtcdrFieldWduShift = 13
	RegisterRtcdrFieldWduMask  = 0xe000
)

// GetWdu Week day units
func (r *RegisterRtcdrType) GetWdu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcdrFieldWduMask) >> RegisterRtcdrFieldWduShift)
}

// SetWdu Week day units
func (r *RegisterRtcdrType) SetWdu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcdrFieldWduMask)|(uint32(value)<<RegisterRtcdrFieldWduShift))
}

const (
	RegisterRtcdrFieldYuShift = 16
	RegisterRtcdrFieldYuMask  = 0xf0000
)

// GetYu Year units in BCD format
func (r *RegisterRtcdrType) GetYu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcdrFieldYuMask) >> RegisterRtcdrFieldYuShift)
}

// SetYu Year units in BCD format
func (r *RegisterRtcdrType) SetYu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcdrFieldYuMask)|(uint32(value)<<RegisterRtcdrFieldYuShift))
}

const (
	RegisterRtcdrFieldYtShift = 20
	RegisterRtcdrFieldYtMask  = 0xf00000
)

// GetYt Year tens in BCD format
func (r *RegisterRtcdrType) GetYt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcdrFieldYtMask) >> RegisterRtcdrFieldYtShift)
}

// SetYt Year tens in BCD format
func (r *RegisterRtcdrType) SetYt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcdrFieldYtMask)|(uint32(value)<<RegisterRtcdrFieldYtShift))
}

// RegisterRtccrType RTC control register
type RegisterRtccrType uint32

func (r *RegisterRtccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtccrFieldWuckselShift = 0
	RegisterRtccrFieldWuckselMask  = 0x7
)

// GetWucksel Wakeup clock selection
func (r *RegisterRtccrType) GetWucksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldWuckselMask) >> RegisterRtccrFieldWuckselShift)
}

// SetWucksel Wakeup clock selection
func (r *RegisterRtccrType) SetWucksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldWuckselMask)|(uint32(value)<<RegisterRtccrFieldWuckselShift))
}

const (
	RegisterRtccrFieldTsedgeShift = 3
	RegisterRtccrFieldTsedgeMask  = 0x8
)

// GetTsedge Time-stamp event active edge TSE must be reset when TSEDGE is changed to avoid unwanted TSF setting.
func (r *RegisterRtccrType) GetTsedge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldTsedgeMask) != 0
}

// SetTsedge Time-stamp event active edge TSE must be reset when TSEDGE is changed to avoid unwanted TSF setting.
func (r *RegisterRtccrType) SetTsedge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldTsedgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldTsedgeMask)
	}
}

const (
	RegisterRtccrFieldRefckonShift = 4
	RegisterRtccrFieldRefckonMask  = 0x10
)

// GetRefckon RTC_REFIN reference clock detection enable (50 or 60Hz) Note: PREDIV_S must be 0x00FF.
func (r *RegisterRtccrType) GetRefckon() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldRefckonMask) != 0
}

// SetRefckon RTC_REFIN reference clock detection enable (50 or 60Hz) Note: PREDIV_S must be 0x00FF.
func (r *RegisterRtccrType) SetRefckon(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldRefckonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldRefckonMask)
	}
}

const (
	RegisterRtccrFieldBypshadShift = 5
	RegisterRtccrFieldBypshadMask  = 0x20
)

// GetBypshad Bypass the shadow registers Note: If the frequency of the APB clock is less than seven times the frequency of RTCCLK, BYPSHAD must be set to 1.
func (r *RegisterRtccrType) GetBypshad() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldBypshadMask) != 0
}

// SetBypshad Bypass the shadow registers Note: If the frequency of the APB clock is less than seven times the frequency of RTCCLK, BYPSHAD must be set to 1.
func (r *RegisterRtccrType) SetBypshad(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldBypshadMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldBypshadMask)
	}
}

const (
	RegisterRtccrFieldFmtShift = 6
	RegisterRtccrFieldFmtMask  = 0x40
)

// GetFmt Hour format
func (r *RegisterRtccrType) GetFmt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldFmtMask) != 0
}

// SetFmt Hour format
func (r *RegisterRtccrType) SetFmt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldFmtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldFmtMask)
	}
}

const (
	RegisterRtccrFieldAlraeShift = 8
	RegisterRtccrFieldAlraeMask  = 0x100
)

// GetAlrae Alarm A enable
func (r *RegisterRtccrType) GetAlrae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldAlraeMask) != 0
}

// SetAlrae Alarm A enable
func (r *RegisterRtccrType) SetAlrae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldAlraeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldAlraeMask)
	}
}

const (
	RegisterRtccrFieldAlrbeShift = 9
	RegisterRtccrFieldAlrbeMask  = 0x200
)

// GetAlrbe Alarm B enable
func (r *RegisterRtccrType) GetAlrbe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldAlrbeMask) != 0
}

// SetAlrbe Alarm B enable
func (r *RegisterRtccrType) SetAlrbe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldAlrbeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldAlrbeMask)
	}
}

const (
	RegisterRtccrFieldWuteShift = 10
	RegisterRtccrFieldWuteMask  = 0x400
)

// GetWute Wakeup timer enable
func (r *RegisterRtccrType) GetWute() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldWuteMask) != 0
}

// SetWute Wakeup timer enable
func (r *RegisterRtccrType) SetWute(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldWuteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldWuteMask)
	}
}

const (
	RegisterRtccrFieldTseShift = 11
	RegisterRtccrFieldTseMask  = 0x800
)

// GetTse timestamp enable
func (r *RegisterRtccrType) GetTse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldTseMask) != 0
}

// SetTse timestamp enable
func (r *RegisterRtccrType) SetTse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldTseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldTseMask)
	}
}

const (
	RegisterRtccrFieldAlraieShift = 12
	RegisterRtccrFieldAlraieMask  = 0x1000
)

// GetAlraie Alarm A interrupt enable
func (r *RegisterRtccrType) GetAlraie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldAlraieMask) != 0
}

// SetAlraie Alarm A interrupt enable
func (r *RegisterRtccrType) SetAlraie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldAlraieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldAlraieMask)
	}
}

const (
	RegisterRtccrFieldAlrbieShift = 13
	RegisterRtccrFieldAlrbieMask  = 0x2000
)

// GetAlrbie Alarm B interrupt enable
func (r *RegisterRtccrType) GetAlrbie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldAlrbieMask) != 0
}

// SetAlrbie Alarm B interrupt enable
func (r *RegisterRtccrType) SetAlrbie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldAlrbieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldAlrbieMask)
	}
}

const (
	RegisterRtccrFieldWutieShift = 14
	RegisterRtccrFieldWutieMask  = 0x4000
)

// GetWutie Wakeup timer interrupt enable
func (r *RegisterRtccrType) GetWutie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldWutieMask) != 0
}

// SetWutie Wakeup timer interrupt enable
func (r *RegisterRtccrType) SetWutie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldWutieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldWutieMask)
	}
}

const (
	RegisterRtccrFieldTsieShift = 15
	RegisterRtccrFieldTsieMask  = 0x8000
)

// GetTsie Time-stamp interrupt enable
func (r *RegisterRtccrType) GetTsie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldTsieMask) != 0
}

// SetTsie Time-stamp interrupt enable
func (r *RegisterRtccrType) SetTsie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldTsieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldTsieMask)
	}
}

const (
	RegisterRtccrFieldAdd1hShift = 16
	RegisterRtccrFieldAdd1hMask  = 0x10000
)

// SetAdd1h Add 1 hour (summer time change) When this bit is set outside initialization mode, 1 hour is added to the calendar time. This bit is always read as 0.
func (r *RegisterRtccrType) SetAdd1h(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldAdd1hMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldAdd1hMask)
	}
}

const (
	RegisterRtccrFieldSub1hShift = 17
	RegisterRtccrFieldSub1hMask  = 0x20000
)

// SetSub1h Subtract 1 hour (winter time change) When this bit is set outside initialization mode, 1 hour is subtracted to the calendar time if the current hour is not 0. This bit is always read as 0. Setting this bit has no effect when current hour is 0.
func (r *RegisterRtccrType) SetSub1h(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldSub1hMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldSub1hMask)
	}
}

const (
	RegisterRtccrFieldBkpShift = 18
	RegisterRtccrFieldBkpMask  = 0x40000
)

// GetBkp Backup This bit can be written by the user to memorize whether the daylight saving time change has been performed or not.
func (r *RegisterRtccrType) GetBkp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldBkpMask) != 0
}

// SetBkp Backup This bit can be written by the user to memorize whether the daylight saving time change has been performed or not.
func (r *RegisterRtccrType) SetBkp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldBkpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldBkpMask)
	}
}

const (
	RegisterRtccrFieldCoselShift = 19
	RegisterRtccrFieldCoselMask  = 0x80000
)

// GetCosel Calibration output selection When COE=1, this bit selects which signal is output on RTC_CALIB. These frequencies are valid for RTCCLK at 32.768 kHz and prescalers at their default values (PREDIV_A=127 and PREDIV_S=255). Refer to Section24.3.15: Calibration clock output
func (r *RegisterRtccrType) GetCosel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldCoselMask) != 0
}

// SetCosel Calibration output selection When COE=1, this bit selects which signal is output on RTC_CALIB. These frequencies are valid for RTCCLK at 32.768 kHz and prescalers at their default values (PREDIV_A=127 and PREDIV_S=255). Refer to Section24.3.15: Calibration clock output
func (r *RegisterRtccrType) SetCosel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldCoselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldCoselMask)
	}
}

const (
	RegisterRtccrFieldPolShift = 20
	RegisterRtccrFieldPolMask  = 0x100000
)

// GetPol Output polarity This bit is used to configure the polarity of RTC_ALARM output
func (r *RegisterRtccrType) GetPol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldPolMask) != 0
}

// SetPol Output polarity This bit is used to configure the polarity of RTC_ALARM output
func (r *RegisterRtccrType) SetPol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldPolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldPolMask)
	}
}

const (
	RegisterRtccrFieldOselShift = 21
	RegisterRtccrFieldOselMask  = 0x600000
)

// GetOsel Output selection These bits are used to select the flag to be routed to RTC_ALARM output
func (r *RegisterRtccrType) GetOsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldOselMask) >> RegisterRtccrFieldOselShift)
}

// SetOsel Output selection These bits are used to select the flag to be routed to RTC_ALARM output
func (r *RegisterRtccrType) SetOsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldOselMask)|(uint32(value)<<RegisterRtccrFieldOselShift))
}

const (
	RegisterRtccrFieldCoeShift = 23
	RegisterRtccrFieldCoeMask  = 0x800000
)

// GetCoe Calibration output enable This bit enables the RTC_CALIB output
func (r *RegisterRtccrType) GetCoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldCoeMask) != 0
}

// SetCoe Calibration output enable This bit enables the RTC_CALIB output
func (r *RegisterRtccrType) SetCoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldCoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldCoeMask)
	}
}

const (
	RegisterRtccrFieldItseShift = 24
	RegisterRtccrFieldItseMask  = 0x1000000
)

// GetItse timestamp on internal event enable
func (r *RegisterRtccrType) GetItse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccrFieldItseMask) != 0
}

// SetItse timestamp on internal event enable
func (r *RegisterRtccrType) SetItse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccrFieldItseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccrFieldItseMask)
	}
}

// RegisterRtcisrType This register is write protected (except for RTC_ISR[13:8] bits). The write access procedure is described in RTC register write protection on page9.
type RegisterRtcisrType uint32

func (r *RegisterRtcisrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcisrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcisrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcisrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcisrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcisrFieldAlrawfShift = 0
	RegisterRtcisrFieldAlrawfMask  = 0x1
)

// GetAlrawf Alarm A write flag This bit is set by hardware when Alarm A values can be changed, after the ALRAE bit has been set to 0 in RTC_CR. It is cleared by hardware in initialization mode.
func (r *RegisterRtcisrType) GetAlrawf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldAlrawfMask) != 0
}

const (
	RegisterRtcisrFieldAlrbwfShift = 1
	RegisterRtcisrFieldAlrbwfMask  = 0x2
)

// GetAlrbwf Alarm B write flag This bit is set by hardware when Alarm B values can be changed, after the ALRBE bit has been set to 0 in RTC_CR. It is cleared by hardware in initialization mode.
func (r *RegisterRtcisrType) GetAlrbwf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldAlrbwfMask) != 0
}

const (
	RegisterRtcisrFieldWutwfShift = 2
	RegisterRtcisrFieldWutwfMask  = 0x4
)

// GetWutwf Wakeup timer write flag This bit is set by hardware up to 2 RTCCLK cycles after the WUTE bit has been set to 0 in RTC_CR, and is cleared up to 2 RTCCLK cycles after the WUTE bit has been set to 1. The wakeup timer values can be changed when WUTE bit is cleared and WUTWF is set.
func (r *RegisterRtcisrType) GetWutwf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldWutwfMask) != 0
}

const (
	RegisterRtcisrFieldShpfShift = 3
	RegisterRtcisrFieldShpfMask  = 0x8
)

// GetShpf Shift operation pending This flag is set by hardware as soon as a shift operation is initiated by a write to the RTC_SHIFTR register. It is cleared by hardware when the corresponding shift operation has been executed. Writing to the SHPF bit has no effect.
func (r *RegisterRtcisrType) GetShpf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldShpfMask) != 0
}

const (
	RegisterRtcisrFieldInitsShift = 4
	RegisterRtcisrFieldInitsMask  = 0x10
)

// GetInits Initialization status flag This bit is set by hardware when the calendar year field is different from 0 (Backup domain reset state).
func (r *RegisterRtcisrType) GetInits() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldInitsMask) != 0
}

const (
	RegisterRtcisrFieldRsfShift = 5
	RegisterRtcisrFieldRsfMask  = 0x20
)

// GetRsf Registers synchronization flag This bit is set by hardware each time the calendar registers are copied into the shadow registers (RTC_SSRx, RTC_TRx and RTC_DRx). This bit is cleared by hardware in initialization mode, while a shift operation is pending (SHPF=1), or when in bypass shadow register mode (BYPSHAD=1). This bit can also be cleared by software. It is cleared either by software or by hardware in initialization mode.
func (r *RegisterRtcisrType) GetRsf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldRsfMask) != 0
}

// SetRsf Registers synchronization flag This bit is set by hardware each time the calendar registers are copied into the shadow registers (RTC_SSRx, RTC_TRx and RTC_DRx). This bit is cleared by hardware in initialization mode, while a shift operation is pending (SHPF=1), or when in bypass shadow register mode (BYPSHAD=1). This bit can also be cleared by software. It is cleared either by software or by hardware in initialization mode.
func (r *RegisterRtcisrType) SetRsf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcisrFieldRsfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcisrFieldRsfMask)
	}
}

const (
	RegisterRtcisrFieldInitfShift = 6
	RegisterRtcisrFieldInitfMask  = 0x40
)

// GetInitf Initialization flag When this bit is set to 1, the RTC is in initialization state, and the time, date and prescaler registers can be updated.
func (r *RegisterRtcisrType) GetInitf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldInitfMask) != 0
}

const (
	RegisterRtcisrFieldInitShift = 7
	RegisterRtcisrFieldInitMask  = 0x80
)

// GetInit Initialization mode
func (r *RegisterRtcisrType) GetInit() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldInitMask) != 0
}

// SetInit Initialization mode
func (r *RegisterRtcisrType) SetInit(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcisrFieldInitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcisrFieldInitMask)
	}
}

const (
	RegisterRtcisrFieldAlrafShift = 8
	RegisterRtcisrFieldAlrafMask  = 0x100
)

// GetAlraf Alarm A flag This flag is set by hardware when the time/date registers (RTC_TR and RTC_DR) match the Alarm A register (RTC_ALRMAR). This flag is cleared by software by writing 0.
func (r *RegisterRtcisrType) GetAlraf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldAlrafMask) != 0
}

// SetAlraf Alarm A flag This flag is set by hardware when the time/date registers (RTC_TR and RTC_DR) match the Alarm A register (RTC_ALRMAR). This flag is cleared by software by writing 0.
func (r *RegisterRtcisrType) SetAlraf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcisrFieldAlrafMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcisrFieldAlrafMask)
	}
}

const (
	RegisterRtcisrFieldAlrbfShift = 9
	RegisterRtcisrFieldAlrbfMask  = 0x200
)

// GetAlrbf Alarm B flag This flag is set by hardware when the time/date registers (RTC_TR and RTC_DR) match the Alarm B register (RTC_ALRMBR). This flag is cleared by software by writing 0.
func (r *RegisterRtcisrType) GetAlrbf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldAlrbfMask) != 0
}

// SetAlrbf Alarm B flag This flag is set by hardware when the time/date registers (RTC_TR and RTC_DR) match the Alarm B register (RTC_ALRMBR). This flag is cleared by software by writing 0.
func (r *RegisterRtcisrType) SetAlrbf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcisrFieldAlrbfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcisrFieldAlrbfMask)
	}
}

const (
	RegisterRtcisrFieldWutfShift = 10
	RegisterRtcisrFieldWutfMask  = 0x400
)

// GetWutf Wakeup timer flag This flag is set by hardware when the wakeup auto-reload counter reaches 0. This flag is cleared by software by writing 0. This flag must be cleared by software at least 1.5 RTCCLK periods before WUTF is set to 1 again.
func (r *RegisterRtcisrType) GetWutf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldWutfMask) != 0
}

// SetWutf Wakeup timer flag This flag is set by hardware when the wakeup auto-reload counter reaches 0. This flag is cleared by software by writing 0. This flag must be cleared by software at least 1.5 RTCCLK periods before WUTF is set to 1 again.
func (r *RegisterRtcisrType) SetWutf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcisrFieldWutfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcisrFieldWutfMask)
	}
}

const (
	RegisterRtcisrFieldTsfShift = 11
	RegisterRtcisrFieldTsfMask  = 0x800
)

// GetTsf Time-stamp flag This flag is set by hardware when a time-stamp event occurs. This flag is cleared by software by writing 0.
func (r *RegisterRtcisrType) GetTsf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldTsfMask) != 0
}

// SetTsf Time-stamp flag This flag is set by hardware when a time-stamp event occurs. This flag is cleared by software by writing 0.
func (r *RegisterRtcisrType) SetTsf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcisrFieldTsfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcisrFieldTsfMask)
	}
}

const (
	RegisterRtcisrFieldTsovfShift = 12
	RegisterRtcisrFieldTsovfMask  = 0x1000
)

// GetTsovf Time-stamp overflow flag This flag is set by hardware when a time-stamp event occurs while TSF is already set. This flag is cleared by software by writing 0. It is recommended to check and then clear TSOVF only after clearing the TSF bit. Otherwise, an overflow might not be noticed if a time-stamp event occurs immediately before the TSF bit is cleared.
func (r *RegisterRtcisrType) GetTsovf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldTsovfMask) != 0
}

// SetTsovf Time-stamp overflow flag This flag is set by hardware when a time-stamp event occurs while TSF is already set. This flag is cleared by software by writing 0. It is recommended to check and then clear TSOVF only after clearing the TSF bit. Otherwise, an overflow might not be noticed if a time-stamp event occurs immediately before the TSF bit is cleared.
func (r *RegisterRtcisrType) SetTsovf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcisrFieldTsovfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcisrFieldTsovfMask)
	}
}

const (
	RegisterRtcisrFieldTamp1fShift = 13
	RegisterRtcisrFieldTamp1fMask  = 0x2000
)

// GetTamp1f RTC_TAMP1 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP1 input. It is cleared by software writing 0
func (r *RegisterRtcisrType) GetTamp1f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldTamp1fMask) != 0
}

// SetTamp1f RTC_TAMP1 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP1 input. It is cleared by software writing 0
func (r *RegisterRtcisrType) SetTamp1f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcisrFieldTamp1fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcisrFieldTamp1fMask)
	}
}

const (
	RegisterRtcisrFieldTamp2fShift = 14
	RegisterRtcisrFieldTamp2fMask  = 0x4000
)

// GetTamp2f RTC_TAMP2 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP2 input. It is cleared by software writing 0
func (r *RegisterRtcisrType) GetTamp2f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldTamp2fMask) != 0
}

// SetTamp2f RTC_TAMP2 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP2 input. It is cleared by software writing 0
func (r *RegisterRtcisrType) SetTamp2f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcisrFieldTamp2fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcisrFieldTamp2fMask)
	}
}

const (
	RegisterRtcisrFieldTamp3fShift = 15
	RegisterRtcisrFieldTamp3fMask  = 0x8000
)

// GetTamp3f RTC_TAMP3 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP3 input. It is cleared by software writing 0
func (r *RegisterRtcisrType) GetTamp3f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldTamp3fMask) != 0
}

// SetTamp3f RTC_TAMP3 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP3 input. It is cleared by software writing 0
func (r *RegisterRtcisrType) SetTamp3f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcisrFieldTamp3fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcisrFieldTamp3fMask)
	}
}

const (
	RegisterRtcisrFieldRecalpfShift = 16
	RegisterRtcisrFieldRecalpfMask  = 0x10000
)

// GetRecalpf Recalibration pending Flag The RECALPF status flag is automatically set to 1 when software writes to the RTC_CALR register, indicating that the RTC_CALR register is blocked. When the new calibration settings are taken into account, this bit returns to 0. Refer to Re-calibration on-the-fly.
func (r *RegisterRtcisrType) GetRecalpf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldRecalpfMask) != 0
}

const (
	RegisterRtcisrFieldItsfShift = 17
	RegisterRtcisrFieldItsfMask  = 0x20000
)

// GetItsf Internal tTime-stamp flag
func (r *RegisterRtcisrType) GetItsf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcisrFieldItsfMask) != 0
}

// SetItsf Internal tTime-stamp flag
func (r *RegisterRtcisrType) SetItsf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcisrFieldItsfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcisrFieldItsfMask)
	}
}

// RegisterRtcprerType This register must be written in initialization mode only. The initialization must be performed in two separate write accesses. Refer to Calendar initialization and configuration on page9.This register is write protected. The write access procedure is described in RTC register write protection on page9.
type RegisterRtcprerType uint32

func (r *RegisterRtcprerType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcprerType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcprerType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcprerType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcprerType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcprerFieldPredivsShift = 0
	RegisterRtcprerFieldPredivsMask  = 0x7fff
)

// GetPredivs Synchronous prescaler factor This is the synchronous division factor: ck_spre frequency = ck_apre frequency/(PREDIV_S+1)
func (r *RegisterRtcprerType) GetPredivs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtcprerFieldPredivsMask) >> RegisterRtcprerFieldPredivsShift)
}

// SetPredivs Synchronous prescaler factor This is the synchronous division factor: ck_spre frequency = ck_apre frequency/(PREDIV_S+1)
func (r *RegisterRtcprerType) SetPredivs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcprerFieldPredivsMask)|(uint32(value)<<RegisterRtcprerFieldPredivsShift))
}

const (
	RegisterRtcprerFieldPredivaShift = 16
	RegisterRtcprerFieldPredivaMask  = 0x7f0000
)

// GetPrediva Asynchronous prescaler factor This is the asynchronous division factor: ck_apre frequency = RTCCLK frequency/(PREDIV_A+1)
func (r *RegisterRtcprerType) GetPrediva() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcprerFieldPredivaMask) >> RegisterRtcprerFieldPredivaShift)
}

// SetPrediva Asynchronous prescaler factor This is the asynchronous division factor: ck_apre frequency = RTCCLK frequency/(PREDIV_A+1)
func (r *RegisterRtcprerType) SetPrediva(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcprerFieldPredivaMask)|(uint32(value)<<RegisterRtcprerFieldPredivaShift))
}

// RegisterRtcwutrType This register can be written only when WUTWF is set to 1 in RTC_ISR.This register is write protected. The write access procedure is described in RTC register write protection on page9.
type RegisterRtcwutrType uint32

func (r *RegisterRtcwutrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcwutrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcwutrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcwutrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcwutrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcwutrFieldWutShift = 0
	RegisterRtcwutrFieldWutMask  = 0xffff
)

// GetWut Wakeup auto-reload value bits When the wakeup timer is enabled (WUTE set to 1), the WUTF flag is set every (WUT[15:0] + 1) ck_wut cycles. The ck_wut period is selected through WUCKSEL[2:0] bits of the RTC_CR register When WUCKSEL[2] = 1, the wakeup timer becomes 17-bits and WUCKSEL[1] effectively becomes WUT[16] the most-significant bit to be reloaded into the timer. The first assertion of WUTF occurs (WUT+1) ck_wut cycles after WUTE is set. Setting WUT[15:0] to 0x0000 with WUCKSEL[2:0] =011 (RTCCLK/2) is forbidden.
func (r *RegisterRtcwutrType) GetWut() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtcwutrFieldWutMask) >> RegisterRtcwutrFieldWutShift)
}

// SetWut Wakeup auto-reload value bits When the wakeup timer is enabled (WUTE set to 1), the WUTF flag is set every (WUT[15:0] + 1) ck_wut cycles. The ck_wut period is selected through WUCKSEL[2:0] bits of the RTC_CR register When WUCKSEL[2] = 1, the wakeup timer becomes 17-bits and WUCKSEL[1] effectively becomes WUT[16] the most-significant bit to be reloaded into the timer. The first assertion of WUTF occurs (WUT+1) ck_wut cycles after WUTE is set. Setting WUT[15:0] to 0x0000 with WUCKSEL[2:0] =011 (RTCCLK/2) is forbidden.
func (r *RegisterRtcwutrType) SetWut(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcwutrFieldWutMask)|(uint32(value)<<RegisterRtcwutrFieldWutShift))
}

// RegisterRtcalrmarType This register can be written only when ALRAWF is set to 1 in RTC_ISR, or in initialization mode.This register is write protected. The write access procedure is described in RTC register write protection on page9.
type RegisterRtcalrmarType uint32

func (r *RegisterRtcalrmarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcalrmarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcalrmarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcalrmarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcalrmarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcalrmarFieldSuShift = 0
	RegisterRtcalrmarFieldSuMask  = 0xf
)

// GetSu Second units in BCD format.
func (r *RegisterRtcalrmarType) GetSu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldSuMask) >> RegisterRtcalrmarFieldSuShift)
}

// SetSu Second units in BCD format.
func (r *RegisterRtcalrmarType) SetSu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldSuMask)|(uint32(value)<<RegisterRtcalrmarFieldSuShift))
}

const (
	RegisterRtcalrmarFieldStShift = 4
	RegisterRtcalrmarFieldStMask  = 0x70
)

// GetSt Second tens in BCD format.
func (r *RegisterRtcalrmarType) GetSt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldStMask) >> RegisterRtcalrmarFieldStShift)
}

// SetSt Second tens in BCD format.
func (r *RegisterRtcalrmarType) SetSt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldStMask)|(uint32(value)<<RegisterRtcalrmarFieldStShift))
}

const (
	RegisterRtcalrmarFieldMsk1Shift = 7
	RegisterRtcalrmarFieldMsk1Mask  = 0x80
)

// GetMsk1 Alarm A seconds mask
func (r *RegisterRtcalrmarType) GetMsk1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldMsk1Mask) != 0
}

// SetMsk1 Alarm A seconds mask
func (r *RegisterRtcalrmarType) SetMsk1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcalrmarFieldMsk1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldMsk1Mask)
	}
}

const (
	RegisterRtcalrmarFieldMnuShift = 8
	RegisterRtcalrmarFieldMnuMask  = 0xf00
)

// GetMnu Minute units in BCD format.
func (r *RegisterRtcalrmarType) GetMnu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldMnuMask) >> RegisterRtcalrmarFieldMnuShift)
}

// SetMnu Minute units in BCD format.
func (r *RegisterRtcalrmarType) SetMnu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldMnuMask)|(uint32(value)<<RegisterRtcalrmarFieldMnuShift))
}

const (
	RegisterRtcalrmarFieldMntShift = 12
	RegisterRtcalrmarFieldMntMask  = 0x7000
)

// GetMnt Minute tens in BCD format.
func (r *RegisterRtcalrmarType) GetMnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldMntMask) >> RegisterRtcalrmarFieldMntShift)
}

// SetMnt Minute tens in BCD format.
func (r *RegisterRtcalrmarType) SetMnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldMntMask)|(uint32(value)<<RegisterRtcalrmarFieldMntShift))
}

const (
	RegisterRtcalrmarFieldMsk2Shift = 15
	RegisterRtcalrmarFieldMsk2Mask  = 0x8000
)

// GetMsk2 Alarm A minutes mask
func (r *RegisterRtcalrmarType) GetMsk2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldMsk2Mask) != 0
}

// SetMsk2 Alarm A minutes mask
func (r *RegisterRtcalrmarType) SetMsk2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcalrmarFieldMsk2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldMsk2Mask)
	}
}

const (
	RegisterRtcalrmarFieldHuShift = 16
	RegisterRtcalrmarFieldHuMask  = 0xf0000
)

// GetHu Hour units in BCD format.
func (r *RegisterRtcalrmarType) GetHu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldHuMask) >> RegisterRtcalrmarFieldHuShift)
}

// SetHu Hour units in BCD format.
func (r *RegisterRtcalrmarType) SetHu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldHuMask)|(uint32(value)<<RegisterRtcalrmarFieldHuShift))
}

const (
	RegisterRtcalrmarFieldHtShift = 20
	RegisterRtcalrmarFieldHtMask  = 0x300000
)

// GetHt Hour tens in BCD format.
func (r *RegisterRtcalrmarType) GetHt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldHtMask) >> RegisterRtcalrmarFieldHtShift)
}

// SetHt Hour tens in BCD format.
func (r *RegisterRtcalrmarType) SetHt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldHtMask)|(uint32(value)<<RegisterRtcalrmarFieldHtShift))
}

const (
	RegisterRtcalrmarFieldPmShift = 22
	RegisterRtcalrmarFieldPmMask  = 0x400000
)

// GetPm AM/PM notation
func (r *RegisterRtcalrmarType) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldPmMask) != 0
}

// SetPm AM/PM notation
func (r *RegisterRtcalrmarType) SetPm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcalrmarFieldPmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldPmMask)
	}
}

const (
	RegisterRtcalrmarFieldMsk3Shift = 23
	RegisterRtcalrmarFieldMsk3Mask  = 0x800000
)

// GetMsk3 Alarm A hours mask
func (r *RegisterRtcalrmarType) GetMsk3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldMsk3Mask) != 0
}

// SetMsk3 Alarm A hours mask
func (r *RegisterRtcalrmarType) SetMsk3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcalrmarFieldMsk3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldMsk3Mask)
	}
}

const (
	RegisterRtcalrmarFieldDuShift = 24
	RegisterRtcalrmarFieldDuMask  = 0xf000000
)

// GetDu Date units or day in BCD format.
func (r *RegisterRtcalrmarType) GetDu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldDuMask) >> RegisterRtcalrmarFieldDuShift)
}

// SetDu Date units or day in BCD format.
func (r *RegisterRtcalrmarType) SetDu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldDuMask)|(uint32(value)<<RegisterRtcalrmarFieldDuShift))
}

const (
	RegisterRtcalrmarFieldDtShift = 28
	RegisterRtcalrmarFieldDtMask  = 0x30000000
)

// GetDt Date tens in BCD format.
func (r *RegisterRtcalrmarType) GetDt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldDtMask) >> RegisterRtcalrmarFieldDtShift)
}

// SetDt Date tens in BCD format.
func (r *RegisterRtcalrmarType) SetDt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldDtMask)|(uint32(value)<<RegisterRtcalrmarFieldDtShift))
}

const (
	RegisterRtcalrmarFieldWdselShift = 30
	RegisterRtcalrmarFieldWdselMask  = 0x40000000
)

// GetWdsel Week day selection
func (r *RegisterRtcalrmarType) GetWdsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldWdselMask) != 0
}

// SetWdsel Week day selection
func (r *RegisterRtcalrmarType) SetWdsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcalrmarFieldWdselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldWdselMask)
	}
}

const (
	RegisterRtcalrmarFieldMsk4Shift = 31
	RegisterRtcalrmarFieldMsk4Mask  = 0x80000000
)

// GetMsk4 Alarm A date mask
func (r *RegisterRtcalrmarType) GetMsk4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmarFieldMsk4Mask) != 0
}

// SetMsk4 Alarm A date mask
func (r *RegisterRtcalrmarType) SetMsk4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcalrmarFieldMsk4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmarFieldMsk4Mask)
	}
}

// RegisterRtcalrmbrType This register can be written only when ALRBWF is set to 1 in RTC_ISR, or in initialization mode.This register is write protected. The write access procedure is described in RTC register write protection on page9.
type RegisterRtcalrmbrType uint32

func (r *RegisterRtcalrmbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcalrmbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcalrmbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcalrmbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcalrmbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcalrmbrFieldSuShift = 0
	RegisterRtcalrmbrFieldSuMask  = 0xf
)

// GetSu Second units in BCD format
func (r *RegisterRtcalrmbrType) GetSu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldSuMask) >> RegisterRtcalrmbrFieldSuShift)
}

// SetSu Second units in BCD format
func (r *RegisterRtcalrmbrType) SetSu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldSuMask)|(uint32(value)<<RegisterRtcalrmbrFieldSuShift))
}

const (
	RegisterRtcalrmbrFieldStShift = 4
	RegisterRtcalrmbrFieldStMask  = 0x70
)

// GetSt Second tens in BCD format
func (r *RegisterRtcalrmbrType) GetSt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldStMask) >> RegisterRtcalrmbrFieldStShift)
}

// SetSt Second tens in BCD format
func (r *RegisterRtcalrmbrType) SetSt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldStMask)|(uint32(value)<<RegisterRtcalrmbrFieldStShift))
}

const (
	RegisterRtcalrmbrFieldMsk1Shift = 7
	RegisterRtcalrmbrFieldMsk1Mask  = 0x80
)

// GetMsk1 Alarm B seconds mask
func (r *RegisterRtcalrmbrType) GetMsk1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldMsk1Mask) != 0
}

// SetMsk1 Alarm B seconds mask
func (r *RegisterRtcalrmbrType) SetMsk1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcalrmbrFieldMsk1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldMsk1Mask)
	}
}

const (
	RegisterRtcalrmbrFieldMnuShift = 8
	RegisterRtcalrmbrFieldMnuMask  = 0xf00
)

// GetMnu Minute units in BCD format
func (r *RegisterRtcalrmbrType) GetMnu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldMnuMask) >> RegisterRtcalrmbrFieldMnuShift)
}

// SetMnu Minute units in BCD format
func (r *RegisterRtcalrmbrType) SetMnu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldMnuMask)|(uint32(value)<<RegisterRtcalrmbrFieldMnuShift))
}

const (
	RegisterRtcalrmbrFieldMntShift = 12
	RegisterRtcalrmbrFieldMntMask  = 0x7000
)

// GetMnt Minute tens in BCD format
func (r *RegisterRtcalrmbrType) GetMnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldMntMask) >> RegisterRtcalrmbrFieldMntShift)
}

// SetMnt Minute tens in BCD format
func (r *RegisterRtcalrmbrType) SetMnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldMntMask)|(uint32(value)<<RegisterRtcalrmbrFieldMntShift))
}

const (
	RegisterRtcalrmbrFieldMsk2Shift = 15
	RegisterRtcalrmbrFieldMsk2Mask  = 0x8000
)

// GetMsk2 Alarm B minutes mask
func (r *RegisterRtcalrmbrType) GetMsk2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldMsk2Mask) != 0
}

// SetMsk2 Alarm B minutes mask
func (r *RegisterRtcalrmbrType) SetMsk2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcalrmbrFieldMsk2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldMsk2Mask)
	}
}

const (
	RegisterRtcalrmbrFieldHuShift = 16
	RegisterRtcalrmbrFieldHuMask  = 0xf0000
)

// GetHu Hour units in BCD format
func (r *RegisterRtcalrmbrType) GetHu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldHuMask) >> RegisterRtcalrmbrFieldHuShift)
}

// SetHu Hour units in BCD format
func (r *RegisterRtcalrmbrType) SetHu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldHuMask)|(uint32(value)<<RegisterRtcalrmbrFieldHuShift))
}

const (
	RegisterRtcalrmbrFieldHtShift = 20
	RegisterRtcalrmbrFieldHtMask  = 0x300000
)

// GetHt Hour tens in BCD format
func (r *RegisterRtcalrmbrType) GetHt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldHtMask) >> RegisterRtcalrmbrFieldHtShift)
}

// SetHt Hour tens in BCD format
func (r *RegisterRtcalrmbrType) SetHt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldHtMask)|(uint32(value)<<RegisterRtcalrmbrFieldHtShift))
}

const (
	RegisterRtcalrmbrFieldPmShift = 22
	RegisterRtcalrmbrFieldPmMask  = 0x400000
)

// GetPm AM/PM notation
func (r *RegisterRtcalrmbrType) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldPmMask) != 0
}

// SetPm AM/PM notation
func (r *RegisterRtcalrmbrType) SetPm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcalrmbrFieldPmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldPmMask)
	}
}

const (
	RegisterRtcalrmbrFieldMsk3Shift = 23
	RegisterRtcalrmbrFieldMsk3Mask  = 0x800000
)

// GetMsk3 Alarm B hours mask
func (r *RegisterRtcalrmbrType) GetMsk3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldMsk3Mask) != 0
}

// SetMsk3 Alarm B hours mask
func (r *RegisterRtcalrmbrType) SetMsk3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcalrmbrFieldMsk3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldMsk3Mask)
	}
}

const (
	RegisterRtcalrmbrFieldDuShift = 24
	RegisterRtcalrmbrFieldDuMask  = 0xf000000
)

// GetDu Date units or day in BCD format
func (r *RegisterRtcalrmbrType) GetDu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldDuMask) >> RegisterRtcalrmbrFieldDuShift)
}

// SetDu Date units or day in BCD format
func (r *RegisterRtcalrmbrType) SetDu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldDuMask)|(uint32(value)<<RegisterRtcalrmbrFieldDuShift))
}

const (
	RegisterRtcalrmbrFieldDtShift = 28
	RegisterRtcalrmbrFieldDtMask  = 0x30000000
)

// GetDt Date tens in BCD format
func (r *RegisterRtcalrmbrType) GetDt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldDtMask) >> RegisterRtcalrmbrFieldDtShift)
}

// SetDt Date tens in BCD format
func (r *RegisterRtcalrmbrType) SetDt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldDtMask)|(uint32(value)<<RegisterRtcalrmbrFieldDtShift))
}

const (
	RegisterRtcalrmbrFieldWdselShift = 30
	RegisterRtcalrmbrFieldWdselMask  = 0x40000000
)

// GetWdsel Week day selection
func (r *RegisterRtcalrmbrType) GetWdsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldWdselMask) != 0
}

// SetWdsel Week day selection
func (r *RegisterRtcalrmbrType) SetWdsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcalrmbrFieldWdselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldWdselMask)
	}
}

const (
	RegisterRtcalrmbrFieldMsk4Shift = 31
	RegisterRtcalrmbrFieldMsk4Mask  = 0x80000000
)

// GetMsk4 Alarm B date mask
func (r *RegisterRtcalrmbrType) GetMsk4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbrFieldMsk4Mask) != 0
}

// SetMsk4 Alarm B date mask
func (r *RegisterRtcalrmbrType) SetMsk4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcalrmbrFieldMsk4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbrFieldMsk4Mask)
	}
}

// RegisterRtcwprType RTC write protection register
type RegisterRtcwprType uint32

func (r *RegisterRtcwprType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcwprType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcwprType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcwprType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcwprType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcwprFieldKeyShift = 0
	RegisterRtcwprFieldKeyMask  = 0xff
)

// GetKey Write protection key This byte is written by software. Reading this byte always returns 0x00. Refer to RTC register write protection for a description of how to unlock RTC register write protection.
func (r *RegisterRtcwprType) GetKey() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcwprFieldKeyMask) >> RegisterRtcwprFieldKeyShift)
}

// SetKey Write protection key This byte is written by software. Reading this byte always returns 0x00. Refer to RTC register write protection for a description of how to unlock RTC register write protection.
func (r *RegisterRtcwprType) SetKey(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcwprFieldKeyMask)|(uint32(value)<<RegisterRtcwprFieldKeyShift))
}

// RegisterRtcssrType RTC sub second register
type RegisterRtcssrType uint32

func (r *RegisterRtcssrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcssrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcssrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcssrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcssrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcssrFieldSsShift = 0
	RegisterRtcssrFieldSsMask  = 0xffff
)

// GetSs Sub second value SS[15:0] is the value in the synchronous prescaler counter. The fraction of a second is given by the formula below: Second fraction = (PREDIV_S - SS) / (PREDIV_S + 1) Note: SS can be larger than PREDIV_S only after a shift operation. In that case, the correct time/date is one second less than as indicated by RTC_TR/RTC_DR.
func (r *RegisterRtcssrType) GetSs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtcssrFieldSsMask) >> RegisterRtcssrFieldSsShift)
}

// SetSs Sub second value SS[15:0] is the value in the synchronous prescaler counter. The fraction of a second is given by the formula below: Second fraction = (PREDIV_S - SS) / (PREDIV_S + 1) Note: SS can be larger than PREDIV_S only after a shift operation. In that case, the correct time/date is one second less than as indicated by RTC_TR/RTC_DR.
func (r *RegisterRtcssrType) SetSs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcssrFieldSsMask)|(uint32(value)<<RegisterRtcssrFieldSsShift))
}

// RegisterRtcshiftrType This register is write protected. The write access procedure is described in RTC register write protection on page9.
type RegisterRtcshiftrType uint32

func (r *RegisterRtcshiftrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcshiftrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcshiftrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcshiftrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcshiftrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcshiftrFieldSubfsShift = 0
	RegisterRtcshiftrFieldSubfsMask  = 0x7fff
)

// GetSubfs Subtract a fraction of a second These bits are write only and is always read as zero. Writing to this bit has no effect when a shift operation is pending (when SHPF=1, in RTC_ISR). The value which is written to SUBFS is added to the synchronous prescaler counter. Since this counter counts down, this operation effectively subtracts from (delays) the clock by: Delay (seconds) = SUBFS / (PREDIV_S + 1) A fraction of a second can effectively be added to the clock (advancing the clock) when the ADD1S function is used in conjunction with SUBFS, effectively advancing the clock by: Advance (seconds) = (1 - (SUBFS / (PREDIV_S + 1))). Note: Writing to SUBFS causes RSF to be cleared. Software can then wait until RSF=1 to be sure that the shadow registers have been updated with the shifted time.
func (r *RegisterRtcshiftrType) GetSubfs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtcshiftrFieldSubfsMask) >> RegisterRtcshiftrFieldSubfsShift)
}

// SetSubfs Subtract a fraction of a second These bits are write only and is always read as zero. Writing to this bit has no effect when a shift operation is pending (when SHPF=1, in RTC_ISR). The value which is written to SUBFS is added to the synchronous prescaler counter. Since this counter counts down, this operation effectively subtracts from (delays) the clock by: Delay (seconds) = SUBFS / (PREDIV_S + 1) A fraction of a second can effectively be added to the clock (advancing the clock) when the ADD1S function is used in conjunction with SUBFS, effectively advancing the clock by: Advance (seconds) = (1 - (SUBFS / (PREDIV_S + 1))). Note: Writing to SUBFS causes RSF to be cleared. Software can then wait until RSF=1 to be sure that the shadow registers have been updated with the shifted time.
func (r *RegisterRtcshiftrType) SetSubfs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcshiftrFieldSubfsMask)|(uint32(value)<<RegisterRtcshiftrFieldSubfsShift))
}

const (
	RegisterRtcshiftrFieldAdd1sShift = 31
	RegisterRtcshiftrFieldAdd1sMask  = 0x80000000
)

// GetAdd1s Add one second This bit is write only and is always read as zero. Writing to this bit has no effect when a shift operation is pending (when SHPF=1, in RTC_ISR). This function is intended to be used with SUBFS (see description below) in order to effectively add a fraction of a second to the clock in an atomic operation.
func (r *RegisterRtcshiftrType) GetAdd1s() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcshiftrFieldAdd1sMask) != 0
}

// SetAdd1s Add one second This bit is write only and is always read as zero. Writing to this bit has no effect when a shift operation is pending (when SHPF=1, in RTC_ISR). This function is intended to be used with SUBFS (see description below) in order to effectively add a fraction of a second to the clock in an atomic operation.
func (r *RegisterRtcshiftrType) SetAdd1s(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcshiftrFieldAdd1sMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcshiftrFieldAdd1sMask)
	}
}

// RegisterRtctstrType The content of this register is valid only when TSF is set to 1 in RTC_ISR. It is cleared when TSF bit is reset.
type RegisterRtctstrType uint32

func (r *RegisterRtctstrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtctstrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtctstrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtctstrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtctstrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtctstrFieldSuShift = 0
	RegisterRtctstrFieldSuMask  = 0xf
)

// GetSu Second units in BCD format.
func (r *RegisterRtctstrType) GetSu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctstrFieldSuMask) >> RegisterRtctstrFieldSuShift)
}

// SetSu Second units in BCD format.
func (r *RegisterRtctstrType) SetSu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctstrFieldSuMask)|(uint32(value)<<RegisterRtctstrFieldSuShift))
}

const (
	RegisterRtctstrFieldStShift = 4
	RegisterRtctstrFieldStMask  = 0x70
)

// GetSt Second tens in BCD format.
func (r *RegisterRtctstrType) GetSt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctstrFieldStMask) >> RegisterRtctstrFieldStShift)
}

// SetSt Second tens in BCD format.
func (r *RegisterRtctstrType) SetSt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctstrFieldStMask)|(uint32(value)<<RegisterRtctstrFieldStShift))
}

const (
	RegisterRtctstrFieldMnuShift = 8
	RegisterRtctstrFieldMnuMask  = 0xf00
)

// GetMnu Minute units in BCD format.
func (r *RegisterRtctstrType) GetMnu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctstrFieldMnuMask) >> RegisterRtctstrFieldMnuShift)
}

// SetMnu Minute units in BCD format.
func (r *RegisterRtctstrType) SetMnu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctstrFieldMnuMask)|(uint32(value)<<RegisterRtctstrFieldMnuShift))
}

const (
	RegisterRtctstrFieldMntShift = 12
	RegisterRtctstrFieldMntMask  = 0x7000
)

// GetMnt Minute tens in BCD format.
func (r *RegisterRtctstrType) GetMnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctstrFieldMntMask) >> RegisterRtctstrFieldMntShift)
}

// SetMnt Minute tens in BCD format.
func (r *RegisterRtctstrType) SetMnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctstrFieldMntMask)|(uint32(value)<<RegisterRtctstrFieldMntShift))
}

const (
	RegisterRtctstrFieldHuShift = 16
	RegisterRtctstrFieldHuMask  = 0xf0000
)

// GetHu Hour units in BCD format.
func (r *RegisterRtctstrType) GetHu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctstrFieldHuMask) >> RegisterRtctstrFieldHuShift)
}

// SetHu Hour units in BCD format.
func (r *RegisterRtctstrType) SetHu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctstrFieldHuMask)|(uint32(value)<<RegisterRtctstrFieldHuShift))
}

const (
	RegisterRtctstrFieldHtShift = 20
	RegisterRtctstrFieldHtMask  = 0x300000
)

// GetHt Hour tens in BCD format.
func (r *RegisterRtctstrType) GetHt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctstrFieldHtMask) >> RegisterRtctstrFieldHtShift)
}

// SetHt Hour tens in BCD format.
func (r *RegisterRtctstrType) SetHt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctstrFieldHtMask)|(uint32(value)<<RegisterRtctstrFieldHtShift))
}

const (
	RegisterRtctstrFieldPmShift = 22
	RegisterRtctstrFieldPmMask  = 0x400000
)

// GetPm AM/PM notation
func (r *RegisterRtctstrType) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctstrFieldPmMask) != 0
}

// SetPm AM/PM notation
func (r *RegisterRtctstrType) SetPm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctstrFieldPmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctstrFieldPmMask)
	}
}

// RegisterRtctsdrType The content of this register is valid only when TSF is set to 1 in RTC_ISR. It is cleared when TSF bit is reset.
type RegisterRtctsdrType uint32

func (r *RegisterRtctsdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtctsdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtctsdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtctsdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtctsdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtctsdrFieldDuShift = 0
	RegisterRtctsdrFieldDuMask  = 0xf
)

// GetDu Date units in BCD format
func (r *RegisterRtctsdrType) GetDu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctsdrFieldDuMask) >> RegisterRtctsdrFieldDuShift)
}

// SetDu Date units in BCD format
func (r *RegisterRtctsdrType) SetDu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctsdrFieldDuMask)|(uint32(value)<<RegisterRtctsdrFieldDuShift))
}

const (
	RegisterRtctsdrFieldDtShift = 4
	RegisterRtctsdrFieldDtMask  = 0x30
)

// GetDt Date tens in BCD format
func (r *RegisterRtctsdrType) GetDt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctsdrFieldDtMask) >> RegisterRtctsdrFieldDtShift)
}

// SetDt Date tens in BCD format
func (r *RegisterRtctsdrType) SetDt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctsdrFieldDtMask)|(uint32(value)<<RegisterRtctsdrFieldDtShift))
}

const (
	RegisterRtctsdrFieldMuShift = 8
	RegisterRtctsdrFieldMuMask  = 0xf00
)

// GetMu Month units in BCD format
func (r *RegisterRtctsdrType) GetMu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctsdrFieldMuMask) >> RegisterRtctsdrFieldMuShift)
}

// SetMu Month units in BCD format
func (r *RegisterRtctsdrType) SetMu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctsdrFieldMuMask)|(uint32(value)<<RegisterRtctsdrFieldMuShift))
}

const (
	RegisterRtctsdrFieldMtShift = 12
	RegisterRtctsdrFieldMtMask  = 0x1000
)

// GetMt Month tens in BCD format
func (r *RegisterRtctsdrType) GetMt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctsdrFieldMtMask) != 0
}

// SetMt Month tens in BCD format
func (r *RegisterRtctsdrType) SetMt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctsdrFieldMtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctsdrFieldMtMask)
	}
}

const (
	RegisterRtctsdrFieldWduShift = 13
	RegisterRtctsdrFieldWduMask  = 0xe000
)

// GetWdu Week day units
func (r *RegisterRtctsdrType) GetWdu() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctsdrFieldWduMask) >> RegisterRtctsdrFieldWduShift)
}

// SetWdu Week day units
func (r *RegisterRtctsdrType) SetWdu(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctsdrFieldWduMask)|(uint32(value)<<RegisterRtctsdrFieldWduShift))
}

// RegisterRtctsssrType The content of this register is valid only when RTC_ISR/TSF is set. It is cleared when the RTC_ISR/TSF bit is reset.
type RegisterRtctsssrType uint32

func (r *RegisterRtctsssrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtctsssrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtctsssrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtctsssrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtctsssrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtctsssrFieldSsShift = 0
	RegisterRtctsssrFieldSsMask  = 0xffff
)

// GetSs Sub second value SS[15:0] is the value of the synchronous prescaler counter when the timestamp event occurred.
func (r *RegisterRtctsssrType) GetSs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtctsssrFieldSsMask) >> RegisterRtctsssrFieldSsShift)
}

// SetSs Sub second value SS[15:0] is the value of the synchronous prescaler counter when the timestamp event occurred.
func (r *RegisterRtctsssrType) SetSs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctsssrFieldSsMask)|(uint32(value)<<RegisterRtctsssrFieldSsShift))
}

// RegisterRtccalrType This register is write protected. The write access procedure is described in RTC register write protection on page9.
type RegisterRtccalrType uint32

func (r *RegisterRtccalrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtccalrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtccalrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtccalrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtccalrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtccalrFieldCalmShift = 0
	RegisterRtccalrFieldCalmMask  = 0x1ff
)

// GetCalm Calibration minus The frequency of the calendar is reduced by masking CALM out of 220 RTCCLK pulses (32 seconds if the input frequency is 32768 Hz). This decreases the frequency of the calendar with a resolution of 0.9537 ppm. To increase the frequency of the calendar, this feature should be used in conjunction with CALP. See Section24.3.12: RTC smooth digital calibration on page13.
func (r *RegisterRtccalrType) GetCalm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtccalrFieldCalmMask) >> RegisterRtccalrFieldCalmShift)
}

// SetCalm Calibration minus The frequency of the calendar is reduced by masking CALM out of 220 RTCCLK pulses (32 seconds if the input frequency is 32768 Hz). This decreases the frequency of the calendar with a resolution of 0.9537 ppm. To increase the frequency of the calendar, this feature should be used in conjunction with CALP. See Section24.3.12: RTC smooth digital calibration on page13.
func (r *RegisterRtccalrType) SetCalm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtccalrFieldCalmMask)|(uint32(value)<<RegisterRtccalrFieldCalmShift))
}

const (
	RegisterRtccalrFieldCalw16Shift = 13
	RegisterRtccalrFieldCalw16Mask  = 0x2000
)

// GetCalw16 Use a 16-second calibration cycle period When CALW16 is set to 1, the 16-second calibration cycle period is selected.This bit must not be set to 1 if CALW8=1. Note: CALM[0] is stuck at 0 when CALW16= 1. Refer to Section24.3.12: RTC smooth digital calibration.
func (r *RegisterRtccalrType) GetCalw16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccalrFieldCalw16Mask) != 0
}

// SetCalw16 Use a 16-second calibration cycle period When CALW16 is set to 1, the 16-second calibration cycle period is selected.This bit must not be set to 1 if CALW8=1. Note: CALM[0] is stuck at 0 when CALW16= 1. Refer to Section24.3.12: RTC smooth digital calibration.
func (r *RegisterRtccalrType) SetCalw16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccalrFieldCalw16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccalrFieldCalw16Mask)
	}
}

const (
	RegisterRtccalrFieldCalw8Shift = 14
	RegisterRtccalrFieldCalw8Mask  = 0x4000
)

// GetCalw8 Use an 8-second calibration cycle period When CALW8 is set to 1, the 8-second calibration cycle period is selected. Note: CALM[1:0] are stuck at 00; when CALW8= 1. Refer to Section24.3.12: RTC smooth digital calibration.
func (r *RegisterRtccalrType) GetCalw8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccalrFieldCalw8Mask) != 0
}

// SetCalw8 Use an 8-second calibration cycle period When CALW8 is set to 1, the 8-second calibration cycle period is selected. Note: CALM[1:0] are stuck at 00; when CALW8= 1. Refer to Section24.3.12: RTC smooth digital calibration.
func (r *RegisterRtccalrType) SetCalw8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccalrFieldCalw8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccalrFieldCalw8Mask)
	}
}

const (
	RegisterRtccalrFieldCalpShift = 15
	RegisterRtccalrFieldCalpMask  = 0x8000
)

// GetCalp Increase frequency of RTC by 488.5 ppm This feature is intended to be used in conjunction with CALM, which lowers the frequency of the calendar with a fine resolution. if the input frequency is 32768 Hz, the number of RTCCLK pulses added during a 32-second window is calculated as follows: (512 * CALP) - CALM. Refer to Section24.3.12: RTC smooth digital calibration.
func (r *RegisterRtccalrType) GetCalp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtccalrFieldCalpMask) != 0
}

// SetCalp Increase frequency of RTC by 488.5 ppm This feature is intended to be used in conjunction with CALM, which lowers the frequency of the calendar with a fine resolution. if the input frequency is 32768 Hz, the number of RTCCLK pulses added during a 32-second window is calculated as follows: (512 * CALP) - CALM. Refer to Section24.3.12: RTC smooth digital calibration.
func (r *RegisterRtccalrType) SetCalp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtccalrFieldCalpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtccalrFieldCalpMask)
	}
}

// RegisterRtctampcrType RTC tamper and alternate function configuration register
type RegisterRtctampcrType uint32

func (r *RegisterRtctampcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtctampcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtctampcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtctampcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtctampcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtctampcrFieldTamp1eShift = 0
	RegisterRtctampcrFieldTamp1eMask  = 0x1
)

// GetTamp1e RTC_TAMP1 input detection enable
func (r *RegisterRtctampcrType) GetTamp1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp1eMask) != 0
}

// SetTamp1e RTC_TAMP1 input detection enable
func (r *RegisterRtctampcrType) SetTamp1e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp1eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp1eMask)
	}
}

const (
	RegisterRtctampcrFieldTamp1trgShift = 1
	RegisterRtctampcrFieldTamp1trgMask  = 0x2
)

// GetTamp1trg Active level for RTC_TAMP1 input If TAMPFLT != 00 if TAMPFLT = 00:
func (r *RegisterRtctampcrType) GetTamp1trg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp1trgMask) != 0
}

// SetTamp1trg Active level for RTC_TAMP1 input If TAMPFLT != 00 if TAMPFLT = 00:
func (r *RegisterRtctampcrType) SetTamp1trg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp1trgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp1trgMask)
	}
}

const (
	RegisterRtctampcrFieldTampieShift = 2
	RegisterRtctampcrFieldTampieMask  = 0x4
)

// GetTampie Tamper interrupt enable
func (r *RegisterRtctampcrType) GetTampie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTampieMask) != 0
}

// SetTampie Tamper interrupt enable
func (r *RegisterRtctampcrType) SetTampie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTampieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTampieMask)
	}
}

const (
	RegisterRtctampcrFieldTamp2eShift = 3
	RegisterRtctampcrFieldTamp2eMask  = 0x8
)

// GetTamp2e RTC_TAMP2 input detection enable
func (r *RegisterRtctampcrType) GetTamp2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp2eMask) != 0
}

// SetTamp2e RTC_TAMP2 input detection enable
func (r *RegisterRtctampcrType) SetTamp2e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp2eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp2eMask)
	}
}

const (
	RegisterRtctampcrFieldTamp2trgShift = 4
	RegisterRtctampcrFieldTamp2trgMask  = 0x10
)

// GetTamp2trg Active level for RTC_TAMP2 input if TAMPFLT != 00: if TAMPFLT = 00:
func (r *RegisterRtctampcrType) GetTamp2trg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp2trgMask) != 0
}

// SetTamp2trg Active level for RTC_TAMP2 input if TAMPFLT != 00: if TAMPFLT = 00:
func (r *RegisterRtctampcrType) SetTamp2trg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp2trgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp2trgMask)
	}
}

const (
	RegisterRtctampcrFieldTamp3eShift = 5
	RegisterRtctampcrFieldTamp3eMask  = 0x20
)

// GetTamp3e RTC_TAMP3 detection enable
func (r *RegisterRtctampcrType) GetTamp3e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp3eMask) != 0
}

// SetTamp3e RTC_TAMP3 detection enable
func (r *RegisterRtctampcrType) SetTamp3e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp3eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp3eMask)
	}
}

const (
	RegisterRtctampcrFieldTamp3trgShift = 6
	RegisterRtctampcrFieldTamp3trgMask  = 0x40
)

// GetTamp3trg Active level for RTC_TAMP3 input if TAMPFLT != 00: if TAMPFLT = 00:
func (r *RegisterRtctampcrType) GetTamp3trg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp3trgMask) != 0
}

// SetTamp3trg Active level for RTC_TAMP3 input if TAMPFLT != 00: if TAMPFLT = 00:
func (r *RegisterRtctampcrType) SetTamp3trg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp3trgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp3trgMask)
	}
}

const (
	RegisterRtctampcrFieldTamptsShift = 7
	RegisterRtctampcrFieldTamptsMask  = 0x80
)

// GetTampts Activate timestamp on tamper detection event TAMPTS is valid even if TSE=0 in the RTC_CR register.
func (r *RegisterRtctampcrType) GetTampts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamptsMask) != 0
}

// SetTampts Activate timestamp on tamper detection event TAMPTS is valid even if TSE=0 in the RTC_CR register.
func (r *RegisterRtctampcrType) SetTampts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamptsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamptsMask)
	}
}

const (
	RegisterRtctampcrFieldTampfreqShift = 8
	RegisterRtctampcrFieldTampfreqMask  = 0x700
)

// GetTampfreq Tamper sampling frequency Determines the frequency at which each of the RTC_TAMPx inputs are sampled.
func (r *RegisterRtctampcrType) GetTampfreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTampfreqMask) >> RegisterRtctampcrFieldTampfreqShift)
}

// SetTampfreq Tamper sampling frequency Determines the frequency at which each of the RTC_TAMPx inputs are sampled.
func (r *RegisterRtctampcrType) SetTampfreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTampfreqMask)|(uint32(value)<<RegisterRtctampcrFieldTampfreqShift))
}

const (
	RegisterRtctampcrFieldTampfltShift = 11
	RegisterRtctampcrFieldTampfltMask  = 0x1800
)

// GetTampflt RTC_TAMPx filter count These bits determines the number of consecutive samples at the specified level (TAMP*TRG) needed to activate a Tamper event. TAMPFLT is valid for each of the RTC_TAMPx inputs.
func (r *RegisterRtctampcrType) GetTampflt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTampfltMask) >> RegisterRtctampcrFieldTampfltShift)
}

// SetTampflt RTC_TAMPx filter count These bits determines the number of consecutive samples at the specified level (TAMP*TRG) needed to activate a Tamper event. TAMPFLT is valid for each of the RTC_TAMPx inputs.
func (r *RegisterRtctampcrType) SetTampflt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTampfltMask)|(uint32(value)<<RegisterRtctampcrFieldTampfltShift))
}

const (
	RegisterRtctampcrFieldTampprchShift = 13
	RegisterRtctampcrFieldTampprchMask  = 0x6000
)

// GetTampprch RTC_TAMPx precharge duration These bit determines the duration of time during which the pull-up/is activated before each sample. TAMPPRCH is valid for each of the RTC_TAMPx inputs.
func (r *RegisterRtctampcrType) GetTampprch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTampprchMask) >> RegisterRtctampcrFieldTampprchShift)
}

// SetTampprch RTC_TAMPx precharge duration These bit determines the duration of time during which the pull-up/is activated before each sample. TAMPPRCH is valid for each of the RTC_TAMPx inputs.
func (r *RegisterRtctampcrType) SetTampprch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTampprchMask)|(uint32(value)<<RegisterRtctampcrFieldTampprchShift))
}

const (
	RegisterRtctampcrFieldTamppudisShift = 15
	RegisterRtctampcrFieldTamppudisMask  = 0x8000
)

// GetTamppudis RTC_TAMPx pull-up disable This bit determines if each of the RTC_TAMPx pins are pre-charged before each sample.
func (r *RegisterRtctampcrType) GetTamppudis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamppudisMask) != 0
}

// SetTamppudis RTC_TAMPx pull-up disable This bit determines if each of the RTC_TAMPx pins are pre-charged before each sample.
func (r *RegisterRtctampcrType) SetTamppudis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamppudisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamppudisMask)
	}
}

const (
	RegisterRtctampcrFieldTamp1ieShift = 16
	RegisterRtctampcrFieldTamp1ieMask  = 0x10000
)

// GetTamp1ie Tamper 1 interrupt enable
func (r *RegisterRtctampcrType) GetTamp1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp1ieMask) != 0
}

// SetTamp1ie Tamper 1 interrupt enable
func (r *RegisterRtctampcrType) SetTamp1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp1ieMask)
	}
}

const (
	RegisterRtctampcrFieldTamp1noeraseShift = 17
	RegisterRtctampcrFieldTamp1noeraseMask  = 0x20000
)

// GetTamp1noerase Tamper 1 no erase
func (r *RegisterRtctampcrType) GetTamp1noerase() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp1noeraseMask) != 0
}

// SetTamp1noerase Tamper 1 no erase
func (r *RegisterRtctampcrType) SetTamp1noerase(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp1noeraseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp1noeraseMask)
	}
}

const (
	RegisterRtctampcrFieldTamp1mfShift = 18
	RegisterRtctampcrFieldTamp1mfMask  = 0x40000
)

// GetTamp1mf Tamper 1 mask flag
func (r *RegisterRtctampcrType) GetTamp1mf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp1mfMask) != 0
}

// SetTamp1mf Tamper 1 mask flag
func (r *RegisterRtctampcrType) SetTamp1mf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp1mfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp1mfMask)
	}
}

const (
	RegisterRtctampcrFieldTamp2ieShift = 19
	RegisterRtctampcrFieldTamp2ieMask  = 0x80000
)

// GetTamp2ie Tamper 2 interrupt enable
func (r *RegisterRtctampcrType) GetTamp2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp2ieMask) != 0
}

// SetTamp2ie Tamper 2 interrupt enable
func (r *RegisterRtctampcrType) SetTamp2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp2ieMask)
	}
}

const (
	RegisterRtctampcrFieldTamp2noeraseShift = 20
	RegisterRtctampcrFieldTamp2noeraseMask  = 0x100000
)

// GetTamp2noerase Tamper 2 no erase
func (r *RegisterRtctampcrType) GetTamp2noerase() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp2noeraseMask) != 0
}

// SetTamp2noerase Tamper 2 no erase
func (r *RegisterRtctampcrType) SetTamp2noerase(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp2noeraseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp2noeraseMask)
	}
}

const (
	RegisterRtctampcrFieldTamp2mfShift = 21
	RegisterRtctampcrFieldTamp2mfMask  = 0x200000
)

// GetTamp2mf Tamper 2 mask flag
func (r *RegisterRtctampcrType) GetTamp2mf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp2mfMask) != 0
}

// SetTamp2mf Tamper 2 mask flag
func (r *RegisterRtctampcrType) SetTamp2mf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp2mfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp2mfMask)
	}
}

const (
	RegisterRtctampcrFieldTamp3ieShift = 22
	RegisterRtctampcrFieldTamp3ieMask  = 0x400000
)

// GetTamp3ie Tamper 3 interrupt enable
func (r *RegisterRtctampcrType) GetTamp3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp3ieMask) != 0
}

// SetTamp3ie Tamper 3 interrupt enable
func (r *RegisterRtctampcrType) SetTamp3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp3ieMask)
	}
}

const (
	RegisterRtctampcrFieldTamp3noeraseShift = 23
	RegisterRtctampcrFieldTamp3noeraseMask  = 0x800000
)

// GetTamp3noerase Tamper 3 no erase
func (r *RegisterRtctampcrType) GetTamp3noerase() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp3noeraseMask) != 0
}

// SetTamp3noerase Tamper 3 no erase
func (r *RegisterRtctampcrType) SetTamp3noerase(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp3noeraseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp3noeraseMask)
	}
}

const (
	RegisterRtctampcrFieldTamp3mfShift = 24
	RegisterRtctampcrFieldTamp3mfMask  = 0x1000000
)

// GetTamp3mf Tamper 3 mask flag
func (r *RegisterRtctampcrType) GetTamp3mf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtctampcrFieldTamp3mfMask) != 0
}

// SetTamp3mf Tamper 3 mask flag
func (r *RegisterRtctampcrType) SetTamp3mf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtctampcrFieldTamp3mfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtctampcrFieldTamp3mfMask)
	}
}

// RegisterRtcalrmassrType This register can be written only when ALRAE is reset in RTC_CR register, or in initialization mode.This register is write protected. The write access procedure is described in RTC register write protection on page9
type RegisterRtcalrmassrType uint32

func (r *RegisterRtcalrmassrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcalrmassrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcalrmassrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcalrmassrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcalrmassrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcalrmassrFieldSsShift = 0
	RegisterRtcalrmassrFieldSsMask  = 0x7fff
)

// GetSs Sub seconds value This value is compared with the contents of the synchronous prescaler counter to determine if Alarm A is to be activated. Only bits 0 up MASKSS-1 are compared.
func (r *RegisterRtcalrmassrType) GetSs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmassrFieldSsMask) >> RegisterRtcalrmassrFieldSsShift)
}

// SetSs Sub seconds value This value is compared with the contents of the synchronous prescaler counter to determine if Alarm A is to be activated. Only bits 0 up MASKSS-1 are compared.
func (r *RegisterRtcalrmassrType) SetSs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmassrFieldSsMask)|(uint32(value)<<RegisterRtcalrmassrFieldSsShift))
}

const (
	RegisterRtcalrmassrFieldMaskssShift = 24
	RegisterRtcalrmassrFieldMaskssMask  = 0xf000000
)

// GetMaskss Mask the most-significant bits starting at this bit ... The overflow bits of the synchronous counter (bits 15) is never compared. This bit can be different from 0 only after a shift operation.
func (r *RegisterRtcalrmassrType) GetMaskss() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmassrFieldMaskssMask) >> RegisterRtcalrmassrFieldMaskssShift)
}

// SetMaskss Mask the most-significant bits starting at this bit ... The overflow bits of the synchronous counter (bits 15) is never compared. This bit can be different from 0 only after a shift operation.
func (r *RegisterRtcalrmassrType) SetMaskss(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmassrFieldMaskssMask)|(uint32(value)<<RegisterRtcalrmassrFieldMaskssShift))
}

// RegisterRtcalrmbssrType This register can be written only when ALRBE is reset in RTC_CR register, or in initialization mode.This register is write protected.The write access procedure is described in Section: RTC register write protection.
type RegisterRtcalrmbssrType uint32

func (r *RegisterRtcalrmbssrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcalrmbssrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcalrmbssrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcalrmbssrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcalrmbssrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcalrmbssrFieldSsShift = 0
	RegisterRtcalrmbssrFieldSsMask  = 0x7fff
)

// GetSs Sub seconds value This value is compared with the contents of the synchronous prescaler counter to determine if Alarm B is to be activated. Only bits 0 up to MASKSS-1 are compared.
func (r *RegisterRtcalrmbssrType) GetSs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbssrFieldSsMask) >> RegisterRtcalrmbssrFieldSsShift)
}

// SetSs Sub seconds value This value is compared with the contents of the synchronous prescaler counter to determine if Alarm B is to be activated. Only bits 0 up to MASKSS-1 are compared.
func (r *RegisterRtcalrmbssrType) SetSs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbssrFieldSsMask)|(uint32(value)<<RegisterRtcalrmbssrFieldSsShift))
}

const (
	RegisterRtcalrmbssrFieldMaskssShift = 24
	RegisterRtcalrmbssrFieldMaskssMask  = 0xf000000
)

// GetMaskss Mask the most-significant bits starting at this bit ... The overflow bits of the synchronous counter (bits 15) is never compared. This bit can be different from 0 only after a shift operation.
func (r *RegisterRtcalrmbssrType) GetMaskss() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtcalrmbssrFieldMaskssMask) >> RegisterRtcalrmbssrFieldMaskssShift)
}

// SetMaskss Mask the most-significant bits starting at this bit ... The overflow bits of the synchronous counter (bits 15) is never compared. This bit can be different from 0 only after a shift operation.
func (r *RegisterRtcalrmbssrType) SetMaskss(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcalrmbssrFieldMaskssMask)|(uint32(value)<<RegisterRtcalrmbssrFieldMaskssShift))
}

// RegisterRtcorType RTC option register
type RegisterRtcorType uint32

func (r *RegisterRtcorType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcorType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcorType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcorType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcorType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcorFieldRtcalarmtypeShift = 0
	RegisterRtcorFieldRtcalarmtypeMask  = 0x1
)

// GetRtcalarmtype RTC_ALARM output type on PC13
func (r *RegisterRtcorType) GetRtcalarmtype() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcorFieldRtcalarmtypeMask) != 0
}

// SetRtcalarmtype RTC_ALARM output type on PC13
func (r *RegisterRtcorType) SetRtcalarmtype(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcorFieldRtcalarmtypeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcorFieldRtcalarmtypeMask)
	}
}

const (
	RegisterRtcorFieldRtcoutrmpShift = 1
	RegisterRtcorFieldRtcoutrmpMask  = 0x2
)

// GetRtcoutrmp RTC_OUT remap
func (r *RegisterRtcorType) GetRtcoutrmp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRtcorFieldRtcoutrmpMask) != 0
}

// SetRtcoutrmp RTC_OUT remap
func (r *RegisterRtcorType) SetRtcoutrmp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRtcorFieldRtcoutrmpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRtcorFieldRtcoutrmpMask)
	}
}

// RegisterRtcbkp0rType RTC backup registers
type RegisterRtcbkp0rType uint32

func (r *RegisterRtcbkp0rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp0rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp0rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp0rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp0rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp0rFieldBkpShift = 0
	RegisterRtcbkp0rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp0rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp0rFieldBkpMask) >> RegisterRtcbkp0rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp0rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp0rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp0rFieldBkpShift))
}

// RegisterRtcbkp1rType RTC backup registers
type RegisterRtcbkp1rType uint32

func (r *RegisterRtcbkp1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp1rFieldBkpShift = 0
	RegisterRtcbkp1rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp1rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp1rFieldBkpMask) >> RegisterRtcbkp1rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp1rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp1rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp1rFieldBkpShift))
}

// RegisterRtcbkp2rType RTC backup registers
type RegisterRtcbkp2rType uint32

func (r *RegisterRtcbkp2rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp2rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp2rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp2rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp2rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp2rFieldBkpShift = 0
	RegisterRtcbkp2rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp2rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp2rFieldBkpMask) >> RegisterRtcbkp2rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp2rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp2rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp2rFieldBkpShift))
}

// RegisterRtcbkp3rType RTC backup registers
type RegisterRtcbkp3rType uint32

func (r *RegisterRtcbkp3rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp3rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp3rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp3rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp3rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp3rFieldBkpShift = 0
	RegisterRtcbkp3rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp3rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp3rFieldBkpMask) >> RegisterRtcbkp3rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp3rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp3rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp3rFieldBkpShift))
}

// RegisterRtcbkp4rType RTC backup registers
type RegisterRtcbkp4rType uint32

func (r *RegisterRtcbkp4rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp4rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp4rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp4rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp4rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp4rFieldBkpShift = 0
	RegisterRtcbkp4rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp4rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp4rFieldBkpMask) >> RegisterRtcbkp4rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp4rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp4rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp4rFieldBkpShift))
}

// RegisterRtcbkp5rType RTC backup registers
type RegisterRtcbkp5rType uint32

func (r *RegisterRtcbkp5rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp5rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp5rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp5rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp5rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp5rFieldBkpShift = 0
	RegisterRtcbkp5rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp5rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp5rFieldBkpMask) >> RegisterRtcbkp5rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp5rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp5rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp5rFieldBkpShift))
}

// RegisterRtcbkp6rType RTC backup registers
type RegisterRtcbkp6rType uint32

func (r *RegisterRtcbkp6rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp6rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp6rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp6rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp6rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp6rFieldBkpShift = 0
	RegisterRtcbkp6rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp6rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp6rFieldBkpMask) >> RegisterRtcbkp6rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp6rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp6rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp6rFieldBkpShift))
}

// RegisterRtcbkp7rType RTC backup registers
type RegisterRtcbkp7rType uint32

func (r *RegisterRtcbkp7rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp7rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp7rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp7rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp7rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp7rFieldBkpShift = 0
	RegisterRtcbkp7rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp7rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp7rFieldBkpMask) >> RegisterRtcbkp7rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp7rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp7rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp7rFieldBkpShift))
}

// RegisterRtcbkp8rType RTC backup registers
type RegisterRtcbkp8rType uint32

func (r *RegisterRtcbkp8rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp8rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp8rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp8rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp8rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp8rFieldBkpShift = 0
	RegisterRtcbkp8rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp8rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp8rFieldBkpMask) >> RegisterRtcbkp8rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp8rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp8rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp8rFieldBkpShift))
}

// RegisterRtcbkp9rType RTC backup registers
type RegisterRtcbkp9rType uint32

func (r *RegisterRtcbkp9rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp9rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp9rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp9rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp9rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp9rFieldBkpShift = 0
	RegisterRtcbkp9rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp9rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp9rFieldBkpMask) >> RegisterRtcbkp9rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp9rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp9rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp9rFieldBkpShift))
}

// RegisterRtcbkp10rType RTC backup registers
type RegisterRtcbkp10rType uint32

func (r *RegisterRtcbkp10rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp10rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp10rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp10rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp10rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp10rFieldBkpShift = 0
	RegisterRtcbkp10rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp10rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp10rFieldBkpMask) >> RegisterRtcbkp10rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp10rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp10rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp10rFieldBkpShift))
}

// RegisterRtcbkp11rType RTC backup registers
type RegisterRtcbkp11rType uint32

func (r *RegisterRtcbkp11rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp11rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp11rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp11rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp11rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp11rFieldBkpShift = 0
	RegisterRtcbkp11rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp11rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp11rFieldBkpMask) >> RegisterRtcbkp11rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp11rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp11rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp11rFieldBkpShift))
}

// RegisterRtcbkp12rType RTC backup registers
type RegisterRtcbkp12rType uint32

func (r *RegisterRtcbkp12rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp12rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp12rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp12rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp12rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp12rFieldBkpShift = 0
	RegisterRtcbkp12rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp12rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp12rFieldBkpMask) >> RegisterRtcbkp12rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp12rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp12rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp12rFieldBkpShift))
}

// RegisterRtcbkp13rType RTC backup registers
type RegisterRtcbkp13rType uint32

func (r *RegisterRtcbkp13rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp13rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp13rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp13rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp13rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp13rFieldBkpShift = 0
	RegisterRtcbkp13rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp13rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp13rFieldBkpMask) >> RegisterRtcbkp13rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp13rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp13rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp13rFieldBkpShift))
}

// RegisterRtcbkp14rType RTC backup registers
type RegisterRtcbkp14rType uint32

func (r *RegisterRtcbkp14rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp14rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp14rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp14rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp14rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp14rFieldBkpShift = 0
	RegisterRtcbkp14rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp14rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp14rFieldBkpMask) >> RegisterRtcbkp14rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp14rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp14rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp14rFieldBkpShift))
}

// RegisterRtcbkp15rType RTC backup registers
type RegisterRtcbkp15rType uint32

func (r *RegisterRtcbkp15rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp15rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp15rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp15rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp15rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp15rFieldBkpShift = 0
	RegisterRtcbkp15rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp15rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp15rFieldBkpMask) >> RegisterRtcbkp15rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp15rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp15rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp15rFieldBkpShift))
}

// RegisterRtcbkp16rType RTC backup registers
type RegisterRtcbkp16rType uint32

func (r *RegisterRtcbkp16rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp16rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp16rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp16rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp16rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp16rFieldBkpShift = 0
	RegisterRtcbkp16rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp16rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp16rFieldBkpMask) >> RegisterRtcbkp16rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp16rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp16rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp16rFieldBkpShift))
}

// RegisterRtcbkp17rType RTC backup registers
type RegisterRtcbkp17rType uint32

func (r *RegisterRtcbkp17rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp17rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp17rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp17rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp17rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp17rFieldBkpShift = 0
	RegisterRtcbkp17rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp17rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp17rFieldBkpMask) >> RegisterRtcbkp17rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp17rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp17rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp17rFieldBkpShift))
}

// RegisterRtcbkp18rType RTC backup registers
type RegisterRtcbkp18rType uint32

func (r *RegisterRtcbkp18rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp18rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp18rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp18rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp18rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp18rFieldBkpShift = 0
	RegisterRtcbkp18rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp18rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp18rFieldBkpMask) >> RegisterRtcbkp18rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp18rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp18rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp18rFieldBkpShift))
}

// RegisterRtcbkp19rType RTC backup registers
type RegisterRtcbkp19rType uint32

func (r *RegisterRtcbkp19rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp19rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp19rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp19rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp19rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp19rFieldBkpShift = 0
	RegisterRtcbkp19rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp19rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp19rFieldBkpMask) >> RegisterRtcbkp19rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp19rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp19rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp19rFieldBkpShift))
}

// RegisterRtcbkp20rType RTC backup registers
type RegisterRtcbkp20rType uint32

func (r *RegisterRtcbkp20rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp20rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp20rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp20rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp20rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp20rFieldBkpShift = 0
	RegisterRtcbkp20rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp20rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp20rFieldBkpMask) >> RegisterRtcbkp20rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp20rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp20rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp20rFieldBkpShift))
}

// RegisterRtcbkp21rType RTC backup registers
type RegisterRtcbkp21rType uint32

func (r *RegisterRtcbkp21rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp21rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp21rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp21rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp21rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp21rFieldBkpShift = 0
	RegisterRtcbkp21rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp21rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp21rFieldBkpMask) >> RegisterRtcbkp21rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp21rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp21rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp21rFieldBkpShift))
}

// RegisterRtcbkp22rType RTC backup registers
type RegisterRtcbkp22rType uint32

func (r *RegisterRtcbkp22rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp22rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp22rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp22rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp22rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp22rFieldBkpShift = 0
	RegisterRtcbkp22rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp22rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp22rFieldBkpMask) >> RegisterRtcbkp22rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp22rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp22rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp22rFieldBkpShift))
}

// RegisterRtcbkp23rType RTC backup registers
type RegisterRtcbkp23rType uint32

func (r *RegisterRtcbkp23rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp23rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp23rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp23rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp23rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp23rFieldBkpShift = 0
	RegisterRtcbkp23rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp23rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp23rFieldBkpMask) >> RegisterRtcbkp23rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp23rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp23rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp23rFieldBkpShift))
}

// RegisterRtcbkp24rType RTC backup registers
type RegisterRtcbkp24rType uint32

func (r *RegisterRtcbkp24rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp24rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp24rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp24rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp24rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp24rFieldBkpShift = 0
	RegisterRtcbkp24rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp24rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp24rFieldBkpMask) >> RegisterRtcbkp24rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp24rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp24rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp24rFieldBkpShift))
}

// RegisterRtcbkp25rType RTC backup registers
type RegisterRtcbkp25rType uint32

func (r *RegisterRtcbkp25rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp25rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp25rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp25rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp25rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp25rFieldBkpShift = 0
	RegisterRtcbkp25rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp25rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp25rFieldBkpMask) >> RegisterRtcbkp25rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp25rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp25rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp25rFieldBkpShift))
}

// RegisterRtcbkp26rType RTC backup registers
type RegisterRtcbkp26rType uint32

func (r *RegisterRtcbkp26rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp26rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp26rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp26rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp26rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp26rFieldBkpShift = 0
	RegisterRtcbkp26rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp26rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp26rFieldBkpMask) >> RegisterRtcbkp26rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp26rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp26rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp26rFieldBkpShift))
}

// RegisterRtcbkp27rType RTC backup registers
type RegisterRtcbkp27rType uint32

func (r *RegisterRtcbkp27rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp27rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp27rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp27rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp27rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp27rFieldBkpShift = 0
	RegisterRtcbkp27rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp27rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp27rFieldBkpMask) >> RegisterRtcbkp27rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp27rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp27rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp27rFieldBkpShift))
}

// RegisterRtcbkp28rType RTC backup registers
type RegisterRtcbkp28rType uint32

func (r *RegisterRtcbkp28rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp28rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp28rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp28rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp28rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp28rFieldBkpShift = 0
	RegisterRtcbkp28rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp28rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp28rFieldBkpMask) >> RegisterRtcbkp28rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp28rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp28rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp28rFieldBkpShift))
}

// RegisterRtcbkp29rType RTC backup registers
type RegisterRtcbkp29rType uint32

func (r *RegisterRtcbkp29rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp29rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp29rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp29rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp29rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp29rFieldBkpShift = 0
	RegisterRtcbkp29rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp29rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp29rFieldBkpMask) >> RegisterRtcbkp29rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp29rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp29rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp29rFieldBkpShift))
}

// RegisterRtcbkp30rType RTC backup registers
type RegisterRtcbkp30rType uint32

func (r *RegisterRtcbkp30rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp30rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp30rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp30rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp30rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp30rFieldBkpShift = 0
	RegisterRtcbkp30rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp30rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp30rFieldBkpMask) >> RegisterRtcbkp30rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp30rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp30rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp30rFieldBkpShift))
}

// RegisterRtcbkp31rType RTC backup registers
type RegisterRtcbkp31rType uint32

func (r *RegisterRtcbkp31rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtcbkp31rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtcbkp31rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtcbkp31rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtcbkp31rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtcbkp31rFieldBkpShift = 0
	RegisterRtcbkp31rFieldBkpMask  = 0xffffffff
)

// GetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp31rType) GetBkp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtcbkp31rFieldBkpMask) >> RegisterRtcbkp31rFieldBkpShift)
}

// SetBkp The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.
func (r *RegisterRtcbkp31rType) SetBkp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtcbkp31rFieldBkpMask)|(uint32(value)<<RegisterRtcbkp31rFieldBkpShift))
}
