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
	Acr       registerAcrType
	Keyr1     registerKeyr1Type
	Optkeyr   registerOptkeyrType
	Cr1       registerCr1Type
	Sr1       registerSr1Type
	Ccr1      registerCcr1Type
	Optcr     registerOptcrType
	Optsrcur  registerOptsrcurType
	Optsrprg  registerOptsrprgType
	Optccr    registerOptccrType
	Prarcur1  registerPrarcur1Type
	Prarprg2  registerPrarprg2Type
	Prarprg1  registerPrarprg1Type
	Scarcur1  registerScarcur1Type
	Scarprg1  registerScarprg1Type
	Wpsncur1r registerWpsncur1rType
	Wpsnprg1r registerWpsnprg1rType
	Bootcurr  registerBootcurrType
	Bootprgr  registerBootprgrType
	_         [8]byte
	Crccr1    registerCrccr1Type
	Crcsadd1r registerCrcsadd1rType
	Crceadd1r registerCrceadd1rType
	Crcdatar  registerCrcdatarType
	Eccfa1r   registerEccfa1rType
	_         [156]byte
	Acr       registerAcrType
	Keyr2     registerKeyr2Type
	Optkeyr   registerOptkeyrType
	Cr2       registerCr2Type
	Sr2       registerSr2Type
	Ccr2      registerCcr2Type
	Optcr     registerOptcrType
	Optsrcur  registerOptsrcurType
	Optsrprg  registerOptsrprgType
	Optccr    registerOptccrType
	Prarcur2  registerPrarcur2Type
	_         [4]byte
	Scarcur2  registerScarcur2Type
	Scarprg2  registerScarprg2Type
	Wpsncur2r registerWpsncur2rType
	Wpsnprg2r registerWpsnprg2rType
	_         [16]byte
	Crccr2    registerCrccr2Type
	Crcsadd2r registerCrcsadd2rType
	Crceadd2r registerCrceadd2rType
	_         [4]byte
	Eccfa2r   registerEccfa2rType
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

// registerKeyr1Type FLASH key register for bank 1
type registerKeyr1Type uint32

const (
	RegisterKeyr1FieldKeyr1Shift = 0
	RegisterKeyr1FieldKeyr1Mask  = 0xffffffff
)

// GetKeyr1 Bank 1 access configuration unlock key
func (r *registerKeyr1Type) GetKeyr1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterKeyr1FieldKeyr1Mask) >> RegisterKeyr1FieldKeyr1Shift)
}

// SetKeyr1 Bank 1 access configuration unlock key
func (r *registerKeyr1Type) SetKeyr1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterKeyr1FieldKeyr1Mask)|(uint32(value)<<RegisterKeyr1FieldKeyr1Shift))
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

// registerCr1Type FLASH control register for bank 1
type registerCr1Type uint32

const (
	RegisterCr1FieldLock1Shift = 0
	RegisterCr1FieldLock1Mask  = 0x1
)

// GetLock1 Bank 1 configuration lock bit
func (r *registerCr1Type) GetLock1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldLock1Mask) != 0
}

// SetLock1 Bank 1 configuration lock bit
func (r *registerCr1Type) SetLock1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldLock1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldLock1Mask)
	}
}

const (
	RegisterCr1FieldPg1Shift = 1
	RegisterCr1FieldPg1Mask  = 0x2
)

// GetPg1 Bank 1 program enable bit
func (r *registerCr1Type) GetPg1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPg1Mask) != 0
}

// SetPg1 Bank 1 program enable bit
func (r *registerCr1Type) SetPg1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldPg1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldPg1Mask)
	}
}

const (
	RegisterCr1FieldSer1Shift = 2
	RegisterCr1FieldSer1Mask  = 0x4
)

// GetSer1 Bank 1 sector erase request
func (r *registerCr1Type) GetSer1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldSer1Mask) != 0
}

// SetSer1 Bank 1 sector erase request
func (r *registerCr1Type) SetSer1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldSer1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldSer1Mask)
	}
}

const (
	RegisterCr1FieldBer1Shift = 3
	RegisterCr1FieldBer1Mask  = 0x8
)

// GetBer1 Bank 1 erase request
func (r *registerCr1Type) GetBer1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldBer1Mask) != 0
}

// SetBer1 Bank 1 erase request
func (r *registerCr1Type) SetBer1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldBer1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldBer1Mask)
	}
}

const (
	RegisterCr1FieldPsize1Shift = 4
	RegisterCr1FieldPsize1Mask  = 0x30
)

// GetPsize1 Bank 1 program size
func (r *registerCr1Type) GetPsize1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPsize1Mask) >> RegisterCr1FieldPsize1Shift)
}

// SetPsize1 Bank 1 program size
func (r *registerCr1Type) SetPsize1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldPsize1Mask)|(uint32(value)<<RegisterCr1FieldPsize1Shift))
}

const (
	RegisterCr1FieldFw1Shift = 6
	RegisterCr1FieldFw1Mask  = 0x40
)

// GetFw1 Bank 1 write forcing control bit
func (r *registerCr1Type) GetFw1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldFw1Mask) != 0
}

// SetFw1 Bank 1 write forcing control bit
func (r *registerCr1Type) SetFw1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldFw1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldFw1Mask)
	}
}

const (
	RegisterCr1FieldStart1Shift = 7
	RegisterCr1FieldStart1Mask  = 0x80
)

// GetStart1 Bank 1 bank or sector erase start control bit
func (r *registerCr1Type) GetStart1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldStart1Mask) != 0
}

// SetStart1 Bank 1 bank or sector erase start control bit
func (r *registerCr1Type) SetStart1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldStart1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldStart1Mask)
	}
}

const (
	RegisterCr1FieldSnb1Shift = 8
	RegisterCr1FieldSnb1Mask  = 0x700
)

// GetSnb1 Bank 1 sector erase selection number
func (r *registerCr1Type) GetSnb1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldSnb1Mask) >> RegisterCr1FieldSnb1Shift)
}

// SetSnb1 Bank 1 sector erase selection number
func (r *registerCr1Type) SetSnb1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldSnb1Mask)|(uint32(value)<<RegisterCr1FieldSnb1Shift))
}

const (
	RegisterCr1FieldCrcenShift = 15
	RegisterCr1FieldCrcenMask  = 0x8000
)

// GetCrcen Bank 1 CRC control bit
func (r *registerCr1Type) GetCrcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCrcenMask) != 0
}

// SetCrcen Bank 1 CRC control bit
func (r *registerCr1Type) SetCrcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldCrcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldCrcenMask)
	}
}

const (
	RegisterCr1FieldEopie1Shift = 16
	RegisterCr1FieldEopie1Mask  = 0x10000
)

// GetEopie1 Bank 1 end-of-program interrupt control bit
func (r *registerCr1Type) GetEopie1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldEopie1Mask) != 0
}

// SetEopie1 Bank 1 end-of-program interrupt control bit
func (r *registerCr1Type) SetEopie1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldEopie1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldEopie1Mask)
	}
}

const (
	RegisterCr1FieldWrperrie1Shift = 17
	RegisterCr1FieldWrperrie1Mask  = 0x20000
)

// GetWrperrie1 Bank 1 write protection error interrupt enable bit
func (r *registerCr1Type) GetWrperrie1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldWrperrie1Mask) != 0
}

// SetWrperrie1 Bank 1 write protection error interrupt enable bit
func (r *registerCr1Type) SetWrperrie1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldWrperrie1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldWrperrie1Mask)
	}
}

const (
	RegisterCr1FieldPgserrie1Shift = 18
	RegisterCr1FieldPgserrie1Mask  = 0x40000
)

// GetPgserrie1 Bank 1 programming sequence error interrupt enable bit
func (r *registerCr1Type) GetPgserrie1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPgserrie1Mask) != 0
}

// SetPgserrie1 Bank 1 programming sequence error interrupt enable bit
func (r *registerCr1Type) SetPgserrie1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldPgserrie1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldPgserrie1Mask)
	}
}

const (
	RegisterCr1FieldStrberrie1Shift = 19
	RegisterCr1FieldStrberrie1Mask  = 0x80000
)

// GetStrberrie1 Bank 1 strobe error interrupt enable bit
func (r *registerCr1Type) GetStrberrie1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldStrberrie1Mask) != 0
}

// SetStrberrie1 Bank 1 strobe error interrupt enable bit
func (r *registerCr1Type) SetStrberrie1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldStrberrie1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldStrberrie1Mask)
	}
}

const (
	RegisterCr1FieldIncerrie1Shift = 21
	RegisterCr1FieldIncerrie1Mask  = 0x200000
)

// GetIncerrie1 Bank 1 inconsistency error interrupt enable bit
func (r *registerCr1Type) GetIncerrie1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldIncerrie1Mask) != 0
}

// SetIncerrie1 Bank 1 inconsistency error interrupt enable bit
func (r *registerCr1Type) SetIncerrie1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldIncerrie1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldIncerrie1Mask)
	}
}

const (
	RegisterCr1FieldOperrie1Shift = 22
	RegisterCr1FieldOperrie1Mask  = 0x400000
)

// GetOperrie1 Bank 1 write/erase error interrupt enable bit
func (r *registerCr1Type) GetOperrie1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldOperrie1Mask) != 0
}

// SetOperrie1 Bank 1 write/erase error interrupt enable bit
func (r *registerCr1Type) SetOperrie1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldOperrie1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldOperrie1Mask)
	}
}

const (
	RegisterCr1FieldRdperrie1Shift = 23
	RegisterCr1FieldRdperrie1Mask  = 0x800000
)

// GetRdperrie1 Bank 1 read protection error interrupt enable bit
func (r *registerCr1Type) GetRdperrie1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldRdperrie1Mask) != 0
}

