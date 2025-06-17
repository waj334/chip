//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package axi

import (
	"unsafe"
	"volatile"
)

var (
	Axi = (*_axi)(unsafe.Pointer(uintptr(0x51000000)))
)

type _axi struct {
	_                  [8144]byte
	Axiperiphid4       RegisterAxiperiphid4Type
	_                  [12]byte
	Axiperiphid0       RegisterAxiperiphid0Type
	Axiperiphid1       RegisterAxiperiphid1Type
	Axiperiphid2       RegisterAxiperiphid2Type
	Axiperiphid3       RegisterAxiperiphid3Type
	Axicompid0         RegisterAxicompid0Type
	Axicompid1         RegisterAxicompid1Type
	Axicompid2         RegisterAxicompid2Type
	Axicompid3         RegisterAxicompid3Type
	_                  [8]byte
	Axitarg1fnmodissbm RegisterAxitarg1fnmodissbmType
	_                  [24]byte
	Axitarg1fnmod2     RegisterAxitarg1fnmod2Type
	_                  [4]byte
	Axitarg1fnmodlb    RegisterAxitarg1fnmodlbType
	_                  [216]byte
	Axitarg1fnmod      RegisterAxitarg1fnmodType
	_                  [3836]byte
	Axitarg2fnmodissbm RegisterAxitarg2fnmodissbmType
	_                  [24]byte
	Axitarg2fnmod2     RegisterAxitarg2fnmod2Type
	_                  [4]byte
	Axitarg2fnmodlb    RegisterAxitarg2fnmodlbType
	_                  [216]byte
	Axitarg2fnmod      RegisterAxitarg2fnmodType
	_                  [3836]byte
	Axitarg3fnmodissbm RegisterAxitarg3fnmodissbmType
	_                  [4092]byte
	Axitarg4fnmodissbm RegisterAxitarg4fnmodissbmType
	_                  [4092]byte
	Axitarg5fnmodissbm RegisterAxitarg5fnmodissbmType
	_                  [4092]byte
	Axitarg6fnmodissbm RegisterAxitarg6fnmodissbmType
	_                  [4096]byte
	Axitarg7fnmodissbm RegisterAxitarg7fnmodissbmType
	_                  [20]byte
	Axitarg7fnmod2     RegisterAxitarg7fnmod2Type
	_                  [224]byte
	Axitarg7fnmod      RegisterAxitarg7fnmodType
	_                  [237336]byte
	Axiini1fnmod2      RegisterAxiini1fnmod2Type
	Axiini1fnmodahb    RegisterAxiini1fnmodahbType
	_                  [212]byte
	Axiini1readqos     RegisterAxiini1readqosType
	Axiini1writeqos    RegisterAxiini1writeqosType
	Axiini1fnmod       RegisterAxiini1fnmodType
	_                  [4084]byte
	Axiini2readqos     RegisterAxiini2readqosType
	Axiini2writeqos    RegisterAxiini2writeqosType
	Axiini2fnmod       RegisterAxiini2fnmodType
	_                  [3864]byte
	Axiini3fnmod2      RegisterAxiini3fnmod2Type
	Axiini3fnmodahb    RegisterAxiini3fnmodahbType
	_                  [212]byte
	Axiini3readqos     RegisterAxiini3readqosType
	Axiini3writeqos    RegisterAxiini3writeqosType
	Axiini3fnmod       RegisterAxiini3fnmodType
	_                  [4084]byte
	Axiini4readqos     RegisterAxiini4readqosType
	Axiini4writeqos    RegisterAxiini4writeqosType
	Axiini4fnmod       RegisterAxiini4fnmodType
	_                  [4084]byte
	Axiini5readqos     RegisterAxiini5readqosType
	Axiini5writeqos    RegisterAxiini5writeqosType
	Axiini5fnmod       RegisterAxiini5fnmodType
	_                  [4084]byte
	Axiini6readqos     RegisterAxiini6readqosType
	Axiini6writeqos    RegisterAxiini6writeqosType
	Axiini6fnmod       RegisterAxiini6fnmodType
}

// RegisterAxiperiphid4Type AXI interconnect - peripheral ID4 register
type RegisterAxiperiphid4Type uint32

func (r *RegisterAxiperiphid4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiperiphid4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiperiphid4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiperiphid4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiperiphid4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiperiphid4FieldJep106conShift = 0
	RegisterAxiperiphid4FieldJep106conMask  = 0xf
)

// GetJep106con JEP106 continuation code
func (r *RegisterAxiperiphid4Type) GetJep106con() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid4FieldJep106conMask) >> RegisterAxiperiphid4FieldJep106conShift)
}

// SetJep106con JEP106 continuation code
func (r *RegisterAxiperiphid4Type) SetJep106con(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid4FieldJep106conMask)|(uint32(value)<<RegisterAxiperiphid4FieldJep106conShift))
}

const (
	RegisterAxiperiphid4FieldKcount4Shift = 4
	RegisterAxiperiphid4FieldKcount4Mask  = 0xf0
)

// GetKcount4 Register file size
func (r *RegisterAxiperiphid4Type) GetKcount4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid4FieldKcount4Mask) >> RegisterAxiperiphid4FieldKcount4Shift)
}

// SetKcount4 Register file size
func (r *RegisterAxiperiphid4Type) SetKcount4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid4FieldKcount4Mask)|(uint32(value)<<RegisterAxiperiphid4FieldKcount4Shift))
}

// RegisterAxiperiphid0Type AXI interconnect - peripheral ID0 register
type RegisterAxiperiphid0Type uint32

func (r *RegisterAxiperiphid0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiperiphid0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiperiphid0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiperiphid0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiperiphid0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiperiphid0FieldPartnumShift = 0
	RegisterAxiperiphid0FieldPartnumMask  = 0xff
)

// GetPartnum Peripheral part number bits 0 to 7
func (r *RegisterAxiperiphid0Type) GetPartnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid0FieldPartnumMask) >> RegisterAxiperiphid0FieldPartnumShift)
}

