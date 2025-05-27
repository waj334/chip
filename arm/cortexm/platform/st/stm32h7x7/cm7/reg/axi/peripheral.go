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
	Axiperiphid4       registerAxiperiphid4Type
	_                  [12]byte
	Axiperiphid0       registerAxiperiphid0Type
	Axiperiphid1       registerAxiperiphid1Type
	Axiperiphid2       registerAxiperiphid2Type
	Axiperiphid3       registerAxiperiphid3Type
	Axicompid0         registerAxicompid0Type
	Axicompid1         registerAxicompid1Type
	Axicompid2         registerAxicompid2Type
	Axicompid3         registerAxicompid3Type
	_                  [8]byte
	Axitarg1fnmodissbm registerAxitarg1fnmodissbmType
	_                  [24]byte
	Axitarg1fnmod2     registerAxitarg1fnmod2Type
	_                  [4]byte
	Axitarg1fnmodlb    registerAxitarg1fnmodlbType
	_                  [216]byte
	Axitarg1fnmod      registerAxitarg1fnmodType
	_                  [3836]byte
	Axitarg2fnmodissbm registerAxitarg2fnmodissbmType
	_                  [24]byte
	Axitarg2fnmod2     registerAxitarg2fnmod2Type
	_                  [4]byte
	Axitarg2fnmodlb    registerAxitarg2fnmodlbType
	_                  [216]byte
	Axitarg2fnmod      registerAxitarg2fnmodType
	_                  [3836]byte
	Axitarg3fnmodissbm registerAxitarg3fnmodissbmType
	_                  [4092]byte
	Axitarg4fnmodissbm registerAxitarg4fnmodissbmType
	_                  [4092]byte
	Axitarg5fnmodissbm registerAxitarg5fnmodissbmType
	_                  [4092]byte
	Axitarg6fnmodissbm registerAxitarg6fnmodissbmType
	_                  [4096]byte
	Axitarg7fnmodissbm registerAxitarg7fnmodissbmType
	_                  [20]byte
	Axitarg7fnmod2     registerAxitarg7fnmod2Type
	_                  [224]byte
	Axitarg7fnmod      registerAxitarg7fnmodType
	_                  [237336]byte
	Axiini1fnmod2      registerAxiini1fnmod2Type
	Axiini1fnmodahb    registerAxiini1fnmodahbType
	_                  [212]byte
	Axiini1readqos     registerAxiini1readqosType
	Axiini1writeqos    registerAxiini1writeqosType
	Axiini1fnmod       registerAxiini1fnmodType
	_                  [4084]byte
	Axiini2readqos     registerAxiini2readqosType
	Axiini2writeqos    registerAxiini2writeqosType
	Axiini2fnmod       registerAxiini2fnmodType
	_                  [3864]byte
	Axiini3fnmod2      registerAxiini3fnmod2Type
	Axiini3fnmodahb    registerAxiini3fnmodahbType
	_                  [212]byte
	Axiini3readqos     registerAxiini3readqosType
	Axiini3writeqos    registerAxiini3writeqosType
	Axiini3fnmod       registerAxiini3fnmodType
	_                  [4084]byte
	Axiini4readqos     registerAxiini4readqosType
	Axiini4writeqos    registerAxiini4writeqosType
	Axiini4fnmod       registerAxiini4fnmodType
	_                  [4084]byte
	Axiini5readqos     registerAxiini5readqosType
	Axiini5writeqos    registerAxiini5writeqosType
	Axiini5fnmod       registerAxiini5fnmodType
	_                  [4084]byte
	Axiini6readqos     registerAxiini6readqosType
	Axiini6writeqos    registerAxiini6writeqosType
	Axiini6fnmod       registerAxiini6fnmodType
}

// registerAxiperiphid4Type AXI interconnect - peripheral ID4 register
type registerAxiperiphid4Type uint32

const (
	RegisterAxiperiphid4FieldJep106conShift = 0
	RegisterAxiperiphid4FieldJep106conMask  = 0xf
)

// GetJep106con JEP106 continuation code
func (r *registerAxiperiphid4Type) GetJep106con() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid4FieldJep106conMask) >> RegisterAxiperiphid4FieldJep106conShift)
}

// SetJep106con JEP106 continuation code
func (r *registerAxiperiphid4Type) SetJep106con(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid4FieldJep106conMask)|(uint32(value)<<RegisterAxiperiphid4FieldJep106conShift))
}

const (
	RegisterAxiperiphid4FieldKcount4Shift = 4
	RegisterAxiperiphid4FieldKcount4Mask  = 0xf0
)

// GetKcount4 Register file size
func (r *registerAxiperiphid4Type) GetKcount4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid4FieldKcount4Mask) >> RegisterAxiperiphid4FieldKcount4Shift)
}

