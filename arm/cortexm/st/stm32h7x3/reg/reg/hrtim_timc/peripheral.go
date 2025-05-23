package hrtim_timc

import (
	"unsafe"
	"volatile"
)

var (
	Hrtim_timc = (*_hrtim_timc)(unsafe.Pointer(uintptr(0x40017580)))
)

type _hrtim_timc struct {
	Timccr    registerTimccrType
	Timcisr   registerTimcisrType
	Timcicr   registerTimcicrType
	Timcdier5 registerTimcdier5Type
	Cntcr     registerCntcrType
	Percr     registerPercrType
	Repcr     registerRepcrType
	Cmp1cr    registerCmp1crType
	Cmp1ccr   registerCmp1ccrType
	Cmp2cr    registerCmp2crType
	Cmp3cr    registerCmp3crType
	Cmp4cr    registerCmp4crType
	Cpt1cr    registerCpt1crType
	Cpt2cr    registerCpt2crType
	Dtcr      registerDtcrType
	Setc1r    registerSetc1rType
	Rstc1r    registerRstc1rType
	Setc2r    registerSetc2rType
	Rstc2r    registerRstc2rType
	Eefcr1    registerEefcr1Type
	Eefcr2    registerEefcr2Type
	Rstcr     registerRstcrType
	Chpcr     registerChpcrType
	Cpt1ccr   registerCpt1ccrType
	Cpt2ccr   registerCpt2ccrType
	Outcr     registerOutcrType
	Fltcr     registerFltcrType
}

// registerTimccrType Timerx Control Register
type registerTimccrType uint32

const (
	RegisterTimccrFieldUpdgatShift = 28
	RegisterTimccrFieldUpdgatMask  = 0xf0000000
)

// GetUpdgat Update Gating
func (r *registerTimccrType) GetUpdgat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldUpdgatMask) >> RegisterTimccrFieldUpdgatShift)
}

// SetUpdgat Update Gating
func (r *registerTimccrType) SetUpdgat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldUpdgatMask)|(uint32(value)<<RegisterTimccrFieldUpdgatShift))
}

const (
	RegisterTimccrFieldPreenShift = 27
	RegisterTimccrFieldPreenMask  = 0x8000000
)

// GetPreen Preload enable
func (r *registerTimccrType) GetPreen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldPreenMask) != 0
}

// SetPreen Preload enable
func (r *registerTimccrType) SetPreen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldPreenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldPreenMask)
	}
}

const (
	RegisterTimccrFieldDacsyncShift = 25
	RegisterTimccrFieldDacsyncMask  = 0x6000000
)

// GetDacsync AC Synchronization
func (r *registerTimccrType) GetDacsync() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldDacsyncMask) >> RegisterTimccrFieldDacsyncShift)
}

// SetDacsync AC Synchronization
func (r *registerTimccrType) SetDacsync(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldDacsyncMask)|(uint32(value)<<RegisterTimccrFieldDacsyncShift))
}

const (
	RegisterTimccrFieldMstuShift = 24
	RegisterTimccrFieldMstuMask  = 0x1000000
)

// GetMstu Master Timer update
func (r *registerTimccrType) GetMstu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldMstuMask) != 0
}

// SetMstu Master Timer update
func (r *registerTimccrType) SetMstu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldMstuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldMstuMask)
	}
}

const (
	RegisterTimccrFieldTeuShift = 23
	RegisterTimccrFieldTeuMask  = 0x800000
)

// GetTeu TEU
func (r *registerTimccrType) GetTeu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldTeuMask) != 0
}

// SetTeu TEU
func (r *registerTimccrType) SetTeu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldTeuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldTeuMask)
	}
}

const (
	RegisterTimccrFieldTduShift = 22
	RegisterTimccrFieldTduMask  = 0x400000
)

// GetTdu TDU
func (r *registerTimccrType) GetTdu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldTduMask) != 0
}

// SetTdu TDU
func (r *registerTimccrType) SetTdu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldTduMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldTduMask)
	}
}

const (
	RegisterTimccrFieldTcuShift = 21
	RegisterTimccrFieldTcuMask  = 0x200000
)

// GetTcu TCU
func (r *registerTimccrType) GetTcu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldTcuMask) != 0
}

// SetTcu TCU
func (r *registerTimccrType) SetTcu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldTcuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldTcuMask)
	}
}

const (
	RegisterTimccrFieldTbuShift = 20
	RegisterTimccrFieldTbuMask  = 0x100000
)

// GetTbu TBU
func (r *registerTimccrType) GetTbu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldTbuMask) != 0
}

// SetTbu TBU
func (r *registerTimccrType) SetTbu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldTbuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldTbuMask)
	}
}

const (
	RegisterTimccrFieldTxrstuShift = 18
	RegisterTimccrFieldTxrstuMask  = 0x40000
)

// GetTxrstu Timerx reset update
func (r *registerTimccrType) GetTxrstu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldTxrstuMask) != 0
}

// SetTxrstu Timerx reset update
func (r *registerTimccrType) SetTxrstu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldTxrstuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldTxrstuMask)
	}
}

const (
	RegisterTimccrFieldTxrepuShift = 17
	RegisterTimccrFieldTxrepuMask  = 0x20000
)

// GetTxrepu Timer x Repetition update
func (r *registerTimccrType) GetTxrepu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldTxrepuMask) != 0
}

// SetTxrepu Timer x Repetition update
func (r *registerTimccrType) SetTxrepu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldTxrepuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldTxrepuMask)
	}
}

const (
	RegisterTimccrFieldDelcmp4Shift = 14
	RegisterTimccrFieldDelcmp4Mask  = 0xc000
)

// GetDelcmp4 Delayed CMP4 mode
func (r *registerTimccrType) GetDelcmp4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldDelcmp4Mask) >> RegisterTimccrFieldDelcmp4Shift)
}

// SetDelcmp4 Delayed CMP4 mode
func (r *registerTimccrType) SetDelcmp4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldDelcmp4Mask)|(uint32(value)<<RegisterTimccrFieldDelcmp4Shift))
}

const (
	RegisterTimccrFieldDelcmp2Shift = 12
	RegisterTimccrFieldDelcmp2Mask  = 0x3000
)

// GetDelcmp2 Delayed CMP2 mode
func (r *registerTimccrType) GetDelcmp2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldDelcmp2Mask) >> RegisterTimccrFieldDelcmp2Shift)
}

// SetDelcmp2 Delayed CMP2 mode
func (r *registerTimccrType) SetDelcmp2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldDelcmp2Mask)|(uint32(value)<<RegisterTimccrFieldDelcmp2Shift))
}

const (
	RegisterTimccrFieldSyncstrtxShift = 11
	RegisterTimccrFieldSyncstrtxMask  = 0x800
)

// GetSyncstrtx Synchronization Starts Timer x
func (r *registerTimccrType) GetSyncstrtx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldSyncstrtxMask) != 0
}

// SetSyncstrtx Synchronization Starts Timer x
func (r *registerTimccrType) SetSyncstrtx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldSyncstrtxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldSyncstrtxMask)
	}
}

const (
	RegisterTimccrFieldSyncrstxShift = 10
	RegisterTimccrFieldSyncrstxMask  = 0x400
)

// GetSyncrstx Synchronization Resets Timer x
func (r *registerTimccrType) GetSyncrstx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldSyncrstxMask) != 0
}

// SetSyncrstx Synchronization Resets Timer x
func (r *registerTimccrType) SetSyncrstx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldSyncrstxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldSyncrstxMask)
	}
}

const (
	RegisterTimccrFieldPshpllShift = 6
	RegisterTimccrFieldPshpllMask  = 0x40
)

// GetPshpll Push-Pull mode enable
func (r *registerTimccrType) GetPshpll() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldPshpllMask) != 0
}

// SetPshpll Push-Pull mode enable
func (r *registerTimccrType) SetPshpll(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldPshpllMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldPshpllMask)
	}
}

const (
	RegisterTimccrFieldHalfShift = 5
	RegisterTimccrFieldHalfMask  = 0x20
)

// GetHalf Half mode enable
func (r *registerTimccrType) GetHalf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldHalfMask) != 0
}

// SetHalf Half mode enable
func (r *registerTimccrType) SetHalf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldHalfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldHalfMask)
	}
}

const (
	RegisterTimccrFieldRetrigShift = 4
	RegisterTimccrFieldRetrigMask  = 0x10
)

// GetRetrig Re-triggerable mode
func (r *registerTimccrType) GetRetrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldRetrigMask) != 0
}

// SetRetrig Re-triggerable mode
func (r *registerTimccrType) SetRetrig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldRetrigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldRetrigMask)
	}
}

const (
	RegisterTimccrFieldContShift = 3
	RegisterTimccrFieldContMask  = 0x8
)

// GetCont Continuous mode
func (r *registerTimccrType) GetCont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldContMask) != 0
}

// SetCont Continuous mode
func (r *registerTimccrType) SetCont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimccrFieldContMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldContMask)
	}
}

const (
	RegisterTimccrFieldCk_pscxShift = 0
	RegisterTimccrFieldCk_pscxMask  = 0x7
)

// GetCk_pscx HRTIM Timer x Clock prescaler
func (r *registerTimccrType) GetCk_pscx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimccrFieldCk_pscxMask) >> RegisterTimccrFieldCk_pscxShift)
}

// SetCk_pscx HRTIM Timer x Clock prescaler
func (r *registerTimccrType) SetCk_pscx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimccrFieldCk_pscxMask)|(uint32(value)<<RegisterTimccrFieldCk_pscxShift))
}

// registerTimcisrType Timerx Interrupt Status Register
type registerTimcisrType uint32

const (
	RegisterTimcisrFieldO2statShift = 19
	RegisterTimcisrFieldO2statMask  = 0x80000
)

// GetO2stat Output 2 State
func (r *registerTimcisrType) GetO2stat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldO2statMask) != 0
}

// SetO2stat Output 2 State
func (r *registerTimcisrType) SetO2stat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldO2statMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldO2statMask)
	}
}

const (
	RegisterTimcisrFieldO1statShift = 18
	RegisterTimcisrFieldO1statMask  = 0x40000
)

// GetO1stat Output 1 State
func (r *registerTimcisrType) GetO1stat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldO1statMask) != 0
}

// SetO1stat Output 1 State
func (r *registerTimcisrType) SetO1stat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldO1statMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldO1statMask)
	}
}

const (
	RegisterTimcisrFieldIppstatShift = 17
	RegisterTimcisrFieldIppstatMask  = 0x20000
)

// GetIppstat Idle Push Pull Status
func (r *registerTimcisrType) GetIppstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldIppstatMask) != 0
}

// SetIppstat Idle Push Pull Status
func (r *registerTimcisrType) SetIppstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldIppstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldIppstatMask)
	}
}

const (
	RegisterTimcisrFieldCppstatShift = 16
	RegisterTimcisrFieldCppstatMask  = 0x10000
)

// GetCppstat Current Push Pull Status
func (r *registerTimcisrType) GetCppstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldCppstatMask) != 0
}

// SetCppstat Current Push Pull Status
func (r *registerTimcisrType) SetCppstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldCppstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldCppstatMask)
	}
}

const (
	RegisterTimcisrFieldDlyprtShift = 14
	RegisterTimcisrFieldDlyprtMask  = 0x4000
)

// GetDlyprt Delayed Protection Flag
func (r *registerTimcisrType) GetDlyprt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldDlyprtMask) != 0
}

// SetDlyprt Delayed Protection Flag
func (r *registerTimcisrType) SetDlyprt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldDlyprtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldDlyprtMask)
	}
}

const (
	RegisterTimcisrFieldRstShift = 13
	RegisterTimcisrFieldRstMask  = 0x2000
)

// GetRst Reset Interrupt Flag
func (r *registerTimcisrType) GetRst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldRstMask) != 0
}

// SetRst Reset Interrupt Flag
func (r *registerTimcisrType) SetRst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldRstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldRstMask)
	}
}

const (
	RegisterTimcisrFieldRstx2Shift = 12
	RegisterTimcisrFieldRstx2Mask  = 0x1000
)

// GetRstx2 Output 2 Reset Interrupt Flag
func (r *registerTimcisrType) GetRstx2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldRstx2Mask) != 0
}

// SetRstx2 Output 2 Reset Interrupt Flag
func (r *registerTimcisrType) SetRstx2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldRstx2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldRstx2Mask)
	}
}

const (
	RegisterTimcisrFieldSetx2Shift = 11
	RegisterTimcisrFieldSetx2Mask  = 0x800
)

// GetSetx2 Output 2 Set Interrupt Flag
func (r *registerTimcisrType) GetSetx2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldSetx2Mask) != 0
}

// SetSetx2 Output 2 Set Interrupt Flag
func (r *registerTimcisrType) SetSetx2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldSetx2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldSetx2Mask)
	}
}

const (
	RegisterTimcisrFieldRstx1Shift = 10
	RegisterTimcisrFieldRstx1Mask  = 0x400
)

// GetRstx1 Output 1 Reset Interrupt Flag
func (r *registerTimcisrType) GetRstx1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldRstx1Mask) != 0
}

// SetRstx1 Output 1 Reset Interrupt Flag
func (r *registerTimcisrType) SetRstx1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldRstx1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldRstx1Mask)
	}
}

const (
	RegisterTimcisrFieldSetx1Shift = 9
	RegisterTimcisrFieldSetx1Mask  = 0x200
)

// GetSetx1 Output 1 Set Interrupt Flag
func (r *registerTimcisrType) GetSetx1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldSetx1Mask) != 0
}

// SetSetx1 Output 1 Set Interrupt Flag
func (r *registerTimcisrType) SetSetx1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldSetx1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldSetx1Mask)
	}
}

const (
	RegisterTimcisrFieldCpt2Shift = 8
	RegisterTimcisrFieldCpt2Mask  = 0x100
)

// GetCpt2 Capture2 Interrupt Flag
func (r *registerTimcisrType) GetCpt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldCpt2Mask) != 0
}

// SetCpt2 Capture2 Interrupt Flag
func (r *registerTimcisrType) SetCpt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldCpt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldCpt2Mask)
	}
}

const (
	RegisterTimcisrFieldCpt1Shift = 7
	RegisterTimcisrFieldCpt1Mask  = 0x80
)

// GetCpt1 Capture1 Interrupt Flag
func (r *registerTimcisrType) GetCpt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldCpt1Mask) != 0
}

// SetCpt1 Capture1 Interrupt Flag
func (r *registerTimcisrType) SetCpt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldCpt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldCpt1Mask)
	}
}

const (
	RegisterTimcisrFieldUpdShift = 6
	RegisterTimcisrFieldUpdMask  = 0x40
)

// GetUpd Update Interrupt Flag
func (r *registerTimcisrType) GetUpd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldUpdMask) != 0
}

// SetUpd Update Interrupt Flag
func (r *registerTimcisrType) SetUpd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldUpdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldUpdMask)
	}
}

const (
	RegisterTimcisrFieldRepShift = 4
	RegisterTimcisrFieldRepMask  = 0x10
)

// GetRep Repetition Interrupt Flag
func (r *registerTimcisrType) GetRep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldRepMask) != 0
}

// SetRep Repetition Interrupt Flag
func (r *registerTimcisrType) SetRep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldRepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldRepMask)
	}
}

const (
	RegisterTimcisrFieldCmp4Shift = 3
	RegisterTimcisrFieldCmp4Mask  = 0x8
)

// GetCmp4 Compare 4 Interrupt Flag
func (r *registerTimcisrType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldCmp4Mask) != 0
}

// SetCmp4 Compare 4 Interrupt Flag
func (r *registerTimcisrType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldCmp4Mask)
	}
}

const (
	RegisterTimcisrFieldCmp3Shift = 2
	RegisterTimcisrFieldCmp3Mask  = 0x4
)

// GetCmp3 Compare 3 Interrupt Flag
func (r *registerTimcisrType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldCmp3Mask) != 0
}

// SetCmp3 Compare 3 Interrupt Flag
func (r *registerTimcisrType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldCmp3Mask)
	}
}

const (
	RegisterTimcisrFieldCmp2Shift = 1
	RegisterTimcisrFieldCmp2Mask  = 0x2
)

