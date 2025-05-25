package fdcan

import (
	"unsafe"
	"volatile"
)

var (
	Fdcan1 = (*_fdcan)(unsafe.Pointer(uintptr(0x4000a000)))
	Fdcan2 = (*_fdcan)(unsafe.Pointer(uintptr(0x4000a400)))
)

type _fdcan struct {
	Fdcan_crel   registerFdcan_crelType
	Fdcan_endn   registerFdcan_endnType
	_            [4]byte
	Fdcan_dbtp   registerFdcan_dbtpType
	Fdcan_test   registerFdcan_testType
	Fdcan_rwd    registerFdcan_rwdType
	Fdcan_cccr   registerFdcan_cccrType
	Fdcan_nbtp   registerFdcan_nbtpType
	Fdcan_tscc   registerFdcan_tsccType
	Fdcan_tscv   registerFdcan_tscvType
	Fdcan_tocc   registerFdcan_toccType
	Fdcan_tocv   registerFdcan_tocvType
	_            [16]byte
	Fdcan_ecr    registerFdcan_ecrType
	Fdcan_psr    registerFdcan_psrType
	Fdcan_tdcr   registerFdcan_tdcrType
	_            [4]byte
	Fdcan_ir     registerFdcan_irType
	Fdcan_ie     registerFdcan_ieType
	Fdcan_ils    registerFdcan_ilsType
	Fdcan_ile    registerFdcan_ileType
	_            [32]byte
	Fdcan_gfc    registerFdcan_gfcType
	Fdcan_sidfc  registerFdcan_sidfcType
	Fdcan_xidfc  registerFdcan_xidfcType
	_            [4]byte
	Fdcan_xidam  registerFdcan_xidamType
	Fdcan_hpms   registerFdcan_hpmsType
	Fdcan_ndat1  registerFdcan_ndat1Type
	Fdcan_ndat2  registerFdcan_ndat2Type
	Fdcan_rxf0c  registerFdcan_rxf0cType
	Fdcan_rxf0s  registerFdcan_rxf0sType
	Fdcan_rxf0a  registerFdcan_rxf0aType
	Fdcan_rxbc   registerFdcan_rxbcType
	Fdcan_rxf1c  registerFdcan_rxf1cType
	Fdcan_rxf1s  registerFdcan_rxf1sType
	Fdcan_rxf1a  registerFdcan_rxf1aType
	Fdcan_rxesc  registerFdcan_rxescType
	Fdcan_txbc   registerFdcan_txbcType
	Fdcan_txfqs  registerFdcan_txfqsType
	Fdcan_txesc  registerFdcan_txescType
	Fdcan_txbrp  registerFdcan_txbrpType
	Fdcan_txbar  registerFdcan_txbarType
	Fdcan_txbcr  registerFdcan_txbcrType
	Fdcan_txbto  registerFdcan_txbtoType
	Fdcan_txbcf  registerFdcan_txbcfType
	Fdcan_txbtie registerFdcan_txbtieType
	Fdcan_txbcie registerFdcan_txbcieType
	_            [8]byte
	Fdcan_txefc  registerFdcan_txefcType
	Fdcan_txefs  registerFdcan_txefsType
	Fdcan_txefa  registerFdcan_txefaType
	_            [4]byte
	Fdcan_tttmc  registerFdcan_tttmcType
	Fdcan_ttrmc  registerFdcan_ttrmcType
	Fdcan_ttocf  registerFdcan_ttocfType
	Fdcan_ttmlm  registerFdcan_ttmlmType
	Fdcan_turcf  registerFdcan_turcfType
	Fdcan_ttocn  registerFdcan_ttocnType
	Can_ttgtp    registerCan_ttgtpType
	Fdcan_tttmk  registerFdcan_tttmkType
	Fdcan_ttir   registerFdcan_ttirType
	Fdcan_ttie   registerFdcan_ttieType
	Fdcan_ttils  registerFdcan_ttilsType
	Fdcan_ttost  registerFdcan_ttostType
	Fdcan_turna  registerFdcan_turnaType
	Fdcan_ttlgt  registerFdcan_ttlgtType
	Fdcan_ttctc  registerFdcan_ttctcType
	Fdcan_ttcpt  registerFdcan_ttcptType
	Fdcan_ttcsm  registerFdcan_ttcsmType
	_            [444]byte
	Fdcan_ttts   registerFdcan_tttsType
}

// registerFdcan_crelType FDCAN Core Release Register
type registerFdcan_crelType uint32

const (
	RegisterFdcan_crelFieldDayShift = 0
	RegisterFdcan_crelFieldDayMask  = 0xff
)

// GetDay Timestamp Day
func (r *registerFdcan_crelType) GetDay() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_crelFieldDayMask) >> RegisterFdcan_crelFieldDayShift)
}

// SetDay Timestamp Day
func (r *registerFdcan_crelType) SetDay(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_crelFieldDayMask)|(uint32(value)<<RegisterFdcan_crelFieldDayShift))
}

const (
	RegisterFdcan_crelFieldMonShift = 8
	RegisterFdcan_crelFieldMonMask  = 0xff00
)

// GetMon Timestamp Month
func (r *registerFdcan_crelType) GetMon() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_crelFieldMonMask) >> RegisterFdcan_crelFieldMonShift)
}

// SetMon Timestamp Month
func (r *registerFdcan_crelType) SetMon(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_crelFieldMonMask)|(uint32(value)<<RegisterFdcan_crelFieldMonShift))
}

const (
	RegisterFdcan_crelFieldYearShift = 16
	RegisterFdcan_crelFieldYearMask  = 0xf0000
)

// GetYear Timestamp Year
func (r *registerFdcan_crelType) GetYear() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_crelFieldYearMask) >> RegisterFdcan_crelFieldYearShift)
}

// SetYear Timestamp Year
func (r *registerFdcan_crelType) SetYear(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_crelFieldYearMask)|(uint32(value)<<RegisterFdcan_crelFieldYearShift))
}

const (
	RegisterFdcan_crelFieldSubstepShift = 20
	RegisterFdcan_crelFieldSubstepMask  = 0xf00000
)

// GetSubstep Sub-step of Core release
func (r *registerFdcan_crelType) GetSubstep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_crelFieldSubstepMask) >> RegisterFdcan_crelFieldSubstepShift)
}

// SetSubstep Sub-step of Core release
func (r *registerFdcan_crelType) SetSubstep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_crelFieldSubstepMask)|(uint32(value)<<RegisterFdcan_crelFieldSubstepShift))
}

const (
	RegisterFdcan_crelFieldStepShift = 24
	RegisterFdcan_crelFieldStepMask  = 0xf000000
)

// GetStep Step of Core release
func (r *registerFdcan_crelType) GetStep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_crelFieldStepMask) >> RegisterFdcan_crelFieldStepShift)
}

// SetStep Step of Core release
func (r *registerFdcan_crelType) SetStep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_crelFieldStepMask)|(uint32(value)<<RegisterFdcan_crelFieldStepShift))
}

const (
	RegisterFdcan_crelFieldRelShift = 28
	RegisterFdcan_crelFieldRelMask  = 0xf0000000
)

// GetRel Core release
func (r *registerFdcan_crelType) GetRel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_crelFieldRelMask) >> RegisterFdcan_crelFieldRelShift)
}

// SetRel Core release
func (r *registerFdcan_crelType) SetRel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_crelFieldRelMask)|(uint32(value)<<RegisterFdcan_crelFieldRelShift))
}

// registerFdcan_endnType FDCAN Core Release Register
type registerFdcan_endnType uint32

const (
	RegisterFdcan_endnFieldEtvShift = 0
	RegisterFdcan_endnFieldEtvMask  = 0xffffffff
)

// GetEtv Endiannes Test Value
func (r *registerFdcan_endnType) GetEtv() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_endnFieldEtvMask) >> RegisterFdcan_endnFieldEtvShift)
}

// SetEtv Endiannes Test Value
func (r *registerFdcan_endnType) SetEtv(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_endnFieldEtvMask)|(uint32(value)<<RegisterFdcan_endnFieldEtvShift))
}

// registerFdcan_dbtpType FDCAN Data Bit Timing and Prescaler Register
type registerFdcan_dbtpType uint32

const (
	RegisterFdcan_dbtpFieldDsjwShift = 0
	RegisterFdcan_dbtpFieldDsjwMask  = 0xf
)

// GetDsjw Synchronization Jump Width
func (r *registerFdcan_dbtpType) GetDsjw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_dbtpFieldDsjwMask) >> RegisterFdcan_dbtpFieldDsjwShift)
}

// SetDsjw Synchronization Jump Width
func (r *registerFdcan_dbtpType) SetDsjw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_dbtpFieldDsjwMask)|(uint32(value)<<RegisterFdcan_dbtpFieldDsjwShift))
}

const (
	RegisterFdcan_dbtpFieldDtseg2Shift = 4
	RegisterFdcan_dbtpFieldDtseg2Mask  = 0xf0
)

// GetDtseg2 Data time segment after sample point
func (r *registerFdcan_dbtpType) GetDtseg2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_dbtpFieldDtseg2Mask) >> RegisterFdcan_dbtpFieldDtseg2Shift)
}

// SetDtseg2 Data time segment after sample point
func (r *registerFdcan_dbtpType) SetDtseg2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_dbtpFieldDtseg2Mask)|(uint32(value)<<RegisterFdcan_dbtpFieldDtseg2Shift))
}

const (
	RegisterFdcan_dbtpFieldDtseg1Shift = 8
	RegisterFdcan_dbtpFieldDtseg1Mask  = 0x1f00
)

// GetDtseg1 Data time segment after sample point
func (r *registerFdcan_dbtpType) GetDtseg1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_dbtpFieldDtseg1Mask) >> RegisterFdcan_dbtpFieldDtseg1Shift)
}

// SetDtseg1 Data time segment after sample point
func (r *registerFdcan_dbtpType) SetDtseg1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_dbtpFieldDtseg1Mask)|(uint32(value)<<RegisterFdcan_dbtpFieldDtseg1Shift))
}

const (
	RegisterFdcan_dbtpFieldDbrpShift = 16
	RegisterFdcan_dbtpFieldDbrpMask  = 0x1f0000
)

// GetDbrp Data BIt Rate Prescaler
func (r *registerFdcan_dbtpType) GetDbrp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_dbtpFieldDbrpMask) >> RegisterFdcan_dbtpFieldDbrpShift)
}

// SetDbrp Data BIt Rate Prescaler
func (r *registerFdcan_dbtpType) SetDbrp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_dbtpFieldDbrpMask)|(uint32(value)<<RegisterFdcan_dbtpFieldDbrpShift))
}

const (
	RegisterFdcan_dbtpFieldTdcShift = 23
	RegisterFdcan_dbtpFieldTdcMask  = 0x800000
)

// GetTdc Transceiver Delay Compensation
func (r *registerFdcan_dbtpType) GetTdc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_dbtpFieldTdcMask) != 0
}

// SetTdc Transceiver Delay Compensation
func (r *registerFdcan_dbtpType) SetTdc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_dbtpFieldTdcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_dbtpFieldTdcMask)
	}
}

// registerFdcan_testType FDCAN Test Register
type registerFdcan_testType uint32

const (
	RegisterFdcan_testFieldLbckShift = 4
	RegisterFdcan_testFieldLbckMask  = 0x10
)

// GetLbck Loop Back mode
func (r *registerFdcan_testType) GetLbck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_testFieldLbckMask) != 0
}

// SetLbck Loop Back mode
func (r *registerFdcan_testType) SetLbck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_testFieldLbckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_testFieldLbckMask)
	}
}

const (
	RegisterFdcan_testFieldTxShift = 5
	RegisterFdcan_testFieldTxMask  = 0x60
)

// GetTx Loop Back mode
func (r *registerFdcan_testType) GetTx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_testFieldTxMask) >> RegisterFdcan_testFieldTxShift)
}

// SetTx Loop Back mode
func (r *registerFdcan_testType) SetTx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_testFieldTxMask)|(uint32(value)<<RegisterFdcan_testFieldTxShift))
}

const (
	RegisterFdcan_testFieldRxShift = 7
	RegisterFdcan_testFieldRxMask  = 0x80
)

// GetRx Control of Transmit Pin
func (r *registerFdcan_testType) GetRx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_testFieldRxMask) != 0
}

// SetRx Control of Transmit Pin
func (r *registerFdcan_testType) SetRx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_testFieldRxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_testFieldRxMask)
	}
}

// registerFdcan_rwdType FDCAN RAM Watchdog Register
type registerFdcan_rwdType uint32

const (
	RegisterFdcan_rwdFieldWdcShift = 0
	RegisterFdcan_rwdFieldWdcMask  = 0xff
)

// GetWdc Watchdog configuration
func (r *registerFdcan_rwdType) GetWdc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rwdFieldWdcMask) >> RegisterFdcan_rwdFieldWdcShift)
}

// SetWdc Watchdog configuration
func (r *registerFdcan_rwdType) SetWdc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rwdFieldWdcMask)|(uint32(value)<<RegisterFdcan_rwdFieldWdcShift))
}

const (
	RegisterFdcan_rwdFieldWdvShift = 8
	RegisterFdcan_rwdFieldWdvMask  = 0xff00
)

// GetWdv Watchdog value
func (r *registerFdcan_rwdType) GetWdv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rwdFieldWdvMask) >> RegisterFdcan_rwdFieldWdvShift)
}

// SetWdv Watchdog value
func (r *registerFdcan_rwdType) SetWdv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rwdFieldWdvMask)|(uint32(value)<<RegisterFdcan_rwdFieldWdvShift))
}

// registerFdcan_cccrType FDCAN CC Control Register
type registerFdcan_cccrType uint32

const (
	RegisterFdcan_cccrFieldInitShift = 0
	RegisterFdcan_cccrFieldInitMask  = 0x1
)

// GetInit Initialization
func (r *registerFdcan_cccrType) GetInit() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldInitMask) != 0
}

// SetInit Initialization
func (r *registerFdcan_cccrType) SetInit(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldInitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldInitMask)
	}
}

const (
	RegisterFdcan_cccrFieldCceShift = 1
	RegisterFdcan_cccrFieldCceMask  = 0x2
)

// GetCce Configuration Change Enable
func (r *registerFdcan_cccrType) GetCce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldCceMask) != 0
}

// SetCce Configuration Change Enable
func (r *registerFdcan_cccrType) SetCce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldCceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldCceMask)
	}
}

const (
	RegisterFdcan_cccrFieldAsmShift = 2
	RegisterFdcan_cccrFieldAsmMask  = 0x4
)

// GetAsm ASM Restricted Operation Mode
func (r *registerFdcan_cccrType) GetAsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldAsmMask) != 0
}

// SetAsm ASM Restricted Operation Mode
func (r *registerFdcan_cccrType) SetAsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldAsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldAsmMask)
	}
}

const (
	RegisterFdcan_cccrFieldCsaShift = 3
	RegisterFdcan_cccrFieldCsaMask  = 0x8
)

// GetCsa Clock Stop Acknowledge
func (r *registerFdcan_cccrType) GetCsa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldCsaMask) != 0
}

// SetCsa Clock Stop Acknowledge
func (r *registerFdcan_cccrType) SetCsa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldCsaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldCsaMask)
	}
}

const (
	RegisterFdcan_cccrFieldCsrShift = 4
	RegisterFdcan_cccrFieldCsrMask  = 0x10
)

// GetCsr Clock Stop Request
func (r *registerFdcan_cccrType) GetCsr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldCsrMask) != 0
}

// SetCsr Clock Stop Request
func (r *registerFdcan_cccrType) SetCsr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldCsrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldCsrMask)
	}
}

const (
	RegisterFdcan_cccrFieldMonShift = 5
	RegisterFdcan_cccrFieldMonMask  = 0x20
)

// GetMon Bus Monitoring Mode
func (r *registerFdcan_cccrType) GetMon() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldMonMask) != 0
}

// SetMon Bus Monitoring Mode
func (r *registerFdcan_cccrType) SetMon(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldMonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldMonMask)
	}
}

const (
	RegisterFdcan_cccrFieldDarShift = 6
	RegisterFdcan_cccrFieldDarMask  = 0x40
)

// GetDar Disable Automatic Retransmission
func (r *registerFdcan_cccrType) GetDar() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldDarMask) != 0
}

// SetDar Disable Automatic Retransmission
func (r *registerFdcan_cccrType) SetDar(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldDarMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldDarMask)
	}
}

const (
	RegisterFdcan_cccrFieldTestShift = 7
	RegisterFdcan_cccrFieldTestMask  = 0x80
)

// GetTest Test Mode Enable
func (r *registerFdcan_cccrType) GetTest() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldTestMask) != 0
}

// SetTest Test Mode Enable
func (r *registerFdcan_cccrType) SetTest(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldTestMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldTestMask)
	}
}

const (
	RegisterFdcan_cccrFieldFdoeShift = 8
	RegisterFdcan_cccrFieldFdoeMask  = 0x100
)

// GetFdoe FD Operation Enable
func (r *registerFdcan_cccrType) GetFdoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldFdoeMask) != 0
}

// SetFdoe FD Operation Enable
func (r *registerFdcan_cccrType) SetFdoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldFdoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldFdoeMask)
	}
}

const (
	RegisterFdcan_cccrFieldBseShift = 9
	RegisterFdcan_cccrFieldBseMask  = 0x200
)

// GetBse FDCAN Bit Rate Switching
func (r *registerFdcan_cccrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldBseMask) != 0
}

// SetBse FDCAN Bit Rate Switching
func (r *registerFdcan_cccrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldBseMask)
	}
}

const (
	RegisterFdcan_cccrFieldPxhdShift = 12
	RegisterFdcan_cccrFieldPxhdMask  = 0x1000
)

// GetPxhd Protocol Exception Handling Disable
func (r *registerFdcan_cccrType) GetPxhd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldPxhdMask) != 0
}

// SetPxhd Protocol Exception Handling Disable
func (r *registerFdcan_cccrType) SetPxhd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldPxhdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldPxhdMask)
	}
}

const (
	RegisterFdcan_cccrFieldEfbiShift = 13
	RegisterFdcan_cccrFieldEfbiMask  = 0x2000
)

// GetEfbi Edge Filtering during Bus Integration
func (r *registerFdcan_cccrType) GetEfbi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldEfbiMask) != 0
}

// SetEfbi Edge Filtering during Bus Integration
func (r *registerFdcan_cccrType) SetEfbi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldEfbiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldEfbiMask)
	}
}

const (
	RegisterFdcan_cccrFieldTxpShift = 14
	RegisterFdcan_cccrFieldTxpMask  = 0x4000
)

// GetTxp TXP
func (r *registerFdcan_cccrType) GetTxp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldTxpMask) != 0
}

// SetTxp TXP
func (r *registerFdcan_cccrType) SetTxp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldTxpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldTxpMask)
	}
}

const (
	RegisterFdcan_cccrFieldNisoShift = 15
	RegisterFdcan_cccrFieldNisoMask  = 0x8000
)

// GetNiso Non ISO Operation
func (r *registerFdcan_cccrType) GetNiso() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_cccrFieldNisoMask) != 0
}

// SetNiso Non ISO Operation
func (r *registerFdcan_cccrType) SetNiso(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_cccrFieldNisoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_cccrFieldNisoMask)
	}
}

// registerFdcan_nbtpType FDCAN Nominal Bit Timing and Prescaler Register
type registerFdcan_nbtpType uint32

const (
	RegisterFdcan_nbtpFieldTseg2Shift = 0
	RegisterFdcan_nbtpFieldTseg2Mask  = 0x7f
)

// GetTseg2 Nominal Time segment after sample point
func (r *registerFdcan_nbtpType) GetTseg2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_nbtpFieldTseg2Mask) >> RegisterFdcan_nbtpFieldTseg2Shift)
}

// SetTseg2 Nominal Time segment after sample point
func (r *registerFdcan_nbtpType) SetTseg2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_nbtpFieldTseg2Mask)|(uint32(value)<<RegisterFdcan_nbtpFieldTseg2Shift))
}

const (
	RegisterFdcan_nbtpFieldNtseg1Shift = 8
	RegisterFdcan_nbtpFieldNtseg1Mask  = 0xff00
)

// GetNtseg1 Nominal Time segment before sample point
func (r *registerFdcan_nbtpType) GetNtseg1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_nbtpFieldNtseg1Mask) >> RegisterFdcan_nbtpFieldNtseg1Shift)
}

// SetNtseg1 Nominal Time segment before sample point
func (r *registerFdcan_nbtpType) SetNtseg1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_nbtpFieldNtseg1Mask)|(uint32(value)<<RegisterFdcan_nbtpFieldNtseg1Shift))
}

const (
	RegisterFdcan_nbtpFieldNbrpShift = 16
	RegisterFdcan_nbtpFieldNbrpMask  = 0x1ff0000
)

// GetNbrp Bit Rate Prescaler
func (r *registerFdcan_nbtpType) GetNbrp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_nbtpFieldNbrpMask) >> RegisterFdcan_nbtpFieldNbrpShift)
}

// SetNbrp Bit Rate Prescaler
func (r *registerFdcan_nbtpType) SetNbrp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_nbtpFieldNbrpMask)|(uint32(value)<<RegisterFdcan_nbtpFieldNbrpShift))
}

const (
	RegisterFdcan_nbtpFieldNsjwShift = 25
	RegisterFdcan_nbtpFieldNsjwMask  = 0xfe000000
)

// GetNsjw NSJW: Nominal (Re)Synchronization Jump Width
func (r *registerFdcan_nbtpType) GetNsjw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_nbtpFieldNsjwMask) >> RegisterFdcan_nbtpFieldNsjwShift)
}

// SetNsjw NSJW: Nominal (Re)Synchronization Jump Width
func (r *registerFdcan_nbtpType) SetNsjw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_nbtpFieldNsjwMask)|(uint32(value)<<RegisterFdcan_nbtpFieldNsjwShift))
}

// registerFdcan_tsccType FDCAN Timestamp Counter Configuration Register
type registerFdcan_tsccType uint32

const (
	RegisterFdcan_tsccFieldTssShift = 0
	RegisterFdcan_tsccFieldTssMask  = 0x3
)

// GetTss Timestamp Select
func (r *registerFdcan_tsccType) GetTss() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tsccFieldTssMask) >> RegisterFdcan_tsccFieldTssShift)
}

// SetTss Timestamp Select
func (r *registerFdcan_tsccType) SetTss(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tsccFieldTssMask)|(uint32(value)<<RegisterFdcan_tsccFieldTssShift))
}

const (
	RegisterFdcan_tsccFieldTcpShift = 16
	RegisterFdcan_tsccFieldTcpMask  = 0xf0000
)

// GetTcp Timestamp Counter Prescaler
func (r *registerFdcan_tsccType) GetTcp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tsccFieldTcpMask) >> RegisterFdcan_tsccFieldTcpShift)
}

// SetTcp Timestamp Counter Prescaler
func (r *registerFdcan_tsccType) SetTcp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tsccFieldTcpMask)|(uint32(value)<<RegisterFdcan_tsccFieldTcpShift))
}

// registerFdcan_tscvType FDCAN Timestamp Counter Value Register
type registerFdcan_tscvType uint32

const (
	RegisterFdcan_tscvFieldTscShift = 0
	RegisterFdcan_tscvFieldTscMask  = 0xffff
)

// GetTsc Timestamp Counter
func (r *registerFdcan_tscvType) GetTsc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tscvFieldTscMask) >> RegisterFdcan_tscvFieldTscShift)
}

// SetTsc Timestamp Counter
func (r *registerFdcan_tscvType) SetTsc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tscvFieldTscMask)|(uint32(value)<<RegisterFdcan_tscvFieldTscShift))
}

// registerFdcan_toccType FDCAN Timeout Counter Configuration Register
type registerFdcan_toccType uint32

