package hrtim_tima

import (
	"unsafe"
	"volatile"
)

var (
	Hrtim_tima = (*_hrtim_tima)(unsafe.Pointer(uintptr(0x40017480)))
)

type _hrtim_tima struct {
	Timacr    registerTimacrType
	Timaisr   registerTimaisrType
	Timaicr   registerTimaicrType
	Timadier5 registerTimadier5Type
	Cntar     registerCntarType
	Perar     registerPerarType
	Repar     registerReparType
	Cmp1ar    registerCmp1arType
	Cmp1car   registerCmp1carType
	Cmp2ar    registerCmp2arType
	Cmp3ar    registerCmp3arType
	Cmp4ar    registerCmp4arType
	Cpt1ar    registerCpt1arType
	Cpt2ar    registerCpt2arType
	Dtar      registerDtarType
	Seta1r    registerSeta1rType
	Rsta1r    registerRsta1rType
	Seta2r    registerSeta2rType
	Rsta2r    registerRsta2rType
	Eefar1    registerEefar1Type
	Eefar2    registerEefar2Type
	Rstar     registerRstarType
	Chpar     registerChparType
	Cpt1acr   registerCpt1acrType
	Cpt2acr   registerCpt2acrType
	Outar     registerOutarType
	Fltar     registerFltarType
}

// registerTimacrType Timerx Control Register
type registerTimacrType uint32

const (
	RegisterTimacrFieldUpdgatShift = 28
	RegisterTimacrFieldUpdgatMask  = 0xf0000000
)

// GetUpdgat Update Gating
func (r *registerTimacrType) GetUpdgat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldUpdgatMask) >> RegisterTimacrFieldUpdgatShift)
}

// SetUpdgat Update Gating
func (r *registerTimacrType) SetUpdgat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldUpdgatMask)|(uint32(value)<<RegisterTimacrFieldUpdgatShift))
}

const (
	RegisterTimacrFieldPreenShift = 27
	RegisterTimacrFieldPreenMask  = 0x8000000
)

// GetPreen Preload enable
func (r *registerTimacrType) GetPreen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldPreenMask) != 0
}

// SetPreen Preload enable
func (r *registerTimacrType) SetPreen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldPreenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldPreenMask)
	}
}

const (
	RegisterTimacrFieldDacsyncShift = 25
	RegisterTimacrFieldDacsyncMask  = 0x6000000
)

// GetDacsync AC Synchronization
func (r *registerTimacrType) GetDacsync() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldDacsyncMask) >> RegisterTimacrFieldDacsyncShift)
}

// SetDacsync AC Synchronization
func (r *registerTimacrType) SetDacsync(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldDacsyncMask)|(uint32(value)<<RegisterTimacrFieldDacsyncShift))
}

const (
	RegisterTimacrFieldMstuShift = 24
	RegisterTimacrFieldMstuMask  = 0x1000000
)

// GetMstu Master Timer update
func (r *registerTimacrType) GetMstu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldMstuMask) != 0
}

// SetMstu Master Timer update
func (r *registerTimacrType) SetMstu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldMstuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldMstuMask)
	}
}

const (
	RegisterTimacrFieldTeuShift = 23
	RegisterTimacrFieldTeuMask  = 0x800000
)

// GetTeu TEU
func (r *registerTimacrType) GetTeu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldTeuMask) != 0
}

// SetTeu TEU
func (r *registerTimacrType) SetTeu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldTeuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldTeuMask)
	}
}

const (
	RegisterTimacrFieldTduShift = 22
	RegisterTimacrFieldTduMask  = 0x400000
)

// GetTdu TDU
func (r *registerTimacrType) GetTdu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldTduMask) != 0
}

// SetTdu TDU
func (r *registerTimacrType) SetTdu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldTduMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldTduMask)
	}
}

const (
	RegisterTimacrFieldTcuShift = 21
	RegisterTimacrFieldTcuMask  = 0x200000
)

// GetTcu TCU
func (r *registerTimacrType) GetTcu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldTcuMask) != 0
}

// SetTcu TCU
func (r *registerTimacrType) SetTcu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldTcuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldTcuMask)
	}
}

const (
	RegisterTimacrFieldTbuShift = 20
	RegisterTimacrFieldTbuMask  = 0x100000
)

// GetTbu TBU
func (r *registerTimacrType) GetTbu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldTbuMask) != 0
}

// SetTbu TBU
func (r *registerTimacrType) SetTbu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldTbuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldTbuMask)
	}
}

const (
	RegisterTimacrFieldTxrstuShift = 18
	RegisterTimacrFieldTxrstuMask  = 0x40000
)

// GetTxrstu Timerx reset update
func (r *registerTimacrType) GetTxrstu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldTxrstuMask) != 0
}

// SetTxrstu Timerx reset update
func (r *registerTimacrType) SetTxrstu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldTxrstuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldTxrstuMask)
	}
}

const (
	RegisterTimacrFieldTxrepuShift = 17
	RegisterTimacrFieldTxrepuMask  = 0x20000
)

// GetTxrepu Timer x Repetition update
func (r *registerTimacrType) GetTxrepu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldTxrepuMask) != 0
}

// SetTxrepu Timer x Repetition update
func (r *registerTimacrType) SetTxrepu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldTxrepuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldTxrepuMask)
	}
}

const (
	RegisterTimacrFieldDelcmp4Shift = 14
	RegisterTimacrFieldDelcmp4Mask  = 0xc000
)

// GetDelcmp4 Delayed CMP4 mode
func (r *registerTimacrType) GetDelcmp4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldDelcmp4Mask) >> RegisterTimacrFieldDelcmp4Shift)
}

// SetDelcmp4 Delayed CMP4 mode
func (r *registerTimacrType) SetDelcmp4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldDelcmp4Mask)|(uint32(value)<<RegisterTimacrFieldDelcmp4Shift))
}

const (
	RegisterTimacrFieldDelcmp2Shift = 12
	RegisterTimacrFieldDelcmp2Mask  = 0x3000
)

// GetDelcmp2 Delayed CMP2 mode
func (r *registerTimacrType) GetDelcmp2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldDelcmp2Mask) >> RegisterTimacrFieldDelcmp2Shift)
}

// SetDelcmp2 Delayed CMP2 mode
func (r *registerTimacrType) SetDelcmp2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldDelcmp2Mask)|(uint32(value)<<RegisterTimacrFieldDelcmp2Shift))
}

const (
	RegisterTimacrFieldSyncstrtxShift = 11
	RegisterTimacrFieldSyncstrtxMask  = 0x800
)

// GetSyncstrtx Synchronization Starts Timer x
func (r *registerTimacrType) GetSyncstrtx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldSyncstrtxMask) != 0
}

// SetSyncstrtx Synchronization Starts Timer x
func (r *registerTimacrType) SetSyncstrtx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldSyncstrtxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldSyncstrtxMask)
	}
}

const (
	RegisterTimacrFieldSyncrstxShift = 10
	RegisterTimacrFieldSyncrstxMask  = 0x400
)

// GetSyncrstx Synchronization Resets Timer x
func (r *registerTimacrType) GetSyncrstx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldSyncrstxMask) != 0
}

// SetSyncrstx Synchronization Resets Timer x
func (r *registerTimacrType) SetSyncrstx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldSyncrstxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldSyncrstxMask)
	}
}

const (
	RegisterTimacrFieldPshpllShift = 6
	RegisterTimacrFieldPshpllMask  = 0x40
)

// GetPshpll Push-Pull mode enable
func (r *registerTimacrType) GetPshpll() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldPshpllMask) != 0
}

// SetPshpll Push-Pull mode enable
func (r *registerTimacrType) SetPshpll(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldPshpllMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldPshpllMask)
	}
}

const (
	RegisterTimacrFieldHalfShift = 5
	RegisterTimacrFieldHalfMask  = 0x20
)

// GetHalf Half mode enable
func (r *registerTimacrType) GetHalf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldHalfMask) != 0
}

// SetHalf Half mode enable
func (r *registerTimacrType) SetHalf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldHalfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldHalfMask)
	}
}

const (
	RegisterTimacrFieldRetrigShift = 4
	RegisterTimacrFieldRetrigMask  = 0x10
)

// GetRetrig Re-triggerable mode
func (r *registerTimacrType) GetRetrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldRetrigMask) != 0
}

// SetRetrig Re-triggerable mode
func (r *registerTimacrType) SetRetrig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldRetrigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldRetrigMask)
	}
}

const (
	RegisterTimacrFieldContShift = 3
	RegisterTimacrFieldContMask  = 0x8
)

// GetCont Continuous mode
func (r *registerTimacrType) GetCont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldContMask) != 0
}

// SetCont Continuous mode
func (r *registerTimacrType) SetCont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimacrFieldContMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldContMask)
	}
}

const (
	RegisterTimacrFieldCk_pscxShift = 0
	RegisterTimacrFieldCk_pscxMask  = 0x7
)

// GetCk_pscx HRTIM Timer x Clock prescaler
func (r *registerTimacrType) GetCk_pscx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimacrFieldCk_pscxMask) >> RegisterTimacrFieldCk_pscxShift)
}

// SetCk_pscx HRTIM Timer x Clock prescaler
func (r *registerTimacrType) SetCk_pscx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimacrFieldCk_pscxMask)|(uint32(value)<<RegisterTimacrFieldCk_pscxShift))
}

// registerTimaisrType Timerx Interrupt Status Register
type registerTimaisrType uint32

const (
	RegisterTimaisrFieldO2statShift = 19
	RegisterTimaisrFieldO2statMask  = 0x80000
)

// GetO2stat Output 2 State
func (r *registerTimaisrType) GetO2stat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldO2statMask) != 0
}

// SetO2stat Output 2 State
func (r *registerTimaisrType) SetO2stat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldO2statMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldO2statMask)
	}
}

const (
	RegisterTimaisrFieldO1statShift = 18
	RegisterTimaisrFieldO1statMask  = 0x40000
)

// GetO1stat Output 1 State
func (r *registerTimaisrType) GetO1stat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldO1statMask) != 0
}

// SetO1stat Output 1 State
func (r *registerTimaisrType) SetO1stat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldO1statMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldO1statMask)
	}
}

const (
	RegisterTimaisrFieldIppstatShift = 17
	RegisterTimaisrFieldIppstatMask  = 0x20000
)

// GetIppstat Idle Push Pull Status
func (r *registerTimaisrType) GetIppstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldIppstatMask) != 0
}

// SetIppstat Idle Push Pull Status
func (r *registerTimaisrType) SetIppstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldIppstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldIppstatMask)
	}
}

const (
	RegisterTimaisrFieldCppstatShift = 16
	RegisterTimaisrFieldCppstatMask  = 0x10000
)

// GetCppstat Current Push Pull Status
func (r *registerTimaisrType) GetCppstat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldCppstatMask) != 0
}

// SetCppstat Current Push Pull Status
func (r *registerTimaisrType) SetCppstat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldCppstatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldCppstatMask)
	}
}

const (
	RegisterTimaisrFieldDlyprtShift = 14
	RegisterTimaisrFieldDlyprtMask  = 0x4000
)

// GetDlyprt Delayed Protection Flag
func (r *registerTimaisrType) GetDlyprt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldDlyprtMask) != 0
}

// SetDlyprt Delayed Protection Flag
func (r *registerTimaisrType) SetDlyprt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldDlyprtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldDlyprtMask)
	}
}

const (
	RegisterTimaisrFieldRstShift = 13
	RegisterTimaisrFieldRstMask  = 0x2000
)

// GetRst Reset Interrupt Flag
func (r *registerTimaisrType) GetRst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldRstMask) != 0
}

// SetRst Reset Interrupt Flag
func (r *registerTimaisrType) SetRst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldRstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldRstMask)
	}
}

const (
	RegisterTimaisrFieldRstx2Shift = 12
	RegisterTimaisrFieldRstx2Mask  = 0x1000
)

// GetRstx2 Output 2 Reset Interrupt Flag
func (r *registerTimaisrType) GetRstx2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldRstx2Mask) != 0
}

// SetRstx2 Output 2 Reset Interrupt Flag
func (r *registerTimaisrType) SetRstx2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldRstx2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldRstx2Mask)
	}
}

const (
	RegisterTimaisrFieldSetx2Shift = 11
	RegisterTimaisrFieldSetx2Mask  = 0x800
)

// GetSetx2 Output 2 Set Interrupt Flag
func (r *registerTimaisrType) GetSetx2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldSetx2Mask) != 0
}

// SetSetx2 Output 2 Set Interrupt Flag
func (r *registerTimaisrType) SetSetx2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldSetx2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldSetx2Mask)
	}
}

const (
	RegisterTimaisrFieldRstx1Shift = 10
	RegisterTimaisrFieldRstx1Mask  = 0x400
)

// GetRstx1 Output 1 Reset Interrupt Flag
func (r *registerTimaisrType) GetRstx1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldRstx1Mask) != 0
}

// SetRstx1 Output 1 Reset Interrupt Flag
func (r *registerTimaisrType) SetRstx1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldRstx1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldRstx1Mask)
	}
}

const (
	RegisterTimaisrFieldSetx1Shift = 9
	RegisterTimaisrFieldSetx1Mask  = 0x200
)

// GetSetx1 Output 1 Set Interrupt Flag
func (r *registerTimaisrType) GetSetx1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldSetx1Mask) != 0
}

// SetSetx1 Output 1 Set Interrupt Flag
func (r *registerTimaisrType) SetSetx1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldSetx1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldSetx1Mask)
	}
}

const (
	RegisterTimaisrFieldCpt2Shift = 8
	RegisterTimaisrFieldCpt2Mask  = 0x100
)

// GetCpt2 Capture2 Interrupt Flag
func (r *registerTimaisrType) GetCpt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldCpt2Mask) != 0
}

// SetCpt2 Capture2 Interrupt Flag
func (r *registerTimaisrType) SetCpt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldCpt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldCpt2Mask)
	}
}

const (
	RegisterTimaisrFieldCpt1Shift = 7
	RegisterTimaisrFieldCpt1Mask  = 0x80
)

// GetCpt1 Capture1 Interrupt Flag
func (r *registerTimaisrType) GetCpt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldCpt1Mask) != 0
}

// SetCpt1 Capture1 Interrupt Flag
func (r *registerTimaisrType) SetCpt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldCpt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldCpt1Mask)
	}
}

const (
	RegisterTimaisrFieldUpdShift = 6
	RegisterTimaisrFieldUpdMask  = 0x40
)

// GetUpd Update Interrupt Flag
func (r *registerTimaisrType) GetUpd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldUpdMask) != 0
}

// SetUpd Update Interrupt Flag
func (r *registerTimaisrType) SetUpd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldUpdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldUpdMask)
	}
}

const (
	RegisterTimaisrFieldRepShift = 4
	RegisterTimaisrFieldRepMask  = 0x10
)

// GetRep Repetition Interrupt Flag
func (r *registerTimaisrType) GetRep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldRepMask) != 0
}

// SetRep Repetition Interrupt Flag
func (r *registerTimaisrType) SetRep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldRepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldRepMask)
	}
}

const (
	RegisterTimaisrFieldCmp4Shift = 3
	RegisterTimaisrFieldCmp4Mask  = 0x8
)

// GetCmp4 Compare 4 Interrupt Flag
func (r *registerTimaisrType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldCmp4Mask) != 0
}

// SetCmp4 Compare 4 Interrupt Flag
func (r *registerTimaisrType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldCmp4Mask)
	}
}

const (
	RegisterTimaisrFieldCmp3Shift = 2
	RegisterTimaisrFieldCmp3Mask  = 0x4
)

// GetCmp3 Compare 3 Interrupt Flag
func (r *registerTimaisrType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldCmp3Mask) != 0
}

// SetCmp3 Compare 3 Interrupt Flag
func (r *registerTimaisrType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldCmp3Mask)
	}
}

const (
	RegisterTimaisrFieldCmp2Shift = 1
	RegisterTimaisrFieldCmp2Mask  = 0x2
)

// GetCmp2 Compare 2 Interrupt Flag
func (r *registerTimaisrType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldCmp2Mask) != 0
}

// SetCmp2 Compare 2 Interrupt Flag
func (r *registerTimaisrType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldCmp2Mask)
	}
}

const (
	RegisterTimaisrFieldCmp1Shift = 0
	RegisterTimaisrFieldCmp1Mask  = 0x1
)

// GetCmp1 Compare 1 Interrupt Flag
func (r *registerTimaisrType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaisrFieldCmp1Mask) != 0
}

// SetCmp1 Compare 1 Interrupt Flag
func (r *registerTimaisrType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaisrFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaisrFieldCmp1Mask)
	}
}

// registerTimaicrType Timerx Interrupt Clear Register
type registerTimaicrType uint32

const (
	RegisterTimaicrFieldDlyprtcShift = 14
	RegisterTimaicrFieldDlyprtcMask  = 0x4000
)

