//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

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
	Fdcancrel   registerFdcancrelType
	Fdcanendn   registerFdcanendnType
	_           [4]byte
	Fdcandbtp   registerFdcandbtpType
	Fdcantest   registerFdcantestType
	Fdcanrwd    registerFdcanrwdType
	Fdcancccr   registerFdcancccrType
	Fdcannbtp   registerFdcannbtpType
	Fdcantscc   registerFdcantsccType
	Fdcantscv   registerFdcantscvType
	Fdcantocc   registerFdcantoccType
	Fdcantocv   registerFdcantocvType
	_           [16]byte
	Fdcanecr    registerFdcanecrType
	Fdcanpsr    registerFdcanpsrType
	Fdcantdcr   registerFdcantdcrType
	_           [4]byte
	Fdcanir     registerFdcanirType
	Fdcanie     registerFdcanieType
	Fdcanils    registerFdcanilsType
	Fdcanile    registerFdcanileType
	_           [32]byte
	Fdcangfc    registerFdcangfcType
	Fdcansidfc  registerFdcansidfcType
	Fdcanxidfc  registerFdcanxidfcType
	_           [4]byte
	Fdcanxidam  registerFdcanxidamType
	Fdcanhpms   registerFdcanhpmsType
	Fdcanndat1  registerFdcanndat1Type
	Fdcanndat2  registerFdcanndat2Type
	Fdcanrxf0c  registerFdcanrxf0cType
	Fdcanrxf0s  registerFdcanrxf0sType
	Fdcanrxf0a  registerFdcanrxf0aType
	Fdcanrxbc   registerFdcanrxbcType
	Fdcanrxf1c  registerFdcanrxf1cType
	Fdcanrxf1s  registerFdcanrxf1sType
	Fdcanrxf1a  registerFdcanrxf1aType
	Fdcanrxesc  registerFdcanrxescType
	Fdcantxbc   registerFdcantxbcType
	Fdcantxfqs  registerFdcantxfqsType
	Fdcantxesc  registerFdcantxescType
	Fdcantxbrp  registerFdcantxbrpType
	Fdcantxbar  registerFdcantxbarType
	Fdcantxbcr  registerFdcantxbcrType
	Fdcantxbto  registerFdcantxbtoType
	Fdcantxbcf  registerFdcantxbcfType
	Fdcantxbtie registerFdcantxbtieType
	Fdcantxbcie registerFdcantxbcieType
	_           [8]byte
	Fdcantxefc  registerFdcantxefcType
	Fdcantxefs  registerFdcantxefsType
	Fdcantxefa  registerFdcantxefaType
	_           [4]byte
	Fdcantttmc  registerFdcantttmcType
	Fdcanttrmc  registerFdcanttrmcType
	Fdcanttocf  registerFdcanttocfType
	Fdcanttmlm  registerFdcanttmlmType
	Fdcanturcf  registerFdcanturcfType
	Fdcanttocn  registerFdcanttocnType
	Canttgtp    registerCanttgtpType
	Fdcantttmk  registerFdcantttmkType
	Fdcanttir   registerFdcanttirType
	Fdcanttie   registerFdcanttieType
	Fdcanttils  registerFdcanttilsType
	Fdcanttost  registerFdcanttostType
	Fdcanturna  registerFdcanturnaType
	Fdcanttlgt  registerFdcanttlgtType
	Fdcanttctc  registerFdcanttctcType
	Fdcanttcpt  registerFdcanttcptType
	Fdcanttcsm  registerFdcanttcsmType
	_           [444]byte
	Fdcanttts   registerFdcantttsType
}

// registerFdcancrelType FDCAN Core Release Register
type registerFdcancrelType uint32

const (
	RegisterFdcancrelFieldDayShift = 0
	RegisterFdcancrelFieldDayMask  = 0xff
)

// GetDay Timestamp Day
func (r *registerFdcancrelType) GetDay() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcancrelFieldDayMask) >> RegisterFdcancrelFieldDayShift)
}

// SetDay Timestamp Day
func (r *registerFdcancrelType) SetDay(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcancrelFieldDayMask)|(uint32(value)<<RegisterFdcancrelFieldDayShift))
}

const (
	RegisterFdcancrelFieldMonShift = 8
	RegisterFdcancrelFieldMonMask  = 0xff00
)

// GetMon Timestamp Month
func (r *registerFdcancrelType) GetMon() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcancrelFieldMonMask) >> RegisterFdcancrelFieldMonShift)
}

// SetMon Timestamp Month
func (r *registerFdcancrelType) SetMon(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcancrelFieldMonMask)|(uint32(value)<<RegisterFdcancrelFieldMonShift))
}

const (
	RegisterFdcancrelFieldYearShift = 16
	RegisterFdcancrelFieldYearMask  = 0xf0000
)

// GetYear Timestamp Year
func (r *registerFdcancrelType) GetYear() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcancrelFieldYearMask) >> RegisterFdcancrelFieldYearShift)
}

// SetYear Timestamp Year
func (r *registerFdcancrelType) SetYear(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcancrelFieldYearMask)|(uint32(value)<<RegisterFdcancrelFieldYearShift))
}

const (
	RegisterFdcancrelFieldSubstepShift = 20
	RegisterFdcancrelFieldSubstepMask  = 0xf00000
)

// GetSubstep Sub-step of Core release
func (r *registerFdcancrelType) GetSubstep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcancrelFieldSubstepMask) >> RegisterFdcancrelFieldSubstepShift)
}

// SetSubstep Sub-step of Core release
func (r *registerFdcancrelType) SetSubstep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcancrelFieldSubstepMask)|(uint32(value)<<RegisterFdcancrelFieldSubstepShift))
}

const (
	RegisterFdcancrelFieldStepShift = 24
	RegisterFdcancrelFieldStepMask  = 0xf000000
)

// GetStep Step of Core release
func (r *registerFdcancrelType) GetStep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcancrelFieldStepMask) >> RegisterFdcancrelFieldStepShift)
}

// SetStep Step of Core release
func (r *registerFdcancrelType) SetStep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcancrelFieldStepMask)|(uint32(value)<<RegisterFdcancrelFieldStepShift))
}

const (
	RegisterFdcancrelFieldRelShift = 28
	RegisterFdcancrelFieldRelMask  = 0xf0000000
)

// GetRel Core release
func (r *registerFdcancrelType) GetRel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcancrelFieldRelMask) >> RegisterFdcancrelFieldRelShift)
}

// SetRel Core release
func (r *registerFdcancrelType) SetRel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcancrelFieldRelMask)|(uint32(value)<<RegisterFdcancrelFieldRelShift))
}

// registerFdcanendnType FDCAN Core Release Register
type registerFdcanendnType uint32

const (
	RegisterFdcanendnFieldEtvShift = 0
	RegisterFdcanendnFieldEtvMask  = 0xffffffff
)

// GetEtv Endiannes Test Value
func (r *registerFdcanendnType) GetEtv() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanendnFieldEtvMask) >> RegisterFdcanendnFieldEtvShift)
}

// SetEtv Endiannes Test Value
func (r *registerFdcanendnType) SetEtv(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanendnFieldEtvMask)|(uint32(value)<<RegisterFdcanendnFieldEtvShift))
}

// registerFdcandbtpType FDCAN Data Bit Timing and Prescaler Register
type registerFdcandbtpType uint32

const (
	RegisterFdcandbtpFieldDsjwShift = 0
	RegisterFdcandbtpFieldDsjwMask  = 0xf
)

// GetDsjw Synchronization Jump Width
func (r *registerFdcandbtpType) GetDsjw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcandbtpFieldDsjwMask) >> RegisterFdcandbtpFieldDsjwShift)
}

// SetDsjw Synchronization Jump Width
func (r *registerFdcandbtpType) SetDsjw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcandbtpFieldDsjwMask)|(uint32(value)<<RegisterFdcandbtpFieldDsjwShift))
}

const (
	RegisterFdcandbtpFieldDtseg2Shift = 4
	RegisterFdcandbtpFieldDtseg2Mask  = 0xf0
)

// GetDtseg2 Data time segment after sample point
func (r *registerFdcandbtpType) GetDtseg2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcandbtpFieldDtseg2Mask) >> RegisterFdcandbtpFieldDtseg2Shift)
}

// SetDtseg2 Data time segment after sample point
func (r *registerFdcandbtpType) SetDtseg2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcandbtpFieldDtseg2Mask)|(uint32(value)<<RegisterFdcandbtpFieldDtseg2Shift))
}

const (
	RegisterFdcandbtpFieldDtseg1Shift = 8
	RegisterFdcandbtpFieldDtseg1Mask  = 0x1f00
)

// GetDtseg1 Data time segment after sample point
func (r *registerFdcandbtpType) GetDtseg1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcandbtpFieldDtseg1Mask) >> RegisterFdcandbtpFieldDtseg1Shift)
}

// SetDtseg1 Data time segment after sample point
func (r *registerFdcandbtpType) SetDtseg1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcandbtpFieldDtseg1Mask)|(uint32(value)<<RegisterFdcandbtpFieldDtseg1Shift))
}

const (
	RegisterFdcandbtpFieldDbrpShift = 16
	RegisterFdcandbtpFieldDbrpMask  = 0x1f0000
)

// GetDbrp Data BIt Rate Prescaler
func (r *registerFdcandbtpType) GetDbrp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcandbtpFieldDbrpMask) >> RegisterFdcandbtpFieldDbrpShift)
}

// SetDbrp Data BIt Rate Prescaler
func (r *registerFdcandbtpType) SetDbrp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcandbtpFieldDbrpMask)|(uint32(value)<<RegisterFdcandbtpFieldDbrpShift))
}

const (
	RegisterFdcandbtpFieldTdcShift = 23
	RegisterFdcandbtpFieldTdcMask  = 0x800000
)

// GetTdc Transceiver Delay Compensation
func (r *registerFdcandbtpType) GetTdc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcandbtpFieldTdcMask) != 0
}

// SetTdc Transceiver Delay Compensation
func (r *registerFdcandbtpType) SetTdc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcandbtpFieldTdcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcandbtpFieldTdcMask)
	}
}

// registerFdcantestType FDCAN Test Register
type registerFdcantestType uint32

const (
	RegisterFdcantestFieldLbckShift = 4
	RegisterFdcantestFieldLbckMask  = 0x10
)

// GetLbck Loop Back mode
func (r *registerFdcantestType) GetLbck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcantestFieldLbckMask) != 0
}

// SetLbck Loop Back mode
func (r *registerFdcantestType) SetLbck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcantestFieldLbckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcantestFieldLbckMask)
	}
}

const (
	RegisterFdcantestFieldTxShift = 5
	RegisterFdcantestFieldTxMask  = 0x60
)

// GetTx Loop Back mode
func (r *registerFdcantestType) GetTx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantestFieldTxMask) >> RegisterFdcantestFieldTxShift)
}

// SetTx Loop Back mode
func (r *registerFdcantestType) SetTx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantestFieldTxMask)|(uint32(value)<<RegisterFdcantestFieldTxShift))
}

const (
	RegisterFdcantestFieldRxShift = 7
	RegisterFdcantestFieldRxMask  = 0x80
)

// GetRx Control of Transmit Pin
func (r *registerFdcantestType) GetRx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcantestFieldRxMask) != 0
}

// SetRx Control of Transmit Pin
func (r *registerFdcantestType) SetRx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcantestFieldRxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcantestFieldRxMask)
	}
}

// registerFdcanrwdType FDCAN RAM Watchdog Register
type registerFdcanrwdType uint32

const (
	RegisterFdcanrwdFieldWdcShift = 0
	RegisterFdcanrwdFieldWdcMask  = 0xff
)

// GetWdc Watchdog configuration
func (r *registerFdcanrwdType) GetWdc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrwdFieldWdcMask) >> RegisterFdcanrwdFieldWdcShift)
}

// SetWdc Watchdog configuration
func (r *registerFdcanrwdType) SetWdc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrwdFieldWdcMask)|(uint32(value)<<RegisterFdcanrwdFieldWdcShift))
}

const (
	RegisterFdcanrwdFieldWdvShift = 8
	RegisterFdcanrwdFieldWdvMask  = 0xff00
)

// GetWdv Watchdog value
func (r *registerFdcanrwdType) GetWdv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrwdFieldWdvMask) >> RegisterFdcanrwdFieldWdvShift)
}

// SetWdv Watchdog value
func (r *registerFdcanrwdType) SetWdv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrwdFieldWdvMask)|(uint32(value)<<RegisterFdcanrwdFieldWdvShift))
}

// registerFdcancccrType FDCAN CC Control Register
type registerFdcancccrType uint32

const (
	RegisterFdcancccrFieldInitShift = 0
	RegisterFdcancccrFieldInitMask  = 0x1
)

// GetInit Initialization
func (r *registerFdcancccrType) GetInit() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldInitMask) != 0
}

// SetInit Initialization
func (r *registerFdcancccrType) SetInit(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldInitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldInitMask)
	}
}

const (
	RegisterFdcancccrFieldCceShift = 1
	RegisterFdcancccrFieldCceMask  = 0x2
)

// GetCce Configuration Change Enable
func (r *registerFdcancccrType) GetCce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldCceMask) != 0
}

// SetCce Configuration Change Enable
func (r *registerFdcancccrType) SetCce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldCceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldCceMask)
	}
}

const (
	RegisterFdcancccrFieldAsmShift = 2
	RegisterFdcancccrFieldAsmMask  = 0x4
)

// GetAsm ASM Restricted Operation Mode
func (r *registerFdcancccrType) GetAsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldAsmMask) != 0
}

// SetAsm ASM Restricted Operation Mode
func (r *registerFdcancccrType) SetAsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldAsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldAsmMask)
	}
}

const (
	RegisterFdcancccrFieldCsaShift = 3
	RegisterFdcancccrFieldCsaMask  = 0x8
)

// GetCsa Clock Stop Acknowledge
func (r *registerFdcancccrType) GetCsa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldCsaMask) != 0
}

// SetCsa Clock Stop Acknowledge
func (r *registerFdcancccrType) SetCsa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldCsaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldCsaMask)
	}
}

const (
	RegisterFdcancccrFieldCsrShift = 4
	RegisterFdcancccrFieldCsrMask  = 0x10
)

// GetCsr Clock Stop Request
func (r *registerFdcancccrType) GetCsr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldCsrMask) != 0
}

// SetCsr Clock Stop Request
func (r *registerFdcancccrType) SetCsr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldCsrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldCsrMask)
	}
}

const (
	RegisterFdcancccrFieldMonShift = 5
	RegisterFdcancccrFieldMonMask  = 0x20
)

// GetMon Bus Monitoring Mode
func (r *registerFdcancccrType) GetMon() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldMonMask) != 0
}

// SetMon Bus Monitoring Mode
func (r *registerFdcancccrType) SetMon(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldMonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldMonMask)
	}
}

const (
	RegisterFdcancccrFieldDarShift = 6
	RegisterFdcancccrFieldDarMask  = 0x40
)

// GetDar Disable Automatic Retransmission
func (r *registerFdcancccrType) GetDar() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldDarMask) != 0
}

// SetDar Disable Automatic Retransmission
func (r *registerFdcancccrType) SetDar(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldDarMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldDarMask)
	}
}

const (
	RegisterFdcancccrFieldTestShift = 7
	RegisterFdcancccrFieldTestMask  = 0x80
)

// GetTest Test Mode Enable
func (r *registerFdcancccrType) GetTest() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldTestMask) != 0
}

// SetTest Test Mode Enable
func (r *registerFdcancccrType) SetTest(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldTestMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldTestMask)
	}
}

const (
	RegisterFdcancccrFieldFdoeShift = 8
	RegisterFdcancccrFieldFdoeMask  = 0x100
)

// GetFdoe FD Operation Enable
func (r *registerFdcancccrType) GetFdoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldFdoeMask) != 0
}

// SetFdoe FD Operation Enable
func (r *registerFdcancccrType) SetFdoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldFdoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldFdoeMask)
	}
}

const (
	RegisterFdcancccrFieldBseShift = 9
	RegisterFdcancccrFieldBseMask  = 0x200
)

// GetBse FDCAN Bit Rate Switching
func (r *registerFdcancccrType) GetBse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldBseMask) != 0
}

// SetBse FDCAN Bit Rate Switching
func (r *registerFdcancccrType) SetBse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldBseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldBseMask)
	}
}

const (
	RegisterFdcancccrFieldPxhdShift = 12
	RegisterFdcancccrFieldPxhdMask  = 0x1000
)

// GetPxhd Protocol Exception Handling Disable
func (r *registerFdcancccrType) GetPxhd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldPxhdMask) != 0
}

// SetPxhd Protocol Exception Handling Disable
func (r *registerFdcancccrType) SetPxhd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldPxhdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldPxhdMask)
	}
}

const (
	RegisterFdcancccrFieldEfbiShift = 13
	RegisterFdcancccrFieldEfbiMask  = 0x2000
)

// GetEfbi Edge Filtering during Bus Integration
func (r *registerFdcancccrType) GetEfbi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldEfbiMask) != 0
}

// SetEfbi Edge Filtering during Bus Integration
func (r *registerFdcancccrType) SetEfbi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldEfbiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldEfbiMask)
	}
}

const (
	RegisterFdcancccrFieldTxpShift = 14
	RegisterFdcancccrFieldTxpMask  = 0x4000
)

// GetTxp TXP
func (r *registerFdcancccrType) GetTxp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldTxpMask) != 0
}

// SetTxp TXP
func (r *registerFdcancccrType) SetTxp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldTxpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldTxpMask)
	}
}

const (
	RegisterFdcancccrFieldNisoShift = 15
	RegisterFdcancccrFieldNisoMask  = 0x8000
)

// GetNiso Non ISO Operation
func (r *registerFdcancccrType) GetNiso() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcancccrFieldNisoMask) != 0
}

// SetNiso Non ISO Operation
func (r *registerFdcancccrType) SetNiso(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcancccrFieldNisoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcancccrFieldNisoMask)
	}
}

// registerFdcannbtpType FDCAN Nominal Bit Timing and Prescaler Register
type registerFdcannbtpType uint32

const (
	RegisterFdcannbtpFieldTseg2Shift = 0
	RegisterFdcannbtpFieldTseg2Mask  = 0x7f
)

// GetTseg2 Nominal Time segment after sample point
func (r *registerFdcannbtpType) GetTseg2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcannbtpFieldTseg2Mask) >> RegisterFdcannbtpFieldTseg2Shift)
}

// SetTseg2 Nominal Time segment after sample point
func (r *registerFdcannbtpType) SetTseg2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcannbtpFieldTseg2Mask)|(uint32(value)<<RegisterFdcannbtpFieldTseg2Shift))
}

const (
	RegisterFdcannbtpFieldNtseg1Shift = 8
	RegisterFdcannbtpFieldNtseg1Mask  = 0xff00
)

// GetNtseg1 Nominal Time segment before sample point
func (r *registerFdcannbtpType) GetNtseg1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcannbtpFieldNtseg1Mask) >> RegisterFdcannbtpFieldNtseg1Shift)
}

// SetNtseg1 Nominal Time segment before sample point
func (r *registerFdcannbtpType) SetNtseg1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcannbtpFieldNtseg1Mask)|(uint32(value)<<RegisterFdcannbtpFieldNtseg1Shift))
}

const (
	RegisterFdcannbtpFieldNbrpShift = 16
	RegisterFdcannbtpFieldNbrpMask  = 0x1ff0000
)

// GetNbrp Bit Rate Prescaler
func (r *registerFdcannbtpType) GetNbrp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcannbtpFieldNbrpMask) >> RegisterFdcannbtpFieldNbrpShift)
}

// SetNbrp Bit Rate Prescaler
func (r *registerFdcannbtpType) SetNbrp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcannbtpFieldNbrpMask)|(uint32(value)<<RegisterFdcannbtpFieldNbrpShift))
}

const (
	RegisterFdcannbtpFieldNsjwShift = 25
	RegisterFdcannbtpFieldNsjwMask  = 0xfe000000
)

// GetNsjw NSJW: Nominal (Re)Synchronization Jump Width
func (r *registerFdcannbtpType) GetNsjw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcannbtpFieldNsjwMask) >> RegisterFdcannbtpFieldNsjwShift)
}

// SetNsjw NSJW: Nominal (Re)Synchronization Jump Width
func (r *registerFdcannbtpType) SetNsjw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcannbtpFieldNsjwMask)|(uint32(value)<<RegisterFdcannbtpFieldNsjwShift))
}

// registerFdcantsccType FDCAN Timestamp Counter Configuration Register
type registerFdcantsccType uint32

const (
	RegisterFdcantsccFieldTssShift = 0
	RegisterFdcantsccFieldTssMask  = 0x3
)

// GetTss Timestamp Select
func (r *registerFdcantsccType) GetTss() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantsccFieldTssMask) >> RegisterFdcantsccFieldTssShift)
}

// SetTss Timestamp Select
func (r *registerFdcantsccType) SetTss(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantsccFieldTssMask)|(uint32(value)<<RegisterFdcantsccFieldTssShift))
}

const (
	RegisterFdcantsccFieldTcpShift = 16
	RegisterFdcantsccFieldTcpMask  = 0xf0000
)

// GetTcp Timestamp Counter Prescaler
func (r *registerFdcantsccType) GetTcp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantsccFieldTcpMask) >> RegisterFdcantsccFieldTcpShift)
}

// SetTcp Timestamp Counter Prescaler
func (r *registerFdcantsccType) SetTcp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantsccFieldTcpMask)|(uint32(value)<<RegisterFdcantsccFieldTcpShift))
}

// registerFdcantscvType FDCAN Timestamp Counter Value Register
type registerFdcantscvType uint32

const (
	RegisterFdcantscvFieldTscShift = 0
	RegisterFdcantscvFieldTscMask  = 0xffff
)

// GetTsc Timestamp Counter
func (r *registerFdcantscvType) GetTsc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantscvFieldTscMask) >> RegisterFdcantscvFieldTscShift)
}

// SetTsc Timestamp Counter
func (r *registerFdcantscvType) SetTsc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantscvFieldTscMask)|(uint32(value)<<RegisterFdcantscvFieldTscShift))
}

// registerFdcantoccType FDCAN Timeout Counter Configuration Register
type registerFdcantoccType uint32

const (
	RegisterFdcantoccFieldEtocShift = 0
	RegisterFdcantoccFieldEtocMask  = 0x1
)

// GetEtoc Enable Timeout Counter
func (r *registerFdcantoccType) GetEtoc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcantoccFieldEtocMask) != 0
}

// SetEtoc Enable Timeout Counter
func (r *registerFdcantoccType) SetEtoc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcantoccFieldEtocMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcantoccFieldEtocMask)
	}
}

const (
	RegisterFdcantoccFieldTosShift = 1
	RegisterFdcantoccFieldTosMask  = 0x6
)