const (
	RegisterFdcan_toccFieldEtocShift = 0
	RegisterFdcan_toccFieldEtocMask  = 0x1
)

// GetEtoc Enable Timeout Counter
func (r *registerFdcan_toccType) GetEtoc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_toccFieldEtocMask) != 0
}

// SetEtoc Enable Timeout Counter
func (r *registerFdcan_toccType) SetEtoc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_toccFieldEtocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_toccFieldEtocMask)
	}
}

const (
	RegisterFdcan_toccFieldTosShift = 1
	RegisterFdcan_toccFieldTosMask  = 0x6
)

// GetTos Timeout Select
func (r *registerFdcan_toccType) GetTos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_toccFieldTosMask) >> RegisterFdcan_toccFieldTosShift)
}

// SetTos Timeout Select
func (r *registerFdcan_toccType) SetTos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_toccFieldTosMask)|(uint32(value)<<RegisterFdcan_toccFieldTosShift))
}

const (
	RegisterFdcan_toccFieldTopShift = 16
	RegisterFdcan_toccFieldTopMask  = 0xffff0000
)

// GetTop Timeout Period
func (r *registerFdcan_toccType) GetTop() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_toccFieldTopMask) >> RegisterFdcan_toccFieldTopShift)
}

// SetTop Timeout Period
func (r *registerFdcan_toccType) SetTop(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_toccFieldTopMask)|(uint32(value)<<RegisterFdcan_toccFieldTopShift))
}

// registerFdcan_tocvType FDCAN Timeout Counter Value Register
type registerFdcan_tocvType uint32

const (
	RegisterFdcan_tocvFieldTocShift = 0
	RegisterFdcan_tocvFieldTocMask  = 0xffff
)

// GetToc Timeout Counter
func (r *registerFdcan_tocvType) GetToc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tocvFieldTocMask) >> RegisterFdcan_tocvFieldTocShift)
}

// SetToc Timeout Counter
func (r *registerFdcan_tocvType) SetToc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tocvFieldTocMask)|(uint32(value)<<RegisterFdcan_tocvFieldTocShift))
}

// registerFdcan_ecrType FDCAN Error Counter Register
type registerFdcan_ecrType uint32

const (
	RegisterFdcan_ecrFieldTecShift = 0
	RegisterFdcan_ecrFieldTecMask  = 0xff
)

// GetTec Transmit Error Counter
func (r *registerFdcan_ecrType) GetTec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ecrFieldTecMask) >> RegisterFdcan_ecrFieldTecShift)
}

// SetTec Transmit Error Counter
func (r *registerFdcan_ecrType) SetTec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ecrFieldTecMask)|(uint32(value)<<RegisterFdcan_ecrFieldTecShift))
}

const (
	RegisterFdcan_ecrFieldTrecShift = 8
	RegisterFdcan_ecrFieldTrecMask  = 0x7f00
)

// GetTrec Receive Error Counter
func (r *registerFdcan_ecrType) GetTrec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ecrFieldTrecMask) >> RegisterFdcan_ecrFieldTrecShift)
}

// SetTrec Receive Error Counter
func (r *registerFdcan_ecrType) SetTrec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ecrFieldTrecMask)|(uint32(value)<<RegisterFdcan_ecrFieldTrecShift))
}

const (
	RegisterFdcan_ecrFieldRpShift = 15
	RegisterFdcan_ecrFieldRpMask  = 0x8000
)

// GetRp Receive Error Passive
func (r *registerFdcan_ecrType) GetRp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ecrFieldRpMask) != 0
}

// SetRp Receive Error Passive
func (r *registerFdcan_ecrType) SetRp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ecrFieldRpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ecrFieldRpMask)
	}
}

const (
	RegisterFdcan_ecrFieldCelShift = 16
	RegisterFdcan_ecrFieldCelMask  = 0xff0000
)

// GetCel AN Error Logging
func (r *registerFdcan_ecrType) GetCel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ecrFieldCelMask) >> RegisterFdcan_ecrFieldCelShift)
}

// SetCel AN Error Logging
func (r *registerFdcan_ecrType) SetCel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ecrFieldCelMask)|(uint32(value)<<RegisterFdcan_ecrFieldCelShift))
}

// registerFdcan_psrType FDCAN Protocol Status Register
type registerFdcan_psrType uint32

const (
	RegisterFdcan_psrFieldLecShift = 0
	RegisterFdcan_psrFieldLecMask  = 0x7
)

// GetLec Last Error Code
func (r *registerFdcan_psrType) GetLec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_psrFieldLecMask) >> RegisterFdcan_psrFieldLecShift)
}

// SetLec Last Error Code
func (r *registerFdcan_psrType) SetLec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_psrFieldLecMask)|(uint32(value)<<RegisterFdcan_psrFieldLecShift))
}

const (
	RegisterFdcan_psrFieldActShift = 3
	RegisterFdcan_psrFieldActMask  = 0x18
)

// GetAct Activity
func (r *registerFdcan_psrType) GetAct() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_psrFieldActMask) >> RegisterFdcan_psrFieldActShift)
}

// SetAct Activity
func (r *registerFdcan_psrType) SetAct(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_psrFieldActMask)|(uint32(value)<<RegisterFdcan_psrFieldActShift))
}

const (
	RegisterFdcan_psrFieldEpShift = 5
	RegisterFdcan_psrFieldEpMask  = 0x20
)

// GetEp Error Passive
func (r *registerFdcan_psrType) GetEp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_psrFieldEpMask) != 0
}

// SetEp Error Passive
func (r *registerFdcan_psrType) SetEp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_psrFieldEpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_psrFieldEpMask)
	}
}

const (
	RegisterFdcan_psrFieldEwShift = 6
	RegisterFdcan_psrFieldEwMask  = 0x40
)

// GetEw Warning Status
func (r *registerFdcan_psrType) GetEw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_psrFieldEwMask) != 0
}

// SetEw Warning Status
func (r *registerFdcan_psrType) SetEw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_psrFieldEwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_psrFieldEwMask)
	}
}

const (
	RegisterFdcan_psrFieldBoShift = 7
	RegisterFdcan_psrFieldBoMask  = 0x80
)

// GetBo Bus_Off Status
func (r *registerFdcan_psrType) GetBo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_psrFieldBoMask) != 0
}

// SetBo Bus_Off Status
func (r *registerFdcan_psrType) SetBo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_psrFieldBoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_psrFieldBoMask)
	}
}

const (
	RegisterFdcan_psrFieldDlecShift = 8
	RegisterFdcan_psrFieldDlecMask  = 0x700
)

// GetDlec Data Last Error Code
func (r *registerFdcan_psrType) GetDlec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_psrFieldDlecMask) >> RegisterFdcan_psrFieldDlecShift)
}

// SetDlec Data Last Error Code
func (r *registerFdcan_psrType) SetDlec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_psrFieldDlecMask)|(uint32(value)<<RegisterFdcan_psrFieldDlecShift))
}

const (
	RegisterFdcan_psrFieldResiShift = 11
	RegisterFdcan_psrFieldResiMask  = 0x800
)

// GetResi ESI flag of last received FDCAN Message
func (r *registerFdcan_psrType) GetResi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_psrFieldResiMask) != 0
}

// SetResi ESI flag of last received FDCAN Message
func (r *registerFdcan_psrType) SetResi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_psrFieldResiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_psrFieldResiMask)
	}
}

const (
	RegisterFdcan_psrFieldRbrsShift = 12
	RegisterFdcan_psrFieldRbrsMask  = 0x1000
)

// GetRbrs BRS flag of last received FDCAN Message
func (r *registerFdcan_psrType) GetRbrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_psrFieldRbrsMask) != 0
}

// SetRbrs BRS flag of last received FDCAN Message
func (r *registerFdcan_psrType) SetRbrs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_psrFieldRbrsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_psrFieldRbrsMask)
	}
}

const (
	RegisterFdcan_psrFieldRedlShift = 13
	RegisterFdcan_psrFieldRedlMask  = 0x2000
)

// GetRedl Received FDCAN Message
func (r *registerFdcan_psrType) GetRedl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_psrFieldRedlMask) != 0
}

// SetRedl Received FDCAN Message
func (r *registerFdcan_psrType) SetRedl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_psrFieldRedlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_psrFieldRedlMask)
	}
}

const (
	RegisterFdcan_psrFieldPxeShift = 14
	RegisterFdcan_psrFieldPxeMask  = 0x4000
)

// GetPxe Protocol Exception Event
func (r *registerFdcan_psrType) GetPxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_psrFieldPxeMask) != 0
}

// SetPxe Protocol Exception Event
func (r *registerFdcan_psrType) SetPxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_psrFieldPxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_psrFieldPxeMask)
	}
}

const (
	RegisterFdcan_psrFieldTdcvShift = 16
	RegisterFdcan_psrFieldTdcvMask  = 0x7f0000
)

// GetTdcv Transmitter Delay Compensation Value
func (r *registerFdcan_psrType) GetTdcv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_psrFieldTdcvMask) >> RegisterFdcan_psrFieldTdcvShift)
}

// SetTdcv Transmitter Delay Compensation Value
func (r *registerFdcan_psrType) SetTdcv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_psrFieldTdcvMask)|(uint32(value)<<RegisterFdcan_psrFieldTdcvShift))
}

// registerFdcan_tdcrType FDCAN Transmitter Delay Compensation Register
type registerFdcan_tdcrType uint32

const (
	RegisterFdcan_tdcrFieldTdcfShift = 0
	RegisterFdcan_tdcrFieldTdcfMask  = 0x7f
)

// GetTdcf Transmitter Delay Compensation Filter Window Length
func (r *registerFdcan_tdcrType) GetTdcf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tdcrFieldTdcfMask) >> RegisterFdcan_tdcrFieldTdcfShift)
}

// SetTdcf Transmitter Delay Compensation Filter Window Length
func (r *registerFdcan_tdcrType) SetTdcf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tdcrFieldTdcfMask)|(uint32(value)<<RegisterFdcan_tdcrFieldTdcfShift))
}

const (
	RegisterFdcan_tdcrFieldTdcoShift = 8
	RegisterFdcan_tdcrFieldTdcoMask  = 0x7f00
)

// GetTdco Transmitter Delay Compensation Offset
func (r *registerFdcan_tdcrType) GetTdco() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tdcrFieldTdcoMask) >> RegisterFdcan_tdcrFieldTdcoShift)
}

// SetTdco Transmitter Delay Compensation Offset
func (r *registerFdcan_tdcrType) SetTdco(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tdcrFieldTdcoMask)|(uint32(value)<<RegisterFdcan_tdcrFieldTdcoShift))
}

// registerFdcan_irType FDCAN Interrupt Register
type registerFdcan_irType uint32

const (
	RegisterFdcan_irFieldRf0nShift = 0
	RegisterFdcan_irFieldRf0nMask  = 0x1
)

// GetRf0n Rx FIFO 0 New Message
func (r *registerFdcan_irType) GetRf0n() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldRf0nMask) != 0
}

// SetRf0n Rx FIFO 0 New Message
func (r *registerFdcan_irType) SetRf0n(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldRf0nMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldRf0nMask)
	}
}

const (
	RegisterFdcan_irFieldRf0wShift = 1
	RegisterFdcan_irFieldRf0wMask  = 0x2
)

// GetRf0w Rx FIFO 0 Full
func (r *registerFdcan_irType) GetRf0w() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldRf0wMask) != 0
}

// SetRf0w Rx FIFO 0 Full
func (r *registerFdcan_irType) SetRf0w(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldRf0wMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldRf0wMask)
	}
}

const (
	RegisterFdcan_irFieldRf0fShift = 2
	RegisterFdcan_irFieldRf0fMask  = 0x4
)

// GetRf0f Rx FIFO 0 Full
func (r *registerFdcan_irType) GetRf0f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldRf0fMask) != 0
}

// SetRf0f Rx FIFO 0 Full
func (r *registerFdcan_irType) SetRf0f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldRf0fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldRf0fMask)
	}
}

const (
	RegisterFdcan_irFieldRf0lShift = 3
	RegisterFdcan_irFieldRf0lMask  = 0x8
)

// GetRf0l Rx FIFO 0 Message Lost
func (r *registerFdcan_irType) GetRf0l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldRf0lMask) != 0
}

// SetRf0l Rx FIFO 0 Message Lost
func (r *registerFdcan_irType) SetRf0l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldRf0lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldRf0lMask)
	}
}

const (
	RegisterFdcan_irFieldRf1nShift = 4
	RegisterFdcan_irFieldRf1nMask  = 0x10
)

// GetRf1n Rx FIFO 1 New Message
func (r *registerFdcan_irType) GetRf1n() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldRf1nMask) != 0
}

// SetRf1n Rx FIFO 1 New Message
func (r *registerFdcan_irType) SetRf1n(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldRf1nMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldRf1nMask)
	}
}

const (
	RegisterFdcan_irFieldRf1wShift = 5
	RegisterFdcan_irFieldRf1wMask  = 0x20
)

// GetRf1w Rx FIFO 1 Watermark Reached
func (r *registerFdcan_irType) GetRf1w() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldRf1wMask) != 0
}

// SetRf1w Rx FIFO 1 Watermark Reached
func (r *registerFdcan_irType) SetRf1w(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldRf1wMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldRf1wMask)
	}
}

const (
	RegisterFdcan_irFieldRf1fShift = 6
	RegisterFdcan_irFieldRf1fMask  = 0x40
)

// GetRf1f Rx FIFO 1 Watermark Reached
func (r *registerFdcan_irType) GetRf1f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldRf1fMask) != 0
}

// SetRf1f Rx FIFO 1 Watermark Reached
func (r *registerFdcan_irType) SetRf1f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldRf1fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldRf1fMask)
	}
}

const (
	RegisterFdcan_irFieldRf1lShift = 7
	RegisterFdcan_irFieldRf1lMask  = 0x80
)

// GetRf1l Rx FIFO 1 Message Lost
func (r *registerFdcan_irType) GetRf1l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldRf1lMask) != 0
}

// SetRf1l Rx FIFO 1 Message Lost
func (r *registerFdcan_irType) SetRf1l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldRf1lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldRf1lMask)
	}
}

const (
	RegisterFdcan_irFieldHpmShift = 8
	RegisterFdcan_irFieldHpmMask  = 0x100
)

// GetHpm High Priority Message
func (r *registerFdcan_irType) GetHpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldHpmMask) != 0
}

// SetHpm High Priority Message
func (r *registerFdcan_irType) SetHpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldHpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldHpmMask)
	}
}

const (
	RegisterFdcan_irFieldTcShift = 9
	RegisterFdcan_irFieldTcMask  = 0x200
)

// GetTc Transmission Completed
func (r *registerFdcan_irType) GetTc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldTcMask) != 0
}

// SetTc Transmission Completed
func (r *registerFdcan_irType) SetTc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldTcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldTcMask)
	}
}

const (
	RegisterFdcan_irFieldTcfShift = 10
	RegisterFdcan_irFieldTcfMask  = 0x400
)

// GetTcf Transmission Cancellation Finished
func (r *registerFdcan_irType) GetTcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldTcfMask) != 0
}

// SetTcf Transmission Cancellation Finished
func (r *registerFdcan_irType) SetTcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldTcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldTcfMask)
	}
}

const (
	RegisterFdcan_irFieldTefShift = 11
	RegisterFdcan_irFieldTefMask  = 0x800
)

// GetTef Tx FIFO Empty
func (r *registerFdcan_irType) GetTef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldTefMask) != 0
}

// SetTef Tx FIFO Empty
func (r *registerFdcan_irType) SetTef(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldTefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldTefMask)
	}
}

const (
	RegisterFdcan_irFieldTefnShift = 12
	RegisterFdcan_irFieldTefnMask  = 0x1000
)

// GetTefn Tx Event FIFO New Entry
func (r *registerFdcan_irType) GetTefn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldTefnMask) != 0
}

// SetTefn Tx Event FIFO New Entry
func (r *registerFdcan_irType) SetTefn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldTefnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldTefnMask)
	}
}

const (
	RegisterFdcan_irFieldTefwShift = 13
	RegisterFdcan_irFieldTefwMask  = 0x2000
)

// GetTefw Tx Event FIFO Watermark Reached
func (r *registerFdcan_irType) GetTefw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldTefwMask) != 0
}

// SetTefw Tx Event FIFO Watermark Reached
func (r *registerFdcan_irType) SetTefw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldTefwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldTefwMask)
	}
}

const (
	RegisterFdcan_irFieldTeffShift = 14
	RegisterFdcan_irFieldTeffMask  = 0x4000
)

// GetTeff Tx Event FIFO Full
func (r *registerFdcan_irType) GetTeff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldTeffMask) != 0
}

// SetTeff Tx Event FIFO Full
func (r *registerFdcan_irType) SetTeff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldTeffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldTeffMask)
	}
}

const (
	RegisterFdcan_irFieldTeflShift = 15
	RegisterFdcan_irFieldTeflMask  = 0x8000
)

// GetTefl Tx Event FIFO Element Lost
func (r *registerFdcan_irType) GetTefl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldTeflMask) != 0
}

// SetTefl Tx Event FIFO Element Lost
func (r *registerFdcan_irType) SetTefl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldTeflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldTeflMask)
	}
}

const (
	RegisterFdcan_irFieldTswShift = 16
	RegisterFdcan_irFieldTswMask  = 0x10000
)

// GetTsw Timestamp Wraparound
func (r *registerFdcan_irType) GetTsw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldTswMask) != 0
}

// SetTsw Timestamp Wraparound
func (r *registerFdcan_irType) SetTsw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldTswMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldTswMask)
	}
}

const (
	RegisterFdcan_irFieldMrafShift = 17
	RegisterFdcan_irFieldMrafMask  = 0x20000
)

// GetMraf Message RAM Access Failure
func (r *registerFdcan_irType) GetMraf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldMrafMask) != 0
}

// SetMraf Message RAM Access Failure
func (r *registerFdcan_irType) SetMraf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldMrafMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldMrafMask)
	}
}

const (
	RegisterFdcan_irFieldTooShift = 18
	RegisterFdcan_irFieldTooMask  = 0x40000
)

// GetToo Timeout Occurred
func (r *registerFdcan_irType) GetToo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldTooMask) != 0
}

// SetToo Timeout Occurred
func (r *registerFdcan_irType) SetToo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldTooMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldTooMask)
	}
}

const (
	RegisterFdcan_irFieldDrxShift = 19
	RegisterFdcan_irFieldDrxMask  = 0x80000
)

// GetDrx Message stored to Dedicated Rx Buffer
func (r *registerFdcan_irType) GetDrx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldDrxMask) != 0
}

// SetDrx Message stored to Dedicated Rx Buffer
func (r *registerFdcan_irType) SetDrx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldDrxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldDrxMask)
	}
}

const (
	RegisterFdcan_irFieldEloShift = 22
	RegisterFdcan_irFieldEloMask  = 0x400000
)

// GetElo Error Logging Overflow
func (r *registerFdcan_irType) GetElo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldEloMask) != 0
}

// SetElo Error Logging Overflow
func (r *registerFdcan_irType) SetElo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldEloMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldEloMask)
	}
}

const (
	RegisterFdcan_irFieldEpShift = 23
	RegisterFdcan_irFieldEpMask  = 0x800000
)

// GetEp Error Passive
func (r *registerFdcan_irType) GetEp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldEpMask) != 0
}

// SetEp Error Passive
func (r *registerFdcan_irType) SetEp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldEpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldEpMask)
	}
}

const (
	RegisterFdcan_irFieldEwShift = 24
	RegisterFdcan_irFieldEwMask  = 0x1000000
)

// GetEw Warning Status
func (r *registerFdcan_irType) GetEw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldEwMask) != 0
}

// SetEw Warning Status
func (r *registerFdcan_irType) SetEw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldEwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldEwMask)
	}
}

const (
	RegisterFdcan_irFieldBoShift = 25
	RegisterFdcan_irFieldBoMask  = 0x2000000
)

// GetBo Bus_Off Status
func (r *registerFdcan_irType) GetBo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldBoMask) != 0
}

// SetBo Bus_Off Status
func (r *registerFdcan_irType) SetBo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldBoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldBoMask)
	}
}

const (
	RegisterFdcan_irFieldWdiShift = 26
	RegisterFdcan_irFieldWdiMask  = 0x4000000
)

// GetWdi Watchdog Interrupt
func (r *registerFdcan_irType) GetWdi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldWdiMask) != 0
}

// SetWdi Watchdog Interrupt
func (r *registerFdcan_irType) SetWdi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldWdiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldWdiMask)
	}
}

const (
	RegisterFdcan_irFieldPeaShift = 27
	RegisterFdcan_irFieldPeaMask  = 0x8000000
)

// GetPea Protocol Error in Arbitration Phase (Nominal Bit Time is used)
func (r *registerFdcan_irType) GetPea() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldPeaMask) != 0
}

// SetPea Protocol Error in Arbitration Phase (Nominal Bit Time is used)
func (r *registerFdcan_irType) SetPea(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldPeaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldPeaMask)
	}
}

const (
	RegisterFdcan_irFieldPedShift = 28
	RegisterFdcan_irFieldPedMask  = 0x10000000
)

// GetPed Protocol Error in Data Phase (Data Bit Time is used)
func (r *registerFdcan_irType) GetPed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldPedMask) != 0
}

// SetPed Protocol Error in Data Phase (Data Bit Time is used)
func (r *registerFdcan_irType) SetPed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldPedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldPedMask)
	}
}

const (
	RegisterFdcan_irFieldAraShift = 29
	RegisterFdcan_irFieldAraMask  = 0x20000000
)

// GetAra Access to Reserved Address
func (r *registerFdcan_irType) GetAra() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_irFieldAraMask) != 0
}

// SetAra Access to Reserved Address
func (r *registerFdcan_irType) SetAra(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_irFieldAraMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_irFieldAraMask)
	}
}

// registerFdcan_ieType FDCAN Interrupt Enable Register
type registerFdcan_ieType uint32

const (
	RegisterFdcan_ieFieldRf0neShift = 0
	RegisterFdcan_ieFieldRf0neMask  = 0x1
)

// GetRf0ne Rx FIFO 0 New Message Enable
func (r *registerFdcan_ieType) GetRf0ne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldRf0neMask) != 0
}

// SetRf0ne Rx FIFO 0 New Message Enable
func (r *registerFdcan_ieType) SetRf0ne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldRf0neMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldRf0neMask)
	}
}

const (
	RegisterFdcan_ieFieldRf0weShift = 1
	RegisterFdcan_ieFieldRf0weMask  = 0x2
)

// GetRf0we Rx FIFO 0 Full Enable
func (r *registerFdcan_ieType) GetRf0we() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldRf0weMask) != 0
}

// SetRf0we Rx FIFO 0 Full Enable
func (r *registerFdcan_ieType) SetRf0we(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldRf0weMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldRf0weMask)
	}
}

const (
	RegisterFdcan_ieFieldRf0feShift = 2
	RegisterFdcan_ieFieldRf0feMask  = 0x4
)

// GetRf0fe Rx FIFO 0 Full Enable
func (r *registerFdcan_ieType) GetRf0fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldRf0feMask) != 0
}

// SetRf0fe Rx FIFO 0 Full Enable
func (r *registerFdcan_ieType) SetRf0fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldRf0feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldRf0feMask)
	}
}

const (
	RegisterFdcan_ieFieldRf0leShift = 3
	RegisterFdcan_ieFieldRf0leMask  = 0x8
)

// GetRf0le Rx FIFO 0 Message Lost Enable
func (r *registerFdcan_ieType) GetRf0le() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldRf0leMask) != 0
}

