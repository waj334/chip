//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package flash

import (
	"unsafe"
	"volatile"
)

var (
	Flash = (*_flash)(unsafe.Pointer(uintptr(0x52002000)))
)

type _flash struct {
	Bank [2]_regsFlash
}

type _regsFlash struct {
	Acr       registerAcrType
	Keyr      registerKeyrType
	Optkeyr   registerOptkeyrType
	Cr        registerCrType
	Sr        registerSrType
	Ccr       registerCcrType
	Optcr     registerOptcrType
	Optsrcur  registerOptsrcurType
	Optsrprg  registerOptsrprgType
	Optccr    registerOptccrType
	Prarcur   registerPrarcurType
	Prarprg2  registerPrarprg2Type
	Prarprg   registerPrarprgType
	Scarcur   registerScarcurType
	Scarprg   registerScarprgType
	Wpsncur1r registerWpsncur1rType
	Wpsnprg1r registerWpsnprg1rType
	Bootcurr  registerBootcurrType
	Bootprgr  registerBootprgrType
	_         [8]byte
	Crccr     registerCrccrType
	Crcsadd1r registerCrcsadd1rType
	Crceadd1r registerCrceadd1rType
	Crcdatar  registerCrcdatarType
	Eccfa1r   registerEccfa1rType
}

// registerAcrType Access control register
type registerAcrType uint32

const (
	RegisterAcrFieldLatencyShift = 0
	RegisterAcrFieldLatencyMask  = 0x7
)

// GetLatency Read latency
func (r *registerAcrType) GetLatency() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAcrFieldLatencyMask) >> RegisterAcrFieldLatencyShift)
}

// SetLatency Read latency
func (r *registerAcrType) SetLatency(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAcrFieldLatencyMask)|(uint32(value)<<RegisterAcrFieldLatencyShift))
}

const (
	RegisterAcrFieldWrhighfreqShift = 4
	RegisterAcrFieldWrhighfreqMask  = 0x30
)

// GetWrhighfreq Flash signal delay
func (r *registerAcrType) GetWrhighfreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAcrFieldWrhighfreqMask) >> RegisterAcrFieldWrhighfreqShift)
}

// SetWrhighfreq Flash signal delay
func (r *registerAcrType) SetWrhighfreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAcrFieldWrhighfreqMask)|(uint32(value)<<RegisterAcrFieldWrhighfreqShift))
}

// registerKeyrType FLASH key register for bank n
type registerKeyrType uint32

const (
	RegisterKeyrFieldKeyrShift = 0
	RegisterKeyrFieldKeyrMask  = 0xffffffff
)

// GetKeyr Bank n access configuration unlock key
func (r *registerKeyrType) GetKeyr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterKeyrFieldKeyrMask) >> RegisterKeyrFieldKeyrShift)
}

// SetKeyr Bank n access configuration unlock key
func (r *registerKeyrType) SetKeyr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterKeyrFieldKeyrMask)|(uint32(value)<<RegisterKeyrFieldKeyrShift))
}

// registerOptkeyrType FLASH option key register
type registerOptkeyrType uint32

const (
	RegisterOptkeyrFieldOptkeyrShift = 0
	RegisterOptkeyrFieldOptkeyrMask  = 0xffffffff
)

// GetOptkeyr Unlock key option bytes
func (r *registerOptkeyrType) GetOptkeyr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOptkeyrFieldOptkeyrMask) >> RegisterOptkeyrFieldOptkeyrShift)
}

// SetOptkeyr Unlock key option bytes
func (r *registerOptkeyrType) SetOptkeyr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptkeyrFieldOptkeyrMask)|(uint32(value)<<RegisterOptkeyrFieldOptkeyrShift))
}

// registerCrType FLASH control register for bank n
type registerCrType uint32

const (
	RegisterCrFieldLockShift = 0
	RegisterCrFieldLockMask  = 0x1
)

// GetLock Bank n configuration lock bit
func (r *registerCrType) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLockMask) != 0
}

// SetLock Bank n configuration lock bit
func (r *registerCrType) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldLockMask)
	}
}

const (
	RegisterCrFieldPgShift = 1
	RegisterCrFieldPgMask  = 0x2
)

// GetPg Bank n program enable bit
func (r *registerCrType) GetPg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPgMask) != 0
}

// SetPg Bank n program enable bit
func (r *registerCrType) SetPg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldPgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPgMask)
	}
}

const (
	RegisterCrFieldSerShift = 2
	RegisterCrFieldSerMask  = 0x4
)

// GetSer Bank n sector erase request
func (r *registerCrType) GetSer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSerMask) != 0
}

// SetSer Bank n sector erase request
func (r *registerCrType) SetSer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldSerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSerMask)
	}
}

const (
	RegisterCrFieldBerShift = 3
	RegisterCrFieldBerMask  = 0x8
)

// GetBer Bank n erase request
func (r *registerCrType) GetBer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldBerMask) != 0
}

// SetBer Bank n erase request
func (r *registerCrType) SetBer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldBerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldBerMask)
	}
}

const (
	RegisterCrFieldPsizeShift = 4
	RegisterCrFieldPsizeMask  = 0x30
)

// GetPsize Bank n program size
func (r *registerCrType) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPsizeMask) >> RegisterCrFieldPsizeShift)
}

// SetPsize Bank n program size
func (r *registerCrType) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPsizeMask)|(uint32(value)<<RegisterCrFieldPsizeShift))
}

const (
	RegisterCrFieldFwShift = 6
	RegisterCrFieldFwMask  = 0x40
)

// GetFw Bank n write forcing control bit
func (r *registerCrType) GetFw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldFwMask) != 0
}

// SetFw Bank n write forcing control bit
func (r *registerCrType) SetFw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldFwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldFwMask)
	}
}

const (
	RegisterCrFieldStartShift = 7
	RegisterCrFieldStartMask  = 0x80
)

// GetStart Bank n bank or sector erase start control bit
func (r *registerCrType) GetStart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldStartMask) != 0
}

// SetStart Bank n bank or sector erase start control bit
func (r *registerCrType) SetStart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldStartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldStartMask)
	}
}

const (
	RegisterCrFieldSnbShift = 8
	RegisterCrFieldSnbMask  = 0x700
)

// GetSnb Bank n sector erase selection number
func (r *registerCrType) GetSnb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSnbMask) >> RegisterCrFieldSnbShift)
}

// SetSnb Bank n sector erase selection number
func (r *registerCrType) SetSnb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSnbMask)|(uint32(value)<<RegisterCrFieldSnbShift))
}

const (
	RegisterCrFieldCrcenShift = 15
	RegisterCrFieldCrcenMask  = 0x8000
)

// GetCrcen Bank n CRC control bit
func (r *registerCrType) GetCrcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCrcenMask) != 0
}

// SetCrcen Bank n CRC control bit
func (r *registerCrType) SetCrcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCrcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCrcenMask)
	}
}

const (
	RegisterCrFieldEopieShift = 16
	RegisterCrFieldEopieMask  = 0x10000
)

// GetEopie Bank n end-of-program interrupt control bit
func (r *registerCrType) GetEopie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEopieMask) != 0
}

// SetEopie Bank n end-of-program interrupt control bit
func (r *registerCrType) SetEopie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEopieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEopieMask)
	}
}

const (
	RegisterCrFieldWrperrieShift = 17
	RegisterCrFieldWrperrieMask  = 0x20000
)