// GetDlyprtc Delayed Protection Flag Clear
func (r *registerTimaicrType) GetDlyprtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldDlyprtcMask) != 0
}

// SetDlyprtc Delayed Protection Flag Clear
func (r *registerTimaicrType) SetDlyprtc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldDlyprtcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldDlyprtcMask)
	}
}

const (
	RegisterTimaicrFieldRstcShift = 13
	RegisterTimaicrFieldRstcMask  = 0x2000
)

// GetRstc Reset Interrupt flag Clear
func (r *registerTimaicrType) GetRstc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldRstcMask) != 0
}

// SetRstc Reset Interrupt flag Clear
func (r *registerTimaicrType) SetRstc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldRstcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldRstcMask)
	}
}

const (
	RegisterTimaicrFieldRstx2cShift = 12
	RegisterTimaicrFieldRstx2cMask  = 0x1000
)

// GetRstx2c Output 2 Reset flag Clear
func (r *registerTimaicrType) GetRstx2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldRstx2cMask) != 0
}

// SetRstx2c Output 2 Reset flag Clear
func (r *registerTimaicrType) SetRstx2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldRstx2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldRstx2cMask)
	}
}

const (
	RegisterTimaicrFieldSet2xcShift = 11
	RegisterTimaicrFieldSet2xcMask  = 0x800
)

// GetSet2xc Output 2 Set flag Clear
func (r *registerTimaicrType) GetSet2xc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldSet2xcMask) != 0
}

// SetSet2xc Output 2 Set flag Clear
func (r *registerTimaicrType) SetSet2xc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldSet2xcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldSet2xcMask)
	}
}

const (
	RegisterTimaicrFieldRstx1cShift = 10
	RegisterTimaicrFieldRstx1cMask  = 0x400
)

// GetRstx1c Output 1 Reset flag Clear
func (r *registerTimaicrType) GetRstx1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldRstx1cMask) != 0
}

// SetRstx1c Output 1 Reset flag Clear
func (r *registerTimaicrType) SetRstx1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldRstx1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldRstx1cMask)
	}
}

const (
	RegisterTimaicrFieldSet1xcShift = 9
	RegisterTimaicrFieldSet1xcMask  = 0x200
)

// GetSet1xc Output 1 Set flag Clear
func (r *registerTimaicrType) GetSet1xc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldSet1xcMask) != 0
}

// SetSet1xc Output 1 Set flag Clear
func (r *registerTimaicrType) SetSet1xc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldSet1xcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldSet1xcMask)
	}
}

const (
	RegisterTimaicrFieldCpt2cShift = 8
	RegisterTimaicrFieldCpt2cMask  = 0x100
)

// GetCpt2c Capture2 Interrupt flag Clear
func (r *registerTimaicrType) GetCpt2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldCpt2cMask) != 0
}

// SetCpt2c Capture2 Interrupt flag Clear
func (r *registerTimaicrType) SetCpt2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldCpt2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldCpt2cMask)
	}
}

const (
	RegisterTimaicrFieldCpt1cShift = 7
	RegisterTimaicrFieldCpt1cMask  = 0x80
)

// GetCpt1c Capture1 Interrupt flag Clear
func (r *registerTimaicrType) GetCpt1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldCpt1cMask) != 0
}

// SetCpt1c Capture1 Interrupt flag Clear
func (r *registerTimaicrType) SetCpt1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldCpt1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldCpt1cMask)
	}
}

const (
	RegisterTimaicrFieldUpdcShift = 6
	RegisterTimaicrFieldUpdcMask  = 0x40
)

// GetUpdc Update Interrupt flag Clear
func (r *registerTimaicrType) GetUpdc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldUpdcMask) != 0
}

// SetUpdc Update Interrupt flag Clear
func (r *registerTimaicrType) SetUpdc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldUpdcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldUpdcMask)
	}
}

const (
	RegisterTimaicrFieldRepcShift = 4
	RegisterTimaicrFieldRepcMask  = 0x10
)

// GetRepc Repetition Interrupt flag Clear
func (r *registerTimaicrType) GetRepc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldRepcMask) != 0
}

// SetRepc Repetition Interrupt flag Clear
func (r *registerTimaicrType) SetRepc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldRepcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldRepcMask)
	}
}

const (
	RegisterTimaicrFieldCmp4cShift = 3
	RegisterTimaicrFieldCmp4cMask  = 0x8
)

// GetCmp4c Compare 4 Interrupt flag Clear
func (r *registerTimaicrType) GetCmp4c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldCmp4cMask) != 0
}

// SetCmp4c Compare 4 Interrupt flag Clear
func (r *registerTimaicrType) SetCmp4c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldCmp4cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldCmp4cMask)
	}
}

const (
	RegisterTimaicrFieldCmp3cShift = 2
	RegisterTimaicrFieldCmp3cMask  = 0x4
)

// GetCmp3c Compare 3 Interrupt flag Clear
func (r *registerTimaicrType) GetCmp3c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldCmp3cMask) != 0
}

// SetCmp3c Compare 3 Interrupt flag Clear
func (r *registerTimaicrType) SetCmp3c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldCmp3cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldCmp3cMask)
	}
}

const (
	RegisterTimaicrFieldCmp2cShift = 1
	RegisterTimaicrFieldCmp2cMask  = 0x2
)

// GetCmp2c Compare 2 Interrupt flag Clear
func (r *registerTimaicrType) GetCmp2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldCmp2cMask) != 0
}

// SetCmp2c Compare 2 Interrupt flag Clear
func (r *registerTimaicrType) SetCmp2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldCmp2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldCmp2cMask)
	}
}

const (
	RegisterTimaicrFieldCmp1cShift = 0
	RegisterTimaicrFieldCmp1cMask  = 0x1
)

// GetCmp1c Compare 1 Interrupt flag Clear
func (r *registerTimaicrType) GetCmp1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimaicrFieldCmp1cMask) != 0
}

// SetCmp1c Compare 1 Interrupt flag Clear
func (r *registerTimaicrType) SetCmp1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimaicrFieldCmp1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimaicrFieldCmp1cMask)
	}
}

// registerTimadier5Type TIMxDIER5
type registerTimadier5Type uint32

const (
	RegisterTimadier5FieldDlyprtdeShift = 30
	RegisterTimadier5FieldDlyprtdeMask  = 0x40000000
)

// GetDlyprtde DLYPRTDE
func (r *registerTimadier5Type) GetDlyprtde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldDlyprtdeMask) != 0
}

// SetDlyprtde DLYPRTDE
func (r *registerTimadier5Type) SetDlyprtde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldDlyprtdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldDlyprtdeMask)
	}
}

const (
	RegisterTimadier5FieldRstdeShift = 29
	RegisterTimadier5FieldRstdeMask  = 0x20000000
)

// GetRstde RSTDE
func (r *registerTimadier5Type) GetRstde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldRstdeMask) != 0
}

// SetRstde RSTDE
func (r *registerTimadier5Type) SetRstde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldRstdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldRstdeMask)
	}
}

const (
	RegisterTimadier5FieldRstx2deShift = 28
	RegisterTimadier5FieldRstx2deMask  = 0x10000000
)

// GetRstx2de RSTx2DE
func (r *registerTimadier5Type) GetRstx2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldRstx2deMask) != 0
}

// SetRstx2de RSTx2DE
func (r *registerTimadier5Type) SetRstx2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldRstx2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldRstx2deMask)
	}
}

const (
	RegisterTimadier5FieldSetx2deShift = 27
	RegisterTimadier5FieldSetx2deMask  = 0x8000000
)

// GetSetx2de SETx2DE
func (r *registerTimadier5Type) GetSetx2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldSetx2deMask) != 0
}

// SetSetx2de SETx2DE
func (r *registerTimadier5Type) SetSetx2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldSetx2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldSetx2deMask)
	}
}

const (
	RegisterTimadier5FieldRstx1deShift = 26
	RegisterTimadier5FieldRstx1deMask  = 0x4000000
)

// GetRstx1de RSTx1DE
func (r *registerTimadier5Type) GetRstx1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldRstx1deMask) != 0
}

// SetRstx1de RSTx1DE
func (r *registerTimadier5Type) SetRstx1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldRstx1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldRstx1deMask)
	}
}

const (
	RegisterTimadier5FieldSet1xdeShift = 25
	RegisterTimadier5FieldSet1xdeMask  = 0x2000000
)

// GetSet1xde SET1xDE
func (r *registerTimadier5Type) GetSet1xde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldSet1xdeMask) != 0
}

// SetSet1xde SET1xDE
func (r *registerTimadier5Type) SetSet1xde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldSet1xdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldSet1xdeMask)
	}
}

const (
	RegisterTimadier5FieldCpt2deShift = 24
	RegisterTimadier5FieldCpt2deMask  = 0x1000000
)

// GetCpt2de CPT2DE
func (r *registerTimadier5Type) GetCpt2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldCpt2deMask) != 0
}

// SetCpt2de CPT2DE
func (r *registerTimadier5Type) SetCpt2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldCpt2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldCpt2deMask)
	}
}

const (
	RegisterTimadier5FieldCpt1deShift = 23
	RegisterTimadier5FieldCpt1deMask  = 0x800000
)

// GetCpt1de CPT1DE
func (r *registerTimadier5Type) GetCpt1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldCpt1deMask) != 0
}

// SetCpt1de CPT1DE
func (r *registerTimadier5Type) SetCpt1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldCpt1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldCpt1deMask)
	}
}

const (
	RegisterTimadier5FieldUpddeShift = 22
	RegisterTimadier5FieldUpddeMask  = 0x400000
)

// GetUpdde UPDDE
func (r *registerTimadier5Type) GetUpdde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldUpddeMask) != 0
}

// SetUpdde UPDDE
func (r *registerTimadier5Type) SetUpdde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldUpddeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldUpddeMask)
	}
}

const (
	RegisterTimadier5FieldRepdeShift = 20
	RegisterTimadier5FieldRepdeMask  = 0x100000
)

// GetRepde REPDE
func (r *registerTimadier5Type) GetRepde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldRepdeMask) != 0
}

// SetRepde REPDE
func (r *registerTimadier5Type) SetRepde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldRepdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldRepdeMask)
	}
}

const (
	RegisterTimadier5FieldCmp4deShift = 19
	RegisterTimadier5FieldCmp4deMask  = 0x80000
)

// GetCmp4de CMP4DE
func (r *registerTimadier5Type) GetCmp4de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldCmp4deMask) != 0
}

// SetCmp4de CMP4DE
func (r *registerTimadier5Type) SetCmp4de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldCmp4deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldCmp4deMask)
	}
}

const (
	RegisterTimadier5FieldCmp3deShift = 18
	RegisterTimadier5FieldCmp3deMask  = 0x40000
)

// GetCmp3de CMP3DE
func (r *registerTimadier5Type) GetCmp3de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldCmp3deMask) != 0
}

// SetCmp3de CMP3DE
func (r *registerTimadier5Type) SetCmp3de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldCmp3deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldCmp3deMask)
	}
}

const (
	RegisterTimadier5FieldCmp2deShift = 17
	RegisterTimadier5FieldCmp2deMask  = 0x20000
)

// GetCmp2de CMP2DE
func (r *registerTimadier5Type) GetCmp2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldCmp2deMask) != 0
}

// SetCmp2de CMP2DE
func (r *registerTimadier5Type) SetCmp2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldCmp2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldCmp2deMask)
	}
}

const (
	RegisterTimadier5FieldCmp1deShift = 16
	RegisterTimadier5FieldCmp1deMask  = 0x10000
)

// GetCmp1de CMP1DE
func (r *registerTimadier5Type) GetCmp1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldCmp1deMask) != 0
}

// SetCmp1de CMP1DE
func (r *registerTimadier5Type) SetCmp1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldCmp1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldCmp1deMask)
	}
}

const (
	RegisterTimadier5FieldDlyprtieShift = 14
	RegisterTimadier5FieldDlyprtieMask  = 0x4000
)

// GetDlyprtie DLYPRTIE
func (r *registerTimadier5Type) GetDlyprtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldDlyprtieMask) != 0
}

// SetDlyprtie DLYPRTIE
func (r *registerTimadier5Type) SetDlyprtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldDlyprtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldDlyprtieMask)
	}
}

const (
	RegisterTimadier5FieldRstieShift = 13
	RegisterTimadier5FieldRstieMask  = 0x2000
)

// GetRstie RSTIE
func (r *registerTimadier5Type) GetRstie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldRstieMask) != 0
}

// SetRstie RSTIE
func (r *registerTimadier5Type) SetRstie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldRstieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldRstieMask)
	}
}

const (
	RegisterTimadier5FieldRstx2ieShift = 12
	RegisterTimadier5FieldRstx2ieMask  = 0x1000
)

// GetRstx2ie RSTx2IE
func (r *registerTimadier5Type) GetRstx2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldRstx2ieMask) != 0
}

// SetRstx2ie RSTx2IE
func (r *registerTimadier5Type) SetRstx2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldRstx2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldRstx2ieMask)
	}
}

const (
	RegisterTimadier5FieldSetx2ieShift = 11
	RegisterTimadier5FieldSetx2ieMask  = 0x800
)

// GetSetx2ie SETx2IE
func (r *registerTimadier5Type) GetSetx2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldSetx2ieMask) != 0
}

// SetSetx2ie SETx2IE
func (r *registerTimadier5Type) SetSetx2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldSetx2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldSetx2ieMask)
	}
}

const (
	RegisterTimadier5FieldRstx1ieShift = 10
	RegisterTimadier5FieldRstx1ieMask  = 0x400
)

// GetRstx1ie RSTx1IE
func (r *registerTimadier5Type) GetRstx1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldRstx1ieMask) != 0
}

// SetRstx1ie RSTx1IE
func (r *registerTimadier5Type) SetRstx1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldRstx1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldRstx1ieMask)
	}
}

const (
	RegisterTimadier5FieldSet1xieShift = 9
	RegisterTimadier5FieldSet1xieMask  = 0x200
)

// GetSet1xie SET1xIE
func (r *registerTimadier5Type) GetSet1xie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldSet1xieMask) != 0
}

// SetSet1xie SET1xIE
func (r *registerTimadier5Type) SetSet1xie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldSet1xieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldSet1xieMask)
	}
}

const (
	RegisterTimadier5FieldCpt2ieShift = 8
	RegisterTimadier5FieldCpt2ieMask  = 0x100
)

// GetCpt2ie CPT2IE
func (r *registerTimadier5Type) GetCpt2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldCpt2ieMask) != 0
}

// SetCpt2ie CPT2IE
func (r *registerTimadier5Type) SetCpt2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldCpt2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldCpt2ieMask)
	}
}

const (
	RegisterTimadier5FieldCpt1ieShift = 7
	RegisterTimadier5FieldCpt1ieMask  = 0x80
)

// GetCpt1ie CPT1IE
func (r *registerTimadier5Type) GetCpt1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldCpt1ieMask) != 0
}

// SetCpt1ie CPT1IE
func (r *registerTimadier5Type) SetCpt1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldCpt1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldCpt1ieMask)
	}
}

const (
	RegisterTimadier5FieldUpdieShift = 6
	RegisterTimadier5FieldUpdieMask  = 0x40
)

// GetUpdie UPDIE
func (r *registerTimadier5Type) GetUpdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldUpdieMask) != 0
}

// SetUpdie UPDIE
func (r *registerTimadier5Type) SetUpdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldUpdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldUpdieMask)
	}
}

const (
	RegisterTimadier5FieldRepieShift = 4
	RegisterTimadier5FieldRepieMask  = 0x10
)

// GetRepie REPIE
func (r *registerTimadier5Type) GetRepie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldRepieMask) != 0
}

// SetRepie REPIE
func (r *registerTimadier5Type) SetRepie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldRepieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldRepieMask)
	}
}

const (
	RegisterTimadier5FieldCmp4ieShift = 3
	RegisterTimadier5FieldCmp4ieMask  = 0x8
)

// GetCmp4ie CMP4IE
func (r *registerTimadier5Type) GetCmp4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldCmp4ieMask) != 0
}

// SetCmp4ie CMP4IE
func (r *registerTimadier5Type) SetCmp4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldCmp4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldCmp4ieMask)
	}
}

const (
	RegisterTimadier5FieldCmp3ieShift = 2
	RegisterTimadier5FieldCmp3ieMask  = 0x4
)

// GetCmp3ie CMP3IE
func (r *registerTimadier5Type) GetCmp3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldCmp3ieMask) != 0
}

// SetCmp3ie CMP3IE
func (r *registerTimadier5Type) SetCmp3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldCmp3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldCmp3ieMask)
	}
}

const (
	RegisterTimadier5FieldCmp2ieShift = 1
	RegisterTimadier5FieldCmp2ieMask  = 0x2
)

