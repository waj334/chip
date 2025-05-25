//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package dfsdm

import (
	"unsafe"
	"volatile"
)

var (
	Dfsdm = (*_dfsdm)(unsafe.Pointer(uintptr(0x40017000)))
)

type _dfsdm struct {
	Ch0cfgr1          registerCh0cfgr1Type
	Ch0cfgr2          registerCh0cfgr2Type
	Ch0awscdr         registerCh0awscdrType
	Ch0wdatr          registerCh0wdatrType
	Ch0datinr         registerCh0datinrType
	Ch0dlyr           registerCh0dlyrType
	_                 [8]byte
	Ch1cfgr1          registerCh1cfgr1Type
	Ch1cfgr2          registerCh1cfgr2Type
	Ch1awscdr         registerCh1awscdrType
	Ch1wdatr          registerCh1wdatrType
	Ch1datinr         registerCh1datinrType
	Ch1dlyr           registerCh1dlyrType
	_                 [8]byte
	Ch2cfgr1          registerCh2cfgr1Type
	Ch2cfgr2          registerCh2cfgr2Type
	Ch2awscdr         registerCh2awscdrType
	Ch2wdatr          registerCh2wdatrType
	Ch2datinr         registerCh2datinrType
	Ch2dlyr           registerCh2dlyrType
	_                 [8]byte
	Ch3cfgr1          registerCh3cfgr1Type
	Ch3cfgr2          registerCh3cfgr2Type
	Ch3awscdr         registerCh3awscdrType
	Ch3wdatr          registerCh3wdatrType
	Ch3datinr         registerCh3datinrType
	Ch3dlyr           registerCh3dlyrType
	_                 [8]byte
	Ch4cfgr1          registerCh4cfgr1Type
	Ch4cfgr2          registerCh4cfgr2Type
	Ch4awscdr         registerCh4awscdrType
	Ch4wdatr          registerCh4wdatrType
	Ch4datinr         registerCh4datinrType
	Ch4dlyr           registerCh4dlyrType
	_                 [8]byte
	Ch5cfgr1          registerCh5cfgr1Type
	Ch5cfgr2          registerCh5cfgr2Type
	Ch5awscdr         registerCh5awscdrType
	Ch5wdatr          registerCh5wdatrType
	Ch5datinr         registerCh5datinrType
	Ch5dlyr           registerCh5dlyrType
	_                 [8]byte
	Ch6cfgr1          registerCh6cfgr1Type
	Ch6cfgr2          registerCh6cfgr2Type
	Ch6awscdr         registerCh6awscdrType
	Ch6wdatr          registerCh6wdatrType
	Ch6datinr         registerCh6datinrType
	Ch6dlyr           registerCh6dlyrType
	_                 [8]byte
	Ch7cfgr1          registerCh7cfgr1Type
	Ch7cfgr2          registerCh7cfgr2Type
	Ch7awscdr         registerCh7awscdrType
	Ch7wdatr          registerCh7wdatrType
	Ch7datinr         registerCh7datinrType
	Ch7dlyr           registerCh7dlyrType
	_                 [8]byte
	Dfsdm_flt0cr1     registerDfsdm_flt0cr1Type
	Dfsdm_flt0cr2     registerDfsdm_flt0cr2Type
	Dfsdm_flt0isr     registerDfsdm_flt0isrType
	Dfsdm_flt0icr     registerDfsdm_flt0icrType
	Dfsdm_flt0jchgr   registerDfsdm_flt0jchgrType
	Dfsdm_flt0fcr     registerDfsdm_flt0fcrType
	Dfsdm_flt0jdatar  registerDfsdm_flt0jdatarType
	Dfsdm_flt0rdatar  registerDfsdm_flt0rdatarType
	Dfsdm_flt0awhtr   registerDfsdm_flt0awhtrType
	Dfsdm_flt0awltr   registerDfsdm_flt0awltrType
	Dfsdm_flt0awsr    registerDfsdm_flt0awsrType
	Dfsdm_flt0awcfr   registerDfsdm_flt0awcfrType
	Dfsdm_flt0exmax   registerDfsdm_flt0exmaxType
	Dfsdm_flt0exmin   registerDfsdm_flt0exminType
	Dfsdm_flt0cnvtimr registerDfsdm_flt0cnvtimrType
	_                 [68]byte
	Dfsdm_flt1cr1     registerDfsdm_flt1cr1Type
	Dfsdm_flt1cr2     registerDfsdm_flt1cr2Type
	Dfsdm_flt1isr     registerDfsdm_flt1isrType
	Dfsdm1_icr        registerDfsdm1_icrType
	Dfsdm_flt1jchgr   registerDfsdm_flt1jchgrType
	Dfsdm1_fcr        registerDfsdm1_fcrType
	Dfsdm_flt1jdatar  registerDfsdm_flt1jdatarType
	Dfsdm_flt1rdatar  registerDfsdm_flt1rdatarType
	Dfsdm_flt1awhtr   registerDfsdm_flt1awhtrType
	Dfsdm_flt1awltr   registerDfsdm_flt1awltrType
	Dfsdm_flt1awsr    registerDfsdm_flt1awsrType
	Dfsdm_flt1awcfr   registerDfsdm_flt1awcfrType
	Dfsdm_flt1exmax   registerDfsdm_flt1exmaxType
	Dfsdm_flt1exmin   registerDfsdm_flt1exminType
	Dfsdm_flt1cnvtimr registerDfsdm_flt1cnvtimrType
	_                 [68]byte
	Dfsdm_flt2cr1     registerDfsdm_flt2cr1Type
	Dfsdm_flt2cr2     registerDfsdm_flt2cr2Type
	Dfsdm_flt2isr     registerDfsdm_flt2isrType
	Dfsdm_flt2icr     registerDfsdm_flt2icrType
	Dfsdm_flt2jchgr   registerDfsdm_flt2jchgrType
	Dfsdm_flt2fcr     registerDfsdm_flt2fcrType
	Dfsdm_flt2jdatar  registerDfsdm_flt2jdatarType
	Dfsdm_flt2rdatar  registerDfsdm_flt2rdatarType
	Dfsdm_flt2awhtr   registerDfsdm_flt2awhtrType
	Dfsdm_flt2awltr   registerDfsdm_flt2awltrType
	Dfsdm_flt2awsr    registerDfsdm_flt2awsrType
	Dfsdm_flt2awcfr   registerDfsdm_flt2awcfrType
	Dfsdm_flt2exmax   registerDfsdm_flt2exmaxType
	Dfsdm_flt2exmin   registerDfsdm_flt2exminType
	Dfsdm_flt2cnvtimr registerDfsdm_flt2cnvtimrType
	_                 [68]byte
	Dfsdm_flt3cr1     registerDfsdm_flt3cr1Type
	Dfsdm_flt3cr2     registerDfsdm_flt3cr2Type
	Dfsdm_flt3isr     registerDfsdm_flt3isrType
	Dfsdm_flt3icr     registerDfsdm_flt3icrType
	Dfsdm_flt3jchgr   registerDfsdm_flt3jchgrType
	Dfsdm_flt3fcr     registerDfsdm_flt3fcrType
	Dfsdm_flt3jdatar  registerDfsdm_flt3jdatarType
	Dfsdm_flt3rdatar  registerDfsdm_flt3rdatarType
	Dfsdm_flt3awhtr   registerDfsdm_flt3awhtrType
	Dfsdm_flt3awltr   registerDfsdm_flt3awltrType
	Dfsdm_flt3awsr    registerDfsdm_flt3awsrType
	Dfsdm_flt3awcfr   registerDfsdm_flt3awcfrType
	Dfsdm_flt3exmax   registerDfsdm_flt3exmaxType
	Dfsdm_flt3exmin   registerDfsdm_flt3exminType
	Dfsdm_flt3cnvtimr registerDfsdm_flt3cnvtimrType
}

// registerCh0cfgr1Type channel configuration y register
type registerCh0cfgr1Type uint32

const (
	RegisterCh0cfgr1FieldSitpShift = 0
	RegisterCh0cfgr1FieldSitpMask  = 0x3
)

// GetSitp SITP
func (r *registerCh0cfgr1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr1FieldSitpMask) >> RegisterCh0cfgr1FieldSitpShift)
}

// SetSitp SITP
func (r *registerCh0cfgr1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr1FieldSitpMask)|(uint32(value)<<RegisterCh0cfgr1FieldSitpShift))
}

const (
	RegisterCh0cfgr1FieldSpickselShift = 2
	RegisterCh0cfgr1FieldSpickselMask  = 0xc
)

// GetSpicksel SPICKSEL
func (r *registerCh0cfgr1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr1FieldSpickselMask) >> RegisterCh0cfgr1FieldSpickselShift)
}

// SetSpicksel SPICKSEL
func (r *registerCh0cfgr1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr1FieldSpickselMask)|(uint32(value)<<RegisterCh0cfgr1FieldSpickselShift))
}

const (
	RegisterCh0cfgr1FieldScdenShift = 5
	RegisterCh0cfgr1FieldScdenMask  = 0x20
)

// GetScden SCDEN
func (r *registerCh0cfgr1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr1FieldScdenMask) != 0
}

// SetScden SCDEN
func (r *registerCh0cfgr1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh0cfgr1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr1FieldScdenMask)
	}
}

const (
	RegisterCh0cfgr1FieldCkabenShift = 6
	RegisterCh0cfgr1FieldCkabenMask  = 0x40
)

// GetCkaben CKABEN
func (r *registerCh0cfgr1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr1FieldCkabenMask) != 0
}

// SetCkaben CKABEN
func (r *registerCh0cfgr1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh0cfgr1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr1FieldCkabenMask)
	}
}

const (
	RegisterCh0cfgr1FieldChenShift = 7
	RegisterCh0cfgr1FieldChenMask  = 0x80
)

// GetChen CHEN
func (r *registerCh0cfgr1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr1FieldChenMask) != 0
}

// SetChen CHEN
func (r *registerCh0cfgr1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh0cfgr1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr1FieldChenMask)
	}
}

const (
	RegisterCh0cfgr1FieldChinselShift = 8
	RegisterCh0cfgr1FieldChinselMask  = 0x100
)

// GetChinsel CHINSEL
func (r *registerCh0cfgr1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr1FieldChinselMask) != 0
}

// SetChinsel CHINSEL
func (r *registerCh0cfgr1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh0cfgr1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr1FieldChinselMask)
	}
}

const (
	RegisterCh0cfgr1FieldDatmpxShift = 12
	RegisterCh0cfgr1FieldDatmpxMask  = 0x3000
)

// GetDatmpx DATMPX
func (r *registerCh0cfgr1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr1FieldDatmpxMask) >> RegisterCh0cfgr1FieldDatmpxShift)
}

// SetDatmpx DATMPX
func (r *registerCh0cfgr1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr1FieldDatmpxMask)|(uint32(value)<<RegisterCh0cfgr1FieldDatmpxShift))
}

const (
	RegisterCh0cfgr1FieldDatpackShift = 14
	RegisterCh0cfgr1FieldDatpackMask  = 0xc000
)

// GetDatpack DATPACK
func (r *registerCh0cfgr1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr1FieldDatpackMask) >> RegisterCh0cfgr1FieldDatpackShift)
}

// SetDatpack DATPACK
func (r *registerCh0cfgr1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr1FieldDatpackMask)|(uint32(value)<<RegisterCh0cfgr1FieldDatpackShift))
}

const (
	RegisterCh0cfgr1FieldCkoutdivShift = 16
	RegisterCh0cfgr1FieldCkoutdivMask  = 0xff0000
)

// GetCkoutdiv CKOUTDIV
func (r *registerCh0cfgr1Type) GetCkoutdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr1FieldCkoutdivMask) >> RegisterCh0cfgr1FieldCkoutdivShift)
}

// SetCkoutdiv CKOUTDIV
func (r *registerCh0cfgr1Type) SetCkoutdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr1FieldCkoutdivMask)|(uint32(value)<<RegisterCh0cfgr1FieldCkoutdivShift))
}

const (
	RegisterCh0cfgr1FieldCkoutsrcShift = 30
	RegisterCh0cfgr1FieldCkoutsrcMask  = 0x40000000
)

// GetCkoutsrc CKOUTSRC
func (r *registerCh0cfgr1Type) GetCkoutsrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr1FieldCkoutsrcMask) != 0
}

// SetCkoutsrc CKOUTSRC
func (r *registerCh0cfgr1Type) SetCkoutsrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh0cfgr1FieldCkoutsrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr1FieldCkoutsrcMask)
	}
}

const (
	RegisterCh0cfgr1FieldDfsdmenShift = 31
	RegisterCh0cfgr1FieldDfsdmenMask  = 0x80000000
)

// GetDfsdmen DFSDMEN
func (r *registerCh0cfgr1Type) GetDfsdmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr1FieldDfsdmenMask) != 0
}

// SetDfsdmen DFSDMEN
func (r *registerCh0cfgr1Type) SetDfsdmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh0cfgr1FieldDfsdmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr1FieldDfsdmenMask)
	}
}

// registerCh0cfgr2Type channel configuration y register
type registerCh0cfgr2Type uint32

const (
	RegisterCh0cfgr2FieldDtrbsShift = 3
	RegisterCh0cfgr2FieldDtrbsMask  = 0xf8
)

// GetDtrbs DTRBS
func (r *registerCh0cfgr2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr2FieldDtrbsMask) >> RegisterCh0cfgr2FieldDtrbsShift)
}

// SetDtrbs DTRBS
func (r *registerCh0cfgr2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr2FieldDtrbsMask)|(uint32(value)<<RegisterCh0cfgr2FieldDtrbsShift))
}

const (
	RegisterCh0cfgr2FieldOffsetShift = 8
	RegisterCh0cfgr2FieldOffsetMask  = 0xffffff00
)

// GetOffset OFFSET
func (r *registerCh0cfgr2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCh0cfgr2FieldOffsetMask) >> RegisterCh0cfgr2FieldOffsetShift)
}

// SetOffset OFFSET
func (r *registerCh0cfgr2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0cfgr2FieldOffsetMask)|(uint32(value)<<RegisterCh0cfgr2FieldOffsetShift))
}

// registerCh0awscdrType analog watchdog and short-circuit detector register
type registerCh0awscdrType uint32

const (
	RegisterCh0awscdrFieldScdtShift = 0
	RegisterCh0awscdrFieldScdtMask  = 0xff
)

// GetScdt SCDT
func (r *registerCh0awscdrType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh0awscdrFieldScdtMask) >> RegisterCh0awscdrFieldScdtShift)
}

// SetScdt SCDT
func (r *registerCh0awscdrType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0awscdrFieldScdtMask)|(uint32(value)<<RegisterCh0awscdrFieldScdtShift))
}

const (
	RegisterCh0awscdrFieldBkscdShift = 12
	RegisterCh0awscdrFieldBkscdMask  = 0xf000
)

// GetBkscd BKSCD
func (r *registerCh0awscdrType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh0awscdrFieldBkscdMask) >> RegisterCh0awscdrFieldBkscdShift)
}

// SetBkscd BKSCD
func (r *registerCh0awscdrType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0awscdrFieldBkscdMask)|(uint32(value)<<RegisterCh0awscdrFieldBkscdShift))
}

const (
	RegisterCh0awscdrFieldAwfosrShift = 16
	RegisterCh0awscdrFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr AWFOSR
func (r *registerCh0awscdrType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh0awscdrFieldAwfosrMask) >> RegisterCh0awscdrFieldAwfosrShift)
}

// SetAwfosr AWFOSR
func (r *registerCh0awscdrType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0awscdrFieldAwfosrMask)|(uint32(value)<<RegisterCh0awscdrFieldAwfosrShift))
}

const (
	RegisterCh0awscdrFieldAwfordShift = 22
	RegisterCh0awscdrFieldAwfordMask  = 0xc00000
)

// GetAwford AWFORD
func (r *registerCh0awscdrType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh0awscdrFieldAwfordMask) >> RegisterCh0awscdrFieldAwfordShift)
}

// SetAwford AWFORD
func (r *registerCh0awscdrType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0awscdrFieldAwfordMask)|(uint32(value)<<RegisterCh0awscdrFieldAwfordShift))
}

// registerCh0wdatrType channel watchdog filter data register
type registerCh0wdatrType uint32

const (
	RegisterCh0wdatrFieldWdataShift = 0
	RegisterCh0wdatrFieldWdataMask  = 0xffff
)

// GetWdata WDATA
func (r *registerCh0wdatrType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh0wdatrFieldWdataMask) >> RegisterCh0wdatrFieldWdataShift)
}

// SetWdata WDATA
func (r *registerCh0wdatrType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0wdatrFieldWdataMask)|(uint32(value)<<RegisterCh0wdatrFieldWdataShift))
}

// registerCh0datinrType channel data input register
type registerCh0datinrType uint32

const (
	RegisterCh0datinrFieldIndat0Shift = 0
	RegisterCh0datinrFieldIndat0Mask  = 0xffff
)

// GetIndat0 INDAT0
func (r *registerCh0datinrType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh0datinrFieldIndat0Mask) >> RegisterCh0datinrFieldIndat0Shift)
}

// SetIndat0 INDAT0
func (r *registerCh0datinrType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0datinrFieldIndat0Mask)|(uint32(value)<<RegisterCh0datinrFieldIndat0Shift))
}

const (
	RegisterCh0datinrFieldIndat1Shift = 16
	RegisterCh0datinrFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 INDAT1
func (r *registerCh0datinrType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh0datinrFieldIndat1Mask) >> RegisterCh0datinrFieldIndat1Shift)
}

// SetIndat1 INDAT1
func (r *registerCh0datinrType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0datinrFieldIndat1Mask)|(uint32(value)<<RegisterCh0datinrFieldIndat1Shift))
}

// registerCh0dlyrType channel y delay register
type registerCh0dlyrType uint32

const (
	RegisterCh0dlyrFieldPlsskpShift = 0
	RegisterCh0dlyrFieldPlsskpMask  = 0x3f
)

// GetPlsskp PLSSKP
func (r *registerCh0dlyrType) GetPlsskp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh0dlyrFieldPlsskpMask) >> RegisterCh0dlyrFieldPlsskpShift)
}

// SetPlsskp PLSSKP
func (r *registerCh0dlyrType) SetPlsskp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh0dlyrFieldPlsskpMask)|(uint32(value)<<RegisterCh0dlyrFieldPlsskpShift))
}

// registerCh1cfgr1Type CH1CFGR1
type registerCh1cfgr1Type uint32

const (
	RegisterCh1cfgr1FieldSitpShift = 0
	RegisterCh1cfgr1FieldSitpMask  = 0x3
)

// GetSitp SITP
func (r *registerCh1cfgr1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh1cfgr1FieldSitpMask) >> RegisterCh1cfgr1FieldSitpShift)
}

// SetSitp SITP
func (r *registerCh1cfgr1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1cfgr1FieldSitpMask)|(uint32(value)<<RegisterCh1cfgr1FieldSitpShift))
}

const (
	RegisterCh1cfgr1FieldSpickselShift = 2
	RegisterCh1cfgr1FieldSpickselMask  = 0xc
)

// GetSpicksel SPICKSEL
func (r *registerCh1cfgr1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh1cfgr1FieldSpickselMask) >> RegisterCh1cfgr1FieldSpickselShift)
}

// SetSpicksel SPICKSEL
func (r *registerCh1cfgr1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1cfgr1FieldSpickselMask)|(uint32(value)<<RegisterCh1cfgr1FieldSpickselShift))
}

const (
	RegisterCh1cfgr1FieldScdenShift = 5
	RegisterCh1cfgr1FieldScdenMask  = 0x20
)

// GetScden SCDEN
func (r *registerCh1cfgr1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh1cfgr1FieldScdenMask) != 0
}

// SetScden SCDEN
func (r *registerCh1cfgr1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh1cfgr1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh1cfgr1FieldScdenMask)
	}
}

const (
	RegisterCh1cfgr1FieldCkabenShift = 6
	RegisterCh1cfgr1FieldCkabenMask  = 0x40
)

// GetCkaben CKABEN
func (r *registerCh1cfgr1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh1cfgr1FieldCkabenMask) != 0
}

// SetCkaben CKABEN
func (r *registerCh1cfgr1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh1cfgr1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh1cfgr1FieldCkabenMask)
	}
}

const (
	RegisterCh1cfgr1FieldChenShift = 7
	RegisterCh1cfgr1FieldChenMask  = 0x80
)

// GetChen CHEN
func (r *registerCh1cfgr1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh1cfgr1FieldChenMask) != 0
}

// SetChen CHEN
func (r *registerCh1cfgr1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh1cfgr1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh1cfgr1FieldChenMask)
	}
}

const (
	RegisterCh1cfgr1FieldChinselShift = 8
	RegisterCh1cfgr1FieldChinselMask  = 0x100
)

// GetChinsel CHINSEL
func (r *registerCh1cfgr1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh1cfgr1FieldChinselMask) != 0
}

// SetChinsel CHINSEL
func (r *registerCh1cfgr1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh1cfgr1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh1cfgr1FieldChinselMask)
	}
}

const (
	RegisterCh1cfgr1FieldDatmpxShift = 12
	RegisterCh1cfgr1FieldDatmpxMask  = 0x3000
)

// GetDatmpx DATMPX
func (r *registerCh1cfgr1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh1cfgr1FieldDatmpxMask) >> RegisterCh1cfgr1FieldDatmpxShift)
}

// SetDatmpx DATMPX
func (r *registerCh1cfgr1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1cfgr1FieldDatmpxMask)|(uint32(value)<<RegisterCh1cfgr1FieldDatmpxShift))
}

const (
	RegisterCh1cfgr1FieldDatpackShift = 14
	RegisterCh1cfgr1FieldDatpackMask  = 0xc000
)

// GetDatpack DATPACK
func (r *registerCh1cfgr1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh1cfgr1FieldDatpackMask) >> RegisterCh1cfgr1FieldDatpackShift)
}

// SetDatpack DATPACK
func (r *registerCh1cfgr1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1cfgr1FieldDatpackMask)|(uint32(value)<<RegisterCh1cfgr1FieldDatpackShift))
}

// registerCh1cfgr2Type CH1CFGR2
type registerCh1cfgr2Type uint32

const (
	RegisterCh1cfgr2FieldDtrbsShift = 3
	RegisterCh1cfgr2FieldDtrbsMask  = 0xf8
)

// GetDtrbs DTRBS
func (r *registerCh1cfgr2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh1cfgr2FieldDtrbsMask) >> RegisterCh1cfgr2FieldDtrbsShift)
}

// SetDtrbs DTRBS
func (r *registerCh1cfgr2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1cfgr2FieldDtrbsMask)|(uint32(value)<<RegisterCh1cfgr2FieldDtrbsShift))
}

const (
	RegisterCh1cfgr2FieldOffsetShift = 8
	RegisterCh1cfgr2FieldOffsetMask  = 0xffffff00
)

// GetOffset OFFSET
func (r *registerCh1cfgr2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCh1cfgr2FieldOffsetMask) >> RegisterCh1cfgr2FieldOffsetShift)
}

// SetOffset OFFSET
func (r *registerCh1cfgr2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1cfgr2FieldOffsetMask)|(uint32(value)<<RegisterCh1cfgr2FieldOffsetShift))
}

// registerCh1awscdrType CH1AWSCDR
type registerCh1awscdrType uint32

const (
	RegisterCh1awscdrFieldScdtShift = 0
	RegisterCh1awscdrFieldScdtMask  = 0xff
)

// GetScdt SCDT
func (r *registerCh1awscdrType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh1awscdrFieldScdtMask) >> RegisterCh1awscdrFieldScdtShift)
}

// SetScdt SCDT
func (r *registerCh1awscdrType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1awscdrFieldScdtMask)|(uint32(value)<<RegisterCh1awscdrFieldScdtShift))
}

const (
	RegisterCh1awscdrFieldBkscdShift = 12
	RegisterCh1awscdrFieldBkscdMask  = 0xf000
)

// GetBkscd BKSCD
func (r *registerCh1awscdrType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh1awscdrFieldBkscdMask) >> RegisterCh1awscdrFieldBkscdShift)
}

// SetBkscd BKSCD
func (r *registerCh1awscdrType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1awscdrFieldBkscdMask)|(uint32(value)<<RegisterCh1awscdrFieldBkscdShift))
}

const (
	RegisterCh1awscdrFieldAwfosrShift = 16
	RegisterCh1awscdrFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr AWFOSR
func (r *registerCh1awscdrType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh1awscdrFieldAwfosrMask) >> RegisterCh1awscdrFieldAwfosrShift)
}

// SetAwfosr AWFOSR
func (r *registerCh1awscdrType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1awscdrFieldAwfosrMask)|(uint32(value)<<RegisterCh1awscdrFieldAwfosrShift))
}

const (
	RegisterCh1awscdrFieldAwfordShift = 22
	RegisterCh1awscdrFieldAwfordMask  = 0xc00000
)

// GetAwford AWFORD
func (r *registerCh1awscdrType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh1awscdrFieldAwfordMask) >> RegisterCh1awscdrFieldAwfordShift)
}

