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
	Acr        registerAcrType
	Keyr1      registerKeyr1Type
	Optkeyr    registerOptkeyrType
	Cr1        registerCr1Type
	Sr1        registerSr1Type
	Ccr1       registerCcr1Type
	Optcr      registerOptcrType
	Optsr_cur  registerOptsr_curType
	Optsr_prg  registerOptsr_prgType
	Optccr     registerOptccrType
	Prar_cur1  registerPrar_cur1Type
	Prar_prg2  registerPrar_prg2Type
	Prar_prg1  registerPrar_prg1Type
	Scar_cur1  registerScar_cur1Type
	Scar_prg1  registerScar_prg1Type
	Wpsn_cur1r registerWpsn_cur1rType
	Wpsn_prg1r registerWpsn_prg1rType
	Boot_curr  registerBoot_currType
	Boot_prgr  registerBoot_prgrType
	_          [8]byte
	Crccr1     registerCrccr1Type
	Crcsadd1r  registerCrcsadd1rType
	Crceadd1r  registerCrceadd1rType
	Crcdatar   registerCrcdatarType
	Ecc_fa1r   registerEcc_fa1rType
	_          [156]byte
	Acr        registerAcrType
	Keyr2      registerKeyr2Type
	Optkeyr    registerOptkeyrType
	Cr2        registerCr2Type
	Sr2        registerSr2Type
	Ccr2       registerCcr2Type
	Optcr      registerOptcrType
	Optsr_cur  registerOptsr_curType
	Optsr_prg  registerOptsr_prgType
	Optccr     registerOptccrType
	Prar_cur2  registerPrar_cur2Type
	_          [4]byte
	Scar_cur2  registerScar_cur2Type
	Scar_prg2  registerScar_prg2Type
	Wpsn_cur2r registerWpsn_cur2rType
	Wpsn_prg2r registerWpsn_prg2rType
	_          [16]byte
	Crccr2     registerCrccr2Type
	Crcsadd2r  registerCrcsadd2rType
	Crceadd2r  registerCrceadd2rType
	_          [4]byte
	Ecc_fa2r   registerEcc_fa2rType
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
	RegisterCr1FieldCrc_enShift = 15
	RegisterCr1FieldCrc_enMask  = 0x8000
)

// GetCrc_en Bank 1 CRC control bit
func (r *registerCr1Type) GetCrc_en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCrc_enMask) != 0
}

// SetCrc_en Bank 1 CRC control bit
func (r *registerCr1Type) SetCrc_en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldCrc_enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldCrc_enMask)
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
	RegisterSr1FieldCrc_busy1Shift = 3
	RegisterSr1FieldCrc_busy1Mask  = 0x8
)

// GetCrc_busy1 Bank 1 CRC busy flag
func (r *registerSr1Type) GetCrc_busy1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr1FieldCrc_busy1Mask) != 0
}

// SetCrc_busy1 Bank 1 CRC busy flag
func (r *registerSr1Type) SetCrc_busy1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr1FieldCrc_busy1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr1FieldCrc_busy1Mask)
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
	RegisterCcr1FieldClr_eop1Shift = 16
	RegisterCcr1FieldClr_eop1Mask  = 0x10000
)

// GetClr_eop1 Bank 1 EOP1 flag clear bit
func (r *registerCcr1Type) GetClr_eop1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClr_eop1Mask) != 0
}

// SetClr_eop1 Bank 1 EOP1 flag clear bit
func (r *registerCcr1Type) SetClr_eop1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClr_eop1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClr_eop1Mask)
	}
}

const (
	RegisterCcr1FieldClr_wrperr1Shift = 17
	RegisterCcr1FieldClr_wrperr1Mask  = 0x20000
)

// GetClr_wrperr1 Bank 1 WRPERR1 flag clear bit
func (r *registerCcr1Type) GetClr_wrperr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClr_wrperr1Mask) != 0
}

// SetClr_wrperr1 Bank 1 WRPERR1 flag clear bit
func (r *registerCcr1Type) SetClr_wrperr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClr_wrperr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClr_wrperr1Mask)
	}
}

const (
	RegisterCcr1FieldClr_pgserr1Shift = 18
	RegisterCcr1FieldClr_pgserr1Mask  = 0x40000
)

// GetClr_pgserr1 Bank 1 PGSERR1 flag clear bi
func (r *registerCcr1Type) GetClr_pgserr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClr_pgserr1Mask) != 0
}

// SetClr_pgserr1 Bank 1 PGSERR1 flag clear bi
func (r *registerCcr1Type) SetClr_pgserr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClr_pgserr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClr_pgserr1Mask)
	}
}

const (
	RegisterCcr1FieldClr_strberr1Shift = 19
	RegisterCcr1FieldClr_strberr1Mask  = 0x80000
)

// GetClr_strberr1 Bank 1 STRBERR1 flag clear bit
func (r *registerCcr1Type) GetClr_strberr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClr_strberr1Mask) != 0
}

// SetClr_strberr1 Bank 1 STRBERR1 flag clear bit
func (r *registerCcr1Type) SetClr_strberr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClr_strberr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClr_strberr1Mask)
	}
}

const (
	RegisterCcr1FieldClr_incerr1Shift = 21
	RegisterCcr1FieldClr_incerr1Mask  = 0x200000
)

// GetClr_incerr1 Bank 1 INCERR1 flag clear bit
func (r *registerCcr1Type) GetClr_incerr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClr_incerr1Mask) != 0
}

// SetClr_incerr1 Bank 1 INCERR1 flag clear bit
func (r *registerCcr1Type) SetClr_incerr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClr_incerr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClr_incerr1Mask)
	}
}

const (
	RegisterCcr1FieldClr_operr1Shift = 22
	RegisterCcr1FieldClr_operr1Mask  = 0x400000
)

// GetClr_operr1 Bank 1 OPERR1 flag clear bit
func (r *registerCcr1Type) GetClr_operr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClr_operr1Mask) != 0
}

// SetClr_operr1 Bank 1 OPERR1 flag clear bit
func (r *registerCcr1Type) SetClr_operr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClr_operr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClr_operr1Mask)
	}
}

const (
	RegisterCcr1FieldClr_rdperr1Shift = 23
	RegisterCcr1FieldClr_rdperr1Mask  = 0x800000
)

// GetClr_rdperr1 Bank 1 RDPERR1 flag clear bit
func (r *registerCcr1Type) GetClr_rdperr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClr_rdperr1Mask) != 0
}

// SetClr_rdperr1 Bank 1 RDPERR1 flag clear bit
func (r *registerCcr1Type) SetClr_rdperr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClr_rdperr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClr_rdperr1Mask)
	}
}

const (
	RegisterCcr1FieldClr_rdserr1Shift = 24
	RegisterCcr1FieldClr_rdserr1Mask  = 0x1000000
)

// GetClr_rdserr1 Bank 1 RDSERR1 flag clear bit
func (r *registerCcr1Type) GetClr_rdserr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClr_rdserr1Mask) != 0
}

// SetClr_rdserr1 Bank 1 RDSERR1 flag clear bit
func (r *registerCcr1Type) SetClr_rdserr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClr_rdserr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClr_rdserr1Mask)
	}
}

const (
	RegisterCcr1FieldClr_sneccerr1Shift = 25
	RegisterCcr1FieldClr_sneccerr1Mask  = 0x2000000
)

// GetClr_sneccerr1 Bank 1 SNECCERR1 flag clear bit
func (r *registerCcr1Type) GetClr_sneccerr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClr_sneccerr1Mask) != 0
}

// SetClr_sneccerr1 Bank 1 SNECCERR1 flag clear bit
func (r *registerCcr1Type) SetClr_sneccerr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClr_sneccerr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClr_sneccerr1Mask)
	}
}

const (
	RegisterCcr1FieldClr_dbeccerr1Shift = 26
	RegisterCcr1FieldClr_dbeccerr1Mask  = 0x4000000
)

// GetClr_dbeccerr1 Bank 1 DBECCERR1 flag clear bit
func (r *registerCcr1Type) GetClr_dbeccerr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClr_dbeccerr1Mask) != 0
}

// SetClr_dbeccerr1 Bank 1 DBECCERR1 flag clear bit
func (r *registerCcr1Type) SetClr_dbeccerr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClr_dbeccerr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClr_dbeccerr1Mask)
	}
}

const (
	RegisterCcr1FieldClr_crcend1Shift = 27
	RegisterCcr1FieldClr_crcend1Mask  = 0x8000000
)

// GetClr_crcend1 Bank 1 CRCEND1 flag clear bit
func (r *registerCcr1Type) GetClr_crcend1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr1FieldClr_crcend1Mask) != 0
}

// SetClr_crcend1 Bank 1 CRCEND1 flag clear bit
func (r *registerCcr1Type) SetClr_crcend1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr1FieldClr_crcend1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr1FieldClr_crcend1Mask)
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
	RegisterOptcrFieldSwap_bankShift = 31
	RegisterOptcrFieldSwap_bankMask  = 0x80000000
)

// GetSwap_bank Bank swapping configuration bit
func (r *registerOptcrType) GetSwap_bank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptcrFieldSwap_bankMask) != 0
}

// SetSwap_bank Bank swapping configuration bit
func (r *registerOptcrType) SetSwap_bank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptcrFieldSwap_bankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptcrFieldSwap_bankMask)
	}
}

// registerOptsr_curType FLASH option status register
type registerOptsr_curType uint32

const (
	RegisterOptsr_curFieldOpt_busyShift = 0
	RegisterOptsr_curFieldOpt_busyMask  = 0x1
)

// GetOpt_busy Option byte change ongoing flag
func (r *registerOptsr_curType) GetOpt_busy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldOpt_busyMask) != 0
}

// SetOpt_busy Option byte change ongoing flag
func (r *registerOptsr_curType) SetOpt_busy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldOpt_busyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldOpt_busyMask)
	}
}

const (
	RegisterOptsr_curFieldBor_levShift = 2
	RegisterOptsr_curFieldBor_levMask  = 0xc
)

// GetBor_lev Brownout level option status bit
func (r *registerOptsr_curType) GetBor_lev() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldBor_levMask) >> RegisterOptsr_curFieldBor_levShift)
}

// SetBor_lev Brownout level option status bit
func (r *registerOptsr_curType) SetBor_lev(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldBor_levMask)|(uint32(value)<<RegisterOptsr_curFieldBor_levShift))
}

const (
	RegisterOptsr_curFieldIwdg1_hwShift = 4
	RegisterOptsr_curFieldIwdg1_hwMask  = 0x10
)

