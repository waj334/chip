package hrtim_timb

import (
	"unsafe"
	"volatile"
)

var (
	Hrtim_timb = (*_hrtim_timb)(unsafe.Pointer(uintptr(0x40017500)))
)

type _hrtim_timb struct {
	Timbcr    registerTimbcrType
	Timbisr   registerTimbisrType
	Timbicr   registerTimbicrType
	Timbdier5 registerTimbdier5Type
	Cntr      registerCntrType
	Perbr     registerPerbrType
	Repbr     registerRepbrType
	Cmp1br    registerCmp1brType
	Cmp1cbr   registerCmp1cbrType
	Cmp2br    registerCmp2brType
	Cmp3br    registerCmp3brType
	Cmp4br    registerCmp4brType
	Cpt1br    registerCpt1brType
	Cpt2br    registerCpt2brType
	Dtbr      registerDtbrType
	Setb1r    registerSetb1rType
	Rstb1r    registerRstb1rType
	Setb2r    registerSetb2rType
	Rstb2r    registerRstb2rType
	Eefbr1    registerEefbr1Type
	Eefbr2    registerEefbr2Type
	Rstbr     registerRstbrType
	Chpbr     registerChpbrType
	Cpt1bcr   registerCpt1bcrType
	Cpt2bcr   registerCpt2bcrType
	Outbr     registerOutbrType
	Fltbr     registerFltbrType
}

// registerTimbcrType Timerx Control Register
type registerTimbcrType uint32

const (
	RegisterTimbcrFieldUpdgatShift = 28
	RegisterTimbcrFieldUpdgatMask  = 0xf0000000
)

// GetUpdgat Update Gating
func (r *registerTimbcrType) GetUpdgat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldUpdgatMask) >> RegisterTimbcrFieldUpdgatShift)
}

// SetUpdgat Update Gating
func (r *registerTimbcrType) SetUpdgat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldUpdgatMask)|(uint32(value)<<RegisterTimbcrFieldUpdgatShift))
}

const (
	RegisterTimbcrFieldPreenShift = 27
	RegisterTimbcrFieldPreenMask  = 0x8000000
)

// GetPreen Preload enable
func (r *registerTimbcrType) GetPreen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldPreenMask) != 0
}

// SetPreen Preload enable
func (r *registerTimbcrType) SetPreen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldPreenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldPreenMask)
	}
}

const (
	RegisterTimbcrFieldDacsyncShift = 25
	RegisterTimbcrFieldDacsyncMask  = 0x6000000
)

// GetDacsync AC Synchronization
func (r *registerTimbcrType) GetDacsync() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldDacsyncMask) >> RegisterTimbcrFieldDacsyncShift)
}

// SetDacsync AC Synchronization
func (r *registerTimbcrType) SetDacsync(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldDacsyncMask)|(uint32(value)<<RegisterTimbcrFieldDacsyncShift))
}

const (
	RegisterTimbcrFieldMstuShift = 24
	RegisterTimbcrFieldMstuMask  = 0x1000000
)

// GetMstu Master Timer update
func (r *registerTimbcrType) GetMstu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldMstuMask) != 0
}

// SetMstu Master Timer update
func (r *registerTimbcrType) SetMstu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldMstuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldMstuMask)
	}
}

const (
	RegisterTimbcrFieldTeuShift = 23
	RegisterTimbcrFieldTeuMask  = 0x800000
)

// GetTeu TEU
func (r *registerTimbcrType) GetTeu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldTeuMask) != 0
}

// SetTeu TEU
func (r *registerTimbcrType) SetTeu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldTeuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldTeuMask)
	}
}

const (
	RegisterTimbcrFieldTduShift = 22
	RegisterTimbcrFieldTduMask  = 0x400000
)

// GetTdu TDU
func (r *registerTimbcrType) GetTdu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldTduMask) != 0
}

// SetTdu TDU
func (r *registerTimbcrType) SetTdu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldTduMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldTduMask)
	}
}

const (
	RegisterTimbcrFieldTcuShift = 21
	RegisterTimbcrFieldTcuMask  = 0x200000
)

// GetTcu TCU
func (r *registerTimbcrType) GetTcu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldTcuMask) != 0
}

// SetTcu TCU
func (r *registerTimbcrType) SetTcu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldTcuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldTcuMask)
	}
}

const (
	RegisterTimbcrFieldTbuShift = 20
	RegisterTimbcrFieldTbuMask  = 0x100000
)

// GetTbu TBU
func (r *registerTimbcrType) GetTbu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldTbuMask) != 0
}

// SetTbu TBU
func (r *registerTimbcrType) SetTbu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldTbuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldTbuMask)
	}
}

const (
	RegisterTimbcrFieldTxrstuShift = 18
	RegisterTimbcrFieldTxrstuMask  = 0x40000
)

// GetTxrstu Timerx reset update
func (r *registerTimbcrType) GetTxrstu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldTxrstuMask) != 0
}

// SetTxrstu Timerx reset update
func (r *registerTimbcrType) SetTxrstu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldTxrstuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldTxrstuMask)
	}
}

const (
	RegisterTimbcrFieldTxrepuShift = 17
	RegisterTimbcrFieldTxrepuMask  = 0x20000
)

// GetTxrepu Timer x Repetition update
func (r *registerTimbcrType) GetTxrepu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldTxrepuMask) != 0
}

// SetTxrepu Timer x Repetition update
func (r *registerTimbcrType) SetTxrepu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldTxrepuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldTxrepuMask)
	}
}

const (
	RegisterTimbcrFieldDelcmp4Shift = 14
	RegisterTimbcrFieldDelcmp4Mask  = 0xc000
)

// GetDelcmp4 Delayed CMP4 mode
func (r *registerTimbcrType) GetDelcmp4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldDelcmp4Mask) >> RegisterTimbcrFieldDelcmp4Shift)
}

// SetDelcmp4 Delayed CMP4 mode
func (r *registerTimbcrType) SetDelcmp4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldDelcmp4Mask)|(uint32(value)<<RegisterTimbcrFieldDelcmp4Shift))
}

const (
	RegisterTimbcrFieldDelcmp2Shift = 12
	RegisterTimbcrFieldDelcmp2Mask  = 0x3000
)

// GetDelcmp2 Delayed CMP2 mode
func (r *registerTimbcrType) GetDelcmp2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldDelcmp2Mask) >> RegisterTimbcrFieldDelcmp2Shift)
}

// SetDelcmp2 Delayed CMP2 mode
func (r *registerTimbcrType) SetDelcmp2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldDelcmp2Mask)|(uint32(value)<<RegisterTimbcrFieldDelcmp2Shift))
}

const (
	RegisterTimbcrFieldSyncstrtxShift = 11
	RegisterTimbcrFieldSyncstrtxMask  = 0x800
)

// GetSyncstrtx Synchronization Starts Timer x
func (r *registerTimbcrType) GetSyncstrtx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldSyncstrtxMask) != 0
}

// SetSyncstrtx Synchronization Starts Timer x
func (r *registerTimbcrType) SetSyncstrtx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldSyncstrtxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldSyncstrtxMask)
	}
}

const (
	RegisterTimbcrFieldSyncrstxShift = 10
	RegisterTimbcrFieldSyncrstxMask  = 0x400
)

// GetSyncrstx Synchronization Resets Timer x
func (r *registerTimbcrType) GetSyncrstx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldSyncrstxMask) != 0
}

// SetSyncrstx Synchronization Resets Timer x
func (r *registerTimbcrType) SetSyncrstx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldSyncrstxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldSyncrstxMask)
	}
}

const (
	RegisterTimbcrFieldPshpllShift = 6
	RegisterTimbcrFieldPshpllMask  = 0x40
)

// GetPshpll Push-Pull mode enable
func (r *registerTimbcrType) GetPshpll() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldPshpllMask) != 0
}

// SetPshpll Push-Pull mode enable
func (r *registerTimbcrType) SetPshpll(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldPshpllMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldPshpllMask)
	}
}

const (
	RegisterTimbcrFieldHalfShift = 5
	RegisterTimbcrFieldHalfMask  = 0x20
)

// GetHalf Half mode enable
func (r *registerTimbcrType) GetHalf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldHalfMask) != 0
}

// SetHalf Half mode enable
func (r *registerTimbcrType) SetHalf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldHalfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldHalfMask)
	}
}

const (
	RegisterTimbcrFieldRetrigShift = 4
	RegisterTimbcrFieldRetrigMask  = 0x10
)

// GetRetrig Re-triggerable mode
func (r *registerTimbcrType) GetRetrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldRetrigMask) != 0
}

// SetRetrig Re-triggerable mode
func (r *registerTimbcrType) SetRetrig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldRetrigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldRetrigMask)
	}
}

const (
	RegisterTimbcrFieldContShift = 3
	RegisterTimbcrFieldContMask  = 0x8
)

// GetCont Continuous mode
func (r *registerTimbcrType) GetCont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldContMask) != 0
}

// SetCont Continuous mode
func (r *registerTimbcrType) SetCont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbcrFieldContMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldContMask)
	}
}

const (
	RegisterTimbcrFieldCk_pscxShift = 0
	RegisterTimbcrFieldCk_pscxMask  = 0x7
)

// GetCk_pscx HRTIM Timer x Clock prescaler
func (r *registerTimbcrType) GetCk_pscx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimbcrFieldCk_pscxMask) >> RegisterTimbcrFieldCk_pscxShift)
}

// SetCk_pscx HRTIM Timer x Clock prescaler
func (r *registerTimbcrType) SetCk_pscx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimbcrFieldCk_pscxMask)|(uint32(value)<<RegisterTimbcrFieldCk_pscxShift))
}

// registerTimbisrType Timerx Interrupt Status Register
type registerTimbisrType uint32

const (
	RegisterTimbisrFieldO2statShift = 19
	RegisterTimbisrFieldO2statMask  = 0x80000
)

// GetO2stat Output 2 State
func (r *registerTimbisrType) GetO2stat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldO2statMask) != 0
}

// SetO2stat Output 2 State
func (r *registerTimbisrType) SetO2stat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldO2statMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldO2statMask)
	}
}

const (
	RegisterTimbisrFieldO1statShift = 18
	RegisterTimbisrFieldO1statMask  = 0x40000
)

// GetO1stat Output 1 State
func (r *registerTimbisrType) GetO1stat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldO1statMask) != 0
}

// SetO1stat Output 1 State
func (r *registerTimbisrType) SetO1stat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldO1statMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldO1statMask)
	}
}

const (
	RegisterTimbisrFieldIppstatShift = 17
	RegisterTimbisrFieldIppstatMask  = 0x20000
)

// GetIppstat Idle Push Pull Status
func (r *registerTimbisrType) GetIppstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldIppstatMask) != 0
}

// SetIppstat Idle Push Pull Status
func (r *registerTimbisrType) SetIppstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldIppstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldIppstatMask)
	}
}

const (
	RegisterTimbisrFieldCppstatShift = 16
	RegisterTimbisrFieldCppstatMask  = 0x10000
)

// GetCppstat Current Push Pull Status
func (r *registerTimbisrType) GetCppstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldCppstatMask) != 0
}

// SetCppstat Current Push Pull Status
func (r *registerTimbisrType) SetCppstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldCppstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldCppstatMask)
	}
}

const (
	RegisterTimbisrFieldDlyprtShift = 14
	RegisterTimbisrFieldDlyprtMask  = 0x4000
)

// GetDlyprt Delayed Protection Flag
func (r *registerTimbisrType) GetDlyprt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldDlyprtMask) != 0
}

// SetDlyprt Delayed Protection Flag
func (r *registerTimbisrType) SetDlyprt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldDlyprtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldDlyprtMask)
	}
}

const (
	RegisterTimbisrFieldRstShift = 13
	RegisterTimbisrFieldRstMask  = 0x2000
)

// GetRst Reset Interrupt Flag
func (r *registerTimbisrType) GetRst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldRstMask) != 0
}

// SetRst Reset Interrupt Flag
func (r *registerTimbisrType) SetRst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldRstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldRstMask)
	}
}

const (
	RegisterTimbisrFieldRstx2Shift = 12
	RegisterTimbisrFieldRstx2Mask  = 0x1000
)

// GetRstx2 Output 2 Reset Interrupt Flag
func (r *registerTimbisrType) GetRstx2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldRstx2Mask) != 0
}

// SetRstx2 Output 2 Reset Interrupt Flag
func (r *registerTimbisrType) SetRstx2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldRstx2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldRstx2Mask)
	}
}

const (
	RegisterTimbisrFieldSetx2Shift = 11
	RegisterTimbisrFieldSetx2Mask  = 0x800
)

// GetSetx2 Output 2 Set Interrupt Flag
func (r *registerTimbisrType) GetSetx2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldSetx2Mask) != 0
}

// SetSetx2 Output 2 Set Interrupt Flag
func (r *registerTimbisrType) SetSetx2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldSetx2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldSetx2Mask)
	}
}

const (
	RegisterTimbisrFieldRstx1Shift = 10
	RegisterTimbisrFieldRstx1Mask  = 0x400
)

// GetRstx1 Output 1 Reset Interrupt Flag
func (r *registerTimbisrType) GetRstx1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldRstx1Mask) != 0
}

// SetRstx1 Output 1 Reset Interrupt Flag
func (r *registerTimbisrType) SetRstx1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldRstx1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldRstx1Mask)
	}
}

const (
	RegisterTimbisrFieldSetx1Shift = 9
	RegisterTimbisrFieldSetx1Mask  = 0x200
)

// GetSetx1 Output 1 Set Interrupt Flag
func (r *registerTimbisrType) GetSetx1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldSetx1Mask) != 0
}

// SetSetx1 Output 1 Set Interrupt Flag
func (r *registerTimbisrType) SetSetx1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldSetx1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldSetx1Mask)
	}
}

const (
	RegisterTimbisrFieldCpt2Shift = 8
	RegisterTimbisrFieldCpt2Mask  = 0x100
)

// GetCpt2 Capture2 Interrupt Flag
func (r *registerTimbisrType) GetCpt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldCpt2Mask) != 0
}

// SetCpt2 Capture2 Interrupt Flag
func (r *registerTimbisrType) SetCpt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldCpt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldCpt2Mask)
	}
}

const (
	RegisterTimbisrFieldCpt1Shift = 7
	RegisterTimbisrFieldCpt1Mask  = 0x80
)

// GetCpt1 Capture1 Interrupt Flag
func (r *registerTimbisrType) GetCpt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldCpt1Mask) != 0
}

// SetCpt1 Capture1 Interrupt Flag
func (r *registerTimbisrType) SetCpt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldCpt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldCpt1Mask)
	}
}

const (
	RegisterTimbisrFieldUpdShift = 6
	RegisterTimbisrFieldUpdMask  = 0x40
)

// GetUpd Update Interrupt Flag
func (r *registerTimbisrType) GetUpd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldUpdMask) != 0
}

// SetUpd Update Interrupt Flag
func (r *registerTimbisrType) SetUpd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldUpdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldUpdMask)
	}
}

const (
	RegisterTimbisrFieldRepShift = 4
	RegisterTimbisrFieldRepMask  = 0x10
)

// GetRep Repetition Interrupt Flag
func (r *registerTimbisrType) GetRep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldRepMask) != 0
}

// SetRep Repetition Interrupt Flag
func (r *registerTimbisrType) SetRep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldRepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldRepMask)
	}
}

const (
	RegisterTimbisrFieldCmp4Shift = 3
	RegisterTimbisrFieldCmp4Mask  = 0x8
)

// GetCmp4 Compare 4 Interrupt Flag
func (r *registerTimbisrType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldCmp4Mask) != 0
}

// SetCmp4 Compare 4 Interrupt Flag
func (r *registerTimbisrType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldCmp4Mask)
	}
}

const (
	RegisterTimbisrFieldCmp3Shift = 2
	RegisterTimbisrFieldCmp3Mask  = 0x4
)

// GetCmp3 Compare 3 Interrupt Flag
func (r *registerTimbisrType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldCmp3Mask) != 0
}

// SetCmp3 Compare 3 Interrupt Flag
func (r *registerTimbisrType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldCmp3Mask)
	}
}

