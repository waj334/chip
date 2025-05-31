//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package otg_hs_host

import (
	"unsafe"
	"volatile"
)

var (
	Otg1_hs_host = (*_otg_hs_host)(unsafe.Pointer(uintptr(0x40040400)))
	Otg2_hs_host = (*_otg_hs_host)(unsafe.Pointer(uintptr(0x40080400)))

	Instances = [2]*_otg_hs_host{
		Otg1_hs_host,
		Otg2_hs_host,
	}
)

type _otg_hs_host struct {
	Otghshcfg       registerOtghshcfgType
	Otghshfir       registerOtghshfirType
	Otghshfnum      registerOtghshfnumType
	_               [4]byte
	Otghshptxsts    registerOtghshptxstsType
	Otghshaint      registerOtghshaintType
	Otghshaintmsk   registerOtghshaintmskType
	_               [36]byte
	Otghshprt       registerOtghshprtType
	_               [188]byte
	Otghshcchar0    registerOtghshcchar0Type
	Otghshcsplt0    registerOtghshcsplt0Type
	Otghshcint0     registerOtghshcint0Type
	Otghshcintmsk0  registerOtghshcintmsk0Type
	Otghshctsiz0    registerOtghshctsiz0Type
	Otghshcdma0     registerOtghshcdma0Type
	_               [8]byte
	Otghshcchar1    registerOtghshcchar1Type
	Otghshcsplt1    registerOtghshcsplt1Type
	Otghshcint1     registerOtghshcint1Type
	Otghshcintmsk1  registerOtghshcintmsk1Type
	Otghshctsiz1    registerOtghshctsiz1Type
	Otghshcdma1     registerOtghshcdma1Type
	_               [8]byte
	Otghshcchar2    registerOtghshcchar2Type
	Otghshcsplt2    registerOtghshcsplt2Type
	Otghshcint2     registerOtghshcint2Type
	Otghshcintmsk2  registerOtghshcintmsk2Type
	Otghshctsiz2    registerOtghshctsiz2Type
	Otghshcdma2     registerOtghshcdma2Type
	_               [8]byte
	Otghshcchar3    registerOtghshcchar3Type
	Otghshcsplt3    registerOtghshcsplt3Type
	Otghshcint3     registerOtghshcint3Type
	Otghshcintmsk3  registerOtghshcintmsk3Type
	Otghshctsiz3    registerOtghshctsiz3Type
	Otghshcdma3     registerOtghshcdma3Type
	_               [8]byte
	Otghshcchar4    registerOtghshcchar4Type
	Otghshcsplt4    registerOtghshcsplt4Type
	Otghshcint4     registerOtghshcint4Type
	Otghshcintmsk4  registerOtghshcintmsk4Type
	Otghshctsiz4    registerOtghshctsiz4Type
	Otghshcdma4     registerOtghshcdma4Type
	_               [8]byte
	Otghshcchar5    registerOtghshcchar5Type
	Otghshcsplt5    registerOtghshcsplt5Type
	Otghshcint5     registerOtghshcint5Type
	Otghshcintmsk5  registerOtghshcintmsk5Type
	Otghshctsiz5    registerOtghshctsiz5Type
	Otghshcdma5     registerOtghshcdma5Type
	_               [8]byte
	Otghshcchar6    registerOtghshcchar6Type
	Otghshcsplt6    registerOtghshcsplt6Type
	Otghshcint6     registerOtghshcint6Type
	Otghshcintmsk6  registerOtghshcintmsk6Type
	Otghshctsiz6    registerOtghshctsiz6Type
	Otghshcdma6     registerOtghshcdma6Type
	_               [8]byte
	Otghshcchar7    registerOtghshcchar7Type
	Otghshcsplt7    registerOtghshcsplt7Type
	Otghshcint7     registerOtghshcint7Type
	Otghshcintmsk7  registerOtghshcintmsk7Type
	Otghshctsiz7    registerOtghshctsiz7Type
	Otghshcdma7     registerOtghshcdma7Type
	_               [8]byte
	Otghshcchar8    registerOtghshcchar8Type
	Otghshcsplt8    registerOtghshcsplt8Type
	Otghshcint8     registerOtghshcint8Type
	Otghshcintmsk8  registerOtghshcintmsk8Type
	Otghshctsiz8    registerOtghshctsiz8Type
	Otghshcdma8     registerOtghshcdma8Type
	_               [8]byte
	Otghshcchar9    registerOtghshcchar9Type
	Otghshcsplt9    registerOtghshcsplt9Type
	Otghshcint9     registerOtghshcint9Type
	Otghshcintmsk9  registerOtghshcintmsk9Type
	Otghshctsiz9    registerOtghshctsiz9Type
	Otghshcdma9     registerOtghshcdma9Type
	_               [8]byte
	Otghshcchar10   registerOtghshcchar10Type
	Otghshcsplt10   registerOtghshcsplt10Type
	Otghshcint10    registerOtghshcint10Type
	Otghshcintmsk10 registerOtghshcintmsk10Type
	Otghshctsiz10   registerOtghshctsiz10Type
	Otghshcdma10    registerOtghshcdma10Type
	_               [8]byte
	Otghshcchar11   registerOtghshcchar11Type
	Otghshcsplt11   registerOtghshcsplt11Type
	Otghshcint11    registerOtghshcint11Type
	Otghshcintmsk11 registerOtghshcintmsk11Type
	Otghshctsiz11   registerOtghshctsiz11Type
	Otghshcdma11    registerOtghshcdma11Type
	Otghshcchar12   registerOtghshcchar12Type
	Otghshcsplt12   registerOtghshcsplt12Type
	Otghshcint12    registerOtghshcint12Type
	Otghshcintmsk12 registerOtghshcintmsk12Type
	Otghshctsiz12   registerOtghshctsiz12Type
	Otghshcdma12    registerOtghshcdma12Type
	Otghshcchar13   registerOtghshcchar13Type
	Otghshcsplt13   registerOtghshcsplt13Type
	Otghshcint13    registerOtghshcint13Type
	Otghshcintmsk13 registerOtghshcintmsk13Type
	Otghshctsiz13   registerOtghshctsiz13Type
	Otghshcdma13    registerOtghshcdma13Type
	Otghshcchar14   registerOtghshcchar14Type
	Otghshcsplt14   registerOtghshcsplt14Type
	Otghshcint14    registerOtghshcint14Type
	Otghshcintmsk14 registerOtghshcintmsk14Type
	Otghshctsiz14   registerOtghshctsiz14Type
	Otghshcdma14    registerOtghshcdma14Type
	Otghshcchar15   registerOtghshcchar15Type
	Otghshcsplt15   registerOtghshcsplt15Type
	Otghshcint15    registerOtghshcint15Type
	Otghshcintmsk15 registerOtghshcintmsk15Type
	Otghshctsiz15   registerOtghshctsiz15Type
	Otghshcdma15    registerOtghshcdma15Type
}

// registerOtghshcfgType OTG_HS host configuration register
type registerOtghshcfgType uint32

const (
	RegisterOtghshcfgFieldFslspcsShift = 0
	RegisterOtghshcfgFieldFslspcsMask  = 0x3
)

// GetFslspcs FS/LS PHY clock select
func (r *registerOtghshcfgType) GetFslspcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcfgFieldFslspcsMask) >> RegisterOtghshcfgFieldFslspcsShift)
}

// SetFslspcs FS/LS PHY clock select
func (r *registerOtghshcfgType) SetFslspcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcfgFieldFslspcsMask)|(uint32(value)<<RegisterOtghshcfgFieldFslspcsShift))
}

const (
	RegisterOtghshcfgFieldFslssShift = 2
	RegisterOtghshcfgFieldFslssMask  = 0x4
)

// GetFslss FS- and LS-only support
func (r *registerOtghshcfgType) GetFslss() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcfgFieldFslssMask) != 0
}

// registerOtghshfirType OTG_HS Host frame interval register
type registerOtghshfirType uint32

const (
	RegisterOtghshfirFieldFrivlShift = 0
	RegisterOtghshfirFieldFrivlMask  = 0xffff
)

// GetFrivl Frame interval
func (r *registerOtghshfirType) GetFrivl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshfirFieldFrivlMask) >> RegisterOtghshfirFieldFrivlShift)
}

// SetFrivl Frame interval
func (r *registerOtghshfirType) SetFrivl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshfirFieldFrivlMask)|(uint32(value)<<RegisterOtghshfirFieldFrivlShift))
}

// registerOtghshfnumType OTG_HS host frame number/frame time remaining register
type registerOtghshfnumType uint32

const (
	RegisterOtghshfnumFieldFrnumShift = 0
	RegisterOtghshfnumFieldFrnumMask  = 0xffff
)

// GetFrnum Frame number
func (r *registerOtghshfnumType) GetFrnum() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshfnumFieldFrnumMask) >> RegisterOtghshfnumFieldFrnumShift)
}

// SetFrnum Frame number
func (r *registerOtghshfnumType) SetFrnum(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshfnumFieldFrnumMask)|(uint32(value)<<RegisterOtghshfnumFieldFrnumShift))
}

const (
	RegisterOtghshfnumFieldFtremShift = 16
	RegisterOtghshfnumFieldFtremMask  = 0xffff0000
)

// GetFtrem Frame time remaining
func (r *registerOtghshfnumType) GetFtrem() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshfnumFieldFtremMask) >> RegisterOtghshfnumFieldFtremShift)
}

// SetFtrem Frame time remaining
func (r *registerOtghshfnumType) SetFtrem(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshfnumFieldFtremMask)|(uint32(value)<<RegisterOtghshfnumFieldFtremShift))
}

// registerOtghshptxstsType OTG_HS_Host periodic transmit FIFO/queue status register
type registerOtghshptxstsType uint32

const (
	RegisterOtghshptxstsFieldPtxfsavlShift = 0
	RegisterOtghshptxstsFieldPtxfsavlMask  = 0xffff
)

// GetPtxfsavl Periodic transmit data FIFO space available
func (r *registerOtghshptxstsType) GetPtxfsavl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshptxstsFieldPtxfsavlMask) >> RegisterOtghshptxstsFieldPtxfsavlShift)
}

// SetPtxfsavl Periodic transmit data FIFO space available
func (r *registerOtghshptxstsType) SetPtxfsavl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshptxstsFieldPtxfsavlMask)|(uint32(value)<<RegisterOtghshptxstsFieldPtxfsavlShift))
}

const (
	RegisterOtghshptxstsFieldPtxqsavShift = 16
	RegisterOtghshptxstsFieldPtxqsavMask  = 0xff0000
)

// GetPtxqsav Periodic transmit request queue space available
func (r *registerOtghshptxstsType) GetPtxqsav() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshptxstsFieldPtxqsavMask) >> RegisterOtghshptxstsFieldPtxqsavShift)
}

const (
	RegisterOtghshptxstsFieldPtxqtopShift = 24
	RegisterOtghshptxstsFieldPtxqtopMask  = 0xff000000
)

// GetPtxqtop Top of the periodic transmit request queue
func (r *registerOtghshptxstsType) GetPtxqtop() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshptxstsFieldPtxqtopMask) >> RegisterOtghshptxstsFieldPtxqtopShift)
}

// registerOtghshaintType OTG_HS Host all channels interrupt register
type registerOtghshaintType uint32

const (
	RegisterOtghshaintFieldHaintShift = 0
	RegisterOtghshaintFieldHaintMask  = 0xffff
)

// GetHaint Channel interrupts
func (r *registerOtghshaintType) GetHaint() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshaintFieldHaintMask) >> RegisterOtghshaintFieldHaintShift)
}

// SetHaint Channel interrupts
func (r *registerOtghshaintType) SetHaint(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshaintFieldHaintMask)|(uint32(value)<<RegisterOtghshaintFieldHaintShift))
}

// registerOtghshaintmskType OTG_HS host all channels interrupt mask register
type registerOtghshaintmskType uint32

const (
	RegisterOtghshaintmskFieldHaintmShift = 0
	RegisterOtghshaintmskFieldHaintmMask  = 0xffff
)

// GetHaintm Channel interrupt mask
func (r *registerOtghshaintmskType) GetHaintm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshaintmskFieldHaintmMask) >> RegisterOtghshaintmskFieldHaintmShift)
}

// SetHaintm Channel interrupt mask
func (r *registerOtghshaintmskType) SetHaintm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshaintmskFieldHaintmMask)|(uint32(value)<<RegisterOtghshaintmskFieldHaintmShift))
}

// registerOtghshprtType OTG_HS host port control and status register
type registerOtghshprtType uint32

const (
	RegisterOtghshprtFieldPcstsShift = 0
	RegisterOtghshprtFieldPcstsMask  = 0x1
)

// GetPcsts Port connect status
func (r *registerOtghshprtType) GetPcsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPcstsMask) != 0
}

const (
	RegisterOtghshprtFieldPcdetShift = 1
	RegisterOtghshprtFieldPcdetMask  = 0x2
)

// GetPcdet Port connect detected
func (r *registerOtghshprtType) GetPcdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPcdetMask) != 0
}

// SetPcdet Port connect detected
func (r *registerOtghshprtType) SetPcdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshprtFieldPcdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshprtFieldPcdetMask)
	}
}

const (
	RegisterOtghshprtFieldPenaShift = 2
	RegisterOtghshprtFieldPenaMask  = 0x4
)

// GetPena Port enable
func (r *registerOtghshprtType) GetPena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPenaMask) != 0
}

// SetPena Port enable
func (r *registerOtghshprtType) SetPena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshprtFieldPenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshprtFieldPenaMask)
	}
}

const (
	RegisterOtghshprtFieldPenchngShift = 3
	RegisterOtghshprtFieldPenchngMask  = 0x8
)

// GetPenchng Port enable/disable change
func (r *registerOtghshprtType) GetPenchng() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPenchngMask) != 0
}

// SetPenchng Port enable/disable change
func (r *registerOtghshprtType) SetPenchng(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshprtFieldPenchngMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshprtFieldPenchngMask)
	}
}

const (
	RegisterOtghshprtFieldPocaShift = 4
	RegisterOtghshprtFieldPocaMask  = 0x10
)

// GetPoca Port overcurrent active
func (r *registerOtghshprtType) GetPoca() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPocaMask) != 0
}

const (
	RegisterOtghshprtFieldPocchngShift = 5
	RegisterOtghshprtFieldPocchngMask  = 0x20
)

// GetPocchng Port overcurrent change
func (r *registerOtghshprtType) GetPocchng() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPocchngMask) != 0
}

// SetPocchng Port overcurrent change
func (r *registerOtghshprtType) SetPocchng(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshprtFieldPocchngMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshprtFieldPocchngMask)
	}
}

const (
	RegisterOtghshprtFieldPresShift = 6
	RegisterOtghshprtFieldPresMask  = 0x40
)

// GetPres Port resume
func (r *registerOtghshprtType) GetPres() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPresMask) != 0
}

// SetPres Port resume
func (r *registerOtghshprtType) SetPres(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshprtFieldPresMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshprtFieldPresMask)
	}
}

const (
	RegisterOtghshprtFieldPsuspShift = 7
	RegisterOtghshprtFieldPsuspMask  = 0x80
)

// GetPsusp Port suspend
func (r *registerOtghshprtType) GetPsusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPsuspMask) != 0
}

// SetPsusp Port suspend
func (r *registerOtghshprtType) SetPsusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshprtFieldPsuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshprtFieldPsuspMask)
	}
}

const (
	RegisterOtghshprtFieldPrstShift = 8
	RegisterOtghshprtFieldPrstMask  = 0x100
)

// GetPrst Port reset
func (r *registerOtghshprtType) GetPrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPrstMask) != 0
}

// SetPrst Port reset
func (r *registerOtghshprtType) SetPrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshprtFieldPrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshprtFieldPrstMask)
	}
}

const (
	RegisterOtghshprtFieldPlstsShift = 10
	RegisterOtghshprtFieldPlstsMask  = 0xc00
)

// GetPlsts Port line status
func (r *registerOtghshprtType) GetPlsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPlstsMask) >> RegisterOtghshprtFieldPlstsShift)
}

const (
	RegisterOtghshprtFieldPpwrShift = 12
	RegisterOtghshprtFieldPpwrMask  = 0x1000
)

// GetPpwr Port power
func (r *registerOtghshprtType) GetPpwr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPpwrMask) != 0
}

// SetPpwr Port power
func (r *registerOtghshprtType) SetPpwr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshprtFieldPpwrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshprtFieldPpwrMask)
	}
}

const (
	RegisterOtghshprtFieldPtctlShift = 13
	RegisterOtghshprtFieldPtctlMask  = 0x1e000
)

// GetPtctl Port test control
func (r *registerOtghshprtType) GetPtctl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPtctlMask) >> RegisterOtghshprtFieldPtctlShift)
}

// SetPtctl Port test control
func (r *registerOtghshprtType) SetPtctl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshprtFieldPtctlMask)|(uint32(value)<<RegisterOtghshprtFieldPtctlShift))
}

const (
	RegisterOtghshprtFieldPspdShift = 17
	RegisterOtghshprtFieldPspdMask  = 0x60000
)

// GetPspd Port speed
func (r *registerOtghshprtType) GetPspd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPspdMask) >> RegisterOtghshprtFieldPspdShift)
}

// registerOtghshcchar0Type OTG_HS host channel-0 characteristics register
type registerOtghshcchar0Type uint32