// SetKcount4 Register file size
func (r *registerAxiperiphid4Type) SetKcount4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid4FieldKcount4Mask)|(uint32(value)<<RegisterAxiperiphid4FieldKcount4Shift))
}

// registerAxiperiphid0Type AXI interconnect - peripheral ID0 register
type registerAxiperiphid0Type uint32

const (
	RegisterAxiperiphid0FieldPartnumShift = 0
	RegisterAxiperiphid0FieldPartnumMask  = 0xff
)

// GetPartnum Peripheral part number bits 0 to 7
func (r *registerAxiperiphid0Type) GetPartnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid0FieldPartnumMask) >> RegisterAxiperiphid0FieldPartnumShift)
}

// SetPartnum Peripheral part number bits 0 to 7
func (r *registerAxiperiphid0Type) SetPartnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid0FieldPartnumMask)|(uint32(value)<<RegisterAxiperiphid0FieldPartnumShift))
}

// registerAxiperiphid1Type AXI interconnect - peripheral ID1 register
type registerAxiperiphid1Type uint32

const (
	RegisterAxiperiphid1FieldPartnumShift = 0
	RegisterAxiperiphid1FieldPartnumMask  = 0xf
)

// GetPartnum Peripheral part number bits 8 to 11
func (r *registerAxiperiphid1Type) GetPartnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid1FieldPartnumMask) >> RegisterAxiperiphid1FieldPartnumShift)
}

// SetPartnum Peripheral part number bits 8 to 11
func (r *registerAxiperiphid1Type) SetPartnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid1FieldPartnumMask)|(uint32(value)<<RegisterAxiperiphid1FieldPartnumShift))
}

const (
	RegisterAxiperiphid1FieldJep106iShift = 4
	RegisterAxiperiphid1FieldJep106iMask  = 0xf0
)

// GetJep106i JEP106 identity bits 0 to 3
func (r *registerAxiperiphid1Type) GetJep106i() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid1FieldJep106iMask) >> RegisterAxiperiphid1FieldJep106iShift)
}

// SetJep106i JEP106 identity bits 0 to 3
func (r *registerAxiperiphid1Type) SetJep106i(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid1FieldJep106iMask)|(uint32(value)<<RegisterAxiperiphid1FieldJep106iShift))
}

// registerAxiperiphid2Type AXI interconnect - peripheral ID2 register
type registerAxiperiphid2Type uint32

const (
	RegisterAxiperiphid2FieldJep106idShift = 0
	RegisterAxiperiphid2FieldJep106idMask  = 0x7
)

// GetJep106id JEP106 Identity bits 4 to 6
func (r *registerAxiperiphid2Type) GetJep106id() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid2FieldJep106idMask) >> RegisterAxiperiphid2FieldJep106idShift)
}

// SetJep106id JEP106 Identity bits 4 to 6
func (r *registerAxiperiphid2Type) SetJep106id(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid2FieldJep106idMask)|(uint32(value)<<RegisterAxiperiphid2FieldJep106idShift))
}

const (
	RegisterAxiperiphid2FieldJedecShift = 3
	RegisterAxiperiphid2FieldJedecMask  = 0x8
)

// GetJedec JEP106 code flag
func (r *registerAxiperiphid2Type) GetJedec() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid2FieldJedecMask) != 0
}

// SetJedec JEP106 code flag
func (r *registerAxiperiphid2Type) SetJedec(value bool) {
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
func (r *registerAxiperiphid2Type) GetRevision() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid2FieldRevisionMask) >> RegisterAxiperiphid2FieldRevisionShift)
}

// SetRevision Peripheral revision number
func (r *registerAxiperiphid2Type) SetRevision(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid2FieldRevisionMask)|(uint32(value)<<RegisterAxiperiphid2FieldRevisionShift))
}

// registerAxiperiphid3Type AXI interconnect - peripheral ID3 register
type registerAxiperiphid3Type uint32

const (
	RegisterAxiperiphid3FieldCustmodnumShift = 0
	RegisterAxiperiphid3FieldCustmodnumMask  = 0xf
)

// GetCustmodnum Customer modification
func (r *registerAxiperiphid3Type) GetCustmodnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid3FieldCustmodnumMask) >> RegisterAxiperiphid3FieldCustmodnumShift)
}

// SetCustmodnum Customer modification
func (r *registerAxiperiphid3Type) SetCustmodnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid3FieldCustmodnumMask)|(uint32(value)<<RegisterAxiperiphid3FieldCustmodnumShift))
}

const (
	RegisterAxiperiphid3FieldRevandShift = 4
	RegisterAxiperiphid3FieldRevandMask  = 0xf0
)

