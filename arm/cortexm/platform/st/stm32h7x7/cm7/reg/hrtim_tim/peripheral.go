//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package hrtim_tim

import (
	"unsafe"
	"volatile"
)

var (
	Hrtim_tima = (*_hrtim_tim)(unsafe.Pointer(uintptr(0x40017480)))
	Hrtim_timb = (*_hrtim_tim)(unsafe.Pointer(uintptr(0x40017500)))
	Hrtim_timc = (*_hrtim_tim)(unsafe.Pointer(uintptr(0x40017580)))
	Hrtim_timd = (*_hrtim_tim)(unsafe.Pointer(uintptr(0x40017600)))
	Hrtim_time = (*_hrtim_tim)(unsafe.Pointer(uintptr(0x40017680)))

	Instances = [5]*_hrtim_tim{
		Hrtim_tima,
		Hrtim_timb,
		Hrtim_timc,
		Hrtim_timd,
		Hrtim_time,
	}
)

type _hrtim_tim struct {
	Timcr    RegisterTimcrType
	Timisr   RegisterTimisrType
	Timicr   RegisterTimicrType
	Timdier5 RegisterTimdier5Type
	Cntr     RegisterCntrType
	Perr     RegisterPerrType
	Repr     RegisterReprType
	Cmp1r    RegisterCmp1rType
	Cmp1cr   RegisterCmp1crType
	Cmp2r    RegisterCmp2rType
	Cmp3r    RegisterCmp3rType
	Cmp4r    RegisterCmp4rType
	Cpt1r    RegisterCpt1rType
	Cpt2r    RegisterCpt2rType
	Dtr      RegisterDtrType
	Set1r    RegisterSet1rType
	Rst1r    RegisterRst1rType
	Set2r    RegisterSet2rType
	Rst2r    RegisterRst2rType
	Eefr1    RegisterEefr1Type
	Eefr2    RegisterEefr2Type
	Rstr     RegisterRstrType
	Chpr     RegisterChprType
	Cpt1cr   RegisterCpt1crType
	Cpt2cr   RegisterCpt2crType
	Outr     RegisterOutrType
	Fltr     RegisterFltrType
}

// RegisterTimcrType Timerx Control Register
type RegisterTimcrType uint32

func (r *RegisterTimcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTimcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTimcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTimcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTimcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTimcrFieldCkpscxShift = 0
	RegisterTimcrFieldCkpscxMask  = 0x7
)

// GetCkpscx HRTIM Timer x Clock prescaler
func (r *RegisterTimcrType) GetCkpscx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldCkpscxMask) >> RegisterTimcrFieldCkpscxShift)
}

// SetCkpscx HRTIM Timer x Clock prescaler
func (r *RegisterTimcrType) SetCkpscx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldCkpscxMask)|(uint32(value)<<RegisterTimcrFieldCkpscxShift))
}

const (
	RegisterTimcrFieldContShift = 3
	RegisterTimcrFieldContMask  = 0x8
)

// GetCont Continuous mode
func (r *RegisterTimcrType) GetCont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldContMask) != 0
}

// SetCont Continuous mode
func (r *RegisterTimcrType) SetCont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldContMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldContMask)
	}
}

const (
	RegisterTimcrFieldRetrigShift = 4
	RegisterTimcrFieldRetrigMask  = 0x10
)

// GetRetrig Re-triggerable mode
func (r *RegisterTimcrType) GetRetrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldRetrigMask) != 0
}

// SetRetrig Re-triggerable mode
func (r *RegisterTimcrType) SetRetrig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldRetrigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldRetrigMask)
	}
}

const (
	RegisterTimcrFieldHalfShift = 5
	RegisterTimcrFieldHalfMask  = 0x20
)

// GetHalf Half mode enable
func (r *RegisterTimcrType) GetHalf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldHalfMask) != 0
}

// SetHalf Half mode enable
func (r *RegisterTimcrType) SetHalf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldHalfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldHalfMask)
	}
}

const (
	RegisterTimcrFieldPshpllShift = 6
	RegisterTimcrFieldPshpllMask  = 0x40
)

// GetPshpll Push-Pull mode enable
func (r *RegisterTimcrType) GetPshpll() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldPshpllMask) != 0
}

// SetPshpll Push-Pull mode enable
func (r *RegisterTimcrType) SetPshpll(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldPshpllMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldPshpllMask)
	}
}

const (
	RegisterTimcrFieldSyncrstxShift = 10
	RegisterTimcrFieldSyncrstxMask  = 0x400
)

// GetSyncrstx Synchronization Resets Timer x
func (r *RegisterTimcrType) GetSyncrstx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldSyncrstxMask) != 0
}

// SetSyncrstx Synchronization Resets Timer x
func (r *RegisterTimcrType) SetSyncrstx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldSyncrstxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldSyncrstxMask)
	}
}

const (
	RegisterTimcrFieldSyncstrtxShift = 11
	RegisterTimcrFieldSyncstrtxMask  = 0x800
)

// GetSyncstrtx Synchronization Starts Timer x
func (r *RegisterTimcrType) GetSyncstrtx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldSyncstrtxMask) != 0
}

// SetSyncstrtx Synchronization Starts Timer x
func (r *RegisterTimcrType) SetSyncstrtx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldSyncstrtxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldSyncstrtxMask)
	}
}

const (
	RegisterTimcrFieldDelcmp2Shift = 12
	RegisterTimcrFieldDelcmp2Mask  = 0x3000
)

// GetDelcmp2 Delayed CMP2 mode
func (r *RegisterTimcrType) GetDelcmp2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldDelcmp2Mask) >> RegisterTimcrFieldDelcmp2Shift)
}

// SetDelcmp2 Delayed CMP2 mode
func (r *RegisterTimcrType) SetDelcmp2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldDelcmp2Mask)|(uint32(value)<<RegisterTimcrFieldDelcmp2Shift))
}

const (
	RegisterTimcrFieldDelcmp4Shift = 14
	RegisterTimcrFieldDelcmp4Mask  = 0xc000
)

// GetDelcmp4 Delayed CMP4 mode
func (r *RegisterTimcrType) GetDelcmp4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldDelcmp4Mask) >> RegisterTimcrFieldDelcmp4Shift)
}

// SetDelcmp4 Delayed CMP4 mode
func (r *RegisterTimcrType) SetDelcmp4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldDelcmp4Mask)|(uint32(value)<<RegisterTimcrFieldDelcmp4Shift))
}

const (
	RegisterTimcrFieldTxrepuShift = 17
	RegisterTimcrFieldTxrepuMask  = 0x20000
)

// GetTxrepu Timer x Repetition update
func (r *RegisterTimcrType) GetTxrepu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldTxrepuMask) != 0
}

// SetTxrepu Timer x Repetition update
func (r *RegisterTimcrType) SetTxrepu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldTxrepuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldTxrepuMask)
	}
}

const (
	RegisterTimcrFieldTxrstuShift = 18
	RegisterTimcrFieldTxrstuMask  = 0x40000
)

// GetTxrstu Timerx reset update
func (r *RegisterTimcrType) GetTxrstu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldTxrstuMask) != 0
}

// SetTxrstu Timerx reset update
func (r *RegisterTimcrType) SetTxrstu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldTxrstuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldTxrstuMask)
	}
}

const (
	RegisterTimcrFieldTbuShift = 20
	RegisterTimcrFieldTbuMask  = 0x100000
)

// GetTbu TBU
func (r *RegisterTimcrType) GetTbu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldTbuMask) != 0
}

// SetTbu TBU
func (r *RegisterTimcrType) SetTbu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldTbuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldTbuMask)
	}
}

const (
	RegisterTimcrFieldTcuShift = 21
	RegisterTimcrFieldTcuMask  = 0x200000
)

// GetTcu TCU
func (r *RegisterTimcrType) GetTcu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldTcuMask) != 0
}

// SetTcu TCU
func (r *RegisterTimcrType) SetTcu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldTcuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldTcuMask)
	}
}

const (
	RegisterTimcrFieldTduShift = 22
	RegisterTimcrFieldTduMask  = 0x400000
)

// GetTdu TDU
func (r *RegisterTimcrType) GetTdu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldTduMask) != 0
}

// SetTdu TDU
func (r *RegisterTimcrType) SetTdu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldTduMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldTduMask)
	}
}

const (
	RegisterTimcrFieldTeuShift = 23
	RegisterTimcrFieldTeuMask  = 0x800000
)

// GetTeu TEU
func (r *RegisterTimcrType) GetTeu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldTeuMask) != 0
}

// SetTeu TEU
func (r *RegisterTimcrType) SetTeu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldTeuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldTeuMask)
	}
}

const (
	RegisterTimcrFieldMstuShift = 24
	RegisterTimcrFieldMstuMask  = 0x1000000
)

// GetMstu Master Timer update
func (r *RegisterTimcrType) GetMstu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldMstuMask) != 0
}

// SetMstu Master Timer update
func (r *RegisterTimcrType) SetMstu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldMstuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldMstuMask)
	}
}

const (
	RegisterTimcrFieldDacsyncShift = 25
	RegisterTimcrFieldDacsyncMask  = 0x6000000
)

// GetDacsync AC Synchronization
func (r *RegisterTimcrType) GetDacsync() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldDacsyncMask) >> RegisterTimcrFieldDacsyncShift)
}

// SetDacsync AC Synchronization
func (r *RegisterTimcrType) SetDacsync(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldDacsyncMask)|(uint32(value)<<RegisterTimcrFieldDacsyncShift))
}

const (
	RegisterTimcrFieldPreenShift = 27
	RegisterTimcrFieldPreenMask  = 0x8000000
)

// GetPreen Preload enable
func (r *RegisterTimcrType) GetPreen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldPreenMask) != 0
}

// SetPreen Preload enable
func (r *RegisterTimcrType) SetPreen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcrFieldPreenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldPreenMask)
	}
}

const (
	RegisterTimcrFieldUpdgatShift = 28
	RegisterTimcrFieldUpdgatMask  = 0xf0000000
)

// GetUpdgat Update Gating
func (r *RegisterTimcrType) GetUpdgat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimcrFieldUpdgatMask) >> RegisterTimcrFieldUpdgatShift)
}

// SetUpdgat Update Gating
func (r *RegisterTimcrType) SetUpdgat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimcrFieldUpdgatMask)|(uint32(value)<<RegisterTimcrFieldUpdgatShift))
}

// RegisterTimisrType Timerx Interrupt Status Register
type RegisterTimisrType uint32

func (r *RegisterTimisrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTimisrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTimisrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTimisrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTimisrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTimisrFieldCmp1Shift = 0
	RegisterTimisrFieldCmp1Mask  = 0x1
)

// GetCmp1 Compare 1 Interrupt Flag
func (r *RegisterTimisrType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldCmp1Mask) != 0
}

// SetCmp1 Compare 1 Interrupt Flag
func (r *RegisterTimisrType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldCmp1Mask)
	}
}

const (
	RegisterTimisrFieldCmp2Shift = 1
	RegisterTimisrFieldCmp2Mask  = 0x2
)

// GetCmp2 Compare 2 Interrupt Flag
func (r *RegisterTimisrType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldCmp2Mask) != 0
}

// SetCmp2 Compare 2 Interrupt Flag
func (r *RegisterTimisrType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldCmp2Mask)
	}
}

const (
	RegisterTimisrFieldCmp3Shift = 2
	RegisterTimisrFieldCmp3Mask  = 0x4
)

// GetCmp3 Compare 3 Interrupt Flag
func (r *RegisterTimisrType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldCmp3Mask) != 0
}

// SetCmp3 Compare 3 Interrupt Flag
func (r *RegisterTimisrType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldCmp3Mask)
	}
}

const (
	RegisterTimisrFieldCmp4Shift = 3
	RegisterTimisrFieldCmp4Mask  = 0x8
)

// GetCmp4 Compare 4 Interrupt Flag
func (r *RegisterTimisrType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldCmp4Mask) != 0
}

// SetCmp4 Compare 4 Interrupt Flag
func (r *RegisterTimisrType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldCmp4Mask)
	}
}

const (
	RegisterTimisrFieldRepShift = 4
	RegisterTimisrFieldRepMask  = 0x10
)

// GetRep Repetition Interrupt Flag
func (r *RegisterTimisrType) GetRep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldRepMask) != 0
}

// SetRep Repetition Interrupt Flag
func (r *RegisterTimisrType) SetRep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldRepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldRepMask)
	}
}

const (
	RegisterTimisrFieldUpdShift = 6
	RegisterTimisrFieldUpdMask  = 0x40
)

// GetUpd Update Interrupt Flag
func (r *RegisterTimisrType) GetUpd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldUpdMask) != 0
}

// SetUpd Update Interrupt Flag
func (r *RegisterTimisrType) SetUpd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldUpdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldUpdMask)
	}
}

const (
	RegisterTimisrFieldCpt1Shift = 7
	RegisterTimisrFieldCpt1Mask  = 0x80
)

// GetCpt1 Capture1 Interrupt Flag
func (r *RegisterTimisrType) GetCpt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldCpt1Mask) != 0
}

// SetCpt1 Capture1 Interrupt Flag
func (r *RegisterTimisrType) SetCpt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldCpt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldCpt1Mask)
	}
}

const (
	RegisterTimisrFieldCpt2Shift = 8
	RegisterTimisrFieldCpt2Mask  = 0x100
)

// GetCpt2 Capture2 Interrupt Flag
func (r *RegisterTimisrType) GetCpt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldCpt2Mask) != 0
}

// SetCpt2 Capture2 Interrupt Flag
func (r *RegisterTimisrType) SetCpt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldCpt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldCpt2Mask)
	}
}

const (
	RegisterTimisrFieldSetx1Shift = 9
	RegisterTimisrFieldSetx1Mask  = 0x200
)

// GetSetx1 Output 1 Set Interrupt Flag
func (r *RegisterTimisrType) GetSetx1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldSetx1Mask) != 0
}

// SetSetx1 Output 1 Set Interrupt Flag
func (r *RegisterTimisrType) SetSetx1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldSetx1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldSetx1Mask)
	}
}

const (
	RegisterTimisrFieldRstx1Shift = 10
	RegisterTimisrFieldRstx1Mask  = 0x400
)

// GetRstx1 Output 1 Reset Interrupt Flag
func (r *RegisterTimisrType) GetRstx1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldRstx1Mask) != 0
}

// SetRstx1 Output 1 Reset Interrupt Flag
func (r *RegisterTimisrType) SetRstx1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldRstx1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldRstx1Mask)
	}
}

const (
	RegisterTimisrFieldSetx2Shift = 11
	RegisterTimisrFieldSetx2Mask  = 0x800
)

// GetSetx2 Output 2 Set Interrupt Flag
func (r *RegisterTimisrType) GetSetx2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldSetx2Mask) != 0
}

// SetSetx2 Output 2 Set Interrupt Flag
func (r *RegisterTimisrType) SetSetx2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldSetx2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldSetx2Mask)
	}
}

const (
	RegisterTimisrFieldRstx2Shift = 12
	RegisterTimisrFieldRstx2Mask  = 0x1000
)

// GetRstx2 Output 2 Reset Interrupt Flag
func (r *RegisterTimisrType) GetRstx2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldRstx2Mask) != 0
}

// SetRstx2 Output 2 Reset Interrupt Flag
func (r *RegisterTimisrType) SetRstx2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldRstx2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldRstx2Mask)
	}
}

const (
	RegisterTimisrFieldRstShift = 13
	RegisterTimisrFieldRstMask  = 0x2000
)

// GetRst Reset Interrupt Flag
func (r *RegisterTimisrType) GetRst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldRstMask) != 0
}

// SetRst Reset Interrupt Flag
func (r *RegisterTimisrType) SetRst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldRstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldRstMask)
	}
}

const (
	RegisterTimisrFieldDlyprtShift = 14
	RegisterTimisrFieldDlyprtMask  = 0x4000
)

// GetDlyprt Delayed Protection Flag
func (r *RegisterTimisrType) GetDlyprt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldDlyprtMask) != 0
}

// SetDlyprt Delayed Protection Flag
func (r *RegisterTimisrType) SetDlyprt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldDlyprtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldDlyprtMask)
	}
}

const (
	RegisterTimisrFieldCppstatShift = 16
	RegisterTimisrFieldCppstatMask  = 0x10000
)

// GetCppstat Current Push Pull Status
func (r *RegisterTimisrType) GetCppstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldCppstatMask) != 0
}

