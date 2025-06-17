//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package hash

import (
	"unsafe"
	"volatile"
)

var (
	Hash = (*_hash)(unsafe.Pointer(uintptr(0x48021400)))
)

type _hash struct {
	Cr      RegisterCrType
	Din     RegisterDinType
	Str     RegisterStrType
	Hr0     RegisterHr0Type
	Hr1     RegisterHr1Type
	Hr2     RegisterHr2Type
	Hr3     RegisterHr3Type
	Hr4     RegisterHr4Type
	Imr     RegisterImrType
	Sr      RegisterSrType
	_       [208]byte
	Csr0    RegisterCsr0Type
	Csr1    RegisterCsr1Type
	Csr2    RegisterCsr2Type
	Csr3    RegisterCsr3Type
	Csr4    RegisterCsr4Type
	Csr5    RegisterCsr5Type
	Csr6    RegisterCsr6Type
	Csr7    RegisterCsr7Type
	Csr8    RegisterCsr8Type
	Csr9    RegisterCsr9Type
	Csr10   RegisterCsr10Type
	Csr11   RegisterCsr11Type
	Csr12   RegisterCsr12Type
	Csr13   RegisterCsr13Type
	Csr14   RegisterCsr14Type
	Csr15   RegisterCsr15Type
	Csr16   RegisterCsr16Type
	Csr17   RegisterCsr17Type
	Csr18   RegisterCsr18Type
	Csr19   RegisterCsr19Type
	Csr20   RegisterCsr20Type
	Csr21   RegisterCsr21Type
	Csr22   RegisterCsr22Type
	Csr23   RegisterCsr23Type
	Csr24   RegisterCsr24Type
	Csr25   RegisterCsr25Type
	Csr26   RegisterCsr26Type
	Csr27   RegisterCsr27Type
	Csr28   RegisterCsr28Type
	Csr29   RegisterCsr29Type
	Csr30   RegisterCsr30Type
	Csr31   RegisterCsr31Type
	Csr32   RegisterCsr32Type
	Csr33   RegisterCsr33Type
	Csr34   RegisterCsr34Type
	Csr35   RegisterCsr35Type
	Csr36   RegisterCsr36Type
	Csr37   RegisterCsr37Type
	Csr38   RegisterCsr38Type
	Csr39   RegisterCsr39Type
	Csr40   RegisterCsr40Type
	Csr41   RegisterCsr41Type
	Csr42   RegisterCsr42Type
	Csr43   RegisterCsr43Type
	Csr44   RegisterCsr44Type
	Csr45   RegisterCsr45Type
	Csr46   RegisterCsr46Type
	Csr47   RegisterCsr47Type
	Csr48   RegisterCsr48Type
	Csr49   RegisterCsr49Type
	Csr50   RegisterCsr50Type
	Csr51   RegisterCsr51Type
	Csr52   RegisterCsr52Type
	Csr53   RegisterCsr53Type
	_       [320]byte
	Hashhr0 RegisterHashhr0Type
	Hashhr1 RegisterHashhr1Type
	Hashhr2 RegisterHashhr2Type
	Hashhr3 RegisterHashhr3Type
	Hashhr4 RegisterHashhr4Type
	Hashhr5 RegisterHashhr5Type
	Hashhr6 RegisterHashhr6Type
	Hashhr7 RegisterHashhr7Type
}

// RegisterCrType control register
type RegisterCrType uint32

func (r *RegisterCrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrFieldInitShift = 2
	RegisterCrFieldInitMask  = 0x4
)

// SetInit Initialize message digest calculation
func (r *RegisterCrType) SetInit(value bool) {
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
func (r *RegisterCrType) GetDmae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDmaeMask) != 0
}

// SetDmae DMA enable
func (r *RegisterCrType) SetDmae(value bool) {
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
func (r *RegisterCrType) GetDatatype() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDatatypeMask) >> RegisterCrFieldDatatypeShift)
}

// SetDatatype Data type selection
func (r *RegisterCrType) SetDatatype(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDatatypeMask)|(uint32(value)<<RegisterCrFieldDatatypeShift))
}

const (
	RegisterCrFieldModeShift = 6
	RegisterCrFieldModeMask  = 0x40
)

// GetMode Mode selection
func (r *RegisterCrType) GetMode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldModeMask) != 0
}

// SetMode Mode selection
func (r *RegisterCrType) SetMode(value bool) {
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
func (r *RegisterCrType) GetAlgo0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAlgo0Mask) != 0
}

// SetAlgo0 Algorithm selection
func (r *RegisterCrType) SetAlgo0(value bool) {
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
func (r *RegisterCrType) GetNbw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldNbwMask) >> RegisterCrFieldNbwShift)
}

const (
	RegisterCrFieldDinneShift = 12
	RegisterCrFieldDinneMask  = 0x1000
)

// GetDinne DIN not empty
func (r *RegisterCrType) GetDinne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDinneMask) != 0
}

const (
	RegisterCrFieldMdmatShift = 13
	RegisterCrFieldMdmatMask  = 0x2000
)

// GetMdmat Multiple DMA Transfers
func (r *RegisterCrType) GetMdmat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldMdmatMask) != 0
}

// SetMdmat Multiple DMA Transfers
func (r *RegisterCrType) SetMdmat(value bool) {
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
func (r *RegisterCrType) GetLkey() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLkeyMask) != 0
}

// SetLkey Long key selection
func (r *RegisterCrType) SetLkey(value bool) {
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
func (r *RegisterCrType) GetAlgo1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAlgo1Mask) != 0
}

// SetAlgo1 ALGO
func (r *RegisterCrType) SetAlgo1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAlgo1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAlgo1Mask)
	}
}

// RegisterDinType data input register
type RegisterDinType uint32

func (r *RegisterDinType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDinType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDinType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDinType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDinType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDinFieldDatainShift = 0
	RegisterDinFieldDatainMask  = 0xffffffff
)

// GetDatain Data input
func (r *RegisterDinType) GetDatain() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDinFieldDatainMask) >> RegisterDinFieldDatainShift)
}

// SetDatain Data input
func (r *RegisterDinType) SetDatain(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinFieldDatainMask)|(uint32(value)<<RegisterDinFieldDatainShift))
}

// RegisterStrType start register
type RegisterStrType uint32

func (r *RegisterStrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterStrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterStrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterStrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterStrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterStrFieldNblwShift = 0
	RegisterStrFieldNblwMask  = 0x1f
)