// GetCmp2 Compare 2 Interrupt Flag
func (r *registerTimcisrType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldCmp2Mask) != 0
}

// SetCmp2 Compare 2 Interrupt Flag
func (r *registerTimcisrType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldCmp2Mask)
	}
}

const (
	RegisterTimcisrFieldCmp1Shift = 0
	RegisterTimcisrFieldCmp1Mask  = 0x1
)

// GetCmp1 Compare 1 Interrupt Flag
func (r *registerTimcisrType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcisrFieldCmp1Mask) != 0
}

// SetCmp1 Compare 1 Interrupt Flag
func (r *registerTimcisrType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcisrFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcisrFieldCmp1Mask)
	}
}

// registerTimcicrType Timerx Interrupt Clear Register
type registerTimcicrType uint32

const (
	RegisterTimcicrFieldDlyprtcShift = 14
	RegisterTimcicrFieldDlyprtcMask  = 0x4000
)

// GetDlyprtc Delayed Protection Flag Clear
func (r *registerTimcicrType) GetDlyprtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldDlyprtcMask) != 0
}

// SetDlyprtc Delayed Protection Flag Clear
func (r *registerTimcicrType) SetDlyprtc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldDlyprtcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldDlyprtcMask)
	}
}

const (
	RegisterTimcicrFieldRstcShift = 13
	RegisterTimcicrFieldRstcMask  = 0x2000
)

// GetRstc Reset Interrupt flag Clear
func (r *registerTimcicrType) GetRstc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldRstcMask) != 0
}

// SetRstc Reset Interrupt flag Clear
func (r *registerTimcicrType) SetRstc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldRstcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldRstcMask)
	}
}

const (
	RegisterTimcicrFieldRstx2cShift = 12
	RegisterTimcicrFieldRstx2cMask  = 0x1000
)

// GetRstx2c Output 2 Reset flag Clear
func (r *registerTimcicrType) GetRstx2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldRstx2cMask) != 0
}

// SetRstx2c Output 2 Reset flag Clear
func (r *registerTimcicrType) SetRstx2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldRstx2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldRstx2cMask)
	}
}

const (
	RegisterTimcicrFieldSet2xcShift = 11
	RegisterTimcicrFieldSet2xcMask  = 0x800
)

// GetSet2xc Output 2 Set flag Clear
func (r *registerTimcicrType) GetSet2xc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldSet2xcMask) != 0
}

// SetSet2xc Output 2 Set flag Clear
func (r *registerTimcicrType) SetSet2xc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldSet2xcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldSet2xcMask)
	}
}

const (
	RegisterTimcicrFieldRstx1cShift = 10
	RegisterTimcicrFieldRstx1cMask  = 0x400
)

// GetRstx1c Output 1 Reset flag Clear
func (r *registerTimcicrType) GetRstx1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldRstx1cMask) != 0
}

// SetRstx1c Output 1 Reset flag Clear
func (r *registerTimcicrType) SetRstx1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldRstx1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldRstx1cMask)
	}
}

const (
	RegisterTimcicrFieldSet1xcShift = 9
	RegisterTimcicrFieldSet1xcMask  = 0x200
)

// GetSet1xc Output 1 Set flag Clear
func (r *registerTimcicrType) GetSet1xc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldSet1xcMask) != 0
}

// SetSet1xc Output 1 Set flag Clear
func (r *registerTimcicrType) SetSet1xc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldSet1xcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldSet1xcMask)
	}
}

const (
	RegisterTimcicrFieldCpt2cShift = 8
	RegisterTimcicrFieldCpt2cMask  = 0x100
)

// GetCpt2c Capture2 Interrupt flag Clear
func (r *registerTimcicrType) GetCpt2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldCpt2cMask) != 0
}

// SetCpt2c Capture2 Interrupt flag Clear
func (r *registerTimcicrType) SetCpt2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldCpt2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldCpt2cMask)
	}
}

const (
	RegisterTimcicrFieldCpt1cShift = 7
	RegisterTimcicrFieldCpt1cMask  = 0x80
)

// GetCpt1c Capture1 Interrupt flag Clear
func (r *registerTimcicrType) GetCpt1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldCpt1cMask) != 0
}

// SetCpt1c Capture1 Interrupt flag Clear
func (r *registerTimcicrType) SetCpt1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldCpt1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldCpt1cMask)
	}
}

const (
	RegisterTimcicrFieldUpdcShift = 6
	RegisterTimcicrFieldUpdcMask  = 0x40
)

// GetUpdc Update Interrupt flag Clear
func (r *registerTimcicrType) GetUpdc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldUpdcMask) != 0
}

// SetUpdc Update Interrupt flag Clear
func (r *registerTimcicrType) SetUpdc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldUpdcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldUpdcMask)
	}
}

const (
	RegisterTimcicrFieldRepcShift = 4
	RegisterTimcicrFieldRepcMask  = 0x10
)

// GetRepc Repetition Interrupt flag Clear
func (r *registerTimcicrType) GetRepc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldRepcMask) != 0
}

// SetRepc Repetition Interrupt flag Clear
func (r *registerTimcicrType) SetRepc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldRepcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldRepcMask)
	}
}

const (
	RegisterTimcicrFieldCmp4cShift = 3
	RegisterTimcicrFieldCmp4cMask  = 0x8
)

// GetCmp4c Compare 4 Interrupt flag Clear
func (r *registerTimcicrType) GetCmp4c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldCmp4cMask) != 0
}

// SetCmp4c Compare 4 Interrupt flag Clear
func (r *registerTimcicrType) SetCmp4c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldCmp4cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldCmp4cMask)
	}
}

const (
	RegisterTimcicrFieldCmp3cShift = 2
	RegisterTimcicrFieldCmp3cMask  = 0x4
)

// GetCmp3c Compare 3 Interrupt flag Clear
func (r *registerTimcicrType) GetCmp3c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldCmp3cMask) != 0
}

// SetCmp3c Compare 3 Interrupt flag Clear
func (r *registerTimcicrType) SetCmp3c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldCmp3cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldCmp3cMask)
	}
}

const (
	RegisterTimcicrFieldCmp2cShift = 1
	RegisterTimcicrFieldCmp2cMask  = 0x2
)

// GetCmp2c Compare 2 Interrupt flag Clear
func (r *registerTimcicrType) GetCmp2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldCmp2cMask) != 0
}

// SetCmp2c Compare 2 Interrupt flag Clear
func (r *registerTimcicrType) SetCmp2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldCmp2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldCmp2cMask)
	}
}

const (
	RegisterTimcicrFieldCmp1cShift = 0
	RegisterTimcicrFieldCmp1cMask  = 0x1
)

// GetCmp1c Compare 1 Interrupt flag Clear
func (r *registerTimcicrType) GetCmp1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcicrFieldCmp1cMask) != 0
}

// SetCmp1c Compare 1 Interrupt flag Clear
func (r *registerTimcicrType) SetCmp1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcicrFieldCmp1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcicrFieldCmp1cMask)
	}
}

// registerTimcdier5Type TIMxDIER5
type registerTimcdier5Type uint32

const (
	RegisterTimcdier5FieldDlyprtdeShift = 30
	RegisterTimcdier5FieldDlyprtdeMask  = 0x40000000
)

// GetDlyprtde DLYPRTDE
func (r *registerTimcdier5Type) GetDlyprtde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldDlyprtdeMask) != 0
}

// SetDlyprtde DLYPRTDE
func (r *registerTimcdier5Type) SetDlyprtde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldDlyprtdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldDlyprtdeMask)
	}
}

const (
	RegisterTimcdier5FieldRstdeShift = 29
	RegisterTimcdier5FieldRstdeMask  = 0x20000000
)

// GetRstde RSTDE
func (r *registerTimcdier5Type) GetRstde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldRstdeMask) != 0
}

// SetRstde RSTDE
func (r *registerTimcdier5Type) SetRstde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldRstdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldRstdeMask)
	}
}

const (
	RegisterTimcdier5FieldRstx2deShift = 28
	RegisterTimcdier5FieldRstx2deMask  = 0x10000000
)

// GetRstx2de RSTx2DE
func (r *registerTimcdier5Type) GetRstx2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldRstx2deMask) != 0
}

// SetRstx2de RSTx2DE
func (r *registerTimcdier5Type) SetRstx2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldRstx2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldRstx2deMask)
	}
}

const (
	RegisterTimcdier5FieldSetx2deShift = 27
	RegisterTimcdier5FieldSetx2deMask  = 0x8000000
)

// GetSetx2de SETx2DE
func (r *registerTimcdier5Type) GetSetx2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldSetx2deMask) != 0
}

// SetSetx2de SETx2DE
func (r *registerTimcdier5Type) SetSetx2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldSetx2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldSetx2deMask)
	}
}

const (
	RegisterTimcdier5FieldRstx1deShift = 26
	RegisterTimcdier5FieldRstx1deMask  = 0x4000000
)

// GetRstx1de RSTx1DE
func (r *registerTimcdier5Type) GetRstx1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldRstx1deMask) != 0
}

// SetRstx1de RSTx1DE
func (r *registerTimcdier5Type) SetRstx1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldRstx1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldRstx1deMask)
	}
}

const (
	RegisterTimcdier5FieldSet1xdeShift = 25
	RegisterTimcdier5FieldSet1xdeMask  = 0x2000000
)

// GetSet1xde SET1xDE
func (r *registerTimcdier5Type) GetSet1xde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldSet1xdeMask) != 0
}

// SetSet1xde SET1xDE
func (r *registerTimcdier5Type) SetSet1xde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldSet1xdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldSet1xdeMask)
	}
}

const (
	RegisterTimcdier5FieldCpt2deShift = 24
	RegisterTimcdier5FieldCpt2deMask  = 0x1000000
)

// GetCpt2de CPT2DE
func (r *registerTimcdier5Type) GetCpt2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldCpt2deMask) != 0
}

// SetCpt2de CPT2DE
func (r *registerTimcdier5Type) SetCpt2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldCpt2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldCpt2deMask)
	}
}

const (
	RegisterTimcdier5FieldCpt1deShift = 23
	RegisterTimcdier5FieldCpt1deMask  = 0x800000
)

// GetCpt1de CPT1DE
func (r *registerTimcdier5Type) GetCpt1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldCpt1deMask) != 0
}

// SetCpt1de CPT1DE
func (r *registerTimcdier5Type) SetCpt1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldCpt1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldCpt1deMask)
	}
}

const (
	RegisterTimcdier5FieldUpddeShift = 22
	RegisterTimcdier5FieldUpddeMask  = 0x400000
)

// GetUpdde UPDDE
func (r *registerTimcdier5Type) GetUpdde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldUpddeMask) != 0
}

// SetUpdde UPDDE
func (r *registerTimcdier5Type) SetUpdde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldUpddeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldUpddeMask)
	}
}

const (
	RegisterTimcdier5FieldRepdeShift = 20
	RegisterTimcdier5FieldRepdeMask  = 0x100000
)

// GetRepde REPDE
func (r *registerTimcdier5Type) GetRepde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldRepdeMask) != 0
}

// SetRepde REPDE
func (r *registerTimcdier5Type) SetRepde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldRepdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldRepdeMask)
	}
}

const (
	RegisterTimcdier5FieldCmp4deShift = 19
	RegisterTimcdier5FieldCmp4deMask  = 0x80000
)

// GetCmp4de CMP4DE
func (r *registerTimcdier5Type) GetCmp4de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldCmp4deMask) != 0
}

// SetCmp4de CMP4DE
func (r *registerTimcdier5Type) SetCmp4de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldCmp4deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldCmp4deMask)
	}
}

const (
	RegisterTimcdier5FieldCmp3deShift = 18
	RegisterTimcdier5FieldCmp3deMask  = 0x40000
)

// GetCmp3de CMP3DE
func (r *registerTimcdier5Type) GetCmp3de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldCmp3deMask) != 0
}

// SetCmp3de CMP3DE
func (r *registerTimcdier5Type) SetCmp3de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldCmp3deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldCmp3deMask)
	}
}

const (
	RegisterTimcdier5FieldCmp2deShift = 17
	RegisterTimcdier5FieldCmp2deMask  = 0x20000
)

// GetCmp2de CMP2DE
func (r *registerTimcdier5Type) GetCmp2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldCmp2deMask) != 0
}

// SetCmp2de CMP2DE
func (r *registerTimcdier5Type) SetCmp2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldCmp2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldCmp2deMask)
	}
}

const (
	RegisterTimcdier5FieldCmp1deShift = 16
	RegisterTimcdier5FieldCmp1deMask  = 0x10000
)

// GetCmp1de CMP1DE
func (r *registerTimcdier5Type) GetCmp1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldCmp1deMask) != 0
}

// SetCmp1de CMP1DE
func (r *registerTimcdier5Type) SetCmp1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldCmp1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldCmp1deMask)
	}
}

const (
	RegisterTimcdier5FieldDlyprtieShift = 14
	RegisterTimcdier5FieldDlyprtieMask  = 0x4000
)

// GetDlyprtie DLYPRTIE
func (r *registerTimcdier5Type) GetDlyprtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldDlyprtieMask) != 0
}

// SetDlyprtie DLYPRTIE
func (r *registerTimcdier5Type) SetDlyprtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldDlyprtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldDlyprtieMask)
	}
}

const (
	RegisterTimcdier5FieldRstieShift = 13
	RegisterTimcdier5FieldRstieMask  = 0x2000
)

// GetRstie RSTIE
func (r *registerTimcdier5Type) GetRstie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldRstieMask) != 0
}

// SetRstie RSTIE
func (r *registerTimcdier5Type) SetRstie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldRstieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldRstieMask)
	}
}

const (
	RegisterTimcdier5FieldRstx2ieShift = 12
	RegisterTimcdier5FieldRstx2ieMask  = 0x1000
)

// GetRstx2ie RSTx2IE
func (r *registerTimcdier5Type) GetRstx2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldRstx2ieMask) != 0
}

// SetRstx2ie RSTx2IE
func (r *registerTimcdier5Type) SetRstx2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldRstx2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldRstx2ieMask)
	}
}

const (
	RegisterTimcdier5FieldSetx2ieShift = 11
	RegisterTimcdier5FieldSetx2ieMask  = 0x800
)

// GetSetx2ie SETx2IE
func (r *registerTimcdier5Type) GetSetx2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldSetx2ieMask) != 0
}

// SetSetx2ie SETx2IE
func (r *registerTimcdier5Type) SetSetx2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldSetx2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldSetx2ieMask)
	}
}

const (
	RegisterTimcdier5FieldRstx1ieShift = 10
	RegisterTimcdier5FieldRstx1ieMask  = 0x400
)

// GetRstx1ie RSTx1IE
func (r *registerTimcdier5Type) GetRstx1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldRstx1ieMask) != 0
}

// SetRstx1ie RSTx1IE
func (r *registerTimcdier5Type) SetRstx1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldRstx1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldRstx1ieMask)
	}
}

const (
	RegisterTimcdier5FieldSet1xieShift = 9
	RegisterTimcdier5FieldSet1xieMask  = 0x200
)

// GetSet1xie SET1xIE
func (r *registerTimcdier5Type) GetSet1xie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldSet1xieMask) != 0
}

// SetSet1xie SET1xIE
func (r *registerTimcdier5Type) SetSet1xie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldSet1xieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldSet1xieMask)
	}
}

const (
	RegisterTimcdier5FieldCpt2ieShift = 8
	RegisterTimcdier5FieldCpt2ieMask  = 0x100
)

// GetCpt2ie CPT2IE
func (r *registerTimcdier5Type) GetCpt2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldCpt2ieMask) != 0
}

// SetCpt2ie CPT2IE
func (r *registerTimcdier5Type) SetCpt2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldCpt2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldCpt2ieMask)
	}
}

const (
	RegisterTimcdier5FieldCpt1ieShift = 7
	RegisterTimcdier5FieldCpt1ieMask  = 0x80
)

// GetCpt1ie CPT1IE
func (r *registerTimcdier5Type) GetCpt1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldCpt1ieMask) != 0
}

// SetCpt1ie CPT1IE
func (r *registerTimcdier5Type) SetCpt1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldCpt1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldCpt1ieMask)
	}
}