// GetIwdg1_hw IWDG1 control option status bit
func (r *registerOptsr_curType) GetIwdg1_hw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldIwdg1_hwMask) != 0
}

// SetIwdg1_hw IWDG1 control option status bit
func (r *registerOptsr_curType) SetIwdg1_hw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldIwdg1_hwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldIwdg1_hwMask)
	}
}

const (
	RegisterOptsr_curFieldNrst_stop_d1Shift = 6
	RegisterOptsr_curFieldNrst_stop_d1Mask  = 0x40
)

// GetNrst_stop_d1 D1 DStop entry reset option status bit
func (r *registerOptsr_curType) GetNrst_stop_d1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldNrst_stop_d1Mask) != 0
}

// SetNrst_stop_d1 D1 DStop entry reset option status bit
func (r *registerOptsr_curType) SetNrst_stop_d1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldNrst_stop_d1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldNrst_stop_d1Mask)
	}
}

const (
	RegisterOptsr_curFieldNrst_stby_d1Shift = 7
	RegisterOptsr_curFieldNrst_stby_d1Mask  = 0x80
)

// GetNrst_stby_d1 D1 DStandby entry reset option status bit
func (r *registerOptsr_curType) GetNrst_stby_d1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldNrst_stby_d1Mask) != 0
}

// SetNrst_stby_d1 D1 DStandby entry reset option status bit
func (r *registerOptsr_curType) SetNrst_stby_d1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldNrst_stby_d1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldNrst_stby_d1Mask)
	}
}

const (
	RegisterOptsr_curFieldRdpShift = 8
	RegisterOptsr_curFieldRdpMask  = 0xff00
)

// GetRdp Readout protection level option status byte
func (r *registerOptsr_curType) GetRdp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldRdpMask) >> RegisterOptsr_curFieldRdpShift)
}

// SetRdp Readout protection level option status byte
func (r *registerOptsr_curType) SetRdp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldRdpMask)|(uint32(value)<<RegisterOptsr_curFieldRdpShift))
}

const (
	RegisterOptsr_curFieldFz_iwdg_stopShift = 17
	RegisterOptsr_curFieldFz_iwdg_stopMask  = 0x20000
)

// GetFz_iwdg_stop IWDG Stop mode freeze option status bit
func (r *registerOptsr_curType) GetFz_iwdg_stop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldFz_iwdg_stopMask) != 0
}

// SetFz_iwdg_stop IWDG Stop mode freeze option status bit
func (r *registerOptsr_curType) SetFz_iwdg_stop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldFz_iwdg_stopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldFz_iwdg_stopMask)
	}
}

const (
	RegisterOptsr_curFieldFz_iwdg_sdbyShift = 18
	RegisterOptsr_curFieldFz_iwdg_sdbyMask  = 0x40000
)

// GetFz_iwdg_sdby IWDG Standby mode freeze option status bit
func (r *registerOptsr_curType) GetFz_iwdg_sdby() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldFz_iwdg_sdbyMask) != 0
}

// SetFz_iwdg_sdby IWDG Standby mode freeze option status bit
func (r *registerOptsr_curType) SetFz_iwdg_sdby(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldFz_iwdg_sdbyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldFz_iwdg_sdbyMask)
	}
}

const (
	RegisterOptsr_curFieldSt_ram_sizeShift = 19
	RegisterOptsr_curFieldSt_ram_sizeMask  = 0x180000
)

// GetSt_ram_size DTCM RAM size option status
func (r *registerOptsr_curType) GetSt_ram_size() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldSt_ram_sizeMask) >> RegisterOptsr_curFieldSt_ram_sizeShift)
}

// SetSt_ram_size DTCM RAM size option status
func (r *registerOptsr_curType) SetSt_ram_size(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldSt_ram_sizeMask)|(uint32(value)<<RegisterOptsr_curFieldSt_ram_sizeShift))
}

const (
	RegisterOptsr_curFieldSecurityShift = 21
	RegisterOptsr_curFieldSecurityMask  = 0x200000
)

// GetSecurity Security enable option status bit
func (r *registerOptsr_curType) GetSecurity() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldSecurityMask) != 0
}

// SetSecurity Security enable option status bit
func (r *registerOptsr_curType) SetSecurity(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldSecurityMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldSecurityMask)
	}
}

const (
	RegisterOptsr_curFieldRss1Shift = 26
	RegisterOptsr_curFieldRss1Mask  = 0x4000000
)

// GetRss1 User option bit 1
func (r *registerOptsr_curType) GetRss1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldRss1Mask) != 0
}

// SetRss1 User option bit 1
func (r *registerOptsr_curType) SetRss1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldRss1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldRss1Mask)
	}
}

const (
	RegisterOptsr_curFieldPerso_okShift = 28
	RegisterOptsr_curFieldPerso_okMask  = 0x10000000
)

// GetPerso_ok Device personalization status bit
func (r *registerOptsr_curType) GetPerso_ok() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldPerso_okMask) != 0
}

// SetPerso_ok Device personalization status bit
func (r *registerOptsr_curType) SetPerso_ok(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldPerso_okMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldPerso_okMask)
	}
}

const (
	RegisterOptsr_curFieldIo_hslvShift = 29
	RegisterOptsr_curFieldIo_hslvMask  = 0x20000000
)

// GetIo_hslv I/O high-speed at low-voltage status bit (PRODUCT_BELOW_25V)
func (r *registerOptsr_curType) GetIo_hslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldIo_hslvMask) != 0
}

// SetIo_hslv I/O high-speed at low-voltage status bit (PRODUCT_BELOW_25V)
func (r *registerOptsr_curType) SetIo_hslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldIo_hslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldIo_hslvMask)
	}
}

const (
	RegisterOptsr_curFieldOptchangeerrShift = 30
	RegisterOptsr_curFieldOptchangeerrMask  = 0x40000000
)

// GetOptchangeerr Option byte change error flag
func (r *registerOptsr_curType) GetOptchangeerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldOptchangeerrMask) != 0
}

// SetOptchangeerr Option byte change error flag
func (r *registerOptsr_curType) SetOptchangeerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldOptchangeerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldOptchangeerrMask)
	}
}

const (
	RegisterOptsr_curFieldSwap_bank_optShift = 31
	RegisterOptsr_curFieldSwap_bank_optMask  = 0x80000000
)

// GetSwap_bank_opt Bank swapping option status bit
func (r *registerOptsr_curType) GetSwap_bank_opt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldSwap_bank_optMask) != 0
}

// SetSwap_bank_opt Bank swapping option status bit
func (r *registerOptsr_curType) SetSwap_bank_opt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldSwap_bank_optMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldSwap_bank_optMask)
	}
}

// registerOptsr_prgType FLASH option status register
type registerOptsr_prgType uint32

const (
	RegisterOptsr_prgFieldBor_levShift = 2
	RegisterOptsr_prgFieldBor_levMask  = 0xc
)

// GetBor_lev BOR reset level option configuration bits
func (r *registerOptsr_prgType) GetBor_lev() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldBor_levMask) >> RegisterOptsr_prgFieldBor_levShift)
}

// SetBor_lev BOR reset level option configuration bits
func (r *registerOptsr_prgType) SetBor_lev(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldBor_levMask)|(uint32(value)<<RegisterOptsr_prgFieldBor_levShift))
}

const (
	RegisterOptsr_prgFieldIwdg1_hwShift = 4
	RegisterOptsr_prgFieldIwdg1_hwMask  = 0x10
)

// GetIwdg1_hw IWDG1 option configuration bit
func (r *registerOptsr_prgType) GetIwdg1_hw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldIwdg1_hwMask) != 0
}

// SetIwdg1_hw IWDG1 option configuration bit
func (r *registerOptsr_prgType) SetIwdg1_hw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldIwdg1_hwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldIwdg1_hwMask)
	}
}

const (
	RegisterOptsr_prgFieldNrst_stop_d1Shift = 6
	RegisterOptsr_prgFieldNrst_stop_d1Mask  = 0x40
)

// GetNrst_stop_d1 Option byte erase after D1 DStop option configuration bit
func (r *registerOptsr_prgType) GetNrst_stop_d1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldNrst_stop_d1Mask) != 0
}

// SetNrst_stop_d1 Option byte erase after D1 DStop option configuration bit
func (r *registerOptsr_prgType) SetNrst_stop_d1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldNrst_stop_d1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldNrst_stop_d1Mask)
	}
}

const (
	RegisterOptsr_prgFieldNrst_stby_d1Shift = 7
	RegisterOptsr_prgFieldNrst_stby_d1Mask  = 0x80
)

// GetNrst_stby_d1 Option byte erase after D1 DStandby option configuration bit
func (r *registerOptsr_prgType) GetNrst_stby_d1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldNrst_stby_d1Mask) != 0
}

// SetNrst_stby_d1 Option byte erase after D1 DStandby option configuration bit
func (r *registerOptsr_prgType) SetNrst_stby_d1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldNrst_stby_d1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldNrst_stby_d1Mask)
	}
}

const (
	RegisterOptsr_prgFieldRdpShift = 8
	RegisterOptsr_prgFieldRdpMask  = 0xff00
)

// GetRdp Readout protection level option configuration byte
func (r *registerOptsr_prgType) GetRdp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldRdpMask) >> RegisterOptsr_prgFieldRdpShift)
}

// SetRdp Readout protection level option configuration byte
func (r *registerOptsr_prgType) SetRdp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldRdpMask)|(uint32(value)<<RegisterOptsr_prgFieldRdpShift))
}

const (
	RegisterOptsr_prgFieldFz_iwdg_stopShift = 17
	RegisterOptsr_prgFieldFz_iwdg_stopMask  = 0x20000
)

// GetFz_iwdg_stop IWDG Stop mode freeze option configuration bit
func (r *registerOptsr_prgType) GetFz_iwdg_stop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldFz_iwdg_stopMask) != 0
}

// SetFz_iwdg_stop IWDG Stop mode freeze option configuration bit
func (r *registerOptsr_prgType) SetFz_iwdg_stop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldFz_iwdg_stopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldFz_iwdg_stopMask)
	}
}

const (
	RegisterOptsr_prgFieldFz_iwdg_sdbyShift = 18
	RegisterOptsr_prgFieldFz_iwdg_sdbyMask  = 0x40000
)

// GetFz_iwdg_sdby IWDG Standby mode freeze option configuration bit
func (r *registerOptsr_prgType) GetFz_iwdg_sdby() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldFz_iwdg_sdbyMask) != 0
}

