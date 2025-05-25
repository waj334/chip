package hash

import (
	"unsafe"
	"volatile"
)

var (
	Hash = (*_hash)(unsafe.Pointer(uintptr(0x48021400)))
)

type _hash struct {
	Cr       registerCrType
	Din      registerDinType
	Str      registerStrType
	Hr0      registerHr0Type
	Hr1      registerHr1Type
	Hr2      registerHr2Type
	Hr3      registerHr3Type
	Hr4      registerHr4Type
	Imr      registerImrType
	Sr       registerSrType
	_        [208]byte
	Csr0     registerCsr0Type
	Csr1     registerCsr1Type
	Csr2     registerCsr2Type
	Csr3     registerCsr3Type
	Csr4     registerCsr4Type
	Csr5     registerCsr5Type
	Csr6     registerCsr6Type
	Csr7     registerCsr7Type
	Csr8     registerCsr8Type
	Csr9     registerCsr9Type
	Csr10    registerCsr10Type
	Csr11    registerCsr11Type
	Csr12    registerCsr12Type
	Csr13    registerCsr13Type
	Csr14    registerCsr14Type
	Csr15    registerCsr15Type
	Csr16    registerCsr16Type
	Csr17    registerCsr17Type
	Csr18    registerCsr18Type
	Csr19    registerCsr19Type
	Csr20    registerCsr20Type
	Csr21    registerCsr21Type
	Csr22    registerCsr22Type
	Csr23    registerCsr23Type
	Csr24    registerCsr24Type
	Csr25    registerCsr25Type
	Csr26    registerCsr26Type
	Csr27    registerCsr27Type
	Csr28    registerCsr28Type
	Csr29    registerCsr29Type
	Csr30    registerCsr30Type
	Csr31    registerCsr31Type
	Csr32    registerCsr32Type
	Csr33    registerCsr33Type
	Csr34    registerCsr34Type
	Csr35    registerCsr35Type
	Csr36    registerCsr36Type
	Csr37    registerCsr37Type
	Csr38    registerCsr38Type
	Csr39    registerCsr39Type
	Csr40    registerCsr40Type
	Csr41    registerCsr41Type
	Csr42    registerCsr42Type
	Csr43    registerCsr43Type
	Csr44    registerCsr44Type
	Csr45    registerCsr45Type
	Csr46    registerCsr46Type
	Csr47    registerCsr47Type
	Csr48    registerCsr48Type
	Csr49    registerCsr49Type
	Csr50    registerCsr50Type
	Csr51    registerCsr51Type
	Csr52    registerCsr52Type
	Csr53    registerCsr53Type
	_        [320]byte
	Hash_hr0 registerHash_hr0Type
	Hash_hr1 registerHash_hr1Type
	Hash_hr2 registerHash_hr2Type
	Hash_hr3 registerHash_hr3Type
	Hash_hr4 registerHash_hr4Type
	Hash_hr5 registerHash_hr5Type
	Hash_hr6 registerHash_hr6Type
	Hash_hr7 registerHash_hr7Type
}

// registerCrType control register
type registerCrType uint32

const (
	RegisterCrFieldInitShift = 2
	RegisterCrFieldInitMask  = 0x4
)

// SetInit Initialize message digest calculation
func (r *registerCrType) SetInit(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldInitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldInitMask)
	}
}

const (
	RegisterCrFieldDmaeShift = 3
	RegisterCrFieldDmaeMask  = 0x8
)

// GetDmae DMA enable
func (r *registerCrType) GetDmae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDmaeMask) != 0
}

// SetDmae DMA enable
func (r *registerCrType) SetDmae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDmaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDmaeMask)
	}
}

const (
	RegisterCrFieldDatatypeShift = 4
	RegisterCrFieldDatatypeMask  = 0x30
)

// GetDatatype Data type selection
func (r *registerCrType) GetDatatype() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDatatypeMask) >> RegisterCrFieldDatatypeShift)
}

// SetDatatype Data type selection
func (r *registerCrType) SetDatatype(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDatatypeMask)|(uint32(value)<<RegisterCrFieldDatatypeShift))
}

const (
	RegisterCrFieldModeShift = 6
	RegisterCrFieldModeMask  = 0x40
)

// GetMode Mode selection
func (r *registerCrType) GetMode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldModeMask) != 0
}

// SetMode Mode selection
func (r *registerCrType) SetMode(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldModeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldModeMask)
	}
}

const (
	RegisterCrFieldAlgo0Shift = 7
	RegisterCrFieldAlgo0Mask  = 0x80
)

// GetAlgo0 Algorithm selection
func (r *registerCrType) GetAlgo0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAlgo0Mask) != 0
}

// SetAlgo0 Algorithm selection
func (r *registerCrType) SetAlgo0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAlgo0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAlgo0Mask)
	}
}

const (
	RegisterCrFieldNbwShift = 8
	RegisterCrFieldNbwMask  = 0xf00
)

// GetNbw Number of words already pushed
func (r *registerCrType) GetNbw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldNbwMask) >> RegisterCrFieldNbwShift)
}

const (
	RegisterCrFieldDinneShift = 12
	RegisterCrFieldDinneMask  = 0x1000
)

// GetDinne DIN not empty
func (r *registerCrType) GetDinne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDinneMask) != 0
}

const (
	RegisterCrFieldMdmatShift = 13
	RegisterCrFieldMdmatMask  = 0x2000
)

// GetMdmat Multiple DMA Transfers
func (r *registerCrType) GetMdmat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldMdmatMask) != 0
}

// SetMdmat Multiple DMA Transfers
func (r *registerCrType) SetMdmat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldMdmatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldMdmatMask)
	}
}

const (
	RegisterCrFieldLkeyShift = 16
	RegisterCrFieldLkeyMask  = 0x10000
)

// GetLkey Long key selection
func (r *registerCrType) GetLkey() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLkeyMask) != 0
}

// SetLkey Long key selection
func (r *registerCrType) SetLkey(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldLkeyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldLkeyMask)
	}
}

