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
	Acr       RegisterAcrType
	Keyr      RegisterKeyrType
	Optkeyr   RegisterOptkeyrType
	Cr        RegisterCrType
	Sr        RegisterSrType
	Ccr       RegisterCcrType
	Optcr     RegisterOptcrType
	Optsrcur  RegisterOptsrcurType
	Optsrprg  RegisterOptsrprgType
	Optccr    RegisterOptccrType
	Prarcur   RegisterPrarcurType
	Prarprg2  RegisterPrarprg2Type
	Prarprg   RegisterPrarprgType
	Scarcur   RegisterScarcurType
	Scarprg   RegisterScarprgType
	Wpsncur1r RegisterWpsncur1rType
	Wpsnprg1r RegisterWpsnprg1rType
	Bootcurr  RegisterBootcurrType
	Bootprgr  RegisterBootprgrType
	_         [8]byte
	Crccr     RegisterCrccrType
	Crcsadd1r RegisterCrcsadd1rType
	Crceadd1r RegisterCrceadd1rType
	Crcdatar  RegisterCrcdatarType
	Eccfa1r   RegisterEccfa1rType
}

// RegisterAcrType Access control register
type RegisterAcrType uint32

func (r *RegisterAcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAcrFieldLatencyShift = 0
	RegisterAcrFieldLatencyMask  = 0x7
)

// GetLatency Read latency
func (r *RegisterAcrType) GetLatency() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAcrFieldLatencyMask) >> RegisterAcrFieldLatencyShift)
}

// SetLatency Read latency
func (r *RegisterAcrType) SetLatency(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAcrFieldLatencyMask)|(uint32(value)<<RegisterAcrFieldLatencyShift))
}

const (
	RegisterAcrFieldWrhighfreqShift = 4
	RegisterAcrFieldWrhighfreqMask  = 0x30
)

// GetWrhighfreq Flash signal delay
func (r *RegisterAcrType) GetWrhighfreq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAcrFieldWrhighfreqMask) >> RegisterAcrFieldWrhighfreqShift)
}

// SetWrhighfreq Flash signal delay
func (r *RegisterAcrType) SetWrhighfreq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAcrFieldWrhighfreqMask)|(uint32(value)<<RegisterAcrFieldWrhighfreqShift))
}

// RegisterKeyrType FLASH key register for bank n
type RegisterKeyrType uint32

func (r *RegisterKeyrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterKeyrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterKeyrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterKeyrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterKeyrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterKeyrFieldKeyrShift = 0
	RegisterKeyrFieldKeyrMask  = 0xffffffff
)

// GetKeyr Bank n access configuration unlock key
func (r *RegisterKeyrType) GetKeyr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterKeyrFieldKeyrMask) >> RegisterKeyrFieldKeyrShift)
}

// SetKeyr Bank n access configuration unlock key
func (r *RegisterKeyrType) SetKeyr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterKeyrFieldKeyrMask)|(uint32(value)<<RegisterKeyrFieldKeyrShift))
}

// RegisterOptkeyrType FLASH option key register
type RegisterOptkeyrType uint32

func (r *RegisterOptkeyrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOptkeyrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOptkeyrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOptkeyrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOptkeyrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOptkeyrFieldOptkeyrShift = 0
	RegisterOptkeyrFieldOptkeyrMask  = 0xffffffff
)

// GetOptkeyr Unlock key option bytes
func (r *RegisterOptkeyrType) GetOptkeyr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOptkeyrFieldOptkeyrMask) >> RegisterOptkeyrFieldOptkeyrShift)
}

// SetOptkeyr Unlock key option bytes
func (r *RegisterOptkeyrType) SetOptkeyr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptkeyrFieldOptkeyrMask)|(uint32(value)<<RegisterOptkeyrFieldOptkeyrShift))
}

// RegisterCrType FLASH control register for bank n
type RegisterCrType uint32

func (r *RegisterCrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrFieldLockShift = 0
	RegisterCrFieldLockMask  = 0x1
)

// GetLock Bank n configuration lock bit
func (r *RegisterCrType) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLockMask) != 0
}

// SetLock Bank n configuration lock bit
func (r *RegisterCrType) SetLock(value bool) {
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
func (r *RegisterCrType) GetPg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPgMask) != 0
}

// SetPg Bank n program enable bit
func (r *RegisterCrType) SetPg(value bool) {
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
func (r *RegisterCrType) GetSer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSerMask) != 0
}

// SetSer Bank n sector erase request
func (r *RegisterCrType) SetSer(value bool) {
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
func (r *RegisterCrType) GetBer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldBerMask) != 0
}

// SetBer Bank n erase request
func (r *RegisterCrType) SetBer(value bool) {
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
func (r *RegisterCrType) GetPsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPsizeMask) >> RegisterCrFieldPsizeShift)
}

// SetPsize Bank n program size
func (r *RegisterCrType) SetPsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPsizeMask)|(uint32(value)<<RegisterCrFieldPsizeShift))
}

const (
	RegisterCrFieldFwShift = 6
	RegisterCrFieldFwMask  = 0x40
)

// GetFw Bank n write forcing control bit
func (r *RegisterCrType) GetFw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldFwMask) != 0
}

// SetFw Bank n write forcing control bit
func (r *RegisterCrType) SetFw(value bool) {
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
func (r *RegisterCrType) GetStart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldStartMask) != 0
}

// SetStart Bank n bank or sector erase start control bit
func (r *RegisterCrType) SetStart(value bool) {
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
func (r *RegisterCrType) GetSnb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSnbMask) >> RegisterCrFieldSnbShift)
}

// SetSnb Bank n sector erase selection number
func (r *RegisterCrType) SetSnb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSnbMask)|(uint32(value)<<RegisterCrFieldSnbShift))
}