// GetRevand Customer version
func (r *registerAxiperiphid3Type) GetRevand() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiperiphid3FieldRevandMask) >> RegisterAxiperiphid3FieldRevandShift)
}

// SetRevand Customer version
func (r *registerAxiperiphid3Type) SetRevand(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiperiphid3FieldRevandMask)|(uint32(value)<<RegisterAxiperiphid3FieldRevandShift))
}

// registerAxicompid0Type AXI interconnect - component ID0 register
type registerAxicompid0Type uint32

const (
	RegisterAxicompid0FieldPreambleShift = 0
	RegisterAxicompid0FieldPreambleMask  = 0xff
)

// GetPreamble Preamble bits 0 to 7
func (r *registerAxicompid0Type) GetPreamble() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxicompid0FieldPreambleMask) >> RegisterAxicompid0FieldPreambleShift)
}

// SetPreamble Preamble bits 0 to 7
func (r *registerAxicompid0Type) SetPreamble(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxicompid0FieldPreambleMask)|(uint32(value)<<RegisterAxicompid0FieldPreambleShift))
}

// registerAxicompid1Type AXI interconnect - component ID1 register
type registerAxicompid1Type uint32

const (
	RegisterAxicompid1FieldPreambleShift = 0
	RegisterAxicompid1FieldPreambleMask  = 0xf
)

// GetPreamble Preamble bits 8 to 11
func (r *registerAxicompid1Type) GetPreamble() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxicompid1FieldPreambleMask) >> RegisterAxicompid1FieldPreambleShift)
}

// SetPreamble Preamble bits 8 to 11
func (r *registerAxicompid1Type) SetPreamble(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxicompid1FieldPreambleMask)|(uint32(value)<<RegisterAxicompid1FieldPreambleShift))
}

const (
	RegisterAxicompid1FieldClassShift = 4
	RegisterAxicompid1FieldClassMask  = 0xf0
)

// GetClass Component class
func (r *registerAxicompid1Type) GetClass() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxicompid1FieldClassMask) >> RegisterAxicompid1FieldClassShift)
}

// SetClass Component class
func (r *registerAxicompid1Type) SetClass(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxicompid1FieldClassMask)|(uint32(value)<<RegisterAxicompid1FieldClassShift))
}

// registerAxicompid2Type AXI interconnect - component ID2 register
type registerAxicompid2Type uint32

const (
	RegisterAxicompid2FieldPreambleShift = 0
	RegisterAxicompid2FieldPreambleMask  = 0xff
)

// GetPreamble Preamble bits 12 to 19
func (r *registerAxicompid2Type) GetPreamble() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxicompid2FieldPreambleMask) >> RegisterAxicompid2FieldPreambleShift)
}

// SetPreamble Preamble bits 12 to 19
func (r *registerAxicompid2Type) SetPreamble(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxicompid2FieldPreambleMask)|(uint32(value)<<RegisterAxicompid2FieldPreambleShift))
}

// registerAxicompid3Type AXI interconnect - component ID3 register
type registerAxicompid3Type uint32

const (
	RegisterAxicompid3FieldPreambleShift = 0
	RegisterAxicompid3FieldPreambleMask  = 0xff
)

// GetPreamble Preamble bits 20 to 27
func (r *registerAxicompid3Type) GetPreamble() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxicompid3FieldPreambleMask) >> RegisterAxicompid3FieldPreambleShift)
}

// SetPreamble Preamble bits 20 to 27
func (r *registerAxicompid3Type) SetPreamble(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxicompid3FieldPreambleMask)|(uint32(value)<<RegisterAxicompid3FieldPreambleShift))
}

// registerAxitarg1fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxitarg1fnmodissbmType uint32

const (
	RegisterAxitarg1fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg1fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg1fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg1fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg1fnmodissbmType) SetReadissoverride(value bool) {
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
func (r *registerAxitarg1fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg1fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *registerAxitarg1fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg1fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg1fnmodissbmFieldWriteissoverrideMask)
	}
}

// registerAxitarg1fnmod2Type AXI interconnect - TARG x bus matrix functionality 2 register
type registerAxitarg1fnmod2Type uint32

const (
	RegisterAxitarg1fnmod2FieldBypassmergeShift = 0
	RegisterAxitarg1fnmod2FieldBypassmergeMask  = 0x1
)

// GetBypassmerge Disable packing of beats to match the output data width
func (r *registerAxitarg1fnmod2Type) GetBypassmerge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg1fnmod2FieldBypassmergeMask) != 0
}

// SetBypassmerge Disable packing of beats to match the output data width
func (r *registerAxitarg1fnmod2Type) SetBypassmerge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg1fnmod2FieldBypassmergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg1fnmod2FieldBypassmergeMask)
	}
}

// registerAxitarg1fnmodlbType AXI interconnect - TARG x long burst functionality modification
type registerAxitarg1fnmodlbType uint32