// GetWrperrie Bank n write protection error interrupt enable bit
func (r *registerCrType) GetWrperrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldWrperrieMask) != 0
}

// SetWrperrie Bank n write protection error interrupt enable bit
func (r *registerCrType) SetWrperrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldWrperrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldWrperrieMask)
	}
}

const (
	RegisterCrFieldPgserrieShift = 18
	RegisterCrFieldPgserrieMask  = 0x40000
)

// GetPgserrie Bank n programming sequence error interrupt enable bit
func (r *registerCrType) GetPgserrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPgserrieMask) != 0
}

// SetPgserrie Bank n programming sequence error interrupt enable bit
func (r *registerCrType) SetPgserrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldPgserrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPgserrieMask)
	}
}

const (
	RegisterCrFieldStrberrieShift = 19
	RegisterCrFieldStrberrieMask  = 0x80000
)

// GetStrberrie Bank n strobe error interrupt enable bit
func (r *registerCrType) GetStrberrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldStrberrieMask) != 0
}

// SetStrberrie Bank n strobe error interrupt enable bit
func (r *registerCrType) SetStrberrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldStrberrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldStrberrieMask)
	}
}

const (
	RegisterCrFieldIncerrieShift = 21
	RegisterCrFieldIncerrieMask  = 0x200000
)

// GetIncerrie Bank n inconsistency error interrupt enable bit
func (r *registerCrType) GetIncerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldIncerrieMask) != 0
}

// SetIncerrie Bank n inconsistency error interrupt enable bit
func (r *registerCrType) SetIncerrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldIncerrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldIncerrieMask)
	}
}

const (
	RegisterCrFieldOperrieShift = 22
	RegisterCrFieldOperrieMask  = 0x400000
)

// GetOperrie Bank n write/erase error interrupt enable bit
func (r *registerCrType) GetOperrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldOperrieMask) != 0
}

// SetOperrie Bank n write/erase error interrupt enable bit
func (r *registerCrType) SetOperrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldOperrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldOperrieMask)
	}
}

const (
	RegisterCrFieldRdperrieShift = 23
	RegisterCrFieldRdperrieMask  = 0x800000
)

// GetRdperrie Bank n read protection error interrupt enable bit
func (r *registerCrType) GetRdperrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRdperrieMask) != 0
}

// SetRdperrie Bank n read protection error interrupt enable bit
func (r *registerCrType) SetRdperrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRdperrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRdperrieMask)
	}
}

const (
	RegisterCrFieldRdserrieShift = 24
	RegisterCrFieldRdserrieMask  = 0x1000000
)

// GetRdserrie Bank n secure error interrupt enable bit
func (r *registerCrType) GetRdserrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRdserrieMask) != 0
}

// SetRdserrie Bank n secure error interrupt enable bit
func (r *registerCrType) SetRdserrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRdserrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRdserrieMask)
	}
}

const (
	RegisterCrFieldSneccerrieShift = 25
	RegisterCrFieldSneccerrieMask  = 0x2000000
)

// GetSneccerrie Bank n ECC single correction error interrupt enable bit
func (r *registerCrType) GetSneccerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSneccerrieMask) != 0
}

// SetSneccerrie Bank n ECC single correction error interrupt enable bit
func (r *registerCrType) SetSneccerrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldSneccerrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSneccerrieMask)
	}
}

const (
	RegisterCrFieldDbeccerrieShift = 26
	RegisterCrFieldDbeccerrieMask  = 0x4000000
)

// GetDbeccerrie Bank n ECC double detection error interrupt enable bit
func (r *registerCrType) GetDbeccerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDbeccerrieMask) != 0
}

// SetDbeccerrie Bank n ECC double detection error interrupt enable bit
func (r *registerCrType) SetDbeccerrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDbeccerrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDbeccerrieMask)
	}
}

const (
	RegisterCrFieldCrcendieShift = 27
	RegisterCrFieldCrcendieMask  = 0x8000000
)

// GetCrcendie Bank n end of CRC calculation interrupt enable bit
func (r *registerCrType) GetCrcendie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCrcendieMask) != 0
}

// SetCrcendie Bank n end of CRC calculation interrupt enable bit
func (r *registerCrType) SetCrcendie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCrcendieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCrcendieMask)
	}
}

// registerSrType FLASH status register for bank n
type registerSrType uint32

const (
	RegisterSrFieldBsyShift = 0
	RegisterSrFieldBsyMask  = 0x1
)

// GetBsy Bank n ongoing program flag
func (r *registerSrType) GetBsy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldBsyMask) != 0
}

// SetBsy Bank n ongoing program flag
func (r *registerSrType) SetBsy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldBsyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldBsyMask)
	}
}

const (
	RegisterSrFieldWbneShift = 1
	RegisterSrFieldWbneMask  = 0x2
)

// GetWbne Bank n write buffer not empty flag
func (r *registerSrType) GetWbne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldWbneMask) != 0
}

// SetWbne Bank n write buffer not empty flag
func (r *registerSrType) SetWbne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldWbneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldWbneMask)
	}
}

const (
	RegisterSrFieldQwShift = 2
	RegisterSrFieldQwMask  = 0x4
)

// GetQw Bank n wait queue flag
func (r *registerSrType) GetQw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldQwMask) != 0
}

// SetQw Bank n wait queue flag
func (r *registerSrType) SetQw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldQwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldQwMask)
	}
}

const (
	RegisterSrFieldCrcbusyShift = 3
	RegisterSrFieldCrcbusyMask  = 0x8
)

// GetCrcbusy Bank n CRC busy flag
func (r *registerSrType) GetCrcbusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCrcbusyMask) != 0
}

// SetCrcbusy Bank n CRC busy flag
func (r *registerSrType) SetCrcbusy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCrcbusyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCrcbusyMask)
	}
}

const (
	RegisterSrFieldEopShift = 16
	RegisterSrFieldEopMask  = 0x10000
)

// GetEop Bank n end-of-program flag
func (r *registerSrType) GetEop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldEopMask) != 0
}

// SetEop Bank n end-of-program flag
func (r *registerSrType) SetEop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldEopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldEopMask)
	}
}

const (
	RegisterSrFieldWrperrShift = 17
	RegisterSrFieldWrperrMask  = 0x20000
)

// GetWrperr Bank n write protection error flag
func (r *registerSrType) GetWrperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldWrperrMask) != 0
}

// SetWrperr Bank n write protection error flag
func (r *registerSrType) SetWrperr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldWrperrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldWrperrMask)
	}
}

const (
	RegisterSrFieldPgserrShift = 18
	RegisterSrFieldPgserrMask  = 0x40000
)

// GetPgserr Bank n programming sequence error flag
func (r *registerSrType) GetPgserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldPgserrMask) != 0
}

// SetPgserr Bank n programming sequence error flag
func (r *registerSrType) SetPgserr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldPgserrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldPgserrMask)
	}
}

const (
	RegisterSrFieldStrberrShift = 19
	RegisterSrFieldStrberrMask  = 0x80000
)

// GetStrberr Bank n strobe error flag
func (r *registerSrType) GetStrberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldStrberrMask) != 0
}

// SetStrberr Bank n strobe error flag
func (r *registerSrType) SetStrberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldStrberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldStrberrMask)
	}
}

const (
	RegisterSrFieldIncerrShift = 21
	RegisterSrFieldIncerrMask  = 0x200000
)