// SetRf0le Rx FIFO 0 Message Lost Enable
func (r *registerFdcan_ieType) SetRf0le(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldRf0leMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldRf0leMask)
	}
}

const (
	RegisterFdcan_ieFieldRf1neShift = 4
	RegisterFdcan_ieFieldRf1neMask  = 0x10
)

// GetRf1ne Rx FIFO 1 New Message Enable
func (r *registerFdcan_ieType) GetRf1ne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldRf1neMask) != 0
}

// SetRf1ne Rx FIFO 1 New Message Enable
func (r *registerFdcan_ieType) SetRf1ne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldRf1neMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldRf1neMask)
	}
}

const (
	RegisterFdcan_ieFieldRf1weShift = 5
	RegisterFdcan_ieFieldRf1weMask  = 0x20
)

// GetRf1we Rx FIFO 1 Watermark Reached Enable
func (r *registerFdcan_ieType) GetRf1we() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldRf1weMask) != 0
}

// SetRf1we Rx FIFO 1 Watermark Reached Enable
func (r *registerFdcan_ieType) SetRf1we(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldRf1weMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldRf1weMask)
	}
}

const (
	RegisterFdcan_ieFieldRf1feShift = 6
	RegisterFdcan_ieFieldRf1feMask  = 0x40
)

// GetRf1fe Rx FIFO 1 Watermark Reached Enable
func (r *registerFdcan_ieType) GetRf1fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldRf1feMask) != 0
}

// SetRf1fe Rx FIFO 1 Watermark Reached Enable
func (r *registerFdcan_ieType) SetRf1fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldRf1feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldRf1feMask)
	}
}

const (
	RegisterFdcan_ieFieldRf1leShift = 7
	RegisterFdcan_ieFieldRf1leMask  = 0x80
)

// GetRf1le Rx FIFO 1 Message Lost Enable
func (r *registerFdcan_ieType) GetRf1le() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldRf1leMask) != 0
}

// SetRf1le Rx FIFO 1 Message Lost Enable
func (r *registerFdcan_ieType) SetRf1le(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldRf1leMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldRf1leMask)
	}
}

const (
	RegisterFdcan_ieFieldHpmeShift = 8
	RegisterFdcan_ieFieldHpmeMask  = 0x100
)

// GetHpme High Priority Message Enable
func (r *registerFdcan_ieType) GetHpme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldHpmeMask) != 0
}

// SetHpme High Priority Message Enable
func (r *registerFdcan_ieType) SetHpme(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldHpmeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldHpmeMask)
	}
}

const (
	RegisterFdcan_ieFieldTceShift = 9
	RegisterFdcan_ieFieldTceMask  = 0x200
)

// GetTce Transmission Completed Enable
func (r *registerFdcan_ieType) GetTce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldTceMask) != 0
}

// SetTce Transmission Completed Enable
func (r *registerFdcan_ieType) SetTce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldTceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldTceMask)
	}
}

const (
	RegisterFdcan_ieFieldTcfeShift = 10
	RegisterFdcan_ieFieldTcfeMask  = 0x400
)

// GetTcfe Transmission Cancellation Finished Enable
func (r *registerFdcan_ieType) GetTcfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldTcfeMask) != 0
}

// SetTcfe Transmission Cancellation Finished Enable
func (r *registerFdcan_ieType) SetTcfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldTcfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldTcfeMask)
	}
}

const (
	RegisterFdcan_ieFieldTefeShift = 11
	RegisterFdcan_ieFieldTefeMask  = 0x800
)

// GetTefe Tx FIFO Empty Enable
func (r *registerFdcan_ieType) GetTefe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldTefeMask) != 0
}

// SetTefe Tx FIFO Empty Enable
func (r *registerFdcan_ieType) SetTefe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldTefeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldTefeMask)
	}
}

const (
	RegisterFdcan_ieFieldTefneShift = 12
	RegisterFdcan_ieFieldTefneMask  = 0x1000
)

// GetTefne Tx Event FIFO New Entry Enable
func (r *registerFdcan_ieType) GetTefne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldTefneMask) != 0
}

// SetTefne Tx Event FIFO New Entry Enable
func (r *registerFdcan_ieType) SetTefne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldTefneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldTefneMask)
	}
}

const (
	RegisterFdcan_ieFieldTefweShift = 13
	RegisterFdcan_ieFieldTefweMask  = 0x2000
)

// GetTefwe Tx Event FIFO Watermark Reached Enable
func (r *registerFdcan_ieType) GetTefwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldTefweMask) != 0
}

// SetTefwe Tx Event FIFO Watermark Reached Enable
func (r *registerFdcan_ieType) SetTefwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldTefweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldTefweMask)
	}
}

const (
	RegisterFdcan_ieFieldTeffeShift = 14
	RegisterFdcan_ieFieldTeffeMask  = 0x4000
)

// GetTeffe Tx Event FIFO Full Enable
func (r *registerFdcan_ieType) GetTeffe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldTeffeMask) != 0
}

// SetTeffe Tx Event FIFO Full Enable
func (r *registerFdcan_ieType) SetTeffe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldTeffeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldTeffeMask)
	}
}

const (
	RegisterFdcan_ieFieldTefleShift = 15
	RegisterFdcan_ieFieldTefleMask  = 0x8000
)

// GetTefle Tx Event FIFO Element Lost Enable
func (r *registerFdcan_ieType) GetTefle() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldTefleMask) != 0
}

// SetTefle Tx Event FIFO Element Lost Enable
func (r *registerFdcan_ieType) SetTefle(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldTefleMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldTefleMask)
	}
}

const (
	RegisterFdcan_ieFieldTsweShift = 16
	RegisterFdcan_ieFieldTsweMask  = 0x10000
)

// GetTswe Timestamp Wraparound Enable
func (r *registerFdcan_ieType) GetTswe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldTsweMask) != 0
}

// SetTswe Timestamp Wraparound Enable
func (r *registerFdcan_ieType) SetTswe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldTsweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldTsweMask)
	}
}

const (
	RegisterFdcan_ieFieldMrafeShift = 17
	RegisterFdcan_ieFieldMrafeMask  = 0x20000
)

// GetMrafe Message RAM Access Failure Enable
func (r *registerFdcan_ieType) GetMrafe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldMrafeMask) != 0
}

// SetMrafe Message RAM Access Failure Enable
func (r *registerFdcan_ieType) SetMrafe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldMrafeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldMrafeMask)
	}
}

const (
	RegisterFdcan_ieFieldTooeShift = 18
	RegisterFdcan_ieFieldTooeMask  = 0x40000
)

// GetTooe Timeout Occurred Enable
func (r *registerFdcan_ieType) GetTooe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldTooeMask) != 0
}

// SetTooe Timeout Occurred Enable
func (r *registerFdcan_ieType) SetTooe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldTooeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldTooeMask)
	}
}

const (
	RegisterFdcan_ieFieldDrxeShift = 19
	RegisterFdcan_ieFieldDrxeMask  = 0x80000
)

// GetDrxe Message stored to Dedicated Rx Buffer Enable
func (r *registerFdcan_ieType) GetDrxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldDrxeMask) != 0
}

// SetDrxe Message stored to Dedicated Rx Buffer Enable
func (r *registerFdcan_ieType) SetDrxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldDrxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldDrxeMask)
	}
}

const (
	RegisterFdcan_ieFieldBeceShift = 20
	RegisterFdcan_ieFieldBeceMask  = 0x100000
)

// GetBece Bit Error Corrected Interrupt Enable
func (r *registerFdcan_ieType) GetBece() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldBeceMask) != 0
}

// SetBece Bit Error Corrected Interrupt Enable
func (r *registerFdcan_ieType) SetBece(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldBeceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldBeceMask)
	}
}

const (
	RegisterFdcan_ieFieldBeueShift = 21
	RegisterFdcan_ieFieldBeueMask  = 0x200000
)

// GetBeue Bit Error Uncorrected Interrupt Enable
func (r *registerFdcan_ieType) GetBeue() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldBeueMask) != 0
}

// SetBeue Bit Error Uncorrected Interrupt Enable
func (r *registerFdcan_ieType) SetBeue(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldBeueMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldBeueMask)
	}
}

const (
	RegisterFdcan_ieFieldEloeShift = 22
	RegisterFdcan_ieFieldEloeMask  = 0x400000
)

// GetEloe Error Logging Overflow Enable
func (r *registerFdcan_ieType) GetEloe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldEloeMask) != 0
}

// SetEloe Error Logging Overflow Enable
func (r *registerFdcan_ieType) SetEloe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldEloeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldEloeMask)
	}
}

const (
	RegisterFdcan_ieFieldEpeShift = 23
	RegisterFdcan_ieFieldEpeMask  = 0x800000
)

// GetEpe Error Passive Enable
func (r *registerFdcan_ieType) GetEpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldEpeMask) != 0
}

// SetEpe Error Passive Enable
func (r *registerFdcan_ieType) SetEpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldEpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldEpeMask)
	}
}

const (
	RegisterFdcan_ieFieldEweShift = 24
	RegisterFdcan_ieFieldEweMask  = 0x1000000
)

// GetEwe Warning Status Enable
func (r *registerFdcan_ieType) GetEwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldEweMask) != 0
}

// SetEwe Warning Status Enable
func (r *registerFdcan_ieType) SetEwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldEweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldEweMask)
	}
}

const (
	RegisterFdcan_ieFieldBoeShift = 25
	RegisterFdcan_ieFieldBoeMask  = 0x2000000
)

// GetBoe Bus_Off Status Enable
func (r *registerFdcan_ieType) GetBoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldBoeMask) != 0
}

// SetBoe Bus_Off Status Enable
func (r *registerFdcan_ieType) SetBoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldBoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldBoeMask)
	}
}

const (
	RegisterFdcan_ieFieldWdieShift = 26
	RegisterFdcan_ieFieldWdieMask  = 0x4000000
)

// GetWdie Watchdog Interrupt Enable
func (r *registerFdcan_ieType) GetWdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldWdieMask) != 0
}

// SetWdie Watchdog Interrupt Enable
func (r *registerFdcan_ieType) SetWdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldWdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldWdieMask)
	}
}

const (
	RegisterFdcan_ieFieldPeaeShift = 27
	RegisterFdcan_ieFieldPeaeMask  = 0x8000000
)

// GetPeae Protocol Error in Arbitration Phase Enable
func (r *registerFdcan_ieType) GetPeae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldPeaeMask) != 0
}

// SetPeae Protocol Error in Arbitration Phase Enable
func (r *registerFdcan_ieType) SetPeae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldPeaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldPeaeMask)
	}
}

const (
	RegisterFdcan_ieFieldPedeShift = 28
	RegisterFdcan_ieFieldPedeMask  = 0x10000000
)

// GetPede Protocol Error in Data Phase Enable
func (r *registerFdcan_ieType) GetPede() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldPedeMask) != 0
}

// SetPede Protocol Error in Data Phase Enable
func (r *registerFdcan_ieType) SetPede(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldPedeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldPedeMask)
	}
}

const (
	RegisterFdcan_ieFieldAraeShift = 29
	RegisterFdcan_ieFieldAraeMask  = 0x20000000
)

// GetArae Access to Reserved Address Enable
func (r *registerFdcan_ieType) GetArae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ieFieldAraeMask) != 0
}

// SetArae Access to Reserved Address Enable
func (r *registerFdcan_ieType) SetArae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ieFieldAraeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ieFieldAraeMask)
	}
}

// registerFdcan_ilsType FDCAN Interrupt Line Select Register
type registerFdcan_ilsType uint32

const (
	RegisterFdcan_ilsFieldRf0nlShift = 0
	RegisterFdcan_ilsFieldRf0nlMask  = 0x1
)

// GetRf0nl Rx FIFO 0 New Message Interrupt Line
func (r *registerFdcan_ilsType) GetRf0nl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldRf0nlMask) != 0
}

// SetRf0nl Rx FIFO 0 New Message Interrupt Line
func (r *registerFdcan_ilsType) SetRf0nl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldRf0nlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldRf0nlMask)
	}
}

const (
	RegisterFdcan_ilsFieldRf0wlShift = 1
	RegisterFdcan_ilsFieldRf0wlMask  = 0x2
)

// GetRf0wl Rx FIFO 0 Watermark Reached Interrupt Line
func (r *registerFdcan_ilsType) GetRf0wl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldRf0wlMask) != 0
}

// SetRf0wl Rx FIFO 0 Watermark Reached Interrupt Line
func (r *registerFdcan_ilsType) SetRf0wl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldRf0wlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldRf0wlMask)
	}
}

const (
	RegisterFdcan_ilsFieldRf0flShift = 2
	RegisterFdcan_ilsFieldRf0flMask  = 0x4
)

// GetRf0fl Rx FIFO 0 Full Interrupt Line
func (r *registerFdcan_ilsType) GetRf0fl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldRf0flMask) != 0
}

// SetRf0fl Rx FIFO 0 Full Interrupt Line
func (r *registerFdcan_ilsType) SetRf0fl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldRf0flMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldRf0flMask)
	}
}

const (
	RegisterFdcan_ilsFieldRf0llShift = 3
	RegisterFdcan_ilsFieldRf0llMask  = 0x8
)

// GetRf0ll Rx FIFO 0 Message Lost Interrupt Line
func (r *registerFdcan_ilsType) GetRf0ll() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldRf0llMask) != 0
}

// SetRf0ll Rx FIFO 0 Message Lost Interrupt Line
func (r *registerFdcan_ilsType) SetRf0ll(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldRf0llMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldRf0llMask)
	}
}

const (
	RegisterFdcan_ilsFieldRf1nlShift = 4
	RegisterFdcan_ilsFieldRf1nlMask  = 0x10
)

// GetRf1nl Rx FIFO 1 New Message Interrupt Line
func (r *registerFdcan_ilsType) GetRf1nl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldRf1nlMask) != 0
}

// SetRf1nl Rx FIFO 1 New Message Interrupt Line
func (r *registerFdcan_ilsType) SetRf1nl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldRf1nlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldRf1nlMask)
	}
}

const (
	RegisterFdcan_ilsFieldRf1wlShift = 5
	RegisterFdcan_ilsFieldRf1wlMask  = 0x20
)

// GetRf1wl Rx FIFO 1 Watermark Reached Interrupt Line
func (r *registerFdcan_ilsType) GetRf1wl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldRf1wlMask) != 0
}

// SetRf1wl Rx FIFO 1 Watermark Reached Interrupt Line
func (r *registerFdcan_ilsType) SetRf1wl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldRf1wlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldRf1wlMask)
	}
}

const (
	RegisterFdcan_ilsFieldRf1flShift = 6
	RegisterFdcan_ilsFieldRf1flMask  = 0x40
)

// GetRf1fl Rx FIFO 1 Full Interrupt Line
func (r *registerFdcan_ilsType) GetRf1fl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldRf1flMask) != 0
}

// SetRf1fl Rx FIFO 1 Full Interrupt Line
func (r *registerFdcan_ilsType) SetRf1fl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldRf1flMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldRf1flMask)
	}
}

const (
	RegisterFdcan_ilsFieldRf1llShift = 7
	RegisterFdcan_ilsFieldRf1llMask  = 0x80
)

// GetRf1ll Rx FIFO 1 Message Lost Interrupt Line
func (r *registerFdcan_ilsType) GetRf1ll() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldRf1llMask) != 0
}

// SetRf1ll Rx FIFO 1 Message Lost Interrupt Line
func (r *registerFdcan_ilsType) SetRf1ll(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldRf1llMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldRf1llMask)
	}
}

const (
	RegisterFdcan_ilsFieldHpmlShift = 8
	RegisterFdcan_ilsFieldHpmlMask  = 0x100
)

// GetHpml High Priority Message Interrupt Line
func (r *registerFdcan_ilsType) GetHpml() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldHpmlMask) != 0
}

// SetHpml High Priority Message Interrupt Line
func (r *registerFdcan_ilsType) SetHpml(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldHpmlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldHpmlMask)
	}
}

const (
	RegisterFdcan_ilsFieldTclShift = 9
	RegisterFdcan_ilsFieldTclMask  = 0x200
)

// GetTcl Transmission Completed Interrupt Line
func (r *registerFdcan_ilsType) GetTcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldTclMask) != 0
}

// SetTcl Transmission Completed Interrupt Line
func (r *registerFdcan_ilsType) SetTcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldTclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldTclMask)
	}
}

const (
	RegisterFdcan_ilsFieldTcflShift = 10
	RegisterFdcan_ilsFieldTcflMask  = 0x400
)

// GetTcfl Transmission Cancellation Finished Interrupt Line
func (r *registerFdcan_ilsType) GetTcfl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldTcflMask) != 0
}

// SetTcfl Transmission Cancellation Finished Interrupt Line
func (r *registerFdcan_ilsType) SetTcfl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldTcflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldTcflMask)
	}
}

const (
	RegisterFdcan_ilsFieldTeflShift = 11
	RegisterFdcan_ilsFieldTeflMask  = 0x800
)

// GetTefl Tx FIFO Empty Interrupt Line
func (r *registerFdcan_ilsType) GetTefl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldTeflMask) != 0
}

// SetTefl Tx FIFO Empty Interrupt Line
func (r *registerFdcan_ilsType) SetTefl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldTeflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldTeflMask)
	}
}

const (
	RegisterFdcan_ilsFieldTefnlShift = 12
	RegisterFdcan_ilsFieldTefnlMask  = 0x1000
)

// GetTefnl Tx Event FIFO New Entry Interrupt Line
func (r *registerFdcan_ilsType) GetTefnl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldTefnlMask) != 0
}

// SetTefnl Tx Event FIFO New Entry Interrupt Line
func (r *registerFdcan_ilsType) SetTefnl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldTefnlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldTefnlMask)
	}
}

const (
	RegisterFdcan_ilsFieldTefwlShift = 13
	RegisterFdcan_ilsFieldTefwlMask  = 0x2000
)

// GetTefwl Tx Event FIFO Watermark Reached Interrupt Line
func (r *registerFdcan_ilsType) GetTefwl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldTefwlMask) != 0
}

// SetTefwl Tx Event FIFO Watermark Reached Interrupt Line
func (r *registerFdcan_ilsType) SetTefwl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldTefwlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldTefwlMask)
	}
}

const (
	RegisterFdcan_ilsFieldTefflShift = 14
	RegisterFdcan_ilsFieldTefflMask  = 0x4000
)

// GetTeffl Tx Event FIFO Full Interrupt Line
func (r *registerFdcan_ilsType) GetTeffl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldTefflMask) != 0
}

// SetTeffl Tx Event FIFO Full Interrupt Line
func (r *registerFdcan_ilsType) SetTeffl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldTefflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldTefflMask)
	}
}

const (
	RegisterFdcan_ilsFieldTefllShift = 15
	RegisterFdcan_ilsFieldTefllMask  = 0x8000
)

// GetTefll Tx Event FIFO Element Lost Interrupt Line
func (r *registerFdcan_ilsType) GetTefll() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldTefllMask) != 0
}

// SetTefll Tx Event FIFO Element Lost Interrupt Line
func (r *registerFdcan_ilsType) SetTefll(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldTefllMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldTefllMask)
	}
}

const (
	RegisterFdcan_ilsFieldTswlShift = 16
	RegisterFdcan_ilsFieldTswlMask  = 0x10000
)

// GetTswl Timestamp Wraparound Interrupt Line
func (r *registerFdcan_ilsType) GetTswl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldTswlMask) != 0
}

// SetTswl Timestamp Wraparound Interrupt Line
func (r *registerFdcan_ilsType) SetTswl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldTswlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldTswlMask)
	}
}

const (
	RegisterFdcan_ilsFieldMraflShift = 17
	RegisterFdcan_ilsFieldMraflMask  = 0x20000
)

// GetMrafl Message RAM Access Failure Interrupt Line
func (r *registerFdcan_ilsType) GetMrafl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldMraflMask) != 0
}

// SetMrafl Message RAM Access Failure Interrupt Line
func (r *registerFdcan_ilsType) SetMrafl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldMraflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldMraflMask)
	}
}

const (
	RegisterFdcan_ilsFieldToolShift = 18
	RegisterFdcan_ilsFieldToolMask  = 0x40000
)

// GetTool Timeout Occurred Interrupt Line
func (r *registerFdcan_ilsType) GetTool() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldToolMask) != 0
}

// SetTool Timeout Occurred Interrupt Line
func (r *registerFdcan_ilsType) SetTool(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldToolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldToolMask)
	}
}

const (
	RegisterFdcan_ilsFieldDrxlShift = 19
	RegisterFdcan_ilsFieldDrxlMask  = 0x80000
)

// GetDrxl Message stored to Dedicated Rx Buffer Interrupt Line
func (r *registerFdcan_ilsType) GetDrxl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldDrxlMask) != 0
}

// SetDrxl Message stored to Dedicated Rx Buffer Interrupt Line
func (r *registerFdcan_ilsType) SetDrxl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldDrxlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldDrxlMask)
	}
}

const (
	RegisterFdcan_ilsFieldBeclShift = 20
	RegisterFdcan_ilsFieldBeclMask  = 0x100000
)

// GetBecl Bit Error Corrected Interrupt Line
func (r *registerFdcan_ilsType) GetBecl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldBeclMask) != 0
}

// SetBecl Bit Error Corrected Interrupt Line
func (r *registerFdcan_ilsType) SetBecl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldBeclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldBeclMask)
	}
}

const (
	RegisterFdcan_ilsFieldBeulShift = 21
	RegisterFdcan_ilsFieldBeulMask  = 0x200000
)

// GetBeul Bit Error Uncorrected Interrupt Line
func (r *registerFdcan_ilsType) GetBeul() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldBeulMask) != 0
}

// SetBeul Bit Error Uncorrected Interrupt Line
func (r *registerFdcan_ilsType) SetBeul(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldBeulMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldBeulMask)
	}
}

const (
	RegisterFdcan_ilsFieldElolShift = 22
	RegisterFdcan_ilsFieldElolMask  = 0x400000
)

// GetElol Error Logging Overflow Interrupt Line
func (r *registerFdcan_ilsType) GetElol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldElolMask) != 0
}

// SetElol Error Logging Overflow Interrupt Line
func (r *registerFdcan_ilsType) SetElol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldElolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldElolMask)
	}
}

const (
	RegisterFdcan_ilsFieldEplShift = 23
	RegisterFdcan_ilsFieldEplMask  = 0x800000
)

// GetEpl Error Passive Interrupt Line
func (r *registerFdcan_ilsType) GetEpl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldEplMask) != 0
}

// SetEpl Error Passive Interrupt Line
func (r *registerFdcan_ilsType) SetEpl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldEplMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldEplMask)
	}
}

const (
	RegisterFdcan_ilsFieldEwlShift = 24
	RegisterFdcan_ilsFieldEwlMask  = 0x1000000
)

// GetEwl Warning Status Interrupt Line
func (r *registerFdcan_ilsType) GetEwl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldEwlMask) != 0
}

// SetEwl Warning Status Interrupt Line
func (r *registerFdcan_ilsType) SetEwl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldEwlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldEwlMask)
	}
}

const (
	RegisterFdcan_ilsFieldBolShift = 25
	RegisterFdcan_ilsFieldBolMask  = 0x2000000
)

// GetBol Bus_Off Status
func (r *registerFdcan_ilsType) GetBol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldBolMask) != 0
}

// SetBol Bus_Off Status
func (r *registerFdcan_ilsType) SetBol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldBolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldBolMask)
	}
}

const (
	RegisterFdcan_ilsFieldWdilShift = 26
	RegisterFdcan_ilsFieldWdilMask  = 0x4000000
)