// SetAwford AWFORD
func (r *registerCh1awscdrType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1awscdrFieldAwfordMask)|(uint32(value)<<RegisterCh1awscdrFieldAwfordShift))
}

// registerCh1wdatrType CH1WDATR
type registerCh1wdatrType uint32

const (
	RegisterCh1wdatrFieldWdataShift = 0
	RegisterCh1wdatrFieldWdataMask  = 0xffff
)

// GetWdata WDATA
func (r *registerCh1wdatrType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh1wdatrFieldWdataMask) >> RegisterCh1wdatrFieldWdataShift)
}

// SetWdata WDATA
func (r *registerCh1wdatrType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1wdatrFieldWdataMask)|(uint32(value)<<RegisterCh1wdatrFieldWdataShift))
}

// registerCh1datinrType CH1DATINR
type registerCh1datinrType uint32

const (
	RegisterCh1datinrFieldIndat0Shift = 0
	RegisterCh1datinrFieldIndat0Mask  = 0xffff
)

// GetIndat0 INDAT0
func (r *registerCh1datinrType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh1datinrFieldIndat0Mask) >> RegisterCh1datinrFieldIndat0Shift)
}

// SetIndat0 INDAT0
func (r *registerCh1datinrType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1datinrFieldIndat0Mask)|(uint32(value)<<RegisterCh1datinrFieldIndat0Shift))
}

const (
	RegisterCh1datinrFieldIndat1Shift = 16
	RegisterCh1datinrFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 INDAT1
func (r *registerCh1datinrType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh1datinrFieldIndat1Mask) >> RegisterCh1datinrFieldIndat1Shift)
}

// SetIndat1 INDAT1
func (r *registerCh1datinrType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1datinrFieldIndat1Mask)|(uint32(value)<<RegisterCh1datinrFieldIndat1Shift))
}

// registerCh1dlyrType channel y delay register
type registerCh1dlyrType uint32

const (
	RegisterCh1dlyrFieldPlsskpShift = 0
	RegisterCh1dlyrFieldPlsskpMask  = 0x3f
)

// GetPlsskp PLSSKP
func (r *registerCh1dlyrType) GetPlsskp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh1dlyrFieldPlsskpMask) >> RegisterCh1dlyrFieldPlsskpShift)
}

// SetPlsskp PLSSKP
func (r *registerCh1dlyrType) SetPlsskp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh1dlyrFieldPlsskpMask)|(uint32(value)<<RegisterCh1dlyrFieldPlsskpShift))
}

// registerCh2cfgr1Type CH2CFGR1
type registerCh2cfgr1Type uint32

const (
	RegisterCh2cfgr1FieldSitpShift = 0
	RegisterCh2cfgr1FieldSitpMask  = 0x3
)

// GetSitp SITP
func (r *registerCh2cfgr1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh2cfgr1FieldSitpMask) >> RegisterCh2cfgr1FieldSitpShift)
}

// SetSitp SITP
func (r *registerCh2cfgr1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2cfgr1FieldSitpMask)|(uint32(value)<<RegisterCh2cfgr1FieldSitpShift))
}

const (
	RegisterCh2cfgr1FieldSpickselShift = 2
	RegisterCh2cfgr1FieldSpickselMask  = 0xc
)

// GetSpicksel SPICKSEL
func (r *registerCh2cfgr1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh2cfgr1FieldSpickselMask) >> RegisterCh2cfgr1FieldSpickselShift)
}

// SetSpicksel SPICKSEL
func (r *registerCh2cfgr1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2cfgr1FieldSpickselMask)|(uint32(value)<<RegisterCh2cfgr1FieldSpickselShift))
}

const (
	RegisterCh2cfgr1FieldScdenShift = 5
	RegisterCh2cfgr1FieldScdenMask  = 0x20
)

// GetScden SCDEN
func (r *registerCh2cfgr1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh2cfgr1FieldScdenMask) != 0
}

// SetScden SCDEN
func (r *registerCh2cfgr1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh2cfgr1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh2cfgr1FieldScdenMask)
	}
}

const (
	RegisterCh2cfgr1FieldCkabenShift = 6
	RegisterCh2cfgr1FieldCkabenMask  = 0x40
)

// GetCkaben CKABEN
func (r *registerCh2cfgr1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh2cfgr1FieldCkabenMask) != 0
}

// SetCkaben CKABEN
func (r *registerCh2cfgr1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh2cfgr1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh2cfgr1FieldCkabenMask)
	}
}

const (
	RegisterCh2cfgr1FieldChenShift = 7
	RegisterCh2cfgr1FieldChenMask  = 0x80
)

// GetChen CHEN
func (r *registerCh2cfgr1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh2cfgr1FieldChenMask) != 0
}

// SetChen CHEN
func (r *registerCh2cfgr1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh2cfgr1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh2cfgr1FieldChenMask)
	}
}

const (
	RegisterCh2cfgr1FieldChinselShift = 8
	RegisterCh2cfgr1FieldChinselMask  = 0x100
)

// GetChinsel CHINSEL
func (r *registerCh2cfgr1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh2cfgr1FieldChinselMask) != 0
}

// SetChinsel CHINSEL
func (r *registerCh2cfgr1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh2cfgr1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh2cfgr1FieldChinselMask)
	}
}

const (
	RegisterCh2cfgr1FieldDatmpxShift = 12
	RegisterCh2cfgr1FieldDatmpxMask  = 0x3000
)

// GetDatmpx DATMPX
func (r *registerCh2cfgr1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh2cfgr1FieldDatmpxMask) >> RegisterCh2cfgr1FieldDatmpxShift)
}

// SetDatmpx DATMPX
func (r *registerCh2cfgr1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2cfgr1FieldDatmpxMask)|(uint32(value)<<RegisterCh2cfgr1FieldDatmpxShift))
}

const (
	RegisterCh2cfgr1FieldDatpackShift = 14
	RegisterCh2cfgr1FieldDatpackMask  = 0xc000
)

// GetDatpack DATPACK
func (r *registerCh2cfgr1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh2cfgr1FieldDatpackMask) >> RegisterCh2cfgr1FieldDatpackShift)
}

// SetDatpack DATPACK
func (r *registerCh2cfgr1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2cfgr1FieldDatpackMask)|(uint32(value)<<RegisterCh2cfgr1FieldDatpackShift))
}

// registerCh2cfgr2Type CH2CFGR2
type registerCh2cfgr2Type uint32

const (
	RegisterCh2cfgr2FieldDtrbsShift = 3
	RegisterCh2cfgr2FieldDtrbsMask  = 0xf8
)

// GetDtrbs DTRBS
func (r *registerCh2cfgr2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh2cfgr2FieldDtrbsMask) >> RegisterCh2cfgr2FieldDtrbsShift)
}

// SetDtrbs DTRBS
func (r *registerCh2cfgr2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2cfgr2FieldDtrbsMask)|(uint32(value)<<RegisterCh2cfgr2FieldDtrbsShift))
}

const (
	RegisterCh2cfgr2FieldOffsetShift = 8
	RegisterCh2cfgr2FieldOffsetMask  = 0xffffff00
)

// GetOffset OFFSET
func (r *registerCh2cfgr2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCh2cfgr2FieldOffsetMask) >> RegisterCh2cfgr2FieldOffsetShift)
}

// SetOffset OFFSET
func (r *registerCh2cfgr2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2cfgr2FieldOffsetMask)|(uint32(value)<<RegisterCh2cfgr2FieldOffsetShift))
}

// registerCh2awscdrType CH2AWSCDR
type registerCh2awscdrType uint32

const (
	RegisterCh2awscdrFieldScdtShift = 0
	RegisterCh2awscdrFieldScdtMask  = 0xff
)

// GetScdt SCDT
func (r *registerCh2awscdrType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh2awscdrFieldScdtMask) >> RegisterCh2awscdrFieldScdtShift)
}

// SetScdt SCDT
func (r *registerCh2awscdrType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2awscdrFieldScdtMask)|(uint32(value)<<RegisterCh2awscdrFieldScdtShift))
}

const (
	RegisterCh2awscdrFieldBkscdShift = 12
	RegisterCh2awscdrFieldBkscdMask  = 0xf000
)

// GetBkscd BKSCD
func (r *registerCh2awscdrType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh2awscdrFieldBkscdMask) >> RegisterCh2awscdrFieldBkscdShift)
}

// SetBkscd BKSCD
func (r *registerCh2awscdrType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2awscdrFieldBkscdMask)|(uint32(value)<<RegisterCh2awscdrFieldBkscdShift))
}

const (
	RegisterCh2awscdrFieldAwfosrShift = 16
	RegisterCh2awscdrFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr AWFOSR
func (r *registerCh2awscdrType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh2awscdrFieldAwfosrMask) >> RegisterCh2awscdrFieldAwfosrShift)
}

// SetAwfosr AWFOSR
func (r *registerCh2awscdrType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2awscdrFieldAwfosrMask)|(uint32(value)<<RegisterCh2awscdrFieldAwfosrShift))
}

const (
	RegisterCh2awscdrFieldAwfordShift = 22
	RegisterCh2awscdrFieldAwfordMask  = 0xc00000
)

// GetAwford AWFORD
func (r *registerCh2awscdrType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh2awscdrFieldAwfordMask) >> RegisterCh2awscdrFieldAwfordShift)
}

// SetAwford AWFORD
func (r *registerCh2awscdrType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2awscdrFieldAwfordMask)|(uint32(value)<<RegisterCh2awscdrFieldAwfordShift))
}

// registerCh2wdatrType CH2WDATR
type registerCh2wdatrType uint32

const (
	RegisterCh2wdatrFieldWdataShift = 0
	RegisterCh2wdatrFieldWdataMask  = 0xffff
)

// GetWdata WDATA
func (r *registerCh2wdatrType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh2wdatrFieldWdataMask) >> RegisterCh2wdatrFieldWdataShift)
}

// SetWdata WDATA
func (r *registerCh2wdatrType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2wdatrFieldWdataMask)|(uint32(value)<<RegisterCh2wdatrFieldWdataShift))
}

// registerCh2datinrType CH2DATINR
type registerCh2datinrType uint32

const (
	RegisterCh2datinrFieldIndat0Shift = 0
	RegisterCh2datinrFieldIndat0Mask  = 0xffff
)

// GetIndat0 INDAT0
func (r *registerCh2datinrType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh2datinrFieldIndat0Mask) >> RegisterCh2datinrFieldIndat0Shift)
}

// SetIndat0 INDAT0
func (r *registerCh2datinrType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2datinrFieldIndat0Mask)|(uint32(value)<<RegisterCh2datinrFieldIndat0Shift))
}

const (
	RegisterCh2datinrFieldIndat1Shift = 16
	RegisterCh2datinrFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 INDAT1
func (r *registerCh2datinrType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh2datinrFieldIndat1Mask) >> RegisterCh2datinrFieldIndat1Shift)
}

// SetIndat1 INDAT1
func (r *registerCh2datinrType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2datinrFieldIndat1Mask)|(uint32(value)<<RegisterCh2datinrFieldIndat1Shift))
}

// registerCh2dlyrType channel y delay register
type registerCh2dlyrType uint32

const (
	RegisterCh2dlyrFieldPlsskpShift = 0
	RegisterCh2dlyrFieldPlsskpMask  = 0x3f
)

// GetPlsskp PLSSKP
func (r *registerCh2dlyrType) GetPlsskp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh2dlyrFieldPlsskpMask) >> RegisterCh2dlyrFieldPlsskpShift)
}

// SetPlsskp PLSSKP
func (r *registerCh2dlyrType) SetPlsskp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh2dlyrFieldPlsskpMask)|(uint32(value)<<RegisterCh2dlyrFieldPlsskpShift))
}

// registerCh3cfgr1Type CH3CFGR1
type registerCh3cfgr1Type uint32

const (
	RegisterCh3cfgr1FieldSitpShift = 0
	RegisterCh3cfgr1FieldSitpMask  = 0x3
)

// GetSitp SITP
func (r *registerCh3cfgr1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh3cfgr1FieldSitpMask) >> RegisterCh3cfgr1FieldSitpShift)
}

// SetSitp SITP
func (r *registerCh3cfgr1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3cfgr1FieldSitpMask)|(uint32(value)<<RegisterCh3cfgr1FieldSitpShift))
}

const (
	RegisterCh3cfgr1FieldSpickselShift = 2
	RegisterCh3cfgr1FieldSpickselMask  = 0xc
)

// GetSpicksel SPICKSEL
func (r *registerCh3cfgr1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh3cfgr1FieldSpickselMask) >> RegisterCh3cfgr1FieldSpickselShift)
}

// SetSpicksel SPICKSEL
func (r *registerCh3cfgr1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3cfgr1FieldSpickselMask)|(uint32(value)<<RegisterCh3cfgr1FieldSpickselShift))
}

const (
	RegisterCh3cfgr1FieldScdenShift = 5
	RegisterCh3cfgr1FieldScdenMask  = 0x20
)

// GetScden SCDEN
func (r *registerCh3cfgr1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh3cfgr1FieldScdenMask) != 0
}

// SetScden SCDEN
func (r *registerCh3cfgr1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh3cfgr1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh3cfgr1FieldScdenMask)
	}
}

const (
	RegisterCh3cfgr1FieldCkabenShift = 6
	RegisterCh3cfgr1FieldCkabenMask  = 0x40
)

// GetCkaben CKABEN
func (r *registerCh3cfgr1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh3cfgr1FieldCkabenMask) != 0
}

// SetCkaben CKABEN
func (r *registerCh3cfgr1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh3cfgr1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh3cfgr1FieldCkabenMask)
	}
}

const (
	RegisterCh3cfgr1FieldChenShift = 7
	RegisterCh3cfgr1FieldChenMask  = 0x80
)

// GetChen CHEN
func (r *registerCh3cfgr1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh3cfgr1FieldChenMask) != 0
}

// SetChen CHEN
func (r *registerCh3cfgr1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh3cfgr1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh3cfgr1FieldChenMask)
	}
}

const (
	RegisterCh3cfgr1FieldChinselShift = 8
	RegisterCh3cfgr1FieldChinselMask  = 0x100
)

// GetChinsel CHINSEL
func (r *registerCh3cfgr1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh3cfgr1FieldChinselMask) != 0
}

// SetChinsel CHINSEL
func (r *registerCh3cfgr1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh3cfgr1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh3cfgr1FieldChinselMask)
	}
}

const (
	RegisterCh3cfgr1FieldDatmpxShift = 12
	RegisterCh3cfgr1FieldDatmpxMask  = 0x3000
)

// GetDatmpx DATMPX
func (r *registerCh3cfgr1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh3cfgr1FieldDatmpxMask) >> RegisterCh3cfgr1FieldDatmpxShift)
}

// SetDatmpx DATMPX
func (r *registerCh3cfgr1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3cfgr1FieldDatmpxMask)|(uint32(value)<<RegisterCh3cfgr1FieldDatmpxShift))
}

const (
	RegisterCh3cfgr1FieldDatpackShift = 14
	RegisterCh3cfgr1FieldDatpackMask  = 0xc000
)

// GetDatpack DATPACK
func (r *registerCh3cfgr1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh3cfgr1FieldDatpackMask) >> RegisterCh3cfgr1FieldDatpackShift)
}

// SetDatpack DATPACK
func (r *registerCh3cfgr1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3cfgr1FieldDatpackMask)|(uint32(value)<<RegisterCh3cfgr1FieldDatpackShift))
}

// registerCh3cfgr2Type CH3CFGR2
type registerCh3cfgr2Type uint32

const (
	RegisterCh3cfgr2FieldDtrbsShift = 3
	RegisterCh3cfgr2FieldDtrbsMask  = 0xf8
)

// GetDtrbs DTRBS
func (r *registerCh3cfgr2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh3cfgr2FieldDtrbsMask) >> RegisterCh3cfgr2FieldDtrbsShift)
}

// SetDtrbs DTRBS
func (r *registerCh3cfgr2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3cfgr2FieldDtrbsMask)|(uint32(value)<<RegisterCh3cfgr2FieldDtrbsShift))
}

const (
	RegisterCh3cfgr2FieldOffsetShift = 8
	RegisterCh3cfgr2FieldOffsetMask  = 0xffffff00
)

// GetOffset OFFSET
func (r *registerCh3cfgr2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCh3cfgr2FieldOffsetMask) >> RegisterCh3cfgr2FieldOffsetShift)
}

// SetOffset OFFSET
func (r *registerCh3cfgr2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3cfgr2FieldOffsetMask)|(uint32(value)<<RegisterCh3cfgr2FieldOffsetShift))
}

// registerCh3awscdrType CH3AWSCDR
type registerCh3awscdrType uint32

const (
	RegisterCh3awscdrFieldScdtShift = 0
	RegisterCh3awscdrFieldScdtMask  = 0xff
)

// GetScdt SCDT
func (r *registerCh3awscdrType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh3awscdrFieldScdtMask) >> RegisterCh3awscdrFieldScdtShift)
}

// SetScdt SCDT
func (r *registerCh3awscdrType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3awscdrFieldScdtMask)|(uint32(value)<<RegisterCh3awscdrFieldScdtShift))
}

const (
	RegisterCh3awscdrFieldBkscdShift = 12
	RegisterCh3awscdrFieldBkscdMask  = 0xf000
)

// GetBkscd BKSCD
func (r *registerCh3awscdrType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh3awscdrFieldBkscdMask) >> RegisterCh3awscdrFieldBkscdShift)
}

// SetBkscd BKSCD
func (r *registerCh3awscdrType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3awscdrFieldBkscdMask)|(uint32(value)<<RegisterCh3awscdrFieldBkscdShift))
}

const (
	RegisterCh3awscdrFieldAwfosrShift = 16
	RegisterCh3awscdrFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr AWFOSR
func (r *registerCh3awscdrType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh3awscdrFieldAwfosrMask) >> RegisterCh3awscdrFieldAwfosrShift)
}

// SetAwfosr AWFOSR
func (r *registerCh3awscdrType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3awscdrFieldAwfosrMask)|(uint32(value)<<RegisterCh3awscdrFieldAwfosrShift))
}

const (
	RegisterCh3awscdrFieldAwfordShift = 22
	RegisterCh3awscdrFieldAwfordMask  = 0xc00000
)

// GetAwford AWFORD
func (r *registerCh3awscdrType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh3awscdrFieldAwfordMask) >> RegisterCh3awscdrFieldAwfordShift)
}

// SetAwford AWFORD
func (r *registerCh3awscdrType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3awscdrFieldAwfordMask)|(uint32(value)<<RegisterCh3awscdrFieldAwfordShift))
}

// registerCh3wdatrType CH3WDATR
type registerCh3wdatrType uint32

const (
	RegisterCh3wdatrFieldWdataShift = 0
	RegisterCh3wdatrFieldWdataMask  = 0xffff
)

// GetWdata WDATA
func (r *registerCh3wdatrType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh3wdatrFieldWdataMask) >> RegisterCh3wdatrFieldWdataShift)
}

// SetWdata WDATA
func (r *registerCh3wdatrType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3wdatrFieldWdataMask)|(uint32(value)<<RegisterCh3wdatrFieldWdataShift))
}

// registerCh3datinrType CH3DATINR
type registerCh3datinrType uint32

const (
	RegisterCh3datinrFieldIndat0Shift = 0
	RegisterCh3datinrFieldIndat0Mask  = 0xffff
)

// GetIndat0 INDAT0
func (r *registerCh3datinrType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh3datinrFieldIndat0Mask) >> RegisterCh3datinrFieldIndat0Shift)
}

// SetIndat0 INDAT0
func (r *registerCh3datinrType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3datinrFieldIndat0Mask)|(uint32(value)<<RegisterCh3datinrFieldIndat0Shift))
}

const (
	RegisterCh3datinrFieldIndat1Shift = 16
	RegisterCh3datinrFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 INDAT1
func (r *registerCh3datinrType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh3datinrFieldIndat1Mask) >> RegisterCh3datinrFieldIndat1Shift)
}

// SetIndat1 INDAT1
func (r *registerCh3datinrType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3datinrFieldIndat1Mask)|(uint32(value)<<RegisterCh3datinrFieldIndat1Shift))
}

// registerCh3dlyrType channel y delay register
type registerCh3dlyrType uint32

const (
	RegisterCh3dlyrFieldPlsskpShift = 0
	RegisterCh3dlyrFieldPlsskpMask  = 0x3f
)

// GetPlsskp PLSSKP
func (r *registerCh3dlyrType) GetPlsskp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh3dlyrFieldPlsskpMask) >> RegisterCh3dlyrFieldPlsskpShift)
}

// SetPlsskp PLSSKP
func (r *registerCh3dlyrType) SetPlsskp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh3dlyrFieldPlsskpMask)|(uint32(value)<<RegisterCh3dlyrFieldPlsskpShift))
}

// registerCh4cfgr1Type CH4CFGR1
type registerCh4cfgr1Type uint32

const (
	RegisterCh4cfgr1FieldSitpShift = 0
	RegisterCh4cfgr1FieldSitpMask  = 0x3
)

// GetSitp SITP
func (r *registerCh4cfgr1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh4cfgr1FieldSitpMask) >> RegisterCh4cfgr1FieldSitpShift)
}

// SetSitp SITP
func (r *registerCh4cfgr1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4cfgr1FieldSitpMask)|(uint32(value)<<RegisterCh4cfgr1FieldSitpShift))
}

const (
	RegisterCh4cfgr1FieldSpickselShift = 2
	RegisterCh4cfgr1FieldSpickselMask  = 0xc
)

// GetSpicksel SPICKSEL
func (r *registerCh4cfgr1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh4cfgr1FieldSpickselMask) >> RegisterCh4cfgr1FieldSpickselShift)
}

// SetSpicksel SPICKSEL
func (r *registerCh4cfgr1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4cfgr1FieldSpickselMask)|(uint32(value)<<RegisterCh4cfgr1FieldSpickselShift))
}

const (
	RegisterCh4cfgr1FieldScdenShift = 5
	RegisterCh4cfgr1FieldScdenMask  = 0x20
)

// GetScden SCDEN
func (r *registerCh4cfgr1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh4cfgr1FieldScdenMask) != 0
}

// SetScden SCDEN
func (r *registerCh4cfgr1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh4cfgr1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh4cfgr1FieldScdenMask)
	}
}

const (
	RegisterCh4cfgr1FieldCkabenShift = 6
	RegisterCh4cfgr1FieldCkabenMask  = 0x40
)

// GetCkaben CKABEN
func (r *registerCh4cfgr1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh4cfgr1FieldCkabenMask) != 0
}

// SetCkaben CKABEN
func (r *registerCh4cfgr1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh4cfgr1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh4cfgr1FieldCkabenMask)
	}
}

const (
	RegisterCh4cfgr1FieldChenShift = 7
	RegisterCh4cfgr1FieldChenMask  = 0x80
)

// GetChen CHEN
func (r *registerCh4cfgr1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh4cfgr1FieldChenMask) != 0
}

// SetChen CHEN
func (r *registerCh4cfgr1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh4cfgr1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh4cfgr1FieldChenMask)
	}
}

const (
	RegisterCh4cfgr1FieldChinselShift = 8
	RegisterCh4cfgr1FieldChinselMask  = 0x100
)

// GetChinsel CHINSEL
func (r *registerCh4cfgr1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh4cfgr1FieldChinselMask) != 0
}

// SetChinsel CHINSEL
func (r *registerCh4cfgr1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh4cfgr1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh4cfgr1FieldChinselMask)
	}
}

const (
	RegisterCh4cfgr1FieldDatmpxShift = 12
	RegisterCh4cfgr1FieldDatmpxMask  = 0x3000
)

// GetDatmpx DATMPX
func (r *registerCh4cfgr1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh4cfgr1FieldDatmpxMask) >> RegisterCh4cfgr1FieldDatmpxShift)
}

// SetDatmpx DATMPX
func (r *registerCh4cfgr1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4cfgr1FieldDatmpxMask)|(uint32(value)<<RegisterCh4cfgr1FieldDatmpxShift))
}

const (
	RegisterCh4cfgr1FieldDatpackShift = 14
	RegisterCh4cfgr1FieldDatpackMask  = 0xc000
)

// GetDatpack DATPACK
func (r *registerCh4cfgr1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh4cfgr1FieldDatpackMask) >> RegisterCh4cfgr1FieldDatpackShift)
}

// SetDatpack DATPACK
func (r *registerCh4cfgr1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4cfgr1FieldDatpackMask)|(uint32(value)<<RegisterCh4cfgr1FieldDatpackShift))
}

// registerCh4cfgr2Type CH4CFGR2
type registerCh4cfgr2Type uint32

const (
	RegisterCh4cfgr2FieldDtrbsShift = 3
	RegisterCh4cfgr2FieldDtrbsMask  = 0xf8
)

// GetDtrbs DTRBS
func (r *registerCh4cfgr2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh4cfgr2FieldDtrbsMask) >> RegisterCh4cfgr2FieldDtrbsShift)
}

// SetDtrbs DTRBS
func (r *registerCh4cfgr2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4cfgr2FieldDtrbsMask)|(uint32(value)<<RegisterCh4cfgr2FieldDtrbsShift))
}