// GetIncerr Bank n inconsistency error flag
func (r *registerSrType) GetIncerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldIncerrMask) != 0
}

// SetIncerr Bank n inconsistency error flag
func (r *registerSrType) SetIncerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldIncerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldIncerrMask)
	}
}

const (
	RegisterSrFieldOperrShift = 22
	RegisterSrFieldOperrMask  = 0x400000
)

// GetOperr Bank n write/erase error flag
func (r *registerSrType) GetOperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldOperrMask) != 0
}

// SetOperr Bank n write/erase error flag
func (r *registerSrType) SetOperr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldOperrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldOperrMask)
	}
}

const (
	RegisterSrFieldRdperrShift = 23
	RegisterSrFieldRdperrMask  = 0x800000
)

// GetRdperr Bank n read protection error flag
func (r *registerSrType) GetRdperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRdperrMask) != 0
}

// SetRdperr Bank n read protection error flag
func (r *registerSrType) SetRdperr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldRdperrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldRdperrMask)
	}
}

const (
	RegisterSrFieldRdserrShift = 24
	RegisterSrFieldRdserrMask  = 0x1000000
)

// GetRdserr Bank n secure error flag
func (r *registerSrType) GetRdserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRdserrMask) != 0
}

// SetRdserr Bank n secure error flag
func (r *registerSrType) SetRdserr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldRdserrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldRdserrMask)
	}
}

const (
	RegisterSrFieldSneccerr1Shift = 25
	RegisterSrFieldSneccerr1Mask  = 0x2000000
)

// GetSneccerr1 Bank n single correction error flag
func (r *registerSrType) GetSneccerr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSneccerr1Mask) != 0
}

// SetSneccerr1 Bank n single correction error flag
func (r *registerSrType) SetSneccerr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldSneccerr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldSneccerr1Mask)
	}
}

const (
	RegisterSrFieldDbeccerrShift = 26
	RegisterSrFieldDbeccerrMask  = 0x4000000
)

// GetDbeccerr Bank n ECC double detection error flag
func (r *registerSrType) GetDbeccerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDbeccerrMask) != 0
}

// SetDbeccerr Bank n ECC double detection error flag
func (r *registerSrType) SetDbeccerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldDbeccerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldDbeccerrMask)
	}
}

const (
	RegisterSrFieldCrcendShift = 27
	RegisterSrFieldCrcendMask  = 0x8000000
)

// GetCrcend Bank n CRC-complete flag
func (r *registerSrType) GetCrcend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCrcendMask) != 0
}

// SetCrcend Bank n CRC-complete flag
func (r *registerSrType) SetCrcend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCrcendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCrcendMask)
	}
}

// registerCcrType FLASH clear control register for bank n
type registerCcrType uint32

const (
	RegisterCcrFieldClreopShift = 16
	RegisterCcrFieldClreopMask  = 0x10000
)

// GetClreop Bank n EOP1 flag clear bit
func (r *registerCcrType) GetClreop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClreopMask) != 0
}

// SetClreop Bank n EOP1 flag clear bit
func (r *registerCcrType) SetClreop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldClreopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldClreopMask)
	}
}

const (
	RegisterCcrFieldClrwrperrShift = 17
	RegisterCcrFieldClrwrperrMask  = 0x20000
)

// GetClrwrperr Bank n WRPERR1 flag clear bit
func (r *registerCcrType) GetClrwrperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrwrperrMask) != 0
}

// SetClrwrperr Bank n WRPERR1 flag clear bit
func (r *registerCcrType) SetClrwrperr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldClrwrperrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldClrwrperrMask)
	}
}

const (
	RegisterCcrFieldClrpgserrShift = 18
	RegisterCcrFieldClrpgserrMask  = 0x40000
)

// GetClrpgserr Bank n PGSERR1 flag clear bi
func (r *registerCcrType) GetClrpgserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrpgserrMask) != 0
}

// SetClrpgserr Bank n PGSERR1 flag clear bi
func (r *registerCcrType) SetClrpgserr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldClrpgserrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldClrpgserrMask)
	}
}

const (
	RegisterCcrFieldClrstrberrShift = 19
	RegisterCcrFieldClrstrberrMask  = 0x80000
)

// GetClrstrberr Bank n STRBERR1 flag clear bit
func (r *registerCcrType) GetClrstrberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrstrberrMask) != 0
}

// SetClrstrberr Bank n STRBERR1 flag clear bit
func (r *registerCcrType) SetClrstrberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldClrstrberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldClrstrberrMask)
	}
}

const (
	RegisterCcrFieldClrincerrShift = 21
	RegisterCcrFieldClrincerrMask  = 0x200000
)

// GetClrincerr Bank n INCERR1 flag clear bit
func (r *registerCcrType) GetClrincerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrincerrMask) != 0
}

// SetClrincerr Bank n INCERR1 flag clear bit
func (r *registerCcrType) SetClrincerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldClrincerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldClrincerrMask)
	}
}

const (
	RegisterCcrFieldClroperrShift = 22
	RegisterCcrFieldClroperrMask  = 0x400000
)

// GetClroperr Bank n OPERR1 flag clear bit
func (r *registerCcrType) GetClroperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClroperrMask) != 0
}

// SetClroperr Bank n OPERR1 flag clear bit
func (r *registerCcrType) SetClroperr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldClroperrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldClroperrMask)
	}
}

const (
	RegisterCcrFieldClrrdperrShift = 23
	RegisterCcrFieldClrrdperrMask  = 0x800000
)

// GetClrrdperr Bank n RDPERR1 flag clear bit
func (r *registerCcrType) GetClrrdperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrrdperrMask) != 0
}

// SetClrrdperr Bank n RDPERR1 flag clear bit
func (r *registerCcrType) SetClrrdperr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldClrrdperrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldClrrdperrMask)
	}
}

const (
	RegisterCcrFieldClrrdserrShift = 24
	RegisterCcrFieldClrrdserrMask  = 0x1000000
)

// GetClrrdserr Bank n RDSERR1 flag clear bit
func (r *registerCcrType) GetClrrdserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrrdserrMask) != 0
}

// SetClrrdserr Bank n RDSERR1 flag clear bit
func (r *registerCcrType) SetClrrdserr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldClrrdserrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldClrrdserrMask)
	}
}

const (
	RegisterCcrFieldClrsneccerrShift = 25
	RegisterCcrFieldClrsneccerrMask  = 0x2000000
)

// GetClrsneccerr Bank n SNECCERR1 flag clear bit
func (r *registerCcrType) GetClrsneccerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrsneccerrMask) != 0
}

// SetClrsneccerr Bank n SNECCERR1 flag clear bit
func (r *registerCcrType) SetClrsneccerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldClrsneccerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldClrsneccerrMask)
	}
}

const (
	RegisterCcrFieldClrdbeccerrShift = 26
	RegisterCcrFieldClrdbeccerrMask  = 0x4000000
)

// GetClrdbeccerr Bank n DBECCERR1 flag clear bit
func (r *registerCcrType) GetClrdbeccerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrdbeccerrMask) != 0
}

// SetClrdbeccerr Bank n DBECCERR1 flag clear bit
func (r *registerCcrType) SetClrdbeccerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldClrdbeccerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldClrdbeccerrMask)
	}
}