const (
	RegisterCrFieldCrcenShift = 15
	RegisterCrFieldCrcenMask  = 0x8000
)

// GetCrcen Bank n CRC control bit
func (r *RegisterCrType) GetCrcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCrcenMask) != 0
}

// SetCrcen Bank n CRC control bit
func (r *RegisterCrType) SetCrcen(value bool) {
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
func (r *RegisterCrType) GetEopie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEopieMask) != 0
}

// SetEopie Bank n end-of-program interrupt control bit
func (r *RegisterCrType) SetEopie(value bool) {
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
func (r *RegisterCrType) GetWrperrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldWrperrieMask) != 0
}

// SetWrperrie Bank n write protection error interrupt enable bit
func (r *RegisterCrType) SetWrperrie(value bool) {
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
func (r *RegisterCrType) GetPgserrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPgserrieMask) != 0
}

// SetPgserrie Bank n programming sequence error interrupt enable bit
func (r *RegisterCrType) SetPgserrie(value bool) {
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
func (r *RegisterCrType) GetStrberrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldStrberrieMask) != 0
}

// SetStrberrie Bank n strobe error interrupt enable bit
func (r *RegisterCrType) SetStrberrie(value bool) {
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
func (r *RegisterCrType) GetIncerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldIncerrieMask) != 0
}

// SetIncerrie Bank n inconsistency error interrupt enable bit
func (r *RegisterCrType) SetIncerrie(value bool) {
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
func (r *RegisterCrType) GetOperrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldOperrieMask) != 0
}

// SetOperrie Bank n write/erase error interrupt enable bit
func (r *RegisterCrType) SetOperrie(value bool) {
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
func (r *RegisterCrType) GetRdperrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRdperrieMask) != 0
}

// SetRdperrie Bank n read protection error interrupt enable bit
func (r *RegisterCrType) SetRdperrie(value bool) {
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
func (r *RegisterCrType) GetRdserrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRdserrieMask) != 0
}

// SetRdserrie Bank n secure error interrupt enable bit
func (r *RegisterCrType) SetRdserrie(value bool) {
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
func (r *RegisterCrType) GetSneccerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSneccerrieMask) != 0
}

// SetSneccerrie Bank n ECC single correction error interrupt enable bit
func (r *RegisterCrType) SetSneccerrie(value bool) {
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
func (r *RegisterCrType) GetDbeccerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDbeccerrieMask) != 0
}

// SetDbeccerrie Bank n ECC double detection error interrupt enable bit
func (r *RegisterCrType) SetDbeccerrie(value bool) {
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
func (r *RegisterCrType) GetCrcendie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCrcendieMask) != 0
}

// SetCrcendie Bank n end of CRC calculation interrupt enable bit
func (r *RegisterCrType) SetCrcendie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCrcendieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCrcendieMask)
	}
}

// RegisterSrType FLASH status register for bank n
type RegisterSrType uint32

func (r *RegisterSrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSrFieldBsyShift = 0
	RegisterSrFieldBsyMask  = 0x1
)

// GetBsy Bank n ongoing program flag
func (r *RegisterSrType) GetBsy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldBsyMask) != 0
}

// SetBsy Bank n ongoing program flag
func (r *RegisterSrType) SetBsy(value bool) {
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
func (r *RegisterSrType) GetWbne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldWbneMask) != 0
}

// SetWbne Bank n write buffer not empty flag
func (r *RegisterSrType) SetWbne(value bool) {
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
func (r *RegisterSrType) GetQw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldQwMask) != 0
}

// SetQw Bank n wait queue flag
func (r *RegisterSrType) SetQw(value bool) {
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
func (r *RegisterSrType) GetCrcbusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCrcbusyMask) != 0
}

// SetCrcbusy Bank n CRC busy flag
func (r *RegisterSrType) SetCrcbusy(value bool) {
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
func (r *RegisterSrType) GetEop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldEopMask) != 0
}

// SetEop Bank n end-of-program flag
func (r *RegisterSrType) SetEop(value bool) {
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
func (r *RegisterSrType) GetWrperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldWrperrMask) != 0
}

// SetWrperr Bank n write protection error flag
func (r *RegisterSrType) SetWrperr(value bool) {
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
func (r *RegisterSrType) GetPgserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldPgserrMask) != 0
}

// SetPgserr Bank n programming sequence error flag
func (r *RegisterSrType) SetPgserr(value bool) {
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
func (r *RegisterSrType) GetStrberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldStrberrMask) != 0
}

// SetStrberr Bank n strobe error flag
func (r *RegisterSrType) SetStrberr(value bool) {
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
func (r *RegisterSrType) GetIncerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldIncerrMask) != 0
}

// SetIncerr Bank n inconsistency error flag
func (r *RegisterSrType) SetIncerr(value bool) {
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
func (r *RegisterSrType) GetOperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldOperrMask) != 0
}

// SetOperr Bank n write/erase error flag
func (r *RegisterSrType) SetOperr(value bool) {
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
func (r *RegisterSrType) GetRdperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRdperrMask) != 0
}

// SetRdperr Bank n read protection error flag
func (r *RegisterSrType) SetRdperr(value bool) {
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
func (r *RegisterSrType) GetRdserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRdserrMask) != 0
}

// SetRdserr Bank n secure error flag
func (r *RegisterSrType) SetRdserr(value bool) {
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
func (r *RegisterSrType) GetSneccerr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSneccerr1Mask) != 0
}

// SetSneccerr1 Bank n single correction error flag
func (r *RegisterSrType) SetSneccerr1(value bool) {
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
func (r *RegisterSrType) GetDbeccerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDbeccerrMask) != 0
}

// SetDbeccerr Bank n ECC double detection error flag
func (r *RegisterSrType) SetDbeccerr(value bool) {
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
func (r *RegisterSrType) GetCrcend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCrcendMask) != 0
}