// SetPartnum Peripheral part number bits 0 to 7
func (r *RegisterAxiperiphid0Type) SetPartnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid0FieldPartnumMask)|(uint32(value)<<RegisterAxiperiphid0FieldPartnumShift))
}

// RegisterAxiperiphid1Type AXI interconnect - peripheral ID1 register
type RegisterAxiperiphid1Type uint32

func (r *RegisterAxiperiphid1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiperiphid1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiperiphid1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiperiphid1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiperiphid1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiperiphid1FieldPartnumShift = 0
	RegisterAxiperiphid1FieldPartnumMask  = 0xf
)

// GetPartnum Peripheral part number bits 8 to 11
func (r *RegisterAxiperiphid1Type) GetPartnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid1FieldPartnumMask) >> RegisterAxiperiphid1FieldPartnumShift)
}

// SetPartnum Peripheral part number bits 8 to 11
func (r *RegisterAxiperiphid1Type) SetPartnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid1FieldPartnumMask)|(uint32(value)<<RegisterAxiperiphid1FieldPartnumShift))
}

const (
	RegisterAxiperiphid1FieldJep106iShift = 4
	RegisterAxiperiphid1FieldJep106iMask  = 0xf0
)

// GetJep106i JEP106 identity bits 0 to 3
func (r *RegisterAxiperiphid1Type) GetJep106i() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid1FieldJep106iMask) >> RegisterAxiperiphid1FieldJep106iShift)
}

// SetJep106i JEP106 identity bits 0 to 3
func (r *RegisterAxiperiphid1Type) SetJep106i(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid1FieldJep106iMask)|(uint32(value)<<RegisterAxiperiphid1FieldJep106iShift))
}

// RegisterAxiperiphid2Type AXI interconnect - peripheral ID2 register
type RegisterAxiperiphid2Type uint32

func (r *RegisterAxiperiphid2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiperiphid2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiperiphid2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiperiphid2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiperiphid2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiperiphid2FieldJep106idShift = 0
	RegisterAxiperiphid2FieldJep106idMask  = 0x7
)

// GetJep106id JEP106 Identity bits 4 to 6
func (r *RegisterAxiperiphid2Type) GetJep106id() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid2FieldJep106idMask) >> RegisterAxiperiphid2FieldJep106idShift)
}

// SetJep106id JEP106 Identity bits 4 to 6
func (r *RegisterAxiperiphid2Type) SetJep106id(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid2FieldJep106idMask)|(uint32(value)<<RegisterAxiperiphid2FieldJep106idShift))
}

const (
	RegisterAxiperiphid2FieldJedecShift = 3
	RegisterAxiperiphid2FieldJedecMask  = 0x8
)

// GetJedec JEP106 code flag
func (r *RegisterAxiperiphid2Type) GetJedec() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid2FieldJedecMask) != 0
}

// SetJedec JEP106 code flag
func (r *RegisterAxiperiphid2Type) SetJedec(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiperiphid2FieldJedecMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid2FieldJedecMask)
	}
}

const (
	RegisterAxiperiphid2FieldRevisionShift = 4
	RegisterAxiperiphid2FieldRevisionMask  = 0xf0
)

// GetRevision Peripheral revision number
func (r *RegisterAxiperiphid2Type) GetRevision() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid2FieldRevisionMask) >> RegisterAxiperiphid2FieldRevisionShift)
}

// SetRevision Peripheral revision number
func (r *RegisterAxiperiphid2Type) SetRevision(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid2FieldRevisionMask)|(uint32(value)<<RegisterAxiperiphid2FieldRevisionShift))
}

// RegisterAxiperiphid3Type AXI interconnect - peripheral ID3 register
type RegisterAxiperiphid3Type uint32

func (r *RegisterAxiperiphid3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiperiphid3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiperiphid3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiperiphid3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiperiphid3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiperiphid3FieldCustmodnumShift = 0
	RegisterAxiperiphid3FieldCustmodnumMask  = 0xf
)

// GetCustmodnum Customer modification
func (r *RegisterAxiperiphid3Type) GetCustmodnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid3FieldCustmodnumMask) >> RegisterAxiperiphid3FieldCustmodnumShift)
}

// SetCustmodnum Customer modification
func (r *RegisterAxiperiphid3Type) SetCustmodnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid3FieldCustmodnumMask)|(uint32(value)<<RegisterAxiperiphid3FieldCustmodnumShift))
}

const (
	RegisterAxiperiphid3FieldRevandShift = 4
	RegisterAxiperiphid3FieldRevandMask  = 0xf0
)

// GetRevand Customer version
func (r *RegisterAxiperiphid3Type) GetRevand() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid3FieldRevandMask) >> RegisterAxiperiphid3FieldRevandShift)
}

// SetRevand Customer version
func (r *RegisterAxiperiphid3Type) SetRevand(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid3FieldRevandMask)|(uint32(value)<<RegisterAxiperiphid3FieldRevandShift))
}

// RegisterAxicompid0Type AXI interconnect - component ID0 register
type RegisterAxicompid0Type uint32

func (r *RegisterAxicompid0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxicompid0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxicompid0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxicompid0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxicompid0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxicompid0FieldPreambleShift = 0
	RegisterAxicompid0FieldPreambleMask  = 0xff
)

// GetPreamble Preamble bits 0 to 7
func (r *RegisterAxicompid0Type) GetPreamble() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxicompid0FieldPreambleMask) >> RegisterAxicompid0FieldPreambleShift)
}

// SetPreamble Preamble bits 0 to 7
func (r *RegisterAxicompid0Type) SetPreamble(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxicompid0FieldPreambleMask)|(uint32(value)<<RegisterAxicompid0FieldPreambleShift))
}

// RegisterAxicompid1Type AXI interconnect - component ID1 register
type RegisterAxicompid1Type uint32