const (
	RegisterCh4cfgr2FieldOffsetShift = 8
	RegisterCh4cfgr2FieldOffsetMask  = 0xffffff00
)

// GetOffset OFFSET
func (r *registerCh4cfgr2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCh4cfgr2FieldOffsetMask) >> RegisterCh4cfgr2FieldOffsetShift)
}

// SetOffset OFFSET
func (r *registerCh4cfgr2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4cfgr2FieldOffsetMask)|(uint32(value)<<RegisterCh4cfgr2FieldOffsetShift))
}

// registerCh4awscdrType CH4AWSCDR
type registerCh4awscdrType uint32

const (
	RegisterCh4awscdrFieldScdtShift = 0
	RegisterCh4awscdrFieldScdtMask  = 0xff
)

// GetScdt SCDT
func (r *registerCh4awscdrType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh4awscdrFieldScdtMask) >> RegisterCh4awscdrFieldScdtShift)
}

// SetScdt SCDT
func (r *registerCh4awscdrType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4awscdrFieldScdtMask)|(uint32(value)<<RegisterCh4awscdrFieldScdtShift))
}

const (
	RegisterCh4awscdrFieldBkscdShift = 12
	RegisterCh4awscdrFieldBkscdMask  = 0xf000
)

// GetBkscd BKSCD
func (r *registerCh4awscdrType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh4awscdrFieldBkscdMask) >> RegisterCh4awscdrFieldBkscdShift)
}

// SetBkscd BKSCD
func (r *registerCh4awscdrType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4awscdrFieldBkscdMask)|(uint32(value)<<RegisterCh4awscdrFieldBkscdShift))
}

const (
	RegisterCh4awscdrFieldAwfosrShift = 16
	RegisterCh4awscdrFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr AWFOSR
func (r *registerCh4awscdrType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh4awscdrFieldAwfosrMask) >> RegisterCh4awscdrFieldAwfosrShift)
}

// SetAwfosr AWFOSR
func (r *registerCh4awscdrType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4awscdrFieldAwfosrMask)|(uint32(value)<<RegisterCh4awscdrFieldAwfosrShift))
}

const (
	RegisterCh4awscdrFieldAwfordShift = 22
	RegisterCh4awscdrFieldAwfordMask  = 0xc00000
)

// GetAwford AWFORD
func (r *registerCh4awscdrType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh4awscdrFieldAwfordMask) >> RegisterCh4awscdrFieldAwfordShift)
}

// SetAwford AWFORD
func (r *registerCh4awscdrType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4awscdrFieldAwfordMask)|(uint32(value)<<RegisterCh4awscdrFieldAwfordShift))
}

// registerCh4wdatrType CH4WDATR
type registerCh4wdatrType uint32

const (
	RegisterCh4wdatrFieldWdataShift = 0
	RegisterCh4wdatrFieldWdataMask  = 0xffff
)

// GetWdata WDATA
func (r *registerCh4wdatrType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh4wdatrFieldWdataMask) >> RegisterCh4wdatrFieldWdataShift)
}

// SetWdata WDATA
func (r *registerCh4wdatrType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4wdatrFieldWdataMask)|(uint32(value)<<RegisterCh4wdatrFieldWdataShift))
}

// registerCh4datinrType CH4DATINR
type registerCh4datinrType uint32

const (
	RegisterCh4datinrFieldIndat0Shift = 0
	RegisterCh4datinrFieldIndat0Mask  = 0xffff
)

// GetIndat0 INDAT0
func (r *registerCh4datinrType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh4datinrFieldIndat0Mask) >> RegisterCh4datinrFieldIndat0Shift)
}

// SetIndat0 INDAT0
func (r *registerCh4datinrType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4datinrFieldIndat0Mask)|(uint32(value)<<RegisterCh4datinrFieldIndat0Shift))
}

const (
	RegisterCh4datinrFieldIndat1Shift = 16
	RegisterCh4datinrFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 INDAT1
func (r *registerCh4datinrType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh4datinrFieldIndat1Mask) >> RegisterCh4datinrFieldIndat1Shift)
}

// SetIndat1 INDAT1
func (r *registerCh4datinrType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4datinrFieldIndat1Mask)|(uint32(value)<<RegisterCh4datinrFieldIndat1Shift))
}

// registerCh4dlyrType channel y delay register
type registerCh4dlyrType uint32

const (
	RegisterCh4dlyrFieldPlsskpShift = 0
	RegisterCh4dlyrFieldPlsskpMask  = 0x3f
)

// GetPlsskp PLSSKP
func (r *registerCh4dlyrType) GetPlsskp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh4dlyrFieldPlsskpMask) >> RegisterCh4dlyrFieldPlsskpShift)
}

// SetPlsskp PLSSKP
func (r *registerCh4dlyrType) SetPlsskp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh4dlyrFieldPlsskpMask)|(uint32(value)<<RegisterCh4dlyrFieldPlsskpShift))
}

// registerCh5cfgr1Type CH5CFGR1
type registerCh5cfgr1Type uint32

const (
	RegisterCh5cfgr1FieldSitpShift = 0
	RegisterCh5cfgr1FieldSitpMask  = 0x3
)

// GetSitp SITP
func (r *registerCh5cfgr1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh5cfgr1FieldSitpMask) >> RegisterCh5cfgr1FieldSitpShift)
}

// SetSitp SITP
func (r *registerCh5cfgr1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5cfgr1FieldSitpMask)|(uint32(value)<<RegisterCh5cfgr1FieldSitpShift))
}

const (
	RegisterCh5cfgr1FieldSpickselShift = 2
	RegisterCh5cfgr1FieldSpickselMask  = 0xc
)

// GetSpicksel SPICKSEL
func (r *registerCh5cfgr1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh5cfgr1FieldSpickselMask) >> RegisterCh5cfgr1FieldSpickselShift)
}

// SetSpicksel SPICKSEL
func (r *registerCh5cfgr1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5cfgr1FieldSpickselMask)|(uint32(value)<<RegisterCh5cfgr1FieldSpickselShift))
}

const (
	RegisterCh5cfgr1FieldScdenShift = 5
	RegisterCh5cfgr1FieldScdenMask  = 0x20
)

// GetScden SCDEN
func (r *registerCh5cfgr1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh5cfgr1FieldScdenMask) != 0
}

// SetScden SCDEN
func (r *registerCh5cfgr1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh5cfgr1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh5cfgr1FieldScdenMask)
	}
}

const (
	RegisterCh5cfgr1FieldCkabenShift = 6
	RegisterCh5cfgr1FieldCkabenMask  = 0x40
)

// GetCkaben CKABEN
func (r *registerCh5cfgr1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh5cfgr1FieldCkabenMask) != 0
}

// SetCkaben CKABEN
func (r *registerCh5cfgr1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh5cfgr1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh5cfgr1FieldCkabenMask)
	}
}

const (
	RegisterCh5cfgr1FieldChenShift = 7
	RegisterCh5cfgr1FieldChenMask  = 0x80
)

// GetChen CHEN
func (r *registerCh5cfgr1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh5cfgr1FieldChenMask) != 0
}

// SetChen CHEN
func (r *registerCh5cfgr1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh5cfgr1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh5cfgr1FieldChenMask)
	}
}

const (
	RegisterCh5cfgr1FieldChinselShift = 8
	RegisterCh5cfgr1FieldChinselMask  = 0x100
)

// GetChinsel CHINSEL
func (r *registerCh5cfgr1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh5cfgr1FieldChinselMask) != 0
}

// SetChinsel CHINSEL
func (r *registerCh5cfgr1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh5cfgr1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh5cfgr1FieldChinselMask)
	}
}

const (
	RegisterCh5cfgr1FieldDatmpxShift = 12
	RegisterCh5cfgr1FieldDatmpxMask  = 0x3000
)

// GetDatmpx DATMPX
func (r *registerCh5cfgr1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh5cfgr1FieldDatmpxMask) >> RegisterCh5cfgr1FieldDatmpxShift)
}

// SetDatmpx DATMPX
func (r *registerCh5cfgr1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5cfgr1FieldDatmpxMask)|(uint32(value)<<RegisterCh5cfgr1FieldDatmpxShift))
}

const (
	RegisterCh5cfgr1FieldDatpackShift = 14
	RegisterCh5cfgr1FieldDatpackMask  = 0xc000
)

// GetDatpack DATPACK
func (r *registerCh5cfgr1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh5cfgr1FieldDatpackMask) >> RegisterCh5cfgr1FieldDatpackShift)
}

// SetDatpack DATPACK
func (r *registerCh5cfgr1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5cfgr1FieldDatpackMask)|(uint32(value)<<RegisterCh5cfgr1FieldDatpackShift))
}

// registerCh5cfgr2Type CH5CFGR2
type registerCh5cfgr2Type uint32

const (
	RegisterCh5cfgr2FieldDtrbsShift = 3
	RegisterCh5cfgr2FieldDtrbsMask  = 0xf8
)

// GetDtrbs DTRBS
func (r *registerCh5cfgr2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh5cfgr2FieldDtrbsMask) >> RegisterCh5cfgr2FieldDtrbsShift)
}

// SetDtrbs DTRBS
func (r *registerCh5cfgr2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5cfgr2FieldDtrbsMask)|(uint32(value)<<RegisterCh5cfgr2FieldDtrbsShift))
}

const (
	RegisterCh5cfgr2FieldOffsetShift = 8
	RegisterCh5cfgr2FieldOffsetMask  = 0xffffff00
)

// GetOffset OFFSET
func (r *registerCh5cfgr2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCh5cfgr2FieldOffsetMask) >> RegisterCh5cfgr2FieldOffsetShift)
}

// SetOffset OFFSET
func (r *registerCh5cfgr2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5cfgr2FieldOffsetMask)|(uint32(value)<<RegisterCh5cfgr2FieldOffsetShift))
}

// registerCh5awscdrType CH5AWSCDR
type registerCh5awscdrType uint32

const (
	RegisterCh5awscdrFieldScdtShift = 0
	RegisterCh5awscdrFieldScdtMask  = 0xff
)

// GetScdt SCDT
func (r *registerCh5awscdrType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh5awscdrFieldScdtMask) >> RegisterCh5awscdrFieldScdtShift)
}

// SetScdt SCDT
func (r *registerCh5awscdrType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5awscdrFieldScdtMask)|(uint32(value)<<RegisterCh5awscdrFieldScdtShift))
}

const (
	RegisterCh5awscdrFieldBkscdShift = 12
	RegisterCh5awscdrFieldBkscdMask  = 0xf000
)

// GetBkscd BKSCD
func (r *registerCh5awscdrType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh5awscdrFieldBkscdMask) >> RegisterCh5awscdrFieldBkscdShift)
}

// SetBkscd BKSCD
func (r *registerCh5awscdrType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5awscdrFieldBkscdMask)|(uint32(value)<<RegisterCh5awscdrFieldBkscdShift))
}

const (
	RegisterCh5awscdrFieldAwfosrShift = 16
	RegisterCh5awscdrFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr AWFOSR
func (r *registerCh5awscdrType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh5awscdrFieldAwfosrMask) >> RegisterCh5awscdrFieldAwfosrShift)
}

// SetAwfosr AWFOSR
func (r *registerCh5awscdrType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5awscdrFieldAwfosrMask)|(uint32(value)<<RegisterCh5awscdrFieldAwfosrShift))
}

const (
	RegisterCh5awscdrFieldAwfordShift = 22
	RegisterCh5awscdrFieldAwfordMask  = 0xc00000
)

// GetAwford AWFORD
func (r *registerCh5awscdrType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh5awscdrFieldAwfordMask) >> RegisterCh5awscdrFieldAwfordShift)
}

// SetAwford AWFORD
func (r *registerCh5awscdrType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5awscdrFieldAwfordMask)|(uint32(value)<<RegisterCh5awscdrFieldAwfordShift))
}

// registerCh5wdatrType CH5WDATR
type registerCh5wdatrType uint32

const (
	RegisterCh5wdatrFieldWdataShift = 0
	RegisterCh5wdatrFieldWdataMask  = 0xffff
)

// GetWdata WDATA
func (r *registerCh5wdatrType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh5wdatrFieldWdataMask) >> RegisterCh5wdatrFieldWdataShift)
}

// SetWdata WDATA
func (r *registerCh5wdatrType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5wdatrFieldWdataMask)|(uint32(value)<<RegisterCh5wdatrFieldWdataShift))
}

// registerCh5datinrType CH5DATINR
type registerCh5datinrType uint32

const (
	RegisterCh5datinrFieldIndat0Shift = 0
	RegisterCh5datinrFieldIndat0Mask  = 0xffff
)

// GetIndat0 INDAT0
func (r *registerCh5datinrType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh5datinrFieldIndat0Mask) >> RegisterCh5datinrFieldIndat0Shift)
}

// SetIndat0 INDAT0
func (r *registerCh5datinrType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5datinrFieldIndat0Mask)|(uint32(value)<<RegisterCh5datinrFieldIndat0Shift))
}

const (
	RegisterCh5datinrFieldIndat1Shift = 16
	RegisterCh5datinrFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 INDAT1
func (r *registerCh5datinrType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh5datinrFieldIndat1Mask) >> RegisterCh5datinrFieldIndat1Shift)
}

// SetIndat1 INDAT1
func (r *registerCh5datinrType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5datinrFieldIndat1Mask)|(uint32(value)<<RegisterCh5datinrFieldIndat1Shift))
}

// registerCh5dlyrType channel y delay register
type registerCh5dlyrType uint32

const (
	RegisterCh5dlyrFieldPlsskpShift = 0
	RegisterCh5dlyrFieldPlsskpMask  = 0x3f
)

// GetPlsskp PLSSKP
func (r *registerCh5dlyrType) GetPlsskp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh5dlyrFieldPlsskpMask) >> RegisterCh5dlyrFieldPlsskpShift)
}

// SetPlsskp PLSSKP
func (r *registerCh5dlyrType) SetPlsskp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh5dlyrFieldPlsskpMask)|(uint32(value)<<RegisterCh5dlyrFieldPlsskpShift))
}

// registerCh6cfgr1Type CH6CFGR1
type registerCh6cfgr1Type uint32

const (
	RegisterCh6cfgr1FieldSitpShift = 0
	RegisterCh6cfgr1FieldSitpMask  = 0x3
)

// GetSitp SITP
func (r *registerCh6cfgr1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh6cfgr1FieldSitpMask) >> RegisterCh6cfgr1FieldSitpShift)
}

// SetSitp SITP
func (r *registerCh6cfgr1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6cfgr1FieldSitpMask)|(uint32(value)<<RegisterCh6cfgr1FieldSitpShift))
}

const (
	RegisterCh6cfgr1FieldSpickselShift = 2
	RegisterCh6cfgr1FieldSpickselMask  = 0xc
)

// GetSpicksel SPICKSEL
func (r *registerCh6cfgr1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh6cfgr1FieldSpickselMask) >> RegisterCh6cfgr1FieldSpickselShift)
}

// SetSpicksel SPICKSEL
func (r *registerCh6cfgr1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6cfgr1FieldSpickselMask)|(uint32(value)<<RegisterCh6cfgr1FieldSpickselShift))
}

const (
	RegisterCh6cfgr1FieldScdenShift = 5
	RegisterCh6cfgr1FieldScdenMask  = 0x20
)

// GetScden SCDEN
func (r *registerCh6cfgr1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh6cfgr1FieldScdenMask) != 0
}

// SetScden SCDEN
func (r *registerCh6cfgr1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh6cfgr1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh6cfgr1FieldScdenMask)
	}
}

const (
	RegisterCh6cfgr1FieldCkabenShift = 6
	RegisterCh6cfgr1FieldCkabenMask  = 0x40
)

// GetCkaben CKABEN
func (r *registerCh6cfgr1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh6cfgr1FieldCkabenMask) != 0
}

// SetCkaben CKABEN
func (r *registerCh6cfgr1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh6cfgr1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh6cfgr1FieldCkabenMask)
	}
}

const (
	RegisterCh6cfgr1FieldChenShift = 7
	RegisterCh6cfgr1FieldChenMask  = 0x80
)

// GetChen CHEN
func (r *registerCh6cfgr1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh6cfgr1FieldChenMask) != 0
}

// SetChen CHEN
func (r *registerCh6cfgr1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh6cfgr1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh6cfgr1FieldChenMask)
	}
}

const (
	RegisterCh6cfgr1FieldChinselShift = 8
	RegisterCh6cfgr1FieldChinselMask  = 0x100
)

// GetChinsel CHINSEL
func (r *registerCh6cfgr1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh6cfgr1FieldChinselMask) != 0
}

// SetChinsel CHINSEL
func (r *registerCh6cfgr1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh6cfgr1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh6cfgr1FieldChinselMask)
	}
}

const (
	RegisterCh6cfgr1FieldDatmpxShift = 12
	RegisterCh6cfgr1FieldDatmpxMask  = 0x3000
)

// GetDatmpx DATMPX
func (r *registerCh6cfgr1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh6cfgr1FieldDatmpxMask) >> RegisterCh6cfgr1FieldDatmpxShift)
}

// SetDatmpx DATMPX
func (r *registerCh6cfgr1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6cfgr1FieldDatmpxMask)|(uint32(value)<<RegisterCh6cfgr1FieldDatmpxShift))
}

const (
	RegisterCh6cfgr1FieldDatpackShift = 14
	RegisterCh6cfgr1FieldDatpackMask  = 0xc000
)

// GetDatpack DATPACK
func (r *registerCh6cfgr1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh6cfgr1FieldDatpackMask) >> RegisterCh6cfgr1FieldDatpackShift)
}

// SetDatpack DATPACK
func (r *registerCh6cfgr1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6cfgr1FieldDatpackMask)|(uint32(value)<<RegisterCh6cfgr1FieldDatpackShift))
}

// registerCh6cfgr2Type CH6CFGR2
type registerCh6cfgr2Type uint32

const (
	RegisterCh6cfgr2FieldDtrbsShift = 3
	RegisterCh6cfgr2FieldDtrbsMask  = 0xf8
)

// GetDtrbs DTRBS
func (r *registerCh6cfgr2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh6cfgr2FieldDtrbsMask) >> RegisterCh6cfgr2FieldDtrbsShift)
}

// SetDtrbs DTRBS
func (r *registerCh6cfgr2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6cfgr2FieldDtrbsMask)|(uint32(value)<<RegisterCh6cfgr2FieldDtrbsShift))
}

const (
	RegisterCh6cfgr2FieldOffsetShift = 8
	RegisterCh6cfgr2FieldOffsetMask  = 0xffffff00
)

// GetOffset OFFSET
func (r *registerCh6cfgr2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCh6cfgr2FieldOffsetMask) >> RegisterCh6cfgr2FieldOffsetShift)
}

// SetOffset OFFSET
func (r *registerCh6cfgr2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6cfgr2FieldOffsetMask)|(uint32(value)<<RegisterCh6cfgr2FieldOffsetShift))
}

// registerCh6awscdrType CH6AWSCDR
type registerCh6awscdrType uint32

const (
	RegisterCh6awscdrFieldScdtShift = 0
	RegisterCh6awscdrFieldScdtMask  = 0xff
)

// GetScdt SCDT
func (r *registerCh6awscdrType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh6awscdrFieldScdtMask) >> RegisterCh6awscdrFieldScdtShift)
}

// SetScdt SCDT
func (r *registerCh6awscdrType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6awscdrFieldScdtMask)|(uint32(value)<<RegisterCh6awscdrFieldScdtShift))
}

const (
	RegisterCh6awscdrFieldBkscdShift = 12
	RegisterCh6awscdrFieldBkscdMask  = 0xf000
)

// GetBkscd BKSCD
func (r *registerCh6awscdrType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh6awscdrFieldBkscdMask) >> RegisterCh6awscdrFieldBkscdShift)
}

// SetBkscd BKSCD
func (r *registerCh6awscdrType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6awscdrFieldBkscdMask)|(uint32(value)<<RegisterCh6awscdrFieldBkscdShift))
}

const (
	RegisterCh6awscdrFieldAwfosrShift = 16
	RegisterCh6awscdrFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr AWFOSR
func (r *registerCh6awscdrType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh6awscdrFieldAwfosrMask) >> RegisterCh6awscdrFieldAwfosrShift)
}

// SetAwfosr AWFOSR
func (r *registerCh6awscdrType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6awscdrFieldAwfosrMask)|(uint32(value)<<RegisterCh6awscdrFieldAwfosrShift))
}

const (
	RegisterCh6awscdrFieldAwfordShift = 22
	RegisterCh6awscdrFieldAwfordMask  = 0xc00000
)

// GetAwford AWFORD
func (r *registerCh6awscdrType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh6awscdrFieldAwfordMask) >> RegisterCh6awscdrFieldAwfordShift)
}

// SetAwford AWFORD
func (r *registerCh6awscdrType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6awscdrFieldAwfordMask)|(uint32(value)<<RegisterCh6awscdrFieldAwfordShift))
}

// registerCh6wdatrType CH6WDATR
type registerCh6wdatrType uint32

const (
	RegisterCh6wdatrFieldWdataShift = 0
	RegisterCh6wdatrFieldWdataMask  = 0xffff
)

// GetWdata WDATA
func (r *registerCh6wdatrType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh6wdatrFieldWdataMask) >> RegisterCh6wdatrFieldWdataShift)
}

// SetWdata WDATA
func (r *registerCh6wdatrType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6wdatrFieldWdataMask)|(uint32(value)<<RegisterCh6wdatrFieldWdataShift))
}

// registerCh6datinrType CH6DATINR
type registerCh6datinrType uint32

const (
	RegisterCh6datinrFieldIndat0Shift = 0
	RegisterCh6datinrFieldIndat0Mask  = 0xffff
)

// GetIndat0 INDAT0
func (r *registerCh6datinrType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh6datinrFieldIndat0Mask) >> RegisterCh6datinrFieldIndat0Shift)
}

// SetIndat0 INDAT0
func (r *registerCh6datinrType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6datinrFieldIndat0Mask)|(uint32(value)<<RegisterCh6datinrFieldIndat0Shift))
}

const (
	RegisterCh6datinrFieldIndat1Shift = 16
	RegisterCh6datinrFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 INDAT1
func (r *registerCh6datinrType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh6datinrFieldIndat1Mask) >> RegisterCh6datinrFieldIndat1Shift)
}

// SetIndat1 INDAT1
func (r *registerCh6datinrType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6datinrFieldIndat1Mask)|(uint32(value)<<RegisterCh6datinrFieldIndat1Shift))
}

// registerCh6dlyrType channel y delay register
type registerCh6dlyrType uint32

const (
	RegisterCh6dlyrFieldPlsskpShift = 0
	RegisterCh6dlyrFieldPlsskpMask  = 0x3f
)

// GetPlsskp PLSSKP
func (r *registerCh6dlyrType) GetPlsskp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh6dlyrFieldPlsskpMask) >> RegisterCh6dlyrFieldPlsskpShift)
}

// SetPlsskp PLSSKP
func (r *registerCh6dlyrType) SetPlsskp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh6dlyrFieldPlsskpMask)|(uint32(value)<<RegisterCh6dlyrFieldPlsskpShift))
}

// registerCh7cfgr1Type CH7CFGR1
type registerCh7cfgr1Type uint32

const (
	RegisterCh7cfgr1FieldSitpShift = 0
	RegisterCh7cfgr1FieldSitpMask  = 0x3
)

// GetSitp SITP
func (r *registerCh7cfgr1Type) GetSitp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh7cfgr1FieldSitpMask) >> RegisterCh7cfgr1FieldSitpShift)
}

// SetSitp SITP
func (r *registerCh7cfgr1Type) SetSitp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7cfgr1FieldSitpMask)|(uint32(value)<<RegisterCh7cfgr1FieldSitpShift))
}

const (
	RegisterCh7cfgr1FieldSpickselShift = 2
	RegisterCh7cfgr1FieldSpickselMask  = 0xc
)

// GetSpicksel SPICKSEL
func (r *registerCh7cfgr1Type) GetSpicksel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh7cfgr1FieldSpickselMask) >> RegisterCh7cfgr1FieldSpickselShift)
}

// SetSpicksel SPICKSEL
func (r *registerCh7cfgr1Type) SetSpicksel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7cfgr1FieldSpickselMask)|(uint32(value)<<RegisterCh7cfgr1FieldSpickselShift))
}

const (
	RegisterCh7cfgr1FieldScdenShift = 5
	RegisterCh7cfgr1FieldScdenMask  = 0x20
)

// GetScden SCDEN
func (r *registerCh7cfgr1Type) GetScden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh7cfgr1FieldScdenMask) != 0
}

// SetScden SCDEN
func (r *registerCh7cfgr1Type) SetScden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh7cfgr1FieldScdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh7cfgr1FieldScdenMask)
	}
}

const (
	RegisterCh7cfgr1FieldCkabenShift = 6
	RegisterCh7cfgr1FieldCkabenMask  = 0x40
)

// GetCkaben CKABEN
func (r *registerCh7cfgr1Type) GetCkaben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh7cfgr1FieldCkabenMask) != 0
}