// SetCppstat Current Push Pull Status
func (r *RegisterTimisrType) SetCppstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldCppstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldCppstatMask)
	}
}

const (
	RegisterTimisrFieldIppstatShift = 17
	RegisterTimisrFieldIppstatMask  = 0x20000
)

// GetIppstat Idle Push Pull Status
func (r *RegisterTimisrType) GetIppstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldIppstatMask) != 0
}

// SetIppstat Idle Push Pull Status
func (r *RegisterTimisrType) SetIppstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldIppstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldIppstatMask)
	}
}

const (
	RegisterTimisrFieldO1statShift = 18
	RegisterTimisrFieldO1statMask  = 0x40000
)

// GetO1stat Output 1 State
func (r *RegisterTimisrType) GetO1stat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldO1statMask) != 0
}

// SetO1stat Output 1 State
func (r *RegisterTimisrType) SetO1stat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldO1statMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldO1statMask)
	}
}

const (
	RegisterTimisrFieldO2statShift = 19
	RegisterTimisrFieldO2statMask  = 0x80000
)

// GetO2stat Output 2 State
func (r *RegisterTimisrType) GetO2stat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimisrFieldO2statMask) != 0
}

// SetO2stat Output 2 State
func (r *RegisterTimisrType) SetO2stat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimisrFieldO2statMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimisrFieldO2statMask)
	}
}

// RegisterTimicrType Timerx Interrupt Clear Register
type RegisterTimicrType uint32

func (r *RegisterTimicrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTimicrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTimicrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTimicrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTimicrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTimicrFieldCmp1cShift = 0
	RegisterTimicrFieldCmp1cMask  = 0x1
)

// GetCmp1c Compare 1 Interrupt flag Clear
func (r *RegisterTimicrType) GetCmp1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldCmp1cMask) != 0
}

// SetCmp1c Compare 1 Interrupt flag Clear
func (r *RegisterTimicrType) SetCmp1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldCmp1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldCmp1cMask)
	}
}

const (
	RegisterTimicrFieldCmp2cShift = 1
	RegisterTimicrFieldCmp2cMask  = 0x2
)

// GetCmp2c Compare 2 Interrupt flag Clear
func (r *RegisterTimicrType) GetCmp2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldCmp2cMask) != 0
}

// SetCmp2c Compare 2 Interrupt flag Clear
func (r *RegisterTimicrType) SetCmp2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldCmp2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldCmp2cMask)
	}
}

const (
	RegisterTimicrFieldCmp3cShift = 2
	RegisterTimicrFieldCmp3cMask  = 0x4
)

// GetCmp3c Compare 3 Interrupt flag Clear
func (r *RegisterTimicrType) GetCmp3c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldCmp3cMask) != 0
}

// SetCmp3c Compare 3 Interrupt flag Clear
func (r *RegisterTimicrType) SetCmp3c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldCmp3cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldCmp3cMask)
	}
}

const (
	RegisterTimicrFieldCmp4cShift = 3
	RegisterTimicrFieldCmp4cMask  = 0x8
)

// GetCmp4c Compare 4 Interrupt flag Clear
func (r *RegisterTimicrType) GetCmp4c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldCmp4cMask) != 0
}

// SetCmp4c Compare 4 Interrupt flag Clear
func (r *RegisterTimicrType) SetCmp4c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldCmp4cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldCmp4cMask)
	}
}

const (
	RegisterTimicrFieldRepcShift = 4
	RegisterTimicrFieldRepcMask  = 0x10
)

// GetRepc Repetition Interrupt flag Clear
func (r *RegisterTimicrType) GetRepc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldRepcMask) != 0
}

// SetRepc Repetition Interrupt flag Clear
func (r *RegisterTimicrType) SetRepc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldRepcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldRepcMask)
	}
}

const (
	RegisterTimicrFieldUpdcShift = 6
	RegisterTimicrFieldUpdcMask  = 0x40
)

// GetUpdc Update Interrupt flag Clear
func (r *RegisterTimicrType) GetUpdc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldUpdcMask) != 0
}

// SetUpdc Update Interrupt flag Clear
func (r *RegisterTimicrType) SetUpdc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldUpdcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldUpdcMask)
	}
}

const (
	RegisterTimicrFieldCpt1cShift = 7
	RegisterTimicrFieldCpt1cMask  = 0x80
)

// GetCpt1c Capture1 Interrupt flag Clear
func (r *RegisterTimicrType) GetCpt1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldCpt1cMask) != 0
}

// SetCpt1c Capture1 Interrupt flag Clear
func (r *RegisterTimicrType) SetCpt1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldCpt1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldCpt1cMask)
	}
}

const (
	RegisterTimicrFieldCpt2cShift = 8
	RegisterTimicrFieldCpt2cMask  = 0x100
)

// GetCpt2c Capture2 Interrupt flag Clear
func (r *RegisterTimicrType) GetCpt2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldCpt2cMask) != 0
}

// SetCpt2c Capture2 Interrupt flag Clear
func (r *RegisterTimicrType) SetCpt2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldCpt2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldCpt2cMask)
	}
}

const (
	RegisterTimicrFieldSet1xcShift = 9
	RegisterTimicrFieldSet1xcMask  = 0x200
)

// GetSet1xc Output 1 Set flag Clear
func (r *RegisterTimicrType) GetSet1xc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldSet1xcMask) != 0
}

// SetSet1xc Output 1 Set flag Clear
func (r *RegisterTimicrType) SetSet1xc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldSet1xcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldSet1xcMask)
	}
}

const (
	RegisterTimicrFieldRstx1cShift = 10
	RegisterTimicrFieldRstx1cMask  = 0x400
)

// GetRstx1c Output 1 Reset flag Clear
func (r *RegisterTimicrType) GetRstx1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldRstx1cMask) != 0
}

// SetRstx1c Output 1 Reset flag Clear
func (r *RegisterTimicrType) SetRstx1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldRstx1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldRstx1cMask)
	}
}

const (
	RegisterTimicrFieldSet2xcShift = 11
	RegisterTimicrFieldSet2xcMask  = 0x800
)

// GetSet2xc Output 2 Set flag Clear
func (r *RegisterTimicrType) GetSet2xc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldSet2xcMask) != 0
}

// SetSet2xc Output 2 Set flag Clear
func (r *RegisterTimicrType) SetSet2xc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldSet2xcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldSet2xcMask)
	}
}

const (
	RegisterTimicrFieldRstx2cShift = 12
	RegisterTimicrFieldRstx2cMask  = 0x1000
)

// GetRstx2c Output 2 Reset flag Clear
func (r *RegisterTimicrType) GetRstx2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldRstx2cMask) != 0
}

// SetRstx2c Output 2 Reset flag Clear
func (r *RegisterTimicrType) SetRstx2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldRstx2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldRstx2cMask)
	}
}

const (
	RegisterTimicrFieldRstcShift = 13
	RegisterTimicrFieldRstcMask  = 0x2000
)

// GetRstc Reset Interrupt flag Clear
func (r *RegisterTimicrType) GetRstc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldRstcMask) != 0
}

// SetRstc Reset Interrupt flag Clear
func (r *RegisterTimicrType) SetRstc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldRstcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldRstcMask)
	}
}

const (
	RegisterTimicrFieldDlyprtcShift = 14
	RegisterTimicrFieldDlyprtcMask  = 0x4000
)

// GetDlyprtc Delayed Protection Flag Clear
func (r *RegisterTimicrType) GetDlyprtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimicrFieldDlyprtcMask) != 0
}

// SetDlyprtc Delayed Protection Flag Clear
func (r *RegisterTimicrType) SetDlyprtc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimicrFieldDlyprtcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimicrFieldDlyprtcMask)
	}
}

// RegisterTimdier5Type TIMxDIER5
type RegisterTimdier5Type uint32

func (r *RegisterTimdier5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTimdier5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTimdier5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTimdier5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTimdier5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTimdier5FieldCmp1ieShift = 0
	RegisterTimdier5FieldCmp1ieMask  = 0x1
)

// GetCmp1ie CMP1IE
func (r *RegisterTimdier5Type) GetCmp1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldCmp1ieMask) != 0
}

// SetCmp1ie CMP1IE
func (r *RegisterTimdier5Type) SetCmp1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldCmp1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldCmp1ieMask)
	}
}

const (
	RegisterTimdier5FieldCmp2ieShift = 1
	RegisterTimdier5FieldCmp2ieMask  = 0x2
)

// GetCmp2ie CMP2IE
func (r *RegisterTimdier5Type) GetCmp2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldCmp2ieMask) != 0
}

// SetCmp2ie CMP2IE
func (r *RegisterTimdier5Type) SetCmp2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldCmp2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldCmp2ieMask)
	}
}

const (
	RegisterTimdier5FieldCmp3ieShift = 2
	RegisterTimdier5FieldCmp3ieMask  = 0x4
)

// GetCmp3ie CMP3IE
func (r *RegisterTimdier5Type) GetCmp3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldCmp3ieMask) != 0
}

// SetCmp3ie CMP3IE
func (r *RegisterTimdier5Type) SetCmp3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldCmp3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldCmp3ieMask)
	}
}

const (
	RegisterTimdier5FieldCmp4ieShift = 3
	RegisterTimdier5FieldCmp4ieMask  = 0x8
)

// GetCmp4ie CMP4IE
func (r *RegisterTimdier5Type) GetCmp4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldCmp4ieMask) != 0
}

// SetCmp4ie CMP4IE
func (r *RegisterTimdier5Type) SetCmp4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldCmp4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldCmp4ieMask)
	}
}

const (
	RegisterTimdier5FieldRepieShift = 4
	RegisterTimdier5FieldRepieMask  = 0x10
)

// GetRepie REPIE
func (r *RegisterTimdier5Type) GetRepie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldRepieMask) != 0
}

// SetRepie REPIE
func (r *RegisterTimdier5Type) SetRepie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldRepieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldRepieMask)
	}
}

const (
	RegisterTimdier5FieldUpdieShift = 6
	RegisterTimdier5FieldUpdieMask  = 0x40
)

// GetUpdie UPDIE
func (r *RegisterTimdier5Type) GetUpdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldUpdieMask) != 0
}

// SetUpdie UPDIE
func (r *RegisterTimdier5Type) SetUpdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldUpdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldUpdieMask)
	}
}

const (
	RegisterTimdier5FieldCpt1ieShift = 7
	RegisterTimdier5FieldCpt1ieMask  = 0x80
)

// GetCpt1ie CPT1IE
func (r *RegisterTimdier5Type) GetCpt1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldCpt1ieMask) != 0
}

// SetCpt1ie CPT1IE
func (r *RegisterTimdier5Type) SetCpt1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldCpt1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldCpt1ieMask)
	}
}

const (
	RegisterTimdier5FieldCpt2ieShift = 8
	RegisterTimdier5FieldCpt2ieMask  = 0x100
)

// GetCpt2ie CPT2IE
func (r *RegisterTimdier5Type) GetCpt2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldCpt2ieMask) != 0
}

// SetCpt2ie CPT2IE
func (r *RegisterTimdier5Type) SetCpt2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldCpt2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldCpt2ieMask)
	}
}

const (
	RegisterTimdier5FieldSet1xieShift = 9
	RegisterTimdier5FieldSet1xieMask  = 0x200
)

// GetSet1xie SET1xIE
func (r *RegisterTimdier5Type) GetSet1xie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldSet1xieMask) != 0
}

// SetSet1xie SET1xIE
func (r *RegisterTimdier5Type) SetSet1xie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldSet1xieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldSet1xieMask)
	}
}

const (
	RegisterTimdier5FieldRstx1ieShift = 10
	RegisterTimdier5FieldRstx1ieMask  = 0x400
)

// GetRstx1ie RSTx1IE
func (r *RegisterTimdier5Type) GetRstx1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldRstx1ieMask) != 0
}

// SetRstx1ie RSTx1IE
func (r *RegisterTimdier5Type) SetRstx1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldRstx1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldRstx1ieMask)
	}
}

const (
	RegisterTimdier5FieldSetx2ieShift = 11
	RegisterTimdier5FieldSetx2ieMask  = 0x800
)

// GetSetx2ie SETx2IE
func (r *RegisterTimdier5Type) GetSetx2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldSetx2ieMask) != 0
}

// SetSetx2ie SETx2IE
func (r *RegisterTimdier5Type) SetSetx2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldSetx2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldSetx2ieMask)
	}
}

const (
	RegisterTimdier5FieldRstx2ieShift = 12
	RegisterTimdier5FieldRstx2ieMask  = 0x1000
)

// GetRstx2ie RSTx2IE
func (r *RegisterTimdier5Type) GetRstx2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldRstx2ieMask) != 0
}

// SetRstx2ie RSTx2IE
func (r *RegisterTimdier5Type) SetRstx2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldRstx2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldRstx2ieMask)
	}
}

const (
	RegisterTimdier5FieldRstieShift = 13
	RegisterTimdier5FieldRstieMask  = 0x2000
)

// GetRstie RSTIE
func (r *RegisterTimdier5Type) GetRstie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldRstieMask) != 0
}

// SetRstie RSTIE
func (r *RegisterTimdier5Type) SetRstie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldRstieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldRstieMask)
	}
}

const (
	RegisterTimdier5FieldDlyprtieShift = 14
	RegisterTimdier5FieldDlyprtieMask  = 0x4000
)

// GetDlyprtie DLYPRTIE
func (r *RegisterTimdier5Type) GetDlyprtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldDlyprtieMask) != 0
}

// SetDlyprtie DLYPRTIE
func (r *RegisterTimdier5Type) SetDlyprtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldDlyprtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldDlyprtieMask)
	}
}

const (
	RegisterTimdier5FieldCmp1deShift = 16
	RegisterTimdier5FieldCmp1deMask  = 0x10000
)

// GetCmp1de CMP1DE
func (r *RegisterTimdier5Type) GetCmp1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldCmp1deMask) != 0
}

// SetCmp1de CMP1DE
func (r *RegisterTimdier5Type) SetCmp1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldCmp1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldCmp1deMask)
	}
}

const (
	RegisterTimdier5FieldCmp2deShift = 17
	RegisterTimdier5FieldCmp2deMask  = 0x20000
)

// GetCmp2de CMP2DE
func (r *RegisterTimdier5Type) GetCmp2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldCmp2deMask) != 0
}

// SetCmp2de CMP2DE
func (r *RegisterTimdier5Type) SetCmp2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldCmp2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldCmp2deMask)
	}
}

const (
	RegisterTimdier5FieldCmp3deShift = 18
	RegisterTimdier5FieldCmp3deMask  = 0x40000
)

// GetCmp3de CMP3DE
func (r *RegisterTimdier5Type) GetCmp3de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldCmp3deMask) != 0
}

// SetCmp3de CMP3DE
func (r *RegisterTimdier5Type) SetCmp3de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldCmp3deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldCmp3deMask)
	}
}

const (
	RegisterTimdier5FieldCmp4deShift = 19
	RegisterTimdier5FieldCmp4deMask  = 0x80000
)

// GetCmp4de CMP4DE
func (r *RegisterTimdier5Type) GetCmp4de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldCmp4deMask) != 0
}

// SetCmp4de CMP4DE
func (r *RegisterTimdier5Type) SetCmp4de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldCmp4deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldCmp4deMask)
	}
}

const (
	RegisterTimdier5FieldRepdeShift = 20
	RegisterTimdier5FieldRepdeMask  = 0x100000
)

// GetRepde REPDE
func (r *RegisterTimdier5Type) GetRepde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldRepdeMask) != 0
}

// SetRepde REPDE
func (r *RegisterTimdier5Type) SetRepde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldRepdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldRepdeMask)
	}
}

const (
	RegisterTimdier5FieldUpddeShift = 22
	RegisterTimdier5FieldUpddeMask  = 0x400000
)

// GetUpdde UPDDE
func (r *RegisterTimdier5Type) GetUpdde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldUpddeMask) != 0
}

// SetUpdde UPDDE
func (r *RegisterTimdier5Type) SetUpdde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldUpddeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldUpddeMask)
	}
}

const (
	RegisterTimdier5FieldCpt1deShift = 23
	RegisterTimdier5FieldCpt1deMask  = 0x800000
)

// GetCpt1de CPT1DE
func (r *RegisterTimdier5Type) GetCpt1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldCpt1deMask) != 0
}

// SetCpt1de CPT1DE
func (r *RegisterTimdier5Type) SetCpt1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldCpt1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldCpt1deMask)
	}
}