const (
	RegisterTimcdier5FieldUpdieShift = 6
	RegisterTimcdier5FieldUpdieMask  = 0x40
)

// GetUpdie UPDIE
func (r *registerTimcdier5Type) GetUpdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldUpdieMask) != 0
}

// SetUpdie UPDIE
func (r *registerTimcdier5Type) SetUpdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldUpdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldUpdieMask)
	}
}

const (
	RegisterTimcdier5FieldRepieShift = 4
	RegisterTimcdier5FieldRepieMask  = 0x10
)

// GetRepie REPIE
func (r *registerTimcdier5Type) GetRepie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldRepieMask) != 0
}

// SetRepie REPIE
func (r *registerTimcdier5Type) SetRepie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldRepieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldRepieMask)
	}
}

const (
	RegisterTimcdier5FieldCmp4ieShift = 3
	RegisterTimcdier5FieldCmp4ieMask  = 0x8
)

// GetCmp4ie CMP4IE
func (r *registerTimcdier5Type) GetCmp4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldCmp4ieMask) != 0
}

// SetCmp4ie CMP4IE
func (r *registerTimcdier5Type) SetCmp4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldCmp4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldCmp4ieMask)
	}
}

const (
	RegisterTimcdier5FieldCmp3ieShift = 2
	RegisterTimcdier5FieldCmp3ieMask  = 0x4
)

// GetCmp3ie CMP3IE
func (r *registerTimcdier5Type) GetCmp3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldCmp3ieMask) != 0
}

// SetCmp3ie CMP3IE
func (r *registerTimcdier5Type) SetCmp3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldCmp3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldCmp3ieMask)
	}
}

const (
	RegisterTimcdier5FieldCmp2ieShift = 1
	RegisterTimcdier5FieldCmp2ieMask  = 0x2
)

// GetCmp2ie CMP2IE
func (r *registerTimcdier5Type) GetCmp2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldCmp2ieMask) != 0
}

// SetCmp2ie CMP2IE
func (r *registerTimcdier5Type) SetCmp2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldCmp2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldCmp2ieMask)
	}
}

const (
	RegisterTimcdier5FieldCmp1ieShift = 0
	RegisterTimcdier5FieldCmp1ieMask  = 0x1
)

// GetCmp1ie CMP1IE
func (r *registerTimcdier5Type) GetCmp1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimcdier5FieldCmp1ieMask) != 0
}

// SetCmp1ie CMP1IE
func (r *registerTimcdier5Type) SetCmp1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimcdier5FieldCmp1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimcdier5FieldCmp1ieMask)
	}
}

// registerCntcrType Timerx Counter Register
type registerCntcrType uint32

const (
	RegisterCntcrFieldCntxShift = 0
	RegisterCntcrFieldCntxMask  = 0xffff
)

// GetCntx Timerx Counter value
func (r *registerCntcrType) GetCntx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntcrFieldCntxMask) >> RegisterCntcrFieldCntxShift)
}

// SetCntx Timerx Counter value
func (r *registerCntcrType) SetCntx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntcrFieldCntxMask)|(uint32(value)<<RegisterCntcrFieldCntxShift))
}

// registerPercrType Timerx Period Register
type registerPercrType uint32

const (
	RegisterPercrFieldPerxShift = 0
	RegisterPercrFieldPerxMask  = 0xffff
)

// GetPerx Timerx Period value
func (r *registerPercrType) GetPerx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPercrFieldPerxMask) >> RegisterPercrFieldPerxShift)
}

// SetPerx Timerx Period value
func (r *registerPercrType) SetPerx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPercrFieldPerxMask)|(uint32(value)<<RegisterPercrFieldPerxShift))
}

// registerRepcrType Timerx Repetition Register
type registerRepcrType uint32

const (
	RegisterRepcrFieldRepxShift = 0
	RegisterRepcrFieldRepxMask  = 0xff
)

// GetRepx Timerx Repetition counter value
func (r *registerRepcrType) GetRepx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRepcrFieldRepxMask) >> RegisterRepcrFieldRepxShift)
}

// SetRepx Timerx Repetition counter value
func (r *registerRepcrType) SetRepx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRepcrFieldRepxMask)|(uint32(value)<<RegisterRepcrFieldRepxShift))
}

// registerCmp1crType Timerx Compare 1 Register
type registerCmp1crType uint32

const (
	RegisterCmp1crFieldCmp1xShift = 0
	RegisterCmp1crFieldCmp1xMask  = 0xffff
)

// GetCmp1x Timerx Compare 1 value
func (r *registerCmp1crType) GetCmp1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1crFieldCmp1xMask) >> RegisterCmp1crFieldCmp1xShift)
}

// SetCmp1x Timerx Compare 1 value
func (r *registerCmp1crType) SetCmp1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1crFieldCmp1xMask)|(uint32(value)<<RegisterCmp1crFieldCmp1xShift))
}

// registerCmp1ccrType Timerx Compare 1 Compound Register
type registerCmp1ccrType uint32

const (
	RegisterCmp1ccrFieldRepxShift = 16
	RegisterCmp1ccrFieldRepxMask  = 0xff0000
)

// GetRepx Timerx Repetition value (aliased from HRTIM_REPx register)
func (r *registerCmp1ccrType) GetRepx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1ccrFieldRepxMask) >> RegisterCmp1ccrFieldRepxShift)
}

// SetRepx Timerx Repetition value (aliased from HRTIM_REPx register)
func (r *registerCmp1ccrType) SetRepx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1ccrFieldRepxMask)|(uint32(value)<<RegisterCmp1ccrFieldRepxShift))
}

const (
	RegisterCmp1ccrFieldCmp1xShift = 0
	RegisterCmp1ccrFieldCmp1xMask  = 0xffff
)

// GetCmp1x Timerx Compare 1 value
func (r *registerCmp1ccrType) GetCmp1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1ccrFieldCmp1xMask) >> RegisterCmp1ccrFieldCmp1xShift)
}

// SetCmp1x Timerx Compare 1 value
func (r *registerCmp1ccrType) SetCmp1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1ccrFieldCmp1xMask)|(uint32(value)<<RegisterCmp1ccrFieldCmp1xShift))
}

// registerCmp2crType Timerx Compare 2 Register
type registerCmp2crType uint32

const (
	RegisterCmp2crFieldCmp2xShift = 0
	RegisterCmp2crFieldCmp2xMask  = 0xffff
)

// GetCmp2x Timerx Compare 2 value
func (r *registerCmp2crType) GetCmp2x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp2crFieldCmp2xMask) >> RegisterCmp2crFieldCmp2xShift)
}

// SetCmp2x Timerx Compare 2 value
func (r *registerCmp2crType) SetCmp2x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp2crFieldCmp2xMask)|(uint32(value)<<RegisterCmp2crFieldCmp2xShift))
}

// registerCmp3crType Timerx Compare 3 Register
type registerCmp3crType uint32

const (
	RegisterCmp3crFieldCmp3xShift = 0
	RegisterCmp3crFieldCmp3xMask  = 0xffff
)

// GetCmp3x Timerx Compare 3 value
func (r *registerCmp3crType) GetCmp3x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp3crFieldCmp3xMask) >> RegisterCmp3crFieldCmp3xShift)
}

// SetCmp3x Timerx Compare 3 value
func (r *registerCmp3crType) SetCmp3x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp3crFieldCmp3xMask)|(uint32(value)<<RegisterCmp3crFieldCmp3xShift))
}

// registerCmp4crType Timerx Compare 4 Register
type registerCmp4crType uint32

const (
	RegisterCmp4crFieldCmp4xShift = 0
	RegisterCmp4crFieldCmp4xMask  = 0xffff
)

// GetCmp4x Timerx Compare 4 value
func (r *registerCmp4crType) GetCmp4x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp4crFieldCmp4xMask) >> RegisterCmp4crFieldCmp4xShift)
}

// SetCmp4x Timerx Compare 4 value
func (r *registerCmp4crType) SetCmp4x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp4crFieldCmp4xMask)|(uint32(value)<<RegisterCmp4crFieldCmp4xShift))
}

// registerCpt1crType Timerx Capture 1 Register
type registerCpt1crType uint32

const (
	RegisterCpt1crFieldCpt1xShift = 0
	RegisterCpt1crFieldCpt1xMask  = 0xffff
)

// GetCpt1x Timerx Capture 1 value
func (r *registerCpt1crType) GetCpt1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpt1crFieldCpt1xMask) >> RegisterCpt1crFieldCpt1xShift)
}

// SetCpt1x Timerx Capture 1 value
func (r *registerCpt1crType) SetCpt1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpt1crFieldCpt1xMask)|(uint32(value)<<RegisterCpt1crFieldCpt1xShift))
}

// registerCpt2crType Timerx Capture 2 Register
type registerCpt2crType uint32

const (
	RegisterCpt2crFieldCpt2xShift = 0
	RegisterCpt2crFieldCpt2xMask  = 0xffff
)

// GetCpt2x Timerx Capture 2 value
func (r *registerCpt2crType) GetCpt2x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpt2crFieldCpt2xMask) >> RegisterCpt2crFieldCpt2xShift)
}

// SetCpt2x Timerx Capture 2 value
func (r *registerCpt2crType) SetCpt2x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpt2crFieldCpt2xMask)|(uint32(value)<<RegisterCpt2crFieldCpt2xShift))
}

// registerDtcrType Timerx Deadtime Register
type registerDtcrType uint32

const (
	RegisterDtcrFieldDtflkxShift = 31
	RegisterDtcrFieldDtflkxMask  = 0x80000000
)

// GetDtflkx Deadtime Falling Lock
func (r *registerDtcrType) GetDtflkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtcrFieldDtflkxMask) != 0
}

// SetDtflkx Deadtime Falling Lock
func (r *registerDtcrType) SetDtflkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtcrFieldDtflkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtcrFieldDtflkxMask)
	}
}

const (
	RegisterDtcrFieldDtfslkxShift = 30
	RegisterDtcrFieldDtfslkxMask  = 0x40000000
)

// GetDtfslkx Deadtime Falling Sign Lock
func (r *registerDtcrType) GetDtfslkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtcrFieldDtfslkxMask) != 0
}

// SetDtfslkx Deadtime Falling Sign Lock
func (r *registerDtcrType) SetDtfslkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtcrFieldDtfslkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtcrFieldDtfslkxMask)
	}
}

const (
	RegisterDtcrFieldSdtfxShift = 25
	RegisterDtcrFieldSdtfxMask  = 0x2000000
)

// GetSdtfx Sign Deadtime Falling value
func (r *registerDtcrType) GetSdtfx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtcrFieldSdtfxMask) != 0
}

// SetSdtfx Sign Deadtime Falling value
func (r *registerDtcrType) SetSdtfx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtcrFieldSdtfxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtcrFieldSdtfxMask)
	}
}

const (
	RegisterDtcrFieldDtfxShift = 16
	RegisterDtcrFieldDtfxMask  = 0x1ff0000
)

// GetDtfx Deadtime Falling value
func (r *registerDtcrType) GetDtfx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDtcrFieldDtfxMask) >> RegisterDtcrFieldDtfxShift)
}

// SetDtfx Deadtime Falling value
func (r *registerDtcrType) SetDtfx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtcrFieldDtfxMask)|(uint32(value)<<RegisterDtcrFieldDtfxShift))
}

const (
	RegisterDtcrFieldDtrlkxShift = 15
	RegisterDtcrFieldDtrlkxMask  = 0x8000
)

// GetDtrlkx Deadtime Rising Lock
func (r *registerDtcrType) GetDtrlkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtcrFieldDtrlkxMask) != 0
}

// SetDtrlkx Deadtime Rising Lock
func (r *registerDtcrType) SetDtrlkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtcrFieldDtrlkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtcrFieldDtrlkxMask)
	}
}

const (
	RegisterDtcrFieldDtrslkxShift = 14
	RegisterDtcrFieldDtrslkxMask  = 0x4000
)

// GetDtrslkx Deadtime Rising Sign Lock
func (r *registerDtcrType) GetDtrslkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtcrFieldDtrslkxMask) != 0
}

// SetDtrslkx Deadtime Rising Sign Lock
func (r *registerDtcrType) SetDtrslkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtcrFieldDtrslkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtcrFieldDtrslkxMask)
	}
}

const (
	RegisterDtcrFieldDtprscShift = 10
	RegisterDtcrFieldDtprscMask  = 0x1c00
)

// GetDtprsc Deadtime Prescaler
func (r *registerDtcrType) GetDtprsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDtcrFieldDtprscMask) >> RegisterDtcrFieldDtprscShift)
}

// SetDtprsc Deadtime Prescaler
func (r *registerDtcrType) SetDtprsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtcrFieldDtprscMask)|(uint32(value)<<RegisterDtcrFieldDtprscShift))
}

const (
	RegisterDtcrFieldSdtrxShift = 9
	RegisterDtcrFieldSdtrxMask  = 0x200
)

// GetSdtrx Sign Deadtime Rising value
func (r *registerDtcrType) GetSdtrx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtcrFieldSdtrxMask) != 0
}

// SetSdtrx Sign Deadtime Rising value
func (r *registerDtcrType) SetSdtrx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtcrFieldSdtrxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtcrFieldSdtrxMask)
	}
}

const (
	RegisterDtcrFieldDtrxShift = 0
	RegisterDtcrFieldDtrxMask  = 0x1ff
)

// GetDtrx Deadtime Rising value
func (r *registerDtcrType) GetDtrx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDtcrFieldDtrxMask) >> RegisterDtcrFieldDtrxShift)
}

// SetDtrx Deadtime Rising value
func (r *registerDtcrType) SetDtrx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtcrFieldDtrxMask)|(uint32(value)<<RegisterDtcrFieldDtrxShift))
}

// registerSetc1rType Timerx Output1 Set Register
type registerSetc1rType uint32

const (
	RegisterSetc1rFieldUpdateShift = 31
	RegisterSetc1rFieldUpdateMask  = 0x80000000
)

// GetUpdate Registers update (transfer preload to active)
func (r *registerSetc1rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldUpdateMask) != 0
}

// SetUpdate Registers update (transfer preload to active)
func (r *registerSetc1rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldUpdateMask)
	}
}

const (
	RegisterSetc1rFieldExtevnt10Shift = 30
	RegisterSetc1rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 External Event 10
func (r *registerSetc1rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 External Event 10
func (r *registerSetc1rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldExtevnt10Mask)
	}
}

const (
	RegisterSetc1rFieldExtevnt9Shift = 29
	RegisterSetc1rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 External Event 9
func (r *registerSetc1rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 External Event 9
func (r *registerSetc1rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldExtevnt9Mask)
	}
}

const (
	RegisterSetc1rFieldExtevnt8Shift = 28
	RegisterSetc1rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 External Event 8
func (r *registerSetc1rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 External Event 8
func (r *registerSetc1rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldExtevnt8Mask)
	}
}

const (
	RegisterSetc1rFieldExtevnt7Shift = 27
	RegisterSetc1rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 External Event 7
func (r *registerSetc1rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 External Event 7
func (r *registerSetc1rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldExtevnt7Mask)
	}
}

const (
	RegisterSetc1rFieldExtevnt6Shift = 26
	RegisterSetc1rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 External Event 6
func (r *registerSetc1rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 External Event 6
func (r *registerSetc1rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldExtevnt6Mask)
	}
}

const (
	RegisterSetc1rFieldExtevnt5Shift = 25
	RegisterSetc1rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 External Event 5
func (r *registerSetc1rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 External Event 5
func (r *registerSetc1rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldExtevnt5Mask)
	}
}

const (
	RegisterSetc1rFieldExtevnt4Shift = 24
	RegisterSetc1rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 External Event 4
func (r *registerSetc1rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 External Event 4
func (r *registerSetc1rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldExtevnt4Mask)
	}
}

const (
	RegisterSetc1rFieldExtevnt3Shift = 23
	RegisterSetc1rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 External Event 3
func (r *registerSetc1rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 External Event 3
func (r *registerSetc1rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldExtevnt3Mask)
	}
}

const (
	RegisterSetc1rFieldExtevnt2Shift = 22
	RegisterSetc1rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 External Event 2
func (r *registerSetc1rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 External Event 2
func (r *registerSetc1rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldExtevnt2Mask)
	}
}