const (
	RegisterOtghshcchar0FieldMpsizShift = 0
	RegisterOtghshcchar0FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar0Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldMpsizMask) >> RegisterOtghshcchar0FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar0Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar0FieldMpsizShift))
}

const (
	RegisterOtghshcchar0FieldEpnumShift = 11
	RegisterOtghshcchar0FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar0Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldEpnumMask) >> RegisterOtghshcchar0FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar0Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar0FieldEpnumShift))
}

const (
	RegisterOtghshcchar0FieldEpdirShift = 15
	RegisterOtghshcchar0FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar0Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar0Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar0FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar0FieldLsdevShift = 17
	RegisterOtghshcchar0FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar0Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar0Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar0FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar0FieldEptypShift = 18
	RegisterOtghshcchar0FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar0Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldEptypMask) >> RegisterOtghshcchar0FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar0Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar0FieldEptypShift))
}

const (
	RegisterOtghshcchar0FieldMcShift = 20
	RegisterOtghshcchar0FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar0Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldMcMask) >> RegisterOtghshcchar0FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar0Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldMcMask)|(uint32(value)<<RegisterOtghshcchar0FieldMcShift))
}

const (
	RegisterOtghshcchar0FieldDadShift = 22
	RegisterOtghshcchar0FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar0Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldDadMask) >> RegisterOtghshcchar0FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar0Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldDadMask)|(uint32(value)<<RegisterOtghshcchar0FieldDadShift))
}

const (
	RegisterOtghshcchar0FieldOddfrmShift = 29
	RegisterOtghshcchar0FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar0Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar0Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar0FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar0FieldChdisShift = 30
	RegisterOtghshcchar0FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar0Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar0Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar0FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar0FieldChenaShift = 31
	RegisterOtghshcchar0FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar0Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar0Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar0FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldChenaMask)
	}
}

// registerOtghshcsplt0Type OTG_HS host channel-0 split control register
type registerOtghshcsplt0Type uint32

const (
	RegisterOtghshcsplt0FieldPrtaddrShift = 0
	RegisterOtghshcsplt0FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt0Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt0FieldPrtaddrMask) >> RegisterOtghshcsplt0FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt0Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt0FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt0FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt0FieldHubaddrShift = 7
	RegisterOtghshcsplt0FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt0Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt0FieldHubaddrMask) >> RegisterOtghshcsplt0FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt0Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt0FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt0FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt0FieldXactposShift = 14
	RegisterOtghshcsplt0FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt0Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt0FieldXactposMask) >> RegisterOtghshcsplt0FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt0Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt0FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt0FieldXactposShift))
}

const (
	RegisterOtghshcsplt0FieldComplspltShift = 16
	RegisterOtghshcsplt0FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt0Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt0FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt0Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt0FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt0FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt0FieldSplitenShift = 31
	RegisterOtghshcsplt0FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt0Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt0FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt0Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt0FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt0FieldSplitenMask)
	}
}

// registerOtghshcint0Type OTG_HS host channel-11 interrupt register
type registerOtghshcint0Type uint32

const (
	RegisterOtghshcint0FieldXfrcShift = 0
	RegisterOtghshcint0FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint0Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint0Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint0FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint0FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint0FieldChhShift = 1
	RegisterOtghshcint0FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint0Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint0Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint0FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint0FieldChhMask)
	}
}

const (
	RegisterOtghshcint0FieldAhberrShift = 2
	RegisterOtghshcint0FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint0Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint0Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint0FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint0FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint0FieldStallShift = 3
	RegisterOtghshcint0FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint0Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint0Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint0FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint0FieldStallMask)
	}
}

const (
	RegisterOtghshcint0FieldNakShift = 4
	RegisterOtghshcint0FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint0Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint0Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint0FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint0FieldNakMask)
	}
}

const (
	RegisterOtghshcint0FieldAckShift = 5
	RegisterOtghshcint0FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint0Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint0Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint0FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint0FieldAckMask)
	}
}

const (
	RegisterOtghshcint0FieldNyetShift = 6
	RegisterOtghshcint0FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint0Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint0Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint0FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint0FieldNyetMask)
	}
}

const (
	RegisterOtghshcint0FieldTxerrShift = 7
	RegisterOtghshcint0FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint0Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint0Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint0FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint0FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint0FieldBberrShift = 8
	RegisterOtghshcint0FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint0Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint0Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint0FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint0FieldBberrMask)
	}
}

const (
	RegisterOtghshcint0FieldFrmorShift = 9
	RegisterOtghshcint0FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint0Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint0Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint0FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint0FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint0FieldDterrShift = 10
	RegisterOtghshcint0FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint0Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint0Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint0FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint0FieldDterrMask)
	}
}

// registerOtghshcintmsk0Type OTG_HS host channel-11 interrupt mask register
type registerOtghshcintmsk0Type uint32

const (
	RegisterOtghshcintmsk0FieldXfrcmShift = 0
	RegisterOtghshcintmsk0FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk0Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk0Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk0FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk0FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk0FieldChhmShift = 1
	RegisterOtghshcintmsk0FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk0Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk0Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk0FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk0FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk0FieldAhberrShift = 2
	RegisterOtghshcintmsk0FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk0Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk0Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk0FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk0FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk0FieldStallmShift = 3
	RegisterOtghshcintmsk0FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk0Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk0Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk0FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk0FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk0FieldNakmShift = 4
	RegisterOtghshcintmsk0FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk0Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk0Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk0FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk0FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk0FieldAckmShift = 5
	RegisterOtghshcintmsk0FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk0Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk0Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk0FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk0FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk0FieldNyetShift = 6
	RegisterOtghshcintmsk0FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtghshcintmsk0Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtghshcintmsk0Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk0FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk0FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk0FieldTxerrmShift = 7
	RegisterOtghshcintmsk0FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtghshcintmsk0Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtghshcintmsk0Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk0FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk0FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk0FieldBberrmShift = 8
	RegisterOtghshcintmsk0FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtghshcintmsk0Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtghshcintmsk0Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk0FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk0FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk0FieldFrmormShift = 9
	RegisterOtghshcintmsk0FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk0Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk0Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk0FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk0FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk0FieldDterrmShift = 10
	RegisterOtghshcintmsk0FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk0Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk0Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk0FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk0FieldDterrmMask)
	}
}

// registerOtghshctsiz0Type OTG_HS host channel-11 transfer size register
type registerOtghshctsiz0Type uint32

const (
	RegisterOtghshctsiz0FieldXfrsizShift = 0
	RegisterOtghshctsiz0FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz0Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz0FieldXfrsizMask) >> RegisterOtghshctsiz0FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz0Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz0FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz0FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz0FieldPktcntShift = 19
	RegisterOtghshctsiz0FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz0Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz0FieldPktcntMask) >> RegisterOtghshctsiz0FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz0Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz0FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz0FieldPktcntShift))
}

const (
	RegisterOtghshctsiz0FieldDpidShift = 29
	RegisterOtghshctsiz0FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz0Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz0FieldDpidMask) >> RegisterOtghshctsiz0FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz0Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz0FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz0FieldDpidShift))
}

// registerOtghshcdma0Type OTG_HS host channel-0 DMA address register
type registerOtghshcdma0Type uint32

const (
	RegisterOtghshcdma0FieldDmaaddrShift = 0
	RegisterOtghshcdma0FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma0Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma0FieldDmaaddrMask) >> RegisterOtghshcdma0FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma0Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma0FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma0FieldDmaaddrShift))
}

// registerOtghshcchar1Type OTG_HS host channel-1 characteristics register
type registerOtghshcchar1Type uint32

const (
	RegisterOtghshcchar1FieldMpsizShift = 0
	RegisterOtghshcchar1FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar1Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldMpsizMask) >> RegisterOtghshcchar1FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar1Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar1FieldMpsizShift))
}

const (
	RegisterOtghshcchar1FieldEpnumShift = 11
	RegisterOtghshcchar1FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar1Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldEpnumMask) >> RegisterOtghshcchar1FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar1Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar1FieldEpnumShift))
}

const (
	RegisterOtghshcchar1FieldEpdirShift = 15
	RegisterOtghshcchar1FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar1Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar1Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar1FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar1FieldLsdevShift = 17
	RegisterOtghshcchar1FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar1Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar1Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar1FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar1FieldEptypShift = 18
	RegisterOtghshcchar1FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar1Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldEptypMask) >> RegisterOtghshcchar1FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar1Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar1FieldEptypShift))
}

const (
	RegisterOtghshcchar1FieldMcShift = 20
	RegisterOtghshcchar1FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar1Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldMcMask) >> RegisterOtghshcchar1FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar1Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldMcMask)|(uint32(value)<<RegisterOtghshcchar1FieldMcShift))
}

const (
	RegisterOtghshcchar1FieldDadShift = 22
	RegisterOtghshcchar1FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar1Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldDadMask) >> RegisterOtghshcchar1FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar1Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldDadMask)|(uint32(value)<<RegisterOtghshcchar1FieldDadShift))
}

const (
	RegisterOtghshcchar1FieldOddfrmShift = 29
	RegisterOtghshcchar1FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar1Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar1Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar1FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar1FieldChdisShift = 30
	RegisterOtghshcchar1FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar1Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar1Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar1FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar1FieldChenaShift = 31
	RegisterOtghshcchar1FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar1Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar1Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar1FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldChenaMask)
	}
}

// registerOtghshcsplt1Type OTG_HS host channel-1 split control register
type registerOtghshcsplt1Type uint32

const (
	RegisterOtghshcsplt1FieldPrtaddrShift = 0
	RegisterOtghshcsplt1FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt1Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt1FieldPrtaddrMask) >> RegisterOtghshcsplt1FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt1Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt1FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt1FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt1FieldHubaddrShift = 7
	RegisterOtghshcsplt1FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt1Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt1FieldHubaddrMask) >> RegisterOtghshcsplt1FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt1Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt1FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt1FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt1FieldXactposShift = 14
	RegisterOtghshcsplt1FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt1Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt1FieldXactposMask) >> RegisterOtghshcsplt1FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt1Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt1FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt1FieldXactposShift))
}

const (
	RegisterOtghshcsplt1FieldComplspltShift = 16
	RegisterOtghshcsplt1FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt1Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt1FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt1Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt1FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt1FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt1FieldSplitenShift = 31
	RegisterOtghshcsplt1FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt1Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt1FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt1Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt1FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt1FieldSplitenMask)
	}
}

// registerOtghshcint1Type OTG_HS host channel-1 interrupt register
type registerOtghshcint1Type uint32

const (
	RegisterOtghshcint1FieldXfrcShift = 0
	RegisterOtghshcint1FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint1Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint1Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint1FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint1FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint1FieldChhShift = 1
	RegisterOtghshcint1FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint1Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint1Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint1FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint1FieldChhMask)
	}
}

const (
	RegisterOtghshcint1FieldAhberrShift = 2
	RegisterOtghshcint1FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint1Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint1Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint1FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint1FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint1FieldStallShift = 3
	RegisterOtghshcint1FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint1Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint1Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint1FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint1FieldStallMask)
	}
}

const (
	RegisterOtghshcint1FieldNakShift = 4
	RegisterOtghshcint1FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint1Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint1Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint1FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint1FieldNakMask)
	}
}

const (
	RegisterOtghshcint1FieldAckShift = 5
	RegisterOtghshcint1FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint1Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint1Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint1FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint1FieldAckMask)
	}
}

const (
	RegisterOtghshcint1FieldNyetShift = 6
	RegisterOtghshcint1FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint1Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint1Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint1FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint1FieldNyetMask)
	}
}

const (
	RegisterOtghshcint1FieldTxerrShift = 7
	RegisterOtghshcint1FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint1Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint1Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint1FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint1FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint1FieldBberrShift = 8
	RegisterOtghshcint1FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint1Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint1Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint1FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint1FieldBberrMask)
	}
}

const (
	RegisterOtghshcint1FieldFrmorShift = 9
	RegisterOtghshcint1FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint1Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint1Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint1FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint1FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint1FieldDterrShift = 10
	RegisterOtghshcint1FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint1Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint1Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint1FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint1FieldDterrMask)
	}
}

// registerOtghshcintmsk1Type OTG_HS host channel-1 interrupt mask register
type registerOtghshcintmsk1Type uint32

const (
	RegisterOtghshcintmsk1FieldXfrcmShift = 0
	RegisterOtghshcintmsk1FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk1Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk1Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk1FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk1FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk1FieldChhmShift = 1
	RegisterOtghshcintmsk1FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk1Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk1Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk1FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk1FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk1FieldAhberrShift = 2
	RegisterOtghshcintmsk1FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk1Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk1Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk1FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk1FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk1FieldStallmShift = 3
	RegisterOtghshcintmsk1FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk1Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk1Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk1FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk1FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk1FieldNakmShift = 4
	RegisterOtghshcintmsk1FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk1Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk1Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk1FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk1FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk1FieldAckmShift = 5
	RegisterOtghshcintmsk1FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk1Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk1Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk1FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk1FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk1FieldNyetShift = 6
	RegisterOtghshcintmsk1FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtghshcintmsk1Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtghshcintmsk1Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk1FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk1FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk1FieldTxerrmShift = 7
	RegisterOtghshcintmsk1FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtghshcintmsk1Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtghshcintmsk1Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk1FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk1FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk1FieldBberrmShift = 8
	RegisterOtghshcintmsk1FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtghshcintmsk1Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtghshcintmsk1Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk1FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk1FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk1FieldFrmormShift = 9
	RegisterOtghshcintmsk1FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk1Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk1Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk1FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk1FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk1FieldDterrmShift = 10
	RegisterOtghshcintmsk1FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk1Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk1Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk1FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk1FieldDterrmMask)
	}
}

// registerOtghshctsiz1Type OTG_HS host channel-1 transfer size register
type registerOtghshctsiz1Type uint32

const (
	RegisterOtghshctsiz1FieldXfrsizShift = 0
	RegisterOtghshctsiz1FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz1Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz1FieldXfrsizMask) >> RegisterOtghshctsiz1FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz1Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz1FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz1FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz1FieldPktcntShift = 19
	RegisterOtghshctsiz1FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz1Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz1FieldPktcntMask) >> RegisterOtghshctsiz1FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz1Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz1FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz1FieldPktcntShift))
}

const (
	RegisterOtghshctsiz1FieldDpidShift = 29
	RegisterOtghshctsiz1FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz1Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz1FieldDpidMask) >> RegisterOtghshctsiz1FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz1Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz1FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz1FieldDpidShift))
}

// registerOtghshcdma1Type OTG_HS host channel-1 DMA address register
type registerOtghshcdma1Type uint32

const (
	RegisterOtghshcdma1FieldDmaaddrShift = 0
	RegisterOtghshcdma1FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma1Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma1FieldDmaaddrMask) >> RegisterOtghshcdma1FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma1Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma1FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma1FieldDmaaddrShift))
}

// registerOtghshcchar2Type OTG_HS host channel-2 characteristics register
type registerOtghshcchar2Type uint32

const (
	RegisterOtghshcchar2FieldMpsizShift = 0
	RegisterOtghshcchar2FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar2Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldMpsizMask) >> RegisterOtghshcchar2FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar2Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar2FieldMpsizShift))
}

const (
	RegisterOtghshcchar2FieldEpnumShift = 11
	RegisterOtghshcchar2FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar2Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldEpnumMask) >> RegisterOtghshcchar2FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar2Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar2FieldEpnumShift))
}

const (
	RegisterOtghshcchar2FieldEpdirShift = 15
	RegisterOtghshcchar2FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar2Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar2Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar2FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar2FieldLsdevShift = 17
	RegisterOtghshcchar2FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar2Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar2Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar2FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar2FieldEptypShift = 18
	RegisterOtghshcchar2FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar2Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldEptypMask) >> RegisterOtghshcchar2FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar2Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar2FieldEptypShift))
}

const (
	RegisterOtghshcchar2FieldMcShift = 20
	RegisterOtghshcchar2FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar2Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldMcMask) >> RegisterOtghshcchar2FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar2Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldMcMask)|(uint32(value)<<RegisterOtghshcchar2FieldMcShift))
}

const (
	RegisterOtghshcchar2FieldDadShift = 22
	RegisterOtghshcchar2FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar2Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldDadMask) >> RegisterOtghshcchar2FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar2Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldDadMask)|(uint32(value)<<RegisterOtghshcchar2FieldDadShift))
}

const (
	RegisterOtghshcchar2FieldOddfrmShift = 29
	RegisterOtghshcchar2FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar2Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar2Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar2FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar2FieldChdisShift = 30
	RegisterOtghshcchar2FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar2Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar2Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar2FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar2FieldChenaShift = 31
	RegisterOtghshcchar2FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar2Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar2Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar2FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldChenaMask)
	}
}

// registerOtghshcsplt2Type OTG_HS host channel-2 split control register
type registerOtghshcsplt2Type uint32

const (
	RegisterOtghshcsplt2FieldPrtaddrShift = 0
	RegisterOtghshcsplt2FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt2Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt2FieldPrtaddrMask) >> RegisterOtghshcsplt2FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt2Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt2FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt2FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt2FieldHubaddrShift = 7
	RegisterOtghshcsplt2FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt2Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt2FieldHubaddrMask) >> RegisterOtghshcsplt2FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt2Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt2FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt2FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt2FieldXactposShift = 14
	RegisterOtghshcsplt2FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt2Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt2FieldXactposMask) >> RegisterOtghshcsplt2FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt2Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt2FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt2FieldXactposShift))
}