const (
	RegisterTimdier5FieldCpt2deShift = 24
	RegisterTimdier5FieldCpt2deMask  = 0x1000000
)

// GetCpt2de CPT2DE
func (r *RegisterTimdier5Type) GetCpt2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldCpt2deMask) != 0
}

// SetCpt2de CPT2DE
func (r *RegisterTimdier5Type) SetCpt2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldCpt2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldCpt2deMask)
	}
}

const (
	RegisterTimdier5FieldSet1xdeShift = 25
	RegisterTimdier5FieldSet1xdeMask  = 0x2000000
)

// GetSet1xde SET1xDE
func (r *RegisterTimdier5Type) GetSet1xde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldSet1xdeMask) != 0
}

// SetSet1xde SET1xDE
func (r *RegisterTimdier5Type) SetSet1xde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldSet1xdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldSet1xdeMask)
	}
}

const (
	RegisterTimdier5FieldRstx1deShift = 26
	RegisterTimdier5FieldRstx1deMask  = 0x4000000
)

// GetRstx1de RSTx1DE
func (r *RegisterTimdier5Type) GetRstx1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldRstx1deMask) != 0
}

// SetRstx1de RSTx1DE
func (r *RegisterTimdier5Type) SetRstx1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldRstx1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldRstx1deMask)
	}
}

const (
	RegisterTimdier5FieldSetx2deShift = 27
	RegisterTimdier5FieldSetx2deMask  = 0x8000000
)

// GetSetx2de SETx2DE
func (r *RegisterTimdier5Type) GetSetx2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldSetx2deMask) != 0
}

// SetSetx2de SETx2DE
func (r *RegisterTimdier5Type) SetSetx2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldSetx2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldSetx2deMask)
	}
}

const (
	RegisterTimdier5FieldRstx2deShift = 28
	RegisterTimdier5FieldRstx2deMask  = 0x10000000
)

// GetRstx2de RSTx2DE
func (r *RegisterTimdier5Type) GetRstx2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldRstx2deMask) != 0
}

// SetRstx2de RSTx2DE
func (r *RegisterTimdier5Type) SetRstx2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldRstx2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldRstx2deMask)
	}
}

const (
	RegisterTimdier5FieldRstdeShift = 29
	RegisterTimdier5FieldRstdeMask  = 0x20000000
)

// GetRstde RSTDE
func (r *RegisterTimdier5Type) GetRstde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldRstdeMask) != 0
}

// SetRstde RSTDE
func (r *RegisterTimdier5Type) SetRstde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldRstdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldRstdeMask)
	}
}

const (
	RegisterTimdier5FieldDlyprtdeShift = 30
	RegisterTimdier5FieldDlyprtdeMask  = 0x40000000
)

// GetDlyprtde DLYPRTDE
func (r *RegisterTimdier5Type) GetDlyprtde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdier5FieldDlyprtdeMask) != 0
}

// SetDlyprtde DLYPRTDE
func (r *RegisterTimdier5Type) SetDlyprtde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdier5FieldDlyprtdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdier5FieldDlyprtdeMask)
	}
}

// RegisterCntrType Timerx Counter Register
type RegisterCntrType uint32

func (r *RegisterCntrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCntrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCntrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCntrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCntrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCntrFieldCntxShift = 0
	RegisterCntrFieldCntxMask  = 0xffff
)

// GetCntx Timerx Counter value
func (r *RegisterCntrType) GetCntx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntrFieldCntxMask) >> RegisterCntrFieldCntxShift)
}

// SetCntx Timerx Counter value
func (r *RegisterCntrType) SetCntx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntrFieldCntxMask)|(uint32(value)<<RegisterCntrFieldCntxShift))
}

// RegisterPerrType Timerx Period Register
type RegisterPerrType uint32

func (r *RegisterPerrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPerrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPerrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPerrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPerrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPerrFieldPerxShift = 0
	RegisterPerrFieldPerxMask  = 0xffff
)

// GetPerx Timerx Period value
func (r *RegisterPerrType) GetPerx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPerrFieldPerxMask) >> RegisterPerrFieldPerxShift)
}

// SetPerx Timerx Period value
func (r *RegisterPerrType) SetPerx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPerrFieldPerxMask)|(uint32(value)<<RegisterPerrFieldPerxShift))
}

// RegisterReprType Timerx Repetition Register
type RegisterReprType uint32

func (r *RegisterReprType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterReprType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterReprType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterReprType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterReprType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterReprFieldRepxShift = 0
	RegisterReprFieldRepxMask  = 0xff
)

// GetRepx Timerx Repetition counter value
func (r *RegisterReprType) GetRepx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterReprFieldRepxMask) >> RegisterReprFieldRepxShift)
}

// SetRepx Timerx Repetition counter value
func (r *RegisterReprType) SetRepx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterReprFieldRepxMask)|(uint32(value)<<RegisterReprFieldRepxShift))
}

// RegisterCmp1rType Timerx Compare 1 Register
type RegisterCmp1rType uint32

func (r *RegisterCmp1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCmp1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCmp1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCmp1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCmp1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCmp1rFieldCmp1xShift = 0
	RegisterCmp1rFieldCmp1xMask  = 0xffff
)

// GetCmp1x Timerx Compare 1 value
func (r *RegisterCmp1rType) GetCmp1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1rFieldCmp1xMask) >> RegisterCmp1rFieldCmp1xShift)
}

// SetCmp1x Timerx Compare 1 value
func (r *RegisterCmp1rType) SetCmp1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1rFieldCmp1xMask)|(uint32(value)<<RegisterCmp1rFieldCmp1xShift))
}

// RegisterCmp1crType Timerx Compare 1 Compound Register
type RegisterCmp1crType uint32

func (r *RegisterCmp1crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCmp1crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCmp1crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCmp1crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCmp1crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCmp1crFieldCmp1xShift = 0
	RegisterCmp1crFieldCmp1xMask  = 0xffff
)

// GetCmp1x Timerx Compare 1 value
func (r *RegisterCmp1crType) GetCmp1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1crFieldCmp1xMask) >> RegisterCmp1crFieldCmp1xShift)
}

// SetCmp1x Timerx Compare 1 value
func (r *RegisterCmp1crType) SetCmp1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1crFieldCmp1xMask)|(uint32(value)<<RegisterCmp1crFieldCmp1xShift))
}

const (
	RegisterCmp1crFieldRepxShift = 16
	RegisterCmp1crFieldRepxMask  = 0xff0000
)

// GetRepx Timerx Repetition value (aliased from HRTIM_REPx register)
func (r *RegisterCmp1crType) GetRepx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1crFieldRepxMask) >> RegisterCmp1crFieldRepxShift)
}

// SetRepx Timerx Repetition value (aliased from HRTIM_REPx register)
func (r *RegisterCmp1crType) SetRepx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1crFieldRepxMask)|(uint32(value)<<RegisterCmp1crFieldRepxShift))
}

// RegisterCmp2rType Timerx Compare 2 Register
type RegisterCmp2rType uint32

func (r *RegisterCmp2rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCmp2rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCmp2rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCmp2rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCmp2rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCmp2rFieldCmp2xShift = 0
	RegisterCmp2rFieldCmp2xMask  = 0xffff
)

// GetCmp2x Timerx Compare 2 value
func (r *RegisterCmp2rType) GetCmp2x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp2rFieldCmp2xMask) >> RegisterCmp2rFieldCmp2xShift)
}

// SetCmp2x Timerx Compare 2 value
func (r *RegisterCmp2rType) SetCmp2x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp2rFieldCmp2xMask)|(uint32(value)<<RegisterCmp2rFieldCmp2xShift))
}

// RegisterCmp3rType Timerx Compare 3 Register
type RegisterCmp3rType uint32

func (r *RegisterCmp3rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCmp3rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCmp3rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCmp3rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCmp3rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCmp3rFieldCmp3xShift = 0
	RegisterCmp3rFieldCmp3xMask  = 0xffff
)

// GetCmp3x Timerx Compare 3 value
func (r *RegisterCmp3rType) GetCmp3x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp3rFieldCmp3xMask) >> RegisterCmp3rFieldCmp3xShift)
}

// SetCmp3x Timerx Compare 3 value
func (r *RegisterCmp3rType) SetCmp3x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp3rFieldCmp3xMask)|(uint32(value)<<RegisterCmp3rFieldCmp3xShift))
}

// RegisterCmp4rType Timerx Compare 4 Register
type RegisterCmp4rType uint32

func (r *RegisterCmp4rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCmp4rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCmp4rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCmp4rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCmp4rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCmp4rFieldCmp4xShift = 0
	RegisterCmp4rFieldCmp4xMask  = 0xffff
)

// GetCmp4x Timerx Compare 4 value
func (r *RegisterCmp4rType) GetCmp4x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp4rFieldCmp4xMask) >> RegisterCmp4rFieldCmp4xShift)
}

// SetCmp4x Timerx Compare 4 value
func (r *RegisterCmp4rType) SetCmp4x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp4rFieldCmp4xMask)|(uint32(value)<<RegisterCmp4rFieldCmp4xShift))
}

// RegisterCpt1rType Timerx Capture 1 Register
type RegisterCpt1rType uint32

func (r *RegisterCpt1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCpt1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCpt1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCpt1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCpt1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCpt1rFieldCpt1xShift = 0
	RegisterCpt1rFieldCpt1xMask  = 0xffff
)

// GetCpt1x Timerx Capture 1 value
func (r *RegisterCpt1rType) GetCpt1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpt1rFieldCpt1xMask) >> RegisterCpt1rFieldCpt1xShift)
}

// SetCpt1x Timerx Capture 1 value
func (r *RegisterCpt1rType) SetCpt1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpt1rFieldCpt1xMask)|(uint32(value)<<RegisterCpt1rFieldCpt1xShift))
}

// RegisterCpt2rType Timerx Capture 2 Register
type RegisterCpt2rType uint32

func (r *RegisterCpt2rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCpt2rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCpt2rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCpt2rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCpt2rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCpt2rFieldCpt2xShift = 0
	RegisterCpt2rFieldCpt2xMask  = 0xffff
)

// GetCpt2x Timerx Capture 2 value
func (r *RegisterCpt2rType) GetCpt2x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpt2rFieldCpt2xMask) >> RegisterCpt2rFieldCpt2xShift)
}

// SetCpt2x Timerx Capture 2 value
func (r *RegisterCpt2rType) SetCpt2x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpt2rFieldCpt2xMask)|(uint32(value)<<RegisterCpt2rFieldCpt2xShift))
}

// RegisterDtrType Timerx Deadtime Register
type RegisterDtrType uint32

func (r *RegisterDtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDtrFieldDtrxShift = 0
	RegisterDtrFieldDtrxMask  = 0x1ff
)

// GetDtrx Deadtime Rising value
func (r *RegisterDtrType) GetDtrx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDtrFieldDtrxMask) >> RegisterDtrFieldDtrxShift)
}

// SetDtrx Deadtime Rising value
func (r *RegisterDtrType) SetDtrx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtrFieldDtrxMask)|(uint32(value)<<RegisterDtrFieldDtrxShift))
}

const (
	RegisterDtrFieldSdtrxShift = 9
	RegisterDtrFieldSdtrxMask  = 0x200
)

// GetSdtrx Sign Deadtime Rising value
func (r *RegisterDtrType) GetSdtrx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtrFieldSdtrxMask) != 0
}

// SetSdtrx Sign Deadtime Rising value
func (r *RegisterDtrType) SetSdtrx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtrFieldSdtrxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtrFieldSdtrxMask)
	}
}

const (
	RegisterDtrFieldDtprscShift = 10
	RegisterDtrFieldDtprscMask  = 0x1c00
)

// GetDtprsc Deadtime Prescaler
func (r *RegisterDtrType) GetDtprsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDtrFieldDtprscMask) >> RegisterDtrFieldDtprscShift)
}

// SetDtprsc Deadtime Prescaler
func (r *RegisterDtrType) SetDtprsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtrFieldDtprscMask)|(uint32(value)<<RegisterDtrFieldDtprscShift))
}

const (
	RegisterDtrFieldDtrslkxShift = 14
	RegisterDtrFieldDtrslkxMask  = 0x4000
)

// GetDtrslkx Deadtime Rising Sign Lock
func (r *RegisterDtrType) GetDtrslkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtrFieldDtrslkxMask) != 0
}

// SetDtrslkx Deadtime Rising Sign Lock
func (r *RegisterDtrType) SetDtrslkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtrFieldDtrslkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtrFieldDtrslkxMask)
	}
}

const (
	RegisterDtrFieldDtrlkxShift = 15
	RegisterDtrFieldDtrlkxMask  = 0x8000
)

// GetDtrlkx Deadtime Rising Lock
func (r *RegisterDtrType) GetDtrlkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtrFieldDtrlkxMask) != 0
}

// SetDtrlkx Deadtime Rising Lock
func (r *RegisterDtrType) SetDtrlkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtrFieldDtrlkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtrFieldDtrlkxMask)
	}
}

const (
	RegisterDtrFieldDtfxShift = 16
	RegisterDtrFieldDtfxMask  = 0x1ff0000
)

// GetDtfx Deadtime Falling value
func (r *RegisterDtrType) GetDtfx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDtrFieldDtfxMask) >> RegisterDtrFieldDtfxShift)
}

// SetDtfx Deadtime Falling value
func (r *RegisterDtrType) SetDtfx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtrFieldDtfxMask)|(uint32(value)<<RegisterDtrFieldDtfxShift))
}

const (
	RegisterDtrFieldSdtfxShift = 25
	RegisterDtrFieldSdtfxMask  = 0x2000000
)

// GetSdtfx Sign Deadtime Falling value
func (r *RegisterDtrType) GetSdtfx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtrFieldSdtfxMask) != 0
}

// SetSdtfx Sign Deadtime Falling value
func (r *RegisterDtrType) SetSdtfx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtrFieldSdtfxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtrFieldSdtfxMask)
	}
}

const (
	RegisterDtrFieldDtfslkxShift = 30
	RegisterDtrFieldDtfslkxMask  = 0x40000000
)

// GetDtfslkx Deadtime Falling Sign Lock
func (r *RegisterDtrType) GetDtfslkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtrFieldDtfslkxMask) != 0
}

// SetDtfslkx Deadtime Falling Sign Lock
func (r *RegisterDtrType) SetDtfslkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtrFieldDtfslkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtrFieldDtfslkxMask)
	}
}

const (
	RegisterDtrFieldDtflkxShift = 31
	RegisterDtrFieldDtflkxMask  = 0x80000000
)

// GetDtflkx Deadtime Falling Lock
func (r *RegisterDtrType) GetDtflkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtrFieldDtflkxMask) != 0
}

// SetDtflkx Deadtime Falling Lock
func (r *RegisterDtrType) SetDtflkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtrFieldDtflkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtrFieldDtflkxMask)
	}
}

// RegisterSet1rType Timerx Output1 Set Register
type RegisterSet1rType uint32

func (r *RegisterSet1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSet1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSet1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSet1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSet1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSet1rFieldSstShift = 0
	RegisterSet1rFieldSstMask  = 0x1
)

// GetSst Software Set trigger
func (r *RegisterSet1rType) GetSst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldSstMask) != 0
}

// SetSst Software Set trigger
func (r *RegisterSet1rType) SetSst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldSstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldSstMask)
	}
}

const (
	RegisterSet1rFieldResyncShift = 1
	RegisterSet1rFieldResyncMask  = 0x2
)

// GetResync Timer A resynchronizaton
func (r *RegisterSet1rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldResyncMask) != 0
}

// SetResync Timer A resynchronizaton
func (r *RegisterSet1rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldResyncMask)
	}
}

const (
	RegisterSet1rFieldPerShift = 2
	RegisterSet1rFieldPerMask  = 0x4
)

// GetPer Timer A Period
func (r *RegisterSet1rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldPerMask) != 0
}

// SetPer Timer A Period
func (r *RegisterSet1rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldPerMask)
	}
}

const (
	RegisterSet1rFieldCmp1Shift = 3
	RegisterSet1rFieldCmp1Mask  = 0x8
)

// GetCmp1 Timer A compare 1
func (r *RegisterSet1rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldCmp1Mask) != 0
}

// SetCmp1 Timer A compare 1
func (r *RegisterSet1rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldCmp1Mask)
	}
}

const (
	RegisterSet1rFieldCmp2Shift = 4
	RegisterSet1rFieldCmp2Mask  = 0x10
)

// GetCmp2 Timer A compare 2
func (r *RegisterSet1rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldCmp2Mask) != 0
}