const (
	RegisterTimbisrFieldCmp2Shift = 1
	RegisterTimbisrFieldCmp2Mask  = 0x2
)

// GetCmp2 Compare 2 Interrupt Flag
func (r *registerTimbisrType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldCmp2Mask) != 0
}

// SetCmp2 Compare 2 Interrupt Flag
func (r *registerTimbisrType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldCmp2Mask)
	}
}

const (
	RegisterTimbisrFieldCmp1Shift = 0
	RegisterTimbisrFieldCmp1Mask  = 0x1
)

// GetCmp1 Compare 1 Interrupt Flag
func (r *registerTimbisrType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbisrFieldCmp1Mask) != 0
}

// SetCmp1 Compare 1 Interrupt Flag
func (r *registerTimbisrType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbisrFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbisrFieldCmp1Mask)
	}
}

// registerTimbicrType Timerx Interrupt Clear Register
type registerTimbicrType uint32

const (
	RegisterTimbicrFieldDlyprtcShift = 14
	RegisterTimbicrFieldDlyprtcMask  = 0x4000
)

// GetDlyprtc Delayed Protection Flag Clear
func (r *registerTimbicrType) GetDlyprtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldDlyprtcMask) != 0
}

// SetDlyprtc Delayed Protection Flag Clear
func (r *registerTimbicrType) SetDlyprtc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldDlyprtcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldDlyprtcMask)
	}
}

const (
	RegisterTimbicrFieldRstcShift = 13
	RegisterTimbicrFieldRstcMask  = 0x2000
)

// GetRstc Reset Interrupt flag Clear
func (r *registerTimbicrType) GetRstc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldRstcMask) != 0
}

// SetRstc Reset Interrupt flag Clear
func (r *registerTimbicrType) SetRstc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldRstcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldRstcMask)
	}
}

const (
	RegisterTimbicrFieldRstx2cShift = 12
	RegisterTimbicrFieldRstx2cMask  = 0x1000
)

// GetRstx2c Output 2 Reset flag Clear
func (r *registerTimbicrType) GetRstx2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldRstx2cMask) != 0
}

// SetRstx2c Output 2 Reset flag Clear
func (r *registerTimbicrType) SetRstx2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldRstx2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldRstx2cMask)
	}
}

const (
	RegisterTimbicrFieldSet2xcShift = 11
	RegisterTimbicrFieldSet2xcMask  = 0x800
)

// GetSet2xc Output 2 Set flag Clear
func (r *registerTimbicrType) GetSet2xc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldSet2xcMask) != 0
}

// SetSet2xc Output 2 Set flag Clear
func (r *registerTimbicrType) SetSet2xc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldSet2xcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldSet2xcMask)
	}
}

const (
	RegisterTimbicrFieldRstx1cShift = 10
	RegisterTimbicrFieldRstx1cMask  = 0x400
)

// GetRstx1c Output 1 Reset flag Clear
func (r *registerTimbicrType) GetRstx1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldRstx1cMask) != 0
}

// SetRstx1c Output 1 Reset flag Clear
func (r *registerTimbicrType) SetRstx1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldRstx1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldRstx1cMask)
	}
}

const (
	RegisterTimbicrFieldSet1xcShift = 9
	RegisterTimbicrFieldSet1xcMask  = 0x200
)

// GetSet1xc Output 1 Set flag Clear
func (r *registerTimbicrType) GetSet1xc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldSet1xcMask) != 0
}

// SetSet1xc Output 1 Set flag Clear
func (r *registerTimbicrType) SetSet1xc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldSet1xcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldSet1xcMask)
	}
}

const (
	RegisterTimbicrFieldCpt2cShift = 8
	RegisterTimbicrFieldCpt2cMask  = 0x100
)

// GetCpt2c Capture2 Interrupt flag Clear
func (r *registerTimbicrType) GetCpt2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldCpt2cMask) != 0
}

// SetCpt2c Capture2 Interrupt flag Clear
func (r *registerTimbicrType) SetCpt2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldCpt2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldCpt2cMask)
	}
}

const (
	RegisterTimbicrFieldCpt1cShift = 7
	RegisterTimbicrFieldCpt1cMask  = 0x80
)

// GetCpt1c Capture1 Interrupt flag Clear
func (r *registerTimbicrType) GetCpt1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldCpt1cMask) != 0
}

// SetCpt1c Capture1 Interrupt flag Clear
func (r *registerTimbicrType) SetCpt1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldCpt1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldCpt1cMask)
	}
}

const (
	RegisterTimbicrFieldUpdcShift = 6
	RegisterTimbicrFieldUpdcMask  = 0x40
)

// GetUpdc Update Interrupt flag Clear
func (r *registerTimbicrType) GetUpdc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldUpdcMask) != 0
}

// SetUpdc Update Interrupt flag Clear
func (r *registerTimbicrType) SetUpdc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldUpdcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldUpdcMask)
	}
}

const (
	RegisterTimbicrFieldRepcShift = 4
	RegisterTimbicrFieldRepcMask  = 0x10
)

// GetRepc Repetition Interrupt flag Clear
func (r *registerTimbicrType) GetRepc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldRepcMask) != 0
}

// SetRepc Repetition Interrupt flag Clear
func (r *registerTimbicrType) SetRepc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldRepcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldRepcMask)
	}
}

const (
	RegisterTimbicrFieldCmp4cShift = 3
	RegisterTimbicrFieldCmp4cMask  = 0x8
)

// GetCmp4c Compare 4 Interrupt flag Clear
func (r *registerTimbicrType) GetCmp4c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldCmp4cMask) != 0
}

// SetCmp4c Compare 4 Interrupt flag Clear
func (r *registerTimbicrType) SetCmp4c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldCmp4cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldCmp4cMask)
	}
}

const (
	RegisterTimbicrFieldCmp3cShift = 2
	RegisterTimbicrFieldCmp3cMask  = 0x4
)

// GetCmp3c Compare 3 Interrupt flag Clear
func (r *registerTimbicrType) GetCmp3c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldCmp3cMask) != 0
}

// SetCmp3c Compare 3 Interrupt flag Clear
func (r *registerTimbicrType) SetCmp3c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldCmp3cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldCmp3cMask)
	}
}

const (
	RegisterTimbicrFieldCmp2cShift = 1
	RegisterTimbicrFieldCmp2cMask  = 0x2
)

// GetCmp2c Compare 2 Interrupt flag Clear
func (r *registerTimbicrType) GetCmp2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldCmp2cMask) != 0
}

// SetCmp2c Compare 2 Interrupt flag Clear
func (r *registerTimbicrType) SetCmp2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldCmp2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldCmp2cMask)
	}
}

const (
	RegisterTimbicrFieldCmp1cShift = 0
	RegisterTimbicrFieldCmp1cMask  = 0x1
)

// GetCmp1c Compare 1 Interrupt flag Clear
func (r *registerTimbicrType) GetCmp1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbicrFieldCmp1cMask) != 0
}

// SetCmp1c Compare 1 Interrupt flag Clear
func (r *registerTimbicrType) SetCmp1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbicrFieldCmp1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbicrFieldCmp1cMask)
	}
}

// registerTimbdier5Type TIMxDIER5
type registerTimbdier5Type uint32

const (
	RegisterTimbdier5FieldDlyprtdeShift = 30
	RegisterTimbdier5FieldDlyprtdeMask  = 0x40000000
)

// GetDlyprtde DLYPRTDE
func (r *registerTimbdier5Type) GetDlyprtde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldDlyprtdeMask) != 0
}

// SetDlyprtde DLYPRTDE
func (r *registerTimbdier5Type) SetDlyprtde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldDlyprtdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldDlyprtdeMask)
	}
}

const (
	RegisterTimbdier5FieldRstdeShift = 29
	RegisterTimbdier5FieldRstdeMask  = 0x20000000
)

// GetRstde RSTDE
func (r *registerTimbdier5Type) GetRstde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldRstdeMask) != 0
}

// SetRstde RSTDE
func (r *registerTimbdier5Type) SetRstde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldRstdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldRstdeMask)
	}
}

const (
	RegisterTimbdier5FieldRstx2deShift = 28
	RegisterTimbdier5FieldRstx2deMask  = 0x10000000
)

// GetRstx2de RSTx2DE
func (r *registerTimbdier5Type) GetRstx2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldRstx2deMask) != 0
}

// SetRstx2de RSTx2DE
func (r *registerTimbdier5Type) SetRstx2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldRstx2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldRstx2deMask)
	}
}

const (
	RegisterTimbdier5FieldSetx2deShift = 27
	RegisterTimbdier5FieldSetx2deMask  = 0x8000000
)

// GetSetx2de SETx2DE
func (r *registerTimbdier5Type) GetSetx2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldSetx2deMask) != 0
}

// SetSetx2de SETx2DE
func (r *registerTimbdier5Type) SetSetx2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldSetx2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldSetx2deMask)
	}
}

const (
	RegisterTimbdier5FieldRstx1deShift = 26
	RegisterTimbdier5FieldRstx1deMask  = 0x4000000
)

// GetRstx1de RSTx1DE
func (r *registerTimbdier5Type) GetRstx1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldRstx1deMask) != 0
}

// SetRstx1de RSTx1DE
func (r *registerTimbdier5Type) SetRstx1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldRstx1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldRstx1deMask)
	}
}

const (
	RegisterTimbdier5FieldSet1xdeShift = 25
	RegisterTimbdier5FieldSet1xdeMask  = 0x2000000
)

// GetSet1xde SET1xDE
func (r *registerTimbdier5Type) GetSet1xde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldSet1xdeMask) != 0
}

// SetSet1xde SET1xDE
func (r *registerTimbdier5Type) SetSet1xde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldSet1xdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldSet1xdeMask)
	}
}

const (
	RegisterTimbdier5FieldCpt2deShift = 24
	RegisterTimbdier5FieldCpt2deMask  = 0x1000000
)

// GetCpt2de CPT2DE
func (r *registerTimbdier5Type) GetCpt2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldCpt2deMask) != 0
}

// SetCpt2de CPT2DE
func (r *registerTimbdier5Type) SetCpt2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldCpt2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldCpt2deMask)
	}
}

const (
	RegisterTimbdier5FieldCpt1deShift = 23
	RegisterTimbdier5FieldCpt1deMask  = 0x800000
)

// GetCpt1de CPT1DE
func (r *registerTimbdier5Type) GetCpt1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldCpt1deMask) != 0
}

// SetCpt1de CPT1DE
func (r *registerTimbdier5Type) SetCpt1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldCpt1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldCpt1deMask)
	}
}

const (
	RegisterTimbdier5FieldUpddeShift = 22
	RegisterTimbdier5FieldUpddeMask  = 0x400000
)

// GetUpdde UPDDE
func (r *registerTimbdier5Type) GetUpdde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldUpddeMask) != 0
}

// SetUpdde UPDDE
func (r *registerTimbdier5Type) SetUpdde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldUpddeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldUpddeMask)
	}
}

const (
	RegisterTimbdier5FieldRepdeShift = 20
	RegisterTimbdier5FieldRepdeMask  = 0x100000
)

// GetRepde REPDE
func (r *registerTimbdier5Type) GetRepde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldRepdeMask) != 0
}

// SetRepde REPDE
func (r *registerTimbdier5Type) SetRepde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldRepdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldRepdeMask)
	}
}

const (
	RegisterTimbdier5FieldCmp4deShift = 19
	RegisterTimbdier5FieldCmp4deMask  = 0x80000
)

// GetCmp4de CMP4DE
func (r *registerTimbdier5Type) GetCmp4de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldCmp4deMask) != 0
}

// SetCmp4de CMP4DE
func (r *registerTimbdier5Type) SetCmp4de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldCmp4deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldCmp4deMask)
	}
}

const (
	RegisterTimbdier5FieldCmp3deShift = 18
	RegisterTimbdier5FieldCmp3deMask  = 0x40000
)

// GetCmp3de CMP3DE
func (r *registerTimbdier5Type) GetCmp3de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldCmp3deMask) != 0
}

// SetCmp3de CMP3DE
func (r *registerTimbdier5Type) SetCmp3de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldCmp3deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldCmp3deMask)
	}
}

const (
	RegisterTimbdier5FieldCmp2deShift = 17
	RegisterTimbdier5FieldCmp2deMask  = 0x20000
)

// GetCmp2de CMP2DE
func (r *registerTimbdier5Type) GetCmp2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldCmp2deMask) != 0
}

// SetCmp2de CMP2DE
func (r *registerTimbdier5Type) SetCmp2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldCmp2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldCmp2deMask)
	}
}

const (
	RegisterTimbdier5FieldCmp1deShift = 16
	RegisterTimbdier5FieldCmp1deMask  = 0x10000
)

// GetCmp1de CMP1DE
func (r *registerTimbdier5Type) GetCmp1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldCmp1deMask) != 0
}

// SetCmp1de CMP1DE
func (r *registerTimbdier5Type) SetCmp1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldCmp1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldCmp1deMask)
	}
}

const (
	RegisterTimbdier5FieldDlyprtieShift = 14
	RegisterTimbdier5FieldDlyprtieMask  = 0x4000
)

// GetDlyprtie DLYPRTIE
func (r *registerTimbdier5Type) GetDlyprtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldDlyprtieMask) != 0
}

// SetDlyprtie DLYPRTIE
func (r *registerTimbdier5Type) SetDlyprtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldDlyprtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldDlyprtieMask)
	}
}

const (
	RegisterTimbdier5FieldRstieShift = 13
	RegisterTimbdier5FieldRstieMask  = 0x2000
)

// GetRstie RSTIE
func (r *registerTimbdier5Type) GetRstie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldRstieMask) != 0
}

// SetRstie RSTIE
func (r *registerTimbdier5Type) SetRstie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldRstieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldRstieMask)
	}
}

const (
	RegisterTimbdier5FieldRstx2ieShift = 12
	RegisterTimbdier5FieldRstx2ieMask  = 0x1000
)

// GetRstx2ie RSTx2IE
func (r *registerTimbdier5Type) GetRstx2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldRstx2ieMask) != 0
}

// SetRstx2ie RSTx2IE
func (r *registerTimbdier5Type) SetRstx2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldRstx2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldRstx2ieMask)
	}
}

const (
	RegisterTimbdier5FieldSetx2ieShift = 11
	RegisterTimbdier5FieldSetx2ieMask  = 0x800
)

// GetSetx2ie SETx2IE
func (r *registerTimbdier5Type) GetSetx2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldSetx2ieMask) != 0
}

// SetSetx2ie SETx2IE
func (r *registerTimbdier5Type) SetSetx2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldSetx2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldSetx2ieMask)
	}
}

const (
	RegisterTimbdier5FieldRstx1ieShift = 10
	RegisterTimbdier5FieldRstx1ieMask  = 0x400
)

// GetRstx1ie RSTx1IE
func (r *registerTimbdier5Type) GetRstx1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldRstx1ieMask) != 0
}

// SetRstx1ie RSTx1IE
func (r *registerTimbdier5Type) SetRstx1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldRstx1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldRstx1ieMask)
	}
}

const (
	RegisterTimbdier5FieldSet1xieShift = 9
	RegisterTimbdier5FieldSet1xieMask  = 0x200
)

// GetSet1xie SET1xIE
func (r *registerTimbdier5Type) GetSet1xie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldSet1xieMask) != 0
}

// SetSet1xie SET1xIE
func (r *registerTimbdier5Type) SetSet1xie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldSet1xieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldSet1xieMask)
	}
}

const (
	RegisterTimbdier5FieldCpt2ieShift = 8
	RegisterTimbdier5FieldCpt2ieMask  = 0x100
)

// GetCpt2ie CPT2IE
func (r *registerTimbdier5Type) GetCpt2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldCpt2ieMask) != 0
}

// SetCpt2ie CPT2IE
func (r *registerTimbdier5Type) SetCpt2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldCpt2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldCpt2ieMask)
	}
}

const (
	RegisterTimbdier5FieldCpt1ieShift = 7
	RegisterTimbdier5FieldCpt1ieMask  = 0x80
)

// GetCpt1ie CPT1IE
func (r *registerTimbdier5Type) GetCpt1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldCpt1ieMask) != 0
}

