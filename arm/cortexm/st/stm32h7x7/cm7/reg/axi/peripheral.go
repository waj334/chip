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
	_                       [8144]byte
	Axi_periph_id_4         registerAxi_periph_id_4Type
	_                       [12]byte
	Axi_periph_id_0         registerAxi_periph_id_0Type
	Axi_periph_id_1         registerAxi_periph_id_1Type
	Axi_periph_id_2         registerAxi_periph_id_2Type
	Axi_periph_id_3         registerAxi_periph_id_3Type
	Axi_comp_id_0           registerAxi_comp_id_0Type
	Axi_comp_id_1           registerAxi_comp_id_1Type
	Axi_comp_id_2           registerAxi_comp_id_2Type
	Axi_comp_id_3           registerAxi_comp_id_3Type
	_                       [8]byte
	Axi_targ1_fn_mod_iss_bm registerAxi_targ1_fn_mod_iss_bmType
	_                       [24]byte
	Axi_targ1_fn_mod2       registerAxi_targ1_fn_mod2Type
	_                       [4]byte
	Axi_targ1_fn_mod_lb     registerAxi_targ1_fn_mod_lbType
	_                       [216]byte
	Axi_targ1_fn_mod        registerAxi_targ1_fn_modType
	_                       [3836]byte
	Axi_targ2_fn_mod_iss_bm registerAxi_targ2_fn_mod_iss_bmType
	_                       [24]byte
	Axi_targ2_fn_mod2       registerAxi_targ2_fn_mod2Type
	_                       [4]byte
	Axi_targ2_fn_mod_lb     registerAxi_targ2_fn_mod_lbType
	_                       [216]byte
	Axi_targ2_fn_mod        registerAxi_targ2_fn_modType
	_                       [3836]byte
	Axi_targ3_fn_mod_iss_bm registerAxi_targ3_fn_mod_iss_bmType
	_                       [4092]byte
	Axi_targ4_fn_mod_iss_bm registerAxi_targ4_fn_mod_iss_bmType
	_                       [4092]byte
	Axi_targ5_fn_mod_iss_bm registerAxi_targ5_fn_mod_iss_bmType
	_                       [4092]byte
	Axi_targ6_fn_mod_iss_bm registerAxi_targ6_fn_mod_iss_bmType
	_                       [4096]byte
	Axi_targ7_fn_mod_iss_bm registerAxi_targ7_fn_mod_iss_bmType
	_                       [20]byte
	Axi_targ7_fn_mod2       registerAxi_targ7_fn_mod2Type
	_                       [224]byte
	Axi_targ7_fn_mod        registerAxi_targ7_fn_modType
	_                       [237336]byte
	Axi_ini1_fn_mod2        registerAxi_ini1_fn_mod2Type
	Axi_ini1_fn_mod_ahb     registerAxi_ini1_fn_mod_ahbType
	_                       [212]byte
	Axi_ini1_read_qos       registerAxi_ini1_read_qosType
	Axi_ini1_write_qos      registerAxi_ini1_write_qosType
	Axi_ini1_fn_mod         registerAxi_ini1_fn_modType
	_                       [4084]byte
	Axi_ini2_read_qos       registerAxi_ini2_read_qosType
	Axi_ini2_write_qos      registerAxi_ini2_write_qosType
	Axi_ini2_fn_mod         registerAxi_ini2_fn_modType
	_                       [3864]byte
	Axi_ini3_fn_mod2        registerAxi_ini3_fn_mod2Type
	Axi_ini3_fn_mod_ahb     registerAxi_ini3_fn_mod_ahbType
	_                       [212]byte
	Axi_ini3_read_qos       registerAxi_ini3_read_qosType
	Axi_ini3_write_qos      registerAxi_ini3_write_qosType
	Axi_ini3_fn_mod         registerAxi_ini3_fn_modType
	_                       [4084]byte
	Axi_ini4_read_qos       registerAxi_ini4_read_qosType
	Axi_ini4_write_qos      registerAxi_ini4_write_qosType
	Axi_ini4_fn_mod         registerAxi_ini4_fn_modType
	_                       [4084]byte
	Axi_ini5_read_qos       registerAxi_ini5_read_qosType
	Axi_ini5_write_qos      registerAxi_ini5_write_qosType
	Axi_ini5_fn_mod         registerAxi_ini5_fn_modType
	_                       [4084]byte
	Axi_ini6_read_qos       registerAxi_ini6_read_qosType
	Axi_ini6_write_qos      registerAxi_ini6_write_qosType
	Axi_ini6_fn_mod         registerAxi_ini6_fn_modType
}

// registerAxi_periph_id_4Type AXI interconnect - peripheral ID4 register
type registerAxi_periph_id_4Type uint32

const (
	RegisterAxi_periph_id_4FieldJep106conShift = 0
	RegisterAxi_periph_id_4FieldJep106conMask  = 0xf
)

// GetJep106con JEP106 continuation code
func (r *registerAxi_periph_id_4Type) GetJep106con() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_periph_id_4FieldJep106conMask) >> RegisterAxi_periph_id_4FieldJep106conShift)
}

// SetJep106con JEP106 continuation code
func (r *registerAxi_periph_id_4Type) SetJep106con(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_periph_id_4FieldJep106conMask)|(uint32(value)<<RegisterAxi_periph_id_4FieldJep106conShift))
}

const (
	RegisterAxi_periph_id_4FieldKcount4Shift = 4
	RegisterAxi_periph_id_4FieldKcount4Mask  = 0xf0
)

// GetKcount4 Register file size
func (r *registerAxi_periph_id_4Type) GetKcount4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_periph_id_4FieldKcount4Mask) >> RegisterAxi_periph_id_4FieldKcount4Shift)
}

// SetKcount4 Register file size
func (r *registerAxi_periph_id_4Type) SetKcount4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_periph_id_4FieldKcount4Mask)|(uint32(value)<<RegisterAxi_periph_id_4FieldKcount4Shift))
}

// registerAxi_periph_id_0Type AXI interconnect - peripheral ID0 register
type registerAxi_periph_id_0Type uint32