func (r *RegisterAxicompid1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxicompid1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxicompid1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxicompid1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxicompid1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxicompid1FieldPreambleShift = 0
	RegisterAxicompid1FieldPreambleMask  = 0xf
)

// GetPreamble Preamble bits 8 to 11
func (r *RegisterAxicompid1Type) GetPreamble() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxicompid1FieldPreambleMask) >> RegisterAxicompid1FieldPreambleShift)
}

// SetPreamble Preamble bits 8 to 11
func (r *RegisterAxicompid1Type) SetPreamble(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxicompid1FieldPreambleMask)|(uint32(value)<<RegisterAxicompid1FieldPreambleShift))
}

const (
	RegisterAxicompid1FieldClassShift = 4
	RegisterAxicompid1FieldClassMask  = 0xf0
)

// GetClass Component class
func (r *RegisterAxicompid1Type) GetClass() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxicompid1FieldClassMask) >> RegisterAxicompid1FieldClassShift)
}

// SetClass Component class
func (r *RegisterAxicompid1Type) SetClass(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxicompid1FieldClassMask)|(uint32(value)<<RegisterAxicompid1FieldClassShift))
}

// RegisterAxicompid2Type AXI interconnect - component ID2 register
type RegisterAxicompid2Type uint32

func (r *RegisterAxicompid2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxicompid2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxicompid2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxicompid2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxicompid2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxicompid2FieldPreambleShift = 0
	RegisterAxicompid2FieldPreambleMask  = 0xff
)

// GetPreamble Preamble bits 12 to 19
func (r *RegisterAxicompid2Type) GetPreamble() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxicompid2FieldPreambleMask) >> RegisterAxicompid2FieldPreambleShift)
}

// SetPreamble Preamble bits 12 to 19
func (r *RegisterAxicompid2Type) SetPreamble(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxicompid2FieldPreambleMask)|(uint32(value)<<RegisterAxicompid2FieldPreambleShift))
}

// RegisterAxicompid3Type AXI interconnect - component ID3 register
type RegisterAxicompid3Type uint32

func (r *RegisterAxicompid3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxicompid3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxicompid3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxicompid3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxicompid3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxicompid3FieldPreambleShift = 0
	RegisterAxicompid3FieldPreambleMask  = 0xff
)

// GetPreamble Preamble bits 20 to 27
func (r *RegisterAxicompid3Type) GetPreamble() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxicompid3FieldPreambleMask) >> RegisterAxicompid3FieldPreambleShift)
}

// SetPreamble Preamble bits 20 to 27
func (r *RegisterAxicompid3Type) SetPreamble(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxicompid3FieldPreambleMask)|(uint32(value)<<RegisterAxicompid3FieldPreambleShift))
}

// RegisterAxitarg1fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type RegisterAxitarg1fnmodissbmType uint32

func (r *RegisterAxitarg1fnmodissbmType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg1fnmodissbmType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg1fnmodissbmType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg1fnmodissbmType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg1fnmodissbmType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg1fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg1fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg1fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg1fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg1fnmodissbmType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg1fnmodissbmFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg1fnmodissbmFieldReadissoverrideMask)
	}
}

const (
	RegisterAxitarg1fnmodissbmFieldWriteissoverrideShift = 1
	RegisterAxitarg1fnmodissbmFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg1fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg1fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg1fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg1fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg1fnmodissbmFieldWriteissoverrideMask)
	}
}

// RegisterAxitarg1fnmod2Type AXI interconnect - TARG x bus matrix functionality 2 register
type RegisterAxitarg1fnmod2Type uint32

func (r *RegisterAxitarg1fnmod2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg1fnmod2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg1fnmod2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg1fnmod2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg1fnmod2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg1fnmod2FieldBypassmergeShift = 0
	RegisterAxitarg1fnmod2FieldBypassmergeMask  = 0x1
)

// GetBypassmerge Disable packing of beats to match the output data width
func (r *RegisterAxitarg1fnmod2Type) GetBypassmerge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg1fnmod2FieldBypassmergeMask) != 0
}

// SetBypassmerge Disable packing of beats to match the output data width
func (r *RegisterAxitarg1fnmod2Type) SetBypassmerge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg1fnmod2FieldBypassmergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg1fnmod2FieldBypassmergeMask)
	}
}

// RegisterAxitarg1fnmodlbType AXI interconnect - TARG x long burst functionality modification
type RegisterAxitarg1fnmodlbType uint32

func (r *RegisterAxitarg1fnmodlbType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg1fnmodlbType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg1fnmodlbType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg1fnmodlbType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg1fnmodlbType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg1fnmodlbFieldFnmodlbShift = 0
	RegisterAxitarg1fnmodlbFieldFnmodlbMask  = 0x1
)

// GetFnmodlb Controls burst breaking of long bursts
func (r *RegisterAxitarg1fnmodlbType) GetFnmodlb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg1fnmodlbFieldFnmodlbMask) != 0
}

// SetFnmodlb Controls burst breaking of long bursts
func (r *RegisterAxitarg1fnmodlbType) SetFnmodlb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg1fnmodlbFieldFnmodlbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg1fnmodlbFieldFnmodlbMask)
	}
}

// RegisterAxitarg1fnmodType AXI interconnect - TARG x long burst functionality modification
type RegisterAxitarg1fnmodType uint32

func (r *RegisterAxitarg1fnmodType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg1fnmodType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg1fnmodType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg1fnmodType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg1fnmodType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg1fnmodFieldReadissoverrideShift = 0
	RegisterAxitarg1fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override AMIB read issuing capability
func (r *RegisterAxitarg1fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg1fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override AMIB read issuing capability
func (r *RegisterAxitarg1fnmodType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg1fnmodFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg1fnmodFieldReadissoverrideMask)
	}
}

const (
	RegisterAxitarg1fnmodFieldWriteissoverrideShift = 1
	RegisterAxitarg1fnmodFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Override AMIB write issuing capability
func (r *RegisterAxitarg1fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg1fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override AMIB write issuing capability
func (r *RegisterAxitarg1fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg1fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg1fnmodFieldWriteissoverrideMask)
	}
}