const (
	RegisterSetc1rFieldExtevnt1Shift = 21
	RegisterSetc1rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 External Event 1
func (r *registerSetc1rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 External Event 1
func (r *registerSetc1rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldExtevnt1Mask)
	}
}

const (
	RegisterSetc1rFieldTimevnt9Shift = 20
	RegisterSetc1rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 Timer Event 9
func (r *registerSetc1rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 Timer Event 9
func (r *registerSetc1rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldTimevnt9Mask)
	}
}

const (
	RegisterSetc1rFieldTimevnt8Shift = 19
	RegisterSetc1rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 Timer Event 8
func (r *registerSetc1rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 Timer Event 8
func (r *registerSetc1rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldTimevnt8Mask)
	}
}

const (
	RegisterSetc1rFieldTimevnt7Shift = 18
	RegisterSetc1rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 Timer Event 7
func (r *registerSetc1rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 Timer Event 7
func (r *registerSetc1rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldTimevnt7Mask)
	}
}

const (
	RegisterSetc1rFieldTimevnt6Shift = 17
	RegisterSetc1rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 Timer Event 6
func (r *registerSetc1rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 Timer Event 6
func (r *registerSetc1rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldTimevnt6Mask)
	}
}

const (
	RegisterSetc1rFieldTimevnt5Shift = 16
	RegisterSetc1rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 Timer Event 5
func (r *registerSetc1rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 Timer Event 5
func (r *registerSetc1rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldTimevnt5Mask)
	}
}

const (
	RegisterSetc1rFieldTimevnt4Shift = 15
	RegisterSetc1rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 Timer Event 4
func (r *registerSetc1rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 Timer Event 4
func (r *registerSetc1rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldTimevnt4Mask)
	}
}

const (
	RegisterSetc1rFieldTimevnt3Shift = 14
	RegisterSetc1rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 Timer Event 3
func (r *registerSetc1rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 Timer Event 3
func (r *registerSetc1rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldTimevnt3Mask)
	}
}

const (
	RegisterSetc1rFieldTimevnt2Shift = 13
	RegisterSetc1rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 Timer Event 2
func (r *registerSetc1rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 Timer Event 2
func (r *registerSetc1rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldTimevnt2Mask)
	}
}

const (
	RegisterSetc1rFieldTimevnt1Shift = 12
	RegisterSetc1rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 Timer Event 1
func (r *registerSetc1rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 Timer Event 1
func (r *registerSetc1rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldTimevnt1Mask)
	}
}

const (
	RegisterSetc1rFieldMstcmp4Shift = 11
	RegisterSetc1rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 Master Compare 4
func (r *registerSetc1rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 Master Compare 4
func (r *registerSetc1rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldMstcmp4Mask)
	}
}

const (
	RegisterSetc1rFieldMstcmp3Shift = 10
	RegisterSetc1rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 Master Compare 3
func (r *registerSetc1rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 Master Compare 3
func (r *registerSetc1rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldMstcmp3Mask)
	}
}

const (
	RegisterSetc1rFieldMstcmp2Shift = 9
	RegisterSetc1rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 Master Compare 2
func (r *registerSetc1rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 Master Compare 2
func (r *registerSetc1rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldMstcmp2Mask)
	}
}

const (
	RegisterSetc1rFieldMstcmp1Shift = 8
	RegisterSetc1rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 Master Compare 1
func (r *registerSetc1rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 Master Compare 1
func (r *registerSetc1rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldMstcmp1Mask)
	}
}

const (
	RegisterSetc1rFieldMstperShift = 7
	RegisterSetc1rFieldMstperMask  = 0x80
)

// GetMstper Master Period
func (r *registerSetc1rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldMstperMask) != 0
}

// SetMstper Master Period
func (r *registerSetc1rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldMstperMask)
	}
}

const (
	RegisterSetc1rFieldCmp4Shift = 6
	RegisterSetc1rFieldCmp4Mask  = 0x40
)

// GetCmp4 Timer A compare 4
func (r *registerSetc1rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldCmp4Mask) != 0
}

// SetCmp4 Timer A compare 4
func (r *registerSetc1rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldCmp4Mask)
	}
}

const (
	RegisterSetc1rFieldCmp3Shift = 5
	RegisterSetc1rFieldCmp3Mask  = 0x20
)

// GetCmp3 Timer A compare 3
func (r *registerSetc1rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldCmp3Mask) != 0
}

// SetCmp3 Timer A compare 3
func (r *registerSetc1rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldCmp3Mask)
	}
}

const (
	RegisterSetc1rFieldCmp2Shift = 4
	RegisterSetc1rFieldCmp2Mask  = 0x10
)

// GetCmp2 Timer A compare 2
func (r *registerSetc1rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldCmp2Mask) != 0
}

// SetCmp2 Timer A compare 2
func (r *registerSetc1rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldCmp2Mask)
	}
}

const (
	RegisterSetc1rFieldCmp1Shift = 3
	RegisterSetc1rFieldCmp1Mask  = 0x8
)

// GetCmp1 Timer A compare 1
func (r *registerSetc1rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldCmp1Mask) != 0
}

// SetCmp1 Timer A compare 1
func (r *registerSetc1rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldCmp1Mask)
	}
}

const (
	RegisterSetc1rFieldPerShift = 2
	RegisterSetc1rFieldPerMask  = 0x4
)

// GetPer Timer A Period
func (r *registerSetc1rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldPerMask) != 0
}

// SetPer Timer A Period
func (r *registerSetc1rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldPerMask)
	}
}

const (
	RegisterSetc1rFieldResyncShift = 1
	RegisterSetc1rFieldResyncMask  = 0x2
)

// GetResync Timer A resynchronizaton
func (r *registerSetc1rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldResyncMask) != 0
}

// SetResync Timer A resynchronizaton
func (r *registerSetc1rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldResyncMask)
	}
}

const (
	RegisterSetc1rFieldSstShift = 0
	RegisterSetc1rFieldSstMask  = 0x1
)

// GetSst Software Set trigger
func (r *registerSetc1rType) GetSst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc1rFieldSstMask) != 0
}

// SetSst Software Set trigger
func (r *registerSetc1rType) SetSst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc1rFieldSstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc1rFieldSstMask)
	}
}

// registerRstc1rType Timerx Output1 Reset Register
type registerRstc1rType uint32

const (
	RegisterRstc1rFieldUpdateShift = 31
	RegisterRstc1rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerRstc1rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerRstc1rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldUpdateMask)
	}
}

const (
	RegisterRstc1rFieldExtevnt10Shift = 30
	RegisterRstc1rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerRstc1rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerRstc1rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldExtevnt10Mask)
	}
}

const (
	RegisterRstc1rFieldExtevnt9Shift = 29
	RegisterRstc1rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerRstc1rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerRstc1rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldExtevnt9Mask)
	}
}

const (
	RegisterRstc1rFieldExtevnt8Shift = 28
	RegisterRstc1rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerRstc1rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerRstc1rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldExtevnt8Mask)
	}
}

const (
	RegisterRstc1rFieldExtevnt7Shift = 27
	RegisterRstc1rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerRstc1rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerRstc1rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldExtevnt7Mask)
	}
}

const (
	RegisterRstc1rFieldExtevnt6Shift = 26
	RegisterRstc1rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerRstc1rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerRstc1rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldExtevnt6Mask)
	}
}

const (
	RegisterRstc1rFieldExtevnt5Shift = 25
	RegisterRstc1rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerRstc1rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerRstc1rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldExtevnt5Mask)
	}
}

const (
	RegisterRstc1rFieldExtevnt4Shift = 24
	RegisterRstc1rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerRstc1rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerRstc1rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldExtevnt4Mask)
	}
}

const (
	RegisterRstc1rFieldExtevnt3Shift = 23
	RegisterRstc1rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerRstc1rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerRstc1rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldExtevnt3Mask)
	}
}

const (
	RegisterRstc1rFieldExtevnt2Shift = 22
	RegisterRstc1rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerRstc1rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerRstc1rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldExtevnt2Mask)
	}
}

const (
	RegisterRstc1rFieldExtevnt1Shift = 21
	RegisterRstc1rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerRstc1rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerRstc1rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldExtevnt1Mask)
	}
}

const (
	RegisterRstc1rFieldTimevnt9Shift = 20
	RegisterRstc1rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerRstc1rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerRstc1rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldTimevnt9Mask)
	}
}

const (
	RegisterRstc1rFieldTimevnt8Shift = 19
	RegisterRstc1rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerRstc1rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerRstc1rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldTimevnt8Mask)
	}
}

const (
	RegisterRstc1rFieldTimevnt7Shift = 18
	RegisterRstc1rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerRstc1rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerRstc1rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldTimevnt7Mask)
	}
}

const (
	RegisterRstc1rFieldTimevnt6Shift = 17
	RegisterRstc1rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerRstc1rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerRstc1rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldTimevnt6Mask)
	}
}

const (
	RegisterRstc1rFieldTimevnt5Shift = 16
	RegisterRstc1rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerRstc1rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerRstc1rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldTimevnt5Mask)
	}
}

const (
	RegisterRstc1rFieldTimevnt4Shift = 15
	RegisterRstc1rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerRstc1rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerRstc1rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldTimevnt4Mask)
	}
}

const (
	RegisterRstc1rFieldTimevnt3Shift = 14
	RegisterRstc1rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerRstc1rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerRstc1rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldTimevnt3Mask)
	}
}

const (
	RegisterRstc1rFieldTimevnt2Shift = 13
	RegisterRstc1rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerRstc1rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerRstc1rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldTimevnt2Mask)
	}
}

const (
	RegisterRstc1rFieldTimevnt1Shift = 12
	RegisterRstc1rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerRstc1rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerRstc1rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldTimevnt1Mask)
	}
}

const (
	RegisterRstc1rFieldMstcmp4Shift = 11
	RegisterRstc1rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerRstc1rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerRstc1rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldMstcmp4Mask)
	}
}

const (
	RegisterRstc1rFieldMstcmp3Shift = 10
	RegisterRstc1rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerRstc1rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerRstc1rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldMstcmp3Mask)
	}
}

const (
	RegisterRstc1rFieldMstcmp2Shift = 9
	RegisterRstc1rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerRstc1rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerRstc1rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldMstcmp2Mask)
	}
}

const (
	RegisterRstc1rFieldMstcmp1Shift = 8
	RegisterRstc1rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerRstc1rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerRstc1rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldMstcmp1Mask)
	}
}

const (
	RegisterRstc1rFieldMstperShift = 7
	RegisterRstc1rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerRstc1rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerRstc1rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldMstperMask)
	}
}

const (
	RegisterRstc1rFieldCmp4Shift = 6
	RegisterRstc1rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerRstc1rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerRstc1rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldCmp4Mask)
	}
}

const (
	RegisterRstc1rFieldCmp3Shift = 5
	RegisterRstc1rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerRstc1rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerRstc1rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldCmp3Mask)
	}
}

const (
	RegisterRstc1rFieldCmp2Shift = 4
	RegisterRstc1rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerRstc1rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerRstc1rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldCmp2Mask)
	}
}

const (
	RegisterRstc1rFieldCmp1Shift = 3
	RegisterRstc1rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerRstc1rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerRstc1rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldCmp1Mask)
	}
}

const (
	RegisterRstc1rFieldPerShift = 2
	RegisterRstc1rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerRstc1rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldPerMask) != 0
}

// SetPer PER
func (r *registerRstc1rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldPerMask)
	}
}

const (
	RegisterRstc1rFieldResyncShift = 1
	RegisterRstc1rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerRstc1rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerRstc1rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldResyncMask)
	}
}

const (
	RegisterRstc1rFieldSrtShift = 0
	RegisterRstc1rFieldSrtMask  = 0x1
)

// GetSrt SRT
func (r *registerRstc1rType) GetSrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc1rFieldSrtMask) != 0
}

// SetSrt SRT
func (r *registerRstc1rType) SetSrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc1rFieldSrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc1rFieldSrtMask)
	}
}

// registerSetc2rType Timerx Output2 Set Register
type registerSetc2rType uint32

const (
	RegisterSetc2rFieldUpdateShift = 31
	RegisterSetc2rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerSetc2rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerSetc2rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldUpdateMask)
	}
}

const (
	RegisterSetc2rFieldExtevnt10Shift = 30
	RegisterSetc2rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerSetc2rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerSetc2rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldExtevnt10Mask)
	}
}

const (
	RegisterSetc2rFieldExtevnt9Shift = 29
	RegisterSetc2rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerSetc2rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerSetc2rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldExtevnt9Mask)
	}
}

const (
	RegisterSetc2rFieldExtevnt8Shift = 28
	RegisterSetc2rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerSetc2rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerSetc2rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldExtevnt8Mask)
	}
}

const (
	RegisterSetc2rFieldExtevnt7Shift = 27
	RegisterSetc2rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerSetc2rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerSetc2rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldExtevnt7Mask)
	}
}

const (
	RegisterSetc2rFieldExtevnt6Shift = 26
	RegisterSetc2rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerSetc2rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerSetc2rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldExtevnt6Mask)
	}
}

const (
	RegisterSetc2rFieldExtevnt5Shift = 25
	RegisterSetc2rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerSetc2rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerSetc2rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldExtevnt5Mask)
	}
}

const (
	RegisterSetc2rFieldExtevnt4Shift = 24
	RegisterSetc2rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerSetc2rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerSetc2rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldExtevnt4Mask)
	}
}

const (
	RegisterSetc2rFieldExtevnt3Shift = 23
	RegisterSetc2rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerSetc2rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerSetc2rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldExtevnt3Mask)
	}
}

const (
	RegisterSetc2rFieldExtevnt2Shift = 22
	RegisterSetc2rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerSetc2rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerSetc2rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldExtevnt2Mask)
	}
}

const (
	RegisterSetc2rFieldExtevnt1Shift = 21
	RegisterSetc2rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerSetc2rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerSetc2rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldExtevnt1Mask)
	}
}

const (
	RegisterSetc2rFieldTimevnt9Shift = 20
	RegisterSetc2rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerSetc2rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerSetc2rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldTimevnt9Mask)
	}
}

const (
	RegisterSetc2rFieldTimevnt8Shift = 19
	RegisterSetc2rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerSetc2rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerSetc2rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldTimevnt8Mask)
	}
}

const (
	RegisterSetc2rFieldTimevnt7Shift = 18
	RegisterSetc2rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerSetc2rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerSetc2rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldTimevnt7Mask)
	}
}

const (
	RegisterSetc2rFieldTimevnt6Shift = 17
	RegisterSetc2rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerSetc2rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerSetc2rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldTimevnt6Mask)
	}
}

const (
	RegisterSetc2rFieldTimevnt5Shift = 16
	RegisterSetc2rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerSetc2rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerSetc2rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldTimevnt5Mask)
	}
}

const (
	RegisterSetc2rFieldTimevnt4Shift = 15
	RegisterSetc2rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerSetc2rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerSetc2rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldTimevnt4Mask)
	}
}

const (
	RegisterSetc2rFieldTimevnt3Shift = 14
	RegisterSetc2rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerSetc2rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerSetc2rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldTimevnt3Mask)
	}
}

const (
	RegisterSetc2rFieldTimevnt2Shift = 13
	RegisterSetc2rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerSetc2rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerSetc2rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldTimevnt2Mask)
	}
}

const (
	RegisterSetc2rFieldTimevnt1Shift = 12
	RegisterSetc2rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerSetc2rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerSetc2rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldTimevnt1Mask)
	}
}

const (
	RegisterSetc2rFieldMstcmp4Shift = 11
	RegisterSetc2rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerSetc2rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerSetc2rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldMstcmp4Mask)
	}
}

const (
	RegisterSetc2rFieldMstcmp3Shift = 10
	RegisterSetc2rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerSetc2rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerSetc2rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldMstcmp3Mask)
	}
}

const (
	RegisterSetc2rFieldMstcmp2Shift = 9
	RegisterSetc2rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerSetc2rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerSetc2rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldMstcmp2Mask)
	}
}

const (
	RegisterSetc2rFieldMstcmp1Shift = 8
	RegisterSetc2rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerSetc2rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerSetc2rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldMstcmp1Mask)
	}
}