// SetRdperrie1 Bank 1 read protection error interrupt enable bit
func (r *registerCr1Type) SetRdperrie1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldRdperrie1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldRdperrie1Mask)
	}
}

const (
	RegisterCr1FieldRdserrie1Shift = 24
	RegisterCr1FieldRdserrie1Mask  = 0x1000000
)

// GetRdserrie1 Bank 1 secure error interrupt enable bit
func (r *registerCr1Type) GetRdserrie1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldRdserrie1Mask) != 0
}

// SetRdserrie1 Bank 1 secure error interrupt enable bit
func (r *registerCr1Type) SetRdserrie1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldRdserrie1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldRdserrie1Mask)
	}
}

const (
	RegisterCr1FieldSneccerrie1Shift = 25
	RegisterCr1FieldSneccerrie1Mask  = 0x2000000
)

// GetSneccerrie1 Bank 1 ECC single correction error interrupt enable bit
func (r *registerCr1Type) GetSneccerrie1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldSneccerrie1Mask) != 0
}

// SetSneccerrie1 Bank 1 ECC single correction error interrupt enable bit
func (r *registerCr1Type) SetSneccerrie1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldSneccerrie1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldSneccerrie1Mask)
	}
}

const (
	RegisterCr1FieldDbeccerrie1Shift = 26
	RegisterCr1FieldDbeccerrie1Mask  = 0x4000000
)

// GetDbeccerrie1 Bank 1 ECC double detection error interrupt enable bit
func (r *registerCr1Type) GetDbeccerrie1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDbeccerrie1Mask) != 0
}

// SetDbeccerrie1 Bank 1 ECC double detection error interrupt enable bit
func (r *registerCr1Type) SetDbeccerrie1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDbeccerrie1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDbeccerrie1Mask)
	}
}

const (
	RegisterCr1FieldCrcendie1Shift = 27
	RegisterCr1FieldCrcendie1Mask  = 0x8000000
)

// GetCrcendie1 Bank 1 end of CRC calculation interrupt enable bit
func (r *registerCr1Type) GetCrcendie1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCrcendie1Mask) != 0
}

// SetCrcendie1 Bank 1 end of CRC calculation interrupt enable bit
func (r *registerCr1Type) SetCrcendie1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldCrcendie1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldCrcendie1Mask)
	}
}

// registerSr1Type FLASH status register for bank 1
type registerSr1Type uint32

const (
	RegisterSr1FieldBsy1Shift = 0
	RegisterSr1FieldBsy1Mask  = 0x1
)

// GetBsy1 Bank 1 ongoing program flag
func (r *registerSr1Type) GetBsy1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldBsy1Mask) != 0
}

// SetBsy1 Bank 1 ongoing program flag
func (r *registerSr1Type) SetBsy1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldBsy1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldBsy1Mask)
	}
}

const (
	RegisterSr1FieldWbne1Shift = 1
	RegisterSr1FieldWbne1Mask  = 0x2
)

// GetWbne1 Bank 1 write buffer not empty flag
func (r *registerSr1Type) GetWbne1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldWbne1Mask) != 0
}

// SetWbne1 Bank 1 write buffer not empty flag
func (r *registerSr1Type) SetWbne1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldWbne1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldWbne1Mask)
	}
}

const (
	RegisterSr1FieldQw1Shift = 2
	RegisterSr1FieldQw1Mask  = 0x4
)

// GetQw1 Bank 1 wait queue flag
func (r *registerSr1Type) GetQw1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldQw1Mask) != 0
}

// SetQw1 Bank 1 wait queue flag
func (r *registerSr1Type) SetQw1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldQw1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldQw1Mask)
	}
}

const (
	RegisterSr1FieldCrcbusy1Shift = 3
	RegisterSr1FieldCrcbusy1Mask  = 0x8
)

// GetCrcbusy1 Bank 1 CRC busy flag
func (r *registerSr1Type) GetCrcbusy1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldCrcbusy1Mask) != 0
}

// SetCrcbusy1 Bank 1 CRC busy flag
func (r *registerSr1Type) SetCrcbusy1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldCrcbusy1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldCrcbusy1Mask)
	}
}

const (
	RegisterSr1FieldEop1Shift = 16
	RegisterSr1FieldEop1Mask  = 0x10000
)

// GetEop1 Bank 1 end-of-program flag
func (r *registerSr1Type) GetEop1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldEop1Mask) != 0
}

// SetEop1 Bank 1 end-of-program flag
func (r *registerSr1Type) SetEop1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldEop1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldEop1Mask)
	}
}

const (
	RegisterSr1FieldWrperr1Shift = 17
	RegisterSr1FieldWrperr1Mask  = 0x20000
)

// GetWrperr1 Bank 1 write protection error flag
func (r *registerSr1Type) GetWrperr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldWrperr1Mask) != 0
}

// SetWrperr1 Bank 1 write protection error flag
func (r *registerSr1Type) SetWrperr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldWrperr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldWrperr1Mask)
	}
}

const (
	RegisterSr1FieldPgserr1Shift = 18
	RegisterSr1FieldPgserr1Mask  = 0x40000
)

// GetPgserr1 Bank 1 programming sequence error flag
func (r *registerSr1Type) GetPgserr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldPgserr1Mask) != 0
}

// SetPgserr1 Bank 1 programming sequence error flag
func (r *registerSr1Type) SetPgserr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldPgserr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldPgserr1Mask)
	}
}

const (
	RegisterSr1FieldStrberr1Shift = 19
	RegisterSr1FieldStrberr1Mask  = 0x80000
)

// GetStrberr1 Bank 1 strobe error flag
func (r *registerSr1Type) GetStrberr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldStrberr1Mask) != 0
}

// SetStrberr1 Bank 1 strobe error flag
func (r *registerSr1Type) SetStrberr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldStrberr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldStrberr1Mask)
	}
}

const (
	RegisterSr1FieldIncerr1Shift = 21
	RegisterSr1FieldIncerr1Mask  = 0x200000
)

// GetIncerr1 Bank 1 inconsistency error flag
func (r *registerSr1Type) GetIncerr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldIncerr1Mask) != 0
}

// SetIncerr1 Bank 1 inconsistency error flag
func (r *registerSr1Type) SetIncerr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldIncerr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldIncerr1Mask)
	}
}

const (
	RegisterSr1FieldOperr1Shift = 22
	RegisterSr1FieldOperr1Mask  = 0x400000
)

// GetOperr1 Bank 1 write/erase error flag
func (r *registerSr1Type) GetOperr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldOperr1Mask) != 0
}

// SetOperr1 Bank 1 write/erase error flag
func (r *registerSr1Type) SetOperr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldOperr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldOperr1Mask)
	}
}

const (
	RegisterSr1FieldRdperr1Shift = 23
	RegisterSr1FieldRdperr1Mask  = 0x800000
)

// GetRdperr1 Bank 1 read protection error flag
func (r *registerSr1Type) GetRdperr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldRdperr1Mask) != 0
}

// SetRdperr1 Bank 1 read protection error flag
func (r *registerSr1Type) SetRdperr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldRdperr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldRdperr1Mask)
	}
}

const (
	RegisterSr1FieldRdserr1Shift = 24
	RegisterSr1FieldRdserr1Mask  = 0x1000000
)

// GetRdserr1 Bank 1 secure error flag
func (r *registerSr1Type) GetRdserr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldRdserr1Mask) != 0
}

// SetRdserr1 Bank 1 secure error flag
func (r *registerSr1Type) SetRdserr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldRdserr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldRdserr1Mask)
	}
}

const (
	RegisterSr1FieldSneccerr11Shift = 25
	RegisterSr1FieldSneccerr11Mask  = 0x2000000
)

// GetSneccerr11 Bank 1 single correction error flag
func (r *registerSr1Type) GetSneccerr11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldSneccerr11Mask) != 0
}

// SetSneccerr11 Bank 1 single correction error flag
func (r *registerSr1Type) SetSneccerr11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldSneccerr11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldSneccerr11Mask)
	}
}

const (
	RegisterSr1FieldDbeccerr1Shift = 26
	RegisterSr1FieldDbeccerr1Mask  = 0x4000000
)

// GetDbeccerr1 Bank 1 ECC double detection error flag
func (r *registerSr1Type) GetDbeccerr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldDbeccerr1Mask) != 0
}

// SetDbeccerr1 Bank 1 ECC double detection error flag
func (r *registerSr1Type) SetDbeccerr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldDbeccerr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldDbeccerr1Mask)
	}
}

const (
	RegisterSr1FieldCrcend1Shift = 27
	RegisterSr1FieldCrcend1Mask  = 0x8000000
)

// GetCrcend1 Bank 1 CRC-complete flag
func (r *registerSr1Type) GetCrcend1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldCrcend1Mask) != 0
}

// SetCrcend1 Bank 1 CRC-complete flag
func (r *registerSr1Type) SetCrcend1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldCrcend1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldCrcend1Mask)
	}
}

// registerCcr1Type FLASH clear control register for bank 1
type registerCcr1Type uint32

const (
	RegisterCcr1FieldClreop1Shift = 16
	RegisterCcr1FieldClreop1Mask  = 0x10000
)

// GetClreop1 Bank 1 EOP1 flag clear bit
func (r *registerCcr1Type) GetClreop1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClreop1Mask) != 0
}

// SetClreop1 Bank 1 EOP1 flag clear bit
func (r *registerCcr1Type) SetClreop1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClreop1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClreop1Mask)
	}
}

const (
	RegisterCcr1FieldClrwrperr1Shift = 17
	RegisterCcr1FieldClrwrperr1Mask  = 0x20000
)

