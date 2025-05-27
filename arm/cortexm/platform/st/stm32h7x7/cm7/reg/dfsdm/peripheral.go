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
	Ch0cfgr1         registerCh0cfgr1Type
	Ch0cfgr2         registerCh0cfgr2Type
	Ch0awscdr        registerCh0awscdrType
	Ch0wdatr         registerCh0wdatrType
	Ch0datinr        registerCh0datinrType
	Ch0dlyr          registerCh0dlyrType
	_                [8]byte
	Ch1cfgr1         registerCh1cfgr1Type
	Ch1cfgr2         registerCh1cfgr2Type
	Ch1awscdr        registerCh1awscdrType
	Ch1wdatr         registerCh1wdatrType
	Ch1datinr        registerCh1datinrType
	Ch1dlyr          registerCh1dlyrType
	_                [8]byte
	Ch2cfgr1         registerCh2cfgr1Type
	Ch2cfgr2         registerCh2cfgr2Type
	Ch2awscdr        registerCh2awscdrType
	Ch2wdatr         registerCh2wdatrType
	Ch2datinr        registerCh2datinrType
	Ch2dlyr          registerCh2dlyrType
	_                [8]byte
	Ch3cfgr1         registerCh3cfgr1Type
	Ch3cfgr2         registerCh3cfgr2Type
	Ch3awscdr        registerCh3awscdrType
	Ch3wdatr         registerCh3wdatrType
	Ch3datinr        registerCh3datinrType
	Ch3dlyr          registerCh3dlyrType
	_                [8]byte
	Ch4cfgr1         registerCh4cfgr1Type
	Ch4cfgr2         registerCh4cfgr2Type
	Ch4awscdr        registerCh4awscdrType
	Ch4wdatr         registerCh4wdatrType
	Ch4datinr        registerCh4datinrType
	Ch4dlyr          registerCh4dlyrType
	_                [8]byte
	Ch5cfgr1         registerCh5cfgr1Type
	Ch5cfgr2         registerCh5cfgr2Type
	Ch5awscdr        registerCh5awscdrType
	Ch5wdatr         registerCh5wdatrType
	Ch5datinr        registerCh5datinrType
	Ch5dlyr          registerCh5dlyrType
	_                [8]byte
	Ch6cfgr1         registerCh6cfgr1Type
	Ch6cfgr2         registerCh6cfgr2Type
	Ch6awscdr        registerCh6awscdrType
	Ch6wdatr         registerCh6wdatrType
	Ch6datinr        registerCh6datinrType
	Ch6dlyr          registerCh6dlyrType
	_                [8]byte
	Ch7cfgr1         registerCh7cfgr1Type
	Ch7cfgr2         registerCh7cfgr2Type
	Ch7awscdr        registerCh7awscdrType
	Ch7wdatr         registerCh7wdatrType
	Ch7datinr        registerCh7datinrType
	Ch7dlyr          registerCh7dlyrType
	_                [8]byte
	Dfsdmflt0cr1     registerDfsdmflt0cr1Type
	Dfsdmflt0cr2     registerDfsdmflt0cr2Type
	Dfsdmflt0isr     registerDfsdmflt0isrType
	Dfsdmflt0icr     registerDfsdmflt0icrType
	Dfsdmflt0jchgr   registerDfsdmflt0jchgrType
	Dfsdmflt0fcr     registerDfsdmflt0fcrType
	Dfsdmflt0jdatar  registerDfsdmflt0jdatarType
	Dfsdmflt0rdatar  registerDfsdmflt0rdatarType
	Dfsdmflt0awhtr   registerDfsdmflt0awhtrType
	Dfsdmflt0awltr   registerDfsdmflt0awltrType
	Dfsdmflt0awsr    registerDfsdmflt0awsrType
	Dfsdmflt0awcfr   registerDfsdmflt0awcfrType
	Dfsdmflt0exmax   registerDfsdmflt0exmaxType
	Dfsdmflt0exmin   registerDfsdmflt0exminType
	Dfsdmflt0cnvtimr registerDfsdmflt0cnvtimrType
	_                [68]byte
	Dfsdmflt1cr1     registerDfsdmflt1cr1Type
	Dfsdmflt1cr2     registerDfsdmflt1cr2Type
	Dfsdmflt1isr     registerDfsdmflt1isrType
	Dfsdm1icr        registerDfsdm1icrType
	Dfsdmflt1jchgr   registerDfsdmflt1jchgrType
	Dfsdm1fcr        registerDfsdm1fcrType
	Dfsdmflt1jdatar  registerDfsdmflt1jdatarType
	Dfsdmflt1rdatar  registerDfsdmflt1rdatarType
	Dfsdmflt1awhtr   registerDfsdmflt1awhtrType
	Dfsdmflt1awltr   registerDfsdmflt1awltrType
	Dfsdmflt1awsr    registerDfsdmflt1awsrType
	Dfsdmflt1awcfr   registerDfsdmflt1awcfrType
	Dfsdmflt1exmax   registerDfsdmflt1exmaxType
	Dfsdmflt1exmin   registerDfsdmflt1exminType
	Dfsdmflt1cnvtimr registerDfsdmflt1cnvtimrType
	_                [68]byte
	Dfsdmflt2cr1     registerDfsdmflt2cr1Type
	Dfsdmflt2cr2     registerDfsdmflt2cr2Type
	Dfsdmflt2isr     registerDfsdmflt2isrType
	Dfsdmflt2icr     registerDfsdmflt2icrType
	Dfsdmflt2jchgr   registerDfsdmflt2jchgrType
	Dfsdmflt2fcr     registerDfsdmflt2fcrType
	Dfsdmflt2jdatar  registerDfsdmflt2jdatarType
	Dfsdmflt2rdatar  registerDfsdmflt2rdatarType
	Dfsdmflt2awhtr   registerDfsdmflt2awhtrType
	Dfsdmflt2awltr   registerDfsdmflt2awltrType
	Dfsdmflt2awsr    registerDfsdmflt2awsrType
	Dfsdmflt2awcfr   registerDfsdmflt2awcfrType
	Dfsdmflt2exmax   registerDfsdmflt2exmaxType
	Dfsdmflt2exmin   registerDfsdmflt2exminType
	Dfsdmflt2cnvtimr registerDfsdmflt2cnvtimrType
	_                [68]byte
	Dfsdmflt3cr1     registerDfsdmflt3cr1Type
	Dfsdmflt3cr2     registerDfsdmflt3cr2Type
	Dfsdmflt3isr     registerDfsdmflt3isrType
	Dfsdmflt3icr     registerDfsdmflt3icrType
	Dfsdmflt3jchgr   registerDfsdmflt3jchgrType
	Dfsdmflt3fcr     registerDfsdmflt3fcrType
	Dfsdmflt3jdatar  registerDfsdmflt3jdatarType
	Dfsdmflt3rdatar  registerDfsdmflt3rdatarType
	Dfsdmflt3awhtr   registerDfsdmflt3awhtrType
	Dfsdmflt3awltr   registerDfsdmflt3awltrType
	Dfsdmflt3awsr    registerDfsdmflt3awsrType
	Dfsdmflt3awcfr   registerDfsdmflt3awcfrType
	Dfsdmflt3exmax   registerDfsdmflt3exmaxType
	Dfsdmflt3exmin   registerDfsdmflt3exminType
	Dfsdmflt3cnvtimr registerDfsdmflt3cnvtimrType
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

// registerDfsdmflt0cr1Type control register 1
type registerDfsdmflt0cr1Type uint32

const (
	RegisterDfsdmflt0cr1FieldDfenShift = 0
	RegisterDfsdmflt0cr1FieldDfenMask  = 0x1
)

// GetDfen DFSDM enable
func (r *registerDfsdmflt0cr1Type) GetDfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldDfenMask) != 0
}

// SetDfen DFSDM enable
func (r *registerDfsdmflt0cr1Type) SetDfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr1FieldDfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldDfenMask)
	}
}

const (
	RegisterDfsdmflt0cr1FieldJswstartShift = 1
	RegisterDfsdmflt0cr1FieldJswstartMask  = 0x2
)

// GetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdmflt0cr1Type) GetJswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldJswstartMask) != 0
}

// SetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdmflt0cr1Type) SetJswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr1FieldJswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldJswstartMask)
	}
}

const (
	RegisterDfsdmflt0cr1FieldJsyncShift = 3
	RegisterDfsdmflt0cr1FieldJsyncMask  = 0x8
)

// GetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdmflt0cr1Type) GetJsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldJsyncMask) != 0
}

// SetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdmflt0cr1Type) SetJsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr1FieldJsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldJsyncMask)
	}
}

const (
	RegisterDfsdmflt0cr1FieldJscanShift = 4
	RegisterDfsdmflt0cr1FieldJscanMask  = 0x10
)

// GetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdmflt0cr1Type) GetJscan() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldJscanMask) != 0
}

// SetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdmflt0cr1Type) SetJscan(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr1FieldJscanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldJscanMask)
	}
}

const (
	RegisterDfsdmflt0cr1FieldJdmaenShift = 5
	RegisterDfsdmflt0cr1FieldJdmaenMask  = 0x20
)

// GetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdmflt0cr1Type) GetJdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldJdmaenMask) != 0
}

// SetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdmflt0cr1Type) SetJdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr1FieldJdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldJdmaenMask)
	}
}

const (
	RegisterDfsdmflt0cr1FieldJextselShift = 8
	RegisterDfsdmflt0cr1FieldJextselMask  = 0x700
)

// GetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdmflt0cr1Type) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldJextselMask) >> RegisterDfsdmflt0cr1FieldJextselShift)
}

// SetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdmflt0cr1Type) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldJextselMask)|(uint32(value)<<RegisterDfsdmflt0cr1FieldJextselShift))
}

const (
	RegisterDfsdmflt0cr1FieldJextenShift = 13
	RegisterDfsdmflt0cr1FieldJextenMask  = 0x6000
)

// GetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdmflt0cr1Type) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldJextenMask) >> RegisterDfsdmflt0cr1FieldJextenShift)
}

// SetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdmflt0cr1Type) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldJextenMask)|(uint32(value)<<RegisterDfsdmflt0cr1FieldJextenShift))
}

const (
	RegisterDfsdmflt0cr1FieldRswstartShift = 17
	RegisterDfsdmflt0cr1FieldRswstartMask  = 0x20000
)

// GetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdmflt0cr1Type) GetRswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldRswstartMask) != 0
}

// SetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdmflt0cr1Type) SetRswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr1FieldRswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldRswstartMask)
	}
}

const (
	RegisterDfsdmflt0cr1FieldRcontShift = 18
	RegisterDfsdmflt0cr1FieldRcontMask  = 0x40000
)

// GetRcont Continuous mode selection for regular conversions
func (r *registerDfsdmflt0cr1Type) GetRcont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldRcontMask) != 0
}

// SetRcont Continuous mode selection for regular conversions
func (r *registerDfsdmflt0cr1Type) SetRcont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr1FieldRcontMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldRcontMask)
	}
}

const (
	RegisterDfsdmflt0cr1FieldRsyncShift = 19
	RegisterDfsdmflt0cr1FieldRsyncMask  = 0x80000
)

// GetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdmflt0cr1Type) GetRsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldRsyncMask) != 0
}

// SetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdmflt0cr1Type) SetRsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr1FieldRsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldRsyncMask)
	}
}

const (
	RegisterDfsdmflt0cr1FieldRdmaenShift = 21
	RegisterDfsdmflt0cr1FieldRdmaenMask  = 0x200000
)

// GetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdmflt0cr1Type) GetRdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldRdmaenMask) != 0
}

// SetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdmflt0cr1Type) SetRdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr1FieldRdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldRdmaenMask)
	}
}

const (
	RegisterDfsdmflt0cr1FieldRchShift = 24
	RegisterDfsdmflt0cr1FieldRchMask  = 0x7000000
)

// GetRch Regular channel selection
func (r *registerDfsdmflt0cr1Type) GetRch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldRchMask) >> RegisterDfsdmflt0cr1FieldRchShift)
}

// SetRch Regular channel selection
func (r *registerDfsdmflt0cr1Type) SetRch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldRchMask)|(uint32(value)<<RegisterDfsdmflt0cr1FieldRchShift))
}

const (
	RegisterDfsdmflt0cr1FieldFastShift = 29
	RegisterDfsdmflt0cr1FieldFastMask  = 0x20000000
)

// GetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdmflt0cr1Type) GetFast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldFastMask) != 0
}

// SetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdmflt0cr1Type) SetFast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr1FieldFastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldFastMask)
	}
}

const (
	RegisterDfsdmflt0cr1FieldAwfselShift = 30
	RegisterDfsdmflt0cr1FieldAwfselMask  = 0x40000000
)

// GetAwfsel Analog watchdog fast mode select
func (r *registerDfsdmflt0cr1Type) GetAwfsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr1FieldAwfselMask) != 0
}

// SetAwfsel Analog watchdog fast mode select
func (r *registerDfsdmflt0cr1Type) SetAwfsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr1FieldAwfselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr1FieldAwfselMask)
	}
}

// registerDfsdmflt0cr2Type control register 2
type registerDfsdmflt0cr2Type uint32

const (
	RegisterDfsdmflt0cr2FieldJeocieShift = 0
	RegisterDfsdmflt0cr2FieldJeocieMask  = 0x1
)

// GetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdmflt0cr2Type) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr2FieldJeocieMask) != 0
}

// SetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdmflt0cr2Type) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr2FieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr2FieldJeocieMask)
	}
}

const (
	RegisterDfsdmflt0cr2FieldReocieShift = 1
	RegisterDfsdmflt0cr2FieldReocieMask  = 0x2
)

// GetReocie Regular end of conversion interrupt enable
func (r *registerDfsdmflt0cr2Type) GetReocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr2FieldReocieMask) != 0
}

// SetReocie Regular end of conversion interrupt enable
func (r *registerDfsdmflt0cr2Type) SetReocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr2FieldReocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr2FieldReocieMask)
	}
}

const (
	RegisterDfsdmflt0cr2FieldJovrieShift = 2
	RegisterDfsdmflt0cr2FieldJovrieMask  = 0x4
)

// GetJovrie Injected data overrun interrupt enable
func (r *registerDfsdmflt0cr2Type) GetJovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr2FieldJovrieMask) != 0
}

// SetJovrie Injected data overrun interrupt enable
func (r *registerDfsdmflt0cr2Type) SetJovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr2FieldJovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr2FieldJovrieMask)
	}
}

const (
	RegisterDfsdmflt0cr2FieldRovrieShift = 3
	RegisterDfsdmflt0cr2FieldRovrieMask  = 0x8
)

// GetRovrie Regular data overrun interrupt enable
func (r *registerDfsdmflt0cr2Type) GetRovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr2FieldRovrieMask) != 0
}

// SetRovrie Regular data overrun interrupt enable
func (r *registerDfsdmflt0cr2Type) SetRovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr2FieldRovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr2FieldRovrieMask)
	}
}

const (
	RegisterDfsdmflt0cr2FieldAwdieShift = 4
	RegisterDfsdmflt0cr2FieldAwdieMask  = 0x10
)

// GetAwdie Analog watchdog interrupt enable
func (r *registerDfsdmflt0cr2Type) GetAwdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr2FieldAwdieMask) != 0
}

// SetAwdie Analog watchdog interrupt enable
func (r *registerDfsdmflt0cr2Type) SetAwdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr2FieldAwdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr2FieldAwdieMask)
	}
}

const (
	RegisterDfsdmflt0cr2FieldScdieShift = 5
	RegisterDfsdmflt0cr2FieldScdieMask  = 0x20
)

// GetScdie Short-circuit detector interrupt enable
func (r *registerDfsdmflt0cr2Type) GetScdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr2FieldScdieMask) != 0
}

// SetScdie Short-circuit detector interrupt enable
func (r *registerDfsdmflt0cr2Type) SetScdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr2FieldScdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr2FieldScdieMask)
	}
}

const (
	RegisterDfsdmflt0cr2FieldCkabieShift = 6
	RegisterDfsdmflt0cr2FieldCkabieMask  = 0x40
)

// GetCkabie Clock absence interrupt enable
func (r *registerDfsdmflt0cr2Type) GetCkabie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr2FieldCkabieMask) != 0
}

// SetCkabie Clock absence interrupt enable
func (r *registerDfsdmflt0cr2Type) SetCkabie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0cr2FieldCkabieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr2FieldCkabieMask)
	}
}