// GetTos Timeout Select
func (r *registerFdcantoccType) GetTos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantoccFieldTosMask) >> RegisterFdcantoccFieldTosShift)
}

// SetTos Timeout Select
func (r *registerFdcantoccType) SetTos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantoccFieldTosMask)|(uint32(value)<<RegisterFdcantoccFieldTosShift))
}

const (
	RegisterFdcantoccFieldTopShift = 16
	RegisterFdcantoccFieldTopMask  = 0xffff0000
)

// GetTop Timeout Period
func (r *registerFdcantoccType) GetTop() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantoccFieldTopMask) >> RegisterFdcantoccFieldTopShift)
}

// SetTop Timeout Period
func (r *registerFdcantoccType) SetTop(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantoccFieldTopMask)|(uint32(value)<<RegisterFdcantoccFieldTopShift))
}

// registerFdcantocvType FDCAN Timeout Counter Value Register
type registerFdcantocvType uint32

const (
	RegisterFdcantocvFieldTocShift = 0
	RegisterFdcantocvFieldTocMask  = 0xffff
)

// GetToc Timeout Counter
func (r *registerFdcantocvType) GetToc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantocvFieldTocMask) >> RegisterFdcantocvFieldTocShift)
}

// SetToc Timeout Counter
func (r *registerFdcantocvType) SetToc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantocvFieldTocMask)|(uint32(value)<<RegisterFdcantocvFieldTocShift))
}

// registerFdcanecrType FDCAN Error Counter Register
type registerFdcanecrType uint32

const (
	RegisterFdcanecrFieldTecShift = 0
	RegisterFdcanecrFieldTecMask  = 0xff
)

// GetTec Transmit Error Counter
func (r *registerFdcanecrType) GetTec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanecrFieldTecMask) >> RegisterFdcanecrFieldTecShift)
}

// SetTec Transmit Error Counter
func (r *registerFdcanecrType) SetTec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanecrFieldTecMask)|(uint32(value)<<RegisterFdcanecrFieldTecShift))
}

const (
	RegisterFdcanecrFieldTrecShift = 8
	RegisterFdcanecrFieldTrecMask  = 0x7f00
)

// GetTrec Receive Error Counter
func (r *registerFdcanecrType) GetTrec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanecrFieldTrecMask) >> RegisterFdcanecrFieldTrecShift)
}

// SetTrec Receive Error Counter
func (r *registerFdcanecrType) SetTrec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanecrFieldTrecMask)|(uint32(value)<<RegisterFdcanecrFieldTrecShift))
}

const (
	RegisterFdcanecrFieldRpShift = 15
	RegisterFdcanecrFieldRpMask  = 0x8000
)

// GetRp Receive Error Passive
func (r *registerFdcanecrType) GetRp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanecrFieldRpMask) != 0
}

// SetRp Receive Error Passive
func (r *registerFdcanecrType) SetRp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanecrFieldRpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanecrFieldRpMask)
	}
}

const (
	RegisterFdcanecrFieldCelShift = 16
	RegisterFdcanecrFieldCelMask  = 0xff0000
)

// GetCel AN Error Logging
func (r *registerFdcanecrType) GetCel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanecrFieldCelMask) >> RegisterFdcanecrFieldCelShift)
}

// SetCel AN Error Logging
func (r *registerFdcanecrType) SetCel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanecrFieldCelMask)|(uint32(value)<<RegisterFdcanecrFieldCelShift))
}

// registerFdcanpsrType FDCAN Protocol Status Register
type registerFdcanpsrType uint32

const (
	RegisterFdcanpsrFieldLecShift = 0
	RegisterFdcanpsrFieldLecMask  = 0x7
)

// GetLec Last Error Code
func (r *registerFdcanpsrType) GetLec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanpsrFieldLecMask) >> RegisterFdcanpsrFieldLecShift)
}

// SetLec Last Error Code
func (r *registerFdcanpsrType) SetLec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanpsrFieldLecMask)|(uint32(value)<<RegisterFdcanpsrFieldLecShift))
}

const (
	RegisterFdcanpsrFieldActShift = 3
	RegisterFdcanpsrFieldActMask  = 0x18
)

// GetAct Activity
func (r *registerFdcanpsrType) GetAct() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanpsrFieldActMask) >> RegisterFdcanpsrFieldActShift)
}

// SetAct Activity
func (r *registerFdcanpsrType) SetAct(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanpsrFieldActMask)|(uint32(value)<<RegisterFdcanpsrFieldActShift))
}

const (
	RegisterFdcanpsrFieldEpShift = 5
	RegisterFdcanpsrFieldEpMask  = 0x20
)

// GetEp Error Passive
func (r *registerFdcanpsrType) GetEp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanpsrFieldEpMask) != 0
}

// SetEp Error Passive
func (r *registerFdcanpsrType) SetEp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanpsrFieldEpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanpsrFieldEpMask)
	}
}

const (
	RegisterFdcanpsrFieldEwShift = 6
	RegisterFdcanpsrFieldEwMask  = 0x40
)

// GetEw Warning Status
func (r *registerFdcanpsrType) GetEw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanpsrFieldEwMask) != 0
}

// SetEw Warning Status
func (r *registerFdcanpsrType) SetEw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanpsrFieldEwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanpsrFieldEwMask)
	}
}

const (
	RegisterFdcanpsrFieldBoShift = 7
	RegisterFdcanpsrFieldBoMask  = 0x80
)

// GetBo Bus_Off Status
func (r *registerFdcanpsrType) GetBo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanpsrFieldBoMask) != 0
}

// SetBo Bus_Off Status
func (r *registerFdcanpsrType) SetBo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanpsrFieldBoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanpsrFieldBoMask)
	}
}

const (
	RegisterFdcanpsrFieldDlecShift = 8
	RegisterFdcanpsrFieldDlecMask  = 0x700
)

// GetDlec Data Last Error Code
func (r *registerFdcanpsrType) GetDlec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanpsrFieldDlecMask) >> RegisterFdcanpsrFieldDlecShift)
}

// SetDlec Data Last Error Code
func (r *registerFdcanpsrType) SetDlec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanpsrFieldDlecMask)|(uint32(value)<<RegisterFdcanpsrFieldDlecShift))
}

const (
	RegisterFdcanpsrFieldResiShift = 11
	RegisterFdcanpsrFieldResiMask  = 0x800
)

// GetResi ESI flag of last received FDCAN Message
func (r *registerFdcanpsrType) GetResi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanpsrFieldResiMask) != 0
}

// SetResi ESI flag of last received FDCAN Message
func (r *registerFdcanpsrType) SetResi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanpsrFieldResiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanpsrFieldResiMask)
	}
}

const (
	RegisterFdcanpsrFieldRbrsShift = 12
	RegisterFdcanpsrFieldRbrsMask  = 0x1000
)

// GetRbrs BRS flag of last received FDCAN Message
func (r *registerFdcanpsrType) GetRbrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanpsrFieldRbrsMask) != 0
}

// SetRbrs BRS flag of last received FDCAN Message
func (r *registerFdcanpsrType) SetRbrs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanpsrFieldRbrsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanpsrFieldRbrsMask)
	}
}

const (
	RegisterFdcanpsrFieldRedlShift = 13
	RegisterFdcanpsrFieldRedlMask  = 0x2000
)

// GetRedl Received FDCAN Message
func (r *registerFdcanpsrType) GetRedl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanpsrFieldRedlMask) != 0
}

// SetRedl Received FDCAN Message
func (r *registerFdcanpsrType) SetRedl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanpsrFieldRedlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanpsrFieldRedlMask)
	}
}

const (
	RegisterFdcanpsrFieldPxeShift = 14
	RegisterFdcanpsrFieldPxeMask  = 0x4000
)

// GetPxe Protocol Exception Event
func (r *registerFdcanpsrType) GetPxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanpsrFieldPxeMask) != 0
}

// SetPxe Protocol Exception Event
func (r *registerFdcanpsrType) SetPxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanpsrFieldPxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanpsrFieldPxeMask)
	}
}

const (
	RegisterFdcanpsrFieldTdcvShift = 16
	RegisterFdcanpsrFieldTdcvMask  = 0x7f0000
)

// GetTdcv Transmitter Delay Compensation Value
func (r *registerFdcanpsrType) GetTdcv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanpsrFieldTdcvMask) >> RegisterFdcanpsrFieldTdcvShift)
}

// SetTdcv Transmitter Delay Compensation Value
func (r *registerFdcanpsrType) SetTdcv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanpsrFieldTdcvMask)|(uint32(value)<<RegisterFdcanpsrFieldTdcvShift))
}

// registerFdcantdcrType FDCAN Transmitter Delay Compensation Register
type registerFdcantdcrType uint32

const (
	RegisterFdcantdcrFieldTdcfShift = 0
	RegisterFdcantdcrFieldTdcfMask  = 0x7f
)

// GetTdcf Transmitter Delay Compensation Filter Window Length
func (r *registerFdcantdcrType) GetTdcf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantdcrFieldTdcfMask) >> RegisterFdcantdcrFieldTdcfShift)
}

// SetTdcf Transmitter Delay Compensation Filter Window Length
func (r *registerFdcantdcrType) SetTdcf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantdcrFieldTdcfMask)|(uint32(value)<<RegisterFdcantdcrFieldTdcfShift))
}

const (
	RegisterFdcantdcrFieldTdcoShift = 8
	RegisterFdcantdcrFieldTdcoMask  = 0x7f00
)

// GetTdco Transmitter Delay Compensation Offset
func (r *registerFdcantdcrType) GetTdco() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantdcrFieldTdcoMask) >> RegisterFdcantdcrFieldTdcoShift)
}

// SetTdco Transmitter Delay Compensation Offset
func (r *registerFdcantdcrType) SetTdco(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantdcrFieldTdcoMask)|(uint32(value)<<RegisterFdcantdcrFieldTdcoShift))
}

// registerFdcanirType FDCAN Interrupt Register
type registerFdcanirType uint32

const (
	RegisterFdcanirFieldRf0nShift = 0
	RegisterFdcanirFieldRf0nMask  = 0x1
)

// GetRf0n Rx FIFO 0 New Message
func (r *registerFdcanirType) GetRf0n() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldRf0nMask) != 0
}

// SetRf0n Rx FIFO 0 New Message
func (r *registerFdcanirType) SetRf0n(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldRf0nMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldRf0nMask)
	}
}

const (
	RegisterFdcanirFieldRf0wShift = 1
	RegisterFdcanirFieldRf0wMask  = 0x2
)

// GetRf0w Rx FIFO 0 Full
func (r *registerFdcanirType) GetRf0w() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldRf0wMask) != 0
}

// SetRf0w Rx FIFO 0 Full
func (r *registerFdcanirType) SetRf0w(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldRf0wMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldRf0wMask)
	}
}

const (
	RegisterFdcanirFieldRf0fShift = 2
	RegisterFdcanirFieldRf0fMask  = 0x4
)

// GetRf0f Rx FIFO 0 Full
func (r *registerFdcanirType) GetRf0f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldRf0fMask) != 0
}

// SetRf0f Rx FIFO 0 Full
func (r *registerFdcanirType) SetRf0f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldRf0fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldRf0fMask)
	}
}

const (
	RegisterFdcanirFieldRf0lShift = 3
	RegisterFdcanirFieldRf0lMask  = 0x8
)

// GetRf0l Rx FIFO 0 Message Lost
func (r *registerFdcanirType) GetRf0l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldRf0lMask) != 0
}

// SetRf0l Rx FIFO 0 Message Lost
func (r *registerFdcanirType) SetRf0l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldRf0lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldRf0lMask)
	}
}

const (
	RegisterFdcanirFieldRf1nShift = 4
	RegisterFdcanirFieldRf1nMask  = 0x10
)

// GetRf1n Rx FIFO 1 New Message
func (r *registerFdcanirType) GetRf1n() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldRf1nMask) != 0
}

// SetRf1n Rx FIFO 1 New Message
func (r *registerFdcanirType) SetRf1n(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldRf1nMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldRf1nMask)
	}
}

const (
	RegisterFdcanirFieldRf1wShift = 5
	RegisterFdcanirFieldRf1wMask  = 0x20
)

// GetRf1w Rx FIFO 1 Watermark Reached
func (r *registerFdcanirType) GetRf1w() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldRf1wMask) != 0
}

// SetRf1w Rx FIFO 1 Watermark Reached
func (r *registerFdcanirType) SetRf1w(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldRf1wMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldRf1wMask)
	}
}

const (
	RegisterFdcanirFieldRf1fShift = 6
	RegisterFdcanirFieldRf1fMask  = 0x40
)

// GetRf1f Rx FIFO 1 Watermark Reached
func (r *registerFdcanirType) GetRf1f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldRf1fMask) != 0
}

// SetRf1f Rx FIFO 1 Watermark Reached
func (r *registerFdcanirType) SetRf1f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldRf1fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldRf1fMask)
	}
}

const (
	RegisterFdcanirFieldRf1lShift = 7
	RegisterFdcanirFieldRf1lMask  = 0x80
)

// GetRf1l Rx FIFO 1 Message Lost
func (r *registerFdcanirType) GetRf1l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldRf1lMask) != 0
}

// SetRf1l Rx FIFO 1 Message Lost
func (r *registerFdcanirType) SetRf1l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldRf1lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldRf1lMask)
	}
}

const (
	RegisterFdcanirFieldHpmShift = 8
	RegisterFdcanirFieldHpmMask  = 0x100
)

// GetHpm High Priority Message
func (r *registerFdcanirType) GetHpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldHpmMask) != 0
}

// SetHpm High Priority Message
func (r *registerFdcanirType) SetHpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldHpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldHpmMask)
	}
}

const (
	RegisterFdcanirFieldTcShift = 9
	RegisterFdcanirFieldTcMask  = 0x200
)

// GetTc Transmission Completed
func (r *registerFdcanirType) GetTc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldTcMask) != 0
}

// SetTc Transmission Completed
func (r *registerFdcanirType) SetTc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldTcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldTcMask)
	}
}

const (
	RegisterFdcanirFieldTcfShift = 10
	RegisterFdcanirFieldTcfMask  = 0x400
)

// GetTcf Transmission Cancellation Finished
func (r *registerFdcanirType) GetTcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldTcfMask) != 0
}

// SetTcf Transmission Cancellation Finished
func (r *registerFdcanirType) SetTcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldTcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldTcfMask)
	}
}

const (
	RegisterFdcanirFieldTefShift = 11
	RegisterFdcanirFieldTefMask  = 0x800
)

// GetTef Tx FIFO Empty
func (r *registerFdcanirType) GetTef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldTefMask) != 0
}

// SetTef Tx FIFO Empty
func (r *registerFdcanirType) SetTef(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldTefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldTefMask)
	}
}

const (
	RegisterFdcanirFieldTefnShift = 12
	RegisterFdcanirFieldTefnMask  = 0x1000
)

// GetTefn Tx Event FIFO New Entry
func (r *registerFdcanirType) GetTefn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldTefnMask) != 0
}

// SetTefn Tx Event FIFO New Entry
func (r *registerFdcanirType) SetTefn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldTefnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldTefnMask)
	}
}

const (
	RegisterFdcanirFieldTefwShift = 13
	RegisterFdcanirFieldTefwMask  = 0x2000
)

// GetTefw Tx Event FIFO Watermark Reached
func (r *registerFdcanirType) GetTefw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldTefwMask) != 0
}

// SetTefw Tx Event FIFO Watermark Reached
func (r *registerFdcanirType) SetTefw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldTefwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldTefwMask)
	}
}

const (
	RegisterFdcanirFieldTeffShift = 14
	RegisterFdcanirFieldTeffMask  = 0x4000
)

// GetTeff Tx Event FIFO Full
func (r *registerFdcanirType) GetTeff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldTeffMask) != 0
}

// SetTeff Tx Event FIFO Full
func (r *registerFdcanirType) SetTeff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldTeffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldTeffMask)
	}
}

const (
	RegisterFdcanirFieldTeflShift = 15
	RegisterFdcanirFieldTeflMask  = 0x8000
)

// GetTefl Tx Event FIFO Element Lost
func (r *registerFdcanirType) GetTefl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldTeflMask) != 0
}

// SetTefl Tx Event FIFO Element Lost
func (r *registerFdcanirType) SetTefl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldTeflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldTeflMask)
	}
}

const (
	RegisterFdcanirFieldTswShift = 16
	RegisterFdcanirFieldTswMask  = 0x10000
)

// GetTsw Timestamp Wraparound
func (r *registerFdcanirType) GetTsw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldTswMask) != 0
}

// SetTsw Timestamp Wraparound
func (r *registerFdcanirType) SetTsw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldTswMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldTswMask)
	}
}

const (
	RegisterFdcanirFieldMrafShift = 17
	RegisterFdcanirFieldMrafMask  = 0x20000
)

// GetMraf Message RAM Access Failure
func (r *registerFdcanirType) GetMraf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldMrafMask) != 0
}

// SetMraf Message RAM Access Failure
func (r *registerFdcanirType) SetMraf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldMrafMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldMrafMask)
	}
}

const (
	RegisterFdcanirFieldTooShift = 18
	RegisterFdcanirFieldTooMask  = 0x40000
)

// GetToo Timeout Occurred
func (r *registerFdcanirType) GetToo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldTooMask) != 0
}

// SetToo Timeout Occurred
func (r *registerFdcanirType) SetToo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldTooMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldTooMask)
	}
}

const (
	RegisterFdcanirFieldDrxShift = 19
	RegisterFdcanirFieldDrxMask  = 0x80000
)

// GetDrx Message stored to Dedicated Rx Buffer
func (r *registerFdcanirType) GetDrx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldDrxMask) != 0
}

// SetDrx Message stored to Dedicated Rx Buffer
func (r *registerFdcanirType) SetDrx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldDrxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldDrxMask)
	}
}

const (
	RegisterFdcanirFieldEloShift = 22
	RegisterFdcanirFieldEloMask  = 0x400000
)

// GetElo Error Logging Overflow
func (r *registerFdcanirType) GetElo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldEloMask) != 0
}

// SetElo Error Logging Overflow
func (r *registerFdcanirType) SetElo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldEloMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldEloMask)
	}
}

const (
	RegisterFdcanirFieldEpShift = 23
	RegisterFdcanirFieldEpMask  = 0x800000
)

// GetEp Error Passive
func (r *registerFdcanirType) GetEp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldEpMask) != 0
}

// SetEp Error Passive
func (r *registerFdcanirType) SetEp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldEpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldEpMask)
	}
}

const (
	RegisterFdcanirFieldEwShift = 24
	RegisterFdcanirFieldEwMask  = 0x1000000
)

// GetEw Warning Status
func (r *registerFdcanirType) GetEw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldEwMask) != 0
}

// SetEw Warning Status
func (r *registerFdcanirType) SetEw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldEwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldEwMask)
	}
}

const (
	RegisterFdcanirFieldBoShift = 25
	RegisterFdcanirFieldBoMask  = 0x2000000
)

// GetBo Bus_Off Status
func (r *registerFdcanirType) GetBo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldBoMask) != 0
}

// SetBo Bus_Off Status
func (r *registerFdcanirType) SetBo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldBoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldBoMask)
	}
}

const (
	RegisterFdcanirFieldWdiShift = 26
	RegisterFdcanirFieldWdiMask  = 0x4000000
)

// GetWdi Watchdog Interrupt
func (r *registerFdcanirType) GetWdi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldWdiMask) != 0
}

// SetWdi Watchdog Interrupt
func (r *registerFdcanirType) SetWdi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldWdiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldWdiMask)
	}
}

const (
	RegisterFdcanirFieldPeaShift = 27
	RegisterFdcanirFieldPeaMask  = 0x8000000
)

// GetPea Protocol Error in Arbitration Phase (Nominal Bit Time is used)
func (r *registerFdcanirType) GetPea() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldPeaMask) != 0
}

// SetPea Protocol Error in Arbitration Phase (Nominal Bit Time is used)
func (r *registerFdcanirType) SetPea(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldPeaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldPeaMask)
	}
}

const (
	RegisterFdcanirFieldPedShift = 28
	RegisterFdcanirFieldPedMask  = 0x10000000
)

// GetPed Protocol Error in Data Phase (Data Bit Time is used)
func (r *registerFdcanirType) GetPed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldPedMask) != 0
}

// SetPed Protocol Error in Data Phase (Data Bit Time is used)
func (r *registerFdcanirType) SetPed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldPedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldPedMask)
	}
}

const (
	RegisterFdcanirFieldAraShift = 29
	RegisterFdcanirFieldAraMask  = 0x20000000
)

// GetAra Access to Reserved Address
func (r *registerFdcanirType) GetAra() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanirFieldAraMask) != 0
}

// SetAra Access to Reserved Address
func (r *registerFdcanirType) SetAra(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanirFieldAraMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanirFieldAraMask)
	}
}

// registerFdcanieType FDCAN Interrupt Enable Register
type registerFdcanieType uint32

const (
	RegisterFdcanieFieldRf0neShift = 0
	RegisterFdcanieFieldRf0neMask  = 0x1
)

// GetRf0ne Rx FIFO 0 New Message Enable
func (r *registerFdcanieType) GetRf0ne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldRf0neMask) != 0
}

// SetRf0ne Rx FIFO 0 New Message Enable
func (r *registerFdcanieType) SetRf0ne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldRf0neMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldRf0neMask)
	}
}

const (
	RegisterFdcanieFieldRf0weShift = 1
	RegisterFdcanieFieldRf0weMask  = 0x2
)

// GetRf0we Rx FIFO 0 Full Enable
func (r *registerFdcanieType) GetRf0we() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldRf0weMask) != 0
}

// SetRf0we Rx FIFO 0 Full Enable
func (r *registerFdcanieType) SetRf0we(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldRf0weMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldRf0weMask)
	}
}

const (
	RegisterFdcanieFieldRf0feShift = 2
	RegisterFdcanieFieldRf0feMask  = 0x4
)

// GetRf0fe Rx FIFO 0 Full Enable
func (r *registerFdcanieType) GetRf0fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldRf0feMask) != 0
}

// SetRf0fe Rx FIFO 0 Full Enable
func (r *registerFdcanieType) SetRf0fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldRf0feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldRf0feMask)
	}
}

const (
	RegisterFdcanieFieldRf0leShift = 3
	RegisterFdcanieFieldRf0leMask  = 0x8
)

// GetRf0le Rx FIFO 0 Message Lost Enable
func (r *registerFdcanieType) GetRf0le() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldRf0leMask) != 0
}

// SetRf0le Rx FIFO 0 Message Lost Enable
func (r *registerFdcanieType) SetRf0le(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldRf0leMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldRf0leMask)
	}
}