// GetClrwrperr1 Bank 1 WRPERR1 flag clear bit
func (r *registerCcr1Type) GetClrwrperr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClrwrperr1Mask) != 0
}

// SetClrwrperr1 Bank 1 WRPERR1 flag clear bit
func (r *registerCcr1Type) SetClrwrperr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClrwrperr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClrwrperr1Mask)
	}
}

const (
	RegisterCcr1FieldClrpgserr1Shift = 18
	RegisterCcr1FieldClrpgserr1Mask  = 0x40000
)

// GetClrpgserr1 Bank 1 PGSERR1 flag clear bi
func (r *registerCcr1Type) GetClrpgserr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClrpgserr1Mask) != 0
}

// SetClrpgserr1 Bank 1 PGSERR1 flag clear bi
func (r *registerCcr1Type) SetClrpgserr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClrpgserr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClrpgserr1Mask)
	}
}

const (
	RegisterCcr1FieldClrstrberr1Shift = 19
	RegisterCcr1FieldClrstrberr1Mask  = 0x80000
)

// GetClrstrberr1 Bank 1 STRBERR1 flag clear bit
func (r *registerCcr1Type) GetClrstrberr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClrstrberr1Mask) != 0
}

// SetClrstrberr1 Bank 1 STRBERR1 flag clear bit
func (r *registerCcr1Type) SetClrstrberr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClrstrberr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClrstrberr1Mask)
	}
}

const (
	RegisterCcr1FieldClrincerr1Shift = 21
	RegisterCcr1FieldClrincerr1Mask  = 0x200000
)

// GetClrincerr1 Bank 1 INCERR1 flag clear bit
func (r *registerCcr1Type) GetClrincerr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClrincerr1Mask) != 0
}

// SetClrincerr1 Bank 1 INCERR1 flag clear bit
func (r *registerCcr1Type) SetClrincerr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClrincerr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClrincerr1Mask)
	}
}

const (
	RegisterCcr1FieldClroperr1Shift = 22
	RegisterCcr1FieldClroperr1Mask  = 0x400000
)

// GetClroperr1 Bank 1 OPERR1 flag clear bit
func (r *registerCcr1Type) GetClroperr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClroperr1Mask) != 0
}

// SetClroperr1 Bank 1 OPERR1 flag clear bit
func (r *registerCcr1Type) SetClroperr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClroperr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClroperr1Mask)
	}
}

const (
	RegisterCcr1FieldClrrdperr1Shift = 23
	RegisterCcr1FieldClrrdperr1Mask  = 0x800000
)

// GetClrrdperr1 Bank 1 RDPERR1 flag clear bit
func (r *registerCcr1Type) GetClrrdperr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClrrdperr1Mask) != 0
}

// SetClrrdperr1 Bank 1 RDPERR1 flag clear bit
func (r *registerCcr1Type) SetClrrdperr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClrrdperr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClrrdperr1Mask)
	}
}

const (
	RegisterCcr1FieldClrrdserr1Shift = 24
	RegisterCcr1FieldClrrdserr1Mask  = 0x1000000
)

// GetClrrdserr1 Bank 1 RDSERR1 flag clear bit
func (r *registerCcr1Type) GetClrrdserr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClrrdserr1Mask) != 0
}

// SetClrrdserr1 Bank 1 RDSERR1 flag clear bit
func (r *registerCcr1Type) SetClrrdserr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClrrdserr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClrrdserr1Mask)
	}
}

const (
	RegisterCcr1FieldClrsneccerr1Shift = 25
	RegisterCcr1FieldClrsneccerr1Mask  = 0x2000000
)

// GetClrsneccerr1 Bank 1 SNECCERR1 flag clear bit
func (r *registerCcr1Type) GetClrsneccerr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClrsneccerr1Mask) != 0
}

// SetClrsneccerr1 Bank 1 SNECCERR1 flag clear bit
func (r *registerCcr1Type) SetClrsneccerr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClrsneccerr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClrsneccerr1Mask)
	}
}

const (
	RegisterCcr1FieldClrdbeccerr1Shift = 26
	RegisterCcr1FieldClrdbeccerr1Mask  = 0x4000000
)

// GetClrdbeccerr1 Bank 1 DBECCERR1 flag clear bit
func (r *registerCcr1Type) GetClrdbeccerr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClrdbeccerr1Mask) != 0
}

// SetClrdbeccerr1 Bank 1 DBECCERR1 flag clear bit
func (r *registerCcr1Type) SetClrdbeccerr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClrdbeccerr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClrdbeccerr1Mask)
	}
}

const (
	RegisterCcr1FieldClrcrcend1Shift = 27
	RegisterCcr1FieldClrcrcend1Mask  = 0x8000000
)

// GetClrcrcend1 Bank 1 CRCEND1 flag clear bit
func (r *registerCcr1Type) GetClrcrcend1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClrcrcend1Mask) != 0
}

// SetClrcrcend1 Bank 1 CRCEND1 flag clear bit
func (r *registerCcr1Type) SetClrcrcend1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClrcrcend1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClrcrcend1Mask)
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
	RegisterOptsrcurFieldNrststopd1Shift = 6
	RegisterOptsrcurFieldNrststopd1Mask  = 0x40
)

// GetNrststopd1 D1 DStop entry reset option status bit
func (r *registerOptsrcurType) GetNrststopd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldNrststopd1Mask) != 0
}

// SetNrststopd1 D1 DStop entry reset option status bit
func (r *registerOptsrcurType) SetNrststopd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldNrststopd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldNrststopd1Mask)
	}
}

const (
	RegisterOptsrcurFieldNrststbyd1Shift = 7
	RegisterOptsrcurFieldNrststbyd1Mask  = 0x80
)

// GetNrststbyd1 D1 DStandby entry reset option status bit
func (r *registerOptsrcurType) GetNrststbyd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldNrststbyd1Mask) != 0
}

// SetNrststbyd1 D1 DStandby entry reset option status bit
func (r *registerOptsrcurType) SetNrststbyd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldNrststbyd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldNrststbyd1Mask)
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

// registerPrarcur1Type FLASH protection address for bank 1
type registerPrarcur1Type uint32

const (
	RegisterPrarcur1FieldProtareastart1Shift = 0
	RegisterPrarcur1FieldProtareastart1Mask  = 0xfff
)

// GetProtareastart1 Bank 1 lowest PCROP protected address
func (r *registerPrarcur1Type) GetProtareastart1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarcur1FieldProtareastart1Mask) >> RegisterPrarcur1FieldProtareastart1Shift)
}

// SetProtareastart1 Bank 1 lowest PCROP protected address
func (r *registerPrarcur1Type) SetProtareastart1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarcur1FieldProtareastart1Mask)|(uint32(value)<<RegisterPrarcur1FieldProtareastart1Shift))
}

const (
	RegisterPrarcur1FieldProtareaend1Shift = 16
	RegisterPrarcur1FieldProtareaend1Mask  = 0xfff0000
)

// GetProtareaend1 Bank 1 highest PCROP protected address
func (r *registerPrarcur1Type) GetProtareaend1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarcur1FieldProtareaend1Mask) >> RegisterPrarcur1FieldProtareaend1Shift)
}

// SetProtareaend1 Bank 1 highest PCROP protected address
func (r *registerPrarcur1Type) SetProtareaend1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarcur1FieldProtareaend1Mask)|(uint32(value)<<RegisterPrarcur1FieldProtareaend1Shift))
}

const (
	RegisterPrarcur1FieldDmep1Shift = 31
	RegisterPrarcur1FieldDmep1Mask  = 0x80000000
)

// GetDmep1 Bank 1 PCROP protected erase enable option status bit
func (r *registerPrarcur1Type) GetDmep1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrarcur1FieldDmep1Mask) != 0
}

// SetDmep1 Bank 1 PCROP protected erase enable option status bit
func (r *registerPrarcur1Type) SetDmep1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrarcur1FieldDmep1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrarcur1FieldDmep1Mask)
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

// registerPrarprg1Type FLASH protection address for bank 1
type registerPrarprg1Type uint32

const (
	RegisterPrarprg1FieldProtareastart1Shift = 0
	RegisterPrarprg1FieldProtareastart1Mask  = 0xfff
)

// GetProtareastart1 Bank 1 lowest PCROP protected address configuration
func (r *registerPrarprg1Type) GetProtareastart1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarprg1FieldProtareastart1Mask) >> RegisterPrarprg1FieldProtareastart1Shift)
}

// SetProtareastart1 Bank 1 lowest PCROP protected address configuration
func (r *registerPrarprg1Type) SetProtareastart1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarprg1FieldProtareastart1Mask)|(uint32(value)<<RegisterPrarprg1FieldProtareastart1Shift))
}

const (
	RegisterPrarprg1FieldProtareaend1Shift = 16
	RegisterPrarprg1FieldProtareaend1Mask  = 0xfff0000
)

// GetProtareaend1 Bank 1 highest PCROP protected address configuration
func (r *registerPrarprg1Type) GetProtareaend1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarprg1FieldProtareaend1Mask) >> RegisterPrarprg1FieldProtareaend1Shift)
}

// SetProtareaend1 Bank 1 highest PCROP protected address configuration
func (r *registerPrarprg1Type) SetProtareaend1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarprg1FieldProtareaend1Mask)|(uint32(value)<<RegisterPrarprg1FieldProtareaend1Shift))
}

const (
	RegisterPrarprg1FieldDmep1Shift = 31
	RegisterPrarprg1FieldDmep1Mask  = 0x80000000
)

// GetDmep1 Bank 1 PCROP protected erase enable option configuration bit
func (r *registerPrarprg1Type) GetDmep1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrarprg1FieldDmep1Mask) != 0
}