const (
	RegisterSetc2rFieldMstperShift = 7
	RegisterSetc2rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerSetc2rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerSetc2rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldMstperMask)
	}
}

const (
	RegisterSetc2rFieldCmp4Shift = 6
	RegisterSetc2rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerSetc2rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerSetc2rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldCmp4Mask)
	}
}

const (
	RegisterSetc2rFieldCmp3Shift = 5
	RegisterSetc2rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerSetc2rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerSetc2rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldCmp3Mask)
	}
}

const (
	RegisterSetc2rFieldCmp2Shift = 4
	RegisterSetc2rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerSetc2rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerSetc2rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldCmp2Mask)
	}
}

const (
	RegisterSetc2rFieldCmp1Shift = 3
	RegisterSetc2rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerSetc2rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerSetc2rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldCmp1Mask)
	}
}

const (
	RegisterSetc2rFieldPerShift = 2
	RegisterSetc2rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerSetc2rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldPerMask) != 0
}

// SetPer PER
func (r *registerSetc2rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldPerMask)
	}
}

const (
	RegisterSetc2rFieldResyncShift = 1
	RegisterSetc2rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerSetc2rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerSetc2rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldResyncMask)
	}
}

const (
	RegisterSetc2rFieldSstShift = 0
	RegisterSetc2rFieldSstMask  = 0x1
)

// GetSst SST
func (r *registerSetc2rType) GetSst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSetc2rFieldSstMask) != 0
}

// SetSst SST
func (r *registerSetc2rType) SetSst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSetc2rFieldSstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSetc2rFieldSstMask)
	}
}

// registerRstc2rType Timerx Output2 Reset Register
type registerRstc2rType uint32

const (
	RegisterRstc2rFieldUpdateShift = 31
	RegisterRstc2rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerRstc2rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerRstc2rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldUpdateMask)
	}
}

const (
	RegisterRstc2rFieldExtevnt10Shift = 30
	RegisterRstc2rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerRstc2rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerRstc2rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldExtevnt10Mask)
	}
}

const (
	RegisterRstc2rFieldExtevnt9Shift = 29
	RegisterRstc2rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerRstc2rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerRstc2rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldExtevnt9Mask)
	}
}

const (
	RegisterRstc2rFieldExtevnt8Shift = 28
	RegisterRstc2rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerRstc2rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerRstc2rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldExtevnt8Mask)
	}
}

const (
	RegisterRstc2rFieldExtevnt7Shift = 27
	RegisterRstc2rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerRstc2rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerRstc2rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldExtevnt7Mask)
	}
}

const (
	RegisterRstc2rFieldExtevnt6Shift = 26
	RegisterRstc2rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerRstc2rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerRstc2rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldExtevnt6Mask)
	}
}

const (
	RegisterRstc2rFieldExtevnt5Shift = 25
	RegisterRstc2rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerRstc2rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerRstc2rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldExtevnt5Mask)
	}
}

const (
	RegisterRstc2rFieldExtevnt4Shift = 24
	RegisterRstc2rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerRstc2rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerRstc2rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldExtevnt4Mask)
	}
}

const (
	RegisterRstc2rFieldExtevnt3Shift = 23
	RegisterRstc2rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerRstc2rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerRstc2rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldExtevnt3Mask)
	}
}

const (
	RegisterRstc2rFieldExtevnt2Shift = 22
	RegisterRstc2rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerRstc2rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerRstc2rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldExtevnt2Mask)
	}
}

const (
	RegisterRstc2rFieldExtevnt1Shift = 21
	RegisterRstc2rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerRstc2rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerRstc2rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldExtevnt1Mask)
	}
}

const (
	RegisterRstc2rFieldTimevnt9Shift = 20
	RegisterRstc2rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerRstc2rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerRstc2rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldTimevnt9Mask)
	}
}

const (
	RegisterRstc2rFieldTimevnt8Shift = 19
	RegisterRstc2rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerRstc2rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerRstc2rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldTimevnt8Mask)
	}
}

const (
	RegisterRstc2rFieldTimevnt7Shift = 18
	RegisterRstc2rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerRstc2rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerRstc2rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldTimevnt7Mask)
	}
}

const (
	RegisterRstc2rFieldTimevnt6Shift = 17
	RegisterRstc2rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerRstc2rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerRstc2rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldTimevnt6Mask)
	}
}

const (
	RegisterRstc2rFieldTimevnt5Shift = 16
	RegisterRstc2rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerRstc2rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerRstc2rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldTimevnt5Mask)
	}
}

const (
	RegisterRstc2rFieldTimevnt4Shift = 15
	RegisterRstc2rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerRstc2rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerRstc2rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldTimevnt4Mask)
	}
}

const (
	RegisterRstc2rFieldTimevnt3Shift = 14
	RegisterRstc2rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerRstc2rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerRstc2rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldTimevnt3Mask)
	}
}

const (
	RegisterRstc2rFieldTimevnt2Shift = 13
	RegisterRstc2rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerRstc2rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerRstc2rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldTimevnt2Mask)
	}
}

const (
	RegisterRstc2rFieldTimevnt1Shift = 12
	RegisterRstc2rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerRstc2rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerRstc2rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldTimevnt1Mask)
	}
}

const (
	RegisterRstc2rFieldMstcmp4Shift = 11
	RegisterRstc2rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerRstc2rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerRstc2rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldMstcmp4Mask)
	}
}

const (
	RegisterRstc2rFieldMstcmp3Shift = 10
	RegisterRstc2rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerRstc2rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerRstc2rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldMstcmp3Mask)
	}
}

const (
	RegisterRstc2rFieldMstcmp2Shift = 9
	RegisterRstc2rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerRstc2rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerRstc2rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldMstcmp2Mask)
	}
}

const (
	RegisterRstc2rFieldMstcmp1Shift = 8
	RegisterRstc2rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerRstc2rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerRstc2rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldMstcmp1Mask)
	}
}

const (
	RegisterRstc2rFieldMstperShift = 7
	RegisterRstc2rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerRstc2rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerRstc2rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldMstperMask)
	}
}

const (
	RegisterRstc2rFieldCmp4Shift = 6
	RegisterRstc2rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerRstc2rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerRstc2rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldCmp4Mask)
	}
}

const (
	RegisterRstc2rFieldCmp3Shift = 5
	RegisterRstc2rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerRstc2rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerRstc2rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldCmp3Mask)
	}
}

const (
	RegisterRstc2rFieldCmp2Shift = 4
	RegisterRstc2rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerRstc2rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerRstc2rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldCmp2Mask)
	}
}

const (
	RegisterRstc2rFieldCmp1Shift = 3
	RegisterRstc2rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerRstc2rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerRstc2rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldCmp1Mask)
	}
}

const (
	RegisterRstc2rFieldPerShift = 2
	RegisterRstc2rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerRstc2rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldPerMask) != 0
}

// SetPer PER
func (r *registerRstc2rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldPerMask)
	}
}

const (
	RegisterRstc2rFieldResyncShift = 1
	RegisterRstc2rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerRstc2rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerRstc2rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldResyncMask)
	}
}

const (
	RegisterRstc2rFieldSrtShift = 0
	RegisterRstc2rFieldSrtMask  = 0x1
)

// GetSrt SRT
func (r *registerRstc2rType) GetSrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstc2rFieldSrtMask) != 0
}

// SetSrt SRT
func (r *registerRstc2rType) SetSrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstc2rFieldSrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstc2rFieldSrtMask)
	}
}

// registerEefcr1Type Timerx External Event Filtering Register 1
type registerEefcr1Type uint32

const (
	RegisterEefcr1FieldEe5fltrShift = 25
	RegisterEefcr1FieldEe5fltrMask  = 0x1e000000
)

// GetEe5fltr External Event 5 filter
func (r *registerEefcr1Type) GetEe5fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefcr1FieldEe5fltrMask) >> RegisterEefcr1FieldEe5fltrShift)
}

// SetEe5fltr External Event 5 filter
func (r *registerEefcr1Type) SetEe5fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefcr1FieldEe5fltrMask)|(uint32(value)<<RegisterEefcr1FieldEe5fltrShift))
}

const (
	RegisterEefcr1FieldEe5ltchShift = 24
	RegisterEefcr1FieldEe5ltchMask  = 0x1000000
)

// GetEe5ltch External Event 5 latch
func (r *registerEefcr1Type) GetEe5ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefcr1FieldEe5ltchMask) != 0
}

// SetEe5ltch External Event 5 latch
func (r *registerEefcr1Type) SetEe5ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefcr1FieldEe5ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefcr1FieldEe5ltchMask)
	}
}

const (
	RegisterEefcr1FieldEe4fltrShift = 19
	RegisterEefcr1FieldEe4fltrMask  = 0x780000
)

// GetEe4fltr External Event 4 filter
func (r *registerEefcr1Type) GetEe4fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefcr1FieldEe4fltrMask) >> RegisterEefcr1FieldEe4fltrShift)
}

// SetEe4fltr External Event 4 filter
func (r *registerEefcr1Type) SetEe4fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefcr1FieldEe4fltrMask)|(uint32(value)<<RegisterEefcr1FieldEe4fltrShift))
}

const (
	RegisterEefcr1FieldEe4ltchShift = 18
	RegisterEefcr1FieldEe4ltchMask  = 0x40000
)

// GetEe4ltch External Event 4 latch
func (r *registerEefcr1Type) GetEe4ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefcr1FieldEe4ltchMask) != 0
}

// SetEe4ltch External Event 4 latch
func (r *registerEefcr1Type) SetEe4ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefcr1FieldEe4ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefcr1FieldEe4ltchMask)
	}
}

const (
	RegisterEefcr1FieldEe3fltrShift = 13
	RegisterEefcr1FieldEe3fltrMask  = 0x1e000
)

// GetEe3fltr External Event 3 filter
func (r *registerEefcr1Type) GetEe3fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefcr1FieldEe3fltrMask) >> RegisterEefcr1FieldEe3fltrShift)
}

// SetEe3fltr External Event 3 filter
func (r *registerEefcr1Type) SetEe3fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefcr1FieldEe3fltrMask)|(uint32(value)<<RegisterEefcr1FieldEe3fltrShift))
}

const (
	RegisterEefcr1FieldEe3ltchShift = 12
	RegisterEefcr1FieldEe3ltchMask  = 0x1000
)

// GetEe3ltch External Event 3 latch
func (r *registerEefcr1Type) GetEe3ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefcr1FieldEe3ltchMask) != 0
}

// SetEe3ltch External Event 3 latch
func (r *registerEefcr1Type) SetEe3ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefcr1FieldEe3ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefcr1FieldEe3ltchMask)
	}
}

const (
	RegisterEefcr1FieldEe2fltrShift = 7
	RegisterEefcr1FieldEe2fltrMask  = 0x780
)

// GetEe2fltr External Event 2 filter
func (r *registerEefcr1Type) GetEe2fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefcr1FieldEe2fltrMask) >> RegisterEefcr1FieldEe2fltrShift)
}

// SetEe2fltr External Event 2 filter
func (r *registerEefcr1Type) SetEe2fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefcr1FieldEe2fltrMask)|(uint32(value)<<RegisterEefcr1FieldEe2fltrShift))
}

const (
	RegisterEefcr1FieldEe2ltchShift = 6
	RegisterEefcr1FieldEe2ltchMask  = 0x40
)

// GetEe2ltch External Event 2 latch
func (r *registerEefcr1Type) GetEe2ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefcr1FieldEe2ltchMask) != 0
}

// SetEe2ltch External Event 2 latch
func (r *registerEefcr1Type) SetEe2ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefcr1FieldEe2ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefcr1FieldEe2ltchMask)
	}
}

const (
	RegisterEefcr1FieldEe1fltrShift = 1
	RegisterEefcr1FieldEe1fltrMask  = 0x1e
)

// GetEe1fltr External Event 1 filter
func (r *registerEefcr1Type) GetEe1fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefcr1FieldEe1fltrMask) >> RegisterEefcr1FieldEe1fltrShift)
}

// SetEe1fltr External Event 1 filter
func (r *registerEefcr1Type) SetEe1fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefcr1FieldEe1fltrMask)|(uint32(value)<<RegisterEefcr1FieldEe1fltrShift))
}

const (
	RegisterEefcr1FieldEe1ltchShift = 0
	RegisterEefcr1FieldEe1ltchMask  = 0x1
)

// GetEe1ltch External Event 1 latch
func (r *registerEefcr1Type) GetEe1ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefcr1FieldEe1ltchMask) != 0
}

// SetEe1ltch External Event 1 latch
func (r *registerEefcr1Type) SetEe1ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefcr1FieldEe1ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefcr1FieldEe1ltchMask)
	}
}

// registerEefcr2Type Timerx External Event Filtering Register 2
type registerEefcr2Type uint32

const (
	RegisterEefcr2FieldEe10fltrShift = 25
	RegisterEefcr2FieldEe10fltrMask  = 0x1e000000
)

// GetEe10fltr External Event 10 filter
func (r *registerEefcr2Type) GetEe10fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefcr2FieldEe10fltrMask) >> RegisterEefcr2FieldEe10fltrShift)
}

// SetEe10fltr External Event 10 filter
func (r *registerEefcr2Type) SetEe10fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefcr2FieldEe10fltrMask)|(uint32(value)<<RegisterEefcr2FieldEe10fltrShift))
}

const (
	RegisterEefcr2FieldEe10ltchShift = 24
	RegisterEefcr2FieldEe10ltchMask  = 0x1000000
)

// GetEe10ltch External Event 10 latch
func (r *registerEefcr2Type) GetEe10ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefcr2FieldEe10ltchMask) != 0
}

// SetEe10ltch External Event 10 latch
func (r *registerEefcr2Type) SetEe10ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefcr2FieldEe10ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefcr2FieldEe10ltchMask)
	}
}

const (
	RegisterEefcr2FieldEe9fltrShift = 19
	RegisterEefcr2FieldEe9fltrMask  = 0x780000
)

// GetEe9fltr External Event 9 filter
func (r *registerEefcr2Type) GetEe9fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefcr2FieldEe9fltrMask) >> RegisterEefcr2FieldEe9fltrShift)
}

// SetEe9fltr External Event 9 filter
func (r *registerEefcr2Type) SetEe9fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefcr2FieldEe9fltrMask)|(uint32(value)<<RegisterEefcr2FieldEe9fltrShift))
}

const (
	RegisterEefcr2FieldEe9ltchShift = 18
	RegisterEefcr2FieldEe9ltchMask  = 0x40000
)

// GetEe9ltch External Event 9 latch
func (r *registerEefcr2Type) GetEe9ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefcr2FieldEe9ltchMask) != 0
}

// SetEe9ltch External Event 9 latch
func (r *registerEefcr2Type) SetEe9ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefcr2FieldEe9ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefcr2FieldEe9ltchMask)
	}
}

const (
	RegisterEefcr2FieldEe8fltrShift = 13
	RegisterEefcr2FieldEe8fltrMask  = 0x1e000
)

// GetEe8fltr External Event 8 filter
func (r *registerEefcr2Type) GetEe8fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefcr2FieldEe8fltrMask) >> RegisterEefcr2FieldEe8fltrShift)
}

// SetEe8fltr External Event 8 filter
func (r *registerEefcr2Type) SetEe8fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefcr2FieldEe8fltrMask)|(uint32(value)<<RegisterEefcr2FieldEe8fltrShift))
}

const (
	RegisterEefcr2FieldEe8ltchShift = 12
	RegisterEefcr2FieldEe8ltchMask  = 0x1000
)

// GetEe8ltch External Event 8 latch
func (r *registerEefcr2Type) GetEe8ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefcr2FieldEe8ltchMask) != 0
}

// SetEe8ltch External Event 8 latch
func (r *registerEefcr2Type) SetEe8ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefcr2FieldEe8ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefcr2FieldEe8ltchMask)
	}
}

const (
	RegisterEefcr2FieldEe7fltrShift = 7
	RegisterEefcr2FieldEe7fltrMask  = 0x780
)

// GetEe7fltr External Event 7 filter
func (r *registerEefcr2Type) GetEe7fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefcr2FieldEe7fltrMask) >> RegisterEefcr2FieldEe7fltrShift)
}

// SetEe7fltr External Event 7 filter
func (r *registerEefcr2Type) SetEe7fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefcr2FieldEe7fltrMask)|(uint32(value)<<RegisterEefcr2FieldEe7fltrShift))
}