const (
	RegisterAxitarg1fnmodlbFieldFnmodlbShift = 0
	RegisterAxitarg1fnmodlbFieldFnmodlbMask  = 0x1
)

// GetFnmodlb Controls burst breaking of long bursts
func (r *registerAxitarg1fnmodlbType) GetFnmodlb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg1fnmodlbFieldFnmodlbMask) != 0
}

// SetFnmodlb Controls burst breaking of long bursts
func (r *registerAxitarg1fnmodlbType) SetFnmodlb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg1fnmodlbFieldFnmodlbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg1fnmodlbFieldFnmodlbMask)
	}
}

// registerAxitarg1fnmodType AXI interconnect - TARG x long burst functionality modification
type registerAxitarg1fnmodType uint32

const (
	RegisterAxitarg1fnmodFieldReadissoverrideShift = 0
	RegisterAxitarg1fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override AMIB read issuing capability
func (r *registerAxitarg1fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg1fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override AMIB read issuing capability
func (r *registerAxitarg1fnmodType) SetReadissoverride(value bool) {
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
func (r *registerAxitarg1fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg1fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override AMIB write issuing capability
func (r *registerAxitarg1fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg1fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg1fnmodFieldWriteissoverrideMask)
	}
}

// registerAxitarg2fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxitarg2fnmodissbmType uint32

const (
	RegisterAxitarg2fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg2fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg2fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg2fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg2fnmodissbmType) SetReadissoverride(value bool) {
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
func (r *registerAxitarg2fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg2fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *registerAxitarg2fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg2fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg2fnmodissbmFieldWriteissoverrideMask)
	}
}

// registerAxitarg2fnmod2Type AXI interconnect - TARG x bus matrix functionality 2 register
type registerAxitarg2fnmod2Type uint32

const (
	RegisterAxitarg2fnmod2FieldBypassmergeShift = 0
	RegisterAxitarg2fnmod2FieldBypassmergeMask  = 0x1
)

// GetBypassmerge Disable packing of beats to match the output data width
func (r *registerAxitarg2fnmod2Type) GetBypassmerge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg2fnmod2FieldBypassmergeMask) != 0
}

// SetBypassmerge Disable packing of beats to match the output data width
func (r *registerAxitarg2fnmod2Type) SetBypassmerge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg2fnmod2FieldBypassmergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg2fnmod2FieldBypassmergeMask)
	}
}

// registerAxitarg2fnmodlbType AXI interconnect - TARG x long burst functionality modification
type registerAxitarg2fnmodlbType uint32

const (
	RegisterAxitarg2fnmodlbFieldFnmodlbShift = 0
	RegisterAxitarg2fnmodlbFieldFnmodlbMask  = 0x1
)

// GetFnmodlb Controls burst breaking of long bursts
func (r *registerAxitarg2fnmodlbType) GetFnmodlb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg2fnmodlbFieldFnmodlbMask) != 0
}

// SetFnmodlb Controls burst breaking of long bursts
func (r *registerAxitarg2fnmodlbType) SetFnmodlb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg2fnmodlbFieldFnmodlbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg2fnmodlbFieldFnmodlbMask)
	}
}

// registerAxitarg2fnmodType AXI interconnect - TARG x long burst functionality modification
type registerAxitarg2fnmodType uint32

const (
	RegisterAxitarg2fnmodFieldReadissoverrideShift = 0
	RegisterAxitarg2fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override AMIB read issuing capability
func (r *registerAxitarg2fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg2fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override AMIB read issuing capability
func (r *registerAxitarg2fnmodType) SetReadissoverride(value bool) {
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
func (r *registerAxitarg2fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg2fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override AMIB write issuing capability
func (r *registerAxitarg2fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg2fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg2fnmodFieldWriteissoverrideMask)
	}
}

// registerAxitarg3fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxitarg3fnmodissbmType uint32

const (
	RegisterAxitarg3fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg3fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg3fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg3fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg3fnmodissbmType) SetReadissoverride(value bool) {
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
func (r *registerAxitarg3fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg3fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *registerAxitarg3fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg3fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg3fnmodissbmFieldWriteissoverrideMask)
	}
}

// registerAxitarg4fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxitarg4fnmodissbmType uint32

const (
	RegisterAxitarg4fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg4fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg4fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg4fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg4fnmodissbmType) SetReadissoverride(value bool) {
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
func (r *registerAxitarg4fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg4fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *registerAxitarg4fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg4fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg4fnmodissbmFieldWriteissoverrideMask)
	}
}

// registerAxitarg5fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxitarg5fnmodissbmType uint32