// GetWdil Watchdog Interrupt Line
func (r *registerFdcan_ilsType) GetWdil() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldWdilMask) != 0
}

// SetWdil Watchdog Interrupt Line
func (r *registerFdcan_ilsType) SetWdil(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldWdilMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldWdilMask)
	}
}

const (
	RegisterFdcan_ilsFieldPealShift = 27
	RegisterFdcan_ilsFieldPealMask  = 0x8000000
)

// GetPeal Protocol Error in Arbitration Phase Line
func (r *registerFdcan_ilsType) GetPeal() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldPealMask) != 0
}

// SetPeal Protocol Error in Arbitration Phase Line
func (r *registerFdcan_ilsType) SetPeal(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldPealMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldPealMask)
	}
}

const (
	RegisterFdcan_ilsFieldPedlShift = 28
	RegisterFdcan_ilsFieldPedlMask  = 0x10000000
)

// GetPedl Protocol Error in Data Phase Line
func (r *registerFdcan_ilsType) GetPedl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldPedlMask) != 0
}

// SetPedl Protocol Error in Data Phase Line
func (r *registerFdcan_ilsType) SetPedl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldPedlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldPedlMask)
	}
}

const (
	RegisterFdcan_ilsFieldAralShift = 29
	RegisterFdcan_ilsFieldAralMask  = 0x20000000
)

// GetAral Access to Reserved Address Line
func (r *registerFdcan_ilsType) GetAral() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ilsFieldAralMask) != 0
}

// SetAral Access to Reserved Address Line
func (r *registerFdcan_ilsType) SetAral(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ilsFieldAralMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ilsFieldAralMask)
	}
}

// registerFdcan_ileType FDCAN Interrupt Line Enable Register
type registerFdcan_ileType uint32

const (
	RegisterFdcan_ileFieldEint0Shift = 0
	RegisterFdcan_ileFieldEint0Mask  = 0x1
)

// GetEint0 Enable Interrupt Line 0
func (r *registerFdcan_ileType) GetEint0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ileFieldEint0Mask) != 0
}

// SetEint0 Enable Interrupt Line 0
func (r *registerFdcan_ileType) SetEint0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ileFieldEint0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ileFieldEint0Mask)
	}
}

const (
	RegisterFdcan_ileFieldEint1Shift = 1
	RegisterFdcan_ileFieldEint1Mask  = 0x2
)

// GetEint1 Enable Interrupt Line 1
func (r *registerFdcan_ileType) GetEint1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ileFieldEint1Mask) != 0
}

// SetEint1 Enable Interrupt Line 1
func (r *registerFdcan_ileType) SetEint1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ileFieldEint1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ileFieldEint1Mask)
	}
}

// registerFdcan_gfcType FDCAN Global Filter Configuration Register
type registerFdcan_gfcType uint32

const (
	RegisterFdcan_gfcFieldRrfeShift = 0
	RegisterFdcan_gfcFieldRrfeMask  = 0x1
)

// GetRrfe Reject Remote Frames Extended
func (r *registerFdcan_gfcType) GetRrfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_gfcFieldRrfeMask) != 0
}

// SetRrfe Reject Remote Frames Extended
func (r *registerFdcan_gfcType) SetRrfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_gfcFieldRrfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_gfcFieldRrfeMask)
	}
}

const (
	RegisterFdcan_gfcFieldRrfsShift = 1
	RegisterFdcan_gfcFieldRrfsMask  = 0x2
)

// GetRrfs Reject Remote Frames Standard
func (r *registerFdcan_gfcType) GetRrfs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_gfcFieldRrfsMask) != 0
}

// SetRrfs Reject Remote Frames Standard
func (r *registerFdcan_gfcType) SetRrfs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_gfcFieldRrfsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_gfcFieldRrfsMask)
	}
}

const (
	RegisterFdcan_gfcFieldAnfeShift = 2
	RegisterFdcan_gfcFieldAnfeMask  = 0xc
)

// GetAnfe Accept Non-matching Frames Extended
func (r *registerFdcan_gfcType) GetAnfe() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_gfcFieldAnfeMask) >> RegisterFdcan_gfcFieldAnfeShift)
}

// SetAnfe Accept Non-matching Frames Extended
func (r *registerFdcan_gfcType) SetAnfe(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_gfcFieldAnfeMask)|(uint32(value)<<RegisterFdcan_gfcFieldAnfeShift))
}

const (
	RegisterFdcan_gfcFieldAnfsShift = 4
	RegisterFdcan_gfcFieldAnfsMask  = 0x30
)

// GetAnfs Accept Non-matching Frames Standard
func (r *registerFdcan_gfcType) GetAnfs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_gfcFieldAnfsMask) >> RegisterFdcan_gfcFieldAnfsShift)
}

// SetAnfs Accept Non-matching Frames Standard
func (r *registerFdcan_gfcType) SetAnfs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_gfcFieldAnfsMask)|(uint32(value)<<RegisterFdcan_gfcFieldAnfsShift))
}

// registerFdcan_sidfcType FDCAN Standard ID Filter Configuration Register
type registerFdcan_sidfcType uint32

const (
	RegisterFdcan_sidfcFieldFlssaShift = 2
	RegisterFdcan_sidfcFieldFlssaMask  = 0xfffc
)

// GetFlssa Filter List Standard Start Address
func (r *registerFdcan_sidfcType) GetFlssa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_sidfcFieldFlssaMask) >> RegisterFdcan_sidfcFieldFlssaShift)
}

// SetFlssa Filter List Standard Start Address
func (r *registerFdcan_sidfcType) SetFlssa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_sidfcFieldFlssaMask)|(uint32(value)<<RegisterFdcan_sidfcFieldFlssaShift))
}

const (
	RegisterFdcan_sidfcFieldLssShift = 16
	RegisterFdcan_sidfcFieldLssMask  = 0xff0000
)

// GetLss List Size Standard
func (r *registerFdcan_sidfcType) GetLss() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_sidfcFieldLssMask) >> RegisterFdcan_sidfcFieldLssShift)
}

// SetLss List Size Standard
func (r *registerFdcan_sidfcType) SetLss(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_sidfcFieldLssMask)|(uint32(value)<<RegisterFdcan_sidfcFieldLssShift))
}

// registerFdcan_xidfcType FDCAN Extended ID Filter Configuration Register
type registerFdcan_xidfcType uint32

const (
	RegisterFdcan_xidfcFieldFlesaShift = 2
	RegisterFdcan_xidfcFieldFlesaMask  = 0xfffc
)

// GetFlesa Filter List Standard Start Address
func (r *registerFdcan_xidfcType) GetFlesa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_xidfcFieldFlesaMask) >> RegisterFdcan_xidfcFieldFlesaShift)
}

// SetFlesa Filter List Standard Start Address
func (r *registerFdcan_xidfcType) SetFlesa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_xidfcFieldFlesaMask)|(uint32(value)<<RegisterFdcan_xidfcFieldFlesaShift))
}

const (
	RegisterFdcan_xidfcFieldLseShift = 16
	RegisterFdcan_xidfcFieldLseMask  = 0xff0000
)

// GetLse List Size Extended
func (r *registerFdcan_xidfcType) GetLse() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_xidfcFieldLseMask) >> RegisterFdcan_xidfcFieldLseShift)
}

// SetLse List Size Extended
func (r *registerFdcan_xidfcType) SetLse(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_xidfcFieldLseMask)|(uint32(value)<<RegisterFdcan_xidfcFieldLseShift))
}

// registerFdcan_xidamType FDCAN Extended ID and Mask Register
type registerFdcan_xidamType uint32

const (
	RegisterFdcan_xidamFieldEidmShift = 0
	RegisterFdcan_xidamFieldEidmMask  = 0x1fffffff
)

// GetEidm Extended ID Mask
func (r *registerFdcan_xidamType) GetEidm() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_xidamFieldEidmMask) >> RegisterFdcan_xidamFieldEidmShift)
}

// SetEidm Extended ID Mask
func (r *registerFdcan_xidamType) SetEidm(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_xidamFieldEidmMask)|(uint32(value)<<RegisterFdcan_xidamFieldEidmShift))
}

// registerFdcan_hpmsType FDCAN High Priority Message Status Register
type registerFdcan_hpmsType uint32

const (
	RegisterFdcan_hpmsFieldBidxShift = 0
	RegisterFdcan_hpmsFieldBidxMask  = 0x3f
)

// GetBidx Buffer Index
func (r *registerFdcan_hpmsType) GetBidx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_hpmsFieldBidxMask) >> RegisterFdcan_hpmsFieldBidxShift)
}

// SetBidx Buffer Index
func (r *registerFdcan_hpmsType) SetBidx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_hpmsFieldBidxMask)|(uint32(value)<<RegisterFdcan_hpmsFieldBidxShift))
}

const (
	RegisterFdcan_hpmsFieldMsiShift = 6
	RegisterFdcan_hpmsFieldMsiMask  = 0xc0
)

// GetMsi Message Storage Indicator
func (r *registerFdcan_hpmsType) GetMsi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_hpmsFieldMsiMask) >> RegisterFdcan_hpmsFieldMsiShift)
}

// SetMsi Message Storage Indicator
func (r *registerFdcan_hpmsType) SetMsi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_hpmsFieldMsiMask)|(uint32(value)<<RegisterFdcan_hpmsFieldMsiShift))
}

const (
	RegisterFdcan_hpmsFieldFidxShift = 8
	RegisterFdcan_hpmsFieldFidxMask  = 0x7f00
)

// GetFidx Filter Index
func (r *registerFdcan_hpmsType) GetFidx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_hpmsFieldFidxMask) >> RegisterFdcan_hpmsFieldFidxShift)
}

// SetFidx Filter Index
func (r *registerFdcan_hpmsType) SetFidx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_hpmsFieldFidxMask)|(uint32(value)<<RegisterFdcan_hpmsFieldFidxShift))
}

const (
	RegisterFdcan_hpmsFieldFlstShift = 15
	RegisterFdcan_hpmsFieldFlstMask  = 0x8000
)

// GetFlst Filter List
func (r *registerFdcan_hpmsType) GetFlst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_hpmsFieldFlstMask) != 0
}

// SetFlst Filter List
func (r *registerFdcan_hpmsType) SetFlst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_hpmsFieldFlstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_hpmsFieldFlstMask)
	}
}

// registerFdcan_ndat1Type FDCAN New Data 1 Register
type registerFdcan_ndat1Type uint32

const (
	RegisterFdcan_ndat1FieldNd0Shift = 0
	RegisterFdcan_ndat1FieldNd0Mask  = 0x1
)

// GetNd0 New data
func (r *registerFdcan_ndat1Type) GetNd0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd0Mask) != 0
}

// SetNd0 New data
func (r *registerFdcan_ndat1Type) SetNd0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd0Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd1Shift = 1
	RegisterFdcan_ndat1FieldNd1Mask  = 0x2
)

// GetNd1 New data
func (r *registerFdcan_ndat1Type) GetNd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd1Mask) != 0
}

// SetNd1 New data
func (r *registerFdcan_ndat1Type) SetNd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd1Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd2Shift = 2
	RegisterFdcan_ndat1FieldNd2Mask  = 0x4
)

// GetNd2 New data
func (r *registerFdcan_ndat1Type) GetNd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd2Mask) != 0
}

// SetNd2 New data
func (r *registerFdcan_ndat1Type) SetNd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd2Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd3Shift = 3
	RegisterFdcan_ndat1FieldNd3Mask  = 0x8
)

// GetNd3 New data
func (r *registerFdcan_ndat1Type) GetNd3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd3Mask) != 0
}

// SetNd3 New data
func (r *registerFdcan_ndat1Type) SetNd3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd3Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd4Shift = 4
	RegisterFdcan_ndat1FieldNd4Mask  = 0x10
)

// GetNd4 New data
func (r *registerFdcan_ndat1Type) GetNd4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd4Mask) != 0
}

// SetNd4 New data
func (r *registerFdcan_ndat1Type) SetNd4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd4Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd5Shift = 5
	RegisterFdcan_ndat1FieldNd5Mask  = 0x20
)

// GetNd5 New data
func (r *registerFdcan_ndat1Type) GetNd5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd5Mask) != 0
}

// SetNd5 New data
func (r *registerFdcan_ndat1Type) SetNd5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd5Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd6Shift = 6
	RegisterFdcan_ndat1FieldNd6Mask  = 0x40
)

// GetNd6 New data
func (r *registerFdcan_ndat1Type) GetNd6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd6Mask) != 0
}

// SetNd6 New data
func (r *registerFdcan_ndat1Type) SetNd6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd6Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd7Shift = 7
	RegisterFdcan_ndat1FieldNd7Mask  = 0x80
)

// GetNd7 New data
func (r *registerFdcan_ndat1Type) GetNd7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd7Mask) != 0
}

// SetNd7 New data
func (r *registerFdcan_ndat1Type) SetNd7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd7Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd8Shift = 8
	RegisterFdcan_ndat1FieldNd8Mask  = 0x100
)

// GetNd8 New data
func (r *registerFdcan_ndat1Type) GetNd8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd8Mask) != 0
}

// SetNd8 New data
func (r *registerFdcan_ndat1Type) SetNd8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd8Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd9Shift = 9
	RegisterFdcan_ndat1FieldNd9Mask  = 0x200
)

// GetNd9 New data
func (r *registerFdcan_ndat1Type) GetNd9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd9Mask) != 0
}

// SetNd9 New data
func (r *registerFdcan_ndat1Type) SetNd9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd9Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd10Shift = 10
	RegisterFdcan_ndat1FieldNd10Mask  = 0x400
)

// GetNd10 New data
func (r *registerFdcan_ndat1Type) GetNd10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd10Mask) != 0
}

// SetNd10 New data
func (r *registerFdcan_ndat1Type) SetNd10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd10Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd11Shift = 11
	RegisterFdcan_ndat1FieldNd11Mask  = 0x800
)

// GetNd11 New data
func (r *registerFdcan_ndat1Type) GetNd11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd11Mask) != 0
}

// SetNd11 New data
func (r *registerFdcan_ndat1Type) SetNd11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd11Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd12Shift = 12
	RegisterFdcan_ndat1FieldNd12Mask  = 0x1000
)

// GetNd12 New data
func (r *registerFdcan_ndat1Type) GetNd12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd12Mask) != 0
}

// SetNd12 New data
func (r *registerFdcan_ndat1Type) SetNd12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd12Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd13Shift = 13
	RegisterFdcan_ndat1FieldNd13Mask  = 0x2000
)

// GetNd13 New data
func (r *registerFdcan_ndat1Type) GetNd13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd13Mask) != 0
}

// SetNd13 New data
func (r *registerFdcan_ndat1Type) SetNd13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd13Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd14Shift = 14
	RegisterFdcan_ndat1FieldNd14Mask  = 0x4000
)

// GetNd14 New data
func (r *registerFdcan_ndat1Type) GetNd14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd14Mask) != 0
}

// SetNd14 New data
func (r *registerFdcan_ndat1Type) SetNd14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd14Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd15Shift = 15
	RegisterFdcan_ndat1FieldNd15Mask  = 0x8000
)

// GetNd15 New data
func (r *registerFdcan_ndat1Type) GetNd15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd15Mask) != 0
}

// SetNd15 New data
func (r *registerFdcan_ndat1Type) SetNd15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd15Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd16Shift = 16
	RegisterFdcan_ndat1FieldNd16Mask  = 0x10000
)

// GetNd16 New data
func (r *registerFdcan_ndat1Type) GetNd16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd16Mask) != 0
}

// SetNd16 New data
func (r *registerFdcan_ndat1Type) SetNd16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd16Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd17Shift = 17
	RegisterFdcan_ndat1FieldNd17Mask  = 0x20000
)

// GetNd17 New data
func (r *registerFdcan_ndat1Type) GetNd17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd17Mask) != 0
}

// SetNd17 New data
func (r *registerFdcan_ndat1Type) SetNd17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd17Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd18Shift = 18
	RegisterFdcan_ndat1FieldNd18Mask  = 0x40000
)

// GetNd18 New data
func (r *registerFdcan_ndat1Type) GetNd18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd18Mask) != 0
}

// SetNd18 New data
func (r *registerFdcan_ndat1Type) SetNd18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd18Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd19Shift = 19
	RegisterFdcan_ndat1FieldNd19Mask  = 0x80000
)

// GetNd19 New data
func (r *registerFdcan_ndat1Type) GetNd19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd19Mask) != 0
}

// SetNd19 New data
func (r *registerFdcan_ndat1Type) SetNd19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd19Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd20Shift = 20
	RegisterFdcan_ndat1FieldNd20Mask  = 0x100000
)

// GetNd20 New data
func (r *registerFdcan_ndat1Type) GetNd20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd20Mask) != 0
}

// SetNd20 New data
func (r *registerFdcan_ndat1Type) SetNd20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd20Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd21Shift = 21
	RegisterFdcan_ndat1FieldNd21Mask  = 0x200000
)

// GetNd21 New data
func (r *registerFdcan_ndat1Type) GetNd21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd21Mask) != 0
}

// SetNd21 New data
func (r *registerFdcan_ndat1Type) SetNd21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd21Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd22Shift = 22
	RegisterFdcan_ndat1FieldNd22Mask  = 0x400000
)

// GetNd22 New data
func (r *registerFdcan_ndat1Type) GetNd22() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd22Mask) != 0
}

// SetNd22 New data
func (r *registerFdcan_ndat1Type) SetNd22(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd22Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd22Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd23Shift = 23
	RegisterFdcan_ndat1FieldNd23Mask  = 0x800000
)

// GetNd23 New data
func (r *registerFdcan_ndat1Type) GetNd23() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd23Mask) != 0
}

// SetNd23 New data
func (r *registerFdcan_ndat1Type) SetNd23(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd23Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd23Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd24Shift = 24
	RegisterFdcan_ndat1FieldNd24Mask  = 0x1000000
)

// GetNd24 New data
func (r *registerFdcan_ndat1Type) GetNd24() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd24Mask) != 0
}

// SetNd24 New data
func (r *registerFdcan_ndat1Type) SetNd24(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd24Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd24Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd25Shift = 25
	RegisterFdcan_ndat1FieldNd25Mask  = 0x2000000
)

// GetNd25 New data
func (r *registerFdcan_ndat1Type) GetNd25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd25Mask) != 0
}

// SetNd25 New data
func (r *registerFdcan_ndat1Type) SetNd25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd25Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd26Shift = 26
	RegisterFdcan_ndat1FieldNd26Mask  = 0x4000000
)

// GetNd26 New data
func (r *registerFdcan_ndat1Type) GetNd26() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd26Mask) != 0
}

// SetNd26 New data
func (r *registerFdcan_ndat1Type) SetNd26(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd26Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd26Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd27Shift = 27
	RegisterFdcan_ndat1FieldNd27Mask  = 0x8000000
)

// GetNd27 New data
func (r *registerFdcan_ndat1Type) GetNd27() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd27Mask) != 0
}

// SetNd27 New data
func (r *registerFdcan_ndat1Type) SetNd27(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd27Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd27Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd28Shift = 28
	RegisterFdcan_ndat1FieldNd28Mask  = 0x10000000
)

// GetNd28 New data
func (r *registerFdcan_ndat1Type) GetNd28() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd28Mask) != 0
}

// SetNd28 New data
func (r *registerFdcan_ndat1Type) SetNd28(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd28Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd28Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd29Shift = 29
	RegisterFdcan_ndat1FieldNd29Mask  = 0x20000000
)

// GetNd29 New data
func (r *registerFdcan_ndat1Type) GetNd29() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd29Mask) != 0
}

// SetNd29 New data
func (r *registerFdcan_ndat1Type) SetNd29(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd29Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd29Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd30Shift = 30
	RegisterFdcan_ndat1FieldNd30Mask  = 0x40000000
)

// GetNd30 New data
func (r *registerFdcan_ndat1Type) GetNd30() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd30Mask) != 0
}

// SetNd30 New data
func (r *registerFdcan_ndat1Type) SetNd30(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd30Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd30Mask)
	}
}

const (
	RegisterFdcan_ndat1FieldNd31Shift = 31
	RegisterFdcan_ndat1FieldNd31Mask  = 0x80000000
)

// GetNd31 New data
func (r *registerFdcan_ndat1Type) GetNd31() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat1FieldNd31Mask) != 0
}

// SetNd31 New data
func (r *registerFdcan_ndat1Type) SetNd31(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat1FieldNd31Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat1FieldNd31Mask)
	}
}

// registerFdcan_ndat2Type FDCAN New Data 2 Register
type registerFdcan_ndat2Type uint32

const (
	RegisterFdcan_ndat2FieldNd32Shift = 0
	RegisterFdcan_ndat2FieldNd32Mask  = 0x1
)

// GetNd32 New data
func (r *registerFdcan_ndat2Type) GetNd32() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd32Mask) != 0
}

// SetNd32 New data
func (r *registerFdcan_ndat2Type) SetNd32(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd32Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd32Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd33Shift = 1
	RegisterFdcan_ndat2FieldNd33Mask  = 0x2
)

// GetNd33 New data
func (r *registerFdcan_ndat2Type) GetNd33() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd33Mask) != 0
}

// SetNd33 New data
func (r *registerFdcan_ndat2Type) SetNd33(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd33Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd33Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd34Shift = 2
	RegisterFdcan_ndat2FieldNd34Mask  = 0x4
)

// GetNd34 New data
func (r *registerFdcan_ndat2Type) GetNd34() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd34Mask) != 0
}

// SetNd34 New data
func (r *registerFdcan_ndat2Type) SetNd34(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd34Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd34Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd35Shift = 3
	RegisterFdcan_ndat2FieldNd35Mask  = 0x8
)

// GetNd35 New data
func (r *registerFdcan_ndat2Type) GetNd35() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd35Mask) != 0
}

// SetNd35 New data
func (r *registerFdcan_ndat2Type) SetNd35(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd35Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd35Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd36Shift = 4
	RegisterFdcan_ndat2FieldNd36Mask  = 0x10
)

// GetNd36 New data
func (r *registerFdcan_ndat2Type) GetNd36() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd36Mask) != 0
}

// SetNd36 New data
func (r *registerFdcan_ndat2Type) SetNd36(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd36Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd36Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd37Shift = 5
	RegisterFdcan_ndat2FieldNd37Mask  = 0x20
)

// GetNd37 New data
func (r *registerFdcan_ndat2Type) GetNd37() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd37Mask) != 0
}

// SetNd37 New data
func (r *registerFdcan_ndat2Type) SetNd37(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd37Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd37Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd38Shift = 6
	RegisterFdcan_ndat2FieldNd38Mask  = 0x40
)

// GetNd38 New data
func (r *registerFdcan_ndat2Type) GetNd38() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd38Mask) != 0
}

// SetNd38 New data
func (r *registerFdcan_ndat2Type) SetNd38(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd38Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd38Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd39Shift = 7
	RegisterFdcan_ndat2FieldNd39Mask  = 0x80
)

// GetNd39 New data
func (r *registerFdcan_ndat2Type) GetNd39() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd39Mask) != 0
}

// SetNd39 New data
func (r *registerFdcan_ndat2Type) SetNd39(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd39Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd39Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd40Shift = 8
	RegisterFdcan_ndat2FieldNd40Mask  = 0x100
)

// GetNd40 New data
func (r *registerFdcan_ndat2Type) GetNd40() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd40Mask) != 0
}

// SetNd40 New data
func (r *registerFdcan_ndat2Type) SetNd40(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd40Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd40Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd41Shift = 9
	RegisterFdcan_ndat2FieldNd41Mask  = 0x200
)

// GetNd41 New data
func (r *registerFdcan_ndat2Type) GetNd41() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd41Mask) != 0
}

// SetNd41 New data
func (r *registerFdcan_ndat2Type) SetNd41(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd41Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd41Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd42Shift = 10
	RegisterFdcan_ndat2FieldNd42Mask  = 0x400
)

// GetNd42 New data
func (r *registerFdcan_ndat2Type) GetNd42() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd42Mask) != 0
}