// SetCrcend Bank n CRC-complete flag
func (r *RegisterSrType) SetCrcend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCrcendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCrcendMask)
	}
}

// RegisterCcrType FLASH clear control register for bank n
type RegisterCcrType uint32

func (r *RegisterCcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcrFieldClreopShift = 16
	RegisterCcrFieldClreopMask  = 0x10000
)

// GetClreop Bank n EOP1 flag clear bit
func (r *RegisterCcrType) GetClreop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClreopMask) != 0
}

// SetClreop Bank n EOP1 flag clear bit
func (r *RegisterCcrType) SetClreop(value bool) {
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
func (r *RegisterCcrType) GetClrwrperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrwrperrMask) != 0
}

// SetClrwrperr Bank n WRPERR1 flag clear bit
func (r *RegisterCcrType) SetClrwrperr(value bool) {
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
func (r *RegisterCcrType) GetClrpgserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrpgserrMask) != 0
}

// SetClrpgserr Bank n PGSERR1 flag clear bi
func (r *RegisterCcrType) SetClrpgserr(value bool) {
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
func (r *RegisterCcrType) GetClrstrberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrstrberrMask) != 0
}

// SetClrstrberr Bank n STRBERR1 flag clear bit
func (r *RegisterCcrType) SetClrstrberr(value bool) {
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
func (r *RegisterCcrType) GetClrincerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrincerrMask) != 0
}

// SetClrincerr Bank n INCERR1 flag clear bit
func (r *RegisterCcrType) SetClrincerr(value bool) {
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
func (r *RegisterCcrType) GetClroperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClroperrMask) != 0
}

// SetClroperr Bank n OPERR1 flag clear bit
func (r *RegisterCcrType) SetClroperr(value bool) {
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
func (r *RegisterCcrType) GetClrrdperr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrrdperrMask) != 0
}

// SetClrrdperr Bank n RDPERR1 flag clear bit
func (r *RegisterCcrType) SetClrrdperr(value bool) {
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
func (r *RegisterCcrType) GetClrrdserr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrrdserrMask) != 0
}

// SetClrrdserr Bank n RDSERR1 flag clear bit
func (r *RegisterCcrType) SetClrrdserr(value bool) {
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
func (r *RegisterCcrType) GetClrsneccerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrsneccerrMask) != 0
}

// SetClrsneccerr Bank n SNECCERR1 flag clear bit
func (r *RegisterCcrType) SetClrsneccerr(value bool) {
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
func (r *RegisterCcrType) GetClrdbeccerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrdbeccerrMask) != 0
}

// SetClrdbeccerr Bank n DBECCERR1 flag clear bit
func (r *RegisterCcrType) SetClrdbeccerr(value bool) {
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
func (r *RegisterCcrType) GetClrcrcend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldClrcrcendMask) != 0
}

// SetClrcrcend Bank n CRCEND1 flag clear bit
func (r *RegisterCcrType) SetClrcrcend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldClrcrcendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldClrcrcendMask)
	}
}

// RegisterOptcrType FLASH option control register
type RegisterOptcrType uint32

func (r *RegisterOptcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOptcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOptcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOptcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOptcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOptcrFieldOptlockShift = 0
	RegisterOptcrFieldOptlockMask  = 0x1
)

// GetOptlock FLASH_OPTCR lock option configuration bit
func (r *RegisterOptcrType) GetOptlock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptcrFieldOptlockMask) != 0
}

// SetOptlock FLASH_OPTCR lock option configuration bit
func (r *RegisterOptcrType) SetOptlock(value bool) {
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
func (r *RegisterOptcrType) GetOptstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptcrFieldOptstartMask) != 0
}

// SetOptstart Option byte start change option configuration bit
func (r *RegisterOptcrType) SetOptstart(value bool) {
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
func (r *RegisterOptcrType) GetMer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptcrFieldMerMask) != 0
}

// SetMer Flash mass erase enable bit
func (r *RegisterOptcrType) SetMer(value bool) {
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
func (r *RegisterOptcrType) GetOptchangeerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptcrFieldOptchangeerrieMask) != 0
}

// SetOptchangeerrie Option byte change error interrupt enable bit
func (r *RegisterOptcrType) SetOptchangeerrie(value bool) {
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
func (r *RegisterOptcrType) GetSwapbank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptcrFieldSwapbankMask) != 0
}

// SetSwapbank Bank swapping configuration bit
func (r *RegisterOptcrType) SetSwapbank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptcrFieldSwapbankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptcrFieldSwapbankMask)
	}
}

// RegisterOptsrcurType FLASH option status register
type RegisterOptsrcurType uint32

func (r *RegisterOptsrcurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOptsrcurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOptsrcurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOptsrcurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOptsrcurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOptsrcurFieldOptbusyShift = 0
	RegisterOptsrcurFieldOptbusyMask  = 0x1
)

// GetOptbusy Option byte change ongoing flag
func (r *RegisterOptsrcurType) GetOptbusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldOptbusyMask) != 0
}

// SetOptbusy Option byte change ongoing flag
func (r *RegisterOptsrcurType) SetOptbusy(value bool) {
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
func (r *RegisterOptsrcurType) GetBorlev() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldBorlevMask) >> RegisterOptsrcurFieldBorlevShift)
}

// SetBorlev Brownout level option status bit
func (r *RegisterOptsrcurType) SetBorlev(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldBorlevMask)|(uint32(value)<<RegisterOptsrcurFieldBorlevShift))
}

const (
	RegisterOptsrcurFieldIwdg1hwShift = 4
	RegisterOptsrcurFieldIwdg1hwMask  = 0x10
)