// RegisterAxitarg2fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type RegisterAxitarg2fnmodissbmType uint32

func (r *RegisterAxitarg2fnmodissbmType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg2fnmodissbmType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg2fnmodissbmType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg2fnmodissbmType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg2fnmodissbmType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg2fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg2fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg2fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg2fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg2fnmodissbmType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg2fnmodissbmFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg2fnmodissbmFieldReadissoverrideMask)
	}
}

const (
	RegisterAxitarg2fnmodissbmFieldWriteissoverrideShift = 1
	RegisterAxitarg2fnmodissbmFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg2fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg2fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg2fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg2fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg2fnmodissbmFieldWriteissoverrideMask)
	}
}

// RegisterAxitarg2fnmod2Type AXI interconnect - TARG x bus matrix functionality 2 register
type RegisterAxitarg2fnmod2Type uint32

func (r *RegisterAxitarg2fnmod2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg2fnmod2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg2fnmod2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg2fnmod2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg2fnmod2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg2fnmod2FieldBypassmergeShift = 0
	RegisterAxitarg2fnmod2FieldBypassmergeMask  = 0x1
)

// GetBypassmerge Disable packing of beats to match the output data width
func (r *RegisterAxitarg2fnmod2Type) GetBypassmerge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg2fnmod2FieldBypassmergeMask) != 0
}

// SetBypassmerge Disable packing of beats to match the output data width
func (r *RegisterAxitarg2fnmod2Type) SetBypassmerge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg2fnmod2FieldBypassmergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg2fnmod2FieldBypassmergeMask)
	}
}

// RegisterAxitarg2fnmodlbType AXI interconnect - TARG x long burst functionality modification
type RegisterAxitarg2fnmodlbType uint32

func (r *RegisterAxitarg2fnmodlbType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg2fnmodlbType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg2fnmodlbType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg2fnmodlbType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg2fnmodlbType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg2fnmodlbFieldFnmodlbShift = 0
	RegisterAxitarg2fnmodlbFieldFnmodlbMask  = 0x1
)

// GetFnmodlb Controls burst breaking of long bursts
func (r *RegisterAxitarg2fnmodlbType) GetFnmodlb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg2fnmodlbFieldFnmodlbMask) != 0
}

// SetFnmodlb Controls burst breaking of long bursts
func (r *RegisterAxitarg2fnmodlbType) SetFnmodlb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg2fnmodlbFieldFnmodlbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg2fnmodlbFieldFnmodlbMask)
	}
}

// RegisterAxitarg2fnmodType AXI interconnect - TARG x long burst functionality modification
type RegisterAxitarg2fnmodType uint32

func (r *RegisterAxitarg2fnmodType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg2fnmodType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg2fnmodType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg2fnmodType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg2fnmodType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg2fnmodFieldReadissoverrideShift = 0
	RegisterAxitarg2fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override AMIB read issuing capability
func (r *RegisterAxitarg2fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg2fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override AMIB read issuing capability
func (r *RegisterAxitarg2fnmodType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg2fnmodFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg2fnmodFieldReadissoverrideMask)
	}
}

const (
	RegisterAxitarg2fnmodFieldWriteissoverrideShift = 1
	RegisterAxitarg2fnmodFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Override AMIB write issuing capability
func (r *RegisterAxitarg2fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg2fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override AMIB write issuing capability
func (r *RegisterAxitarg2fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg2fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg2fnmodFieldWriteissoverrideMask)
	}
}

// RegisterAxitarg3fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type RegisterAxitarg3fnmodissbmType uint32

func (r *RegisterAxitarg3fnmodissbmType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg3fnmodissbmType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg3fnmodissbmType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg3fnmodissbmType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg3fnmodissbmType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg3fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg3fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg3fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg3fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg3fnmodissbmType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg3fnmodissbmFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg3fnmodissbmFieldReadissoverrideMask)
	}
}

const (
	RegisterAxitarg3fnmodissbmFieldWriteissoverrideShift = 1
	RegisterAxitarg3fnmodissbmFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg3fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg3fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg3fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg3fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg3fnmodissbmFieldWriteissoverrideMask)
	}
}

// RegisterAxitarg4fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type RegisterAxitarg4fnmodissbmType uint32

func (r *RegisterAxitarg4fnmodissbmType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg4fnmodissbmType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg4fnmodissbmType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg4fnmodissbmType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg4fnmodissbmType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg4fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg4fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg4fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg4fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg4fnmodissbmType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg4fnmodissbmFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg4fnmodissbmFieldReadissoverrideMask)
	}
}

const (
	RegisterAxitarg4fnmodissbmFieldWriteissoverrideShift = 1
	RegisterAxitarg4fnmodissbmFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg4fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg4fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg4fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg4fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg4fnmodissbmFieldWriteissoverrideMask)
	}
}

// RegisterAxitarg5fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type RegisterAxitarg5fnmodissbmType uint32

func (r *RegisterAxitarg5fnmodissbmType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg5fnmodissbmType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg5fnmodissbmType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg5fnmodissbmType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg5fnmodissbmType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg5fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg5fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg5fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg5fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg5fnmodissbmType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg5fnmodissbmFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg5fnmodissbmFieldReadissoverrideMask)
	}
}

const (
	RegisterAxitarg5fnmodissbmFieldWriteissoverrideShift = 1
	RegisterAxitarg5fnmodissbmFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg5fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg5fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg5fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg5fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg5fnmodissbmFieldWriteissoverrideMask)
	}
}

// RegisterAxitarg6fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type RegisterAxitarg6fnmodissbmType uint32

func (r *RegisterAxitarg6fnmodissbmType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg6fnmodissbmType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg6fnmodissbmType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg6fnmodissbmType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg6fnmodissbmType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg6fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg6fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg6fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg6fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg6fnmodissbmType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg6fnmodissbmFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg6fnmodissbmFieldReadissoverrideMask)
	}
}