// SetCkaben CKABEN
func (r *registerCh7cfgr1Type) SetCkaben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh7cfgr1FieldCkabenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh7cfgr1FieldCkabenMask)
	}
}

const (
	RegisterCh7cfgr1FieldChenShift = 7
	RegisterCh7cfgr1FieldChenMask  = 0x80
)

// GetChen CHEN
func (r *registerCh7cfgr1Type) GetChen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh7cfgr1FieldChenMask) != 0
}

// SetChen CHEN
func (r *registerCh7cfgr1Type) SetChen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh7cfgr1FieldChenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh7cfgr1FieldChenMask)
	}
}

const (
	RegisterCh7cfgr1FieldChinselShift = 8
	RegisterCh7cfgr1FieldChinselMask  = 0x100
)

// GetChinsel CHINSEL
func (r *registerCh7cfgr1Type) GetChinsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCh7cfgr1FieldChinselMask) != 0
}

// SetChinsel CHINSEL
func (r *registerCh7cfgr1Type) SetChinsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCh7cfgr1FieldChinselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCh7cfgr1FieldChinselMask)
	}
}

const (
	RegisterCh7cfgr1FieldDatmpxShift = 12
	RegisterCh7cfgr1FieldDatmpxMask  = 0x3000
)

// GetDatmpx DATMPX
func (r *registerCh7cfgr1Type) GetDatmpx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh7cfgr1FieldDatmpxMask) >> RegisterCh7cfgr1FieldDatmpxShift)
}

// SetDatmpx DATMPX
func (r *registerCh7cfgr1Type) SetDatmpx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7cfgr1FieldDatmpxMask)|(uint32(value)<<RegisterCh7cfgr1FieldDatmpxShift))
}

const (
	RegisterCh7cfgr1FieldDatpackShift = 14
	RegisterCh7cfgr1FieldDatpackMask  = 0xc000
)

// GetDatpack DATPACK
func (r *registerCh7cfgr1Type) GetDatpack() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh7cfgr1FieldDatpackMask) >> RegisterCh7cfgr1FieldDatpackShift)
}

// SetDatpack DATPACK
func (r *registerCh7cfgr1Type) SetDatpack(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7cfgr1FieldDatpackMask)|(uint32(value)<<RegisterCh7cfgr1FieldDatpackShift))
}

// registerCh7cfgr2Type CH7CFGR2
type registerCh7cfgr2Type uint32

const (
	RegisterCh7cfgr2FieldDtrbsShift = 3
	RegisterCh7cfgr2FieldDtrbsMask  = 0xf8
)

// GetDtrbs DTRBS
func (r *registerCh7cfgr2Type) GetDtrbs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh7cfgr2FieldDtrbsMask) >> RegisterCh7cfgr2FieldDtrbsShift)
}

// SetDtrbs DTRBS
func (r *registerCh7cfgr2Type) SetDtrbs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7cfgr2FieldDtrbsMask)|(uint32(value)<<RegisterCh7cfgr2FieldDtrbsShift))
}

const (
	RegisterCh7cfgr2FieldOffsetShift = 8
	RegisterCh7cfgr2FieldOffsetMask  = 0xffffff00
)

// GetOffset OFFSET
func (r *registerCh7cfgr2Type) GetOffset() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCh7cfgr2FieldOffsetMask) >> RegisterCh7cfgr2FieldOffsetShift)
}

// SetOffset OFFSET
func (r *registerCh7cfgr2Type) SetOffset(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7cfgr2FieldOffsetMask)|(uint32(value)<<RegisterCh7cfgr2FieldOffsetShift))
}

// registerCh7awscdrType CH7AWSCDR
type registerCh7awscdrType uint32

const (
	RegisterCh7awscdrFieldScdtShift = 0
	RegisterCh7awscdrFieldScdtMask  = 0xff
)

// GetScdt SCDT
func (r *registerCh7awscdrType) GetScdt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh7awscdrFieldScdtMask) >> RegisterCh7awscdrFieldScdtShift)
}

// SetScdt SCDT
func (r *registerCh7awscdrType) SetScdt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7awscdrFieldScdtMask)|(uint32(value)<<RegisterCh7awscdrFieldScdtShift))
}

const (
	RegisterCh7awscdrFieldBkscdShift = 12
	RegisterCh7awscdrFieldBkscdMask  = 0xf000
)

// GetBkscd BKSCD
func (r *registerCh7awscdrType) GetBkscd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh7awscdrFieldBkscdMask) >> RegisterCh7awscdrFieldBkscdShift)
}

// SetBkscd BKSCD
func (r *registerCh7awscdrType) SetBkscd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7awscdrFieldBkscdMask)|(uint32(value)<<RegisterCh7awscdrFieldBkscdShift))
}

const (
	RegisterCh7awscdrFieldAwfosrShift = 16
	RegisterCh7awscdrFieldAwfosrMask  = 0x1f0000
)

// GetAwfosr AWFOSR
func (r *registerCh7awscdrType) GetAwfosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh7awscdrFieldAwfosrMask) >> RegisterCh7awscdrFieldAwfosrShift)
}

// SetAwfosr AWFOSR
func (r *registerCh7awscdrType) SetAwfosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7awscdrFieldAwfosrMask)|(uint32(value)<<RegisterCh7awscdrFieldAwfosrShift))
}

const (
	RegisterCh7awscdrFieldAwfordShift = 22
	RegisterCh7awscdrFieldAwfordMask  = 0xc00000
)

// GetAwford AWFORD
func (r *registerCh7awscdrType) GetAwford() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh7awscdrFieldAwfordMask) >> RegisterCh7awscdrFieldAwfordShift)
}

// SetAwford AWFORD
func (r *registerCh7awscdrType) SetAwford(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7awscdrFieldAwfordMask)|(uint32(value)<<RegisterCh7awscdrFieldAwfordShift))
}

// registerCh7wdatrType CH7WDATR
type registerCh7wdatrType uint32

const (
	RegisterCh7wdatrFieldWdataShift = 0
	RegisterCh7wdatrFieldWdataMask  = 0xffff
)

// GetWdata WDATA
func (r *registerCh7wdatrType) GetWdata() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh7wdatrFieldWdataMask) >> RegisterCh7wdatrFieldWdataShift)
}

// SetWdata WDATA
func (r *registerCh7wdatrType) SetWdata(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7wdatrFieldWdataMask)|(uint32(value)<<RegisterCh7wdatrFieldWdataShift))
}

// registerCh7datinrType CH7DATINR
type registerCh7datinrType uint32

const (
	RegisterCh7datinrFieldIndat0Shift = 0
	RegisterCh7datinrFieldIndat0Mask  = 0xffff
)

// GetIndat0 INDAT0
func (r *registerCh7datinrType) GetIndat0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh7datinrFieldIndat0Mask) >> RegisterCh7datinrFieldIndat0Shift)
}

// SetIndat0 INDAT0
func (r *registerCh7datinrType) SetIndat0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7datinrFieldIndat0Mask)|(uint32(value)<<RegisterCh7datinrFieldIndat0Shift))
}

const (
	RegisterCh7datinrFieldIndat1Shift = 16
	RegisterCh7datinrFieldIndat1Mask  = 0xffff0000
)

// GetIndat1 INDAT1
func (r *registerCh7datinrType) GetIndat1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCh7datinrFieldIndat1Mask) >> RegisterCh7datinrFieldIndat1Shift)
}

// SetIndat1 INDAT1
func (r *registerCh7datinrType) SetIndat1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7datinrFieldIndat1Mask)|(uint32(value)<<RegisterCh7datinrFieldIndat1Shift))
}

// registerCh7dlyrType channel y delay register
type registerCh7dlyrType uint32

const (
	RegisterCh7dlyrFieldPlsskpShift = 0
	RegisterCh7dlyrFieldPlsskpMask  = 0x3f
)

// GetPlsskp PLSSKP
func (r *registerCh7dlyrType) GetPlsskp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCh7dlyrFieldPlsskpMask) >> RegisterCh7dlyrFieldPlsskpShift)
}

// SetPlsskp PLSSKP
func (r *registerCh7dlyrType) SetPlsskp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCh7dlyrFieldPlsskpMask)|(uint32(value)<<RegisterCh7dlyrFieldPlsskpShift))
}

// registerDfsdm_flt0cr1Type control register 1
type registerDfsdm_flt0cr1Type uint32

const (
	RegisterDfsdm_flt0cr1FieldDfenShift = 0
	RegisterDfsdm_flt0cr1FieldDfenMask  = 0x1
)

// GetDfen DFSDM enable
func (r *registerDfsdm_flt0cr1Type) GetDfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldDfenMask) != 0
}

// SetDfen DFSDM enable
func (r *registerDfsdm_flt0cr1Type) SetDfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr1FieldDfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldDfenMask)
	}
}

const (
	RegisterDfsdm_flt0cr1FieldJswstartShift = 1
	RegisterDfsdm_flt0cr1FieldJswstartMask  = 0x2
)

// GetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm_flt0cr1Type) GetJswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldJswstartMask) != 0
}

// SetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm_flt0cr1Type) SetJswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr1FieldJswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldJswstartMask)
	}
}

const (
	RegisterDfsdm_flt0cr1FieldJsyncShift = 3
	RegisterDfsdm_flt0cr1FieldJsyncMask  = 0x8
)

// GetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm_flt0cr1Type) GetJsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldJsyncMask) != 0
}

// SetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm_flt0cr1Type) SetJsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr1FieldJsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldJsyncMask)
	}
}

const (
	RegisterDfsdm_flt0cr1FieldJscanShift = 4
	RegisterDfsdm_flt0cr1FieldJscanMask  = 0x10
)

// GetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm_flt0cr1Type) GetJscan() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldJscanMask) != 0
}

// SetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm_flt0cr1Type) SetJscan(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr1FieldJscanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldJscanMask)
	}
}

const (
	RegisterDfsdm_flt0cr1FieldJdmaenShift = 5
	RegisterDfsdm_flt0cr1FieldJdmaenMask  = 0x20
)

// GetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm_flt0cr1Type) GetJdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldJdmaenMask) != 0
}

// SetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm_flt0cr1Type) SetJdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr1FieldJdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldJdmaenMask)
	}
}

const (
	RegisterDfsdm_flt0cr1FieldJextselShift = 8
	RegisterDfsdm_flt0cr1FieldJextselMask  = 0x700
)

// GetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm_flt0cr1Type) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldJextselMask) >> RegisterDfsdm_flt0cr1FieldJextselShift)
}

// SetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm_flt0cr1Type) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldJextselMask)|(uint32(value)<<RegisterDfsdm_flt0cr1FieldJextselShift))
}

const (
	RegisterDfsdm_flt0cr1FieldJextenShift = 13
	RegisterDfsdm_flt0cr1FieldJextenMask  = 0x6000
)

// GetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm_flt0cr1Type) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldJextenMask) >> RegisterDfsdm_flt0cr1FieldJextenShift)
}

// SetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm_flt0cr1Type) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldJextenMask)|(uint32(value)<<RegisterDfsdm_flt0cr1FieldJextenShift))
}

const (
	RegisterDfsdm_flt0cr1FieldRswstartShift = 17
	RegisterDfsdm_flt0cr1FieldRswstartMask  = 0x20000
)

// GetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm_flt0cr1Type) GetRswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldRswstartMask) != 0
}

// SetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm_flt0cr1Type) SetRswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr1FieldRswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldRswstartMask)
	}
}

const (
	RegisterDfsdm_flt0cr1FieldRcontShift = 18
	RegisterDfsdm_flt0cr1FieldRcontMask  = 0x40000
)

// GetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm_flt0cr1Type) GetRcont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldRcontMask) != 0
}

// SetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm_flt0cr1Type) SetRcont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr1FieldRcontMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldRcontMask)
	}
}

const (
	RegisterDfsdm_flt0cr1FieldRsyncShift = 19
	RegisterDfsdm_flt0cr1FieldRsyncMask  = 0x80000
)

// GetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm_flt0cr1Type) GetRsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldRsyncMask) != 0
}

// SetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm_flt0cr1Type) SetRsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr1FieldRsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldRsyncMask)
	}
}

const (
	RegisterDfsdm_flt0cr1FieldRdmaenShift = 21
	RegisterDfsdm_flt0cr1FieldRdmaenMask  = 0x200000
)

// GetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm_flt0cr1Type) GetRdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldRdmaenMask) != 0
}

// SetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm_flt0cr1Type) SetRdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr1FieldRdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldRdmaenMask)
	}
}

const (
	RegisterDfsdm_flt0cr1FieldRchShift = 24
	RegisterDfsdm_flt0cr1FieldRchMask  = 0x7000000
)

// GetRch Regular channel selection
func (r *registerDfsdm_flt0cr1Type) GetRch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldRchMask) >> RegisterDfsdm_flt0cr1FieldRchShift)
}

// SetRch Regular channel selection
func (r *registerDfsdm_flt0cr1Type) SetRch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldRchMask)|(uint32(value)<<RegisterDfsdm_flt0cr1FieldRchShift))
}

const (
	RegisterDfsdm_flt0cr1FieldFastShift = 29
	RegisterDfsdm_flt0cr1FieldFastMask  = 0x20000000
)

// GetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm_flt0cr1Type) GetFast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldFastMask) != 0
}

// SetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm_flt0cr1Type) SetFast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr1FieldFastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldFastMask)
	}
}

const (
	RegisterDfsdm_flt0cr1FieldAwfselShift = 30
	RegisterDfsdm_flt0cr1FieldAwfselMask  = 0x40000000
)

// GetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm_flt0cr1Type) GetAwfsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr1FieldAwfselMask) != 0
}

// SetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm_flt0cr1Type) SetAwfsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr1FieldAwfselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr1FieldAwfselMask)
	}
}

// registerDfsdm_flt0cr2Type control register 2
type registerDfsdm_flt0cr2Type uint32

const (
	RegisterDfsdm_flt0cr2FieldJeocieShift = 0
	RegisterDfsdm_flt0cr2FieldJeocieMask  = 0x1
)

// GetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm_flt0cr2Type) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr2FieldJeocieMask) != 0
}

// SetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm_flt0cr2Type) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr2FieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr2FieldJeocieMask)
	}
}

const (
	RegisterDfsdm_flt0cr2FieldReocieShift = 1
	RegisterDfsdm_flt0cr2FieldReocieMask  = 0x2
)

// GetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm_flt0cr2Type) GetReocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr2FieldReocieMask) != 0
}

// SetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm_flt0cr2Type) SetReocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr2FieldReocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr2FieldReocieMask)
	}
}

const (
	RegisterDfsdm_flt0cr2FieldJovrieShift = 2
	RegisterDfsdm_flt0cr2FieldJovrieMask  = 0x4
)

// GetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm_flt0cr2Type) GetJovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr2FieldJovrieMask) != 0
}

// SetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm_flt0cr2Type) SetJovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr2FieldJovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr2FieldJovrieMask)
	}
}

const (
	RegisterDfsdm_flt0cr2FieldRovrieShift = 3
	RegisterDfsdm_flt0cr2FieldRovrieMask  = 0x8
)

// GetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm_flt0cr2Type) GetRovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr2FieldRovrieMask) != 0
}

// SetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm_flt0cr2Type) SetRovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr2FieldRovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr2FieldRovrieMask)
	}
}

const (
	RegisterDfsdm_flt0cr2FieldAwdieShift = 4
	RegisterDfsdm_flt0cr2FieldAwdieMask  = 0x10
)

// GetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm_flt0cr2Type) GetAwdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr2FieldAwdieMask) != 0
}

// SetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm_flt0cr2Type) SetAwdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr2FieldAwdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr2FieldAwdieMask)
	}
}

const (
	RegisterDfsdm_flt0cr2FieldScdieShift = 5
	RegisterDfsdm_flt0cr2FieldScdieMask  = 0x20
)

// GetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm_flt0cr2Type) GetScdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr2FieldScdieMask) != 0
}

// SetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm_flt0cr2Type) SetScdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr2FieldScdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr2FieldScdieMask)
	}
}

const (
	RegisterDfsdm_flt0cr2FieldCkabieShift = 6
	RegisterDfsdm_flt0cr2FieldCkabieMask  = 0x40
)

// GetCkabie Clock absence interrupt enable
func (r *registerDfsdm_flt0cr2Type) GetCkabie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr2FieldCkabieMask) != 0
}

// SetCkabie Clock absence interrupt enable
func (r *registerDfsdm_flt0cr2Type) SetCkabie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0cr2FieldCkabieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr2FieldCkabieMask)
	}
}

const (
	RegisterDfsdm_flt0cr2FieldExchShift = 8
	RegisterDfsdm_flt0cr2FieldExchMask  = 0xff00
)

// GetExch Extremes detector channel selection
func (r *registerDfsdm_flt0cr2Type) GetExch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr2FieldExchMask) >> RegisterDfsdm_flt0cr2FieldExchShift)
}

// SetExch Extremes detector channel selection
func (r *registerDfsdm_flt0cr2Type) SetExch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr2FieldExchMask)|(uint32(value)<<RegisterDfsdm_flt0cr2FieldExchShift))
}

const (
	RegisterDfsdm_flt0cr2FieldAwdchShift = 16
	RegisterDfsdm_flt0cr2FieldAwdchMask  = 0xff0000
)

// GetAwdch Analog watchdog channel selection
func (r *registerDfsdm_flt0cr2Type) GetAwdch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cr2FieldAwdchMask) >> RegisterDfsdm_flt0cr2FieldAwdchShift)
}

// SetAwdch Analog watchdog channel selection
func (r *registerDfsdm_flt0cr2Type) SetAwdch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cr2FieldAwdchMask)|(uint32(value)<<RegisterDfsdm_flt0cr2FieldAwdchShift))
}

// registerDfsdm_flt0isrType interrupt and status register
type registerDfsdm_flt0isrType uint32

const (
	RegisterDfsdm_flt0isrFieldJeocfShift = 0
	RegisterDfsdm_flt0isrFieldJeocfMask  = 0x1
)

// GetJeocf End of injected conversion flag
func (r *registerDfsdm_flt0isrType) GetJeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0isrFieldJeocfMask) != 0
}

// SetJeocf End of injected conversion flag
func (r *registerDfsdm_flt0isrType) SetJeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0isrFieldJeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0isrFieldJeocfMask)
	}
}

const (
	RegisterDfsdm_flt0isrFieldReocfShift = 1
	RegisterDfsdm_flt0isrFieldReocfMask  = 0x2
)

// GetReocf End of regular conversion flag
func (r *registerDfsdm_flt0isrType) GetReocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0isrFieldReocfMask) != 0
}

// SetReocf End of regular conversion flag
func (r *registerDfsdm_flt0isrType) SetReocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0isrFieldReocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0isrFieldReocfMask)
	}
}

const (
	RegisterDfsdm_flt0isrFieldJovrfShift = 2
	RegisterDfsdm_flt0isrFieldJovrfMask  = 0x4
)

// GetJovrf Injected conversion overrun flag
func (r *registerDfsdm_flt0isrType) GetJovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0isrFieldJovrfMask) != 0
}

// SetJovrf Injected conversion overrun flag
func (r *registerDfsdm_flt0isrType) SetJovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0isrFieldJovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0isrFieldJovrfMask)
	}
}

const (
	RegisterDfsdm_flt0isrFieldRovrfShift = 3
	RegisterDfsdm_flt0isrFieldRovrfMask  = 0x8
)

// GetRovrf Regular conversion overrun flag
func (r *registerDfsdm_flt0isrType) GetRovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0isrFieldRovrfMask) != 0
}

// SetRovrf Regular conversion overrun flag
func (r *registerDfsdm_flt0isrType) SetRovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0isrFieldRovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0isrFieldRovrfMask)
	}
}

const (
	RegisterDfsdm_flt0isrFieldAwdfShift = 4
	RegisterDfsdm_flt0isrFieldAwdfMask  = 0x10
)

// GetAwdf Analog watchdog
func (r *registerDfsdm_flt0isrType) GetAwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0isrFieldAwdfMask) != 0
}

// SetAwdf Analog watchdog
func (r *registerDfsdm_flt0isrType) SetAwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0isrFieldAwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0isrFieldAwdfMask)
	}
}

const (
	RegisterDfsdm_flt0isrFieldJcipShift = 13
	RegisterDfsdm_flt0isrFieldJcipMask  = 0x2000
)

// GetJcip Injected conversion in progress status
func (r *registerDfsdm_flt0isrType) GetJcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0isrFieldJcipMask) != 0
}

// SetJcip Injected conversion in progress status
func (r *registerDfsdm_flt0isrType) SetJcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0isrFieldJcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0isrFieldJcipMask)
	}
}

const (
	RegisterDfsdm_flt0isrFieldRcipShift = 14
	RegisterDfsdm_flt0isrFieldRcipMask  = 0x4000
)

// GetRcip Regular conversion in progress status
func (r *registerDfsdm_flt0isrType) GetRcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0isrFieldRcipMask) != 0
}

// SetRcip Regular conversion in progress status
func (r *registerDfsdm_flt0isrType) SetRcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0isrFieldRcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0isrFieldRcipMask)
	}
}

const (
	RegisterDfsdm_flt0isrFieldCkabfShift = 16
	RegisterDfsdm_flt0isrFieldCkabfMask  = 0xff0000
)

// GetCkabf Clock absence flag
func (r *registerDfsdm_flt0isrType) GetCkabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0isrFieldCkabfMask) >> RegisterDfsdm_flt0isrFieldCkabfShift)
}

// SetCkabf Clock absence flag
func (r *registerDfsdm_flt0isrType) SetCkabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0isrFieldCkabfMask)|(uint32(value)<<RegisterDfsdm_flt0isrFieldCkabfShift))
}

const (
	RegisterDfsdm_flt0isrFieldScdfShift = 24
	RegisterDfsdm_flt0isrFieldScdfMask  = 0xff000000
)

// GetScdf short-circuit detector flag
func (r *registerDfsdm_flt0isrType) GetScdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0isrFieldScdfMask) >> RegisterDfsdm_flt0isrFieldScdfShift)
}

// SetScdf short-circuit detector flag
func (r *registerDfsdm_flt0isrType) SetScdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0isrFieldScdfMask)|(uint32(value)<<RegisterDfsdm_flt0isrFieldScdfShift))
}

// registerDfsdm_flt0icrType interrupt flag clear register
type registerDfsdm_flt0icrType uint32

const (
	RegisterDfsdm_flt0icrFieldClrjovrfShift = 2
	RegisterDfsdm_flt0icrFieldClrjovrfMask  = 0x4
)

// GetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm_flt0icrType) GetClrjovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0icrFieldClrjovrfMask) != 0
}

// SetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm_flt0icrType) SetClrjovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0icrFieldClrjovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0icrFieldClrjovrfMask)
	}
}

const (
	RegisterDfsdm_flt0icrFieldClrrovrfShift = 3
	RegisterDfsdm_flt0icrFieldClrrovrfMask  = 0x8
)

// GetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm_flt0icrType) GetClrrovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0icrFieldClrrovrfMask) != 0
}

// SetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm_flt0icrType) SetClrrovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0icrFieldClrrovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0icrFieldClrrovrfMask)
	}
}

const (
	RegisterDfsdm_flt0icrFieldClrckabfShift = 16
	RegisterDfsdm_flt0icrFieldClrckabfMask  = 0xff0000
)

// GetClrckabf Clear the clock absence flag
func (r *registerDfsdm_flt0icrType) GetClrckabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0icrFieldClrckabfMask) >> RegisterDfsdm_flt0icrFieldClrckabfShift)
}

// SetClrckabf Clear the clock absence flag
func (r *registerDfsdm_flt0icrType) SetClrckabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0icrFieldClrckabfMask)|(uint32(value)<<RegisterDfsdm_flt0icrFieldClrckabfShift))
}

const (
	RegisterDfsdm_flt0icrFieldClrscdfShift = 24
	RegisterDfsdm_flt0icrFieldClrscdfMask  = 0xff000000
)

// GetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm_flt0icrType) GetClrscdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0icrFieldClrscdfMask) >> RegisterDfsdm_flt0icrFieldClrscdfShift)
}

// SetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm_flt0icrType) SetClrscdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0icrFieldClrscdfMask)|(uint32(value)<<RegisterDfsdm_flt0icrFieldClrscdfShift))
}

// registerDfsdm_flt0jchgrType injected channel group selection register
type registerDfsdm_flt0jchgrType uint32

const (
	RegisterDfsdm_flt0jchgrFieldJchgShift = 0
	RegisterDfsdm_flt0jchgrFieldJchgMask  = 0xff
)

// GetJchg Injected channel group selection
func (r *registerDfsdm_flt0jchgrType) GetJchg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0jchgrFieldJchgMask) >> RegisterDfsdm_flt0jchgrFieldJchgShift)
}

// SetJchg Injected channel group selection
func (r *registerDfsdm_flt0jchgrType) SetJchg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0jchgrFieldJchgMask)|(uint32(value)<<RegisterDfsdm_flt0jchgrFieldJchgShift))
}