const (
	RegisterFdcanieFieldRf1neShift = 4
	RegisterFdcanieFieldRf1neMask  = 0x10
)

// GetRf1ne Rx FIFO 1 New Message Enable
func (r *registerFdcanieType) GetRf1ne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldRf1neMask) != 0
}

// SetRf1ne Rx FIFO 1 New Message Enable
func (r *registerFdcanieType) SetRf1ne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldRf1neMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldRf1neMask)
	}
}

const (
	RegisterFdcanieFieldRf1weShift = 5
	RegisterFdcanieFieldRf1weMask  = 0x20
)

// GetRf1we Rx FIFO 1 Watermark Reached Enable
func (r *registerFdcanieType) GetRf1we() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldRf1weMask) != 0
}

// SetRf1we Rx FIFO 1 Watermark Reached Enable
func (r *registerFdcanieType) SetRf1we(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldRf1weMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldRf1weMask)
	}
}

const (
	RegisterFdcanieFieldRf1feShift = 6
	RegisterFdcanieFieldRf1feMask  = 0x40
)

// GetRf1fe Rx FIFO 1 Watermark Reached Enable
func (r *registerFdcanieType) GetRf1fe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldRf1feMask) != 0
}

// SetRf1fe Rx FIFO 1 Watermark Reached Enable
func (r *registerFdcanieType) SetRf1fe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldRf1feMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldRf1feMask)
	}
}

const (
	RegisterFdcanieFieldRf1leShift = 7
	RegisterFdcanieFieldRf1leMask  = 0x80
)

// GetRf1le Rx FIFO 1 Message Lost Enable
func (r *registerFdcanieType) GetRf1le() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldRf1leMask) != 0
}

// SetRf1le Rx FIFO 1 Message Lost Enable
func (r *registerFdcanieType) SetRf1le(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldRf1leMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldRf1leMask)
	}
}

const (
	RegisterFdcanieFieldHpmeShift = 8
	RegisterFdcanieFieldHpmeMask  = 0x100
)

// GetHpme High Priority Message Enable
func (r *registerFdcanieType) GetHpme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldHpmeMask) != 0
}

// SetHpme High Priority Message Enable
func (r *registerFdcanieType) SetHpme(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldHpmeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldHpmeMask)
	}
}

const (
	RegisterFdcanieFieldTceShift = 9
	RegisterFdcanieFieldTceMask  = 0x200
)

// GetTce Transmission Completed Enable
func (r *registerFdcanieType) GetTce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldTceMask) != 0
}

// SetTce Transmission Completed Enable
func (r *registerFdcanieType) SetTce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldTceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldTceMask)
	}
}

const (
	RegisterFdcanieFieldTcfeShift = 10
	RegisterFdcanieFieldTcfeMask  = 0x400
)

// GetTcfe Transmission Cancellation Finished Enable
func (r *registerFdcanieType) GetTcfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldTcfeMask) != 0
}

// SetTcfe Transmission Cancellation Finished Enable
func (r *registerFdcanieType) SetTcfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldTcfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldTcfeMask)
	}
}

const (
	RegisterFdcanieFieldTefeShift = 11
	RegisterFdcanieFieldTefeMask  = 0x800
)

// GetTefe Tx FIFO Empty Enable
func (r *registerFdcanieType) GetTefe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldTefeMask) != 0
}

// SetTefe Tx FIFO Empty Enable
func (r *registerFdcanieType) SetTefe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldTefeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldTefeMask)
	}
}

const (
	RegisterFdcanieFieldTefneShift = 12
	RegisterFdcanieFieldTefneMask  = 0x1000
)

// GetTefne Tx Event FIFO New Entry Enable
func (r *registerFdcanieType) GetTefne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldTefneMask) != 0
}

// SetTefne Tx Event FIFO New Entry Enable
func (r *registerFdcanieType) SetTefne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldTefneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldTefneMask)
	}
}

const (
	RegisterFdcanieFieldTefweShift = 13
	RegisterFdcanieFieldTefweMask  = 0x2000
)

// GetTefwe Tx Event FIFO Watermark Reached Enable
func (r *registerFdcanieType) GetTefwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldTefweMask) != 0
}

// SetTefwe Tx Event FIFO Watermark Reached Enable
func (r *registerFdcanieType) SetTefwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldTefweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldTefweMask)
	}
}

const (
	RegisterFdcanieFieldTeffeShift = 14
	RegisterFdcanieFieldTeffeMask  = 0x4000
)

// GetTeffe Tx Event FIFO Full Enable
func (r *registerFdcanieType) GetTeffe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldTeffeMask) != 0
}

// SetTeffe Tx Event FIFO Full Enable
func (r *registerFdcanieType) SetTeffe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldTeffeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldTeffeMask)
	}
}

const (
	RegisterFdcanieFieldTefleShift = 15
	RegisterFdcanieFieldTefleMask  = 0x8000
)

// GetTefle Tx Event FIFO Element Lost Enable
func (r *registerFdcanieType) GetTefle() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldTefleMask) != 0
}

// SetTefle Tx Event FIFO Element Lost Enable
func (r *registerFdcanieType) SetTefle(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldTefleMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldTefleMask)
	}
}

const (
	RegisterFdcanieFieldTsweShift = 16
	RegisterFdcanieFieldTsweMask  = 0x10000
)

// GetTswe Timestamp Wraparound Enable
func (r *registerFdcanieType) GetTswe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldTsweMask) != 0
}

// SetTswe Timestamp Wraparound Enable
func (r *registerFdcanieType) SetTswe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldTsweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldTsweMask)
	}
}

const (
	RegisterFdcanieFieldMrafeShift = 17
	RegisterFdcanieFieldMrafeMask  = 0x20000
)

// GetMrafe Message RAM Access Failure Enable
func (r *registerFdcanieType) GetMrafe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldMrafeMask) != 0
}

// SetMrafe Message RAM Access Failure Enable
func (r *registerFdcanieType) SetMrafe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldMrafeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldMrafeMask)
	}
}

const (
	RegisterFdcanieFieldTooeShift = 18
	RegisterFdcanieFieldTooeMask  = 0x40000
)

// GetTooe Timeout Occurred Enable
func (r *registerFdcanieType) GetTooe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldTooeMask) != 0
}

// SetTooe Timeout Occurred Enable
func (r *registerFdcanieType) SetTooe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldTooeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldTooeMask)
	}
}

const (
	RegisterFdcanieFieldDrxeShift = 19
	RegisterFdcanieFieldDrxeMask  = 0x80000
)

// GetDrxe Message stored to Dedicated Rx Buffer Enable
func (r *registerFdcanieType) GetDrxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldDrxeMask) != 0
}

// SetDrxe Message stored to Dedicated Rx Buffer Enable
func (r *registerFdcanieType) SetDrxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldDrxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldDrxeMask)
	}
}

const (
	RegisterFdcanieFieldBeceShift = 20
	RegisterFdcanieFieldBeceMask  = 0x100000
)

// GetBece Bit Error Corrected Interrupt Enable
func (r *registerFdcanieType) GetBece() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldBeceMask) != 0
}

// SetBece Bit Error Corrected Interrupt Enable
func (r *registerFdcanieType) SetBece(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldBeceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldBeceMask)
	}
}

const (
	RegisterFdcanieFieldBeueShift = 21
	RegisterFdcanieFieldBeueMask  = 0x200000
)

// GetBeue Bit Error Uncorrected Interrupt Enable
func (r *registerFdcanieType) GetBeue() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldBeueMask) != 0
}

// SetBeue Bit Error Uncorrected Interrupt Enable
func (r *registerFdcanieType) SetBeue(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldBeueMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldBeueMask)
	}
}

const (
	RegisterFdcanieFieldEloeShift = 22
	RegisterFdcanieFieldEloeMask  = 0x400000
)

// GetEloe Error Logging Overflow Enable
func (r *registerFdcanieType) GetEloe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldEloeMask) != 0
}

// SetEloe Error Logging Overflow Enable
func (r *registerFdcanieType) SetEloe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldEloeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldEloeMask)
	}
}

const (
	RegisterFdcanieFieldEpeShift = 23
	RegisterFdcanieFieldEpeMask  = 0x800000
)

// GetEpe Error Passive Enable
func (r *registerFdcanieType) GetEpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldEpeMask) != 0
}

// SetEpe Error Passive Enable
func (r *registerFdcanieType) SetEpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldEpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldEpeMask)
	}
}

const (
	RegisterFdcanieFieldEweShift = 24
	RegisterFdcanieFieldEweMask  = 0x1000000
)

// GetEwe Warning Status Enable
func (r *registerFdcanieType) GetEwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldEweMask) != 0
}

// SetEwe Warning Status Enable
func (r *registerFdcanieType) SetEwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldEweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldEweMask)
	}
}

const (
	RegisterFdcanieFieldBoeShift = 25
	RegisterFdcanieFieldBoeMask  = 0x2000000
)

// GetBoe Bus_Off Status Enable
func (r *registerFdcanieType) GetBoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldBoeMask) != 0
}

// SetBoe Bus_Off Status Enable
func (r *registerFdcanieType) SetBoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldBoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldBoeMask)
	}
}

const (
	RegisterFdcanieFieldWdieShift = 26
	RegisterFdcanieFieldWdieMask  = 0x4000000
)

// GetWdie Watchdog Interrupt Enable
func (r *registerFdcanieType) GetWdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldWdieMask) != 0
}

// SetWdie Watchdog Interrupt Enable
func (r *registerFdcanieType) SetWdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldWdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldWdieMask)
	}
}

const (
	RegisterFdcanieFieldPeaeShift = 27
	RegisterFdcanieFieldPeaeMask  = 0x8000000
)

// GetPeae Protocol Error in Arbitration Phase Enable
func (r *registerFdcanieType) GetPeae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldPeaeMask) != 0
}

// SetPeae Protocol Error in Arbitration Phase Enable
func (r *registerFdcanieType) SetPeae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldPeaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldPeaeMask)
	}
}

const (
	RegisterFdcanieFieldPedeShift = 28
	RegisterFdcanieFieldPedeMask  = 0x10000000
)

// GetPede Protocol Error in Data Phase Enable
func (r *registerFdcanieType) GetPede() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldPedeMask) != 0
}

// SetPede Protocol Error in Data Phase Enable
func (r *registerFdcanieType) SetPede(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldPedeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldPedeMask)
	}
}

const (
	RegisterFdcanieFieldAraeShift = 29
	RegisterFdcanieFieldAraeMask  = 0x20000000
)

// GetArae Access to Reserved Address Enable
func (r *registerFdcanieType) GetArae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanieFieldAraeMask) != 0
}

// SetArae Access to Reserved Address Enable
func (r *registerFdcanieType) SetArae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanieFieldAraeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanieFieldAraeMask)
	}
}

// registerFdcanilsType FDCAN Interrupt Line Select Register
type registerFdcanilsType uint32

const (
	RegisterFdcanilsFieldRf0nlShift = 0
	RegisterFdcanilsFieldRf0nlMask  = 0x1
)

// GetRf0nl Rx FIFO 0 New Message Interrupt Line
func (r *registerFdcanilsType) GetRf0nl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldRf0nlMask) != 0
}

// SetRf0nl Rx FIFO 0 New Message Interrupt Line
func (r *registerFdcanilsType) SetRf0nl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldRf0nlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldRf0nlMask)
	}
}

const (
	RegisterFdcanilsFieldRf0wlShift = 1
	RegisterFdcanilsFieldRf0wlMask  = 0x2
)

// GetRf0wl Rx FIFO 0 Watermark Reached Interrupt Line
func (r *registerFdcanilsType) GetRf0wl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldRf0wlMask) != 0
}

// SetRf0wl Rx FIFO 0 Watermark Reached Interrupt Line
func (r *registerFdcanilsType) SetRf0wl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldRf0wlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldRf0wlMask)
	}
}

const (
	RegisterFdcanilsFieldRf0flShift = 2
	RegisterFdcanilsFieldRf0flMask  = 0x4
)

// GetRf0fl Rx FIFO 0 Full Interrupt Line
func (r *registerFdcanilsType) GetRf0fl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldRf0flMask) != 0
}

// SetRf0fl Rx FIFO 0 Full Interrupt Line
func (r *registerFdcanilsType) SetRf0fl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldRf0flMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldRf0flMask)
	}
}

const (
	RegisterFdcanilsFieldRf0llShift = 3
	RegisterFdcanilsFieldRf0llMask  = 0x8
)

// GetRf0ll Rx FIFO 0 Message Lost Interrupt Line
func (r *registerFdcanilsType) GetRf0ll() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldRf0llMask) != 0
}

// SetRf0ll Rx FIFO 0 Message Lost Interrupt Line
func (r *registerFdcanilsType) SetRf0ll(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldRf0llMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldRf0llMask)
	}
}

const (
	RegisterFdcanilsFieldRf1nlShift = 4
	RegisterFdcanilsFieldRf1nlMask  = 0x10
)

// GetRf1nl Rx FIFO 1 New Message Interrupt Line
func (r *registerFdcanilsType) GetRf1nl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldRf1nlMask) != 0
}

// SetRf1nl Rx FIFO 1 New Message Interrupt Line
func (r *registerFdcanilsType) SetRf1nl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldRf1nlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldRf1nlMask)
	}
}

const (
	RegisterFdcanilsFieldRf1wlShift = 5
	RegisterFdcanilsFieldRf1wlMask  = 0x20
)

// GetRf1wl Rx FIFO 1 Watermark Reached Interrupt Line
func (r *registerFdcanilsType) GetRf1wl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldRf1wlMask) != 0
}

// SetRf1wl Rx FIFO 1 Watermark Reached Interrupt Line
func (r *registerFdcanilsType) SetRf1wl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldRf1wlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldRf1wlMask)
	}
}

const (
	RegisterFdcanilsFieldRf1flShift = 6
	RegisterFdcanilsFieldRf1flMask  = 0x40
)

// GetRf1fl Rx FIFO 1 Full Interrupt Line
func (r *registerFdcanilsType) GetRf1fl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldRf1flMask) != 0
}

// SetRf1fl Rx FIFO 1 Full Interrupt Line
func (r *registerFdcanilsType) SetRf1fl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldRf1flMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldRf1flMask)
	}
}

const (
	RegisterFdcanilsFieldRf1llShift = 7
	RegisterFdcanilsFieldRf1llMask  = 0x80
)

// GetRf1ll Rx FIFO 1 Message Lost Interrupt Line
func (r *registerFdcanilsType) GetRf1ll() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldRf1llMask) != 0
}

// SetRf1ll Rx FIFO 1 Message Lost Interrupt Line
func (r *registerFdcanilsType) SetRf1ll(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldRf1llMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldRf1llMask)
	}
}

const (
	RegisterFdcanilsFieldHpmlShift = 8
	RegisterFdcanilsFieldHpmlMask  = 0x100
)

// GetHpml High Priority Message Interrupt Line
func (r *registerFdcanilsType) GetHpml() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldHpmlMask) != 0
}

// SetHpml High Priority Message Interrupt Line
func (r *registerFdcanilsType) SetHpml(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldHpmlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldHpmlMask)
	}
}

const (
	RegisterFdcanilsFieldTclShift = 9
	RegisterFdcanilsFieldTclMask  = 0x200
)

// GetTcl Transmission Completed Interrupt Line
func (r *registerFdcanilsType) GetTcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldTclMask) != 0
}

// SetTcl Transmission Completed Interrupt Line
func (r *registerFdcanilsType) SetTcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldTclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldTclMask)
	}
}

const (
	RegisterFdcanilsFieldTcflShift = 10
	RegisterFdcanilsFieldTcflMask  = 0x400
)

// GetTcfl Transmission Cancellation Finished Interrupt Line
func (r *registerFdcanilsType) GetTcfl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldTcflMask) != 0
}

// SetTcfl Transmission Cancellation Finished Interrupt Line
func (r *registerFdcanilsType) SetTcfl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldTcflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldTcflMask)
	}
}

const (
	RegisterFdcanilsFieldTeflShift = 11
	RegisterFdcanilsFieldTeflMask  = 0x800
)

// GetTefl Tx FIFO Empty Interrupt Line
func (r *registerFdcanilsType) GetTefl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldTeflMask) != 0
}

// SetTefl Tx FIFO Empty Interrupt Line
func (r *registerFdcanilsType) SetTefl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldTeflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldTeflMask)
	}
}

const (
	RegisterFdcanilsFieldTefnlShift = 12
	RegisterFdcanilsFieldTefnlMask  = 0x1000
)

// GetTefnl Tx Event FIFO New Entry Interrupt Line
func (r *registerFdcanilsType) GetTefnl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldTefnlMask) != 0
}

// SetTefnl Tx Event FIFO New Entry Interrupt Line
func (r *registerFdcanilsType) SetTefnl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldTefnlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldTefnlMask)
	}
}

const (
	RegisterFdcanilsFieldTefwlShift = 13
	RegisterFdcanilsFieldTefwlMask  = 0x2000
)

// GetTefwl Tx Event FIFO Watermark Reached Interrupt Line
func (r *registerFdcanilsType) GetTefwl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldTefwlMask) != 0
}

// SetTefwl Tx Event FIFO Watermark Reached Interrupt Line
func (r *registerFdcanilsType) SetTefwl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldTefwlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldTefwlMask)
	}
}

const (
	RegisterFdcanilsFieldTefflShift = 14
	RegisterFdcanilsFieldTefflMask  = 0x4000
)

// GetTeffl Tx Event FIFO Full Interrupt Line
func (r *registerFdcanilsType) GetTeffl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldTefflMask) != 0
}

// SetTeffl Tx Event FIFO Full Interrupt Line
func (r *registerFdcanilsType) SetTeffl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldTefflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldTefflMask)
	}
}

const (
	RegisterFdcanilsFieldTefllShift = 15
	RegisterFdcanilsFieldTefllMask  = 0x8000
)

// GetTefll Tx Event FIFO Element Lost Interrupt Line
func (r *registerFdcanilsType) GetTefll() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldTefllMask) != 0
}

// SetTefll Tx Event FIFO Element Lost Interrupt Line
func (r *registerFdcanilsType) SetTefll(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldTefllMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldTefllMask)
	}
}

const (
	RegisterFdcanilsFieldTswlShift = 16
	RegisterFdcanilsFieldTswlMask  = 0x10000
)

// GetTswl Timestamp Wraparound Interrupt Line
func (r *registerFdcanilsType) GetTswl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldTswlMask) != 0
}

// SetTswl Timestamp Wraparound Interrupt Line
func (r *registerFdcanilsType) SetTswl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldTswlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldTswlMask)
	}
}

const (
	RegisterFdcanilsFieldMraflShift = 17
	RegisterFdcanilsFieldMraflMask  = 0x20000
)

// GetMrafl Message RAM Access Failure Interrupt Line
func (r *registerFdcanilsType) GetMrafl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldMraflMask) != 0
}

// SetMrafl Message RAM Access Failure Interrupt Line
func (r *registerFdcanilsType) SetMrafl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldMraflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldMraflMask)
	}
}

const (
	RegisterFdcanilsFieldToolShift = 18
	RegisterFdcanilsFieldToolMask  = 0x40000
)

// GetTool Timeout Occurred Interrupt Line
func (r *registerFdcanilsType) GetTool() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldToolMask) != 0
}

// SetTool Timeout Occurred Interrupt Line
func (r *registerFdcanilsType) SetTool(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldToolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldToolMask)
	}
}

const (
	RegisterFdcanilsFieldDrxlShift = 19
	RegisterFdcanilsFieldDrxlMask  = 0x80000
)

// GetDrxl Message stored to Dedicated Rx Buffer Interrupt Line
func (r *registerFdcanilsType) GetDrxl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldDrxlMask) != 0
}

// SetDrxl Message stored to Dedicated Rx Buffer Interrupt Line
func (r *registerFdcanilsType) SetDrxl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldDrxlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldDrxlMask)
	}
}

const (
	RegisterFdcanilsFieldBeclShift = 20
	RegisterFdcanilsFieldBeclMask  = 0x100000
)

// GetBecl Bit Error Corrected Interrupt Line
func (r *registerFdcanilsType) GetBecl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldBeclMask) != 0
}

// SetBecl Bit Error Corrected Interrupt Line
func (r *registerFdcanilsType) SetBecl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldBeclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldBeclMask)
	}
}

const (
	RegisterFdcanilsFieldBeulShift = 21
	RegisterFdcanilsFieldBeulMask  = 0x200000
)

// GetBeul Bit Error Uncorrected Interrupt Line
func (r *registerFdcanilsType) GetBeul() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldBeulMask) != 0
}

// SetBeul Bit Error Uncorrected Interrupt Line
func (r *registerFdcanilsType) SetBeul(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldBeulMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldBeulMask)
	}
}

const (
	RegisterFdcanilsFieldElolShift = 22
	RegisterFdcanilsFieldElolMask  = 0x400000
)

// GetElol Error Logging Overflow Interrupt Line
func (r *registerFdcanilsType) GetElol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldElolMask) != 0
}

// SetElol Error Logging Overflow Interrupt Line
func (r *registerFdcanilsType) SetElol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldElolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldElolMask)
	}
}

const (
	RegisterFdcanilsFieldEplShift = 23
	RegisterFdcanilsFieldEplMask  = 0x800000
)

// GetEpl Error Passive Interrupt Line
func (r *registerFdcanilsType) GetEpl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldEplMask) != 0
}

// SetEpl Error Passive Interrupt Line
func (r *registerFdcanilsType) SetEpl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldEplMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldEplMask)
	}
}

const (
	RegisterFdcanilsFieldEwlShift = 24
	RegisterFdcanilsFieldEwlMask  = 0x1000000
)

// GetEwl Warning Status Interrupt Line
func (r *registerFdcanilsType) GetEwl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldEwlMask) != 0
}

// SetEwl Warning Status Interrupt Line
func (r *registerFdcanilsType) SetEwl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldEwlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldEwlMask)
	}
}

const (
	RegisterFdcanilsFieldBolShift = 25
	RegisterFdcanilsFieldBolMask  = 0x2000000
)

// GetBol Bus_Off Status
func (r *registerFdcanilsType) GetBol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldBolMask) != 0
}

// SetBol Bus_Off Status
func (r *registerFdcanilsType) SetBol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldBolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldBolMask)
	}
}

const (
	RegisterFdcanilsFieldWdilShift = 26
	RegisterFdcanilsFieldWdilMask  = 0x4000000
)

// GetWdil Watchdog Interrupt Line
func (r *registerFdcanilsType) GetWdil() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldWdilMask) != 0
}