// SetCmp2 Timer A compare 2
func (r *RegisterSet1rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldCmp2Mask)
	}
}

const (
	RegisterSet1rFieldCmp3Shift = 5
	RegisterSet1rFieldCmp3Mask  = 0x20
)

// GetCmp3 Timer A compare 3
func (r *RegisterSet1rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldCmp3Mask) != 0
}

// SetCmp3 Timer A compare 3
func (r *RegisterSet1rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldCmp3Mask)
	}
}

const (
	RegisterSet1rFieldCmp4Shift = 6
	RegisterSet1rFieldCmp4Mask  = 0x40
)

// GetCmp4 Timer A compare 4
func (r *RegisterSet1rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldCmp4Mask) != 0
}

// SetCmp4 Timer A compare 4
func (r *RegisterSet1rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldCmp4Mask)
	}
}

const (
	RegisterSet1rFieldMstperShift = 7
	RegisterSet1rFieldMstperMask  = 0x80
)

// GetMstper Master Period
func (r *RegisterSet1rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldMstperMask) != 0
}

// SetMstper Master Period
func (r *RegisterSet1rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldMstperMask)
	}
}

const (
	RegisterSet1rFieldMstcmp1Shift = 8
	RegisterSet1rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 Master Compare 1
func (r *RegisterSet1rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 Master Compare 1
func (r *RegisterSet1rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldMstcmp1Mask)
	}
}

const (
	RegisterSet1rFieldMstcmp2Shift = 9
	RegisterSet1rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 Master Compare 2
func (r *RegisterSet1rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 Master Compare 2
func (r *RegisterSet1rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldMstcmp2Mask)
	}
}

const (
	RegisterSet1rFieldMstcmp3Shift = 10
	RegisterSet1rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 Master Compare 3
func (r *RegisterSet1rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 Master Compare 3
func (r *RegisterSet1rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldMstcmp3Mask)
	}
}

const (
	RegisterSet1rFieldMstcmp4Shift = 11
	RegisterSet1rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 Master Compare 4
func (r *RegisterSet1rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 Master Compare 4
func (r *RegisterSet1rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldMstcmp4Mask)
	}
}

const (
	RegisterSet1rFieldTimevnt1Shift = 12
	RegisterSet1rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 Timer Event 1
func (r *RegisterSet1rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 Timer Event 1
func (r *RegisterSet1rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldTimevnt1Mask)
	}
}

const (
	RegisterSet1rFieldTimevnt2Shift = 13
	RegisterSet1rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 Timer Event 2
func (r *RegisterSet1rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 Timer Event 2
func (r *RegisterSet1rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldTimevnt2Mask)
	}
}

const (
	RegisterSet1rFieldTimevnt3Shift = 14
	RegisterSet1rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 Timer Event 3
func (r *RegisterSet1rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 Timer Event 3
func (r *RegisterSet1rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldTimevnt3Mask)
	}
}

const (
	RegisterSet1rFieldTimevnt4Shift = 15
	RegisterSet1rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 Timer Event 4
func (r *RegisterSet1rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 Timer Event 4
func (r *RegisterSet1rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldTimevnt4Mask)
	}
}

const (
	RegisterSet1rFieldTimevnt5Shift = 16
	RegisterSet1rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 Timer Event 5
func (r *RegisterSet1rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 Timer Event 5
func (r *RegisterSet1rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldTimevnt5Mask)
	}
}

const (
	RegisterSet1rFieldTimevnt6Shift = 17
	RegisterSet1rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 Timer Event 6
func (r *RegisterSet1rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 Timer Event 6
func (r *RegisterSet1rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldTimevnt6Mask)
	}
}

const (
	RegisterSet1rFieldTimevnt7Shift = 18
	RegisterSet1rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 Timer Event 7
func (r *RegisterSet1rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 Timer Event 7
func (r *RegisterSet1rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldTimevnt7Mask)
	}
}

const (
	RegisterSet1rFieldTimevnt8Shift = 19
	RegisterSet1rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 Timer Event 8
func (r *RegisterSet1rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 Timer Event 8
func (r *RegisterSet1rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldTimevnt8Mask)
	}
}

const (
	RegisterSet1rFieldTimevnt9Shift = 20
	RegisterSet1rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 Timer Event 9
func (r *RegisterSet1rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 Timer Event 9
func (r *RegisterSet1rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldTimevnt9Mask)
	}
}

const (
	RegisterSet1rFieldExtevnt1Shift = 21
	RegisterSet1rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 External Event 1
func (r *RegisterSet1rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 External Event 1
func (r *RegisterSet1rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldExtevnt1Mask)
	}
}

const (
	RegisterSet1rFieldExtevnt2Shift = 22
	RegisterSet1rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 External Event 2
func (r *RegisterSet1rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 External Event 2
func (r *RegisterSet1rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldExtevnt2Mask)
	}
}

const (
	RegisterSet1rFieldExtevnt3Shift = 23
	RegisterSet1rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 External Event 3
func (r *RegisterSet1rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 External Event 3
func (r *RegisterSet1rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldExtevnt3Mask)
	}
}

const (
	RegisterSet1rFieldExtevnt4Shift = 24
	RegisterSet1rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 External Event 4
func (r *RegisterSet1rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 External Event 4
func (r *RegisterSet1rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldExtevnt4Mask)
	}
}

const (
	RegisterSet1rFieldExtevnt5Shift = 25
	RegisterSet1rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 External Event 5
func (r *RegisterSet1rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 External Event 5
func (r *RegisterSet1rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldExtevnt5Mask)
	}
}

const (
	RegisterSet1rFieldExtevnt6Shift = 26
	RegisterSet1rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 External Event 6
func (r *RegisterSet1rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 External Event 6
func (r *RegisterSet1rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldExtevnt6Mask)
	}
}

const (
	RegisterSet1rFieldExtevnt7Shift = 27
	RegisterSet1rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 External Event 7
func (r *RegisterSet1rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 External Event 7
func (r *RegisterSet1rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldExtevnt7Mask)
	}
}

const (
	RegisterSet1rFieldExtevnt8Shift = 28
	RegisterSet1rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 External Event 8
func (r *RegisterSet1rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 External Event 8
func (r *RegisterSet1rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldExtevnt8Mask)
	}
}

const (
	RegisterSet1rFieldExtevnt9Shift = 29
	RegisterSet1rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 External Event 9
func (r *RegisterSet1rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 External Event 9
func (r *RegisterSet1rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldExtevnt9Mask)
	}
}

const (
	RegisterSet1rFieldExtevnt10Shift = 30
	RegisterSet1rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 External Event 10
func (r *RegisterSet1rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 External Event 10
func (r *RegisterSet1rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldExtevnt10Mask)
	}
}

const (
	RegisterSet1rFieldUpdateShift = 31
	RegisterSet1rFieldUpdateMask  = 0x80000000
)

// GetUpdate Registers update (transfer preload to active)
func (r *RegisterSet1rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet1rFieldUpdateMask) != 0
}

// SetUpdate Registers update (transfer preload to active)
func (r *RegisterSet1rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet1rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet1rFieldUpdateMask)
	}
}

// RegisterRst1rType Timerx Output1 Reset Register
type RegisterRst1rType uint32

func (r *RegisterRst1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRst1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRst1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRst1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRst1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRst1rFieldSrtShift = 0
	RegisterRst1rFieldSrtMask  = 0x1
)

// GetSrt SRT
func (r *RegisterRst1rType) GetSrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldSrtMask) != 0
}

// SetSrt SRT
func (r *RegisterRst1rType) SetSrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldSrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldSrtMask)
	}
}

const (
	RegisterRst1rFieldResyncShift = 1
	RegisterRst1rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *RegisterRst1rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *RegisterRst1rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldResyncMask)
	}
}

const (
	RegisterRst1rFieldPerShift = 2
	RegisterRst1rFieldPerMask  = 0x4
)

// GetPer PER
func (r *RegisterRst1rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldPerMask) != 0
}

// SetPer PER
func (r *RegisterRst1rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldPerMask)
	}
}

const (
	RegisterRst1rFieldCmp1Shift = 3
	RegisterRst1rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *RegisterRst1rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *RegisterRst1rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldCmp1Mask)
	}
}

const (
	RegisterRst1rFieldCmp2Shift = 4
	RegisterRst1rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *RegisterRst1rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *RegisterRst1rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldCmp2Mask)
	}
}

const (
	RegisterRst1rFieldCmp3Shift = 5
	RegisterRst1rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *RegisterRst1rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *RegisterRst1rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldCmp3Mask)
	}
}

const (
	RegisterRst1rFieldCmp4Shift = 6
	RegisterRst1rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *RegisterRst1rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *RegisterRst1rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldCmp4Mask)
	}
}

const (
	RegisterRst1rFieldMstperShift = 7
	RegisterRst1rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *RegisterRst1rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *RegisterRst1rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldMstperMask)
	}
}

const (
	RegisterRst1rFieldMstcmp1Shift = 8
	RegisterRst1rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *RegisterRst1rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *RegisterRst1rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldMstcmp1Mask)
	}
}

const (
	RegisterRst1rFieldMstcmp2Shift = 9
	RegisterRst1rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *RegisterRst1rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *RegisterRst1rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldMstcmp2Mask)
	}
}

const (
	RegisterRst1rFieldMstcmp3Shift = 10
	RegisterRst1rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *RegisterRst1rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *RegisterRst1rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldMstcmp3Mask)
	}
}

const (
	RegisterRst1rFieldMstcmp4Shift = 11
	RegisterRst1rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *RegisterRst1rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *RegisterRst1rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldMstcmp4Mask)
	}
}

const (
	RegisterRst1rFieldTimevnt1Shift = 12
	RegisterRst1rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *RegisterRst1rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *RegisterRst1rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldTimevnt1Mask)
	}
}

const (
	RegisterRst1rFieldTimevnt2Shift = 13
	RegisterRst1rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *RegisterRst1rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *RegisterRst1rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldTimevnt2Mask)
	}
}

const (
	RegisterRst1rFieldTimevnt3Shift = 14
	RegisterRst1rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *RegisterRst1rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *RegisterRst1rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldTimevnt3Mask)
	}
}

const (
	RegisterRst1rFieldTimevnt4Shift = 15
	RegisterRst1rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *RegisterRst1rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *RegisterRst1rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldTimevnt4Mask)
	}
}

const (
	RegisterRst1rFieldTimevnt5Shift = 16
	RegisterRst1rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *RegisterRst1rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *RegisterRst1rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldTimevnt5Mask)
	}
}

const (
	RegisterRst1rFieldTimevnt6Shift = 17
	RegisterRst1rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *RegisterRst1rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *RegisterRst1rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldTimevnt6Mask)
	}
}

const (
	RegisterRst1rFieldTimevnt7Shift = 18
	RegisterRst1rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *RegisterRst1rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *RegisterRst1rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldTimevnt7Mask)
	}
}

const (
	RegisterRst1rFieldTimevnt8Shift = 19
	RegisterRst1rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *RegisterRst1rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *RegisterRst1rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldTimevnt8Mask)
	}
}

const (
	RegisterRst1rFieldTimevnt9Shift = 20
	RegisterRst1rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *RegisterRst1rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *RegisterRst1rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldTimevnt9Mask)
	}
}

const (
	RegisterRst1rFieldExtevnt1Shift = 21
	RegisterRst1rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *RegisterRst1rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *RegisterRst1rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldExtevnt1Mask)
	}
}

const (
	RegisterRst1rFieldExtevnt2Shift = 22
	RegisterRst1rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *RegisterRst1rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *RegisterRst1rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldExtevnt2Mask)
	}
}

const (
	RegisterRst1rFieldExtevnt3Shift = 23
	RegisterRst1rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *RegisterRst1rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *RegisterRst1rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldExtevnt3Mask)
	}
}

const (
	RegisterRst1rFieldExtevnt4Shift = 24
	RegisterRst1rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *RegisterRst1rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *RegisterRst1rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldExtevnt4Mask)
	}
}

const (
	RegisterRst1rFieldExtevnt5Shift = 25
	RegisterRst1rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *RegisterRst1rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *RegisterRst1rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldExtevnt5Mask)
	}
}

const (
	RegisterRst1rFieldExtevnt6Shift = 26
	RegisterRst1rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *RegisterRst1rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *RegisterRst1rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldExtevnt6Mask)
	}
}

const (
	RegisterRst1rFieldExtevnt7Shift = 27
	RegisterRst1rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *RegisterRst1rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *RegisterRst1rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldExtevnt7Mask)
	}
}

const (
	RegisterRst1rFieldExtevnt8Shift = 28
	RegisterRst1rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *RegisterRst1rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *RegisterRst1rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldExtevnt8Mask)
	}
}

const (
	RegisterRst1rFieldExtevnt9Shift = 29
	RegisterRst1rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *RegisterRst1rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *RegisterRst1rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldExtevnt9Mask)
	}
}

const (
	RegisterRst1rFieldExtevnt10Shift = 30
	RegisterRst1rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *RegisterRst1rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *RegisterRst1rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldExtevnt10Mask)
	}
}

const (
	RegisterRst1rFieldUpdateShift = 31
	RegisterRst1rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *RegisterRst1rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst1rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *RegisterRst1rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst1rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst1rFieldUpdateMask)
	}
}

// RegisterSet2rType Timerx Output2 Set Register
type RegisterSet2rType uint32

func (r *RegisterSet2rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSet2rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSet2rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSet2rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSet2rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSet2rFieldSstShift = 0
	RegisterSet2rFieldSstMask  = 0x1
)

// GetSst SST
func (r *RegisterSet2rType) GetSst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldSstMask) != 0
}

// SetSst SST
func (r *RegisterSet2rType) SetSst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldSstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldSstMask)
	}
}

const (
	RegisterSet2rFieldResyncShift = 1
	RegisterSet2rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *RegisterSet2rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *RegisterSet2rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldResyncMask)
	}
}

const (
	RegisterSet2rFieldPerShift = 2
	RegisterSet2rFieldPerMask  = 0x4
)

// GetPer PER
func (r *RegisterSet2rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldPerMask) != 0
}

// SetPer PER
func (r *RegisterSet2rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldPerMask)
	}
}

const (
	RegisterSet2rFieldCmp1Shift = 3
	RegisterSet2rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *RegisterSet2rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *RegisterSet2rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldCmp1Mask)
	}
}

const (
	RegisterSet2rFieldCmp2Shift = 4
	RegisterSet2rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *RegisterSet2rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *RegisterSet2rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldCmp2Mask)
	}
}

const (
	RegisterSet2rFieldCmp3Shift = 5
	RegisterSet2rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *RegisterSet2rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *RegisterSet2rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldCmp3Mask)
	}
}

const (
	RegisterSet2rFieldCmp4Shift = 6
	RegisterSet2rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *RegisterSet2rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *RegisterSet2rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldCmp4Mask)
	}
}

const (
	RegisterSet2rFieldMstperShift = 7
	RegisterSet2rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *RegisterSet2rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *RegisterSet2rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldMstperMask)
	}
}

const (
	RegisterSet2rFieldMstcmp1Shift = 8
	RegisterSet2rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *RegisterSet2rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *RegisterSet2rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldMstcmp1Mask)
	}
}

const (
	RegisterSet2rFieldMstcmp2Shift = 9
	RegisterSet2rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *RegisterSet2rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *RegisterSet2rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldMstcmp2Mask)
	}
}

const (
	RegisterSet2rFieldMstcmp3Shift = 10
	RegisterSet2rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *RegisterSet2rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *RegisterSet2rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldMstcmp3Mask)
	}
}

const (
	RegisterSet2rFieldMstcmp4Shift = 11
	RegisterSet2rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *RegisterSet2rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *RegisterSet2rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldMstcmp4Mask)
	}
}

const (
	RegisterSet2rFieldTimevnt1Shift = 12
	RegisterSet2rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *RegisterSet2rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *RegisterSet2rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldTimevnt1Mask)
	}
}

const (
	RegisterSet2rFieldTimevnt2Shift = 13
	RegisterSet2rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *RegisterSet2rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *RegisterSet2rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldTimevnt2Mask)
	}
}

const (
	RegisterSet2rFieldTimevnt3Shift = 14
	RegisterSet2rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *RegisterSet2rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *RegisterSet2rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldTimevnt3Mask)
	}
}

const (
	RegisterSet2rFieldTimevnt4Shift = 15
	RegisterSet2rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *RegisterSet2rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *RegisterSet2rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldTimevnt4Mask)
	}
}