// registerDfsdm_flt0fcrType filter control register
type registerDfsdm_flt0fcrType uint32

const (
	RegisterDfsdm_flt0fcrFieldIosrShift = 0
	RegisterDfsdm_flt0fcrFieldIosrMask  = 0xff
)

// GetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm_flt0fcrType) GetIosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0fcrFieldIosrMask) >> RegisterDfsdm_flt0fcrFieldIosrShift)
}

// SetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm_flt0fcrType) SetIosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0fcrFieldIosrMask)|(uint32(value)<<RegisterDfsdm_flt0fcrFieldIosrShift))
}

const (
	RegisterDfsdm_flt0fcrFieldFosrShift = 16
	RegisterDfsdm_flt0fcrFieldFosrMask  = 0x3ff0000
)

// GetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm_flt0fcrType) GetFosr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0fcrFieldFosrMask) >> RegisterDfsdm_flt0fcrFieldFosrShift)
}

// SetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm_flt0fcrType) SetFosr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0fcrFieldFosrMask)|(uint32(value)<<RegisterDfsdm_flt0fcrFieldFosrShift))
}

const (
	RegisterDfsdm_flt0fcrFieldFordShift = 29
	RegisterDfsdm_flt0fcrFieldFordMask  = 0xe0000000
)

// GetFord Sinc filter order
func (r *registerDfsdm_flt0fcrType) GetFord() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0fcrFieldFordMask) >> RegisterDfsdm_flt0fcrFieldFordShift)
}

// SetFord Sinc filter order
func (r *registerDfsdm_flt0fcrType) SetFord(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0fcrFieldFordMask)|(uint32(value)<<RegisterDfsdm_flt0fcrFieldFordShift))
}

// registerDfsdm_flt0jdatarType data register for injected group
type registerDfsdm_flt0jdatarType uint32

const (
	RegisterDfsdm_flt0jdatarFieldJdatachShift = 0
	RegisterDfsdm_flt0jdatarFieldJdatachMask  = 0x7
)

// GetJdatach Injected channel most recently converted
func (r *registerDfsdm_flt0jdatarType) GetJdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0jdatarFieldJdatachMask) >> RegisterDfsdm_flt0jdatarFieldJdatachShift)
}

// SetJdatach Injected channel most recently converted
func (r *registerDfsdm_flt0jdatarType) SetJdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0jdatarFieldJdatachMask)|(uint32(value)<<RegisterDfsdm_flt0jdatarFieldJdatachShift))
}

const (
	RegisterDfsdm_flt0jdatarFieldJdataShift = 8
	RegisterDfsdm_flt0jdatarFieldJdataMask  = 0xffffff00
)

// GetJdata Injected group conversion data
func (r *registerDfsdm_flt0jdatarType) GetJdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0jdatarFieldJdataMask) >> RegisterDfsdm_flt0jdatarFieldJdataShift)
}

// SetJdata Injected group conversion data
func (r *registerDfsdm_flt0jdatarType) SetJdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0jdatarFieldJdataMask)|(uint32(value)<<RegisterDfsdm_flt0jdatarFieldJdataShift))
}

// registerDfsdm_flt0rdatarType data register for the regular channel
type registerDfsdm_flt0rdatarType uint32

const (
	RegisterDfsdm_flt0rdatarFieldRdatachShift = 0
	RegisterDfsdm_flt0rdatarFieldRdatachMask  = 0x7
)

// GetRdatach Regular channel most recently converted
func (r *registerDfsdm_flt0rdatarType) GetRdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0rdatarFieldRdatachMask) >> RegisterDfsdm_flt0rdatarFieldRdatachShift)
}

// SetRdatach Regular channel most recently converted
func (r *registerDfsdm_flt0rdatarType) SetRdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0rdatarFieldRdatachMask)|(uint32(value)<<RegisterDfsdm_flt0rdatarFieldRdatachShift))
}

const (
	RegisterDfsdm_flt0rdatarFieldRpendShift = 4
	RegisterDfsdm_flt0rdatarFieldRpendMask  = 0x10
)

// GetRpend Regular channel pending data
func (r *registerDfsdm_flt0rdatarType) GetRpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0rdatarFieldRpendMask) != 0
}

// SetRpend Regular channel pending data
func (r *registerDfsdm_flt0rdatarType) SetRpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt0rdatarFieldRpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0rdatarFieldRpendMask)
	}
}

const (
	RegisterDfsdm_flt0rdatarFieldRdataShift = 8
	RegisterDfsdm_flt0rdatarFieldRdataMask  = 0xffffff00
)

// GetRdata Regular channel conversion data
func (r *registerDfsdm_flt0rdatarType) GetRdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0rdatarFieldRdataMask) >> RegisterDfsdm_flt0rdatarFieldRdataShift)
}

// SetRdata Regular channel conversion data
func (r *registerDfsdm_flt0rdatarType) SetRdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0rdatarFieldRdataMask)|(uint32(value)<<RegisterDfsdm_flt0rdatarFieldRdataShift))
}

// registerDfsdm_flt0awhtrType analog watchdog high threshold register
type registerDfsdm_flt0awhtrType uint32

const (
	RegisterDfsdm_flt0awhtrFieldBkawhShift = 0
	RegisterDfsdm_flt0awhtrFieldBkawhMask  = 0xf
)

// GetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm_flt0awhtrType) GetBkawh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0awhtrFieldBkawhMask) >> RegisterDfsdm_flt0awhtrFieldBkawhShift)
}

// SetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm_flt0awhtrType) SetBkawh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0awhtrFieldBkawhMask)|(uint32(value)<<RegisterDfsdm_flt0awhtrFieldBkawhShift))
}

const (
	RegisterDfsdm_flt0awhtrFieldAwhtShift = 8
	RegisterDfsdm_flt0awhtrFieldAwhtMask  = 0xffffff00
)

// GetAwht Analog watchdog high threshold
func (r *registerDfsdm_flt0awhtrType) GetAwht() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0awhtrFieldAwhtMask) >> RegisterDfsdm_flt0awhtrFieldAwhtShift)
}

// SetAwht Analog watchdog high threshold
func (r *registerDfsdm_flt0awhtrType) SetAwht(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0awhtrFieldAwhtMask)|(uint32(value)<<RegisterDfsdm_flt0awhtrFieldAwhtShift))
}

// registerDfsdm_flt0awltrType analog watchdog low threshold register
type registerDfsdm_flt0awltrType uint32

const (
	RegisterDfsdm_flt0awltrFieldBkawlShift = 0
	RegisterDfsdm_flt0awltrFieldBkawlMask  = 0xf
)

// GetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm_flt0awltrType) GetBkawl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0awltrFieldBkawlMask) >> RegisterDfsdm_flt0awltrFieldBkawlShift)
}

// SetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm_flt0awltrType) SetBkawl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0awltrFieldBkawlMask)|(uint32(value)<<RegisterDfsdm_flt0awltrFieldBkawlShift))
}

const (
	RegisterDfsdm_flt0awltrFieldAwltShift = 8
	RegisterDfsdm_flt0awltrFieldAwltMask  = 0xffffff00
)

// GetAwlt Analog watchdog low threshold
func (r *registerDfsdm_flt0awltrType) GetAwlt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0awltrFieldAwltMask) >> RegisterDfsdm_flt0awltrFieldAwltShift)
}

// SetAwlt Analog watchdog low threshold
func (r *registerDfsdm_flt0awltrType) SetAwlt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0awltrFieldAwltMask)|(uint32(value)<<RegisterDfsdm_flt0awltrFieldAwltShift))
}

// registerDfsdm_flt0awsrType analog watchdog status register
type registerDfsdm_flt0awsrType uint32

const (
	RegisterDfsdm_flt0awsrFieldAwltfShift = 0
	RegisterDfsdm_flt0awsrFieldAwltfMask  = 0xff
)

// GetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm_flt0awsrType) GetAwltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0awsrFieldAwltfMask) >> RegisterDfsdm_flt0awsrFieldAwltfShift)
}

// SetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm_flt0awsrType) SetAwltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0awsrFieldAwltfMask)|(uint32(value)<<RegisterDfsdm_flt0awsrFieldAwltfShift))
}

const (
	RegisterDfsdm_flt0awsrFieldAwhtfShift = 8
	RegisterDfsdm_flt0awsrFieldAwhtfMask  = 0xff00
)

// GetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm_flt0awsrType) GetAwhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0awsrFieldAwhtfMask) >> RegisterDfsdm_flt0awsrFieldAwhtfShift)
}

// SetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm_flt0awsrType) SetAwhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0awsrFieldAwhtfMask)|(uint32(value)<<RegisterDfsdm_flt0awsrFieldAwhtfShift))
}

// registerDfsdm_flt0awcfrType analog watchdog clear flag register
type registerDfsdm_flt0awcfrType uint32

const (
	RegisterDfsdm_flt0awcfrFieldClrawltfShift = 0
	RegisterDfsdm_flt0awcfrFieldClrawltfMask  = 0xff
)

// GetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm_flt0awcfrType) GetClrawltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0awcfrFieldClrawltfMask) >> RegisterDfsdm_flt0awcfrFieldClrawltfShift)
}

// SetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm_flt0awcfrType) SetClrawltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0awcfrFieldClrawltfMask)|(uint32(value)<<RegisterDfsdm_flt0awcfrFieldClrawltfShift))
}

const (
	RegisterDfsdm_flt0awcfrFieldClrawhtfShift = 8
	RegisterDfsdm_flt0awcfrFieldClrawhtfMask  = 0xff00
)

// GetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm_flt0awcfrType) GetClrawhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0awcfrFieldClrawhtfMask) >> RegisterDfsdm_flt0awcfrFieldClrawhtfShift)
}

// SetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm_flt0awcfrType) SetClrawhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0awcfrFieldClrawhtfMask)|(uint32(value)<<RegisterDfsdm_flt0awcfrFieldClrawhtfShift))
}

// registerDfsdm_flt0exmaxType Extremes detector maximum register
type registerDfsdm_flt0exmaxType uint32

const (
	RegisterDfsdm_flt0exmaxFieldExmaxchShift = 0
	RegisterDfsdm_flt0exmaxFieldExmaxchMask  = 0x7
)

// GetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm_flt0exmaxType) GetExmaxch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0exmaxFieldExmaxchMask) >> RegisterDfsdm_flt0exmaxFieldExmaxchShift)
}

// SetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm_flt0exmaxType) SetExmaxch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0exmaxFieldExmaxchMask)|(uint32(value)<<RegisterDfsdm_flt0exmaxFieldExmaxchShift))
}

const (
	RegisterDfsdm_flt0exmaxFieldExmaxShift = 8
	RegisterDfsdm_flt0exmaxFieldExmaxMask  = 0xffffff00
)

// GetExmax Extremes detector maximum value
func (r *registerDfsdm_flt0exmaxType) GetExmax() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0exmaxFieldExmaxMask) >> RegisterDfsdm_flt0exmaxFieldExmaxShift)
}

// SetExmax Extremes detector maximum value
func (r *registerDfsdm_flt0exmaxType) SetExmax(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0exmaxFieldExmaxMask)|(uint32(value)<<RegisterDfsdm_flt0exmaxFieldExmaxShift))
}

// registerDfsdm_flt0exminType Extremes detector minimum register
type registerDfsdm_flt0exminType uint32

const (
	RegisterDfsdm_flt0exminFieldExminchShift = 0
	RegisterDfsdm_flt0exminFieldExminchMask  = 0x7
)

// GetExminch Extremes detector minimum data channel
func (r *registerDfsdm_flt0exminType) GetExminch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0exminFieldExminchMask) >> RegisterDfsdm_flt0exminFieldExminchShift)
}

// SetExminch Extremes detector minimum data channel
func (r *registerDfsdm_flt0exminType) SetExminch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0exminFieldExminchMask)|(uint32(value)<<RegisterDfsdm_flt0exminFieldExminchShift))
}

const (
	RegisterDfsdm_flt0exminFieldExminShift = 8
	RegisterDfsdm_flt0exminFieldExminMask  = 0xffffff00
)

// GetExmin EXMIN
func (r *registerDfsdm_flt0exminType) GetExmin() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0exminFieldExminMask) >> RegisterDfsdm_flt0exminFieldExminShift)
}

// SetExmin EXMIN
func (r *registerDfsdm_flt0exminType) SetExmin(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0exminFieldExminMask)|(uint32(value)<<RegisterDfsdm_flt0exminFieldExminShift))
}

// registerDfsdm_flt0cnvtimrType conversion timer register
type registerDfsdm_flt0cnvtimrType uint32

const (
	RegisterDfsdm_flt0cnvtimrFieldCnvcntShift = 4
	RegisterDfsdm_flt0cnvtimrFieldCnvcntMask  = 0xfffffff0
)

// GetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdm_flt0cnvtimrType) GetCnvcnt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt0cnvtimrFieldCnvcntMask) >> RegisterDfsdm_flt0cnvtimrFieldCnvcntShift)
}

// SetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdm_flt0cnvtimrType) SetCnvcnt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt0cnvtimrFieldCnvcntMask)|(uint32(value)<<RegisterDfsdm_flt0cnvtimrFieldCnvcntShift))
}

// registerDfsdm_flt1cr1Type control register 1
type registerDfsdm_flt1cr1Type uint32

const (
	RegisterDfsdm_flt1cr1FieldDfenShift = 0
	RegisterDfsdm_flt1cr1FieldDfenMask  = 0x1
)

// GetDfen DFSDM enable
func (r *registerDfsdm_flt1cr1Type) GetDfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldDfenMask) != 0
}

// SetDfen DFSDM enable
func (r *registerDfsdm_flt1cr1Type) SetDfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr1FieldDfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldDfenMask)
	}
}

const (
	RegisterDfsdm_flt1cr1FieldJswstartShift = 1
	RegisterDfsdm_flt1cr1FieldJswstartMask  = 0x2
)

// GetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm_flt1cr1Type) GetJswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldJswstartMask) != 0
}

// SetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm_flt1cr1Type) SetJswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr1FieldJswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldJswstartMask)
	}
}

const (
	RegisterDfsdm_flt1cr1FieldJsyncShift = 3
	RegisterDfsdm_flt1cr1FieldJsyncMask  = 0x8
)

// GetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm_flt1cr1Type) GetJsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldJsyncMask) != 0
}

// SetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm_flt1cr1Type) SetJsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr1FieldJsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldJsyncMask)
	}
}

const (
	RegisterDfsdm_flt1cr1FieldJscanShift = 4
	RegisterDfsdm_flt1cr1FieldJscanMask  = 0x10
)

// GetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm_flt1cr1Type) GetJscan() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldJscanMask) != 0
}

// SetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm_flt1cr1Type) SetJscan(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr1FieldJscanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldJscanMask)
	}
}

const (
	RegisterDfsdm_flt1cr1FieldJdmaenShift = 5
	RegisterDfsdm_flt1cr1FieldJdmaenMask  = 0x20
)

// GetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm_flt1cr1Type) GetJdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldJdmaenMask) != 0
}

// SetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm_flt1cr1Type) SetJdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr1FieldJdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldJdmaenMask)
	}
}

const (
	RegisterDfsdm_flt1cr1FieldJextselShift = 8
	RegisterDfsdm_flt1cr1FieldJextselMask  = 0x700
)

// GetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm_flt1cr1Type) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldJextselMask) >> RegisterDfsdm_flt1cr1FieldJextselShift)
}

// SetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm_flt1cr1Type) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldJextselMask)|(uint32(value)<<RegisterDfsdm_flt1cr1FieldJextselShift))
}

const (
	RegisterDfsdm_flt1cr1FieldJextenShift = 13
	RegisterDfsdm_flt1cr1FieldJextenMask  = 0x6000
)

// GetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm_flt1cr1Type) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldJextenMask) >> RegisterDfsdm_flt1cr1FieldJextenShift)
}

// SetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm_flt1cr1Type) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldJextenMask)|(uint32(value)<<RegisterDfsdm_flt1cr1FieldJextenShift))
}

const (
	RegisterDfsdm_flt1cr1FieldRswstartShift = 17
	RegisterDfsdm_flt1cr1FieldRswstartMask  = 0x20000
)

// GetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm_flt1cr1Type) GetRswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldRswstartMask) != 0
}

// SetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm_flt1cr1Type) SetRswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr1FieldRswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldRswstartMask)
	}
}

const (
	RegisterDfsdm_flt1cr1FieldRcontShift = 18
	RegisterDfsdm_flt1cr1FieldRcontMask  = 0x40000
)

// GetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm_flt1cr1Type) GetRcont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldRcontMask) != 0
}

// SetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm_flt1cr1Type) SetRcont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr1FieldRcontMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldRcontMask)
	}
}

const (
	RegisterDfsdm_flt1cr1FieldRsyncShift = 19
	RegisterDfsdm_flt1cr1FieldRsyncMask  = 0x80000
)

// GetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm_flt1cr1Type) GetRsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldRsyncMask) != 0
}

// SetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm_flt1cr1Type) SetRsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr1FieldRsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldRsyncMask)
	}
}

const (
	RegisterDfsdm_flt1cr1FieldRdmaenShift = 21
	RegisterDfsdm_flt1cr1FieldRdmaenMask  = 0x200000
)

// GetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm_flt1cr1Type) GetRdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldRdmaenMask) != 0
}

// SetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm_flt1cr1Type) SetRdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr1FieldRdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldRdmaenMask)
	}
}

const (
	RegisterDfsdm_flt1cr1FieldRchShift = 24
	RegisterDfsdm_flt1cr1FieldRchMask  = 0x7000000
)

// GetRch Regular channel selection
func (r *registerDfsdm_flt1cr1Type) GetRch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldRchMask) >> RegisterDfsdm_flt1cr1FieldRchShift)
}

// SetRch Regular channel selection
func (r *registerDfsdm_flt1cr1Type) SetRch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldRchMask)|(uint32(value)<<RegisterDfsdm_flt1cr1FieldRchShift))
}

const (
	RegisterDfsdm_flt1cr1FieldFastShift = 29
	RegisterDfsdm_flt1cr1FieldFastMask  = 0x20000000
)

// GetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm_flt1cr1Type) GetFast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldFastMask) != 0
}

// SetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm_flt1cr1Type) SetFast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr1FieldFastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldFastMask)
	}
}

const (
	RegisterDfsdm_flt1cr1FieldAwfselShift = 30
	RegisterDfsdm_flt1cr1FieldAwfselMask  = 0x40000000
)

// GetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm_flt1cr1Type) GetAwfsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr1FieldAwfselMask) != 0
}

// SetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm_flt1cr1Type) SetAwfsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr1FieldAwfselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr1FieldAwfselMask)
	}
}

// registerDfsdm_flt1cr2Type control register 2
type registerDfsdm_flt1cr2Type uint32

const (
	RegisterDfsdm_flt1cr2FieldJeocieShift = 0
	RegisterDfsdm_flt1cr2FieldJeocieMask  = 0x1
)

// GetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm_flt1cr2Type) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr2FieldJeocieMask) != 0
}

// SetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm_flt1cr2Type) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr2FieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr2FieldJeocieMask)
	}
}

const (
	RegisterDfsdm_flt1cr2FieldReocieShift = 1
	RegisterDfsdm_flt1cr2FieldReocieMask  = 0x2
)

// GetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm_flt1cr2Type) GetReocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr2FieldReocieMask) != 0
}

// SetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm_flt1cr2Type) SetReocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr2FieldReocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr2FieldReocieMask)
	}
}

const (
	RegisterDfsdm_flt1cr2FieldJovrieShift = 2
	RegisterDfsdm_flt1cr2FieldJovrieMask  = 0x4
)

// GetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm_flt1cr2Type) GetJovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr2FieldJovrieMask) != 0
}

// SetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm_flt1cr2Type) SetJovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr2FieldJovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr2FieldJovrieMask)
	}
}

const (
	RegisterDfsdm_flt1cr2FieldRovrieShift = 3
	RegisterDfsdm_flt1cr2FieldRovrieMask  = 0x8
)

// GetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm_flt1cr2Type) GetRovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr2FieldRovrieMask) != 0
}

// SetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm_flt1cr2Type) SetRovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr2FieldRovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr2FieldRovrieMask)
	}
}

const (
	RegisterDfsdm_flt1cr2FieldAwdieShift = 4
	RegisterDfsdm_flt1cr2FieldAwdieMask  = 0x10
)

// GetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm_flt1cr2Type) GetAwdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr2FieldAwdieMask) != 0
}

// SetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm_flt1cr2Type) SetAwdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr2FieldAwdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr2FieldAwdieMask)
	}
}

const (
	RegisterDfsdm_flt1cr2FieldScdieShift = 5
	RegisterDfsdm_flt1cr2FieldScdieMask  = 0x20
)

// GetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm_flt1cr2Type) GetScdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr2FieldScdieMask) != 0
}

// SetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm_flt1cr2Type) SetScdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr2FieldScdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr2FieldScdieMask)
	}
}

const (
	RegisterDfsdm_flt1cr2FieldCkabieShift = 6
	RegisterDfsdm_flt1cr2FieldCkabieMask  = 0x40
)

// GetCkabie Clock absence interrupt enable
func (r *registerDfsdm_flt1cr2Type) GetCkabie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr2FieldCkabieMask) != 0
}

// SetCkabie Clock absence interrupt enable
func (r *registerDfsdm_flt1cr2Type) SetCkabie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1cr2FieldCkabieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr2FieldCkabieMask)
	}
}

const (
	RegisterDfsdm_flt1cr2FieldExchShift = 8
	RegisterDfsdm_flt1cr2FieldExchMask  = 0xff00
)

// GetExch Extremes detector channel selection
func (r *registerDfsdm_flt1cr2Type) GetExch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr2FieldExchMask) >> RegisterDfsdm_flt1cr2FieldExchShift)
}

// SetExch Extremes detector channel selection
func (r *registerDfsdm_flt1cr2Type) SetExch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr2FieldExchMask)|(uint32(value)<<RegisterDfsdm_flt1cr2FieldExchShift))
}

const (
	RegisterDfsdm_flt1cr2FieldAwdchShift = 16
	RegisterDfsdm_flt1cr2FieldAwdchMask  = 0xff0000
)

// GetAwdch Analog watchdog channel selection
func (r *registerDfsdm_flt1cr2Type) GetAwdch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cr2FieldAwdchMask) >> RegisterDfsdm_flt1cr2FieldAwdchShift)
}

// SetAwdch Analog watchdog channel selection
func (r *registerDfsdm_flt1cr2Type) SetAwdch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cr2FieldAwdchMask)|(uint32(value)<<RegisterDfsdm_flt1cr2FieldAwdchShift))
}

// registerDfsdm_flt1isrType interrupt and status register
type registerDfsdm_flt1isrType uint32

const (
	RegisterDfsdm_flt1isrFieldJeocfShift = 0
	RegisterDfsdm_flt1isrFieldJeocfMask  = 0x1
)

// GetJeocf End of injected conversion flag
func (r *registerDfsdm_flt1isrType) GetJeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1isrFieldJeocfMask) != 0
}

// SetJeocf End of injected conversion flag
func (r *registerDfsdm_flt1isrType) SetJeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1isrFieldJeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1isrFieldJeocfMask)
	}
}

const (
	RegisterDfsdm_flt1isrFieldReocfShift = 1
	RegisterDfsdm_flt1isrFieldReocfMask  = 0x2
)

// GetReocf End of regular conversion flag
func (r *registerDfsdm_flt1isrType) GetReocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1isrFieldReocfMask) != 0
}

// SetReocf End of regular conversion flag
func (r *registerDfsdm_flt1isrType) SetReocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1isrFieldReocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1isrFieldReocfMask)
	}
}

const (
	RegisterDfsdm_flt1isrFieldJovrfShift = 2
	RegisterDfsdm_flt1isrFieldJovrfMask  = 0x4
)

// GetJovrf Injected conversion overrun flag
func (r *registerDfsdm_flt1isrType) GetJovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1isrFieldJovrfMask) != 0
}

// SetJovrf Injected conversion overrun flag
func (r *registerDfsdm_flt1isrType) SetJovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1isrFieldJovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1isrFieldJovrfMask)
	}
}

const (
	RegisterDfsdm_flt1isrFieldRovrfShift = 3
	RegisterDfsdm_flt1isrFieldRovrfMask  = 0x8
)

// GetRovrf Regular conversion overrun flag
func (r *registerDfsdm_flt1isrType) GetRovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1isrFieldRovrfMask) != 0
}

// SetRovrf Regular conversion overrun flag
func (r *registerDfsdm_flt1isrType) SetRovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1isrFieldRovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1isrFieldRovrfMask)
	}
}

const (
	RegisterDfsdm_flt1isrFieldAwdfShift = 4
	RegisterDfsdm_flt1isrFieldAwdfMask  = 0x10
)

// GetAwdf Analog watchdog
func (r *registerDfsdm_flt1isrType) GetAwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1isrFieldAwdfMask) != 0
}