const (
	RegisterAxi_periph_id_0FieldPartnumShift = 0
	RegisterAxi_periph_id_0FieldPartnumMask  = 0xff
)

// GetPartnum Peripheral part number bits 0 to 7
func (r *registerAxi_periph_id_0Type) GetPartnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_periph_id_0FieldPartnumMask) >> RegisterAxi_periph_id_0FieldPartnumShift)
}

// SetPartnum Peripheral part number bits 0 to 7
func (r *registerAxi_periph_id_0Type) SetPartnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_periph_id_0FieldPartnumMask)|(uint32(value)<<RegisterAxi_periph_id_0FieldPartnumShift))
}

// registerAxi_periph_id_1Type AXI interconnect - peripheral ID1 register
type registerAxi_periph_id_1Type uint32

const (
	RegisterAxi_periph_id_1FieldPartnumShift = 0
	RegisterAxi_periph_id_1FieldPartnumMask  = 0xf
)

// GetPartnum Peripheral part number bits 8 to 11
func (r *registerAxi_periph_id_1Type) GetPartnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_periph_id_1FieldPartnumMask) >> RegisterAxi_periph_id_1FieldPartnumShift)
}

// SetPartnum Peripheral part number bits 8 to 11
func (r *registerAxi_periph_id_1Type) SetPartnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_periph_id_1FieldPartnumMask)|(uint32(value)<<RegisterAxi_periph_id_1FieldPartnumShift))
}

const (
	RegisterAxi_periph_id_1FieldJep106iShift = 4
	RegisterAxi_periph_id_1FieldJep106iMask  = 0xf0
)

// GetJep106i JEP106 identity bits 0 to 3
func (r *registerAxi_periph_id_1Type) GetJep106i() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_periph_id_1FieldJep106iMask) >> RegisterAxi_periph_id_1FieldJep106iShift)
}

// SetJep106i JEP106 identity bits 0 to 3
func (r *registerAxi_periph_id_1Type) SetJep106i(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_periph_id_1FieldJep106iMask)|(uint32(value)<<RegisterAxi_periph_id_1FieldJep106iShift))
}

// registerAxi_periph_id_2Type AXI interconnect - peripheral ID2 register
type registerAxi_periph_id_2Type uint32

const (
	RegisterAxi_periph_id_2FieldJep106idShift = 0
	RegisterAxi_periph_id_2FieldJep106idMask  = 0x7
)

// GetJep106id JEP106 Identity bits 4 to 6
func (r *registerAxi_periph_id_2Type) GetJep106id() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_periph_id_2FieldJep106idMask) >> RegisterAxi_periph_id_2FieldJep106idShift)
}

// SetJep106id JEP106 Identity bits 4 to 6
func (r *registerAxi_periph_id_2Type) SetJep106id(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_periph_id_2FieldJep106idMask)|(uint32(value)<<RegisterAxi_periph_id_2FieldJep106idShift))
}

const (
	RegisterAxi_periph_id_2FieldJedecShift = 3
	RegisterAxi_periph_id_2FieldJedecMask  = 0x8
)

// GetJedec JEP106 code flag
func (r *registerAxi_periph_id_2Type) GetJedec() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_periph_id_2FieldJedecMask) != 0
}

// SetJedec JEP106 code flag
func (r *registerAxi_periph_id_2Type) SetJedec(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_periph_id_2FieldJedecMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_periph_id_2FieldJedecMask)
	}
}

const (
	RegisterAxi_periph_id_2FieldRevisionShift = 4
	RegisterAxi_periph_id_2FieldRevisionMask  = 0xf0
)

// GetRevision Peripheral revision number
func (r *registerAxi_periph_id_2Type) GetRevision() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_periph_id_2FieldRevisionMask) >> RegisterAxi_periph_id_2FieldRevisionShift)
}

// SetRevision Peripheral revision number
func (r *registerAxi_periph_id_2Type) SetRevision(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_periph_id_2FieldRevisionMask)|(uint32(value)<<RegisterAxi_periph_id_2FieldRevisionShift))
}

// registerAxi_periph_id_3Type AXI interconnect - peripheral ID3 register
type registerAxi_periph_id_3Type uint32

const (
	RegisterAxi_periph_id_3FieldCust_mod_numShift = 0
	RegisterAxi_periph_id_3FieldCust_mod_numMask  = 0xf
)

// GetCust_mod_num Customer modification
func (r *registerAxi_periph_id_3Type) GetCust_mod_num() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_periph_id_3FieldCust_mod_numMask) >> RegisterAxi_periph_id_3FieldCust_mod_numShift)
}

// SetCust_mod_num Customer modification
func (r *registerAxi_periph_id_3Type) SetCust_mod_num(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_periph_id_3FieldCust_mod_numMask)|(uint32(value)<<RegisterAxi_periph_id_3FieldCust_mod_numShift))
}

const (
	RegisterAxi_periph_id_3FieldRev_andShift = 4
	RegisterAxi_periph_id_3FieldRev_andMask  = 0xf0
)

// GetRev_and Customer version
func (r *registerAxi_periph_id_3Type) GetRev_and() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_periph_id_3FieldRev_andMask) >> RegisterAxi_periph_id_3FieldRev_andShift)
}

// SetRev_and Customer version
func (r *registerAxi_periph_id_3Type) SetRev_and(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_periph_id_3FieldRev_andMask)|(uint32(value)<<RegisterAxi_periph_id_3FieldRev_andShift))
}

// registerAxi_comp_id_0Type AXI interconnect - component ID0 register
type registerAxi_comp_id_0Type uint32

const (
	RegisterAxi_comp_id_0FieldPreambleShift = 0
	RegisterAxi_comp_id_0FieldPreambleMask  = 0xff
)

// GetPreamble Preamble bits 0 to 7
func (r *registerAxi_comp_id_0Type) GetPreamble() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_comp_id_0FieldPreambleMask) >> RegisterAxi_comp_id_0FieldPreambleShift)
}

// SetPreamble Preamble bits 0 to 7
func (r *registerAxi_comp_id_0Type) SetPreamble(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_comp_id_0FieldPreambleMask)|(uint32(value)<<RegisterAxi_comp_id_0FieldPreambleShift))
}