const (
	RegisterDfsdmflt0cr2FieldExchShift = 8
	RegisterDfsdmflt0cr2FieldExchMask  = 0xff00
)

// GetExch Extremes detector channel selection
func (r *registerDfsdmflt0cr2Type) GetExch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr2FieldExchMask) >> RegisterDfsdmflt0cr2FieldExchShift)
}

// SetExch Extremes detector channel selection
func (r *registerDfsdmflt0cr2Type) SetExch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr2FieldExchMask)|(uint32(value)<<RegisterDfsdmflt0cr2FieldExchShift))
}

const (
	RegisterDfsdmflt0cr2FieldAwdchShift = 16
	RegisterDfsdmflt0cr2FieldAwdchMask  = 0xff0000
)

// GetAwdch Analog watchdog channel selection
func (r *registerDfsdmflt0cr2Type) GetAwdch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cr2FieldAwdchMask) >> RegisterDfsdmflt0cr2FieldAwdchShift)
}

// SetAwdch Analog watchdog channel selection
func (r *registerDfsdmflt0cr2Type) SetAwdch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cr2FieldAwdchMask)|(uint32(value)<<RegisterDfsdmflt0cr2FieldAwdchShift))
}

// registerDfsdmflt0isrType interrupt and status register
type registerDfsdmflt0isrType uint32

const (
	RegisterDfsdmflt0isrFieldJeocfShift = 0
	RegisterDfsdmflt0isrFieldJeocfMask  = 0x1
)

// GetJeocf End of injected conversion flag
func (r *registerDfsdmflt0isrType) GetJeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0isrFieldJeocfMask) != 0
}

// SetJeocf End of injected conversion flag
func (r *registerDfsdmflt0isrType) SetJeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0isrFieldJeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0isrFieldJeocfMask)
	}
}

const (
	RegisterDfsdmflt0isrFieldReocfShift = 1
	RegisterDfsdmflt0isrFieldReocfMask  = 0x2
)

// GetReocf End of regular conversion flag
func (r *registerDfsdmflt0isrType) GetReocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0isrFieldReocfMask) != 0
}

// SetReocf End of regular conversion flag
func (r *registerDfsdmflt0isrType) SetReocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0isrFieldReocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0isrFieldReocfMask)
	}
}

const (
	RegisterDfsdmflt0isrFieldJovrfShift = 2
	RegisterDfsdmflt0isrFieldJovrfMask  = 0x4
)

// GetJovrf Injected conversion overrun flag
func (r *registerDfsdmflt0isrType) GetJovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0isrFieldJovrfMask) != 0
}

// SetJovrf Injected conversion overrun flag
func (r *registerDfsdmflt0isrType) SetJovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0isrFieldJovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0isrFieldJovrfMask)
	}
}

const (
	RegisterDfsdmflt0isrFieldRovrfShift = 3
	RegisterDfsdmflt0isrFieldRovrfMask  = 0x8
)

// GetRovrf Regular conversion overrun flag
func (r *registerDfsdmflt0isrType) GetRovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0isrFieldRovrfMask) != 0
}

// SetRovrf Regular conversion overrun flag
func (r *registerDfsdmflt0isrType) SetRovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0isrFieldRovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0isrFieldRovrfMask)
	}
}

const (
	RegisterDfsdmflt0isrFieldAwdfShift = 4
	RegisterDfsdmflt0isrFieldAwdfMask  = 0x10
)

// GetAwdf Analog watchdog
func (r *registerDfsdmflt0isrType) GetAwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0isrFieldAwdfMask) != 0
}

// SetAwdf Analog watchdog
func (r *registerDfsdmflt0isrType) SetAwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0isrFieldAwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0isrFieldAwdfMask)
	}
}

const (
	RegisterDfsdmflt0isrFieldJcipShift = 13
	RegisterDfsdmflt0isrFieldJcipMask  = 0x2000
)

// GetJcip Injected conversion in progress status
func (r *registerDfsdmflt0isrType) GetJcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0isrFieldJcipMask) != 0
}

// SetJcip Injected conversion in progress status
func (r *registerDfsdmflt0isrType) SetJcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0isrFieldJcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0isrFieldJcipMask)
	}
}

const (
	RegisterDfsdmflt0isrFieldRcipShift = 14
	RegisterDfsdmflt0isrFieldRcipMask  = 0x4000
)

// GetRcip Regular conversion in progress status
func (r *registerDfsdmflt0isrType) GetRcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0isrFieldRcipMask) != 0
}

// SetRcip Regular conversion in progress status
func (r *registerDfsdmflt0isrType) SetRcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0isrFieldRcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0isrFieldRcipMask)
	}
}

const (
	RegisterDfsdmflt0isrFieldCkabfShift = 16
	RegisterDfsdmflt0isrFieldCkabfMask  = 0xff0000
)

// GetCkabf Clock absence flag
func (r *registerDfsdmflt0isrType) GetCkabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0isrFieldCkabfMask) >> RegisterDfsdmflt0isrFieldCkabfShift)
}

// SetCkabf Clock absence flag
func (r *registerDfsdmflt0isrType) SetCkabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0isrFieldCkabfMask)|(uint32(value)<<RegisterDfsdmflt0isrFieldCkabfShift))
}

const (
	RegisterDfsdmflt0isrFieldScdfShift = 24
	RegisterDfsdmflt0isrFieldScdfMask  = 0xff000000
)

// GetScdf short-circuit detector flag
func (r *registerDfsdmflt0isrType) GetScdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0isrFieldScdfMask) >> RegisterDfsdmflt0isrFieldScdfShift)
}

// SetScdf short-circuit detector flag
func (r *registerDfsdmflt0isrType) SetScdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0isrFieldScdfMask)|(uint32(value)<<RegisterDfsdmflt0isrFieldScdfShift))
}

// registerDfsdmflt0icrType interrupt flag clear register
type registerDfsdmflt0icrType uint32

const (
	RegisterDfsdmflt0icrFieldClrjovrfShift = 2
	RegisterDfsdmflt0icrFieldClrjovrfMask  = 0x4
)

// GetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdmflt0icrType) GetClrjovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0icrFieldClrjovrfMask) != 0
}

// SetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdmflt0icrType) SetClrjovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0icrFieldClrjovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0icrFieldClrjovrfMask)
	}
}

const (
	RegisterDfsdmflt0icrFieldClrrovrfShift = 3
	RegisterDfsdmflt0icrFieldClrrovrfMask  = 0x8
)

// GetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdmflt0icrType) GetClrrovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0icrFieldClrrovrfMask) != 0
}

// SetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdmflt0icrType) SetClrrovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0icrFieldClrrovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0icrFieldClrrovrfMask)
	}
}

const (
	RegisterDfsdmflt0icrFieldClrckabfShift = 16
	RegisterDfsdmflt0icrFieldClrckabfMask  = 0xff0000
)

// GetClrckabf Clear the clock absence flag
func (r *registerDfsdmflt0icrType) GetClrckabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0icrFieldClrckabfMask) >> RegisterDfsdmflt0icrFieldClrckabfShift)
}

// SetClrckabf Clear the clock absence flag
func (r *registerDfsdmflt0icrType) SetClrckabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0icrFieldClrckabfMask)|(uint32(value)<<RegisterDfsdmflt0icrFieldClrckabfShift))
}

const (
	RegisterDfsdmflt0icrFieldClrscdfShift = 24
	RegisterDfsdmflt0icrFieldClrscdfMask  = 0xff000000
)

// GetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdmflt0icrType) GetClrscdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0icrFieldClrscdfMask) >> RegisterDfsdmflt0icrFieldClrscdfShift)
}

// SetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdmflt0icrType) SetClrscdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0icrFieldClrscdfMask)|(uint32(value)<<RegisterDfsdmflt0icrFieldClrscdfShift))
}

// registerDfsdmflt0jchgrType injected channel group selection register
type registerDfsdmflt0jchgrType uint32

const (
	RegisterDfsdmflt0jchgrFieldJchgShift = 0
	RegisterDfsdmflt0jchgrFieldJchgMask  = 0xff
)

// GetJchg Injected channel group selection
func (r *registerDfsdmflt0jchgrType) GetJchg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0jchgrFieldJchgMask) >> RegisterDfsdmflt0jchgrFieldJchgShift)
}

// SetJchg Injected channel group selection
func (r *registerDfsdmflt0jchgrType) SetJchg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0jchgrFieldJchgMask)|(uint32(value)<<RegisterDfsdmflt0jchgrFieldJchgShift))
}

// registerDfsdmflt0fcrType filter control register
type registerDfsdmflt0fcrType uint32

const (
	RegisterDfsdmflt0fcrFieldIosrShift = 0
	RegisterDfsdmflt0fcrFieldIosrMask  = 0xff
)

// GetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdmflt0fcrType) GetIosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0fcrFieldIosrMask) >> RegisterDfsdmflt0fcrFieldIosrShift)
}

// SetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdmflt0fcrType) SetIosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0fcrFieldIosrMask)|(uint32(value)<<RegisterDfsdmflt0fcrFieldIosrShift))
}

const (
	RegisterDfsdmflt0fcrFieldFosrShift = 16
	RegisterDfsdmflt0fcrFieldFosrMask  = 0x3ff0000
)

// GetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdmflt0fcrType) GetFosr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0fcrFieldFosrMask) >> RegisterDfsdmflt0fcrFieldFosrShift)
}

// SetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdmflt0fcrType) SetFosr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0fcrFieldFosrMask)|(uint32(value)<<RegisterDfsdmflt0fcrFieldFosrShift))
}

const (
	RegisterDfsdmflt0fcrFieldFordShift = 29
	RegisterDfsdmflt0fcrFieldFordMask  = 0xe0000000
)

// GetFord Sinc filter order
func (r *registerDfsdmflt0fcrType) GetFord() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0fcrFieldFordMask) >> RegisterDfsdmflt0fcrFieldFordShift)
}

// SetFord Sinc filter order
func (r *registerDfsdmflt0fcrType) SetFord(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0fcrFieldFordMask)|(uint32(value)<<RegisterDfsdmflt0fcrFieldFordShift))
}

// registerDfsdmflt0jdatarType data register for injected group
type registerDfsdmflt0jdatarType uint32

const (
	RegisterDfsdmflt0jdatarFieldJdatachShift = 0
	RegisterDfsdmflt0jdatarFieldJdatachMask  = 0x7
)

// GetJdatach Injected channel most recently converted
func (r *registerDfsdmflt0jdatarType) GetJdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0jdatarFieldJdatachMask) >> RegisterDfsdmflt0jdatarFieldJdatachShift)
}

// SetJdatach Injected channel most recently converted
func (r *registerDfsdmflt0jdatarType) SetJdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0jdatarFieldJdatachMask)|(uint32(value)<<RegisterDfsdmflt0jdatarFieldJdatachShift))
}

const (
	RegisterDfsdmflt0jdatarFieldJdataShift = 8
	RegisterDfsdmflt0jdatarFieldJdataMask  = 0xffffff00
)

// GetJdata Injected group conversion data
func (r *registerDfsdmflt0jdatarType) GetJdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0jdatarFieldJdataMask) >> RegisterDfsdmflt0jdatarFieldJdataShift)
}

// SetJdata Injected group conversion data
func (r *registerDfsdmflt0jdatarType) SetJdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0jdatarFieldJdataMask)|(uint32(value)<<RegisterDfsdmflt0jdatarFieldJdataShift))
}

// registerDfsdmflt0rdatarType data register for the regular channel
type registerDfsdmflt0rdatarType uint32

const (
	RegisterDfsdmflt0rdatarFieldRdatachShift = 0
	RegisterDfsdmflt0rdatarFieldRdatachMask  = 0x7
)

// GetRdatach Regular channel most recently converted
func (r *registerDfsdmflt0rdatarType) GetRdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0rdatarFieldRdatachMask) >> RegisterDfsdmflt0rdatarFieldRdatachShift)
}

// SetRdatach Regular channel most recently converted
func (r *registerDfsdmflt0rdatarType) SetRdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0rdatarFieldRdatachMask)|(uint32(value)<<RegisterDfsdmflt0rdatarFieldRdatachShift))
}

const (
	RegisterDfsdmflt0rdatarFieldRpendShift = 4
	RegisterDfsdmflt0rdatarFieldRpendMask  = 0x10
)

// GetRpend Regular channel pending data
func (r *registerDfsdmflt0rdatarType) GetRpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0rdatarFieldRpendMask) != 0
}

// SetRpend Regular channel pending data
func (r *registerDfsdmflt0rdatarType) SetRpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt0rdatarFieldRpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0rdatarFieldRpendMask)
	}
}

const (
	RegisterDfsdmflt0rdatarFieldRdataShift = 8
	RegisterDfsdmflt0rdatarFieldRdataMask  = 0xffffff00
)

// GetRdata Regular channel conversion data
func (r *registerDfsdmflt0rdatarType) GetRdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0rdatarFieldRdataMask) >> RegisterDfsdmflt0rdatarFieldRdataShift)
}

// SetRdata Regular channel conversion data
func (r *registerDfsdmflt0rdatarType) SetRdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0rdatarFieldRdataMask)|(uint32(value)<<RegisterDfsdmflt0rdatarFieldRdataShift))
}

// registerDfsdmflt0awhtrType analog watchdog high threshold register
type registerDfsdmflt0awhtrType uint32

const (
	RegisterDfsdmflt0awhtrFieldBkawhShift = 0
	RegisterDfsdmflt0awhtrFieldBkawhMask  = 0xf
)

// GetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdmflt0awhtrType) GetBkawh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0awhtrFieldBkawhMask) >> RegisterDfsdmflt0awhtrFieldBkawhShift)
}

// SetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdmflt0awhtrType) SetBkawh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0awhtrFieldBkawhMask)|(uint32(value)<<RegisterDfsdmflt0awhtrFieldBkawhShift))
}

const (
	RegisterDfsdmflt0awhtrFieldAwhtShift = 8
	RegisterDfsdmflt0awhtrFieldAwhtMask  = 0xffffff00
)

// GetAwht Analog watchdog high threshold
func (r *registerDfsdmflt0awhtrType) GetAwht() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0awhtrFieldAwhtMask) >> RegisterDfsdmflt0awhtrFieldAwhtShift)
}

// SetAwht Analog watchdog high threshold
func (r *registerDfsdmflt0awhtrType) SetAwht(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0awhtrFieldAwhtMask)|(uint32(value)<<RegisterDfsdmflt0awhtrFieldAwhtShift))
}

// registerDfsdmflt0awltrType analog watchdog low threshold register
type registerDfsdmflt0awltrType uint32

const (
	RegisterDfsdmflt0awltrFieldBkawlShift = 0
	RegisterDfsdmflt0awltrFieldBkawlMask  = 0xf
)

// GetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdmflt0awltrType) GetBkawl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0awltrFieldBkawlMask) >> RegisterDfsdmflt0awltrFieldBkawlShift)
}

// SetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdmflt0awltrType) SetBkawl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0awltrFieldBkawlMask)|(uint32(value)<<RegisterDfsdmflt0awltrFieldBkawlShift))
}

const (
	RegisterDfsdmflt0awltrFieldAwltShift = 8
	RegisterDfsdmflt0awltrFieldAwltMask  = 0xffffff00
)

// GetAwlt Analog watchdog low threshold
func (r *registerDfsdmflt0awltrType) GetAwlt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0awltrFieldAwltMask) >> RegisterDfsdmflt0awltrFieldAwltShift)
}

// SetAwlt Analog watchdog low threshold
func (r *registerDfsdmflt0awltrType) SetAwlt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0awltrFieldAwltMask)|(uint32(value)<<RegisterDfsdmflt0awltrFieldAwltShift))
}

// registerDfsdmflt0awsrType analog watchdog status register
type registerDfsdmflt0awsrType uint32

const (
	RegisterDfsdmflt0awsrFieldAwltfShift = 0
	RegisterDfsdmflt0awsrFieldAwltfMask  = 0xff
)

// GetAwltf Analog watchdog low threshold flag
func (r *registerDfsdmflt0awsrType) GetAwltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0awsrFieldAwltfMask) >> RegisterDfsdmflt0awsrFieldAwltfShift)
}

// SetAwltf Analog watchdog low threshold flag
func (r *registerDfsdmflt0awsrType) SetAwltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0awsrFieldAwltfMask)|(uint32(value)<<RegisterDfsdmflt0awsrFieldAwltfShift))
}

