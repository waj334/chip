package dfsdm

import (
	"unsafe"
	"volatile"
)

var (
	Dfsdm = (*_dfsdm)(unsafe.Pointer(uintptr(0x40017000)))
)

type _dfsdm struct {
	Dfsdm_chcfg0r1  registerDfsdm_chcfg0r1Type
	Dfsdm_chcfg1r1  registerDfsdm_chcfg1r1Type
	Dfsdm_chcfg2r1  registerDfsdm_chcfg2r1Type
	Dfsdm_chcfg3r1  registerDfsdm_chcfg3r1Type
	Dfsdm_chcfg4r1  registerDfsdm_chcfg4r1Type
	Dfsdm_chcfg5r1  registerDfsdm_chcfg5r1Type
	Dfsdm_chcfg6r1  registerDfsdm_chcfg6r1Type
	Dfsdm_chcfg7r1  registerDfsdm_chcfg7r1Type
	Dfsdm_chcfg0r2  registerDfsdm_chcfg0r2Type
	Dfsdm_chcfg1r2  registerDfsdm_chcfg1r2Type
	Dfsdm_chcfg2r2  registerDfsdm_chcfg2r2Type
	Dfsdm_chcfg3r2  registerDfsdm_chcfg3r2Type
	Dfsdm_chcfg4r2  registerDfsdm_chcfg4r2Type
	Dfsdm_chcfg5r2  registerDfsdm_chcfg5r2Type
	Dfsdm_chcfg6r2  registerDfsdm_chcfg6r2Type
	Dfsdm_chcfg7r2  registerDfsdm_chcfg7r2Type
	Dfsdm_awscd0r   registerDfsdm_awscd0rType
	Dfsdm_awscd1r   registerDfsdm_awscd1rType
	Dfsdm_awscd2r   registerDfsdm_awscd2rType
	Dfsdm_awscd3r   registerDfsdm_awscd3rType
	Dfsdm_awscd4r   registerDfsdm_awscd4rType
	Dfsdm_awscd5r   registerDfsdm_awscd5rType
	Dfsdm_awscd6r   registerDfsdm_awscd6rType
	Dfsdm_awscd7r   registerDfsdm_awscd7rType
	Dfsdm_chwdat0r  registerDfsdm_chwdat0rType
	Dfsdm_chwdat1r  registerDfsdm_chwdat1rType
	Dfsdm_chwdat2r  registerDfsdm_chwdat2rType
	Dfsdm_chwdat3r  registerDfsdm_chwdat3rType
	Dfsdm_chwdat4r  registerDfsdm_chwdat4rType
	Dfsdm_chwdat5r  registerDfsdm_chwdat5rType
	Dfsdm_chwdat6r  registerDfsdm_chwdat6rType
	Dfsdm_chwdat7r  registerDfsdm_chwdat7rType
	Dfsdm_chdatin0r registerDfsdm_chdatin0rType
	Dfsdm_chdatin1r registerDfsdm_chdatin1rType
	Dfsdm_chdatin2r registerDfsdm_chdatin2rType
	Dfsdm_chdatin3r registerDfsdm_chdatin3rType
	Dfsdm_chdatin4r registerDfsdm_chdatin4rType
	Dfsdm_chdatin5r registerDfsdm_chdatin5rType
	Dfsdm_chdatin6r registerDfsdm_chdatin6rType
	Dfsdm_chdatin7r registerDfsdm_chdatin7rType
	Dfsdm0_cr1      registerDfsdm0_cr1Type
	Dfsdm1_cr1      registerDfsdm1_cr1Type
	Dfsdm2_cr1      registerDfsdm2_cr1Type
	Dfsdm3_cr1      registerDfsdm3_cr1Type
	Dfsdm0_cr2      registerDfsdm0_cr2Type
	Dfsdm1_cr2      registerDfsdm1_cr2Type
	Dfsdm2_cr2      registerDfsdm2_cr2Type
	Dfsdm3_cr2      registerDfsdm3_cr2Type
	Dfsdm0_isr      registerDfsdm0_isrType
	Dfsdm1_isr      registerDfsdm1_isrType
	Dfsdm2_isr      registerDfsdm2_isrType
	Dfsdm3_isr      registerDfsdm3_isrType
	Dfsdm0_icr      registerDfsdm0_icrType
	Dfsdm1_icr      registerDfsdm1_icrType
	Dfsdm2_icr      registerDfsdm2_icrType
	Dfsdm3_icr      registerDfsdm3_icrType
	Dfsdm0_jchgr    registerDfsdm0_jchgrType
	Dfsdm1_jchgr    registerDfsdm1_jchgrType
	Dfsdm2_jchgr    registerDfsdm2_jchgrType
	Dfsdm3_jchgr    registerDfsdm3_jchgrType
	Dfsdm0_fcr      registerDfsdm0_fcrType
	Dfsdm1_fcr      registerDfsdm1_fcrType
	Dfsdm2_fcr      registerDfsdm2_fcrType
	Dfsdm3_fcr      registerDfsdm3_fcrType
	Dfsdm0_jdatar   registerDfsdm0_jdatarType
	Dfsdm1_jdatar   registerDfsdm1_jdatarType
	Dfsdm2_jdatar   registerDfsdm2_jdatarType
	Dfsdm3_jdatar   registerDfsdm3_jdatarType
	Dfsdm0_rdatar   registerDfsdm0_rdatarType
	Dfsdm1_rdatar   registerDfsdm1_rdatarType
	Dfsdm2_rdatar   registerDfsdm2_rdatarType
	Dfsdm3_rdatar   registerDfsdm3_rdatarType
	Dfsdm0_awhtr    registerDfsdm0_awhtrType
	Dfsdm1_awhtr    registerDfsdm1_awhtrType
	Dfsdm2_awhtr    registerDfsdm2_awhtrType
	Dfsdm3_awhtr    registerDfsdm3_awhtrType
	Dfsdm0_awltr    registerDfsdm0_awltrType
	Dfsdm1_awltr    registerDfsdm1_awltrType
	Dfsdm2_awltr    registerDfsdm2_awltrType
	Dfsdm3_awltr    registerDfsdm3_awltrType
	Dfsdm0_awsr     registerDfsdm0_awsrType
	Dfsdm1_awsr     registerDfsdm1_awsrType
	Dfsdm2_awsr     registerDfsdm2_awsrType
	Dfsdm3_awsr     registerDfsdm3_awsrType
	Dfsdm0_awcfr    registerDfsdm0_awcfrType
	Dfsdm1_awcfr    registerDfsdm1_awcfrType
	Dfsdm2_awcfr    registerDfsdm2_awcfrType
	Dfsdm3_awcfr    registerDfsdm3_awcfrType
	Dfsdm0_exmax    registerDfsdm0_exmaxType
	Dfsdm1_exmax    registerDfsdm1_exmaxType
	Dfsdm2_exmax    registerDfsdm2_exmaxType
	Dfsdm3_exmax    registerDfsdm3_exmaxType
	Dfsdm0_exmin    registerDfsdm0_exminType
	Dfsdm1_exmin    registerDfsdm1_exminType
	Dfsdm2_exmin    registerDfsdm2_exminType
	Dfsdm3_exmin    registerDfsdm3_exminType
	Dfsdm0_cnvtimr  registerDfsdm0_cnvtimrType
	Dfsdm1_cnvtimr  registerDfsdm1_cnvtimrType
	Dfsdm2_cnvtimr  registerDfsdm2_cnvtimrType
	Dfsdm3_cnvtimr  registerDfsdm3_cnvtimrType
}

// registerDfsdm_chcfg0r1Type DFSDM channel configuration 0 register 1
type registerDfsdm_chcfg0r1Type uint32

const (
	RegisterDfsdm_chcfg0r1FieldSitpShift = 0
	RegisterDfsdm_chcfg0r1FieldSitpMask  = 0x3
)

// GetSitp Serial interface type for channel 0
func (r *registerDfsdm_chcfg0r1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r1FieldSitpMask) >> RegisterDfsdm_chcfg0r1FieldSitpShift)
}

// SetSitp Serial interface type for channel 0
func (r *registerDfsdm_chcfg0r1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r1FieldSitpMask)|(uint32(value)<<RegisterDfsdm_chcfg0r1FieldSitpShift))
}

const (
	RegisterDfsdm_chcfg0r1FieldSpickselShift = 2
	RegisterDfsdm_chcfg0r1FieldSpickselMask  = 0xc
)

// GetSpicksel SPI clock select for channel 0
func (r *registerDfsdm_chcfg0r1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r1FieldSpickselMask) >> RegisterDfsdm_chcfg0r1FieldSpickselShift)
}

// SetSpicksel SPI clock select for channel 0
func (r *registerDfsdm_chcfg0r1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r1FieldSpickselMask)|(uint32(value)<<RegisterDfsdm_chcfg0r1FieldSpickselShift))
}

const (
	RegisterDfsdm_chcfg0r1FieldScdenShift = 5
	RegisterDfsdm_chcfg0r1FieldScdenMask  = 0x20
)

// GetScden Short-circuit detector enable on channel 0
func (r *registerDfsdm_chcfg0r1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r1FieldScdenMask) != 0
}

// SetScden Short-circuit detector enable on channel 0
func (r *registerDfsdm_chcfg0r1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg0r1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r1FieldScdenMask)
	}
}

const (
	RegisterDfsdm_chcfg0r1FieldCkabenShift = 6
	RegisterDfsdm_chcfg0r1FieldCkabenMask  = 0x40
)

// GetCkaben Clock absence detector enable on channel 0
func (r *registerDfsdm_chcfg0r1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r1FieldCkabenMask) != 0
}

// SetCkaben Clock absence detector enable on channel 0
func (r *registerDfsdm_chcfg0r1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg0r1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r1FieldCkabenMask)
	}
}

const (
	RegisterDfsdm_chcfg0r1FieldChenShift = 7
	RegisterDfsdm_chcfg0r1FieldChenMask  = 0x80
)

// GetChen Channel 0 enable
func (r *registerDfsdm_chcfg0r1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r1FieldChenMask) != 0
}

// SetChen Channel 0 enable
func (r *registerDfsdm_chcfg0r1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg0r1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r1FieldChenMask)
	}
}

const (
	RegisterDfsdm_chcfg0r1FieldChinselShift = 8
	RegisterDfsdm_chcfg0r1FieldChinselMask  = 0x100
)

// GetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg0r1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r1FieldChinselMask) != 0
}

// SetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg0r1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg0r1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r1FieldChinselMask)
	}
}

const (
	RegisterDfsdm_chcfg0r1FieldDatmpxShift = 12
	RegisterDfsdm_chcfg0r1FieldDatmpxMask  = 0x3000
)

// GetDatmpx Input data multiplexer for channel 0
func (r *registerDfsdm_chcfg0r1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r1FieldDatmpxMask) >> RegisterDfsdm_chcfg0r1FieldDatmpxShift)
}

// SetDatmpx Input data multiplexer for channel 0
func (r *registerDfsdm_chcfg0r1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r1FieldDatmpxMask)|(uint32(value)<<RegisterDfsdm_chcfg0r1FieldDatmpxShift))
}

const (
	RegisterDfsdm_chcfg0r1FieldDatpackShift = 14
	RegisterDfsdm_chcfg0r1FieldDatpackMask  = 0xc000
)

// GetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg0r1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r1FieldDatpackMask) >> RegisterDfsdm_chcfg0r1FieldDatpackShift)
}

// SetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg0r1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r1FieldDatpackMask)|(uint32(value)<<RegisterDfsdm_chcfg0r1FieldDatpackShift))
}

const (
	RegisterDfsdm_chcfg0r1FieldCkoutdivShift = 16
	RegisterDfsdm_chcfg0r1FieldCkoutdivMask  = 0xff0000
)

// GetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg0r1Type) GetCkoutdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r1FieldCkoutdivMask) >> RegisterDfsdm_chcfg0r1FieldCkoutdivShift)
}

// SetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg0r1Type) SetCkoutdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r1FieldCkoutdivMask)|(uint32(value)<<RegisterDfsdm_chcfg0r1FieldCkoutdivShift))
}

const (
	RegisterDfsdm_chcfg0r1FieldCkoutsrcShift = 30
	RegisterDfsdm_chcfg0r1FieldCkoutsrcMask  = 0x40000000
)

// GetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg0r1Type) GetCkoutsrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r1FieldCkoutsrcMask) != 0
}

// SetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg0r1Type) SetCkoutsrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg0r1FieldCkoutsrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r1FieldCkoutsrcMask)
	}
}

const (
	RegisterDfsdm_chcfg0r1FieldDfsdmenShift = 31
	RegisterDfsdm_chcfg0r1FieldDfsdmenMask  = 0x80000000
)

// GetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg0r1Type) GetDfsdmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r1FieldDfsdmenMask) != 0
}

// SetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg0r1Type) SetDfsdmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg0r1FieldDfsdmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r1FieldDfsdmenMask)
	}
}

// registerDfsdm_chcfg1r1Type DFSDM channel configuration 1 register 1
type registerDfsdm_chcfg1r1Type uint32

const (
	RegisterDfsdm_chcfg1r1FieldSitpShift = 0
	RegisterDfsdm_chcfg1r1FieldSitpMask  = 0x3
)

// GetSitp Serial interface type for channel 1
func (r *registerDfsdm_chcfg1r1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r1FieldSitpMask) >> RegisterDfsdm_chcfg1r1FieldSitpShift)
}

// SetSitp Serial interface type for channel 1
func (r *registerDfsdm_chcfg1r1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r1FieldSitpMask)|(uint32(value)<<RegisterDfsdm_chcfg1r1FieldSitpShift))
}

const (
	RegisterDfsdm_chcfg1r1FieldSpickselShift = 2
	RegisterDfsdm_chcfg1r1FieldSpickselMask  = 0xc
)

// GetSpicksel SPI clock select for channel 1
func (r *registerDfsdm_chcfg1r1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r1FieldSpickselMask) >> RegisterDfsdm_chcfg1r1FieldSpickselShift)
}

// SetSpicksel SPI clock select for channel 1
func (r *registerDfsdm_chcfg1r1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r1FieldSpickselMask)|(uint32(value)<<RegisterDfsdm_chcfg1r1FieldSpickselShift))
}

const (
	RegisterDfsdm_chcfg1r1FieldScdenShift = 5
	RegisterDfsdm_chcfg1r1FieldScdenMask  = 0x20
)

// GetScden Short-circuit detector enable on channel 1
func (r *registerDfsdm_chcfg1r1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r1FieldScdenMask) != 0
}

// SetScden Short-circuit detector enable on channel 1
func (r *registerDfsdm_chcfg1r1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg1r1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r1FieldScdenMask)
	}
}

const (
	RegisterDfsdm_chcfg1r1FieldCkabenShift = 6
	RegisterDfsdm_chcfg1r1FieldCkabenMask  = 0x40
)

// GetCkaben Clock absence detector enable on channel 1
func (r *registerDfsdm_chcfg1r1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r1FieldCkabenMask) != 0
}

// SetCkaben Clock absence detector enable on channel 1
func (r *registerDfsdm_chcfg1r1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg1r1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r1FieldCkabenMask)
	}
}

const (
	RegisterDfsdm_chcfg1r1FieldChenShift = 7
	RegisterDfsdm_chcfg1r1FieldChenMask  = 0x80
)

// GetChen Channel 1 enable
func (r *registerDfsdm_chcfg1r1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r1FieldChenMask) != 0
}

// SetChen Channel 1 enable
func (r *registerDfsdm_chcfg1r1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg1r1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r1FieldChenMask)
	}
}

const (
	RegisterDfsdm_chcfg1r1FieldChinselShift = 8
	RegisterDfsdm_chcfg1r1FieldChinselMask  = 0x100
)

// GetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg1r1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r1FieldChinselMask) != 0
}

// SetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg1r1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg1r1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r1FieldChinselMask)
	}
}

const (
	RegisterDfsdm_chcfg1r1FieldDatmpxShift = 12
	RegisterDfsdm_chcfg1r1FieldDatmpxMask  = 0x3000
)

// GetDatmpx Input data multiplexer for channel 1
func (r *registerDfsdm_chcfg1r1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r1FieldDatmpxMask) >> RegisterDfsdm_chcfg1r1FieldDatmpxShift)
}

// SetDatmpx Input data multiplexer for channel 1
func (r *registerDfsdm_chcfg1r1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r1FieldDatmpxMask)|(uint32(value)<<RegisterDfsdm_chcfg1r1FieldDatmpxShift))
}

const (
	RegisterDfsdm_chcfg1r1FieldDatpackShift = 14
	RegisterDfsdm_chcfg1r1FieldDatpackMask  = 0xc000
)

// GetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg1r1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r1FieldDatpackMask) >> RegisterDfsdm_chcfg1r1FieldDatpackShift)
}

// SetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg1r1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r1FieldDatpackMask)|(uint32(value)<<RegisterDfsdm_chcfg1r1FieldDatpackShift))
}

const (
	RegisterDfsdm_chcfg1r1FieldCkoutdivShift = 16
	RegisterDfsdm_chcfg1r1FieldCkoutdivMask  = 0xff0000
)

// GetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg1r1Type) GetCkoutdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r1FieldCkoutdivMask) >> RegisterDfsdm_chcfg1r1FieldCkoutdivShift)
}

// SetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg1r1Type) SetCkoutdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r1FieldCkoutdivMask)|(uint32(value)<<RegisterDfsdm_chcfg1r1FieldCkoutdivShift))
}

const (
	RegisterDfsdm_chcfg1r1FieldCkoutsrcShift = 30
	RegisterDfsdm_chcfg1r1FieldCkoutsrcMask  = 0x40000000
)

// GetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg1r1Type) GetCkoutsrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r1FieldCkoutsrcMask) != 0
}

// SetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg1r1Type) SetCkoutsrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg1r1FieldCkoutsrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r1FieldCkoutsrcMask)
	}
}

const (
	RegisterDfsdm_chcfg1r1FieldDfsdmenShift = 31
	RegisterDfsdm_chcfg1r1FieldDfsdmenMask  = 0x80000000
)

// GetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg1r1Type) GetDfsdmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r1FieldDfsdmenMask) != 0
}

// SetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg1r1Type) SetDfsdmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg1r1FieldDfsdmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r1FieldDfsdmenMask)
	}
}

// registerDfsdm_chcfg2r1Type DFSDM channel configuration 2 register 1
type registerDfsdm_chcfg2r1Type uint32

const (
	RegisterDfsdm_chcfg2r1FieldSitpShift = 0
	RegisterDfsdm_chcfg2r1FieldSitpMask  = 0x3
)

// GetSitp Serial interface type for channel 2
func (r *registerDfsdm_chcfg2r1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r1FieldSitpMask) >> RegisterDfsdm_chcfg2r1FieldSitpShift)
}

// SetSitp Serial interface type for channel 2
func (r *registerDfsdm_chcfg2r1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r1FieldSitpMask)|(uint32(value)<<RegisterDfsdm_chcfg2r1FieldSitpShift))
}

const (
	RegisterDfsdm_chcfg2r1FieldSpickselShift = 2
	RegisterDfsdm_chcfg2r1FieldSpickselMask  = 0xc
)

// GetSpicksel SPI clock select for channel 2
func (r *registerDfsdm_chcfg2r1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r1FieldSpickselMask) >> RegisterDfsdm_chcfg2r1FieldSpickselShift)
}

// SetSpicksel SPI clock select for channel 2
func (r *registerDfsdm_chcfg2r1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r1FieldSpickselMask)|(uint32(value)<<RegisterDfsdm_chcfg2r1FieldSpickselShift))
}

const (
	RegisterDfsdm_chcfg2r1FieldScdenShift = 5
	RegisterDfsdm_chcfg2r1FieldScdenMask  = 0x20
)

// GetScden Short-circuit detector enable on channel 2
func (r *registerDfsdm_chcfg2r1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r1FieldScdenMask) != 0
}

// SetScden Short-circuit detector enable on channel 2
func (r *registerDfsdm_chcfg2r1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg2r1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r1FieldScdenMask)
	}
}

const (
	RegisterDfsdm_chcfg2r1FieldCkabenShift = 6
	RegisterDfsdm_chcfg2r1FieldCkabenMask  = 0x40
)

// GetCkaben Clock absence detector enable on channel 2
func (r *registerDfsdm_chcfg2r1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r1FieldCkabenMask) != 0
}

// SetCkaben Clock absence detector enable on channel 2
func (r *registerDfsdm_chcfg2r1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg2r1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r1FieldCkabenMask)
	}
}

const (
	RegisterDfsdm_chcfg2r1FieldChenShift = 7
	RegisterDfsdm_chcfg2r1FieldChenMask  = 0x80
)

// GetChen Channel 2 enable
func (r *registerDfsdm_chcfg2r1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r1FieldChenMask) != 0
}

// SetChen Channel 2 enable
func (r *registerDfsdm_chcfg2r1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg2r1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r1FieldChenMask)
	}
}

const (
	RegisterDfsdm_chcfg2r1FieldChinselShift = 8
	RegisterDfsdm_chcfg2r1FieldChinselMask  = 0x100
)

// GetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg2r1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r1FieldChinselMask) != 0
}

// SetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg2r1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg2r1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r1FieldChinselMask)
	}
}

const (
	RegisterDfsdm_chcfg2r1FieldDatmpxShift = 12
	RegisterDfsdm_chcfg2r1FieldDatmpxMask  = 0x3000
)

// GetDatmpx Input data multiplexer for channel 2
func (r *registerDfsdm_chcfg2r1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r1FieldDatmpxMask) >> RegisterDfsdm_chcfg2r1FieldDatmpxShift)
}

// SetDatmpx Input data multiplexer for channel 2
func (r *registerDfsdm_chcfg2r1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r1FieldDatmpxMask)|(uint32(value)<<RegisterDfsdm_chcfg2r1FieldDatmpxShift))
}

const (
	RegisterDfsdm_chcfg2r1FieldDatpackShift = 14
	RegisterDfsdm_chcfg2r1FieldDatpackMask  = 0xc000
)

// GetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg2r1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r1FieldDatpackMask) >> RegisterDfsdm_chcfg2r1FieldDatpackShift)
}

// SetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg2r1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r1FieldDatpackMask)|(uint32(value)<<RegisterDfsdm_chcfg2r1FieldDatpackShift))
}

const (
	RegisterDfsdm_chcfg2r1FieldCkoutdivShift = 16
	RegisterDfsdm_chcfg2r1FieldCkoutdivMask  = 0xff0000
)

// GetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg2r1Type) GetCkoutdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r1FieldCkoutdivMask) >> RegisterDfsdm_chcfg2r1FieldCkoutdivShift)
}

// SetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg2r1Type) SetCkoutdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r1FieldCkoutdivMask)|(uint32(value)<<RegisterDfsdm_chcfg2r1FieldCkoutdivShift))
}

const (
	RegisterDfsdm_chcfg2r1FieldCkoutsrcShift = 30
	RegisterDfsdm_chcfg2r1FieldCkoutsrcMask  = 0x40000000
)

// GetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg2r1Type) GetCkoutsrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r1FieldCkoutsrcMask) != 0
}

// SetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg2r1Type) SetCkoutsrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg2r1FieldCkoutsrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r1FieldCkoutsrcMask)
	}
}

const (
	RegisterDfsdm_chcfg2r1FieldDfsdmenShift = 31
	RegisterDfsdm_chcfg2r1FieldDfsdmenMask  = 0x80000000
)

// GetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg2r1Type) GetDfsdmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r1FieldDfsdmenMask) != 0
}

// SetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg2r1Type) SetDfsdmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg2r1FieldDfsdmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r1FieldDfsdmenMask)
	}
}

// registerDfsdm_chcfg3r1Type DFSDM channel configuration 3 register 1
type registerDfsdm_chcfg3r1Type uint32

const (
	RegisterDfsdm_chcfg3r1FieldSitpShift = 0
	RegisterDfsdm_chcfg3r1FieldSitpMask  = 0x3
)

// GetSitp Serial interface type for channel 3
func (r *registerDfsdm_chcfg3r1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r1FieldSitpMask) >> RegisterDfsdm_chcfg3r1FieldSitpShift)
}

// SetSitp Serial interface type for channel 3
func (r *registerDfsdm_chcfg3r1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r1FieldSitpMask)|(uint32(value)<<RegisterDfsdm_chcfg3r1FieldSitpShift))
}

const (
	RegisterDfsdm_chcfg3r1FieldSpickselShift = 2
	RegisterDfsdm_chcfg3r1FieldSpickselMask  = 0xc
)

// GetSpicksel SPI clock select for channel 3
func (r *registerDfsdm_chcfg3r1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r1FieldSpickselMask) >> RegisterDfsdm_chcfg3r1FieldSpickselShift)
}

// SetSpicksel SPI clock select for channel 3
func (r *registerDfsdm_chcfg3r1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r1FieldSpickselMask)|(uint32(value)<<RegisterDfsdm_chcfg3r1FieldSpickselShift))
}

const (
	RegisterDfsdm_chcfg3r1FieldScdenShift = 5
	RegisterDfsdm_chcfg3r1FieldScdenMask  = 0x20
)

// GetScden Short-circuit detector enable on channel 3
func (r *registerDfsdm_chcfg3r1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r1FieldScdenMask) != 0
}

// SetScden Short-circuit detector enable on channel 3
func (r *registerDfsdm_chcfg3r1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg3r1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r1FieldScdenMask)
	}
}

const (
	RegisterDfsdm_chcfg3r1FieldCkabenShift = 6
	RegisterDfsdm_chcfg3r1FieldCkabenMask  = 0x40
)

// GetCkaben Clock absence detector enable on channel 3
func (r *registerDfsdm_chcfg3r1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r1FieldCkabenMask) != 0
}

// SetCkaben Clock absence detector enable on channel 3
func (r *registerDfsdm_chcfg3r1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg3r1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r1FieldCkabenMask)
	}
}

const (
	RegisterDfsdm_chcfg3r1FieldChenShift = 7
	RegisterDfsdm_chcfg3r1FieldChenMask  = 0x80
)

// GetChen Channel 3 enable
func (r *registerDfsdm_chcfg3r1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r1FieldChenMask) != 0
}

// SetChen Channel 3 enable
func (r *registerDfsdm_chcfg3r1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg3r1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r1FieldChenMask)
	}
}

const (
	RegisterDfsdm_chcfg3r1FieldChinselShift = 8
	RegisterDfsdm_chcfg3r1FieldChinselMask  = 0x100
)

// GetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg3r1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r1FieldChinselMask) != 0
}

// SetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg3r1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg3r1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r1FieldChinselMask)
	}
}

const (
	RegisterDfsdm_chcfg3r1FieldDatmpxShift = 12
	RegisterDfsdm_chcfg3r1FieldDatmpxMask  = 0x3000
)

// GetDatmpx Input data multiplexer for channel 3
func (r *registerDfsdm_chcfg3r1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r1FieldDatmpxMask) >> RegisterDfsdm_chcfg3r1FieldDatmpxShift)
}

// SetDatmpx Input data multiplexer for channel 3
func (r *registerDfsdm_chcfg3r1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r1FieldDatmpxMask)|(uint32(value)<<RegisterDfsdm_chcfg3r1FieldDatmpxShift))
}

const (
	RegisterDfsdm_chcfg3r1FieldDatpackShift = 14
	RegisterDfsdm_chcfg3r1FieldDatpackMask  = 0xc000
)

// GetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg3r1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r1FieldDatpackMask) >> RegisterDfsdm_chcfg3r1FieldDatpackShift)
}

// SetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg3r1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r1FieldDatpackMask)|(uint32(value)<<RegisterDfsdm_chcfg3r1FieldDatpackShift))
}

const (
	RegisterDfsdm_chcfg3r1FieldCkoutdivShift = 16
	RegisterDfsdm_chcfg3r1FieldCkoutdivMask  = 0xff0000
)

// GetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg3r1Type) GetCkoutdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r1FieldCkoutdivMask) >> RegisterDfsdm_chcfg3r1FieldCkoutdivShift)
}

// SetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg3r1Type) SetCkoutdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r1FieldCkoutdivMask)|(uint32(value)<<RegisterDfsdm_chcfg3r1FieldCkoutdivShift))
}

const (
	RegisterDfsdm_chcfg3r1FieldCkoutsrcShift = 30
	RegisterDfsdm_chcfg3r1FieldCkoutsrcMask  = 0x40000000
)

// GetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg3r1Type) GetCkoutsrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r1FieldCkoutsrcMask) != 0
}

// SetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg3r1Type) SetCkoutsrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg3r1FieldCkoutsrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r1FieldCkoutsrcMask)
	}
}

const (
	RegisterDfsdm_chcfg3r1FieldDfsdmenShift = 31
	RegisterDfsdm_chcfg3r1FieldDfsdmenMask  = 0x80000000
)

// GetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg3r1Type) GetDfsdmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r1FieldDfsdmenMask) != 0
}

// SetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg3r1Type) SetDfsdmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg3r1FieldDfsdmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r1FieldDfsdmenMask)
	}
}

// registerDfsdm_chcfg4r1Type DFSDM channel configuration 4 register 1
type registerDfsdm_chcfg4r1Type uint32

const (
	RegisterDfsdm_chcfg4r1FieldSitpShift = 0
	RegisterDfsdm_chcfg4r1FieldSitpMask  = 0x3
)

// GetSitp Serial interface type for channel 4
func (r *registerDfsdm_chcfg4r1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r1FieldSitpMask) >> RegisterDfsdm_chcfg4r1FieldSitpShift)
}

// SetSitp Serial interface type for channel 4
func (r *registerDfsdm_chcfg4r1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r1FieldSitpMask)|(uint32(value)<<RegisterDfsdm_chcfg4r1FieldSitpShift))
}

const (
	RegisterDfsdm_chcfg4r1FieldSpickselShift = 2
	RegisterDfsdm_chcfg4r1FieldSpickselMask  = 0xc
)

// GetSpicksel SPI clock select for channel 4
func (r *registerDfsdm_chcfg4r1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r1FieldSpickselMask) >> RegisterDfsdm_chcfg4r1FieldSpickselShift)
}

// SetSpicksel SPI clock select for channel 4
func (r *registerDfsdm_chcfg4r1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r1FieldSpickselMask)|(uint32(value)<<RegisterDfsdm_chcfg4r1FieldSpickselShift))
}

const (
	RegisterDfsdm_chcfg4r1FieldScdenShift = 5
	RegisterDfsdm_chcfg4r1FieldScdenMask  = 0x20
)

// GetScden Short-circuit detector enable on channel 4
func (r *registerDfsdm_chcfg4r1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r1FieldScdenMask) != 0
}

// SetScden Short-circuit detector enable on channel 4
func (r *registerDfsdm_chcfg4r1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg4r1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r1FieldScdenMask)
	}
}

const (
	RegisterDfsdm_chcfg4r1FieldCkabenShift = 6
	RegisterDfsdm_chcfg4r1FieldCkabenMask  = 0x40
)

// GetCkaben Clock absence detector enable on channel 4
func (r *registerDfsdm_chcfg4r1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r1FieldCkabenMask) != 0
}

// SetCkaben Clock absence detector enable on channel 4
func (r *registerDfsdm_chcfg4r1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg4r1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r1FieldCkabenMask)
	}
}

const (
	RegisterDfsdm_chcfg4r1FieldChenShift = 7
	RegisterDfsdm_chcfg4r1FieldChenMask  = 0x80
)

// GetChen Channel 4 enable
func (r *registerDfsdm_chcfg4r1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r1FieldChenMask) != 0
}

// SetChen Channel 4 enable
func (r *registerDfsdm_chcfg4r1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg4r1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r1FieldChenMask)
	}
}

const (
	RegisterDfsdm_chcfg4r1FieldChinselShift = 8
	RegisterDfsdm_chcfg4r1FieldChinselMask  = 0x100
)

// GetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg4r1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r1FieldChinselMask) != 0
}

// SetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg4r1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg4r1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r1FieldChinselMask)
	}
}

const (
	RegisterDfsdm_chcfg4r1FieldDatmpxShift = 12
	RegisterDfsdm_chcfg4r1FieldDatmpxMask  = 0x3000
)

// GetDatmpx Input data multiplexer for channel 4
func (r *registerDfsdm_chcfg4r1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r1FieldDatmpxMask) >> RegisterDfsdm_chcfg4r1FieldDatmpxShift)
}

// SetDatmpx Input data multiplexer for channel 4
func (r *registerDfsdm_chcfg4r1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r1FieldDatmpxMask)|(uint32(value)<<RegisterDfsdm_chcfg4r1FieldDatmpxShift))
}

const (
	RegisterDfsdm_chcfg4r1FieldDatpackShift = 14
	RegisterDfsdm_chcfg4r1FieldDatpackMask  = 0xc000
)

// GetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg4r1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r1FieldDatpackMask) >> RegisterDfsdm_chcfg4r1FieldDatpackShift)
}

// SetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg4r1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r1FieldDatpackMask)|(uint32(value)<<RegisterDfsdm_chcfg4r1FieldDatpackShift))
}

const (
	RegisterDfsdm_chcfg4r1FieldCkoutdivShift = 16
	RegisterDfsdm_chcfg4r1FieldCkoutdivMask  = 0xff0000
)

// GetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg4r1Type) GetCkoutdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r1FieldCkoutdivMask) >> RegisterDfsdm_chcfg4r1FieldCkoutdivShift)
}

// SetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg4r1Type) SetCkoutdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r1FieldCkoutdivMask)|(uint32(value)<<RegisterDfsdm_chcfg4r1FieldCkoutdivShift))
}

const (
	RegisterDfsdm_chcfg4r1FieldCkoutsrcShift = 30
	RegisterDfsdm_chcfg4r1FieldCkoutsrcMask  = 0x40000000
)

// GetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg4r1Type) GetCkoutsrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r1FieldCkoutsrcMask) != 0
}

// SetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg4r1Type) SetCkoutsrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg4r1FieldCkoutsrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r1FieldCkoutsrcMask)
	}
}

const (
	RegisterDfsdm_chcfg4r1FieldDfsdmenShift = 31
	RegisterDfsdm_chcfg4r1FieldDfsdmenMask  = 0x80000000
)

// GetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg4r1Type) GetDfsdmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r1FieldDfsdmenMask) != 0
}

// SetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg4r1Type) SetDfsdmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg4r1FieldDfsdmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r1FieldDfsdmenMask)
	}
}

// registerDfsdm_chcfg5r1Type DFSDM channel configuration 5 register 1
type registerDfsdm_chcfg5r1Type uint32

const (
	RegisterDfsdm_chcfg5r1FieldSitpShift = 0
	RegisterDfsdm_chcfg5r1FieldSitpMask  = 0x3
)

// GetSitp Serial interface type for channel 5
func (r *registerDfsdm_chcfg5r1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r1FieldSitpMask) >> RegisterDfsdm_chcfg5r1FieldSitpShift)
}

// SetSitp Serial interface type for channel 5
func (r *registerDfsdm_chcfg5r1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r1FieldSitpMask)|(uint32(value)<<RegisterDfsdm_chcfg5r1FieldSitpShift))
}

const (
	RegisterDfsdm_chcfg5r1FieldSpickselShift = 2
	RegisterDfsdm_chcfg5r1FieldSpickselMask  = 0xc
)

// GetSpicksel SPI clock select for channel 5
func (r *registerDfsdm_chcfg5r1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r1FieldSpickselMask) >> RegisterDfsdm_chcfg5r1FieldSpickselShift)
}

// SetSpicksel SPI clock select for channel 5
func (r *registerDfsdm_chcfg5r1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r1FieldSpickselMask)|(uint32(value)<<RegisterDfsdm_chcfg5r1FieldSpickselShift))
}

const (
	RegisterDfsdm_chcfg5r1FieldScdenShift = 5
	RegisterDfsdm_chcfg5r1FieldScdenMask  = 0x20
)

// GetScden Short-circuit detector enable on channel 5
func (r *registerDfsdm_chcfg5r1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r1FieldScdenMask) != 0
}

// SetScden Short-circuit detector enable on channel 5
func (r *registerDfsdm_chcfg5r1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg5r1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r1FieldScdenMask)
	}
}

const (
	RegisterDfsdm_chcfg5r1FieldCkabenShift = 6
	RegisterDfsdm_chcfg5r1FieldCkabenMask  = 0x40
)

// GetCkaben Clock absence detector enable on channel 5
func (r *registerDfsdm_chcfg5r1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r1FieldCkabenMask) != 0
}

// SetCkaben Clock absence detector enable on channel 5
func (r *registerDfsdm_chcfg5r1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg5r1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r1FieldCkabenMask)
	}
}

const (
	RegisterDfsdm_chcfg5r1FieldChenShift = 7
	RegisterDfsdm_chcfg5r1FieldChenMask  = 0x80
)

// GetChen Channel 5 enable
func (r *registerDfsdm_chcfg5r1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r1FieldChenMask) != 0
}

// SetChen Channel 5 enable
func (r *registerDfsdm_chcfg5r1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg5r1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r1FieldChenMask)
	}
}

const (
	RegisterDfsdm_chcfg5r1FieldChinselShift = 8
	RegisterDfsdm_chcfg5r1FieldChinselMask  = 0x100
)

// GetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg5r1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r1FieldChinselMask) != 0
}

// SetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg5r1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg5r1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r1FieldChinselMask)
	}
}

const (
	RegisterDfsdm_chcfg5r1FieldDatmpxShift = 12
	RegisterDfsdm_chcfg5r1FieldDatmpxMask  = 0x3000
)

// GetDatmpx Input data multiplexer for channel 5
func (r *registerDfsdm_chcfg5r1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r1FieldDatmpxMask) >> RegisterDfsdm_chcfg5r1FieldDatmpxShift)
}

// SetDatmpx Input data multiplexer for channel 5
func (r *registerDfsdm_chcfg5r1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r1FieldDatmpxMask)|(uint32(value)<<RegisterDfsdm_chcfg5r1FieldDatmpxShift))
}

const (
	RegisterDfsdm_chcfg5r1FieldDatpackShift = 14
	RegisterDfsdm_chcfg5r1FieldDatpackMask  = 0xc000
)

// GetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg5r1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r1FieldDatpackMask) >> RegisterDfsdm_chcfg5r1FieldDatpackShift)
}

// SetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg5r1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r1FieldDatpackMask)|(uint32(value)<<RegisterDfsdm_chcfg5r1FieldDatpackShift))
}

const (
	RegisterDfsdm_chcfg5r1FieldCkoutdivShift = 16
	RegisterDfsdm_chcfg5r1FieldCkoutdivMask  = 0xff0000
)

// GetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg5r1Type) GetCkoutdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r1FieldCkoutdivMask) >> RegisterDfsdm_chcfg5r1FieldCkoutdivShift)
}

// SetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg5r1Type) SetCkoutdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r1FieldCkoutdivMask)|(uint32(value)<<RegisterDfsdm_chcfg5r1FieldCkoutdivShift))
}

const (
	RegisterDfsdm_chcfg5r1FieldCkoutsrcShift = 30
	RegisterDfsdm_chcfg5r1FieldCkoutsrcMask  = 0x40000000
)

// GetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg5r1Type) GetCkoutsrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r1FieldCkoutsrcMask) != 0
}

// SetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg5r1Type) SetCkoutsrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg5r1FieldCkoutsrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r1FieldCkoutsrcMask)
	}
}

const (
	RegisterDfsdm_chcfg5r1FieldDfsdmenShift = 31
	RegisterDfsdm_chcfg5r1FieldDfsdmenMask  = 0x80000000
)

// GetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg5r1Type) GetDfsdmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r1FieldDfsdmenMask) != 0
}

// SetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg5r1Type) SetDfsdmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg5r1FieldDfsdmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r1FieldDfsdmenMask)
	}
}

// registerDfsdm_chcfg6r1Type DFSDM channel configuration 6 register 1
type registerDfsdm_chcfg6r1Type uint32

const (
	RegisterDfsdm_chcfg6r1FieldSitpShift = 0
	RegisterDfsdm_chcfg6r1FieldSitpMask  = 0x3
)

// GetSitp Serial interface type for channel 6
func (r *registerDfsdm_chcfg6r1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r1FieldSitpMask) >> RegisterDfsdm_chcfg6r1FieldSitpShift)
}

// SetSitp Serial interface type for channel 6
func (r *registerDfsdm_chcfg6r1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r1FieldSitpMask)|(uint32(value)<<RegisterDfsdm_chcfg6r1FieldSitpShift))
}

const (
	RegisterDfsdm_chcfg6r1FieldSpickselShift = 2
	RegisterDfsdm_chcfg6r1FieldSpickselMask  = 0xc
)

// GetSpicksel SPI clock select for channel 6
func (r *registerDfsdm_chcfg6r1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r1FieldSpickselMask) >> RegisterDfsdm_chcfg6r1FieldSpickselShift)
}

// SetSpicksel SPI clock select for channel 6
func (r *registerDfsdm_chcfg6r1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r1FieldSpickselMask)|(uint32(value)<<RegisterDfsdm_chcfg6r1FieldSpickselShift))
}

const (
	RegisterDfsdm_chcfg6r1FieldScdenShift = 5
	RegisterDfsdm_chcfg6r1FieldScdenMask  = 0x20
)

// GetScden Short-circuit detector enable on channel 6
func (r *registerDfsdm_chcfg6r1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r1FieldScdenMask) != 0
}

// SetScden Short-circuit detector enable on channel 6
func (r *registerDfsdm_chcfg6r1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg6r1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r1FieldScdenMask)
	}
}

const (
	RegisterDfsdm_chcfg6r1FieldCkabenShift = 6
	RegisterDfsdm_chcfg6r1FieldCkabenMask  = 0x40
)

// GetCkaben Clock absence detector enable on channel 6
func (r *registerDfsdm_chcfg6r1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r1FieldCkabenMask) != 0
}

// SetCkaben Clock absence detector enable on channel 6
func (r *registerDfsdm_chcfg6r1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg6r1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r1FieldCkabenMask)
	}
}

const (
	RegisterDfsdm_chcfg6r1FieldChenShift = 7
	RegisterDfsdm_chcfg6r1FieldChenMask  = 0x80
)

// GetChen Channel 6 enable
func (r *registerDfsdm_chcfg6r1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r1FieldChenMask) != 0
}

// SetChen Channel 6 enable
func (r *registerDfsdm_chcfg6r1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg6r1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r1FieldChenMask)
	}
}

const (
	RegisterDfsdm_chcfg6r1FieldChinselShift = 8
	RegisterDfsdm_chcfg6r1FieldChinselMask  = 0x100
)

// GetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg6r1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r1FieldChinselMask) != 0
}

// SetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg6r1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg6r1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r1FieldChinselMask)
	}
}

const (
	RegisterDfsdm_chcfg6r1FieldDatmpxShift = 12
	RegisterDfsdm_chcfg6r1FieldDatmpxMask  = 0x3000
)

// GetDatmpx Input data multiplexer for channel 6
func (r *registerDfsdm_chcfg6r1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r1FieldDatmpxMask) >> RegisterDfsdm_chcfg6r1FieldDatmpxShift)
}

// SetDatmpx Input data multiplexer for channel 6
func (r *registerDfsdm_chcfg6r1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r1FieldDatmpxMask)|(uint32(value)<<RegisterDfsdm_chcfg6r1FieldDatmpxShift))
}

const (
	RegisterDfsdm_chcfg6r1FieldDatpackShift = 14
	RegisterDfsdm_chcfg6r1FieldDatpackMask  = 0xc000
)

// GetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg6r1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r1FieldDatpackMask) >> RegisterDfsdm_chcfg6r1FieldDatpackShift)
}

// SetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg6r1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r1FieldDatpackMask)|(uint32(value)<<RegisterDfsdm_chcfg6r1FieldDatpackShift))
}

const (
	RegisterDfsdm_chcfg6r1FieldCkoutdivShift = 16
	RegisterDfsdm_chcfg6r1FieldCkoutdivMask  = 0xff0000
)

// GetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg6r1Type) GetCkoutdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r1FieldCkoutdivMask) >> RegisterDfsdm_chcfg6r1FieldCkoutdivShift)
}

// SetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg6r1Type) SetCkoutdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r1FieldCkoutdivMask)|(uint32(value)<<RegisterDfsdm_chcfg6r1FieldCkoutdivShift))
}

const (
	RegisterDfsdm_chcfg6r1FieldCkoutsrcShift = 30
	RegisterDfsdm_chcfg6r1FieldCkoutsrcMask  = 0x40000000
)

// GetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg6r1Type) GetCkoutsrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r1FieldCkoutsrcMask) != 0
}

// SetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg6r1Type) SetCkoutsrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg6r1FieldCkoutsrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r1FieldCkoutsrcMask)
	}
}

const (
	RegisterDfsdm_chcfg6r1FieldDfsdmenShift = 31
	RegisterDfsdm_chcfg6r1FieldDfsdmenMask  = 0x80000000
)

// GetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg6r1Type) GetDfsdmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r1FieldDfsdmenMask) != 0
}

// SetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg6r1Type) SetDfsdmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg6r1FieldDfsdmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r1FieldDfsdmenMask)
	}
}

// registerDfsdm_chcfg7r1Type DFSDM channel configuration 7 register 1
type registerDfsdm_chcfg7r1Type uint32

const (
	RegisterDfsdm_chcfg7r1FieldSitpShift = 0
	RegisterDfsdm_chcfg7r1FieldSitpMask  = 0x3
)

// GetSitp Serial interface type for channel 7
func (r *registerDfsdm_chcfg7r1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r1FieldSitpMask) >> RegisterDfsdm_chcfg7r1FieldSitpShift)
}

// SetSitp Serial interface type for channel 7
func (r *registerDfsdm_chcfg7r1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r1FieldSitpMask)|(uint32(value)<<RegisterDfsdm_chcfg7r1FieldSitpShift))
}

const (
	RegisterDfsdm_chcfg7r1FieldSpickselShift = 2
	RegisterDfsdm_chcfg7r1FieldSpickselMask  = 0xc
)

// GetSpicksel SPI clock select for channel 7
func (r *registerDfsdm_chcfg7r1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r1FieldSpickselMask) >> RegisterDfsdm_chcfg7r1FieldSpickselShift)
}

// SetSpicksel SPI clock select for channel 7
func (r *registerDfsdm_chcfg7r1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r1FieldSpickselMask)|(uint32(value)<<RegisterDfsdm_chcfg7r1FieldSpickselShift))
}

const (
	RegisterDfsdm_chcfg7r1FieldScdenShift = 5
	RegisterDfsdm_chcfg7r1FieldScdenMask  = 0x20
)

// GetScden Short-circuit detector enable on channel 7
func (r *registerDfsdm_chcfg7r1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r1FieldScdenMask) != 0
}

// SetScden Short-circuit detector enable on channel 7
func (r *registerDfsdm_chcfg7r1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg7r1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r1FieldScdenMask)
	}
}

const (
	RegisterDfsdm_chcfg7r1FieldCkabenShift = 6
	RegisterDfsdm_chcfg7r1FieldCkabenMask  = 0x40
)

// GetCkaben Clock absence detector enable on channel 7
func (r *registerDfsdm_chcfg7r1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r1FieldCkabenMask) != 0
}

// SetCkaben Clock absence detector enable on channel 7
func (r *registerDfsdm_chcfg7r1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg7r1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r1FieldCkabenMask)
	}
}

const (
	RegisterDfsdm_chcfg7r1FieldChenShift = 7
	RegisterDfsdm_chcfg7r1FieldChenMask  = 0x80
)

// GetChen Channel 7 enable
func (r *registerDfsdm_chcfg7r1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r1FieldChenMask) != 0
}

// SetChen Channel 7 enable
func (r *registerDfsdm_chcfg7r1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg7r1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r1FieldChenMask)
	}
}

const (
	RegisterDfsdm_chcfg7r1FieldChinselShift = 8
	RegisterDfsdm_chcfg7r1FieldChinselMask  = 0x100
)

// GetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg7r1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r1FieldChinselMask) != 0
}

// SetChinsel Channel inputs selection
func (r *registerDfsdm_chcfg7r1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg7r1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r1FieldChinselMask)
	}
}

const (
	RegisterDfsdm_chcfg7r1FieldDatmpxShift = 12
	RegisterDfsdm_chcfg7r1FieldDatmpxMask  = 0x3000
)

// GetDatmpx Input data multiplexer for channel 7
func (r *registerDfsdm_chcfg7r1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r1FieldDatmpxMask) >> RegisterDfsdm_chcfg7r1FieldDatmpxShift)
}

// SetDatmpx Input data multiplexer for channel 7
func (r *registerDfsdm_chcfg7r1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r1FieldDatmpxMask)|(uint32(value)<<RegisterDfsdm_chcfg7r1FieldDatmpxShift))
}

const (
	RegisterDfsdm_chcfg7r1FieldDatpackShift = 14
	RegisterDfsdm_chcfg7r1FieldDatpackMask  = 0xc000
)

// GetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg7r1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r1FieldDatpackMask) >> RegisterDfsdm_chcfg7r1FieldDatpackShift)
}

// SetDatpack Data packing mode in DFSDM_CHDATINyR register
func (r *registerDfsdm_chcfg7r1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r1FieldDatpackMask)|(uint32(value)<<RegisterDfsdm_chcfg7r1FieldDatpackShift))
}

const (
	RegisterDfsdm_chcfg7r1FieldCkoutdivShift = 16
	RegisterDfsdm_chcfg7r1FieldCkoutdivMask  = 0xff0000
)

// GetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg7r1Type) GetCkoutdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r1FieldCkoutdivMask) >> RegisterDfsdm_chcfg7r1FieldCkoutdivShift)
}

// SetCkoutdiv Output serial clock divider
func (r *registerDfsdm_chcfg7r1Type) SetCkoutdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r1FieldCkoutdivMask)|(uint32(value)<<RegisterDfsdm_chcfg7r1FieldCkoutdivShift))
}

const (
	RegisterDfsdm_chcfg7r1FieldCkoutsrcShift = 30
	RegisterDfsdm_chcfg7r1FieldCkoutsrcMask  = 0x40000000
)

// GetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg7r1Type) GetCkoutsrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r1FieldCkoutsrcMask) != 0
}

// SetCkoutsrc Output serial clock source selection
func (r *registerDfsdm_chcfg7r1Type) SetCkoutsrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg7r1FieldCkoutsrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r1FieldCkoutsrcMask)
	}
}

const (
	RegisterDfsdm_chcfg7r1FieldDfsdmenShift = 31
	RegisterDfsdm_chcfg7r1FieldDfsdmenMask  = 0x80000000
)

// GetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg7r1Type) GetDfsdmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r1FieldDfsdmenMask) != 0
}

// SetDfsdmen Global enable for DFSDM interface
func (r *registerDfsdm_chcfg7r1Type) SetDfsdmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_chcfg7r1FieldDfsdmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r1FieldDfsdmenMask)
	}
}

// registerDfsdm_chcfg0r2Type DFSDM channel configuration 0 register 2
type registerDfsdm_chcfg0r2Type uint32

const (
	RegisterDfsdm_chcfg0r2FieldDtrbsShift = 3
	RegisterDfsdm_chcfg0r2FieldDtrbsMask  = 0xf8
)

// GetDtrbs Data right bit-shift for channel 0
func (r *registerDfsdm_chcfg0r2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r2FieldDtrbsMask) >> RegisterDfsdm_chcfg0r2FieldDtrbsShift)
}

// SetDtrbs Data right bit-shift for channel 0
func (r *registerDfsdm_chcfg0r2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r2FieldDtrbsMask)|(uint32(value)<<RegisterDfsdm_chcfg0r2FieldDtrbsShift))
}

const (
	RegisterDfsdm_chcfg0r2FieldOffsetShift = 8
	RegisterDfsdm_chcfg0r2FieldOffsetMask  = 0xffffff00
)

// GetOffset 24-bit calibration offset for channel 0
func (r *registerDfsdm_chcfg0r2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg0r2FieldOffsetMask) >> RegisterDfsdm_chcfg0r2FieldOffsetShift)
}

// SetOffset 24-bit calibration offset for channel 0
func (r *registerDfsdm_chcfg0r2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg0r2FieldOffsetMask)|(uint32(value)<<RegisterDfsdm_chcfg0r2FieldOffsetShift))
}

// registerDfsdm_chcfg1r2Type DFSDM channel configuration 1 register 2
type registerDfsdm_chcfg1r2Type uint32

const (
	RegisterDfsdm_chcfg1r2FieldDtrbsShift = 3
	RegisterDfsdm_chcfg1r2FieldDtrbsMask  = 0xf8
)

// GetDtrbs Data right bit-shift for channel 1
func (r *registerDfsdm_chcfg1r2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r2FieldDtrbsMask) >> RegisterDfsdm_chcfg1r2FieldDtrbsShift)
}