const (
	RegisterCcrFieldClrcrcendShift = 27
	RegisterCcrFieldClrcrcendMask  = 0x8000000
)

// GetClrcrcend Bank n CRCEND1 flag clear bit
func (r *registerCcrType) GetClrcrcend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrcrcendMask) != 0
}

// SetClrcrcend Bank n CRCEND1 flag clear bit
func (r *registerCcrType) SetClrcrcend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldClrcrcendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldClrcrcendMask)
	}
}

// registerOptcrType FLASH option control register
type registerOptcrType uint32

const (
	RegisterOptcrFieldOptlockShift = 0
	RegisterOptcrFieldOptlockMask  = 0x1
)

// GetOptlock FLASH_OPTCR lock option configuration bit
func (r *registerOptcrType) GetOptlock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptcrFieldOptlockMask) != 0
}

// SetOptlock FLASH_OPTCR lock option configuration bit
func (r *registerOptcrType) SetOptlock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptcrFieldOptlockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptcrFieldOptlockMask)
	}
}

const (
	RegisterOptcrFieldOptstartShift = 1
	RegisterOptcrFieldOptstartMask  = 0x2
)

// GetOptstart Option byte start change option configuration bit
func (r *registerOptcrType) GetOptstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptcrFieldOptstartMask) != 0
}

// SetOptstart Option byte start change option configuration bit
func (r *registerOptcrType) SetOptstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptcrFieldOptstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptcrFieldOptstartMask)
	}
}

const (
	RegisterOptcrFieldMerShift = 4
	RegisterOptcrFieldMerMask  = 0x10
)

// GetMer Flash mass erase enable bit
func (r *registerOptcrType) GetMer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptcrFieldMerMask) != 0
}

// SetMer Flash mass erase enable bit
func (r *registerOptcrType) SetMer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptcrFieldMerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptcrFieldMerMask)
	}
}

const (
	RegisterOptcrFieldOptchangeerrieShift = 30
	RegisterOptcrFieldOptchangeerrieMask  = 0x40000000
)

// GetOptchangeerrie Option byte change error interrupt enable bit
func (r *registerOptcrType) GetOptchangeerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptcrFieldOptchangeerrieMask) != 0
}

// SetOptchangeerrie Option byte change error interrupt enable bit
func (r *registerOptcrType) SetOptchangeerrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptcrFieldOptchangeerrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptcrFieldOptchangeerrieMask)
	}
}

const (
	RegisterOptcrFieldSwapbankShift = 31
	RegisterOptcrFieldSwapbankMask  = 0x80000000
)

// GetSwapbank Bank swapping configuration bit
func (r *registerOptcrType) GetSwapbank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptcrFieldSwapbankMask) != 0
}

// SetSwapbank Bank swapping configuration bit
func (r *registerOptcrType) SetSwapbank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptcrFieldSwapbankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptcrFieldSwapbankMask)
	}
}

// registerOptsrcurType FLASH option status register
type registerOptsrcurType uint32

const (
	RegisterOptsrcurFieldOptbusyShift = 0
	RegisterOptsrcurFieldOptbusyMask  = 0x1
)

// GetOptbusy Option byte change ongoing flag
func (r *registerOptsrcurType) GetOptbusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldOptbusyMask) != 0
}

// SetOptbusy Option byte change ongoing flag
func (r *registerOptsrcurType) SetOptbusy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldOptbusyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldOptbusyMask)
	}
}

const (
	RegisterOptsrcurFieldBorlevShift = 2
	RegisterOptsrcurFieldBorlevMask  = 0xc
)

// GetBorlev Brownout level option status bit
func (r *registerOptsrcurType) GetBorlev() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldBorlevMask) >> RegisterOptsrcurFieldBorlevShift)
}

// SetBorlev Brownout level option status bit
func (r *registerOptsrcurType) SetBorlev(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldBorlevMask)|(uint32(value)<<RegisterOptsrcurFieldBorlevShift))
}

const (
	RegisterOptsrcurFieldIwdg1hwShift = 4
	RegisterOptsrcurFieldIwdg1hwMask  = 0x10
)

// GetIwdg1hw IWDG1 control option status bit
func (r *registerOptsrcurType) GetIwdg1hw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldIwdg1hwMask) != 0
}

// SetIwdg1hw IWDG1 control option status bit
func (r *registerOptsrcurType) SetIwdg1hw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldIwdg1hwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldIwdg1hwMask)
	}
}

const (
	RegisterOptsrcurFieldNrststopdShift = 6
	RegisterOptsrcurFieldNrststopdMask  = 0x40
)

// GetNrststopd D1 DStop entry reset option status bit
func (r *registerOptsrcurType) GetNrststopd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldNrststopdMask) != 0
}

// SetNrststopd D1 DStop entry reset option status bit
func (r *registerOptsrcurType) SetNrststopd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldNrststopdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldNrststopdMask)
	}
}

const (
	RegisterOptsrcurFieldNrststbydShift = 7
	RegisterOptsrcurFieldNrststbydMask  = 0x80
)

// GetNrststbyd D1 DStandby entry reset option status bit
func (r *registerOptsrcurType) GetNrststbyd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldNrststbydMask) != 0
}

// SetNrststbyd D1 DStandby entry reset option status bit
func (r *registerOptsrcurType) SetNrststbyd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldNrststbydMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldNrststbydMask)
	}
}

const (
	RegisterOptsrcurFieldRdpShift = 8
	RegisterOptsrcurFieldRdpMask  = 0xff00
)

// GetRdp Readout protection level option status byte
func (r *registerOptsrcurType) GetRdp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldRdpMask) >> RegisterOptsrcurFieldRdpShift)
}

// SetRdp Readout protection level option status byte
func (r *registerOptsrcurType) SetRdp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldRdpMask)|(uint32(value)<<RegisterOptsrcurFieldRdpShift))
}

const (
	RegisterOptsrcurFieldFziwdgstopShift = 17
	RegisterOptsrcurFieldFziwdgstopMask  = 0x20000
)

// GetFziwdgstop IWDG Stop mode freeze option status bit
func (r *registerOptsrcurType) GetFziwdgstop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldFziwdgstopMask) != 0
}

// SetFziwdgstop IWDG Stop mode freeze option status bit
func (r *registerOptsrcurType) SetFziwdgstop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldFziwdgstopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldFziwdgstopMask)
	}
}

const (
	RegisterOptsrcurFieldFziwdgsdbyShift = 18
	RegisterOptsrcurFieldFziwdgsdbyMask  = 0x40000
)

// GetFziwdgsdby IWDG Standby mode freeze option status bit
func (r *registerOptsrcurType) GetFziwdgsdby() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldFziwdgsdbyMask) != 0
}

// SetFziwdgsdby IWDG Standby mode freeze option status bit
func (r *registerOptsrcurType) SetFziwdgsdby(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldFziwdgsdbyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldFziwdgsdbyMask)
	}
}

const (
	RegisterOptsrcurFieldStramsizeShift = 19
	RegisterOptsrcurFieldStramsizeMask  = 0x180000
)

// GetStramsize DTCM RAM size option status
func (r *registerOptsrcurType) GetStramsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldStramsizeMask) >> RegisterOptsrcurFieldStramsizeShift)
}

// SetStramsize DTCM RAM size option status
func (r *registerOptsrcurType) SetStramsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldStramsizeMask)|(uint32(value)<<RegisterOptsrcurFieldStramsizeShift))
}