// SetCpt1ie CPT1IE
func (r *registerTimbdier5Type) SetCpt1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldCpt1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldCpt1ieMask)
	}
}

const (
	RegisterTimbdier5FieldUpdieShift = 6
	RegisterTimbdier5FieldUpdieMask  = 0x40
)

// GetUpdie UPDIE
func (r *registerTimbdier5Type) GetUpdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldUpdieMask) != 0
}

// SetUpdie UPDIE
func (r *registerTimbdier5Type) SetUpdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldUpdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldUpdieMask)
	}
}

const (
	RegisterTimbdier5FieldRepieShift = 4
	RegisterTimbdier5FieldRepieMask  = 0x10
)

// GetRepie REPIE
func (r *registerTimbdier5Type) GetRepie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldRepieMask) != 0
}

// SetRepie REPIE
func (r *registerTimbdier5Type) SetRepie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldRepieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldRepieMask)
	}
}

const (
	RegisterTimbdier5FieldCmp4ieShift = 3
	RegisterTimbdier5FieldCmp4ieMask  = 0x8
)

// GetCmp4ie CMP4IE
func (r *registerTimbdier5Type) GetCmp4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldCmp4ieMask) != 0
}

// SetCmp4ie CMP4IE
func (r *registerTimbdier5Type) SetCmp4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldCmp4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldCmp4ieMask)
	}
}

const (
	RegisterTimbdier5FieldCmp3ieShift = 2
	RegisterTimbdier5FieldCmp3ieMask  = 0x4
)

// GetCmp3ie CMP3IE
func (r *registerTimbdier5Type) GetCmp3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldCmp3ieMask) != 0
}

// SetCmp3ie CMP3IE
func (r *registerTimbdier5Type) SetCmp3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldCmp3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldCmp3ieMask)
	}
}

const (
	RegisterTimbdier5FieldCmp2ieShift = 1
	RegisterTimbdier5FieldCmp2ieMask  = 0x2
)

// GetCmp2ie CMP2IE
func (r *registerTimbdier5Type) GetCmp2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldCmp2ieMask) != 0
}

// SetCmp2ie CMP2IE
func (r *registerTimbdier5Type) SetCmp2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldCmp2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldCmp2ieMask)
	}
}

const (
	RegisterTimbdier5FieldCmp1ieShift = 0
	RegisterTimbdier5FieldCmp1ieMask  = 0x1
)

// GetCmp1ie CMP1IE
func (r *registerTimbdier5Type) GetCmp1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimbdier5FieldCmp1ieMask) != 0
}

// SetCmp1ie CMP1IE
func (r *registerTimbdier5Type) SetCmp1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimbdier5FieldCmp1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimbdier5FieldCmp1ieMask)
	}
}

// registerCntrType Timerx Counter Register
type registerCntrType uint32

const (
	RegisterCntrFieldCntxShift = 0
	RegisterCntrFieldCntxMask  = 0xffff
)

// GetCntx Timerx Counter value
func (r *registerCntrType) GetCntx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntrFieldCntxMask) >> RegisterCntrFieldCntxShift)
}

// SetCntx Timerx Counter value
func (r *registerCntrType) SetCntx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntrFieldCntxMask)|(uint32(value)<<RegisterCntrFieldCntxShift))
}

// registerPerbrType Timerx Period Register
type registerPerbrType uint32

const (
	RegisterPerbrFieldPerxShift = 0
	RegisterPerbrFieldPerxMask  = 0xffff
)

// GetPerx Timerx Period value
func (r *registerPerbrType) GetPerx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPerbrFieldPerxMask) >> RegisterPerbrFieldPerxShift)
}

// SetPerx Timerx Period value
func (r *registerPerbrType) SetPerx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPerbrFieldPerxMask)|(uint32(value)<<RegisterPerbrFieldPerxShift))
}

// registerRepbrType Timerx Repetition Register
type registerRepbrType uint32

const (
	RegisterRepbrFieldRepxShift = 0
	RegisterRepbrFieldRepxMask  = 0xff
)

// GetRepx Timerx Repetition counter value
func (r *registerRepbrType) GetRepx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRepbrFieldRepxMask) >> RegisterRepbrFieldRepxShift)
}

// SetRepx Timerx Repetition counter value
func (r *registerRepbrType) SetRepx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRepbrFieldRepxMask)|(uint32(value)<<RegisterRepbrFieldRepxShift))
}

// registerCmp1brType Timerx Compare 1 Register
type registerCmp1brType uint32

const (
	RegisterCmp1brFieldCmp1xShift = 0
	RegisterCmp1brFieldCmp1xMask  = 0xffff
)

// GetCmp1x Timerx Compare 1 value
func (r *registerCmp1brType) GetCmp1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1brFieldCmp1xMask) >> RegisterCmp1brFieldCmp1xShift)
}

// SetCmp1x Timerx Compare 1 value
func (r *registerCmp1brType) SetCmp1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1brFieldCmp1xMask)|(uint32(value)<<RegisterCmp1brFieldCmp1xShift))
}

// registerCmp1cbrType Timerx Compare 1 Compound Register
type registerCmp1cbrType uint32

const (
	RegisterCmp1cbrFieldRepxShift = 16
	RegisterCmp1cbrFieldRepxMask  = 0xff0000
)

// GetRepx Timerx Repetition value (aliased from HRTIM_REPx register)
func (r *registerCmp1cbrType) GetRepx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1cbrFieldRepxMask) >> RegisterCmp1cbrFieldRepxShift)
}

// SetRepx Timerx Repetition value (aliased from HRTIM_REPx register)
func (r *registerCmp1cbrType) SetRepx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1cbrFieldRepxMask)|(uint32(value)<<RegisterCmp1cbrFieldRepxShift))
}

const (
	RegisterCmp1cbrFieldCmp1xShift = 0
	RegisterCmp1cbrFieldCmp1xMask  = 0xffff
)

// GetCmp1x Timerx Compare 1 value
func (r *registerCmp1cbrType) GetCmp1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1cbrFieldCmp1xMask) >> RegisterCmp1cbrFieldCmp1xShift)
}

// SetCmp1x Timerx Compare 1 value
func (r *registerCmp1cbrType) SetCmp1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1cbrFieldCmp1xMask)|(uint32(value)<<RegisterCmp1cbrFieldCmp1xShift))
}

// registerCmp2brType Timerx Compare 2 Register
type registerCmp2brType uint32

const (
	RegisterCmp2brFieldCmp2xShift = 0
	RegisterCmp2brFieldCmp2xMask  = 0xffff
)

// GetCmp2x Timerx Compare 2 value
func (r *registerCmp2brType) GetCmp2x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp2brFieldCmp2xMask) >> RegisterCmp2brFieldCmp2xShift)
}

// SetCmp2x Timerx Compare 2 value
func (r *registerCmp2brType) SetCmp2x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp2brFieldCmp2xMask)|(uint32(value)<<RegisterCmp2brFieldCmp2xShift))
}

// registerCmp3brType Timerx Compare 3 Register
type registerCmp3brType uint32

const (
	RegisterCmp3brFieldCmp3xShift = 0
	RegisterCmp3brFieldCmp3xMask  = 0xffff
)

// GetCmp3x Timerx Compare 3 value
func (r *registerCmp3brType) GetCmp3x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp3brFieldCmp3xMask) >> RegisterCmp3brFieldCmp3xShift)
}

// SetCmp3x Timerx Compare 3 value
func (r *registerCmp3brType) SetCmp3x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp3brFieldCmp3xMask)|(uint32(value)<<RegisterCmp3brFieldCmp3xShift))
}

// registerCmp4brType Timerx Compare 4 Register
type registerCmp4brType uint32

const (
	RegisterCmp4brFieldCmp4xShift = 0
	RegisterCmp4brFieldCmp4xMask  = 0xffff
)

// GetCmp4x Timerx Compare 4 value
func (r *registerCmp4brType) GetCmp4x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp4brFieldCmp4xMask) >> RegisterCmp4brFieldCmp4xShift)
}

// SetCmp4x Timerx Compare 4 value
func (r *registerCmp4brType) SetCmp4x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp4brFieldCmp4xMask)|(uint32(value)<<RegisterCmp4brFieldCmp4xShift))
}

// registerCpt1brType Timerx Capture 1 Register
type registerCpt1brType uint32

const (
	RegisterCpt1brFieldCpt1xShift = 0
	RegisterCpt1brFieldCpt1xMask  = 0xffff
)

// GetCpt1x Timerx Capture 1 value
func (r *registerCpt1brType) GetCpt1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpt1brFieldCpt1xMask) >> RegisterCpt1brFieldCpt1xShift)
}

// SetCpt1x Timerx Capture 1 value
func (r *registerCpt1brType) SetCpt1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpt1brFieldCpt1xMask)|(uint32(value)<<RegisterCpt1brFieldCpt1xShift))
}

// registerCpt2brType Timerx Capture 2 Register
type registerCpt2brType uint32

const (
	RegisterCpt2brFieldCpt2xShift = 0
	RegisterCpt2brFieldCpt2xMask  = 0xffff
)

// GetCpt2x Timerx Capture 2 value
func (r *registerCpt2brType) GetCpt2x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpt2brFieldCpt2xMask) >> RegisterCpt2brFieldCpt2xShift)
}

// SetCpt2x Timerx Capture 2 value
func (r *registerCpt2brType) SetCpt2x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpt2brFieldCpt2xMask)|(uint32(value)<<RegisterCpt2brFieldCpt2xShift))
}

// registerDtbrType Timerx Deadtime Register
type registerDtbrType uint32

const (
	RegisterDtbrFieldDtflkxShift = 31
	RegisterDtbrFieldDtflkxMask  = 0x80000000
)

// GetDtflkx Deadtime Falling Lock
func (r *registerDtbrType) GetDtflkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtbrFieldDtflkxMask) != 0
}

// SetDtflkx Deadtime Falling Lock
func (r *registerDtbrType) SetDtflkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtbrFieldDtflkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtbrFieldDtflkxMask)
	}
}

const (
	RegisterDtbrFieldDtfslkxShift = 30
	RegisterDtbrFieldDtfslkxMask  = 0x40000000
)

// GetDtfslkx Deadtime Falling Sign Lock
func (r *registerDtbrType) GetDtfslkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtbrFieldDtfslkxMask) != 0
}

// SetDtfslkx Deadtime Falling Sign Lock
func (r *registerDtbrType) SetDtfslkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtbrFieldDtfslkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtbrFieldDtfslkxMask)
	}
}

const (
	RegisterDtbrFieldSdtfxShift = 25
	RegisterDtbrFieldSdtfxMask  = 0x2000000
)

// GetSdtfx Sign Deadtime Falling value
func (r *registerDtbrType) GetSdtfx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtbrFieldSdtfxMask) != 0
}

// SetSdtfx Sign Deadtime Falling value
func (r *registerDtbrType) SetSdtfx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtbrFieldSdtfxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtbrFieldSdtfxMask)
	}
}

const (
	RegisterDtbrFieldDtfxShift = 16
	RegisterDtbrFieldDtfxMask  = 0x1ff0000
)

// GetDtfx Deadtime Falling value
func (r *registerDtbrType) GetDtfx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDtbrFieldDtfxMask) >> RegisterDtbrFieldDtfxShift)
}

// SetDtfx Deadtime Falling value
func (r *registerDtbrType) SetDtfx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtbrFieldDtfxMask)|(uint32(value)<<RegisterDtbrFieldDtfxShift))
}

const (
	RegisterDtbrFieldDtrlkxShift = 15
	RegisterDtbrFieldDtrlkxMask  = 0x8000
)

// GetDtrlkx Deadtime Rising Lock
func (r *registerDtbrType) GetDtrlkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtbrFieldDtrlkxMask) != 0
}

// SetDtrlkx Deadtime Rising Lock
func (r *registerDtbrType) SetDtrlkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtbrFieldDtrlkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtbrFieldDtrlkxMask)
	}
}

const (
	RegisterDtbrFieldDtrslkxShift = 14
	RegisterDtbrFieldDtrslkxMask  = 0x4000
)

// GetDtrslkx Deadtime Rising Sign Lock
func (r *registerDtbrType) GetDtrslkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtbrFieldDtrslkxMask) != 0
}

// SetDtrslkx Deadtime Rising Sign Lock
func (r *registerDtbrType) SetDtrslkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtbrFieldDtrslkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtbrFieldDtrslkxMask)
	}
}

const (
	RegisterDtbrFieldDtprscShift = 10
	RegisterDtbrFieldDtprscMask  = 0x1c00
)

// GetDtprsc Deadtime Prescaler
func (r *registerDtbrType) GetDtprsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDtbrFieldDtprscMask) >> RegisterDtbrFieldDtprscShift)
}

// SetDtprsc Deadtime Prescaler
func (r *registerDtbrType) SetDtprsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtbrFieldDtprscMask)|(uint32(value)<<RegisterDtbrFieldDtprscShift))
}

const (
	RegisterDtbrFieldSdtrxShift = 9
	RegisterDtbrFieldSdtrxMask  = 0x200
)

// GetSdtrx Sign Deadtime Rising value
func (r *registerDtbrType) GetSdtrx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtbrFieldSdtrxMask) != 0
}

// SetSdtrx Sign Deadtime Rising value
func (r *registerDtbrType) SetSdtrx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtbrFieldSdtrxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtbrFieldSdtrxMask)
	}
}

const (
	RegisterDtbrFieldDtrxShift = 0
	RegisterDtbrFieldDtrxMask  = 0x1ff
)

// GetDtrx Deadtime Rising value
func (r *registerDtbrType) GetDtrx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDtbrFieldDtrxMask) >> RegisterDtbrFieldDtrxShift)
}

// SetDtrx Deadtime Rising value
func (r *registerDtbrType) SetDtrx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtbrFieldDtrxMask)|(uint32(value)<<RegisterDtbrFieldDtrxShift))
}

// registerSetb1rType Timerx Output1 Set Register
type registerSetb1rType uint32

const (
	RegisterSetb1rFieldUpdateShift = 31
	RegisterSetb1rFieldUpdateMask  = 0x80000000
)

// GetUpdate Registers update (transfer preload to active)
func (r *registerSetb1rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldUpdateMask) != 0
}

// SetUpdate Registers update (transfer preload to active)
func (r *registerSetb1rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldUpdateMask)
	}
}

const (
	RegisterSetb1rFieldExtevnt10Shift = 30
	RegisterSetb1rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 External Event 10
func (r *registerSetb1rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 External Event 10
func (r *registerSetb1rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldExtevnt10Mask)
	}
}

const (
	RegisterSetb1rFieldExtevnt9Shift = 29
	RegisterSetb1rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 External Event 9
func (r *registerSetb1rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 External Event 9
func (r *registerSetb1rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldExtevnt9Mask)
	}
}

const (
	RegisterSetb1rFieldExtevnt8Shift = 28
	RegisterSetb1rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 External Event 8
func (r *registerSetb1rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 External Event 8
func (r *registerSetb1rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldExtevnt8Mask)
	}
}

const (
	RegisterSetb1rFieldExtevnt7Shift = 27
	RegisterSetb1rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 External Event 7
func (r *registerSetb1rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 External Event 7
func (r *registerSetb1rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldExtevnt7Mask)
	}
}

const (
	RegisterSetb1rFieldExtevnt6Shift = 26
	RegisterSetb1rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 External Event 6
func (r *registerSetb1rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 External Event 6
func (r *registerSetb1rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldExtevnt6Mask)
	}
}

const (
	RegisterSetb1rFieldExtevnt5Shift = 25
	RegisterSetb1rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 External Event 5
func (r *registerSetb1rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 External Event 5
func (r *registerSetb1rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldExtevnt5Mask)
	}
}

const (
	RegisterSetb1rFieldExtevnt4Shift = 24
	RegisterSetb1rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 External Event 4
func (r *registerSetb1rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 External Event 4
func (r *registerSetb1rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldExtevnt4Mask)
	}
}

const (
	RegisterSetb1rFieldExtevnt3Shift = 23
	RegisterSetb1rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 External Event 3
func (r *registerSetb1rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 External Event 3
func (r *registerSetb1rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldExtevnt3Mask)
	}
}

const (
	RegisterSetb1rFieldExtevnt2Shift = 22
	RegisterSetb1rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 External Event 2
func (r *registerSetb1rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 External Event 2
func (r *registerSetb1rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldExtevnt2Mask)
	}
}