// SetFz_iwdg_sdby IWDG Standby mode freeze option configuration bit
func (r *registerOptsr_prgType) SetFz_iwdg_sdby(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldFz_iwdg_sdbyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldFz_iwdg_sdbyMask)
	}
}

const (
	RegisterOptsr_prgFieldSt_ram_sizeShift = 19
	RegisterOptsr_prgFieldSt_ram_sizeMask  = 0x180000
)

// GetSt_ram_size DTCM size select option configuration bits
func (r *registerOptsr_prgType) GetSt_ram_size() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldSt_ram_sizeMask) >> RegisterOptsr_prgFieldSt_ram_sizeShift)
}

// SetSt_ram_size DTCM size select option configuration bits
func (r *registerOptsr_prgType) SetSt_ram_size(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldSt_ram_sizeMask)|(uint32(value)<<RegisterOptsr_prgFieldSt_ram_sizeShift))
}

const (
	RegisterOptsr_prgFieldSecurityShift = 21
	RegisterOptsr_prgFieldSecurityMask  = 0x200000
)

// GetSecurity Security option configuration bit
func (r *registerOptsr_prgType) GetSecurity() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldSecurityMask) != 0
}

// SetSecurity Security option configuration bit
func (r *registerOptsr_prgType) SetSecurity(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldSecurityMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldSecurityMask)
	}
}

const (
	RegisterOptsr_prgFieldRss1Shift = 26
	RegisterOptsr_prgFieldRss1Mask  = 0x4000000
)

// GetRss1 User option configuration bit 1
func (r *registerOptsr_prgType) GetRss1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldRss1Mask) != 0
}

// SetRss1 User option configuration bit 1
func (r *registerOptsr_prgType) SetRss1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldRss1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldRss1Mask)
	}
}

const (
	RegisterOptsr_prgFieldRss2Shift = 27
	RegisterOptsr_prgFieldRss2Mask  = 0x8000000
)

// GetRss2 User option configuration bit 2
func (r *registerOptsr_prgType) GetRss2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldRss2Mask) != 0
}

// SetRss2 User option configuration bit 2
func (r *registerOptsr_prgType) SetRss2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldRss2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldRss2Mask)
	}
}

const (
	RegisterOptsr_prgFieldIo_hslvShift = 29
	RegisterOptsr_prgFieldIo_hslvMask  = 0x20000000
)

// GetIo_hslv I/O high-speed at low-voltage (PRODUCT_BELOW_25V)
func (r *registerOptsr_prgType) GetIo_hslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldIo_hslvMask) != 0
}

// SetIo_hslv I/O high-speed at low-voltage (PRODUCT_BELOW_25V)
func (r *registerOptsr_prgType) SetIo_hslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldIo_hslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldIo_hslvMask)
	}
}

const (
	RegisterOptsr_prgFieldSwap_bank_optShift = 31
	RegisterOptsr_prgFieldSwap_bank_optMask  = 0x80000000
)

// GetSwap_bank_opt Bank swapping option configuration bit
func (r *registerOptsr_prgType) GetSwap_bank_opt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldSwap_bank_optMask) != 0
}

// SetSwap_bank_opt Bank swapping option configuration bit
func (r *registerOptsr_prgType) SetSwap_bank_opt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldSwap_bank_optMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldSwap_bank_optMask)
	}
}

// registerOptccrType FLASH option clear control register
type registerOptccrType uint32

const (
	RegisterOptccrFieldClr_optchangeerrShift = 30
	RegisterOptccrFieldClr_optchangeerrMask  = 0x40000000
)

// GetClr_optchangeerr OPTCHANGEERR reset bit
func (r *registerOptccrType) GetClr_optchangeerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptccrFieldClr_optchangeerrMask) != 0
}

// SetClr_optchangeerr OPTCHANGEERR reset bit
func (r *registerOptccrType) SetClr_optchangeerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptccrFieldClr_optchangeerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptccrFieldClr_optchangeerrMask)
	}
}

// registerPrar_cur1Type FLASH protection address for bank 1
type registerPrar_cur1Type uint32

const (
	RegisterPrar_cur1FieldProt_area_start1Shift = 0
	RegisterPrar_cur1FieldProt_area_start1Mask  = 0xfff
)

// GetProt_area_start1 Bank 1 lowest PCROP protected address
func (r *registerPrar_cur1Type) GetProt_area_start1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrar_cur1FieldProt_area_start1Mask) >> RegisterPrar_cur1FieldProt_area_start1Shift)
}

// SetProt_area_start1 Bank 1 lowest PCROP protected address
func (r *registerPrar_cur1Type) SetProt_area_start1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrar_cur1FieldProt_area_start1Mask)|(uint32(value)<<RegisterPrar_cur1FieldProt_area_start1Shift))
}

const (
	RegisterPrar_cur1FieldProt_area_end1Shift = 16
	RegisterPrar_cur1FieldProt_area_end1Mask  = 0xfff0000
)

// GetProt_area_end1 Bank 1 highest PCROP protected address
func (r *registerPrar_cur1Type) GetProt_area_end1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrar_cur1FieldProt_area_end1Mask) >> RegisterPrar_cur1FieldProt_area_end1Shift)
}

// SetProt_area_end1 Bank 1 highest PCROP protected address
func (r *registerPrar_cur1Type) SetProt_area_end1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrar_cur1FieldProt_area_end1Mask)|(uint32(value)<<RegisterPrar_cur1FieldProt_area_end1Shift))
}

const (
	RegisterPrar_cur1FieldDmep1Shift = 31
	RegisterPrar_cur1FieldDmep1Mask  = 0x80000000
)

// GetDmep1 Bank 1 PCROP protected erase enable option status bit
func (r *registerPrar_cur1Type) GetDmep1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrar_cur1FieldDmep1Mask) != 0
}

// SetDmep1 Bank 1 PCROP protected erase enable option status bit
func (r *registerPrar_cur1Type) SetDmep1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrar_cur1FieldDmep1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrar_cur1FieldDmep1Mask)
	}
}

// registerPrar_prg2Type FLASH protection address for bank 2
type registerPrar_prg2Type uint32

const (
	RegisterPrar_prg2FieldProt_area_start2Shift = 0
	RegisterPrar_prg2FieldProt_area_start2Mask  = 0xfff
)

// GetProt_area_start2 Bank 2 lowest PCROP protected address configuration
func (r *registerPrar_prg2Type) GetProt_area_start2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrar_prg2FieldProt_area_start2Mask) >> RegisterPrar_prg2FieldProt_area_start2Shift)
}

// SetProt_area_start2 Bank 2 lowest PCROP protected address configuration
func (r *registerPrar_prg2Type) SetProt_area_start2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrar_prg2FieldProt_area_start2Mask)|(uint32(value)<<RegisterPrar_prg2FieldProt_area_start2Shift))
}

const (
	RegisterPrar_prg2FieldProt_area_end2Shift = 16
	RegisterPrar_prg2FieldProt_area_end2Mask  = 0xfff0000
)

// GetProt_area_end2 Bank 2 highest PCROP protected address configuration
func (r *registerPrar_prg2Type) GetProt_area_end2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrar_prg2FieldProt_area_end2Mask) >> RegisterPrar_prg2FieldProt_area_end2Shift)
}

// SetProt_area_end2 Bank 2 highest PCROP protected address configuration
func (r *registerPrar_prg2Type) SetProt_area_end2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrar_prg2FieldProt_area_end2Mask)|(uint32(value)<<RegisterPrar_prg2FieldProt_area_end2Shift))
}

const (
	RegisterPrar_prg2FieldDmep2Shift = 31
	RegisterPrar_prg2FieldDmep2Mask  = 0x80000000
)

// GetDmep2 Bank 2 PCROP protected erase enable option configuration bit
func (r *registerPrar_prg2Type) GetDmep2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrar_prg2FieldDmep2Mask) != 0
}

// SetDmep2 Bank 2 PCROP protected erase enable option configuration bit
func (r *registerPrar_prg2Type) SetDmep2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrar_prg2FieldDmep2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrar_prg2FieldDmep2Mask)
	}
}

// registerPrar_prg1Type FLASH protection address for bank 1
type registerPrar_prg1Type uint32

const (
	RegisterPrar_prg1FieldProt_area_start1Shift = 0
	RegisterPrar_prg1FieldProt_area_start1Mask  = 0xfff
)

// GetProt_area_start1 Bank 1 lowest PCROP protected address configuration
func (r *registerPrar_prg1Type) GetProt_area_start1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrar_prg1FieldProt_area_start1Mask) >> RegisterPrar_prg1FieldProt_area_start1Shift)
}

// SetProt_area_start1 Bank 1 lowest PCROP protected address configuration
func (r *registerPrar_prg1Type) SetProt_area_start1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrar_prg1FieldProt_area_start1Mask)|(uint32(value)<<RegisterPrar_prg1FieldProt_area_start1Shift))
}

const (
	RegisterPrar_prg1FieldProt_area_end1Shift = 16
	RegisterPrar_prg1FieldProt_area_end1Mask  = 0xfff0000
)

// GetProt_area_end1 Bank 1 highest PCROP protected address configuration
func (r *registerPrar_prg1Type) GetProt_area_end1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrar_prg1FieldProt_area_end1Mask) >> RegisterPrar_prg1FieldProt_area_end1Shift)
}

// SetProt_area_end1 Bank 1 highest PCROP protected address configuration
func (r *registerPrar_prg1Type) SetProt_area_end1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrar_prg1FieldProt_area_end1Mask)|(uint32(value)<<RegisterPrar_prg1FieldProt_area_end1Shift))
}

const (
	RegisterPrar_prg1FieldDmep1Shift = 31
	RegisterPrar_prg1FieldDmep1Mask  = 0x80000000
)

// GetDmep1 Bank 1 PCROP protected erase enable option configuration bit
func (r *registerPrar_prg1Type) GetDmep1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrar_prg1FieldDmep1Mask) != 0
}

// SetDmep1 Bank 1 PCROP protected erase enable option configuration bit
func (r *registerPrar_prg1Type) SetDmep1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrar_prg1FieldDmep1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrar_prg1FieldDmep1Mask)
	}
}

// registerScar_cur1Type FLASH secure address for bank 1
type registerScar_cur1Type uint32

const (
	RegisterScar_cur1FieldSec_area_start1Shift = 0
	RegisterScar_cur1FieldSec_area_start1Mask  = 0xfff
)