const (
	RegisterDfsdmflt0awsrFieldAwhtfShift = 8
	RegisterDfsdmflt0awsrFieldAwhtfMask  = 0xff00
)

// GetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdmflt0awsrType) GetAwhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0awsrFieldAwhtfMask) >> RegisterDfsdmflt0awsrFieldAwhtfShift)
}

// SetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdmflt0awsrType) SetAwhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0awsrFieldAwhtfMask)|(uint32(value)<<RegisterDfsdmflt0awsrFieldAwhtfShift))
}

// registerDfsdmflt0awcfrType analog watchdog clear flag register
type registerDfsdmflt0awcfrType uint32

const (
	RegisterDfsdmflt0awcfrFieldClrawltfShift = 0
	RegisterDfsdmflt0awcfrFieldClrawltfMask  = 0xff
)

// GetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdmflt0awcfrType) GetClrawltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0awcfrFieldClrawltfMask) >> RegisterDfsdmflt0awcfrFieldClrawltfShift)
}

// SetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdmflt0awcfrType) SetClrawltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0awcfrFieldClrawltfMask)|(uint32(value)<<RegisterDfsdmflt0awcfrFieldClrawltfShift))
}

const (
	RegisterDfsdmflt0awcfrFieldClrawhtfShift = 8
	RegisterDfsdmflt0awcfrFieldClrawhtfMask  = 0xff00
)

// GetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdmflt0awcfrType) GetClrawhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0awcfrFieldClrawhtfMask) >> RegisterDfsdmflt0awcfrFieldClrawhtfShift)
}

// SetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdmflt0awcfrType) SetClrawhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0awcfrFieldClrawhtfMask)|(uint32(value)<<RegisterDfsdmflt0awcfrFieldClrawhtfShift))
}

// registerDfsdmflt0exmaxType Extremes detector maximum register
type registerDfsdmflt0exmaxType uint32

const (
	RegisterDfsdmflt0exmaxFieldExmaxchShift = 0
	RegisterDfsdmflt0exmaxFieldExmaxchMask  = 0x7
)

// GetExmaxch Extremes detector maximum data channel
func (r *registerDfsdmflt0exmaxType) GetExmaxch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0exmaxFieldExmaxchMask) >> RegisterDfsdmflt0exmaxFieldExmaxchShift)
}

// SetExmaxch Extremes detector maximum data channel
func (r *registerDfsdmflt0exmaxType) SetExmaxch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0exmaxFieldExmaxchMask)|(uint32(value)<<RegisterDfsdmflt0exmaxFieldExmaxchShift))
}

const (
	RegisterDfsdmflt0exmaxFieldExmaxShift = 8
	RegisterDfsdmflt0exmaxFieldExmaxMask  = 0xffffff00
)

// GetExmax Extremes detector maximum value
func (r *registerDfsdmflt0exmaxType) GetExmax() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0exmaxFieldExmaxMask) >> RegisterDfsdmflt0exmaxFieldExmaxShift)
}

// SetExmax Extremes detector maximum value
func (r *registerDfsdmflt0exmaxType) SetExmax(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0exmaxFieldExmaxMask)|(uint32(value)<<RegisterDfsdmflt0exmaxFieldExmaxShift))
}

// registerDfsdmflt0exminType Extremes detector minimum register
type registerDfsdmflt0exminType uint32

const (
	RegisterDfsdmflt0exminFieldExminchShift = 0
	RegisterDfsdmflt0exminFieldExminchMask  = 0x7
)

// GetExminch Extremes detector minimum data channel
func (r *registerDfsdmflt0exminType) GetExminch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0exminFieldExminchMask) >> RegisterDfsdmflt0exminFieldExminchShift)
}

// SetExminch Extremes detector minimum data channel
func (r *registerDfsdmflt0exminType) SetExminch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0exminFieldExminchMask)|(uint32(value)<<RegisterDfsdmflt0exminFieldExminchShift))
}

const (
	RegisterDfsdmflt0exminFieldExminShift = 8
	RegisterDfsdmflt0exminFieldExminMask  = 0xffffff00
)

// GetExmin EXMIN
func (r *registerDfsdmflt0exminType) GetExmin() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0exminFieldExminMask) >> RegisterDfsdmflt0exminFieldExminShift)
}

// SetExmin EXMIN
func (r *registerDfsdmflt0exminType) SetExmin(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0exminFieldExminMask)|(uint32(value)<<RegisterDfsdmflt0exminFieldExminShift))
}

// registerDfsdmflt0cnvtimrType conversion timer register
type registerDfsdmflt0cnvtimrType uint32

const (
	RegisterDfsdmflt0cnvtimrFieldCnvcntShift = 4
	RegisterDfsdmflt0cnvtimrFieldCnvcntMask  = 0xfffffff0
)

// GetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdmflt0cnvtimrType) GetCnvcnt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt0cnvtimrFieldCnvcntMask) >> RegisterDfsdmflt0cnvtimrFieldCnvcntShift)
}

// SetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdmflt0cnvtimrType) SetCnvcnt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt0cnvtimrFieldCnvcntMask)|(uint32(value)<<RegisterDfsdmflt0cnvtimrFieldCnvcntShift))
}

// registerDfsdmflt1cr1Type control register 1
type registerDfsdmflt1cr1Type uint32

const (
	RegisterDfsdmflt1cr1FieldDfenShift = 0
	RegisterDfsdmflt1cr1FieldDfenMask  = 0x1
)

// GetDfen DFSDM enable
func (r *registerDfsdmflt1cr1Type) GetDfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldDfenMask) != 0
}

// SetDfen DFSDM enable
func (r *registerDfsdmflt1cr1Type) SetDfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr1FieldDfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldDfenMask)
	}
}

const (
	RegisterDfsdmflt1cr1FieldJswstartShift = 1
	RegisterDfsdmflt1cr1FieldJswstartMask  = 0x2
)

// GetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdmflt1cr1Type) GetJswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldJswstartMask) != 0
}

// SetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdmflt1cr1Type) SetJswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr1FieldJswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldJswstartMask)
	}
}

const (
	RegisterDfsdmflt1cr1FieldJsyncShift = 3
	RegisterDfsdmflt1cr1FieldJsyncMask  = 0x8
)

// GetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdmflt1cr1Type) GetJsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldJsyncMask) != 0
}

// SetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdmflt1cr1Type) SetJsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr1FieldJsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldJsyncMask)
	}
}

const (
	RegisterDfsdmflt1cr1FieldJscanShift = 4
	RegisterDfsdmflt1cr1FieldJscanMask  = 0x10
)

// GetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdmflt1cr1Type) GetJscan() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldJscanMask) != 0
}

// SetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdmflt1cr1Type) SetJscan(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr1FieldJscanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldJscanMask)
	}
}

const (
	RegisterDfsdmflt1cr1FieldJdmaenShift = 5
	RegisterDfsdmflt1cr1FieldJdmaenMask  = 0x20
)

// GetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdmflt1cr1Type) GetJdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldJdmaenMask) != 0
}

// SetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdmflt1cr1Type) SetJdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr1FieldJdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldJdmaenMask)
	}
}

const (
	RegisterDfsdmflt1cr1FieldJextselShift = 8
	RegisterDfsdmflt1cr1FieldJextselMask  = 0x700
)

// GetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdmflt1cr1Type) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldJextselMask) >> RegisterDfsdmflt1cr1FieldJextselShift)
}

// SetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdmflt1cr1Type) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldJextselMask)|(uint32(value)<<RegisterDfsdmflt1cr1FieldJextselShift))
}

const (
	RegisterDfsdmflt1cr1FieldJextenShift = 13
	RegisterDfsdmflt1cr1FieldJextenMask  = 0x6000
)

// GetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdmflt1cr1Type) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldJextenMask) >> RegisterDfsdmflt1cr1FieldJextenShift)
}

// SetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdmflt1cr1Type) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldJextenMask)|(uint32(value)<<RegisterDfsdmflt1cr1FieldJextenShift))
}

const (
	RegisterDfsdmflt1cr1FieldRswstartShift = 17
	RegisterDfsdmflt1cr1FieldRswstartMask  = 0x20000
)

// GetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdmflt1cr1Type) GetRswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldRswstartMask) != 0
}

// SetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdmflt1cr1Type) SetRswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr1FieldRswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldRswstartMask)
	}
}

const (
	RegisterDfsdmflt1cr1FieldRcontShift = 18
	RegisterDfsdmflt1cr1FieldRcontMask  = 0x40000
)

// GetRcont Continuous mode selection for regular conversions
func (r *registerDfsdmflt1cr1Type) GetRcont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldRcontMask) != 0
}

// SetRcont Continuous mode selection for regular conversions
func (r *registerDfsdmflt1cr1Type) SetRcont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr1FieldRcontMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldRcontMask)
	}
}

const (
	RegisterDfsdmflt1cr1FieldRsyncShift = 19
	RegisterDfsdmflt1cr1FieldRsyncMask  = 0x80000
)

// GetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdmflt1cr1Type) GetRsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldRsyncMask) != 0
}

// SetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdmflt1cr1Type) SetRsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr1FieldRsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldRsyncMask)
	}
}

const (
	RegisterDfsdmflt1cr1FieldRdmaenShift = 21
	RegisterDfsdmflt1cr1FieldRdmaenMask  = 0x200000
)

// GetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdmflt1cr1Type) GetRdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldRdmaenMask) != 0
}

// SetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdmflt1cr1Type) SetRdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr1FieldRdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldRdmaenMask)
	}
}

const (
	RegisterDfsdmflt1cr1FieldRchShift = 24
	RegisterDfsdmflt1cr1FieldRchMask  = 0x7000000
)

// GetRch Regular channel selection
func (r *registerDfsdmflt1cr1Type) GetRch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldRchMask) >> RegisterDfsdmflt1cr1FieldRchShift)
}

// SetRch Regular channel selection
func (r *registerDfsdmflt1cr1Type) SetRch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldRchMask)|(uint32(value)<<RegisterDfsdmflt1cr1FieldRchShift))
}

const (
	RegisterDfsdmflt1cr1FieldFastShift = 29
	RegisterDfsdmflt1cr1FieldFastMask  = 0x20000000
)

// GetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdmflt1cr1Type) GetFast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldFastMask) != 0
}

// SetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdmflt1cr1Type) SetFast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr1FieldFastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldFastMask)
	}
}

const (
	RegisterDfsdmflt1cr1FieldAwfselShift = 30
	RegisterDfsdmflt1cr1FieldAwfselMask  = 0x40000000
)

// GetAwfsel Analog watchdog fast mode select
func (r *registerDfsdmflt1cr1Type) GetAwfsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr1FieldAwfselMask) != 0
}

// SetAwfsel Analog watchdog fast mode select
func (r *registerDfsdmflt1cr1Type) SetAwfsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr1FieldAwfselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr1FieldAwfselMask)
	}
}

// registerDfsdmflt1cr2Type control register 2
type registerDfsdmflt1cr2Type uint32

const (
	RegisterDfsdmflt1cr2FieldJeocieShift = 0
	RegisterDfsdmflt1cr2FieldJeocieMask  = 0x1
)

// GetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdmflt1cr2Type) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr2FieldJeocieMask) != 0
}

// SetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdmflt1cr2Type) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr2FieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr2FieldJeocieMask)
	}
}

const (
	RegisterDfsdmflt1cr2FieldReocieShift = 1
	RegisterDfsdmflt1cr2FieldReocieMask  = 0x2
)

// GetReocie Regular end of conversion interrupt enable
func (r *registerDfsdmflt1cr2Type) GetReocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr2FieldReocieMask) != 0
}

// SetReocie Regular end of conversion interrupt enable
func (r *registerDfsdmflt1cr2Type) SetReocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr2FieldReocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr2FieldReocieMask)
	}
}

const (
	RegisterDfsdmflt1cr2FieldJovrieShift = 2
	RegisterDfsdmflt1cr2FieldJovrieMask  = 0x4
)

// GetJovrie Injected data overrun interrupt enable
func (r *registerDfsdmflt1cr2Type) GetJovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr2FieldJovrieMask) != 0
}

// SetJovrie Injected data overrun interrupt enable
func (r *registerDfsdmflt1cr2Type) SetJovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr2FieldJovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr2FieldJovrieMask)
	}
}

const (
	RegisterDfsdmflt1cr2FieldRovrieShift = 3
	RegisterDfsdmflt1cr2FieldRovrieMask  = 0x8
)

// GetRovrie Regular data overrun interrupt enable
func (r *registerDfsdmflt1cr2Type) GetRovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr2FieldRovrieMask) != 0
}

// SetRovrie Regular data overrun interrupt enable
func (r *registerDfsdmflt1cr2Type) SetRovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr2FieldRovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr2FieldRovrieMask)
	}
}

const (
	RegisterDfsdmflt1cr2FieldAwdieShift = 4
	RegisterDfsdmflt1cr2FieldAwdieMask  = 0x10
)

// GetAwdie Analog watchdog interrupt enable
func (r *registerDfsdmflt1cr2Type) GetAwdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr2FieldAwdieMask) != 0
}

// SetAwdie Analog watchdog interrupt enable
func (r *registerDfsdmflt1cr2Type) SetAwdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr2FieldAwdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr2FieldAwdieMask)
	}
}

const (
	RegisterDfsdmflt1cr2FieldScdieShift = 5
	RegisterDfsdmflt1cr2FieldScdieMask  = 0x20
)

// GetScdie Short-circuit detector interrupt enable
func (r *registerDfsdmflt1cr2Type) GetScdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr2FieldScdieMask) != 0
}

// SetScdie Short-circuit detector interrupt enable
func (r *registerDfsdmflt1cr2Type) SetScdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr2FieldScdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr2FieldScdieMask)
	}
}

const (
	RegisterDfsdmflt1cr2FieldCkabieShift = 6
	RegisterDfsdmflt1cr2FieldCkabieMask  = 0x40
)

// GetCkabie Clock absence interrupt enable
func (r *registerDfsdmflt1cr2Type) GetCkabie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr2FieldCkabieMask) != 0
}

// SetCkabie Clock absence interrupt enable
func (r *registerDfsdmflt1cr2Type) SetCkabie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1cr2FieldCkabieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr2FieldCkabieMask)
	}
}

const (
	RegisterDfsdmflt1cr2FieldExchShift = 8
	RegisterDfsdmflt1cr2FieldExchMask  = 0xff00
)

// GetExch Extremes detector channel selection
func (r *registerDfsdmflt1cr2Type) GetExch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr2FieldExchMask) >> RegisterDfsdmflt1cr2FieldExchShift)
}

// SetExch Extremes detector channel selection
func (r *registerDfsdmflt1cr2Type) SetExch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr2FieldExchMask)|(uint32(value)<<RegisterDfsdmflt1cr2FieldExchShift))
}

const (
	RegisterDfsdmflt1cr2FieldAwdchShift = 16
	RegisterDfsdmflt1cr2FieldAwdchMask  = 0xff0000
)

// GetAwdch Analog watchdog channel selection
func (r *registerDfsdmflt1cr2Type) GetAwdch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cr2FieldAwdchMask) >> RegisterDfsdmflt1cr2FieldAwdchShift)
}

// SetAwdch Analog watchdog channel selection
func (r *registerDfsdmflt1cr2Type) SetAwdch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cr2FieldAwdchMask)|(uint32(value)<<RegisterDfsdmflt1cr2FieldAwdchShift))
}

// registerDfsdmflt1isrType interrupt and status register
type registerDfsdmflt1isrType uint32

const (
	RegisterDfsdmflt1isrFieldJeocfShift = 0
	RegisterDfsdmflt1isrFieldJeocfMask  = 0x1
)

// GetJeocf End of injected conversion flag
func (r *registerDfsdmflt1isrType) GetJeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1isrFieldJeocfMask) != 0
}

// SetJeocf End of injected conversion flag
func (r *registerDfsdmflt1isrType) SetJeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1isrFieldJeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1isrFieldJeocfMask)
	}
}

const (
	RegisterDfsdmflt1isrFieldReocfShift = 1
	RegisterDfsdmflt1isrFieldReocfMask  = 0x2
)

// GetReocf End of regular conversion flag
func (r *registerDfsdmflt1isrType) GetReocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1isrFieldReocfMask) != 0
}

// SetReocf End of regular conversion flag
func (r *registerDfsdmflt1isrType) SetReocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1isrFieldReocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1isrFieldReocfMask)
	}
}