const (
	RegisterCrFieldAlgo1Shift = 18
	RegisterCrFieldAlgo1Mask  = 0x40000
)

// GetAlgo1 ALGO
func (r *registerCrType) GetAlgo1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAlgo1Mask) != 0
}

// SetAlgo1 ALGO
func (r *registerCrType) SetAlgo1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAlgo1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAlgo1Mask)
	}
}

// registerDinType data input register
type registerDinType uint32

const (
	RegisterDinFieldDatainShift = 0
	RegisterDinFieldDatainMask  = 0xffffffff
)

// GetDatain Data input
func (r *registerDinType) GetDatain() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDinFieldDatainMask) >> RegisterDinFieldDatainShift)
}

// SetDatain Data input
func (r *registerDinType) SetDatain(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinFieldDatainMask)|(uint32(value)<<RegisterDinFieldDatainShift))
}

// registerStrType start register
type registerStrType uint32

const (
	RegisterStrFieldNblwShift = 0
	RegisterStrFieldNblwMask  = 0x1f
)

// GetNblw Number of valid bits in the last word of the message
func (r *registerStrType) GetNblw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterStrFieldNblwMask) >> RegisterStrFieldNblwShift)
}

// SetNblw Number of valid bits in the last word of the message
func (r *registerStrType) SetNblw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterStrFieldNblwMask)|(uint32(value)<<RegisterStrFieldNblwShift))
}

const (
	RegisterStrFieldDcalShift = 8
	RegisterStrFieldDcalMask  = 0x100
)

// SetDcal Digest calculation
func (r *registerStrType) SetDcal(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStrFieldDcalMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStrFieldDcalMask)
	}
}

// registerHr0Type digest registers
type registerHr0Type uint32

const (
	RegisterHr0FieldH0Shift = 0
	RegisterHr0FieldH0Mask  = 0xffffffff
)

// GetH0 H0
func (r *registerHr0Type) GetH0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHr0FieldH0Mask) >> RegisterHr0FieldH0Shift)
}

// SetH0 H0
func (r *registerHr0Type) SetH0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHr0FieldH0Mask)|(uint32(value)<<RegisterHr0FieldH0Shift))
}

// registerHr1Type digest registers
type registerHr1Type uint32

const (
	RegisterHr1FieldH1Shift = 0
	RegisterHr1FieldH1Mask  = 0xffffffff
)

// GetH1 H1
func (r *registerHr1Type) GetH1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHr1FieldH1Mask) >> RegisterHr1FieldH1Shift)
}

// SetH1 H1
func (r *registerHr1Type) SetH1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHr1FieldH1Mask)|(uint32(value)<<RegisterHr1FieldH1Shift))
}

// registerHr2Type digest registers
type registerHr2Type uint32

const (
	RegisterHr2FieldH2Shift = 0
	RegisterHr2FieldH2Mask  = 0xffffffff
)

// GetH2 H2
func (r *registerHr2Type) GetH2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHr2FieldH2Mask) >> RegisterHr2FieldH2Shift)
}

// SetH2 H2
func (r *registerHr2Type) SetH2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHr2FieldH2Mask)|(uint32(value)<<RegisterHr2FieldH2Shift))
}

// registerHr3Type digest registers
type registerHr3Type uint32

const (
	RegisterHr3FieldH3Shift = 0
	RegisterHr3FieldH3Mask  = 0xffffffff
)

// GetH3 H3
func (r *registerHr3Type) GetH3() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHr3FieldH3Mask) >> RegisterHr3FieldH3Shift)
}

// SetH3 H3
func (r *registerHr3Type) SetH3(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHr3FieldH3Mask)|(uint32(value)<<RegisterHr3FieldH3Shift))
}

// registerHr4Type digest registers
type registerHr4Type uint32

const (
	RegisterHr4FieldH4Shift = 0
	RegisterHr4FieldH4Mask  = 0xffffffff
)

// GetH4 H4
func (r *registerHr4Type) GetH4() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHr4FieldH4Mask) >> RegisterHr4FieldH4Shift)
}

// SetH4 H4
func (r *registerHr4Type) SetH4(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHr4FieldH4Mask)|(uint32(value)<<RegisterHr4FieldH4Shift))
}

// registerImrType interrupt enable register
type registerImrType uint32

const (
	RegisterImrFieldDinieShift = 0
	RegisterImrFieldDinieMask  = 0x1
)

// GetDinie Data input interrupt enable
func (r *registerImrType) GetDinie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldDinieMask) != 0
}

// SetDinie Data input interrupt enable
func (r *registerImrType) SetDinie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImrFieldDinieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImrFieldDinieMask)
	}
}

const (
	RegisterImrFieldDcieShift = 1
	RegisterImrFieldDcieMask  = 0x2
)

// GetDcie Digest calculation completion interrupt enable
func (r *registerImrType) GetDcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldDcieMask) != 0
}

// SetDcie Digest calculation completion interrupt enable
func (r *registerImrType) SetDcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImrFieldDcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImrFieldDcieMask)
	}
}

// registerSrType status register
type registerSrType uint32

const (
	RegisterSrFieldDinisShift = 0
	RegisterSrFieldDinisMask  = 0x1
)

// GetDinis Data input interrupt status
func (r *registerSrType) GetDinis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDinisMask) != 0
}

// SetDinis Data input interrupt status
func (r *registerSrType) SetDinis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldDinisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldDinisMask)
	}
}

const (
	RegisterSrFieldDcisShift = 1
	RegisterSrFieldDcisMask  = 0x2
)

// GetDcis Digest calculation completion interrupt status
func (r *registerSrType) GetDcis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDcisMask) != 0
}

// SetDcis Digest calculation completion interrupt status
func (r *registerSrType) SetDcis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldDcisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldDcisMask)
	}
}

const (
	RegisterSrFieldDmasShift = 2
	RegisterSrFieldDmasMask  = 0x4
)

// GetDmas DMA Status
func (r *registerSrType) GetDmas() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDmasMask) != 0
}

const (
	RegisterSrFieldBusyShift = 3
	RegisterSrFieldBusyMask  = 0x8
)