const (
	RegisterOtghshcsplt2FieldComplspltShift = 16
	RegisterOtghshcsplt2FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt2Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt2FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt2Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt2FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt2FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt2FieldSplitenShift = 31
	RegisterOtghshcsplt2FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt2Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt2FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt2Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt2FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt2FieldSplitenMask)
	}
}

// registerOtghshcint2Type OTG_HS host channel-2 interrupt register
type registerOtghshcint2Type uint32

const (
	RegisterOtghshcint2FieldXfrcShift = 0
	RegisterOtghshcint2FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint2Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint2Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint2FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint2FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint2FieldChhShift = 1
	RegisterOtghshcint2FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint2Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint2Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint2FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint2FieldChhMask)
	}
}

const (
	RegisterOtghshcint2FieldAhberrShift = 2
	RegisterOtghshcint2FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint2Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint2Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint2FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint2FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint2FieldStallShift = 3
	RegisterOtghshcint2FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint2Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint2Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint2FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint2FieldStallMask)
	}
}

const (
	RegisterOtghshcint2FieldNakShift = 4
	RegisterOtghshcint2FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint2Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint2Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint2FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint2FieldNakMask)
	}
}

const (
	RegisterOtghshcint2FieldAckShift = 5
	RegisterOtghshcint2FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint2Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint2Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint2FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint2FieldAckMask)
	}
}

const (
	RegisterOtghshcint2FieldNyetShift = 6
	RegisterOtghshcint2FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint2Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint2Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint2FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint2FieldNyetMask)
	}
}

const (
	RegisterOtghshcint2FieldTxerrShift = 7
	RegisterOtghshcint2FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint2Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint2Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint2FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint2FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint2FieldBberrShift = 8
	RegisterOtghshcint2FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint2Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint2Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint2FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint2FieldBberrMask)
	}
}

const (
	RegisterOtghshcint2FieldFrmorShift = 9
	RegisterOtghshcint2FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint2Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint2Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint2FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint2FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint2FieldDterrShift = 10
	RegisterOtghshcint2FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint2Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint2Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint2FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint2FieldDterrMask)
	}
}

// registerOtghshcintmsk2Type OTG_HS host channel-2 interrupt mask register
type registerOtghshcintmsk2Type uint32

const (
	RegisterOtghshcintmsk2FieldXfrcmShift = 0
	RegisterOtghshcintmsk2FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk2Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk2Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk2FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk2FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk2FieldChhmShift = 1
	RegisterOtghshcintmsk2FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk2Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk2Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk2FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk2FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk2FieldAhberrShift = 2
	RegisterOtghshcintmsk2FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk2Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk2Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk2FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk2FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk2FieldStallmShift = 3
	RegisterOtghshcintmsk2FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk2Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk2Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk2FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk2FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk2FieldNakmShift = 4
	RegisterOtghshcintmsk2FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk2Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk2Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk2FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk2FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk2FieldAckmShift = 5
	RegisterOtghshcintmsk2FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk2Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk2Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk2FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk2FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk2FieldNyetShift = 6
	RegisterOtghshcintmsk2FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtghshcintmsk2Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtghshcintmsk2Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk2FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk2FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk2FieldTxerrmShift = 7
	RegisterOtghshcintmsk2FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtghshcintmsk2Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtghshcintmsk2Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk2FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk2FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk2FieldBberrmShift = 8
	RegisterOtghshcintmsk2FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtghshcintmsk2Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtghshcintmsk2Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk2FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk2FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk2FieldFrmormShift = 9
	RegisterOtghshcintmsk2FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk2Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk2Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk2FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk2FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk2FieldDterrmShift = 10
	RegisterOtghshcintmsk2FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk2Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk2Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk2FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk2FieldDterrmMask)
	}
}

// registerOtghshctsiz2Type OTG_HS host channel-2 transfer size register
type registerOtghshctsiz2Type uint32

const (
	RegisterOtghshctsiz2FieldXfrsizShift = 0
	RegisterOtghshctsiz2FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz2Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz2FieldXfrsizMask) >> RegisterOtghshctsiz2FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz2Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz2FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz2FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz2FieldPktcntShift = 19
	RegisterOtghshctsiz2FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz2Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz2FieldPktcntMask) >> RegisterOtghshctsiz2FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz2Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz2FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz2FieldPktcntShift))
}

const (
	RegisterOtghshctsiz2FieldDpidShift = 29
	RegisterOtghshctsiz2FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz2Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz2FieldDpidMask) >> RegisterOtghshctsiz2FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz2Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz2FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz2FieldDpidShift))
}

// registerOtghshcdma2Type OTG_HS host channel-2 DMA address register
type registerOtghshcdma2Type uint32

const (
	RegisterOtghshcdma2FieldDmaaddrShift = 0
	RegisterOtghshcdma2FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma2Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma2FieldDmaaddrMask) >> RegisterOtghshcdma2FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma2Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma2FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma2FieldDmaaddrShift))
}

// registerOtghshcchar3Type OTG_HS host channel-3 characteristics register
type registerOtghshcchar3Type uint32

const (
	RegisterOtghshcchar3FieldMpsizShift = 0
	RegisterOtghshcchar3FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar3Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldMpsizMask) >> RegisterOtghshcchar3FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar3Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar3FieldMpsizShift))
}

const (
	RegisterOtghshcchar3FieldEpnumShift = 11
	RegisterOtghshcchar3FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar3Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldEpnumMask) >> RegisterOtghshcchar3FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar3Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar3FieldEpnumShift))
}

const (
	RegisterOtghshcchar3FieldEpdirShift = 15
	RegisterOtghshcchar3FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar3Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar3Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar3FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar3FieldLsdevShift = 17
	RegisterOtghshcchar3FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar3Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar3Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar3FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar3FieldEptypShift = 18
	RegisterOtghshcchar3FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar3Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldEptypMask) >> RegisterOtghshcchar3FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar3Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar3FieldEptypShift))
}

const (
	RegisterOtghshcchar3FieldMcShift = 20
	RegisterOtghshcchar3FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar3Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldMcMask) >> RegisterOtghshcchar3FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar3Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldMcMask)|(uint32(value)<<RegisterOtghshcchar3FieldMcShift))
}

const (
	RegisterOtghshcchar3FieldDadShift = 22
	RegisterOtghshcchar3FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar3Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldDadMask) >> RegisterOtghshcchar3FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar3Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldDadMask)|(uint32(value)<<RegisterOtghshcchar3FieldDadShift))
}

const (
	RegisterOtghshcchar3FieldOddfrmShift = 29
	RegisterOtghshcchar3FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar3Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar3Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar3FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar3FieldChdisShift = 30
	RegisterOtghshcchar3FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar3Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar3Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar3FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar3FieldChenaShift = 31
	RegisterOtghshcchar3FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar3Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar3Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar3FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldChenaMask)
	}
}

// registerOtghshcsplt3Type OTG_HS host channel-3 split control register
type registerOtghshcsplt3Type uint32

const (
	RegisterOtghshcsplt3FieldPrtaddrShift = 0
	RegisterOtghshcsplt3FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt3Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt3FieldPrtaddrMask) >> RegisterOtghshcsplt3FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt3Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt3FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt3FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt3FieldHubaddrShift = 7
	RegisterOtghshcsplt3FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt3Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt3FieldHubaddrMask) >> RegisterOtghshcsplt3FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt3Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt3FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt3FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt3FieldXactposShift = 14
	RegisterOtghshcsplt3FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt3Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt3FieldXactposMask) >> RegisterOtghshcsplt3FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt3Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt3FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt3FieldXactposShift))
}

const (
	RegisterOtghshcsplt3FieldComplspltShift = 16
	RegisterOtghshcsplt3FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt3Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt3FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt3Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt3FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt3FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt3FieldSplitenShift = 31
	RegisterOtghshcsplt3FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt3Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt3FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt3Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt3FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt3FieldSplitenMask)
	}
}

// registerOtghshcint3Type OTG_HS host channel-3 interrupt register
type registerOtghshcint3Type uint32

const (
	RegisterOtghshcint3FieldXfrcShift = 0
	RegisterOtghshcint3FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint3Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint3Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint3FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint3FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint3FieldChhShift = 1
	RegisterOtghshcint3FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint3Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint3Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint3FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint3FieldChhMask)
	}
}

const (
	RegisterOtghshcint3FieldAhberrShift = 2
	RegisterOtghshcint3FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint3Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint3Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint3FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint3FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint3FieldStallShift = 3
	RegisterOtghshcint3FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint3Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint3Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint3FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint3FieldStallMask)
	}
}

const (
	RegisterOtghshcint3FieldNakShift = 4
	RegisterOtghshcint3FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint3Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint3Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint3FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint3FieldNakMask)
	}
}

const (
	RegisterOtghshcint3FieldAckShift = 5
	RegisterOtghshcint3FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint3Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint3Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint3FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint3FieldAckMask)
	}
}

const (
	RegisterOtghshcint3FieldNyetShift = 6
	RegisterOtghshcint3FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint3Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint3Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint3FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint3FieldNyetMask)
	}
}

const (
	RegisterOtghshcint3FieldTxerrShift = 7
	RegisterOtghshcint3FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint3Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint3Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint3FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint3FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint3FieldBberrShift = 8
	RegisterOtghshcint3FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint3Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint3Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint3FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint3FieldBberrMask)
	}
}

const (
	RegisterOtghshcint3FieldFrmorShift = 9
	RegisterOtghshcint3FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint3Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint3Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint3FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint3FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint3FieldDterrShift = 10
	RegisterOtghshcint3FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint3Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint3Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint3FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint3FieldDterrMask)
	}
}

// registerOtghshcintmsk3Type OTG_HS host channel-3 interrupt mask register
type registerOtghshcintmsk3Type uint32

const (
	RegisterOtghshcintmsk3FieldXfrcmShift = 0
	RegisterOtghshcintmsk3FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk3Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk3Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk3FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk3FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk3FieldChhmShift = 1
	RegisterOtghshcintmsk3FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk3Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk3Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk3FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk3FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk3FieldAhberrShift = 2
	RegisterOtghshcintmsk3FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk3Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk3Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk3FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk3FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk3FieldStallmShift = 3
	RegisterOtghshcintmsk3FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk3Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk3Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk3FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk3FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk3FieldNakmShift = 4
	RegisterOtghshcintmsk3FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk3Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk3Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk3FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk3FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk3FieldAckmShift = 5
	RegisterOtghshcintmsk3FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk3Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk3Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk3FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk3FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk3FieldNyetShift = 6
	RegisterOtghshcintmsk3FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtghshcintmsk3Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtghshcintmsk3Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk3FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk3FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk3FieldTxerrmShift = 7
	RegisterOtghshcintmsk3FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtghshcintmsk3Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtghshcintmsk3Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk3FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk3FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk3FieldBberrmShift = 8
	RegisterOtghshcintmsk3FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtghshcintmsk3Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtghshcintmsk3Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk3FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk3FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk3FieldFrmormShift = 9
	RegisterOtghshcintmsk3FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk3Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk3Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk3FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk3FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk3FieldDterrmShift = 10
	RegisterOtghshcintmsk3FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk3Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk3Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk3FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk3FieldDterrmMask)
	}
}

// registerOtghshctsiz3Type OTG_HS host channel-3 transfer size register
type registerOtghshctsiz3Type uint32

const (
	RegisterOtghshctsiz3FieldXfrsizShift = 0
	RegisterOtghshctsiz3FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz3Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz3FieldXfrsizMask) >> RegisterOtghshctsiz3FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz3Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz3FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz3FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz3FieldPktcntShift = 19
	RegisterOtghshctsiz3FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz3Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz3FieldPktcntMask) >> RegisterOtghshctsiz3FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz3Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz3FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz3FieldPktcntShift))
}

const (
	RegisterOtghshctsiz3FieldDpidShift = 29
	RegisterOtghshctsiz3FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz3Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz3FieldDpidMask) >> RegisterOtghshctsiz3FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz3Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz3FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz3FieldDpidShift))
}

// registerOtghshcdma3Type OTG_HS host channel-3 DMA address register
type registerOtghshcdma3Type uint32

const (
	RegisterOtghshcdma3FieldDmaaddrShift = 0
	RegisterOtghshcdma3FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma3Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma3FieldDmaaddrMask) >> RegisterOtghshcdma3FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma3Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma3FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma3FieldDmaaddrShift))
}

// registerOtghshcchar4Type OTG_HS host channel-4 characteristics register
type registerOtghshcchar4Type uint32

const (
	RegisterOtghshcchar4FieldMpsizShift = 0
	RegisterOtghshcchar4FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar4Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldMpsizMask) >> RegisterOtghshcchar4FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar4Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar4FieldMpsizShift))
}

const (
	RegisterOtghshcchar4FieldEpnumShift = 11
	RegisterOtghshcchar4FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar4Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldEpnumMask) >> RegisterOtghshcchar4FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar4Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar4FieldEpnumShift))
}

const (
	RegisterOtghshcchar4FieldEpdirShift = 15
	RegisterOtghshcchar4FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar4Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar4Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar4FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar4FieldLsdevShift = 17
	RegisterOtghshcchar4FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar4Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar4Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar4FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar4FieldEptypShift = 18
	RegisterOtghshcchar4FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar4Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldEptypMask) >> RegisterOtghshcchar4FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar4Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar4FieldEptypShift))
}

const (
	RegisterOtghshcchar4FieldMcShift = 20
	RegisterOtghshcchar4FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar4Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldMcMask) >> RegisterOtghshcchar4FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar4Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldMcMask)|(uint32(value)<<RegisterOtghshcchar4FieldMcShift))
}

const (
	RegisterOtghshcchar4FieldDadShift = 22
	RegisterOtghshcchar4FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar4Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldDadMask) >> RegisterOtghshcchar4FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar4Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldDadMask)|(uint32(value)<<RegisterOtghshcchar4FieldDadShift))
}

const (
	RegisterOtghshcchar4FieldOddfrmShift = 29
	RegisterOtghshcchar4FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar4Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar4Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar4FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar4FieldChdisShift = 30
	RegisterOtghshcchar4FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar4Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar4Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar4FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar4FieldChenaShift = 31
	RegisterOtghshcchar4FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar4Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar4Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar4FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldChenaMask)
	}
}

// registerOtghshcsplt4Type OTG_HS host channel-4 split control register
type registerOtghshcsplt4Type uint32

const (
	RegisterOtghshcsplt4FieldPrtaddrShift = 0
	RegisterOtghshcsplt4FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt4Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt4FieldPrtaddrMask) >> RegisterOtghshcsplt4FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt4Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt4FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt4FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt4FieldHubaddrShift = 7
	RegisterOtghshcsplt4FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt4Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt4FieldHubaddrMask) >> RegisterOtghshcsplt4FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt4Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt4FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt4FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt4FieldXactposShift = 14
	RegisterOtghshcsplt4FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt4Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt4FieldXactposMask) >> RegisterOtghshcsplt4FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt4Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt4FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt4FieldXactposShift))
}

const (
	RegisterOtghshcsplt4FieldComplspltShift = 16
	RegisterOtghshcsplt4FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt4Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt4FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt4Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt4FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt4FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt4FieldSplitenShift = 31
	RegisterOtghshcsplt4FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt4Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt4FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt4Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt4FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt4FieldSplitenMask)
	}
}

// registerOtghshcint4Type OTG_HS host channel-4 interrupt register
type registerOtghshcint4Type uint32

const (
	RegisterOtghshcint4FieldXfrcShift = 0
	RegisterOtghshcint4FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint4Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint4Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint4FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint4FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint4FieldChhShift = 1
	RegisterOtghshcint4FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint4Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint4Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint4FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint4FieldChhMask)
	}
}

const (
	RegisterOtghshcint4FieldAhberrShift = 2
	RegisterOtghshcint4FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint4Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint4Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint4FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint4FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint4FieldStallShift = 3
	RegisterOtghshcint4FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint4Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint4Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint4FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint4FieldStallMask)
	}
}

const (
	RegisterOtghshcint4FieldNakShift = 4
	RegisterOtghshcint4FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint4Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint4Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint4FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint4FieldNakMask)
	}
}

const (
	RegisterOtghshcint4FieldAckShift = 5
	RegisterOtghshcint4FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint4Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint4Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint4FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint4FieldAckMask)
	}
}

const (
	RegisterOtghshcint4FieldNyetShift = 6
	RegisterOtghshcint4FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint4Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint4Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint4FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint4FieldNyetMask)
	}
}

const (
	RegisterOtghshcint4FieldTxerrShift = 7
	RegisterOtghshcint4FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint4Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint4Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint4FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint4FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint4FieldBberrShift = 8
	RegisterOtghshcint4FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint4Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint4Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint4FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint4FieldBberrMask)
	}
}

const (
	RegisterOtghshcint4FieldFrmorShift = 9
	RegisterOtghshcint4FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint4Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint4Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint4FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint4FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint4FieldDterrShift = 10
	RegisterOtghshcint4FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint4Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint4Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint4FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint4FieldDterrMask)
	}
}

// registerOtghshcintmsk4Type OTG_HS host channel-4 interrupt mask register
type registerOtghshcintmsk4Type uint32

