package hrtim_timd

import (
	"unsafe"
	"volatile"
)

var (
	Hrtim_timd = (*_hrtim_timd)(unsafe.Pointer(uintptr(0x40017600)))
)

type _hrtim_timd struct {
	Timdcr    registerTimdcrType
	Timdisr   registerTimdisrType
	Timdicr   registerTimdicrType
	Timddier5 registerTimddier5Type
	Cntdr     registerCntdrType
	Perdr     registerPerdrType
	Repdr     registerRepdrType
	Cmp1dr    registerCmp1drType
	Cmp1cdr   registerCmp1cdrType
	Cmp2dr    registerCmp2drType
	Cmp3dr    registerCmp3drType
	Cmp4dr    registerCmp4drType
	Cpt1dr    registerCpt1drType
	Cpt2dr    registerCpt2drType
	Dtdr      registerDtdrType
	Setd1r    registerSetd1rType
	Rstd1r    registerRstd1rType
	Setd2r    registerSetd2rType
	Rstd2r    registerRstd2rType
	Eefdr1    registerEefdr1Type
	Eefdr2    registerEefdr2Type
	Rstdr     registerRstdrType
	Chpdr     registerChpdrType
	Cpt1dcr   registerCpt1dcrType
	Cpt2dcr   registerCpt2dcrType
	Outdr     registerOutdrType
	Fltdr     registerFltdrType
}

// registerTimdcrType Timerx Control Register
type registerTimdcrType uint32

const (
	RegisterTimdcrFieldUpdgatShift = 28
	RegisterTimdcrFieldUpdgatMask  = 0xf0000000
)

// GetUpdgat Update Gating
func (r *registerTimdcrType) GetUpdgat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldUpdgatMask) >> RegisterTimdcrFieldUpdgatShift)
}

// SetUpdgat Update Gating
func (r *registerTimdcrType) SetUpdgat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldUpdgatMask)|(uint32(value)<<RegisterTimdcrFieldUpdgatShift))
}

const (
	RegisterTimdcrFieldPreenShift = 27
	RegisterTimdcrFieldPreenMask  = 0x8000000
)

// GetPreen Preload enable
func (r *registerTimdcrType) GetPreen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldPreenMask) != 0
}

// SetPreen Preload enable
func (r *registerTimdcrType) SetPreen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldPreenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldPreenMask)
	}
}

const (
	RegisterTimdcrFieldDacsyncShift = 25
	RegisterTimdcrFieldDacsyncMask  = 0x6000000
)

// GetDacsync AC Synchronization
func (r *registerTimdcrType) GetDacsync() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldDacsyncMask) >> RegisterTimdcrFieldDacsyncShift)
}

// SetDacsync AC Synchronization
func (r *registerTimdcrType) SetDacsync(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldDacsyncMask)|(uint32(value)<<RegisterTimdcrFieldDacsyncShift))
}

const (
	RegisterTimdcrFieldMstuShift = 24
	RegisterTimdcrFieldMstuMask  = 0x1000000
)

// GetMstu Master Timer update
func (r *registerTimdcrType) GetMstu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldMstuMask) != 0
}

// SetMstu Master Timer update
func (r *registerTimdcrType) SetMstu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldMstuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldMstuMask)
	}
}

const (
	RegisterTimdcrFieldTeuShift = 23
	RegisterTimdcrFieldTeuMask  = 0x800000
)

// GetTeu TEU
func (r *registerTimdcrType) GetTeu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldTeuMask) != 0
}

// SetTeu TEU
func (r *registerTimdcrType) SetTeu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldTeuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldTeuMask)
	}
}

const (
	RegisterTimdcrFieldTduShift = 22
	RegisterTimdcrFieldTduMask  = 0x400000
)

// GetTdu TDU
func (r *registerTimdcrType) GetTdu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldTduMask) != 0
}

// SetTdu TDU
func (r *registerTimdcrType) SetTdu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldTduMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldTduMask)
	}
}

const (
	RegisterTimdcrFieldTcuShift = 21
	RegisterTimdcrFieldTcuMask  = 0x200000
)

// GetTcu TCU
func (r *registerTimdcrType) GetTcu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldTcuMask) != 0
}

// SetTcu TCU
func (r *registerTimdcrType) SetTcu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldTcuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldTcuMask)
	}
}

const (
	RegisterTimdcrFieldTbuShift = 20
	RegisterTimdcrFieldTbuMask  = 0x100000
)

// GetTbu TBU
func (r *registerTimdcrType) GetTbu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldTbuMask) != 0
}

// SetTbu TBU
func (r *registerTimdcrType) SetTbu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldTbuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldTbuMask)
	}
}

const (
	RegisterTimdcrFieldTxrstuShift = 18
	RegisterTimdcrFieldTxrstuMask  = 0x40000
)

// GetTxrstu Timerx reset update
func (r *registerTimdcrType) GetTxrstu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldTxrstuMask) != 0
}

// SetTxrstu Timerx reset update
func (r *registerTimdcrType) SetTxrstu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldTxrstuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldTxrstuMask)
	}
}

const (
	RegisterTimdcrFieldTxrepuShift = 17
	RegisterTimdcrFieldTxrepuMask  = 0x20000
)

// GetTxrepu Timer x Repetition update
func (r *registerTimdcrType) GetTxrepu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldTxrepuMask) != 0
}

// SetTxrepu Timer x Repetition update
func (r *registerTimdcrType) SetTxrepu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldTxrepuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldTxrepuMask)
	}
}

const (
	RegisterTimdcrFieldDelcmp4Shift = 14
	RegisterTimdcrFieldDelcmp4Mask  = 0xc000
)

// GetDelcmp4 Delayed CMP4 mode
func (r *registerTimdcrType) GetDelcmp4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldDelcmp4Mask) >> RegisterTimdcrFieldDelcmp4Shift)
}

// SetDelcmp4 Delayed CMP4 mode
func (r *registerTimdcrType) SetDelcmp4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldDelcmp4Mask)|(uint32(value)<<RegisterTimdcrFieldDelcmp4Shift))
}

const (
	RegisterTimdcrFieldDelcmp2Shift = 12
	RegisterTimdcrFieldDelcmp2Mask  = 0x3000
)

// GetDelcmp2 Delayed CMP2 mode
func (r *registerTimdcrType) GetDelcmp2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldDelcmp2Mask) >> RegisterTimdcrFieldDelcmp2Shift)
}

// SetDelcmp2 Delayed CMP2 mode
func (r *registerTimdcrType) SetDelcmp2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldDelcmp2Mask)|(uint32(value)<<RegisterTimdcrFieldDelcmp2Shift))
}

const (
	RegisterTimdcrFieldSyncstrtxShift = 11
	RegisterTimdcrFieldSyncstrtxMask  = 0x800
)

// GetSyncstrtx Synchronization Starts Timer x
func (r *registerTimdcrType) GetSyncstrtx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldSyncstrtxMask) != 0
}

// SetSyncstrtx Synchronization Starts Timer x
func (r *registerTimdcrType) SetSyncstrtx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldSyncstrtxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldSyncstrtxMask)
	}
}

const (
	RegisterTimdcrFieldSyncrstxShift = 10
	RegisterTimdcrFieldSyncrstxMask  = 0x400
)

// GetSyncrstx Synchronization Resets Timer x
func (r *registerTimdcrType) GetSyncrstx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldSyncrstxMask) != 0
}

// SetSyncrstx Synchronization Resets Timer x
func (r *registerTimdcrType) SetSyncrstx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldSyncrstxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldSyncrstxMask)
	}
}

const (
	RegisterTimdcrFieldPshpllShift = 6
	RegisterTimdcrFieldPshpllMask  = 0x40
)

// GetPshpll Push-Pull mode enable
func (r *registerTimdcrType) GetPshpll() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldPshpllMask) != 0
}

// SetPshpll Push-Pull mode enable
func (r *registerTimdcrType) SetPshpll(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldPshpllMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldPshpllMask)
	}
}

const (
	RegisterTimdcrFieldHalfShift = 5
	RegisterTimdcrFieldHalfMask  = 0x20
)

// GetHalf Half mode enable
func (r *registerTimdcrType) GetHalf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldHalfMask) != 0
}

// SetHalf Half mode enable
func (r *registerTimdcrType) SetHalf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldHalfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldHalfMask)
	}
}

const (
	RegisterTimdcrFieldRetrigShift = 4
	RegisterTimdcrFieldRetrigMask  = 0x10
)

// GetRetrig Re-triggerable mode
func (r *registerTimdcrType) GetRetrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldRetrigMask) != 0
}

// SetRetrig Re-triggerable mode
func (r *registerTimdcrType) SetRetrig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldRetrigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldRetrigMask)
	}
}

const (
	RegisterTimdcrFieldContShift = 3
	RegisterTimdcrFieldContMask  = 0x8
)

// GetCont Continuous mode
func (r *registerTimdcrType) GetCont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldContMask) != 0
}

// SetCont Continuous mode
func (r *registerTimdcrType) SetCont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdcrFieldContMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldContMask)
	}
}

const (
	RegisterTimdcrFieldCk_pscxShift = 0
	RegisterTimdcrFieldCk_pscxMask  = 0x7
)

// GetCk_pscx HRTIM Timer x Clock prescaler
func (r *registerTimdcrType) GetCk_pscx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimdcrFieldCk_pscxMask) >> RegisterTimdcrFieldCk_pscxShift)
}

// SetCk_pscx HRTIM Timer x Clock prescaler
func (r *registerTimdcrType) SetCk_pscx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimdcrFieldCk_pscxMask)|(uint32(value)<<RegisterTimdcrFieldCk_pscxShift))
}

// registerTimdisrType Timerx Interrupt Status Register
type registerTimdisrType uint32

const (
	RegisterTimdisrFieldO2statShift = 19
	RegisterTimdisrFieldO2statMask  = 0x80000
)

// GetO2stat Output 2 State
func (r *registerTimdisrType) GetO2stat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldO2statMask) != 0
}

// SetO2stat Output 2 State
func (r *registerTimdisrType) SetO2stat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldO2statMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldO2statMask)
	}
}

const (
	RegisterTimdisrFieldO1statShift = 18
	RegisterTimdisrFieldO1statMask  = 0x40000
)

// GetO1stat Output 1 State
func (r *registerTimdisrType) GetO1stat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldO1statMask) != 0
}

// SetO1stat Output 1 State
func (r *registerTimdisrType) SetO1stat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldO1statMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldO1statMask)
	}
}

const (
	RegisterTimdisrFieldIppstatShift = 17
	RegisterTimdisrFieldIppstatMask  = 0x20000
)

// GetIppstat Idle Push Pull Status
func (r *registerTimdisrType) GetIppstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldIppstatMask) != 0
}

// SetIppstat Idle Push Pull Status
func (r *registerTimdisrType) SetIppstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldIppstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldIppstatMask)
	}
}

const (
	RegisterTimdisrFieldCppstatShift = 16
	RegisterTimdisrFieldCppstatMask  = 0x10000
)

// GetCppstat Current Push Pull Status
func (r *registerTimdisrType) GetCppstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldCppstatMask) != 0
}

// SetCppstat Current Push Pull Status
func (r *registerTimdisrType) SetCppstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldCppstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldCppstatMask)
	}
}

const (
	RegisterTimdisrFieldDlyprtShift = 14
	RegisterTimdisrFieldDlyprtMask  = 0x4000
)

// GetDlyprt Delayed Protection Flag
func (r *registerTimdisrType) GetDlyprt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldDlyprtMask) != 0
}

// SetDlyprt Delayed Protection Flag
func (r *registerTimdisrType) SetDlyprt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldDlyprtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldDlyprtMask)
	}
}

const (
	RegisterTimdisrFieldRstShift = 13
	RegisterTimdisrFieldRstMask  = 0x2000
)

// GetRst Reset Interrupt Flag
func (r *registerTimdisrType) GetRst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldRstMask) != 0
}

// SetRst Reset Interrupt Flag
func (r *registerTimdisrType) SetRst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldRstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldRstMask)
	}
}

const (
	RegisterTimdisrFieldRstx2Shift = 12
	RegisterTimdisrFieldRstx2Mask  = 0x1000
)

// GetRstx2 Output 2 Reset Interrupt Flag
func (r *registerTimdisrType) GetRstx2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldRstx2Mask) != 0
}

// SetRstx2 Output 2 Reset Interrupt Flag
func (r *registerTimdisrType) SetRstx2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldRstx2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldRstx2Mask)
	}
}

const (
	RegisterTimdisrFieldSetx2Shift = 11
	RegisterTimdisrFieldSetx2Mask  = 0x800
)

// GetSetx2 Output 2 Set Interrupt Flag
func (r *registerTimdisrType) GetSetx2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldSetx2Mask) != 0
}

// SetSetx2 Output 2 Set Interrupt Flag
func (r *registerTimdisrType) SetSetx2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldSetx2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldSetx2Mask)
	}
}

const (
	RegisterTimdisrFieldRstx1Shift = 10
	RegisterTimdisrFieldRstx1Mask  = 0x400
)

// GetRstx1 Output 1 Reset Interrupt Flag
func (r *registerTimdisrType) GetRstx1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldRstx1Mask) != 0
}

// SetRstx1 Output 1 Reset Interrupt Flag
func (r *registerTimdisrType) SetRstx1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldRstx1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldRstx1Mask)
	}
}

const (
	RegisterTimdisrFieldSetx1Shift = 9
	RegisterTimdisrFieldSetx1Mask  = 0x200
)

// GetSetx1 Output 1 Set Interrupt Flag
func (r *registerTimdisrType) GetSetx1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldSetx1Mask) != 0
}

// SetSetx1 Output 1 Set Interrupt Flag
func (r *registerTimdisrType) SetSetx1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldSetx1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldSetx1Mask)
	}
}

const (
	RegisterTimdisrFieldCpt2Shift = 8
	RegisterTimdisrFieldCpt2Mask  = 0x100
)

// GetCpt2 Capture2 Interrupt Flag
func (r *registerTimdisrType) GetCpt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldCpt2Mask) != 0
}

// SetCpt2 Capture2 Interrupt Flag
func (r *registerTimdisrType) SetCpt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldCpt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldCpt2Mask)
	}
}

const (
	RegisterTimdisrFieldCpt1Shift = 7
	RegisterTimdisrFieldCpt1Mask  = 0x80
)

// GetCpt1 Capture1 Interrupt Flag
func (r *registerTimdisrType) GetCpt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldCpt1Mask) != 0
}

// SetCpt1 Capture1 Interrupt Flag
func (r *registerTimdisrType) SetCpt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldCpt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldCpt1Mask)
	}
}

const (
	RegisterTimdisrFieldUpdShift = 6
	RegisterTimdisrFieldUpdMask  = 0x40
)

// GetUpd Update Interrupt Flag
func (r *registerTimdisrType) GetUpd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldUpdMask) != 0
}

// SetUpd Update Interrupt Flag
func (r *registerTimdisrType) SetUpd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldUpdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldUpdMask)
	}
}

const (
	RegisterTimdisrFieldRepShift = 4
	RegisterTimdisrFieldRepMask  = 0x10
)

// GetRep Repetition Interrupt Flag
func (r *registerTimdisrType) GetRep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldRepMask) != 0
}

// SetRep Repetition Interrupt Flag
func (r *registerTimdisrType) SetRep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldRepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldRepMask)
	}
}

const (
	RegisterTimdisrFieldCmp4Shift = 3
	RegisterTimdisrFieldCmp4Mask  = 0x8
)

// GetCmp4 Compare 4 Interrupt Flag
func (r *registerTimdisrType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldCmp4Mask) != 0
}

// SetCmp4 Compare 4 Interrupt Flag
func (r *registerTimdisrType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldCmp4Mask)
	}
}

const (
	RegisterTimdisrFieldCmp3Shift = 2
	RegisterTimdisrFieldCmp3Mask  = 0x4
)

// GetCmp3 Compare 3 Interrupt Flag
func (r *registerTimdisrType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldCmp3Mask) != 0
}

// SetCmp3 Compare 3 Interrupt Flag
func (r *registerTimdisrType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldCmp3Mask)
	}
}