const (
	RegisterAxitarg5fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg5fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg5fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg5fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg5fnmodissbmType) SetReadissoverride(value bool) {
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
func (r *registerAxitarg5fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg5fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *registerAxitarg5fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg5fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg5fnmodissbmFieldWriteissoverrideMask)
	}
}

// registerAxitarg6fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxitarg6fnmodissbmType uint32

const (
	RegisterAxitarg6fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg6fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg6fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg6fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg6fnmodissbmType) SetReadissoverride(value bool) {
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
func (r *registerAxitarg6fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg6fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *registerAxitarg6fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg6fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg6fnmodissbmFieldWriteissoverrideMask)
	}
}

// registerAxitarg7fnmodissbmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxitarg7fnmodissbmType uint32

const (
	RegisterAxitarg7fnmodissbmFieldReadissoverrideShift = 0
	RegisterAxitarg7fnmodissbmFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg7fnmodissbmType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg7fnmodissbmFieldReadissoverrideMask) != 0
}

// SetReadissoverride READ_ISS_OVERRIDE
func (r *registerAxitarg7fnmodissbmType) SetReadissoverride(value bool) {
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
func (r *registerAxitarg7fnmodissbmType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg7fnmodissbmFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Switch matrix write issuing override for target
func (r *registerAxitarg7fnmodissbmType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg7fnmodissbmFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg7fnmodissbmFieldWriteissoverrideMask)
	}
}

// registerAxitarg7fnmod2Type AXI interconnect - TARG x bus matrix functionality 2 register
type registerAxitarg7fnmod2Type uint32

const (
	RegisterAxitarg7fnmod2FieldBypassmergeShift = 0
	RegisterAxitarg7fnmod2FieldBypassmergeMask  = 0x1
)

// GetBypassmerge Disable packing of beats to match the output data width
func (r *registerAxitarg7fnmod2Type) GetBypassmerge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg7fnmod2FieldBypassmergeMask) != 0
}

// SetBypassmerge Disable packing of beats to match the output data width
func (r *registerAxitarg7fnmod2Type) SetBypassmerge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg7fnmod2FieldBypassmergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg7fnmod2FieldBypassmergeMask)
	}
}

// registerAxitarg7fnmodType AXI interconnect - TARG x long burst functionality modification
type registerAxitarg7fnmodType uint32

const (
	RegisterAxitarg7fnmodFieldReadissoverrideShift = 0
	RegisterAxitarg7fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override AMIB read issuing capability
func (r *registerAxitarg7fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg7fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override AMIB read issuing capability
func (r *registerAxitarg7fnmodType) SetReadissoverride(value bool) {
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
func (r *registerAxitarg7fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxitarg7fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override AMIB write issuing capability
func (r *registerAxitarg7fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxitarg7fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxitarg7fnmodFieldWriteissoverrideMask)
	}
}

// registerAxiini1fnmod2Type AXI interconnect - INI x functionality modification 2 register
type registerAxiini1fnmod2Type uint32

const (
	RegisterAxiini1fnmod2FieldBypassmergeShift = 0
	RegisterAxiini1fnmod2FieldBypassmergeMask  = 0x1
)

// GetBypassmerge Disables alteration of transactions by the up-sizer unless required by the protocol
func (r *registerAxiini1fnmod2Type) GetBypassmerge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1fnmod2FieldBypassmergeMask) != 0
}

// SetBypassmerge Disables alteration of transactions by the up-sizer unless required by the protocol
func (r *registerAxiini1fnmod2Type) SetBypassmerge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini1fnmod2FieldBypassmergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini1fnmod2FieldBypassmergeMask)
	}
}

// registerAxiini1fnmodahbType AXI interconnect - INI x AHB functionality modification register
type registerAxiini1fnmodahbType uint32

const (
	RegisterAxiini1fnmodahbFieldRdincoverrideShift = 0
	RegisterAxiini1fnmodahbFieldRdincoverrideMask  = 0x1
)

// GetRdincoverride Converts all AHB-Lite write transactions to a series of single beat AXI
func (r *registerAxiini1fnmodahbType) GetRdincoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1fnmodahbFieldRdincoverrideMask) != 0
}

// SetRdincoverride Converts all AHB-Lite write transactions to a series of single beat AXI
func (r *registerAxiini1fnmodahbType) SetRdincoverride(value bool) {
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
func (r *registerAxiini1fnmodahbType) GetWrincoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1fnmodahbFieldWrincoverrideMask) != 0
}

// SetWrincoverride Converts all AHB-Lite read transactions to a series of single beat AXI
func (r *registerAxiini1fnmodahbType) SetWrincoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini1fnmodahbFieldWrincoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini1fnmodahbFieldWrincoverrideMask)
	}
}

// registerAxiini1readqosType AXI interconnect - INI x read QoS register
type registerAxiini1readqosType uint32