// GetBusy Busy bit
func (r *registerSrType) GetBusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldBusyMask) != 0
}

// registerCsr0Type context swap registers
type registerCsr0Type uint32

const (
	RegisterCsr0FieldCsr0Shift = 0
	RegisterCsr0FieldCsr0Mask  = 0xffffffff
)

// GetCsr0 CSR0
func (r *registerCsr0Type) GetCsr0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr0FieldCsr0Mask) >> RegisterCsr0FieldCsr0Shift)
}

// SetCsr0 CSR0
func (r *registerCsr0Type) SetCsr0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr0FieldCsr0Mask)|(uint32(value)<<RegisterCsr0FieldCsr0Shift))
}

// registerCsr1Type context swap registers
type registerCsr1Type uint32

const (
	RegisterCsr1FieldCsr1Shift = 0
	RegisterCsr1FieldCsr1Mask  = 0xffffffff
)

// GetCsr1 CSR1
func (r *registerCsr1Type) GetCsr1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr1FieldCsr1Mask) >> RegisterCsr1FieldCsr1Shift)
}

// SetCsr1 CSR1
func (r *registerCsr1Type) SetCsr1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr1FieldCsr1Mask)|(uint32(value)<<RegisterCsr1FieldCsr1Shift))
}

// registerCsr2Type context swap registers
type registerCsr2Type uint32

const (
	RegisterCsr2FieldCsr2Shift = 0
	RegisterCsr2FieldCsr2Mask  = 0xffffffff
)

// GetCsr2 CSR2
func (r *registerCsr2Type) GetCsr2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr2FieldCsr2Mask) >> RegisterCsr2FieldCsr2Shift)
}

// SetCsr2 CSR2
func (r *registerCsr2Type) SetCsr2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr2FieldCsr2Mask)|(uint32(value)<<RegisterCsr2FieldCsr2Shift))
}

// registerCsr3Type context swap registers
type registerCsr3Type uint32

const (
	RegisterCsr3FieldCsr3Shift = 0
	RegisterCsr3FieldCsr3Mask  = 0xffffffff
)

// GetCsr3 CSR3
func (r *registerCsr3Type) GetCsr3() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr3FieldCsr3Mask) >> RegisterCsr3FieldCsr3Shift)
}

// SetCsr3 CSR3
func (r *registerCsr3Type) SetCsr3(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr3FieldCsr3Mask)|(uint32(value)<<RegisterCsr3FieldCsr3Shift))
}

// registerCsr4Type context swap registers
type registerCsr4Type uint32

const (
	RegisterCsr4FieldCsr4Shift = 0
	RegisterCsr4FieldCsr4Mask  = 0xffffffff
)

// GetCsr4 CSR4
func (r *registerCsr4Type) GetCsr4() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr4FieldCsr4Mask) >> RegisterCsr4FieldCsr4Shift)
}

// SetCsr4 CSR4
func (r *registerCsr4Type) SetCsr4(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr4FieldCsr4Mask)|(uint32(value)<<RegisterCsr4FieldCsr4Shift))
}

// registerCsr5Type context swap registers
type registerCsr5Type uint32

const (
	RegisterCsr5FieldCsr5Shift = 0
	RegisterCsr5FieldCsr5Mask  = 0xffffffff
)

// GetCsr5 CSR5
func (r *registerCsr5Type) GetCsr5() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr5FieldCsr5Mask) >> RegisterCsr5FieldCsr5Shift)
}

// SetCsr5 CSR5
func (r *registerCsr5Type) SetCsr5(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr5FieldCsr5Mask)|(uint32(value)<<RegisterCsr5FieldCsr5Shift))
}

// registerCsr6Type context swap registers
type registerCsr6Type uint32

const (
	RegisterCsr6FieldCsr6Shift = 0
	RegisterCsr6FieldCsr6Mask  = 0xffffffff
)

// GetCsr6 CSR6
func (r *registerCsr6Type) GetCsr6() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr6FieldCsr6Mask) >> RegisterCsr6FieldCsr6Shift)
}

// SetCsr6 CSR6
func (r *registerCsr6Type) SetCsr6(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr6FieldCsr6Mask)|(uint32(value)<<RegisterCsr6FieldCsr6Shift))
}

// registerCsr7Type context swap registers
type registerCsr7Type uint32

const (
	RegisterCsr7FieldCsr7Shift = 0
	RegisterCsr7FieldCsr7Mask  = 0xffffffff
)

// GetCsr7 CSR7
func (r *registerCsr7Type) GetCsr7() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr7FieldCsr7Mask) >> RegisterCsr7FieldCsr7Shift)
}

// SetCsr7 CSR7
func (r *registerCsr7Type) SetCsr7(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr7FieldCsr7Mask)|(uint32(value)<<RegisterCsr7FieldCsr7Shift))
}

// registerCsr8Type context swap registers
type registerCsr8Type uint32

const (
	RegisterCsr8FieldCsr8Shift = 0
	RegisterCsr8FieldCsr8Mask  = 0xffffffff
)

// GetCsr8 CSR8
func (r *registerCsr8Type) GetCsr8() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr8FieldCsr8Mask) >> RegisterCsr8FieldCsr8Shift)
}

// SetCsr8 CSR8
func (r *registerCsr8Type) SetCsr8(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr8FieldCsr8Mask)|(uint32(value)<<RegisterCsr8FieldCsr8Shift))
}

// registerCsr9Type context swap registers
type registerCsr9Type uint32

const (
	RegisterCsr9FieldCsr9Shift = 0
	RegisterCsr9FieldCsr9Mask  = 0xffffffff
)

// GetCsr9 CSR9
func (r *registerCsr9Type) GetCsr9() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr9FieldCsr9Mask) >> RegisterCsr9FieldCsr9Shift)
}

// SetCsr9 CSR9
func (r *registerCsr9Type) SetCsr9(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr9FieldCsr9Mask)|(uint32(value)<<RegisterCsr9FieldCsr9Shift))
}

// registerCsr10Type context swap registers
type registerCsr10Type uint32