// GetSec_area_start1 Bank 1 lowest secure protected address
func (r *registerScar_cur1Type) GetSec_area_start1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScar_cur1FieldSec_area_start1Mask) >> RegisterScar_cur1FieldSec_area_start1Shift)
}

// SetSec_area_start1 Bank 1 lowest secure protected address
func (r *registerScar_cur1Type) SetSec_area_start1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScar_cur1FieldSec_area_start1Mask)|(uint32(value)<<RegisterScar_cur1FieldSec_area_start1Shift))
}

const (
	RegisterScar_cur1FieldSec_area_end1Shift = 16
	RegisterScar_cur1FieldSec_area_end1Mask  = 0xfff0000
)

// GetSec_area_end1 Bank 1 highest secure protected address
func (r *registerScar_cur1Type) GetSec_area_end1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScar_cur1FieldSec_area_end1Mask) >> RegisterScar_cur1FieldSec_area_end1Shift)
}

// SetSec_area_end1 Bank 1 highest secure protected address
func (r *registerScar_cur1Type) SetSec_area_end1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScar_cur1FieldSec_area_end1Mask)|(uint32(value)<<RegisterScar_cur1FieldSec_area_end1Shift))
}

const (
	RegisterScar_cur1FieldDmes1Shift = 31
	RegisterScar_cur1FieldDmes1Mask  = 0x80000000
)

// GetDmes1 Bank 1 secure protected erase enable option status bit
func (r *registerScar_cur1Type) GetDmes1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScar_cur1FieldDmes1Mask) != 0
}

// SetDmes1 Bank 1 secure protected erase enable option status bit
func (r *registerScar_cur1Type) SetDmes1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScar_cur1FieldDmes1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScar_cur1FieldDmes1Mask)
	}
}

// registerScar_prg1Type FLASH secure address for bank 1
type registerScar_prg1Type uint32

const (
	RegisterScar_prg1FieldSec_area_start1Shift = 0
	RegisterScar_prg1FieldSec_area_start1Mask  = 0xfff
)

// GetSec_area_start1 Bank 1 lowest secure protected address configuration
func (r *registerScar_prg1Type) GetSec_area_start1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScar_prg1FieldSec_area_start1Mask) >> RegisterScar_prg1FieldSec_area_start1Shift)
}

// SetSec_area_start1 Bank 1 lowest secure protected address configuration
func (r *registerScar_prg1Type) SetSec_area_start1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScar_prg1FieldSec_area_start1Mask)|(uint32(value)<<RegisterScar_prg1FieldSec_area_start1Shift))
}

const (
	RegisterScar_prg1FieldSec_area_end1Shift = 16
	RegisterScar_prg1FieldSec_area_end1Mask  = 0xfff0000
)

// GetSec_area_end1 Bank 1 highest secure protected address configuration
func (r *registerScar_prg1Type) GetSec_area_end1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScar_prg1FieldSec_area_end1Mask) >> RegisterScar_prg1FieldSec_area_end1Shift)
}

// SetSec_area_end1 Bank 1 highest secure protected address configuration
func (r *registerScar_prg1Type) SetSec_area_end1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScar_prg1FieldSec_area_end1Mask)|(uint32(value)<<RegisterScar_prg1FieldSec_area_end1Shift))
}

const (
	RegisterScar_prg1FieldDmes1Shift = 31
	RegisterScar_prg1FieldDmes1Mask  = 0x80000000
)

// GetDmes1 Bank 1 secure protected erase enable option configuration bit
func (r *registerScar_prg1Type) GetDmes1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScar_prg1FieldDmes1Mask) != 0
}

// SetDmes1 Bank 1 secure protected erase enable option configuration bit
func (r *registerScar_prg1Type) SetDmes1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScar_prg1FieldDmes1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScar_prg1FieldDmes1Mask)
	}
}

// registerWpsn_cur1rType FLASH write sector protection for bank 1
type registerWpsn_cur1rType uint32

const (
	RegisterWpsn_cur1rFieldWrpsn1Shift = 0
	RegisterWpsn_cur1rFieldWrpsn1Mask  = 0xff
)

// GetWrpsn1 Bank 1 sector write protection option status byte
func (r *registerWpsn_cur1rType) GetWrpsn1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWpsn_cur1rFieldWrpsn1Mask) >> RegisterWpsn_cur1rFieldWrpsn1Shift)
}

// SetWrpsn1 Bank 1 sector write protection option status byte
func (r *registerWpsn_cur1rType) SetWrpsn1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWpsn_cur1rFieldWrpsn1Mask)|(uint32(value)<<RegisterWpsn_cur1rFieldWrpsn1Shift))
}

// registerWpsn_prg1rType FLASH write sector protection for bank 1
type registerWpsn_prg1rType uint32

const (
	RegisterWpsn_prg1rFieldWrpsn1Shift = 0
	RegisterWpsn_prg1rFieldWrpsn1Mask  = 0xff
)

// GetWrpsn1 Bank 1 sector write protection configuration byte
func (r *registerWpsn_prg1rType) GetWrpsn1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWpsn_prg1rFieldWrpsn1Mask) >> RegisterWpsn_prg1rFieldWrpsn1Shift)
}

// SetWrpsn1 Bank 1 sector write protection configuration byte
func (r *registerWpsn_prg1rType) SetWrpsn1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWpsn_prg1rFieldWrpsn1Mask)|(uint32(value)<<RegisterWpsn_prg1rFieldWrpsn1Shift))
}

// registerBoot_currType FLASH register with boot address
type registerBoot_currType uint32

const (
	RegisterBoot_currFieldBoot_add0Shift = 0
	RegisterBoot_currFieldBoot_add0Mask  = 0xffff
)

// GetBoot_add0 Boot address 0
func (r *registerBoot_currType) GetBoot_add0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBoot_currFieldBoot_add0Mask) >> RegisterBoot_currFieldBoot_add0Shift)
}

// SetBoot_add0 Boot address 0
func (r *registerBoot_currType) SetBoot_add0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBoot_currFieldBoot_add0Mask)|(uint32(value)<<RegisterBoot_currFieldBoot_add0Shift))
}

const (
	RegisterBoot_currFieldBoot_add1Shift = 16
	RegisterBoot_currFieldBoot_add1Mask  = 0xffff0000
)

// GetBoot_add1 Boot address 1
func (r *registerBoot_currType) GetBoot_add1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBoot_currFieldBoot_add1Mask) >> RegisterBoot_currFieldBoot_add1Shift)
}

// SetBoot_add1 Boot address 1
func (r *registerBoot_currType) SetBoot_add1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBoot_currFieldBoot_add1Mask)|(uint32(value)<<RegisterBoot_currFieldBoot_add1Shift))
}

// registerBoot_prgrType FLASH register with boot address
type registerBoot_prgrType uint32

const (
	RegisterBoot_prgrFieldBoot_add0Shift = 0
	RegisterBoot_prgrFieldBoot_add0Mask  = 0xffff
)

// GetBoot_add0 Boot address 0
func (r *registerBoot_prgrType) GetBoot_add0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBoot_prgrFieldBoot_add0Mask) >> RegisterBoot_prgrFieldBoot_add0Shift)
}

// SetBoot_add0 Boot address 0
func (r *registerBoot_prgrType) SetBoot_add0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBoot_prgrFieldBoot_add0Mask)|(uint32(value)<<RegisterBoot_prgrFieldBoot_add0Shift))
}

const (
	RegisterBoot_prgrFieldBoot_add1Shift = 16
	RegisterBoot_prgrFieldBoot_add1Mask  = 0xffff0000
)

// GetBoot_add1 Boot address 1
func (r *registerBoot_prgrType) GetBoot_add1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBoot_prgrFieldBoot_add1Mask) >> RegisterBoot_prgrFieldBoot_add1Shift)
}

// SetBoot_add1 Boot address 1
func (r *registerBoot_prgrType) SetBoot_add1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBoot_prgrFieldBoot_add1Mask)|(uint32(value)<<RegisterBoot_prgrFieldBoot_add1Shift))
}

// registerCrccr1Type FLASH CRC control register for bank 1
type registerCrccr1Type uint32

const (
	RegisterCrccr1FieldCrc_sectShift = 0
	RegisterCrccr1FieldCrc_sectMask  = 0x7
)

// GetCrc_sect Bank 1 CRC sector number
func (r *registerCrccr1Type) GetCrc_sect() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldCrc_sectMask) >> RegisterCrccr1FieldCrc_sectShift)
}

// SetCrc_sect Bank 1 CRC sector number
func (r *registerCrccr1Type) SetCrc_sect(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldCrc_sectMask)|(uint32(value)<<RegisterCrccr1FieldCrc_sectShift))
}

const (
	RegisterCrccr1FieldAll_bankShift = 7
	RegisterCrccr1FieldAll_bankMask  = 0x80
)

// GetAll_bank Bank 1 CRC select bit
func (r *registerCrccr1Type) GetAll_bank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldAll_bankMask) != 0
}

// SetAll_bank Bank 1 CRC select bit
func (r *registerCrccr1Type) SetAll_bank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr1FieldAll_bankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldAll_bankMask)
	}
}

const (
	RegisterCrccr1FieldCrc_by_sectShift = 8
	RegisterCrccr1FieldCrc_by_sectMask  = 0x100
)

// GetCrc_by_sect Bank 1 CRC sector mode select bit
func (r *registerCrccr1Type) GetCrc_by_sect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldCrc_by_sectMask) != 0
}

// SetCrc_by_sect Bank 1 CRC sector mode select bit
func (r *registerCrccr1Type) SetCrc_by_sect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr1FieldCrc_by_sectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldCrc_by_sectMask)
	}
}

const (
	RegisterCrccr1FieldAdd_sectShift = 9
	RegisterCrccr1FieldAdd_sectMask  = 0x200
)

// GetAdd_sect Bank 1 CRC sector select bit
func (r *registerCrccr1Type) GetAdd_sect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldAdd_sectMask) != 0
}

// SetAdd_sect Bank 1 CRC sector select bit
func (r *registerCrccr1Type) SetAdd_sect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr1FieldAdd_sectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldAdd_sectMask)
	}
}

const (
	RegisterCrccr1FieldClean_sectShift = 10
	RegisterCrccr1FieldClean_sectMask  = 0x400
)

// GetClean_sect Bank 1 CRC sector list clear bit
func (r *registerCrccr1Type) GetClean_sect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldClean_sectMask) != 0
}

// SetClean_sect Bank 1 CRC sector list clear bit
func (r *registerCrccr1Type) SetClean_sect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr1FieldClean_sectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldClean_sectMask)
	}
}