// GetIwdg1hw IWDG1 control option status bit
func (r *RegisterOptsrcurType) GetIwdg1hw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldIwdg1hwMask) != 0
}

// SetIwdg1hw IWDG1 control option status bit
func (r *RegisterOptsrcurType) SetIwdg1hw(value bool) {
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
func (r *RegisterOptsrcurType) GetNrststopd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldNrststopdMask) != 0
}

// SetNrststopd D1 DStop entry reset option status bit
func (r *RegisterOptsrcurType) SetNrststopd(value bool) {
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
func (r *RegisterOptsrcurType) GetNrststbyd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldNrststbydMask) != 0
}

// SetNrststbyd D1 DStandby entry reset option status bit
func (r *RegisterOptsrcurType) SetNrststbyd(value bool) {
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
func (r *RegisterOptsrcurType) GetRdp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldRdpMask) >> RegisterOptsrcurFieldRdpShift)
}

// SetRdp Readout protection level option status byte
func (r *RegisterOptsrcurType) SetRdp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldRdpMask)|(uint32(value)<<RegisterOptsrcurFieldRdpShift))
}

const (
	RegisterOptsrcurFieldFziwdgstopShift = 17
	RegisterOptsrcurFieldFziwdgstopMask  = 0x20000
)

// GetFziwdgstop IWDG Stop mode freeze option status bit
func (r *RegisterOptsrcurType) GetFziwdgstop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldFziwdgstopMask) != 0
}

// SetFziwdgstop IWDG Stop mode freeze option status bit
func (r *RegisterOptsrcurType) SetFziwdgstop(value bool) {
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
func (r *RegisterOptsrcurType) GetFziwdgsdby() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldFziwdgsdbyMask) != 0
}

// SetFziwdgsdby IWDG Standby mode freeze option status bit
func (r *RegisterOptsrcurType) SetFziwdgsdby(value bool) {
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
func (r *RegisterOptsrcurType) GetStramsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldStramsizeMask) >> RegisterOptsrcurFieldStramsizeShift)
}

// SetStramsize DTCM RAM size option status
func (r *RegisterOptsrcurType) SetStramsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldStramsizeMask)|(uint32(value)<<RegisterOptsrcurFieldStramsizeShift))
}

const (
	RegisterOptsrcurFieldSecurityShift = 21
	RegisterOptsrcurFieldSecurityMask  = 0x200000
)

// GetSecurity Security enable option status bit
func (r *RegisterOptsrcurType) GetSecurity() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldSecurityMask) != 0
}

// SetSecurity Security enable option status bit
func (r *RegisterOptsrcurType) SetSecurity(value bool) {
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
func (r *RegisterOptsrcurType) GetRss1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldRss1Mask) != 0
}

// SetRss1 User option bit 1
func (r *RegisterOptsrcurType) SetRss1(value bool) {
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
func (r *RegisterOptsrcurType) GetPersook() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldPersookMask) != 0
}

// SetPersook Device personalization status bit
func (r *RegisterOptsrcurType) SetPersook(value bool) {
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
func (r *RegisterOptsrcurType) GetIohslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldIohslvMask) != 0
}

// SetIohslv I/O high-speed at low-voltage status bit (PRODUCT_BELOW_25V)
func (r *RegisterOptsrcurType) SetIohslv(value bool) {
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
func (r *RegisterOptsrcurType) GetOptchangeerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldOptchangeerrMask) != 0
}

// SetOptchangeerr Option byte change error flag
func (r *RegisterOptsrcurType) SetOptchangeerr(value bool) {
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
func (r *RegisterOptsrcurType) GetSwapbankopt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldSwapbankoptMask) != 0
}

// SetSwapbankopt Bank swapping option status bit
func (r *RegisterOptsrcurType) SetSwapbankopt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldSwapbankoptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldSwapbankoptMask)
	}
}

// RegisterOptsrprgType FLASH option status register
type RegisterOptsrprgType uint32

func (r *RegisterOptsrprgType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOptsrprgType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOptsrprgType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOptsrprgType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOptsrprgType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOptsrprgFieldBorlevShift = 2
	RegisterOptsrprgFieldBorlevMask  = 0xc
)

// GetBorlev BOR reset level option configuration bits
func (r *RegisterOptsrprgType) GetBorlev() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldBorlevMask) >> RegisterOptsrprgFieldBorlevShift)
}

// SetBorlev BOR reset level option configuration bits
func (r *RegisterOptsrprgType) SetBorlev(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldBorlevMask)|(uint32(value)<<RegisterOptsrprgFieldBorlevShift))
}

const (
	RegisterOptsrprgFieldIwdg1hwShift = 4
	RegisterOptsrprgFieldIwdg1hwMask  = 0x10
)

// GetIwdg1hw IWDG1 option configuration bit
func (r *RegisterOptsrprgType) GetIwdg1hw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldIwdg1hwMask) != 0
}

// SetIwdg1hw IWDG1 option configuration bit
func (r *RegisterOptsrprgType) SetIwdg1hw(value bool) {
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
func (r *RegisterOptsrprgType) GetNrststopd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldNrststopd1Mask) != 0
}

// SetNrststopd1 Option byte erase after D1 DStop option configuration bit
func (r *RegisterOptsrprgType) SetNrststopd1(value bool) {
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
func (r *RegisterOptsrprgType) GetNrststbyd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldNrststbyd1Mask) != 0
}

// SetNrststbyd1 Option byte erase after D1 DStandby option configuration bit
func (r *RegisterOptsrprgType) SetNrststbyd1(value bool) {
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
func (r *RegisterOptsrprgType) GetRdp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldRdpMask) >> RegisterOptsrprgFieldRdpShift)
}