// GetNblw Number of valid bits in the last word of the message
func (r *RegisterStrType) GetNblw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterStrFieldNblwMask) >> RegisterStrFieldNblwShift)
}

// SetNblw Number of valid bits in the last word of the message
func (r *RegisterStrType) SetNblw(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterStrFieldNblwMask)|(uint32(value)<<RegisterStrFieldNblwShift))
}

const (
	RegisterStrFieldDcalShift = 8
	RegisterStrFieldDcalMask  = 0x100
)

// SetDcal Digest calculation
func (r *RegisterStrType) SetDcal(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStrFieldDcalMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStrFieldDcalMask)
	}
}

// RegisterHr0Type digest registers
type RegisterHr0Type uint32

func (r *RegisterHr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHr0FieldH0Shift = 0
	RegisterHr0FieldH0Mask  = 0xffffffff
)

// GetH0 H0
func (r *RegisterHr0Type) GetH0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHr0FieldH0Mask) >> RegisterHr0FieldH0Shift)
}

// SetH0 H0
func (r *RegisterHr0Type) SetH0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHr0FieldH0Mask)|(uint32(value)<<RegisterHr0FieldH0Shift))
}

// RegisterHr1Type digest registers
type RegisterHr1Type uint32

func (r *RegisterHr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHr1FieldH1Shift = 0
	RegisterHr1FieldH1Mask  = 0xffffffff
)

// GetH1 H1
func (r *RegisterHr1Type) GetH1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHr1FieldH1Mask) >> RegisterHr1FieldH1Shift)
}

// SetH1 H1
func (r *RegisterHr1Type) SetH1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHr1FieldH1Mask)|(uint32(value)<<RegisterHr1FieldH1Shift))
}

// RegisterHr2Type digest registers
type RegisterHr2Type uint32

func (r *RegisterHr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHr2FieldH2Shift = 0
	RegisterHr2FieldH2Mask  = 0xffffffff
)

// GetH2 H2
func (r *RegisterHr2Type) GetH2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHr2FieldH2Mask) >> RegisterHr2FieldH2Shift)
}

// SetH2 H2
func (r *RegisterHr2Type) SetH2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHr2FieldH2Mask)|(uint32(value)<<RegisterHr2FieldH2Shift))
}

// RegisterHr3Type digest registers
type RegisterHr3Type uint32

func (r *RegisterHr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHr3FieldH3Shift = 0
	RegisterHr3FieldH3Mask  = 0xffffffff
)

// GetH3 H3
func (r *RegisterHr3Type) GetH3() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHr3FieldH3Mask) >> RegisterHr3FieldH3Shift)
}

// SetH3 H3
func (r *RegisterHr3Type) SetH3(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHr3FieldH3Mask)|(uint32(value)<<RegisterHr3FieldH3Shift))
}

// RegisterHr4Type digest registers
type RegisterHr4Type uint32

func (r *RegisterHr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHr4FieldH4Shift = 0
	RegisterHr4FieldH4Mask  = 0xffffffff
)

// GetH4 H4
func (r *RegisterHr4Type) GetH4() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHr4FieldH4Mask) >> RegisterHr4FieldH4Shift)
}

// SetH4 H4
func (r *RegisterHr4Type) SetH4(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHr4FieldH4Mask)|(uint32(value)<<RegisterHr4FieldH4Shift))
}

// RegisterImrType interrupt enable register
type RegisterImrType uint32

func (r *RegisterImrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterImrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterImrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterImrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterImrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterImrFieldDinieShift = 0
	RegisterImrFieldDinieMask  = 0x1
)

// GetDinie Data input interrupt enable
func (r *RegisterImrType) GetDinie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldDinieMask) != 0
}

// SetDinie Data input interrupt enable
func (r *RegisterImrType) SetDinie(value bool) {
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
func (r *RegisterImrType) GetDcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldDcieMask) != 0
}

// SetDcie Digest calculation completion interrupt enable
func (r *RegisterImrType) SetDcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImrFieldDcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImrFieldDcieMask)
	}
}

// RegisterSrType status register
type RegisterSrType uint32

func (r *RegisterSrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSrFieldDinisShift = 0
	RegisterSrFieldDinisMask  = 0x1
)

// GetDinis Data input interrupt status
func (r *RegisterSrType) GetDinis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDinisMask) != 0
}

// SetDinis Data input interrupt status
func (r *RegisterSrType) SetDinis(value bool) {
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
func (r *RegisterSrType) GetDcis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDcisMask) != 0
}

// SetDcis Digest calculation completion interrupt status
func (r *RegisterSrType) SetDcis(value bool) {
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
func (r *RegisterSrType) GetDmas() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDmasMask) != 0
}

const (
	RegisterSrFieldBusyShift = 3
	RegisterSrFieldBusyMask  = 0x8
)

// GetBusy Busy bit
func (r *RegisterSrType) GetBusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldBusyMask) != 0
}

// RegisterCsr0Type context swap registers
type RegisterCsr0Type uint32

func (r *RegisterCsr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr0FieldCsr0Shift = 0
	RegisterCsr0FieldCsr0Mask  = 0xffffffff
)

// GetCsr0 CSR0
func (r *RegisterCsr0Type) GetCsr0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr0FieldCsr0Mask) >> RegisterCsr0FieldCsr0Shift)
}

// SetCsr0 CSR0
func (r *RegisterCsr0Type) SetCsr0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr0FieldCsr0Mask)|(uint32(value)<<RegisterCsr0FieldCsr0Shift))
}

// RegisterCsr1Type context swap registers
type RegisterCsr1Type uint32

func (r *RegisterCsr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr1FieldCsr1Shift = 0
	RegisterCsr1FieldCsr1Mask  = 0xffffffff
)

// GetCsr1 CSR1
func (r *RegisterCsr1Type) GetCsr1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr1FieldCsr1Mask) >> RegisterCsr1FieldCsr1Shift)
}

// SetCsr1 CSR1
func (r *RegisterCsr1Type) SetCsr1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr1FieldCsr1Mask)|(uint32(value)<<RegisterCsr1FieldCsr1Shift))
}

// RegisterCsr2Type context swap registers
type RegisterCsr2Type uint32

func (r *RegisterCsr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr2FieldCsr2Shift = 0
	RegisterCsr2FieldCsr2Mask  = 0xffffffff
)