const (
	RegisterCrccr1FieldStart_crcShift = 16
	RegisterCrccr1FieldStart_crcMask  = 0x10000
)

// GetStart_crc Bank 1 CRC start bit
func (r *registerCrccr1Type) GetStart_crc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldStart_crcMask) != 0
}

// SetStart_crc Bank 1 CRC start bit
func (r *registerCrccr1Type) SetStart_crc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr1FieldStart_crcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldStart_crcMask)
	}
}

const (
	RegisterCrccr1FieldClean_crcShift = 17
	RegisterCrccr1FieldClean_crcMask  = 0x20000
)

// GetClean_crc Bank 1 CRC clear bit
func (r *registerCrccr1Type) GetClean_crc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldClean_crcMask) != 0
}

// SetClean_crc Bank 1 CRC clear bit
func (r *registerCrccr1Type) SetClean_crc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr1FieldClean_crcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldClean_crcMask)
	}
}

const (
	RegisterCrccr1FieldCrc_burstShift = 20
	RegisterCrccr1FieldCrc_burstMask  = 0x300000
)

// GetCrc_burst Bank 1 CRC burst size
func (r *registerCrccr1Type) GetCrc_burst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrccr1FieldCrc_burstMask) >> RegisterCrccr1FieldCrc_burstShift)
}

// SetCrc_burst Bank 1 CRC burst size
func (r *registerCrccr1Type) SetCrc_burst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrccr1FieldCrc_burstMask)|(uint32(value)<<RegisterCrccr1FieldCrc_burstShift))
}

// registerCrcsadd1rType FLASH CRC start address register for bank 1
type registerCrcsadd1rType uint32

const (
	RegisterCrcsadd1rFieldCrc_start_addrShift = 0
	RegisterCrcsadd1rFieldCrc_start_addrMask  = 0xffffffff
)

// GetCrc_start_addr CRC start address on bank 1
func (r *registerCrcsadd1rType) GetCrc_start_addr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrcsadd1rFieldCrc_start_addrMask) >> RegisterCrcsadd1rFieldCrc_start_addrShift)
}

// SetCrc_start_addr CRC start address on bank 1
func (r *registerCrcsadd1rType) SetCrc_start_addr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrcsadd1rFieldCrc_start_addrMask)|(uint32(value)<<RegisterCrcsadd1rFieldCrc_start_addrShift))
}

// registerCrceadd1rType FLASH CRC end address register for bank 1
type registerCrceadd1rType uint32

const (
	RegisterCrceadd1rFieldCrc_end_addrShift = 0
	RegisterCrceadd1rFieldCrc_end_addrMask  = 0xffffffff
)

// GetCrc_end_addr CRC end address on bank 1
func (r *registerCrceadd1rType) GetCrc_end_addr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrceadd1rFieldCrc_end_addrMask) >> RegisterCrceadd1rFieldCrc_end_addrShift)
}

// SetCrc_end_addr CRC end address on bank 1
func (r *registerCrceadd1rType) SetCrc_end_addr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrceadd1rFieldCrc_end_addrMask)|(uint32(value)<<RegisterCrceadd1rFieldCrc_end_addrShift))
}

// registerCrcdatarType FLASH CRC data register
type registerCrcdatarType uint32

const (
	RegisterCrcdatarFieldCrc_dataShift = 0
	RegisterCrcdatarFieldCrc_dataMask  = 0xffffffff
)

// GetCrc_data CRC result
func (r *registerCrcdatarType) GetCrc_data() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrcdatarFieldCrc_dataMask) >> RegisterCrcdatarFieldCrc_dataShift)
}

// SetCrc_data CRC result
func (r *registerCrcdatarType) SetCrc_data(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrcdatarFieldCrc_dataMask)|(uint32(value)<<RegisterCrcdatarFieldCrc_dataShift))
}

// registerEcc_fa1rType FLASH ECC fail address for bank 1
type registerEcc_fa1rType uint32

const (
	RegisterEcc_fa1rFieldFail_ecc_addr1Shift = 0
	RegisterEcc_fa1rFieldFail_ecc_addr1Mask  = 0x7fff
)

// GetFail_ecc_addr1 Bank 1 ECC error address
func (r *registerEcc_fa1rType) GetFail_ecc_addr1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterEcc_fa1rFieldFail_ecc_addr1Mask) >> RegisterEcc_fa1rFieldFail_ecc_addr1Shift)
}

// SetFail_ecc_addr1 Bank 1 ECC error address
func (r *registerEcc_fa1rType) SetFail_ecc_addr1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEcc_fa1rFieldFail_ecc_addr1Mask)|(uint32(value)<<RegisterEcc_fa1rFieldFail_ecc_addr1Shift))
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
	RegisterCr2FieldCrc_enShift = 15
	RegisterCr2FieldCrc_enMask  = 0x8000
)

// GetCrc_en Bank 2 CRC control bit
func (r *registerCr2Type) GetCrc_en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCrc_enMask) != 0
}

// SetCrc_en Bank 2 CRC control bit
func (r *registerCr2Type) SetCrc_en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldCrc_enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldCrc_enMask)
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
	RegisterSr2FieldCrc_busy2Shift = 3
	RegisterSr2FieldCrc_busy2Mask  = 0x8
)

// GetCrc_busy2 Bank 2 CRC busy flag
func (r *registerSr2Type) GetCrc_busy2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSr2FieldCrc_busy2Mask) != 0
}

// SetCrc_busy2 Bank 2 CRC busy flag
func (r *registerSr2Type) SetCrc_busy2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSr2FieldCrc_busy2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSr2FieldCrc_busy2Mask)
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
	RegisterCcr2FieldClr_eop2Shift = 16
	RegisterCcr2FieldClr_eop2Mask  = 0x10000
)

// GetClr_eop2 Bank 1 EOP1 flag clear bit
func (r *registerCcr2Type) GetClr_eop2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClr_eop2Mask) != 0
}

// SetClr_eop2 Bank 1 EOP1 flag clear bit
func (r *registerCcr2Type) SetClr_eop2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClr_eop2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClr_eop2Mask)
	}
}

const (
	RegisterCcr2FieldClr_wrperr2Shift = 17
	RegisterCcr2FieldClr_wrperr2Mask  = 0x20000
)

// GetClr_wrperr2 Bank 2 WRPERR1 flag clear bit
func (r *registerCcr2Type) GetClr_wrperr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClr_wrperr2Mask) != 0
}

// SetClr_wrperr2 Bank 2 WRPERR1 flag clear bit
func (r *registerCcr2Type) SetClr_wrperr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClr_wrperr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClr_wrperr2Mask)
	}
}

const (
	RegisterCcr2FieldClr_pgserr2Shift = 18
	RegisterCcr2FieldClr_pgserr2Mask  = 0x40000
)

// GetClr_pgserr2 Bank 2 PGSERR1 flag clear bi
func (r *registerCcr2Type) GetClr_pgserr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClr_pgserr2Mask) != 0
}

// SetClr_pgserr2 Bank 2 PGSERR1 flag clear bi
func (r *registerCcr2Type) SetClr_pgserr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClr_pgserr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClr_pgserr2Mask)
	}
}

const (
	RegisterCcr2FieldClr_strberr2Shift = 19
	RegisterCcr2FieldClr_strberr2Mask  = 0x80000
)

// GetClr_strberr2 Bank 2 STRBERR1 flag clear bit
func (r *registerCcr2Type) GetClr_strberr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClr_strberr2Mask) != 0
}

// SetClr_strberr2 Bank 2 STRBERR1 flag clear bit
func (r *registerCcr2Type) SetClr_strberr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClr_strberr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClr_strberr2Mask)
	}
}

const (
	RegisterCcr2FieldClr_incerr2Shift = 21
	RegisterCcr2FieldClr_incerr2Mask  = 0x200000
)

// GetClr_incerr2 Bank 2 INCERR1 flag clear bit
func (r *registerCcr2Type) GetClr_incerr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClr_incerr2Mask) != 0
}

// SetClr_incerr2 Bank 2 INCERR1 flag clear bit
func (r *registerCcr2Type) SetClr_incerr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClr_incerr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClr_incerr2Mask)
	}
}

const (
	RegisterCcr2FieldClr_operr2Shift = 22
	RegisterCcr2FieldClr_operr2Mask  = 0x400000
)

// GetClr_operr2 Bank 2 OPERR1 flag clear bit
func (r *registerCcr2Type) GetClr_operr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClr_operr2Mask) != 0
}

// SetClr_operr2 Bank 2 OPERR1 flag clear bit
func (r *registerCcr2Type) SetClr_operr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClr_operr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClr_operr2Mask)
	}
}

const (
	RegisterCcr2FieldClr_rdperr2Shift = 23
	RegisterCcr2FieldClr_rdperr2Mask  = 0x800000
)

// GetClr_rdperr2 Bank 2 RDPERR1 flag clear bit
func (r *registerCcr2Type) GetClr_rdperr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClr_rdperr2Mask) != 0
}

// SetClr_rdperr2 Bank 2 RDPERR1 flag clear bit
func (r *registerCcr2Type) SetClr_rdperr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClr_rdperr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClr_rdperr2Mask)
	}
}

const (
	RegisterCcr2FieldClr_rdserr1Shift = 24
	RegisterCcr2FieldClr_rdserr1Mask  = 0x1000000
)

// GetClr_rdserr1 Bank 1 RDSERR1 flag clear bit
func (r *registerCcr2Type) GetClr_rdserr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClr_rdserr1Mask) != 0
}

// SetClr_rdserr1 Bank 1 RDSERR1 flag clear bit
func (r *registerCcr2Type) SetClr_rdserr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClr_rdserr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClr_rdserr1Mask)
	}
}

const (
	RegisterCcr2FieldClr_sneccerr2Shift = 25
	RegisterCcr2FieldClr_sneccerr2Mask  = 0x2000000
)

// GetClr_sneccerr2 Bank 2 SNECCERR1 flag clear bit
func (r *registerCcr2Type) GetClr_sneccerr2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClr_sneccerr2Mask) != 0
}

// SetClr_sneccerr2 Bank 2 SNECCERR1 flag clear bit
func (r *registerCcr2Type) SetClr_sneccerr2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClr_sneccerr2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClr_sneccerr2Mask)
	}
}

const (
	RegisterCcr2FieldClr_dbeccerr1Shift = 26
	RegisterCcr2FieldClr_dbeccerr1Mask  = 0x4000000
)