const (
	RegisterSetb1rFieldExtevnt1Shift = 21
	RegisterSetb1rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 External Event 1
func (r *registerSetb1rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 External Event 1
func (r *registerSetb1rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldExtevnt1Mask)
	}
}

const (
	RegisterSetb1rFieldTimevnt9Shift = 20
	RegisterSetb1rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 Timer Event 9
func (r *registerSetb1rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 Timer Event 9
func (r *registerSetb1rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldTimevnt9Mask)
	}
}

const (
	RegisterSetb1rFieldTimevnt8Shift = 19
	RegisterSetb1rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 Timer Event 8
func (r *registerSetb1rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 Timer Event 8
func (r *registerSetb1rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldTimevnt8Mask)
	}
}

const (
	RegisterSetb1rFieldTimevnt7Shift = 18
	RegisterSetb1rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 Timer Event 7
func (r *registerSetb1rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 Timer Event 7
func (r *registerSetb1rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldTimevnt7Mask)
	}
}

const (
	RegisterSetb1rFieldTimevnt6Shift = 17
	RegisterSetb1rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 Timer Event 6
func (r *registerSetb1rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 Timer Event 6
func (r *registerSetb1rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldTimevnt6Mask)
	}
}

const (
	RegisterSetb1rFieldTimevnt5Shift = 16
	RegisterSetb1rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 Timer Event 5
func (r *registerSetb1rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 Timer Event 5
func (r *registerSetb1rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldTimevnt5Mask)
	}
}

const (
	RegisterSetb1rFieldTimevnt4Shift = 15
	RegisterSetb1rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 Timer Event 4
func (r *registerSetb1rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 Timer Event 4
func (r *registerSetb1rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldTimevnt4Mask)
	}
}

const (
	RegisterSetb1rFieldTimevnt3Shift = 14
	RegisterSetb1rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 Timer Event 3
func (r *registerSetb1rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 Timer Event 3
func (r *registerSetb1rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldTimevnt3Mask)
	}
}

const (
	RegisterSetb1rFieldTimevnt2Shift = 13
	RegisterSetb1rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 Timer Event 2
func (r *registerSetb1rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 Timer Event 2
func (r *registerSetb1rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldTimevnt2Mask)
	}
}

const (
	RegisterSetb1rFieldTimevnt1Shift = 12
	RegisterSetb1rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 Timer Event 1
func (r *registerSetb1rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 Timer Event 1
func (r *registerSetb1rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldTimevnt1Mask)
	}
}

const (
	RegisterSetb1rFieldMstcmp4Shift = 11
	RegisterSetb1rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 Master Compare 4
func (r *registerSetb1rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 Master Compare 4
func (r *registerSetb1rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldMstcmp4Mask)
	}
}

const (
	RegisterSetb1rFieldMstcmp3Shift = 10
	RegisterSetb1rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 Master Compare 3
func (r *registerSetb1rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 Master Compare 3
func (r *registerSetb1rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldMstcmp3Mask)
	}
}

const (
	RegisterSetb1rFieldMstcmp2Shift = 9
	RegisterSetb1rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 Master Compare 2
func (r *registerSetb1rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 Master Compare 2
func (r *registerSetb1rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldMstcmp2Mask)
	}
}

const (
	RegisterSetb1rFieldMstcmp1Shift = 8
	RegisterSetb1rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 Master Compare 1
func (r *registerSetb1rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 Master Compare 1
func (r *registerSetb1rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldMstcmp1Mask)
	}
}

const (
	RegisterSetb1rFieldMstperShift = 7
	RegisterSetb1rFieldMstperMask  = 0x80
)

// GetMstper Master Period
func (r *registerSetb1rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldMstperMask) != 0
}

// SetMstper Master Period
func (r *registerSetb1rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldMstperMask)
	}
}

const (
	RegisterSetb1rFieldCmp4Shift = 6
	RegisterSetb1rFieldCmp4Mask  = 0x40
)

// GetCmp4 Timer A compare 4
func (r *registerSetb1rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldCmp4Mask) != 0
}

// SetCmp4 Timer A compare 4
func (r *registerSetb1rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldCmp4Mask)
	}
}

const (
	RegisterSetb1rFieldCmp3Shift = 5
	RegisterSetb1rFieldCmp3Mask  = 0x20
)

// GetCmp3 Timer A compare 3
func (r *registerSetb1rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldCmp3Mask) != 0
}

// SetCmp3 Timer A compare 3
func (r *registerSetb1rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldCmp3Mask)
	}
}

const (
	RegisterSetb1rFieldCmp2Shift = 4
	RegisterSetb1rFieldCmp2Mask  = 0x10
)

// GetCmp2 Timer A compare 2
func (r *registerSetb1rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldCmp2Mask) != 0
}

// SetCmp2 Timer A compare 2
func (r *registerSetb1rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldCmp2Mask)
	}
}

const (
	RegisterSetb1rFieldCmp1Shift = 3
	RegisterSetb1rFieldCmp1Mask  = 0x8
)

// GetCmp1 Timer A compare 1
func (r *registerSetb1rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldCmp1Mask) != 0
}

// SetCmp1 Timer A compare 1
func (r *registerSetb1rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldCmp1Mask)
	}
}

const (
	RegisterSetb1rFieldPerShift = 2
	RegisterSetb1rFieldPerMask  = 0x4
)

// GetPer Timer A Period
func (r *registerSetb1rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldPerMask) != 0
}

// SetPer Timer A Period
func (r *registerSetb1rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldPerMask)
	}
}

const (
	RegisterSetb1rFieldResyncShift = 1
	RegisterSetb1rFieldResyncMask  = 0x2
)

// GetResync Timer A resynchronizaton
func (r *registerSetb1rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldResyncMask) != 0
}

// SetResync Timer A resynchronizaton
func (r *registerSetb1rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldResyncMask)
	}
}

const (
	RegisterSetb1rFieldSstShift = 0
	RegisterSetb1rFieldSstMask  = 0x1
)

// GetSst Software Set trigger
func (r *registerSetb1rType) GetSst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb1rFieldSstMask) != 0
}

// SetSst Software Set trigger
func (r *registerSetb1rType) SetSst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb1rFieldSstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb1rFieldSstMask)
	}
}

// registerRstb1rType Timerx Output1 Reset Register
type registerRstb1rType uint32

const (
	RegisterRstb1rFieldUpdateShift = 31
	RegisterRstb1rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerRstb1rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerRstb1rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldUpdateMask)
	}
}

const (
	RegisterRstb1rFieldExtevnt10Shift = 30
	RegisterRstb1rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerRstb1rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerRstb1rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldExtevnt10Mask)
	}
}

const (
	RegisterRstb1rFieldExtevnt9Shift = 29
	RegisterRstb1rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerRstb1rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerRstb1rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldExtevnt9Mask)
	}
}

const (
	RegisterRstb1rFieldExtevnt8Shift = 28
	RegisterRstb1rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerRstb1rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerRstb1rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldExtevnt8Mask)
	}
}

const (
	RegisterRstb1rFieldExtevnt7Shift = 27
	RegisterRstb1rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerRstb1rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerRstb1rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldExtevnt7Mask)
	}
}

const (
	RegisterRstb1rFieldExtevnt6Shift = 26
	RegisterRstb1rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerRstb1rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerRstb1rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldExtevnt6Mask)
	}
}

const (
	RegisterRstb1rFieldExtevnt5Shift = 25
	RegisterRstb1rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerRstb1rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerRstb1rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldExtevnt5Mask)
	}
}

const (
	RegisterRstb1rFieldExtevnt4Shift = 24
	RegisterRstb1rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerRstb1rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerRstb1rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldExtevnt4Mask)
	}
}

const (
	RegisterRstb1rFieldExtevnt3Shift = 23
	RegisterRstb1rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerRstb1rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerRstb1rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldExtevnt3Mask)
	}
}

const (
	RegisterRstb1rFieldExtevnt2Shift = 22
	RegisterRstb1rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerRstb1rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerRstb1rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldExtevnt2Mask)
	}
}

const (
	RegisterRstb1rFieldExtevnt1Shift = 21
	RegisterRstb1rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerRstb1rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerRstb1rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldExtevnt1Mask)
	}
}

const (
	RegisterRstb1rFieldTimevnt9Shift = 20
	RegisterRstb1rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerRstb1rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerRstb1rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldTimevnt9Mask)
	}
}

const (
	RegisterRstb1rFieldTimevnt8Shift = 19
	RegisterRstb1rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerRstb1rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerRstb1rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldTimevnt8Mask)
	}
}

const (
	RegisterRstb1rFieldTimevnt7Shift = 18
	RegisterRstb1rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerRstb1rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerRstb1rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldTimevnt7Mask)
	}
}

const (
	RegisterRstb1rFieldTimevnt6Shift = 17
	RegisterRstb1rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerRstb1rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerRstb1rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldTimevnt6Mask)
	}
}

const (
	RegisterRstb1rFieldTimevnt5Shift = 16
	RegisterRstb1rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerRstb1rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerRstb1rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldTimevnt5Mask)
	}
}

const (
	RegisterRstb1rFieldTimevnt4Shift = 15
	RegisterRstb1rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerRstb1rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerRstb1rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldTimevnt4Mask)
	}
}

const (
	RegisterRstb1rFieldTimevnt3Shift = 14
	RegisterRstb1rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerRstb1rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerRstb1rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldTimevnt3Mask)
	}
}

const (
	RegisterRstb1rFieldTimevnt2Shift = 13
	RegisterRstb1rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerRstb1rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerRstb1rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldTimevnt2Mask)
	}
}

const (
	RegisterRstb1rFieldTimevnt1Shift = 12
	RegisterRstb1rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerRstb1rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerRstb1rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldTimevnt1Mask)
	}
}

const (
	RegisterRstb1rFieldMstcmp4Shift = 11
	RegisterRstb1rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerRstb1rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerRstb1rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldMstcmp4Mask)
	}
}

const (
	RegisterRstb1rFieldMstcmp3Shift = 10
	RegisterRstb1rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerRstb1rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerRstb1rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldMstcmp3Mask)
	}
}

const (
	RegisterRstb1rFieldMstcmp2Shift = 9
	RegisterRstb1rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerRstb1rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerRstb1rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldMstcmp2Mask)
	}
}

const (
	RegisterRstb1rFieldMstcmp1Shift = 8
	RegisterRstb1rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerRstb1rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerRstb1rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldMstcmp1Mask)
	}
}

const (
	RegisterRstb1rFieldMstperShift = 7
	RegisterRstb1rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerRstb1rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerRstb1rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldMstperMask)
	}
}

const (
	RegisterRstb1rFieldCmp4Shift = 6
	RegisterRstb1rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerRstb1rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerRstb1rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldCmp4Mask)
	}
}

const (
	RegisterRstb1rFieldCmp3Shift = 5
	RegisterRstb1rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerRstb1rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerRstb1rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldCmp3Mask)
	}
}

const (
	RegisterRstb1rFieldCmp2Shift = 4
	RegisterRstb1rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerRstb1rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerRstb1rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldCmp2Mask)
	}
}

const (
	RegisterRstb1rFieldCmp1Shift = 3
	RegisterRstb1rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerRstb1rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerRstb1rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldCmp1Mask)
	}
}

const (
	RegisterRstb1rFieldPerShift = 2
	RegisterRstb1rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerRstb1rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldPerMask) != 0
}

// SetPer PER
func (r *registerRstb1rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldPerMask)
	}
}

const (
	RegisterRstb1rFieldResyncShift = 1
	RegisterRstb1rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerRstb1rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerRstb1rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldResyncMask)
	}
}

const (
	RegisterRstb1rFieldSrtShift = 0
	RegisterRstb1rFieldSrtMask  = 0x1
)

// GetSrt SRT
func (r *registerRstb1rType) GetSrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb1rFieldSrtMask) != 0
}

// SetSrt SRT
func (r *registerRstb1rType) SetSrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb1rFieldSrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb1rFieldSrtMask)
	}
}

// registerSetb2rType Timerx Output2 Set Register
type registerSetb2rType uint32

const (
	RegisterSetb2rFieldUpdateShift = 31
	RegisterSetb2rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerSetb2rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerSetb2rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldUpdateMask)
	}
}

const (
	RegisterSetb2rFieldExtevnt10Shift = 30
	RegisterSetb2rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerSetb2rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerSetb2rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldExtevnt10Mask)
	}
}

const (
	RegisterSetb2rFieldExtevnt9Shift = 29
	RegisterSetb2rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerSetb2rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerSetb2rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldExtevnt9Mask)
	}
}

const (
	RegisterSetb2rFieldExtevnt8Shift = 28
	RegisterSetb2rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerSetb2rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerSetb2rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldExtevnt8Mask)
	}
}

const (
	RegisterSetb2rFieldExtevnt7Shift = 27
	RegisterSetb2rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerSetb2rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerSetb2rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldExtevnt7Mask)
	}
}

const (
	RegisterSetb2rFieldExtevnt6Shift = 26
	RegisterSetb2rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerSetb2rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerSetb2rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldExtevnt6Mask)
	}
}

const (
	RegisterSetb2rFieldExtevnt5Shift = 25
	RegisterSetb2rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerSetb2rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerSetb2rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldExtevnt5Mask)
	}
}

const (
	RegisterSetb2rFieldExtevnt4Shift = 24
	RegisterSetb2rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerSetb2rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerSetb2rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldExtevnt4Mask)
	}
}

const (
	RegisterSetb2rFieldExtevnt3Shift = 23
	RegisterSetb2rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerSetb2rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerSetb2rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldExtevnt3Mask)
	}
}

const (
	RegisterSetb2rFieldExtevnt2Shift = 22
	RegisterSetb2rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerSetb2rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerSetb2rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldExtevnt2Mask)
	}
}

const (
	RegisterSetb2rFieldExtevnt1Shift = 21
	RegisterSetb2rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerSetb2rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerSetb2rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldExtevnt1Mask)
	}
}

const (
	RegisterSetb2rFieldTimevnt9Shift = 20
	RegisterSetb2rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerSetb2rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerSetb2rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldTimevnt9Mask)
	}
}

const (
	RegisterSetb2rFieldTimevnt8Shift = 19
	RegisterSetb2rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerSetb2rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerSetb2rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldTimevnt8Mask)
	}
}

const (
	RegisterSetb2rFieldTimevnt7Shift = 18
	RegisterSetb2rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerSetb2rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerSetb2rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldTimevnt7Mask)
	}
}

const (
	RegisterSetb2rFieldTimevnt6Shift = 17
	RegisterSetb2rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerSetb2rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerSetb2rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldTimevnt6Mask)
	}
}

const (
	RegisterSetb2rFieldTimevnt5Shift = 16
	RegisterSetb2rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerSetb2rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerSetb2rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldTimevnt5Mask)
	}
}

const (
	RegisterSetb2rFieldTimevnt4Shift = 15
	RegisterSetb2rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerSetb2rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerSetb2rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldTimevnt4Mask)
	}
}

const (
	RegisterSetb2rFieldTimevnt3Shift = 14
	RegisterSetb2rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerSetb2rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerSetb2rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldTimevnt3Mask)
	}
}

const (
	RegisterSetb2rFieldTimevnt2Shift = 13
	RegisterSetb2rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerSetb2rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerSetb2rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldTimevnt2Mask)
	}
}

const (
	RegisterSetb2rFieldTimevnt1Shift = 12
	RegisterSetb2rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerSetb2rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerSetb2rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldTimevnt1Mask)
	}
}

const (
	RegisterSetb2rFieldMstcmp4Shift = 11
	RegisterSetb2rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerSetb2rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerSetb2rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldMstcmp4Mask)
	}
}

const (
	RegisterSetb2rFieldMstcmp3Shift = 10
	RegisterSetb2rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerSetb2rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerSetb2rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldMstcmp3Mask)
	}
}

const (
	RegisterSetb2rFieldMstcmp2Shift = 9
	RegisterSetb2rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerSetb2rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerSetb2rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldMstcmp2Mask)
	}
}