// SetDmep1 Bank 1 PCROP protected erase enable option configuration bit
func (r *registerPrarprg1Type) SetDmep1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrarprg1FieldDmep1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrarprg1FieldDmep1Mask)
	}
}

// registerScarcur1Type FLASH secure address for bank 1
type registerScarcur1Type uint32

const (
	RegisterScarcur1FieldSecareastart1Shift = 0
	RegisterScarcur1FieldSecareastart1Mask  = 0xfff
)

// GetSecareastart1 Bank 1 lowest secure protected address
func (r *registerScarcur1Type) GetSecareastart1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarcur1FieldSecareastart1Mask) >> RegisterScarcur1FieldSecareastart1Shift)
}

// SetSecareastart1 Bank 1 lowest secure protected address
func (r *registerScarcur1Type) SetSecareastart1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarcur1FieldSecareastart1Mask)|(uint32(value)<<RegisterScarcur1FieldSecareastart1Shift))
}

const (
	RegisterScarcur1FieldSecareaend1Shift = 16
	RegisterScarcur1FieldSecareaend1Mask  = 0xfff0000
)

// GetSecareaend1 Bank 1 highest secure protected address
func (r *registerScarcur1Type) GetSecareaend1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarcur1FieldSecareaend1Mask) >> RegisterScarcur1FieldSecareaend1Shift)
}

// SetSecareaend1 Bank 1 highest secure protected address
func (r *registerScarcur1Type) SetSecareaend1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarcur1FieldSecareaend1Mask)|(uint32(value)<<RegisterScarcur1FieldSecareaend1Shift))
}

const (
	RegisterScarcur1FieldDmes1Shift = 31
	RegisterScarcur1FieldDmes1Mask  = 0x80000000
)

// GetDmes1 Bank 1 secure protected erase enable option status bit
func (r *registerScarcur1Type) GetDmes1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScarcur1FieldDmes1Mask) != 0
}

// SetDmes1 Bank 1 secure protected erase enable option status bit
func (r *registerScarcur1Type) SetDmes1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScarcur1FieldDmes1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScarcur1FieldDmes1Mask)
	}
}

// registerScarprg1Type FLASH secure address for bank 1
type registerScarprg1Type uint32

const (
	RegisterScarprg1FieldSecareastart1Shift = 0
	RegisterScarprg1FieldSecareastart1Mask  = 0xfff
)

// GetSecareastart1 Bank 1 lowest secure protected address configuration
func (r *registerScarprg1Type) GetSecareastart1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarprg1FieldSecareastart1Mask) >> RegisterScarprg1FieldSecareastart1Shift)
}

// SetSecareastart1 Bank 1 lowest secure protected address configuration
func (r *registerScarprg1Type) SetSecareastart1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarprg1FieldSecareastart1Mask)|(uint32(value)<<RegisterScarprg1FieldSecareastart1Shift))
}

const (
	RegisterScarprg1FieldSecareaend1Shift = 16
	RegisterScarprg1FieldSecareaend1Mask  = 0xfff0000
)

// GetSecareaend1 Bank 1 highest secure protected address configuration
func (r *registerScarprg1Type) GetSecareaend1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarprg1FieldSecareaend1Mask) >> RegisterScarprg1FieldSecareaend1Shift)
}

// SetSecareaend1 Bank 1 highest secure protected address configuration
func (r *registerScarprg1Type) SetSecareaend1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarprg1FieldSecareaend1Mask)|(uint32(value)<<RegisterScarprg1FieldSecareaend1Shift))
}

const (
	RegisterScarprg1FieldDmes1Shift = 31
	RegisterScarprg1FieldDmes1Mask  = 0x80000000
)

// GetDmes1 Bank 1 secure protected erase enable option configuration bit
func (r *registerScarprg1Type) GetDmes1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScarprg1FieldDmes1Mask) != 0
}

// SetDmes1 Bank 1 secure protected erase enable option configuration bit
func (r *registerScarprg1Type) SetDmes1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScarprg1FieldDmes1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScarprg1FieldDmes1Mask)
	}
}

// registerWpsncur1rType FLASH write sector protection for bank 1
type registerWpsncur1rType uint32

const (
	RegisterWpsncur1rFieldWrpsn1Shift = 0
	RegisterWpsncur1rFieldWrpsn1Mask  = 0xff
)

// GetWrpsn1 Bank 1 sector write protection option status byte
func (r *registerWpsncur1rType) GetWrpsn1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWpsncur1rFieldWrpsn1Mask) >> RegisterWpsncur1rFieldWrpsn1Shift)
}

// SetWrpsn1 Bank 1 sector write protection option status byte
func (r *registerWpsncur1rType) SetWrpsn1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWpsncur1rFieldWrpsn1Mask)|(uint32(value)<<RegisterWpsncur1rFieldWrpsn1Shift))
}

// registerWpsnprg1rType FLASH write sector protection for bank 1
type registerWpsnprg1rType uint32

const (
	RegisterWpsnprg1rFieldWrpsn1Shift = 0
	RegisterWpsnprg1rFieldWrpsn1Mask  = 0xff
)

// GetWrpsn1 Bank 1 sector write protection configuration byte
func (r *registerWpsnprg1rType) GetWrpsn1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWpsnprg1rFieldWrpsn1Mask) >> RegisterWpsnprg1rFieldWrpsn1Shift)
}

// SetWrpsn1 Bank 1 sector write protection configuration byte
func (r *registerWpsnprg1rType) SetWrpsn1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWpsnprg1rFieldWrpsn1Mask)|(uint32(value)<<RegisterWpsnprg1rFieldWrpsn1Shift))
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

// registerCrccr1Type FLASH CRC control register for bank 1
type registerCrccr1Type uint32

const (
	RegisterCrccr1FieldCrcsectShift = 0
	RegisterCrccr1FieldCrcsectMask  = 0x7
)

// GetCrcsect Bank 1 CRC sector number
func (r *registerCrccr1Type) GetCrcsect() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldCrcsectMask) >> RegisterCrccr1FieldCrcsectShift)
}

// SetCrcsect Bank 1 CRC sector number
func (r *registerCrccr1Type) SetCrcsect(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldCrcsectMask)|(uint32(value)<<RegisterCrccr1FieldCrcsectShift))
}

const (
	RegisterCrccr1FieldAllbankShift = 7
	RegisterCrccr1FieldAllbankMask  = 0x80
)

// GetAllbank Bank 1 CRC select bit
func (r *registerCrccr1Type) GetAllbank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldAllbankMask) != 0
}

// SetAllbank Bank 1 CRC select bit
func (r *registerCrccr1Type) SetAllbank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr1FieldAllbankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldAllbankMask)
	}
}

const (
	RegisterCrccr1FieldCrcbysectShift = 8
	RegisterCrccr1FieldCrcbysectMask  = 0x100
)

// GetCrcbysect Bank 1 CRC sector mode select bit
func (r *registerCrccr1Type) GetCrcbysect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldCrcbysectMask) != 0
}

// SetCrcbysect Bank 1 CRC sector mode select bit
func (r *registerCrccr1Type) SetCrcbysect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr1FieldCrcbysectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldCrcbysectMask)
	}
}

const (
	RegisterCrccr1FieldAddsectShift = 9
	RegisterCrccr1FieldAddsectMask  = 0x200
)

// GetAddsect Bank 1 CRC sector select bit
func (r *registerCrccr1Type) GetAddsect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldAddsectMask) != 0
}

// SetAddsect Bank 1 CRC sector select bit
func (r *registerCrccr1Type) SetAddsect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr1FieldAddsectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldAddsectMask)
	}
}

const (
	RegisterCrccr1FieldCleansectShift = 10
	RegisterCrccr1FieldCleansectMask  = 0x400
)

// GetCleansect Bank 1 CRC sector list clear bit
func (r *registerCrccr1Type) GetCleansect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldCleansectMask) != 0
}

// SetCleansect Bank 1 CRC sector list clear bit
func (r *registerCrccr1Type) SetCleansect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr1FieldCleansectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldCleansectMask)
	}
}

const (
	RegisterCrccr1FieldStartcrcShift = 16
	RegisterCrccr1FieldStartcrcMask  = 0x10000
)

// GetStartcrc Bank 1 CRC start bit
func (r *registerCrccr1Type) GetStartcrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldStartcrcMask) != 0
}

// SetStartcrc Bank 1 CRC start bit
func (r *registerCrccr1Type) SetStartcrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr1FieldStartcrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldStartcrcMask)
	}
}

const (
	RegisterCrccr1FieldCleancrcShift = 17
	RegisterCrccr1FieldCleancrcMask  = 0x20000
)

// GetCleancrc Bank 1 CRC clear bit
func (r *registerCrccr1Type) GetCleancrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldCleancrcMask) != 0
}

// SetCleancrc Bank 1 CRC clear bit
func (r *registerCrccr1Type) SetCleancrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr1FieldCleancrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldCleancrcMask)
	}
}

const (
	RegisterCrccr1FieldCrcburstShift = 20
	RegisterCrccr1FieldCrcburstMask  = 0x300000
)

// GetCrcburst Bank 1 CRC burst size
func (r *registerCrccr1Type) GetCrcburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldCrcburstMask) >> RegisterCrccr1FieldCrcburstShift)
}

// SetCrcburst Bank 1 CRC burst size
func (r *registerCrccr1Type) SetCrcburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldCrcburstMask)|(uint32(value)<<RegisterCrccr1FieldCrcburstShift))
}

// registerCrcsadd1rType FLASH CRC start address register for bank 1
type registerCrcsadd1rType uint32

const (
	RegisterCrcsadd1rFieldCrcstartaddrShift = 0
	RegisterCrcsadd1rFieldCrcstartaddrMask  = 0xffffffff
)