const (
	RegisterOtghshcintmsk4FieldXfrcmShift = 0
	RegisterOtghshcintmsk4FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk4Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk4Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk4FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk4FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk4FieldChhmShift = 1
	RegisterOtghshcintmsk4FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk4Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk4Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk4FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk4FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk4FieldAhberrShift = 2
	RegisterOtghshcintmsk4FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk4Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk4Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk4FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk4FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk4FieldStallmShift = 3
	RegisterOtghshcintmsk4FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk4Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk4Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk4FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk4FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk4FieldNakmShift = 4
	RegisterOtghshcintmsk4FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk4Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk4Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk4FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk4FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk4FieldAckmShift = 5
	RegisterOtghshcintmsk4FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk4Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk4Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk4FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk4FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk4FieldNyetShift = 6
	RegisterOtghshcintmsk4FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtghshcintmsk4Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtghshcintmsk4Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk4FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk4FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk4FieldTxerrmShift = 7
	RegisterOtghshcintmsk4FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtghshcintmsk4Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtghshcintmsk4Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk4FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk4FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk4FieldBberrmShift = 8
	RegisterOtghshcintmsk4FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtghshcintmsk4Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtghshcintmsk4Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk4FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk4FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk4FieldFrmormShift = 9
	RegisterOtghshcintmsk4FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk4Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk4Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk4FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk4FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk4FieldDterrmShift = 10
	RegisterOtghshcintmsk4FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk4Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk4Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk4FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk4FieldDterrmMask)
	}
}

// registerOtghshctsiz4Type OTG_HS host channel-4 transfer size register
type registerOtghshctsiz4Type uint32

const (
	RegisterOtghshctsiz4FieldXfrsizShift = 0
	RegisterOtghshctsiz4FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz4Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz4FieldXfrsizMask) >> RegisterOtghshctsiz4FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz4Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz4FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz4FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz4FieldPktcntShift = 19
	RegisterOtghshctsiz4FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz4Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz4FieldPktcntMask) >> RegisterOtghshctsiz4FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz4Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz4FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz4FieldPktcntShift))
}

const (
	RegisterOtghshctsiz4FieldDpidShift = 29
	RegisterOtghshctsiz4FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz4Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz4FieldDpidMask) >> RegisterOtghshctsiz4FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz4Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz4FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz4FieldDpidShift))
}

// registerOtghshcdma4Type OTG_HS host channel-4 DMA address register
type registerOtghshcdma4Type uint32

const (
	RegisterOtghshcdma4FieldDmaaddrShift = 0
	RegisterOtghshcdma4FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma4Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma4FieldDmaaddrMask) >> RegisterOtghshcdma4FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma4Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma4FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma4FieldDmaaddrShift))
}

// registerOtghshcchar5Type OTG_HS host channel-5 characteristics register
type registerOtghshcchar5Type uint32

const (
	RegisterOtghshcchar5FieldMpsizShift = 0
	RegisterOtghshcchar5FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar5Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldMpsizMask) >> RegisterOtghshcchar5FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar5Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar5FieldMpsizShift))
}

const (
	RegisterOtghshcchar5FieldEpnumShift = 11
	RegisterOtghshcchar5FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar5Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldEpnumMask) >> RegisterOtghshcchar5FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar5Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar5FieldEpnumShift))
}

const (
	RegisterOtghshcchar5FieldEpdirShift = 15
	RegisterOtghshcchar5FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar5Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar5Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar5FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar5FieldLsdevShift = 17
	RegisterOtghshcchar5FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar5Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar5Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar5FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar5FieldEptypShift = 18
	RegisterOtghshcchar5FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar5Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldEptypMask) >> RegisterOtghshcchar5FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar5Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar5FieldEptypShift))
}

const (
	RegisterOtghshcchar5FieldMcShift = 20
	RegisterOtghshcchar5FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar5Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldMcMask) >> RegisterOtghshcchar5FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar5Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldMcMask)|(uint32(value)<<RegisterOtghshcchar5FieldMcShift))
}

const (
	RegisterOtghshcchar5FieldDadShift = 22
	RegisterOtghshcchar5FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar5Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldDadMask) >> RegisterOtghshcchar5FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar5Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldDadMask)|(uint32(value)<<RegisterOtghshcchar5FieldDadShift))
}

const (
	RegisterOtghshcchar5FieldOddfrmShift = 29
	RegisterOtghshcchar5FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar5Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar5Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar5FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar5FieldChdisShift = 30
	RegisterOtghshcchar5FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar5Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar5Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar5FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar5FieldChenaShift = 31
	RegisterOtghshcchar5FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar5Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar5Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar5FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldChenaMask)
	}
}

// registerOtghshcsplt5Type OTG_HS host channel-5 split control register
type registerOtghshcsplt5Type uint32

const (
	RegisterOtghshcsplt5FieldPrtaddrShift = 0
	RegisterOtghshcsplt5FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt5Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt5FieldPrtaddrMask) >> RegisterOtghshcsplt5FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt5Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt5FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt5FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt5FieldHubaddrShift = 7
	RegisterOtghshcsplt5FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt5Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt5FieldHubaddrMask) >> RegisterOtghshcsplt5FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt5Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt5FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt5FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt5FieldXactposShift = 14
	RegisterOtghshcsplt5FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt5Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt5FieldXactposMask) >> RegisterOtghshcsplt5FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt5Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt5FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt5FieldXactposShift))
}

const (
	RegisterOtghshcsplt5FieldComplspltShift = 16
	RegisterOtghshcsplt5FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt5Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt5FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt5Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt5FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt5FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt5FieldSplitenShift = 31
	RegisterOtghshcsplt5FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt5Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt5FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt5Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt5FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt5FieldSplitenMask)
	}
}

// registerOtghshcint5Type OTG_HS host channel-5 interrupt register
type registerOtghshcint5Type uint32

const (
	RegisterOtghshcint5FieldXfrcShift = 0
	RegisterOtghshcint5FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint5Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint5Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint5FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint5FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint5FieldChhShift = 1
	RegisterOtghshcint5FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint5Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint5Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint5FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint5FieldChhMask)
	}
}

const (
	RegisterOtghshcint5FieldAhberrShift = 2
	RegisterOtghshcint5FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint5Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint5Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint5FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint5FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint5FieldStallShift = 3
	RegisterOtghshcint5FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint5Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint5Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint5FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint5FieldStallMask)
	}
}

const (
	RegisterOtghshcint5FieldNakShift = 4
	RegisterOtghshcint5FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint5Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint5Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint5FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint5FieldNakMask)
	}
}

const (
	RegisterOtghshcint5FieldAckShift = 5
	RegisterOtghshcint5FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint5Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint5Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint5FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint5FieldAckMask)
	}
}

const (
	RegisterOtghshcint5FieldNyetShift = 6
	RegisterOtghshcint5FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint5Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint5Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint5FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint5FieldNyetMask)
	}
}

const (
	RegisterOtghshcint5FieldTxerrShift = 7
	RegisterOtghshcint5FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint5Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint5Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint5FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint5FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint5FieldBberrShift = 8
	RegisterOtghshcint5FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint5Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint5Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint5FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint5FieldBberrMask)
	}
}

const (
	RegisterOtghshcint5FieldFrmorShift = 9
	RegisterOtghshcint5FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint5Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint5Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint5FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint5FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint5FieldDterrShift = 10
	RegisterOtghshcint5FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint5Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint5Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint5FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint5FieldDterrMask)
	}
}

// registerOtghshcintmsk5Type OTG_HS host channel-5 interrupt mask register
type registerOtghshcintmsk5Type uint32

const (
	RegisterOtghshcintmsk5FieldXfrcmShift = 0
	RegisterOtghshcintmsk5FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk5Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk5Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk5FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk5FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk5FieldChhmShift = 1
	RegisterOtghshcintmsk5FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk5Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk5Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk5FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk5FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk5FieldAhberrShift = 2
	RegisterOtghshcintmsk5FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk5Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk5Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk5FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk5FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk5FieldStallmShift = 3
	RegisterOtghshcintmsk5FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk5Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk5Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk5FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk5FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk5FieldNakmShift = 4
	RegisterOtghshcintmsk5FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk5Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk5Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk5FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk5FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk5FieldAckmShift = 5
	RegisterOtghshcintmsk5FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk5Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk5Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk5FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk5FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk5FieldNyetShift = 6
	RegisterOtghshcintmsk5FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtghshcintmsk5Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtghshcintmsk5Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk5FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk5FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk5FieldTxerrmShift = 7
	RegisterOtghshcintmsk5FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtghshcintmsk5Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtghshcintmsk5Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk5FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk5FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk5FieldBberrmShift = 8
	RegisterOtghshcintmsk5FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtghshcintmsk5Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtghshcintmsk5Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk5FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk5FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk5FieldFrmormShift = 9
	RegisterOtghshcintmsk5FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk5Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk5Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk5FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk5FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk5FieldDterrmShift = 10
	RegisterOtghshcintmsk5FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk5Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk5Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk5FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk5FieldDterrmMask)
	}
}

// registerOtghshctsiz5Type OTG_HS host channel-5 transfer size register
type registerOtghshctsiz5Type uint32

const (
	RegisterOtghshctsiz5FieldXfrsizShift = 0
	RegisterOtghshctsiz5FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz5Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz5FieldXfrsizMask) >> RegisterOtghshctsiz5FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz5Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz5FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz5FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz5FieldPktcntShift = 19
	RegisterOtghshctsiz5FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz5Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz5FieldPktcntMask) >> RegisterOtghshctsiz5FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz5Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz5FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz5FieldPktcntShift))
}

const (
	RegisterOtghshctsiz5FieldDpidShift = 29
	RegisterOtghshctsiz5FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz5Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz5FieldDpidMask) >> RegisterOtghshctsiz5FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz5Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz5FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz5FieldDpidShift))
}

// registerOtghshcdma5Type OTG_HS host channel-5 DMA address register
type registerOtghshcdma5Type uint32

const (
	RegisterOtghshcdma5FieldDmaaddrShift = 0
	RegisterOtghshcdma5FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma5Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma5FieldDmaaddrMask) >> RegisterOtghshcdma5FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma5Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma5FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma5FieldDmaaddrShift))
}

// registerOtghshcchar6Type OTG_HS host channel-6 characteristics register
type registerOtghshcchar6Type uint32

const (
	RegisterOtghshcchar6FieldMpsizShift = 0
	RegisterOtghshcchar6FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar6Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldMpsizMask) >> RegisterOtghshcchar6FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar6Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar6FieldMpsizShift))
}

const (
	RegisterOtghshcchar6FieldEpnumShift = 11
	RegisterOtghshcchar6FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar6Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldEpnumMask) >> RegisterOtghshcchar6FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar6Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar6FieldEpnumShift))
}

const (
	RegisterOtghshcchar6FieldEpdirShift = 15
	RegisterOtghshcchar6FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar6Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar6Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar6FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar6FieldLsdevShift = 17
	RegisterOtghshcchar6FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar6Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar6Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar6FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar6FieldEptypShift = 18
	RegisterOtghshcchar6FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar6Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldEptypMask) >> RegisterOtghshcchar6FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar6Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar6FieldEptypShift))
}

const (
	RegisterOtghshcchar6FieldMcShift = 20
	RegisterOtghshcchar6FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar6Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldMcMask) >> RegisterOtghshcchar6FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar6Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldMcMask)|(uint32(value)<<RegisterOtghshcchar6FieldMcShift))
}

const (
	RegisterOtghshcchar6FieldDadShift = 22
	RegisterOtghshcchar6FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar6Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldDadMask) >> RegisterOtghshcchar6FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar6Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldDadMask)|(uint32(value)<<RegisterOtghshcchar6FieldDadShift))
}

const (
	RegisterOtghshcchar6FieldOddfrmShift = 29
	RegisterOtghshcchar6FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar6Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar6Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar6FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar6FieldChdisShift = 30
	RegisterOtghshcchar6FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar6Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar6Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar6FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar6FieldChenaShift = 31
	RegisterOtghshcchar6FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar6Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar6Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar6FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldChenaMask)
	}
}

// registerOtghshcsplt6Type OTG_HS host channel-6 split control register
type registerOtghshcsplt6Type uint32

const (
	RegisterOtghshcsplt6FieldPrtaddrShift = 0
	RegisterOtghshcsplt6FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt6Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt6FieldPrtaddrMask) >> RegisterOtghshcsplt6FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt6Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt6FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt6FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt6FieldHubaddrShift = 7
	RegisterOtghshcsplt6FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt6Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt6FieldHubaddrMask) >> RegisterOtghshcsplt6FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt6Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt6FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt6FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt6FieldXactposShift = 14
	RegisterOtghshcsplt6FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt6Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt6FieldXactposMask) >> RegisterOtghshcsplt6FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt6Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt6FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt6FieldXactposShift))
}

const (
	RegisterOtghshcsplt6FieldComplspltShift = 16
	RegisterOtghshcsplt6FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt6Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt6FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt6Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt6FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt6FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt6FieldSplitenShift = 31
	RegisterOtghshcsplt6FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt6Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt6FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt6Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt6FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt6FieldSplitenMask)
	}
}

// registerOtghshcint6Type OTG_HS host channel-6 interrupt register
type registerOtghshcint6Type uint32

const (
	RegisterOtghshcint6FieldXfrcShift = 0
	RegisterOtghshcint6FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint6Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint6Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint6FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint6FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint6FieldChhShift = 1
	RegisterOtghshcint6FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint6Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint6Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint6FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint6FieldChhMask)
	}
}

const (
	RegisterOtghshcint6FieldAhberrShift = 2
	RegisterOtghshcint6FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint6Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint6Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint6FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint6FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint6FieldStallShift = 3
	RegisterOtghshcint6FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint6Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint6Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint6FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint6FieldStallMask)
	}
}

const (
	RegisterOtghshcint6FieldNakShift = 4
	RegisterOtghshcint6FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint6Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint6Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint6FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint6FieldNakMask)
	}
}

const (
	RegisterOtghshcint6FieldAckShift = 5
	RegisterOtghshcint6FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint6Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint6Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint6FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint6FieldAckMask)
	}
}

const (
	RegisterOtghshcint6FieldNyetShift = 6
	RegisterOtghshcint6FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint6Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint6Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint6FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint6FieldNyetMask)
	}
}

const (
	RegisterOtghshcint6FieldTxerrShift = 7
	RegisterOtghshcint6FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint6Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint6Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint6FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint6FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint6FieldBberrShift = 8
	RegisterOtghshcint6FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint6Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint6Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint6FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint6FieldBberrMask)
	}
}

const (
	RegisterOtghshcint6FieldFrmorShift = 9
	RegisterOtghshcint6FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint6Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint6Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint6FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint6FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint6FieldDterrShift = 10
	RegisterOtghshcint6FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint6Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint6Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint6FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint6FieldDterrMask)
	}
}

// registerOtghshcintmsk6Type OTG_HS host channel-6 interrupt mask register
type registerOtghshcintmsk6Type uint32

const (
	RegisterOtghshcintmsk6FieldXfrcmShift = 0
	RegisterOtghshcintmsk6FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk6Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk6Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk6FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk6FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk6FieldChhmShift = 1
	RegisterOtghshcintmsk6FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk6Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk6Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk6FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk6FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk6FieldAhberrShift = 2
	RegisterOtghshcintmsk6FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk6Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk6Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk6FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk6FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk6FieldStallmShift = 3
	RegisterOtghshcintmsk6FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk6Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk6Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk6FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk6FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk6FieldNakmShift = 4
	RegisterOtghshcintmsk6FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk6Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk6Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk6FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk6FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk6FieldAckmShift = 5
	RegisterOtghshcintmsk6FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk6Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk6Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk6FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk6FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk6FieldNyetShift = 6
	RegisterOtghshcintmsk6FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtghshcintmsk6Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtghshcintmsk6Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk6FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk6FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk6FieldTxerrmShift = 7
	RegisterOtghshcintmsk6FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtghshcintmsk6Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtghshcintmsk6Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk6FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk6FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk6FieldBberrmShift = 8
	RegisterOtghshcintmsk6FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtghshcintmsk6Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtghshcintmsk6Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk6FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk6FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk6FieldFrmormShift = 9
	RegisterOtghshcintmsk6FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk6Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk6Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk6FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk6FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk6FieldDterrmShift = 10
	RegisterOtghshcintmsk6FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk6Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk6Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk6FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk6FieldDterrmMask)
	}
}

// registerOtghshctsiz6Type OTG_HS host channel-6 transfer size register
type registerOtghshctsiz6Type uint32

const (
	RegisterOtghshctsiz6FieldXfrsizShift = 0
	RegisterOtghshctsiz6FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz6Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz6FieldXfrsizMask) >> RegisterOtghshctsiz6FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz6Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz6FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz6FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz6FieldPktcntShift = 19
	RegisterOtghshctsiz6FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz6Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz6FieldPktcntMask) >> RegisterOtghshctsiz6FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz6Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz6FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz6FieldPktcntShift))
}

const (
	RegisterOtghshctsiz6FieldDpidShift = 29
	RegisterOtghshctsiz6FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz6Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz6FieldDpidMask) >> RegisterOtghshctsiz6FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz6Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz6FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz6FieldDpidShift))
}

// registerOtghshcdma6Type OTG_HS host channel-6 DMA address register
type registerOtghshcdma6Type uint32

const (
	RegisterOtghshcdma6FieldDmaaddrShift = 0
	RegisterOtghshcdma6FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma6Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma6FieldDmaaddrMask) >> RegisterOtghshcdma6FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma6Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma6FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma6FieldDmaaddrShift))
}

// registerOtghshcchar7Type OTG_HS host channel-7 characteristics register
type registerOtghshcchar7Type uint32

const (
	RegisterOtghshcchar7FieldMpsizShift = 0
	RegisterOtghshcchar7FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar7Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldMpsizMask) >> RegisterOtghshcchar7FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar7Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar7FieldMpsizShift))
}

const (
	RegisterOtghshcchar7FieldEpnumShift = 11
	RegisterOtghshcchar7FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar7Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldEpnumMask) >> RegisterOtghshcchar7FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar7Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar7FieldEpnumShift))
}