const (
	RegisterDfsdmflt1isrFieldJovrfShift = 2
	RegisterDfsdmflt1isrFieldJovrfMask  = 0x4
)

// GetJovrf Injected conversion overrun flag
func (r *registerDfsdmflt1isrType) GetJovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1isrFieldJovrfMask) != 0
}

// SetJovrf Injected conversion overrun flag
func (r *registerDfsdmflt1isrType) SetJovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1isrFieldJovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1isrFieldJovrfMask)
	}
}

const (
	RegisterDfsdmflt1isrFieldRovrfShift = 3
	RegisterDfsdmflt1isrFieldRovrfMask  = 0x8
)

// GetRovrf Regular conversion overrun flag
func (r *registerDfsdmflt1isrType) GetRovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1isrFieldRovrfMask) != 0
}

// SetRovrf Regular conversion overrun flag
func (r *registerDfsdmflt1isrType) SetRovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1isrFieldRovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1isrFieldRovrfMask)
	}
}

const (
	RegisterDfsdmflt1isrFieldAwdfShift = 4
	RegisterDfsdmflt1isrFieldAwdfMask  = 0x10
)

// GetAwdf Analog watchdog
func (r *registerDfsdmflt1isrType) GetAwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1isrFieldAwdfMask) != 0
}

// SetAwdf Analog watchdog
func (r *registerDfsdmflt1isrType) SetAwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1isrFieldAwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1isrFieldAwdfMask)
	}
}

const (
	RegisterDfsdmflt1isrFieldJcipShift = 13
	RegisterDfsdmflt1isrFieldJcipMask  = 0x2000
)

// GetJcip Injected conversion in progress status
func (r *registerDfsdmflt1isrType) GetJcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1isrFieldJcipMask) != 0
}

// SetJcip Injected conversion in progress status
func (r *registerDfsdmflt1isrType) SetJcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1isrFieldJcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1isrFieldJcipMask)
	}
}

const (
	RegisterDfsdmflt1isrFieldRcipShift = 14
	RegisterDfsdmflt1isrFieldRcipMask  = 0x4000
)

// GetRcip Regular conversion in progress status
func (r *registerDfsdmflt1isrType) GetRcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1isrFieldRcipMask) != 0
}

// SetRcip Regular conversion in progress status
func (r *registerDfsdmflt1isrType) SetRcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1isrFieldRcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1isrFieldRcipMask)
	}
}

const (
	RegisterDfsdmflt1isrFieldCkabfShift = 16
	RegisterDfsdmflt1isrFieldCkabfMask  = 0xff0000
)

// GetCkabf Clock absence flag
func (r *registerDfsdmflt1isrType) GetCkabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1isrFieldCkabfMask) >> RegisterDfsdmflt1isrFieldCkabfShift)
}

// SetCkabf Clock absence flag
func (r *registerDfsdmflt1isrType) SetCkabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1isrFieldCkabfMask)|(uint32(value)<<RegisterDfsdmflt1isrFieldCkabfShift))
}

const (
	RegisterDfsdmflt1isrFieldScdfShift = 24
	RegisterDfsdmflt1isrFieldScdfMask  = 0xff000000
)

// GetScdf short-circuit detector flag
func (r *registerDfsdmflt1isrType) GetScdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1isrFieldScdfMask) >> RegisterDfsdmflt1isrFieldScdfShift)
}

// SetScdf short-circuit detector flag
func (r *registerDfsdmflt1isrType) SetScdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1isrFieldScdfMask)|(uint32(value)<<RegisterDfsdmflt1isrFieldScdfShift))
}

// registerDfsdm1icrType interrupt flag clear register
type registerDfsdm1icrType uint32

const (
	RegisterDfsdm1icrFieldClrjovrfShift = 2
	RegisterDfsdm1icrFieldClrjovrfMask  = 0x4
)

// GetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm1icrType) GetClrjovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1icrFieldClrjovrfMask) != 0
}

// SetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdm1icrType) SetClrjovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1icrFieldClrjovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1icrFieldClrjovrfMask)
	}
}

const (
	RegisterDfsdm1icrFieldClrrovrfShift = 3
	RegisterDfsdm1icrFieldClrrovrfMask  = 0x8
)

// GetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm1icrType) GetClrrovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1icrFieldClrrovrfMask) != 0
}

// SetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdm1icrType) SetClrrovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdm1icrFieldClrrovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1icrFieldClrrovrfMask)
	}
}

const (
	RegisterDfsdm1icrFieldClrckabfShift = 16
	RegisterDfsdm1icrFieldClrckabfMask  = 0xff0000
)

// GetClrckabf Clear the clock absence flag
func (r *registerDfsdm1icrType) GetClrckabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1icrFieldClrckabfMask) >> RegisterDfsdm1icrFieldClrckabfShift)
}

// SetClrckabf Clear the clock absence flag
func (r *registerDfsdm1icrType) SetClrckabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1icrFieldClrckabfMask)|(uint32(value)<<RegisterDfsdm1icrFieldClrckabfShift))
}

const (
	RegisterDfsdm1icrFieldClrscdfShift = 24
	RegisterDfsdm1icrFieldClrscdfMask  = 0xff000000
)

// GetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm1icrType) GetClrscdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1icrFieldClrscdfMask) >> RegisterDfsdm1icrFieldClrscdfShift)
}

// SetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdm1icrType) SetClrscdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1icrFieldClrscdfMask)|(uint32(value)<<RegisterDfsdm1icrFieldClrscdfShift))
}

// registerDfsdmflt1jchgrType injected channel group selection register
type registerDfsdmflt1jchgrType uint32

const (
	RegisterDfsdmflt1jchgrFieldJchgShift = 0
	RegisterDfsdmflt1jchgrFieldJchgMask  = 0xff
)

// GetJchg Injected channel group selection
func (r *registerDfsdmflt1jchgrType) GetJchg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1jchgrFieldJchgMask) >> RegisterDfsdmflt1jchgrFieldJchgShift)
}

// SetJchg Injected channel group selection
func (r *registerDfsdmflt1jchgrType) SetJchg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1jchgrFieldJchgMask)|(uint32(value)<<RegisterDfsdmflt1jchgrFieldJchgShift))
}

// registerDfsdm1fcrType filter control register
type registerDfsdm1fcrType uint32

const (
	RegisterDfsdm1fcrFieldIosrShift = 0
	RegisterDfsdm1fcrFieldIosrMask  = 0xff
)

// GetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm1fcrType) GetIosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1fcrFieldIosrMask) >> RegisterDfsdm1fcrFieldIosrShift)
}

// SetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdm1fcrType) SetIosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1fcrFieldIosrMask)|(uint32(value)<<RegisterDfsdm1fcrFieldIosrShift))
}

const (
	RegisterDfsdm1fcrFieldFosrShift = 16
	RegisterDfsdm1fcrFieldFosrMask  = 0x3ff0000
)

// GetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm1fcrType) GetFosr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1fcrFieldFosrMask) >> RegisterDfsdm1fcrFieldFosrShift)
}

// SetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdm1fcrType) SetFosr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1fcrFieldFosrMask)|(uint32(value)<<RegisterDfsdm1fcrFieldFosrShift))
}

const (
	RegisterDfsdm1fcrFieldFordShift = 29
	RegisterDfsdm1fcrFieldFordMask  = 0xe0000000
)

// GetFord Sinc filter order
func (r *registerDfsdm1fcrType) GetFord() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdm1fcrFieldFordMask) >> RegisterDfsdm1fcrFieldFordShift)
}

// SetFord Sinc filter order
func (r *registerDfsdm1fcrType) SetFord(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdm1fcrFieldFordMask)|(uint32(value)<<RegisterDfsdm1fcrFieldFordShift))
}

// registerDfsdmflt1jdatarType data register for injected group
type registerDfsdmflt1jdatarType uint32

const (
	RegisterDfsdmflt1jdatarFieldJdatachShift = 0
	RegisterDfsdmflt1jdatarFieldJdatachMask  = 0x7
)

// GetJdatach Injected channel most recently converted
func (r *registerDfsdmflt1jdatarType) GetJdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1jdatarFieldJdatachMask) >> RegisterDfsdmflt1jdatarFieldJdatachShift)
}

// SetJdatach Injected channel most recently converted
func (r *registerDfsdmflt1jdatarType) SetJdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1jdatarFieldJdatachMask)|(uint32(value)<<RegisterDfsdmflt1jdatarFieldJdatachShift))
}

const (
	RegisterDfsdmflt1jdatarFieldJdataShift = 8
	RegisterDfsdmflt1jdatarFieldJdataMask  = 0xffffff00
)

// GetJdata Injected group conversion data
func (r *registerDfsdmflt1jdatarType) GetJdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1jdatarFieldJdataMask) >> RegisterDfsdmflt1jdatarFieldJdataShift)
}

// SetJdata Injected group conversion data
func (r *registerDfsdmflt1jdatarType) SetJdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1jdatarFieldJdataMask)|(uint32(value)<<RegisterDfsdmflt1jdatarFieldJdataShift))
}

// registerDfsdmflt1rdatarType data register for the regular channel
type registerDfsdmflt1rdatarType uint32

const (
	RegisterDfsdmflt1rdatarFieldRdatachShift = 0
	RegisterDfsdmflt1rdatarFieldRdatachMask  = 0x7
)

// GetRdatach Regular channel most recently converted
func (r *registerDfsdmflt1rdatarType) GetRdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1rdatarFieldRdatachMask) >> RegisterDfsdmflt1rdatarFieldRdatachShift)
}

// SetRdatach Regular channel most recently converted
func (r *registerDfsdmflt1rdatarType) SetRdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1rdatarFieldRdatachMask)|(uint32(value)<<RegisterDfsdmflt1rdatarFieldRdatachShift))
}

const (
	RegisterDfsdmflt1rdatarFieldRpendShift = 4
	RegisterDfsdmflt1rdatarFieldRpendMask  = 0x10
)

// GetRpend Regular channel pending data
func (r *registerDfsdmflt1rdatarType) GetRpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1rdatarFieldRpendMask) != 0
}

// SetRpend Regular channel pending data
func (r *registerDfsdmflt1rdatarType) SetRpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt1rdatarFieldRpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1rdatarFieldRpendMask)
	}
}

const (
	RegisterDfsdmflt1rdatarFieldRdataShift = 8
	RegisterDfsdmflt1rdatarFieldRdataMask  = 0xffffff00
)

// GetRdata Regular channel conversion data
func (r *registerDfsdmflt1rdatarType) GetRdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1rdatarFieldRdataMask) >> RegisterDfsdmflt1rdatarFieldRdataShift)
}

// SetRdata Regular channel conversion data
func (r *registerDfsdmflt1rdatarType) SetRdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1rdatarFieldRdataMask)|(uint32(value)<<RegisterDfsdmflt1rdatarFieldRdataShift))
}

// registerDfsdmflt1awhtrType analog watchdog high threshold register
type registerDfsdmflt1awhtrType uint32

const (
	RegisterDfsdmflt1awhtrFieldBkawhShift = 0
	RegisterDfsdmflt1awhtrFieldBkawhMask  = 0xf
)

// GetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdmflt1awhtrType) GetBkawh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1awhtrFieldBkawhMask) >> RegisterDfsdmflt1awhtrFieldBkawhShift)
}

// SetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdmflt1awhtrType) SetBkawh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1awhtrFieldBkawhMask)|(uint32(value)<<RegisterDfsdmflt1awhtrFieldBkawhShift))
}

const (
	RegisterDfsdmflt1awhtrFieldAwhtShift = 8
	RegisterDfsdmflt1awhtrFieldAwhtMask  = 0xffffff00
)

// GetAwht Analog watchdog high threshold
func (r *registerDfsdmflt1awhtrType) GetAwht() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1awhtrFieldAwhtMask) >> RegisterDfsdmflt1awhtrFieldAwhtShift)
}

// SetAwht Analog watchdog high threshold
func (r *registerDfsdmflt1awhtrType) SetAwht(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1awhtrFieldAwhtMask)|(uint32(value)<<RegisterDfsdmflt1awhtrFieldAwhtShift))
}

// registerDfsdmflt1awltrType analog watchdog low threshold register
type registerDfsdmflt1awltrType uint32

const (
	RegisterDfsdmflt1awltrFieldBkawlShift = 0
	RegisterDfsdmflt1awltrFieldBkawlMask  = 0xf
)

// GetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdmflt1awltrType) GetBkawl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1awltrFieldBkawlMask) >> RegisterDfsdmflt1awltrFieldBkawlShift)
}

// SetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdmflt1awltrType) SetBkawl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1awltrFieldBkawlMask)|(uint32(value)<<RegisterDfsdmflt1awltrFieldBkawlShift))
}

const (
	RegisterDfsdmflt1awltrFieldAwltShift = 8
	RegisterDfsdmflt1awltrFieldAwltMask  = 0xffffff00
)

// GetAwlt Analog watchdog low threshold
func (r *registerDfsdmflt1awltrType) GetAwlt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1awltrFieldAwltMask) >> RegisterDfsdmflt1awltrFieldAwltShift)
}

// SetAwlt Analog watchdog low threshold
func (r *registerDfsdmflt1awltrType) SetAwlt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1awltrFieldAwltMask)|(uint32(value)<<RegisterDfsdmflt1awltrFieldAwltShift))
}

// registerDfsdmflt1awsrType analog watchdog status register
type registerDfsdmflt1awsrType uint32

const (
	RegisterDfsdmflt1awsrFieldAwltfShift = 0
	RegisterDfsdmflt1awsrFieldAwltfMask  = 0xff
)

// GetAwltf Analog watchdog low threshold flag
func (r *registerDfsdmflt1awsrType) GetAwltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1awsrFieldAwltfMask) >> RegisterDfsdmflt1awsrFieldAwltfShift)
}

// SetAwltf Analog watchdog low threshold flag
func (r *registerDfsdmflt1awsrType) SetAwltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1awsrFieldAwltfMask)|(uint32(value)<<RegisterDfsdmflt1awsrFieldAwltfShift))
}

const (
	RegisterDfsdmflt1awsrFieldAwhtfShift = 8
	RegisterDfsdmflt1awsrFieldAwhtfMask  = 0xff00
)

// GetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdmflt1awsrType) GetAwhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1awsrFieldAwhtfMask) >> RegisterDfsdmflt1awsrFieldAwhtfShift)
}

// SetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdmflt1awsrType) SetAwhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1awsrFieldAwhtfMask)|(uint32(value)<<RegisterDfsdmflt1awsrFieldAwhtfShift))
}

// registerDfsdmflt1awcfrType analog watchdog clear flag register
type registerDfsdmflt1awcfrType uint32

const (
	RegisterDfsdmflt1awcfrFieldClrawltfShift = 0
	RegisterDfsdmflt1awcfrFieldClrawltfMask  = 0xff
)

// GetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdmflt1awcfrType) GetClrawltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1awcfrFieldClrawltfMask) >> RegisterDfsdmflt1awcfrFieldClrawltfShift)
}

// SetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdmflt1awcfrType) SetClrawltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1awcfrFieldClrawltfMask)|(uint32(value)<<RegisterDfsdmflt1awcfrFieldClrawltfShift))
}

const (
	RegisterDfsdmflt1awcfrFieldClrawhtfShift = 8
	RegisterDfsdmflt1awcfrFieldClrawhtfMask  = 0xff00
)

// GetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdmflt1awcfrType) GetClrawhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1awcfrFieldClrawhtfMask) >> RegisterDfsdmflt1awcfrFieldClrawhtfShift)
}

// SetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdmflt1awcfrType) SetClrawhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1awcfrFieldClrawhtfMask)|(uint32(value)<<RegisterDfsdmflt1awcfrFieldClrawhtfShift))
}

// registerDfsdmflt1exmaxType Extremes detector maximum register
type registerDfsdmflt1exmaxType uint32

const (
	RegisterDfsdmflt1exmaxFieldExmaxchShift = 0
	RegisterDfsdmflt1exmaxFieldExmaxchMask  = 0x7
)

// GetExmaxch Extremes detector maximum data channel
func (r *registerDfsdmflt1exmaxType) GetExmaxch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1exmaxFieldExmaxchMask) >> RegisterDfsdmflt1exmaxFieldExmaxchShift)
}

// SetExmaxch Extremes detector maximum data channel
func (r *registerDfsdmflt1exmaxType) SetExmaxch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1exmaxFieldExmaxchMask)|(uint32(value)<<RegisterDfsdmflt1exmaxFieldExmaxchShift))
}