const (
	RegisterTimdisrFieldCmp2Shift = 1
	RegisterTimdisrFieldCmp2Mask  = 0x2
)

// GetCmp2 Compare 2 Interrupt Flag
func (r *registerTimdisrType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldCmp2Mask) != 0
}

// SetCmp2 Compare 2 Interrupt Flag
func (r *registerTimdisrType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldCmp2Mask)
	}
}

const (
	RegisterTimdisrFieldCmp1Shift = 0
	RegisterTimdisrFieldCmp1Mask  = 0x1
)

// GetCmp1 Compare 1 Interrupt Flag
func (r *registerTimdisrType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdisrFieldCmp1Mask) != 0
}

// SetCmp1 Compare 1 Interrupt Flag
func (r *registerTimdisrType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdisrFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdisrFieldCmp1Mask)
	}
}

// registerTimdicrType Timerx Interrupt Clear Register
type registerTimdicrType uint32

const (
	RegisterTimdicrFieldDlyprtcShift = 14
	RegisterTimdicrFieldDlyprtcMask  = 0x4000
)

// GetDlyprtc Delayed Protection Flag Clear
func (r *registerTimdicrType) GetDlyprtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldDlyprtcMask) != 0
}

// SetDlyprtc Delayed Protection Flag Clear
func (r *registerTimdicrType) SetDlyprtc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldDlyprtcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldDlyprtcMask)
	}
}

const (
	RegisterTimdicrFieldRstcShift = 13
	RegisterTimdicrFieldRstcMask  = 0x2000
)

// GetRstc Reset Interrupt flag Clear
func (r *registerTimdicrType) GetRstc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldRstcMask) != 0
}

// SetRstc Reset Interrupt flag Clear
func (r *registerTimdicrType) SetRstc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldRstcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldRstcMask)
	}
}

const (
	RegisterTimdicrFieldRstx2cShift = 12
	RegisterTimdicrFieldRstx2cMask  = 0x1000
)

// GetRstx2c Output 2 Reset flag Clear
func (r *registerTimdicrType) GetRstx2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldRstx2cMask) != 0
}

// SetRstx2c Output 2 Reset flag Clear
func (r *registerTimdicrType) SetRstx2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldRstx2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldRstx2cMask)
	}
}

const (
	RegisterTimdicrFieldSet2xcShift = 11
	RegisterTimdicrFieldSet2xcMask  = 0x800
)

// GetSet2xc Output 2 Set flag Clear
func (r *registerTimdicrType) GetSet2xc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldSet2xcMask) != 0
}

// SetSet2xc Output 2 Set flag Clear
func (r *registerTimdicrType) SetSet2xc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldSet2xcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldSet2xcMask)
	}
}

const (
	RegisterTimdicrFieldRstx1cShift = 10
	RegisterTimdicrFieldRstx1cMask  = 0x400
)

// GetRstx1c Output 1 Reset flag Clear
func (r *registerTimdicrType) GetRstx1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldRstx1cMask) != 0
}

// SetRstx1c Output 1 Reset flag Clear
func (r *registerTimdicrType) SetRstx1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldRstx1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldRstx1cMask)
	}
}

const (
	RegisterTimdicrFieldSet1xcShift = 9
	RegisterTimdicrFieldSet1xcMask  = 0x200
)

// GetSet1xc Output 1 Set flag Clear
func (r *registerTimdicrType) GetSet1xc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldSet1xcMask) != 0
}

// SetSet1xc Output 1 Set flag Clear
func (r *registerTimdicrType) SetSet1xc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldSet1xcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldSet1xcMask)
	}
}

const (
	RegisterTimdicrFieldCpt2cShift = 8
	RegisterTimdicrFieldCpt2cMask  = 0x100
)

// GetCpt2c Capture2 Interrupt flag Clear
func (r *registerTimdicrType) GetCpt2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldCpt2cMask) != 0
}

// SetCpt2c Capture2 Interrupt flag Clear
func (r *registerTimdicrType) SetCpt2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldCpt2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldCpt2cMask)
	}
}

const (
	RegisterTimdicrFieldCpt1cShift = 7
	RegisterTimdicrFieldCpt1cMask  = 0x80
)

// GetCpt1c Capture1 Interrupt flag Clear
func (r *registerTimdicrType) GetCpt1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldCpt1cMask) != 0
}

// SetCpt1c Capture1 Interrupt flag Clear
func (r *registerTimdicrType) SetCpt1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldCpt1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldCpt1cMask)
	}
}

const (
	RegisterTimdicrFieldUpdcShift = 6
	RegisterTimdicrFieldUpdcMask  = 0x40
)

// GetUpdc Update Interrupt flag Clear
func (r *registerTimdicrType) GetUpdc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldUpdcMask) != 0
}

// SetUpdc Update Interrupt flag Clear
func (r *registerTimdicrType) SetUpdc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldUpdcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldUpdcMask)
	}
}

const (
	RegisterTimdicrFieldRepcShift = 4
	RegisterTimdicrFieldRepcMask  = 0x10
)

// GetRepc Repetition Interrupt flag Clear
func (r *registerTimdicrType) GetRepc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldRepcMask) != 0
}

// SetRepc Repetition Interrupt flag Clear
func (r *registerTimdicrType) SetRepc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldRepcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldRepcMask)
	}
}

const (
	RegisterTimdicrFieldCmp4cShift = 3
	RegisterTimdicrFieldCmp4cMask  = 0x8
)

// GetCmp4c Compare 4 Interrupt flag Clear
func (r *registerTimdicrType) GetCmp4c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldCmp4cMask) != 0
}

// SetCmp4c Compare 4 Interrupt flag Clear
func (r *registerTimdicrType) SetCmp4c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldCmp4cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldCmp4cMask)
	}
}

const (
	RegisterTimdicrFieldCmp3cShift = 2
	RegisterTimdicrFieldCmp3cMask  = 0x4
)

// GetCmp3c Compare 3 Interrupt flag Clear
func (r *registerTimdicrType) GetCmp3c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldCmp3cMask) != 0
}

// SetCmp3c Compare 3 Interrupt flag Clear
func (r *registerTimdicrType) SetCmp3c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldCmp3cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldCmp3cMask)
	}
}

const (
	RegisterTimdicrFieldCmp2cShift = 1
	RegisterTimdicrFieldCmp2cMask  = 0x2
)

// GetCmp2c Compare 2 Interrupt flag Clear
func (r *registerTimdicrType) GetCmp2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldCmp2cMask) != 0
}

// SetCmp2c Compare 2 Interrupt flag Clear
func (r *registerTimdicrType) SetCmp2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldCmp2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldCmp2cMask)
	}
}

const (
	RegisterTimdicrFieldCmp1cShift = 0
	RegisterTimdicrFieldCmp1cMask  = 0x1
)

// GetCmp1c Compare 1 Interrupt flag Clear
func (r *registerTimdicrType) GetCmp1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimdicrFieldCmp1cMask) != 0
}

// SetCmp1c Compare 1 Interrupt flag Clear
func (r *registerTimdicrType) SetCmp1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimdicrFieldCmp1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimdicrFieldCmp1cMask)
	}
}

// registerTimddier5Type TIMxDIER5
type registerTimddier5Type uint32

const (
	RegisterTimddier5FieldDlyprtdeShift = 30
	RegisterTimddier5FieldDlyprtdeMask  = 0x40000000
)

// GetDlyprtde DLYPRTDE
func (r *registerTimddier5Type) GetDlyprtde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldDlyprtdeMask) != 0
}

// SetDlyprtde DLYPRTDE
func (r *registerTimddier5Type) SetDlyprtde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldDlyprtdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldDlyprtdeMask)
	}
}

const (
	RegisterTimddier5FieldRstdeShift = 29
	RegisterTimddier5FieldRstdeMask  = 0x20000000
)

// GetRstde RSTDE
func (r *registerTimddier5Type) GetRstde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldRstdeMask) != 0
}

// SetRstde RSTDE
func (r *registerTimddier5Type) SetRstde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldRstdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldRstdeMask)
	}
}

const (
	RegisterTimddier5FieldRstx2deShift = 28
	RegisterTimddier5FieldRstx2deMask  = 0x10000000
)

// GetRstx2de RSTx2DE
func (r *registerTimddier5Type) GetRstx2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldRstx2deMask) != 0
}

// SetRstx2de RSTx2DE
func (r *registerTimddier5Type) SetRstx2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldRstx2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldRstx2deMask)
	}
}

const (
	RegisterTimddier5FieldSetx2deShift = 27
	RegisterTimddier5FieldSetx2deMask  = 0x8000000
)

// GetSetx2de SETx2DE
func (r *registerTimddier5Type) GetSetx2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldSetx2deMask) != 0
}

// SetSetx2de SETx2DE
func (r *registerTimddier5Type) SetSetx2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldSetx2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldSetx2deMask)
	}
}

const (
	RegisterTimddier5FieldRstx1deShift = 26
	RegisterTimddier5FieldRstx1deMask  = 0x4000000
)

// GetRstx1de RSTx1DE
func (r *registerTimddier5Type) GetRstx1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldRstx1deMask) != 0
}

// SetRstx1de RSTx1DE
func (r *registerTimddier5Type) SetRstx1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldRstx1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldRstx1deMask)
	}
}

const (
	RegisterTimddier5FieldSet1xdeShift = 25
	RegisterTimddier5FieldSet1xdeMask  = 0x2000000
)

// GetSet1xde SET1xDE
func (r *registerTimddier5Type) GetSet1xde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldSet1xdeMask) != 0
}

// SetSet1xde SET1xDE
func (r *registerTimddier5Type) SetSet1xde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldSet1xdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldSet1xdeMask)
	}
}

const (
	RegisterTimddier5FieldCpt2deShift = 24
	RegisterTimddier5FieldCpt2deMask  = 0x1000000
)

// GetCpt2de CPT2DE
func (r *registerTimddier5Type) GetCpt2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldCpt2deMask) != 0
}

// SetCpt2de CPT2DE
func (r *registerTimddier5Type) SetCpt2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldCpt2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldCpt2deMask)
	}
}

const (
	RegisterTimddier5FieldCpt1deShift = 23
	RegisterTimddier5FieldCpt1deMask  = 0x800000
)

// GetCpt1de CPT1DE
func (r *registerTimddier5Type) GetCpt1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldCpt1deMask) != 0
}

// SetCpt1de CPT1DE
func (r *registerTimddier5Type) SetCpt1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldCpt1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldCpt1deMask)
	}
}

const (
	RegisterTimddier5FieldUpddeShift = 22
	RegisterTimddier5FieldUpddeMask  = 0x400000
)

// GetUpdde UPDDE
func (r *registerTimddier5Type) GetUpdde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldUpddeMask) != 0
}

// SetUpdde UPDDE
func (r *registerTimddier5Type) SetUpdde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldUpddeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldUpddeMask)
	}
}

const (
	RegisterTimddier5FieldRepdeShift = 20
	RegisterTimddier5FieldRepdeMask  = 0x100000
)

// GetRepde REPDE
func (r *registerTimddier5Type) GetRepde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldRepdeMask) != 0
}

// SetRepde REPDE
func (r *registerTimddier5Type) SetRepde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldRepdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldRepdeMask)
	}
}

const (
	RegisterTimddier5FieldCmp4deShift = 19
	RegisterTimddier5FieldCmp4deMask  = 0x80000
)

// GetCmp4de CMP4DE
func (r *registerTimddier5Type) GetCmp4de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldCmp4deMask) != 0
}

// SetCmp4de CMP4DE
func (r *registerTimddier5Type) SetCmp4de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldCmp4deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldCmp4deMask)
	}
}

const (
	RegisterTimddier5FieldCmp3deShift = 18
	RegisterTimddier5FieldCmp3deMask  = 0x40000
)

// GetCmp3de CMP3DE
func (r *registerTimddier5Type) GetCmp3de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldCmp3deMask) != 0
}

// SetCmp3de CMP3DE
func (r *registerTimddier5Type) SetCmp3de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldCmp3deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldCmp3deMask)
	}
}

const (
	RegisterTimddier5FieldCmp2deShift = 17
	RegisterTimddier5FieldCmp2deMask  = 0x20000
)

// GetCmp2de CMP2DE
func (r *registerTimddier5Type) GetCmp2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldCmp2deMask) != 0
}

// SetCmp2de CMP2DE
func (r *registerTimddier5Type) SetCmp2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldCmp2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldCmp2deMask)
	}
}

const (
	RegisterTimddier5FieldCmp1deShift = 16
	RegisterTimddier5FieldCmp1deMask  = 0x10000
)

// GetCmp1de CMP1DE
func (r *registerTimddier5Type) GetCmp1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldCmp1deMask) != 0
}

// SetCmp1de CMP1DE
func (r *registerTimddier5Type) SetCmp1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldCmp1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldCmp1deMask)
	}
}

const (
	RegisterTimddier5FieldDlyprtieShift = 14
	RegisterTimddier5FieldDlyprtieMask  = 0x4000
)

// GetDlyprtie DLYPRTIE
func (r *registerTimddier5Type) GetDlyprtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldDlyprtieMask) != 0
}

// SetDlyprtie DLYPRTIE
func (r *registerTimddier5Type) SetDlyprtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldDlyprtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldDlyprtieMask)
	}
}

const (
	RegisterTimddier5FieldRstieShift = 13
	RegisterTimddier5FieldRstieMask  = 0x2000
)

// GetRstie RSTIE
func (r *registerTimddier5Type) GetRstie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldRstieMask) != 0
}

// SetRstie RSTIE
func (r *registerTimddier5Type) SetRstie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldRstieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldRstieMask)
	}
}

const (
	RegisterTimddier5FieldRstx2ieShift = 12
	RegisterTimddier5FieldRstx2ieMask  = 0x1000
)

// GetRstx2ie RSTx2IE
func (r *registerTimddier5Type) GetRstx2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldRstx2ieMask) != 0
}

// SetRstx2ie RSTx2IE
func (r *registerTimddier5Type) SetRstx2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldRstx2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldRstx2ieMask)
	}
}

const (
	RegisterTimddier5FieldSetx2ieShift = 11
	RegisterTimddier5FieldSetx2ieMask  = 0x800
)

// GetSetx2ie SETx2IE
func (r *registerTimddier5Type) GetSetx2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldSetx2ieMask) != 0
}

// SetSetx2ie SETx2IE
func (r *registerTimddier5Type) SetSetx2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldSetx2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldSetx2ieMask)
	}
}

const (
	RegisterTimddier5FieldRstx1ieShift = 10
	RegisterTimddier5FieldRstx1ieMask  = 0x400
)

// GetRstx1ie RSTx1IE
func (r *registerTimddier5Type) GetRstx1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldRstx1ieMask) != 0
}

// SetRstx1ie RSTx1IE
func (r *registerTimddier5Type) SetRstx1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldRstx1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldRstx1ieMask)
	}
}

const (
	RegisterTimddier5FieldSet1xieShift = 9
	RegisterTimddier5FieldSet1xieMask  = 0x200
)

// GetSet1xie SET1xIE
func (r *registerTimddier5Type) GetSet1xie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldSet1xieMask) != 0
}

// SetSet1xie SET1xIE
func (r *registerTimddier5Type) SetSet1xie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldSet1xieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldSet1xieMask)
	}
}

const (
	RegisterTimddier5FieldCpt2ieShift = 8
	RegisterTimddier5FieldCpt2ieMask  = 0x100
)

// GetCpt2ie CPT2IE
func (r *registerTimddier5Type) GetCpt2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldCpt2ieMask) != 0
}

// SetCpt2ie CPT2IE
func (r *registerTimddier5Type) SetCpt2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldCpt2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldCpt2ieMask)
	}
}

const (
	RegisterTimddier5FieldCpt1ieShift = 7
	RegisterTimddier5FieldCpt1ieMask  = 0x80
)

// GetCpt1ie CPT1IE
func (r *registerTimddier5Type) GetCpt1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldCpt1ieMask) != 0
}

// SetCpt1ie CPT1IE
func (r *registerTimddier5Type) SetCpt1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldCpt1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldCpt1ieMask)
	}
}