const (
	RegisterAxitarg6fnmodissbmFieldWriteissoverrideShift = 1
	RegisterAxitarg6fnmodissbmFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg6fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg6fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg6fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg6fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg6fnmodissbmFieldWriteissoverrideMask)
	}
}

// RegisterAxitarg7fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type RegisterAxitarg7fnmodissbmType uint32

func (r *RegisterAxitarg7fnmodissbmType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg7fnmodissbmType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg7fnmodissbmType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg7fnmodissbmType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg7fnmodissbmType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg7fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg7fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg7fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg7fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *RegisterAxitarg7fnmodissbmType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg7fnmodissbmFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg7fnmodissbmFieldReadissoverrideMask)
	}
}

const (
	RegisterAxitarg7fnmodissbmFieldWriteissoverrideShift = 1
	RegisterAxitarg7fnmodissbmFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg7fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg7fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *RegisterAxitarg7fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg7fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg7fnmodissbmFieldWriteissoverrideMask)
	}
}

// RegisterAxitarg7fnmod2Type AXI interconnect - TARG x bus matrix functionality 2 register
type RegisterAxitarg7fnmod2Type uint32

func (r *RegisterAxitarg7fnmod2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg7fnmod2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg7fnmod2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg7fnmod2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg7fnmod2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg7fnmod2FieldBypassmergeShift = 0
	RegisterAxitarg7fnmod2FieldBypassmergeMask  = 0x1
)

// GetBypassmerge Disable packing of beats to match the output data width
func (r *RegisterAxitarg7fnmod2Type) GetBypassmerge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg7fnmod2FieldBypassmergeMask) != 0
}

// SetBypassmerge Disable packing of beats to match the output data width
func (r *RegisterAxitarg7fnmod2Type) SetBypassmerge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg7fnmod2FieldBypassmergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg7fnmod2FieldBypassmergeMask)
	}
}

// RegisterAxitarg7fnmodType AXI interconnect - TARG x long burst functionality modification
type RegisterAxitarg7fnmodType uint32

func (r *RegisterAxitarg7fnmodType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxitarg7fnmodType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxitarg7fnmodType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxitarg7fnmodType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxitarg7fnmodType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxitarg7fnmodFieldReadissoverrideShift = 0
	RegisterAxitarg7fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override AMIB read issuing capability
func (r *RegisterAxitarg7fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg7fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override AMIB read issuing capability
func (r *RegisterAxitarg7fnmodType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg7fnmodFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg7fnmodFieldReadissoverrideMask)
	}
}

const (
	RegisterAxitarg7fnmodFieldWriteissoverrideShift = 1
	RegisterAxitarg7fnmodFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Override AMIB write issuing capability
func (r *RegisterAxitarg7fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg7fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override AMIB write issuing capability
func (r *RegisterAxitarg7fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg7fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg7fnmodFieldWriteissoverrideMask)
	}
}

// RegisterAxiini1fnmod2Type AXI interconnect - INI x functionality modification 2 register
type RegisterAxiini1fnmod2Type uint32

func (r *RegisterAxiini1fnmod2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini1fnmod2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini1fnmod2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini1fnmod2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini1fnmod2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini1fnmod2FieldBypassmergeShift = 0
	RegisterAxiini1fnmod2FieldBypassmergeMask  = 0x1
)

// GetBypassmerge Disables alteration of transactions by the up-sizer unless required by the protocol
func (r *RegisterAxiini1fnmod2Type) GetBypassmerge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1fnmod2FieldBypassmergeMask) != 0
}

// SetBypassmerge Disables alteration of transactions by the up-sizer unless required by the protocol
func (r *RegisterAxiini1fnmod2Type) SetBypassmerge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini1fnmod2FieldBypassmergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini1fnmod2FieldBypassmergeMask)
	}
}

// RegisterAxiini1fnmodahbType AXI interconnect - INI x AHB functionality modification register
type RegisterAxiini1fnmodahbType uint32

func (r *RegisterAxiini1fnmodahbType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini1fnmodahbType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini1fnmodahbType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini1fnmodahbType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini1fnmodahbType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini1fnmodahbFieldRdincoverrideShift = 0
	RegisterAxiini1fnmodahbFieldRdincoverrideMask  = 0x1
)

// GetRdincoverride Converts all AHB-Lite write transactions to a series of single beat AXI
func (r *RegisterAxiini1fnmodahbType) GetRdincoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1fnmodahbFieldRdincoverrideMask) != 0
}

// SetRdincoverride Converts all AHB-Lite write transactions to a series of single beat AXI
func (r *RegisterAxiini1fnmodahbType) SetRdincoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini1fnmodahbFieldRdincoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini1fnmodahbFieldRdincoverrideMask)
	}
}

const (
	RegisterAxiini1fnmodahbFieldWrincoverrideShift = 1
	RegisterAxiini1fnmodahbFieldWrincoverrideMask  = 0x2
)

// GetWrincoverride Converts all AHB-Lite read transactions to a series of single beat AXI
func (r *RegisterAxiini1fnmodahbType) GetWrincoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1fnmodahbFieldWrincoverrideMask) != 0
}

// SetWrincoverride Converts all AHB-Lite read transactions to a series of single beat AXI
func (r *RegisterAxiini1fnmodahbType) SetWrincoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini1fnmodahbFieldWrincoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini1fnmodahbFieldWrincoverrideMask)
	}
}

// RegisterAxiini1readqosType AXI interconnect - INI x read QoS register
type RegisterAxiini1readqosType uint32

func (r *RegisterAxiini1readqosType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini1readqosType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini1readqosType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini1readqosType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini1readqosType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini1readqosFieldArqosShift = 0
	RegisterAxiini1readqosFieldArqosMask  = 0xf
)