const (
	RegisterOtghshcchar7FieldEpdirShift = 15
	RegisterOtghshcchar7FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar7Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar7Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar7FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar7FieldLsdevShift = 17
	RegisterOtghshcchar7FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar7Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar7Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar7FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar7FieldEptypShift = 18
	RegisterOtghshcchar7FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar7Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldEptypMask) >> RegisterOtghshcchar7FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar7Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar7FieldEptypShift))
}

const (
	RegisterOtghshcchar7FieldMcShift = 20
	RegisterOtghshcchar7FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar7Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldMcMask) >> RegisterOtghshcchar7FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar7Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldMcMask)|(uint32(value)<<RegisterOtghshcchar7FieldMcShift))
}

const (
	RegisterOtghshcchar7FieldDadShift = 22
	RegisterOtghshcchar7FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar7Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldDadMask) >> RegisterOtghshcchar7FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar7Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldDadMask)|(uint32(value)<<RegisterOtghshcchar7FieldDadShift))
}

const (
	RegisterOtghshcchar7FieldOddfrmShift = 29
	RegisterOtghshcchar7FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar7Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar7Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar7FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar7FieldChdisShift = 30
	RegisterOtghshcchar7FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar7Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar7Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar7FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar7FieldChenaShift = 31
	RegisterOtghshcchar7FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar7Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar7Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar7FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldChenaMask)
	}
}

// registerOtghshcsplt7Type OTG_HS host channel-7 split control register
type registerOtghshcsplt7Type uint32

const (
	RegisterOtghshcsplt7FieldPrtaddrShift = 0
	RegisterOtghshcsplt7FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt7Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt7FieldPrtaddrMask) >> RegisterOtghshcsplt7FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt7Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt7FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt7FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt7FieldHubaddrShift = 7
	RegisterOtghshcsplt7FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt7Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt7FieldHubaddrMask) >> RegisterOtghshcsplt7FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt7Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt7FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt7FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt7FieldXactposShift = 14
	RegisterOtghshcsplt7FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt7Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt7FieldXactposMask) >> RegisterOtghshcsplt7FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt7Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt7FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt7FieldXactposShift))
}

const (
	RegisterOtghshcsplt7FieldComplspltShift = 16
	RegisterOtghshcsplt7FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt7Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt7FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt7Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt7FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt7FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt7FieldSplitenShift = 31
	RegisterOtghshcsplt7FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt7Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt7FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt7Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt7FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt7FieldSplitenMask)
	}
}

// registerOtghshcint7Type OTG_HS host channel-7 interrupt register
type registerOtghshcint7Type uint32

const (
	RegisterOtghshcint7FieldXfrcShift = 0
	RegisterOtghshcint7FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint7Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint7Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint7FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint7FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint7FieldChhShift = 1
	RegisterOtghshcint7FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint7Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint7Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint7FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint7FieldChhMask)
	}
}

const (
	RegisterOtghshcint7FieldAhberrShift = 2
	RegisterOtghshcint7FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint7Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint7Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint7FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint7FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint7FieldStallShift = 3
	RegisterOtghshcint7FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint7Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint7Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint7FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint7FieldStallMask)
	}
}

const (
	RegisterOtghshcint7FieldNakShift = 4
	RegisterOtghshcint7FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint7Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint7Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint7FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint7FieldNakMask)
	}
}

const (
	RegisterOtghshcint7FieldAckShift = 5
	RegisterOtghshcint7FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint7Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint7Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint7FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint7FieldAckMask)
	}
}

const (
	RegisterOtghshcint7FieldNyetShift = 6
	RegisterOtghshcint7FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint7Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint7Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint7FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint7FieldNyetMask)
	}
}

const (
	RegisterOtghshcint7FieldTxerrShift = 7
	RegisterOtghshcint7FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint7Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint7Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint7FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint7FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint7FieldBberrShift = 8
	RegisterOtghshcint7FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint7Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint7Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint7FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint7FieldBberrMask)
	}
}

const (
	RegisterOtghshcint7FieldFrmorShift = 9
	RegisterOtghshcint7FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint7Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint7Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint7FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint7FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint7FieldDterrShift = 10
	RegisterOtghshcint7FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint7Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint7Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint7FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint7FieldDterrMask)
	}
}

// registerOtghshcintmsk7Type OTG_HS host channel-7 interrupt mask register
type registerOtghshcintmsk7Type uint32

const (
	RegisterOtghshcintmsk7FieldXfrcmShift = 0
	RegisterOtghshcintmsk7FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk7Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk7Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk7FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk7FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk7FieldChhmShift = 1
	RegisterOtghshcintmsk7FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk7Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk7Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk7FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk7FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk7FieldAhberrShift = 2
	RegisterOtghshcintmsk7FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk7Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk7Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk7FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk7FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk7FieldStallmShift = 3
	RegisterOtghshcintmsk7FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk7Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk7Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk7FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk7FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk7FieldNakmShift = 4
	RegisterOtghshcintmsk7FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk7Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk7Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk7FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk7FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk7FieldAckmShift = 5
	RegisterOtghshcintmsk7FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk7Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk7Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk7FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk7FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk7FieldNyetShift = 6
	RegisterOtghshcintmsk7FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtghshcintmsk7Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtghshcintmsk7Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk7FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk7FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk7FieldTxerrmShift = 7
	RegisterOtghshcintmsk7FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtghshcintmsk7Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtghshcintmsk7Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk7FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk7FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk7FieldBberrmShift = 8
	RegisterOtghshcintmsk7FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtghshcintmsk7Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtghshcintmsk7Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk7FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk7FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk7FieldFrmormShift = 9
	RegisterOtghshcintmsk7FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk7Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk7Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk7FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk7FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk7FieldDterrmShift = 10
	RegisterOtghshcintmsk7FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk7Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk7Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk7FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk7FieldDterrmMask)
	}
}

// registerOtghshctsiz7Type OTG_HS host channel-7 transfer size register
type registerOtghshctsiz7Type uint32

const (
	RegisterOtghshctsiz7FieldXfrsizShift = 0
	RegisterOtghshctsiz7FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz7Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz7FieldXfrsizMask) >> RegisterOtghshctsiz7FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz7Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz7FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz7FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz7FieldPktcntShift = 19
	RegisterOtghshctsiz7FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz7Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz7FieldPktcntMask) >> RegisterOtghshctsiz7FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz7Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz7FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz7FieldPktcntShift))
}

const (
	RegisterOtghshctsiz7FieldDpidShift = 29
	RegisterOtghshctsiz7FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz7Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz7FieldDpidMask) >> RegisterOtghshctsiz7FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz7Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz7FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz7FieldDpidShift))
}

// registerOtghshcdma7Type OTG_HS host channel-7 DMA address register
type registerOtghshcdma7Type uint32

const (
	RegisterOtghshcdma7FieldDmaaddrShift = 0
	RegisterOtghshcdma7FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma7Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma7FieldDmaaddrMask) >> RegisterOtghshcdma7FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma7Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma7FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma7FieldDmaaddrShift))
}

// registerOtghshcchar8Type OTG_HS host channel-8 characteristics register
type registerOtghshcchar8Type uint32

const (
	RegisterOtghshcchar8FieldMpsizShift = 0
	RegisterOtghshcchar8FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar8Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldMpsizMask) >> RegisterOtghshcchar8FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar8Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar8FieldMpsizShift))
}

const (
	RegisterOtghshcchar8FieldEpnumShift = 11
	RegisterOtghshcchar8FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar8Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldEpnumMask) >> RegisterOtghshcchar8FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar8Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar8FieldEpnumShift))
}

const (
	RegisterOtghshcchar8FieldEpdirShift = 15
	RegisterOtghshcchar8FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar8Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar8Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar8FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar8FieldLsdevShift = 17
	RegisterOtghshcchar8FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar8Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar8Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar8FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar8FieldEptypShift = 18
	RegisterOtghshcchar8FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar8Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldEptypMask) >> RegisterOtghshcchar8FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar8Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar8FieldEptypShift))
}

const (
	RegisterOtghshcchar8FieldMcShift = 20
	RegisterOtghshcchar8FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar8Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldMcMask) >> RegisterOtghshcchar8FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar8Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldMcMask)|(uint32(value)<<RegisterOtghshcchar8FieldMcShift))
}

const (
	RegisterOtghshcchar8FieldDadShift = 22
	RegisterOtghshcchar8FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar8Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldDadMask) >> RegisterOtghshcchar8FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar8Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldDadMask)|(uint32(value)<<RegisterOtghshcchar8FieldDadShift))
}

const (
	RegisterOtghshcchar8FieldOddfrmShift = 29
	RegisterOtghshcchar8FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar8Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar8Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar8FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar8FieldChdisShift = 30
	RegisterOtghshcchar8FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar8Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar8Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar8FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar8FieldChenaShift = 31
	RegisterOtghshcchar8FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar8Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar8Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar8FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldChenaMask)
	}
}

// registerOtghshcsplt8Type OTG_HS host channel-8 split control register
type registerOtghshcsplt8Type uint32

const (
	RegisterOtghshcsplt8FieldPrtaddrShift = 0
	RegisterOtghshcsplt8FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt8Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt8FieldPrtaddrMask) >> RegisterOtghshcsplt8FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt8Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt8FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt8FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt8FieldHubaddrShift = 7
	RegisterOtghshcsplt8FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt8Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt8FieldHubaddrMask) >> RegisterOtghshcsplt8FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt8Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt8FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt8FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt8FieldXactposShift = 14
	RegisterOtghshcsplt8FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt8Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt8FieldXactposMask) >> RegisterOtghshcsplt8FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt8Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt8FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt8FieldXactposShift))
}

const (
	RegisterOtghshcsplt8FieldComplspltShift = 16
	RegisterOtghshcsplt8FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt8Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt8FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt8Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt8FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt8FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt8FieldSplitenShift = 31
	RegisterOtghshcsplt8FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt8Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt8FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt8Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt8FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt8FieldSplitenMask)
	}
}

// registerOtghshcint8Type OTG_HS host channel-8 interrupt register
type registerOtghshcint8Type uint32

const (
	RegisterOtghshcint8FieldXfrcShift = 0
	RegisterOtghshcint8FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint8Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint8Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint8FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint8FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint8FieldChhShift = 1
	RegisterOtghshcint8FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint8Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint8Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint8FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint8FieldChhMask)
	}
}

const (
	RegisterOtghshcint8FieldAhberrShift = 2
	RegisterOtghshcint8FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint8Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint8Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint8FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint8FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint8FieldStallShift = 3
	RegisterOtghshcint8FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint8Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint8Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint8FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint8FieldStallMask)
	}
}

const (
	RegisterOtghshcint8FieldNakShift = 4
	RegisterOtghshcint8FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint8Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint8Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint8FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint8FieldNakMask)
	}
}

const (
	RegisterOtghshcint8FieldAckShift = 5
	RegisterOtghshcint8FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint8Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint8Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint8FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint8FieldAckMask)
	}
}

const (
	RegisterOtghshcint8FieldNyetShift = 6
	RegisterOtghshcint8FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint8Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint8Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint8FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint8FieldNyetMask)
	}
}

const (
	RegisterOtghshcint8FieldTxerrShift = 7
	RegisterOtghshcint8FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint8Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint8Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint8FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint8FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint8FieldBberrShift = 8
	RegisterOtghshcint8FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint8Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint8Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint8FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint8FieldBberrMask)
	}
}

const (
	RegisterOtghshcint8FieldFrmorShift = 9
	RegisterOtghshcint8FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint8Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint8Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint8FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint8FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint8FieldDterrShift = 10
	RegisterOtghshcint8FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint8Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint8Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint8FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint8FieldDterrMask)
	}
}

// registerOtghshcintmsk8Type OTG_HS host channel-8 interrupt mask register
type registerOtghshcintmsk8Type uint32

const (
	RegisterOtghshcintmsk8FieldXfrcmShift = 0
	RegisterOtghshcintmsk8FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk8Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk8Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk8FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk8FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk8FieldChhmShift = 1
	RegisterOtghshcintmsk8FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk8Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk8Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk8FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk8FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk8FieldAhberrShift = 2
	RegisterOtghshcintmsk8FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk8Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk8Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk8FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk8FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk8FieldStallmShift = 3
	RegisterOtghshcintmsk8FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk8Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk8Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk8FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk8FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk8FieldNakmShift = 4
	RegisterOtghshcintmsk8FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk8Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk8Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk8FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk8FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk8FieldAckmShift = 5
	RegisterOtghshcintmsk8FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk8Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk8Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk8FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk8FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk8FieldNyetShift = 6
	RegisterOtghshcintmsk8FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtghshcintmsk8Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtghshcintmsk8Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk8FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk8FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk8FieldTxerrmShift = 7
	RegisterOtghshcintmsk8FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtghshcintmsk8Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtghshcintmsk8Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk8FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk8FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk8FieldBberrmShift = 8
	RegisterOtghshcintmsk8FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtghshcintmsk8Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtghshcintmsk8Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk8FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk8FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk8FieldFrmormShift = 9
	RegisterOtghshcintmsk8FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk8Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk8Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk8FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk8FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk8FieldDterrmShift = 10
	RegisterOtghshcintmsk8FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk8Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk8Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk8FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk8FieldDterrmMask)
	}
}

// registerOtghshctsiz8Type OTG_HS host channel-8 transfer size register
type registerOtghshctsiz8Type uint32

const (
	RegisterOtghshctsiz8FieldXfrsizShift = 0
	RegisterOtghshctsiz8FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz8Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz8FieldXfrsizMask) >> RegisterOtghshctsiz8FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz8Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz8FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz8FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz8FieldPktcntShift = 19
	RegisterOtghshctsiz8FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz8Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz8FieldPktcntMask) >> RegisterOtghshctsiz8FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz8Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz8FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz8FieldPktcntShift))
}

const (
	RegisterOtghshctsiz8FieldDpidShift = 29
	RegisterOtghshctsiz8FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz8Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz8FieldDpidMask) >> RegisterOtghshctsiz8FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz8Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz8FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz8FieldDpidShift))
}

// registerOtghshcdma8Type OTG_HS host channel-8 DMA address register
type registerOtghshcdma8Type uint32

const (
	RegisterOtghshcdma8FieldDmaaddrShift = 0
	RegisterOtghshcdma8FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma8Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma8FieldDmaaddrMask) >> RegisterOtghshcdma8FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma8Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma8FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma8FieldDmaaddrShift))
}

// registerOtghshcchar9Type OTG_HS host channel-9 characteristics register
type registerOtghshcchar9Type uint32

const (
	RegisterOtghshcchar9FieldMpsizShift = 0
	RegisterOtghshcchar9FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar9Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldMpsizMask) >> RegisterOtghshcchar9FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar9Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar9FieldMpsizShift))
}

const (
	RegisterOtghshcchar9FieldEpnumShift = 11
	RegisterOtghshcchar9FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar9Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldEpnumMask) >> RegisterOtghshcchar9FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar9Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar9FieldEpnumShift))
}

const (
	RegisterOtghshcchar9FieldEpdirShift = 15
	RegisterOtghshcchar9FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar9Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar9Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar9FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar9FieldLsdevShift = 17
	RegisterOtghshcchar9FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar9Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar9Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar9FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar9FieldEptypShift = 18
	RegisterOtghshcchar9FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar9Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldEptypMask) >> RegisterOtghshcchar9FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar9Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar9FieldEptypShift))
}

const (
	RegisterOtghshcchar9FieldMcShift = 20
	RegisterOtghshcchar9FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar9Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldMcMask) >> RegisterOtghshcchar9FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar9Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldMcMask)|(uint32(value)<<RegisterOtghshcchar9FieldMcShift))
}

const (
	RegisterOtghshcchar9FieldDadShift = 22
	RegisterOtghshcchar9FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar9Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldDadMask) >> RegisterOtghshcchar9FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar9Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldDadMask)|(uint32(value)<<RegisterOtghshcchar9FieldDadShift))
}

const (
	RegisterOtghshcchar9FieldOddfrmShift = 29
	RegisterOtghshcchar9FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar9Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar9Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar9FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar9FieldChdisShift = 30
	RegisterOtghshcchar9FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar9Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar9Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar9FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar9FieldChenaShift = 31
	RegisterOtghshcchar9FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar9Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar9Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar9FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldChenaMask)
	}
}

// registerOtghshcsplt9Type OTG_HS host channel-9 split control register
type registerOtghshcsplt9Type uint32

const (
	RegisterOtghshcsplt9FieldPrtaddrShift = 0
	RegisterOtghshcsplt9FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt9Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt9FieldPrtaddrMask) >> RegisterOtghshcsplt9FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt9Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt9FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt9FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt9FieldHubaddrShift = 7
	RegisterOtghshcsplt9FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt9Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt9FieldHubaddrMask) >> RegisterOtghshcsplt9FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt9Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt9FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt9FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt9FieldXactposShift = 14
	RegisterOtghshcsplt9FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt9Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt9FieldXactposMask) >> RegisterOtghshcsplt9FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt9Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt9FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt9FieldXactposShift))
}

const (
	RegisterOtghshcsplt9FieldComplspltShift = 16
	RegisterOtghshcsplt9FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt9Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt9FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt9Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt9FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt9FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt9FieldSplitenShift = 31
	RegisterOtghshcsplt9FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt9Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt9FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt9Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt9FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt9FieldSplitenMask)
	}
}

// registerOtghshcint9Type OTG_HS host channel-9 interrupt register
type registerOtghshcint9Type uint32

const (
	RegisterOtghshcint9FieldXfrcShift = 0
	RegisterOtghshcint9FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint9Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint9Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint9FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint9FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint9FieldChhShift = 1
	RegisterOtghshcint9FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint9Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint9Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint9FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint9FieldChhMask)
	}
}