// GetCmp2ie CMP2IE
func (r *registerTimadier5Type) GetCmp2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldCmp2ieMask) != 0
}

// SetCmp2ie CMP2IE
func (r *registerTimadier5Type) SetCmp2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldCmp2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldCmp2ieMask)
	}
}

const (
	RegisterTimadier5FieldCmp1ieShift = 0
	RegisterTimadier5FieldCmp1ieMask  = 0x1
)

// GetCmp1ie CMP1IE
func (r *registerTimadier5Type) GetCmp1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimadier5FieldCmp1ieMask) != 0
}

// SetCmp1ie CMP1IE
func (r *registerTimadier5Type) SetCmp1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimadier5FieldCmp1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimadier5FieldCmp1ieMask)
	}
}

// registerCntarType Timerx Counter Register
type registerCntarType uint32

const (
	RegisterCntarFieldCntxShift = 0
	RegisterCntarFieldCntxMask  = 0xffff
)

// GetCntx Timerx Counter value
func (r *registerCntarType) GetCntx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntarFieldCntxMask) >> RegisterCntarFieldCntxShift)
}

// SetCntx Timerx Counter value
func (r *registerCntarType) SetCntx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntarFieldCntxMask)|(uint32(value)<<RegisterCntarFieldCntxShift))
}

// registerPerarType Timerx Period Register
type registerPerarType uint32

const (
	RegisterPerarFieldPerxShift = 0
	RegisterPerarFieldPerxMask  = 0xffff
)

// GetPerx Timerx Period value
func (r *registerPerarType) GetPerx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPerarFieldPerxMask) >> RegisterPerarFieldPerxShift)
}

// SetPerx Timerx Period value
func (r *registerPerarType) SetPerx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPerarFieldPerxMask)|(uint32(value)<<RegisterPerarFieldPerxShift))
}

// registerReparType Timerx Repetition Register
type registerReparType uint32

const (
	RegisterReparFieldRepxShift = 0
	RegisterReparFieldRepxMask  = 0xff
)

// GetRepx Timerx Repetition counter value
func (r *registerReparType) GetRepx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterReparFieldRepxMask) >> RegisterReparFieldRepxShift)
}

// SetRepx Timerx Repetition counter value
func (r *registerReparType) SetRepx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterReparFieldRepxMask)|(uint32(value)<<RegisterReparFieldRepxShift))
}

// registerCmp1arType Timerx Compare 1 Register
type registerCmp1arType uint32

const (
	RegisterCmp1arFieldCmp1xShift = 0
	RegisterCmp1arFieldCmp1xMask  = 0xffff
)

// GetCmp1x Timerx Compare 1 value
func (r *registerCmp1arType) GetCmp1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1arFieldCmp1xMask) >> RegisterCmp1arFieldCmp1xShift)
}

// SetCmp1x Timerx Compare 1 value
func (r *registerCmp1arType) SetCmp1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1arFieldCmp1xMask)|(uint32(value)<<RegisterCmp1arFieldCmp1xShift))
}

// registerCmp1carType Timerx Compare 1 Compound Register
type registerCmp1carType uint32

const (
	RegisterCmp1carFieldRepxShift = 16
	RegisterCmp1carFieldRepxMask  = 0xff0000
)

// GetRepx Timerx Repetition value (aliased from HRTIM_REPx register)
func (r *registerCmp1carType) GetRepx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1carFieldRepxMask) >> RegisterCmp1carFieldRepxShift)
}

// SetRepx Timerx Repetition value (aliased from HRTIM_REPx register)
func (r *registerCmp1carType) SetRepx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1carFieldRepxMask)|(uint32(value)<<RegisterCmp1carFieldRepxShift))
}

const (
	RegisterCmp1carFieldCmp1xShift = 0
	RegisterCmp1carFieldCmp1xMask  = 0xffff
)

// GetCmp1x Timerx Compare 1 value
func (r *registerCmp1carType) GetCmp1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp1carFieldCmp1xMask) >> RegisterCmp1carFieldCmp1xShift)
}

// SetCmp1x Timerx Compare 1 value
func (r *registerCmp1carType) SetCmp1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp1carFieldCmp1xMask)|(uint32(value)<<RegisterCmp1carFieldCmp1xShift))
}

// registerCmp2arType Timerx Compare 2 Register
type registerCmp2arType uint32

const (
	RegisterCmp2arFieldCmp2xShift = 0
	RegisterCmp2arFieldCmp2xMask  = 0xffff
)

// GetCmp2x Timerx Compare 2 value
func (r *registerCmp2arType) GetCmp2x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp2arFieldCmp2xMask) >> RegisterCmp2arFieldCmp2xShift)
}

// SetCmp2x Timerx Compare 2 value
func (r *registerCmp2arType) SetCmp2x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp2arFieldCmp2xMask)|(uint32(value)<<RegisterCmp2arFieldCmp2xShift))
}

// registerCmp3arType Timerx Compare 3 Register
type registerCmp3arType uint32

const (
	RegisterCmp3arFieldCmp3xShift = 0
	RegisterCmp3arFieldCmp3xMask  = 0xffff
)

// GetCmp3x Timerx Compare 3 value
func (r *registerCmp3arType) GetCmp3x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp3arFieldCmp3xMask) >> RegisterCmp3arFieldCmp3xShift)
}

// SetCmp3x Timerx Compare 3 value
func (r *registerCmp3arType) SetCmp3x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp3arFieldCmp3xMask)|(uint32(value)<<RegisterCmp3arFieldCmp3xShift))
}

// registerCmp4arType Timerx Compare 4 Register
type registerCmp4arType uint32

const (
	RegisterCmp4arFieldCmp4xShift = 0
	RegisterCmp4arFieldCmp4xMask  = 0xffff
)

// GetCmp4x Timerx Compare 4 value
func (r *registerCmp4arType) GetCmp4x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmp4arFieldCmp4xMask) >> RegisterCmp4arFieldCmp4xShift)
}

// SetCmp4x Timerx Compare 4 value
func (r *registerCmp4arType) SetCmp4x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmp4arFieldCmp4xMask)|(uint32(value)<<RegisterCmp4arFieldCmp4xShift))
}

// registerCpt1arType Timerx Capture 1 Register
type registerCpt1arType uint32

const (
	RegisterCpt1arFieldCpt1xShift = 0
	RegisterCpt1arFieldCpt1xMask  = 0xffff
)

// GetCpt1x Timerx Capture 1 value
func (r *registerCpt1arType) GetCpt1x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpt1arFieldCpt1xMask) >> RegisterCpt1arFieldCpt1xShift)
}

// SetCpt1x Timerx Capture 1 value
func (r *registerCpt1arType) SetCpt1x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpt1arFieldCpt1xMask)|(uint32(value)<<RegisterCpt1arFieldCpt1xShift))
}

// registerCpt2arType Timerx Capture 2 Register
type registerCpt2arType uint32

const (
	RegisterCpt2arFieldCpt2xShift = 0
	RegisterCpt2arFieldCpt2xMask  = 0xffff
)

// GetCpt2x Timerx Capture 2 value
func (r *registerCpt2arType) GetCpt2x() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpt2arFieldCpt2xMask) >> RegisterCpt2arFieldCpt2xShift)
}

// SetCpt2x Timerx Capture 2 value
func (r *registerCpt2arType) SetCpt2x(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpt2arFieldCpt2xMask)|(uint32(value)<<RegisterCpt2arFieldCpt2xShift))
}

// registerDtarType Timerx Deadtime Register
type registerDtarType uint32

const (
	RegisterDtarFieldDtflkxShift = 31
	RegisterDtarFieldDtflkxMask  = 0x80000000
)

// GetDtflkx Deadtime Falling Lock
func (r *registerDtarType) GetDtflkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtarFieldDtflkxMask) != 0
}

// SetDtflkx Deadtime Falling Lock
func (r *registerDtarType) SetDtflkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtarFieldDtflkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtarFieldDtflkxMask)
	}
}

const (
	RegisterDtarFieldDtfslkxShift = 30
	RegisterDtarFieldDtfslkxMask  = 0x40000000
)

// GetDtfslkx Deadtime Falling Sign Lock
func (r *registerDtarType) GetDtfslkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtarFieldDtfslkxMask) != 0
}

// SetDtfslkx Deadtime Falling Sign Lock
func (r *registerDtarType) SetDtfslkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtarFieldDtfslkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtarFieldDtfslkxMask)
	}
}

const (
	RegisterDtarFieldSdtfxShift = 25
	RegisterDtarFieldSdtfxMask  = 0x2000000
)

// GetSdtfx Sign Deadtime Falling value
func (r *registerDtarType) GetSdtfx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtarFieldSdtfxMask) != 0
}

// SetSdtfx Sign Deadtime Falling value
func (r *registerDtarType) SetSdtfx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtarFieldSdtfxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtarFieldSdtfxMask)
	}
}

const (
	RegisterDtarFieldDtfxShift = 16
	RegisterDtarFieldDtfxMask  = 0x1ff0000
)

// GetDtfx Deadtime Falling value
func (r *registerDtarType) GetDtfx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDtarFieldDtfxMask) >> RegisterDtarFieldDtfxShift)
}

// SetDtfx Deadtime Falling value
func (r *registerDtarType) SetDtfx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtarFieldDtfxMask)|(uint32(value)<<RegisterDtarFieldDtfxShift))
}

const (
	RegisterDtarFieldDtrlkxShift = 15
	RegisterDtarFieldDtrlkxMask  = 0x8000
)

// GetDtrlkx Deadtime Rising Lock
func (r *registerDtarType) GetDtrlkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtarFieldDtrlkxMask) != 0
}

// SetDtrlkx Deadtime Rising Lock
func (r *registerDtarType) SetDtrlkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtarFieldDtrlkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtarFieldDtrlkxMask)
	}
}

const (
	RegisterDtarFieldDtrslkxShift = 14
	RegisterDtarFieldDtrslkxMask  = 0x4000
)

// GetDtrslkx Deadtime Rising Sign Lock
func (r *registerDtarType) GetDtrslkx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtarFieldDtrslkxMask) != 0
}

// SetDtrslkx Deadtime Rising Sign Lock
func (r *registerDtarType) SetDtrslkx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtarFieldDtrslkxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtarFieldDtrslkxMask)
	}
}

const (
	RegisterDtarFieldDtprscShift = 10
	RegisterDtarFieldDtprscMask  = 0x1c00
)

// GetDtprsc Deadtime Prescaler
func (r *registerDtarType) GetDtprsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDtarFieldDtprscMask) >> RegisterDtarFieldDtprscShift)
}

// SetDtprsc Deadtime Prescaler
func (r *registerDtarType) SetDtprsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtarFieldDtprscMask)|(uint32(value)<<RegisterDtarFieldDtprscShift))
}

const (
	RegisterDtarFieldSdtrxShift = 9
	RegisterDtarFieldSdtrxMask  = 0x200
)

// GetSdtrx Sign Deadtime Rising value
func (r *registerDtarType) GetSdtrx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtarFieldSdtrxMask) != 0
}

// SetSdtrx Sign Deadtime Rising value
func (r *registerDtarType) SetSdtrx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtarFieldSdtrxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtarFieldSdtrxMask)
	}
}

const (
	RegisterDtarFieldDtrxShift = 0
	RegisterDtarFieldDtrxMask  = 0x1ff
)

// GetDtrx Deadtime Rising value
func (r *registerDtarType) GetDtrx() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDtarFieldDtrxMask) >> RegisterDtarFieldDtrxShift)
}

// SetDtrx Deadtime Rising value
func (r *registerDtarType) SetDtrx(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtarFieldDtrxMask)|(uint32(value)<<RegisterDtarFieldDtrxShift))
}

// registerSeta1rType Timerx Output1 Set Register
type registerSeta1rType uint32

const (
	RegisterSeta1rFieldUpdateShift = 31
	RegisterSeta1rFieldUpdateMask  = 0x80000000
)

// GetUpdate Registers update (transfer preload to active)
func (r *registerSeta1rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldUpdateMask) != 0
}

// SetUpdate Registers update (transfer preload to active)
func (r *registerSeta1rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldUpdateMask)
	}
}

const (
	RegisterSeta1rFieldExtevnt10Shift = 30
	RegisterSeta1rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 External Event 10
func (r *registerSeta1rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 External Event 10
func (r *registerSeta1rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldExtevnt10Mask)
	}
}

const (
	RegisterSeta1rFieldExtevnt9Shift = 29
	RegisterSeta1rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 External Event 9
func (r *registerSeta1rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 External Event 9
func (r *registerSeta1rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldExtevnt9Mask)
	}
}

const (
	RegisterSeta1rFieldExtevnt8Shift = 28
	RegisterSeta1rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 External Event 8
func (r *registerSeta1rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 External Event 8
func (r *registerSeta1rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldExtevnt8Mask)
	}
}

const (
	RegisterSeta1rFieldExtevnt7Shift = 27
	RegisterSeta1rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 External Event 7
func (r *registerSeta1rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 External Event 7
func (r *registerSeta1rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldExtevnt7Mask)
	}
}

const (
	RegisterSeta1rFieldExtevnt6Shift = 26
	RegisterSeta1rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 External Event 6
func (r *registerSeta1rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 External Event 6
func (r *registerSeta1rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldExtevnt6Mask)
	}
}

const (
	RegisterSeta1rFieldExtevnt5Shift = 25
	RegisterSeta1rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 External Event 5
func (r *registerSeta1rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 External Event 5
func (r *registerSeta1rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldExtevnt5Mask)
	}
}

const (
	RegisterSeta1rFieldExtevnt4Shift = 24
	RegisterSeta1rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 External Event 4
func (r *registerSeta1rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 External Event 4
func (r *registerSeta1rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldExtevnt4Mask)
	}
}

const (
	RegisterSeta1rFieldExtevnt3Shift = 23
	RegisterSeta1rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 External Event 3
func (r *registerSeta1rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 External Event 3
func (r *registerSeta1rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldExtevnt3Mask)
	}
}

const (
	RegisterSeta1rFieldExtevnt2Shift = 22
	RegisterSeta1rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 External Event 2
func (r *registerSeta1rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 External Event 2
func (r *registerSeta1rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldExtevnt2Mask)
	}
}

const (
	RegisterSeta1rFieldExtevnt1Shift = 21
	RegisterSeta1rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 External Event 1
func (r *registerSeta1rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 External Event 1
func (r *registerSeta1rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldExtevnt1Mask)
	}
}

const (
	RegisterSeta1rFieldTimevnt9Shift = 20
	RegisterSeta1rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 Timer Event 9
func (r *registerSeta1rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 Timer Event 9
func (r *registerSeta1rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldTimevnt9Mask)
	}
}

const (
	RegisterSeta1rFieldTimevnt8Shift = 19
	RegisterSeta1rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 Timer Event 8
func (r *registerSeta1rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 Timer Event 8
func (r *registerSeta1rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldTimevnt8Mask)
	}
}

const (
	RegisterSeta1rFieldTimevnt7Shift = 18
	RegisterSeta1rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 Timer Event 7
func (r *registerSeta1rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 Timer Event 7
func (r *registerSeta1rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldTimevnt7Mask)
	}
}

const (
	RegisterSeta1rFieldTimevnt6Shift = 17
	RegisterSeta1rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 Timer Event 6
func (r *registerSeta1rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 Timer Event 6
func (r *registerSeta1rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldTimevnt6Mask)
	}
}

const (
	RegisterSeta1rFieldTimevnt5Shift = 16
	RegisterSeta1rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 Timer Event 5
func (r *registerSeta1rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 Timer Event 5
func (r *registerSeta1rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldTimevnt5Mask)
	}
}

const (
	RegisterSeta1rFieldTimevnt4Shift = 15
	RegisterSeta1rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 Timer Event 4
func (r *registerSeta1rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 Timer Event 4
func (r *registerSeta1rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldTimevnt4Mask)
	}
}

const (
	RegisterSeta1rFieldTimevnt3Shift = 14
	RegisterSeta1rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 Timer Event 3
func (r *registerSeta1rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 Timer Event 3
func (r *registerSeta1rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldTimevnt3Mask)
	}
}

const (
	RegisterSeta1rFieldTimevnt2Shift = 13
	RegisterSeta1rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 Timer Event 2
func (r *registerSeta1rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 Timer Event 2
func (r *registerSeta1rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldTimevnt2Mask)
	}
}

const (
	RegisterSeta1rFieldTimevnt1Shift = 12
	RegisterSeta1rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 Timer Event 1
func (r *registerSeta1rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 Timer Event 1
func (r *registerSeta1rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldTimevnt1Mask)
	}
}

const (
	RegisterSeta1rFieldMstcmp4Shift = 11
	RegisterSeta1rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 Master Compare 4
func (r *registerSeta1rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 Master Compare 4
func (r *registerSeta1rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldMstcmp4Mask)
	}
}

const (
	RegisterSeta1rFieldMstcmp3Shift = 10
	RegisterSeta1rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 Master Compare 3
func (r *registerSeta1rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 Master Compare 3
func (r *registerSeta1rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldMstcmp3Mask)
	}
}