const (
	RegisterEefcr2FieldEe7ltchShift = 6
	RegisterEefcr2FieldEe7ltchMask  = 0x40
)

// GetEe7ltch External Event 7 latch
func (r *registerEefcr2Type) GetEe7ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefcr2FieldEe7ltchMask) != 0
}

// SetEe7ltch External Event 7 latch
func (r *registerEefcr2Type) SetEe7ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefcr2FieldEe7ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefcr2FieldEe7ltchMask)
	}
}

const (
	RegisterEefcr2FieldEe6fltrShift = 1
	RegisterEefcr2FieldEe6fltrMask  = 0x1e
)

// GetEe6fltr External Event 6 filter
func (r *registerEefcr2Type) GetEe6fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefcr2FieldEe6fltrMask) >> RegisterEefcr2FieldEe6fltrShift)
}

// SetEe6fltr External Event 6 filter
func (r *registerEefcr2Type) SetEe6fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefcr2FieldEe6fltrMask)|(uint32(value)<<RegisterEefcr2FieldEe6fltrShift))
}

const (
	RegisterEefcr2FieldEe6ltchShift = 0
	RegisterEefcr2FieldEe6ltchMask  = 0x1
)

// GetEe6ltch External Event 6 latch
func (r *registerEefcr2Type) GetEe6ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefcr2FieldEe6ltchMask) != 0
}

// SetEe6ltch External Event 6 latch
func (r *registerEefcr2Type) SetEe6ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefcr2FieldEe6ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefcr2FieldEe6ltchMask)
	}
}

// registerRstcrType TimerA Reset Register
type registerRstcrType uint32

const (
	RegisterRstcrFieldTimecmp4Shift = 30
	RegisterRstcrFieldTimecmp4Mask  = 0x40000000
)

// GetTimecmp4 Timer E Compare 4
func (r *registerRstcrType) GetTimecmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldTimecmp4Mask) != 0
}

// SetTimecmp4 Timer E Compare 4
func (r *registerRstcrType) SetTimecmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldTimecmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldTimecmp4Mask)
	}
}

const (
	RegisterRstcrFieldTimecmp2Shift = 29
	RegisterRstcrFieldTimecmp2Mask  = 0x20000000
)

// GetTimecmp2 Timer E Compare 2
func (r *registerRstcrType) GetTimecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldTimecmp2Mask) != 0
}

// SetTimecmp2 Timer E Compare 2
func (r *registerRstcrType) SetTimecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldTimecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldTimecmp2Mask)
	}
}

const (
	RegisterRstcrFieldTimecmp1Shift = 28
	RegisterRstcrFieldTimecmp1Mask  = 0x10000000
)

// GetTimecmp1 Timer E Compare 1
func (r *registerRstcrType) GetTimecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldTimecmp1Mask) != 0
}

// SetTimecmp1 Timer E Compare 1
func (r *registerRstcrType) SetTimecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldTimecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldTimecmp1Mask)
	}
}

const (
	RegisterRstcrFieldTimdcmp4Shift = 27
	RegisterRstcrFieldTimdcmp4Mask  = 0x8000000
)

// GetTimdcmp4 Timer D Compare 4
func (r *registerRstcrType) GetTimdcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldTimdcmp4Mask) != 0
}

// SetTimdcmp4 Timer D Compare 4
func (r *registerRstcrType) SetTimdcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldTimdcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldTimdcmp4Mask)
	}
}

const (
	RegisterRstcrFieldTimdcmp2Shift = 26
	RegisterRstcrFieldTimdcmp2Mask  = 0x4000000
)

// GetTimdcmp2 Timer D Compare 2
func (r *registerRstcrType) GetTimdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldTimdcmp2Mask) != 0
}

// SetTimdcmp2 Timer D Compare 2
func (r *registerRstcrType) SetTimdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldTimdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldTimdcmp2Mask)
	}
}

const (
	RegisterRstcrFieldTimdcmp1Shift = 25
	RegisterRstcrFieldTimdcmp1Mask  = 0x2000000
)

// GetTimdcmp1 Timer D Compare 1
func (r *registerRstcrType) GetTimdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldTimdcmp1Mask) != 0
}

// SetTimdcmp1 Timer D Compare 1
func (r *registerRstcrType) SetTimdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldTimdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldTimdcmp1Mask)
	}
}

const (
	RegisterRstcrFieldTimbcmp4Shift = 24
	RegisterRstcrFieldTimbcmp4Mask  = 0x1000000
)

// GetTimbcmp4 Timer B Compare 4
func (r *registerRstcrType) GetTimbcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldTimbcmp4Mask) != 0
}

// SetTimbcmp4 Timer B Compare 4
func (r *registerRstcrType) SetTimbcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldTimbcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldTimbcmp4Mask)
	}
}

const (
	RegisterRstcrFieldTimbcmp2Shift = 23
	RegisterRstcrFieldTimbcmp2Mask  = 0x800000
)

// GetTimbcmp2 Timer B Compare 2
func (r *registerRstcrType) GetTimbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldTimbcmp2Mask) != 0
}

// SetTimbcmp2 Timer B Compare 2
func (r *registerRstcrType) SetTimbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldTimbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldTimbcmp2Mask)
	}
}

const (
	RegisterRstcrFieldTimbcmp1Shift = 22
	RegisterRstcrFieldTimbcmp1Mask  = 0x400000
)

// GetTimbcmp1 Timer B Compare 1
func (r *registerRstcrType) GetTimbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldTimbcmp1Mask) != 0
}

// SetTimbcmp1 Timer B Compare 1
func (r *registerRstcrType) SetTimbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldTimbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldTimbcmp1Mask)
	}
}

const (
	RegisterRstcrFieldTimacmp4Shift = 21
	RegisterRstcrFieldTimacmp4Mask  = 0x200000
)

// GetTimacmp4 Timer A Compare 4
func (r *registerRstcrType) GetTimacmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldTimacmp4Mask) != 0
}

// SetTimacmp4 Timer A Compare 4
func (r *registerRstcrType) SetTimacmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldTimacmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldTimacmp4Mask)
	}
}

const (
	RegisterRstcrFieldTimacmp2Shift = 20
	RegisterRstcrFieldTimacmp2Mask  = 0x100000
)

// GetTimacmp2 Timer A Compare 2
func (r *registerRstcrType) GetTimacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldTimacmp2Mask) != 0
}

// SetTimacmp2 Timer A Compare 2
func (r *registerRstcrType) SetTimacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldTimacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldTimacmp2Mask)
	}
}

const (
	RegisterRstcrFieldTimacmp1Shift = 19
	RegisterRstcrFieldTimacmp1Mask  = 0x80000
)

// GetTimacmp1 Timer A Compare 1
func (r *registerRstcrType) GetTimacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldTimacmp1Mask) != 0
}

// SetTimacmp1 Timer A Compare 1
func (r *registerRstcrType) SetTimacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldTimacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldTimacmp1Mask)
	}
}

const (
	RegisterRstcrFieldExtevnt10Shift = 18
	RegisterRstcrFieldExtevnt10Mask  = 0x40000
)

// GetExtevnt10 External Event 10
func (r *registerRstcrType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldExtevnt10Mask) != 0
}

// SetExtevnt10 External Event 10
func (r *registerRstcrType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldExtevnt10Mask)
	}
}

const (
	RegisterRstcrFieldExtevnt9Shift = 17
	RegisterRstcrFieldExtevnt9Mask  = 0x20000
)

// GetExtevnt9 External Event 9
func (r *registerRstcrType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldExtevnt9Mask) != 0
}

// SetExtevnt9 External Event 9
func (r *registerRstcrType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldExtevnt9Mask)
	}
}

const (
	RegisterRstcrFieldExtevnt8Shift = 16
	RegisterRstcrFieldExtevnt8Mask  = 0x10000
)

// GetExtevnt8 External Event 8
func (r *registerRstcrType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldExtevnt8Mask) != 0
}

// SetExtevnt8 External Event 8
func (r *registerRstcrType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldExtevnt8Mask)
	}
}

const (
	RegisterRstcrFieldExtevnt7Shift = 15
	RegisterRstcrFieldExtevnt7Mask  = 0x8000
)

// GetExtevnt7 External Event 7
func (r *registerRstcrType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldExtevnt7Mask) != 0
}

// SetExtevnt7 External Event 7
func (r *registerRstcrType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldExtevnt7Mask)
	}
}

const (
	RegisterRstcrFieldExtevnt6Shift = 14
	RegisterRstcrFieldExtevnt6Mask  = 0x4000
)

// GetExtevnt6 External Event 6
func (r *registerRstcrType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldExtevnt6Mask) != 0
}

// SetExtevnt6 External Event 6
func (r *registerRstcrType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldExtevnt6Mask)
	}
}

const (
	RegisterRstcrFieldExtevnt5Shift = 13
	RegisterRstcrFieldExtevnt5Mask  = 0x2000
)

// GetExtevnt5 External Event 5
func (r *registerRstcrType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldExtevnt5Mask) != 0
}

// SetExtevnt5 External Event 5
func (r *registerRstcrType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldExtevnt5Mask)
	}
}

const (
	RegisterRstcrFieldExtevnt4Shift = 12
	RegisterRstcrFieldExtevnt4Mask  = 0x1000
)

// GetExtevnt4 External Event 4
func (r *registerRstcrType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldExtevnt4Mask) != 0
}

// SetExtevnt4 External Event 4
func (r *registerRstcrType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldExtevnt4Mask)
	}
}

const (
	RegisterRstcrFieldExtevnt3Shift = 11
	RegisterRstcrFieldExtevnt3Mask  = 0x800
)

// GetExtevnt3 External Event 3
func (r *registerRstcrType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldExtevnt3Mask) != 0
}

// SetExtevnt3 External Event 3
func (r *registerRstcrType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldExtevnt3Mask)
	}
}

const (
	RegisterRstcrFieldExtevnt2Shift = 10
	RegisterRstcrFieldExtevnt2Mask  = 0x400
)

// GetExtevnt2 External Event 2
func (r *registerRstcrType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldExtevnt2Mask) != 0
}

// SetExtevnt2 External Event 2
func (r *registerRstcrType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldExtevnt2Mask)
	}
}

const (
	RegisterRstcrFieldExtevnt1Shift = 9
	RegisterRstcrFieldExtevnt1Mask  = 0x200
)

// GetExtevnt1 External Event 1
func (r *registerRstcrType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldExtevnt1Mask) != 0
}

// SetExtevnt1 External Event 1
func (r *registerRstcrType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldExtevnt1Mask)
	}
}

const (
	RegisterRstcrFieldMstcmp4Shift = 8
	RegisterRstcrFieldMstcmp4Mask  = 0x100
)

// GetMstcmp4 Master compare 4
func (r *registerRstcrType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldMstcmp4Mask) != 0
}

// SetMstcmp4 Master compare 4
func (r *registerRstcrType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldMstcmp4Mask)
	}
}

const (
	RegisterRstcrFieldMstcmp3Shift = 7
	RegisterRstcrFieldMstcmp3Mask  = 0x80
)

// GetMstcmp3 Master compare 3
func (r *registerRstcrType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldMstcmp3Mask) != 0
}

// SetMstcmp3 Master compare 3
func (r *registerRstcrType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldMstcmp3Mask)
	}
}

const (
	RegisterRstcrFieldMstcmp2Shift = 6
	RegisterRstcrFieldMstcmp2Mask  = 0x40
)

// GetMstcmp2 Master compare 2
func (r *registerRstcrType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldMstcmp2Mask) != 0
}

// SetMstcmp2 Master compare 2
func (r *registerRstcrType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldMstcmp2Mask)
	}
}

const (
	RegisterRstcrFieldMstcmp1Shift = 5
	RegisterRstcrFieldMstcmp1Mask  = 0x20
)

// GetMstcmp1 Master compare 1
func (r *registerRstcrType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldMstcmp1Mask) != 0
}

// SetMstcmp1 Master compare 1
func (r *registerRstcrType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldMstcmp1Mask)
	}
}

const (
	RegisterRstcrFieldMstperShift = 4
	RegisterRstcrFieldMstperMask  = 0x10
)

// GetMstper Master timer Period
func (r *registerRstcrType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldMstperMask) != 0
}

// SetMstper Master timer Period
func (r *registerRstcrType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldMstperMask)
	}
}

const (
	RegisterRstcrFieldCmp4Shift = 3
	RegisterRstcrFieldCmp4Mask  = 0x8
)

// GetCmp4 Timer A compare 4 reset
func (r *registerRstcrType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldCmp4Mask) != 0
}

// SetCmp4 Timer A compare 4 reset
func (r *registerRstcrType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldCmp4Mask)
	}
}

const (
	RegisterRstcrFieldCmp2Shift = 2
	RegisterRstcrFieldCmp2Mask  = 0x4
)

// GetCmp2 Timer A compare 2 reset
func (r *registerRstcrType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldCmp2Mask) != 0
}

// SetCmp2 Timer A compare 2 reset
func (r *registerRstcrType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldCmp2Mask)
	}
}

const (
	RegisterRstcrFieldUpdtShift = 1
	RegisterRstcrFieldUpdtMask  = 0x2
)

// GetUpdt Timer A Update reset
func (r *registerRstcrType) GetUpdt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstcrFieldUpdtMask) != 0
}

// SetUpdt Timer A Update reset
func (r *registerRstcrType) SetUpdt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstcrFieldUpdtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstcrFieldUpdtMask)
	}
}

// registerChpcrType Timerx Chopper Register
type registerChpcrType uint32

const (
	RegisterChpcrFieldStrtpwShift = 7
	RegisterChpcrFieldStrtpwMask  = 0x780
)

// GetStrtpw STRTPW
func (r *registerChpcrType) GetStrtpw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChpcrFieldStrtpwMask) >> RegisterChpcrFieldStrtpwShift)
}

// SetStrtpw STRTPW
func (r *registerChpcrType) SetStrtpw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChpcrFieldStrtpwMask)|(uint32(value)<<RegisterChpcrFieldStrtpwShift))
}

const (
	RegisterChpcrFieldChpdtyShift = 4
	RegisterChpcrFieldChpdtyMask  = 0x70
)

// GetChpdty Timerx chopper duty cycle value
func (r *registerChpcrType) GetChpdty() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChpcrFieldChpdtyMask) >> RegisterChpcrFieldChpdtyShift)
}

// SetChpdty Timerx chopper duty cycle value
func (r *registerChpcrType) SetChpdty(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChpcrFieldChpdtyMask)|(uint32(value)<<RegisterChpcrFieldChpdtyShift))
}

const (
	RegisterChpcrFieldChpfrqShift = 0
	RegisterChpcrFieldChpfrqMask  = 0xf
)

// GetChpfrq Timerx carrier frequency value
func (r *registerChpcrType) GetChpfrq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChpcrFieldChpfrqMask) >> RegisterChpcrFieldChpfrqShift)
}

// SetChpfrq Timerx carrier frequency value
func (r *registerChpcrType) SetChpfrq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChpcrFieldChpfrqMask)|(uint32(value)<<RegisterChpcrFieldChpfrqShift))
}

// registerCpt1ccrType Timerx Capture 2 Control Register
type registerCpt1ccrType uint32

const (
	RegisterCpt1ccrFieldTecmp2Shift = 31
	RegisterCpt1ccrFieldTecmp2Mask  = 0x80000000
)

// GetTecmp2 Timer E Compare 2
func (r *registerCpt1ccrType) GetTecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTecmp2Mask) != 0
}

// SetTecmp2 Timer E Compare 2
func (r *registerCpt1ccrType) SetTecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTecmp2Mask)
	}
}

const (
	RegisterCpt1ccrFieldTecmp1Shift = 30
	RegisterCpt1ccrFieldTecmp1Mask  = 0x40000000
)

// GetTecmp1 Timer E Compare 1
func (r *registerCpt1ccrType) GetTecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTecmp1Mask) != 0
}

// SetTecmp1 Timer E Compare 1
func (r *registerCpt1ccrType) SetTecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTecmp1Mask)
	}
}

const (
	RegisterCpt1ccrFieldTe1rstShift = 29
	RegisterCpt1ccrFieldTe1rstMask  = 0x20000000
)