const (
	RegisterOtghshcint9FieldAhberrShift = 2
	RegisterOtghshcint9FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint9Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint9Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint9FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint9FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint9FieldStallShift = 3
	RegisterOtghshcint9FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint9Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint9Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint9FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint9FieldStallMask)
	}
}

const (
	RegisterOtghshcint9FieldNakShift = 4
	RegisterOtghshcint9FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint9Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint9Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint9FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint9FieldNakMask)
	}
}

const (
	RegisterOtghshcint9FieldAckShift = 5
	RegisterOtghshcint9FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint9Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint9Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint9FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint9FieldAckMask)
	}
}

const (
	RegisterOtghshcint9FieldNyetShift = 6
	RegisterOtghshcint9FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint9Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint9Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint9FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint9FieldNyetMask)
	}
}

const (
	RegisterOtghshcint9FieldTxerrShift = 7
	RegisterOtghshcint9FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint9Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint9Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint9FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint9FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint9FieldBberrShift = 8
	RegisterOtghshcint9FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint9Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint9Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint9FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint9FieldBberrMask)
	}
}

const (
	RegisterOtghshcint9FieldFrmorShift = 9
	RegisterOtghshcint9FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint9Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint9Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint9FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint9FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint9FieldDterrShift = 10
	RegisterOtghshcint9FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint9Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint9Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint9FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint9FieldDterrMask)
	}
}

// registerOtghshcintmsk9Type OTG_HS host channel-9 interrupt mask register
type registerOtghshcintmsk9Type uint32

const (
	RegisterOtghshcintmsk9FieldXfrcmShift = 0
	RegisterOtghshcintmsk9FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk9Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk9Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk9FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk9FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk9FieldChhmShift = 1
	RegisterOtghshcintmsk9FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk9Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk9Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk9FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk9FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk9FieldAhberrShift = 2
	RegisterOtghshcintmsk9FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk9Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk9Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk9FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk9FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk9FieldStallmShift = 3
	RegisterOtghshcintmsk9FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk9Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk9Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk9FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk9FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk9FieldNakmShift = 4
	RegisterOtghshcintmsk9FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk9Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk9Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk9FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk9FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk9FieldAckmShift = 5
	RegisterOtghshcintmsk9FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk9Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk9Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk9FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk9FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk9FieldNyetShift = 6
	RegisterOtghshcintmsk9FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtghshcintmsk9Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtghshcintmsk9Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk9FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk9FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk9FieldTxerrmShift = 7
	RegisterOtghshcintmsk9FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtghshcintmsk9Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtghshcintmsk9Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk9FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk9FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk9FieldBberrmShift = 8
	RegisterOtghshcintmsk9FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtghshcintmsk9Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtghshcintmsk9Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk9FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk9FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk9FieldFrmormShift = 9
	RegisterOtghshcintmsk9FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk9Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk9Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk9FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk9FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk9FieldDterrmShift = 10
	RegisterOtghshcintmsk9FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk9Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk9Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk9FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk9FieldDterrmMask)
	}
}

// registerOtghshctsiz9Type OTG_HS host channel-9 transfer size register
type registerOtghshctsiz9Type uint32

const (
	RegisterOtghshctsiz9FieldXfrsizShift = 0
	RegisterOtghshctsiz9FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz9Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz9FieldXfrsizMask) >> RegisterOtghshctsiz9FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz9Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz9FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz9FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz9FieldPktcntShift = 19
	RegisterOtghshctsiz9FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz9Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz9FieldPktcntMask) >> RegisterOtghshctsiz9FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz9Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz9FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz9FieldPktcntShift))
}

const (
	RegisterOtghshctsiz9FieldDpidShift = 29
	RegisterOtghshctsiz9FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz9Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz9FieldDpidMask) >> RegisterOtghshctsiz9FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz9Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz9FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz9FieldDpidShift))
}

// registerOtghshcdma9Type OTG_HS host channel-9 DMA address register
type registerOtghshcdma9Type uint32

const (
	RegisterOtghshcdma9FieldDmaaddrShift = 0
	RegisterOtghshcdma9FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma9Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma9FieldDmaaddrMask) >> RegisterOtghshcdma9FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma9Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma9FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma9FieldDmaaddrShift))
}

// registerOtghshcchar10Type OTG_HS host channel-10 characteristics register
type registerOtghshcchar10Type uint32

const (
	RegisterOtghshcchar10FieldMpsizShift = 0
	RegisterOtghshcchar10FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar10Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldMpsizMask) >> RegisterOtghshcchar10FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar10Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar10FieldMpsizShift))
}

const (
	RegisterOtghshcchar10FieldEpnumShift = 11
	RegisterOtghshcchar10FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar10Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldEpnumMask) >> RegisterOtghshcchar10FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar10Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar10FieldEpnumShift))
}

const (
	RegisterOtghshcchar10FieldEpdirShift = 15
	RegisterOtghshcchar10FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar10Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar10Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar10FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar10FieldLsdevShift = 17
	RegisterOtghshcchar10FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar10Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar10Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar10FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar10FieldEptypShift = 18
	RegisterOtghshcchar10FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar10Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldEptypMask) >> RegisterOtghshcchar10FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar10Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar10FieldEptypShift))
}

const (
	RegisterOtghshcchar10FieldMcShift = 20
	RegisterOtghshcchar10FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar10Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldMcMask) >> RegisterOtghshcchar10FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar10Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldMcMask)|(uint32(value)<<RegisterOtghshcchar10FieldMcShift))
}

const (
	RegisterOtghshcchar10FieldDadShift = 22
	RegisterOtghshcchar10FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar10Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldDadMask) >> RegisterOtghshcchar10FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar10Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldDadMask)|(uint32(value)<<RegisterOtghshcchar10FieldDadShift))
}

const (
	RegisterOtghshcchar10FieldOddfrmShift = 29
	RegisterOtghshcchar10FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar10Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar10Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar10FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar10FieldChdisShift = 30
	RegisterOtghshcchar10FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar10Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar10Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar10FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar10FieldChenaShift = 31
	RegisterOtghshcchar10FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar10Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar10Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar10FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldChenaMask)
	}
}

// registerOtghshcsplt10Type OTG_HS host channel-10 split control register
type registerOtghshcsplt10Type uint32

const (
	RegisterOtghshcsplt10FieldPrtaddrShift = 0
	RegisterOtghshcsplt10FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt10Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt10FieldPrtaddrMask) >> RegisterOtghshcsplt10FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt10Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt10FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt10FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt10FieldHubaddrShift = 7
	RegisterOtghshcsplt10FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt10Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt10FieldHubaddrMask) >> RegisterOtghshcsplt10FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt10Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt10FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt10FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt10FieldXactposShift = 14
	RegisterOtghshcsplt10FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt10Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt10FieldXactposMask) >> RegisterOtghshcsplt10FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt10Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt10FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt10FieldXactposShift))
}

const (
	RegisterOtghshcsplt10FieldComplspltShift = 16
	RegisterOtghshcsplt10FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt10Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt10FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt10Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt10FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt10FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt10FieldSplitenShift = 31
	RegisterOtghshcsplt10FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt10Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt10FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt10Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt10FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt10FieldSplitenMask)
	}
}

// registerOtghshcint10Type OTG_HS host channel-10 interrupt register
type registerOtghshcint10Type uint32

const (
	RegisterOtghshcint10FieldXfrcShift = 0
	RegisterOtghshcint10FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint10Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint10Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint10FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint10FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint10FieldChhShift = 1
	RegisterOtghshcint10FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint10Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint10Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint10FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint10FieldChhMask)
	}
}

const (
	RegisterOtghshcint10FieldAhberrShift = 2
	RegisterOtghshcint10FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint10Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint10Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint10FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint10FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint10FieldStallShift = 3
	RegisterOtghshcint10FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint10Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint10Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint10FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint10FieldStallMask)
	}
}

const (
	RegisterOtghshcint10FieldNakShift = 4
	RegisterOtghshcint10FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint10Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint10Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint10FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint10FieldNakMask)
	}
}

const (
	RegisterOtghshcint10FieldAckShift = 5
	RegisterOtghshcint10FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint10Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint10Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint10FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint10FieldAckMask)
	}
}

const (
	RegisterOtghshcint10FieldNyetShift = 6
	RegisterOtghshcint10FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint10Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint10Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint10FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint10FieldNyetMask)
	}
}

const (
	RegisterOtghshcint10FieldTxerrShift = 7
	RegisterOtghshcint10FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint10Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint10Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint10FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint10FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint10FieldBberrShift = 8
	RegisterOtghshcint10FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint10Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint10Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint10FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint10FieldBberrMask)
	}
}

const (
	RegisterOtghshcint10FieldFrmorShift = 9
	RegisterOtghshcint10FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint10Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint10Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint10FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint10FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint10FieldDterrShift = 10
	RegisterOtghshcint10FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint10Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint10Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint10FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint10FieldDterrMask)
	}
}

// registerOtghshcintmsk10Type OTG_HS host channel-10 interrupt mask register
type registerOtghshcintmsk10Type uint32

const (
	RegisterOtghshcintmsk10FieldXfrcmShift = 0
	RegisterOtghshcintmsk10FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk10Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk10Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk10FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk10FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk10FieldChhmShift = 1
	RegisterOtghshcintmsk10FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk10Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk10Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk10FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk10FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk10FieldAhberrShift = 2
	RegisterOtghshcintmsk10FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk10Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk10Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk10FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk10FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk10FieldStallmShift = 3
	RegisterOtghshcintmsk10FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk10Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk10Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk10FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk10FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk10FieldNakmShift = 4
	RegisterOtghshcintmsk10FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk10Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk10Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk10FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk10FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk10FieldAckmShift = 5
	RegisterOtghshcintmsk10FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk10Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk10Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk10FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk10FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk10FieldNyetShift = 6
	RegisterOtghshcintmsk10FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtghshcintmsk10Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtghshcintmsk10Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk10FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk10FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk10FieldTxerrmShift = 7
	RegisterOtghshcintmsk10FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtghshcintmsk10Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtghshcintmsk10Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk10FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk10FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk10FieldBberrmShift = 8
	RegisterOtghshcintmsk10FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtghshcintmsk10Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtghshcintmsk10Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk10FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk10FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk10FieldFrmormShift = 9
	RegisterOtghshcintmsk10FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk10Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk10Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk10FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk10FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk10FieldDterrmShift = 10
	RegisterOtghshcintmsk10FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk10Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk10Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk10FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk10FieldDterrmMask)
	}
}

// registerOtghshctsiz10Type OTG_HS host channel-10 transfer size register
type registerOtghshctsiz10Type uint32

const (
	RegisterOtghshctsiz10FieldXfrsizShift = 0
	RegisterOtghshctsiz10FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz10Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz10FieldXfrsizMask) >> RegisterOtghshctsiz10FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz10Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz10FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz10FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz10FieldPktcntShift = 19
	RegisterOtghshctsiz10FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz10Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz10FieldPktcntMask) >> RegisterOtghshctsiz10FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz10Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz10FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz10FieldPktcntShift))
}

const (
	RegisterOtghshctsiz10FieldDpidShift = 29
	RegisterOtghshctsiz10FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz10Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz10FieldDpidMask) >> RegisterOtghshctsiz10FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz10Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz10FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz10FieldDpidShift))
}

// registerOtghshcdma10Type OTG_HS host channel-10 DMA address register
type registerOtghshcdma10Type uint32

const (
	RegisterOtghshcdma10FieldDmaaddrShift = 0
	RegisterOtghshcdma10FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma10Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma10FieldDmaaddrMask) >> RegisterOtghshcdma10FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma10Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma10FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma10FieldDmaaddrShift))
}

// registerOtghshcchar11Type OTG_HS host channel-11 characteristics register
type registerOtghshcchar11Type uint32

const (
	RegisterOtghshcchar11FieldMpsizShift = 0
	RegisterOtghshcchar11FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar11Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldMpsizMask) >> RegisterOtghshcchar11FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar11Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar11FieldMpsizShift))
}

const (
	RegisterOtghshcchar11FieldEpnumShift = 11
	RegisterOtghshcchar11FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar11Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldEpnumMask) >> RegisterOtghshcchar11FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar11Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar11FieldEpnumShift))
}

const (
	RegisterOtghshcchar11FieldEpdirShift = 15
	RegisterOtghshcchar11FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar11Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar11Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar11FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar11FieldLsdevShift = 17
	RegisterOtghshcchar11FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar11Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar11Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar11FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar11FieldEptypShift = 18
	RegisterOtghshcchar11FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar11Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldEptypMask) >> RegisterOtghshcchar11FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar11Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar11FieldEptypShift))
}

const (
	RegisterOtghshcchar11FieldMcShift = 20
	RegisterOtghshcchar11FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar11Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldMcMask) >> RegisterOtghshcchar11FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar11Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldMcMask)|(uint32(value)<<RegisterOtghshcchar11FieldMcShift))
}

const (
	RegisterOtghshcchar11FieldDadShift = 22
	RegisterOtghshcchar11FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar11Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldDadMask) >> RegisterOtghshcchar11FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar11Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldDadMask)|(uint32(value)<<RegisterOtghshcchar11FieldDadShift))
}

const (
	RegisterOtghshcchar11FieldOddfrmShift = 29
	RegisterOtghshcchar11FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar11Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar11Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar11FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar11FieldChdisShift = 30
	RegisterOtghshcchar11FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar11Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar11Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar11FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar11FieldChenaShift = 31
	RegisterOtghshcchar11FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar11Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar11Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar11FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldChenaMask)
	}
}

// registerOtghshcsplt11Type OTG_HS host channel-11 split control register
type registerOtghshcsplt11Type uint32

const (
	RegisterOtghshcsplt11FieldPrtaddrShift = 0
	RegisterOtghshcsplt11FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt11Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt11FieldPrtaddrMask) >> RegisterOtghshcsplt11FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt11Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt11FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt11FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt11FieldHubaddrShift = 7
	RegisterOtghshcsplt11FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt11Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt11FieldHubaddrMask) >> RegisterOtghshcsplt11FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt11Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt11FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt11FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt11FieldXactposShift = 14
	RegisterOtghshcsplt11FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt11Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt11FieldXactposMask) >> RegisterOtghshcsplt11FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt11Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt11FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt11FieldXactposShift))
}

const (
	RegisterOtghshcsplt11FieldComplspltShift = 16
	RegisterOtghshcsplt11FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt11Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt11FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt11Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt11FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt11FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt11FieldSplitenShift = 31
	RegisterOtghshcsplt11FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt11Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt11FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt11Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt11FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt11FieldSplitenMask)
	}
}

// registerOtghshcint11Type OTG_HS host channel-11 interrupt register
type registerOtghshcint11Type uint32

const (
	RegisterOtghshcint11FieldXfrcShift = 0
	RegisterOtghshcint11FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint11Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint11Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint11FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint11FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint11FieldChhShift = 1
	RegisterOtghshcint11FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint11Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint11Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint11FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint11FieldChhMask)
	}
}

const (
	RegisterOtghshcint11FieldAhberrShift = 2
	RegisterOtghshcint11FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint11Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint11Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint11FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint11FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint11FieldStallShift = 3
	RegisterOtghshcint11FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint11Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint11Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint11FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint11FieldStallMask)
	}
}

const (
	RegisterOtghshcint11FieldNakShift = 4
	RegisterOtghshcint11FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint11Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint11Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint11FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint11FieldNakMask)
	}
}

const (
	RegisterOtghshcint11FieldAckShift = 5
	RegisterOtghshcint11FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint11Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint11Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint11FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint11FieldAckMask)
	}
}

const (
	RegisterOtghshcint11FieldNyetShift = 6
	RegisterOtghshcint11FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint11Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint11Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint11FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint11FieldNyetMask)
	}
}

const (
	RegisterOtghshcint11FieldTxerrShift = 7
	RegisterOtghshcint11FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint11Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint11Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint11FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint11FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint11FieldBberrShift = 8
	RegisterOtghshcint11FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint11Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint11Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint11FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint11FieldBberrMask)
	}
}

const (
	RegisterOtghshcint11FieldFrmorShift = 9
	RegisterOtghshcint11FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint11Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint11Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint11FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint11FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint11FieldDterrShift = 10
	RegisterOtghshcint11FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint11Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint11Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint11FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint11FieldDterrMask)
	}
}

// registerOtghshcintmsk11Type OTG_HS host channel-11 interrupt mask register
type registerOtghshcintmsk11Type uint32

const (
	RegisterOtghshcintmsk11FieldXfrcmShift = 0
	RegisterOtghshcintmsk11FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk11Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk11Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk11FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk11FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk11FieldChhmShift = 1
	RegisterOtghshcintmsk11FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk11Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk11Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk11FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk11FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk11FieldAhberrShift = 2
	RegisterOtghshcintmsk11FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk11Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk11Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk11FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk11FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk11FieldStallmShift = 3
	RegisterOtghshcintmsk11FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk11Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk11Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk11FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk11FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk11FieldNakmShift = 4
	RegisterOtghshcintmsk11FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk11Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk11Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk11FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk11FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk11FieldAckmShift = 5
	RegisterOtghshcintmsk11FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk11Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk11Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk11FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk11FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk11FieldNyetShift = 6
	RegisterOtghshcintmsk11FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtghshcintmsk11Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtghshcintmsk11Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk11FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk11FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk11FieldTxerrmShift = 7
	RegisterOtghshcintmsk11FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtghshcintmsk11Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtghshcintmsk11Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk11FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk11FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk11FieldBberrmShift = 8
	RegisterOtghshcintmsk11FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtghshcintmsk11Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtghshcintmsk11Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk11FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk11FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk11FieldFrmormShift = 9
	RegisterOtghshcintmsk11FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk11Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk11Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk11FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk11FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk11FieldDterrmShift = 10
	RegisterOtghshcintmsk11FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk11Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk11Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk11FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk11FieldDterrmMask)
	}
}