// GetArqos Read channel QoS setting
func (r *RegisterAxiini1readqosType) GetArqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1readqosFieldArqosMask) >> RegisterAxiini1readqosFieldArqosShift)
}

// SetArqos Read channel QoS setting
func (r *RegisterAxiini1readqosType) SetArqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini1readqosFieldArqosMask)|(uint32(value)<<RegisterAxiini1readqosFieldArqosShift))
}

// RegisterAxiini1writeqosType AXI interconnect - INI x write QoS register
type RegisterAxiini1writeqosType uint32

func (r *RegisterAxiini1writeqosType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini1writeqosType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini1writeqosType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini1writeqosType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini1writeqosType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini1writeqosFieldAwqosShift = 0
	RegisterAxiini1writeqosFieldAwqosMask  = 0xf
)

// GetAwqos Write channel QoS setting
func (r *RegisterAxiini1writeqosType) GetAwqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1writeqosFieldAwqosMask) >> RegisterAxiini1writeqosFieldAwqosShift)
}

// SetAwqos Write channel QoS setting
func (r *RegisterAxiini1writeqosType) SetAwqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini1writeqosFieldAwqosMask)|(uint32(value)<<RegisterAxiini1writeqosFieldAwqosShift))
}

// RegisterAxiini1fnmodType AXI interconnect - INI x issuing functionality modification register
type RegisterAxiini1fnmodType uint32

func (r *RegisterAxiini1fnmodType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini1fnmodType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini1fnmodType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini1fnmodType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini1fnmodType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini1fnmodFieldReadissoverrideShift = 0
	RegisterAxiini1fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override ASIB read issuing capability
func (r *RegisterAxiini1fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override ASIB read issuing capability
func (r *RegisterAxiini1fnmodType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini1fnmodFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini1fnmodFieldReadissoverrideMask)
	}
}

const (
	RegisterAxiini1fnmodFieldWriteissoverrideShift = 1
	RegisterAxiini1fnmodFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Override ASIB write issuing capability
func (r *RegisterAxiini1fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override ASIB write issuing capability
func (r *RegisterAxiini1fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini1fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini1fnmodFieldWriteissoverrideMask)
	}
}

// RegisterAxiini2readqosType AXI interconnect - INI x read QoS register
type RegisterAxiini2readqosType uint32

func (r *RegisterAxiini2readqosType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini2readqosType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini2readqosType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini2readqosType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini2readqosType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini2readqosFieldArqosShift = 0
	RegisterAxiini2readqosFieldArqosMask  = 0xf
)

// GetArqos Read channel QoS setting
func (r *RegisterAxiini2readqosType) GetArqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini2readqosFieldArqosMask) >> RegisterAxiini2readqosFieldArqosShift)
}

// SetArqos Read channel QoS setting
func (r *RegisterAxiini2readqosType) SetArqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini2readqosFieldArqosMask)|(uint32(value)<<RegisterAxiini2readqosFieldArqosShift))
}

// RegisterAxiini2writeqosType AXI interconnect - INI x write QoS register
type RegisterAxiini2writeqosType uint32

func (r *RegisterAxiini2writeqosType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini2writeqosType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini2writeqosType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini2writeqosType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini2writeqosType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini2writeqosFieldAwqosShift = 0
	RegisterAxiini2writeqosFieldAwqosMask  = 0xf
)

// GetAwqos Write channel QoS setting
func (r *RegisterAxiini2writeqosType) GetAwqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini2writeqosFieldAwqosMask) >> RegisterAxiini2writeqosFieldAwqosShift)
}

// SetAwqos Write channel QoS setting
func (r *RegisterAxiini2writeqosType) SetAwqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini2writeqosFieldAwqosMask)|(uint32(value)<<RegisterAxiini2writeqosFieldAwqosShift))
}

// RegisterAxiini2fnmodType AXI interconnect - INI x issuing functionality modification register
type RegisterAxiini2fnmodType uint32

func (r *RegisterAxiini2fnmodType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini2fnmodType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini2fnmodType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini2fnmodType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini2fnmodType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini2fnmodFieldReadissoverrideShift = 0
	RegisterAxiini2fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override ASIB read issuing capability
func (r *RegisterAxiini2fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini2fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override ASIB read issuing capability
func (r *RegisterAxiini2fnmodType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini2fnmodFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini2fnmodFieldReadissoverrideMask)
	}
}

const (
	RegisterAxiini2fnmodFieldWriteissoverrideShift = 1
	RegisterAxiini2fnmodFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Override ASIB write issuing capability
func (r *RegisterAxiini2fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini2fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override ASIB write issuing capability
func (r *RegisterAxiini2fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini2fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini2fnmodFieldWriteissoverrideMask)
	}
}

// RegisterAxiini3fnmod2Type AXI interconnect - INI x functionality modification 2 register
type RegisterAxiini3fnmod2Type uint32

func (r *RegisterAxiini3fnmod2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini3fnmod2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini3fnmod2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini3fnmod2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini3fnmod2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini3fnmod2FieldBypassmergeShift = 0
	RegisterAxiini3fnmod2FieldBypassmergeMask  = 0x1
)

// GetBypassmerge Disables alteration of transactions by the up-sizer unless required by the protocol
func (r *RegisterAxiini3fnmod2Type) GetBypassmerge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3fnmod2FieldBypassmergeMask) != 0
}

// SetBypassmerge Disables alteration of transactions by the up-sizer unless required by the protocol
func (r *RegisterAxiini3fnmod2Type) SetBypassmerge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini3fnmod2FieldBypassmergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini3fnmod2FieldBypassmergeMask)
	}
}

// RegisterAxiini3fnmodahbType AXI interconnect - INI x AHB functionality modification register
type RegisterAxiini3fnmodahbType uint32