// SetWdil Watchdog Interrupt Line
func (r *registerFdcanilsType) SetWdil(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldWdilMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldWdilMask)
	}
}

const (
	RegisterFdcanilsFieldPealShift = 27
	RegisterFdcanilsFieldPealMask  = 0x8000000
)

// GetPeal Protocol Error in Arbitration Phase Line
func (r *registerFdcanilsType) GetPeal() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldPealMask) != 0
}

// SetPeal Protocol Error in Arbitration Phase Line
func (r *registerFdcanilsType) SetPeal(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldPealMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldPealMask)
	}
}

const (
	RegisterFdcanilsFieldPedlShift = 28
	RegisterFdcanilsFieldPedlMask  = 0x10000000
)

// GetPedl Protocol Error in Data Phase Line
func (r *registerFdcanilsType) GetPedl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldPedlMask) != 0
}

// SetPedl Protocol Error in Data Phase Line
func (r *registerFdcanilsType) SetPedl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldPedlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldPedlMask)
	}
}

const (
	RegisterFdcanilsFieldAralShift = 29
	RegisterFdcanilsFieldAralMask  = 0x20000000
)

// GetAral Access to Reserved Address Line
func (r *registerFdcanilsType) GetAral() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanilsFieldAralMask) != 0
}

// SetAral Access to Reserved Address Line
func (r *registerFdcanilsType) SetAral(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanilsFieldAralMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanilsFieldAralMask)
	}
}

// registerFdcanileType FDCAN Interrupt Line Enable Register
type registerFdcanileType uint32

const (
	RegisterFdcanileFieldEint0Shift = 0
	RegisterFdcanileFieldEint0Mask  = 0x1
)

// GetEint0 Enable Interrupt Line 0
func (r *registerFdcanileType) GetEint0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanileFieldEint0Mask) != 0
}

// SetEint0 Enable Interrupt Line 0
func (r *registerFdcanileType) SetEint0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanileFieldEint0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanileFieldEint0Mask)
	}
}

const (
	RegisterFdcanileFieldEint1Shift = 1
	RegisterFdcanileFieldEint1Mask  = 0x2
)

// GetEint1 Enable Interrupt Line 1
func (r *registerFdcanileType) GetEint1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanileFieldEint1Mask) != 0
}

// SetEint1 Enable Interrupt Line 1
func (r *registerFdcanileType) SetEint1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanileFieldEint1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanileFieldEint1Mask)
	}
}

// registerFdcangfcType FDCAN Global Filter Configuration Register
type registerFdcangfcType uint32

const (
	RegisterFdcangfcFieldRrfeShift = 0
	RegisterFdcangfcFieldRrfeMask  = 0x1
)

// GetRrfe Reject Remote Frames Extended
func (r *registerFdcangfcType) GetRrfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcangfcFieldRrfeMask) != 0
}

// SetRrfe Reject Remote Frames Extended
func (r *registerFdcangfcType) SetRrfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcangfcFieldRrfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcangfcFieldRrfeMask)
	}
}

const (
	RegisterFdcangfcFieldRrfsShift = 1
	RegisterFdcangfcFieldRrfsMask  = 0x2
)

// GetRrfs Reject Remote Frames Standard
func (r *registerFdcangfcType) GetRrfs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcangfcFieldRrfsMask) != 0
}

// SetRrfs Reject Remote Frames Standard
func (r *registerFdcangfcType) SetRrfs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcangfcFieldRrfsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcangfcFieldRrfsMask)
	}
}

const (
	RegisterFdcangfcFieldAnfeShift = 2
	RegisterFdcangfcFieldAnfeMask  = 0xc
)

// GetAnfe Accept Non-matching Frames Extended
func (r *registerFdcangfcType) GetAnfe() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcangfcFieldAnfeMask) >> RegisterFdcangfcFieldAnfeShift)
}

// SetAnfe Accept Non-matching Frames Extended
func (r *registerFdcangfcType) SetAnfe(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcangfcFieldAnfeMask)|(uint32(value)<<RegisterFdcangfcFieldAnfeShift))
}

const (
	RegisterFdcangfcFieldAnfsShift = 4
	RegisterFdcangfcFieldAnfsMask  = 0x30
)

// GetAnfs Accept Non-matching Frames Standard
func (r *registerFdcangfcType) GetAnfs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcangfcFieldAnfsMask) >> RegisterFdcangfcFieldAnfsShift)
}

// SetAnfs Accept Non-matching Frames Standard
func (r *registerFdcangfcType) SetAnfs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcangfcFieldAnfsMask)|(uint32(value)<<RegisterFdcangfcFieldAnfsShift))
}

// registerFdcansidfcType FDCAN Standard ID Filter Configuration Register
type registerFdcansidfcType uint32

const (
	RegisterFdcansidfcFieldFlssaShift = 2
	RegisterFdcansidfcFieldFlssaMask  = 0xfffc
)

// GetFlssa Filter List Standard Start Address
func (r *registerFdcansidfcType) GetFlssa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcansidfcFieldFlssaMask) >> RegisterFdcansidfcFieldFlssaShift)
}

// SetFlssa Filter List Standard Start Address
func (r *registerFdcansidfcType) SetFlssa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcansidfcFieldFlssaMask)|(uint32(value)<<RegisterFdcansidfcFieldFlssaShift))
}

const (
	RegisterFdcansidfcFieldLssShift = 16
	RegisterFdcansidfcFieldLssMask  = 0xff0000
)

// GetLss List Size Standard
func (r *registerFdcansidfcType) GetLss() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcansidfcFieldLssMask) >> RegisterFdcansidfcFieldLssShift)
}

// SetLss List Size Standard
func (r *registerFdcansidfcType) SetLss(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcansidfcFieldLssMask)|(uint32(value)<<RegisterFdcansidfcFieldLssShift))
}

// registerFdcanxidfcType FDCAN Extended ID Filter Configuration Register
type registerFdcanxidfcType uint32

const (
	RegisterFdcanxidfcFieldFlesaShift = 2
	RegisterFdcanxidfcFieldFlesaMask  = 0xfffc
)

// GetFlesa Filter List Standard Start Address
func (r *registerFdcanxidfcType) GetFlesa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanxidfcFieldFlesaMask) >> RegisterFdcanxidfcFieldFlesaShift)
}

// SetFlesa Filter List Standard Start Address
func (r *registerFdcanxidfcType) SetFlesa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanxidfcFieldFlesaMask)|(uint32(value)<<RegisterFdcanxidfcFieldFlesaShift))
}

const (
	RegisterFdcanxidfcFieldLseShift = 16
	RegisterFdcanxidfcFieldLseMask  = 0xff0000
)

// GetLse List Size Extended
func (r *registerFdcanxidfcType) GetLse() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanxidfcFieldLseMask) >> RegisterFdcanxidfcFieldLseShift)
}

// SetLse List Size Extended
func (r *registerFdcanxidfcType) SetLse(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanxidfcFieldLseMask)|(uint32(value)<<RegisterFdcanxidfcFieldLseShift))
}

// registerFdcanxidamType FDCAN Extended ID and Mask Register
type registerFdcanxidamType uint32

const (
	RegisterFdcanxidamFieldEidmShift = 0
	RegisterFdcanxidamFieldEidmMask  = 0x1fffffff
)

// GetEidm Extended ID Mask
func (r *registerFdcanxidamType) GetEidm() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanxidamFieldEidmMask) >> RegisterFdcanxidamFieldEidmShift)
}

// SetEidm Extended ID Mask
func (r *registerFdcanxidamType) SetEidm(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanxidamFieldEidmMask)|(uint32(value)<<RegisterFdcanxidamFieldEidmShift))
}

// registerFdcanhpmsType FDCAN High Priority Message Status Register
type registerFdcanhpmsType uint32

const (
	RegisterFdcanhpmsFieldBidxShift = 0
	RegisterFdcanhpmsFieldBidxMask  = 0x3f
)

// GetBidx Buffer Index
func (r *registerFdcanhpmsType) GetBidx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanhpmsFieldBidxMask) >> RegisterFdcanhpmsFieldBidxShift)
}

// SetBidx Buffer Index
func (r *registerFdcanhpmsType) SetBidx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanhpmsFieldBidxMask)|(uint32(value)<<RegisterFdcanhpmsFieldBidxShift))
}

const (
	RegisterFdcanhpmsFieldMsiShift = 6
	RegisterFdcanhpmsFieldMsiMask  = 0xc0
)

// GetMsi Message Storage Indicator
func (r *registerFdcanhpmsType) GetMsi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanhpmsFieldMsiMask) >> RegisterFdcanhpmsFieldMsiShift)
}

// SetMsi Message Storage Indicator
func (r *registerFdcanhpmsType) SetMsi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanhpmsFieldMsiMask)|(uint32(value)<<RegisterFdcanhpmsFieldMsiShift))
}

const (
	RegisterFdcanhpmsFieldFidxShift = 8
	RegisterFdcanhpmsFieldFidxMask  = 0x7f00
)

// GetFidx Filter Index
func (r *registerFdcanhpmsType) GetFidx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanhpmsFieldFidxMask) >> RegisterFdcanhpmsFieldFidxShift)
}

// SetFidx Filter Index
func (r *registerFdcanhpmsType) SetFidx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanhpmsFieldFidxMask)|(uint32(value)<<RegisterFdcanhpmsFieldFidxShift))
}

const (
	RegisterFdcanhpmsFieldFlstShift = 15
	RegisterFdcanhpmsFieldFlstMask  = 0x8000
)

// GetFlst Filter List
func (r *registerFdcanhpmsType) GetFlst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanhpmsFieldFlstMask) != 0
}

// SetFlst Filter List
func (r *registerFdcanhpmsType) SetFlst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanhpmsFieldFlstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanhpmsFieldFlstMask)
	}
}

// registerFdcanndat1Type FDCAN New Data 1 Register
type registerFdcanndat1Type uint32

const (
	RegisterFdcanndat1FieldNd0Shift = 0
	RegisterFdcanndat1FieldNd0Mask  = 0x1
)

// GetNd0 New data
func (r *registerFdcanndat1Type) GetNd0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd0Mask) != 0
}

// SetNd0 New data
func (r *registerFdcanndat1Type) SetNd0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd0Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd1Shift = 1
	RegisterFdcanndat1FieldNd1Mask  = 0x2
)

// GetNd1 New data
func (r *registerFdcanndat1Type) GetNd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd1Mask) != 0
}

// SetNd1 New data
func (r *registerFdcanndat1Type) SetNd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd1Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd2Shift = 2
	RegisterFdcanndat1FieldNd2Mask  = 0x4
)

// GetNd2 New data
func (r *registerFdcanndat1Type) GetNd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd2Mask) != 0
}

// SetNd2 New data
func (r *registerFdcanndat1Type) SetNd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd2Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd3Shift = 3
	RegisterFdcanndat1FieldNd3Mask  = 0x8
)

// GetNd3 New data
func (r *registerFdcanndat1Type) GetNd3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd3Mask) != 0
}

// SetNd3 New data
func (r *registerFdcanndat1Type) SetNd3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd3Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd4Shift = 4
	RegisterFdcanndat1FieldNd4Mask  = 0x10
)

// GetNd4 New data
func (r *registerFdcanndat1Type) GetNd4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd4Mask) != 0
}

// SetNd4 New data
func (r *registerFdcanndat1Type) SetNd4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd4Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd5Shift = 5
	RegisterFdcanndat1FieldNd5Mask  = 0x20
)

// GetNd5 New data
func (r *registerFdcanndat1Type) GetNd5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd5Mask) != 0
}

// SetNd5 New data
func (r *registerFdcanndat1Type) SetNd5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd5Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd6Shift = 6
	RegisterFdcanndat1FieldNd6Mask  = 0x40
)

// GetNd6 New data
func (r *registerFdcanndat1Type) GetNd6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd6Mask) != 0
}

// SetNd6 New data
func (r *registerFdcanndat1Type) SetNd6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd6Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd7Shift = 7
	RegisterFdcanndat1FieldNd7Mask  = 0x80
)

// GetNd7 New data
func (r *registerFdcanndat1Type) GetNd7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd7Mask) != 0
}

// SetNd7 New data
func (r *registerFdcanndat1Type) SetNd7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd7Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd8Shift = 8
	RegisterFdcanndat1FieldNd8Mask  = 0x100
)

// GetNd8 New data
func (r *registerFdcanndat1Type) GetNd8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd8Mask) != 0
}

// SetNd8 New data
func (r *registerFdcanndat1Type) SetNd8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd8Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd9Shift = 9
	RegisterFdcanndat1FieldNd9Mask  = 0x200
)

// GetNd9 New data
func (r *registerFdcanndat1Type) GetNd9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd9Mask) != 0
}

// SetNd9 New data
func (r *registerFdcanndat1Type) SetNd9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd9Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd10Shift = 10
	RegisterFdcanndat1FieldNd10Mask  = 0x400
)

// GetNd10 New data
func (r *registerFdcanndat1Type) GetNd10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd10Mask) != 0
}

// SetNd10 New data
func (r *registerFdcanndat1Type) SetNd10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd10Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd11Shift = 11
	RegisterFdcanndat1FieldNd11Mask  = 0x800
)

// GetNd11 New data
func (r *registerFdcanndat1Type) GetNd11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd11Mask) != 0
}

// SetNd11 New data
func (r *registerFdcanndat1Type) SetNd11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd11Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd12Shift = 12
	RegisterFdcanndat1FieldNd12Mask  = 0x1000
)

// GetNd12 New data
func (r *registerFdcanndat1Type) GetNd12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd12Mask) != 0
}

// SetNd12 New data
func (r *registerFdcanndat1Type) SetNd12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd12Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd13Shift = 13
	RegisterFdcanndat1FieldNd13Mask  = 0x2000
)

// GetNd13 New data
func (r *registerFdcanndat1Type) GetNd13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd13Mask) != 0
}

// SetNd13 New data
func (r *registerFdcanndat1Type) SetNd13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd13Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd14Shift = 14
	RegisterFdcanndat1FieldNd14Mask  = 0x4000
)

// GetNd14 New data
func (r *registerFdcanndat1Type) GetNd14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd14Mask) != 0
}

// SetNd14 New data
func (r *registerFdcanndat1Type) SetNd14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd14Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd15Shift = 15
	RegisterFdcanndat1FieldNd15Mask  = 0x8000
)

// GetNd15 New data
func (r *registerFdcanndat1Type) GetNd15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd15Mask) != 0
}

// SetNd15 New data
func (r *registerFdcanndat1Type) SetNd15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd15Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd16Shift = 16
	RegisterFdcanndat1FieldNd16Mask  = 0x10000
)

// GetNd16 New data
func (r *registerFdcanndat1Type) GetNd16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd16Mask) != 0
}

// SetNd16 New data
func (r *registerFdcanndat1Type) SetNd16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd16Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd17Shift = 17
	RegisterFdcanndat1FieldNd17Mask  = 0x20000
)

// GetNd17 New data
func (r *registerFdcanndat1Type) GetNd17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd17Mask) != 0
}

// SetNd17 New data
func (r *registerFdcanndat1Type) SetNd17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd17Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd18Shift = 18
	RegisterFdcanndat1FieldNd18Mask  = 0x40000
)

// GetNd18 New data
func (r *registerFdcanndat1Type) GetNd18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd18Mask) != 0
}

// SetNd18 New data
func (r *registerFdcanndat1Type) SetNd18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd18Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd19Shift = 19
	RegisterFdcanndat1FieldNd19Mask  = 0x80000
)

// GetNd19 New data
func (r *registerFdcanndat1Type) GetNd19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd19Mask) != 0
}

// SetNd19 New data
func (r *registerFdcanndat1Type) SetNd19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd19Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd20Shift = 20
	RegisterFdcanndat1FieldNd20Mask  = 0x100000
)

// GetNd20 New data
func (r *registerFdcanndat1Type) GetNd20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd20Mask) != 0
}

// SetNd20 New data
func (r *registerFdcanndat1Type) SetNd20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd20Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd21Shift = 21
	RegisterFdcanndat1FieldNd21Mask  = 0x200000
)

// GetNd21 New data
func (r *registerFdcanndat1Type) GetNd21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd21Mask) != 0
}

// SetNd21 New data
func (r *registerFdcanndat1Type) SetNd21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd21Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd22Shift = 22
	RegisterFdcanndat1FieldNd22Mask  = 0x400000
)

// GetNd22 New data
func (r *registerFdcanndat1Type) GetNd22() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd22Mask) != 0
}

// SetNd22 New data
func (r *registerFdcanndat1Type) SetNd22(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd22Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd22Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd23Shift = 23
	RegisterFdcanndat1FieldNd23Mask  = 0x800000
)

// GetNd23 New data
func (r *registerFdcanndat1Type) GetNd23() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd23Mask) != 0
}

// SetNd23 New data
func (r *registerFdcanndat1Type) SetNd23(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd23Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd23Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd24Shift = 24
	RegisterFdcanndat1FieldNd24Mask  = 0x1000000
)

// GetNd24 New data
func (r *registerFdcanndat1Type) GetNd24() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd24Mask) != 0
}

// SetNd24 New data
func (r *registerFdcanndat1Type) SetNd24(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd24Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd24Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd25Shift = 25
	RegisterFdcanndat1FieldNd25Mask  = 0x2000000
)

// GetNd25 New data
func (r *registerFdcanndat1Type) GetNd25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd25Mask) != 0
}

// SetNd25 New data
func (r *registerFdcanndat1Type) SetNd25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd25Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd26Shift = 26
	RegisterFdcanndat1FieldNd26Mask  = 0x4000000
)

// GetNd26 New data
func (r *registerFdcanndat1Type) GetNd26() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd26Mask) != 0
}

// SetNd26 New data
func (r *registerFdcanndat1Type) SetNd26(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd26Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd26Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd27Shift = 27
	RegisterFdcanndat1FieldNd27Mask  = 0x8000000
)

// GetNd27 New data
func (r *registerFdcanndat1Type) GetNd27() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd27Mask) != 0
}

// SetNd27 New data
func (r *registerFdcanndat1Type) SetNd27(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd27Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd27Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd28Shift = 28
	RegisterFdcanndat1FieldNd28Mask  = 0x10000000
)

// GetNd28 New data
func (r *registerFdcanndat1Type) GetNd28() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd28Mask) != 0
}

// SetNd28 New data
func (r *registerFdcanndat1Type) SetNd28(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd28Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd28Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd29Shift = 29
	RegisterFdcanndat1FieldNd29Mask  = 0x20000000
)

// GetNd29 New data
func (r *registerFdcanndat1Type) GetNd29() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd29Mask) != 0
}

// SetNd29 New data
func (r *registerFdcanndat1Type) SetNd29(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd29Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd29Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd30Shift = 30
	RegisterFdcanndat1FieldNd30Mask  = 0x40000000
)

// GetNd30 New data
func (r *registerFdcanndat1Type) GetNd30() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd30Mask) != 0
}

// SetNd30 New data
func (r *registerFdcanndat1Type) SetNd30(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd30Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd30Mask)
	}
}

const (
	RegisterFdcanndat1FieldNd31Shift = 31
	RegisterFdcanndat1FieldNd31Mask  = 0x80000000
)

// GetNd31 New data
func (r *registerFdcanndat1Type) GetNd31() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat1FieldNd31Mask) != 0
}

// SetNd31 New data
func (r *registerFdcanndat1Type) SetNd31(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat1FieldNd31Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat1FieldNd31Mask)
	}
}

// registerFdcanndat2Type FDCAN New Data 2 Register
type registerFdcanndat2Type uint32

const (
	RegisterFdcanndat2FieldNd32Shift = 0
	RegisterFdcanndat2FieldNd32Mask  = 0x1
)

// GetNd32 New data
func (r *registerFdcanndat2Type) GetNd32() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd32Mask) != 0
}

// SetNd32 New data
func (r *registerFdcanndat2Type) SetNd32(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd32Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd32Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd33Shift = 1
	RegisterFdcanndat2FieldNd33Mask  = 0x2
)

// GetNd33 New data
func (r *registerFdcanndat2Type) GetNd33() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd33Mask) != 0
}

// SetNd33 New data
func (r *registerFdcanndat2Type) SetNd33(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd33Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd33Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd34Shift = 2
	RegisterFdcanndat2FieldNd34Mask  = 0x4
)

// GetNd34 New data
func (r *registerFdcanndat2Type) GetNd34() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd34Mask) != 0
}

// SetNd34 New data
func (r *registerFdcanndat2Type) SetNd34(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd34Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd34Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd35Shift = 3
	RegisterFdcanndat2FieldNd35Mask  = 0x8
)

// GetNd35 New data
func (r *registerFdcanndat2Type) GetNd35() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd35Mask) != 0
}

// SetNd35 New data
func (r *registerFdcanndat2Type) SetNd35(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd35Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd35Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd36Shift = 4
	RegisterFdcanndat2FieldNd36Mask  = 0x10
)

// GetNd36 New data
func (r *registerFdcanndat2Type) GetNd36() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd36Mask) != 0
}

// SetNd36 New data
func (r *registerFdcanndat2Type) SetNd36(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd36Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd36Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd37Shift = 5
	RegisterFdcanndat2FieldNd37Mask  = 0x20
)

// GetNd37 New data
func (r *registerFdcanndat2Type) GetNd37() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd37Mask) != 0
}

// SetNd37 New data
func (r *registerFdcanndat2Type) SetNd37(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd37Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd37Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd38Shift = 6
	RegisterFdcanndat2FieldNd38Mask  = 0x40
)

// GetNd38 New data
func (r *registerFdcanndat2Type) GetNd38() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd38Mask) != 0
}

// SetNd38 New data
func (r *registerFdcanndat2Type) SetNd38(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd38Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd38Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd39Shift = 7
	RegisterFdcanndat2FieldNd39Mask  = 0x80
)

// GetNd39 New data
func (r *registerFdcanndat2Type) GetNd39() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd39Mask) != 0
}

// SetNd39 New data
func (r *registerFdcanndat2Type) SetNd39(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd39Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd39Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd40Shift = 8
	RegisterFdcanndat2FieldNd40Mask  = 0x100
)

// GetNd40 New data
func (r *registerFdcanndat2Type) GetNd40() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd40Mask) != 0
}

// SetNd40 New data
func (r *registerFdcanndat2Type) SetNd40(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd40Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd40Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd41Shift = 9
	RegisterFdcanndat2FieldNd41Mask  = 0x200
)

// GetNd41 New data
func (r *registerFdcanndat2Type) GetNd41() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd41Mask) != 0
}

// SetNd41 New data
func (r *registerFdcanndat2Type) SetNd41(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd41Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd41Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd42Shift = 10
	RegisterFdcanndat2FieldNd42Mask  = 0x400
)

// GetNd42 New data
func (r *registerFdcanndat2Type) GetNd42() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd42Mask) != 0
}

// SetNd42 New data
func (r *registerFdcanndat2Type) SetNd42(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd42Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd42Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd43Shift = 11
	RegisterFdcanndat2FieldNd43Mask  = 0x800
)