// GetCsr2 CSR2
func (r *RegisterCsr2Type) GetCsr2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr2FieldCsr2Mask) >> RegisterCsr2FieldCsr2Shift)
}

// SetCsr2 CSR2
func (r *RegisterCsr2Type) SetCsr2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr2FieldCsr2Mask)|(uint32(value)<<RegisterCsr2FieldCsr2Shift))
}

// RegisterCsr3Type context swap registers
type RegisterCsr3Type uint32

func (r *RegisterCsr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr3FieldCsr3Shift = 0
	RegisterCsr3FieldCsr3Mask  = 0xffffffff
)

// GetCsr3 CSR3
func (r *RegisterCsr3Type) GetCsr3() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr3FieldCsr3Mask) >> RegisterCsr3FieldCsr3Shift)
}

// SetCsr3 CSR3
func (r *RegisterCsr3Type) SetCsr3(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr3FieldCsr3Mask)|(uint32(value)<<RegisterCsr3FieldCsr3Shift))
}

// RegisterCsr4Type context swap registers
type RegisterCsr4Type uint32

func (r *RegisterCsr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr4FieldCsr4Shift = 0
	RegisterCsr4FieldCsr4Mask  = 0xffffffff
)

// GetCsr4 CSR4
func (r *RegisterCsr4Type) GetCsr4() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr4FieldCsr4Mask) >> RegisterCsr4FieldCsr4Shift)
}

// SetCsr4 CSR4
func (r *RegisterCsr4Type) SetCsr4(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr4FieldCsr4Mask)|(uint32(value)<<RegisterCsr4FieldCsr4Shift))
}

// RegisterCsr5Type context swap registers
type RegisterCsr5Type uint32

func (r *RegisterCsr5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr5FieldCsr5Shift = 0
	RegisterCsr5FieldCsr5Mask  = 0xffffffff
)

// GetCsr5 CSR5
func (r *RegisterCsr5Type) GetCsr5() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr5FieldCsr5Mask) >> RegisterCsr5FieldCsr5Shift)
}

// SetCsr5 CSR5
func (r *RegisterCsr5Type) SetCsr5(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr5FieldCsr5Mask)|(uint32(value)<<RegisterCsr5FieldCsr5Shift))
}

// RegisterCsr6Type context swap registers
type RegisterCsr6Type uint32

func (r *RegisterCsr6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr6FieldCsr6Shift = 0
	RegisterCsr6FieldCsr6Mask  = 0xffffffff
)

// GetCsr6 CSR6
func (r *RegisterCsr6Type) GetCsr6() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr6FieldCsr6Mask) >> RegisterCsr6FieldCsr6Shift)
}

// SetCsr6 CSR6
func (r *RegisterCsr6Type) SetCsr6(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr6FieldCsr6Mask)|(uint32(value)<<RegisterCsr6FieldCsr6Shift))
}

// RegisterCsr7Type context swap registers
type RegisterCsr7Type uint32

func (r *RegisterCsr7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr7FieldCsr7Shift = 0
	RegisterCsr7FieldCsr7Mask  = 0xffffffff
)

// GetCsr7 CSR7
func (r *RegisterCsr7Type) GetCsr7() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr7FieldCsr7Mask) >> RegisterCsr7FieldCsr7Shift)
}

// SetCsr7 CSR7
func (r *RegisterCsr7Type) SetCsr7(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr7FieldCsr7Mask)|(uint32(value)<<RegisterCsr7FieldCsr7Shift))
}

// RegisterCsr8Type context swap registers
type RegisterCsr8Type uint32

func (r *RegisterCsr8Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr8Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr8Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr8Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr8Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr8FieldCsr8Shift = 0
	RegisterCsr8FieldCsr8Mask  = 0xffffffff
)

// GetCsr8 CSR8
func (r *RegisterCsr8Type) GetCsr8() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr8FieldCsr8Mask) >> RegisterCsr8FieldCsr8Shift)
}

// SetCsr8 CSR8
func (r *RegisterCsr8Type) SetCsr8(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr8FieldCsr8Mask)|(uint32(value)<<RegisterCsr8FieldCsr8Shift))
}

// RegisterCsr9Type context swap registers
type RegisterCsr9Type uint32

func (r *RegisterCsr9Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr9Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr9Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr9Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr9Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr9FieldCsr9Shift = 0
	RegisterCsr9FieldCsr9Mask  = 0xffffffff
)

// GetCsr9 CSR9
func (r *RegisterCsr9Type) GetCsr9() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr9FieldCsr9Mask) >> RegisterCsr9FieldCsr9Shift)
}

// SetCsr9 CSR9
func (r *RegisterCsr9Type) SetCsr9(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr9FieldCsr9Mask)|(uint32(value)<<RegisterCsr9FieldCsr9Shift))
}

// RegisterCsr10Type context swap registers
type RegisterCsr10Type uint32

func (r *RegisterCsr10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr10FieldCsr10Shift = 0
	RegisterCsr10FieldCsr10Mask  = 0xffffffff
)

// GetCsr10 CSR10
func (r *RegisterCsr10Type) GetCsr10() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr10FieldCsr10Mask) >> RegisterCsr10FieldCsr10Shift)
}

// SetCsr10 CSR10
func (r *RegisterCsr10Type) SetCsr10(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr10FieldCsr10Mask)|(uint32(value)<<RegisterCsr10FieldCsr10Shift))
}

// RegisterCsr11Type context swap registers
type RegisterCsr11Type uint32

func (r *RegisterCsr11Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr11Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr11Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr11Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr11Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr11FieldCsr11Shift = 0
	RegisterCsr11FieldCsr11Mask  = 0xffffffff
)

// GetCsr11 CSR11
func (r *RegisterCsr11Type) GetCsr11() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr11FieldCsr11Mask) >> RegisterCsr11FieldCsr11Shift)
}

// SetCsr11 CSR11
func (r *RegisterCsr11Type) SetCsr11(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr11FieldCsr11Mask)|(uint32(value)<<RegisterCsr11FieldCsr11Shift))
}

// RegisterCsr12Type context swap registers
type RegisterCsr12Type uint32

func (r *RegisterCsr12Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr12Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr12Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr12Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr12Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr12FieldCsr12Shift = 0
	RegisterCsr12FieldCsr12Mask  = 0xffffffff
)

// GetCsr12 CSR12
func (r *RegisterCsr12Type) GetCsr12() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr12FieldCsr12Mask) >> RegisterCsr12FieldCsr12Shift)
}