// SetAwdf Analog watchdog
func (r *registerDfsdm_flt1isrType) SetAwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1isrFieldAwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1isrFieldAwdfMask)
	}
}

const (
	RegisterDfsdm_flt1isrFieldJcipShift = 13
	RegisterDfsdm_flt1isrFieldJcipMask  = 0x2000
)

// GetJcip Injected conversion in progress status
func (r *registerDfsdm_flt1isrType) GetJcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1isrFieldJcipMask) != 0
}

// SetJcip Injected conversion in progress status
func (r *registerDfsdm_flt1isrType) SetJcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1isrFieldJcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1isrFieldJcipMask)
	}
}

const (
	RegisterDfsdm_flt1isrFieldRcipShift = 14
	RegisterDfsdm_flt1isrFieldRcipMask  = 0x4000
)

// GetRcip Regular conversion in progress status
func (r *registerDfsdm_flt1isrType) GetRcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1isrFieldRcipMask) != 0
}

// SetRcip Regular conversion in progress status
func (r *registerDfsdm_flt1isrType) SetRcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1isrFieldRcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1isrFieldRcipMask)
	}
}

const (
	RegisterDfsdm_flt1isrFieldCkabfShift = 16
	RegisterDfsdm_flt1isrFieldCkabfMask  = 0xff0000
)

// GetCkabf Clock absence flag
func (r *registerDfsdm_flt1isrType) GetCkabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1isrFieldCkabfMask) >> RegisterDfsdm_flt1isrFieldCkabfShift)
}

// SetCkabf Clock absence flag
func (r *registerDfsdm_flt1isrType) SetCkabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1isrFieldCkabfMask)|(uint32(value)<<RegisterDfsdm_flt1isrFieldCkabfShift))
}

const (
	RegisterDfsdm_flt1isrFieldScdfShift = 24
	RegisterDfsdm_flt1isrFieldScdfMask  = 0xff000000
)

// GetScdf short-circuit detector flag
func (r *registerDfsdm_flt1isrType) GetScdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1isrFieldScdfMask) >> RegisterDfsdm_flt1isrFieldScdfShift)
}

// SetScdf short-circuit detector flag
func (r *registerDfsdm_flt1isrType) SetScdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1isrFieldScdfMask)|(uint32(value)<<RegisterDfsdm_flt1isrFieldScdfShift))
}

// registerDfsdm1_icrType interrupt flag clear register
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

// registerDfsdm_flt1jchgrType injected channel group selection register
type registerDfsdm_flt1jchgrType uint32

const (
	RegisterDfsdm_flt1jchgrFieldJchgShift = 0
	RegisterDfsdm_flt1jchgrFieldJchgMask  = 0xff
)

// GetJchg Injected channel group selection
func (r *registerDfsdm_flt1jchgrType) GetJchg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1jchgrFieldJchgMask) >> RegisterDfsdm_flt1jchgrFieldJchgShift)
}

// SetJchg Injected channel group selection
func (r *registerDfsdm_flt1jchgrType) SetJchg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1jchgrFieldJchgMask)|(uint32(value)<<RegisterDfsdm_flt1jchgrFieldJchgShift))
}

// registerDfsdm1_fcrType filter control register
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

// registerDfsdm_flt1jdatarType data register for injected group
type registerDfsdm_flt1jdatarType uint32

const (
	RegisterDfsdm_flt1jdatarFieldJdatachShift = 0
	RegisterDfsdm_flt1jdatarFieldJdatachMask  = 0x7
)

// GetJdatach Injected channel most recently converted
func (r *registerDfsdm_flt1jdatarType) GetJdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1jdatarFieldJdatachMask) >> RegisterDfsdm_flt1jdatarFieldJdatachShift)
}

// SetJdatach Injected channel most recently converted
func (r *registerDfsdm_flt1jdatarType) SetJdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1jdatarFieldJdatachMask)|(uint32(value)<<RegisterDfsdm_flt1jdatarFieldJdatachShift))
}

const (
	RegisterDfsdm_flt1jdatarFieldJdataShift = 8
	RegisterDfsdm_flt1jdatarFieldJdataMask  = 0xffffff00
)

// GetJdata Injected group conversion data
func (r *registerDfsdm_flt1jdatarType) GetJdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1jdatarFieldJdataMask) >> RegisterDfsdm_flt1jdatarFieldJdataShift)
}

// SetJdata Injected group conversion data
func (r *registerDfsdm_flt1jdatarType) SetJdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1jdatarFieldJdataMask)|(uint32(value)<<RegisterDfsdm_flt1jdatarFieldJdataShift))
}

// registerDfsdm_flt1rdatarType data register for the regular channel
type registerDfsdm_flt1rdatarType uint32

const (
	RegisterDfsdm_flt1rdatarFieldRdatachShift = 0
	RegisterDfsdm_flt1rdatarFieldRdatachMask  = 0x7
)

// GetRdatach Regular channel most recently converted
func (r *registerDfsdm_flt1rdatarType) GetRdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1rdatarFieldRdatachMask) >> RegisterDfsdm_flt1rdatarFieldRdatachShift)
}

// SetRdatach Regular channel most recently converted
func (r *registerDfsdm_flt1rdatarType) SetRdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1rdatarFieldRdatachMask)|(uint32(value)<<RegisterDfsdm_flt1rdatarFieldRdatachShift))
}

const (
	RegisterDfsdm_flt1rdatarFieldRpendShift = 4
	RegisterDfsdm_flt1rdatarFieldRpendMask  = 0x10
)

// GetRpend Regular channel pending data
func (r *registerDfsdm_flt1rdatarType) GetRpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1rdatarFieldRpendMask) != 0
}

// SetRpend Regular channel pending data
func (r *registerDfsdm_flt1rdatarType) SetRpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt1rdatarFieldRpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1rdatarFieldRpendMask)
	}
}

const (
	RegisterDfsdm_flt1rdatarFieldRdataShift = 8
	RegisterDfsdm_flt1rdatarFieldRdataMask  = 0xffffff00
)

// GetRdata Regular channel conversion data
func (r *registerDfsdm_flt1rdatarType) GetRdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1rdatarFieldRdataMask) >> RegisterDfsdm_flt1rdatarFieldRdataShift)
}

// SetRdata Regular channel conversion data
func (r *registerDfsdm_flt1rdatarType) SetRdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1rdatarFieldRdataMask)|(uint32(value)<<RegisterDfsdm_flt1rdatarFieldRdataShift))
}

// registerDfsdm_flt1awhtrType analog watchdog high threshold register
type registerDfsdm_flt1awhtrType uint32

const (
	RegisterDfsdm_flt1awhtrFieldBkawhShift = 0
	RegisterDfsdm_flt1awhtrFieldBkawhMask  = 0xf
)

// GetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm_flt1awhtrType) GetBkawh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1awhtrFieldBkawhMask) >> RegisterDfsdm_flt1awhtrFieldBkawhShift)
}

// SetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm_flt1awhtrType) SetBkawh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1awhtrFieldBkawhMask)|(uint32(value)<<RegisterDfsdm_flt1awhtrFieldBkawhShift))
}

const (
	RegisterDfsdm_flt1awhtrFieldAwhtShift = 8
	RegisterDfsdm_flt1awhtrFieldAwhtMask  = 0xffffff00
)

// GetAwht Analog watchdog high threshold
func (r *registerDfsdm_flt1awhtrType) GetAwht() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1awhtrFieldAwhtMask) >> RegisterDfsdm_flt1awhtrFieldAwhtShift)
}

// SetAwht Analog watchdog high threshold
func (r *registerDfsdm_flt1awhtrType) SetAwht(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1awhtrFieldAwhtMask)|(uint32(value)<<RegisterDfsdm_flt1awhtrFieldAwhtShift))
}

// registerDfsdm_flt1awltrType analog watchdog low threshold register
type registerDfsdm_flt1awltrType uint32

const (
	RegisterDfsdm_flt1awltrFieldBkawlShift = 0
	RegisterDfsdm_flt1awltrFieldBkawlMask  = 0xf
)

// GetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm_flt1awltrType) GetBkawl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1awltrFieldBkawlMask) >> RegisterDfsdm_flt1awltrFieldBkawlShift)
}

// SetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm_flt1awltrType) SetBkawl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1awltrFieldBkawlMask)|(uint32(value)<<RegisterDfsdm_flt1awltrFieldBkawlShift))
}

const (
	RegisterDfsdm_flt1awltrFieldAwltShift = 8
	RegisterDfsdm_flt1awltrFieldAwltMask  = 0xffffff00
)

// GetAwlt Analog watchdog low threshold
func (r *registerDfsdm_flt1awltrType) GetAwlt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1awltrFieldAwltMask) >> RegisterDfsdm_flt1awltrFieldAwltShift)
}

// SetAwlt Analog watchdog low threshold
func (r *registerDfsdm_flt1awltrType) SetAwlt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1awltrFieldAwltMask)|(uint32(value)<<RegisterDfsdm_flt1awltrFieldAwltShift))
}

// registerDfsdm_flt1awsrType analog watchdog status register
type registerDfsdm_flt1awsrType uint32

const (
	RegisterDfsdm_flt1awsrFieldAwltfShift = 0
	RegisterDfsdm_flt1awsrFieldAwltfMask  = 0xff
)

// GetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm_flt1awsrType) GetAwltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1awsrFieldAwltfMask) >> RegisterDfsdm_flt1awsrFieldAwltfShift)
}

// SetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm_flt1awsrType) SetAwltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1awsrFieldAwltfMask)|(uint32(value)<<RegisterDfsdm_flt1awsrFieldAwltfShift))
}

const (
	RegisterDfsdm_flt1awsrFieldAwhtfShift = 8
	RegisterDfsdm_flt1awsrFieldAwhtfMask  = 0xff00
)

// GetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm_flt1awsrType) GetAwhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1awsrFieldAwhtfMask) >> RegisterDfsdm_flt1awsrFieldAwhtfShift)
}

// SetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm_flt1awsrType) SetAwhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1awsrFieldAwhtfMask)|(uint32(value)<<RegisterDfsdm_flt1awsrFieldAwhtfShift))
}

// registerDfsdm_flt1awcfrType analog watchdog clear flag register
type registerDfsdm_flt1awcfrType uint32

const (
	RegisterDfsdm_flt1awcfrFieldClrawltfShift = 0
	RegisterDfsdm_flt1awcfrFieldClrawltfMask  = 0xff
)

// GetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm_flt1awcfrType) GetClrawltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1awcfrFieldClrawltfMask) >> RegisterDfsdm_flt1awcfrFieldClrawltfShift)
}

// SetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm_flt1awcfrType) SetClrawltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1awcfrFieldClrawltfMask)|(uint32(value)<<RegisterDfsdm_flt1awcfrFieldClrawltfShift))
}

const (
	RegisterDfsdm_flt1awcfrFieldClrawhtfShift = 8
	RegisterDfsdm_flt1awcfrFieldClrawhtfMask  = 0xff00
)

// GetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm_flt1awcfrType) GetClrawhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1awcfrFieldClrawhtfMask) >> RegisterDfsdm_flt1awcfrFieldClrawhtfShift)
}

// SetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm_flt1awcfrType) SetClrawhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1awcfrFieldClrawhtfMask)|(uint32(value)<<RegisterDfsdm_flt1awcfrFieldClrawhtfShift))
}

// registerDfsdm_flt1exmaxType Extremes detector maximum register
type registerDfsdm_flt1exmaxType uint32

const (
	RegisterDfsdm_flt1exmaxFieldExmaxchShift = 0
	RegisterDfsdm_flt1exmaxFieldExmaxchMask  = 0x7
)

// GetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm_flt1exmaxType) GetExmaxch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1exmaxFieldExmaxchMask) >> RegisterDfsdm_flt1exmaxFieldExmaxchShift)
}

// SetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm_flt1exmaxType) SetExmaxch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1exmaxFieldExmaxchMask)|(uint32(value)<<RegisterDfsdm_flt1exmaxFieldExmaxchShift))
}

const (
	RegisterDfsdm_flt1exmaxFieldExmaxShift = 8
	RegisterDfsdm_flt1exmaxFieldExmaxMask  = 0xffffff00
)

// GetExmax Extremes detector maximum value
func (r *registerDfsdm_flt1exmaxType) GetExmax() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1exmaxFieldExmaxMask) >> RegisterDfsdm_flt1exmaxFieldExmaxShift)
}

// SetExmax Extremes detector maximum value
func (r *registerDfsdm_flt1exmaxType) SetExmax(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1exmaxFieldExmaxMask)|(uint32(value)<<RegisterDfsdm_flt1exmaxFieldExmaxShift))
}

// registerDfsdm_flt1exminType Extremes detector minimum register
type registerDfsdm_flt1exminType uint32

const (
	RegisterDfsdm_flt1exminFieldExminchShift = 0
	RegisterDfsdm_flt1exminFieldExminchMask  = 0x7
)

// GetExminch Extremes detector minimum data channel
func (r *registerDfsdm_flt1exminType) GetExminch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1exminFieldExminchMask) >> RegisterDfsdm_flt1exminFieldExminchShift)
}

// SetExminch Extremes detector minimum data channel
func (r *registerDfsdm_flt1exminType) SetExminch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1exminFieldExminchMask)|(uint32(value)<<RegisterDfsdm_flt1exminFieldExminchShift))
}

const (
	RegisterDfsdm_flt1exminFieldExminShift = 8
	RegisterDfsdm_flt1exminFieldExminMask  = 0xffffff00
)

// GetExmin EXMIN
func (r *registerDfsdm_flt1exminType) GetExmin() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1exminFieldExminMask) >> RegisterDfsdm_flt1exminFieldExminShift)
}

// SetExmin EXMIN
func (r *registerDfsdm_flt1exminType) SetExmin(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1exminFieldExminMask)|(uint32(value)<<RegisterDfsdm_flt1exminFieldExminShift))
}

// registerDfsdm_flt1cnvtimrType conversion timer register
type registerDfsdm_flt1cnvtimrType uint32

const (
	RegisterDfsdm_flt1cnvtimrFieldCnvcntShift = 4
	RegisterDfsdm_flt1cnvtimrFieldCnvcntMask  = 0xfffffff0
)

// GetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdm_flt1cnvtimrType) GetCnvcnt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt1cnvtimrFieldCnvcntMask) >> RegisterDfsdm_flt1cnvtimrFieldCnvcntShift)
}

// SetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdm_flt1cnvtimrType) SetCnvcnt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt1cnvtimrFieldCnvcntMask)|(uint32(value)<<RegisterDfsdm_flt1cnvtimrFieldCnvcntShift))
}

// registerDfsdm_flt2cr1Type control register 1
type registerDfsdm_flt2cr1Type uint32

const (
	RegisterDfsdm_flt2cr1FieldDfenShift = 0
	RegisterDfsdm_flt2cr1FieldDfenMask  = 0x1
)

// GetDfen DFSDM enable
func (r *registerDfsdm_flt2cr1Type) GetDfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldDfenMask) != 0
}

// SetDfen DFSDM enable
func (r *registerDfsdm_flt2cr1Type) SetDfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr1FieldDfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldDfenMask)
	}
}

const (
	RegisterDfsdm_flt2cr1FieldJswstartShift = 1
	RegisterDfsdm_flt2cr1FieldJswstartMask  = 0x2
)

// GetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm_flt2cr1Type) GetJswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldJswstartMask) != 0
}

// SetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm_flt2cr1Type) SetJswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr1FieldJswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldJswstartMask)
	}
}

const (
	RegisterDfsdm_flt2cr1FieldJsyncShift = 3
	RegisterDfsdm_flt2cr1FieldJsyncMask  = 0x8
)

// GetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm_flt2cr1Type) GetJsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldJsyncMask) != 0
}

// SetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm_flt2cr1Type) SetJsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr1FieldJsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldJsyncMask)
	}
}

const (
	RegisterDfsdm_flt2cr1FieldJscanShift = 4
	RegisterDfsdm_flt2cr1FieldJscanMask  = 0x10
)

// GetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm_flt2cr1Type) GetJscan() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldJscanMask) != 0
}

// SetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm_flt2cr1Type) SetJscan(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr1FieldJscanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldJscanMask)
	}
}

const (
	RegisterDfsdm_flt2cr1FieldJdmaenShift = 5
	RegisterDfsdm_flt2cr1FieldJdmaenMask  = 0x20
)

// GetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm_flt2cr1Type) GetJdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldJdmaenMask) != 0
}

// SetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm_flt2cr1Type) SetJdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr1FieldJdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldJdmaenMask)
	}
}

const (
	RegisterDfsdm_flt2cr1FieldJextselShift = 8
	RegisterDfsdm_flt2cr1FieldJextselMask  = 0x700
)

// GetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm_flt2cr1Type) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldJextselMask) >> RegisterDfsdm_flt2cr1FieldJextselShift)
}

// SetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm_flt2cr1Type) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldJextselMask)|(uint32(value)<<RegisterDfsdm_flt2cr1FieldJextselShift))
}

const (
	RegisterDfsdm_flt2cr1FieldJextenShift = 13
	RegisterDfsdm_flt2cr1FieldJextenMask  = 0x6000
)

// GetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm_flt2cr1Type) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldJextenMask) >> RegisterDfsdm_flt2cr1FieldJextenShift)
}

// SetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm_flt2cr1Type) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldJextenMask)|(uint32(value)<<RegisterDfsdm_flt2cr1FieldJextenShift))
}

const (
	RegisterDfsdm_flt2cr1FieldRswstartShift = 17
	RegisterDfsdm_flt2cr1FieldRswstartMask  = 0x20000
)

// GetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm_flt2cr1Type) GetRswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldRswstartMask) != 0
}

// SetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm_flt2cr1Type) SetRswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr1FieldRswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldRswstartMask)
	}
}

const (
	RegisterDfsdm_flt2cr1FieldRcontShift = 18
	RegisterDfsdm_flt2cr1FieldRcontMask  = 0x40000
)

// GetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm_flt2cr1Type) GetRcont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldRcontMask) != 0
}

// SetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm_flt2cr1Type) SetRcont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr1FieldRcontMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldRcontMask)
	}
}

const (
	RegisterDfsdm_flt2cr1FieldRsyncShift = 19
	RegisterDfsdm_flt2cr1FieldRsyncMask  = 0x80000
)

// GetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm_flt2cr1Type) GetRsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldRsyncMask) != 0
}

// SetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm_flt2cr1Type) SetRsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr1FieldRsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldRsyncMask)
	}
}

const (
	RegisterDfsdm_flt2cr1FieldRdmaenShift = 21
	RegisterDfsdm_flt2cr1FieldRdmaenMask  = 0x200000
)

// GetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm_flt2cr1Type) GetRdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldRdmaenMask) != 0
}

// SetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm_flt2cr1Type) SetRdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr1FieldRdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldRdmaenMask)
	}
}

const (
	RegisterDfsdm_flt2cr1FieldRchShift = 24
	RegisterDfsdm_flt2cr1FieldRchMask  = 0x7000000
)

// GetRch Regular channel selection
func (r *registerDfsdm_flt2cr1Type) GetRch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldRchMask) >> RegisterDfsdm_flt2cr1FieldRchShift)
}

// SetRch Regular channel selection
func (r *registerDfsdm_flt2cr1Type) SetRch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldRchMask)|(uint32(value)<<RegisterDfsdm_flt2cr1FieldRchShift))
}

const (
	RegisterDfsdm_flt2cr1FieldFastShift = 29
	RegisterDfsdm_flt2cr1FieldFastMask  = 0x20000000
)

// GetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm_flt2cr1Type) GetFast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldFastMask) != 0
}

// SetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm_flt2cr1Type) SetFast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr1FieldFastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldFastMask)
	}
}

const (
	RegisterDfsdm_flt2cr1FieldAwfselShift = 30
	RegisterDfsdm_flt2cr1FieldAwfselMask  = 0x40000000
)

// GetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm_flt2cr1Type) GetAwfsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr1FieldAwfselMask) != 0
}

// SetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm_flt2cr1Type) SetAwfsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr1FieldAwfselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr1FieldAwfselMask)
	}
}

// registerDfsdm_flt2cr2Type control register 2
type registerDfsdm_flt2cr2Type uint32

const (
	RegisterDfsdm_flt2cr2FieldJeocieShift = 0
	RegisterDfsdm_flt2cr2FieldJeocieMask  = 0x1
)

// GetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm_flt2cr2Type) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr2FieldJeocieMask) != 0
}

// SetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm_flt2cr2Type) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr2FieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr2FieldJeocieMask)
	}
}

const (
	RegisterDfsdm_flt2cr2FieldReocieShift = 1
	RegisterDfsdm_flt2cr2FieldReocieMask  = 0x2
)

// GetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm_flt2cr2Type) GetReocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr2FieldReocieMask) != 0
}

// SetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm_flt2cr2Type) SetReocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr2FieldReocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr2FieldReocieMask)
	}
}

const (
	RegisterDfsdm_flt2cr2FieldJovrieShift = 2
	RegisterDfsdm_flt2cr2FieldJovrieMask  = 0x4
)

// GetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm_flt2cr2Type) GetJovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr2FieldJovrieMask) != 0
}

// SetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm_flt2cr2Type) SetJovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr2FieldJovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr2FieldJovrieMask)
	}
}

const (
	RegisterDfsdm_flt2cr2FieldRovrieShift = 3
	RegisterDfsdm_flt2cr2FieldRovrieMask  = 0x8
)

// GetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm_flt2cr2Type) GetRovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr2FieldRovrieMask) != 0
}

// SetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm_flt2cr2Type) SetRovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr2FieldRovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr2FieldRovrieMask)
	}
}

const (
	RegisterDfsdm_flt2cr2FieldAwdieShift = 4
	RegisterDfsdm_flt2cr2FieldAwdieMask  = 0x10
)

// GetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm_flt2cr2Type) GetAwdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr2FieldAwdieMask) != 0
}

// SetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm_flt2cr2Type) SetAwdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr2FieldAwdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr2FieldAwdieMask)
	}
}

const (
	RegisterDfsdm_flt2cr2FieldScdieShift = 5
	RegisterDfsdm_flt2cr2FieldScdieMask  = 0x20
)

// GetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm_flt2cr2Type) GetScdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr2FieldScdieMask) != 0
}

// SetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm_flt2cr2Type) SetScdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr2FieldScdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr2FieldScdieMask)
	}
}

const (
	RegisterDfsdm_flt2cr2FieldCkabieShift = 6
	RegisterDfsdm_flt2cr2FieldCkabieMask  = 0x40
)

// GetCkabie Clock absence interrupt enable
func (r *registerDfsdm_flt2cr2Type) GetCkabie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr2FieldCkabieMask) != 0
}

// SetCkabie Clock absence interrupt enable
func (r *registerDfsdm_flt2cr2Type) SetCkabie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2cr2FieldCkabieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr2FieldCkabieMask)
	}
}

const (
	RegisterDfsdm_flt2cr2FieldExchShift = 8
	RegisterDfsdm_flt2cr2FieldExchMask  = 0xff00
)

// GetExch Extremes detector channel selection
func (r *registerDfsdm_flt2cr2Type) GetExch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr2FieldExchMask) >> RegisterDfsdm_flt2cr2FieldExchShift)
}

// SetExch Extremes detector channel selection
func (r *registerDfsdm_flt2cr2Type) SetExch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr2FieldExchMask)|(uint32(value)<<RegisterDfsdm_flt2cr2FieldExchShift))
}

const (
	RegisterDfsdm_flt2cr2FieldAwdchShift = 16
	RegisterDfsdm_flt2cr2FieldAwdchMask  = 0xff0000
)

// GetAwdch Analog watchdog channel selection
func (r *registerDfsdm_flt2cr2Type) GetAwdch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cr2FieldAwdchMask) >> RegisterDfsdm_flt2cr2FieldAwdchShift)
}

// SetAwdch Analog watchdog channel selection
func (r *registerDfsdm_flt2cr2Type) SetAwdch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cr2FieldAwdchMask)|(uint32(value)<<RegisterDfsdm_flt2cr2FieldAwdchShift))
}

// registerDfsdm_flt2isrType interrupt and status register
type registerDfsdm_flt2isrType uint32

const (
	RegisterDfsdm_flt2isrFieldJeocfShift = 0
	RegisterDfsdm_flt2isrFieldJeocfMask  = 0x1
)

// GetJeocf End of injected conversion flag
func (r *registerDfsdm_flt2isrType) GetJeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2isrFieldJeocfMask) != 0
}

// SetJeocf End of injected conversion flag
func (r *registerDfsdm_flt2isrType) SetJeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2isrFieldJeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2isrFieldJeocfMask)
	}
}

const (
	RegisterDfsdm_flt2isrFieldReocfShift = 1
	RegisterDfsdm_flt2isrFieldReocfMask  = 0x2
)

// GetReocf End of regular conversion flag
func (r *registerDfsdm_flt2isrType) GetReocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2isrFieldReocfMask) != 0
}