const (
	RegisterTimddier5FieldUpdieShift = 6
	RegisterTimddier5FieldUpdieMask  = 0x40
)

// GetUpdie UPDIE
func (r *registerTimddier5Type) GetUpdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldUpdieMask) != 0
}

// SetUpdie UPDIE
func (r *registerTimddier5Type) SetUpdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldUpdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldUpdieMask)
	}
}

const (
	RegisterTimddier5FieldRepieShift = 4
	RegisterTimddier5FieldRepieMask  = 0x10
)

// GetRepie REPIE
func (r *registerTimddier5Type) GetRepie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldRepieMask) != 0
}

// SetRepie REPIE
func (r *registerTimddier5Type) SetRepie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldRepieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldRepieMask)
	}
}

const (
	RegisterTimddier5FieldCmp4ieShift = 3
	RegisterTimddier5FieldCmp4ieMask  = 0x8
)

// GetCmp4ie CMP4IE
func (r *registerTimddier5Type) GetCmp4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldCmp4ieMask) != 0
}

// SetCmp4ie CMP4IE
func (r *registerTimddier5Type) SetCmp4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldCmp4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldCmp4ieMask)
	}
}

const (
	RegisterTimddier5FieldCmp3ieShift = 2
	RegisterTimddier5FieldCmp3ieMask  = 0x4
)

// GetCmp3ie CMP3IE
func (r *registerTimddier5Type) GetCmp3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldCmp3ieMask) != 0
}

// SetCmp3ie CMP3IE
func (r *registerTimddier5Type) SetCmp3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldCmp3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldCmp3ieMask)
	}
}

const (
	RegisterTimddier5FieldCmp2ieShift = 1
	RegisterTimddier5FieldCmp2ieMask  = 0x2
)

// GetCmp2ie CMP2IE
func (r *registerTimddier5Type) GetCmp2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldCmp2ieMask) != 0
}

// SetCmp2ie CMP2IE
func (r *registerTimddier5Type) SetCmp2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldCmp2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldCmp2ieMask)
	}
}

const (
	RegisterTimddier5FieldCmp1ieShift = 0
	RegisterTimddier5FieldCmp1ieMask  = 0x1
)

// GetCmp1ie CMP1IE
func (r *registerTimddier5Type) GetCmp1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimddier5FieldCmp1ieMask) != 0
}

// SetCmp1ie CMP1IE
func (r *registerTimddier5Type) SetCmp1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimddier5FieldCmp1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimddier5FieldCmp1ieMask)
	}
}

// registerCntdrType Timerx Counter Register
type registerCntdrType uint32

const (
	RegisterCntdrFieldCntxShift = 0
	RegisterCntdrFieldCntxMask  = 0xffff
)

// GetCntx Timerx Counter value
func (r *registerCntdrType) GetCntx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntdrFieldCntxMask) >> RegisterCntdrFieldCntxShift)
}

// SetCntx Timerx Counter value
func (r *registerCntdrType) SetCntx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntdrFieldCntxMask)|(uint32(value)<<RegisterCntdrFieldCntxShift))
}

// registerPerdrType Timerx Period Register
type registerPerdrType uint32

const (
	RegisterPerdrFieldPerxShift = 0
	RegisterPerdrFieldPerxMask  = 0xffff
)

// GetPerx Timerx Period value
func (r *registerPerdrType) GetPerx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPerdrFieldPerxMask) >> RegisterPerdrFieldPerxShift)
}

// SetPerx Timerx Period value
func (r *registerPerdrType) SetPerx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPerdrFieldPerxMask)|(uint32(value)<<RegisterPerdrFieldPerxShift))
}

// registerRepdrType Timerx Repetition Register
type registerRepdrType uint32

const (
	RegisterRepdrFieldRepxShift = 0
	RegisterRepdrFieldRepxMask  = 0xff
)

// GetRepx Timerx Repetition counter value
func (r *registerRepdrType) GetRepx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRepdrFieldRepxMask) >> RegisterRepdrFieldRepxShift)
}

// SetRepx Timerx Repetition counter value
func (r *registerRepdrType) SetRepx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRepdrFieldRepxMask)|(uint32(value)<<RegisterRepdrFieldRepxShift))
}

// registerCmp1drType Timerx Compare 1 Register
type registerCmp1drType uint32

const (
	RegisterCmp1drFieldCmp1xShift = 0
	RegisterCmp1drFieldCmp1xMask  = 0xffff
)

// GetCmp1x Timerx Compare 1 value
func (r *registerCmp1drType) GetCmp1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1drFieldCmp1xMask) >> RegisterCmp1drFieldCmp1xShift)
}

// SetCmp1x Timerx Compare 1 value
func (r *registerCmp1drType) SetCmp1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1drFieldCmp1xMask)|(uint32(value)<<RegisterCmp1drFieldCmp1xShift))
}

// registerCmp1cdrType Timerx Compare 1 Compound Register
type registerCmp1cdrType uint32

const (
	RegisterCmp1cdrFieldRepxShift = 16
	RegisterCmp1cdrFieldRepxMask  = 0xff0000
)

// GetRepx Timerx Repetition value (aliased from HRTIM_REPx register)
func (r *registerCmp1cdrType) GetRepx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1cdrFieldRepxMask) >> RegisterCmp1cdrFieldRepxShift)
}

// SetRepx Timerx Repetition value (aliased from HRTIM_REPx register)
func (r *registerCmp1cdrType) SetRepx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1cdrFieldRepxMask)|(uint32(value)<<RegisterCmp1cdrFieldRepxShift))
}

const (
	RegisterCmp1cdrFieldCmp1xShift = 0
	RegisterCmp1cdrFieldCmp1xMask  = 0xffff
)

// GetCmp1x Timerx Compare 1 value
func (r *registerCmp1cdrType) GetCmp1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1cdrFieldCmp1xMask) >> RegisterCmp1cdrFieldCmp1xShift)
}

// SetCmp1x Timerx Compare 1 value
func (r *registerCmp1cdrType) SetCmp1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1cdrFieldCmp1xMask)|(uint32(value)<<RegisterCmp1cdrFieldCmp1xShift))
}

// registerCmp2drType Timerx Compare 2 Register
type registerCmp2drType uint32

const (
	RegisterCmp2drFieldCmp2xShift = 0
	RegisterCmp2drFieldCmp2xMask  = 0xffff
)

// GetCmp2x Timerx Compare 2 value
func (r *registerCmp2drType) GetCmp2x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp2drFieldCmp2xMask) >> RegisterCmp2drFieldCmp2xShift)
}

// SetCmp2x Timerx Compare 2 value
func (r *registerCmp2drType) SetCmp2x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp2drFieldCmp2xMask)|(uint32(value)<<RegisterCmp2drFieldCmp2xShift))
}

// registerCmp3drType Timerx Compare 3 Register
type registerCmp3drType uint32

const (
	RegisterCmp3drFieldCmp3xShift = 0
	RegisterCmp3drFieldCmp3xMask  = 0xffff
)

// GetCmp3x Timerx Compare 3 value
func (r *registerCmp3drType) GetCmp3x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp3drFieldCmp3xMask) >> RegisterCmp3drFieldCmp3xShift)
}

// SetCmp3x Timerx Compare 3 value
func (r *registerCmp3drType) SetCmp3x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp3drFieldCmp3xMask)|(uint32(value)<<RegisterCmp3drFieldCmp3xShift))
}

// registerCmp4drType Timerx Compare 4 Register
type registerCmp4drType uint32

const (
	RegisterCmp4drFieldCmp4xShift = 0
	RegisterCmp4drFieldCmp4xMask  = 0xffff
)

// GetCmp4x Timerx Compare 4 value
func (r *registerCmp4drType) GetCmp4x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp4drFieldCmp4xMask) >> RegisterCmp4drFieldCmp4xShift)
}

// SetCmp4x Timerx Compare 4 value
func (r *registerCmp4drType) SetCmp4x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp4drFieldCmp4xMask)|(uint32(value)<<RegisterCmp4drFieldCmp4xShift))
}

// registerCpt1drType Timerx Capture 1 Register
type registerCpt1drType uint32

const (
	RegisterCpt1drFieldCpt1xShift = 0
	RegisterCpt1drFieldCpt1xMask  = 0xffff
)

// GetCpt1x Timerx Capture 1 value
func (r *registerCpt1drType) GetCpt1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpt1drFieldCpt1xMask) >> RegisterCpt1drFieldCpt1xShift)
}

// SetCpt1x Timerx Capture 1 value
func (r *registerCpt1drType) SetCpt1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpt1drFieldCpt1xMask)|(uint32(value)<<RegisterCpt1drFieldCpt1xShift))
}

// registerCpt2drType Timerx Capture 2 Register
type registerCpt2drType uint32

const (
	RegisterCpt2drFieldCpt2xShift = 0
	RegisterCpt2drFieldCpt2xMask  = 0xffff
)

// GetCpt2x Timerx Capture 2 value
func (r *registerCpt2drType) GetCpt2x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpt2drFieldCpt2xMask) >> RegisterCpt2drFieldCpt2xShift)
}

// SetCpt2x Timerx Capture 2 value
func (r *registerCpt2drType) SetCpt2x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpt2drFieldCpt2xMask)|(uint32(value)<<RegisterCpt2drFieldCpt2xShift))
}

// registerDtdrType Timerx Deadtime Register
type registerDtdrType uint32

const (
	RegisterDtdrFieldDtflkxShift = 31
	RegisterDtdrFieldDtflkxMask  = 0x80000000
)

// GetDtflkx Deadtime Falling Lock
func (r *registerDtdrType) GetDtflkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtdrFieldDtflkxMask) != 0
}

// SetDtflkx Deadtime Falling Lock
func (r *registerDtdrType) SetDtflkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtdrFieldDtflkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtdrFieldDtflkxMask)
	}
}

const (
	RegisterDtdrFieldDtfslkxShift = 30
	RegisterDtdrFieldDtfslkxMask  = 0x40000000
)

// GetDtfslkx Deadtime Falling Sign Lock
func (r *registerDtdrType) GetDtfslkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtdrFieldDtfslkxMask) != 0
}

// SetDtfslkx Deadtime Falling Sign Lock
func (r *registerDtdrType) SetDtfslkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtdrFieldDtfslkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtdrFieldDtfslkxMask)
	}
}

const (
	RegisterDtdrFieldSdtfxShift = 25
	RegisterDtdrFieldSdtfxMask  = 0x2000000
)

// GetSdtfx Sign Deadtime Falling value
func (r *registerDtdrType) GetSdtfx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtdrFieldSdtfxMask) != 0
}

// SetSdtfx Sign Deadtime Falling value
func (r *registerDtdrType) SetSdtfx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtdrFieldSdtfxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtdrFieldSdtfxMask)
	}
}

const (
	RegisterDtdrFieldDtfxShift = 16
	RegisterDtdrFieldDtfxMask  = 0x1ff0000
)

// GetDtfx Deadtime Falling value
func (r *registerDtdrType) GetDtfx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDtdrFieldDtfxMask) >> RegisterDtdrFieldDtfxShift)
}

// SetDtfx Deadtime Falling value
func (r *registerDtdrType) SetDtfx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtdrFieldDtfxMask)|(uint32(value)<<RegisterDtdrFieldDtfxShift))
}

const (
	RegisterDtdrFieldDtrlkxShift = 15
	RegisterDtdrFieldDtrlkxMask  = 0x8000
)

// GetDtrlkx Deadtime Rising Lock
func (r *registerDtdrType) GetDtrlkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtdrFieldDtrlkxMask) != 0
}

// SetDtrlkx Deadtime Rising Lock
func (r *registerDtdrType) SetDtrlkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtdrFieldDtrlkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtdrFieldDtrlkxMask)
	}
}

const (
	RegisterDtdrFieldDtrslkxShift = 14
	RegisterDtdrFieldDtrslkxMask  = 0x4000
)

// GetDtrslkx Deadtime Rising Sign Lock
func (r *registerDtdrType) GetDtrslkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtdrFieldDtrslkxMask) != 0
}

// SetDtrslkx Deadtime Rising Sign Lock
func (r *registerDtdrType) SetDtrslkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtdrFieldDtrslkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtdrFieldDtrslkxMask)
	}
}

const (
	RegisterDtdrFieldDtprscShift = 10
	RegisterDtdrFieldDtprscMask  = 0x1c00
)

// GetDtprsc Deadtime Prescaler
func (r *registerDtdrType) GetDtprsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDtdrFieldDtprscMask) >> RegisterDtdrFieldDtprscShift)
}

// SetDtprsc Deadtime Prescaler
func (r *registerDtdrType) SetDtprsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtdrFieldDtprscMask)|(uint32(value)<<RegisterDtdrFieldDtprscShift))
}

const (
	RegisterDtdrFieldSdtrxShift = 9
	RegisterDtdrFieldSdtrxMask  = 0x200
)

// GetSdtrx Sign Deadtime Rising value
func (r *registerDtdrType) GetSdtrx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtdrFieldSdtrxMask) != 0
}

// SetSdtrx Sign Deadtime Rising value
func (r *registerDtdrType) SetSdtrx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtdrFieldSdtrxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtdrFieldSdtrxMask)
	}
}

const (
	RegisterDtdrFieldDtrxShift = 0
	RegisterDtdrFieldDtrxMask  = 0x1ff
)

// GetDtrx Deadtime Rising value
func (r *registerDtdrType) GetDtrx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDtdrFieldDtrxMask) >> RegisterDtdrFieldDtrxShift)
}

// SetDtrx Deadtime Rising value
func (r *registerDtdrType) SetDtrx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtdrFieldDtrxMask)|(uint32(value)<<RegisterDtdrFieldDtrxShift))
}

// registerSetd1rType Timerx Output1 Set Register
type registerSetd1rType uint32

const (
	RegisterSetd1rFieldUpdateShift = 31
	RegisterSetd1rFieldUpdateMask  = 0x80000000
)

// GetUpdate Registers update (transfer preload to active)
func (r *registerSetd1rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldUpdateMask) != 0
}

// SetUpdate Registers update (transfer preload to active)
func (r *registerSetd1rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldUpdateMask)
	}
}

const (
	RegisterSetd1rFieldExtevnt10Shift = 30
	RegisterSetd1rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 External Event 10
func (r *registerSetd1rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 External Event 10
func (r *registerSetd1rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldExtevnt10Mask)
	}
}

const (
	RegisterSetd1rFieldExtevnt9Shift = 29
	RegisterSetd1rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 External Event 9
func (r *registerSetd1rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 External Event 9
func (r *registerSetd1rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldExtevnt9Mask)
	}
}

const (
	RegisterSetd1rFieldExtevnt8Shift = 28
	RegisterSetd1rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 External Event 8
func (r *registerSetd1rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 External Event 8
func (r *registerSetd1rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldExtevnt8Mask)
	}
}

const (
	RegisterSetd1rFieldExtevnt7Shift = 27
	RegisterSetd1rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 External Event 7
func (r *registerSetd1rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 External Event 7
func (r *registerSetd1rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldExtevnt7Mask)
	}
}

const (
	RegisterSetd1rFieldExtevnt6Shift = 26
	RegisterSetd1rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 External Event 6
func (r *registerSetd1rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 External Event 6
func (r *registerSetd1rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldExtevnt6Mask)
	}
}

const (
	RegisterSetd1rFieldExtevnt5Shift = 25
	RegisterSetd1rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 External Event 5
func (r *registerSetd1rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 External Event 5
func (r *registerSetd1rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldExtevnt5Mask)
	}
}

const (
	RegisterSetd1rFieldExtevnt4Shift = 24
	RegisterSetd1rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 External Event 4
func (r *registerSetd1rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 External Event 4
func (r *registerSetd1rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldExtevnt4Mask)
	}
}

const (
	RegisterSetd1rFieldExtevnt3Shift = 23
	RegisterSetd1rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 External Event 3
func (r *registerSetd1rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 External Event 3
func (r *registerSetd1rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldExtevnt3Mask)
	}
}

const (
	RegisterSetd1rFieldExtevnt2Shift = 22
	RegisterSetd1rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 External Event 2
func (r *registerSetd1rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 External Event 2
func (r *registerSetd1rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldExtevnt2Mask)
	}
}