// SetCsr12 CSR12
func (r *RegisterCsr12Type) SetCsr12(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr12FieldCsr12Mask)|(uint32(value)<<RegisterCsr12FieldCsr12Shift))
}

// RegisterCsr13Type context swap registers
type RegisterCsr13Type uint32

func (r *RegisterCsr13Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr13Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr13Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr13Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr13Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr13FieldCsr13Shift = 0
	RegisterCsr13FieldCsr13Mask  = 0xffffffff
)

// GetCsr13 CSR13
func (r *RegisterCsr13Type) GetCsr13() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr13FieldCsr13Mask) >> RegisterCsr13FieldCsr13Shift)
}

// SetCsr13 CSR13
func (r *RegisterCsr13Type) SetCsr13(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr13FieldCsr13Mask)|(uint32(value)<<RegisterCsr13FieldCsr13Shift))
}

// RegisterCsr14Type context swap registers
type RegisterCsr14Type uint32

func (r *RegisterCsr14Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr14Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr14Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr14Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr14Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr14FieldCsr14Shift = 0
	RegisterCsr14FieldCsr14Mask  = 0xffffffff
)

// GetCsr14 CSR14
func (r *RegisterCsr14Type) GetCsr14() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr14FieldCsr14Mask) >> RegisterCsr14FieldCsr14Shift)
}

// SetCsr14 CSR14
func (r *RegisterCsr14Type) SetCsr14(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr14FieldCsr14Mask)|(uint32(value)<<RegisterCsr14FieldCsr14Shift))
}

// RegisterCsr15Type context swap registers
type RegisterCsr15Type uint32

func (r *RegisterCsr15Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr15Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr15Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr15Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr15Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr15FieldCsr15Shift = 0
	RegisterCsr15FieldCsr15Mask  = 0xffffffff
)

// GetCsr15 CSR15
func (r *RegisterCsr15Type) GetCsr15() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr15FieldCsr15Mask) >> RegisterCsr15FieldCsr15Shift)
}

// SetCsr15 CSR15
func (r *RegisterCsr15Type) SetCsr15(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr15FieldCsr15Mask)|(uint32(value)<<RegisterCsr15FieldCsr15Shift))
}

// RegisterCsr16Type context swap registers
type RegisterCsr16Type uint32

func (r *RegisterCsr16Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr16Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr16Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr16Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr16Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr16FieldCsr16Shift = 0
	RegisterCsr16FieldCsr16Mask  = 0xffffffff
)

// GetCsr16 CSR16
func (r *RegisterCsr16Type) GetCsr16() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr16FieldCsr16Mask) >> RegisterCsr16FieldCsr16Shift)
}

// SetCsr16 CSR16
func (r *RegisterCsr16Type) SetCsr16(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr16FieldCsr16Mask)|(uint32(value)<<RegisterCsr16FieldCsr16Shift))
}

// RegisterCsr17Type context swap registers
type RegisterCsr17Type uint32

func (r *RegisterCsr17Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr17Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr17Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr17Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr17Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr17FieldCsr17Shift = 0
	RegisterCsr17FieldCsr17Mask  = 0xffffffff
)

// GetCsr17 CSR17
func (r *RegisterCsr17Type) GetCsr17() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr17FieldCsr17Mask) >> RegisterCsr17FieldCsr17Shift)
}

// SetCsr17 CSR17
func (r *RegisterCsr17Type) SetCsr17(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr17FieldCsr17Mask)|(uint32(value)<<RegisterCsr17FieldCsr17Shift))
}

// RegisterCsr18Type context swap registers
type RegisterCsr18Type uint32

func (r *RegisterCsr18Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr18Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr18Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr18Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr18Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr18FieldCsr18Shift = 0
	RegisterCsr18FieldCsr18Mask  = 0xffffffff
)

// GetCsr18 CSR18
func (r *RegisterCsr18Type) GetCsr18() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr18FieldCsr18Mask) >> RegisterCsr18FieldCsr18Shift)
}

// SetCsr18 CSR18
func (r *RegisterCsr18Type) SetCsr18(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr18FieldCsr18Mask)|(uint32(value)<<RegisterCsr18FieldCsr18Shift))
}

// RegisterCsr19Type context swap registers
type RegisterCsr19Type uint32

func (r *RegisterCsr19Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr19Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr19Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr19Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr19Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr19FieldCsr19Shift = 0
	RegisterCsr19FieldCsr19Mask  = 0xffffffff
)

// GetCsr19 CSR19
func (r *RegisterCsr19Type) GetCsr19() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr19FieldCsr19Mask) >> RegisterCsr19FieldCsr19Shift)
}

// SetCsr19 CSR19
func (r *RegisterCsr19Type) SetCsr19(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr19FieldCsr19Mask)|(uint32(value)<<RegisterCsr19FieldCsr19Shift))
}

// RegisterCsr20Type context swap registers
type RegisterCsr20Type uint32

func (r *RegisterCsr20Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr20Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr20Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr20Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr20Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr20FieldCsr20Shift = 0
	RegisterCsr20FieldCsr20Mask  = 0xffffffff
)

// GetCsr20 CSR20
func (r *RegisterCsr20Type) GetCsr20() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr20FieldCsr20Mask) >> RegisterCsr20FieldCsr20Shift)
}

// SetCsr20 CSR20
func (r *RegisterCsr20Type) SetCsr20(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr20FieldCsr20Mask)|(uint32(value)<<RegisterCsr20FieldCsr20Shift))
}

// RegisterCsr21Type context swap registers
type RegisterCsr21Type uint32

func (r *RegisterCsr21Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr21Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr21Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr21Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr21Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr21FieldCsr21Shift = 0
	RegisterCsr21FieldCsr21Mask  = 0xffffffff
)

// GetCsr21 CSR21
func (r *RegisterCsr21Type) GetCsr21() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr21FieldCsr21Mask) >> RegisterCsr21FieldCsr21Shift)
}

// SetCsr21 CSR21
func (r *RegisterCsr21Type) SetCsr21(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr21FieldCsr21Mask)|(uint32(value)<<RegisterCsr21FieldCsr21Shift))
}

// RegisterCsr22Type context swap registers
type RegisterCsr22Type uint32

func (r *RegisterCsr22Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr22Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr22Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr22Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr22Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr22FieldCsr22Shift = 0
	RegisterCsr22FieldCsr22Mask  = 0xffffffff
)