const (
	RegisterCsr10FieldCsr10Shift = 0
	RegisterCsr10FieldCsr10Mask  = 0xffffffff
)

// GetCsr10 CSR10
func (r *registerCsr10Type) GetCsr10() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr10FieldCsr10Mask) >> RegisterCsr10FieldCsr10Shift)
}

// SetCsr10 CSR10
func (r *registerCsr10Type) SetCsr10(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr10FieldCsr10Mask)|(uint32(value)<<RegisterCsr10FieldCsr10Shift))
}

// registerCsr11Type context swap registers
type registerCsr11Type uint32

const (
	RegisterCsr11FieldCsr11Shift = 0
	RegisterCsr11FieldCsr11Mask  = 0xffffffff
)

// GetCsr11 CSR11
func (r *registerCsr11Type) GetCsr11() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr11FieldCsr11Mask) >> RegisterCsr11FieldCsr11Shift)
}

// SetCsr11 CSR11
func (r *registerCsr11Type) SetCsr11(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr11FieldCsr11Mask)|(uint32(value)<<RegisterCsr11FieldCsr11Shift))
}

// registerCsr12Type context swap registers
type registerCsr12Type uint32

const (
	RegisterCsr12FieldCsr12Shift = 0
	RegisterCsr12FieldCsr12Mask  = 0xffffffff
)

// GetCsr12 CSR12
func (r *registerCsr12Type) GetCsr12() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr12FieldCsr12Mask) >> RegisterCsr12FieldCsr12Shift)
}

// SetCsr12 CSR12
func (r *registerCsr12Type) SetCsr12(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr12FieldCsr12Mask)|(uint32(value)<<RegisterCsr12FieldCsr12Shift))
}

// registerCsr13Type context swap registers
type registerCsr13Type uint32

const (
	RegisterCsr13FieldCsr13Shift = 0
	RegisterCsr13FieldCsr13Mask  = 0xffffffff
)

// GetCsr13 CSR13
func (r *registerCsr13Type) GetCsr13() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr13FieldCsr13Mask) >> RegisterCsr13FieldCsr13Shift)
}

// SetCsr13 CSR13
func (r *registerCsr13Type) SetCsr13(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr13FieldCsr13Mask)|(uint32(value)<<RegisterCsr13FieldCsr13Shift))
}

// registerCsr14Type context swap registers
type registerCsr14Type uint32

const (
	RegisterCsr14FieldCsr14Shift = 0
	RegisterCsr14FieldCsr14Mask  = 0xffffffff
)

// GetCsr14 CSR14
func (r *registerCsr14Type) GetCsr14() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr14FieldCsr14Mask) >> RegisterCsr14FieldCsr14Shift)
}

// SetCsr14 CSR14
func (r *registerCsr14Type) SetCsr14(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr14FieldCsr14Mask)|(uint32(value)<<RegisterCsr14FieldCsr14Shift))
}

// registerCsr15Type context swap registers
type registerCsr15Type uint32

const (
	RegisterCsr15FieldCsr15Shift = 0
	RegisterCsr15FieldCsr15Mask  = 0xffffffff
)

// GetCsr15 CSR15
func (r *registerCsr15Type) GetCsr15() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr15FieldCsr15Mask) >> RegisterCsr15FieldCsr15Shift)
}

// SetCsr15 CSR15
func (r *registerCsr15Type) SetCsr15(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr15FieldCsr15Mask)|(uint32(value)<<RegisterCsr15FieldCsr15Shift))
}

// registerCsr16Type context swap registers
type registerCsr16Type uint32

const (
	RegisterCsr16FieldCsr16Shift = 0
	RegisterCsr16FieldCsr16Mask  = 0xffffffff
)

// GetCsr16 CSR16
func (r *registerCsr16Type) GetCsr16() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr16FieldCsr16Mask) >> RegisterCsr16FieldCsr16Shift)
}

// SetCsr16 CSR16
func (r *registerCsr16Type) SetCsr16(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr16FieldCsr16Mask)|(uint32(value)<<RegisterCsr16FieldCsr16Shift))
}

// registerCsr17Type context swap registers
type registerCsr17Type uint32

const (
	RegisterCsr17FieldCsr17Shift = 0
	RegisterCsr17FieldCsr17Mask  = 0xffffffff
)

// GetCsr17 CSR17
func (r *registerCsr17Type) GetCsr17() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr17FieldCsr17Mask) >> RegisterCsr17FieldCsr17Shift)
}

// SetCsr17 CSR17
func (r *registerCsr17Type) SetCsr17(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr17FieldCsr17Mask)|(uint32(value)<<RegisterCsr17FieldCsr17Shift))
}

// registerCsr18Type context swap registers
type registerCsr18Type uint32

const (
	RegisterCsr18FieldCsr18Shift = 0
	RegisterCsr18FieldCsr18Mask  = 0xffffffff
)

// GetCsr18 CSR18
func (r *registerCsr18Type) GetCsr18() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr18FieldCsr18Mask) >> RegisterCsr18FieldCsr18Shift)
}

// SetCsr18 CSR18
func (r *registerCsr18Type) SetCsr18(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr18FieldCsr18Mask)|(uint32(value)<<RegisterCsr18FieldCsr18Shift))
}

// registerCsr19Type context swap registers
type registerCsr19Type uint32

const (
	RegisterCsr19FieldCsr19Shift = 0
	RegisterCsr19FieldCsr19Mask  = 0xffffffff
)

// GetCsr19 CSR19
func (r *registerCsr19Type) GetCsr19() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr19FieldCsr19Mask) >> RegisterCsr19FieldCsr19Shift)
}

// SetCsr19 CSR19
func (r *registerCsr19Type) SetCsr19(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr19FieldCsr19Mask)|(uint32(value)<<RegisterCsr19FieldCsr19Shift))
}

// registerCsr20Type context swap registers
type registerCsr20Type uint32

const (
	RegisterCsr20FieldCsr20Shift = 0
	RegisterCsr20FieldCsr20Mask  = 0xffffffff
)

// GetCsr20 CSR20
func (r *registerCsr20Type) GetCsr20() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr20FieldCsr20Mask) >> RegisterCsr20FieldCsr20Shift)
}