const (
	RegisterSetd1rFieldExtevnt1Shift = 21
	RegisterSetd1rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 External Event 1
func (r *registerSetd1rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 External Event 1
func (r *registerSetd1rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldExtevnt1Mask)
	}
}

const (
	RegisterSetd1rFieldTimevnt9Shift = 20
	RegisterSetd1rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 Timer Event 9
func (r *registerSetd1rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 Timer Event 9
func (r *registerSetd1rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldTimevnt9Mask)
	}
}

const (
	RegisterSetd1rFieldTimevnt8Shift = 19
	RegisterSetd1rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 Timer Event 8
func (r *registerSetd1rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 Timer Event 8
func (r *registerSetd1rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldTimevnt8Mask)
	}
}

const (
	RegisterSetd1rFieldTimevnt7Shift = 18
	RegisterSetd1rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 Timer Event 7
func (r *registerSetd1rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 Timer Event 7
func (r *registerSetd1rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldTimevnt7Mask)
	}
}

const (
	RegisterSetd1rFieldTimevnt6Shift = 17
	RegisterSetd1rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 Timer Event 6
func (r *registerSetd1rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 Timer Event 6
func (r *registerSetd1rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldTimevnt6Mask)
	}
}

const (
	RegisterSetd1rFieldTimevnt5Shift = 16
	RegisterSetd1rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 Timer Event 5
func (r *registerSetd1rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 Timer Event 5
func (r *registerSetd1rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldTimevnt5Mask)
	}
}

const (
	RegisterSetd1rFieldTimevnt4Shift = 15
	RegisterSetd1rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 Timer Event 4
func (r *registerSetd1rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 Timer Event 4
func (r *registerSetd1rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldTimevnt4Mask)
	}
}

const (
	RegisterSetd1rFieldTimevnt3Shift = 14
	RegisterSetd1rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 Timer Event 3
func (r *registerSetd1rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 Timer Event 3
func (r *registerSetd1rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldTimevnt3Mask)
	}
}

const (
	RegisterSetd1rFieldTimevnt2Shift = 13
	RegisterSetd1rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 Timer Event 2
func (r *registerSetd1rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 Timer Event 2
func (r *registerSetd1rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldTimevnt2Mask)
	}
}

const (
	RegisterSetd1rFieldTimevnt1Shift = 12
	RegisterSetd1rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 Timer Event 1
func (r *registerSetd1rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 Timer Event 1
func (r *registerSetd1rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldTimevnt1Mask)
	}
}

const (
	RegisterSetd1rFieldMstcmp4Shift = 11
	RegisterSetd1rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 Master Compare 4
func (r *registerSetd1rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 Master Compare 4
func (r *registerSetd1rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldMstcmp4Mask)
	}
}

const (
	RegisterSetd1rFieldMstcmp3Shift = 10
	RegisterSetd1rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 Master Compare 3
func (r *registerSetd1rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 Master Compare 3
func (r *registerSetd1rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldMstcmp3Mask)
	}
}

const (
	RegisterSetd1rFieldMstcmp2Shift = 9
	RegisterSetd1rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 Master Compare 2
func (r *registerSetd1rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 Master Compare 2
func (r *registerSetd1rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldMstcmp2Mask)
	}
}

const (
	RegisterSetd1rFieldMstcmp1Shift = 8
	RegisterSetd1rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 Master Compare 1
func (r *registerSetd1rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 Master Compare 1
func (r *registerSetd1rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldMstcmp1Mask)
	}
}

const (
	RegisterSetd1rFieldMstperShift = 7
	RegisterSetd1rFieldMstperMask  = 0x80
)

// GetMstper Master Period
func (r *registerSetd1rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldMstperMask) != 0
}

// SetMstper Master Period
func (r *registerSetd1rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldMstperMask)
	}
}

const (
	RegisterSetd1rFieldCmp4Shift = 6
	RegisterSetd1rFieldCmp4Mask  = 0x40
)

// GetCmp4 Timer A compare 4
func (r *registerSetd1rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldCmp4Mask) != 0
}

// SetCmp4 Timer A compare 4
func (r *registerSetd1rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldCmp4Mask)
	}
}

const (
	RegisterSetd1rFieldCmp3Shift = 5
	RegisterSetd1rFieldCmp3Mask  = 0x20
)

// GetCmp3 Timer A compare 3
func (r *registerSetd1rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldCmp3Mask) != 0
}

// SetCmp3 Timer A compare 3
func (r *registerSetd1rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldCmp3Mask)
	}
}

const (
	RegisterSetd1rFieldCmp2Shift = 4
	RegisterSetd1rFieldCmp2Mask  = 0x10
)

// GetCmp2 Timer A compare 2
func (r *registerSetd1rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldCmp2Mask) != 0
}

// SetCmp2 Timer A compare 2
func (r *registerSetd1rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldCmp2Mask)
	}
}

const (
	RegisterSetd1rFieldCmp1Shift = 3
	RegisterSetd1rFieldCmp1Mask  = 0x8
)

// GetCmp1 Timer A compare 1
func (r *registerSetd1rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldCmp1Mask) != 0
}

// SetCmp1 Timer A compare 1
func (r *registerSetd1rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldCmp1Mask)
	}
}

const (
	RegisterSetd1rFieldPerShift = 2
	RegisterSetd1rFieldPerMask  = 0x4
)

// GetPer Timer A Period
func (r *registerSetd1rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldPerMask) != 0
}

// SetPer Timer A Period
func (r *registerSetd1rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldPerMask)
	}
}

const (
	RegisterSetd1rFieldResyncShift = 1
	RegisterSetd1rFieldResyncMask  = 0x2
)

// GetResync Timer A resynchronizaton
func (r *registerSetd1rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldResyncMask) != 0
}

// SetResync Timer A resynchronizaton
func (r *registerSetd1rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldResyncMask)
	}
}

const (
	RegisterSetd1rFieldSstShift = 0
	RegisterSetd1rFieldSstMask  = 0x1
)

// GetSst Software Set trigger
func (r *registerSetd1rType) GetSst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd1rFieldSstMask) != 0
}

// SetSst Software Set trigger
func (r *registerSetd1rType) SetSst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd1rFieldSstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd1rFieldSstMask)
	}
}

// registerRstd1rType Timerx Output1 Reset Register
type registerRstd1rType uint32

const (
	RegisterRstd1rFieldUpdateShift = 31
	RegisterRstd1rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerRstd1rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerRstd1rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldUpdateMask)
	}
}

const (
	RegisterRstd1rFieldExtevnt10Shift = 30
	RegisterRstd1rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerRstd1rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerRstd1rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldExtevnt10Mask)
	}
}

const (
	RegisterRstd1rFieldExtevnt9Shift = 29
	RegisterRstd1rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerRstd1rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerRstd1rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldExtevnt9Mask)
	}
}

const (
	RegisterRstd1rFieldExtevnt8Shift = 28
	RegisterRstd1rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerRstd1rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerRstd1rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldExtevnt8Mask)
	}
}

const (
	RegisterRstd1rFieldExtevnt7Shift = 27
	RegisterRstd1rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerRstd1rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerRstd1rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldExtevnt7Mask)
	}
}

const (
	RegisterRstd1rFieldExtevnt6Shift = 26
	RegisterRstd1rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerRstd1rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerRstd1rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldExtevnt6Mask)
	}
}

const (
	RegisterRstd1rFieldExtevnt5Shift = 25
	RegisterRstd1rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerRstd1rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerRstd1rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldExtevnt5Mask)
	}
}

const (
	RegisterRstd1rFieldExtevnt4Shift = 24
	RegisterRstd1rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerRstd1rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerRstd1rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldExtevnt4Mask)
	}
}

const (
	RegisterRstd1rFieldExtevnt3Shift = 23
	RegisterRstd1rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerRstd1rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerRstd1rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldExtevnt3Mask)
	}
}

const (
	RegisterRstd1rFieldExtevnt2Shift = 22
	RegisterRstd1rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerRstd1rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerRstd1rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldExtevnt2Mask)
	}
}

const (
	RegisterRstd1rFieldExtevnt1Shift = 21
	RegisterRstd1rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerRstd1rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerRstd1rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldExtevnt1Mask)
	}
}

const (
	RegisterRstd1rFieldTimevnt9Shift = 20
	RegisterRstd1rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerRstd1rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerRstd1rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldTimevnt9Mask)
	}
}

const (
	RegisterRstd1rFieldTimevnt8Shift = 19
	RegisterRstd1rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerRstd1rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerRstd1rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldTimevnt8Mask)
	}
}

const (
	RegisterRstd1rFieldTimevnt7Shift = 18
	RegisterRstd1rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerRstd1rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerRstd1rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldTimevnt7Mask)
	}
}

const (
	RegisterRstd1rFieldTimevnt6Shift = 17
	RegisterRstd1rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerRstd1rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerRstd1rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldTimevnt6Mask)
	}
}

const (
	RegisterRstd1rFieldTimevnt5Shift = 16
	RegisterRstd1rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerRstd1rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerRstd1rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldTimevnt5Mask)
	}
}

const (
	RegisterRstd1rFieldTimevnt4Shift = 15
	RegisterRstd1rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerRstd1rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerRstd1rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldTimevnt4Mask)
	}
}

const (
	RegisterRstd1rFieldTimevnt3Shift = 14
	RegisterRstd1rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerRstd1rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerRstd1rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldTimevnt3Mask)
	}
}

const (
	RegisterRstd1rFieldTimevnt2Shift = 13
	RegisterRstd1rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerRstd1rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerRstd1rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldTimevnt2Mask)
	}
}

const (
	RegisterRstd1rFieldTimevnt1Shift = 12
	RegisterRstd1rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerRstd1rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerRstd1rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldTimevnt1Mask)
	}
}

const (
	RegisterRstd1rFieldMstcmp4Shift = 11
	RegisterRstd1rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerRstd1rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerRstd1rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldMstcmp4Mask)
	}
}

const (
	RegisterRstd1rFieldMstcmp3Shift = 10
	RegisterRstd1rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerRstd1rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerRstd1rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldMstcmp3Mask)
	}
}

const (
	RegisterRstd1rFieldMstcmp2Shift = 9
	RegisterRstd1rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerRstd1rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerRstd1rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldMstcmp2Mask)
	}
}

const (
	RegisterRstd1rFieldMstcmp1Shift = 8
	RegisterRstd1rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerRstd1rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerRstd1rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldMstcmp1Mask)
	}
}

const (
	RegisterRstd1rFieldMstperShift = 7
	RegisterRstd1rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerRstd1rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerRstd1rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldMstperMask)
	}
}

const (
	RegisterRstd1rFieldCmp4Shift = 6
	RegisterRstd1rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerRstd1rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerRstd1rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldCmp4Mask)
	}
}

const (
	RegisterRstd1rFieldCmp3Shift = 5
	RegisterRstd1rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerRstd1rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerRstd1rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldCmp3Mask)
	}
}

const (
	RegisterRstd1rFieldCmp2Shift = 4
	RegisterRstd1rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerRstd1rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerRstd1rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldCmp2Mask)
	}
}

const (
	RegisterRstd1rFieldCmp1Shift = 3
	RegisterRstd1rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerRstd1rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerRstd1rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldCmp1Mask)
	}
}

const (
	RegisterRstd1rFieldPerShift = 2
	RegisterRstd1rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerRstd1rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldPerMask) != 0
}

// SetPer PER
func (r *registerRstd1rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldPerMask)
	}
}

const (
	RegisterRstd1rFieldResyncShift = 1
	RegisterRstd1rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerRstd1rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerRstd1rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldResyncMask)
	}
}

const (
	RegisterRstd1rFieldSrtShift = 0
	RegisterRstd1rFieldSrtMask  = 0x1
)

// GetSrt SRT
func (r *registerRstd1rType) GetSrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd1rFieldSrtMask) != 0
}

// SetSrt SRT
func (r *registerRstd1rType) SetSrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd1rFieldSrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd1rFieldSrtMask)
	}
}

// registerSetd2rType Timerx Output2 Set Register
type registerSetd2rType uint32

const (
	RegisterSetd2rFieldUpdateShift = 31
	RegisterSetd2rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerSetd2rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerSetd2rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldUpdateMask)
	}
}

const (
	RegisterSetd2rFieldExtevnt10Shift = 30
	RegisterSetd2rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerSetd2rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerSetd2rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldExtevnt10Mask)
	}
}

const (
	RegisterSetd2rFieldExtevnt9Shift = 29
	RegisterSetd2rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerSetd2rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerSetd2rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldExtevnt9Mask)
	}
}

const (
	RegisterSetd2rFieldExtevnt8Shift = 28
	RegisterSetd2rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerSetd2rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerSetd2rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldExtevnt8Mask)
	}
}

const (
	RegisterSetd2rFieldExtevnt7Shift = 27
	RegisterSetd2rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerSetd2rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerSetd2rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldExtevnt7Mask)
	}
}

const (
	RegisterSetd2rFieldExtevnt6Shift = 26
	RegisterSetd2rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerSetd2rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerSetd2rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldExtevnt6Mask)
	}
}

const (
	RegisterSetd2rFieldExtevnt5Shift = 25
	RegisterSetd2rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerSetd2rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerSetd2rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldExtevnt5Mask)
	}
}

const (
	RegisterSetd2rFieldExtevnt4Shift = 24
	RegisterSetd2rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerSetd2rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerSetd2rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldExtevnt4Mask)
	}
}

const (
	RegisterSetd2rFieldExtevnt3Shift = 23
	RegisterSetd2rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerSetd2rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerSetd2rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldExtevnt3Mask)
	}
}

const (
	RegisterSetd2rFieldExtevnt2Shift = 22
	RegisterSetd2rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerSetd2rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerSetd2rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldExtevnt2Mask)
	}
}

const (
	RegisterSetd2rFieldExtevnt1Shift = 21
	RegisterSetd2rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerSetd2rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerSetd2rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldExtevnt1Mask)
	}
}

const (
	RegisterSetd2rFieldTimevnt9Shift = 20
	RegisterSetd2rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerSetd2rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerSetd2rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldTimevnt9Mask)
	}
}

const (
	RegisterSetd2rFieldTimevnt8Shift = 19
	RegisterSetd2rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerSetd2rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerSetd2rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldTimevnt8Mask)
	}
}

const (
	RegisterSetd2rFieldTimevnt7Shift = 18
	RegisterSetd2rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerSetd2rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerSetd2rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldTimevnt7Mask)
	}
}

const (
	RegisterSetd2rFieldTimevnt6Shift = 17
	RegisterSetd2rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerSetd2rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerSetd2rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldTimevnt6Mask)
	}
}

const (
	RegisterSetd2rFieldTimevnt5Shift = 16
	RegisterSetd2rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerSetd2rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerSetd2rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldTimevnt5Mask)
	}
}

const (
	RegisterSetd2rFieldTimevnt4Shift = 15
	RegisterSetd2rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerSetd2rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerSetd2rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldTimevnt4Mask)
	}
}

const (
	RegisterSetd2rFieldTimevnt3Shift = 14
	RegisterSetd2rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerSetd2rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerSetd2rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldTimevnt3Mask)
	}
}

const (
	RegisterSetd2rFieldTimevnt2Shift = 13
	RegisterSetd2rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerSetd2rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerSetd2rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldTimevnt2Mask)
	}
}

const (
	RegisterSetd2rFieldTimevnt1Shift = 12
	RegisterSetd2rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerSetd2rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerSetd2rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldTimevnt1Mask)
	}
}

const (
	RegisterSetd2rFieldMstcmp4Shift = 11
	RegisterSetd2rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerSetd2rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerSetd2rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldMstcmp4Mask)
	}
}

const (
	RegisterSetd2rFieldMstcmp3Shift = 10
	RegisterSetd2rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerSetd2rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerSetd2rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldMstcmp3Mask)
	}
}

const (
	RegisterSetd2rFieldMstcmp2Shift = 9
	RegisterSetd2rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerSetd2rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerSetd2rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldMstcmp2Mask)
	}
}

const (
	RegisterSetd2rFieldMstcmp1Shift = 8
	RegisterSetd2rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerSetd2rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerSetd2rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldMstcmp1Mask)
	}
}