const (
	RegisterOptsrcurFieldSecurityShift = 21
	RegisterOptsrcurFieldSecurityMask  = 0x200000
)

// GetSecurity Security enable option status bit
func (r *registerOptsrcurType) GetSecurity() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldSecurityMask) != 0
}

// SetSecurity Security enable option status bit
func (r *registerOptsrcurType) SetSecurity(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldSecurityMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldSecurityMask)
	}
}

const (
	RegisterOptsrcurFieldRss1Shift = 26
	RegisterOptsrcurFieldRss1Mask  = 0x4000000
)

// GetRss1 User option bit 1
func (r *registerOptsrcurType) GetRss1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldRss1Mask) != 0
}

// SetRss1 User option bit 1
func (r *registerOptsrcurType) SetRss1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldRss1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldRss1Mask)
	}
}

const (
	RegisterOptsrcurFieldPersookShift = 28
	RegisterOptsrcurFieldPersookMask  = 0x10000000
)

// GetPersook Device personalization status bit
func (r *registerOptsrcurType) GetPersook() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldPersookMask) != 0
}

// SetPersook Device personalization status bit
func (r *registerOptsrcurType) SetPersook(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldPersookMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldPersookMask)
	}
}

const (
	RegisterOptsrcurFieldIohslvShift = 29
	RegisterOptsrcurFieldIohslvMask  = 0x20000000
)

// GetIohslv I/O high-speed at low-voltage status bit (PRODUCT_BELOW_25V)
func (r *registerOptsrcurType) GetIohslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldIohslvMask) != 0
}

// SetIohslv I/O high-speed at low-voltage status bit (PRODUCT_BELOW_25V)
func (r *registerOptsrcurType) SetIohslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldIohslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldIohslvMask)
	}
}

const (
	RegisterOptsrcurFieldOptchangeerrShift = 30
	RegisterOptsrcurFieldOptchangeerrMask  = 0x40000000
)

// GetOptchangeerr Option byte change error flag
func (r *registerOptsrcurType) GetOptchangeerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldOptchangeerrMask) != 0
}

// SetOptchangeerr Option byte change error flag
func (r *registerOptsrcurType) SetOptchangeerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldOptchangeerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldOptchangeerrMask)
	}
}

const (
	RegisterOptsrcurFieldSwapbankoptShift = 31
	RegisterOptsrcurFieldSwapbankoptMask  = 0x80000000
)

// GetSwapbankopt Bank swapping option status bit
func (r *registerOptsrcurType) GetSwapbankopt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldSwapbankoptMask) != 0
}

// SetSwapbankopt Bank swapping option status bit
func (r *registerOptsrcurType) SetSwapbankopt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldSwapbankoptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldSwapbankoptMask)
	}
}

// registerOptsrprgType FLASH option status register
type registerOptsrprgType uint32

const (
	RegisterOptsrprgFieldBorlevShift = 2
	RegisterOptsrprgFieldBorlevMask  = 0xc
)

// GetBorlev BOR reset level option configuration bits
func (r *registerOptsrprgType) GetBorlev() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldBorlevMask) >> RegisterOptsrprgFieldBorlevShift)
}

// SetBorlev BOR reset level option configuration bits
func (r *registerOptsrprgType) SetBorlev(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldBorlevMask)|(uint32(value)<<RegisterOptsrprgFieldBorlevShift))
}

const (
	RegisterOptsrprgFieldIwdg1hwShift = 4
	RegisterOptsrprgFieldIwdg1hwMask  = 0x10
)

// GetIwdg1hw IWDG1 option configuration bit
func (r *registerOptsrprgType) GetIwdg1hw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldIwdg1hwMask) != 0
}

// SetIwdg1hw IWDG1 option configuration bit
func (r *registerOptsrprgType) SetIwdg1hw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrprgFieldIwdg1hwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldIwdg1hwMask)
	}
}

const (
	RegisterOptsrprgFieldNrststopd1Shift = 6
	RegisterOptsrprgFieldNrststopd1Mask  = 0x40
)

// GetNrststopd1 Option byte erase after D1 DStop option configuration bit
func (r *registerOptsrprgType) GetNrststopd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldNrststopd1Mask) != 0
}

// SetNrststopd1 Option byte erase after D1 DStop option configuration bit
func (r *registerOptsrprgType) SetNrststopd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrprgFieldNrststopd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldNrststopd1Mask)
	}
}

const (
	RegisterOptsrprgFieldNrststbyd1Shift = 7
	RegisterOptsrprgFieldNrststbyd1Mask  = 0x80
)

// GetNrststbyd1 Option byte erase after D1 DStandby option configuration bit
func (r *registerOptsrprgType) GetNrststbyd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldNrststbyd1Mask) != 0
}

// SetNrststbyd1 Option byte erase after D1 DStandby option configuration bit
func (r *registerOptsrprgType) SetNrststbyd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrprgFieldNrststbyd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldNrststbyd1Mask)
	}
}

const (
	RegisterOptsrprgFieldRdpShift = 8
	RegisterOptsrprgFieldRdpMask  = 0xff00
)

// GetRdp Readout protection level option configuration byte
func (r *registerOptsrprgType) GetRdp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldRdpMask) >> RegisterOptsrprgFieldRdpShift)
}

// SetRdp Readout protection level option configuration byte
func (r *registerOptsrprgType) SetRdp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldRdpMask)|(uint32(value)<<RegisterOptsrprgFieldRdpShift))
}

const (
	RegisterOptsrprgFieldFziwdgstopShift = 17
	RegisterOptsrprgFieldFziwdgstopMask  = 0x20000
)

// GetFziwdgstop IWDG Stop mode freeze option configuration bit
func (r *registerOptsrprgType) GetFziwdgstop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldFziwdgstopMask) != 0
}

// SetFziwdgstop IWDG Stop mode freeze option configuration bit
func (r *registerOptsrprgType) SetFziwdgstop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrprgFieldFziwdgstopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldFziwdgstopMask)
	}
}

const (
	RegisterOptsrprgFieldFziwdgsdbyShift = 18
	RegisterOptsrprgFieldFziwdgsdbyMask  = 0x40000
)

// GetFziwdgsdby IWDG Standby mode freeze option configuration bit
func (r *registerOptsrprgType) GetFziwdgsdby() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldFziwdgsdbyMask) != 0
}

// SetFziwdgsdby IWDG Standby mode freeze option configuration bit
func (r *registerOptsrprgType) SetFziwdgsdby(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrprgFieldFziwdgsdbyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldFziwdgsdbyMask)
	}
}

const (
	RegisterOptsrprgFieldStramsizeShift = 19
	RegisterOptsrprgFieldStramsizeMask  = 0x180000
)

// GetStramsize DTCM size select option configuration bits
func (r *registerOptsrprgType) GetStramsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldStramsizeMask) >> RegisterOptsrprgFieldStramsizeShift)
}

// SetStramsize DTCM size select option configuration bits
func (r *registerOptsrprgType) SetStramsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldStramsizeMask)|(uint32(value)<<RegisterOptsrprgFieldStramsizeShift))
}

const (
	RegisterOptsrprgFieldSecurityShift = 21
	RegisterOptsrprgFieldSecurityMask  = 0x200000
)