// SetReocf End of regular conversion flag
func (r *registerDfsdm_flt2isrType) SetReocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2isrFieldReocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2isrFieldReocfMask)
	}
}

const (
	RegisterDfsdm_flt2isrFieldJovrfShift = 2
	RegisterDfsdm_flt2isrFieldJovrfMask  = 0x4
)

// GetJovrf Injected conversion overrun flag
func (r *registerDfsdm_flt2isrType) GetJovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2isrFieldJovrfMask) != 0
}

// SetJovrf Injected conversion overrun flag
func (r *registerDfsdm_flt2isrType) SetJovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2isrFieldJovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2isrFieldJovrfMask)
	}
}

const (
	RegisterDfsdm_flt2isrFieldRovrfShift = 3
	RegisterDfsdm_flt2isrFieldRovrfMask  = 0x8
)

// GetRovrf Regular conversion overrun flag
func (r *registerDfsdm_flt2isrType) GetRovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2isrFieldRovrfMask) != 0
}

// SetRovrf Regular conversion overrun flag
func (r *registerDfsdm_flt2isrType) SetRovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2isrFieldRovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2isrFieldRovrfMask)
	}
}

const (
	RegisterDfsdm_flt2isrFieldAwdfShift = 4
	RegisterDfsdm_flt2isrFieldAwdfMask  = 0x10
)

// GetAwdf Analog watchdog
func (r *registerDfsdm_flt2isrType) GetAwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2isrFieldAwdfMask) != 0
}

// SetAwdf Analog watchdog
func (r *registerDfsdm_flt2isrType) SetAwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2isrFieldAwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2isrFieldAwdfMask)
	}
}

const (
	RegisterDfsdm_flt2isrFieldJcipShift = 13
	RegisterDfsdm_flt2isrFieldJcipMask  = 0x2000
)

// GetJcip Injected conversion in progress status
func (r *registerDfsdm_flt2isrType) GetJcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2isrFieldJcipMask) != 0
}

// SetJcip Injected conversion in progress status
func (r *registerDfsdm_flt2isrType) SetJcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2isrFieldJcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2isrFieldJcipMask)
	}
}

const (
	RegisterDfsdm_flt2isrFieldRcipShift = 14
	RegisterDfsdm_flt2isrFieldRcipMask  = 0x4000
)

// GetRcip Regular conversion in progress status
func (r *registerDfsdm_flt2isrType) GetRcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2isrFieldRcipMask) != 0
}

// SetRcip Regular conversion in progress status
func (r *registerDfsdm_flt2isrType) SetRcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2isrFieldRcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2isrFieldRcipMask)
	}
}

const (
	RegisterDfsdm_flt2isrFieldCkabfShift = 16
	RegisterDfsdm_flt2isrFieldCkabfMask  = 0xff0000
)

// GetCkabf Clock absence flag
func (r *registerDfsdm_flt2isrType) GetCkabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2isrFieldCkabfMask) >> RegisterDfsdm_flt2isrFieldCkabfShift)
}

// SetCkabf Clock absence flag
func (r *registerDfsdm_flt2isrType) SetCkabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2isrFieldCkabfMask)|(uint32(value)<<RegisterDfsdm_flt2isrFieldCkabfShift))
}

const (
	RegisterDfsdm_flt2isrFieldScdfShift = 24
	RegisterDfsdm_flt2isrFieldScdfMask  = 0xff000000
)

// GetScdf short-circuit detector flag
func (r *registerDfsdm_flt2isrType) GetScdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2isrFieldScdfMask) >> RegisterDfsdm_flt2isrFieldScdfShift)
}

// SetScdf short-circuit detector flag
func (r *registerDfsdm_flt2isrType) SetScdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2isrFieldScdfMask)|(uint32(value)<<RegisterDfsdm_flt2isrFieldScdfShift))
}

// registerDfsdm_flt2icrType interrupt flag clear register
type registerDfsdm_flt2icrType uint32

const (
	RegisterDfsdm_flt2icrFieldClrjovrfShift = 2
	RegisterDfsdm_flt2icrFieldClrjovrfMask  = 0x4
)

// GetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm_flt2icrType) GetClrjovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2icrFieldClrjovrfMask) != 0
}

// SetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm_flt2icrType) SetClrjovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2icrFieldClrjovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2icrFieldClrjovrfMask)
	}
}

const (
	RegisterDfsdm_flt2icrFieldClrrovrfShift = 3
	RegisterDfsdm_flt2icrFieldClrrovrfMask  = 0x8
)

// GetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm_flt2icrType) GetClrrovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2icrFieldClrrovrfMask) != 0
}

// SetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm_flt2icrType) SetClrrovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2icrFieldClrrovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2icrFieldClrrovrfMask)
	}
}

const (
	RegisterDfsdm_flt2icrFieldClrckabfShift = 16
	RegisterDfsdm_flt2icrFieldClrckabfMask  = 0xff0000
)

// GetClrckabf Clear the clock absence flag
func (r *registerDfsdm_flt2icrType) GetClrckabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2icrFieldClrckabfMask) >> RegisterDfsdm_flt2icrFieldClrckabfShift)
}

// SetClrckabf Clear the clock absence flag
func (r *registerDfsdm_flt2icrType) SetClrckabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2icrFieldClrckabfMask)|(uint32(value)<<RegisterDfsdm_flt2icrFieldClrckabfShift))
}

const (
	RegisterDfsdm_flt2icrFieldClrscdfShift = 24
	RegisterDfsdm_flt2icrFieldClrscdfMask  = 0xff000000
)

// GetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm_flt2icrType) GetClrscdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2icrFieldClrscdfMask) >> RegisterDfsdm_flt2icrFieldClrscdfShift)
}

// SetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm_flt2icrType) SetClrscdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2icrFieldClrscdfMask)|(uint32(value)<<RegisterDfsdm_flt2icrFieldClrscdfShift))
}

// registerDfsdm_flt2jchgrType injected channel group selection register
type registerDfsdm_flt2jchgrType uint32

const (
	RegisterDfsdm_flt2jchgrFieldJchgShift = 0
	RegisterDfsdm_flt2jchgrFieldJchgMask  = 0xff
)

// GetJchg Injected channel group selection
func (r *registerDfsdm_flt2jchgrType) GetJchg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2jchgrFieldJchgMask) >> RegisterDfsdm_flt2jchgrFieldJchgShift)
}

// SetJchg Injected channel group selection
func (r *registerDfsdm_flt2jchgrType) SetJchg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2jchgrFieldJchgMask)|(uint32(value)<<RegisterDfsdm_flt2jchgrFieldJchgShift))
}

// registerDfsdm_flt2fcrType filter control register
type registerDfsdm_flt2fcrType uint32

const (
	RegisterDfsdm_flt2fcrFieldIosrShift = 0
	RegisterDfsdm_flt2fcrFieldIosrMask  = 0xff
)

// GetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm_flt2fcrType) GetIosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2fcrFieldIosrMask) >> RegisterDfsdm_flt2fcrFieldIosrShift)
}

// SetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm_flt2fcrType) SetIosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2fcrFieldIosrMask)|(uint32(value)<<RegisterDfsdm_flt2fcrFieldIosrShift))
}

const (
	RegisterDfsdm_flt2fcrFieldFosrShift = 16
	RegisterDfsdm_flt2fcrFieldFosrMask  = 0x3ff0000
)

// GetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm_flt2fcrType) GetFosr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2fcrFieldFosrMask) >> RegisterDfsdm_flt2fcrFieldFosrShift)
}

// SetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm_flt2fcrType) SetFosr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2fcrFieldFosrMask)|(uint32(value)<<RegisterDfsdm_flt2fcrFieldFosrShift))
}

const (
	RegisterDfsdm_flt2fcrFieldFordShift = 29
	RegisterDfsdm_flt2fcrFieldFordMask  = 0xe0000000
)

// GetFord Sinc filter order
func (r *registerDfsdm_flt2fcrType) GetFord() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2fcrFieldFordMask) >> RegisterDfsdm_flt2fcrFieldFordShift)
}

// SetFord Sinc filter order
func (r *registerDfsdm_flt2fcrType) SetFord(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2fcrFieldFordMask)|(uint32(value)<<RegisterDfsdm_flt2fcrFieldFordShift))
}

// registerDfsdm_flt2jdatarType data register for injected group
type registerDfsdm_flt2jdatarType uint32

const (
	RegisterDfsdm_flt2jdatarFieldJdatachShift = 0
	RegisterDfsdm_flt2jdatarFieldJdatachMask  = 0x7
)

// GetJdatach Injected channel most recently converted
func (r *registerDfsdm_flt2jdatarType) GetJdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2jdatarFieldJdatachMask) >> RegisterDfsdm_flt2jdatarFieldJdatachShift)
}

// SetJdatach Injected channel most recently converted
func (r *registerDfsdm_flt2jdatarType) SetJdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2jdatarFieldJdatachMask)|(uint32(value)<<RegisterDfsdm_flt2jdatarFieldJdatachShift))
}

const (
	RegisterDfsdm_flt2jdatarFieldJdataShift = 8
	RegisterDfsdm_flt2jdatarFieldJdataMask  = 0xffffff00
)

// GetJdata Injected group conversion data
func (r *registerDfsdm_flt2jdatarType) GetJdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2jdatarFieldJdataMask) >> RegisterDfsdm_flt2jdatarFieldJdataShift)
}

// SetJdata Injected group conversion data
func (r *registerDfsdm_flt2jdatarType) SetJdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2jdatarFieldJdataMask)|(uint32(value)<<RegisterDfsdm_flt2jdatarFieldJdataShift))
}

// registerDfsdm_flt2rdatarType data register for the regular channel
type registerDfsdm_flt2rdatarType uint32

const (
	RegisterDfsdm_flt2rdatarFieldRdatachShift = 0
	RegisterDfsdm_flt2rdatarFieldRdatachMask  = 0x7
)

// GetRdatach Regular channel most recently converted
func (r *registerDfsdm_flt2rdatarType) GetRdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2rdatarFieldRdatachMask) >> RegisterDfsdm_flt2rdatarFieldRdatachShift)
}

// SetRdatach Regular channel most recently converted
func (r *registerDfsdm_flt2rdatarType) SetRdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2rdatarFieldRdatachMask)|(uint32(value)<<RegisterDfsdm_flt2rdatarFieldRdatachShift))
}

const (
	RegisterDfsdm_flt2rdatarFieldRpendShift = 4
	RegisterDfsdm_flt2rdatarFieldRpendMask  = 0x10
)

// GetRpend Regular channel pending data
func (r *registerDfsdm_flt2rdatarType) GetRpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2rdatarFieldRpendMask) != 0
}

// SetRpend Regular channel pending data
func (r *registerDfsdm_flt2rdatarType) SetRpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt2rdatarFieldRpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2rdatarFieldRpendMask)
	}
}

const (
	RegisterDfsdm_flt2rdatarFieldRdataShift = 8
	RegisterDfsdm_flt2rdatarFieldRdataMask  = 0xffffff00
)

// GetRdata Regular channel conversion data
func (r *registerDfsdm_flt2rdatarType) GetRdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2rdatarFieldRdataMask) >> RegisterDfsdm_flt2rdatarFieldRdataShift)
}

// SetRdata Regular channel conversion data
func (r *registerDfsdm_flt2rdatarType) SetRdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2rdatarFieldRdataMask)|(uint32(value)<<RegisterDfsdm_flt2rdatarFieldRdataShift))
}

// registerDfsdm_flt2awhtrType analog watchdog high threshold register
type registerDfsdm_flt2awhtrType uint32

const (
	RegisterDfsdm_flt2awhtrFieldBkawhShift = 0
	RegisterDfsdm_flt2awhtrFieldBkawhMask  = 0xf
)

// GetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm_flt2awhtrType) GetBkawh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2awhtrFieldBkawhMask) >> RegisterDfsdm_flt2awhtrFieldBkawhShift)
}

// SetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm_flt2awhtrType) SetBkawh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2awhtrFieldBkawhMask)|(uint32(value)<<RegisterDfsdm_flt2awhtrFieldBkawhShift))
}

const (
	RegisterDfsdm_flt2awhtrFieldAwhtShift = 8
	RegisterDfsdm_flt2awhtrFieldAwhtMask  = 0xffffff00
)

// GetAwht Analog watchdog high threshold
func (r *registerDfsdm_flt2awhtrType) GetAwht() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2awhtrFieldAwhtMask) >> RegisterDfsdm_flt2awhtrFieldAwhtShift)
}

// SetAwht Analog watchdog high threshold
func (r *registerDfsdm_flt2awhtrType) SetAwht(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2awhtrFieldAwhtMask)|(uint32(value)<<RegisterDfsdm_flt2awhtrFieldAwhtShift))
}

// registerDfsdm_flt2awltrType analog watchdog low threshold register
type registerDfsdm_flt2awltrType uint32

const (
	RegisterDfsdm_flt2awltrFieldBkawlShift = 0
	RegisterDfsdm_flt2awltrFieldBkawlMask  = 0xf
)

// GetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm_flt2awltrType) GetBkawl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2awltrFieldBkawlMask) >> RegisterDfsdm_flt2awltrFieldBkawlShift)
}

// SetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm_flt2awltrType) SetBkawl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2awltrFieldBkawlMask)|(uint32(value)<<RegisterDfsdm_flt2awltrFieldBkawlShift))
}

const (
	RegisterDfsdm_flt2awltrFieldAwltShift = 8
	RegisterDfsdm_flt2awltrFieldAwltMask  = 0xffffff00
)

// GetAwlt Analog watchdog low threshold
func (r *registerDfsdm_flt2awltrType) GetAwlt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2awltrFieldAwltMask) >> RegisterDfsdm_flt2awltrFieldAwltShift)
}

// SetAwlt Analog watchdog low threshold
func (r *registerDfsdm_flt2awltrType) SetAwlt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2awltrFieldAwltMask)|(uint32(value)<<RegisterDfsdm_flt2awltrFieldAwltShift))
}

// registerDfsdm_flt2awsrType analog watchdog status register
type registerDfsdm_flt2awsrType uint32

const (
	RegisterDfsdm_flt2awsrFieldAwltfShift = 0
	RegisterDfsdm_flt2awsrFieldAwltfMask  = 0xff
)

// GetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm_flt2awsrType) GetAwltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2awsrFieldAwltfMask) >> RegisterDfsdm_flt2awsrFieldAwltfShift)
}

// SetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm_flt2awsrType) SetAwltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2awsrFieldAwltfMask)|(uint32(value)<<RegisterDfsdm_flt2awsrFieldAwltfShift))
}

const (
	RegisterDfsdm_flt2awsrFieldAwhtfShift = 8
	RegisterDfsdm_flt2awsrFieldAwhtfMask  = 0xff00
)

// GetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm_flt2awsrType) GetAwhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2awsrFieldAwhtfMask) >> RegisterDfsdm_flt2awsrFieldAwhtfShift)
}

// SetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm_flt2awsrType) SetAwhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2awsrFieldAwhtfMask)|(uint32(value)<<RegisterDfsdm_flt2awsrFieldAwhtfShift))
}

// registerDfsdm_flt2awcfrType analog watchdog clear flag register
type registerDfsdm_flt2awcfrType uint32

const (
	RegisterDfsdm_flt2awcfrFieldClrawltfShift = 0
	RegisterDfsdm_flt2awcfrFieldClrawltfMask  = 0xff
)

// GetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm_flt2awcfrType) GetClrawltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2awcfrFieldClrawltfMask) >> RegisterDfsdm_flt2awcfrFieldClrawltfShift)
}

// SetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm_flt2awcfrType) SetClrawltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2awcfrFieldClrawltfMask)|(uint32(value)<<RegisterDfsdm_flt2awcfrFieldClrawltfShift))
}

const (
	RegisterDfsdm_flt2awcfrFieldClrawhtfShift = 8
	RegisterDfsdm_flt2awcfrFieldClrawhtfMask  = 0xff00
)

// GetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm_flt2awcfrType) GetClrawhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2awcfrFieldClrawhtfMask) >> RegisterDfsdm_flt2awcfrFieldClrawhtfShift)
}

// SetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm_flt2awcfrType) SetClrawhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2awcfrFieldClrawhtfMask)|(uint32(value)<<RegisterDfsdm_flt2awcfrFieldClrawhtfShift))
}

// registerDfsdm_flt2exmaxType Extremes detector maximum register
type registerDfsdm_flt2exmaxType uint32

const (
	RegisterDfsdm_flt2exmaxFieldExmaxchShift = 0
	RegisterDfsdm_flt2exmaxFieldExmaxchMask  = 0x7
)

// GetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm_flt2exmaxType) GetExmaxch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2exmaxFieldExmaxchMask) >> RegisterDfsdm_flt2exmaxFieldExmaxchShift)
}

// SetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm_flt2exmaxType) SetExmaxch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2exmaxFieldExmaxchMask)|(uint32(value)<<RegisterDfsdm_flt2exmaxFieldExmaxchShift))
}

const (
	RegisterDfsdm_flt2exmaxFieldExmaxShift = 8
	RegisterDfsdm_flt2exmaxFieldExmaxMask  = 0xffffff00
)

// GetExmax Extremes detector maximum value
func (r *registerDfsdm_flt2exmaxType) GetExmax() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2exmaxFieldExmaxMask) >> RegisterDfsdm_flt2exmaxFieldExmaxShift)
}

// SetExmax Extremes detector maximum value
func (r *registerDfsdm_flt2exmaxType) SetExmax(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2exmaxFieldExmaxMask)|(uint32(value)<<RegisterDfsdm_flt2exmaxFieldExmaxShift))
}

// registerDfsdm_flt2exminType Extremes detector minimum register
type registerDfsdm_flt2exminType uint32

const (
	RegisterDfsdm_flt2exminFieldExminchShift = 0
	RegisterDfsdm_flt2exminFieldExminchMask  = 0x7
)

// GetExminch Extremes detector minimum data channel
func (r *registerDfsdm_flt2exminType) GetExminch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2exminFieldExminchMask) >> RegisterDfsdm_flt2exminFieldExminchShift)
}

// SetExminch Extremes detector minimum data channel
func (r *registerDfsdm_flt2exminType) SetExminch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2exminFieldExminchMask)|(uint32(value)<<RegisterDfsdm_flt2exminFieldExminchShift))
}

const (
	RegisterDfsdm_flt2exminFieldExminShift = 8
	RegisterDfsdm_flt2exminFieldExminMask  = 0xffffff00
)

// GetExmin EXMIN
func (r *registerDfsdm_flt2exminType) GetExmin() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2exminFieldExminMask) >> RegisterDfsdm_flt2exminFieldExminShift)
}

// SetExmin EXMIN
func (r *registerDfsdm_flt2exminType) SetExmin(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2exminFieldExminMask)|(uint32(value)<<RegisterDfsdm_flt2exminFieldExminShift))
}

// registerDfsdm_flt2cnvtimrType conversion timer register
type registerDfsdm_flt2cnvtimrType uint32

const (
	RegisterDfsdm_flt2cnvtimrFieldCnvcntShift = 4
	RegisterDfsdm_flt2cnvtimrFieldCnvcntMask  = 0xfffffff0
)

// GetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdm_flt2cnvtimrType) GetCnvcnt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt2cnvtimrFieldCnvcntMask) >> RegisterDfsdm_flt2cnvtimrFieldCnvcntShift)
}

// SetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdm_flt2cnvtimrType) SetCnvcnt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt2cnvtimrFieldCnvcntMask)|(uint32(value)<<RegisterDfsdm_flt2cnvtimrFieldCnvcntShift))
}

// registerDfsdm_flt3cr1Type control register 1
type registerDfsdm_flt3cr1Type uint32

const (
	RegisterDfsdm_flt3cr1FieldDfenShift = 0
	RegisterDfsdm_flt3cr1FieldDfenMask  = 0x1
)

// GetDfen DFSDM enable
func (r *registerDfsdm_flt3cr1Type) GetDfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldDfenMask) != 0
}

// SetDfen DFSDM enable
func (r *registerDfsdm_flt3cr1Type) SetDfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr1FieldDfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldDfenMask)
	}
}

const (
	RegisterDfsdm_flt3cr1FieldJswstartShift = 1
	RegisterDfsdm_flt3cr1FieldJswstartMask  = 0x2
)

// GetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm_flt3cr1Type) GetJswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldJswstartMask) != 0
}

// SetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdm_flt3cr1Type) SetJswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr1FieldJswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldJswstartMask)
	}
}

const (
	RegisterDfsdm_flt3cr1FieldJsyncShift = 3
	RegisterDfsdm_flt3cr1FieldJsyncMask  = 0x8
)

// GetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm_flt3cr1Type) GetJsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldJsyncMask) != 0
}

// SetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdm_flt3cr1Type) SetJsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr1FieldJsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldJsyncMask)
	}
}

const (
	RegisterDfsdm_flt3cr1FieldJscanShift = 4
	RegisterDfsdm_flt3cr1FieldJscanMask  = 0x10
)

// GetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm_flt3cr1Type) GetJscan() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldJscanMask) != 0
}

// SetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdm_flt3cr1Type) SetJscan(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr1FieldJscanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldJscanMask)
	}
}

const (
	RegisterDfsdm_flt3cr1FieldJdmaenShift = 5
	RegisterDfsdm_flt3cr1FieldJdmaenMask  = 0x20
)

// GetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm_flt3cr1Type) GetJdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldJdmaenMask) != 0
}

// SetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdm_flt3cr1Type) SetJdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr1FieldJdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldJdmaenMask)
	}
}

const (
	RegisterDfsdm_flt3cr1FieldJextselShift = 8
	RegisterDfsdm_flt3cr1FieldJextselMask  = 0x700
)

// GetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm_flt3cr1Type) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldJextselMask) >> RegisterDfsdm_flt3cr1FieldJextselShift)
}

// SetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdm_flt3cr1Type) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldJextselMask)|(uint32(value)<<RegisterDfsdm_flt3cr1FieldJextselShift))
}

const (
	RegisterDfsdm_flt3cr1FieldJextenShift = 13
	RegisterDfsdm_flt3cr1FieldJextenMask  = 0x6000
)

// GetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm_flt3cr1Type) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldJextenMask) >> RegisterDfsdm_flt3cr1FieldJextenShift)
}

// SetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdm_flt3cr1Type) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldJextenMask)|(uint32(value)<<RegisterDfsdm_flt3cr1FieldJextenShift))
}

const (
	RegisterDfsdm_flt3cr1FieldRswstartShift = 17
	RegisterDfsdm_flt3cr1FieldRswstartMask  = 0x20000
)

// GetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm_flt3cr1Type) GetRswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldRswstartMask) != 0
}

// SetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdm_flt3cr1Type) SetRswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr1FieldRswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldRswstartMask)
	}
}

const (
	RegisterDfsdm_flt3cr1FieldRcontShift = 18
	RegisterDfsdm_flt3cr1FieldRcontMask  = 0x40000
)

// GetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm_flt3cr1Type) GetRcont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldRcontMask) != 0
}

// SetRcont Continuous mode selection for regular conversions
func (r *registerDfsdm_flt3cr1Type) SetRcont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr1FieldRcontMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldRcontMask)
	}
}

const (
	RegisterDfsdm_flt3cr1FieldRsyncShift = 19
	RegisterDfsdm_flt3cr1FieldRsyncMask  = 0x80000
)

// GetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm_flt3cr1Type) GetRsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldRsyncMask) != 0
}

// SetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdm_flt3cr1Type) SetRsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr1FieldRsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldRsyncMask)
	}
}

const (
	RegisterDfsdm_flt3cr1FieldRdmaenShift = 21
	RegisterDfsdm_flt3cr1FieldRdmaenMask  = 0x200000
)

// GetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm_flt3cr1Type) GetRdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldRdmaenMask) != 0
}

// SetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdm_flt3cr1Type) SetRdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr1FieldRdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldRdmaenMask)
	}
}

const (
	RegisterDfsdm_flt3cr1FieldRchShift = 24
	RegisterDfsdm_flt3cr1FieldRchMask  = 0x7000000
)

// GetRch Regular channel selection
func (r *registerDfsdm_flt3cr1Type) GetRch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldRchMask) >> RegisterDfsdm_flt3cr1FieldRchShift)
}

// SetRch Regular channel selection
func (r *registerDfsdm_flt3cr1Type) SetRch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldRchMask)|(uint32(value)<<RegisterDfsdm_flt3cr1FieldRchShift))
}

const (
	RegisterDfsdm_flt3cr1FieldFastShift = 29
	RegisterDfsdm_flt3cr1FieldFastMask  = 0x20000000
)

// GetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm_flt3cr1Type) GetFast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldFastMask) != 0
}

// SetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdm_flt3cr1Type) SetFast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr1FieldFastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldFastMask)
	}
}

const (
	RegisterDfsdm_flt3cr1FieldAwfselShift = 30
	RegisterDfsdm_flt3cr1FieldAwfselMask  = 0x40000000
)

// GetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm_flt3cr1Type) GetAwfsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr1FieldAwfselMask) != 0
}