const (
	RegisterDfsdmflt1exmaxFieldExmaxShift = 8
	RegisterDfsdmflt1exmaxFieldExmaxMask  = 0xffffff00
)

// GetExmax Extremes detector maximum value
func (r *registerDfsdmflt1exmaxType) GetExmax() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1exmaxFieldExmaxMask) >> RegisterDfsdmflt1exmaxFieldExmaxShift)
}

// SetExmax Extremes detector maximum value
func (r *registerDfsdmflt1exmaxType) SetExmax(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1exmaxFieldExmaxMask)|(uint32(value)<<RegisterDfsdmflt1exmaxFieldExmaxShift))
}

// registerDfsdmflt1exminType Extremes detector minimum register
type registerDfsdmflt1exminType uint32

const (
	RegisterDfsdmflt1exminFieldExminchShift = 0
	RegisterDfsdmflt1exminFieldExminchMask  = 0x7
)

// GetExminch Extremes detector minimum data channel
func (r *registerDfsdmflt1exminType) GetExminch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1exminFieldExminchMask) >> RegisterDfsdmflt1exminFieldExminchShift)
}

// SetExminch Extremes detector minimum data channel
func (r *registerDfsdmflt1exminType) SetExminch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1exminFieldExminchMask)|(uint32(value)<<RegisterDfsdmflt1exminFieldExminchShift))
}

const (
	RegisterDfsdmflt1exminFieldExminShift = 8
	RegisterDfsdmflt1exminFieldExminMask  = 0xffffff00
)

// GetExmin EXMIN
func (r *registerDfsdmflt1exminType) GetExmin() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1exminFieldExminMask) >> RegisterDfsdmflt1exminFieldExminShift)
}

// SetExmin EXMIN
func (r *registerDfsdmflt1exminType) SetExmin(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1exminFieldExminMask)|(uint32(value)<<RegisterDfsdmflt1exminFieldExminShift))
}

// registerDfsdmflt1cnvtimrType conversion timer register
type registerDfsdmflt1cnvtimrType uint32

const (
	RegisterDfsdmflt1cnvtimrFieldCnvcntShift = 4
	RegisterDfsdmflt1cnvtimrFieldCnvcntMask  = 0xfffffff0
)

// GetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdmflt1cnvtimrType) GetCnvcnt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt1cnvtimrFieldCnvcntMask) >> RegisterDfsdmflt1cnvtimrFieldCnvcntShift)
}

// SetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdmflt1cnvtimrType) SetCnvcnt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt1cnvtimrFieldCnvcntMask)|(uint32(value)<<RegisterDfsdmflt1cnvtimrFieldCnvcntShift))
}

// registerDfsdmflt2cr1Type control register 1
type registerDfsdmflt2cr1Type uint32

const (
	RegisterDfsdmflt2cr1FieldDfenShift = 0
	RegisterDfsdmflt2cr1FieldDfenMask  = 0x1
)

// GetDfen DFSDM enable
func (r *registerDfsdmflt2cr1Type) GetDfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldDfenMask) != 0
}

// SetDfen DFSDM enable
func (r *registerDfsdmflt2cr1Type) SetDfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr1FieldDfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldDfenMask)
	}
}

const (
	RegisterDfsdmflt2cr1FieldJswstartShift = 1
	RegisterDfsdmflt2cr1FieldJswstartMask  = 0x2
)

// GetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdmflt2cr1Type) GetJswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldJswstartMask) != 0
}

// SetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdmflt2cr1Type) SetJswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr1FieldJswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldJswstartMask)
	}
}

const (
	RegisterDfsdmflt2cr1FieldJsyncShift = 3
	RegisterDfsdmflt2cr1FieldJsyncMask  = 0x8
)

// GetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdmflt2cr1Type) GetJsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldJsyncMask) != 0
}

// SetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdmflt2cr1Type) SetJsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr1FieldJsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldJsyncMask)
	}
}

const (
	RegisterDfsdmflt2cr1FieldJscanShift = 4
	RegisterDfsdmflt2cr1FieldJscanMask  = 0x10
)

// GetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdmflt2cr1Type) GetJscan() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldJscanMask) != 0
}

// SetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdmflt2cr1Type) SetJscan(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr1FieldJscanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldJscanMask)
	}
}

const (
	RegisterDfsdmflt2cr1FieldJdmaenShift = 5
	RegisterDfsdmflt2cr1FieldJdmaenMask  = 0x20
)

// GetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdmflt2cr1Type) GetJdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldJdmaenMask) != 0
}

// SetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdmflt2cr1Type) SetJdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr1FieldJdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldJdmaenMask)
	}
}

const (
	RegisterDfsdmflt2cr1FieldJextselShift = 8
	RegisterDfsdmflt2cr1FieldJextselMask  = 0x700
)

// GetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdmflt2cr1Type) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldJextselMask) >> RegisterDfsdmflt2cr1FieldJextselShift)
}

// SetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdmflt2cr1Type) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldJextselMask)|(uint32(value)<<RegisterDfsdmflt2cr1FieldJextselShift))
}

const (
	RegisterDfsdmflt2cr1FieldJextenShift = 13
	RegisterDfsdmflt2cr1FieldJextenMask  = 0x6000
)

// GetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdmflt2cr1Type) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldJextenMask) >> RegisterDfsdmflt2cr1FieldJextenShift)
}

// SetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdmflt2cr1Type) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldJextenMask)|(uint32(value)<<RegisterDfsdmflt2cr1FieldJextenShift))
}

const (
	RegisterDfsdmflt2cr1FieldRswstartShift = 17
	RegisterDfsdmflt2cr1FieldRswstartMask  = 0x20000
)

// GetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdmflt2cr1Type) GetRswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldRswstartMask) != 0
}

// SetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdmflt2cr1Type) SetRswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr1FieldRswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldRswstartMask)
	}
}

const (
	RegisterDfsdmflt2cr1FieldRcontShift = 18
	RegisterDfsdmflt2cr1FieldRcontMask  = 0x40000
)

// GetRcont Continuous mode selection for regular conversions
func (r *registerDfsdmflt2cr1Type) GetRcont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldRcontMask) != 0
}

// SetRcont Continuous mode selection for regular conversions
func (r *registerDfsdmflt2cr1Type) SetRcont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr1FieldRcontMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldRcontMask)
	}
}

const (
	RegisterDfsdmflt2cr1FieldRsyncShift = 19
	RegisterDfsdmflt2cr1FieldRsyncMask  = 0x80000
)

// GetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdmflt2cr1Type) GetRsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldRsyncMask) != 0
}

// SetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdmflt2cr1Type) SetRsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr1FieldRsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldRsyncMask)
	}
}

const (
	RegisterDfsdmflt2cr1FieldRdmaenShift = 21
	RegisterDfsdmflt2cr1FieldRdmaenMask  = 0x200000
)

// GetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdmflt2cr1Type) GetRdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldRdmaenMask) != 0
}

// SetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdmflt2cr1Type) SetRdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr1FieldRdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldRdmaenMask)
	}
}

const (
	RegisterDfsdmflt2cr1FieldRchShift = 24
	RegisterDfsdmflt2cr1FieldRchMask  = 0x7000000
)

// GetRch Regular channel selection
func (r *registerDfsdmflt2cr1Type) GetRch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldRchMask) >> RegisterDfsdmflt2cr1FieldRchShift)
}

// SetRch Regular channel selection
func (r *registerDfsdmflt2cr1Type) SetRch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldRchMask)|(uint32(value)<<RegisterDfsdmflt2cr1FieldRchShift))
}

const (
	RegisterDfsdmflt2cr1FieldFastShift = 29
	RegisterDfsdmflt2cr1FieldFastMask  = 0x20000000
)

// GetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdmflt2cr1Type) GetFast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldFastMask) != 0
}

// SetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdmflt2cr1Type) SetFast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr1FieldFastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldFastMask)
	}
}

const (
	RegisterDfsdmflt2cr1FieldAwfselShift = 30
	RegisterDfsdmflt2cr1FieldAwfselMask  = 0x40000000
)

// GetAwfsel Analog watchdog fast mode select
func (r *registerDfsdmflt2cr1Type) GetAwfsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr1FieldAwfselMask) != 0
}

// SetAwfsel Analog watchdog fast mode select
func (r *registerDfsdmflt2cr1Type) SetAwfsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr1FieldAwfselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr1FieldAwfselMask)
	}
}

// registerDfsdmflt2cr2Type control register 2
type registerDfsdmflt2cr2Type uint32

const (
	RegisterDfsdmflt2cr2FieldJeocieShift = 0
	RegisterDfsdmflt2cr2FieldJeocieMask  = 0x1
)

// GetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdmflt2cr2Type) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr2FieldJeocieMask) != 0
}

// SetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdmflt2cr2Type) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr2FieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr2FieldJeocieMask)
	}
}

const (
	RegisterDfsdmflt2cr2FieldReocieShift = 1
	RegisterDfsdmflt2cr2FieldReocieMask  = 0x2
)

// GetReocie Regular end of conversion interrupt enable
func (r *registerDfsdmflt2cr2Type) GetReocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr2FieldReocieMask) != 0
}

// SetReocie Regular end of conversion interrupt enable
func (r *registerDfsdmflt2cr2Type) SetReocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr2FieldReocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr2FieldReocieMask)
	}
}

const (
	RegisterDfsdmflt2cr2FieldJovrieShift = 2
	RegisterDfsdmflt2cr2FieldJovrieMask  = 0x4
)

// GetJovrie Injected data overrun interrupt enable
func (r *registerDfsdmflt2cr2Type) GetJovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr2FieldJovrieMask) != 0
}

// SetJovrie Injected data overrun interrupt enable
func (r *registerDfsdmflt2cr2Type) SetJovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr2FieldJovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr2FieldJovrieMask)
	}
}

const (
	RegisterDfsdmflt2cr2FieldRovrieShift = 3
	RegisterDfsdmflt2cr2FieldRovrieMask  = 0x8
)

// GetRovrie Regular data overrun interrupt enable
func (r *registerDfsdmflt2cr2Type) GetRovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr2FieldRovrieMask) != 0
}

// SetRovrie Regular data overrun interrupt enable
func (r *registerDfsdmflt2cr2Type) SetRovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr2FieldRovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr2FieldRovrieMask)
	}
}

const (
	RegisterDfsdmflt2cr2FieldAwdieShift = 4
	RegisterDfsdmflt2cr2FieldAwdieMask  = 0x10
)

// GetAwdie Analog watchdog interrupt enable
func (r *registerDfsdmflt2cr2Type) GetAwdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr2FieldAwdieMask) != 0
}

// SetAwdie Analog watchdog interrupt enable
func (r *registerDfsdmflt2cr2Type) SetAwdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr2FieldAwdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr2FieldAwdieMask)
	}
}

const (
	RegisterDfsdmflt2cr2FieldScdieShift = 5
	RegisterDfsdmflt2cr2FieldScdieMask  = 0x20
)

// GetScdie Short-circuit detector interrupt enable
func (r *registerDfsdmflt2cr2Type) GetScdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr2FieldScdieMask) != 0
}

// SetScdie Short-circuit detector interrupt enable
func (r *registerDfsdmflt2cr2Type) SetScdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr2FieldScdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr2FieldScdieMask)
	}
}

const (
	RegisterDfsdmflt2cr2FieldCkabieShift = 6
	RegisterDfsdmflt2cr2FieldCkabieMask  = 0x40
)

// GetCkabie Clock absence interrupt enable
func (r *registerDfsdmflt2cr2Type) GetCkabie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr2FieldCkabieMask) != 0
}

// SetCkabie Clock absence interrupt enable
func (r *registerDfsdmflt2cr2Type) SetCkabie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2cr2FieldCkabieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr2FieldCkabieMask)
	}
}

const (
	RegisterDfsdmflt2cr2FieldExchShift = 8
	RegisterDfsdmflt2cr2FieldExchMask  = 0xff00
)

// GetExch Extremes detector channel selection
func (r *registerDfsdmflt2cr2Type) GetExch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr2FieldExchMask) >> RegisterDfsdmflt2cr2FieldExchShift)
}

// SetExch Extremes detector channel selection
func (r *registerDfsdmflt2cr2Type) SetExch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr2FieldExchMask)|(uint32(value)<<RegisterDfsdmflt2cr2FieldExchShift))
}

const (
	RegisterDfsdmflt2cr2FieldAwdchShift = 16
	RegisterDfsdmflt2cr2FieldAwdchMask  = 0xff0000
)

// GetAwdch Analog watchdog channel selection
func (r *registerDfsdmflt2cr2Type) GetAwdch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cr2FieldAwdchMask) >> RegisterDfsdmflt2cr2FieldAwdchShift)
}

// SetAwdch Analog watchdog channel selection
func (r *registerDfsdmflt2cr2Type) SetAwdch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cr2FieldAwdchMask)|(uint32(value)<<RegisterDfsdmflt2cr2FieldAwdchShift))
}

// registerDfsdmflt2isrType interrupt and status register
type registerDfsdmflt2isrType uint32

const (
	RegisterDfsdmflt2isrFieldJeocfShift = 0
	RegisterDfsdmflt2isrFieldJeocfMask  = 0x1
)

// GetJeocf End of injected conversion flag
func (r *registerDfsdmflt2isrType) GetJeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2isrFieldJeocfMask) != 0
}

// SetJeocf End of injected conversion flag
func (r *registerDfsdmflt2isrType) SetJeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2isrFieldJeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2isrFieldJeocfMask)
	}
}

const (
	RegisterDfsdmflt2isrFieldReocfShift = 1
	RegisterDfsdmflt2isrFieldReocfMask  = 0x2
)

// GetReocf End of regular conversion flag
func (r *registerDfsdmflt2isrType) GetReocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2isrFieldReocfMask) != 0
}

// SetReocf End of regular conversion flag
func (r *registerDfsdmflt2isrType) SetReocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2isrFieldReocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2isrFieldReocfMask)
	}
}

const (
	RegisterDfsdmflt2isrFieldJovrfShift = 2
	RegisterDfsdmflt2isrFieldJovrfMask  = 0x4
)

// GetJovrf Injected conversion overrun flag
func (r *registerDfsdmflt2isrType) GetJovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2isrFieldJovrfMask) != 0
}

// SetJovrf Injected conversion overrun flag
func (r *registerDfsdmflt2isrType) SetJovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2isrFieldJovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2isrFieldJovrfMask)
	}
}

const (
	RegisterDfsdmflt2isrFieldRovrfShift = 3
	RegisterDfsdmflt2isrFieldRovrfMask  = 0x8
)

// GetRovrf Regular conversion overrun flag
func (r *registerDfsdmflt2isrType) GetRovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2isrFieldRovrfMask) != 0
}

// SetRovrf Regular conversion overrun flag
func (r *registerDfsdmflt2isrType) SetRovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2isrFieldRovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2isrFieldRovrfMask)
	}
}

const (
	RegisterDfsdmflt2isrFieldAwdfShift = 4
	RegisterDfsdmflt2isrFieldAwdfMask  = 0x10
)

// GetAwdf Analog watchdog
func (r *registerDfsdmflt2isrType) GetAwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2isrFieldAwdfMask) != 0
}

// SetAwdf Analog watchdog
func (r *registerDfsdmflt2isrType) SetAwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2isrFieldAwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2isrFieldAwdfMask)
	}
}

const (
	RegisterDfsdmflt2isrFieldJcipShift = 13
	RegisterDfsdmflt2isrFieldJcipMask  = 0x2000
)

// GetJcip Injected conversion in progress status
func (r *registerDfsdmflt2isrType) GetJcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2isrFieldJcipMask) != 0
}

// SetJcip Injected conversion in progress status
func (r *registerDfsdmflt2isrType) SetJcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2isrFieldJcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2isrFieldJcipMask)
	}
}

const (
	RegisterDfsdmflt2isrFieldRcipShift = 14
	RegisterDfsdmflt2isrFieldRcipMask  = 0x4000
)

// GetRcip Regular conversion in progress status
func (r *registerDfsdmflt2isrType) GetRcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2isrFieldRcipMask) != 0
}

// SetRcip Regular conversion in progress status
func (r *registerDfsdmflt2isrType) SetRcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2isrFieldRcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2isrFieldRcipMask)
	}
}