// SetDtrbs Data right bit-shift for channel 1
func (r *registerDfsdm_chcfg1r2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r2FieldDtrbsMask)|(uint32(value)<<RegisterDfsdm_chcfg1r2FieldDtrbsShift))
}

const (
	RegisterDfsdm_chcfg1r2FieldOffsetShift = 8
	RegisterDfsdm_chcfg1r2FieldOffsetMask  = 0xffffff00
)

// GetOffset 24-bit calibration offset for channel 1
func (r *registerDfsdm_chcfg1r2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg1r2FieldOffsetMask) >> RegisterDfsdm_chcfg1r2FieldOffsetShift)
}

// SetOffset 24-bit calibration offset for channel 1
func (r *registerDfsdm_chcfg1r2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg1r2FieldOffsetMask)|(uint32(value)<<RegisterDfsdm_chcfg1r2FieldOffsetShift))
}

// registerDfsdm_chcfg2r2Type DFSDM channel configuration 2 register 2
type registerDfsdm_chcfg2r2Type uint32

const (
	RegisterDfsdm_chcfg2r2FieldDtrbsShift = 3
	RegisterDfsdm_chcfg2r2FieldDtrbsMask  = 0xf8
)

// GetDtrbs Data right bit-shift for channel 2
func (r *registerDfsdm_chcfg2r2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r2FieldDtrbsMask) >> RegisterDfsdm_chcfg2r2FieldDtrbsShift)
}

// SetDtrbs Data right bit-shift for channel 2
func (r *registerDfsdm_chcfg2r2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r2FieldDtrbsMask)|(uint32(value)<<RegisterDfsdm_chcfg2r2FieldDtrbsShift))
}

const (
	RegisterDfsdm_chcfg2r2FieldOffsetShift = 8
	RegisterDfsdm_chcfg2r2FieldOffsetMask  = 0xffffff00
)

// GetOffset 24-bit calibration offset for channel 2
func (r *registerDfsdm_chcfg2r2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg2r2FieldOffsetMask) >> RegisterDfsdm_chcfg2r2FieldOffsetShift)
}

// SetOffset 24-bit calibration offset for channel 2
func (r *registerDfsdm_chcfg2r2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg2r2FieldOffsetMask)|(uint32(value)<<RegisterDfsdm_chcfg2r2FieldOffsetShift))
}

// registerDfsdm_chcfg3r2Type DFSDM channel configuration 3 register 2
type registerDfsdm_chcfg3r2Type uint32

const (
	RegisterDfsdm_chcfg3r2FieldDtrbsShift = 3
	RegisterDfsdm_chcfg3r2FieldDtrbsMask  = 0xf8
)

// GetDtrbs Data right bit-shift for channel 3
func (r *registerDfsdm_chcfg3r2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r2FieldDtrbsMask) >> RegisterDfsdm_chcfg3r2FieldDtrbsShift)
}

// SetDtrbs Data right bit-shift for channel 3
func (r *registerDfsdm_chcfg3r2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r2FieldDtrbsMask)|(uint32(value)<<RegisterDfsdm_chcfg3r2FieldDtrbsShift))
}

const (
	RegisterDfsdm_chcfg3r2FieldOffsetShift = 8
	RegisterDfsdm_chcfg3r2FieldOffsetMask  = 0xffffff00
)

// GetOffset 24-bit calibration offset for channel 3
func (r *registerDfsdm_chcfg3r2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg3r2FieldOffsetMask) >> RegisterDfsdm_chcfg3r2FieldOffsetShift)
}

// SetOffset 24-bit calibration offset for channel 3
func (r *registerDfsdm_chcfg3r2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg3r2FieldOffsetMask)|(uint32(value)<<RegisterDfsdm_chcfg3r2FieldOffsetShift))
}

// registerDfsdm_chcfg4r2Type DFSDM channel configuration 4 register 2
type registerDfsdm_chcfg4r2Type uint32

const (
	RegisterDfsdm_chcfg4r2FieldDtrbsShift = 3
	RegisterDfsdm_chcfg4r2FieldDtrbsMask  = 0xf8
)

// GetDtrbs Data right bit-shift for channel 4
func (r *registerDfsdm_chcfg4r2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r2FieldDtrbsMask) >> RegisterDfsdm_chcfg4r2FieldDtrbsShift)
}

// SetDtrbs Data right bit-shift for channel 4
func (r *registerDfsdm_chcfg4r2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r2FieldDtrbsMask)|(uint32(value)<<RegisterDfsdm_chcfg4r2FieldDtrbsShift))
}

const (
	RegisterDfsdm_chcfg4r2FieldOffsetShift = 8
	RegisterDfsdm_chcfg4r2FieldOffsetMask  = 0xffffff00
)

// GetOffset 24-bit calibration offset for channel 4
func (r *registerDfsdm_chcfg4r2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg4r2FieldOffsetMask) >> RegisterDfsdm_chcfg4r2FieldOffsetShift)
}

// SetOffset 24-bit calibration offset for channel 4
func (r *registerDfsdm_chcfg4r2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg4r2FieldOffsetMask)|(uint32(value)<<RegisterDfsdm_chcfg4r2FieldOffsetShift))
}

// registerDfsdm_chcfg5r2Type DFSDM channel configuration 5 register 2
type registerDfsdm_chcfg5r2Type uint32

const (
	RegisterDfsdm_chcfg5r2FieldDtrbsShift = 3
	RegisterDfsdm_chcfg5r2FieldDtrbsMask  = 0xf8
)

// GetDtrbs Data right bit-shift for channel 5
func (r *registerDfsdm_chcfg5r2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r2FieldDtrbsMask) >> RegisterDfsdm_chcfg5r2FieldDtrbsShift)
}

// SetDtrbs Data right bit-shift for channel 5
func (r *registerDfsdm_chcfg5r2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r2FieldDtrbsMask)|(uint32(value)<<RegisterDfsdm_chcfg5r2FieldDtrbsShift))
}

const (
	RegisterDfsdm_chcfg5r2FieldOffsetShift = 8
	RegisterDfsdm_chcfg5r2FieldOffsetMask  = 0xffffff00
)

// GetOffset 24-bit calibration offset for channel 5
func (r *registerDfsdm_chcfg5r2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg5r2FieldOffsetMask) >> RegisterDfsdm_chcfg5r2FieldOffsetShift)
}

// SetOffset 24-bit calibration offset for channel 5
func (r *registerDfsdm_chcfg5r2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg5r2FieldOffsetMask)|(uint32(value)<<RegisterDfsdm_chcfg5r2FieldOffsetShift))
}

// registerDfsdm_chcfg6r2Type DFSDM channel configuration 6 register 2
type registerDfsdm_chcfg6r2Type uint32

const (
	RegisterDfsdm_chcfg6r2FieldDtrbsShift = 3
	RegisterDfsdm_chcfg6r2FieldDtrbsMask  = 0xf8
)

// GetDtrbs Data right bit-shift for channel 6
func (r *registerDfsdm_chcfg6r2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r2FieldDtrbsMask) >> RegisterDfsdm_chcfg6r2FieldDtrbsShift)
}

// SetDtrbs Data right bit-shift for channel 6
func (r *registerDfsdm_chcfg6r2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r2FieldDtrbsMask)|(uint32(value)<<RegisterDfsdm_chcfg6r2FieldDtrbsShift))
}

const (
	RegisterDfsdm_chcfg6r2FieldOffsetShift = 8
	RegisterDfsdm_chcfg6r2FieldOffsetMask  = 0xffffff00
)

// GetOffset 24-bit calibration offset for channel 6
func (r *registerDfsdm_chcfg6r2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg6r2FieldOffsetMask) >> RegisterDfsdm_chcfg6r2FieldOffsetShift)
}

// SetOffset 24-bit calibration offset for channel 6
func (r *registerDfsdm_chcfg6r2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg6r2FieldOffsetMask)|(uint32(value)<<RegisterDfsdm_chcfg6r2FieldOffsetShift))
}

// registerDfsdm_chcfg7r2Type DFSDM channel configuration 7 register 2
type registerDfsdm_chcfg7r2Type uint32

const (
	RegisterDfsdm_chcfg7r2FieldDtrbsShift = 3
	RegisterDfsdm_chcfg7r2FieldDtrbsMask  = 0xf8
)

// GetDtrbs Data right bit-shift for channel 7
func (r *registerDfsdm_chcfg7r2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r2FieldDtrbsMask) >> RegisterDfsdm_chcfg7r2FieldDtrbsShift)
}

// SetDtrbs Data right bit-shift for channel 7
func (r *registerDfsdm_chcfg7r2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r2FieldDtrbsMask)|(uint32(value)<<RegisterDfsdm_chcfg7r2FieldDtrbsShift))
}

const (
	RegisterDfsdm_chcfg7r2FieldOffsetShift = 8
	RegisterDfsdm_chcfg7r2FieldOffsetMask  = 0xffffff00
)

// GetOffset 24-bit calibration offset for channel 7
func (r *registerDfsdm_chcfg7r2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chcfg7r2FieldOffsetMask) >> RegisterDfsdm_chcfg7r2FieldOffsetShift)
}

// SetOffset 24-bit calibration offset for channel 7
func (r *registerDfsdm_chcfg7r2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chcfg7r2FieldOffsetMask)|(uint32(value)<<RegisterDfsdm_chcfg7r2FieldOffsetShift))
}

// registerDfsdm_awscd0rType DFSDM analog watchdog and short-circuit detector register
type registerDfsdm_awscd0rType uint32

const (
	RegisterDfsdm_awscd0rFieldScdtShift = 0
	RegisterDfsdm_awscd0rFieldScdtMask  = 0xff
)

// GetScdt short-circuit detector threshold for channel 0
func (r *registerDfsdm_awscd0rType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd0rFieldScdtMask) >> RegisterDfsdm_awscd0rFieldScdtShift)
}

// SetScdt short-circuit detector threshold for channel 0
func (r *registerDfsdm_awscd0rType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd0rFieldScdtMask)|(uint32(value)<<RegisterDfsdm_awscd0rFieldScdtShift))
}

const (
	RegisterDfsdm_awscd0rFieldBkscdShift = 12
	RegisterDfsdm_awscd0rFieldBkscdMask  = 0xf000
)

// GetBkscd Break signal assignment for short-circuit detector on channel 0
func (r *registerDfsdm_awscd0rType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd0rFieldBkscdMask) >> RegisterDfsdm_awscd0rFieldBkscdShift)
}

// SetBkscd Break signal assignment for short-circuit detector on channel 0
func (r *registerDfsdm_awscd0rType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd0rFieldBkscdMask)|(uint32(value)<<RegisterDfsdm_awscd0rFieldBkscdShift))
}

const (
	RegisterDfsdm_awscd0rFieldAwfosrShift = 16
	RegisterDfsdm_awscd0rFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 0
func (r *registerDfsdm_awscd0rType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd0rFieldAwfosrMask) >> RegisterDfsdm_awscd0rFieldAwfosrShift)
}

// SetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 0
func (r *registerDfsdm_awscd0rType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd0rFieldAwfosrMask)|(uint32(value)<<RegisterDfsdm_awscd0rFieldAwfosrShift))
}

const (
	RegisterDfsdm_awscd0rFieldAwfordShift = 22
	RegisterDfsdm_awscd0rFieldAwfordMask  = 0xc00000
)

// GetAwford Analog watchdog Sinc filter order on channel 0
func (r *registerDfsdm_awscd0rType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd0rFieldAwfordMask) >> RegisterDfsdm_awscd0rFieldAwfordShift)
}

// SetAwford Analog watchdog Sinc filter order on channel 0
func (r *registerDfsdm_awscd0rType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd0rFieldAwfordMask)|(uint32(value)<<RegisterDfsdm_awscd0rFieldAwfordShift))
}

// registerDfsdm_awscd1rType DFSDM analog watchdog and short-circuit detector register
type registerDfsdm_awscd1rType uint32

const (
	RegisterDfsdm_awscd1rFieldScdtShift = 0
	RegisterDfsdm_awscd1rFieldScdtMask  = 0xff
)

// GetScdt short-circuit detector threshold for channel 1
func (r *registerDfsdm_awscd1rType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd1rFieldScdtMask) >> RegisterDfsdm_awscd1rFieldScdtShift)
}

// SetScdt short-circuit detector threshold for channel 1
func (r *registerDfsdm_awscd1rType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd1rFieldScdtMask)|(uint32(value)<<RegisterDfsdm_awscd1rFieldScdtShift))
}

const (
	RegisterDfsdm_awscd1rFieldBkscdShift = 12
	RegisterDfsdm_awscd1rFieldBkscdMask  = 0xf000
)

// GetBkscd Break signal assignment for short-circuit detector on channel 1
func (r *registerDfsdm_awscd1rType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd1rFieldBkscdMask) >> RegisterDfsdm_awscd1rFieldBkscdShift)
}

// SetBkscd Break signal assignment for short-circuit detector on channel 1
func (r *registerDfsdm_awscd1rType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd1rFieldBkscdMask)|(uint32(value)<<RegisterDfsdm_awscd1rFieldBkscdShift))
}

const (
	RegisterDfsdm_awscd1rFieldAwfosrShift = 16
	RegisterDfsdm_awscd1rFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 1
func (r *registerDfsdm_awscd1rType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd1rFieldAwfosrMask) >> RegisterDfsdm_awscd1rFieldAwfosrShift)
}

// SetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 1
func (r *registerDfsdm_awscd1rType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd1rFieldAwfosrMask)|(uint32(value)<<RegisterDfsdm_awscd1rFieldAwfosrShift))
}

const (
	RegisterDfsdm_awscd1rFieldAwfordShift = 22
	RegisterDfsdm_awscd1rFieldAwfordMask  = 0xc00000
)

// GetAwford Analog watchdog Sinc filter order on channel 1
func (r *registerDfsdm_awscd1rType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd1rFieldAwfordMask) >> RegisterDfsdm_awscd1rFieldAwfordShift)
}

// SetAwford Analog watchdog Sinc filter order on channel 1
func (r *registerDfsdm_awscd1rType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd1rFieldAwfordMask)|(uint32(value)<<RegisterDfsdm_awscd1rFieldAwfordShift))
}

// registerDfsdm_awscd2rType DFSDM analog watchdog and short-circuit detector register
type registerDfsdm_awscd2rType uint32

const (
	RegisterDfsdm_awscd2rFieldScdtShift = 0
	RegisterDfsdm_awscd2rFieldScdtMask  = 0xff
)

// GetScdt short-circuit detector threshold for channel 2
func (r *registerDfsdm_awscd2rType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd2rFieldScdtMask) >> RegisterDfsdm_awscd2rFieldScdtShift)
}

// SetScdt short-circuit detector threshold for channel 2
func (r *registerDfsdm_awscd2rType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd2rFieldScdtMask)|(uint32(value)<<RegisterDfsdm_awscd2rFieldScdtShift))
}

const (
	RegisterDfsdm_awscd2rFieldBkscdShift = 12
	RegisterDfsdm_awscd2rFieldBkscdMask  = 0xf000
)

// GetBkscd Break signal assignment for short-circuit detector on channel 2
func (r *registerDfsdm_awscd2rType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd2rFieldBkscdMask) >> RegisterDfsdm_awscd2rFieldBkscdShift)
}

// SetBkscd Break signal assignment for short-circuit detector on channel 2
func (r *registerDfsdm_awscd2rType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd2rFieldBkscdMask)|(uint32(value)<<RegisterDfsdm_awscd2rFieldBkscdShift))
}

const (
	RegisterDfsdm_awscd2rFieldAwfosrShift = 16
	RegisterDfsdm_awscd2rFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 2
func (r *registerDfsdm_awscd2rType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd2rFieldAwfosrMask) >> RegisterDfsdm_awscd2rFieldAwfosrShift)
}

// SetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 2
func (r *registerDfsdm_awscd2rType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd2rFieldAwfosrMask)|(uint32(value)<<RegisterDfsdm_awscd2rFieldAwfosrShift))
}

const (
	RegisterDfsdm_awscd2rFieldAwfordShift = 22
	RegisterDfsdm_awscd2rFieldAwfordMask  = 0xc00000
)

// GetAwford Analog watchdog Sinc filter order on channel 2
func (r *registerDfsdm_awscd2rType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd2rFieldAwfordMask) >> RegisterDfsdm_awscd2rFieldAwfordShift)
}

// SetAwford Analog watchdog Sinc filter order on channel 2
func (r *registerDfsdm_awscd2rType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd2rFieldAwfordMask)|(uint32(value)<<RegisterDfsdm_awscd2rFieldAwfordShift))
}

// registerDfsdm_awscd3rType DFSDM analog watchdog and short-circuit detector register
type registerDfsdm_awscd3rType uint32

const (
	RegisterDfsdm_awscd3rFieldScdtShift = 0
	RegisterDfsdm_awscd3rFieldScdtMask  = 0xff
)

// GetScdt short-circuit detector threshold for channel 3
func (r *registerDfsdm_awscd3rType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd3rFieldScdtMask) >> RegisterDfsdm_awscd3rFieldScdtShift)
}

// SetScdt short-circuit detector threshold for channel 3
func (r *registerDfsdm_awscd3rType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd3rFieldScdtMask)|(uint32(value)<<RegisterDfsdm_awscd3rFieldScdtShift))
}

const (
	RegisterDfsdm_awscd3rFieldBkscdShift = 12
	RegisterDfsdm_awscd3rFieldBkscdMask  = 0xf000
)

// GetBkscd Break signal assignment for short-circuit detector on channel 3
func (r *registerDfsdm_awscd3rType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd3rFieldBkscdMask) >> RegisterDfsdm_awscd3rFieldBkscdShift)
}

// SetBkscd Break signal assignment for short-circuit detector on channel 3
func (r *registerDfsdm_awscd3rType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd3rFieldBkscdMask)|(uint32(value)<<RegisterDfsdm_awscd3rFieldBkscdShift))
}

const (
	RegisterDfsdm_awscd3rFieldAwfosrShift = 16
	RegisterDfsdm_awscd3rFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 3
func (r *registerDfsdm_awscd3rType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd3rFieldAwfosrMask) >> RegisterDfsdm_awscd3rFieldAwfosrShift)
}

// SetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 3
func (r *registerDfsdm_awscd3rType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd3rFieldAwfosrMask)|(uint32(value)<<RegisterDfsdm_awscd3rFieldAwfosrShift))
}

const (
	RegisterDfsdm_awscd3rFieldAwfordShift = 22
	RegisterDfsdm_awscd3rFieldAwfordMask  = 0xc00000
)

// GetAwford Analog watchdog Sinc filter order on channel 3
func (r *registerDfsdm_awscd3rType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd3rFieldAwfordMask) >> RegisterDfsdm_awscd3rFieldAwfordShift)
}

// SetAwford Analog watchdog Sinc filter order on channel 3
func (r *registerDfsdm_awscd3rType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd3rFieldAwfordMask)|(uint32(value)<<RegisterDfsdm_awscd3rFieldAwfordShift))
}

// registerDfsdm_awscd4rType DFSDM analog watchdog and short-circuit detector register
type registerDfsdm_awscd4rType uint32

const (
	RegisterDfsdm_awscd4rFieldScdtShift = 0
	RegisterDfsdm_awscd4rFieldScdtMask  = 0xff
)

// GetScdt short-circuit detector threshold for channel 4
func (r *registerDfsdm_awscd4rType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd4rFieldScdtMask) >> RegisterDfsdm_awscd4rFieldScdtShift)
}

// SetScdt short-circuit detector threshold for channel 4
func (r *registerDfsdm_awscd4rType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd4rFieldScdtMask)|(uint32(value)<<RegisterDfsdm_awscd4rFieldScdtShift))
}

const (
	RegisterDfsdm_awscd4rFieldBkscdShift = 12
	RegisterDfsdm_awscd4rFieldBkscdMask  = 0xf000
)

// GetBkscd Break signal assignment for short-circuit detector on channel 4
func (r *registerDfsdm_awscd4rType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd4rFieldBkscdMask) >> RegisterDfsdm_awscd4rFieldBkscdShift)
}

// SetBkscd Break signal assignment for short-circuit detector on channel 4
func (r *registerDfsdm_awscd4rType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd4rFieldBkscdMask)|(uint32(value)<<RegisterDfsdm_awscd4rFieldBkscdShift))
}

const (
	RegisterDfsdm_awscd4rFieldAwfosrShift = 16
	RegisterDfsdm_awscd4rFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 4
func (r *registerDfsdm_awscd4rType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd4rFieldAwfosrMask) >> RegisterDfsdm_awscd4rFieldAwfosrShift)
}

// SetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 4
func (r *registerDfsdm_awscd4rType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd4rFieldAwfosrMask)|(uint32(value)<<RegisterDfsdm_awscd4rFieldAwfosrShift))
}

const (
	RegisterDfsdm_awscd4rFieldAwfordShift = 22
	RegisterDfsdm_awscd4rFieldAwfordMask  = 0xc00000
)

// GetAwford Analog watchdog Sinc filter order on channel 4
func (r *registerDfsdm_awscd4rType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd4rFieldAwfordMask) >> RegisterDfsdm_awscd4rFieldAwfordShift)
}

// SetAwford Analog watchdog Sinc filter order on channel 4
func (r *registerDfsdm_awscd4rType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd4rFieldAwfordMask)|(uint32(value)<<RegisterDfsdm_awscd4rFieldAwfordShift))
}

// registerDfsdm_awscd5rType DFSDM analog watchdog and short-circuit detector register
type registerDfsdm_awscd5rType uint32

const (
	RegisterDfsdm_awscd5rFieldScdtShift = 0
	RegisterDfsdm_awscd5rFieldScdtMask  = 0xff
)

// GetScdt short-circuit detector threshold for channel 5
func (r *registerDfsdm_awscd5rType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd5rFieldScdtMask) >> RegisterDfsdm_awscd5rFieldScdtShift)
}

// SetScdt short-circuit detector threshold for channel 5
func (r *registerDfsdm_awscd5rType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd5rFieldScdtMask)|(uint32(value)<<RegisterDfsdm_awscd5rFieldScdtShift))
}

const (
	RegisterDfsdm_awscd5rFieldBkscdShift = 12
	RegisterDfsdm_awscd5rFieldBkscdMask  = 0xf000
)

// GetBkscd Break signal assignment for short-circuit detector on channel 5
func (r *registerDfsdm_awscd5rType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd5rFieldBkscdMask) >> RegisterDfsdm_awscd5rFieldBkscdShift)
}

// SetBkscd Break signal assignment for short-circuit detector on channel 5
func (r *registerDfsdm_awscd5rType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd5rFieldBkscdMask)|(uint32(value)<<RegisterDfsdm_awscd5rFieldBkscdShift))
}

const (
	RegisterDfsdm_awscd5rFieldAwfosrShift = 16
	RegisterDfsdm_awscd5rFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 5
func (r *registerDfsdm_awscd5rType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd5rFieldAwfosrMask) >> RegisterDfsdm_awscd5rFieldAwfosrShift)
}

// SetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 5
func (r *registerDfsdm_awscd5rType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd5rFieldAwfosrMask)|(uint32(value)<<RegisterDfsdm_awscd5rFieldAwfosrShift))
}

const (
	RegisterDfsdm_awscd5rFieldAwfordShift = 22
	RegisterDfsdm_awscd5rFieldAwfordMask  = 0xc00000
)

// GetAwford Analog watchdog Sinc filter order on channel 5
func (r *registerDfsdm_awscd5rType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd5rFieldAwfordMask) >> RegisterDfsdm_awscd5rFieldAwfordShift)
}

// SetAwford Analog watchdog Sinc filter order on channel 5
func (r *registerDfsdm_awscd5rType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd5rFieldAwfordMask)|(uint32(value)<<RegisterDfsdm_awscd5rFieldAwfordShift))
}

// registerDfsdm_awscd6rType DFSDM analog watchdog and short-circuit detector register
type registerDfsdm_awscd6rType uint32

const (
	RegisterDfsdm_awscd6rFieldScdtShift = 0
	RegisterDfsdm_awscd6rFieldScdtMask  = 0xff
)

// GetScdt short-circuit detector threshold for channel 6
func (r *registerDfsdm_awscd6rType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd6rFieldScdtMask) >> RegisterDfsdm_awscd6rFieldScdtShift)
}

// SetScdt short-circuit detector threshold for channel 6
func (r *registerDfsdm_awscd6rType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd6rFieldScdtMask)|(uint32(value)<<RegisterDfsdm_awscd6rFieldScdtShift))
}

const (
	RegisterDfsdm_awscd6rFieldBkscdShift = 12
	RegisterDfsdm_awscd6rFieldBkscdMask  = 0xf000
)

// GetBkscd Break signal assignment for short-circuit detector on channel 6
func (r *registerDfsdm_awscd6rType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd6rFieldBkscdMask) >> RegisterDfsdm_awscd6rFieldBkscdShift)
}

// SetBkscd Break signal assignment for short-circuit detector on channel 6
func (r *registerDfsdm_awscd6rType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd6rFieldBkscdMask)|(uint32(value)<<RegisterDfsdm_awscd6rFieldBkscdShift))
}

const (
	RegisterDfsdm_awscd6rFieldAwfosrShift = 16
	RegisterDfsdm_awscd6rFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 6
func (r *registerDfsdm_awscd6rType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd6rFieldAwfosrMask) >> RegisterDfsdm_awscd6rFieldAwfosrShift)
}

// SetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 6
func (r *registerDfsdm_awscd6rType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd6rFieldAwfosrMask)|(uint32(value)<<RegisterDfsdm_awscd6rFieldAwfosrShift))
}

const (
	RegisterDfsdm_awscd6rFieldAwfordShift = 22
	RegisterDfsdm_awscd6rFieldAwfordMask  = 0xc00000
)

// GetAwford Analog watchdog Sinc filter order on channel 6
func (r *registerDfsdm_awscd6rType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd6rFieldAwfordMask) >> RegisterDfsdm_awscd6rFieldAwfordShift)
}

// SetAwford Analog watchdog Sinc filter order on channel 6
func (r *registerDfsdm_awscd6rType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd6rFieldAwfordMask)|(uint32(value)<<RegisterDfsdm_awscd6rFieldAwfordShift))
}

// registerDfsdm_awscd7rType DFSDM analog watchdog and short-circuit detector register
type registerDfsdm_awscd7rType uint32

