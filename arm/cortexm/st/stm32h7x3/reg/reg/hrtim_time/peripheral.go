package hrtim_time

import (
	"unsafe"
	"volatile"
)

var (
	Hrtim_time = (*_hrtim_time)(unsafe.Pointer(uintptr(0x40017680)))
)

type _hrtim_time struct {
	Timecr    registerTimecrType
	Timeisr   registerTimeisrType
	Timeicr   registerTimeicrType
	Timedier5 registerTimedier5Type
	Cnter     registerCnterType
	Perer     registerPererType
	Reper     registerReperType
	Cmp1er    registerCmp1erType
	Cmp1cer   registerCmp1cerType
	Cmp2er    registerCmp2erType
	Cmp3er    registerCmp3erType
	Cmp4er    registerCmp4erType
	Cpt1er    registerCpt1erType
	Cpt2er    registerCpt2erType
	Dter      registerDterType
	Sete1r    registerSete1rType
	Rste1r    registerRste1rType
	Sete2r    registerSete2rType
	Rste2r    registerRste2rType
	Eefer1    registerEefer1Type
	Eefer2    registerEefer2Type
	Rster     registerRsterType
	Chper     registerChperType
	Cpt1ecr   registerCpt1ecrType
	Cpt2ecr   registerCpt2ecrType
	Outer     registerOuterType
	Flter     registerFlterType
}

// registerTimecrType Timerx Control Register
type registerTimecrType uint32

const (
	RegisterTimecrFieldUpdgatShift = 28
	RegisterTimecrFieldUpdgatMask  = 0xf0000000
)

// GetUpdgat Update Gating
func (r *registerTimecrType) GetUpdgat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldUpdgatMask) >> RegisterTimecrFieldUpdgatShift)
}

// SetUpdgat Update Gating
func (r *registerTimecrType) SetUpdgat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldUpdgatMask)|(uint32(value)<<RegisterTimecrFieldUpdgatShift))
}

const (
	RegisterTimecrFieldPreenShift = 27
	RegisterTimecrFieldPreenMask  = 0x8000000
)

// GetPreen Preload enable
func (r *registerTimecrType) GetPreen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldPreenMask) != 0
}

// SetPreen Preload enable
func (r *registerTimecrType) SetPreen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldPreenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldPreenMask)
	}
}

const (
	RegisterTimecrFieldDacsyncShift = 25
	RegisterTimecrFieldDacsyncMask  = 0x6000000
)

// GetDacsync AC Synchronization
func (r *registerTimecrType) GetDacsync() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldDacsyncMask) >> RegisterTimecrFieldDacsyncShift)
}

// SetDacsync AC Synchronization
func (r *registerTimecrType) SetDacsync(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldDacsyncMask)|(uint32(value)<<RegisterTimecrFieldDacsyncShift))
}

const (
	RegisterTimecrFieldMstuShift = 24
	RegisterTimecrFieldMstuMask  = 0x1000000
)

// GetMstu Master Timer update
func (r *registerTimecrType) GetMstu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldMstuMask) != 0
}

// SetMstu Master Timer update
func (r *registerTimecrType) SetMstu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldMstuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldMstuMask)
	}
}

const (
	RegisterTimecrFieldTeuShift = 23
	RegisterTimecrFieldTeuMask  = 0x800000
)

// GetTeu TEU
func (r *registerTimecrType) GetTeu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldTeuMask) != 0
}

// SetTeu TEU
func (r *registerTimecrType) SetTeu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldTeuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldTeuMask)
	}
}

const (
	RegisterTimecrFieldTduShift = 22
	RegisterTimecrFieldTduMask  = 0x400000
)

// GetTdu TDU
func (r *registerTimecrType) GetTdu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldTduMask) != 0
}

// SetTdu TDU
func (r *registerTimecrType) SetTdu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldTduMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldTduMask)
	}
}

const (
	RegisterTimecrFieldTcuShift = 21
	RegisterTimecrFieldTcuMask  = 0x200000
)

// GetTcu TCU
func (r *registerTimecrType) GetTcu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldTcuMask) != 0
}

// SetTcu TCU
func (r *registerTimecrType) SetTcu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldTcuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldTcuMask)
	}
}

const (
	RegisterTimecrFieldTbuShift = 20
	RegisterTimecrFieldTbuMask  = 0x100000
)

// GetTbu TBU
func (r *registerTimecrType) GetTbu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldTbuMask) != 0
}

// SetTbu TBU
func (r *registerTimecrType) SetTbu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldTbuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldTbuMask)
	}
}

const (
	RegisterTimecrFieldTxrstuShift = 18
	RegisterTimecrFieldTxrstuMask  = 0x40000
)

// GetTxrstu Timerx reset update
func (r *registerTimecrType) GetTxrstu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldTxrstuMask) != 0
}

// SetTxrstu Timerx reset update
func (r *registerTimecrType) SetTxrstu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldTxrstuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldTxrstuMask)
	}
}

const (
	RegisterTimecrFieldTxrepuShift = 17
	RegisterTimecrFieldTxrepuMask  = 0x20000
)

// GetTxrepu Timer x Repetition update
func (r *registerTimecrType) GetTxrepu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldTxrepuMask) != 0
}

// SetTxrepu Timer x Repetition update
func (r *registerTimecrType) SetTxrepu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldTxrepuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldTxrepuMask)
	}
}

const (
	RegisterTimecrFieldDelcmp4Shift = 14
	RegisterTimecrFieldDelcmp4Mask  = 0xc000
)

// GetDelcmp4 Delayed CMP4 mode
func (r *registerTimecrType) GetDelcmp4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldDelcmp4Mask) >> RegisterTimecrFieldDelcmp4Shift)
}

// SetDelcmp4 Delayed CMP4 mode
func (r *registerTimecrType) SetDelcmp4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldDelcmp4Mask)|(uint32(value)<<RegisterTimecrFieldDelcmp4Shift))
}

const (
	RegisterTimecrFieldDelcmp2Shift = 12
	RegisterTimecrFieldDelcmp2Mask  = 0x3000
)

// GetDelcmp2 Delayed CMP2 mode
func (r *registerTimecrType) GetDelcmp2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldDelcmp2Mask) >> RegisterTimecrFieldDelcmp2Shift)
}

// SetDelcmp2 Delayed CMP2 mode
func (r *registerTimecrType) SetDelcmp2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldDelcmp2Mask)|(uint32(value)<<RegisterTimecrFieldDelcmp2Shift))
}

const (
	RegisterTimecrFieldSyncstrtxShift = 11
	RegisterTimecrFieldSyncstrtxMask  = 0x800
)

// GetSyncstrtx Synchronization Starts Timer x
func (r *registerTimecrType) GetSyncstrtx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldSyncstrtxMask) != 0
}

// SetSyncstrtx Synchronization Starts Timer x
func (r *registerTimecrType) SetSyncstrtx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldSyncstrtxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldSyncstrtxMask)
	}
}

const (
	RegisterTimecrFieldSyncrstxShift = 10
	RegisterTimecrFieldSyncrstxMask  = 0x400
)

// GetSyncrstx Synchronization Resets Timer x
func (r *registerTimecrType) GetSyncrstx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldSyncrstxMask) != 0
}

// SetSyncrstx Synchronization Resets Timer x
func (r *registerTimecrType) SetSyncrstx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldSyncrstxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldSyncrstxMask)
	}
}

const (
	RegisterTimecrFieldPshpllShift = 6
	RegisterTimecrFieldPshpllMask  = 0x40
)

// GetPshpll Push-Pull mode enable
func (r *registerTimecrType) GetPshpll() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldPshpllMask) != 0
}

// SetPshpll Push-Pull mode enable
func (r *registerTimecrType) SetPshpll(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldPshpllMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldPshpllMask)
	}
}

const (
	RegisterTimecrFieldHalfShift = 5
	RegisterTimecrFieldHalfMask  = 0x20
)

// GetHalf Half mode enable
func (r *registerTimecrType) GetHalf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldHalfMask) != 0
}

// SetHalf Half mode enable
func (r *registerTimecrType) SetHalf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldHalfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldHalfMask)
	}
}

const (
	RegisterTimecrFieldRetrigShift = 4
	RegisterTimecrFieldRetrigMask  = 0x10
)

// GetRetrig Re-triggerable mode
func (r *registerTimecrType) GetRetrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldRetrigMask) != 0
}

// SetRetrig Re-triggerable mode
func (r *registerTimecrType) SetRetrig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldRetrigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldRetrigMask)
	}
}

const (
	RegisterTimecrFieldContShift = 3
	RegisterTimecrFieldContMask  = 0x8
)

// GetCont Continuous mode
func (r *registerTimecrType) GetCont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldContMask) != 0
}

// SetCont Continuous mode
func (r *registerTimecrType) SetCont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimecrFieldContMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldContMask)
	}
}

const (
	RegisterTimecrFieldCk_pscxShift = 0
	RegisterTimecrFieldCk_pscxMask  = 0x7
)

// GetCk_pscx HRTIM Timer x Clock prescaler
func (r *registerTimecrType) GetCk_pscx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimecrFieldCk_pscxMask) >> RegisterTimecrFieldCk_pscxShift)
}

// SetCk_pscx HRTIM Timer x Clock prescaler
func (r *registerTimecrType) SetCk_pscx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimecrFieldCk_pscxMask)|(uint32(value)<<RegisterTimecrFieldCk_pscxShift))
}

// registerTimeisrType Timerx Interrupt Status Register
type registerTimeisrType uint32

const (
	RegisterTimeisrFieldO2statShift = 19
	RegisterTimeisrFieldO2statMask  = 0x80000
)

// GetO2stat Output 2 State
func (r *registerTimeisrType) GetO2stat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldO2statMask) != 0
}

// SetO2stat Output 2 State
func (r *registerTimeisrType) SetO2stat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldO2statMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldO2statMask)
	}
}

const (
	RegisterTimeisrFieldO1statShift = 18
	RegisterTimeisrFieldO1statMask  = 0x40000
)

// GetO1stat Output 1 State
func (r *registerTimeisrType) GetO1stat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldO1statMask) != 0
}

// SetO1stat Output 1 State
func (r *registerTimeisrType) SetO1stat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldO1statMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldO1statMask)
	}
}

const (
	RegisterTimeisrFieldIppstatShift = 17
	RegisterTimeisrFieldIppstatMask  = 0x20000
)

// GetIppstat Idle Push Pull Status
func (r *registerTimeisrType) GetIppstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldIppstatMask) != 0
}

// SetIppstat Idle Push Pull Status
func (r *registerTimeisrType) SetIppstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldIppstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldIppstatMask)
	}
}

const (
	RegisterTimeisrFieldCppstatShift = 16
	RegisterTimeisrFieldCppstatMask  = 0x10000
)

// GetCppstat Current Push Pull Status
func (r *registerTimeisrType) GetCppstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldCppstatMask) != 0
}

// SetCppstat Current Push Pull Status
func (r *registerTimeisrType) SetCppstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldCppstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldCppstatMask)
	}
}

const (
	RegisterTimeisrFieldDlyprtShift = 14
	RegisterTimeisrFieldDlyprtMask  = 0x4000
)

// GetDlyprt Delayed Protection Flag
func (r *registerTimeisrType) GetDlyprt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldDlyprtMask) != 0
}

// SetDlyprt Delayed Protection Flag
func (r *registerTimeisrType) SetDlyprt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldDlyprtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldDlyprtMask)
	}
}

const (
	RegisterTimeisrFieldRstShift = 13
	RegisterTimeisrFieldRstMask  = 0x2000
)

// GetRst Reset Interrupt Flag
func (r *registerTimeisrType) GetRst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldRstMask) != 0
}

// SetRst Reset Interrupt Flag
func (r *registerTimeisrType) SetRst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldRstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldRstMask)
	}
}

const (
	RegisterTimeisrFieldRstx2Shift = 12
	RegisterTimeisrFieldRstx2Mask  = 0x1000
)

// GetRstx2 Output 2 Reset Interrupt Flag
func (r *registerTimeisrType) GetRstx2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldRstx2Mask) != 0
}

// SetRstx2 Output 2 Reset Interrupt Flag
func (r *registerTimeisrType) SetRstx2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldRstx2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldRstx2Mask)
	}
}

const (
	RegisterTimeisrFieldSetx2Shift = 11
	RegisterTimeisrFieldSetx2Mask  = 0x800
)

// GetSetx2 Output 2 Set Interrupt Flag
func (r *registerTimeisrType) GetSetx2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldSetx2Mask) != 0
}

// SetSetx2 Output 2 Set Interrupt Flag
func (r *registerTimeisrType) SetSetx2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldSetx2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldSetx2Mask)
	}
}

const (
	RegisterTimeisrFieldRstx1Shift = 10
	RegisterTimeisrFieldRstx1Mask  = 0x400
)

// GetRstx1 Output 1 Reset Interrupt Flag
func (r *registerTimeisrType) GetRstx1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldRstx1Mask) != 0
}

// SetRstx1 Output 1 Reset Interrupt Flag
func (r *registerTimeisrType) SetRstx1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldRstx1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldRstx1Mask)
	}
}

const (
	RegisterTimeisrFieldSetx1Shift = 9
	RegisterTimeisrFieldSetx1Mask  = 0x200
)

// GetSetx1 Output 1 Set Interrupt Flag
func (r *registerTimeisrType) GetSetx1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldSetx1Mask) != 0
}

// SetSetx1 Output 1 Set Interrupt Flag
func (r *registerTimeisrType) SetSetx1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldSetx1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldSetx1Mask)
	}
}

const (
	RegisterTimeisrFieldCpt2Shift = 8
	RegisterTimeisrFieldCpt2Mask  = 0x100
)

// GetCpt2 Capture2 Interrupt Flag
func (r *registerTimeisrType) GetCpt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldCpt2Mask) != 0
}

// SetCpt2 Capture2 Interrupt Flag
func (r *registerTimeisrType) SetCpt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldCpt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldCpt2Mask)
	}
}

const (
	RegisterTimeisrFieldCpt1Shift = 7
	RegisterTimeisrFieldCpt1Mask  = 0x80
)

// GetCpt1 Capture1 Interrupt Flag
func (r *registerTimeisrType) GetCpt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldCpt1Mask) != 0
}

// SetCpt1 Capture1 Interrupt Flag
func (r *registerTimeisrType) SetCpt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldCpt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldCpt1Mask)
	}
}

const (
	RegisterTimeisrFieldUpdShift = 6
	RegisterTimeisrFieldUpdMask  = 0x40
)

// GetUpd Update Interrupt Flag
func (r *registerTimeisrType) GetUpd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldUpdMask) != 0
}

// SetUpd Update Interrupt Flag
func (r *registerTimeisrType) SetUpd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldUpdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldUpdMask)
	}
}

const (
	RegisterTimeisrFieldRepShift = 4
	RegisterTimeisrFieldRepMask  = 0x10
)

// GetRep Repetition Interrupt Flag
func (r *registerTimeisrType) GetRep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldRepMask) != 0
}

// SetRep Repetition Interrupt Flag
func (r *registerTimeisrType) SetRep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldRepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldRepMask)
	}
}

const (
	RegisterTimeisrFieldCmp4Shift = 3
	RegisterTimeisrFieldCmp4Mask  = 0x8
)

// GetCmp4 Compare 4 Interrupt Flag
func (r *registerTimeisrType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldCmp4Mask) != 0
}

// SetCmp4 Compare 4 Interrupt Flag
func (r *registerTimeisrType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldCmp4Mask)
	}
}

const (
	RegisterTimeisrFieldCmp3Shift = 2
	RegisterTimeisrFieldCmp3Mask  = 0x4
)

// GetCmp3 Compare 3 Interrupt Flag
func (r *registerTimeisrType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldCmp3Mask) != 0
}