const (
	RegisterDfsdmflt2isrFieldCkabfShift = 16
	RegisterDfsdmflt2isrFieldCkabfMask  = 0xff0000
)

// GetCkabf Clock absence flag
func (r *registerDfsdmflt2isrType) GetCkabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2isrFieldCkabfMask) >> RegisterDfsdmflt2isrFieldCkabfShift)
}

// SetCkabf Clock absence flag
func (r *registerDfsdmflt2isrType) SetCkabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2isrFieldCkabfMask)|(uint32(value)<<RegisterDfsdmflt2isrFieldCkabfShift))
}

const (
	RegisterDfsdmflt2isrFieldScdfShift = 24
	RegisterDfsdmflt2isrFieldScdfMask  = 0xff000000
)

// GetScdf short-circuit detector flag
func (r *registerDfsdmflt2isrType) GetScdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2isrFieldScdfMask) >> RegisterDfsdmflt2isrFieldScdfShift)
}

// SetScdf short-circuit detector flag
func (r *registerDfsdmflt2isrType) SetScdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2isrFieldScdfMask)|(uint32(value)<<RegisterDfsdmflt2isrFieldScdfShift))
}

// registerDfsdmflt2icrType interrupt flag clear register
type registerDfsdmflt2icrType uint32

const (
	RegisterDfsdmflt2icrFieldClrjovrfShift = 2
	RegisterDfsdmflt2icrFieldClrjovrfMask  = 0x4
)

// GetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdmflt2icrType) GetClrjovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2icrFieldClrjovrfMask) != 0
}

// SetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdmflt2icrType) SetClrjovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2icrFieldClrjovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2icrFieldClrjovrfMask)
	}
}

const (
	RegisterDfsdmflt2icrFieldClrrovrfShift = 3
	RegisterDfsdmflt2icrFieldClrrovrfMask  = 0x8
)

// GetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdmflt2icrType) GetClrrovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2icrFieldClrrovrfMask) != 0
}

// SetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdmflt2icrType) SetClrrovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2icrFieldClrrovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2icrFieldClrrovrfMask)
	}
}

const (
	RegisterDfsdmflt2icrFieldClrckabfShift = 16
	RegisterDfsdmflt2icrFieldClrckabfMask  = 0xff0000
)

// GetClrckabf Clear the clock absence flag
func (r *registerDfsdmflt2icrType) GetClrckabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2icrFieldClrckabfMask) >> RegisterDfsdmflt2icrFieldClrckabfShift)
}

// SetClrckabf Clear the clock absence flag
func (r *registerDfsdmflt2icrType) SetClrckabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2icrFieldClrckabfMask)|(uint32(value)<<RegisterDfsdmflt2icrFieldClrckabfShift))
}

const (
	RegisterDfsdmflt2icrFieldClrscdfShift = 24
	RegisterDfsdmflt2icrFieldClrscdfMask  = 0xff000000
)

// GetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdmflt2icrType) GetClrscdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2icrFieldClrscdfMask) >> RegisterDfsdmflt2icrFieldClrscdfShift)
}

// SetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdmflt2icrType) SetClrscdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2icrFieldClrscdfMask)|(uint32(value)<<RegisterDfsdmflt2icrFieldClrscdfShift))
}

// registerDfsdmflt2jchgrType injected channel group selection register
type registerDfsdmflt2jchgrType uint32

const (
	RegisterDfsdmflt2jchgrFieldJchgShift = 0
	RegisterDfsdmflt2jchgrFieldJchgMask  = 0xff
)

// GetJchg Injected channel group selection
func (r *registerDfsdmflt2jchgrType) GetJchg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2jchgrFieldJchgMask) >> RegisterDfsdmflt2jchgrFieldJchgShift)
}

// SetJchg Injected channel group selection
func (r *registerDfsdmflt2jchgrType) SetJchg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2jchgrFieldJchgMask)|(uint32(value)<<RegisterDfsdmflt2jchgrFieldJchgShift))
}

// registerDfsdmflt2fcrType filter control register
type registerDfsdmflt2fcrType uint32

const (
	RegisterDfsdmflt2fcrFieldIosrShift = 0
	RegisterDfsdmflt2fcrFieldIosrMask  = 0xff
)

// GetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdmflt2fcrType) GetIosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2fcrFieldIosrMask) >> RegisterDfsdmflt2fcrFieldIosrShift)
}

// SetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdmflt2fcrType) SetIosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2fcrFieldIosrMask)|(uint32(value)<<RegisterDfsdmflt2fcrFieldIosrShift))
}

const (
	RegisterDfsdmflt2fcrFieldFosrShift = 16
	RegisterDfsdmflt2fcrFieldFosrMask  = 0x3ff0000
)

// GetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdmflt2fcrType) GetFosr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2fcrFieldFosrMask) >> RegisterDfsdmflt2fcrFieldFosrShift)
}

// SetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdmflt2fcrType) SetFosr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2fcrFieldFosrMask)|(uint32(value)<<RegisterDfsdmflt2fcrFieldFosrShift))
}

const (
	RegisterDfsdmflt2fcrFieldFordShift = 29
	RegisterDfsdmflt2fcrFieldFordMask  = 0xe0000000
)

// GetFord Sinc filter order
func (r *registerDfsdmflt2fcrType) GetFord() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2fcrFieldFordMask) >> RegisterDfsdmflt2fcrFieldFordShift)
}

// SetFord Sinc filter order
func (r *registerDfsdmflt2fcrType) SetFord(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2fcrFieldFordMask)|(uint32(value)<<RegisterDfsdmflt2fcrFieldFordShift))
}

// registerDfsdmflt2jdatarType data register for injected group
type registerDfsdmflt2jdatarType uint32

const (
	RegisterDfsdmflt2jdatarFieldJdatachShift = 0
	RegisterDfsdmflt2jdatarFieldJdatachMask  = 0x7
)

// GetJdatach Injected channel most recently converted
func (r *registerDfsdmflt2jdatarType) GetJdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2jdatarFieldJdatachMask) >> RegisterDfsdmflt2jdatarFieldJdatachShift)
}

// SetJdatach Injected channel most recently converted
func (r *registerDfsdmflt2jdatarType) SetJdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2jdatarFieldJdatachMask)|(uint32(value)<<RegisterDfsdmflt2jdatarFieldJdatachShift))
}

const (
	RegisterDfsdmflt2jdatarFieldJdataShift = 8
	RegisterDfsdmflt2jdatarFieldJdataMask  = 0xffffff00
)

// GetJdata Injected group conversion data
func (r *registerDfsdmflt2jdatarType) GetJdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2jdatarFieldJdataMask) >> RegisterDfsdmflt2jdatarFieldJdataShift)
}

// SetJdata Injected group conversion data
func (r *registerDfsdmflt2jdatarType) SetJdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2jdatarFieldJdataMask)|(uint32(value)<<RegisterDfsdmflt2jdatarFieldJdataShift))
}

// registerDfsdmflt2rdatarType data register for the regular channel
type registerDfsdmflt2rdatarType uint32

const (
	RegisterDfsdmflt2rdatarFieldRdatachShift = 0
	RegisterDfsdmflt2rdatarFieldRdatachMask  = 0x7
)

// GetRdatach Regular channel most recently converted
func (r *registerDfsdmflt2rdatarType) GetRdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2rdatarFieldRdatachMask) >> RegisterDfsdmflt2rdatarFieldRdatachShift)
}

// SetRdatach Regular channel most recently converted
func (r *registerDfsdmflt2rdatarType) SetRdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2rdatarFieldRdatachMask)|(uint32(value)<<RegisterDfsdmflt2rdatarFieldRdatachShift))
}

const (
	RegisterDfsdmflt2rdatarFieldRpendShift = 4
	RegisterDfsdmflt2rdatarFieldRpendMask  = 0x10
)

// GetRpend Regular channel pending data
func (r *registerDfsdmflt2rdatarType) GetRpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2rdatarFieldRpendMask) != 0
}

// SetRpend Regular channel pending data
func (r *registerDfsdmflt2rdatarType) SetRpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt2rdatarFieldRpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2rdatarFieldRpendMask)
	}
}

const (
	RegisterDfsdmflt2rdatarFieldRdataShift = 8
	RegisterDfsdmflt2rdatarFieldRdataMask  = 0xffffff00
)

// GetRdata Regular channel conversion data
func (r *registerDfsdmflt2rdatarType) GetRdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2rdatarFieldRdataMask) >> RegisterDfsdmflt2rdatarFieldRdataShift)
}

// SetRdata Regular channel conversion data
func (r *registerDfsdmflt2rdatarType) SetRdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2rdatarFieldRdataMask)|(uint32(value)<<RegisterDfsdmflt2rdatarFieldRdataShift))
}

// registerDfsdmflt2awhtrType analog watchdog high threshold register
type registerDfsdmflt2awhtrType uint32

const (
	RegisterDfsdmflt2awhtrFieldBkawhShift = 0
	RegisterDfsdmflt2awhtrFieldBkawhMask  = 0xf
)

// GetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdmflt2awhtrType) GetBkawh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2awhtrFieldBkawhMask) >> RegisterDfsdmflt2awhtrFieldBkawhShift)
}

// SetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdmflt2awhtrType) SetBkawh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2awhtrFieldBkawhMask)|(uint32(value)<<RegisterDfsdmflt2awhtrFieldBkawhShift))
}

const (
	RegisterDfsdmflt2awhtrFieldAwhtShift = 8
	RegisterDfsdmflt2awhtrFieldAwhtMask  = 0xffffff00
)

// GetAwht Analog watchdog high threshold
func (r *registerDfsdmflt2awhtrType) GetAwht() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2awhtrFieldAwhtMask) >> RegisterDfsdmflt2awhtrFieldAwhtShift)
}

// SetAwht Analog watchdog high threshold
func (r *registerDfsdmflt2awhtrType) SetAwht(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2awhtrFieldAwhtMask)|(uint32(value)<<RegisterDfsdmflt2awhtrFieldAwhtShift))
}

// registerDfsdmflt2awltrType analog watchdog low threshold register
type registerDfsdmflt2awltrType uint32

const (
	RegisterDfsdmflt2awltrFieldBkawlShift = 0
	RegisterDfsdmflt2awltrFieldBkawlMask  = 0xf
)

// GetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdmflt2awltrType) GetBkawl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2awltrFieldBkawlMask) >> RegisterDfsdmflt2awltrFieldBkawlShift)
}

// SetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdmflt2awltrType) SetBkawl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2awltrFieldBkawlMask)|(uint32(value)<<RegisterDfsdmflt2awltrFieldBkawlShift))
}

const (
	RegisterDfsdmflt2awltrFieldAwltShift = 8
	RegisterDfsdmflt2awltrFieldAwltMask  = 0xffffff00
)

// GetAwlt Analog watchdog low threshold
func (r *registerDfsdmflt2awltrType) GetAwlt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2awltrFieldAwltMask) >> RegisterDfsdmflt2awltrFieldAwltShift)
}

// SetAwlt Analog watchdog low threshold
func (r *registerDfsdmflt2awltrType) SetAwlt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2awltrFieldAwltMask)|(uint32(value)<<RegisterDfsdmflt2awltrFieldAwltShift))
}

// registerDfsdmflt2awsrType analog watchdog status register
type registerDfsdmflt2awsrType uint32

const (
	RegisterDfsdmflt2awsrFieldAwltfShift = 0
	RegisterDfsdmflt2awsrFieldAwltfMask  = 0xff
)

// GetAwltf Analog watchdog low threshold flag
func (r *registerDfsdmflt2awsrType) GetAwltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2awsrFieldAwltfMask) >> RegisterDfsdmflt2awsrFieldAwltfShift)
}

// SetAwltf Analog watchdog low threshold flag
func (r *registerDfsdmflt2awsrType) SetAwltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2awsrFieldAwltfMask)|(uint32(value)<<RegisterDfsdmflt2awsrFieldAwltfShift))
}

const (
	RegisterDfsdmflt2awsrFieldAwhtfShift = 8
	RegisterDfsdmflt2awsrFieldAwhtfMask  = 0xff00
)

// GetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdmflt2awsrType) GetAwhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2awsrFieldAwhtfMask) >> RegisterDfsdmflt2awsrFieldAwhtfShift)
}

// SetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdmflt2awsrType) SetAwhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2awsrFieldAwhtfMask)|(uint32(value)<<RegisterDfsdmflt2awsrFieldAwhtfShift))
}

// registerDfsdmflt2awcfrType analog watchdog clear flag register
type registerDfsdmflt2awcfrType uint32

const (
	RegisterDfsdmflt2awcfrFieldClrawltfShift = 0
	RegisterDfsdmflt2awcfrFieldClrawltfMask  = 0xff
)

// GetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdmflt2awcfrType) GetClrawltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2awcfrFieldClrawltfMask) >> RegisterDfsdmflt2awcfrFieldClrawltfShift)
}

// SetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdmflt2awcfrType) SetClrawltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2awcfrFieldClrawltfMask)|(uint32(value)<<RegisterDfsdmflt2awcfrFieldClrawltfShift))
}

const (
	RegisterDfsdmflt2awcfrFieldClrawhtfShift = 8
	RegisterDfsdmflt2awcfrFieldClrawhtfMask  = 0xff00
)

// GetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdmflt2awcfrType) GetClrawhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2awcfrFieldClrawhtfMask) >> RegisterDfsdmflt2awcfrFieldClrawhtfShift)
}

// SetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdmflt2awcfrType) SetClrawhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2awcfrFieldClrawhtfMask)|(uint32(value)<<RegisterDfsdmflt2awcfrFieldClrawhtfShift))
}

// registerDfsdmflt2exmaxType Extremes detector maximum register
type registerDfsdmflt2exmaxType uint32

const (
	RegisterDfsdmflt2exmaxFieldExmaxchShift = 0
	RegisterDfsdmflt2exmaxFieldExmaxchMask  = 0x7
)

// GetExmaxch Extremes detector maximum data channel
func (r *registerDfsdmflt2exmaxType) GetExmaxch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2exmaxFieldExmaxchMask) >> RegisterDfsdmflt2exmaxFieldExmaxchShift)
}

// SetExmaxch Extremes detector maximum data channel
func (r *registerDfsdmflt2exmaxType) SetExmaxch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2exmaxFieldExmaxchMask)|(uint32(value)<<RegisterDfsdmflt2exmaxFieldExmaxchShift))
}

const (
	RegisterDfsdmflt2exmaxFieldExmaxShift = 8
	RegisterDfsdmflt2exmaxFieldExmaxMask  = 0xffffff00
)

// GetExmax Extremes detector maximum value
func (r *registerDfsdmflt2exmaxType) GetExmax() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2exmaxFieldExmaxMask) >> RegisterDfsdmflt2exmaxFieldExmaxShift)
}

// SetExmax Extremes detector maximum value
func (r *registerDfsdmflt2exmaxType) SetExmax(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2exmaxFieldExmaxMask)|(uint32(value)<<RegisterDfsdmflt2exmaxFieldExmaxShift))
}

// registerDfsdmflt2exminType Extremes detector minimum register
type registerDfsdmflt2exminType uint32

const (
	RegisterDfsdmflt2exminFieldExminchShift = 0
	RegisterDfsdmflt2exminFieldExminchMask  = 0x7
)

// GetExminch Extremes detector minimum data channel
func (r *registerDfsdmflt2exminType) GetExminch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2exminFieldExminchMask) >> RegisterDfsdmflt2exminFieldExminchShift)
}

// SetExminch Extremes detector minimum data channel
func (r *registerDfsdmflt2exminType) SetExminch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2exminFieldExminchMask)|(uint32(value)<<RegisterDfsdmflt2exminFieldExminchShift))
}

const (
	RegisterDfsdmflt2exminFieldExminShift = 8
	RegisterDfsdmflt2exminFieldExminMask  = 0xffffff00
)

// GetExmin EXMIN
func (r *registerDfsdmflt2exminType) GetExmin() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2exminFieldExminMask) >> RegisterDfsdmflt2exminFieldExminShift)
}

// SetExmin EXMIN
func (r *registerDfsdmflt2exminType) SetExmin(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2exminFieldExminMask)|(uint32(value)<<RegisterDfsdmflt2exminFieldExminShift))
}

// registerDfsdmflt2cnvtimrType conversion timer register
type registerDfsdmflt2cnvtimrType uint32

const (
	RegisterDfsdmflt2cnvtimrFieldCnvcntShift = 4
	RegisterDfsdmflt2cnvtimrFieldCnvcntMask  = 0xfffffff0
)

// GetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdmflt2cnvtimrType) GetCnvcnt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt2cnvtimrFieldCnvcntMask) >> RegisterDfsdmflt2cnvtimrFieldCnvcntShift)
}

// SetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdmflt2cnvtimrType) SetCnvcnt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt2cnvtimrFieldCnvcntMask)|(uint32(value)<<RegisterDfsdmflt2cnvtimrFieldCnvcntShift))
}

// registerDfsdmflt3cr1Type control register 1
type registerDfsdmflt3cr1Type uint32

const (
	RegisterDfsdmflt3cr1FieldDfenShift = 0
	RegisterDfsdmflt3cr1FieldDfenMask  = 0x1
)

// GetDfen DFSDM enable
func (r *registerDfsdmflt3cr1Type) GetDfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldDfenMask) != 0
}

// SetDfen DFSDM enable
func (r *registerDfsdmflt3cr1Type) SetDfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr1FieldDfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldDfenMask)
	}
}

const (
	RegisterDfsdmflt3cr1FieldJswstartShift = 1
	RegisterDfsdmflt3cr1FieldJswstartMask  = 0x2
)

// GetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdmflt3cr1Type) GetJswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldJswstartMask) != 0
}

// SetJswstart Start a conversion of the injected group of channels
func (r *registerDfsdmflt3cr1Type) SetJswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr1FieldJswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldJswstartMask)
	}
}

const (
	RegisterDfsdmflt3cr1FieldJsyncShift = 3
	RegisterDfsdmflt3cr1FieldJsyncMask  = 0x8
)

// GetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdmflt3cr1Type) GetJsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldJsyncMask) != 0
}

// SetJsync Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
func (r *registerDfsdmflt3cr1Type) SetJsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr1FieldJsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldJsyncMask)
	}
}

const (
	RegisterDfsdmflt3cr1FieldJscanShift = 4
	RegisterDfsdmflt3cr1FieldJscanMask  = 0x10
)

// GetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdmflt3cr1Type) GetJscan() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldJscanMask) != 0
}

// SetJscan Scanning conversion mode for injected conversions
func (r *registerDfsdmflt3cr1Type) SetJscan(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr1FieldJscanMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldJscanMask)
	}
}

const (
	RegisterDfsdmflt3cr1FieldJdmaenShift = 5
	RegisterDfsdmflt3cr1FieldJdmaenMask  = 0x20
)

// GetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdmflt3cr1Type) GetJdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldJdmaenMask) != 0
}

// SetJdmaen DMA channel enabled to read data for the injected channel group
func (r *registerDfsdmflt3cr1Type) SetJdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr1FieldJdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldJdmaenMask)
	}
}

const (
	RegisterDfsdmflt3cr1FieldJextselShift = 8
	RegisterDfsdmflt3cr1FieldJextselMask  = 0x700
)

// GetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdmflt3cr1Type) GetJextsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldJextselMask) >> RegisterDfsdmflt3cr1FieldJextselShift)
}

// SetJextsel Trigger signal selection for launching injected conversions
func (r *registerDfsdmflt3cr1Type) SetJextsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldJextselMask)|(uint32(value)<<RegisterDfsdmflt3cr1FieldJextselShift))
}

const (
	RegisterDfsdmflt3cr1FieldJextenShift = 13
	RegisterDfsdmflt3cr1FieldJextenMask  = 0x6000
)

// GetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdmflt3cr1Type) GetJexten() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldJextenMask) >> RegisterDfsdmflt3cr1FieldJextenShift)
}

// SetJexten Trigger enable and trigger edge selection for injected conversions
func (r *registerDfsdmflt3cr1Type) SetJexten(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldJextenMask)|(uint32(value)<<RegisterDfsdmflt3cr1FieldJextenShift))
}

const (
	RegisterDfsdmflt3cr1FieldRswstartShift = 17
	RegisterDfsdmflt3cr1FieldRswstartMask  = 0x20000
)

// GetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdmflt3cr1Type) GetRswstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldRswstartMask) != 0
}

// SetRswstart Software start of a conversion on the regular channel
func (r *registerDfsdmflt3cr1Type) SetRswstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr1FieldRswstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldRswstartMask)
	}
}

const (
	RegisterDfsdmflt3cr1FieldRcontShift = 18
	RegisterDfsdmflt3cr1FieldRcontMask  = 0x40000
)

// GetRcont Continuous mode selection for regular conversions
func (r *registerDfsdmflt3cr1Type) GetRcont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldRcontMask) != 0
}

// SetRcont Continuous mode selection for regular conversions
func (r *registerDfsdmflt3cr1Type) SetRcont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr1FieldRcontMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldRcontMask)
	}
}

const (
	RegisterDfsdmflt3cr1FieldRsyncShift = 19
	RegisterDfsdmflt3cr1FieldRsyncMask  = 0x80000
)

// GetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdmflt3cr1Type) GetRsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldRsyncMask) != 0
}

// SetRsync Launch regular conversion synchronously with DFSDM0
func (r *registerDfsdmflt3cr1Type) SetRsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr1FieldRsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldRsyncMask)
	}
}

const (
	RegisterDfsdmflt3cr1FieldRdmaenShift = 21
	RegisterDfsdmflt3cr1FieldRdmaenMask  = 0x200000
)

// GetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdmflt3cr1Type) GetRdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldRdmaenMask) != 0
}

// SetRdmaen DMA channel enabled to read data for the regular conversion
func (r *registerDfsdmflt3cr1Type) SetRdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr1FieldRdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldRdmaenMask)
	}
}

const (
	RegisterDfsdmflt3cr1FieldRchShift = 24
	RegisterDfsdmflt3cr1FieldRchMask  = 0x7000000
)

// GetRch Regular channel selection
func (r *registerDfsdmflt3cr1Type) GetRch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldRchMask) >> RegisterDfsdmflt3cr1FieldRchShift)
}

// SetRch Regular channel selection
func (r *registerDfsdmflt3cr1Type) SetRch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldRchMask)|(uint32(value)<<RegisterDfsdmflt3cr1FieldRchShift))
}

const (
	RegisterDfsdmflt3cr1FieldFastShift = 29
	RegisterDfsdmflt3cr1FieldFastMask  = 0x20000000
)

// GetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdmflt3cr1Type) GetFast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldFastMask) != 0
}

// SetFast Fast conversion mode selection for regular conversions
func (r *registerDfsdmflt3cr1Type) SetFast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr1FieldFastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldFastMask)
	}
}

const (
	RegisterDfsdmflt3cr1FieldAwfselShift = 30
	RegisterDfsdmflt3cr1FieldAwfselMask  = 0x40000000
)

// GetAwfsel Analog watchdog fast mode select
func (r *registerDfsdmflt3cr1Type) GetAwfsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr1FieldAwfselMask) != 0
}

// SetAwfsel Analog watchdog fast mode select
func (r *registerDfsdmflt3cr1Type) SetAwfsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr1FieldAwfselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr1FieldAwfselMask)
	}
}

// registerDfsdmflt3cr2Type control register 2
type registerDfsdmflt3cr2Type uint32

const (
	RegisterDfsdmflt3cr2FieldJeocieShift = 0
	RegisterDfsdmflt3cr2FieldJeocieMask  = 0x1
)

// GetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdmflt3cr2Type) GetJeocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr2FieldJeocieMask) != 0
}

// SetJeocie Injected end of conversion interrupt enable
func (r *registerDfsdmflt3cr2Type) SetJeocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr2FieldJeocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr2FieldJeocieMask)
	}
}

const (
	RegisterDfsdmflt3cr2FieldReocieShift = 1
	RegisterDfsdmflt3cr2FieldReocieMask  = 0x2
)

// GetReocie Regular end of conversion interrupt enable
func (r *registerDfsdmflt3cr2Type) GetReocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr2FieldReocieMask) != 0
}

// SetReocie Regular end of conversion interrupt enable
func (r *registerDfsdmflt3cr2Type) SetReocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr2FieldReocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr2FieldReocieMask)
	}
}

const (
	RegisterDfsdmflt3cr2FieldJovrieShift = 2
	RegisterDfsdmflt3cr2FieldJovrieMask  = 0x4
)

// GetJovrie Injected data overrun interrupt enable
func (r *registerDfsdmflt3cr2Type) GetJovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr2FieldJovrieMask) != 0
}

// SetJovrie Injected data overrun interrupt enable
func (r *registerDfsdmflt3cr2Type) SetJovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr2FieldJovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr2FieldJovrieMask)
	}
}

const (
	RegisterDfsdmflt3cr2FieldRovrieShift = 3
	RegisterDfsdmflt3cr2FieldRovrieMask  = 0x8
)

// GetRovrie Regular data overrun interrupt enable
func (r *registerDfsdmflt3cr2Type) GetRovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr2FieldRovrieMask) != 0
}

// SetRovrie Regular data overrun interrupt enable
func (r *registerDfsdmflt3cr2Type) SetRovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr2FieldRovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr2FieldRovrieMask)
	}
}

const (
	RegisterDfsdmflt3cr2FieldAwdieShift = 4
	RegisterDfsdmflt3cr2FieldAwdieMask  = 0x10
)

// GetAwdie Analog watchdog interrupt enable
func (r *registerDfsdmflt3cr2Type) GetAwdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr2FieldAwdieMask) != 0
}

// SetAwdie Analog watchdog interrupt enable
func (r *registerDfsdmflt3cr2Type) SetAwdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr2FieldAwdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr2FieldAwdieMask)
	}
}

const (
	RegisterDfsdmflt3cr2FieldScdieShift = 5
	RegisterDfsdmflt3cr2FieldScdieMask  = 0x20
)

// GetScdie Short-circuit detector interrupt enable
func (r *registerDfsdmflt3cr2Type) GetScdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr2FieldScdieMask) != 0
}

// SetScdie Short-circuit detector interrupt enable
func (r *registerDfsdmflt3cr2Type) SetScdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr2FieldScdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr2FieldScdieMask)
	}
}

const (
	RegisterDfsdmflt3cr2FieldCkabieShift = 6
	RegisterDfsdmflt3cr2FieldCkabieMask  = 0x40
)

// GetCkabie Clock absence interrupt enable
func (r *registerDfsdmflt3cr2Type) GetCkabie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr2FieldCkabieMask) != 0
}

// SetCkabie Clock absence interrupt enable
func (r *registerDfsdmflt3cr2Type) SetCkabie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3cr2FieldCkabieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr2FieldCkabieMask)
	}
}

const (
	RegisterDfsdmflt3cr2FieldExchShift = 8
	RegisterDfsdmflt3cr2FieldExchMask  = 0xff00
)

// GetExch Extremes detector channel selection
func (r *registerDfsdmflt3cr2Type) GetExch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr2FieldExchMask) >> RegisterDfsdmflt3cr2FieldExchShift)
}

// SetExch Extremes detector channel selection
func (r *registerDfsdmflt3cr2Type) SetExch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr2FieldExchMask)|(uint32(value)<<RegisterDfsdmflt3cr2FieldExchShift))
}

const (
	RegisterDfsdmflt3cr2FieldAwdchShift = 16
	RegisterDfsdmflt3cr2FieldAwdchMask  = 0xff0000
)

// GetAwdch Analog watchdog channel selection
func (r *registerDfsdmflt3cr2Type) GetAwdch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cr2FieldAwdchMask) >> RegisterDfsdmflt3cr2FieldAwdchShift)
}

// SetAwdch Analog watchdog channel selection
func (r *registerDfsdmflt3cr2Type) SetAwdch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cr2FieldAwdchMask)|(uint32(value)<<RegisterDfsdmflt3cr2FieldAwdchShift))
}

// registerDfsdmflt3isrType interrupt and status register
type registerDfsdmflt3isrType uint32

const (
	RegisterDfsdmflt3isrFieldJeocfShift = 0
	RegisterDfsdmflt3isrFieldJeocfMask  = 0x1
)

// GetJeocf End of injected conversion flag
func (r *registerDfsdmflt3isrType) GetJeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3isrFieldJeocfMask) != 0
}

// SetJeocf End of injected conversion flag
func (r *registerDfsdmflt3isrType) SetJeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3isrFieldJeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3isrFieldJeocfMask)
	}
}

const (
	RegisterDfsdmflt3isrFieldReocfShift = 1
	RegisterDfsdmflt3isrFieldReocfMask  = 0x2
)

// GetReocf End of regular conversion flag
func (r *registerDfsdmflt3isrType) GetReocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3isrFieldReocfMask) != 0
}

// SetReocf End of regular conversion flag
func (r *registerDfsdmflt3isrType) SetReocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3isrFieldReocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3isrFieldReocfMask)
	}
}

const (
	RegisterDfsdmflt3isrFieldJovrfShift = 2
	RegisterDfsdmflt3isrFieldJovrfMask  = 0x4
)

// GetJovrf Injected conversion overrun flag
func (r *registerDfsdmflt3isrType) GetJovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3isrFieldJovrfMask) != 0
}

// SetJovrf Injected conversion overrun flag
func (r *registerDfsdmflt3isrType) SetJovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3isrFieldJovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3isrFieldJovrfMask)
	}
}

const (
	RegisterDfsdmflt3isrFieldRovrfShift = 3
	RegisterDfsdmflt3isrFieldRovrfMask  = 0x8
)

// GetRovrf Regular conversion overrun flag
func (r *registerDfsdmflt3isrType) GetRovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3isrFieldRovrfMask) != 0
}

// SetRovrf Regular conversion overrun flag
func (r *registerDfsdmflt3isrType) SetRovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3isrFieldRovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3isrFieldRovrfMask)
	}
}

const (
	RegisterDfsdmflt3isrFieldAwdfShift = 4
	RegisterDfsdmflt3isrFieldAwdfMask  = 0x10
)

// GetAwdf Analog watchdog
func (r *registerDfsdmflt3isrType) GetAwdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3isrFieldAwdfMask) != 0
}

// SetAwdf Analog watchdog
func (r *registerDfsdmflt3isrType) SetAwdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3isrFieldAwdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3isrFieldAwdfMask)
	}
}

const (
	RegisterDfsdmflt3isrFieldJcipShift = 13
	RegisterDfsdmflt3isrFieldJcipMask  = 0x2000
)

// GetJcip Injected conversion in progress status
func (r *registerDfsdmflt3isrType) GetJcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3isrFieldJcipMask) != 0
}

// SetJcip Injected conversion in progress status
func (r *registerDfsdmflt3isrType) SetJcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3isrFieldJcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3isrFieldJcipMask)
	}
}

const (
	RegisterDfsdmflt3isrFieldRcipShift = 14
	RegisterDfsdmflt3isrFieldRcipMask  = 0x4000
)

// GetRcip Regular conversion in progress status
func (r *registerDfsdmflt3isrType) GetRcip() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3isrFieldRcipMask) != 0
}

// SetRcip Regular conversion in progress status
func (r *registerDfsdmflt3isrType) SetRcip(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3isrFieldRcipMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3isrFieldRcipMask)
	}
}

const (
	RegisterDfsdmflt3isrFieldCkabfShift = 16
	RegisterDfsdmflt3isrFieldCkabfMask  = 0xff0000
)

// GetCkabf Clock absence flag
func (r *registerDfsdmflt3isrType) GetCkabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3isrFieldCkabfMask) >> RegisterDfsdmflt3isrFieldCkabfShift)
}

// SetCkabf Clock absence flag
func (r *registerDfsdmflt3isrType) SetCkabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3isrFieldCkabfMask)|(uint32(value)<<RegisterDfsdmflt3isrFieldCkabfShift))
}

const (
	RegisterDfsdmflt3isrFieldScdfShift = 24
	RegisterDfsdmflt3isrFieldScdfMask  = 0xff000000
)

// GetScdf short-circuit detector flag
func (r *registerDfsdmflt3isrType) GetScdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3isrFieldScdfMask) >> RegisterDfsdmflt3isrFieldScdfShift)
}

// SetScdf short-circuit detector flag
func (r *registerDfsdmflt3isrType) SetScdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3isrFieldScdfMask)|(uint32(value)<<RegisterDfsdmflt3isrFieldScdfShift))
}

// registerDfsdmflt3icrType interrupt flag clear register
type registerDfsdmflt3icrType uint32

const (
	RegisterDfsdmflt3icrFieldClrjovrfShift = 2
	RegisterDfsdmflt3icrFieldClrjovrfMask  = 0x4
)

// GetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdmflt3icrType) GetClrjovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3icrFieldClrjovrfMask) != 0
}

// SetClrjovrf Clear the injected conversion overrun flag
func (r *registerDfsdmflt3icrType) SetClrjovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3icrFieldClrjovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3icrFieldClrjovrfMask)
	}
}

const (
	RegisterDfsdmflt3icrFieldClrrovrfShift = 3
	RegisterDfsdmflt3icrFieldClrrovrfMask  = 0x8
)

// GetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdmflt3icrType) GetClrrovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3icrFieldClrrovrfMask) != 0
}

// SetClrrovrf Clear the regular conversion overrun flag
func (r *registerDfsdmflt3icrType) SetClrrovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3icrFieldClrrovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3icrFieldClrrovrfMask)
	}
}

const (
	RegisterDfsdmflt3icrFieldClrckabfShift = 16
	RegisterDfsdmflt3icrFieldClrckabfMask  = 0xff0000
)

// GetClrckabf Clear the clock absence flag
func (r *registerDfsdmflt3icrType) GetClrckabf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3icrFieldClrckabfMask) >> RegisterDfsdmflt3icrFieldClrckabfShift)
}

// SetClrckabf Clear the clock absence flag
func (r *registerDfsdmflt3icrType) SetClrckabf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3icrFieldClrckabfMask)|(uint32(value)<<RegisterDfsdmflt3icrFieldClrckabfShift))
}

const (
	RegisterDfsdmflt3icrFieldClrscdfShift = 24
	RegisterDfsdmflt3icrFieldClrscdfMask  = 0xff000000
)

// GetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdmflt3icrType) GetClrscdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3icrFieldClrscdfMask) >> RegisterDfsdmflt3icrFieldClrscdfShift)
}

// SetClrscdf Clear the short-circuit detector flag
func (r *registerDfsdmflt3icrType) SetClrscdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3icrFieldClrscdfMask)|(uint32(value)<<RegisterDfsdmflt3icrFieldClrscdfShift))
}

// registerDfsdmflt3jchgrType injected channel group selection register
type registerDfsdmflt3jchgrType uint32

const (
	RegisterDfsdmflt3jchgrFieldJchgShift = 0
	RegisterDfsdmflt3jchgrFieldJchgMask  = 0xff
)

// GetJchg Injected channel group selection
func (r *registerDfsdmflt3jchgrType) GetJchg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3jchgrFieldJchgMask) >> RegisterDfsdmflt3jchgrFieldJchgShift)
}

// SetJchg Injected channel group selection
func (r *registerDfsdmflt3jchgrType) SetJchg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3jchgrFieldJchgMask)|(uint32(value)<<RegisterDfsdmflt3jchgrFieldJchgShift))
}

// registerDfsdmflt3fcrType filter control register
type registerDfsdmflt3fcrType uint32

const (
	RegisterDfsdmflt3fcrFieldIosrShift = 0
	RegisterDfsdmflt3fcrFieldIosrMask  = 0xff
)

// GetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdmflt3fcrType) GetIosr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3fcrFieldIosrMask) >> RegisterDfsdmflt3fcrFieldIosrShift)
}

// SetIosr Integrator oversampling ratio (averaging length)
func (r *registerDfsdmflt3fcrType) SetIosr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3fcrFieldIosrMask)|(uint32(value)<<RegisterDfsdmflt3fcrFieldIosrShift))
}

const (
	RegisterDfsdmflt3fcrFieldFosrShift = 16
	RegisterDfsdmflt3fcrFieldFosrMask  = 0x3ff0000
)

// GetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdmflt3fcrType) GetFosr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3fcrFieldFosrMask) >> RegisterDfsdmflt3fcrFieldFosrShift)
}

// SetFosr Sinc filter oversampling ratio (decimation rate)
func (r *registerDfsdmflt3fcrType) SetFosr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3fcrFieldFosrMask)|(uint32(value)<<RegisterDfsdmflt3fcrFieldFosrShift))
}

const (
	RegisterDfsdmflt3fcrFieldFordShift = 29
	RegisterDfsdmflt3fcrFieldFordMask  = 0xe0000000
)

// GetFord Sinc filter order
func (r *registerDfsdmflt3fcrType) GetFord() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3fcrFieldFordMask) >> RegisterDfsdmflt3fcrFieldFordShift)
}

// SetFord Sinc filter order
func (r *registerDfsdmflt3fcrType) SetFord(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3fcrFieldFordMask)|(uint32(value)<<RegisterDfsdmflt3fcrFieldFordShift))
}

// registerDfsdmflt3jdatarType data register for injected group
type registerDfsdmflt3jdatarType uint32

const (
	RegisterDfsdmflt3jdatarFieldJdatachShift = 0
	RegisterDfsdmflt3jdatarFieldJdatachMask  = 0x7
)

// GetJdatach Injected channel most recently converted
func (r *registerDfsdmflt3jdatarType) GetJdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3jdatarFieldJdatachMask) >> RegisterDfsdmflt3jdatarFieldJdatachShift)
}

// SetJdatach Injected channel most recently converted
func (r *registerDfsdmflt3jdatarType) SetJdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3jdatarFieldJdatachMask)|(uint32(value)<<RegisterDfsdmflt3jdatarFieldJdatachShift))
}

const (
	RegisterDfsdmflt3jdatarFieldJdataShift = 8
	RegisterDfsdmflt3jdatarFieldJdataMask  = 0xffffff00
)

// GetJdata Injected group conversion data
func (r *registerDfsdmflt3jdatarType) GetJdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3jdatarFieldJdataMask) >> RegisterDfsdmflt3jdatarFieldJdataShift)
}

// SetJdata Injected group conversion data
func (r *registerDfsdmflt3jdatarType) SetJdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3jdatarFieldJdataMask)|(uint32(value)<<RegisterDfsdmflt3jdatarFieldJdataShift))
}

// registerDfsdmflt3rdatarType data register for the regular channel
type registerDfsdmflt3rdatarType uint32

const (
	RegisterDfsdmflt3rdatarFieldRdatachShift = 0
	RegisterDfsdmflt3rdatarFieldRdatachMask  = 0x7
)

// GetRdatach Regular channel most recently converted
func (r *registerDfsdmflt3rdatarType) GetRdatach() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3rdatarFieldRdatachMask) >> RegisterDfsdmflt3rdatarFieldRdatachShift)
}

// SetRdatach Regular channel most recently converted
func (r *registerDfsdmflt3rdatarType) SetRdatach(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3rdatarFieldRdatachMask)|(uint32(value)<<RegisterDfsdmflt3rdatarFieldRdatachShift))
}

const (
	RegisterDfsdmflt3rdatarFieldRpendShift = 4
	RegisterDfsdmflt3rdatarFieldRpendMask  = 0x10
)

// GetRpend Regular channel pending data
func (r *registerDfsdmflt3rdatarType) GetRpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3rdatarFieldRpendMask) != 0
}

// SetRpend Regular channel pending data
func (r *registerDfsdmflt3rdatarType) SetRpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDfsdmflt3rdatarFieldRpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3rdatarFieldRpendMask)
	}
}

const (
	RegisterDfsdmflt3rdatarFieldRdataShift = 8
	RegisterDfsdmflt3rdatarFieldRdataMask  = 0xffffff00
)

// GetRdata Regular channel conversion data
func (r *registerDfsdmflt3rdatarType) GetRdata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3rdatarFieldRdataMask) >> RegisterDfsdmflt3rdatarFieldRdataShift)
}

// SetRdata Regular channel conversion data
func (r *registerDfsdmflt3rdatarType) SetRdata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3rdatarFieldRdataMask)|(uint32(value)<<RegisterDfsdmflt3rdatarFieldRdataShift))
}

// registerDfsdmflt3awhtrType analog watchdog high threshold register
type registerDfsdmflt3awhtrType uint32

const (
	RegisterDfsdmflt3awhtrFieldBkawhShift = 0
	RegisterDfsdmflt3awhtrFieldBkawhMask  = 0xf
)

// GetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdmflt3awhtrType) GetBkawh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3awhtrFieldBkawhMask) >> RegisterDfsdmflt3awhtrFieldBkawhShift)
}

// SetBkawh Break signal assignment to analog watchdog high threshold event
func (r *registerDfsdmflt3awhtrType) SetBkawh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3awhtrFieldBkawhMask)|(uint32(value)<<RegisterDfsdmflt3awhtrFieldBkawhShift))
}

const (
	RegisterDfsdmflt3awhtrFieldAwhtShift = 8
	RegisterDfsdmflt3awhtrFieldAwhtMask  = 0xffffff00
)

// GetAwht Analog watchdog high threshold
func (r *registerDfsdmflt3awhtrType) GetAwht() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3awhtrFieldAwhtMask) >> RegisterDfsdmflt3awhtrFieldAwhtShift)
}

// SetAwht Analog watchdog high threshold
func (r *registerDfsdmflt3awhtrType) SetAwht(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3awhtrFieldAwhtMask)|(uint32(value)<<RegisterDfsdmflt3awhtrFieldAwhtShift))
}

// registerDfsdmflt3awltrType analog watchdog low threshold register
type registerDfsdmflt3awltrType uint32

const (
	RegisterDfsdmflt3awltrFieldBkawlShift = 0
	RegisterDfsdmflt3awltrFieldBkawlMask  = 0xf
)

// GetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdmflt3awltrType) GetBkawl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3awltrFieldBkawlMask) >> RegisterDfsdmflt3awltrFieldBkawlShift)
}

// SetBkawl Break signal assignment to analog watchdog low threshold event
func (r *registerDfsdmflt3awltrType) SetBkawl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3awltrFieldBkawlMask)|(uint32(value)<<RegisterDfsdmflt3awltrFieldBkawlShift))
}

const (
	RegisterDfsdmflt3awltrFieldAwltShift = 8
	RegisterDfsdmflt3awltrFieldAwltMask  = 0xffffff00
)

// GetAwlt Analog watchdog low threshold
func (r *registerDfsdmflt3awltrType) GetAwlt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3awltrFieldAwltMask) >> RegisterDfsdmflt3awltrFieldAwltShift)
}

// SetAwlt Analog watchdog low threshold
func (r *registerDfsdmflt3awltrType) SetAwlt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3awltrFieldAwltMask)|(uint32(value)<<RegisterDfsdmflt3awltrFieldAwltShift))
}

// registerDfsdmflt3awsrType analog watchdog status register
type registerDfsdmflt3awsrType uint32

const (
	RegisterDfsdmflt3awsrFieldAwltfShift = 0
	RegisterDfsdmflt3awsrFieldAwltfMask  = 0xff
)

// GetAwltf Analog watchdog low threshold flag
func (r *registerDfsdmflt3awsrType) GetAwltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3awsrFieldAwltfMask) >> RegisterDfsdmflt3awsrFieldAwltfShift)
}

// SetAwltf Analog watchdog low threshold flag
func (r *registerDfsdmflt3awsrType) SetAwltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3awsrFieldAwltfMask)|(uint32(value)<<RegisterDfsdmflt3awsrFieldAwltfShift))
}

const (
	RegisterDfsdmflt3awsrFieldAwhtfShift = 8
	RegisterDfsdmflt3awsrFieldAwhtfMask  = 0xff00
)

// GetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdmflt3awsrType) GetAwhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3awsrFieldAwhtfMask) >> RegisterDfsdmflt3awsrFieldAwhtfShift)
}

// SetAwhtf Analog watchdog high threshold flag
func (r *registerDfsdmflt3awsrType) SetAwhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3awsrFieldAwhtfMask)|(uint32(value)<<RegisterDfsdmflt3awsrFieldAwhtfShift))
}

// registerDfsdmflt3awcfrType analog watchdog clear flag register
type registerDfsdmflt3awcfrType uint32

const (
	RegisterDfsdmflt3awcfrFieldClrawltfShift = 0
	RegisterDfsdmflt3awcfrFieldClrawltfMask  = 0xff
)

// GetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdmflt3awcfrType) GetClrawltf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3awcfrFieldClrawltfMask) >> RegisterDfsdmflt3awcfrFieldClrawltfShift)
}

// SetClrawltf Clear the analog watchdog low threshold flag
func (r *registerDfsdmflt3awcfrType) SetClrawltf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3awcfrFieldClrawltfMask)|(uint32(value)<<RegisterDfsdmflt3awcfrFieldClrawltfShift))
}

const (
	RegisterDfsdmflt3awcfrFieldClrawhtfShift = 8
	RegisterDfsdmflt3awcfrFieldClrawhtfMask  = 0xff00
)

// GetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdmflt3awcfrType) GetClrawhtf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3awcfrFieldClrawhtfMask) >> RegisterDfsdmflt3awcfrFieldClrawhtfShift)
}

// SetClrawhtf Clear the analog watchdog high threshold flag
func (r *registerDfsdmflt3awcfrType) SetClrawhtf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3awcfrFieldClrawhtfMask)|(uint32(value)<<RegisterDfsdmflt3awcfrFieldClrawhtfShift))
}

// registerDfsdmflt3exmaxType Extremes detector maximum register
type registerDfsdmflt3exmaxType uint32

const (
	RegisterDfsdmflt3exmaxFieldExmaxchShift = 0
	RegisterDfsdmflt3exmaxFieldExmaxchMask  = 0x7
)

// GetExmaxch Extremes detector maximum data channel
func (r *registerDfsdmflt3exmaxType) GetExmaxch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3exmaxFieldExmaxchMask) >> RegisterDfsdmflt3exmaxFieldExmaxchShift)
}

// SetExmaxch Extremes detector maximum data channel
func (r *registerDfsdmflt3exmaxType) SetExmaxch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3exmaxFieldExmaxchMask)|(uint32(value)<<RegisterDfsdmflt3exmaxFieldExmaxchShift))
}

const (
	RegisterDfsdmflt3exmaxFieldExmaxShift = 8
	RegisterDfsdmflt3exmaxFieldExmaxMask  = 0xffffff00
)

// GetExmax Extremes detector maximum value
func (r *registerDfsdmflt3exmaxType) GetExmax() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3exmaxFieldExmaxMask) >> RegisterDfsdmflt3exmaxFieldExmaxShift)
}

// SetExmax Extremes detector maximum value
func (r *registerDfsdmflt3exmaxType) SetExmax(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3exmaxFieldExmaxMask)|(uint32(value)<<RegisterDfsdmflt3exmaxFieldExmaxShift))
}

// registerDfsdmflt3exminType Extremes detector minimum register
type registerDfsdmflt3exminType uint32

const (
	RegisterDfsdmflt3exminFieldExminchShift = 0
	RegisterDfsdmflt3exminFieldExminchMask  = 0x7
)

// GetExminch Extremes detector minimum data channel
func (r *registerDfsdmflt3exminType) GetExminch() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3exminFieldExminchMask) >> RegisterDfsdmflt3exminFieldExminchShift)
}

// SetExminch Extremes detector minimum data channel
func (r *registerDfsdmflt3exminType) SetExminch(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3exminFieldExminchMask)|(uint32(value)<<RegisterDfsdmflt3exminFieldExminchShift))
}

const (
	RegisterDfsdmflt3exminFieldExminShift = 8
	RegisterDfsdmflt3exminFieldExminMask  = 0xffffff00
)

// GetExmin EXMIN
func (r *registerDfsdmflt3exminType) GetExmin() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3exminFieldExminMask) >> RegisterDfsdmflt3exminFieldExminShift)
}

// SetExmin EXMIN
func (r *registerDfsdmflt3exminType) SetExmin(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3exminFieldExminMask)|(uint32(value)<<RegisterDfsdmflt3exminFieldExminShift))
}

// registerDfsdmflt3cnvtimrType conversion timer register
type registerDfsdmflt3cnvtimrType uint32

const (
	RegisterDfsdmflt3cnvtimrFieldCnvcntShift = 4
	RegisterDfsdmflt3cnvtimrFieldCnvcntMask  = 0xfffffff0
)

// GetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdmflt3cnvtimrType) GetCnvcnt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDfsdmflt3cnvtimrFieldCnvcntMask) >> RegisterDfsdmflt3cnvtimrFieldCnvcntShift)
}

// SetCnvcnt 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
func (r *registerDfsdmflt3cnvtimrType) SetCnvcnt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDfsdmflt3cnvtimrFieldCnvcntMask)|(uint32(value)<<RegisterDfsdmflt3cnvtimrFieldCnvcntShift))
}