// SetRdp Readout protection level option configuration byte
func (r *RegisterOptsrprgType) SetRdp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldRdpMask)|(uint32(value)<<RegisterOptsrprgFieldRdpShift))
}

const (
	RegisterOptsrprgFieldFziwdgstopShift = 17
	RegisterOptsrprgFieldFziwdgstopMask  = 0x20000
)

// GetFziwdgstop IWDG Stop mode freeze option configuration bit
func (r *RegisterOptsrprgType) GetFziwdgstop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldFziwdgstopMask) != 0
}

// SetFziwdgstop IWDG Stop mode freeze option configuration bit
func (r *RegisterOptsrprgType) SetFziwdgstop(value bool) {
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
func (r *RegisterOptsrprgType) GetFziwdgsdby() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldFziwdgsdbyMask) != 0
}

// SetFziwdgsdby IWDG Standby mode freeze option configuration bit
func (r *RegisterOptsrprgType) SetFziwdgsdby(value bool) {
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
func (r *RegisterOptsrprgType) GetStramsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldStramsizeMask) >> RegisterOptsrprgFieldStramsizeShift)
}

// SetStramsize DTCM size select option configuration bits
func (r *RegisterOptsrprgType) SetStramsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldStramsizeMask)|(uint32(value)<<RegisterOptsrprgFieldStramsizeShift))
}

const (
	RegisterOptsrprgFieldSecurityShift = 21
	RegisterOptsrprgFieldSecurityMask  = 0x200000
)

// GetSecurity Security option configuration bit
func (r *RegisterOptsrprgType) GetSecurity() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldSecurityMask) != 0
}

// SetSecurity Security option configuration bit
func (r *RegisterOptsrprgType) SetSecurity(value bool) {
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
func (r *RegisterOptsrprgType) GetRss1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldRss1Mask) != 0
}

// SetRss1 User option configuration bit 1
func (r *RegisterOptsrprgType) SetRss1(value bool) {
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
func (r *RegisterOptsrprgType) GetRss2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldRss2Mask) != 0
}

// SetRss2 User option configuration bit 2
func (r *RegisterOptsrprgType) SetRss2(value bool) {
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
func (r *RegisterOptsrprgType) GetIohslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldIohslvMask) != 0
}

// SetIohslv I/O high-speed at low-voltage (PRODUCT_BELOW_25V)
func (r *RegisterOptsrprgType) SetIohslv(value bool) {
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
func (r *RegisterOptsrprgType) GetSwapbankopt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrprgFieldSwapbankoptMask) != 0
}

// SetSwapbankopt Bank swapping option configuration bit
func (r *RegisterOptsrprgType) SetSwapbankopt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrprgFieldSwapbankoptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrprgFieldSwapbankoptMask)
	}
}

// RegisterOptccrType FLASH option clear control register
type RegisterOptccrType uint32

func (r *RegisterOptccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOptccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOptccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOptccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOptccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOptccrFieldClroptchangeerrShift = 30
	RegisterOptccrFieldClroptchangeerrMask  = 0x40000000
)

// GetClroptchangeerr OPTCHANGEERR reset bit
func (r *RegisterOptccrType) GetClroptchangeerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptccrFieldClroptchangeerrMask) != 0
}

// SetClroptchangeerr OPTCHANGEERR reset bit
func (r *RegisterOptccrType) SetClroptchangeerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptccrFieldClroptchangeerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptccrFieldClroptchangeerrMask)
	}
}

// RegisterPrarcurType FLASH protection address for bank n
type RegisterPrarcurType uint32

func (r *RegisterPrarcurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPrarcurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPrarcurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPrarcurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPrarcurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPrarcurFieldProtareastartShift = 0
	RegisterPrarcurFieldProtareastartMask  = 0xfff
)

// GetProtareastart Bank n lowest PCROP protected address
func (r *RegisterPrarcurType) GetProtareastart() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarcurFieldProtareastartMask) >> RegisterPrarcurFieldProtareastartShift)
}

// SetProtareastart Bank n lowest PCROP protected address
func (r *RegisterPrarcurType) SetProtareastart(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarcurFieldProtareastartMask)|(uint32(value)<<RegisterPrarcurFieldProtareastartShift))
}

const (
	RegisterPrarcurFieldProtareaendShift = 16
	RegisterPrarcurFieldProtareaendMask  = 0xfff0000
)

// GetProtareaend Bank n highest PCROP protected address
func (r *RegisterPrarcurType) GetProtareaend() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarcurFieldProtareaendMask) >> RegisterPrarcurFieldProtareaendShift)
}

// SetProtareaend Bank n highest PCROP protected address
func (r *RegisterPrarcurType) SetProtareaend(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarcurFieldProtareaendMask)|(uint32(value)<<RegisterPrarcurFieldProtareaendShift))
}

const (
	RegisterPrarcurFieldDmepShift = 31
	RegisterPrarcurFieldDmepMask  = 0x80000000
)

// GetDmep Bank n PCROP protected erase enable option status bit
func (r *RegisterPrarcurType) GetDmep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrarcurFieldDmepMask) != 0
}

// SetDmep Bank n PCROP protected erase enable option status bit
func (r *RegisterPrarcurType) SetDmep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrarcurFieldDmepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrarcurFieldDmepMask)
	}
}

// RegisterPrarprg2Type FLASH protection address for bank 2
type RegisterPrarprg2Type uint32

func (r *RegisterPrarprg2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPrarprg2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPrarprg2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPrarprg2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPrarprg2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPrarprg2FieldProtareastart2Shift = 0
	RegisterPrarprg2FieldProtareastart2Mask  = 0xfff
)

// GetProtareastart2 Bank 2 lowest PCROP protected address configuration
func (r *RegisterPrarprg2Type) GetProtareastart2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarprg2FieldProtareastart2Mask) >> RegisterPrarprg2FieldProtareastart2Shift)
}