// GetClr_dbeccerr1 Bank 1 DBECCERR1 flag clear bit
func (r *registerCcr2Type) GetClr_dbeccerr1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClr_dbeccerr1Mask) != 0
}

// SetClr_dbeccerr1 Bank 1 DBECCERR1 flag clear bit
func (r *registerCcr2Type) SetClr_dbeccerr1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClr_dbeccerr1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClr_dbeccerr1Mask)
	}
}

const (
	RegisterCcr2FieldClr_crcend2Shift = 27
	RegisterCcr2FieldClr_crcend2Mask  = 0x8000000
)

// GetClr_crcend2 Bank 2 CRCEND1 flag clear bit
func (r *registerCcr2Type) GetClr_crcend2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcr2FieldClr_crcend2Mask) != 0
}

// SetClr_crcend2 Bank 2 CRCEND1 flag clear bit
func (r *registerCcr2Type) SetClr_crcend2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcr2FieldClr_crcend2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcr2FieldClr_crcend2Mask)
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
	RegisterOptcrFieldSwap_bankShift = 31
	RegisterOptcrFieldSwap_bankMask  = 0x80000000
)

// GetSwap_bank Bank swapping configuration bit
func (r *registerOptcrType) GetSwap_bank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptcrFieldSwap_bankMask) != 0
}

// SetSwap_bank Bank swapping configuration bit
func (r *registerOptcrType) SetSwap_bank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptcrFieldSwap_bankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptcrFieldSwap_bankMask)
	}
}

// registerOptsr_curType FLASH option status register
type registerOptsr_curType uint32

const (
	RegisterOptsr_curFieldOpt_busyShift = 0
	RegisterOptsr_curFieldOpt_busyMask  = 0x1
)

// GetOpt_busy Option byte change ongoing flag
func (r *registerOptsr_curType) GetOpt_busy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldOpt_busyMask) != 0
}

// SetOpt_busy Option byte change ongoing flag
func (r *registerOptsr_curType) SetOpt_busy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldOpt_busyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldOpt_busyMask)
	}
}

const (
	RegisterOptsr_curFieldBor_levShift = 2
	RegisterOptsr_curFieldBor_levMask  = 0xc
)

// GetBor_lev Brownout level option status bit
func (r *registerOptsr_curType) GetBor_lev() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldBor_levMask) >> RegisterOptsr_curFieldBor_levShift)
}

// SetBor_lev Brownout level option status bit
func (r *registerOptsr_curType) SetBor_lev(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldBor_levMask)|(uint32(value)<<RegisterOptsr_curFieldBor_levShift))
}

const (
	RegisterOptsr_curFieldIwdg1_hwShift = 4
	RegisterOptsr_curFieldIwdg1_hwMask  = 0x10
)

// GetIwdg1_hw IWDG1 control option status bit
func (r *registerOptsr_curType) GetIwdg1_hw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldIwdg1_hwMask) != 0
}

// SetIwdg1_hw IWDG1 control option status bit
func (r *registerOptsr_curType) SetIwdg1_hw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldIwdg1_hwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldIwdg1_hwMask)
	}
}

const (
	RegisterOptsr_curFieldNrst_stop_d1Shift = 6
	RegisterOptsr_curFieldNrst_stop_d1Mask  = 0x40
)

// GetNrst_stop_d1 D1 DStop entry reset option status bit
func (r *registerOptsr_curType) GetNrst_stop_d1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldNrst_stop_d1Mask) != 0
}

// SetNrst_stop_d1 D1 DStop entry reset option status bit
func (r *registerOptsr_curType) SetNrst_stop_d1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldNrst_stop_d1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldNrst_stop_d1Mask)
	}
}

const (
	RegisterOptsr_curFieldNrst_stby_d1Shift = 7
	RegisterOptsr_curFieldNrst_stby_d1Mask  = 0x80
)

// GetNrst_stby_d1 D1 DStandby entry reset option status bit
func (r *registerOptsr_curType) GetNrst_stby_d1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldNrst_stby_d1Mask) != 0
}

// SetNrst_stby_d1 D1 DStandby entry reset option status bit
func (r *registerOptsr_curType) SetNrst_stby_d1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldNrst_stby_d1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldNrst_stby_d1Mask)
	}
}

const (
	RegisterOptsr_curFieldRdpShift = 8
	RegisterOptsr_curFieldRdpMask  = 0xff00
)

// GetRdp Readout protection level option status byte
func (r *registerOptsr_curType) GetRdp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldRdpMask) >> RegisterOptsr_curFieldRdpShift)
}

// SetRdp Readout protection level option status byte
func (r *registerOptsr_curType) SetRdp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldRdpMask)|(uint32(value)<<RegisterOptsr_curFieldRdpShift))
}

const (
	RegisterOptsr_curFieldFz_iwdg_stopShift = 17
	RegisterOptsr_curFieldFz_iwdg_stopMask  = 0x20000
)

// GetFz_iwdg_stop IWDG Stop mode freeze option status bit
func (r *registerOptsr_curType) GetFz_iwdg_stop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldFz_iwdg_stopMask) != 0
}

// SetFz_iwdg_stop IWDG Stop mode freeze option status bit
func (r *registerOptsr_curType) SetFz_iwdg_stop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldFz_iwdg_stopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldFz_iwdg_stopMask)
	}
}

const (
	RegisterOptsr_curFieldFz_iwdg_sdbyShift = 18
	RegisterOptsr_curFieldFz_iwdg_sdbyMask  = 0x40000
)

// GetFz_iwdg_sdby IWDG Standby mode freeze option status bit
func (r *registerOptsr_curType) GetFz_iwdg_sdby() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldFz_iwdg_sdbyMask) != 0
}

// SetFz_iwdg_sdby IWDG Standby mode freeze option status bit
func (r *registerOptsr_curType) SetFz_iwdg_sdby(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldFz_iwdg_sdbyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldFz_iwdg_sdbyMask)
	}
}

const (
	RegisterOptsr_curFieldSt_ram_sizeShift = 19
	RegisterOptsr_curFieldSt_ram_sizeMask  = 0x180000
)

// GetSt_ram_size DTCM RAM size option status
func (r *registerOptsr_curType) GetSt_ram_size() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldSt_ram_sizeMask) >> RegisterOptsr_curFieldSt_ram_sizeShift)
}

// SetSt_ram_size DTCM RAM size option status
func (r *registerOptsr_curType) SetSt_ram_size(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldSt_ram_sizeMask)|(uint32(value)<<RegisterOptsr_curFieldSt_ram_sizeShift))
}

const (
	RegisterOptsr_curFieldSecurityShift = 21
	RegisterOptsr_curFieldSecurityMask  = 0x200000
)

// GetSecurity Security enable option status bit
func (r *registerOptsr_curType) GetSecurity() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldSecurityMask) != 0
}

// SetSecurity Security enable option status bit
func (r *registerOptsr_curType) SetSecurity(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldSecurityMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldSecurityMask)
	}
}

const (
	RegisterOptsr_curFieldRss1Shift = 26
	RegisterOptsr_curFieldRss1Mask  = 0x4000000
)

// GetRss1 User option bit 1
func (r *registerOptsr_curType) GetRss1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldRss1Mask) != 0
}

// SetRss1 User option bit 1
func (r *registerOptsr_curType) SetRss1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldRss1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldRss1Mask)
	}
}

const (
	RegisterOptsr_curFieldPerso_okShift = 28
	RegisterOptsr_curFieldPerso_okMask  = 0x10000000
)

// GetPerso_ok Device personalization status bit
func (r *registerOptsr_curType) GetPerso_ok() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldPerso_okMask) != 0
}

// SetPerso_ok Device personalization status bit
func (r *registerOptsr_curType) SetPerso_ok(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldPerso_okMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldPerso_okMask)
	}
}

const (
	RegisterOptsr_curFieldIo_hslvShift = 29
	RegisterOptsr_curFieldIo_hslvMask  = 0x20000000
)

// GetIo_hslv I/O high-speed at low-voltage status bit (PRODUCT_BELOW_25V)
func (r *registerOptsr_curType) GetIo_hslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldIo_hslvMask) != 0
}

// SetIo_hslv I/O high-speed at low-voltage status bit (PRODUCT_BELOW_25V)
func (r *registerOptsr_curType) SetIo_hslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldIo_hslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldIo_hslvMask)
	}
}

const (
	RegisterOptsr_curFieldOptchangeerrShift = 30
	RegisterOptsr_curFieldOptchangeerrMask  = 0x40000000
)

// GetOptchangeerr Option byte change error flag
func (r *registerOptsr_curType) GetOptchangeerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldOptchangeerrMask) != 0
}

// SetOptchangeerr Option byte change error flag
func (r *registerOptsr_curType) SetOptchangeerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldOptchangeerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldOptchangeerrMask)
	}
}

const (
	RegisterOptsr_curFieldSwap_bank_optShift = 31
	RegisterOptsr_curFieldSwap_bank_optMask  = 0x80000000
)

// GetSwap_bank_opt Bank swapping option status bit
func (r *registerOptsr_curType) GetSwap_bank_opt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_curFieldSwap_bank_optMask) != 0
}

// SetSwap_bank_opt Bank swapping option status bit
func (r *registerOptsr_curType) SetSwap_bank_opt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_curFieldSwap_bank_optMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_curFieldSwap_bank_optMask)
	}
}

// registerOptsr_prgType FLASH option status register
type registerOptsr_prgType uint32

const (
	RegisterOptsr_prgFieldBor_levShift = 2
	RegisterOptsr_prgFieldBor_levMask  = 0xc
)

// GetBor_lev BOR reset level option configuration bits
func (r *registerOptsr_prgType) GetBor_lev() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldBor_levMask) >> RegisterOptsr_prgFieldBor_levShift)
}

// SetBor_lev BOR reset level option configuration bits
func (r *registerOptsr_prgType) SetBor_lev(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldBor_levMask)|(uint32(value)<<RegisterOptsr_prgFieldBor_levShift))
}

const (
	RegisterOptsr_prgFieldIwdg1_hwShift = 4
	RegisterOptsr_prgFieldIwdg1_hwMask  = 0x10
)

// GetIwdg1_hw IWDG1 option configuration bit
func (r *registerOptsr_prgType) GetIwdg1_hw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldIwdg1_hwMask) != 0
}

// SetIwdg1_hw IWDG1 option configuration bit
func (r *registerOptsr_prgType) SetIwdg1_hw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldIwdg1_hwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldIwdg1_hwMask)
	}
}