// SetCmp3 Compare 3 Interrupt Flag
func (r *registerTimeisrType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldCmp3Mask)
	}
}

const (
	RegisterTimeisrFieldCmp2Shift = 1
	RegisterTimeisrFieldCmp2Mask  = 0x2
)

// GetCmp2 Compare 2 Interrupt Flag
func (r *registerTimeisrType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldCmp2Mask) != 0
}

// SetCmp2 Compare 2 Interrupt Flag
func (r *registerTimeisrType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldCmp2Mask)
	}
}

const (
	RegisterTimeisrFieldCmp1Shift = 0
	RegisterTimeisrFieldCmp1Mask  = 0x1
)

// GetCmp1 Compare 1 Interrupt Flag
func (r *registerTimeisrType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeisrFieldCmp1Mask) != 0
}

// SetCmp1 Compare 1 Interrupt Flag
func (r *registerTimeisrType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeisrFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeisrFieldCmp1Mask)
	}
}

// registerTimeicrType Timerx Interrupt Clear Register
type registerTimeicrType uint32

const (
	RegisterTimeicrFieldDlyprtcShift = 14
	RegisterTimeicrFieldDlyprtcMask  = 0x4000
)

// GetDlyprtc Delayed Protection Flag Clear
func (r *registerTimeicrType) GetDlyprtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldDlyprtcMask) != 0
}

// SetDlyprtc Delayed Protection Flag Clear
func (r *registerTimeicrType) SetDlyprtc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldDlyprtcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldDlyprtcMask)
	}
}

const (
	RegisterTimeicrFieldRstcShift = 13
	RegisterTimeicrFieldRstcMask  = 0x2000
)

// GetRstc Reset Interrupt flag Clear
func (r *registerTimeicrType) GetRstc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldRstcMask) != 0
}

// SetRstc Reset Interrupt flag Clear
func (r *registerTimeicrType) SetRstc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldRstcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldRstcMask)
	}
}

const (
	RegisterTimeicrFieldRstx2cShift = 12
	RegisterTimeicrFieldRstx2cMask  = 0x1000
)

// GetRstx2c Output 2 Reset flag Clear
func (r *registerTimeicrType) GetRstx2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldRstx2cMask) != 0
}

// SetRstx2c Output 2 Reset flag Clear
func (r *registerTimeicrType) SetRstx2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldRstx2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldRstx2cMask)
	}
}

const (
	RegisterTimeicrFieldSet2xcShift = 11
	RegisterTimeicrFieldSet2xcMask  = 0x800
)

// GetSet2xc Output 2 Set flag Clear
func (r *registerTimeicrType) GetSet2xc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldSet2xcMask) != 0
}

// SetSet2xc Output 2 Set flag Clear
func (r *registerTimeicrType) SetSet2xc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldSet2xcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldSet2xcMask)
	}
}

const (
	RegisterTimeicrFieldRstx1cShift = 10
	RegisterTimeicrFieldRstx1cMask  = 0x400
)

// GetRstx1c Output 1 Reset flag Clear
func (r *registerTimeicrType) GetRstx1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldRstx1cMask) != 0
}

// SetRstx1c Output 1 Reset flag Clear
func (r *registerTimeicrType) SetRstx1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldRstx1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldRstx1cMask)
	}
}

const (
	RegisterTimeicrFieldSet1xcShift = 9
	RegisterTimeicrFieldSet1xcMask  = 0x200
)

// GetSet1xc Output 1 Set flag Clear
func (r *registerTimeicrType) GetSet1xc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldSet1xcMask) != 0
}

// SetSet1xc Output 1 Set flag Clear
func (r *registerTimeicrType) SetSet1xc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldSet1xcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldSet1xcMask)
	}
}

const (
	RegisterTimeicrFieldCpt2cShift = 8
	RegisterTimeicrFieldCpt2cMask  = 0x100
)

// GetCpt2c Capture2 Interrupt flag Clear
func (r *registerTimeicrType) GetCpt2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldCpt2cMask) != 0
}

// SetCpt2c Capture2 Interrupt flag Clear
func (r *registerTimeicrType) SetCpt2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldCpt2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldCpt2cMask)
	}
}

const (
	RegisterTimeicrFieldCpt1cShift = 7
	RegisterTimeicrFieldCpt1cMask  = 0x80
)

// GetCpt1c Capture1 Interrupt flag Clear
func (r *registerTimeicrType) GetCpt1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldCpt1cMask) != 0
}

// SetCpt1c Capture1 Interrupt flag Clear
func (r *registerTimeicrType) SetCpt1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldCpt1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldCpt1cMask)
	}
}

const (
	RegisterTimeicrFieldUpdcShift = 6
	RegisterTimeicrFieldUpdcMask  = 0x40
)

// GetUpdc Update Interrupt flag Clear
func (r *registerTimeicrType) GetUpdc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldUpdcMask) != 0
}

// SetUpdc Update Interrupt flag Clear
func (r *registerTimeicrType) SetUpdc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldUpdcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldUpdcMask)
	}
}

const (
	RegisterTimeicrFieldRepcShift = 4
	RegisterTimeicrFieldRepcMask  = 0x10
)

// GetRepc Repetition Interrupt flag Clear
func (r *registerTimeicrType) GetRepc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldRepcMask) != 0
}

// SetRepc Repetition Interrupt flag Clear
func (r *registerTimeicrType) SetRepc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldRepcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldRepcMask)
	}
}

const (
	RegisterTimeicrFieldCmp4cShift = 3
	RegisterTimeicrFieldCmp4cMask  = 0x8
)

// GetCmp4c Compare 4 Interrupt flag Clear
func (r *registerTimeicrType) GetCmp4c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldCmp4cMask) != 0
}

// SetCmp4c Compare 4 Interrupt flag Clear
func (r *registerTimeicrType) SetCmp4c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldCmp4cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldCmp4cMask)
	}
}

const (
	RegisterTimeicrFieldCmp3cShift = 2
	RegisterTimeicrFieldCmp3cMask  = 0x4
)

// GetCmp3c Compare 3 Interrupt flag Clear
func (r *registerTimeicrType) GetCmp3c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldCmp3cMask) != 0
}

// SetCmp3c Compare 3 Interrupt flag Clear
func (r *registerTimeicrType) SetCmp3c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldCmp3cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldCmp3cMask)
	}
}

const (
	RegisterTimeicrFieldCmp2cShift = 1
	RegisterTimeicrFieldCmp2cMask  = 0x2
)

// GetCmp2c Compare 2 Interrupt flag Clear
func (r *registerTimeicrType) GetCmp2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldCmp2cMask) != 0
}

// SetCmp2c Compare 2 Interrupt flag Clear
func (r *registerTimeicrType) SetCmp2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldCmp2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldCmp2cMask)
	}
}

const (
	RegisterTimeicrFieldCmp1cShift = 0
	RegisterTimeicrFieldCmp1cMask  = 0x1
)

// GetCmp1c Compare 1 Interrupt flag Clear
func (r *registerTimeicrType) GetCmp1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeicrFieldCmp1cMask) != 0
}

// SetCmp1c Compare 1 Interrupt flag Clear
func (r *registerTimeicrType) SetCmp1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeicrFieldCmp1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeicrFieldCmp1cMask)
	}
}

// registerTimedier5Type TIMxDIER5
type registerTimedier5Type uint32

const (
	RegisterTimedier5FieldDlyprtdeShift = 30
	RegisterTimedier5FieldDlyprtdeMask  = 0x40000000
)

// GetDlyprtde DLYPRTDE
func (r *registerTimedier5Type) GetDlyprtde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldDlyprtdeMask) != 0
}

// SetDlyprtde DLYPRTDE
func (r *registerTimedier5Type) SetDlyprtde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldDlyprtdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldDlyprtdeMask)
	}
}

const (
	RegisterTimedier5FieldRstdeShift = 29
	RegisterTimedier5FieldRstdeMask  = 0x20000000
)

// GetRstde RSTDE
func (r *registerTimedier5Type) GetRstde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldRstdeMask) != 0
}

// SetRstde RSTDE
func (r *registerTimedier5Type) SetRstde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldRstdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldRstdeMask)
	}
}

const (
	RegisterTimedier5FieldRstx2deShift = 28
	RegisterTimedier5FieldRstx2deMask  = 0x10000000
)

// GetRstx2de RSTx2DE
func (r *registerTimedier5Type) GetRstx2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldRstx2deMask) != 0
}

// SetRstx2de RSTx2DE
func (r *registerTimedier5Type) SetRstx2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldRstx2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldRstx2deMask)
	}
}

const (
	RegisterTimedier5FieldSetx2deShift = 27
	RegisterTimedier5FieldSetx2deMask  = 0x8000000
)

// GetSetx2de SETx2DE
func (r *registerTimedier5Type) GetSetx2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldSetx2deMask) != 0
}

// SetSetx2de SETx2DE
func (r *registerTimedier5Type) SetSetx2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldSetx2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldSetx2deMask)
	}
}

const (
	RegisterTimedier5FieldRstx1deShift = 26
	RegisterTimedier5FieldRstx1deMask  = 0x4000000
)

// GetRstx1de RSTx1DE
func (r *registerTimedier5Type) GetRstx1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldRstx1deMask) != 0
}

// SetRstx1de RSTx1DE
func (r *registerTimedier5Type) SetRstx1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldRstx1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldRstx1deMask)
	}
}

const (
	RegisterTimedier5FieldSet1xdeShift = 25
	RegisterTimedier5FieldSet1xdeMask  = 0x2000000
)

// GetSet1xde SET1xDE
func (r *registerTimedier5Type) GetSet1xde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldSet1xdeMask) != 0
}

// SetSet1xde SET1xDE
func (r *registerTimedier5Type) SetSet1xde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldSet1xdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldSet1xdeMask)
	}
}

const (
	RegisterTimedier5FieldCpt2deShift = 24
	RegisterTimedier5FieldCpt2deMask  = 0x1000000
)

// GetCpt2de CPT2DE
func (r *registerTimedier5Type) GetCpt2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldCpt2deMask) != 0
}

// SetCpt2de CPT2DE
func (r *registerTimedier5Type) SetCpt2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldCpt2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldCpt2deMask)
	}
}

const (
	RegisterTimedier5FieldCpt1deShift = 23
	RegisterTimedier5FieldCpt1deMask  = 0x800000
)

// GetCpt1de CPT1DE
func (r *registerTimedier5Type) GetCpt1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldCpt1deMask) != 0
}

// SetCpt1de CPT1DE
func (r *registerTimedier5Type) SetCpt1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldCpt1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldCpt1deMask)
	}
}

const (
	RegisterTimedier5FieldUpddeShift = 22
	RegisterTimedier5FieldUpddeMask  = 0x400000
)

// GetUpdde UPDDE
func (r *registerTimedier5Type) GetUpdde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldUpddeMask) != 0
}

// SetUpdde UPDDE
func (r *registerTimedier5Type) SetUpdde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldUpddeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldUpddeMask)
	}
}

const (
	RegisterTimedier5FieldRepdeShift = 20
	RegisterTimedier5FieldRepdeMask  = 0x100000
)

// GetRepde REPDE
func (r *registerTimedier5Type) GetRepde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldRepdeMask) != 0
}

// SetRepde REPDE
func (r *registerTimedier5Type) SetRepde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldRepdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldRepdeMask)
	}
}

const (
	RegisterTimedier5FieldCmp4deShift = 19
	RegisterTimedier5FieldCmp4deMask  = 0x80000
)

// GetCmp4de CMP4DE
func (r *registerTimedier5Type) GetCmp4de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldCmp4deMask) != 0
}

// SetCmp4de CMP4DE
func (r *registerTimedier5Type) SetCmp4de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldCmp4deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldCmp4deMask)
	}
}

const (
	RegisterTimedier5FieldCmp3deShift = 18
	RegisterTimedier5FieldCmp3deMask  = 0x40000
)

// GetCmp3de CMP3DE
func (r *registerTimedier5Type) GetCmp3de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldCmp3deMask) != 0
}

// SetCmp3de CMP3DE
func (r *registerTimedier5Type) SetCmp3de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldCmp3deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldCmp3deMask)
	}
}

const (
	RegisterTimedier5FieldCmp2deShift = 17
	RegisterTimedier5FieldCmp2deMask  = 0x20000
)

// GetCmp2de CMP2DE
func (r *registerTimedier5Type) GetCmp2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldCmp2deMask) != 0
}

// SetCmp2de CMP2DE
func (r *registerTimedier5Type) SetCmp2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldCmp2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldCmp2deMask)
	}
}

const (
	RegisterTimedier5FieldCmp1deShift = 16
	RegisterTimedier5FieldCmp1deMask  = 0x10000
)

// GetCmp1de CMP1DE
func (r *registerTimedier5Type) GetCmp1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldCmp1deMask) != 0
}

// SetCmp1de CMP1DE
func (r *registerTimedier5Type) SetCmp1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldCmp1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldCmp1deMask)
	}
}

const (
	RegisterTimedier5FieldDlyprtieShift = 14
	RegisterTimedier5FieldDlyprtieMask  = 0x4000
)

// GetDlyprtie DLYPRTIE
func (r *registerTimedier5Type) GetDlyprtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldDlyprtieMask) != 0
}

// SetDlyprtie DLYPRTIE
func (r *registerTimedier5Type) SetDlyprtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldDlyprtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldDlyprtieMask)
	}
}

const (
	RegisterTimedier5FieldRstieShift = 13
	RegisterTimedier5FieldRstieMask  = 0x2000
)

// GetRstie RSTIE
func (r *registerTimedier5Type) GetRstie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldRstieMask) != 0
}

// SetRstie RSTIE
func (r *registerTimedier5Type) SetRstie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldRstieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldRstieMask)
	}
}

const (
	RegisterTimedier5FieldRstx2ieShift = 12
	RegisterTimedier5FieldRstx2ieMask  = 0x1000
)

// GetRstx2ie RSTx2IE
func (r *registerTimedier5Type) GetRstx2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldRstx2ieMask) != 0
}

// SetRstx2ie RSTx2IE
func (r *registerTimedier5Type) SetRstx2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldRstx2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldRstx2ieMask)
	}
}

const (
	RegisterTimedier5FieldSetx2ieShift = 11
	RegisterTimedier5FieldSetx2ieMask  = 0x800
)

// GetSetx2ie SETx2IE
func (r *registerTimedier5Type) GetSetx2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldSetx2ieMask) != 0
}

// SetSetx2ie SETx2IE
func (r *registerTimedier5Type) SetSetx2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldSetx2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldSetx2ieMask)
	}
}

const (
	RegisterTimedier5FieldRstx1ieShift = 10
	RegisterTimedier5FieldRstx1ieMask  = 0x400
)

// GetRstx1ie RSTx1IE
func (r *registerTimedier5Type) GetRstx1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldRstx1ieMask) != 0
}

// SetRstx1ie RSTx1IE
func (r *registerTimedier5Type) SetRstx1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldRstx1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldRstx1ieMask)
	}
}

const (
	RegisterTimedier5FieldSet1xieShift = 9
	RegisterTimedier5FieldSet1xieMask  = 0x200
)

// GetSet1xie SET1xIE
func (r *registerTimedier5Type) GetSet1xie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldSet1xieMask) != 0
}

// SetSet1xie SET1xIE
func (r *registerTimedier5Type) SetSet1xie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldSet1xieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldSet1xieMask)
	}
}

const (
	RegisterTimedier5FieldCpt2ieShift = 8
	RegisterTimedier5FieldCpt2ieMask  = 0x100
)

// GetCpt2ie CPT2IE
func (r *registerTimedier5Type) GetCpt2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldCpt2ieMask) != 0
}

// SetCpt2ie CPT2IE
func (r *registerTimedier5Type) SetCpt2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldCpt2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldCpt2ieMask)
	}
}