// SetProtareastart2 Bank 2 lowest PCROP protected address configuration
func (r *RegisterPrarprg2Type) SetProtareastart2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarprg2FieldProtareastart2Mask)|(uint32(value)<<RegisterPrarprg2FieldProtareastart2Shift))
}

const (
	RegisterPrarprg2FieldProtareaend2Shift = 16
	RegisterPrarprg2FieldProtareaend2Mask  = 0xfff0000
)

// GetProtareaend2 Bank 2 highest PCROP protected address configuration
func (r *RegisterPrarprg2Type) GetProtareaend2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarprg2FieldProtareaend2Mask) >> RegisterPrarprg2FieldProtareaend2Shift)
}

// SetProtareaend2 Bank 2 highest PCROP protected address configuration
func (r *RegisterPrarprg2Type) SetProtareaend2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarprg2FieldProtareaend2Mask)|(uint32(value)<<RegisterPrarprg2FieldProtareaend2Shift))
}

const (
	RegisterPrarprg2FieldDmep2Shift = 31
	RegisterPrarprg2FieldDmep2Mask  = 0x80000000
)

// GetDmep2 Bank 2 PCROP protected erase enable option configuration bit
func (r *RegisterPrarprg2Type) GetDmep2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrarprg2FieldDmep2Mask) != 0
}

// SetDmep2 Bank 2 PCROP protected erase enable option configuration bit
func (r *RegisterPrarprg2Type) SetDmep2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrarprg2FieldDmep2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrarprg2FieldDmep2Mask)
	}
}

// RegisterPrarprgType FLASH protection address for bank n
type RegisterPrarprgType uint32

func (r *RegisterPrarprgType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPrarprgType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPrarprgType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPrarprgType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPrarprgType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPrarprgFieldProtareastartShift = 0
	RegisterPrarprgFieldProtareastartMask  = 0xfff
)

// GetProtareastart Bank n lowest PCROP protected address configuration
func (r *RegisterPrarprgType) GetProtareastart() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarprgFieldProtareastartMask) >> RegisterPrarprgFieldProtareastartShift)
}

// SetProtareastart Bank n lowest PCROP protected address configuration
func (r *RegisterPrarprgType) SetProtareastart(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarprgFieldProtareastartMask)|(uint32(value)<<RegisterPrarprgFieldProtareastartShift))
}

const (
	RegisterPrarprgFieldProtareaendShift = 16
	RegisterPrarprgFieldProtareaendMask  = 0xfff0000
)

// GetProtareaend Bank n highest PCROP protected address configuration
func (r *RegisterPrarprgType) GetProtareaend() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarprgFieldProtareaendMask) >> RegisterPrarprgFieldProtareaendShift)
}

// SetProtareaend Bank n highest PCROP protected address configuration
func (r *RegisterPrarprgType) SetProtareaend(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarprgFieldProtareaendMask)|(uint32(value)<<RegisterPrarprgFieldProtareaendShift))
}

const (
	RegisterPrarprgFieldDmepShift = 31
	RegisterPrarprgFieldDmepMask  = 0x80000000
)

// GetDmep Bank n PCROP protected erase enable option configuration bit
func (r *RegisterPrarprgType) GetDmep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrarprgFieldDmepMask) != 0
}

// SetDmep Bank n PCROP protected erase enable option configuration bit
func (r *RegisterPrarprgType) SetDmep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrarprgFieldDmepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrarprgFieldDmepMask)
	}
}

// RegisterScarcurType FLASH secure address for bank n
type RegisterScarcurType uint32

func (r *RegisterScarcurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterScarcurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterScarcurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterScarcurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterScarcurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterScarcurFieldSecareastartShift = 0
	RegisterScarcurFieldSecareastartMask  = 0xfff
)

// GetSecareastart Bank n lowest secure protected address
func (r *RegisterScarcurType) GetSecareastart() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarcurFieldSecareastartMask) >> RegisterScarcurFieldSecareastartShift)
}

// SetSecareastart Bank n lowest secure protected address
func (r *RegisterScarcurType) SetSecareastart(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarcurFieldSecareastartMask)|(uint32(value)<<RegisterScarcurFieldSecareastartShift))
}

const (
	RegisterScarcurFieldSecareaendShift = 16
	RegisterScarcurFieldSecareaendMask  = 0xfff0000
)

// GetSecareaend Bank n highest secure protected address
func (r *RegisterScarcurType) GetSecareaend() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarcurFieldSecareaendMask) >> RegisterScarcurFieldSecareaendShift)
}

// SetSecareaend Bank n highest secure protected address
func (r *RegisterScarcurType) SetSecareaend(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarcurFieldSecareaendMask)|(uint32(value)<<RegisterScarcurFieldSecareaendShift))
}

const (
	RegisterScarcurFieldDmesShift = 31
	RegisterScarcurFieldDmesMask  = 0x80000000
)

// GetDmes Bank n secure protected erase enable option status bit
func (r *RegisterScarcurType) GetDmes() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScarcurFieldDmesMask) != 0
}

// SetDmes Bank n secure protected erase enable option status bit
func (r *RegisterScarcurType) SetDmes(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScarcurFieldDmesMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScarcurFieldDmesMask)
	}
}

// RegisterScarprgType FLASH secure address for bank n
type RegisterScarprgType uint32

func (r *RegisterScarprgType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterScarprgType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterScarprgType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterScarprgType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterScarprgType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterScarprgFieldSecareastartShift = 0
	RegisterScarprgFieldSecareastartMask  = 0xfff
)

// GetSecareastart Bank n lowest secure protected address configuration
func (r *RegisterScarprgType) GetSecareastart() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarprgFieldSecareastartMask) >> RegisterScarprgFieldSecareastartShift)
}