// GetSecurity Security option configuration bit
func (r *registerOptsrprgType) GetSecurity() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldSecurityMask) != 0
}

// SetSecurity Security option configuration bit
func (r *registerOptsrprgType) SetSecurity(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrprgFieldSecurityMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldSecurityMask)
	}
}

const (
	RegisterOptsrprgFieldRss1Shift = 26
	RegisterOptsrprgFieldRss1Mask  = 0x4000000
)

// GetRss1 User option configuration bit 1
func (r *registerOptsrprgType) GetRss1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldRss1Mask) != 0
}

// SetRss1 User option configuration bit 1
func (r *registerOptsrprgType) SetRss1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrprgFieldRss1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldRss1Mask)
	}
}

const (
	RegisterOptsrprgFieldRss2Shift = 27
	RegisterOptsrprgFieldRss2Mask  = 0x8000000
)

// GetRss2 User option configuration bit 2
func (r *registerOptsrprgType) GetRss2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldRss2Mask) != 0
}

// SetRss2 User option configuration bit 2
func (r *registerOptsrprgType) SetRss2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrprgFieldRss2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldRss2Mask)
	}
}

const (
	RegisterOptsrprgFieldIohslvShift = 29
	RegisterOptsrprgFieldIohslvMask  = 0x20000000
)

// GetIohslv I/O high-speed at low-voltage (PRODUCT_BELOW_25V)
func (r *registerOptsrprgType) GetIohslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldIohslvMask) != 0
}

// SetIohslv I/O high-speed at low-voltage (PRODUCT_BELOW_25V)
func (r *registerOptsrprgType) SetIohslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrprgFieldIohslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldIohslvMask)
	}
}

const (
	RegisterOptsrprgFieldSwapbankoptShift = 31
	RegisterOptsrprgFieldSwapbankoptMask  = 0x80000000
)

// GetSwapbankopt Bank swapping option configuration bit
func (r *registerOptsrprgType) GetSwapbankopt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldSwapbankoptMask) != 0
}

// SetSwapbankopt Bank swapping option configuration bit
func (r *registerOptsrprgType) SetSwapbankopt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrprgFieldSwapbankoptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldSwapbankoptMask)
	}
}

// registerOptccrType FLASH option clear control register
type registerOptccrType uint32

const (
	RegisterOptccrFieldClroptchangeerrShift = 30
	RegisterOptccrFieldClroptchangeerrMask  = 0x40000000
)

// GetClroptchangeerr OPTCHANGEERR reset bit
func (r *registerOptccrType) GetClroptchangeerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptccrFieldClroptchangeerrMask) != 0
}

// SetClroptchangeerr OPTCHANGEERR reset bit
func (r *registerOptccrType) SetClroptchangeerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptccrFieldClroptchangeerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptccrFieldClroptchangeerrMask)
	}
}

// registerPrarcurType FLASH protection address for bank n
type registerPrarcurType uint32

const (
	RegisterPrarcurFieldProtareastartShift = 0
	RegisterPrarcurFieldProtareastartMask  = 0xfff
)

// GetProtareastart Bank n lowest PCROP protected address
func (r *registerPrarcurType) GetProtareastart() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarcurFieldProtareastartMask) >> RegisterPrarcurFieldProtareastartShift)
}

// SetProtareastart Bank n lowest PCROP protected address
func (r *registerPrarcurType) SetProtareastart(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarcurFieldProtareastartMask)|(uint32(value)<<RegisterPrarcurFieldProtareastartShift))
}

const (
	RegisterPrarcurFieldProtareaendShift = 16
	RegisterPrarcurFieldProtareaendMask  = 0xfff0000
)

// GetProtareaend Bank n highest PCROP protected address
func (r *registerPrarcurType) GetProtareaend() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarcurFieldProtareaendMask) >> RegisterPrarcurFieldProtareaendShift)
}

// SetProtareaend Bank n highest PCROP protected address
func (r *registerPrarcurType) SetProtareaend(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarcurFieldProtareaendMask)|(uint32(value)<<RegisterPrarcurFieldProtareaendShift))
}

const (
	RegisterPrarcurFieldDmepShift = 31
	RegisterPrarcurFieldDmepMask  = 0x80000000
)

// GetDmep Bank n PCROP protected erase enable option status bit
func (r *registerPrarcurType) GetDmep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrarcurFieldDmepMask) != 0
}

// SetDmep Bank n PCROP protected erase enable option status bit
func (r *registerPrarcurType) SetDmep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrarcurFieldDmepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrarcurFieldDmepMask)
	}
}

// registerPrarprg2Type FLASH protection address for bank 2
type registerPrarprg2Type uint32

const (
	RegisterPrarprg2FieldProtareastart2Shift = 0
	RegisterPrarprg2FieldProtareastart2Mask  = 0xfff
)

// GetProtareastart2 Bank 2 lowest PCROP protected address configuration
func (r *registerPrarprg2Type) GetProtareastart2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarprg2FieldProtareastart2Mask) >> RegisterPrarprg2FieldProtareastart2Shift)
}

// SetProtareastart2 Bank 2 lowest PCROP protected address configuration
func (r *registerPrarprg2Type) SetProtareastart2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarprg2FieldProtareastart2Mask)|(uint32(value)<<RegisterPrarprg2FieldProtareastart2Shift))
}

const (
	RegisterPrarprg2FieldProtareaend2Shift = 16
	RegisterPrarprg2FieldProtareaend2Mask  = 0xfff0000
)

// GetProtareaend2 Bank 2 highest PCROP protected address configuration
func (r *registerPrarprg2Type) GetProtareaend2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarprg2FieldProtareaend2Mask) >> RegisterPrarprg2FieldProtareaend2Shift)
}

// SetProtareaend2 Bank 2 highest PCROP protected address configuration
func (r *registerPrarprg2Type) SetProtareaend2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarprg2FieldProtareaend2Mask)|(uint32(value)<<RegisterPrarprg2FieldProtareaend2Shift))
}

const (
	RegisterPrarprg2FieldDmep2Shift = 31
	RegisterPrarprg2FieldDmep2Mask  = 0x80000000
)

// GetDmep2 Bank 2 PCROP protected erase enable option configuration bit
func (r *registerPrarprg2Type) GetDmep2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrarprg2FieldDmep2Mask) != 0
}

// SetDmep2 Bank 2 PCROP protected erase enable option configuration bit
func (r *registerPrarprg2Type) SetDmep2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrarprg2FieldDmep2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrarprg2FieldDmep2Mask)
	}
}

// registerPrarprgType FLASH protection address for bank n
type registerPrarprgType uint32

const (
	RegisterPrarprgFieldProtareastartShift = 0
	RegisterPrarprgFieldProtareastartMask  = 0xfff
)

// GetProtareastart Bank n lowest PCROP protected address configuration
func (r *registerPrarprgType) GetProtareastart() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarprgFieldProtareastartMask) >> RegisterPrarprgFieldProtareastartShift)
}

// SetProtareastart Bank n lowest PCROP protected address configuration
func (r *registerPrarprgType) SetProtareastart(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarprgFieldProtareastartMask)|(uint32(value)<<RegisterPrarprgFieldProtareastartShift))
}

const (
	RegisterPrarprgFieldProtareaendShift = 16
	RegisterPrarprgFieldProtareaendMask  = 0xfff0000
)