// registerAxi_comp_id_1Type AXI interconnect - component ID1 register
type registerAxi_comp_id_1Type uint32

const (
	RegisterAxi_comp_id_1FieldPreambleShift = 0
	RegisterAxi_comp_id_1FieldPreambleMask  = 0xf
)

// GetPreamble Preamble bits 8 to 11
func (r *registerAxi_comp_id_1Type) GetPreamble() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_comp_id_1FieldPreambleMask) >> RegisterAxi_comp_id_1FieldPreambleShift)
}

// SetPreamble Preamble bits 8 to 11
func (r *registerAxi_comp_id_1Type) SetPreamble(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_comp_id_1FieldPreambleMask)|(uint32(value)<<RegisterAxi_comp_id_1FieldPreambleShift))
}

const (
	RegisterAxi_comp_id_1FieldClassShift = 4
	RegisterAxi_comp_id_1FieldClassMask  = 0xf0
)

// GetClass Component class
func (r *registerAxi_comp_id_1Type) GetClass() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_comp_id_1FieldClassMask) >> RegisterAxi_comp_id_1FieldClassShift)
}

// SetClass Component class
func (r *registerAxi_comp_id_1Type) SetClass(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_comp_id_1FieldClassMask)|(uint32(value)<<RegisterAxi_comp_id_1FieldClassShift))
}

// registerAxi_comp_id_2Type AXI interconnect - component ID2 register
type registerAxi_comp_id_2Type uint32

const (
	RegisterAxi_comp_id_2FieldPreambleShift = 0
	RegisterAxi_comp_id_2FieldPreambleMask  = 0xff
)

// GetPreamble Preamble bits 12 to 19
func (r *registerAxi_comp_id_2Type) GetPreamble() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_comp_id_2FieldPreambleMask) >> RegisterAxi_comp_id_2FieldPreambleShift)
}

// SetPreamble Preamble bits 12 to 19
func (r *registerAxi_comp_id_2Type) SetPreamble(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_comp_id_2FieldPreambleMask)|(uint32(value)<<RegisterAxi_comp_id_2FieldPreambleShift))
}

// registerAxi_comp_id_3Type AXI interconnect - component ID3 register
type registerAxi_comp_id_3Type uint32

const (
	RegisterAxi_comp_id_3FieldPreambleShift = 0
	RegisterAxi_comp_id_3FieldPreambleMask  = 0xff
)

// GetPreamble Preamble bits 20 to 27
func (r *registerAxi_comp_id_3Type) GetPreamble() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_comp_id_3FieldPreambleMask) >> RegisterAxi_comp_id_3FieldPreambleShift)
}

// SetPreamble Preamble bits 20 to 27
func (r *registerAxi_comp_id_3Type) SetPreamble(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_comp_id_3FieldPreambleMask)|(uint32(value)<<RegisterAxi_comp_id_3FieldPreambleShift))
}

// registerAxi_targ1_fn_mod_iss_bmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxi_targ1_fn_mod_iss_bmType uint32

const (
	RegisterAxi_targ1_fn_mod_iss_bmFieldRead_iss_overrideShift = 0
	RegisterAxi_targ1_fn_mod_iss_bmFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ1_fn_mod_iss_bmType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ1_fn_mod_iss_bmFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ1_fn_mod_iss_bmType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ1_fn_mod_iss_bmFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ1_fn_mod_iss_bmFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_targ1_fn_mod_iss_bmFieldWrite_iss_overrideShift = 1
	RegisterAxi_targ1_fn_mod_iss_bmFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ1_fn_mod_iss_bmType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ1_fn_mod_iss_bmFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ1_fn_mod_iss_bmType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ1_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ1_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	}
}

// registerAxi_targ1_fn_mod2Type AXI interconnect - TARG x bus matrix functionality 2 register
type registerAxi_targ1_fn_mod2Type uint32

const (
	RegisterAxi_targ1_fn_mod2FieldBypass_mergeShift = 0
	RegisterAxi_targ1_fn_mod2FieldBypass_mergeMask  = 0x1
)

// GetBypass_merge Disable packing of beats to match the output data width
func (r *registerAxi_targ1_fn_mod2Type) GetBypass_merge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ1_fn_mod2FieldBypass_mergeMask) != 0
}

// SetBypass_merge Disable packing of beats to match the output data width
func (r *registerAxi_targ1_fn_mod2Type) SetBypass_merge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ1_fn_mod2FieldBypass_mergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ1_fn_mod2FieldBypass_mergeMask)
	}
}

// registerAxi_targ1_fn_mod_lbType AXI interconnect - TARG x long burst functionality modification
type registerAxi_targ1_fn_mod_lbType uint32

const (
	RegisterAxi_targ1_fn_mod_lbFieldFn_mod_lbShift = 0
	RegisterAxi_targ1_fn_mod_lbFieldFn_mod_lbMask  = 0x1
)

// GetFn_mod_lb Controls burst breaking of long bursts
func (r *registerAxi_targ1_fn_mod_lbType) GetFn_mod_lb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ1_fn_mod_lbFieldFn_mod_lbMask) != 0
}

// SetFn_mod_lb Controls burst breaking of long bursts
func (r *registerAxi_targ1_fn_mod_lbType) SetFn_mod_lb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ1_fn_mod_lbFieldFn_mod_lbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ1_fn_mod_lbFieldFn_mod_lbMask)
	}
}

// registerAxi_targ1_fn_modType AXI interconnect - TARG x long burst functionality modification
type registerAxi_targ1_fn_modType uint32