// SetAwfsel Analog watchdog fast mode select
func (r *registerDfsdm_flt3cr1Type) SetAwfsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr1FieldAwfselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr1FieldAwfselMask)
	}
}

// registerDfsdm_flt3cr2Type control register 2
type registerDfsdm_flt3cr2Type uint32

const (
	RegisterDfsdm_flt3cr2FieldJeocieShift = 0
	RegisterDfsdm_flt3cr2FieldJeocieMask  = 0x1
)

// GetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm_flt3cr2Type) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr2FieldJeocieMask) != 0
}

// SetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdm_flt3cr2Type) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr2FieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr2FieldJeocieMask)
	}
}

const (
	RegisterDfsdm_flt3cr2FieldReocieShift = 1
	RegisterDfsdm_flt3cr2FieldReocieMask  = 0x2
)

// GetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm_flt3cr2Type) GetReocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr2FieldReocieMask) != 0
}

// SetReocie Regular end of conversion interrupt enable
func (r *registerDfsdm_flt3cr2Type) SetReocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr2FieldReocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr2FieldReocieMask)
	}
}

const (
	RegisterDfsdm_flt3cr2FieldJovrieShift = 2
	RegisterDfsdm_flt3cr2FieldJovrieMask  = 0x4
)

// GetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm_flt3cr2Type) GetJovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr2FieldJovrieMask) != 0
}

// SetJovrie Injected data overrun interrupt enable
func (r *registerDfsdm_flt3cr2Type) SetJovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr2FieldJovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr2FieldJovrieMask)
	}
}

const (
	RegisterDfsdm_flt3cr2FieldRovrieShift = 3
	RegisterDfsdm_flt3cr2FieldRovrieMask  = 0x8
)

// GetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm_flt3cr2Type) GetRovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr2FieldRovrieMask) != 0
}

// SetRovrie Regular data overrun interrupt enable
func (r *registerDfsdm_flt3cr2Type) SetRovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr2FieldRovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr2FieldRovrieMask)
	}
}

const (
	RegisterDfsdm_flt3cr2FieldAwdieShift = 4
	RegisterDfsdm_flt3cr2FieldAwdieMask  = 0x10
)

// GetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm_flt3cr2Type) GetAwdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr2FieldAwdieMask) != 0
}

// SetAwdie Analog watchdog interrupt enable
func (r *registerDfsdm_flt3cr2Type) SetAwdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr2FieldAwdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr2FieldAwdieMask)
	}
}

const (
	RegisterDfsdm_flt3cr2FieldScdieShift = 5
	RegisterDfsdm_flt3cr2FieldScdieMask  = 0x20
)

// GetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm_flt3cr2Type) GetScdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr2FieldScdieMask) != 0
}

// SetScdie Short-circuit detector interrupt enable
func (r *registerDfsdm_flt3cr2Type) SetScdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr2FieldScdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr2FieldScdieMask)
	}
}

const (
	RegisterDfsdm_flt3cr2FieldCkabieShift = 6
	RegisterDfsdm_flt3cr2FieldCkabieMask  = 0x40
)

// GetCkabie Clock absence interrupt enable
func (r *registerDfsdm_flt3cr2Type) GetCkabie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr2FieldCkabieMask) != 0
}

// SetCkabie Clock absence interrupt enable
func (r *registerDfsdm_flt3cr2Type) SetCkabie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3cr2FieldCkabieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr2FieldCkabieMask)
	}
}

const (
	RegisterDfsdm_flt3cr2FieldExchShift = 8
	RegisterDfsdm_flt3cr2FieldExchMask  = 0xff00
)

// GetExch Extremes detector channel selection
func (r *registerDfsdm_flt3cr2Type) GetExch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr2FieldExchMask) >> RegisterDfsdm_flt3cr2FieldExchShift)
}

// SetExch Extremes detector channel selection
func (r *registerDfsdm_flt3cr2Type) SetExch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr2FieldExchMask)|(uint32(value)<<RegisterDfsdm_flt3cr2FieldExchShift))
}

const (
	RegisterDfsdm_flt3cr2FieldAwdchShift = 16
	RegisterDfsdm_flt3cr2FieldAwdchMask  = 0xff0000
)

// GetAwdch Analog watchdog channel selection
func (r *registerDfsdm_flt3cr2Type) GetAwdch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cr2FieldAwdchMask) >> RegisterDfsdm_flt3cr2FieldAwdchShift)
}

// SetAwdch Analog watchdog channel selection
func (r *registerDfsdm_flt3cr2Type) SetAwdch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cr2FieldAwdchMask)|(uint32(value)<<RegisterDfsdm_flt3cr2FieldAwdchShift))
}

// registerDfsdm_flt3isrType interrupt and status register
type registerDfsdm_flt3isrType uint32

const (
	RegisterDfsdm_flt3isrFieldJeocfShift = 0
	RegisterDfsdm_flt3isrFieldJeocfMask  = 0x1
)

// GetJeocf End of injected conversion flag
func (r *registerDfsdm_flt3isrType) GetJeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3isrFieldJeocfMask) != 0
}

// SetJeocf End of injected conversion flag
func (r *registerDfsdm_flt3isrType) SetJeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3isrFieldJeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3isrFieldJeocfMask)
	}
}

const (
	RegisterDfsdm_flt3isrFieldReocfShift = 1
	RegisterDfsdm_flt3isrFieldReocfMask  = 0x2
)

// GetReocf End of regular conversion flag
func (r *registerDfsdm_flt3isrType) GetReocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3isrFieldReocfMask) != 0
}

// SetReocf End of regular conversion flag
func (r *registerDfsdm_flt3isrType) SetReocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3isrFieldReocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3isrFieldReocfMask)
	}
}

const (
	RegisterDfsdm_flt3isrFieldJovrfShift = 2
	RegisterDfsdm_flt3isrFieldJovrfMask  = 0x4
)

// GetJovrf Injected conversion overrun flag
func (r *registerDfsdm_flt3isrType) GetJovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3isrFieldJovrfMask) != 0
}

// SetJovrf Injected conversion overrun flag
func (r *registerDfsdm_flt3isrType) SetJovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3isrFieldJovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3isrFieldJovrfMask)
	}
}

const (
	RegisterDfsdm_flt3isrFieldRovrfShift = 3
	RegisterDfsdm_flt3isrFieldRovrfMask  = 0x8
)

// GetRovrf Regular conversion overrun flag
func (r *registerDfsdm_flt3isrType) GetRovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3isrFieldRovrfMask) != 0
}

// SetRovrf Regular conversion overrun flag
func (r *registerDfsdm_flt3isrType) SetRovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3isrFieldRovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3isrFieldRovrfMask)
	}
}

const (
	RegisterDfsdm_flt3isrFieldAwdfShift = 4
	RegisterDfsdm_flt3isrFieldAwdfMask  = 0x10
)

// GetAwdf Analog watchdog
func (r *registerDfsdm_flt3isrType) GetAwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3isrFieldAwdfMask) != 0
}

// SetAwdf Analog watchdog
func (r *registerDfsdm_flt3isrType) SetAwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3isrFieldAwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3isrFieldAwdfMask)
	}
}

const (
	RegisterDfsdm_flt3isrFieldJcipShift = 13
	RegisterDfsdm_flt3isrFieldJcipMask  = 0x2000
)

// GetJcip Injected conversion in progress status
func (r *registerDfsdm_flt3isrType) GetJcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3isrFieldJcipMask) != 0
}

// SetJcip Injected conversion in progress status
func (r *registerDfsdm_flt3isrType) SetJcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3isrFieldJcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3isrFieldJcipMask)
	}
}

const (
	RegisterDfsdm_flt3isrFieldRcipShift = 14
	RegisterDfsdm_flt3isrFieldRcipMask  = 0x4000
)

// GetRcip Regular conversion in progress status
func (r *registerDfsdm_flt3isrType) GetRcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3isrFieldRcipMask) != 0
}

// SetRcip Regular conversion in progress status
func (r *registerDfsdm_flt3isrType) SetRcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3isrFieldRcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3isrFieldRcipMask)
	}
}

const (
	RegisterDfsdm_flt3isrFieldCkabfShift = 16
	RegisterDfsdm_flt3isrFieldCkabfMask  = 0xff0000
)

// GetCkabf Clock absence flag
func (r *registerDfsdm_flt3isrType) GetCkabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3isrFieldCkabfMask) >> RegisterDfsdm_flt3isrFieldCkabfShift)
}

// SetCkabf Clock absence flag
func (r *registerDfsdm_flt3isrType) SetCkabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3isrFieldCkabfMask)|(uint32(value)<<RegisterDfsdm_flt3isrFieldCkabfShift))
}

const (
	RegisterDfsdm_flt3isrFieldScdfShift = 24
	RegisterDfsdm_flt3isrFieldScdfMask  = 0xff000000
)

// GetScdf short-circuit detector flag
func (r *registerDfsdm_flt3isrType) GetScdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3isrFieldScdfMask) >> RegisterDfsdm_flt3isrFieldScdfShift)
}

// SetScdf short-circuit detector flag
func (r *registerDfsdm_flt3isrType) SetScdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3isrFieldScdfMask)|(uint32(value)<<RegisterDfsdm_flt3isrFieldScdfShift))
}

// registerDfsdm_flt3icrType interrupt flag clear register
type registerDfsdm_flt3icrType uint32

const (
	RegisterDfsdm_flt3icrFieldClrjovrfShift = 2
	RegisterDfsdm_flt3icrFieldClrjovrfMask  = 0x4
)

// GetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm_flt3icrType) GetClrjovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3icrFieldClrjovrfMask) != 0
}

// SetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm_flt3icrType) SetClrjovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3icrFieldClrjovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3icrFieldClrjovrfMask)
	}
}

const (
	RegisterDfsdm_flt3icrFieldClrrovrfShift = 3
	RegisterDfsdm_flt3icrFieldClrrovrfMask  = 0x8
)

// GetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm_flt3icrType) GetClrrovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3icrFieldClrrovrfMask) != 0
}

// SetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm_flt3icrType) SetClrrovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3icrFieldClrrovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3icrFieldClrrovrfMask)
	}
}

const (
	RegisterDfsdm_flt3icrFieldClrckabfShift = 16
	RegisterDfsdm_flt3icrFieldClrckabfMask  = 0xff0000
)

// GetClrckabf Clear the clock absence flag
func (r *registerDfsdm_flt3icrType) GetClrckabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3icrFieldClrckabfMask) >> RegisterDfsdm_flt3icrFieldClrckabfShift)
}

// SetClrckabf Clear the clock absence flag
func (r *registerDfsdm_flt3icrType) SetClrckabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3icrFieldClrckabfMask)|(uint32(value)<<RegisterDfsdm_flt3icrFieldClrckabfShift))
}

const (
	RegisterDfsdm_flt3icrFieldClrscdfShift = 24
	RegisterDfsdm_flt3icrFieldClrscdfMask  = 0xff000000
)

// GetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm_flt3icrType) GetClrscdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3icrFieldClrscdfMask) >> RegisterDfsdm_flt3icrFieldClrscdfShift)
}

// SetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm_flt3icrType) SetClrscdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3icrFieldClrscdfMask)|(uint32(value)<<RegisterDfsdm_flt3icrFieldClrscdfShift))
}

// registerDfsdm_flt3jchgrType injected channel group selection register
type registerDfsdm_flt3jchgrType uint32

const (
	RegisterDfsdm_flt3jchgrFieldJchgShift = 0
	RegisterDfsdm_flt3jchgrFieldJchgMask  = 0xff
)

// GetJchg Injected channel group selection
func (r *registerDfsdm_flt3jchgrType) GetJchg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3jchgrFieldJchgMask) >> RegisterDfsdm_flt3jchgrFieldJchgShift)
}

// SetJchg Injected channel group selection
func (r *registerDfsdm_flt3jchgrType) SetJchg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3jchgrFieldJchgMask)|(uint32(value)<<RegisterDfsdm_flt3jchgrFieldJchgShift))
}

// registerDfsdm_flt3fcrType filter control register
type registerDfsdm_flt3fcrType uint32

const (
	RegisterDfsdm_flt3fcrFieldIosrShift = 0
	RegisterDfsdm_flt3fcrFieldIosrMask  = 0xff
)

// GetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm_flt3fcrType) GetIosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3fcrFieldIosrMask) >> RegisterDfsdm_flt3fcrFieldIosrShift)
}

// SetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm_flt3fcrType) SetIosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3fcrFieldIosrMask)|(uint32(value)<<RegisterDfsdm_flt3fcrFieldIosrShift))
}

const (
	RegisterDfsdm_flt3fcrFieldFosrShift = 16
	RegisterDfsdm_flt3fcrFieldFosrMask  = 0x3ff0000
)

// GetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm_flt3fcrType) GetFosr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3fcrFieldFosrMask) >> RegisterDfsdm_flt3fcrFieldFosrShift)
}

// SetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm_flt3fcrType) SetFosr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3fcrFieldFosrMask)|(uint32(value)<<RegisterDfsdm_flt3fcrFieldFosrShift))
}

const (
	RegisterDfsdm_flt3fcrFieldFordShift = 29
	RegisterDfsdm_flt3fcrFieldFordMask  = 0xe0000000
)

// GetFord Sinc filter order
func (r *registerDfsdm_flt3fcrType) GetFord() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3fcrFieldFordMask) >> RegisterDfsdm_flt3fcrFieldFordShift)
}

// SetFord Sinc filter order
func (r *registerDfsdm_flt3fcrType) SetFord(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3fcrFieldFordMask)|(uint32(value)<<RegisterDfsdm_flt3fcrFieldFordShift))
}

// registerDfsdm_flt3jdatarType data register for injected group
type registerDfsdm_flt3jdatarType uint32

const (
	RegisterDfsdm_flt3jdatarFieldJdatachShift = 0
	RegisterDfsdm_flt3jdatarFieldJdatachMask  = 0x7
)

// GetJdatach Injected channel most recently converted
func (r *registerDfsdm_flt3jdatarType) GetJdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3jdatarFieldJdatachMask) >> RegisterDfsdm_flt3jdatarFieldJdatachShift)
}

// SetJdatach Injected channel most recently converted
func (r *registerDfsdm_flt3jdatarType) SetJdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3jdatarFieldJdatachMask)|(uint32(value)<<RegisterDfsdm_flt3jdatarFieldJdatachShift))
}

const (
	RegisterDfsdm_flt3jdatarFieldJdataShift = 8
	RegisterDfsdm_flt3jdatarFieldJdataMask  = 0xffffff00
)

// GetJdata Injected group conversion data
func (r *registerDfsdm_flt3jdatarType) GetJdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3jdatarFieldJdataMask) >> RegisterDfsdm_flt3jdatarFieldJdataShift)
}

// SetJdata Injected group conversion data
func (r *registerDfsdm_flt3jdatarType) SetJdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3jdatarFieldJdataMask)|(uint32(value)<<RegisterDfsdm_flt3jdatarFieldJdataShift))
}

// registerDfsdm_flt3rdatarType data register for the regular channel
type registerDfsdm_flt3rdatarType uint32

const (
	RegisterDfsdm_flt3rdatarFieldRdatachShift = 0
	RegisterDfsdm_flt3rdatarFieldRdatachMask  = 0x7
)

// GetRdatach Regular channel most recently converted
func (r *registerDfsdm_flt3rdatarType) GetRdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3rdatarFieldRdatachMask) >> RegisterDfsdm_flt3rdatarFieldRdatachShift)
}

// SetRdatach Regular channel most recently converted
func (r *registerDfsdm_flt3rdatarType) SetRdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3rdatarFieldRdatachMask)|(uint32(value)<<RegisterDfsdm_flt3rdatarFieldRdatachShift))
}

const (
	RegisterDfsdm_flt3rdatarFieldRpendShift = 4
	RegisterDfsdm_flt3rdatarFieldRpendMask  = 0x10
)

// GetRpend Regular channel pending data
func (r *registerDfsdm_flt3rdatarType) GetRpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3rdatarFieldRpendMask) != 0
}

// SetRpend Regular channel pending data
func (r *registerDfsdm_flt3rdatarType) SetRpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm_flt3rdatarFieldRpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3rdatarFieldRpendMask)
	}
}

const (
	RegisterDfsdm_flt3rdatarFieldRdataShift = 8
	RegisterDfsdm_flt3rdatarFieldRdataMask  = 0xffffff00
)

// GetRdata Regular channel conversion data
func (r *registerDfsdm_flt3rdatarType) GetRdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3rdatarFieldRdataMask) >> RegisterDfsdm_flt3rdatarFieldRdataShift)
}

// SetRdata Regular channel conversion data
func (r *registerDfsdm_flt3rdatarType) SetRdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3rdatarFieldRdataMask)|(uint32(value)<<RegisterDfsdm_flt3rdatarFieldRdataShift))
}

// registerDfsdm_flt3awhtrType analog watchdog high threshold register
type registerDfsdm_flt3awhtrType uint32

const (
	RegisterDfsdm_flt3awhtrFieldBkawhShift = 0
	RegisterDfsdm_flt3awhtrFieldBkawhMask  = 0xf
)

// GetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm_flt3awhtrType) GetBkawh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3awhtrFieldBkawhMask) >> RegisterDfsdm_flt3awhtrFieldBkawhShift)
}

// SetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdm_flt3awhtrType) SetBkawh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3awhtrFieldBkawhMask)|(uint32(value)<<RegisterDfsdm_flt3awhtrFieldBkawhShift))
}

const (
	RegisterDfsdm_flt3awhtrFieldAwhtShift = 8
	RegisterDfsdm_flt3awhtrFieldAwhtMask  = 0xffffff00
)

// GetAwht Analog watchdog high threshold
func (r *registerDfsdm_flt3awhtrType) GetAwht() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3awhtrFieldAwhtMask) >> RegisterDfsdm_flt3awhtrFieldAwhtShift)
}

// SetAwht Analog watchdog high threshold
func (r *registerDfsdm_flt3awhtrType) SetAwht(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3awhtrFieldAwhtMask)|(uint32(value)<<RegisterDfsdm_flt3awhtrFieldAwhtShift))
}

// registerDfsdm_flt3awltrType analog watchdog low threshold register
type registerDfsdm_flt3awltrType uint32

const (
	RegisterDfsdm_flt3awltrFieldBkawlShift = 0
	RegisterDfsdm_flt3awltrFieldBkawlMask  = 0xf
)

// GetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm_flt3awltrType) GetBkawl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3awltrFieldBkawlMask) >> RegisterDfsdm_flt3awltrFieldBkawlShift)
}

// SetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdm_flt3awltrType) SetBkawl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3awltrFieldBkawlMask)|(uint32(value)<<RegisterDfsdm_flt3awltrFieldBkawlShift))
}

const (
	RegisterDfsdm_flt3awltrFieldAwltShift = 8
	RegisterDfsdm_flt3awltrFieldAwltMask  = 0xffffff00
)

// GetAwlt Analog watchdog low threshold
func (r *registerDfsdm_flt3awltrType) GetAwlt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3awltrFieldAwltMask) >> RegisterDfsdm_flt3awltrFieldAwltShift)
}

// SetAwlt Analog watchdog low threshold
func (r *registerDfsdm_flt3awltrType) SetAwlt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3awltrFieldAwltMask)|(uint32(value)<<RegisterDfsdm_flt3awltrFieldAwltShift))
}

// registerDfsdm_flt3awsrType analog watchdog status register
type registerDfsdm_flt3awsrType uint32

const (
	RegisterDfsdm_flt3awsrFieldAwltfShift = 0
	RegisterDfsdm_flt3awsrFieldAwltfMask  = 0xff
)

// GetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm_flt3awsrType) GetAwltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3awsrFieldAwltfMask) >> RegisterDfsdm_flt3awsrFieldAwltfShift)
}

// SetAwltf Analog watchdog low threshold flag
func (r *registerDfsdm_flt3awsrType) SetAwltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3awsrFieldAwltfMask)|(uint32(value)<<RegisterDfsdm_flt3awsrFieldAwltfShift))
}

const (
	RegisterDfsdm_flt3awsrFieldAwhtfShift = 8
	RegisterDfsdm_flt3awsrFieldAwhtfMask  = 0xff00
)

// GetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm_flt3awsrType) GetAwhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3awsrFieldAwhtfMask) >> RegisterDfsdm_flt3awsrFieldAwhtfShift)
}

// SetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdm_flt3awsrType) SetAwhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3awsrFieldAwhtfMask)|(uint32(value)<<RegisterDfsdm_flt3awsrFieldAwhtfShift))
}

// registerDfsdm_flt3awcfrType analog watchdog clear flag register
type registerDfsdm_flt3awcfrType uint32

const (
	RegisterDfsdm_flt3awcfrFieldClrawltfShift = 0
	RegisterDfsdm_flt3awcfrFieldClrawltfMask  = 0xff
)

// GetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm_flt3awcfrType) GetClrawltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3awcfrFieldClrawltfMask) >> RegisterDfsdm_flt3awcfrFieldClrawltfShift)
}

// SetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdm_flt3awcfrType) SetClrawltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3awcfrFieldClrawltfMask)|(uint32(value)<<RegisterDfsdm_flt3awcfrFieldClrawltfShift))
}

const (
	RegisterDfsdm_flt3awcfrFieldClrawhtfShift = 8
	RegisterDfsdm_flt3awcfrFieldClrawhtfMask  = 0xff00
)

// GetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm_flt3awcfrType) GetClrawhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3awcfrFieldClrawhtfMask) >> RegisterDfsdm_flt3awcfrFieldClrawhtfShift)
}

// SetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdm_flt3awcfrType) SetClrawhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3awcfrFieldClrawhtfMask)|(uint32(value)<<RegisterDfsdm_flt3awcfrFieldClrawhtfShift))
}

// registerDfsdm_flt3exmaxType Extremes detector maximum register
type registerDfsdm_flt3exmaxType uint32

const (
	RegisterDfsdm_flt3exmaxFieldExmaxchShift = 0
	RegisterDfsdm_flt3exmaxFieldExmaxchMask  = 0x7
)

// GetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm_flt3exmaxType) GetExmaxch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3exmaxFieldExmaxchMask) >> RegisterDfsdm_flt3exmaxFieldExmaxchShift)
}

// SetExmaxch Extremes detector maximum data channel
func (r *registerDfsdm_flt3exmaxType) SetExmaxch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3exmaxFieldExmaxchMask)|(uint32(value)<<RegisterDfsdm_flt3exmaxFieldExmaxchShift))
}

const (
	RegisterDfsdm_flt3exmaxFieldExmaxShift = 8
	RegisterDfsdm_flt3exmaxFieldExmaxMask  = 0xffffff00
)

// GetExmax Extremes detector maximum value
func (r *registerDfsdm_flt3exmaxType) GetExmax() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3exmaxFieldExmaxMask) >> RegisterDfsdm_flt3exmaxFieldExmaxShift)
}

// SetExmax Extremes detector maximum value
func (r *registerDfsdm_flt3exmaxType) SetExmax(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3exmaxFieldExmaxMask)|(uint32(value)<<RegisterDfsdm_flt3exmaxFieldExmaxShift))
}

// registerDfsdm_flt3exminType Extremes detector minimum register
type registerDfsdm_flt3exminType uint32

const (
	RegisterDfsdm_flt3exminFieldExminchShift = 0
	RegisterDfsdm_flt3exminFieldExminchMask  = 0x7
)

// GetExminch Extremes detector minimum data channel
func (r *registerDfsdm_flt3exminType) GetExminch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3exminFieldExminchMask) >> RegisterDfsdm_flt3exminFieldExminchShift)
}

// SetExminch Extremes detector minimum data channel
func (r *registerDfsdm_flt3exminType) SetExminch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3exminFieldExminchMask)|(uint32(value)<<RegisterDfsdm_flt3exminFieldExminchShift))
}

const (
	RegisterDfsdm_flt3exminFieldExminShift = 8
	RegisterDfsdm_flt3exminFieldExminMask  = 0xffffff00
)

// GetExmin EXMIN
func (r *registerDfsdm_flt3exminType) GetExmin() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3exminFieldExminMask) >> RegisterDfsdm_flt3exminFieldExminShift)
}

// SetExmin EXMIN
func (r *registerDfsdm_flt3exminType) SetExmin(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3exminFieldExminMask)|(uint32(value)<<RegisterDfsdm_flt3exminFieldExminShift))
}

// registerDfsdm_flt3cnvtimrType conversion timer register
type registerDfsdm_flt3cnvtimrType uint32

const (
	RegisterDfsdm_flt3cnvtimrFieldCnvcntShift = 4
	RegisterDfsdm_flt3cnvtimrFieldCnvcntMask  = 0xfffffff0
)

// GetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdm_flt3cnvtimrType) GetCnvcnt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm_flt3cnvtimrFieldCnvcntMask) >> RegisterDfsdm_flt3cnvtimrFieldCnvcntShift)
}

// SetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdm_flt3cnvtimrType) SetCnvcnt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm_flt3cnvtimrFieldCnvcntMask)|(uint32(value)<<RegisterDfsdm_flt3cnvtimrFieldCnvcntShift))
}