// SetCsr20 CSR20
func (r *registerCsr20Type) SetCsr20(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr20FieldCsr20Mask)|(uint32(value)<<RegisterCsr20FieldCsr20Shift))
}

// registerCsr21Type context swap registers
type registerCsr21Type uint32

const (
	RegisterCsr21FieldCsr21Shift = 0
	RegisterCsr21FieldCsr21Mask  = 0xffffffff
)

// GetCsr21 CSR21
func (r *registerCsr21Type) GetCsr21() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr21FieldCsr21Mask) >> RegisterCsr21FieldCsr21Shift)
}

// SetCsr21 CSR21
func (r *registerCsr21Type) SetCsr21(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr21FieldCsr21Mask)|(uint32(value)<<RegisterCsr21FieldCsr21Shift))
}

// registerCsr22Type context swap registers
type registerCsr22Type uint32

const (
	RegisterCsr22FieldCsr22Shift = 0
	RegisterCsr22FieldCsr22Mask  = 0xffffffff
)

// GetCsr22 CSR22
func (r *registerCsr22Type) GetCsr22() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr22FieldCsr22Mask) >> RegisterCsr22FieldCsr22Shift)
}

// SetCsr22 CSR22
func (r *registerCsr22Type) SetCsr22(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr22FieldCsr22Mask)|(uint32(value)<<RegisterCsr22FieldCsr22Shift))
}

// registerCsr23Type context swap registers
type registerCsr23Type uint32

const (
	RegisterCsr23FieldCsr23Shift = 0
	RegisterCsr23FieldCsr23Mask  = 0xffffffff
)

// GetCsr23 CSR23
func (r *registerCsr23Type) GetCsr23() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr23FieldCsr23Mask) >> RegisterCsr23FieldCsr23Shift)
}

// SetCsr23 CSR23
func (r *registerCsr23Type) SetCsr23(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr23FieldCsr23Mask)|(uint32(value)<<RegisterCsr23FieldCsr23Shift))
}

// registerCsr24Type context swap registers
type registerCsr24Type uint32

const (
	RegisterCsr24FieldCsr24Shift = 0
	RegisterCsr24FieldCsr24Mask  = 0xffffffff
)

// GetCsr24 CSR24
func (r *registerCsr24Type) GetCsr24() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr24FieldCsr24Mask) >> RegisterCsr24FieldCsr24Shift)
}

// SetCsr24 CSR24
func (r *registerCsr24Type) SetCsr24(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr24FieldCsr24Mask)|(uint32(value)<<RegisterCsr24FieldCsr24Shift))
}

// registerCsr25Type context swap registers
type registerCsr25Type uint32

const (
	RegisterCsr25FieldCsr25Shift = 0
	RegisterCsr25FieldCsr25Mask  = 0xffffffff
)

// GetCsr25 CSR25
func (r *registerCsr25Type) GetCsr25() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr25FieldCsr25Mask) >> RegisterCsr25FieldCsr25Shift)
}

// SetCsr25 CSR25
func (r *registerCsr25Type) SetCsr25(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr25FieldCsr25Mask)|(uint32(value)<<RegisterCsr25FieldCsr25Shift))
}

// registerCsr26Type context swap registers
type registerCsr26Type uint32

const (
	RegisterCsr26FieldCsr26Shift = 0
	RegisterCsr26FieldCsr26Mask  = 0xffffffff
)

// GetCsr26 CSR26
func (r *registerCsr26Type) GetCsr26() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr26FieldCsr26Mask) >> RegisterCsr26FieldCsr26Shift)
}

// SetCsr26 CSR26
func (r *registerCsr26Type) SetCsr26(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr26FieldCsr26Mask)|(uint32(value)<<RegisterCsr26FieldCsr26Shift))
}

// registerCsr27Type context swap registers
type registerCsr27Type uint32

const (
	RegisterCsr27FieldCsr27Shift = 0
	RegisterCsr27FieldCsr27Mask  = 0xffffffff
)

// GetCsr27 CSR27
func (r *registerCsr27Type) GetCsr27() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr27FieldCsr27Mask) >> RegisterCsr27FieldCsr27Shift)
}

// SetCsr27 CSR27
func (r *registerCsr27Type) SetCsr27(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr27FieldCsr27Mask)|(uint32(value)<<RegisterCsr27FieldCsr27Shift))
}

// registerCsr28Type context swap registers
type registerCsr28Type uint32

const (
	RegisterCsr28FieldCsr28Shift = 0
	RegisterCsr28FieldCsr28Mask  = 0xffffffff
)

// GetCsr28 CSR28
func (r *registerCsr28Type) GetCsr28() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr28FieldCsr28Mask) >> RegisterCsr28FieldCsr28Shift)
}

// SetCsr28 CSR28
func (r *registerCsr28Type) SetCsr28(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr28FieldCsr28Mask)|(uint32(value)<<RegisterCsr28FieldCsr28Shift))
}

// registerCsr29Type context swap registers
type registerCsr29Type uint32

const (
	RegisterCsr29FieldCsr29Shift = 0
	RegisterCsr29FieldCsr29Mask  = 0xffffffff
)

// GetCsr29 CSR29
func (r *registerCsr29Type) GetCsr29() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr29FieldCsr29Mask) >> RegisterCsr29FieldCsr29Shift)
}

// SetCsr29 CSR29
func (r *registerCsr29Type) SetCsr29(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr29FieldCsr29Mask)|(uint32(value)<<RegisterCsr29FieldCsr29Shift))
}

// registerCsr30Type context swap registers
type registerCsr30Type uint32

const (
	RegisterCsr30FieldCsr30Shift = 0
	RegisterCsr30FieldCsr30Mask  = 0xffffffff
)

// GetCsr30 CSR30
func (r *registerCsr30Type) GetCsr30() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr30FieldCsr30Mask) >> RegisterCsr30FieldCsr30Shift)
}

// SetCsr30 CSR30
func (r *registerCsr30Type) SetCsr30(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr30FieldCsr30Mask)|(uint32(value)<<RegisterCsr30FieldCsr30Shift))
}