// SetNd42 New data
func (r *registerFdcan_ndat2Type) SetNd42(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd42Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd42Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd43Shift = 11
	RegisterFdcan_ndat2FieldNd43Mask  = 0x800
)

// GetNd43 New data
func (r *registerFdcan_ndat2Type) GetNd43() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd43Mask) != 0
}

// SetNd43 New data
func (r *registerFdcan_ndat2Type) SetNd43(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd43Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd43Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd44Shift = 12
	RegisterFdcan_ndat2FieldNd44Mask  = 0x1000
)

// GetNd44 New data
func (r *registerFdcan_ndat2Type) GetNd44() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd44Mask) != 0
}

// SetNd44 New data
func (r *registerFdcan_ndat2Type) SetNd44(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd44Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd44Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd45Shift = 13
	RegisterFdcan_ndat2FieldNd45Mask  = 0x2000
)

// GetNd45 New data
func (r *registerFdcan_ndat2Type) GetNd45() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd45Mask) != 0
}

// SetNd45 New data
func (r *registerFdcan_ndat2Type) SetNd45(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd45Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd45Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd46Shift = 14
	RegisterFdcan_ndat2FieldNd46Mask  = 0x4000
)

// GetNd46 New data
func (r *registerFdcan_ndat2Type) GetNd46() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd46Mask) != 0
}

// SetNd46 New data
func (r *registerFdcan_ndat2Type) SetNd46(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd46Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd46Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd47Shift = 15
	RegisterFdcan_ndat2FieldNd47Mask  = 0x8000
)

// GetNd47 New data
func (r *registerFdcan_ndat2Type) GetNd47() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd47Mask) != 0
}

// SetNd47 New data
func (r *registerFdcan_ndat2Type) SetNd47(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd47Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd47Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd48Shift = 16
	RegisterFdcan_ndat2FieldNd48Mask  = 0x10000
)

// GetNd48 New data
func (r *registerFdcan_ndat2Type) GetNd48() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd48Mask) != 0
}

// SetNd48 New data
func (r *registerFdcan_ndat2Type) SetNd48(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd48Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd48Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd49Shift = 17
	RegisterFdcan_ndat2FieldNd49Mask  = 0x20000
)

// GetNd49 New data
func (r *registerFdcan_ndat2Type) GetNd49() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd49Mask) != 0
}

// SetNd49 New data
func (r *registerFdcan_ndat2Type) SetNd49(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd49Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd49Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd50Shift = 18
	RegisterFdcan_ndat2FieldNd50Mask  = 0x40000
)

// GetNd50 New data
func (r *registerFdcan_ndat2Type) GetNd50() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd50Mask) != 0
}

// SetNd50 New data
func (r *registerFdcan_ndat2Type) SetNd50(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd50Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd50Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd51Shift = 19
	RegisterFdcan_ndat2FieldNd51Mask  = 0x80000
)

// GetNd51 New data
func (r *registerFdcan_ndat2Type) GetNd51() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd51Mask) != 0
}

// SetNd51 New data
func (r *registerFdcan_ndat2Type) SetNd51(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd51Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd51Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd52Shift = 20
	RegisterFdcan_ndat2FieldNd52Mask  = 0x100000
)

// GetNd52 New data
func (r *registerFdcan_ndat2Type) GetNd52() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd52Mask) != 0
}

// SetNd52 New data
func (r *registerFdcan_ndat2Type) SetNd52(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd52Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd52Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd53Shift = 21
	RegisterFdcan_ndat2FieldNd53Mask  = 0x200000
)

// GetNd53 New data
func (r *registerFdcan_ndat2Type) GetNd53() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd53Mask) != 0
}

// SetNd53 New data
func (r *registerFdcan_ndat2Type) SetNd53(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd53Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd53Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd54Shift = 22
	RegisterFdcan_ndat2FieldNd54Mask  = 0x400000
)

// GetNd54 New data
func (r *registerFdcan_ndat2Type) GetNd54() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd54Mask) != 0
}

// SetNd54 New data
func (r *registerFdcan_ndat2Type) SetNd54(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd54Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd54Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd55Shift = 23
	RegisterFdcan_ndat2FieldNd55Mask  = 0x800000
)

// GetNd55 New data
func (r *registerFdcan_ndat2Type) GetNd55() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd55Mask) != 0
}

// SetNd55 New data
func (r *registerFdcan_ndat2Type) SetNd55(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd55Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd55Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd56Shift = 24
	RegisterFdcan_ndat2FieldNd56Mask  = 0x1000000
)

// GetNd56 New data
func (r *registerFdcan_ndat2Type) GetNd56() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd56Mask) != 0
}

// SetNd56 New data
func (r *registerFdcan_ndat2Type) SetNd56(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd56Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd56Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd57Shift = 25
	RegisterFdcan_ndat2FieldNd57Mask  = 0x2000000
)

// GetNd57 New data
func (r *registerFdcan_ndat2Type) GetNd57() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd57Mask) != 0
}

// SetNd57 New data
func (r *registerFdcan_ndat2Type) SetNd57(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd57Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd57Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd58Shift = 26
	RegisterFdcan_ndat2FieldNd58Mask  = 0x4000000
)

// GetNd58 New data
func (r *registerFdcan_ndat2Type) GetNd58() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd58Mask) != 0
}

// SetNd58 New data
func (r *registerFdcan_ndat2Type) SetNd58(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd58Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd58Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd59Shift = 27
	RegisterFdcan_ndat2FieldNd59Mask  = 0x8000000
)

// GetNd59 New data
func (r *registerFdcan_ndat2Type) GetNd59() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd59Mask) != 0
}

// SetNd59 New data
func (r *registerFdcan_ndat2Type) SetNd59(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd59Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd59Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd60Shift = 28
	RegisterFdcan_ndat2FieldNd60Mask  = 0x10000000
)

// GetNd60 New data
func (r *registerFdcan_ndat2Type) GetNd60() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd60Mask) != 0
}

// SetNd60 New data
func (r *registerFdcan_ndat2Type) SetNd60(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd60Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd60Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd61Shift = 29
	RegisterFdcan_ndat2FieldNd61Mask  = 0x20000000
)

// GetNd61 New data
func (r *registerFdcan_ndat2Type) GetNd61() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd61Mask) != 0
}

// SetNd61 New data
func (r *registerFdcan_ndat2Type) SetNd61(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd61Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd61Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd62Shift = 30
	RegisterFdcan_ndat2FieldNd62Mask  = 0x40000000
)

// GetNd62 New data
func (r *registerFdcan_ndat2Type) GetNd62() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd62Mask) != 0
}

// SetNd62 New data
func (r *registerFdcan_ndat2Type) SetNd62(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd62Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd62Mask)
	}
}

const (
	RegisterFdcan_ndat2FieldNd63Shift = 31
	RegisterFdcan_ndat2FieldNd63Mask  = 0x80000000
)

// GetNd63 New data
func (r *registerFdcan_ndat2Type) GetNd63() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ndat2FieldNd63Mask) != 0
}

// SetNd63 New data
func (r *registerFdcan_ndat2Type) SetNd63(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ndat2FieldNd63Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ndat2FieldNd63Mask)
	}
}

// registerFdcan_rxf0cType FDCAN Rx FIFO 0 Configuration Register
type registerFdcan_rxf0cType uint32

const (
	RegisterFdcan_rxf0cFieldF0saShift = 2
	RegisterFdcan_rxf0cFieldF0saMask  = 0xfffc
)

// GetF0sa Rx FIFO 0 Start Address
func (r *registerFdcan_rxf0cType) GetF0sa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf0cFieldF0saMask) >> RegisterFdcan_rxf0cFieldF0saShift)
}

// SetF0sa Rx FIFO 0 Start Address
func (r *registerFdcan_rxf0cType) SetF0sa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf0cFieldF0saMask)|(uint32(value)<<RegisterFdcan_rxf0cFieldF0saShift))
}

const (
	RegisterFdcan_rxf0cFieldF0sShift = 16
	RegisterFdcan_rxf0cFieldF0sMask  = 0xff0000
)

// GetF0s Rx FIFO 0 Size
func (r *registerFdcan_rxf0cType) GetF0s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf0cFieldF0sMask) >> RegisterFdcan_rxf0cFieldF0sShift)
}

// SetF0s Rx FIFO 0 Size
func (r *registerFdcan_rxf0cType) SetF0s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf0cFieldF0sMask)|(uint32(value)<<RegisterFdcan_rxf0cFieldF0sShift))
}

const (
	RegisterFdcan_rxf0cFieldF0wmShift = 24
	RegisterFdcan_rxf0cFieldF0wmMask  = 0xff000000
)

// GetF0wm FIFO 0 Watermark
func (r *registerFdcan_rxf0cType) GetF0wm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf0cFieldF0wmMask) >> RegisterFdcan_rxf0cFieldF0wmShift)
}

// SetF0wm FIFO 0 Watermark
func (r *registerFdcan_rxf0cType) SetF0wm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf0cFieldF0wmMask)|(uint32(value)<<RegisterFdcan_rxf0cFieldF0wmShift))
}

// registerFdcan_rxf0sType FDCAN Rx FIFO 0 Status Register
type registerFdcan_rxf0sType uint32

const (
	RegisterFdcan_rxf0sFieldF0flShift = 0
	RegisterFdcan_rxf0sFieldF0flMask  = 0x7f
)

// GetF0fl Rx FIFO 0 Fill Level
func (r *registerFdcan_rxf0sType) GetF0fl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf0sFieldF0flMask) >> RegisterFdcan_rxf0sFieldF0flShift)
}

// SetF0fl Rx FIFO 0 Fill Level
func (r *registerFdcan_rxf0sType) SetF0fl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf0sFieldF0flMask)|(uint32(value)<<RegisterFdcan_rxf0sFieldF0flShift))
}

const (
	RegisterFdcan_rxf0sFieldF0gShift = 8
	RegisterFdcan_rxf0sFieldF0gMask  = 0x3f00
)

// GetF0g Rx FIFO 0 Get Index
func (r *registerFdcan_rxf0sType) GetF0g() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf0sFieldF0gMask) >> RegisterFdcan_rxf0sFieldF0gShift)
}

// SetF0g Rx FIFO 0 Get Index
func (r *registerFdcan_rxf0sType) SetF0g(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf0sFieldF0gMask)|(uint32(value)<<RegisterFdcan_rxf0sFieldF0gShift))
}

const (
	RegisterFdcan_rxf0sFieldF0pShift = 16
	RegisterFdcan_rxf0sFieldF0pMask  = 0x3f0000
)

// GetF0p Rx FIFO 0 Put Index
func (r *registerFdcan_rxf0sType) GetF0p() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf0sFieldF0pMask) >> RegisterFdcan_rxf0sFieldF0pShift)
}

// SetF0p Rx FIFO 0 Put Index
func (r *registerFdcan_rxf0sType) SetF0p(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf0sFieldF0pMask)|(uint32(value)<<RegisterFdcan_rxf0sFieldF0pShift))
}

const (
	RegisterFdcan_rxf0sFieldF0fShift = 24
	RegisterFdcan_rxf0sFieldF0fMask  = 0x1000000
)

// GetF0f Rx FIFO 0 Full
func (r *registerFdcan_rxf0sType) GetF0f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf0sFieldF0fMask) != 0
}

// SetF0f Rx FIFO 0 Full
func (r *registerFdcan_rxf0sType) SetF0f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_rxf0sFieldF0fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf0sFieldF0fMask)
	}
}

const (
	RegisterFdcan_rxf0sFieldRf0lShift = 25
	RegisterFdcan_rxf0sFieldRf0lMask  = 0x2000000
)

// GetRf0l Rx FIFO 0 Message Lost
func (r *registerFdcan_rxf0sType) GetRf0l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf0sFieldRf0lMask) != 0
}

// SetRf0l Rx FIFO 0 Message Lost
func (r *registerFdcan_rxf0sType) SetRf0l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_rxf0sFieldRf0lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf0sFieldRf0lMask)
	}
}

// registerFdcan_rxf0aType CAN Rx FIFO 0 Acknowledge Register
type registerFdcan_rxf0aType uint32

const (
	RegisterFdcan_rxf0aFieldFa01Shift = 0
	RegisterFdcan_rxf0aFieldFa01Mask  = 0x3f
)

// GetFa01 Rx FIFO 0 Acknowledge Index
func (r *registerFdcan_rxf0aType) GetFa01() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf0aFieldFa01Mask) >> RegisterFdcan_rxf0aFieldFa01Shift)
}

// SetFa01 Rx FIFO 0 Acknowledge Index
func (r *registerFdcan_rxf0aType) SetFa01(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf0aFieldFa01Mask)|(uint32(value)<<RegisterFdcan_rxf0aFieldFa01Shift))
}

// registerFdcan_rxbcType FDCAN Rx Buffer Configuration Register
type registerFdcan_rxbcType uint32

const (
	RegisterFdcan_rxbcFieldRbsaShift = 2
	RegisterFdcan_rxbcFieldRbsaMask  = 0xfffc
)

// GetRbsa Rx Buffer Start Address
func (r *registerFdcan_rxbcType) GetRbsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxbcFieldRbsaMask) >> RegisterFdcan_rxbcFieldRbsaShift)
}

// SetRbsa Rx Buffer Start Address
func (r *registerFdcan_rxbcType) SetRbsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxbcFieldRbsaMask)|(uint32(value)<<RegisterFdcan_rxbcFieldRbsaShift))
}

// registerFdcan_rxf1cType FDCAN Rx FIFO 1 Configuration Register
type registerFdcan_rxf1cType uint32

const (
	RegisterFdcan_rxf1cFieldF1saShift = 2
	RegisterFdcan_rxf1cFieldF1saMask  = 0xfffc
)

// GetF1sa Rx FIFO 1 Start Address
func (r *registerFdcan_rxf1cType) GetF1sa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf1cFieldF1saMask) >> RegisterFdcan_rxf1cFieldF1saShift)
}

// SetF1sa Rx FIFO 1 Start Address
func (r *registerFdcan_rxf1cType) SetF1sa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf1cFieldF1saMask)|(uint32(value)<<RegisterFdcan_rxf1cFieldF1saShift))
}

const (
	RegisterFdcan_rxf1cFieldF1sShift = 16
	RegisterFdcan_rxf1cFieldF1sMask  = 0x7f0000
)

// GetF1s Rx FIFO 1 Size
func (r *registerFdcan_rxf1cType) GetF1s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf1cFieldF1sMask) >> RegisterFdcan_rxf1cFieldF1sShift)
}

// SetF1s Rx FIFO 1 Size
func (r *registerFdcan_rxf1cType) SetF1s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf1cFieldF1sMask)|(uint32(value)<<RegisterFdcan_rxf1cFieldF1sShift))
}

const (
	RegisterFdcan_rxf1cFieldF1wmShift = 24
	RegisterFdcan_rxf1cFieldF1wmMask  = 0x7f000000
)

// GetF1wm Rx FIFO 1 Watermark
func (r *registerFdcan_rxf1cType) GetF1wm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf1cFieldF1wmMask) >> RegisterFdcan_rxf1cFieldF1wmShift)
}

// SetF1wm Rx FIFO 1 Watermark
func (r *registerFdcan_rxf1cType) SetF1wm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf1cFieldF1wmMask)|(uint32(value)<<RegisterFdcan_rxf1cFieldF1wmShift))
}

// registerFdcan_rxf1sType FDCAN Rx FIFO 1 Status Register
type registerFdcan_rxf1sType uint32

const (
	RegisterFdcan_rxf1sFieldF1flShift = 0
	RegisterFdcan_rxf1sFieldF1flMask  = 0x7f
)

// GetF1fl Rx FIFO 1 Fill Level
func (r *registerFdcan_rxf1sType) GetF1fl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf1sFieldF1flMask) >> RegisterFdcan_rxf1sFieldF1flShift)
}

// SetF1fl Rx FIFO 1 Fill Level
func (r *registerFdcan_rxf1sType) SetF1fl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf1sFieldF1flMask)|(uint32(value)<<RegisterFdcan_rxf1sFieldF1flShift))
}

const (
	RegisterFdcan_rxf1sFieldF1giShift = 8
	RegisterFdcan_rxf1sFieldF1giMask  = 0x7f00
)

// GetF1gi Rx FIFO 1 Get Index
func (r *registerFdcan_rxf1sType) GetF1gi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf1sFieldF1giMask) >> RegisterFdcan_rxf1sFieldF1giShift)
}

// SetF1gi Rx FIFO 1 Get Index
func (r *registerFdcan_rxf1sType) SetF1gi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf1sFieldF1giMask)|(uint32(value)<<RegisterFdcan_rxf1sFieldF1giShift))
}

const (
	RegisterFdcan_rxf1sFieldF1piShift = 16
	RegisterFdcan_rxf1sFieldF1piMask  = 0x7f0000
)

// GetF1pi Rx FIFO 1 Put Index
func (r *registerFdcan_rxf1sType) GetF1pi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf1sFieldF1piMask) >> RegisterFdcan_rxf1sFieldF1piShift)
}

// SetF1pi Rx FIFO 1 Put Index
func (r *registerFdcan_rxf1sType) SetF1pi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf1sFieldF1piMask)|(uint32(value)<<RegisterFdcan_rxf1sFieldF1piShift))
}

const (
	RegisterFdcan_rxf1sFieldF1fShift = 24
	RegisterFdcan_rxf1sFieldF1fMask  = 0x1000000
)

// GetF1f Rx FIFO 1 Full
func (r *registerFdcan_rxf1sType) GetF1f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf1sFieldF1fMask) != 0
}

// SetF1f Rx FIFO 1 Full
func (r *registerFdcan_rxf1sType) SetF1f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_rxf1sFieldF1fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf1sFieldF1fMask)
	}
}

const (
	RegisterFdcan_rxf1sFieldRf1lShift = 25
	RegisterFdcan_rxf1sFieldRf1lMask  = 0x2000000
)

// GetRf1l Rx FIFO 1 Message Lost
func (r *registerFdcan_rxf1sType) GetRf1l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf1sFieldRf1lMask) != 0
}

// SetRf1l Rx FIFO 1 Message Lost
func (r *registerFdcan_rxf1sType) SetRf1l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_rxf1sFieldRf1lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf1sFieldRf1lMask)
	}
}

const (
	RegisterFdcan_rxf1sFieldDmsShift = 30
	RegisterFdcan_rxf1sFieldDmsMask  = 0xc0000000
)

// GetDms Debug Message Status
func (r *registerFdcan_rxf1sType) GetDms() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf1sFieldDmsMask) >> RegisterFdcan_rxf1sFieldDmsShift)
}

// SetDms Debug Message Status
func (r *registerFdcan_rxf1sType) SetDms(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf1sFieldDmsMask)|(uint32(value)<<RegisterFdcan_rxf1sFieldDmsShift))
}

// registerFdcan_rxf1aType FDCAN Rx FIFO 1 Acknowledge Register
type registerFdcan_rxf1aType uint32

const (
	RegisterFdcan_rxf1aFieldF1aiShift = 0
	RegisterFdcan_rxf1aFieldF1aiMask  = 0x3f
)

// GetF1ai Rx FIFO 1 Acknowledge Index
func (r *registerFdcan_rxf1aType) GetF1ai() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxf1aFieldF1aiMask) >> RegisterFdcan_rxf1aFieldF1aiShift)
}

// SetF1ai Rx FIFO 1 Acknowledge Index
func (r *registerFdcan_rxf1aType) SetF1ai(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxf1aFieldF1aiMask)|(uint32(value)<<RegisterFdcan_rxf1aFieldF1aiShift))
}

// registerFdcan_rxescType FDCAN Rx Buffer Element Size Configuration Register
type registerFdcan_rxescType uint32

const (
	RegisterFdcan_rxescFieldF0dsShift = 0
	RegisterFdcan_rxescFieldF0dsMask  = 0x7
)

// GetF0ds Rx FIFO 1 Data Field Size:
func (r *registerFdcan_rxescType) GetF0ds() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxescFieldF0dsMask) >> RegisterFdcan_rxescFieldF0dsShift)
}

// SetF0ds Rx FIFO 1 Data Field Size:
func (r *registerFdcan_rxescType) SetF0ds(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxescFieldF0dsMask)|(uint32(value)<<RegisterFdcan_rxescFieldF0dsShift))
}

const (
	RegisterFdcan_rxescFieldF1dsShift = 4
	RegisterFdcan_rxescFieldF1dsMask  = 0x70
)

// GetF1ds Rx FIFO 0 Data Field Size:
func (r *registerFdcan_rxescType) GetF1ds() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxescFieldF1dsMask) >> RegisterFdcan_rxescFieldF1dsShift)
}

// SetF1ds Rx FIFO 0 Data Field Size:
func (r *registerFdcan_rxescType) SetF1ds(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxescFieldF1dsMask)|(uint32(value)<<RegisterFdcan_rxescFieldF1dsShift))
}

const (
	RegisterFdcan_rxescFieldRbdsShift = 8
	RegisterFdcan_rxescFieldRbdsMask  = 0x700
)

// GetRbds Rx Buffer Data Field Size:
func (r *registerFdcan_rxescType) GetRbds() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_rxescFieldRbdsMask) >> RegisterFdcan_rxescFieldRbdsShift)
}

// SetRbds Rx Buffer Data Field Size:
func (r *registerFdcan_rxescType) SetRbds(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_rxescFieldRbdsMask)|(uint32(value)<<RegisterFdcan_rxescFieldRbdsShift))
}

// registerFdcan_txbcType FDCAN Tx Buffer Configuration Register
type registerFdcan_txbcType uint32

const (
	RegisterFdcan_txbcFieldTbsaShift = 2
	RegisterFdcan_txbcFieldTbsaMask  = 0xfffc
)

// GetTbsa Tx Buffers Start Address
func (r *registerFdcan_txbcType) GetTbsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txbcFieldTbsaMask) >> RegisterFdcan_txbcFieldTbsaShift)
}

// SetTbsa Tx Buffers Start Address
func (r *registerFdcan_txbcType) SetTbsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txbcFieldTbsaMask)|(uint32(value)<<RegisterFdcan_txbcFieldTbsaShift))
}

const (
	RegisterFdcan_txbcFieldNdtbShift = 16
	RegisterFdcan_txbcFieldNdtbMask  = 0x3f0000
)

// GetNdtb Number of Dedicated Transmit Buffers
func (r *registerFdcan_txbcType) GetNdtb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txbcFieldNdtbMask) >> RegisterFdcan_txbcFieldNdtbShift)
}

// SetNdtb Number of Dedicated Transmit Buffers
func (r *registerFdcan_txbcType) SetNdtb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txbcFieldNdtbMask)|(uint32(value)<<RegisterFdcan_txbcFieldNdtbShift))
}

const (
	RegisterFdcan_txbcFieldTfqsShift = 24
	RegisterFdcan_txbcFieldTfqsMask  = 0x3f000000
)

// GetTfqs Transmit FIFO/Queue Size
func (r *registerFdcan_txbcType) GetTfqs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txbcFieldTfqsMask) >> RegisterFdcan_txbcFieldTfqsShift)
}

// SetTfqs Transmit FIFO/Queue Size
func (r *registerFdcan_txbcType) SetTfqs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txbcFieldTfqsMask)|(uint32(value)<<RegisterFdcan_txbcFieldTfqsShift))
}

const (
	RegisterFdcan_txbcFieldTfqmShift = 30
	RegisterFdcan_txbcFieldTfqmMask  = 0x40000000
)

// GetTfqm Tx FIFO/Queue Mode
func (r *registerFdcan_txbcType) GetTfqm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txbcFieldTfqmMask) != 0
}

// SetTfqm Tx FIFO/Queue Mode
func (r *registerFdcan_txbcType) SetTfqm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_txbcFieldTfqmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txbcFieldTfqmMask)
	}
}

// registerFdcan_txfqsType FDCAN Tx FIFO/Queue Status Register
type registerFdcan_txfqsType uint32