const (
	RegisterAxiini1readqosFieldArqosShift = 0
	RegisterAxiini1readqosFieldArqosMask  = 0xf
)

// GetArqos Read channel QoS setting
func (r *registerAxiini1readqosType) GetArqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1readqosFieldArqosMask) >> RegisterAxiini1readqosFieldArqosShift)
}

// SetArqos Read channel QoS setting
func (r *registerAxiini1readqosType) SetArqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini1readqosFieldArqosMask)|(uint32(value)<<RegisterAxiini1readqosFieldArqosShift))
}

// registerAxiini1writeqosType AXI interconnect - INI x write QoS register
type registerAxiini1writeqosType uint32

const (
	RegisterAxiini1writeqosFieldAwqosShift = 0
	RegisterAxiini1writeqosFieldAwqosMask  = 0xf
)

// GetAwqos Write channel QoS setting
func (r *registerAxiini1writeqosType) GetAwqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1writeqosFieldAwqosMask) >> RegisterAxiini1writeqosFieldAwqosShift)
}

// SetAwqos Write channel QoS setting
func (r *registerAxiini1writeqosType) SetAwqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini1writeqosFieldAwqosMask)|(uint32(value)<<RegisterAxiini1writeqosFieldAwqosShift))
}

// registerAxiini1fnmodType AXI interconnect - INI x issuing functionality modification register
type registerAxiini1fnmodType uint32

const (
	RegisterAxiini1fnmodFieldReadissoverrideShift = 0
	RegisterAxiini1fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override ASIB read issuing capability
func (r *registerAxiini1fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override ASIB read issuing capability
func (r *registerAxiini1fnmodType) SetReadissoverride(value bool) {
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
func (r *registerAxiini1fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini1fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override ASIB write issuing capability
func (r *registerAxiini1fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini1fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini1fnmodFieldWriteissoverrideMask)
	}
}

// registerAxiini2readqosType AXI interconnect - INI x read QoS register
type registerAxiini2readqosType uint32

const (
	RegisterAxiini2readqosFieldArqosShift = 0
	RegisterAxiini2readqosFieldArqosMask  = 0xf
)

// GetArqos Read channel QoS setting
func (r *registerAxiini2readqosType) GetArqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini2readqosFieldArqosMask) >> RegisterAxiini2readqosFieldArqosShift)
}

// SetArqos Read channel QoS setting
func (r *registerAxiini2readqosType) SetArqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini2readqosFieldArqosMask)|(uint32(value)<<RegisterAxiini2readqosFieldArqosShift))
}

// registerAxiini2writeqosType AXI interconnect - INI x write QoS register
type registerAxiini2writeqosType uint32

const (
	RegisterAxiini2writeqosFieldAwqosShift = 0
	RegisterAxiini2writeqosFieldAwqosMask  = 0xf
)

// GetAwqos Write channel QoS setting
func (r *registerAxiini2writeqosType) GetAwqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini2writeqosFieldAwqosMask) >> RegisterAxiini2writeqosFieldAwqosShift)
}

// SetAwqos Write channel QoS setting
func (r *registerAxiini2writeqosType) SetAwqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini2writeqosFieldAwqosMask)|(uint32(value)<<RegisterAxiini2writeqosFieldAwqosShift))
}

// registerAxiini2fnmodType AXI interconnect - INI x issuing functionality modification register
type registerAxiini2fnmodType uint32

const (
	RegisterAxiini2fnmodFieldReadissoverrideShift = 0
	RegisterAxiini2fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override ASIB read issuing capability
func (r *registerAxiini2fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini2fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override ASIB read issuing capability
func (r *registerAxiini2fnmodType) SetReadissoverride(value bool) {
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
func (r *registerAxiini2fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini2fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override ASIB write issuing capability
func (r *registerAxiini2fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini2fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini2fnmodFieldWriteissoverrideMask)
	}
}

// registerAxiini3fnmod2Type AXI interconnect - INI x functionality modification 2 register
type registerAxiini3fnmod2Type uint32

const (
	RegisterAxiini3fnmod2FieldBypassmergeShift = 0
	RegisterAxiini3fnmod2FieldBypassmergeMask  = 0x1
)

// GetBypassmerge Disables alteration of transactions by the up-sizer unless required by the protocol
func (r *registerAxiini3fnmod2Type) GetBypassmerge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3fnmod2FieldBypassmergeMask) != 0
}

// SetBypassmerge Disables alteration of transactions by the up-sizer unless required by the protocol
func (r *registerAxiini3fnmod2Type) SetBypassmerge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini3fnmod2FieldBypassmergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini3fnmod2FieldBypassmergeMask)
	}
}