const (
	RegisterTimedier5FieldCpt1ieShift = 7
	RegisterTimedier5FieldCpt1ieMask  = 0x80
)

// GetCpt1ie CPT1IE
func (r *registerTimedier5Type) GetCpt1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldCpt1ieMask) != 0
}

// SetCpt1ie CPT1IE
func (r *registerTimedier5Type) SetCpt1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldCpt1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldCpt1ieMask)
	}
}

const (
	RegisterTimedier5FieldUpdieShift = 6
	RegisterTimedier5FieldUpdieMask  = 0x40
)

// GetUpdie UPDIE
func (r *registerTimedier5Type) GetUpdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldUpdieMask) != 0
}

// SetUpdie UPDIE
func (r *registerTimedier5Type) SetUpdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldUpdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldUpdieMask)
	}
}

const (
	RegisterTimedier5FieldRepieShift = 4
	RegisterTimedier5FieldRepieMask  = 0x10
)

// GetRepie REPIE
func (r *registerTimedier5Type) GetRepie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldRepieMask) != 0
}

// SetRepie REPIE
func (r *registerTimedier5Type) SetRepie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldRepieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldRepieMask)
	}
}

const (
	RegisterTimedier5FieldCmp4ieShift = 3
	RegisterTimedier5FieldCmp4ieMask  = 0x8
)

// GetCmp4ie CMP4IE
func (r *registerTimedier5Type) GetCmp4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldCmp4ieMask) != 0
}

// SetCmp4ie CMP4IE
func (r *registerTimedier5Type) SetCmp4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldCmp4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldCmp4ieMask)
	}
}

const (
	RegisterTimedier5FieldCmp3ieShift = 2
	RegisterTimedier5FieldCmp3ieMask  = 0x4
)

// GetCmp3ie CMP3IE
func (r *registerTimedier5Type) GetCmp3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldCmp3ieMask) != 0
}

// SetCmp3ie CMP3IE
func (r *registerTimedier5Type) SetCmp3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldCmp3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldCmp3ieMask)
	}
}

const (
	RegisterTimedier5FieldCmp2ieShift = 1
	RegisterTimedier5FieldCmp2ieMask  = 0x2
)

// GetCmp2ie CMP2IE
func (r *registerTimedier5Type) GetCmp2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldCmp2ieMask) != 0
}

// SetCmp2ie CMP2IE
func (r *registerTimedier5Type) SetCmp2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldCmp2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldCmp2ieMask)
	}
}

const (
	RegisterTimedier5FieldCmp1ieShift = 0
	RegisterTimedier5FieldCmp1ieMask  = 0x1
)

// GetCmp1ie CMP1IE
func (r *registerTimedier5Type) GetCmp1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimedier5FieldCmp1ieMask) != 0
}

// SetCmp1ie CMP1IE
func (r *registerTimedier5Type) SetCmp1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimedier5FieldCmp1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimedier5FieldCmp1ieMask)
	}
}

// registerCnterType Timerx Counter Register
type registerCnterType uint32

const (
	RegisterCnterFieldCntxShift = 0
	RegisterCnterFieldCntxMask  = 0xffff
)

// GetCntx Timerx Counter value
func (r *registerCnterType) GetCntx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCnterFieldCntxMask) >> RegisterCnterFieldCntxShift)
}

// SetCntx Timerx Counter value
func (r *registerCnterType) SetCntx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCnterFieldCntxMask)|(uint32(value)<<RegisterCnterFieldCntxShift))
}

// registerPererType Timerx Period Register
type registerPererType uint32

const (
	RegisterPererFieldPerxShift = 0
	RegisterPererFieldPerxMask  = 0xffff
)

// GetPerx Timerx Period value
func (r *registerPererType) GetPerx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPererFieldPerxMask) >> RegisterPererFieldPerxShift)
}

// SetPerx Timerx Period value
func (r *registerPererType) SetPerx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPererFieldPerxMask)|(uint32(value)<<RegisterPererFieldPerxShift))
}

// registerReperType Timerx Repetition Register
type registerReperType uint32

const (
	RegisterReperFieldRepxShift = 0
	RegisterReperFieldRepxMask  = 0xff
)

// GetRepx Timerx Repetition counter value
func (r *registerReperType) GetRepx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterReperFieldRepxMask) >> RegisterReperFieldRepxShift)
}

// SetRepx Timerx Repetition counter value
func (r *registerReperType) SetRepx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterReperFieldRepxMask)|(uint32(value)<<RegisterReperFieldRepxShift))
}

// registerCmp1erType Timerx Compare 1 Register
type registerCmp1erType uint32

const (
	RegisterCmp1erFieldCmp1xShift = 0
	RegisterCmp1erFieldCmp1xMask  = 0xffff
)

// GetCmp1x Timerx Compare 1 value
func (r *registerCmp1erType) GetCmp1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1erFieldCmp1xMask) >> RegisterCmp1erFieldCmp1xShift)
}

// SetCmp1x Timerx Compare 1 value
func (r *registerCmp1erType) SetCmp1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1erFieldCmp1xMask)|(uint32(value)<<RegisterCmp1erFieldCmp1xShift))
}

// registerCmp1cerType Timerx Compare 1 Compound Register
type registerCmp1cerType uint32

const (
	RegisterCmp1cerFieldRepxShift = 16
	RegisterCmp1cerFieldRepxMask  = 0xff0000
)

// GetRepx Timerx Repetition value (aliased from HRTIM_REPx register)
func (r *registerCmp1cerType) GetRepx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1cerFieldRepxMask) >> RegisterCmp1cerFieldRepxShift)
}

// SetRepx Timerx Repetition value (aliased from HRTIM_REPx register)
func (r *registerCmp1cerType) SetRepx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1cerFieldRepxMask)|(uint32(value)<<RegisterCmp1cerFieldRepxShift))
}

const (
	RegisterCmp1cerFieldCmp1xShift = 0
	RegisterCmp1cerFieldCmp1xMask  = 0xffff
)

// GetCmp1x Timerx Compare 1 value
func (r *registerCmp1cerType) GetCmp1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1cerFieldCmp1xMask) >> RegisterCmp1cerFieldCmp1xShift)
}

// SetCmp1x Timerx Compare 1 value
func (r *registerCmp1cerType) SetCmp1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1cerFieldCmp1xMask)|(uint32(value)<<RegisterCmp1cerFieldCmp1xShift))
}

// registerCmp2erType Timerx Compare 2 Register
type registerCmp2erType uint32

const (
	RegisterCmp2erFieldCmp2xShift = 0
	RegisterCmp2erFieldCmp2xMask  = 0xffff
)

// GetCmp2x Timerx Compare 2 value
func (r *registerCmp2erType) GetCmp2x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp2erFieldCmp2xMask) >> RegisterCmp2erFieldCmp2xShift)
}

// SetCmp2x Timerx Compare 2 value
func (r *registerCmp2erType) SetCmp2x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp2erFieldCmp2xMask)|(uint32(value)<<RegisterCmp2erFieldCmp2xShift))
}

// registerCmp3erType Timerx Compare 3 Register
type registerCmp3erType uint32

const (
	RegisterCmp3erFieldCmp3xShift = 0
	RegisterCmp3erFieldCmp3xMask  = 0xffff
)

// GetCmp3x Timerx Compare 3 value
func (r *registerCmp3erType) GetCmp3x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp3erFieldCmp3xMask) >> RegisterCmp3erFieldCmp3xShift)
}

// SetCmp3x Timerx Compare 3 value
func (r *registerCmp3erType) SetCmp3x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp3erFieldCmp3xMask)|(uint32(value)<<RegisterCmp3erFieldCmp3xShift))
}

// registerCmp4erType Timerx Compare 4 Register
type registerCmp4erType uint32

const (
	RegisterCmp4erFieldCmp4xShift = 0
	RegisterCmp4erFieldCmp4xMask  = 0xffff
)

// GetCmp4x Timerx Compare 4 value
func (r *registerCmp4erType) GetCmp4x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp4erFieldCmp4xMask) >> RegisterCmp4erFieldCmp4xShift)
}

// SetCmp4x Timerx Compare 4 value
func (r *registerCmp4erType) SetCmp4x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp4erFieldCmp4xMask)|(uint32(value)<<RegisterCmp4erFieldCmp4xShift))
}

// registerCpt1erType Timerx Capture 1 Register
type registerCpt1erType uint32

const (
	RegisterCpt1erFieldCpt1xShift = 0
	RegisterCpt1erFieldCpt1xMask  = 0xffff
)

// GetCpt1x Timerx Capture 1 value
func (r *registerCpt1erType) GetCpt1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpt1erFieldCpt1xMask) >> RegisterCpt1erFieldCpt1xShift)
}

// SetCpt1x Timerx Capture 1 value
func (r *registerCpt1erType) SetCpt1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpt1erFieldCpt1xMask)|(uint32(value)<<RegisterCpt1erFieldCpt1xShift))
}

// registerCpt2erType Timerx Capture 2 Register
type registerCpt2erType uint32

const (
	RegisterCpt2erFieldCpt2xShift = 0
	RegisterCpt2erFieldCpt2xMask  = 0xffff
)

// GetCpt2x Timerx Capture 2 value
func (r *registerCpt2erType) GetCpt2x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpt2erFieldCpt2xMask) >> RegisterCpt2erFieldCpt2xShift)
}

// SetCpt2x Timerx Capture 2 value
func (r *registerCpt2erType) SetCpt2x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpt2erFieldCpt2xMask)|(uint32(value)<<RegisterCpt2erFieldCpt2xShift))
}

// registerDterType Timerx Deadtime Register
type registerDterType uint32

const (
	RegisterDterFieldDtflkxShift = 31
	RegisterDterFieldDtflkxMask  = 0x80000000
)

// GetDtflkx Deadtime Falling Lock
func (r *registerDterType) GetDtflkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDterFieldDtflkxMask) != 0
}

// SetDtflkx Deadtime Falling Lock
func (r *registerDterType) SetDtflkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDterFieldDtflkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDterFieldDtflkxMask)
	}
}

const (
	RegisterDterFieldDtfslkxShift = 30
	RegisterDterFieldDtfslkxMask  = 0x40000000
)

// GetDtfslkx Deadtime Falling Sign Lock
func (r *registerDterType) GetDtfslkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDterFieldDtfslkxMask) != 0
}

// SetDtfslkx Deadtime Falling Sign Lock
func (r *registerDterType) SetDtfslkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDterFieldDtfslkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDterFieldDtfslkxMask)
	}
}

const (
	RegisterDterFieldSdtfxShift = 25
	RegisterDterFieldSdtfxMask  = 0x2000000
)

// GetSdtfx Sign Deadtime Falling value
func (r *registerDterType) GetSdtfx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDterFieldSdtfxMask) != 0
}

// SetSdtfx Sign Deadtime Falling value
func (r *registerDterType) SetSdtfx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDterFieldSdtfxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDterFieldSdtfxMask)
	}
}

const (
	RegisterDterFieldDtfxShift = 16
	RegisterDterFieldDtfxMask  = 0x1ff0000
)

// GetDtfx Deadtime Falling value
func (r *registerDterType) GetDtfx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDterFieldDtfxMask) >> RegisterDterFieldDtfxShift)
}

// SetDtfx Deadtime Falling value
func (r *registerDterType) SetDtfx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDterFieldDtfxMask)|(uint32(value)<<RegisterDterFieldDtfxShift))
}

const (
	RegisterDterFieldDtrlkxShift = 15
	RegisterDterFieldDtrlkxMask  = 0x8000
)

// GetDtrlkx Deadtime Rising Lock
func (r *registerDterType) GetDtrlkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDterFieldDtrlkxMask) != 0
}

// SetDtrlkx Deadtime Rising Lock
func (r *registerDterType) SetDtrlkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDterFieldDtrlkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDterFieldDtrlkxMask)
	}
}

const (
	RegisterDterFieldDtrslkxShift = 14
	RegisterDterFieldDtrslkxMask  = 0x4000
)

// GetDtrslkx Deadtime Rising Sign Lock
func (r *registerDterType) GetDtrslkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDterFieldDtrslkxMask) != 0
}

// SetDtrslkx Deadtime Rising Sign Lock
func (r *registerDterType) SetDtrslkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDterFieldDtrslkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDterFieldDtrslkxMask)
	}
}

const (
	RegisterDterFieldDtprscShift = 10
	RegisterDterFieldDtprscMask  = 0x1c00
)

// GetDtprsc Deadtime Prescaler
func (r *registerDterType) GetDtprsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDterFieldDtprscMask) >> RegisterDterFieldDtprscShift)
}

// SetDtprsc Deadtime Prescaler
func (r *registerDterType) SetDtprsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDterFieldDtprscMask)|(uint32(value)<<RegisterDterFieldDtprscShift))
}

const (
	RegisterDterFieldSdtrxShift = 9
	RegisterDterFieldSdtrxMask  = 0x200
)

// GetSdtrx Sign Deadtime Rising value
func (r *registerDterType) GetSdtrx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDterFieldSdtrxMask) != 0
}

// SetSdtrx Sign Deadtime Rising value
func (r *registerDterType) SetSdtrx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDterFieldSdtrxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDterFieldSdtrxMask)
	}
}

const (
	RegisterDterFieldDtrxShift = 0
	RegisterDterFieldDtrxMask  = 0x1ff
)

// GetDtrx Deadtime Rising value
func (r *registerDterType) GetDtrx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDterFieldDtrxMask) >> RegisterDterFieldDtrxShift)
}

// SetDtrx Deadtime Rising value
func (r *registerDterType) SetDtrx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDterFieldDtrxMask)|(uint32(value)<<RegisterDterFieldDtrxShift))
}

// registerSete1rType Timerx Output1 Set Register
type registerSete1rType uint32

const (
	RegisterSete1rFieldUpdateShift = 31
	RegisterSete1rFieldUpdateMask  = 0x80000000
)

// GetUpdate Registers update (transfer preload to active)
func (r *registerSete1rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldUpdateMask) != 0
}

// SetUpdate Registers update (transfer preload to active)
func (r *registerSete1rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldUpdateMask)
	}
}

const (
	RegisterSete1rFieldExtevnt10Shift = 30
	RegisterSete1rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 External Event 10
func (r *registerSete1rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 External Event 10
func (r *registerSete1rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldExtevnt10Mask)
	}
}

const (
	RegisterSete1rFieldExtevnt9Shift = 29
	RegisterSete1rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 External Event 9
func (r *registerSete1rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 External Event 9
func (r *registerSete1rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldExtevnt9Mask)
	}
}

const (
	RegisterSete1rFieldExtevnt8Shift = 28
	RegisterSete1rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 External Event 8
func (r *registerSete1rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 External Event 8
func (r *registerSete1rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldExtevnt8Mask)
	}
}

const (
	RegisterSete1rFieldExtevnt7Shift = 27
	RegisterSete1rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 External Event 7
func (r *registerSete1rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 External Event 7
func (r *registerSete1rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldExtevnt7Mask)
	}
}

const (
	RegisterSete1rFieldExtevnt6Shift = 26
	RegisterSete1rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 External Event 6
func (r *registerSete1rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 External Event 6
func (r *registerSete1rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldExtevnt6Mask)
	}
}

const (
	RegisterSete1rFieldExtevnt5Shift = 25
	RegisterSete1rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 External Event 5
func (r *registerSete1rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 External Event 5
func (r *registerSete1rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldExtevnt5Mask)
	}
}

const (
	RegisterSete1rFieldExtevnt4Shift = 24
	RegisterSete1rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 External Event 4
func (r *registerSete1rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 External Event 4
func (r *registerSete1rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldExtevnt4Mask)
	}
}

const (
	RegisterSete1rFieldExtevnt3Shift = 23
	RegisterSete1rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 External Event 3
func (r *registerSete1rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 External Event 3
func (r *registerSete1rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldExtevnt3Mask)
	}
}