const (
	RegisterFdcan_txfqsFieldTfflShift = 0
	RegisterFdcan_txfqsFieldTfflMask  = 0x3f
)

// GetTffl Tx FIFO Free Level
func (r *registerFdcan_txfqsType) GetTffl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txfqsFieldTfflMask) >> RegisterFdcan_txfqsFieldTfflShift)
}

// SetTffl Tx FIFO Free Level
func (r *registerFdcan_txfqsType) SetTffl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txfqsFieldTfflMask)|(uint32(value)<<RegisterFdcan_txfqsFieldTfflShift))
}

const (
	RegisterFdcan_txfqsFieldTfgiShift = 8
	RegisterFdcan_txfqsFieldTfgiMask  = 0x1f00
)

// GetTfgi TFGI
func (r *registerFdcan_txfqsType) GetTfgi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txfqsFieldTfgiMask) >> RegisterFdcan_txfqsFieldTfgiShift)
}

// SetTfgi TFGI
func (r *registerFdcan_txfqsType) SetTfgi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txfqsFieldTfgiMask)|(uint32(value)<<RegisterFdcan_txfqsFieldTfgiShift))
}

const (
	RegisterFdcan_txfqsFieldTfqpiShift = 16
	RegisterFdcan_txfqsFieldTfqpiMask  = 0x1f0000
)

// GetTfqpi Tx FIFO/Queue Put Index
func (r *registerFdcan_txfqsType) GetTfqpi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txfqsFieldTfqpiMask) >> RegisterFdcan_txfqsFieldTfqpiShift)
}

// SetTfqpi Tx FIFO/Queue Put Index
func (r *registerFdcan_txfqsType) SetTfqpi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txfqsFieldTfqpiMask)|(uint32(value)<<RegisterFdcan_txfqsFieldTfqpiShift))
}

const (
	RegisterFdcan_txfqsFieldTfqfShift = 21
	RegisterFdcan_txfqsFieldTfqfMask  = 0x200000
)

// GetTfqf Tx FIFO/Queue Full
func (r *registerFdcan_txfqsType) GetTfqf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txfqsFieldTfqfMask) != 0
}

// SetTfqf Tx FIFO/Queue Full
func (r *registerFdcan_txfqsType) SetTfqf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_txfqsFieldTfqfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txfqsFieldTfqfMask)
	}
}

// registerFdcan_txescType FDCAN Tx Buffer Element Size Configuration Register
type registerFdcan_txescType uint32

const (
	RegisterFdcan_txescFieldTbdsShift = 0
	RegisterFdcan_txescFieldTbdsMask  = 0x7
)

// GetTbds Tx Buffer Data Field Size:
func (r *registerFdcan_txescType) GetTbds() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txescFieldTbdsMask) >> RegisterFdcan_txescFieldTbdsShift)
}

// SetTbds Tx Buffer Data Field Size:
func (r *registerFdcan_txescType) SetTbds(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txescFieldTbdsMask)|(uint32(value)<<RegisterFdcan_txescFieldTbdsShift))
}

// registerFdcan_txbrpType FDCAN Tx Buffer Request Pending Register
type registerFdcan_txbrpType uint32

const (
	RegisterFdcan_txbrpFieldTrpShift = 0
	RegisterFdcan_txbrpFieldTrpMask  = 0xffffffff
)

// GetTrp Transmission Request Pending
func (r *registerFdcan_txbrpType) GetTrp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txbrpFieldTrpMask) >> RegisterFdcan_txbrpFieldTrpShift)
}

// SetTrp Transmission Request Pending
func (r *registerFdcan_txbrpType) SetTrp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txbrpFieldTrpMask)|(uint32(value)<<RegisterFdcan_txbrpFieldTrpShift))
}

// registerFdcan_txbarType FDCAN Tx Buffer Add Request Register
type registerFdcan_txbarType uint32

const (
	RegisterFdcan_txbarFieldArShift = 0
	RegisterFdcan_txbarFieldArMask  = 0xffffffff
)

// GetAr Add Request
func (r *registerFdcan_txbarType) GetAr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txbarFieldArMask) >> RegisterFdcan_txbarFieldArShift)
}

// SetAr Add Request
func (r *registerFdcan_txbarType) SetAr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txbarFieldArMask)|(uint32(value)<<RegisterFdcan_txbarFieldArShift))
}

// registerFdcan_txbcrType FDCAN Tx Buffer Cancellation Request Register
type registerFdcan_txbcrType uint32

const (
	RegisterFdcan_txbcrFieldCrShift = 0
	RegisterFdcan_txbcrFieldCrMask  = 0xffffffff
)

// GetCr Cancellation Request
func (r *registerFdcan_txbcrType) GetCr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txbcrFieldCrMask) >> RegisterFdcan_txbcrFieldCrShift)
}

// SetCr Cancellation Request
func (r *registerFdcan_txbcrType) SetCr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txbcrFieldCrMask)|(uint32(value)<<RegisterFdcan_txbcrFieldCrShift))
}

// registerFdcan_txbtoType FDCAN Tx Buffer Transmission Occurred Register
type registerFdcan_txbtoType uint32

const (
	RegisterFdcan_txbtoFieldToShift = 0
	RegisterFdcan_txbtoFieldToMask  = 0xffffffff
)

// GetTo Transmission Occurred.
func (r *registerFdcan_txbtoType) GetTo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txbtoFieldToMask) >> RegisterFdcan_txbtoFieldToShift)
}

// SetTo Transmission Occurred.
func (r *registerFdcan_txbtoType) SetTo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txbtoFieldToMask)|(uint32(value)<<RegisterFdcan_txbtoFieldToShift))
}

// registerFdcan_txbcfType FDCAN Tx Buffer Cancellation Finished Register
type registerFdcan_txbcfType uint32

const (
	RegisterFdcan_txbcfFieldCfShift = 0
	RegisterFdcan_txbcfFieldCfMask  = 0xffffffff
)

// GetCf Cancellation Finished
func (r *registerFdcan_txbcfType) GetCf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txbcfFieldCfMask) >> RegisterFdcan_txbcfFieldCfShift)
}

// SetCf Cancellation Finished
func (r *registerFdcan_txbcfType) SetCf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txbcfFieldCfMask)|(uint32(value)<<RegisterFdcan_txbcfFieldCfShift))
}

// registerFdcan_txbtieType FDCAN Tx Buffer Transmission Interrupt Enable Register
type registerFdcan_txbtieType uint32

const (
	RegisterFdcan_txbtieFieldTieShift = 0
	RegisterFdcan_txbtieFieldTieMask  = 0xffffffff
)

// GetTie Transmission Interrupt Enable
func (r *registerFdcan_txbtieType) GetTie() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txbtieFieldTieMask) >> RegisterFdcan_txbtieFieldTieShift)
}

// SetTie Transmission Interrupt Enable
func (r *registerFdcan_txbtieType) SetTie(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txbtieFieldTieMask)|(uint32(value)<<RegisterFdcan_txbtieFieldTieShift))
}

// registerFdcan_txbcieType FDCAN Tx Buffer Cancellation Finished Interrupt Enable Register
type registerFdcan_txbcieType uint32

const (
	RegisterFdcan_txbcieFieldCfShift = 0
	RegisterFdcan_txbcieFieldCfMask  = 0xffffffff
)

// GetCf Cancellation Finished Interrupt Enable
func (r *registerFdcan_txbcieType) GetCf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txbcieFieldCfMask) >> RegisterFdcan_txbcieFieldCfShift)
}

// SetCf Cancellation Finished Interrupt Enable
func (r *registerFdcan_txbcieType) SetCf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txbcieFieldCfMask)|(uint32(value)<<RegisterFdcan_txbcieFieldCfShift))
}

// registerFdcan_txefcType FDCAN Tx Event FIFO Configuration Register
type registerFdcan_txefcType uint32

const (
	RegisterFdcan_txefcFieldEfsaShift = 2
	RegisterFdcan_txefcFieldEfsaMask  = 0xfffc
)

// GetEfsa Event FIFO Start Address
func (r *registerFdcan_txefcType) GetEfsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txefcFieldEfsaMask) >> RegisterFdcan_txefcFieldEfsaShift)
}

// SetEfsa Event FIFO Start Address
func (r *registerFdcan_txefcType) SetEfsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txefcFieldEfsaMask)|(uint32(value)<<RegisterFdcan_txefcFieldEfsaShift))
}

const (
	RegisterFdcan_txefcFieldEfsShift = 16
	RegisterFdcan_txefcFieldEfsMask  = 0x3f0000
)

// GetEfs Event FIFO Size
func (r *registerFdcan_txefcType) GetEfs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txefcFieldEfsMask) >> RegisterFdcan_txefcFieldEfsShift)
}

// SetEfs Event FIFO Size
func (r *registerFdcan_txefcType) SetEfs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txefcFieldEfsMask)|(uint32(value)<<RegisterFdcan_txefcFieldEfsShift))
}

const (
	RegisterFdcan_txefcFieldEfwmShift = 24
	RegisterFdcan_txefcFieldEfwmMask  = 0x3f000000
)

// GetEfwm Event FIFO Watermark
func (r *registerFdcan_txefcType) GetEfwm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txefcFieldEfwmMask) >> RegisterFdcan_txefcFieldEfwmShift)
}

// SetEfwm Event FIFO Watermark
func (r *registerFdcan_txefcType) SetEfwm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txefcFieldEfwmMask)|(uint32(value)<<RegisterFdcan_txefcFieldEfwmShift))
}

// registerFdcan_txefsType FDCAN Tx Event FIFO Status Register
type registerFdcan_txefsType uint32

const (
	RegisterFdcan_txefsFieldEfflShift = 0
	RegisterFdcan_txefsFieldEfflMask  = 0x3f
)

// GetEffl Event FIFO Fill Level
func (r *registerFdcan_txefsType) GetEffl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txefsFieldEfflMask) >> RegisterFdcan_txefsFieldEfflShift)
}

// SetEffl Event FIFO Fill Level
func (r *registerFdcan_txefsType) SetEffl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txefsFieldEfflMask)|(uint32(value)<<RegisterFdcan_txefsFieldEfflShift))
}

const (
	RegisterFdcan_txefsFieldEfgiShift = 8
	RegisterFdcan_txefsFieldEfgiMask  = 0x1f00
)

// GetEfgi Event FIFO Get Index.
func (r *registerFdcan_txefsType) GetEfgi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txefsFieldEfgiMask) >> RegisterFdcan_txefsFieldEfgiShift)
}

// SetEfgi Event FIFO Get Index.
func (r *registerFdcan_txefsType) SetEfgi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txefsFieldEfgiMask)|(uint32(value)<<RegisterFdcan_txefsFieldEfgiShift))
}

const (
	RegisterFdcan_txefsFieldEfpiShift = 16
	RegisterFdcan_txefsFieldEfpiMask  = 0x1f0000
)

// GetEfpi Event FIFO put index.
func (r *registerFdcan_txefsType) GetEfpi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txefsFieldEfpiMask) >> RegisterFdcan_txefsFieldEfpiShift)
}

// SetEfpi Event FIFO put index.
func (r *registerFdcan_txefsType) SetEfpi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txefsFieldEfpiMask)|(uint32(value)<<RegisterFdcan_txefsFieldEfpiShift))
}

const (
	RegisterFdcan_txefsFieldEffShift = 24
	RegisterFdcan_txefsFieldEffMask  = 0x1000000
)

// GetEff Event FIFO Full.
func (r *registerFdcan_txefsType) GetEff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txefsFieldEffMask) != 0
}

// SetEff Event FIFO Full.
func (r *registerFdcan_txefsType) SetEff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_txefsFieldEffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txefsFieldEffMask)
	}
}

const (
	RegisterFdcan_txefsFieldTeflShift = 25
	RegisterFdcan_txefsFieldTeflMask  = 0x2000000
)

// GetTefl Tx Event FIFO Element Lost.
func (r *registerFdcan_txefsType) GetTefl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txefsFieldTeflMask) != 0
}

// SetTefl Tx Event FIFO Element Lost.
func (r *registerFdcan_txefsType) SetTefl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_txefsFieldTeflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txefsFieldTeflMask)
	}
}

// registerFdcan_txefaType FDCAN Tx Event FIFO Acknowledge Register
type registerFdcan_txefaType uint32

const (
	RegisterFdcan_txefaFieldEfaiShift = 0
	RegisterFdcan_txefaFieldEfaiMask  = 0x1f
)

// GetEfai Event FIFO Acknowledge Index
func (r *registerFdcan_txefaType) GetEfai() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_txefaFieldEfaiMask) >> RegisterFdcan_txefaFieldEfaiShift)
}

// SetEfai Event FIFO Acknowledge Index
func (r *registerFdcan_txefaType) SetEfai(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_txefaFieldEfaiMask)|(uint32(value)<<RegisterFdcan_txefaFieldEfaiShift))
}

// registerFdcan_tttmcType FDCAN TT Trigger Memory Configuration Register
type registerFdcan_tttmcType uint32

const (
	RegisterFdcan_tttmcFieldTmsaShift = 2
	RegisterFdcan_tttmcFieldTmsaMask  = 0xfffc
)

// GetTmsa Trigger Memory Start Address
func (r *registerFdcan_tttmcType) GetTmsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tttmcFieldTmsaMask) >> RegisterFdcan_tttmcFieldTmsaShift)
}

// SetTmsa Trigger Memory Start Address
func (r *registerFdcan_tttmcType) SetTmsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tttmcFieldTmsaMask)|(uint32(value)<<RegisterFdcan_tttmcFieldTmsaShift))
}

const (
	RegisterFdcan_tttmcFieldTmeShift = 16
	RegisterFdcan_tttmcFieldTmeMask  = 0x7f0000
)

// GetTme Trigger Memory Elements
func (r *registerFdcan_tttmcType) GetTme() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tttmcFieldTmeMask) >> RegisterFdcan_tttmcFieldTmeShift)
}

// SetTme Trigger Memory Elements
func (r *registerFdcan_tttmcType) SetTme(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tttmcFieldTmeMask)|(uint32(value)<<RegisterFdcan_tttmcFieldTmeShift))
}

// registerFdcan_ttrmcType FDCAN TT Reference Message Configuration Register
type registerFdcan_ttrmcType uint32

const (
	RegisterFdcan_ttrmcFieldRidShift = 0
	RegisterFdcan_ttrmcFieldRidMask  = 0x1fffffff
)

// GetRid Reference Identifier.
func (r *registerFdcan_ttrmcType) GetRid() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttrmcFieldRidMask) >> RegisterFdcan_ttrmcFieldRidShift)
}

// SetRid Reference Identifier.
func (r *registerFdcan_ttrmcType) SetRid(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttrmcFieldRidMask)|(uint32(value)<<RegisterFdcan_ttrmcFieldRidShift))
}

const (
	RegisterFdcan_ttrmcFieldXtdShift = 30
	RegisterFdcan_ttrmcFieldXtdMask  = 0x40000000
)

// GetXtd Extended Identifier
func (r *registerFdcan_ttrmcType) GetXtd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttrmcFieldXtdMask) != 0
}

// SetXtd Extended Identifier
func (r *registerFdcan_ttrmcType) SetXtd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttrmcFieldXtdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttrmcFieldXtdMask)
	}
}

const (
	RegisterFdcan_ttrmcFieldRmpsShift = 31
	RegisterFdcan_ttrmcFieldRmpsMask  = 0x80000000
)

// GetRmps Reference Message Payload Select
func (r *registerFdcan_ttrmcType) GetRmps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttrmcFieldRmpsMask) != 0
}

// SetRmps Reference Message Payload Select
func (r *registerFdcan_ttrmcType) SetRmps(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttrmcFieldRmpsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttrmcFieldRmpsMask)
	}
}

// registerFdcan_ttocfType FDCAN TT Operation Configuration Register
type registerFdcan_ttocfType uint32

const (
	RegisterFdcan_ttocfFieldOmShift = 0
	RegisterFdcan_ttocfFieldOmMask  = 0x3
)

// GetOm Operation Mode
func (r *registerFdcan_ttocfType) GetOm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocfFieldOmMask) >> RegisterFdcan_ttocfFieldOmShift)
}

// SetOm Operation Mode
func (r *registerFdcan_ttocfType) SetOm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocfFieldOmMask)|(uint32(value)<<RegisterFdcan_ttocfFieldOmShift))
}

const (
	RegisterFdcan_ttocfFieldGenShift = 3
	RegisterFdcan_ttocfFieldGenMask  = 0x8
)

// GetGen Gap Enable
func (r *registerFdcan_ttocfType) GetGen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocfFieldGenMask) != 0
}

// SetGen Gap Enable
func (r *registerFdcan_ttocfType) SetGen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocfFieldGenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocfFieldGenMask)
	}
}

const (
	RegisterFdcan_ttocfFieldTmShift = 4
	RegisterFdcan_ttocfFieldTmMask  = 0x10
)

// GetTm Time Master
func (r *registerFdcan_ttocfType) GetTm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocfFieldTmMask) != 0
}

// SetTm Time Master
func (r *registerFdcan_ttocfType) SetTm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocfFieldTmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocfFieldTmMask)
	}
}

const (
	RegisterFdcan_ttocfFieldLdsdlShift = 5
	RegisterFdcan_ttocfFieldLdsdlMask  = 0xe0
)

// GetLdsdl LD of Synchronization Deviation Limit
func (r *registerFdcan_ttocfType) GetLdsdl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocfFieldLdsdlMask) >> RegisterFdcan_ttocfFieldLdsdlShift)
}

// SetLdsdl LD of Synchronization Deviation Limit
func (r *registerFdcan_ttocfType) SetLdsdl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocfFieldLdsdlMask)|(uint32(value)<<RegisterFdcan_ttocfFieldLdsdlShift))
}

const (
	RegisterFdcan_ttocfFieldIrtoShift = 8
	RegisterFdcan_ttocfFieldIrtoMask  = 0x7f00
)

// GetIrto Initial Reference Trigger Offset
func (r *registerFdcan_ttocfType) GetIrto() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocfFieldIrtoMask) >> RegisterFdcan_ttocfFieldIrtoShift)
}

// SetIrto Initial Reference Trigger Offset
func (r *registerFdcan_ttocfType) SetIrto(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocfFieldIrtoMask)|(uint32(value)<<RegisterFdcan_ttocfFieldIrtoShift))
}

const (
	RegisterFdcan_ttocfFieldEecsShift = 15
	RegisterFdcan_ttocfFieldEecsMask  = 0x8000
)

// GetEecs Enable External Clock Synchronization
func (r *registerFdcan_ttocfType) GetEecs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocfFieldEecsMask) != 0
}

// SetEecs Enable External Clock Synchronization
func (r *registerFdcan_ttocfType) SetEecs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocfFieldEecsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocfFieldEecsMask)
	}
}

const (
	RegisterFdcan_ttocfFieldAwlShift = 16
	RegisterFdcan_ttocfFieldAwlMask  = 0xff0000
)

// GetAwl Application Watchdog Limit
func (r *registerFdcan_ttocfType) GetAwl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocfFieldAwlMask) >> RegisterFdcan_ttocfFieldAwlShift)
}

// SetAwl Application Watchdog Limit
func (r *registerFdcan_ttocfType) SetAwl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocfFieldAwlMask)|(uint32(value)<<RegisterFdcan_ttocfFieldAwlShift))
}

const (
	RegisterFdcan_ttocfFieldEgtfShift = 24
	RegisterFdcan_ttocfFieldEgtfMask  = 0x1000000
)

// GetEgtf Enable Global Time Filtering
func (r *registerFdcan_ttocfType) GetEgtf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocfFieldEgtfMask) != 0
}

// SetEgtf Enable Global Time Filtering
func (r *registerFdcan_ttocfType) SetEgtf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocfFieldEgtfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocfFieldEgtfMask)
	}
}

const (
	RegisterFdcan_ttocfFieldEccShift = 25
	RegisterFdcan_ttocfFieldEccMask  = 0x2000000
)

// GetEcc Enable Clock Calibration
func (r *registerFdcan_ttocfType) GetEcc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocfFieldEccMask) != 0
}

// SetEcc Enable Clock Calibration
func (r *registerFdcan_ttocfType) SetEcc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocfFieldEccMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocfFieldEccMask)
	}
}

const (
	RegisterFdcan_ttocfFieldEvtpShift = 26
	RegisterFdcan_ttocfFieldEvtpMask  = 0x4000000
)

// GetEvtp Event Trigger Polarity
func (r *registerFdcan_ttocfType) GetEvtp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocfFieldEvtpMask) != 0
}

// SetEvtp Event Trigger Polarity
func (r *registerFdcan_ttocfType) SetEvtp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocfFieldEvtpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocfFieldEvtpMask)
	}
}

// registerFdcan_ttmlmType FDCAN TT Matrix Limits Register
type registerFdcan_ttmlmType uint32

const (
	RegisterFdcan_ttmlmFieldCcmShift = 0
	RegisterFdcan_ttmlmFieldCcmMask  = 0x3f
)

// GetCcm Cycle Count Max
func (r *registerFdcan_ttmlmType) GetCcm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttmlmFieldCcmMask) >> RegisterFdcan_ttmlmFieldCcmShift)
}

// SetCcm Cycle Count Max
func (r *registerFdcan_ttmlmType) SetCcm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttmlmFieldCcmMask)|(uint32(value)<<RegisterFdcan_ttmlmFieldCcmShift))
}

const (
	RegisterFdcan_ttmlmFieldCssShift = 6
	RegisterFdcan_ttmlmFieldCssMask  = 0xc0
)

// GetCss Cycle Start Synchronization
func (r *registerFdcan_ttmlmType) GetCss() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttmlmFieldCssMask) >> RegisterFdcan_ttmlmFieldCssShift)
}

// SetCss Cycle Start Synchronization
func (r *registerFdcan_ttmlmType) SetCss(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttmlmFieldCssMask)|(uint32(value)<<RegisterFdcan_ttmlmFieldCssShift))
}

const (
	RegisterFdcan_ttmlmFieldTxewShift = 8
	RegisterFdcan_ttmlmFieldTxewMask  = 0xf00
)

// GetTxew Tx Enable Window
func (r *registerFdcan_ttmlmType) GetTxew() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttmlmFieldTxewMask) >> RegisterFdcan_ttmlmFieldTxewShift)
}

// SetTxew Tx Enable Window
func (r *registerFdcan_ttmlmType) SetTxew(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttmlmFieldTxewMask)|(uint32(value)<<RegisterFdcan_ttmlmFieldTxewShift))
}

const (
	RegisterFdcan_ttmlmFieldEnttShift = 16
	RegisterFdcan_ttmlmFieldEnttMask  = 0xfff0000
)

// GetEntt Expected Number of Tx Triggers
func (r *registerFdcan_ttmlmType) GetEntt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttmlmFieldEnttMask) >> RegisterFdcan_ttmlmFieldEnttShift)
}

// SetEntt Expected Number of Tx Triggers
func (r *registerFdcan_ttmlmType) SetEntt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttmlmFieldEnttMask)|(uint32(value)<<RegisterFdcan_ttmlmFieldEnttShift))
}

// registerFdcan_turcfType FDCAN TUR Configuration Register
type registerFdcan_turcfType uint32

const (
	RegisterFdcan_turcfFieldNclShift = 0
	RegisterFdcan_turcfFieldNclMask  = 0xffff
)

// GetNcl Numerator Configuration Low.
func (r *registerFdcan_turcfType) GetNcl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_turcfFieldNclMask) >> RegisterFdcan_turcfFieldNclShift)
}

// SetNcl Numerator Configuration Low.
func (r *registerFdcan_turcfType) SetNcl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_turcfFieldNclMask)|(uint32(value)<<RegisterFdcan_turcfFieldNclShift))
}

const (
	RegisterFdcan_turcfFieldDcShift = 16
	RegisterFdcan_turcfFieldDcMask  = 0x3fff0000
)