const (
	RegisterSeta1rFieldMstcmp2Shift = 9
	RegisterSeta1rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 Master Compare 2
func (r *registerSeta1rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 Master Compare 2
func (r *registerSeta1rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldMstcmp2Mask)
	}
}

const (
	RegisterSeta1rFieldMstcmp1Shift = 8
	RegisterSeta1rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 Master Compare 1
func (r *registerSeta1rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 Master Compare 1
func (r *registerSeta1rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldMstcmp1Mask)
	}
}

const (
	RegisterSeta1rFieldMstperShift = 7
	RegisterSeta1rFieldMstperMask  = 0x80
)

// GetMstper Master Period
func (r *registerSeta1rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldMstperMask) != 0
}

// SetMstper Master Period
func (r *registerSeta1rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldMstperMask)
	}
}

const (
	RegisterSeta1rFieldCmp4Shift = 6
	RegisterSeta1rFieldCmp4Mask  = 0x40
)

// GetCmp4 Timer A compare 4
func (r *registerSeta1rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldCmp4Mask) != 0
}

// SetCmp4 Timer A compare 4
func (r *registerSeta1rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldCmp4Mask)
	}
}

const (
	RegisterSeta1rFieldCmp3Shift = 5
	RegisterSeta1rFieldCmp3Mask  = 0x20
)

// GetCmp3 Timer A compare 3
func (r *registerSeta1rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldCmp3Mask) != 0
}

// SetCmp3 Timer A compare 3
func (r *registerSeta1rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldCmp3Mask)
	}
}

const (
	RegisterSeta1rFieldCmp2Shift = 4
	RegisterSeta1rFieldCmp2Mask  = 0x10
)

// GetCmp2 Timer A compare 2
func (r *registerSeta1rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldCmp2Mask) != 0
}

// SetCmp2 Timer A compare 2
func (r *registerSeta1rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldCmp2Mask)
	}
}

const (
	RegisterSeta1rFieldCmp1Shift = 3
	RegisterSeta1rFieldCmp1Mask  = 0x8
)

// GetCmp1 Timer A compare 1
func (r *registerSeta1rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldCmp1Mask) != 0
}

// SetCmp1 Timer A compare 1
func (r *registerSeta1rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldCmp1Mask)
	}
}

const (
	RegisterSeta1rFieldPerShift = 2
	RegisterSeta1rFieldPerMask  = 0x4
)

// GetPer Timer A Period
func (r *registerSeta1rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldPerMask) != 0
}

// SetPer Timer A Period
func (r *registerSeta1rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldPerMask)
	}
}

const (
	RegisterSeta1rFieldResyncShift = 1
	RegisterSeta1rFieldResyncMask  = 0x2
)

// GetResync Timer A resynchronizaton
func (r *registerSeta1rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldResyncMask) != 0
}

// SetResync Timer A resynchronizaton
func (r *registerSeta1rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldResyncMask)
	}
}

const (
	RegisterSeta1rFieldSstShift = 0
	RegisterSeta1rFieldSstMask  = 0x1
)

// GetSst Software Set trigger
func (r *registerSeta1rType) GetSst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta1rFieldSstMask) != 0
}

// SetSst Software Set trigger
func (r *registerSeta1rType) SetSst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta1rFieldSstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta1rFieldSstMask)
	}
}

// registerRsta1rType Timerx Output1 Reset Register
type registerRsta1rType uint32

const (
	RegisterRsta1rFieldUpdateShift = 31
	RegisterRsta1rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerRsta1rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerRsta1rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldUpdateMask)
	}
}

const (
	RegisterRsta1rFieldExtevnt10Shift = 30
	RegisterRsta1rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerRsta1rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerRsta1rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldExtevnt10Mask)
	}
}

const (
	RegisterRsta1rFieldExtevnt9Shift = 29
	RegisterRsta1rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerRsta1rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerRsta1rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldExtevnt9Mask)
	}
}

const (
	RegisterRsta1rFieldExtevnt8Shift = 28
	RegisterRsta1rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerRsta1rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerRsta1rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldExtevnt8Mask)
	}
}

const (
	RegisterRsta1rFieldExtevnt7Shift = 27
	RegisterRsta1rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerRsta1rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerRsta1rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldExtevnt7Mask)
	}
}

const (
	RegisterRsta1rFieldExtevnt6Shift = 26
	RegisterRsta1rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerRsta1rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerRsta1rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldExtevnt6Mask)
	}
}

const (
	RegisterRsta1rFieldExtevnt5Shift = 25
	RegisterRsta1rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerRsta1rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerRsta1rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldExtevnt5Mask)
	}
}

const (
	RegisterRsta1rFieldExtevnt4Shift = 24
	RegisterRsta1rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerRsta1rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerRsta1rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldExtevnt4Mask)
	}
}

const (
	RegisterRsta1rFieldExtevnt3Shift = 23
	RegisterRsta1rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerRsta1rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerRsta1rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldExtevnt3Mask)
	}
}

const (
	RegisterRsta1rFieldExtevnt2Shift = 22
	RegisterRsta1rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerRsta1rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerRsta1rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldExtevnt2Mask)
	}
}

const (
	RegisterRsta1rFieldExtevnt1Shift = 21
	RegisterRsta1rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerRsta1rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerRsta1rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldExtevnt1Mask)
	}
}

const (
	RegisterRsta1rFieldTimevnt9Shift = 20
	RegisterRsta1rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerRsta1rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerRsta1rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldTimevnt9Mask)
	}
}

const (
	RegisterRsta1rFieldTimevnt8Shift = 19
	RegisterRsta1rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerRsta1rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerRsta1rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldTimevnt8Mask)
	}
}

const (
	RegisterRsta1rFieldTimevnt7Shift = 18
	RegisterRsta1rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerRsta1rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerRsta1rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldTimevnt7Mask)
	}
}

const (
	RegisterRsta1rFieldTimevnt6Shift = 17
	RegisterRsta1rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerRsta1rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerRsta1rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldTimevnt6Mask)
	}
}

const (
	RegisterRsta1rFieldTimevnt5Shift = 16
	RegisterRsta1rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerRsta1rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerRsta1rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldTimevnt5Mask)
	}
}

const (
	RegisterRsta1rFieldTimevnt4Shift = 15
	RegisterRsta1rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerRsta1rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerRsta1rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldTimevnt4Mask)
	}
}

const (
	RegisterRsta1rFieldTimevnt3Shift = 14
	RegisterRsta1rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerRsta1rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerRsta1rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldTimevnt3Mask)
	}
}

const (
	RegisterRsta1rFieldTimevnt2Shift = 13
	RegisterRsta1rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerRsta1rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerRsta1rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldTimevnt2Mask)
	}
}

const (
	RegisterRsta1rFieldTimevnt1Shift = 12
	RegisterRsta1rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerRsta1rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerRsta1rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldTimevnt1Mask)
	}
}

const (
	RegisterRsta1rFieldMstcmp4Shift = 11
	RegisterRsta1rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerRsta1rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerRsta1rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldMstcmp4Mask)
	}
}

const (
	RegisterRsta1rFieldMstcmp3Shift = 10
	RegisterRsta1rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerRsta1rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerRsta1rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldMstcmp3Mask)
	}
}

const (
	RegisterRsta1rFieldMstcmp2Shift = 9
	RegisterRsta1rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerRsta1rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerRsta1rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldMstcmp2Mask)
	}
}

const (
	RegisterRsta1rFieldMstcmp1Shift = 8
	RegisterRsta1rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerRsta1rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerRsta1rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldMstcmp1Mask)
	}
}

const (
	RegisterRsta1rFieldMstperShift = 7
	RegisterRsta1rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerRsta1rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerRsta1rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldMstperMask)
	}
}

const (
	RegisterRsta1rFieldCmp4Shift = 6
	RegisterRsta1rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerRsta1rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerRsta1rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldCmp4Mask)
	}
}

const (
	RegisterRsta1rFieldCmp3Shift = 5
	RegisterRsta1rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerRsta1rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerRsta1rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldCmp3Mask)
	}
}

const (
	RegisterRsta1rFieldCmp2Shift = 4
	RegisterRsta1rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerRsta1rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerRsta1rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldCmp2Mask)
	}
}

const (
	RegisterRsta1rFieldCmp1Shift = 3
	RegisterRsta1rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerRsta1rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerRsta1rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldCmp1Mask)
	}
}

const (
	RegisterRsta1rFieldPerShift = 2
	RegisterRsta1rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerRsta1rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldPerMask) != 0
}

// SetPer PER
func (r *registerRsta1rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldPerMask)
	}
}

const (
	RegisterRsta1rFieldResyncShift = 1
	RegisterRsta1rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerRsta1rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerRsta1rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldResyncMask)
	}
}

const (
	RegisterRsta1rFieldSrtShift = 0
	RegisterRsta1rFieldSrtMask  = 0x1
)

// GetSrt SRT
func (r *registerRsta1rType) GetSrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta1rFieldSrtMask) != 0
}

// SetSrt SRT
func (r *registerRsta1rType) SetSrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta1rFieldSrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta1rFieldSrtMask)
	}
}

// registerSeta2rType Timerx Output2 Set Register
type registerSeta2rType uint32

const (
	RegisterSeta2rFieldUpdateShift = 31
	RegisterSeta2rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerSeta2rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerSeta2rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldUpdateMask)
	}
}

const (
	RegisterSeta2rFieldExtevnt10Shift = 30
	RegisterSeta2rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerSeta2rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerSeta2rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldExtevnt10Mask)
	}
}

const (
	RegisterSeta2rFieldExtevnt9Shift = 29
	RegisterSeta2rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerSeta2rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerSeta2rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldExtevnt9Mask)
	}
}

const (
	RegisterSeta2rFieldExtevnt8Shift = 28
	RegisterSeta2rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerSeta2rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerSeta2rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldExtevnt8Mask)
	}
}

const (
	RegisterSeta2rFieldExtevnt7Shift = 27
	RegisterSeta2rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerSeta2rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerSeta2rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldExtevnt7Mask)
	}
}

const (
	RegisterSeta2rFieldExtevnt6Shift = 26
	RegisterSeta2rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerSeta2rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerSeta2rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldExtevnt6Mask)
	}
}

const (
	RegisterSeta2rFieldExtevnt5Shift = 25
	RegisterSeta2rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerSeta2rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerSeta2rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldExtevnt5Mask)
	}
}

const (
	RegisterSeta2rFieldExtevnt4Shift = 24
	RegisterSeta2rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerSeta2rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerSeta2rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldExtevnt4Mask)
	}
}

const (
	RegisterSeta2rFieldExtevnt3Shift = 23
	RegisterSeta2rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerSeta2rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerSeta2rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldExtevnt3Mask)
	}
}

const (
	RegisterSeta2rFieldExtevnt2Shift = 22
	RegisterSeta2rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerSeta2rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerSeta2rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldExtevnt2Mask)
	}
}

const (
	RegisterSeta2rFieldExtevnt1Shift = 21
	RegisterSeta2rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerSeta2rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerSeta2rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldExtevnt1Mask)
	}
}

const (
	RegisterSeta2rFieldTimevnt9Shift = 20
	RegisterSeta2rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerSeta2rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerSeta2rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldTimevnt9Mask)
	}
}

const (
	RegisterSeta2rFieldTimevnt8Shift = 19
	RegisterSeta2rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerSeta2rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerSeta2rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldTimevnt8Mask)
	}
}

const (
	RegisterSeta2rFieldTimevnt7Shift = 18
	RegisterSeta2rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerSeta2rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerSeta2rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldTimevnt7Mask)
	}
}

const (
	RegisterSeta2rFieldTimevnt6Shift = 17
	RegisterSeta2rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerSeta2rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerSeta2rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldTimevnt6Mask)
	}
}

const (
	RegisterSeta2rFieldTimevnt5Shift = 16
	RegisterSeta2rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerSeta2rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerSeta2rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldTimevnt5Mask)
	}
}

const (
	RegisterSeta2rFieldTimevnt4Shift = 15
	RegisterSeta2rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerSeta2rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerSeta2rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldTimevnt4Mask)
	}
}

const (
	RegisterSeta2rFieldTimevnt3Shift = 14
	RegisterSeta2rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerSeta2rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerSeta2rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldTimevnt3Mask)
	}
}

const (
	RegisterSeta2rFieldTimevnt2Shift = 13
	RegisterSeta2rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerSeta2rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerSeta2rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldTimevnt2Mask)
	}
}

const (
	RegisterSeta2rFieldTimevnt1Shift = 12
	RegisterSeta2rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerSeta2rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerSeta2rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldTimevnt1Mask)
	}
}

const (
	RegisterSeta2rFieldMstcmp4Shift = 11
	RegisterSeta2rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerSeta2rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerSeta2rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldMstcmp4Mask)
	}
}

const (
	RegisterSeta2rFieldMstcmp3Shift = 10
	RegisterSeta2rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerSeta2rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerSeta2rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldMstcmp3Mask)
	}
}

const (
	RegisterSeta2rFieldMstcmp2Shift = 9
	RegisterSeta2rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerSeta2rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerSeta2rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldMstcmp2Mask)
	}
}

const (
	RegisterSeta2rFieldMstcmp1Shift = 8
	RegisterSeta2rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerSeta2rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerSeta2rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldMstcmp1Mask)
	}
}

const (
	RegisterSeta2rFieldMstperShift = 7
	RegisterSeta2rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerSeta2rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerSeta2rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldMstperMask)
	}
}

const (
	RegisterSeta2rFieldCmp4Shift = 6
	RegisterSeta2rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerSeta2rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerSeta2rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldCmp4Mask)
	}
}

const (
	RegisterSeta2rFieldCmp3Shift = 5
	RegisterSeta2rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerSeta2rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerSeta2rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldCmp3Mask)
	}
}

const (
	RegisterSeta2rFieldCmp2Shift = 4
	RegisterSeta2rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerSeta2rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerSeta2rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldCmp2Mask)
	}
}

const (
	RegisterSeta2rFieldCmp1Shift = 3
	RegisterSeta2rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerSeta2rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerSeta2rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldCmp1Mask)
	}
}

const (
	RegisterSeta2rFieldPerShift = 2
	RegisterSeta2rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerSeta2rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldPerMask) != 0
}

// SetPer PER
func (r *registerSeta2rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldPerMask)
	}
}

const (
	RegisterSeta2rFieldResyncShift = 1
	RegisterSeta2rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerSeta2rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerSeta2rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldResyncMask)
	}
}

const (
	RegisterSeta2rFieldSstShift = 0
	RegisterSeta2rFieldSstMask  = 0x1
)

// GetSst SST
func (r *registerSeta2rType) GetSst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSeta2rFieldSstMask) != 0
}

// SetSst SST
func (r *registerSeta2rType) SetSst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSeta2rFieldSstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSeta2rFieldSstMask)
	}
}

// registerRsta2rType Timerx Output2 Reset Register
type registerRsta2rType uint32

const (
	RegisterRsta2rFieldUpdateShift = 31
	RegisterRsta2rFieldUpdateMask  = 0x80000000
)

// GetUpdate UPDATE
func (r *registerRsta2rType) GetUpdate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldUpdateMask) != 0
}

// SetUpdate UPDATE
func (r *registerRsta2rType) SetUpdate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldUpdateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldUpdateMask)
	}
}

const (
	RegisterRsta2rFieldExtevnt10Shift = 30
	RegisterRsta2rFieldExtevnt10Mask  = 0x40000000
)

// GetExtevnt10 EXTEVNT10
func (r *registerRsta2rType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldExtevnt10Mask) != 0
}

// SetExtevnt10 EXTEVNT10
func (r *registerRsta2rType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldExtevnt10Mask)
	}
}

const (
	RegisterRsta2rFieldExtevnt9Shift = 29
	RegisterRsta2rFieldExtevnt9Mask  = 0x20000000
)

// GetExtevnt9 EXTEVNT9
func (r *registerRsta2rType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldExtevnt9Mask) != 0
}

// SetExtevnt9 EXTEVNT9
func (r *registerRsta2rType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldExtevnt9Mask)
	}
}

const (
	RegisterRsta2rFieldExtevnt8Shift = 28
	RegisterRsta2rFieldExtevnt8Mask  = 0x10000000
)

// GetExtevnt8 EXTEVNT8
func (r *registerRsta2rType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldExtevnt8Mask) != 0
}

// SetExtevnt8 EXTEVNT8
func (r *registerRsta2rType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldExtevnt8Mask)
	}
}

const (
	RegisterRsta2rFieldExtevnt7Shift = 27
	RegisterRsta2rFieldExtevnt7Mask  = 0x8000000
)

// GetExtevnt7 EXTEVNT7
func (r *registerRsta2rType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldExtevnt7Mask) != 0
}

// SetExtevnt7 EXTEVNT7
func (r *registerRsta2rType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldExtevnt7Mask)
	}
}

