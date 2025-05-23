package pf

import (
	"unsafe"
	"volatile"
)

var (
	Pf = (*_pf)(unsafe.Pointer(uintptr(0xe000ed78)))
)

type _pf struct {
	Clidr  registerClidrType
	Ctr    registerCtrType
	Ccsidr registerCcsidrType
}

// registerClidrType Cache Level ID register
type registerClidrType uint32

const (
	RegisterClidrFieldCl1Shift = 0
	RegisterClidrFieldCl1Mask  = 0x7
)

// GetCl1 CL1
func (r *registerClidrType) GetCl1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldCl1Mask) >> RegisterClidrFieldCl1Shift)
}

// SetCl1 CL1
func (r *registerClidrType) SetCl1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldCl1Mask)|(uint32(value)<<RegisterClidrFieldCl1Shift))
}

const (
	RegisterClidrFieldCl2Shift = 3
	RegisterClidrFieldCl2Mask  = 0x38
)

// GetCl2 CL2
func (r *registerClidrType) GetCl2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldCl2Mask) >> RegisterClidrFieldCl2Shift)
}

// SetCl2 CL2
func (r *registerClidrType) SetCl2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldCl2Mask)|(uint32(value)<<RegisterClidrFieldCl2Shift))
}

const (
	RegisterClidrFieldCl3Shift = 6
	RegisterClidrFieldCl3Mask  = 0x1c0
)

// GetCl3 CL3
func (r *registerClidrType) GetCl3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldCl3Mask) >> RegisterClidrFieldCl3Shift)
}

// SetCl3 CL3
func (r *registerClidrType) SetCl3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldCl3Mask)|(uint32(value)<<RegisterClidrFieldCl3Shift))
}

const (
	RegisterClidrFieldCl4Shift = 9
	RegisterClidrFieldCl4Mask  = 0xe00
)

// GetCl4 CL4
func (r *registerClidrType) GetCl4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldCl4Mask) >> RegisterClidrFieldCl4Shift)
}

// SetCl4 CL4
func (r *registerClidrType) SetCl4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldCl4Mask)|(uint32(value)<<RegisterClidrFieldCl4Shift))
}

const (
	RegisterClidrFieldCl5Shift = 12
	RegisterClidrFieldCl5Mask  = 0x7000
)

// GetCl5 CL5
func (r *registerClidrType) GetCl5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldCl5Mask) >> RegisterClidrFieldCl5Shift)
}

// SetCl5 CL5
func (r *registerClidrType) SetCl5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldCl5Mask)|(uint32(value)<<RegisterClidrFieldCl5Shift))
}

const (
	RegisterClidrFieldCl6Shift = 15
	RegisterClidrFieldCl6Mask  = 0x38000
)

// GetCl6 CL6
func (r *registerClidrType) GetCl6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldCl6Mask) >> RegisterClidrFieldCl6Shift)
}

// SetCl6 CL6
func (r *registerClidrType) SetCl6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldCl6Mask)|(uint32(value)<<RegisterClidrFieldCl6Shift))
}

const (
	RegisterClidrFieldCl7Shift = 18
	RegisterClidrFieldCl7Mask  = 0x1c0000
)

// GetCl7 CL7
func (r *registerClidrType) GetCl7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldCl7Mask) >> RegisterClidrFieldCl7Shift)
}

// SetCl7 CL7
func (r *registerClidrType) SetCl7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldCl7Mask)|(uint32(value)<<RegisterClidrFieldCl7Shift))
}

const (
	RegisterClidrFieldLouisShift = 21
	RegisterClidrFieldLouisMask  = 0xe00000
)

// GetLouis LoUIS
func (r *registerClidrType) GetLouis() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldLouisMask) >> RegisterClidrFieldLouisShift)
}

// SetLouis LoUIS
func (r *registerClidrType) SetLouis(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldLouisMask)|(uint32(value)<<RegisterClidrFieldLouisShift))
}

const (
	RegisterClidrFieldLocShift = 24
	RegisterClidrFieldLocMask  = 0x7000000
)

// GetLoc LoC
func (r *registerClidrType) GetLoc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldLocMask) >> RegisterClidrFieldLocShift)
}

// SetLoc LoC
func (r *registerClidrType) SetLoc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldLocMask)|(uint32(value)<<RegisterClidrFieldLocShift))
}

const (
	RegisterClidrFieldLouShift = 27
	RegisterClidrFieldLouMask  = 0x38000000
)

// GetLou LoU
func (r *registerClidrType) GetLou() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClidrFieldLouMask) >> RegisterClidrFieldLouShift)
}

// SetLou LoU
func (r *registerClidrType) SetLou(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClidrFieldLouMask)|(uint32(value)<<RegisterClidrFieldLouShift))
}

// registerCtrType Cache Type register
type registerCtrType uint32

const (
	RegisterCtrField_iminlineShift = 0
	RegisterCtrField_iminlineMask  = 0xf
)