// GetDc Denominator Configuration.
func (r *registerFdcan_turcfType) GetDc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_turcfFieldDcMask) >> RegisterFdcan_turcfFieldDcShift)
}

// SetDc Denominator Configuration.
func (r *registerFdcan_turcfType) SetDc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_turcfFieldDcMask)|(uint32(value)<<RegisterFdcan_turcfFieldDcShift))
}

const (
	RegisterFdcan_turcfFieldEltShift = 31
	RegisterFdcan_turcfFieldEltMask  = 0x80000000
)

// GetElt Enable Local Time
func (r *registerFdcan_turcfType) GetElt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_turcfFieldEltMask) != 0
}

// SetElt Enable Local Time
func (r *registerFdcan_turcfType) SetElt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_turcfFieldEltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_turcfFieldEltMask)
	}
}

// registerFdcan_ttocnType FDCAN TT Operation Control Register
type registerFdcan_ttocnType uint32

const (
	RegisterFdcan_ttocnFieldSgtShift = 0
	RegisterFdcan_ttocnFieldSgtMask  = 0x1
)

// GetSgt Set Global time
func (r *registerFdcan_ttocnType) GetSgt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldSgtMask) != 0
}

// SetSgt Set Global time
func (r *registerFdcan_ttocnType) SetSgt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocnFieldSgtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldSgtMask)
	}
}

const (
	RegisterFdcan_ttocnFieldEcsShift = 1
	RegisterFdcan_ttocnFieldEcsMask  = 0x2
)

// GetEcs External Clock Synchronization
func (r *registerFdcan_ttocnType) GetEcs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldEcsMask) != 0
}

// SetEcs External Clock Synchronization
func (r *registerFdcan_ttocnType) SetEcs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocnFieldEcsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldEcsMask)
	}
}

const (
	RegisterFdcan_ttocnFieldSwpShift = 2
	RegisterFdcan_ttocnFieldSwpMask  = 0x4
)

// GetSwp Stop Watch Polarity
func (r *registerFdcan_ttocnType) GetSwp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldSwpMask) != 0
}

// SetSwp Stop Watch Polarity
func (r *registerFdcan_ttocnType) SetSwp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocnFieldSwpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldSwpMask)
	}
}

const (
	RegisterFdcan_ttocnFieldSwsShift = 3
	RegisterFdcan_ttocnFieldSwsMask  = 0x18
)

// GetSws Stop Watch Source.
func (r *registerFdcan_ttocnType) GetSws() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldSwsMask) >> RegisterFdcan_ttocnFieldSwsShift)
}

// SetSws Stop Watch Source.
func (r *registerFdcan_ttocnType) SetSws(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldSwsMask)|(uint32(value)<<RegisterFdcan_ttocnFieldSwsShift))
}

const (
	RegisterFdcan_ttocnFieldRtieShift = 5
	RegisterFdcan_ttocnFieldRtieMask  = 0x20
)

// GetRtie Register Time Mark Interrupt Pulse Enable
func (r *registerFdcan_ttocnType) GetRtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldRtieMask) != 0
}

// SetRtie Register Time Mark Interrupt Pulse Enable
func (r *registerFdcan_ttocnType) SetRtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocnFieldRtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldRtieMask)
	}
}

const (
	RegisterFdcan_ttocnFieldTmcShift = 6
	RegisterFdcan_ttocnFieldTmcMask  = 0xc0
)

// GetTmc Register Time Mark Compare
func (r *registerFdcan_ttocnType) GetTmc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldTmcMask) >> RegisterFdcan_ttocnFieldTmcShift)
}

// SetTmc Register Time Mark Compare
func (r *registerFdcan_ttocnType) SetTmc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldTmcMask)|(uint32(value)<<RegisterFdcan_ttocnFieldTmcShift))
}

const (
	RegisterFdcan_ttocnFieldTtieShift = 8
	RegisterFdcan_ttocnFieldTtieMask  = 0x100
)

// GetTtie Trigger Time Mark Interrupt Pulse Enable
func (r *registerFdcan_ttocnType) GetTtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldTtieMask) != 0
}

// SetTtie Trigger Time Mark Interrupt Pulse Enable
func (r *registerFdcan_ttocnType) SetTtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocnFieldTtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldTtieMask)
	}
}

const (
	RegisterFdcan_ttocnFieldGcsShift = 9
	RegisterFdcan_ttocnFieldGcsMask  = 0x200
)

// GetGcs Gap Control Select
func (r *registerFdcan_ttocnType) GetGcs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldGcsMask) != 0
}

// SetGcs Gap Control Select
func (r *registerFdcan_ttocnType) SetGcs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocnFieldGcsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldGcsMask)
	}
}

const (
	RegisterFdcan_ttocnFieldFgpShift = 10
	RegisterFdcan_ttocnFieldFgpMask  = 0x400
)

// GetFgp Finish Gap.
func (r *registerFdcan_ttocnType) GetFgp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldFgpMask) != 0
}

// SetFgp Finish Gap.
func (r *registerFdcan_ttocnType) SetFgp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocnFieldFgpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldFgpMask)
	}
}

const (
	RegisterFdcan_ttocnFieldTmgShift = 11
	RegisterFdcan_ttocnFieldTmgMask  = 0x800
)

// GetTmg Time Mark Gap
func (r *registerFdcan_ttocnType) GetTmg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldTmgMask) != 0
}

// SetTmg Time Mark Gap
func (r *registerFdcan_ttocnType) SetTmg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocnFieldTmgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldTmgMask)
	}
}

const (
	RegisterFdcan_ttocnFieldNigShift = 12
	RegisterFdcan_ttocnFieldNigMask  = 0x1000
)

// GetNig Next is Gap
func (r *registerFdcan_ttocnType) GetNig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldNigMask) != 0
}

// SetNig Next is Gap
func (r *registerFdcan_ttocnType) SetNig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocnFieldNigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldNigMask)
	}
}

const (
	RegisterFdcan_ttocnFieldEscnShift = 13
	RegisterFdcan_ttocnFieldEscnMask  = 0x2000
)

// GetEscn External Synchronization Control
func (r *registerFdcan_ttocnType) GetEscn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldEscnMask) != 0
}

// SetEscn External Synchronization Control
func (r *registerFdcan_ttocnType) SetEscn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocnFieldEscnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldEscnMask)
	}
}

const (
	RegisterFdcan_ttocnFieldLckcShift = 15
	RegisterFdcan_ttocnFieldLckcMask  = 0x8000
)

// GetLckc TT Operation Control Register Locked
func (r *registerFdcan_ttocnType) GetLckc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttocnFieldLckcMask) != 0
}

// SetLckc TT Operation Control Register Locked
func (r *registerFdcan_ttocnType) SetLckc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttocnFieldLckcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttocnFieldLckcMask)
	}
}

// registerCan_ttgtpType FDCAN TT Global Time Preset Register
type registerCan_ttgtpType uint32

const (
	RegisterCan_ttgtpFieldNclShift = 0
	RegisterCan_ttgtpFieldNclMask  = 0xffff
)

// GetNcl Time Preset
func (r *registerCan_ttgtpType) GetNcl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCan_ttgtpFieldNclMask) >> RegisterCan_ttgtpFieldNclShift)
}

// SetNcl Time Preset
func (r *registerCan_ttgtpType) SetNcl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCan_ttgtpFieldNclMask)|(uint32(value)<<RegisterCan_ttgtpFieldNclShift))
}

const (
	RegisterCan_ttgtpFieldCtpShift = 16
	RegisterCan_ttgtpFieldCtpMask  = 0xffff0000
)

// GetCtp Cycle Time Target Phase
func (r *registerCan_ttgtpType) GetCtp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCan_ttgtpFieldCtpMask) >> RegisterCan_ttgtpFieldCtpShift)
}

// SetCtp Cycle Time Target Phase
func (r *registerCan_ttgtpType) SetCtp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCan_ttgtpFieldCtpMask)|(uint32(value)<<RegisterCan_ttgtpFieldCtpShift))
}

// registerFdcan_tttmkType FDCAN TT Time Mark Register
type registerFdcan_tttmkType uint32

const (
	RegisterFdcan_tttmkFieldTmShift = 0
	RegisterFdcan_tttmkFieldTmMask  = 0xffff
)

// GetTm Time Mark
func (r *registerFdcan_tttmkType) GetTm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tttmkFieldTmMask) >> RegisterFdcan_tttmkFieldTmShift)
}

// SetTm Time Mark
func (r *registerFdcan_tttmkType) SetTm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tttmkFieldTmMask)|(uint32(value)<<RegisterFdcan_tttmkFieldTmShift))
}

const (
	RegisterFdcan_tttmkFieldTiccShift = 16
	RegisterFdcan_tttmkFieldTiccMask  = 0x7f0000
)

// GetTicc Time Mark Cycle Code
func (r *registerFdcan_tttmkType) GetTicc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tttmkFieldTiccMask) >> RegisterFdcan_tttmkFieldTiccShift)
}

// SetTicc Time Mark Cycle Code
func (r *registerFdcan_tttmkType) SetTicc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tttmkFieldTiccMask)|(uint32(value)<<RegisterFdcan_tttmkFieldTiccShift))
}

const (
	RegisterFdcan_tttmkFieldLckmShift = 31
	RegisterFdcan_tttmkFieldLckmMask  = 0x80000000
)

// GetLckm TT Time Mark Register Locked
func (r *registerFdcan_tttmkType) GetLckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tttmkFieldLckmMask) != 0
}

// SetLckm TT Time Mark Register Locked
func (r *registerFdcan_tttmkType) SetLckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_tttmkFieldLckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tttmkFieldLckmMask)
	}
}

// registerFdcan_ttirType FDCAN TT Interrupt Register
type registerFdcan_ttirType uint32

const (
	RegisterFdcan_ttirFieldSbcShift = 0
	RegisterFdcan_ttirFieldSbcMask  = 0x1
)

// GetSbc Start of Basic Cycle
func (r *registerFdcan_ttirType) GetSbc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldSbcMask) != 0
}

// SetSbc Start of Basic Cycle
func (r *registerFdcan_ttirType) SetSbc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldSbcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldSbcMask)
	}
}

const (
	RegisterFdcan_ttirFieldSmcShift = 1
	RegisterFdcan_ttirFieldSmcMask  = 0x2
)

// GetSmc Start of Matrix Cycle
func (r *registerFdcan_ttirType) GetSmc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldSmcMask) != 0
}

// SetSmc Start of Matrix Cycle
func (r *registerFdcan_ttirType) SetSmc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldSmcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldSmcMask)
	}
}

const (
	RegisterFdcan_ttirFieldCsmShift = 2
	RegisterFdcan_ttirFieldCsmMask  = 0x4
)

// GetCsm Change of Synchronization Mode
func (r *registerFdcan_ttirType) GetCsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldCsmMask) != 0
}

// SetCsm Change of Synchronization Mode
func (r *registerFdcan_ttirType) SetCsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldCsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldCsmMask)
	}
}

const (
	RegisterFdcan_ttirFieldSogShift = 3
	RegisterFdcan_ttirFieldSogMask  = 0x8
)

// GetSog Start of Gap
func (r *registerFdcan_ttirType) GetSog() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldSogMask) != 0
}

// SetSog Start of Gap
func (r *registerFdcan_ttirType) SetSog(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldSogMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldSogMask)
	}
}

const (
	RegisterFdcan_ttirFieldRtmiShift = 4
	RegisterFdcan_ttirFieldRtmiMask  = 0x10
)

// GetRtmi Register Time Mark Interrupt.
func (r *registerFdcan_ttirType) GetRtmi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldRtmiMask) != 0
}

// SetRtmi Register Time Mark Interrupt.
func (r *registerFdcan_ttirType) SetRtmi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldRtmiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldRtmiMask)
	}
}

const (
	RegisterFdcan_ttirFieldTtmiShift = 5
	RegisterFdcan_ttirFieldTtmiMask  = 0x20
)

// GetTtmi Trigger Time Mark Event Internal
func (r *registerFdcan_ttirType) GetTtmi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldTtmiMask) != 0
}

// SetTtmi Trigger Time Mark Event Internal
func (r *registerFdcan_ttirType) SetTtmi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldTtmiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldTtmiMask)
	}
}

const (
	RegisterFdcan_ttirFieldSweShift = 6
	RegisterFdcan_ttirFieldSweMask  = 0x40
)

// GetSwe Stop Watch Event
func (r *registerFdcan_ttirType) GetSwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldSweMask) != 0
}

// SetSwe Stop Watch Event
func (r *registerFdcan_ttirType) SetSwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldSweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldSweMask)
	}
}

const (
	RegisterFdcan_ttirFieldGtwShift = 7
	RegisterFdcan_ttirFieldGtwMask  = 0x80
)

// GetGtw Global Time Wrap
func (r *registerFdcan_ttirType) GetGtw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldGtwMask) != 0
}

// SetGtw Global Time Wrap
func (r *registerFdcan_ttirType) SetGtw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldGtwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldGtwMask)
	}
}

const (
	RegisterFdcan_ttirFieldGtdShift = 8
	RegisterFdcan_ttirFieldGtdMask  = 0x100
)

// GetGtd Global Time Discontinuity
func (r *registerFdcan_ttirType) GetGtd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldGtdMask) != 0
}

// SetGtd Global Time Discontinuity
func (r *registerFdcan_ttirType) SetGtd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldGtdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldGtdMask)
	}
}

const (
	RegisterFdcan_ttirFieldGteShift = 9
	RegisterFdcan_ttirFieldGteMask  = 0x200
)

// GetGte Global Time Error
func (r *registerFdcan_ttirType) GetGte() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldGteMask) != 0
}

// SetGte Global Time Error
func (r *registerFdcan_ttirType) SetGte(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldGteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldGteMask)
	}
}

const (
	RegisterFdcan_ttirFieldTxuShift = 10
	RegisterFdcan_ttirFieldTxuMask  = 0x400
)

// GetTxu Tx Count Underflow
func (r *registerFdcan_ttirType) GetTxu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldTxuMask) != 0
}

// SetTxu Tx Count Underflow
func (r *registerFdcan_ttirType) SetTxu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldTxuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldTxuMask)
	}
}

const (
	RegisterFdcan_ttirFieldTxoShift = 11
	RegisterFdcan_ttirFieldTxoMask  = 0x800
)

// GetTxo Tx Count Overflow
func (r *registerFdcan_ttirType) GetTxo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldTxoMask) != 0
}

// SetTxo Tx Count Overflow
func (r *registerFdcan_ttirType) SetTxo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldTxoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldTxoMask)
	}
}

const (
	RegisterFdcan_ttirFieldSe1Shift = 12
	RegisterFdcan_ttirFieldSe1Mask  = 0x1000
)

// GetSe1 Scheduling Error 1
func (r *registerFdcan_ttirType) GetSe1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldSe1Mask) != 0
}

// SetSe1 Scheduling Error 1
func (r *registerFdcan_ttirType) SetSe1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldSe1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldSe1Mask)
	}
}

const (
	RegisterFdcan_ttirFieldSe2Shift = 13
	RegisterFdcan_ttirFieldSe2Mask  = 0x2000
)

// GetSe2 Scheduling Error 2
func (r *registerFdcan_ttirType) GetSe2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldSe2Mask) != 0
}

// SetSe2 Scheduling Error 2
func (r *registerFdcan_ttirType) SetSe2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldSe2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldSe2Mask)
	}
}

const (
	RegisterFdcan_ttirFieldElcShift = 14
	RegisterFdcan_ttirFieldElcMask  = 0x4000
)

// GetElc Error Level Changed.
func (r *registerFdcan_ttirType) GetElc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldElcMask) != 0
}

// SetElc Error Level Changed.
func (r *registerFdcan_ttirType) SetElc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldElcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldElcMask)
	}
}

const (
	RegisterFdcan_ttirFieldIwtgShift = 15
	RegisterFdcan_ttirFieldIwtgMask  = 0x8000
)

// GetIwtg Initialization Watch Trigger
func (r *registerFdcan_ttirType) GetIwtg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldIwtgMask) != 0
}

// SetIwtg Initialization Watch Trigger
func (r *registerFdcan_ttirType) SetIwtg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldIwtgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldIwtgMask)
	}
}

const (
	RegisterFdcan_ttirFieldWtShift = 16
	RegisterFdcan_ttirFieldWtMask  = 0x10000
)

// GetWt Watch Trigger
func (r *registerFdcan_ttirType) GetWt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldWtMask) != 0
}

// SetWt Watch Trigger
func (r *registerFdcan_ttirType) SetWt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldWtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldWtMask)
	}
}

const (
	RegisterFdcan_ttirFieldAwShift = 17
	RegisterFdcan_ttirFieldAwMask  = 0x20000
)

// GetAw Application Watchdog
func (r *registerFdcan_ttirType) GetAw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldAwMask) != 0
}

// SetAw Application Watchdog
func (r *registerFdcan_ttirType) SetAw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldAwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldAwMask)
	}
}

const (
	RegisterFdcan_ttirFieldCerShift = 18
	RegisterFdcan_ttirFieldCerMask  = 0x40000
)

// GetCer Configuration Error
func (r *registerFdcan_ttirType) GetCer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttirFieldCerMask) != 0
}

// SetCer Configuration Error
func (r *registerFdcan_ttirType) SetCer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttirFieldCerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttirFieldCerMask)
	}
}

// registerFdcan_ttieType FDCAN TT Interrupt Enable Register
type registerFdcan_ttieType uint32

const (
	RegisterFdcan_ttieFieldSbceShift = 0
	RegisterFdcan_ttieFieldSbceMask  = 0x1
)

// GetSbce Start of Basic Cycle Interrupt Enable
func (r *registerFdcan_ttieType) GetSbce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldSbceMask) != 0
}

// SetSbce Start of Basic Cycle Interrupt Enable
func (r *registerFdcan_ttieType) SetSbce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldSbceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldSbceMask)
	}
}

const (
	RegisterFdcan_ttieFieldSmceShift = 1
	RegisterFdcan_ttieFieldSmceMask  = 0x2
)

// GetSmce Start of Matrix Cycle Interrupt Enable
func (r *registerFdcan_ttieType) GetSmce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldSmceMask) != 0
}

// SetSmce Start of Matrix Cycle Interrupt Enable
func (r *registerFdcan_ttieType) SetSmce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldSmceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldSmceMask)
	}
}

const (
	RegisterFdcan_ttieFieldCsmeShift = 2
	RegisterFdcan_ttieFieldCsmeMask  = 0x4
)

// GetCsme Change of Synchronization Mode Interrupt Enable
func (r *registerFdcan_ttieType) GetCsme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldCsmeMask) != 0
}

// SetCsme Change of Synchronization Mode Interrupt Enable
func (r *registerFdcan_ttieType) SetCsme(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldCsmeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldCsmeMask)
	}
}

const (
	RegisterFdcan_ttieFieldSogeShift = 3
	RegisterFdcan_ttieFieldSogeMask  = 0x8
)

// GetSoge Start of Gap Interrupt Enable
func (r *registerFdcan_ttieType) GetSoge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldSogeMask) != 0
}

// SetSoge Start of Gap Interrupt Enable
func (r *registerFdcan_ttieType) SetSoge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldSogeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldSogeMask)
	}
}

const (
	RegisterFdcan_ttieFieldRtmieShift = 4
	RegisterFdcan_ttieFieldRtmieMask  = 0x10
)

// GetRtmie Register Time Mark Interrupt Enable
func (r *registerFdcan_ttieType) GetRtmie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldRtmieMask) != 0
}

// SetRtmie Register Time Mark Interrupt Enable
func (r *registerFdcan_ttieType) SetRtmie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldRtmieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldRtmieMask)
	}
}

const (
	RegisterFdcan_ttieFieldTtmieShift = 5
	RegisterFdcan_ttieFieldTtmieMask  = 0x20
)

// GetTtmie Trigger Time Mark Event Internal Interrupt Enable
func (r *registerFdcan_ttieType) GetTtmie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldTtmieMask) != 0
}

// SetTtmie Trigger Time Mark Event Internal Interrupt Enable
func (r *registerFdcan_ttieType) SetTtmie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldTtmieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldTtmieMask)
	}
}

const (
	RegisterFdcan_ttieFieldSweeShift = 6
	RegisterFdcan_ttieFieldSweeMask  = 0x40
)

// GetSwee Stop Watch Event Interrupt Enable
func (r *registerFdcan_ttieType) GetSwee() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldSweeMask) != 0
}

// SetSwee Stop Watch Event Interrupt Enable
func (r *registerFdcan_ttieType) SetSwee(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldSweeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldSweeMask)
	}
}

const (
	RegisterFdcan_ttieFieldGtweShift = 7
	RegisterFdcan_ttieFieldGtweMask  = 0x80
)

// GetGtwe Global Time Wrap Interrupt Enable
func (r *registerFdcan_ttieType) GetGtwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldGtweMask) != 0
}

// SetGtwe Global Time Wrap Interrupt Enable
func (r *registerFdcan_ttieType) SetGtwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldGtweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldGtweMask)
	}
}

const (
	RegisterFdcan_ttieFieldGtdeShift = 8
	RegisterFdcan_ttieFieldGtdeMask  = 0x100
)

// GetGtde Global Time Discontinuity Interrupt Enable
func (r *registerFdcan_ttieType) GetGtde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldGtdeMask) != 0
}

// SetGtde Global Time Discontinuity Interrupt Enable
func (r *registerFdcan_ttieType) SetGtde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldGtdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldGtdeMask)
	}
}

const (
	RegisterFdcan_ttieFieldGteeShift = 9
	RegisterFdcan_ttieFieldGteeMask  = 0x200
)

// GetGtee Global Time Error Interrupt Enable
func (r *registerFdcan_ttieType) GetGtee() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldGteeMask) != 0
}

// SetGtee Global Time Error Interrupt Enable
func (r *registerFdcan_ttieType) SetGtee(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldGteeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldGteeMask)
	}
}

const (
	RegisterFdcan_ttieFieldTxueShift = 10
	RegisterFdcan_ttieFieldTxueMask  = 0x400
)

// GetTxue Tx Count Underflow Interrupt Enable
func (r *registerFdcan_ttieType) GetTxue() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldTxueMask) != 0
}

// SetTxue Tx Count Underflow Interrupt Enable
func (r *registerFdcan_ttieType) SetTxue(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldTxueMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldTxueMask)
	}
}

const (
	RegisterFdcan_ttieFieldTxoeShift = 11
	RegisterFdcan_ttieFieldTxoeMask  = 0x800
)

// GetTxoe Tx Count Overflow Interrupt Enable
func (r *registerFdcan_ttieType) GetTxoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldTxoeMask) != 0
}

// SetTxoe Tx Count Overflow Interrupt Enable
func (r *registerFdcan_ttieType) SetTxoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldTxoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldTxoeMask)
	}
}

const (
	RegisterFdcan_ttieFieldSe1eShift = 12
	RegisterFdcan_ttieFieldSe1eMask  = 0x1000
)

// GetSe1e Scheduling Error 1 Interrupt Enable
func (r *registerFdcan_ttieType) GetSe1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldSe1eMask) != 0
}

// SetSe1e Scheduling Error 1 Interrupt Enable
func (r *registerFdcan_ttieType) SetSe1e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldSe1eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldSe1eMask)
	}
}

const (
	RegisterFdcan_ttieFieldSe2eShift = 13
	RegisterFdcan_ttieFieldSe2eMask  = 0x2000
)

// GetSe2e Scheduling Error 2 Interrupt Enable
func (r *registerFdcan_ttieType) GetSe2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldSe2eMask) != 0
}