const (
	RegisterSete1rFieldExtevnt2Shift = 22
	RegisterSete1rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 External Event 2
func (r *registerSete1rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 External Event 2
func (r *registerSete1rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldExtevnt2Mask)
	}
}

const (
	RegisterSete1rFieldExtevnt1Shift = 21
	RegisterSete1rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 External Event 1
func (r *registerSete1rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 External Event 1
func (r *registerSete1rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldExtevnt1Mask)
	}
}

const (
	RegisterSete1rFieldTimevnt9Shift = 20
	RegisterSete1rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 Timer Event 9
func (r *registerSete1rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 Timer Event 9
func (r *registerSete1rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldTimevnt9Mask)
	}
}

const (
	RegisterSete1rFieldTimevnt8Shift = 19
	RegisterSete1rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 Timer Event 8
func (r *registerSete1rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 Timer Event 8
func (r *registerSete1rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldTimevnt8Mask)
	}
}

const (
	RegisterSete1rFieldTimevnt7Shift = 18
	RegisterSete1rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 Timer Event 7
func (r *registerSete1rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 Timer Event 7
func (r *registerSete1rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldTimevnt7Mask)
	}
}

const (
	RegisterSete1rFieldTimevnt6Shift = 17
	RegisterSete1rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 Timer Event 6
func (r *registerSete1rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 Timer Event 6
func (r *registerSete1rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldTimevnt6Mask)
	}
}

const (
	RegisterSete1rFieldTimevnt5Shift = 16
	RegisterSete1rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 Timer Event 5
func (r *registerSete1rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 Timer Event 5
func (r *registerSete1rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldTimevnt5Mask)
	}
}

const (
	RegisterSete1rFieldTimevnt4Shift = 15
	RegisterSete1rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 Timer Event 4
func (r *registerSete1rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 Timer Event 4
func (r *registerSete1rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldTimevnt4Mask)
	}
}

const (
	RegisterSete1rFieldTimevnt3Shift = 14
	RegisterSete1rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 Timer Event 3
func (r *registerSete1rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 Timer Event 3
func (r *registerSete1rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldTimevnt3Mask)
	}
}

const (
	RegisterSete1rFieldTimevnt2Shift = 13
	RegisterSete1rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 Timer Event 2
func (r *registerSete1rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 Timer Event 2
func (r *registerSete1rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldTimevnt2Mask)
	}
}

const (
	RegisterSete1rFieldTimevnt1Shift = 12
	RegisterSete1rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 Timer Event 1
func (r *registerSete1rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 Timer Event 1
func (r *registerSete1rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldTimevnt1Mask)
	}
}

const (
	RegisterSete1rFieldMstcmp4Shift = 11
	RegisterSete1rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 Master Compare 4
func (r *registerSete1rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 Master Compare 4
func (r *registerSete1rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldMstcmp4Mask)
	}
}

const (
	RegisterSete1rFieldMstcmp3Shift = 10
	RegisterSete1rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 Master Compare 3
func (r *registerSete1rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 Master Compare 3
func (r *registerSete1rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldMstcmp3Mask)
	}
}

const (
	RegisterSete1rFieldMstcmp2Shift = 9
	RegisterSete1rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 Master Compare 2
func (r *registerSete1rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 Master Compare 2
func (r *registerSete1rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldMstcmp2Mask)
	}
}

const (
	RegisterSete1rFieldMstcmp1Shift = 8
	RegisterSete1rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 Master Compare 1
func (r *registerSete1rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 Master Compare 1
func (r *registerSete1rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldMstcmp1Mask)
	}
}

const (
	RegisterSete1rFieldMstperShift = 7
	RegisterSete1rFieldMstperMask  = 0x80
)

// GetMstper Master Period
func (r *registerSete1rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldMstperMask) != 0
}

// SetMstper Master Period
func (r *registerSete1rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldMstperMask)
	}
}

const (
	RegisterSete1rFieldCmp4Shift = 6
	RegisterSete1rFieldCmp4Mask  = 0x40
)

// GetCmp4 Timer A compare 4
func (r *registerSete1rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldCmp4Mask) != 0
}

// SetCmp4 Timer A compare 4
func (r *registerSete1rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldCmp4Mask)
	}
}

const (
	RegisterSete1rFieldCmp3Shift = 5
	RegisterSete1rFieldCmp3Mask  = 0x20
)

// GetCmp3 Timer A compare 3
func (r *registerSete1rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldCmp3Mask) != 0
}

// SetCmp3 Timer A compare 3
func (r *registerSete1rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldCmp3Mask)
	}
}

const (
	RegisterSete1rFieldCmp2Shift = 4
	RegisterSete1rFieldCmp2Mask  = 0x10
)

// GetCmp2 Timer A compare 2
func (r *registerSete1rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldCmp2Mask) != 0
}

// SetCmp2 Timer A compare 2
func (r *registerSete1rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldCmp2Mask)
	}
}

const (
	RegisterSete1rFieldCmp1Shift = 3
	RegisterSete1rFieldCmp1Mask  = 0x8
)

// GetCmp1 Timer A compare 1
func (r *registerSete1rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldCmp1Mask) != 0
}

// SetCmp1 Timer A compare 1
func (r *registerSete1rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldCmp1Mask)
	}
}

const (
	RegisterSete1rFieldPerShift = 2
	RegisterSete1rFieldPerMask  = 0x4
)

// GetPer Timer A Period
func (r *registerSete1rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldPerMask) != 0
}

// SetPer Timer A Period
func (r *registerSete1rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldPerMask)
	}
}

const (
	RegisterSete1rFieldResyncShift = 1
	RegisterSete1rFieldResyncMask  = 0x2
)

// GetResync Timer A resynchronizaton
func (r *registerSete1rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldResyncMask) != 0
}

// SetResync Timer A resynchronizaton
func (r *registerSete1rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldResyncMask)
	}
}

const (
	RegisterSete1rFieldSstShift = 0
	RegisterSete1rFieldSstMask  = 0x1
)

// GetSst Software Set trigger
func (r *registerSete1rType) GetSst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete1rFieldSstMask) != 0
}

// SetSst Software Set trigger
func (r *registerSete1rType) SetSst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete1rFieldSstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete1rFieldSstMask)
	}
}

// registerRste1rType Timerx Output1 Reset Register
type registerRste1rType uint32

const (
	RegisterRste1rFieldUpdateShift = 31
	RegisterRste1rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerRste1rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerRste1rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldUpdateMask)
	}
}

const (
	RegisterRste1rFieldExtevnt10Shift = 30
	RegisterRste1rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerRste1rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerRste1rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldExtevnt10Mask)
	}
}

const (
	RegisterRste1rFieldExtevnt9Shift = 29
	RegisterRste1rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerRste1rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerRste1rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldExtevnt9Mask)
	}
}

const (
	RegisterRste1rFieldExtevnt8Shift = 28
	RegisterRste1rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerRste1rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerRste1rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldExtevnt8Mask)
	}
}

const (
	RegisterRste1rFieldExtevnt7Shift = 27
	RegisterRste1rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerRste1rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerRste1rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldExtevnt7Mask)
	}
}

const (
	RegisterRste1rFieldExtevnt6Shift = 26
	RegisterRste1rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerRste1rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerRste1rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldExtevnt6Mask)
	}
}

const (
	RegisterRste1rFieldExtevnt5Shift = 25
	RegisterRste1rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerRste1rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerRste1rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldExtevnt5Mask)
	}
}

const (
	RegisterRste1rFieldExtevnt4Shift = 24
	RegisterRste1rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerRste1rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerRste1rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldExtevnt4Mask)
	}
}

const (
	RegisterRste1rFieldExtevnt3Shift = 23
	RegisterRste1rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerRste1rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerRste1rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldExtevnt3Mask)
	}
}

const (
	RegisterRste1rFieldExtevnt2Shift = 22
	RegisterRste1rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerRste1rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerRste1rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldExtevnt2Mask)
	}
}

const (
	RegisterRste1rFieldExtevnt1Shift = 21
	RegisterRste1rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerRste1rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerRste1rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldExtevnt1Mask)
	}
}

const (
	RegisterRste1rFieldTimevnt9Shift = 20
	RegisterRste1rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerRste1rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerRste1rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldTimevnt9Mask)
	}
}

const (
	RegisterRste1rFieldTimevnt8Shift = 19
	RegisterRste1rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerRste1rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerRste1rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldTimevnt8Mask)
	}
}

const (
	RegisterRste1rFieldTimevnt7Shift = 18
	RegisterRste1rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerRste1rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerRste1rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldTimevnt7Mask)
	}
}

const (
	RegisterRste1rFieldTimevnt6Shift = 17
	RegisterRste1rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerRste1rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerRste1rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldTimevnt6Mask)
	}
}

const (
	RegisterRste1rFieldTimevnt5Shift = 16
	RegisterRste1rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerRste1rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerRste1rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldTimevnt5Mask)
	}
}

const (
	RegisterRste1rFieldTimevnt4Shift = 15
	RegisterRste1rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerRste1rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerRste1rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldTimevnt4Mask)
	}
}

const (
	RegisterRste1rFieldTimevnt3Shift = 14
	RegisterRste1rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerRste1rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerRste1rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldTimevnt3Mask)
	}
}

const (
	RegisterRste1rFieldTimevnt2Shift = 13
	RegisterRste1rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerRste1rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerRste1rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldTimevnt2Mask)
	}
}

const (
	RegisterRste1rFieldTimevnt1Shift = 12
	RegisterRste1rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerRste1rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerRste1rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldTimevnt1Mask)
	}
}

const (
	RegisterRste1rFieldMstcmp4Shift = 11
	RegisterRste1rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerRste1rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerRste1rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldMstcmp4Mask)
	}
}

const (
	RegisterRste1rFieldMstcmp3Shift = 10
	RegisterRste1rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerRste1rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerRste1rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldMstcmp3Mask)
	}
}

const (
	RegisterRste1rFieldMstcmp2Shift = 9
	RegisterRste1rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerRste1rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerRste1rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldMstcmp2Mask)
	}
}

const (
	RegisterRste1rFieldMstcmp1Shift = 8
	RegisterRste1rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerRste1rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerRste1rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldMstcmp1Mask)
	}
}

const (
	RegisterRste1rFieldMstperShift = 7
	RegisterRste1rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerRste1rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerRste1rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldMstperMask)
	}
}

const (
	RegisterRste1rFieldCmp4Shift = 6
	RegisterRste1rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerRste1rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerRste1rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldCmp4Mask)
	}
}

const (
	RegisterRste1rFieldCmp3Shift = 5
	RegisterRste1rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerRste1rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerRste1rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldCmp3Mask)
	}
}

const (
	RegisterRste1rFieldCmp2Shift = 4
	RegisterRste1rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerRste1rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerRste1rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldCmp2Mask)
	}
}

const (
	RegisterRste1rFieldCmp1Shift = 3
	RegisterRste1rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerRste1rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerRste1rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldCmp1Mask)
	}
}

const (
	RegisterRste1rFieldPerShift = 2
	RegisterRste1rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerRste1rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldPerMask) != 0
}

// SetPer PER
func (r *registerRste1rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldPerMask)
	}
}

const (
	RegisterRste1rFieldResyncShift = 1
	RegisterRste1rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerRste1rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerRste1rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldResyncMask)
	}
}

const (
	RegisterRste1rFieldSrtShift = 0
	RegisterRste1rFieldSrtMask  = 0x1
)

// GetSrt SRT
func (r *registerRste1rType) GetSrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste1rFieldSrtMask) != 0
}

// SetSrt SRT
func (r *registerRste1rType) SetSrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste1rFieldSrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste1rFieldSrtMask)
	}
}

// registerSete2rType Timerx Output2 Set Register
type registerSete2rType uint32

const (
	RegisterSete2rFieldUpdateShift = 31
	RegisterSete2rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerSete2rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerSete2rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldUpdateMask)
	}
}

const (
	RegisterSete2rFieldExtevnt10Shift = 30
	RegisterSete2rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerSete2rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerSete2rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldExtevnt10Mask)
	}
}

const (
	RegisterSete2rFieldExtevnt9Shift = 29
	RegisterSete2rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerSete2rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerSete2rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldExtevnt9Mask)
	}
}

const (
	RegisterSete2rFieldExtevnt8Shift = 28
	RegisterSete2rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerSete2rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerSete2rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldExtevnt8Mask)
	}
}

const (
	RegisterSete2rFieldExtevnt7Shift = 27
	RegisterSete2rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerSete2rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerSete2rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldExtevnt7Mask)
	}
}

const (
	RegisterSete2rFieldExtevnt6Shift = 26
	RegisterSete2rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerSete2rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerSete2rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldExtevnt6Mask)
	}
}

const (
	RegisterSete2rFieldExtevnt5Shift = 25
	RegisterSete2rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerSete2rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerSete2rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldExtevnt5Mask)
	}
}

const (
	RegisterSete2rFieldExtevnt4Shift = 24
	RegisterSete2rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerSete2rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerSete2rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldExtevnt4Mask)
	}
}

const (
	RegisterSete2rFieldExtevnt3Shift = 23
	RegisterSete2rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerSete2rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerSete2rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldExtevnt3Mask)
	}
}

const (
	RegisterSete2rFieldExtevnt2Shift = 22
	RegisterSete2rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerSete2rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerSete2rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldExtevnt2Mask)
	}
}

const (
	RegisterSete2rFieldExtevnt1Shift = 21
	RegisterSete2rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerSete2rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerSete2rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldExtevnt1Mask)
	}
}

const (
	RegisterSete2rFieldTimevnt9Shift = 20
	RegisterSete2rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerSete2rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerSete2rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldTimevnt9Mask)
	}
}

const (
	RegisterSete2rFieldTimevnt8Shift = 19
	RegisterSete2rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerSete2rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerSete2rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldTimevnt8Mask)
	}
}

const (
	RegisterSete2rFieldTimevnt7Shift = 18
	RegisterSete2rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerSete2rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerSete2rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldTimevnt7Mask)
	}
}

const (
	RegisterSete2rFieldTimevnt6Shift = 17
	RegisterSete2rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerSete2rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerSete2rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldTimevnt6Mask)
	}
}

const (
	RegisterSete2rFieldTimevnt5Shift = 16
	RegisterSete2rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerSete2rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerSete2rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldTimevnt5Mask)
	}
}

const (
	RegisterSete2rFieldTimevnt4Shift = 15
	RegisterSete2rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerSete2rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerSete2rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldTimevnt4Mask)
	}
}

const (
	RegisterSete2rFieldTimevnt3Shift = 14
	RegisterSete2rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerSete2rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerSete2rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldTimevnt3Mask)
	}
}

const (
	RegisterSete2rFieldTimevnt2Shift = 13
	RegisterSete2rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerSete2rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerSete2rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldTimevnt2Mask)
	}
}

const (
	RegisterSete2rFieldTimevnt1Shift = 12
	RegisterSete2rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerSete2rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerSete2rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldTimevnt1Mask)
	}
}

const (
	RegisterSete2rFieldMstcmp4Shift = 11
	RegisterSete2rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerSete2rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerSete2rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldMstcmp4Mask)
	}
}

const (
	RegisterSete2rFieldMstcmp3Shift = 10
	RegisterSete2rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerSete2rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerSete2rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldMstcmp3Mask)
	}
}

const (
	RegisterSete2rFieldMstcmp2Shift = 9
	RegisterSete2rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerSete2rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerSete2rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldMstcmp2Mask)
	}
}

const (
	RegisterSete2rFieldMstcmp1Shift = 8
	RegisterSete2rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerSete2rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerSete2rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldMstcmp1Mask)
	}
}

const (
	RegisterSete2rFieldMstperShift = 7
	RegisterSete2rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerSete2rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerSete2rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldMstperMask)
	}
}