// GetCrcstartaddr CRC start address on bank 1
func (r *registerCrcsadd1rType) GetCrcstartaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrcsadd1rFieldCrcstartaddrMask) >> RegisterCrcsadd1rFieldCrcstartaddrShift)
}

// SetCrcstartaddr CRC start address on bank 1
func (r *registerCrcsadd1rType) SetCrcstartaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrcsadd1rFieldCrcstartaddrMask)|(uint32(value)<<RegisterCrcsadd1rFieldCrcstartaddrShift))
}

// registerCrceadd1rType FLASH CRC end address register for bank 1
type registerCrceadd1rType uint32

const (
	RegisterCrceadd1rFieldCrcendaddrShift = 0
	RegisterCrceadd1rFieldCrcendaddrMask  = 0xffffffff
)

// GetCrcendaddr CRC end address on bank 1
func (r *registerCrceadd1rType) GetCrcendaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrceadd1rFieldCrcendaddrMask) >> RegisterCrceadd1rFieldCrcendaddrShift)
}

// SetCrcendaddr CRC end address on bank 1
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

// registerEccfa1rType FLASH ECC fail address for bank 1
type registerEccfa1rType uint32

const (
	RegisterEccfa1rFieldFaileccaddr1Shift = 0
	RegisterEccfa1rFieldFaileccaddr1Mask  = 0x7fff
)

// GetFaileccaddr1 Bank 1 ECC error address
func (r *registerEccfa1rType) GetFaileccaddr1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterEccfa1rFieldFaileccaddr1Mask) >> RegisterEccfa1rFieldFaileccaddr1Shift)
}

// SetFaileccaddr1 Bank 1 ECC error address
func (r *registerEccfa1rType) SetFaileccaddr1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEccfa1rFieldFaileccaddr1Mask)|(uint32(value)<<RegisterEccfa1rFieldFaileccaddr1Shift))
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

// registerKeyr2Type FLASH key register for bank 2
type registerKeyr2Type uint32

const (
	RegisterKeyr2FieldKeyr2Shift = 0
	RegisterKeyr2FieldKeyr2Mask  = 0xffffffff
)

// GetKeyr2 Bank 2 access configuration unlock key
func (r *registerKeyr2Type) GetKeyr2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterKeyr2FieldKeyr2Mask) >> RegisterKeyr2FieldKeyr2Shift)
}

// SetKeyr2 Bank 2 access configuration unlock key
func (r *registerKeyr2Type) SetKeyr2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterKeyr2FieldKeyr2Mask)|(uint32(value)<<RegisterKeyr2FieldKeyr2Shift))
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

// registerCr2Type FLASH control register for bank 2
type registerCr2Type uint32

const (
	RegisterCr2FieldLock2Shift = 0
	RegisterCr2FieldLock2Mask  = 0x1
)

// GetLock2 Bank 2 configuration lock bit
func (r *registerCr2Type) GetLock2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldLock2Mask) != 0
}

// SetLock2 Bank 2 configuration lock bit
func (r *registerCr2Type) SetLock2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldLock2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldLock2Mask)
	}
}

const (
	RegisterCr2FieldPg2Shift = 1
	RegisterCr2FieldPg2Mask  = 0x2
)

// GetPg2 Bank 2 program enable bit
func (r *registerCr2Type) GetPg2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldPg2Mask) != 0
}

// SetPg2 Bank 2 program enable bit
func (r *registerCr2Type) SetPg2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldPg2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldPg2Mask)
	}
}

const (
	RegisterCr2FieldSer2Shift = 2
	RegisterCr2FieldSer2Mask  = 0x4
)

// GetSer2 Bank 2 sector erase request
func (r *registerCr2Type) GetSer2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSer2Mask) != 0
}

// SetSer2 Bank 2 sector erase request
func (r *registerCr2Type) SetSer2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSer2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSer2Mask)
	}
}

const (
	RegisterCr2FieldBer2Shift = 3
	RegisterCr2FieldBer2Mask  = 0x8
)

// GetBer2 Bank 2 erase request
func (r *registerCr2Type) GetBer2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldBer2Mask) != 0
}

// SetBer2 Bank 2 erase request
func (r *registerCr2Type) SetBer2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldBer2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldBer2Mask)
	}
}

const (
	RegisterCr2FieldPsize2Shift = 4
	RegisterCr2FieldPsize2Mask  = 0x30
)

// GetPsize2 Bank 2 program size
func (r *registerCr2Type) GetPsize2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldPsize2Mask) >> RegisterCr2FieldPsize2Shift)
}

// SetPsize2 Bank 2 program size
func (r *registerCr2Type) SetPsize2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldPsize2Mask)|(uint32(value)<<RegisterCr2FieldPsize2Shift))
}

const (
	RegisterCr2FieldFw2Shift = 6
	RegisterCr2FieldFw2Mask  = 0x40
)

// GetFw2 Bank 2 write forcing control bit
func (r *registerCr2Type) GetFw2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldFw2Mask) != 0
}

// SetFw2 Bank 2 write forcing control bit
func (r *registerCr2Type) SetFw2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldFw2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldFw2Mask)
	}
}

const (
	RegisterCr2FieldStart2Shift = 7
	RegisterCr2FieldStart2Mask  = 0x80
)

// GetStart2 Bank 2 bank or sector erase start control bit
func (r *registerCr2Type) GetStart2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldStart2Mask) != 0
}

// SetStart2 Bank 2 bank or sector erase start control bit
func (r *registerCr2Type) SetStart2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldStart2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldStart2Mask)
	}
}

const (
	RegisterCr2FieldSnb2Shift = 8
	RegisterCr2FieldSnb2Mask  = 0x700
)

// GetSnb2 Bank 2 sector erase selection number
func (r *registerCr2Type) GetSnb2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSnb2Mask) >> RegisterCr2FieldSnb2Shift)
}

// SetSnb2 Bank 2 sector erase selection number
func (r *registerCr2Type) SetSnb2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSnb2Mask)|(uint32(value)<<RegisterCr2FieldSnb2Shift))
}

const (
	RegisterCr2FieldCrcenShift = 15
	RegisterCr2FieldCrcenMask  = 0x8000
)

// GetCrcen Bank 2 CRC control bit
func (r *registerCr2Type) GetCrcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCrcenMask) != 0
}

// SetCrcen Bank 2 CRC control bit
func (r *registerCr2Type) SetCrcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldCrcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldCrcenMask)
	}
}

const (
	RegisterCr2FieldEopie2Shift = 16
	RegisterCr2FieldEopie2Mask  = 0x10000
)

// GetEopie2 Bank 2 end-of-program interrupt control bit
func (r *registerCr2Type) GetEopie2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldEopie2Mask) != 0
}

// SetEopie2 Bank 2 end-of-program interrupt control bit
func (r *registerCr2Type) SetEopie2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldEopie2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldEopie2Mask)
	}
}

const (
	RegisterCr2FieldWrperrie2Shift = 17
	RegisterCr2FieldWrperrie2Mask  = 0x20000
)

// GetWrperrie2 Bank 2 write protection error interrupt enable bit
func (r *registerCr2Type) GetWrperrie2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldWrperrie2Mask) != 0
}

// SetWrperrie2 Bank 2 write protection error interrupt enable bit
func (r *registerCr2Type) SetWrperrie2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldWrperrie2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldWrperrie2Mask)
	}
}

const (
	RegisterCr2FieldPgserrie2Shift = 18
	RegisterCr2FieldPgserrie2Mask  = 0x40000
)

// GetPgserrie2 Bank 2 programming sequence error interrupt enable bit
func (r *registerCr2Type) GetPgserrie2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldPgserrie2Mask) != 0
}

// SetPgserrie2 Bank 2 programming sequence error interrupt enable bit
func (r *registerCr2Type) SetPgserrie2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldPgserrie2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldPgserrie2Mask)
	}
}

const (
	RegisterCr2FieldStrberrie2Shift = 19
	RegisterCr2FieldStrberrie2Mask  = 0x80000
)

// GetStrberrie2 Bank 2 strobe error interrupt enable bit
func (r *registerCr2Type) GetStrberrie2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldStrberrie2Mask) != 0
}

// SetStrberrie2 Bank 2 strobe error interrupt enable bit
func (r *registerCr2Type) SetStrberrie2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldStrberrie2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldStrberrie2Mask)
	}
}

const (
	RegisterCr2FieldIncerrie2Shift = 21
	RegisterCr2FieldIncerrie2Mask  = 0x200000
)

// GetIncerrie2 Bank 2 inconsistency error interrupt enable bit
func (r *registerCr2Type) GetIncerrie2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldIncerrie2Mask) != 0
}

// SetIncerrie2 Bank 2 inconsistency error interrupt enable bit
func (r *registerCr2Type) SetIncerrie2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldIncerrie2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldIncerrie2Mask)
	}
}

const (
	RegisterCr2FieldOperrie2Shift = 22
	RegisterCr2FieldOperrie2Mask  = 0x400000
)

// GetOperrie2 Bank 2 write/erase error interrupt enable bit
func (r *registerCr2Type) GetOperrie2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldOperrie2Mask) != 0
}

// SetOperrie2 Bank 2 write/erase error interrupt enable bit
func (r *registerCr2Type) SetOperrie2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldOperrie2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldOperrie2Mask)
	}
}

const (
	RegisterCr2FieldRdperrie2Shift = 23
	RegisterCr2FieldRdperrie2Mask  = 0x800000
)

// GetRdperrie2 Bank 2 read protection error interrupt enable bit
func (r *registerCr2Type) GetRdperrie2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldRdperrie2Mask) != 0
}