// GetProtareaend Bank n highest PCROP protected address configuration
func (r *registerPrarprgType) GetProtareaend() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarprgFieldProtareaendMask) >> RegisterPrarprgFieldProtareaendShift)
}

// SetProtareaend Bank n highest PCROP protected address configuration
func (r *registerPrarprgType) SetProtareaend(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarprgFieldProtareaendMask)|(uint32(value)<<RegisterPrarprgFieldProtareaendShift))
}

const (
	RegisterPrarprgFieldDmepShift = 31
	RegisterPrarprgFieldDmepMask  = 0x80000000
)

// GetDmep Bank n PCROP protected erase enable option configuration bit
func (r *registerPrarprgType) GetDmep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrarprgFieldDmepMask) != 0
}

// SetDmep Bank n PCROP protected erase enable option configuration bit
func (r *registerPrarprgType) SetDmep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrarprgFieldDmepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrarprgFieldDmepMask)
	}
}

// registerScarcurType FLASH secure address for bank n
type registerScarcurType uint32

const (
	RegisterScarcurFieldSecareastartShift = 0
	RegisterScarcurFieldSecareastartMask  = 0xfff
)

// GetSecareastart Bank n lowest secure protected address
func (r *registerScarcurType) GetSecareastart() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarcurFieldSecareastartMask) >> RegisterScarcurFieldSecareastartShift)
}

// SetSecareastart Bank n lowest secure protected address
func (r *registerScarcurType) SetSecareastart(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarcurFieldSecareastartMask)|(uint32(value)<<RegisterScarcurFieldSecareastartShift))
}

const (
	RegisterScarcurFieldSecareaendShift = 16
	RegisterScarcurFieldSecareaendMask  = 0xfff0000
)

// GetSecareaend Bank n highest secure protected address
func (r *registerScarcurType) GetSecareaend() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarcurFieldSecareaendMask) >> RegisterScarcurFieldSecareaendShift)
}

// SetSecareaend Bank n highest secure protected address
func (r *registerScarcurType) SetSecareaend(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarcurFieldSecareaendMask)|(uint32(value)<<RegisterScarcurFieldSecareaendShift))
}

const (
	RegisterScarcurFieldDmesShift = 31
	RegisterScarcurFieldDmesMask  = 0x80000000
)

// GetDmes Bank n secure protected erase enable option status bit
func (r *registerScarcurType) GetDmes() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScarcurFieldDmesMask) != 0
}

// SetDmes Bank n secure protected erase enable option status bit
func (r *registerScarcurType) SetDmes(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScarcurFieldDmesMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScarcurFieldDmesMask)
	}
}

// registerScarprgType FLASH secure address for bank n
type registerScarprgType uint32

const (
	RegisterScarprgFieldSecareastartShift = 0
	RegisterScarprgFieldSecareastartMask  = 0xfff
)

// GetSecareastart Bank n lowest secure protected address configuration
func (r *registerScarprgType) GetSecareastart() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarprgFieldSecareastartMask) >> RegisterScarprgFieldSecareastartShift)
}

// SetSecareastart Bank n lowest secure protected address configuration
func (r *registerScarprgType) SetSecareastart(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarprgFieldSecareastartMask)|(uint32(value)<<RegisterScarprgFieldSecareastartShift))
}

const (
	RegisterScarprgFieldSecareaendShift = 16
	RegisterScarprgFieldSecareaendMask  = 0xfff0000
)

// GetSecareaend Bank n highest secure protected address configuration
func (r *registerScarprgType) GetSecareaend() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarprgFieldSecareaendMask) >> RegisterScarprgFieldSecareaendShift)
}

// SetSecareaend Bank n highest secure protected address configuration
func (r *registerScarprgType) SetSecareaend(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarprgFieldSecareaendMask)|(uint32(value)<<RegisterScarprgFieldSecareaendShift))
}

const (
	RegisterScarprgFieldDmesShift = 31
	RegisterScarprgFieldDmesMask  = 0x80000000
)

// GetDmes Bank n secure protected erase enable option configuration bit
func (r *registerScarprgType) GetDmes() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScarprgFieldDmesMask) != 0
}

// SetDmes Bank n secure protected erase enable option configuration bit
func (r *registerScarprgType) SetDmes(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScarprgFieldDmesMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScarprgFieldDmesMask)
	}
}

// registerWpsncur1rType FLASH write sector protection for bank n
type registerWpsncur1rType uint32

const (
	RegisterWpsncur1rFieldWrpsnShift = 0
	RegisterWpsncur1rFieldWrpsnMask  = 0xff
)

// GetWrpsn Bank n sector write protection option status byte
func (r *registerWpsncur1rType) GetWrpsn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWpsncur1rFieldWrpsnMask) >> RegisterWpsncur1rFieldWrpsnShift)
}

// SetWrpsn Bank n sector write protection option status byte
func (r *registerWpsncur1rType) SetWrpsn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWpsncur1rFieldWrpsnMask)|(uint32(value)<<RegisterWpsncur1rFieldWrpsnShift))
}

// registerWpsnprg1rType FLASH write sector protection for bank n
type registerWpsnprg1rType uint32

const (
	RegisterWpsnprg1rFieldWrpsnShift = 0
	RegisterWpsnprg1rFieldWrpsnMask  = 0xff
)

// GetWrpsn Bank n sector write protection configuration byte
func (r *registerWpsnprg1rType) GetWrpsn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWpsnprg1rFieldWrpsnMask) >> RegisterWpsnprg1rFieldWrpsnShift)
}

// SetWrpsn Bank n sector write protection configuration byte
func (r *registerWpsnprg1rType) SetWrpsn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWpsnprg1rFieldWrpsnMask)|(uint32(value)<<RegisterWpsnprg1rFieldWrpsnShift))
}

// registerBootcurrType FLASH register with boot address
type registerBootcurrType uint32

const (
	RegisterBootcurrFieldBootadd0Shift = 0
	RegisterBootcurrFieldBootadd0Mask  = 0xffff
)

// GetBootadd0 Boot address 0
func (r *registerBootcurrType) GetBootadd0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBootcurrFieldBootadd0Mask) >> RegisterBootcurrFieldBootadd0Shift)
}

// SetBootadd0 Boot address 0
func (r *registerBootcurrType) SetBootadd0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBootcurrFieldBootadd0Mask)|(uint32(value)<<RegisterBootcurrFieldBootadd0Shift))
}

const (
	RegisterBootcurrFieldBootadd1Shift = 16
	RegisterBootcurrFieldBootadd1Mask  = 0xffff0000
)

// GetBootadd1 Boot address 1
func (r *registerBootcurrType) GetBootadd1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBootcurrFieldBootadd1Mask) >> RegisterBootcurrFieldBootadd1Shift)
}

// SetBootadd1 Boot address 1
func (r *registerBootcurrType) SetBootadd1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBootcurrFieldBootadd1Mask)|(uint32(value)<<RegisterBootcurrFieldBootadd1Shift))
}

// registerBootprgrType FLASH register with boot address
type registerBootprgrType uint32

const (
	RegisterBootprgrFieldBootadd0Shift = 0
	RegisterBootprgrFieldBootadd0Mask  = 0xffff
)

// GetBootadd0 Boot address 0
func (r *registerBootprgrType) GetBootadd0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBootprgrFieldBootadd0Mask) >> RegisterBootprgrFieldBootadd0Shift)
}