const (
	RegisterSete2rFieldCmp4Shift = 6
	RegisterSete2rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerSete2rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerSete2rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldCmp4Mask)
	}
}

const (
	RegisterSete2rFieldCmp3Shift = 5
	RegisterSete2rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerSete2rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerSete2rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldCmp3Mask)
	}
}

const (
	RegisterSete2rFieldCmp2Shift = 4
	RegisterSete2rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerSete2rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerSete2rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldCmp2Mask)
	}
}

const (
	RegisterSete2rFieldCmp1Shift = 3
	RegisterSete2rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerSete2rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerSete2rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldCmp1Mask)
	}
}

const (
	RegisterSete2rFieldPerShift = 2
	RegisterSete2rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerSete2rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldPerMask) != 0
}

// SetPer PER
func (r *registerSete2rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldPerMask)
	}
}

const (
	RegisterSete2rFieldResyncShift = 1
	RegisterSete2rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerSete2rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerSete2rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldResyncMask)
	}
}

const (
	RegisterSete2rFieldSstShift = 0
	RegisterSete2rFieldSstMask  = 0x1
)

// GetSst SST
func (r *registerSete2rType) GetSst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSete2rFieldSstMask) != 0
}

// SetSst SST
func (r *registerSete2rType) SetSst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSete2rFieldSstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSete2rFieldSstMask)
	}
}

// registerRste2rType Timerx Output2 Reset Register
type registerRste2rType uint32

const (
	RegisterRste2rFieldUpdateShift = 31
	RegisterRste2rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerRste2rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerRste2rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldUpdateMask)
	}
}

const (
	RegisterRste2rFieldExtevnt10Shift = 30
	RegisterRste2rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerRste2rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerRste2rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldExtevnt10Mask)
	}
}

const (
	RegisterRste2rFieldExtevnt9Shift = 29
	RegisterRste2rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerRste2rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerRste2rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldExtevnt9Mask)
	}
}

const (
	RegisterRste2rFieldExtevnt8Shift = 28
	RegisterRste2rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerRste2rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerRste2rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldExtevnt8Mask)
	}
}

const (
	RegisterRste2rFieldExtevnt7Shift = 27
	RegisterRste2rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerRste2rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerRste2rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldExtevnt7Mask)
	}
}

const (
	RegisterRste2rFieldExtevnt6Shift = 26
	RegisterRste2rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerRste2rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerRste2rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldExtevnt6Mask)
	}
}

const (
	RegisterRste2rFieldExtevnt5Shift = 25
	RegisterRste2rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerRste2rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerRste2rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldExtevnt5Mask)
	}
}

const (
	RegisterRste2rFieldExtevnt4Shift = 24
	RegisterRste2rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerRste2rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerRste2rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldExtevnt4Mask)
	}
}

const (
	RegisterRste2rFieldExtevnt3Shift = 23
	RegisterRste2rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerRste2rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerRste2rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldExtevnt3Mask)
	}
}

const (
	RegisterRste2rFieldExtevnt2Shift = 22
	RegisterRste2rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerRste2rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerRste2rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldExtevnt2Mask)
	}
}

const (
	RegisterRste2rFieldExtevnt1Shift = 21
	RegisterRste2rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerRste2rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerRste2rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldExtevnt1Mask)
	}
}

const (
	RegisterRste2rFieldTimevnt9Shift = 20
	RegisterRste2rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerRste2rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerRste2rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldTimevnt9Mask)
	}
}

const (
	RegisterRste2rFieldTimevnt8Shift = 19
	RegisterRste2rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerRste2rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerRste2rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldTimevnt8Mask)
	}
}

const (
	RegisterRste2rFieldTimevnt7Shift = 18
	RegisterRste2rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerRste2rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerRste2rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldTimevnt7Mask)
	}
}

const (
	RegisterRste2rFieldTimevnt6Shift = 17
	RegisterRste2rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerRste2rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerRste2rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldTimevnt6Mask)
	}
}

const (
	RegisterRste2rFieldTimevnt5Shift = 16
	RegisterRste2rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerRste2rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerRste2rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldTimevnt5Mask)
	}
}

const (
	RegisterRste2rFieldTimevnt4Shift = 15
	RegisterRste2rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerRste2rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerRste2rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldTimevnt4Mask)
	}
}

const (
	RegisterRste2rFieldTimevnt3Shift = 14
	RegisterRste2rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerRste2rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerRste2rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldTimevnt3Mask)
	}
}

const (
	RegisterRste2rFieldTimevnt2Shift = 13
	RegisterRste2rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerRste2rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerRste2rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldTimevnt2Mask)
	}
}

const (
	RegisterRste2rFieldTimevnt1Shift = 12
	RegisterRste2rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerRste2rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerRste2rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldTimevnt1Mask)
	}
}

const (
	RegisterRste2rFieldMstcmp4Shift = 11
	RegisterRste2rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerRste2rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerRste2rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldMstcmp4Mask)
	}
}

const (
	RegisterRste2rFieldMstcmp3Shift = 10
	RegisterRste2rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerRste2rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerRste2rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldMstcmp3Mask)
	}
}

const (
	RegisterRste2rFieldMstcmp2Shift = 9
	RegisterRste2rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerRste2rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerRste2rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldMstcmp2Mask)
	}
}

const (
	RegisterRste2rFieldMstcmp1Shift = 8
	RegisterRste2rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerRste2rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerRste2rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldMstcmp1Mask)
	}
}

const (
	RegisterRste2rFieldMstperShift = 7
	RegisterRste2rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerRste2rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerRste2rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldMstperMask)
	}
}

const (
	RegisterRste2rFieldCmp4Shift = 6
	RegisterRste2rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerRste2rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerRste2rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldCmp4Mask)
	}
}

const (
	RegisterRste2rFieldCmp3Shift = 5
	RegisterRste2rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerRste2rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerRste2rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldCmp3Mask)
	}
}

const (
	RegisterRste2rFieldCmp2Shift = 4
	RegisterRste2rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerRste2rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerRste2rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldCmp2Mask)
	}
}

const (
	RegisterRste2rFieldCmp1Shift = 3
	RegisterRste2rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerRste2rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerRste2rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldCmp1Mask)
	}
}

const (
	RegisterRste2rFieldPerShift = 2
	RegisterRste2rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerRste2rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldPerMask) != 0
}

// SetPer PER
func (r *registerRste2rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldPerMask)
	}
}

const (
	RegisterRste2rFieldResyncShift = 1
	RegisterRste2rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerRste2rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerRste2rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldResyncMask)
	}
}

const (
	RegisterRste2rFieldSrtShift = 0
	RegisterRste2rFieldSrtMask  = 0x1
)

// GetSrt SRT
func (r *registerRste2rType) GetSrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRste2rFieldSrtMask) != 0
}

// SetSrt SRT
func (r *registerRste2rType) SetSrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRste2rFieldSrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRste2rFieldSrtMask)
	}
}

// registerEefer1Type Timerx External Event Filtering Register 1
type registerEefer1Type uint32

const (
	RegisterEefer1FieldEe5fltrShift = 25
	RegisterEefer1FieldEe5fltrMask  = 0x1e000000
)

// GetEe5fltr External Event 5 filter
func (r *registerEefer1Type) GetEe5fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefer1FieldEe5fltrMask) >> RegisterEefer1FieldEe5fltrShift)
}

// SetEe5fltr External Event 5 filter
func (r *registerEefer1Type) SetEe5fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefer1FieldEe5fltrMask)|(uint32(value)<<RegisterEefer1FieldEe5fltrShift))
}

const (
	RegisterEefer1FieldEe5ltchShift = 24
	RegisterEefer1FieldEe5ltchMask  = 0x1000000
)

// GetEe5ltch External Event 5 latch
func (r *registerEefer1Type) GetEe5ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefer1FieldEe5ltchMask) != 0
}

// SetEe5ltch External Event 5 latch
func (r *registerEefer1Type) SetEe5ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefer1FieldEe5ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefer1FieldEe5ltchMask)
	}
}

const (
	RegisterEefer1FieldEe4fltrShift = 19
	RegisterEefer1FieldEe4fltrMask  = 0x780000
)

// GetEe4fltr External Event 4 filter
func (r *registerEefer1Type) GetEe4fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefer1FieldEe4fltrMask) >> RegisterEefer1FieldEe4fltrShift)
}

// SetEe4fltr External Event 4 filter
func (r *registerEefer1Type) SetEe4fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefer1FieldEe4fltrMask)|(uint32(value)<<RegisterEefer1FieldEe4fltrShift))
}

const (
	RegisterEefer1FieldEe4ltchShift = 18
	RegisterEefer1FieldEe4ltchMask  = 0x40000
)

// GetEe4ltch External Event 4 latch
func (r *registerEefer1Type) GetEe4ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefer1FieldEe4ltchMask) != 0
}

// SetEe4ltch External Event 4 latch
func (r *registerEefer1Type) SetEe4ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefer1FieldEe4ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefer1FieldEe4ltchMask)
	}
}

const (
	RegisterEefer1FieldEe3fltrShift = 13
	RegisterEefer1FieldEe3fltrMask  = 0x1e000
)

// GetEe3fltr External Event 3 filter
func (r *registerEefer1Type) GetEe3fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefer1FieldEe3fltrMask) >> RegisterEefer1FieldEe3fltrShift)
}

// SetEe3fltr External Event 3 filter
func (r *registerEefer1Type) SetEe3fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefer1FieldEe3fltrMask)|(uint32(value)<<RegisterEefer1FieldEe3fltrShift))
}

const (
	RegisterEefer1FieldEe3ltchShift = 12
	RegisterEefer1FieldEe3ltchMask  = 0x1000
)

// GetEe3ltch External Event 3 latch
func (r *registerEefer1Type) GetEe3ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefer1FieldEe3ltchMask) != 0
}

// SetEe3ltch External Event 3 latch
func (r *registerEefer1Type) SetEe3ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefer1FieldEe3ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefer1FieldEe3ltchMask)
	}
}

const (
	RegisterEefer1FieldEe2fltrShift = 7
	RegisterEefer1FieldEe2fltrMask  = 0x780
)

// GetEe2fltr External Event 2 filter
func (r *registerEefer1Type) GetEe2fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefer1FieldEe2fltrMask) >> RegisterEefer1FieldEe2fltrShift)
}

// SetEe2fltr External Event 2 filter
func (r *registerEefer1Type) SetEe2fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefer1FieldEe2fltrMask)|(uint32(value)<<RegisterEefer1FieldEe2fltrShift))
}

const (
	RegisterEefer1FieldEe2ltchShift = 6
	RegisterEefer1FieldEe2ltchMask  = 0x40
)

// GetEe2ltch External Event 2 latch
func (r *registerEefer1Type) GetEe2ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefer1FieldEe2ltchMask) != 0
}

// SetEe2ltch External Event 2 latch
func (r *registerEefer1Type) SetEe2ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefer1FieldEe2ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefer1FieldEe2ltchMask)
	}
}

const (
	RegisterEefer1FieldEe1fltrShift = 1
	RegisterEefer1FieldEe1fltrMask  = 0x1e
)

// GetEe1fltr External Event 1 filter
func (r *registerEefer1Type) GetEe1fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefer1FieldEe1fltrMask) >> RegisterEefer1FieldEe1fltrShift)
}

// SetEe1fltr External Event 1 filter
func (r *registerEefer1Type) SetEe1fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefer1FieldEe1fltrMask)|(uint32(value)<<RegisterEefer1FieldEe1fltrShift))
}

const (
	RegisterEefer1FieldEe1ltchShift = 0
	RegisterEefer1FieldEe1ltchMask  = 0x1
)

// GetEe1ltch External Event 1 latch
func (r *registerEefer1Type) GetEe1ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefer1FieldEe1ltchMask) != 0
}

// SetEe1ltch External Event 1 latch
func (r *registerEefer1Type) SetEe1ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefer1FieldEe1ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefer1FieldEe1ltchMask)
	}
}

// registerEefer2Type Timerx External Event Filtering Register 2
type registerEefer2Type uint32

const (
	RegisterEefer2FieldEe10fltrShift = 25
	RegisterEefer2FieldEe10fltrMask  = 0x1e000000
)

// GetEe10fltr External Event 10 filter
func (r *registerEefer2Type) GetEe10fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefer2FieldEe10fltrMask) >> RegisterEefer2FieldEe10fltrShift)
}

// SetEe10fltr External Event 10 filter
func (r *registerEefer2Type) SetEe10fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefer2FieldEe10fltrMask)|(uint32(value)<<RegisterEefer2FieldEe10fltrShift))
}

const (
	RegisterEefer2FieldEe10ltchShift = 24
	RegisterEefer2FieldEe10ltchMask  = 0x1000000
)

// GetEe10ltch External Event 10 latch
func (r *registerEefer2Type) GetEe10ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefer2FieldEe10ltchMask) != 0
}

// SetEe10ltch External Event 10 latch
func (r *registerEefer2Type) SetEe10ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefer2FieldEe10ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefer2FieldEe10ltchMask)
	}
}

const (
	RegisterEefer2FieldEe9fltrShift = 19
	RegisterEefer2FieldEe9fltrMask  = 0x780000
)

// GetEe9fltr External Event 9 filter
func (r *registerEefer2Type) GetEe9fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefer2FieldEe9fltrMask) >> RegisterEefer2FieldEe9fltrShift)
}

// SetEe9fltr External Event 9 filter
func (r *registerEefer2Type) SetEe9fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefer2FieldEe9fltrMask)|(uint32(value)<<RegisterEefer2FieldEe9fltrShift))
}

const (
	RegisterEefer2FieldEe9ltchShift = 18
	RegisterEefer2FieldEe9ltchMask  = 0x40000
)

// GetEe9ltch External Event 9 latch
func (r *registerEefer2Type) GetEe9ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefer2FieldEe9ltchMask) != 0
}

// SetEe9ltch External Event 9 latch
func (r *registerEefer2Type) SetEe9ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefer2FieldEe9ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefer2FieldEe9ltchMask)
	}
}

const (
	RegisterEefer2FieldEe8fltrShift = 13
	RegisterEefer2FieldEe8fltrMask  = 0x1e000
)

// GetEe8fltr External Event 8 filter
func (r *registerEefer2Type) GetEe8fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefer2FieldEe8fltrMask) >> RegisterEefer2FieldEe8fltrShift)
}

// SetEe8fltr External Event 8 filter
func (r *registerEefer2Type) SetEe8fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefer2FieldEe8fltrMask)|(uint32(value)<<RegisterEefer2FieldEe8fltrShift))
}

const (
	RegisterEefer2FieldEe8ltchShift = 12
	RegisterEefer2FieldEe8ltchMask  = 0x1000
)

// GetEe8ltch External Event 8 latch
func (r *registerEefer2Type) GetEe8ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefer2FieldEe8ltchMask) != 0
}

// SetEe8ltch External Event 8 latch
func (r *registerEefer2Type) SetEe8ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefer2FieldEe8ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefer2FieldEe8ltchMask)
	}
}

const (
	RegisterEefer2FieldEe7fltrShift = 7
	RegisterEefer2FieldEe7fltrMask  = 0x780
)

// GetEe7fltr External Event 7 filter
func (r *registerEefer2Type) GetEe7fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefer2FieldEe7fltrMask) >> RegisterEefer2FieldEe7fltrShift)
}

// SetEe7fltr External Event 7 filter
func (r *registerEefer2Type) SetEe7fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefer2FieldEe7fltrMask)|(uint32(value)<<RegisterEefer2FieldEe7fltrShift))
}

const (
	RegisterEefer2FieldEe7ltchShift = 6
	RegisterEefer2FieldEe7ltchMask  = 0x40
)

// GetEe7ltch External Event 7 latch
func (r *registerEefer2Type) GetEe7ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefer2FieldEe7ltchMask) != 0
}

// SetEe7ltch External Event 7 latch
func (r *registerEefer2Type) SetEe7ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefer2FieldEe7ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefer2FieldEe7ltchMask)
	}
}