const (
	RegisterSetd2rFieldMstperShift = 7
	RegisterSetd2rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerSetd2rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerSetd2rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldMstperMask)
	}
}

const (
	RegisterSetd2rFieldCmp4Shift = 6
	RegisterSetd2rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerSetd2rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerSetd2rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldCmp4Mask)
	}
}

const (
	RegisterSetd2rFieldCmp3Shift = 5
	RegisterSetd2rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerSetd2rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerSetd2rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldCmp3Mask)
	}
}

const (
	RegisterSetd2rFieldCmp2Shift = 4
	RegisterSetd2rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerSetd2rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerSetd2rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldCmp2Mask)
	}
}

const (
	RegisterSetd2rFieldCmp1Shift = 3
	RegisterSetd2rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerSetd2rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerSetd2rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldCmp1Mask)
	}
}

const (
	RegisterSetd2rFieldPerShift = 2
	RegisterSetd2rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerSetd2rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldPerMask) != 0
}

// SetPer PER
func (r *registerSetd2rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldPerMask)
	}
}

const (
	RegisterSetd2rFieldResyncShift = 1
	RegisterSetd2rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerSetd2rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerSetd2rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldResyncMask)
	}
}

const (
	RegisterSetd2rFieldSstShift = 0
	RegisterSetd2rFieldSstMask  = 0x1
)

// GetSst SST
func (r *registerSetd2rType) GetSst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetd2rFieldSstMask) != 0
}

// SetSst SST
func (r *registerSetd2rType) SetSst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetd2rFieldSstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetd2rFieldSstMask)
	}
}

// registerRstd2rType Timerx Output2 Reset Register
type registerRstd2rType uint32

const (
	RegisterRstd2rFieldUpdateShift = 31
	RegisterRstd2rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerRstd2rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerRstd2rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldUpdateMask)
	}
}

const (
	RegisterRstd2rFieldExtevnt10Shift = 30
	RegisterRstd2rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerRstd2rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerRstd2rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldExtevnt10Mask)
	}
}

const (
	RegisterRstd2rFieldExtevnt9Shift = 29
	RegisterRstd2rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerRstd2rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerRstd2rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldExtevnt9Mask)
	}
}

const (
	RegisterRstd2rFieldExtevnt8Shift = 28
	RegisterRstd2rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerRstd2rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerRstd2rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldExtevnt8Mask)
	}
}

const (
	RegisterRstd2rFieldExtevnt7Shift = 27
	RegisterRstd2rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerRstd2rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerRstd2rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldExtevnt7Mask)
	}
}

const (
	RegisterRstd2rFieldExtevnt6Shift = 26
	RegisterRstd2rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerRstd2rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerRstd2rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldExtevnt6Mask)
	}
}

const (
	RegisterRstd2rFieldExtevnt5Shift = 25
	RegisterRstd2rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerRstd2rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerRstd2rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldExtevnt5Mask)
	}
}

const (
	RegisterRstd2rFieldExtevnt4Shift = 24
	RegisterRstd2rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerRstd2rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerRstd2rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldExtevnt4Mask)
	}
}

const (
	RegisterRstd2rFieldExtevnt3Shift = 23
	RegisterRstd2rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerRstd2rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerRstd2rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldExtevnt3Mask)
	}
}

const (
	RegisterRstd2rFieldExtevnt2Shift = 22
	RegisterRstd2rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerRstd2rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerRstd2rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldExtevnt2Mask)
	}
}

const (
	RegisterRstd2rFieldExtevnt1Shift = 21
	RegisterRstd2rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerRstd2rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerRstd2rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldExtevnt1Mask)
	}
}

const (
	RegisterRstd2rFieldTimevnt9Shift = 20
	RegisterRstd2rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerRstd2rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerRstd2rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldTimevnt9Mask)
	}
}

const (
	RegisterRstd2rFieldTimevnt8Shift = 19
	RegisterRstd2rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerRstd2rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerRstd2rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldTimevnt8Mask)
	}
}

const (
	RegisterRstd2rFieldTimevnt7Shift = 18
	RegisterRstd2rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerRstd2rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerRstd2rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldTimevnt7Mask)
	}
}

const (
	RegisterRstd2rFieldTimevnt6Shift = 17
	RegisterRstd2rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerRstd2rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerRstd2rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldTimevnt6Mask)
	}
}

const (
	RegisterRstd2rFieldTimevnt5Shift = 16
	RegisterRstd2rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerRstd2rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerRstd2rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldTimevnt5Mask)
	}
}

const (
	RegisterRstd2rFieldTimevnt4Shift = 15
	RegisterRstd2rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerRstd2rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerRstd2rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldTimevnt4Mask)
	}
}

const (
	RegisterRstd2rFieldTimevnt3Shift = 14
	RegisterRstd2rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerRstd2rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerRstd2rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldTimevnt3Mask)
	}
}

const (
	RegisterRstd2rFieldTimevnt2Shift = 13
	RegisterRstd2rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerRstd2rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerRstd2rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldTimevnt2Mask)
	}
}

const (
	RegisterRstd2rFieldTimevnt1Shift = 12
	RegisterRstd2rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerRstd2rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerRstd2rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldTimevnt1Mask)
	}
}

const (
	RegisterRstd2rFieldMstcmp4Shift = 11
	RegisterRstd2rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerRstd2rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerRstd2rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldMstcmp4Mask)
	}
}

const (
	RegisterRstd2rFieldMstcmp3Shift = 10
	RegisterRstd2rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerRstd2rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerRstd2rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldMstcmp3Mask)
	}
}

const (
	RegisterRstd2rFieldMstcmp2Shift = 9
	RegisterRstd2rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerRstd2rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerRstd2rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldMstcmp2Mask)
	}
}

const (
	RegisterRstd2rFieldMstcmp1Shift = 8
	RegisterRstd2rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerRstd2rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerRstd2rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldMstcmp1Mask)
	}
}

const (
	RegisterRstd2rFieldMstperShift = 7
	RegisterRstd2rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerRstd2rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerRstd2rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldMstperMask)
	}
}

const (
	RegisterRstd2rFieldCmp4Shift = 6
	RegisterRstd2rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerRstd2rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerRstd2rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldCmp4Mask)
	}
}

const (
	RegisterRstd2rFieldCmp3Shift = 5
	RegisterRstd2rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerRstd2rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerRstd2rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldCmp3Mask)
	}
}

const (
	RegisterRstd2rFieldCmp2Shift = 4
	RegisterRstd2rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerRstd2rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerRstd2rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldCmp2Mask)
	}
}

const (
	RegisterRstd2rFieldCmp1Shift = 3
	RegisterRstd2rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerRstd2rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerRstd2rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldCmp1Mask)
	}
}

const (
	RegisterRstd2rFieldPerShift = 2
	RegisterRstd2rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerRstd2rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldPerMask) != 0
}

// SetPer PER
func (r *registerRstd2rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldPerMask)
	}
}

const (
	RegisterRstd2rFieldResyncShift = 1
	RegisterRstd2rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerRstd2rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerRstd2rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldResyncMask)
	}
}

const (
	RegisterRstd2rFieldSrtShift = 0
	RegisterRstd2rFieldSrtMask  = 0x1
)

// GetSrt SRT
func (r *registerRstd2rType) GetSrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstd2rFieldSrtMask) != 0
}

// SetSrt SRT
func (r *registerRstd2rType) SetSrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstd2rFieldSrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstd2rFieldSrtMask)
	}
}

// registerEefdr1Type Timerx External Event Filtering Register 1
type registerEefdr1Type uint32

const (
	RegisterEefdr1FieldEe5fltrShift = 25
	RegisterEefdr1FieldEe5fltrMask  = 0x1e000000
)

// GetEe5fltr External Event 5 filter
func (r *registerEefdr1Type) GetEe5fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefdr1FieldEe5fltrMask) >> RegisterEefdr1FieldEe5fltrShift)
}

// SetEe5fltr External Event 5 filter
func (r *registerEefdr1Type) SetEe5fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefdr1FieldEe5fltrMask)|(uint32(value)<<RegisterEefdr1FieldEe5fltrShift))
}

const (
	RegisterEefdr1FieldEe5ltchShift = 24
	RegisterEefdr1FieldEe5ltchMask  = 0x1000000
)

// GetEe5ltch External Event 5 latch
func (r *registerEefdr1Type) GetEe5ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefdr1FieldEe5ltchMask) != 0
}

// SetEe5ltch External Event 5 latch
func (r *registerEefdr1Type) SetEe5ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefdr1FieldEe5ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefdr1FieldEe5ltchMask)
	}
}

const (
	RegisterEefdr1FieldEe4fltrShift = 19
	RegisterEefdr1FieldEe4fltrMask  = 0x780000
)

// GetEe4fltr External Event 4 filter
func (r *registerEefdr1Type) GetEe4fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefdr1FieldEe4fltrMask) >> RegisterEefdr1FieldEe4fltrShift)
}

// SetEe4fltr External Event 4 filter
func (r *registerEefdr1Type) SetEe4fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefdr1FieldEe4fltrMask)|(uint32(value)<<RegisterEefdr1FieldEe4fltrShift))
}

const (
	RegisterEefdr1FieldEe4ltchShift = 18
	RegisterEefdr1FieldEe4ltchMask  = 0x40000
)

// GetEe4ltch External Event 4 latch
func (r *registerEefdr1Type) GetEe4ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefdr1FieldEe4ltchMask) != 0
}

// SetEe4ltch External Event 4 latch
func (r *registerEefdr1Type) SetEe4ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefdr1FieldEe4ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefdr1FieldEe4ltchMask)
	}
}

const (
	RegisterEefdr1FieldEe3fltrShift = 13
	RegisterEefdr1FieldEe3fltrMask  = 0x1e000
)

// GetEe3fltr External Event 3 filter
func (r *registerEefdr1Type) GetEe3fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefdr1FieldEe3fltrMask) >> RegisterEefdr1FieldEe3fltrShift)
}

// SetEe3fltr External Event 3 filter
func (r *registerEefdr1Type) SetEe3fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefdr1FieldEe3fltrMask)|(uint32(value)<<RegisterEefdr1FieldEe3fltrShift))
}

const (
	RegisterEefdr1FieldEe3ltchShift = 12
	RegisterEefdr1FieldEe3ltchMask  = 0x1000
)

// GetEe3ltch External Event 3 latch
func (r *registerEefdr1Type) GetEe3ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefdr1FieldEe3ltchMask) != 0
}

// SetEe3ltch External Event 3 latch
func (r *registerEefdr1Type) SetEe3ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefdr1FieldEe3ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefdr1FieldEe3ltchMask)
	}
}

const (
	RegisterEefdr1FieldEe2fltrShift = 7
	RegisterEefdr1FieldEe2fltrMask  = 0x780
)

// GetEe2fltr External Event 2 filter
func (r *registerEefdr1Type) GetEe2fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefdr1FieldEe2fltrMask) >> RegisterEefdr1FieldEe2fltrShift)
}

// SetEe2fltr External Event 2 filter
func (r *registerEefdr1Type) SetEe2fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefdr1FieldEe2fltrMask)|(uint32(value)<<RegisterEefdr1FieldEe2fltrShift))
}

const (
	RegisterEefdr1FieldEe2ltchShift = 6
	RegisterEefdr1FieldEe2ltchMask  = 0x40
)

// GetEe2ltch External Event 2 latch
func (r *registerEefdr1Type) GetEe2ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefdr1FieldEe2ltchMask) != 0
}

// SetEe2ltch External Event 2 latch
func (r *registerEefdr1Type) SetEe2ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefdr1FieldEe2ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefdr1FieldEe2ltchMask)
	}
}

const (
	RegisterEefdr1FieldEe1fltrShift = 1
	RegisterEefdr1FieldEe1fltrMask  = 0x1e
)

// GetEe1fltr External Event 1 filter
func (r *registerEefdr1Type) GetEe1fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefdr1FieldEe1fltrMask) >> RegisterEefdr1FieldEe1fltrShift)
}

// SetEe1fltr External Event 1 filter
func (r *registerEefdr1Type) SetEe1fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefdr1FieldEe1fltrMask)|(uint32(value)<<RegisterEefdr1FieldEe1fltrShift))
}

const (
	RegisterEefdr1FieldEe1ltchShift = 0
	RegisterEefdr1FieldEe1ltchMask  = 0x1
)

// GetEe1ltch External Event 1 latch
func (r *registerEefdr1Type) GetEe1ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefdr1FieldEe1ltchMask) != 0
}

// SetEe1ltch External Event 1 latch
func (r *registerEefdr1Type) SetEe1ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefdr1FieldEe1ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefdr1FieldEe1ltchMask)
	}
}

// registerEefdr2Type Timerx External Event Filtering Register 2
type registerEefdr2Type uint32

const (
	RegisterEefdr2FieldEe10fltrShift = 25
	RegisterEefdr2FieldEe10fltrMask  = 0x1e000000
)

// GetEe10fltr External Event 10 filter
func (r *registerEefdr2Type) GetEe10fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefdr2FieldEe10fltrMask) >> RegisterEefdr2FieldEe10fltrShift)
}

// SetEe10fltr External Event 10 filter
func (r *registerEefdr2Type) SetEe10fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefdr2FieldEe10fltrMask)|(uint32(value)<<RegisterEefdr2FieldEe10fltrShift))
}

const (
	RegisterEefdr2FieldEe10ltchShift = 24
	RegisterEefdr2FieldEe10ltchMask  = 0x1000000
)

// GetEe10ltch External Event 10 latch
func (r *registerEefdr2Type) GetEe10ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefdr2FieldEe10ltchMask) != 0
}

// SetEe10ltch External Event 10 latch
func (r *registerEefdr2Type) SetEe10ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefdr2FieldEe10ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefdr2FieldEe10ltchMask)
	}
}

const (
	RegisterEefdr2FieldEe9fltrShift = 19
	RegisterEefdr2FieldEe9fltrMask  = 0x780000
)

// GetEe9fltr External Event 9 filter
func (r *registerEefdr2Type) GetEe9fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefdr2FieldEe9fltrMask) >> RegisterEefdr2FieldEe9fltrShift)
}

// SetEe9fltr External Event 9 filter
func (r *registerEefdr2Type) SetEe9fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefdr2FieldEe9fltrMask)|(uint32(value)<<RegisterEefdr2FieldEe9fltrShift))
}

const (
	RegisterEefdr2FieldEe9ltchShift = 18
	RegisterEefdr2FieldEe9ltchMask  = 0x40000
)

// GetEe9ltch External Event 9 latch
func (r *registerEefdr2Type) GetEe9ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefdr2FieldEe9ltchMask) != 0
}

// SetEe9ltch External Event 9 latch
func (r *registerEefdr2Type) SetEe9ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefdr2FieldEe9ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefdr2FieldEe9ltchMask)
	}
}

const (
	RegisterEefdr2FieldEe8fltrShift = 13
	RegisterEefdr2FieldEe8fltrMask  = 0x1e000
)

// GetEe8fltr External Event 8 filter
func (r *registerEefdr2Type) GetEe8fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefdr2FieldEe8fltrMask) >> RegisterEefdr2FieldEe8fltrShift)
}

// SetEe8fltr External Event 8 filter
func (r *registerEefdr2Type) SetEe8fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefdr2FieldEe8fltrMask)|(uint32(value)<<RegisterEefdr2FieldEe8fltrShift))
}

const (
	RegisterEefdr2FieldEe8ltchShift = 12
	RegisterEefdr2FieldEe8ltchMask  = 0x1000
)

// GetEe8ltch External Event 8 latch
func (r *registerEefdr2Type) GetEe8ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefdr2FieldEe8ltchMask) != 0
}

// SetEe8ltch External Event 8 latch
func (r *registerEefdr2Type) SetEe8ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefdr2FieldEe8ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefdr2FieldEe8ltchMask)
	}
}

const (
	RegisterEefdr2FieldEe7fltrShift = 7
	RegisterEefdr2FieldEe7fltrMask  = 0x780
)

// GetEe7fltr External Event 7 filter
func (r *registerEefdr2Type) GetEe7fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefdr2FieldEe7fltrMask) >> RegisterEefdr2FieldEe7fltrShift)
}

// SetEe7fltr External Event 7 filter
func (r *registerEefdr2Type) SetEe7fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefdr2FieldEe7fltrMask)|(uint32(value)<<RegisterEefdr2FieldEe7fltrShift))
}