// SetSe2e Scheduling Error 2 Interrupt Enable
func (r *registerFdcan_ttieType) SetSe2e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldSe2eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldSe2eMask)
	}
}

const (
	RegisterFdcan_ttieFieldElceShift = 14
	RegisterFdcan_ttieFieldElceMask  = 0x4000
)

// GetElce Change Error Level Interrupt Enable
func (r *registerFdcan_ttieType) GetElce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldElceMask) != 0
}

// SetElce Change Error Level Interrupt Enable
func (r *registerFdcan_ttieType) SetElce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldElceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldElceMask)
	}
}

const (
	RegisterFdcan_ttieFieldIwtgeShift = 15
	RegisterFdcan_ttieFieldIwtgeMask  = 0x8000
)

// GetIwtge Initialization Watch Trigger Interrupt Enable
func (r *registerFdcan_ttieType) GetIwtge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldIwtgeMask) != 0
}

// SetIwtge Initialization Watch Trigger Interrupt Enable
func (r *registerFdcan_ttieType) SetIwtge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldIwtgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldIwtgeMask)
	}
}

const (
	RegisterFdcan_ttieFieldWteShift = 16
	RegisterFdcan_ttieFieldWteMask  = 0x10000
)

// GetWte Watch Trigger Interrupt Enable
func (r *registerFdcan_ttieType) GetWte() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldWteMask) != 0
}

// SetWte Watch Trigger Interrupt Enable
func (r *registerFdcan_ttieType) SetWte(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldWteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldWteMask)
	}
}

const (
	RegisterFdcan_ttieFieldAweShift = 17
	RegisterFdcan_ttieFieldAweMask  = 0x20000
)

// GetAwe Application Watchdog Interrupt Enable
func (r *registerFdcan_ttieType) GetAwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldAweMask) != 0
}

// SetAwe Application Watchdog Interrupt Enable
func (r *registerFdcan_ttieType) SetAwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldAweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldAweMask)
	}
}

const (
	RegisterFdcan_ttieFieldCereShift = 18
	RegisterFdcan_ttieFieldCereMask  = 0x40000
)

// GetCere Configuration Error Interrupt Enable
func (r *registerFdcan_ttieType) GetCere() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttieFieldCereMask) != 0
}

// SetCere Configuration Error Interrupt Enable
func (r *registerFdcan_ttieType) SetCere(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttieFieldCereMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttieFieldCereMask)
	}
}

// registerFdcan_ttilsType FDCAN TT Interrupt Line Select Register
type registerFdcan_ttilsType uint32

const (
	RegisterFdcan_ttilsFieldSbclShift = 0
	RegisterFdcan_ttilsFieldSbclMask  = 0x1
)

// GetSbcl Start of Basic Cycle Interrupt Line
func (r *registerFdcan_ttilsType) GetSbcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldSbclMask) != 0
}

// SetSbcl Start of Basic Cycle Interrupt Line
func (r *registerFdcan_ttilsType) SetSbcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldSbclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldSbclMask)
	}
}

const (
	RegisterFdcan_ttilsFieldSmclShift = 1
	RegisterFdcan_ttilsFieldSmclMask  = 0x2
)

// GetSmcl Start of Matrix Cycle Interrupt Line
func (r *registerFdcan_ttilsType) GetSmcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldSmclMask) != 0
}

// SetSmcl Start of Matrix Cycle Interrupt Line
func (r *registerFdcan_ttilsType) SetSmcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldSmclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldSmclMask)
	}
}

const (
	RegisterFdcan_ttilsFieldCsmlShift = 2
	RegisterFdcan_ttilsFieldCsmlMask  = 0x4
)

// GetCsml Change of Synchronization Mode Interrupt Line
func (r *registerFdcan_ttilsType) GetCsml() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldCsmlMask) != 0
}

// SetCsml Change of Synchronization Mode Interrupt Line
func (r *registerFdcan_ttilsType) SetCsml(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldCsmlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldCsmlMask)
	}
}

const (
	RegisterFdcan_ttilsFieldSoglShift = 3
	RegisterFdcan_ttilsFieldSoglMask  = 0x8
)

// GetSogl Start of Gap Interrupt Line
func (r *registerFdcan_ttilsType) GetSogl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldSoglMask) != 0
}

// SetSogl Start of Gap Interrupt Line
func (r *registerFdcan_ttilsType) SetSogl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldSoglMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldSoglMask)
	}
}

const (
	RegisterFdcan_ttilsFieldRtmilShift = 4
	RegisterFdcan_ttilsFieldRtmilMask  = 0x10
)

// GetRtmil Register Time Mark Interrupt Line
func (r *registerFdcan_ttilsType) GetRtmil() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldRtmilMask) != 0
}

// SetRtmil Register Time Mark Interrupt Line
func (r *registerFdcan_ttilsType) SetRtmil(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldRtmilMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldRtmilMask)
	}
}

const (
	RegisterFdcan_ttilsFieldTtmilShift = 5
	RegisterFdcan_ttilsFieldTtmilMask  = 0x20
)

// GetTtmil Trigger Time Mark Event Internal Interrupt Line
func (r *registerFdcan_ttilsType) GetTtmil() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldTtmilMask) != 0
}

// SetTtmil Trigger Time Mark Event Internal Interrupt Line
func (r *registerFdcan_ttilsType) SetTtmil(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldTtmilMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldTtmilMask)
	}
}

const (
	RegisterFdcan_ttilsFieldSwelShift = 6
	RegisterFdcan_ttilsFieldSwelMask  = 0x40
)

// GetSwel Stop Watch Event Interrupt Line
func (r *registerFdcan_ttilsType) GetSwel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldSwelMask) != 0
}

// SetSwel Stop Watch Event Interrupt Line
func (r *registerFdcan_ttilsType) SetSwel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldSwelMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldSwelMask)
	}
}

const (
	RegisterFdcan_ttilsFieldGtwlShift = 7
	RegisterFdcan_ttilsFieldGtwlMask  = 0x80
)

// GetGtwl Global Time Wrap Interrupt Line
func (r *registerFdcan_ttilsType) GetGtwl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldGtwlMask) != 0
}

// SetGtwl Global Time Wrap Interrupt Line
func (r *registerFdcan_ttilsType) SetGtwl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldGtwlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldGtwlMask)
	}
}

const (
	RegisterFdcan_ttilsFieldGtdlShift = 8
	RegisterFdcan_ttilsFieldGtdlMask  = 0x100
)

// GetGtdl Global Time Discontinuity Interrupt Line
func (r *registerFdcan_ttilsType) GetGtdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldGtdlMask) != 0
}

// SetGtdl Global Time Discontinuity Interrupt Line
func (r *registerFdcan_ttilsType) SetGtdl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldGtdlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldGtdlMask)
	}
}

const (
	RegisterFdcan_ttilsFieldGtelShift = 9
	RegisterFdcan_ttilsFieldGtelMask  = 0x200
)

// GetGtel Global Time Error Interrupt Line
func (r *registerFdcan_ttilsType) GetGtel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldGtelMask) != 0
}

// SetGtel Global Time Error Interrupt Line
func (r *registerFdcan_ttilsType) SetGtel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldGtelMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldGtelMask)
	}
}

const (
	RegisterFdcan_ttilsFieldTxulShift = 10
	RegisterFdcan_ttilsFieldTxulMask  = 0x400
)

// GetTxul Tx Count Underflow Interrupt Line
func (r *registerFdcan_ttilsType) GetTxul() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldTxulMask) != 0
}

// SetTxul Tx Count Underflow Interrupt Line
func (r *registerFdcan_ttilsType) SetTxul(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldTxulMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldTxulMask)
	}
}

const (
	RegisterFdcan_ttilsFieldTxolShift = 11
	RegisterFdcan_ttilsFieldTxolMask  = 0x800
)

// GetTxol Tx Count Overflow Interrupt Line
func (r *registerFdcan_ttilsType) GetTxol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldTxolMask) != 0
}

// SetTxol Tx Count Overflow Interrupt Line
func (r *registerFdcan_ttilsType) SetTxol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldTxolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldTxolMask)
	}
}

const (
	RegisterFdcan_ttilsFieldSe1lShift = 12
	RegisterFdcan_ttilsFieldSe1lMask  = 0x1000
)

// GetSe1l Scheduling Error 1 Interrupt Line
func (r *registerFdcan_ttilsType) GetSe1l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldSe1lMask) != 0
}

// SetSe1l Scheduling Error 1 Interrupt Line
func (r *registerFdcan_ttilsType) SetSe1l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldSe1lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldSe1lMask)
	}
}

const (
	RegisterFdcan_ttilsFieldSe2lShift = 13
	RegisterFdcan_ttilsFieldSe2lMask  = 0x2000
)

// GetSe2l Scheduling Error 2 Interrupt Line
func (r *registerFdcan_ttilsType) GetSe2l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldSe2lMask) != 0
}

// SetSe2l Scheduling Error 2 Interrupt Line
func (r *registerFdcan_ttilsType) SetSe2l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldSe2lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldSe2lMask)
	}
}

const (
	RegisterFdcan_ttilsFieldElclShift = 14
	RegisterFdcan_ttilsFieldElclMask  = 0x4000
)

// GetElcl Change Error Level Interrupt Line
func (r *registerFdcan_ttilsType) GetElcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldElclMask) != 0
}

// SetElcl Change Error Level Interrupt Line
func (r *registerFdcan_ttilsType) SetElcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldElclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldElclMask)
	}
}

const (
	RegisterFdcan_ttilsFieldIwtglShift = 15
	RegisterFdcan_ttilsFieldIwtglMask  = 0x8000
)

// GetIwtgl Initialization Watch Trigger Interrupt Line
func (r *registerFdcan_ttilsType) GetIwtgl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldIwtglMask) != 0
}

// SetIwtgl Initialization Watch Trigger Interrupt Line
func (r *registerFdcan_ttilsType) SetIwtgl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldIwtglMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldIwtglMask)
	}
}

const (
	RegisterFdcan_ttilsFieldWtlShift = 16
	RegisterFdcan_ttilsFieldWtlMask  = 0x10000
)

// GetWtl Watch Trigger Interrupt Line
func (r *registerFdcan_ttilsType) GetWtl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldWtlMask) != 0
}

// SetWtl Watch Trigger Interrupt Line
func (r *registerFdcan_ttilsType) SetWtl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldWtlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldWtlMask)
	}
}

const (
	RegisterFdcan_ttilsFieldAwlShift = 17
	RegisterFdcan_ttilsFieldAwlMask  = 0x20000
)

// GetAwl Application Watchdog Interrupt Line
func (r *registerFdcan_ttilsType) GetAwl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldAwlMask) != 0
}

// SetAwl Application Watchdog Interrupt Line
func (r *registerFdcan_ttilsType) SetAwl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldAwlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldAwlMask)
	}
}

const (
	RegisterFdcan_ttilsFieldCerlShift = 18
	RegisterFdcan_ttilsFieldCerlMask  = 0x40000
)

// GetCerl Configuration Error Interrupt Line
func (r *registerFdcan_ttilsType) GetCerl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttilsFieldCerlMask) != 0
}

// SetCerl Configuration Error Interrupt Line
func (r *registerFdcan_ttilsType) SetCerl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttilsFieldCerlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttilsFieldCerlMask)
	}
}

// registerFdcan_ttostType FDCAN TT Operation Status Register
type registerFdcan_ttostType uint32

const (
	RegisterFdcan_ttostFieldElShift = 0
	RegisterFdcan_ttostFieldElMask  = 0x3
)

// GetEl Error Level
func (r *registerFdcan_ttostType) GetEl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldElMask) >> RegisterFdcan_ttostFieldElShift)
}

// SetEl Error Level
func (r *registerFdcan_ttostType) SetEl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldElMask)|(uint32(value)<<RegisterFdcan_ttostFieldElShift))
}

const (
	RegisterFdcan_ttostFieldMsShift = 2
	RegisterFdcan_ttostFieldMsMask  = 0xc
)

// GetMs Master State.
func (r *registerFdcan_ttostType) GetMs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldMsMask) >> RegisterFdcan_ttostFieldMsShift)
}

// SetMs Master State.
func (r *registerFdcan_ttostType) SetMs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldMsMask)|(uint32(value)<<RegisterFdcan_ttostFieldMsShift))
}

const (
	RegisterFdcan_ttostFieldSysShift = 4
	RegisterFdcan_ttostFieldSysMask  = 0x30
)

// GetSys Synchronization State
func (r *registerFdcan_ttostType) GetSys() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldSysMask) >> RegisterFdcan_ttostFieldSysShift)
}

// SetSys Synchronization State
func (r *registerFdcan_ttostType) SetSys(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldSysMask)|(uint32(value)<<RegisterFdcan_ttostFieldSysShift))
}

const (
	RegisterFdcan_ttostFieldGtpShift = 6
	RegisterFdcan_ttostFieldGtpMask  = 0x40
)

// GetGtp Quality of Global Time Phase
func (r *registerFdcan_ttostType) GetGtp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldGtpMask) != 0
}

// SetGtp Quality of Global Time Phase
func (r *registerFdcan_ttostType) SetGtp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttostFieldGtpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldGtpMask)
	}
}

const (
	RegisterFdcan_ttostFieldQcsShift = 7
	RegisterFdcan_ttostFieldQcsMask  = 0x80
)

// GetQcs Quality of Clock Speed
func (r *registerFdcan_ttostType) GetQcs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldQcsMask) != 0
}

// SetQcs Quality of Clock Speed
func (r *registerFdcan_ttostType) SetQcs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttostFieldQcsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldQcsMask)
	}
}

const (
	RegisterFdcan_ttostFieldRtoShift = 8
	RegisterFdcan_ttostFieldRtoMask  = 0xff00
)

// GetRto Reference Trigger Offset
func (r *registerFdcan_ttostType) GetRto() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldRtoMask) >> RegisterFdcan_ttostFieldRtoShift)
}

// SetRto Reference Trigger Offset
func (r *registerFdcan_ttostType) SetRto(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldRtoMask)|(uint32(value)<<RegisterFdcan_ttostFieldRtoShift))
}

const (
	RegisterFdcan_ttostFieldWgtdShift = 22
	RegisterFdcan_ttostFieldWgtdMask  = 0x400000
)

// GetWgtd Wait for Global Time Discontinuity
func (r *registerFdcan_ttostType) GetWgtd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldWgtdMask) != 0
}

// SetWgtd Wait for Global Time Discontinuity
func (r *registerFdcan_ttostType) SetWgtd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttostFieldWgtdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldWgtdMask)
	}
}

const (
	RegisterFdcan_ttostFieldGfiShift = 23
	RegisterFdcan_ttostFieldGfiMask  = 0x800000
)

// GetGfi Gap Finished Indicator.
func (r *registerFdcan_ttostType) GetGfi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldGfiMask) != 0
}

// SetGfi Gap Finished Indicator.
func (r *registerFdcan_ttostType) SetGfi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttostFieldGfiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldGfiMask)
	}
}

const (
	RegisterFdcan_ttostFieldTmpShift = 24
	RegisterFdcan_ttostFieldTmpMask  = 0x7000000
)

// GetTmp Time Master Priority
func (r *registerFdcan_ttostType) GetTmp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldTmpMask) >> RegisterFdcan_ttostFieldTmpShift)
}

// SetTmp Time Master Priority
func (r *registerFdcan_ttostType) SetTmp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldTmpMask)|(uint32(value)<<RegisterFdcan_ttostFieldTmpShift))
}

const (
	RegisterFdcan_ttostFieldGsiShift = 27
	RegisterFdcan_ttostFieldGsiMask  = 0x8000000
)

// GetGsi Gap Started Indicator.
func (r *registerFdcan_ttostType) GetGsi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldGsiMask) != 0
}

// SetGsi Gap Started Indicator.
func (r *registerFdcan_ttostType) SetGsi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttostFieldGsiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldGsiMask)
	}
}

const (
	RegisterFdcan_ttostFieldWfeShift = 28
	RegisterFdcan_ttostFieldWfeMask  = 0x10000000
)

// GetWfe Wait for Event
func (r *registerFdcan_ttostType) GetWfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldWfeMask) != 0
}

// SetWfe Wait for Event
func (r *registerFdcan_ttostType) SetWfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttostFieldWfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldWfeMask)
	}
}

const (
	RegisterFdcan_ttostFieldAweShift = 29
	RegisterFdcan_ttostFieldAweMask  = 0x20000000
)

// GetAwe Application Watchdog Event
func (r *registerFdcan_ttostType) GetAwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldAweMask) != 0
}

// SetAwe Application Watchdog Event
func (r *registerFdcan_ttostType) SetAwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttostFieldAweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldAweMask)
	}
}

const (
	RegisterFdcan_ttostFieldWecsShift = 30
	RegisterFdcan_ttostFieldWecsMask  = 0x40000000
)

// GetWecs Wait for External Clock Synchronization
func (r *registerFdcan_ttostType) GetWecs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldWecsMask) != 0
}

// SetWecs Wait for External Clock Synchronization
func (r *registerFdcan_ttostType) SetWecs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttostFieldWecsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldWecsMask)
	}
}

const (
	RegisterFdcan_ttostFieldSplShift = 31
	RegisterFdcan_ttostFieldSplMask  = 0x80000000
)

// GetSpl Schedule Phase Lock
func (r *registerFdcan_ttostType) GetSpl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttostFieldSplMask) != 0
}

// SetSpl Schedule Phase Lock
func (r *registerFdcan_ttostType) SetSpl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcan_ttostFieldSplMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttostFieldSplMask)
	}
}

// registerFdcan_turnaType FDCAN TUR Numerator Actual Register
type registerFdcan_turnaType uint32

const (
	RegisterFdcan_turnaFieldNavShift = 0
	RegisterFdcan_turnaFieldNavMask  = 0x3ffff
)

// GetNav Numerator Actual Value
func (r *registerFdcan_turnaType) GetNav() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_turnaFieldNavMask) >> RegisterFdcan_turnaFieldNavShift)
}

// SetNav Numerator Actual Value
func (r *registerFdcan_turnaType) SetNav(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_turnaFieldNavMask)|(uint32(value)<<RegisterFdcan_turnaFieldNavShift))
}

// registerFdcan_ttlgtType FDCAN TT Local and Global Time Register
type registerFdcan_ttlgtType uint32

const (
	RegisterFdcan_ttlgtFieldLtShift = 0
	RegisterFdcan_ttlgtFieldLtMask  = 0xffff
)

// GetLt Local Time
func (r *registerFdcan_ttlgtType) GetLt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttlgtFieldLtMask) >> RegisterFdcan_ttlgtFieldLtShift)
}

// SetLt Local Time
func (r *registerFdcan_ttlgtType) SetLt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttlgtFieldLtMask)|(uint32(value)<<RegisterFdcan_ttlgtFieldLtShift))
}

const (
	RegisterFdcan_ttlgtFieldGtShift = 16
	RegisterFdcan_ttlgtFieldGtMask  = 0xffff0000
)

// GetGt Global Time
func (r *registerFdcan_ttlgtType) GetGt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttlgtFieldGtMask) >> RegisterFdcan_ttlgtFieldGtShift)
}

// SetGt Global Time
func (r *registerFdcan_ttlgtType) SetGt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttlgtFieldGtMask)|(uint32(value)<<RegisterFdcan_ttlgtFieldGtShift))
}

// registerFdcan_ttctcType FDCAN TT Cycle Time and Count Register
type registerFdcan_ttctcType uint32

const (
	RegisterFdcan_ttctcFieldCtShift = 0
	RegisterFdcan_ttctcFieldCtMask  = 0xffff
)

// GetCt Cycle Time
func (r *registerFdcan_ttctcType) GetCt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttctcFieldCtMask) >> RegisterFdcan_ttctcFieldCtShift)
}

// SetCt Cycle Time
func (r *registerFdcan_ttctcType) SetCt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttctcFieldCtMask)|(uint32(value)<<RegisterFdcan_ttctcFieldCtShift))
}

const (
	RegisterFdcan_ttctcFieldCcShift = 16
	RegisterFdcan_ttctcFieldCcMask  = 0x3f0000
)

// GetCc Cycle Count
func (r *registerFdcan_ttctcType) GetCc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttctcFieldCcMask) >> RegisterFdcan_ttctcFieldCcShift)
}

// SetCc Cycle Count
func (r *registerFdcan_ttctcType) SetCc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttctcFieldCcMask)|(uint32(value)<<RegisterFdcan_ttctcFieldCcShift))
}

// registerFdcan_ttcptType FDCAN TT Capture Time Register
type registerFdcan_ttcptType uint32

const (
	RegisterFdcan_ttcptFieldCtShift = 0
	RegisterFdcan_ttcptFieldCtMask  = 0x3f
)

// GetCt Cycle Count Value
func (r *registerFdcan_ttcptType) GetCt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttcptFieldCtMask) >> RegisterFdcan_ttcptFieldCtShift)
}

// SetCt Cycle Count Value
func (r *registerFdcan_ttcptType) SetCt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttcptFieldCtMask)|(uint32(value)<<RegisterFdcan_ttcptFieldCtShift))
}

const (
	RegisterFdcan_ttcptFieldSwvShift = 16
	RegisterFdcan_ttcptFieldSwvMask  = 0xffff0000
)

// GetSwv Stop Watch Value
func (r *registerFdcan_ttcptType) GetSwv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttcptFieldSwvMask) >> RegisterFdcan_ttcptFieldSwvShift)
}

// SetSwv Stop Watch Value
func (r *registerFdcan_ttcptType) SetSwv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttcptFieldSwvMask)|(uint32(value)<<RegisterFdcan_ttcptFieldSwvShift))
}

// registerFdcan_ttcsmType FDCAN TT Cycle Sync Mark Register
type registerFdcan_ttcsmType uint32

const (
	RegisterFdcan_ttcsmFieldCsmShift = 0
	RegisterFdcan_ttcsmFieldCsmMask  = 0xffff
)

// GetCsm Cycle Sync Mark
func (r *registerFdcan_ttcsmType) GetCsm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_ttcsmFieldCsmMask) >> RegisterFdcan_ttcsmFieldCsmShift)
}

// SetCsm Cycle Sync Mark
func (r *registerFdcan_ttcsmType) SetCsm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_ttcsmFieldCsmMask)|(uint32(value)<<RegisterFdcan_ttcsmFieldCsmShift))
}

// registerFdcan_tttsType FDCAN TT Trigger Select Register
type registerFdcan_tttsType uint32

const (
	RegisterFdcan_tttsFieldSwtdelShift = 0
	RegisterFdcan_tttsFieldSwtdelMask  = 0x3
)

// GetSwtdel Stop watch trigger input selection
func (r *registerFdcan_tttsType) GetSwtdel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tttsFieldSwtdelMask) >> RegisterFdcan_tttsFieldSwtdelShift)
}

// SetSwtdel Stop watch trigger input selection
func (r *registerFdcan_tttsType) SetSwtdel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tttsFieldSwtdelMask)|(uint32(value)<<RegisterFdcan_tttsFieldSwtdelShift))
}

const (
	RegisterFdcan_tttsFieldEvtselShift = 4
	RegisterFdcan_tttsFieldEvtselMask  = 0x30
)

// GetEvtsel Event trigger input selection
func (r *registerFdcan_tttsType) GetEvtsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcan_tttsFieldEvtselMask) >> RegisterFdcan_tttsFieldEvtselShift)
}

// SetEvtsel Event trigger input selection
func (r *registerFdcan_tttsType) SetEvtsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcan_tttsFieldEvtselMask)|(uint32(value)<<RegisterFdcan_tttsFieldEvtselShift))
}