const (
	RegisterRsta2rFieldExtevnt6Shift = 26
	RegisterRsta2rFieldExtevnt6Mask  = 0x4000000
)

// GetExtevnt6 EXTEVNT6
func (r *registerRsta2rType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldExtevnt6Mask) != 0
}

// SetExtevnt6 EXTEVNT6
func (r *registerRsta2rType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldExtevnt6Mask)
	}
}

const (
	RegisterRsta2rFieldExtevnt5Shift = 25
	RegisterRsta2rFieldExtevnt5Mask  = 0x2000000
)

// GetExtevnt5 EXTEVNT5
func (r *registerRsta2rType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldExtevnt5Mask) != 0
}

// SetExtevnt5 EXTEVNT5
func (r *registerRsta2rType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldExtevnt5Mask)
	}
}

const (
	RegisterRsta2rFieldExtevnt4Shift = 24
	RegisterRsta2rFieldExtevnt4Mask  = 0x1000000
)

// GetExtevnt4 EXTEVNT4
func (r *registerRsta2rType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldExtevnt4Mask) != 0
}

// SetExtevnt4 EXTEVNT4
func (r *registerRsta2rType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldExtevnt4Mask)
	}
}

const (
	RegisterRsta2rFieldExtevnt3Shift = 23
	RegisterRsta2rFieldExtevnt3Mask  = 0x800000
)

// GetExtevnt3 EXTEVNT3
func (r *registerRsta2rType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldExtevnt3Mask) != 0
}

// SetExtevnt3 EXTEVNT3
func (r *registerRsta2rType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldExtevnt3Mask)
	}
}

const (
	RegisterRsta2rFieldExtevnt2Shift = 22
	RegisterRsta2rFieldExtevnt2Mask  = 0x400000
)

// GetExtevnt2 EXTEVNT2
func (r *registerRsta2rType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldExtevnt2Mask) != 0
}

// SetExtevnt2 EXTEVNT2
func (r *registerRsta2rType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldExtevnt2Mask)
	}
}

const (
	RegisterRsta2rFieldExtevnt1Shift = 21
	RegisterRsta2rFieldExtevnt1Mask  = 0x200000
)

// GetExtevnt1 EXTEVNT1
func (r *registerRsta2rType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldExtevnt1Mask) != 0
}

// SetExtevnt1 EXTEVNT1
func (r *registerRsta2rType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldExtevnt1Mask)
	}
}

const (
	RegisterRsta2rFieldTimevnt9Shift = 20
	RegisterRsta2rFieldTimevnt9Mask  = 0x100000
)

// GetTimevnt9 TIMEVNT9
func (r *registerRsta2rType) GetTimevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldTimevnt9Mask) != 0
}

// SetTimevnt9 TIMEVNT9
func (r *registerRsta2rType) SetTimevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldTimevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldTimevnt9Mask)
	}
}

const (
	RegisterRsta2rFieldTimevnt8Shift = 19
	RegisterRsta2rFieldTimevnt8Mask  = 0x80000
)

// GetTimevnt8 TIMEVNT8
func (r *registerRsta2rType) GetTimevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldTimevnt8Mask) != 0
}

// SetTimevnt8 TIMEVNT8
func (r *registerRsta2rType) SetTimevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldTimevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldTimevnt8Mask)
	}
}

const (
	RegisterRsta2rFieldTimevnt7Shift = 18
	RegisterRsta2rFieldTimevnt7Mask  = 0x40000
)

// GetTimevnt7 TIMEVNT7
func (r *registerRsta2rType) GetTimevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldTimevnt7Mask) != 0
}

// SetTimevnt7 TIMEVNT7
func (r *registerRsta2rType) SetTimevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldTimevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldTimevnt7Mask)
	}
}

const (
	RegisterRsta2rFieldTimevnt6Shift = 17
	RegisterRsta2rFieldTimevnt6Mask  = 0x20000
)

// GetTimevnt6 TIMEVNT6
func (r *registerRsta2rType) GetTimevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldTimevnt6Mask) != 0
}

// SetTimevnt6 TIMEVNT6
func (r *registerRsta2rType) SetTimevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldTimevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldTimevnt6Mask)
	}
}

const (
	RegisterRsta2rFieldTimevnt5Shift = 16
	RegisterRsta2rFieldTimevnt5Mask  = 0x10000
)

// GetTimevnt5 TIMEVNT5
func (r *registerRsta2rType) GetTimevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldTimevnt5Mask) != 0
}

// SetTimevnt5 TIMEVNT5
func (r *registerRsta2rType) SetTimevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldTimevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldTimevnt5Mask)
	}
}

const (
	RegisterRsta2rFieldTimevnt4Shift = 15
	RegisterRsta2rFieldTimevnt4Mask  = 0x8000
)

// GetTimevnt4 TIMEVNT4
func (r *registerRsta2rType) GetTimevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldTimevnt4Mask) != 0
}

// SetTimevnt4 TIMEVNT4
func (r *registerRsta2rType) SetTimevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldTimevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldTimevnt4Mask)
	}
}

const (
	RegisterRsta2rFieldTimevnt3Shift = 14
	RegisterRsta2rFieldTimevnt3Mask  = 0x4000
)

// GetTimevnt3 TIMEVNT3
func (r *registerRsta2rType) GetTimevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldTimevnt3Mask) != 0
}

// SetTimevnt3 TIMEVNT3
func (r *registerRsta2rType) SetTimevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldTimevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldTimevnt3Mask)
	}
}

const (
	RegisterRsta2rFieldTimevnt2Shift = 13
	RegisterRsta2rFieldTimevnt2Mask  = 0x2000
)

// GetTimevnt2 TIMEVNT2
func (r *registerRsta2rType) GetTimevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldTimevnt2Mask) != 0
}

// SetTimevnt2 TIMEVNT2
func (r *registerRsta2rType) SetTimevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldTimevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldTimevnt2Mask)
	}
}

const (
	RegisterRsta2rFieldTimevnt1Shift = 12
	RegisterRsta2rFieldTimevnt1Mask  = 0x1000
)

// GetTimevnt1 TIMEVNT1
func (r *registerRsta2rType) GetTimevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldTimevnt1Mask) != 0
}

// SetTimevnt1 TIMEVNT1
func (r *registerRsta2rType) SetTimevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldTimevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldTimevnt1Mask)
	}
}

const (
	RegisterRsta2rFieldMstcmp4Shift = 11
	RegisterRsta2rFieldMstcmp4Mask  = 0x800
)

// GetMstcmp4 MSTCMP4
func (r *registerRsta2rType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldMstcmp4Mask) != 0
}

// SetMstcmp4 MSTCMP4
func (r *registerRsta2rType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldMstcmp4Mask)
	}
}

const (
	RegisterRsta2rFieldMstcmp3Shift = 10
	RegisterRsta2rFieldMstcmp3Mask  = 0x400
)

// GetMstcmp3 MSTCMP3
func (r *registerRsta2rType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldMstcmp3Mask) != 0
}

// SetMstcmp3 MSTCMP3
func (r *registerRsta2rType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldMstcmp3Mask)
	}
}

const (
	RegisterRsta2rFieldMstcmp2Shift = 9
	RegisterRsta2rFieldMstcmp2Mask  = 0x200
)

// GetMstcmp2 MSTCMP2
func (r *registerRsta2rType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldMstcmp2Mask) != 0
}

// SetMstcmp2 MSTCMP2
func (r *registerRsta2rType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldMstcmp2Mask)
	}
}

const (
	RegisterRsta2rFieldMstcmp1Shift = 8
	RegisterRsta2rFieldMstcmp1Mask  = 0x100
)

// GetMstcmp1 MSTCMP1
func (r *registerRsta2rType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldMstcmp1Mask) != 0
}

// SetMstcmp1 MSTCMP1
func (r *registerRsta2rType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldMstcmp1Mask)
	}
}

const (
	RegisterRsta2rFieldMstperShift = 7
	RegisterRsta2rFieldMstperMask  = 0x80
)

// GetMstper MSTPER
func (r *registerRsta2rType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldMstperMask) != 0
}

// SetMstper MSTPER
func (r *registerRsta2rType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldMstperMask)
	}
}

const (
	RegisterRsta2rFieldCmp4Shift = 6
	RegisterRsta2rFieldCmp4Mask  = 0x40
)

// GetCmp4 CMP4
func (r *registerRsta2rType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldCmp4Mask) != 0
}

// SetCmp4 CMP4
func (r *registerRsta2rType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldCmp4Mask)
	}
}

const (
	RegisterRsta2rFieldCmp3Shift = 5
	RegisterRsta2rFieldCmp3Mask  = 0x20
)

// GetCmp3 CMP3
func (r *registerRsta2rType) GetCmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldCmp3Mask) != 0
}

// SetCmp3 CMP3
func (r *registerRsta2rType) SetCmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldCmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldCmp3Mask)
	}
}

const (
	RegisterRsta2rFieldCmp2Shift = 4
	RegisterRsta2rFieldCmp2Mask  = 0x10
)

// GetCmp2 CMP2
func (r *registerRsta2rType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldCmp2Mask) != 0
}

// SetCmp2 CMP2
func (r *registerRsta2rType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldCmp2Mask)
	}
}

const (
	RegisterRsta2rFieldCmp1Shift = 3
	RegisterRsta2rFieldCmp1Mask  = 0x8
)

// GetCmp1 CMP1
func (r *registerRsta2rType) GetCmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldCmp1Mask) != 0
}

// SetCmp1 CMP1
func (r *registerRsta2rType) SetCmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldCmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldCmp1Mask)
	}
}

const (
	RegisterRsta2rFieldPerShift = 2
	RegisterRsta2rFieldPerMask  = 0x4
)

// GetPer PER
func (r *registerRsta2rType) GetPer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldPerMask) != 0
}

// SetPer PER
func (r *registerRsta2rType) SetPer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldPerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldPerMask)
	}
}

const (
	RegisterRsta2rFieldResyncShift = 1
	RegisterRsta2rFieldResyncMask  = 0x2
)

// GetResync RESYNC
func (r *registerRsta2rType) GetResync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldResyncMask) != 0
}

// SetResync RESYNC
func (r *registerRsta2rType) SetResync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldResyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldResyncMask)
	}
}

const (
	RegisterRsta2rFieldSrtShift = 0
	RegisterRsta2rFieldSrtMask  = 0x1
)

// GetSrt SRT
func (r *registerRsta2rType) GetSrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsta2rFieldSrtMask) != 0
}

// SetSrt SRT
func (r *registerRsta2rType) SetSrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsta2rFieldSrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsta2rFieldSrtMask)
	}
}

// registerEefar1Type Timerx External Event Filtering Register 1
type registerEefar1Type uint32

const (
	RegisterEefar1FieldEe5fltrShift = 25
	RegisterEefar1FieldEe5fltrMask  = 0x1e000000
)

// GetEe5fltr External Event 5 filter
func (r *registerEefar1Type) GetEe5fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefar1FieldEe5fltrMask) >> RegisterEefar1FieldEe5fltrShift)
}

// SetEe5fltr External Event 5 filter
func (r *registerEefar1Type) SetEe5fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefar1FieldEe5fltrMask)|(uint32(value)<<RegisterEefar1FieldEe5fltrShift))
}

const (
	RegisterEefar1FieldEe5ltchShift = 24
	RegisterEefar1FieldEe5ltchMask  = 0x1000000
)

// GetEe5ltch External Event 5 latch
func (r *registerEefar1Type) GetEe5ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefar1FieldEe5ltchMask) != 0
}

// SetEe5ltch External Event 5 latch
func (r *registerEefar1Type) SetEe5ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefar1FieldEe5ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefar1FieldEe5ltchMask)
	}
}

const (
	RegisterEefar1FieldEe4fltrShift = 19
	RegisterEefar1FieldEe4fltrMask  = 0x780000
)

// GetEe4fltr External Event 4 filter
func (r *registerEefar1Type) GetEe4fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefar1FieldEe4fltrMask) >> RegisterEefar1FieldEe4fltrShift)
}

// SetEe4fltr External Event 4 filter
func (r *registerEefar1Type) SetEe4fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefar1FieldEe4fltrMask)|(uint32(value)<<RegisterEefar1FieldEe4fltrShift))
}

const (
	RegisterEefar1FieldEe4ltchShift = 18
	RegisterEefar1FieldEe4ltchMask  = 0x40000
)

// GetEe4ltch External Event 4 latch
func (r *registerEefar1Type) GetEe4ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefar1FieldEe4ltchMask) != 0
}

// SetEe4ltch External Event 4 latch
func (r *registerEefar1Type) SetEe4ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefar1FieldEe4ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefar1FieldEe4ltchMask)
	}
}

const (
	RegisterEefar1FieldEe3fltrShift = 13
	RegisterEefar1FieldEe3fltrMask  = 0x1e000
)

// GetEe3fltr External Event 3 filter
func (r *registerEefar1Type) GetEe3fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefar1FieldEe3fltrMask) >> RegisterEefar1FieldEe3fltrShift)
}

// SetEe3fltr External Event 3 filter
func (r *registerEefar1Type) SetEe3fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefar1FieldEe3fltrMask)|(uint32(value)<<RegisterEefar1FieldEe3fltrShift))
}

const (
	RegisterEefar1FieldEe3ltchShift = 12
	RegisterEefar1FieldEe3ltchMask  = 0x1000
)

// GetEe3ltch External Event 3 latch
func (r *registerEefar1Type) GetEe3ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefar1FieldEe3ltchMask) != 0
}

// SetEe3ltch External Event 3 latch
func (r *registerEefar1Type) SetEe3ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefar1FieldEe3ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefar1FieldEe3ltchMask)
	}
}

const (
	RegisterEefar1FieldEe2fltrShift = 7
	RegisterEefar1FieldEe2fltrMask  = 0x780
)

// GetEe2fltr External Event 2 filter
func (r *registerEefar1Type) GetEe2fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefar1FieldEe2fltrMask) >> RegisterEefar1FieldEe2fltrShift)
}

// SetEe2fltr External Event 2 filter
func (r *registerEefar1Type) SetEe2fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefar1FieldEe2fltrMask)|(uint32(value)<<RegisterEefar1FieldEe2fltrShift))
}

const (
	RegisterEefar1FieldEe2ltchShift = 6
	RegisterEefar1FieldEe2ltchMask  = 0x40
)

// GetEe2ltch External Event 2 latch
func (r *registerEefar1Type) GetEe2ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefar1FieldEe2ltchMask) != 0
}

// SetEe2ltch External Event 2 latch
func (r *registerEefar1Type) SetEe2ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefar1FieldEe2ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefar1FieldEe2ltchMask)
	}
}

const (
	RegisterEefar1FieldEe1fltrShift = 1
	RegisterEefar1FieldEe1fltrMask  = 0x1e
)

// GetEe1fltr External Event 1 filter
func (r *registerEefar1Type) GetEe1fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefar1FieldEe1fltrMask) >> RegisterEefar1FieldEe1fltrShift)
}

// SetEe1fltr External Event 1 filter
func (r *registerEefar1Type) SetEe1fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefar1FieldEe1fltrMask)|(uint32(value)<<RegisterEefar1FieldEe1fltrShift))
}

const (
	RegisterEefar1FieldEe1ltchShift = 0
	RegisterEefar1FieldEe1ltchMask  = 0x1
)

// GetEe1ltch External Event 1 latch
func (r *registerEefar1Type) GetEe1ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefar1FieldEe1ltchMask) != 0
}

// SetEe1ltch External Event 1 latch
func (r *registerEefar1Type) SetEe1ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefar1FieldEe1ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefar1FieldEe1ltchMask)
	}
}

// registerEefar2Type Timerx External Event Filtering Register 2
type registerEefar2Type uint32

const (
	RegisterEefar2FieldEe10fltrShift = 25
	RegisterEefar2FieldEe10fltrMask  = 0x1e000000
)

// GetEe10fltr External Event 10 filter
func (r *registerEefar2Type) GetEe10fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefar2FieldEe10fltrMask) >> RegisterEefar2FieldEe10fltrShift)
}

// SetEe10fltr External Event 10 filter
func (r *registerEefar2Type) SetEe10fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefar2FieldEe10fltrMask)|(uint32(value)<<RegisterEefar2FieldEe10fltrShift))
}

const (
	RegisterEefar2FieldEe10ltchShift = 24
	RegisterEefar2FieldEe10ltchMask  = 0x1000000
)

// GetEe10ltch External Event 10 latch
func (r *registerEefar2Type) GetEe10ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefar2FieldEe10ltchMask) != 0
}

// SetEe10ltch External Event 10 latch
func (r *registerEefar2Type) SetEe10ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefar2FieldEe10ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefar2FieldEe10ltchMask)
	}
}

const (
	RegisterEefar2FieldEe9fltrShift = 19
	RegisterEefar2FieldEe9fltrMask  = 0x780000
)