// Get_iminline IminLine
func (r *registerCtrType) Get_iminline() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrField_iminlineMask) >> RegisterCtrField_iminlineShift)
}

// Set_iminline IminLine
func (r *registerCtrType) Set_iminline(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrField_iminlineMask)|(uint32(value)<<RegisterCtrField_iminlineShift))
}

const (
	RegisterCtrFieldDminlineShift = 16
	RegisterCtrFieldDminlineMask  = 0xf0000
)

// GetDminline DMinLine
func (r *registerCtrType) GetDminline() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldDminlineMask) >> RegisterCtrFieldDminlineShift)
}

// SetDminline DMinLine
func (r *registerCtrType) SetDminline(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldDminlineMask)|(uint32(value)<<RegisterCtrFieldDminlineShift))
}

const (
	RegisterCtrFieldErgShift = 20
	RegisterCtrFieldErgMask  = 0xf00000
)

// GetErg ERG
func (r *registerCtrType) GetErg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldErgMask) >> RegisterCtrFieldErgShift)
}

// SetErg ERG
func (r *registerCtrType) SetErg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldErgMask)|(uint32(value)<<RegisterCtrFieldErgShift))
}

const (
	RegisterCtrFieldCwgShift = 24
	RegisterCtrFieldCwgMask  = 0xf000000
)

// GetCwg CWG
func (r *registerCtrType) GetCwg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldCwgMask) >> RegisterCtrFieldCwgShift)
}

// SetCwg CWG
func (r *registerCtrType) SetCwg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldCwgMask)|(uint32(value)<<RegisterCtrFieldCwgShift))
}

const (
	RegisterCtrFieldFormatShift = 29
	RegisterCtrFieldFormatMask  = 0xe0000000
)

// GetFormat Format
func (r *registerCtrType) GetFormat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCtrFieldFormatMask) >> RegisterCtrFieldFormatShift)
}

// SetFormat Format
func (r *registerCtrType) SetFormat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCtrFieldFormatMask)|(uint32(value)<<RegisterCtrFieldFormatShift))
}

// registerCcsidrType Cache Size ID register
type registerCcsidrType uint32

const (
	RegisterCcsidrFieldLinesizeShift = 0
	RegisterCcsidrFieldLinesizeMask  = 0x7
)

// GetLinesize LineSize
func (r *registerCcsidrType) GetLinesize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldLinesizeMask) >> RegisterCcsidrFieldLinesizeShift)
}

// SetLinesize LineSize
func (r *registerCcsidrType) SetLinesize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldLinesizeMask)|(uint32(value)<<RegisterCcsidrFieldLinesizeShift))
}

const (
	RegisterCcsidrFieldAssociativityShift = 3
	RegisterCcsidrFieldAssociativityMask  = 0x1ff8
)

// GetAssociativity Associativity
func (r *registerCcsidrType) GetAssociativity() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldAssociativityMask) >> RegisterCcsidrFieldAssociativityShift)
}

// SetAssociativity Associativity
func (r *registerCcsidrType) SetAssociativity(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldAssociativityMask)|(uint32(value)<<RegisterCcsidrFieldAssociativityShift))
}

const (
	RegisterCcsidrFieldNumsetsShift = 13
	RegisterCcsidrFieldNumsetsMask  = 0xfffe000
)

// GetNumsets NumSets
func (r *registerCcsidrType) GetNumsets() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldNumsetsMask) >> RegisterCcsidrFieldNumsetsShift)
}

// SetNumsets NumSets
func (r *registerCcsidrType) SetNumsets(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldNumsetsMask)|(uint32(value)<<RegisterCcsidrFieldNumsetsShift))
}

const (
	RegisterCcsidrFieldWaShift = 28
	RegisterCcsidrFieldWaMask  = 0x10000000
)

// GetWa WA
func (r *registerCcsidrType) GetWa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldWaMask) != 0
}

// SetWa WA
func (r *registerCcsidrType) SetWa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcsidrFieldWaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldWaMask)
	}
}

const (
	RegisterCcsidrFieldRaShift = 29
	RegisterCcsidrFieldRaMask  = 0x20000000
)

// GetRa RA
func (r *registerCcsidrType) GetRa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldRaMask) != 0
}

// SetRa RA
func (r *registerCcsidrType) SetRa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcsidrFieldRaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldRaMask)
	}
}

const (
	RegisterCcsidrFieldWbShift = 30
	RegisterCcsidrFieldWbMask  = 0x40000000
)

// GetWb WB
func (r *registerCcsidrType) GetWb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldWbMask) != 0
}

// SetWb WB
func (r *registerCcsidrType) SetWb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcsidrFieldWbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldWbMask)
	}
}

const (
	RegisterCcsidrFieldWtShift = 31
	RegisterCcsidrFieldWtMask  = 0x80000000
)

// GetWt WT
func (r *registerCcsidrType) GetWt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcsidrFieldWtMask) != 0
}

// SetWt WT
func (r *registerCcsidrType) SetWt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcsidrFieldWtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcsidrFieldWtMask)
	}
}