// GetCsr22 CSR22
func (r *RegisterCsr22Type) GetCsr22() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr22FieldCsr22Mask) >> RegisterCsr22FieldCsr22Shift)
}

// SetCsr22 CSR22
func (r *RegisterCsr22Type) SetCsr22(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr22FieldCsr22Mask)|(uint32(value)<<RegisterCsr22FieldCsr22Shift))
}

// RegisterCsr23Type context swap registers
type RegisterCsr23Type uint32

func (r *RegisterCsr23Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr23Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr23Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr23Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr23Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr23FieldCsr23Shift = 0
	RegisterCsr23FieldCsr23Mask  = 0xffffffff
)

// GetCsr23 CSR23
func (r *RegisterCsr23Type) GetCsr23() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr23FieldCsr23Mask) >> RegisterCsr23FieldCsr23Shift)
}

// SetCsr23 CSR23
func (r *RegisterCsr23Type) SetCsr23(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr23FieldCsr23Mask)|(uint32(value)<<RegisterCsr23FieldCsr23Shift))
}

// RegisterCsr24Type context swap registers
type RegisterCsr24Type uint32

func (r *RegisterCsr24Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr24Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr24Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr24Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr24Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr24FieldCsr24Shift = 0
	RegisterCsr24FieldCsr24Mask  = 0xffffffff
)

// GetCsr24 CSR24
func (r *RegisterCsr24Type) GetCsr24() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr24FieldCsr24Mask) >> RegisterCsr24FieldCsr24Shift)
}

// SetCsr24 CSR24
func (r *RegisterCsr24Type) SetCsr24(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr24FieldCsr24Mask)|(uint32(value)<<RegisterCsr24FieldCsr24Shift))
}

// RegisterCsr25Type context swap registers
type RegisterCsr25Type uint32

func (r *RegisterCsr25Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr25Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr25Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr25Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr25Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr25FieldCsr25Shift = 0
	RegisterCsr25FieldCsr25Mask  = 0xffffffff
)

// GetCsr25 CSR25
func (r *RegisterCsr25Type) GetCsr25() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr25FieldCsr25Mask) >> RegisterCsr25FieldCsr25Shift)
}

// SetCsr25 CSR25
func (r *RegisterCsr25Type) SetCsr25(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr25FieldCsr25Mask)|(uint32(value)<<RegisterCsr25FieldCsr25Shift))
}

// RegisterCsr26Type context swap registers
type RegisterCsr26Type uint32

func (r *RegisterCsr26Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr26Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr26Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr26Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr26Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr26FieldCsr26Shift = 0
	RegisterCsr26FieldCsr26Mask  = 0xffffffff
)

// GetCsr26 CSR26
func (r *RegisterCsr26Type) GetCsr26() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr26FieldCsr26Mask) >> RegisterCsr26FieldCsr26Shift)
}

// SetCsr26 CSR26
func (r *RegisterCsr26Type) SetCsr26(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr26FieldCsr26Mask)|(uint32(value)<<RegisterCsr26FieldCsr26Shift))
}

// RegisterCsr27Type context swap registers
type RegisterCsr27Type uint32

func (r *RegisterCsr27Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr27Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr27Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr27Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr27Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr27FieldCsr27Shift = 0
	RegisterCsr27FieldCsr27Mask  = 0xffffffff
)

// GetCsr27 CSR27
func (r *RegisterCsr27Type) GetCsr27() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr27FieldCsr27Mask) >> RegisterCsr27FieldCsr27Shift)
}

// SetCsr27 CSR27
func (r *RegisterCsr27Type) SetCsr27(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr27FieldCsr27Mask)|(uint32(value)<<RegisterCsr27FieldCsr27Shift))
}

// RegisterCsr28Type context swap registers
type RegisterCsr28Type uint32

func (r *RegisterCsr28Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr28Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr28Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr28Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr28Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr28FieldCsr28Shift = 0
	RegisterCsr28FieldCsr28Mask  = 0xffffffff
)

// GetCsr28 CSR28
func (r *RegisterCsr28Type) GetCsr28() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr28FieldCsr28Mask) >> RegisterCsr28FieldCsr28Shift)
}

// SetCsr28 CSR28
func (r *RegisterCsr28Type) SetCsr28(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr28FieldCsr28Mask)|(uint32(value)<<RegisterCsr28FieldCsr28Shift))
}

// RegisterCsr29Type context swap registers
type RegisterCsr29Type uint32

func (r *RegisterCsr29Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr29Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr29Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr29Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr29Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr29FieldCsr29Shift = 0
	RegisterCsr29FieldCsr29Mask  = 0xffffffff
)

// GetCsr29 CSR29
func (r *RegisterCsr29Type) GetCsr29() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr29FieldCsr29Mask) >> RegisterCsr29FieldCsr29Shift)
}

// SetCsr29 CSR29
func (r *RegisterCsr29Type) SetCsr29(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr29FieldCsr29Mask)|(uint32(value)<<RegisterCsr29FieldCsr29Shift))
}

// RegisterCsr30Type context swap registers
type RegisterCsr30Type uint32

func (r *RegisterCsr30Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr30Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr30Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr30Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr30Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr30FieldCsr30Shift = 0
	RegisterCsr30FieldCsr30Mask  = 0xffffffff
)

// GetCsr30 CSR30
func (r *RegisterCsr30Type) GetCsr30() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr30FieldCsr30Mask) >> RegisterCsr30FieldCsr30Shift)
}

// SetCsr30 CSR30
func (r *RegisterCsr30Type) SetCsr30(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr30FieldCsr30Mask)|(uint32(value)<<RegisterCsr30FieldCsr30Shift))
}

// RegisterCsr31Type context swap registers
type RegisterCsr31Type uint32

func (r *RegisterCsr31Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr31Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr31Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr31Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr31Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr31FieldCsr31Shift = 0
	RegisterCsr31FieldCsr31Mask  = 0xffffffff
)

// GetCsr31 CSR31
func (r *RegisterCsr31Type) GetCsr31() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr31FieldCsr31Mask) >> RegisterCsr31FieldCsr31Shift)
}

// SetCsr31 CSR31
func (r *RegisterCsr31Type) SetCsr31(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr31FieldCsr31Mask)|(uint32(value)<<RegisterCsr31FieldCsr31Shift))
}

// RegisterCsr32Type context swap registers
type RegisterCsr32Type uint32

func (r *RegisterCsr32Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr32Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr32Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr32Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr32Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr32FieldCsr32Shift = 0
	RegisterCsr32FieldCsr32Mask  = 0xffffffff
)