const (
	RegisterEefdr2FieldEe7ltchShift = 6
	RegisterEefdr2FieldEe7ltchMask  = 0x40
)

// GetEe7ltch External Event 7 latch
func (r *registerEefdr2Type) GetEe7ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefdr2FieldEe7ltchMask) != 0
}

// SetEe7ltch External Event 7 latch
func (r *registerEefdr2Type) SetEe7ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefdr2FieldEe7ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefdr2FieldEe7ltchMask)
	}
}

const (
	RegisterEefdr2FieldEe6fltrShift = 1
	RegisterEefdr2FieldEe6fltrMask  = 0x1e
)

// GetEe6fltr External Event 6 filter
func (r *registerEefdr2Type) GetEe6fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefdr2FieldEe6fltrMask) >> RegisterEefdr2FieldEe6fltrShift)
}

// SetEe6fltr External Event 6 filter
func (r *registerEefdr2Type) SetEe6fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefdr2FieldEe6fltrMask)|(uint32(value)<<RegisterEefdr2FieldEe6fltrShift))
}

const (
	RegisterEefdr2FieldEe6ltchShift = 0
	RegisterEefdr2FieldEe6ltchMask  = 0x1
)

// GetEe6ltch External Event 6 latch
func (r *registerEefdr2Type) GetEe6ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefdr2FieldEe6ltchMask) != 0
}

// SetEe6ltch External Event 6 latch
func (r *registerEefdr2Type) SetEe6ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefdr2FieldEe6ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefdr2FieldEe6ltchMask)
	}
}

// registerRstdrType TimerA Reset Register
type registerRstdrType uint32

const (
	RegisterRstdrFieldTimecmp4Shift = 30
	RegisterRstdrFieldTimecmp4Mask  = 0x40000000
)

// GetTimecmp4 Timer E Compare 4
func (r *registerRstdrType) GetTimecmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldTimecmp4Mask) != 0
}

// SetTimecmp4 Timer E Compare 4
func (r *registerRstdrType) SetTimecmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldTimecmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldTimecmp4Mask)
	}
}

const (
	RegisterRstdrFieldTimecmp2Shift = 29
	RegisterRstdrFieldTimecmp2Mask  = 0x20000000
)

// GetTimecmp2 Timer E Compare 2
func (r *registerRstdrType) GetTimecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldTimecmp2Mask) != 0
}

// SetTimecmp2 Timer E Compare 2
func (r *registerRstdrType) SetTimecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldTimecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldTimecmp2Mask)
	}
}

const (
	RegisterRstdrFieldTimecmp1Shift = 28
	RegisterRstdrFieldTimecmp1Mask  = 0x10000000
)

// GetTimecmp1 Timer E Compare 1
func (r *registerRstdrType) GetTimecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldTimecmp1Mask) != 0
}

// SetTimecmp1 Timer E Compare 1
func (r *registerRstdrType) SetTimecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldTimecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldTimecmp1Mask)
	}
}

const (
	RegisterRstdrFieldTimccmp4Shift = 27
	RegisterRstdrFieldTimccmp4Mask  = 0x8000000
)

// GetTimccmp4 Timer C Compare 4
func (r *registerRstdrType) GetTimccmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldTimccmp4Mask) != 0
}

// SetTimccmp4 Timer C Compare 4
func (r *registerRstdrType) SetTimccmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldTimccmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldTimccmp4Mask)
	}
}

const (
	RegisterRstdrFieldTimccmp2Shift = 26
	RegisterRstdrFieldTimccmp2Mask  = 0x4000000
)

// GetTimccmp2 Timer C Compare 2
func (r *registerRstdrType) GetTimccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldTimccmp2Mask) != 0
}

// SetTimccmp2 Timer C Compare 2
func (r *registerRstdrType) SetTimccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldTimccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldTimccmp2Mask)
	}
}

const (
	RegisterRstdrFieldTimccmp1Shift = 25
	RegisterRstdrFieldTimccmp1Mask  = 0x2000000
)

// GetTimccmp1 Timer C Compare 1
func (r *registerRstdrType) GetTimccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldTimccmp1Mask) != 0
}

// SetTimccmp1 Timer C Compare 1
func (r *registerRstdrType) SetTimccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldTimccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldTimccmp1Mask)
	}
}

const (
	RegisterRstdrFieldTimbcmp4Shift = 24
	RegisterRstdrFieldTimbcmp4Mask  = 0x1000000
)

// GetTimbcmp4 Timer B Compare 4
func (r *registerRstdrType) GetTimbcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldTimbcmp4Mask) != 0
}

// SetTimbcmp4 Timer B Compare 4
func (r *registerRstdrType) SetTimbcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldTimbcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldTimbcmp4Mask)
	}
}

const (
	RegisterRstdrFieldTimbcmp2Shift = 23
	RegisterRstdrFieldTimbcmp2Mask  = 0x800000
)

// GetTimbcmp2 Timer B Compare 2
func (r *registerRstdrType) GetTimbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldTimbcmp2Mask) != 0
}

// SetTimbcmp2 Timer B Compare 2
func (r *registerRstdrType) SetTimbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldTimbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldTimbcmp2Mask)
	}
}

const (
	RegisterRstdrFieldTimbcmp1Shift = 22
	RegisterRstdrFieldTimbcmp1Mask  = 0x400000
)

// GetTimbcmp1 Timer B Compare 1
func (r *registerRstdrType) GetTimbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldTimbcmp1Mask) != 0
}

// SetTimbcmp1 Timer B Compare 1
func (r *registerRstdrType) SetTimbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldTimbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldTimbcmp1Mask)
	}
}

const (
	RegisterRstdrFieldTimacmp4Shift = 21
	RegisterRstdrFieldTimacmp4Mask  = 0x200000
)

// GetTimacmp4 Timer A Compare 4
func (r *registerRstdrType) GetTimacmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldTimacmp4Mask) != 0
}

// SetTimacmp4 Timer A Compare 4
func (r *registerRstdrType) SetTimacmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldTimacmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldTimacmp4Mask)
	}
}

const (
	RegisterRstdrFieldTimacmp2Shift = 20
	RegisterRstdrFieldTimacmp2Mask  = 0x100000
)

// GetTimacmp2 Timer A Compare 2
func (r *registerRstdrType) GetTimacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldTimacmp2Mask) != 0
}

// SetTimacmp2 Timer A Compare 2
func (r *registerRstdrType) SetTimacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldTimacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldTimacmp2Mask)
	}
}

const (
	RegisterRstdrFieldTimacmp1Shift = 19
	RegisterRstdrFieldTimacmp1Mask  = 0x80000
)

// GetTimacmp1 Timer A Compare 1
func (r *registerRstdrType) GetTimacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldTimacmp1Mask) != 0
}

// SetTimacmp1 Timer A Compare 1
func (r *registerRstdrType) SetTimacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldTimacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldTimacmp1Mask)
	}
}

const (
	RegisterRstdrFieldExtevnt10Shift = 18
	RegisterRstdrFieldExtevnt10Mask  = 0x40000
)

// GetExtevnt10 External Event 10
func (r *registerRstdrType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldExtevnt10Mask) != 0
}

// SetExtevnt10 External Event 10
func (r *registerRstdrType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldExtevnt10Mask)
	}
}

const (
	RegisterRstdrFieldExtevnt9Shift = 17
	RegisterRstdrFieldExtevnt9Mask  = 0x20000
)

// GetExtevnt9 External Event 9
func (r *registerRstdrType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldExtevnt9Mask) != 0
}

// SetExtevnt9 External Event 9
func (r *registerRstdrType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldExtevnt9Mask)
	}
}

const (
	RegisterRstdrFieldExtevnt8Shift = 16
	RegisterRstdrFieldExtevnt8Mask  = 0x10000
)

// GetExtevnt8 External Event 8
func (r *registerRstdrType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldExtevnt8Mask) != 0
}

// SetExtevnt8 External Event 8
func (r *registerRstdrType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldExtevnt8Mask)
	}
}

const (
	RegisterRstdrFieldExtevnt7Shift = 15
	RegisterRstdrFieldExtevnt7Mask  = 0x8000
)

// GetExtevnt7 External Event 7
func (r *registerRstdrType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldExtevnt7Mask) != 0
}

// SetExtevnt7 External Event 7
func (r *registerRstdrType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldExtevnt7Mask)
	}
}

const (
	RegisterRstdrFieldExtevnt6Shift = 14
	RegisterRstdrFieldExtevnt6Mask  = 0x4000
)

// GetExtevnt6 External Event 6
func (r *registerRstdrType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldExtevnt6Mask) != 0
}

// SetExtevnt6 External Event 6
func (r *registerRstdrType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldExtevnt6Mask)
	}
}

const (
	RegisterRstdrFieldExtevnt5Shift = 13
	RegisterRstdrFieldExtevnt5Mask  = 0x2000
)

// GetExtevnt5 External Event 5
func (r *registerRstdrType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldExtevnt5Mask) != 0
}

// SetExtevnt5 External Event 5
func (r *registerRstdrType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldExtevnt5Mask)
	}
}

const (
	RegisterRstdrFieldExtevnt4Shift = 12
	RegisterRstdrFieldExtevnt4Mask  = 0x1000
)

// GetExtevnt4 External Event 4
func (r *registerRstdrType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldExtevnt4Mask) != 0
}

// SetExtevnt4 External Event 4
func (r *registerRstdrType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldExtevnt4Mask)
	}
}

const (
	RegisterRstdrFieldExtevnt3Shift = 11
	RegisterRstdrFieldExtevnt3Mask  = 0x800
)

// GetExtevnt3 External Event 3
func (r *registerRstdrType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldExtevnt3Mask) != 0
}

// SetExtevnt3 External Event 3
func (r *registerRstdrType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldExtevnt3Mask)
	}
}

const (
	RegisterRstdrFieldExtevnt2Shift = 10
	RegisterRstdrFieldExtevnt2Mask  = 0x400
)

// GetExtevnt2 External Event 2
func (r *registerRstdrType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldExtevnt2Mask) != 0
}

// SetExtevnt2 External Event 2
func (r *registerRstdrType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldExtevnt2Mask)
	}
}

const (
	RegisterRstdrFieldExtevnt1Shift = 9
	RegisterRstdrFieldExtevnt1Mask  = 0x200
)

// GetExtevnt1 External Event 1
func (r *registerRstdrType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldExtevnt1Mask) != 0
}

// SetExtevnt1 External Event 1
func (r *registerRstdrType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldExtevnt1Mask)
	}
}

const (
	RegisterRstdrFieldMstcmp4Shift = 8
	RegisterRstdrFieldMstcmp4Mask  = 0x100
)

// GetMstcmp4 Master compare 4
func (r *registerRstdrType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldMstcmp4Mask) != 0
}

// SetMstcmp4 Master compare 4
func (r *registerRstdrType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldMstcmp4Mask)
	}
}

const (
	RegisterRstdrFieldMstcmp3Shift = 7
	RegisterRstdrFieldMstcmp3Mask  = 0x80
)

// GetMstcmp3 Master compare 3
func (r *registerRstdrType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldMstcmp3Mask) != 0
}

// SetMstcmp3 Master compare 3
func (r *registerRstdrType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldMstcmp3Mask)
	}
}

const (
	RegisterRstdrFieldMstcmp2Shift = 6
	RegisterRstdrFieldMstcmp2Mask  = 0x40
)

// GetMstcmp2 Master compare 2
func (r *registerRstdrType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldMstcmp2Mask) != 0
}

// SetMstcmp2 Master compare 2
func (r *registerRstdrType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldMstcmp2Mask)
	}
}

const (
	RegisterRstdrFieldMstcmp1Shift = 5
	RegisterRstdrFieldMstcmp1Mask  = 0x20
)

// GetMstcmp1 Master compare 1
func (r *registerRstdrType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldMstcmp1Mask) != 0
}

// SetMstcmp1 Master compare 1
func (r *registerRstdrType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldMstcmp1Mask)
	}
}

const (
	RegisterRstdrFieldMstperShift = 4
	RegisterRstdrFieldMstperMask  = 0x10
)

// GetMstper Master timer Period
func (r *registerRstdrType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldMstperMask) != 0
}

// SetMstper Master timer Period
func (r *registerRstdrType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldMstperMask)
	}
}

const (
	RegisterRstdrFieldCmp4Shift = 3
	RegisterRstdrFieldCmp4Mask  = 0x8
)

// GetCmp4 Timer A compare 4 reset
func (r *registerRstdrType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldCmp4Mask) != 0
}

// SetCmp4 Timer A compare 4 reset
func (r *registerRstdrType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldCmp4Mask)
	}
}

const (
	RegisterRstdrFieldCmp2Shift = 2
	RegisterRstdrFieldCmp2Mask  = 0x4
)

// GetCmp2 Timer A compare 2 reset
func (r *registerRstdrType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldCmp2Mask) != 0
}

// SetCmp2 Timer A compare 2 reset
func (r *registerRstdrType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldCmp2Mask)
	}
}

const (
	RegisterRstdrFieldUpdtShift = 1
	RegisterRstdrFieldUpdtMask  = 0x2
)

// GetUpdt Timer A Update reset
func (r *registerRstdrType) GetUpdt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstdrFieldUpdtMask) != 0
}

// SetUpdt Timer A Update reset
func (r *registerRstdrType) SetUpdt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstdrFieldUpdtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstdrFieldUpdtMask)
	}
}

// registerChpdrType Timerx Chopper Register
type registerChpdrType uint32

const (
	RegisterChpdrFieldStrtpwShift = 7
	RegisterChpdrFieldStrtpwMask  = 0x780
)

// GetStrtpw STRTPW
func (r *registerChpdrType) GetStrtpw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChpdrFieldStrtpwMask) >> RegisterChpdrFieldStrtpwShift)
}

// SetStrtpw STRTPW
func (r *registerChpdrType) SetStrtpw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChpdrFieldStrtpwMask)|(uint32(value)<<RegisterChpdrFieldStrtpwShift))
}

const (
	RegisterChpdrFieldChpdtyShift = 4
	RegisterChpdrFieldChpdtyMask  = 0x70
)

// GetChpdty Timerx chopper duty cycle value
func (r *registerChpdrType) GetChpdty() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChpdrFieldChpdtyMask) >> RegisterChpdrFieldChpdtyShift)
}

// SetChpdty Timerx chopper duty cycle value
func (r *registerChpdrType) SetChpdty(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChpdrFieldChpdtyMask)|(uint32(value)<<RegisterChpdrFieldChpdtyShift))
}

const (
	RegisterChpdrFieldChpfrqShift = 0
	RegisterChpdrFieldChpfrqMask  = 0xf
)

// GetChpfrq Timerx carrier frequency value
func (r *registerChpdrType) GetChpfrq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChpdrFieldChpfrqMask) >> RegisterChpdrFieldChpfrqShift)
}

// SetChpfrq Timerx carrier frequency value
func (r *registerChpdrType) SetChpfrq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChpdrFieldChpfrqMask)|(uint32(value)<<RegisterChpdrFieldChpfrqShift))
}

// registerCpt1dcrType Timerx Capture 2 Control Register
type registerCpt1dcrType uint32

const (
	RegisterCpt1dcrFieldTecmp2Shift = 31
	RegisterCpt1dcrFieldTecmp2Mask  = 0x80000000
)

// GetTecmp2 Timer E Compare 2
func (r *registerCpt1dcrType) GetTecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTecmp2Mask) != 0
}

// SetTecmp2 Timer E Compare 2
func (r *registerCpt1dcrType) SetTecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTecmp2Mask)
	}
}

const (
	RegisterCpt1dcrFieldTecmp1Shift = 30
	RegisterCpt1dcrFieldTecmp1Mask  = 0x40000000
)

// GetTecmp1 Timer E Compare 1
func (r *registerCpt1dcrType) GetTecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTecmp1Mask) != 0
}

// SetTecmp1 Timer E Compare 1
func (r *registerCpt1dcrType) SetTecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTecmp1Mask)
	}
}

const (
	RegisterCpt1dcrFieldTe1rstShift = 29
	RegisterCpt1dcrFieldTe1rstMask  = 0x20000000
)