const (
	RegisterSetb2rFieldMstcmp1Shift = 8
	RegisterSetb2rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerSetb2rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerSetb2rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldMstcmp1Mask)
	}
}

const (
	RegisterSetb2rFieldMstperShift = 7
	RegisterSetb2rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerSetb2rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerSetb2rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldMstperMask)
	}
}

const (
	RegisterSetb2rFieldCmp4Shift = 6
	RegisterSetb2rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerSetb2rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerSetb2rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldCmp4Mask)
	}
}

const (
	RegisterSetb2rFieldCmp3Shift = 5
	RegisterSetb2rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerSetb2rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerSetb2rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldCmp3Mask)
	}
}

const (
	RegisterSetb2rFieldCmp2Shift = 4
	RegisterSetb2rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerSetb2rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerSetb2rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldCmp2Mask)
	}
}

const (
	RegisterSetb2rFieldCmp1Shift = 3
	RegisterSetb2rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerSetb2rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerSetb2rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldCmp1Mask)
	}
}

const (
	RegisterSetb2rFieldPerShift = 2
	RegisterSetb2rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerSetb2rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldPerMask) != 0
}

// SetPer PER
func (r *registerSetb2rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldPerMask)
	}
}

const (
	RegisterSetb2rFieldResyncShift = 1
	RegisterSetb2rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerSetb2rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerSetb2rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldResyncMask)
	}
}

const (
	RegisterSetb2rFieldSstShift = 0
	RegisterSetb2rFieldSstMask  = 0x1
)

// GetSst SST
func (r *registerSetb2rType) GetSst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetb2rFieldSstMask) != 0
}

// SetSst SST
func (r *registerSetb2rType) SetSst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetb2rFieldSstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetb2rFieldSstMask)
	}
}

// registerRstb2rType Timerx Output2 Reset Register
type registerRstb2rType uint32

const (
	RegisterRstb2rFieldUpdateShift = 31
	RegisterRstb2rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerRstb2rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerRstb2rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldUpdateMask)
	}
}

const (
	RegisterRstb2rFieldExtevnt10Shift = 30
	RegisterRstb2rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerRstb2rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerRstb2rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldExtevnt10Mask)
	}
}

const (
	RegisterRstb2rFieldExtevnt9Shift = 29
	RegisterRstb2rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerRstb2rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerRstb2rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldExtevnt9Mask)
	}
}

const (
	RegisterRstb2rFieldExtevnt8Shift = 28
	RegisterRstb2rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerRstb2rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerRstb2rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldExtevnt8Mask)
	}
}

const (
	RegisterRstb2rFieldExtevnt7Shift = 27
	RegisterRstb2rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerRstb2rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerRstb2rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldExtevnt7Mask)
	}
}

const (
	RegisterRstb2rFieldExtevnt6Shift = 26
	RegisterRstb2rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerRstb2rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerRstb2rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldExtevnt6Mask)
	}
}

const (
	RegisterRstb2rFieldExtevnt5Shift = 25
	RegisterRstb2rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerRstb2rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerRstb2rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldExtevnt5Mask)
	}
}

const (
	RegisterRstb2rFieldExtevnt4Shift = 24
	RegisterRstb2rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerRstb2rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerRstb2rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldExtevnt4Mask)
	}
}

const (
	RegisterRstb2rFieldExtevnt3Shift = 23
	RegisterRstb2rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerRstb2rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerRstb2rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldExtevnt3Mask)
	}
}

const (
	RegisterRstb2rFieldExtevnt2Shift = 22
	RegisterRstb2rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerRstb2rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerRstb2rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldExtevnt2Mask)
	}
}

const (
	RegisterRstb2rFieldExtevnt1Shift = 21
	RegisterRstb2rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerRstb2rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerRstb2rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldExtevnt1Mask)
	}
}

const (
	RegisterRstb2rFieldTimevnt9Shift = 20
	RegisterRstb2rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerRstb2rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerRstb2rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldTimevnt9Mask)
	}
}

const (
	RegisterRstb2rFieldTimevnt8Shift = 19
	RegisterRstb2rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerRstb2rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerRstb2rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldTimevnt8Mask)
	}
}

const (
	RegisterRstb2rFieldTimevnt7Shift = 18
	RegisterRstb2rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerRstb2rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerRstb2rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldTimevnt7Mask)
	}
}

const (
	RegisterRstb2rFieldTimevnt6Shift = 17
	RegisterRstb2rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerRstb2rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerRstb2rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldTimevnt6Mask)
	}
}

const (
	RegisterRstb2rFieldTimevnt5Shift = 16
	RegisterRstb2rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerRstb2rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerRstb2rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldTimevnt5Mask)
	}
}

const (
	RegisterRstb2rFieldTimevnt4Shift = 15
	RegisterRstb2rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerRstb2rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerRstb2rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldTimevnt4Mask)
	}
}

const (
	RegisterRstb2rFieldTimevnt3Shift = 14
	RegisterRstb2rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerRstb2rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerRstb2rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldTimevnt3Mask)
	}
}

const (
	RegisterRstb2rFieldTimevnt2Shift = 13
	RegisterRstb2rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerRstb2rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerRstb2rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldTimevnt2Mask)
	}
}

const (
	RegisterRstb2rFieldTimevnt1Shift = 12
	RegisterRstb2rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerRstb2rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerRstb2rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldTimevnt1Mask)
	}
}

const (
	RegisterRstb2rFieldMstcmp4Shift = 11
	RegisterRstb2rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerRstb2rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerRstb2rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldMstcmp4Mask)
	}
}

const (
	RegisterRstb2rFieldMstcmp3Shift = 10
	RegisterRstb2rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerRstb2rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerRstb2rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldMstcmp3Mask)
	}
}

const (
	RegisterRstb2rFieldMstcmp2Shift = 9
	RegisterRstb2rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerRstb2rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerRstb2rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldMstcmp2Mask)
	}
}

const (
	RegisterRstb2rFieldMstcmp1Shift = 8
	RegisterRstb2rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerRstb2rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerRstb2rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldMstcmp1Mask)
	}
}

const (
	RegisterRstb2rFieldMstperShift = 7
	RegisterRstb2rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerRstb2rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerRstb2rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldMstperMask)
	}
}

const (
	RegisterRstb2rFieldCmp4Shift = 6
	RegisterRstb2rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerRstb2rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerRstb2rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldCmp4Mask)
	}
}

const (
	RegisterRstb2rFieldCmp3Shift = 5
	RegisterRstb2rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerRstb2rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerRstb2rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldCmp3Mask)
	}
}

const (
	RegisterRstb2rFieldCmp2Shift = 4
	RegisterRstb2rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerRstb2rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerRstb2rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldCmp2Mask)
	}
}

const (
	RegisterRstb2rFieldCmp1Shift = 3
	RegisterRstb2rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerRstb2rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerRstb2rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldCmp1Mask)
	}
}

const (
	RegisterRstb2rFieldPerShift = 2
	RegisterRstb2rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerRstb2rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldPerMask) != 0
}

// SetPer PER
func (r *registerRstb2rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldPerMask)
	}
}

const (
	RegisterRstb2rFieldResyncShift = 1
	RegisterRstb2rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerRstb2rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerRstb2rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldResyncMask)
	}
}

const (
	RegisterRstb2rFieldSrtShift = 0
	RegisterRstb2rFieldSrtMask  = 0x1
)

// GetSrt SRT
func (r *registerRstb2rType) GetSrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstb2rFieldSrtMask) != 0
}

// SetSrt SRT
func (r *registerRstb2rType) SetSrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstb2rFieldSrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstb2rFieldSrtMask)
	}
}

// registerEefbr1Type Timerx External Event Filtering Register 1
type registerEefbr1Type uint32

const (
	RegisterEefbr1FieldEe5fltrShift = 25
	RegisterEefbr1FieldEe5fltrMask  = 0x1e000000
)

// GetEe5fltr External Event 5 filter
func (r *registerEefbr1Type) GetEe5fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefbr1FieldEe5fltrMask) >> RegisterEefbr1FieldEe5fltrShift)
}

// SetEe5fltr External Event 5 filter
func (r *registerEefbr1Type) SetEe5fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefbr1FieldEe5fltrMask)|(uint32(value)<<RegisterEefbr1FieldEe5fltrShift))
}

const (
	RegisterEefbr1FieldEe5ltchShift = 24
	RegisterEefbr1FieldEe5ltchMask  = 0x1000000
)

// GetEe5ltch External Event 5 latch
func (r *registerEefbr1Type) GetEe5ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefbr1FieldEe5ltchMask) != 0
}

// SetEe5ltch External Event 5 latch
func (r *registerEefbr1Type) SetEe5ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefbr1FieldEe5ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefbr1FieldEe5ltchMask)
	}
}

const (
	RegisterEefbr1FieldEe4fltrShift = 19
	RegisterEefbr1FieldEe4fltrMask  = 0x780000
)

// GetEe4fltr External Event 4 filter
func (r *registerEefbr1Type) GetEe4fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefbr1FieldEe4fltrMask) >> RegisterEefbr1FieldEe4fltrShift)
}

// SetEe4fltr External Event 4 filter
func (r *registerEefbr1Type) SetEe4fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefbr1FieldEe4fltrMask)|(uint32(value)<<RegisterEefbr1FieldEe4fltrShift))
}

const (
	RegisterEefbr1FieldEe4ltchShift = 18
	RegisterEefbr1FieldEe4ltchMask  = 0x40000
)

// GetEe4ltch External Event 4 latch
func (r *registerEefbr1Type) GetEe4ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefbr1FieldEe4ltchMask) != 0
}

// SetEe4ltch External Event 4 latch
func (r *registerEefbr1Type) SetEe4ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefbr1FieldEe4ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefbr1FieldEe4ltchMask)
	}
}

const (
	RegisterEefbr1FieldEe3fltrShift = 13
	RegisterEefbr1FieldEe3fltrMask  = 0x1e000
)

// GetEe3fltr External Event 3 filter
func (r *registerEefbr1Type) GetEe3fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefbr1FieldEe3fltrMask) >> RegisterEefbr1FieldEe3fltrShift)
}

// SetEe3fltr External Event 3 filter
func (r *registerEefbr1Type) SetEe3fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefbr1FieldEe3fltrMask)|(uint32(value)<<RegisterEefbr1FieldEe3fltrShift))
}

const (
	RegisterEefbr1FieldEe3ltchShift = 12
	RegisterEefbr1FieldEe3ltchMask  = 0x1000
)

// GetEe3ltch External Event 3 latch
func (r *registerEefbr1Type) GetEe3ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefbr1FieldEe3ltchMask) != 0
}

// SetEe3ltch External Event 3 latch
func (r *registerEefbr1Type) SetEe3ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefbr1FieldEe3ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefbr1FieldEe3ltchMask)
	}
}

const (
	RegisterEefbr1FieldEe2fltrShift = 7
	RegisterEefbr1FieldEe2fltrMask  = 0x780
)

// GetEe2fltr External Event 2 filter
func (r *registerEefbr1Type) GetEe2fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefbr1FieldEe2fltrMask) >> RegisterEefbr1FieldEe2fltrShift)
}

// SetEe2fltr External Event 2 filter
func (r *registerEefbr1Type) SetEe2fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefbr1FieldEe2fltrMask)|(uint32(value)<<RegisterEefbr1FieldEe2fltrShift))
}

const (
	RegisterEefbr1FieldEe2ltchShift = 6
	RegisterEefbr1FieldEe2ltchMask  = 0x40
)

// GetEe2ltch External Event 2 latch
func (r *registerEefbr1Type) GetEe2ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefbr1FieldEe2ltchMask) != 0
}

// SetEe2ltch External Event 2 latch
func (r *registerEefbr1Type) SetEe2ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefbr1FieldEe2ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefbr1FieldEe2ltchMask)
	}
}

const (
	RegisterEefbr1FieldEe1fltrShift = 1
	RegisterEefbr1FieldEe1fltrMask  = 0x1e
)

// GetEe1fltr External Event 1 filter
func (r *registerEefbr1Type) GetEe1fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefbr1FieldEe1fltrMask) >> RegisterEefbr1FieldEe1fltrShift)
}

// SetEe1fltr External Event 1 filter
func (r *registerEefbr1Type) SetEe1fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefbr1FieldEe1fltrMask)|(uint32(value)<<RegisterEefbr1FieldEe1fltrShift))
}

const (
	RegisterEefbr1FieldEe1ltchShift = 0
	RegisterEefbr1FieldEe1ltchMask  = 0x1
)

// GetEe1ltch External Event 1 latch
func (r *registerEefbr1Type) GetEe1ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefbr1FieldEe1ltchMask) != 0
}

// SetEe1ltch External Event 1 latch
func (r *registerEefbr1Type) SetEe1ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefbr1FieldEe1ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefbr1FieldEe1ltchMask)
	}
}

// registerEefbr2Type Timerx External Event Filtering Register 2
type registerEefbr2Type uint32

const (
	RegisterEefbr2FieldEe10fltrShift = 25
	RegisterEefbr2FieldEe10fltrMask  = 0x1e000000
)

// GetEe10fltr External Event 10 filter
func (r *registerEefbr2Type) GetEe10fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefbr2FieldEe10fltrMask) >> RegisterEefbr2FieldEe10fltrShift)
}

// SetEe10fltr External Event 10 filter
func (r *registerEefbr2Type) SetEe10fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefbr2FieldEe10fltrMask)|(uint32(value)<<RegisterEefbr2FieldEe10fltrShift))
}

const (
	RegisterEefbr2FieldEe10ltchShift = 24
	RegisterEefbr2FieldEe10ltchMask  = 0x1000000
)

// GetEe10ltch External Event 10 latch
func (r *registerEefbr2Type) GetEe10ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefbr2FieldEe10ltchMask) != 0
}

// SetEe10ltch External Event 10 latch
func (r *registerEefbr2Type) SetEe10ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefbr2FieldEe10ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefbr2FieldEe10ltchMask)
	}
}

const (
	RegisterEefbr2FieldEe9fltrShift = 19
	RegisterEefbr2FieldEe9fltrMask  = 0x780000
)

// GetEe9fltr External Event 9 filter
func (r *registerEefbr2Type) GetEe9fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefbr2FieldEe9fltrMask) >> RegisterEefbr2FieldEe9fltrShift)
}

// SetEe9fltr External Event 9 filter
func (r *registerEefbr2Type) SetEe9fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefbr2FieldEe9fltrMask)|(uint32(value)<<RegisterEefbr2FieldEe9fltrShift))
}

const (
	RegisterEefbr2FieldEe9ltchShift = 18
	RegisterEefbr2FieldEe9ltchMask  = 0x40000
)

// GetEe9ltch External Event 9 latch
func (r *registerEefbr2Type) GetEe9ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefbr2FieldEe9ltchMask) != 0
}

// SetEe9ltch External Event 9 latch
func (r *registerEefbr2Type) SetEe9ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefbr2FieldEe9ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefbr2FieldEe9ltchMask)
	}
}

const (
	RegisterEefbr2FieldEe8fltrShift = 13
	RegisterEefbr2FieldEe8fltrMask  = 0x1e000
)

// GetEe8fltr External Event 8 filter
func (r *registerEefbr2Type) GetEe8fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefbr2FieldEe8fltrMask) >> RegisterEefbr2FieldEe8fltrShift)
}

// SetEe8fltr External Event 8 filter
func (r *registerEefbr2Type) SetEe8fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefbr2FieldEe8fltrMask)|(uint32(value)<<RegisterEefbr2FieldEe8fltrShift))
}

const (
	RegisterEefbr2FieldEe8ltchShift = 12
	RegisterEefbr2FieldEe8ltchMask  = 0x1000
)

// GetEe8ltch External Event 8 latch
func (r *registerEefbr2Type) GetEe8ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefbr2FieldEe8ltchMask) != 0
}

// SetEe8ltch External Event 8 latch
func (r *registerEefbr2Type) SetEe8ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefbr2FieldEe8ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefbr2FieldEe8ltchMask)
	}
}

const (
	RegisterEefbr2FieldEe7fltrShift = 7
	RegisterEefbr2FieldEe7fltrMask  = 0x780
)

// GetEe7fltr External Event 7 filter
func (r *registerEefbr2Type) GetEe7fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefbr2FieldEe7fltrMask) >> RegisterEefbr2FieldEe7fltrShift)
}