// registerCsr31Type context swap registers
type registerCsr31Type uint32

const (
	RegisterCsr31FieldCsr31Shift = 0
	RegisterCsr31FieldCsr31Mask  = 0xffffffff
)

// GetCsr31 CSR31
func (r *registerCsr31Type) GetCsr31() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr31FieldCsr31Mask) >> RegisterCsr31FieldCsr31Shift)
}

// SetCsr31 CSR31
func (r *registerCsr31Type) SetCsr31(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr31FieldCsr31Mask)|(uint32(value)<<RegisterCsr31FieldCsr31Shift))
}

// registerCsr32Type context swap registers
type registerCsr32Type uint32

const (
	RegisterCsr32FieldCsr32Shift = 0
	RegisterCsr32FieldCsr32Mask  = 0xffffffff
)

// GetCsr32 CSR32
func (r *registerCsr32Type) GetCsr32() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr32FieldCsr32Mask) >> RegisterCsr32FieldCsr32Shift)
}

// SetCsr32 CSR32
func (r *registerCsr32Type) SetCsr32(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr32FieldCsr32Mask)|(uint32(value)<<RegisterCsr32FieldCsr32Shift))
}

// registerCsr33Type context swap registers
type registerCsr33Type uint32

const (
	RegisterCsr33FieldCsr33Shift = 0
	RegisterCsr33FieldCsr33Mask  = 0xffffffff
)

// GetCsr33 CSR33
func (r *registerCsr33Type) GetCsr33() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr33FieldCsr33Mask) >> RegisterCsr33FieldCsr33Shift)
}

// SetCsr33 CSR33
func (r *registerCsr33Type) SetCsr33(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr33FieldCsr33Mask)|(uint32(value)<<RegisterCsr33FieldCsr33Shift))
}

// registerCsr34Type context swap registers
type registerCsr34Type uint32

const (
	RegisterCsr34FieldCsr34Shift = 0
	RegisterCsr34FieldCsr34Mask  = 0xffffffff
)

// GetCsr34 CSR34
func (r *registerCsr34Type) GetCsr34() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr34FieldCsr34Mask) >> RegisterCsr34FieldCsr34Shift)
}

// SetCsr34 CSR34
func (r *registerCsr34Type) SetCsr34(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr34FieldCsr34Mask)|(uint32(value)<<RegisterCsr34FieldCsr34Shift))
}

// registerCsr35Type context swap registers
type registerCsr35Type uint32

const (
	RegisterCsr35FieldCsr35Shift = 0
	RegisterCsr35FieldCsr35Mask  = 0xffffffff
)

// GetCsr35 CSR35
func (r *registerCsr35Type) GetCsr35() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr35FieldCsr35Mask) >> RegisterCsr35FieldCsr35Shift)
}

// SetCsr35 CSR35
func (r *registerCsr35Type) SetCsr35(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr35FieldCsr35Mask)|(uint32(value)<<RegisterCsr35FieldCsr35Shift))
}

// registerCsr36Type context swap registers
type registerCsr36Type uint32

const (
	RegisterCsr36FieldCsr36Shift = 0
	RegisterCsr36FieldCsr36Mask  = 0xffffffff
)

// GetCsr36 CSR36
func (r *registerCsr36Type) GetCsr36() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr36FieldCsr36Mask) >> RegisterCsr36FieldCsr36Shift)
}

// SetCsr36 CSR36
func (r *registerCsr36Type) SetCsr36(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr36FieldCsr36Mask)|(uint32(value)<<RegisterCsr36FieldCsr36Shift))
}

// registerCsr37Type context swap registers
type registerCsr37Type uint32

const (
	RegisterCsr37FieldCsr37Shift = 0
	RegisterCsr37FieldCsr37Mask  = 0xffffffff
)

// GetCsr37 CSR37
func (r *registerCsr37Type) GetCsr37() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr37FieldCsr37Mask) >> RegisterCsr37FieldCsr37Shift)
}

// SetCsr37 CSR37
func (r *registerCsr37Type) SetCsr37(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr37FieldCsr37Mask)|(uint32(value)<<RegisterCsr37FieldCsr37Shift))
}

// registerCsr38Type context swap registers
type registerCsr38Type uint32

const (
	RegisterCsr38FieldCsr38Shift = 0
	RegisterCsr38FieldCsr38Mask  = 0xffffffff
)

// GetCsr38 CSR38
func (r *registerCsr38Type) GetCsr38() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr38FieldCsr38Mask) >> RegisterCsr38FieldCsr38Shift)
}

// SetCsr38 CSR38
func (r *registerCsr38Type) SetCsr38(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr38FieldCsr38Mask)|(uint32(value)<<RegisterCsr38FieldCsr38Shift))
}

// registerCsr39Type context swap registers
type registerCsr39Type uint32

const (
	RegisterCsr39FieldCsr39Shift = 0
	RegisterCsr39FieldCsr39Mask  = 0xffffffff
)

// GetCsr39 CSR39
func (r *registerCsr39Type) GetCsr39() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr39FieldCsr39Mask) >> RegisterCsr39FieldCsr39Shift)
}

// SetCsr39 CSR39
func (r *registerCsr39Type) SetCsr39(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr39FieldCsr39Mask)|(uint32(value)<<RegisterCsr39FieldCsr39Shift))
}

// registerCsr40Type context swap registers
type registerCsr40Type uint32

const (
	RegisterCsr40FieldCsr40Shift = 0
	RegisterCsr40FieldCsr40Mask  = 0xffffffff
)

// GetCsr40 CSR40
func (r *registerCsr40Type) GetCsr40() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr40FieldCsr40Mask) >> RegisterCsr40FieldCsr40Shift)
}

// SetCsr40 CSR40
func (r *registerCsr40Type) SetCsr40(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr40FieldCsr40Mask)|(uint32(value)<<RegisterCsr40FieldCsr40Shift))
}

// registerCsr41Type context swap registers
type registerCsr41Type uint32

const (
	RegisterCsr41FieldCsr41Shift = 0
	RegisterCsr41FieldCsr41Mask  = 0xffffffff
)