// registerOtghshctsiz11Type OTG_HS host channel-11 transfer size register
type registerOtghshctsiz11Type uint32

const (
	RegisterOtghshctsiz11FieldXfrsizShift = 0
	RegisterOtghshctsiz11FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz11Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz11FieldXfrsizMask) >> RegisterOtghshctsiz11FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz11Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz11FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz11FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz11FieldPktcntShift = 19
	RegisterOtghshctsiz11FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz11Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz11FieldPktcntMask) >> RegisterOtghshctsiz11FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz11Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz11FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz11FieldPktcntShift))
}

const (
	RegisterOtghshctsiz11FieldDpidShift = 29
	RegisterOtghshctsiz11FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz11Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz11FieldDpidMask) >> RegisterOtghshctsiz11FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz11Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz11FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz11FieldDpidShift))
}

// registerOtghshcdma11Type OTG_HS host channel-11 DMA address register
type registerOtghshcdma11Type uint32

const (
	RegisterOtghshcdma11FieldDmaaddrShift = 0
	RegisterOtghshcdma11FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma11Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma11FieldDmaaddrMask) >> RegisterOtghshcdma11FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma11Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma11FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma11FieldDmaaddrShift))
}

// registerOtghshcchar12Type OTG_HS host channel-12 characteristics register
type registerOtghshcchar12Type uint32

const (
	RegisterOtghshcchar12FieldMpsizShift = 0
	RegisterOtghshcchar12FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar12Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldMpsizMask) >> RegisterOtghshcchar12FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar12Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar12FieldMpsizShift))
}

const (
	RegisterOtghshcchar12FieldEpnumShift = 11
	RegisterOtghshcchar12FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar12Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldEpnumMask) >> RegisterOtghshcchar12FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar12Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar12FieldEpnumShift))
}

const (
	RegisterOtghshcchar12FieldEpdirShift = 15
	RegisterOtghshcchar12FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar12Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar12Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar12FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar12FieldLsdevShift = 17
	RegisterOtghshcchar12FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar12Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar12Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar12FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar12FieldEptypShift = 18
	RegisterOtghshcchar12FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar12Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldEptypMask) >> RegisterOtghshcchar12FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar12Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar12FieldEptypShift))
}

const (
	RegisterOtghshcchar12FieldMcShift = 20
	RegisterOtghshcchar12FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar12Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldMcMask) >> RegisterOtghshcchar12FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar12Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldMcMask)|(uint32(value)<<RegisterOtghshcchar12FieldMcShift))
}

const (
	RegisterOtghshcchar12FieldDadShift = 22
	RegisterOtghshcchar12FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar12Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldDadMask) >> RegisterOtghshcchar12FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar12Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldDadMask)|(uint32(value)<<RegisterOtghshcchar12FieldDadShift))
}

const (
	RegisterOtghshcchar12FieldOddfrmShift = 29
	RegisterOtghshcchar12FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar12Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar12Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar12FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar12FieldChdisShift = 30
	RegisterOtghshcchar12FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar12Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar12Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar12FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar12FieldChenaShift = 31
	RegisterOtghshcchar12FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar12Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar12Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar12FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldChenaMask)
	}
}

// registerOtghshcsplt12Type OTG_HS host channel-12 split control register
type registerOtghshcsplt12Type uint32

const (
	RegisterOtghshcsplt12FieldPrtaddrShift = 0
	RegisterOtghshcsplt12FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt12Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt12FieldPrtaddrMask) >> RegisterOtghshcsplt12FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt12Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt12FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt12FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt12FieldHubaddrShift = 7
	RegisterOtghshcsplt12FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt12Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt12FieldHubaddrMask) >> RegisterOtghshcsplt12FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt12Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt12FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt12FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt12FieldXactposShift = 14
	RegisterOtghshcsplt12FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt12Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt12FieldXactposMask) >> RegisterOtghshcsplt12FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt12Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt12FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt12FieldXactposShift))
}

const (
	RegisterOtghshcsplt12FieldComplspltShift = 16
	RegisterOtghshcsplt12FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt12Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt12FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt12Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt12FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt12FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt12FieldSplitenShift = 31
	RegisterOtghshcsplt12FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt12Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt12FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt12Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt12FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt12FieldSplitenMask)
	}
}

// registerOtghshcint12Type OTG_HS host channel-12 interrupt register
type registerOtghshcint12Type uint32

const (
	RegisterOtghshcint12FieldXfrcShift = 0
	RegisterOtghshcint12FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint12Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint12Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint12FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint12FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint12FieldChhShift = 1
	RegisterOtghshcint12FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint12Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint12Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint12FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint12FieldChhMask)
	}
}

const (
	RegisterOtghshcint12FieldAhberrShift = 2
	RegisterOtghshcint12FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint12Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint12Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint12FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint12FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint12FieldStallShift = 3
	RegisterOtghshcint12FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint12Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint12Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint12FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint12FieldStallMask)
	}
}

const (
	RegisterOtghshcint12FieldNakShift = 4
	RegisterOtghshcint12FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint12Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint12Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint12FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint12FieldNakMask)
	}
}

const (
	RegisterOtghshcint12FieldAckShift = 5
	RegisterOtghshcint12FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint12Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint12Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint12FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint12FieldAckMask)
	}
}

const (
	RegisterOtghshcint12FieldNyetShift = 6
	RegisterOtghshcint12FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint12Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint12Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint12FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint12FieldNyetMask)
	}
}

const (
	RegisterOtghshcint12FieldTxerrShift = 7
	RegisterOtghshcint12FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint12Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint12Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint12FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint12FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint12FieldBberrShift = 8
	RegisterOtghshcint12FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint12Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint12Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint12FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint12FieldBberrMask)
	}
}

const (
	RegisterOtghshcint12FieldFrmorShift = 9
	RegisterOtghshcint12FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint12Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint12Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint12FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint12FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint12FieldDterrShift = 10
	RegisterOtghshcint12FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint12Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint12Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint12FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint12FieldDterrMask)
	}
}

// registerOtghshcintmsk12Type OTG_HS host channel-12 interrupt mask register
type registerOtghshcintmsk12Type uint32

const (
	RegisterOtghshcintmsk12FieldXfrcmShift = 0
	RegisterOtghshcintmsk12FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk12Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk12Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk12FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk12FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk12FieldChhmShift = 1
	RegisterOtghshcintmsk12FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk12Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk12Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk12FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk12FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk12FieldAhberrShift = 2
	RegisterOtghshcintmsk12FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk12Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk12Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk12FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk12FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk12FieldStallmShift = 3
	RegisterOtghshcintmsk12FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk12Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk12Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk12FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk12FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk12FieldNakmShift = 4
	RegisterOtghshcintmsk12FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk12Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk12Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk12FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk12FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk12FieldAckmShift = 5
	RegisterOtghshcintmsk12FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk12Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk12Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk12FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk12FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk12FieldNyetShift = 6
	RegisterOtghshcintmsk12FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcintmsk12Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcintmsk12Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk12FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk12FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk12FieldTxerrmShift = 7
	RegisterOtghshcintmsk12FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error
func (r *registerOtghshcintmsk12Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error
func (r *registerOtghshcintmsk12Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk12FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk12FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk12FieldBberrmShift = 8
	RegisterOtghshcintmsk12FieldBberrmMask  = 0x100
)

// GetBberrm Babble error
func (r *registerOtghshcintmsk12Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldBberrmMask) != 0
}

// SetBberrm Babble error
func (r *registerOtghshcintmsk12Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk12FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk12FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk12FieldFrmormShift = 9
	RegisterOtghshcintmsk12FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk12Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk12Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk12FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk12FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk12FieldDterrmShift = 10
	RegisterOtghshcintmsk12FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk12Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk12Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk12FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk12FieldDterrmMask)
	}
}

// registerOtghshctsiz12Type OTG_HS host channel-12 transfer size register
type registerOtghshctsiz12Type uint32

const (
	RegisterOtghshctsiz12FieldXfrsizShift = 0
	RegisterOtghshctsiz12FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz12Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz12FieldXfrsizMask) >> RegisterOtghshctsiz12FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz12Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz12FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz12FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz12FieldPktcntShift = 19
	RegisterOtghshctsiz12FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz12Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz12FieldPktcntMask) >> RegisterOtghshctsiz12FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz12Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz12FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz12FieldPktcntShift))
}

const (
	RegisterOtghshctsiz12FieldDpidShift = 29
	RegisterOtghshctsiz12FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz12Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz12FieldDpidMask) >> RegisterOtghshctsiz12FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz12Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz12FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz12FieldDpidShift))
}

// registerOtghshcdma12Type OTG_HS host channel-12 DMA address register
type registerOtghshcdma12Type uint32

const (
	RegisterOtghshcdma12FieldDmaaddrShift = 0
	RegisterOtghshcdma12FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma12Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma12FieldDmaaddrMask) >> RegisterOtghshcdma12FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma12Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma12FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma12FieldDmaaddrShift))
}

// registerOtghshcchar13Type OTG_HS host channel-13 characteristics register
type registerOtghshcchar13Type uint32

const (
	RegisterOtghshcchar13FieldMpsizShift = 0
	RegisterOtghshcchar13FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar13Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldMpsizMask) >> RegisterOtghshcchar13FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar13Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar13FieldMpsizShift))
}

const (
	RegisterOtghshcchar13FieldEpnumShift = 11
	RegisterOtghshcchar13FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar13Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldEpnumMask) >> RegisterOtghshcchar13FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar13Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar13FieldEpnumShift))
}

const (
	RegisterOtghshcchar13FieldEpdirShift = 15
	RegisterOtghshcchar13FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar13Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar13Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar13FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar13FieldLsdevShift = 17
	RegisterOtghshcchar13FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar13Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar13Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar13FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar13FieldEptypShift = 18
	RegisterOtghshcchar13FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar13Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldEptypMask) >> RegisterOtghshcchar13FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar13Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar13FieldEptypShift))
}

const (
	RegisterOtghshcchar13FieldMcShift = 20
	RegisterOtghshcchar13FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar13Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldMcMask) >> RegisterOtghshcchar13FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar13Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldMcMask)|(uint32(value)<<RegisterOtghshcchar13FieldMcShift))
}

const (
	RegisterOtghshcchar13FieldDadShift = 22
	RegisterOtghshcchar13FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar13Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldDadMask) >> RegisterOtghshcchar13FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar13Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldDadMask)|(uint32(value)<<RegisterOtghshcchar13FieldDadShift))
}

const (
	RegisterOtghshcchar13FieldOddfrmShift = 29
	RegisterOtghshcchar13FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar13Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar13Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar13FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar13FieldChdisShift = 30
	RegisterOtghshcchar13FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar13Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar13Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar13FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar13FieldChenaShift = 31
	RegisterOtghshcchar13FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar13Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar13Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar13FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldChenaMask)
	}
}

// registerOtghshcsplt13Type OTG_HS host channel-13 split control register
type registerOtghshcsplt13Type uint32

const (
	RegisterOtghshcsplt13FieldPrtaddrShift = 0
	RegisterOtghshcsplt13FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt13Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt13FieldPrtaddrMask) >> RegisterOtghshcsplt13FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt13Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt13FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt13FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt13FieldHubaddrShift = 7
	RegisterOtghshcsplt13FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt13Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt13FieldHubaddrMask) >> RegisterOtghshcsplt13FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt13Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt13FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt13FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt13FieldXactposShift = 14
	RegisterOtghshcsplt13FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt13Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt13FieldXactposMask) >> RegisterOtghshcsplt13FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt13Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt13FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt13FieldXactposShift))
}

const (
	RegisterOtghshcsplt13FieldComplspltShift = 16
	RegisterOtghshcsplt13FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt13Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt13FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt13Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt13FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt13FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt13FieldSplitenShift = 31
	RegisterOtghshcsplt13FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt13Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt13FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt13Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt13FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt13FieldSplitenMask)
	}
}

// registerOtghshcint13Type OTG_HS host channel-13 interrupt register
type registerOtghshcint13Type uint32

const (
	RegisterOtghshcint13FieldXfrcShift = 0
	RegisterOtghshcint13FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint13Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint13Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint13FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint13FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint13FieldChhShift = 1
	RegisterOtghshcint13FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint13Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint13Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint13FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint13FieldChhMask)
	}
}

const (
	RegisterOtghshcint13FieldAhberrShift = 2
	RegisterOtghshcint13FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint13Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint13Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint13FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint13FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint13FieldStallShift = 3
	RegisterOtghshcint13FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint13Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint13Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint13FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint13FieldStallMask)
	}
}

const (
	RegisterOtghshcint13FieldNakShift = 4
	RegisterOtghshcint13FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint13Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint13Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint13FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint13FieldNakMask)
	}
}

const (
	RegisterOtghshcint13FieldAckShift = 5
	RegisterOtghshcint13FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint13Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint13Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint13FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint13FieldAckMask)
	}
}

const (
	RegisterOtghshcint13FieldNyetShift = 6
	RegisterOtghshcint13FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint13Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint13Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint13FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint13FieldNyetMask)
	}
}

const (
	RegisterOtghshcint13FieldTxerrShift = 7
	RegisterOtghshcint13FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint13Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint13Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint13FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint13FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint13FieldBberrShift = 8
	RegisterOtghshcint13FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint13Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint13Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint13FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint13FieldBberrMask)
	}
}

const (
	RegisterOtghshcint13FieldFrmorShift = 9
	RegisterOtghshcint13FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint13Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint13Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint13FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint13FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint13FieldDterrShift = 10
	RegisterOtghshcint13FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint13Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint13Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint13FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint13FieldDterrMask)
	}
}

// registerOtghshcintmsk13Type OTG_HS host channel-13 interrupt mask register
type registerOtghshcintmsk13Type uint32

const (
	RegisterOtghshcintmsk13FieldXfrcmShift = 0
	RegisterOtghshcintmsk13FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk13Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk13Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk13FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk13FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk13FieldChhmShift = 1
	RegisterOtghshcintmsk13FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk13Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk13Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk13FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk13FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk13FieldAhberrShift = 2
	RegisterOtghshcintmsk13FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk13Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk13Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk13FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk13FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk13FieldStallmShift = 3
	RegisterOtghshcintmsk13FieldStallmMask  = 0x8
)

// GetStallm STALLM response received interrupt mask
func (r *registerOtghshcintmsk13Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldStallmMask) != 0
}

// SetStallm STALLM response received interrupt mask
func (r *registerOtghshcintmsk13Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk13FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk13FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk13FieldNakmShift = 4
	RegisterOtghshcintmsk13FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk13Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk13Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk13FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk13FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk13FieldAckmShift = 5
	RegisterOtghshcintmsk13FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk13Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk13Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk13FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk13FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk13FieldNyetShift = 6
	RegisterOtghshcintmsk13FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcintmsk13Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcintmsk13Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk13FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk13FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk13FieldTxerrmShift = 7
	RegisterOtghshcintmsk13FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error
func (r *registerOtghshcintmsk13Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error
func (r *registerOtghshcintmsk13Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk13FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk13FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk13FieldBberrmShift = 8
	RegisterOtghshcintmsk13FieldBberrmMask  = 0x100
)

// GetBberrm Babble error
func (r *registerOtghshcintmsk13Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldBberrmMask) != 0
}

// SetBberrm Babble error
func (r *registerOtghshcintmsk13Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk13FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk13FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk13FieldFrmormShift = 9
	RegisterOtghshcintmsk13FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk13Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk13Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk13FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk13FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk13FieldDterrmShift = 10
	RegisterOtghshcintmsk13FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk13Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk13Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk13FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk13FieldDterrmMask)
	}
}

// registerOtghshctsiz13Type OTG_HS host channel-13 transfer size register
type registerOtghshctsiz13Type uint32

const (
	RegisterOtghshctsiz13FieldXfrsizShift = 0
	RegisterOtghshctsiz13FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz13Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz13FieldXfrsizMask) >> RegisterOtghshctsiz13FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz13Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz13FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz13FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz13FieldPktcntShift = 19
	RegisterOtghshctsiz13FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz13Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz13FieldPktcntMask) >> RegisterOtghshctsiz13FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz13Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz13FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz13FieldPktcntShift))
}

const (
	RegisterOtghshctsiz13FieldDpidShift = 29
	RegisterOtghshctsiz13FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz13Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz13FieldDpidMask) >> RegisterOtghshctsiz13FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz13Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz13FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz13FieldDpidShift))
}

// registerOtghshcdma13Type OTG_HS host channel-13 DMA address register
type registerOtghshcdma13Type uint32

const (
	RegisterOtghshcdma13FieldDmaaddrShift = 0
	RegisterOtghshcdma13FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma13Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma13FieldDmaaddrMask) >> RegisterOtghshcdma13FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma13Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma13FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma13FieldDmaaddrShift))
}

// registerOtghshcchar14Type OTG_HS host channel-14 characteristics register
type registerOtghshcchar14Type uint32

const (
	RegisterOtghshcchar14FieldMpsizShift = 0
	RegisterOtghshcchar14FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar14Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldMpsizMask) >> RegisterOtghshcchar14FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar14Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar14FieldMpsizShift))
}

const (
	RegisterOtghshcchar14FieldEpnumShift = 11
	RegisterOtghshcchar14FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar14Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldEpnumMask) >> RegisterOtghshcchar14FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar14Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar14FieldEpnumShift))
}