// SetEe7fltr External Event 7 filter
func (r *registerEefbr2Type) SetEe7fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefbr2FieldEe7fltrMask)|(uint32(value)<<RegisterEefbr2FieldEe7fltrShift))
}

const (
	RegisterEefbr2FieldEe7ltchShift = 6
	RegisterEefbr2FieldEe7ltchMask  = 0x40
)

// GetEe7ltch External Event 7 latch
func (r *registerEefbr2Type) GetEe7ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefbr2FieldEe7ltchMask) != 0
}

// SetEe7ltch External Event 7 latch
func (r *registerEefbr2Type) SetEe7ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefbr2FieldEe7ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefbr2FieldEe7ltchMask)
	}
}

const (
	RegisterEefbr2FieldEe6fltrShift = 1
	RegisterEefbr2FieldEe6fltrMask  = 0x1e
)

// GetEe6fltr External Event 6 filter
func (r *registerEefbr2Type) GetEe6fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefbr2FieldEe6fltrMask) >> RegisterEefbr2FieldEe6fltrShift)
}

// SetEe6fltr External Event 6 filter
func (r *registerEefbr2Type) SetEe6fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefbr2FieldEe6fltrMask)|(uint32(value)<<RegisterEefbr2FieldEe6fltrShift))
}

const (
	RegisterEefbr2FieldEe6ltchShift = 0
	RegisterEefbr2FieldEe6ltchMask  = 0x1
)

// GetEe6ltch External Event 6 latch
func (r *registerEefbr2Type) GetEe6ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefbr2FieldEe6ltchMask) != 0
}

// SetEe6ltch External Event 6 latch
func (r *registerEefbr2Type) SetEe6ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefbr2FieldEe6ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefbr2FieldEe6ltchMask)
	}
}

// registerRstbrType TimerA Reset Register
type registerRstbrType uint32

const (
	RegisterRstbrFieldTimecmp4Shift = 30
	RegisterRstbrFieldTimecmp4Mask  = 0x40000000
)

// GetTimecmp4 Timer E Compare 4
func (r *registerRstbrType) GetTimecmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldTimecmp4Mask) != 0
}

// SetTimecmp4 Timer E Compare 4
func (r *registerRstbrType) SetTimecmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldTimecmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldTimecmp4Mask)
	}
}

const (
	RegisterRstbrFieldTimecmp2Shift = 29
	RegisterRstbrFieldTimecmp2Mask  = 0x20000000
)

// GetTimecmp2 Timer E Compare 2
func (r *registerRstbrType) GetTimecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldTimecmp2Mask) != 0
}

// SetTimecmp2 Timer E Compare 2
func (r *registerRstbrType) SetTimecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldTimecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldTimecmp2Mask)
	}
}

const (
	RegisterRstbrFieldTimecmp1Shift = 28
	RegisterRstbrFieldTimecmp1Mask  = 0x10000000
)

// GetTimecmp1 Timer E Compare 1
func (r *registerRstbrType) GetTimecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldTimecmp1Mask) != 0
}

// SetTimecmp1 Timer E Compare 1
func (r *registerRstbrType) SetTimecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldTimecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldTimecmp1Mask)
	}
}

const (
	RegisterRstbrFieldTimdcmp4Shift = 27
	RegisterRstbrFieldTimdcmp4Mask  = 0x8000000
)

// GetTimdcmp4 Timer D Compare 4
func (r *registerRstbrType) GetTimdcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldTimdcmp4Mask) != 0
}

// SetTimdcmp4 Timer D Compare 4
func (r *registerRstbrType) SetTimdcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldTimdcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldTimdcmp4Mask)
	}
}

const (
	RegisterRstbrFieldTimdcmp2Shift = 26
	RegisterRstbrFieldTimdcmp2Mask  = 0x4000000
)

// GetTimdcmp2 Timer D Compare 2
func (r *registerRstbrType) GetTimdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldTimdcmp2Mask) != 0
}

// SetTimdcmp2 Timer D Compare 2
func (r *registerRstbrType) SetTimdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldTimdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldTimdcmp2Mask)
	}
}

const (
	RegisterRstbrFieldTimdcmp1Shift = 25
	RegisterRstbrFieldTimdcmp1Mask  = 0x2000000
)

// GetTimdcmp1 Timer D Compare 1
func (r *registerRstbrType) GetTimdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldTimdcmp1Mask) != 0
}

// SetTimdcmp1 Timer D Compare 1
func (r *registerRstbrType) SetTimdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldTimdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldTimdcmp1Mask)
	}
}

const (
	RegisterRstbrFieldTimccmp4Shift = 24
	RegisterRstbrFieldTimccmp4Mask  = 0x1000000
)

// GetTimccmp4 Timer C Compare 4
func (r *registerRstbrType) GetTimccmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldTimccmp4Mask) != 0
}

// SetTimccmp4 Timer C Compare 4
func (r *registerRstbrType) SetTimccmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldTimccmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldTimccmp4Mask)
	}
}

const (
	RegisterRstbrFieldTimccmp2Shift = 23
	RegisterRstbrFieldTimccmp2Mask  = 0x800000
)

// GetTimccmp2 Timer C Compare 2
func (r *registerRstbrType) GetTimccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldTimccmp2Mask) != 0
}

// SetTimccmp2 Timer C Compare 2
func (r *registerRstbrType) SetTimccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldTimccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldTimccmp2Mask)
	}
}

const (
	RegisterRstbrFieldTimccmp1Shift = 22
	RegisterRstbrFieldTimccmp1Mask  = 0x400000
)

// GetTimccmp1 Timer C Compare 1
func (r *registerRstbrType) GetTimccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldTimccmp1Mask) != 0
}

// SetTimccmp1 Timer C Compare 1
func (r *registerRstbrType) SetTimccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldTimccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldTimccmp1Mask)
	}
}

const (
	RegisterRstbrFieldTimacmp4Shift = 21
	RegisterRstbrFieldTimacmp4Mask  = 0x200000
)

// GetTimacmp4 Timer A Compare 4
func (r *registerRstbrType) GetTimacmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldTimacmp4Mask) != 0
}

// SetTimacmp4 Timer A Compare 4
func (r *registerRstbrType) SetTimacmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldTimacmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldTimacmp4Mask)
	}
}

const (
	RegisterRstbrFieldTimacmp2Shift = 20
	RegisterRstbrFieldTimacmp2Mask  = 0x100000
)

// GetTimacmp2 Timer A Compare 2
func (r *registerRstbrType) GetTimacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldTimacmp2Mask) != 0
}

// SetTimacmp2 Timer A Compare 2
func (r *registerRstbrType) SetTimacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldTimacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldTimacmp2Mask)
	}
}

const (
	RegisterRstbrFieldTimacmp1Shift = 19
	RegisterRstbrFieldTimacmp1Mask  = 0x80000
)

// GetTimacmp1 Timer A Compare 1
func (r *registerRstbrType) GetTimacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldTimacmp1Mask) != 0
}

// SetTimacmp1 Timer A Compare 1
func (r *registerRstbrType) SetTimacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldTimacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldTimacmp1Mask)
	}
}

const (
	RegisterRstbrFieldExtevnt10Shift = 18
	RegisterRstbrFieldExtevnt10Mask  = 0x40000
)

// GetExtevnt10 External Event 10
func (r *registerRstbrType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldExtevnt10Mask) != 0
}

// SetExtevnt10 External Event 10
func (r *registerRstbrType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldExtevnt10Mask)
	}
}

const (
	RegisterRstbrFieldExtevnt9Shift = 17
	RegisterRstbrFieldExtevnt9Mask  = 0x20000
)

// GetExtevnt9 External Event 9
func (r *registerRstbrType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldExtevnt9Mask) != 0
}

// SetExtevnt9 External Event 9
func (r *registerRstbrType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldExtevnt9Mask)
	}
}

const (
	RegisterRstbrFieldExtevnt8Shift = 16
	RegisterRstbrFieldExtevnt8Mask  = 0x10000
)

// GetExtevnt8 External Event 8
func (r *registerRstbrType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldExtevnt8Mask) != 0
}

// SetExtevnt8 External Event 8
func (r *registerRstbrType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldExtevnt8Mask)
	}
}

const (
	RegisterRstbrFieldExtevnt7Shift = 15
	RegisterRstbrFieldExtevnt7Mask  = 0x8000
)

// GetExtevnt7 External Event 7
func (r *registerRstbrType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldExtevnt7Mask) != 0
}

// SetExtevnt7 External Event 7
func (r *registerRstbrType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldExtevnt7Mask)
	}
}

const (
	RegisterRstbrFieldExtevnt6Shift = 14
	RegisterRstbrFieldExtevnt6Mask  = 0x4000
)

// GetExtevnt6 External Event 6
func (r *registerRstbrType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldExtevnt6Mask) != 0
}

// SetExtevnt6 External Event 6
func (r *registerRstbrType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldExtevnt6Mask)
	}
}

const (
	RegisterRstbrFieldExtevnt5Shift = 13
	RegisterRstbrFieldExtevnt5Mask  = 0x2000
)

// GetExtevnt5 External Event 5
func (r *registerRstbrType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldExtevnt5Mask) != 0
}

// SetExtevnt5 External Event 5
func (r *registerRstbrType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldExtevnt5Mask)
	}
}

const (
	RegisterRstbrFieldExtevnt4Shift = 12
	RegisterRstbrFieldExtevnt4Mask  = 0x1000
)

// GetExtevnt4 External Event 4
func (r *registerRstbrType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldExtevnt4Mask) != 0
}

// SetExtevnt4 External Event 4
func (r *registerRstbrType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldExtevnt4Mask)
	}
}

const (
	RegisterRstbrFieldExtevnt3Shift = 11
	RegisterRstbrFieldExtevnt3Mask  = 0x800
)

// GetExtevnt3 External Event 3
func (r *registerRstbrType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldExtevnt3Mask) != 0
}

// SetExtevnt3 External Event 3
func (r *registerRstbrType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldExtevnt3Mask)
	}
}

const (
	RegisterRstbrFieldExtevnt2Shift = 10
	RegisterRstbrFieldExtevnt2Mask  = 0x400
)

// GetExtevnt2 External Event 2
func (r *registerRstbrType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldExtevnt2Mask) != 0
}

// SetExtevnt2 External Event 2
func (r *registerRstbrType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldExtevnt2Mask)
	}
}

const (
	RegisterRstbrFieldExtevnt1Shift = 9
	RegisterRstbrFieldExtevnt1Mask  = 0x200
)

// GetExtevnt1 External Event 1
func (r *registerRstbrType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldExtevnt1Mask) != 0
}

// SetExtevnt1 External Event 1
func (r *registerRstbrType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldExtevnt1Mask)
	}
}

const (
	RegisterRstbrFieldMstcmp4Shift = 8
	RegisterRstbrFieldMstcmp4Mask  = 0x100
)

// GetMstcmp4 Master compare 4
func (r *registerRstbrType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldMstcmp4Mask) != 0
}

// SetMstcmp4 Master compare 4
func (r *registerRstbrType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldMstcmp4Mask)
	}
}

const (
	RegisterRstbrFieldMstcmp3Shift = 7
	RegisterRstbrFieldMstcmp3Mask  = 0x80
)

// GetMstcmp3 Master compare 3
func (r *registerRstbrType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldMstcmp3Mask) != 0
}

// SetMstcmp3 Master compare 3
func (r *registerRstbrType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldMstcmp3Mask)
	}
}

const (
	RegisterRstbrFieldMstcmp2Shift = 6
	RegisterRstbrFieldMstcmp2Mask  = 0x40
)

// GetMstcmp2 Master compare 2
func (r *registerRstbrType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldMstcmp2Mask) != 0
}

// SetMstcmp2 Master compare 2
func (r *registerRstbrType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldMstcmp2Mask)
	}
}

const (
	RegisterRstbrFieldMstcmp1Shift = 5
	RegisterRstbrFieldMstcmp1Mask  = 0x20
)

// GetMstcmp1 Master compare 1
func (r *registerRstbrType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldMstcmp1Mask) != 0
}

// SetMstcmp1 Master compare 1
func (r *registerRstbrType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldMstcmp1Mask)
	}
}

const (
	RegisterRstbrFieldMstperShift = 4
	RegisterRstbrFieldMstperMask  = 0x10
)

// GetMstper Master timer Period
func (r *registerRstbrType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldMstperMask) != 0
}

// SetMstper Master timer Period
func (r *registerRstbrType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldMstperMask)
	}
}

const (
	RegisterRstbrFieldCmp4Shift = 3
	RegisterRstbrFieldCmp4Mask  = 0x8
)

// GetCmp4 Timer A compare 4 reset
func (r *registerRstbrType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldCmp4Mask) != 0
}

// SetCmp4 Timer A compare 4 reset
func (r *registerRstbrType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldCmp4Mask)
	}
}

const (
	RegisterRstbrFieldCmp2Shift = 2
	RegisterRstbrFieldCmp2Mask  = 0x4
)

// GetCmp2 Timer A compare 2 reset
func (r *registerRstbrType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldCmp2Mask) != 0
}

// SetCmp2 Timer A compare 2 reset
func (r *registerRstbrType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldCmp2Mask)
	}
}

const (
	RegisterRstbrFieldUpdtShift = 1
	RegisterRstbrFieldUpdtMask  = 0x2
)

// GetUpdt Timer A Update reset
func (r *registerRstbrType) GetUpdt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstbrFieldUpdtMask) != 0
}

// SetUpdt Timer A Update reset
func (r *registerRstbrType) SetUpdt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstbrFieldUpdtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstbrFieldUpdtMask)
	}
}

// registerChpbrType Timerx Chopper Register
type registerChpbrType uint32

const (
	RegisterChpbrFieldStrtpwShift = 7
	RegisterChpbrFieldStrtpwMask  = 0x780
)

// GetStrtpw STRTPW
func (r *registerChpbrType) GetStrtpw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChpbrFieldStrtpwMask) >> RegisterChpbrFieldStrtpwShift)
}

// SetStrtpw STRTPW
func (r *registerChpbrType) SetStrtpw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChpbrFieldStrtpwMask)|(uint32(value)<<RegisterChpbrFieldStrtpwShift))
}

const (
	RegisterChpbrFieldChpdtyShift = 4
	RegisterChpbrFieldChpdtyMask  = 0x70
)

// GetChpdty Timerx chopper duty cycle value
func (r *registerChpbrType) GetChpdty() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChpbrFieldChpdtyMask) >> RegisterChpbrFieldChpdtyShift)
}

// SetChpdty Timerx chopper duty cycle value
func (r *registerChpbrType) SetChpdty(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChpbrFieldChpdtyMask)|(uint32(value)<<RegisterChpbrFieldChpdtyShift))
}

const (
	RegisterChpbrFieldChpfrqShift = 0
	RegisterChpbrFieldChpfrqMask  = 0xf
)

// GetChpfrq Timerx carrier frequency value
func (r *registerChpbrType) GetChpfrq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChpbrFieldChpfrqMask) >> RegisterChpbrFieldChpfrqShift)
}

// SetChpfrq Timerx carrier frequency value
func (r *registerChpbrType) SetChpfrq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChpbrFieldChpfrqMask)|(uint32(value)<<RegisterChpbrFieldChpfrqShift))
}

// registerCpt1bcrType Timerx Capture 2 Control Register
type registerCpt1bcrType uint32

const (
	RegisterCpt1bcrFieldTecmp2Shift = 31
	RegisterCpt1bcrFieldTecmp2Mask  = 0x80000000
)

// GetTecmp2 Timer E Compare 2
func (r *registerCpt1bcrType) GetTecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTecmp2Mask) != 0
}

// SetTecmp2 Timer E Compare 2
func (r *registerCpt1bcrType) SetTecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTecmp2Mask)
	}
}

const (
	RegisterCpt1bcrFieldTecmp1Shift = 30
	RegisterCpt1bcrFieldTecmp1Mask  = 0x40000000
)

// GetTecmp1 Timer E Compare 1
func (r *registerCpt1bcrType) GetTecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTecmp1Mask) != 0
}

// SetTecmp1 Timer E Compare 1
func (r *registerCpt1bcrType) SetTecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTecmp1Mask)
	}
}

const (
	RegisterCpt1bcrFieldTe1rstShift = 29
	RegisterCpt1bcrFieldTe1rstMask  = 0x20000000
)