const (
	RegisterSet2rFieldTimevnt5Shift = 16
	RegisterSet2rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *RegisterSet2rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *RegisterSet2rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldTimevnt5Mask)
	}
}

const (
	RegisterSet2rFieldTimevnt6Shift = 17
	RegisterSet2rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *RegisterSet2rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *RegisterSet2rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldTimevnt6Mask)
	}
}

const (
	RegisterSet2rFieldTimevnt7Shift = 18
	RegisterSet2rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *RegisterSet2rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *RegisterSet2rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldTimevnt7Mask)
	}
}

const (
	RegisterSet2rFieldTimevnt8Shift = 19
	RegisterSet2rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *RegisterSet2rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *RegisterSet2rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldTimevnt8Mask)
	}
}

const (
	RegisterSet2rFieldTimevnt9Shift = 20
	RegisterSet2rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *RegisterSet2rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *RegisterSet2rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldTimevnt9Mask)
	}
}

const (
	RegisterSet2rFieldExtevnt1Shift = 21
	RegisterSet2rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *RegisterSet2rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *RegisterSet2rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldExtevnt1Mask)
	}
}

const (
	RegisterSet2rFieldExtevnt2Shift = 22
	RegisterSet2rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *RegisterSet2rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *RegisterSet2rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldExtevnt2Mask)
	}
}

const (
	RegisterSet2rFieldExtevnt3Shift = 23
	RegisterSet2rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *RegisterSet2rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *RegisterSet2rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldExtevnt3Mask)
	}
}

const (
	RegisterSet2rFieldExtevnt4Shift = 24
	RegisterSet2rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *RegisterSet2rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *RegisterSet2rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldExtevnt4Mask)
	}
}

const (
	RegisterSet2rFieldExtevnt5Shift = 25
	RegisterSet2rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *RegisterSet2rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *RegisterSet2rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldExtevnt5Mask)
	}
}

const (
	RegisterSet2rFieldExtevnt6Shift = 26
	RegisterSet2rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *RegisterSet2rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *RegisterSet2rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldExtevnt6Mask)
	}
}

const (
	RegisterSet2rFieldExtevnt7Shift = 27
	RegisterSet2rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *RegisterSet2rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *RegisterSet2rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldExtevnt7Mask)
	}
}

const (
	RegisterSet2rFieldExtevnt8Shift = 28
	RegisterSet2rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *RegisterSet2rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *RegisterSet2rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldExtevnt8Mask)
	}
}

const (
	RegisterSet2rFieldExtevnt9Shift = 29
	RegisterSet2rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *RegisterSet2rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *RegisterSet2rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldExtevnt9Mask)
	}
}

const (
	RegisterSet2rFieldExtevnt10Shift = 30
	RegisterSet2rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *RegisterSet2rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *RegisterSet2rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldExtevnt10Mask)
	}
}

const (
	RegisterSet2rFieldUpdateShift = 31
	RegisterSet2rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *RegisterSet2rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSet2rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *RegisterSet2rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSet2rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSet2rFieldUpdateMask)
	}
}

// RegisterRst2rType Timerx Output2 Reset Register
type RegisterRst2rType uint32

func (r *RegisterRst2rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRst2rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRst2rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRst2rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRst2rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRst2rFieldSrtShift = 0
	RegisterRst2rFieldSrtMask  = 0x1
)

// GetSrt SRT
func (r *RegisterRst2rType) GetSrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldSrtMask) != 0
}

// SetSrt SRT
func (r *RegisterRst2rType) SetSrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldSrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldSrtMask)
	}
}

const (
	RegisterRst2rFieldResyncShift = 1
	RegisterRst2rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *RegisterRst2rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *RegisterRst2rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldResyncMask)
	}
}

const (
	RegisterRst2rFieldPerShift = 2
	RegisterRst2rFieldPerMask  = 0x4
)

// GetPer PER
func (r *RegisterRst2rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldPerMask) != 0
}

// SetPer PER
func (r *RegisterRst2rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldPerMask)
	}
}

const (
	RegisterRst2rFieldCmp1Shift = 3
	RegisterRst2rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *RegisterRst2rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *RegisterRst2rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldCmp1Mask)
	}
}

const (
	RegisterRst2rFieldCmp2Shift = 4
	RegisterRst2rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *RegisterRst2rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *RegisterRst2rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldCmp2Mask)
	}
}

const (
	RegisterRst2rFieldCmp3Shift = 5
	RegisterRst2rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *RegisterRst2rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *RegisterRst2rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldCmp3Mask)
	}
}

const (
	RegisterRst2rFieldCmp4Shift = 6
	RegisterRst2rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *RegisterRst2rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *RegisterRst2rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldCmp4Mask)
	}
}

const (
	RegisterRst2rFieldMstperShift = 7
	RegisterRst2rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *RegisterRst2rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *RegisterRst2rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldMstperMask)
	}
}

const (
	RegisterRst2rFieldMstcmp1Shift = 8
	RegisterRst2rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *RegisterRst2rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *RegisterRst2rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldMstcmp1Mask)
	}
}

const (
	RegisterRst2rFieldMstcmp2Shift = 9
	RegisterRst2rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *RegisterRst2rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *RegisterRst2rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldMstcmp2Mask)
	}
}

const (
	RegisterRst2rFieldMstcmp3Shift = 10
	RegisterRst2rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *RegisterRst2rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *RegisterRst2rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldMstcmp3Mask)
	}
}

const (
	RegisterRst2rFieldMstcmp4Shift = 11
	RegisterRst2rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *RegisterRst2rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *RegisterRst2rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldMstcmp4Mask)
	}
}

const (
	RegisterRst2rFieldTimevnt1Shift = 12
	RegisterRst2rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *RegisterRst2rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *RegisterRst2rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldTimevnt1Mask)
	}
}

const (
	RegisterRst2rFieldTimevnt2Shift = 13
	RegisterRst2rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *RegisterRst2rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *RegisterRst2rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldTimevnt2Mask)
	}
}

const (
	RegisterRst2rFieldTimevnt3Shift = 14
	RegisterRst2rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *RegisterRst2rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *RegisterRst2rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldTimevnt3Mask)
	}
}

const (
	RegisterRst2rFieldTimevnt4Shift = 15
	RegisterRst2rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *RegisterRst2rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *RegisterRst2rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldTimevnt4Mask)
	}
}

const (
	RegisterRst2rFieldTimevnt5Shift = 16
	RegisterRst2rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *RegisterRst2rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *RegisterRst2rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldTimevnt5Mask)
	}
}

const (
	RegisterRst2rFieldTimevnt6Shift = 17
	RegisterRst2rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *RegisterRst2rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *RegisterRst2rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldTimevnt6Mask)
	}
}

const (
	RegisterRst2rFieldTimevnt7Shift = 18
	RegisterRst2rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *RegisterRst2rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *RegisterRst2rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldTimevnt7Mask)
	}
}

const (
	RegisterRst2rFieldTimevnt8Shift = 19
	RegisterRst2rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *RegisterRst2rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *RegisterRst2rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldTimevnt8Mask)
	}
}

const (
	RegisterRst2rFieldTimevnt9Shift = 20
	RegisterRst2rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *RegisterRst2rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *RegisterRst2rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldTimevnt9Mask)
	}
}

const (
	RegisterRst2rFieldExtevnt1Shift = 21
	RegisterRst2rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *RegisterRst2rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *RegisterRst2rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldExtevnt1Mask)
	}
}

const (
	RegisterRst2rFieldExtevnt2Shift = 22
	RegisterRst2rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *RegisterRst2rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *RegisterRst2rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldExtevnt2Mask)
	}
}

const (
	RegisterRst2rFieldExtevnt3Shift = 23
	RegisterRst2rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *RegisterRst2rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *RegisterRst2rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldExtevnt3Mask)
	}
}

const (
	RegisterRst2rFieldExtevnt4Shift = 24
	RegisterRst2rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *RegisterRst2rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *RegisterRst2rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldExtevnt4Mask)
	}
}

const (
	RegisterRst2rFieldExtevnt5Shift = 25
	RegisterRst2rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *RegisterRst2rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *RegisterRst2rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldExtevnt5Mask)
	}
}

const (
	RegisterRst2rFieldExtevnt6Shift = 26
	RegisterRst2rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *RegisterRst2rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *RegisterRst2rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldExtevnt6Mask)
	}
}

const (
	RegisterRst2rFieldExtevnt7Shift = 27
	RegisterRst2rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *RegisterRst2rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *RegisterRst2rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldExtevnt7Mask)
	}
}

const (
	RegisterRst2rFieldExtevnt8Shift = 28
	RegisterRst2rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *RegisterRst2rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *RegisterRst2rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldExtevnt8Mask)
	}
}

const (
	RegisterRst2rFieldExtevnt9Shift = 29
	RegisterRst2rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *RegisterRst2rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *RegisterRst2rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldExtevnt9Mask)
	}
}

const (
	RegisterRst2rFieldExtevnt10Shift = 30
	RegisterRst2rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *RegisterRst2rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *RegisterRst2rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldExtevnt10Mask)
	}
}

const (
	RegisterRst2rFieldUpdateShift = 31
	RegisterRst2rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *RegisterRst2rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRst2rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *RegisterRst2rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRst2rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRst2rFieldUpdateMask)
	}
}

// RegisterEefr1Type Timerx External Event Filtering Register 1
type RegisterEefr1Type uint32

func (r *RegisterEefr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterEefr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterEefr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterEefr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterEefr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterEefr1FieldEe1ltchShift = 0
	RegisterEefr1FieldEe1ltchMask  = 0x1
)

// GetEe1ltch External Event 1 latch
func (r *RegisterEefr1Type) GetEe1ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefr1FieldEe1ltchMask) != 0
}

// SetEe1ltch External Event 1 latch
func (r *RegisterEefr1Type) SetEe1ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefr1FieldEe1ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefr1FieldEe1ltchMask)
	}
}

const (
	RegisterEefr1FieldEe1fltrShift = 1
	RegisterEefr1FieldEe1fltrMask  = 0x1e
)

// GetEe1fltr External Event 1 filter
func (r *RegisterEefr1Type) GetEe1fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefr1FieldEe1fltrMask) >> RegisterEefr1FieldEe1fltrShift)
}

// SetEe1fltr External Event 1 filter
func (r *RegisterEefr1Type) SetEe1fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefr1FieldEe1fltrMask)|(uint32(value)<<RegisterEefr1FieldEe1fltrShift))
}

const (
	RegisterEefr1FieldEe2ltchShift = 6
	RegisterEefr1FieldEe2ltchMask  = 0x40
)

// GetEe2ltch External Event 2 latch
func (r *RegisterEefr1Type) GetEe2ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefr1FieldEe2ltchMask) != 0
}

// SetEe2ltch External Event 2 latch
func (r *RegisterEefr1Type) SetEe2ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefr1FieldEe2ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefr1FieldEe2ltchMask)
	}
}

const (
	RegisterEefr1FieldEe2fltrShift = 7
	RegisterEefr1FieldEe2fltrMask  = 0x780
)

// GetEe2fltr External Event 2 filter
func (r *RegisterEefr1Type) GetEe2fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefr1FieldEe2fltrMask) >> RegisterEefr1FieldEe2fltrShift)
}

// SetEe2fltr External Event 2 filter
func (r *RegisterEefr1Type) SetEe2fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefr1FieldEe2fltrMask)|(uint32(value)<<RegisterEefr1FieldEe2fltrShift))
}

const (
	RegisterEefr1FieldEe3ltchShift = 12
	RegisterEefr1FieldEe3ltchMask  = 0x1000
)

// GetEe3ltch External Event 3 latch
func (r *RegisterEefr1Type) GetEe3ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefr1FieldEe3ltchMask) != 0
}

// SetEe3ltch External Event 3 latch
func (r *RegisterEefr1Type) SetEe3ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefr1FieldEe3ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefr1FieldEe3ltchMask)
	}
}

const (
	RegisterEefr1FieldEe3fltrShift = 13
	RegisterEefr1FieldEe3fltrMask  = 0x1e000
)

// GetEe3fltr External Event 3 filter
func (r *RegisterEefr1Type) GetEe3fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefr1FieldEe3fltrMask) >> RegisterEefr1FieldEe3fltrShift)
}

// SetEe3fltr External Event 3 filter
func (r *RegisterEefr1Type) SetEe3fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefr1FieldEe3fltrMask)|(uint32(value)<<RegisterEefr1FieldEe3fltrShift))
}

const (
	RegisterEefr1FieldEe4ltchShift = 18
	RegisterEefr1FieldEe4ltchMask  = 0x40000
)

// GetEe4ltch External Event 4 latch
func (r *RegisterEefr1Type) GetEe4ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefr1FieldEe4ltchMask) != 0
}

// SetEe4ltch External Event 4 latch
func (r *RegisterEefr1Type) SetEe4ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefr1FieldEe4ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefr1FieldEe4ltchMask)
	}
}

const (
	RegisterEefr1FieldEe4fltrShift = 19
	RegisterEefr1FieldEe4fltrMask  = 0x780000
)

// GetEe4fltr External Event 4 filter
func (r *RegisterEefr1Type) GetEe4fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefr1FieldEe4fltrMask) >> RegisterEefr1FieldEe4fltrShift)
}

// SetEe4fltr External Event 4 filter
func (r *RegisterEefr1Type) SetEe4fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefr1FieldEe4fltrMask)|(uint32(value)<<RegisterEefr1FieldEe4fltrShift))
}

const (
	RegisterEefr1FieldEe5ltchShift = 24
	RegisterEefr1FieldEe5ltchMask  = 0x1000000
)

// GetEe5ltch External Event 5 latch
func (r *RegisterEefr1Type) GetEe5ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefr1FieldEe5ltchMask) != 0
}

// SetEe5ltch External Event 5 latch
func (r *RegisterEefr1Type) SetEe5ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefr1FieldEe5ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefr1FieldEe5ltchMask)
	}
}

const (
	RegisterEefr1FieldEe5fltrShift = 25
	RegisterEefr1FieldEe5fltrMask  = 0x1e000000
)

// GetEe5fltr External Event 5 filter
func (r *RegisterEefr1Type) GetEe5fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefr1FieldEe5fltrMask) >> RegisterEefr1FieldEe5fltrShift)
}

// SetEe5fltr External Event 5 filter
func (r *RegisterEefr1Type) SetEe5fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefr1FieldEe5fltrMask)|(uint32(value)<<RegisterEefr1FieldEe5fltrShift))
}

// RegisterEefr2Type Timerx External Event Filtering Register 2
type RegisterEefr2Type uint32

func (r *RegisterEefr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterEefr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterEefr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterEefr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterEefr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterEefr2FieldEe6ltchShift = 0
	RegisterEefr2FieldEe6ltchMask  = 0x1
)

// GetEe6ltch External Event 6 latch
func (r *RegisterEefr2Type) GetEe6ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefr2FieldEe6ltchMask) != 0
}

// SetEe6ltch External Event 6 latch
func (r *RegisterEefr2Type) SetEe6ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefr2FieldEe6ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefr2FieldEe6ltchMask)
	}
}

const (
	RegisterEefr2FieldEe6fltrShift = 1
	RegisterEefr2FieldEe6fltrMask  = 0x1e
)

// GetEe6fltr External Event 6 filter
func (r *RegisterEefr2Type) GetEe6fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefr2FieldEe6fltrMask) >> RegisterEefr2FieldEe6fltrShift)
}

// SetEe6fltr External Event 6 filter
func (r *RegisterEefr2Type) SetEe6fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefr2FieldEe6fltrMask)|(uint32(value)<<RegisterEefr2FieldEe6fltrShift))
}

const (
	RegisterEefr2FieldEe7ltchShift = 6
	RegisterEefr2FieldEe7ltchMask  = 0x40
)

// GetEe7ltch External Event 7 latch
func (r *RegisterEefr2Type) GetEe7ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefr2FieldEe7ltchMask) != 0
}

// SetEe7ltch External Event 7 latch
func (r *RegisterEefr2Type) SetEe7ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefr2FieldEe7ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefr2FieldEe7ltchMask)
	}
}

const (
	RegisterEefr2FieldEe7fltrShift = 7
	RegisterEefr2FieldEe7fltrMask  = 0x780
)

// GetEe7fltr External Event 7 filter
func (r *RegisterEefr2Type) GetEe7fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefr2FieldEe7fltrMask) >> RegisterEefr2FieldEe7fltrShift)
}

// SetEe7fltr External Event 7 filter
func (r *RegisterEefr2Type) SetEe7fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefr2FieldEe7fltrMask)|(uint32(value)<<RegisterEefr2FieldEe7fltrShift))
}