// SetRdperrie2 Bank 2 read protection error interrupt enable bit
func (r *registerCr2Type) SetRdperrie2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldRdperrie2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldRdperrie2Mask)
	}
}

const (
	RegisterCr2FieldRdserrie2Shift = 24
	RegisterCr2FieldRdserrie2Mask  = 0x1000000
)

// GetRdserrie2 Bank 2 secure error interrupt enable bit
func (r *registerCr2Type) GetRdserrie2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldRdserrie2Mask) != 0
}

// SetRdserrie2 Bank 2 secure error interrupt enable bit
func (r *registerCr2Type) SetRdserrie2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldRdserrie2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldRdserrie2Mask)
	}
}

const (
	RegisterCr2FieldSneccerrie2Shift = 25
	RegisterCr2FieldSneccerrie2Mask  = 0x2000000
)

// GetSneccerrie2 Bank 2 ECC single correction error interrupt enable bit
func (r *registerCr2Type) GetSneccerrie2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSneccerrie2Mask) != 0
}

// SetSneccerrie2 Bank 2 ECC single correction error interrupt enable bit
func (r *registerCr2Type) SetSneccerrie2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSneccerrie2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSneccerrie2Mask)
	}
}

const (
	RegisterCr2FieldDbeccerrie2Shift = 26
	RegisterCr2FieldDbeccerrie2Mask  = 0x4000000
)

// GetDbeccerrie2 Bank 2 ECC double detection error interrupt enable bit
func (r *registerCr2Type) GetDbeccerrie2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldDbeccerrie2Mask) != 0
}

// SetDbeccerrie2 Bank 2 ECC double detection error interrupt enable bit
func (r *registerCr2Type) SetDbeccerrie2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldDbeccerrie2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldDbeccerrie2Mask)
	}
}

const (
	RegisterCr2FieldCrcendie2Shift = 27
	RegisterCr2FieldCrcendie2Mask  = 0x8000000
)

// GetCrcendie2 Bank 2 end of CRC calculation interrupt enable bit
func (r *registerCr2Type) GetCrcendie2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCrcendie2Mask) != 0
}

// SetCrcendie2 Bank 2 end of CRC calculation interrupt enable bit
func (r *registerCr2Type) SetCrcendie2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldCrcendie2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldCrcendie2Mask)
	}
}

// registerSr2Type FLASH status register for bank 2
type registerSr2Type uint32

const (
	RegisterSr2FieldBsy2Shift = 0
	RegisterSr2FieldBsy2Mask  = 0x1
)

// GetBsy2 Bank 2 ongoing program flag
func (r *registerSr2Type) GetBsy2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldBsy2Mask) != 0
}

// SetBsy2 Bank 2 ongoing program flag
func (r *registerSr2Type) SetBsy2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldBsy2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldBsy2Mask)
	}
}

const (
	RegisterSr2FieldWbne2Shift = 1
	RegisterSr2FieldWbne2Mask  = 0x2
)

// GetWbne2 Bank 2 write buffer not empty flag
func (r *registerSr2Type) GetWbne2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldWbne2Mask) != 0
}

// SetWbne2 Bank 2 write buffer not empty flag
func (r *registerSr2Type) SetWbne2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldWbne2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldWbne2Mask)
	}
}

const (
	RegisterSr2FieldQw2Shift = 2
	RegisterSr2FieldQw2Mask  = 0x4
)

// GetQw2 Bank 2 wait queue flag
func (r *registerSr2Type) GetQw2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldQw2Mask) != 0
}

// SetQw2 Bank 2 wait queue flag
func (r *registerSr2Type) SetQw2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldQw2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldQw2Mask)
	}
}

const (
	RegisterSr2FieldCrcbusy2Shift = 3
	RegisterSr2FieldCrcbusy2Mask  = 0x8
)

// GetCrcbusy2 Bank 2 CRC busy flag
func (r *registerSr2Type) GetCrcbusy2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldCrcbusy2Mask) != 0
}

// SetCrcbusy2 Bank 2 CRC busy flag
func (r *registerSr2Type) SetCrcbusy2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldCrcbusy2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldCrcbusy2Mask)
	}
}

const (
	RegisterSr2FieldEop2Shift = 16
	RegisterSr2FieldEop2Mask  = 0x10000
)

// GetEop2 Bank 2 end-of-program flag
func (r *registerSr2Type) GetEop2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldEop2Mask) != 0
}

// SetEop2 Bank 2 end-of-program flag
func (r *registerSr2Type) SetEop2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldEop2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldEop2Mask)
	}
}

const (
	RegisterSr2FieldWrperr2Shift = 17
	RegisterSr2FieldWrperr2Mask  = 0x20000
)

// GetWrperr2 Bank 2 write protection error flag
func (r *registerSr2Type) GetWrperr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldWrperr2Mask) != 0
}

// SetWrperr2 Bank 2 write protection error flag
func (r *registerSr2Type) SetWrperr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldWrperr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldWrperr2Mask)
	}
}

const (
	RegisterSr2FieldPgserr2Shift = 18
	RegisterSr2FieldPgserr2Mask  = 0x40000
)

// GetPgserr2 Bank 2 programming sequence error flag
func (r *registerSr2Type) GetPgserr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldPgserr2Mask) != 0
}

// SetPgserr2 Bank 2 programming sequence error flag
func (r *registerSr2Type) SetPgserr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldPgserr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldPgserr2Mask)
	}
}

const (
	RegisterSr2FieldStrberr2Shift = 19
	RegisterSr2FieldStrberr2Mask  = 0x80000
)

// GetStrberr2 Bank 2 strobe error flag
func (r *registerSr2Type) GetStrberr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldStrberr2Mask) != 0
}

// SetStrberr2 Bank 2 strobe error flag
func (r *registerSr2Type) SetStrberr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldStrberr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldStrberr2Mask)
	}
}

const (
	RegisterSr2FieldIncerr2Shift = 21
	RegisterSr2FieldIncerr2Mask  = 0x200000
)

// GetIncerr2 Bank 2 inconsistency error flag
func (r *registerSr2Type) GetIncerr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldIncerr2Mask) != 0
}

// SetIncerr2 Bank 2 inconsistency error flag
func (r *registerSr2Type) SetIncerr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldIncerr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldIncerr2Mask)
	}
}

const (
	RegisterSr2FieldOperr2Shift = 22
	RegisterSr2FieldOperr2Mask  = 0x400000
)

// GetOperr2 Bank 2 write/erase error flag
func (r *registerSr2Type) GetOperr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldOperr2Mask) != 0
}

// SetOperr2 Bank 2 write/erase error flag
func (r *registerSr2Type) SetOperr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldOperr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldOperr2Mask)
	}
}

const (
	RegisterSr2FieldRdperr2Shift = 23
	RegisterSr2FieldRdperr2Mask  = 0x800000
)

// GetRdperr2 Bank 2 read protection error flag
func (r *registerSr2Type) GetRdperr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldRdperr2Mask) != 0
}

// SetRdperr2 Bank 2 read protection error flag
func (r *registerSr2Type) SetRdperr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldRdperr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldRdperr2Mask)
	}
}

const (
	RegisterSr2FieldRdserr2Shift = 24
	RegisterSr2FieldRdserr2Mask  = 0x1000000
)

// GetRdserr2 Bank 2 secure error flag
func (r *registerSr2Type) GetRdserr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldRdserr2Mask) != 0
}

// SetRdserr2 Bank 2 secure error flag
func (r *registerSr2Type) SetRdserr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldRdserr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldRdserr2Mask)
	}
}

const (
	RegisterSr2FieldSneccerr2Shift = 25
	RegisterSr2FieldSneccerr2Mask  = 0x2000000
)

// GetSneccerr2 Bank 2 single correction error flag
func (r *registerSr2Type) GetSneccerr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldSneccerr2Mask) != 0
}

// SetSneccerr2 Bank 2 single correction error flag
func (r *registerSr2Type) SetSneccerr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldSneccerr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldSneccerr2Mask)
	}
}

const (
	RegisterSr2FieldDbeccerr2Shift = 26
	RegisterSr2FieldDbeccerr2Mask  = 0x4000000
)

// GetDbeccerr2 Bank 2 ECC double detection error flag
func (r *registerSr2Type) GetDbeccerr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldDbeccerr2Mask) != 0
}

// SetDbeccerr2 Bank 2 ECC double detection error flag
func (r *registerSr2Type) SetDbeccerr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldDbeccerr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldDbeccerr2Mask)
	}
}

const (
	RegisterSr2FieldCrcend2Shift = 27
	RegisterSr2FieldCrcend2Mask  = 0x8000000
)

// GetCrcend2 Bank 2 CRC-complete flag
func (r *registerSr2Type) GetCrcend2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldCrcend2Mask) != 0
}

// SetCrcend2 Bank 2 CRC-complete flag
func (r *registerSr2Type) SetCrcend2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldCrcend2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldCrcend2Mask)
	}
}

// registerCcr2Type FLASH clear control register for bank 2
type registerCcr2Type uint32

const (
	RegisterCcr2FieldClreop2Shift = 16
	RegisterCcr2FieldClreop2Mask  = 0x10000
)

// GetClreop2 Bank 1 EOP1 flag clear bit
func (r *registerCcr2Type) GetClreop2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClreop2Mask) != 0
}

// SetClreop2 Bank 1 EOP1 flag clear bit
func (r *registerCcr2Type) SetClreop2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClreop2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClreop2Mask)
	}
}

const (
	RegisterCcr2FieldClrwrperr2Shift = 17
	RegisterCcr2FieldClrwrperr2Mask  = 0x20000
)

// GetClrwrperr2 Bank 2 WRPERR1 flag clear bit
func (r *registerCcr2Type) GetClrwrperr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClrwrperr2Mask) != 0
}