const (
	RegisterOtghshcchar14FieldEpdirShift = 15
	RegisterOtghshcchar14FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar14Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar14Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar14FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar14FieldLsdevShift = 17
	RegisterOtghshcchar14FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar14Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar14Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar14FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar14FieldEptypShift = 18
	RegisterOtghshcchar14FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar14Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldEptypMask) >> RegisterOtghshcchar14FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar14Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar14FieldEptypShift))
}

const (
	RegisterOtghshcchar14FieldMcShift = 20
	RegisterOtghshcchar14FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar14Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldMcMask) >> RegisterOtghshcchar14FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar14Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldMcMask)|(uint32(value)<<RegisterOtghshcchar14FieldMcShift))
}

const (
	RegisterOtghshcchar14FieldDadShift = 22
	RegisterOtghshcchar14FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar14Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldDadMask) >> RegisterOtghshcchar14FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar14Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldDadMask)|(uint32(value)<<RegisterOtghshcchar14FieldDadShift))
}

const (
	RegisterOtghshcchar14FieldOddfrmShift = 29
	RegisterOtghshcchar14FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar14Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar14Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar14FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar14FieldChdisShift = 30
	RegisterOtghshcchar14FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar14Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar14Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar14FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar14FieldChenaShift = 31
	RegisterOtghshcchar14FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar14Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar14Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar14FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldChenaMask)
	}
}

// registerOtghshcsplt14Type OTG_HS host channel-14 split control register
type registerOtghshcsplt14Type uint32

const (
	RegisterOtghshcsplt14FieldPrtaddrShift = 0
	RegisterOtghshcsplt14FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt14Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt14FieldPrtaddrMask) >> RegisterOtghshcsplt14FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt14Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt14FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt14FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt14FieldHubaddrShift = 7
	RegisterOtghshcsplt14FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt14Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt14FieldHubaddrMask) >> RegisterOtghshcsplt14FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt14Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt14FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt14FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt14FieldXactposShift = 14
	RegisterOtghshcsplt14FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt14Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt14FieldXactposMask) >> RegisterOtghshcsplt14FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt14Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt14FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt14FieldXactposShift))
}

const (
	RegisterOtghshcsplt14FieldComplspltShift = 16
	RegisterOtghshcsplt14FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt14Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt14FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt14Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt14FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt14FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt14FieldSplitenShift = 31
	RegisterOtghshcsplt14FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt14Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt14FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt14Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt14FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt14FieldSplitenMask)
	}
}

// registerOtghshcint14Type OTG_HS host channel-14 interrupt register
type registerOtghshcint14Type uint32

const (
	RegisterOtghshcint14FieldXfrcShift = 0
	RegisterOtghshcint14FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint14Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint14Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint14FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint14FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint14FieldChhShift = 1
	RegisterOtghshcint14FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint14Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint14Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint14FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint14FieldChhMask)
	}
}

const (
	RegisterOtghshcint14FieldAhberrShift = 2
	RegisterOtghshcint14FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint14Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint14Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint14FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint14FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint14FieldStallShift = 3
	RegisterOtghshcint14FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint14Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint14Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint14FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint14FieldStallMask)
	}
}

const (
	RegisterOtghshcint14FieldNakShift = 4
	RegisterOtghshcint14FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint14Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint14Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint14FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint14FieldNakMask)
	}
}

const (
	RegisterOtghshcint14FieldAckShift = 5
	RegisterOtghshcint14FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint14Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint14Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint14FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint14FieldAckMask)
	}
}

const (
	RegisterOtghshcint14FieldNyetShift = 6
	RegisterOtghshcint14FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint14Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint14Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint14FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint14FieldNyetMask)
	}
}

const (
	RegisterOtghshcint14FieldTxerrShift = 7
	RegisterOtghshcint14FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint14Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint14Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint14FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint14FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint14FieldBberrShift = 8
	RegisterOtghshcint14FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint14Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint14Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint14FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint14FieldBberrMask)
	}
}

const (
	RegisterOtghshcint14FieldFrmorShift = 9
	RegisterOtghshcint14FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint14Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint14Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint14FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint14FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint14FieldDterrShift = 10
	RegisterOtghshcint14FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint14Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint14Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint14FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint14FieldDterrMask)
	}
}

// registerOtghshcintmsk14Type OTG_HS host channel-14 interrupt mask register
type registerOtghshcintmsk14Type uint32

const (
	RegisterOtghshcintmsk14FieldXfrcmShift = 0
	RegisterOtghshcintmsk14FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk14Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk14Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk14FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk14FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk14FieldChhmShift = 1
	RegisterOtghshcintmsk14FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk14Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk14Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk14FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk14FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk14FieldAhberrShift = 2
	RegisterOtghshcintmsk14FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk14Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk14Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk14FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk14FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk14FieldStallmShift = 3
	RegisterOtghshcintmsk14FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk14Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtghshcintmsk14Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk14FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk14FieldStallmMask)
	}
}

const (
	RegisterOtghshcintmsk14FieldNakmShift = 4
	RegisterOtghshcintmsk14FieldNakmMask  = 0x10
)

// GetNakm NAKM response received interrupt mask
func (r *registerOtghshcintmsk14Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldNakmMask) != 0
}

// SetNakm NAKM response received interrupt mask
func (r *registerOtghshcintmsk14Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk14FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk14FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk14FieldAckmShift = 5
	RegisterOtghshcintmsk14FieldAckmMask  = 0x20
)

// GetAckm ACKM response received/transmitted interrupt mask
func (r *registerOtghshcintmsk14Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldAckmMask) != 0
}

// SetAckm ACKM response received/transmitted interrupt mask
func (r *registerOtghshcintmsk14Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk14FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk14FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk14FieldNyetShift = 6
	RegisterOtghshcintmsk14FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcintmsk14Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcintmsk14Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk14FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk14FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk14FieldTxerrmShift = 7
	RegisterOtghshcintmsk14FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error
func (r *registerOtghshcintmsk14Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error
func (r *registerOtghshcintmsk14Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk14FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk14FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk14FieldBberrmShift = 8
	RegisterOtghshcintmsk14FieldBberrmMask  = 0x100
)

// GetBberrm Babble error
func (r *registerOtghshcintmsk14Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldBberrmMask) != 0
}

// SetBberrm Babble error
func (r *registerOtghshcintmsk14Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk14FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk14FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk14FieldFrmormShift = 9
	RegisterOtghshcintmsk14FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk14Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk14Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk14FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk14FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk14FieldDterrmShift = 10
	RegisterOtghshcintmsk14FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk14Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk14Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk14FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk14FieldDterrmMask)
	}
}

// registerOtghshctsiz14Type OTG_HS host channel-14 transfer size register
type registerOtghshctsiz14Type uint32

const (
	RegisterOtghshctsiz14FieldXfrsizShift = 0
	RegisterOtghshctsiz14FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz14Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz14FieldXfrsizMask) >> RegisterOtghshctsiz14FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz14Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz14FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz14FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz14FieldPktcntShift = 19
	RegisterOtghshctsiz14FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz14Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz14FieldPktcntMask) >> RegisterOtghshctsiz14FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz14Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz14FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz14FieldPktcntShift))
}

const (
	RegisterOtghshctsiz14FieldDpidShift = 29
	RegisterOtghshctsiz14FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz14Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz14FieldDpidMask) >> RegisterOtghshctsiz14FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz14Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz14FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz14FieldDpidShift))
}

// registerOtghshcdma14Type OTG_HS host channel-14 DMA address register
type registerOtghshcdma14Type uint32

const (
	RegisterOtghshcdma14FieldDmaaddrShift = 0
	RegisterOtghshcdma14FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma14Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma14FieldDmaaddrMask) >> RegisterOtghshcdma14FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma14Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma14FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma14FieldDmaaddrShift))
}

// registerOtghshcchar15Type OTG_HS host channel-15 characteristics register
type registerOtghshcchar15Type uint32

const (
	RegisterOtghshcchar15FieldMpsizShift = 0
	RegisterOtghshcchar15FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtghshcchar15Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldMpsizMask) >> RegisterOtghshcchar15FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtghshcchar15Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar15FieldMpsizShift))
}

const (
	RegisterOtghshcchar15FieldEpnumShift = 11
	RegisterOtghshcchar15FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtghshcchar15Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldEpnumMask) >> RegisterOtghshcchar15FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtghshcchar15Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar15FieldEpnumShift))
}

const (
	RegisterOtghshcchar15FieldEpdirShift = 15
	RegisterOtghshcchar15FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtghshcchar15Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtghshcchar15Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar15FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldEpdirMask)
	}
}

const (
	RegisterOtghshcchar15FieldLsdevShift = 17
	RegisterOtghshcchar15FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtghshcchar15Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtghshcchar15Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar15FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldLsdevMask)
	}
}

const (
	RegisterOtghshcchar15FieldEptypShift = 18
	RegisterOtghshcchar15FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtghshcchar15Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldEptypMask) >> RegisterOtghshcchar15FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtghshcchar15Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar15FieldEptypShift))
}

const (
	RegisterOtghshcchar15FieldMcShift = 20
	RegisterOtghshcchar15FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar15Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldMcMask) >> RegisterOtghshcchar15FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtghshcchar15Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldMcMask)|(uint32(value)<<RegisterOtghshcchar15FieldMcShift))
}

const (
	RegisterOtghshcchar15FieldDadShift = 22
	RegisterOtghshcchar15FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtghshcchar15Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldDadMask) >> RegisterOtghshcchar15FieldDadShift)
}

// SetDad Device address
func (r *registerOtghshcchar15Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldDadMask)|(uint32(value)<<RegisterOtghshcchar15FieldDadShift))
}

const (
	RegisterOtghshcchar15FieldOddfrmShift = 29
	RegisterOtghshcchar15FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtghshcchar15Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtghshcchar15Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar15FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldOddfrmMask)
	}
}

const (
	RegisterOtghshcchar15FieldChdisShift = 30
	RegisterOtghshcchar15FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtghshcchar15Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtghshcchar15Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar15FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldChdisMask)
	}
}

const (
	RegisterOtghshcchar15FieldChenaShift = 31
	RegisterOtghshcchar15FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtghshcchar15Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtghshcchar15Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar15FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldChenaMask)
	}
}

// registerOtghshcsplt15Type OTG_HS host channel-15 split control register
type registerOtghshcsplt15Type uint32

const (
	RegisterOtghshcsplt15FieldPrtaddrShift = 0
	RegisterOtghshcsplt15FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtghshcsplt15Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt15FieldPrtaddrMask) >> RegisterOtghshcsplt15FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtghshcsplt15Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt15FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt15FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt15FieldHubaddrShift = 7
	RegisterOtghshcsplt15FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtghshcsplt15Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt15FieldHubaddrMask) >> RegisterOtghshcsplt15FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtghshcsplt15Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt15FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt15FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt15FieldXactposShift = 14
	RegisterOtghshcsplt15FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtghshcsplt15Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt15FieldXactposMask) >> RegisterOtghshcsplt15FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtghshcsplt15Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt15FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt15FieldXactposShift))
}

const (
	RegisterOtghshcsplt15FieldComplspltShift = 16
	RegisterOtghshcsplt15FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtghshcsplt15Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt15FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtghshcsplt15Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt15FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt15FieldComplspltMask)
	}
}

const (
	RegisterOtghshcsplt15FieldSplitenShift = 31
	RegisterOtghshcsplt15FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtghshcsplt15Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt15FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtghshcsplt15Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt15FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt15FieldSplitenMask)
	}
}

// registerOtghshcint15Type OTG_HS host channel-15 interrupt register
type registerOtghshcint15Type uint32

const (
	RegisterOtghshcint15FieldXfrcShift = 0
	RegisterOtghshcint15FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtghshcint15Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtghshcint15Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint15FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint15FieldXfrcMask)
	}
}

const (
	RegisterOtghshcint15FieldChhShift = 1
	RegisterOtghshcint15FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtghshcint15Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtghshcint15Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint15FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint15FieldChhMask)
	}
}

const (
	RegisterOtghshcint15FieldAhberrShift = 2
	RegisterOtghshcint15FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcint15Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcint15Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint15FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint15FieldAhberrMask)
	}
}

const (
	RegisterOtghshcint15FieldStallShift = 3
	RegisterOtghshcint15FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtghshcint15Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtghshcint15Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint15FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint15FieldStallMask)
	}
}

const (
	RegisterOtghshcint15FieldNakShift = 4
	RegisterOtghshcint15FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtghshcint15Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtghshcint15Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint15FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint15FieldNakMask)
	}
}

const (
	RegisterOtghshcint15FieldAckShift = 5
	RegisterOtghshcint15FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint15Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtghshcint15Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint15FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint15FieldAckMask)
	}
}

const (
	RegisterOtghshcint15FieldNyetShift = 6
	RegisterOtghshcint15FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcint15Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcint15Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint15FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint15FieldNyetMask)
	}
}

const (
	RegisterOtghshcint15FieldTxerrShift = 7
	RegisterOtghshcint15FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtghshcint15Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtghshcint15Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint15FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint15FieldTxerrMask)
	}
}

const (
	RegisterOtghshcint15FieldBberrShift = 8
	RegisterOtghshcint15FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtghshcint15Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtghshcint15Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint15FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint15FieldBberrMask)
	}
}

const (
	RegisterOtghshcint15FieldFrmorShift = 9
	RegisterOtghshcint15FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtghshcint15Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtghshcint15Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint15FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint15FieldFrmorMask)
	}
}

const (
	RegisterOtghshcint15FieldDterrShift = 10
	RegisterOtghshcint15FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtghshcint15Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtghshcint15Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint15FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint15FieldDterrMask)
	}
}

// registerOtghshcintmsk15Type OTG_HS host channel-15 interrupt mask register
type registerOtghshcintmsk15Type uint32

const (
	RegisterOtghshcintmsk15FieldXfrcmShift = 0
	RegisterOtghshcintmsk15FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk15Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtghshcintmsk15Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk15FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk15FieldXfrcmMask)
	}
}

const (
	RegisterOtghshcintmsk15FieldChhmShift = 1
	RegisterOtghshcintmsk15FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtghshcintmsk15Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtghshcintmsk15Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk15FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk15FieldChhmMask)
	}
}

const (
	RegisterOtghshcintmsk15FieldAhberrShift = 2
	RegisterOtghshcintmsk15FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtghshcintmsk15Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtghshcintmsk15Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk15FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk15FieldAhberrMask)
	}
}

const (
	RegisterOtghshcintmsk15FieldStallShift = 3
	RegisterOtghshcintmsk15FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt mask
func (r *registerOtghshcintmsk15Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldStallMask) != 0
}

// SetStall STALL response received interrupt mask
func (r *registerOtghshcintmsk15Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk15FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk15FieldStallMask)
	}
}

const (
	RegisterOtghshcintmsk15FieldNakmShift = 4
	RegisterOtghshcintmsk15FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk15Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtghshcintmsk15Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk15FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk15FieldNakmMask)
	}
}

const (
	RegisterOtghshcintmsk15FieldAckmShift = 5
	RegisterOtghshcintmsk15FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk15Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtghshcintmsk15Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk15FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk15FieldAckmMask)
	}
}

const (
	RegisterOtghshcintmsk15FieldNyetShift = 6
	RegisterOtghshcintmsk15FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtghshcintmsk15Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtghshcintmsk15Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk15FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk15FieldNyetMask)
	}
}

const (
	RegisterOtghshcintmsk15FieldTxerrmShift = 7
	RegisterOtghshcintmsk15FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error
func (r *registerOtghshcintmsk15Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error
func (r *registerOtghshcintmsk15Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk15FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk15FieldTxerrmMask)
	}
}

const (
	RegisterOtghshcintmsk15FieldBberrmShift = 8
	RegisterOtghshcintmsk15FieldBberrmMask  = 0x100
)

// GetBberrm Babble error
func (r *registerOtghshcintmsk15Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldBberrmMask) != 0
}

// SetBberrm Babble error
func (r *registerOtghshcintmsk15Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk15FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk15FieldBberrmMask)
	}
}

const (
	RegisterOtghshcintmsk15FieldFrmormShift = 9
	RegisterOtghshcintmsk15FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk15Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtghshcintmsk15Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk15FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk15FieldFrmormMask)
	}
}

const (
	RegisterOtghshcintmsk15FieldDterrmShift = 10
	RegisterOtghshcintmsk15FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtghshcintmsk15Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtghshcintmsk15Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk15FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk15FieldDterrmMask)
	}
}

// registerOtghshctsiz15Type OTG_HS host channel-15 transfer size register
type registerOtghshctsiz15Type uint32

const (
	RegisterOtghshctsiz15FieldXfrsizShift = 0
	RegisterOtghshctsiz15FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtghshctsiz15Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz15FieldXfrsizMask) >> RegisterOtghshctsiz15FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtghshctsiz15Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz15FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz15FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz15FieldPktcntShift = 19
	RegisterOtghshctsiz15FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtghshctsiz15Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz15FieldPktcntMask) >> RegisterOtghshctsiz15FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtghshctsiz15Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz15FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz15FieldPktcntShift))
}

const (
	RegisterOtghshctsiz15FieldDpidShift = 29
	RegisterOtghshctsiz15FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtghshctsiz15Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz15FieldDpidMask) >> RegisterOtghshctsiz15FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtghshctsiz15Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz15FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz15FieldDpidShift))
}

// registerOtghshcdma15Type OTG_HS host channel-15 DMA address register
type registerOtghshcdma15Type uint32

const (
	RegisterOtghshcdma15FieldDmaaddrShift = 0
	RegisterOtghshcdma15FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtghshcdma15Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma15FieldDmaaddrMask) >> RegisterOtghshcdma15FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtghshcdma15Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma15FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma15FieldDmaaddrShift))
}