// GetNd43 New data
func (r *registerFdcanndat2Type) GetNd43() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd43Mask) != 0
}

// SetNd43 New data
func (r *registerFdcanndat2Type) SetNd43(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd43Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd43Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd44Shift = 12
	RegisterFdcanndat2FieldNd44Mask  = 0x1000
)

// GetNd44 New data
func (r *registerFdcanndat2Type) GetNd44() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd44Mask) != 0
}

// SetNd44 New data
func (r *registerFdcanndat2Type) SetNd44(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd44Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd44Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd45Shift = 13
	RegisterFdcanndat2FieldNd45Mask  = 0x2000
)

// GetNd45 New data
func (r *registerFdcanndat2Type) GetNd45() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd45Mask) != 0
}

// SetNd45 New data
func (r *registerFdcanndat2Type) SetNd45(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd45Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd45Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd46Shift = 14
	RegisterFdcanndat2FieldNd46Mask  = 0x4000
)

// GetNd46 New data
func (r *registerFdcanndat2Type) GetNd46() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd46Mask) != 0
}

// SetNd46 New data
func (r *registerFdcanndat2Type) SetNd46(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd46Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd46Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd47Shift = 15
	RegisterFdcanndat2FieldNd47Mask  = 0x8000
)

// GetNd47 New data
func (r *registerFdcanndat2Type) GetNd47() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd47Mask) != 0
}

// SetNd47 New data
func (r *registerFdcanndat2Type) SetNd47(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd47Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd47Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd48Shift = 16
	RegisterFdcanndat2FieldNd48Mask  = 0x10000
)

// GetNd48 New data
func (r *registerFdcanndat2Type) GetNd48() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd48Mask) != 0
}

// SetNd48 New data
func (r *registerFdcanndat2Type) SetNd48(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd48Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd48Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd49Shift = 17
	RegisterFdcanndat2FieldNd49Mask  = 0x20000
)

// GetNd49 New data
func (r *registerFdcanndat2Type) GetNd49() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd49Mask) != 0
}

// SetNd49 New data
func (r *registerFdcanndat2Type) SetNd49(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd49Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd49Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd50Shift = 18
	RegisterFdcanndat2FieldNd50Mask  = 0x40000
)

// GetNd50 New data
func (r *registerFdcanndat2Type) GetNd50() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd50Mask) != 0
}

// SetNd50 New data
func (r *registerFdcanndat2Type) SetNd50(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd50Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd50Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd51Shift = 19
	RegisterFdcanndat2FieldNd51Mask  = 0x80000
)

// GetNd51 New data
func (r *registerFdcanndat2Type) GetNd51() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd51Mask) != 0
}

// SetNd51 New data
func (r *registerFdcanndat2Type) SetNd51(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd51Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd51Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd52Shift = 20
	RegisterFdcanndat2FieldNd52Mask  = 0x100000
)

// GetNd52 New data
func (r *registerFdcanndat2Type) GetNd52() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd52Mask) != 0
}

// SetNd52 New data
func (r *registerFdcanndat2Type) SetNd52(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd52Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd52Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd53Shift = 21
	RegisterFdcanndat2FieldNd53Mask  = 0x200000
)

// GetNd53 New data
func (r *registerFdcanndat2Type) GetNd53() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd53Mask) != 0
}

// SetNd53 New data
func (r *registerFdcanndat2Type) SetNd53(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd53Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd53Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd54Shift = 22
	RegisterFdcanndat2FieldNd54Mask  = 0x400000
)

// GetNd54 New data
func (r *registerFdcanndat2Type) GetNd54() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd54Mask) != 0
}

// SetNd54 New data
func (r *registerFdcanndat2Type) SetNd54(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd54Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd54Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd55Shift = 23
	RegisterFdcanndat2FieldNd55Mask  = 0x800000
)

// GetNd55 New data
func (r *registerFdcanndat2Type) GetNd55() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd55Mask) != 0
}

// SetNd55 New data
func (r *registerFdcanndat2Type) SetNd55(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd55Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd55Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd56Shift = 24
	RegisterFdcanndat2FieldNd56Mask  = 0x1000000
)

// GetNd56 New data
func (r *registerFdcanndat2Type) GetNd56() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd56Mask) != 0
}

// SetNd56 New data
func (r *registerFdcanndat2Type) SetNd56(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd56Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd56Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd57Shift = 25
	RegisterFdcanndat2FieldNd57Mask  = 0x2000000
)

// GetNd57 New data
func (r *registerFdcanndat2Type) GetNd57() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd57Mask) != 0
}

// SetNd57 New data
func (r *registerFdcanndat2Type) SetNd57(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd57Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd57Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd58Shift = 26
	RegisterFdcanndat2FieldNd58Mask  = 0x4000000
)

// GetNd58 New data
func (r *registerFdcanndat2Type) GetNd58() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd58Mask) != 0
}

// SetNd58 New data
func (r *registerFdcanndat2Type) SetNd58(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd58Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd58Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd59Shift = 27
	RegisterFdcanndat2FieldNd59Mask  = 0x8000000
)

// GetNd59 New data
func (r *registerFdcanndat2Type) GetNd59() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd59Mask) != 0
}

// SetNd59 New data
func (r *registerFdcanndat2Type) SetNd59(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd59Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd59Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd60Shift = 28
	RegisterFdcanndat2FieldNd60Mask  = 0x10000000
)

// GetNd60 New data
func (r *registerFdcanndat2Type) GetNd60() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd60Mask) != 0
}

// SetNd60 New data
func (r *registerFdcanndat2Type) SetNd60(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd60Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd60Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd61Shift = 29
	RegisterFdcanndat2FieldNd61Mask  = 0x20000000
)

// GetNd61 New data
func (r *registerFdcanndat2Type) GetNd61() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd61Mask) != 0
}

// SetNd61 New data
func (r *registerFdcanndat2Type) SetNd61(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd61Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd61Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd62Shift = 30
	RegisterFdcanndat2FieldNd62Mask  = 0x40000000
)

// GetNd62 New data
func (r *registerFdcanndat2Type) GetNd62() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd62Mask) != 0
}

// SetNd62 New data
func (r *registerFdcanndat2Type) SetNd62(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd62Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd62Mask)
	}
}

const (
	RegisterFdcanndat2FieldNd63Shift = 31
	RegisterFdcanndat2FieldNd63Mask  = 0x80000000
)

// GetNd63 New data
func (r *registerFdcanndat2Type) GetNd63() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanndat2FieldNd63Mask) != 0
}

// SetNd63 New data
func (r *registerFdcanndat2Type) SetNd63(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanndat2FieldNd63Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanndat2FieldNd63Mask)
	}
}

// registerFdcanrxf0cType FDCAN Rx FIFO 0 Configuration Register
type registerFdcanrxf0cType uint32

const (
	RegisterFdcanrxf0cFieldF0saShift = 2
	RegisterFdcanrxf0cFieldF0saMask  = 0xfffc
)

// GetF0sa Rx FIFO 0 Start Address
func (r *registerFdcanrxf0cType) GetF0sa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf0cFieldF0saMask) >> RegisterFdcanrxf0cFieldF0saShift)
}

// SetF0sa Rx FIFO 0 Start Address
func (r *registerFdcanrxf0cType) SetF0sa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf0cFieldF0saMask)|(uint32(value)<<RegisterFdcanrxf0cFieldF0saShift))
}

const (
	RegisterFdcanrxf0cFieldF0sShift = 16
	RegisterFdcanrxf0cFieldF0sMask  = 0xff0000
)

// GetF0s Rx FIFO 0 Size
func (r *registerFdcanrxf0cType) GetF0s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf0cFieldF0sMask) >> RegisterFdcanrxf0cFieldF0sShift)
}

// SetF0s Rx FIFO 0 Size
func (r *registerFdcanrxf0cType) SetF0s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf0cFieldF0sMask)|(uint32(value)<<RegisterFdcanrxf0cFieldF0sShift))
}

const (
	RegisterFdcanrxf0cFieldF0wmShift = 24
	RegisterFdcanrxf0cFieldF0wmMask  = 0xff000000
)

// GetF0wm FIFO 0 Watermark
func (r *registerFdcanrxf0cType) GetF0wm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf0cFieldF0wmMask) >> RegisterFdcanrxf0cFieldF0wmShift)
}

// SetF0wm FIFO 0 Watermark
func (r *registerFdcanrxf0cType) SetF0wm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf0cFieldF0wmMask)|(uint32(value)<<RegisterFdcanrxf0cFieldF0wmShift))
}

// registerFdcanrxf0sType FDCAN Rx FIFO 0 Status Register
type registerFdcanrxf0sType uint32

const (
	RegisterFdcanrxf0sFieldF0flShift = 0
	RegisterFdcanrxf0sFieldF0flMask  = 0x7f
)

// GetF0fl Rx FIFO 0 Fill Level
func (r *registerFdcanrxf0sType) GetF0fl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf0sFieldF0flMask) >> RegisterFdcanrxf0sFieldF0flShift)
}

// SetF0fl Rx FIFO 0 Fill Level
func (r *registerFdcanrxf0sType) SetF0fl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf0sFieldF0flMask)|(uint32(value)<<RegisterFdcanrxf0sFieldF0flShift))
}

const (
	RegisterFdcanrxf0sFieldF0gShift = 8
	RegisterFdcanrxf0sFieldF0gMask  = 0x3f00
)

// GetF0g Rx FIFO 0 Get Index
func (r *registerFdcanrxf0sType) GetF0g() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf0sFieldF0gMask) >> RegisterFdcanrxf0sFieldF0gShift)
}

// SetF0g Rx FIFO 0 Get Index
func (r *registerFdcanrxf0sType) SetF0g(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf0sFieldF0gMask)|(uint32(value)<<RegisterFdcanrxf0sFieldF0gShift))
}

const (
	RegisterFdcanrxf0sFieldF0pShift = 16
	RegisterFdcanrxf0sFieldF0pMask  = 0x3f0000
)

// GetF0p Rx FIFO 0 Put Index
func (r *registerFdcanrxf0sType) GetF0p() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf0sFieldF0pMask) >> RegisterFdcanrxf0sFieldF0pShift)
}

// SetF0p Rx FIFO 0 Put Index
func (r *registerFdcanrxf0sType) SetF0p(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf0sFieldF0pMask)|(uint32(value)<<RegisterFdcanrxf0sFieldF0pShift))
}

const (
	RegisterFdcanrxf0sFieldF0fShift = 24
	RegisterFdcanrxf0sFieldF0fMask  = 0x1000000
)

// GetF0f Rx FIFO 0 Full
func (r *registerFdcanrxf0sType) GetF0f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf0sFieldF0fMask) != 0
}

// SetF0f Rx FIFO 0 Full
func (r *registerFdcanrxf0sType) SetF0f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanrxf0sFieldF0fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf0sFieldF0fMask)
	}
}

const (
	RegisterFdcanrxf0sFieldRf0lShift = 25
	RegisterFdcanrxf0sFieldRf0lMask  = 0x2000000
)

// GetRf0l Rx FIFO 0 Message Lost
func (r *registerFdcanrxf0sType) GetRf0l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf0sFieldRf0lMask) != 0
}

// SetRf0l Rx FIFO 0 Message Lost
func (r *registerFdcanrxf0sType) SetRf0l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanrxf0sFieldRf0lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf0sFieldRf0lMask)
	}
}

// registerFdcanrxf0aType CAN Rx FIFO 0 Acknowledge Register
type registerFdcanrxf0aType uint32

const (
	RegisterFdcanrxf0aFieldFa01Shift = 0
	RegisterFdcanrxf0aFieldFa01Mask  = 0x3f
)

// GetFa01 Rx FIFO 0 Acknowledge Index
func (r *registerFdcanrxf0aType) GetFa01() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf0aFieldFa01Mask) >> RegisterFdcanrxf0aFieldFa01Shift)
}

// SetFa01 Rx FIFO 0 Acknowledge Index
func (r *registerFdcanrxf0aType) SetFa01(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf0aFieldFa01Mask)|(uint32(value)<<RegisterFdcanrxf0aFieldFa01Shift))
}

// registerFdcanrxbcType FDCAN Rx Buffer Configuration Register
type registerFdcanrxbcType uint32

const (
	RegisterFdcanrxbcFieldRbsaShift = 2
	RegisterFdcanrxbcFieldRbsaMask  = 0xfffc
)

// GetRbsa Rx Buffer Start Address
func (r *registerFdcanrxbcType) GetRbsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxbcFieldRbsaMask) >> RegisterFdcanrxbcFieldRbsaShift)
}

// SetRbsa Rx Buffer Start Address
func (r *registerFdcanrxbcType) SetRbsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxbcFieldRbsaMask)|(uint32(value)<<RegisterFdcanrxbcFieldRbsaShift))
}

// registerFdcanrxf1cType FDCAN Rx FIFO 1 Configuration Register
type registerFdcanrxf1cType uint32

const (
	RegisterFdcanrxf1cFieldF1saShift = 2
	RegisterFdcanrxf1cFieldF1saMask  = 0xfffc
)

// GetF1sa Rx FIFO 1 Start Address
func (r *registerFdcanrxf1cType) GetF1sa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf1cFieldF1saMask) >> RegisterFdcanrxf1cFieldF1saShift)
}

// SetF1sa Rx FIFO 1 Start Address
func (r *registerFdcanrxf1cType) SetF1sa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf1cFieldF1saMask)|(uint32(value)<<RegisterFdcanrxf1cFieldF1saShift))
}

const (
	RegisterFdcanrxf1cFieldF1sShift = 16
	RegisterFdcanrxf1cFieldF1sMask  = 0x7f0000
)

// GetF1s Rx FIFO 1 Size
func (r *registerFdcanrxf1cType) GetF1s() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf1cFieldF1sMask) >> RegisterFdcanrxf1cFieldF1sShift)
}

// SetF1s Rx FIFO 1 Size
func (r *registerFdcanrxf1cType) SetF1s(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf1cFieldF1sMask)|(uint32(value)<<RegisterFdcanrxf1cFieldF1sShift))
}

const (
	RegisterFdcanrxf1cFieldF1wmShift = 24
	RegisterFdcanrxf1cFieldF1wmMask  = 0x7f000000
)

// GetF1wm Rx FIFO 1 Watermark
func (r *registerFdcanrxf1cType) GetF1wm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf1cFieldF1wmMask) >> RegisterFdcanrxf1cFieldF1wmShift)
}

// SetF1wm Rx FIFO 1 Watermark
func (r *registerFdcanrxf1cType) SetF1wm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf1cFieldF1wmMask)|(uint32(value)<<RegisterFdcanrxf1cFieldF1wmShift))
}

// registerFdcanrxf1sType FDCAN Rx FIFO 1 Status Register
type registerFdcanrxf1sType uint32

const (
	RegisterFdcanrxf1sFieldF1flShift = 0
	RegisterFdcanrxf1sFieldF1flMask  = 0x7f
)

// GetF1fl Rx FIFO 1 Fill Level
func (r *registerFdcanrxf1sType) GetF1fl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf1sFieldF1flMask) >> RegisterFdcanrxf1sFieldF1flShift)
}

// SetF1fl Rx FIFO 1 Fill Level
func (r *registerFdcanrxf1sType) SetF1fl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf1sFieldF1flMask)|(uint32(value)<<RegisterFdcanrxf1sFieldF1flShift))
}

const (
	RegisterFdcanrxf1sFieldF1giShift = 8
	RegisterFdcanrxf1sFieldF1giMask  = 0x7f00
)

// GetF1gi Rx FIFO 1 Get Index
func (r *registerFdcanrxf1sType) GetF1gi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf1sFieldF1giMask) >> RegisterFdcanrxf1sFieldF1giShift)
}

// SetF1gi Rx FIFO 1 Get Index
func (r *registerFdcanrxf1sType) SetF1gi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf1sFieldF1giMask)|(uint32(value)<<RegisterFdcanrxf1sFieldF1giShift))
}

const (
	RegisterFdcanrxf1sFieldF1piShift = 16
	RegisterFdcanrxf1sFieldF1piMask  = 0x7f0000
)

// GetF1pi Rx FIFO 1 Put Index
func (r *registerFdcanrxf1sType) GetF1pi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf1sFieldF1piMask) >> RegisterFdcanrxf1sFieldF1piShift)
}

// SetF1pi Rx FIFO 1 Put Index
func (r *registerFdcanrxf1sType) SetF1pi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf1sFieldF1piMask)|(uint32(value)<<RegisterFdcanrxf1sFieldF1piShift))
}

const (
	RegisterFdcanrxf1sFieldF1fShift = 24
	RegisterFdcanrxf1sFieldF1fMask  = 0x1000000
)

// GetF1f Rx FIFO 1 Full
func (r *registerFdcanrxf1sType) GetF1f() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf1sFieldF1fMask) != 0
}

// SetF1f Rx FIFO 1 Full
func (r *registerFdcanrxf1sType) SetF1f(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanrxf1sFieldF1fMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf1sFieldF1fMask)
	}
}

const (
	RegisterFdcanrxf1sFieldRf1lShift = 25
	RegisterFdcanrxf1sFieldRf1lMask  = 0x2000000
)

// GetRf1l Rx FIFO 1 Message Lost
func (r *registerFdcanrxf1sType) GetRf1l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf1sFieldRf1lMask) != 0
}

// SetRf1l Rx FIFO 1 Message Lost
func (r *registerFdcanrxf1sType) SetRf1l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanrxf1sFieldRf1lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf1sFieldRf1lMask)
	}
}

const (
	RegisterFdcanrxf1sFieldDmsShift = 30
	RegisterFdcanrxf1sFieldDmsMask  = 0xc0000000
)

// GetDms Debug Message Status
func (r *registerFdcanrxf1sType) GetDms() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf1sFieldDmsMask) >> RegisterFdcanrxf1sFieldDmsShift)
}

// SetDms Debug Message Status
func (r *registerFdcanrxf1sType) SetDms(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf1sFieldDmsMask)|(uint32(value)<<RegisterFdcanrxf1sFieldDmsShift))
}

// registerFdcanrxf1aType FDCAN Rx FIFO 1 Acknowledge Register
type registerFdcanrxf1aType uint32

const (
	RegisterFdcanrxf1aFieldF1aiShift = 0
	RegisterFdcanrxf1aFieldF1aiMask  = 0x3f
)

// GetF1ai Rx FIFO 1 Acknowledge Index
func (r *registerFdcanrxf1aType) GetF1ai() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxf1aFieldF1aiMask) >> RegisterFdcanrxf1aFieldF1aiShift)
}

// SetF1ai Rx FIFO 1 Acknowledge Index
func (r *registerFdcanrxf1aType) SetF1ai(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxf1aFieldF1aiMask)|(uint32(value)<<RegisterFdcanrxf1aFieldF1aiShift))
}

// registerFdcanrxescType FDCAN Rx Buffer Element Size Configuration Register
type registerFdcanrxescType uint32

const (
	RegisterFdcanrxescFieldF0dsShift = 0
	RegisterFdcanrxescFieldF0dsMask  = 0x7
)

// GetF0ds Rx FIFO 1 Data Field Size:
func (r *registerFdcanrxescType) GetF0ds() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxescFieldF0dsMask) >> RegisterFdcanrxescFieldF0dsShift)
}

// SetF0ds Rx FIFO 1 Data Field Size:
func (r *registerFdcanrxescType) SetF0ds(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxescFieldF0dsMask)|(uint32(value)<<RegisterFdcanrxescFieldF0dsShift))
}

const (
	RegisterFdcanrxescFieldF1dsShift = 4
	RegisterFdcanrxescFieldF1dsMask  = 0x70
)

// GetF1ds Rx FIFO 0 Data Field Size:
func (r *registerFdcanrxescType) GetF1ds() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxescFieldF1dsMask) >> RegisterFdcanrxescFieldF1dsShift)
}

// SetF1ds Rx FIFO 0 Data Field Size:
func (r *registerFdcanrxescType) SetF1ds(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxescFieldF1dsMask)|(uint32(value)<<RegisterFdcanrxescFieldF1dsShift))
}

const (
	RegisterFdcanrxescFieldRbdsShift = 8
	RegisterFdcanrxescFieldRbdsMask  = 0x700
)

// GetRbds Rx Buffer Data Field Size:
func (r *registerFdcanrxescType) GetRbds() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanrxescFieldRbdsMask) >> RegisterFdcanrxescFieldRbdsShift)
}

// SetRbds Rx Buffer Data Field Size:
func (r *registerFdcanrxescType) SetRbds(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanrxescFieldRbdsMask)|(uint32(value)<<RegisterFdcanrxescFieldRbdsShift))
}

// registerFdcantxbcType FDCAN Tx Buffer Configuration Register
type registerFdcantxbcType uint32

const (
	RegisterFdcantxbcFieldTbsaShift = 2
	RegisterFdcantxbcFieldTbsaMask  = 0xfffc
)

// GetTbsa Tx Buffers Start Address
func (r *registerFdcantxbcType) GetTbsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxbcFieldTbsaMask) >> RegisterFdcantxbcFieldTbsaShift)
}

// SetTbsa Tx Buffers Start Address
func (r *registerFdcantxbcType) SetTbsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxbcFieldTbsaMask)|(uint32(value)<<RegisterFdcantxbcFieldTbsaShift))
}

const (
	RegisterFdcantxbcFieldNdtbShift = 16
	RegisterFdcantxbcFieldNdtbMask  = 0x3f0000
)

// GetNdtb Number of Dedicated Transmit Buffers
func (r *registerFdcantxbcType) GetNdtb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxbcFieldNdtbMask) >> RegisterFdcantxbcFieldNdtbShift)
}

// SetNdtb Number of Dedicated Transmit Buffers
func (r *registerFdcantxbcType) SetNdtb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxbcFieldNdtbMask)|(uint32(value)<<RegisterFdcantxbcFieldNdtbShift))
}

const (
	RegisterFdcantxbcFieldTfqsShift = 24
	RegisterFdcantxbcFieldTfqsMask  = 0x3f000000
)

// GetTfqs Transmit FIFO/Queue Size
func (r *registerFdcantxbcType) GetTfqs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxbcFieldTfqsMask) >> RegisterFdcantxbcFieldTfqsShift)
}

// SetTfqs Transmit FIFO/Queue Size
func (r *registerFdcantxbcType) SetTfqs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxbcFieldTfqsMask)|(uint32(value)<<RegisterFdcantxbcFieldTfqsShift))
}

const (
	RegisterFdcantxbcFieldTfqmShift = 30
	RegisterFdcantxbcFieldTfqmMask  = 0x40000000
)

// GetTfqm Tx FIFO/Queue Mode
func (r *registerFdcantxbcType) GetTfqm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxbcFieldTfqmMask) != 0
}

// SetTfqm Tx FIFO/Queue Mode
func (r *registerFdcantxbcType) SetTfqm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcantxbcFieldTfqmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxbcFieldTfqmMask)
	}
}

// registerFdcantxfqsType FDCAN Tx FIFO/Queue Status Register
type registerFdcantxfqsType uint32