// GetCsr41 CSR41
func (r *registerCsr41Type) GetCsr41() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr41FieldCsr41Mask) >> RegisterCsr41FieldCsr41Shift)
}

// SetCsr41 CSR41
func (r *registerCsr41Type) SetCsr41(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr41FieldCsr41Mask)|(uint32(value)<<RegisterCsr41FieldCsr41Shift))
}

// registerCsr42Type context swap registers
type registerCsr42Type uint32

const (
	RegisterCsr42FieldCsr42Shift = 0
	RegisterCsr42FieldCsr42Mask  = 0xffffffff
)

// GetCsr42 CSR42
func (r *registerCsr42Type) GetCsr42() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr42FieldCsr42Mask) >> RegisterCsr42FieldCsr42Shift)
}

// SetCsr42 CSR42
func (r *registerCsr42Type) SetCsr42(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr42FieldCsr42Mask)|(uint32(value)<<RegisterCsr42FieldCsr42Shift))
}

// registerCsr43Type context swap registers
type registerCsr43Type uint32

const (
	RegisterCsr43FieldCsr43Shift = 0
	RegisterCsr43FieldCsr43Mask  = 0xffffffff
)

// GetCsr43 CSR43
func (r *registerCsr43Type) GetCsr43() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr43FieldCsr43Mask) >> RegisterCsr43FieldCsr43Shift)
}

// SetCsr43 CSR43
func (r *registerCsr43Type) SetCsr43(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr43FieldCsr43Mask)|(uint32(value)<<RegisterCsr43FieldCsr43Shift))
}

// registerCsr44Type context swap registers
type registerCsr44Type uint32

const (
	RegisterCsr44FieldCsr44Shift = 0
	RegisterCsr44FieldCsr44Mask  = 0xffffffff
)

// GetCsr44 CSR44
func (r *registerCsr44Type) GetCsr44() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr44FieldCsr44Mask) >> RegisterCsr44FieldCsr44Shift)
}

// SetCsr44 CSR44
func (r *registerCsr44Type) SetCsr44(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr44FieldCsr44Mask)|(uint32(value)<<RegisterCsr44FieldCsr44Shift))
}

// registerCsr45Type context swap registers
type registerCsr45Type uint32

const (
	RegisterCsr45FieldCsr45Shift = 0
	RegisterCsr45FieldCsr45Mask  = 0xffffffff
)

// GetCsr45 CSR45
func (r *registerCsr45Type) GetCsr45() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr45FieldCsr45Mask) >> RegisterCsr45FieldCsr45Shift)
}

// SetCsr45 CSR45
func (r *registerCsr45Type) SetCsr45(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr45FieldCsr45Mask)|(uint32(value)<<RegisterCsr45FieldCsr45Shift))
}

// registerCsr46Type context swap registers
type registerCsr46Type uint32

const (
	RegisterCsr46FieldCsr46Shift = 0
	RegisterCsr46FieldCsr46Mask  = 0xffffffff
)

// GetCsr46 CSR46
func (r *registerCsr46Type) GetCsr46() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr46FieldCsr46Mask) >> RegisterCsr46FieldCsr46Shift)
}

// SetCsr46 CSR46
func (r *registerCsr46Type) SetCsr46(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr46FieldCsr46Mask)|(uint32(value)<<RegisterCsr46FieldCsr46Shift))
}

// registerCsr47Type context swap registers
type registerCsr47Type uint32

const (
	RegisterCsr47FieldCsr47Shift = 0
	RegisterCsr47FieldCsr47Mask  = 0xffffffff
)

// GetCsr47 CSR47
func (r *registerCsr47Type) GetCsr47() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr47FieldCsr47Mask) >> RegisterCsr47FieldCsr47Shift)
}

// SetCsr47 CSR47
func (r *registerCsr47Type) SetCsr47(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr47FieldCsr47Mask)|(uint32(value)<<RegisterCsr47FieldCsr47Shift))
}

// registerCsr48Type context swap registers
type registerCsr48Type uint32

const (
	RegisterCsr48FieldCsr48Shift = 0
	RegisterCsr48FieldCsr48Mask  = 0xffffffff
)

// GetCsr48 CSR48
func (r *registerCsr48Type) GetCsr48() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr48FieldCsr48Mask) >> RegisterCsr48FieldCsr48Shift)
}

// SetCsr48 CSR48
func (r *registerCsr48Type) SetCsr48(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr48FieldCsr48Mask)|(uint32(value)<<RegisterCsr48FieldCsr48Shift))
}

// registerCsr49Type context swap registers
type registerCsr49Type uint32

const (
	RegisterCsr49FieldCsr49Shift = 0
	RegisterCsr49FieldCsr49Mask  = 0xffffffff
)

// GetCsr49 CSR49
func (r *registerCsr49Type) GetCsr49() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr49FieldCsr49Mask) >> RegisterCsr49FieldCsr49Shift)
}

// SetCsr49 CSR49
func (r *registerCsr49Type) SetCsr49(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr49FieldCsr49Mask)|(uint32(value)<<RegisterCsr49FieldCsr49Shift))
}

// registerCsr50Type context swap registers
type registerCsr50Type uint32

const (
	RegisterCsr50FieldCsr50Shift = 0
	RegisterCsr50FieldCsr50Mask  = 0xffffffff
)

// GetCsr50 CSR50
func (r *registerCsr50Type) GetCsr50() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr50FieldCsr50Mask) >> RegisterCsr50FieldCsr50Shift)
}

// SetCsr50 CSR50
func (r *registerCsr50Type) SetCsr50(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr50FieldCsr50Mask)|(uint32(value)<<RegisterCsr50FieldCsr50Shift))
}

// registerCsr51Type context swap registers
type registerCsr51Type uint32

const (
	RegisterCsr51FieldCsr51Shift = 0
	RegisterCsr51FieldCsr51Mask  = 0xffffffff
)

// GetCsr51 CSR51
func (r *registerCsr51Type) GetCsr51() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr51FieldCsr51Mask) >> RegisterCsr51FieldCsr51Shift)
}