const (
	RegisterDfsdm_awscd7rFieldScdtShift = 0
	RegisterDfsdm_awscd7rFieldScdtMask  = 0xff
)

// GetScdt short-circuit detector threshold for channel 7
func (r *registerDfsdm_awscd7rType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd7rFieldScdtMask) >> RegisterDfsdm_awscd7rFieldScdtShift)
}

// SetScdt short-circuit detector threshold for channel 7
func (r *registerDfsdm_awscd7rType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd7rFieldScdtMask)|(uint32(value)<<RegisterDfsdm_awscd7rFieldScdtShift))
}

const (
	RegisterDfsdm_awscd7rFieldBkscdShift = 12
	RegisterDfsdm_awscd7rFieldBkscdMask  = 0xf000
)

// GetBkscd Break signal assignment for short-circuit detector on channel 7
func (r *registerDfsdm_awscd7rType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd7rFieldBkscdMask) >> RegisterDfsdm_awscd7rFieldBkscdShift)
}

// SetBkscd Break signal assignment for short-circuit detector on channel 7
func (r *registerDfsdm_awscd7rType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd7rFieldBkscdMask)|(uint32(value)<<RegisterDfsdm_awscd7rFieldBkscdShift))
}

const (
	RegisterDfsdm_awscd7rFieldAwfosrShift = 16
	RegisterDfsdm_awscd7rFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 7
func (r *registerDfsdm_awscd7rType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd7rFieldAwfosrMask) >> RegisterDfsdm_awscd7rFieldAwfosrShift)
}

// SetAwfosr Analog watchdog filter oversampling ratio (decimation rate) on channel 7
func (r *registerDfsdm_awscd7rType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd7rFieldAwfosrMask)|(uint32(value)<<RegisterDfsdm_awscd7rFieldAwfosrShift))
}

const (
	RegisterDfsdm_awscd7rFieldAwfordShift = 22
	RegisterDfsdm_awscd7rFieldAwfordMask  = 0xc00000
)

// GetAwford Analog watchdog Sinc filter order on channel 7
func (r *registerDfsdm_awscd7rType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_awscd7rFieldAwfordMask) >> RegisterDfsdm_awscd7rFieldAwfordShift)
}

// SetAwford Analog watchdog Sinc filter order on channel 7
func (r *registerDfsdm_awscd7rType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_awscd7rFieldAwfordMask)|(uint32(value)<<RegisterDfsdm_awscd7rFieldAwfordShift))
}

// registerDfsdm_chwdat0rType DFSDM channel watchdog filter data register
type registerDfsdm_chwdat0rType uint32

const (
	RegisterDfsdm_chwdat0rFieldWdataShift = 0
	RegisterDfsdm_chwdat0rFieldWdataMask  = 0xffff
)

// GetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat0rType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chwdat0rFieldWdataMask) >> RegisterDfsdm_chwdat0rFieldWdataShift)
}

// SetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat0rType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chwdat0rFieldWdataMask)|(uint32(value)<<RegisterDfsdm_chwdat0rFieldWdataShift))
}

// registerDfsdm_chwdat1rType DFSDM channel watchdog filter data register
type registerDfsdm_chwdat1rType uint32

const (
	RegisterDfsdm_chwdat1rFieldWdataShift = 0
	RegisterDfsdm_chwdat1rFieldWdataMask  = 0xffff
)

// GetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat1rType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chwdat1rFieldWdataMask) >> RegisterDfsdm_chwdat1rFieldWdataShift)
}

// SetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat1rType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chwdat1rFieldWdataMask)|(uint32(value)<<RegisterDfsdm_chwdat1rFieldWdataShift))
}

// registerDfsdm_chwdat2rType DFSDM channel watchdog filter data register
type registerDfsdm_chwdat2rType uint32

const (
	RegisterDfsdm_chwdat2rFieldWdataShift = 0
	RegisterDfsdm_chwdat2rFieldWdataMask  = 0xffff
)

// GetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat2rType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chwdat2rFieldWdataMask) >> RegisterDfsdm_chwdat2rFieldWdataShift)
}

// SetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat2rType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chwdat2rFieldWdataMask)|(uint32(value)<<RegisterDfsdm_chwdat2rFieldWdataShift))
}

// registerDfsdm_chwdat3rType DFSDM channel watchdog filter data register
type registerDfsdm_chwdat3rType uint32

const (
	RegisterDfsdm_chwdat3rFieldWdataShift = 0
	RegisterDfsdm_chwdat3rFieldWdataMask  = 0xffff
)

// GetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat3rType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chwdat3rFieldWdataMask) >> RegisterDfsdm_chwdat3rFieldWdataShift)
}

// SetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat3rType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chwdat3rFieldWdataMask)|(uint32(value)<<RegisterDfsdm_chwdat3rFieldWdataShift))
}

// registerDfsdm_chwdat4rType DFSDM channel watchdog filter data register
type registerDfsdm_chwdat4rType uint32

const (
	RegisterDfsdm_chwdat4rFieldWdataShift = 0
	RegisterDfsdm_chwdat4rFieldWdataMask  = 0xffff
)

// GetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat4rType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chwdat4rFieldWdataMask) >> RegisterDfsdm_chwdat4rFieldWdataShift)
}

// SetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat4rType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chwdat4rFieldWdataMask)|(uint32(value)<<RegisterDfsdm_chwdat4rFieldWdataShift))
}

// registerDfsdm_chwdat5rType DFSDM channel watchdog filter data register
type registerDfsdm_chwdat5rType uint32

const (
	RegisterDfsdm_chwdat5rFieldWdataShift = 0
	RegisterDfsdm_chwdat5rFieldWdataMask  = 0xffff
)

// GetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat5rType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chwdat5rFieldWdataMask) >> RegisterDfsdm_chwdat5rFieldWdataShift)
}

// SetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat5rType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chwdat5rFieldWdataMask)|(uint32(value)<<RegisterDfsdm_chwdat5rFieldWdataShift))
}

// registerDfsdm_chwdat6rType DFSDM channel watchdog filter data register
type registerDfsdm_chwdat6rType uint32

const (
	RegisterDfsdm_chwdat6rFieldWdataShift = 0
	RegisterDfsdm_chwdat6rFieldWdataMask  = 0xffff
)

// GetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat6rType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chwdat6rFieldWdataMask) >> RegisterDfsdm_chwdat6rFieldWdataShift)
}

// SetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat6rType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chwdat6rFieldWdataMask)|(uint32(value)<<RegisterDfsdm_chwdat6rFieldWdataShift))
}

// registerDfsdm_chwdat7rType DFSDM channel watchdog filter data register
type registerDfsdm_chwdat7rType uint32

const (
	RegisterDfsdm_chwdat7rFieldWdataShift = 0
	RegisterDfsdm_chwdat7rFieldWdataMask  = 0xffff
)

// GetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat7rType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chwdat7rFieldWdataMask) >> RegisterDfsdm_chwdat7rFieldWdataShift)
}

// SetWdata Input channel y watchdog data
func (r *registerDfsdm_chwdat7rType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chwdat7rFieldWdataMask)|(uint32(value)<<RegisterDfsdm_chwdat7rFieldWdataShift))
}

// registerDfsdm_chdatin0rType DFSDM channel data input register
type registerDfsdm_chdatin0rType uint32

const (
	RegisterDfsdm_chdatin0rFieldIndat0Shift = 0
	RegisterDfsdm_chdatin0rFieldIndat0Mask  = 0xffff
)

// GetIndat0 Input data for channel 0
func (r *registerDfsdm_chdatin0rType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin0rFieldIndat0Mask) >> RegisterDfsdm_chdatin0rFieldIndat0Shift)
}

// SetIndat0 Input data for channel 0
func (r *registerDfsdm_chdatin0rType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin0rFieldIndat0Mask)|(uint32(value)<<RegisterDfsdm_chdatin0rFieldIndat0Shift))
}

const (
	RegisterDfsdm_chdatin0rFieldIndat1Shift = 16
	RegisterDfsdm_chdatin0rFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 Input data for channel 1
func (r *registerDfsdm_chdatin0rType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin0rFieldIndat1Mask) >> RegisterDfsdm_chdatin0rFieldIndat1Shift)
}

// SetIndat1 Input data for channel 1
func (r *registerDfsdm_chdatin0rType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin0rFieldIndat1Mask)|(uint32(value)<<RegisterDfsdm_chdatin0rFieldIndat1Shift))
}

// registerDfsdm_chdatin1rType DFSDM channel data input register
type registerDfsdm_chdatin1rType uint32

const (
	RegisterDfsdm_chdatin1rFieldIndat0Shift = 0
	RegisterDfsdm_chdatin1rFieldIndat0Mask  = 0xffff
)

// GetIndat0 Input data for channel 1
func (r *registerDfsdm_chdatin1rType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin1rFieldIndat0Mask) >> RegisterDfsdm_chdatin1rFieldIndat0Shift)
}

// SetIndat0 Input data for channel 1
func (r *registerDfsdm_chdatin1rType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin1rFieldIndat0Mask)|(uint32(value)<<RegisterDfsdm_chdatin1rFieldIndat0Shift))
}

const (
	RegisterDfsdm_chdatin1rFieldIndat1Shift = 16
	RegisterDfsdm_chdatin1rFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 Input data for channel 2
func (r *registerDfsdm_chdatin1rType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin1rFieldIndat1Mask) >> RegisterDfsdm_chdatin1rFieldIndat1Shift)
}

// SetIndat1 Input data for channel 2
func (r *registerDfsdm_chdatin1rType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin1rFieldIndat1Mask)|(uint32(value)<<RegisterDfsdm_chdatin1rFieldIndat1Shift))
}

// registerDfsdm_chdatin2rType DFSDM channel data input register
type registerDfsdm_chdatin2rType uint32

const (
	RegisterDfsdm_chdatin2rFieldIndat0Shift = 0
	RegisterDfsdm_chdatin2rFieldIndat0Mask  = 0xffff
)

// GetIndat0 Input data for channel 2
func (r *registerDfsdm_chdatin2rType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin2rFieldIndat0Mask) >> RegisterDfsdm_chdatin2rFieldIndat0Shift)
}

// SetIndat0 Input data for channel 2
func (r *registerDfsdm_chdatin2rType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin2rFieldIndat0Mask)|(uint32(value)<<RegisterDfsdm_chdatin2rFieldIndat0Shift))
}

const (
	RegisterDfsdm_chdatin2rFieldIndat1Shift = 16
	RegisterDfsdm_chdatin2rFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 Input data for channel 3
func (r *registerDfsdm_chdatin2rType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin2rFieldIndat1Mask) >> RegisterDfsdm_chdatin2rFieldIndat1Shift)
}

// SetIndat1 Input data for channel 3
func (r *registerDfsdm_chdatin2rType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin2rFieldIndat1Mask)|(uint32(value)<<RegisterDfsdm_chdatin2rFieldIndat1Shift))
}

// registerDfsdm_chdatin3rType DFSDM channel data input register
type registerDfsdm_chdatin3rType uint32

const (
	RegisterDfsdm_chdatin3rFieldIndat0Shift = 0
	RegisterDfsdm_chdatin3rFieldIndat0Mask  = 0xffff
)

// GetIndat0 Input data for channel 3
func (r *registerDfsdm_chdatin3rType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin3rFieldIndat0Mask) >> RegisterDfsdm_chdatin3rFieldIndat0Shift)
}

// SetIndat0 Input data for channel 3
func (r *registerDfsdm_chdatin3rType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin3rFieldIndat0Mask)|(uint32(value)<<RegisterDfsdm_chdatin3rFieldIndat0Shift))
}

const (
	RegisterDfsdm_chdatin3rFieldIndat1Shift = 16
	RegisterDfsdm_chdatin3rFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 Input data for channel 4
func (r *registerDfsdm_chdatin3rType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin3rFieldIndat1Mask) >> RegisterDfsdm_chdatin3rFieldIndat1Shift)
}

// SetIndat1 Input data for channel 4
func (r *registerDfsdm_chdatin3rType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin3rFieldIndat1Mask)|(uint32(value)<<RegisterDfsdm_chdatin3rFieldIndat1Shift))
}

// registerDfsdm_chdatin4rType DFSDM channel data input register
type registerDfsdm_chdatin4rType uint32

const (
	RegisterDfsdm_chdatin4rFieldIndat0Shift = 0
	RegisterDfsdm_chdatin4rFieldIndat0Mask  = 0xffff
)

// GetIndat0 Input data for channel 4
func (r *registerDfsdm_chdatin4rType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin4rFieldIndat0Mask) >> RegisterDfsdm_chdatin4rFieldIndat0Shift)
}

// SetIndat0 Input data for channel 4
func (r *registerDfsdm_chdatin4rType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin4rFieldIndat0Mask)|(uint32(value)<<RegisterDfsdm_chdatin4rFieldIndat0Shift))
}

const (
	RegisterDfsdm_chdatin4rFieldIndat1Shift = 16
	RegisterDfsdm_chdatin4rFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 Input data for channel 5
func (r *registerDfsdm_chdatin4rType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin4rFieldIndat1Mask) >> RegisterDfsdm_chdatin4rFieldIndat1Shift)
}

// SetIndat1 Input data for channel 5
func (r *registerDfsdm_chdatin4rType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin4rFieldIndat1Mask)|(uint32(value)<<RegisterDfsdm_chdatin4rFieldIndat1Shift))
}

// registerDfsdm_chdatin5rType DFSDM channel data input register
type registerDfsdm_chdatin5rType uint32

const (
	RegisterDfsdm_chdatin5rFieldIndat0Shift = 0
	RegisterDfsdm_chdatin5rFieldIndat0Mask  = 0xffff
)

// GetIndat0 Input data for channel 5
func (r *registerDfsdm_chdatin5rType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin5rFieldIndat0Mask) >> RegisterDfsdm_chdatin5rFieldIndat0Shift)
}

// SetIndat0 Input data for channel 5
func (r *registerDfsdm_chdatin5rType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin5rFieldIndat0Mask)|(uint32(value)<<RegisterDfsdm_chdatin5rFieldIndat0Shift))
}

const (
	RegisterDfsdm_chdatin5rFieldIndat1Shift = 16
	RegisterDfsdm_chdatin5rFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 Input data for channel 6
func (r *registerDfsdm_chdatin5rType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin5rFieldIndat1Mask) >> RegisterDfsdm_chdatin5rFieldIndat1Shift)
}

// SetIndat1 Input data for channel 6
func (r *registerDfsdm_chdatin5rType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin5rFieldIndat1Mask)|(uint32(value)<<RegisterDfsdm_chdatin5rFieldIndat1Shift))
}

// registerDfsdm_chdatin6rType DFSDM channel data input register
type registerDfsdm_chdatin6rType uint32

const (
	RegisterDfsdm_chdatin6rFieldIndat0Shift = 0
	RegisterDfsdm_chdatin6rFieldIndat0Mask  = 0xffff
)

// GetIndat0 Input data for channel 6
func (r *registerDfsdm_chdatin6rType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin6rFieldIndat0Mask) >> RegisterDfsdm_chdatin6rFieldIndat0Shift)
}

// SetIndat0 Input data for channel 6
func (r *registerDfsdm_chdatin6rType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin6rFieldIndat0Mask)|(uint32(value)<<RegisterDfsdm_chdatin6rFieldIndat0Shift))
}

const (
	RegisterDfsdm_chdatin6rFieldIndat1Shift = 16
	RegisterDfsdm_chdatin6rFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 Input data for channel 7
func (r *registerDfsdm_chdatin6rType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin6rFieldIndat1Mask) >> RegisterDfsdm_chdatin6rFieldIndat1Shift)
}

// SetIndat1 Input data for channel 7
func (r *registerDfsdm_chdatin6rType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin6rFieldIndat1Mask)|(uint32(value)<<RegisterDfsdm_chdatin6rFieldIndat1Shift))
}

// registerDfsdm_chdatin7rType DFSDM channel data input register
type registerDfsdm_chdatin7rType uint32

const (
	RegisterDfsdm_chdatin7rFieldIndat0Shift = 0
	RegisterDfsdm_chdatin7rFieldIndat0Mask  = 0xffff
)

// GetIndat0 Input data for channel 7
func (r *registerDfsdm_chdatin7rType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin7rFieldIndat0Mask) >> RegisterDfsdm_chdatin7rFieldIndat0Shift)
}

// SetIndat0 Input data for channel 7
func (r *registerDfsdm_chdatin7rType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin7rFieldIndat0Mask)|(uint32(value)<<RegisterDfsdm_chdatin7rFieldIndat0Shift))
}

const (
	RegisterDfsdm_chdatin7rFieldIndat1Shift = 16
	RegisterDfsdm_chdatin7rFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 Input data for channel 8
func (r *registerDfsdm_chdatin7rType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_chdatin7rFieldIndat1Mask) >> RegisterDfsdm_chdatin7rFieldIndat1Shift)
}

// SetIndat1 Input data for channel 8
func (r *registerDfsdm_chdatin7rType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_chdatin7rFieldIndat1Mask)|(uint32(value)<<RegisterDfsdm_chdatin7rFieldIndat1Shift))
}

// registerDfsdm0_cr1Type DFSDM control register 1
type registerDfsdm0_cr1Type uint32

const (
	RegisterDfsdm0_cr1FieldDfenShift = 0
	RegisterDfsdm0_cr1FieldDfenMask  = 0x1
)

// GetDfen DFSDM enable
func (r *registerDfsdm0_cr1Type) GetDfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldDfenMask) != 0
}

// SetDfen DFSDM enable
func (r *registerDfsdm0_cr1Type) SetDfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr1FieldDfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldDfenMask)
	}
}

const (
	RegisterDfsdm0_cr1FieldJswstartShift = 1
	RegisterDfsdm0_cr1FieldJswstartMask  = 0x2
)

// GetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm0_cr1Type) GetJswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldJswstartMask) != 0
}

// SetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm0_cr1Type) SetJswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr1FieldJswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldJswstartMask)
	}
}

const (
	RegisterDfsdm0_cr1FieldJsyncShift = 3
	RegisterDfsdm0_cr1FieldJsyncMask  = 0x8
)

// GetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm0_cr1Type) GetJsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldJsyncMask) != 0
}

// SetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm0_cr1Type) SetJsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr1FieldJsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldJsyncMask)
	}
}

const (
	RegisterDfsdm0_cr1FieldJscanShift = 4
	RegisterDfsdm0_cr1FieldJscanMask  = 0x10
)

// GetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm0_cr1Type) GetJscan() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldJscanMask) != 0
}

// SetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm0_cr1Type) SetJscan(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr1FieldJscanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldJscanMask)
	}
}

const (
	RegisterDfsdm0_cr1FieldJdmaenShift = 5
	RegisterDfsdm0_cr1FieldJdmaenMask  = 0x20
)

// GetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm0_cr1Type) GetJdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldJdmaenMask) != 0
}

// SetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm0_cr1Type) SetJdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr1FieldJdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldJdmaenMask)
	}
}

const (
	RegisterDfsdm0_cr1FieldJextselShift = 8
	RegisterDfsdm0_cr1FieldJextselMask  = 0x1f00
)

// GetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm0_cr1Type) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldJextselMask) >> RegisterDfsdm0_cr1FieldJextselShift)
}

// SetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm0_cr1Type) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldJextselMask)|(uint32(value)<<RegisterDfsdm0_cr1FieldJextselShift))
}

const (
	RegisterDfsdm0_cr1FieldJextenShift = 13
	RegisterDfsdm0_cr1FieldJextenMask  = 0x6000
)

// GetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm0_cr1Type) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldJextenMask) >> RegisterDfsdm0_cr1FieldJextenShift)
}

// SetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm0_cr1Type) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldJextenMask)|(uint32(value)<<RegisterDfsdm0_cr1FieldJextenShift))
}

const (
	RegisterDfsdm0_cr1FieldRswstartShift = 17
	RegisterDfsdm0_cr1FieldRswstartMask  = 0x20000
)

// GetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm0_cr1Type) GetRswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldRswstartMask) != 0
}

// SetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm0_cr1Type) SetRswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr1FieldRswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldRswstartMask)
	}
}

const (
	RegisterDfsdm0_cr1FieldRcontShift = 18
	RegisterDfsdm0_cr1FieldRcontMask  = 0x40000
)

// GetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm0_cr1Type) GetRcont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldRcontMask) != 0
}

// SetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm0_cr1Type) SetRcont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr1FieldRcontMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldRcontMask)
	}
}

const (
	RegisterDfsdm0_cr1FieldRsyncShift = 19
	RegisterDfsdm0_cr1FieldRsyncMask  = 0x80000
)

// GetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm0_cr1Type) GetRsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldRsyncMask) != 0
}

// SetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm0_cr1Type) SetRsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr1FieldRsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldRsyncMask)
	}
}

const (
	RegisterDfsdm0_cr1FieldRdmaenShift = 21
	RegisterDfsdm0_cr1FieldRdmaenMask  = 0x200000
)

// GetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm0_cr1Type) GetRdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldRdmaenMask) != 0
}

// SetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm0_cr1Type) SetRdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr1FieldRdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldRdmaenMask)
	}
}

const (
	RegisterDfsdm0_cr1FieldRchShift = 24
	RegisterDfsdm0_cr1FieldRchMask  = 0x7000000
)

// GetRch Regular channel selection
func (r *registerDfsdm0_cr1Type) GetRch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldRchMask) >> RegisterDfsdm0_cr1FieldRchShift)
}

// SetRch Regular channel selection
func (r *registerDfsdm0_cr1Type) SetRch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldRchMask)|(uint32(value)<<RegisterDfsdm0_cr1FieldRchShift))
}

const (
	RegisterDfsdm0_cr1FieldFastShift = 29
	RegisterDfsdm0_cr1FieldFastMask  = 0x20000000
)

// GetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm0_cr1Type) GetFast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldFastMask) != 0
}

// SetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm0_cr1Type) SetFast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr1FieldFastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldFastMask)
	}
}

const (
	RegisterDfsdm0_cr1FieldAwfselShift = 30
	RegisterDfsdm0_cr1FieldAwfselMask  = 0x40000000
)

// GetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm0_cr1Type) GetAwfsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr1FieldAwfselMask) != 0
}

// SetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm0_cr1Type) SetAwfsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr1FieldAwfselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr1FieldAwfselMask)
	}
}

// registerDfsdm1_cr1Type DFSDM control register 1
type registerDfsdm1_cr1Type uint32

const (
	RegisterDfsdm1_cr1FieldDfenShift = 0
	RegisterDfsdm1_cr1FieldDfenMask  = 0x1
)

// GetDfen DFSDM enable
func (r *registerDfsdm1_cr1Type) GetDfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldDfenMask) != 0
}

// SetDfen DFSDM enable
func (r *registerDfsdm1_cr1Type) SetDfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr1FieldDfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldDfenMask)
	}
}

const (
	RegisterDfsdm1_cr1FieldJswstartShift = 1
	RegisterDfsdm1_cr1FieldJswstartMask  = 0x2
)

// GetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm1_cr1Type) GetJswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldJswstartMask) != 0
}

// SetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm1_cr1Type) SetJswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr1FieldJswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldJswstartMask)
	}
}

const (
	RegisterDfsdm1_cr1FieldJsyncShift = 3
	RegisterDfsdm1_cr1FieldJsyncMask  = 0x8
)

// GetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm1_cr1Type) GetJsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldJsyncMask) != 0
}

// SetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm1_cr1Type) SetJsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr1FieldJsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldJsyncMask)
	}
}

const (
	RegisterDfsdm1_cr1FieldJscanShift = 4
	RegisterDfsdm1_cr1FieldJscanMask  = 0x10
)

// GetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm1_cr1Type) GetJscan() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldJscanMask) != 0
}

// SetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm1_cr1Type) SetJscan(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr1FieldJscanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldJscanMask)
	}
}

const (
	RegisterDfsdm1_cr1FieldJdmaenShift = 5
	RegisterDfsdm1_cr1FieldJdmaenMask  = 0x20
)

// GetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm1_cr1Type) GetJdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldJdmaenMask) != 0
}

// SetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm1_cr1Type) SetJdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr1FieldJdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldJdmaenMask)
	}
}

const (
	RegisterDfsdm1_cr1FieldJextselShift = 8
	RegisterDfsdm1_cr1FieldJextselMask  = 0x1f00
)

// GetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm1_cr1Type) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldJextselMask) >> RegisterDfsdm1_cr1FieldJextselShift)
}

// SetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm1_cr1Type) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldJextselMask)|(uint32(value)<<RegisterDfsdm1_cr1FieldJextselShift))
}

const (
	RegisterDfsdm1_cr1FieldJextenShift = 13
	RegisterDfsdm1_cr1FieldJextenMask  = 0x6000
)

// GetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm1_cr1Type) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldJextenMask) >> RegisterDfsdm1_cr1FieldJextenShift)
}

// SetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm1_cr1Type) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldJextenMask)|(uint32(value)<<RegisterDfsdm1_cr1FieldJextenShift))
}

const (
	RegisterDfsdm1_cr1FieldRswstartShift = 17
	RegisterDfsdm1_cr1FieldRswstartMask  = 0x20000
)

// GetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm1_cr1Type) GetRswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldRswstartMask) != 0
}

// SetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm1_cr1Type) SetRswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr1FieldRswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldRswstartMask)
	}
}

const (
	RegisterDfsdm1_cr1FieldRcontShift = 18
	RegisterDfsdm1_cr1FieldRcontMask  = 0x40000
)

// GetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm1_cr1Type) GetRcont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldRcontMask) != 0
}

// SetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm1_cr1Type) SetRcont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr1FieldRcontMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldRcontMask)
	}
}

const (
	RegisterDfsdm1_cr1FieldRsyncShift = 19
	RegisterDfsdm1_cr1FieldRsyncMask  = 0x80000
)

// GetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm1_cr1Type) GetRsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldRsyncMask) != 0
}

// SetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm1_cr1Type) SetRsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr1FieldRsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldRsyncMask)
	}
}

const (
	RegisterDfsdm1_cr1FieldRdmaenShift = 21
	RegisterDfsdm1_cr1FieldRdmaenMask  = 0x200000
)

// GetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm1_cr1Type) GetRdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldRdmaenMask) != 0
}

// SetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm1_cr1Type) SetRdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr1FieldRdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldRdmaenMask)
	}
}