func (r *RegisterAxiini3fnmodahbType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini3fnmodahbType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini3fnmodahbType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini3fnmodahbType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini3fnmodahbType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini3fnmodahbFieldRdincoverrideShift = 0
	RegisterAxiini3fnmodahbFieldRdincoverrideMask  = 0x1
)

// GetRdincoverride Converts all AHB-Lite write transactions to a series of single beat AXI
func (r *RegisterAxiini3fnmodahbType) GetRdincoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3fnmodahbFieldRdincoverrideMask) != 0
}

// SetRdincoverride Converts all AHB-Lite write transactions to a series of single beat AXI
func (r *RegisterAxiini3fnmodahbType) SetRdincoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini3fnmodahbFieldRdincoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini3fnmodahbFieldRdincoverrideMask)
	}
}

const (
	RegisterAxiini3fnmodahbFieldWrincoverrideShift = 1
	RegisterAxiini3fnmodahbFieldWrincoverrideMask  = 0x2
)

// GetWrincoverride Converts all AHB-Lite read transactions to a series of single beat AXI
func (r *RegisterAxiini3fnmodahbType) GetWrincoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3fnmodahbFieldWrincoverrideMask) != 0
}

// SetWrincoverride Converts all AHB-Lite read transactions to a series of single beat AXI
func (r *RegisterAxiini3fnmodahbType) SetWrincoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini3fnmodahbFieldWrincoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini3fnmodahbFieldWrincoverrideMask)
	}
}

// RegisterAxiini3readqosType AXI interconnect - INI x read QoS register
type RegisterAxiini3readqosType uint32

func (r *RegisterAxiini3readqosType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini3readqosType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini3readqosType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini3readqosType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini3readqosType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini3readqosFieldArqosShift = 0
	RegisterAxiini3readqosFieldArqosMask  = 0xf
)

// GetArqos Read channel QoS setting
func (r *RegisterAxiini3readqosType) GetArqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3readqosFieldArqosMask) >> RegisterAxiini3readqosFieldArqosShift)
}

// SetArqos Read channel QoS setting
func (r *RegisterAxiini3readqosType) SetArqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini3readqosFieldArqosMask)|(uint32(value)<<RegisterAxiini3readqosFieldArqosShift))
}

// RegisterAxiini3writeqosType AXI interconnect - INI x write QoS register
type RegisterAxiini3writeqosType uint32

func (r *RegisterAxiini3writeqosType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini3writeqosType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini3writeqosType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini3writeqosType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini3writeqosType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini3writeqosFieldAwqosShift = 0
	RegisterAxiini3writeqosFieldAwqosMask  = 0xf
)

// GetAwqos Write channel QoS setting
func (r *RegisterAxiini3writeqosType) GetAwqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3writeqosFieldAwqosMask) >> RegisterAxiini3writeqosFieldAwqosShift)
}

// SetAwqos Write channel QoS setting
func (r *RegisterAxiini3writeqosType) SetAwqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini3writeqosFieldAwqosMask)|(uint32(value)<<RegisterAxiini3writeqosFieldAwqosShift))
}

// RegisterAxiini3fnmodType AXI interconnect - INI x issuing functionality modification register
type RegisterAxiini3fnmodType uint32

func (r *RegisterAxiini3fnmodType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini3fnmodType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini3fnmodType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini3fnmodType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini3fnmodType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini3fnmodFieldReadissoverrideShift = 0
	RegisterAxiini3fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override ASIB read issuing capability
func (r *RegisterAxiini3fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override ASIB read issuing capability
func (r *RegisterAxiini3fnmodType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini3fnmodFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini3fnmodFieldReadissoverrideMask)
	}
}

const (
	RegisterAxiini3fnmodFieldWriteissoverrideShift = 1
	RegisterAxiini3fnmodFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Override ASIB write issuing capability
func (r *RegisterAxiini3fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override ASIB write issuing capability
func (r *RegisterAxiini3fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini3fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini3fnmodFieldWriteissoverrideMask)
	}
}

// RegisterAxiini4readqosType AXI interconnect - INI x read QoS register
type RegisterAxiini4readqosType uint32

func (r *RegisterAxiini4readqosType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini4readqosType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini4readqosType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini4readqosType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini4readqosType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini4readqosFieldArqosShift = 0
	RegisterAxiini4readqosFieldArqosMask  = 0xf
)

// GetArqos Read channel QoS setting
func (r *RegisterAxiini4readqosType) GetArqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini4readqosFieldArqosMask) >> RegisterAxiini4readqosFieldArqosShift)
}

// SetArqos Read channel QoS setting
func (r *RegisterAxiini4readqosType) SetArqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini4readqosFieldArqosMask)|(uint32(value)<<RegisterAxiini4readqosFieldArqosShift))
}

// RegisterAxiini4writeqosType AXI interconnect - INI x write QoS register
type RegisterAxiini4writeqosType uint32

func (r *RegisterAxiini4writeqosType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini4writeqosType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini4writeqosType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini4writeqosType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini4writeqosType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini4writeqosFieldAwqosShift = 0
	RegisterAxiini4writeqosFieldAwqosMask  = 0xf
)

// GetAwqos Write channel QoS setting
func (r *RegisterAxiini4writeqosType) GetAwqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini4writeqosFieldAwqosMask) >> RegisterAxiini4writeqosFieldAwqosShift)
}

// SetAwqos Write channel QoS setting
func (r *RegisterAxiini4writeqosType) SetAwqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini4writeqosFieldAwqosMask)|(uint32(value)<<RegisterAxiini4writeqosFieldAwqosShift))
}

// RegisterAxiini4fnmodType AXI interconnect - INI x issuing functionality modification register
type RegisterAxiini4fnmodType uint32