// GetTe1rst Timer E output 1 Reset
func (r *registerCpt1dcrType) GetTe1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTe1rstMask) != 0
}

// SetTe1rst Timer E output 1 Reset
func (r *registerCpt1dcrType) SetTe1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTe1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTe1rstMask)
	}
}

const (
	RegisterCpt1dcrFieldTe1setShift = 28
	RegisterCpt1dcrFieldTe1setMask  = 0x10000000
)

// GetTe1set Timer E output 1 Set
func (r *registerCpt1dcrType) GetTe1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTe1setMask) != 0
}

// SetTe1set Timer E output 1 Set
func (r *registerCpt1dcrType) SetTe1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTe1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTe1setMask)
	}
}

const (
	RegisterCpt1dcrFieldTccmp2Shift = 23
	RegisterCpt1dcrFieldTccmp2Mask  = 0x800000
)

// GetTccmp2 Timer C Compare 2
func (r *registerCpt1dcrType) GetTccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTccmp2Mask) != 0
}

// SetTccmp2 Timer C Compare 2
func (r *registerCpt1dcrType) SetTccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTccmp2Mask)
	}
}

const (
	RegisterCpt1dcrFieldTccmp1Shift = 22
	RegisterCpt1dcrFieldTccmp1Mask  = 0x400000
)

// GetTccmp1 Timer C Compare 1
func (r *registerCpt1dcrType) GetTccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTccmp1Mask) != 0
}

// SetTccmp1 Timer C Compare 1
func (r *registerCpt1dcrType) SetTccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTccmp1Mask)
	}
}

const (
	RegisterCpt1dcrFieldTc1rstShift = 21
	RegisterCpt1dcrFieldTc1rstMask  = 0x200000
)

// GetTc1rst Timer C output 1 Reset
func (r *registerCpt1dcrType) GetTc1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTc1rstMask) != 0
}

// SetTc1rst Timer C output 1 Reset
func (r *registerCpt1dcrType) SetTc1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTc1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTc1rstMask)
	}
}

const (
	RegisterCpt1dcrFieldTc1setShift = 20
	RegisterCpt1dcrFieldTc1setMask  = 0x100000
)

// GetTc1set Timer C output 1 Set
func (r *registerCpt1dcrType) GetTc1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTc1setMask) != 0
}

// SetTc1set Timer C output 1 Set
func (r *registerCpt1dcrType) SetTc1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTc1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTc1setMask)
	}
}

const (
	RegisterCpt1dcrFieldTbcmp2Shift = 19
	RegisterCpt1dcrFieldTbcmp2Mask  = 0x80000
)

// GetTbcmp2 Timer B Compare 2
func (r *registerCpt1dcrType) GetTbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTbcmp2Mask) != 0
}

// SetTbcmp2 Timer B Compare 2
func (r *registerCpt1dcrType) SetTbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTbcmp2Mask)
	}
}

const (
	RegisterCpt1dcrFieldTbcmp1Shift = 18
	RegisterCpt1dcrFieldTbcmp1Mask  = 0x40000
)

// GetTbcmp1 Timer B Compare 1
func (r *registerCpt1dcrType) GetTbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTbcmp1Mask) != 0
}

// SetTbcmp1 Timer B Compare 1
func (r *registerCpt1dcrType) SetTbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTbcmp1Mask)
	}
}

const (
	RegisterCpt1dcrFieldTb1rstShift = 17
	RegisterCpt1dcrFieldTb1rstMask  = 0x20000
)

// GetTb1rst Timer B output 1 Reset
func (r *registerCpt1dcrType) GetTb1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTb1rstMask) != 0
}

// SetTb1rst Timer B output 1 Reset
func (r *registerCpt1dcrType) SetTb1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTb1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTb1rstMask)
	}
}

const (
	RegisterCpt1dcrFieldTb1setShift = 16
	RegisterCpt1dcrFieldTb1setMask  = 0x10000
)

// GetTb1set Timer B output 1 Set
func (r *registerCpt1dcrType) GetTb1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTb1setMask) != 0
}

// SetTb1set Timer B output 1 Set
func (r *registerCpt1dcrType) SetTb1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTb1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTb1setMask)
	}
}

const (
	RegisterCpt1dcrFieldTacmp2Shift = 15
	RegisterCpt1dcrFieldTacmp2Mask  = 0x8000
)

// GetTacmp2 Timer A Compare 2
func (r *registerCpt1dcrType) GetTacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTacmp2Mask) != 0
}

// SetTacmp2 Timer A Compare 2
func (r *registerCpt1dcrType) SetTacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTacmp2Mask)
	}
}

const (
	RegisterCpt1dcrFieldTacmp1Shift = 14
	RegisterCpt1dcrFieldTacmp1Mask  = 0x4000
)

// GetTacmp1 Timer A Compare 1
func (r *registerCpt1dcrType) GetTacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTacmp1Mask) != 0
}

// SetTacmp1 Timer A Compare 1
func (r *registerCpt1dcrType) SetTacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTacmp1Mask)
	}
}

const (
	RegisterCpt1dcrFieldTa1rstShift = 13
	RegisterCpt1dcrFieldTa1rstMask  = 0x2000
)

// GetTa1rst Timer A output 1 Reset
func (r *registerCpt1dcrType) GetTa1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTa1rstMask) != 0
}

// SetTa1rst Timer A output 1 Reset
func (r *registerCpt1dcrType) SetTa1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTa1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTa1rstMask)
	}
}

const (
	RegisterCpt1dcrFieldTa1setShift = 12
	RegisterCpt1dcrFieldTa1setMask  = 0x1000
)

// GetTa1set Timer A output 1 Set
func (r *registerCpt1dcrType) GetTa1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldTa1setMask) != 0
}

// SetTa1set Timer A output 1 Set
func (r *registerCpt1dcrType) SetTa1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldTa1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldTa1setMask)
	}
}

const (
	RegisterCpt1dcrFieldExev10cptShift = 11
	RegisterCpt1dcrFieldExev10cptMask  = 0x800
)

// GetExev10cpt External Event 10 Capture
func (r *registerCpt1dcrType) GetExev10cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldExev10cptMask) != 0
}

// SetExev10cpt External Event 10 Capture
func (r *registerCpt1dcrType) SetExev10cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldExev10cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldExev10cptMask)
	}
}

const (
	RegisterCpt1dcrFieldExev9cptShift = 10
	RegisterCpt1dcrFieldExev9cptMask  = 0x400
)

// GetExev9cpt External Event 9 Capture
func (r *registerCpt1dcrType) GetExev9cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldExev9cptMask) != 0
}

// SetExev9cpt External Event 9 Capture
func (r *registerCpt1dcrType) SetExev9cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldExev9cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldExev9cptMask)
	}
}

const (
	RegisterCpt1dcrFieldExev8cptShift = 9
	RegisterCpt1dcrFieldExev8cptMask  = 0x200
)

// GetExev8cpt External Event 8 Capture
func (r *registerCpt1dcrType) GetExev8cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldExev8cptMask) != 0
}

// SetExev8cpt External Event 8 Capture
func (r *registerCpt1dcrType) SetExev8cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldExev8cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldExev8cptMask)
	}
}

const (
	RegisterCpt1dcrFieldExev7cptShift = 8
	RegisterCpt1dcrFieldExev7cptMask  = 0x100
)

// GetExev7cpt External Event 7 Capture
func (r *registerCpt1dcrType) GetExev7cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldExev7cptMask) != 0
}

// SetExev7cpt External Event 7 Capture
func (r *registerCpt1dcrType) SetExev7cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldExev7cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldExev7cptMask)
	}
}

const (
	RegisterCpt1dcrFieldExev6cptShift = 7
	RegisterCpt1dcrFieldExev6cptMask  = 0x80
)

// GetExev6cpt External Event 6 Capture
func (r *registerCpt1dcrType) GetExev6cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldExev6cptMask) != 0
}

// SetExev6cpt External Event 6 Capture
func (r *registerCpt1dcrType) SetExev6cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldExev6cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldExev6cptMask)
	}
}

const (
	RegisterCpt1dcrFieldExev5cptShift = 6
	RegisterCpt1dcrFieldExev5cptMask  = 0x40
)

// GetExev5cpt External Event 5 Capture
func (r *registerCpt1dcrType) GetExev5cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldExev5cptMask) != 0
}

// SetExev5cpt External Event 5 Capture
func (r *registerCpt1dcrType) SetExev5cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldExev5cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldExev5cptMask)
	}
}

const (
	RegisterCpt1dcrFieldExev4cptShift = 5
	RegisterCpt1dcrFieldExev4cptMask  = 0x20
)

// GetExev4cpt External Event 4 Capture
func (r *registerCpt1dcrType) GetExev4cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldExev4cptMask) != 0
}

// SetExev4cpt External Event 4 Capture
func (r *registerCpt1dcrType) SetExev4cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldExev4cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldExev4cptMask)
	}
}

const (
	RegisterCpt1dcrFieldExev3cptShift = 4
	RegisterCpt1dcrFieldExev3cptMask  = 0x10
)

// GetExev3cpt External Event 3 Capture
func (r *registerCpt1dcrType) GetExev3cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldExev3cptMask) != 0
}

// SetExev3cpt External Event 3 Capture
func (r *registerCpt1dcrType) SetExev3cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldExev3cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldExev3cptMask)
	}
}

const (
	RegisterCpt1dcrFieldExev2cptShift = 3
	RegisterCpt1dcrFieldExev2cptMask  = 0x8
)

// GetExev2cpt External Event 2 Capture
func (r *registerCpt1dcrType) GetExev2cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldExev2cptMask) != 0
}

// SetExev2cpt External Event 2 Capture
func (r *registerCpt1dcrType) SetExev2cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldExev2cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldExev2cptMask)
	}
}

const (
	RegisterCpt1dcrFieldExev1cptShift = 2
	RegisterCpt1dcrFieldExev1cptMask  = 0x4
)

// GetExev1cpt External Event 1 Capture
func (r *registerCpt1dcrType) GetExev1cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldExev1cptMask) != 0
}

// SetExev1cpt External Event 1 Capture
func (r *registerCpt1dcrType) SetExev1cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldExev1cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldExev1cptMask)
	}
}

const (
	RegisterCpt1dcrFieldUdpcptShift = 1
	RegisterCpt1dcrFieldUdpcptMask  = 0x2
)

// GetUdpcpt Update Capture
func (r *registerCpt1dcrType) GetUdpcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldUdpcptMask) != 0
}

// SetUdpcpt Update Capture
func (r *registerCpt1dcrType) SetUdpcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldUdpcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldUdpcptMask)
	}
}

const (
	RegisterCpt1dcrFieldSwcptShift = 0
	RegisterCpt1dcrFieldSwcptMask  = 0x1
)

// GetSwcpt Software Capture
func (r *registerCpt1dcrType) GetSwcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1dcrFieldSwcptMask) != 0
}

// SetSwcpt Software Capture
func (r *registerCpt1dcrType) SetSwcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1dcrFieldSwcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1dcrFieldSwcptMask)
	}
}

// registerCpt2dcrType CPT2xCR
type registerCpt2dcrType uint32

const (
	RegisterCpt2dcrFieldTecmp2Shift = 31
	RegisterCpt2dcrFieldTecmp2Mask  = 0x80000000
)

// GetTecmp2 Timer E Compare 2
func (r *registerCpt2dcrType) GetTecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTecmp2Mask) != 0
}

// SetTecmp2 Timer E Compare 2
func (r *registerCpt2dcrType) SetTecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTecmp2Mask)
	}
}

const (
	RegisterCpt2dcrFieldTecmp1Shift = 30
	RegisterCpt2dcrFieldTecmp1Mask  = 0x40000000
)

// GetTecmp1 Timer E Compare 1
func (r *registerCpt2dcrType) GetTecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTecmp1Mask) != 0
}

// SetTecmp1 Timer E Compare 1
func (r *registerCpt2dcrType) SetTecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTecmp1Mask)
	}
}

const (
	RegisterCpt2dcrFieldTe1rstShift = 29
	RegisterCpt2dcrFieldTe1rstMask  = 0x20000000
)

// GetTe1rst Timer E output 1 Reset
func (r *registerCpt2dcrType) GetTe1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTe1rstMask) != 0
}

// SetTe1rst Timer E output 1 Reset
func (r *registerCpt2dcrType) SetTe1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTe1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTe1rstMask)
	}
}

const (
	RegisterCpt2dcrFieldTe1setShift = 28
	RegisterCpt2dcrFieldTe1setMask  = 0x10000000
)

// GetTe1set Timer E output 1 Set
func (r *registerCpt2dcrType) GetTe1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTe1setMask) != 0
}

// SetTe1set Timer E output 1 Set
func (r *registerCpt2dcrType) SetTe1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTe1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTe1setMask)
	}
}

const (
	RegisterCpt2dcrFieldTccmp2Shift = 23
	RegisterCpt2dcrFieldTccmp2Mask  = 0x800000
)

// GetTccmp2 Timer C Compare 2
func (r *registerCpt2dcrType) GetTccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTccmp2Mask) != 0
}

// SetTccmp2 Timer C Compare 2
func (r *registerCpt2dcrType) SetTccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTccmp2Mask)
	}
}

const (
	RegisterCpt2dcrFieldTccmp1Shift = 22
	RegisterCpt2dcrFieldTccmp1Mask  = 0x400000
)

// GetTccmp1 Timer C Compare 1
func (r *registerCpt2dcrType) GetTccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTccmp1Mask) != 0
}

// SetTccmp1 Timer C Compare 1
func (r *registerCpt2dcrType) SetTccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTccmp1Mask)
	}
}

const (
	RegisterCpt2dcrFieldTc1rstShift = 21
	RegisterCpt2dcrFieldTc1rstMask  = 0x200000
)

// GetTc1rst Timer C output 1 Reset
func (r *registerCpt2dcrType) GetTc1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTc1rstMask) != 0
}

// SetTc1rst Timer C output 1 Reset
func (r *registerCpt2dcrType) SetTc1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTc1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTc1rstMask)
	}
}

const (
	RegisterCpt2dcrFieldTc1setShift = 20
	RegisterCpt2dcrFieldTc1setMask  = 0x100000
)

// GetTc1set Timer C output 1 Set
func (r *registerCpt2dcrType) GetTc1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTc1setMask) != 0
}

// SetTc1set Timer C output 1 Set
func (r *registerCpt2dcrType) SetTc1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTc1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTc1setMask)
	}
}

const (
	RegisterCpt2dcrFieldTbcmp2Shift = 19
	RegisterCpt2dcrFieldTbcmp2Mask  = 0x80000
)

// GetTbcmp2 Timer B Compare 2
func (r *registerCpt2dcrType) GetTbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTbcmp2Mask) != 0
}

// SetTbcmp2 Timer B Compare 2
func (r *registerCpt2dcrType) SetTbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTbcmp2Mask)
	}
}

const (
	RegisterCpt2dcrFieldTbcmp1Shift = 18
	RegisterCpt2dcrFieldTbcmp1Mask  = 0x40000
)

// GetTbcmp1 Timer B Compare 1
func (r *registerCpt2dcrType) GetTbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTbcmp1Mask) != 0
}

// SetTbcmp1 Timer B Compare 1
func (r *registerCpt2dcrType) SetTbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTbcmp1Mask)
	}
}

const (
	RegisterCpt2dcrFieldTb1rstShift = 17
	RegisterCpt2dcrFieldTb1rstMask  = 0x20000
)

// GetTb1rst Timer B output 1 Reset
func (r *registerCpt2dcrType) GetTb1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTb1rstMask) != 0
}

// SetTb1rst Timer B output 1 Reset
func (r *registerCpt2dcrType) SetTb1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTb1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTb1rstMask)
	}
}

const (
	RegisterCpt2dcrFieldTb1setShift = 16
	RegisterCpt2dcrFieldTb1setMask  = 0x10000
)

// GetTb1set Timer B output 1 Set
func (r *registerCpt2dcrType) GetTb1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTb1setMask) != 0
}

// SetTb1set Timer B output 1 Set
func (r *registerCpt2dcrType) SetTb1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTb1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTb1setMask)
	}
}