const (
	RegisterEefr2FieldEe8ltchShift = 12
	RegisterEefr2FieldEe8ltchMask  = 0x1000
)

// GetEe8ltch External Event 8 latch
func (r *RegisterEefr2Type) GetEe8ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefr2FieldEe8ltchMask) != 0
}

// SetEe8ltch External Event 8 latch
func (r *RegisterEefr2Type) SetEe8ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefr2FieldEe8ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefr2FieldEe8ltchMask)
	}
}

const (
	RegisterEefr2FieldEe8fltrShift = 13
	RegisterEefr2FieldEe8fltrMask  = 0x1e000
)

// GetEe8fltr External Event 8 filter
func (r *RegisterEefr2Type) GetEe8fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefr2FieldEe8fltrMask) >> RegisterEefr2FieldEe8fltrShift)
}

// SetEe8fltr External Event 8 filter
func (r *RegisterEefr2Type) SetEe8fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefr2FieldEe8fltrMask)|(uint32(value)<<RegisterEefr2FieldEe8fltrShift))
}

const (
	RegisterEefr2FieldEe9ltchShift = 18
	RegisterEefr2FieldEe9ltchMask  = 0x40000
)

// GetEe9ltch External Event 9 latch
func (r *RegisterEefr2Type) GetEe9ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefr2FieldEe9ltchMask) != 0
}

// SetEe9ltch External Event 9 latch
func (r *RegisterEefr2Type) SetEe9ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefr2FieldEe9ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefr2FieldEe9ltchMask)
	}
}

const (
	RegisterEefr2FieldEe9fltrShift = 19
	RegisterEefr2FieldEe9fltrMask  = 0x780000
)

// GetEe9fltr External Event 9 filter
func (r *RegisterEefr2Type) GetEe9fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefr2FieldEe9fltrMask) >> RegisterEefr2FieldEe9fltrShift)
}

// SetEe9fltr External Event 9 filter
func (r *RegisterEefr2Type) SetEe9fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefr2FieldEe9fltrMask)|(uint32(value)<<RegisterEefr2FieldEe9fltrShift))
}

const (
	RegisterEefr2FieldEe10ltchShift = 24
	RegisterEefr2FieldEe10ltchMask  = 0x1000000
)

// GetEe10ltch External Event 10 latch
func (r *RegisterEefr2Type) GetEe10ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefr2FieldEe10ltchMask) != 0
}

// SetEe10ltch External Event 10 latch
func (r *RegisterEefr2Type) SetEe10ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefr2FieldEe10ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefr2FieldEe10ltchMask)
	}
}

const (
	RegisterEefr2FieldEe10fltrShift = 25
	RegisterEefr2FieldEe10fltrMask  = 0x1e000000
)

// GetEe10fltr External Event 10 filter
func (r *RegisterEefr2Type) GetEe10fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefr2FieldEe10fltrMask) >> RegisterEefr2FieldEe10fltrShift)
}

// SetEe10fltr External Event 10 filter
func (r *RegisterEefr2Type) SetEe10fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefr2FieldEe10fltrMask)|(uint32(value)<<RegisterEefr2FieldEe10fltrShift))
}

// RegisterRstrType TimerA Reset Register
type RegisterRstrType uint32

func (r *RegisterRstrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRstrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRstrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRstrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRstrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRstrFieldUpdtShift = 1
	RegisterRstrFieldUpdtMask  = 0x2
)

// GetUpdt Timer A Update reset
func (r *RegisterRstrType) GetUpdt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldUpdtMask) != 0
}

// SetUpdt Timer A Update reset
func (r *RegisterRstrType) SetUpdt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldUpdtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldUpdtMask)
	}
}

const (
	RegisterRstrFieldCmp2Shift = 2
	RegisterRstrFieldCmp2Mask  = 0x4
)

// GetCmp2 Timer A compare 2 reset
func (r *RegisterRstrType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldCmp2Mask) != 0
}

// SetCmp2 Timer A compare 2 reset
func (r *RegisterRstrType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldCmp2Mask)
	}
}

const (
	RegisterRstrFieldCmp4Shift = 3
	RegisterRstrFieldCmp4Mask  = 0x8
)

// GetCmp4 Timer A compare 4 reset
func (r *RegisterRstrType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldCmp4Mask) != 0
}

// SetCmp4 Timer A compare 4 reset
func (r *RegisterRstrType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldCmp4Mask)
	}
}

const (
	RegisterRstrFieldMstperShift = 4
	RegisterRstrFieldMstperMask  = 0x10
)

// GetMstper Master timer Period
func (r *RegisterRstrType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldMstperMask) != 0
}

// SetMstper Master timer Period
func (r *RegisterRstrType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldMstperMask)
	}
}

const (
	RegisterRstrFieldMstcmp1Shift = 5
	RegisterRstrFieldMstcmp1Mask  = 0x20
)

// GetMstcmp1 Master compare 1
func (r *RegisterRstrType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldMstcmp1Mask) != 0
}

// SetMstcmp1 Master compare 1
func (r *RegisterRstrType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldMstcmp1Mask)
	}
}

const (
	RegisterRstrFieldMstcmp2Shift = 6
	RegisterRstrFieldMstcmp2Mask  = 0x40
)

// GetMstcmp2 Master compare 2
func (r *RegisterRstrType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldMstcmp2Mask) != 0
}

// SetMstcmp2 Master compare 2
func (r *RegisterRstrType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldMstcmp2Mask)
	}
}

const (
	RegisterRstrFieldMstcmp3Shift = 7
	RegisterRstrFieldMstcmp3Mask  = 0x80
)

// GetMstcmp3 Master compare 3
func (r *RegisterRstrType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldMstcmp3Mask) != 0
}

// SetMstcmp3 Master compare 3
func (r *RegisterRstrType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldMstcmp3Mask)
	}
}

const (
	RegisterRstrFieldMstcmp4Shift = 8
	RegisterRstrFieldMstcmp4Mask  = 0x100
)

// GetMstcmp4 Master compare 4
func (r *RegisterRstrType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldMstcmp4Mask) != 0
}

// SetMstcmp4 Master compare 4
func (r *RegisterRstrType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldMstcmp4Mask)
	}
}

const (
	RegisterRstrFieldExtevnt1Shift = 9
	RegisterRstrFieldExtevnt1Mask  = 0x200
)

// GetExtevnt1 External Event 1
func (r *RegisterRstrType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldExtevnt1Mask) != 0
}

// SetExtevnt1 External Event 1
func (r *RegisterRstrType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldExtevnt1Mask)
	}
}

const (
	RegisterRstrFieldExtevnt2Shift = 10
	RegisterRstrFieldExtevnt2Mask  = 0x400
)

// GetExtevnt2 External Event 2
func (r *RegisterRstrType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldExtevnt2Mask) != 0
}

// SetExtevnt2 External Event 2
func (r *RegisterRstrType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldExtevnt2Mask)
	}
}

const (
	RegisterRstrFieldExtevnt3Shift = 11
	RegisterRstrFieldExtevnt3Mask  = 0x800
)

// GetExtevnt3 External Event 3
func (r *RegisterRstrType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldExtevnt3Mask) != 0
}

// SetExtevnt3 External Event 3
func (r *RegisterRstrType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldExtevnt3Mask)
	}
}

const (
	RegisterRstrFieldExtevnt4Shift = 12
	RegisterRstrFieldExtevnt4Mask  = 0x1000
)

// GetExtevnt4 External Event 4
func (r *RegisterRstrType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldExtevnt4Mask) != 0
}

// SetExtevnt4 External Event 4
func (r *RegisterRstrType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldExtevnt4Mask)
	}
}

const (
	RegisterRstrFieldExtevnt5Shift = 13
	RegisterRstrFieldExtevnt5Mask  = 0x2000
)

// GetExtevnt5 External Event 5
func (r *RegisterRstrType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldExtevnt5Mask) != 0
}

// SetExtevnt5 External Event 5
func (r *RegisterRstrType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldExtevnt5Mask)
	}
}

const (
	RegisterRstrFieldExtevnt6Shift = 14
	RegisterRstrFieldExtevnt6Mask  = 0x4000
)

// GetExtevnt6 External Event 6
func (r *RegisterRstrType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldExtevnt6Mask) != 0
}

// SetExtevnt6 External Event 6
func (r *RegisterRstrType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldExtevnt6Mask)
	}
}

const (
	RegisterRstrFieldExtevnt7Shift = 15
	RegisterRstrFieldExtevnt7Mask  = 0x8000
)

// GetExtevnt7 External Event 7
func (r *RegisterRstrType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldExtevnt7Mask) != 0
}

// SetExtevnt7 External Event 7
func (r *RegisterRstrType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldExtevnt7Mask)
	}
}

const (
	RegisterRstrFieldExtevnt8Shift = 16
	RegisterRstrFieldExtevnt8Mask  = 0x10000
)

// GetExtevnt8 External Event 8
func (r *RegisterRstrType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldExtevnt8Mask) != 0
}

// SetExtevnt8 External Event 8
func (r *RegisterRstrType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldExtevnt8Mask)
	}
}

const (
	RegisterRstrFieldExtevnt9Shift = 17
	RegisterRstrFieldExtevnt9Mask  = 0x20000
)

// GetExtevnt9 External Event 9
func (r *RegisterRstrType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldExtevnt9Mask) != 0
}

// SetExtevnt9 External Event 9
func (r *RegisterRstrType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldExtevnt9Mask)
	}
}

const (
	RegisterRstrFieldExtevnt10Shift = 18
	RegisterRstrFieldExtevnt10Mask  = 0x40000
)

// GetExtevnt10 External Event 10
func (r *RegisterRstrType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldExtevnt10Mask) != 0
}

// SetExtevnt10 External Event 10
func (r *RegisterRstrType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldExtevnt10Mask)
	}
}

const (
	RegisterRstrFieldTimcmp1Shift = 19
	RegisterRstrFieldTimcmp1Mask  = 0x80000
)

// GetTimcmp1 Timer Compare 1
func (r *RegisterRstrType) GetTimcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldTimcmp1Mask) != 0
}

// SetTimcmp1 Timer Compare 1
func (r *RegisterRstrType) SetTimcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldTimcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldTimcmp1Mask)
	}
}

const (
	RegisterRstrFieldTimcmp2Shift = 20
	RegisterRstrFieldTimcmp2Mask  = 0x100000
)

// GetTimcmp2 Timer Compare 2
func (r *RegisterRstrType) GetTimcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldTimcmp2Mask) != 0
}

// SetTimcmp2 Timer Compare 2
func (r *RegisterRstrType) SetTimcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldTimcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldTimcmp2Mask)
	}
}

const (
	RegisterRstrFieldTimcmp4Shift = 21
	RegisterRstrFieldTimcmp4Mask  = 0x200000
)

// GetTimcmp4 Timer Compare 4
func (r *RegisterRstrType) GetTimcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldTimcmp4Mask) != 0
}

// SetTimcmp4 Timer Compare 4
func (r *RegisterRstrType) SetTimcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldTimcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldTimcmp4Mask)
	}
}

const (
	RegisterRstrFieldTimccmp1Shift = 22
	RegisterRstrFieldTimccmp1Mask  = 0x400000
)

// GetTimccmp1 Timer C Compare 1
func (r *RegisterRstrType) GetTimccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldTimccmp1Mask) != 0
}

// SetTimccmp1 Timer C Compare 1
func (r *RegisterRstrType) SetTimccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldTimccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldTimccmp1Mask)
	}
}

const (
	RegisterRstrFieldTimccmp2Shift = 23
	RegisterRstrFieldTimccmp2Mask  = 0x800000
)

// GetTimccmp2 Timer C Compare 2
func (r *RegisterRstrType) GetTimccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldTimccmp2Mask) != 0
}

// SetTimccmp2 Timer C Compare 2
func (r *RegisterRstrType) SetTimccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldTimccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldTimccmp2Mask)
	}
}

const (
	RegisterRstrFieldTimccmp4Shift = 24
	RegisterRstrFieldTimccmp4Mask  = 0x1000000
)

// GetTimccmp4 Timer C Compare 4
func (r *RegisterRstrType) GetTimccmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldTimccmp4Mask) != 0
}

// SetTimccmp4 Timer C Compare 4
func (r *RegisterRstrType) SetTimccmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldTimccmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldTimccmp4Mask)
	}
}

const (
	RegisterRstrFieldTimdcmp1Shift = 25
	RegisterRstrFieldTimdcmp1Mask  = 0x2000000
)

// GetTimdcmp1 Timer D Compare 1
func (r *RegisterRstrType) GetTimdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldTimdcmp1Mask) != 0
}

// SetTimdcmp1 Timer D Compare 1
func (r *RegisterRstrType) SetTimdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldTimdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldTimdcmp1Mask)
	}
}

const (
	RegisterRstrFieldTimdcmp2Shift = 26
	RegisterRstrFieldTimdcmp2Mask  = 0x4000000
)

// GetTimdcmp2 Timer D Compare 2
func (r *RegisterRstrType) GetTimdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldTimdcmp2Mask) != 0
}

// SetTimdcmp2 Timer D Compare 2
func (r *RegisterRstrType) SetTimdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldTimdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldTimdcmp2Mask)
	}
}

const (
	RegisterRstrFieldTimdcmp4Shift = 27
	RegisterRstrFieldTimdcmp4Mask  = 0x8000000
)

// GetTimdcmp4 Timer D Compare 4
func (r *RegisterRstrType) GetTimdcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldTimdcmp4Mask) != 0
}

// SetTimdcmp4 Timer D Compare 4
func (r *RegisterRstrType) SetTimdcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldTimdcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldTimdcmp4Mask)
	}
}

const (
	RegisterRstrFieldTimecmp1Shift = 28
	RegisterRstrFieldTimecmp1Mask  = 0x10000000
)

// GetTimecmp1 Timer E Compare 1
func (r *RegisterRstrType) GetTimecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldTimecmp1Mask) != 0
}

// SetTimecmp1 Timer E Compare 1
func (r *RegisterRstrType) SetTimecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldTimecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldTimecmp1Mask)
	}
}

const (
	RegisterRstrFieldTimecmp2Shift = 29
	RegisterRstrFieldTimecmp2Mask  = 0x20000000
)

// GetTimecmp2 Timer E Compare 2
func (r *RegisterRstrType) GetTimecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldTimecmp2Mask) != 0
}

// SetTimecmp2 Timer E Compare 2
func (r *RegisterRstrType) SetTimecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldTimecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldTimecmp2Mask)
	}
}

const (
	RegisterRstrFieldTimecmp4Shift = 30
	RegisterRstrFieldTimecmp4Mask  = 0x40000000
)

// GetTimecmp4 Timer E Compare 4
func (r *RegisterRstrType) GetTimecmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstrFieldTimecmp4Mask) != 0
}

// SetTimecmp4 Timer E Compare 4
func (r *RegisterRstrType) SetTimecmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstrFieldTimecmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstrFieldTimecmp4Mask)
	}
}

// RegisterChprType Timerx Chopper Register
type RegisterChprType uint32

func (r *RegisterChprType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterChprType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterChprType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterChprType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterChprType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterChprFieldChpfrqShift = 0
	RegisterChprFieldChpfrqMask  = 0xf
)

// GetChpfrq Timerx carrier frequency value
func (r *RegisterChprType) GetChpfrq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChprFieldChpfrqMask) >> RegisterChprFieldChpfrqShift)
}

// SetChpfrq Timerx carrier frequency value
func (r *RegisterChprType) SetChpfrq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChprFieldChpfrqMask)|(uint32(value)<<RegisterChprFieldChpfrqShift))
}

const (
	RegisterChprFieldChpdtyShift = 4
	RegisterChprFieldChpdtyMask  = 0x70
)

// GetChpdty Timerx chopper duty cycle value
func (r *RegisterChprType) GetChpdty() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChprFieldChpdtyMask) >> RegisterChprFieldChpdtyShift)
}

// SetChpdty Timerx chopper duty cycle value
func (r *RegisterChprType) SetChpdty(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChprFieldChpdtyMask)|(uint32(value)<<RegisterChprFieldChpdtyShift))
}

const (
	RegisterChprFieldStrtpwShift = 7
	RegisterChprFieldStrtpwMask  = 0x780
)

// GetStrtpw STRTPW
func (r *RegisterChprType) GetStrtpw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChprFieldStrtpwMask) >> RegisterChprFieldStrtpwShift)
}

// SetStrtpw STRTPW
func (r *RegisterChprType) SetStrtpw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChprFieldStrtpwMask)|(uint32(value)<<RegisterChprFieldStrtpwShift))
}