// SetBootadd0 Boot address 0
func (r *registerBootprgrType) SetBootadd0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBootprgrFieldBootadd0Mask)|(uint32(value)<<RegisterBootprgrFieldBootadd0Shift))
}

const (
	RegisterBootprgrFieldBootadd1Shift = 16
	RegisterBootprgrFieldBootadd1Mask  = 0xffff0000
)

// GetBootadd1 Boot address 1
func (r *registerBootprgrType) GetBootadd1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBootprgrFieldBootadd1Mask) >> RegisterBootprgrFieldBootadd1Shift)
}

// SetBootadd1 Boot address 1
func (r *registerBootprgrType) SetBootadd1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBootprgrFieldBootadd1Mask)|(uint32(value)<<RegisterBootprgrFieldBootadd1Shift))
}

// registerCrccrType FLASH CRC control register for bank n
type registerCrccrType uint32

const (
	RegisterCrccrFieldCrcsectShift = 0
	RegisterCrccrFieldCrcsectMask  = 0x7
)

// GetCrcsect Bank n CRC sector number
func (r *registerCrccrType) GetCrcsect() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldCrcsectMask) >> RegisterCrccrFieldCrcsectShift)
}

// SetCrcsect Bank n CRC sector number
func (r *registerCrccrType) SetCrcsect(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrccrFieldCrcsectMask)|(uint32(value)<<RegisterCrccrFieldCrcsectShift))
}

const (
	RegisterCrccrFieldAllbankShift = 7
	RegisterCrccrFieldAllbankMask  = 0x80
)

// GetAllbank Bank n CRC select bit
func (r *registerCrccrType) GetAllbank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldAllbankMask) != 0
}

// SetAllbank Bank n CRC select bit
func (r *registerCrccrType) SetAllbank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccrFieldAllbankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccrFieldAllbankMask)
	}
}

const (
	RegisterCrccrFieldCrcbysectShift = 8
	RegisterCrccrFieldCrcbysectMask  = 0x100
)

// GetCrcbysect Bank n CRC sector mode select bit
func (r *registerCrccrType) GetCrcbysect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldCrcbysectMask) != 0
}

// SetCrcbysect Bank n CRC sector mode select bit
func (r *registerCrccrType) SetCrcbysect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccrFieldCrcbysectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccrFieldCrcbysectMask)
	}
}

const (
	RegisterCrccrFieldAddsectShift = 9
	RegisterCrccrFieldAddsectMask  = 0x200
)

// GetAddsect Bank n CRC sector select bit
func (r *registerCrccrType) GetAddsect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldAddsectMask) != 0
}

// SetAddsect Bank n CRC sector select bit
func (r *registerCrccrType) SetAddsect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccrFieldAddsectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccrFieldAddsectMask)
	}
}

const (
	RegisterCrccrFieldCleansectShift = 10
	RegisterCrccrFieldCleansectMask  = 0x400
)

// GetCleansect Bank n CRC sector list clear bit
func (r *registerCrccrType) GetCleansect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldCleansectMask) != 0
}

// SetCleansect Bank n CRC sector list clear bit
func (r *registerCrccrType) SetCleansect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccrFieldCleansectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccrFieldCleansectMask)
	}
}

const (
	RegisterCrccrFieldStartcrcShift = 16
	RegisterCrccrFieldStartcrcMask  = 0x10000
)

// GetStartcrc Bank n CRC start bit
func (r *registerCrccrType) GetStartcrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldStartcrcMask) != 0
}

// SetStartcrc Bank n CRC start bit
func (r *registerCrccrType) SetStartcrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccrFieldStartcrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccrFieldStartcrcMask)
	}
}

const (
	RegisterCrccrFieldCleancrcShift = 17
	RegisterCrccrFieldCleancrcMask  = 0x20000
)

// GetCleancrc Bank n CRC clear bit
func (r *registerCrccrType) GetCleancrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldCleancrcMask) != 0
}

// SetCleancrc Bank n CRC clear bit
func (r *registerCrccrType) SetCleancrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccrFieldCleancrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccrFieldCleancrcMask)
	}
}

const (
	RegisterCrccrFieldCrcburstShift = 20
	RegisterCrccrFieldCrcburstMask  = 0x300000
)

// GetCrcburst Bank n CRC burst size
func (r *registerCrccrType) GetCrcburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldCrcburstMask) >> RegisterCrccrFieldCrcburstShift)
}

// SetCrcburst Bank n CRC burst size
func (r *registerCrccrType) SetCrcburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrccrFieldCrcburstMask)|(uint32(value)<<RegisterCrccrFieldCrcburstShift))
}

// registerCrcsadd1rType FLASH CRC start address register for bank n
type registerCrcsadd1rType uint32

const (
	RegisterCrcsadd1rFieldCrcstartaddrShift = 0
	RegisterCrcsadd1rFieldCrcstartaddrMask  = 0xffffffff
)

// GetCrcstartaddr CRC start address on bank n
func (r *registerCrcsadd1rType) GetCrcstartaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrcsadd1rFieldCrcstartaddrMask) >> RegisterCrcsadd1rFieldCrcstartaddrShift)
}

// SetCrcstartaddr CRC start address on bank n
func (r *registerCrcsadd1rType) SetCrcstartaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrcsadd1rFieldCrcstartaddrMask)|(uint32(value)<<RegisterCrcsadd1rFieldCrcstartaddrShift))
}

// registerCrceadd1rType FLASH CRC end address register for bank n
type registerCrceadd1rType uint32

const (
	RegisterCrceadd1rFieldCrcendaddrShift = 0
	RegisterCrceadd1rFieldCrcendaddrMask  = 0xffffffff
)

// GetCrcendaddr CRC end address on bank n
func (r *registerCrceadd1rType) GetCrcendaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrceadd1rFieldCrcendaddrMask) >> RegisterCrceadd1rFieldCrcendaddrShift)
}

// SetCrcendaddr CRC end address on bank n
func (r *registerCrceadd1rType) SetCrcendaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrceadd1rFieldCrcendaddrMask)|(uint32(value)<<RegisterCrceadd1rFieldCrcendaddrShift))
}

// registerCrcdatarType FLASH CRC data register
type registerCrcdatarType uint32

const (
	RegisterCrcdatarFieldCrcdataShift = 0
	RegisterCrcdatarFieldCrcdataMask  = 0xffffffff
)

// GetCrcdata CRC result
func (r *registerCrcdatarType) GetCrcdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrcdatarFieldCrcdataMask) >> RegisterCrcdatarFieldCrcdataShift)
}

// SetCrcdata CRC result
func (r *registerCrcdatarType) SetCrcdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrcdatarFieldCrcdataMask)|(uint32(value)<<RegisterCrcdatarFieldCrcdataShift))
}

// registerEccfa1rType FLASH ECC fail address for bank n
type registerEccfa1rType uint32

const (
	RegisterEccfa1rFieldFaileccaddrShift = 0
	RegisterEccfa1rFieldFaileccaddrMask  = 0x7fff
)

// GetFaileccaddr Bank n ECC error address
func (r *registerEccfa1rType) GetFaileccaddr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterEccfa1rFieldFaileccaddrMask) >> RegisterEccfa1rFieldFaileccaddrShift)
}

// SetFaileccaddr Bank n ECC error address
func (r *registerEccfa1rType) SetFaileccaddr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEccfa1rFieldFaileccaddrMask)|(uint32(value)<<RegisterEccfa1rFieldFaileccaddrShift))
}