// GetTe1rst Timer E output 1 Reset
func (r *registerCpt1bcrType) GetTe1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTe1rstMask) != 0
}

// SetTe1rst Timer E output 1 Reset
func (r *registerCpt1bcrType) SetTe1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTe1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTe1rstMask)
	}
}

const (
	RegisterCpt1bcrFieldTe1setShift = 28
	RegisterCpt1bcrFieldTe1setMask  = 0x10000000
)

// GetTe1set Timer E output 1 Set
func (r *registerCpt1bcrType) GetTe1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTe1setMask) != 0
}

// SetTe1set Timer E output 1 Set
func (r *registerCpt1bcrType) SetTe1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTe1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTe1setMask)
	}
}

const (
	RegisterCpt1bcrFieldTdcmp2Shift = 27
	RegisterCpt1bcrFieldTdcmp2Mask  = 0x8000000
)

// GetTdcmp2 Timer D Compare 2
func (r *registerCpt1bcrType) GetTdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTdcmp2Mask) != 0
}

// SetTdcmp2 Timer D Compare 2
func (r *registerCpt1bcrType) SetTdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTdcmp2Mask)
	}
}

const (
	RegisterCpt1bcrFieldTdcmp1Shift = 26
	RegisterCpt1bcrFieldTdcmp1Mask  = 0x4000000
)

// GetTdcmp1 Timer D Compare 1
func (r *registerCpt1bcrType) GetTdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTdcmp1Mask) != 0
}

// SetTdcmp1 Timer D Compare 1
func (r *registerCpt1bcrType) SetTdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTdcmp1Mask)
	}
}

const (
	RegisterCpt1bcrFieldTd1rstShift = 25
	RegisterCpt1bcrFieldTd1rstMask  = 0x2000000
)

// GetTd1rst Timer D output 1 Reset
func (r *registerCpt1bcrType) GetTd1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTd1rstMask) != 0
}

// SetTd1rst Timer D output 1 Reset
func (r *registerCpt1bcrType) SetTd1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTd1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTd1rstMask)
	}
}

const (
	RegisterCpt1bcrFieldTd1setShift = 24
	RegisterCpt1bcrFieldTd1setMask  = 0x1000000
)

// GetTd1set Timer D output 1 Set
func (r *registerCpt1bcrType) GetTd1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTd1setMask) != 0
}

// SetTd1set Timer D output 1 Set
func (r *registerCpt1bcrType) SetTd1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTd1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTd1setMask)
	}
}

const (
	RegisterCpt1bcrFieldTccmp2Shift = 23
	RegisterCpt1bcrFieldTccmp2Mask  = 0x800000
)

// GetTccmp2 Timer C Compare 2
func (r *registerCpt1bcrType) GetTccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTccmp2Mask) != 0
}

// SetTccmp2 Timer C Compare 2
func (r *registerCpt1bcrType) SetTccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTccmp2Mask)
	}
}

const (
	RegisterCpt1bcrFieldTccmp1Shift = 22
	RegisterCpt1bcrFieldTccmp1Mask  = 0x400000
)

// GetTccmp1 Timer C Compare 1
func (r *registerCpt1bcrType) GetTccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTccmp1Mask) != 0
}

// SetTccmp1 Timer C Compare 1
func (r *registerCpt1bcrType) SetTccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTccmp1Mask)
	}
}

const (
	RegisterCpt1bcrFieldTc1rstShift = 21
	RegisterCpt1bcrFieldTc1rstMask  = 0x200000
)

// GetTc1rst Timer C output 1 Reset
func (r *registerCpt1bcrType) GetTc1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTc1rstMask) != 0
}

// SetTc1rst Timer C output 1 Reset
func (r *registerCpt1bcrType) SetTc1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTc1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTc1rstMask)
	}
}

const (
	RegisterCpt1bcrFieldTc1setShift = 20
	RegisterCpt1bcrFieldTc1setMask  = 0x100000
)

// GetTc1set Timer C output 1 Set
func (r *registerCpt1bcrType) GetTc1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTc1setMask) != 0
}

// SetTc1set Timer C output 1 Set
func (r *registerCpt1bcrType) SetTc1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTc1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTc1setMask)
	}
}

const (
	RegisterCpt1bcrFieldTacmp2Shift = 15
	RegisterCpt1bcrFieldTacmp2Mask  = 0x8000
)

// GetTacmp2 Timer A Compare 2
func (r *registerCpt1bcrType) GetTacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTacmp2Mask) != 0
}

// SetTacmp2 Timer A Compare 2
func (r *registerCpt1bcrType) SetTacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTacmp2Mask)
	}
}

const (
	RegisterCpt1bcrFieldTacmp1Shift = 14
	RegisterCpt1bcrFieldTacmp1Mask  = 0x4000
)

// GetTacmp1 Timer A Compare 1
func (r *registerCpt1bcrType) GetTacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTacmp1Mask) != 0
}

// SetTacmp1 Timer A Compare 1
func (r *registerCpt1bcrType) SetTacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTacmp1Mask)
	}
}

const (
	RegisterCpt1bcrFieldTa1rstShift = 13
	RegisterCpt1bcrFieldTa1rstMask  = 0x2000
)

// GetTa1rst Timer A output 1 Reset
func (r *registerCpt1bcrType) GetTa1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTa1rstMask) != 0
}

// SetTa1rst Timer A output 1 Reset
func (r *registerCpt1bcrType) SetTa1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTa1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTa1rstMask)
	}
}

const (
	RegisterCpt1bcrFieldTa1setShift = 12
	RegisterCpt1bcrFieldTa1setMask  = 0x1000
)

// GetTa1set Timer A output 1 Set
func (r *registerCpt1bcrType) GetTa1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldTa1setMask) != 0
}

// SetTa1set Timer A output 1 Set
func (r *registerCpt1bcrType) SetTa1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldTa1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldTa1setMask)
	}
}

const (
	RegisterCpt1bcrFieldExev10cptShift = 11
	RegisterCpt1bcrFieldExev10cptMask  = 0x800
)

// GetExev10cpt External Event 10 Capture
func (r *registerCpt1bcrType) GetExev10cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldExev10cptMask) != 0
}

// SetExev10cpt External Event 10 Capture
func (r *registerCpt1bcrType) SetExev10cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldExev10cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldExev10cptMask)
	}
}

const (
	RegisterCpt1bcrFieldExev9cptShift = 10
	RegisterCpt1bcrFieldExev9cptMask  = 0x400
)

// GetExev9cpt External Event 9 Capture
func (r *registerCpt1bcrType) GetExev9cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldExev9cptMask) != 0
}

// SetExev9cpt External Event 9 Capture
func (r *registerCpt1bcrType) SetExev9cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldExev9cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldExev9cptMask)
	}
}

const (
	RegisterCpt1bcrFieldExev8cptShift = 9
	RegisterCpt1bcrFieldExev8cptMask  = 0x200
)

// GetExev8cpt External Event 8 Capture
func (r *registerCpt1bcrType) GetExev8cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldExev8cptMask) != 0
}

// SetExev8cpt External Event 8 Capture
func (r *registerCpt1bcrType) SetExev8cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldExev8cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldExev8cptMask)
	}
}

const (
	RegisterCpt1bcrFieldExev7cptShift = 8
	RegisterCpt1bcrFieldExev7cptMask  = 0x100
)

// GetExev7cpt External Event 7 Capture
func (r *registerCpt1bcrType) GetExev7cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldExev7cptMask) != 0
}

// SetExev7cpt External Event 7 Capture
func (r *registerCpt1bcrType) SetExev7cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldExev7cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldExev7cptMask)
	}
}

const (
	RegisterCpt1bcrFieldExev6cptShift = 7
	RegisterCpt1bcrFieldExev6cptMask  = 0x80
)

// GetExev6cpt External Event 6 Capture
func (r *registerCpt1bcrType) GetExev6cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldExev6cptMask) != 0
}

// SetExev6cpt External Event 6 Capture
func (r *registerCpt1bcrType) SetExev6cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldExev6cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldExev6cptMask)
	}
}

const (
	RegisterCpt1bcrFieldExev5cptShift = 6
	RegisterCpt1bcrFieldExev5cptMask  = 0x40
)

// GetExev5cpt External Event 5 Capture
func (r *registerCpt1bcrType) GetExev5cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldExev5cptMask) != 0
}

// SetExev5cpt External Event 5 Capture
func (r *registerCpt1bcrType) SetExev5cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldExev5cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldExev5cptMask)
	}
}

const (
	RegisterCpt1bcrFieldExev4cptShift = 5
	RegisterCpt1bcrFieldExev4cptMask  = 0x20
)

// GetExev4cpt External Event 4 Capture
func (r *registerCpt1bcrType) GetExev4cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldExev4cptMask) != 0
}

// SetExev4cpt External Event 4 Capture
func (r *registerCpt1bcrType) SetExev4cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldExev4cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldExev4cptMask)
	}
}

const (
	RegisterCpt1bcrFieldExev3cptShift = 4
	RegisterCpt1bcrFieldExev3cptMask  = 0x10
)

// GetExev3cpt External Event 3 Capture
func (r *registerCpt1bcrType) GetExev3cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldExev3cptMask) != 0
}

// SetExev3cpt External Event 3 Capture
func (r *registerCpt1bcrType) SetExev3cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldExev3cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldExev3cptMask)
	}
}

const (
	RegisterCpt1bcrFieldExev2cptShift = 3
	RegisterCpt1bcrFieldExev2cptMask  = 0x8
)

// GetExev2cpt External Event 2 Capture
func (r *registerCpt1bcrType) GetExev2cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldExev2cptMask) != 0
}

// SetExev2cpt External Event 2 Capture
func (r *registerCpt1bcrType) SetExev2cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldExev2cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldExev2cptMask)
	}
}

const (
	RegisterCpt1bcrFieldExev1cptShift = 2
	RegisterCpt1bcrFieldExev1cptMask  = 0x4
)

// GetExev1cpt External Event 1 Capture
func (r *registerCpt1bcrType) GetExev1cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldExev1cptMask) != 0
}

// SetExev1cpt External Event 1 Capture
func (r *registerCpt1bcrType) SetExev1cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldExev1cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldExev1cptMask)
	}
}

const (
	RegisterCpt1bcrFieldUdpcptShift = 1
	RegisterCpt1bcrFieldUdpcptMask  = 0x2
)

// GetUdpcpt Update Capture
func (r *registerCpt1bcrType) GetUdpcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldUdpcptMask) != 0
}

// SetUdpcpt Update Capture
func (r *registerCpt1bcrType) SetUdpcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldUdpcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldUdpcptMask)
	}
}

const (
	RegisterCpt1bcrFieldSwcptShift = 0
	RegisterCpt1bcrFieldSwcptMask  = 0x1
)

// GetSwcpt Software Capture
func (r *registerCpt1bcrType) GetSwcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1bcrFieldSwcptMask) != 0
}

// SetSwcpt Software Capture
func (r *registerCpt1bcrType) SetSwcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1bcrFieldSwcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1bcrFieldSwcptMask)
	}
}

// registerCpt2bcrType CPT2xCR
type registerCpt2bcrType uint32

const (
	RegisterCpt2bcrFieldTecmp2Shift = 31
	RegisterCpt2bcrFieldTecmp2Mask  = 0x80000000
)

// GetTecmp2 Timer E Compare 2
func (r *registerCpt2bcrType) GetTecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTecmp2Mask) != 0
}

// SetTecmp2 Timer E Compare 2
func (r *registerCpt2bcrType) SetTecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTecmp2Mask)
	}
}

const (
	RegisterCpt2bcrFieldTecmp1Shift = 30
	RegisterCpt2bcrFieldTecmp1Mask  = 0x40000000
)

// GetTecmp1 Timer E Compare 1
func (r *registerCpt2bcrType) GetTecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTecmp1Mask) != 0
}

// SetTecmp1 Timer E Compare 1
func (r *registerCpt2bcrType) SetTecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTecmp1Mask)
	}
}

const (
	RegisterCpt2bcrFieldTe1rstShift = 29
	RegisterCpt2bcrFieldTe1rstMask  = 0x20000000
)

// GetTe1rst Timer E output 1 Reset
func (r *registerCpt2bcrType) GetTe1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTe1rstMask) != 0
}

// SetTe1rst Timer E output 1 Reset
func (r *registerCpt2bcrType) SetTe1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTe1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTe1rstMask)
	}
}

const (
	RegisterCpt2bcrFieldTe1setShift = 28
	RegisterCpt2bcrFieldTe1setMask  = 0x10000000
)

// GetTe1set Timer E output 1 Set
func (r *registerCpt2bcrType) GetTe1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTe1setMask) != 0
}

// SetTe1set Timer E output 1 Set
func (r *registerCpt2bcrType) SetTe1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTe1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTe1setMask)
	}
}

const (
	RegisterCpt2bcrFieldTdcmp2Shift = 27
	RegisterCpt2bcrFieldTdcmp2Mask  = 0x8000000
)

// GetTdcmp2 Timer D Compare 2
func (r *registerCpt2bcrType) GetTdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTdcmp2Mask) != 0
}

// SetTdcmp2 Timer D Compare 2
func (r *registerCpt2bcrType) SetTdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTdcmp2Mask)
	}
}

const (
	RegisterCpt2bcrFieldTdcmp1Shift = 26
	RegisterCpt2bcrFieldTdcmp1Mask  = 0x4000000
)

// GetTdcmp1 Timer D Compare 1
func (r *registerCpt2bcrType) GetTdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTdcmp1Mask) != 0
}

// SetTdcmp1 Timer D Compare 1
func (r *registerCpt2bcrType) SetTdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTdcmp1Mask)
	}
}

const (
	RegisterCpt2bcrFieldTd1rstShift = 25
	RegisterCpt2bcrFieldTd1rstMask  = 0x2000000
)

// GetTd1rst Timer D output 1 Reset
func (r *registerCpt2bcrType) GetTd1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTd1rstMask) != 0
}

// SetTd1rst Timer D output 1 Reset
func (r *registerCpt2bcrType) SetTd1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTd1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTd1rstMask)
	}
}

const (
	RegisterCpt2bcrFieldTd1setShift = 24
	RegisterCpt2bcrFieldTd1setMask  = 0x1000000
)

// GetTd1set Timer D output 1 Set
func (r *registerCpt2bcrType) GetTd1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTd1setMask) != 0
}

// SetTd1set Timer D output 1 Set
func (r *registerCpt2bcrType) SetTd1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTd1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTd1setMask)
	}
}

const (
	RegisterCpt2bcrFieldTccmp2Shift = 23
	RegisterCpt2bcrFieldTccmp2Mask  = 0x800000
)

// GetTccmp2 Timer C Compare 2
func (r *registerCpt2bcrType) GetTccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTccmp2Mask) != 0
}

// SetTccmp2 Timer C Compare 2
func (r *registerCpt2bcrType) SetTccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTccmp2Mask)
	}
}

const (
	RegisterCpt2bcrFieldTccmp1Shift = 22
	RegisterCpt2bcrFieldTccmp1Mask  = 0x400000
)

// GetTccmp1 Timer C Compare 1
func (r *registerCpt2bcrType) GetTccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTccmp1Mask) != 0
}

// SetTccmp1 Timer C Compare 1
func (r *registerCpt2bcrType) SetTccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTccmp1Mask)
	}
}

const (
	RegisterCpt2bcrFieldTc1rstShift = 21
	RegisterCpt2bcrFieldTc1rstMask  = 0x200000
)

// GetTc1rst Timer C output 1 Reset
func (r *registerCpt2bcrType) GetTc1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTc1rstMask) != 0
}

// SetTc1rst Timer C output 1 Reset
func (r *registerCpt2bcrType) SetTc1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTc1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTc1rstMask)
	}
}

const (
	RegisterCpt2bcrFieldTc1setShift = 20
	RegisterCpt2bcrFieldTc1setMask  = 0x100000
)

// GetTc1set Timer C output 1 Set
func (r *registerCpt2bcrType) GetTc1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTc1setMask) != 0
}

// SetTc1set Timer C output 1 Set
func (r *registerCpt2bcrType) SetTc1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTc1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTc1setMask)
	}
}