// SetClrwrperr2 Bank 2 WRPERR1 flag clear bit
func (r *registerCcr2Type) SetClrwrperr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClrwrperr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClrwrperr2Mask)
	}
}

const (
	RegisterCcr2FieldClrpgserr2Shift = 18
	RegisterCcr2FieldClrpgserr2Mask  = 0x40000
)

// GetClrpgserr2 Bank 2 PGSERR1 flag clear bi
func (r *registerCcr2Type) GetClrpgserr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClrpgserr2Mask) != 0
}

// SetClrpgserr2 Bank 2 PGSERR1 flag clear bi
func (r *registerCcr2Type) SetClrpgserr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClrpgserr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClrpgserr2Mask)
	}
}

const (
	RegisterCcr2FieldClrstrberr2Shift = 19
	RegisterCcr2FieldClrstrberr2Mask  = 0x80000
)

// GetClrstrberr2 Bank 2 STRBERR1 flag clear bit
func (r *registerCcr2Type) GetClrstrberr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClrstrberr2Mask) != 0
}

// SetClrstrberr2 Bank 2 STRBERR1 flag clear bit
func (r *registerCcr2Type) SetClrstrberr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClrstrberr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClrstrberr2Mask)
	}
}

const (
	RegisterCcr2FieldClrincerr2Shift = 21
	RegisterCcr2FieldClrincerr2Mask  = 0x200000
)

// GetClrincerr2 Bank 2 INCERR1 flag clear bit
func (r *registerCcr2Type) GetClrincerr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClrincerr2Mask) != 0
}

// SetClrincerr2 Bank 2 INCERR1 flag clear bit
func (r *registerCcr2Type) SetClrincerr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClrincerr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClrincerr2Mask)
	}
}

const (
	RegisterCcr2FieldClroperr2Shift = 22
	RegisterCcr2FieldClroperr2Mask  = 0x400000
)

// GetClroperr2 Bank 2 OPERR1 flag clear bit
func (r *registerCcr2Type) GetClroperr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClroperr2Mask) != 0
}

// SetClroperr2 Bank 2 OPERR1 flag clear bit
func (r *registerCcr2Type) SetClroperr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClroperr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClroperr2Mask)
	}
}

const (
	RegisterCcr2FieldClrrdperr2Shift = 23
	RegisterCcr2FieldClrrdperr2Mask  = 0x800000
)

// GetClrrdperr2 Bank 2 RDPERR1 flag clear bit
func (r *registerCcr2Type) GetClrrdperr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClrrdperr2Mask) != 0
}

// SetClrrdperr2 Bank 2 RDPERR1 flag clear bit
func (r *registerCcr2Type) SetClrrdperr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClrrdperr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClrrdperr2Mask)
	}
}

const (
	RegisterCcr2FieldClrrdserr1Shift = 24
	RegisterCcr2FieldClrrdserr1Mask  = 0x1000000
)

// GetClrrdserr1 Bank 1 RDSERR1 flag clear bit
func (r *registerCcr2Type) GetClrrdserr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClrrdserr1Mask) != 0
}

// SetClrrdserr1 Bank 1 RDSERR1 flag clear bit
func (r *registerCcr2Type) SetClrrdserr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClrrdserr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClrrdserr1Mask)
	}
}

const (
	RegisterCcr2FieldClrsneccerr2Shift = 25
	RegisterCcr2FieldClrsneccerr2Mask  = 0x2000000
)

// GetClrsneccerr2 Bank 2 SNECCERR1 flag clear bit
func (r *registerCcr2Type) GetClrsneccerr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClrsneccerr2Mask) != 0
}

// SetClrsneccerr2 Bank 2 SNECCERR1 flag clear bit
func (r *registerCcr2Type) SetClrsneccerr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClrsneccerr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClrsneccerr2Mask)
	}
}

const (
	RegisterCcr2FieldClrdbeccerr1Shift = 26
	RegisterCcr2FieldClrdbeccerr1Mask  = 0x4000000
)

// GetClrdbeccerr1 Bank 1 DBECCERR1 flag clear bit
func (r *registerCcr2Type) GetClrdbeccerr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClrdbeccerr1Mask) != 0
}

// SetClrdbeccerr1 Bank 1 DBECCERR1 flag clear bit
func (r *registerCcr2Type) SetClrdbeccerr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClrdbeccerr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClrdbeccerr1Mask)
	}
}

const (
	RegisterCcr2FieldClrcrcend2Shift = 27
	RegisterCcr2FieldClrcrcend2Mask  = 0x8000000
)

// GetClrcrcend2 Bank 2 CRCEND1 flag clear bit
func (r *registerCcr2Type) GetClrcrcend2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClrcrcend2Mask) != 0
}

// SetClrcrcend2 Bank 2 CRCEND1 flag clear bit
func (r *registerCcr2Type) SetClrcrcend2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClrcrcend2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClrcrcend2Mask)
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
	RegisterOptsrcurFieldNrststopd1Shift = 6
	RegisterOptsrcurFieldNrststopd1Mask  = 0x40
)

// GetNrststopd1 D1 DStop entry reset option status bit
func (r *registerOptsrcurType) GetNrststopd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldNrststopd1Mask) != 0
}

// SetNrststopd1 D1 DStop entry reset option status bit
func (r *registerOptsrcurType) SetNrststopd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldNrststopd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldNrststopd1Mask)
	}
}

const (
	RegisterOptsrcurFieldNrststbyd1Shift = 7
	RegisterOptsrcurFieldNrststbyd1Mask  = 0x80
)

// GetNrststbyd1 D1 DStandby entry reset option status bit
func (r *registerOptsrcurType) GetNrststbyd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsrcurFieldNrststbyd1Mask) != 0
}

// SetNrststbyd1 D1 DStandby entry reset option status bit
func (r *registerOptsrcurType) SetNrststbyd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsrcurFieldNrststbyd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsrcurFieldNrststbyd1Mask)
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

// registerPrarcur2Type FLASH protection address for bank 1
type registerPrarcur2Type uint32

const (
	RegisterPrarcur2FieldProtareastart2Shift = 0
	RegisterPrarcur2FieldProtareastart2Mask  = 0xfff
)

// GetProtareastart2 Bank 2 lowest PCROP protected address
func (r *registerPrarcur2Type) GetProtareastart2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarcur2FieldProtareastart2Mask) >> RegisterPrarcur2FieldProtareastart2Shift)
}

// SetProtareastart2 Bank 2 lowest PCROP protected address
func (r *registerPrarcur2Type) SetProtareastart2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarcur2FieldProtareastart2Mask)|(uint32(value)<<RegisterPrarcur2FieldProtareastart2Shift))
}

const (
	RegisterPrarcur2FieldProtareaend2Shift = 16
	RegisterPrarcur2FieldProtareaend2Mask  = 0xfff0000
)

// GetProtareaend2 Bank 2 highest PCROP protected address
func (r *registerPrarcur2Type) GetProtareaend2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrarcur2FieldProtareaend2Mask) >> RegisterPrarcur2FieldProtareaend2Shift)
}

// SetProtareaend2 Bank 2 highest PCROP protected address
func (r *registerPrarcur2Type) SetProtareaend2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrarcur2FieldProtareaend2Mask)|(uint32(value)<<RegisterPrarcur2FieldProtareaend2Shift))
}

const (
	RegisterPrarcur2FieldDmep2Shift = 31
	RegisterPrarcur2FieldDmep2Mask  = 0x80000000
)

// GetDmep2 Bank 2 PCROP protected erase enable option status bit
func (r *registerPrarcur2Type) GetDmep2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrarcur2FieldDmep2Mask) != 0
}

// SetDmep2 Bank 2 PCROP protected erase enable option status bit
func (r *registerPrarcur2Type) SetDmep2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrarcur2FieldDmep2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrarcur2FieldDmep2Mask)
	}
}

// registerScarcur2Type FLASH secure address for bank 2
type registerScarcur2Type uint32

const (
	RegisterScarcur2FieldSecareastart2Shift = 0
	RegisterScarcur2FieldSecareastart2Mask  = 0xfff
)

// GetSecareastart2 Bank 2 lowest secure protected address
func (r *registerScarcur2Type) GetSecareastart2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarcur2FieldSecareastart2Mask) >> RegisterScarcur2FieldSecareastart2Shift)
}

// SetSecareastart2 Bank 2 lowest secure protected address
func (r *registerScarcur2Type) SetSecareastart2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarcur2FieldSecareastart2Mask)|(uint32(value)<<RegisterScarcur2FieldSecareastart2Shift))
}

const (
	RegisterScarcur2FieldSecareaend2Shift = 16
	RegisterScarcur2FieldSecareaend2Mask  = 0xfff0000
)

// GetSecareaend2 Bank 2 highest secure protected address
func (r *registerScarcur2Type) GetSecareaend2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarcur2FieldSecareaend2Mask) >> RegisterScarcur2FieldSecareaend2Shift)
}

// SetSecareaend2 Bank 2 highest secure protected address
func (r *registerScarcur2Type) SetSecareaend2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarcur2FieldSecareaend2Mask)|(uint32(value)<<RegisterScarcur2FieldSecareaend2Shift))
}

const (
	RegisterScarcur2FieldDmes2Shift = 31
	RegisterScarcur2FieldDmes2Mask  = 0x80000000
)

// GetDmes2 Bank 2 secure protected erase enable option status bit
func (r *registerScarcur2Type) GetDmes2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScarcur2FieldDmes2Mask) != 0
}

// SetDmes2 Bank 2 secure protected erase enable option status bit
func (r *registerScarcur2Type) SetDmes2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScarcur2FieldDmes2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScarcur2FieldDmes2Mask)
	}
}