const (
	RegisterEefer2FieldEe6fltrShift = 1
	RegisterEefer2FieldEe6fltrMask  = 0x1e
)

// GetEe6fltr External Event 6 filter
func (r *registerEefer2Type) GetEe6fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefer2FieldEe6fltrMask) >> RegisterEefer2FieldEe6fltrShift)
}

// SetEe6fltr External Event 6 filter
func (r *registerEefer2Type) SetEe6fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefer2FieldEe6fltrMask)|(uint32(value)<<RegisterEefer2FieldEe6fltrShift))
}

const (
	RegisterEefer2FieldEe6ltchShift = 0
	RegisterEefer2FieldEe6ltchMask  = 0x1
)

// GetEe6ltch External Event 6 latch
func (r *registerEefer2Type) GetEe6ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefer2FieldEe6ltchMask) != 0
}

// SetEe6ltch External Event 6 latch
func (r *registerEefer2Type) SetEe6ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefer2FieldEe6ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefer2FieldEe6ltchMask)
	}
}

// registerRsterType TimerA Reset Register
type registerRsterType uint32

const (
	RegisterRsterFieldTimdcmp4Shift = 30
	RegisterRsterFieldTimdcmp4Mask  = 0x40000000
)

// GetTimdcmp4 Timer D Compare 4
func (r *registerRsterType) GetTimdcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldTimdcmp4Mask) != 0
}

// SetTimdcmp4 Timer D Compare 4
func (r *registerRsterType) SetTimdcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldTimdcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldTimdcmp4Mask)
	}
}

const (
	RegisterRsterFieldTimdcmp2Shift = 29
	RegisterRsterFieldTimdcmp2Mask  = 0x20000000
)

// GetTimdcmp2 Timer D Compare 2
func (r *registerRsterType) GetTimdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldTimdcmp2Mask) != 0
}

// SetTimdcmp2 Timer D Compare 2
func (r *registerRsterType) SetTimdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldTimdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldTimdcmp2Mask)
	}
}

const (
	RegisterRsterFieldTimdcmp1Shift = 28
	RegisterRsterFieldTimdcmp1Mask  = 0x10000000
)

// GetTimdcmp1 Timer D Compare 1
func (r *registerRsterType) GetTimdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldTimdcmp1Mask) != 0
}

// SetTimdcmp1 Timer D Compare 1
func (r *registerRsterType) SetTimdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldTimdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldTimdcmp1Mask)
	}
}

const (
	RegisterRsterFieldTimccmp4Shift = 27
	RegisterRsterFieldTimccmp4Mask  = 0x8000000
)

// GetTimccmp4 Timer C Compare 4
func (r *registerRsterType) GetTimccmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldTimccmp4Mask) != 0
}

// SetTimccmp4 Timer C Compare 4
func (r *registerRsterType) SetTimccmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldTimccmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldTimccmp4Mask)
	}
}

const (
	RegisterRsterFieldTimccmp2Shift = 26
	RegisterRsterFieldTimccmp2Mask  = 0x4000000
)

// GetTimccmp2 Timer C Compare 2
func (r *registerRsterType) GetTimccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldTimccmp2Mask) != 0
}

// SetTimccmp2 Timer C Compare 2
func (r *registerRsterType) SetTimccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldTimccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldTimccmp2Mask)
	}
}

const (
	RegisterRsterFieldTimccmp1Shift = 25
	RegisterRsterFieldTimccmp1Mask  = 0x2000000
)

// GetTimccmp1 Timer C Compare 1
func (r *registerRsterType) GetTimccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldTimccmp1Mask) != 0
}

// SetTimccmp1 Timer C Compare 1
func (r *registerRsterType) SetTimccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldTimccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldTimccmp1Mask)
	}
}

const (
	RegisterRsterFieldTimbcmp4Shift = 24
	RegisterRsterFieldTimbcmp4Mask  = 0x1000000
)

// GetTimbcmp4 Timer B Compare 4
func (r *registerRsterType) GetTimbcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldTimbcmp4Mask) != 0
}

// SetTimbcmp4 Timer B Compare 4
func (r *registerRsterType) SetTimbcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldTimbcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldTimbcmp4Mask)
	}
}

const (
	RegisterRsterFieldTimbcmp2Shift = 23
	RegisterRsterFieldTimbcmp2Mask  = 0x800000
)

// GetTimbcmp2 Timer B Compare 2
func (r *registerRsterType) GetTimbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldTimbcmp2Mask) != 0
}

// SetTimbcmp2 Timer B Compare 2
func (r *registerRsterType) SetTimbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldTimbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldTimbcmp2Mask)
	}
}

const (
	RegisterRsterFieldTimbcmp1Shift = 22
	RegisterRsterFieldTimbcmp1Mask  = 0x400000
)

// GetTimbcmp1 Timer B Compare 1
func (r *registerRsterType) GetTimbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldTimbcmp1Mask) != 0
}

// SetTimbcmp1 Timer B Compare 1
func (r *registerRsterType) SetTimbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldTimbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldTimbcmp1Mask)
	}
}

const (
	RegisterRsterFieldTimacmp4Shift = 21
	RegisterRsterFieldTimacmp4Mask  = 0x200000
)

// GetTimacmp4 Timer A Compare 4
func (r *registerRsterType) GetTimacmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldTimacmp4Mask) != 0
}

// SetTimacmp4 Timer A Compare 4
func (r *registerRsterType) SetTimacmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldTimacmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldTimacmp4Mask)
	}
}

const (
	RegisterRsterFieldTimacmp2Shift = 20
	RegisterRsterFieldTimacmp2Mask  = 0x100000
)

// GetTimacmp2 Timer A Compare 2
func (r *registerRsterType) GetTimacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldTimacmp2Mask) != 0
}

// SetTimacmp2 Timer A Compare 2
func (r *registerRsterType) SetTimacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldTimacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldTimacmp2Mask)
	}
}

const (
	RegisterRsterFieldTimacmp1Shift = 19
	RegisterRsterFieldTimacmp1Mask  = 0x80000
)

// GetTimacmp1 Timer A Compare 1
func (r *registerRsterType) GetTimacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldTimacmp1Mask) != 0
}

// SetTimacmp1 Timer A Compare 1
func (r *registerRsterType) SetTimacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldTimacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldTimacmp1Mask)
	}
}

const (
	RegisterRsterFieldExtevnt10Shift = 18
	RegisterRsterFieldExtevnt10Mask  = 0x40000
)

// GetExtevnt10 External Event 10
func (r *registerRsterType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldExtevnt10Mask) != 0
}

// SetExtevnt10 External Event 10
func (r *registerRsterType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldExtevnt10Mask)
	}
}

const (
	RegisterRsterFieldExtevnt9Shift = 17
	RegisterRsterFieldExtevnt9Mask  = 0x20000
)

// GetExtevnt9 External Event 9
func (r *registerRsterType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldExtevnt9Mask) != 0
}

// SetExtevnt9 External Event 9
func (r *registerRsterType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldExtevnt9Mask)
	}
}

const (
	RegisterRsterFieldExtevnt8Shift = 16
	RegisterRsterFieldExtevnt8Mask  = 0x10000
)

// GetExtevnt8 External Event 8
func (r *registerRsterType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldExtevnt8Mask) != 0
}

// SetExtevnt8 External Event 8
func (r *registerRsterType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldExtevnt8Mask)
	}
}

const (
	RegisterRsterFieldExtevnt7Shift = 15
	RegisterRsterFieldExtevnt7Mask  = 0x8000
)

// GetExtevnt7 External Event 7
func (r *registerRsterType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldExtevnt7Mask) != 0
}

// SetExtevnt7 External Event 7
func (r *registerRsterType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldExtevnt7Mask)
	}
}

const (
	RegisterRsterFieldExtevnt6Shift = 14
	RegisterRsterFieldExtevnt6Mask  = 0x4000
)

// GetExtevnt6 External Event 6
func (r *registerRsterType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldExtevnt6Mask) != 0
}

// SetExtevnt6 External Event 6
func (r *registerRsterType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldExtevnt6Mask)
	}
}

const (
	RegisterRsterFieldExtevnt5Shift = 13
	RegisterRsterFieldExtevnt5Mask  = 0x2000
)

// GetExtevnt5 External Event 5
func (r *registerRsterType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldExtevnt5Mask) != 0
}

// SetExtevnt5 External Event 5
func (r *registerRsterType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldExtevnt5Mask)
	}
}

const (
	RegisterRsterFieldExtevnt4Shift = 12
	RegisterRsterFieldExtevnt4Mask  = 0x1000
)

// GetExtevnt4 External Event 4
func (r *registerRsterType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldExtevnt4Mask) != 0
}

// SetExtevnt4 External Event 4
func (r *registerRsterType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldExtevnt4Mask)
	}
}

const (
	RegisterRsterFieldExtevnt3Shift = 11
	RegisterRsterFieldExtevnt3Mask  = 0x800
)

// GetExtevnt3 External Event 3
func (r *registerRsterType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldExtevnt3Mask) != 0
}

// SetExtevnt3 External Event 3
func (r *registerRsterType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldExtevnt3Mask)
	}
}

const (
	RegisterRsterFieldExtevnt2Shift = 10
	RegisterRsterFieldExtevnt2Mask  = 0x400
)

// GetExtevnt2 External Event 2
func (r *registerRsterType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldExtevnt2Mask) != 0
}

// SetExtevnt2 External Event 2
func (r *registerRsterType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldExtevnt2Mask)
	}
}

const (
	RegisterRsterFieldExtevnt1Shift = 9
	RegisterRsterFieldExtevnt1Mask  = 0x200
)

// GetExtevnt1 External Event 1
func (r *registerRsterType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldExtevnt1Mask) != 0
}

// SetExtevnt1 External Event 1
func (r *registerRsterType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldExtevnt1Mask)
	}
}

const (
	RegisterRsterFieldMstcmp4Shift = 8
	RegisterRsterFieldMstcmp4Mask  = 0x100
)

// GetMstcmp4 Master compare 4
func (r *registerRsterType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldMstcmp4Mask) != 0
}

// SetMstcmp4 Master compare 4
func (r *registerRsterType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldMstcmp4Mask)
	}
}

const (
	RegisterRsterFieldMstcmp3Shift = 7
	RegisterRsterFieldMstcmp3Mask  = 0x80
)

// GetMstcmp3 Master compare 3
func (r *registerRsterType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldMstcmp3Mask) != 0
}

// SetMstcmp3 Master compare 3
func (r *registerRsterType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldMstcmp3Mask)
	}
}

const (
	RegisterRsterFieldMstcmp2Shift = 6
	RegisterRsterFieldMstcmp2Mask  = 0x40
)

// GetMstcmp2 Master compare 2
func (r *registerRsterType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldMstcmp2Mask) != 0
}

// SetMstcmp2 Master compare 2
func (r *registerRsterType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldMstcmp2Mask)
	}
}

const (
	RegisterRsterFieldMstcmp1Shift = 5
	RegisterRsterFieldMstcmp1Mask  = 0x20
)

// GetMstcmp1 Master compare 1
func (r *registerRsterType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldMstcmp1Mask) != 0
}

// SetMstcmp1 Master compare 1
func (r *registerRsterType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldMstcmp1Mask)
	}
}

const (
	RegisterRsterFieldMstperShift = 4
	RegisterRsterFieldMstperMask  = 0x10
)

// GetMstper Master timer Period
func (r *registerRsterType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldMstperMask) != 0
}

// SetMstper Master timer Period
func (r *registerRsterType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldMstperMask)
	}
}

const (
	RegisterRsterFieldCmp4Shift = 3
	RegisterRsterFieldCmp4Mask  = 0x8
)

// GetCmp4 Timer A compare 4 reset
func (r *registerRsterType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldCmp4Mask) != 0
}

// SetCmp4 Timer A compare 4 reset
func (r *registerRsterType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldCmp4Mask)
	}
}

const (
	RegisterRsterFieldCmp2Shift = 2
	RegisterRsterFieldCmp2Mask  = 0x4
)

// GetCmp2 Timer A compare 2 reset
func (r *registerRsterType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldCmp2Mask) != 0
}

// SetCmp2 Timer A compare 2 reset
func (r *registerRsterType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldCmp2Mask)
	}
}

const (
	RegisterRsterFieldUpdtShift = 1
	RegisterRsterFieldUpdtMask  = 0x2
)

// GetUpdt Timer A Update reset
func (r *registerRsterType) GetUpdt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsterFieldUpdtMask) != 0
}

// SetUpdt Timer A Update reset
func (r *registerRsterType) SetUpdt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsterFieldUpdtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsterFieldUpdtMask)
	}
}

// registerChperType Timerx Chopper Register
type registerChperType uint32

const (
	RegisterChperFieldStrtpwShift = 7
	RegisterChperFieldStrtpwMask  = 0x780
)

// GetStrtpw STRTPW
func (r *registerChperType) GetStrtpw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChperFieldStrtpwMask) >> RegisterChperFieldStrtpwShift)
}

// SetStrtpw STRTPW
func (r *registerChperType) SetStrtpw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChperFieldStrtpwMask)|(uint32(value)<<RegisterChperFieldStrtpwShift))
}

const (
	RegisterChperFieldChpdtyShift = 4
	RegisterChperFieldChpdtyMask  = 0x70
)

// GetChpdty Timerx chopper duty cycle value
func (r *registerChperType) GetChpdty() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChperFieldChpdtyMask) >> RegisterChperFieldChpdtyShift)
}

// SetChpdty Timerx chopper duty cycle value
func (r *registerChperType) SetChpdty(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChperFieldChpdtyMask)|(uint32(value)<<RegisterChperFieldChpdtyShift))
}

const (
	RegisterChperFieldChpfrqShift = 0
	RegisterChperFieldChpfrqMask  = 0xf
)

// GetChpfrq Timerx carrier frequency value
func (r *registerChperType) GetChpfrq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChperFieldChpfrqMask) >> RegisterChperFieldChpfrqShift)
}

// SetChpfrq Timerx carrier frequency value
func (r *registerChperType) SetChpfrq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChperFieldChpfrqMask)|(uint32(value)<<RegisterChperFieldChpfrqShift))
}

// registerCpt1ecrType Timerx Capture 2 Control Register
type registerCpt1ecrType uint32

const (
	RegisterCpt1ecrFieldTdcmp2Shift = 27
	RegisterCpt1ecrFieldTdcmp2Mask  = 0x8000000
)

// GetTdcmp2 Timer D Compare 2
func (r *registerCpt1ecrType) GetTdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTdcmp2Mask) != 0
}

// SetTdcmp2 Timer D Compare 2
func (r *registerCpt1ecrType) SetTdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTdcmp2Mask)
	}
}

const (
	RegisterCpt1ecrFieldTdcmp1Shift = 26
	RegisterCpt1ecrFieldTdcmp1Mask  = 0x4000000
)

// GetTdcmp1 Timer D Compare 1
func (r *registerCpt1ecrType) GetTdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTdcmp1Mask) != 0
}

// SetTdcmp1 Timer D Compare 1
func (r *registerCpt1ecrType) SetTdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTdcmp1Mask)
	}
}

const (
	RegisterCpt1ecrFieldTd1rstShift = 25
	RegisterCpt1ecrFieldTd1rstMask  = 0x2000000
)

// GetTd1rst Timer D output 1 Reset
func (r *registerCpt1ecrType) GetTd1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTd1rstMask) != 0
}

// SetTd1rst Timer D output 1 Reset
func (r *registerCpt1ecrType) SetTd1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTd1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTd1rstMask)
	}
}

const (
	RegisterCpt1ecrFieldTd1setShift = 24
	RegisterCpt1ecrFieldTd1setMask  = 0x1000000
)