// registerAxiini3fnmodahbType AXI interconnect - INI x AHB functionality modification register
type registerAxiini3fnmodahbType uint32

const (
	RegisterAxiini3fnmodahbFieldRdincoverrideShift = 0
	RegisterAxiini3fnmodahbFieldRdincoverrideMask  = 0x1
)

// GetRdincoverride Converts all AHB-Lite write transactions to a series of single beat AXI
func (r *registerAxiini3fnmodahbType) GetRdincoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3fnmodahbFieldRdincoverrideMask) != 0
}

// SetRdincoverride Converts all AHB-Lite write transactions to a series of single beat AXI
func (r *registerAxiini3fnmodahbType) SetRdincoverride(value bool) {
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
func (r *registerAxiini3fnmodahbType) GetWrincoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3fnmodahbFieldWrincoverrideMask) != 0
}

// SetWrincoverride Converts all AHB-Lite read transactions to a series of single beat AXI
func (r *registerAxiini3fnmodahbType) SetWrincoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini3fnmodahbFieldWrincoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini3fnmodahbFieldWrincoverrideMask)
	}
}

// registerAxiini3readqosType AXI interconnect - INI x read QoS register
type registerAxiini3readqosType uint32

const (
	RegisterAxiini3readqosFieldArqosShift = 0
	RegisterAxiini3readqosFieldArqosMask  = 0xf
)

// GetArqos Read channel QoS setting
func (r *registerAxiini3readqosType) GetArqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3readqosFieldArqosMask) >> RegisterAxiini3readqosFieldArqosShift)
}

// SetArqos Read channel QoS setting
func (r *registerAxiini3readqosType) SetArqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini3readqosFieldArqosMask)|(uint32(value)<<RegisterAxiini3readqosFieldArqosShift))
}

// registerAxiini3writeqosType AXI interconnect - INI x write QoS register
type registerAxiini3writeqosType uint32

const (
	RegisterAxiini3writeqosFieldAwqosShift = 0
	RegisterAxiini3writeqosFieldAwqosMask  = 0xf
)

// GetAwqos Write channel QoS setting
func (r *registerAxiini3writeqosType) GetAwqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3writeqosFieldAwqosMask) >> RegisterAxiini3writeqosFieldAwqosShift)
}

// SetAwqos Write channel QoS setting
func (r *registerAxiini3writeqosType) SetAwqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini3writeqosFieldAwqosMask)|(uint32(value)<<RegisterAxiini3writeqosFieldAwqosShift))
}

// registerAxiini3fnmodType AXI interconnect - INI x issuing functionality modification register
type registerAxiini3fnmodType uint32

const (
	RegisterAxiini3fnmodFieldReadissoverrideShift = 0
	RegisterAxiini3fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override ASIB read issuing capability
func (r *registerAxiini3fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override ASIB read issuing capability
func (r *registerAxiini3fnmodType) SetReadissoverride(value bool) {
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
func (r *registerAxiini3fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini3fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override ASIB write issuing capability
func (r *registerAxiini3fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini3fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini3fnmodFieldWriteissoverrideMask)
	}
}

// registerAxiini4readqosType AXI interconnect - INI x read QoS register
type registerAxiini4readqosType uint32

const (
	RegisterAxiini4readqosFieldArqosShift = 0
	RegisterAxiini4readqosFieldArqosMask  = 0xf
)

// GetArqos Read channel QoS setting
func (r *registerAxiini4readqosType) GetArqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini4readqosFieldArqosMask) >> RegisterAxiini4readqosFieldArqosShift)
}

// SetArqos Read channel QoS setting
func (r *registerAxiini4readqosType) SetArqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini4readqosFieldArqosMask)|(uint32(value)<<RegisterAxiini4readqosFieldArqosShift))
}

// registerAxiini4writeqosType AXI interconnect - INI x write QoS register
type registerAxiini4writeqosType uint32

const (
	RegisterAxiini4writeqosFieldAwqosShift = 0
	RegisterAxiini4writeqosFieldAwqosMask  = 0xf
)

// GetAwqos Write channel QoS setting
func (r *registerAxiini4writeqosType) GetAwqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini4writeqosFieldAwqosMask) >> RegisterAxiini4writeqosFieldAwqosShift)
}

// SetAwqos Write channel QoS setting
func (r *registerAxiini4writeqosType) SetAwqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini4writeqosFieldAwqosMask)|(uint32(value)<<RegisterAxiini4writeqosFieldAwqosShift))
}

// registerAxiini4fnmodType AXI interconnect - INI x issuing functionality modification register
type registerAxiini4fnmodType uint32