// GetCsr32 CSR32
func (r *RegisterCsr32Type) GetCsr32() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr32FieldCsr32Mask) >> RegisterCsr32FieldCsr32Shift)
}

// SetCsr32 CSR32
func (r *RegisterCsr32Type) SetCsr32(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr32FieldCsr32Mask)|(uint32(value)<<RegisterCsr32FieldCsr32Shift))
}

// RegisterCsr33Type context swap registers
type RegisterCsr33Type uint32

func (r *RegisterCsr33Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr33Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr33Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr33Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr33Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr33FieldCsr33Shift = 0
	RegisterCsr33FieldCsr33Mask  = 0xffffffff
)

// GetCsr33 CSR33
func (r *RegisterCsr33Type) GetCsr33() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr33FieldCsr33Mask) >> RegisterCsr33FieldCsr33Shift)
}

// SetCsr33 CSR33
func (r *RegisterCsr33Type) SetCsr33(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr33FieldCsr33Mask)|(uint32(value)<<RegisterCsr33FieldCsr33Shift))
}

// RegisterCsr34Type context swap registers
type RegisterCsr34Type uint32

func (r *RegisterCsr34Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr34Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr34Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr34Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr34Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr34FieldCsr34Shift = 0
	RegisterCsr34FieldCsr34Mask  = 0xffffffff
)

// GetCsr34 CSR34
func (r *RegisterCsr34Type) GetCsr34() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr34FieldCsr34Mask) >> RegisterCsr34FieldCsr34Shift)
}

// SetCsr34 CSR34
func (r *RegisterCsr34Type) SetCsr34(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr34FieldCsr34Mask)|(uint32(value)<<RegisterCsr34FieldCsr34Shift))
}

// RegisterCsr35Type context swap registers
type RegisterCsr35Type uint32

func (r *RegisterCsr35Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr35Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr35Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr35Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr35Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr35FieldCsr35Shift = 0
	RegisterCsr35FieldCsr35Mask  = 0xffffffff
)

// GetCsr35 CSR35
func (r *RegisterCsr35Type) GetCsr35() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr35FieldCsr35Mask) >> RegisterCsr35FieldCsr35Shift)
}

// SetCsr35 CSR35
func (r *RegisterCsr35Type) SetCsr35(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr35FieldCsr35Mask)|(uint32(value)<<RegisterCsr35FieldCsr35Shift))
}

// RegisterCsr36Type context swap registers
type RegisterCsr36Type uint32

func (r *RegisterCsr36Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr36Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr36Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr36Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr36Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr36FieldCsr36Shift = 0
	RegisterCsr36FieldCsr36Mask  = 0xffffffff
)

// GetCsr36 CSR36
func (r *RegisterCsr36Type) GetCsr36() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr36FieldCsr36Mask) >> RegisterCsr36FieldCsr36Shift)
}

// SetCsr36 CSR36
func (r *RegisterCsr36Type) SetCsr36(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr36FieldCsr36Mask)|(uint32(value)<<RegisterCsr36FieldCsr36Shift))
}

// RegisterCsr37Type context swap registers
type RegisterCsr37Type uint32

func (r *RegisterCsr37Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr37Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr37Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr37Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr37Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr37FieldCsr37Shift = 0
	RegisterCsr37FieldCsr37Mask  = 0xffffffff
)

// GetCsr37 CSR37
func (r *RegisterCsr37Type) GetCsr37() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr37FieldCsr37Mask) >> RegisterCsr37FieldCsr37Shift)
}

// SetCsr37 CSR37
func (r *RegisterCsr37Type) SetCsr37(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr37FieldCsr37Mask)|(uint32(value)<<RegisterCsr37FieldCsr37Shift))
}

// RegisterCsr38Type context swap registers
type RegisterCsr38Type uint32

func (r *RegisterCsr38Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr38Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr38Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr38Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr38Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr38FieldCsr38Shift = 0
	RegisterCsr38FieldCsr38Mask  = 0xffffffff
)

// GetCsr38 CSR38
func (r *RegisterCsr38Type) GetCsr38() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr38FieldCsr38Mask) >> RegisterCsr38FieldCsr38Shift)
}

// SetCsr38 CSR38
func (r *RegisterCsr38Type) SetCsr38(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr38FieldCsr38Mask)|(uint32(value)<<RegisterCsr38FieldCsr38Shift))
}

// RegisterCsr39Type context swap registers
type RegisterCsr39Type uint32

func (r *RegisterCsr39Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr39Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr39Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr39Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr39Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr39FieldCsr39Shift = 0
	RegisterCsr39FieldCsr39Mask  = 0xffffffff
)

// GetCsr39 CSR39
func (r *RegisterCsr39Type) GetCsr39() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr39FieldCsr39Mask) >> RegisterCsr39FieldCsr39Shift)
}

// SetCsr39 CSR39
func (r *RegisterCsr39Type) SetCsr39(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr39FieldCsr39Mask)|(uint32(value)<<RegisterCsr39FieldCsr39Shift))
}

// RegisterCsr40Type context swap registers
type RegisterCsr40Type uint32

func (r *RegisterCsr40Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr40Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr40Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr40Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr40Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr40FieldCsr40Shift = 0
	RegisterCsr40FieldCsr40Mask  = 0xffffffff
)

// GetCsr40 CSR40
func (r *RegisterCsr40Type) GetCsr40() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr40FieldCsr40Mask) >> RegisterCsr40FieldCsr40Shift)
}

// SetCsr40 CSR40
func (r *RegisterCsr40Type) SetCsr40(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr40FieldCsr40Mask)|(uint32(value)<<RegisterCsr40FieldCsr40Shift))
}

// RegisterCsr41Type context swap registers
type RegisterCsr41Type uint32

func (r *RegisterCsr41Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr41Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr41Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr41Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr41Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr41FieldCsr41Shift = 0
	RegisterCsr41FieldCsr41Mask  = 0xffffffff
)

// GetCsr41 CSR41
func (r *RegisterCsr41Type) GetCsr41() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr41FieldCsr41Mask) >> RegisterCsr41FieldCsr41Shift)
}

// SetCsr41 CSR41
func (r *RegisterCsr41Type) SetCsr41(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr41FieldCsr41Mask)|(uint32(value)<<RegisterCsr41FieldCsr41Shift))
}

// RegisterCsr42Type context swap registers
type RegisterCsr42Type uint32

func (r *RegisterCsr42Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr42Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr42Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr42Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr42Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr42FieldCsr42Shift = 0
	RegisterCsr42FieldCsr42Mask  = 0xffffffff
)