const (
	RegisterFdcantxfqsFieldTfflShift = 0
	RegisterFdcantxfqsFieldTfflMask  = 0x3f
)

// GetTffl Tx FIFO Free Level
func (r *registerFdcantxfqsType) GetTffl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxfqsFieldTfflMask) >> RegisterFdcantxfqsFieldTfflShift)
}

// SetTffl Tx FIFO Free Level
func (r *registerFdcantxfqsType) SetTffl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxfqsFieldTfflMask)|(uint32(value)<<RegisterFdcantxfqsFieldTfflShift))
}

const (
	RegisterFdcantxfqsFieldTfgiShift = 8
	RegisterFdcantxfqsFieldTfgiMask  = 0x1f00
)

// GetTfgi TFGI
func (r *registerFdcantxfqsType) GetTfgi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxfqsFieldTfgiMask) >> RegisterFdcantxfqsFieldTfgiShift)
}

// SetTfgi TFGI
func (r *registerFdcantxfqsType) SetTfgi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxfqsFieldTfgiMask)|(uint32(value)<<RegisterFdcantxfqsFieldTfgiShift))
}

const (
	RegisterFdcantxfqsFieldTfqpiShift = 16
	RegisterFdcantxfqsFieldTfqpiMask  = 0x1f0000
)

// GetTfqpi Tx FIFO/Queue Put Index
func (r *registerFdcantxfqsType) GetTfqpi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxfqsFieldTfqpiMask) >> RegisterFdcantxfqsFieldTfqpiShift)
}

// SetTfqpi Tx FIFO/Queue Put Index
func (r *registerFdcantxfqsType) SetTfqpi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxfqsFieldTfqpiMask)|(uint32(value)<<RegisterFdcantxfqsFieldTfqpiShift))
}

const (
	RegisterFdcantxfqsFieldTfqfShift = 21
	RegisterFdcantxfqsFieldTfqfMask  = 0x200000
)

// GetTfqf Tx FIFO/Queue Full
func (r *registerFdcantxfqsType) GetTfqf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxfqsFieldTfqfMask) != 0
}

// SetTfqf Tx FIFO/Queue Full
func (r *registerFdcantxfqsType) SetTfqf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcantxfqsFieldTfqfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxfqsFieldTfqfMask)
	}
}

// registerFdcantxescType FDCAN Tx Buffer Element Size Configuration Register
type registerFdcantxescType uint32

const (
	RegisterFdcantxescFieldTbdsShift = 0
	RegisterFdcantxescFieldTbdsMask  = 0x7
)

// GetTbds Tx Buffer Data Field Size:
func (r *registerFdcantxescType) GetTbds() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxescFieldTbdsMask) >> RegisterFdcantxescFieldTbdsShift)
}

// SetTbds Tx Buffer Data Field Size:
func (r *registerFdcantxescType) SetTbds(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxescFieldTbdsMask)|(uint32(value)<<RegisterFdcantxescFieldTbdsShift))
}

// registerFdcantxbrpType FDCAN Tx Buffer Request Pending Register
type registerFdcantxbrpType uint32

const (
	RegisterFdcantxbrpFieldTrpShift = 0
	RegisterFdcantxbrpFieldTrpMask  = 0xffffffff
)

// GetTrp Transmission Request Pending
func (r *registerFdcantxbrpType) GetTrp() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxbrpFieldTrpMask) >> RegisterFdcantxbrpFieldTrpShift)
}

// SetTrp Transmission Request Pending
func (r *registerFdcantxbrpType) SetTrp(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxbrpFieldTrpMask)|(uint32(value)<<RegisterFdcantxbrpFieldTrpShift))
}

// registerFdcantxbarType FDCAN Tx Buffer Add Request Register
type registerFdcantxbarType uint32

const (
	RegisterFdcantxbarFieldArShift = 0
	RegisterFdcantxbarFieldArMask  = 0xffffffff
)

// GetAr Add Request
func (r *registerFdcantxbarType) GetAr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxbarFieldArMask) >> RegisterFdcantxbarFieldArShift)
}

// SetAr Add Request
func (r *registerFdcantxbarType) SetAr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxbarFieldArMask)|(uint32(value)<<RegisterFdcantxbarFieldArShift))
}

// registerFdcantxbcrType FDCAN Tx Buffer Cancellation Request Register
type registerFdcantxbcrType uint32

const (
	RegisterFdcantxbcrFieldCrShift = 0
	RegisterFdcantxbcrFieldCrMask  = 0xffffffff
)

// GetCr Cancellation Request
func (r *registerFdcantxbcrType) GetCr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxbcrFieldCrMask) >> RegisterFdcantxbcrFieldCrShift)
}

// SetCr Cancellation Request
func (r *registerFdcantxbcrType) SetCr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxbcrFieldCrMask)|(uint32(value)<<RegisterFdcantxbcrFieldCrShift))
}

// registerFdcantxbtoType FDCAN Tx Buffer Transmission Occurred Register
type registerFdcantxbtoType uint32

const (
	RegisterFdcantxbtoFieldToShift = 0
	RegisterFdcantxbtoFieldToMask  = 0xffffffff
)

// GetTo Transmission Occurred.
func (r *registerFdcantxbtoType) GetTo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxbtoFieldToMask) >> RegisterFdcantxbtoFieldToShift)
}

// SetTo Transmission Occurred.
func (r *registerFdcantxbtoType) SetTo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxbtoFieldToMask)|(uint32(value)<<RegisterFdcantxbtoFieldToShift))
}

// registerFdcantxbcfType FDCAN Tx Buffer Cancellation Finished Register
type registerFdcantxbcfType uint32

const (
	RegisterFdcantxbcfFieldCfShift = 0
	RegisterFdcantxbcfFieldCfMask  = 0xffffffff
)

// GetCf Cancellation Finished
func (r *registerFdcantxbcfType) GetCf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxbcfFieldCfMask) >> RegisterFdcantxbcfFieldCfShift)
}

// SetCf Cancellation Finished
func (r *registerFdcantxbcfType) SetCf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxbcfFieldCfMask)|(uint32(value)<<RegisterFdcantxbcfFieldCfShift))
}

// registerFdcantxbtieType FDCAN Tx Buffer Transmission Interrupt Enable Register
type registerFdcantxbtieType uint32

const (
	RegisterFdcantxbtieFieldTieShift = 0
	RegisterFdcantxbtieFieldTieMask  = 0xffffffff
)

// GetTie Transmission Interrupt Enable
func (r *registerFdcantxbtieType) GetTie() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxbtieFieldTieMask) >> RegisterFdcantxbtieFieldTieShift)
}

// SetTie Transmission Interrupt Enable
func (r *registerFdcantxbtieType) SetTie(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxbtieFieldTieMask)|(uint32(value)<<RegisterFdcantxbtieFieldTieShift))
}

// registerFdcantxbcieType FDCAN Tx Buffer Cancellation Finished Interrupt Enable Register
type registerFdcantxbcieType uint32

const (
	RegisterFdcantxbcieFieldCfShift = 0
	RegisterFdcantxbcieFieldCfMask  = 0xffffffff
)

// GetCf Cancellation Finished Interrupt Enable
func (r *registerFdcantxbcieType) GetCf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxbcieFieldCfMask) >> RegisterFdcantxbcieFieldCfShift)
}

// SetCf Cancellation Finished Interrupt Enable
func (r *registerFdcantxbcieType) SetCf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxbcieFieldCfMask)|(uint32(value)<<RegisterFdcantxbcieFieldCfShift))
}

// registerFdcantxefcType FDCAN Tx Event FIFO Configuration Register
type registerFdcantxefcType uint32

const (
	RegisterFdcantxefcFieldEfsaShift = 2
	RegisterFdcantxefcFieldEfsaMask  = 0xfffc
)

// GetEfsa Event FIFO Start Address
func (r *registerFdcantxefcType) GetEfsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxefcFieldEfsaMask) >> RegisterFdcantxefcFieldEfsaShift)
}

// SetEfsa Event FIFO Start Address
func (r *registerFdcantxefcType) SetEfsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxefcFieldEfsaMask)|(uint32(value)<<RegisterFdcantxefcFieldEfsaShift))
}

const (
	RegisterFdcantxefcFieldEfsShift = 16
	RegisterFdcantxefcFieldEfsMask  = 0x3f0000
)

// GetEfs Event FIFO Size
func (r *registerFdcantxefcType) GetEfs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxefcFieldEfsMask) >> RegisterFdcantxefcFieldEfsShift)
}

// SetEfs Event FIFO Size
func (r *registerFdcantxefcType) SetEfs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxefcFieldEfsMask)|(uint32(value)<<RegisterFdcantxefcFieldEfsShift))
}

const (
	RegisterFdcantxefcFieldEfwmShift = 24
	RegisterFdcantxefcFieldEfwmMask  = 0x3f000000
)

// GetEfwm Event FIFO Watermark
func (r *registerFdcantxefcType) GetEfwm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxefcFieldEfwmMask) >> RegisterFdcantxefcFieldEfwmShift)
}

// SetEfwm Event FIFO Watermark
func (r *registerFdcantxefcType) SetEfwm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxefcFieldEfwmMask)|(uint32(value)<<RegisterFdcantxefcFieldEfwmShift))
}

// registerFdcantxefsType FDCAN Tx Event FIFO Status Register
type registerFdcantxefsType uint32

const (
	RegisterFdcantxefsFieldEfflShift = 0
	RegisterFdcantxefsFieldEfflMask  = 0x3f
)

// GetEffl Event FIFO Fill Level
func (r *registerFdcantxefsType) GetEffl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxefsFieldEfflMask) >> RegisterFdcantxefsFieldEfflShift)
}

// SetEffl Event FIFO Fill Level
func (r *registerFdcantxefsType) SetEffl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxefsFieldEfflMask)|(uint32(value)<<RegisterFdcantxefsFieldEfflShift))
}

const (
	RegisterFdcantxefsFieldEfgiShift = 8
	RegisterFdcantxefsFieldEfgiMask  = 0x1f00
)

// GetEfgi Event FIFO Get Index.
func (r *registerFdcantxefsType) GetEfgi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxefsFieldEfgiMask) >> RegisterFdcantxefsFieldEfgiShift)
}

// SetEfgi Event FIFO Get Index.
func (r *registerFdcantxefsType) SetEfgi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxefsFieldEfgiMask)|(uint32(value)<<RegisterFdcantxefsFieldEfgiShift))
}

const (
	RegisterFdcantxefsFieldEfpiShift = 16
	RegisterFdcantxefsFieldEfpiMask  = 0x1f0000
)

// GetEfpi Event FIFO put index.
func (r *registerFdcantxefsType) GetEfpi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxefsFieldEfpiMask) >> RegisterFdcantxefsFieldEfpiShift)
}

// SetEfpi Event FIFO put index.
func (r *registerFdcantxefsType) SetEfpi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxefsFieldEfpiMask)|(uint32(value)<<RegisterFdcantxefsFieldEfpiShift))
}

const (
	RegisterFdcantxefsFieldEffShift = 24
	RegisterFdcantxefsFieldEffMask  = 0x1000000
)

// GetEff Event FIFO Full.
func (r *registerFdcantxefsType) GetEff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxefsFieldEffMask) != 0
}

// SetEff Event FIFO Full.
func (r *registerFdcantxefsType) SetEff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcantxefsFieldEffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxefsFieldEffMask)
	}
}

const (
	RegisterFdcantxefsFieldTeflShift = 25
	RegisterFdcantxefsFieldTeflMask  = 0x2000000
)

// GetTefl Tx Event FIFO Element Lost.
func (r *registerFdcantxefsType) GetTefl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxefsFieldTeflMask) != 0
}

// SetTefl Tx Event FIFO Element Lost.
func (r *registerFdcantxefsType) SetTefl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcantxefsFieldTeflMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxefsFieldTeflMask)
	}
}

// registerFdcantxefaType FDCAN Tx Event FIFO Acknowledge Register
type registerFdcantxefaType uint32

const (
	RegisterFdcantxefaFieldEfaiShift = 0
	RegisterFdcantxefaFieldEfaiMask  = 0x1f
)

// GetEfai Event FIFO Acknowledge Index
func (r *registerFdcantxefaType) GetEfai() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantxefaFieldEfaiMask) >> RegisterFdcantxefaFieldEfaiShift)
}

// SetEfai Event FIFO Acknowledge Index
func (r *registerFdcantxefaType) SetEfai(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantxefaFieldEfaiMask)|(uint32(value)<<RegisterFdcantxefaFieldEfaiShift))
}

// registerFdcantttmcType FDCAN TT Trigger Memory Configuration Register
type registerFdcantttmcType uint32

const (
	RegisterFdcantttmcFieldTmsaShift = 2
	RegisterFdcantttmcFieldTmsaMask  = 0xfffc
)

// GetTmsa Trigger Memory Start Address
func (r *registerFdcantttmcType) GetTmsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantttmcFieldTmsaMask) >> RegisterFdcantttmcFieldTmsaShift)
}

// SetTmsa Trigger Memory Start Address
func (r *registerFdcantttmcType) SetTmsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantttmcFieldTmsaMask)|(uint32(value)<<RegisterFdcantttmcFieldTmsaShift))
}

const (
	RegisterFdcantttmcFieldTmeShift = 16
	RegisterFdcantttmcFieldTmeMask  = 0x7f0000
)

// GetTme Trigger Memory Elements
func (r *registerFdcantttmcType) GetTme() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantttmcFieldTmeMask) >> RegisterFdcantttmcFieldTmeShift)
}

// SetTme Trigger Memory Elements
func (r *registerFdcantttmcType) SetTme(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantttmcFieldTmeMask)|(uint32(value)<<RegisterFdcantttmcFieldTmeShift))
}

// registerFdcanttrmcType FDCAN TT Reference Message Configuration Register
type registerFdcanttrmcType uint32

const (
	RegisterFdcanttrmcFieldRidShift = 0
	RegisterFdcanttrmcFieldRidMask  = 0x1fffffff
)

// GetRid Reference Identifier.
func (r *registerFdcanttrmcType) GetRid() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttrmcFieldRidMask) >> RegisterFdcanttrmcFieldRidShift)
}

// SetRid Reference Identifier.
func (r *registerFdcanttrmcType) SetRid(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttrmcFieldRidMask)|(uint32(value)<<RegisterFdcanttrmcFieldRidShift))
}

const (
	RegisterFdcanttrmcFieldXtdShift = 30
	RegisterFdcanttrmcFieldXtdMask  = 0x40000000
)

// GetXtd Extended Identifier
func (r *registerFdcanttrmcType) GetXtd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttrmcFieldXtdMask) != 0
}

// SetXtd Extended Identifier
func (r *registerFdcanttrmcType) SetXtd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttrmcFieldXtdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttrmcFieldXtdMask)
	}
}

const (
	RegisterFdcanttrmcFieldRmpsShift = 31
	RegisterFdcanttrmcFieldRmpsMask  = 0x80000000
)

// GetRmps Reference Message Payload Select
func (r *registerFdcanttrmcType) GetRmps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttrmcFieldRmpsMask) != 0
}

// SetRmps Reference Message Payload Select
func (r *registerFdcanttrmcType) SetRmps(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttrmcFieldRmpsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttrmcFieldRmpsMask)
	}
}

// registerFdcanttocfType FDCAN TT Operation Configuration Register
type registerFdcanttocfType uint32

const (
	RegisterFdcanttocfFieldOmShift = 0
	RegisterFdcanttocfFieldOmMask  = 0x3
)

// GetOm Operation Mode
func (r *registerFdcanttocfType) GetOm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocfFieldOmMask) >> RegisterFdcanttocfFieldOmShift)
}

// SetOm Operation Mode
func (r *registerFdcanttocfType) SetOm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocfFieldOmMask)|(uint32(value)<<RegisterFdcanttocfFieldOmShift))
}

const (
	RegisterFdcanttocfFieldGenShift = 3
	RegisterFdcanttocfFieldGenMask  = 0x8
)

// GetGen Gap Enable
func (r *registerFdcanttocfType) GetGen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocfFieldGenMask) != 0
}

// SetGen Gap Enable
func (r *registerFdcanttocfType) SetGen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocfFieldGenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocfFieldGenMask)
	}
}

const (
	RegisterFdcanttocfFieldTmShift = 4
	RegisterFdcanttocfFieldTmMask  = 0x10
)

// GetTm Time Master
func (r *registerFdcanttocfType) GetTm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocfFieldTmMask) != 0
}

// SetTm Time Master
func (r *registerFdcanttocfType) SetTm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocfFieldTmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocfFieldTmMask)
	}
}

const (
	RegisterFdcanttocfFieldLdsdlShift = 5
	RegisterFdcanttocfFieldLdsdlMask  = 0xe0
)

// GetLdsdl LD of Synchronization Deviation Limit
func (r *registerFdcanttocfType) GetLdsdl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocfFieldLdsdlMask) >> RegisterFdcanttocfFieldLdsdlShift)
}

// SetLdsdl LD of Synchronization Deviation Limit
func (r *registerFdcanttocfType) SetLdsdl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocfFieldLdsdlMask)|(uint32(value)<<RegisterFdcanttocfFieldLdsdlShift))
}

const (
	RegisterFdcanttocfFieldIrtoShift = 8
	RegisterFdcanttocfFieldIrtoMask  = 0x7f00
)

// GetIrto Initial Reference Trigger Offset
func (r *registerFdcanttocfType) GetIrto() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocfFieldIrtoMask) >> RegisterFdcanttocfFieldIrtoShift)
}

// SetIrto Initial Reference Trigger Offset
func (r *registerFdcanttocfType) SetIrto(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocfFieldIrtoMask)|(uint32(value)<<RegisterFdcanttocfFieldIrtoShift))
}

const (
	RegisterFdcanttocfFieldEecsShift = 15
	RegisterFdcanttocfFieldEecsMask  = 0x8000
)

// GetEecs Enable External Clock Synchronization
func (r *registerFdcanttocfType) GetEecs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocfFieldEecsMask) != 0
}

// SetEecs Enable External Clock Synchronization
func (r *registerFdcanttocfType) SetEecs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocfFieldEecsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocfFieldEecsMask)
	}
}

const (
	RegisterFdcanttocfFieldAwlShift = 16
	RegisterFdcanttocfFieldAwlMask  = 0xff0000
)

// GetAwl Application Watchdog Limit
func (r *registerFdcanttocfType) GetAwl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocfFieldAwlMask) >> RegisterFdcanttocfFieldAwlShift)
}

// SetAwl Application Watchdog Limit
func (r *registerFdcanttocfType) SetAwl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocfFieldAwlMask)|(uint32(value)<<RegisterFdcanttocfFieldAwlShift))
}

const (
	RegisterFdcanttocfFieldEgtfShift = 24
	RegisterFdcanttocfFieldEgtfMask  = 0x1000000
)

// GetEgtf Enable Global Time Filtering
func (r *registerFdcanttocfType) GetEgtf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocfFieldEgtfMask) != 0
}

// SetEgtf Enable Global Time Filtering
func (r *registerFdcanttocfType) SetEgtf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocfFieldEgtfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocfFieldEgtfMask)
	}
}

const (
	RegisterFdcanttocfFieldEccShift = 25
	RegisterFdcanttocfFieldEccMask  = 0x2000000
)

// GetEcc Enable Clock Calibration
func (r *registerFdcanttocfType) GetEcc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocfFieldEccMask) != 0
}

// SetEcc Enable Clock Calibration
func (r *registerFdcanttocfType) SetEcc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocfFieldEccMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocfFieldEccMask)
	}
}

const (
	RegisterFdcanttocfFieldEvtpShift = 26
	RegisterFdcanttocfFieldEvtpMask  = 0x4000000
)

// GetEvtp Event Trigger Polarity
func (r *registerFdcanttocfType) GetEvtp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocfFieldEvtpMask) != 0
}

// SetEvtp Event Trigger Polarity
func (r *registerFdcanttocfType) SetEvtp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocfFieldEvtpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocfFieldEvtpMask)
	}
}

// registerFdcanttmlmType FDCAN TT Matrix Limits Register
type registerFdcanttmlmType uint32

const (
	RegisterFdcanttmlmFieldCcmShift = 0
	RegisterFdcanttmlmFieldCcmMask  = 0x3f
)

// GetCcm Cycle Count Max
func (r *registerFdcanttmlmType) GetCcm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttmlmFieldCcmMask) >> RegisterFdcanttmlmFieldCcmShift)
}

// SetCcm Cycle Count Max
func (r *registerFdcanttmlmType) SetCcm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttmlmFieldCcmMask)|(uint32(value)<<RegisterFdcanttmlmFieldCcmShift))
}

const (
	RegisterFdcanttmlmFieldCssShift = 6
	RegisterFdcanttmlmFieldCssMask  = 0xc0
)

// GetCss Cycle Start Synchronization
func (r *registerFdcanttmlmType) GetCss() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttmlmFieldCssMask) >> RegisterFdcanttmlmFieldCssShift)
}

// SetCss Cycle Start Synchronization
func (r *registerFdcanttmlmType) SetCss(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttmlmFieldCssMask)|(uint32(value)<<RegisterFdcanttmlmFieldCssShift))
}

const (
	RegisterFdcanttmlmFieldTxewShift = 8
	RegisterFdcanttmlmFieldTxewMask  = 0xf00
)

// GetTxew Tx Enable Window
func (r *registerFdcanttmlmType) GetTxew() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttmlmFieldTxewMask) >> RegisterFdcanttmlmFieldTxewShift)
}

// SetTxew Tx Enable Window
func (r *registerFdcanttmlmType) SetTxew(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttmlmFieldTxewMask)|(uint32(value)<<RegisterFdcanttmlmFieldTxewShift))
}

const (
	RegisterFdcanttmlmFieldEnttShift = 16
	RegisterFdcanttmlmFieldEnttMask  = 0xfff0000
)

// GetEntt Expected Number of Tx Triggers
func (r *registerFdcanttmlmType) GetEntt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttmlmFieldEnttMask) >> RegisterFdcanttmlmFieldEnttShift)
}

// SetEntt Expected Number of Tx Triggers
func (r *registerFdcanttmlmType) SetEntt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttmlmFieldEnttMask)|(uint32(value)<<RegisterFdcanttmlmFieldEnttShift))
}

// registerFdcanturcfType FDCAN TUR Configuration Register
type registerFdcanturcfType uint32

const (
	RegisterFdcanturcfFieldNclShift = 0
	RegisterFdcanturcfFieldNclMask  = 0xffff
)

// GetNcl Numerator Configuration Low.
func (r *registerFdcanturcfType) GetNcl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanturcfFieldNclMask) >> RegisterFdcanturcfFieldNclShift)
}

// SetNcl Numerator Configuration Low.
func (r *registerFdcanturcfType) SetNcl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanturcfFieldNclMask)|(uint32(value)<<RegisterFdcanturcfFieldNclShift))
}