const (
	RegisterAxi_targ1_fn_modFieldRead_iss_overrideShift = 0
	RegisterAxi_targ1_fn_modFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override Override AMIB read issuing capability
func (r *registerAxi_targ1_fn_modType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ1_fn_modFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override Override AMIB read issuing capability
func (r *registerAxi_targ1_fn_modType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ1_fn_modFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ1_fn_modFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_targ1_fn_modFieldWrite_iss_overrideShift = 1
	RegisterAxi_targ1_fn_modFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Override AMIB write issuing capability
func (r *registerAxi_targ1_fn_modType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ1_fn_modFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Override AMIB write issuing capability
func (r *registerAxi_targ1_fn_modType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ1_fn_modFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ1_fn_modFieldWrite_iss_overrideMask)
	}
}

// registerAxi_targ2_fn_mod_iss_bmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxi_targ2_fn_mod_iss_bmType uint32

const (
	RegisterAxi_targ2_fn_mod_iss_bmFieldRead_iss_overrideShift = 0
	RegisterAxi_targ2_fn_mod_iss_bmFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ2_fn_mod_iss_bmType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ2_fn_mod_iss_bmFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ2_fn_mod_iss_bmType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ2_fn_mod_iss_bmFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ2_fn_mod_iss_bmFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_targ2_fn_mod_iss_bmFieldWrite_iss_overrideShift = 1
	RegisterAxi_targ2_fn_mod_iss_bmFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ2_fn_mod_iss_bmType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ2_fn_mod_iss_bmFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ2_fn_mod_iss_bmType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ2_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ2_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	}
}

// registerAxi_targ2_fn_mod2Type AXI interconnect - TARG x bus matrix functionality 2 register
type registerAxi_targ2_fn_mod2Type uint32

const (
	RegisterAxi_targ2_fn_mod2FieldBypass_mergeShift = 0
	RegisterAxi_targ2_fn_mod2FieldBypass_mergeMask  = 0x1
)

// GetBypass_merge Disable packing of beats to match the output data width
func (r *registerAxi_targ2_fn_mod2Type) GetBypass_merge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ2_fn_mod2FieldBypass_mergeMask) != 0
}

// SetBypass_merge Disable packing of beats to match the output data width
func (r *registerAxi_targ2_fn_mod2Type) SetBypass_merge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ2_fn_mod2FieldBypass_mergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ2_fn_mod2FieldBypass_mergeMask)
	}
}

// registerAxi_targ2_fn_mod_lbType AXI interconnect - TARG x long burst functionality modification
type registerAxi_targ2_fn_mod_lbType uint32

const (
	RegisterAxi_targ2_fn_mod_lbFieldFn_mod_lbShift = 0
	RegisterAxi_targ2_fn_mod_lbFieldFn_mod_lbMask  = 0x1
)

// GetFn_mod_lb Controls burst breaking of long bursts
func (r *registerAxi_targ2_fn_mod_lbType) GetFn_mod_lb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ2_fn_mod_lbFieldFn_mod_lbMask) != 0
}

// SetFn_mod_lb Controls burst breaking of long bursts
func (r *registerAxi_targ2_fn_mod_lbType) SetFn_mod_lb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ2_fn_mod_lbFieldFn_mod_lbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ2_fn_mod_lbFieldFn_mod_lbMask)
	}
}

// registerAxi_targ2_fn_modType AXI interconnect - TARG x long burst functionality modification
type registerAxi_targ2_fn_modType uint32

const (
	RegisterAxi_targ2_fn_modFieldRead_iss_overrideShift = 0
	RegisterAxi_targ2_fn_modFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override Override AMIB read issuing capability
func (r *registerAxi_targ2_fn_modType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ2_fn_modFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override Override AMIB read issuing capability
func (r *registerAxi_targ2_fn_modType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ2_fn_modFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ2_fn_modFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_targ2_fn_modFieldWrite_iss_overrideShift = 1
	RegisterAxi_targ2_fn_modFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Override AMIB write issuing capability
func (r *registerAxi_targ2_fn_modType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ2_fn_modFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Override AMIB write issuing capability
func (r *registerAxi_targ2_fn_modType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ2_fn_modFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ2_fn_modFieldWrite_iss_overrideMask)
	}
}

// registerAxi_targ3_fn_mod_iss_bmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxi_targ3_fn_mod_iss_bmType uint32

const (
	RegisterAxi_targ3_fn_mod_iss_bmFieldRead_iss_overrideShift = 0
	RegisterAxi_targ3_fn_mod_iss_bmFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ3_fn_mod_iss_bmType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ3_fn_mod_iss_bmFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ3_fn_mod_iss_bmType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ3_fn_mod_iss_bmFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ3_fn_mod_iss_bmFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_targ3_fn_mod_iss_bmFieldWrite_iss_overrideShift = 1
	RegisterAxi_targ3_fn_mod_iss_bmFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ3_fn_mod_iss_bmType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ3_fn_mod_iss_bmFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ3_fn_mod_iss_bmType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ3_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ3_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	}
}

// registerAxi_targ4_fn_mod_iss_bmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxi_targ4_fn_mod_iss_bmType uint32

const (
	RegisterAxi_targ4_fn_mod_iss_bmFieldRead_iss_overrideShift = 0
	RegisterAxi_targ4_fn_mod_iss_bmFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ4_fn_mod_iss_bmType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ4_fn_mod_iss_bmFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ4_fn_mod_iss_bmType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ4_fn_mod_iss_bmFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ4_fn_mod_iss_bmFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_targ4_fn_mod_iss_bmFieldWrite_iss_overrideShift = 1
	RegisterAxi_targ4_fn_mod_iss_bmFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ4_fn_mod_iss_bmType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ4_fn_mod_iss_bmFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ4_fn_mod_iss_bmType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ4_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ4_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	}
}

// registerAxi_targ5_fn_mod_iss_bmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxi_targ5_fn_mod_iss_bmType uint32

const (
	RegisterAxi_targ5_fn_mod_iss_bmFieldRead_iss_overrideShift = 0
	RegisterAxi_targ5_fn_mod_iss_bmFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ5_fn_mod_iss_bmType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ5_fn_mod_iss_bmFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ5_fn_mod_iss_bmType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ5_fn_mod_iss_bmFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ5_fn_mod_iss_bmFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_targ5_fn_mod_iss_bmFieldWrite_iss_overrideShift = 1
	RegisterAxi_targ5_fn_mod_iss_bmFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ5_fn_mod_iss_bmType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ5_fn_mod_iss_bmFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ5_fn_mod_iss_bmType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ5_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ5_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	}
}

// registerAxi_targ6_fn_mod_iss_bmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxi_targ6_fn_mod_iss_bmType uint32