// GetTe1rst Timer E output 1 Reset
func (r *registerCpt1ccrType) GetTe1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTe1rstMask) != 0
}

// SetTe1rst Timer E output 1 Reset
func (r *registerCpt1ccrType) SetTe1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTe1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTe1rstMask)
	}
}

const (
	RegisterCpt1ccrFieldTe1setShift = 28
	RegisterCpt1ccrFieldTe1setMask  = 0x10000000
)

// GetTe1set Timer E output 1 Set
func (r *registerCpt1ccrType) GetTe1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTe1setMask) != 0
}

// SetTe1set Timer E output 1 Set
func (r *registerCpt1ccrType) SetTe1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTe1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTe1setMask)
	}
}

const (
	RegisterCpt1ccrFieldTdcmp2Shift = 27
	RegisterCpt1ccrFieldTdcmp2Mask  = 0x8000000
)

// GetTdcmp2 Timer D Compare 2
func (r *registerCpt1ccrType) GetTdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTdcmp2Mask) != 0
}

// SetTdcmp2 Timer D Compare 2
func (r *registerCpt1ccrType) SetTdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTdcmp2Mask)
	}
}

const (
	RegisterCpt1ccrFieldTdcmp1Shift = 26
	RegisterCpt1ccrFieldTdcmp1Mask  = 0x4000000
)

// GetTdcmp1 Timer D Compare 1
func (r *registerCpt1ccrType) GetTdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTdcmp1Mask) != 0
}

// SetTdcmp1 Timer D Compare 1
func (r *registerCpt1ccrType) SetTdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTdcmp1Mask)
	}
}

const (
	RegisterCpt1ccrFieldTd1rstShift = 25
	RegisterCpt1ccrFieldTd1rstMask  = 0x2000000
)

// GetTd1rst Timer D output 1 Reset
func (r *registerCpt1ccrType) GetTd1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTd1rstMask) != 0
}

// SetTd1rst Timer D output 1 Reset
func (r *registerCpt1ccrType) SetTd1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTd1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTd1rstMask)
	}
}

const (
	RegisterCpt1ccrFieldTd1setShift = 24
	RegisterCpt1ccrFieldTd1setMask  = 0x1000000
)

// GetTd1set Timer D output 1 Set
func (r *registerCpt1ccrType) GetTd1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTd1setMask) != 0
}

// SetTd1set Timer D output 1 Set
func (r *registerCpt1ccrType) SetTd1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTd1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTd1setMask)
	}
}

const (
	RegisterCpt1ccrFieldTbcmp2Shift = 19
	RegisterCpt1ccrFieldTbcmp2Mask  = 0x80000
)

// GetTbcmp2 Timer B Compare 2
func (r *registerCpt1ccrType) GetTbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTbcmp2Mask) != 0
}

// SetTbcmp2 Timer B Compare 2
func (r *registerCpt1ccrType) SetTbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTbcmp2Mask)
	}
}

const (
	RegisterCpt1ccrFieldTbcmp1Shift = 18
	RegisterCpt1ccrFieldTbcmp1Mask  = 0x40000
)

// GetTbcmp1 Timer B Compare 1
func (r *registerCpt1ccrType) GetTbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTbcmp1Mask) != 0
}

// SetTbcmp1 Timer B Compare 1
func (r *registerCpt1ccrType) SetTbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTbcmp1Mask)
	}
}

const (
	RegisterCpt1ccrFieldTb1rstShift = 17
	RegisterCpt1ccrFieldTb1rstMask  = 0x20000
)

// GetTb1rst Timer B output 1 Reset
func (r *registerCpt1ccrType) GetTb1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTb1rstMask) != 0
}

// SetTb1rst Timer B output 1 Reset
func (r *registerCpt1ccrType) SetTb1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTb1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTb1rstMask)
	}
}

const (
	RegisterCpt1ccrFieldTb1setShift = 16
	RegisterCpt1ccrFieldTb1setMask  = 0x10000
)

// GetTb1set Timer B output 1 Set
func (r *registerCpt1ccrType) GetTb1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTb1setMask) != 0
}

// SetTb1set Timer B output 1 Set
func (r *registerCpt1ccrType) SetTb1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTb1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTb1setMask)
	}
}

const (
	RegisterCpt1ccrFieldTacmp2Shift = 15
	RegisterCpt1ccrFieldTacmp2Mask  = 0x8000
)

// GetTacmp2 Timer A Compare 2
func (r *registerCpt1ccrType) GetTacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTacmp2Mask) != 0
}

// SetTacmp2 Timer A Compare 2
func (r *registerCpt1ccrType) SetTacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTacmp2Mask)
	}
}

const (
	RegisterCpt1ccrFieldTacmp1Shift = 14
	RegisterCpt1ccrFieldTacmp1Mask  = 0x4000
)

// GetTacmp1 Timer A Compare 1
func (r *registerCpt1ccrType) GetTacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTacmp1Mask) != 0
}

// SetTacmp1 Timer A Compare 1
func (r *registerCpt1ccrType) SetTacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTacmp1Mask)
	}
}

const (
	RegisterCpt1ccrFieldTa1rstShift = 13
	RegisterCpt1ccrFieldTa1rstMask  = 0x2000
)

// GetTa1rst Timer A output 1 Reset
func (r *registerCpt1ccrType) GetTa1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTa1rstMask) != 0
}

// SetTa1rst Timer A output 1 Reset
func (r *registerCpt1ccrType) SetTa1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTa1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTa1rstMask)
	}
}

const (
	RegisterCpt1ccrFieldTa1setShift = 12
	RegisterCpt1ccrFieldTa1setMask  = 0x1000
)

// GetTa1set Timer A output 1 Set
func (r *registerCpt1ccrType) GetTa1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldTa1setMask) != 0
}

// SetTa1set Timer A output 1 Set
func (r *registerCpt1ccrType) SetTa1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldTa1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldTa1setMask)
	}
}

const (
	RegisterCpt1ccrFieldExev10cptShift = 11
	RegisterCpt1ccrFieldExev10cptMask  = 0x800
)

// GetExev10cpt External Event 10 Capture
func (r *registerCpt1ccrType) GetExev10cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldExev10cptMask) != 0
}

// SetExev10cpt External Event 10 Capture
func (r *registerCpt1ccrType) SetExev10cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldExev10cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldExev10cptMask)
	}
}

const (
	RegisterCpt1ccrFieldExev9cptShift = 10
	RegisterCpt1ccrFieldExev9cptMask  = 0x400
)

// GetExev9cpt External Event 9 Capture
func (r *registerCpt1ccrType) GetExev9cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldExev9cptMask) != 0
}

// SetExev9cpt External Event 9 Capture
func (r *registerCpt1ccrType) SetExev9cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldExev9cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldExev9cptMask)
	}
}

const (
	RegisterCpt1ccrFieldExev8cptShift = 9
	RegisterCpt1ccrFieldExev8cptMask  = 0x200
)

// GetExev8cpt External Event 8 Capture
func (r *registerCpt1ccrType) GetExev8cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldExev8cptMask) != 0
}

// SetExev8cpt External Event 8 Capture
func (r *registerCpt1ccrType) SetExev8cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldExev8cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldExev8cptMask)
	}
}

const (
	RegisterCpt1ccrFieldExev7cptShift = 8
	RegisterCpt1ccrFieldExev7cptMask  = 0x100
)

// GetExev7cpt External Event 7 Capture
func (r *registerCpt1ccrType) GetExev7cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldExev7cptMask) != 0
}

// SetExev7cpt External Event 7 Capture
func (r *registerCpt1ccrType) SetExev7cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldExev7cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldExev7cptMask)
	}
}

const (
	RegisterCpt1ccrFieldExev6cptShift = 7
	RegisterCpt1ccrFieldExev6cptMask  = 0x80
)

// GetExev6cpt External Event 6 Capture
func (r *registerCpt1ccrType) GetExev6cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldExev6cptMask) != 0
}

// SetExev6cpt External Event 6 Capture
func (r *registerCpt1ccrType) SetExev6cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldExev6cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldExev6cptMask)
	}
}

const (
	RegisterCpt1ccrFieldExev5cptShift = 6
	RegisterCpt1ccrFieldExev5cptMask  = 0x40
)

// GetExev5cpt External Event 5 Capture
func (r *registerCpt1ccrType) GetExev5cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldExev5cptMask) != 0
}

// SetExev5cpt External Event 5 Capture
func (r *registerCpt1ccrType) SetExev5cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldExev5cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldExev5cptMask)
	}
}

const (
	RegisterCpt1ccrFieldExev4cptShift = 5
	RegisterCpt1ccrFieldExev4cptMask  = 0x20
)

// GetExev4cpt External Event 4 Capture
func (r *registerCpt1ccrType) GetExev4cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldExev4cptMask) != 0
}

// SetExev4cpt External Event 4 Capture
func (r *registerCpt1ccrType) SetExev4cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldExev4cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldExev4cptMask)
	}
}

const (
	RegisterCpt1ccrFieldExev3cptShift = 4
	RegisterCpt1ccrFieldExev3cptMask  = 0x10
)

// GetExev3cpt External Event 3 Capture
func (r *registerCpt1ccrType) GetExev3cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldExev3cptMask) != 0
}

// SetExev3cpt External Event 3 Capture
func (r *registerCpt1ccrType) SetExev3cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldExev3cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldExev3cptMask)
	}
}

const (
	RegisterCpt1ccrFieldExev2cptShift = 3
	RegisterCpt1ccrFieldExev2cptMask  = 0x8
)

// GetExev2cpt External Event 2 Capture
func (r *registerCpt1ccrType) GetExev2cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldExev2cptMask) != 0
}

// SetExev2cpt External Event 2 Capture
func (r *registerCpt1ccrType) SetExev2cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldExev2cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldExev2cptMask)
	}
}

const (
	RegisterCpt1ccrFieldExev1cptShift = 2
	RegisterCpt1ccrFieldExev1cptMask  = 0x4
)

// GetExev1cpt External Event 1 Capture
func (r *registerCpt1ccrType) GetExev1cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldExev1cptMask) != 0
}

// SetExev1cpt External Event 1 Capture
func (r *registerCpt1ccrType) SetExev1cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldExev1cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldExev1cptMask)
	}
}

const (
	RegisterCpt1ccrFieldUdpcptShift = 1
	RegisterCpt1ccrFieldUdpcptMask  = 0x2
)

// GetUdpcpt Update Capture
func (r *registerCpt1ccrType) GetUdpcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldUdpcptMask) != 0
}

// SetUdpcpt Update Capture
func (r *registerCpt1ccrType) SetUdpcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldUdpcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldUdpcptMask)
	}
}

const (
	RegisterCpt1ccrFieldSwcptShift = 0
	RegisterCpt1ccrFieldSwcptMask  = 0x1
)

// GetSwcpt Software Capture
func (r *registerCpt1ccrType) GetSwcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1ccrFieldSwcptMask) != 0
}

// SetSwcpt Software Capture
func (r *registerCpt1ccrType) SetSwcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1ccrFieldSwcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1ccrFieldSwcptMask)
	}
}

// registerCpt2ccrType CPT2xCR
type registerCpt2ccrType uint32

const (
	RegisterCpt2ccrFieldTecmp2Shift = 31
	RegisterCpt2ccrFieldTecmp2Mask  = 0x80000000
)

// GetTecmp2 Timer E Compare 2
func (r *registerCpt2ccrType) GetTecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTecmp2Mask) != 0
}

// SetTecmp2 Timer E Compare 2
func (r *registerCpt2ccrType) SetTecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTecmp2Mask)
	}
}

const (
	RegisterCpt2ccrFieldTecmp1Shift = 30
	RegisterCpt2ccrFieldTecmp1Mask  = 0x40000000
)

// GetTecmp1 Timer E Compare 1
func (r *registerCpt2ccrType) GetTecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTecmp1Mask) != 0
}

// SetTecmp1 Timer E Compare 1
func (r *registerCpt2ccrType) SetTecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTecmp1Mask)
	}
}

const (
	RegisterCpt2ccrFieldTe1rstShift = 29
	RegisterCpt2ccrFieldTe1rstMask  = 0x20000000
)

// GetTe1rst Timer E output 1 Reset
func (r *registerCpt2ccrType) GetTe1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTe1rstMask) != 0
}

// SetTe1rst Timer E output 1 Reset
func (r *registerCpt2ccrType) SetTe1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTe1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTe1rstMask)
	}
}

const (
	RegisterCpt2ccrFieldTe1setShift = 28
	RegisterCpt2ccrFieldTe1setMask  = 0x10000000
)

// GetTe1set Timer E output 1 Set
func (r *registerCpt2ccrType) GetTe1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTe1setMask) != 0
}

// SetTe1set Timer E output 1 Set
func (r *registerCpt2ccrType) SetTe1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTe1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTe1setMask)
	}
}

const (
	RegisterCpt2ccrFieldTdcmp2Shift = 27
	RegisterCpt2ccrFieldTdcmp2Mask  = 0x8000000
)

// GetTdcmp2 Timer D Compare 2
func (r *registerCpt2ccrType) GetTdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTdcmp2Mask) != 0
}

// SetTdcmp2 Timer D Compare 2
func (r *registerCpt2ccrType) SetTdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTdcmp2Mask)
	}
}

const (
	RegisterCpt2ccrFieldTdcmp1Shift = 26
	RegisterCpt2ccrFieldTdcmp1Mask  = 0x4000000
)

// GetTdcmp1 Timer D Compare 1
func (r *registerCpt2ccrType) GetTdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTdcmp1Mask) != 0
}

// SetTdcmp1 Timer D Compare 1
func (r *registerCpt2ccrType) SetTdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTdcmp1Mask)
	}
}

const (
	RegisterCpt2ccrFieldTd1rstShift = 25
	RegisterCpt2ccrFieldTd1rstMask  = 0x2000000
)

// GetTd1rst Timer D output 1 Reset
func (r *registerCpt2ccrType) GetTd1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTd1rstMask) != 0
}

// SetTd1rst Timer D output 1 Reset
func (r *registerCpt2ccrType) SetTd1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTd1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTd1rstMask)
	}
}

const (
	RegisterCpt2ccrFieldTd1setShift = 24
	RegisterCpt2ccrFieldTd1setMask  = 0x1000000
)

// GetTd1set Timer D output 1 Set
func (r *registerCpt2ccrType) GetTd1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTd1setMask) != 0
}

// SetTd1set Timer D output 1 Set
func (r *registerCpt2ccrType) SetTd1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTd1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTd1setMask)
	}
}

const (
	RegisterCpt2ccrFieldTbcmp2Shift = 19
	RegisterCpt2ccrFieldTbcmp2Mask  = 0x80000
)

// GetTbcmp2 Timer B Compare 2
func (r *registerCpt2ccrType) GetTbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTbcmp2Mask) != 0
}

// SetTbcmp2 Timer B Compare 2
func (r *registerCpt2ccrType) SetTbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTbcmp2Mask)
	}
}

const (
	RegisterCpt2ccrFieldTbcmp1Shift = 18
	RegisterCpt2ccrFieldTbcmp1Mask  = 0x40000
)

// GetTbcmp1 Timer B Compare 1
func (r *registerCpt2ccrType) GetTbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTbcmp1Mask) != 0
}

// SetTbcmp1 Timer B Compare 1
func (r *registerCpt2ccrType) SetTbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTbcmp1Mask)
	}
}

const (
	RegisterCpt2ccrFieldTb1rstShift = 17
	RegisterCpt2ccrFieldTb1rstMask  = 0x20000
)

// GetTb1rst Timer B output 1 Reset
func (r *registerCpt2ccrType) GetTb1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTb1rstMask) != 0
}

// SetTb1rst Timer B output 1 Reset
func (r *registerCpt2ccrType) SetTb1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTb1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTb1rstMask)
	}
}

const (
	RegisterCpt2ccrFieldTb1setShift = 16
	RegisterCpt2ccrFieldTb1setMask  = 0x10000
)

// GetTb1set Timer B output 1 Set
func (r *registerCpt2ccrType) GetTb1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTb1setMask) != 0
}

// SetTb1set Timer B output 1 Set
func (r *registerCpt2ccrType) SetTb1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTb1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTb1setMask)
	}
}