// RegisterCpt1crType Timerx Capture 2 Control Register
type RegisterCpt1crType uint32

func (r *RegisterCpt1crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCpt1crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCpt1crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCpt1crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCpt1crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCpt1crFieldSwcptShift = 0
	RegisterCpt1crFieldSwcptMask  = 0x1
)

// GetSwcpt Software Capture
func (r *RegisterCpt1crType) GetSwcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldSwcptMask) != 0
}

// SetSwcpt Software Capture
func (r *RegisterCpt1crType) SetSwcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldSwcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldSwcptMask)
	}
}

const (
	RegisterCpt1crFieldUdpcptShift = 1
	RegisterCpt1crFieldUdpcptMask  = 0x2
)

// GetUdpcpt Update Capture
func (r *RegisterCpt1crType) GetUdpcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldUdpcptMask) != 0
}

// SetUdpcpt Update Capture
func (r *RegisterCpt1crType) SetUdpcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldUdpcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldUdpcptMask)
	}
}

const (
	RegisterCpt1crFieldExev1cptShift = 2
	RegisterCpt1crFieldExev1cptMask  = 0x4
)

// GetExev1cpt External Event 1 Capture
func (r *RegisterCpt1crType) GetExev1cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldExev1cptMask) != 0
}

// SetExev1cpt External Event 1 Capture
func (r *RegisterCpt1crType) SetExev1cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldExev1cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldExev1cptMask)
	}
}

const (
	RegisterCpt1crFieldExev2cptShift = 3
	RegisterCpt1crFieldExev2cptMask  = 0x8
)

// GetExev2cpt External Event 2 Capture
func (r *RegisterCpt1crType) GetExev2cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldExev2cptMask) != 0
}

// SetExev2cpt External Event 2 Capture
func (r *RegisterCpt1crType) SetExev2cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldExev2cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldExev2cptMask)
	}
}

const (
	RegisterCpt1crFieldExev3cptShift = 4
	RegisterCpt1crFieldExev3cptMask  = 0x10
)

// GetExev3cpt External Event 3 Capture
func (r *RegisterCpt1crType) GetExev3cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldExev3cptMask) != 0
}

// SetExev3cpt External Event 3 Capture
func (r *RegisterCpt1crType) SetExev3cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldExev3cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldExev3cptMask)
	}
}

const (
	RegisterCpt1crFieldExev4cptShift = 5
	RegisterCpt1crFieldExev4cptMask  = 0x20
)

// GetExev4cpt External Event 4 Capture
func (r *RegisterCpt1crType) GetExev4cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldExev4cptMask) != 0
}

// SetExev4cpt External Event 4 Capture
func (r *RegisterCpt1crType) SetExev4cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldExev4cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldExev4cptMask)
	}
}

const (
	RegisterCpt1crFieldExev5cptShift = 6
	RegisterCpt1crFieldExev5cptMask  = 0x40
)

// GetExev5cpt External Event 5 Capture
func (r *RegisterCpt1crType) GetExev5cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldExev5cptMask) != 0
}

// SetExev5cpt External Event 5 Capture
func (r *RegisterCpt1crType) SetExev5cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldExev5cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldExev5cptMask)
	}
}

const (
	RegisterCpt1crFieldExev6cptShift = 7
	RegisterCpt1crFieldExev6cptMask  = 0x80
)

// GetExev6cpt External Event 6 Capture
func (r *RegisterCpt1crType) GetExev6cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldExev6cptMask) != 0
}

// SetExev6cpt External Event 6 Capture
func (r *RegisterCpt1crType) SetExev6cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldExev6cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldExev6cptMask)
	}
}

const (
	RegisterCpt1crFieldExev7cptShift = 8
	RegisterCpt1crFieldExev7cptMask  = 0x100
)

// GetExev7cpt External Event 7 Capture
func (r *RegisterCpt1crType) GetExev7cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldExev7cptMask) != 0
}

// SetExev7cpt External Event 7 Capture
func (r *RegisterCpt1crType) SetExev7cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldExev7cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldExev7cptMask)
	}
}

const (
	RegisterCpt1crFieldExev8cptShift = 9
	RegisterCpt1crFieldExev8cptMask  = 0x200
)

// GetExev8cpt External Event 8 Capture
func (r *RegisterCpt1crType) GetExev8cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldExev8cptMask) != 0
}

// SetExev8cpt External Event 8 Capture
func (r *RegisterCpt1crType) SetExev8cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldExev8cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldExev8cptMask)
	}
}

const (
	RegisterCpt1crFieldExev9cptShift = 10
	RegisterCpt1crFieldExev9cptMask  = 0x400
)

// GetExev9cpt External Event 9 Capture
func (r *RegisterCpt1crType) GetExev9cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldExev9cptMask) != 0
}

// SetExev9cpt External Event 9 Capture
func (r *RegisterCpt1crType) SetExev9cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldExev9cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldExev9cptMask)
	}
}

const (
	RegisterCpt1crFieldExev10cptShift = 11
	RegisterCpt1crFieldExev10cptMask  = 0x800
)

// GetExev10cpt External Event 10 Capture
func (r *RegisterCpt1crType) GetExev10cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldExev10cptMask) != 0
}

// SetExev10cpt External Event 10 Capture
func (r *RegisterCpt1crType) SetExev10cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldExev10cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldExev10cptMask)
	}
}

const (
	RegisterCpt1crFieldT1setShift = 16
	RegisterCpt1crFieldT1setMask  = 0x10000
)

// GetT1set Timer output 1 Set
func (r *RegisterCpt1crType) GetT1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldT1setMask) != 0
}

// SetT1set Timer output 1 Set
func (r *RegisterCpt1crType) SetT1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldT1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldT1setMask)
	}
}

const (
	RegisterCpt1crFieldT1rstShift = 17
	RegisterCpt1crFieldT1rstMask  = 0x20000
)

// GetT1rst Timer output 1 Reset
func (r *RegisterCpt1crType) GetT1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldT1rstMask) != 0
}

// SetT1rst Timer output 1 Reset
func (r *RegisterCpt1crType) SetT1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldT1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldT1rstMask)
	}
}

const (
	RegisterCpt1crFieldTcmp1Shift = 18
	RegisterCpt1crFieldTcmp1Mask  = 0x40000
)

// GetTcmp1 Timer Compare 1
func (r *RegisterCpt1crType) GetTcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTcmp1Mask) != 0
}

// SetTcmp1 Timer Compare 1
func (r *RegisterCpt1crType) SetTcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTcmp1Mask)
	}
}

const (
	RegisterCpt1crFieldTcmp2Shift = 19
	RegisterCpt1crFieldTcmp2Mask  = 0x80000
)

// GetTcmp2 Timer Compare 2
func (r *RegisterCpt1crType) GetTcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTcmp2Mask) != 0
}

// SetTcmp2 Timer Compare 2
func (r *RegisterCpt1crType) SetTcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTcmp2Mask)
	}
}

const (
	RegisterCpt1crFieldTc1setShift = 20
	RegisterCpt1crFieldTc1setMask  = 0x100000
)

// GetTc1set Timer C output 1 Set
func (r *RegisterCpt1crType) GetTc1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTc1setMask) != 0
}

// SetTc1set Timer C output 1 Set
func (r *RegisterCpt1crType) SetTc1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTc1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTc1setMask)
	}
}

const (
	RegisterCpt1crFieldTc1rstShift = 21
	RegisterCpt1crFieldTc1rstMask  = 0x200000
)

// GetTc1rst Timer C output 1 Reset
func (r *RegisterCpt1crType) GetTc1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTc1rstMask) != 0
}

// SetTc1rst Timer C output 1 Reset
func (r *RegisterCpt1crType) SetTc1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTc1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTc1rstMask)
	}
}

const (
	RegisterCpt1crFieldTccmp1Shift = 22
	RegisterCpt1crFieldTccmp1Mask  = 0x400000
)

// GetTccmp1 Timer C Compare 1
func (r *RegisterCpt1crType) GetTccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTccmp1Mask) != 0
}

// SetTccmp1 Timer C Compare 1
func (r *RegisterCpt1crType) SetTccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTccmp1Mask)
	}
}

const (
	RegisterCpt1crFieldTccmp2Shift = 23
	RegisterCpt1crFieldTccmp2Mask  = 0x800000
)

// GetTccmp2 Timer C Compare 2
func (r *RegisterCpt1crType) GetTccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTccmp2Mask) != 0
}

// SetTccmp2 Timer C Compare 2
func (r *RegisterCpt1crType) SetTccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTccmp2Mask)
	}
}

const (
	RegisterCpt1crFieldTd1setShift = 24
	RegisterCpt1crFieldTd1setMask  = 0x1000000
)

// GetTd1set Timer D output 1 Set
func (r *RegisterCpt1crType) GetTd1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTd1setMask) != 0
}

// SetTd1set Timer D output 1 Set
func (r *RegisterCpt1crType) SetTd1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTd1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTd1setMask)
	}
}

const (
	RegisterCpt1crFieldTd1rstShift = 25
	RegisterCpt1crFieldTd1rstMask  = 0x2000000
)

// GetTd1rst Timer D output 1 Reset
func (r *RegisterCpt1crType) GetTd1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTd1rstMask) != 0
}

// SetTd1rst Timer D output 1 Reset
func (r *RegisterCpt1crType) SetTd1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTd1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTd1rstMask)
	}
}

const (
	RegisterCpt1crFieldTdcmp1Shift = 26
	RegisterCpt1crFieldTdcmp1Mask  = 0x4000000
)

// GetTdcmp1 Timer D Compare 1
func (r *RegisterCpt1crType) GetTdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTdcmp1Mask) != 0
}

// SetTdcmp1 Timer D Compare 1
func (r *RegisterCpt1crType) SetTdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTdcmp1Mask)
	}
}

const (
	RegisterCpt1crFieldTdcmp2Shift = 27
	RegisterCpt1crFieldTdcmp2Mask  = 0x8000000
)

// GetTdcmp2 Timer D Compare 2
func (r *RegisterCpt1crType) GetTdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTdcmp2Mask) != 0
}

// SetTdcmp2 Timer D Compare 2
func (r *RegisterCpt1crType) SetTdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTdcmp2Mask)
	}
}

const (
	RegisterCpt1crFieldTe1setShift = 28
	RegisterCpt1crFieldTe1setMask  = 0x10000000
)

// GetTe1set Timer E output 1 Set
func (r *RegisterCpt1crType) GetTe1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTe1setMask) != 0
}

// SetTe1set Timer E output 1 Set
func (r *RegisterCpt1crType) SetTe1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTe1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTe1setMask)
	}
}

const (
	RegisterCpt1crFieldTe1rstShift = 29
	RegisterCpt1crFieldTe1rstMask  = 0x20000000
)

// GetTe1rst Timer E output 1 Reset
func (r *RegisterCpt1crType) GetTe1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTe1rstMask) != 0
}

// SetTe1rst Timer E output 1 Reset
func (r *RegisterCpt1crType) SetTe1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTe1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTe1rstMask)
	}
}

const (
	RegisterCpt1crFieldTecmp1Shift = 30
	RegisterCpt1crFieldTecmp1Mask  = 0x40000000
)

// GetTecmp1 Timer E Compare 1
func (r *RegisterCpt1crType) GetTecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTecmp1Mask) != 0
}

// SetTecmp1 Timer E Compare 1
func (r *RegisterCpt1crType) SetTecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTecmp1Mask)
	}
}

const (
	RegisterCpt1crFieldTecmp2Shift = 31
	RegisterCpt1crFieldTecmp2Mask  = 0x80000000
)

// GetTecmp2 Timer E Compare 2
func (r *RegisterCpt1crType) GetTecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldTecmp2Mask) != 0
}

// SetTecmp2 Timer E Compare 2
func (r *RegisterCpt1crType) SetTecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1crFieldTecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldTecmp2Mask)
	}
}

// RegisterCpt2crType CPT2xCR
type RegisterCpt2crType uint32

func (r *RegisterCpt2crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCpt2crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCpt2crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCpt2crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCpt2crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCpt2crFieldSwcptShift = 0
	RegisterCpt2crFieldSwcptMask  = 0x1
)

// GetSwcpt Software Capture
func (r *RegisterCpt2crType) GetSwcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldSwcptMask) != 0
}

// SetSwcpt Software Capture
func (r *RegisterCpt2crType) SetSwcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldSwcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldSwcptMask)
	}
}

const (
	RegisterCpt2crFieldUdpcptShift = 1
	RegisterCpt2crFieldUdpcptMask  = 0x2
)

// GetUdpcpt Update Capture
func (r *RegisterCpt2crType) GetUdpcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldUdpcptMask) != 0
}

// SetUdpcpt Update Capture
func (r *RegisterCpt2crType) SetUdpcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldUdpcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldUdpcptMask)
	}
}

const (
	RegisterCpt2crFieldExev1cptShift = 2
	RegisterCpt2crFieldExev1cptMask  = 0x4
)

// GetExev1cpt External Event 1 Capture
func (r *RegisterCpt2crType) GetExev1cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldExev1cptMask) != 0
}

// SetExev1cpt External Event 1 Capture
func (r *RegisterCpt2crType) SetExev1cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldExev1cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldExev1cptMask)
	}
}

const (
	RegisterCpt2crFieldExev2cptShift = 3
	RegisterCpt2crFieldExev2cptMask  = 0x8
)

// GetExev2cpt External Event 2 Capture
func (r *RegisterCpt2crType) GetExev2cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldExev2cptMask) != 0
}

// SetExev2cpt External Event 2 Capture
func (r *RegisterCpt2crType) SetExev2cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldExev2cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldExev2cptMask)
	}
}

const (
	RegisterCpt2crFieldExev3cptShift = 4
	RegisterCpt2crFieldExev3cptMask  = 0x10
)

// GetExev3cpt External Event 3 Capture
func (r *RegisterCpt2crType) GetExev3cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldExev3cptMask) != 0
}

// SetExev3cpt External Event 3 Capture
func (r *RegisterCpt2crType) SetExev3cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldExev3cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldExev3cptMask)
	}
}

const (
	RegisterCpt2crFieldExev4cptShift = 5
	RegisterCpt2crFieldExev4cptMask  = 0x20
)

// GetExev4cpt External Event 4 Capture
func (r *RegisterCpt2crType) GetExev4cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldExev4cptMask) != 0
}

// SetExev4cpt External Event 4 Capture
func (r *RegisterCpt2crType) SetExev4cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldExev4cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldExev4cptMask)
	}
}

const (
	RegisterCpt2crFieldExev5cptShift = 6
	RegisterCpt2crFieldExev5cptMask  = 0x40
)

// GetExev5cpt External Event 5 Capture
func (r *RegisterCpt2crType) GetExev5cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldExev5cptMask) != 0
}

// SetExev5cpt External Event 5 Capture
func (r *RegisterCpt2crType) SetExev5cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldExev5cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldExev5cptMask)
	}
}

const (
	RegisterCpt2crFieldExev6cptShift = 7
	RegisterCpt2crFieldExev6cptMask  = 0x80
)

// GetExev6cpt External Event 6 Capture
func (r *RegisterCpt2crType) GetExev6cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldExev6cptMask) != 0
}

// SetExev6cpt External Event 6 Capture
func (r *RegisterCpt2crType) SetExev6cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldExev6cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldExev6cptMask)
	}
}

const (
	RegisterCpt2crFieldExev7cptShift = 8
	RegisterCpt2crFieldExev7cptMask  = 0x100
)

// GetExev7cpt External Event 7 Capture
func (r *RegisterCpt2crType) GetExev7cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldExev7cptMask) != 0
}

// SetExev7cpt External Event 7 Capture
func (r *RegisterCpt2crType) SetExev7cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldExev7cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldExev7cptMask)
	}
}

const (
	RegisterCpt2crFieldExev8cptShift = 9
	RegisterCpt2crFieldExev8cptMask  = 0x200
)

// GetExev8cpt External Event 8 Capture
func (r *RegisterCpt2crType) GetExev8cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldExev8cptMask) != 0
}

// SetExev8cpt External Event 8 Capture
func (r *RegisterCpt2crType) SetExev8cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldExev8cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldExev8cptMask)
	}
}

const (
	RegisterCpt2crFieldExev9cptShift = 10
	RegisterCpt2crFieldExev9cptMask  = 0x400
)