// GetCsr42 CSR42
func (r *RegisterCsr42Type) GetCsr42() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr42FieldCsr42Mask) >> RegisterCsr42FieldCsr42Shift)
}

// SetCsr42 CSR42
func (r *RegisterCsr42Type) SetCsr42(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr42FieldCsr42Mask)|(uint32(value)<<RegisterCsr42FieldCsr42Shift))
}

// RegisterCsr43Type context swap registers
type RegisterCsr43Type uint32

func (r *RegisterCsr43Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr43Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr43Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr43Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr43Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr43FieldCsr43Shift = 0
	RegisterCsr43FieldCsr43Mask  = 0xffffffff
)

// GetCsr43 CSR43
func (r *RegisterCsr43Type) GetCsr43() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr43FieldCsr43Mask) >> RegisterCsr43FieldCsr43Shift)
}

// SetCsr43 CSR43
func (r *RegisterCsr43Type) SetCsr43(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr43FieldCsr43Mask)|(uint32(value)<<RegisterCsr43FieldCsr43Shift))
}

// RegisterCsr44Type context swap registers
type RegisterCsr44Type uint32

func (r *RegisterCsr44Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr44Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr44Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr44Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr44Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr44FieldCsr44Shift = 0
	RegisterCsr44FieldCsr44Mask  = 0xffffffff
)

// GetCsr44 CSR44
func (r *RegisterCsr44Type) GetCsr44() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr44FieldCsr44Mask) >> RegisterCsr44FieldCsr44Shift)
}

// SetCsr44 CSR44
func (r *RegisterCsr44Type) SetCsr44(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr44FieldCsr44Mask)|(uint32(value)<<RegisterCsr44FieldCsr44Shift))
}

// RegisterCsr45Type context swap registers
type RegisterCsr45Type uint32

func (r *RegisterCsr45Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr45Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr45Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr45Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr45Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr45FieldCsr45Shift = 0
	RegisterCsr45FieldCsr45Mask  = 0xffffffff
)

// GetCsr45 CSR45
func (r *RegisterCsr45Type) GetCsr45() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr45FieldCsr45Mask) >> RegisterCsr45FieldCsr45Shift)
}

// SetCsr45 CSR45
func (r *RegisterCsr45Type) SetCsr45(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr45FieldCsr45Mask)|(uint32(value)<<RegisterCsr45FieldCsr45Shift))
}

// RegisterCsr46Type context swap registers
type RegisterCsr46Type uint32

func (r *RegisterCsr46Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr46Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr46Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr46Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr46Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr46FieldCsr46Shift = 0
	RegisterCsr46FieldCsr46Mask  = 0xffffffff
)

// GetCsr46 CSR46
func (r *RegisterCsr46Type) GetCsr46() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr46FieldCsr46Mask) >> RegisterCsr46FieldCsr46Shift)
}

// SetCsr46 CSR46
func (r *RegisterCsr46Type) SetCsr46(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr46FieldCsr46Mask)|(uint32(value)<<RegisterCsr46FieldCsr46Shift))
}

// RegisterCsr47Type context swap registers
type RegisterCsr47Type uint32

func (r *RegisterCsr47Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr47Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr47Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr47Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr47Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr47FieldCsr47Shift = 0
	RegisterCsr47FieldCsr47Mask  = 0xffffffff
)

// GetCsr47 CSR47
func (r *RegisterCsr47Type) GetCsr47() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr47FieldCsr47Mask) >> RegisterCsr47FieldCsr47Shift)
}

// SetCsr47 CSR47
func (r *RegisterCsr47Type) SetCsr47(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr47FieldCsr47Mask)|(uint32(value)<<RegisterCsr47FieldCsr47Shift))
}

// RegisterCsr48Type context swap registers
type RegisterCsr48Type uint32

func (r *RegisterCsr48Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr48Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr48Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr48Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr48Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr48FieldCsr48Shift = 0
	RegisterCsr48FieldCsr48Mask  = 0xffffffff
)

// GetCsr48 CSR48
func (r *RegisterCsr48Type) GetCsr48() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr48FieldCsr48Mask) >> RegisterCsr48FieldCsr48Shift)
}

// SetCsr48 CSR48
func (r *RegisterCsr48Type) SetCsr48(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr48FieldCsr48Mask)|(uint32(value)<<RegisterCsr48FieldCsr48Shift))
}

// RegisterCsr49Type context swap registers
type RegisterCsr49Type uint32

func (r *RegisterCsr49Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr49Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr49Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr49Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr49Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr49FieldCsr49Shift = 0
	RegisterCsr49FieldCsr49Mask  = 0xffffffff
)

// GetCsr49 CSR49
func (r *RegisterCsr49Type) GetCsr49() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr49FieldCsr49Mask) >> RegisterCsr49FieldCsr49Shift)
}

// SetCsr49 CSR49
func (r *RegisterCsr49Type) SetCsr49(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr49FieldCsr49Mask)|(uint32(value)<<RegisterCsr49FieldCsr49Shift))
}

// RegisterCsr50Type context swap registers
type RegisterCsr50Type uint32

func (r *RegisterCsr50Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr50Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr50Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr50Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr50Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr50FieldCsr50Shift = 0
	RegisterCsr50FieldCsr50Mask  = 0xffffffff
)

// GetCsr50 CSR50
func (r *RegisterCsr50Type) GetCsr50() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr50FieldCsr50Mask) >> RegisterCsr50FieldCsr50Shift)
}

// SetCsr50 CSR50
func (r *RegisterCsr50Type) SetCsr50(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr50FieldCsr50Mask)|(uint32(value)<<RegisterCsr50FieldCsr50Shift))
}

// RegisterCsr51Type context swap registers
type RegisterCsr51Type uint32

func (r *RegisterCsr51Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr51Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr51Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr51Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr51Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr51FieldCsr51Shift = 0
	RegisterCsr51FieldCsr51Mask  = 0xffffffff
)

// GetCsr51 CSR51
func (r *RegisterCsr51Type) GetCsr51() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr51FieldCsr51Mask) >> RegisterCsr51FieldCsr51Shift)
}

// SetCsr51 CSR51
func (r *RegisterCsr51Type) SetCsr51(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr51FieldCsr51Mask)|(uint32(value)<<RegisterCsr51FieldCsr51Shift))
}

// RegisterCsr52Type context swap registers
type RegisterCsr52Type uint32

func (r *RegisterCsr52Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr52Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr52Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr52Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr52Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr52FieldCsr52Shift = 0
	RegisterCsr52FieldCsr52Mask  = 0xffffffff
)