// GetTd1set Timer D output 1 Set
func (r *registerCpt1ecrType) GetTd1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTd1setMask) != 0
}

// SetTd1set Timer D output 1 Set
func (r *registerCpt1ecrType) SetTd1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTd1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTd1setMask)
	}
}

const (
	RegisterCpt1ecrFieldTccmp2Shift = 23
	RegisterCpt1ecrFieldTccmp2Mask  = 0x800000
)

// GetTccmp2 Timer C Compare 2
func (r *registerCpt1ecrType) GetTccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTccmp2Mask) != 0
}

// SetTccmp2 Timer C Compare 2
func (r *registerCpt1ecrType) SetTccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTccmp2Mask)
	}
}

const (
	RegisterCpt1ecrFieldTccmp1Shift = 22
	RegisterCpt1ecrFieldTccmp1Mask  = 0x400000
)

// GetTccmp1 Timer C Compare 1
func (r *registerCpt1ecrType) GetTccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTccmp1Mask) != 0
}

// SetTccmp1 Timer C Compare 1
func (r *registerCpt1ecrType) SetTccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTccmp1Mask)
	}
}

const (
	RegisterCpt1ecrFieldTc1rstShift = 21
	RegisterCpt1ecrFieldTc1rstMask  = 0x200000
)

// GetTc1rst Timer C output 1 Reset
func (r *registerCpt1ecrType) GetTc1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTc1rstMask) != 0
}

// SetTc1rst Timer C output 1 Reset
func (r *registerCpt1ecrType) SetTc1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTc1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTc1rstMask)
	}
}

const (
	RegisterCpt1ecrFieldTc1setShift = 20
	RegisterCpt1ecrFieldTc1setMask  = 0x100000
)

// GetTc1set Timer C output 1 Set
func (r *registerCpt1ecrType) GetTc1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTc1setMask) != 0
}

// SetTc1set Timer C output 1 Set
func (r *registerCpt1ecrType) SetTc1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTc1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTc1setMask)
	}
}

const (
	RegisterCpt1ecrFieldTbcmp2Shift = 19
	RegisterCpt1ecrFieldTbcmp2Mask  = 0x80000
)

// GetTbcmp2 Timer B Compare 2
func (r *registerCpt1ecrType) GetTbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTbcmp2Mask) != 0
}

// SetTbcmp2 Timer B Compare 2
func (r *registerCpt1ecrType) SetTbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTbcmp2Mask)
	}
}

const (
	RegisterCpt1ecrFieldTbcmp1Shift = 18
	RegisterCpt1ecrFieldTbcmp1Mask  = 0x40000
)

// GetTbcmp1 Timer B Compare 1
func (r *registerCpt1ecrType) GetTbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTbcmp1Mask) != 0
}

// SetTbcmp1 Timer B Compare 1
func (r *registerCpt1ecrType) SetTbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTbcmp1Mask)
	}
}

const (
	RegisterCpt1ecrFieldTb1rstShift = 17
	RegisterCpt1ecrFieldTb1rstMask  = 0x20000
)

// GetTb1rst Timer B output 1 Reset
func (r *registerCpt1ecrType) GetTb1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTb1rstMask) != 0
}

// SetTb1rst Timer B output 1 Reset
func (r *registerCpt1ecrType) SetTb1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTb1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTb1rstMask)
	}
}

const (
	RegisterCpt1ecrFieldTb1setShift = 16
	RegisterCpt1ecrFieldTb1setMask  = 0x10000
)

// GetTb1set Timer B output 1 Set
func (r *registerCpt1ecrType) GetTb1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTb1setMask) != 0
}

// SetTb1set Timer B output 1 Set
func (r *registerCpt1ecrType) SetTb1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTb1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTb1setMask)
	}
}

const (
	RegisterCpt1ecrFieldTacmp2Shift = 15
	RegisterCpt1ecrFieldTacmp2Mask  = 0x8000
)

// GetTacmp2 Timer A Compare 2
func (r *registerCpt1ecrType) GetTacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTacmp2Mask) != 0
}

// SetTacmp2 Timer A Compare 2
func (r *registerCpt1ecrType) SetTacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTacmp2Mask)
	}
}

const (
	RegisterCpt1ecrFieldTacmp1Shift = 14
	RegisterCpt1ecrFieldTacmp1Mask  = 0x4000
)

// GetTacmp1 Timer A Compare 1
func (r *registerCpt1ecrType) GetTacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTacmp1Mask) != 0
}

// SetTacmp1 Timer A Compare 1
func (r *registerCpt1ecrType) SetTacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTacmp1Mask)
	}
}

const (
	RegisterCpt1ecrFieldTa1rstShift = 13
	RegisterCpt1ecrFieldTa1rstMask  = 0x2000
)

// GetTa1rst Timer A output 1 Reset
func (r *registerCpt1ecrType) GetTa1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTa1rstMask) != 0
}

// SetTa1rst Timer A output 1 Reset
func (r *registerCpt1ecrType) SetTa1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTa1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTa1rstMask)
	}
}

const (
	RegisterCpt1ecrFieldTa1setShift = 12
	RegisterCpt1ecrFieldTa1setMask  = 0x1000
)

// GetTa1set Timer A output 1 Set
func (r *registerCpt1ecrType) GetTa1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldTa1setMask) != 0
}

// SetTa1set Timer A output 1 Set
func (r *registerCpt1ecrType) SetTa1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldTa1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldTa1setMask)
	}
}

const (
	RegisterCpt1ecrFieldExev10cptShift = 11
	RegisterCpt1ecrFieldExev10cptMask  = 0x800
)

// GetExev10cpt External Event 10 Capture
func (r *registerCpt1ecrType) GetExev10cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldExev10cptMask) != 0
}

// SetExev10cpt External Event 10 Capture
func (r *registerCpt1ecrType) SetExev10cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldExev10cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldExev10cptMask)
	}
}

const (
	RegisterCpt1ecrFieldExev9cptShift = 10
	RegisterCpt1ecrFieldExev9cptMask  = 0x400
)

// GetExev9cpt External Event 9 Capture
func (r *registerCpt1ecrType) GetExev9cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldExev9cptMask) != 0
}

// SetExev9cpt External Event 9 Capture
func (r *registerCpt1ecrType) SetExev9cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldExev9cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldExev9cptMask)
	}
}

const (
	RegisterCpt1ecrFieldExev8cptShift = 9
	RegisterCpt1ecrFieldExev8cptMask  = 0x200
)

// GetExev8cpt External Event 8 Capture
func (r *registerCpt1ecrType) GetExev8cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldExev8cptMask) != 0
}

// SetExev8cpt External Event 8 Capture
func (r *registerCpt1ecrType) SetExev8cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldExev8cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldExev8cptMask)
	}
}

const (
	RegisterCpt1ecrFieldExev7cptShift = 8
	RegisterCpt1ecrFieldExev7cptMask  = 0x100
)

// GetExev7cpt External Event 7 Capture
func (r *registerCpt1ecrType) GetExev7cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldExev7cptMask) != 0
}

// SetExev7cpt External Event 7 Capture
func (r *registerCpt1ecrType) SetExev7cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldExev7cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldExev7cptMask)
	}
}

const (
	RegisterCpt1ecrFieldExev6cptShift = 7
	RegisterCpt1ecrFieldExev6cptMask  = 0x80
)

// GetExev6cpt External Event 6 Capture
func (r *registerCpt1ecrType) GetExev6cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldExev6cptMask) != 0
}

// SetExev6cpt External Event 6 Capture
func (r *registerCpt1ecrType) SetExev6cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldExev6cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldExev6cptMask)
	}
}

const (
	RegisterCpt1ecrFieldExev5cptShift = 6
	RegisterCpt1ecrFieldExev5cptMask  = 0x40
)

// GetExev5cpt External Event 5 Capture
func (r *registerCpt1ecrType) GetExev5cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldExev5cptMask) != 0
}

// SetExev5cpt External Event 5 Capture
func (r *registerCpt1ecrType) SetExev5cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldExev5cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldExev5cptMask)
	}
}

const (
	RegisterCpt1ecrFieldExev4cptShift = 5
	RegisterCpt1ecrFieldExev4cptMask  = 0x20
)

// GetExev4cpt External Event 4 Capture
func (r *registerCpt1ecrType) GetExev4cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldExev4cptMask) != 0
}

// SetExev4cpt External Event 4 Capture
func (r *registerCpt1ecrType) SetExev4cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldExev4cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldExev4cptMask)
	}
}

const (
	RegisterCpt1ecrFieldExev3cptShift = 4
	RegisterCpt1ecrFieldExev3cptMask  = 0x10
)

// GetExev3cpt External Event 3 Capture
func (r *registerCpt1ecrType) GetExev3cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldExev3cptMask) != 0
}

// SetExev3cpt External Event 3 Capture
func (r *registerCpt1ecrType) SetExev3cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldExev3cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldExev3cptMask)
	}
}

const (
	RegisterCpt1ecrFieldExev2cptShift = 3
	RegisterCpt1ecrFieldExev2cptMask  = 0x8
)

// GetExev2cpt External Event 2 Capture
func (r *registerCpt1ecrType) GetExev2cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldExev2cptMask) != 0
}

// SetExev2cpt External Event 2 Capture
func (r *registerCpt1ecrType) SetExev2cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldExev2cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldExev2cptMask)
	}
}

const (
	RegisterCpt1ecrFieldExev1cptShift = 2
	RegisterCpt1ecrFieldExev1cptMask  = 0x4
)

// GetExev1cpt External Event 1 Capture
func (r *registerCpt1ecrType) GetExev1cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldExev1cptMask) != 0
}

// SetExev1cpt External Event 1 Capture
func (r *registerCpt1ecrType) SetExev1cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldExev1cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldExev1cptMask)
	}
}

const (
	RegisterCpt1ecrFieldUdpcptShift = 1
	RegisterCpt1ecrFieldUdpcptMask  = 0x2
)

// GetUdpcpt Update Capture
func (r *registerCpt1ecrType) GetUdpcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldUdpcptMask) != 0
}

// SetUdpcpt Update Capture
func (r *registerCpt1ecrType) SetUdpcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldUdpcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldUdpcptMask)
	}
}

const (
	RegisterCpt1ecrFieldSwcptShift = 0
	RegisterCpt1ecrFieldSwcptMask  = 0x1
)

// GetSwcpt Software Capture
func (r *registerCpt1ecrType) GetSwcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ecrFieldSwcptMask) != 0
}

// SetSwcpt Software Capture
func (r *registerCpt1ecrType) SetSwcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ecrFieldSwcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ecrFieldSwcptMask)
	}
}

// registerCpt2ecrType CPT2xCR
type registerCpt2ecrType uint32

const (
	RegisterCpt2ecrFieldTdcmp2Shift = 27
	RegisterCpt2ecrFieldTdcmp2Mask  = 0x8000000
)

// GetTdcmp2 Timer D Compare 2
func (r *registerCpt2ecrType) GetTdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTdcmp2Mask) != 0
}

// SetTdcmp2 Timer D Compare 2
func (r *registerCpt2ecrType) SetTdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTdcmp2Mask)
	}
}

const (
	RegisterCpt2ecrFieldTdcmp1Shift = 26
	RegisterCpt2ecrFieldTdcmp1Mask  = 0x4000000
)

// GetTdcmp1 Timer D Compare 1
func (r *registerCpt2ecrType) GetTdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTdcmp1Mask) != 0
}

// SetTdcmp1 Timer D Compare 1
func (r *registerCpt2ecrType) SetTdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTdcmp1Mask)
	}
}

const (
	RegisterCpt2ecrFieldTd1rstShift = 25
	RegisterCpt2ecrFieldTd1rstMask  = 0x2000000
)

// GetTd1rst Timer D output 1 Reset
func (r *registerCpt2ecrType) GetTd1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTd1rstMask) != 0
}

// SetTd1rst Timer D output 1 Reset
func (r *registerCpt2ecrType) SetTd1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTd1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTd1rstMask)
	}
}

const (
	RegisterCpt2ecrFieldTd1setShift = 24
	RegisterCpt2ecrFieldTd1setMask  = 0x1000000
)

// GetTd1set Timer D output 1 Set
func (r *registerCpt2ecrType) GetTd1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTd1setMask) != 0
}

// SetTd1set Timer D output 1 Set
func (r *registerCpt2ecrType) SetTd1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTd1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTd1setMask)
	}
}

const (
	RegisterCpt2ecrFieldTccmp2Shift = 23
	RegisterCpt2ecrFieldTccmp2Mask  = 0x800000
)

// GetTccmp2 Timer C Compare 2
func (r *registerCpt2ecrType) GetTccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTccmp2Mask) != 0
}

// SetTccmp2 Timer C Compare 2
func (r *registerCpt2ecrType) SetTccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTccmp2Mask)
	}
}

const (
	RegisterCpt2ecrFieldTccmp1Shift = 22
	RegisterCpt2ecrFieldTccmp1Mask  = 0x400000
)

// GetTccmp1 Timer C Compare 1
func (r *registerCpt2ecrType) GetTccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTccmp1Mask) != 0
}

// SetTccmp1 Timer C Compare 1
func (r *registerCpt2ecrType) SetTccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTccmp1Mask)
	}
}

const (
	RegisterCpt2ecrFieldTc1rstShift = 21
	RegisterCpt2ecrFieldTc1rstMask  = 0x200000
)

// GetTc1rst Timer C output 1 Reset
func (r *registerCpt2ecrType) GetTc1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTc1rstMask) != 0
}

// SetTc1rst Timer C output 1 Reset
func (r *registerCpt2ecrType) SetTc1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTc1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTc1rstMask)
	}
}

const (
	RegisterCpt2ecrFieldTc1setShift = 20
	RegisterCpt2ecrFieldTc1setMask  = 0x100000
)

// GetTc1set Timer C output 1 Set
func (r *registerCpt2ecrType) GetTc1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTc1setMask) != 0
}

// SetTc1set Timer C output 1 Set
func (r *registerCpt2ecrType) SetTc1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTc1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTc1setMask)
	}
}

const (
	RegisterCpt2ecrFieldTbcmp2Shift = 19
	RegisterCpt2ecrFieldTbcmp2Mask  = 0x80000
)

// GetTbcmp2 Timer B Compare 2
func (r *registerCpt2ecrType) GetTbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTbcmp2Mask) != 0
}

// SetTbcmp2 Timer B Compare 2
func (r *registerCpt2ecrType) SetTbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTbcmp2Mask)
	}
}

const (
	RegisterCpt2ecrFieldTbcmp1Shift = 18
	RegisterCpt2ecrFieldTbcmp1Mask  = 0x40000
)

// GetTbcmp1 Timer B Compare 1
func (r *registerCpt2ecrType) GetTbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTbcmp1Mask) != 0
}

// SetTbcmp1 Timer B Compare 1
func (r *registerCpt2ecrType) SetTbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTbcmp1Mask)
	}
}

const (
	RegisterCpt2ecrFieldTb1rstShift = 17
	RegisterCpt2ecrFieldTb1rstMask  = 0x20000
)

// GetTb1rst Timer B output 1 Reset
func (r *registerCpt2ecrType) GetTb1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTb1rstMask) != 0
}

// SetTb1rst Timer B output 1 Reset
func (r *registerCpt2ecrType) SetTb1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTb1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTb1rstMask)
	}
}

const (
	RegisterCpt2ecrFieldTb1setShift = 16
	RegisterCpt2ecrFieldTb1setMask  = 0x10000
)

// GetTb1set Timer B output 1 Set
func (r *registerCpt2ecrType) GetTb1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTb1setMask) != 0
}

// SetTb1set Timer B output 1 Set
func (r *registerCpt2ecrType) SetTb1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTb1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTb1setMask)
	}
}

const (
	RegisterCpt2ecrFieldTacmp2Shift = 15
	RegisterCpt2ecrFieldTacmp2Mask  = 0x8000
)