// GetExev9cpt External Event 9 Capture
func (r *RegisterCpt2crType) GetExev9cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldExev9cptMask) != 0
}

// SetExev9cpt External Event 9 Capture
func (r *RegisterCpt2crType) SetExev9cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldExev9cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldExev9cptMask)
	}
}

const (
	RegisterCpt2crFieldExev10cptShift = 11
	RegisterCpt2crFieldExev10cptMask  = 0x800
)

// GetExev10cpt External Event 10 Capture
func (r *RegisterCpt2crType) GetExev10cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldExev10cptMask) != 0
}

// SetExev10cpt External Event 10 Capture
func (r *RegisterCpt2crType) SetExev10cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldExev10cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldExev10cptMask)
	}
}

const (
	RegisterCpt2crFieldT1setShift = 16
	RegisterCpt2crFieldT1setMask  = 0x10000
)

// GetT1set Timer output 1 Set
func (r *RegisterCpt2crType) GetT1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldT1setMask) != 0
}

// SetT1set Timer output 1 Set
func (r *RegisterCpt2crType) SetT1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldT1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldT1setMask)
	}
}

const (
	RegisterCpt2crFieldT1rstShift = 17
	RegisterCpt2crFieldT1rstMask  = 0x20000
)

// GetT1rst Timer output 1 Reset
func (r *RegisterCpt2crType) GetT1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldT1rstMask) != 0
}

// SetT1rst Timer output 1 Reset
func (r *RegisterCpt2crType) SetT1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldT1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldT1rstMask)
	}
}

const (
	RegisterCpt2crFieldTcmp1Shift = 18
	RegisterCpt2crFieldTcmp1Mask  = 0x40000
)

// GetTcmp1 Timer Compare 1
func (r *RegisterCpt2crType) GetTcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTcmp1Mask) != 0
}

// SetTcmp1 Timer Compare 1
func (r *RegisterCpt2crType) SetTcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTcmp1Mask)
	}
}

const (
	RegisterCpt2crFieldTcmp2Shift = 19
	RegisterCpt2crFieldTcmp2Mask  = 0x80000
)

// GetTcmp2 Timer Compare 2
func (r *RegisterCpt2crType) GetTcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTcmp2Mask) != 0
}

// SetTcmp2 Timer Compare 2
func (r *RegisterCpt2crType) SetTcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTcmp2Mask)
	}
}

const (
	RegisterCpt2crFieldTc1setShift = 20
	RegisterCpt2crFieldTc1setMask  = 0x100000
)

// GetTc1set Timer C output 1 Set
func (r *RegisterCpt2crType) GetTc1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTc1setMask) != 0
}

// SetTc1set Timer C output 1 Set
func (r *RegisterCpt2crType) SetTc1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTc1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTc1setMask)
	}
}

const (
	RegisterCpt2crFieldTc1rstShift = 21
	RegisterCpt2crFieldTc1rstMask  = 0x200000
)

// GetTc1rst Timer C output 1 Reset
func (r *RegisterCpt2crType) GetTc1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTc1rstMask) != 0
}

// SetTc1rst Timer C output 1 Reset
func (r *RegisterCpt2crType) SetTc1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTc1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTc1rstMask)
	}
}

const (
	RegisterCpt2crFieldTccmp1Shift = 22
	RegisterCpt2crFieldTccmp1Mask  = 0x400000
)

// GetTccmp1 Timer C Compare 1
func (r *RegisterCpt2crType) GetTccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTccmp1Mask) != 0
}

// SetTccmp1 Timer C Compare 1
func (r *RegisterCpt2crType) SetTccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTccmp1Mask)
	}
}

const (
	RegisterCpt2crFieldTccmp2Shift = 23
	RegisterCpt2crFieldTccmp2Mask  = 0x800000
)

// GetTccmp2 Timer C Compare 2
func (r *RegisterCpt2crType) GetTccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTccmp2Mask) != 0
}

// SetTccmp2 Timer C Compare 2
func (r *RegisterCpt2crType) SetTccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTccmp2Mask)
	}
}

const (
	RegisterCpt2crFieldTd1setShift = 24
	RegisterCpt2crFieldTd1setMask  = 0x1000000
)

// GetTd1set Timer D output 1 Set
func (r *RegisterCpt2crType) GetTd1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTd1setMask) != 0
}

// SetTd1set Timer D output 1 Set
func (r *RegisterCpt2crType) SetTd1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTd1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTd1setMask)
	}
}

const (
	RegisterCpt2crFieldTd1rstShift = 25
	RegisterCpt2crFieldTd1rstMask  = 0x2000000
)

// GetTd1rst Timer D output 1 Reset
func (r *RegisterCpt2crType) GetTd1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTd1rstMask) != 0
}

// SetTd1rst Timer D output 1 Reset
func (r *RegisterCpt2crType) SetTd1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTd1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTd1rstMask)
	}
}

const (
	RegisterCpt2crFieldTdcmp1Shift = 26
	RegisterCpt2crFieldTdcmp1Mask  = 0x4000000
)

// GetTdcmp1 Timer D Compare 1
func (r *RegisterCpt2crType) GetTdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTdcmp1Mask) != 0
}

// SetTdcmp1 Timer D Compare 1
func (r *RegisterCpt2crType) SetTdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTdcmp1Mask)
	}
}

const (
	RegisterCpt2crFieldTdcmp2Shift = 27
	RegisterCpt2crFieldTdcmp2Mask  = 0x8000000
)

// GetTdcmp2 Timer D Compare 2
func (r *RegisterCpt2crType) GetTdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTdcmp2Mask) != 0
}

// SetTdcmp2 Timer D Compare 2
func (r *RegisterCpt2crType) SetTdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTdcmp2Mask)
	}
}

const (
	RegisterCpt2crFieldTe1setShift = 28
	RegisterCpt2crFieldTe1setMask  = 0x10000000
)

// GetTe1set Timer E output 1 Set
func (r *RegisterCpt2crType) GetTe1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTe1setMask) != 0
}

// SetTe1set Timer E output 1 Set
func (r *RegisterCpt2crType) SetTe1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTe1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTe1setMask)
	}
}

const (
	RegisterCpt2crFieldTe1rstShift = 29
	RegisterCpt2crFieldTe1rstMask  = 0x20000000
)

// GetTe1rst Timer E output 1 Reset
func (r *RegisterCpt2crType) GetTe1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTe1rstMask) != 0
}

// SetTe1rst Timer E output 1 Reset
func (r *RegisterCpt2crType) SetTe1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTe1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTe1rstMask)
	}
}

const (
	RegisterCpt2crFieldTecmp1Shift = 30
	RegisterCpt2crFieldTecmp1Mask  = 0x40000000
)

// GetTecmp1 Timer E Compare 1
func (r *RegisterCpt2crType) GetTecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTecmp1Mask) != 0
}

// SetTecmp1 Timer E Compare 1
func (r *RegisterCpt2crType) SetTecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTecmp1Mask)
	}
}

const (
	RegisterCpt2crFieldTecmp2Shift = 31
	RegisterCpt2crFieldTecmp2Mask  = 0x80000000
)

// GetTecmp2 Timer E Compare 2
func (r *RegisterCpt2crType) GetTecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldTecmp2Mask) != 0
}

// SetTecmp2 Timer E Compare 2
func (r *RegisterCpt2crType) SetTecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2crFieldTecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldTecmp2Mask)
	}
}

// RegisterOutrType Timerx Output Register
type RegisterOutrType uint32

func (r *RegisterOutrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOutrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOutrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOutrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOutrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOutrFieldPol1Shift = 1
	RegisterOutrFieldPol1Mask  = 0x2
)

// GetPol1 Output 1 polarity
func (r *RegisterOutrType) GetPol1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldPol1Mask) != 0
}

// SetPol1 Output 1 polarity
func (r *RegisterOutrType) SetPol1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutrFieldPol1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldPol1Mask)
	}
}

const (
	RegisterOutrFieldIdlem1Shift = 2
	RegisterOutrFieldIdlem1Mask  = 0x4
)

// GetIdlem1 Output 1 Idle mode
func (r *RegisterOutrType) GetIdlem1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldIdlem1Mask) != 0
}

// SetIdlem1 Output 1 Idle mode
func (r *RegisterOutrType) SetIdlem1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutrFieldIdlem1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldIdlem1Mask)
	}
}

const (
	RegisterOutrFieldIdles1Shift = 3
	RegisterOutrFieldIdles1Mask  = 0x8
)

// GetIdles1 Output 1 Idle State
func (r *RegisterOutrType) GetIdles1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldIdles1Mask) != 0
}

// SetIdles1 Output 1 Idle State
func (r *RegisterOutrType) SetIdles1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutrFieldIdles1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldIdles1Mask)
	}
}

const (
	RegisterOutrFieldFault1Shift = 4
	RegisterOutrFieldFault1Mask  = 0x30
)

// GetFault1 Output 1 Fault state
func (r *RegisterOutrType) GetFault1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldFault1Mask) >> RegisterOutrFieldFault1Shift)
}

// SetFault1 Output 1 Fault state
func (r *RegisterOutrType) SetFault1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldFault1Mask)|(uint32(value)<<RegisterOutrFieldFault1Shift))
}

const (
	RegisterOutrFieldChp1Shift = 6
	RegisterOutrFieldChp1Mask  = 0x40
)

// GetChp1 Output 1 Chopper enable
func (r *RegisterOutrType) GetChp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldChp1Mask) != 0
}

// SetChp1 Output 1 Chopper enable
func (r *RegisterOutrType) SetChp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutrFieldChp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldChp1Mask)
	}
}

const (
	RegisterOutrFieldDidl1Shift = 7
	RegisterOutrFieldDidl1Mask  = 0x80
)

// GetDidl1 Output 1 Deadtime upon burst mode Idle entry
func (r *RegisterOutrType) GetDidl1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldDidl1Mask) != 0
}

// SetDidl1 Output 1 Deadtime upon burst mode Idle entry
func (r *RegisterOutrType) SetDidl1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutrFieldDidl1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldDidl1Mask)
	}
}

const (
	RegisterOutrFieldDtenShift = 8
	RegisterOutrFieldDtenMask  = 0x100
)

// GetDten Deadtime enable
func (r *RegisterOutrType) GetDten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldDtenMask) != 0
}

// SetDten Deadtime enable
func (r *RegisterOutrType) SetDten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutrFieldDtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldDtenMask)
	}
}

const (
	RegisterOutrFieldDlyprtenShift = 9
	RegisterOutrFieldDlyprtenMask  = 0x200
)

// GetDlyprten Delayed Protection Enable
func (r *RegisterOutrType) GetDlyprten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldDlyprtenMask) != 0
}

// SetDlyprten Delayed Protection Enable
func (r *RegisterOutrType) SetDlyprten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutrFieldDlyprtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldDlyprtenMask)
	}
}

const (
	RegisterOutrFieldDlyprtShift = 10
	RegisterOutrFieldDlyprtMask  = 0x1c00
)

// GetDlyprt Delayed Protection
func (r *RegisterOutrType) GetDlyprt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldDlyprtMask) >> RegisterOutrFieldDlyprtShift)
}

// SetDlyprt Delayed Protection
func (r *RegisterOutrType) SetDlyprt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldDlyprtMask)|(uint32(value)<<RegisterOutrFieldDlyprtShift))
}

const (
	RegisterOutrFieldPol2Shift = 17
	RegisterOutrFieldPol2Mask  = 0x20000
)

// GetPol2 Output 2 polarity
func (r *RegisterOutrType) GetPol2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldPol2Mask) != 0
}

// SetPol2 Output 2 polarity
func (r *RegisterOutrType) SetPol2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutrFieldPol2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldPol2Mask)
	}
}

const (
	RegisterOutrFieldIdlem2Shift = 18
	RegisterOutrFieldIdlem2Mask  = 0x40000
)

// GetIdlem2 Output 2 Idle mode
func (r *RegisterOutrType) GetIdlem2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldIdlem2Mask) != 0
}

// SetIdlem2 Output 2 Idle mode
func (r *RegisterOutrType) SetIdlem2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutrFieldIdlem2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldIdlem2Mask)
	}
}

const (
	RegisterOutrFieldIdles2Shift = 19
	RegisterOutrFieldIdles2Mask  = 0x80000
)

// GetIdles2 Output 2 Idle State
func (r *RegisterOutrType) GetIdles2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldIdles2Mask) != 0
}

// SetIdles2 Output 2 Idle State
func (r *RegisterOutrType) SetIdles2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutrFieldIdles2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldIdles2Mask)
	}
}

const (
	RegisterOutrFieldFault2Shift = 20
	RegisterOutrFieldFault2Mask  = 0x300000
)

// GetFault2 Output 2 Fault state
func (r *RegisterOutrType) GetFault2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldFault2Mask) >> RegisterOutrFieldFault2Shift)
}

// SetFault2 Output 2 Fault state
func (r *RegisterOutrType) SetFault2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldFault2Mask)|(uint32(value)<<RegisterOutrFieldFault2Shift))
}

const (
	RegisterOutrFieldChp2Shift = 22
	RegisterOutrFieldChp2Mask  = 0x400000
)

// GetChp2 Output 2 Chopper enable
func (r *RegisterOutrType) GetChp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldChp2Mask) != 0
}

// SetChp2 Output 2 Chopper enable
func (r *RegisterOutrType) SetChp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutrFieldChp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldChp2Mask)
	}
}

const (
	RegisterOutrFieldDidl2Shift = 23
	RegisterOutrFieldDidl2Mask  = 0x800000
)

// GetDidl2 Output 2 Deadtime upon burst mode Idle entry
func (r *RegisterOutrType) GetDidl2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutrFieldDidl2Mask) != 0
}

// SetDidl2 Output 2 Deadtime upon burst mode Idle entry
func (r *RegisterOutrType) SetDidl2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutrFieldDidl2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutrFieldDidl2Mask)
	}
}

// RegisterFltrType Timerx Fault Register
type RegisterFltrType uint32

func (r *RegisterFltrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterFltrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterFltrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterFltrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterFltrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterFltrFieldFlt1enShift = 0
	RegisterFltrFieldFlt1enMask  = 0x1
)

// GetFlt1en Fault 1 enable
func (r *RegisterFltrType) GetFlt1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltrFieldFlt1enMask) != 0
}

// SetFlt1en Fault 1 enable
func (r *RegisterFltrType) SetFlt1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltrFieldFlt1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltrFieldFlt1enMask)
	}
}

const (
	RegisterFltrFieldFlt2enShift = 1
	RegisterFltrFieldFlt2enMask  = 0x2
)

// GetFlt2en Fault 2 enable
func (r *RegisterFltrType) GetFlt2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltrFieldFlt2enMask) != 0
}

// SetFlt2en Fault 2 enable
func (r *RegisterFltrType) SetFlt2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltrFieldFlt2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltrFieldFlt2enMask)
	}
}

const (
	RegisterFltrFieldFlt3enShift = 2
	RegisterFltrFieldFlt3enMask  = 0x4
)

// GetFlt3en Fault 3 enable
func (r *RegisterFltrType) GetFlt3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltrFieldFlt3enMask) != 0
}

// SetFlt3en Fault 3 enable
func (r *RegisterFltrType) SetFlt3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltrFieldFlt3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltrFieldFlt3enMask)
	}
}

const (
	RegisterFltrFieldFlt4enShift = 3
	RegisterFltrFieldFlt4enMask  = 0x8
)

// GetFlt4en Fault 4 enable
func (r *RegisterFltrType) GetFlt4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltrFieldFlt4enMask) != 0
}

// SetFlt4en Fault 4 enable
func (r *RegisterFltrType) SetFlt4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltrFieldFlt4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltrFieldFlt4enMask)
	}
}

const (
	RegisterFltrFieldFlt5enShift = 4
	RegisterFltrFieldFlt5enMask  = 0x10
)

// GetFlt5en Fault 5 enable
func (r *RegisterFltrType) GetFlt5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltrFieldFlt5enMask) != 0
}

// SetFlt5en Fault 5 enable
func (r *RegisterFltrType) SetFlt5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltrFieldFlt5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltrFieldFlt5enMask)
	}
}

const (
	RegisterFltrFieldFltlckShift = 31
	RegisterFltrFieldFltlckMask  = 0x80000000
)

// GetFltlck Fault sources Lock
func (r *RegisterFltrType) GetFltlck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltrFieldFltlckMask) != 0
}

// SetFltlck Fault sources Lock
func (r *RegisterFltrType) SetFltlck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltrFieldFltlckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltrFieldFltlckMask)
	}
}