const (
	RegisterAxi_targ6_fn_mod_iss_bmFieldRead_iss_overrideShift = 0
	RegisterAxi_targ6_fn_mod_iss_bmFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ6_fn_mod_iss_bmType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ6_fn_mod_iss_bmFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ6_fn_mod_iss_bmType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ6_fn_mod_iss_bmFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ6_fn_mod_iss_bmFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_targ6_fn_mod_iss_bmFieldWrite_iss_overrideShift = 1
	RegisterAxi_targ6_fn_mod_iss_bmFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ6_fn_mod_iss_bmType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ6_fn_mod_iss_bmFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ6_fn_mod_iss_bmType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ6_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ6_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	}
}

// registerAxi_targ7_fn_mod_iss_bmType AXI interconnect - TARG x bus matrix issuing functionality register
type registerAxi_targ7_fn_mod_iss_bmType uint32

const (
	RegisterAxi_targ7_fn_mod_iss_bmFieldRead_iss_overrideShift = 0
	RegisterAxi_targ7_fn_mod_iss_bmFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ7_fn_mod_iss_bmType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ7_fn_mod_iss_bmFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override READ_ISS_OVERRIDE
func (r *registerAxi_targ7_fn_mod_iss_bmType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ7_fn_mod_iss_bmFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ7_fn_mod_iss_bmFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_targ7_fn_mod_iss_bmFieldWrite_iss_overrideShift = 1
	RegisterAxi_targ7_fn_mod_iss_bmFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ7_fn_mod_iss_bmType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ7_fn_mod_iss_bmFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Switch matrix write issuing override for target
func (r *registerAxi_targ7_fn_mod_iss_bmType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ7_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ7_fn_mod_iss_bmFieldWrite_iss_overrideMask)
	}
}

// registerAxi_targ7_fn_mod2Type AXI interconnect - TARG x bus matrix functionality 2 register
type registerAxi_targ7_fn_mod2Type uint32

const (
	RegisterAxi_targ7_fn_mod2FieldBypass_mergeShift = 0
	RegisterAxi_targ7_fn_mod2FieldBypass_mergeMask  = 0x1
)

// GetBypass_merge Disable packing of beats to match the output data width
func (r *registerAxi_targ7_fn_mod2Type) GetBypass_merge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ7_fn_mod2FieldBypass_mergeMask) != 0
}

// SetBypass_merge Disable packing of beats to match the output data width
func (r *registerAxi_targ7_fn_mod2Type) SetBypass_merge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ7_fn_mod2FieldBypass_mergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ7_fn_mod2FieldBypass_mergeMask)
	}
}

// registerAxi_targ7_fn_modType AXI interconnect - TARG x long burst functionality modification
type registerAxi_targ7_fn_modType uint32

const (
	RegisterAxi_targ7_fn_modFieldRead_iss_overrideShift = 0
	RegisterAxi_targ7_fn_modFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override Override AMIB read issuing capability
func (r *registerAxi_targ7_fn_modType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ7_fn_modFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override Override AMIB read issuing capability
func (r *registerAxi_targ7_fn_modType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ7_fn_modFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ7_fn_modFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_targ7_fn_modFieldWrite_iss_overrideShift = 1
	RegisterAxi_targ7_fn_modFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Override AMIB write issuing capability
func (r *registerAxi_targ7_fn_modType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_targ7_fn_modFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Override AMIB write issuing capability
func (r *registerAxi_targ7_fn_modType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_targ7_fn_modFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_targ7_fn_modFieldWrite_iss_overrideMask)
	}
}

// registerAxi_ini1_fn_mod2Type AXI interconnect - INI x functionality modification 2 register
type registerAxi_ini1_fn_mod2Type uint32

const (
	RegisterAxi_ini1_fn_mod2FieldBypass_mergeShift = 0
	RegisterAxi_ini1_fn_mod2FieldBypass_mergeMask  = 0x1
)

// GetBypass_merge Disables alteration of transactions by the up-sizer unless required by the protocol
func (r *registerAxi_ini1_fn_mod2Type) GetBypass_merge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini1_fn_mod2FieldBypass_mergeMask) != 0
}

// SetBypass_merge Disables alteration of transactions by the up-sizer unless required by the protocol
func (r *registerAxi_ini1_fn_mod2Type) SetBypass_merge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini1_fn_mod2FieldBypass_mergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini1_fn_mod2FieldBypass_mergeMask)
	}
}

// registerAxi_ini1_fn_mod_ahbType AXI interconnect - INI x AHB functionality modification register
type registerAxi_ini1_fn_mod_ahbType uint32

const (
	RegisterAxi_ini1_fn_mod_ahbFieldRd_inc_overrideShift = 0
	RegisterAxi_ini1_fn_mod_ahbFieldRd_inc_overrideMask  = 0x1
)

// GetRd_inc_override Converts all AHB-Lite write transactions to a series of single beat AXI
func (r *registerAxi_ini1_fn_mod_ahbType) GetRd_inc_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini1_fn_mod_ahbFieldRd_inc_overrideMask) != 0
}

// SetRd_inc_override Converts all AHB-Lite write transactions to a series of single beat AXI
func (r *registerAxi_ini1_fn_mod_ahbType) SetRd_inc_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini1_fn_mod_ahbFieldRd_inc_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini1_fn_mod_ahbFieldRd_inc_overrideMask)
	}
}

const (
	RegisterAxi_ini1_fn_mod_ahbFieldWr_inc_overrideShift = 1
	RegisterAxi_ini1_fn_mod_ahbFieldWr_inc_overrideMask  = 0x2
)

// GetWr_inc_override Converts all AHB-Lite read transactions to a series of single beat AXI
func (r *registerAxi_ini1_fn_mod_ahbType) GetWr_inc_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini1_fn_mod_ahbFieldWr_inc_overrideMask) != 0
}

// SetWr_inc_override Converts all AHB-Lite read transactions to a series of single beat AXI
func (r *registerAxi_ini1_fn_mod_ahbType) SetWr_inc_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini1_fn_mod_ahbFieldWr_inc_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini1_fn_mod_ahbFieldWr_inc_overrideMask)
	}
}

// registerAxi_ini1_read_qosType AXI interconnect - INI x read QoS register
type registerAxi_ini1_read_qosType uint32