// GetTacmp2 Timer A Compare 2
func (r *registerCpt2ecrType) GetTacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTacmp2Mask) != 0
}

// SetTacmp2 Timer A Compare 2
func (r *registerCpt2ecrType) SetTacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTacmp2Mask)
	}
}

const (
	RegisterCpt2ecrFieldTacmp1Shift = 14
	RegisterCpt2ecrFieldTacmp1Mask  = 0x4000
)

// GetTacmp1 Timer A Compare 1
func (r *registerCpt2ecrType) GetTacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTacmp1Mask) != 0
}

// SetTacmp1 Timer A Compare 1
func (r *registerCpt2ecrType) SetTacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTacmp1Mask)
	}
}

const (
	RegisterCpt2ecrFieldTa1rstShift = 13
	RegisterCpt2ecrFieldTa1rstMask  = 0x2000
)

// GetTa1rst Timer A output 1 Reset
func (r *registerCpt2ecrType) GetTa1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTa1rstMask) != 0
}

// SetTa1rst Timer A output 1 Reset
func (r *registerCpt2ecrType) SetTa1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTa1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTa1rstMask)
	}
}

const (
	RegisterCpt2ecrFieldTa1setShift = 12
	RegisterCpt2ecrFieldTa1setMask  = 0x1000
)

// GetTa1set Timer A output 1 Set
func (r *registerCpt2ecrType) GetTa1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldTa1setMask) != 0
}

// SetTa1set Timer A output 1 Set
func (r *registerCpt2ecrType) SetTa1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldTa1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldTa1setMask)
	}
}

const (
	RegisterCpt2ecrFieldExev10cptShift = 11
	RegisterCpt2ecrFieldExev10cptMask  = 0x800
)

// GetExev10cpt External Event 10 Capture
func (r *registerCpt2ecrType) GetExev10cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldExev10cptMask) != 0
}

// SetExev10cpt External Event 10 Capture
func (r *registerCpt2ecrType) SetExev10cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldExev10cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldExev10cptMask)
	}
}

const (
	RegisterCpt2ecrFieldExev9cptShift = 10
	RegisterCpt2ecrFieldExev9cptMask  = 0x400
)

// GetExev9cpt External Event 9 Capture
func (r *registerCpt2ecrType) GetExev9cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldExev9cptMask) != 0
}

// SetExev9cpt External Event 9 Capture
func (r *registerCpt2ecrType) SetExev9cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldExev9cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldExev9cptMask)
	}
}

const (
	RegisterCpt2ecrFieldExev8cptShift = 9
	RegisterCpt2ecrFieldExev8cptMask  = 0x200
)

// GetExev8cpt External Event 8 Capture
func (r *registerCpt2ecrType) GetExev8cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldExev8cptMask) != 0
}

// SetExev8cpt External Event 8 Capture
func (r *registerCpt2ecrType) SetExev8cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldExev8cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldExev8cptMask)
	}
}

const (
	RegisterCpt2ecrFieldExev7cptShift = 8
	RegisterCpt2ecrFieldExev7cptMask  = 0x100
)

// GetExev7cpt External Event 7 Capture
func (r *registerCpt2ecrType) GetExev7cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldExev7cptMask) != 0
}

// SetExev7cpt External Event 7 Capture
func (r *registerCpt2ecrType) SetExev7cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldExev7cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldExev7cptMask)
	}
}

const (
	RegisterCpt2ecrFieldExev6cptShift = 7
	RegisterCpt2ecrFieldExev6cptMask  = 0x80
)

// GetExev6cpt External Event 6 Capture
func (r *registerCpt2ecrType) GetExev6cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldExev6cptMask) != 0
}

// SetExev6cpt External Event 6 Capture
func (r *registerCpt2ecrType) SetExev6cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldExev6cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldExev6cptMask)
	}
}

const (
	RegisterCpt2ecrFieldExev5cptShift = 6
	RegisterCpt2ecrFieldExev5cptMask  = 0x40
)

// GetExev5cpt External Event 5 Capture
func (r *registerCpt2ecrType) GetExev5cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldExev5cptMask) != 0
}

// SetExev5cpt External Event 5 Capture
func (r *registerCpt2ecrType) SetExev5cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldExev5cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldExev5cptMask)
	}
}

const (
	RegisterCpt2ecrFieldExev4cptShift = 5
	RegisterCpt2ecrFieldExev4cptMask  = 0x20
)

// GetExev4cpt External Event 4 Capture
func (r *registerCpt2ecrType) GetExev4cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldExev4cptMask) != 0
}

// SetExev4cpt External Event 4 Capture
func (r *registerCpt2ecrType) SetExev4cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldExev4cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldExev4cptMask)
	}
}

const (
	RegisterCpt2ecrFieldExev3cptShift = 4
	RegisterCpt2ecrFieldExev3cptMask  = 0x10
)

// GetExev3cpt External Event 3 Capture
func (r *registerCpt2ecrType) GetExev3cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldExev3cptMask) != 0
}

// SetExev3cpt External Event 3 Capture
func (r *registerCpt2ecrType) SetExev3cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldExev3cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldExev3cptMask)
	}
}

const (
	RegisterCpt2ecrFieldExev2cptShift = 3
	RegisterCpt2ecrFieldExev2cptMask  = 0x8
)

// GetExev2cpt External Event 2 Capture
func (r *registerCpt2ecrType) GetExev2cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldExev2cptMask) != 0
}

// SetExev2cpt External Event 2 Capture
func (r *registerCpt2ecrType) SetExev2cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldExev2cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldExev2cptMask)
	}
}

const (
	RegisterCpt2ecrFieldExev1cptShift = 2
	RegisterCpt2ecrFieldExev1cptMask  = 0x4
)

// GetExev1cpt External Event 1 Capture
func (r *registerCpt2ecrType) GetExev1cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldExev1cptMask) != 0
}

// SetExev1cpt External Event 1 Capture
func (r *registerCpt2ecrType) SetExev1cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldExev1cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldExev1cptMask)
	}
}

const (
	RegisterCpt2ecrFieldUdpcptShift = 1
	RegisterCpt2ecrFieldUdpcptMask  = 0x2
)

// GetUdpcpt Update Capture
func (r *registerCpt2ecrType) GetUdpcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldUdpcptMask) != 0
}

// SetUdpcpt Update Capture
func (r *registerCpt2ecrType) SetUdpcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldUdpcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldUdpcptMask)
	}
}

const (
	RegisterCpt2ecrFieldSwcptShift = 0
	RegisterCpt2ecrFieldSwcptMask  = 0x1
)

// GetSwcpt Software Capture
func (r *registerCpt2ecrType) GetSwcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ecrFieldSwcptMask) != 0
}

// SetSwcpt Software Capture
func (r *registerCpt2ecrType) SetSwcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ecrFieldSwcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ecrFieldSwcptMask)
	}
}

// registerOuterType Timerx Output Register
type registerOuterType uint32

const (
	RegisterOuterFieldDidl2Shift = 23
	RegisterOuterFieldDidl2Mask  = 0x800000
)

// GetDidl2 Output 2 Deadtime upon burst mode Idle entry
func (r *registerOuterType) GetDidl2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldDidl2Mask) != 0
}

// SetDidl2 Output 2 Deadtime upon burst mode Idle entry
func (r *registerOuterType) SetDidl2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOuterFieldDidl2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldDidl2Mask)
	}
}

const (
	RegisterOuterFieldChp2Shift = 22
	RegisterOuterFieldChp2Mask  = 0x400000
)

// GetChp2 Output 2 Chopper enable
func (r *registerOuterType) GetChp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldChp2Mask) != 0
}

// SetChp2 Output 2 Chopper enable
func (r *registerOuterType) SetChp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOuterFieldChp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldChp2Mask)
	}
}

const (
	RegisterOuterFieldFault2Shift = 20
	RegisterOuterFieldFault2Mask  = 0x300000
)

// GetFault2 Output 2 Fault state
func (r *registerOuterType) GetFault2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldFault2Mask) >> RegisterOuterFieldFault2Shift)
}

// SetFault2 Output 2 Fault state
func (r *registerOuterType) SetFault2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldFault2Mask)|(uint32(value)<<RegisterOuterFieldFault2Shift))
}

const (
	RegisterOuterFieldIdles2Shift = 19
	RegisterOuterFieldIdles2Mask  = 0x80000
)

// GetIdles2 Output 2 Idle State
func (r *registerOuterType) GetIdles2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldIdles2Mask) != 0
}

// SetIdles2 Output 2 Idle State
func (r *registerOuterType) SetIdles2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOuterFieldIdles2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldIdles2Mask)
	}
}

const (
	RegisterOuterFieldIdlem2Shift = 18
	RegisterOuterFieldIdlem2Mask  = 0x40000
)

// GetIdlem2 Output 2 Idle mode
func (r *registerOuterType) GetIdlem2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldIdlem2Mask) != 0
}

// SetIdlem2 Output 2 Idle mode
func (r *registerOuterType) SetIdlem2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOuterFieldIdlem2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldIdlem2Mask)
	}
}

const (
	RegisterOuterFieldPol2Shift = 17
	RegisterOuterFieldPol2Mask  = 0x20000
)

// GetPol2 Output 2 polarity
func (r *registerOuterType) GetPol2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldPol2Mask) != 0
}

// SetPol2 Output 2 polarity
func (r *registerOuterType) SetPol2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOuterFieldPol2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldPol2Mask)
	}
}

const (
	RegisterOuterFieldDlyprtShift = 10
	RegisterOuterFieldDlyprtMask  = 0x1c00
)

// GetDlyprt Delayed Protection
func (r *registerOuterType) GetDlyprt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldDlyprtMask) >> RegisterOuterFieldDlyprtShift)
}

// SetDlyprt Delayed Protection
func (r *registerOuterType) SetDlyprt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldDlyprtMask)|(uint32(value)<<RegisterOuterFieldDlyprtShift))
}

const (
	RegisterOuterFieldDlyprtenShift = 9
	RegisterOuterFieldDlyprtenMask  = 0x200
)

// GetDlyprten Delayed Protection Enable
func (r *registerOuterType) GetDlyprten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldDlyprtenMask) != 0
}

// SetDlyprten Delayed Protection Enable
func (r *registerOuterType) SetDlyprten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOuterFieldDlyprtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldDlyprtenMask)
	}
}

const (
	RegisterOuterFieldDtenShift = 8
	RegisterOuterFieldDtenMask  = 0x100
)

// GetDten Deadtime enable
func (r *registerOuterType) GetDten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldDtenMask) != 0
}

// SetDten Deadtime enable
func (r *registerOuterType) SetDten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOuterFieldDtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldDtenMask)
	}
}

const (
	RegisterOuterFieldDidl1Shift = 7
	RegisterOuterFieldDidl1Mask  = 0x80
)

// GetDidl1 Output 1 Deadtime upon burst mode Idle entry
func (r *registerOuterType) GetDidl1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldDidl1Mask) != 0
}

// SetDidl1 Output 1 Deadtime upon burst mode Idle entry
func (r *registerOuterType) SetDidl1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOuterFieldDidl1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldDidl1Mask)
	}
}

const (
	RegisterOuterFieldChp1Shift = 6
	RegisterOuterFieldChp1Mask  = 0x40
)

// GetChp1 Output 1 Chopper enable
func (r *registerOuterType) GetChp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldChp1Mask) != 0
}

// SetChp1 Output 1 Chopper enable
func (r *registerOuterType) SetChp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOuterFieldChp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldChp1Mask)
	}
}

const (
	RegisterOuterFieldFault1Shift = 4
	RegisterOuterFieldFault1Mask  = 0x30
)

// GetFault1 Output 1 Fault state
func (r *registerOuterType) GetFault1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldFault1Mask) >> RegisterOuterFieldFault1Shift)
}

// SetFault1 Output 1 Fault state
func (r *registerOuterType) SetFault1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldFault1Mask)|(uint32(value)<<RegisterOuterFieldFault1Shift))
}

const (
	RegisterOuterFieldIdles1Shift = 3
	RegisterOuterFieldIdles1Mask  = 0x8
)

// GetIdles1 Output 1 Idle State
func (r *registerOuterType) GetIdles1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldIdles1Mask) != 0
}

// SetIdles1 Output 1 Idle State
func (r *registerOuterType) SetIdles1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOuterFieldIdles1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldIdles1Mask)
	}
}

const (
	RegisterOuterFieldIdlem1Shift = 2
	RegisterOuterFieldIdlem1Mask  = 0x4
)

// GetIdlem1 Output 1 Idle mode
func (r *registerOuterType) GetIdlem1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldIdlem1Mask) != 0
}

// SetIdlem1 Output 1 Idle mode
func (r *registerOuterType) SetIdlem1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOuterFieldIdlem1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldIdlem1Mask)
	}
}

const (
	RegisterOuterFieldPol1Shift = 1
	RegisterOuterFieldPol1Mask  = 0x2
)

// GetPol1 Output 1 polarity
func (r *registerOuterType) GetPol1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOuterFieldPol1Mask) != 0
}

// SetPol1 Output 1 polarity
func (r *registerOuterType) SetPol1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOuterFieldPol1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOuterFieldPol1Mask)
	}
}

// registerFlterType Timerx Fault Register
type registerFlterType uint32

const (
	RegisterFlterFieldFltlckShift = 31
	RegisterFlterFieldFltlckMask  = 0x80000000
)

// GetFltlck Fault sources Lock
func (r *registerFlterType) GetFltlck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFlterFieldFltlckMask) != 0
}

// SetFltlck Fault sources Lock
func (r *registerFlterType) SetFltlck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFlterFieldFltlckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFlterFieldFltlckMask)
	}
}

const (
	RegisterFlterFieldFlt5enShift = 4
	RegisterFlterFieldFlt5enMask  = 0x10
)

// GetFlt5en Fault 5 enable
func (r *registerFlterType) GetFlt5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFlterFieldFlt5enMask) != 0
}

// SetFlt5en Fault 5 enable
func (r *registerFlterType) SetFlt5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFlterFieldFlt5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFlterFieldFlt5enMask)
	}
}

const (
	RegisterFlterFieldFlt4enShift = 3
	RegisterFlterFieldFlt4enMask  = 0x8
)

// GetFlt4en Fault 4 enable
func (r *registerFlterType) GetFlt4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFlterFieldFlt4enMask) != 0
}

// SetFlt4en Fault 4 enable
func (r *registerFlterType) SetFlt4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFlterFieldFlt4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFlterFieldFlt4enMask)
	}
}

const (
	RegisterFlterFieldFlt3enShift = 2
	RegisterFlterFieldFlt3enMask  = 0x4
)

// GetFlt3en Fault 3 enable
func (r *registerFlterType) GetFlt3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFlterFieldFlt3enMask) != 0
}

// SetFlt3en Fault 3 enable
func (r *registerFlterType) SetFlt3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFlterFieldFlt3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFlterFieldFlt3enMask)
	}
}

const (
	RegisterFlterFieldFlt2enShift = 1
	RegisterFlterFieldFlt2enMask  = 0x2
)

// GetFlt2en Fault 2 enable
func (r *registerFlterType) GetFlt2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFlterFieldFlt2enMask) != 0
}

// SetFlt2en Fault 2 enable
func (r *registerFlterType) SetFlt2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFlterFieldFlt2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFlterFieldFlt2enMask)
	}
}

const (
	RegisterFlterFieldFlt1enShift = 0
	RegisterFlterFieldFlt1enMask  = 0x1
)

// GetFlt1en Fault 1 enable
func (r *registerFlterType) GetFlt1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFlterFieldFlt1enMask) != 0
}

// SetFlt1en Fault 1 enable
func (r *registerFlterType) SetFlt1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFlterFieldFlt1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFlterFieldFlt1enMask)
	}
}