const (
	RegisterDfsdm1_cr1FieldRchShift = 24
	RegisterDfsdm1_cr1FieldRchMask  = 0x7000000
)

// GetRch Regular channel selection
func (r *registerDfsdm1_cr1Type) GetRch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldRchMask) >> RegisterDfsdm1_cr1FieldRchShift)
}

// SetRch Regular channel selection
func (r *registerDfsdm1_cr1Type) SetRch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldRchMask)|(uint32(value)<<RegisterDfsdm1_cr1FieldRchShift))
}

const (
	RegisterDfsdm1_cr1FieldFastShift = 29
	RegisterDfsdm1_cr1FieldFastMask  = 0x20000000
)

// GetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm1_cr1Type) GetFast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldFastMask) != 0
}

// SetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm1_cr1Type) SetFast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr1FieldFastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldFastMask)
	}
}

const (
	RegisterDfsdm1_cr1FieldAwfselShift = 30
	RegisterDfsdm1_cr1FieldAwfselMask  = 0x40000000
)

// GetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm1_cr1Type) GetAwfsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr1FieldAwfselMask) != 0
}

// SetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm1_cr1Type) SetAwfsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr1FieldAwfselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr1FieldAwfselMask)
	}
}

// registerDfsdm2_cr1Type DFSDM control register 1
type registerDfsdm2_cr1Type uint32

const (
	RegisterDfsdm2_cr1FieldDfenShift = 0
	RegisterDfsdm2_cr1FieldDfenMask  = 0x1
)

// GetDfen DFSDM enable
func (r *registerDfsdm2_cr1Type) GetDfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldDfenMask) != 0
}

// SetDfen DFSDM enable
func (r *registerDfsdm2_cr1Type) SetDfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr1FieldDfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldDfenMask)
	}
}

const (
	RegisterDfsdm2_cr1FieldJswstartShift = 1
	RegisterDfsdm2_cr1FieldJswstartMask  = 0x2
)

// GetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm2_cr1Type) GetJswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldJswstartMask) != 0
}

// SetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm2_cr1Type) SetJswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr1FieldJswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldJswstartMask)
	}
}

const (
	RegisterDfsdm2_cr1FieldJsyncShift = 3
	RegisterDfsdm2_cr1FieldJsyncMask  = 0x8
)

// GetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm2_cr1Type) GetJsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldJsyncMask) != 0
}

// SetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm2_cr1Type) SetJsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr1FieldJsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldJsyncMask)
	}
}

const (
	RegisterDfsdm2_cr1FieldJscanShift = 4
	RegisterDfsdm2_cr1FieldJscanMask  = 0x10
)

// GetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm2_cr1Type) GetJscan() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldJscanMask) != 0
}

// SetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm2_cr1Type) SetJscan(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr1FieldJscanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldJscanMask)
	}
}

const (
	RegisterDfsdm2_cr1FieldJdmaenShift = 5
	RegisterDfsdm2_cr1FieldJdmaenMask  = 0x20
)

// GetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm2_cr1Type) GetJdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldJdmaenMask) != 0
}

// SetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm2_cr1Type) SetJdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr1FieldJdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldJdmaenMask)
	}
}

const (
	RegisterDfsdm2_cr1FieldJextselShift = 8
	RegisterDfsdm2_cr1FieldJextselMask  = 0x1f00
)

// GetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm2_cr1Type) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldJextselMask) >> RegisterDfsdm2_cr1FieldJextselShift)
}

// SetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm2_cr1Type) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldJextselMask)|(uint32(value)<<RegisterDfsdm2_cr1FieldJextselShift))
}

const (
	RegisterDfsdm2_cr1FieldJextenShift = 13
	RegisterDfsdm2_cr1FieldJextenMask  = 0x6000
)

// GetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm2_cr1Type) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldJextenMask) >> RegisterDfsdm2_cr1FieldJextenShift)
}

// SetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm2_cr1Type) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldJextenMask)|(uint32(value)<<RegisterDfsdm2_cr1FieldJextenShift))
}

const (
	RegisterDfsdm2_cr1FieldRswstartShift = 17
	RegisterDfsdm2_cr1FieldRswstartMask  = 0x20000
)

// GetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm2_cr1Type) GetRswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldRswstartMask) != 0
}

// SetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm2_cr1Type) SetRswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr1FieldRswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldRswstartMask)
	}
}

const (
	RegisterDfsdm2_cr1FieldRcontShift = 18
	RegisterDfsdm2_cr1FieldRcontMask  = 0x40000
)

// GetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm2_cr1Type) GetRcont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldRcontMask) != 0
}

// SetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm2_cr1Type) SetRcont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr1FieldRcontMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldRcontMask)
	}
}

const (
	RegisterDfsdm2_cr1FieldRsyncShift = 19
	RegisterDfsdm2_cr1FieldRsyncMask  = 0x80000
)

// GetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm2_cr1Type) GetRsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldRsyncMask) != 0
}

// SetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm2_cr1Type) SetRsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr1FieldRsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldRsyncMask)
	}
}

const (
	RegisterDfsdm2_cr1FieldRdmaenShift = 21
	RegisterDfsdm2_cr1FieldRdmaenMask  = 0x200000
)

// GetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm2_cr1Type) GetRdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldRdmaenMask) != 0
}

// SetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm2_cr1Type) SetRdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr1FieldRdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldRdmaenMask)
	}
}

const (
	RegisterDfsdm2_cr1FieldRchShift = 24
	RegisterDfsdm2_cr1FieldRchMask  = 0x7000000
)

// GetRch Regular channel selection
func (r *registerDfsdm2_cr1Type) GetRch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldRchMask) >> RegisterDfsdm2_cr1FieldRchShift)
}

// SetRch Regular channel selection
func (r *registerDfsdm2_cr1Type) SetRch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldRchMask)|(uint32(value)<<RegisterDfsdm2_cr1FieldRchShift))
}

const (
	RegisterDfsdm2_cr1FieldFastShift = 29
	RegisterDfsdm2_cr1FieldFastMask  = 0x20000000
)

// GetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm2_cr1Type) GetFast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldFastMask) != 0
}

// SetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm2_cr1Type) SetFast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr1FieldFastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldFastMask)
	}
}

const (
	RegisterDfsdm2_cr1FieldAwfselShift = 30
	RegisterDfsdm2_cr1FieldAwfselMask  = 0x40000000
)

// GetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm2_cr1Type) GetAwfsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr1FieldAwfselMask) != 0
}

// SetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm2_cr1Type) SetAwfsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr1FieldAwfselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr1FieldAwfselMask)
	}
}

// registerDfsdm3_cr1Type DFSDM control register 1
type registerDfsdm3_cr1Type uint32

const (
	RegisterDfsdm3_cr1FieldDfenShift = 0
	RegisterDfsdm3_cr1FieldDfenMask  = 0x1
)

// GetDfen DFSDM enable
func (r *registerDfsdm3_cr1Type) GetDfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldDfenMask) != 0
}

// SetDfen DFSDM enable
func (r *registerDfsdm3_cr1Type) SetDfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr1FieldDfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldDfenMask)
	}
}

const (
	RegisterDfsdm3_cr1FieldJswstartShift = 1
	RegisterDfsdm3_cr1FieldJswstartMask  = 0x2
)

// GetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm3_cr1Type) GetJswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldJswstartMask) != 0
}

// SetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm3_cr1Type) SetJswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr1FieldJswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldJswstartMask)
	}
}

const (
	RegisterDfsdm3_cr1FieldJsyncShift = 3
	RegisterDfsdm3_cr1FieldJsyncMask  = 0x8
)

// GetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm3_cr1Type) GetJsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldJsyncMask) != 0
}

// SetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm3_cr1Type) SetJsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr1FieldJsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldJsyncMask)
	}
}

const (
	RegisterDfsdm3_cr1FieldJscanShift = 4
	RegisterDfsdm3_cr1FieldJscanMask  = 0x10
)

// GetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm3_cr1Type) GetJscan() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldJscanMask) != 0
}

// SetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm3_cr1Type) SetJscan(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr1FieldJscanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldJscanMask)
	}
}

const (
	RegisterDfsdm3_cr1FieldJdmaenShift = 5
	RegisterDfsdm3_cr1FieldJdmaenMask  = 0x20
)

// GetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm3_cr1Type) GetJdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldJdmaenMask) != 0
}

// SetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm3_cr1Type) SetJdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr1FieldJdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldJdmaenMask)
	}
}

const (
	RegisterDfsdm3_cr1FieldJextselShift = 8
	RegisterDfsdm3_cr1FieldJextselMask  = 0x1f00
)

// GetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm3_cr1Type) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldJextselMask) >> RegisterDfsdm3_cr1FieldJextselShift)
}

// SetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm3_cr1Type) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldJextselMask)|(uint32(value)<<RegisterDfsdm3_cr1FieldJextselShift))
}

const (
	RegisterDfsdm3_cr1FieldJextenShift = 13
	RegisterDfsdm3_cr1FieldJextenMask  = 0x6000
)

// GetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm3_cr1Type) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldJextenMask) >> RegisterDfsdm3_cr1FieldJextenShift)
}

// SetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm3_cr1Type) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldJextenMask)|(uint32(value)<<RegisterDfsdm3_cr1FieldJextenShift))
}

const (
	RegisterDfsdm3_cr1FieldRswstartShift = 17
	RegisterDfsdm3_cr1FieldRswstartMask  = 0x20000
)

// GetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm3_cr1Type) GetRswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldRswstartMask) != 0
}

// SetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm3_cr1Type) SetRswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr1FieldRswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldRswstartMask)
	}
}

const (
	RegisterDfsdm3_cr1FieldRcontShift = 18
	RegisterDfsdm3_cr1FieldRcontMask  = 0x40000
)

// GetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm3_cr1Type) GetRcont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldRcontMask) != 0
}

// SetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm3_cr1Type) SetRcont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr1FieldRcontMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldRcontMask)
	}
}

const (
	RegisterDfsdm3_cr1FieldRsyncShift = 19
	RegisterDfsdm3_cr1FieldRsyncMask  = 0x80000
)

// GetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm3_cr1Type) GetRsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldRsyncMask) != 0
}

// SetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm3_cr1Type) SetRsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr1FieldRsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldRsyncMask)
	}
}

const (
	RegisterDfsdm3_cr1FieldRdmaenShift = 21
	RegisterDfsdm3_cr1FieldRdmaenMask  = 0x200000
)

// GetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm3_cr1Type) GetRdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldRdmaenMask) != 0
}

// SetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm3_cr1Type) SetRdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr1FieldRdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldRdmaenMask)
	}
}

const (
	RegisterDfsdm3_cr1FieldRchShift = 24
	RegisterDfsdm3_cr1FieldRchMask  = 0x7000000
)

// GetRch Regular channel selection
func (r *registerDfsdm3_cr1Type) GetRch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldRchMask) >> RegisterDfsdm3_cr1FieldRchShift)
}

// SetRch Regular channel selection
func (r *registerDfsdm3_cr1Type) SetRch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldRchMask)|(uint32(value)<<RegisterDfsdm3_cr1FieldRchShift))
}

const (
	RegisterDfsdm3_cr1FieldFastShift = 29
	RegisterDfsdm3_cr1FieldFastMask  = 0x20000000
)

// GetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm3_cr1Type) GetFast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldFastMask) != 0
}

// SetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm3_cr1Type) SetFast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr1FieldFastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldFastMask)
	}
}

const (
	RegisterDfsdm3_cr1FieldAwfselShift = 30
	RegisterDfsdm3_cr1FieldAwfselMask  = 0x40000000
)

// GetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm3_cr1Type) GetAwfsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr1FieldAwfselMask) != 0
}

// SetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm3_cr1Type) SetAwfsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr1FieldAwfselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr1FieldAwfselMask)
	}
}

// registerDfsdm0_cr2Type DFSDM control register 2
type registerDfsdm0_cr2Type uint32

const (
	RegisterDfsdm0_cr2FieldJeocieShift = 0
	RegisterDfsdm0_cr2FieldJeocieMask  = 0x1
)

// GetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm0_cr2Type) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr2FieldJeocieMask) != 0
}

// SetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm0_cr2Type) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr2FieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr2FieldJeocieMask)
	}
}

const (
	RegisterDfsdm0_cr2FieldReocieShift = 1
	RegisterDfsdm0_cr2FieldReocieMask  = 0x2
)

// GetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm0_cr2Type) GetReocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr2FieldReocieMask) != 0
}

// SetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm0_cr2Type) SetReocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr2FieldReocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr2FieldReocieMask)
	}
}

const (
	RegisterDfsdm0_cr2FieldJovrieShift = 2
	RegisterDfsdm0_cr2FieldJovrieMask  = 0x4
)

// GetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm0_cr2Type) GetJovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr2FieldJovrieMask) != 0
}

// SetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm0_cr2Type) SetJovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr2FieldJovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr2FieldJovrieMask)
	}
}

const (
	RegisterDfsdm0_cr2FieldRovrieShift = 3
	RegisterDfsdm0_cr2FieldRovrieMask  = 0x8
)

// GetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm0_cr2Type) GetRovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr2FieldRovrieMask) != 0
}

// SetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm0_cr2Type) SetRovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr2FieldRovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr2FieldRovrieMask)
	}
}

const (
	RegisterDfsdm0_cr2FieldAwdieShift = 4
	RegisterDfsdm0_cr2FieldAwdieMask  = 0x10
)

// GetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm0_cr2Type) GetAwdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr2FieldAwdieMask) != 0
}

// SetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm0_cr2Type) SetAwdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr2FieldAwdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr2FieldAwdieMask)
	}
}

const (
	RegisterDfsdm0_cr2FieldScdieShift = 5
	RegisterDfsdm0_cr2FieldScdieMask  = 0x20
)

// GetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm0_cr2Type) GetScdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr2FieldScdieMask) != 0
}

// SetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm0_cr2Type) SetScdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr2FieldScdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr2FieldScdieMask)
	}
}

const (
	RegisterDfsdm0_cr2FieldCkabieShift = 6
	RegisterDfsdm0_cr2FieldCkabieMask  = 0x40
)

// GetCkabie Clock absence interrupt enable
func (r *registerDfsdm0_cr2Type) GetCkabie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr2FieldCkabieMask) != 0
}

// SetCkabie Clock absence interrupt enable
func (r *registerDfsdm0_cr2Type) SetCkabie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_cr2FieldCkabieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr2FieldCkabieMask)
	}
}

const (
	RegisterDfsdm0_cr2FieldExchShift = 8
	RegisterDfsdm0_cr2FieldExchMask  = 0xff00
)

// GetExch Extremes detector channel selection
func (r *registerDfsdm0_cr2Type) GetExch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr2FieldExchMask) >> RegisterDfsdm0_cr2FieldExchShift)
}

// SetExch Extremes detector channel selection
func (r *registerDfsdm0_cr2Type) SetExch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr2FieldExchMask)|(uint32(value)<<RegisterDfsdm0_cr2FieldExchShift))
}

const (
	RegisterDfsdm0_cr2FieldAwdchShift = 16
	RegisterDfsdm0_cr2FieldAwdchMask  = 0xff0000
)

// GetAwdch Analog watchdog channel selection
func (r *registerDfsdm0_cr2Type) GetAwdch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cr2FieldAwdchMask) >> RegisterDfsdm0_cr2FieldAwdchShift)
}

// SetAwdch Analog watchdog channel selection
func (r *registerDfsdm0_cr2Type) SetAwdch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cr2FieldAwdchMask)|(uint32(value)<<RegisterDfsdm0_cr2FieldAwdchShift))
}

// registerDfsdm1_cr2Type DFSDM control register 2
type registerDfsdm1_cr2Type uint32

const (
	RegisterDfsdm1_cr2FieldJeocieShift = 0
	RegisterDfsdm1_cr2FieldJeocieMask  = 0x1
)

// GetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm1_cr2Type) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr2FieldJeocieMask) != 0
}

// SetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm1_cr2Type) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr2FieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr2FieldJeocieMask)
	}
}

const (
	RegisterDfsdm1_cr2FieldReocieShift = 1
	RegisterDfsdm1_cr2FieldReocieMask  = 0x2
)

// GetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm1_cr2Type) GetReocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr2FieldReocieMask) != 0
}

// SetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm1_cr2Type) SetReocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr2FieldReocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr2FieldReocieMask)
	}
}

const (
	RegisterDfsdm1_cr2FieldJovrieShift = 2
	RegisterDfsdm1_cr2FieldJovrieMask  = 0x4
)

// GetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm1_cr2Type) GetJovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr2FieldJovrieMask) != 0
}

// SetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm1_cr2Type) SetJovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr2FieldJovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr2FieldJovrieMask)
	}
}

const (
	RegisterDfsdm1_cr2FieldRovrieShift = 3
	RegisterDfsdm1_cr2FieldRovrieMask  = 0x8
)

// GetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm1_cr2Type) GetRovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr2FieldRovrieMask) != 0
}

// SetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm1_cr2Type) SetRovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr2FieldRovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr2FieldRovrieMask)
	}
}

const (
	RegisterDfsdm1_cr2FieldAwdieShift = 4
	RegisterDfsdm1_cr2FieldAwdieMask  = 0x10
)

// GetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm1_cr2Type) GetAwdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr2FieldAwdieMask) != 0
}

// SetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm1_cr2Type) SetAwdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr2FieldAwdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr2FieldAwdieMask)
	}
}

const (
	RegisterDfsdm1_cr2FieldScdieShift = 5
	RegisterDfsdm1_cr2FieldScdieMask  = 0x20
)

// GetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm1_cr2Type) GetScdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr2FieldScdieMask) != 0
}

// SetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm1_cr2Type) SetScdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr2FieldScdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr2FieldScdieMask)
	}
}

const (
	RegisterDfsdm1_cr2FieldCkabieShift = 6
	RegisterDfsdm1_cr2FieldCkabieMask  = 0x40
)

// GetCkabie Clock absence interrupt enable
func (r *registerDfsdm1_cr2Type) GetCkabie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr2FieldCkabieMask) != 0
}

// SetCkabie Clock absence interrupt enable
func (r *registerDfsdm1_cr2Type) SetCkabie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_cr2FieldCkabieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr2FieldCkabieMask)
	}
}

const (
	RegisterDfsdm1_cr2FieldExchShift = 8
	RegisterDfsdm1_cr2FieldExchMask  = 0xff00
)

// GetExch Extremes detector channel selection
func (r *registerDfsdm1_cr2Type) GetExch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr2FieldExchMask) >> RegisterDfsdm1_cr2FieldExchShift)
}

// SetExch Extremes detector channel selection
func (r *registerDfsdm1_cr2Type) SetExch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr2FieldExchMask)|(uint32(value)<<RegisterDfsdm1_cr2FieldExchShift))
}

const (
	RegisterDfsdm1_cr2FieldAwdchShift = 16
	RegisterDfsdm1_cr2FieldAwdchMask  = 0xff0000
)

// GetAwdch Analog watchdog channel selection
func (r *registerDfsdm1_cr2Type) GetAwdch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cr2FieldAwdchMask) >> RegisterDfsdm1_cr2FieldAwdchShift)
}

// SetAwdch Analog watchdog channel selection
func (r *registerDfsdm1_cr2Type) SetAwdch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cr2FieldAwdchMask)|(uint32(value)<<RegisterDfsdm1_cr2FieldAwdchShift))
}

// registerDfsdm2_cr2Type DFSDM control register 2
type registerDfsdm2_cr2Type uint32

const (
	RegisterDfsdm2_cr2FieldJeocieShift = 0
	RegisterDfsdm2_cr2FieldJeocieMask  = 0x1
)

// GetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm2_cr2Type) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr2FieldJeocieMask) != 0
}

// SetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm2_cr2Type) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr2FieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr2FieldJeocieMask)
	}
}

const (
	RegisterDfsdm2_cr2FieldReocieShift = 1
	RegisterDfsdm2_cr2FieldReocieMask  = 0x2
)

// GetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm2_cr2Type) GetReocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr2FieldReocieMask) != 0
}

// SetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm2_cr2Type) SetReocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr2FieldReocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr2FieldReocieMask)
	}
}

const (
	RegisterDfsdm2_cr2FieldJovrieShift = 2
	RegisterDfsdm2_cr2FieldJovrieMask  = 0x4
)

// GetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm2_cr2Type) GetJovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr2FieldJovrieMask) != 0
}

// SetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm2_cr2Type) SetJovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr2FieldJovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr2FieldJovrieMask)
	}
}

const (
	RegisterDfsdm2_cr2FieldRovrieShift = 3
	RegisterDfsdm2_cr2FieldRovrieMask  = 0x8
)

// GetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm2_cr2Type) GetRovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr2FieldRovrieMask) != 0
}

// SetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm2_cr2Type) SetRovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr2FieldRovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr2FieldRovrieMask)
	}
}

const (
	RegisterDfsdm2_cr2FieldAwdieShift = 4
	RegisterDfsdm2_cr2FieldAwdieMask  = 0x10
)

// GetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm2_cr2Type) GetAwdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr2FieldAwdieMask) != 0
}

// SetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm2_cr2Type) SetAwdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr2FieldAwdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr2FieldAwdieMask)
	}
}

const (
	RegisterDfsdm2_cr2FieldScdieShift = 5
	RegisterDfsdm2_cr2FieldScdieMask  = 0x20
)

// GetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm2_cr2Type) GetScdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr2FieldScdieMask) != 0
}

// SetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm2_cr2Type) SetScdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr2FieldScdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr2FieldScdieMask)
	}
}

const (
	RegisterDfsdm2_cr2FieldCkabieShift = 6
	RegisterDfsdm2_cr2FieldCkabieMask  = 0x40
)

// GetCkabie Clock absence interrupt enable
func (r *registerDfsdm2_cr2Type) GetCkabie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr2FieldCkabieMask) != 0
}

// SetCkabie Clock absence interrupt enable
func (r *registerDfsdm2_cr2Type) SetCkabie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_cr2FieldCkabieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr2FieldCkabieMask)
	}
}

const (
	RegisterDfsdm2_cr2FieldExchShift = 8
	RegisterDfsdm2_cr2FieldExchMask  = 0xff00
)

// GetExch Extremes detector channel selection
func (r *registerDfsdm2_cr2Type) GetExch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr2FieldExchMask) >> RegisterDfsdm2_cr2FieldExchShift)
}

// SetExch Extremes detector channel selection
func (r *registerDfsdm2_cr2Type) SetExch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr2FieldExchMask)|(uint32(value)<<RegisterDfsdm2_cr2FieldExchShift))
}

const (
	RegisterDfsdm2_cr2FieldAwdchShift = 16
	RegisterDfsdm2_cr2FieldAwdchMask  = 0xff0000
)

// GetAwdch Analog watchdog channel selection
func (r *registerDfsdm2_cr2Type) GetAwdch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cr2FieldAwdchMask) >> RegisterDfsdm2_cr2FieldAwdchShift)
}

// SetAwdch Analog watchdog channel selection
func (r *registerDfsdm2_cr2Type) SetAwdch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cr2FieldAwdchMask)|(uint32(value)<<RegisterDfsdm2_cr2FieldAwdchShift))
}

// registerDfsdm3_cr2Type DFSDM control register 2
type registerDfsdm3_cr2Type uint32

const (
	RegisterDfsdm3_cr2FieldJeocieShift = 0
	RegisterDfsdm3_cr2FieldJeocieMask  = 0x1
)

// GetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm3_cr2Type) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr2FieldJeocieMask) != 0
}

// SetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm3_cr2Type) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr2FieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr2FieldJeocieMask)
	}
}

const (
	RegisterDfsdm3_cr2FieldReocieShift = 1
	RegisterDfsdm3_cr2FieldReocieMask  = 0x2
)

// GetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm3_cr2Type) GetReocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr2FieldReocieMask) != 0
}

// SetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm3_cr2Type) SetReocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr2FieldReocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr2FieldReocieMask)
	}
}

const (
	RegisterDfsdm3_cr2FieldJovrieShift = 2
	RegisterDfsdm3_cr2FieldJovrieMask  = 0x4
)

// GetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm3_cr2Type) GetJovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr2FieldJovrieMask) != 0
}

// SetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm3_cr2Type) SetJovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr2FieldJovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr2FieldJovrieMask)
	}
}

const (
	RegisterDfsdm3_cr2FieldRovrieShift = 3
	RegisterDfsdm3_cr2FieldRovrieMask  = 0x8
)

// GetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm3_cr2Type) GetRovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr2FieldRovrieMask) != 0
}

// SetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm3_cr2Type) SetRovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr2FieldRovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr2FieldRovrieMask)
	}
}

const (
	RegisterDfsdm3_cr2FieldAwdieShift = 4
	RegisterDfsdm3_cr2FieldAwdieMask  = 0x10
)

// GetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm3_cr2Type) GetAwdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr2FieldAwdieMask) != 0
}

// SetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm3_cr2Type) SetAwdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr2FieldAwdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr2FieldAwdieMask)
	}
}

const (
	RegisterDfsdm3_cr2FieldScdieShift = 5
	RegisterDfsdm3_cr2FieldScdieMask  = 0x20
)

// GetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm3_cr2Type) GetScdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr2FieldScdieMask) != 0
}

// SetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm3_cr2Type) SetScdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr2FieldScdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr2FieldScdieMask)
	}
}

const (
	RegisterDfsdm3_cr2FieldCkabieShift = 6
	RegisterDfsdm3_cr2FieldCkabieMask  = 0x40
)

// GetCkabie Clock absence interrupt enable
func (r *registerDfsdm3_cr2Type) GetCkabie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr2FieldCkabieMask) != 0
}

// SetCkabie Clock absence interrupt enable
func (r *registerDfsdm3_cr2Type) SetCkabie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_cr2FieldCkabieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr2FieldCkabieMask)
	}
}

const (
	RegisterDfsdm3_cr2FieldExchShift = 8
	RegisterDfsdm3_cr2FieldExchMask  = 0xff00
)

// GetExch Extremes detector channel selection
func (r *registerDfsdm3_cr2Type) GetExch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr2FieldExchMask) >> RegisterDfsdm3_cr2FieldExchShift)
}

// SetExch Extremes detector channel selection
func (r *registerDfsdm3_cr2Type) SetExch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr2FieldExchMask)|(uint32(value)<<RegisterDfsdm3_cr2FieldExchShift))
}