// SetSecareastart Bank n lowest secure protected address configuration
func (r *RegisterScarprgType) SetSecareastart(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarprgFieldSecareastartMask)|(uint32(value)<<RegisterScarprgFieldSecareastartShift))
}

const (
	RegisterScarprgFieldSecareaendShift = 16
	RegisterScarprgFieldSecareaendMask  = 0xfff0000
)

// GetSecareaend Bank n highest secure protected address configuration
func (r *RegisterScarprgType) GetSecareaend() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarprgFieldSecareaendMask) >> RegisterScarprgFieldSecareaendShift)
}

// SetSecareaend Bank n highest secure protected address configuration
func (r *RegisterScarprgType) SetSecareaend(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarprgFieldSecareaendMask)|(uint32(value)<<RegisterScarprgFieldSecareaendShift))
}

const (
	RegisterScarprgFieldDmesShift = 31
	RegisterScarprgFieldDmesMask  = 0x80000000
)

// GetDmes Bank n secure protected erase enable option configuration bit
func (r *RegisterScarprgType) GetDmes() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScarprgFieldDmesMask) != 0
}

// SetDmes Bank n secure protected erase enable option configuration bit
func (r *RegisterScarprgType) SetDmes(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScarprgFieldDmesMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScarprgFieldDmesMask)
	}
}

// RegisterWpsncur1rType FLASH write sector protection for bank n
type RegisterWpsncur1rType uint32

func (r *RegisterWpsncur1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterWpsncur1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterWpsncur1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterWpsncur1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterWpsncur1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterWpsncur1rFieldWrpsnShift = 0
	RegisterWpsncur1rFieldWrpsnMask  = 0xff
)

// GetWrpsn Bank n sector write protection option status byte
func (r *RegisterWpsncur1rType) GetWrpsn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWpsncur1rFieldWrpsnMask) >> RegisterWpsncur1rFieldWrpsnShift)
}

// SetWrpsn Bank n sector write protection option status byte
func (r *RegisterWpsncur1rType) SetWrpsn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWpsncur1rFieldWrpsnMask)|(uint32(value)<<RegisterWpsncur1rFieldWrpsnShift))
}

// RegisterWpsnprg1rType FLASH write sector protection for bank n
type RegisterWpsnprg1rType uint32

func (r *RegisterWpsnprg1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterWpsnprg1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterWpsnprg1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterWpsnprg1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterWpsnprg1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterWpsnprg1rFieldWrpsnShift = 0
	RegisterWpsnprg1rFieldWrpsnMask  = 0xff
)

// GetWrpsn Bank n sector write protection configuration byte
func (r *RegisterWpsnprg1rType) GetWrpsn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWpsnprg1rFieldWrpsnMask) >> RegisterWpsnprg1rFieldWrpsnShift)
}

// SetWrpsn Bank n sector write protection configuration byte
func (r *RegisterWpsnprg1rType) SetWrpsn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWpsnprg1rFieldWrpsnMask)|(uint32(value)<<RegisterWpsnprg1rFieldWrpsnShift))
}

// RegisterBootcurrType FLASH register with boot address
type RegisterBootcurrType uint32

func (r *RegisterBootcurrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBootcurrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBootcurrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBootcurrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBootcurrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBootcurrFieldBootadd0Shift = 0
	RegisterBootcurrFieldBootadd0Mask  = 0xffff
)

// GetBootadd0 Boot address 0
func (r *RegisterBootcurrType) GetBootadd0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBootcurrFieldBootadd0Mask) >> RegisterBootcurrFieldBootadd0Shift)
}

// SetBootadd0 Boot address 0
func (r *RegisterBootcurrType) SetBootadd0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBootcurrFieldBootadd0Mask)|(uint32(value)<<RegisterBootcurrFieldBootadd0Shift))
}

const (
	RegisterBootcurrFieldBootadd1Shift = 16
	RegisterBootcurrFieldBootadd1Mask  = 0xffff0000
)

// GetBootadd1 Boot address 1
func (r *RegisterBootcurrType) GetBootadd1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBootcurrFieldBootadd1Mask) >> RegisterBootcurrFieldBootadd1Shift)
}

// SetBootadd1 Boot address 1
func (r *RegisterBootcurrType) SetBootadd1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBootcurrFieldBootadd1Mask)|(uint32(value)<<RegisterBootcurrFieldBootadd1Shift))
}

// RegisterBootprgrType FLASH register with boot address
type RegisterBootprgrType uint32

func (r *RegisterBootprgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBootprgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBootprgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBootprgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBootprgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBootprgrFieldBootadd0Shift = 0
	RegisterBootprgrFieldBootadd0Mask  = 0xffff
)

// GetBootadd0 Boot address 0
func (r *RegisterBootprgrType) GetBootadd0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBootprgrFieldBootadd0Mask) >> RegisterBootprgrFieldBootadd0Shift)
}

// SetBootadd0 Boot address 0
func (r *RegisterBootprgrType) SetBootadd0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBootprgrFieldBootadd0Mask)|(uint32(value)<<RegisterBootprgrFieldBootadd0Shift))
}

const (
	RegisterBootprgrFieldBootadd1Shift = 16
	RegisterBootprgrFieldBootadd1Mask  = 0xffff0000
)

// GetBootadd1 Boot address 1
func (r *RegisterBootprgrType) GetBootadd1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBootprgrFieldBootadd1Mask) >> RegisterBootprgrFieldBootadd1Shift)
}

// SetBootadd1 Boot address 1
func (r *RegisterBootprgrType) SetBootadd1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBootprgrFieldBootadd1Mask)|(uint32(value)<<RegisterBootprgrFieldBootadd1Shift))
}

// RegisterCrccrType FLASH CRC control register for bank n
type RegisterCrccrType uint32