// GetCsr52 CSR52
func (r *RegisterCsr52Type) GetCsr52() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr52FieldCsr52Mask) >> RegisterCsr52FieldCsr52Shift)
}

// SetCsr52 CSR52
func (r *RegisterCsr52Type) SetCsr52(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr52FieldCsr52Mask)|(uint32(value)<<RegisterCsr52FieldCsr52Shift))
}

// RegisterCsr53Type context swap registers
type RegisterCsr53Type uint32

func (r *RegisterCsr53Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsr53Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsr53Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsr53Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsr53Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsr53FieldCsr53Shift = 0
	RegisterCsr53FieldCsr53Mask  = 0xffffffff
)

// GetCsr53 CSR53
func (r *RegisterCsr53Type) GetCsr53() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsr53FieldCsr53Mask) >> RegisterCsr53FieldCsr53Shift)
}

// SetCsr53 CSR53
func (r *RegisterCsr53Type) SetCsr53(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr53FieldCsr53Mask)|(uint32(value)<<RegisterCsr53FieldCsr53Shift))
}

// RegisterHashhr0Type HASH digest register
type RegisterHashhr0Type uint32

func (r *RegisterHashhr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHashhr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHashhr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHashhr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHashhr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHashhr0FieldH0Shift = 0
	RegisterHashhr0FieldH0Mask  = 0xffffffff
)

// GetH0 H0
func (r *RegisterHashhr0Type) GetH0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHashhr0FieldH0Mask) >> RegisterHashhr0FieldH0Shift)
}

// SetH0 H0
func (r *RegisterHashhr0Type) SetH0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHashhr0FieldH0Mask)|(uint32(value)<<RegisterHashhr0FieldH0Shift))
}

// RegisterHashhr1Type read-only
type RegisterHashhr1Type uint32

func (r *RegisterHashhr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHashhr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHashhr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHashhr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHashhr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHashhr1FieldH1Shift = 0
	RegisterHashhr1FieldH1Mask  = 0xffffffff
)

// GetH1 H1
func (r *RegisterHashhr1Type) GetH1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHashhr1FieldH1Mask) >> RegisterHashhr1FieldH1Shift)
}

// SetH1 H1
func (r *RegisterHashhr1Type) SetH1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHashhr1FieldH1Mask)|(uint32(value)<<RegisterHashhr1FieldH1Shift))
}

// RegisterHashhr2Type read-only
type RegisterHashhr2Type uint32

func (r *RegisterHashhr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHashhr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHashhr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHashhr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHashhr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHashhr2FieldH2Shift = 0
	RegisterHashhr2FieldH2Mask  = 0xffffffff
)

// GetH2 H2
func (r *RegisterHashhr2Type) GetH2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHashhr2FieldH2Mask) >> RegisterHashhr2FieldH2Shift)
}

// SetH2 H2
func (r *RegisterHashhr2Type) SetH2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHashhr2FieldH2Mask)|(uint32(value)<<RegisterHashhr2FieldH2Shift))
}

// RegisterHashhr3Type read-only
type RegisterHashhr3Type uint32

func (r *RegisterHashhr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHashhr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHashhr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHashhr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHashhr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHashhr3FieldH3Shift = 0
	RegisterHashhr3FieldH3Mask  = 0xffffffff
)

// GetH3 H3
func (r *RegisterHashhr3Type) GetH3() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHashhr3FieldH3Mask) >> RegisterHashhr3FieldH3Shift)
}

// SetH3 H3
func (r *RegisterHashhr3Type) SetH3(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHashhr3FieldH3Mask)|(uint32(value)<<RegisterHashhr3FieldH3Shift))
}

// RegisterHashhr4Type read-only
type RegisterHashhr4Type uint32

func (r *RegisterHashhr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHashhr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHashhr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHashhr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHashhr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHashhr4FieldH4Shift = 0
	RegisterHashhr4FieldH4Mask  = 0xffffffff
)

// GetH4 H4
func (r *RegisterHashhr4Type) GetH4() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHashhr4FieldH4Mask) >> RegisterHashhr4FieldH4Shift)
}

// SetH4 H4
func (r *RegisterHashhr4Type) SetH4(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHashhr4FieldH4Mask)|(uint32(value)<<RegisterHashhr4FieldH4Shift))
}

// RegisterHashhr5Type read-only
type RegisterHashhr5Type uint32

func (r *RegisterHashhr5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHashhr5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHashhr5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHashhr5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHashhr5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHashhr5FieldH5Shift = 0
	RegisterHashhr5FieldH5Mask  = 0xffffffff
)

// GetH5 H5
func (r *RegisterHashhr5Type) GetH5() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHashhr5FieldH5Mask) >> RegisterHashhr5FieldH5Shift)
}

// SetH5 H5
func (r *RegisterHashhr5Type) SetH5(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHashhr5FieldH5Mask)|(uint32(value)<<RegisterHashhr5FieldH5Shift))
}

// RegisterHashhr6Type read-only
type RegisterHashhr6Type uint32

func (r *RegisterHashhr6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHashhr6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHashhr6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHashhr6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHashhr6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHashhr6FieldH6Shift = 0
	RegisterHashhr6FieldH6Mask  = 0xffffffff
)

// GetH6 H6
func (r *RegisterHashhr6Type) GetH6() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHashhr6FieldH6Mask) >> RegisterHashhr6FieldH6Shift)
}

// SetH6 H6
func (r *RegisterHashhr6Type) SetH6(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHashhr6FieldH6Mask)|(uint32(value)<<RegisterHashhr6FieldH6Shift))
}

// RegisterHashhr7Type read-only
type RegisterHashhr7Type uint32

func (r *RegisterHashhr7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHashhr7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHashhr7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHashhr7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHashhr7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHashhr7FieldH7Shift = 0
	RegisterHashhr7FieldH7Mask  = 0xffffffff
)

// GetH7 H7
func (r *RegisterHashhr7Type) GetH7() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHashhr7FieldH7Mask) >> RegisterHashhr7FieldH7Shift)
}

// SetH7 H7
func (r *RegisterHashhr7Type) SetH7(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHashhr7FieldH7Mask)|(uint32(value)<<RegisterHashhr7FieldH7Shift))
}