const (
	RegisterAxi_ini1_read_qosFieldAr_qosShift = 0
	RegisterAxi_ini1_read_qosFieldAr_qosMask  = 0xf
)

// GetAr_qos Read channel QoS setting
func (r *registerAxi_ini1_read_qosType) GetAr_qos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini1_read_qosFieldAr_qosMask) >> RegisterAxi_ini1_read_qosFieldAr_qosShift)
}

// SetAr_qos Read channel QoS setting
func (r *registerAxi_ini1_read_qosType) SetAr_qos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini1_read_qosFieldAr_qosMask)|(uint32(value)<<RegisterAxi_ini1_read_qosFieldAr_qosShift))
}

// registerAxi_ini1_write_qosType AXI interconnect - INI x write QoS register
type registerAxi_ini1_write_qosType uint32

const (
	RegisterAxi_ini1_write_qosFieldAw_qosShift = 0
	RegisterAxi_ini1_write_qosFieldAw_qosMask  = 0xf
)

// GetAw_qos Write channel QoS setting
func (r *registerAxi_ini1_write_qosType) GetAw_qos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini1_write_qosFieldAw_qosMask) >> RegisterAxi_ini1_write_qosFieldAw_qosShift)
}

// SetAw_qos Write channel QoS setting
func (r *registerAxi_ini1_write_qosType) SetAw_qos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini1_write_qosFieldAw_qosMask)|(uint32(value)<<RegisterAxi_ini1_write_qosFieldAw_qosShift))
}

// registerAxi_ini1_fn_modType AXI interconnect - INI x issuing functionality modification register
type registerAxi_ini1_fn_modType uint32

const (
	RegisterAxi_ini1_fn_modFieldRead_iss_overrideShift = 0
	RegisterAxi_ini1_fn_modFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override Override ASIB read issuing capability
func (r *registerAxi_ini1_fn_modType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini1_fn_modFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override Override ASIB read issuing capability
func (r *registerAxi_ini1_fn_modType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini1_fn_modFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini1_fn_modFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_ini1_fn_modFieldWrite_iss_overrideShift = 1
	RegisterAxi_ini1_fn_modFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Override ASIB write issuing capability
func (r *registerAxi_ini1_fn_modType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini1_fn_modFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Override ASIB write issuing capability
func (r *registerAxi_ini1_fn_modType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini1_fn_modFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini1_fn_modFieldWrite_iss_overrideMask)
	}
}

// registerAxi_ini2_read_qosType AXI interconnect - INI x read QoS register
type registerAxi_ini2_read_qosType uint32

const (
	RegisterAxi_ini2_read_qosFieldAr_qosShift = 0
	RegisterAxi_ini2_read_qosFieldAr_qosMask  = 0xf
)

// GetAr_qos Read channel QoS setting
func (r *registerAxi_ini2_read_qosType) GetAr_qos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini2_read_qosFieldAr_qosMask) >> RegisterAxi_ini2_read_qosFieldAr_qosShift)
}

// SetAr_qos Read channel QoS setting
func (r *registerAxi_ini2_read_qosType) SetAr_qos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini2_read_qosFieldAr_qosMask)|(uint32(value)<<RegisterAxi_ini2_read_qosFieldAr_qosShift))
}

// registerAxi_ini2_write_qosType AXI interconnect - INI x write QoS register
type registerAxi_ini2_write_qosType uint32

const (
	RegisterAxi_ini2_write_qosFieldAw_qosShift = 0
	RegisterAxi_ini2_write_qosFieldAw_qosMask  = 0xf
)

// GetAw_qos Write channel QoS setting
func (r *registerAxi_ini2_write_qosType) GetAw_qos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini2_write_qosFieldAw_qosMask) >> RegisterAxi_ini2_write_qosFieldAw_qosShift)
}

// SetAw_qos Write channel QoS setting
func (r *registerAxi_ini2_write_qosType) SetAw_qos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini2_write_qosFieldAw_qosMask)|(uint32(value)<<RegisterAxi_ini2_write_qosFieldAw_qosShift))
}

// registerAxi_ini2_fn_modType AXI interconnect - INI x issuing functionality modification register
type registerAxi_ini2_fn_modType uint32

const (
	RegisterAxi_ini2_fn_modFieldRead_iss_overrideShift = 0
	RegisterAxi_ini2_fn_modFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override Override ASIB read issuing capability
func (r *registerAxi_ini2_fn_modType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini2_fn_modFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override Override ASIB read issuing capability
func (r *registerAxi_ini2_fn_modType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini2_fn_modFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini2_fn_modFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_ini2_fn_modFieldWrite_iss_overrideShift = 1
	RegisterAxi_ini2_fn_modFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Override ASIB write issuing capability
func (r *registerAxi_ini2_fn_modType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini2_fn_modFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Override ASIB write issuing capability
func (r *registerAxi_ini2_fn_modType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini2_fn_modFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini2_fn_modFieldWrite_iss_overrideMask)
	}
}

// registerAxi_ini3_fn_mod2Type AXI interconnect - INI x functionality modification 2 register
type registerAxi_ini3_fn_mod2Type uint32

const (
	RegisterAxi_ini3_fn_mod2FieldBypass_mergeShift = 0
	RegisterAxi_ini3_fn_mod2FieldBypass_mergeMask  = 0x1
)

// GetBypass_merge Disables alteration of transactions by the up-sizer unless required by the protocol
func (r *registerAxi_ini3_fn_mod2Type) GetBypass_merge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini3_fn_mod2FieldBypass_mergeMask) != 0
}

// SetBypass_merge Disables alteration of transactions by the up-sizer unless required by the protocol
func (r *registerAxi_ini3_fn_mod2Type) SetBypass_merge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini3_fn_mod2FieldBypass_mergeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini3_fn_mod2FieldBypass_mergeMask)
	}
}

// registerAxi_ini3_fn_mod_ahbType AXI interconnect - INI x AHB functionality modification register
type registerAxi_ini3_fn_mod_ahbType uint32

const (
	RegisterAxi_ini3_fn_mod_ahbFieldRd_inc_overrideShift = 0
	RegisterAxi_ini3_fn_mod_ahbFieldRd_inc_overrideMask  = 0x1
)