// GetEe9fltr External Event 9 filter
func (r *registerEefar2Type) GetEe9fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefar2FieldEe9fltrMask) >> RegisterEefar2FieldEe9fltrShift)
}

// SetEe9fltr External Event 9 filter
func (r *registerEefar2Type) SetEe9fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefar2FieldEe9fltrMask)|(uint32(value)<<RegisterEefar2FieldEe9fltrShift))
}

const (
	RegisterEefar2FieldEe9ltchShift = 18
	RegisterEefar2FieldEe9ltchMask  = 0x40000
)

// GetEe9ltch External Event 9 latch
func (r *registerEefar2Type) GetEe9ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefar2FieldEe9ltchMask) != 0
}

// SetEe9ltch External Event 9 latch
func (r *registerEefar2Type) SetEe9ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefar2FieldEe9ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefar2FieldEe9ltchMask)
	}
}

const (
	RegisterEefar2FieldEe8fltrShift = 13
	RegisterEefar2FieldEe8fltrMask  = 0x1e000
)

// GetEe8fltr External Event 8 filter
func (r *registerEefar2Type) GetEe8fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefar2FieldEe8fltrMask) >> RegisterEefar2FieldEe8fltrShift)
}

// SetEe8fltr External Event 8 filter
func (r *registerEefar2Type) SetEe8fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefar2FieldEe8fltrMask)|(uint32(value)<<RegisterEefar2FieldEe8fltrShift))
}

const (
	RegisterEefar2FieldEe8ltchShift = 12
	RegisterEefar2FieldEe8ltchMask  = 0x1000
)

// GetEe8ltch External Event 8 latch
func (r *registerEefar2Type) GetEe8ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefar2FieldEe8ltchMask) != 0
}

// SetEe8ltch External Event 8 latch
func (r *registerEefar2Type) SetEe8ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefar2FieldEe8ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefar2FieldEe8ltchMask)
	}
}

const (
	RegisterEefar2FieldEe7fltrShift = 7
	RegisterEefar2FieldEe7fltrMask  = 0x780
)

// GetEe7fltr External Event 7 filter
func (r *registerEefar2Type) GetEe7fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefar2FieldEe7fltrMask) >> RegisterEefar2FieldEe7fltrShift)
}

// SetEe7fltr External Event 7 filter
func (r *registerEefar2Type) SetEe7fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefar2FieldEe7fltrMask)|(uint32(value)<<RegisterEefar2FieldEe7fltrShift))
}

const (
	RegisterEefar2FieldEe7ltchShift = 6
	RegisterEefar2FieldEe7ltchMask  = 0x40
)

// GetEe7ltch External Event 7 latch
func (r *registerEefar2Type) GetEe7ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefar2FieldEe7ltchMask) != 0
}

// SetEe7ltch External Event 7 latch
func (r *registerEefar2Type) SetEe7ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefar2FieldEe7ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefar2FieldEe7ltchMask)
	}
}

const (
	RegisterEefar2FieldEe6fltrShift = 1
	RegisterEefar2FieldEe6fltrMask  = 0x1e
)

// GetEe6fltr External Event 6 filter
func (r *registerEefar2Type) GetEe6fltr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterEefar2FieldEe6fltrMask) >> RegisterEefar2FieldEe6fltrShift)
}

// SetEe6fltr External Event 6 filter
func (r *registerEefar2Type) SetEe6fltr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEefar2FieldEe6fltrMask)|(uint32(value)<<RegisterEefar2FieldEe6fltrShift))
}

const (
	RegisterEefar2FieldEe6ltchShift = 0
	RegisterEefar2FieldEe6ltchMask  = 0x1
)

// GetEe6ltch External Event 6 latch
func (r *registerEefar2Type) GetEe6ltch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEefar2FieldEe6ltchMask) != 0
}

// SetEe6ltch External Event 6 latch
func (r *registerEefar2Type) SetEe6ltch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEefar2FieldEe6ltchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEefar2FieldEe6ltchMask)
	}
}

// registerRstarType TimerA Reset Register
type registerRstarType uint32

const (
	RegisterRstarFieldTimecmp4Shift = 30
	RegisterRstarFieldTimecmp4Mask  = 0x40000000
)

// GetTimecmp4 Timer E Compare 4
func (r *registerRstarType) GetTimecmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldTimecmp4Mask) != 0
}

// SetTimecmp4 Timer E Compare 4
func (r *registerRstarType) SetTimecmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldTimecmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldTimecmp4Mask)
	}
}

const (
	RegisterRstarFieldTimecmp2Shift = 29
	RegisterRstarFieldTimecmp2Mask  = 0x20000000
)

// GetTimecmp2 Timer E Compare 2
func (r *registerRstarType) GetTimecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldTimecmp2Mask) != 0
}

// SetTimecmp2 Timer E Compare 2
func (r *registerRstarType) SetTimecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldTimecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldTimecmp2Mask)
	}
}

const (
	RegisterRstarFieldTimecmp1Shift = 28
	RegisterRstarFieldTimecmp1Mask  = 0x10000000
)

// GetTimecmp1 Timer E Compare 1
func (r *registerRstarType) GetTimecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldTimecmp1Mask) != 0
}

// SetTimecmp1 Timer E Compare 1
func (r *registerRstarType) SetTimecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldTimecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldTimecmp1Mask)
	}
}

const (
	RegisterRstarFieldTimdcmp4Shift = 27
	RegisterRstarFieldTimdcmp4Mask  = 0x8000000
)

// GetTimdcmp4 Timer D Compare 4
func (r *registerRstarType) GetTimdcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldTimdcmp4Mask) != 0
}

// SetTimdcmp4 Timer D Compare 4
func (r *registerRstarType) SetTimdcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldTimdcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldTimdcmp4Mask)
	}
}

const (
	RegisterRstarFieldTimdcmp2Shift = 26
	RegisterRstarFieldTimdcmp2Mask  = 0x4000000
)

// GetTimdcmp2 Timer D Compare 2
func (r *registerRstarType) GetTimdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldTimdcmp2Mask) != 0
}

// SetTimdcmp2 Timer D Compare 2
func (r *registerRstarType) SetTimdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldTimdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldTimdcmp2Mask)
	}
}

const (
	RegisterRstarFieldTimdcmp1Shift = 25
	RegisterRstarFieldTimdcmp1Mask  = 0x2000000
)

// GetTimdcmp1 Timer D Compare 1
func (r *registerRstarType) GetTimdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldTimdcmp1Mask) != 0
}

// SetTimdcmp1 Timer D Compare 1
func (r *registerRstarType) SetTimdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldTimdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldTimdcmp1Mask)
	}
}

const (
	RegisterRstarFieldTimccmp4Shift = 24
	RegisterRstarFieldTimccmp4Mask  = 0x1000000
)

// GetTimccmp4 Timer C Compare 4
func (r *registerRstarType) GetTimccmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldTimccmp4Mask) != 0
}

// SetTimccmp4 Timer C Compare 4
func (r *registerRstarType) SetTimccmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldTimccmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldTimccmp4Mask)
	}
}

const (
	RegisterRstarFieldTimccmp2Shift = 23
	RegisterRstarFieldTimccmp2Mask  = 0x800000
)

// GetTimccmp2 Timer C Compare 2
func (r *registerRstarType) GetTimccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldTimccmp2Mask) != 0
}

// SetTimccmp2 Timer C Compare 2
func (r *registerRstarType) SetTimccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldTimccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldTimccmp2Mask)
	}
}

const (
	RegisterRstarFieldTimccmp1Shift = 22
	RegisterRstarFieldTimccmp1Mask  = 0x400000
)

// GetTimccmp1 Timer C Compare 1
func (r *registerRstarType) GetTimccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldTimccmp1Mask) != 0
}

// SetTimccmp1 Timer C Compare 1
func (r *registerRstarType) SetTimccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldTimccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldTimccmp1Mask)
	}
}

const (
	RegisterRstarFieldTimbcmp4Shift = 21
	RegisterRstarFieldTimbcmp4Mask  = 0x200000
)

// GetTimbcmp4 Timer B Compare 4
func (r *registerRstarType) GetTimbcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldTimbcmp4Mask) != 0
}

// SetTimbcmp4 Timer B Compare 4
func (r *registerRstarType) SetTimbcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldTimbcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldTimbcmp4Mask)
	}
}

const (
	RegisterRstarFieldTimbcmp2Shift = 20
	RegisterRstarFieldTimbcmp2Mask  = 0x100000
)

// GetTimbcmp2 Timer B Compare 2
func (r *registerRstarType) GetTimbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldTimbcmp2Mask) != 0
}

// SetTimbcmp2 Timer B Compare 2
func (r *registerRstarType) SetTimbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldTimbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldTimbcmp2Mask)
	}
}

const (
	RegisterRstarFieldTimbcmp1Shift = 19
	RegisterRstarFieldTimbcmp1Mask  = 0x80000
)

// GetTimbcmp1 Timer B Compare 1
func (r *registerRstarType) GetTimbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldTimbcmp1Mask) != 0
}

// SetTimbcmp1 Timer B Compare 1
func (r *registerRstarType) SetTimbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldTimbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldTimbcmp1Mask)
	}
}

const (
	RegisterRstarFieldExtevnt10Shift = 18
	RegisterRstarFieldExtevnt10Mask  = 0x40000
)

// GetExtevnt10 External Event 10
func (r *registerRstarType) GetExtevnt10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldExtevnt10Mask) != 0
}

// SetExtevnt10 External Event 10
func (r *registerRstarType) SetExtevnt10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldExtevnt10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldExtevnt10Mask)
	}
}

const (
	RegisterRstarFieldExtevnt9Shift = 17
	RegisterRstarFieldExtevnt9Mask  = 0x20000
)

// GetExtevnt9 External Event 9
func (r *registerRstarType) GetExtevnt9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldExtevnt9Mask) != 0
}

// SetExtevnt9 External Event 9
func (r *registerRstarType) SetExtevnt9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldExtevnt9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldExtevnt9Mask)
	}
}

const (
	RegisterRstarFieldExtevnt8Shift = 16
	RegisterRstarFieldExtevnt8Mask  = 0x10000
)

// GetExtevnt8 External Event 8
func (r *registerRstarType) GetExtevnt8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldExtevnt8Mask) != 0
}

// SetExtevnt8 External Event 8
func (r *registerRstarType) SetExtevnt8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldExtevnt8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldExtevnt8Mask)
	}
}

const (
	RegisterRstarFieldExtevnt7Shift = 15
	RegisterRstarFieldExtevnt7Mask  = 0x8000
)

// GetExtevnt7 External Event 7
func (r *registerRstarType) GetExtevnt7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldExtevnt7Mask) != 0
}

// SetExtevnt7 External Event 7
func (r *registerRstarType) SetExtevnt7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldExtevnt7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldExtevnt7Mask)
	}
}

const (
	RegisterRstarFieldExtevnt6Shift = 14
	RegisterRstarFieldExtevnt6Mask  = 0x4000
)

// GetExtevnt6 External Event 6
func (r *registerRstarType) GetExtevnt6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldExtevnt6Mask) != 0
}

// SetExtevnt6 External Event 6
func (r *registerRstarType) SetExtevnt6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldExtevnt6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldExtevnt6Mask)
	}
}

const (
	RegisterRstarFieldExtevnt5Shift = 13
	RegisterRstarFieldExtevnt5Mask  = 0x2000
)

// GetExtevnt5 External Event 5
func (r *registerRstarType) GetExtevnt5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldExtevnt5Mask) != 0
}

// SetExtevnt5 External Event 5
func (r *registerRstarType) SetExtevnt5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldExtevnt5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldExtevnt5Mask)
	}
}

const (
	RegisterRstarFieldExtevnt4Shift = 12
	RegisterRstarFieldExtevnt4Mask  = 0x1000
)

// GetExtevnt4 External Event 4
func (r *registerRstarType) GetExtevnt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldExtevnt4Mask) != 0
}

// SetExtevnt4 External Event 4
func (r *registerRstarType) SetExtevnt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldExtevnt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldExtevnt4Mask)
	}
}

const (
	RegisterRstarFieldExtevnt3Shift = 11
	RegisterRstarFieldExtevnt3Mask  = 0x800
)

// GetExtevnt3 External Event 3
func (r *registerRstarType) GetExtevnt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldExtevnt3Mask) != 0
}

// SetExtevnt3 External Event 3
func (r *registerRstarType) SetExtevnt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldExtevnt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldExtevnt3Mask)
	}
}

const (
	RegisterRstarFieldExtevnt2Shift = 10
	RegisterRstarFieldExtevnt2Mask  = 0x400
)

// GetExtevnt2 External Event 2
func (r *registerRstarType) GetExtevnt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldExtevnt2Mask) != 0
}

// SetExtevnt2 External Event 2
func (r *registerRstarType) SetExtevnt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldExtevnt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldExtevnt2Mask)
	}
}

const (
	RegisterRstarFieldExtevnt1Shift = 9
	RegisterRstarFieldExtevnt1Mask  = 0x200
)

// GetExtevnt1 External Event 1
func (r *registerRstarType) GetExtevnt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldExtevnt1Mask) != 0
}

// SetExtevnt1 External Event 1
func (r *registerRstarType) SetExtevnt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldExtevnt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldExtevnt1Mask)
	}
}

const (
	RegisterRstarFieldMstcmp4Shift = 8
	RegisterRstarFieldMstcmp4Mask  = 0x100
)

// GetMstcmp4 Master compare 4
func (r *registerRstarType) GetMstcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldMstcmp4Mask) != 0
}

// SetMstcmp4 Master compare 4
func (r *registerRstarType) SetMstcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldMstcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldMstcmp4Mask)
	}
}

const (
	RegisterRstarFieldMstcmp3Shift = 7
	RegisterRstarFieldMstcmp3Mask  = 0x80
)

// GetMstcmp3 Master compare 3
func (r *registerRstarType) GetMstcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldMstcmp3Mask) != 0
}

// SetMstcmp3 Master compare 3
func (r *registerRstarType) SetMstcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldMstcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldMstcmp3Mask)
	}
}

const (
	RegisterRstarFieldMstcmp2Shift = 6
	RegisterRstarFieldMstcmp2Mask  = 0x40
)

// GetMstcmp2 Master compare 2
func (r *registerRstarType) GetMstcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldMstcmp2Mask) != 0
}

// SetMstcmp2 Master compare 2
func (r *registerRstarType) SetMstcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldMstcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldMstcmp2Mask)
	}
}

const (
	RegisterRstarFieldMstcmp1Shift = 5
	RegisterRstarFieldMstcmp1Mask  = 0x20
)

// GetMstcmp1 Master compare 1
func (r *registerRstarType) GetMstcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldMstcmp1Mask) != 0
}

// SetMstcmp1 Master compare 1
func (r *registerRstarType) SetMstcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldMstcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldMstcmp1Mask)
	}
}

const (
	RegisterRstarFieldMstperShift = 4
	RegisterRstarFieldMstperMask  = 0x10
)

// GetMstper Master timer Period
func (r *registerRstarType) GetMstper() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldMstperMask) != 0
}

// SetMstper Master timer Period
func (r *registerRstarType) SetMstper(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldMstperMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldMstperMask)
	}
}

const (
	RegisterRstarFieldCmp4Shift = 3
	RegisterRstarFieldCmp4Mask  = 0x8
)

// GetCmp4 Timer A compare 4 reset
func (r *registerRstarType) GetCmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldCmp4Mask) != 0
}

// SetCmp4 Timer A compare 4 reset
func (r *registerRstarType) SetCmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldCmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldCmp4Mask)
	}
}

const (
	RegisterRstarFieldCmp2Shift = 2
	RegisterRstarFieldCmp2Mask  = 0x4
)

// GetCmp2 Timer A compare 2 reset
func (r *registerRstarType) GetCmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldCmp2Mask) != 0
}

// SetCmp2 Timer A compare 2 reset
func (r *registerRstarType) SetCmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldCmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldCmp2Mask)
	}
}

const (
	RegisterRstarFieldUpdtShift = 1
	RegisterRstarFieldUpdtMask  = 0x2
)

// GetUpdt Timer A Update reset
func (r *registerRstarType) GetUpdt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRstarFieldUpdtMask) != 0
}

// SetUpdt Timer A Update reset
func (r *registerRstarType) SetUpdt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRstarFieldUpdtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRstarFieldUpdtMask)
	}
}

// registerChparType Timerx Chopper Register
type registerChparType uint32

const (
	RegisterChparFieldStrtpwShift = 7
	RegisterChparFieldStrtpwMask  = 0x780
)

// GetStrtpw STRTPW
func (r *registerChparType) GetStrtpw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChparFieldStrtpwMask) >> RegisterChparFieldStrtpwShift)
}

// SetStrtpw STRTPW
func (r *registerChparType) SetStrtpw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChparFieldStrtpwMask)|(uint32(value)<<RegisterChparFieldStrtpwShift))
}