const (
	RegisterDfsdm3_cr2FieldAwdchShift = 16
	RegisterDfsdm3_cr2FieldAwdchMask  = 0xff0000
)

// GetAwdch Analog watchdog channel selection
func (r *registerDfsdm3_cr2Type) GetAwdch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cr2FieldAwdchMask) >> RegisterDfsdm3_cr2FieldAwdchShift)
}

// SetAwdch Analog watchdog channel selection
func (r *registerDfsdm3_cr2Type) SetAwdch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cr2FieldAwdchMask)|(uint32(value)<<RegisterDfsdm3_cr2FieldAwdchShift))
}

// registerDfsdm0_isrType DFSDM interrupt and status register
type registerDfsdm0_isrType uint32

const (
	RegisterDfsdm0_isrFieldJeocfShift = 0
	RegisterDfsdm0_isrFieldJeocfMask  = 0x1
)

// GetJeocf End of injected conversion flag
func (r *registerDfsdm0_isrType) GetJeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_isrFieldJeocfMask) != 0
}

// SetJeocf End of injected conversion flag
func (r *registerDfsdm0_isrType) SetJeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_isrFieldJeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_isrFieldJeocfMask)
	}
}

const (
	RegisterDfsdm0_isrFieldReocfShift = 1
	RegisterDfsdm0_isrFieldReocfMask  = 0x2
)

// GetReocf End of regular conversion flag
func (r *registerDfsdm0_isrType) GetReocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_isrFieldReocfMask) != 0
}

// SetReocf End of regular conversion flag
func (r *registerDfsdm0_isrType) SetReocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_isrFieldReocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_isrFieldReocfMask)
	}
}

const (
	RegisterDfsdm0_isrFieldJovrfShift = 2
	RegisterDfsdm0_isrFieldJovrfMask  = 0x4
)

// GetJovrf Injected conversion overrun flag
func (r *registerDfsdm0_isrType) GetJovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_isrFieldJovrfMask) != 0
}

// SetJovrf Injected conversion overrun flag
func (r *registerDfsdm0_isrType) SetJovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_isrFieldJovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_isrFieldJovrfMask)
	}
}

const (
	RegisterDfsdm0_isrFieldRovrfShift = 3
	RegisterDfsdm0_isrFieldRovrfMask  = 0x8
)

// GetRovrf Regular conversion overrun flag
func (r *registerDfsdm0_isrType) GetRovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_isrFieldRovrfMask) != 0
}

// SetRovrf Regular conversion overrun flag
func (r *registerDfsdm0_isrType) SetRovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_isrFieldRovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_isrFieldRovrfMask)
	}
}

const (
	RegisterDfsdm0_isrFieldAwdfShift = 4
	RegisterDfsdm0_isrFieldAwdfMask  = 0x10
)

// GetAwdf Analog watchdog
func (r *registerDfsdm0_isrType) GetAwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_isrFieldAwdfMask) != 0
}

// SetAwdf Analog watchdog
func (r *registerDfsdm0_isrType) SetAwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_isrFieldAwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_isrFieldAwdfMask)
	}
}

const (
	RegisterDfsdm0_isrFieldJcipShift = 13
	RegisterDfsdm0_isrFieldJcipMask  = 0x2000
)

// GetJcip Injected conversion in progress status
func (r *registerDfsdm0_isrType) GetJcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_isrFieldJcipMask) != 0
}

// SetJcip Injected conversion in progress status
func (r *registerDfsdm0_isrType) SetJcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_isrFieldJcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_isrFieldJcipMask)
	}
}

const (
	RegisterDfsdm0_isrFieldRcipShift = 14
	RegisterDfsdm0_isrFieldRcipMask  = 0x4000
)

// GetRcip Regular conversion in progress status
func (r *registerDfsdm0_isrType) GetRcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_isrFieldRcipMask) != 0
}

// SetRcip Regular conversion in progress status
func (r *registerDfsdm0_isrType) SetRcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_isrFieldRcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_isrFieldRcipMask)
	}
}

const (
	RegisterDfsdm0_isrFieldCkabfShift = 16
	RegisterDfsdm0_isrFieldCkabfMask  = 0xff0000
)

// GetCkabf Clock absence flag
func (r *registerDfsdm0_isrType) GetCkabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_isrFieldCkabfMask) >> RegisterDfsdm0_isrFieldCkabfShift)
}

// SetCkabf Clock absence flag
func (r *registerDfsdm0_isrType) SetCkabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_isrFieldCkabfMask)|(uint32(value)<<RegisterDfsdm0_isrFieldCkabfShift))
}

const (
	RegisterDfsdm0_isrFieldScdfShift = 24
	RegisterDfsdm0_isrFieldScdfMask  = 0xff000000
)

// GetScdf short-circuit detector flag
func (r *registerDfsdm0_isrType) GetScdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_isrFieldScdfMask) >> RegisterDfsdm0_isrFieldScdfShift)
}

// SetScdf short-circuit detector flag
func (r *registerDfsdm0_isrType) SetScdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_isrFieldScdfMask)|(uint32(value)<<RegisterDfsdm0_isrFieldScdfShift))
}

// registerDfsdm1_isrType DFSDM interrupt and status register
type registerDfsdm1_isrType uint32

const (
	RegisterDfsdm1_isrFieldJeocfShift = 0
	RegisterDfsdm1_isrFieldJeocfMask  = 0x1
)

// GetJeocf End of injected conversion flag
func (r *registerDfsdm1_isrType) GetJeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_isrFieldJeocfMask) != 0
}

// SetJeocf End of injected conversion flag
func (r *registerDfsdm1_isrType) SetJeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_isrFieldJeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_isrFieldJeocfMask)
	}
}

const (
	RegisterDfsdm1_isrFieldReocfShift = 1
	RegisterDfsdm1_isrFieldReocfMask  = 0x2
)

// GetReocf End of regular conversion flag
func (r *registerDfsdm1_isrType) GetReocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_isrFieldReocfMask) != 0
}

// SetReocf End of regular conversion flag
func (r *registerDfsdm1_isrType) SetReocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_isrFieldReocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_isrFieldReocfMask)
	}
}

const (
	RegisterDfsdm1_isrFieldJovrfShift = 2
	RegisterDfsdm1_isrFieldJovrfMask  = 0x4
)

// GetJovrf Injected conversion overrun flag
func (r *registerDfsdm1_isrType) GetJovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_isrFieldJovrfMask) != 0
}

// SetJovrf Injected conversion overrun flag
func (r *registerDfsdm1_isrType) SetJovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_isrFieldJovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_isrFieldJovrfMask)
	}
}

const (
	RegisterDfsdm1_isrFieldRovrfShift = 3
	RegisterDfsdm1_isrFieldRovrfMask  = 0x8
)

// GetRovrf Regular conversion overrun flag
func (r *registerDfsdm1_isrType) GetRovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_isrFieldRovrfMask) != 0
}

// SetRovrf Regular conversion overrun flag
func (r *registerDfsdm1_isrType) SetRovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_isrFieldRovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_isrFieldRovrfMask)
	}
}

const (
	RegisterDfsdm1_isrFieldAwdfShift = 4
	RegisterDfsdm1_isrFieldAwdfMask  = 0x10
)

// GetAwdf Analog watchdog
func (r *registerDfsdm1_isrType) GetAwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_isrFieldAwdfMask) != 0
}

// SetAwdf Analog watchdog
func (r *registerDfsdm1_isrType) SetAwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_isrFieldAwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_isrFieldAwdfMask)
	}
}

const (
	RegisterDfsdm1_isrFieldJcipShift = 13
	RegisterDfsdm1_isrFieldJcipMask  = 0x2000
)

// GetJcip Injected conversion in progress status
func (r *registerDfsdm1_isrType) GetJcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_isrFieldJcipMask) != 0
}

// SetJcip Injected conversion in progress status
func (r *registerDfsdm1_isrType) SetJcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_isrFieldJcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_isrFieldJcipMask)
	}
}

const (
	RegisterDfsdm1_isrFieldRcipShift = 14
	RegisterDfsdm1_isrFieldRcipMask  = 0x4000
)

// GetRcip Regular conversion in progress status
func (r *registerDfsdm1_isrType) GetRcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_isrFieldRcipMask) != 0
}

// SetRcip Regular conversion in progress status
func (r *registerDfsdm1_isrType) SetRcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_isrFieldRcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_isrFieldRcipMask)
	}
}

const (
	RegisterDfsdm1_isrFieldCkabfShift = 16
	RegisterDfsdm1_isrFieldCkabfMask  = 0xff0000
)

// GetCkabf Clock absence flag
func (r *registerDfsdm1_isrType) GetCkabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_isrFieldCkabfMask) >> RegisterDfsdm1_isrFieldCkabfShift)
}

// SetCkabf Clock absence flag
func (r *registerDfsdm1_isrType) SetCkabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_isrFieldCkabfMask)|(uint32(value)<<RegisterDfsdm1_isrFieldCkabfShift))
}

const (
	RegisterDfsdm1_isrFieldScdfShift = 24
	RegisterDfsdm1_isrFieldScdfMask  = 0xff000000
)

// GetScdf short-circuit detector flag
func (r *registerDfsdm1_isrType) GetScdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_isrFieldScdfMask) >> RegisterDfsdm1_isrFieldScdfShift)
}

// SetScdf short-circuit detector flag
func (r *registerDfsdm1_isrType) SetScdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_isrFieldScdfMask)|(uint32(value)<<RegisterDfsdm1_isrFieldScdfShift))
}

// registerDfsdm2_isrType DFSDM interrupt and status register
type registerDfsdm2_isrType uint32

const (
	RegisterDfsdm2_isrFieldJeocfShift = 0
	RegisterDfsdm2_isrFieldJeocfMask  = 0x1
)

// GetJeocf End of injected conversion flag
func (r *registerDfsdm2_isrType) GetJeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_isrFieldJeocfMask) != 0
}

// SetJeocf End of injected conversion flag
func (r *registerDfsdm2_isrType) SetJeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_isrFieldJeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_isrFieldJeocfMask)
	}
}

const (
	RegisterDfsdm2_isrFieldReocfShift = 1
	RegisterDfsdm2_isrFieldReocfMask  = 0x2
)

// GetReocf End of regular conversion flag
func (r *registerDfsdm2_isrType) GetReocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_isrFieldReocfMask) != 0
}

// SetReocf End of regular conversion flag
func (r *registerDfsdm2_isrType) SetReocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_isrFieldReocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_isrFieldReocfMask)
	}
}

const (
	RegisterDfsdm2_isrFieldJovrfShift = 2
	RegisterDfsdm2_isrFieldJovrfMask  = 0x4
)

// GetJovrf Injected conversion overrun flag
func (r *registerDfsdm2_isrType) GetJovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_isrFieldJovrfMask) != 0
}

// SetJovrf Injected conversion overrun flag
func (r *registerDfsdm2_isrType) SetJovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_isrFieldJovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_isrFieldJovrfMask)
	}
}

const (
	RegisterDfsdm2_isrFieldRovrfShift = 3
	RegisterDfsdm2_isrFieldRovrfMask  = 0x8
)

// GetRovrf Regular conversion overrun flag
func (r *registerDfsdm2_isrType) GetRovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_isrFieldRovrfMask) != 0
}

// SetRovrf Regular conversion overrun flag
func (r *registerDfsdm2_isrType) SetRovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_isrFieldRovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_isrFieldRovrfMask)
	}
}

const (
	RegisterDfsdm2_isrFieldAwdfShift = 4
	RegisterDfsdm2_isrFieldAwdfMask  = 0x10
)

// GetAwdf Analog watchdog
func (r *registerDfsdm2_isrType) GetAwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_isrFieldAwdfMask) != 0
}

// SetAwdf Analog watchdog
func (r *registerDfsdm2_isrType) SetAwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_isrFieldAwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_isrFieldAwdfMask)
	}
}

const (
	RegisterDfsdm2_isrFieldJcipShift = 13
	RegisterDfsdm2_isrFieldJcipMask  = 0x2000
)

// GetJcip Injected conversion in progress status
func (r *registerDfsdm2_isrType) GetJcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_isrFieldJcipMask) != 0
}

// SetJcip Injected conversion in progress status
func (r *registerDfsdm2_isrType) SetJcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_isrFieldJcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_isrFieldJcipMask)
	}
}

const (
	RegisterDfsdm2_isrFieldRcipShift = 14
	RegisterDfsdm2_isrFieldRcipMask  = 0x4000
)

// GetRcip Regular conversion in progress status
func (r *registerDfsdm2_isrType) GetRcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_isrFieldRcipMask) != 0
}

// SetRcip Regular conversion in progress status
func (r *registerDfsdm2_isrType) SetRcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_isrFieldRcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_isrFieldRcipMask)
	}
}

const (
	RegisterDfsdm2_isrFieldCkabfShift = 16
	RegisterDfsdm2_isrFieldCkabfMask  = 0xff0000
)

// GetCkabf Clock absence flag
func (r *registerDfsdm2_isrType) GetCkabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_isrFieldCkabfMask) >> RegisterDfsdm2_isrFieldCkabfShift)
}

// SetCkabf Clock absence flag
func (r *registerDfsdm2_isrType) SetCkabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_isrFieldCkabfMask)|(uint32(value)<<RegisterDfsdm2_isrFieldCkabfShift))
}

const (
	RegisterDfsdm2_isrFieldScdfShift = 24
	RegisterDfsdm2_isrFieldScdfMask  = 0xff000000
)

// GetScdf short-circuit detector flag
func (r *registerDfsdm2_isrType) GetScdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_isrFieldScdfMask) >> RegisterDfsdm2_isrFieldScdfShift)
}

// SetScdf short-circuit detector flag
func (r *registerDfsdm2_isrType) SetScdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_isrFieldScdfMask)|(uint32(value)<<RegisterDfsdm2_isrFieldScdfShift))
}

// registerDfsdm3_isrType DFSDM interrupt and status register
type registerDfsdm3_isrType uint32

const (
	RegisterDfsdm3_isrFieldJeocfShift = 0
	RegisterDfsdm3_isrFieldJeocfMask  = 0x1
)

// GetJeocf End of injected conversion flag
func (r *registerDfsdm3_isrType) GetJeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_isrFieldJeocfMask) != 0
}

// SetJeocf End of injected conversion flag
func (r *registerDfsdm3_isrType) SetJeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_isrFieldJeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_isrFieldJeocfMask)
	}
}

const (
	RegisterDfsdm3_isrFieldReocfShift = 1
	RegisterDfsdm3_isrFieldReocfMask  = 0x2
)

// GetReocf End of regular conversion flag
func (r *registerDfsdm3_isrType) GetReocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_isrFieldReocfMask) != 0
}

// SetReocf End of regular conversion flag
func (r *registerDfsdm3_isrType) SetReocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_isrFieldReocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_isrFieldReocfMask)
	}
}

const (
	RegisterDfsdm3_isrFieldJovrfShift = 2
	RegisterDfsdm3_isrFieldJovrfMask  = 0x4
)

// GetJovrf Injected conversion overrun flag
func (r *registerDfsdm3_isrType) GetJovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_isrFieldJovrfMask) != 0
}

// SetJovrf Injected conversion overrun flag
func (r *registerDfsdm3_isrType) SetJovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_isrFieldJovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_isrFieldJovrfMask)
	}
}

const (
	RegisterDfsdm3_isrFieldRovrfShift = 3
	RegisterDfsdm3_isrFieldRovrfMask  = 0x8
)

// GetRovrf Regular conversion overrun flag
func (r *registerDfsdm3_isrType) GetRovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_isrFieldRovrfMask) != 0
}

// SetRovrf Regular conversion overrun flag
func (r *registerDfsdm3_isrType) SetRovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_isrFieldRovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_isrFieldRovrfMask)
	}
}

const (
	RegisterDfsdm3_isrFieldAwdfShift = 4
	RegisterDfsdm3_isrFieldAwdfMask  = 0x10
)

// GetAwdf Analog watchdog
func (r *registerDfsdm3_isrType) GetAwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_isrFieldAwdfMask) != 0
}

// SetAwdf Analog watchdog
func (r *registerDfsdm3_isrType) SetAwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_isrFieldAwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_isrFieldAwdfMask)
	}
}

const (
	RegisterDfsdm3_isrFieldJcipShift = 13
	RegisterDfsdm3_isrFieldJcipMask  = 0x2000
)

// GetJcip Injected conversion in progress status
func (r *registerDfsdm3_isrType) GetJcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_isrFieldJcipMask) != 0
}

// SetJcip Injected conversion in progress status
func (r *registerDfsdm3_isrType) SetJcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_isrFieldJcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_isrFieldJcipMask)
	}
}

const (
	RegisterDfsdm3_isrFieldRcipShift = 14
	RegisterDfsdm3_isrFieldRcipMask  = 0x4000
)

// GetRcip Regular conversion in progress status
func (r *registerDfsdm3_isrType) GetRcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_isrFieldRcipMask) != 0
}

// SetRcip Regular conversion in progress status
func (r *registerDfsdm3_isrType) SetRcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_isrFieldRcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_isrFieldRcipMask)
	}
}

const (
	RegisterDfsdm3_isrFieldCkabfShift = 16
	RegisterDfsdm3_isrFieldCkabfMask  = 0xff0000
)

// GetCkabf Clock absence flag
func (r *registerDfsdm3_isrType) GetCkabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_isrFieldCkabfMask) >> RegisterDfsdm3_isrFieldCkabfShift)
}

// SetCkabf Clock absence flag
func (r *registerDfsdm3_isrType) SetCkabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_isrFieldCkabfMask)|(uint32(value)<<RegisterDfsdm3_isrFieldCkabfShift))
}

const (
	RegisterDfsdm3_isrFieldScdfShift = 24
	RegisterDfsdm3_isrFieldScdfMask  = 0xff000000
)

// GetScdf short-circuit detector flag
func (r *registerDfsdm3_isrType) GetScdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_isrFieldScdfMask) >> RegisterDfsdm3_isrFieldScdfShift)
}

// SetScdf short-circuit detector flag
func (r *registerDfsdm3_isrType) SetScdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_isrFieldScdfMask)|(uint32(value)<<RegisterDfsdm3_isrFieldScdfShift))
}

// registerDfsdm0_icrType DFSDM interrupt flag clear register
type registerDfsdm0_icrType uint32

const (
	RegisterDfsdm0_icrFieldClrjovrfShift = 2
	RegisterDfsdm0_icrFieldClrjovrfMask  = 0x4
)

// GetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm0_icrType) GetClrjovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_icrFieldClrjovrfMask) != 0
}

// SetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm0_icrType) SetClrjovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_icrFieldClrjovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_icrFieldClrjovrfMask)
	}
}

const (
	RegisterDfsdm0_icrFieldClrrovrfShift = 3
	RegisterDfsdm0_icrFieldClrrovrfMask  = 0x8
)

// GetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm0_icrType) GetClrrovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_icrFieldClrrovrfMask) != 0
}

// SetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm0_icrType) SetClrrovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_icrFieldClrrovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_icrFieldClrrovrfMask)
	}
}

const (
	RegisterDfsdm0_icrFieldClrckabfShift = 16
	RegisterDfsdm0_icrFieldClrckabfMask  = 0xff0000
)

// GetClrckabf Clear the clock absence flag
func (r *registerDfsdm0_icrType) GetClrckabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_icrFieldClrckabfMask) >> RegisterDfsdm0_icrFieldClrckabfShift)
}

// SetClrckabf Clear the clock absence flag
func (r *registerDfsdm0_icrType) SetClrckabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_icrFieldClrckabfMask)|(uint32(value)<<RegisterDfsdm0_icrFieldClrckabfShift))
}

const (
	RegisterDfsdm0_icrFieldClrscdfShift = 24
	RegisterDfsdm0_icrFieldClrscdfMask  = 0xff000000
)

// GetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm0_icrType) GetClrscdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_icrFieldClrscdfMask) >> RegisterDfsdm0_icrFieldClrscdfShift)
}

// SetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm0_icrType) SetClrscdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_icrFieldClrscdfMask)|(uint32(value)<<RegisterDfsdm0_icrFieldClrscdfShift))
}

// registerDfsdm1_icrType DFSDM interrupt flag clear register
type registerDfsdm1_icrType uint32

const (
	RegisterDfsdm1_icrFieldClrjovrfShift = 2
	RegisterDfsdm1_icrFieldClrjovrfMask  = 0x4
)

// GetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm1_icrType) GetClrjovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_icrFieldClrjovrfMask) != 0
}

// SetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm1_icrType) SetClrjovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_icrFieldClrjovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_icrFieldClrjovrfMask)
	}
}

const (
	RegisterDfsdm1_icrFieldClrrovrfShift = 3
	RegisterDfsdm1_icrFieldClrrovrfMask  = 0x8
)

// GetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm1_icrType) GetClrrovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_icrFieldClrrovrfMask) != 0
}

// SetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm1_icrType) SetClrrovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_icrFieldClrrovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_icrFieldClrrovrfMask)
	}
}

const (
	RegisterDfsdm1_icrFieldClrckabfShift = 16
	RegisterDfsdm1_icrFieldClrckabfMask  = 0xff0000
)

// GetClrckabf Clear the clock absence flag
func (r *registerDfsdm1_icrType) GetClrckabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_icrFieldClrckabfMask) >> RegisterDfsdm1_icrFieldClrckabfShift)
}

// SetClrckabf Clear the clock absence flag
func (r *registerDfsdm1_icrType) SetClrckabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_icrFieldClrckabfMask)|(uint32(value)<<RegisterDfsdm1_icrFieldClrckabfShift))
}

const (
	RegisterDfsdm1_icrFieldClrscdfShift = 24
	RegisterDfsdm1_icrFieldClrscdfMask  = 0xff000000
)

// GetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm1_icrType) GetClrscdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_icrFieldClrscdfMask) >> RegisterDfsdm1_icrFieldClrscdfShift)
}

// SetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm1_icrType) SetClrscdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_icrFieldClrscdfMask)|(uint32(value)<<RegisterDfsdm1_icrFieldClrscdfShift))
}

// registerDfsdm2_icrType DFSDM interrupt flag clear register
type registerDfsdm2_icrType uint32

const (
	RegisterDfsdm2_icrFieldClrjovrfShift = 2
	RegisterDfsdm2_icrFieldClrjovrfMask  = 0x4
)

// GetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm2_icrType) GetClrjovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_icrFieldClrjovrfMask) != 0
}

// SetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm2_icrType) SetClrjovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_icrFieldClrjovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_icrFieldClrjovrfMask)
	}
}

const (
	RegisterDfsdm2_icrFieldClrrovrfShift = 3
	RegisterDfsdm2_icrFieldClrrovrfMask  = 0x8
)

// GetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm2_icrType) GetClrrovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_icrFieldClrrovrfMask) != 0
}

// SetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm2_icrType) SetClrrovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_icrFieldClrrovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_icrFieldClrrovrfMask)
	}
}

const (
	RegisterDfsdm2_icrFieldClrckabfShift = 16
	RegisterDfsdm2_icrFieldClrckabfMask  = 0xff0000
)

// GetClrckabf Clear the clock absence flag
func (r *registerDfsdm2_icrType) GetClrckabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_icrFieldClrckabfMask) >> RegisterDfsdm2_icrFieldClrckabfShift)
}

// SetClrckabf Clear the clock absence flag
func (r *registerDfsdm2_icrType) SetClrckabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_icrFieldClrckabfMask)|(uint32(value)<<RegisterDfsdm2_icrFieldClrckabfShift))
}

const (
	RegisterDfsdm2_icrFieldClrscdfShift = 24
	RegisterDfsdm2_icrFieldClrscdfMask  = 0xff000000
)

// GetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm2_icrType) GetClrscdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_icrFieldClrscdfMask) >> RegisterDfsdm2_icrFieldClrscdfShift)
}

// SetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm2_icrType) SetClrscdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_icrFieldClrscdfMask)|(uint32(value)<<RegisterDfsdm2_icrFieldClrscdfShift))
}

// registerDfsdm3_icrType DFSDM interrupt flag clear register
type registerDfsdm3_icrType uint32

const (
	RegisterDfsdm3_icrFieldClrjovrfShift = 2
	RegisterDfsdm3_icrFieldClrjovrfMask  = 0x4
)

// GetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm3_icrType) GetClrjovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_icrFieldClrjovrfMask) != 0
}

// SetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm3_icrType) SetClrjovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_icrFieldClrjovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_icrFieldClrjovrfMask)
	}
}

const (
	RegisterDfsdm3_icrFieldClrrovrfShift = 3
	RegisterDfsdm3_icrFieldClrrovrfMask  = 0x8
)

// GetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm3_icrType) GetClrrovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_icrFieldClrrovrfMask) != 0
}

// SetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm3_icrType) SetClrrovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_icrFieldClrrovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_icrFieldClrrovrfMask)
	}
}

const (
	RegisterDfsdm3_icrFieldClrckabfShift = 16
	RegisterDfsdm3_icrFieldClrckabfMask  = 0xff0000
)

// GetClrckabf Clear the clock absence flag
func (r *registerDfsdm3_icrType) GetClrckabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_icrFieldClrckabfMask) >> RegisterDfsdm3_icrFieldClrckabfShift)
}

// SetClrckabf Clear the clock absence flag
func (r *registerDfsdm3_icrType) SetClrckabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_icrFieldClrckabfMask)|(uint32(value)<<RegisterDfsdm3_icrFieldClrckabfShift))
}

const (
	RegisterDfsdm3_icrFieldClrscdfShift = 24
	RegisterDfsdm3_icrFieldClrscdfMask  = 0xff000000
)

// GetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm3_icrType) GetClrscdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_icrFieldClrscdfMask) >> RegisterDfsdm3_icrFieldClrscdfShift)
}

// SetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm3_icrType) SetClrscdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_icrFieldClrscdfMask)|(uint32(value)<<RegisterDfsdm3_icrFieldClrscdfShift))
}

// registerDfsdm0_jchgrType DFSDM injected channel group selection register
type registerDfsdm0_jchgrType uint32

const (
	RegisterDfsdm0_jchgrFieldJchgShift = 0
	RegisterDfsdm0_jchgrFieldJchgMask  = 0xff
)

// GetJchg Injected channel group selection
func (r *registerDfsdm0_jchgrType) GetJchg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_jchgrFieldJchgMask) >> RegisterDfsdm0_jchgrFieldJchgShift)
}