func (r *RegisterCrccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrccrFieldCrcsectShift = 0
	RegisterCrccrFieldCrcsectMask  = 0x7
)

// GetCrcsect Bank n CRC sector number
func (r *RegisterCrccrType) GetCrcsect() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldCrcsectMask) >> RegisterCrccrFieldCrcsectShift)
}

// SetCrcsect Bank n CRC sector number
func (r *RegisterCrccrType) SetCrcsect(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrccrFieldCrcsectMask)|(uint32(value)<<RegisterCrccrFieldCrcsectShift))
}

const (
	RegisterCrccrFieldAllbankShift = 7
	RegisterCrccrFieldAllbankMask  = 0x80
)

// GetAllbank Bank n CRC select bit
func (r *RegisterCrccrType) GetAllbank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldAllbankMask) != 0
}

// SetAllbank Bank n CRC select bit
func (r *RegisterCrccrType) SetAllbank(value bool) {
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
func (r *RegisterCrccrType) GetCrcbysect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldCrcbysectMask) != 0
}

// SetCrcbysect Bank n CRC sector mode select bit
func (r *RegisterCrccrType) SetCrcbysect(value bool) {
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
func (r *RegisterCrccrType) GetAddsect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldAddsectMask) != 0
}

// SetAddsect Bank n CRC sector select bit
func (r *RegisterCrccrType) SetAddsect(value bool) {
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
func (r *RegisterCrccrType) GetCleansect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldCleansectMask) != 0
}

// SetCleansect Bank n CRC sector list clear bit
func (r *RegisterCrccrType) SetCleansect(value bool) {
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
func (r *RegisterCrccrType) GetStartcrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldStartcrcMask) != 0
}

// SetStartcrc Bank n CRC start bit
func (r *RegisterCrccrType) SetStartcrc(value bool) {
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
func (r *RegisterCrccrType) GetCleancrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldCleancrcMask) != 0
}

// SetCleancrc Bank n CRC clear bit
func (r *RegisterCrccrType) SetCleancrc(value bool) {
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
func (r *RegisterCrccrType) GetCrcburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrccrFieldCrcburstMask) >> RegisterCrccrFieldCrcburstShift)
}

// SetCrcburst Bank n CRC burst size
func (r *RegisterCrccrType) SetCrcburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrccrFieldCrcburstMask)|(uint32(value)<<RegisterCrccrFieldCrcburstShift))
}

// RegisterCrcsadd1rType FLASH CRC start address register for bank n
type RegisterCrcsadd1rType uint32

func (r *RegisterCrcsadd1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrcsadd1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrcsadd1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrcsadd1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrcsadd1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrcsadd1rFieldCrcstartaddrShift = 0
	RegisterCrcsadd1rFieldCrcstartaddrMask  = 0xffffffff
)

// GetCrcstartaddr CRC start address on bank n
func (r *RegisterCrcsadd1rType) GetCrcstartaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrcsadd1rFieldCrcstartaddrMask) >> RegisterCrcsadd1rFieldCrcstartaddrShift)
}

// SetCrcstartaddr CRC start address on bank n
func (r *RegisterCrcsadd1rType) SetCrcstartaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrcsadd1rFieldCrcstartaddrMask)|(uint32(value)<<RegisterCrcsadd1rFieldCrcstartaddrShift))
}

// RegisterCrceadd1rType FLASH CRC end address register for bank n
type RegisterCrceadd1rType uint32

func (r *RegisterCrceadd1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrceadd1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrceadd1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrceadd1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrceadd1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrceadd1rFieldCrcendaddrShift = 0
	RegisterCrceadd1rFieldCrcendaddrMask  = 0xffffffff
)

// GetCrcendaddr CRC end address on bank n
func (r *RegisterCrceadd1rType) GetCrcendaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrceadd1rFieldCrcendaddrMask) >> RegisterCrceadd1rFieldCrcendaddrShift)
}

// SetCrcendaddr CRC end address on bank n
func (r *RegisterCrceadd1rType) SetCrcendaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrceadd1rFieldCrcendaddrMask)|(uint32(value)<<RegisterCrceadd1rFieldCrcendaddrShift))
}

// RegisterCrcdatarType FLASH CRC data register
type RegisterCrcdatarType uint32

func (r *RegisterCrcdatarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrcdatarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrcdatarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrcdatarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrcdatarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrcdatarFieldCrcdataShift = 0
	RegisterCrcdatarFieldCrcdataMask  = 0xffffffff
)

// GetCrcdata CRC result
func (r *RegisterCrcdatarType) GetCrcdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrcdatarFieldCrcdataMask) >> RegisterCrcdatarFieldCrcdataShift)
}

// SetCrcdata CRC result
func (r *RegisterCrcdatarType) SetCrcdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrcdatarFieldCrcdataMask)|(uint32(value)<<RegisterCrcdatarFieldCrcdataShift))
}

// RegisterEccfa1rType FLASH ECC fail address for bank n
type RegisterEccfa1rType uint32

func (r *RegisterEccfa1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterEccfa1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterEccfa1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterEccfa1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterEccfa1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterEccfa1rFieldFaileccaddrShift = 0
	RegisterEccfa1rFieldFaileccaddrMask  = 0x7fff
)

// GetFaileccaddr Bank n ECC error address
func (r *RegisterEccfa1rType) GetFaileccaddr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterEccfa1rFieldFaileccaddrMask) >> RegisterEccfa1rFieldFaileccaddrShift)
}

// SetFaileccaddr Bank n ECC error address
func (r *RegisterEccfa1rType) SetFaileccaddr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEccfa1rFieldFaileccaddrMask)|(uint32(value)<<RegisterEccfa1rFieldFaileccaddrShift))
}