const (
	RegisterChparFieldChpdtyShift = 4
	RegisterChparFieldChpdtyMask  = 0x70
)

// GetChpdty Timerx chopper duty cycle value
func (r *registerChparType) GetChpdty() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChparFieldChpdtyMask) >> RegisterChparFieldChpdtyShift)
}

// SetChpdty Timerx chopper duty cycle value
func (r *registerChparType) SetChpdty(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChparFieldChpdtyMask)|(uint32(value)<<RegisterChparFieldChpdtyShift))
}

const (
	RegisterChparFieldChpfrqShift = 0
	RegisterChparFieldChpfrqMask  = 0xf
)

// GetChpfrq Timerx carrier frequency value
func (r *registerChparType) GetChpfrq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterChparFieldChpfrqMask) >> RegisterChparFieldChpfrqShift)
}

// SetChpfrq Timerx carrier frequency value
func (r *registerChparType) SetChpfrq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterChparFieldChpfrqMask)|(uint32(value)<<RegisterChparFieldChpfrqShift))
}

// registerCpt1acrType Timerx Capture 2 Control Register
type registerCpt1acrType uint32

const (
	RegisterCpt1acrFieldTecmp2Shift = 31
	RegisterCpt1acrFieldTecmp2Mask  = 0x80000000
)

// GetTecmp2 Timer E Compare 2
func (r *registerCpt1acrType) GetTecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTecmp2Mask) != 0
}

// SetTecmp2 Timer E Compare 2
func (r *registerCpt1acrType) SetTecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTecmp2Mask)
	}
}

const (
	RegisterCpt1acrFieldTecmp1Shift = 30
	RegisterCpt1acrFieldTecmp1Mask  = 0x40000000
)

// GetTecmp1 Timer E Compare 1
func (r *registerCpt1acrType) GetTecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTecmp1Mask) != 0
}

// SetTecmp1 Timer E Compare 1
func (r *registerCpt1acrType) SetTecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTecmp1Mask)
	}
}

const (
	RegisterCpt1acrFieldTe1rstShift = 29
	RegisterCpt1acrFieldTe1rstMask  = 0x20000000
)

// GetTe1rst Timer E output 1 Reset
func (r *registerCpt1acrType) GetTe1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTe1rstMask) != 0
}

// SetTe1rst Timer E output 1 Reset
func (r *registerCpt1acrType) SetTe1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTe1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTe1rstMask)
	}
}

const (
	RegisterCpt1acrFieldTe1setShift = 28
	RegisterCpt1acrFieldTe1setMask  = 0x10000000
)

// GetTe1set Timer E output 1 Set
func (r *registerCpt1acrType) GetTe1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTe1setMask) != 0
}

// SetTe1set Timer E output 1 Set
func (r *registerCpt1acrType) SetTe1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTe1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTe1setMask)
	}
}

const (
	RegisterCpt1acrFieldTdcmp2Shift = 27
	RegisterCpt1acrFieldTdcmp2Mask  = 0x8000000
)

// GetTdcmp2 Timer D Compare 2
func (r *registerCpt1acrType) GetTdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTdcmp2Mask) != 0
}

// SetTdcmp2 Timer D Compare 2
func (r *registerCpt1acrType) SetTdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTdcmp2Mask)
	}
}

const (
	RegisterCpt1acrFieldTdcmp1Shift = 26
	RegisterCpt1acrFieldTdcmp1Mask  = 0x4000000
)

// GetTdcmp1 Timer D Compare 1
func (r *registerCpt1acrType) GetTdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTdcmp1Mask) != 0
}

// SetTdcmp1 Timer D Compare 1
func (r *registerCpt1acrType) SetTdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTdcmp1Mask)
	}
}

const (
	RegisterCpt1acrFieldTd1rstShift = 25
	RegisterCpt1acrFieldTd1rstMask  = 0x2000000
)

// GetTd1rst Timer D output 1 Reset
func (r *registerCpt1acrType) GetTd1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTd1rstMask) != 0
}

// SetTd1rst Timer D output 1 Reset
func (r *registerCpt1acrType) SetTd1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTd1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTd1rstMask)
	}
}

const (
	RegisterCpt1acrFieldTd1setShift = 24
	RegisterCpt1acrFieldTd1setMask  = 0x1000000
)

// GetTd1set Timer D output 1 Set
func (r *registerCpt1acrType) GetTd1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTd1setMask) != 0
}

// SetTd1set Timer D output 1 Set
func (r *registerCpt1acrType) SetTd1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTd1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTd1setMask)
	}
}

const (
	RegisterCpt1acrFieldTccmp2Shift = 23
	RegisterCpt1acrFieldTccmp2Mask  = 0x800000
)

// GetTccmp2 Timer C Compare 2
func (r *registerCpt1acrType) GetTccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTccmp2Mask) != 0
}

// SetTccmp2 Timer C Compare 2
func (r *registerCpt1acrType) SetTccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTccmp2Mask)
	}
}

const (
	RegisterCpt1acrFieldTccmp1Shift = 22
	RegisterCpt1acrFieldTccmp1Mask  = 0x400000
)

// GetTccmp1 Timer C Compare 1
func (r *registerCpt1acrType) GetTccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTccmp1Mask) != 0
}

// SetTccmp1 Timer C Compare 1
func (r *registerCpt1acrType) SetTccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTccmp1Mask)
	}
}

const (
	RegisterCpt1acrFieldTc1rstShift = 21
	RegisterCpt1acrFieldTc1rstMask  = 0x200000
)

// GetTc1rst Timer C output 1 Reset
func (r *registerCpt1acrType) GetTc1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTc1rstMask) != 0
}

// SetTc1rst Timer C output 1 Reset
func (r *registerCpt1acrType) SetTc1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTc1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTc1rstMask)
	}
}

const (
	RegisterCpt1acrFieldTc1setShift = 20
	RegisterCpt1acrFieldTc1setMask  = 0x100000
)

// GetTc1set Timer C output 1 Set
func (r *registerCpt1acrType) GetTc1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTc1setMask) != 0
}

// SetTc1set Timer C output 1 Set
func (r *registerCpt1acrType) SetTc1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTc1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTc1setMask)
	}
}

const (
	RegisterCpt1acrFieldTbcmp2Shift = 19
	RegisterCpt1acrFieldTbcmp2Mask  = 0x80000
)

// GetTbcmp2 Timer B Compare 2
func (r *registerCpt1acrType) GetTbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTbcmp2Mask) != 0
}

// SetTbcmp2 Timer B Compare 2
func (r *registerCpt1acrType) SetTbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTbcmp2Mask)
	}
}

const (
	RegisterCpt1acrFieldTbcmp1Shift = 18
	RegisterCpt1acrFieldTbcmp1Mask  = 0x40000
)

// GetTbcmp1 Timer B Compare 1
func (r *registerCpt1acrType) GetTbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTbcmp1Mask) != 0
}

// SetTbcmp1 Timer B Compare 1
func (r *registerCpt1acrType) SetTbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTbcmp1Mask)
	}
}

const (
	RegisterCpt1acrFieldTb1rstShift = 17
	RegisterCpt1acrFieldTb1rstMask  = 0x20000
)

// GetTb1rst Timer B output 1 Reset
func (r *registerCpt1acrType) GetTb1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTb1rstMask) != 0
}

// SetTb1rst Timer B output 1 Reset
func (r *registerCpt1acrType) SetTb1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTb1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTb1rstMask)
	}
}

const (
	RegisterCpt1acrFieldTb1setShift = 16
	RegisterCpt1acrFieldTb1setMask  = 0x10000
)

// GetTb1set Timer B output 1 Set
func (r *registerCpt1acrType) GetTb1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldTb1setMask) != 0
}

// SetTb1set Timer B output 1 Set
func (r *registerCpt1acrType) SetTb1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldTb1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldTb1setMask)
	}
}

const (
	RegisterCpt1acrFieldExev10cptShift = 11
	RegisterCpt1acrFieldExev10cptMask  = 0x800
)

// GetExev10cpt External Event 10 Capture
func (r *registerCpt1acrType) GetExev10cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldExev10cptMask) != 0
}

// SetExev10cpt External Event 10 Capture
func (r *registerCpt1acrType) SetExev10cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldExev10cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldExev10cptMask)
	}
}

const (
	RegisterCpt1acrFieldExev9cptShift = 10
	RegisterCpt1acrFieldExev9cptMask  = 0x400
)

// GetExev9cpt External Event 9 Capture
func (r *registerCpt1acrType) GetExev9cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldExev9cptMask) != 0
}

// SetExev9cpt External Event 9 Capture
func (r *registerCpt1acrType) SetExev9cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldExev9cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldExev9cptMask)
	}
}

const (
	RegisterCpt1acrFieldExev8cptShift = 9
	RegisterCpt1acrFieldExev8cptMask  = 0x200
)

// GetExev8cpt External Event 8 Capture
func (r *registerCpt1acrType) GetExev8cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldExev8cptMask) != 0
}

// SetExev8cpt External Event 8 Capture
func (r *registerCpt1acrType) SetExev8cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldExev8cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldExev8cptMask)
	}
}

const (
	RegisterCpt1acrFieldExev7cptShift = 8
	RegisterCpt1acrFieldExev7cptMask  = 0x100
)

// GetExev7cpt External Event 7 Capture
func (r *registerCpt1acrType) GetExev7cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldExev7cptMask) != 0
}

// SetExev7cpt External Event 7 Capture
func (r *registerCpt1acrType) SetExev7cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldExev7cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldExev7cptMask)
	}
}

const (
	RegisterCpt1acrFieldExev6cptShift = 7
	RegisterCpt1acrFieldExev6cptMask  = 0x80
)

// GetExev6cpt External Event 6 Capture
func (r *registerCpt1acrType) GetExev6cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldExev6cptMask) != 0
}

// SetExev6cpt External Event 6 Capture
func (r *registerCpt1acrType) SetExev6cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldExev6cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldExev6cptMask)
	}
}

const (
	RegisterCpt1acrFieldExev5cptShift = 6
	RegisterCpt1acrFieldExev5cptMask  = 0x40
)

// GetExev5cpt External Event 5 Capture
func (r *registerCpt1acrType) GetExev5cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldExev5cptMask) != 0
}

// SetExev5cpt External Event 5 Capture
func (r *registerCpt1acrType) SetExev5cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldExev5cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldExev5cptMask)
	}
}

const (
	RegisterCpt1acrFieldExev4cptShift = 5
	RegisterCpt1acrFieldExev4cptMask  = 0x20
)

// GetExev4cpt External Event 4 Capture
func (r *registerCpt1acrType) GetExev4cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldExev4cptMask) != 0
}

// SetExev4cpt External Event 4 Capture
func (r *registerCpt1acrType) SetExev4cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldExev4cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldExev4cptMask)
	}
}

const (
	RegisterCpt1acrFieldExev3cptShift = 4
	RegisterCpt1acrFieldExev3cptMask  = 0x10
)

// GetExev3cpt External Event 3 Capture
func (r *registerCpt1acrType) GetExev3cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldExev3cptMask) != 0
}

// SetExev3cpt External Event 3 Capture
func (r *registerCpt1acrType) SetExev3cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldExev3cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldExev3cptMask)
	}
}

const (
	RegisterCpt1acrFieldExev2cptShift = 3
	RegisterCpt1acrFieldExev2cptMask  = 0x8
)

// GetExev2cpt External Event 2 Capture
func (r *registerCpt1acrType) GetExev2cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldExev2cptMask) != 0
}

// SetExev2cpt External Event 2 Capture
func (r *registerCpt1acrType) SetExev2cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldExev2cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldExev2cptMask)
	}
}

const (
	RegisterCpt1acrFieldExev1cptShift = 2
	RegisterCpt1acrFieldExev1cptMask  = 0x4
)

// GetExev1cpt External Event 1 Capture
func (r *registerCpt1acrType) GetExev1cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldExev1cptMask) != 0
}

// SetExev1cpt External Event 1 Capture
func (r *registerCpt1acrType) SetExev1cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldExev1cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldExev1cptMask)
	}
}

const (
	RegisterCpt1acrFieldUdpcptShift = 1
	RegisterCpt1acrFieldUdpcptMask  = 0x2
)

// GetUdpcpt Update Capture
func (r *registerCpt1acrType) GetUdpcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldUdpcptMask) != 0
}

// SetUdpcpt Update Capture
func (r *registerCpt1acrType) SetUdpcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldUdpcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldUdpcptMask)
	}
}

const (
	RegisterCpt1acrFieldSwcptShift = 0
	RegisterCpt1acrFieldSwcptMask  = 0x1
)

// GetSwcpt Software Capture
func (r *registerCpt1acrType) GetSwcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt1acrFieldSwcptMask) != 0
}

// SetSwcpt Software Capture
func (r *registerCpt1acrType) SetSwcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt1acrFieldSwcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt1acrFieldSwcptMask)
	}
}

// registerCpt2acrType CPT2xCR
type registerCpt2acrType uint32

const (
	RegisterCpt2acrFieldTecmp2Shift = 31
	RegisterCpt2acrFieldTecmp2Mask  = 0x80000000
)

// GetTecmp2 Timer E Compare 2
func (r *registerCpt2acrType) GetTecmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTecmp2Mask) != 0
}

// SetTecmp2 Timer E Compare 2
func (r *registerCpt2acrType) SetTecmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTecmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTecmp2Mask)
	}
}

const (
	RegisterCpt2acrFieldTecmp1Shift = 30
	RegisterCpt2acrFieldTecmp1Mask  = 0x40000000
)

// GetTecmp1 Timer E Compare 1
func (r *registerCpt2acrType) GetTecmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTecmp1Mask) != 0
}

// SetTecmp1 Timer E Compare 1
func (r *registerCpt2acrType) SetTecmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTecmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTecmp1Mask)
	}
}

const (
	RegisterCpt2acrFieldTe1rstShift = 29
	RegisterCpt2acrFieldTe1rstMask  = 0x20000000
)

// GetTe1rst Timer E output 1 Reset
func (r *registerCpt2acrType) GetTe1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTe1rstMask) != 0
}

// SetTe1rst Timer E output 1 Reset
func (r *registerCpt2acrType) SetTe1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTe1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTe1rstMask)
	}
}

const (
	RegisterCpt2acrFieldTe1setShift = 28
	RegisterCpt2acrFieldTe1setMask  = 0x10000000
)

// GetTe1set Timer E output 1 Set
func (r *registerCpt2acrType) GetTe1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTe1setMask) != 0
}

// SetTe1set Timer E output 1 Set
func (r *registerCpt2acrType) SetTe1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTe1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTe1setMask)
	}
}

const (
	RegisterCpt2acrFieldTdcmp2Shift = 27
	RegisterCpt2acrFieldTdcmp2Mask  = 0x8000000
)

// GetTdcmp2 Timer D Compare 2
func (r *registerCpt2acrType) GetTdcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTdcmp2Mask) != 0
}

// SetTdcmp2 Timer D Compare 2
func (r *registerCpt2acrType) SetTdcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTdcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTdcmp2Mask)
	}
}

const (
	RegisterCpt2acrFieldTdcmp1Shift = 26
	RegisterCpt2acrFieldTdcmp1Mask  = 0x4000000
)

// GetTdcmp1 Timer D Compare 1
func (r *registerCpt2acrType) GetTdcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTdcmp1Mask) != 0
}

// SetTdcmp1 Timer D Compare 1
func (r *registerCpt2acrType) SetTdcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTdcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTdcmp1Mask)
	}
}

const (
	RegisterCpt2acrFieldTd1rstShift = 25
	RegisterCpt2acrFieldTd1rstMask  = 0x2000000
)

// GetTd1rst Timer D output 1 Reset
func (r *registerCpt2acrType) GetTd1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTd1rstMask) != 0
}

// SetTd1rst Timer D output 1 Reset
func (r *registerCpt2acrType) SetTd1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTd1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTd1rstMask)
	}
}

const (
	RegisterCpt2acrFieldTd1setShift = 24
	RegisterCpt2acrFieldTd1setMask  = 0x1000000
)

// GetTd1set Timer D output 1 Set
func (r *registerCpt2acrType) GetTd1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTd1setMask) != 0
}

// SetTd1set Timer D output 1 Set
func (r *registerCpt2acrType) SetTd1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTd1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTd1setMask)
	}
}

const (
	RegisterCpt2acrFieldTccmp2Shift = 23
	RegisterCpt2acrFieldTccmp2Mask  = 0x800000
)

// GetTccmp2 Timer C Compare 2
func (r *registerCpt2acrType) GetTccmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTccmp2Mask) != 0
}

// SetTccmp2 Timer C Compare 2
func (r *registerCpt2acrType) SetTccmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTccmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTccmp2Mask)
	}
}

const (
	RegisterCpt2acrFieldTccmp1Shift = 22
	RegisterCpt2acrFieldTccmp1Mask  = 0x400000
)