func (r *RegisterAxiini4fnmodType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini4fnmodType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini4fnmodType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini4fnmodType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini4fnmodType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini4fnmodFieldReadissoverrideShift = 0
	RegisterAxiini4fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override ASIB read issuing capability
func (r *RegisterAxiini4fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini4fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override ASIB read issuing capability
func (r *RegisterAxiini4fnmodType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini4fnmodFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini4fnmodFieldReadissoverrideMask)
	}
}

const (
	RegisterAxiini4fnmodFieldWriteissoverrideShift = 1
	RegisterAxiini4fnmodFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Override ASIB write issuing capability
func (r *RegisterAxiini4fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini4fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override ASIB write issuing capability
func (r *RegisterAxiini4fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini4fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini4fnmodFieldWriteissoverrideMask)
	}
}

// RegisterAxiini5readqosType AXI interconnect - INI x read QoS register
type RegisterAxiini5readqosType uint32

func (r *RegisterAxiini5readqosType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini5readqosType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini5readqosType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini5readqosType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini5readqosType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini5readqosFieldArqosShift = 0
	RegisterAxiini5readqosFieldArqosMask  = 0xf
)

// GetArqos Read channel QoS setting
func (r *RegisterAxiini5readqosType) GetArqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini5readqosFieldArqosMask) >> RegisterAxiini5readqosFieldArqosShift)
}

// SetArqos Read channel QoS setting
func (r *RegisterAxiini5readqosType) SetArqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini5readqosFieldArqosMask)|(uint32(value)<<RegisterAxiini5readqosFieldArqosShift))
}

// RegisterAxiini5writeqosType AXI interconnect - INI x write QoS register
type RegisterAxiini5writeqosType uint32

func (r *RegisterAxiini5writeqosType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini5writeqosType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini5writeqosType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini5writeqosType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini5writeqosType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini5writeqosFieldAwqosShift = 0
	RegisterAxiini5writeqosFieldAwqosMask  = 0xf
)

// GetAwqos Write channel QoS setting
func (r *RegisterAxiini5writeqosType) GetAwqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini5writeqosFieldAwqosMask) >> RegisterAxiini5writeqosFieldAwqosShift)
}

// SetAwqos Write channel QoS setting
func (r *RegisterAxiini5writeqosType) SetAwqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini5writeqosFieldAwqosMask)|(uint32(value)<<RegisterAxiini5writeqosFieldAwqosShift))
}

// RegisterAxiini5fnmodType AXI interconnect - INI x issuing functionality modification register
type RegisterAxiini5fnmodType uint32

func (r *RegisterAxiini5fnmodType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini5fnmodType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini5fnmodType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini5fnmodType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini5fnmodType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini5fnmodFieldReadissoverrideShift = 0
	RegisterAxiini5fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override ASIB read issuing capability
func (r *RegisterAxiini5fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini5fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override ASIB read issuing capability
func (r *RegisterAxiini5fnmodType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini5fnmodFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini5fnmodFieldReadissoverrideMask)
	}
}

const (
	RegisterAxiini5fnmodFieldWriteissoverrideShift = 1
	RegisterAxiini5fnmodFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Override ASIB write issuing capability
func (r *RegisterAxiini5fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini5fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override ASIB write issuing capability
func (r *RegisterAxiini5fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini5fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini5fnmodFieldWriteissoverrideMask)
	}
}

// RegisterAxiini6readqosType AXI interconnect - INI x read QoS register
type RegisterAxiini6readqosType uint32

func (r *RegisterAxiini6readqosType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini6readqosType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini6readqosType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini6readqosType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini6readqosType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini6readqosFieldArqosShift = 0
	RegisterAxiini6readqosFieldArqosMask  = 0xf
)

// GetArqos Read channel QoS setting
func (r *RegisterAxiini6readqosType) GetArqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini6readqosFieldArqosMask) >> RegisterAxiini6readqosFieldArqosShift)
}

// SetArqos Read channel QoS setting
func (r *RegisterAxiini6readqosType) SetArqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini6readqosFieldArqosMask)|(uint32(value)<<RegisterAxiini6readqosFieldArqosShift))
}

// RegisterAxiini6writeqosType AXI interconnect - INI x write QoS register
type RegisterAxiini6writeqosType uint32

func (r *RegisterAxiini6writeqosType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini6writeqosType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini6writeqosType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini6writeqosType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini6writeqosType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini6writeqosFieldAwqosShift = 0
	RegisterAxiini6writeqosFieldAwqosMask  = 0xf
)

// GetAwqos Write channel QoS setting
func (r *RegisterAxiini6writeqosType) GetAwqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini6writeqosFieldAwqosMask) >> RegisterAxiini6writeqosFieldAwqosShift)
}

// SetAwqos Write channel QoS setting
func (r *RegisterAxiini6writeqosType) SetAwqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini6writeqosFieldAwqosMask)|(uint32(value)<<RegisterAxiini6writeqosFieldAwqosShift))
}

// RegisterAxiini6fnmodType AXI interconnect - INI x issuing functionality modification register
type RegisterAxiini6fnmodType uint32

func (r *RegisterAxiini6fnmodType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAxiini6fnmodType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAxiini6fnmodType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAxiini6fnmodType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAxiini6fnmodType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAxiini6fnmodFieldReadissoverrideShift = 0
	RegisterAxiini6fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override ASIB read issuing capability
func (r *RegisterAxiini6fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini6fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override ASIB read issuing capability
func (r *RegisterAxiini6fnmodType) SetReadissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini6fnmodFieldReadissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini6fnmodFieldReadissoverrideMask)
	}
}

const (
	RegisterAxiini6fnmodFieldWriteissoverrideShift = 1
	RegisterAxiini6fnmodFieldWriteissoverrideMask  = 0x2
)

// GetWriteissoverride Override ASIB write issuing capability
func (r *RegisterAxiini6fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini6fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override ASIB write issuing capability
func (r *RegisterAxiini6fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini6fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini6fnmodFieldWriteissoverrideMask)
	}
}