// registerScarprg2Type FLASH secure address for bank 2
type registerScarprg2Type uint32

const (
	RegisterScarprg2FieldSecareastart2Shift = 0
	RegisterScarprg2FieldSecareastart2Mask  = 0xfff
)

// GetSecareastart2 Bank 2 lowest secure protected address configuration
func (r *registerScarprg2Type) GetSecareastart2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarprg2FieldSecareastart2Mask) >> RegisterScarprg2FieldSecareastart2Shift)
}

// SetSecareastart2 Bank 2 lowest secure protected address configuration
func (r *registerScarprg2Type) SetSecareastart2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarprg2FieldSecareastart2Mask)|(uint32(value)<<RegisterScarprg2FieldSecareastart2Shift))
}

const (
	RegisterScarprg2FieldSecareaend2Shift = 16
	RegisterScarprg2FieldSecareaend2Mask  = 0xfff0000
)

// GetSecareaend2 Bank 2 highest secure protected address configuration
func (r *registerScarprg2Type) GetSecareaend2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScarprg2FieldSecareaend2Mask) >> RegisterScarprg2FieldSecareaend2Shift)
}

// SetSecareaend2 Bank 2 highest secure protected address configuration
func (r *registerScarprg2Type) SetSecareaend2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScarprg2FieldSecareaend2Mask)|(uint32(value)<<RegisterScarprg2FieldSecareaend2Shift))
}

const (
	RegisterScarprg2FieldDmes2Shift = 31
	RegisterScarprg2FieldDmes2Mask  = 0x80000000
)

// GetDmes2 Bank 2 secure protected erase enable option configuration bit
func (r *registerScarprg2Type) GetDmes2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScarprg2FieldDmes2Mask) != 0
}

// SetDmes2 Bank 2 secure protected erase enable option configuration bit
func (r *registerScarprg2Type) SetDmes2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScarprg2FieldDmes2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScarprg2FieldDmes2Mask)
	}
}

// registerWpsncur2rType FLASH write sector protection for bank 2
type registerWpsncur2rType uint32

const (
	RegisterWpsncur2rFieldWrpsn2Shift = 0
	RegisterWpsncur2rFieldWrpsn2Mask  = 0xff
)

// GetWrpsn2 Bank 2 sector write protection option status byte
func (r *registerWpsncur2rType) GetWrpsn2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWpsncur2rFieldWrpsn2Mask) >> RegisterWpsncur2rFieldWrpsn2Shift)
}

// SetWrpsn2 Bank 2 sector write protection option status byte
func (r *registerWpsncur2rType) SetWrpsn2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWpsncur2rFieldWrpsn2Mask)|(uint32(value)<<RegisterWpsncur2rFieldWrpsn2Shift))
}

// registerWpsnprg2rType FLASH write sector protection for bank 2
type registerWpsnprg2rType uint32

const (
	RegisterWpsnprg2rFieldWrpsn2Shift = 0
	RegisterWpsnprg2rFieldWrpsn2Mask  = 0xff
)

// GetWrpsn2 Bank 2 sector write protection configuration byte
func (r *registerWpsnprg2rType) GetWrpsn2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWpsnprg2rFieldWrpsn2Mask) >> RegisterWpsnprg2rFieldWrpsn2Shift)
}

// SetWrpsn2 Bank 2 sector write protection configuration byte
func (r *registerWpsnprg2rType) SetWrpsn2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWpsnprg2rFieldWrpsn2Mask)|(uint32(value)<<RegisterWpsnprg2rFieldWrpsn2Shift))
}

// registerCrccr2Type FLASH CRC control register for bank 1
type registerCrccr2Type uint32

const (
	RegisterCrccr2FieldCrcsectShift = 0
	RegisterCrccr2FieldCrcsectMask  = 0x7
)

// GetCrcsect Bank 2 CRC sector number
func (r *registerCrccr2Type) GetCrcsect() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldCrcsectMask) >> RegisterCrccr2FieldCrcsectShift)
}

// SetCrcsect Bank 2 CRC sector number
func (r *registerCrccr2Type) SetCrcsect(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldCrcsectMask)|(uint32(value)<<RegisterCrccr2FieldCrcsectShift))
}

const (
	RegisterCrccr2FieldAllbankShift = 7
	RegisterCrccr2FieldAllbankMask  = 0x80
)

// GetAllbank Bank 2 CRC select bit
func (r *registerCrccr2Type) GetAllbank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldAllbankMask) != 0
}

// SetAllbank Bank 2 CRC select bit
func (r *registerCrccr2Type) SetAllbank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr2FieldAllbankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldAllbankMask)
	}
}

const (
	RegisterCrccr2FieldCrcbysectShift = 8
	RegisterCrccr2FieldCrcbysectMask  = 0x100
)

// GetCrcbysect Bank 2 CRC sector mode select bit
func (r *registerCrccr2Type) GetCrcbysect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldCrcbysectMask) != 0
}

// SetCrcbysect Bank 2 CRC sector mode select bit
func (r *registerCrccr2Type) SetCrcbysect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr2FieldCrcbysectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldCrcbysectMask)
	}
}

const (
	RegisterCrccr2FieldAddsectShift = 9
	RegisterCrccr2FieldAddsectMask  = 0x200
)

// GetAddsect Bank 2 CRC sector select bit
func (r *registerCrccr2Type) GetAddsect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldAddsectMask) != 0
}

// SetAddsect Bank 2 CRC sector select bit
func (r *registerCrccr2Type) SetAddsect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr2FieldAddsectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldAddsectMask)
	}
}

const (
	RegisterCrccr2FieldCleansectShift = 10
	RegisterCrccr2FieldCleansectMask  = 0x400
)

// GetCleansect Bank 2 CRC sector list clear bit
func (r *registerCrccr2Type) GetCleansect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldCleansectMask) != 0
}

// SetCleansect Bank 2 CRC sector list clear bit
func (r *registerCrccr2Type) SetCleansect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr2FieldCleansectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldCleansectMask)
	}
}

const (
	RegisterCrccr2FieldStartcrcShift = 16
	RegisterCrccr2FieldStartcrcMask  = 0x10000
)

// GetStartcrc Bank 2 CRC start bit
func (r *registerCrccr2Type) GetStartcrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldStartcrcMask) != 0
}

// SetStartcrc Bank 2 CRC start bit
func (r *registerCrccr2Type) SetStartcrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr2FieldStartcrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldStartcrcMask)
	}
}

const (
	RegisterCrccr2FieldCleancrcShift = 17
	RegisterCrccr2FieldCleancrcMask  = 0x20000
)

// GetCleancrc Bank 2 CRC clear bit
func (r *registerCrccr2Type) GetCleancrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldCleancrcMask) != 0
}

// SetCleancrc Bank 2 CRC clear bit
func (r *registerCrccr2Type) SetCleancrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr2FieldCleancrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldCleancrcMask)
	}
}

const (
	RegisterCrccr2FieldCrcburstShift = 20
	RegisterCrccr2FieldCrcburstMask  = 0x300000
)

// GetCrcburst Bank 2 CRC burst size
func (r *registerCrccr2Type) GetCrcburst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldCrcburstMask) >> RegisterCrccr2FieldCrcburstShift)
}

// SetCrcburst Bank 2 CRC burst size
func (r *registerCrccr2Type) SetCrcburst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldCrcburstMask)|(uint32(value)<<RegisterCrccr2FieldCrcburstShift))
}

// registerCrcsadd2rType FLASH CRC start address register for bank 2
type registerCrcsadd2rType uint32

const (
	RegisterCrcsadd2rFieldCrcstartaddrShift = 0
	RegisterCrcsadd2rFieldCrcstartaddrMask  = 0xffffffff
)

// GetCrcstartaddr CRC start address on bank 2
func (r *registerCrcsadd2rType) GetCrcstartaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrcsadd2rFieldCrcstartaddrMask) >> RegisterCrcsadd2rFieldCrcstartaddrShift)
}

// SetCrcstartaddr CRC start address on bank 2
func (r *registerCrcsadd2rType) SetCrcstartaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrcsadd2rFieldCrcstartaddrMask)|(uint32(value)<<RegisterCrcsadd2rFieldCrcstartaddrShift))
}

// registerCrceadd2rType FLASH CRC end address register for bank 2
type registerCrceadd2rType uint32

const (
	RegisterCrceadd2rFieldCrcendaddrShift = 0
	RegisterCrceadd2rFieldCrcendaddrMask  = 0xffffffff
)

// GetCrcendaddr CRC end address on bank 2
func (r *registerCrceadd2rType) GetCrcendaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrceadd2rFieldCrcendaddrMask) >> RegisterCrceadd2rFieldCrcendaddrShift)
}

// SetCrcendaddr CRC end address on bank 2
func (r *registerCrceadd2rType) SetCrcendaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrceadd2rFieldCrcendaddrMask)|(uint32(value)<<RegisterCrceadd2rFieldCrcendaddrShift))
}

// registerEccfa2rType FLASH ECC fail address for bank 2
type registerEccfa2rType uint32

const (
	RegisterEccfa2rFieldFaileccaddr2Shift = 0
	RegisterEccfa2rFieldFaileccaddr2Mask  = 0x7fff
)

// GetFaileccaddr2 Bank 2 ECC error address
func (r *registerEccfa2rType) GetFaileccaddr2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterEccfa2rFieldFaileccaddr2Mask) >> RegisterEccfa2rFieldFaileccaddr2Shift)
}

// SetFaileccaddr2 Bank 2 ECC error address
func (r *registerEccfa2rType) SetFaileccaddr2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEccfa2rFieldFaileccaddr2Mask)|(uint32(value)<<RegisterEccfa2rFieldFaileccaddr2Shift))
}