const (
	RegisterCpt2ccrFieldTacmp2Shift = 15
	RegisterCpt2ccrFieldTacmp2Mask  = 0x8000
)

// GetTacmp2 Timer A Compare 2
func (r *registerCpt2ccrType) GetTacmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTacmp2Mask) != 0
}

// SetTacmp2 Timer A Compare 2
func (r *registerCpt2ccrType) SetTacmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTacmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTacmp2Mask)
	}
}

const (
	RegisterCpt2ccrFieldTacmp1Shift = 14
	RegisterCpt2ccrFieldTacmp1Mask  = 0x4000
)

// GetTacmp1 Timer A Compare 1
func (r *registerCpt2ccrType) GetTacmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTacmp1Mask) != 0
}

// SetTacmp1 Timer A Compare 1
func (r *registerCpt2ccrType) SetTacmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTacmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTacmp1Mask)
	}
}

const (
	RegisterCpt2ccrFieldTa1rstShift = 13
	RegisterCpt2ccrFieldTa1rstMask  = 0x2000
)

// GetTa1rst Timer A output 1 Reset
func (r *registerCpt2ccrType) GetTa1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTa1rstMask) != 0
}

// SetTa1rst Timer A output 1 Reset
func (r *registerCpt2ccrType) SetTa1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTa1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTa1rstMask)
	}
}

const (
	RegisterCpt2ccrFieldTa1setShift = 12
	RegisterCpt2ccrFieldTa1setMask  = 0x1000
)

// GetTa1set Timer A output 1 Set
func (r *registerCpt2ccrType) GetTa1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldTa1setMask) != 0
}

// SetTa1set Timer A output 1 Set
func (r *registerCpt2ccrType) SetTa1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldTa1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldTa1setMask)
	}
}

const (
	RegisterCpt2ccrFieldExev10cptShift = 11
	RegisterCpt2ccrFieldExev10cptMask  = 0x800
)

// GetExev10cpt External Event 10 Capture
func (r *registerCpt2ccrType) GetExev10cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldExev10cptMask) != 0
}

// SetExev10cpt External Event 10 Capture
func (r *registerCpt2ccrType) SetExev10cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldExev10cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldExev10cptMask)
	}
}

const (
	RegisterCpt2ccrFieldExev9cptShift = 10
	RegisterCpt2ccrFieldExev9cptMask  = 0x400
)

// GetExev9cpt External Event 9 Capture
func (r *registerCpt2ccrType) GetExev9cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldExev9cptMask) != 0
}

// SetExev9cpt External Event 9 Capture
func (r *registerCpt2ccrType) SetExev9cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldExev9cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldExev9cptMask)
	}
}

const (
	RegisterCpt2ccrFieldExev8cptShift = 9
	RegisterCpt2ccrFieldExev8cptMask  = 0x200
)

// GetExev8cpt External Event 8 Capture
func (r *registerCpt2ccrType) GetExev8cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldExev8cptMask) != 0
}

// SetExev8cpt External Event 8 Capture
func (r *registerCpt2ccrType) SetExev8cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldExev8cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldExev8cptMask)
	}
}

const (
	RegisterCpt2ccrFieldExev7cptShift = 8
	RegisterCpt2ccrFieldExev7cptMask  = 0x100
)

// GetExev7cpt External Event 7 Capture
func (r *registerCpt2ccrType) GetExev7cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldExev7cptMask) != 0
}

// SetExev7cpt External Event 7 Capture
func (r *registerCpt2ccrType) SetExev7cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldExev7cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldExev7cptMask)
	}
}

const (
	RegisterCpt2ccrFieldExev6cptShift = 7
	RegisterCpt2ccrFieldExev6cptMask  = 0x80
)

// GetExev6cpt External Event 6 Capture
func (r *registerCpt2ccrType) GetExev6cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldExev6cptMask) != 0
}

// SetExev6cpt External Event 6 Capture
func (r *registerCpt2ccrType) SetExev6cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldExev6cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldExev6cptMask)
	}
}

const (
	RegisterCpt2ccrFieldExev5cptShift = 6
	RegisterCpt2ccrFieldExev5cptMask  = 0x40
)

// GetExev5cpt External Event 5 Capture
func (r *registerCpt2ccrType) GetExev5cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldExev5cptMask) != 0
}

// SetExev5cpt External Event 5 Capture
func (r *registerCpt2ccrType) SetExev5cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldExev5cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldExev5cptMask)
	}
}

const (
	RegisterCpt2ccrFieldExev4cptShift = 5
	RegisterCpt2ccrFieldExev4cptMask  = 0x20
)

// GetExev4cpt External Event 4 Capture
func (r *registerCpt2ccrType) GetExev4cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldExev4cptMask) != 0
}

// SetExev4cpt External Event 4 Capture
func (r *registerCpt2ccrType) SetExev4cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldExev4cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldExev4cptMask)
	}
}

const (
	RegisterCpt2ccrFieldExev3cptShift = 4
	RegisterCpt2ccrFieldExev3cptMask  = 0x10
)

// GetExev3cpt External Event 3 Capture
func (r *registerCpt2ccrType) GetExev3cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldExev3cptMask) != 0
}

// SetExev3cpt External Event 3 Capture
func (r *registerCpt2ccrType) SetExev3cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldExev3cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldExev3cptMask)
	}
}

const (
	RegisterCpt2ccrFieldExev2cptShift = 3
	RegisterCpt2ccrFieldExev2cptMask  = 0x8
)

// GetExev2cpt External Event 2 Capture
func (r *registerCpt2ccrType) GetExev2cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldExev2cptMask) != 0
}

// SetExev2cpt External Event 2 Capture
func (r *registerCpt2ccrType) SetExev2cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldExev2cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldExev2cptMask)
	}
}

const (
	RegisterCpt2ccrFieldExev1cptShift = 2
	RegisterCpt2ccrFieldExev1cptMask  = 0x4
)

// GetExev1cpt External Event 1 Capture
func (r *registerCpt2ccrType) GetExev1cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldExev1cptMask) != 0
}

// SetExev1cpt External Event 1 Capture
func (r *registerCpt2ccrType) SetExev1cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldExev1cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldExev1cptMask)
	}
}

const (
	RegisterCpt2ccrFieldUdpcptShift = 1
	RegisterCpt2ccrFieldUdpcptMask  = 0x2
)

// GetUdpcpt Update Capture
func (r *registerCpt2ccrType) GetUdpcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldUdpcptMask) != 0
}

// SetUdpcpt Update Capture
func (r *registerCpt2ccrType) SetUdpcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldUdpcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldUdpcptMask)
	}
}

const (
	RegisterCpt2ccrFieldSwcptShift = 0
	RegisterCpt2ccrFieldSwcptMask  = 0x1
)

// GetSwcpt Software Capture
func (r *registerCpt2ccrType) GetSwcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2ccrFieldSwcptMask) != 0
}

// SetSwcpt Software Capture
func (r *registerCpt2ccrType) SetSwcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2ccrFieldSwcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2ccrFieldSwcptMask)
	}
}

// registerOutcrType Timerx Output Register
type registerOutcrType uint32

const (
	RegisterOutcrFieldDidl2Shift = 23
	RegisterOutcrFieldDidl2Mask  = 0x800000
)

// GetDidl2 Output 2 Deadtime upon burst mode Idle entry
func (r *registerOutcrType) GetDidl2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldDidl2Mask) != 0
}

// SetDidl2 Output 2 Deadtime upon burst mode Idle entry
func (r *registerOutcrType) SetDidl2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutcrFieldDidl2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldDidl2Mask)
	}
}

const (
	RegisterOutcrFieldChp2Shift = 22
	RegisterOutcrFieldChp2Mask  = 0x400000
)

// GetChp2 Output 2 Chopper enable
func (r *registerOutcrType) GetChp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldChp2Mask) != 0
}

// SetChp2 Output 2 Chopper enable
func (r *registerOutcrType) SetChp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutcrFieldChp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldChp2Mask)
	}
}

const (
	RegisterOutcrFieldFault2Shift = 20
	RegisterOutcrFieldFault2Mask  = 0x300000
)

// GetFault2 Output 2 Fault state
func (r *registerOutcrType) GetFault2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldFault2Mask) >> RegisterOutcrFieldFault2Shift)
}

// SetFault2 Output 2 Fault state
func (r *registerOutcrType) SetFault2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldFault2Mask)|(uint32(value)<<RegisterOutcrFieldFault2Shift))
}

const (
	RegisterOutcrFieldIdles2Shift = 19
	RegisterOutcrFieldIdles2Mask  = 0x80000
)

// GetIdles2 Output 2 Idle State
func (r *registerOutcrType) GetIdles2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldIdles2Mask) != 0
}

// SetIdles2 Output 2 Idle State
func (r *registerOutcrType) SetIdles2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutcrFieldIdles2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldIdles2Mask)
	}
}

const (
	RegisterOutcrFieldIdlem2Shift = 18
	RegisterOutcrFieldIdlem2Mask  = 0x40000
)

// GetIdlem2 Output 2 Idle mode
func (r *registerOutcrType) GetIdlem2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldIdlem2Mask) != 0
}

// SetIdlem2 Output 2 Idle mode
func (r *registerOutcrType) SetIdlem2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutcrFieldIdlem2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldIdlem2Mask)
	}
}

const (
	RegisterOutcrFieldPol2Shift = 17
	RegisterOutcrFieldPol2Mask  = 0x20000
)

// GetPol2 Output 2 polarity
func (r *registerOutcrType) GetPol2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldPol2Mask) != 0
}

// SetPol2 Output 2 polarity
func (r *registerOutcrType) SetPol2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutcrFieldPol2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldPol2Mask)
	}
}

const (
	RegisterOutcrFieldDlyprtShift = 10
	RegisterOutcrFieldDlyprtMask  = 0x1c00
)

// GetDlyprt Delayed Protection
func (r *registerOutcrType) GetDlyprt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldDlyprtMask) >> RegisterOutcrFieldDlyprtShift)
}

// SetDlyprt Delayed Protection
func (r *registerOutcrType) SetDlyprt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldDlyprtMask)|(uint32(value)<<RegisterOutcrFieldDlyprtShift))
}

const (
	RegisterOutcrFieldDlyprtenShift = 9
	RegisterOutcrFieldDlyprtenMask  = 0x200
)

// GetDlyprten Delayed Protection Enable
func (r *registerOutcrType) GetDlyprten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldDlyprtenMask) != 0
}

// SetDlyprten Delayed Protection Enable
func (r *registerOutcrType) SetDlyprten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutcrFieldDlyprtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldDlyprtenMask)
	}
}

const (
	RegisterOutcrFieldDtenShift = 8
	RegisterOutcrFieldDtenMask  = 0x100
)

// GetDten Deadtime enable
func (r *registerOutcrType) GetDten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldDtenMask) != 0
}

// SetDten Deadtime enable
func (r *registerOutcrType) SetDten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutcrFieldDtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldDtenMask)
	}
}

const (
	RegisterOutcrFieldDidl1Shift = 7
	RegisterOutcrFieldDidl1Mask  = 0x80
)

// GetDidl1 Output 1 Deadtime upon burst mode Idle entry
func (r *registerOutcrType) GetDidl1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldDidl1Mask) != 0
}

// SetDidl1 Output 1 Deadtime upon burst mode Idle entry
func (r *registerOutcrType) SetDidl1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutcrFieldDidl1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldDidl1Mask)
	}
}

const (
	RegisterOutcrFieldChp1Shift = 6
	RegisterOutcrFieldChp1Mask  = 0x40
)

// GetChp1 Output 1 Chopper enable
func (r *registerOutcrType) GetChp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldChp1Mask) != 0
}

// SetChp1 Output 1 Chopper enable
func (r *registerOutcrType) SetChp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutcrFieldChp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldChp1Mask)
	}
}

const (
	RegisterOutcrFieldFault1Shift = 4
	RegisterOutcrFieldFault1Mask  = 0x30
)

// GetFault1 Output 1 Fault state
func (r *registerOutcrType) GetFault1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldFault1Mask) >> RegisterOutcrFieldFault1Shift)
}

// SetFault1 Output 1 Fault state
func (r *registerOutcrType) SetFault1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldFault1Mask)|(uint32(value)<<RegisterOutcrFieldFault1Shift))
}

const (
	RegisterOutcrFieldIdles1Shift = 3
	RegisterOutcrFieldIdles1Mask  = 0x8
)

// GetIdles1 Output 1 Idle State
func (r *registerOutcrType) GetIdles1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldIdles1Mask) != 0
}

// SetIdles1 Output 1 Idle State
func (r *registerOutcrType) SetIdles1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutcrFieldIdles1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldIdles1Mask)
	}
}

const (
	RegisterOutcrFieldIdlem1Shift = 2
	RegisterOutcrFieldIdlem1Mask  = 0x4
)

// GetIdlem1 Output 1 Idle mode
func (r *registerOutcrType) GetIdlem1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldIdlem1Mask) != 0
}

// SetIdlem1 Output 1 Idle mode
func (r *registerOutcrType) SetIdlem1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutcrFieldIdlem1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldIdlem1Mask)
	}
}

const (
	RegisterOutcrFieldPol1Shift = 1
	RegisterOutcrFieldPol1Mask  = 0x2
)

// GetPol1 Output 1 polarity
func (r *registerOutcrType) GetPol1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutcrFieldPol1Mask) != 0
}

// SetPol1 Output 1 polarity
func (r *registerOutcrType) SetPol1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutcrFieldPol1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutcrFieldPol1Mask)
	}
}

// registerFltcrType Timerx Fault Register
type registerFltcrType uint32

const (
	RegisterFltcrFieldFltlckShift = 31
	RegisterFltcrFieldFltlckMask  = 0x80000000
)

// GetFltlck Fault sources Lock
func (r *registerFltcrType) GetFltlck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltcrFieldFltlckMask) != 0
}

// SetFltlck Fault sources Lock
func (r *registerFltcrType) SetFltlck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltcrFieldFltlckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltcrFieldFltlckMask)
	}
}

const (
	RegisterFltcrFieldFlt5enShift = 4
	RegisterFltcrFieldFlt5enMask  = 0x10
)

// GetFlt5en Fault 5 enable
func (r *registerFltcrType) GetFlt5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltcrFieldFlt5enMask) != 0
}

// SetFlt5en Fault 5 enable
func (r *registerFltcrType) SetFlt5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltcrFieldFlt5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltcrFieldFlt5enMask)
	}
}

const (
	RegisterFltcrFieldFlt4enShift = 3
	RegisterFltcrFieldFlt4enMask  = 0x8
)

// GetFlt4en Fault 4 enable
func (r *registerFltcrType) GetFlt4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltcrFieldFlt4enMask) != 0
}

// SetFlt4en Fault 4 enable
func (r *registerFltcrType) SetFlt4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltcrFieldFlt4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltcrFieldFlt4enMask)
	}
}

const (
	RegisterFltcrFieldFlt3enShift = 2
	RegisterFltcrFieldFlt3enMask  = 0x4
)

// GetFlt3en Fault 3 enable
func (r *registerFltcrType) GetFlt3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltcrFieldFlt3enMask) != 0
}

// SetFlt3en Fault 3 enable
func (r *registerFltcrType) SetFlt3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltcrFieldFlt3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltcrFieldFlt3enMask)
	}
}

const (
	RegisterFltcrFieldFlt2enShift = 1
	RegisterFltcrFieldFlt2enMask  = 0x2
)

// GetFlt2en Fault 2 enable
func (r *registerFltcrType) GetFlt2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltcrFieldFlt2enMask) != 0
}

// SetFlt2en Fault 2 enable
func (r *registerFltcrType) SetFlt2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltcrFieldFlt2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltcrFieldFlt2enMask)
	}
}

const (
	RegisterFltcrFieldFlt1enShift = 0
	RegisterFltcrFieldFlt1enMask  = 0x1
)

// GetFlt1en Fault 1 enable
func (r *registerFltcrType) GetFlt1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltcrFieldFlt1enMask) != 0
}

// SetFlt1en Fault 1 enable
func (r *registerFltcrType) SetFlt1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltcrFieldFlt1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltcrFieldFlt1enMask)
	}
}