const (
	RegisterCpt2dcrFieldTacmp2Shift = 15
	RegisterCpt2dcrFieldTacmp2Mask  = 0x8000
)

// GetTacmp2 Timer A Compare 2
func (r *registerCpt2dcrType) GetTacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTacmp2Mask) != 0
}

// SetTacmp2 Timer A Compare 2
func (r *registerCpt2dcrType) SetTacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTacmp2Mask)
	}
}

const (
	RegisterCpt2dcrFieldTacmp1Shift = 14
	RegisterCpt2dcrFieldTacmp1Mask  = 0x4000
)

// GetTacmp1 Timer A Compare 1
func (r *registerCpt2dcrType) GetTacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTacmp1Mask) != 0
}

// SetTacmp1 Timer A Compare 1
func (r *registerCpt2dcrType) SetTacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTacmp1Mask)
	}
}

const (
	RegisterCpt2dcrFieldTa1rstShift = 13
	RegisterCpt2dcrFieldTa1rstMask  = 0x2000
)

// GetTa1rst Timer A output 1 Reset
func (r *registerCpt2dcrType) GetTa1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTa1rstMask) != 0
}

// SetTa1rst Timer A output 1 Reset
func (r *registerCpt2dcrType) SetTa1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTa1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTa1rstMask)
	}
}

const (
	RegisterCpt2dcrFieldTa1setShift = 12
	RegisterCpt2dcrFieldTa1setMask  = 0x1000
)

// GetTa1set Timer A output 1 Set
func (r *registerCpt2dcrType) GetTa1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldTa1setMask) != 0
}

// SetTa1set Timer A output 1 Set
func (r *registerCpt2dcrType) SetTa1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldTa1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldTa1setMask)
	}
}

const (
	RegisterCpt2dcrFieldExev10cptShift = 11
	RegisterCpt2dcrFieldExev10cptMask  = 0x800
)

// GetExev10cpt External Event 10 Capture
func (r *registerCpt2dcrType) GetExev10cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldExev10cptMask) != 0
}

// SetExev10cpt External Event 10 Capture
func (r *registerCpt2dcrType) SetExev10cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldExev10cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldExev10cptMask)
	}
}

const (
	RegisterCpt2dcrFieldExev9cptShift = 10
	RegisterCpt2dcrFieldExev9cptMask  = 0x400
)

// GetExev9cpt External Event 9 Capture
func (r *registerCpt2dcrType) GetExev9cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldExev9cptMask) != 0
}

// SetExev9cpt External Event 9 Capture
func (r *registerCpt2dcrType) SetExev9cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldExev9cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldExev9cptMask)
	}
}

const (
	RegisterCpt2dcrFieldExev8cptShift = 9
	RegisterCpt2dcrFieldExev8cptMask  = 0x200
)

// GetExev8cpt External Event 8 Capture
func (r *registerCpt2dcrType) GetExev8cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldExev8cptMask) != 0
}

// SetExev8cpt External Event 8 Capture
func (r *registerCpt2dcrType) SetExev8cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldExev8cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldExev8cptMask)
	}
}

const (
	RegisterCpt2dcrFieldExev7cptShift = 8
	RegisterCpt2dcrFieldExev7cptMask  = 0x100
)

// GetExev7cpt External Event 7 Capture
func (r *registerCpt2dcrType) GetExev7cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldExev7cptMask) != 0
}

// SetExev7cpt External Event 7 Capture
func (r *registerCpt2dcrType) SetExev7cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldExev7cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldExev7cptMask)
	}
}

const (
	RegisterCpt2dcrFieldExev6cptShift = 7
	RegisterCpt2dcrFieldExev6cptMask  = 0x80
)

// GetExev6cpt External Event 6 Capture
func (r *registerCpt2dcrType) GetExev6cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldExev6cptMask) != 0
}

// SetExev6cpt External Event 6 Capture
func (r *registerCpt2dcrType) SetExev6cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldExev6cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldExev6cptMask)
	}
}

const (
	RegisterCpt2dcrFieldExev5cptShift = 6
	RegisterCpt2dcrFieldExev5cptMask  = 0x40
)

// GetExev5cpt External Event 5 Capture
func (r *registerCpt2dcrType) GetExev5cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldExev5cptMask) != 0
}

// SetExev5cpt External Event 5 Capture
func (r *registerCpt2dcrType) SetExev5cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldExev5cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldExev5cptMask)
	}
}

const (
	RegisterCpt2dcrFieldExev4cptShift = 5
	RegisterCpt2dcrFieldExev4cptMask  = 0x20
)

// GetExev4cpt External Event 4 Capture
func (r *registerCpt2dcrType) GetExev4cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldExev4cptMask) != 0
}

// SetExev4cpt External Event 4 Capture
func (r *registerCpt2dcrType) SetExev4cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldExev4cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldExev4cptMask)
	}
}

const (
	RegisterCpt2dcrFieldExev3cptShift = 4
	RegisterCpt2dcrFieldExev3cptMask  = 0x10
)

// GetExev3cpt External Event 3 Capture
func (r *registerCpt2dcrType) GetExev3cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldExev3cptMask) != 0
}

// SetExev3cpt External Event 3 Capture
func (r *registerCpt2dcrType) SetExev3cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldExev3cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldExev3cptMask)
	}
}

const (
	RegisterCpt2dcrFieldExev2cptShift = 3
	RegisterCpt2dcrFieldExev2cptMask  = 0x8
)

// GetExev2cpt External Event 2 Capture
func (r *registerCpt2dcrType) GetExev2cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldExev2cptMask) != 0
}

// SetExev2cpt External Event 2 Capture
func (r *registerCpt2dcrType) SetExev2cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldExev2cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldExev2cptMask)
	}
}

const (
	RegisterCpt2dcrFieldExev1cptShift = 2
	RegisterCpt2dcrFieldExev1cptMask  = 0x4
)

// GetExev1cpt External Event 1 Capture
func (r *registerCpt2dcrType) GetExev1cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldExev1cptMask) != 0
}

// SetExev1cpt External Event 1 Capture
func (r *registerCpt2dcrType) SetExev1cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldExev1cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldExev1cptMask)
	}
}

const (
	RegisterCpt2dcrFieldUdpcptShift = 1
	RegisterCpt2dcrFieldUdpcptMask  = 0x2
)

// GetUdpcpt Update Capture
func (r *registerCpt2dcrType) GetUdpcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldUdpcptMask) != 0
}

// SetUdpcpt Update Capture
func (r *registerCpt2dcrType) SetUdpcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldUdpcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldUdpcptMask)
	}
}

const (
	RegisterCpt2dcrFieldSwcptShift = 0
	RegisterCpt2dcrFieldSwcptMask  = 0x1
)

// GetSwcpt Software Capture
func (r *registerCpt2dcrType) GetSwcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2dcrFieldSwcptMask) != 0
}

// SetSwcpt Software Capture
func (r *registerCpt2dcrType) SetSwcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2dcrFieldSwcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2dcrFieldSwcptMask)
	}
}

// registerOutdrType Timerx Output Register
type registerOutdrType uint32

const (
	RegisterOutdrFieldDidl2Shift = 23
	RegisterOutdrFieldDidl2Mask  = 0x800000
)

// GetDidl2 Output 2 Deadtime upon burst mode Idle entry
func (r *registerOutdrType) GetDidl2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldDidl2Mask) != 0
}

// SetDidl2 Output 2 Deadtime upon burst mode Idle entry
func (r *registerOutdrType) SetDidl2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutdrFieldDidl2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldDidl2Mask)
	}
}

const (
	RegisterOutdrFieldChp2Shift = 22
	RegisterOutdrFieldChp2Mask  = 0x400000
)

// GetChp2 Output 2 Chopper enable
func (r *registerOutdrType) GetChp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldChp2Mask) != 0
}

// SetChp2 Output 2 Chopper enable
func (r *registerOutdrType) SetChp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutdrFieldChp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldChp2Mask)
	}
}

const (
	RegisterOutdrFieldFault2Shift = 20
	RegisterOutdrFieldFault2Mask  = 0x300000
)

// GetFault2 Output 2 Fault state
func (r *registerOutdrType) GetFault2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldFault2Mask) >> RegisterOutdrFieldFault2Shift)
}

// SetFault2 Output 2 Fault state
func (r *registerOutdrType) SetFault2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldFault2Mask)|(uint32(value)<<RegisterOutdrFieldFault2Shift))
}

const (
	RegisterOutdrFieldIdles2Shift = 19
	RegisterOutdrFieldIdles2Mask  = 0x80000
)

// GetIdles2 Output 2 Idle State
func (r *registerOutdrType) GetIdles2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldIdles2Mask) != 0
}

// SetIdles2 Output 2 Idle State
func (r *registerOutdrType) SetIdles2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutdrFieldIdles2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldIdles2Mask)
	}
}

const (
	RegisterOutdrFieldIdlem2Shift = 18
	RegisterOutdrFieldIdlem2Mask  = 0x40000
)

// GetIdlem2 Output 2 Idle mode
func (r *registerOutdrType) GetIdlem2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldIdlem2Mask) != 0
}

// SetIdlem2 Output 2 Idle mode
func (r *registerOutdrType) SetIdlem2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutdrFieldIdlem2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldIdlem2Mask)
	}
}

const (
	RegisterOutdrFieldPol2Shift = 17
	RegisterOutdrFieldPol2Mask  = 0x20000
)

// GetPol2 Output 2 polarity
func (r *registerOutdrType) GetPol2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldPol2Mask) != 0
}

// SetPol2 Output 2 polarity
func (r *registerOutdrType) SetPol2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutdrFieldPol2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldPol2Mask)
	}
}

const (
	RegisterOutdrFieldDlyprtShift = 10
	RegisterOutdrFieldDlyprtMask  = 0x1c00
)

// GetDlyprt Delayed Protection
func (r *registerOutdrType) GetDlyprt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldDlyprtMask) >> RegisterOutdrFieldDlyprtShift)
}

// SetDlyprt Delayed Protection
func (r *registerOutdrType) SetDlyprt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldDlyprtMask)|(uint32(value)<<RegisterOutdrFieldDlyprtShift))
}

const (
	RegisterOutdrFieldDlyprtenShift = 9
	RegisterOutdrFieldDlyprtenMask  = 0x200
)

// GetDlyprten Delayed Protection Enable
func (r *registerOutdrType) GetDlyprten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldDlyprtenMask) != 0
}

// SetDlyprten Delayed Protection Enable
func (r *registerOutdrType) SetDlyprten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutdrFieldDlyprtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldDlyprtenMask)
	}
}

const (
	RegisterOutdrFieldDtenShift = 8
	RegisterOutdrFieldDtenMask  = 0x100
)

// GetDten Deadtime enable
func (r *registerOutdrType) GetDten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldDtenMask) != 0
}

// SetDten Deadtime enable
func (r *registerOutdrType) SetDten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutdrFieldDtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldDtenMask)
	}
}

const (
	RegisterOutdrFieldDidl1Shift = 7
	RegisterOutdrFieldDidl1Mask  = 0x80
)

// GetDidl1 Output 1 Deadtime upon burst mode Idle entry
func (r *registerOutdrType) GetDidl1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldDidl1Mask) != 0
}

// SetDidl1 Output 1 Deadtime upon burst mode Idle entry
func (r *registerOutdrType) SetDidl1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutdrFieldDidl1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldDidl1Mask)
	}
}

const (
	RegisterOutdrFieldChp1Shift = 6
	RegisterOutdrFieldChp1Mask  = 0x40
)

// GetChp1 Output 1 Chopper enable
func (r *registerOutdrType) GetChp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldChp1Mask) != 0
}

// SetChp1 Output 1 Chopper enable
func (r *registerOutdrType) SetChp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutdrFieldChp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldChp1Mask)
	}
}

const (
	RegisterOutdrFieldFault1Shift = 4
	RegisterOutdrFieldFault1Mask  = 0x30
)

// GetFault1 Output 1 Fault state
func (r *registerOutdrType) GetFault1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldFault1Mask) >> RegisterOutdrFieldFault1Shift)
}

// SetFault1 Output 1 Fault state
func (r *registerOutdrType) SetFault1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldFault1Mask)|(uint32(value)<<RegisterOutdrFieldFault1Shift))
}

const (
	RegisterOutdrFieldIdles1Shift = 3
	RegisterOutdrFieldIdles1Mask  = 0x8
)

// GetIdles1 Output 1 Idle State
func (r *registerOutdrType) GetIdles1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldIdles1Mask) != 0
}

// SetIdles1 Output 1 Idle State
func (r *registerOutdrType) SetIdles1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutdrFieldIdles1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldIdles1Mask)
	}
}

const (
	RegisterOutdrFieldIdlem1Shift = 2
	RegisterOutdrFieldIdlem1Mask  = 0x4
)

// GetIdlem1 Output 1 Idle mode
func (r *registerOutdrType) GetIdlem1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldIdlem1Mask) != 0
}

// SetIdlem1 Output 1 Idle mode
func (r *registerOutdrType) SetIdlem1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutdrFieldIdlem1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldIdlem1Mask)
	}
}

const (
	RegisterOutdrFieldPol1Shift = 1
	RegisterOutdrFieldPol1Mask  = 0x2
)

// GetPol1 Output 1 polarity
func (r *registerOutdrType) GetPol1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutdrFieldPol1Mask) != 0
}

// SetPol1 Output 1 polarity
func (r *registerOutdrType) SetPol1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutdrFieldPol1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutdrFieldPol1Mask)
	}
}

// registerFltdrType Timerx Fault Register
type registerFltdrType uint32

const (
	RegisterFltdrFieldFltlckShift = 31
	RegisterFltdrFieldFltlckMask  = 0x80000000
)

// GetFltlck Fault sources Lock
func (r *registerFltdrType) GetFltlck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltdrFieldFltlckMask) != 0
}

// SetFltlck Fault sources Lock
func (r *registerFltdrType) SetFltlck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltdrFieldFltlckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltdrFieldFltlckMask)
	}
}

const (
	RegisterFltdrFieldFlt5enShift = 4
	RegisterFltdrFieldFlt5enMask  = 0x10
)

// GetFlt5en Fault 5 enable
func (r *registerFltdrType) GetFlt5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltdrFieldFlt5enMask) != 0
}

// SetFlt5en Fault 5 enable
func (r *registerFltdrType) SetFlt5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltdrFieldFlt5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltdrFieldFlt5enMask)
	}
}

const (
	RegisterFltdrFieldFlt4enShift = 3
	RegisterFltdrFieldFlt4enMask  = 0x8
)

// GetFlt4en Fault 4 enable
func (r *registerFltdrType) GetFlt4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltdrFieldFlt4enMask) != 0
}

// SetFlt4en Fault 4 enable
func (r *registerFltdrType) SetFlt4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltdrFieldFlt4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltdrFieldFlt4enMask)
	}
}

const (
	RegisterFltdrFieldFlt3enShift = 2
	RegisterFltdrFieldFlt3enMask  = 0x4
)

// GetFlt3en Fault 3 enable
func (r *registerFltdrType) GetFlt3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltdrFieldFlt3enMask) != 0
}

// SetFlt3en Fault 3 enable
func (r *registerFltdrType) SetFlt3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltdrFieldFlt3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltdrFieldFlt3enMask)
	}
}

const (
	RegisterFltdrFieldFlt2enShift = 1
	RegisterFltdrFieldFlt2enMask  = 0x2
)

// GetFlt2en Fault 2 enable
func (r *registerFltdrType) GetFlt2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltdrFieldFlt2enMask) != 0
}

// SetFlt2en Fault 2 enable
func (r *registerFltdrType) SetFlt2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltdrFieldFlt2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltdrFieldFlt2enMask)
	}
}

const (
	RegisterFltdrFieldFlt1enShift = 0
	RegisterFltdrFieldFlt1enMask  = 0x1
)

// GetFlt1en Fault 1 enable
func (r *registerFltdrType) GetFlt1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltdrFieldFlt1enMask) != 0
}

// SetFlt1en Fault 1 enable
func (r *registerFltdrType) SetFlt1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltdrFieldFlt1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltdrFieldFlt1enMask)
	}
}