// SetCsr51 CSR51
func (r *registerCsr51Type) SetCsr51(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr51FieldCsr51Mask)|(uint32(value)<<RegisterCsr51FieldCsr51Shift))
}

// registerCsr52Type context swap registers
type registerCsr52Type uint32

const (
	RegisterCsr52FieldCsr52Shift = 0
	RegisterCsr52FieldCsr52Mask  = 0xffffffff
)

// GetCsr52 CSR52
func (r *registerCsr52Type) GetCsr52() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr52FieldCsr52Mask) >> RegisterCsr52FieldCsr52Shift)
}

// SetCsr52 CSR52
func (r *registerCsr52Type) SetCsr52(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr52FieldCsr52Mask)|(uint32(value)<<RegisterCsr52FieldCsr52Shift))
}

// registerCsr53Type context swap registers
type registerCsr53Type uint32

const (
	RegisterCsr53FieldCsr53Shift = 0
	RegisterCsr53FieldCsr53Mask  = 0xffffffff
)

// GetCsr53 CSR53
func (r *registerCsr53Type) GetCsr53() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr53FieldCsr53Mask) >> RegisterCsr53FieldCsr53Shift)
}

// SetCsr53 CSR53
func (r *registerCsr53Type) SetCsr53(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr53FieldCsr53Mask)|(uint32(value)<<RegisterCsr53FieldCsr53Shift))
}

// registerHash_hr0Type HASH digest register
type registerHash_hr0Type uint32

const (
	RegisterHash_hr0FieldH0Shift = 0
	RegisterHash_hr0FieldH0Mask  = 0xffffffff
)

// GetH0 H0
func (r *registerHash_hr0Type) GetH0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHash_hr0FieldH0Mask) >> RegisterHash_hr0FieldH0Shift)
}

// SetH0 H0
func (r *registerHash_hr0Type) SetH0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHash_hr0FieldH0Mask)|(uint32(value)<<RegisterHash_hr0FieldH0Shift))
}

// registerHash_hr1Type read-only
type registerHash_hr1Type uint32

const (
	RegisterHash_hr1FieldH1Shift = 0
	RegisterHash_hr1FieldH1Mask  = 0xffffffff
)

// GetH1 H1
func (r *registerHash_hr1Type) GetH1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHash_hr1FieldH1Mask) >> RegisterHash_hr1FieldH1Shift)
}

// SetH1 H1
func (r *registerHash_hr1Type) SetH1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHash_hr1FieldH1Mask)|(uint32(value)<<RegisterHash_hr1FieldH1Shift))
}

// registerHash_hr2Type read-only
type registerHash_hr2Type uint32

const (
	RegisterHash_hr2FieldH2Shift = 0
	RegisterHash_hr2FieldH2Mask  = 0xffffffff
)

// GetH2 H2
func (r *registerHash_hr2Type) GetH2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHash_hr2FieldH2Mask) >> RegisterHash_hr2FieldH2Shift)
}

// SetH2 H2
func (r *registerHash_hr2Type) SetH2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHash_hr2FieldH2Mask)|(uint32(value)<<RegisterHash_hr2FieldH2Shift))
}

// registerHash_hr3Type read-only
type registerHash_hr3Type uint32

const (
	RegisterHash_hr3FieldH3Shift = 0
	RegisterHash_hr3FieldH3Mask  = 0xffffffff
)

// GetH3 H3
func (r *registerHash_hr3Type) GetH3() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHash_hr3FieldH3Mask) >> RegisterHash_hr3FieldH3Shift)
}

// SetH3 H3
func (r *registerHash_hr3Type) SetH3(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHash_hr3FieldH3Mask)|(uint32(value)<<RegisterHash_hr3FieldH3Shift))
}

// registerHash_hr4Type read-only
type registerHash_hr4Type uint32

const (
	RegisterHash_hr4FieldH4Shift = 0
	RegisterHash_hr4FieldH4Mask  = 0xffffffff
)

// GetH4 H4
func (r *registerHash_hr4Type) GetH4() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHash_hr4FieldH4Mask) >> RegisterHash_hr4FieldH4Shift)
}

// SetH4 H4
func (r *registerHash_hr4Type) SetH4(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHash_hr4FieldH4Mask)|(uint32(value)<<RegisterHash_hr4FieldH4Shift))
}

// registerHash_hr5Type read-only
type registerHash_hr5Type uint32

const (
	RegisterHash_hr5FieldH5Shift = 0
	RegisterHash_hr5FieldH5Mask  = 0xffffffff
)

// GetH5 H5
func (r *registerHash_hr5Type) GetH5() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHash_hr5FieldH5Mask) >> RegisterHash_hr5FieldH5Shift)
}

// SetH5 H5
func (r *registerHash_hr5Type) SetH5(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHash_hr5FieldH5Mask)|(uint32(value)<<RegisterHash_hr5FieldH5Shift))
}

// registerHash_hr6Type read-only
type registerHash_hr6Type uint32

const (
	RegisterHash_hr6FieldH6Shift = 0
	RegisterHash_hr6FieldH6Mask  = 0xffffffff
)

// GetH6 H6
func (r *registerHash_hr6Type) GetH6() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHash_hr6FieldH6Mask) >> RegisterHash_hr6FieldH6Shift)
}

// SetH6 H6
func (r *registerHash_hr6Type) SetH6(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHash_hr6FieldH6Mask)|(uint32(value)<<RegisterHash_hr6FieldH6Shift))
}

// registerHash_hr7Type read-only
type registerHash_hr7Type uint32

const (
	RegisterHash_hr7FieldH7Shift = 0
	RegisterHash_hr7FieldH7Mask  = 0xffffffff
)

// GetH7 H7
func (r *registerHash_hr7Type) GetH7() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHash_hr7FieldH7Mask) >> RegisterHash_hr7FieldH7Shift)
}

// SetH7 H7
func (r *registerHash_hr7Type) SetH7(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHash_hr7FieldH7Mask)|(uint32(value)<<RegisterHash_hr7FieldH7Shift))
}