// GetTccmp1 Timer C Compare 1
func (r *registerCpt2acrType) GetTccmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTccmp1Mask) != 0
}

// SetTccmp1 Timer C Compare 1
func (r *registerCpt2acrType) SetTccmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTccmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTccmp1Mask)
	}
}

const (
	RegisterCpt2acrFieldTc1rstShift = 21
	RegisterCpt2acrFieldTc1rstMask  = 0x200000
)

// GetTc1rst Timer C output 1 Reset
func (r *registerCpt2acrType) GetTc1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTc1rstMask) != 0
}

// SetTc1rst Timer C output 1 Reset
func (r *registerCpt2acrType) SetTc1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTc1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTc1rstMask)
	}
}

const (
	RegisterCpt2acrFieldTc1setShift = 20
	RegisterCpt2acrFieldTc1setMask  = 0x100000
)

// GetTc1set Timer C output 1 Set
func (r *registerCpt2acrType) GetTc1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTc1setMask) != 0
}

// SetTc1set Timer C output 1 Set
func (r *registerCpt2acrType) SetTc1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTc1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTc1setMask)
	}
}

const (
	RegisterCpt2acrFieldTbcmp2Shift = 19
	RegisterCpt2acrFieldTbcmp2Mask  = 0x80000
)

// GetTbcmp2 Timer B Compare 2
func (r *registerCpt2acrType) GetTbcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTbcmp2Mask) != 0
}

// SetTbcmp2 Timer B Compare 2
func (r *registerCpt2acrType) SetTbcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTbcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTbcmp2Mask)
	}
}

const (
	RegisterCpt2acrFieldTbcmp1Shift = 18
	RegisterCpt2acrFieldTbcmp1Mask  = 0x40000
)

// GetTbcmp1 Timer B Compare 1
func (r *registerCpt2acrType) GetTbcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTbcmp1Mask) != 0
}

// SetTbcmp1 Timer B Compare 1
func (r *registerCpt2acrType) SetTbcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTbcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTbcmp1Mask)
	}
}

const (
	RegisterCpt2acrFieldTb1rstShift = 17
	RegisterCpt2acrFieldTb1rstMask  = 0x20000
)

// GetTb1rst Timer B output 1 Reset
func (r *registerCpt2acrType) GetTb1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTb1rstMask) != 0
}

// SetTb1rst Timer B output 1 Reset
func (r *registerCpt2acrType) SetTb1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTb1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTb1rstMask)
	}
}

const (
	RegisterCpt2acrFieldTb1setShift = 16
	RegisterCpt2acrFieldTb1setMask  = 0x10000
)

// GetTb1set Timer B output 1 Set
func (r *registerCpt2acrType) GetTb1set() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldTb1setMask) != 0
}

// SetTb1set Timer B output 1 Set
func (r *registerCpt2acrType) SetTb1set(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldTb1setMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldTb1setMask)
	}
}

const (
	RegisterCpt2acrFieldExev10cptShift = 11
	RegisterCpt2acrFieldExev10cptMask  = 0x800
)

// GetExev10cpt External Event 10 Capture
func (r *registerCpt2acrType) GetExev10cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldExev10cptMask) != 0
}

// SetExev10cpt External Event 10 Capture
func (r *registerCpt2acrType) SetExev10cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldExev10cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldExev10cptMask)
	}
}

const (
	RegisterCpt2acrFieldExev9cptShift = 10
	RegisterCpt2acrFieldExev9cptMask  = 0x400
)

// GetExev9cpt External Event 9 Capture
func (r *registerCpt2acrType) GetExev9cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldExev9cptMask) != 0
}

// SetExev9cpt External Event 9 Capture
func (r *registerCpt2acrType) SetExev9cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldExev9cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldExev9cptMask)
	}
}

const (
	RegisterCpt2acrFieldExev8cptShift = 9
	RegisterCpt2acrFieldExev8cptMask  = 0x200
)

// GetExev8cpt External Event 8 Capture
func (r *registerCpt2acrType) GetExev8cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldExev8cptMask) != 0
}

// SetExev8cpt External Event 8 Capture
func (r *registerCpt2acrType) SetExev8cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldExev8cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldExev8cptMask)
	}
}

const (
	RegisterCpt2acrFieldExev7cptShift = 8
	RegisterCpt2acrFieldExev7cptMask  = 0x100
)

// GetExev7cpt External Event 7 Capture
func (r *registerCpt2acrType) GetExev7cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldExev7cptMask) != 0
}

// SetExev7cpt External Event 7 Capture
func (r *registerCpt2acrType) SetExev7cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldExev7cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldExev7cptMask)
	}
}

const (
	RegisterCpt2acrFieldExev6cptShift = 7
	RegisterCpt2acrFieldExev6cptMask  = 0x80
)

// GetExev6cpt External Event 6 Capture
func (r *registerCpt2acrType) GetExev6cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldExev6cptMask) != 0
}

// SetExev6cpt External Event 6 Capture
func (r *registerCpt2acrType) SetExev6cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldExev6cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldExev6cptMask)
	}
}

const (
	RegisterCpt2acrFieldExev5cptShift = 6
	RegisterCpt2acrFieldExev5cptMask  = 0x40
)

// GetExev5cpt External Event 5 Capture
func (r *registerCpt2acrType) GetExev5cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldExev5cptMask) != 0
}

// SetExev5cpt External Event 5 Capture
func (r *registerCpt2acrType) SetExev5cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldExev5cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldExev5cptMask)
	}
}

const (
	RegisterCpt2acrFieldExev4cptShift = 5
	RegisterCpt2acrFieldExev4cptMask  = 0x20
)

// GetExev4cpt External Event 4 Capture
func (r *registerCpt2acrType) GetExev4cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldExev4cptMask) != 0
}

// SetExev4cpt External Event 4 Capture
func (r *registerCpt2acrType) SetExev4cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldExev4cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldExev4cptMask)
	}
}

const (
	RegisterCpt2acrFieldExev3cptShift = 4
	RegisterCpt2acrFieldExev3cptMask  = 0x10
)

// GetExev3cpt External Event 3 Capture
func (r *registerCpt2acrType) GetExev3cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldExev3cptMask) != 0
}

// SetExev3cpt External Event 3 Capture
func (r *registerCpt2acrType) SetExev3cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldExev3cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldExev3cptMask)
	}
}

const (
	RegisterCpt2acrFieldExev2cptShift = 3
	RegisterCpt2acrFieldExev2cptMask  = 0x8
)

// GetExev2cpt External Event 2 Capture
func (r *registerCpt2acrType) GetExev2cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldExev2cptMask) != 0
}

// SetExev2cpt External Event 2 Capture
func (r *registerCpt2acrType) SetExev2cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldExev2cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldExev2cptMask)
	}
}

const (
	RegisterCpt2acrFieldExev1cptShift = 2
	RegisterCpt2acrFieldExev1cptMask  = 0x4
)

// GetExev1cpt External Event 1 Capture
func (r *registerCpt2acrType) GetExev1cpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldExev1cptMask) != 0
}

// SetExev1cpt External Event 1 Capture
func (r *registerCpt2acrType) SetExev1cpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldExev1cptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldExev1cptMask)
	}
}

const (
	RegisterCpt2acrFieldUdpcptShift = 1
	RegisterCpt2acrFieldUdpcptMask  = 0x2
)

// GetUdpcpt Update Capture
func (r *registerCpt2acrType) GetUdpcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldUdpcptMask) != 0
}

// SetUdpcpt Update Capture
func (r *registerCpt2acrType) SetUdpcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldUdpcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldUdpcptMask)
	}
}

const (
	RegisterCpt2acrFieldSwcptShift = 0
	RegisterCpt2acrFieldSwcptMask  = 0x1
)

// GetSwcpt Software Capture
func (r *registerCpt2acrType) GetSwcpt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpt2acrFieldSwcptMask) != 0
}

// SetSwcpt Software Capture
func (r *registerCpt2acrType) SetSwcpt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpt2acrFieldSwcptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpt2acrFieldSwcptMask)
	}
}

// registerOutarType Timerx Output Register
type registerOutarType uint32

const (
	RegisterOutarFieldDidl2Shift = 23
	RegisterOutarFieldDidl2Mask  = 0x800000
)

// GetDidl2 Output 2 Deadtime upon burst mode Idle entry
func (r *registerOutarType) GetDidl2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldDidl2Mask) != 0
}

// SetDidl2 Output 2 Deadtime upon burst mode Idle entry
func (r *registerOutarType) SetDidl2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutarFieldDidl2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldDidl2Mask)
	}
}

const (
	RegisterOutarFieldChp2Shift = 22
	RegisterOutarFieldChp2Mask  = 0x400000
)

// GetChp2 Output 2 Chopper enable
func (r *registerOutarType) GetChp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldChp2Mask) != 0
}

// SetChp2 Output 2 Chopper enable
func (r *registerOutarType) SetChp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutarFieldChp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldChp2Mask)
	}
}

const (
	RegisterOutarFieldFault2Shift = 20
	RegisterOutarFieldFault2Mask  = 0x300000
)

// GetFault2 Output 2 Fault state
func (r *registerOutarType) GetFault2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldFault2Mask) >> RegisterOutarFieldFault2Shift)
}

// SetFault2 Output 2 Fault state
func (r *registerOutarType) SetFault2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldFault2Mask)|(uint32(value)<<RegisterOutarFieldFault2Shift))
}

const (
	RegisterOutarFieldIdles2Shift = 19
	RegisterOutarFieldIdles2Mask  = 0x80000
)

// GetIdles2 Output 2 Idle State
func (r *registerOutarType) GetIdles2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldIdles2Mask) != 0
}

// SetIdles2 Output 2 Idle State
func (r *registerOutarType) SetIdles2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutarFieldIdles2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldIdles2Mask)
	}
}

const (
	RegisterOutarFieldIdlem2Shift = 18
	RegisterOutarFieldIdlem2Mask  = 0x40000
)

// GetIdlem2 Output 2 Idle mode
func (r *registerOutarType) GetIdlem2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldIdlem2Mask) != 0
}

// SetIdlem2 Output 2 Idle mode
func (r *registerOutarType) SetIdlem2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutarFieldIdlem2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldIdlem2Mask)
	}
}

const (
	RegisterOutarFieldPol2Shift = 17
	RegisterOutarFieldPol2Mask  = 0x20000
)

// GetPol2 Output 2 polarity
func (r *registerOutarType) GetPol2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldPol2Mask) != 0
}

// SetPol2 Output 2 polarity
func (r *registerOutarType) SetPol2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutarFieldPol2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldPol2Mask)
	}
}

const (
	RegisterOutarFieldDlyprtShift = 10
	RegisterOutarFieldDlyprtMask  = 0x1c00
)

// GetDlyprt Delayed Protection
func (r *registerOutarType) GetDlyprt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldDlyprtMask) >> RegisterOutarFieldDlyprtShift)
}

// SetDlyprt Delayed Protection
func (r *registerOutarType) SetDlyprt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldDlyprtMask)|(uint32(value)<<RegisterOutarFieldDlyprtShift))
}

const (
	RegisterOutarFieldDlyprtenShift = 9
	RegisterOutarFieldDlyprtenMask  = 0x200
)

// GetDlyprten Delayed Protection Enable
func (r *registerOutarType) GetDlyprten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldDlyprtenMask) != 0
}

// SetDlyprten Delayed Protection Enable
func (r *registerOutarType) SetDlyprten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutarFieldDlyprtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldDlyprtenMask)
	}
}

const (
	RegisterOutarFieldDtenShift = 8
	RegisterOutarFieldDtenMask  = 0x100
)

// GetDten Deadtime enable
func (r *registerOutarType) GetDten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldDtenMask) != 0
}

// SetDten Deadtime enable
func (r *registerOutarType) SetDten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutarFieldDtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldDtenMask)
	}
}

const (
	RegisterOutarFieldDidl1Shift = 7
	RegisterOutarFieldDidl1Mask  = 0x80
)

// GetDidl1 Output 1 Deadtime upon burst mode Idle entry
func (r *registerOutarType) GetDidl1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldDidl1Mask) != 0
}

// SetDidl1 Output 1 Deadtime upon burst mode Idle entry
func (r *registerOutarType) SetDidl1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutarFieldDidl1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldDidl1Mask)
	}
}

const (
	RegisterOutarFieldChp1Shift = 6
	RegisterOutarFieldChp1Mask  = 0x40
)

// GetChp1 Output 1 Chopper enable
func (r *registerOutarType) GetChp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldChp1Mask) != 0
}

// SetChp1 Output 1 Chopper enable
func (r *registerOutarType) SetChp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutarFieldChp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldChp1Mask)
	}
}

const (
	RegisterOutarFieldFault1Shift = 4
	RegisterOutarFieldFault1Mask  = 0x30
)

// GetFault1 Output 1 Fault state
func (r *registerOutarType) GetFault1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldFault1Mask) >> RegisterOutarFieldFault1Shift)
}

// SetFault1 Output 1 Fault state
func (r *registerOutarType) SetFault1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldFault1Mask)|(uint32(value)<<RegisterOutarFieldFault1Shift))
}

const (
	RegisterOutarFieldIdles1Shift = 3
	RegisterOutarFieldIdles1Mask  = 0x8
)

// GetIdles1 Output 1 Idle State
func (r *registerOutarType) GetIdles1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldIdles1Mask) != 0
}

// SetIdles1 Output 1 Idle State
func (r *registerOutarType) SetIdles1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutarFieldIdles1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldIdles1Mask)
	}
}

const (
	RegisterOutarFieldIdlem1Shift = 2
	RegisterOutarFieldIdlem1Mask  = 0x4
)

// GetIdlem1 Output 1 Idle mode
func (r *registerOutarType) GetIdlem1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldIdlem1Mask) != 0
}

// SetIdlem1 Output 1 Idle mode
func (r *registerOutarType) SetIdlem1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutarFieldIdlem1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldIdlem1Mask)
	}
}

const (
	RegisterOutarFieldPol1Shift = 1
	RegisterOutarFieldPol1Mask  = 0x2
)

// GetPol1 Output 1 polarity
func (r *registerOutarType) GetPol1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOutarFieldPol1Mask) != 0
}

// SetPol1 Output 1 polarity
func (r *registerOutarType) SetPol1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOutarFieldPol1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOutarFieldPol1Mask)
	}
}

// registerFltarType Timerx Fault Register
type registerFltarType uint32

const (
	RegisterFltarFieldFltlckShift = 31
	RegisterFltarFieldFltlckMask  = 0x80000000
)

// GetFltlck Fault sources Lock
func (r *registerFltarType) GetFltlck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltarFieldFltlckMask) != 0
}

// SetFltlck Fault sources Lock
func (r *registerFltarType) SetFltlck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltarFieldFltlckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltarFieldFltlckMask)
	}
}

const (
	RegisterFltarFieldFlt5enShift = 4
	RegisterFltarFieldFlt5enMask  = 0x10
)

// GetFlt5en Fault 5 enable
func (r *registerFltarType) GetFlt5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltarFieldFlt5enMask) != 0
}

// SetFlt5en Fault 5 enable
func (r *registerFltarType) SetFlt5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltarFieldFlt5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltarFieldFlt5enMask)
	}
}

const (
	RegisterFltarFieldFlt4enShift = 3
	RegisterFltarFieldFlt4enMask  = 0x8
)

// GetFlt4en Fault 4 enable
func (r *registerFltarType) GetFlt4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltarFieldFlt4enMask) != 0
}

// SetFlt4en Fault 4 enable
func (r *registerFltarType) SetFlt4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltarFieldFlt4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltarFieldFlt4enMask)
	}
}

const (
	RegisterFltarFieldFlt3enShift = 2
	RegisterFltarFieldFlt3enMask  = 0x4
)

// GetFlt3en Fault 3 enable
func (r *registerFltarType) GetFlt3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltarFieldFlt3enMask) != 0
}

// SetFlt3en Fault 3 enable
func (r *registerFltarType) SetFlt3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltarFieldFlt3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltarFieldFlt3enMask)
	}
}

const (
	RegisterFltarFieldFlt2enShift = 1
	RegisterFltarFieldFlt2enMask  = 0x2
)

// GetFlt2en Fault 2 enable
func (r *registerFltarType) GetFlt2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltarFieldFlt2enMask) != 0
}

// SetFlt2en Fault 2 enable
func (r *registerFltarType) SetFlt2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltarFieldFlt2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltarFieldFlt2enMask)
	}
}

const (
	RegisterFltarFieldFlt1enShift = 0
	RegisterFltarFieldFlt1enMask  = 0x1
)

// GetFlt1en Fault 1 enable
func (r *registerFltarType) GetFlt1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFltarFieldFlt1enMask) != 0
}

// SetFlt1en Fault 1 enable
func (r *registerFltarType) SetFlt1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFltarFieldFlt1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFltarFieldFlt1enMask)
	}
}