const (
	RegisterOptsr_prgFieldNrst_stop_d1Shift = 6
	RegisterOptsr_prgFieldNrst_stop_d1Mask  = 0x40
)

// GetNrst_stop_d1 Option byte erase after D1 DStop option configuration bit
func (r *registerOptsr_prgType) GetNrst_stop_d1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldNrst_stop_d1Mask) != 0
}

// SetNrst_stop_d1 Option byte erase after D1 DStop option configuration bit
func (r *registerOptsr_prgType) SetNrst_stop_d1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldNrst_stop_d1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldNrst_stop_d1Mask)
	}
}

const (
	RegisterOptsr_prgFieldNrst_stby_d1Shift = 7
	RegisterOptsr_prgFieldNrst_stby_d1Mask  = 0x80
)

// GetNrst_stby_d1 Option byte erase after D1 DStandby option configuration bit
func (r *registerOptsr_prgType) GetNrst_stby_d1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldNrst_stby_d1Mask) != 0
}

// SetNrst_stby_d1 Option byte erase after D1 DStandby option configuration bit
func (r *registerOptsr_prgType) SetNrst_stby_d1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldNrst_stby_d1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldNrst_stby_d1Mask)
	}
}

const (
	RegisterOptsr_prgFieldRdpShift = 8
	RegisterOptsr_prgFieldRdpMask  = 0xff00
)

// GetRdp Readout protection level option configuration byte
func (r *registerOptsr_prgType) GetRdp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldRdpMask) >> RegisterOptsr_prgFieldRdpShift)
}

// SetRdp Readout protection level option configuration byte
func (r *registerOptsr_prgType) SetRdp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldRdpMask)|(uint32(value)<<RegisterOptsr_prgFieldRdpShift))
}

const (
	RegisterOptsr_prgFieldFz_iwdg_stopShift = 17
	RegisterOptsr_prgFieldFz_iwdg_stopMask  = 0x20000
)

// GetFz_iwdg_stop IWDG Stop mode freeze option configuration bit
func (r *registerOptsr_prgType) GetFz_iwdg_stop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldFz_iwdg_stopMask) != 0
}

// SetFz_iwdg_stop IWDG Stop mode freeze option configuration bit
func (r *registerOptsr_prgType) SetFz_iwdg_stop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldFz_iwdg_stopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldFz_iwdg_stopMask)
	}
}

const (
	RegisterOptsr_prgFieldFz_iwdg_sdbyShift = 18
	RegisterOptsr_prgFieldFz_iwdg_sdbyMask  = 0x40000
)

// GetFz_iwdg_sdby IWDG Standby mode freeze option configuration bit
func (r *registerOptsr_prgType) GetFz_iwdg_sdby() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldFz_iwdg_sdbyMask) != 0
}

// SetFz_iwdg_sdby IWDG Standby mode freeze option configuration bit
func (r *registerOptsr_prgType) SetFz_iwdg_sdby(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldFz_iwdg_sdbyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldFz_iwdg_sdbyMask)
	}
}

const (
	RegisterOptsr_prgFieldSt_ram_sizeShift = 19
	RegisterOptsr_prgFieldSt_ram_sizeMask  = 0x180000
)

// GetSt_ram_size DTCM size select option configuration bits
func (r *registerOptsr_prgType) GetSt_ram_size() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldSt_ram_sizeMask) >> RegisterOptsr_prgFieldSt_ram_sizeShift)
}

// SetSt_ram_size DTCM size select option configuration bits
func (r *registerOptsr_prgType) SetSt_ram_size(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldSt_ram_sizeMask)|(uint32(value)<<RegisterOptsr_prgFieldSt_ram_sizeShift))
}

const (
	RegisterOptsr_prgFieldSecurityShift = 21
	RegisterOptsr_prgFieldSecurityMask  = 0x200000
)

// GetSecurity Security option configuration bit
func (r *registerOptsr_prgType) GetSecurity() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldSecurityMask) != 0
}

// SetSecurity Security option configuration bit
func (r *registerOptsr_prgType) SetSecurity(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldSecurityMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldSecurityMask)
	}
}

const (
	RegisterOptsr_prgFieldRss1Shift = 26
	RegisterOptsr_prgFieldRss1Mask  = 0x4000000
)

// GetRss1 User option configuration bit 1
func (r *registerOptsr_prgType) GetRss1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldRss1Mask) != 0
}

// SetRss1 User option configuration bit 1
func (r *registerOptsr_prgType) SetRss1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldRss1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldRss1Mask)
	}
}

const (
	RegisterOptsr_prgFieldRss2Shift = 27
	RegisterOptsr_prgFieldRss2Mask  = 0x8000000
)

// GetRss2 User option configuration bit 2
func (r *registerOptsr_prgType) GetRss2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldRss2Mask) != 0
}

// SetRss2 User option configuration bit 2
func (r *registerOptsr_prgType) SetRss2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldRss2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldRss2Mask)
	}
}

const (
	RegisterOptsr_prgFieldIo_hslvShift = 29
	RegisterOptsr_prgFieldIo_hslvMask  = 0x20000000
)

// GetIo_hslv I/O high-speed at low-voltage (PRODUCT_BELOW_25V)
func (r *registerOptsr_prgType) GetIo_hslv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldIo_hslvMask) != 0
}

// SetIo_hslv I/O high-speed at low-voltage (PRODUCT_BELOW_25V)
func (r *registerOptsr_prgType) SetIo_hslv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldIo_hslvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldIo_hslvMask)
	}
}

const (
	RegisterOptsr_prgFieldSwap_bank_optShift = 31
	RegisterOptsr_prgFieldSwap_bank_optMask  = 0x80000000
)

// GetSwap_bank_opt Bank swapping option configuration bit
func (r *registerOptsr_prgType) GetSwap_bank_opt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptsr_prgFieldSwap_bank_optMask) != 0
}

// SetSwap_bank_opt Bank swapping option configuration bit
func (r *registerOptsr_prgType) SetSwap_bank_opt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptsr_prgFieldSwap_bank_optMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptsr_prgFieldSwap_bank_optMask)
	}
}

// registerOptccrType FLASH option clear control register
type registerOptccrType uint32

const (
	RegisterOptccrFieldClr_optchangeerrShift = 30
	RegisterOptccrFieldClr_optchangeerrMask  = 0x40000000
)

// GetClr_optchangeerr OPTCHANGEERR reset bit
func (r *registerOptccrType) GetClr_optchangeerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOptccrFieldClr_optchangeerrMask) != 0
}

// SetClr_optchangeerr OPTCHANGEERR reset bit
func (r *registerOptccrType) SetClr_optchangeerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOptccrFieldClr_optchangeerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOptccrFieldClr_optchangeerrMask)
	}
}

// registerPrar_cur2Type FLASH protection address for bank 1
type registerPrar_cur2Type uint32

const (
	RegisterPrar_cur2FieldProt_area_start2Shift = 0
	RegisterPrar_cur2FieldProt_area_start2Mask  = 0xfff
)

// GetProt_area_start2 Bank 2 lowest PCROP protected address
func (r *registerPrar_cur2Type) GetProt_area_start2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrar_cur2FieldProt_area_start2Mask) >> RegisterPrar_cur2FieldProt_area_start2Shift)
}

// SetProt_area_start2 Bank 2 lowest PCROP protected address
func (r *registerPrar_cur2Type) SetProt_area_start2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrar_cur2FieldProt_area_start2Mask)|(uint32(value)<<RegisterPrar_cur2FieldProt_area_start2Shift))
}

const (
	RegisterPrar_cur2FieldProt_area_end2Shift = 16
	RegisterPrar_cur2FieldProt_area_end2Mask  = 0xfff0000
)

// GetProt_area_end2 Bank 2 highest PCROP protected address
func (r *registerPrar_cur2Type) GetProt_area_end2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPrar_cur2FieldProt_area_end2Mask) >> RegisterPrar_cur2FieldProt_area_end2Shift)
}

// SetProt_area_end2 Bank 2 highest PCROP protected address
func (r *registerPrar_cur2Type) SetProt_area_end2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrar_cur2FieldProt_area_end2Mask)|(uint32(value)<<RegisterPrar_cur2FieldProt_area_end2Shift))
}

const (
	RegisterPrar_cur2FieldDmep2Shift = 31
	RegisterPrar_cur2FieldDmep2Mask  = 0x80000000
)

// GetDmep2 Bank 2 PCROP protected erase enable option status bit
func (r *registerPrar_cur2Type) GetDmep2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPrar_cur2FieldDmep2Mask) != 0
}

// SetDmep2 Bank 2 PCROP protected erase enable option status bit
func (r *registerPrar_cur2Type) SetDmep2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPrar_cur2FieldDmep2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPrar_cur2FieldDmep2Mask)
	}
}

// registerScar_cur2Type FLASH secure address for bank 2
type registerScar_cur2Type uint32

const (
	RegisterScar_cur2FieldSec_area_start2Shift = 0
	RegisterScar_cur2FieldSec_area_start2Mask  = 0xfff
)

// GetSec_area_start2 Bank 2 lowest secure protected address
func (r *registerScar_cur2Type) GetSec_area_start2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScar_cur2FieldSec_area_start2Mask) >> RegisterScar_cur2FieldSec_area_start2Shift)
}

// SetSec_area_start2 Bank 2 lowest secure protected address
func (r *registerScar_cur2Type) SetSec_area_start2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScar_cur2FieldSec_area_start2Mask)|(uint32(value)<<RegisterScar_cur2FieldSec_area_start2Shift))
}

const (
	RegisterScar_cur2FieldSec_area_end2Shift = 16
	RegisterScar_cur2FieldSec_area_end2Mask  = 0xfff0000
)

// GetSec_area_end2 Bank 2 highest secure protected address
func (r *registerScar_cur2Type) GetSec_area_end2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScar_cur2FieldSec_area_end2Mask) >> RegisterScar_cur2FieldSec_area_end2Shift)
}

// SetSec_area_end2 Bank 2 highest secure protected address
func (r *registerScar_cur2Type) SetSec_area_end2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScar_cur2FieldSec_area_end2Mask)|(uint32(value)<<RegisterScar_cur2FieldSec_area_end2Shift))
}

const (
	RegisterScar_cur2FieldDmes2Shift = 31
	RegisterScar_cur2FieldDmes2Mask  = 0x80000000
)

// GetDmes2 Bank 2 secure protected erase enable option status bit
func (r *registerScar_cur2Type) GetDmes2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScar_cur2FieldDmes2Mask) != 0
}