const (
	RegisterFdcanturcfFieldDcShift = 16
	RegisterFdcanturcfFieldDcMask  = 0x3fff0000
)

// GetDc Denominator Configuration.
func (r *registerFdcanturcfType) GetDc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanturcfFieldDcMask) >> RegisterFdcanturcfFieldDcShift)
}

// SetDc Denominator Configuration.
func (r *registerFdcanturcfType) SetDc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanturcfFieldDcMask)|(uint32(value)<<RegisterFdcanturcfFieldDcShift))
}

const (
	RegisterFdcanturcfFieldEltShift = 31
	RegisterFdcanturcfFieldEltMask  = 0x80000000
)

// GetElt Enable Local Time
func (r *registerFdcanturcfType) GetElt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanturcfFieldEltMask) != 0
}

// SetElt Enable Local Time
func (r *registerFdcanturcfType) SetElt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanturcfFieldEltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanturcfFieldEltMask)
	}
}

// registerFdcanttocnType FDCAN TT Operation Control Register
type registerFdcanttocnType uint32

const (
	RegisterFdcanttocnFieldSgtShift = 0
	RegisterFdcanttocnFieldSgtMask  = 0x1
)

// GetSgt Set Global time
func (r *registerFdcanttocnType) GetSgt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldSgtMask) != 0
}

// SetSgt Set Global time
func (r *registerFdcanttocnType) SetSgt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocnFieldSgtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldSgtMask)
	}
}

const (
	RegisterFdcanttocnFieldEcsShift = 1
	RegisterFdcanttocnFieldEcsMask  = 0x2
)

// GetEcs External Clock Synchronization
func (r *registerFdcanttocnType) GetEcs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldEcsMask) != 0
}

// SetEcs External Clock Synchronization
func (r *registerFdcanttocnType) SetEcs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocnFieldEcsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldEcsMask)
	}
}

const (
	RegisterFdcanttocnFieldSwpShift = 2
	RegisterFdcanttocnFieldSwpMask  = 0x4
)

// GetSwp Stop Watch Polarity
func (r *registerFdcanttocnType) GetSwp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldSwpMask) != 0
}

// SetSwp Stop Watch Polarity
func (r *registerFdcanttocnType) SetSwp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocnFieldSwpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldSwpMask)
	}
}

const (
	RegisterFdcanttocnFieldSwsShift = 3
	RegisterFdcanttocnFieldSwsMask  = 0x18
)

// GetSws Stop Watch Source.
func (r *registerFdcanttocnType) GetSws() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldSwsMask) >> RegisterFdcanttocnFieldSwsShift)
}

// SetSws Stop Watch Source.
func (r *registerFdcanttocnType) SetSws(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldSwsMask)|(uint32(value)<<RegisterFdcanttocnFieldSwsShift))
}

const (
	RegisterFdcanttocnFieldRtieShift = 5
	RegisterFdcanttocnFieldRtieMask  = 0x20
)

// GetRtie Register Time Mark Interrupt Pulse Enable
func (r *registerFdcanttocnType) GetRtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldRtieMask) != 0
}

// SetRtie Register Time Mark Interrupt Pulse Enable
func (r *registerFdcanttocnType) SetRtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocnFieldRtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldRtieMask)
	}
}

const (
	RegisterFdcanttocnFieldTmcShift = 6
	RegisterFdcanttocnFieldTmcMask  = 0xc0
)

// GetTmc Register Time Mark Compare
func (r *registerFdcanttocnType) GetTmc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldTmcMask) >> RegisterFdcanttocnFieldTmcShift)
}

// SetTmc Register Time Mark Compare
func (r *registerFdcanttocnType) SetTmc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldTmcMask)|(uint32(value)<<RegisterFdcanttocnFieldTmcShift))
}

const (
	RegisterFdcanttocnFieldTtieShift = 8
	RegisterFdcanttocnFieldTtieMask  = 0x100
)

// GetTtie Trigger Time Mark Interrupt Pulse Enable
func (r *registerFdcanttocnType) GetTtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldTtieMask) != 0
}

// SetTtie Trigger Time Mark Interrupt Pulse Enable
func (r *registerFdcanttocnType) SetTtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocnFieldTtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldTtieMask)
	}
}

const (
	RegisterFdcanttocnFieldGcsShift = 9
	RegisterFdcanttocnFieldGcsMask  = 0x200
)

// GetGcs Gap Control Select
func (r *registerFdcanttocnType) GetGcs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldGcsMask) != 0
}

// SetGcs Gap Control Select
func (r *registerFdcanttocnType) SetGcs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocnFieldGcsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldGcsMask)
	}
}

const (
	RegisterFdcanttocnFieldFgpShift = 10
	RegisterFdcanttocnFieldFgpMask  = 0x400
)

// GetFgp Finish Gap.
func (r *registerFdcanttocnType) GetFgp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldFgpMask) != 0
}

// SetFgp Finish Gap.
func (r *registerFdcanttocnType) SetFgp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocnFieldFgpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldFgpMask)
	}
}

const (
	RegisterFdcanttocnFieldTmgShift = 11
	RegisterFdcanttocnFieldTmgMask  = 0x800
)

// GetTmg Time Mark Gap
func (r *registerFdcanttocnType) GetTmg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldTmgMask) != 0
}

// SetTmg Time Mark Gap
func (r *registerFdcanttocnType) SetTmg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocnFieldTmgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldTmgMask)
	}
}

const (
	RegisterFdcanttocnFieldNigShift = 12
	RegisterFdcanttocnFieldNigMask  = 0x1000
)

// GetNig Next is Gap
func (r *registerFdcanttocnType) GetNig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldNigMask) != 0
}

// SetNig Next is Gap
func (r *registerFdcanttocnType) SetNig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocnFieldNigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldNigMask)
	}
}

const (
	RegisterFdcanttocnFieldEscnShift = 13
	RegisterFdcanttocnFieldEscnMask  = 0x2000
)

// GetEscn External Synchronization Control
func (r *registerFdcanttocnType) GetEscn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldEscnMask) != 0
}

// SetEscn External Synchronization Control
func (r *registerFdcanttocnType) SetEscn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocnFieldEscnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldEscnMask)
	}
}

const (
	RegisterFdcanttocnFieldLckcShift = 15
	RegisterFdcanttocnFieldLckcMask  = 0x8000
)

// GetLckc TT Operation Control Register Locked
func (r *registerFdcanttocnType) GetLckc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttocnFieldLckcMask) != 0
}

// SetLckc TT Operation Control Register Locked
func (r *registerFdcanttocnType) SetLckc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttocnFieldLckcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttocnFieldLckcMask)
	}
}

// registerCanttgtpType FDCAN TT Global Time Preset Register
type registerCanttgtpType uint32

const (
	RegisterCanttgtpFieldNclShift = 0
	RegisterCanttgtpFieldNclMask  = 0xffff
)

// GetNcl Time Preset
func (r *registerCanttgtpType) GetNcl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCanttgtpFieldNclMask) >> RegisterCanttgtpFieldNclShift)
}

// SetNcl Time Preset
func (r *registerCanttgtpType) SetNcl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCanttgtpFieldNclMask)|(uint32(value)<<RegisterCanttgtpFieldNclShift))
}

const (
	RegisterCanttgtpFieldCtpShift = 16
	RegisterCanttgtpFieldCtpMask  = 0xffff0000
)

// GetCtp Cycle Time Target Phase
func (r *registerCanttgtpType) GetCtp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCanttgtpFieldCtpMask) >> RegisterCanttgtpFieldCtpShift)
}

// SetCtp Cycle Time Target Phase
func (r *registerCanttgtpType) SetCtp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCanttgtpFieldCtpMask)|(uint32(value)<<RegisterCanttgtpFieldCtpShift))
}

// registerFdcantttmkType FDCAN TT Time Mark Register
type registerFdcantttmkType uint32

const (
	RegisterFdcantttmkFieldTmShift = 0
	RegisterFdcantttmkFieldTmMask  = 0xffff
)

// GetTm Time Mark
func (r *registerFdcantttmkType) GetTm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantttmkFieldTmMask) >> RegisterFdcantttmkFieldTmShift)
}

// SetTm Time Mark
func (r *registerFdcantttmkType) SetTm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantttmkFieldTmMask)|(uint32(value)<<RegisterFdcantttmkFieldTmShift))
}

const (
	RegisterFdcantttmkFieldTiccShift = 16
	RegisterFdcantttmkFieldTiccMask  = 0x7f0000
)

// GetTicc Time Mark Cycle Code
func (r *registerFdcantttmkType) GetTicc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantttmkFieldTiccMask) >> RegisterFdcantttmkFieldTiccShift)
}

// SetTicc Time Mark Cycle Code
func (r *registerFdcantttmkType) SetTicc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantttmkFieldTiccMask)|(uint32(value)<<RegisterFdcantttmkFieldTiccShift))
}

const (
	RegisterFdcantttmkFieldLckmShift = 31
	RegisterFdcantttmkFieldLckmMask  = 0x80000000
)

// GetLckm TT Time Mark Register Locked
func (r *registerFdcantttmkType) GetLckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcantttmkFieldLckmMask) != 0
}

// SetLckm TT Time Mark Register Locked
func (r *registerFdcantttmkType) SetLckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcantttmkFieldLckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcantttmkFieldLckmMask)
	}
}

// registerFdcanttirType FDCAN TT Interrupt Register
type registerFdcanttirType uint32

const (
	RegisterFdcanttirFieldSbcShift = 0
	RegisterFdcanttirFieldSbcMask  = 0x1
)

// GetSbc Start of Basic Cycle
func (r *registerFdcanttirType) GetSbc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldSbcMask) != 0
}

// SetSbc Start of Basic Cycle
func (r *registerFdcanttirType) SetSbc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldSbcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldSbcMask)
	}
}

const (
	RegisterFdcanttirFieldSmcShift = 1
	RegisterFdcanttirFieldSmcMask  = 0x2
)

// GetSmc Start of Matrix Cycle
func (r *registerFdcanttirType) GetSmc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldSmcMask) != 0
}

// SetSmc Start of Matrix Cycle
func (r *registerFdcanttirType) SetSmc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldSmcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldSmcMask)
	}
}

const (
	RegisterFdcanttirFieldCsmShift = 2
	RegisterFdcanttirFieldCsmMask  = 0x4
)

// GetCsm Change of Synchronization Mode
func (r *registerFdcanttirType) GetCsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldCsmMask) != 0
}

// SetCsm Change of Synchronization Mode
func (r *registerFdcanttirType) SetCsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldCsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldCsmMask)
	}
}

const (
	RegisterFdcanttirFieldSogShift = 3
	RegisterFdcanttirFieldSogMask  = 0x8
)

// GetSog Start of Gap
func (r *registerFdcanttirType) GetSog() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldSogMask) != 0
}

// SetSog Start of Gap
func (r *registerFdcanttirType) SetSog(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldSogMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldSogMask)
	}
}

const (
	RegisterFdcanttirFieldRtmiShift = 4
	RegisterFdcanttirFieldRtmiMask  = 0x10
)

// GetRtmi Register Time Mark Interrupt.
func (r *registerFdcanttirType) GetRtmi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldRtmiMask) != 0
}

// SetRtmi Register Time Mark Interrupt.
func (r *registerFdcanttirType) SetRtmi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldRtmiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldRtmiMask)
	}
}

const (
	RegisterFdcanttirFieldTtmiShift = 5
	RegisterFdcanttirFieldTtmiMask  = 0x20
)

// GetTtmi Trigger Time Mark Event Internal
func (r *registerFdcanttirType) GetTtmi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldTtmiMask) != 0
}

// SetTtmi Trigger Time Mark Event Internal
func (r *registerFdcanttirType) SetTtmi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldTtmiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldTtmiMask)
	}
}

const (
	RegisterFdcanttirFieldSweShift = 6
	RegisterFdcanttirFieldSweMask  = 0x40
)

// GetSwe Stop Watch Event
func (r *registerFdcanttirType) GetSwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldSweMask) != 0
}

// SetSwe Stop Watch Event
func (r *registerFdcanttirType) SetSwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldSweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldSweMask)
	}
}

const (
	RegisterFdcanttirFieldGtwShift = 7
	RegisterFdcanttirFieldGtwMask  = 0x80
)

// GetGtw Global Time Wrap
func (r *registerFdcanttirType) GetGtw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldGtwMask) != 0
}

// SetGtw Global Time Wrap
func (r *registerFdcanttirType) SetGtw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldGtwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldGtwMask)
	}
}

const (
	RegisterFdcanttirFieldGtdShift = 8
	RegisterFdcanttirFieldGtdMask  = 0x100
)

// GetGtd Global Time Discontinuity
func (r *registerFdcanttirType) GetGtd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldGtdMask) != 0
}

// SetGtd Global Time Discontinuity
func (r *registerFdcanttirType) SetGtd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldGtdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldGtdMask)
	}
}

const (
	RegisterFdcanttirFieldGteShift = 9
	RegisterFdcanttirFieldGteMask  = 0x200
)

// GetGte Global Time Error
func (r *registerFdcanttirType) GetGte() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldGteMask) != 0
}

// SetGte Global Time Error
func (r *registerFdcanttirType) SetGte(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldGteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldGteMask)
	}
}

const (
	RegisterFdcanttirFieldTxuShift = 10
	RegisterFdcanttirFieldTxuMask  = 0x400
)

// GetTxu Tx Count Underflow
func (r *registerFdcanttirType) GetTxu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldTxuMask) != 0
}

// SetTxu Tx Count Underflow
func (r *registerFdcanttirType) SetTxu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldTxuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldTxuMask)
	}
}

const (
	RegisterFdcanttirFieldTxoShift = 11
	RegisterFdcanttirFieldTxoMask  = 0x800
)

// GetTxo Tx Count Overflow
func (r *registerFdcanttirType) GetTxo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldTxoMask) != 0
}

// SetTxo Tx Count Overflow
func (r *registerFdcanttirType) SetTxo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldTxoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldTxoMask)
	}
}

const (
	RegisterFdcanttirFieldSe1Shift = 12
	RegisterFdcanttirFieldSe1Mask  = 0x1000
)

// GetSe1 Scheduling Error 1
func (r *registerFdcanttirType) GetSe1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldSe1Mask) != 0
}

// SetSe1 Scheduling Error 1
func (r *registerFdcanttirType) SetSe1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldSe1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldSe1Mask)
	}
}

const (
	RegisterFdcanttirFieldSe2Shift = 13
	RegisterFdcanttirFieldSe2Mask  = 0x2000
)

// GetSe2 Scheduling Error 2
func (r *registerFdcanttirType) GetSe2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldSe2Mask) != 0
}

// SetSe2 Scheduling Error 2
func (r *registerFdcanttirType) SetSe2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldSe2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldSe2Mask)
	}
}

const (
	RegisterFdcanttirFieldElcShift = 14
	RegisterFdcanttirFieldElcMask  = 0x4000
)

// GetElc Error Level Changed.
func (r *registerFdcanttirType) GetElc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldElcMask) != 0
}

// SetElc Error Level Changed.
func (r *registerFdcanttirType) SetElc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldElcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldElcMask)
	}
}

const (
	RegisterFdcanttirFieldIwtgShift = 15
	RegisterFdcanttirFieldIwtgMask  = 0x8000
)

// GetIwtg Initialization Watch Trigger
func (r *registerFdcanttirType) GetIwtg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldIwtgMask) != 0
}

// SetIwtg Initialization Watch Trigger
func (r *registerFdcanttirType) SetIwtg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldIwtgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldIwtgMask)
	}
}

const (
	RegisterFdcanttirFieldWtShift = 16
	RegisterFdcanttirFieldWtMask  = 0x10000
)

// GetWt Watch Trigger
func (r *registerFdcanttirType) GetWt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldWtMask) != 0
}

// SetWt Watch Trigger
func (r *registerFdcanttirType) SetWt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldWtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldWtMask)
	}
}

const (
	RegisterFdcanttirFieldAwShift = 17
	RegisterFdcanttirFieldAwMask  = 0x20000
)

// GetAw Application Watchdog
func (r *registerFdcanttirType) GetAw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldAwMask) != 0
}

// SetAw Application Watchdog
func (r *registerFdcanttirType) SetAw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldAwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldAwMask)
	}
}

const (
	RegisterFdcanttirFieldCerShift = 18
	RegisterFdcanttirFieldCerMask  = 0x40000
)

// GetCer Configuration Error
func (r *registerFdcanttirType) GetCer() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttirFieldCerMask) != 0
}

// SetCer Configuration Error
func (r *registerFdcanttirType) SetCer(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttirFieldCerMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttirFieldCerMask)
	}
}

// registerFdcanttieType FDCAN TT Interrupt Enable Register
type registerFdcanttieType uint32

const (
	RegisterFdcanttieFieldSbceShift = 0
	RegisterFdcanttieFieldSbceMask  = 0x1
)

// GetSbce Start of Basic Cycle Interrupt Enable
func (r *registerFdcanttieType) GetSbce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldSbceMask) != 0
}

// SetSbce Start of Basic Cycle Interrupt Enable
func (r *registerFdcanttieType) SetSbce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldSbceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldSbceMask)
	}
}

const (
	RegisterFdcanttieFieldSmceShift = 1
	RegisterFdcanttieFieldSmceMask  = 0x2
)

// GetSmce Start of Matrix Cycle Interrupt Enable
func (r *registerFdcanttieType) GetSmce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldSmceMask) != 0
}

// SetSmce Start of Matrix Cycle Interrupt Enable
func (r *registerFdcanttieType) SetSmce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldSmceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldSmceMask)
	}
}

const (
	RegisterFdcanttieFieldCsmeShift = 2
	RegisterFdcanttieFieldCsmeMask  = 0x4
)

// GetCsme Change of Synchronization Mode Interrupt Enable
func (r *registerFdcanttieType) GetCsme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldCsmeMask) != 0
}

// SetCsme Change of Synchronization Mode Interrupt Enable
func (r *registerFdcanttieType) SetCsme(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldCsmeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldCsmeMask)
	}
}

const (
	RegisterFdcanttieFieldSogeShift = 3
	RegisterFdcanttieFieldSogeMask  = 0x8
)

// GetSoge Start of Gap Interrupt Enable
func (r *registerFdcanttieType) GetSoge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldSogeMask) != 0
}

// SetSoge Start of Gap Interrupt Enable
func (r *registerFdcanttieType) SetSoge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldSogeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldSogeMask)
	}
}

const (
	RegisterFdcanttieFieldRtmieShift = 4
	RegisterFdcanttieFieldRtmieMask  = 0x10
)

// GetRtmie Register Time Mark Interrupt Enable
func (r *registerFdcanttieType) GetRtmie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldRtmieMask) != 0
}

// SetRtmie Register Time Mark Interrupt Enable
func (r *registerFdcanttieType) SetRtmie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldRtmieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldRtmieMask)
	}
}

const (
	RegisterFdcanttieFieldTtmieShift = 5
	RegisterFdcanttieFieldTtmieMask  = 0x20
)

// GetTtmie Trigger Time Mark Event Internal Interrupt Enable
func (r *registerFdcanttieType) GetTtmie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldTtmieMask) != 0
}

// SetTtmie Trigger Time Mark Event Internal Interrupt Enable
func (r *registerFdcanttieType) SetTtmie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldTtmieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldTtmieMask)
	}
}

const (
	RegisterFdcanttieFieldSweeShift = 6
	RegisterFdcanttieFieldSweeMask  = 0x40
)

// GetSwee Stop Watch Event Interrupt Enable
func (r *registerFdcanttieType) GetSwee() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldSweeMask) != 0
}

// SetSwee Stop Watch Event Interrupt Enable
func (r *registerFdcanttieType) SetSwee(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldSweeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldSweeMask)
	}
}

const (
	RegisterFdcanttieFieldGtweShift = 7
	RegisterFdcanttieFieldGtweMask  = 0x80
)

// GetGtwe Global Time Wrap Interrupt Enable
func (r *registerFdcanttieType) GetGtwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldGtweMask) != 0
}

// SetGtwe Global Time Wrap Interrupt Enable
func (r *registerFdcanttieType) SetGtwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldGtweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldGtweMask)
	}
}

const (
	RegisterFdcanttieFieldGtdeShift = 8
	RegisterFdcanttieFieldGtdeMask  = 0x100
)

// GetGtde Global Time Discontinuity Interrupt Enable
func (r *registerFdcanttieType) GetGtde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldGtdeMask) != 0
}

// SetGtde Global Time Discontinuity Interrupt Enable
func (r *registerFdcanttieType) SetGtde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldGtdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldGtdeMask)
	}
}

const (
	RegisterFdcanttieFieldGteeShift = 9
	RegisterFdcanttieFieldGteeMask  = 0x200
)

// GetGtee Global Time Error Interrupt Enable
func (r *registerFdcanttieType) GetGtee() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldGteeMask) != 0
}

// SetGtee Global Time Error Interrupt Enable
func (r *registerFdcanttieType) SetGtee(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldGteeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldGteeMask)
	}
}

const (
	RegisterFdcanttieFieldTxueShift = 10
	RegisterFdcanttieFieldTxueMask  = 0x400
)

// GetTxue Tx Count Underflow Interrupt Enable
func (r *registerFdcanttieType) GetTxue() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldTxueMask) != 0
}

// SetTxue Tx Count Underflow Interrupt Enable
func (r *registerFdcanttieType) SetTxue(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldTxueMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldTxueMask)
	}
}

const (
	RegisterFdcanttieFieldTxoeShift = 11
	RegisterFdcanttieFieldTxoeMask  = 0x800
)

// GetTxoe Tx Count Overflow Interrupt Enable
func (r *registerFdcanttieType) GetTxoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldTxoeMask) != 0
}

// SetTxoe Tx Count Overflow Interrupt Enable
func (r *registerFdcanttieType) SetTxoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldTxoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldTxoeMask)
	}
}

const (
	RegisterFdcanttieFieldSe1eShift = 12
	RegisterFdcanttieFieldSe1eMask  = 0x1000
)

// GetSe1e Scheduling Error 1 Interrupt Enable
func (r *registerFdcanttieType) GetSe1e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldSe1eMask) != 0
}

// SetSe1e Scheduling Error 1 Interrupt Enable
func (r *registerFdcanttieType) SetSe1e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldSe1eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldSe1eMask)
	}
}

const (
	RegisterFdcanttieFieldSe2eShift = 13
	RegisterFdcanttieFieldSe2eMask  = 0x2000
)