// SetJchg Injected channel group selection
func (r *registerDfsdm0_jchgrType) SetJchg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_jchgrFieldJchgMask)|(uint32(value)<<RegisterDfsdm0_jchgrFieldJchgShift))
}

// registerDfsdm1_jchgrType DFSDM injected channel group selection register
type registerDfsdm1_jchgrType uint32

const (
	RegisterDfsdm1_jchgrFieldJchgShift = 0
	RegisterDfsdm1_jchgrFieldJchgMask  = 0xff
)

// GetJchg Injected channel group selection
func (r *registerDfsdm1_jchgrType) GetJchg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_jchgrFieldJchgMask) >> RegisterDfsdm1_jchgrFieldJchgShift)
}

// SetJchg Injected channel group selection
func (r *registerDfsdm1_jchgrType) SetJchg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_jchgrFieldJchgMask)|(uint32(value)<<RegisterDfsdm1_jchgrFieldJchgShift))
}

// registerDfsdm2_jchgrType DFSDM injected channel group selection register
type registerDfsdm2_jchgrType uint32

const (
	RegisterDfsdm2_jchgrFieldJchgShift = 0
	RegisterDfsdm2_jchgrFieldJchgMask  = 0xff
)

// GetJchg Injected channel group selection
func (r *registerDfsdm2_jchgrType) GetJchg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_jchgrFieldJchgMask) >> RegisterDfsdm2_jchgrFieldJchgShift)
}

// SetJchg Injected channel group selection
func (r *registerDfsdm2_jchgrType) SetJchg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_jchgrFieldJchgMask)|(uint32(value)<<RegisterDfsdm2_jchgrFieldJchgShift))
}

// registerDfsdm3_jchgrType DFSDM injected channel group selection register
type registerDfsdm3_jchgrType uint32

const (
	RegisterDfsdm3_jchgrFieldJchgShift = 0
	RegisterDfsdm3_jchgrFieldJchgMask  = 0xff
)

// GetJchg Injected channel group selection
func (r *registerDfsdm3_jchgrType) GetJchg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_jchgrFieldJchgMask) >> RegisterDfsdm3_jchgrFieldJchgShift)
}

// SetJchg Injected channel group selection
func (r *registerDfsdm3_jchgrType) SetJchg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_jchgrFieldJchgMask)|(uint32(value)<<RegisterDfsdm3_jchgrFieldJchgShift))
}

// registerDfsdm0_fcrType DFSDM filter control register
type registerDfsdm0_fcrType uint32

const (
	RegisterDfsdm0_fcrFieldIosrShift = 0
	RegisterDfsdm0_fcrFieldIosrMask  = 0xff
)

// GetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm0_fcrType) GetIosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_fcrFieldIosrMask) >> RegisterDfsdm0_fcrFieldIosrShift)
}

// SetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm0_fcrType) SetIosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_fcrFieldIosrMask)|(uint32(value)<<RegisterDfsdm0_fcrFieldIosrShift))
}

const (
	RegisterDfsdm0_fcrFieldFosrShift = 16
	RegisterDfsdm0_fcrFieldFosrMask  = 0x3ff0000
)

// GetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm0_fcrType) GetFosr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_fcrFieldFosrMask) >> RegisterDfsdm0_fcrFieldFosrShift)
}

// SetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm0_fcrType) SetFosr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_fcrFieldFosrMask)|(uint32(value)<<RegisterDfsdm0_fcrFieldFosrShift))
}

const (
	RegisterDfsdm0_fcrFieldFordShift = 29
	RegisterDfsdm0_fcrFieldFordMask  = 0xe0000000
)

// GetFord Sinc filter order
func (r *registerDfsdm0_fcrType) GetFord() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_fcrFieldFordMask) >> RegisterDfsdm0_fcrFieldFordShift)
}

// SetFord Sinc filter order
func (r *registerDfsdm0_fcrType) SetFord(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_fcrFieldFordMask)|(uint32(value)<<RegisterDfsdm0_fcrFieldFordShift))
}

// registerDfsdm1_fcrType DFSDM filter control register
type registerDfsdm1_fcrType uint32

const (
	RegisterDfsdm1_fcrFieldIosrShift = 0
	RegisterDfsdm1_fcrFieldIosrMask  = 0xff
)

// GetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm1_fcrType) GetIosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_fcrFieldIosrMask) >> RegisterDfsdm1_fcrFieldIosrShift)
}

// SetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm1_fcrType) SetIosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_fcrFieldIosrMask)|(uint32(value)<<RegisterDfsdm1_fcrFieldIosrShift))
}

const (
	RegisterDfsdm1_fcrFieldFosrShift = 16
	RegisterDfsdm1_fcrFieldFosrMask  = 0x3ff0000
)

// GetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm1_fcrType) GetFosr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_fcrFieldFosrMask) >> RegisterDfsdm1_fcrFieldFosrShift)
}

// SetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm1_fcrType) SetFosr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_fcrFieldFosrMask)|(uint32(value)<<RegisterDfsdm1_fcrFieldFosrShift))
}

const (
	RegisterDfsdm1_fcrFieldFordShift = 29
	RegisterDfsdm1_fcrFieldFordMask  = 0xe0000000
)

// GetFord Sinc filter order
func (r *registerDfsdm1_fcrType) GetFord() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_fcrFieldFordMask) >> RegisterDfsdm1_fcrFieldFordShift)
}

// SetFord Sinc filter order
func (r *registerDfsdm1_fcrType) SetFord(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_fcrFieldFordMask)|(uint32(value)<<RegisterDfsdm1_fcrFieldFordShift))
}

// registerDfsdm2_fcrType DFSDM filter control register
type registerDfsdm2_fcrType uint32

const (
	RegisterDfsdm2_fcrFieldIosrShift = 0
	RegisterDfsdm2_fcrFieldIosrMask  = 0xff
)

// GetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm2_fcrType) GetIosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_fcrFieldIosrMask) >> RegisterDfsdm2_fcrFieldIosrShift)
}

// SetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm2_fcrType) SetIosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_fcrFieldIosrMask)|(uint32(value)<<RegisterDfsdm2_fcrFieldIosrShift))
}

const (
	RegisterDfsdm2_fcrFieldFosrShift = 16
	RegisterDfsdm2_fcrFieldFosrMask  = 0x3ff0000
)

// GetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm2_fcrType) GetFosr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_fcrFieldFosrMask) >> RegisterDfsdm2_fcrFieldFosrShift)
}

// SetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm2_fcrType) SetFosr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_fcrFieldFosrMask)|(uint32(value)<<RegisterDfsdm2_fcrFieldFosrShift))
}

const (
	RegisterDfsdm2_fcrFieldFordShift = 29
	RegisterDfsdm2_fcrFieldFordMask  = 0xe0000000
)

// GetFord Sinc filter order
func (r *registerDfsdm2_fcrType) GetFord() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_fcrFieldFordMask) >> RegisterDfsdm2_fcrFieldFordShift)
}

// SetFord Sinc filter order
func (r *registerDfsdm2_fcrType) SetFord(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_fcrFieldFordMask)|(uint32(value)<<RegisterDfsdm2_fcrFieldFordShift))
}

// registerDfsdm3_fcrType DFSDM filter control register
type registerDfsdm3_fcrType uint32

const (
	RegisterDfsdm3_fcrFieldIosrShift = 0
	RegisterDfsdm3_fcrFieldIosrMask  = 0xff
)

// GetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm3_fcrType) GetIosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_fcrFieldIosrMask) >> RegisterDfsdm3_fcrFieldIosrShift)
}

// SetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm3_fcrType) SetIosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_fcrFieldIosrMask)|(uint32(value)<<RegisterDfsdm3_fcrFieldIosrShift))
}

const (
	RegisterDfsdm3_fcrFieldFosrShift = 16
	RegisterDfsdm3_fcrFieldFosrMask  = 0x3ff0000
)

// GetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm3_fcrType) GetFosr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_fcrFieldFosrMask) >> RegisterDfsdm3_fcrFieldFosrShift)
}

// SetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm3_fcrType) SetFosr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_fcrFieldFosrMask)|(uint32(value)<<RegisterDfsdm3_fcrFieldFosrShift))
}

const (
	RegisterDfsdm3_fcrFieldFordShift = 29
	RegisterDfsdm3_fcrFieldFordMask  = 0xe0000000
)

// GetFord Sinc filter order
func (r *registerDfsdm3_fcrType) GetFord() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_fcrFieldFordMask) >> RegisterDfsdm3_fcrFieldFordShift)
}

// SetFord Sinc filter order
func (r *registerDfsdm3_fcrType) SetFord(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_fcrFieldFordMask)|(uint32(value)<<RegisterDfsdm3_fcrFieldFordShift))
}

// registerDfsdm0_jdatarType DFSDM data register for injected group
type registerDfsdm0_jdatarType uint32

const (
	RegisterDfsdm0_jdatarFieldJdatachShift = 0
	RegisterDfsdm0_jdatarFieldJdatachMask  = 0x7
)

// GetJdatach Injected channel most recently converted
func (r *registerDfsdm0_jdatarType) GetJdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_jdatarFieldJdatachMask) >> RegisterDfsdm0_jdatarFieldJdatachShift)
}

// SetJdatach Injected channel most recently converted
func (r *registerDfsdm0_jdatarType) SetJdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_jdatarFieldJdatachMask)|(uint32(value)<<RegisterDfsdm0_jdatarFieldJdatachShift))
}

const (
	RegisterDfsdm0_jdatarFieldJdataShift = 8
	RegisterDfsdm0_jdatarFieldJdataMask  = 0xffffff00
)

// GetJdata Injected group conversion data
func (r *registerDfsdm0_jdatarType) GetJdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_jdatarFieldJdataMask) >> RegisterDfsdm0_jdatarFieldJdataShift)
}

// SetJdata Injected group conversion data
func (r *registerDfsdm0_jdatarType) SetJdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_jdatarFieldJdataMask)|(uint32(value)<<RegisterDfsdm0_jdatarFieldJdataShift))
}

// registerDfsdm1_jdatarType DFSDM data register for injected group
type registerDfsdm1_jdatarType uint32

const (
	RegisterDfsdm1_jdatarFieldJdatachShift = 0
	RegisterDfsdm1_jdatarFieldJdatachMask  = 0x7
)

// GetJdatach Injected channel most recently converted
func (r *registerDfsdm1_jdatarType) GetJdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_jdatarFieldJdatachMask) >> RegisterDfsdm1_jdatarFieldJdatachShift)
}

// SetJdatach Injected channel most recently converted
func (r *registerDfsdm1_jdatarType) SetJdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_jdatarFieldJdatachMask)|(uint32(value)<<RegisterDfsdm1_jdatarFieldJdatachShift))
}

const (
	RegisterDfsdm1_jdatarFieldJdataShift = 8
	RegisterDfsdm1_jdatarFieldJdataMask  = 0xffffff00
)

// GetJdata Injected group conversion data
func (r *registerDfsdm1_jdatarType) GetJdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_jdatarFieldJdataMask) >> RegisterDfsdm1_jdatarFieldJdataShift)
}

// SetJdata Injected group conversion data
func (r *registerDfsdm1_jdatarType) SetJdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_jdatarFieldJdataMask)|(uint32(value)<<RegisterDfsdm1_jdatarFieldJdataShift))
}

// registerDfsdm2_jdatarType DFSDM data register for injected group
type registerDfsdm2_jdatarType uint32

const (
	RegisterDfsdm2_jdatarFieldJdatachShift = 0
	RegisterDfsdm2_jdatarFieldJdatachMask  = 0x7
)

// GetJdatach Injected channel most recently converted
func (r *registerDfsdm2_jdatarType) GetJdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_jdatarFieldJdatachMask) >> RegisterDfsdm2_jdatarFieldJdatachShift)
}

// SetJdatach Injected channel most recently converted
func (r *registerDfsdm2_jdatarType) SetJdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_jdatarFieldJdatachMask)|(uint32(value)<<RegisterDfsdm2_jdatarFieldJdatachShift))
}

const (
	RegisterDfsdm2_jdatarFieldJdataShift = 8
	RegisterDfsdm2_jdatarFieldJdataMask  = 0xffffff00
)

// GetJdata Injected group conversion data
func (r *registerDfsdm2_jdatarType) GetJdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_jdatarFieldJdataMask) >> RegisterDfsdm2_jdatarFieldJdataShift)
}

// SetJdata Injected group conversion data
func (r *registerDfsdm2_jdatarType) SetJdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_jdatarFieldJdataMask)|(uint32(value)<<RegisterDfsdm2_jdatarFieldJdataShift))
}

// registerDfsdm3_jdatarType DFSDM data register for injected group
type registerDfsdm3_jdatarType uint32

const (
	RegisterDfsdm3_jdatarFieldJdatachShift = 0
	RegisterDfsdm3_jdatarFieldJdatachMask  = 0x7
)

// GetJdatach Injected channel most recently converted
func (r *registerDfsdm3_jdatarType) GetJdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_jdatarFieldJdatachMask) >> RegisterDfsdm3_jdatarFieldJdatachShift)
}

// SetJdatach Injected channel most recently converted
func (r *registerDfsdm3_jdatarType) SetJdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_jdatarFieldJdatachMask)|(uint32(value)<<RegisterDfsdm3_jdatarFieldJdatachShift))
}

const (
	RegisterDfsdm3_jdatarFieldJdataShift = 8
	RegisterDfsdm3_jdatarFieldJdataMask  = 0xffffff00
)

// GetJdata Injected group conversion data
func (r *registerDfsdm3_jdatarType) GetJdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_jdatarFieldJdataMask) >> RegisterDfsdm3_jdatarFieldJdataShift)
}

// SetJdata Injected group conversion data
func (r *registerDfsdm3_jdatarType) SetJdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_jdatarFieldJdataMask)|(uint32(value)<<RegisterDfsdm3_jdatarFieldJdataShift))
}

// registerDfsdm0_rdatarType DFSDM data register for the regular channel
type registerDfsdm0_rdatarType uint32

const (
	RegisterDfsdm0_rdatarFieldRdatachShift = 0
	RegisterDfsdm0_rdatarFieldRdatachMask  = 0x7
)

// GetRdatach Regular channel most recently converted
func (r *registerDfsdm0_rdatarType) GetRdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_rdatarFieldRdatachMask) >> RegisterDfsdm0_rdatarFieldRdatachShift)
}

// SetRdatach Regular channel most recently converted
func (r *registerDfsdm0_rdatarType) SetRdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_rdatarFieldRdatachMask)|(uint32(value)<<RegisterDfsdm0_rdatarFieldRdatachShift))
}

const (
	RegisterDfsdm0_rdatarFieldRpendShift = 4
	RegisterDfsdm0_rdatarFieldRpendMask  = 0x10
)

// GetRpend Regular channel pending data
func (r *registerDfsdm0_rdatarType) GetRpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_rdatarFieldRpendMask) != 0
}

// SetRpend Regular channel pending data
func (r *registerDfsdm0_rdatarType) SetRpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm0_rdatarFieldRpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_rdatarFieldRpendMask)
	}
}

const (
	RegisterDfsdm0_rdatarFieldRdataShift = 8
	RegisterDfsdm0_rdatarFieldRdataMask  = 0xffffff00
)

// GetRdata Regular channel conversion data
func (r *registerDfsdm0_rdatarType) GetRdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_rdatarFieldRdataMask) >> RegisterDfsdm0_rdatarFieldRdataShift)
}

// SetRdata Regular channel conversion data
func (r *registerDfsdm0_rdatarType) SetRdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_rdatarFieldRdataMask)|(uint32(value)<<RegisterDfsdm0_rdatarFieldRdataShift))
}

// registerDfsdm1_rdatarType DFSDM data register for the regular channel
type registerDfsdm1_rdatarType uint32

const (
	RegisterDfsdm1_rdatarFieldRdatachShift = 0
	RegisterDfsdm1_rdatarFieldRdatachMask  = 0x7
)

// GetRdatach Regular channel most recently converted
func (r *registerDfsdm1_rdatarType) GetRdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_rdatarFieldRdatachMask) >> RegisterDfsdm1_rdatarFieldRdatachShift)
}

// SetRdatach Regular channel most recently converted
func (r *registerDfsdm1_rdatarType) SetRdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_rdatarFieldRdatachMask)|(uint32(value)<<RegisterDfsdm1_rdatarFieldRdatachShift))
}

const (
	RegisterDfsdm1_rdatarFieldRpendShift = 4
	RegisterDfsdm1_rdatarFieldRpendMask  = 0x10
)

// GetRpend Regular channel pending data
func (r *registerDfsdm1_rdatarType) GetRpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_rdatarFieldRpendMask) != 0
}

// SetRpend Regular channel pending data
func (r *registerDfsdm1_rdatarType) SetRpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1_rdatarFieldRpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_rdatarFieldRpendMask)
	}
}

const (
	RegisterDfsdm1_rdatarFieldRdataShift = 8
	RegisterDfsdm1_rdatarFieldRdataMask  = 0xffffff00
)

// GetRdata Regular channel conversion data
func (r *registerDfsdm1_rdatarType) GetRdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_rdatarFieldRdataMask) >> RegisterDfsdm1_rdatarFieldRdataShift)
}

// SetRdata Regular channel conversion data
func (r *registerDfsdm1_rdatarType) SetRdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_rdatarFieldRdataMask)|(uint32(value)<<RegisterDfsdm1_rdatarFieldRdataShift))
}

// registerDfsdm2_rdatarType DFSDM data register for the regular channel
type registerDfsdm2_rdatarType uint32

const (
	RegisterDfsdm2_rdatarFieldRdatachShift = 0
	RegisterDfsdm2_rdatarFieldRdatachMask  = 0x7
)

// GetRdatach Regular channel most recently converted
func (r *registerDfsdm2_rdatarType) GetRdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_rdatarFieldRdatachMask) >> RegisterDfsdm2_rdatarFieldRdatachShift)
}

// SetRdatach Regular channel most recently converted
func (r *registerDfsdm2_rdatarType) SetRdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_rdatarFieldRdatachMask)|(uint32(value)<<RegisterDfsdm2_rdatarFieldRdatachShift))
}

const (
	RegisterDfsdm2_rdatarFieldRpendShift = 4
	RegisterDfsdm2_rdatarFieldRpendMask  = 0x10
)

// GetRpend Regular channel pending data
func (r *registerDfsdm2_rdatarType) GetRpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_rdatarFieldRpendMask) != 0
}

// SetRpend Regular channel pending data
func (r *registerDfsdm2_rdatarType) SetRpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm2_rdatarFieldRpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_rdatarFieldRpendMask)
	}
}

const (
	RegisterDfsdm2_rdatarFieldRdataShift = 8
	RegisterDfsdm2_rdatarFieldRdataMask  = 0xffffff00
)

// GetRdata Regular channel conversion data
func (r *registerDfsdm2_rdatarType) GetRdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_rdatarFieldRdataMask) >> RegisterDfsdm2_rdatarFieldRdataShift)
}

// SetRdata Regular channel conversion data
func (r *registerDfsdm2_rdatarType) SetRdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_rdatarFieldRdataMask)|(uint32(value)<<RegisterDfsdm2_rdatarFieldRdataShift))
}

// registerDfsdm3_rdatarType DFSDM data register for the regular channel
type registerDfsdm3_rdatarType uint32

const (
	RegisterDfsdm3_rdatarFieldRdatachShift = 0
	RegisterDfsdm3_rdatarFieldRdatachMask  = 0x7
)

// GetRdatach Regular channel most recently converted
func (r *registerDfsdm3_rdatarType) GetRdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_rdatarFieldRdatachMask) >> RegisterDfsdm3_rdatarFieldRdatachShift)
}

// SetRdatach Regular channel most recently converted
func (r *registerDfsdm3_rdatarType) SetRdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_rdatarFieldRdatachMask)|(uint32(value)<<RegisterDfsdm3_rdatarFieldRdatachShift))
}

const (
	RegisterDfsdm3_rdatarFieldRpendShift = 4
	RegisterDfsdm3_rdatarFieldRpendMask  = 0x10
)

// GetRpend Regular channel pending data
func (r *registerDfsdm3_rdatarType) GetRpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_rdatarFieldRpendMask) != 0
}

// SetRpend Regular channel pending data
func (r *registerDfsdm3_rdatarType) SetRpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm3_rdatarFieldRpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_rdatarFieldRpendMask)
	}
}

const (
	RegisterDfsdm3_rdatarFieldRdataShift = 8
	RegisterDfsdm3_rdatarFieldRdataMask  = 0xffffff00
)

// GetRdata Regular channel conversion data
func (r *registerDfsdm3_rdatarType) GetRdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_rdatarFieldRdataMask) >> RegisterDfsdm3_rdatarFieldRdataShift)
}

// SetRdata Regular channel conversion data
func (r *registerDfsdm3_rdatarType) SetRdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_rdatarFieldRdataMask)|(uint32(value)<<RegisterDfsdm3_rdatarFieldRdataShift))
}

// registerDfsdm0_awhtrType DFSDM analog watchdog high threshold register
type registerDfsdm0_awhtrType uint32

const (
	RegisterDfsdm0_awhtrFieldBkawhShift = 0
	RegisterDfsdm0_awhtrFieldBkawhMask  = 0xf
)

// GetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm0_awhtrType) GetBkawh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_awhtrFieldBkawhMask) >> RegisterDfsdm0_awhtrFieldBkawhShift)
}

// SetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm0_awhtrType) SetBkawh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_awhtrFieldBkawhMask)|(uint32(value)<<RegisterDfsdm0_awhtrFieldBkawhShift))
}

const (
	RegisterDfsdm0_awhtrFieldAwhtShift = 8
	RegisterDfsdm0_awhtrFieldAwhtMask  = 0xffffff00
)

// GetAwht Analog watchdog high threshold
func (r *registerDfsdm0_awhtrType) GetAwht() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_awhtrFieldAwhtMask) >> RegisterDfsdm0_awhtrFieldAwhtShift)
}

// SetAwht Analog watchdog high threshold
func (r *registerDfsdm0_awhtrType) SetAwht(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_awhtrFieldAwhtMask)|(uint32(value)<<RegisterDfsdm0_awhtrFieldAwhtShift))
}

// registerDfsdm1_awhtrType DFSDM analog watchdog high threshold register
type registerDfsdm1_awhtrType uint32

const (
	RegisterDfsdm1_awhtrFieldBkawhShift = 0
	RegisterDfsdm1_awhtrFieldBkawhMask  = 0xf
)

// GetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm1_awhtrType) GetBkawh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_awhtrFieldBkawhMask) >> RegisterDfsdm1_awhtrFieldBkawhShift)
}

// SetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm1_awhtrType) SetBkawh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_awhtrFieldBkawhMask)|(uint32(value)<<RegisterDfsdm1_awhtrFieldBkawhShift))
}

const (
	RegisterDfsdm1_awhtrFieldAwhtShift = 8
	RegisterDfsdm1_awhtrFieldAwhtMask  = 0xffffff00
)

// GetAwht Analog watchdog high threshold
func (r *registerDfsdm1_awhtrType) GetAwht() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_awhtrFieldAwhtMask) >> RegisterDfsdm1_awhtrFieldAwhtShift)
}

// SetAwht Analog watchdog high threshold
func (r *registerDfsdm1_awhtrType) SetAwht(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_awhtrFieldAwhtMask)|(uint32(value)<<RegisterDfsdm1_awhtrFieldAwhtShift))
}

// registerDfsdm2_awhtrType DFSDM analog watchdog high threshold register
type registerDfsdm2_awhtrType uint32

const (
	RegisterDfsdm2_awhtrFieldBkawhShift = 0
	RegisterDfsdm2_awhtrFieldBkawhMask  = 0xf
)

// GetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm2_awhtrType) GetBkawh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_awhtrFieldBkawhMask) >> RegisterDfsdm2_awhtrFieldBkawhShift)
}

// SetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm2_awhtrType) SetBkawh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_awhtrFieldBkawhMask)|(uint32(value)<<RegisterDfsdm2_awhtrFieldBkawhShift))
}

const (
	RegisterDfsdm2_awhtrFieldAwhtShift = 8
	RegisterDfsdm2_awhtrFieldAwhtMask  = 0xffffff00
)

// GetAwht Analog watchdog high threshold
func (r *registerDfsdm2_awhtrType) GetAwht() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_awhtrFieldAwhtMask) >> RegisterDfsdm2_awhtrFieldAwhtShift)
}

// SetAwht Analog watchdog high threshold
func (r *registerDfsdm2_awhtrType) SetAwht(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_awhtrFieldAwhtMask)|(uint32(value)<<RegisterDfsdm2_awhtrFieldAwhtShift))
}

// registerDfsdm3_awhtrType DFSDM analog watchdog high threshold register
type registerDfsdm3_awhtrType uint32

const (
	RegisterDfsdm3_awhtrFieldBkawhShift = 0
	RegisterDfsdm3_awhtrFieldBkawhMask  = 0xf
)

// GetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm3_awhtrType) GetBkawh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_awhtrFieldBkawhMask) >> RegisterDfsdm3_awhtrFieldBkawhShift)
}

// SetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm3_awhtrType) SetBkawh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_awhtrFieldBkawhMask)|(uint32(value)<<RegisterDfsdm3_awhtrFieldBkawhShift))
}

const (
	RegisterDfsdm3_awhtrFieldAwhtShift = 8
	RegisterDfsdm3_awhtrFieldAwhtMask  = 0xffffff00
)

// GetAwht Analog watchdog high threshold
func (r *registerDfsdm3_awhtrType) GetAwht() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_awhtrFieldAwhtMask) >> RegisterDfsdm3_awhtrFieldAwhtShift)
}

// SetAwht Analog watchdog high threshold
func (r *registerDfsdm3_awhtrType) SetAwht(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_awhtrFieldAwhtMask)|(uint32(value)<<RegisterDfsdm3_awhtrFieldAwhtShift))
}

// registerDfsdm0_awltrType DFSDM analog watchdog low threshold register
type registerDfsdm0_awltrType uint32

const (
	RegisterDfsdm0_awltrFieldBkawlShift = 0
	RegisterDfsdm0_awltrFieldBkawlMask  = 0xf
)

// GetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm0_awltrType) GetBkawl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_awltrFieldBkawlMask) >> RegisterDfsdm0_awltrFieldBkawlShift)
}

// SetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm0_awltrType) SetBkawl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_awltrFieldBkawlMask)|(uint32(value)<<RegisterDfsdm0_awltrFieldBkawlShift))
}