// SetDmes2 Bank 2 secure protected erase enable option status bit
func (r *registerScar_cur2Type) SetDmes2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScar_cur2FieldDmes2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScar_cur2FieldDmes2Mask)
	}
}

// registerScar_prg2Type FLASH secure address for bank 2
type registerScar_prg2Type uint32

const (
	RegisterScar_prg2FieldSec_area_start2Shift = 0
	RegisterScar_prg2FieldSec_area_start2Mask  = 0xfff
)

// GetSec_area_start2 Bank 2 lowest secure protected address configuration
func (r *registerScar_prg2Type) GetSec_area_start2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScar_prg2FieldSec_area_start2Mask) >> RegisterScar_prg2FieldSec_area_start2Shift)
}

// SetSec_area_start2 Bank 2 lowest secure protected address configuration
func (r *registerScar_prg2Type) SetSec_area_start2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScar_prg2FieldSec_area_start2Mask)|(uint32(value)<<RegisterScar_prg2FieldSec_area_start2Shift))
}

const (
	RegisterScar_prg2FieldSec_area_end2Shift = 16
	RegisterScar_prg2FieldSec_area_end2Mask  = 0xfff0000
)

// GetSec_area_end2 Bank 2 highest secure protected address configuration
func (r *registerScar_prg2Type) GetSec_area_end2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterScar_prg2FieldSec_area_end2Mask) >> RegisterScar_prg2FieldSec_area_end2Shift)
}

// SetSec_area_end2 Bank 2 highest secure protected address configuration
func (r *registerScar_prg2Type) SetSec_area_end2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterScar_prg2FieldSec_area_end2Mask)|(uint32(value)<<RegisterScar_prg2FieldSec_area_end2Shift))
}

const (
	RegisterScar_prg2FieldDmes2Shift = 31
	RegisterScar_prg2FieldDmes2Mask  = 0x80000000
)

// GetDmes2 Bank 2 secure protected erase enable option configuration bit
func (r *registerScar_prg2Type) GetDmes2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterScar_prg2FieldDmes2Mask) != 0
}

// SetDmes2 Bank 2 secure protected erase enable option configuration bit
func (r *registerScar_prg2Type) SetDmes2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterScar_prg2FieldDmes2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterScar_prg2FieldDmes2Mask)
	}
}

// registerWpsn_cur2rType FLASH write sector protection for bank 2
type registerWpsn_cur2rType uint32

const (
	RegisterWpsn_cur2rFieldWrpsn2Shift = 0
	RegisterWpsn_cur2rFieldWrpsn2Mask  = 0xff
)

// GetWrpsn2 Bank 2 sector write protection option status byte
func (r *registerWpsn_cur2rType) GetWrpsn2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWpsn_cur2rFieldWrpsn2Mask) >> RegisterWpsn_cur2rFieldWrpsn2Shift)
}

// SetWrpsn2 Bank 2 sector write protection option status byte
func (r *registerWpsn_cur2rType) SetWrpsn2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWpsn_cur2rFieldWrpsn2Mask)|(uint32(value)<<RegisterWpsn_cur2rFieldWrpsn2Shift))
}

// registerWpsn_prg2rType FLASH write sector protection for bank 2
type registerWpsn_prg2rType uint32

const (
	RegisterWpsn_prg2rFieldWrpsn2Shift = 0
	RegisterWpsn_prg2rFieldWrpsn2Mask  = 0xff
)

// GetWrpsn2 Bank 2 sector write protection configuration byte
func (r *registerWpsn_prg2rType) GetWrpsn2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWpsn_prg2rFieldWrpsn2Mask) >> RegisterWpsn_prg2rFieldWrpsn2Shift)
}

// SetWrpsn2 Bank 2 sector write protection configuration byte
func (r *registerWpsn_prg2rType) SetWrpsn2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWpsn_prg2rFieldWrpsn2Mask)|(uint32(value)<<RegisterWpsn_prg2rFieldWrpsn2Shift))
}

// registerCrccr2Type FLASH CRC control register for bank 1
type registerCrccr2Type uint32

const (
	RegisterCrccr2FieldCrc_sectShift = 0
	RegisterCrccr2FieldCrc_sectMask  = 0x7
)

// GetCrc_sect Bank 2 CRC sector number
func (r *registerCrccr2Type) GetCrc_sect() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldCrc_sectMask) >> RegisterCrccr2FieldCrc_sectShift)
}

// SetCrc_sect Bank 2 CRC sector number
func (r *registerCrccr2Type) SetCrc_sect(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldCrc_sectMask)|(uint32(value)<<RegisterCrccr2FieldCrc_sectShift))
}

const (
	RegisterCrccr2FieldAll_bankShift = 7
	RegisterCrccr2FieldAll_bankMask  = 0x80
)

// GetAll_bank Bank 2 CRC select bit
func (r *registerCrccr2Type) GetAll_bank() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldAll_bankMask) != 0
}

// SetAll_bank Bank 2 CRC select bit
func (r *registerCrccr2Type) SetAll_bank(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr2FieldAll_bankMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldAll_bankMask)
	}
}

const (
	RegisterCrccr2FieldCrc_by_sectShift = 8
	RegisterCrccr2FieldCrc_by_sectMask  = 0x100
)

// GetCrc_by_sect Bank 2 CRC sector mode select bit
func (r *registerCrccr2Type) GetCrc_by_sect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldCrc_by_sectMask) != 0
}

// SetCrc_by_sect Bank 2 CRC sector mode select bit
func (r *registerCrccr2Type) SetCrc_by_sect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr2FieldCrc_by_sectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldCrc_by_sectMask)
	}
}

const (
	RegisterCrccr2FieldAdd_sectShift = 9
	RegisterCrccr2FieldAdd_sectMask  = 0x200
)

// GetAdd_sect Bank 2 CRC sector select bit
func (r *registerCrccr2Type) GetAdd_sect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldAdd_sectMask) != 0
}

// SetAdd_sect Bank 2 CRC sector select bit
func (r *registerCrccr2Type) SetAdd_sect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr2FieldAdd_sectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldAdd_sectMask)
	}
}

const (
	RegisterCrccr2FieldClean_sectShift = 10
	RegisterCrccr2FieldClean_sectMask  = 0x400
)

// GetClean_sect Bank 2 CRC sector list clear bit
func (r *registerCrccr2Type) GetClean_sect() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldClean_sectMask) != 0
}

// SetClean_sect Bank 2 CRC sector list clear bit
func (r *registerCrccr2Type) SetClean_sect(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr2FieldClean_sectMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldClean_sectMask)
	}
}

const (
	RegisterCrccr2FieldStart_crcShift = 16
	RegisterCrccr2FieldStart_crcMask  = 0x10000
)

// GetStart_crc Bank 2 CRC start bit
func (r *registerCrccr2Type) GetStart_crc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldStart_crcMask) != 0
}

// SetStart_crc Bank 2 CRC start bit
func (r *registerCrccr2Type) SetStart_crc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr2FieldStart_crcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldStart_crcMask)
	}
}

const (
	RegisterCrccr2FieldClean_crcShift = 17
	RegisterCrccr2FieldClean_crcMask  = 0x20000
)

// GetClean_crc Bank 2 CRC clear bit
func (r *registerCrccr2Type) GetClean_crc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldClean_crcMask) != 0
}

// SetClean_crc Bank 2 CRC clear bit
func (r *registerCrccr2Type) SetClean_crc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrccr2FieldClean_crcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldClean_crcMask)
	}
}

const (
	RegisterCrccr2FieldCrc_burstShift = 20
	RegisterCrccr2FieldCrc_burstMask  = 0x300000
)

// GetCrc_burst Bank 2 CRC burst size
func (r *registerCrccr2Type) GetCrc_burst() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrccr2FieldCrc_burstMask) >> RegisterCrccr2FieldCrc_burstShift)
}

// SetCrc_burst Bank 2 CRC burst size
func (r *registerCrccr2Type) SetCrc_burst(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrccr2FieldCrc_burstMask)|(uint32(value)<<RegisterCrccr2FieldCrc_burstShift))
}

// registerCrcsadd2rType FLASH CRC start address register for bank 2
type registerCrcsadd2rType uint32

const (
	RegisterCrcsadd2rFieldCrc_start_addrShift = 0
	RegisterCrcsadd2rFieldCrc_start_addrMask  = 0xffffffff
)

// GetCrc_start_addr CRC start address on bank 2
func (r *registerCrcsadd2rType) GetCrc_start_addr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrcsadd2rFieldCrc_start_addrMask) >> RegisterCrcsadd2rFieldCrc_start_addrShift)
}

// SetCrc_start_addr CRC start address on bank 2
func (r *registerCrcsadd2rType) SetCrc_start_addr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrcsadd2rFieldCrc_start_addrMask)|(uint32(value)<<RegisterCrcsadd2rFieldCrc_start_addrShift))
}

// registerCrceadd2rType FLASH CRC end address register for bank 2
type registerCrceadd2rType uint32

const (
	RegisterCrceadd2rFieldCrc_end_addrShift = 0
	RegisterCrceadd2rFieldCrc_end_addrMask  = 0xffffffff
)

// GetCrc_end_addr CRC end address on bank 2
func (r *registerCrceadd2rType) GetCrc_end_addr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrceadd2rFieldCrc_end_addrMask) >> RegisterCrceadd2rFieldCrc_end_addrShift)
}

// SetCrc_end_addr CRC end address on bank 2
func (r *registerCrceadd2rType) SetCrc_end_addr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrceadd2rFieldCrc_end_addrMask)|(uint32(value)<<RegisterCrceadd2rFieldCrc_end_addrShift))
}

// registerEcc_fa2rType FLASH ECC fail address for bank 2
type registerEcc_fa2rType uint32

const (
	RegisterEcc_fa2rFieldFail_ecc_addr2Shift = 0
	RegisterEcc_fa2rFieldFail_ecc_addr2Mask  = 0x7fff
)

// GetFail_ecc_addr2 Bank 2 ECC error address
func (r *registerEcc_fa2rType) GetFail_ecc_addr2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterEcc_fa2rFieldFail_ecc_addr2Mask) >> RegisterEcc_fa2rFieldFail_ecc_addr2Shift)
}

// SetFail_ecc_addr2 Bank 2 ECC error address
func (r *registerEcc_fa2rType) SetFail_ecc_addr2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEcc_fa2rFieldFail_ecc_addr2Mask)|(uint32(value)<<RegisterEcc_fa2rFieldFail_ecc_addr2Shift))
}