const (
	RegisterCpt2bcrFieldTacmp2Shift = 15
	RegisterCpt2bcrFieldTacmp2Mask  = 0x8000
)

// GetTacmp2 Timer A Compare 2
func (r *registerCpt2bcrType) GetTacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTacmp2Mask) != 0
}

// SetTacmp2 Timer A Compare 2
func (r *registerCpt2bcrType) SetTacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTacmp2Mask)
	}
}

const (
	RegisterCpt2bcrFieldTacmp1Shift = 14
	RegisterCpt2bcrFieldTacmp1Mask  = 0x4000
)

// GetTacmp1 Timer A Compare 1
func (r *registerCpt2bcrType) GetTacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTacmp1Mask) != 0
}

// SetTacmp1 Timer A Compare 1
func (r *registerCpt2bcrType) SetTacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTacmp1Mask)
	}
}

const (
	RegisterCpt2bcrFieldTa1rstShift = 13
	RegisterCpt2bcrFieldTa1rstMask  = 0x2000
)

// GetTa1rst Timer A output 1 Reset
func (r *registerCpt2bcrType) GetTa1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTa1rstMask) != 0
}

// SetTa1rst Timer A output 1 Reset
func (r *registerCpt2bcrType) SetTa1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTa1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTa1rstMask)
	}
}

const (
	RegisterCpt2bcrFieldTa1setShift = 12
	RegisterCpt2bcrFieldTa1setMask  = 0x1000
)

// GetTa1set Timer A output 1 Set
func (r *registerCpt2bcrType) GetTa1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldTa1setMask) != 0
}

// SetTa1set Timer A output 1 Set
func (r *registerCpt2bcrType) SetTa1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldTa1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldTa1setMask)
	}
}

const (
	RegisterCpt2bcrFieldExev10cptShift = 11
	RegisterCpt2bcrFieldExev10cptMask  = 0x800
)

// GetExev10cpt External Event 10 Capture
func (r *registerCpt2bcrType) GetExev10cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldExev10cptMask) != 0
}

// SetExev10cpt External Event 10 Capture
func (r *registerCpt2bcrType) SetExev10cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldExev10cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldExev10cptMask)
	}
}

const (
	RegisterCpt2bcrFieldExev9cptShift = 10
	RegisterCpt2bcrFieldExev9cptMask  = 0x400
)

// GetExev9cpt External Event 9 Capture
func (r *registerCpt2bcrType) GetExev9cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldExev9cptMask) != 0
}

// SetExev9cpt External Event 9 Capture
func (r *registerCpt2bcrType) SetExev9cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldExev9cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldExev9cptMask)
	}
}

const (
	RegisterCpt2bcrFieldExev8cptShift = 9
	RegisterCpt2bcrFieldExev8cptMask  = 0x200
)

// GetExev8cpt External Event 8 Capture
func (r *registerCpt2bcrType) GetExev8cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldExev8cptMask) != 0
}

// SetExev8cpt External Event 8 Capture
func (r *registerCpt2bcrType) SetExev8cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldExev8cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldExev8cptMask)
	}
}

const (
	RegisterCpt2bcrFieldExev7cptShift = 8
	RegisterCpt2bcrFieldExev7cptMask  = 0x100
)

// GetExev7cpt External Event 7 Capture
func (r *registerCpt2bcrType) GetExev7cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldExev7cptMask) != 0
}

// SetExev7cpt External Event 7 Capture
func (r *registerCpt2bcrType) SetExev7cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldExev7cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldExev7cptMask)
	}
}

const (
	RegisterCpt2bcrFieldExev6cptShift = 7
	RegisterCpt2bcrFieldExev6cptMask  = 0x80
)

// GetExev6cpt External Event 6 Capture
func (r *registerCpt2bcrType) GetExev6cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldExev6cptMask) != 0
}

// SetExev6cpt External Event 6 Capture
func (r *registerCpt2bcrType) SetExev6cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldExev6cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldExev6cptMask)
	}
}

const (
	RegisterCpt2bcrFieldExev5cptShift = 6
	RegisterCpt2bcrFieldExev5cptMask  = 0x40
)

// GetExev5cpt External Event 5 Capture
func (r *registerCpt2bcrType) GetExev5cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldExev5cptMask) != 0
}

// SetExev5cpt External Event 5 Capture
func (r *registerCpt2bcrType) SetExev5cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldExev5cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldExev5cptMask)
	}
}

const (
	RegisterCpt2bcrFieldExev4cptShift = 5
	RegisterCpt2bcrFieldExev4cptMask  = 0x20
)

// GetExev4cpt External Event 4 Capture
func (r *registerCpt2bcrType) GetExev4cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldExev4cptMask) != 0
}

// SetExev4cpt External Event 4 Capture
func (r *registerCpt2bcrType) SetExev4cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldExev4cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldExev4cptMask)
	}
}

const (
	RegisterCpt2bcrFieldExev3cptShift = 4
	RegisterCpt2bcrFieldExev3cptMask  = 0x10
)

// GetExev3cpt External Event 3 Capture
func (r *registerCpt2bcrType) GetExev3cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldExev3cptMask) != 0
}

// SetExev3cpt External Event 3 Capture
func (r *registerCpt2bcrType) SetExev3cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldExev3cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldExev3cptMask)
	}
}

const (
	RegisterCpt2bcrFieldExev2cptShift = 3
	RegisterCpt2bcrFieldExev2cptMask  = 0x8
)

// GetExev2cpt External Event 2 Capture
func (r *registerCpt2bcrType) GetExev2cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldExev2cptMask) != 0
}

// SetExev2cpt External Event 2 Capture
func (r *registerCpt2bcrType) SetExev2cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldExev2cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldExev2cptMask)
	}
}

const (
	RegisterCpt2bcrFieldExev1cptShift = 2
	RegisterCpt2bcrFieldExev1cptMask  = 0x4
)

// GetExev1cpt External Event 1 Capture
func (r *registerCpt2bcrType) GetExev1cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldExev1cptMask) != 0
}

// SetExev1cpt External Event 1 Capture
func (r *registerCpt2bcrType) SetExev1cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldExev1cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldExev1cptMask)
	}
}

const (
	RegisterCpt2bcrFieldUdpcptShift = 1
	RegisterCpt2bcrFieldUdpcptMask  = 0x2
)

// GetUdpcpt Update Capture
func (r *registerCpt2bcrType) GetUdpcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldUdpcptMask) != 0
}

// SetUdpcpt Update Capture
func (r *registerCpt2bcrType) SetUdpcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldUdpcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldUdpcptMask)
	}
}

const (
	RegisterCpt2bcrFieldSwcptShift = 0
	RegisterCpt2bcrFieldSwcptMask  = 0x1
)

// GetSwcpt Software Capture
func (r *registerCpt2bcrType) GetSwcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2bcrFieldSwcptMask) != 0
}

// SetSwcpt Software Capture
func (r *registerCpt2bcrType) SetSwcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2bcrFieldSwcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2bcrFieldSwcptMask)
	}
}

// registerOutbrType Timerx Output Register
type registerOutbrType uint32

const (
	RegisterOutbrFieldDidl2Shift = 23
	RegisterOutbrFieldDidl2Mask  = 0x800000
)

// GetDidl2 Output 2 Deadtime upon burst mode Idle entry
func (r *registerOutbrType) GetDidl2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldDidl2Mask) != 0
}

// SetDidl2 Output 2 Deadtime upon burst mode Idle entry
func (r *registerOutbrType) SetDidl2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutbrFieldDidl2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldDidl2Mask)
	}
}

const (
	RegisterOutbrFieldChp2Shift = 22
	RegisterOutbrFieldChp2Mask  = 0x400000
)

// GetChp2 Output 2 Chopper enable
func (r *registerOutbrType) GetChp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldChp2Mask) != 0
}

// SetChp2 Output 2 Chopper enable
func (r *registerOutbrType) SetChp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutbrFieldChp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldChp2Mask)
	}
}

const (
	RegisterOutbrFieldFault2Shift = 20
	RegisterOutbrFieldFault2Mask  = 0x300000
)

// GetFault2 Output 2 Fault state
func (r *registerOutbrType) GetFault2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldFault2Mask) >> RegisterOutbrFieldFault2Shift)
}

// SetFault2 Output 2 Fault state
func (r *registerOutbrType) SetFault2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldFault2Mask)|(uint32(value)<<RegisterOutbrFieldFault2Shift))
}

const (
	RegisterOutbrFieldIdles2Shift = 19
	RegisterOutbrFieldIdles2Mask  = 0x80000
)

// GetIdles2 Output 2 Idle State
func (r *registerOutbrType) GetIdles2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldIdles2Mask) != 0
}

// SetIdles2 Output 2 Idle State
func (r *registerOutbrType) SetIdles2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutbrFieldIdles2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldIdles2Mask)
	}
}

const (
	RegisterOutbrFieldIdlem2Shift = 18
	RegisterOutbrFieldIdlem2Mask  = 0x40000
)

// GetIdlem2 Output 2 Idle mode
func (r *registerOutbrType) GetIdlem2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldIdlem2Mask) != 0
}

// SetIdlem2 Output 2 Idle mode
func (r *registerOutbrType) SetIdlem2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutbrFieldIdlem2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldIdlem2Mask)
	}
}

const (
	RegisterOutbrFieldPol2Shift = 17
	RegisterOutbrFieldPol2Mask  = 0x20000
)

// GetPol2 Output 2 polarity
func (r *registerOutbrType) GetPol2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldPol2Mask) != 0
}

// SetPol2 Output 2 polarity
func (r *registerOutbrType) SetPol2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutbrFieldPol2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldPol2Mask)
	}
}

const (
	RegisterOutbrFieldDlyprtShift = 10
	RegisterOutbrFieldDlyprtMask  = 0x1c00
)

// GetDlyprt Delayed Protection
func (r *registerOutbrType) GetDlyprt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldDlyprtMask) >> RegisterOutbrFieldDlyprtShift)
}

// SetDlyprt Delayed Protection
func (r *registerOutbrType) SetDlyprt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldDlyprtMask)|(uint32(value)<<RegisterOutbrFieldDlyprtShift))
}

const (
	RegisterOutbrFieldDlyprtenShift = 9
	RegisterOutbrFieldDlyprtenMask  = 0x200
)

// GetDlyprten Delayed Protection Enable
func (r *registerOutbrType) GetDlyprten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldDlyprtenMask) != 0
}

// SetDlyprten Delayed Protection Enable
func (r *registerOutbrType) SetDlyprten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutbrFieldDlyprtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldDlyprtenMask)
	}
}

const (
	RegisterOutbrFieldDtenShift = 8
	RegisterOutbrFieldDtenMask  = 0x100
)

// GetDten Deadtime enable
func (r *registerOutbrType) GetDten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldDtenMask) != 0
}

// SetDten Deadtime enable
func (r *registerOutbrType) SetDten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutbrFieldDtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldDtenMask)
	}
}

const (
	RegisterOutbrFieldDidl1Shift = 7
	RegisterOutbrFieldDidl1Mask  = 0x80
)

// GetDidl1 Output 1 Deadtime upon burst mode Idle entry
func (r *registerOutbrType) GetDidl1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldDidl1Mask) != 0
}

// SetDidl1 Output 1 Deadtime upon burst mode Idle entry
func (r *registerOutbrType) SetDidl1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutbrFieldDidl1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldDidl1Mask)
	}
}

const (
	RegisterOutbrFieldChp1Shift = 6
	RegisterOutbrFieldChp1Mask  = 0x40
)

// GetChp1 Output 1 Chopper enable
func (r *registerOutbrType) GetChp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldChp1Mask) != 0
}

// SetChp1 Output 1 Chopper enable
func (r *registerOutbrType) SetChp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutbrFieldChp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldChp1Mask)
	}
}

const (
	RegisterOutbrFieldFault1Shift = 4
	RegisterOutbrFieldFault1Mask  = 0x30
)

// GetFault1 Output 1 Fault state
func (r *registerOutbrType) GetFault1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldFault1Mask) >> RegisterOutbrFieldFault1Shift)
}

// SetFault1 Output 1 Fault state
func (r *registerOutbrType) SetFault1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldFault1Mask)|(uint32(value)<<RegisterOutbrFieldFault1Shift))
}

const (
	RegisterOutbrFieldIdles1Shift = 3
	RegisterOutbrFieldIdles1Mask  = 0x8
)

// GetIdles1 Output 1 Idle State
func (r *registerOutbrType) GetIdles1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldIdles1Mask) != 0
}

// SetIdles1 Output 1 Idle State
func (r *registerOutbrType) SetIdles1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutbrFieldIdles1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldIdles1Mask)
	}
}

const (
	RegisterOutbrFieldIdlem1Shift = 2
	RegisterOutbrFieldIdlem1Mask  = 0x4
)

// GetIdlem1 Output 1 Idle mode
func (r *registerOutbrType) GetIdlem1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldIdlem1Mask) != 0
}

// SetIdlem1 Output 1 Idle mode
func (r *registerOutbrType) SetIdlem1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutbrFieldIdlem1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldIdlem1Mask)
	}
}

const (
	RegisterOutbrFieldPol1Shift = 1
	RegisterOutbrFieldPol1Mask  = 0x2
)

// GetPol1 Output 1 polarity
func (r *registerOutbrType) GetPol1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutbrFieldPol1Mask) != 0
}

// SetPol1 Output 1 polarity
func (r *registerOutbrType) SetPol1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutbrFieldPol1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutbrFieldPol1Mask)
	}
}

// registerFltbrType Timerx Fault Register
type registerFltbrType uint32

const (
	RegisterFltbrFieldFltlckShift = 31
	RegisterFltbrFieldFltlckMask  = 0x80000000
)

// GetFltlck Fault sources Lock
func (r *registerFltbrType) GetFltlck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltbrFieldFltlckMask) != 0
}

// SetFltlck Fault sources Lock
func (r *registerFltbrType) SetFltlck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltbrFieldFltlckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltbrFieldFltlckMask)
	}
}

const (
	RegisterFltbrFieldFlt5enShift = 4
	RegisterFltbrFieldFlt5enMask  = 0x10
)

// GetFlt5en Fault 5 enable
func (r *registerFltbrType) GetFlt5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltbrFieldFlt5enMask) != 0
}

// SetFlt5en Fault 5 enable
func (r *registerFltbrType) SetFlt5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltbrFieldFlt5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltbrFieldFlt5enMask)
	}
}

const (
	RegisterFltbrFieldFlt4enShift = 3
	RegisterFltbrFieldFlt4enMask  = 0x8
)

// GetFlt4en Fault 4 enable
func (r *registerFltbrType) GetFlt4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltbrFieldFlt4enMask) != 0
}

// SetFlt4en Fault 4 enable
func (r *registerFltbrType) SetFlt4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltbrFieldFlt4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltbrFieldFlt4enMask)
	}
}

const (
	RegisterFltbrFieldFlt3enShift = 2
	RegisterFltbrFieldFlt3enMask  = 0x4
)

// GetFlt3en Fault 3 enable
func (r *registerFltbrType) GetFlt3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltbrFieldFlt3enMask) != 0
}

// SetFlt3en Fault 3 enable
func (r *registerFltbrType) SetFlt3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltbrFieldFlt3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltbrFieldFlt3enMask)
	}
}

const (
	RegisterFltbrFieldFlt2enShift = 1
	RegisterFltbrFieldFlt2enMask  = 0x2
)

// GetFlt2en Fault 2 enable
func (r *registerFltbrType) GetFlt2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltbrFieldFlt2enMask) != 0
}

// SetFlt2en Fault 2 enable
func (r *registerFltbrType) SetFlt2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltbrFieldFlt2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltbrFieldFlt2enMask)
	}
}

const (
	RegisterFltbrFieldFlt1enShift = 0
	RegisterFltbrFieldFlt1enMask  = 0x1
)

// GetFlt1en Fault 1 enable
func (r *registerFltbrType) GetFlt1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltbrFieldFlt1enMask) != 0
}

// SetFlt1en Fault 1 enable
func (r *registerFltbrType) SetFlt1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltbrFieldFlt1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltbrFieldFlt1enMask)
	}
}