const (
	RegisterAxiini4fnmodFieldReadissoverrideShift = 0
	RegisterAxiini4fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override ASIB read issuing capability
func (r *registerAxiini4fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini4fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override ASIB read issuing capability
func (r *registerAxiini4fnmodType) SetReadissoverride(value bool) {
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
func (r *registerAxiini4fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini4fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override ASIB write issuing capability
func (r *registerAxiini4fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini4fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini4fnmodFieldWriteissoverrideMask)
	}
}

// registerAxiini5readqosType AXI interconnect - INI x read QoS register
type registerAxiini5readqosType uint32

const (
	RegisterAxiini5readqosFieldArqosShift = 0
	RegisterAxiini5readqosFieldArqosMask  = 0xf
)

// GetArqos Read channel QoS setting
func (r *registerAxiini5readqosType) GetArqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini5readqosFieldArqosMask) >> RegisterAxiini5readqosFieldArqosShift)
}

// SetArqos Read channel QoS setting
func (r *registerAxiini5readqosType) SetArqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini5readqosFieldArqosMask)|(uint32(value)<<RegisterAxiini5readqosFieldArqosShift))
}

// registerAxiini5writeqosType AXI interconnect - INI x write QoS register
type registerAxiini5writeqosType uint32

const (
	RegisterAxiini5writeqosFieldAwqosShift = 0
	RegisterAxiini5writeqosFieldAwqosMask  = 0xf
)

// GetAwqos Write channel QoS setting
func (r *registerAxiini5writeqosType) GetAwqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini5writeqosFieldAwqosMask) >> RegisterAxiini5writeqosFieldAwqosShift)
}

// SetAwqos Write channel QoS setting
func (r *registerAxiini5writeqosType) SetAwqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini5writeqosFieldAwqosMask)|(uint32(value)<<RegisterAxiini5writeqosFieldAwqosShift))
}

// registerAxiini5fnmodType AXI interconnect - INI x issuing functionality modification register
type registerAxiini5fnmodType uint32

const (
	RegisterAxiini5fnmodFieldReadissoverrideShift = 0
	RegisterAxiini5fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override ASIB read issuing capability
func (r *registerAxiini5fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini5fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override ASIB read issuing capability
func (r *registerAxiini5fnmodType) SetReadissoverride(value bool) {
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
func (r *registerAxiini5fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini5fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override ASIB write issuing capability
func (r *registerAxiini5fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini5fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini5fnmodFieldWriteissoverrideMask)
	}
}

// registerAxiini6readqosType AXI interconnect - INI x read QoS register
type registerAxiini6readqosType uint32

const (
	RegisterAxiini6readqosFieldArqosShift = 0
	RegisterAxiini6readqosFieldArqosMask  = 0xf
)

// GetArqos Read channel QoS setting
func (r *registerAxiini6readqosType) GetArqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini6readqosFieldArqosMask) >> RegisterAxiini6readqosFieldArqosShift)
}

// SetArqos Read channel QoS setting
func (r *registerAxiini6readqosType) SetArqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini6readqosFieldArqosMask)|(uint32(value)<<RegisterAxiini6readqosFieldArqosShift))
}

// registerAxiini6writeqosType AXI interconnect - INI x write QoS register
type registerAxiini6writeqosType uint32

const (
	RegisterAxiini6writeqosFieldAwqosShift = 0
	RegisterAxiini6writeqosFieldAwqosMask  = 0xf
)

// GetAwqos Write channel QoS setting
func (r *registerAxiini6writeqosType) GetAwqos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxiini6writeqosFieldAwqosMask) >> RegisterAxiini6writeqosFieldAwqosShift)
}

// SetAwqos Write channel QoS setting
func (r *registerAxiini6writeqosType) SetAwqos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxiini6writeqosFieldAwqosMask)|(uint32(value)<<RegisterAxiini6writeqosFieldAwqosShift))
}

// registerAxiini6fnmodType AXI interconnect - INI x issuing functionality modification register
type registerAxiini6fnmodType uint32

const (
	RegisterAxiini6fnmodFieldReadissoverrideShift = 0
	RegisterAxiini6fnmodFieldReadissoverrideMask  = 0x1
)

// GetReadissoverride Override ASIB read issuing capability
func (r *registerAxiini6fnmodType) GetReadissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini6fnmodFieldReadissoverrideMask) != 0
}

// SetReadissoverride Override ASIB read issuing capability
func (r *registerAxiini6fnmodType) SetReadissoverride(value bool) {
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
func (r *registerAxiini6fnmodType) GetWriteissoverride() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxiini6fnmodFieldWriteissoverrideMask) != 0
}

// SetWriteissoverride Override ASIB write issuing capability
func (r *registerAxiini6fnmodType) SetWriteissoverride(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxiini6fnmodFieldWriteissoverrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxiini6fnmodFieldWriteissoverrideMask)
	}
}