// GetSe2e Scheduling Error 2 Interrupt Enable
func (r *registerFdcanttieType) GetSe2e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldSe2eMask) != 0
}

// SetSe2e Scheduling Error 2 Interrupt Enable
func (r *registerFdcanttieType) SetSe2e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldSe2eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldSe2eMask)
	}
}

const (
	RegisterFdcanttieFieldElceShift = 14
	RegisterFdcanttieFieldElceMask  = 0x4000
)

// GetElce Change Error Level Interrupt Enable
func (r *registerFdcanttieType) GetElce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldElceMask) != 0
}

// SetElce Change Error Level Interrupt Enable
func (r *registerFdcanttieType) SetElce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldElceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldElceMask)
	}
}

const (
	RegisterFdcanttieFieldIwtgeShift = 15
	RegisterFdcanttieFieldIwtgeMask  = 0x8000
)

// GetIwtge Initialization Watch Trigger Interrupt Enable
func (r *registerFdcanttieType) GetIwtge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldIwtgeMask) != 0
}

// SetIwtge Initialization Watch Trigger Interrupt Enable
func (r *registerFdcanttieType) SetIwtge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldIwtgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldIwtgeMask)
	}
}

const (
	RegisterFdcanttieFieldWteShift = 16
	RegisterFdcanttieFieldWteMask  = 0x10000
)

// GetWte Watch Trigger Interrupt Enable
func (r *registerFdcanttieType) GetWte() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldWteMask) != 0
}

// SetWte Watch Trigger Interrupt Enable
func (r *registerFdcanttieType) SetWte(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldWteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldWteMask)
	}
}

const (
	RegisterFdcanttieFieldAweShift = 17
	RegisterFdcanttieFieldAweMask  = 0x20000
)

// GetAwe Application Watchdog Interrupt Enable
func (r *registerFdcanttieType) GetAwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldAweMask) != 0
}

// SetAwe Application Watchdog Interrupt Enable
func (r *registerFdcanttieType) SetAwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldAweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldAweMask)
	}
}

const (
	RegisterFdcanttieFieldCereShift = 18
	RegisterFdcanttieFieldCereMask  = 0x40000
)

// GetCere Configuration Error Interrupt Enable
func (r *registerFdcanttieType) GetCere() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttieFieldCereMask) != 0
}

// SetCere Configuration Error Interrupt Enable
func (r *registerFdcanttieType) SetCere(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttieFieldCereMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttieFieldCereMask)
	}
}

// registerFdcanttilsType FDCAN TT Interrupt Line Select Register
type registerFdcanttilsType uint32

const (
	RegisterFdcanttilsFieldSbclShift = 0
	RegisterFdcanttilsFieldSbclMask  = 0x1
)

// GetSbcl Start of Basic Cycle Interrupt Line
func (r *registerFdcanttilsType) GetSbcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldSbclMask) != 0
}

// SetSbcl Start of Basic Cycle Interrupt Line
func (r *registerFdcanttilsType) SetSbcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldSbclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldSbclMask)
	}
}

const (
	RegisterFdcanttilsFieldSmclShift = 1
	RegisterFdcanttilsFieldSmclMask  = 0x2
)

// GetSmcl Start of Matrix Cycle Interrupt Line
func (r *registerFdcanttilsType) GetSmcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldSmclMask) != 0
}

// SetSmcl Start of Matrix Cycle Interrupt Line
func (r *registerFdcanttilsType) SetSmcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldSmclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldSmclMask)
	}
}

const (
	RegisterFdcanttilsFieldCsmlShift = 2
	RegisterFdcanttilsFieldCsmlMask  = 0x4
)

// GetCsml Change of Synchronization Mode Interrupt Line
func (r *registerFdcanttilsType) GetCsml() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldCsmlMask) != 0
}

// SetCsml Change of Synchronization Mode Interrupt Line
func (r *registerFdcanttilsType) SetCsml(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldCsmlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldCsmlMask)
	}
}

const (
	RegisterFdcanttilsFieldSoglShift = 3
	RegisterFdcanttilsFieldSoglMask  = 0x8
)

// GetSogl Start of Gap Interrupt Line
func (r *registerFdcanttilsType) GetSogl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldSoglMask) != 0
}

// SetSogl Start of Gap Interrupt Line
func (r *registerFdcanttilsType) SetSogl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldSoglMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldSoglMask)
	}
}

const (
	RegisterFdcanttilsFieldRtmilShift = 4
	RegisterFdcanttilsFieldRtmilMask  = 0x10
)

// GetRtmil Register Time Mark Interrupt Line
func (r *registerFdcanttilsType) GetRtmil() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldRtmilMask) != 0
}

// SetRtmil Register Time Mark Interrupt Line
func (r *registerFdcanttilsType) SetRtmil(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldRtmilMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldRtmilMask)
	}
}

const (
	RegisterFdcanttilsFieldTtmilShift = 5
	RegisterFdcanttilsFieldTtmilMask  = 0x20
)

// GetTtmil Trigger Time Mark Event Internal Interrupt Line
func (r *registerFdcanttilsType) GetTtmil() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldTtmilMask) != 0
}

// SetTtmil Trigger Time Mark Event Internal Interrupt Line
func (r *registerFdcanttilsType) SetTtmil(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldTtmilMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldTtmilMask)
	}
}

const (
	RegisterFdcanttilsFieldSwelShift = 6
	RegisterFdcanttilsFieldSwelMask  = 0x40
)

// GetSwel Stop Watch Event Interrupt Line
func (r *registerFdcanttilsType) GetSwel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldSwelMask) != 0
}

// SetSwel Stop Watch Event Interrupt Line
func (r *registerFdcanttilsType) SetSwel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldSwelMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldSwelMask)
	}
}

const (
	RegisterFdcanttilsFieldGtwlShift = 7
	RegisterFdcanttilsFieldGtwlMask  = 0x80
)

// GetGtwl Global Time Wrap Interrupt Line
func (r *registerFdcanttilsType) GetGtwl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldGtwlMask) != 0
}

// SetGtwl Global Time Wrap Interrupt Line
func (r *registerFdcanttilsType) SetGtwl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldGtwlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldGtwlMask)
	}
}

const (
	RegisterFdcanttilsFieldGtdlShift = 8
	RegisterFdcanttilsFieldGtdlMask  = 0x100
)

// GetGtdl Global Time Discontinuity Interrupt Line
func (r *registerFdcanttilsType) GetGtdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldGtdlMask) != 0
}

// SetGtdl Global Time Discontinuity Interrupt Line
func (r *registerFdcanttilsType) SetGtdl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldGtdlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldGtdlMask)
	}
}

const (
	RegisterFdcanttilsFieldGtelShift = 9
	RegisterFdcanttilsFieldGtelMask  = 0x200
)

// GetGtel Global Time Error Interrupt Line
func (r *registerFdcanttilsType) GetGtel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldGtelMask) != 0
}

// SetGtel Global Time Error Interrupt Line
func (r *registerFdcanttilsType) SetGtel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldGtelMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldGtelMask)
	}
}

const (
	RegisterFdcanttilsFieldTxulShift = 10
	RegisterFdcanttilsFieldTxulMask  = 0x400
)

// GetTxul Tx Count Underflow Interrupt Line
func (r *registerFdcanttilsType) GetTxul() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldTxulMask) != 0
}

// SetTxul Tx Count Underflow Interrupt Line
func (r *registerFdcanttilsType) SetTxul(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldTxulMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldTxulMask)
	}
}

const (
	RegisterFdcanttilsFieldTxolShift = 11
	RegisterFdcanttilsFieldTxolMask  = 0x800
)

// GetTxol Tx Count Overflow Interrupt Line
func (r *registerFdcanttilsType) GetTxol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldTxolMask) != 0
}

// SetTxol Tx Count Overflow Interrupt Line
func (r *registerFdcanttilsType) SetTxol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldTxolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldTxolMask)
	}
}

const (
	RegisterFdcanttilsFieldSe1lShift = 12
	RegisterFdcanttilsFieldSe1lMask  = 0x1000
)

// GetSe1l Scheduling Error 1 Interrupt Line
func (r *registerFdcanttilsType) GetSe1l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldSe1lMask) != 0
}

// SetSe1l Scheduling Error 1 Interrupt Line
func (r *registerFdcanttilsType) SetSe1l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldSe1lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldSe1lMask)
	}
}

const (
	RegisterFdcanttilsFieldSe2lShift = 13
	RegisterFdcanttilsFieldSe2lMask  = 0x2000
)

// GetSe2l Scheduling Error 2 Interrupt Line
func (r *registerFdcanttilsType) GetSe2l() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldSe2lMask) != 0
}

// SetSe2l Scheduling Error 2 Interrupt Line
func (r *registerFdcanttilsType) SetSe2l(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldSe2lMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldSe2lMask)
	}
}

const (
	RegisterFdcanttilsFieldElclShift = 14
	RegisterFdcanttilsFieldElclMask  = 0x4000
)

// GetElcl Change Error Level Interrupt Line
func (r *registerFdcanttilsType) GetElcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldElclMask) != 0
}

// SetElcl Change Error Level Interrupt Line
func (r *registerFdcanttilsType) SetElcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldElclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldElclMask)
	}
}

const (
	RegisterFdcanttilsFieldIwtglShift = 15
	RegisterFdcanttilsFieldIwtglMask  = 0x8000
)

// GetIwtgl Initialization Watch Trigger Interrupt Line
func (r *registerFdcanttilsType) GetIwtgl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldIwtglMask) != 0
}

// SetIwtgl Initialization Watch Trigger Interrupt Line
func (r *registerFdcanttilsType) SetIwtgl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldIwtglMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldIwtglMask)
	}
}

const (
	RegisterFdcanttilsFieldWtlShift = 16
	RegisterFdcanttilsFieldWtlMask  = 0x10000
)

// GetWtl Watch Trigger Interrupt Line
func (r *registerFdcanttilsType) GetWtl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldWtlMask) != 0
}

// SetWtl Watch Trigger Interrupt Line
func (r *registerFdcanttilsType) SetWtl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldWtlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldWtlMask)
	}
}

const (
	RegisterFdcanttilsFieldAwlShift = 17
	RegisterFdcanttilsFieldAwlMask  = 0x20000
)

// GetAwl Application Watchdog Interrupt Line
func (r *registerFdcanttilsType) GetAwl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldAwlMask) != 0
}

// SetAwl Application Watchdog Interrupt Line
func (r *registerFdcanttilsType) SetAwl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldAwlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldAwlMask)
	}
}

const (
	RegisterFdcanttilsFieldCerlShift = 18
	RegisterFdcanttilsFieldCerlMask  = 0x40000
)

// GetCerl Configuration Error Interrupt Line
func (r *registerFdcanttilsType) GetCerl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttilsFieldCerlMask) != 0
}

// SetCerl Configuration Error Interrupt Line
func (r *registerFdcanttilsType) SetCerl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttilsFieldCerlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttilsFieldCerlMask)
	}
}

// registerFdcanttostType FDCAN TT Operation Status Register
type registerFdcanttostType uint32

const (
	RegisterFdcanttostFieldElShift = 0
	RegisterFdcanttostFieldElMask  = 0x3
)

// GetEl Error Level
func (r *registerFdcanttostType) GetEl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldElMask) >> RegisterFdcanttostFieldElShift)
}

// SetEl Error Level
func (r *registerFdcanttostType) SetEl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldElMask)|(uint32(value)<<RegisterFdcanttostFieldElShift))
}

const (
	RegisterFdcanttostFieldMsShift = 2
	RegisterFdcanttostFieldMsMask  = 0xc
)

// GetMs Master State.
func (r *registerFdcanttostType) GetMs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldMsMask) >> RegisterFdcanttostFieldMsShift)
}

// SetMs Master State.
func (r *registerFdcanttostType) SetMs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldMsMask)|(uint32(value)<<RegisterFdcanttostFieldMsShift))
}

const (
	RegisterFdcanttostFieldSysShift = 4
	RegisterFdcanttostFieldSysMask  = 0x30
)

// GetSys Synchronization State
func (r *registerFdcanttostType) GetSys() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldSysMask) >> RegisterFdcanttostFieldSysShift)
}

// SetSys Synchronization State
func (r *registerFdcanttostType) SetSys(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldSysMask)|(uint32(value)<<RegisterFdcanttostFieldSysShift))
}

const (
	RegisterFdcanttostFieldGtpShift = 6
	RegisterFdcanttostFieldGtpMask  = 0x40
)

// GetGtp Quality of Global Time Phase
func (r *registerFdcanttostType) GetGtp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldGtpMask) != 0
}

// SetGtp Quality of Global Time Phase
func (r *registerFdcanttostType) SetGtp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttostFieldGtpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldGtpMask)
	}
}

const (
	RegisterFdcanttostFieldQcsShift = 7
	RegisterFdcanttostFieldQcsMask  = 0x80
)

// GetQcs Quality of Clock Speed
func (r *registerFdcanttostType) GetQcs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldQcsMask) != 0
}

// SetQcs Quality of Clock Speed
func (r *registerFdcanttostType) SetQcs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttostFieldQcsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldQcsMask)
	}
}

const (
	RegisterFdcanttostFieldRtoShift = 8
	RegisterFdcanttostFieldRtoMask  = 0xff00
)

// GetRto Reference Trigger Offset
func (r *registerFdcanttostType) GetRto() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldRtoMask) >> RegisterFdcanttostFieldRtoShift)
}

// SetRto Reference Trigger Offset
func (r *registerFdcanttostType) SetRto(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldRtoMask)|(uint32(value)<<RegisterFdcanttostFieldRtoShift))
}

const (
	RegisterFdcanttostFieldWgtdShift = 22
	RegisterFdcanttostFieldWgtdMask  = 0x400000
)

// GetWgtd Wait for Global Time Discontinuity
func (r *registerFdcanttostType) GetWgtd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldWgtdMask) != 0
}

// SetWgtd Wait for Global Time Discontinuity
func (r *registerFdcanttostType) SetWgtd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttostFieldWgtdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldWgtdMask)
	}
}

const (
	RegisterFdcanttostFieldGfiShift = 23
	RegisterFdcanttostFieldGfiMask  = 0x800000
)

// GetGfi Gap Finished Indicator.
func (r *registerFdcanttostType) GetGfi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldGfiMask) != 0
}

// SetGfi Gap Finished Indicator.
func (r *registerFdcanttostType) SetGfi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttostFieldGfiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldGfiMask)
	}
}

const (
	RegisterFdcanttostFieldTmpShift = 24
	RegisterFdcanttostFieldTmpMask  = 0x7000000
)

// GetTmp Time Master Priority
func (r *registerFdcanttostType) GetTmp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldTmpMask) >> RegisterFdcanttostFieldTmpShift)
}

// SetTmp Time Master Priority
func (r *registerFdcanttostType) SetTmp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldTmpMask)|(uint32(value)<<RegisterFdcanttostFieldTmpShift))
}

const (
	RegisterFdcanttostFieldGsiShift = 27
	RegisterFdcanttostFieldGsiMask  = 0x8000000
)

// GetGsi Gap Started Indicator.
func (r *registerFdcanttostType) GetGsi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldGsiMask) != 0
}

// SetGsi Gap Started Indicator.
func (r *registerFdcanttostType) SetGsi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttostFieldGsiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldGsiMask)
	}
}

const (
	RegisterFdcanttostFieldWfeShift = 28
	RegisterFdcanttostFieldWfeMask  = 0x10000000
)

// GetWfe Wait for Event
func (r *registerFdcanttostType) GetWfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldWfeMask) != 0
}

// SetWfe Wait for Event
func (r *registerFdcanttostType) SetWfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttostFieldWfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldWfeMask)
	}
}

const (
	RegisterFdcanttostFieldAweShift = 29
	RegisterFdcanttostFieldAweMask  = 0x20000000
)

// GetAwe Application Watchdog Event
func (r *registerFdcanttostType) GetAwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldAweMask) != 0
}

// SetAwe Application Watchdog Event
func (r *registerFdcanttostType) SetAwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttostFieldAweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldAweMask)
	}
}

const (
	RegisterFdcanttostFieldWecsShift = 30
	RegisterFdcanttostFieldWecsMask  = 0x40000000
)

// GetWecs Wait for External Clock Synchronization
func (r *registerFdcanttostType) GetWecs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldWecsMask) != 0
}

// SetWecs Wait for External Clock Synchronization
func (r *registerFdcanttostType) SetWecs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttostFieldWecsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldWecsMask)
	}
}

const (
	RegisterFdcanttostFieldSplShift = 31
	RegisterFdcanttostFieldSplMask  = 0x80000000
)

// GetSpl Schedule Phase Lock
func (r *registerFdcanttostType) GetSpl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttostFieldSplMask) != 0
}

// SetSpl Schedule Phase Lock
func (r *registerFdcanttostType) SetSpl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFdcanttostFieldSplMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttostFieldSplMask)
	}
}

// registerFdcanturnaType FDCAN TUR Numerator Actual Register
type registerFdcanturnaType uint32

const (
	RegisterFdcanturnaFieldNavShift = 0
	RegisterFdcanturnaFieldNavMask  = 0x3ffff
)

// GetNav Numerator Actual Value
func (r *registerFdcanturnaType) GetNav() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanturnaFieldNavMask) >> RegisterFdcanturnaFieldNavShift)
}

// SetNav Numerator Actual Value
func (r *registerFdcanturnaType) SetNav(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanturnaFieldNavMask)|(uint32(value)<<RegisterFdcanturnaFieldNavShift))
}

// registerFdcanttlgtType FDCAN TT Local and Global Time Register
type registerFdcanttlgtType uint32

const (
	RegisterFdcanttlgtFieldLtShift = 0
	RegisterFdcanttlgtFieldLtMask  = 0xffff
)

// GetLt Local Time
func (r *registerFdcanttlgtType) GetLt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttlgtFieldLtMask) >> RegisterFdcanttlgtFieldLtShift)
}

// SetLt Local Time
func (r *registerFdcanttlgtType) SetLt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttlgtFieldLtMask)|(uint32(value)<<RegisterFdcanttlgtFieldLtShift))
}

const (
	RegisterFdcanttlgtFieldGtShift = 16
	RegisterFdcanttlgtFieldGtMask  = 0xffff0000
)

// GetGt Global Time
func (r *registerFdcanttlgtType) GetGt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttlgtFieldGtMask) >> RegisterFdcanttlgtFieldGtShift)
}

// SetGt Global Time
func (r *registerFdcanttlgtType) SetGt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttlgtFieldGtMask)|(uint32(value)<<RegisterFdcanttlgtFieldGtShift))
}

// registerFdcanttctcType FDCAN TT Cycle Time and Count Register
type registerFdcanttctcType uint32

const (
	RegisterFdcanttctcFieldCtShift = 0
	RegisterFdcanttctcFieldCtMask  = 0xffff
)

// GetCt Cycle Time
func (r *registerFdcanttctcType) GetCt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttctcFieldCtMask) >> RegisterFdcanttctcFieldCtShift)
}

// SetCt Cycle Time
func (r *registerFdcanttctcType) SetCt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttctcFieldCtMask)|(uint32(value)<<RegisterFdcanttctcFieldCtShift))
}

const (
	RegisterFdcanttctcFieldCcShift = 16
	RegisterFdcanttctcFieldCcMask  = 0x3f0000
)

// GetCc Cycle Count
func (r *registerFdcanttctcType) GetCc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttctcFieldCcMask) >> RegisterFdcanttctcFieldCcShift)
}

// SetCc Cycle Count
func (r *registerFdcanttctcType) SetCc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttctcFieldCcMask)|(uint32(value)<<RegisterFdcanttctcFieldCcShift))
}

// registerFdcanttcptType FDCAN TT Capture Time Register
type registerFdcanttcptType uint32

const (
	RegisterFdcanttcptFieldCtShift = 0
	RegisterFdcanttcptFieldCtMask  = 0x3f
)

// GetCt Cycle Count Value
func (r *registerFdcanttcptType) GetCt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttcptFieldCtMask) >> RegisterFdcanttcptFieldCtShift)
}

// SetCt Cycle Count Value
func (r *registerFdcanttcptType) SetCt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttcptFieldCtMask)|(uint32(value)<<RegisterFdcanttcptFieldCtShift))
}

const (
	RegisterFdcanttcptFieldSwvShift = 16
	RegisterFdcanttcptFieldSwvMask  = 0xffff0000
)

// GetSwv Stop Watch Value
func (r *registerFdcanttcptType) GetSwv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttcptFieldSwvMask) >> RegisterFdcanttcptFieldSwvShift)
}

// SetSwv Stop Watch Value
func (r *registerFdcanttcptType) SetSwv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttcptFieldSwvMask)|(uint32(value)<<RegisterFdcanttcptFieldSwvShift))
}

// registerFdcanttcsmType FDCAN TT Cycle Sync Mark Register
type registerFdcanttcsmType uint32

const (
	RegisterFdcanttcsmFieldCsmShift = 0
	RegisterFdcanttcsmFieldCsmMask  = 0xffff
)

// GetCsm Cycle Sync Mark
func (r *registerFdcanttcsmType) GetCsm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterFdcanttcsmFieldCsmMask) >> RegisterFdcanttcsmFieldCsmShift)
}

// SetCsm Cycle Sync Mark
func (r *registerFdcanttcsmType) SetCsm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcanttcsmFieldCsmMask)|(uint32(value)<<RegisterFdcanttcsmFieldCsmShift))
}

// registerFdcantttsType FDCAN TT Trigger Select Register
type registerFdcantttsType uint32

const (
	RegisterFdcantttsFieldSwtdelShift = 0
	RegisterFdcantttsFieldSwtdelMask  = 0x3
)

// GetSwtdel Stop watch trigger input selection
func (r *registerFdcantttsType) GetSwtdel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantttsFieldSwtdelMask) >> RegisterFdcantttsFieldSwtdelShift)
}

// SetSwtdel Stop watch trigger input selection
func (r *registerFdcantttsType) SetSwtdel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantttsFieldSwtdelMask)|(uint32(value)<<RegisterFdcantttsFieldSwtdelShift))
}

const (
	RegisterFdcantttsFieldEvtselShift = 4
	RegisterFdcantttsFieldEvtselMask  = 0x30
)

// GetEvtsel Event trigger input selection
func (r *registerFdcantttsType) GetEvtsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterFdcantttsFieldEvtselMask) >> RegisterFdcantttsFieldEvtselShift)
}

// SetEvtsel Event trigger input selection
func (r *registerFdcantttsType) SetEvtsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFdcantttsFieldEvtselMask)|(uint32(value)<<RegisterFdcantttsFieldEvtselShift))
}