const (
	RegisterDfsdm0_awltrFieldAwltShift = 8
	RegisterDfsdm0_awltrFieldAwltMask  = 0xffffff00
)

// GetAwlt Analog watchdog low threshold
func (r *registerDfsdm0_awltrType) GetAwlt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_awltrFieldAwltMask) >> RegisterDfsdm0_awltrFieldAwltShift)
}

// SetAwlt Analog watchdog low threshold
func (r *registerDfsdm0_awltrType) SetAwlt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_awltrFieldAwltMask)|(uint32(value)<<RegisterDfsdm0_awltrFieldAwltShift))
}

// registerDfsdm1_awltrType DFSDM analog watchdog low threshold register
type registerDfsdm1_awltrType uint32

const (
	RegisterDfsdm1_awltrFieldBkawlShift = 0
	RegisterDfsdm1_awltrFieldBkawlMask  = 0xf
)

// GetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm1_awltrType) GetBkawl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_awltrFieldBkawlMask) >> RegisterDfsdm1_awltrFieldBkawlShift)
}

// SetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm1_awltrType) SetBkawl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_awltrFieldBkawlMask)|(uint32(value)<<RegisterDfsdm1_awltrFieldBkawlShift))
}

const (
	RegisterDfsdm1_awltrFieldAwltShift = 8
	RegisterDfsdm1_awltrFieldAwltMask  = 0xffffff00
)

// GetAwlt Analog watchdog low threshold
func (r *registerDfsdm1_awltrType) GetAwlt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_awltrFieldAwltMask) >> RegisterDfsdm1_awltrFieldAwltShift)
}

// SetAwlt Analog watchdog low threshold
func (r *registerDfsdm1_awltrType) SetAwlt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_awltrFieldAwltMask)|(uint32(value)<<RegisterDfsdm1_awltrFieldAwltShift))
}

// registerDfsdm2_awltrType DFSDM analog watchdog low threshold register
type registerDfsdm2_awltrType uint32

const (
	RegisterDfsdm2_awltrFieldBkawlShift = 0
	RegisterDfsdm2_awltrFieldBkawlMask  = 0xf
)

// GetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm2_awltrType) GetBkawl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_awltrFieldBkawlMask) >> RegisterDfsdm2_awltrFieldBkawlShift)
}

// SetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm2_awltrType) SetBkawl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_awltrFieldBkawlMask)|(uint32(value)<<RegisterDfsdm2_awltrFieldBkawlShift))
}

const (
	RegisterDfsdm2_awltrFieldAwltShift = 8
	RegisterDfsdm2_awltrFieldAwltMask  = 0xffffff00
)

// GetAwlt Analog watchdog low threshold
func (r *registerDfsdm2_awltrType) GetAwlt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_awltrFieldAwltMask) >> RegisterDfsdm2_awltrFieldAwltShift)
}

// SetAwlt Analog watchdog low threshold
func (r *registerDfsdm2_awltrType) SetAwlt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_awltrFieldAwltMask)|(uint32(value)<<RegisterDfsdm2_awltrFieldAwltShift))
}

// registerDfsdm3_awltrType DFSDM analog watchdog low threshold register
type registerDfsdm3_awltrType uint32

const (
	RegisterDfsdm3_awltrFieldBkawlShift = 0
	RegisterDfsdm3_awltrFieldBkawlMask  = 0xf
)

// GetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm3_awltrType) GetBkawl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_awltrFieldBkawlMask) >> RegisterDfsdm3_awltrFieldBkawlShift)
}

// SetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm3_awltrType) SetBkawl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_awltrFieldBkawlMask)|(uint32(value)<<RegisterDfsdm3_awltrFieldBkawlShift))
}

const (
	RegisterDfsdm3_awltrFieldAwltShift = 8
	RegisterDfsdm3_awltrFieldAwltMask  = 0xffffff00
)

// GetAwlt Analog watchdog low threshold
func (r *registerDfsdm3_awltrType) GetAwlt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_awltrFieldAwltMask) >> RegisterDfsdm3_awltrFieldAwltShift)
}

// SetAwlt Analog watchdog low threshold
func (r *registerDfsdm3_awltrType) SetAwlt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_awltrFieldAwltMask)|(uint32(value)<<RegisterDfsdm3_awltrFieldAwltShift))
}

// registerDfsdm0_awsrType DFSDM analog watchdog status register
type registerDfsdm0_awsrType uint32

const (
	RegisterDfsdm0_awsrFieldAwltfShift = 0
	RegisterDfsdm0_awsrFieldAwltfMask  = 0xff
)

// GetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm0_awsrType) GetAwltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_awsrFieldAwltfMask) >> RegisterDfsdm0_awsrFieldAwltfShift)
}

// SetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm0_awsrType) SetAwltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_awsrFieldAwltfMask)|(uint32(value)<<RegisterDfsdm0_awsrFieldAwltfShift))
}

const (
	RegisterDfsdm0_awsrFieldAwhtfShift = 8
	RegisterDfsdm0_awsrFieldAwhtfMask  = 0xff00
)

// GetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm0_awsrType) GetAwhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_awsrFieldAwhtfMask) >> RegisterDfsdm0_awsrFieldAwhtfShift)
}

// SetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm0_awsrType) SetAwhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_awsrFieldAwhtfMask)|(uint32(value)<<RegisterDfsdm0_awsrFieldAwhtfShift))
}

// registerDfsdm1_awsrType DFSDM analog watchdog status register
type registerDfsdm1_awsrType uint32

const (
	RegisterDfsdm1_awsrFieldAwltfShift = 0
	RegisterDfsdm1_awsrFieldAwltfMask  = 0xff
)

// GetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm1_awsrType) GetAwltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_awsrFieldAwltfMask) >> RegisterDfsdm1_awsrFieldAwltfShift)
}

// SetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm1_awsrType) SetAwltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_awsrFieldAwltfMask)|(uint32(value)<<RegisterDfsdm1_awsrFieldAwltfShift))
}

const (
	RegisterDfsdm1_awsrFieldAwhtfShift = 8
	RegisterDfsdm1_awsrFieldAwhtfMask  = 0xff00
)

// GetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm1_awsrType) GetAwhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_awsrFieldAwhtfMask) >> RegisterDfsdm1_awsrFieldAwhtfShift)
}

// SetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm1_awsrType) SetAwhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_awsrFieldAwhtfMask)|(uint32(value)<<RegisterDfsdm1_awsrFieldAwhtfShift))
}

// registerDfsdm2_awsrType DFSDM analog watchdog status register
type registerDfsdm2_awsrType uint32

const (
	RegisterDfsdm2_awsrFieldAwltfShift = 0
	RegisterDfsdm2_awsrFieldAwltfMask  = 0xff
)

// GetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm2_awsrType) GetAwltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_awsrFieldAwltfMask) >> RegisterDfsdm2_awsrFieldAwltfShift)
}

// SetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm2_awsrType) SetAwltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_awsrFieldAwltfMask)|(uint32(value)<<RegisterDfsdm2_awsrFieldAwltfShift))
}

const (
	RegisterDfsdm2_awsrFieldAwhtfShift = 8
	RegisterDfsdm2_awsrFieldAwhtfMask  = 0xff00
)

// GetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm2_awsrType) GetAwhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_awsrFieldAwhtfMask) >> RegisterDfsdm2_awsrFieldAwhtfShift)
}

// SetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm2_awsrType) SetAwhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_awsrFieldAwhtfMask)|(uint32(value)<<RegisterDfsdm2_awsrFieldAwhtfShift))
}

// registerDfsdm3_awsrType DFSDM analog watchdog status register
type registerDfsdm3_awsrType uint32

const (
	RegisterDfsdm3_awsrFieldAwltfShift = 0
	RegisterDfsdm3_awsrFieldAwltfMask  = 0xff
)

// GetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm3_awsrType) GetAwltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_awsrFieldAwltfMask) >> RegisterDfsdm3_awsrFieldAwltfShift)
}

// SetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm3_awsrType) SetAwltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_awsrFieldAwltfMask)|(uint32(value)<<RegisterDfsdm3_awsrFieldAwltfShift))
}

const (
	RegisterDfsdm3_awsrFieldAwhtfShift = 8
	RegisterDfsdm3_awsrFieldAwhtfMask  = 0xff00
)

// GetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm3_awsrType) GetAwhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_awsrFieldAwhtfMask) >> RegisterDfsdm3_awsrFieldAwhtfShift)
}

// SetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm3_awsrType) SetAwhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_awsrFieldAwhtfMask)|(uint32(value)<<RegisterDfsdm3_awsrFieldAwhtfShift))
}

// registerDfsdm0_awcfrType DFSDM analog watchdog clear flag register
type registerDfsdm0_awcfrType uint32

const (
	RegisterDfsdm0_awcfrFieldClrawltfShift = 0
	RegisterDfsdm0_awcfrFieldClrawltfMask  = 0xff
)

// GetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm0_awcfrType) GetClrawltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_awcfrFieldClrawltfMask) >> RegisterDfsdm0_awcfrFieldClrawltfShift)
}

// SetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm0_awcfrType) SetClrawltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_awcfrFieldClrawltfMask)|(uint32(value)<<RegisterDfsdm0_awcfrFieldClrawltfShift))
}

const (
	RegisterDfsdm0_awcfrFieldClrawhtfShift = 8
	RegisterDfsdm0_awcfrFieldClrawhtfMask  = 0xff00
)

// GetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm0_awcfrType) GetClrawhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_awcfrFieldClrawhtfMask) >> RegisterDfsdm0_awcfrFieldClrawhtfShift)
}

// SetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm0_awcfrType) SetClrawhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_awcfrFieldClrawhtfMask)|(uint32(value)<<RegisterDfsdm0_awcfrFieldClrawhtfShift))
}

// registerDfsdm1_awcfrType DFSDM analog watchdog clear flag register
type registerDfsdm1_awcfrType uint32

const (
	RegisterDfsdm1_awcfrFieldClrawltfShift = 0
	RegisterDfsdm1_awcfrFieldClrawltfMask  = 0xff
)

// GetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm1_awcfrType) GetClrawltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_awcfrFieldClrawltfMask) >> RegisterDfsdm1_awcfrFieldClrawltfShift)
}

// SetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm1_awcfrType) SetClrawltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_awcfrFieldClrawltfMask)|(uint32(value)<<RegisterDfsdm1_awcfrFieldClrawltfShift))
}

const (
	RegisterDfsdm1_awcfrFieldClrawhtfShift = 8
	RegisterDfsdm1_awcfrFieldClrawhtfMask  = 0xff00
)

// GetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm1_awcfrType) GetClrawhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_awcfrFieldClrawhtfMask) >> RegisterDfsdm1_awcfrFieldClrawhtfShift)
}

// SetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm1_awcfrType) SetClrawhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_awcfrFieldClrawhtfMask)|(uint32(value)<<RegisterDfsdm1_awcfrFieldClrawhtfShift))
}

// registerDfsdm2_awcfrType DFSDM analog watchdog clear flag register
type registerDfsdm2_awcfrType uint32

const (
	RegisterDfsdm2_awcfrFieldClrawltfShift = 0
	RegisterDfsdm2_awcfrFieldClrawltfMask  = 0xff
)

// GetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm2_awcfrType) GetClrawltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_awcfrFieldClrawltfMask) >> RegisterDfsdm2_awcfrFieldClrawltfShift)
}

// SetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm2_awcfrType) SetClrawltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_awcfrFieldClrawltfMask)|(uint32(value)<<RegisterDfsdm2_awcfrFieldClrawltfShift))
}

const (
	RegisterDfsdm2_awcfrFieldClrawhtfShift = 8
	RegisterDfsdm2_awcfrFieldClrawhtfMask  = 0xff00
)

// GetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm2_awcfrType) GetClrawhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_awcfrFieldClrawhtfMask) >> RegisterDfsdm2_awcfrFieldClrawhtfShift)
}

// SetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm2_awcfrType) SetClrawhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_awcfrFieldClrawhtfMask)|(uint32(value)<<RegisterDfsdm2_awcfrFieldClrawhtfShift))
}

// registerDfsdm3_awcfrType DFSDM analog watchdog clear flag register
type registerDfsdm3_awcfrType uint32

const (
	RegisterDfsdm3_awcfrFieldClrawltfShift = 0
	RegisterDfsdm3_awcfrFieldClrawltfMask  = 0xff
)

// GetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm3_awcfrType) GetClrawltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_awcfrFieldClrawltfMask) >> RegisterDfsdm3_awcfrFieldClrawltfShift)
}

// SetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm3_awcfrType) SetClrawltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_awcfrFieldClrawltfMask)|(uint32(value)<<RegisterDfsdm3_awcfrFieldClrawltfShift))
}

const (
	RegisterDfsdm3_awcfrFieldClrawhtfShift = 8
	RegisterDfsdm3_awcfrFieldClrawhtfMask  = 0xff00
)

// GetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm3_awcfrType) GetClrawhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_awcfrFieldClrawhtfMask) >> RegisterDfsdm3_awcfrFieldClrawhtfShift)
}

// SetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm3_awcfrType) SetClrawhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_awcfrFieldClrawhtfMask)|(uint32(value)<<RegisterDfsdm3_awcfrFieldClrawhtfShift))
}

// registerDfsdm0_exmaxType DFSDM Extremes detector maximum register
type registerDfsdm0_exmaxType uint32

const (
	RegisterDfsdm0_exmaxFieldExmaxchShift = 0
	RegisterDfsdm0_exmaxFieldExmaxchMask  = 0x7
)

// GetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm0_exmaxType) GetExmaxch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_exmaxFieldExmaxchMask) >> RegisterDfsdm0_exmaxFieldExmaxchShift)
}

// SetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm0_exmaxType) SetExmaxch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_exmaxFieldExmaxchMask)|(uint32(value)<<RegisterDfsdm0_exmaxFieldExmaxchShift))
}

const (
	RegisterDfsdm0_exmaxFieldExmaxShift = 8
	RegisterDfsdm0_exmaxFieldExmaxMask  = 0xffffff00
)

// GetExmax Extremes detector maximum value
func (r *registerDfsdm0_exmaxType) GetExmax() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_exmaxFieldExmaxMask) >> RegisterDfsdm0_exmaxFieldExmaxShift)
}

// SetExmax Extremes detector maximum value
func (r *registerDfsdm0_exmaxType) SetExmax(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_exmaxFieldExmaxMask)|(uint32(value)<<RegisterDfsdm0_exmaxFieldExmaxShift))
}

// registerDfsdm1_exmaxType DFSDM Extremes detector maximum register
type registerDfsdm1_exmaxType uint32

const (
	RegisterDfsdm1_exmaxFieldExmaxchShift = 0
	RegisterDfsdm1_exmaxFieldExmaxchMask  = 0x7
)

// GetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm1_exmaxType) GetExmaxch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_exmaxFieldExmaxchMask) >> RegisterDfsdm1_exmaxFieldExmaxchShift)
}

// SetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm1_exmaxType) SetExmaxch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_exmaxFieldExmaxchMask)|(uint32(value)<<RegisterDfsdm1_exmaxFieldExmaxchShift))
}

const (
	RegisterDfsdm1_exmaxFieldExmaxShift = 8
	RegisterDfsdm1_exmaxFieldExmaxMask  = 0xffffff00
)

// GetExmax Extremes detector maximum value
func (r *registerDfsdm1_exmaxType) GetExmax() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_exmaxFieldExmaxMask) >> RegisterDfsdm1_exmaxFieldExmaxShift)
}

// SetExmax Extremes detector maximum value
func (r *registerDfsdm1_exmaxType) SetExmax(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_exmaxFieldExmaxMask)|(uint32(value)<<RegisterDfsdm1_exmaxFieldExmaxShift))
}

// registerDfsdm2_exmaxType DFSDM Extremes detector maximum register
type registerDfsdm2_exmaxType uint32

const (
	RegisterDfsdm2_exmaxFieldExmaxchShift = 0
	RegisterDfsdm2_exmaxFieldExmaxchMask  = 0x7
)

// GetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm2_exmaxType) GetExmaxch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_exmaxFieldExmaxchMask) >> RegisterDfsdm2_exmaxFieldExmaxchShift)
}

// SetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm2_exmaxType) SetExmaxch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_exmaxFieldExmaxchMask)|(uint32(value)<<RegisterDfsdm2_exmaxFieldExmaxchShift))
}

const (
	RegisterDfsdm2_exmaxFieldExmaxShift = 8
	RegisterDfsdm2_exmaxFieldExmaxMask  = 0xffffff00
)

// GetExmax Extremes detector maximum value
func (r *registerDfsdm2_exmaxType) GetExmax() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_exmaxFieldExmaxMask) >> RegisterDfsdm2_exmaxFieldExmaxShift)
}

// SetExmax Extremes detector maximum value
func (r *registerDfsdm2_exmaxType) SetExmax(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_exmaxFieldExmaxMask)|(uint32(value)<<RegisterDfsdm2_exmaxFieldExmaxShift))
}

// registerDfsdm3_exmaxType DFSDM Extremes detector maximum register
type registerDfsdm3_exmaxType uint32

const (
	RegisterDfsdm3_exmaxFieldExmaxchShift = 0
	RegisterDfsdm3_exmaxFieldExmaxchMask  = 0x7
)

// GetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm3_exmaxType) GetExmaxch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_exmaxFieldExmaxchMask) >> RegisterDfsdm3_exmaxFieldExmaxchShift)
}

// SetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm3_exmaxType) SetExmaxch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_exmaxFieldExmaxchMask)|(uint32(value)<<RegisterDfsdm3_exmaxFieldExmaxchShift))
}

const (
	RegisterDfsdm3_exmaxFieldExmaxShift = 8
	RegisterDfsdm3_exmaxFieldExmaxMask  = 0xffffff00
)

// GetExmax Extremes detector maximum value
func (r *registerDfsdm3_exmaxType) GetExmax() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_exmaxFieldExmaxMask) >> RegisterDfsdm3_exmaxFieldExmaxShift)
}

// SetExmax Extremes detector maximum value
func (r *registerDfsdm3_exmaxType) SetExmax(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_exmaxFieldExmaxMask)|(uint32(value)<<RegisterDfsdm3_exmaxFieldExmaxShift))
}

// registerDfsdm0_exminType DFSDM Extremes detector minimum register
type registerDfsdm0_exminType uint32

const (
	RegisterDfsdm0_exminFieldExminchShift = 0
	RegisterDfsdm0_exminFieldExminchMask  = 0x7
)

// GetExminch Extremes detector minimum data channel
func (r *registerDfsdm0_exminType) GetExminch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_exminFieldExminchMask) >> RegisterDfsdm0_exminFieldExminchShift)
}

// SetExminch Extremes detector minimum data channel
func (r *registerDfsdm0_exminType) SetExminch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_exminFieldExminchMask)|(uint32(value)<<RegisterDfsdm0_exminFieldExminchShift))
}

const (
	RegisterDfsdm0_exminFieldExminShift = 8
	RegisterDfsdm0_exminFieldExminMask  = 0xffffff00
)

// GetExmin Extremes detector minimum value
func (r *registerDfsdm0_exminType) GetExmin() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_exminFieldExminMask) >> RegisterDfsdm0_exminFieldExminShift)
}

// SetExmin Extremes detector minimum value
func (r *registerDfsdm0_exminType) SetExmin(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_exminFieldExminMask)|(uint32(value)<<RegisterDfsdm0_exminFieldExminShift))
}

// registerDfsdm1_exminType DFSDM Extremes detector minimum register
type registerDfsdm1_exminType uint32

const (
	RegisterDfsdm1_exminFieldExminchShift = 0
	RegisterDfsdm1_exminFieldExminchMask  = 0x7
)

// GetExminch Extremes detector minimum data channel
func (r *registerDfsdm1_exminType) GetExminch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_exminFieldExminchMask) >> RegisterDfsdm1_exminFieldExminchShift)
}

// SetExminch Extremes detector minimum data channel
func (r *registerDfsdm1_exminType) SetExminch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_exminFieldExminchMask)|(uint32(value)<<RegisterDfsdm1_exminFieldExminchShift))
}

const (
	RegisterDfsdm1_exminFieldExminShift = 8
	RegisterDfsdm1_exminFieldExminMask  = 0xffffff00
)

// GetExmin Extremes detector minimum value
func (r *registerDfsdm1_exminType) GetExmin() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_exminFieldExminMask) >> RegisterDfsdm1_exminFieldExminShift)
}

// SetExmin Extremes detector minimum value
func (r *registerDfsdm1_exminType) SetExmin(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_exminFieldExminMask)|(uint32(value)<<RegisterDfsdm1_exminFieldExminShift))
}

// registerDfsdm2_exminType DFSDM Extremes detector minimum register
type registerDfsdm2_exminType uint32

const (
	RegisterDfsdm2_exminFieldExminchShift = 0
	RegisterDfsdm2_exminFieldExminchMask  = 0x7
)

// GetExminch Extremes detector minimum data channel
func (r *registerDfsdm2_exminType) GetExminch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_exminFieldExminchMask) >> RegisterDfsdm2_exminFieldExminchShift)
}

// SetExminch Extremes detector minimum data channel
func (r *registerDfsdm2_exminType) SetExminch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_exminFieldExminchMask)|(uint32(value)<<RegisterDfsdm2_exminFieldExminchShift))
}

const (
	RegisterDfsdm2_exminFieldExminShift = 8
	RegisterDfsdm2_exminFieldExminMask  = 0xffffff00
)

// GetExmin Extremes detector minimum value
func (r *registerDfsdm2_exminType) GetExmin() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_exminFieldExminMask) >> RegisterDfsdm2_exminFieldExminShift)
}

// SetExmin Extremes detector minimum value
func (r *registerDfsdm2_exminType) SetExmin(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_exminFieldExminMask)|(uint32(value)<<RegisterDfsdm2_exminFieldExminShift))
}

// registerDfsdm3_exminType DFSDM Extremes detector minimum register
type registerDfsdm3_exminType uint32

const (
	RegisterDfsdm3_exminFieldExminchShift = 0
	RegisterDfsdm3_exminFieldExminchMask  = 0x7
)

// GetExminch Extremes detector minimum data channel
func (r *registerDfsdm3_exminType) GetExminch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_exminFieldExminchMask) >> RegisterDfsdm3_exminFieldExminchShift)
}

// SetExminch Extremes detector minimum data channel
func (r *registerDfsdm3_exminType) SetExminch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_exminFieldExminchMask)|(uint32(value)<<RegisterDfsdm3_exminFieldExminchShift))
}

const (
	RegisterDfsdm3_exminFieldExminShift = 8
	RegisterDfsdm3_exminFieldExminMask  = 0xffffff00
)

// GetExmin Extremes detector minimum value
func (r *registerDfsdm3_exminType) GetExmin() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_exminFieldExminMask) >> RegisterDfsdm3_exminFieldExminShift)
}

// SetExmin Extremes detector minimum value
func (r *registerDfsdm3_exminType) SetExmin(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_exminFieldExminMask)|(uint32(value)<<RegisterDfsdm3_exminFieldExminShift))
}

// registerDfsdm0_cnvtimrType DFSDM conversion timer register
type registerDfsdm0_cnvtimrType uint32

const (
	RegisterDfsdm0_cnvtimrFieldCnvcntShift = 4
	RegisterDfsdm0_cnvtimrFieldCnvcntMask  = 0xfffffff0
)

// GetCnvcnt 28-bit timer counting conversion time
func (r *registerDfsdm0_cnvtimrType) GetCnvcnt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm0_cnvtimrFieldCnvcntMask) >> RegisterDfsdm0_cnvtimrFieldCnvcntShift)
}

// SetCnvcnt 28-bit timer counting conversion time
func (r *registerDfsdm0_cnvtimrType) SetCnvcnt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm0_cnvtimrFieldCnvcntMask)|(uint32(value)<<RegisterDfsdm0_cnvtimrFieldCnvcntShift))
}

// registerDfsdm1_cnvtimrType DFSDM conversion timer register
type registerDfsdm1_cnvtimrType uint32

const (
	RegisterDfsdm1_cnvtimrFieldCnvcntShift = 4
	RegisterDfsdm1_cnvtimrFieldCnvcntMask  = 0xfffffff0
)

// GetCnvcnt 28-bit timer counting conversion time
func (r *registerDfsdm1_cnvtimrType) GetCnvcnt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1_cnvtimrFieldCnvcntMask) >> RegisterDfsdm1_cnvtimrFieldCnvcntShift)
}

// SetCnvcnt 28-bit timer counting conversion time
func (r *registerDfsdm1_cnvtimrType) SetCnvcnt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1_cnvtimrFieldCnvcntMask)|(uint32(value)<<RegisterDfsdm1_cnvtimrFieldCnvcntShift))
}

// registerDfsdm2_cnvtimrType DFSDM conversion timer register
type registerDfsdm2_cnvtimrType uint32

const (
	RegisterDfsdm2_cnvtimrFieldCnvcntShift = 4
	RegisterDfsdm2_cnvtimrFieldCnvcntMask  = 0xfffffff0
)

// GetCnvcnt 28-bit timer counting conversion time
func (r *registerDfsdm2_cnvtimrType) GetCnvcnt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm2_cnvtimrFieldCnvcntMask) >> RegisterDfsdm2_cnvtimrFieldCnvcntShift)
}

// SetCnvcnt 28-bit timer counting conversion time
func (r *registerDfsdm2_cnvtimrType) SetCnvcnt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm2_cnvtimrFieldCnvcntMask)|(uint32(value)<<RegisterDfsdm2_cnvtimrFieldCnvcntShift))
}

// registerDfsdm3_cnvtimrType DFSDM conversion timer register
type registerDfsdm3_cnvtimrType uint32

const (
	RegisterDfsdm3_cnvtimrFieldCnvcntShift = 4
	RegisterDfsdm3_cnvtimrFieldCnvcntMask  = 0xfffffff0
)

// GetCnvcnt 28-bit timer counting conversion time
func (r *registerDfsdm3_cnvtimrType) GetCnvcnt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm3_cnvtimrFieldCnvcntMask) >> RegisterDfsdm3_cnvtimrFieldCnvcntShift)
}

// SetCnvcnt 28-bit timer counting conversion time
func (r *registerDfsdm3_cnvtimrType) SetCnvcnt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm3_cnvtimrFieldCnvcntMask)|(uint32(value)<<RegisterDfsdm3_cnvtimrFieldCnvcntShift))
}