// GetRd_inc_override Converts all AHB-Lite write transactions to a series of single beat AXI
func (r *registerAxi_ini3_fn_mod_ahbType) GetRd_inc_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini3_fn_mod_ahbFieldRd_inc_overrideMask) != 0
}

// SetRd_inc_override Converts all AHB-Lite write transactions to a series of single beat AXI
func (r *registerAxi_ini3_fn_mod_ahbType) SetRd_inc_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini3_fn_mod_ahbFieldRd_inc_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini3_fn_mod_ahbFieldRd_inc_overrideMask)
	}
}

const (
	RegisterAxi_ini3_fn_mod_ahbFieldWr_inc_overrideShift = 1
	RegisterAxi_ini3_fn_mod_ahbFieldWr_inc_overrideMask  = 0x2
)

// GetWr_inc_override Converts all AHB-Lite read transactions to a series of single beat AXI
func (r *registerAxi_ini3_fn_mod_ahbType) GetWr_inc_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini3_fn_mod_ahbFieldWr_inc_overrideMask) != 0
}

// SetWr_inc_override Converts all AHB-Lite read transactions to a series of single beat AXI
func (r *registerAxi_ini3_fn_mod_ahbType) SetWr_inc_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini3_fn_mod_ahbFieldWr_inc_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini3_fn_mod_ahbFieldWr_inc_overrideMask)
	}
}

// registerAxi_ini3_read_qosType AXI interconnect - INI x read QoS register
type registerAxi_ini3_read_qosType uint32

const (
	RegisterAxi_ini3_read_qosFieldAr_qosShift = 0
	RegisterAxi_ini3_read_qosFieldAr_qosMask  = 0xf
)

// GetAr_qos Read channel QoS setting
func (r *registerAxi_ini3_read_qosType) GetAr_qos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini3_read_qosFieldAr_qosMask) >> RegisterAxi_ini3_read_qosFieldAr_qosShift)
}

// SetAr_qos Read channel QoS setting
func (r *registerAxi_ini3_read_qosType) SetAr_qos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini3_read_qosFieldAr_qosMask)|(uint32(value)<<RegisterAxi_ini3_read_qosFieldAr_qosShift))
}

// registerAxi_ini3_write_qosType AXI interconnect - INI x write QoS register
type registerAxi_ini3_write_qosType uint32

const (
	RegisterAxi_ini3_write_qosFieldAw_qosShift = 0
	RegisterAxi_ini3_write_qosFieldAw_qosMask  = 0xf
)

// GetAw_qos Write channel QoS setting
func (r *registerAxi_ini3_write_qosType) GetAw_qos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini3_write_qosFieldAw_qosMask) >> RegisterAxi_ini3_write_qosFieldAw_qosShift)
}

// SetAw_qos Write channel QoS setting
func (r *registerAxi_ini3_write_qosType) SetAw_qos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini3_write_qosFieldAw_qosMask)|(uint32(value)<<RegisterAxi_ini3_write_qosFieldAw_qosShift))
}

// registerAxi_ini3_fn_modType AXI interconnect - INI x issuing functionality modification register
type registerAxi_ini3_fn_modType uint32

const (
	RegisterAxi_ini3_fn_modFieldRead_iss_overrideShift = 0
	RegisterAxi_ini3_fn_modFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override Override ASIB read issuing capability
func (r *registerAxi_ini3_fn_modType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini3_fn_modFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override Override ASIB read issuing capability
func (r *registerAxi_ini3_fn_modType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini3_fn_modFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini3_fn_modFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_ini3_fn_modFieldWrite_iss_overrideShift = 1
	RegisterAxi_ini3_fn_modFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Override ASIB write issuing capability
func (r *registerAxi_ini3_fn_modType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini3_fn_modFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Override ASIB write issuing capability
func (r *registerAxi_ini3_fn_modType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini3_fn_modFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini3_fn_modFieldWrite_iss_overrideMask)
	}
}

// registerAxi_ini4_read_qosType AXI interconnect - INI x read QoS register
type registerAxi_ini4_read_qosType uint32

const (
	RegisterAxi_ini4_read_qosFieldAr_qosShift = 0
	RegisterAxi_ini4_read_qosFieldAr_qosMask  = 0xf
)

// GetAr_qos Read channel QoS setting
func (r *registerAxi_ini4_read_qosType) GetAr_qos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini4_read_qosFieldAr_qosMask) >> RegisterAxi_ini4_read_qosFieldAr_qosShift)
}

// SetAr_qos Read channel QoS setting
func (r *registerAxi_ini4_read_qosType) SetAr_qos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini4_read_qosFieldAr_qosMask)|(uint32(value)<<RegisterAxi_ini4_read_qosFieldAr_qosShift))
}

// registerAxi_ini4_write_qosType AXI interconnect - INI x write QoS register
type registerAxi_ini4_write_qosType uint32

const (
	RegisterAxi_ini4_write_qosFieldAw_qosShift = 0
	RegisterAxi_ini4_write_qosFieldAw_qosMask  = 0xf
)

// GetAw_qos Write channel QoS setting
func (r *registerAxi_ini4_write_qosType) GetAw_qos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini4_write_qosFieldAw_qosMask) >> RegisterAxi_ini4_write_qosFieldAw_qosShift)
}

// SetAw_qos Write channel QoS setting
func (r *registerAxi_ini4_write_qosType) SetAw_qos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini4_write_qosFieldAw_qosMask)|(uint32(value)<<RegisterAxi_ini4_write_qosFieldAw_qosShift))
}

// registerAxi_ini4_fn_modType AXI interconnect - INI x issuing functionality modification register
type registerAxi_ini4_fn_modType uint32

const (
	RegisterAxi_ini4_fn_modFieldRead_iss_overrideShift = 0
	RegisterAxi_ini4_fn_modFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override Override ASIB read issuing capability
func (r *registerAxi_ini4_fn_modType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini4_fn_modFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override Override ASIB read issuing capability
func (r *registerAxi_ini4_fn_modType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini4_fn_modFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini4_fn_modFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_ini4_fn_modFieldWrite_iss_overrideShift = 1
	RegisterAxi_ini4_fn_modFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Override ASIB write issuing capability
func (r *registerAxi_ini4_fn_modType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini4_fn_modFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Override ASIB write issuing capability
func (r *registerAxi_ini4_fn_modType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini4_fn_modFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini4_fn_modFieldWrite_iss_overrideMask)
	}
}

// registerAxi_ini5_read_qosType AXI interconnect - INI x read QoS register
type registerAxi_ini5_read_qosType uint32

const (
	RegisterAxi_ini5_read_qosFieldAr_qosShift = 0
	RegisterAxi_ini5_read_qosFieldAr_qosMask  = 0xf
)

// GetAr_qos Read channel QoS setting
func (r *registerAxi_ini5_read_qosType) GetAr_qos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini5_read_qosFieldAr_qosMask) >> RegisterAxi_ini5_read_qosFieldAr_qosShift)
}

// SetAr_qos Read channel QoS setting
func (r *registerAxi_ini5_read_qosType) SetAr_qos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini5_read_qosFieldAr_qosMask)|(uint32(value)<<RegisterAxi_ini5_read_qosFieldAr_qosShift))
}

// registerAxi_ini5_write_qosType AXI interconnect - INI x write QoS register
type registerAxi_ini5_write_qosType uint32

const (
	RegisterAxi_ini5_write_qosFieldAw_qosShift = 0
	RegisterAxi_ini5_write_qosFieldAw_qosMask  = 0xf
)

// GetAw_qos Write channel QoS setting
func (r *registerAxi_ini5_write_qosType) GetAw_qos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini5_write_qosFieldAw_qosMask) >> RegisterAxi_ini5_write_qosFieldAw_qosShift)
}

// SetAw_qos Write channel QoS setting
func (r *registerAxi_ini5_write_qosType) SetAw_qos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini5_write_qosFieldAw_qosMask)|(uint32(value)<<RegisterAxi_ini5_write_qosFieldAw_qosShift))
}

// registerAxi_ini5_fn_modType AXI interconnect - INI x issuing functionality modification register
type registerAxi_ini5_fn_modType uint32

const (
	RegisterAxi_ini5_fn_modFieldRead_iss_overrideShift = 0
	RegisterAxi_ini5_fn_modFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override Override ASIB read issuing capability
func (r *registerAxi_ini5_fn_modType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini5_fn_modFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override Override ASIB read issuing capability
func (r *registerAxi_ini5_fn_modType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini5_fn_modFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini5_fn_modFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_ini5_fn_modFieldWrite_iss_overrideShift = 1
	RegisterAxi_ini5_fn_modFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Override ASIB write issuing capability
func (r *registerAxi_ini5_fn_modType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini5_fn_modFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Override ASIB write issuing capability
func (r *registerAxi_ini5_fn_modType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini5_fn_modFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini5_fn_modFieldWrite_iss_overrideMask)
	}
}

// registerAxi_ini6_read_qosType AXI interconnect - INI x read QoS register
type registerAxi_ini6_read_qosType uint32

const (
	RegisterAxi_ini6_read_qosFieldAr_qosShift = 0
	RegisterAxi_ini6_read_qosFieldAr_qosMask  = 0xf
)

// GetAr_qos Read channel QoS setting
func (r *registerAxi_ini6_read_qosType) GetAr_qos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini6_read_qosFieldAr_qosMask) >> RegisterAxi_ini6_read_qosFieldAr_qosShift)
}

// SetAr_qos Read channel QoS setting
func (r *registerAxi_ini6_read_qosType) SetAr_qos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini6_read_qosFieldAr_qosMask)|(uint32(value)<<RegisterAxi_ini6_read_qosFieldAr_qosShift))
}

// registerAxi_ini6_write_qosType AXI interconnect - INI x write QoS register
type registerAxi_ini6_write_qosType uint32

const (
	RegisterAxi_ini6_write_qosFieldAw_qosShift = 0
	RegisterAxi_ini6_write_qosFieldAw_qosMask  = 0xf
)

// GetAw_qos Write channel QoS setting
func (r *registerAxi_ini6_write_qosType) GetAw_qos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini6_write_qosFieldAw_qosMask) >> RegisterAxi_ini6_write_qosFieldAw_qosShift)
}

// SetAw_qos Write channel QoS setting
func (r *registerAxi_ini6_write_qosType) SetAw_qos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini6_write_qosFieldAw_qosMask)|(uint32(value)<<RegisterAxi_ini6_write_qosFieldAw_qosShift))
}

// registerAxi_ini6_fn_modType AXI interconnect - INI x issuing functionality modification register
type registerAxi_ini6_fn_modType uint32

const (
	RegisterAxi_ini6_fn_modFieldRead_iss_overrideShift = 0
	RegisterAxi_ini6_fn_modFieldRead_iss_overrideMask  = 0x1
)

// GetRead_iss_override Override ASIB read issuing capability
func (r *registerAxi_ini6_fn_modType) GetRead_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini6_fn_modFieldRead_iss_overrideMask) != 0
}

// SetRead_iss_override Override ASIB read issuing capability
func (r *registerAxi_ini6_fn_modType) SetRead_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini6_fn_modFieldRead_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini6_fn_modFieldRead_iss_overrideMask)
	}
}

const (
	RegisterAxi_ini6_fn_modFieldWrite_iss_overrideShift = 1
	RegisterAxi_ini6_fn_modFieldWrite_iss_overrideMask  = 0x2
)

// GetWrite_iss_override Override ASIB write issuing capability
func (r *registerAxi_ini6_fn_modType) GetWrite_iss_override() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAxi_ini6_fn_modFieldWrite_iss_overrideMask) != 0
}

// SetWrite_iss_override Override ASIB write issuing capability
func (r *registerAxi_ini6_fn_modType) SetWrite_iss_override(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAxi_ini6_fn_modFieldWrite_iss_overrideMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAxi_ini6_fn_modFieldWrite_iss_overrideMask)
	}
}
